package kl

import (
	"fmt"
	"runtime"
)

type Cora struct {
	ControlFlow
}

func (c *Cora) Call(f Obj, args ...Obj) Obj {
	c.TailApply(f, args...)
	return c.trampoline(c)
}

func (e *Cora) Try(f Obj) (res tryResult) {
	panic("not implement")
}

func (e *Cora) Eval(exp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Panic:", r)
			fmt.Println("Recovered in Eval:", ObjString(exp))
			fmt.Println(string(buf[:n]))
			res = Nil
		}
	}()
	res = e.evalExp(e, exp, Nil)
	return
}

func NewCora() *Cora {
	var e Cora
	for _, prim := range AllPrimitives {
		sym := MakeSymbol(prim.Name)
		CoraSet(sym, Obj(&prim.scmHead))
	}
	CoraSet(MakeSymbol("car"), CoraValue(MakeSymbol("hd")))
	CoraSet(MakeSymbol("cdr"), CoraValue(MakeSymbol("tl")))
	priv := &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "gensym", Required: 1, Function: PrimGenSym}
	CoraSet(MakeSymbol("gensym"), Obj(&priv.scmHead))
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2, Function: CoraSet}
	CoraSet(MakeSymbol("set"), Obj(&priv.scmHead))
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load", Required: 1, Function: e.primLoadFile}
	CoraSet(MakeSymbol("load"), Obj(&priv.scmHead))
	return &e
}

func (e *Cora) primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return loadFile(&e.ControlFlow, e, true, path)
}

func (e *Cora) eval(ctl *ControlFlow) {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	// fmt.Println("eval === exp:", ObjString(exp))

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure, scmHeadPrimitive:
		e.Return(exp)
		return
	case scmHeadSymbol:
		if val, ok := envGet(env, exp); ok {
			e.Return(val)
			return
		}
		sym := mustSymbol(exp)
		if sym.cora == nil {
			// TODO: xxx
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + sym.str)
		}
		e.Return(sym.cora)
		return
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		tl := pair.cdr // handle special form
		switch pair.car {
		case symQuote:
			// Extension to make vm work.
			// TODO: remove it later
			e.Return(car(tl))
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				ctl.evalIf(e, car(tl), cadr(tl), caddr(tl), env)
				return
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			ctl.evalExp(e, car(tl), env)
			e.TailEval(cadr(tl), env)
			return
		case symLambda: // (lambda (x y) x)
			e.Return(makeProcedure(car(tl), cadr(tl), env))
			return
		}
	}

	savePOS := e.pos
	args := exp
	for *args == scmHeadPair {
		v := ctl.evalExp(e, car(args), env)
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
	return
}
