package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp68096 := MakeNative(func(__e *ControlFlow) {
		V451 := __e.Get(1)
		_ = V451
		tmp68124 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68125 := Call(__e, tmp68124, Nil, V451)

		if True == tmp68125 {
			__e.Return(Nil)
			return
		} else {
			tmp68122 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68123 := Call(__e, tmp68122, V451)

			var ifres68116 Obj

			if True == tmp68123 {
				tmp68118 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68119 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68120 := Call(__e, tmp68119, V451)

				tmp68121 := Call(__e, tmp68118, tmp68120)

				var ifres68117 Obj

				if True == tmp68121 {
					ifres68117 = True

				} else {
					ifres68117 = False

				}

				ifres68116 = ifres68117

			} else {
				ifres68116 = False

			}

			if True == ifres68116 {
				tmp68100 := MakeNative(func(__e *ControlFlow) {
					DecArity := __e.Get(1)
					_ = DecArity
					tmp68101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4initialise__arity__table)

					tmp68102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp68104 := Call(__e, tmp68103, V451)

					tmp68105 := Call(__e, tmp68102, tmp68104)

					__e.TailApply(tmp68101, tmp68105)
					return

				}, 1)

				tmp68106 := Call(__e, PrimNS1Value(symns2_1value), symput)

				tmp68107 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68108 := Call(__e, tmp68107, V451)

				tmp68109 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68110 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68111 := Call(__e, tmp68110, V451)

				tmp68112 := Call(__e, tmp68109, tmp68111)

				tmp68113 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp68114 := Call(__e, tmp68113, sym_dproperty_1vector_d)

				tmp68115 := Call(__e, tmp68106, tmp68108, symarity, tmp68112, tmp68114)

				__e.TailApply(tmp68100, tmp68115)
				return

			} else {
				tmp68099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp68099, symshen_4initialise__arity__table)
				return

			}

		}

	}, 1)

	tmp68126 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__arity__table, tmp68096)

	_ = tmp68126

	tmp68127 := MakeNative(func(__e *ControlFlow) {
		V453 := __e.Get(1)
		_ = V453
		tmp68128 := MakeNative(func(__e *ControlFlow) {
			tmp68129 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp68130 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp68131 := Call(__e, tmp68130, sym_dproperty_1vector_d)

			__e.TailApply(tmp68129, V453, symarity, tmp68131)
			return

		}, 0)

		tmp68132 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(-1))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp68128, tmp68132)
		return

	}, 1)

	tmp68133 := Call(__e, PrimNS1Value(symns2_1set), symarity, tmp68127)

	_ = tmp68133

	tmp68134 := MakeNative(func(__e *ControlFlow) {
		V455 := __e.Get(1)
		_ = V455
		tmp68135 := MakeNative(func(__e *ControlFlow) {
			Shen := __e.Get(1)
			_ = Shen
			tmp68136 := MakeNative(func(__e *ControlFlow) {
				External := __e.Get(1)
				_ = External
				tmp68137 := MakeNative(func(__e *ControlFlow) {
					Place := __e.Get(1)
					_ = Place
					__e.Return(V455)
					return
				}, 1)

				tmp68138 := Call(__e, PrimNS1Value(symns2_1value), symput)

				tmp68139 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

				tmp68140 := Call(__e, tmp68139, V455, External)

				tmp68141 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp68142 := Call(__e, tmp68141, sym_dproperty_1vector_d)

				tmp68143 := Call(__e, tmp68138, Shen, symshen_4external_1symbols, tmp68140, tmp68142)

				__e.TailApply(tmp68137, tmp68143)
				return

			}, 1)

			tmp68144 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp68145 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp68146 := Call(__e, tmp68145, sym_dproperty_1vector_d)

			tmp68147 := Call(__e, tmp68144, Shen, symshen_4external_1symbols, tmp68146)

			__e.TailApply(tmp68136, tmp68147)
			return

		}, 1)

		tmp68148 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp68149 := Call(__e, tmp68148, MakeString("shen"))

		__e.TailApply(tmp68135, tmp68149)
		return

	}, 1)

	tmp68150 := Call(__e, PrimNS1Value(symns2_1set), symsystemf, tmp68134)

	_ = tmp68150

	tmp68151 := MakeNative(func(__e *ControlFlow) {
		V458 := __e.Get(1)
		_ = V458
		V459 := __e.Get(2)
		_ = V459
		tmp68154 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp68155 := Call(__e, tmp68154, V458, V459)

		if True == tmp68155 {
			__e.Return(V459)
			return
		} else {
			tmp68153 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp68153, V458, V459)
			return

		}

	}, 2)

	tmp68156 := Call(__e, PrimNS1Value(symns2_1set), symadjoin, tmp68151)

	_ = tmp68156

	tmp68157 := MakeNative(func(__e *ControlFlow) {
		V461 := __e.Get(1)
		_ = V461
		tmp68178 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68179 := Call(__e, tmp68178, sympackage, V461)

		if True == tmp68179 {
			__e.Return(Nil)
			return
		} else {
			tmp68176 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68177 := Call(__e, tmp68176, symreceive, V461)

			if True == tmp68177 {
				__e.Return(Nil)
				return
			} else {
				tmp68160 := MakeNative(func(__e *ControlFlow) {
					ArityF := __e.Get(1)
					_ = ArityF
					tmp68172 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp68173 := Call(__e, tmp68172, ArityF, MakeNumber(-1))

					if True == tmp68173 {
						__e.Return(Nil)
						return
					} else {
						tmp68170 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp68171 := Call(__e, tmp68170, ArityF, MakeNumber(0))

						if True == tmp68171 {
							__e.Return(Nil)
							return
						} else {
							tmp68163 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp68164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp68165 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

							tmp68166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1form)

							tmp68167 := Call(__e, tmp68166, V461, ArityF)

							tmp68168 := Call(__e, tmp68165, tmp68167)

							tmp68169 := Call(__e, tmp68164, V461, tmp68168)

							__e.TailApply(tmp68163, tmp68169, Nil)
							return

						}

					}

				}, 1)

				tmp68174 := Call(__e, PrimNS1Value(symns2_1value), symarity)

				tmp68175 := Call(__e, tmp68174, V461)

				__e.TailApply(tmp68160, tmp68175)
				return

			}

		}

	}, 1)

	tmp68180 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form_1entry, tmp68157)

	_ = tmp68180

	tmp68181 := MakeNative(func(__e *ControlFlow) {
		V464 := __e.Get(1)
		_ = V464
		V465 := __e.Get(2)
		_ = V465
		tmp68197 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68198 := Call(__e, tmp68197, MakeNumber(0), V465)

		if True == tmp68198 {
			__e.Return(V464)
			return
		} else {
			tmp68183 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp68184 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68185 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68186 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68187 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1form)

				tmp68188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1end)

				tmp68189 := Call(__e, tmp68188, V464, X)

				tmp68190 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp68191 := Call(__e, tmp68190, V465, MakeNumber(1))

				tmp68192 := Call(__e, tmp68187, tmp68189, tmp68191)

				tmp68193 := Call(__e, tmp68186, tmp68192, Nil)

				tmp68194 := Call(__e, tmp68185, X, tmp68193)

				__e.TailApply(tmp68184, symlambda, tmp68194)
				return

			}, 1)

			tmp68195 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp68196 := Call(__e, tmp68195, symV)

			__e.TailApply(tmp68183, tmp68196)
			return

		}

	}, 2)

	tmp68199 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form, tmp68181)

	_ = tmp68199

	tmp68200 := MakeNative(func(__e *ControlFlow) {
		V468 := __e.Get(1)
		_ = V468
		V469 := __e.Get(2)
		_ = V469
		tmp68208 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68209 := Call(__e, tmp68208, V468)

		if True == tmp68209 {
			tmp68205 := Call(__e, PrimNS1Value(symns2_1value), symappend)

			tmp68206 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68207 := Call(__e, tmp68206, V469, Nil)

			__e.TailApply(tmp68205, V468, tmp68207)
			return

		} else {
			tmp68202 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68203 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp68204 := Call(__e, tmp68203, V469, Nil)

			__e.TailApply(tmp68202, V468, tmp68204)
			return

		}

	}, 2)

	tmp68210 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1end, tmp68200)

	_ = tmp68210

	tmp68211 := MakeNative(func(__e *ControlFlow) {
		V471 := __e.Get(1)
		_ = V471
		tmp68221 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68222 := Call(__e, tmp68221, V471)

		if True == tmp68222 {
			tmp68214 := Call(__e, PrimNS1Value(symns2_1value), symput)

			tmp68215 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp68216 := Call(__e, tmp68215, V471)

			tmp68217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68218 := Call(__e, tmp68217, V471)

			tmp68219 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp68220 := Call(__e, tmp68219, sym_dproperty_1vector_d)

			__e.TailApply(tmp68214, tmp68216, symshen_4lambda_1form, tmp68218, tmp68220)
			return

		} else {
			tmp68213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp68213, symshen_4set_1lambda_1form_1entry)
			return

		}

	}, 1)

	tmp68223 := Call(__e, PrimNS1Value(symns2_1set), symshen_4set_1lambda_1form_1entry, tmp68211)

	_ = tmp68223

	tmp68224 := MakeNative(func(__e *ControlFlow) {
		V473 := __e.Get(1)
		_ = V473
		tmp68225 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp68226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp68227 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp68228 := Call(__e, tmp68227, symshen_4_dspecial_d)

		tmp68229 := Call(__e, tmp68226, V473, tmp68228)

		tmp68230 := Call(__e, tmp68225, symshen_4_dspecial_d, tmp68229)

		_ = tmp68230

		__e.Return(V473)
		return

	}, 1)

	tmp68231 := Call(__e, PrimNS1Value(symns2_1set), symspecialise, tmp68224)

	_ = tmp68231

	tmp68232 := MakeNative(func(__e *ControlFlow) {
		V475 := __e.Get(1)
		_ = V475
		tmp68233 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp68234 := Call(__e, PrimNS1Value(symns2_1value), symremove)

		tmp68235 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp68236 := Call(__e, tmp68235, symshen_4_dspecial_d)

		tmp68237 := Call(__e, tmp68234, V475, tmp68236)

		tmp68238 := Call(__e, tmp68233, symshen_4_dspecial_d, tmp68237)

		_ = tmp68238

		__e.Return(V475)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symunspecialise, tmp68232)
	return

}, 0)
