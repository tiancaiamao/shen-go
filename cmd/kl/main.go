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

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var ctl kl.ControlFlow
	kl.BindSymbolFunc(kl.MakeSymbol("bc->go"), bcToGo)
	kl.BindSymbolFunc(kl.MakeSymbol("make-code-generator"), makeCodeGenerator)

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
		res := kl.Eval(&ctl, sexp)
		fmt.Println(kl.ObjString(res))
	}
}
