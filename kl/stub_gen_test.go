package kl

var Load = MakeNative(func(__e Evaluator) {
	gen5 := MakeNative(func(__e Evaluator) {
		gen1 := MakeNative(func(__e Evaluator) {
			e := __e.Get(1)
			_ = e
			__e.Return(MakeNumber(2))
			return
		}, 1)
		gen3 := MakeNative(func(__e Evaluator) {
			gen2 := Call(__e, __e.Global(symsimple_1error), MakeString("xxx"))

			__e.TailApply(__e.Global(sym_7), MakeNumber(2), gen2)

			return

		}, 0)
		gen4 := Try(__e, gen3).Catch(gen1)
		__e.TailApply(__e.Global(sym_7), MakeNumber(1), gen4)

		return

	}, 0)
	Call(__e, __e.Global(symdefun), MakeSymbol("f"), gen5)

	gen8 := MakeNative(func(__e Evaluator) {
		n := __e.Get(1)
		_ = n
		gen7 := Call(__e, __e.Global(sym_a), n, MakeNumber(0))

		if True == gen7 {
			__e.Return(n)
			return
		} else {
			gen6 := Call(__e, __e.Global(sym_1), n, MakeNumber(1))

			__e.TailApply(__e.Global(symrecur), gen6)

			return

		}

	}, 1)
	Call(__e, __e.Global(symdefun), MakeSymbol("recur"), gen8)

	gen12 := MakeNative(func(__e Evaluator) {
		sum := __e.Get(1)
		_ = sum
		n := __e.Get(2)
		_ = n
		gen11 := Call(__e, __e.Global(sym_a), n, MakeNumber(0))

		if True == gen11 {
			__e.Return(sum)
			return
		} else {
			gen9 := Call(__e, __e.Global(sym_d), sum, n)

			gen10 := Call(__e, __e.Global(sym_1), n, MakeNumber(1))

			__e.TailApply(__e.Global(symfact0), gen9, gen10)

			return

		}

	}, 2)
	Call(__e, __e.Global(symdefun), MakeSymbol("fact0"), gen12)

	gen13 := MakeNative(func(__e Evaluator) {
		n := __e.Get(1)
		_ = n
		__e.TailApply(__e.Global(symfact0), MakeNumber(1), n)

		return
	}, 1)
	Call(__e, __e.Global(symdefun), MakeSymbol("fact"), gen13)

	gen18 := MakeNative(func(__e Evaluator) {
		res := __e.Get(1)
		_ = res
		l := __e.Get(2)
		_ = l
		gen17 := Call(__e, __e.Global(sym_a), l, Nil)

		if True == gen17 {
			__e.Return(res)
			return
		} else {
			gen14 := Call(__e, __e.Global(symhd), l)

			gen15 := Call(__e, __e.Global(symcons), gen14, res)

			gen16 := Call(__e, __e.Global(symtl), l)

			__e.TailApply(__e.Global(symreverse_1help), gen15, gen16)

			return

		}

	}, 2)
	Call(__e, __e.Global(symdefun), MakeSymbol("reverse-help"), gen18)

	gen19 := MakeNative(func(__e Evaluator) {
		l := __e.Get(1)
		_ = l
		__e.TailApply(__e.Global(symreverse_1help), Nil, l)

		return
	}, 1)
	Call(__e, __e.Global(symdefun), MakeSymbol("reverse"), gen19)

	gen25 := MakeNative(func(__e Evaluator) {
		res := __e.Get(1)
		_ = res
		f := __e.Get(2)
		_ = f
		l := __e.Get(3)
		_ = l
		gen24 := Call(__e, __e.Global(symcons_2), l)

		if True == gen24 {
			gen20 := Call(__e, __e.Global(symhd), l)

			gen21 := Call(__e, f, gen20)

			gen22 := Call(__e, __e.Global(symcons), gen21, res)

			gen23 := Call(__e, __e.Global(symtl), l)

			__e.TailApply(__e.Global(symmap_1help), gen22, f, gen23)

			return

		} else {
			__e.TailApply(__e.Global(symreverse), res)

			return
		}

	}, 3)
	Call(__e, __e.Global(symdefun), MakeSymbol("map-help"), gen25)

	gen26 := MakeNative(func(__e Evaluator) {
		f := __e.Get(1)
		_ = f
		l := __e.Get(2)
		_ = l
		__e.TailApply(__e.Global(symmap_1help), Nil, f, l)

		return
	}, 2)
	__e.TailApply(__e.Global(symdefun), MakeSymbol("map"), gen26)

	return

}, 0)

var sym_a = MakeSymbol("=")
var sym_1 = MakeSymbol("-")
var symreverse_1help = MakeSymbol("reverse-help")
var symsimple_1error = MakeSymbol("simple-error")
var sym_7 = MakeSymbol("+")
var symrecur = MakeSymbol("recur")
var sym_d = MakeSymbol("*")
var symmap_1help = MakeSymbol("map-help")
var symreverse = MakeSymbol("reverse")
var symfact0 = MakeSymbol("fact0")
var symtl = MakeSymbol("tl")
var symcons_2 = MakeSymbol("cons?")
var symdefun = MakeSymbol("defun")
var symhd = MakeSymbol("hd")
var symcons = MakeSymbol("cons")
