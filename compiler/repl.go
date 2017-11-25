package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
	"github.com/tiancaiamao/shen-go/vm"
)

func main() {
	vm.StdBC, _ = os.Create("bytecode.log")
	vm.StdDebug, _ = os.Create("debug.log")
	vm.StdIn, _ = os.Open("fifo")

	r := kl.NewSexpReader(vm.StdIn)
	m := vm.New()
	for {
		sexp, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(vm.StdErr, err)
			continue
		}

		var a vm.Assember
		err = a.FromSexp(sexp)
		if err != nil {
			fmt.Fprintln(vm.StdErr, err)
			continue
		}
		code := a.Comiple()
		v, err := m.Run(code)
		if err != nil {
			fmt.Fprintln(vm.StdErr, err)
			m.Debug()
			m.Reset()
			continue
		}

		fmt.Fprintln(vm.StdOut, kl.ObjString(v))
		m.Debug()
		fmt.Fprintln(vm.StdBC)
	}

}
