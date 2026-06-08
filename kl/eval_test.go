package kl

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	type testCase struct {
		name   string
		input  string
		output string
	}
	cases := []testCase{
		{
			name: "let-variable-shadow",
			input: `(do (defun f (a b)
					   (let a 3 a)) (f 4 5))`,
			output: "3",
		},

		{
			name: "let variable shadow",
			input: `(let Result 123
     (let Result 456
	  (if (= Result 456)
	      true
	      Result)))`,
			output: "true",
		},

		{
			name:   "trap-let",
			input:  "(trap-error (let X 666 42) (lambda E (cons --> (cons A ()))))",
			output: "42",
		},

		{
			name:   "curry-partial",
			input:  `((lambda x (lambda y (lambda z (+ x z)))) 1 2 3)`,
			output: "4",
		},

		{
			name:   "curry",
			input:  `(do (defun f (x y z) y) ((f 1 2) 3))`,
			output: "2",
		},

		{
			name: "curry1",
			input: `(do (defun f (x)
		 (do (defun ignore (z w)
		   (lambda y
		      z)) (ignore)))
		(((f 1) 2 3) 4))`,
			output: "2",
		},

		{
			name: "fib10",
			input: `(do (defun fib (i)
	(if (= i 0)
	    1
	    (if (= i 1)
		1
		(+ (fib (- i 1)) (fib (- i 2))))))
	(fib 10))`,
			output: "89",
		},

		{
			name: "proper tail call",
			input: `(do (defun sum (r i)
	  (if (= i 0)
	      r
	      (sum (+ r 1) (- i 1))))
	(sum 0 5000000))`,
			output: "5000000",
		},

		{
			name:   "do in args",
			input:  `(+ (do 1 (do 2 3)) 4)`,
			output: "7",
		},

		// testCase{
		// 	name:   "partial primitive",
		// 	input:  `(+ (+ (+ 1 2) 3) 4)`,
		// 	output: "10",
		// },

		{
			name:   "do in tail call",
			input:  `(do (defun f (x y z) (do 1 (do 2 z))) (f 1 2 3))`,
			output: "3",
		},

		{
			name:   "basic func call",
			input:  `(do (defun id (x) x) (id (do 1 (do 2 42))))`,
			output: "42",
		},

		{
			name:   "identify function",
			input:  `(do (defun id (x) x) (id 42))`,
			output: "42",
		},

		{
			name:   "basic set",
			input:  `(do (set x 42) (value x))`,
			output: "42",
		},

		{
			name:   "basic if",
			input:  `(if true 1 2)`,
			output: "1",
		},

		{
			name:   "basic lambda",
			input:  `((lambda x (lambda y (lambda z z))) 1 2 3)`,
			output: "3",
		},

		{
			name:   "basic do",
			input:  `(do 1 2)`,
			output: "2",
		},

		{
			name:   "basic primitive",
			input:  `(+ 3 7)`,
			output: "10",
		},

		// testCase{
		// 	name:   "partial primitive1",
		// 	input:  `((+ 1) 2)`,
		// 	output: "3",
		// },

		// testCase{
		// 	name:   "partial primitive2",
		// 	input:  `(((+) 1) 2)`,
		// 	output: "3",
		// },
	}

	var ctx ControlFlow
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := evalString(&ctx, c.input)
			if ObjString(res) != c.output {
				fmt.Println("input is:", c.input)
				fmt.Println("output is:", ObjString(res))
				t.Fail()
			}
			if ctx.pos != 0 {
				fmt.Println("unexpected sp after evaluation:", ctx.pos)
				t.Fail()
			}
			if len(ctx.data) != 0 {
				fmt.Println("unexpected stack after evaluation:", len(ctx.data))
				t.Fail()
			}
		})
	}
}

// func TestYYY(t *testing.T) {
// 	var ctx ControlFlow
// 	evalString(&ctx, "(defun return (x) (lambda k (k x)))")
// 	evalString(&ctx, "(defun add1 (n) (return (+ n 1)))")
// 	res := evalString(&ctx, "(add1 4 (lambda x x))")
// 	if res != MakeInteger(5) {
// 		t.Fail()
// 	}
// }

func evalString(ctx *ControlFlow, exp string) Obj {
	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		panic(err)
	}
	return Eval(ctx, sexp)
}

