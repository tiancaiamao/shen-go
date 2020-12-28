package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

var __initExprs []kl.Obj

func regist(e *kl.KLambda) {
	for _, expr := range __initExprs {
		kl.Call(e, expr)
	}
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e kl.KLambda
	regist(&e)

	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	e.Override()
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.repl"), kl.Nil))
	return
}
