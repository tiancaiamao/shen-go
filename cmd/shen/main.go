package main

import (
	// "fmt"
	// "os"
	// "io"
	"flag"
	"net/http"
	_ "net/http/pprof"

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

	e := kl.NewKLambda()
	Regist(e)

	kl.Eval(e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	kl.Eval(e, kl.Cons(kl.MakeSymbol("shen.repl"), kl.Nil))

	// r := kl.NewSexpReader(os.Stdin, false)
	// for i := 0; ; i++ {
	// 	fmt.Printf("%d #> ", i)
	// 	sexp, err := r.Read()
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println("read error:", err)
	// 		}
	// 		break
	// 	}

	// 	res := e.Eval(sexp)
	// 	fmt.Println(kl.ObjString(res))
	// }

	return
}
