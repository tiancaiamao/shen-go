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

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e cora.ControlFlow
	cora.PrimNS1Set(symMacroExpand, cora.Nil)
	if err := cora.NCall(cora.PrimNS1Value(cora.MakeSymbol("cora.init"))); cora.IsError(err) {
		os.Exit(-1)
	}
	if err := cora.NCall(cora.PrimNS1Value(cora.MakeSymbol("kl.init"))); cora.IsError(err) {
		os.Exit(-1)
	}
	r := cora.NewSexpReader(os.Stdin, false)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		res := evalKL(&e, sexp)
		fmt.Println(cora.ObjString(res))
	}
}

var (
	symMacroExpand = cora.MakeSymbol("macroexpand")
)

var symKLToCora = cora.MakeSymbol("kl->cora")

func evalKL(e *cora.ControlFlow, exp cora.Obj) cora.Obj {
	exp1 := cora.NCall(cora.PrimNS1Value(symKLToCora), cora.Nil, exp)

	// fmt.Println("evalKL with ===", kl.ObjString(exp1))
	res := cora.Neval(exp1)
	return res
}
