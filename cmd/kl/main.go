package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"plugin"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool
var shen bool
var pluginFile string

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&shen, "shen", false, "run shen REPL")
	flag.StringVar(&pluginFile, "plugin", "", "load plugin")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	e := kl.NewEvaluator()

	if len(pluginFile) > 0 {
		p, err := plugin.Open(pluginFile)
		if err != nil {
			fmt.Println("load plugin fail:", err)
			return
		}
		f, err := p.Lookup("Regist")
		if err != nil {
			fmt.Println("Regist function is not provided!")
			return
		}
		regist, ok := f.(func(*kl.Evaluator))
		if !ok {
			fmt.Println("Regist function format wrong!")
			return
		}
		regist(e)
	}

	if shen {
		e.BootstrapShen()
		e.Eval(kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
		e.Eval(kl.Cons(kl.MakeSymbol("shen.repl"), kl.Nil))
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
