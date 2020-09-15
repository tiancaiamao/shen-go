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
	functionTable map[string]Obj
	Silence       bool

	nativeFunc map[string]Obj
}

func NewEvaluator() *Evaluator {
	var e Evaluator

	e.functionTable = make(map[string]Obj, len(AllPrimitives))
	for _, prim := range AllPrimitives {
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

	e.nativeFunc = make(map[string]Obj)

	PrimSet(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	PrimSet(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	PrimSet(MakeSymbol("*home-directory*"), MakeString(dir))
	PrimSet(MakeSymbol("*language*"), MakeString("Go"))
	PrimSet(MakeSymbol("*implementation*"), MakeString("fusion"))
	PrimSet(MakeSymbol("*relase*"), MakeString(runtime.Version()))
	PrimSet(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	PrimSet(MakeSymbol("*porters*"), MakeString("Arthur Mao"))
	PrimSet(MakeSymbol("*port*"), MakeString("1.0.0"))

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

		res := e.evalExp(exp, nil)
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
	// ctl.stack[bp : len(ctl.stack)] is current stack frame
	// 
	// ---------------------------
	//  | bp            |
	//  | arg1 arg2 ... |
	// ---------------------------
	//
	// save old bp, update bp to len(ctl.stack)
	// push arg1, arg2 ...
	// -----------------------------------
	//  | saveBP        | bp           |
	//  | ...           | arg1 arg2 ...|
	// -----------------------------------
	// 
	// Call()
	// recover the old frame: stack <- stack[:bp], bp <-saveBP

	ctl := &e.ControlFlow
	saveBP := ctl.bp

	ctl.bp = len(ctl.stack)
	ctl.stack = append(ctl.stack, f)
	ctl.stack = append(ctl.stack, args...)
	ctl.kind = ControlFlowApply
	ret := e.trampoline()

	ctl.stack = ctl.stack[:ctl.bp]
	ctl.bp = saveBP
	return ret
}

func (e *Evaluator) RegistNativeCall(name string, f Obj) {
	_ = MustNative(f)
	e.nativeFunc[name] = f
}

func (e *Evaluator) BootstrapShen() {
	e.LoadFile("ShenOSKernel-22.2/klambda/toplevel.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/dict.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/core.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/sys.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/sequent.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/yacc.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/reader.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/prolog.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/track.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/load.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/writer.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/macros.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/declarations.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/t-star.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/types.kl")
	e.LoadFile("ShenOSKernel-22.2/klambda/init.kl")

	// override
	isSymbol := MakePrimitive("symbol?", 1, PrimIsSymbol)
	e.functionTable["symbol?"] = Obj(&isSymbol.scmHead)
}
