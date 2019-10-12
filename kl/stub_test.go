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

func init() {
	__defun_recur = MakeNative(func(ctx *Trampoline, args ...Obj) {
		n := args[0]
		if equal(n, MakeInteger(0)) == True {
			ctx.Return(n)
			return
		} else {
			n = primNumberSubtract(n, MakeInteger(1))
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
			sum = primNumberMultiply(sum, n)
			n = primNumberSubtract(n, MakeInteger(1))
			ctx.TailApply(__defun_fact, sum, n)
			return
		}
	}, 2)

	__defun_fact0 = MakeNative(func(ctx *Trampoline, args ...Obj) {
		res := Call(__defun_fact, MakeNumber(1), args[0])
		ctx.Return(res)
		return
	}, 1)
}
