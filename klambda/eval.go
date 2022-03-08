package klambda

import (
	"fmt"
)

type controlFlowKind int

const (
	controlFlowReturn controlFlowKind = iota
	controlFlowApply
)

type ControlFlow struct {
	// controlFlowReturn: result = data[0]
	// controlFlowApply: fn, args = data[0], data[1], data[2] ...
	kind controlFlowKind

	// data[pos : len(data)] is the arguments to current function.
	// Why it needs to be a stack?
	// Most of the time, an array is sufficient. But when eval expression like
	// (f (g1 h) (g2 n) ...), the stack is used to store the temporary arguments
	stack []Obj
	pos   int

	val Obj
	pc  inst
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("eval error ==", ObjString(exp))
			panic(r)
		}
	}()
	var cc Compiler
	env := &Env{
		parent: nil,
		args:   Nil,
	}
	c := cc.compile(exp, env, true)
	if cc.local > 0 {
		// Special handle for top level (let x v ...) expression
		ctx.stack = append(ctx.stack, Nil) // Dummy closure
		for i := 0; i < cc.local; i++ {    // Reserve space for local variables
			ctx.stack = append(ctx.stack, Nil)
		}
	}
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
		n = -1
		idx := 0
		for args := e.args; args != Nil; args = cdr(args) {
			if car(args) == s {
				n = idx
				// return // local variable
			}
			// n++
			idx++
		}
		if n >= 0 {
			return
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

type Compiler struct {
	freeVars []posRef
	local    int
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
		case symLambda: // (lambda X X)
			args := Nil
			if car(tl) != Nil {
				args = Cons(car(tl), args)
			}
			body := cadr(tl)
			return cc.compileLambda(args, body, env)
		case symLet: // (let a b c)
			val := cc.compile(cadr(tl), env, false)
			// Actually, the compile Env of the body is *not* changed!
			// fmt.Println("before ... and args ===", ObjString(env.args))
			env1 := &Env{
				parent: env.parent,
				args:   listAppend(env.args, cons(car(tl), Nil)), // cons(car(tl), env.args),
			}
			// fmt.Println("compile let body env ==", ObjString(env1.args))
			cc.local++
			body := caddr(tl)
			op := cc.compile(body, env1, tail)
			idx := listLength(env.args)
			return genLetInst(val, op, idx)
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
	return genMakeClosureInst(op, nargs, collectFVs, cc1.local)
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

func compileOrMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (or x y) => (if x true (if y true false))
	x := Car(exp)
	y := cadr(exp)
	tmp := Cons(symIf, Cons(y, Cons(True, Cons(False, Nil))))
	exp = Cons(symIf, Cons(x, Cons(True, Cons(tmp, Nil))))
	return cc.compile(exp, env, tail)
}

func compileAndMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
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

func compileLetMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
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
	// fmt.Println("after rewrite trap get ===", ObjString(exp))
	return cc.compile(exp, env, tail)
}

func compileTypeMacro(cc *Compiler, exp Obj, env *Env, tail bool) inst {
	// (type exp _) => exp
	return cc.compile(Car(exp), env, tail)
}
