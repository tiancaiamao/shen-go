package klambda

var Load = MakeNative(func(__e *ControlFlow) {
	tmp176 := MakeNative(func(__e *ControlFlow) {
		tmp177 := MakeNative(func(__e *ControlFlow) {
			tmp178 := Call(__e, PrimNS2Value(symsimple_1error), MakeString("xxx"))

			__e.TailApply(PrimNS1Value(sym_7), MakeNumber(2), tmp178)
			return

		}, 0)

		tmp179 := MakeNative(func(__e *ControlFlow) {
			e := __e.Get(1)
			_ = e
			__e.Return(MakeNumber(2))
			return
		}, 1)

		tmp180 := Call(__e, PrimNS1Value(symtry_1catch), tmp177, tmp179)

		__e.TailApply(PrimNS1Value(sym_7), MakeNumber(1), tmp180)
		return

	}, 0)

	tmp181 := Call(__e, PrimNS1Value(symset), symf, tmp176)

	_ = tmp181

	tmp182 := MakeNative(func(__e *ControlFlow) {
		n := __e.Get(1)
		_ = n
		tmp185 := Call(__e, PrimNS1Value(sym_a), n, MakeNumber(0))

		if True == tmp185 {
			__e.Return(n)
			return
		} else {
			tmp184 := Call(__e, PrimNS1Value(sym_1), n, MakeNumber(1))

			__e.TailApply(PrimNS1Value(symrecur), tmp184)
			return

		}

	}, 1)

	tmp186 := Call(__e, PrimNS1Value(symset), symrecur, tmp182)

	_ = tmp186

	tmp187 := MakeNative(func(__e *ControlFlow) {
		sum := __e.Get(1)
		_ = sum
		n := __e.Get(2)
		_ = n
		tmp191 := Call(__e, PrimNS1Value(sym_a), n, MakeNumber(0))

		if True == tmp191 {
			__e.Return(sum)
			return
		} else {
			tmp189 := Call(__e, PrimNS1Value(sym_d), sum, n)

			tmp190 := Call(__e, PrimNS1Value(sym_1), n, MakeNumber(1))

			__e.TailApply(PrimNS1Value(symfact0), tmp189, tmp190)
			return

		}

	}, 2)

	tmp192 := Call(__e, PrimNS1Value(symset), symfact0, tmp187)

	_ = tmp192

	tmp193 := MakeNative(func(__e *ControlFlow) {
		n := __e.Get(1)
		_ = n
		__e.TailApply(PrimNS1Value(symfact0), MakeNumber(1), n)
		return
	}, 1)

	tmp194 := Call(__e, PrimNS1Value(symset), symfact, tmp193)

	_ = tmp194

	tmp195 := MakeNative(func(__e *ControlFlow) {
		res := __e.Get(1)
		_ = res
		l := __e.Get(2)
		_ = l
		tmp200 := Call(__e, PrimNS1Value(sym_a), l, Nil)

		if True == tmp200 {
			__e.Return(res)
			return
		} else {
			tmp197 := Call(__e, PrimNS1Value(symcar), l)

			tmp198 := Call(__e, PrimNS1Value(symcons), tmp197, res)

			tmp199 := Call(__e, PrimNS1Value(symcdr), l)

			__e.TailApply(PrimNS1Value(symreverse_1help), tmp198, tmp199)
			return

		}

	}, 2)

	tmp201 := Call(__e, PrimNS1Value(symset), symreverse_1help, tmp195)

	_ = tmp201

	tmp202 := MakeNative(func(__e *ControlFlow) {
		l := __e.Get(1)
		_ = l
		__e.TailApply(PrimNS1Value(symreverse_1help), Nil, l)
		return
	}, 1)

	tmp203 := Call(__e, PrimNS1Value(symset), symreverse, tmp202)

	_ = tmp203

	tmp204 := MakeNative(func(__e *ControlFlow) {
		res := __e.Get(1)
		_ = res
		f := __e.Get(2)
		_ = f
		l := __e.Get(3)
		_ = l
		tmp210 := Call(__e, PrimNS1Value(symcons_2), l)

		if True == tmp210 {
			tmp206 := Call(__e, PrimNS1Value(symcar), l)

			tmp207 := Call(__e, f, tmp206)

			tmp208 := Call(__e, PrimNS1Value(symcons), tmp207, res)

			tmp209 := Call(__e, PrimNS1Value(symcdr), l)

			__e.TailApply(PrimNS1Value(symmap_1help), tmp208, f, tmp209)
			return

		} else {
			__e.TailApply(PrimNS1Value(symreverse), res)
			return
		}

	}, 3)

	tmp211 := Call(__e, PrimNS1Value(symset), symmap_1help, tmp204)

	_ = tmp211

	tmp212 := MakeNative(func(__e *ControlFlow) {
		f := __e.Get(1)
		_ = f
		l := __e.Get(2)
		_ = l
		__e.TailApply(PrimNS1Value(symmap_1help), Nil, f, l)
		return
	}, 2)

	__e.TailApply(PrimNS1Value(symset), symmap, tmp212)
	return

}, 0)

var symcons = MakeSymbol("cons")
var symreverse_1help = MakeSymbol("reverse-help")
var symmap = MakeSymbol("map")
var symf = MakeSymbol("f")
var sym_a = MakeSymbol("=")
var symrecur = MakeSymbol("recur")
var sym_d = MakeSymbol("*")
var symfact0 = MakeSymbol("fact0")
var symcar = MakeSymbol("car")
var symsimple_1error = MakeSymbol("simple-error")
var symtry_1catch = MakeSymbol("try-catch")
var symfact = MakeSymbol("fact")
var sym_7 = MakeSymbol("+")
var symset = MakeSymbol("set")
var symreverse = MakeSymbol("reverse")
var symcons_2 = MakeSymbol("cons?")
var symmap_1help = MakeSymbol("map-help")
var sym_1 = MakeSymbol("-")
var symcdr = MakeSymbol("cdr")
