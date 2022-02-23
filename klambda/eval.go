package klambda

import (
	"fmt"
)

type controlFlowKind int

const (
	controlFlowReturn controlFlowKind = iota
	controlFlowApply
)

type inst = func(ctx *ControlFlow, ebp, esp int)

type ControlFlow struct {
	// controlFlowReturn: result = data[0]
	// controlFlowApply: fn, args = data[0], data[1], data[2] ...
	kind controlFlowKind

	// data[pos : len(data)] is the arguments to current function.
	// Why it needs to be a stack?
	// Most of the time, an array is sufficient. But when eval expression like
	// (f (g1 h) (g2 n) ...), the stack is used to store the temporary arguments
	stack []Obj
	pos  int

	val   Obj
	pc    inst
}

func (ctl *ControlFlow) TailApply(f Obj, args ...Obj) {
	ctl.stack = ctl.stack[:ctl.pos]
	ctl.stack = append(ctl.stack, f)
	// TODO: Change the function signature.
	// Using args ... is not good for performance, there is a copy here.
	ctl.stack = append(ctl.stack, args...)
	ctl.kind = controlFlowApply
}

func (ctl *ControlFlow) Return(result Obj) {
	ctl.stack = ctl.stack[:ctl.pos]
	ctl.stack = append(ctl.stack, result)
	ctl.kind = controlFlowReturn
}

func (ctl *ControlFlow) Get(n int) Obj {
	return ctl.stack[ctl.pos+n]
}

// trampoline is introduced for tail call optimization.
func trampoline(ctl *ControlFlow) Obj {
	for ctl.kind != controlFlowReturn {
		switch ctl.kind {
		case controlFlowApply:
			apply(ctl)
		}
	}
	ret := ctl.stack[ctl.pos]
	ctl.stack = ctl.stack[:ctl.pos]
	return ret
}

func makeTempSymbols(n int) []Obj {
	ret := make([]Obj, n)
	for i := 0; i < n; i++ {
		ret[i] = MakeSymbol(fmt.Sprintf("tmp%d", i))
	}
	return ret
}

