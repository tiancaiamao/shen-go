package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"

	"github.com/tiancaiamao/shen-go/runtime"
)

var pprof bool
var boot string

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.StringVar(&boot, "boot", "", "use bootstrap file for testing")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}
	runtime.Boot = boot
	runtime.BootstrapCora()
	runtime.RegistNativeCall("primitive.load-file", 1, loadFile)

	m := runtime.NewVM()
	r := runtime.NewSexpReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		// var a vm.Assember
		// a.FromSexp(sexp)
		// code := a.Comiple()
		// res := m.Eval(code)
		// fmt.Println(kl.ObjString(res))

		res := m.Eval(sexp)
		fmt.Println(runtime.ObjString(res))
	}

	// m.Eval(runtime.Cons(runtime.MakeSymbol("shen.shen"), runtime.Nil))
	return
}

func loadFile(args ...runtime.Obj) runtime.Obj {
	file := runtime.GetString(args[0])
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(runtime.PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return runtime.MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return runtime.MakeError(err.Error())
	}
	defer f.Close()

	r := runtime.NewSexpReader(f)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return runtime.MakeError(err.Error())
			}
			break
		}

		res := runtime.Eval(exp)

		if runtime.IsError(res) {
			return res
		}
	}
	return args[0]
}
