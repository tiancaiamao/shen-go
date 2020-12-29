package kl

import (
	// "fmt"
	"strings"
	"testing"
)

func TestTrampoline(t *testing.T) {
	var kl KLambda
	// Make sure no OOM in the tail apply case by using Trampoine
	recur := ShenFunc(MakeSymbol("recur"))
	_ = Call(&kl, recur, MakeNumber(100000))
}

func TestFact(t *testing.T) {
	var kl KLambda
	fn := ShenFunc(MakeSymbol("fact"))
	res := Call(&kl, fn, MakeInteger(5))
	n := mustInteger(res)
	if n != 120 {
		t.Fail()
	}
}

// func TestPartialApply(t *testing.T) {
// 	var kl KLambda
// 	fn := ShenFunc(MakeSymbol("fact0"))
// 	res := Call(&kl, MakeNative(func(_e Evaluator, args ...Obj) {
// 		_e.Return(fn)
// 		return
// 	}, 0), MakeInteger(5))
// 	if mustInteger(res) != 120 {
// 		t.Fail()
// 	}
// }

var __initExprs []Obj

func TestTryCatch(t *testing.T) {
	var kl KLambda
	// (trap-error (+ 2 (simple-error "xxx")) (lambda X (error-to-string X)))
	// res := Try(e, MakeNative(func(_e Evaluator, args ...Obj) {
	// 	regXX := MakeString("xxx")
	// 	simpleErr := ShenFunc(MakeSymbol("simple-error"))
	// 	regYY := Call(e, simpleErr, regXX)
	// 	res := PrimNumberAdd(MakeNumber(2), regYY)
	// 	_e.Return(res)
	// 	return
	// }, 0)).Catch(MakeNative(func(_e Evaluator, args ...Obj) {
	// 	err := args[0]
	// 	res := PrimErrorToString(err)
	// 	_e.Return(res)
	// 	return
	// }, 1))

	// if mustString(res) != "xxx" {
	// 	t.Fail()
	// }

	// res = Call(e, __defun__trycatch)
	// if mustInteger(res) != 3 {
	// 	t.Fail()
	// }

	// res = Try(e, MakeNative(func(_e Evaluator, args ...Obj) {
	// 	_e.Return(MakeNumber(42))
	// 	return
	// }, 0)).Catch(MakeNative(func(_e Evaluator, args ...Obj) {
	// 	err := args[0]
	// 	res := PrimErrorToString(err)
	// 	_e.Return(res)
	// 	return
	// }, 1))
	// if mustInteger(res) != 42 {
	// 	t.Fail()
	// }

	res := Eval(&kl, cons(MakeSymbol("f"), Nil))
	if mustInteger(res) != 3 {
		t.Fail()
	}
}

func TestFusion(t *testing.T) {
	var kl KLambda
	// e.RegistNativeCall("map", __defun__map)
	expect := Cons(MakeNumber(2), Cons(MakeNumber(3), Cons(MakeNumber(4), Nil)))

	// Partial apply, generated code and intepreted code fusion!!
	cases := []string{
		"(map (lambda x (+ x 1)) (cons 1 (cons 2 (cons 3 ()))))",
		"(map (+ 1) (cons 1 (cons 2 (cons 3 ()))))",
		"((map (+ 1)) (cons 1 (cons 2 (cons 3 ()))))",
		"((map) (+ 1) (cons 1 (cons 2 (cons 3 ()))))",
	}

	for _, input := range cases {
		r := NewSexpReader(strings.NewReader(input), false)
		sexp, err := r.Read()
		if err != nil {
			t.Error(err)
		}

		obtain := Eval(&kl, sexp)
		if equal(expect, obtain) == False {
			t.Fail()
		}
	}
}
