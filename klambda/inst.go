package klambda

import (
	"fmt"
)

var _ inst = constInst{}
var _ inst = &ifInst{}
var _ inst = &doInst{}
var _ inst = localRefInst{}
var _ inst = closureRefInst{}
var _ inst = globalRefInst{}
var _ inst = klGlobalRefInst{}
var _ inst = letInst{}
var _ inst = &makeClosureInst{}
var _ inst = &callInst{}

type inst interface {
	Exec(ctx *ControlFlow, ebp, esp int)
	// TODO: String()  print out the struct for debugging
	GenGoCode(gen *GoCodeGenerator) int
	// TODO: GenCCode()   generate C code
}

type constInst struct {
	val Obj
}

func (v constInst) Exec(ctx *ControlFlow, ebp, esp int) {
	ctx.val = v.val
	ctx.pc = nil
}

type ifInst struct {
	a inst
	b inst
	c inst
}

func (i *ifInst) Exec(ctx *ControlFlow, ebp, esp int) {
	for ctx.pc = i.a; ctx.pc != nil; {
		ctx.pc.Exec(ctx, ebp, esp)
	}

	switch ctx.val {
	case True:
		ctx.pc = i.b
	case False:
		ctx.pc = i.c
	default:
		panic("test branch for if must be a boolean")
	}
}

type doInst struct {
	op1  inst
	op2  inst
	tail bool
}

func (d *doInst) Exec(ctx *ControlFlow, ebp, esp int) {
	for ctx.pc = d.op1; ctx.pc != nil; {
		ctx.pc.Exec(ctx, ebp, esp)
	}
	ctx.pc = d.op2
}

type localRefInst struct {
	int
	isLetVar bool
}

func (l localRefInst) Exec(ctx *ControlFlow, ebp, esp int) {
	ctx.val = ctx.stack[ebp+l.int]
	ctx.pc = nil
}

type closureRefInst struct {
	m int
	n int
}

func (c closureRefInst) Exec(ctx *ControlFlow, ebp, esp int) {
	clo := mustClosure(ctx.stack[ebp])
	for i := 1; i < c.m; i++ {
		clo = clo.parent
	}
	ctx.val = clo.freeVars[c.n]
	ctx.pc = nil
}

type globalRefInst struct {
	sym Obj
}

func (g globalRefInst) Exec(ctx *ControlFlow, ebp, esp int) {
	s := mustSymbol(g.sym)
	if s.cora == nil {
		fmt.Println("NOT DEFINED")
		panic("undefined symbol:" + s.str)
	}
	ctx.val = s.cora
	ctx.pc = nil
}

type klGlobalRefInst struct {
	sym Obj
}

func (ref klGlobalRefInst) Exec(ctx *ControlFlow, ebp, esp int) {
	s := mustSymbol(ref.sym)
	if s.function == nil {
		fmt.Println("NOT DEFINED")
		panic("undefined symbol:" + s.str)
	}
	ctx.val = s.function
	ctx.pc = nil
}

type letInst struct {
	arg  inst
	body inst
	idx  int
}

func (l letInst) Exec(ctx *ControlFlow, ebp, esp int) {
	for ctx.pc = l.arg; ctx.pc != nil; {
		ctx.pc.Exec(ctx, ebp, esp)
	}

	ctx.stack[1+ebp+l.idx] = ctx.val
	ctx.pc = l.body
}

type makeClosureInst struct {
	op     inst
	nargs  int
	mark   []posRef
	nlocal int
}

func (m *makeClosureInst) Exec(ctx *ControlFlow, ebp, esp int) {
	var parent *scmClosure
	if esp > ebp {
		if *(ctx.stack[ebp]) == scmHeadClosure {
			parent = mustClosure(ctx.stack[ebp])
		}
	}
	// Save some variables into the closure
	// Because they are free variables.
	var freeVars map[int]Obj
	if len(m.mark) > 0 {
		freeVars = make(map[int]Obj, len(m.mark))
		for _, p := range m.mark {
			if p.Up != 1 {
				panic("should never here")
			}
			k := p.Offset
			freeVars[k] = ctx.stack[ebp+k+1]
		}
	}
	ctx.val = MakeClosure(m.op, ebp, esp, m.nargs, parent, freeVars, m.nlocal)
	ctx.pc = nil
}

type callInst struct {
	insts    []inst
	debugExp Obj
	tail     bool
}

func (c *callInst) Exec(ctx *ControlFlow, ebp, esp int) {
	saveSP := ctx.pos
	ctx.pos = len(ctx.stack)
	for _, inst := range c.insts {
		for ctx.pc = inst; ctx.pc != nil; {
			ctx.pc.Exec(ctx, ebp, esp)
		}
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
		if c.tail && *f == scmHeadClosure {
			clo := mustClosure(f)
			nargs := len(ctx.stack[ctx.pos:])
			copy(ctx.stack[saveSP:saveSP+nargs], ctx.stack[ctx.pos:])
			ctx.stack = ctx.stack[:saveSP+nargs]
			if clo.nlocal > 0 {
				for i := 0; i < clo.nlocal; i++ {
					ctx.stack = append(ctx.stack, Nil)
				}
			}
			// fmt.Println(".... in tail call ", clo.nlocal, ctx.stack)
			ctx.pos = saveSP
			ctx.pc = clo.code
			return
		}

		myCall(ctx)
		recoverCall(ctx, saveSP)
	}
}

type klGlobalSetInst struct {
	name Obj
	f    inst
}

func (s klGlobalSetInst) Exec(ctx *ControlFlow, ebp, esp int) {
	for ctx.pc = s.f; ctx.pc != nil; {
		ctx.pc.Exec(ctx, ebp, esp)
	}

	ctx.val = PrimNS2Set(s.name, ctx.val)
	ctx.pc = nil
}

func genConstInst(v Obj) inst {
	return constInst{v}
}

func genIfInst(a, b, c inst) inst {
	return &ifInst{a, b, c}
}

func genDoInst(op1, op2 inst, tail bool) inst {
	return &doInst{op1, op2, tail}
}

func genLocalRefInst(idx int, isLetVar bool) inst {
	return localRefInst{idx, isLetVar}
}

func genClosureRefInst(m, n int) inst {
	return closureRefInst{m, n}
}

func genGlobalRefInst(sym Obj) inst {
	return globalRefInst{sym}
}

func genLetInst(arg, body inst, idx int) inst {
	return letInst{arg, body, idx}
}

func genMakeClosureInst(op inst, nargs int, mark []posRef, nlocal int) inst {
	return &makeClosureInst{op, nargs, mark, nlocal}
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
	return &callInst{insts, debugExp, tail}
}

func genKLGlobalRefInst(sym Obj) inst {
	return klGlobalRefInst{sym}
}

func genKLGlobalSetInst(fName Obj, f inst) inst {
	return klGlobalSetInst{fName, f}
}

func myCall(ctx *ControlFlow) {
retry:
	f := ctx.stack[ctx.pos]
	switch *f {

	case scmHeadNative:
		// fn := MustNative(f)
		// fn.fn(ctx)

		ctx.kind = controlFlowApply
		ctx.val = trampoline(ctx)

	case scmHeadClosure:
		clo := mustClosure(f)
		if clo.nlocal > 0 {
			for i := 0; i < clo.nlocal; i++ {
				ctx.stack = append(ctx.stack, Nil)
			}
			// fmt.Println(".... in call ", clo.nlocal, ctx.stack)
		}

		for ctx.pc = clo.code; ctx.pc != nil; {
			ctx.pc.Exec(ctx, ctx.pos, len(ctx.stack))
		}

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
