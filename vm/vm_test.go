package vm

import (
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestProcedureCall(t *testing.T) {
	var a1, a2 Assember
	// ((lambda x lambda y (+ x y)) 1 2)
	a2.GRAB()
	a2.GRAB()
	a2.ACCESS(0)
	a2.ACCESS(1)
	a2.ADD()
	a2.RETURN()
	code := a2.Comiple()
	tmp := Procedure{
		ScmRaw: kl.Make_raw(),
		code:   code,
	}
	proc := tmp.ScmRaw.Object()

	vm := NewVM()
	a1.MARK(3)
	vm.stack[vm.mark+2] = proc
	a1.CONST(kl.Make_integer(1))
	a1.PUSHARG(1)
	a1.CONST(kl.Make_integer(2))
	a1.PUSHARG(0)
	a1.APPLY()
	a1.HALT()
	code = a1.Comiple()

	vm.Run(code)
	if kl.PrimEqual(vm.stack[vm.top-1], kl.Make_integer(3)) != kl.True {
		t.Error("failed!")
	}
}
