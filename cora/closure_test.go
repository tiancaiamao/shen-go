package cora

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestXXX(t *testing.T) {
	type testCase struct {
		name   string
		input  string
		output string
	}
	cases := []testCase{
		testCase{
			name:   "curry-partial",
			input:  `((lambda (x) (lambda (y z) (+ x z))) 1 2 3)`,
			output: "4",
		},

		testCase{
			name:   "curry",
			input:  `((((lambda (x y z) y)) 1 2) 3)`,
			output: "2",
		},

		testCase{
			name: "curry1",
			input: `(do (set (quote f)
		(lambda (x)
		 (lambda (z w)
		   (lambda (y)
		      z))))
		(((f 1) 2 3) 4))`,
			output: "2",
		},

		testCase{
			name: "fib10",
			input: `(do (set (quote fib)
	(lambda (i)
	(if (= i 0)
	    1
	    (if (= i 1)
		1
		(+ (fib (- i 1)) (fib (- i 2)))))))
	(fib 10))`,
			output: "89",
		},

		testCase{
			name: "proper tail call",
			input: `(do (set (quote sum)
	(lambda (r i)
	  (if (= i 0)
	      r
	      (sum (+ r 1) (- i 1)))))
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
			input:  `((lambda (x y z) (do 1 (do 2 z))) 1 2 3)`,
			output: "3",
		},

		testCase{
			name:   "basic func call",
			input:  `(do (set (quote id) (lambda (x) x)) (id (do 1 (do 2 42))))`,
			output: "42",
		},

		testCase{
			name:   "identify function",
			input:  `(do (set (quote id) (lambda (x) x)) (id 42))`,
			output: "42",
		},

		testCase{
			name:   "basic set",
			input:  `(do (set (quote x) 42) x)`,
			output: "42",
		},

		testCase{
			name:   "basic if",
			input:  `(if true 1 2)`,
			output: "1",
		},

		testCase{
			name:   "basic lambda",
			input:  `((lambda (x y z) z) 1 2 3)`,
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res :=evalString(c.input)
			if ObjString(res) != c.output {
				fmt.Println("input is:", c.input)
				fmt.Println("output is:", ObjString(res))
				t.Fail()
			}
		})
	}
}


func TestYYY(t *testing.T) {
	evalString("(set 'return (lambda (x) (lambda (k) (k x))))")
	evalString("(set 'add1 (lambda (n) (return (+ n 1))))")
	res := evalString("(add1 4 (lambda (x) x))")
	if res != MakeInteger(5) {
		t.Fail()
	}
}

func evalString(exp string) Obj {
	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		panic(err)
	}
	return Neval(sexp)
}
