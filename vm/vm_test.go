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
	a.CONST(kl.MakeInteger(2))
	a.CONST(kl.MakeInteger(1))
	a.FREEZE(6)
	a.GRAB()
	a.GRAB()
	a.ACCESS(1)
	a.ACCESS(0)
	a.PRIMCALL(23)
	a.RETURN()
	a.TAILAPPLY()
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
	runTest(vm, "(do (defun f (a b) (+ a b)) (f 1 2))", kl.MakeInteger(3), t)
	runTest(vm, `(do (defun fact (n) (if (= n 0) 1 (* n (fact (- n 1))))) (fact 5))`, kl.MakeInteger(120), t)
	runTest(vm, "(let x 3 (let y 5 (- y x)))", kl.MakeInteger(2), t)
	runTest(vm, "((- 6) 2)", kl.MakeInteger(4), t)
	runTest(vm, "((freeze 1))", kl.MakeInteger(1), t)
	runTest(vm, "(symbol? (lambda x x))", kl.False, t)
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
	// runTest(vm, "(trap-error ((lambda X (simple-error \"aa\")) 42) (lambda E false))", kl.False, t)
}

func TestNativeCall(t *testing.T) {
	vm := New()
	vm.RegistNativeCall("hello", 0, helloWorld)

	runTest(vm, "(native hello)", kl.MakeString("hello world"), t)
	runTest(vm, "(native bbc)", kl.MakeError("unknown native function:bbc"), t)
	runTest(vm, "(native hello 1)", kl.MakeError("wrong arity for native hello"), t)
}

func helloWorld(...kl.Obj) kl.Obj {
	return kl.MakeString("hello world")
}

func TestOrder(t *testing.T) {
	vm := New()
	runTest(vm, "(defun f (x y) y)", kl.MakeSymbol("f"), t)
	runTest(vm, "((lambda D ((lambda Fill D) (f 1 2))) 42)", kl.MakeInteger(42), t)
	runTest(vm, "(let D 42 (let Fill (f 1 2) D))", kl.MakeInteger(42), t)
}

func TestReverse(t *testing.T) {
	vm := New()
	res := kl.Cons(kl.MakeInteger(3), kl.Cons(kl.MakeInteger(2), kl.Cons(kl.MakeInteger(1), kl.Nil)))
	runTest(vm, "(defun reverse-h (L R) (if (= L ()) R (reverse-h (tl L) (cons (hd L) R))))", kl.MakeSymbol("reverse-h"), t)
	runTest(vm, "(defun reverse (X) (reverse-h X ()))", kl.MakeSymbol("reverse"), t)
	runTest(vm, "(reverse (cons 1 (cons 2 (cons 3 ()))))", res, t)
}

func TestB521(t *testing.T) {
	vm := New()
	runTest(vm, "(+ (let X 7 (let Y 2 (if (if (= X 7) (< Y 0) (<= 0 Y)) 77 88))) 99)", kl.MakeInteger(187), t)
	runTest(vm, "(if (= (+ 7 (* 2 4)) (- 20 (+ (+ 1 1) (+ (+ 1 1) 1)))) (+ 1 (+ 1 (+ 1 (+ 1 (+ 1 10))))) 0)", kl.MakeInteger(15), t)
	runTest(vm, `(cons (let F (lambda H (lambda V (* H V)))
           (let K (lambda X (+ X 5))
              (let X 15
                (let G (lambda X (+ 1 X))
                  (K (G (let G 3 (F G X))))))))
          ())`, kl.Cons(kl.MakeInteger(51), kl.Nil), t)
	runTest(vm, `(let n 5
      (let a 1
        (let a (* a n)
          (let n (- n 1)
            (let a (* a n)
              (let n (- n 1)
                (let a (* a n)
                  (let n (- n 1)
                    (let a (* a n)
                      a)))))))))`, kl.MakeInteger(120), t)
	runTest(vm, `(let F (lambda P
               (- (<-address
                    (<-address (<-address (<-address (<-address P 0) 0) 1) 0)
                    (<-address (<-address P 1) (<-address (<-address P 0) 4)))
                  (<-address
                    (<-address P (<-address P 2))
                    (<-address (<-address P 0) (<-address P 4)))))
  (let X (absvector 6)
  (let Y (absvector 7)
      (do (address-> X 0 Y)
      (do (address-> X 1 X)
      (do (address-> Y 0 X)
      (do (address-> Y 1 -4421)
      (do (address-> X 2 0)
      (do (address-> X 3 -37131)
      (do (address-> X 4 4)
      (do (address-> X 5 6)
      (do (address-> Y 2 -55151)
      (do (address-> Y 3 -32000911)
      (do (address-> Y 4 5)
      (do (address-> Y 5 55)
      (do (address-> Y 6 -36)
          (* (F X) 2)))))))))))))))))`, kl.MakeInteger(-182), t)
	runTest(vm, "((lambda Y ((lambda F (F (F Y))) (lambda Y Y))) 4)", kl.MakeInteger(4), t)
	runTest(vm, "(((((lambda X (lambda Y (lambda Z (lambda W (lambda U (+ X (+ Y (+ Z (+ W U))))))))) 5) 6 7) 8) 9)", kl.MakeInteger(35), t)
	runTest(vm, `
(let F (+ 1)
(let G (- 1)
(let T (lambda X (- X 1))
(let J (lambda X (- X 1))
(let I (- 1)
(let H (- 1)
(let X 80
(let A (F X)
(let B (G X)
(let C (H (I (J (T X))))
(* A (* B (+ C 0)))))))))))))
`, kl.MakeInteger(-499122), t)
	runTest(vm, `(let a 5
		(let b 4
			(let c (*)
				(let f (cons)
					(if (or (> (c a b) 15)
						(= (c a b) 20))
						(f a b)
						42)))))`, kl.Cons(kl.MakeInteger(5), kl.MakeInteger(4)), t)
}

