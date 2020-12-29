package kl

var Load = MakeNative(func(__e Evaluator, __args ...Obj) {
	gen5 := MakeNative(func(__e Evaluator, __args ...Obj) {
		gen1 := MakeNative(func(__e Evaluator, __args ...Obj) {
			e := __args[0]
			_ = e
			__e.Return(MakeNumber(2))
			return
		}, 1)
		gen3 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen2 := Call(__e, ShenFunc(symsimple_1error), MakeString("xxx"))

			__e.TailApply(ShenFunc(sym_7), MakeNumber(2), gen2)

			return

		}, 0)
		gen4 := Try(__e, gen3).Catch(gen1)
		__e.TailApply(ShenFunc(sym_7), MakeNumber(1), gen4)

		return

	}, 0)
	Call(__e, ShenFunc(symdefun), MakeSymbol("f"), gen5)

	gen8 := MakeNative(func(__e Evaluator, __args ...Obj) {
		n := __args[0]
		_ = n
		gen7 := Call(__e, ShenFunc(sym_a), n, MakeNumber(0))

		if True == gen7 {
			__e.Return(n)
			return
		} else {
			gen6 := Call(__e, ShenFunc(sym_1), n, MakeNumber(1))

			__e.TailApply(ShenFunc(symrecur), gen6)

			return

		}

	}, 1)
	Call(__e, ShenFunc(symdefun), MakeSymbol("recur"), gen8)

	gen12 := MakeNative(func(__e Evaluator, __args ...Obj) {
		sum := __args[0]
		_ = sum
		n := __args[1]
		_ = n
		gen11 := Call(__e, ShenFunc(sym_a), n, MakeNumber(0))

		if True == gen11 {
			__e.Return(sum)
			return
		} else {
			gen9 := Call(__e, ShenFunc(sym_d), sum, n)

			gen10 := Call(__e, ShenFunc(sym_1), n, MakeNumber(1))

			__e.TailApply(ShenFunc(symfact0), gen9, gen10)

			return

		}

	}, 2)
	Call(__e, ShenFunc(symdefun), MakeSymbol("fact0"), gen12)

	gen13 := MakeNative(func(__e Evaluator, __args ...Obj) {
		n := __args[0]
		_ = n
		__e.TailApply(ShenFunc(symfact0), MakeNumber(1), n)

		return
	}, 1)
	Call(__e, ShenFunc(symdefun), MakeSymbol("fact"), gen13)

	gen18 := MakeNative(func(__e Evaluator, __args ...Obj) {
		res := __args[0]
		_ = res
		l := __args[1]
		_ = l
		gen17 := Call(__e, ShenFunc(sym_a), l, Nil)

		if True == gen17 {
			__e.Return(res)
			return
		} else {
			gen14 := Call(__e, ShenFunc(symhd), l)

			gen15 := Call(__e, ShenFunc(symcons), gen14, res)

			gen16 := Call(__e, ShenFunc(symtl), l)

			__e.TailApply(ShenFunc(symreverse_1help), gen15, gen16)

			return

		}

	}, 2)
	Call(__e, ShenFunc(symdefun), MakeSymbol("reverse-help"), gen18)

	gen19 := MakeNative(func(__e Evaluator, __args ...Obj) {
		l := __args[0]
		_ = l
		__e.TailApply(ShenFunc(symreverse_1help), Nil, l)

		return
	}, 1)
	Call(__e, ShenFunc(symdefun), MakeSymbol("reverse"), gen19)

	gen25 := MakeNative(func(__e Evaluator, __args ...Obj) {
		res := __args[0]
		_ = res
		f := __args[1]
		_ = f
		l := __args[2]
		_ = l
		gen24 := Call(__e, ShenFunc(symcons_2), l)

		if True == gen24 {
			gen20 := Call(__e, ShenFunc(symhd), l)

			gen21 := Call(__e, f, gen20)

			gen22 := Call(__e, ShenFunc(symcons), gen21, res)

			gen23 := Call(__e, ShenFunc(symtl), l)

			__e.TailApply(ShenFunc(symmap_1help), gen22, f, gen23)

			return

		} else {
			__e.TailApply(ShenFunc(symreverse), res)

			return
		}

	}, 3)
	Call(__e, ShenFunc(symdefun), MakeSymbol("map-help"), gen25)

	gen26 := MakeNative(func(__e Evaluator, __args ...Obj) {
		f := __args[0]
		_ = f
		l := __args[1]
		_ = l
		__e.TailApply(ShenFunc(symmap_1help), Nil, f, l)

		return
	}, 2)
	__e.TailApply(ShenFunc(symdefun), MakeSymbol("map"), gen26)

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
