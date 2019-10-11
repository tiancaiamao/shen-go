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

// type Trampoline struct {
// 	pc   *scmFunc
// 	args []Obj
// }

// func (t *Trampoline) Return(res Obj) {
// 	t.pc = nil
// 	t.args = t.args[:0]
// 	t.args = append(t.args, res)
// }

// func (t *Trampoline) TailCall(f Obj, args ...Obj) {
// 	t.pc = MustFunc(f)
// 	t.args = t.args[:0]
// 	t.args = append(t.args, args...)
// }

// func (t *Trampoline) partialApply() {
// 	provided := len(t.pc.captured) + len(t.args)
// 	required := t.pc.require
// 	if provided == required {
// 		t.pc.fn(t, t.args...)
// 		return
// 	}

// 	tmp := make([]Obj, 0, t.pc.require)
// 	tmp = append(tmp, t.pc.captured...)
// 	if provided < required {
// 		tmp = append(tmp, t.args...)
// 		t.Return(MakeFunc(t.pc.fn, t.pc.require, tmp...))
// 		return
// 	}

// 	tmp = append(tmp, t.args[:provided-required]...)
// 	fn := Call(MakeFunc(t.pc.fn, t.pc.require), tmp...)
// 	t.TailCall(fn, t.args[provided-required:]...)
// 	return
// }

// func (t *Trampoline) Run(pc *scmFunc, args ...Obj) Obj {
// 	t.pc = pc
// 	t.args = args
// 	for t.pc != nil {
// 		t.partialApply()
// 	}
// 	return t.args[0]
// }

// func Call(f Obj, args ...Obj) Obj {
// 	fn := MustFunc(f)
// 	var ctx Trampoline
// 	return ctx.Run(fn, args...)
// }

// func BuiltinEQ(x, y Obj) Obj {
// 	return equal(x, y)
// }

// func BuiltinSub(x, y Obj) Obj {
// 	x1 := mustNumber(x)
// 	y1 := mustNumber(y)
// 	return MakeNumber(x1.val - y1.val)
// }

// func BuiltinAdd(x, y Obj) Obj {
// 	x1 := mustNumber(x)
// 	y1 := mustNumber(y)
// 	return MakeNumber(x1.val + y1.val)
// }

// func BuiltinMul(x, y Obj) Obj {
// 	x1 := mustNumber(x)
// 	y1 := mustNumber(y)
// 	return MakeNumber(x1.val * y1.val)
// }

// func BuiltinCons(x, y Obj) Obj {
// 	return cons(x, y)
// }

// func BuiltinHD(x Obj) Obj {
// 	return car(x)
// }

// func BuiltinTL(x Obj) Obj {
// 	return cdr(x)
// }

// func BuiltinConsP(x Obj) Obj {
// 	return primIsPair(x)
// }

// func BuiltinSet(symbol, val Obj) Obj {
// 	sym := mustSymbol(symbol)
// 	symVal := &symbolArray[sym.offset]
// 	symVal.value = val
// 	return val
// }

// func BuiltinValue(arg Obj) Obj {
// 	sym := mustSymbol(arg)
// 	symVal := &symbolArray[sym.offset]
// 	if symVal.value != nil {
// 		return symVal.value
// 	}
// 	return MakeError(fmt.Sprintf("variable %s not bound", symVal.str))
// }