func apply(ctl *ControlFlow) {
	f := ctl.stack[ctl.pos]
	args := ctl.stack[ctl.pos+1:]

	if *f == scmHeadNative {
		fn := MustNative(f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided == required {
			if len(fn.captured) > 0 {
				// Captured data should come fisrt, then the arguments.
				ctl.stack = append(ctl.stack, make([]Obj, len(fn.captured))...)
				dst := ctl.stack[len(ctl.stack)-len(args):]
				copy(dst, args)
				dst = ctl.stack[ctl.pos+1 : ctl.pos+1+len(fn.captured)]
				copy(dst, fn.captured)
			}
			fn.fn(ctl)
			return
		}

		if provided < required {
			tmp := make([]Obj, 0, fn.require)
			tmp = append(tmp, fn.captured...)
			tmp = append(tmp, args...)
			ctl.Return(MakeNative(fn.fn, fn.require, tmp...))
			return
		}

		// Trick: instead of adjust the arguments, we use a new call stack here.
		savePos := ctl.pos
		ctl.pos = len(ctl.stack)
		ctl.stack = append(ctl.stack, f) // Actually, it's not f any more.
		ctl.stack = append(ctl.stack, fn.captured...)
		taken := required - len(fn.captured)
		ctl.stack = append(ctl.stack, ctl.stack[savePos+1:savePos+1+taken]...)
		fn.fn(ctl)
		f := ctl.stack[ctl.pos]
		ctl.pos = savePos
		args = args[taken:]
		ctl.TailApply(f, args...)
		return
	}
	panic("can't apply object")
}

var symDebug = MakeSymbol("*debug-eval*")

func Eval(ctx *ControlFlow, exp Obj) Obj {
	return ctx.Eval(exp)
}

func (ctx *ControlFlow) Eval(exp Obj) Obj {
	var cc Compiler
	c := cc.compile(exp, nil, true)
	run(ctx, c)
	ctx.stack = ctx.stack[:ctx.pos]
	return ctx.val
}

type Env struct {
	parent *Env
	args   Obj
}

func (env *Env) findVariable(s Obj) (m int, n int) {
	e := env
	for e != nil {
		n = 0
		for args := e.args; args != Nil; args = cdr(args) {
			if car(args) == s {
				return // local variable
			}
			n++
		}
		m = m + 1
		e = e.parent
	}
	if e == nil {
		m = -1 // global variable
	}
	return // closure variable
}

type posRef struct {
	Up     int
	Offset int
}

type Compiler struct{
	freeVars []posRef
}

func (cc *Compiler) compile(exp Obj, env *Env, tail bool) inst {
	switch *exp {
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull:
		return genConstInst(exp)
	case scmHeadSymbol:
		m, n := env.findVariable(exp)
		if m == 0 {
			// local[0] is the closure object, the real args begins from i+1
			return genLocalRefInst(n + 1)
		}
		if m > 0 {
			cc.freeVars = append(cc.freeVars, posRef{Up: m, Offset: n})
			return genClosureRefInst(m, n)
		}

		return genConstInst(exp)
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		tl := pair.cdr // handle special form
		switch pair.car {
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				x := cc.compile(car(tl), env, false)
				y := cc.compile(cadr(tl), env, tail)
				z := cc.compile(caddr(tl), env, tail)
				return genIfInst(x, y, z)
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			x := cc.compile(car(tl), env, false)
			y := cc.compile(cadr(tl), env, tail)
			return genDoInst(x, y, tail)
		case symLambda: // (lambda (x y) x)
			args := Cons(car(tl), Nil)
			body := cadr(tl)
			return cc.compileLambda(args, body, env)
		}
	}

	var cached [5]inst
	insts := cached[:0]
	if IsSymbol(Car(exp)) {
		if fn, ok := klMacro[Car(exp)]; ok {
			return fn(cc, Cdr(exp), env, tail)
		}
		inst := cc.compileSymbolInCall(Car(exp), env, false)
		insts = append(insts, inst)
		exp = Cdr(exp)
	}

	for IsCons(exp) {
		inst := cc.compile(Car(exp), env, false)
		insts = append(insts, inst)
		exp = Cdr(exp)
	}
	return genCallInst(insts, exp, tail)
}

func (cc *Compiler) compileLambda(args, body Obj, env *Env) inst {
	env1 := &Env{
		parent: env,
		args:   args,
	}

	var cc1 Compiler
	op := cc1.compile(body, env1, true)
	nargs := listLength(args)
	collectFVs := cc1.freeVars[:0]
	for _, p := range cc1.freeVars {
		if p.Up == 1 {
			collectFVs = append(collectFVs, p)
		} else {
			// Collect to freeVars and delete from this one.
			cc.freeVars = append(cc.freeVars, posRef{Up: p.Up - 1, Offset: p.Offset})
		}
	}
	return genMakeClosureInst(op, nargs, collectFVs)
}

func (cc *Compiler) compileSymbolInCall(exp Obj, env *Env, tail bool) inst {
	m, n := env.findVariable(exp)
	if m == 0 {
		// local[0] is the closure object, the real args begins from i+1
		return genLocalRefInst(n + 1)
	}
	if m > 0 {
		cc.freeVars = append(cc.freeVars, posRef{Up: m, Offset: n})
		return genClosureRefInst(m, n)
	}
	return genKLGlobalRefInst(exp)
}

func  compileOrMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (or x y) => (if x true (if y true false))
	x := Car(exp)
	y := cadr(exp)
	tmp := Cons(symIf, Cons(y, Cons(True, Cons(False, Nil))))
	exp = Cons(symIf, Cons(x, Cons(True, Cons(tmp, Nil))))
	return cc.compile(exp, env, tail)
}

func  compileAndMacro(cc *Compiler,exp Obj, env *Env, tail bool) inst {
	// (and x y) => (if x (if y true false) false)
	x := Car(exp)
	y := cadr(exp)
	tmp := Cons(symIf, Cons(y, Cons(True, Cons(False, Nil))))
	exp = Cons(symIf, Cons(x, Cons(tmp, Cons(False, Nil))))
	return cc.compile(exp, env, tail)
}

func compileCondMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	exp = compileCondMacroHelp(exp)
	return cc.compile(exp, env, tail)
}

func compileCondMacroHelp(exp Obj) Obj {
	// (cond (x y) ...)
	if exp == Nil {
		return Cons(MakeSymbol("simple-error"), Cons(MakeString("no cond match"), Nil))
	}
	xy := Car(exp)
	x := Car(xy)
	y := cadr(xy)
	tmp := compileCondMacroHelp(Cdr(exp))
	return Cons(symIf, Cons(x, Cons(y, Cons(tmp, Nil))))
}

func compileDefunMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (defun f (x y) ..) => ((fn 'set) (quote f) (lambda (x y) ...))
	fName := Car(exp)
	args := cadr(exp)
	body := caddr(exp)
	inst := cc.compileLambda(args, body, env)
	return genKLGlobalSetInst(fName, inst)
}

func genKLGlobalRefInst(sym Obj) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		s := mustSymbol(sym)
		if s.function == nil {
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + s.str)
		}
		ctx.val = s.function
	}
}

func genKLGlobalSetInst(fName Obj, f inst) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		f(ctx, ebp, esp)
		ctx.val = PrimNS2Set(fName, ctx.val)
	}
}

func  compileLetMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (let x y z) => ((lambda x z) y)
	x := Car(exp)
	y := cadr(exp)
	z := Cdr(Cdr(exp))
	fn := Cons(symLambda, Cons(x, z))
	exp = Cons(fn, Cons(y, Nil))
	return cc.compile(exp, env, tail)
}

func compileFreezeMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (freeze body) => (lambda () body)
	return cc.compileLambda(Nil, Car(exp), env)
}

func compileTrapErrorMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (trap-error body handler) => (try-catch (lambda () body) handler)
	body := Car(exp)
	handler := Cdr(exp)
	action := Cons(symLambda, Cons(Nil, Cons(body, Nil)))
	exp = Cons(MakeSymbol("try-catch"), Cons(action, handler))
	return cc.compile(exp, env, tail)
}

func  compileTypeMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (type exp _) => exp
	return cc.compile(Car(exp), env, tail)
}

func genConstInst(v Obj) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		ctx.val = v
	}
}

func genIfInst(a, b, c inst) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		a(ctx, ebp, esp)
		switch ctx.val {
		case True:
			b(ctx, ebp, esp)
		case False:
			c(ctx, ebp, esp)
		default:
			panic("test branch for if must be a boolean")
		}
	}
}

func genDoInst(op1, op2 inst, tail bool) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		op1(ctx, ebp, esp)
		if tail {
			ctx.pc = op2
			return
		}
		op2(ctx, ebp, esp)
	}
}

func genLocalRefInst(idx int) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		ctx.val = ctx.stack[ebp+idx]
	}
}

func genClosureRefInst(m, n int) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		clo := mustClosure(ctx.stack[ebp])
		for i := 1; i < m; i++ {
			clo = clo.parent
		}
		ctx.val = clo.freeVars[n]
	}
}

func genGlobalRefInst(sym Obj) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		s := mustSymbol(sym)
		if s.cora == nil {
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + s.str)
		}
		ctx.val = s.cora
	}
}

func genMakeClosureInst(op func(ctx *ControlFlow, ebp, esp int), nargs int, mark []posRef) func(ctx *ControlFlow, ebp, esp int) {
	return func(ctx *ControlFlow, ebp, esp int) {
		var parent *scmClosure
		if esp > ebp {
			if *(ctx.stack[ebp]) == scmHeadClosure {
				parent = mustClosure(ctx.stack[ebp])
			}
		}
		// Save some variables into the closure
		// Because they are free variables.
		var freeVars map[int]Obj
		if len(mark) > 0 {
			freeVars = make(map[int]Obj, len(mark))
			for _, p := range mark {
				if p.Up != 1 {
					panic("should never here")
				}
				k := p.Offset
				freeVars[k] = ctx.stack[ebp+k+1]
			}
		}
		ctx.val = MakeClosure(op, ebp, esp, nargs, parent, freeVars)
	}
}

