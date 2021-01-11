package kl

var Load = MakeNative(func(__e *ControlFlow) {
	gen99 := MakeNative(func(__e *ControlFlow) {
		gen96 := MakeNative(func(__e *ControlFlow) {
			gen95 := Call(__e, PrimNS1Value(symsimple_1error), MakeString("xxx"))

			__e.TailApply(PrimNS1Value(sym_7), MakeNumber(2), gen95)

			return

		}, 0)

		gen97 := MakeNative(func(__e *ControlFlow) {
			e := __e.Get(1)
			_ = e
			__e.Return(MakeNumber(2))
			return
		}, 1)

		gen98 := Call(__e, PrimNS1Value(symtry_1catch), gen96, gen97)

		__e.TailApply(PrimNS1Value(sym_7), MakeNumber(1), gen98)

		return

	}, 0)

	Call(__e, PrimNS1Value(symset), symf, gen99)

	gen102 := MakeNative(func(__e *ControlFlow) {
		n := __e.Get(1)
		_ = n
		gen101 := Call(__e, PrimNS1Value(sym_a), n, MakeNumber(0))

		if True == gen101 {
			__e.Return(n)
			return
		} else {
			gen100 := Call(__e, PrimNS1Value(sym_1), n, MakeNumber(1))

			__e.TailApply(PrimNS1Value(symrecur), gen100)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symset), symrecur, gen102)

	gen106 := MakeNative(func(__e *ControlFlow) {
		sum := __e.Get(1)
		_ = sum
		n := __e.Get(2)
		_ = n
		gen105 := Call(__e, PrimNS1Value(sym_a), n, MakeNumber(0))

		if True == gen105 {
			__e.Return(sum)
			return
		} else {
			gen103 := Call(__e, PrimNS1Value(sym_d), sum, n)

			gen104 := Call(__e, PrimNS1Value(sym_1), n, MakeNumber(1))

			__e.TailApply(PrimNS1Value(symfact0), gen103, gen104)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symset), symfact0, gen106)

	gen107 := MakeNative(func(__e *ControlFlow) {
		n := __e.Get(1)
		_ = n
		__e.TailApply(PrimNS1Value(symfact0), MakeNumber(1), n)

		return
	}, 1)

	Call(__e, PrimNS1Value(symset), symfact, gen107)

	gen112 := MakeNative(func(__e *ControlFlow) {
		res := __e.Get(1)
		_ = res
		l := __e.Get(2)
		_ = l
		gen111 := Call(__e, PrimNS1Value(sym_a), l, Nil)

		if True == gen111 {
			__e.Return(res)
			return
		} else {
			gen108 := Call(__e, PrimNS1Value(symhd), l)

			gen109 := Call(__e, PrimNS1Value(symcons), gen108, res)

			gen110 := Call(__e, PrimNS1Value(symtl), l)

			__e.TailApply(PrimNS1Value(symreverse_1help), gen109, gen110)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symset), symreverse_1help, gen112)

	gen113 := MakeNative(func(__e *ControlFlow) {
		l := __e.Get(1)
		_ = l
		__e.TailApply(PrimNS1Value(symreverse_1help), Nil, l)

		return
	}, 1)

	Call(__e, PrimNS1Value(symset), symreverse, gen113)

	gen119 := MakeNative(func(__e *ControlFlow) {
		res := __e.Get(1)
		_ = res
		f := __e.Get(2)
		_ = f
		l := __e.Get(3)
		_ = l
		gen118 := Call(__e, PrimNS1Value(symcons_2), l)

		if True == gen118 {
			gen114 := Call(__e, PrimNS1Value(symhd), l)

			gen115 := Call(__e, f, gen114)

			gen116 := Call(__e, PrimNS1Value(symcons), gen115, res)

			gen117 := Call(__e, PrimNS1Value(symtl), l)

			__e.TailApply(PrimNS1Value(symmap_1help), gen116, f, gen117)

			return

		} else {
			__e.TailApply(PrimNS1Value(symreverse), res)

			return
		}

	}, 3)

	Call(__e, PrimNS1Value(symset), symmap_1help, gen119)

	gen120 := MakeNative(func(__e *ControlFlow) {
		f := __e.Get(1)
		_ = f
		l := __e.Get(2)
		_ = l
		__e.TailApply(PrimNS1Value(symmap_1help), Nil, f, l)

		return
	}, 2)

	__e.TailApply(PrimNS1Value(symset), symmap, gen120)

	return

}, 0)
var symcons_2 = MakeSymbol("cons?")
var symf = MakeSymbol("f")
var symrecur = MakeSymbol("recur")
var symfact0 = MakeSymbol("fact0")
var sym_1 = MakeSymbol("-")
var symcons = MakeSymbol("cons")
var symreverse_1help = MakeSymbol("reverse-help")
var symmap = MakeSymbol("map")
var sym_7 = MakeSymbol("+")
var symset = MakeSymbol("set")
var sym_a = MakeSymbol("=")
var symfact = MakeSymbol("fact")
var symtl = MakeSymbol("cdr")
var symsimple_1error = MakeSymbol("simple-error")
var symtry_1catch = MakeSymbol("try-catch")
var sym_d = MakeSymbol("*")
var symhd = MakeSymbol("car")
var symreverse = MakeSymbol("reverse")
var symmap_1help = MakeSymbol("map-help")
