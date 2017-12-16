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
	vm.BootstrapCompiler()

	r := kl.NewSexpReader(vm.StdIn)
	m := vm.New()

	m.RegistNativeCall(kl.MakePrimitive("load-bytecode", 1, m.LoadBytecode))
	m.RegistNativeCall(kl.MakePrimitive("primitive?", 1, kl.NativeIsPrimitive))
	m.RegistNativeCall(kl.MakePrimitive("primitive-arity", 1, kl.NativePrimitiveArity))
	m.RegistNativeCall(kl.MakePrimitive("primitive-id", 1, kl.NativePrimitiveID))

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
