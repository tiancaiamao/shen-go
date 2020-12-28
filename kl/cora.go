package kl

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"plugin"
)

type Cora struct {
	ControlFlow
}

func primLoadSo(e Evaluator) {
	pluginPath := GetString(e.Get(1))
	p, err := plugin.Open(pluginPath)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}

	entry, err := p.Lookup("Main")
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}

	f, ok := entry.(func(__e Evaluator, __args ...Obj))
	if !ok {
		e.Return(MakeError("plugin Main should be func(__e Evaluator, __args ...Obj)"))
		return
	}

	res := Call(e, MakeNative(f, 0))
	e.Return(res)
	return
}

func readFileAsSexp(e Evaluator) {
	filePath := mustString(e.Get(1))
	f, err := os.Open(filePath)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()

	ret := Nil
	r := NewSexpReader(f, true)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				e.Return(MakeError(err.Error()))
				return
			}
			break
		}
		ret = cons(exp, ret)
	}
	e.Return(reverse(ret))
	return
}

func writeSexpToFile(e Evaluator) {
	filePath := mustString(e.Get(1))
	str := ObjString(e.Get(2))
	err := ioutil.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	e.Return(Nil)
	return
}

func (e *Cora) eval() {
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
