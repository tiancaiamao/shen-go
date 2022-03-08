package klambda

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
		testCase{
			name: "let-variable-shadow",
			input: `(do (defun f (a b)
					   (let a 3 a)) (f 4 5))`,
			output: "3",
		},

		testCase{
			name: "let variable shadow",
			input: `(let Result 123
     (let Result 456
	  (if (= Result 456)
	      true
	      Result)))`,
			output: "true",
		},

		testCase{
			name:   "trap-let",
			input:  "(trap-error (let X 666 42) (lambda E (cons --> (cons A ()))))",
			output: "42",
		},

		testCase{
			name:   "curry-partial",
			input:  `((lambda x (lambda y (lambda z (+ x z)))) 1 2 3)`,
			output: "4",
		},

		testCase{
			name:   "curry",
			input:  `(do (defun f (x y z) y) ((f 1 2) 3))`,
			output: "2",
		},

		testCase{
			name: "curry1",
			input: `(do (defun f (x)
		 (do (defun ignore (z w)
		   (lambda y
		      z)) (ignore)))
		(((f 1) 2 3) 4))`,
			output: "2",
		},

		testCase{
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

		testCase{
			name: "proper tail call",
			input: `(do (defun sum (r i)
	  (if (= i 0)
	      r
	      (sum (+ r 1) (- i 1))))
	(sum 0 5000000))`,
			output: "5000000",
		},

		testCase{
			name:   "do in args",
			input:  `(+ (do 1 (do 2 3)) 4)`,
			output: "7",
		},

		testCase{
			name:   "partial primitive",
			input:  `(+ (+ (+ 1 2) 3) 4)`,
			output: "10",
		},

		testCase{
			name:   "do in tail call",
			input:  `(do (defun f (x y z) (do 1 (do 2 z))) (f 1 2 3))`,
			output: "3",
		},

		testCase{
			name:   "basic func call",
			input:  `(do (defun id (x) x) (id (do 1 (do 2 42))))`,
			output: "42",
		},

		testCase{
			name:   "identify function",
			input:  `(do (defun id (x) x) (id 42))`,
			output: "42",
		},

		testCase{
			name:   "basic set",
			input:  `(do (set x 42) (value x))`,
			output: "42",
		},

		testCase{
			name:   "basic if",
			input:  `(if true 1 2)`,
			output: "1",
		},

		testCase{
			name:   "basic lambda",
			input:  `((lambda x (lambda y (lambda z z))) 1 2 3)`,
			output: "3",
		},

		testCase{
			name:   "basic do",
			input:  `(do 1 2)`,
			output: "2",
		},

		testCase{
			name:   "basic primitive",
			input:  `(+ 3 7)`,
			output: "10",
		},

		testCase{
			name:   "partial primitive1",
			input:  `((+ 1) 2)`,
			output: "3",
		},

		testCase{
			name:   "partial primitive2",
			input:  `(((+) 1) 2)`,
			output: "3",
		},
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
			if len(ctx.stack) != 0 {
				fmt.Println("unexpected stack after evaluation:", len(ctx.stack))
				t.Fail()
			}
		})
	}
}

func TestYYY(t *testing.T) {
	var ctx ControlFlow
	evalString(&ctx, "(defun return (x) (lambda k (k x)))")
	evalString(&ctx, "(defun add1 (n) (return (+ n 1)))")
	res := evalString(&ctx, "(add1 4 (lambda x x))")
	if res != MakeInteger(5) {
		t.Fail()
	}
}

func evalString(ctx *ControlFlow, exp string) Obj {
	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		panic(err)
	}
	return ctx.Eval(sexp)
}
