package kl

import (
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
