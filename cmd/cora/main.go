package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/tiancaiamao/shen-go/cora"
	"github.com/tiancaiamao/shen-go/lib/babashka"
)

var (
	pprof     bool
	quiet     bool
	evalStr   string
	evalFile  string
	bindInput string
)

func init() {
	flag.StringVar(&bindInput, "i", "", "eval a one-liner expression, bind *input* to lines from stdin")
	flag.StringVar(&evalStr, "e", "", "eval a expression")
	flag.StringVar(&evalFile, "f", "", "eval a file")

	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&quiet, "quiet", false, "donot load init file")
}

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
	cora.PrimNS1Set(cora.MakeSymbol("make-code-generator"), makeCodeGenerator)
	cora.PrimNS1Set(cora.MakeSymbol("cg:bc->go"), bcToGo)
	babashka.Init()

	if !quiet {
		cora.Init(&e, true)
		if cora.IsError(e.Get(0)) {
			os.Exit(-1)
		}
	}

	if bindInput != "" {
		input := cora.Nil
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			input = cora.Cons(cora.MakeString(line), input)
		}
		cora.Reverse(input)
		cora.PrimNS1Set(cora.MakeSymbol("*input*"), input)
		evalStr = bindInput
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
		res := cora.NCall(load, cora.MakeString(evalFile))
		fmt.Println(cora.ObjString(res))
		return
	}

	// Start with no arguments, run REPL.
	if len(os.Args) == 1 {
		r := cora.NewSexpReader(os.Stdin, true)
		repl(r, &e, false)
		return
	}

	// Other case, run as shebang script.
	// cora script.cora arg1 arg2 ...
	shebang(&e)
	return
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

func repl(r *cora.SexpReader, e *cora.ControlFlow, batch bool) {
	for i := 0; ; i++ {
		if !batch {
			fmt.Printf("%d #> ", i)
		}
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

		if !batch {
			fmt.Println(cora.ObjString(res))
		}
	}
}

func shebang(e *cora.ControlFlow) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	// Ignore the shebang line:
	// #!/usr/bin/env cora
	//
	// (followed by cora script ...)
	//
	r1 := bufio.NewReader(f)
	_, err = r1.ReadBytes('\n')
	if err != nil {
		fmt.Println("invalid script")
		return
	}

	args := cora.Nil
	for _, arg := range os.Args[1:] {
		args = cora.Cons(cora.MakeString(arg), args)
	}
	args = cora.Reverse(args)
	cora.PrimNS1Set(cora.MakeSymbol("*command-line-args*"), args)

	r := cora.NewSexpReader(r1, true)
	repl(r, e, true)
}
