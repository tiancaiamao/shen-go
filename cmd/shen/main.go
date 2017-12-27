package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

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
		vm.Debug = true
	}

	vm.BootstrapShen()

	m := vm.New()
	m.Eval(kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
	return

	r := kl.NewSexpReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		// var a vm.Assember
		// a.FromSexp(sexp)
		// code := a.Comiple()
		// res := m.Eval(code)
		// fmt.Println(kl.ObjString(res))

		res := m.Eval(sexp)
		fmt.Println(kl.ObjString(res))
	}

}
