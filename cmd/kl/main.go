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
var script string

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.StringVar(&script, "s", "", "bootstrap script")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	e := kl.NewEvaluator()
	if script != "" {
		e.LoadFile(script)
		e.Eval(kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
		return
	}

	r := kl.NewSexpReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}

		res := e.Eval(sexp)
		fmt.Println(kl.ObjString(res))
	}
}
