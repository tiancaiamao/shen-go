package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool
var shen bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.BoolVar(&shen, "shen", false, "run shen REPL")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	e := kl.NewEvaluator()
	if shen {
		bootstrapShen(e)
		e.Eval(kl.Cons(kl.MakeSymbol("shen.shen"), kl.Nil))
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

func bootstrapShen(e *kl.Evaluator) {
	e.LoadFile("ShenOSKernel-20.1/klambda/toplevel.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/core.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/sys.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/sequent.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/yacc.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/reader.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/prolog.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/track.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/load.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/writer.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/macros.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/declarations.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/t-star.kl")
	e.LoadFile("ShenOSKernel-20.1/klambda/types.kl")
}
