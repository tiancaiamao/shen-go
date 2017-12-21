package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	"github.com/tiancaiamao/shen-go/kl"
	"github.com/tiancaiamao/shen-go/vm"
)

var pprof bool
var boot bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&boot, "boot", false, "use bootstrap file for testing")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}
	if boot {
		vm.Boot = true
	}

	vm.Bootstrap()

	m := vm.New()
	m.Eval(kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
}
