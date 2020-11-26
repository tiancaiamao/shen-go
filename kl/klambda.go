package kl

import (
	"fmt"
	"os"
	"runtime"
	// "runtime/debug"
)

type KLambda struct {
	ControlFlow
}

func NewKLambda() *KLambda {
	var e KLambda
	for _, prim := range AllPrimitives {
		sym := MakeSymbol(prim.Name)
		BindSymbolFunc(sym, Obj(&prim.scmHead))
	}
	// Overload for primitive set and value.
	tmp := &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1, Function: e.primEvalKL}
	BindSymbolFunc(MakeSymbol("eval-kl"), Obj(&tmp.scmHead))
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1, Function: e.primLoadFile}
	BindSymbolFunc(MakeSymbol("load-file"), Obj(&tmp.scmHead))

	PrimSet(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	PrimSet(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	PrimSet(MakeSymbol("*home-directory*"), MakeString(dir))
	PrimSet(MakeSymbol("*language*"), MakeString("Go"))
	PrimSet(MakeSymbol("*implementation*"), MakeString("AOT+interpreter"))
	PrimSet(MakeSymbol("*relase*"), MakeString(runtime.Version()))
	PrimSet(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	PrimSet(MakeSymbol("*porters*"), MakeString("Arthur Mao"))
	PrimSet(MakeSymbol("*port*"), MakeString("1.0.0-rc1"))

	// Extended by shen-go implementation
	PrimSet(MakeSymbol("*package-path*"), MakeString(PackagePath()))
	return &e
}

func (e *KLambda) primEvalKL(args ...Obj) Obj {
	return evalExp(e, args[0], Nil)
}

func (e *KLambda) primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return loadFile(e, false, path)
}

func (e *KLambda) LoadFile(file string) Obj {
	return loadFile(e, false, file)
}

func (e *KLambda) eval() {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	// fmt.Println("eval exp = ", ObjString(exp))
	// debug.PrintStack()

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure, scmHeadPrimitive:
		e.Return(exp)
		return
	case scmHeadSymbol:
		// fmt.Println("eval symbol ====", ObjString(exp))
		if val, ok := envGet(env, exp); ok {
			exp = val
		}
		e.Return(exp)
		return
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		exp = pair.cdr // handle special form
		switch pair.car {
		case symDefun: // (defun f (x y) z)
			proc := makeProcedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			BindSymbolFunc(funName, proc)
			e.Return(funName)
			return
		case symLambda: // (lambda x x)
			e.Return(makeProcedure(car(exp), cadr(exp), env))
			return
		case symFreeze: // (freeze body)
			e.Return(makeProcedure(Nil, car(exp), env))
			return
		case symLet: // (let x y z)
			args := evalExp(e, cadr(exp), env)
			newEnv := envExtend(env, []Obj{car(exp)}, []Obj{args})
			e.TailEval(caddr(exp), newEnv)
			return
		case symAnd:
			evalAnd(e, car(exp), cadr(exp), env)
			return
		case symOr:
			evalOr(e, car(exp), cadr(exp), env)
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(e, car(exp), cadr(exp), caddr(exp), env)
				return
			} // if may also be a function for partial apply
		case symCond: // (cond (false 1) (true 2))
			e.evalCond(exp, env)
			return
		case symTrapError: // (trap-error ~body ~handler)
			e.evalTrapError(exp, env)
			return
		case symDo: // (do A A)
			evalExp(e, car(exp), env)
			e.TailEval(cadr(exp), env)
			return
		}
	}

	savePOS := e.pos
	fn := e.evalFunction(pair.car, env)
	e.data = e.data[:e.pos]
	e.data = append(e.data, fn)
	e.pos++

	args := pair.cdr
	for *args == scmHeadPair {
		v := evalExp(e, car(args), env)
		if *v == scmHeadError {
			e.pos = savePOS
			e.TailApply(fn, v)
			return
		}
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
	return
}

func (e *KLambda) evalCond(l Obj, env Obj) {
	for *l == scmHeadPair {
		curr := car(l)
		if evalExp(e, car(curr), env) == True {
			e.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	e.Return(Nil)
	return
}

func (e *KLambda) evalTrapError(exp Obj, env Obj) {
	savePOS := e.pos
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					e.pos = savePOS
					e.data = e.data[:e.pos]
					handle := e.evalFunction(cadr(exp), env)
					e.TailApply(handle, val)
					return
				}
			}
			// Unexpected panic?
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Panic:", err)
			fmt.Println("Recovered in trap-error:", ObjString(exp))
			fmt.Println(string(buf[:n]))
		}
	}()
	v := evalExp(e, car(exp), env)
	e.Return(v)
	return
}

func (e *KLambda) evalFunction(fn Obj, env Obj) Obj {
	if ok, sym := isSymbol(fn); ok {
		if sym.function != nil {
			return sym.function
		}

		if proc, ok := envGet(env, fn); ok {
			return proc
		}
	}

	switch *fn {
	case scmHeadPrimitive, scmHeadProcedure, scmHeadNative:
		return fn
	case scmHeadPair:
		return evalExp(e, fn, env)
	}
	panic(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func (e *KLambda) RegistNativeCall(name string, f Obj) {
	_ = MustNative(f)
	BindSymbolFunc(MakeSymbol(name), f)
}

func (e *KLambda) BootstrapShen() {
	e.LoadFile("shen-sources-shen-22.3/klambda/toplevel.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/dict.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/core.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/sys.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/sequent.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/yacc.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/reader.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/prolog.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/track.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/load.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/writer.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/macros.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/declarations.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/t-star.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/types.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/init.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/extension-features.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/extension-launcher.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/extension-factorise-defun.kl")
	e.LoadFile("shen-sources-shen-22.3/klambda/extension-programmable-pattern-matching.kl")

	// override
	isSymbol := MakePrimitive("symbol?", 1, PrimIsSymbol)
	BindSymbolFunc(MakeSymbol("symbol?"), Obj(&isSymbol.scmHead))
}
