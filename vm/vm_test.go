package vm

import (
	// "strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestProcedureCall(t *testing.T) {
	var a Assember
	// ((lambda x lambda y (+ x y)) 1 2)
	a.GRAB(6)
	a.GRAB(4)
	a.ACCESS(0)
	a.ACCESS(1)
	a.ADD()
	a.RETURN()
	a.RETURN()
	a.CONST(kl.Make_integer(1))
	a.PUSHARG()
	a.CONST(kl.Make_integer(2))
	a.PUSHARG()
	a.APPLY()
	a.HALT()
	code := a.Comiple()

	vm := NewVM()
	vm.Run(code)
	if kl.PrimEqual(vm.stack[vm.top-1], kl.Make_integer(3)) != kl.True {
		t.Error("failed!")
	}
}

// func TestSimple(t *testing.T) {
// 	str := "((iGrab (iGrab (iAccess 1) (iAccess 0) (iPrimCall 2) (iReturn)) (iReturn)) (iConst 1) (iPushArg) (iConst 2) (iPushArg) (iApply) (iHalt))"
// 	r := kl.NewSexpReader(strings.NewReader(str))
// 	sexp, err := r.Read()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	var a Assember
// 	a.FromSexp(sexp)
// 	code := a.Comiple()

// 	vm := NewVM()
// 	vm.Run(code)
// 	vm.debug()
// }