// TestBytecodeVM validates key VM features: closure capture, trap-error,
// let-shadowing, currying, and tail calls through compiled defuns.
func TestBytecodeVM(t *testing.T) {
	type tc struct {
		name  string
		input string
		want  string
	}
	cases := []tc{
		{
			"compiled tak",
			`(do (defun tak (X Y Z)
			       (if (not (< Y X)) Z
			           (tak (tak (- X 1) Y Z)
			                (tak (- Y 1) Z X)
			                (tak (- Z 1) X Y))))
			     (tak 12 6 3))`,
			"4", // tak(12,6,3) = 4 per standard definition
		},
		{
			"closure upval capture",
			`(do (defun make-adder (X) (lambda Y (+ X Y)))
			     ((make-adder 10) 5))`,
			"15",
		},
		{
			"let shadowing in compiled fn",
			`(do (defun f (X)
			       (let X 99
			            X))
			     (f 1))`,
			"99",
		},
		{
			"trap-error in compiled fn",
			`(do (defun safe (X)
			       (trap-error (simple-error "oops") (lambda E (error-to-string E))))
			     (safe 1))`,
			`"oops"`,
		},
		{
			"compiled fib",
			`(do (defun fib (N)
			       (if (= N 0) 0
			           (if (= N 1) 1
			               (+ (fib (- N 1)) (fib (- N 2))))))
			     (fib 10))`,
			"55",
		},
		{
			"compiled tail-recursive sum",
			`(do (defun sum (R I)
			       (if (= I 0) R (sum (+ R 1) (- I 1))))
			     (sum 0 1000000))`,
			"1000000",
		},
		{
			"currying through compiled fn",
			`(do (defun add (X Y) (+ X Y))
			     ((add 3) 4))`,
			"7",
		},
		// P1 regression: defun inside defun must capture outer lexical scope.
		{
			"nested defun captures outer local",
			`(do (defun outer (X)
			       (do (defun inner (Y) (+ X Y))
			           (inner 10)))
			     (outer 5))`,
			"15",
		},
		// P1 regression: (if non-bool ...) must signal an error, not silently
		// treat any non-False value as truthy.  Wrap in a compiled defun so the
		// OP_JUMP_FALSE path is exercised (not just the tree-walker's evalIf).
		{
			"if on non-boolean panics",
			`(do (defun bad-if () (if 42 1 2))
			     (trap-error (bad-if) (lambda E (error-to-string E))))`,
			`"if requires a boolean"`,
		},
		// P2 regression: a symbol that has a global function binding must be
		// looked up as OP_LOAD_GLOBAL even when a same-named local is in scope.
		{
			"global fn takes precedence over local in call position",
			`(do (defun id (X) X)
			     (do (defun call-id (id) (id id))
			         (call-id 42)))`,
			"42",
		},
		// Multi-level closure: closure inside closure must chain upvalue capture.
		{
			"multi-level closure upval chain",
			`(do (defun make-adder2 (X)
			       (lambda Y (lambda Z (+ X (+ Y Z)))))
			     (((make-adder2 1) 2) 3))`,
			"6",
		},
		// Over-application of a compiled (bytecode) function.
		{
			"over-application of bytecode fn",
			`(do (defun add (X Y) (+ X Y))
			     (add 3 4))`,
			"7",
		},
		// VM-path float comparisons through compiled defuns.
		{
			"compiled float comparisons",
			`(do (defun lt (X Y) (< X Y))
			     (do (defun le (X Y) (<= X Y))
			         (do (defun gt (X Y) (> X Y))
			             (do (defun ge (X Y) (>= X Y))
			                 (if (lt 1.5 2.0)
			                     (if (le 2.0 2.0)
			                         (if (gt 3.0 1.5)
			                             (if (ge 2.0 2.0) true false)
			                             false)
			                         false)
			                     false)))))`,
			"true",
		},
	}
	var ctx ControlFlow
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ObjString(evalString(&ctx, c.input))
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestTypeIsMacro(t *testing.T) {
	var ctx ControlFlow
	// type is implemented as a macro rather than a primitive.
	// This is because the second argument is not evaluated.
	// See also https://github.com/tiancaiamao/shen-go/pull/30
	// This shouldn't something like
	// 	"Panic: can't apply non function: undefined"
	res := evalString(&ctx, "(type (cons 1 ()) (undefined func))")
	if ObjString(res) != "(1)" {
		t.Fail()
	}
}
