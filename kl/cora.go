package kl

import (
	"fmt"
)

type Cora struct {
	ControlFlow
}

func (e *Cora) eval() {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	// fmt.Println("eval === exp:", ObjString(exp))

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure /* , scmHeadPrimitive */ :
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
			e.Return(car(tl))
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(e, car(tl), cadr(tl), caddr(tl), env)
				return
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			evalExp(e, car(tl), env)
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
		v := evalExp(e, car(args), env)
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
	return
}


func (e *Cora) Global(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.cora != nil {
		return sym.cora
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}
