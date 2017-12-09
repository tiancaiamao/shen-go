package vm

import (
	"os"
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestProcedureCall(t *testing.T) {
	var a Assember
	// ((lambda x (lambda y (+ x y))) 1 2)
	a.GRAB(6)
	a.GRAB(4)
	a.ACCESS(1)
	a.ACCESS(0)
	a.PRIMCALL(23)
	a.RETURN()
	a.RETURN()
	a.CONST(kl.MakeInteger(1))
	a.CONST(kl.MakeInteger(2))
	a.TAILAPPLY(2)
	a.HALT()
	code := a.Comiple()

	vm := New()
	o := vm.Run(code)
	if kl.PrimEqual(o, kl.MakeInteger(3)) != kl.True {
		t.Error("failed!")
	}
}

func TestVM(t *testing.T) {
	vm := New()
	runTest(vm, "(* 4 7)", kl.MakeInteger(28), t)
	runTest(vm, "(if true 3 2)", kl.MakeInteger(3), t)
	runTest(vm, "(cond (true 1) (false 2))", kl.MakeInteger(1), t)
	runTest(vm, "(((lambda x (lambda y x)) (lambda z z)) 2 3)", kl.MakeInteger(3), t)
	runTest(vm, `(do (defun fact (n) (if (= n 0) 1 (* n (fact (- n 1))))) (fact 5))`, kl.MakeInteger(120), t)
	runTest(vm, "(let x 3 (let y 5 (+ x y)))", kl.MakeInteger(8), t)
	runTest(vm, "((+ 1) 2)", kl.MakeInteger(3), t)
	runTest(vm, "(let V 4 (let V1 (+ V 1) V1))", kl.MakeInteger(5), t)
	runTest(vm, "((freeze 1))", kl.MakeInteger(1), t)
	runTest(vm, "(do (defun f (a b) (+ a b)) (f 1 2))", kl.MakeInteger(3), t)
	runTest(vm, `(trap-error (simple-error "asd") (lambda X (error-to-string X)))`, kl.MakeString("asd"), t)
	runTest(vm, `(trap-error 42 (lambda X 42))`, kl.MakeInteger(42), t)
}

func runTest(vm *VM, input string, result kl.Obj, t *testing.T) {
	r := kl.NewSexpReader(strings.NewReader(input))
	sexp, err := r.Read()
	if err != nil {
		t.Error(err)
	}

	code, err := klToByteCode(sexp)
	if err != nil {
		t.Error("input to kl bytecode fail", input)
	}

	o := vm.Run(code)
	if kl.PrimEqual(o, result) != kl.True {
		t.Error("failed!", input)
	}
}

func TestTrapError(t *testing.T) {
	vm := New()
	// If exception handle is not quit leave VM a clean state,
	// run it twice, something goes wrong.
	runTest(vm, "(trap-error (value XXX) (lambda E ((freeze 42))))", kl.MakeInteger(42), t)
	runTest(vm, "(trap-error (value XXX) (lambda E ((freeze 42))))", kl.MakeInteger(42), t)

	// Test for a  bug fix that vm.code and code are mess.
	runTest(vm, "(defun thaw (F) (F))", kl.MakeSymbol("thaw"), t)
	runTest(vm, "(defun value/or (V2876 V2877) (trap-error (value V2876) (lambda E (thaw V2877))))", kl.MakeSymbol("value/or"), t)
	runTest(vm, "(value/or XXX (freeze 42))", kl.MakeInteger(42), t)
}

func init() {
	BootstrapCompiler()
	StdDebug, _ = os.Open("/dev/null")
	StdBC, _ = os.Open("/dev/null")
}

func TestKLToBytecode(t *testing.T) {
	tests := [][2]string{
		[2]string{
			"(cons 1 ())",
			"((iConst 1) (iConst ()) (iPrimCall 34) (iHalt))"},
		[2]string{
			"(+ 1 2)",
			"((iConst 1) (iConst 2) (iPrimCall 23) (iHalt))",
		},
		[2]string{
			"(defun f () 1)",
			"((iFreeze (iConst 1) (iReturn)) (iConst f) (iDefun) (iHalt))",
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
