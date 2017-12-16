package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

type Evaluator struct {
	symbolTable   map[string]Obj
	functionTable map[string]Obj
	Silence       bool

	nativeFunc map[string]*ScmPrimitive
}

func NewEvaluator() *Evaluator {
	var e Evaluator
	e.symbolTable = make(map[string]Obj)

	e.functionTable = make(map[string]Obj, len(allPrimitives))
	for _, prim := range allPrimitives {
		e.functionTable[prim.Name] = Obj(&prim.scmHead)
	}
	// Overload for primitive set and value.
	tmp := &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2, Function: e.primSet}
	e.functionTable["set"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1, Function: e.primValue}
	e.functionTable["value"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1, Function: e.primEvalKL}
	e.functionTable["eval-kl"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1, Function: e.primLoadFile}
	e.functionTable["load-file"] = Obj(&tmp.scmHead)

	e.nativeFunc = make(map[string]*ScmPrimitive)
	e.RegistNativeCall(MakePrimitive("primitive?", 1, NativeIsPrimitive))
	e.RegistNativeCall(MakePrimitive("primitive-arity", 1, NativePrimitiveArity))
	e.RegistNativeCall(MakePrimitive("primitive-id", 1, NativePrimitiveID))

	e.symbolTable["*stinput*"] = MakeStream(os.Stdin)
	e.symbolTable["*stoutput*"] = MakeStream(os.Stdout)
	dir, _ := os.Getwd()
	e.symbolTable["*home-directory*"] = MakeString(dir)
	e.symbolTable["*language*"] = MakeString("Go")
	e.symbolTable["*implementation*"] = MakeString("interpreter")
	e.symbolTable["*relase*"] = MakeString(runtime.Version())
	e.symbolTable["*os*"] = MakeString(runtime.GOOS)
	e.symbolTable["*porters*"] = MakeString("Arthur Mao")
	e.symbolTable["*port*"] = MakeString("0.0.1")

	// Extended by shen-go implementation
	e.symbolTable["*package-path*"] = MakeString(PackagePath())
	return &e
}

func (e *Evaluator) primSet(args ...Obj) Obj {
	return PrimSet(e.symbolTable, args[0], args[1])
}

func (e *Evaluator) primValue(args ...Obj) Obj {
	return PrimValue(e.symbolTable, args[0])
}

func (e *Evaluator) primEvalKL(args ...Obj) Obj {
	return e.trampoline(args[0], nil)
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

		res := e.trampoline(exp, nil)
		if *res == scmHeadError {
			return res
		}
		if !e.Silence {
			fmt.Printf("%#v\n", (*scmHead)(res))
		}
	}
	return Nil
}

func (e *Evaluator) Eval(exp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Recovered in Eval:", ObjString(exp))
			fmt.Println(string(buf[:n]))
			res = Nil
		}
	}()
	res = e.trampoline(exp, nil)
	return
}

func (e *Evaluator) RegistNativeCall(prim *ScmPrimitive) {
	e.nativeFunc[prim.Name] = prim
}

func (e *Evaluator) BootstrapShen() {
	e.LoadFile("ShenOSKernel-20.1/klambda/toplevel.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/core.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/sys.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/sequent.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/yacc.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/reader.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/prolog.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/track.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/load.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/writer.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/macros.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/declarations.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/t-star.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/types.kl")

	// override
	isSymbol := MakePrimitive("symbol?", 1, primIsSymbol)
	e.functionTable["symbol?"] = Obj(&isSymbol.scmHead)
}
