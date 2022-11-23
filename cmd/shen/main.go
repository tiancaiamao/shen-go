package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func regist(e *kl.ControlFlow) {
	for _, init := range []kl.Obj{
		TopLevelMain,
		CoreMain,
		SysMain,
		SequentMain,
		YaccMain,
		ReaderMain,
		PrologMain,
		TrackMain,
		LoadMain,
		WriterMain,
		MacrosMain,
		DeclarationsMain,
		TStarMain,
		TypesMain,
	} {
		res := kl.Call(e, init)
		if kl.IsError(res) {
			fmt.Println("load ...fail")
		}
	}
}

var ns2_1set kl.Obj
var try_1catch kl.Obj

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
	try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))

	var e kl.ControlFlow
	regist(&e)
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
}
