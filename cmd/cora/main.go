package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/tiancaiamao/shen-go/cora"
)

var (
	pprof bool
	quiet bool
)

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&quiet, "quiet", false, "donot load init file")
}

var symMacroExpand = cora.MakeSymbol("macroexpand")

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e cora.ControlFlow
	cora.PrimNS1Set(symMacroExpand, cora.Nil)
	cora.PrimNS1Set(cora.MakeSymbol("make-code-generator"), makeCodeGenerator)
	cora.PrimNS1Set(cora.MakeSymbol("cg:bc->go"), bcToGo)

	if !quiet {
		err := cora.Call(&e, cora.PrimNS1Value(cora.MakeSymbol("cora.init")))
		if cora.IsError(err) {
			os.Exit(-1)
		}
	}
	repl(&e)
}

func repl(e *cora.ControlFlow) {
	r := cora.NewSexpReader(os.Stdin, true)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}

		expand := cora.PrimNS1Value(symMacroExpand)
		if expand != cora.Nil {
			sexp = cora.Call(e, expand, sexp)
		}
		// fmt.Println("after macroexpand = ", cora.ObjString(sexp))

		res := cora.Eval(e, sexp)
		fmt.Println(cora.ObjString(res))
	}
}
