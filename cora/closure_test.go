package cora

import (
	"fmt"
	"io"
	"strings"
	"testing"
	// "unsafe"
)

// func primitive(fn func(args ...Obj) Obj, nargs int) Obj {
// 	tmp := scmPrimitive{
// 		scmHead: scmHeadPrimitive,
// 		fn:      fn,
// 		params:  nargs,
// 	}
// 	return Obj(&tmp.scmHead)
// }

// func myadd(args ...Obj) Obj {
// 	x := mustInteger(args[0])
// 	y := mustInteger(args[1])
// 	// fmt.Println("in add ...", x, y)
// 	return MakeInteger(x + y)
// }

// func mymul(args ...Obj) Obj {
// 	x := mustInteger(args[0])
// 	y := mustInteger(args[1])
// 	return MakeInteger(x * y)
// }

// func mysub(args ...Obj) Obj {
// 	x := mustInteger(args[0])
// 	y := mustInteger(args[1])
// 	// fmt.Println("in sub - ", x, y)
// 	return MakeInteger(x - y)
// }

// func myeq(args ...Obj) Obj {
// 	// fmt.Println("EQ== ", ObjString(args[0]), ObjString(args[1]))
// 	return equal(args[0], args[1])
// }

// func myset(args ...Obj) Obj {
// 	return PrimNS1Set(args[0], args[1])
// }

func TestXXX(t *testing.T) {
	// exp := `((lambda (x) (lambda (y z) (+ x z)))
	// 1 2 3)`

	// 	exp := `((((lambda (x y z)
	// y)) 1 2) 3)`

	// 	exp := `(do (set (quote f)
	// (lambda (x)
	//  (lambda (z w)
	//    (lambda (y)
	//       z))))
	// (((f 1) 2 3) 4))`

	// exp := `(do (set (quote fib)
	// (lambda (i)
	// (if (= i 0)
	//     1
	//     (if (= i 1)
	//         1
	//         (+ (fib (- i 1)) (fib (- i 2)))))))
	// (fib 10))`

	exp := `(do (set (quote sum)
	(lambda (r i)
	  (if (= i 0)
	      r
	      (sum (+ r 1) (- i 1)))))
	(sum 0 20000000))`

	// exp := `(+ (do 1 (do 2 3)) 4)`

	// exp := `(+ (+ (+ 1 2) 3) 4)`
	// exp := `(do (set (quote id) (lambda (x) x)) (id 42))`
	// exp := `(do (set (quote x) 42) x)`
	// exp := `(if true 1 2)`
	// exp := `((lambda (x y z) z) 1 2 3)`
	// exp := `(do 1 2)`
	// exp := `(+ 3 7)`
	// exp := `((+ 1) 2)`
	// exp := `(((+) 1) 2)`

	// PrimNS1Set(MakeSymbol("set"), primitive(myset, 2))

	// PrimNS1Set(MakeSymbol("+"), primitive(myadd, 2))
	// PrimNS1Set(MakeSymbol("*"), primitive(mymul, 2))
	// PrimNS1Set(MakeSymbol("-"), primitive(mysub, 2))
	// PrimNS1Set(MakeSymbol("="), primitive(myeq, 2))

	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println("input is:", ObjString(sexp))
	res := Neval(sexp)
	fmt.Println("output is:", ObjString(res))
}
