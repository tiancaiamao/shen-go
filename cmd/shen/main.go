package main

import (
	"flag"
	// "fmt"
	// "io"
	"net/http"
	_ "net/http/pprof"
	// "os"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

var __initExprs []kl.Obj

func Regist(e *kl.KLambda) {
	for _, expr := range __initExprs {
		kl.Call(e, expr)
	}
}

// func Regist(e *kl.KLambda) {
// 	kl.Call(e, kl.MakeNative(MainToplevel, 0))
// 	kl.Call(e, kl.MakeNative(MainDict, 0))
// 	kl.Call(e, kl.MakeNative(MainCore, 0))
// 	kl.Call(e, kl.MakeNative(MainSys, 0))
// 	kl.Call(e, kl.MakeNative(MainSequent, 0))
// 	kl.Call(e, kl.MakeNative(MainYacc, 0))
// 	kl.Call(e, kl.MakeNative(MainReader, 0))
// 	kl.Call(e, kl.MakeNative(MainProlog, 0))
// 	kl.Call(e, kl.MakeNative(MainTrack, 0))
// 	kl.Call(e, kl.MakeNative(MainLoad, 0))
// 	kl.Call(e, kl.MakeNative(MainWriter, 0))
// 	kl.Call(e, kl.MakeNative(MainMacros, 0))
// 	kl.Call(e, kl.MakeNative(MainDeclarations, 0))
// 	kl.Call(e, kl.MakeNative(MainTStar, 0))
// 	kl.Call(e, kl.MakeNative(MainTypes, 0))
// 	kl.Call(e, kl.MakeNative(MainInit, 0))
// 	kl.Call(e, kl.MakeNative(MainExtensionFeatures, 0))
// 	kl.Call(e, kl.MakeNative(MainExtensionLauncher, 0))
// 	kl.Call(e, kl.MakeNative(MainExtensionFactoriseDefun, 0))
// 	// kl.Call(e, kl.MakeNative(MainExtensionProgrammablePatternMatching, 0))
// }

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	e := kl.NewKLambda()
	Regist(e)

	kl.Eval(e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	e.Override()
	kl.Eval(e, kl.Cons(kl.MakeSymbol("shen.repl"), kl.Nil))

	// r := kl.NewSexpReader(os.Stdin, false)
	// for i := 0; ; i++ {
	// 	fmt.Printf("%d #> ", i)
	// 	sexp, err := r.Read()
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println("read error:", err)
	// 		}
	// 		break
	// 	}

	// 	res := kl.Eval(e, sexp)
	// 	fmt.Println(kl.ObjString(res))
	// }

	return
}

// func Regist(e *kl.KLambda) {
// 	for _, def := range __initDefs {
// 		if def.name == "symbol?" {
// 			continue
// 		}
// 		e.RegistNativeCall(def.name, def.value)
// 	}

// 	for _, expr := range __initExprs {
// 		kl.Call(e, expr)
// 	}
// }
