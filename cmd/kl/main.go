package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool
var shen bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&shen, "shen", false, "run shen REPL")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e kl.KLambda
	if shen {
		e.BootstrapShen()
		kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
		kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.repl"), kl.Nil))
		return
	}

	r := kl.NewSexpReader(os.Stdin, false)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}

		res := kl.Eval(&e, sexp)
		fmt.Println(kl.ObjString(res))
	}
}