func getRequired(o Obj) int {
	switch *o {
	case scmHeadClosure:
		return mustClosure(o).params
	case scmHeadCurry:
		return mustCurry(o).params
	case scmHeadNative:
		n := MustNative(o)
		return n.require - len(n.captured)
	}
	panic("not implemented" + ObjString(o))
}

func recoverCall(ctx *ControlFlow, saveSP int) {
	ctx.stack = ctx.stack[:ctx.pos]
	ctx.pos = saveSP
}


func Call(ctx *ControlFlow, f Obj, args ...Obj) Obj {
	return ctx.Call(f, args...)
}

func (ctx *ControlFlow) Call(f Obj, args ...Obj) Obj {
	saveSP := ctx.pos
	ctx.pos = len(ctx.stack)
	ctx.stack = append(ctx.stack, f)
	for _, arg := range args {
		ctx.stack = append(ctx.stack, arg)
	}

	myCall(ctx)
	recoverCall(ctx, saveSP)
	return ctx.val
}

func genCallInst(insts []inst, debugExp Obj, tail bool) inst {
	return func(ctx *ControlFlow, ebp, esp int) {
		saveSP := ctx.pos
		ctx.pos = len(ctx.stack)
		for _, inst := range insts {
			inst(ctx, ebp, esp)
			ctx.stack = append(ctx.stack, ctx.val)
		}

	again:

		f := ctx.stack[ctx.pos]
		required := getRequired(f)
		provided := len(ctx.stack[ctx.pos+1:])

		if required < provided {
			// Prepare a new stack for calling...
			saveSP1 := ctx.pos
			ctx.pos = len(ctx.stack)
			for i := saveSP1; i < saveSP1+1+required; i++ {
				ctx.stack = append(ctx.stack, ctx.stack[i])
			}
			// Make the call
			myCall(ctx)
			// Recover the stack
			ctx.stack[saveSP1] = ctx.val
			to := saveSP1 + 1
			from := saveSP1 + 1 + required
			for from < saveSP1+1+provided {
				ctx.stack[to] = ctx.stack[from]
				from++
				to++
			}
			ctx.stack = ctx.stack[:to]
			ctx.pos = saveSP1

			goto again
		}

		if required > provided {
			upvalues := make([]Obj, provided)
			copy(upvalues, ctx.stack[ctx.pos+1:])
			ctx.val = MakeCurry(f, upvalues, required-provided)
			recoverCall(ctx, saveSP)
			return
		}

		if required == provided {
			if tail {
				f := ctx.stack[ctx.pos]
				if *f == scmHeadClosure {
					clo := mustClosure(f)
					nargs := len(ctx.stack[ctx.pos:])
					copy(ctx.stack[saveSP:saveSP+nargs], ctx.stack[ctx.pos:])
					ctx.stack = ctx.stack[:saveSP+nargs]
					ctx.pos = saveSP
					ctx.pc = clo.code
					return
				}
			}

			myCall(ctx)
			recoverCall(ctx, saveSP)
		}
	}
}

func run(ctx *ControlFlow, code inst) {
	ctx.pc = code
	for ctx.pc != nil {
		tmp := ctx.pc
		ctx.pc = nil
		tmp(ctx, ctx.pos, len(ctx.stack))
	}
}

func myCall(ctx *ControlFlow) {
retry:
	f := ctx.stack[ctx.pos]
	switch *f {

	case scmHeadNative:
		fn := MustNative(f)
		ctx.kind = controlFlowApply
		fn.fn(ctx)
		ctx.val = ctx.Get(0)

	case scmHeadClosure:
		clo := mustClosure(f)
		run(ctx, clo.code)

	case scmHeadCurry:
		curry := mustCurry(f)
		saved := append([]Obj{}, ctx.stack[ctx.pos+1:]...)
		ctx.stack = ctx.stack[:ctx.pos]
		ctx.stack = append(ctx.stack, curry.origin)
		ctx.stack = append(ctx.stack, curry.upvalue...)
		ctx.stack = append(ctx.stack, saved...)
		f = ctx.stack[ctx.pos]
		goto retry
	default:
		panic("invalid call operation on:" + ObjString(f))
	}
}
