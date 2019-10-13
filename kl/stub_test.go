package kl

import (
	"testing"
)

func TestTrampoline(t *testing.T) {
	// Make sure no OOM in the tail apply case by using Trampoine
	_ = Call(__defun_recur, MakeNumber(100000))
}

func TestFact(t *testing.T) {
	res := Call(__defun_fact0, MakeInteger(5))
	n := mustInteger(res)
	if n != 120 {
		t.Fail()
	}
}

func TestPartialApply(t *testing.T) {
	fn := Call(__defun_fact0)
	res := Call(fn, MakeInteger(5))
	if mustInteger(res) != 120 {
		t.Fail()
	}

	res = Call(MakeNative(func(ctx *Trampoline, args ...Obj) {
		ctx.Return(__defun_fact0)
		return
	}, 0), MakeInteger(5))
	if mustInteger(res) != 120 {
		t.Fail()
	}
}

func TestNativeCall(t *testing.T) {
	e := NewEvaluator()
	e.RegistNativeCall("fact", 1, __defun_fact0)
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
	__defun_recur = MakeNative(func(ctx *Trampoline, args ...Obj) {
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

	__defun_fact = MakeNative(func(ctx *Trampoline, args ...Obj) {
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

	__defun_fact0 = MakeNative(func(ctx *Trampoline, args ...Obj) {
		res := Call(__defun_fact, MakeNumber(1), args[0])
		ctx.Return(res)
		return
	}, 1)

	// (define trycatch
	//         -> (trap-error (+ 4 (simple-error "xxx")) (/. E (error-to-string E))))
	__defun__trycatch = MakeNative(func(__ctx *Trampoline, __args ...Obj) {
		reg94576 := MakeNative(func(__ctx *Trampoline, __args ...Obj) {
			ignore := __args[0]
			_ = ignore
			reg94577 := MakeNumber(4)
			reg94578 := MakeString("xxx")
			reg94579 := PrimSimpleError(reg94578)
			reg94580 := PrimNumberAdd(reg94577, reg94579)
			__ctx.Return(reg94580)
			return
		}, 1)
		reg94581 := MakeNative(func(__ctx *Trampoline, __args ...Obj) {
			E := __args[0]
			_ = E
			reg94582 := PrimErrorToString(E)
			__ctx.Return(reg94582)
			return
		}, 1)
		reg94575 := Try(reg94576).Catch(reg94581)
		__ctx.Return(reg94575)
		return
	}, 0)
}

func TestTryCatch(t *testing.T) {
	// (trap-error (+ 2 (simple-error "xxx")) (lambda X (error-to-string X)))
	res := Try(MakeNative(func(ctx *Trampoline, args ...Obj) {
		regXX := MakeString("xxx")
		regYY := PrimSimpleError(regXX)
		res := PrimNumberAdd(MakeNumber(2), regYY)
		ctx.Return(res)
		return
	}, 1)).Catch(MakeNative(func(ctx *Trampoline, args ...Obj) {
		err := args[0]
		res := PrimErrorToString(err)
		ctx.Return(res)
		return
	}, 1))

	if mustString(res) != "xxx" {
		t.Fail()
	}

	res = Call(__defun__trycatch)
	if mustString(res) != "xxx" {
		t.Fail()
	}

	res = Try(MakeNative(func(ctx *Trampoline, args ...Obj) {
		ctx.Return(MakeNumber(42))
		return
	}, 1)).Catch(MakeNative(func(ctx *Trampoline, args ...Obj) {
		err := args[0]
		res := PrimErrorToString(err)
		ctx.Return(res)
		return
	}, 1))
	if mustNumber(res).val != 42 {
		t.Fail()
	}
}