func TestCPSFib(t *testing.T) {
	vm := New()
	runTest(vm, `
(defun fib (n k)
  (if (or (= n 0) (= n 1))
      (k 1)
    (fib (- n 1)
         (lambda w
           (fib (- n 2) (lambda v
                          (k (+ w v))))))))
`, kl.MakeSymbol("fib"), t)
	runTest(vm, "(fib 10 (lambda X X))", kl.MakeInteger(89), t)
}

func TestCall(t *testing.T) {
	vm := New()
	runTest(vm, "(defun a (u v w x) (if (= u 0) (b v w x) (a (- u 1) v w x)))", kl.MakeSymbol("a"), t)
	runTest(vm, "(defun b (q r x) (let p (* q r) (e (* q r) p x)))", kl.MakeSymbol("b"), t)
	runTest(vm, "(defun c (x) (* 5 x))", kl.MakeSymbol("c"), t)
	runTest(vm, "(defun e (n p x) (if (= n 0) (c p) (o (- n 1) p x)))", kl.MakeSymbol("e"), t)
	runTest(vm, "(defun o (n p x) (if (= 0 n) (c x) (e (- n 1) p x)))", kl.MakeSymbol("o"), t)
	runTest(vm, "(let X 5 (a 3 2 1 X))", kl.MakeInteger(10), t)
}

func TestVectorLength(t *testing.T) {
	vm := New()
	runTest(vm, "(defun visit (V N) (trap-error (do (<-address V N) true) (lambda X false)))", kl.MakeSymbol("visit"), t)
	runTest(vm, `
(defun vector-length-h (V N)
  (if (visit V N)
      (vector-length-h V (+ N 1))
    N))`, kl.MakeSymbol("vector-length-h"), t)
	runTest(vm, "(defun vector-length (V) (vector-length-h V 0))", kl.MakeSymbol("vector-length"), t)
	runTest(vm, "(vector-length (absvector 5))", kl.MakeInteger(5), t)
}

func TestCountLeaves(t *testing.T) {
	vm := New()
	runTest(vm, `(defun count-leaves (p)
  (if (cons? p)
      (+ (count-leaves (hd p))
         (count-leaves (tl p)))
    1))
`, kl.MakeSymbol("count-leaves"), t)
	runTest(vm, `(count-leaves
 (cons
  (cons 0 (cons 0 0))
  (cons
   (cons (cons (cons 0 (cons 0 0)) 0) 0)
   (cons
    (cons (cons 0 0) (cons 0 (cons 0 0)))
    (cons (cons 0 0) 0)))))
`, kl.MakeInteger(16), t)
}

func TestFreeze(t *testing.T) {
	vm := New()
	runTest(vm, "(defun thaw (X) (X))", kl.MakeSymbol("thaw"), t)
	runTest(vm, `(defun add-ths (T1 T2 T3 T4)
		(+ (+ (thaw T1) (thaw T2))
			(+ (thaw T3) (thaw T4))))`, kl.MakeSymbol("add-ths"), t)
	runTest(vm, "(add-ths (freeze 5) (freeze 17) (freeze 7) (freeze 9))", kl.MakeInteger(38), t)
}

func TestFrancisFrenandez(t *testing.T) {
	vm := New()
	runTest(vm, "(defun to-bool (X) (not (= X false)))", kl.MakeSymbol("to-bool"), t)
	runTest(vm, `(and (to-bool
      (+ ((if (not (to-bool (cons 1 2)))
              true
            (let F1 3
                 (let F2 (lambda X (+ X 4))
                      F2))) 5)
        6)) false)`, kl.False, t)
}

func init() {
	StdDebug, _ = os.Open("/dev/null")
	StdBC, _ = os.Open("/dev/null")
}

func TestKLToBytecode(t *testing.T) {
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
