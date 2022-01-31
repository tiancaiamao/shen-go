package cora

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrampoline(t *testing.T) {
	var kl ControlFlow
	Call(&kl, Load)
	// Make sure no OOM in the tail apply case by using Trampoine
	recur := PrimNS1Value(MakeSymbol("recur"))
	_ = Call(&kl, recur, MakeNumber(100000))
}

func TestFact(t *testing.T) {
	var kl ControlFlow
	Call(&kl, Load)
	fn := PrimNS1Value(MakeSymbol("fact"))
	res := Call(&kl, fn, MakeInteger(5))
	n := mustInteger(res)
	if n != 120 {
		t.Fail()
	}
}

func TestPartialApply(t *testing.T) {
	var kl ControlFlow
	Call(&kl, Load)
	fn := PrimNS1Value(MakeSymbol("fact0"))
	fn1 := Call(&kl, fn, MakeInteger(1))
	res := Call(&kl, fn1, MakeInteger(5))
	if mustInteger(res) != 120 {
		t.Fail()
	}
}

func TestTryCatch(t *testing.T) {
	var kl ControlFlow
	klInit0(&kl)
	Call(&kl, Load)
	// (trap-error (+ 2 (simple-error "xxx")) (lambda X (error-to-string X)))
	exp := `(try-catch (lambda () (+ 2 ((fn (quote simple-error)) "xxx"))) (lambda (X) ((fn (quote error-to-string)) X)))`
	r := NewSexpReader(strings.NewReader(exp), false)
	sexp, err := r.Read()
	if err != nil {
		t.Error(err)
	}
	res := Eval(&kl, sexp)
	if mustString(res) != "xxx" {
		fmt.Println(ObjString(res))
		t.Fail()
	}

	res = Eval(&kl, cons(MakeSymbol("f"), Nil))
	if mustInteger(res) != 3 {
		t.Fail()
	}
}

func TestFusion(t *testing.T) {
	var kl ControlFlow
	Call(&kl, Load)
	expect := Cons(MakeNumber(2), Cons(MakeNumber(3), Cons(MakeNumber(4), Nil)))

	// Partial apply, generated code and intepreted code fusion!!
	cases := []string{
		"(map (lambda (x) (+ x 1)) (cons 1 (cons 2 (cons 3 ()))))",
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
