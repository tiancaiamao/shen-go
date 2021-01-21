package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp63050 := MakeNative(func(__e *ControlFlow) {
		V152 := __e.Get(1)
		_ = V152
		V153 := __e.Get(2)
		_ = V153
		tmp63051 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

		tmp63052 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp63053 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5define_6)

			__e.TailApply(tmp63053, X)
			return

		}, 1)

		tmp63054 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp63055 := Call(__e, tmp63054, V152, V153)

		tmp63056 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp63057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4shen_1syntax_1error)

			__e.TailApply(tmp63057, V152, X)
			return

		}, 1)

		__e.TailApply(tmp63051, tmp63052, tmp63055, tmp63056)
		return

	}, 2)

	tmp63058 := Call(__e, PrimNS1Value(symns2_1set), symshen_4shen_1_6kl, tmp63050)

	_ = tmp63058

	tmp63059 := MakeNative(func(__e *ControlFlow) {
		V160 := __e.Get(1)
		_ = V160
		V161 := __e.Get(2)
		_ = V161
		tmp63079 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63080 := Call(__e, tmp63079, V161)

		if True == tmp63080 {
			tmp63066 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp63067 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp63068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp63069 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp63070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp63071 := Call(__e, PrimNS1Value(symns2_1value), symshen_4next_150)

			tmp63072 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp63073 := Call(__e, tmp63072, V161)

			tmp63074 := Call(__e, tmp63071, MakeNumber(50), tmp63073)

			tmp63075 := Call(__e, tmp63070, tmp63074, MakeString("\n"), symshen_4a)

			tmp63076 := Call(__e, tmp63069, MakeString(" here:\n\n "), tmp63075)

			tmp63077 := Call(__e, tmp63068, V160, tmp63076, symshen_4a)

			tmp63078 := Call(__e, tmp63067, MakeString("syntax error in "), tmp63077)

			__e.TailApply(tmp63066, tmp63078)
			return

		} else {
			tmp63061 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp63062 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp63063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp63064 := Call(__e, tmp63063, V160, MakeString("\n"), symshen_4a)

			tmp63065 := Call(__e, tmp63062, MakeString("syntax error in "), tmp63064)

			__e.TailApply(tmp63061, tmp63065)
			return

		}

	}, 2)

	tmp63081 := Call(__e, PrimNS1Value(symns2_1set), symshen_4shen_1syntax_1error, tmp63059)

	_ = tmp63081

	tmp63082 := MakeNative(func(__e *ControlFlow) {
		V163 := __e.Get(1)
		_ = V163
		tmp63083 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp63116 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63117 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63118 := Call(__e, tmp63117)

			tmp63119 := Call(__e, tmp63116, YaccParse, tmp63118)

			if True == tmp63119 {
				tmp63085 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5name_6 := __e.Get(1)
					_ = Parse__shen_4_5name_6
					tmp63108 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63109 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63110 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63111 := Call(__e, tmp63110)

					tmp63112 := Call(__e, tmp63109, tmp63111, Parse__shen_4_5name_6)

					tmp63113 := Call(__e, tmp63108, tmp63112)

					if True == tmp63113 {
						tmp63088 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rules_6 := __e.Get(1)
							_ = Parse__shen_4_5rules_6
							tmp63100 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp63101 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63102 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp63103 := Call(__e, tmp63102)

							tmp63104 := Call(__e, tmp63101, tmp63103, Parse__shen_4_5rules_6)

							tmp63105 := Call(__e, tmp63100, tmp63104)

							if True == tmp63105 {
								tmp63091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp63092 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63093 := Call(__e, tmp63092, Parse__shen_4_5rules_6)

								tmp63094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compile__to__machine__code)

								tmp63095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp63096 := Call(__e, tmp63095, Parse__shen_4_5name_6)

								tmp63097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp63098 := Call(__e, tmp63097, Parse__shen_4_5rules_6)

								tmp63099 := Call(__e, tmp63094, tmp63096, tmp63098)

								__e.TailApply(tmp63091, tmp63093, tmp63099)
								return

							} else {
								tmp63090 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp63090)
								return

							}

						}, 1)

						tmp63106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rules_6)

						tmp63107 := Call(__e, tmp63106, Parse__shen_4_5name_6)

						__e.TailApply(tmp63088, tmp63107)
						return

					} else {
						tmp63087 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63087)
						return

					}

				}, 1)

				tmp63114 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5name_6)

				tmp63115 := Call(__e, tmp63114, V163)

				__e.TailApply(tmp63085, tmp63115)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp63120 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5name_6 := __e.Get(1)
			_ = Parse__shen_4_5name_6
			tmp63154 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp63155 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63156 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63157 := Call(__e, tmp63156)

			tmp63158 := Call(__e, tmp63155, tmp63157, Parse__shen_4_5name_6)

			tmp63159 := Call(__e, tmp63154, tmp63158)

			if True == tmp63159 {
				tmp63123 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_6
					tmp63146 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63147 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63148 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63149 := Call(__e, tmp63148)

					tmp63150 := Call(__e, tmp63147, tmp63149, Parse__shen_4_5signature_6)

					tmp63151 := Call(__e, tmp63146, tmp63150)

					if True == tmp63151 {
						tmp63126 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rules_6 := __e.Get(1)
							_ = Parse__shen_4_5rules_6
							tmp63138 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp63139 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63140 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp63141 := Call(__e, tmp63140)

							tmp63142 := Call(__e, tmp63139, tmp63141, Parse__shen_4_5rules_6)

							tmp63143 := Call(__e, tmp63138, tmp63142)

							if True == tmp63143 {
								tmp63129 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp63130 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63131 := Call(__e, tmp63130, Parse__shen_4_5rules_6)

								tmp63132 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compile__to__machine__code)

								tmp63133 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp63134 := Call(__e, tmp63133, Parse__shen_4_5name_6)

								tmp63135 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp63136 := Call(__e, tmp63135, Parse__shen_4_5rules_6)

								tmp63137 := Call(__e, tmp63132, tmp63134, tmp63136)

								__e.TailApply(tmp63129, tmp63131, tmp63137)
								return

							} else {
								tmp63128 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp63128)
								return

							}

						}, 1)

						tmp63144 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rules_6)

						tmp63145 := Call(__e, tmp63144, Parse__shen_4_5signature_6)

						__e.TailApply(tmp63126, tmp63145)
						return

					} else {
						tmp63125 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63125)
						return

					}

				}, 1)

				tmp63152 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_6)

				tmp63153 := Call(__e, tmp63152, Parse__shen_4_5name_6)

				__e.TailApply(tmp63123, tmp63153)
				return

			} else {
				tmp63122 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp63122)
				return

			}

		}, 1)

		tmp63160 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5name_6)

		tmp63161 := Call(__e, tmp63160, V163)

		tmp63162 := Call(__e, tmp63120, tmp63161)

		__e.TailApply(tmp63083, tmp63162)
		return

	}, 1)

	tmp63163 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5define_6, tmp63082)

	_ = tmp63163

	tmp63164 := MakeNative(func(__e *ControlFlow) {
		V165 := __e.Get(1)
		_ = V165
		tmp63192 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63193 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp63194 := Call(__e, tmp63193, V165)

		tmp63195 := Call(__e, tmp63192, tmp63194)

		if True == tmp63195 {
			tmp63167 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp63168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp63169 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp63170 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp63171 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp63172 := Call(__e, tmp63171, V165)

				tmp63173 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp63174 := Call(__e, tmp63173, V165)

				tmp63175 := Call(__e, tmp63170, tmp63172, tmp63174)

				tmp63176 := Call(__e, tmp63169, tmp63175)

				tmp63188 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

				tmp63189 := Call(__e, tmp63188, Parse__X)

				var ifres63182 Obj

				if True == tmp63189 {
					tmp63184 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63185 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sysfunc_2)

					tmp63186 := Call(__e, tmp63185, Parse__X)

					tmp63187 := Call(__e, tmp63184, tmp63186)

					var ifres63183 Obj

					if True == tmp63187 {
						ifres63183 = True

					} else {
						ifres63183 = False

					}

					ifres63182 = ifres63183

				} else {
					ifres63182 = False

				}

				var ifres63177 Obj

				if True == ifres63182 {
					ifres63177 = Parse__X

				} else {
					tmp63178 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					tmp63179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp63180 := Call(__e, tmp63179, Parse__X, MakeString(" is not a legitimate function name.\n"), symshen_4a)

					tmp63181 := Call(__e, tmp63178, tmp63180)

					ifres63177 = tmp63181

				}

				__e.TailApply(tmp63168, tmp63176, ifres63177)
				return

			}, 1)

			tmp63190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp63191 := Call(__e, tmp63190, V165)

			__e.TailApply(tmp63167, tmp63191)
			return

		} else {
			tmp63166 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp63166)
			return

		}

	}, 1)

	tmp63196 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5name_6, tmp63164)

	_ = tmp63196

	tmp63197 := MakeNative(func(__e *ControlFlow) {
		V167 := __e.Get(1)
		_ = V167
		tmp63198 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp63199 := Call(__e, PrimNS1Value(symns2_1value), symget)

		tmp63200 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp63201 := Call(__e, tmp63200, MakeString("shen"))

		tmp63202 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp63203 := Call(__e, tmp63202, sym_dproperty_1vector_d)

		tmp63204 := Call(__e, tmp63199, tmp63201, symshen_4external_1symbols, tmp63203)

		__e.TailApply(tmp63198, V167, tmp63204)
		return

	}, 1)

	tmp63205 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sysfunc_2, tmp63197)

	_ = tmp63205

	tmp63206 := MakeNative(func(__e *ControlFlow) {
		V171 := __e.Get(1)
		_ = V171
		tmp63261 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63262 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp63263 := Call(__e, tmp63262, V171)

		tmp63264 := Call(__e, tmp63261, tmp63263)

		var ifres63255 Obj

		if True == tmp63264 {
			tmp63257 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63258 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp63259 := Call(__e, tmp63258, V171)

			tmp63260 := Call(__e, tmp63257, sym_i, tmp63259)

			var ifres63256 Obj

			if True == tmp63260 {
				ifres63256 = True

			} else {
				ifres63256 = False

			}

			ifres63255 = ifres63256

		} else {
			ifres63255 = False

		}

		if True == ifres63255 {
			tmp63209 := MakeNative(func(__e *ControlFlow) {
				NewStream168 := __e.Get(1)
				_ = NewStream168
				tmp63210 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_1help_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_1help_6
					tmp63241 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63242 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63243 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63244 := Call(__e, tmp63243)

					tmp63245 := Call(__e, tmp63242, tmp63244, Parse__shen_4_5signature_1help_6)

					tmp63246 := Call(__e, tmp63241, tmp63245)

					if True == tmp63246 {
						tmp63237 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp63238 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63239 := Call(__e, tmp63238, Parse__shen_4_5signature_1help_6)

						tmp63240 := Call(__e, tmp63237, tmp63239)

						var ifres63231 Obj

						if True == tmp63240 {
							tmp63233 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63234 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp63235 := Call(__e, tmp63234, Parse__shen_4_5signature_1help_6)

							tmp63236 := Call(__e, tmp63233, sym_j, tmp63235)

							var ifres63232 Obj

							if True == tmp63236 {
								ifres63232 = True

							} else {
								ifres63232 = False

							}

							ifres63231 = ifres63232

						} else {
							ifres63231 = False

						}

						if True == ifres63231 {
							tmp63215 := MakeNative(func(__e *ControlFlow) {
								NewStream169 := __e.Get(1)
								_ = NewStream169
								tmp63216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp63217 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63218 := Call(__e, tmp63217, NewStream169)

								tmp63219 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

								tmp63220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

								tmp63221 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp63222 := Call(__e, tmp63221, Parse__shen_4_5signature_1help_6)

								tmp63223 := Call(__e, tmp63220, tmp63222)

								tmp63224 := Call(__e, tmp63219, tmp63223)

								__e.TailApply(tmp63216, tmp63218, tmp63224)
								return

							}, 1)

							tmp63225 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp63226 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							tmp63227 := Call(__e, tmp63226, Parse__shen_4_5signature_1help_6)

							tmp63228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp63229 := Call(__e, tmp63228, Parse__shen_4_5signature_1help_6)

							tmp63230 := Call(__e, tmp63225, tmp63227, tmp63229)

							__e.TailApply(tmp63215, tmp63230)
							return

						} else {
							tmp63214 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp63214)
							return

						}

					} else {
						tmp63212 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63212)
						return

					}

				}, 1)

				tmp63247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_1help_6)

				tmp63248 := Call(__e, tmp63247, NewStream168)

				__e.TailApply(tmp63210, tmp63248)
				return

			}, 1)

			tmp63249 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp63250 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp63251 := Call(__e, tmp63250, V171)

			tmp63252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp63253 := Call(__e, tmp63252, V171)

			tmp63254 := Call(__e, tmp63249, tmp63251, tmp63253)

			__e.TailApply(tmp63209, tmp63254)
			return

		} else {
			tmp63208 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp63208)
			return

		}

	}, 1)

	tmp63265 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5signature_6, tmp63206)

	_ = tmp63265

	tmp63266 := MakeNative(func(__e *ControlFlow) {
		V173 := __e.Get(1)
		_ = V173
		tmp63267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4active_1cons)

		tmp63268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type_1h)

		tmp63269 := Call(__e, tmp63268, V173)

		__e.TailApply(tmp63267, tmp63269)
		return

	}, 1)

	tmp63270 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1type, tmp63266)

	_ = tmp63270

	tmp63271 := MakeNative(func(__e *ControlFlow) {
		V175 := __e.Get(1)
		_ = V175
		tmp63330 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63331 := Call(__e, tmp63330, V175)

		var ifres63298 Obj

		if True == tmp63331 {
			tmp63326 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp63327 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63328 := Call(__e, tmp63327, V175)

			tmp63329 := Call(__e, tmp63326, tmp63328)

			var ifres63300 Obj

			if True == tmp63329 {
				tmp63320 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp63321 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63322 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63323 := Call(__e, tmp63322, V175)

				tmp63324 := Call(__e, tmp63321, tmp63323)

				tmp63325 := Call(__e, tmp63320, tmp63324)

				var ifres63302 Obj

				if True == tmp63325 {
					tmp63312 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63313 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63314 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63315 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63316 := Call(__e, tmp63315, V175)

					tmp63317 := Call(__e, tmp63314, tmp63316)

					tmp63318 := Call(__e, tmp63313, tmp63317)

					tmp63319 := Call(__e, tmp63312, Nil, tmp63318)

					var ifres63304 Obj

					if True == tmp63319 {
						tmp63306 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp63307 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63308 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63309 := Call(__e, tmp63308, V175)

						tmp63310 := Call(__e, tmp63307, tmp63309)

						tmp63311 := Call(__e, tmp63306, tmp63310, symbar_b)

						var ifres63305 Obj

						if True == tmp63311 {
							ifres63305 = True

						} else {
							ifres63305 = False

						}

						ifres63304 = ifres63305

					} else {
						ifres63304 = False

					}

					var ifres63303 Obj

					if True == ifres63304 {
						ifres63303 = True

					} else {
						ifres63303 = False

					}

					ifres63302 = ifres63303

				} else {
					ifres63302 = False

				}

				var ifres63301 Obj

				if True == ifres63302 {
					ifres63301 = True

				} else {
					ifres63301 = False

				}

				ifres63300 = ifres63301

			} else {
				ifres63300 = False

			}

			var ifres63299 Obj

			if True == ifres63300 {
				ifres63299 = True

			} else {
				ifres63299 = False

			}

			ifres63298 = ifres63299

		} else {
			ifres63298 = False

		}

		if True == ifres63298 {
			tmp63285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp63286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4active_1cons)

			tmp63287 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp63288 := Call(__e, tmp63287, V175)

			tmp63289 := Call(__e, tmp63286, tmp63288)

			tmp63290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4active_1cons)

			tmp63291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp63292 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63293 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63294 := Call(__e, tmp63293, V175)

			tmp63295 := Call(__e, tmp63292, tmp63294)

			tmp63296 := Call(__e, tmp63291, tmp63295)

			tmp63297 := Call(__e, tmp63290, tmp63296)

			__e.TailApply(tmp63285, tmp63289, tmp63297)
			return

		} else {
			tmp63283 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp63284 := Call(__e, tmp63283, V175)

			if True == tmp63284 {
				tmp63274 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp63275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4active_1cons)

				tmp63276 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp63277 := Call(__e, tmp63276, V175)

				tmp63278 := Call(__e, tmp63275, tmp63277)

				tmp63279 := Call(__e, PrimNS1Value(symns2_1value), symshen_4active_1cons)

				tmp63280 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63281 := Call(__e, tmp63280, V175)

				tmp63282 := Call(__e, tmp63279, tmp63281)

				__e.TailApply(tmp63274, tmp63278, tmp63282)
				return

			} else {
				__e.Return(V175)
				return
			}

		}

	}, 1)

	tmp63332 := Call(__e, PrimNS1Value(symns2_1set), symshen_4active_1cons, tmp63271)

	_ = tmp63332

	tmp63333 := MakeNative(func(__e *ControlFlow) {
		V177 := __e.Get(1)
		_ = V177
		tmp63458 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63459 := Call(__e, tmp63458, V177)

		var ifres63414 Obj

		if True == tmp63459 {
			tmp63454 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp63455 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63456 := Call(__e, tmp63455, V177)

			tmp63457 := Call(__e, tmp63454, tmp63456)

			var ifres63416 Obj

			if True == tmp63457 {
				tmp63448 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp63449 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp63450 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63451 := Call(__e, tmp63450, V177)

				tmp63452 := Call(__e, tmp63449, tmp63451)

				tmp63453 := Call(__e, tmp63448, sym_1_1_6, tmp63452)

				var ifres63418 Obj

				if True == tmp63453 {
					tmp63442 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp63443 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63444 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63445 := Call(__e, tmp63444, V177)

					tmp63446 := Call(__e, tmp63443, tmp63445)

					tmp63447 := Call(__e, tmp63442, tmp63446)

					var ifres63420 Obj

					if True == tmp63447 {
						tmp63434 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp63435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63436 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63437 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63438 := Call(__e, tmp63437, V177)

						tmp63439 := Call(__e, tmp63436, tmp63438)

						tmp63440 := Call(__e, tmp63435, tmp63439)

						tmp63441 := Call(__e, tmp63434, tmp63440)

						var ifres63422 Obj

						if True == tmp63441 {
							tmp63424 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63425 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp63426 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63427 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63428 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63429 := Call(__e, tmp63428, V177)

							tmp63430 := Call(__e, tmp63427, tmp63429)

							tmp63431 := Call(__e, tmp63426, tmp63430)

							tmp63432 := Call(__e, tmp63425, tmp63431)

							tmp63433 := Call(__e, tmp63424, sym_1_1_6, tmp63432)

							var ifres63423 Obj

							if True == tmp63433 {
								ifres63423 = True

							} else {
								ifres63423 = False

							}

							ifres63422 = ifres63423

						} else {
							ifres63422 = False

						}

						var ifres63421 Obj

						if True == ifres63422 {
							ifres63421 = True

						} else {
							ifres63421 = False

						}

						ifres63420 = ifres63421

					} else {
						ifres63420 = False

					}

					var ifres63419 Obj

					if True == ifres63420 {
						ifres63419 = True

					} else {
						ifres63419 = False

					}

					ifres63418 = ifres63419

				} else {
					ifres63418 = False

				}

				var ifres63417 Obj

				if True == ifres63418 {
					ifres63417 = True

				} else {
					ifres63417 = False

				}

				ifres63416 = ifres63417

			} else {
				ifres63416 = False

			}

			var ifres63415 Obj

			if True == ifres63416 {
				ifres63415 = True

			} else {
				ifres63415 = False

			}

			ifres63414 = ifres63415

		} else {
			ifres63414 = False

		}

		if True == ifres63414 {
			tmp63401 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type_1h)

			tmp63402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp63403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp63404 := Call(__e, tmp63403, V177)

			tmp63405 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp63406 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp63407 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63408 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp63409 := Call(__e, tmp63408, V177)

			tmp63410 := Call(__e, tmp63407, tmp63409)

			tmp63411 := Call(__e, tmp63406, tmp63410, Nil)

			tmp63412 := Call(__e, tmp63405, sym_1_1_6, tmp63411)

			tmp63413 := Call(__e, tmp63402, tmp63404, tmp63412)

			__e.TailApply(tmp63401, tmp63413)
			return

		} else {
			tmp63399 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp63400 := Call(__e, tmp63399, V177)

			var ifres63355 Obj

			if True == tmp63400 {
				tmp63395 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp63396 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63397 := Call(__e, tmp63396, V177)

				tmp63398 := Call(__e, tmp63395, tmp63397)

				var ifres63357 Obj

				if True == tmp63398 {
					tmp63389 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63390 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp63391 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp63392 := Call(__e, tmp63391, V177)

					tmp63393 := Call(__e, tmp63390, tmp63392)

					tmp63394 := Call(__e, tmp63389, sym_d, tmp63393)

					var ifres63359 Obj

					if True == tmp63394 {
						tmp63383 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp63384 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63385 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp63386 := Call(__e, tmp63385, V177)

						tmp63387 := Call(__e, tmp63384, tmp63386)

						tmp63388 := Call(__e, tmp63383, tmp63387)

						var ifres63361 Obj

						if True == tmp63388 {
							tmp63375 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp63376 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63377 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63378 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp63379 := Call(__e, tmp63378, V177)

							tmp63380 := Call(__e, tmp63377, tmp63379)

							tmp63381 := Call(__e, tmp63376, tmp63380)

							tmp63382 := Call(__e, tmp63375, tmp63381)

							var ifres63363 Obj

							if True == tmp63382 {
								tmp63365 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp63366 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63367 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp63368 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp63369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp63370 := Call(__e, tmp63369, V177)

								tmp63371 := Call(__e, tmp63368, tmp63370)

								tmp63372 := Call(__e, tmp63367, tmp63371)

								tmp63373 := Call(__e, tmp63366, tmp63372)

								tmp63374 := Call(__e, tmp63365, sym_d, tmp63373)

								var ifres63364 Obj

								if True == tmp63374 {
									ifres63364 = True

								} else {
									ifres63364 = False

								}

								ifres63363 = ifres63364

							} else {
								ifres63363 = False

							}

							var ifres63362 Obj

							if True == ifres63363 {
								ifres63362 = True

							} else {
								ifres63362 = False

							}

							ifres63361 = ifres63362

						} else {
							ifres63361 = False

						}

						var ifres63360 Obj

						if True == ifres63361 {
							ifres63360 = True

						} else {
							ifres63360 = False

						}

						ifres63359 = ifres63360

					} else {
						ifres63359 = False

					}

					var ifres63358 Obj

					if True == ifres63359 {
						ifres63358 = True

					} else {
						ifres63358 = False

					}

					ifres63357 = ifres63358

				} else {
					ifres63357 = False

				}

				var ifres63356 Obj

				if True == ifres63357 {
					ifres63356 = True

				} else {
					ifres63356 = False

				}

				ifres63355 = ifres63356

			} else {
				ifres63355 = False

			}

			if True == ifres63355 {
				tmp63342 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type_1h)

				tmp63343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp63344 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp63345 := Call(__e, tmp63344, V177)

				tmp63346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp63347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp63348 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63349 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp63350 := Call(__e, tmp63349, V177)

				tmp63351 := Call(__e, tmp63348, tmp63350)

				tmp63352 := Call(__e, tmp63347, tmp63351, Nil)

				tmp63353 := Call(__e, tmp63346, sym_d, tmp63352)

				tmp63354 := Call(__e, tmp63343, tmp63345, tmp63353)

				__e.TailApply(tmp63342, tmp63354)
				return

			} else {
				tmp63340 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp63341 := Call(__e, tmp63340, V177)

				if True == tmp63341 {
					tmp63337 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					tmp63338 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						tmp63339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type_1h)

						__e.TailApply(tmp63339, Z)
						return

					}, 1)

					__e.TailApply(tmp63337, tmp63338, V177)
					return

				} else {
					__e.Return(V177)
					return
				}

			}

		}

	}, 1)

	tmp63460 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1type_1h, tmp63333)

	_ = tmp63460

	tmp63461 := MakeNative(func(__e *ControlFlow) {
		V179 := __e.Get(1)
		_ = V179
		tmp63462 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp63478 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63479 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63480 := Call(__e, tmp63479)

			tmp63481 := Call(__e, tmp63478, YaccParse, tmp63480)

			if True == tmp63481 {
				tmp63464 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp63470 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63471 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63472 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63473 := Call(__e, tmp63472)

					tmp63474 := Call(__e, tmp63471, tmp63473, Parse___5e_6)

					tmp63475 := Call(__e, tmp63470, tmp63474)

					if True == tmp63475 {
						tmp63467 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp63468 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63469 := Call(__e, tmp63468, Parse___5e_6)

						__e.TailApply(tmp63467, tmp63469, Nil)
						return

					} else {
						tmp63466 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63466)
						return

					}

				}, 1)

				tmp63476 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp63477 := Call(__e, tmp63476, V179)

				__e.TailApply(tmp63464, tmp63477)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp63523 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp63524 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp63525 := Call(__e, tmp63524, V179)

		tmp63526 := Call(__e, tmp63523, tmp63525)

		var ifres63482 Obj

		if True == tmp63526 {
			tmp63485 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp63486 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_1help_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_1help_6
					tmp63506 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63507 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63508 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63509 := Call(__e, tmp63508)

					tmp63510 := Call(__e, tmp63507, tmp63509, Parse__shen_4_5signature_1help_6)

					tmp63511 := Call(__e, tmp63506, tmp63510)

					if True == tmp63511 {
						tmp63498 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						tmp63499 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

						tmp63500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp63501 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp63502 := Call(__e, tmp63501, sym_j, Nil)

						tmp63503 := Call(__e, tmp63500, sym_i, tmp63502)

						tmp63504 := Call(__e, tmp63499, Parse__X, tmp63503)

						tmp63505 := Call(__e, tmp63498, tmp63504)

						if True == tmp63505 {
							tmp63491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp63492 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp63493 := Call(__e, tmp63492, Parse__shen_4_5signature_1help_6)

							tmp63494 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp63495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp63496 := Call(__e, tmp63495, Parse__shen_4_5signature_1help_6)

							tmp63497 := Call(__e, tmp63494, Parse__X, tmp63496)

							__e.TailApply(tmp63491, tmp63493, tmp63497)
							return

						} else {
							tmp63490 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp63490)
							return

						}

					} else {
						tmp63488 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63488)
						return

					}

				}, 1)

				tmp63512 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_1help_6)

				tmp63513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp63514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp63515 := Call(__e, tmp63514, V179)

				tmp63516 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp63517 := Call(__e, tmp63516, V179)

				tmp63518 := Call(__e, tmp63513, tmp63515, tmp63517)

				tmp63519 := Call(__e, tmp63512, tmp63518)

				__e.TailApply(tmp63486, tmp63519)
				return

			}, 1)

			tmp63520 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp63521 := Call(__e, tmp63520, V179)

			tmp63522 := Call(__e, tmp63485, tmp63521)

			ifres63482 = tmp63522

		} else {
			tmp63483 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63484 := Call(__e, tmp63483)

			ifres63482 = tmp63484

		}

		__e.TailApply(tmp63462, ifres63482)
		return

	}, 1)

	tmp63527 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5signature_1help_6, tmp63461)

	_ = tmp63527

	tmp63528 := MakeNative(func(__e *ControlFlow) {
		V181 := __e.Get(1)
		_ = V181
		tmp63529 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp63551 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63552 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63553 := Call(__e, tmp63552)

			tmp63554 := Call(__e, tmp63551, YaccParse, tmp63553)

			if True == tmp63554 {
				tmp63531 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rule_6 := __e.Get(1)
					_ = Parse__shen_4_5rule_6
					tmp63543 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63544 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63545 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63546 := Call(__e, tmp63545)

					tmp63547 := Call(__e, tmp63544, tmp63546, Parse__shen_4_5rule_6)

					tmp63548 := Call(__e, tmp63543, tmp63547)

					if True == tmp63548 {
						tmp63534 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp63535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63536 := Call(__e, tmp63535, Parse__shen_4_5rule_6)

						tmp63537 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp63538 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise)

						tmp63539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp63540 := Call(__e, tmp63539, Parse__shen_4_5rule_6)

						tmp63541 := Call(__e, tmp63538, tmp63540)

						tmp63542 := Call(__e, tmp63537, tmp63541, Nil)

						__e.TailApply(tmp63534, tmp63536, tmp63542)
						return

					} else {
						tmp63533 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63533)
						return

					}

				}, 1)

				tmp63549 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rule_6)

				tmp63550 := Call(__e, tmp63549, V181)

				__e.TailApply(tmp63531, tmp63550)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp63555 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5rule_6 := __e.Get(1)
			_ = Parse__shen_4_5rule_6
			tmp63580 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp63581 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63582 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63583 := Call(__e, tmp63582)

			tmp63584 := Call(__e, tmp63581, tmp63583, Parse__shen_4_5rule_6)

			tmp63585 := Call(__e, tmp63580, tmp63584)

			if True == tmp63585 {
				tmp63558 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rules_6 := __e.Get(1)
					_ = Parse__shen_4_5rules_6
					tmp63572 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63573 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63574 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63575 := Call(__e, tmp63574)

					tmp63576 := Call(__e, tmp63573, tmp63575, Parse__shen_4_5rules_6)

					tmp63577 := Call(__e, tmp63572, tmp63576)

					if True == tmp63577 {
						tmp63561 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp63562 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63563 := Call(__e, tmp63562, Parse__shen_4_5rules_6)

						tmp63564 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp63565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise)

						tmp63566 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp63567 := Call(__e, tmp63566, Parse__shen_4_5rule_6)

						tmp63568 := Call(__e, tmp63565, tmp63567)

						tmp63569 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp63570 := Call(__e, tmp63569, Parse__shen_4_5rules_6)

						tmp63571 := Call(__e, tmp63564, tmp63568, tmp63570)

						__e.TailApply(tmp63561, tmp63563, tmp63571)
						return

					} else {
						tmp63560 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63560)
						return

					}

				}, 1)

				tmp63578 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rules_6)

				tmp63579 := Call(__e, tmp63578, Parse__shen_4_5rule_6)

				__e.TailApply(tmp63558, tmp63579)
				return

			} else {
				tmp63557 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp63557)
				return

			}

		}, 1)

		tmp63586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rule_6)

		tmp63587 := Call(__e, tmp63586, V181)

		tmp63588 := Call(__e, tmp63555, tmp63587)

		__e.TailApply(tmp63529, tmp63588)
		return

	}, 1)

	tmp63589 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rules_6, tmp63528)

	_ = tmp63589

	tmp63590 := MakeNative(func(__e *ControlFlow) {
		V189 := __e.Get(1)
		_ = V189
		tmp63591 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp63809 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63810 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63811 := Call(__e, tmp63810)

			tmp63812 := Call(__e, tmp63809, YaccParse, tmp63811)

			if True == tmp63812 {
				tmp63593 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp63752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63753 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63754 := Call(__e, tmp63753)

					tmp63755 := Call(__e, tmp63752, YaccParse, tmp63754)

					if True == tmp63755 {
						tmp63595 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp63653 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63654 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp63655 := Call(__e, tmp63654)

							tmp63656 := Call(__e, tmp63653, YaccParse, tmp63655)

							if True == tmp63656 {
								tmp63597 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5patterns_6 := __e.Get(1)
									_ = Parse__shen_4_5patterns_6
									tmp63645 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp63646 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp63647 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp63648 := Call(__e, tmp63647)

									tmp63649 := Call(__e, tmp63646, tmp63648, Parse__shen_4_5patterns_6)

									tmp63650 := Call(__e, tmp63645, tmp63649)

									if True == tmp63650 {
										tmp63641 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp63642 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp63643 := Call(__e, tmp63642, Parse__shen_4_5patterns_6)

										tmp63644 := Call(__e, tmp63641, tmp63643)

										var ifres63635 Obj

										if True == tmp63644 {
											tmp63637 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp63638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

											tmp63639 := Call(__e, tmp63638, Parse__shen_4_5patterns_6)

											tmp63640 := Call(__e, tmp63637, sym_5_1, tmp63639)

											var ifres63636 Obj

											if True == tmp63640 {
												ifres63636 = True

											} else {
												ifres63636 = False

											}

											ifres63635 = ifres63636

										} else {
											ifres63635 = False

										}

										if True == ifres63635 {
											tmp63602 := MakeNative(func(__e *ControlFlow) {
												NewStream187 := __e.Get(1)
												_ = NewStream187
												tmp63603 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5action_6 := __e.Get(1)
													_ = Parse__shen_4_5action_6
													tmp63621 := Call(__e, PrimNS1Value(symns2_1value), symnot)

													tmp63622 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp63623 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp63624 := Call(__e, tmp63623)

													tmp63625 := Call(__e, tmp63622, tmp63624, Parse__shen_4_5action_6)

													tmp63626 := Call(__e, tmp63621, tmp63625)

													if True == tmp63626 {
														tmp63606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														tmp63607 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp63608 := Call(__e, tmp63607, Parse__shen_4_5action_6)

														tmp63609 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp63610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp63611 := Call(__e, tmp63610, Parse__shen_4_5patterns_6)

														tmp63612 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp63613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp63614 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp63615 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp63616 := Call(__e, tmp63615, Parse__shen_4_5action_6)

														tmp63617 := Call(__e, tmp63614, tmp63616, Nil)

														tmp63618 := Call(__e, tmp63613, symshen_4choicepoint_b, tmp63617)

														tmp63619 := Call(__e, tmp63612, tmp63618, Nil)

														tmp63620 := Call(__e, tmp63609, tmp63611, tmp63619)

														__e.TailApply(tmp63606, tmp63608, tmp63620)
														return

													} else {
														tmp63605 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp63605)
														return

													}

												}, 1)

												tmp63627 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5action_6)

												tmp63628 := Call(__e, tmp63627, NewStream187)

												__e.TailApply(tmp63603, tmp63628)
												return

											}, 1)

											tmp63629 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

											tmp63630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

											tmp63631 := Call(__e, tmp63630, Parse__shen_4_5patterns_6)

											tmp63632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

											tmp63633 := Call(__e, tmp63632, Parse__shen_4_5patterns_6)

											tmp63634 := Call(__e, tmp63629, tmp63631, tmp63633)

											__e.TailApply(tmp63602, tmp63634)
											return

										} else {
											tmp63601 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											__e.TailApply(tmp63601)
											return

										}

									} else {
										tmp63599 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp63599)
										return

									}

								}, 1)

								tmp63651 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5patterns_6)

								tmp63652 := Call(__e, tmp63651, V189)

								__e.TailApply(tmp63597, tmp63652)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp63657 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5patterns_6 := __e.Get(1)
							_ = Parse__shen_4_5patterns_6
							tmp63743 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp63744 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63745 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp63746 := Call(__e, tmp63745)

							tmp63747 := Call(__e, tmp63744, tmp63746, Parse__shen_4_5patterns_6)

							tmp63748 := Call(__e, tmp63743, tmp63747)

							if True == tmp63748 {
								tmp63739 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp63740 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63741 := Call(__e, tmp63740, Parse__shen_4_5patterns_6)

								tmp63742 := Call(__e, tmp63739, tmp63741)

								var ifres63733 Obj

								if True == tmp63742 {
									tmp63735 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp63736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

									tmp63737 := Call(__e, tmp63736, Parse__shen_4_5patterns_6)

									tmp63738 := Call(__e, tmp63735, sym_5_1, tmp63737)

									var ifres63734 Obj

									if True == tmp63738 {
										ifres63734 = True

									} else {
										ifres63734 = False

									}

									ifres63733 = ifres63734

								} else {
									ifres63733 = False

								}

								if True == ifres63733 {
									tmp63662 := MakeNative(func(__e *ControlFlow) {
										NewStream185 := __e.Get(1)
										_ = NewStream185
										tmp63663 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5action_6 := __e.Get(1)
											_ = Parse__shen_4_5action_6
											tmp63719 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp63720 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp63721 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp63722 := Call(__e, tmp63721)

											tmp63723 := Call(__e, tmp63720, tmp63722, Parse__shen_4_5action_6)

											tmp63724 := Call(__e, tmp63719, tmp63723)

											if True == tmp63724 {
												tmp63715 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp63716 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp63717 := Call(__e, tmp63716, Parse__shen_4_5action_6)

												tmp63718 := Call(__e, tmp63715, tmp63717)

												var ifres63709 Obj

												if True == tmp63718 {
													tmp63711 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp63712 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

													tmp63713 := Call(__e, tmp63712, Parse__shen_4_5action_6)

													tmp63714 := Call(__e, tmp63711, symwhere, tmp63713)

													var ifres63710 Obj

													if True == tmp63714 {
														ifres63710 = True

													} else {
														ifres63710 = False

													}

													ifres63709 = ifres63710

												} else {
													ifres63709 = False

												}

												if True == ifres63709 {
													tmp63668 := MakeNative(func(__e *ControlFlow) {
														NewStream186 := __e.Get(1)
														_ = NewStream186
														tmp63669 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5guard_6 := __e.Get(1)
															_ = Parse__shen_4_5guard_6
															tmp63695 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp63696 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp63697 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp63698 := Call(__e, tmp63697)

															tmp63699 := Call(__e, tmp63696, tmp63698, Parse__shen_4_5guard_6)

															tmp63700 := Call(__e, tmp63695, tmp63699)

															if True == tmp63700 {
																tmp63672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																tmp63673 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp63674 := Call(__e, tmp63673, Parse__shen_4_5guard_6)

																tmp63675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63676 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp63677 := Call(__e, tmp63676, Parse__shen_4_5patterns_6)

																tmp63678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63679 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63680 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp63682 := Call(__e, tmp63681, Parse__shen_4_5guard_6)

																tmp63683 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63684 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63685 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp63686 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp63687 := Call(__e, tmp63686, Parse__shen_4_5action_6)

																tmp63688 := Call(__e, tmp63685, tmp63687, Nil)

																tmp63689 := Call(__e, tmp63684, symshen_4choicepoint_b, tmp63688)

																tmp63690 := Call(__e, tmp63683, tmp63689, Nil)

																tmp63691 := Call(__e, tmp63680, tmp63682, tmp63690)

																tmp63692 := Call(__e, tmp63679, symwhere, tmp63691)

																tmp63693 := Call(__e, tmp63678, tmp63692, Nil)

																tmp63694 := Call(__e, tmp63675, tmp63677, tmp63693)

																__e.TailApply(tmp63672, tmp63674, tmp63694)
																return

															} else {
																tmp63671 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp63671)
																return

															}

														}, 1)

														tmp63701 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5guard_6)

														tmp63702 := Call(__e, tmp63701, NewStream186)

														__e.TailApply(tmp63669, tmp63702)
														return

													}, 1)

													tmp63703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

													tmp63704 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

													tmp63705 := Call(__e, tmp63704, Parse__shen_4_5action_6)

													tmp63706 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

													tmp63707 := Call(__e, tmp63706, Parse__shen_4_5action_6)

													tmp63708 := Call(__e, tmp63703, tmp63705, tmp63707)

													__e.TailApply(tmp63668, tmp63708)
													return

												} else {
													tmp63667 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													__e.TailApply(tmp63667)
													return

												}

											} else {
												tmp63665 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp63665)
												return

											}

										}, 1)

										tmp63725 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5action_6)

										tmp63726 := Call(__e, tmp63725, NewStream185)

										__e.TailApply(tmp63663, tmp63726)
										return

									}, 1)

									tmp63727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

									tmp63728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

									tmp63729 := Call(__e, tmp63728, Parse__shen_4_5patterns_6)

									tmp63730 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp63731 := Call(__e, tmp63730, Parse__shen_4_5patterns_6)

									tmp63732 := Call(__e, tmp63727, tmp63729, tmp63731)

									__e.TailApply(tmp63662, tmp63732)
									return

								} else {
									tmp63661 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									__e.TailApply(tmp63661)
									return

								}

							} else {
								tmp63659 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp63659)
								return

							}

						}, 1)

						tmp63749 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5patterns_6)

						tmp63750 := Call(__e, tmp63749, V189)

						tmp63751 := Call(__e, tmp63657, tmp63750)

						__e.TailApply(tmp63595, tmp63751)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp63756 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5patterns_6 := __e.Get(1)
					_ = Parse__shen_4_5patterns_6
					tmp63800 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63801 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63802 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63803 := Call(__e, tmp63802)

					tmp63804 := Call(__e, tmp63801, tmp63803, Parse__shen_4_5patterns_6)

					tmp63805 := Call(__e, tmp63800, tmp63804)

					if True == tmp63805 {
						tmp63796 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp63797 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63798 := Call(__e, tmp63797, Parse__shen_4_5patterns_6)

						tmp63799 := Call(__e, tmp63796, tmp63798)

						var ifres63790 Obj

						if True == tmp63799 {
							tmp63792 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63793 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp63794 := Call(__e, tmp63793, Parse__shen_4_5patterns_6)

							tmp63795 := Call(__e, tmp63792, sym_1_6, tmp63794)

							var ifres63791 Obj

							if True == tmp63795 {
								ifres63791 = True

							} else {
								ifres63791 = False

							}

							ifres63790 = ifres63791

						} else {
							ifres63790 = False

						}

						if True == ifres63790 {
							tmp63761 := MakeNative(func(__e *ControlFlow) {
								NewStream184 := __e.Get(1)
								_ = NewStream184
								tmp63762 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5action_6 := __e.Get(1)
									_ = Parse__shen_4_5action_6
									tmp63776 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp63777 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp63778 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp63779 := Call(__e, tmp63778)

									tmp63780 := Call(__e, tmp63777, tmp63779, Parse__shen_4_5action_6)

									tmp63781 := Call(__e, tmp63776, tmp63780)

									if True == tmp63781 {
										tmp63765 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp63766 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp63767 := Call(__e, tmp63766, Parse__shen_4_5action_6)

										tmp63768 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp63769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp63770 := Call(__e, tmp63769, Parse__shen_4_5patterns_6)

										tmp63771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp63772 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp63773 := Call(__e, tmp63772, Parse__shen_4_5action_6)

										tmp63774 := Call(__e, tmp63771, tmp63773, Nil)

										tmp63775 := Call(__e, tmp63768, tmp63770, tmp63774)

										__e.TailApply(tmp63765, tmp63767, tmp63775)
										return

									} else {
										tmp63764 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp63764)
										return

									}

								}, 1)

								tmp63782 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5action_6)

								tmp63783 := Call(__e, tmp63782, NewStream184)

								__e.TailApply(tmp63762, tmp63783)
								return

							}, 1)

							tmp63784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp63785 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							tmp63786 := Call(__e, tmp63785, Parse__shen_4_5patterns_6)

							tmp63787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp63788 := Call(__e, tmp63787, Parse__shen_4_5patterns_6)

							tmp63789 := Call(__e, tmp63784, tmp63786, tmp63788)

							__e.TailApply(tmp63761, tmp63789)
							return

						} else {
							tmp63760 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp63760)
							return

						}

					} else {
						tmp63758 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63758)
						return

					}

				}, 1)

				tmp63806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5patterns_6)

				tmp63807 := Call(__e, tmp63806, V189)

				tmp63808 := Call(__e, tmp63756, tmp63807)

				__e.TailApply(tmp63593, tmp63808)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp63813 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5patterns_6 := __e.Get(1)
			_ = Parse__shen_4_5patterns_6
			tmp63895 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp63896 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63897 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63898 := Call(__e, tmp63897)

			tmp63899 := Call(__e, tmp63896, tmp63898, Parse__shen_4_5patterns_6)

			tmp63900 := Call(__e, tmp63895, tmp63899)

			if True == tmp63900 {
				tmp63891 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp63892 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp63893 := Call(__e, tmp63892, Parse__shen_4_5patterns_6)

				tmp63894 := Call(__e, tmp63891, tmp63893)

				var ifres63885 Obj

				if True == tmp63894 {
					tmp63887 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp63889 := Call(__e, tmp63888, Parse__shen_4_5patterns_6)

					tmp63890 := Call(__e, tmp63887, sym_1_6, tmp63889)

					var ifres63886 Obj

					if True == tmp63890 {
						ifres63886 = True

					} else {
						ifres63886 = False

					}

					ifres63885 = ifres63886

				} else {
					ifres63885 = False

				}

				if True == ifres63885 {
					tmp63818 := MakeNative(func(__e *ControlFlow) {
						NewStream182 := __e.Get(1)
						_ = NewStream182
						tmp63819 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5action_6 := __e.Get(1)
							_ = Parse__shen_4_5action_6
							tmp63871 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp63872 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp63873 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp63874 := Call(__e, tmp63873)

							tmp63875 := Call(__e, tmp63872, tmp63874, Parse__shen_4_5action_6)

							tmp63876 := Call(__e, tmp63871, tmp63875)

							if True == tmp63876 {
								tmp63867 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp63868 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp63869 := Call(__e, tmp63868, Parse__shen_4_5action_6)

								tmp63870 := Call(__e, tmp63867, tmp63869)

								var ifres63861 Obj

								if True == tmp63870 {
									tmp63863 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp63864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

									tmp63865 := Call(__e, tmp63864, Parse__shen_4_5action_6)

									tmp63866 := Call(__e, tmp63863, symwhere, tmp63865)

									var ifres63862 Obj

									if True == tmp63866 {
										ifres63862 = True

									} else {
										ifres63862 = False

									}

									ifres63861 = ifres63862

								} else {
									ifres63861 = False

								}

								if True == ifres63861 {
									tmp63824 := MakeNative(func(__e *ControlFlow) {
										NewStream183 := __e.Get(1)
										_ = NewStream183
										tmp63825 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5guard_6 := __e.Get(1)
											_ = Parse__shen_4_5guard_6
											tmp63847 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp63848 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp63849 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp63850 := Call(__e, tmp63849)

											tmp63851 := Call(__e, tmp63848, tmp63850, Parse__shen_4_5guard_6)

											tmp63852 := Call(__e, tmp63847, tmp63851)

											if True == tmp63852 {
												tmp63828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp63829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp63830 := Call(__e, tmp63829, Parse__shen_4_5guard_6)

												tmp63831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp63832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp63833 := Call(__e, tmp63832, Parse__shen_4_5patterns_6)

												tmp63834 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp63835 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp63836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp63837 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp63838 := Call(__e, tmp63837, Parse__shen_4_5guard_6)

												tmp63839 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp63840 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp63841 := Call(__e, tmp63840, Parse__shen_4_5action_6)

												tmp63842 := Call(__e, tmp63839, tmp63841, Nil)

												tmp63843 := Call(__e, tmp63836, tmp63838, tmp63842)

												tmp63844 := Call(__e, tmp63835, symwhere, tmp63843)

												tmp63845 := Call(__e, tmp63834, tmp63844, Nil)

												tmp63846 := Call(__e, tmp63831, tmp63833, tmp63845)

												__e.TailApply(tmp63828, tmp63830, tmp63846)
												return

											} else {
												tmp63827 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp63827)
												return

											}

										}, 1)

										tmp63853 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5guard_6)

										tmp63854 := Call(__e, tmp63853, NewStream183)

										__e.TailApply(tmp63825, tmp63854)
										return

									}, 1)

									tmp63855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

									tmp63856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

									tmp63857 := Call(__e, tmp63856, Parse__shen_4_5action_6)

									tmp63858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp63859 := Call(__e, tmp63858, Parse__shen_4_5action_6)

									tmp63860 := Call(__e, tmp63855, tmp63857, tmp63859)

									__e.TailApply(tmp63824, tmp63860)
									return

								} else {
									tmp63823 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									__e.TailApply(tmp63823)
									return

								}

							} else {
								tmp63821 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp63821)
								return

							}

						}, 1)

						tmp63877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5action_6)

						tmp63878 := Call(__e, tmp63877, NewStream182)

						__e.TailApply(tmp63819, tmp63878)
						return

					}, 1)

					tmp63879 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp63880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp63881 := Call(__e, tmp63880, Parse__shen_4_5patterns_6)

					tmp63882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp63883 := Call(__e, tmp63882, Parse__shen_4_5patterns_6)

					tmp63884 := Call(__e, tmp63879, tmp63881, tmp63883)

					__e.TailApply(tmp63818, tmp63884)
					return

				} else {
					tmp63817 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp63817)
					return

				}

			} else {
				tmp63815 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp63815)
				return

			}

		}, 1)

		tmp63901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5patterns_6)

		tmp63902 := Call(__e, tmp63901, V189)

		tmp63903 := Call(__e, tmp63813, tmp63902)

		__e.TailApply(tmp63591, tmp63903)
		return

	}, 1)

	tmp63904 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rule_6, tmp63590)

	_ = tmp63904

	tmp63905 := MakeNative(func(__e *ControlFlow) {
		V192 := __e.Get(1)
		_ = V192
		V193 := __e.Get(2)
		_ = V193
		tmp63908 := Call(__e, V192, V193)

		if True == tmp63908 {
			tmp63907 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp63907)
			return

		} else {
			__e.Return(V193)
			return
		}

	}, 2)

	tmp63909 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fail__if, tmp63905)

	_ = tmp63909

	tmp63910 := MakeNative(func(__e *ControlFlow) {
		V199 := __e.Get(1)
		_ = V199
		tmp63912 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp63913 := Call(__e, PrimNS1Value(symns2_1value), symfail)

		tmp63914 := Call(__e, tmp63913)

		tmp63915 := Call(__e, tmp63912, V199, tmp63914)

		if True == tmp63915 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp63916 := Call(__e, PrimNS1Value(symns2_1set), symshen_4succeeds_2, tmp63910)

	_ = tmp63916

	tmp63917 := MakeNative(func(__e *ControlFlow) {
		V202 := __e.Get(1)
		_ = V202
		V203 := __e.Get(2)
		_ = V203
		tmp63918 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp63919 := Call(__e, tmp63918, symshen_4_dcustom_1pattern_1compiler_d)

		__e.TailApply(tmp63919, V202, V203)
		return

	}, 2)

	tmp63920 := Call(__e, PrimNS1Value(symns2_1set), symshen_4custom_1pattern_1compiler, tmp63917)

	_ = tmp63920

	tmp63921 := MakeNative(func(__e *ControlFlow) {
		V205 := __e.Get(1)
		_ = V205
		tmp63922 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp63923 := Call(__e, tmp63922, symshen_4_dcustom_1pattern_1reducer_d)

		__e.TailApply(tmp63923, V205)
		return

	}, 1)

	tmp63924 := Call(__e, PrimNS1Value(symns2_1set), symshen_4custom_1pattern_1reducer, tmp63921)

	_ = tmp63924

	tmp63925 := MakeNative(func(__e *ControlFlow) {
		V207 := __e.Get(1)
		_ = V207
		tmp63926 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp63942 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63943 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63944 := Call(__e, tmp63943)

			tmp63945 := Call(__e, tmp63942, YaccParse, tmp63944)

			if True == tmp63945 {
				tmp63928 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp63934 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63935 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63936 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63937 := Call(__e, tmp63936)

					tmp63938 := Call(__e, tmp63935, tmp63937, Parse___5e_6)

					tmp63939 := Call(__e, tmp63934, tmp63938)

					if True == tmp63939 {
						tmp63931 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp63932 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63933 := Call(__e, tmp63932, Parse___5e_6)

						__e.TailApply(tmp63931, tmp63933, Nil)
						return

					} else {
						tmp63930 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63930)
						return

					}

				}, 1)

				tmp63940 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp63941 := Call(__e, tmp63940, V207)

				__e.TailApply(tmp63928, tmp63941)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp63946 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp63969 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp63970 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp63971 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp63972 := Call(__e, tmp63971)

			tmp63973 := Call(__e, tmp63970, tmp63972, Parse__shen_4_5pattern_6)

			tmp63974 := Call(__e, tmp63969, tmp63973)

			if True == tmp63974 {
				tmp63949 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5patterns_6 := __e.Get(1)
					_ = Parse__shen_4_5patterns_6
					tmp63961 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp63962 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp63963 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp63964 := Call(__e, tmp63963)

					tmp63965 := Call(__e, tmp63962, tmp63964, Parse__shen_4_5patterns_6)

					tmp63966 := Call(__e, tmp63961, tmp63965)

					if True == tmp63966 {
						tmp63952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp63953 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp63954 := Call(__e, tmp63953, Parse__shen_4_5patterns_6)

						tmp63955 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp63956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp63957 := Call(__e, tmp63956, Parse__shen_4_5pattern_6)

						tmp63958 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp63959 := Call(__e, tmp63958, Parse__shen_4_5patterns_6)

						tmp63960 := Call(__e, tmp63955, tmp63957, tmp63959)

						__e.TailApply(tmp63952, tmp63954, tmp63960)
						return

					} else {
						tmp63951 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp63951)
						return

					}

				}, 1)

				tmp63967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5patterns_6)

				tmp63968 := Call(__e, tmp63967, Parse__shen_4_5pattern_6)

				__e.TailApply(tmp63949, tmp63968)
				return

			} else {
				tmp63948 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp63948)
				return

			}

		}, 1)

		tmp63975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern_6)

		tmp63976 := Call(__e, tmp63975, V207)

		tmp63977 := Call(__e, tmp63946, tmp63976)

		__e.TailApply(tmp63926, tmp63977)
		return

	}, 1)

	tmp63978 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5patterns_6, tmp63925)

	_ = tmp63978

	tmp63979 := MakeNative(func(__e *ControlFlow) {
		V220 := __e.Get(1)
		_ = V220
		tmp63980 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp64443 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp64444 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64445 := Call(__e, tmp64444)

			tmp64446 := Call(__e, tmp64443, YaccParse, tmp64445)

			if True == tmp64446 {
				tmp63982 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp64340 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp64341 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp64342 := Call(__e, tmp64341)

					tmp64343 := Call(__e, tmp64340, YaccParse, tmp64342)

					if True == tmp64343 {
						tmp63984 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp64237 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp64238 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp64239 := Call(__e, tmp64238)

							tmp64240 := Call(__e, tmp64237, YaccParse, tmp64239)

							if True == tmp64240 {
								tmp63986 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp64134 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp64135 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp64136 := Call(__e, tmp64135)

									tmp64137 := Call(__e, tmp64134, YaccParse, tmp64136)

									if True == tmp64137 {
										tmp63988 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp64040 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp64041 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp64042 := Call(__e, tmp64041)

											tmp64043 := Call(__e, tmp64040, YaccParse, tmp64042)

											if True == tmp64043 {
												tmp63990 := MakeNative(func(__e *ControlFlow) {
													YaccParse := __e.Get(1)
													_ = YaccParse
													tmp64008 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp64009 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp64010 := Call(__e, tmp64009)

													tmp64011 := Call(__e, tmp64008, YaccParse, tmp64010)

													if True == tmp64011 {
														tmp63992 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5simple__pattern_6 := __e.Get(1)
															_ = Parse__shen_4_5simple__pattern_6
															tmp64000 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp64001 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp64002 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															tmp64003 := Call(__e, tmp64002)

															tmp64004 := Call(__e, tmp64001, tmp64003, Parse__shen_4_5simple__pattern_6)

															tmp64005 := Call(__e, tmp64000, tmp64004)

															if True == tmp64005 {
																tmp63995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

																tmp63996 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp63997 := Call(__e, tmp63996, Parse__shen_4_5simple__pattern_6)

																tmp63998 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

																tmp63999 := Call(__e, tmp63998, Parse__shen_4_5simple__pattern_6)

																__e.TailApply(tmp63995, tmp63997, tmp63999)
																return

															} else {
																tmp63994 := Call(__e, PrimNS1Value(symns2_1value), symfail)

																__e.TailApply(tmp63994)
																return

															}

														}, 1)

														tmp64006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5simple__pattern_6)

														tmp64007 := Call(__e, tmp64006, V220)

														__e.TailApply(tmp63992, tmp64007)
														return

													} else {
														__e.Return(YaccParse)
														return
													}

												}, 1)

												tmp64036 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp64037 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp64038 := Call(__e, tmp64037, V220)

												tmp64039 := Call(__e, tmp64036, tmp64038)

												var ifres64012 Obj

												if True == tmp64039 {
													tmp64015 := MakeNative(func(__e *ControlFlow) {
														Parse__X := __e.Get(1)
														_ = Parse__X
														tmp64031 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp64032 := Call(__e, tmp64031, Parse__X)

														if True == tmp64032 {
															tmp64018 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64019 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp64020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

															tmp64022 := Call(__e, tmp64021, V220)

															tmp64023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

															tmp64024 := Call(__e, tmp64023, V220)

															tmp64025 := Call(__e, tmp64020, tmp64022, tmp64024)

															tmp64026 := Call(__e, tmp64019, tmp64025)

															tmp64027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4custom_1pattern_1compiler)

															tmp64028 := MakeNative(func(__e *ControlFlow) {
																tmp64029 := Call(__e, PrimNS1Value(symns2_1value), symshen_4constructor_1error)

																__e.TailApply(tmp64029, Parse__X)
																return

															}, 0)

															tmp64030 := Call(__e, tmp64027, Parse__X, tmp64028)

															__e.TailApply(tmp64018, tmp64026, tmp64030)
															return

														} else {
															tmp64017 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															__e.TailApply(tmp64017)
															return

														}

													}, 1)

													tmp64033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

													tmp64034 := Call(__e, tmp64033, V220)

													tmp64035 := Call(__e, tmp64015, tmp64034)

													ifres64012 = tmp64035

												} else {
													tmp64013 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													tmp64014 := Call(__e, tmp64013)

													ifres64012 = tmp64014

												}

												__e.TailApply(tmp63990, ifres64012)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp64130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp64131 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp64132 := Call(__e, tmp64131, V220)

										tmp64133 := Call(__e, tmp64130, tmp64132)

										var ifres64124 Obj

										if True == tmp64133 {
											tmp64126 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp64127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

											tmp64128 := Call(__e, tmp64127, V220)

											tmp64129 := Call(__e, tmp64126, tmp64128)

											var ifres64125 Obj

											if True == tmp64129 {
												ifres64125 = True

											} else {
												ifres64125 = False

											}

											ifres64124 = ifres64125

										} else {
											ifres64124 = False

										}

										var ifres64044 Obj

										if True == ifres64124 {
											tmp64114 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp64115 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp64116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

											tmp64117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

											tmp64118 := Call(__e, tmp64117, V220)

											tmp64119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

											tmp64120 := Call(__e, tmp64119, V220)

											tmp64121 := Call(__e, tmp64116, tmp64118, tmp64120)

											tmp64122 := Call(__e, tmp64115, tmp64121)

											tmp64123 := Call(__e, tmp64114, tmp64122)

											var ifres64102 Obj

											if True == tmp64123 {
												tmp64104 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp64105 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

												tmp64106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp64107 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

												tmp64108 := Call(__e, tmp64107, V220)

												tmp64109 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp64110 := Call(__e, tmp64109, V220)

												tmp64111 := Call(__e, tmp64106, tmp64108, tmp64110)

												tmp64112 := Call(__e, tmp64105, tmp64111)

												tmp64113 := Call(__e, tmp64104, symvector, tmp64112)

												var ifres64103 Obj

												if True == tmp64113 {
													ifres64103 = True

												} else {
													ifres64103 = False

												}

												ifres64102 = ifres64103

											} else {
												ifres64102 = False

											}

											var ifres64047 Obj

											if True == ifres64102 {
												tmp64050 := MakeNative(func(__e *ControlFlow) {
													NewStream217 := __e.Get(1)
													_ = NewStream217
													tmp64079 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp64080 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp64081 := Call(__e, tmp64080, NewStream217)

													tmp64082 := Call(__e, tmp64079, tmp64081)

													var ifres64073 Obj

													if True == tmp64082 {
														tmp64075 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp64076 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

														tmp64077 := Call(__e, tmp64076, NewStream217)

														tmp64078 := Call(__e, tmp64075, MakeNumber(0), tmp64077)

														var ifres64074 Obj

														if True == tmp64078 {
															ifres64074 = True

														} else {
															ifres64074 = False

														}

														ifres64073 = ifres64074

													} else {
														ifres64073 = False

													}

													if True == ifres64073 {
														tmp64053 := MakeNative(func(__e *ControlFlow) {
															NewStream218 := __e.Get(1)
															_ = NewStream218
															tmp64054 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64055 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp64056 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

															tmp64058 := Call(__e, tmp64057, V220)

															tmp64059 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

															tmp64060 := Call(__e, tmp64059, V220)

															tmp64061 := Call(__e, tmp64056, tmp64058, tmp64060)

															tmp64062 := Call(__e, tmp64055, tmp64061)

															tmp64063 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp64064 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp64065 := Call(__e, tmp64064, MakeNumber(0), Nil)

															tmp64066 := Call(__e, tmp64063, symvector, tmp64065)

															__e.TailApply(tmp64054, tmp64062, tmp64066)
															return

														}, 1)

														tmp64067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

														tmp64068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

														tmp64069 := Call(__e, tmp64068, NewStream217)

														tmp64070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

														tmp64071 := Call(__e, tmp64070, NewStream217)

														tmp64072 := Call(__e, tmp64067, tmp64069, tmp64071)

														__e.TailApply(tmp64053, tmp64072)
														return

													} else {
														tmp64052 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														__e.TailApply(tmp64052)
														return

													}

												}, 1)

												tmp64083 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp64084 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

												tmp64085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp64086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

												tmp64087 := Call(__e, tmp64086, V220)

												tmp64088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp64089 := Call(__e, tmp64088, V220)

												tmp64090 := Call(__e, tmp64085, tmp64087, tmp64089)

												tmp64091 := Call(__e, tmp64084, tmp64090)

												tmp64092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp64093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp64094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

												tmp64095 := Call(__e, tmp64094, V220)

												tmp64096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp64097 := Call(__e, tmp64096, V220)

												tmp64098 := Call(__e, tmp64093, tmp64095, tmp64097)

												tmp64099 := Call(__e, tmp64092, tmp64098)

												tmp64100 := Call(__e, tmp64083, tmp64091, tmp64099)

												tmp64101 := Call(__e, tmp64050, tmp64100)

												ifres64047 = tmp64101

											} else {
												tmp64048 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												tmp64049 := Call(__e, tmp64048)

												ifres64047 = tmp64049

											}

											ifres64044 = ifres64047

										} else {
											tmp64045 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp64046 := Call(__e, tmp64045)

											ifres64044 = tmp64046

										}

										__e.TailApply(tmp63988, ifres64044)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp64233 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp64234 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp64235 := Call(__e, tmp64234, V220)

								tmp64236 := Call(__e, tmp64233, tmp64235)

								var ifres64227 Obj

								if True == tmp64236 {
									tmp64229 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp64230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

									tmp64231 := Call(__e, tmp64230, V220)

									tmp64232 := Call(__e, tmp64229, tmp64231)

									var ifres64228 Obj

									if True == tmp64232 {
										ifres64228 = True

									} else {
										ifres64228 = False

									}

									ifres64227 = ifres64228

								} else {
									ifres64227 = False

								}

								var ifres64138 Obj

								if True == ifres64227 {
									tmp64217 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp64218 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp64219 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

									tmp64220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

									tmp64221 := Call(__e, tmp64220, V220)

									tmp64222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp64223 := Call(__e, tmp64222, V220)

									tmp64224 := Call(__e, tmp64219, tmp64221, tmp64223)

									tmp64225 := Call(__e, tmp64218, tmp64224)

									tmp64226 := Call(__e, tmp64217, tmp64225)

									var ifres64205 Obj

									if True == tmp64226 {
										tmp64207 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp64208 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

										tmp64209 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp64210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

										tmp64211 := Call(__e, tmp64210, V220)

										tmp64212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp64213 := Call(__e, tmp64212, V220)

										tmp64214 := Call(__e, tmp64209, tmp64211, tmp64213)

										tmp64215 := Call(__e, tmp64208, tmp64214)

										tmp64216 := Call(__e, tmp64207, sym_8s, tmp64215)

										var ifres64206 Obj

										if True == tmp64216 {
											ifres64206 = True

										} else {
											ifres64206 = False

										}

										ifres64205 = ifres64206

									} else {
										ifres64205 = False

									}

									var ifres64141 Obj

									if True == ifres64205 {
										tmp64144 := MakeNative(func(__e *ControlFlow) {
											NewStream215 := __e.Get(1)
											_ = NewStream215
											tmp64145 := MakeNative(func(__e *ControlFlow) {
												Parse__shen_4_5pattern1_6 := __e.Get(1)
												_ = Parse__shen_4_5pattern1_6
												tmp64178 := Call(__e, PrimNS1Value(symns2_1value), symnot)

												tmp64179 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp64180 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												tmp64181 := Call(__e, tmp64180)

												tmp64182 := Call(__e, tmp64179, tmp64181, Parse__shen_4_5pattern1_6)

												tmp64183 := Call(__e, tmp64178, tmp64182)

												if True == tmp64183 {
													tmp64148 := MakeNative(func(__e *ControlFlow) {
														Parse__shen_4_5pattern2_6 := __e.Get(1)
														_ = Parse__shen_4_5pattern2_6
														tmp64170 := Call(__e, PrimNS1Value(symns2_1value), symnot)

														tmp64171 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp64172 := Call(__e, PrimNS1Value(symns2_1value), symfail)

														tmp64173 := Call(__e, tmp64172)

														tmp64174 := Call(__e, tmp64171, tmp64173, Parse__shen_4_5pattern2_6)

														tmp64175 := Call(__e, tmp64170, tmp64174)

														if True == tmp64175 {
															tmp64151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64152 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp64153 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

															tmp64154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

															tmp64155 := Call(__e, tmp64154, V220)

															tmp64156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

															tmp64157 := Call(__e, tmp64156, V220)

															tmp64158 := Call(__e, tmp64153, tmp64155, tmp64157)

															tmp64159 := Call(__e, tmp64152, tmp64158)

															tmp64160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp64161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp64162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

															tmp64163 := Call(__e, tmp64162, Parse__shen_4_5pattern1_6)

															tmp64164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp64165 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

															tmp64166 := Call(__e, tmp64165, Parse__shen_4_5pattern2_6)

															tmp64167 := Call(__e, tmp64164, tmp64166, Nil)

															tmp64168 := Call(__e, tmp64161, tmp64163, tmp64167)

															tmp64169 := Call(__e, tmp64160, sym_8s, tmp64168)

															__e.TailApply(tmp64151, tmp64159, tmp64169)
															return

														} else {
															tmp64150 := Call(__e, PrimNS1Value(symns2_1value), symfail)

															__e.TailApply(tmp64150)
															return

														}

													}, 1)

													tmp64176 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern2_6)

													tmp64177 := Call(__e, tmp64176, Parse__shen_4_5pattern1_6)

													__e.TailApply(tmp64148, tmp64177)
													return

												} else {
													tmp64147 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													__e.TailApply(tmp64147)
													return

												}

											}, 1)

											tmp64184 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern1_6)

											tmp64185 := Call(__e, tmp64184, NewStream215)

											__e.TailApply(tmp64145, tmp64185)
											return

										}, 1)

										tmp64186 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp64187 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

										tmp64188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp64189 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

										tmp64190 := Call(__e, tmp64189, V220)

										tmp64191 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp64192 := Call(__e, tmp64191, V220)

										tmp64193 := Call(__e, tmp64188, tmp64190, tmp64192)

										tmp64194 := Call(__e, tmp64187, tmp64193)

										tmp64195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp64196 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp64197 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

										tmp64198 := Call(__e, tmp64197, V220)

										tmp64199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp64200 := Call(__e, tmp64199, V220)

										tmp64201 := Call(__e, tmp64196, tmp64198, tmp64200)

										tmp64202 := Call(__e, tmp64195, tmp64201)

										tmp64203 := Call(__e, tmp64186, tmp64194, tmp64202)

										tmp64204 := Call(__e, tmp64144, tmp64203)

										ifres64141 = tmp64204

									} else {
										tmp64142 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										tmp64143 := Call(__e, tmp64142)

										ifres64141 = tmp64143

									}

									ifres64138 = ifres64141

								} else {
									tmp64139 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp64140 := Call(__e, tmp64139)

									ifres64138 = tmp64140

								}

								__e.TailApply(tmp63986, ifres64138)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp64336 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp64337 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp64338 := Call(__e, tmp64337, V220)

						tmp64339 := Call(__e, tmp64336, tmp64338)

						var ifres64330 Obj

						if True == tmp64339 {
							tmp64332 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp64333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp64334 := Call(__e, tmp64333, V220)

							tmp64335 := Call(__e, tmp64332, tmp64334)

							var ifres64331 Obj

							if True == tmp64335 {
								ifres64331 = True

							} else {
								ifres64331 = False

							}

							ifres64330 = ifres64331

						} else {
							ifres64330 = False

						}

						var ifres64241 Obj

						if True == ifres64330 {
							tmp64320 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp64321 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp64322 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp64323 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp64324 := Call(__e, tmp64323, V220)

							tmp64325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp64326 := Call(__e, tmp64325, V220)

							tmp64327 := Call(__e, tmp64322, tmp64324, tmp64326)

							tmp64328 := Call(__e, tmp64321, tmp64327)

							tmp64329 := Call(__e, tmp64320, tmp64328)

							var ifres64308 Obj

							if True == tmp64329 {
								tmp64310 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp64311 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

								tmp64312 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp64313 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

								tmp64314 := Call(__e, tmp64313, V220)

								tmp64315 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp64316 := Call(__e, tmp64315, V220)

								tmp64317 := Call(__e, tmp64312, tmp64314, tmp64316)

								tmp64318 := Call(__e, tmp64311, tmp64317)

								tmp64319 := Call(__e, tmp64310, sym_8v, tmp64318)

								var ifres64309 Obj

								if True == tmp64319 {
									ifres64309 = True

								} else {
									ifres64309 = False

								}

								ifres64308 = ifres64309

							} else {
								ifres64308 = False

							}

							var ifres64244 Obj

							if True == ifres64308 {
								tmp64247 := MakeNative(func(__e *ControlFlow) {
									NewStream213 := __e.Get(1)
									_ = NewStream213
									tmp64248 := MakeNative(func(__e *ControlFlow) {
										Parse__shen_4_5pattern1_6 := __e.Get(1)
										_ = Parse__shen_4_5pattern1_6
										tmp64281 := Call(__e, PrimNS1Value(symns2_1value), symnot)

										tmp64282 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp64283 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										tmp64284 := Call(__e, tmp64283)

										tmp64285 := Call(__e, tmp64282, tmp64284, Parse__shen_4_5pattern1_6)

										tmp64286 := Call(__e, tmp64281, tmp64285)

										if True == tmp64286 {
											tmp64251 := MakeNative(func(__e *ControlFlow) {
												Parse__shen_4_5pattern2_6 := __e.Get(1)
												_ = Parse__shen_4_5pattern2_6
												tmp64273 := Call(__e, PrimNS1Value(symns2_1value), symnot)

												tmp64274 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp64275 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												tmp64276 := Call(__e, tmp64275)

												tmp64277 := Call(__e, tmp64274, tmp64276, Parse__shen_4_5pattern2_6)

												tmp64278 := Call(__e, tmp64273, tmp64277)

												if True == tmp64278 {
													tmp64254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

													tmp64255 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp64256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

													tmp64257 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

													tmp64258 := Call(__e, tmp64257, V220)

													tmp64259 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

													tmp64260 := Call(__e, tmp64259, V220)

													tmp64261 := Call(__e, tmp64256, tmp64258, tmp64260)

													tmp64262 := Call(__e, tmp64255, tmp64261)

													tmp64263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													tmp64264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													tmp64265 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

													tmp64266 := Call(__e, tmp64265, Parse__shen_4_5pattern1_6)

													tmp64267 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													tmp64268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

													tmp64269 := Call(__e, tmp64268, Parse__shen_4_5pattern2_6)

													tmp64270 := Call(__e, tmp64267, tmp64269, Nil)

													tmp64271 := Call(__e, tmp64264, tmp64266, tmp64270)

													tmp64272 := Call(__e, tmp64263, sym_8v, tmp64271)

													__e.TailApply(tmp64254, tmp64262, tmp64272)
													return

												} else {
													tmp64253 := Call(__e, PrimNS1Value(symns2_1value), symfail)

													__e.TailApply(tmp64253)
													return

												}

											}, 1)

											tmp64279 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern2_6)

											tmp64280 := Call(__e, tmp64279, Parse__shen_4_5pattern1_6)

											__e.TailApply(tmp64251, tmp64280)
											return

										} else {
											tmp64250 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											__e.TailApply(tmp64250)
											return

										}

									}, 1)

									tmp64287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern1_6)

									tmp64288 := Call(__e, tmp64287, NewStream213)

									__e.TailApply(tmp64248, tmp64288)
									return

								}, 1)

								tmp64289 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp64290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

								tmp64291 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp64292 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

								tmp64293 := Call(__e, tmp64292, V220)

								tmp64294 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp64295 := Call(__e, tmp64294, V220)

								tmp64296 := Call(__e, tmp64291, tmp64293, tmp64295)

								tmp64297 := Call(__e, tmp64290, tmp64296)

								tmp64298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp64299 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp64300 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

								tmp64301 := Call(__e, tmp64300, V220)

								tmp64302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp64303 := Call(__e, tmp64302, V220)

								tmp64304 := Call(__e, tmp64299, tmp64301, tmp64303)

								tmp64305 := Call(__e, tmp64298, tmp64304)

								tmp64306 := Call(__e, tmp64289, tmp64297, tmp64305)

								tmp64307 := Call(__e, tmp64247, tmp64306)

								ifres64244 = tmp64307

							} else {
								tmp64245 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								tmp64246 := Call(__e, tmp64245)

								ifres64244 = tmp64246

							}

							ifres64241 = ifres64244

						} else {
							tmp64242 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp64243 := Call(__e, tmp64242)

							ifres64241 = tmp64243

						}

						__e.TailApply(tmp63984, ifres64241)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp64439 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp64440 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64441 := Call(__e, tmp64440, V220)

				tmp64442 := Call(__e, tmp64439, tmp64441)

				var ifres64433 Obj

				if True == tmp64442 {
					tmp64435 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp64436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp64437 := Call(__e, tmp64436, V220)

					tmp64438 := Call(__e, tmp64435, tmp64437)

					var ifres64434 Obj

					if True == tmp64438 {
						ifres64434 = True

					} else {
						ifres64434 = False

					}

					ifres64433 = ifres64434

				} else {
					ifres64433 = False

				}

				var ifres64344 Obj

				if True == ifres64433 {
					tmp64423 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp64424 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64425 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp64426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp64427 := Call(__e, tmp64426, V220)

					tmp64428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp64429 := Call(__e, tmp64428, V220)

					tmp64430 := Call(__e, tmp64425, tmp64427, tmp64429)

					tmp64431 := Call(__e, tmp64424, tmp64430)

					tmp64432 := Call(__e, tmp64423, tmp64431)

					var ifres64411 Obj

					if True == tmp64432 {
						tmp64413 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp64414 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

						tmp64415 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp64416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

						tmp64417 := Call(__e, tmp64416, V220)

						tmp64418 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp64419 := Call(__e, tmp64418, V220)

						tmp64420 := Call(__e, tmp64415, tmp64417, tmp64419)

						tmp64421 := Call(__e, tmp64414, tmp64420)

						tmp64422 := Call(__e, tmp64413, symcons, tmp64421)

						var ifres64412 Obj

						if True == tmp64422 {
							ifres64412 = True

						} else {
							ifres64412 = False

						}

						ifres64411 = ifres64412

					} else {
						ifres64411 = False

					}

					var ifres64347 Obj

					if True == ifres64411 {
						tmp64350 := MakeNative(func(__e *ControlFlow) {
							NewStream211 := __e.Get(1)
							_ = NewStream211
							tmp64351 := MakeNative(func(__e *ControlFlow) {
								Parse__shen_4_5pattern1_6 := __e.Get(1)
								_ = Parse__shen_4_5pattern1_6
								tmp64384 := Call(__e, PrimNS1Value(symns2_1value), symnot)

								tmp64385 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp64386 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								tmp64387 := Call(__e, tmp64386)

								tmp64388 := Call(__e, tmp64385, tmp64387, Parse__shen_4_5pattern1_6)

								tmp64389 := Call(__e, tmp64384, tmp64388)

								if True == tmp64389 {
									tmp64354 := MakeNative(func(__e *ControlFlow) {
										Parse__shen_4_5pattern2_6 := __e.Get(1)
										_ = Parse__shen_4_5pattern2_6
										tmp64376 := Call(__e, PrimNS1Value(symns2_1value), symnot)

										tmp64377 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp64378 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										tmp64379 := Call(__e, tmp64378)

										tmp64380 := Call(__e, tmp64377, tmp64379, Parse__shen_4_5pattern2_6)

										tmp64381 := Call(__e, tmp64376, tmp64380)

										if True == tmp64381 {
											tmp64357 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

											tmp64358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp64359 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

											tmp64360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

											tmp64361 := Call(__e, tmp64360, V220)

											tmp64362 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

											tmp64363 := Call(__e, tmp64362, V220)

											tmp64364 := Call(__e, tmp64359, tmp64361, tmp64363)

											tmp64365 := Call(__e, tmp64358, tmp64364)

											tmp64366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp64367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp64368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

											tmp64369 := Call(__e, tmp64368, Parse__shen_4_5pattern1_6)

											tmp64370 := Call(__e, PrimNS1Value(symns2_1value), symcons)

											tmp64371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

											tmp64372 := Call(__e, tmp64371, Parse__shen_4_5pattern2_6)

											tmp64373 := Call(__e, tmp64370, tmp64372, Nil)

											tmp64374 := Call(__e, tmp64367, tmp64369, tmp64373)

											tmp64375 := Call(__e, tmp64366, symcons, tmp64374)

											__e.TailApply(tmp64357, tmp64365, tmp64375)
											return

										} else {
											tmp64356 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											__e.TailApply(tmp64356)
											return

										}

									}, 1)

									tmp64382 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern2_6)

									tmp64383 := Call(__e, tmp64382, Parse__shen_4_5pattern1_6)

									__e.TailApply(tmp64354, tmp64383)
									return

								} else {
									tmp64353 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									__e.TailApply(tmp64353)
									return

								}

							}, 1)

							tmp64390 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern1_6)

							tmp64391 := Call(__e, tmp64390, NewStream211)

							__e.TailApply(tmp64351, tmp64391)
							return

						}, 1)

						tmp64392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp64393 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

						tmp64394 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp64395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

						tmp64396 := Call(__e, tmp64395, V220)

						tmp64397 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp64398 := Call(__e, tmp64397, V220)

						tmp64399 := Call(__e, tmp64394, tmp64396, tmp64398)

						tmp64400 := Call(__e, tmp64393, tmp64399)

						tmp64401 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp64402 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp64403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

						tmp64404 := Call(__e, tmp64403, V220)

						tmp64405 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp64406 := Call(__e, tmp64405, V220)

						tmp64407 := Call(__e, tmp64402, tmp64404, tmp64406)

						tmp64408 := Call(__e, tmp64401, tmp64407)

						tmp64409 := Call(__e, tmp64392, tmp64400, tmp64408)

						tmp64410 := Call(__e, tmp64350, tmp64409)

						ifres64347 = tmp64410

					} else {
						tmp64348 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						tmp64349 := Call(__e, tmp64348)

						ifres64347 = tmp64349

					}

					ifres64344 = ifres64347

				} else {
					tmp64345 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp64346 := Call(__e, tmp64345)

					ifres64344 = tmp64346

				}

				__e.TailApply(tmp63982, ifres64344)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp64542 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64543 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp64544 := Call(__e, tmp64543, V220)

		tmp64545 := Call(__e, tmp64542, tmp64544)

		var ifres64536 Obj

		if True == tmp64545 {
			tmp64538 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp64539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp64540 := Call(__e, tmp64539, V220)

			tmp64541 := Call(__e, tmp64538, tmp64540)

			var ifres64537 Obj

			if True == tmp64541 {
				ifres64537 = True

			} else {
				ifres64537 = False

			}

			ifres64536 = ifres64537

		} else {
			ifres64536 = False

		}

		var ifres64447 Obj

		if True == ifres64536 {
			tmp64526 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp64527 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp64528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp64529 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp64530 := Call(__e, tmp64529, V220)

			tmp64531 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp64532 := Call(__e, tmp64531, V220)

			tmp64533 := Call(__e, tmp64528, tmp64530, tmp64532)

			tmp64534 := Call(__e, tmp64527, tmp64533)

			tmp64535 := Call(__e, tmp64526, tmp64534)

			var ifres64514 Obj

			if True == tmp64535 {
				tmp64516 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp64517 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

				tmp64518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

				tmp64520 := Call(__e, tmp64519, V220)

				tmp64521 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64522 := Call(__e, tmp64521, V220)

				tmp64523 := Call(__e, tmp64518, tmp64520, tmp64522)

				tmp64524 := Call(__e, tmp64517, tmp64523)

				tmp64525 := Call(__e, tmp64516, sym_8p, tmp64524)

				var ifres64515 Obj

				if True == tmp64525 {
					ifres64515 = True

				} else {
					ifres64515 = False

				}

				ifres64514 = ifres64515

			} else {
				ifres64514 = False

			}

			var ifres64450 Obj

			if True == ifres64514 {
				tmp64453 := MakeNative(func(__e *ControlFlow) {
					NewStream209 := __e.Get(1)
					_ = NewStream209
					tmp64454 := MakeNative(func(__e *ControlFlow) {
						Parse__shen_4_5pattern1_6 := __e.Get(1)
						_ = Parse__shen_4_5pattern1_6
						tmp64487 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						tmp64488 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp64489 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						tmp64490 := Call(__e, tmp64489)

						tmp64491 := Call(__e, tmp64488, tmp64490, Parse__shen_4_5pattern1_6)

						tmp64492 := Call(__e, tmp64487, tmp64491)

						if True == tmp64492 {
							tmp64457 := MakeNative(func(__e *ControlFlow) {
								Parse__shen_4_5pattern2_6 := __e.Get(1)
								_ = Parse__shen_4_5pattern2_6
								tmp64479 := Call(__e, PrimNS1Value(symns2_1value), symnot)

								tmp64480 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp64481 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								tmp64482 := Call(__e, tmp64481)

								tmp64483 := Call(__e, tmp64480, tmp64482, Parse__shen_4_5pattern2_6)

								tmp64484 := Call(__e, tmp64479, tmp64483)

								if True == tmp64484 {
									tmp64460 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

									tmp64461 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp64462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

									tmp64463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

									tmp64464 := Call(__e, tmp64463, V220)

									tmp64465 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp64466 := Call(__e, tmp64465, V220)

									tmp64467 := Call(__e, tmp64462, tmp64464, tmp64466)

									tmp64468 := Call(__e, tmp64461, tmp64467)

									tmp64469 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp64470 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp64471 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp64472 := Call(__e, tmp64471, Parse__shen_4_5pattern1_6)

									tmp64473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp64474 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

									tmp64475 := Call(__e, tmp64474, Parse__shen_4_5pattern2_6)

									tmp64476 := Call(__e, tmp64473, tmp64475, Nil)

									tmp64477 := Call(__e, tmp64470, tmp64472, tmp64476)

									tmp64478 := Call(__e, tmp64469, sym_8p, tmp64477)

									__e.TailApply(tmp64460, tmp64468, tmp64478)
									return

								} else {
									tmp64459 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									__e.TailApply(tmp64459)
									return

								}

							}, 1)

							tmp64485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern2_6)

							tmp64486 := Call(__e, tmp64485, Parse__shen_4_5pattern1_6)

							__e.TailApply(tmp64457, tmp64486)
							return

						} else {
							tmp64456 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp64456)
							return

						}

					}, 1)

					tmp64493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern1_6)

					tmp64494 := Call(__e, tmp64493, NewStream209)

					__e.TailApply(tmp64454, tmp64494)
					return

				}, 1)

				tmp64495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp64497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

				tmp64499 := Call(__e, tmp64498, V220)

				tmp64500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64501 := Call(__e, tmp64500, V220)

				tmp64502 := Call(__e, tmp64497, tmp64499, tmp64501)

				tmp64503 := Call(__e, tmp64496, tmp64502)

				tmp64504 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64505 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64506 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

				tmp64507 := Call(__e, tmp64506, V220)

				tmp64508 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64509 := Call(__e, tmp64508, V220)

				tmp64510 := Call(__e, tmp64505, tmp64507, tmp64509)

				tmp64511 := Call(__e, tmp64504, tmp64510)

				tmp64512 := Call(__e, tmp64495, tmp64503, tmp64511)

				tmp64513 := Call(__e, tmp64453, tmp64512)

				ifres64450 = tmp64513

			} else {
				tmp64451 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				tmp64452 := Call(__e, tmp64451)

				ifres64450 = tmp64452

			}

			ifres64447 = ifres64450

		} else {
			tmp64448 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64449 := Call(__e, tmp64448)

			ifres64447 = tmp64449

		}

		__e.TailApply(tmp63980, ifres64447)
		return

	}, 1)

	tmp64546 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern_6, tmp63979)

	_ = tmp64546

	tmp64547 := MakeNative(func(__e *ControlFlow) {
		V222 := __e.Get(1)
		_ = V222
		tmp64548 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		tmp64549 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp64550 := Call(__e, tmp64549, V222, MakeString(" is not a legitimate constructor\n"), symshen_4a)

		__e.TailApply(tmp64548, tmp64550)
		return

	}, 1)

	tmp64551 := Call(__e, PrimNS1Value(symns2_1set), symshen_4constructor_1error, tmp64547)

	_ = tmp64551

	tmp64552 := MakeNative(func(__e *ControlFlow) {
		V224 := __e.Get(1)
		_ = V224
		tmp64553 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp64583 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp64584 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64585 := Call(__e, tmp64584)

			tmp64586 := Call(__e, tmp64583, YaccParse, tmp64585)

			if True == tmp64586 {
				tmp64579 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp64580 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64581 := Call(__e, tmp64580, V224)

				tmp64582 := Call(__e, tmp64579, tmp64581)

				if True == tmp64582 {
					tmp64557 := MakeNative(func(__e *ControlFlow) {
						Parse__X := __e.Get(1)
						_ = Parse__X
						tmp64569 := Call(__e, PrimNS1Value(symns2_1value), symnot)

						tmp64570 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

						tmp64571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp64572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp64573 := Call(__e, tmp64572, sym_5_1, Nil)

						tmp64574 := Call(__e, tmp64571, sym_1_6, tmp64573)

						tmp64575 := Call(__e, tmp64570, Parse__X, tmp64574)

						tmp64576 := Call(__e, tmp64569, tmp64575)

						if True == tmp64576 {
							tmp64560 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp64561 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp64562 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp64563 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							tmp64564 := Call(__e, tmp64563, V224)

							tmp64565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp64566 := Call(__e, tmp64565, V224)

							tmp64567 := Call(__e, tmp64562, tmp64564, tmp64566)

							tmp64568 := Call(__e, tmp64561, tmp64567)

							__e.TailApply(tmp64560, tmp64568, Parse__X)
							return

						} else {
							tmp64559 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp64559)
							return

						}

					}, 1)

					tmp64577 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp64578 := Call(__e, tmp64577, V224)

					__e.TailApply(tmp64557, tmp64578)
					return

				} else {
					tmp64556 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp64556)
					return

				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp64609 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64610 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp64611 := Call(__e, tmp64610, V224)

		tmp64612 := Call(__e, tmp64609, tmp64611)

		var ifres64587 Obj

		if True == tmp64612 {
			tmp64590 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp64604 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp64605 := Call(__e, tmp64604, Parse__X, sym__)

				if True == tmp64605 {
					tmp64593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp64594 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp64596 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp64597 := Call(__e, tmp64596, V224)

					tmp64598 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp64599 := Call(__e, tmp64598, V224)

					tmp64600 := Call(__e, tmp64595, tmp64597, tmp64599)

					tmp64601 := Call(__e, tmp64594, tmp64600)

					tmp64602 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

					tmp64603 := Call(__e, tmp64602, symParse__Y)

					__e.TailApply(tmp64593, tmp64601, tmp64603)
					return

				} else {
					tmp64592 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp64592)
					return

				}

			}, 1)

			tmp64606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp64607 := Call(__e, tmp64606, V224)

			tmp64608 := Call(__e, tmp64590, tmp64607)

			ifres64587 = tmp64608

		} else {
			tmp64588 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64589 := Call(__e, tmp64588)

			ifres64587 = tmp64589

		}

		__e.TailApply(tmp64553, ifres64587)
		return

	}, 1)

	tmp64613 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5simple__pattern_6, tmp64552)

	_ = tmp64613

	tmp64614 := MakeNative(func(__e *ControlFlow) {
		V226 := __e.Get(1)
		_ = V226
		tmp64615 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp64623 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp64624 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp64625 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64626 := Call(__e, tmp64625)

			tmp64627 := Call(__e, tmp64624, tmp64626, Parse__shen_4_5pattern_6)

			tmp64628 := Call(__e, tmp64623, tmp64627)

			if True == tmp64628 {
				tmp64618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64619 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64620 := Call(__e, tmp64619, Parse__shen_4_5pattern_6)

				tmp64621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64622 := Call(__e, tmp64621, Parse__shen_4_5pattern_6)

				__e.TailApply(tmp64618, tmp64620, tmp64622)
				return

			} else {
				tmp64617 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp64617)
				return

			}

		}, 1)

		tmp64629 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern_6)

		tmp64630 := Call(__e, tmp64629, V226)

		__e.TailApply(tmp64615, tmp64630)
		return

	}, 1)

	tmp64631 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern1_6, tmp64614)

	_ = tmp64631

	tmp64632 := MakeNative(func(__e *ControlFlow) {
		V228 := __e.Get(1)
		_ = V228
		tmp64633 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp64641 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp64642 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp64643 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp64644 := Call(__e, tmp64643)

			tmp64645 := Call(__e, tmp64642, tmp64644, Parse__shen_4_5pattern_6)

			tmp64646 := Call(__e, tmp64641, tmp64645)

			if True == tmp64646 {
				tmp64636 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64638 := Call(__e, tmp64637, Parse__shen_4_5pattern_6)

				tmp64639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64640 := Call(__e, tmp64639, Parse__shen_4_5pattern_6)

				__e.TailApply(tmp64636, tmp64638, tmp64640)
				return

			} else {
				tmp64635 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp64635)
				return

			}

		}, 1)

		tmp64647 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5pattern_6)

		tmp64648 := Call(__e, tmp64647, V228)

		__e.TailApply(tmp64633, tmp64648)
		return

	}, 1)

	tmp64649 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern2_6, tmp64632)

	_ = tmp64649

	tmp64650 := MakeNative(func(__e *ControlFlow) {
		V230 := __e.Get(1)
		_ = V230
		tmp64665 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64666 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp64667 := Call(__e, tmp64666, V230)

		tmp64668 := Call(__e, tmp64665, tmp64667)

		if True == tmp64668 {
			tmp64653 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp64654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64655 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64656 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp64658 := Call(__e, tmp64657, V230)

				tmp64659 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64660 := Call(__e, tmp64659, V230)

				tmp64661 := Call(__e, tmp64656, tmp64658, tmp64660)

				tmp64662 := Call(__e, tmp64655, tmp64661)

				__e.TailApply(tmp64654, tmp64662, Parse__X)
				return

			}, 1)

			tmp64663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp64664 := Call(__e, tmp64663, V230)

			__e.TailApply(tmp64653, tmp64664)
			return

		} else {
			tmp64652 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp64652)
			return

		}

	}, 1)

	tmp64669 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5action_6, tmp64650)

	_ = tmp64669

	tmp64670 := MakeNative(func(__e *ControlFlow) {
		V232 := __e.Get(1)
		_ = V232
		tmp64685 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp64687 := Call(__e, tmp64686, V232)

		tmp64688 := Call(__e, tmp64685, tmp64687)

		if True == tmp64688 {
			tmp64673 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp64674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64675 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64676 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp64677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

				tmp64678 := Call(__e, tmp64677, V232)

				tmp64679 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp64680 := Call(__e, tmp64679, V232)

				tmp64681 := Call(__e, tmp64676, tmp64678, tmp64680)

				tmp64682 := Call(__e, tmp64675, tmp64681)

				__e.TailApply(tmp64674, tmp64682, Parse__X)
				return

			}, 1)

			tmp64683 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp64684 := Call(__e, tmp64683, V232)

			__e.TailApply(tmp64673, tmp64684)
			return

		} else {
			tmp64672 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp64672)
			return

		}

	}, 1)

	tmp64689 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5guard_6, tmp64670)

	_ = tmp64689

	tmp64690 := MakeNative(func(__e *ControlFlow) {
		V235 := __e.Get(1)
		_ = V235
		V236 := __e.Get(2)
		_ = V236
		tmp64691 := MakeNative(func(__e *ControlFlow) {
			Lambda_7 := __e.Get(1)
			_ = Lambda_7
			tmp64692 := MakeNative(func(__e *ControlFlow) {
				KL := __e.Get(1)
				_ = KL
				tmp64693 := MakeNative(func(__e *ControlFlow) {
					Record := __e.Get(1)
					_ = Record
					__e.Return(KL)
					return
				}, 1)

				tmp64694 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1source)

				tmp64695 := Call(__e, tmp64694, V235, KL)

				__e.TailApply(tmp64693, tmp64695)
				return

			}, 1)

			tmp64696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compile__to__kl)

			tmp64697 := Call(__e, tmp64696, V235, Lambda_7)

			__e.TailApply(tmp64692, tmp64697)
			return

		}, 1)

		tmp64698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compile__to__lambda_7)

		tmp64699 := Call(__e, tmp64698, V235, V236)

		__e.TailApply(tmp64691, tmp64699)
		return

	}, 2)

	tmp64700 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__machine__code, tmp64690)

	_ = tmp64700

	tmp64701 := MakeNative(func(__e *ControlFlow) {
		V241 := __e.Get(1)
		_ = V241
		V242 := __e.Get(2)
		_ = V242
		tmp64706 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp64707 := Call(__e, tmp64706, symshen_4_dinstalling_1kl_d)

		if True == tmp64707 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp64703 := Call(__e, PrimNS1Value(symns2_1value), symput)

			tmp64704 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp64705 := Call(__e, tmp64704, sym_dproperty_1vector_d)

			__e.TailApply(tmp64703, V241, symshen_4source, V242, tmp64705)
			return

		}

	}, 2)

	tmp64708 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1source, tmp64701)

	_ = tmp64708

	tmp64709 := MakeNative(func(__e *ControlFlow) {
		V245 := __e.Get(1)
		_ = V245
		V246 := __e.Get(2)
		_ = V246
		tmp64710 := MakeNative(func(__e *ControlFlow) {
			Arity := __e.Get(1)
			_ = Arity
			tmp64711 := MakeNative(func(__e *ControlFlow) {
				UpDateSymbolTable := __e.Get(1)
				_ = UpDateSymbolTable
				tmp64712 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp64713 := MakeNative(func(__e *ControlFlow) {
						Variables := __e.Get(1)
						_ = Variables
						tmp64714 := MakeNative(func(__e *ControlFlow) {
							Strip := __e.Get(1)
							_ = Strip
							tmp64715 := MakeNative(func(__e *ControlFlow) {
								Abstractions := __e.Get(1)
								_ = Abstractions
								tmp64716 := MakeNative(func(__e *ControlFlow) {
									Applications := __e.Get(1)
									_ = Applications
									tmp64717 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp64718 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp64719 := Call(__e, tmp64718, Applications, Nil)

									__e.TailApply(tmp64717, Variables, tmp64719)
									return

								}, 1)

								tmp64720 := Call(__e, PrimNS1Value(symns2_1value), symmap)

								tmp64721 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									tmp64722 := Call(__e, PrimNS1Value(symns2_1value), symshen_4application__build)

									__e.TailApply(tmp64722, Variables, X)
									return

								}, 1)

								tmp64723 := Call(__e, tmp64720, tmp64721, Abstractions)

								__e.TailApply(tmp64716, tmp64723)
								return

							}, 1)

							tmp64724 := Call(__e, PrimNS1Value(symns2_1value), symmap)

							tmp64725 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp64726 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abstract__rule)

								__e.TailApply(tmp64726, X)
								return

							}, 1)

							tmp64727 := Call(__e, tmp64724, tmp64725, Strip)

							__e.TailApply(tmp64715, tmp64727)
							return

						}, 1)

						tmp64728 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						tmp64729 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							tmp64730 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1protect)

							__e.TailApply(tmp64730, X)
							return

						}, 1)

						tmp64731 := Call(__e, tmp64728, tmp64729, V246)

						__e.TailApply(tmp64714, tmp64731)
						return

					}, 1)

					tmp64732 := Call(__e, PrimNS1Value(symns2_1value), symshen_4parameters)

					tmp64733 := Call(__e, tmp64732, Arity)

					__e.TailApply(tmp64713, tmp64733)
					return

				}, 1)

				tmp64734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

				tmp64735 := MakeNative(func(__e *ControlFlow) {
					Rule := __e.Get(1)
					_ = Rule
					tmp64736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4free__variable__check)

					__e.TailApply(tmp64736, V245, Rule)
					return

				}, 1)

				tmp64737 := Call(__e, tmp64734, tmp64735, V246)

				__e.TailApply(tmp64712, tmp64737)
				return

			}, 1)

			tmp64738 := Call(__e, PrimNS1Value(symns2_1value), symshen_4update_1symbol_1table)

			tmp64739 := Call(__e, tmp64738, V245, Arity)

			__e.TailApply(tmp64711, tmp64739)
			return

		}, 1)

		tmp64740 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck)

		tmp64741 := Call(__e, tmp64740, V245, V246)

		__e.TailApply(tmp64710, tmp64741)
		return

	}, 2)

	tmp64742 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__lambda_7, tmp64709)

	_ = tmp64742

	tmp64743 := MakeNative(func(__e *ControlFlow) {
		V249 := __e.Get(1)
		_ = V249
		V250 := __e.Get(2)
		_ = V250
		tmp64752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp64753 := Call(__e, tmp64752, MakeNumber(0), V250)

		if True == tmp64753 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp64745 := Call(__e, PrimNS1Value(symns2_1value), symput)

			tmp64746 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

			tmp64747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1form)

			tmp64748 := Call(__e, tmp64747, V249, V250)

			tmp64749 := Call(__e, tmp64746, tmp64748)

			tmp64750 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp64751 := Call(__e, tmp64750, sym_dproperty_1vector_d)

			__e.TailApply(tmp64745, V249, symshen_4lambda_1form, tmp64749, tmp64751)
			return

		}

	}, 2)

	tmp64754 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update_1symbol_1table, tmp64743)

	_ = tmp64754

	tmp64755 := MakeNative(func(__e *ControlFlow) {
		V253 := __e.Get(1)
		_ = V253
		V254 := __e.Get(2)
		_ = V254
		tmp64785 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64786 := Call(__e, tmp64785, V254)

		var ifres64771 Obj

		if True == tmp64786 {
			tmp64781 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp64782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp64783 := Call(__e, tmp64782, V254)

			tmp64784 := Call(__e, tmp64781, tmp64783)

			var ifres64773 Obj

			if True == tmp64784 {
				tmp64775 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp64776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64778 := Call(__e, tmp64777, V254)

				tmp64779 := Call(__e, tmp64776, tmp64778)

				tmp64780 := Call(__e, tmp64775, Nil, tmp64779)

				var ifres64774 Obj

				if True == tmp64780 {
					ifres64774 = True

				} else {
					ifres64774 = False

				}

				ifres64773 = ifres64774

			} else {
				ifres64773 = False

			}

			var ifres64772 Obj

			if True == ifres64773 {
				ifres64772 = True

			} else {
				ifres64772 = False

			}

			ifres64771 = ifres64772

		} else {
			ifres64771 = False

		}

		if True == ifres64771 {
			tmp64758 := MakeNative(func(__e *ControlFlow) {
				Bound := __e.Get(1)
				_ = Bound
				tmp64759 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp64760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4free__variable__warnings)

					__e.TailApply(tmp64760, V253, Free)
					return

				}, 1)

				tmp64761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

				tmp64762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64764 := Call(__e, tmp64763, V254)

				tmp64765 := Call(__e, tmp64762, tmp64764)

				tmp64766 := Call(__e, tmp64761, Bound, tmp64765)

				__e.TailApply(tmp64759, tmp64766)
				return

			}, 1)

			tmp64767 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

			tmp64768 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp64769 := Call(__e, tmp64768, V254)

			tmp64770 := Call(__e, tmp64767, tmp64769)

			__e.TailApply(tmp64758, tmp64770)
			return

		} else {
			tmp64757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp64757, symshen_4free__variable__check)
			return

		}

	}, 2)

	tmp64787 := Call(__e, PrimNS1Value(symns2_1set), symshen_4free__variable__check, tmp64755)

	_ = tmp64787

	tmp64788 := MakeNative(func(__e *ControlFlow) {
		V256 := __e.Get(1)
		_ = V256
		tmp64803 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

		tmp64804 := Call(__e, tmp64803, V256)

		if True == tmp64804 {
			tmp64802 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp64802, V256, Nil)
			return

		} else {
			tmp64800 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp64801 := Call(__e, tmp64800, V256)

			if True == tmp64801 {
				tmp64791 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp64792 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				tmp64793 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64794 := Call(__e, tmp64793, V256)

				tmp64795 := Call(__e, tmp64792, tmp64794)

				tmp64796 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				tmp64797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64798 := Call(__e, tmp64797, V256)

				tmp64799 := Call(__e, tmp64796, tmp64798)

				__e.TailApply(tmp64791, tmp64795, tmp64799)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp64805 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract__vars, tmp64788)

	_ = tmp64805

	tmp64806 := MakeNative(func(__e *ControlFlow) {
		V268 := __e.Get(1)
		_ = V268
		V269 := __e.Get(2)
		_ = V269
		tmp64966 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp64967 := Call(__e, tmp64966, V269)

		var ifres64946 Obj

		if True == tmp64967 {
			tmp64962 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp64963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp64964 := Call(__e, tmp64963, V269)

			tmp64965 := Call(__e, tmp64962, tmp64964)

			var ifres64948 Obj

			if True == tmp64965 {
				tmp64956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp64957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64959 := Call(__e, tmp64958, V269)

				tmp64960 := Call(__e, tmp64957, tmp64959)

				tmp64961 := Call(__e, tmp64956, Nil, tmp64960)

				var ifres64950 Obj

				if True == tmp64961 {
					tmp64952 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp64953 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64954 := Call(__e, tmp64953, V269)

					tmp64955 := Call(__e, tmp64952, tmp64954, symprotect)

					var ifres64951 Obj

					if True == tmp64955 {
						ifres64951 = True

					} else {
						ifres64951 = False

					}

					ifres64950 = ifres64951

				} else {
					ifres64950 = False

				}

				var ifres64949 Obj

				if True == ifres64950 {
					ifres64949 = True

				} else {
					ifres64949 = False

				}

				ifres64948 = ifres64949

			} else {
				ifres64948 = False

			}

			var ifres64947 Obj

			if True == ifres64948 {
				ifres64947 = True

			} else {
				ifres64947 = False

			}

			ifres64946 = ifres64947

		} else {
			ifres64946 = False

		}

		if True == ifres64946 {
			__e.Return(Nil)
			return
		} else {
			tmp64944 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp64945 := Call(__e, tmp64944, V269)

			var ifres64938 Obj

			if True == tmp64945 {
				tmp64940 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp64941 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp64942 := Call(__e, tmp64941, V269, V268)

				tmp64943 := Call(__e, tmp64940, tmp64942)

				var ifres64939 Obj

				if True == tmp64943 {
					ifres64939 = True

				} else {
					ifres64939 = False

				}

				ifres64938 = ifres64939

			} else {
				ifres64938 = False

			}

			if True == ifres64938 {
				tmp64937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				__e.TailApply(tmp64937, V269, Nil)
				return

			} else {
				tmp64935 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp64936 := Call(__e, tmp64935, V269)

				var ifres64905 Obj

				if True == tmp64936 {
					tmp64931 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp64932 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64933 := Call(__e, tmp64932, V269)

					tmp64934 := Call(__e, tmp64931, symlambda, tmp64933)

					var ifres64907 Obj

					if True == tmp64934 {
						tmp64927 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp64928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64929 := Call(__e, tmp64928, V269)

						tmp64930 := Call(__e, tmp64927, tmp64929)

						var ifres64909 Obj

						if True == tmp64930 {
							tmp64921 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp64922 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp64923 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp64924 := Call(__e, tmp64923, V269)

							tmp64925 := Call(__e, tmp64922, tmp64924)

							tmp64926 := Call(__e, tmp64921, tmp64925)

							var ifres64911 Obj

							if True == tmp64926 {
								tmp64913 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp64914 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp64915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp64916 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp64917 := Call(__e, tmp64916, V269)

								tmp64918 := Call(__e, tmp64915, tmp64917)

								tmp64919 := Call(__e, tmp64914, tmp64918)

								tmp64920 := Call(__e, tmp64913, Nil, tmp64919)

								var ifres64912 Obj

								if True == tmp64920 {
									ifres64912 = True

								} else {
									ifres64912 = False

								}

								ifres64911 = ifres64912

							} else {
								ifres64911 = False

							}

							var ifres64910 Obj

							if True == ifres64911 {
								ifres64910 = True

							} else {
								ifres64910 = False

							}

							ifres64909 = ifres64910

						} else {
							ifres64909 = False

						}

						var ifres64908 Obj

						if True == ifres64909 {
							ifres64908 = True

						} else {
							ifres64908 = False

						}

						ifres64907 = ifres64908

					} else {
						ifres64907 = False

					}

					var ifres64906 Obj

					if True == ifres64907 {
						ifres64906 = True

					} else {
						ifres64906 = False

					}

					ifres64905 = ifres64906

				} else {
					ifres64905 = False

				}

				if True == ifres64905 {
					tmp64892 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

					tmp64893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp64894 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64895 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp64896 := Call(__e, tmp64895, V269)

					tmp64897 := Call(__e, tmp64894, tmp64896)

					tmp64898 := Call(__e, tmp64893, tmp64897, V268)

					tmp64899 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp64900 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp64901 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp64902 := Call(__e, tmp64901, V269)

					tmp64903 := Call(__e, tmp64900, tmp64902)

					tmp64904 := Call(__e, tmp64899, tmp64903)

					__e.TailApply(tmp64892, tmp64898, tmp64904)
					return

				} else {
					tmp64890 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp64891 := Call(__e, tmp64890, V269)

					var ifres64848 Obj

					if True == tmp64891 {
						tmp64886 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp64887 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp64888 := Call(__e, tmp64887, V269)

						tmp64889 := Call(__e, tmp64886, symlet, tmp64888)

						var ifres64850 Obj

						if True == tmp64889 {
							tmp64882 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp64883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp64884 := Call(__e, tmp64883, V269)

							tmp64885 := Call(__e, tmp64882, tmp64884)

							var ifres64852 Obj

							if True == tmp64885 {
								tmp64876 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp64877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp64878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp64879 := Call(__e, tmp64878, V269)

								tmp64880 := Call(__e, tmp64877, tmp64879)

								tmp64881 := Call(__e, tmp64876, tmp64880)

								var ifres64854 Obj

								if True == tmp64881 {
									tmp64868 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp64869 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp64870 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp64871 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp64872 := Call(__e, tmp64871, V269)

									tmp64873 := Call(__e, tmp64870, tmp64872)

									tmp64874 := Call(__e, tmp64869, tmp64873)

									tmp64875 := Call(__e, tmp64868, tmp64874)

									var ifres64856 Obj

									if True == tmp64875 {
										tmp64858 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp64859 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp64860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp64861 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp64862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp64863 := Call(__e, tmp64862, V269)

										tmp64864 := Call(__e, tmp64861, tmp64863)

										tmp64865 := Call(__e, tmp64860, tmp64864)

										tmp64866 := Call(__e, tmp64859, tmp64865)

										tmp64867 := Call(__e, tmp64858, Nil, tmp64866)

										var ifres64857 Obj

										if True == tmp64867 {
											ifres64857 = True

										} else {
											ifres64857 = False

										}

										ifres64856 = ifres64857

									} else {
										ifres64856 = False

									}

									var ifres64855 Obj

									if True == ifres64856 {
										ifres64855 = True

									} else {
										ifres64855 = False

									}

									ifres64854 = ifres64855

								} else {
									ifres64854 = False

								}

								var ifres64853 Obj

								if True == ifres64854 {
									ifres64853 = True

								} else {
									ifres64853 = False

								}

								ifres64852 = ifres64853

							} else {
								ifres64852 = False

							}

							var ifres64851 Obj

							if True == ifres64852 {
								ifres64851 = True

							} else {
								ifres64851 = False

							}

							ifres64850 = ifres64851

						} else {
							ifres64850 = False

						}

						var ifres64849 Obj

						if True == ifres64850 {
							ifres64849 = True

						} else {
							ifres64849 = False

						}

						ifres64848 = ifres64849

					} else {
						ifres64848 = False

					}

					if True == ifres64848 {
						tmp64823 := Call(__e, PrimNS1Value(symns2_1value), symunion)

						tmp64824 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

						tmp64825 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp64826 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64827 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64828 := Call(__e, tmp64827, V269)

						tmp64829 := Call(__e, tmp64826, tmp64828)

						tmp64830 := Call(__e, tmp64825, tmp64829)

						tmp64831 := Call(__e, tmp64824, V268, tmp64830)

						tmp64832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

						tmp64833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp64834 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp64835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64836 := Call(__e, tmp64835, V269)

						tmp64837 := Call(__e, tmp64834, tmp64836)

						tmp64838 := Call(__e, tmp64833, tmp64837, V268)

						tmp64839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp64840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64841 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64842 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp64843 := Call(__e, tmp64842, V269)

						tmp64844 := Call(__e, tmp64841, tmp64843)

						tmp64845 := Call(__e, tmp64840, tmp64844)

						tmp64846 := Call(__e, tmp64839, tmp64845)

						tmp64847 := Call(__e, tmp64832, tmp64838, tmp64846)

						__e.TailApply(tmp64823, tmp64831, tmp64847)
						return

					} else {
						tmp64821 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp64822 := Call(__e, tmp64821, V269)

						if True == tmp64822 {
							tmp64812 := Call(__e, PrimNS1Value(symns2_1value), symunion)

							tmp64813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

							tmp64814 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp64815 := Call(__e, tmp64814, V269)

							tmp64816 := Call(__e, tmp64813, V268, tmp64815)

							tmp64817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__free__vars)

							tmp64818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp64819 := Call(__e, tmp64818, V269)

							tmp64820 := Call(__e, tmp64817, V268, tmp64819)

							__e.TailApply(tmp64812, tmp64816, tmp64820)
							return

						} else {
							__e.Return(Nil)
							return
						}

					}

				}

			}

		}

	}, 2)

	tmp64968 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract__free__vars, tmp64806)

	_ = tmp64968

	tmp64969 := MakeNative(func(__e *ControlFlow) {
		V274 := __e.Get(1)
		_ = V274
		V275 := __e.Get(2)
		_ = V275
		tmp64982 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp64983 := Call(__e, tmp64982, Nil, V275)

		if True == tmp64983 {
			__e.Return(sym__)
			return
		} else {
			tmp64971 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp64972 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp64973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp64974 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp64975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp64976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list__variables)

			tmp64977 := Call(__e, tmp64976, V275)

			tmp64978 := Call(__e, tmp64975, tmp64977, MakeString(""), symshen_4a)

			tmp64979 := Call(__e, tmp64974, MakeString(": "), tmp64978)

			tmp64980 := Call(__e, tmp64973, V274, tmp64979, symshen_4a)

			tmp64981 := Call(__e, tmp64972, MakeString("error: the following variables are free in "), tmp64980)

			__e.TailApply(tmp64971, tmp64981)
			return

		}

	}, 2)

	tmp64984 := Call(__e, PrimNS1Value(symns2_1set), symshen_4free__variable__warnings, tmp64969)

	_ = tmp64984

	tmp64985 := MakeNative(func(__e *ControlFlow) {
		V277 := __e.Get(1)
		_ = V277
		tmp65013 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65014 := Call(__e, tmp65013, V277)

		var ifres65007 Obj

		if True == tmp65014 {
			tmp65009 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp65010 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65011 := Call(__e, tmp65010, V277)

			tmp65012 := Call(__e, tmp65009, Nil, tmp65011)

			var ifres65008 Obj

			if True == tmp65012 {
				ifres65008 = True

			} else {
				ifres65008 = False

			}

			ifres65007 = ifres65008

		} else {
			ifres65007 = False

		}

		if True == ifres65007 {
			tmp65002 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp65003 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			tmp65004 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65005 := Call(__e, tmp65004, V277)

			tmp65006 := Call(__e, tmp65003, tmp65005)

			__e.TailApply(tmp65002, tmp65006, MakeString("."))
			return

		} else {
			tmp65000 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65001 := Call(__e, tmp65000, V277)

			if True == tmp65001 {
				tmp64989 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp64990 := Call(__e, PrimNS1Value(symns2_1value), symstr)

				tmp64991 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp64992 := Call(__e, tmp64991, V277)

				tmp64993 := Call(__e, tmp64990, tmp64992)

				tmp64994 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp64995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list__variables)

				tmp64996 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp64997 := Call(__e, tmp64996, V277)

				tmp64998 := Call(__e, tmp64995, tmp64997)

				tmp64999 := Call(__e, tmp64994, MakeString(", "), tmp64998)

				__e.TailApply(tmp64989, tmp64993, tmp64999)
				return

			} else {
				tmp64988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp64988, symshen_4list__variables)
				return

			}

		}

	}, 1)

	tmp65015 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list__variables, tmp64985)

	_ = tmp65015

	tmp65016 := MakeNative(func(__e *ControlFlow) {
		V279 := __e.Get(1)
		_ = V279
		tmp65049 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65050 := Call(__e, tmp65049, V279)

		var ifres65029 Obj

		if True == tmp65050 {
			tmp65045 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65046 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65047 := Call(__e, tmp65046, V279)

			tmp65048 := Call(__e, tmp65045, tmp65047)

			var ifres65031 Obj

			if True == tmp65048 {
				tmp65039 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65040 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65041 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65042 := Call(__e, tmp65041, V279)

				tmp65043 := Call(__e, tmp65040, tmp65042)

				tmp65044 := Call(__e, tmp65039, Nil, tmp65043)

				var ifres65033 Obj

				if True == tmp65044 {
					tmp65035 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp65036 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65037 := Call(__e, tmp65036, V279)

					tmp65038 := Call(__e, tmp65035, tmp65037, symprotect)

					var ifres65034 Obj

					if True == tmp65038 {
						ifres65034 = True

					} else {
						ifres65034 = False

					}

					ifres65033 = ifres65034

				} else {
					ifres65033 = False

				}

				var ifres65032 Obj

				if True == ifres65033 {
					ifres65032 = True

				} else {
					ifres65032 = False

				}

				ifres65031 = ifres65032

			} else {
				ifres65031 = False

			}

			var ifres65030 Obj

			if True == ifres65031 {
				ifres65030 = True

			} else {
				ifres65030 = False

			}

			ifres65029 = ifres65030

		} else {
			ifres65029 = False

		}

		if True == ifres65029 {
			tmp65024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1protect)

			tmp65025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65027 := Call(__e, tmp65026, V279)

			tmp65028 := Call(__e, tmp65025, tmp65027)

			__e.TailApply(tmp65024, tmp65028)
			return

		} else {
			tmp65022 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65023 := Call(__e, tmp65022, V279)

			if True == tmp65023 {
				tmp65019 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp65020 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp65021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1protect)

					__e.TailApply(tmp65021, Z)
					return

				}, 1)

				__e.TailApply(tmp65019, tmp65020, V279)
				return

			} else {
				__e.Return(V279)
				return
			}

		}

	}, 1)

	tmp65051 := Call(__e, PrimNS1Value(symns2_1set), symshen_4strip_1protect, tmp65016)

	_ = tmp65051

	tmp65052 := MakeNative(func(__e *ControlFlow) {
		V281 := __e.Get(1)
		_ = V281
		tmp65080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65081 := Call(__e, tmp65080, V281)

		var ifres65066 Obj

		if True == tmp65081 {
			tmp65076 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65077 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65078 := Call(__e, tmp65077, V281)

			tmp65079 := Call(__e, tmp65076, tmp65078)

			var ifres65068 Obj

			if True == tmp65079 {
				tmp65070 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65072 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65073 := Call(__e, tmp65072, V281)

				tmp65074 := Call(__e, tmp65071, tmp65073)

				tmp65075 := Call(__e, tmp65070, Nil, tmp65074)

				var ifres65069 Obj

				if True == tmp65075 {
					ifres65069 = True

				} else {
					ifres65069 = False

				}

				ifres65068 = ifres65069

			} else {
				ifres65068 = False

			}

			var ifres65067 Obj

			if True == ifres65068 {
				ifres65067 = True

			} else {
				ifres65067 = False

			}

			ifres65066 = ifres65067

		} else {
			ifres65066 = False

		}

		if True == ifres65066 {
			tmp65055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__help)

			tmp65056 := Call(__e, PrimNS1Value(symns2_1value), symshen_4flatten)

			tmp65057 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65058 := Call(__e, tmp65057, V281)

			tmp65059 := Call(__e, tmp65056, tmp65058)

			tmp65060 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65061 := Call(__e, tmp65060, V281)

			tmp65062 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65063 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65064 := Call(__e, tmp65063, V281)

			tmp65065 := Call(__e, tmp65062, tmp65064)

			__e.TailApply(tmp65055, tmp65059, tmp65061, tmp65065)
			return

		} else {
			tmp65054 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp65054, symshen_4linearise)
			return

		}

	}, 1)

	tmp65082 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise, tmp65052)

	_ = tmp65082

	tmp65083 := MakeNative(func(__e *ControlFlow) {
		V283 := __e.Get(1)
		_ = V283
		tmp65098 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65099 := Call(__e, tmp65098, Nil, V283)

		if True == tmp65099 {
			__e.Return(Nil)
			return
		} else {
			tmp65096 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65097 := Call(__e, tmp65096, V283)

			if True == tmp65097 {
				tmp65087 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp65088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4flatten)

				tmp65089 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65090 := Call(__e, tmp65089, V283)

				tmp65091 := Call(__e, tmp65088, tmp65090)

				tmp65092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4flatten)

				tmp65093 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65094 := Call(__e, tmp65093, V283)

				tmp65095 := Call(__e, tmp65092, tmp65094)

				__e.TailApply(tmp65087, tmp65091, tmp65095)
				return

			} else {
				tmp65086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				__e.TailApply(tmp65086, V283, Nil)
				return

			}

		}

	}, 1)

	tmp65100 := Call(__e, PrimNS1Value(symns2_1set), symshen_4flatten, tmp65083)

	_ = tmp65100

	tmp65101 := MakeNative(func(__e *ControlFlow) {
		V287 := __e.Get(1)
		_ = V287
		V288 := __e.Get(2)
		_ = V288
		V289 := __e.Get(3)
		_ = V289
		tmp65154 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65155 := Call(__e, tmp65154, Nil, V287)

		if True == tmp65155 {
			tmp65151 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65152 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65153 := Call(__e, tmp65152, V289, Nil)

			__e.TailApply(tmp65151, V288, tmp65153)
			return

		} else {
			tmp65149 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65150 := Call(__e, tmp65149, V287)

			if True == tmp65150 {
				tmp65145 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp65146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65147 := Call(__e, tmp65146, V287)

				tmp65148 := Call(__e, tmp65145, tmp65147)

				var ifres65137 Obj

				if True == tmp65148 {
					tmp65139 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					tmp65140 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65141 := Call(__e, tmp65140, V287)

					tmp65142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65143 := Call(__e, tmp65142, V287)

					tmp65144 := Call(__e, tmp65139, tmp65141, tmp65143)

					var ifres65138 Obj

					if True == tmp65144 {
						ifres65138 = True

					} else {
						ifres65138 = False

					}

					ifres65137 = ifres65138

				} else {
					ifres65137 = False

				}

				if True == ifres65137 {
					tmp65109 := MakeNative(func(__e *ControlFlow) {
						Var := __e.Get(1)
						_ = Var
						tmp65110 := MakeNative(func(__e *ControlFlow) {
							NewAction := __e.Get(1)
							_ = NewAction
							tmp65111 := MakeNative(func(__e *ControlFlow) {
								NewPatts := __e.Get(1)
								_ = NewPatts
								tmp65112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__help)

								tmp65113 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp65114 := Call(__e, tmp65113, V287)

								__e.TailApply(tmp65112, tmp65114, NewPatts, NewAction)
								return

							}, 1)

							tmp65115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__X)

							tmp65116 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp65117 := Call(__e, tmp65116, V287)

							tmp65118 := Call(__e, tmp65115, tmp65117, Var, V288)

							__e.TailApply(tmp65111, tmp65118)
							return

						}, 1)

						tmp65119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65123 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65124 := Call(__e, tmp65123, V287)

						tmp65125 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65126 := Call(__e, tmp65125, Var, Nil)

						tmp65127 := Call(__e, tmp65122, tmp65124, tmp65126)

						tmp65128 := Call(__e, tmp65121, sym_a, tmp65127)

						tmp65129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65130 := Call(__e, tmp65129, V289, Nil)

						tmp65131 := Call(__e, tmp65120, tmp65128, tmp65130)

						tmp65132 := Call(__e, tmp65119, symwhere, tmp65131)

						__e.TailApply(tmp65110, tmp65132)
						return

					}, 1)

					tmp65133 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

					tmp65134 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65135 := Call(__e, tmp65134, V287)

					tmp65136 := Call(__e, tmp65133, tmp65135)

					__e.TailApply(tmp65109, tmp65136)
					return

				} else {
					tmp65106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__help)

					tmp65107 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65108 := Call(__e, tmp65107, V287)

					__e.TailApply(tmp65106, tmp65108, V288, V289)
					return

				}

			} else {
				tmp65104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp65104, symshen_4linearise__help)
				return

			}

		}

	}, 3)

	tmp65156 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise__help, tmp65101)

	_ = tmp65156

	tmp65157 := MakeNative(func(__e *ControlFlow) {
		V302 := __e.Get(1)
		_ = V302
		V303 := __e.Get(2)
		_ = V303
		V304 := __e.Get(3)
		_ = V304
		tmp65182 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65183 := Call(__e, tmp65182, V304, V302)

		if True == tmp65183 {
			__e.Return(V303)
			return
		} else {
			tmp65180 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65181 := Call(__e, tmp65180, V304)

			if True == tmp65181 {
				tmp65160 := MakeNative(func(__e *ControlFlow) {
					L := __e.Get(1)
					_ = L
					tmp65172 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp65173 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65174 := Call(__e, tmp65173, V304)

					tmp65175 := Call(__e, tmp65172, L, tmp65174)

					if True == tmp65175 {
						tmp65165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65166 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65167 := Call(__e, tmp65166, V304)

						tmp65168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__X)

						tmp65169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65170 := Call(__e, tmp65169, V304)

						tmp65171 := Call(__e, tmp65168, V302, V303, tmp65170)

						__e.TailApply(tmp65165, tmp65167, tmp65171)
						return

					} else {
						tmp65162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65163 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65164 := Call(__e, tmp65163, V304)

						__e.TailApply(tmp65162, L, tmp65164)
						return

					}

				}, 1)

				tmp65176 := Call(__e, PrimNS1Value(symns2_1value), symshen_4linearise__X)

				tmp65177 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65178 := Call(__e, tmp65177, V304)

				tmp65179 := Call(__e, tmp65176, V302, V303, tmp65178)

				__e.TailApply(tmp65160, tmp65179)
				return

			} else {
				__e.Return(V304)
				return
			}

		}

	}, 3)

	tmp65184 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise__X, tmp65157)

	_ = tmp65184

	tmp65185 := MakeNative(func(__e *ControlFlow) {
		V307 := __e.Get(1)
		_ = V307
		V308 := __e.Get(2)
		_ = V308
		tmp65331 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65332 := Call(__e, tmp65331, V308)

		var ifres65301 Obj

		if True == tmp65332 {
			tmp65327 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65329 := Call(__e, tmp65328, V308)

			tmp65330 := Call(__e, tmp65327, tmp65329)

			var ifres65303 Obj

			if True == tmp65330 {
				tmp65321 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp65322 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65324 := Call(__e, tmp65323, V308)

				tmp65325 := Call(__e, tmp65322, tmp65324)

				tmp65326 := Call(__e, tmp65321, tmp65325)

				var ifres65305 Obj

				if True == tmp65326 {
					tmp65313 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp65314 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65315 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65316 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65317 := Call(__e, tmp65316, V308)

					tmp65318 := Call(__e, tmp65315, tmp65317)

					tmp65319 := Call(__e, tmp65314, tmp65318)

					tmp65320 := Call(__e, tmp65313, Nil, tmp65319)

					var ifres65307 Obj

					if True == tmp65320 {
						tmp65309 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp65310 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65311 := Call(__e, tmp65310, V308)

						tmp65312 := Call(__e, tmp65309, Nil, tmp65311)

						var ifres65308 Obj

						if True == tmp65312 {
							ifres65308 = True

						} else {
							ifres65308 = False

						}

						ifres65307 = ifres65308

					} else {
						ifres65307 = False

					}

					var ifres65306 Obj

					if True == ifres65307 {
						ifres65306 = True

					} else {
						ifres65306 = False

					}

					ifres65305 = ifres65306

				} else {
					ifres65305 = False

				}

				var ifres65304 Obj

				if True == ifres65305 {
					ifres65304 = True

				} else {
					ifres65304 = False

				}

				ifres65303 = ifres65304

			} else {
				ifres65303 = False

			}

			var ifres65302 Obj

			if True == ifres65303 {
				ifres65302 = True

			} else {
				ifres65302 = False

			}

			ifres65301 = ifres65302

		} else {
			ifres65301 = False

		}

		if True == ifres65301 {
			tmp65284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck_1action)

			tmp65285 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65286 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65287 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65288 := Call(__e, tmp65287, V308)

			tmp65289 := Call(__e, tmp65286, tmp65288)

			tmp65290 := Call(__e, tmp65285, tmp65289)

			tmp65291 := Call(__e, tmp65284, tmp65290)

			_ = tmp65291

			tmp65292 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck_1name)

			tmp65293 := Call(__e, PrimNS1Value(symns2_1value), symarity)

			tmp65294 := Call(__e, tmp65293, V307)

			tmp65295 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp65296 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65298 := Call(__e, tmp65297, V308)

			tmp65299 := Call(__e, tmp65296, tmp65298)

			tmp65300 := Call(__e, tmp65295, tmp65299)

			__e.TailApply(tmp65292, V307, tmp65294, tmp65300)
			return

		} else {
			tmp65282 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65283 := Call(__e, tmp65282, V308)

			var ifres65222 Obj

			if True == tmp65283 {
				tmp65278 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp65279 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65280 := Call(__e, tmp65279, V308)

				tmp65281 := Call(__e, tmp65278, tmp65280)

				var ifres65224 Obj

				if True == tmp65281 {
					tmp65272 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp65273 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65274 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65275 := Call(__e, tmp65274, V308)

					tmp65276 := Call(__e, tmp65273, tmp65275)

					tmp65277 := Call(__e, tmp65272, tmp65276)

					var ifres65226 Obj

					if True == tmp65277 {
						tmp65264 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp65265 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65266 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65267 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65268 := Call(__e, tmp65267, V308)

						tmp65269 := Call(__e, tmp65266, tmp65268)

						tmp65270 := Call(__e, tmp65265, tmp65269)

						tmp65271 := Call(__e, tmp65264, Nil, tmp65270)

						var ifres65228 Obj

						if True == tmp65271 {
							tmp65260 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp65261 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65262 := Call(__e, tmp65261, V308)

							tmp65263 := Call(__e, tmp65260, tmp65262)

							var ifres65230 Obj

							if True == tmp65263 {
								tmp65254 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp65255 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp65256 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp65257 := Call(__e, tmp65256, V308)

								tmp65258 := Call(__e, tmp65255, tmp65257)

								tmp65259 := Call(__e, tmp65254, tmp65258)

								var ifres65232 Obj

								if True == tmp65259 {
									tmp65246 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp65247 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65248 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65250 := Call(__e, tmp65249, V308)

									tmp65251 := Call(__e, tmp65248, tmp65250)

									tmp65252 := Call(__e, tmp65247, tmp65251)

									tmp65253 := Call(__e, tmp65246, tmp65252)

									var ifres65234 Obj

									if True == tmp65253 {
										tmp65236 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp65237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65238 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65239 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp65240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65241 := Call(__e, tmp65240, V308)

										tmp65242 := Call(__e, tmp65239, tmp65241)

										tmp65243 := Call(__e, tmp65238, tmp65242)

										tmp65244 := Call(__e, tmp65237, tmp65243)

										tmp65245 := Call(__e, tmp65236, Nil, tmp65244)

										var ifres65235 Obj

										if True == tmp65245 {
											ifres65235 = True

										} else {
											ifres65235 = False

										}

										ifres65234 = ifres65235

									} else {
										ifres65234 = False

									}

									var ifres65233 Obj

									if True == ifres65234 {
										ifres65233 = True

									} else {
										ifres65233 = False

									}

									ifres65232 = ifres65233

								} else {
									ifres65232 = False

								}

								var ifres65231 Obj

								if True == ifres65232 {
									ifres65231 = True

								} else {
									ifres65231 = False

								}

								ifres65230 = ifres65231

							} else {
								ifres65230 = False

							}

							var ifres65229 Obj

							if True == ifres65230 {
								ifres65229 = True

							} else {
								ifres65229 = False

							}

							ifres65228 = ifres65229

						} else {
							ifres65228 = False

						}

						var ifres65227 Obj

						if True == ifres65228 {
							ifres65227 = True

						} else {
							ifres65227 = False

						}

						ifres65226 = ifres65227

					} else {
						ifres65226 = False

					}

					var ifres65225 Obj

					if True == ifres65226 {
						ifres65225 = True

					} else {
						ifres65225 = False

					}

					ifres65224 = ifres65225

				} else {
					ifres65224 = False

				}

				var ifres65223 Obj

				if True == ifres65224 {
					ifres65223 = True

				} else {
					ifres65223 = False

				}

				ifres65222 = ifres65223

			} else {
				ifres65222 = False

			}

			if True == ifres65222 {
				tmp65206 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65207 := Call(__e, PrimNS1Value(symns2_1value), symlength)

				tmp65208 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65209 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65210 := Call(__e, tmp65209, V308)

				tmp65211 := Call(__e, tmp65208, tmp65210)

				tmp65212 := Call(__e, tmp65207, tmp65211)

				tmp65213 := Call(__e, PrimNS1Value(symns2_1value), symlength)

				tmp65214 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65215 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65216 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65217 := Call(__e, tmp65216, V308)

				tmp65218 := Call(__e, tmp65215, tmp65217)

				tmp65219 := Call(__e, tmp65214, tmp65218)

				tmp65220 := Call(__e, tmp65213, tmp65219)

				tmp65221 := Call(__e, tmp65206, tmp65212, tmp65220)

				if True == tmp65221 {
					tmp65195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck_1action)

					tmp65196 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65197 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65198 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65199 := Call(__e, tmp65198, V308)

					tmp65200 := Call(__e, tmp65197, tmp65199)

					tmp65201 := Call(__e, tmp65196, tmp65200)

					tmp65202 := Call(__e, tmp65195, tmp65201)

					_ = tmp65202

					tmp65203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck)

					tmp65204 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65205 := Call(__e, tmp65204, V308)

					__e.TailApply(tmp65203, V307, tmp65205)
					return

				} else {
					tmp65190 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					tmp65191 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp65192 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp65193 := Call(__e, tmp65192, V307, MakeString("\n"), symshen_4a)

					tmp65194 := Call(__e, tmp65191, MakeString("arity error in "), tmp65193)

					__e.TailApply(tmp65190, tmp65194)
					return

				}

			} else {
				tmp65188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp65188, symshen_4aritycheck)
				return

			}

		}

	}, 2)

	tmp65333 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck, tmp65185)

	_ = tmp65333

	tmp65334 := MakeNative(func(__e *ControlFlow) {
		V321 := __e.Get(1)
		_ = V321
		V322 := __e.Get(2)
		_ = V322
		V323 := __e.Get(3)
		_ = V323
		tmp65347 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65348 := Call(__e, tmp65347, MakeNumber(-1), V322)

		if True == tmp65348 {
			__e.Return(V323)
			return
		} else {
			tmp65345 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp65346 := Call(__e, tmp65345, V323, V322)

			if True == tmp65346 {
				__e.Return(V323)
				return
			} else {
				tmp65337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp65338 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp65339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp65340 := Call(__e, tmp65339, V321, MakeString(" can cause errors.\n"), symshen_4a)

				tmp65341 := Call(__e, tmp65338, MakeString("\nwarning: changing the arity of "), tmp65340)

				tmp65342 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp65343 := Call(__e, tmp65342)

				tmp65344 := Call(__e, tmp65337, tmp65341, tmp65343)

				_ = tmp65344

				__e.Return(V323)
				return

			}

		}

	}, 3)

	tmp65349 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck_1name, tmp65334)

	_ = tmp65349

	tmp65350 := MakeNative(func(__e *ControlFlow) {
		V329 := __e.Get(1)
		_ = V329
		tmp65361 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65362 := Call(__e, tmp65361, V329)

		if True == tmp65362 {
			tmp65352 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aah)

			tmp65353 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65354 := Call(__e, tmp65353, V329)

			tmp65355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65356 := Call(__e, tmp65355, V329)

			tmp65357 := Call(__e, tmp65352, tmp65354, tmp65356)

			_ = tmp65357

			tmp65358 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

			tmp65359 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp65360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4aritycheck_1action)

				__e.TailApply(tmp65360, Y)
				return

			}, 1)

			__e.TailApply(tmp65358, tmp65359, V329)
			return

		} else {
			__e.Return(symshen_4skip)
			return
		}

	}, 1)

	tmp65363 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck_1action, tmp65350)

	_ = tmp65363

	tmp65364 := MakeNative(func(__e *ControlFlow) {
		V332 := __e.Get(1)
		_ = V332
		V333 := __e.Get(2)
		_ = V333
		tmp65365 := MakeNative(func(__e *ControlFlow) {
			Arity := __e.Get(1)
			_ = Arity
			tmp65366 := MakeNative(func(__e *ControlFlow) {
				Len := __e.Get(1)
				_ = Len
				tmp65390 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

				tmp65391 := Call(__e, tmp65390, Arity, MakeNumber(-1))

				var ifres65386 Obj

				if True == tmp65391 {
					tmp65388 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp65389 := Call(__e, tmp65388, Len, Arity)

					var ifres65387 Obj

					if True == tmp65389 {
						ifres65387 = True

					} else {
						ifres65387 = False

					}

					ifres65386 = ifres65387

				} else {
					ifres65386 = False

				}

				if True == ifres65386 {
					tmp65368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					tmp65369 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp65370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp65371 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp65372 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp65373 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp65374 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp65376 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp65377 := Call(__e, tmp65376, Len, MakeNumber(1))

					var ifres65375 Obj

					if True == tmp65377 {
						ifres65375 = MakeString("s")

					} else {
						ifres65375 = MakeString("")

					}

					tmp65378 := Call(__e, tmp65374, ifres65375, MakeString(".\n"), symshen_4a)

					tmp65379 := Call(__e, tmp65373, MakeString(" argument"), tmp65378)

					tmp65380 := Call(__e, tmp65372, Len, tmp65379, symshen_4a)

					tmp65381 := Call(__e, tmp65371, MakeString(" might not like "), tmp65380)

					tmp65382 := Call(__e, tmp65370, V332, tmp65381, symshen_4a)

					tmp65383 := Call(__e, tmp65369, MakeString("warning: "), tmp65382)

					tmp65384 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					tmp65385 := Call(__e, tmp65384)

					__e.TailApply(tmp65368, tmp65383, tmp65385)
					return

				} else {
					__e.Return(symshen_4skip)
					return
				}

			}, 1)

			tmp65392 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp65393 := Call(__e, tmp65392, V333)

			__e.TailApply(tmp65366, tmp65393)
			return

		}, 1)

		tmp65394 := Call(__e, PrimNS1Value(symns2_1value), symarity)

		tmp65395 := Call(__e, tmp65394, V332)

		__e.TailApply(tmp65365, tmp65395)
		return

	}, 2)

	tmp65396 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aah, tmp65364)

	_ = tmp65396

	tmp65397 := MakeNative(func(__e *ControlFlow) {
		V335 := __e.Get(1)
		_ = V335
		tmp65421 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65422 := Call(__e, tmp65421, V335)

		var ifres65407 Obj

		if True == tmp65422 {
			tmp65417 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65418 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65419 := Call(__e, tmp65418, V335)

			tmp65420 := Call(__e, tmp65417, tmp65419)

			var ifres65409 Obj

			if True == tmp65420 {
				tmp65411 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65413 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65414 := Call(__e, tmp65413, V335)

				tmp65415 := Call(__e, tmp65412, tmp65414)

				tmp65416 := Call(__e, tmp65411, Nil, tmp65415)

				var ifres65410 Obj

				if True == tmp65416 {
					ifres65410 = True

				} else {
					ifres65410 = False

				}

				ifres65409 = ifres65410

			} else {
				ifres65409 = False

			}

			var ifres65408 Obj

			if True == ifres65409 {
				ifres65408 = True

			} else {
				ifres65408 = False

			}

			ifres65407 = ifres65408

		} else {
			ifres65407 = False

		}

		if True == ifres65407 {
			tmp65400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abstraction__build)

			tmp65401 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65402 := Call(__e, tmp65401, V335)

			tmp65403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65404 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65405 := Call(__e, tmp65404, V335)

			tmp65406 := Call(__e, tmp65403, tmp65405)

			__e.TailApply(tmp65400, tmp65402, tmp65406)
			return

		} else {
			tmp65399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp65399, symshen_4abstract__rule)
			return

		}

	}, 1)

	tmp65423 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abstract__rule, tmp65397)

	_ = tmp65423

	tmp65424 := MakeNative(func(__e *ControlFlow) {
		V338 := __e.Get(1)
		_ = V338
		V339 := __e.Get(2)
		_ = V339
		tmp65441 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65442 := Call(__e, tmp65441, Nil, V338)

		if True == tmp65442 {
			__e.Return(V339)
			return
		} else {
			tmp65439 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65440 := Call(__e, tmp65439, V338)

			if True == tmp65440 {
				tmp65428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65430 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65431 := Call(__e, tmp65430, V338)

				tmp65432 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abstraction__build)

				tmp65434 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65435 := Call(__e, tmp65434, V338)

				tmp65436 := Call(__e, tmp65433, tmp65435, V339)

				tmp65437 := Call(__e, tmp65432, tmp65436, Nil)

				tmp65438 := Call(__e, tmp65429, tmp65431, tmp65437)

				__e.TailApply(tmp65428, sym_c_4, tmp65438)
				return

			} else {
				tmp65427 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp65427, symshen_4abstraction__build)
				return

			}

		}

	}, 2)

	tmp65443 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abstraction__build, tmp65424)

	_ = tmp65443

	tmp65444 := MakeNative(func(__e *ControlFlow) {
		V341 := __e.Get(1)
		_ = V341
		tmp65453 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65454 := Call(__e, tmp65453, MakeNumber(0), V341)

		if True == tmp65454 {
			__e.Return(Nil)
			return
		} else {
			tmp65446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65447 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp65448 := Call(__e, tmp65447, symV)

			tmp65449 := Call(__e, PrimNS1Value(symns2_1value), symshen_4parameters)

			tmp65450 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp65451 := Call(__e, tmp65450, V341, MakeNumber(1))

			tmp65452 := Call(__e, tmp65449, tmp65451)

			__e.TailApply(tmp65446, tmp65448, tmp65452)
			return

		}

	}, 1)

	tmp65455 := Call(__e, PrimNS1Value(symns2_1set), symshen_4parameters, tmp65444)

	_ = tmp65455

	tmp65456 := MakeNative(func(__e *ControlFlow) {
		V344 := __e.Get(1)
		_ = V344
		V345 := __e.Get(2)
		_ = V345
		tmp65471 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp65472 := Call(__e, tmp65471, Nil, V344)

		if True == tmp65472 {
			__e.Return(V345)
			return
		} else {
			tmp65469 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65470 := Call(__e, tmp65469, V344)

			if True == tmp65470 {
				tmp65460 := Call(__e, PrimNS1Value(symns2_1value), symshen_4application__build)

				tmp65461 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65462 := Call(__e, tmp65461, V344)

				tmp65463 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65464 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65465 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65466 := Call(__e, tmp65465, V344)

				tmp65467 := Call(__e, tmp65464, tmp65466, Nil)

				tmp65468 := Call(__e, tmp65463, V345, tmp65467)

				__e.TailApply(tmp65460, tmp65462, tmp65468)
				return

			} else {
				tmp65459 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp65459, symshen_4application__build)
				return

			}

		}

	}, 2)

	tmp65473 := Call(__e, PrimNS1Value(symns2_1set), symshen_4application__build, tmp65456)

	_ = tmp65473

	tmp65474 := MakeNative(func(__e *ControlFlow) {
		V348 := __e.Get(1)
		_ = V348
		V349 := __e.Get(2)
		_ = V349
		tmp65539 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65540 := Call(__e, tmp65539, V349)

		var ifres65525 Obj

		if True == tmp65540 {
			tmp65535 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65536 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65537 := Call(__e, tmp65536, V349)

			tmp65538 := Call(__e, tmp65535, tmp65537)

			var ifres65527 Obj

			if True == tmp65538 {
				tmp65529 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65530 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65531 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65532 := Call(__e, tmp65531, V349)

				tmp65533 := Call(__e, tmp65530, tmp65532)

				tmp65534 := Call(__e, tmp65529, Nil, tmp65533)

				var ifres65528 Obj

				if True == tmp65534 {
					ifres65528 = True

				} else {
					ifres65528 = False

				}

				ifres65527 = ifres65528

			} else {
				ifres65527 = False

			}

			var ifres65526 Obj

			if True == ifres65527 {
				ifres65526 = True

			} else {
				ifres65526 = False

			}

			ifres65525 = ifres65526

		} else {
			ifres65525 = False

		}

		if True == ifres65525 {
			tmp65477 := MakeNative(func(__e *ControlFlow) {
				Arity := __e.Get(1)
				_ = Arity
				tmp65478 := MakeNative(func(__e *ControlFlow) {
					Reduce := __e.Get(1)
					_ = Reduce
					tmp65479 := MakeNative(func(__e *ControlFlow) {
						CondExpression := __e.Get(1)
						_ = CondExpression
						tmp65480 := MakeNative(func(__e *ControlFlow) {
							TypeTable := __e.Get(1)
							_ = TypeTable
							tmp65481 := MakeNative(func(__e *ControlFlow) {
								TypedCondExpression := __e.Get(1)
								_ = TypedCondExpression
								tmp65482 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65483 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65484 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65485 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp65486 := Call(__e, tmp65485, V349)

								tmp65487 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65488 := Call(__e, tmp65487, TypedCondExpression, Nil)

								tmp65489 := Call(__e, tmp65484, tmp65486, tmp65488)

								tmp65490 := Call(__e, tmp65483, V348, tmp65489)

								__e.TailApply(tmp65482, symdefun, tmp65490)
								return

							}, 1)

							tmp65496 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

							tmp65497 := Call(__e, tmp65496, symshen_4_doptimise_d)

							var ifres65491 Obj

							if True == tmp65497 {
								tmp65492 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

								tmp65493 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp65494 := Call(__e, tmp65493, V349)

								tmp65495 := Call(__e, tmp65492, tmp65494, TypeTable, CondExpression)

								ifres65491 = tmp65495

							} else {
								ifres65491 = CondExpression

							}

							__e.TailApply(tmp65481, ifres65491)
							return

						}, 1)

						tmp65505 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

						tmp65506 := Call(__e, tmp65505, symshen_4_doptimise_d)

						var ifres65498 Obj

						if True == tmp65506 {
							tmp65499 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typextable)

							tmp65500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4get_1type)

							tmp65501 := Call(__e, tmp65500, V348)

							tmp65502 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp65503 := Call(__e, tmp65502, V349)

							tmp65504 := Call(__e, tmp65499, tmp65501, tmp65503)

							ifres65498 = tmp65504

						} else {
							ifres65498 = symshen_4skip

						}

						__e.TailApply(tmp65480, ifres65498)
						return

					}, 1)

					tmp65507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cond_1expression)

					tmp65508 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65509 := Call(__e, tmp65508, V349)

					tmp65510 := Call(__e, tmp65507, V348, tmp65509, Reduce)

					__e.TailApply(tmp65479, tmp65510)
					return

				}, 1)

				tmp65511 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp65512 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp65513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce)

					__e.TailApply(tmp65513, X)
					return

				}, 1)

				tmp65514 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65515 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65516 := Call(__e, tmp65515, V349)

				tmp65517 := Call(__e, tmp65514, tmp65516)

				tmp65518 := Call(__e, tmp65511, tmp65512, tmp65517)

				__e.TailApply(tmp65478, tmp65518)
				return

			}, 1)

			tmp65519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4store_1arity)

			tmp65520 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp65521 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65522 := Call(__e, tmp65521, V349)

			tmp65523 := Call(__e, tmp65520, tmp65522)

			tmp65524 := Call(__e, tmp65519, V348, tmp65523)

			__e.TailApply(tmp65477, tmp65524)
			return

		} else {
			tmp65476 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp65476, symshen_4compile__to__kl)
			return

		}

	}, 2)

	tmp65541 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__kl, tmp65474)

	_ = tmp65541

	tmp65542 := MakeNative(func(__e *ControlFlow) {
		V355 := __e.Get(1)
		_ = V355
		tmp65553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65554 := Call(__e, tmp65553, V355)

		if True == tmp65554 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp65544 := MakeNative(func(__e *ControlFlow) {
				FType := __e.Get(1)
				_ = FType
				tmp65547 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				tmp65548 := Call(__e, tmp65547, FType)

				if True == tmp65548 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp65546 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					__e.TailApply(tmp65546, FType)
					return

				}

			}, 1)

			tmp65549 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

			tmp65550 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp65551 := Call(__e, tmp65550, symshen_4_dsignedfuncs_d)

			tmp65552 := Call(__e, tmp65549, V355, tmp65551)

			__e.TailApply(tmp65544, tmp65552)
			return

		}

	}, 1)

	tmp65555 := Call(__e, PrimNS1Value(symns2_1set), symshen_4get_1type, tmp65542)

	_ = tmp65555

	tmp65556 := MakeNative(func(__e *ControlFlow) {
		V366 := __e.Get(1)
		_ = V366
		V367 := __e.Get(2)
		_ = V367
		tmp65625 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65626 := Call(__e, tmp65625, V366)

		var ifres65589 Obj

		if True == tmp65626 {
			tmp65621 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65622 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65623 := Call(__e, tmp65622, V366)

			tmp65624 := Call(__e, tmp65621, tmp65623)

			var ifres65591 Obj

			if True == tmp65624 {
				tmp65615 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65616 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65617 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65618 := Call(__e, tmp65617, V366)

				tmp65619 := Call(__e, tmp65616, tmp65618)

				tmp65620 := Call(__e, tmp65615, sym_1_1_6, tmp65619)

				var ifres65593 Obj

				if True == tmp65620 {
					tmp65609 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp65610 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65612 := Call(__e, tmp65611, V366)

					tmp65613 := Call(__e, tmp65610, tmp65612)

					tmp65614 := Call(__e, tmp65609, tmp65613)

					var ifres65595 Obj

					if True == tmp65614 {
						tmp65601 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp65602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65603 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65604 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65605 := Call(__e, tmp65604, V366)

						tmp65606 := Call(__e, tmp65603, tmp65605)

						tmp65607 := Call(__e, tmp65602, tmp65606)

						tmp65608 := Call(__e, tmp65601, Nil, tmp65607)

						var ifres65597 Obj

						if True == tmp65608 {
							tmp65599 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp65600 := Call(__e, tmp65599, V367)

							var ifres65598 Obj

							if True == tmp65600 {
								ifres65598 = True

							} else {
								ifres65598 = False

							}

							ifres65597 = ifres65598

						} else {
							ifres65597 = False

						}

						var ifres65596 Obj

						if True == ifres65597 {
							ifres65596 = True

						} else {
							ifres65596 = False

						}

						ifres65595 = ifres65596

					} else {
						ifres65595 = False

					}

					var ifres65594 Obj

					if True == ifres65595 {
						ifres65594 = True

					} else {
						ifres65594 = False

					}

					ifres65593 = ifres65594

				} else {
					ifres65593 = False

				}

				var ifres65592 Obj

				if True == ifres65593 {
					ifres65592 = True

				} else {
					ifres65592 = False

				}

				ifres65591 = ifres65592

			} else {
				ifres65591 = False

			}

			var ifres65590 Obj

			if True == ifres65591 {
				ifres65590 = True

			} else {
				ifres65590 = False

			}

			ifres65589 = ifres65590

		} else {
			ifres65589 = False

		}

		if True == ifres65589 {
			tmp65585 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp65586 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65587 := Call(__e, tmp65586, V366)

			tmp65588 := Call(__e, tmp65585, tmp65587)

			if True == tmp65588 {
				tmp65576 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typextable)

				tmp65577 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65578 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65579 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65580 := Call(__e, tmp65579, V366)

				tmp65581 := Call(__e, tmp65578, tmp65580)

				tmp65582 := Call(__e, tmp65577, tmp65581)

				tmp65583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65584 := Call(__e, tmp65583, V367)

				__e.TailApply(tmp65576, tmp65582, tmp65584)
				return

			} else {
				tmp65559 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65561 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65562 := Call(__e, tmp65561, V367)

				tmp65563 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65564 := Call(__e, tmp65563, V366)

				tmp65565 := Call(__e, tmp65560, tmp65562, tmp65564)

				tmp65566 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typextable)

				tmp65567 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65570 := Call(__e, tmp65569, V366)

				tmp65571 := Call(__e, tmp65568, tmp65570)

				tmp65572 := Call(__e, tmp65567, tmp65571)

				tmp65573 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65574 := Call(__e, tmp65573, V367)

				tmp65575 := Call(__e, tmp65566, tmp65572, tmp65574)

				__e.TailApply(tmp65559, tmp65565, tmp65575)
				return

			}

		} else {
			__e.Return(Nil)
			return
		}

	}, 2)

	tmp65627 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typextable, tmp65556)

	_ = tmp65627

	tmp65628 := MakeNative(func(__e *ControlFlow) {
		V371 := __e.Get(1)
		_ = V371
		V372 := __e.Get(2)
		_ = V372
		V373 := __e.Get(3)
		_ = V373
		tmp65831 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp65832 := Call(__e, tmp65831, V373)

		var ifres65789 Obj

		if True == tmp65832 {
			tmp65827 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp65828 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65829 := Call(__e, tmp65828, V373)

			tmp65830 := Call(__e, tmp65827, symlet, tmp65829)

			var ifres65791 Obj

			if True == tmp65830 {
				tmp65823 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp65824 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65825 := Call(__e, tmp65824, V373)

				tmp65826 := Call(__e, tmp65823, tmp65825)

				var ifres65793 Obj

				if True == tmp65826 {
					tmp65817 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp65818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65819 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65820 := Call(__e, tmp65819, V373)

					tmp65821 := Call(__e, tmp65818, tmp65820)

					tmp65822 := Call(__e, tmp65817, tmp65821)

					var ifres65795 Obj

					if True == tmp65822 {
						tmp65809 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp65810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65811 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65812 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65813 := Call(__e, tmp65812, V373)

						tmp65814 := Call(__e, tmp65811, tmp65813)

						tmp65815 := Call(__e, tmp65810, tmp65814)

						tmp65816 := Call(__e, tmp65809, tmp65815)

						var ifres65797 Obj

						if True == tmp65816 {
							tmp65799 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp65800 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65802 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65804 := Call(__e, tmp65803, V373)

							tmp65805 := Call(__e, tmp65802, tmp65804)

							tmp65806 := Call(__e, tmp65801, tmp65805)

							tmp65807 := Call(__e, tmp65800, tmp65806)

							tmp65808 := Call(__e, tmp65799, Nil, tmp65807)

							var ifres65798 Obj

							if True == tmp65808 {
								ifres65798 = True

							} else {
								ifres65798 = False

							}

							ifres65797 = ifres65798

						} else {
							ifres65797 = False

						}

						var ifres65796 Obj

						if True == ifres65797 {
							ifres65796 = True

						} else {
							ifres65796 = False

						}

						ifres65795 = ifres65796

					} else {
						ifres65795 = False

					}

					var ifres65794 Obj

					if True == ifres65795 {
						ifres65794 = True

					} else {
						ifres65794 = False

					}

					ifres65793 = ifres65794

				} else {
					ifres65793 = False

				}

				var ifres65792 Obj

				if True == ifres65793 {
					ifres65792 = True

				} else {
					ifres65792 = False

				}

				ifres65791 = ifres65792

			} else {
				ifres65791 = False

			}

			var ifres65790 Obj

			if True == ifres65791 {
				ifres65790 = True

			} else {
				ifres65790 = False

			}

			ifres65789 = ifres65790

		} else {
			ifres65789 = False

		}

		if True == ifres65789 {
			tmp65754 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65756 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65757 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65758 := Call(__e, tmp65757, V373)

			tmp65759 := Call(__e, tmp65756, tmp65758)

			tmp65760 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

			tmp65762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65765 := Call(__e, tmp65764, V373)

			tmp65766 := Call(__e, tmp65763, tmp65765)

			tmp65767 := Call(__e, tmp65762, tmp65766)

			tmp65768 := Call(__e, tmp65761, V371, V372, tmp65767)

			tmp65769 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65770 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

			tmp65771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65772 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65773 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65774 := Call(__e, tmp65773, V373)

			tmp65775 := Call(__e, tmp65772, tmp65774)

			tmp65776 := Call(__e, tmp65771, tmp65775, V371)

			tmp65777 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp65778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65779 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp65781 := Call(__e, tmp65780, V373)

			tmp65782 := Call(__e, tmp65779, tmp65781)

			tmp65783 := Call(__e, tmp65778, tmp65782)

			tmp65784 := Call(__e, tmp65777, tmp65783)

			tmp65785 := Call(__e, tmp65770, tmp65776, V372, tmp65784)

			tmp65786 := Call(__e, tmp65769, tmp65785, Nil)

			tmp65787 := Call(__e, tmp65760, tmp65768, tmp65786)

			tmp65788 := Call(__e, tmp65755, tmp65759, tmp65787)

			__e.TailApply(tmp65754, symlet, tmp65788)
			return

		} else {
			tmp65752 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp65753 := Call(__e, tmp65752, V373)

			var ifres65722 Obj

			if True == tmp65753 {
				tmp65748 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp65749 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65750 := Call(__e, tmp65749, V373)

				tmp65751 := Call(__e, tmp65748, symlambda, tmp65750)

				var ifres65724 Obj

				if True == tmp65751 {
					tmp65744 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp65745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65746 := Call(__e, tmp65745, V373)

					tmp65747 := Call(__e, tmp65744, tmp65746)

					var ifres65726 Obj

					if True == tmp65747 {
						tmp65738 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp65739 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65740 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65741 := Call(__e, tmp65740, V373)

						tmp65742 := Call(__e, tmp65739, tmp65741)

						tmp65743 := Call(__e, tmp65738, tmp65742)

						var ifres65728 Obj

						if True == tmp65743 {
							tmp65730 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp65731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65732 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65733 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65734 := Call(__e, tmp65733, V373)

							tmp65735 := Call(__e, tmp65732, tmp65734)

							tmp65736 := Call(__e, tmp65731, tmp65735)

							tmp65737 := Call(__e, tmp65730, Nil, tmp65736)

							var ifres65729 Obj

							if True == tmp65737 {
								ifres65729 = True

							} else {
								ifres65729 = False

							}

							ifres65728 = ifres65729

						} else {
							ifres65728 = False

						}

						var ifres65727 Obj

						if True == ifres65728 {
							ifres65727 = True

						} else {
							ifres65727 = False

						}

						ifres65726 = ifres65727

					} else {
						ifres65726 = False

					}

					var ifres65725 Obj

					if True == ifres65726 {
						ifres65725 = True

					} else {
						ifres65725 = False

					}

					ifres65724 = ifres65725

				} else {
					ifres65724 = False

				}

				var ifres65723 Obj

				if True == ifres65724 {
					ifres65723 = True

				} else {
					ifres65723 = False

				}

				ifres65722 = ifres65723

			} else {
				ifres65722 = False

			}

			if True == ifres65722 {
				tmp65699 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65700 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65701 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65703 := Call(__e, tmp65702, V373)

				tmp65704 := Call(__e, tmp65701, tmp65703)

				tmp65705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65706 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

				tmp65707 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65708 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65709 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65710 := Call(__e, tmp65709, V373)

				tmp65711 := Call(__e, tmp65708, tmp65710)

				tmp65712 := Call(__e, tmp65707, tmp65711, V371)

				tmp65713 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp65714 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp65716 := Call(__e, tmp65715, V373)

				tmp65717 := Call(__e, tmp65714, tmp65716)

				tmp65718 := Call(__e, tmp65713, tmp65717)

				tmp65719 := Call(__e, tmp65706, tmp65712, V372, tmp65718)

				tmp65720 := Call(__e, tmp65705, tmp65719, Nil)

				tmp65721 := Call(__e, tmp65700, tmp65704, tmp65720)

				__e.TailApply(tmp65699, symlambda, tmp65721)
				return

			} else {
				tmp65697 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp65698 := Call(__e, tmp65697, V373)

				var ifres65691 Obj

				if True == tmp65698 {
					tmp65693 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp65694 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp65695 := Call(__e, tmp65694, V373)

					tmp65696 := Call(__e, tmp65693, symcond, tmp65695)

					var ifres65692 Obj

					if True == tmp65696 {
						ifres65692 = True

					} else {
						ifres65692 = False

					}

					ifres65691 = ifres65692

				} else {
					ifres65691 = False

				}

				if True == ifres65691 {
					tmp65672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp65673 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					tmp65674 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						tmp65675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65676 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

						tmp65677 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65678 := Call(__e, tmp65677, Y)

						tmp65679 := Call(__e, tmp65676, V371, V372, tmp65678)

						tmp65680 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

						tmp65682 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65684 := Call(__e, tmp65683, Y)

						tmp65685 := Call(__e, tmp65682, tmp65684)

						tmp65686 := Call(__e, tmp65681, V371, V372, tmp65685)

						tmp65687 := Call(__e, tmp65680, tmp65686, Nil)

						__e.TailApply(tmp65675, tmp65679, tmp65687)
						return

					}, 1)

					tmp65688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp65689 := Call(__e, tmp65688, V373)

					tmp65690 := Call(__e, tmp65673, tmp65674, tmp65689)

					__e.TailApply(tmp65672, symcond, tmp65690)
					return

				} else {
					tmp65670 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp65671 := Call(__e, tmp65670, V373)

					if True == tmp65671 {
						tmp65650 := MakeNative(func(__e *ControlFlow) {
							NewTable := __e.Get(1)
							_ = NewTable
							tmp65651 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp65652 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp65653 := Call(__e, tmp65652, V373)

							tmp65654 := Call(__e, PrimNS1Value(symns2_1value), symmap)

							tmp65655 := MakeNative(func(__e *ControlFlow) {
								Y := __e.Get(1)
								_ = Y
								tmp65656 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assign_1types)

								tmp65657 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								tmp65658 := Call(__e, tmp65657, V372, NewTable)

								__e.TailApply(tmp65656, V371, tmp65658, Y)
								return

							}, 1)

							tmp65659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp65660 := Call(__e, tmp65659, V373)

							tmp65661 := Call(__e, tmp65654, tmp65655, tmp65660)

							__e.TailApply(tmp65651, tmp65653, tmp65661)
							return

						}, 1)

						tmp65662 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typextable)

						tmp65663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4get_1type)

						tmp65664 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp65665 := Call(__e, tmp65664, V373)

						tmp65666 := Call(__e, tmp65663, tmp65665)

						tmp65667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp65668 := Call(__e, tmp65667, V373)

						tmp65669 := Call(__e, tmp65662, tmp65666, tmp65668)

						__e.TailApply(tmp65650, tmp65669)
						return

					} else {
						tmp65633 := MakeNative(func(__e *ControlFlow) {
							AtomType := __e.Get(1)
							_ = AtomType
							tmp65646 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp65647 := Call(__e, tmp65646, AtomType)

							if True == tmp65647 {
								tmp65639 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65640 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65641 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp65642 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp65643 := Call(__e, tmp65642, AtomType)

								tmp65644 := Call(__e, tmp65641, tmp65643, Nil)

								tmp65645 := Call(__e, tmp65640, V373, tmp65644)

								__e.TailApply(tmp65639, symtype, tmp65645)
								return

							} else {
								tmp65637 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

								tmp65638 := Call(__e, tmp65637, V373, V371)

								if True == tmp65638 {
									__e.Return(V373)
									return
								} else {
									tmp65636 := Call(__e, PrimNS1Value(symns2_1value), symshen_4atom_1type)

									__e.TailApply(tmp65636, V373)
									return

								}

							}

						}, 1)

						tmp65648 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

						tmp65649 := Call(__e, tmp65648, V373, V372)

						__e.TailApply(tmp65633, tmp65649)
						return

					}

				}

			}

		}

	}, 3)

	tmp65833 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assign_1types, tmp65628)

	_ = tmp65833

	tmp65834 := MakeNative(func(__e *ControlFlow) {
		V375 := __e.Get(1)
		_ = V375
		tmp65865 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

		tmp65866 := Call(__e, tmp65865, V375)

		if True == tmp65866 {
			tmp65860 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65863 := Call(__e, tmp65862, symstring, Nil)

			tmp65864 := Call(__e, tmp65861, V375, tmp65863)

			__e.TailApply(tmp65860, symtype, tmp65864)
			return

		} else {
			tmp65858 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

			tmp65859 := Call(__e, tmp65858, V375)

			if True == tmp65859 {
				tmp65853 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65854 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65855 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp65856 := Call(__e, tmp65855, symnumber, Nil)

				tmp65857 := Call(__e, tmp65854, V375, tmp65856)

				__e.TailApply(tmp65853, symtype, tmp65857)
				return

			} else {
				tmp65851 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

				tmp65852 := Call(__e, tmp65851, V375)

				if True == tmp65852 {
					tmp65846 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp65847 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp65848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp65849 := Call(__e, tmp65848, symboolean, Nil)

					tmp65850 := Call(__e, tmp65847, V375, tmp65849)

					__e.TailApply(tmp65846, symtype, tmp65850)
					return

				} else {
					tmp65844 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

					tmp65845 := Call(__e, tmp65844, V375)

					if True == tmp65845 {
						tmp65839 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65840 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65841 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp65842 := Call(__e, tmp65841, symsymbol, Nil)

						tmp65843 := Call(__e, tmp65840, V375, tmp65842)

						__e.TailApply(tmp65839, symtype, tmp65843)
						return

					} else {
						__e.Return(V375)
						return
					}

				}

			}

		}

	}, 1)

	tmp65867 := Call(__e, PrimNS1Value(symns2_1set), symshen_4atom_1type, tmp65834)

	_ = tmp65867

	tmp65868 := MakeNative(func(__e *ControlFlow) {
		V380 := __e.Get(1)
		_ = V380
		V381 := __e.Get(2)
		_ = V381
		tmp65873 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp65874 := Call(__e, tmp65873, symshen_4_dinstalling_1kl_d)

		if True == tmp65874 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp65870 := Call(__e, PrimNS1Value(symns2_1value), symput)

			tmp65871 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp65872 := Call(__e, tmp65871, sym_dproperty_1vector_d)

			__e.TailApply(tmp65870, V380, symarity, V381, tmp65872)
			return

		}

	}, 2)

	tmp65875 := Call(__e, PrimNS1Value(symns2_1set), symshen_4store_1arity, tmp65868)

	_ = tmp65875

	tmp65876 := MakeNative(func(__e *ControlFlow) {
		V383 := __e.Get(1)
		_ = V383
		tmp65877 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp65878 := Call(__e, tmp65877, symshen_4_dteststack_d, Nil)

		_ = tmp65878

		tmp65879 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp65880 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65881 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65882 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65883 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			tmp65884 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp65885 := Call(__e, tmp65884, symshen_4_dteststack_d)

			tmp65886 := Call(__e, tmp65883, tmp65885)

			tmp65887 := Call(__e, tmp65882, symshen_4tests, tmp65886)

			tmp65888 := Call(__e, tmp65881, sym_h, tmp65887)

			tmp65889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp65890 := Call(__e, tmp65889, Result, Nil)

			__e.TailApply(tmp65880, tmp65888, tmp65890)
			return

		}, 1)

		tmp65891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

		tmp65892 := Call(__e, tmp65891, V383)

		__e.TailApply(tmp65879, tmp65892)
		return

	}, 1)

	tmp65893 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reduce, tmp65876)

	_ = tmp65893

	tmp65894 := MakeNative(func(__e *ControlFlow) {
		V385 := __e.Get(1)
		_ = V385
		tmp67043 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp67044 := Call(__e, tmp67043, V385)

		var ifres66921 Obj

		if True == tmp67044 {
			tmp67039 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp67040 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp67041 := Call(__e, tmp67040, V385)

			tmp67042 := Call(__e, tmp67039, tmp67041)

			var ifres66923 Obj

			if True == tmp67042 {
				tmp67033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp67034 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67035 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67036 := Call(__e, tmp67035, V385)

				tmp67037 := Call(__e, tmp67034, tmp67036)

				tmp67038 := Call(__e, tmp67033, sym_c_4, tmp67037)

				var ifres66925 Obj

				if True == tmp67038 {
					tmp67027 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67029 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67030 := Call(__e, tmp67029, V385)

					tmp67031 := Call(__e, tmp67028, tmp67030)

					tmp67032 := Call(__e, tmp67027, tmp67031)

					var ifres66927 Obj

					if True == tmp67032 {
						tmp67019 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67020 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67021 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67022 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67023 := Call(__e, tmp67022, V385)

						tmp67024 := Call(__e, tmp67021, tmp67023)

						tmp67025 := Call(__e, tmp67020, tmp67024)

						tmp67026 := Call(__e, tmp67019, tmp67025)

						var ifres66929 Obj

						if True == tmp67026 {
							tmp67009 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp67010 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67011 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67012 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67013 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67014 := Call(__e, tmp67013, V385)

							tmp67015 := Call(__e, tmp67012, tmp67014)

							tmp67016 := Call(__e, tmp67011, tmp67015)

							tmp67017 := Call(__e, tmp67010, tmp67016)

							tmp67018 := Call(__e, tmp67009, symcons, tmp67017)

							var ifres66931 Obj

							if True == tmp67018 {
								tmp66999 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67000 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67001 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67002 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67003 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67004 := Call(__e, tmp67003, V385)

								tmp67005 := Call(__e, tmp67002, tmp67004)

								tmp67006 := Call(__e, tmp67001, tmp67005)

								tmp67007 := Call(__e, tmp67000, tmp67006)

								tmp67008 := Call(__e, tmp66999, tmp67007)

								var ifres66933 Obj

								if True == tmp67008 {
									tmp66987 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp66988 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66989 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66990 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66991 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66993 := Call(__e, tmp66992, V385)

									tmp66994 := Call(__e, tmp66991, tmp66993)

									tmp66995 := Call(__e, tmp66990, tmp66994)

									tmp66996 := Call(__e, tmp66989, tmp66995)

									tmp66997 := Call(__e, tmp66988, tmp66996)

									tmp66998 := Call(__e, tmp66987, tmp66997)

									var ifres66935 Obj

									if True == tmp66998 {
										tmp66973 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp66974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66975 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66976 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66977 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66979 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66980 := Call(__e, tmp66979, V385)

										tmp66981 := Call(__e, tmp66978, tmp66980)

										tmp66982 := Call(__e, tmp66977, tmp66981)

										tmp66983 := Call(__e, tmp66976, tmp66982)

										tmp66984 := Call(__e, tmp66975, tmp66983)

										tmp66985 := Call(__e, tmp66974, tmp66984)

										tmp66986 := Call(__e, tmp66973, Nil, tmp66985)

										var ifres66937 Obj

										if True == tmp66986 {
											tmp66965 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66967 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66968 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66969 := Call(__e, tmp66968, V385)

											tmp66970 := Call(__e, tmp66967, tmp66969)

											tmp66971 := Call(__e, tmp66966, tmp66970)

											tmp66972 := Call(__e, tmp66965, tmp66971)

											var ifres66939 Obj

											if True == tmp66972 {
												tmp66955 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp66956 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66959 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66960 := Call(__e, tmp66959, V385)

												tmp66961 := Call(__e, tmp66958, tmp66960)

												tmp66962 := Call(__e, tmp66957, tmp66961)

												tmp66963 := Call(__e, tmp66956, tmp66962)

												tmp66964 := Call(__e, tmp66955, Nil, tmp66963)

												var ifres66941 Obj

												if True == tmp66964 {
													tmp66951 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp66952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66953 := Call(__e, tmp66952, V385)

													tmp66954 := Call(__e, tmp66951, tmp66953)

													var ifres66943 Obj

													if True == tmp66954 {
														tmp66945 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp66946 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66947 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66948 := Call(__e, tmp66947, V385)

														tmp66949 := Call(__e, tmp66946, tmp66948)

														tmp66950 := Call(__e, tmp66945, Nil, tmp66949)

														var ifres66944 Obj

														if True == tmp66950 {
															ifres66944 = True

														} else {
															ifres66944 = False

														}

														ifres66943 = ifres66944

													} else {
														ifres66943 = False

													}

													var ifres66942 Obj

													if True == ifres66943 {
														ifres66942 = True

													} else {
														ifres66942 = False

													}

													ifres66941 = ifres66942

												} else {
													ifres66941 = False

												}

												var ifres66940 Obj

												if True == ifres66941 {
													ifres66940 = True

												} else {
													ifres66940 = False

												}

												ifres66939 = ifres66940

											} else {
												ifres66939 = False

											}

											var ifres66938 Obj

											if True == ifres66939 {
												ifres66938 = True

											} else {
												ifres66938 = False

											}

											ifres66937 = ifres66938

										} else {
											ifres66937 = False

										}

										var ifres66936 Obj

										if True == ifres66937 {
											ifres66936 = True

										} else {
											ifres66936 = False

										}

										ifres66935 = ifres66936

									} else {
										ifres66935 = False

									}

									var ifres66934 Obj

									if True == ifres66935 {
										ifres66934 = True

									} else {
										ifres66934 = False

									}

									ifres66933 = ifres66934

								} else {
									ifres66933 = False

								}

								var ifres66932 Obj

								if True == ifres66933 {
									ifres66932 = True

								} else {
									ifres66932 = False

								}

								ifres66931 = ifres66932

							} else {
								ifres66931 = False

							}

							var ifres66930 Obj

							if True == ifres66931 {
								ifres66930 = True

							} else {
								ifres66930 = False

							}

							ifres66929 = ifres66930

						} else {
							ifres66929 = False

						}

						var ifres66928 Obj

						if True == ifres66929 {
							ifres66928 = True

						} else {
							ifres66928 = False

						}

						ifres66927 = ifres66928

					} else {
						ifres66927 = False

					}

					var ifres66926 Obj

					if True == ifres66927 {
						ifres66926 = True

					} else {
						ifres66926 = False

					}

					ifres66925 = ifres66926

				} else {
					ifres66925 = False

				}

				var ifres66924 Obj

				if True == ifres66925 {
					ifres66924 = True

				} else {
					ifres66924 = False

				}

				ifres66923 = ifres66924

			} else {
				ifres66923 = False

			}

			var ifres66922 Obj

			if True == ifres66923 {
				ifres66922 = True

			} else {
				ifres66922 = False

			}

			ifres66921 = ifres66922

		} else {
			ifres66921 = False

		}

		if True == ifres66921 {
			tmp66842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

			tmp66843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66845 := Call(__e, tmp66844, V385)

			tmp66846 := Call(__e, tmp66843, symcons_2, tmp66845)

			tmp66847 := Call(__e, tmp66842, tmp66846)

			_ = tmp66847

			tmp66848 := MakeNative(func(__e *ControlFlow) {
				Abstraction := __e.Get(1)
				_ = Abstraction
				tmp66849 := MakeNative(func(__e *ControlFlow) {
					Application := __e.Get(1)
					_ = Application
					tmp66850 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

					__e.TailApply(tmp66850, Application)
					return

				}, 1)

				tmp66851 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66852 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66853 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66854 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66856 := Call(__e, tmp66855, V385)

				tmp66857 := Call(__e, tmp66854, symhd, tmp66856)

				tmp66858 := Call(__e, tmp66853, tmp66857, Nil)

				tmp66859 := Call(__e, tmp66852, Abstraction, tmp66858)

				tmp66860 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66863 := Call(__e, tmp66862, V385)

				tmp66864 := Call(__e, tmp66861, symtl, tmp66863)

				tmp66865 := Call(__e, tmp66860, tmp66864, Nil)

				tmp66866 := Call(__e, tmp66851, tmp66859, tmp66865)

				__e.TailApply(tmp66849, tmp66866)
				return

			}, 1)

			tmp66867 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66868 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66869 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66870 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66871 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66873 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66874 := Call(__e, tmp66873, V385)

			tmp66875 := Call(__e, tmp66872, tmp66874)

			tmp66876 := Call(__e, tmp66871, tmp66875)

			tmp66877 := Call(__e, tmp66870, tmp66876)

			tmp66878 := Call(__e, tmp66869, tmp66877)

			tmp66879 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66880 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66881 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66882 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66885 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66887 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66888 := Call(__e, tmp66887, V385)

			tmp66889 := Call(__e, tmp66886, tmp66888)

			tmp66890 := Call(__e, tmp66885, tmp66889)

			tmp66891 := Call(__e, tmp66884, tmp66890)

			tmp66892 := Call(__e, tmp66883, tmp66891)

			tmp66893 := Call(__e, tmp66882, tmp66892)

			tmp66894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp66895 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

			tmp66896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66898 := Call(__e, tmp66897, V385)

			tmp66899 := Call(__e, tmp66896, tmp66898)

			tmp66900 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66901 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66902 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66903 := Call(__e, tmp66902, V385)

			tmp66904 := Call(__e, tmp66901, tmp66903)

			tmp66905 := Call(__e, tmp66900, tmp66904)

			tmp66906 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66907 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66908 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp66909 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp66910 := Call(__e, tmp66909, V385)

			tmp66911 := Call(__e, tmp66908, tmp66910)

			tmp66912 := Call(__e, tmp66907, tmp66911)

			tmp66913 := Call(__e, tmp66906, tmp66912)

			tmp66914 := Call(__e, tmp66895, tmp66899, tmp66905, tmp66913)

			tmp66915 := Call(__e, tmp66894, tmp66914, Nil)

			tmp66916 := Call(__e, tmp66881, tmp66893, tmp66915)

			tmp66917 := Call(__e, tmp66880, sym_c_4, tmp66916)

			tmp66918 := Call(__e, tmp66879, tmp66917, Nil)

			tmp66919 := Call(__e, tmp66868, tmp66878, tmp66918)

			tmp66920 := Call(__e, tmp66867, sym_c_4, tmp66919)

			__e.TailApply(tmp66848, tmp66920)
			return

		} else {
			tmp66840 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp66841 := Call(__e, tmp66840, V385)

			var ifres66718 Obj

			if True == tmp66841 {
				tmp66836 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp66837 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66838 := Call(__e, tmp66837, V385)

				tmp66839 := Call(__e, tmp66836, tmp66838)

				var ifres66720 Obj

				if True == tmp66839 {
					tmp66830 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp66831 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66832 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66833 := Call(__e, tmp66832, V385)

					tmp66834 := Call(__e, tmp66831, tmp66833)

					tmp66835 := Call(__e, tmp66830, sym_c_4, tmp66834)

					var ifres66722 Obj

					if True == tmp66835 {
						tmp66824 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp66825 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66826 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66827 := Call(__e, tmp66826, V385)

						tmp66828 := Call(__e, tmp66825, tmp66827)

						tmp66829 := Call(__e, tmp66824, tmp66828)

						var ifres66724 Obj

						if True == tmp66829 {
							tmp66816 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp66817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp66819 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66820 := Call(__e, tmp66819, V385)

							tmp66821 := Call(__e, tmp66818, tmp66820)

							tmp66822 := Call(__e, tmp66817, tmp66821)

							tmp66823 := Call(__e, tmp66816, tmp66822)

							var ifres66726 Obj

							if True == tmp66823 {
								tmp66806 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp66807 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66808 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66809 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66810 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66811 := Call(__e, tmp66810, V385)

								tmp66812 := Call(__e, tmp66809, tmp66811)

								tmp66813 := Call(__e, tmp66808, tmp66812)

								tmp66814 := Call(__e, tmp66807, tmp66813)

								tmp66815 := Call(__e, tmp66806, sym_8p, tmp66814)

								var ifres66728 Obj

								if True == tmp66815 {
									tmp66796 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp66797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66798 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66799 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66800 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66801 := Call(__e, tmp66800, V385)

									tmp66802 := Call(__e, tmp66799, tmp66801)

									tmp66803 := Call(__e, tmp66798, tmp66802)

									tmp66804 := Call(__e, tmp66797, tmp66803)

									tmp66805 := Call(__e, tmp66796, tmp66804)

									var ifres66730 Obj

									if True == tmp66805 {
										tmp66784 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp66785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66786 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66787 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66788 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66789 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66790 := Call(__e, tmp66789, V385)

										tmp66791 := Call(__e, tmp66788, tmp66790)

										tmp66792 := Call(__e, tmp66787, tmp66791)

										tmp66793 := Call(__e, tmp66786, tmp66792)

										tmp66794 := Call(__e, tmp66785, tmp66793)

										tmp66795 := Call(__e, tmp66784, tmp66794)

										var ifres66732 Obj

										if True == tmp66795 {
											tmp66770 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp66771 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66772 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66773 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66774 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66775 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66777 := Call(__e, tmp66776, V385)

											tmp66778 := Call(__e, tmp66775, tmp66777)

											tmp66779 := Call(__e, tmp66774, tmp66778)

											tmp66780 := Call(__e, tmp66773, tmp66779)

											tmp66781 := Call(__e, tmp66772, tmp66780)

											tmp66782 := Call(__e, tmp66771, tmp66781)

											tmp66783 := Call(__e, tmp66770, Nil, tmp66782)

											var ifres66734 Obj

											if True == tmp66783 {
												tmp66762 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp66763 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66766 := Call(__e, tmp66765, V385)

												tmp66767 := Call(__e, tmp66764, tmp66766)

												tmp66768 := Call(__e, tmp66763, tmp66767)

												tmp66769 := Call(__e, tmp66762, tmp66768)

												var ifres66736 Obj

												if True == tmp66769 {
													tmp66752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp66753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66754 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66756 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp66757 := Call(__e, tmp66756, V385)

													tmp66758 := Call(__e, tmp66755, tmp66757)

													tmp66759 := Call(__e, tmp66754, tmp66758)

													tmp66760 := Call(__e, tmp66753, tmp66759)

													tmp66761 := Call(__e, tmp66752, Nil, tmp66760)

													var ifres66738 Obj

													if True == tmp66761 {
														tmp66748 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp66749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66750 := Call(__e, tmp66749, V385)

														tmp66751 := Call(__e, tmp66748, tmp66750)

														var ifres66740 Obj

														if True == tmp66751 {
															tmp66742 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp66743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66744 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66745 := Call(__e, tmp66744, V385)

															tmp66746 := Call(__e, tmp66743, tmp66745)

															tmp66747 := Call(__e, tmp66742, Nil, tmp66746)

															var ifres66741 Obj

															if True == tmp66747 {
																ifres66741 = True

															} else {
																ifres66741 = False

															}

															ifres66740 = ifres66741

														} else {
															ifres66740 = False

														}

														var ifres66739 Obj

														if True == ifres66740 {
															ifres66739 = True

														} else {
															ifres66739 = False

														}

														ifres66738 = ifres66739

													} else {
														ifres66738 = False

													}

													var ifres66737 Obj

													if True == ifres66738 {
														ifres66737 = True

													} else {
														ifres66737 = False

													}

													ifres66736 = ifres66737

												} else {
													ifres66736 = False

												}

												var ifres66735 Obj

												if True == ifres66736 {
													ifres66735 = True

												} else {
													ifres66735 = False

												}

												ifres66734 = ifres66735

											} else {
												ifres66734 = False

											}

											var ifres66733 Obj

											if True == ifres66734 {
												ifres66733 = True

											} else {
												ifres66733 = False

											}

											ifres66732 = ifres66733

										} else {
											ifres66732 = False

										}

										var ifres66731 Obj

										if True == ifres66732 {
											ifres66731 = True

										} else {
											ifres66731 = False

										}

										ifres66730 = ifres66731

									} else {
										ifres66730 = False

									}

									var ifres66729 Obj

									if True == ifres66730 {
										ifres66729 = True

									} else {
										ifres66729 = False

									}

									ifres66728 = ifres66729

								} else {
									ifres66728 = False

								}

								var ifres66727 Obj

								if True == ifres66728 {
									ifres66727 = True

								} else {
									ifres66727 = False

								}

								ifres66726 = ifres66727

							} else {
								ifres66726 = False

							}

							var ifres66725 Obj

							if True == ifres66726 {
								ifres66725 = True

							} else {
								ifres66725 = False

							}

							ifres66724 = ifres66725

						} else {
							ifres66724 = False

						}

						var ifres66723 Obj

						if True == ifres66724 {
							ifres66723 = True

						} else {
							ifres66723 = False

						}

						ifres66722 = ifres66723

					} else {
						ifres66722 = False

					}

					var ifres66721 Obj

					if True == ifres66722 {
						ifres66721 = True

					} else {
						ifres66721 = False

					}

					ifres66720 = ifres66721

				} else {
					ifres66720 = False

				}

				var ifres66719 Obj

				if True == ifres66720 {
					ifres66719 = True

				} else {
					ifres66719 = False

				}

				ifres66718 = ifres66719

			} else {
				ifres66718 = False

			}

			if True == ifres66718 {
				tmp66639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

				tmp66640 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66642 := Call(__e, tmp66641, V385)

				tmp66643 := Call(__e, tmp66640, symtuple_2, tmp66642)

				tmp66644 := Call(__e, tmp66639, tmp66643)

				_ = tmp66644

				tmp66645 := MakeNative(func(__e *ControlFlow) {
					Abstraction := __e.Get(1)
					_ = Abstraction
					tmp66646 := MakeNative(func(__e *ControlFlow) {
						Application := __e.Get(1)
						_ = Application
						tmp66647 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

						__e.TailApply(tmp66647, Application)
						return

					}, 1)

					tmp66648 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66649 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66650 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66651 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66652 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66653 := Call(__e, tmp66652, V385)

					tmp66654 := Call(__e, tmp66651, symfst, tmp66653)

					tmp66655 := Call(__e, tmp66650, tmp66654, Nil)

					tmp66656 := Call(__e, tmp66649, Abstraction, tmp66655)

					tmp66657 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66658 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66660 := Call(__e, tmp66659, V385)

					tmp66661 := Call(__e, tmp66658, symsnd, tmp66660)

					tmp66662 := Call(__e, tmp66657, tmp66661, Nil)

					tmp66663 := Call(__e, tmp66648, tmp66656, tmp66662)

					__e.TailApply(tmp66646, tmp66663)
					return

				}, 1)

				tmp66664 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66666 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66668 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66669 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66670 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66671 := Call(__e, tmp66670, V385)

				tmp66672 := Call(__e, tmp66669, tmp66671)

				tmp66673 := Call(__e, tmp66668, tmp66672)

				tmp66674 := Call(__e, tmp66667, tmp66673)

				tmp66675 := Call(__e, tmp66666, tmp66674)

				tmp66676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66681 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66682 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66684 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66685 := Call(__e, tmp66684, V385)

				tmp66686 := Call(__e, tmp66683, tmp66685)

				tmp66687 := Call(__e, tmp66682, tmp66686)

				tmp66688 := Call(__e, tmp66681, tmp66687)

				tmp66689 := Call(__e, tmp66680, tmp66688)

				tmp66690 := Call(__e, tmp66679, tmp66689)

				tmp66691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp66692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

				tmp66693 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66694 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66695 := Call(__e, tmp66694, V385)

				tmp66696 := Call(__e, tmp66693, tmp66695)

				tmp66697 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66698 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66699 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66700 := Call(__e, tmp66699, V385)

				tmp66701 := Call(__e, tmp66698, tmp66700)

				tmp66702 := Call(__e, tmp66697, tmp66701)

				tmp66703 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66704 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66705 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp66706 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp66707 := Call(__e, tmp66706, V385)

				tmp66708 := Call(__e, tmp66705, tmp66707)

				tmp66709 := Call(__e, tmp66704, tmp66708)

				tmp66710 := Call(__e, tmp66703, tmp66709)

				tmp66711 := Call(__e, tmp66692, tmp66696, tmp66702, tmp66710)

				tmp66712 := Call(__e, tmp66691, tmp66711, Nil)

				tmp66713 := Call(__e, tmp66678, tmp66690, tmp66712)

				tmp66714 := Call(__e, tmp66677, sym_c_4, tmp66713)

				tmp66715 := Call(__e, tmp66676, tmp66714, Nil)

				tmp66716 := Call(__e, tmp66665, tmp66675, tmp66715)

				tmp66717 := Call(__e, tmp66664, sym_c_4, tmp66716)

				__e.TailApply(tmp66645, tmp66717)
				return

			} else {
				tmp66637 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp66638 := Call(__e, tmp66637, V385)

				var ifres66515 Obj

				if True == tmp66638 {
					tmp66633 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp66634 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66635 := Call(__e, tmp66634, V385)

					tmp66636 := Call(__e, tmp66633, tmp66635)

					var ifres66517 Obj

					if True == tmp66636 {
						tmp66627 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp66628 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66629 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66630 := Call(__e, tmp66629, V385)

						tmp66631 := Call(__e, tmp66628, tmp66630)

						tmp66632 := Call(__e, tmp66627, sym_c_4, tmp66631)

						var ifres66519 Obj

						if True == tmp66632 {
							tmp66621 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp66622 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp66623 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66624 := Call(__e, tmp66623, V385)

							tmp66625 := Call(__e, tmp66622, tmp66624)

							tmp66626 := Call(__e, tmp66621, tmp66625)

							var ifres66521 Obj

							if True == tmp66626 {
								tmp66613 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp66614 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66615 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66616 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66617 := Call(__e, tmp66616, V385)

								tmp66618 := Call(__e, tmp66615, tmp66617)

								tmp66619 := Call(__e, tmp66614, tmp66618)

								tmp66620 := Call(__e, tmp66613, tmp66619)

								var ifres66523 Obj

								if True == tmp66620 {
									tmp66603 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp66604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66605 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66606 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66607 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66608 := Call(__e, tmp66607, V385)

									tmp66609 := Call(__e, tmp66606, tmp66608)

									tmp66610 := Call(__e, tmp66605, tmp66609)

									tmp66611 := Call(__e, tmp66604, tmp66610)

									tmp66612 := Call(__e, tmp66603, sym_8v, tmp66611)

									var ifres66525 Obj

									if True == tmp66612 {
										tmp66593 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp66594 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66595 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66596 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66597 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66598 := Call(__e, tmp66597, V385)

										tmp66599 := Call(__e, tmp66596, tmp66598)

										tmp66600 := Call(__e, tmp66595, tmp66599)

										tmp66601 := Call(__e, tmp66594, tmp66600)

										tmp66602 := Call(__e, tmp66593, tmp66601)

										var ifres66527 Obj

										if True == tmp66602 {
											tmp66581 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66582 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66584 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66585 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66586 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66587 := Call(__e, tmp66586, V385)

											tmp66588 := Call(__e, tmp66585, tmp66587)

											tmp66589 := Call(__e, tmp66584, tmp66588)

											tmp66590 := Call(__e, tmp66583, tmp66589)

											tmp66591 := Call(__e, tmp66582, tmp66590)

											tmp66592 := Call(__e, tmp66581, tmp66591)

											var ifres66529 Obj

											if True == tmp66592 {
												tmp66567 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp66568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66570 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66571 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66572 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66573 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66574 := Call(__e, tmp66573, V385)

												tmp66575 := Call(__e, tmp66572, tmp66574)

												tmp66576 := Call(__e, tmp66571, tmp66575)

												tmp66577 := Call(__e, tmp66570, tmp66576)

												tmp66578 := Call(__e, tmp66569, tmp66577)

												tmp66579 := Call(__e, tmp66568, tmp66578)

												tmp66580 := Call(__e, tmp66567, Nil, tmp66579)

												var ifres66531 Obj

												if True == tmp66580 {
													tmp66559 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp66560 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66562 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp66563 := Call(__e, tmp66562, V385)

													tmp66564 := Call(__e, tmp66561, tmp66563)

													tmp66565 := Call(__e, tmp66560, tmp66564)

													tmp66566 := Call(__e, tmp66559, tmp66565)

													var ifres66533 Obj

													if True == tmp66566 {
														tmp66549 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp66550 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66551 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66552 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66553 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp66554 := Call(__e, tmp66553, V385)

														tmp66555 := Call(__e, tmp66552, tmp66554)

														tmp66556 := Call(__e, tmp66551, tmp66555)

														tmp66557 := Call(__e, tmp66550, tmp66556)

														tmp66558 := Call(__e, tmp66549, Nil, tmp66557)

														var ifres66535 Obj

														if True == tmp66558 {
															tmp66545 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp66546 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66547 := Call(__e, tmp66546, V385)

															tmp66548 := Call(__e, tmp66545, tmp66547)

															var ifres66537 Obj

															if True == tmp66548 {
																tmp66539 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp66540 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp66541 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp66542 := Call(__e, tmp66541, V385)

																tmp66543 := Call(__e, tmp66540, tmp66542)

																tmp66544 := Call(__e, tmp66539, Nil, tmp66543)

																var ifres66538 Obj

																if True == tmp66544 {
																	ifres66538 = True

																} else {
																	ifres66538 = False

																}

																ifres66537 = ifres66538

															} else {
																ifres66537 = False

															}

															var ifres66536 Obj

															if True == ifres66537 {
																ifres66536 = True

															} else {
																ifres66536 = False

															}

															ifres66535 = ifres66536

														} else {
															ifres66535 = False

														}

														var ifres66534 Obj

														if True == ifres66535 {
															ifres66534 = True

														} else {
															ifres66534 = False

														}

														ifres66533 = ifres66534

													} else {
														ifres66533 = False

													}

													var ifres66532 Obj

													if True == ifres66533 {
														ifres66532 = True

													} else {
														ifres66532 = False

													}

													ifres66531 = ifres66532

												} else {
													ifres66531 = False

												}

												var ifres66530 Obj

												if True == ifres66531 {
													ifres66530 = True

												} else {
													ifres66530 = False

												}

												ifres66529 = ifres66530

											} else {
												ifres66529 = False

											}

											var ifres66528 Obj

											if True == ifres66529 {
												ifres66528 = True

											} else {
												ifres66528 = False

											}

											ifres66527 = ifres66528

										} else {
											ifres66527 = False

										}

										var ifres66526 Obj

										if True == ifres66527 {
											ifres66526 = True

										} else {
											ifres66526 = False

										}

										ifres66525 = ifres66526

									} else {
										ifres66525 = False

									}

									var ifres66524 Obj

									if True == ifres66525 {
										ifres66524 = True

									} else {
										ifres66524 = False

									}

									ifres66523 = ifres66524

								} else {
									ifres66523 = False

								}

								var ifres66522 Obj

								if True == ifres66523 {
									ifres66522 = True

								} else {
									ifres66522 = False

								}

								ifres66521 = ifres66522

							} else {
								ifres66521 = False

							}

							var ifres66520 Obj

							if True == ifres66521 {
								ifres66520 = True

							} else {
								ifres66520 = False

							}

							ifres66519 = ifres66520

						} else {
							ifres66519 = False

						}

						var ifres66518 Obj

						if True == ifres66519 {
							ifres66518 = True

						} else {
							ifres66518 = False

						}

						ifres66517 = ifres66518

					} else {
						ifres66517 = False

					}

					var ifres66516 Obj

					if True == ifres66517 {
						ifres66516 = True

					} else {
						ifres66516 = False

					}

					ifres66515 = ifres66516

				} else {
					ifres66515 = False

				}

				if True == ifres66515 {
					tmp66436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

					tmp66437 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66438 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66439 := Call(__e, tmp66438, V385)

					tmp66440 := Call(__e, tmp66437, symshen_4_7vector_2, tmp66439)

					tmp66441 := Call(__e, tmp66436, tmp66440)

					_ = tmp66441

					tmp66442 := MakeNative(func(__e *ControlFlow) {
						Abstraction := __e.Get(1)
						_ = Abstraction
						tmp66443 := MakeNative(func(__e *ControlFlow) {
							Application := __e.Get(1)
							_ = Application
							tmp66444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

							__e.TailApply(tmp66444, Application)
							return

						}, 1)

						tmp66445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66448 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66449 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66450 := Call(__e, tmp66449, V385)

						tmp66451 := Call(__e, tmp66448, symhdv, tmp66450)

						tmp66452 := Call(__e, tmp66447, tmp66451, Nil)

						tmp66453 := Call(__e, tmp66446, Abstraction, tmp66452)

						tmp66454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66457 := Call(__e, tmp66456, V385)

						tmp66458 := Call(__e, tmp66455, symtlv, tmp66457)

						tmp66459 := Call(__e, tmp66454, tmp66458, Nil)

						tmp66460 := Call(__e, tmp66445, tmp66453, tmp66459)

						__e.TailApply(tmp66443, tmp66460)
						return

					}, 1)

					tmp66461 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66462 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66463 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66465 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66466 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66467 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66468 := Call(__e, tmp66467, V385)

					tmp66469 := Call(__e, tmp66466, tmp66468)

					tmp66470 := Call(__e, tmp66465, tmp66469)

					tmp66471 := Call(__e, tmp66464, tmp66470)

					tmp66472 := Call(__e, tmp66463, tmp66471)

					tmp66473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66475 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66476 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66478 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66479 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66481 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66482 := Call(__e, tmp66481, V385)

					tmp66483 := Call(__e, tmp66480, tmp66482)

					tmp66484 := Call(__e, tmp66479, tmp66483)

					tmp66485 := Call(__e, tmp66478, tmp66484)

					tmp66486 := Call(__e, tmp66477, tmp66485)

					tmp66487 := Call(__e, tmp66476, tmp66486)

					tmp66488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp66489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

					tmp66490 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66491 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66492 := Call(__e, tmp66491, V385)

					tmp66493 := Call(__e, tmp66490, tmp66492)

					tmp66494 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66495 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66496 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66497 := Call(__e, tmp66496, V385)

					tmp66498 := Call(__e, tmp66495, tmp66497)

					tmp66499 := Call(__e, tmp66494, tmp66498)

					tmp66500 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66501 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66502 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp66503 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp66504 := Call(__e, tmp66503, V385)

					tmp66505 := Call(__e, tmp66502, tmp66504)

					tmp66506 := Call(__e, tmp66501, tmp66505)

					tmp66507 := Call(__e, tmp66500, tmp66506)

					tmp66508 := Call(__e, tmp66489, tmp66493, tmp66499, tmp66507)

					tmp66509 := Call(__e, tmp66488, tmp66508, Nil)

					tmp66510 := Call(__e, tmp66475, tmp66487, tmp66509)

					tmp66511 := Call(__e, tmp66474, sym_c_4, tmp66510)

					tmp66512 := Call(__e, tmp66473, tmp66511, Nil)

					tmp66513 := Call(__e, tmp66462, tmp66472, tmp66512)

					tmp66514 := Call(__e, tmp66461, sym_c_4, tmp66513)

					__e.TailApply(tmp66442, tmp66514)
					return

				} else {
					tmp66434 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp66435 := Call(__e, tmp66434, V385)

					var ifres66312 Obj

					if True == tmp66435 {
						tmp66430 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp66431 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66432 := Call(__e, tmp66431, V385)

						tmp66433 := Call(__e, tmp66430, tmp66432)

						var ifres66314 Obj

						if True == tmp66433 {
							tmp66424 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp66425 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66426 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66427 := Call(__e, tmp66426, V385)

							tmp66428 := Call(__e, tmp66425, tmp66427)

							tmp66429 := Call(__e, tmp66424, sym_c_4, tmp66428)

							var ifres66316 Obj

							if True == tmp66429 {
								tmp66418 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp66419 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66420 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66421 := Call(__e, tmp66420, V385)

								tmp66422 := Call(__e, tmp66419, tmp66421)

								tmp66423 := Call(__e, tmp66418, tmp66422)

								var ifres66318 Obj

								if True == tmp66423 {
									tmp66410 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp66411 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66413 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66414 := Call(__e, tmp66413, V385)

									tmp66415 := Call(__e, tmp66412, tmp66414)

									tmp66416 := Call(__e, tmp66411, tmp66415)

									tmp66417 := Call(__e, tmp66410, tmp66416)

									var ifres66320 Obj

									if True == tmp66417 {
										tmp66400 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp66401 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66402 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66403 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66405 := Call(__e, tmp66404, V385)

										tmp66406 := Call(__e, tmp66403, tmp66405)

										tmp66407 := Call(__e, tmp66402, tmp66406)

										tmp66408 := Call(__e, tmp66401, tmp66407)

										tmp66409 := Call(__e, tmp66400, sym_8s, tmp66408)

										var ifres66322 Obj

										if True == tmp66409 {
											tmp66390 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66391 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66392 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66393 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66394 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66395 := Call(__e, tmp66394, V385)

											tmp66396 := Call(__e, tmp66393, tmp66395)

											tmp66397 := Call(__e, tmp66392, tmp66396)

											tmp66398 := Call(__e, tmp66391, tmp66397)

											tmp66399 := Call(__e, tmp66390, tmp66398)

											var ifres66324 Obj

											if True == tmp66399 {
												tmp66378 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp66379 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66380 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66381 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66382 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66383 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66384 := Call(__e, tmp66383, V385)

												tmp66385 := Call(__e, tmp66382, tmp66384)

												tmp66386 := Call(__e, tmp66381, tmp66385)

												tmp66387 := Call(__e, tmp66380, tmp66386)

												tmp66388 := Call(__e, tmp66379, tmp66387)

												tmp66389 := Call(__e, tmp66378, tmp66388)

												var ifres66326 Obj

												if True == tmp66389 {
													tmp66364 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp66365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66366 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66367 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66368 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp66369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66370 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp66371 := Call(__e, tmp66370, V385)

													tmp66372 := Call(__e, tmp66369, tmp66371)

													tmp66373 := Call(__e, tmp66368, tmp66372)

													tmp66374 := Call(__e, tmp66367, tmp66373)

													tmp66375 := Call(__e, tmp66366, tmp66374)

													tmp66376 := Call(__e, tmp66365, tmp66375)

													tmp66377 := Call(__e, tmp66364, Nil, tmp66376)

													var ifres66328 Obj

													if True == tmp66377 {
														tmp66356 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp66357 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66358 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66359 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp66360 := Call(__e, tmp66359, V385)

														tmp66361 := Call(__e, tmp66358, tmp66360)

														tmp66362 := Call(__e, tmp66357, tmp66361)

														tmp66363 := Call(__e, tmp66356, tmp66362)

														var ifres66330 Obj

														if True == tmp66363 {
															tmp66346 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp66347 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66348 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66349 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66350 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp66351 := Call(__e, tmp66350, V385)

															tmp66352 := Call(__e, tmp66349, tmp66351)

															tmp66353 := Call(__e, tmp66348, tmp66352)

															tmp66354 := Call(__e, tmp66347, tmp66353)

															tmp66355 := Call(__e, tmp66346, Nil, tmp66354)

															var ifres66332 Obj

															if True == tmp66355 {
																tmp66342 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp66343 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp66344 := Call(__e, tmp66343, V385)

																tmp66345 := Call(__e, tmp66342, tmp66344)

																var ifres66334 Obj

																if True == tmp66345 {
																	tmp66336 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp66337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp66338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp66339 := Call(__e, tmp66338, V385)

																	tmp66340 := Call(__e, tmp66337, tmp66339)

																	tmp66341 := Call(__e, tmp66336, Nil, tmp66340)

																	var ifres66335 Obj

																	if True == tmp66341 {
																		ifres66335 = True

																	} else {
																		ifres66335 = False

																	}

																	ifres66334 = ifres66335

																} else {
																	ifres66334 = False

																}

																var ifres66333 Obj

																if True == ifres66334 {
																	ifres66333 = True

																} else {
																	ifres66333 = False

																}

																ifres66332 = ifres66333

															} else {
																ifres66332 = False

															}

															var ifres66331 Obj

															if True == ifres66332 {
																ifres66331 = True

															} else {
																ifres66331 = False

															}

															ifres66330 = ifres66331

														} else {
															ifres66330 = False

														}

														var ifres66329 Obj

														if True == ifres66330 {
															ifres66329 = True

														} else {
															ifres66329 = False

														}

														ifres66328 = ifres66329

													} else {
														ifres66328 = False

													}

													var ifres66327 Obj

													if True == ifres66328 {
														ifres66327 = True

													} else {
														ifres66327 = False

													}

													ifres66326 = ifres66327

												} else {
													ifres66326 = False

												}

												var ifres66325 Obj

												if True == ifres66326 {
													ifres66325 = True

												} else {
													ifres66325 = False

												}

												ifres66324 = ifres66325

											} else {
												ifres66324 = False

											}

											var ifres66323 Obj

											if True == ifres66324 {
												ifres66323 = True

											} else {
												ifres66323 = False

											}

											ifres66322 = ifres66323

										} else {
											ifres66322 = False

										}

										var ifres66321 Obj

										if True == ifres66322 {
											ifres66321 = True

										} else {
											ifres66321 = False

										}

										ifres66320 = ifres66321

									} else {
										ifres66320 = False

									}

									var ifres66319 Obj

									if True == ifres66320 {
										ifres66319 = True

									} else {
										ifres66319 = False

									}

									ifres66318 = ifres66319

								} else {
									ifres66318 = False

								}

								var ifres66317 Obj

								if True == ifres66318 {
									ifres66317 = True

								} else {
									ifres66317 = False

								}

								ifres66316 = ifres66317

							} else {
								ifres66316 = False

							}

							var ifres66315 Obj

							if True == ifres66316 {
								ifres66315 = True

							} else {
								ifres66315 = False

							}

							ifres66314 = ifres66315

						} else {
							ifres66314 = False

						}

						var ifres66313 Obj

						if True == ifres66314 {
							ifres66313 = True

						} else {
							ifres66313 = False

						}

						ifres66312 = ifres66313

					} else {
						ifres66312 = False

					}

					if True == ifres66312 {
						tmp66227 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

						tmp66228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66229 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66230 := Call(__e, tmp66229, V385)

						tmp66231 := Call(__e, tmp66228, symshen_4_7string_2, tmp66230)

						tmp66232 := Call(__e, tmp66227, tmp66231)

						_ = tmp66232

						tmp66233 := MakeNative(func(__e *ControlFlow) {
							Abstraction := __e.Get(1)
							_ = Abstraction
							tmp66234 := MakeNative(func(__e *ControlFlow) {
								Application := __e.Get(1)
								_ = Application
								tmp66235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

								__e.TailApply(tmp66235, Application)
								return

							}, 1)

							tmp66236 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66237 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66238 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66239 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66240 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66241 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66242 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp66243 := Call(__e, tmp66242, V385)

							tmp66244 := Call(__e, tmp66241, tmp66243)

							tmp66245 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66246 := Call(__e, tmp66245, MakeNumber(0), Nil)

							tmp66247 := Call(__e, tmp66240, tmp66244, tmp66246)

							tmp66248 := Call(__e, tmp66239, sympos, tmp66247)

							tmp66249 := Call(__e, tmp66238, tmp66248, Nil)

							tmp66250 := Call(__e, tmp66237, Abstraction, tmp66249)

							tmp66251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66252 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp66253 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp66254 := Call(__e, tmp66253, V385)

							tmp66255 := Call(__e, tmp66252, symtlstr, tmp66254)

							tmp66256 := Call(__e, tmp66251, tmp66255, Nil)

							tmp66257 := Call(__e, tmp66236, tmp66250, tmp66256)

							__e.TailApply(tmp66234, tmp66257)
							return

						}, 1)

						tmp66258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66260 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66261 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66262 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66263 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66264 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66265 := Call(__e, tmp66264, V385)

						tmp66266 := Call(__e, tmp66263, tmp66265)

						tmp66267 := Call(__e, tmp66262, tmp66266)

						tmp66268 := Call(__e, tmp66261, tmp66267)

						tmp66269 := Call(__e, tmp66260, tmp66268)

						tmp66270 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66272 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66273 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66274 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66275 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66276 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66277 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66278 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66279 := Call(__e, tmp66278, V385)

						tmp66280 := Call(__e, tmp66277, tmp66279)

						tmp66281 := Call(__e, tmp66276, tmp66280)

						tmp66282 := Call(__e, tmp66275, tmp66281)

						tmp66283 := Call(__e, tmp66274, tmp66282)

						tmp66284 := Call(__e, tmp66273, tmp66283)

						tmp66285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp66286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

						tmp66287 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66289 := Call(__e, tmp66288, V385)

						tmp66290 := Call(__e, tmp66287, tmp66289)

						tmp66291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66292 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66293 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66294 := Call(__e, tmp66293, V385)

						tmp66295 := Call(__e, tmp66292, tmp66294)

						tmp66296 := Call(__e, tmp66291, tmp66295)

						tmp66297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66299 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp66300 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp66301 := Call(__e, tmp66300, V385)

						tmp66302 := Call(__e, tmp66299, tmp66301)

						tmp66303 := Call(__e, tmp66298, tmp66302)

						tmp66304 := Call(__e, tmp66297, tmp66303)

						tmp66305 := Call(__e, tmp66286, tmp66290, tmp66296, tmp66304)

						tmp66306 := Call(__e, tmp66285, tmp66305, Nil)

						tmp66307 := Call(__e, tmp66272, tmp66284, tmp66306)

						tmp66308 := Call(__e, tmp66271, sym_c_4, tmp66307)

						tmp66309 := Call(__e, tmp66270, tmp66308, Nil)

						tmp66310 := Call(__e, tmp66259, tmp66269, tmp66309)

						tmp66311 := Call(__e, tmp66258, sym_c_4, tmp66310)

						__e.TailApply(tmp66233, tmp66311)
						return

					} else {
						tmp66225 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp66226 := Call(__e, tmp66225, V385)

						var ifres66157 Obj

						if True == tmp66226 {
							tmp66221 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp66222 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp66223 := Call(__e, tmp66222, V385)

							tmp66224 := Call(__e, tmp66221, tmp66223)

							var ifres66159 Obj

							if True == tmp66224 {
								tmp66215 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp66216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66217 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66218 := Call(__e, tmp66217, V385)

								tmp66219 := Call(__e, tmp66216, tmp66218)

								tmp66220 := Call(__e, tmp66215, sym_c_4, tmp66219)

								var ifres66161 Obj

								if True == tmp66220 {
									tmp66209 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp66210 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp66211 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66212 := Call(__e, tmp66211, V385)

									tmp66213 := Call(__e, tmp66210, tmp66212)

									tmp66214 := Call(__e, tmp66209, tmp66213)

									var ifres66163 Obj

									if True == tmp66214 {
										tmp66201 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp66202 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66203 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66204 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66205 := Call(__e, tmp66204, V385)

										tmp66206 := Call(__e, tmp66203, tmp66205)

										tmp66207 := Call(__e, tmp66202, tmp66206)

										tmp66208 := Call(__e, tmp66201, tmp66207)

										var ifres66165 Obj

										if True == tmp66208 {
											tmp66193 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66194 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66195 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66196 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66197 := Call(__e, tmp66196, V385)

											tmp66198 := Call(__e, tmp66195, tmp66197)

											tmp66199 := Call(__e, tmp66194, tmp66198)

											tmp66200 := Call(__e, tmp66193, tmp66199)

											var ifres66167 Obj

											if True == tmp66200 {
												tmp66183 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp66184 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66187 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66188 := Call(__e, tmp66187, V385)

												tmp66189 := Call(__e, tmp66186, tmp66188)

												tmp66190 := Call(__e, tmp66185, tmp66189)

												tmp66191 := Call(__e, tmp66184, tmp66190)

												tmp66192 := Call(__e, tmp66183, Nil, tmp66191)

												var ifres66169 Obj

												if True == tmp66192 {
													tmp66179 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp66180 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66181 := Call(__e, tmp66180, V385)

													tmp66182 := Call(__e, tmp66179, tmp66181)

													var ifres66171 Obj

													if True == tmp66182 {
														tmp66173 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp66174 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66176 := Call(__e, tmp66175, V385)

														tmp66177 := Call(__e, tmp66174, tmp66176)

														tmp66178 := Call(__e, tmp66173, Nil, tmp66177)

														var ifres66172 Obj

														if True == tmp66178 {
															ifres66172 = True

														} else {
															ifres66172 = False

														}

														ifres66171 = ifres66172

													} else {
														ifres66171 = False

													}

													var ifres66170 Obj

													if True == ifres66171 {
														ifres66170 = True

													} else {
														ifres66170 = False

													}

													ifres66169 = ifres66170

												} else {
													ifres66169 = False

												}

												var ifres66168 Obj

												if True == ifres66169 {
													ifres66168 = True

												} else {
													ifres66168 = False

												}

												ifres66167 = ifres66168

											} else {
												ifres66167 = False

											}

											var ifres66166 Obj

											if True == ifres66167 {
												ifres66166 = True

											} else {
												ifres66166 = False

											}

											ifres66165 = ifres66166

										} else {
											ifres66165 = False

										}

										var ifres66164 Obj

										if True == ifres66165 {
											ifres66164 = True

										} else {
											ifres66164 = False

										}

										ifres66163 = ifres66164

									} else {
										ifres66163 = False

									}

									var ifres66162 Obj

									if True == ifres66163 {
										ifres66162 = True

									} else {
										ifres66162 = False

									}

									ifres66161 = ifres66162

								} else {
									ifres66161 = False

								}

								var ifres66160 Obj

								if True == ifres66161 {
									ifres66160 = True

								} else {
									ifres66160 = False

								}

								ifres66159 = ifres66160

							} else {
								ifres66159 = False

							}

							var ifres66158 Obj

							if True == ifres66159 {
								ifres66158 = True

							} else {
								ifres66158 = False

							}

							ifres66157 = ifres66158

						} else {
							ifres66157 = False

						}

						if True == ifres66157 {
							tmp66156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4custom_1pattern_1reducer)

							__e.TailApply(tmp66156, V385)
							return

						} else {
							tmp66154 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp66155 := Call(__e, tmp66154, V385)

							var ifres66084 Obj

							if True == tmp66155 {
								tmp66150 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp66151 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66152 := Call(__e, tmp66151, V385)

								tmp66153 := Call(__e, tmp66150, tmp66152)

								var ifres66086 Obj

								if True == tmp66153 {
									tmp66144 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp66145 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66147 := Call(__e, tmp66146, V385)

									tmp66148 := Call(__e, tmp66145, tmp66147)

									tmp66149 := Call(__e, tmp66144, sym_c_4, tmp66148)

									var ifres66088 Obj

									if True == tmp66149 {
										tmp66138 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp66139 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp66140 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66141 := Call(__e, tmp66140, V385)

										tmp66142 := Call(__e, tmp66139, tmp66141)

										tmp66143 := Call(__e, tmp66138, tmp66142)

										var ifres66090 Obj

										if True == tmp66143 {
											tmp66130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66132 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66133 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66134 := Call(__e, tmp66133, V385)

											tmp66135 := Call(__e, tmp66132, tmp66134)

											tmp66136 := Call(__e, tmp66131, tmp66135)

											tmp66137 := Call(__e, tmp66130, tmp66136)

											var ifres66092 Obj

											if True == tmp66137 {
												tmp66120 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp66121 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66124 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66125 := Call(__e, tmp66124, V385)

												tmp66126 := Call(__e, tmp66123, tmp66125)

												tmp66127 := Call(__e, tmp66122, tmp66126)

												tmp66128 := Call(__e, tmp66121, tmp66127)

												tmp66129 := Call(__e, tmp66120, Nil, tmp66128)

												var ifres66094 Obj

												if True == tmp66129 {
													tmp66116 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp66117 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66118 := Call(__e, tmp66117, V385)

													tmp66119 := Call(__e, tmp66116, tmp66118)

													var ifres66096 Obj

													if True == tmp66119 {
														tmp66110 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp66111 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66112 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66113 := Call(__e, tmp66112, V385)

														tmp66114 := Call(__e, tmp66111, tmp66113)

														tmp66115 := Call(__e, tmp66110, Nil, tmp66114)

														var ifres66098 Obj

														if True == tmp66115 {
															tmp66100 := Call(__e, PrimNS1Value(symns2_1value), symnot)

															tmp66101 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

															tmp66102 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp66103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66104 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp66105 := Call(__e, tmp66104, V385)

															tmp66106 := Call(__e, tmp66103, tmp66105)

															tmp66107 := Call(__e, tmp66102, tmp66106)

															tmp66108 := Call(__e, tmp66101, tmp66107)

															tmp66109 := Call(__e, tmp66100, tmp66108)

															var ifres66099 Obj

															if True == tmp66109 {
																ifres66099 = True

															} else {
																ifres66099 = False

															}

															ifres66098 = ifres66099

														} else {
															ifres66098 = False

														}

														var ifres66097 Obj

														if True == ifres66098 {
															ifres66097 = True

														} else {
															ifres66097 = False

														}

														ifres66096 = ifres66097

													} else {
														ifres66096 = False

													}

													var ifres66095 Obj

													if True == ifres66096 {
														ifres66095 = True

													} else {
														ifres66095 = False

													}

													ifres66094 = ifres66095

												} else {
													ifres66094 = False

												}

												var ifres66093 Obj

												if True == ifres66094 {
													ifres66093 = True

												} else {
													ifres66093 = False

												}

												ifres66092 = ifres66093

											} else {
												ifres66092 = False

											}

											var ifres66091 Obj

											if True == ifres66092 {
												ifres66091 = True

											} else {
												ifres66091 = False

											}

											ifres66090 = ifres66091

										} else {
											ifres66090 = False

										}

										var ifres66089 Obj

										if True == ifres66090 {
											ifres66089 = True

										} else {
											ifres66089 = False

										}

										ifres66088 = ifres66089

									} else {
										ifres66088 = False

									}

									var ifres66087 Obj

									if True == ifres66088 {
										ifres66087 = True

									} else {
										ifres66087 = False

									}

									ifres66086 = ifres66087

								} else {
									ifres66086 = False

								}

								var ifres66085 Obj

								if True == ifres66086 {
									ifres66085 = True

								} else {
									ifres66085 = False

								}

								ifres66084 = ifres66085

							} else {
								ifres66084 = False

							}

							if True == ifres66084 {
								tmp66061 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

								tmp66062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp66063 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp66064 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66066 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66067 := Call(__e, tmp66066, V385)

								tmp66068 := Call(__e, tmp66065, tmp66067)

								tmp66069 := Call(__e, tmp66064, tmp66068)

								tmp66070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66071 := Call(__e, tmp66070, V385)

								tmp66072 := Call(__e, tmp66063, tmp66069, tmp66071)

								tmp66073 := Call(__e, tmp66062, sym_a, tmp66072)

								tmp66074 := Call(__e, tmp66061, tmp66073)

								_ = tmp66074

								tmp66075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

								tmp66076 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66077 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp66079 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp66080 := Call(__e, tmp66079, V385)

								tmp66081 := Call(__e, tmp66078, tmp66080)

								tmp66082 := Call(__e, tmp66077, tmp66081)

								tmp66083 := Call(__e, tmp66076, tmp66082)

								__e.TailApply(tmp66075, tmp66083)
								return

							} else {
								tmp66059 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp66060 := Call(__e, tmp66059, V385)

								var ifres66001 Obj

								if True == tmp66060 {
									tmp66055 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp66056 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp66057 := Call(__e, tmp66056, V385)

									tmp66058 := Call(__e, tmp66055, tmp66057)

									var ifres66003 Obj

									if True == tmp66058 {
										tmp66049 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp66050 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66051 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp66052 := Call(__e, tmp66051, V385)

										tmp66053 := Call(__e, tmp66050, tmp66052)

										tmp66054 := Call(__e, tmp66049, sym_c_4, tmp66053)

										var ifres66005 Obj

										if True == tmp66054 {
											tmp66043 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp66044 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp66045 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp66046 := Call(__e, tmp66045, V385)

											tmp66047 := Call(__e, tmp66044, tmp66046)

											tmp66048 := Call(__e, tmp66043, tmp66047)

											var ifres66007 Obj

											if True == tmp66048 {
												tmp66035 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp66036 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp66038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp66039 := Call(__e, tmp66038, V385)

												tmp66040 := Call(__e, tmp66037, tmp66039)

												tmp66041 := Call(__e, tmp66036, tmp66040)

												tmp66042 := Call(__e, tmp66035, tmp66041)

												var ifres66009 Obj

												if True == tmp66042 {
													tmp66025 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp66026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66027 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp66029 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp66030 := Call(__e, tmp66029, V385)

													tmp66031 := Call(__e, tmp66028, tmp66030)

													tmp66032 := Call(__e, tmp66027, tmp66031)

													tmp66033 := Call(__e, tmp66026, tmp66032)

													tmp66034 := Call(__e, tmp66025, Nil, tmp66033)

													var ifres66011 Obj

													if True == tmp66034 {
														tmp66021 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp66022 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp66023 := Call(__e, tmp66022, V385)

														tmp66024 := Call(__e, tmp66021, tmp66023)

														var ifres66013 Obj

														if True == tmp66024 {
															tmp66015 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp66016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66017 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp66018 := Call(__e, tmp66017, V385)

															tmp66019 := Call(__e, tmp66016, tmp66018)

															tmp66020 := Call(__e, tmp66015, Nil, tmp66019)

															var ifres66014 Obj

															if True == tmp66020 {
																ifres66014 = True

															} else {
																ifres66014 = False

															}

															ifres66013 = ifres66014

														} else {
															ifres66013 = False

														}

														var ifres66012 Obj

														if True == ifres66013 {
															ifres66012 = True

														} else {
															ifres66012 = False

														}

														ifres66011 = ifres66012

													} else {
														ifres66011 = False

													}

													var ifres66010 Obj

													if True == ifres66011 {
														ifres66010 = True

													} else {
														ifres66010 = False

													}

													ifres66009 = ifres66010

												} else {
													ifres66009 = False

												}

												var ifres66008 Obj

												if True == ifres66009 {
													ifres66008 = True

												} else {
													ifres66008 = False

												}

												ifres66007 = ifres66008

											} else {
												ifres66007 = False

											}

											var ifres66006 Obj

											if True == ifres66007 {
												ifres66006 = True

											} else {
												ifres66006 = False

											}

											ifres66005 = ifres66006

										} else {
											ifres66005 = False

										}

										var ifres66004 Obj

										if True == ifres66005 {
											ifres66004 = True

										} else {
											ifres66004 = False

										}

										ifres66003 = ifres66004

									} else {
										ifres66003 = False

									}

									var ifres66002 Obj

									if True == ifres66003 {
										ifres66002 = True

									} else {
										ifres66002 = False

									}

									ifres66001 = ifres66002

								} else {
									ifres66001 = False

								}

								if True == ifres66001 {
									tmp65980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

									tmp65981 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

									tmp65982 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65983 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65984 := Call(__e, tmp65983, V385)

									tmp65985 := Call(__e, tmp65982, tmp65984)

									tmp65986 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65988 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65989 := Call(__e, tmp65988, V385)

									tmp65990 := Call(__e, tmp65987, tmp65989)

									tmp65991 := Call(__e, tmp65986, tmp65990)

									tmp65992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65994 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp65995 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp65996 := Call(__e, tmp65995, V385)

									tmp65997 := Call(__e, tmp65994, tmp65996)

									tmp65998 := Call(__e, tmp65993, tmp65997)

									tmp65999 := Call(__e, tmp65992, tmp65998)

									tmp66000 := Call(__e, tmp65981, tmp65985, tmp65991, tmp65999)

									__e.TailApply(tmp65980, tmp66000)
									return

								} else {
									tmp65978 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp65979 := Call(__e, tmp65978, V385)

									var ifres65948 Obj

									if True == tmp65979 {
										tmp65974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp65975 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp65976 := Call(__e, tmp65975, V385)

										tmp65977 := Call(__e, tmp65974, symwhere, tmp65976)

										var ifres65950 Obj

										if True == tmp65977 {
											tmp65970 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp65971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp65972 := Call(__e, tmp65971, V385)

											tmp65973 := Call(__e, tmp65970, tmp65972)

											var ifres65952 Obj

											if True == tmp65973 {
												tmp65964 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp65965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp65966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp65967 := Call(__e, tmp65966, V385)

												tmp65968 := Call(__e, tmp65965, tmp65967)

												tmp65969 := Call(__e, tmp65964, tmp65968)

												var ifres65954 Obj

												if True == tmp65969 {
													tmp65956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp65957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp65958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp65959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp65960 := Call(__e, tmp65959, V385)

													tmp65961 := Call(__e, tmp65958, tmp65960)

													tmp65962 := Call(__e, tmp65957, tmp65961)

													tmp65963 := Call(__e, tmp65956, Nil, tmp65962)

													var ifres65955 Obj

													if True == tmp65963 {
														ifres65955 = True

													} else {
														ifres65955 = False

													}

													ifres65954 = ifres65955

												} else {
													ifres65954 = False

												}

												var ifres65953 Obj

												if True == ifres65954 {
													ifres65953 = True

												} else {
													ifres65953 = False

												}

												ifres65952 = ifres65953

											} else {
												ifres65952 = False

											}

											var ifres65951 Obj

											if True == ifres65952 {
												ifres65951 = True

											} else {
												ifres65951 = False

											}

											ifres65950 = ifres65951

										} else {
											ifres65950 = False

										}

										var ifres65949 Obj

										if True == ifres65950 {
											ifres65949 = True

										} else {
											ifres65949 = False

										}

										ifres65948 = ifres65949

									} else {
										ifres65948 = False

									}

									if True == ifres65948 {
										tmp65935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add__test)

										tmp65936 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp65937 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65938 := Call(__e, tmp65937, V385)

										tmp65939 := Call(__e, tmp65936, tmp65938)

										tmp65940 := Call(__e, tmp65935, tmp65939)

										_ = tmp65940

										tmp65941 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

										tmp65942 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp65943 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp65945 := Call(__e, tmp65944, V385)

										tmp65946 := Call(__e, tmp65943, tmp65945)

										tmp65947 := Call(__e, tmp65942, tmp65946)

										__e.TailApply(tmp65941, tmp65947)
										return

									} else {
										tmp65933 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp65934 := Call(__e, tmp65933, V385)

										var ifres65919 Obj

										if True == tmp65934 {
											tmp65929 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp65930 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp65931 := Call(__e, tmp65930, V385)

											tmp65932 := Call(__e, tmp65929, tmp65931)

											var ifres65921 Obj

											if True == tmp65932 {
												tmp65923 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp65924 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp65925 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp65926 := Call(__e, tmp65925, V385)

												tmp65927 := Call(__e, tmp65924, tmp65926)

												tmp65928 := Call(__e, tmp65923, Nil, tmp65927)

												var ifres65922 Obj

												if True == tmp65928 {
													ifres65922 = True

												} else {
													ifres65922 = False

												}

												ifres65921 = ifres65922

											} else {
												ifres65921 = False

											}

											var ifres65920 Obj

											if True == ifres65921 {
												ifres65920 = True

											} else {
												ifres65920 = False

											}

											ifres65919 = ifres65920

										} else {
											ifres65919 = False

										}

										if True == ifres65919 {
											tmp65904 := MakeNative(func(__e *ControlFlow) {
												Z := __e.Get(1)
												_ = Z
												tmp65911 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp65912 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp65913 := Call(__e, tmp65912, V385)

												tmp65914 := Call(__e, tmp65911, tmp65913, Z)

												if True == tmp65914 {
													__e.Return(V385)
													return
												} else {
													tmp65906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

													tmp65907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

													tmp65908 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp65909 := Call(__e, tmp65908, V385)

													tmp65910 := Call(__e, tmp65907, Z, tmp65909)

													__e.TailApply(tmp65906, tmp65910)
													return

												}

											}, 1)

											tmp65915 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reduce__help)

											tmp65916 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp65917 := Call(__e, tmp65916, V385)

											tmp65918 := Call(__e, tmp65915, tmp65917)

											__e.TailApply(tmp65904, tmp65918)
											return

										} else {
											__e.Return(V385)
											return
										}

									}

								}

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp67045 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reduce__help, tmp65894)

	_ = tmp67045

	tmp67046 := MakeNative(func(__e *ControlFlow) {
		V387 := __e.Get(1)
		_ = V387
		tmp67049 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp67050 := Call(__e, tmp67049, MakeString(""), V387)

		if True == tmp67050 {
			__e.Return(False)
			return
		} else {
			tmp67048 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

			__e.TailApply(tmp67048, V387)
			return

		}

	}, 1)

	tmp67051 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_7string_2, tmp67046)

	_ = tmp67051

	tmp67052 := MakeNative(func(__e *ControlFlow) {
		V389 := __e.Get(1)
		_ = V389
		tmp67059 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		tmp67060 := Call(__e, tmp67059, V389)

		if True == tmp67060 {
			tmp67055 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp67056 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp67057 := Call(__e, tmp67056, V389, MakeNumber(0))

			tmp67058 := Call(__e, tmp67055, tmp67057, MakeNumber(0))

			if True == tmp67058 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp67061 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_7vector_2, tmp67052)

	_ = tmp67061

	tmp67062 := MakeNative(func(__e *ControlFlow) {
		V402 := __e.Get(1)
		_ = V402
		V403 := __e.Get(2)
		_ = V403
		V404 := __e.Get(3)
		_ = V404
		tmp67193 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp67194 := Call(__e, tmp67193, V404, V403)

		if True == tmp67194 {
			__e.Return(V402)
			return
		} else {
			tmp67191 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp67192 := Call(__e, tmp67191, V404)

			var ifres67153 Obj

			if True == tmp67192 {
				tmp67187 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp67188 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67189 := Call(__e, tmp67188, V404)

				tmp67190 := Call(__e, tmp67187, symlambda, tmp67189)

				var ifres67155 Obj

				if True == tmp67190 {
					tmp67183 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67184 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67185 := Call(__e, tmp67184, V404)

					tmp67186 := Call(__e, tmp67183, tmp67185)

					var ifres67157 Obj

					if True == tmp67186 {
						tmp67177 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67178 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67180 := Call(__e, tmp67179, V404)

						tmp67181 := Call(__e, tmp67178, tmp67180)

						tmp67182 := Call(__e, tmp67177, tmp67181)

						var ifres67159 Obj

						if True == tmp67182 {
							tmp67169 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp67170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67172 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67173 := Call(__e, tmp67172, V404)

							tmp67174 := Call(__e, tmp67171, tmp67173)

							tmp67175 := Call(__e, tmp67170, tmp67174)

							tmp67176 := Call(__e, tmp67169, Nil, tmp67175)

							var ifres67161 Obj

							if True == tmp67176 {
								tmp67163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clash_2)

								tmp67164 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67165 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67166 := Call(__e, tmp67165, V404)

								tmp67167 := Call(__e, tmp67164, tmp67166)

								tmp67168 := Call(__e, tmp67163, tmp67167, V403)

								var ifres67162 Obj

								if True == tmp67168 {
									ifres67162 = True

								} else {
									ifres67162 = False

								}

								ifres67161 = ifres67162

							} else {
								ifres67161 = False

							}

							var ifres67160 Obj

							if True == ifres67161 {
								ifres67160 = True

							} else {
								ifres67160 = False

							}

							ifres67159 = ifres67160

						} else {
							ifres67159 = False

						}

						var ifres67158 Obj

						if True == ifres67159 {
							ifres67158 = True

						} else {
							ifres67158 = False

						}

						ifres67157 = ifres67158

					} else {
						ifres67157 = False

					}

					var ifres67156 Obj

					if True == ifres67157 {
						ifres67156 = True

					} else {
						ifres67156 = False

					}

					ifres67155 = ifres67156

				} else {
					ifres67155 = False

				}

				var ifres67154 Obj

				if True == ifres67155 {
					ifres67154 = True

				} else {
					ifres67154 = False

				}

				ifres67153 = ifres67154

			} else {
				ifres67153 = False

			}

			if True == ifres67153 {
				__e.Return(V404)
				return
			} else {
				tmp67151 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp67152 := Call(__e, tmp67151, V404)

				var ifres67101 Obj

				if True == tmp67152 {
					tmp67147 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp67148 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67149 := Call(__e, tmp67148, V404)

					tmp67150 := Call(__e, tmp67147, symlet, tmp67149)

					var ifres67103 Obj

					if True == tmp67150 {
						tmp67143 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67144 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67145 := Call(__e, tmp67144, V404)

						tmp67146 := Call(__e, tmp67143, tmp67145)

						var ifres67105 Obj

						if True == tmp67146 {
							tmp67137 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67138 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67139 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67140 := Call(__e, tmp67139, V404)

							tmp67141 := Call(__e, tmp67138, tmp67140)

							tmp67142 := Call(__e, tmp67137, tmp67141)

							var ifres67107 Obj

							if True == tmp67142 {
								tmp67129 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67132 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67133 := Call(__e, tmp67132, V404)

								tmp67134 := Call(__e, tmp67131, tmp67133)

								tmp67135 := Call(__e, tmp67130, tmp67134)

								tmp67136 := Call(__e, tmp67129, tmp67135)

								var ifres67109 Obj

								if True == tmp67136 {
									tmp67119 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp67120 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67121 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67124 := Call(__e, tmp67123, V404)

									tmp67125 := Call(__e, tmp67122, tmp67124)

									tmp67126 := Call(__e, tmp67121, tmp67125)

									tmp67127 := Call(__e, tmp67120, tmp67126)

									tmp67128 := Call(__e, tmp67119, Nil, tmp67127)

									var ifres67111 Obj

									if True == tmp67128 {
										tmp67113 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clash_2)

										tmp67114 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67115 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67116 := Call(__e, tmp67115, V404)

										tmp67117 := Call(__e, tmp67114, tmp67116)

										tmp67118 := Call(__e, tmp67113, tmp67117, V403)

										var ifres67112 Obj

										if True == tmp67118 {
											ifres67112 = True

										} else {
											ifres67112 = False

										}

										ifres67111 = ifres67112

									} else {
										ifres67111 = False

									}

									var ifres67110 Obj

									if True == ifres67111 {
										ifres67110 = True

									} else {
										ifres67110 = False

									}

									ifres67109 = ifres67110

								} else {
									ifres67109 = False

								}

								var ifres67108 Obj

								if True == ifres67109 {
									ifres67108 = True

								} else {
									ifres67108 = False

								}

								ifres67107 = ifres67108

							} else {
								ifres67107 = False

							}

							var ifres67106 Obj

							if True == ifres67107 {
								ifres67106 = True

							} else {
								ifres67106 = False

							}

							ifres67105 = ifres67106

						} else {
							ifres67105 = False

						}

						var ifres67104 Obj

						if True == ifres67105 {
							ifres67104 = True

						} else {
							ifres67104 = False

						}

						ifres67103 = ifres67104

					} else {
						ifres67103 = False

					}

					var ifres67102 Obj

					if True == ifres67103 {
						ifres67102 = True

					} else {
						ifres67102 = False

					}

					ifres67101 = ifres67102

				} else {
					ifres67101 = False

				}

				if True == ifres67101 {
					tmp67078 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67080 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67081 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67082 := Call(__e, tmp67081, V404)

					tmp67083 := Call(__e, tmp67080, tmp67082)

					tmp67084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

					tmp67086 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67087 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67088 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67089 := Call(__e, tmp67088, V404)

					tmp67090 := Call(__e, tmp67087, tmp67089)

					tmp67091 := Call(__e, tmp67086, tmp67090)

					tmp67092 := Call(__e, tmp67085, V402, V403, tmp67091)

					tmp67093 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67095 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67096 := Call(__e, tmp67095, V404)

					tmp67097 := Call(__e, tmp67094, tmp67096)

					tmp67098 := Call(__e, tmp67093, tmp67097)

					tmp67099 := Call(__e, tmp67084, tmp67092, tmp67098)

					tmp67100 := Call(__e, tmp67079, tmp67083, tmp67099)

					__e.TailApply(tmp67078, symlet, tmp67100)
					return

				} else {
					tmp67076 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67077 := Call(__e, tmp67076, V404)

					if True == tmp67077 {
						tmp67067 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

						tmp67069 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67070 := Call(__e, tmp67069, V404)

						tmp67071 := Call(__e, tmp67068, V402, V403, tmp67070)

						tmp67072 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

						tmp67073 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67074 := Call(__e, tmp67073, V404)

						tmp67075 := Call(__e, tmp67072, V402, V403, tmp67074)

						__e.TailApply(tmp67067, tmp67071, tmp67075)
						return

					} else {
						__e.Return(V404)
						return
					}

				}

			}

		}

	}, 3)

	tmp67195 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ebr, tmp67062)

	_ = tmp67195

	tmp67196 := MakeNative(func(__e *ControlFlow) {
		V416 := __e.Get(1)
		_ = V416
		V417 := __e.Get(2)
		_ = V417
		tmp67211 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp67212 := Call(__e, tmp67211, V417, V416)

		if True == tmp67212 {
			__e.Return(True)
			return
		} else {
			tmp67209 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp67210 := Call(__e, tmp67209, V417)

			if True == tmp67210 {
				tmp67205 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clash_2)

				tmp67206 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67207 := Call(__e, tmp67206, V417)

				tmp67208 := Call(__e, tmp67205, V416, tmp67207)

				if True == tmp67208 {
					__e.Return(True)
					return
				} else {
					tmp67201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4clash_2)

					tmp67202 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67203 := Call(__e, tmp67202, V417)

					tmp67204 := Call(__e, tmp67201, V416, tmp67203)

					if True == tmp67204 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp67213 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clash_2, tmp67196)

	_ = tmp67213

	tmp67214 := MakeNative(func(__e *ControlFlow) {
		V419 := __e.Get(1)
		_ = V419
		tmp67215 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp67216 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp67217 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp67218 := Call(__e, tmp67217, symshen_4_dteststack_d)

		tmp67219 := Call(__e, tmp67216, V419, tmp67218)

		__e.TailApply(tmp67215, symshen_4_dteststack_d, tmp67219)
		return

	}, 1)

	tmp67220 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add__test, tmp67214)

	_ = tmp67220

	tmp67221 := MakeNative(func(__e *ControlFlow) {
		V423 := __e.Get(1)
		_ = V423
		V424 := __e.Get(2)
		_ = V424
		V425 := __e.Get(3)
		_ = V425
		tmp67222 := MakeNative(func(__e *ControlFlow) {
			Err := __e.Get(1)
			_ = Err
			tmp67223 := MakeNative(func(__e *ControlFlow) {
				Cases := __e.Get(1)
				_ = Cases
				tmp67224 := MakeNative(func(__e *ControlFlow) {
					EncodeChoices := __e.Get(1)
					_ = EncodeChoices
					tmp67225 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cond_1form)

					__e.TailApply(tmp67225, EncodeChoices)
					return

				}, 1)

				tmp67226 := Call(__e, PrimNS1Value(symns2_1value), symshen_4encode_1choices)

				tmp67227 := Call(__e, tmp67226, Cases, V423)

				__e.TailApply(tmp67224, tmp67227)
				return

			}, 1)

			tmp67228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4case_1form)

			tmp67229 := Call(__e, tmp67228, V425, Err)

			__e.TailApply(tmp67223, tmp67229)
			return

		}, 1)

		tmp67230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4err_1condition)

		tmp67231 := Call(__e, tmp67230, V423)

		__e.TailApply(tmp67222, tmp67231)
		return

	}, 3)

	tmp67232 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cond_1expression, tmp67221)

	_ = tmp67232

	tmp67233 := MakeNative(func(__e *ControlFlow) {
		V429 := __e.Get(1)
		_ = V429
		tmp67273 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp67274 := Call(__e, tmp67273, V429)

		var ifres67241 Obj

		if True == tmp67274 {
			tmp67269 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp67270 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp67271 := Call(__e, tmp67270, V429)

			tmp67272 := Call(__e, tmp67269, tmp67271)

			var ifres67243 Obj

			if True == tmp67272 {
				tmp67263 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp67264 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67265 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67266 := Call(__e, tmp67265, V429)

				tmp67267 := Call(__e, tmp67264, tmp67266)

				tmp67268 := Call(__e, tmp67263, True, tmp67267)

				var ifres67245 Obj

				if True == tmp67268 {
					tmp67257 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67259 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67260 := Call(__e, tmp67259, V429)

					tmp67261 := Call(__e, tmp67258, tmp67260)

					tmp67262 := Call(__e, tmp67257, tmp67261)

					var ifres67247 Obj

					if True == tmp67262 {
						tmp67249 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp67250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67251 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67252 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67253 := Call(__e, tmp67252, V429)

						tmp67254 := Call(__e, tmp67251, tmp67253)

						tmp67255 := Call(__e, tmp67250, tmp67254)

						tmp67256 := Call(__e, tmp67249, Nil, tmp67255)

						var ifres67248 Obj

						if True == tmp67256 {
							ifres67248 = True

						} else {
							ifres67248 = False

						}

						ifres67247 = ifres67248

					} else {
						ifres67247 = False

					}

					var ifres67246 Obj

					if True == ifres67247 {
						ifres67246 = True

					} else {
						ifres67246 = False

					}

					ifres67245 = ifres67246

				} else {
					ifres67245 = False

				}

				var ifres67244 Obj

				if True == ifres67245 {
					ifres67244 = True

				} else {
					ifres67244 = False

				}

				ifres67243 = ifres67244

			} else {
				ifres67243 = False

			}

			var ifres67242 Obj

			if True == ifres67243 {
				ifres67242 = True

			} else {
				ifres67242 = False

			}

			ifres67241 = ifres67242

		} else {
			ifres67241 = False

		}

		if True == ifres67241 {
			tmp67236 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp67237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp67238 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp67239 := Call(__e, tmp67238, V429)

			tmp67240 := Call(__e, tmp67237, tmp67239)

			__e.TailApply(tmp67236, tmp67240)
			return

		} else {
			tmp67235 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp67235, symcond, V429)
			return

		}

	}, 1)

	tmp67275 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cond_1form, tmp67233)

	_ = tmp67275

	tmp67276 := MakeNative(func(__e *ControlFlow) {
		V434 := __e.Get(1)
		_ = V434
		V435 := __e.Get(2)
		_ = V435
		tmp67732 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp67733 := Call(__e, tmp67732, Nil, V434)

		if True == tmp67733 {
			__e.Return(Nil)
			return
		} else {
			tmp67730 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp67731 := Call(__e, tmp67730, V434)

			var ifres67644 Obj

			if True == tmp67731 {
				tmp67726 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp67727 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67728 := Call(__e, tmp67727, V434)

				tmp67729 := Call(__e, tmp67726, tmp67728)

				var ifres67646 Obj

				if True == tmp67729 {
					tmp67720 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp67721 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67722 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67723 := Call(__e, tmp67722, V434)

					tmp67724 := Call(__e, tmp67721, tmp67723)

					tmp67725 := Call(__e, tmp67720, True, tmp67724)

					var ifres67648 Obj

					if True == tmp67725 {
						tmp67714 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67716 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67717 := Call(__e, tmp67716, V434)

						tmp67718 := Call(__e, tmp67715, tmp67717)

						tmp67719 := Call(__e, tmp67714, tmp67718)

						var ifres67650 Obj

						if True == tmp67719 {
							tmp67706 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67707 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67708 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67709 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67710 := Call(__e, tmp67709, V434)

							tmp67711 := Call(__e, tmp67708, tmp67710)

							tmp67712 := Call(__e, tmp67707, tmp67711)

							tmp67713 := Call(__e, tmp67706, tmp67712)

							var ifres67652 Obj

							if True == tmp67713 {
								tmp67696 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp67697 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67698 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67699 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67700 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67701 := Call(__e, tmp67700, V434)

								tmp67702 := Call(__e, tmp67699, tmp67701)

								tmp67703 := Call(__e, tmp67698, tmp67702)

								tmp67704 := Call(__e, tmp67697, tmp67703)

								tmp67705 := Call(__e, tmp67696, symshen_4choicepoint_b, tmp67704)

								var ifres67654 Obj

								if True == tmp67705 {
									tmp67686 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp67687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67688 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67689 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67690 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67691 := Call(__e, tmp67690, V434)

									tmp67692 := Call(__e, tmp67689, tmp67691)

									tmp67693 := Call(__e, tmp67688, tmp67692)

									tmp67694 := Call(__e, tmp67687, tmp67693)

									tmp67695 := Call(__e, tmp67686, tmp67694)

									var ifres67656 Obj

									if True == tmp67695 {
										tmp67674 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp67675 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67676 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67677 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67678 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67680 := Call(__e, tmp67679, V434)

										tmp67681 := Call(__e, tmp67678, tmp67680)

										tmp67682 := Call(__e, tmp67677, tmp67681)

										tmp67683 := Call(__e, tmp67676, tmp67682)

										tmp67684 := Call(__e, tmp67675, tmp67683)

										tmp67685 := Call(__e, tmp67674, Nil, tmp67684)

										var ifres67658 Obj

										if True == tmp67685 {
											tmp67666 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp67667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67669 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67670 := Call(__e, tmp67669, V434)

											tmp67671 := Call(__e, tmp67668, tmp67670)

											tmp67672 := Call(__e, tmp67667, tmp67671)

											tmp67673 := Call(__e, tmp67666, Nil, tmp67672)

											var ifres67660 Obj

											if True == tmp67673 {
												tmp67662 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67663 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67664 := Call(__e, tmp67663, V434)

												tmp67665 := Call(__e, tmp67662, Nil, tmp67664)

												var ifres67661 Obj

												if True == tmp67665 {
													ifres67661 = True

												} else {
													ifres67661 = False

												}

												ifres67660 = ifres67661

											} else {
												ifres67660 = False

											}

											var ifres67659 Obj

											if True == ifres67660 {
												ifres67659 = True

											} else {
												ifres67659 = False

											}

											ifres67658 = ifres67659

										} else {
											ifres67658 = False

										}

										var ifres67657 Obj

										if True == ifres67658 {
											ifres67657 = True

										} else {
											ifres67657 = False

										}

										ifres67656 = ifres67657

									} else {
										ifres67656 = False

									}

									var ifres67655 Obj

									if True == ifres67656 {
										ifres67655 = True

									} else {
										ifres67655 = False

									}

									ifres67654 = ifres67655

								} else {
									ifres67654 = False

								}

								var ifres67653 Obj

								if True == ifres67654 {
									ifres67653 = True

								} else {
									ifres67653 = False

								}

								ifres67652 = ifres67653

							} else {
								ifres67652 = False

							}

							var ifres67651 Obj

							if True == ifres67652 {
								ifres67651 = True

							} else {
								ifres67651 = False

							}

							ifres67650 = ifres67651

						} else {
							ifres67650 = False

						}

						var ifres67649 Obj

						if True == ifres67650 {
							ifres67649 = True

						} else {
							ifres67649 = False

						}

						ifres67648 = ifres67649

					} else {
						ifres67648 = False

					}

					var ifres67647 Obj

					if True == ifres67648 {
						ifres67647 = True

					} else {
						ifres67647 = False

					}

					ifres67646 = ifres67647

				} else {
					ifres67646 = False

				}

				var ifres67645 Obj

				if True == ifres67646 {
					ifres67645 = True

				} else {
					ifres67645 = False

				}

				ifres67644 = ifres67645

			} else {
				ifres67644 = False

			}

			if True == ifres67644 {
				tmp67594 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67596 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67599 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67600 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp67602 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67603 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp67604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67605 := Call(__e, tmp67604, V434)

				tmp67606 := Call(__e, tmp67603, tmp67605)

				tmp67607 := Call(__e, tmp67602, tmp67606)

				tmp67608 := Call(__e, tmp67601, tmp67607)

				tmp67609 := Call(__e, tmp67600, tmp67608)

				tmp67610 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67611 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67612 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67614 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67615 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67617 := Call(__e, tmp67616, symfail, Nil)

				tmp67618 := Call(__e, tmp67615, tmp67617, Nil)

				tmp67619 := Call(__e, tmp67614, symResult, tmp67618)

				tmp67620 := Call(__e, tmp67613, sym_a, tmp67619)

				tmp67621 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67631 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp67632 := Call(__e, tmp67631, symshen_4_dinstalling_1kl_d)

				var ifres67622 Obj

				if True == tmp67632 {
					tmp67627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67629 := Call(__e, tmp67628, V435, Nil)

					tmp67630 := Call(__e, tmp67627, symshen_4sys_1error, tmp67629)

					ifres67622 = tmp67630

				} else {
					tmp67623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67625 := Call(__e, tmp67624, V435, Nil)

					tmp67626 := Call(__e, tmp67623, symshen_4f__error, tmp67625)

					ifres67622 = tmp67626

				}

				tmp67633 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67634 := Call(__e, tmp67633, symResult, Nil)

				tmp67635 := Call(__e, tmp67621, ifres67622, tmp67634)

				tmp67636 := Call(__e, tmp67612, tmp67620, tmp67635)

				tmp67637 := Call(__e, tmp67611, symif, tmp67636)

				tmp67638 := Call(__e, tmp67610, tmp67637, Nil)

				tmp67639 := Call(__e, tmp67599, tmp67609, tmp67638)

				tmp67640 := Call(__e, tmp67598, symResult, tmp67639)

				tmp67641 := Call(__e, tmp67597, symlet, tmp67640)

				tmp67642 := Call(__e, tmp67596, tmp67641, Nil)

				tmp67643 := Call(__e, tmp67595, True, tmp67642)

				__e.TailApply(tmp67594, tmp67643, Nil)
				return

			} else {
				tmp67592 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp67593 := Call(__e, tmp67592, V434)

				var ifres67512 Obj

				if True == tmp67593 {
					tmp67588 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67589 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67590 := Call(__e, tmp67589, V434)

					tmp67591 := Call(__e, tmp67588, tmp67590)

					var ifres67514 Obj

					if True == tmp67591 {
						tmp67582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp67583 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67584 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67585 := Call(__e, tmp67584, V434)

						tmp67586 := Call(__e, tmp67583, tmp67585)

						tmp67587 := Call(__e, tmp67582, True, tmp67586)

						var ifres67516 Obj

						if True == tmp67587 {
							tmp67576 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67577 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67578 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67579 := Call(__e, tmp67578, V434)

							tmp67580 := Call(__e, tmp67577, tmp67579)

							tmp67581 := Call(__e, tmp67576, tmp67580)

							var ifres67518 Obj

							if True == tmp67581 {
								tmp67568 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67569 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67570 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67571 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67572 := Call(__e, tmp67571, V434)

								tmp67573 := Call(__e, tmp67570, tmp67572)

								tmp67574 := Call(__e, tmp67569, tmp67573)

								tmp67575 := Call(__e, tmp67568, tmp67574)

								var ifres67520 Obj

								if True == tmp67575 {
									tmp67558 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp67559 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67560 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67562 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67563 := Call(__e, tmp67562, V434)

									tmp67564 := Call(__e, tmp67561, tmp67563)

									tmp67565 := Call(__e, tmp67560, tmp67564)

									tmp67566 := Call(__e, tmp67559, tmp67565)

									tmp67567 := Call(__e, tmp67558, symshen_4choicepoint_b, tmp67566)

									var ifres67522 Obj

									if True == tmp67567 {
										tmp67548 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp67549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67551 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67552 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67553 := Call(__e, tmp67552, V434)

										tmp67554 := Call(__e, tmp67551, tmp67553)

										tmp67555 := Call(__e, tmp67550, tmp67554)

										tmp67556 := Call(__e, tmp67549, tmp67555)

										tmp67557 := Call(__e, tmp67548, tmp67556)

										var ifres67524 Obj

										if True == tmp67557 {
											tmp67536 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp67537 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67538 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67539 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67540 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67541 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67542 := Call(__e, tmp67541, V434)

											tmp67543 := Call(__e, tmp67540, tmp67542)

											tmp67544 := Call(__e, tmp67539, tmp67543)

											tmp67545 := Call(__e, tmp67538, tmp67544)

											tmp67546 := Call(__e, tmp67537, tmp67545)

											tmp67547 := Call(__e, tmp67536, Nil, tmp67546)

											var ifres67526 Obj

											if True == tmp67547 {
												tmp67528 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67529 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67530 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67531 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67532 := Call(__e, tmp67531, V434)

												tmp67533 := Call(__e, tmp67530, tmp67532)

												tmp67534 := Call(__e, tmp67529, tmp67533)

												tmp67535 := Call(__e, tmp67528, Nil, tmp67534)

												var ifres67527 Obj

												if True == tmp67535 {
													ifres67527 = True

												} else {
													ifres67527 = False

												}

												ifres67526 = ifres67527

											} else {
												ifres67526 = False

											}

											var ifres67525 Obj

											if True == ifres67526 {
												ifres67525 = True

											} else {
												ifres67525 = False

											}

											ifres67524 = ifres67525

										} else {
											ifres67524 = False

										}

										var ifres67523 Obj

										if True == ifres67524 {
											ifres67523 = True

										} else {
											ifres67523 = False

										}

										ifres67522 = ifres67523

									} else {
										ifres67522 = False

									}

									var ifres67521 Obj

									if True == ifres67522 {
										ifres67521 = True

									} else {
										ifres67521 = False

									}

									ifres67520 = ifres67521

								} else {
									ifres67520 = False

								}

								var ifres67519 Obj

								if True == ifres67520 {
									ifres67519 = True

								} else {
									ifres67519 = False

								}

								ifres67518 = ifres67519

							} else {
								ifres67518 = False

							}

							var ifres67517 Obj

							if True == ifres67518 {
								ifres67517 = True

							} else {
								ifres67517 = False

							}

							ifres67516 = ifres67517

						} else {
							ifres67516 = False

						}

						var ifres67515 Obj

						if True == ifres67516 {
							ifres67515 = True

						} else {
							ifres67515 = False

						}

						ifres67514 = ifres67515

					} else {
						ifres67514 = False

					}

					var ifres67513 Obj

					if True == ifres67514 {
						ifres67513 = True

					} else {
						ifres67513 = False

					}

					ifres67512 = ifres67513

				} else {
					ifres67512 = False

				}

				if True == ifres67512 {
					tmp67467 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67468 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67469 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67470 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67471 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67472 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67473 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67474 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67475 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67477 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67478 := Call(__e, tmp67477, V434)

					tmp67479 := Call(__e, tmp67476, tmp67478)

					tmp67480 := Call(__e, tmp67475, tmp67479)

					tmp67481 := Call(__e, tmp67474, tmp67480)

					tmp67482 := Call(__e, tmp67473, tmp67481)

					tmp67483 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67484 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67485 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67486 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67487 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67489 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67490 := Call(__e, tmp67489, symfail, Nil)

					tmp67491 := Call(__e, tmp67488, tmp67490, Nil)

					tmp67492 := Call(__e, tmp67487, symResult, tmp67491)

					tmp67493 := Call(__e, tmp67486, sym_a, tmp67492)

					tmp67494 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cond_1form)

					tmp67496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4encode_1choices)

					tmp67497 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67498 := Call(__e, tmp67497, V434)

					tmp67499 := Call(__e, tmp67496, tmp67498, V435)

					tmp67500 := Call(__e, tmp67495, tmp67499)

					tmp67501 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67502 := Call(__e, tmp67501, symResult, Nil)

					tmp67503 := Call(__e, tmp67494, tmp67500, tmp67502)

					tmp67504 := Call(__e, tmp67485, tmp67493, tmp67503)

					tmp67505 := Call(__e, tmp67484, symif, tmp67504)

					tmp67506 := Call(__e, tmp67483, tmp67505, Nil)

					tmp67507 := Call(__e, tmp67472, tmp67482, tmp67506)

					tmp67508 := Call(__e, tmp67471, symResult, tmp67507)

					tmp67509 := Call(__e, tmp67470, symlet, tmp67508)

					tmp67510 := Call(__e, tmp67469, tmp67509, Nil)

					tmp67511 := Call(__e, tmp67468, True, tmp67510)

					__e.TailApply(tmp67467, tmp67511, Nil)
					return

				} else {
					tmp67465 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67466 := Call(__e, tmp67465, V434)

					var ifres67393 Obj

					if True == tmp67466 {
						tmp67461 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67463 := Call(__e, tmp67462, V434)

						tmp67464 := Call(__e, tmp67461, tmp67463)

						var ifres67395 Obj

						if True == tmp67464 {
							tmp67455 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67457 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67458 := Call(__e, tmp67457, V434)

							tmp67459 := Call(__e, tmp67456, tmp67458)

							tmp67460 := Call(__e, tmp67455, tmp67459)

							var ifres67397 Obj

							if True == tmp67460 {
								tmp67447 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67448 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67449 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67450 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67451 := Call(__e, tmp67450, V434)

								tmp67452 := Call(__e, tmp67449, tmp67451)

								tmp67453 := Call(__e, tmp67448, tmp67452)

								tmp67454 := Call(__e, tmp67447, tmp67453)

								var ifres67399 Obj

								if True == tmp67454 {
									tmp67437 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp67438 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67439 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67440 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67441 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67442 := Call(__e, tmp67441, V434)

									tmp67443 := Call(__e, tmp67440, tmp67442)

									tmp67444 := Call(__e, tmp67439, tmp67443)

									tmp67445 := Call(__e, tmp67438, tmp67444)

									tmp67446 := Call(__e, tmp67437, symshen_4choicepoint_b, tmp67445)

									var ifres67401 Obj

									if True == tmp67446 {
										tmp67427 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp67428 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67429 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67430 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67431 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67432 := Call(__e, tmp67431, V434)

										tmp67433 := Call(__e, tmp67430, tmp67432)

										tmp67434 := Call(__e, tmp67429, tmp67433)

										tmp67435 := Call(__e, tmp67428, tmp67434)

										tmp67436 := Call(__e, tmp67427, tmp67435)

										var ifres67403 Obj

										if True == tmp67436 {
											tmp67415 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp67416 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67417 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67418 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67419 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67420 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67421 := Call(__e, tmp67420, V434)

											tmp67422 := Call(__e, tmp67419, tmp67421)

											tmp67423 := Call(__e, tmp67418, tmp67422)

											tmp67424 := Call(__e, tmp67417, tmp67423)

											tmp67425 := Call(__e, tmp67416, tmp67424)

											tmp67426 := Call(__e, tmp67415, Nil, tmp67425)

											var ifres67405 Obj

											if True == tmp67426 {
												tmp67407 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67408 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67409 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67410 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67411 := Call(__e, tmp67410, V434)

												tmp67412 := Call(__e, tmp67409, tmp67411)

												tmp67413 := Call(__e, tmp67408, tmp67412)

												tmp67414 := Call(__e, tmp67407, Nil, tmp67413)

												var ifres67406 Obj

												if True == tmp67414 {
													ifres67406 = True

												} else {
													ifres67406 = False

												}

												ifres67405 = ifres67406

											} else {
												ifres67405 = False

											}

											var ifres67404 Obj

											if True == ifres67405 {
												ifres67404 = True

											} else {
												ifres67404 = False

											}

											ifres67403 = ifres67404

										} else {
											ifres67403 = False

										}

										var ifres67402 Obj

										if True == ifres67403 {
											ifres67402 = True

										} else {
											ifres67402 = False

										}

										ifres67401 = ifres67402

									} else {
										ifres67401 = False

									}

									var ifres67400 Obj

									if True == ifres67401 {
										ifres67400 = True

									} else {
										ifres67400 = False

									}

									ifres67399 = ifres67400

								} else {
									ifres67399 = False

								}

								var ifres67398 Obj

								if True == ifres67399 {
									ifres67398 = True

								} else {
									ifres67398 = False

								}

								ifres67397 = ifres67398

							} else {
								ifres67397 = False

							}

							var ifres67396 Obj

							if True == ifres67397 {
								ifres67396 = True

							} else {
								ifres67396 = False

							}

							ifres67395 = ifres67396

						} else {
							ifres67395 = False

						}

						var ifres67394 Obj

						if True == ifres67395 {
							ifres67394 = True

						} else {
							ifres67394 = False

						}

						ifres67393 = ifres67394

					} else {
						ifres67393 = False

					}

					if True == ifres67393 {
						tmp67316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67324 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cond_1form)

						tmp67325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4encode_1choices)

						tmp67326 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67327 := Call(__e, tmp67326, V434)

						tmp67328 := Call(__e, tmp67325, tmp67327, V435)

						tmp67329 := Call(__e, tmp67324, tmp67328)

						tmp67330 := Call(__e, tmp67323, tmp67329, Nil)

						tmp67331 := Call(__e, tmp67322, symfreeze, tmp67330)

						tmp67332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67335 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67336 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67337 := Call(__e, tmp67336, V434)

						tmp67338 := Call(__e, tmp67335, tmp67337)

						tmp67339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67343 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67344 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67345 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67347 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67348 := Call(__e, tmp67347, V434)

						tmp67349 := Call(__e, tmp67346, tmp67348)

						tmp67350 := Call(__e, tmp67345, tmp67349)

						tmp67351 := Call(__e, tmp67344, tmp67350)

						tmp67352 := Call(__e, tmp67343, tmp67351)

						tmp67353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67360 := Call(__e, tmp67359, symfail, Nil)

						tmp67361 := Call(__e, tmp67358, tmp67360, Nil)

						tmp67362 := Call(__e, tmp67357, symResult, tmp67361)

						tmp67363 := Call(__e, tmp67356, sym_a, tmp67362)

						tmp67364 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67365 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67367 := Call(__e, tmp67366, symFreeze, Nil)

						tmp67368 := Call(__e, tmp67365, symthaw, tmp67367)

						tmp67369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67370 := Call(__e, tmp67369, symResult, Nil)

						tmp67371 := Call(__e, tmp67364, tmp67368, tmp67370)

						tmp67372 := Call(__e, tmp67355, tmp67363, tmp67371)

						tmp67373 := Call(__e, tmp67354, symif, tmp67372)

						tmp67374 := Call(__e, tmp67353, tmp67373, Nil)

						tmp67375 := Call(__e, tmp67342, tmp67352, tmp67374)

						tmp67376 := Call(__e, tmp67341, symResult, tmp67375)

						tmp67377 := Call(__e, tmp67340, symlet, tmp67376)

						tmp67378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67379 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67380 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67381 := Call(__e, tmp67380, symFreeze, Nil)

						tmp67382 := Call(__e, tmp67379, symthaw, tmp67381)

						tmp67383 := Call(__e, tmp67378, tmp67382, Nil)

						tmp67384 := Call(__e, tmp67339, tmp67377, tmp67383)

						tmp67385 := Call(__e, tmp67334, tmp67338, tmp67384)

						tmp67386 := Call(__e, tmp67333, symif, tmp67385)

						tmp67387 := Call(__e, tmp67332, tmp67386, Nil)

						tmp67388 := Call(__e, tmp67321, tmp67331, tmp67387)

						tmp67389 := Call(__e, tmp67320, symFreeze, tmp67388)

						tmp67390 := Call(__e, tmp67319, symlet, tmp67389)

						tmp67391 := Call(__e, tmp67318, tmp67390, Nil)

						tmp67392 := Call(__e, tmp67317, True, tmp67391)

						__e.TailApply(tmp67316, tmp67392, Nil)
						return

					} else {
						tmp67314 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67315 := Call(__e, tmp67314, V434)

						var ifres67290 Obj

						if True == tmp67315 {
							tmp67310 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67311 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67312 := Call(__e, tmp67311, V434)

							tmp67313 := Call(__e, tmp67310, tmp67312)

							var ifres67292 Obj

							if True == tmp67313 {
								tmp67304 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67305 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67306 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67307 := Call(__e, tmp67306, V434)

								tmp67308 := Call(__e, tmp67305, tmp67307)

								tmp67309 := Call(__e, tmp67304, tmp67308)

								var ifres67294 Obj

								if True == tmp67309 {
									tmp67296 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp67297 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67299 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67300 := Call(__e, tmp67299, V434)

									tmp67301 := Call(__e, tmp67298, tmp67300)

									tmp67302 := Call(__e, tmp67297, tmp67301)

									tmp67303 := Call(__e, tmp67296, Nil, tmp67302)

									var ifres67295 Obj

									if True == tmp67303 {
										ifres67295 = True

									} else {
										ifres67295 = False

									}

									ifres67294 = ifres67295

								} else {
									ifres67294 = False

								}

								var ifres67293 Obj

								if True == ifres67294 {
									ifres67293 = True

								} else {
									ifres67293 = False

								}

								ifres67292 = ifres67293

							} else {
								ifres67292 = False

							}

							var ifres67291 Obj

							if True == ifres67292 {
								ifres67291 = True

							} else {
								ifres67291 = False

							}

							ifres67290 = ifres67291

						} else {
							ifres67290 = False

						}

						if True == ifres67290 {
							tmp67283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp67284 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67285 := Call(__e, tmp67284, V434)

							tmp67286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4encode_1choices)

							tmp67287 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp67288 := Call(__e, tmp67287, V434)

							tmp67289 := Call(__e, tmp67286, tmp67288, V435)

							__e.TailApply(tmp67283, tmp67285, tmp67289)
							return

						} else {
							tmp67282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(tmp67282, symshen_4encode_1choices)
							return

						}

					}

				}

			}

		}

	}, 2)

	tmp67734 := Call(__e, PrimNS1Value(symns2_1set), symshen_4encode_1choices, tmp67276)

	_ = tmp67734

	tmp67735 := MakeNative(func(__e *ControlFlow) {
		V442 := __e.Get(1)
		_ = V442
		V443 := __e.Get(2)
		_ = V443
		tmp68051 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp68052 := Call(__e, tmp68051, Nil, V442)

		if True == tmp68052 {
			tmp68050 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp68050, V443, Nil)
			return

		} else {
			tmp68048 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68049 := Call(__e, tmp68048, V442)

			var ifres67924 Obj

			if True == tmp68049 {
				tmp68044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp68045 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68046 := Call(__e, tmp68045, V442)

				tmp68047 := Call(__e, tmp68044, tmp68046)

				var ifres67926 Obj

				if True == tmp68047 {
					tmp68038 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp68039 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68040 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp68041 := Call(__e, tmp68040, V442)

					tmp68042 := Call(__e, tmp68039, tmp68041)

					tmp68043 := Call(__e, tmp68038, tmp68042)

					var ifres67928 Obj

					if True == tmp68043 {
						tmp68030 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp68031 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68032 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68033 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp68034 := Call(__e, tmp68033, V442)

						tmp68035 := Call(__e, tmp68032, tmp68034)

						tmp68036 := Call(__e, tmp68031, tmp68035)

						tmp68037 := Call(__e, tmp68030, sym_h, tmp68036)

						var ifres67930 Obj

						if True == tmp68037 {
							tmp68022 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp68023 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp68024 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp68025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp68026 := Call(__e, tmp68025, V442)

							tmp68027 := Call(__e, tmp68024, tmp68026)

							tmp68028 := Call(__e, tmp68023, tmp68027)

							tmp68029 := Call(__e, tmp68022, tmp68028)

							var ifres67932 Obj

							if True == tmp68029 {
								tmp68012 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp68013 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp68014 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp68015 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp68016 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp68017 := Call(__e, tmp68016, V442)

								tmp68018 := Call(__e, tmp68015, tmp68017)

								tmp68019 := Call(__e, tmp68014, tmp68018)

								tmp68020 := Call(__e, tmp68013, tmp68019)

								tmp68021 := Call(__e, tmp68012, symshen_4tests, tmp68020)

								var ifres67934 Obj

								if True == tmp68021 {
									tmp68002 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp68003 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68004 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp68005 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp68006 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp68007 := Call(__e, tmp68006, V442)

									tmp68008 := Call(__e, tmp68005, tmp68007)

									tmp68009 := Call(__e, tmp68004, tmp68008)

									tmp68010 := Call(__e, tmp68003, tmp68009)

									tmp68011 := Call(__e, tmp68002, Nil, tmp68010)

									var ifres67936 Obj

									if True == tmp68011 {
										tmp67996 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp67997 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67999 := Call(__e, tmp67998, V442)

										tmp68000 := Call(__e, tmp67997, tmp67999)

										tmp68001 := Call(__e, tmp67996, tmp68000)

										var ifres67938 Obj

										if True == tmp68001 {
											tmp67988 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp67989 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67991 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67992 := Call(__e, tmp67991, V442)

											tmp67993 := Call(__e, tmp67990, tmp67992)

											tmp67994 := Call(__e, tmp67989, tmp67993)

											tmp67995 := Call(__e, tmp67988, tmp67994)

											var ifres67940 Obj

											if True == tmp67995 {
												tmp67978 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67979 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67980 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67981 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67982 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67983 := Call(__e, tmp67982, V442)

												tmp67984 := Call(__e, tmp67981, tmp67983)

												tmp67985 := Call(__e, tmp67980, tmp67984)

												tmp67986 := Call(__e, tmp67979, tmp67985)

												tmp67987 := Call(__e, tmp67978, symshen_4choicepoint_b, tmp67986)

												var ifres67942 Obj

												if True == tmp67987 {
													tmp67968 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp67969 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp67970 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp67971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp67972 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp67973 := Call(__e, tmp67972, V442)

													tmp67974 := Call(__e, tmp67971, tmp67973)

													tmp67975 := Call(__e, tmp67970, tmp67974)

													tmp67976 := Call(__e, tmp67969, tmp67975)

													tmp67977 := Call(__e, tmp67968, tmp67976)

													var ifres67944 Obj

													if True == tmp67977 {
														tmp67956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp67957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp67958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp67959 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp67960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp67961 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp67962 := Call(__e, tmp67961, V442)

														tmp67963 := Call(__e, tmp67960, tmp67962)

														tmp67964 := Call(__e, tmp67959, tmp67963)

														tmp67965 := Call(__e, tmp67958, tmp67964)

														tmp67966 := Call(__e, tmp67957, tmp67965)

														tmp67967 := Call(__e, tmp67956, Nil, tmp67966)

														var ifres67946 Obj

														if True == tmp67967 {
															tmp67948 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp67949 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp67950 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp67951 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp67952 := Call(__e, tmp67951, V442)

															tmp67953 := Call(__e, tmp67950, tmp67952)

															tmp67954 := Call(__e, tmp67949, tmp67953)

															tmp67955 := Call(__e, tmp67948, Nil, tmp67954)

															var ifres67947 Obj

															if True == tmp67955 {
																ifres67947 = True

															} else {
																ifres67947 = False

															}

															ifres67946 = ifres67947

														} else {
															ifres67946 = False

														}

														var ifres67945 Obj

														if True == ifres67946 {
															ifres67945 = True

														} else {
															ifres67945 = False

														}

														ifres67944 = ifres67945

													} else {
														ifres67944 = False

													}

													var ifres67943 Obj

													if True == ifres67944 {
														ifres67943 = True

													} else {
														ifres67943 = False

													}

													ifres67942 = ifres67943

												} else {
													ifres67942 = False

												}

												var ifres67941 Obj

												if True == ifres67942 {
													ifres67941 = True

												} else {
													ifres67941 = False

												}

												ifres67940 = ifres67941

											} else {
												ifres67940 = False

											}

											var ifres67939 Obj

											if True == ifres67940 {
												ifres67939 = True

											} else {
												ifres67939 = False

											}

											ifres67938 = ifres67939

										} else {
											ifres67938 = False

										}

										var ifres67937 Obj

										if True == ifres67938 {
											ifres67937 = True

										} else {
											ifres67937 = False

										}

										ifres67936 = ifres67937

									} else {
										ifres67936 = False

									}

									var ifres67935 Obj

									if True == ifres67936 {
										ifres67935 = True

									} else {
										ifres67935 = False

									}

									ifres67934 = ifres67935

								} else {
									ifres67934 = False

								}

								var ifres67933 Obj

								if True == ifres67934 {
									ifres67933 = True

								} else {
									ifres67933 = False

								}

								ifres67932 = ifres67933

							} else {
								ifres67932 = False

							}

							var ifres67931 Obj

							if True == ifres67932 {
								ifres67931 = True

							} else {
								ifres67931 = False

							}

							ifres67930 = ifres67931

						} else {
							ifres67930 = False

						}

						var ifres67929 Obj

						if True == ifres67930 {
							ifres67929 = True

						} else {
							ifres67929 = False

						}

						ifres67928 = ifres67929

					} else {
						ifres67928 = False

					}

					var ifres67927 Obj

					if True == ifres67928 {
						ifres67927 = True

					} else {
						ifres67927 = False

					}

					ifres67926 = ifres67927

				} else {
					ifres67926 = False

				}

				var ifres67925 Obj

				if True == ifres67926 {
					ifres67925 = True

				} else {
					ifres67925 = False

				}

				ifres67924 = ifres67925

			} else {
				ifres67924 = False

			}

			if True == ifres67924 {
				tmp67913 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67914 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp67915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp67916 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp67917 := Call(__e, tmp67916, V442)

				tmp67918 := Call(__e, tmp67915, tmp67917)

				tmp67919 := Call(__e, tmp67914, True, tmp67918)

				tmp67920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4case_1form)

				tmp67921 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp67922 := Call(__e, tmp67921, V442)

				tmp67923 := Call(__e, tmp67920, tmp67922, V443)

				__e.TailApply(tmp67913, tmp67919, tmp67923)
				return

			} else {
				tmp67911 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp67912 := Call(__e, tmp67911, V442)

				var ifres67835 Obj

				if True == tmp67912 {
					tmp67907 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67908 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67909 := Call(__e, tmp67908, V442)

					tmp67910 := Call(__e, tmp67907, tmp67909)

					var ifres67837 Obj

					if True == tmp67910 {
						tmp67901 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67902 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67903 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67904 := Call(__e, tmp67903, V442)

						tmp67905 := Call(__e, tmp67902, tmp67904)

						tmp67906 := Call(__e, tmp67901, tmp67905)

						var ifres67839 Obj

						if True == tmp67906 {
							tmp67893 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp67894 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67895 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67897 := Call(__e, tmp67896, V442)

							tmp67898 := Call(__e, tmp67895, tmp67897)

							tmp67899 := Call(__e, tmp67894, tmp67898)

							tmp67900 := Call(__e, tmp67893, sym_h, tmp67899)

							var ifres67841 Obj

							if True == tmp67900 {
								tmp67885 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp67886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp67887 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67888 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67889 := Call(__e, tmp67888, V442)

								tmp67890 := Call(__e, tmp67887, tmp67889)

								tmp67891 := Call(__e, tmp67886, tmp67890)

								tmp67892 := Call(__e, tmp67885, tmp67891)

								var ifres67843 Obj

								if True == tmp67892 {
									tmp67875 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp67876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67878 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67879 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67880 := Call(__e, tmp67879, V442)

									tmp67881 := Call(__e, tmp67878, tmp67880)

									tmp67882 := Call(__e, tmp67877, tmp67881)

									tmp67883 := Call(__e, tmp67876, tmp67882)

									tmp67884 := Call(__e, tmp67875, symshen_4tests, tmp67883)

									var ifres67845 Obj

									if True == tmp67884 {
										tmp67865 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp67866 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67868 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67869 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67870 := Call(__e, tmp67869, V442)

										tmp67871 := Call(__e, tmp67868, tmp67870)

										tmp67872 := Call(__e, tmp67867, tmp67871)

										tmp67873 := Call(__e, tmp67866, tmp67872)

										tmp67874 := Call(__e, tmp67865, Nil, tmp67873)

										var ifres67847 Obj

										if True == tmp67874 {
											tmp67859 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp67860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67861 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67862 := Call(__e, tmp67861, V442)

											tmp67863 := Call(__e, tmp67860, tmp67862)

											tmp67864 := Call(__e, tmp67859, tmp67863)

											var ifres67849 Obj

											if True == tmp67864 {
												tmp67851 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67854 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67855 := Call(__e, tmp67854, V442)

												tmp67856 := Call(__e, tmp67853, tmp67855)

												tmp67857 := Call(__e, tmp67852, tmp67856)

												tmp67858 := Call(__e, tmp67851, Nil, tmp67857)

												var ifres67850 Obj

												if True == tmp67858 {
													ifres67850 = True

												} else {
													ifres67850 = False

												}

												ifres67849 = ifres67850

											} else {
												ifres67849 = False

											}

											var ifres67848 Obj

											if True == ifres67849 {
												ifres67848 = True

											} else {
												ifres67848 = False

											}

											ifres67847 = ifres67848

										} else {
											ifres67847 = False

										}

										var ifres67846 Obj

										if True == ifres67847 {
											ifres67846 = True

										} else {
											ifres67846 = False

										}

										ifres67845 = ifres67846

									} else {
										ifres67845 = False

									}

									var ifres67844 Obj

									if True == ifres67845 {
										ifres67844 = True

									} else {
										ifres67844 = False

									}

									ifres67843 = ifres67844

								} else {
									ifres67843 = False

								}

								var ifres67842 Obj

								if True == ifres67843 {
									ifres67842 = True

								} else {
									ifres67842 = False

								}

								ifres67841 = ifres67842

							} else {
								ifres67841 = False

							}

							var ifres67840 Obj

							if True == ifres67841 {
								ifres67840 = True

							} else {
								ifres67840 = False

							}

							ifres67839 = ifres67840

						} else {
							ifres67839 = False

						}

						var ifres67838 Obj

						if True == ifres67839 {
							ifres67838 = True

						} else {
							ifres67838 = False

						}

						ifres67837 = ifres67838

					} else {
						ifres67837 = False

					}

					var ifres67836 Obj

					if True == ifres67837 {
						ifres67836 = True

					} else {
						ifres67836 = False

					}

					ifres67835 = ifres67836

				} else {
					ifres67835 = False

				}

				if True == ifres67835 {
					tmp67828 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67829 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp67830 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp67831 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp67832 := Call(__e, tmp67831, V442)

					tmp67833 := Call(__e, tmp67830, tmp67832)

					tmp67834 := Call(__e, tmp67829, True, tmp67833)

					__e.TailApply(tmp67828, tmp67834, Nil)
					return

				} else {
					tmp67826 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp67827 := Call(__e, tmp67826, V442)

					var ifres67762 Obj

					if True == tmp67827 {
						tmp67822 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp67823 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67824 := Call(__e, tmp67823, V442)

						tmp67825 := Call(__e, tmp67822, tmp67824)

						var ifres67764 Obj

						if True == tmp67825 {
							tmp67816 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp67817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67818 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp67819 := Call(__e, tmp67818, V442)

							tmp67820 := Call(__e, tmp67817, tmp67819)

							tmp67821 := Call(__e, tmp67816, tmp67820)

							var ifres67766 Obj

							if True == tmp67821 {
								tmp67808 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp67809 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67810 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67811 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp67812 := Call(__e, tmp67811, V442)

								tmp67813 := Call(__e, tmp67810, tmp67812)

								tmp67814 := Call(__e, tmp67809, tmp67813)

								tmp67815 := Call(__e, tmp67808, sym_h, tmp67814)

								var ifres67768 Obj

								if True == tmp67815 {
									tmp67800 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp67801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp67802 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67803 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp67804 := Call(__e, tmp67803, V442)

									tmp67805 := Call(__e, tmp67802, tmp67804)

									tmp67806 := Call(__e, tmp67801, tmp67805)

									tmp67807 := Call(__e, tmp67800, tmp67806)

									var ifres67770 Obj

									if True == tmp67807 {
										tmp67790 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp67791 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67792 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp67793 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp67795 := Call(__e, tmp67794, V442)

										tmp67796 := Call(__e, tmp67793, tmp67795)

										tmp67797 := Call(__e, tmp67792, tmp67796)

										tmp67798 := Call(__e, tmp67791, tmp67797)

										tmp67799 := Call(__e, tmp67790, symshen_4tests, tmp67798)

										var ifres67772 Obj

										if True == tmp67799 {
											tmp67784 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp67785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp67786 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp67787 := Call(__e, tmp67786, V442)

											tmp67788 := Call(__e, tmp67785, tmp67787)

											tmp67789 := Call(__e, tmp67784, tmp67788)

											var ifres67774 Obj

											if True == tmp67789 {
												tmp67776 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp67777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp67779 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp67780 := Call(__e, tmp67779, V442)

												tmp67781 := Call(__e, tmp67778, tmp67780)

												tmp67782 := Call(__e, tmp67777, tmp67781)

												tmp67783 := Call(__e, tmp67776, Nil, tmp67782)

												var ifres67775 Obj

												if True == tmp67783 {
													ifres67775 = True

												} else {
													ifres67775 = False

												}

												ifres67774 = ifres67775

											} else {
												ifres67774 = False

											}

											var ifres67773 Obj

											if True == ifres67774 {
												ifres67773 = True

											} else {
												ifres67773 = False

											}

											ifres67772 = ifres67773

										} else {
											ifres67772 = False

										}

										var ifres67771 Obj

										if True == ifres67772 {
											ifres67771 = True

										} else {
											ifres67771 = False

										}

										ifres67770 = ifres67771

									} else {
										ifres67770 = False

									}

									var ifres67769 Obj

									if True == ifres67770 {
										ifres67769 = True

									} else {
										ifres67769 = False

									}

									ifres67768 = ifres67769

								} else {
									ifres67768 = False

								}

								var ifres67767 Obj

								if True == ifres67768 {
									ifres67767 = True

								} else {
									ifres67767 = False

								}

								ifres67766 = ifres67767

							} else {
								ifres67766 = False

							}

							var ifres67765 Obj

							if True == ifres67766 {
								ifres67765 = True

							} else {
								ifres67765 = False

							}

							ifres67764 = ifres67765

						} else {
							ifres67764 = False

						}

						var ifres67763 Obj

						if True == ifres67764 {
							ifres67763 = True

						} else {
							ifres67763 = False

						}

						ifres67762 = ifres67763

					} else {
						ifres67762 = False

					}

					if True == ifres67762 {
						tmp67741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67742 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp67743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4embed_1and)

						tmp67744 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67746 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67747 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67748 := Call(__e, tmp67747, V442)

						tmp67749 := Call(__e, tmp67746, tmp67748)

						tmp67750 := Call(__e, tmp67745, tmp67749)

						tmp67751 := Call(__e, tmp67744, tmp67750)

						tmp67752 := Call(__e, tmp67743, tmp67751)

						tmp67753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67754 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp67755 := Call(__e, tmp67754, V442)

						tmp67756 := Call(__e, tmp67753, tmp67755)

						tmp67757 := Call(__e, tmp67742, tmp67752, tmp67756)

						tmp67758 := Call(__e, PrimNS1Value(symns2_1value), symshen_4case_1form)

						tmp67759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp67760 := Call(__e, tmp67759, V442)

						tmp67761 := Call(__e, tmp67758, tmp67760, V443)

						__e.TailApply(tmp67741, tmp67757, tmp67761)
						return

					} else {
						tmp67740 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp67740, symshen_4case_1form)
						return

					}

				}

			}

		}

	}, 2)

	tmp68053 := Call(__e, PrimNS1Value(symns2_1set), symshen_4case_1form, tmp67735)

	_ = tmp68053

	tmp68054 := MakeNative(func(__e *ControlFlow) {
		V445 := __e.Get(1)
		_ = V445
		tmp68078 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp68079 := Call(__e, tmp68078, V445)

		var ifres68072 Obj

		if True == tmp68079 {
			tmp68074 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp68075 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp68076 := Call(__e, tmp68075, V445)

			tmp68077 := Call(__e, tmp68074, Nil, tmp68076)

			var ifres68073 Obj

			if True == tmp68077 {
				ifres68073 = True

			} else {
				ifres68073 = False

			}

			ifres68072 = ifres68073

		} else {
			ifres68072 = False

		}

		if True == ifres68072 {
			tmp68071 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp68071, V445)
			return

		} else {
			tmp68069 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp68070 := Call(__e, tmp68069, V445)

			if True == tmp68070 {
				tmp68058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68059 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68060 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp68061 := Call(__e, tmp68060, V445)

				tmp68062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp68063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4embed_1and)

				tmp68064 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp68065 := Call(__e, tmp68064, V445)

				tmp68066 := Call(__e, tmp68063, tmp68065)

				tmp68067 := Call(__e, tmp68062, tmp68066, Nil)

				tmp68068 := Call(__e, tmp68059, tmp68061, tmp68067)

				__e.TailApply(tmp68058, symand, tmp68068)
				return

			} else {
				tmp68057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp68057, symshen_4embed_1and)
				return

			}

		}

	}, 1)

	tmp68080 := Call(__e, PrimNS1Value(symns2_1set), symshen_4embed_1and, tmp68054)

	_ = tmp68080

	tmp68081 := MakeNative(func(__e *ControlFlow) {
		V447 := __e.Get(1)
		_ = V447
		tmp68082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp68083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp68084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp68085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp68086 := Call(__e, tmp68085, V447, Nil)

		tmp68087 := Call(__e, tmp68084, symshen_4f__error, tmp68086)

		tmp68088 := Call(__e, tmp68083, tmp68087, Nil)

		__e.TailApply(tmp68082, True, tmp68088)
		return

	}, 1)

	tmp68089 := Call(__e, PrimNS1Value(symns2_1set), symshen_4err_1condition, tmp68081)

	_ = tmp68089

	tmp68090 := MakeNative(func(__e *ControlFlow) {
		V449 := __e.Get(1)
		_ = V449
		tmp68091 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		tmp68092 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp68093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp68094 := Call(__e, tmp68093, V449, MakeString(": unexpected argument\n"), symshen_4a)

		tmp68095 := Call(__e, tmp68092, MakeString("system function "), tmp68094)

		__e.TailApply(tmp68091, tmp68095)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4sys_1error, tmp68090)
	return

}, 0)
