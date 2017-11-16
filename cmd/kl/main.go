package main

import (
	"flag"
	"fmt"
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
	if script != "" {
		kl.LoadFile(script)
	}

	r := kl.NewSexpReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		res := kl.Eval(sexp)
		fmt.Println(kl.ObjString(res))
	}
}
