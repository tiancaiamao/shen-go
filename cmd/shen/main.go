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

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	e := kl.NewEvaluator()
	Regist(e)

	e.Eval(kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
	return
}
