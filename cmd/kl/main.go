package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/tiancaiamao/shen-go/klambda"
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

	var e klambda.ControlFlow
	// klambda.PrimNS1Set(symMacroExpand, klambda.Nil)

	// klambda.Init(&e, true)
	// klambda.KLInit(&e, true)
	// klambda.PrimNS2Set(klambda.MakeSymbol("try-catch"), klambda.MakeNative(klambda.PrimTryCatch(true), 2))

	r := klambda.NewSexpReader(os.Stdin, false)
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
		fmt.Println(klambda.ObjString(res))
	}
}

var (
	symMacroExpand = klambda.MakeSymbol("macroexpand")
)
