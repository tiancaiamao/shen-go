package main

import (
	"fmt"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
	"github.com/tiancaiamao/shen-go/vm"
)

func main() {
	vm.StdBC, _ = os.Create("bytecode.log")
	vm.StdDebug, _ = os.Create("debug.log")

	kl.LoadFile("shen.kl")
	kl.LoadFile("cmd/shen/primitive.kl")
	kl.LoadFile("cmd/shen/de-bruijn.kl")
	kl.LoadFile("cmd/shen/compile.kl")

	r := kl.NewSexpReader(vm.StdIn)
	m := vm.New()
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		res := m.Eval(sexp)
		fmt.Println(kl.ObjString(res))
	}
}
