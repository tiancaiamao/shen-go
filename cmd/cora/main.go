package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	"github.com/tiancaiamao/shen-go/kl"
)

var (
	pprof bool
	quiet bool
)

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&quiet, "quiet", false, "donot load init file")
}

var symMacroExpand = kl.MakeSymbol("macroexpand")

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e kl.Cora
	kl.CoraSet(symMacroExpand, kl.Nil)

	if !quiet {
		home, err := os.UserHomeDir()
		if err == nil {
			path := filepath.Join(home, ".cora", "init.cora")
			if _, err := os.Stat(path); err == nil {
				res := kl.Call(&e, e.Global(kl.MakeSymbol("load")), kl.MakeString(path))
				if kl.IsError(res) {
					os.Exit(-1)
				}
			}
		}
	}
	repl(&e)
}

func repl(e *kl.Cora) {
	r := kl.NewSexpReader(os.Stdin, true)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}

		expand := e.Global(symMacroExpand)
		if expand != kl.Nil {
			sexp = kl.Call(e, expand, sexp)
		}
		// fmt.Println("after macroexpand = ", kl.ObjString(sexp))

		res := kl.Eval(e, sexp)
		fmt.Println(kl.ObjString(res))
	}
}
