package vm

import (
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestProcedureCall(t *testing.T) {
	var a Assember
	// (lambda x (lambda y (+ x y)) 1 2)
	a.GRAB(6)
	a.GRAB(4)
	a.ACCESS(0)
	a.ACCESS(1)
	a.PRIMCALL(23)
	a.RETURN()
	a.RETURN()
	a.CONST(kl.MakeInteger(1))
	a.CONST(kl.MakeInteger(2))
	a.TAILAPPLY(2)
	a.HALT()
	code := a.Comiple()

	vm := New()
	o, err := vm.Run(code)
	if err != nil || kl.PrimEqual(o, kl.MakeInteger(3)) != kl.True {
		t.Error("failed!")
	}
}

func TestPartialApply(t *testing.T) {
	// (((lambda x (lambda y x)) (lambda z z)) 2 3)
	// str := "((iGrab (iGrab (iAccess 1) (iReturn)) (iReturn)) (iGrab (iAccess 0) (iReturn)) (iPushArg) (iApply) (iConst 2) (iPushArg) (iConst 3) (iPushArg) (iApply) (iHalt))"

	// [do [defun f [a b] [+ a b]] [f 1 2]]
	// str := "((iGrab (iGrab (iAccess 1) (iAccess 0) (iPrimCall 2) (iReturn)) (iReturn)) (iConst f) (iDefun) (iPop) (iConst f) (iGetF) (iConst 1) (iPushArg) (iConst 2) (iPushArg) (iApply) (iHalt))"

	// (if true 3 2)
	// str := "((iConst true) (iJF (iConst 3)) (iJMP (iConst 2)) (iHalt))"

	// (* 4 7)
	// str := "((iConst 4) (iConst 7) (iPrimCall 21) (iHalt))"

	// (do (defun fact (n)
	// 	(if (= n 0)
	// 		1
	// 		(* n (fact (- n 1)))))
	// 	(fact 5))
	// str := "((iGrab (iAccess 0) (iConst 0) (iPrimCall 19) (iJF (iConst 1)) (iJMP (iAccess 0) (iConst fact) (iGetF) (iAccess 0) (iConst 1) (iPrimCall 20) (iPushArg) (iApply) (iPrimCall 21)) (iReturn)) (iConst fact) (iDefun) (iPop) (iConst fact) (iGetF) (iConst 5) (iPushArg) (iTailApply) (iHalt))"
	// str := "((iGrab (iGrab (iAccess 1) (iConst 0) (iPrimCall 19) (iJF (iAccess 0)) (iJMP (iConst fact) (iGetF) (iAccess 1) (iConst 1) (iPrimCall 20) (iPushArg) (iAccess 1) (iAccess 0) (iPrimCall 21) (iPushArg) (iApply)) (iReturn)) (iReturn)) (iConst fact) (iDefun) (iPop) (iConst fact) (iGetF) (iConst 5) (iPushArg) (iConst 1) (iPushArg) (iTailApply) (iHalt))"
	// str := "((iGrab (iAccess 0) (iConst 0) (iPrimCall 19) (iJF (iConst 1)) (iJMP (iAccess 0) (iConst fact) (iGetF) (iAccess 0) (iConst 1) (iPrimCall 20) (iPushArg) (iApply) (iPrimCall 21)) (iReturn)) (iConst fact) (iDefun) (iPop) (iConst fact) (iGetF) (iConst 5) (iPushArg) (iApply) (iHalt))"

	// [let x 3 [let y 5 [+ x y]]]
	str := "((iGrab (iGrab (iAccess 1) (iAccess 0) (iPrimCall 23) (iReturn)) (iConst 5) (iApply 1) (iReturn)) (iConst 3) (iApply 1) (iHalt))"

	// ((freeze 1))
	// str := "((iFreeze (iConst 1) (iReturn)) (iApply) (iHalt))"

	// [cond [true 1] [false 2]]
	// str := `((iConst true) (iJF (iConst 1)) (iJMP (iConst false) (iJF (iConst 2)) (iJMP (iConst "no match cond") (iPrimCall 18))) (iHalt))`

	// ((+ 1) 2)
	// str := "((iGrab (iConst 1) (iAccess 0) (iPrimCall 23) (iReturn)) (iConst 2) (iPushArg) (iApply) (iHalt))"

	// [trap-error [simple-error "asd"] [lambda X 42]]
	// str := `((iSetJmp (iGrab (iConst 42) (iReturn) (iReturn))) (iConst "asd") (iPrimCall 18) (iReturn) (iHalt))`

	// str := `((iConst "asd") (iPrimCall 18) (iReturn) (iHalt))`

	r := kl.NewSexpReader(strings.NewReader(str))
	sexp, err := r.Read()
	if err != nil {
		t.Error(err)
	}

	var a Assember
	a.FromSexp(sexp)
	code := a.Comiple()

	vm := New()
	o, err := vm.Run(code)
	if err != nil || kl.PrimEqual(o, kl.MakeInteger(8)) != kl.True {
		t.Error("failed!")
	}
}

func TestKLToBytecode(t *testing.T) {
	BootstrapCompiler()

	tests := [][2]string{
		[2]string{
			"(cons 1 ())",
			"((iConst 1) (iConst ()) (iPrimCall 34) (iReturn) (iHalt))"},
		[2]string{
			"(+ 1 2)",
			"((iConst 1) (iConst 2) (iPrimCall 23) (iReturn) (iHalt))",
		},
		[2]string{
			"(defun f () 1)",
			"((iFreeze (iConst 1) (iReturn)) (iConst f) (iDefun) (iReturn) (iHalt))",
		},
	}
	for _, test := range tests {
		testKLToBytecode(t, test[0], test[1])
	}

}

func testKLToBytecode(t *testing.T, input, expect string) {
	r1 := kl.NewSexpReader(strings.NewReader(input))
	r2 := kl.NewSexpReader(strings.NewReader(expect))

	klambda, err := r1.Read()
	if err != nil {
		t.Fatal(err)
	}
	bc := klToSexpByteCode(klambda)

	bc1, err := r2.Read()
	if err != nil {
		t.Fatal(err)
	}

	if kl.PrimEqual(bc, bc1) != kl.True {
		t.Errorf("input:%s\n expect:%s\n get:%s\n", input, expect, kl.ObjString(bc))
	}
}
