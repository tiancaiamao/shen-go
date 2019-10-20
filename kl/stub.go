package kl

import (
	"unsafe"
)

type scmNative struct {
	scmHead
	fn       func(*Evaluator, *ControlFlow, ...Obj)
	require  int
	captured []Obj
}

func MakeNative(fn func(*Evaluator, *ControlFlow, ...Obj), require int, captured ...Obj) Obj {
	tmp := scmNative{
		scmHead:  scmHeadNative,
		fn:       fn,
		require:  require,
		captured: captured,
	}
	return &tmp.scmHead
}

func MustNative(o Obj) *scmNative {
	if *o != scmHeadNative {
		panic("mustNative")
	}
	return (*scmNative)(unsafe.Pointer(o))
}

type tryResult struct {
	e    *Evaluator
	data Obj
}

func (e *Evaluator) Try(f Obj) (res tryResult) {
	defer func() {
		if err := recover(); err != nil {
			val := err.(Obj)
			res = tryResult{e: e, data: val}
		}
	}()
	MustNative(f)
	val := e.Call(f)
	res = tryResult{e: e, data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {
	if IsError(t.data) {
		return t.e.Call(f, t.data)
	}
	return t.data
}
