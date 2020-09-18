package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

type Evaluator struct {
	ControlFlow
	Silence bool
}

func NewEvaluator() *Evaluator {
	var e Evaluator
	for _, prim := range AllPrimitives {
		sym := MakeSymbol(prim.Name)
		BindSymbolFunc(sym, Obj(&prim.scmHead))
	}
	// Overload for primitive set and value.
	tmp := &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2, Function: e.primSet}
	BindSymbolFunc(MakeSymbol("set"), Obj(&tmp.scmHead))
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1, Function: e.primValue}
	BindSymbolFunc(MakeSymbol("value"), Obj(&tmp.scmHead))
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1, Function: e.primEvalKL}
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

func (e *Evaluator) primSet(args ...Obj) Obj {
	return PrimSet(args[0], args[1])
}

func (e *Evaluator) primValue(args ...Obj) Obj {
	return PrimValue(args[0])
}

func (e *Evaluator) primEvalKL(args ...Obj) Obj {
	// fmt.Println("eval-kl: ", ObjString(args[0]))
	return e.evalExp(args[0], Nil)
}

func (e *Evaluator) primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return e.LoadFile(path)
}

func (e *Evaluator) LoadFile(file string) Obj {
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return MakeError(err.Error())
			}
			break
		}

		res := e.evalExp(exp, Nil)
		if *res == scmHeadError {
			return res
		}
	}
	if !e.Silence {
		fmt.Println(filePath)
	}
	return MakeString(file)
}

func (e *Evaluator) Eval(exp Obj) (res Obj) {
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
	res = e.evalExp(exp, Nil)
	return
}

func (e *Evaluator) Call(f Obj, args ...Obj) Obj {
	e.TailApply(f, args...)
	return e.trampoline()
}

func (e *Evaluator) RegistNativeCall(name string, f Obj) {
	_ = MustNative(f)
	BindSymbolFunc(MakeSymbol(name), f)
}

func (e *Evaluator) BootstrapShen() {
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
