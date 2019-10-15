package kl

import (
	"strings"
	"testing"
)

func TestTrampoline(t *testing.T) {
	// Make sure no OOM in the tail apply case by using Trampoine
	_ = NewEvaluator().Call(__defun_recur, MakeNumber(100000))
}

func TestFact(t *testing.T) {
	res := NewEvaluator().Call(__defun_fact0, MakeInteger(5))
	n := mustInteger(res)
	if n != 120 {
		t.Fail()
	}
}

func TestPartialApply(t *testing.T) {
	e := NewEvaluator()
	fn := e.Call(__defun_fact0)
	res := e.Call(fn, MakeInteger(5))
	if mustInteger(res) != 120 {
		t.Fail()
	}

	res = e.Call(MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		ctx.Return(__defun_fact0)
		return
	}, 0), MakeInteger(5))
	if mustInteger(res) != 120 {
		t.Fail()
	}
}

func TestNativeCall(t *testing.T) {
	e := NewEvaluator()
	e.RegistNativeCall("fact", __defun_fact0)
	res := e.Eval(cons(MakeSymbol("fact"), cons(MakeInteger(5), Nil)))
	if mustInteger(res) != 120 {
		t.Fail()
	}
}

var __defun_fact Obj
var __defun_fact0 Obj
var __defun_recur Obj
var __defun__trycatch Obj
var __defun__reverse Obj
var __defun__shen_4reverse__help Obj
var __defun__map Obj
var __defun__shen_4map_1h Obj

func init() {
	__defun_recur = MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		n := args[0]
		if equal(n, MakeInteger(0)) == True {
			ctx.Return(n)
			return
		} else {
			n = PrimNumberSubtract(n, MakeInteger(1))
			ctx.TailApply(__defun_recur, n)
			return
		}
	}, 1)

	__defun_fact = MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		sum := args[0]
		n := args[1]
		if equal(n, MakeInteger(0)) == True {
			ctx.Return(sum)
			return
		} else {
			sum = PrimNumberMultiply(sum, n)
			n = PrimNumberSubtract(n, MakeInteger(1))
			ctx.TailApply(__defun_fact, sum, n)
			return
		}
	}, 2)

	__defun_fact0 = MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		res := _e.Call(__defun_fact, MakeNumber(1), args[0])
		ctx.Return(res)
		return
	}, 1)

	// (define trycatch
	//         -> (trap-error (+ 4 (simple-error "xxx")) (/. E (error-to-string E))))
	__defun__trycatch = MakeNative(func(___e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg94576 := MakeNative(func(___e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
			ignore := __args[0]
			_ = ignore
			reg94577 := MakeNumber(4)
			reg94578 := MakeString("xxx")
			reg94579 := PrimSimpleError(reg94578)
			reg94580 := PrimNumberAdd(reg94577, reg94579)
			__ctx.Return(reg94580)
			return
		}, 1)
		reg94581 := MakeNative(func(___e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg94582 := PrimErrorToString(E)
			__ctx.Return(reg94582)
			return
		}, 1)
		reg94575 := ___e.Try(reg94576).Catch(reg94581)
		__ctx.Return(reg94575)
		return
	}, 0)

	__defun__map = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3161 := __args[0]
		_ = V3161
		V3162 := __args[1]
		_ = V3162
		reg100263 := Nil
		__ctx.TailApply(__defun__shen_4map_1h, V3161, V3162, reg100263)
		return
	}, 2)
	__defun__shen_4map_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3168 := __args[0]
		_ = V3168
		V3169 := __args[1]
		_ = V3169
		V3170 := __args[2]
		_ = V3170
		reg100265 := Nil
		reg100266 := PrimEqual(reg100265, V3169)
		if reg100266 == True {
			__ctx.TailApply(__defun__reverse, V3170)
			return
		} else {
			reg100268 := PrimIsPair(V3169)
			if reg100268 == True {
				reg100269 := PrimTail(V3169)
				reg100270 := PrimHead(V3169)
				reg100271 := __e.Call(V3168, reg100270)
				reg100272 := PrimCons(reg100271, V3170)
				__ctx.TailApply(__defun__shen_4map_1h, V3168, reg100269, reg100272)
				return
			}
		}
	}, 3)
	__defun__reverse = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3121 := __args[0]
		_ = V3121
		reg100177 := Nil
		__ctx.TailApply(__defun__shen_4reverse__help, V3121, reg100177)
		return
	}, 1)
	__defun__shen_4reverse__help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3124 := __args[0]
		_ = V3124
		V3125 := __args[1]
		_ = V3125
		reg100179 := Nil
		reg100180 := PrimEqual(reg100179, V3124)
		if reg100180 == True {
			__ctx.Return(V3125)
			return
		} else {
			reg100181 := PrimIsPair(V3124)
			if reg100181 == True {
				reg100182 := PrimTail(V3124)
				reg100183 := PrimHead(V3124)
				reg100184 := PrimCons(reg100183, V3125)
				__ctx.TailApply(__defun__shen_4reverse__help, reg100182, reg100184)
				return
			}
		}
	}, 2)
}

func TestTryCatch(t *testing.T) {
	e := NewEvaluator()
	// (trap-error (+ 2 (simple-error "xxx")) (lambda X (error-to-string X)))
	res := e.Try(MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		regXX := MakeString("xxx")
		regYY := PrimSimpleError(regXX)
		res := PrimNumberAdd(MakeNumber(2), regYY)
		ctx.Return(res)
		return
	}, 1)).Catch(MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		err := args[0]
		res := PrimErrorToString(err)
		ctx.Return(res)
		return
	}, 1))

	if mustString(res) != "xxx" {
		t.Fail()
	}

	res = e.Call(__defun__trycatch)
	if mustString(res) != "xxx" {
		t.Fail()
	}

	res = e.Try(MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		ctx.Return(MakeNumber(42))
		return
	}, 1)).Catch(MakeNative(func(_e *Evaluator, ctx *ControlFlow, args ...Obj) {
		err := args[0]
		res := PrimErrorToString(err)
		ctx.Return(res)
		return
	}, 1))
	if mustNumber(res).val != 42 {
		t.Fail()
	}
}

func TestFusion(t *testing.T) {
	e := NewEvaluator()
	e.RegistNativeCall("map", __defun__map)
	expect := Cons(MakeNumber(2), Cons(MakeNumber(3), Cons(MakeNumber(4), Nil)))

	// Partial apply, generated code and intepreted code fusion!!
	cases := []string{
		"(map (lambda x (+ x 1)) (cons 1 (cons 2 (cons 3 ()))))",
		"(map (+ 1) (cons 1 (cons 2 (cons 3 ()))))",
		"((map (+ 1)) (cons 1 (cons 2 (cons 3 ()))))",
		"((map) (+ 1) (cons 1 (cons 2 (cons 3 ()))))",
	}

	for _, input := range cases {
		r := NewSexpReader(strings.NewReader(input))
		sexp, err := r.Read()
		if err != nil {
			t.Error(err)
		}

		obtain := e.Eval(sexp)
		if equal(expect, obtain) == False {
			t.Fail()
		}
	}
}
