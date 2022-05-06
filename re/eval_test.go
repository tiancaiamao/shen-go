package re

import (
	"testing"
	"strings"
	"fmt"
	"io"
)

func TestCompile(t *testing.T) {
	vm := New()	
	// input := "42"
	// input := `"hello world"`
	// input := "(if false 23 42)"
	// input := "(do 1 42)"
	// input := "(do (set (quote x) 3) (value (quote x)))"
	// input := "(+ 4 7)"
	// input := "((lambda (x) x) 42)"
	// input := "((lambda (x) (+ x 3)) 1)"
	//  input := "(do (set (quote f) (lambda (x y) (+ x y))) (do 42 (f 3 5) ))"
	// input :=  `(do (set (quote id) (lambda  (x) x)) (id 42))`
	// input := `(+ (do 1 (do 2 3)) 4)`
	input := `(do (set (quote f) (lambda (x y z) (do 1 (do 2 z)))) (f 1 2 3))`

/*	input := `(do (set (quote fib) (lambda (i)
	(if (= i 0)
	    1
	    (if (= i 1)
		1
		(+ (fib (- i 1)) (fib (- i 2)))))))
	(fib 10))`

	input := `(do (set (quote sum) (lambda (r i)
	  (if (= i 0)
	      r
	      (sum (+ r 1) (- i 1)))))
	(sum 0 5000000))`
*/	
	
	r := NewSexpReader(strings.NewReader(input), false)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		t.Error("read error", err)
	}
	
	res := eval(vm, sexp)
	fmt.Println(res)
}