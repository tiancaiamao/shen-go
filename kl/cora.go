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
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load-so", Required: 1, Function: e.primLoadSo}
	CoraSet(MakeSymbol("load-so"), Obj(&priv.scmHead))
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-sexp", Required: 1, Function: readFileAsSexp}
	CoraSet(MakeSymbol("read-file-as-sexp"), Obj(&priv.scmHead))
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "write-sexp-to-file", Required: 2, Function: writeSexpToFile}
	CoraSet(MakeSymbol("write-sexp-to-file"), Obj(&priv.scmHead))
	priv = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "apply", Required: 2, Function: e.primApply}
	CoraSet(MakeSymbol("apply"), Obj(&priv.scmHead))
	return &e
}

func (e *Cora) primApply(args ...Obj) Obj {
	fn := args[0]
	l := args[1]
	objs := ListToSlice(l)
	return Call(e, fn, objs...)
}

func (e *Cora) primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return loadFile(e, true, path)
}

func (e *Cora) primLoadSo(args ...Obj) Obj {
	pluginPath := GetString(args[0])
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return MakeError(err.Error())
	}

	entry, err := p.Lookup("Main")
	if err != nil {
		return MakeError(err.Error())
	}

	f, ok := entry.(func(__e Evaluator, __args ...Obj))
	if !ok {
		return MakeError("plugin Main should be func(__e Evaluator, __args ...Obj)")
	}

	res := Call(e, MakeNative(f, 0))

	return res
}

func readFileAsSexp(args ...Obj) Obj {
	filePath := mustString(args[0])
	f, err := os.Open(filePath)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()

	ret := Nil
	r := NewSexpReader(f, true)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return MakeError(err.Error())
			}
			break
		}
		ret = cons(exp, ret)
	}
	return reverse(ret)
}

func writeSexpToFile(args ...Obj) Obj {
	filePath := mustString(args[0])
	str := ObjString(args[1])
	err := ioutil.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		return MakeError(err.Error())
	}
	return Nil
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
