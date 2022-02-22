package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/tiancaiamao/shen-go/klambda"
	"github.com/tiancaiamao/cora/cora_go"
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

	cora.Init(&e, true)
	klambda.KLInit(&e, true)
	cora.PrimNS2Set(cora.MakeSymbol("try-catch"), cora.MakeNative(cora.PrimTryCatch(true), 2))

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
		res := klambda.EvalKL(sexp)
		fmt.Println(cora.ObjString(res))
	}
}

var (
	symMacroExpand = cora.MakeSymbol("macroexpand")
)
