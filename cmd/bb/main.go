package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/tiancaiamao/shen-go/cora"
)

var (
	pprof    bool
	quiet    bool
	evalStr  string
	evalFile string
)

func init() {
	flag.StringVar(&evalStr, "e", "", "eval a string")
	flag.StringVar(&evalFile, "f", "", "eval a file")

	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&quiet, "quiet", false, "donot load init file")
}

var symQuote = cora.MakeSymbol("quote")
var symMacroExpand = cora.MakeSymbol("macroexpand")

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	// Override the 'load' function
	cora.PrimNS1Set(cora.MakeSymbol("load"), cora.MakeNative(cora.PrimLoadFile(true), 1))

	var e cora.ControlFlow
	cora.PrimNS1Set(symMacroExpand, cora.Nil)
	// cora.PrimNS1Set(cora.MakeSymbol("make-code-generator"), makeCodeGenerator)
	// cora.PrimNS1Set(cora.MakeSymbol("cg:bc->go"), bcToGo)

	if !quiet {
		cora.CoraInit(&e, true)
		if cora.IsError(e.Get(0)) {
			os.Exit(-1)
		}
	}

	if evalStr != "" {
		sexp, err := readStringAsSexp(evalStr)
		if err != nil {
			os.Exit(-1)
		}
		sexp = macroExpand(sexp)
		res := cora.Neval(sexp)
		fmt.Println(cora.ObjString(res))
		return
	}

	if evalFile != "" {
		load := cora.PrimNS1Value(cora.MakeSymbol("load"))
		res := cora.Call(&e, load, cora.MakeString(evalFile))
		fmt.Println(cora.ObjString(res))
		return
	}

	repl(&e)
}

func readStringAsSexp(str string) (cora.Obj, error) {
	r := cora.NewSexpReader(strings.NewReader(evalStr), true)
	return r.Read()
}

func macroExpand(sexp cora.Obj) cora.Obj {
	expand := cora.PrimNS1Value(symMacroExpand)
	if expand != cora.Nil {
		sexp = cora.NCall(expand, sexp)
	}
	return sexp
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

		sexp = macroExpand(sexp)
		// fmt.Println("after macroexpand = ", cora.ObjString(sexp))

		res := cora.Neval(sexp)
		fmt.Println(cora.ObjString(res))
	}
}
