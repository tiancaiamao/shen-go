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

	bootstrapCompiler()

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

func bootstrapCompiler() {
	kl.LoadFile("ShenOSKernel-20.1/klambda/toplevel.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/core.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/sys.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/sequent.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/yacc.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/reader.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/prolog.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/track.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/load.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/writer.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/macros.kl")
	kl.LoadFile("ShenOSKernel-20.1/klambda/declarations.kl")
	kl.LoadFile("cmd/shen/primitive.kl")
	kl.LoadFile("cmd/shen/de-bruijn.kl")
	kl.LoadFile("cmd/shen/compile.kl")
}
