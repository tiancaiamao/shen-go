package kl

import (
	"unsafe"
)

type scmNative struct {
	scmHead
	fn       func(*Trampoline, ...Obj)
	require  int
	captured []Obj
}

func MakeNative(fn func(*Trampoline, ...Obj), require int, captured ...Obj) Obj {
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

// Trampoline are used by native code to implement tail call.
type Trampoline struct {
	kind controlFlowKind
	// arguments for apply
	f    Obj
	args []Obj
	// return result
	result Obj
}

func (t *Trampoline) TailApply(f Obj, args ...Obj) {
	t.f = f
	t.args = args
	t.kind = controlFlowApply
}

func (t *Trampoline) Return(result Obj) {
	t.result = result
	t.kind = controlFlowReturn
}

// Call is used in native function for non-tail-call.
// Similar to Evaluator.trampoline, but not as general as it.
func Call(f Obj, args ...Obj) Obj {
	var t = Trampoline{
		kind: controlFlowApply,
		f:    f,
		args: args,
	}
	for t.kind != controlFlowReturn {
		fn := MustNative(t.f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided == required {
			fn.fn(&t, t.args...)
			continue
		}

		tmp := make([]Obj, 0, fn.require)
		tmp = append(tmp, fn.captured...)
		if provided < required {
			tmp = append(tmp, t.args...)
			return MakeNative(fn.fn, fn.require, tmp...)
		}

		taken := required - len(tmp)
		tmp = append(tmp, t.args[:taken]...)
		t.f = Call(t.f, tmp...)
		t.args = t.args[taken:]
	}
	return t.result
}

type tryResult struct {
	data Obj
}

func Try(f Obj) (res tryResult) {
	defer func() {
		if err := recover(); err != nil {
			val := err.(Obj)
			res = tryResult{data: val}
		}
	}()
	MustNative(f)
	val := Call(f, Nil)
	res = tryResult{data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {
	if IsError(t.data) {
		return Call(f, t.data)
	}
	return t.data
}

func PrimSimpleError(args ...Obj) Obj {
	str := mustString(args[0])
	panic(MakeError(str))
}
