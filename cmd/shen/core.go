package main

import . "github.com/tiancaiamao/shen-go/cora"

var CoreMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp86 := MakeNative(func(__e *ControlFlow) {
		V152 := __e.Get(1)
		_ = V152
		V153 := __e.Get(2)
		_ = V153
		tmp87 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5define_6), X)
			return
		}, 1)

		tmp88 := PrimCons(V152, V153)

		tmp89 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4shen_1syntax_1error), V152, X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symcompile), tmp87, tmp88, tmp89)
		return

	}, 2)

	tmp90 := Call(__e, PrimNS1Value(symns2_1set), symshen_4shen_1_6kl, tmp86)

	_ = tmp90

	tmp91 := MakeNative(func(__e *ControlFlow) {
		V160 := __e.Get(1)
		_ = V160
		V161 := __e.Get(2)
		_ = V161
		tmp101 := PrimIsPair(V161)

		if True == tmp101 {
			tmp95 := PrimHead(V161)

			tmp96 := Call(__e, PrimNS2Value(symshen_4next_150), MakeNumber(50), tmp95)

			tmp97 := Call(__e, PrimNS2Value(symshen_4app), tmp96, MakeString("\n"), symshen_4a)

			tmp98 := PrimStringConcat(MakeString(" here:\n\n "), tmp97)

			tmp99 := Call(__e, PrimNS2Value(symshen_4app), V160, tmp98, symshen_4a)

			tmp100 := PrimStringConcat(MakeString("syntax error in "), tmp99)

			__e.Return(PrimSimpleError(tmp100))
			return

		} else {
			tmp93 := Call(__e, PrimNS2Value(symshen_4app), V160, MakeString("\n"), symshen_4a)

			tmp94 := PrimStringConcat(MakeString("syntax error in "), tmp93)

			__e.Return(PrimSimpleError(tmp94))
			return

		}

	}, 2)

	tmp102 := Call(__e, PrimNS1Value(symns2_1set), symshen_4shen_1syntax_1error, tmp91)

	_ = tmp102

	tmp103 := MakeNative(func(__e *ControlFlow) {
		V163 := __e.Get(1)
		_ = V163
		tmp104 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp122 := Call(__e, PrimNS2Value(symfail))

			tmp123 := PrimEqual(YaccParse, tmp122)

			if True == tmp123 {
				tmp106 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5name_6 := __e.Get(1)
					_ = Parse__shen_4_5name_6
					tmp118 := Call(__e, PrimNS2Value(symfail))

					tmp119 := PrimEqual(tmp118, Parse__shen_4_5name_6)

					tmp120 := PrimNot(tmp119)

					if True == tmp120 {
						tmp108 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rules_6 := __e.Get(1)
							_ = Parse__shen_4_5rules_6
							tmp114 := Call(__e, PrimNS2Value(symfail))

							tmp115 := PrimEqual(tmp114, Parse__shen_4_5rules_6)

							tmp116 := PrimNot(tmp115)

							if True == tmp116 {
								tmp110 := PrimHead(Parse__shen_4_5rules_6)

								tmp111 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5name_6)

								tmp112 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rules_6)

								tmp113 := Call(__e, PrimNS2Value(symshen_4compile__to__machine__code), tmp111, tmp112)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp110, tmp113)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp117 := Call(__e, PrimNS2Value(symshen_4_5rules_6), Parse__shen_4_5name_6)

						__e.TailApply(tmp108, tmp117)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp121 := Call(__e, PrimNS2Value(symshen_4_5name_6), V163)

				__e.TailApply(tmp106, tmp121)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp124 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5name_6 := __e.Get(1)
			_ = Parse__shen_4_5name_6
			tmp142 := Call(__e, PrimNS2Value(symfail))

			tmp143 := PrimEqual(tmp142, Parse__shen_4_5name_6)

			tmp144 := PrimNot(tmp143)

			if True == tmp144 {
				tmp126 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_6
					tmp138 := Call(__e, PrimNS2Value(symfail))

					tmp139 := PrimEqual(tmp138, Parse__shen_4_5signature_6)

					tmp140 := PrimNot(tmp139)

					if True == tmp140 {
						tmp128 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5rules_6 := __e.Get(1)
							_ = Parse__shen_4_5rules_6
							tmp134 := Call(__e, PrimNS2Value(symfail))

							tmp135 := PrimEqual(tmp134, Parse__shen_4_5rules_6)

							tmp136 := PrimNot(tmp135)

							if True == tmp136 {
								tmp130 := PrimHead(Parse__shen_4_5rules_6)

								tmp131 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5name_6)

								tmp132 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rules_6)

								tmp133 := Call(__e, PrimNS2Value(symshen_4compile__to__machine__code), tmp131, tmp132)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp130, tmp133)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp137 := Call(__e, PrimNS2Value(symshen_4_5rules_6), Parse__shen_4_5signature_6)

						__e.TailApply(tmp128, tmp137)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp141 := Call(__e, PrimNS2Value(symshen_4_5signature_6), Parse__shen_4_5name_6)

				__e.TailApply(tmp126, tmp141)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp145 := Call(__e, PrimNS2Value(symshen_4_5name_6), V163)

		tmp146 := Call(__e, tmp124, tmp145)

		__e.TailApply(tmp104, tmp146)
		return

	}, 1)

	tmp147 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5define_6, tmp103)

	_ = tmp147

	tmp148 := MakeNative(func(__e *ControlFlow) {
		V165 := __e.Get(1)
		_ = V165
		tmp164 := PrimHead(V165)

		tmp165 := PrimIsPair(tmp164)

		if True == tmp165 {
			tmp150 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp151 := Call(__e, PrimNS2Value(symshen_4tlhd), V165)

				tmp152 := Call(__e, PrimNS2Value(symshen_4hdtl), V165)

				tmp153 := Call(__e, PrimNS2Value(symshen_4pair), tmp151, tmp152)

				tmp154 := PrimHead(tmp153)

				tmp162 := PrimIsSymbol(Parse__X)

				var ifres158 Obj

				if True == tmp162 {
					tmp160 := Call(__e, PrimNS2Value(symshen_4sysfunc_2), Parse__X)

					tmp161 := PrimNot(tmp160)

					var ifres159 Obj

					if True == tmp161 {
						ifres159 = True

					} else {
						ifres159 = False

					}

					ifres158 = ifres159

				} else {
					ifres158 = False

				}

				var ifres155 Obj

				if True == ifres158 {
					ifres155 = Parse__X

				} else {
					tmp156 := Call(__e, PrimNS2Value(symshen_4app), Parse__X, MakeString(" is not a legitimate function name.\n"), symshen_4a)

					tmp157 := PrimSimpleError(tmp156)

					ifres155 = tmp157

				}

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp154, ifres155)
				return

			}, 1)

			tmp163 := Call(__e, PrimNS2Value(symshen_4hdhd), V165)

			__e.TailApply(tmp150, tmp163)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp166 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5name_6, tmp148)

	_ = tmp166

	tmp167 := MakeNative(func(__e *ControlFlow) {
		V167 := __e.Get(1)
		_ = V167
		tmp168 := PrimIntern(MakeString("shen"))

		tmp169 := PrimNS3Value(sym_dproperty_1vector_d)

		tmp170 := Call(__e, PrimNS2Value(symget), tmp168, symshen_4external_1symbols, tmp169)

		__e.TailApply(PrimNS2Value(symelement_2), V167, tmp170)
		return

	}, 1)

	tmp171 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sysfunc_2, tmp167)

	_ = tmp171

	tmp172 := MakeNative(func(__e *ControlFlow) {
		V171 := __e.Get(1)
		_ = V171
		tmp203 := PrimHead(V171)

		tmp204 := PrimIsPair(tmp203)

		var ifres199 Obj

		if True == tmp204 {
			tmp201 := Call(__e, PrimNS2Value(symshen_4hdhd), V171)

			tmp202 := PrimEqual(sym_i, tmp201)

			var ifres200 Obj

			if True == tmp202 {
				ifres200 = True

			} else {
				ifres200 = False

			}

			ifres199 = ifres200

		} else {
			ifres199 = False

		}

		if True == ifres199 {
			tmp174 := MakeNative(func(__e *ControlFlow) {
				NewStream168 := __e.Get(1)
				_ = NewStream168
				tmp175 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_1help_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_1help_6
					tmp192 := Call(__e, PrimNS2Value(symfail))

					tmp193 := PrimEqual(tmp192, Parse__shen_4_5signature_1help_6)

					tmp194 := PrimNot(tmp193)

					if True == tmp194 {
						tmp190 := PrimHead(Parse__shen_4_5signature_1help_6)

						tmp191 := PrimIsPair(tmp190)

						var ifres186 Obj

						if True == tmp191 {
							tmp188 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5signature_1help_6)

							tmp189 := PrimEqual(sym_j, tmp188)

							var ifres187 Obj

							if True == tmp189 {
								ifres187 = True

							} else {
								ifres187 = False

							}

							ifres186 = ifres187

						} else {
							ifres186 = False

						}

						if True == ifres186 {
							tmp178 := MakeNative(func(__e *ControlFlow) {
								NewStream169 := __e.Get(1)
								_ = NewStream169
								tmp179 := PrimHead(NewStream169)

								tmp180 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

								tmp181 := Call(__e, PrimNS2Value(symshen_4curry_1type), tmp180)

								tmp182 := Call(__e, PrimNS2Value(symshen_4demodulate), tmp181)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp179, tmp182)
								return

							}, 1)

							tmp183 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5signature_1help_6)

							tmp184 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

							tmp185 := Call(__e, PrimNS2Value(symshen_4pair), tmp183, tmp184)

							__e.TailApply(tmp178, tmp185)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp195 := Call(__e, PrimNS2Value(symshen_4_5signature_1help_6), NewStream168)

				__e.TailApply(tmp175, tmp195)
				return

			}, 1)

			tmp196 := Call(__e, PrimNS2Value(symshen_4tlhd), V171)

			tmp197 := Call(__e, PrimNS2Value(symshen_4hdtl), V171)

			tmp198 := Call(__e, PrimNS2Value(symshen_4pair), tmp196, tmp197)

			__e.TailApply(tmp174, tmp198)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp205 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5signature_6, tmp172)

	_ = tmp205

	tmp206 := MakeNative(func(__e *ControlFlow) {
		V173 := __e.Get(1)
		_ = V173
		tmp207 := Call(__e, PrimNS2Value(symshen_4curry_1type_1h), V173)

		__e.TailApply(PrimNS2Value(symshen_4active_1cons), tmp207)
		return

	}, 1)

	tmp208 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1type, tmp206)

	_ = tmp208

	tmp209 := MakeNative(func(__e *ControlFlow) {
		V175 := __e.Get(1)
		_ = V175
		tmp243 := PrimIsPair(V175)

		var ifres223 Obj

		if True == tmp243 {
			tmp241 := PrimTail(V175)

			tmp242 := PrimIsPair(tmp241)

			var ifres225 Obj

			if True == tmp242 {
				tmp238 := PrimTail(V175)

				tmp239 := PrimTail(tmp238)

				tmp240 := PrimIsPair(tmp239)

				var ifres227 Obj

				if True == tmp240 {
					tmp234 := PrimTail(V175)

					tmp235 := PrimTail(tmp234)

					tmp236 := PrimTail(tmp235)

					tmp237 := PrimEqual(Nil, tmp236)

					var ifres229 Obj

					if True == tmp237 {
						tmp231 := PrimTail(V175)

						tmp232 := PrimHead(tmp231)

						tmp233 := PrimEqual(tmp232, symbar_b)

						var ifres230 Obj

						if True == tmp233 {
							ifres230 = True

						} else {
							ifres230 = False

						}

						ifres229 = ifres230

					} else {
						ifres229 = False

					}

					var ifres228 Obj

					if True == ifres229 {
						ifres228 = True

					} else {
						ifres228 = False

					}

					ifres227 = ifres228

				} else {
					ifres227 = False

				}

				var ifres226 Obj

				if True == ifres227 {
					ifres226 = True

				} else {
					ifres226 = False

				}

				ifres225 = ifres226

			} else {
				ifres225 = False

			}

			var ifres224 Obj

			if True == ifres225 {
				ifres224 = True

			} else {
				ifres224 = False

			}

			ifres223 = ifres224

		} else {
			ifres223 = False

		}

		if True == ifres223 {
			tmp217 := PrimHead(V175)

			tmp218 := Call(__e, PrimNS2Value(symshen_4active_1cons), tmp217)

			tmp219 := PrimTail(V175)

			tmp220 := PrimTail(tmp219)

			tmp221 := PrimHead(tmp220)

			tmp222 := Call(__e, PrimNS2Value(symshen_4active_1cons), tmp221)

			__e.Return(PrimCons(tmp218, tmp222))
			return

		} else {
			tmp216 := PrimIsPair(V175)

			if True == tmp216 {
				tmp212 := PrimHead(V175)

				tmp213 := Call(__e, PrimNS2Value(symshen_4active_1cons), tmp212)

				tmp214 := PrimTail(V175)

				tmp215 := Call(__e, PrimNS2Value(symshen_4active_1cons), tmp214)

				__e.Return(PrimCons(tmp213, tmp215))
				return

			} else {
				__e.Return(V175)
				return
			}

		}

	}, 1)

	tmp244 := Call(__e, PrimNS1Value(symns2_1set), symshen_4active_1cons, tmp209)

	_ = tmp244

	tmp245 := MakeNative(func(__e *ControlFlow) {
		V177 := __e.Get(1)
		_ = V177
		tmp318 := PrimIsPair(V177)

		var ifres291 Obj

		if True == tmp318 {
			tmp316 := PrimTail(V177)

			tmp317 := PrimIsPair(tmp316)

			var ifres293 Obj

			if True == tmp317 {
				tmp313 := PrimTail(V177)

				tmp314 := PrimHead(tmp313)

				tmp315 := PrimEqual(sym_1_1_6, tmp314)

				var ifres295 Obj

				if True == tmp315 {
					tmp310 := PrimTail(V177)

					tmp311 := PrimTail(tmp310)

					tmp312 := PrimIsPair(tmp311)

					var ifres297 Obj

					if True == tmp312 {
						tmp306 := PrimTail(V177)

						tmp307 := PrimTail(tmp306)

						tmp308 := PrimTail(tmp307)

						tmp309 := PrimIsPair(tmp308)

						var ifres299 Obj

						if True == tmp309 {
							tmp301 := PrimTail(V177)

							tmp302 := PrimTail(tmp301)

							tmp303 := PrimTail(tmp302)

							tmp304 := PrimHead(tmp303)

							tmp305 := PrimEqual(sym_1_1_6, tmp304)

							var ifres300 Obj

							if True == tmp305 {
								ifres300 = True

							} else {
								ifres300 = False

							}

							ifres299 = ifres300

						} else {
							ifres299 = False

						}

						var ifres298 Obj

						if True == ifres299 {
							ifres298 = True

						} else {
							ifres298 = False

						}

						ifres297 = ifres298

					} else {
						ifres297 = False

					}

					var ifres296 Obj

					if True == ifres297 {
						ifres296 = True

					} else {
						ifres296 = False

					}

					ifres295 = ifres296

				} else {
					ifres295 = False

				}

				var ifres294 Obj

				if True == ifres295 {
					ifres294 = True

				} else {
					ifres294 = False

				}

				ifres293 = ifres294

			} else {
				ifres293 = False

			}

			var ifres292 Obj

			if True == ifres293 {
				ifres292 = True

			} else {
				ifres292 = False

			}

			ifres291 = ifres292

		} else {
			ifres291 = False

		}

		if True == ifres291 {
			tmp285 := PrimHead(V177)

			tmp286 := PrimTail(V177)

			tmp287 := PrimTail(tmp286)

			tmp288 := PrimCons(tmp287, Nil)

			tmp289 := PrimCons(sym_1_1_6, tmp288)

			tmp290 := PrimCons(tmp285, tmp289)

			__e.TailApply(PrimNS2Value(symshen_4curry_1type_1h), tmp290)
			return

		} else {
			tmp284 := PrimIsPair(V177)

			var ifres257 Obj

			if True == tmp284 {
				tmp282 := PrimTail(V177)

				tmp283 := PrimIsPair(tmp282)

				var ifres259 Obj

				if True == tmp283 {
					tmp279 := PrimTail(V177)

					tmp280 := PrimHead(tmp279)

					tmp281 := PrimEqual(sym_d, tmp280)

					var ifres261 Obj

					if True == tmp281 {
						tmp276 := PrimTail(V177)

						tmp277 := PrimTail(tmp276)

						tmp278 := PrimIsPair(tmp277)

						var ifres263 Obj

						if True == tmp278 {
							tmp272 := PrimTail(V177)

							tmp273 := PrimTail(tmp272)

							tmp274 := PrimTail(tmp273)

							tmp275 := PrimIsPair(tmp274)

							var ifres265 Obj

							if True == tmp275 {
								tmp267 := PrimTail(V177)

								tmp268 := PrimTail(tmp267)

								tmp269 := PrimTail(tmp268)

								tmp270 := PrimHead(tmp269)

								tmp271 := PrimEqual(sym_d, tmp270)

								var ifres266 Obj

								if True == tmp271 {
									ifres266 = True

								} else {
									ifres266 = False

								}

								ifres265 = ifres266

							} else {
								ifres265 = False

							}

							var ifres264 Obj

							if True == ifres265 {
								ifres264 = True

							} else {
								ifres264 = False

							}

							ifres263 = ifres264

						} else {
							ifres263 = False

						}

						var ifres262 Obj

						if True == ifres263 {
							ifres262 = True

						} else {
							ifres262 = False

						}

						ifres261 = ifres262

					} else {
						ifres261 = False

					}

					var ifres260 Obj

					if True == ifres261 {
						ifres260 = True

					} else {
						ifres260 = False

					}

					ifres259 = ifres260

				} else {
					ifres259 = False

				}

				var ifres258 Obj

				if True == ifres259 {
					ifres258 = True

				} else {
					ifres258 = False

				}

				ifres257 = ifres258

			} else {
				ifres257 = False

			}

			if True == ifres257 {
				tmp251 := PrimHead(V177)

				tmp252 := PrimTail(V177)

				tmp253 := PrimTail(tmp252)

				tmp254 := PrimCons(tmp253, Nil)

				tmp255 := PrimCons(sym_d, tmp254)

				tmp256 := PrimCons(tmp251, tmp255)

				__e.TailApply(PrimNS2Value(symshen_4curry_1type_1h), tmp256)
				return

			} else {
				tmp250 := PrimIsPair(V177)

				if True == tmp250 {
					tmp249 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimNS2Value(symshen_4curry_1type_1h), Z)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp249, V177)
					return

				} else {
					__e.Return(V177)
					return
				}

			}

		}

	}, 1)

	tmp319 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1type_1h, tmp245)

	_ = tmp319

	tmp320 := MakeNative(func(__e *ControlFlow) {
		V179 := __e.Get(1)
		_ = V179
		tmp321 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp330 := Call(__e, PrimNS2Value(symfail))

			tmp331 := PrimEqual(YaccParse, tmp330)

			if True == tmp331 {
				tmp323 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp326 := Call(__e, PrimNS2Value(symfail))

					tmp327 := PrimEqual(tmp326, Parse___5e_6)

					tmp328 := PrimNot(tmp327)

					if True == tmp328 {
						tmp325 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp325, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp329 := Call(__e, PrimNS2Value(sym_5e_6), V179)

				__e.TailApply(tmp323, tmp329)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp354 := PrimHead(V179)

		tmp355 := PrimIsPair(tmp354)

		var ifres332 Obj

		if True == tmp355 {
			tmp334 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp335 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5signature_1help_6 := __e.Get(1)
					_ = Parse__shen_4_5signature_1help_6
					tmp345 := Call(__e, PrimNS2Value(symfail))

					tmp346 := PrimEqual(tmp345, Parse__shen_4_5signature_1help_6)

					tmp347 := PrimNot(tmp346)

					if True == tmp347 {
						tmp341 := PrimCons(sym_j, Nil)

						tmp342 := PrimCons(sym_i, tmp341)

						tmp343 := Call(__e, PrimNS2Value(symelement_2), Parse__X, tmp342)

						tmp344 := PrimNot(tmp343)

						if True == tmp344 {
							tmp338 := PrimHead(Parse__shen_4_5signature_1help_6)

							tmp339 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

							tmp340 := PrimCons(Parse__X, tmp339)

							__e.TailApply(PrimNS2Value(symshen_4pair), tmp338, tmp340)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp348 := Call(__e, PrimNS2Value(symshen_4tlhd), V179)

				tmp349 := Call(__e, PrimNS2Value(symshen_4hdtl), V179)

				tmp350 := Call(__e, PrimNS2Value(symshen_4pair), tmp348, tmp349)

				tmp351 := Call(__e, PrimNS2Value(symshen_4_5signature_1help_6), tmp350)

				__e.TailApply(tmp335, tmp351)
				return

			}, 1)

			tmp352 := Call(__e, PrimNS2Value(symshen_4hdhd), V179)

			tmp353 := Call(__e, tmp334, tmp352)

			ifres332 = tmp353

		} else {
			tmp333 := Call(__e, PrimNS2Value(symfail))

			ifres332 = tmp333

		}

		__e.TailApply(tmp321, ifres332)
		return

	}, 1)

	tmp356 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5signature_1help_6, tmp320)

	_ = tmp356

	tmp357 := MakeNative(func(__e *ControlFlow) {
		V181 := __e.Get(1)
		_ = V181
		tmp358 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp370 := Call(__e, PrimNS2Value(symfail))

			tmp371 := PrimEqual(YaccParse, tmp370)

			if True == tmp371 {
				tmp360 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rule_6 := __e.Get(1)
					_ = Parse__shen_4_5rule_6
					tmp366 := Call(__e, PrimNS2Value(symfail))

					tmp367 := PrimEqual(tmp366, Parse__shen_4_5rule_6)

					tmp368 := PrimNot(tmp367)

					if True == tmp368 {
						tmp362 := PrimHead(Parse__shen_4_5rule_6)

						tmp363 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rule_6)

						tmp364 := Call(__e, PrimNS2Value(symshen_4linearise), tmp363)

						tmp365 := PrimCons(tmp364, Nil)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp362, tmp365)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp369 := Call(__e, PrimNS2Value(symshen_4_5rule_6), V181)

				__e.TailApply(tmp360, tmp369)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp372 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5rule_6 := __e.Get(1)
			_ = Parse__shen_4_5rule_6
			tmp385 := Call(__e, PrimNS2Value(symfail))

			tmp386 := PrimEqual(tmp385, Parse__shen_4_5rule_6)

			tmp387 := PrimNot(tmp386)

			if True == tmp387 {
				tmp374 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rules_6 := __e.Get(1)
					_ = Parse__shen_4_5rules_6
					tmp381 := Call(__e, PrimNS2Value(symfail))

					tmp382 := PrimEqual(tmp381, Parse__shen_4_5rules_6)

					tmp383 := PrimNot(tmp382)

					if True == tmp383 {
						tmp376 := PrimHead(Parse__shen_4_5rules_6)

						tmp377 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rule_6)

						tmp378 := Call(__e, PrimNS2Value(symshen_4linearise), tmp377)

						tmp379 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rules_6)

						tmp380 := PrimCons(tmp378, tmp379)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp376, tmp380)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp384 := Call(__e, PrimNS2Value(symshen_4_5rules_6), Parse__shen_4_5rule_6)

				__e.TailApply(tmp374, tmp384)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp388 := Call(__e, PrimNS2Value(symshen_4_5rule_6), V181)

		tmp389 := Call(__e, tmp372, tmp388)

		__e.TailApply(tmp358, tmp389)
		return

	}, 1)

	tmp390 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rules_6, tmp357)

	_ = tmp390

	tmp391 := MakeNative(func(__e *ControlFlow) {
		V189 := __e.Get(1)
		_ = V189
		tmp392 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp513 := Call(__e, PrimNS2Value(symfail))

			tmp514 := PrimEqual(YaccParse, tmp513)

			if True == tmp514 {
				tmp394 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp482 := Call(__e, PrimNS2Value(symfail))

					tmp483 := PrimEqual(YaccParse, tmp482)

					if True == tmp483 {
						tmp396 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp428 := Call(__e, PrimNS2Value(symfail))

							tmp429 := PrimEqual(YaccParse, tmp428)

							if True == tmp429 {
								tmp398 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5patterns_6 := __e.Get(1)
									_ = Parse__shen_4_5patterns_6
									tmp424 := Call(__e, PrimNS2Value(symfail))

									tmp425 := PrimEqual(tmp424, Parse__shen_4_5patterns_6)

									tmp426 := PrimNot(tmp425)

									if True == tmp426 {
										tmp422 := PrimHead(Parse__shen_4_5patterns_6)

										tmp423 := PrimIsPair(tmp422)

										var ifres418 Obj

										if True == tmp423 {
											tmp420 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5patterns_6)

											tmp421 := PrimEqual(sym_5_1, tmp420)

											var ifres419 Obj

											if True == tmp421 {
												ifres419 = True

											} else {
												ifres419 = False

											}

											ifres418 = ifres419

										} else {
											ifres418 = False

										}

										if True == ifres418 {
											tmp401 := MakeNative(func(__e *ControlFlow) {
												NewStream187 := __e.Get(1)
												_ = NewStream187
												tmp402 := MakeNative(func(__e *ControlFlow) {
													Parse__shen_4_5action_6 := __e.Get(1)
													_ = Parse__shen_4_5action_6
													tmp411 := Call(__e, PrimNS2Value(symfail))

													tmp412 := PrimEqual(tmp411, Parse__shen_4_5action_6)

													tmp413 := PrimNot(tmp412)

													if True == tmp413 {
														tmp404 := PrimHead(Parse__shen_4_5action_6)

														tmp405 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

														tmp406 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

														tmp407 := PrimCons(tmp406, Nil)

														tmp408 := PrimCons(symshen_4choicepoint_b, tmp407)

														tmp409 := PrimCons(tmp408, Nil)

														tmp410 := PrimCons(tmp405, tmp409)

														__e.TailApply(PrimNS2Value(symshen_4pair), tmp404, tmp410)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp414 := Call(__e, PrimNS2Value(symshen_4_5action_6), NewStream187)

												__e.TailApply(tmp402, tmp414)
												return

											}, 1)

											tmp415 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5patterns_6)

											tmp416 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

											tmp417 := Call(__e, PrimNS2Value(symshen_4pair), tmp415, tmp416)

											__e.TailApply(tmp401, tmp417)
											return

										} else {
											__e.TailApply(PrimNS2Value(symfail))
											return
										}

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp427 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V189)

								__e.TailApply(tmp398, tmp427)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp430 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5patterns_6 := __e.Get(1)
							_ = Parse__shen_4_5patterns_6
							tmp477 := Call(__e, PrimNS2Value(symfail))

							tmp478 := PrimEqual(tmp477, Parse__shen_4_5patterns_6)

							tmp479 := PrimNot(tmp478)

							if True == tmp479 {
								tmp475 := PrimHead(Parse__shen_4_5patterns_6)

								tmp476 := PrimIsPair(tmp475)

								var ifres471 Obj

								if True == tmp476 {
									tmp473 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5patterns_6)

									tmp474 := PrimEqual(sym_5_1, tmp473)

									var ifres472 Obj

									if True == tmp474 {
										ifres472 = True

									} else {
										ifres472 = False

									}

									ifres471 = ifres472

								} else {
									ifres471 = False

								}

								if True == ifres471 {
									tmp433 := MakeNative(func(__e *ControlFlow) {
										NewStream185 := __e.Get(1)
										_ = NewStream185
										tmp434 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5action_6 := __e.Get(1)
											_ = Parse__shen_4_5action_6
											tmp464 := Call(__e, PrimNS2Value(symfail))

											tmp465 := PrimEqual(tmp464, Parse__shen_4_5action_6)

											tmp466 := PrimNot(tmp465)

											if True == tmp466 {
												tmp462 := PrimHead(Parse__shen_4_5action_6)

												tmp463 := PrimIsPair(tmp462)

												var ifres458 Obj

												if True == tmp463 {
													tmp460 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5action_6)

													tmp461 := PrimEqual(symwhere, tmp460)

													var ifres459 Obj

													if True == tmp461 {
														ifres459 = True

													} else {
														ifres459 = False

													}

													ifres458 = ifres459

												} else {
													ifres458 = False

												}

												if True == ifres458 {
													tmp437 := MakeNative(func(__e *ControlFlow) {
														NewStream186 := __e.Get(1)
														_ = NewStream186
														tmp438 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5guard_6 := __e.Get(1)
															_ = Parse__shen_4_5guard_6
															tmp451 := Call(__e, PrimNS2Value(symfail))

															tmp452 := PrimEqual(tmp451, Parse__shen_4_5guard_6)

															tmp453 := PrimNot(tmp452)

															if True == tmp453 {
																tmp440 := PrimHead(Parse__shen_4_5guard_6)

																tmp441 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

																tmp442 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5guard_6)

																tmp443 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

																tmp444 := PrimCons(tmp443, Nil)

																tmp445 := PrimCons(symshen_4choicepoint_b, tmp444)

																tmp446 := PrimCons(tmp445, Nil)

																tmp447 := PrimCons(tmp442, tmp446)

																tmp448 := PrimCons(symwhere, tmp447)

																tmp449 := PrimCons(tmp448, Nil)

																tmp450 := PrimCons(tmp441, tmp449)

																__e.TailApply(PrimNS2Value(symshen_4pair), tmp440, tmp450)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp454 := Call(__e, PrimNS2Value(symshen_4_5guard_6), NewStream186)

														__e.TailApply(tmp438, tmp454)
														return

													}, 1)

													tmp455 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5action_6)

													tmp456 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

													tmp457 := Call(__e, PrimNS2Value(symshen_4pair), tmp455, tmp456)

													__e.TailApply(tmp437, tmp457)
													return

												} else {
													__e.TailApply(PrimNS2Value(symfail))
													return
												}

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp467 := Call(__e, PrimNS2Value(symshen_4_5action_6), NewStream185)

										__e.TailApply(tmp434, tmp467)
										return

									}, 1)

									tmp468 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5patterns_6)

									tmp469 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

									tmp470 := Call(__e, PrimNS2Value(symshen_4pair), tmp468, tmp469)

									__e.TailApply(tmp433, tmp470)
									return

								} else {
									__e.TailApply(PrimNS2Value(symfail))
									return
								}

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp480 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V189)

						tmp481 := Call(__e, tmp430, tmp480)

						__e.TailApply(tmp396, tmp481)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp484 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5patterns_6 := __e.Get(1)
					_ = Parse__shen_4_5patterns_6
					tmp508 := Call(__e, PrimNS2Value(symfail))

					tmp509 := PrimEqual(tmp508, Parse__shen_4_5patterns_6)

					tmp510 := PrimNot(tmp509)

					if True == tmp510 {
						tmp506 := PrimHead(Parse__shen_4_5patterns_6)

						tmp507 := PrimIsPair(tmp506)

						var ifres502 Obj

						if True == tmp507 {
							tmp504 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5patterns_6)

							tmp505 := PrimEqual(sym_1_6, tmp504)

							var ifres503 Obj

							if True == tmp505 {
								ifres503 = True

							} else {
								ifres503 = False

							}

							ifres502 = ifres503

						} else {
							ifres502 = False

						}

						if True == ifres502 {
							tmp487 := MakeNative(func(__e *ControlFlow) {
								NewStream184 := __e.Get(1)
								_ = NewStream184
								tmp488 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5action_6 := __e.Get(1)
									_ = Parse__shen_4_5action_6
									tmp495 := Call(__e, PrimNS2Value(symfail))

									tmp496 := PrimEqual(tmp495, Parse__shen_4_5action_6)

									tmp497 := PrimNot(tmp496)

									if True == tmp497 {
										tmp490 := PrimHead(Parse__shen_4_5action_6)

										tmp491 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

										tmp492 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

										tmp493 := PrimCons(tmp492, Nil)

										tmp494 := PrimCons(tmp491, tmp493)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp490, tmp494)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp498 := Call(__e, PrimNS2Value(symshen_4_5action_6), NewStream184)

								__e.TailApply(tmp488, tmp498)
								return

							}, 1)

							tmp499 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5patterns_6)

							tmp500 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

							tmp501 := Call(__e, PrimNS2Value(symshen_4pair), tmp499, tmp500)

							__e.TailApply(tmp487, tmp501)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp511 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V189)

				tmp512 := Call(__e, tmp484, tmp511)

				__e.TailApply(tmp394, tmp512)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp515 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5patterns_6 := __e.Get(1)
			_ = Parse__shen_4_5patterns_6
			tmp560 := Call(__e, PrimNS2Value(symfail))

			tmp561 := PrimEqual(tmp560, Parse__shen_4_5patterns_6)

			tmp562 := PrimNot(tmp561)

			if True == tmp562 {
				tmp558 := PrimHead(Parse__shen_4_5patterns_6)

				tmp559 := PrimIsPair(tmp558)

				var ifres554 Obj

				if True == tmp559 {
					tmp556 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5patterns_6)

					tmp557 := PrimEqual(sym_1_6, tmp556)

					var ifres555 Obj

					if True == tmp557 {
						ifres555 = True

					} else {
						ifres555 = False

					}

					ifres554 = ifres555

				} else {
					ifres554 = False

				}

				if True == ifres554 {
					tmp518 := MakeNative(func(__e *ControlFlow) {
						NewStream182 := __e.Get(1)
						_ = NewStream182
						tmp519 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5action_6 := __e.Get(1)
							_ = Parse__shen_4_5action_6
							tmp547 := Call(__e, PrimNS2Value(symfail))

							tmp548 := PrimEqual(tmp547, Parse__shen_4_5action_6)

							tmp549 := PrimNot(tmp548)

							if True == tmp549 {
								tmp545 := PrimHead(Parse__shen_4_5action_6)

								tmp546 := PrimIsPair(tmp545)

								var ifres541 Obj

								if True == tmp546 {
									tmp543 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5action_6)

									tmp544 := PrimEqual(symwhere, tmp543)

									var ifres542 Obj

									if True == tmp544 {
										ifres542 = True

									} else {
										ifres542 = False

									}

									ifres541 = ifres542

								} else {
									ifres541 = False

								}

								if True == ifres541 {
									tmp522 := MakeNative(func(__e *ControlFlow) {
										NewStream183 := __e.Get(1)
										_ = NewStream183
										tmp523 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5guard_6 := __e.Get(1)
											_ = Parse__shen_4_5guard_6
											tmp534 := Call(__e, PrimNS2Value(symfail))

											tmp535 := PrimEqual(tmp534, Parse__shen_4_5guard_6)

											tmp536 := PrimNot(tmp535)

											if True == tmp536 {
												tmp525 := PrimHead(Parse__shen_4_5guard_6)

												tmp526 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

												tmp527 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5guard_6)

												tmp528 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

												tmp529 := PrimCons(tmp528, Nil)

												tmp530 := PrimCons(tmp527, tmp529)

												tmp531 := PrimCons(symwhere, tmp530)

												tmp532 := PrimCons(tmp531, Nil)

												tmp533 := PrimCons(tmp526, tmp532)

												__e.TailApply(PrimNS2Value(symshen_4pair), tmp525, tmp533)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp537 := Call(__e, PrimNS2Value(symshen_4_5guard_6), NewStream183)

										__e.TailApply(tmp523, tmp537)
										return

									}, 1)

									tmp538 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5action_6)

									tmp539 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5action_6)

									tmp540 := Call(__e, PrimNS2Value(symshen_4pair), tmp538, tmp539)

									__e.TailApply(tmp522, tmp540)
									return

								} else {
									__e.TailApply(PrimNS2Value(symfail))
									return
								}

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp550 := Call(__e, PrimNS2Value(symshen_4_5action_6), NewStream182)

						__e.TailApply(tmp519, tmp550)
						return

					}, 1)

					tmp551 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5patterns_6)

					tmp552 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

					tmp553 := Call(__e, PrimNS2Value(symshen_4pair), tmp551, tmp552)

					__e.TailApply(tmp518, tmp553)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp563 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), V189)

		tmp564 := Call(__e, tmp515, tmp563)

		__e.TailApply(tmp392, tmp564)
		return

	}, 1)

	tmp565 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5rule_6, tmp391)

	_ = tmp565

	tmp566 := MakeNative(func(__e *ControlFlow) {
		V192 := __e.Get(1)
		_ = V192
		V193 := __e.Get(2)
		_ = V193
		tmp568 := Call(__e, V192, V193)

		if True == tmp568 {
			__e.TailApply(PrimNS2Value(symfail))
			return
		} else {
			__e.Return(V193)
			return
		}

	}, 2)

	tmp569 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fail__if, tmp566)

	_ = tmp569

	tmp570 := MakeNative(func(__e *ControlFlow) {
		V199 := __e.Get(1)
		_ = V199
		tmp572 := Call(__e, PrimNS2Value(symfail))

		tmp573 := PrimEqual(V199, tmp572)

		if True == tmp573 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp574 := Call(__e, PrimNS1Value(symns2_1set), symshen_4succeeds_2, tmp570)

	_ = tmp574

	tmp575 := MakeNative(func(__e *ControlFlow) {
		V202 := __e.Get(1)
		_ = V202
		V203 := __e.Get(2)
		_ = V203
		tmp576 := PrimNS3Value(symshen_4_dcustom_1pattern_1compiler_d)

		__e.TailApply(tmp576, V202, V203)
		return

	}, 2)

	tmp577 := Call(__e, PrimNS1Value(symns2_1set), symshen_4custom_1pattern_1compiler, tmp575)

	_ = tmp577

	tmp578 := MakeNative(func(__e *ControlFlow) {
		V205 := __e.Get(1)
		_ = V205
		tmp579 := PrimNS3Value(symshen_4_dcustom_1pattern_1reducer_d)

		__e.TailApply(tmp579, V205)
		return

	}, 1)

	tmp580 := Call(__e, PrimNS1Value(symns2_1set), symshen_4custom_1pattern_1reducer, tmp578)

	_ = tmp580

	tmp581 := MakeNative(func(__e *ControlFlow) {
		V207 := __e.Get(1)
		_ = V207
		tmp582 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp591 := Call(__e, PrimNS2Value(symfail))

			tmp592 := PrimEqual(YaccParse, tmp591)

			if True == tmp592 {
				tmp584 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp587 := Call(__e, PrimNS2Value(symfail))

					tmp588 := PrimEqual(tmp587, Parse___5e_6)

					tmp589 := PrimNot(tmp588)

					if True == tmp589 {
						tmp586 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp586, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp590 := Call(__e, PrimNS2Value(sym_5e_6), V207)

				__e.TailApply(tmp584, tmp590)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp593 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp605 := Call(__e, PrimNS2Value(symfail))

			tmp606 := PrimEqual(tmp605, Parse__shen_4_5pattern_6)

			tmp607 := PrimNot(tmp606)

			if True == tmp607 {
				tmp595 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5patterns_6 := __e.Get(1)
					_ = Parse__shen_4_5patterns_6
					tmp601 := Call(__e, PrimNS2Value(symfail))

					tmp602 := PrimEqual(tmp601, Parse__shen_4_5patterns_6)

					tmp603 := PrimNot(tmp602)

					if True == tmp603 {
						tmp597 := PrimHead(Parse__shen_4_5patterns_6)

						tmp598 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern_6)

						tmp599 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5patterns_6)

						tmp600 := PrimCons(tmp598, tmp599)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp597, tmp600)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp604 := Call(__e, PrimNS2Value(symshen_4_5patterns_6), Parse__shen_4_5pattern_6)

				__e.TailApply(tmp595, tmp604)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp608 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V207)

		tmp609 := Call(__e, tmp593, tmp608)

		__e.TailApply(tmp582, tmp609)
		return

	}, 1)

	tmp610 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5patterns_6, tmp581)

	_ = tmp610

	tmp611 := MakeNative(func(__e *ControlFlow) {
		V220 := __e.Get(1)
		_ = V220
		tmp612 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp869 := Call(__e, PrimNS2Value(symfail))

			tmp870 := PrimEqual(YaccParse, tmp869)

			if True == tmp870 {
				tmp614 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp813 := Call(__e, PrimNS2Value(symfail))

					tmp814 := PrimEqual(YaccParse, tmp813)

					if True == tmp814 {
						tmp616 := MakeNative(func(__e *ControlFlow) {
							YaccParse := __e.Get(1)
							_ = YaccParse
							tmp757 := Call(__e, PrimNS2Value(symfail))

							tmp758 := PrimEqual(YaccParse, tmp757)

							if True == tmp758 {
								tmp618 := MakeNative(func(__e *ControlFlow) {
									YaccParse := __e.Get(1)
									_ = YaccParse
									tmp701 := Call(__e, PrimNS2Value(symfail))

									tmp702 := PrimEqual(YaccParse, tmp701)

									if True == tmp702 {
										tmp620 := MakeNative(func(__e *ControlFlow) {
											YaccParse := __e.Get(1)
											_ = YaccParse
											tmp649 := Call(__e, PrimNS2Value(symfail))

											tmp650 := PrimEqual(YaccParse, tmp649)

											if True == tmp650 {
												tmp622 := MakeNative(func(__e *ControlFlow) {
													YaccParse := __e.Get(1)
													_ = YaccParse
													tmp632 := Call(__e, PrimNS2Value(symfail))

													tmp633 := PrimEqual(YaccParse, tmp632)

													if True == tmp633 {
														tmp624 := MakeNative(func(__e *ControlFlow) {
															Parse__shen_4_5simple__pattern_6 := __e.Get(1)
															_ = Parse__shen_4_5simple__pattern_6
															tmp628 := Call(__e, PrimNS2Value(symfail))

															tmp629 := PrimEqual(tmp628, Parse__shen_4_5simple__pattern_6)

															tmp630 := PrimNot(tmp629)

															if True == tmp630 {
																tmp626 := PrimHead(Parse__shen_4_5simple__pattern_6)

																tmp627 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5simple__pattern_6)

																__e.TailApply(PrimNS2Value(symshen_4pair), tmp626, tmp627)
																return

															} else {
																__e.TailApply(PrimNS2Value(symfail))
																return
															}

														}, 1)

														tmp631 := Call(__e, PrimNS2Value(symshen_4_5simple__pattern_6), V220)

														__e.TailApply(tmp624, tmp631)
														return

													} else {
														__e.Return(YaccParse)
														return
													}

												}, 1)

												tmp647 := PrimHead(V220)

												tmp648 := PrimIsPair(tmp647)

												var ifres634 Obj

												if True == tmp648 {
													tmp636 := MakeNative(func(__e *ControlFlow) {
														Parse__X := __e.Get(1)
														_ = Parse__X
														tmp644 := PrimIsPair(Parse__X)

														if True == tmp644 {
															tmp638 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

															tmp639 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

															tmp640 := Call(__e, PrimNS2Value(symshen_4pair), tmp638, tmp639)

															tmp641 := PrimHead(tmp640)

															tmp642 := MakeNative(func(__e *ControlFlow) {
																__e.TailApply(PrimNS2Value(symshen_4constructor_1error), Parse__X)
																return
															}, 0)

															tmp643 := Call(__e, PrimNS2Value(symshen_4custom_1pattern_1compiler), Parse__X, tmp642)

															__e.TailApply(PrimNS2Value(symshen_4pair), tmp641, tmp643)
															return

														} else {
															__e.TailApply(PrimNS2Value(symfail))
															return
														}

													}, 1)

													tmp645 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

													tmp646 := Call(__e, tmp636, tmp645)

													ifres634 = tmp646

												} else {
													tmp635 := Call(__e, PrimNS2Value(symfail))

													ifres634 = tmp635

												}

												__e.TailApply(tmp622, ifres634)
												return

											} else {
												__e.Return(YaccParse)
												return
											}

										}, 1)

										tmp699 := PrimHead(V220)

										tmp700 := PrimIsPair(tmp699)

										var ifres695 Obj

										if True == tmp700 {
											tmp697 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

											tmp698 := PrimIsPair(tmp697)

											var ifres696 Obj

											if True == tmp698 {
												ifres696 = True

											} else {
												ifres696 = False

											}

											ifres695 = ifres696

										} else {
											ifres695 = False

										}

										var ifres651 Obj

										if True == ifres695 {
											tmp690 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

											tmp691 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

											tmp692 := Call(__e, PrimNS2Value(symshen_4pair), tmp690, tmp691)

											tmp693 := PrimHead(tmp692)

											tmp694 := PrimIsPair(tmp693)

											var ifres683 Obj

											if True == tmp694 {
												tmp685 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

												tmp686 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

												tmp687 := Call(__e, PrimNS2Value(symshen_4pair), tmp685, tmp686)

												tmp688 := Call(__e, PrimNS2Value(symshen_4hdhd), tmp687)

												tmp689 := PrimEqual(symvector, tmp688)

												var ifres684 Obj

												if True == tmp689 {
													ifres684 = True

												} else {
													ifres684 = False

												}

												ifres683 = ifres684

											} else {
												ifres683 = False

											}

											var ifres653 Obj

											if True == ifres683 {
												tmp655 := MakeNative(func(__e *ControlFlow) {
													NewStream217 := __e.Get(1)
													_ = NewStream217
													tmp671 := PrimHead(NewStream217)

													tmp672 := PrimIsPair(tmp671)

													var ifres667 Obj

													if True == tmp672 {
														tmp669 := Call(__e, PrimNS2Value(symshen_4hdhd), NewStream217)

														tmp670 := PrimEqual(MakeNumber(0), tmp669)

														var ifres668 Obj

														if True == tmp670 {
															ifres668 = True

														} else {
															ifres668 = False

														}

														ifres667 = ifres668

													} else {
														ifres667 = False

													}

													if True == ifres667 {
														tmp657 := MakeNative(func(__e *ControlFlow) {
															NewStream218 := __e.Get(1)
															_ = NewStream218
															tmp658 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

															tmp659 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

															tmp660 := Call(__e, PrimNS2Value(symshen_4pair), tmp658, tmp659)

															tmp661 := PrimHead(tmp660)

															tmp662 := PrimCons(MakeNumber(0), Nil)

															tmp663 := PrimCons(symvector, tmp662)

															__e.TailApply(PrimNS2Value(symshen_4pair), tmp661, tmp663)
															return

														}, 1)

														tmp664 := Call(__e, PrimNS2Value(symshen_4tlhd), NewStream217)

														tmp665 := Call(__e, PrimNS2Value(symshen_4hdtl), NewStream217)

														tmp666 := Call(__e, PrimNS2Value(symshen_4pair), tmp664, tmp665)

														__e.TailApply(tmp657, tmp666)
														return

													} else {
														__e.TailApply(PrimNS2Value(symfail))
														return
													}

												}, 1)

												tmp673 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

												tmp674 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

												tmp675 := Call(__e, PrimNS2Value(symshen_4pair), tmp673, tmp674)

												tmp676 := Call(__e, PrimNS2Value(symshen_4tlhd), tmp675)

												tmp677 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

												tmp678 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

												tmp679 := Call(__e, PrimNS2Value(symshen_4pair), tmp677, tmp678)

												tmp680 := Call(__e, PrimNS2Value(symshen_4hdtl), tmp679)

												tmp681 := Call(__e, PrimNS2Value(symshen_4pair), tmp676, tmp680)

												tmp682 := Call(__e, tmp655, tmp681)

												ifres653 = tmp682

											} else {
												tmp654 := Call(__e, PrimNS2Value(symfail))

												ifres653 = tmp654

											}

											ifres651 = ifres653

										} else {
											tmp652 := Call(__e, PrimNS2Value(symfail))

											ifres651 = tmp652

										}

										__e.TailApply(tmp620, ifres651)
										return

									} else {
										__e.Return(YaccParse)
										return
									}

								}, 1)

								tmp755 := PrimHead(V220)

								tmp756 := PrimIsPair(tmp755)

								var ifres751 Obj

								if True == tmp756 {
									tmp753 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

									tmp754 := PrimIsPair(tmp753)

									var ifres752 Obj

									if True == tmp754 {
										ifres752 = True

									} else {
										ifres752 = False

									}

									ifres751 = ifres752

								} else {
									ifres751 = False

								}

								var ifres703 Obj

								if True == ifres751 {
									tmp746 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

									tmp747 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

									tmp748 := Call(__e, PrimNS2Value(symshen_4pair), tmp746, tmp747)

									tmp749 := PrimHead(tmp748)

									tmp750 := PrimIsPair(tmp749)

									var ifres739 Obj

									if True == tmp750 {
										tmp741 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

										tmp742 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

										tmp743 := Call(__e, PrimNS2Value(symshen_4pair), tmp741, tmp742)

										tmp744 := Call(__e, PrimNS2Value(symshen_4hdhd), tmp743)

										tmp745 := PrimEqual(sym_8s, tmp744)

										var ifres740 Obj

										if True == tmp745 {
											ifres740 = True

										} else {
											ifres740 = False

										}

										ifres739 = ifres740

									} else {
										ifres739 = False

									}

									var ifres705 Obj

									if True == ifres739 {
										tmp707 := MakeNative(func(__e *ControlFlow) {
											NewStream215 := __e.Get(1)
											_ = NewStream215
											tmp708 := MakeNative(func(__e *ControlFlow) {
												Parse__shen_4_5pattern1_6 := __e.Get(1)
												_ = Parse__shen_4_5pattern1_6
												tmp725 := Call(__e, PrimNS2Value(symfail))

												tmp726 := PrimEqual(tmp725, Parse__shen_4_5pattern1_6)

												tmp727 := PrimNot(tmp726)

												if True == tmp727 {
													tmp710 := MakeNative(func(__e *ControlFlow) {
														Parse__shen_4_5pattern2_6 := __e.Get(1)
														_ = Parse__shen_4_5pattern2_6
														tmp721 := Call(__e, PrimNS2Value(symfail))

														tmp722 := PrimEqual(tmp721, Parse__shen_4_5pattern2_6)

														tmp723 := PrimNot(tmp722)

														if True == tmp723 {
															tmp712 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

															tmp713 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

															tmp714 := Call(__e, PrimNS2Value(symshen_4pair), tmp712, tmp713)

															tmp715 := PrimHead(tmp714)

															tmp716 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern1_6)

															tmp717 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern2_6)

															tmp718 := PrimCons(tmp717, Nil)

															tmp719 := PrimCons(tmp716, tmp718)

															tmp720 := PrimCons(sym_8s, tmp719)

															__e.TailApply(PrimNS2Value(symshen_4pair), tmp715, tmp720)
															return

														} else {
															__e.TailApply(PrimNS2Value(symfail))
															return
														}

													}, 1)

													tmp724 := Call(__e, PrimNS2Value(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

													__e.TailApply(tmp710, tmp724)
													return

												} else {
													__e.TailApply(PrimNS2Value(symfail))
													return
												}

											}, 1)

											tmp728 := Call(__e, PrimNS2Value(symshen_4_5pattern1_6), NewStream215)

											__e.TailApply(tmp708, tmp728)
											return

										}, 1)

										tmp729 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

										tmp730 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

										tmp731 := Call(__e, PrimNS2Value(symshen_4pair), tmp729, tmp730)

										tmp732 := Call(__e, PrimNS2Value(symshen_4tlhd), tmp731)

										tmp733 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

										tmp734 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

										tmp735 := Call(__e, PrimNS2Value(symshen_4pair), tmp733, tmp734)

										tmp736 := Call(__e, PrimNS2Value(symshen_4hdtl), tmp735)

										tmp737 := Call(__e, PrimNS2Value(symshen_4pair), tmp732, tmp736)

										tmp738 := Call(__e, tmp707, tmp737)

										ifres705 = tmp738

									} else {
										tmp706 := Call(__e, PrimNS2Value(symfail))

										ifres705 = tmp706

									}

									ifres703 = ifres705

								} else {
									tmp704 := Call(__e, PrimNS2Value(symfail))

									ifres703 = tmp704

								}

								__e.TailApply(tmp618, ifres703)
								return

							} else {
								__e.Return(YaccParse)
								return
							}

						}, 1)

						tmp811 := PrimHead(V220)

						tmp812 := PrimIsPair(tmp811)

						var ifres807 Obj

						if True == tmp812 {
							tmp809 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

							tmp810 := PrimIsPair(tmp809)

							var ifres808 Obj

							if True == tmp810 {
								ifres808 = True

							} else {
								ifres808 = False

							}

							ifres807 = ifres808

						} else {
							ifres807 = False

						}

						var ifres759 Obj

						if True == ifres807 {
							tmp802 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

							tmp803 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

							tmp804 := Call(__e, PrimNS2Value(symshen_4pair), tmp802, tmp803)

							tmp805 := PrimHead(tmp804)

							tmp806 := PrimIsPair(tmp805)

							var ifres795 Obj

							if True == tmp806 {
								tmp797 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

								tmp798 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

								tmp799 := Call(__e, PrimNS2Value(symshen_4pair), tmp797, tmp798)

								tmp800 := Call(__e, PrimNS2Value(symshen_4hdhd), tmp799)

								tmp801 := PrimEqual(sym_8v, tmp800)

								var ifres796 Obj

								if True == tmp801 {
									ifres796 = True

								} else {
									ifres796 = False

								}

								ifres795 = ifres796

							} else {
								ifres795 = False

							}

							var ifres761 Obj

							if True == ifres795 {
								tmp763 := MakeNative(func(__e *ControlFlow) {
									NewStream213 := __e.Get(1)
									_ = NewStream213
									tmp764 := MakeNative(func(__e *ControlFlow) {
										Parse__shen_4_5pattern1_6 := __e.Get(1)
										_ = Parse__shen_4_5pattern1_6
										tmp781 := Call(__e, PrimNS2Value(symfail))

										tmp782 := PrimEqual(tmp781, Parse__shen_4_5pattern1_6)

										tmp783 := PrimNot(tmp782)

										if True == tmp783 {
											tmp766 := MakeNative(func(__e *ControlFlow) {
												Parse__shen_4_5pattern2_6 := __e.Get(1)
												_ = Parse__shen_4_5pattern2_6
												tmp777 := Call(__e, PrimNS2Value(symfail))

												tmp778 := PrimEqual(tmp777, Parse__shen_4_5pattern2_6)

												tmp779 := PrimNot(tmp778)

												if True == tmp779 {
													tmp768 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

													tmp769 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

													tmp770 := Call(__e, PrimNS2Value(symshen_4pair), tmp768, tmp769)

													tmp771 := PrimHead(tmp770)

													tmp772 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern1_6)

													tmp773 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern2_6)

													tmp774 := PrimCons(tmp773, Nil)

													tmp775 := PrimCons(tmp772, tmp774)

													tmp776 := PrimCons(sym_8v, tmp775)

													__e.TailApply(PrimNS2Value(symshen_4pair), tmp771, tmp776)
													return

												} else {
													__e.TailApply(PrimNS2Value(symfail))
													return
												}

											}, 1)

											tmp780 := Call(__e, PrimNS2Value(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

											__e.TailApply(tmp766, tmp780)
											return

										} else {
											__e.TailApply(PrimNS2Value(symfail))
											return
										}

									}, 1)

									tmp784 := Call(__e, PrimNS2Value(symshen_4_5pattern1_6), NewStream213)

									__e.TailApply(tmp764, tmp784)
									return

								}, 1)

								tmp785 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

								tmp786 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

								tmp787 := Call(__e, PrimNS2Value(symshen_4pair), tmp785, tmp786)

								tmp788 := Call(__e, PrimNS2Value(symshen_4tlhd), tmp787)

								tmp789 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

								tmp790 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

								tmp791 := Call(__e, PrimNS2Value(symshen_4pair), tmp789, tmp790)

								tmp792 := Call(__e, PrimNS2Value(symshen_4hdtl), tmp791)

								tmp793 := Call(__e, PrimNS2Value(symshen_4pair), tmp788, tmp792)

								tmp794 := Call(__e, tmp763, tmp793)

								ifres761 = tmp794

							} else {
								tmp762 := Call(__e, PrimNS2Value(symfail))

								ifres761 = tmp762

							}

							ifres759 = ifres761

						} else {
							tmp760 := Call(__e, PrimNS2Value(symfail))

							ifres759 = tmp760

						}

						__e.TailApply(tmp616, ifres759)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp867 := PrimHead(V220)

				tmp868 := PrimIsPair(tmp867)

				var ifres863 Obj

				if True == tmp868 {
					tmp865 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

					tmp866 := PrimIsPair(tmp865)

					var ifres864 Obj

					if True == tmp866 {
						ifres864 = True

					} else {
						ifres864 = False

					}

					ifres863 = ifres864

				} else {
					ifres863 = False

				}

				var ifres815 Obj

				if True == ifres863 {
					tmp858 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

					tmp859 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

					tmp860 := Call(__e, PrimNS2Value(symshen_4pair), tmp858, tmp859)

					tmp861 := PrimHead(tmp860)

					tmp862 := PrimIsPair(tmp861)

					var ifres851 Obj

					if True == tmp862 {
						tmp853 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

						tmp854 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

						tmp855 := Call(__e, PrimNS2Value(symshen_4pair), tmp853, tmp854)

						tmp856 := Call(__e, PrimNS2Value(symshen_4hdhd), tmp855)

						tmp857 := PrimEqual(symcons, tmp856)

						var ifres852 Obj

						if True == tmp857 {
							ifres852 = True

						} else {
							ifres852 = False

						}

						ifres851 = ifres852

					} else {
						ifres851 = False

					}

					var ifres817 Obj

					if True == ifres851 {
						tmp819 := MakeNative(func(__e *ControlFlow) {
							NewStream211 := __e.Get(1)
							_ = NewStream211
							tmp820 := MakeNative(func(__e *ControlFlow) {
								Parse__shen_4_5pattern1_6 := __e.Get(1)
								_ = Parse__shen_4_5pattern1_6
								tmp837 := Call(__e, PrimNS2Value(symfail))

								tmp838 := PrimEqual(tmp837, Parse__shen_4_5pattern1_6)

								tmp839 := PrimNot(tmp838)

								if True == tmp839 {
									tmp822 := MakeNative(func(__e *ControlFlow) {
										Parse__shen_4_5pattern2_6 := __e.Get(1)
										_ = Parse__shen_4_5pattern2_6
										tmp833 := Call(__e, PrimNS2Value(symfail))

										tmp834 := PrimEqual(tmp833, Parse__shen_4_5pattern2_6)

										tmp835 := PrimNot(tmp834)

										if True == tmp835 {
											tmp824 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

											tmp825 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

											tmp826 := Call(__e, PrimNS2Value(symshen_4pair), tmp824, tmp825)

											tmp827 := PrimHead(tmp826)

											tmp828 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern1_6)

											tmp829 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern2_6)

											tmp830 := PrimCons(tmp829, Nil)

											tmp831 := PrimCons(tmp828, tmp830)

											tmp832 := PrimCons(symcons, tmp831)

											__e.TailApply(PrimNS2Value(symshen_4pair), tmp827, tmp832)
											return

										} else {
											__e.TailApply(PrimNS2Value(symfail))
											return
										}

									}, 1)

									tmp836 := Call(__e, PrimNS2Value(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

									__e.TailApply(tmp822, tmp836)
									return

								} else {
									__e.TailApply(PrimNS2Value(symfail))
									return
								}

							}, 1)

							tmp840 := Call(__e, PrimNS2Value(symshen_4_5pattern1_6), NewStream211)

							__e.TailApply(tmp820, tmp840)
							return

						}, 1)

						tmp841 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

						tmp842 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

						tmp843 := Call(__e, PrimNS2Value(symshen_4pair), tmp841, tmp842)

						tmp844 := Call(__e, PrimNS2Value(symshen_4tlhd), tmp843)

						tmp845 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

						tmp846 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

						tmp847 := Call(__e, PrimNS2Value(symshen_4pair), tmp845, tmp846)

						tmp848 := Call(__e, PrimNS2Value(symshen_4hdtl), tmp847)

						tmp849 := Call(__e, PrimNS2Value(symshen_4pair), tmp844, tmp848)

						tmp850 := Call(__e, tmp819, tmp849)

						ifres817 = tmp850

					} else {
						tmp818 := Call(__e, PrimNS2Value(symfail))

						ifres817 = tmp818

					}

					ifres815 = ifres817

				} else {
					tmp816 := Call(__e, PrimNS2Value(symfail))

					ifres815 = tmp816

				}

				__e.TailApply(tmp614, ifres815)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp923 := PrimHead(V220)

		tmp924 := PrimIsPair(tmp923)

		var ifres919 Obj

		if True == tmp924 {
			tmp921 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

			tmp922 := PrimIsPair(tmp921)

			var ifres920 Obj

			if True == tmp922 {
				ifres920 = True

			} else {
				ifres920 = False

			}

			ifres919 = ifres920

		} else {
			ifres919 = False

		}

		var ifres871 Obj

		if True == ifres919 {
			tmp914 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

			tmp915 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

			tmp916 := Call(__e, PrimNS2Value(symshen_4pair), tmp914, tmp915)

			tmp917 := PrimHead(tmp916)

			tmp918 := PrimIsPair(tmp917)

			var ifres907 Obj

			if True == tmp918 {
				tmp909 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

				tmp910 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

				tmp911 := Call(__e, PrimNS2Value(symshen_4pair), tmp909, tmp910)

				tmp912 := Call(__e, PrimNS2Value(symshen_4hdhd), tmp911)

				tmp913 := PrimEqual(sym_8p, tmp912)

				var ifres908 Obj

				if True == tmp913 {
					ifres908 = True

				} else {
					ifres908 = False

				}

				ifres907 = ifres908

			} else {
				ifres907 = False

			}

			var ifres873 Obj

			if True == ifres907 {
				tmp875 := MakeNative(func(__e *ControlFlow) {
					NewStream209 := __e.Get(1)
					_ = NewStream209
					tmp876 := MakeNative(func(__e *ControlFlow) {
						Parse__shen_4_5pattern1_6 := __e.Get(1)
						_ = Parse__shen_4_5pattern1_6
						tmp893 := Call(__e, PrimNS2Value(symfail))

						tmp894 := PrimEqual(tmp893, Parse__shen_4_5pattern1_6)

						tmp895 := PrimNot(tmp894)

						if True == tmp895 {
							tmp878 := MakeNative(func(__e *ControlFlow) {
								Parse__shen_4_5pattern2_6 := __e.Get(1)
								_ = Parse__shen_4_5pattern2_6
								tmp889 := Call(__e, PrimNS2Value(symfail))

								tmp890 := PrimEqual(tmp889, Parse__shen_4_5pattern2_6)

								tmp891 := PrimNot(tmp890)

								if True == tmp891 {
									tmp880 := Call(__e, PrimNS2Value(symshen_4tlhd), V220)

									tmp881 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

									tmp882 := Call(__e, PrimNS2Value(symshen_4pair), tmp880, tmp881)

									tmp883 := PrimHead(tmp882)

									tmp884 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern1_6)

									tmp885 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern2_6)

									tmp886 := PrimCons(tmp885, Nil)

									tmp887 := PrimCons(tmp884, tmp886)

									tmp888 := PrimCons(sym_8p, tmp887)

									__e.TailApply(PrimNS2Value(symshen_4pair), tmp883, tmp888)
									return

								} else {
									__e.TailApply(PrimNS2Value(symfail))
									return
								}

							}, 1)

							tmp892 := Call(__e, PrimNS2Value(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

							__e.TailApply(tmp878, tmp892)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					}, 1)

					tmp896 := Call(__e, PrimNS2Value(symshen_4_5pattern1_6), NewStream209)

					__e.TailApply(tmp876, tmp896)
					return

				}, 1)

				tmp897 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

				tmp898 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

				tmp899 := Call(__e, PrimNS2Value(symshen_4pair), tmp897, tmp898)

				tmp900 := Call(__e, PrimNS2Value(symshen_4tlhd), tmp899)

				tmp901 := Call(__e, PrimNS2Value(symshen_4hdhd), V220)

				tmp902 := Call(__e, PrimNS2Value(symshen_4hdtl), V220)

				tmp903 := Call(__e, PrimNS2Value(symshen_4pair), tmp901, tmp902)

				tmp904 := Call(__e, PrimNS2Value(symshen_4hdtl), tmp903)

				tmp905 := Call(__e, PrimNS2Value(symshen_4pair), tmp900, tmp904)

				tmp906 := Call(__e, tmp875, tmp905)

				ifres873 = tmp906

			} else {
				tmp874 := Call(__e, PrimNS2Value(symfail))

				ifres873 = tmp874

			}

			ifres871 = ifres873

		} else {
			tmp872 := Call(__e, PrimNS2Value(symfail))

			ifres871 = tmp872

		}

		__e.TailApply(tmp612, ifres871)
		return

	}, 1)

	tmp925 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern_6, tmp611)

	_ = tmp925

	tmp926 := MakeNative(func(__e *ControlFlow) {
		V222 := __e.Get(1)
		_ = V222
		tmp927 := Call(__e, PrimNS2Value(symshen_4app), V222, MakeString(" is not a legitimate constructor\n"), symshen_4a)

		__e.Return(PrimSimpleError(tmp927))
		return

	}, 1)

	tmp928 := Call(__e, PrimNS1Value(symns2_1set), symshen_4constructor_1error, tmp926)

	_ = tmp928

	tmp929 := MakeNative(func(__e *ControlFlow) {
		V224 := __e.Get(1)
		_ = V224
		tmp930 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp946 := Call(__e, PrimNS2Value(symfail))

			tmp947 := PrimEqual(YaccParse, tmp946)

			if True == tmp947 {
				tmp944 := PrimHead(V224)

				tmp945 := PrimIsPair(tmp944)

				if True == tmp945 {
					tmp933 := MakeNative(func(__e *ControlFlow) {
						Parse__X := __e.Get(1)
						_ = Parse__X
						tmp939 := PrimCons(sym_5_1, Nil)

						tmp940 := PrimCons(sym_1_6, tmp939)

						tmp941 := Call(__e, PrimNS2Value(symelement_2), Parse__X, tmp940)

						tmp942 := PrimNot(tmp941)

						if True == tmp942 {
							tmp935 := Call(__e, PrimNS2Value(symshen_4tlhd), V224)

							tmp936 := Call(__e, PrimNS2Value(symshen_4hdtl), V224)

							tmp937 := Call(__e, PrimNS2Value(symshen_4pair), tmp935, tmp936)

							tmp938 := PrimHead(tmp937)

							__e.TailApply(PrimNS2Value(symshen_4pair), tmp938, Parse__X)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					}, 1)

					tmp943 := Call(__e, PrimNS2Value(symshen_4hdhd), V224)

					__e.TailApply(tmp933, tmp943)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp960 := PrimHead(V224)

		tmp961 := PrimIsPair(tmp960)

		var ifres948 Obj

		if True == tmp961 {
			tmp950 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp957 := PrimEqual(Parse__X, sym__)

				if True == tmp957 {
					tmp952 := Call(__e, PrimNS2Value(symshen_4tlhd), V224)

					tmp953 := Call(__e, PrimNS2Value(symshen_4hdtl), V224)

					tmp954 := Call(__e, PrimNS2Value(symshen_4pair), tmp952, tmp953)

					tmp955 := PrimHead(tmp954)

					tmp956 := Call(__e, PrimNS2Value(symgensym), symParse__Y)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp955, tmp956)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp958 := Call(__e, PrimNS2Value(symshen_4hdhd), V224)

			tmp959 := Call(__e, tmp950, tmp958)

			ifres948 = tmp959

		} else {
			tmp949 := Call(__e, PrimNS2Value(symfail))

			ifres948 = tmp949

		}

		__e.TailApply(tmp930, ifres948)
		return

	}, 1)

	tmp962 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5simple__pattern_6, tmp929)

	_ = tmp962

	tmp963 := MakeNative(func(__e *ControlFlow) {
		V226 := __e.Get(1)
		_ = V226
		tmp964 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp968 := Call(__e, PrimNS2Value(symfail))

			tmp969 := PrimEqual(tmp968, Parse__shen_4_5pattern_6)

			tmp970 := PrimNot(tmp969)

			if True == tmp970 {
				tmp966 := PrimHead(Parse__shen_4_5pattern_6)

				tmp967 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp966, tmp967)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp971 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V226)

		__e.TailApply(tmp964, tmp971)
		return

	}, 1)

	tmp972 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern1_6, tmp963)

	_ = tmp972

	tmp973 := MakeNative(func(__e *ControlFlow) {
		V228 := __e.Get(1)
		_ = V228
		tmp974 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5pattern_6 := __e.Get(1)
			_ = Parse__shen_4_5pattern_6
			tmp978 := Call(__e, PrimNS2Value(symfail))

			tmp979 := PrimEqual(tmp978, Parse__shen_4_5pattern_6)

			tmp980 := PrimNot(tmp979)

			if True == tmp980 {
				tmp976 := PrimHead(Parse__shen_4_5pattern_6)

				tmp977 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp976, tmp977)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp981 := Call(__e, PrimNS2Value(symshen_4_5pattern_6), V228)

		__e.TailApply(tmp974, tmp981)
		return

	}, 1)

	tmp982 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5pattern2_6, tmp973)

	_ = tmp982

	tmp983 := MakeNative(func(__e *ControlFlow) {
		V230 := __e.Get(1)
		_ = V230
		tmp991 := PrimHead(V230)

		tmp992 := PrimIsPair(tmp991)

		if True == tmp992 {
			tmp985 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp986 := Call(__e, PrimNS2Value(symshen_4tlhd), V230)

				tmp987 := Call(__e, PrimNS2Value(symshen_4hdtl), V230)

				tmp988 := Call(__e, PrimNS2Value(symshen_4pair), tmp986, tmp987)

				tmp989 := PrimHead(tmp988)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp989, Parse__X)
				return

			}, 1)

			tmp990 := Call(__e, PrimNS2Value(symshen_4hdhd), V230)

			__e.TailApply(tmp985, tmp990)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp993 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5action_6, tmp983)

	_ = tmp993

	tmp994 := MakeNative(func(__e *ControlFlow) {
		V232 := __e.Get(1)
		_ = V232
		tmp1002 := PrimHead(V232)

		tmp1003 := PrimIsPair(tmp1002)

		if True == tmp1003 {
			tmp996 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp997 := Call(__e, PrimNS2Value(symshen_4tlhd), V232)

				tmp998 := Call(__e, PrimNS2Value(symshen_4hdtl), V232)

				tmp999 := Call(__e, PrimNS2Value(symshen_4pair), tmp997, tmp998)

				tmp1000 := PrimHead(tmp999)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp1000, Parse__X)
				return

			}, 1)

			tmp1001 := Call(__e, PrimNS2Value(symshen_4hdhd), V232)

			__e.TailApply(tmp996, tmp1001)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp1004 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5guard_6, tmp994)

	_ = tmp1004

	tmp1005 := MakeNative(func(__e *ControlFlow) {
		V235 := __e.Get(1)
		_ = V235
		V236 := __e.Get(2)
		_ = V236
		tmp1006 := MakeNative(func(__e *ControlFlow) {
			Lambda_7 := __e.Get(1)
			_ = Lambda_7
			tmp1007 := MakeNative(func(__e *ControlFlow) {
				KL := __e.Get(1)
				_ = KL
				tmp1008 := MakeNative(func(__e *ControlFlow) {
					Record := __e.Get(1)
					_ = Record
					__e.Return(KL)
					return
				}, 1)

				tmp1009 := Call(__e, PrimNS2Value(symshen_4record_1source), V235, KL)

				__e.TailApply(tmp1008, tmp1009)
				return

			}, 1)

			tmp1010 := Call(__e, PrimNS2Value(symshen_4compile__to__kl), V235, Lambda_7)

			__e.TailApply(tmp1007, tmp1010)
			return

		}, 1)

		tmp1011 := Call(__e, PrimNS2Value(symshen_4compile__to__lambda_7), V235, V236)

		__e.TailApply(tmp1006, tmp1011)
		return

	}, 2)

	tmp1012 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__machine__code, tmp1005)

	_ = tmp1012

	tmp1013 := MakeNative(func(__e *ControlFlow) {
		V241 := __e.Get(1)
		_ = V241
		V242 := __e.Get(2)
		_ = V242
		tmp1016 := PrimNS3Value(symshen_4_dinstalling_1kl_d)

		if True == tmp1016 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp1015 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), V241, symshen_4source, V242, tmp1015)
			return

		}

	}, 2)

	tmp1017 := Call(__e, PrimNS1Value(symns2_1set), symshen_4record_1source, tmp1013)

	_ = tmp1017

	tmp1018 := MakeNative(func(__e *ControlFlow) {
		V245 := __e.Get(1)
		_ = V245
		V246 := __e.Get(2)
		_ = V246
		tmp1019 := MakeNative(func(__e *ControlFlow) {
			Arity := __e.Get(1)
			_ = Arity
			tmp1020 := MakeNative(func(__e *ControlFlow) {
				UpDateSymbolTable := __e.Get(1)
				_ = UpDateSymbolTable
				tmp1021 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					tmp1022 := MakeNative(func(__e *ControlFlow) {
						Variables := __e.Get(1)
						_ = Variables
						tmp1023 := MakeNative(func(__e *ControlFlow) {
							Strip := __e.Get(1)
							_ = Strip
							tmp1024 := MakeNative(func(__e *ControlFlow) {
								Abstractions := __e.Get(1)
								_ = Abstractions
								tmp1025 := MakeNative(func(__e *ControlFlow) {
									Applications := __e.Get(1)
									_ = Applications
									tmp1026 := PrimCons(Applications, Nil)

									__e.Return(PrimCons(Variables, tmp1026))
									return

								}, 1)

								tmp1027 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									__e.TailApply(PrimNS2Value(symshen_4application__build), Variables, X)
									return
								}, 1)

								tmp1028 := Call(__e, PrimNS2Value(symmap), tmp1027, Abstractions)

								__e.TailApply(tmp1025, tmp1028)
								return

							}, 1)

							tmp1029 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								__e.TailApply(PrimNS2Value(symshen_4abstract__rule), X)
								return
							}, 1)

							tmp1030 := Call(__e, PrimNS2Value(symmap), tmp1029, Strip)

							__e.TailApply(tmp1024, tmp1030)
							return

						}, 1)

						tmp1031 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(PrimNS2Value(symshen_4strip_1protect), X)
							return
						}, 1)

						tmp1032 := Call(__e, PrimNS2Value(symmap), tmp1031, V246)

						__e.TailApply(tmp1023, tmp1032)
						return

					}, 1)

					tmp1033 := Call(__e, PrimNS2Value(symshen_4parameters), Arity)

					__e.TailApply(tmp1022, tmp1033)
					return

				}, 1)

				tmp1034 := MakeNative(func(__e *ControlFlow) {
					Rule := __e.Get(1)
					_ = Rule
					__e.TailApply(PrimNS2Value(symshen_4free__variable__check), V245, Rule)
					return
				}, 1)

				tmp1035 := Call(__e, PrimNS2Value(symshen_4for_1each), tmp1034, V246)

				__e.TailApply(tmp1021, tmp1035)
				return

			}, 1)

			tmp1036 := Call(__e, PrimNS2Value(symshen_4update_1symbol_1table), V245, Arity)

			__e.TailApply(tmp1020, tmp1036)
			return

		}, 1)

		tmp1037 := Call(__e, PrimNS2Value(symshen_4aritycheck), V245, V246)

		__e.TailApply(tmp1019, tmp1037)
		return

	}, 2)

	tmp1038 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__lambda_7, tmp1018)

	_ = tmp1038

	tmp1039 := MakeNative(func(__e *ControlFlow) {
		V249 := __e.Get(1)
		_ = V249
		V250 := __e.Get(2)
		_ = V250
		tmp1044 := PrimEqual(MakeNumber(0), V250)

		if True == tmp1044 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp1041 := Call(__e, PrimNS2Value(symshen_4lambda_1form), V249, V250)

			tmp1042 := Call(__e, PrimNS2Value(symeval_1kl), tmp1041)

			tmp1043 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), V249, symshen_4lambda_1form, tmp1042, tmp1043)
			return

		}

	}, 2)

	tmp1045 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update_1symbol_1table, tmp1039)

	_ = tmp1045

	tmp1046 := MakeNative(func(__e *ControlFlow) {
		V253 := __e.Get(1)
		_ = V253
		V254 := __e.Get(2)
		_ = V254
		tmp1064 := PrimIsPair(V254)

		var ifres1055 Obj

		if True == tmp1064 {
			tmp1062 := PrimTail(V254)

			tmp1063 := PrimIsPair(tmp1062)

			var ifres1057 Obj

			if True == tmp1063 {
				tmp1059 := PrimTail(V254)

				tmp1060 := PrimTail(tmp1059)

				tmp1061 := PrimEqual(Nil, tmp1060)

				var ifres1058 Obj

				if True == tmp1061 {
					ifres1058 = True

				} else {
					ifres1058 = False

				}

				ifres1057 = ifres1058

			} else {
				ifres1057 = False

			}

			var ifres1056 Obj

			if True == ifres1057 {
				ifres1056 = True

			} else {
				ifres1056 = False

			}

			ifres1055 = ifres1056

		} else {
			ifres1055 = False

		}

		if True == ifres1055 {
			tmp1048 := MakeNative(func(__e *ControlFlow) {
				Bound := __e.Get(1)
				_ = Bound
				tmp1049 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					__e.TailApply(PrimNS2Value(symshen_4free__variable__warnings), V253, Free)
					return
				}, 1)

				tmp1050 := PrimTail(V254)

				tmp1051 := PrimHead(tmp1050)

				tmp1052 := Call(__e, PrimNS2Value(symshen_4extract__free__vars), Bound, tmp1051)

				__e.TailApply(tmp1049, tmp1052)
				return

			}, 1)

			tmp1053 := PrimHead(V254)

			tmp1054 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp1053)

			__e.TailApply(tmp1048, tmp1054)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4free__variable__check)
			return
		}

	}, 2)

	tmp1065 := Call(__e, PrimNS1Value(symns2_1set), symshen_4free__variable__check, tmp1046)

	_ = tmp1065

	tmp1066 := MakeNative(func(__e *ControlFlow) {
		V256 := __e.Get(1)
		_ = V256
		tmp1074 := PrimIsVariable(V256)

		if True == tmp1074 {
			__e.Return(PrimCons(V256, Nil))
			return
		} else {
			tmp1073 := PrimIsPair(V256)

			if True == tmp1073 {
				tmp1069 := PrimHead(V256)

				tmp1070 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp1069)

				tmp1071 := PrimTail(V256)

				tmp1072 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp1071)

				__e.TailApply(PrimNS2Value(symunion), tmp1070, tmp1072)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp1075 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract__vars, tmp1066)

	_ = tmp1075

	tmp1076 := MakeNative(func(__e *ControlFlow) {
		V268 := __e.Get(1)
		_ = V268
		V269 := __e.Get(2)
		_ = V269
		tmp1170 := PrimIsPair(V269)

		var ifres1157 Obj

		if True == tmp1170 {
			tmp1168 := PrimTail(V269)

			tmp1169 := PrimIsPair(tmp1168)

			var ifres1159 Obj

			if True == tmp1169 {
				tmp1165 := PrimTail(V269)

				tmp1166 := PrimTail(tmp1165)

				tmp1167 := PrimEqual(Nil, tmp1166)

				var ifres1161 Obj

				if True == tmp1167 {
					tmp1163 := PrimHead(V269)

					tmp1164 := PrimEqual(tmp1163, symprotect)

					var ifres1162 Obj

					if True == tmp1164 {
						ifres1162 = True

					} else {
						ifres1162 = False

					}

					ifres1161 = ifres1162

				} else {
					ifres1161 = False

				}

				var ifres1160 Obj

				if True == ifres1161 {
					ifres1160 = True

				} else {
					ifres1160 = False

				}

				ifres1159 = ifres1160

			} else {
				ifres1159 = False

			}

			var ifres1158 Obj

			if True == ifres1159 {
				ifres1158 = True

			} else {
				ifres1158 = False

			}

			ifres1157 = ifres1158

		} else {
			ifres1157 = False

		}

		if True == ifres1157 {
			__e.Return(Nil)
			return
		} else {
			tmp1156 := PrimIsVariable(V269)

			var ifres1152 Obj

			if True == tmp1156 {
				tmp1154 := Call(__e, PrimNS2Value(symelement_2), V269, V268)

				tmp1155 := PrimNot(tmp1154)

				var ifres1153 Obj

				if True == tmp1155 {
					ifres1153 = True

				} else {
					ifres1153 = False

				}

				ifres1152 = ifres1153

			} else {
				ifres1152 = False

			}

			if True == ifres1152 {
				__e.Return(PrimCons(V269, Nil))
				return
			} else {
				tmp1151 := PrimIsPair(V269)

				var ifres1132 Obj

				if True == tmp1151 {
					tmp1149 := PrimHead(V269)

					tmp1150 := PrimEqual(symlambda, tmp1149)

					var ifres1134 Obj

					if True == tmp1150 {
						tmp1147 := PrimTail(V269)

						tmp1148 := PrimIsPair(tmp1147)

						var ifres1136 Obj

						if True == tmp1148 {
							tmp1144 := PrimTail(V269)

							tmp1145 := PrimTail(tmp1144)

							tmp1146 := PrimIsPair(tmp1145)

							var ifres1138 Obj

							if True == tmp1146 {
								tmp1140 := PrimTail(V269)

								tmp1141 := PrimTail(tmp1140)

								tmp1142 := PrimTail(tmp1141)

								tmp1143 := PrimEqual(Nil, tmp1142)

								var ifres1139 Obj

								if True == tmp1143 {
									ifres1139 = True

								} else {
									ifres1139 = False

								}

								ifres1138 = ifres1139

							} else {
								ifres1138 = False

							}

							var ifres1137 Obj

							if True == ifres1138 {
								ifres1137 = True

							} else {
								ifres1137 = False

							}

							ifres1136 = ifres1137

						} else {
							ifres1136 = False

						}

						var ifres1135 Obj

						if True == ifres1136 {
							ifres1135 = True

						} else {
							ifres1135 = False

						}

						ifres1134 = ifres1135

					} else {
						ifres1134 = False

					}

					var ifres1133 Obj

					if True == ifres1134 {
						ifres1133 = True

					} else {
						ifres1133 = False

					}

					ifres1132 = ifres1133

				} else {
					ifres1132 = False

				}

				if True == ifres1132 {
					tmp1126 := PrimTail(V269)

					tmp1127 := PrimHead(tmp1126)

					tmp1128 := PrimCons(tmp1127, V268)

					tmp1129 := PrimTail(V269)

					tmp1130 := PrimTail(tmp1129)

					tmp1131 := PrimHead(tmp1130)

					__e.TailApply(PrimNS2Value(symshen_4extract__free__vars), tmp1128, tmp1131)
					return

				} else {
					tmp1125 := PrimIsPair(V269)

					var ifres1099 Obj

					if True == tmp1125 {
						tmp1123 := PrimHead(V269)

						tmp1124 := PrimEqual(symlet, tmp1123)

						var ifres1101 Obj

						if True == tmp1124 {
							tmp1121 := PrimTail(V269)

							tmp1122 := PrimIsPair(tmp1121)

							var ifres1103 Obj

							if True == tmp1122 {
								tmp1118 := PrimTail(V269)

								tmp1119 := PrimTail(tmp1118)

								tmp1120 := PrimIsPair(tmp1119)

								var ifres1105 Obj

								if True == tmp1120 {
									tmp1114 := PrimTail(V269)

									tmp1115 := PrimTail(tmp1114)

									tmp1116 := PrimTail(tmp1115)

									tmp1117 := PrimIsPair(tmp1116)

									var ifres1107 Obj

									if True == tmp1117 {
										tmp1109 := PrimTail(V269)

										tmp1110 := PrimTail(tmp1109)

										tmp1111 := PrimTail(tmp1110)

										tmp1112 := PrimTail(tmp1111)

										tmp1113 := PrimEqual(Nil, tmp1112)

										var ifres1108 Obj

										if True == tmp1113 {
											ifres1108 = True

										} else {
											ifres1108 = False

										}

										ifres1107 = ifres1108

									} else {
										ifres1107 = False

									}

									var ifres1106 Obj

									if True == ifres1107 {
										ifres1106 = True

									} else {
										ifres1106 = False

									}

									ifres1105 = ifres1106

								} else {
									ifres1105 = False

								}

								var ifres1104 Obj

								if True == ifres1105 {
									ifres1104 = True

								} else {
									ifres1104 = False

								}

								ifres1103 = ifres1104

							} else {
								ifres1103 = False

							}

							var ifres1102 Obj

							if True == ifres1103 {
								ifres1102 = True

							} else {
								ifres1102 = False

							}

							ifres1101 = ifres1102

						} else {
							ifres1101 = False

						}

						var ifres1100 Obj

						if True == ifres1101 {
							ifres1100 = True

						} else {
							ifres1100 = False

						}

						ifres1099 = ifres1100

					} else {
						ifres1099 = False

					}

					if True == ifres1099 {
						tmp1087 := PrimTail(V269)

						tmp1088 := PrimTail(tmp1087)

						tmp1089 := PrimHead(tmp1088)

						tmp1090 := Call(__e, PrimNS2Value(symshen_4extract__free__vars), V268, tmp1089)

						tmp1091 := PrimTail(V269)

						tmp1092 := PrimHead(tmp1091)

						tmp1093 := PrimCons(tmp1092, V268)

						tmp1094 := PrimTail(V269)

						tmp1095 := PrimTail(tmp1094)

						tmp1096 := PrimTail(tmp1095)

						tmp1097 := PrimHead(tmp1096)

						tmp1098 := Call(__e, PrimNS2Value(symshen_4extract__free__vars), tmp1093, tmp1097)

						__e.TailApply(PrimNS2Value(symunion), tmp1090, tmp1098)
						return

					} else {
						tmp1086 := PrimIsPair(V269)

						if True == tmp1086 {
							tmp1082 := PrimHead(V269)

							tmp1083 := Call(__e, PrimNS2Value(symshen_4extract__free__vars), V268, tmp1082)

							tmp1084 := PrimTail(V269)

							tmp1085 := Call(__e, PrimNS2Value(symshen_4extract__free__vars), V268, tmp1084)

							__e.TailApply(PrimNS2Value(symunion), tmp1083, tmp1085)
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

	tmp1171 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract__free__vars, tmp1076)

	_ = tmp1171

	tmp1172 := MakeNative(func(__e *ControlFlow) {
		V274 := __e.Get(1)
		_ = V274
		V275 := __e.Get(2)
		_ = V275
		tmp1179 := PrimEqual(Nil, V275)

		if True == tmp1179 {
			__e.Return(sym__)
			return
		} else {
			tmp1174 := Call(__e, PrimNS2Value(symshen_4list__variables), V275)

			tmp1175 := Call(__e, PrimNS2Value(symshen_4app), tmp1174, MakeString(""), symshen_4a)

			tmp1176 := PrimStringConcat(MakeString(": "), tmp1175)

			tmp1177 := Call(__e, PrimNS2Value(symshen_4app), V274, tmp1176, symshen_4a)

			tmp1178 := PrimStringConcat(MakeString("error: the following variables are free in "), tmp1177)

			__e.Return(PrimSimpleError(tmp1178))
			return

		}

	}, 2)

	tmp1180 := Call(__e, PrimNS1Value(symns2_1set), symshen_4free__variable__warnings, tmp1172)

	_ = tmp1180

	tmp1181 := MakeNative(func(__e *ControlFlow) {
		V277 := __e.Get(1)
		_ = V277
		tmp1196 := PrimIsPair(V277)

		var ifres1192 Obj

		if True == tmp1196 {
			tmp1194 := PrimTail(V277)

			tmp1195 := PrimEqual(Nil, tmp1194)

			var ifres1193 Obj

			if True == tmp1195 {
				ifres1193 = True

			} else {
				ifres1193 = False

			}

			ifres1192 = ifres1193

		} else {
			ifres1192 = False

		}

		if True == ifres1192 {
			tmp1190 := PrimHead(V277)

			tmp1191 := PrimStr(tmp1190)

			__e.Return(PrimStringConcat(tmp1191, MakeString(".")))
			return

		} else {
			tmp1189 := PrimIsPair(V277)

			if True == tmp1189 {
				tmp1184 := PrimHead(V277)

				tmp1185 := PrimStr(tmp1184)

				tmp1186 := PrimTail(V277)

				tmp1187 := Call(__e, PrimNS2Value(symshen_4list__variables), tmp1186)

				tmp1188 := PrimStringConcat(MakeString(", "), tmp1187)

				__e.Return(PrimStringConcat(tmp1185, tmp1188))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4list__variables)
				return
			}

		}

	}, 1)

	tmp1197 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list__variables, tmp1181)

	_ = tmp1197

	tmp1198 := MakeNative(func(__e *ControlFlow) {
		V279 := __e.Get(1)
		_ = V279
		tmp1218 := PrimIsPair(V279)

		var ifres1205 Obj

		if True == tmp1218 {
			tmp1216 := PrimTail(V279)

			tmp1217 := PrimIsPair(tmp1216)

			var ifres1207 Obj

			if True == tmp1217 {
				tmp1213 := PrimTail(V279)

				tmp1214 := PrimTail(tmp1213)

				tmp1215 := PrimEqual(Nil, tmp1214)

				var ifres1209 Obj

				if True == tmp1215 {
					tmp1211 := PrimHead(V279)

					tmp1212 := PrimEqual(tmp1211, symprotect)

					var ifres1210 Obj

					if True == tmp1212 {
						ifres1210 = True

					} else {
						ifres1210 = False

					}

					ifres1209 = ifres1210

				} else {
					ifres1209 = False

				}

				var ifres1208 Obj

				if True == ifres1209 {
					ifres1208 = True

				} else {
					ifres1208 = False

				}

				ifres1207 = ifres1208

			} else {
				ifres1207 = False

			}

			var ifres1206 Obj

			if True == ifres1207 {
				ifres1206 = True

			} else {
				ifres1206 = False

			}

			ifres1205 = ifres1206

		} else {
			ifres1205 = False

		}

		if True == ifres1205 {
			tmp1203 := PrimTail(V279)

			tmp1204 := PrimHead(tmp1203)

			__e.TailApply(PrimNS2Value(symshen_4strip_1protect), tmp1204)
			return

		} else {
			tmp1202 := PrimIsPair(V279)

			if True == tmp1202 {
				tmp1201 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4strip_1protect), Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp1201, V279)
				return

			} else {
				__e.Return(V279)
				return
			}

		}

	}, 1)

	tmp1219 := Call(__e, PrimNS1Value(symns2_1set), symshen_4strip_1protect, tmp1198)

	_ = tmp1219

	tmp1220 := MakeNative(func(__e *ControlFlow) {
		V281 := __e.Get(1)
		_ = V281
		tmp1236 := PrimIsPair(V281)

		var ifres1227 Obj

		if True == tmp1236 {
			tmp1234 := PrimTail(V281)

			tmp1235 := PrimIsPair(tmp1234)

			var ifres1229 Obj

			if True == tmp1235 {
				tmp1231 := PrimTail(V281)

				tmp1232 := PrimTail(tmp1231)

				tmp1233 := PrimEqual(Nil, tmp1232)

				var ifres1230 Obj

				if True == tmp1233 {
					ifres1230 = True

				} else {
					ifres1230 = False

				}

				ifres1229 = ifres1230

			} else {
				ifres1229 = False

			}

			var ifres1228 Obj

			if True == ifres1229 {
				ifres1228 = True

			} else {
				ifres1228 = False

			}

			ifres1227 = ifres1228

		} else {
			ifres1227 = False

		}

		if True == ifres1227 {
			tmp1222 := PrimHead(V281)

			tmp1223 := Call(__e, PrimNS2Value(symshen_4flatten), tmp1222)

			tmp1224 := PrimHead(V281)

			tmp1225 := PrimTail(V281)

			tmp1226 := PrimHead(tmp1225)

			__e.TailApply(PrimNS2Value(symshen_4linearise__help), tmp1223, tmp1224, tmp1226)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4linearise)
			return
		}

	}, 1)

	tmp1237 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise, tmp1220)

	_ = tmp1237

	tmp1238 := MakeNative(func(__e *ControlFlow) {
		V283 := __e.Get(1)
		_ = V283
		tmp1246 := PrimEqual(Nil, V283)

		if True == tmp1246 {
			__e.Return(Nil)
			return
		} else {
			tmp1245 := PrimIsPair(V283)

			if True == tmp1245 {
				tmp1241 := PrimHead(V283)

				tmp1242 := Call(__e, PrimNS2Value(symshen_4flatten), tmp1241)

				tmp1243 := PrimTail(V283)

				tmp1244 := Call(__e, PrimNS2Value(symshen_4flatten), tmp1243)

				__e.TailApply(PrimNS2Value(symappend), tmp1242, tmp1244)
				return

			} else {
				__e.Return(PrimCons(V283, Nil))
				return
			}

		}

	}, 1)

	tmp1247 := Call(__e, PrimNS1Value(symns2_1set), symshen_4flatten, tmp1238)

	_ = tmp1247

	tmp1248 := MakeNative(func(__e *ControlFlow) {
		V287 := __e.Get(1)
		_ = V287
		V288 := __e.Get(2)
		_ = V288
		V289 := __e.Get(3)
		_ = V289
		tmp1277 := PrimEqual(Nil, V287)

		if True == tmp1277 {
			tmp1276 := PrimCons(V289, Nil)

			__e.Return(PrimCons(V288, tmp1276))
			return

		} else {
			tmp1275 := PrimIsPair(V287)

			if True == tmp1275 {
				tmp1273 := PrimHead(V287)

				tmp1274 := PrimIsVariable(tmp1273)

				var ifres1268 Obj

				if True == tmp1274 {
					tmp1270 := PrimHead(V287)

					tmp1271 := PrimTail(V287)

					tmp1272 := Call(__e, PrimNS2Value(symelement_2), tmp1270, tmp1271)

					var ifres1269 Obj

					if True == tmp1272 {
						ifres1269 = True

					} else {
						ifres1269 = False

					}

					ifres1268 = ifres1269

				} else {
					ifres1268 = False

				}

				if True == ifres1268 {
					tmp1253 := MakeNative(func(__e *ControlFlow) {
						Var := __e.Get(1)
						_ = Var
						tmp1254 := MakeNative(func(__e *ControlFlow) {
							NewAction := __e.Get(1)
							_ = NewAction
							tmp1255 := MakeNative(func(__e *ControlFlow) {
								NewPatts := __e.Get(1)
								_ = NewPatts
								tmp1256 := PrimTail(V287)

								__e.TailApply(PrimNS2Value(symshen_4linearise__help), tmp1256, NewPatts, NewAction)
								return

							}, 1)

							tmp1257 := PrimHead(V287)

							tmp1258 := Call(__e, PrimNS2Value(symshen_4linearise__X), tmp1257, Var, V288)

							__e.TailApply(tmp1255, tmp1258)
							return

						}, 1)

						tmp1259 := PrimHead(V287)

						tmp1260 := PrimCons(Var, Nil)

						tmp1261 := PrimCons(tmp1259, tmp1260)

						tmp1262 := PrimCons(sym_a, tmp1261)

						tmp1263 := PrimCons(V289, Nil)

						tmp1264 := PrimCons(tmp1262, tmp1263)

						tmp1265 := PrimCons(symwhere, tmp1264)

						__e.TailApply(tmp1254, tmp1265)
						return

					}, 1)

					tmp1266 := PrimHead(V287)

					tmp1267 := Call(__e, PrimNS2Value(symgensym), tmp1266)

					__e.TailApply(tmp1253, tmp1267)
					return

				} else {
					tmp1252 := PrimTail(V287)

					__e.TailApply(PrimNS2Value(symshen_4linearise__help), tmp1252, V288, V289)
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4linearise__help)
				return
			}

		}

	}, 3)

	tmp1278 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise__help, tmp1248)

	_ = tmp1278

	tmp1279 := MakeNative(func(__e *ControlFlow) {
		V302 := __e.Get(1)
		_ = V302
		V303 := __e.Get(2)
		_ = V303
		V304 := __e.Get(3)
		_ = V304
		tmp1293 := PrimEqual(V304, V302)

		if True == tmp1293 {
			__e.Return(V303)
			return
		} else {
			tmp1292 := PrimIsPair(V304)

			if True == tmp1292 {
				tmp1282 := MakeNative(func(__e *ControlFlow) {
					L := __e.Get(1)
					_ = L
					tmp1288 := PrimHead(V304)

					tmp1289 := PrimEqual(L, tmp1288)

					if True == tmp1289 {
						tmp1285 := PrimHead(V304)

						tmp1286 := PrimTail(V304)

						tmp1287 := Call(__e, PrimNS2Value(symshen_4linearise__X), V302, V303, tmp1286)

						__e.Return(PrimCons(tmp1285, tmp1287))
						return

					} else {
						tmp1284 := PrimTail(V304)

						__e.Return(PrimCons(L, tmp1284))
						return

					}

				}, 1)

				tmp1290 := PrimHead(V304)

				tmp1291 := Call(__e, PrimNS2Value(symshen_4linearise__X), V302, V303, tmp1290)

				__e.TailApply(tmp1282, tmp1291)
				return

			} else {
				__e.Return(V304)
				return
			}

		}

	}, 3)

	tmp1294 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise__X, tmp1279)

	_ = tmp1294

	tmp1295 := MakeNative(func(__e *ControlFlow) {
		V307 := __e.Get(1)
		_ = V307
		V308 := __e.Get(2)
		_ = V308
		tmp1379 := PrimIsPair(V308)

		var ifres1360 Obj

		if True == tmp1379 {
			tmp1377 := PrimHead(V308)

			tmp1378 := PrimIsPair(tmp1377)

			var ifres1362 Obj

			if True == tmp1378 {
				tmp1374 := PrimHead(V308)

				tmp1375 := PrimTail(tmp1374)

				tmp1376 := PrimIsPair(tmp1375)

				var ifres1364 Obj

				if True == tmp1376 {
					tmp1370 := PrimHead(V308)

					tmp1371 := PrimTail(tmp1370)

					tmp1372 := PrimTail(tmp1371)

					tmp1373 := PrimEqual(Nil, tmp1372)

					var ifres1366 Obj

					if True == tmp1373 {
						tmp1368 := PrimTail(V308)

						tmp1369 := PrimEqual(Nil, tmp1368)

						var ifres1367 Obj

						if True == tmp1369 {
							ifres1367 = True

						} else {
							ifres1367 = False

						}

						ifres1366 = ifres1367

					} else {
						ifres1366 = False

					}

					var ifres1365 Obj

					if True == ifres1366 {
						ifres1365 = True

					} else {
						ifres1365 = False

					}

					ifres1364 = ifres1365

				} else {
					ifres1364 = False

				}

				var ifres1363 Obj

				if True == ifres1364 {
					ifres1363 = True

				} else {
					ifres1363 = False

				}

				ifres1362 = ifres1363

			} else {
				ifres1362 = False

			}

			var ifres1361 Obj

			if True == ifres1362 {
				ifres1361 = True

			} else {
				ifres1361 = False

			}

			ifres1360 = ifres1361

		} else {
			ifres1360 = False

		}

		if True == ifres1360 {
			tmp1352 := PrimHead(V308)

			tmp1353 := PrimTail(tmp1352)

			tmp1354 := PrimHead(tmp1353)

			tmp1355 := Call(__e, PrimNS2Value(symshen_4aritycheck_1action), tmp1354)

			_ = tmp1355

			tmp1356 := Call(__e, PrimNS2Value(symarity), V307)

			tmp1357 := PrimHead(V308)

			tmp1358 := PrimHead(tmp1357)

			tmp1359 := Call(__e, PrimNS2Value(symlength), tmp1358)

			__e.TailApply(PrimNS2Value(symshen_4aritycheck_1name), V307, tmp1356, tmp1359)
			return

		} else {
			tmp1351 := PrimIsPair(V308)

			var ifres1314 Obj

			if True == tmp1351 {
				tmp1349 := PrimHead(V308)

				tmp1350 := PrimIsPair(tmp1349)

				var ifres1316 Obj

				if True == tmp1350 {
					tmp1346 := PrimHead(V308)

					tmp1347 := PrimTail(tmp1346)

					tmp1348 := PrimIsPair(tmp1347)

					var ifres1318 Obj

					if True == tmp1348 {
						tmp1342 := PrimHead(V308)

						tmp1343 := PrimTail(tmp1342)

						tmp1344 := PrimTail(tmp1343)

						tmp1345 := PrimEqual(Nil, tmp1344)

						var ifres1320 Obj

						if True == tmp1345 {
							tmp1340 := PrimTail(V308)

							tmp1341 := PrimIsPair(tmp1340)

							var ifres1322 Obj

							if True == tmp1341 {
								tmp1337 := PrimTail(V308)

								tmp1338 := PrimHead(tmp1337)

								tmp1339 := PrimIsPair(tmp1338)

								var ifres1324 Obj

								if True == tmp1339 {
									tmp1333 := PrimTail(V308)

									tmp1334 := PrimHead(tmp1333)

									tmp1335 := PrimTail(tmp1334)

									tmp1336 := PrimIsPair(tmp1335)

									var ifres1326 Obj

									if True == tmp1336 {
										tmp1328 := PrimTail(V308)

										tmp1329 := PrimHead(tmp1328)

										tmp1330 := PrimTail(tmp1329)

										tmp1331 := PrimTail(tmp1330)

										tmp1332 := PrimEqual(Nil, tmp1331)

										var ifres1327 Obj

										if True == tmp1332 {
											ifres1327 = True

										} else {
											ifres1327 = False

										}

										ifres1326 = ifres1327

									} else {
										ifres1326 = False

									}

									var ifres1325 Obj

									if True == ifres1326 {
										ifres1325 = True

									} else {
										ifres1325 = False

									}

									ifres1324 = ifres1325

								} else {
									ifres1324 = False

								}

								var ifres1323 Obj

								if True == ifres1324 {
									ifres1323 = True

								} else {
									ifres1323 = False

								}

								ifres1322 = ifres1323

							} else {
								ifres1322 = False

							}

							var ifres1321 Obj

							if True == ifres1322 {
								ifres1321 = True

							} else {
								ifres1321 = False

							}

							ifres1320 = ifres1321

						} else {
							ifres1320 = False

						}

						var ifres1319 Obj

						if True == ifres1320 {
							ifres1319 = True

						} else {
							ifres1319 = False

						}

						ifres1318 = ifres1319

					} else {
						ifres1318 = False

					}

					var ifres1317 Obj

					if True == ifres1318 {
						ifres1317 = True

					} else {
						ifres1317 = False

					}

					ifres1316 = ifres1317

				} else {
					ifres1316 = False

				}

				var ifres1315 Obj

				if True == ifres1316 {
					ifres1315 = True

				} else {
					ifres1315 = False

				}

				ifres1314 = ifres1315

			} else {
				ifres1314 = False

			}

			if True == ifres1314 {
				tmp1306 := PrimHead(V308)

				tmp1307 := PrimHead(tmp1306)

				tmp1308 := Call(__e, PrimNS2Value(symlength), tmp1307)

				tmp1309 := PrimTail(V308)

				tmp1310 := PrimHead(tmp1309)

				tmp1311 := PrimHead(tmp1310)

				tmp1312 := Call(__e, PrimNS2Value(symlength), tmp1311)

				tmp1313 := PrimEqual(tmp1308, tmp1312)

				if True == tmp1313 {
					tmp1301 := PrimHead(V308)

					tmp1302 := PrimTail(tmp1301)

					tmp1303 := PrimHead(tmp1302)

					tmp1304 := Call(__e, PrimNS2Value(symshen_4aritycheck_1action), tmp1303)

					_ = tmp1304

					tmp1305 := PrimTail(V308)

					__e.TailApply(PrimNS2Value(symshen_4aritycheck), V307, tmp1305)
					return

				} else {
					tmp1299 := Call(__e, PrimNS2Value(symshen_4app), V307, MakeString("\n"), symshen_4a)

					tmp1300 := PrimStringConcat(MakeString("arity error in "), tmp1299)

					__e.Return(PrimSimpleError(tmp1300))
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4aritycheck)
				return
			}

		}

	}, 2)

	tmp1380 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck, tmp1295)

	_ = tmp1380

	tmp1381 := MakeNative(func(__e *ControlFlow) {
		V321 := __e.Get(1)
		_ = V321
		V322 := __e.Get(2)
		_ = V322
		V323 := __e.Get(3)
		_ = V323
		tmp1389 := PrimEqual(MakeNumber(-1), V322)

		if True == tmp1389 {
			__e.Return(V323)
			return
		} else {
			tmp1388 := PrimEqual(V323, V322)

			if True == tmp1388 {
				__e.Return(V323)
				return
			} else {
				tmp1384 := Call(__e, PrimNS2Value(symshen_4app), V321, MakeString(" can cause errors.\n"), symshen_4a)

				tmp1385 := PrimStringConcat(MakeString("\nwarning: changing the arity of "), tmp1384)

				tmp1386 := Call(__e, PrimNS2Value(symstoutput))

				tmp1387 := Call(__e, PrimNS2Value(symshen_4prhush), tmp1385, tmp1386)

				_ = tmp1387

				__e.Return(V323)
				return

			}

		}

	}, 3)

	tmp1390 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck_1name, tmp1381)

	_ = tmp1390

	tmp1391 := MakeNative(func(__e *ControlFlow) {
		V329 := __e.Get(1)
		_ = V329
		tmp1397 := PrimIsPair(V329)

		if True == tmp1397 {
			tmp1393 := PrimHead(V329)

			tmp1394 := PrimTail(V329)

			tmp1395 := Call(__e, PrimNS2Value(symshen_4aah), tmp1393, tmp1394)

			_ = tmp1395

			tmp1396 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4aritycheck_1action), Y)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symshen_4for_1each), tmp1396, V329)
			return

		} else {
			__e.Return(symshen_4skip)
			return
		}

	}, 1)

	tmp1398 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aritycheck_1action, tmp1391)

	_ = tmp1398

	tmp1399 := MakeNative(func(__e *ControlFlow) {
		V332 := __e.Get(1)
		_ = V332
		V333 := __e.Get(2)
		_ = V333
		tmp1400 := MakeNative(func(__e *ControlFlow) {
			Arity := __e.Get(1)
			_ = Arity
			tmp1401 := MakeNative(func(__e *ControlFlow) {
				Len := __e.Get(1)
				_ = Len
				tmp1415 := PrimGreatThan(Arity, MakeNumber(-1))

				var ifres1412 Obj

				if True == tmp1415 {
					tmp1414 := PrimGreatThan(Len, Arity)

					var ifres1413 Obj

					if True == tmp1414 {
						ifres1413 = True

					} else {
						ifres1413 = False

					}

					ifres1412 = ifres1413

				} else {
					ifres1412 = False

				}

				if True == ifres1412 {
					tmp1404 := PrimGreatThan(Len, MakeNumber(1))

					var ifres1403 Obj

					if True == tmp1404 {
						ifres1403 = MakeString("s")

					} else {
						ifres1403 = MakeString("")

					}

					tmp1405 := Call(__e, PrimNS2Value(symshen_4app), ifres1403, MakeString(".\n"), symshen_4a)

					tmp1406 := PrimStringConcat(MakeString(" argument"), tmp1405)

					tmp1407 := Call(__e, PrimNS2Value(symshen_4app), Len, tmp1406, symshen_4a)

					tmp1408 := PrimStringConcat(MakeString(" might not like "), tmp1407)

					tmp1409 := Call(__e, PrimNS2Value(symshen_4app), V332, tmp1408, symshen_4a)

					tmp1410 := PrimStringConcat(MakeString("warning: "), tmp1409)

					tmp1411 := Call(__e, PrimNS2Value(symstoutput))

					__e.TailApply(PrimNS2Value(symshen_4prhush), tmp1410, tmp1411)
					return

				} else {
					__e.Return(symshen_4skip)
					return
				}

			}, 1)

			tmp1416 := Call(__e, PrimNS2Value(symlength), V333)

			__e.TailApply(tmp1401, tmp1416)
			return

		}, 1)

		tmp1417 := Call(__e, PrimNS2Value(symarity), V332)

		__e.TailApply(tmp1400, tmp1417)
		return

	}, 2)

	tmp1418 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aah, tmp1399)

	_ = tmp1418

	tmp1419 := MakeNative(func(__e *ControlFlow) {
		V335 := __e.Get(1)
		_ = V335
		tmp1433 := PrimIsPair(V335)

		var ifres1424 Obj

		if True == tmp1433 {
			tmp1431 := PrimTail(V335)

			tmp1432 := PrimIsPair(tmp1431)

			var ifres1426 Obj

			if True == tmp1432 {
				tmp1428 := PrimTail(V335)

				tmp1429 := PrimTail(tmp1428)

				tmp1430 := PrimEqual(Nil, tmp1429)

				var ifres1427 Obj

				if True == tmp1430 {
					ifres1427 = True

				} else {
					ifres1427 = False

				}

				ifres1426 = ifres1427

			} else {
				ifres1426 = False

			}

			var ifres1425 Obj

			if True == ifres1426 {
				ifres1425 = True

			} else {
				ifres1425 = False

			}

			ifres1424 = ifres1425

		} else {
			ifres1424 = False

		}

		if True == ifres1424 {
			tmp1421 := PrimHead(V335)

			tmp1422 := PrimTail(V335)

			tmp1423 := PrimHead(tmp1422)

			__e.TailApply(PrimNS2Value(symshen_4abstraction__build), tmp1421, tmp1423)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4abstract__rule)
			return
		}

	}, 1)

	tmp1434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abstract__rule, tmp1419)

	_ = tmp1434

	tmp1435 := MakeNative(func(__e *ControlFlow) {
		V338 := __e.Get(1)
		_ = V338
		V339 := __e.Get(2)
		_ = V339
		tmp1444 := PrimEqual(Nil, V338)

		if True == tmp1444 {
			__e.Return(V339)
			return
		} else {
			tmp1443 := PrimIsPair(V338)

			if True == tmp1443 {
				tmp1438 := PrimHead(V338)

				tmp1439 := PrimTail(V338)

				tmp1440 := Call(__e, PrimNS2Value(symshen_4abstraction__build), tmp1439, V339)

				tmp1441 := PrimCons(tmp1440, Nil)

				tmp1442 := PrimCons(tmp1438, tmp1441)

				__e.Return(PrimCons(sym_c_4, tmp1442))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4abstraction__build)
				return
			}

		}

	}, 2)

	tmp1445 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abstraction__build, tmp1435)

	_ = tmp1445

	tmp1446 := MakeNative(func(__e *ControlFlow) {
		V341 := __e.Get(1)
		_ = V341
		tmp1451 := PrimEqual(MakeNumber(0), V341)

		if True == tmp1451 {
			__e.Return(Nil)
			return
		} else {
			tmp1448 := Call(__e, PrimNS2Value(symgensym), symV)

			tmp1449 := PrimNumberSubtract(V341, MakeNumber(1))

			tmp1450 := Call(__e, PrimNS2Value(symshen_4parameters), tmp1449)

			__e.Return(PrimCons(tmp1448, tmp1450))
			return

		}

	}, 1)

	tmp1452 := Call(__e, PrimNS1Value(symns2_1set), symshen_4parameters, tmp1446)

	_ = tmp1452

	tmp1453 := MakeNative(func(__e *ControlFlow) {
		V344 := __e.Get(1)
		_ = V344
		V345 := __e.Get(2)
		_ = V345
		tmp1461 := PrimEqual(Nil, V344)

		if True == tmp1461 {
			__e.Return(V345)
			return
		} else {
			tmp1460 := PrimIsPair(V344)

			if True == tmp1460 {
				tmp1456 := PrimTail(V344)

				tmp1457 := PrimHead(V344)

				tmp1458 := PrimCons(tmp1457, Nil)

				tmp1459 := PrimCons(V345, tmp1458)

				__e.TailApply(PrimNS2Value(symshen_4application__build), tmp1456, tmp1459)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4application__build)
				return
			}

		}

	}, 2)

	tmp1462 := Call(__e, PrimNS1Value(symns2_1set), symshen_4application__build, tmp1453)

	_ = tmp1462

	tmp1463 := MakeNative(func(__e *ControlFlow) {
		V348 := __e.Get(1)
		_ = V348
		V349 := __e.Get(2)
		_ = V349
		tmp1501 := PrimIsPair(V349)

		var ifres1492 Obj

		if True == tmp1501 {
			tmp1499 := PrimTail(V349)

			tmp1500 := PrimIsPair(tmp1499)

			var ifres1494 Obj

			if True == tmp1500 {
				tmp1496 := PrimTail(V349)

				tmp1497 := PrimTail(tmp1496)

				tmp1498 := PrimEqual(Nil, tmp1497)

				var ifres1495 Obj

				if True == tmp1498 {
					ifres1495 = True

				} else {
					ifres1495 = False

				}

				ifres1494 = ifres1495

			} else {
				ifres1494 = False

			}

			var ifres1493 Obj

			if True == ifres1494 {
				ifres1493 = True

			} else {
				ifres1493 = False

			}

			ifres1492 = ifres1493

		} else {
			ifres1492 = False

		}

		if True == ifres1492 {
			tmp1465 := MakeNative(func(__e *ControlFlow) {
				Arity := __e.Get(1)
				_ = Arity
				tmp1466 := MakeNative(func(__e *ControlFlow) {
					Reduce := __e.Get(1)
					_ = Reduce
					tmp1467 := MakeNative(func(__e *ControlFlow) {
						CondExpression := __e.Get(1)
						_ = CondExpression
						tmp1468 := MakeNative(func(__e *ControlFlow) {
							TypeTable := __e.Get(1)
							_ = TypeTable
							tmp1469 := MakeNative(func(__e *ControlFlow) {
								TypedCondExpression := __e.Get(1)
								_ = TypedCondExpression
								tmp1470 := PrimHead(V349)

								tmp1471 := PrimCons(TypedCondExpression, Nil)

								tmp1472 := PrimCons(tmp1470, tmp1471)

								tmp1473 := PrimCons(V348, tmp1472)

								__e.Return(PrimCons(symdefun, tmp1473))
								return

							}, 1)

							tmp1477 := PrimNS3Value(symshen_4_doptimise_d)

							var ifres1474 Obj

							if True == tmp1477 {
								tmp1475 := PrimHead(V349)

								tmp1476 := Call(__e, PrimNS2Value(symshen_4assign_1types), tmp1475, TypeTable, CondExpression)

								ifres1474 = tmp1476

							} else {
								ifres1474 = CondExpression

							}

							__e.TailApply(tmp1469, ifres1474)
							return

						}, 1)

						tmp1482 := PrimNS3Value(symshen_4_doptimise_d)

						var ifres1478 Obj

						if True == tmp1482 {
							tmp1479 := Call(__e, PrimNS2Value(symshen_4get_1type), V348)

							tmp1480 := PrimHead(V349)

							tmp1481 := Call(__e, PrimNS2Value(symshen_4typextable), tmp1479, tmp1480)

							ifres1478 = tmp1481

						} else {
							ifres1478 = symshen_4skip

						}

						__e.TailApply(tmp1468, ifres1478)
						return

					}, 1)

					tmp1483 := PrimHead(V349)

					tmp1484 := Call(__e, PrimNS2Value(symshen_4cond_1expression), V348, tmp1483, Reduce)

					__e.TailApply(tmp1467, tmp1484)
					return

				}, 1)

				tmp1485 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4reduce), X)
					return
				}, 1)

				tmp1486 := PrimTail(V349)

				tmp1487 := PrimHead(tmp1486)

				tmp1488 := Call(__e, PrimNS2Value(symmap), tmp1485, tmp1487)

				__e.TailApply(tmp1466, tmp1488)
				return

			}, 1)

			tmp1489 := PrimHead(V349)

			tmp1490 := Call(__e, PrimNS2Value(symlength), tmp1489)

			tmp1491 := Call(__e, PrimNS2Value(symshen_4store_1arity), V348, tmp1490)

			__e.TailApply(tmp1465, tmp1491)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4compile__to__kl)
			return
		}

	}, 2)

	tmp1502 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__to__kl, tmp1463)

	_ = tmp1502

	tmp1503 := MakeNative(func(__e *ControlFlow) {
		V355 := __e.Get(1)
		_ = V355
		tmp1510 := PrimIsPair(V355)

		if True == tmp1510 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp1505 := MakeNative(func(__e *ControlFlow) {
				FType := __e.Get(1)
				_ = FType
				tmp1507 := Call(__e, PrimNS2Value(symempty_2), FType)

				if True == tmp1507 {
					__e.Return(symshen_4skip)
					return
				} else {
					__e.Return(PrimTail(FType))
					return
				}

			}, 1)

			tmp1508 := PrimNS3Value(symshen_4_dsignedfuncs_d)

			tmp1509 := Call(__e, PrimNS2Value(symassoc), V355, tmp1508)

			__e.TailApply(tmp1505, tmp1509)
			return

		}

	}, 1)

	tmp1511 := Call(__e, PrimNS1Value(symns2_1set), symshen_4get_1type, tmp1503)

	_ = tmp1511

	tmp1512 := MakeNative(func(__e *ControlFlow) {
		V366 := __e.Get(1)
		_ = V366
		V367 := __e.Get(2)
		_ = V367
		tmp1552 := PrimIsPair(V366)

		var ifres1529 Obj

		if True == tmp1552 {
			tmp1550 := PrimTail(V366)

			tmp1551 := PrimIsPair(tmp1550)

			var ifres1531 Obj

			if True == tmp1551 {
				tmp1547 := PrimTail(V366)

				tmp1548 := PrimHead(tmp1547)

				tmp1549 := PrimEqual(sym_1_1_6, tmp1548)

				var ifres1533 Obj

				if True == tmp1549 {
					tmp1544 := PrimTail(V366)

					tmp1545 := PrimTail(tmp1544)

					tmp1546 := PrimIsPair(tmp1545)

					var ifres1535 Obj

					if True == tmp1546 {
						tmp1540 := PrimTail(V366)

						tmp1541 := PrimTail(tmp1540)

						tmp1542 := PrimTail(tmp1541)

						tmp1543 := PrimEqual(Nil, tmp1542)

						var ifres1537 Obj

						if True == tmp1543 {
							tmp1539 := PrimIsPair(V367)

							var ifres1538 Obj

							if True == tmp1539 {
								ifres1538 = True

							} else {
								ifres1538 = False

							}

							ifres1537 = ifres1538

						} else {
							ifres1537 = False

						}

						var ifres1536 Obj

						if True == ifres1537 {
							ifres1536 = True

						} else {
							ifres1536 = False

						}

						ifres1535 = ifres1536

					} else {
						ifres1535 = False

					}

					var ifres1534 Obj

					if True == ifres1535 {
						ifres1534 = True

					} else {
						ifres1534 = False

					}

					ifres1533 = ifres1534

				} else {
					ifres1533 = False

				}

				var ifres1532 Obj

				if True == ifres1533 {
					ifres1532 = True

				} else {
					ifres1532 = False

				}

				ifres1531 = ifres1532

			} else {
				ifres1531 = False

			}

			var ifres1530 Obj

			if True == ifres1531 {
				ifres1530 = True

			} else {
				ifres1530 = False

			}

			ifres1529 = ifres1530

		} else {
			ifres1529 = False

		}

		if True == ifres1529 {
			tmp1527 := PrimHead(V366)

			tmp1528 := PrimIsVariable(tmp1527)

			if True == tmp1528 {
				tmp1523 := PrimTail(V366)

				tmp1524 := PrimTail(tmp1523)

				tmp1525 := PrimHead(tmp1524)

				tmp1526 := PrimTail(V367)

				__e.TailApply(PrimNS2Value(symshen_4typextable), tmp1525, tmp1526)
				return

			} else {
				tmp1515 := PrimHead(V367)

				tmp1516 := PrimHead(V366)

				tmp1517 := PrimCons(tmp1515, tmp1516)

				tmp1518 := PrimTail(V366)

				tmp1519 := PrimTail(tmp1518)

				tmp1520 := PrimHead(tmp1519)

				tmp1521 := PrimTail(V367)

				tmp1522 := Call(__e, PrimNS2Value(symshen_4typextable), tmp1520, tmp1521)

				__e.Return(PrimCons(tmp1517, tmp1522))
				return

			}

		} else {
			__e.Return(Nil)
			return
		}

	}, 2)

	tmp1553 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typextable, tmp1512)

	_ = tmp1553

	tmp1554 := MakeNative(func(__e *ControlFlow) {
		V371 := __e.Get(1)
		_ = V371
		V372 := __e.Get(2)
		_ = V372
		V373 := __e.Get(3)
		_ = V373
		tmp1667 := PrimIsPair(V373)

		var ifres1641 Obj

		if True == tmp1667 {
			tmp1665 := PrimHead(V373)

			tmp1666 := PrimEqual(symlet, tmp1665)

			var ifres1643 Obj

			if True == tmp1666 {
				tmp1663 := PrimTail(V373)

				tmp1664 := PrimIsPair(tmp1663)

				var ifres1645 Obj

				if True == tmp1664 {
					tmp1660 := PrimTail(V373)

					tmp1661 := PrimTail(tmp1660)

					tmp1662 := PrimIsPair(tmp1661)

					var ifres1647 Obj

					if True == tmp1662 {
						tmp1656 := PrimTail(V373)

						tmp1657 := PrimTail(tmp1656)

						tmp1658 := PrimTail(tmp1657)

						tmp1659 := PrimIsPair(tmp1658)

						var ifres1649 Obj

						if True == tmp1659 {
							tmp1651 := PrimTail(V373)

							tmp1652 := PrimTail(tmp1651)

							tmp1653 := PrimTail(tmp1652)

							tmp1654 := PrimTail(tmp1653)

							tmp1655 := PrimEqual(Nil, tmp1654)

							var ifres1650 Obj

							if True == tmp1655 {
								ifres1650 = True

							} else {
								ifres1650 = False

							}

							ifres1649 = ifres1650

						} else {
							ifres1649 = False

						}

						var ifres1648 Obj

						if True == ifres1649 {
							ifres1648 = True

						} else {
							ifres1648 = False

						}

						ifres1647 = ifres1648

					} else {
						ifres1647 = False

					}

					var ifres1646 Obj

					if True == ifres1647 {
						ifres1646 = True

					} else {
						ifres1646 = False

					}

					ifres1645 = ifres1646

				} else {
					ifres1645 = False

				}

				var ifres1644 Obj

				if True == ifres1645 {
					ifres1644 = True

				} else {
					ifres1644 = False

				}

				ifres1643 = ifres1644

			} else {
				ifres1643 = False

			}

			var ifres1642 Obj

			if True == ifres1643 {
				ifres1642 = True

			} else {
				ifres1642 = False

			}

			ifres1641 = ifres1642

		} else {
			ifres1641 = False

		}

		if True == ifres1641 {
			tmp1624 := PrimTail(V373)

			tmp1625 := PrimHead(tmp1624)

			tmp1626 := PrimTail(V373)

			tmp1627 := PrimTail(tmp1626)

			tmp1628 := PrimHead(tmp1627)

			tmp1629 := Call(__e, PrimNS2Value(symshen_4assign_1types), V371, V372, tmp1628)

			tmp1630 := PrimTail(V373)

			tmp1631 := PrimHead(tmp1630)

			tmp1632 := PrimCons(tmp1631, V371)

			tmp1633 := PrimTail(V373)

			tmp1634 := PrimTail(tmp1633)

			tmp1635 := PrimTail(tmp1634)

			tmp1636 := PrimHead(tmp1635)

			tmp1637 := Call(__e, PrimNS2Value(symshen_4assign_1types), tmp1632, V372, tmp1636)

			tmp1638 := PrimCons(tmp1637, Nil)

			tmp1639 := PrimCons(tmp1629, tmp1638)

			tmp1640 := PrimCons(tmp1625, tmp1639)

			__e.Return(PrimCons(symlet, tmp1640))
			return

		} else {
			tmp1623 := PrimIsPair(V373)

			var ifres1604 Obj

			if True == tmp1623 {
				tmp1621 := PrimHead(V373)

				tmp1622 := PrimEqual(symlambda, tmp1621)

				var ifres1606 Obj

				if True == tmp1622 {
					tmp1619 := PrimTail(V373)

					tmp1620 := PrimIsPair(tmp1619)

					var ifres1608 Obj

					if True == tmp1620 {
						tmp1616 := PrimTail(V373)

						tmp1617 := PrimTail(tmp1616)

						tmp1618 := PrimIsPair(tmp1617)

						var ifres1610 Obj

						if True == tmp1618 {
							tmp1612 := PrimTail(V373)

							tmp1613 := PrimTail(tmp1612)

							tmp1614 := PrimTail(tmp1613)

							tmp1615 := PrimEqual(Nil, tmp1614)

							var ifres1611 Obj

							if True == tmp1615 {
								ifres1611 = True

							} else {
								ifres1611 = False

							}

							ifres1610 = ifres1611

						} else {
							ifres1610 = False

						}

						var ifres1609 Obj

						if True == ifres1610 {
							ifres1609 = True

						} else {
							ifres1609 = False

						}

						ifres1608 = ifres1609

					} else {
						ifres1608 = False

					}

					var ifres1607 Obj

					if True == ifres1608 {
						ifres1607 = True

					} else {
						ifres1607 = False

					}

					ifres1606 = ifres1607

				} else {
					ifres1606 = False

				}

				var ifres1605 Obj

				if True == ifres1606 {
					ifres1605 = True

				} else {
					ifres1605 = False

				}

				ifres1604 = ifres1605

			} else {
				ifres1604 = False

			}

			if True == ifres1604 {
				tmp1593 := PrimTail(V373)

				tmp1594 := PrimHead(tmp1593)

				tmp1595 := PrimTail(V373)

				tmp1596 := PrimHead(tmp1595)

				tmp1597 := PrimCons(tmp1596, V371)

				tmp1598 := PrimTail(V373)

				tmp1599 := PrimTail(tmp1598)

				tmp1600 := PrimHead(tmp1599)

				tmp1601 := Call(__e, PrimNS2Value(symshen_4assign_1types), tmp1597, V372, tmp1600)

				tmp1602 := PrimCons(tmp1601, Nil)

				tmp1603 := PrimCons(tmp1594, tmp1602)

				__e.Return(PrimCons(symlambda, tmp1603))
				return

			} else {
				tmp1592 := PrimIsPair(V373)

				var ifres1588 Obj

				if True == tmp1592 {
					tmp1590 := PrimHead(V373)

					tmp1591 := PrimEqual(symcond, tmp1590)

					var ifres1589 Obj

					if True == tmp1591 {
						ifres1589 = True

					} else {
						ifres1589 = False

					}

					ifres1588 = ifres1589

				} else {
					ifres1588 = False

				}

				if True == ifres1588 {
					tmp1579 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						tmp1580 := PrimHead(Y)

						tmp1581 := Call(__e, PrimNS2Value(symshen_4assign_1types), V371, V372, tmp1580)

						tmp1582 := PrimTail(Y)

						tmp1583 := PrimHead(tmp1582)

						tmp1584 := Call(__e, PrimNS2Value(symshen_4assign_1types), V371, V372, tmp1583)

						tmp1585 := PrimCons(tmp1584, Nil)

						__e.Return(PrimCons(tmp1581, tmp1585))
						return

					}, 1)

					tmp1586 := PrimTail(V373)

					tmp1587 := Call(__e, PrimNS2Value(symmap), tmp1579, tmp1586)

					__e.Return(PrimCons(symcond, tmp1587))
					return

				} else {
					tmp1578 := PrimIsPair(V373)

					if True == tmp1578 {
						tmp1568 := MakeNative(func(__e *ControlFlow) {
							NewTable := __e.Get(1)
							_ = NewTable
							tmp1569 := PrimHead(V373)

							tmp1570 := MakeNative(func(__e *ControlFlow) {
								Y := __e.Get(1)
								_ = Y
								tmp1571 := Call(__e, PrimNS2Value(symappend), V372, NewTable)

								__e.TailApply(PrimNS2Value(symshen_4assign_1types), V371, tmp1571, Y)
								return

							}, 1)

							tmp1572 := PrimTail(V373)

							tmp1573 := Call(__e, PrimNS2Value(symmap), tmp1570, tmp1572)

							__e.Return(PrimCons(tmp1569, tmp1573))
							return

						}, 1)

						tmp1574 := PrimHead(V373)

						tmp1575 := Call(__e, PrimNS2Value(symshen_4get_1type), tmp1574)

						tmp1576 := PrimTail(V373)

						tmp1577 := Call(__e, PrimNS2Value(symshen_4typextable), tmp1575, tmp1576)

						__e.TailApply(tmp1568, tmp1577)
						return

					} else {
						tmp1559 := MakeNative(func(__e *ControlFlow) {
							AtomType := __e.Get(1)
							_ = AtomType
							tmp1566 := PrimIsPair(AtomType)

							if True == tmp1566 {
								tmp1563 := PrimTail(AtomType)

								tmp1564 := PrimCons(tmp1563, Nil)

								tmp1565 := PrimCons(V373, tmp1564)

								__e.Return(PrimCons(symtype, tmp1565))
								return

							} else {
								tmp1562 := Call(__e, PrimNS2Value(symelement_2), V373, V371)

								if True == tmp1562 {
									__e.Return(V373)
									return
								} else {
									__e.TailApply(PrimNS2Value(symshen_4atom_1type), V373)
									return
								}

							}

						}, 1)

						tmp1567 := Call(__e, PrimNS2Value(symassoc), V373, V372)

						__e.TailApply(tmp1559, tmp1567)
						return

					}

				}

			}

		}

	}, 3)

	tmp1668 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assign_1types, tmp1554)

	_ = tmp1668

	tmp1669 := MakeNative(func(__e *ControlFlow) {
		V375 := __e.Get(1)
		_ = V375
		tmp1685 := PrimIsString(V375)

		if True == tmp1685 {
			tmp1683 := PrimCons(symstring, Nil)

			tmp1684 := PrimCons(V375, tmp1683)

			__e.Return(PrimCons(symtype, tmp1684))
			return

		} else {
			tmp1682 := PrimIsNumber(V375)

			if True == tmp1682 {
				tmp1680 := PrimCons(symnumber, Nil)

				tmp1681 := PrimCons(V375, tmp1680)

				__e.Return(PrimCons(symtype, tmp1681))
				return

			} else {
				tmp1679 := Call(__e, PrimNS2Value(symboolean_2), V375)

				if True == tmp1679 {
					tmp1677 := PrimCons(symboolean, Nil)

					tmp1678 := PrimCons(V375, tmp1677)

					__e.Return(PrimCons(symtype, tmp1678))
					return

				} else {
					tmp1676 := PrimIsSymbol(V375)

					if True == tmp1676 {
						tmp1674 := PrimCons(symsymbol, Nil)

						tmp1675 := PrimCons(V375, tmp1674)

						__e.Return(PrimCons(symtype, tmp1675))
						return

					} else {
						__e.Return(V375)
						return
					}

				}

			}

		}

	}, 1)

	tmp1686 := Call(__e, PrimNS1Value(symns2_1set), symshen_4atom_1type, tmp1669)

	_ = tmp1686

	tmp1687 := MakeNative(func(__e *ControlFlow) {
		V380 := __e.Get(1)
		_ = V380
		V381 := __e.Get(2)
		_ = V381
		tmp1690 := PrimNS3Value(symshen_4_dinstalling_1kl_d)

		if True == tmp1690 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp1689 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), V380, symarity, V381, tmp1689)
			return

		}

	}, 2)

	tmp1691 := Call(__e, PrimNS1Value(symns2_1set), symshen_4store_1arity, tmp1687)

	_ = tmp1691

	tmp1692 := MakeNative(func(__e *ControlFlow) {
		V383 := __e.Get(1)
		_ = V383
		tmp1693 := PrimNS3Set(symshen_4_dteststack_d, Nil)

		_ = tmp1693

		tmp1694 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp1695 := PrimNS3Value(symshen_4_dteststack_d)

			tmp1696 := Call(__e, PrimNS2Value(symreverse), tmp1695)

			tmp1697 := PrimCons(symshen_4tests, tmp1696)

			tmp1698 := PrimCons(sym_h, tmp1697)

			tmp1699 := PrimCons(Result, Nil)

			__e.Return(PrimCons(tmp1698, tmp1699))
			return

		}, 1)

		tmp1700 := Call(__e, PrimNS2Value(symshen_4reduce__help), V383)

		__e.TailApply(tmp1694, tmp1700)
		return

	}, 1)

	tmp1701 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reduce, tmp1692)

	_ = tmp1701

	tmp1702 := MakeNative(func(__e *ControlFlow) {
		V385 := __e.Get(1)
		_ = V385
		tmp2359 := PrimIsPair(V385)

		var ifres2286 Obj

		if True == tmp2359 {
			tmp2357 := PrimHead(V385)

			tmp2358 := PrimIsPair(tmp2357)

			var ifres2288 Obj

			if True == tmp2358 {
				tmp2354 := PrimHead(V385)

				tmp2355 := PrimHead(tmp2354)

				tmp2356 := PrimEqual(sym_c_4, tmp2355)

				var ifres2290 Obj

				if True == tmp2356 {
					tmp2351 := PrimHead(V385)

					tmp2352 := PrimTail(tmp2351)

					tmp2353 := PrimIsPair(tmp2352)

					var ifres2292 Obj

					if True == tmp2353 {
						tmp2347 := PrimHead(V385)

						tmp2348 := PrimTail(tmp2347)

						tmp2349 := PrimHead(tmp2348)

						tmp2350 := PrimIsPair(tmp2349)

						var ifres2294 Obj

						if True == tmp2350 {
							tmp2342 := PrimHead(V385)

							tmp2343 := PrimTail(tmp2342)

							tmp2344 := PrimHead(tmp2343)

							tmp2345 := PrimHead(tmp2344)

							tmp2346 := PrimEqual(symcons, tmp2345)

							var ifres2296 Obj

							if True == tmp2346 {
								tmp2337 := PrimHead(V385)

								tmp2338 := PrimTail(tmp2337)

								tmp2339 := PrimHead(tmp2338)

								tmp2340 := PrimTail(tmp2339)

								tmp2341 := PrimIsPair(tmp2340)

								var ifres2298 Obj

								if True == tmp2341 {
									tmp2331 := PrimHead(V385)

									tmp2332 := PrimTail(tmp2331)

									tmp2333 := PrimHead(tmp2332)

									tmp2334 := PrimTail(tmp2333)

									tmp2335 := PrimTail(tmp2334)

									tmp2336 := PrimIsPair(tmp2335)

									var ifres2300 Obj

									if True == tmp2336 {
										tmp2324 := PrimHead(V385)

										tmp2325 := PrimTail(tmp2324)

										tmp2326 := PrimHead(tmp2325)

										tmp2327 := PrimTail(tmp2326)

										tmp2328 := PrimTail(tmp2327)

										tmp2329 := PrimTail(tmp2328)

										tmp2330 := PrimEqual(Nil, tmp2329)

										var ifres2302 Obj

										if True == tmp2330 {
											tmp2320 := PrimHead(V385)

											tmp2321 := PrimTail(tmp2320)

											tmp2322 := PrimTail(tmp2321)

											tmp2323 := PrimIsPair(tmp2322)

											var ifres2304 Obj

											if True == tmp2323 {
												tmp2315 := PrimHead(V385)

												tmp2316 := PrimTail(tmp2315)

												tmp2317 := PrimTail(tmp2316)

												tmp2318 := PrimTail(tmp2317)

												tmp2319 := PrimEqual(Nil, tmp2318)

												var ifres2306 Obj

												if True == tmp2319 {
													tmp2313 := PrimTail(V385)

													tmp2314 := PrimIsPair(tmp2313)

													var ifres2308 Obj

													if True == tmp2314 {
														tmp2310 := PrimTail(V385)

														tmp2311 := PrimTail(tmp2310)

														tmp2312 := PrimEqual(Nil, tmp2311)

														var ifres2309 Obj

														if True == tmp2312 {
															ifres2309 = True

														} else {
															ifres2309 = False

														}

														ifres2308 = ifres2309

													} else {
														ifres2308 = False

													}

													var ifres2307 Obj

													if True == ifres2308 {
														ifres2307 = True

													} else {
														ifres2307 = False

													}

													ifres2306 = ifres2307

												} else {
													ifres2306 = False

												}

												var ifres2305 Obj

												if True == ifres2306 {
													ifres2305 = True

												} else {
													ifres2305 = False

												}

												ifres2304 = ifres2305

											} else {
												ifres2304 = False

											}

											var ifres2303 Obj

											if True == ifres2304 {
												ifres2303 = True

											} else {
												ifres2303 = False

											}

											ifres2302 = ifres2303

										} else {
											ifres2302 = False

										}

										var ifres2301 Obj

										if True == ifres2302 {
											ifres2301 = True

										} else {
											ifres2301 = False

										}

										ifres2300 = ifres2301

									} else {
										ifres2300 = False

									}

									var ifres2299 Obj

									if True == ifres2300 {
										ifres2299 = True

									} else {
										ifres2299 = False

									}

									ifres2298 = ifres2299

								} else {
									ifres2298 = False

								}

								var ifres2297 Obj

								if True == ifres2298 {
									ifres2297 = True

								} else {
									ifres2297 = False

								}

								ifres2296 = ifres2297

							} else {
								ifres2296 = False

							}

							var ifres2295 Obj

							if True == ifres2296 {
								ifres2295 = True

							} else {
								ifres2295 = False

							}

							ifres2294 = ifres2295

						} else {
							ifres2294 = False

						}

						var ifres2293 Obj

						if True == ifres2294 {
							ifres2293 = True

						} else {
							ifres2293 = False

						}

						ifres2292 = ifres2293

					} else {
						ifres2292 = False

					}

					var ifres2291 Obj

					if True == ifres2292 {
						ifres2291 = True

					} else {
						ifres2291 = False

					}

					ifres2290 = ifres2291

				} else {
					ifres2290 = False

				}

				var ifres2289 Obj

				if True == ifres2290 {
					ifres2289 = True

				} else {
					ifres2289 = False

				}

				ifres2288 = ifres2289

			} else {
				ifres2288 = False

			}

			var ifres2287 Obj

			if True == ifres2288 {
				ifres2287 = True

			} else {
				ifres2287 = False

			}

			ifres2286 = ifres2287

		} else {
			ifres2286 = False

		}

		if True == ifres2286 {
			tmp2246 := PrimTail(V385)

			tmp2247 := PrimCons(symcons_2, tmp2246)

			tmp2248 := Call(__e, PrimNS2Value(symshen_4add__test), tmp2247)

			_ = tmp2248

			tmp2249 := MakeNative(func(__e *ControlFlow) {
				Abstraction := __e.Get(1)
				_ = Abstraction
				tmp2250 := MakeNative(func(__e *ControlFlow) {
					Application := __e.Get(1)
					_ = Application
					__e.TailApply(PrimNS2Value(symshen_4reduce__help), Application)
					return
				}, 1)

				tmp2251 := PrimTail(V385)

				tmp2252 := PrimCons(symhd, tmp2251)

				tmp2253 := PrimCons(tmp2252, Nil)

				tmp2254 := PrimCons(Abstraction, tmp2253)

				tmp2255 := PrimTail(V385)

				tmp2256 := PrimCons(symtl, tmp2255)

				tmp2257 := PrimCons(tmp2256, Nil)

				tmp2258 := PrimCons(tmp2254, tmp2257)

				__e.TailApply(tmp2250, tmp2258)
				return

			}, 1)

			tmp2259 := PrimHead(V385)

			tmp2260 := PrimTail(tmp2259)

			tmp2261 := PrimHead(tmp2260)

			tmp2262 := PrimTail(tmp2261)

			tmp2263 := PrimHead(tmp2262)

			tmp2264 := PrimHead(V385)

			tmp2265 := PrimTail(tmp2264)

			tmp2266 := PrimHead(tmp2265)

			tmp2267 := PrimTail(tmp2266)

			tmp2268 := PrimTail(tmp2267)

			tmp2269 := PrimHead(tmp2268)

			tmp2270 := PrimTail(V385)

			tmp2271 := PrimHead(tmp2270)

			tmp2272 := PrimHead(V385)

			tmp2273 := PrimTail(tmp2272)

			tmp2274 := PrimHead(tmp2273)

			tmp2275 := PrimHead(V385)

			tmp2276 := PrimTail(tmp2275)

			tmp2277 := PrimTail(tmp2276)

			tmp2278 := PrimHead(tmp2277)

			tmp2279 := Call(__e, PrimNS2Value(symshen_4ebr), tmp2271, tmp2274, tmp2278)

			tmp2280 := PrimCons(tmp2279, Nil)

			tmp2281 := PrimCons(tmp2269, tmp2280)

			tmp2282 := PrimCons(sym_c_4, tmp2281)

			tmp2283 := PrimCons(tmp2282, Nil)

			tmp2284 := PrimCons(tmp2263, tmp2283)

			tmp2285 := PrimCons(sym_c_4, tmp2284)

			__e.TailApply(tmp2249, tmp2285)
			return

		} else {
			tmp2245 := PrimIsPair(V385)

			var ifres2172 Obj

			if True == tmp2245 {
				tmp2243 := PrimHead(V385)

				tmp2244 := PrimIsPair(tmp2243)

				var ifres2174 Obj

				if True == tmp2244 {
					tmp2240 := PrimHead(V385)

					tmp2241 := PrimHead(tmp2240)

					tmp2242 := PrimEqual(sym_c_4, tmp2241)

					var ifres2176 Obj

					if True == tmp2242 {
						tmp2237 := PrimHead(V385)

						tmp2238 := PrimTail(tmp2237)

						tmp2239 := PrimIsPair(tmp2238)

						var ifres2178 Obj

						if True == tmp2239 {
							tmp2233 := PrimHead(V385)

							tmp2234 := PrimTail(tmp2233)

							tmp2235 := PrimHead(tmp2234)

							tmp2236 := PrimIsPair(tmp2235)

							var ifres2180 Obj

							if True == tmp2236 {
								tmp2228 := PrimHead(V385)

								tmp2229 := PrimTail(tmp2228)

								tmp2230 := PrimHead(tmp2229)

								tmp2231 := PrimHead(tmp2230)

								tmp2232 := PrimEqual(sym_8p, tmp2231)

								var ifres2182 Obj

								if True == tmp2232 {
									tmp2223 := PrimHead(V385)

									tmp2224 := PrimTail(tmp2223)

									tmp2225 := PrimHead(tmp2224)

									tmp2226 := PrimTail(tmp2225)

									tmp2227 := PrimIsPair(tmp2226)

									var ifres2184 Obj

									if True == tmp2227 {
										tmp2217 := PrimHead(V385)

										tmp2218 := PrimTail(tmp2217)

										tmp2219 := PrimHead(tmp2218)

										tmp2220 := PrimTail(tmp2219)

										tmp2221 := PrimTail(tmp2220)

										tmp2222 := PrimIsPair(tmp2221)

										var ifres2186 Obj

										if True == tmp2222 {
											tmp2210 := PrimHead(V385)

											tmp2211 := PrimTail(tmp2210)

											tmp2212 := PrimHead(tmp2211)

											tmp2213 := PrimTail(tmp2212)

											tmp2214 := PrimTail(tmp2213)

											tmp2215 := PrimTail(tmp2214)

											tmp2216 := PrimEqual(Nil, tmp2215)

											var ifres2188 Obj

											if True == tmp2216 {
												tmp2206 := PrimHead(V385)

												tmp2207 := PrimTail(tmp2206)

												tmp2208 := PrimTail(tmp2207)

												tmp2209 := PrimIsPair(tmp2208)

												var ifres2190 Obj

												if True == tmp2209 {
													tmp2201 := PrimHead(V385)

													tmp2202 := PrimTail(tmp2201)

													tmp2203 := PrimTail(tmp2202)

													tmp2204 := PrimTail(tmp2203)

													tmp2205 := PrimEqual(Nil, tmp2204)

													var ifres2192 Obj

													if True == tmp2205 {
														tmp2199 := PrimTail(V385)

														tmp2200 := PrimIsPair(tmp2199)

														var ifres2194 Obj

														if True == tmp2200 {
															tmp2196 := PrimTail(V385)

															tmp2197 := PrimTail(tmp2196)

															tmp2198 := PrimEqual(Nil, tmp2197)

															var ifres2195 Obj

															if True == tmp2198 {
																ifres2195 = True

															} else {
																ifres2195 = False

															}

															ifres2194 = ifres2195

														} else {
															ifres2194 = False

														}

														var ifres2193 Obj

														if True == ifres2194 {
															ifres2193 = True

														} else {
															ifres2193 = False

														}

														ifres2192 = ifres2193

													} else {
														ifres2192 = False

													}

													var ifres2191 Obj

													if True == ifres2192 {
														ifres2191 = True

													} else {
														ifres2191 = False

													}

													ifres2190 = ifres2191

												} else {
													ifres2190 = False

												}

												var ifres2189 Obj

												if True == ifres2190 {
													ifres2189 = True

												} else {
													ifres2189 = False

												}

												ifres2188 = ifres2189

											} else {
												ifres2188 = False

											}

											var ifres2187 Obj

											if True == ifres2188 {
												ifres2187 = True

											} else {
												ifres2187 = False

											}

											ifres2186 = ifres2187

										} else {
											ifres2186 = False

										}

										var ifres2185 Obj

										if True == ifres2186 {
											ifres2185 = True

										} else {
											ifres2185 = False

										}

										ifres2184 = ifres2185

									} else {
										ifres2184 = False

									}

									var ifres2183 Obj

									if True == ifres2184 {
										ifres2183 = True

									} else {
										ifres2183 = False

									}

									ifres2182 = ifres2183

								} else {
									ifres2182 = False

								}

								var ifres2181 Obj

								if True == ifres2182 {
									ifres2181 = True

								} else {
									ifres2181 = False

								}

								ifres2180 = ifres2181

							} else {
								ifres2180 = False

							}

							var ifres2179 Obj

							if True == ifres2180 {
								ifres2179 = True

							} else {
								ifres2179 = False

							}

							ifres2178 = ifres2179

						} else {
							ifres2178 = False

						}

						var ifres2177 Obj

						if True == ifres2178 {
							ifres2177 = True

						} else {
							ifres2177 = False

						}

						ifres2176 = ifres2177

					} else {
						ifres2176 = False

					}

					var ifres2175 Obj

					if True == ifres2176 {
						ifres2175 = True

					} else {
						ifres2175 = False

					}

					ifres2174 = ifres2175

				} else {
					ifres2174 = False

				}

				var ifres2173 Obj

				if True == ifres2174 {
					ifres2173 = True

				} else {
					ifres2173 = False

				}

				ifres2172 = ifres2173

			} else {
				ifres2172 = False

			}

			if True == ifres2172 {
				tmp2132 := PrimTail(V385)

				tmp2133 := PrimCons(symtuple_2, tmp2132)

				tmp2134 := Call(__e, PrimNS2Value(symshen_4add__test), tmp2133)

				_ = tmp2134

				tmp2135 := MakeNative(func(__e *ControlFlow) {
					Abstraction := __e.Get(1)
					_ = Abstraction
					tmp2136 := MakeNative(func(__e *ControlFlow) {
						Application := __e.Get(1)
						_ = Application
						__e.TailApply(PrimNS2Value(symshen_4reduce__help), Application)
						return
					}, 1)

					tmp2137 := PrimTail(V385)

					tmp2138 := PrimCons(symfst, tmp2137)

					tmp2139 := PrimCons(tmp2138, Nil)

					tmp2140 := PrimCons(Abstraction, tmp2139)

					tmp2141 := PrimTail(V385)

					tmp2142 := PrimCons(symsnd, tmp2141)

					tmp2143 := PrimCons(tmp2142, Nil)

					tmp2144 := PrimCons(tmp2140, tmp2143)

					__e.TailApply(tmp2136, tmp2144)
					return

				}, 1)

				tmp2145 := PrimHead(V385)

				tmp2146 := PrimTail(tmp2145)

				tmp2147 := PrimHead(tmp2146)

				tmp2148 := PrimTail(tmp2147)

				tmp2149 := PrimHead(tmp2148)

				tmp2150 := PrimHead(V385)

				tmp2151 := PrimTail(tmp2150)

				tmp2152 := PrimHead(tmp2151)

				tmp2153 := PrimTail(tmp2152)

				tmp2154 := PrimTail(tmp2153)

				tmp2155 := PrimHead(tmp2154)

				tmp2156 := PrimTail(V385)

				tmp2157 := PrimHead(tmp2156)

				tmp2158 := PrimHead(V385)

				tmp2159 := PrimTail(tmp2158)

				tmp2160 := PrimHead(tmp2159)

				tmp2161 := PrimHead(V385)

				tmp2162 := PrimTail(tmp2161)

				tmp2163 := PrimTail(tmp2162)

				tmp2164 := PrimHead(tmp2163)

				tmp2165 := Call(__e, PrimNS2Value(symshen_4ebr), tmp2157, tmp2160, tmp2164)

				tmp2166 := PrimCons(tmp2165, Nil)

				tmp2167 := PrimCons(tmp2155, tmp2166)

				tmp2168 := PrimCons(sym_c_4, tmp2167)

				tmp2169 := PrimCons(tmp2168, Nil)

				tmp2170 := PrimCons(tmp2149, tmp2169)

				tmp2171 := PrimCons(sym_c_4, tmp2170)

				__e.TailApply(tmp2135, tmp2171)
				return

			} else {
				tmp2131 := PrimIsPair(V385)

				var ifres2058 Obj

				if True == tmp2131 {
					tmp2129 := PrimHead(V385)

					tmp2130 := PrimIsPair(tmp2129)

					var ifres2060 Obj

					if True == tmp2130 {
						tmp2126 := PrimHead(V385)

						tmp2127 := PrimHead(tmp2126)

						tmp2128 := PrimEqual(sym_c_4, tmp2127)

						var ifres2062 Obj

						if True == tmp2128 {
							tmp2123 := PrimHead(V385)

							tmp2124 := PrimTail(tmp2123)

							tmp2125 := PrimIsPair(tmp2124)

							var ifres2064 Obj

							if True == tmp2125 {
								tmp2119 := PrimHead(V385)

								tmp2120 := PrimTail(tmp2119)

								tmp2121 := PrimHead(tmp2120)

								tmp2122 := PrimIsPair(tmp2121)

								var ifres2066 Obj

								if True == tmp2122 {
									tmp2114 := PrimHead(V385)

									tmp2115 := PrimTail(tmp2114)

									tmp2116 := PrimHead(tmp2115)

									tmp2117 := PrimHead(tmp2116)

									tmp2118 := PrimEqual(sym_8v, tmp2117)

									var ifres2068 Obj

									if True == tmp2118 {
										tmp2109 := PrimHead(V385)

										tmp2110 := PrimTail(tmp2109)

										tmp2111 := PrimHead(tmp2110)

										tmp2112 := PrimTail(tmp2111)

										tmp2113 := PrimIsPair(tmp2112)

										var ifres2070 Obj

										if True == tmp2113 {
											tmp2103 := PrimHead(V385)

											tmp2104 := PrimTail(tmp2103)

											tmp2105 := PrimHead(tmp2104)

											tmp2106 := PrimTail(tmp2105)

											tmp2107 := PrimTail(tmp2106)

											tmp2108 := PrimIsPair(tmp2107)

											var ifres2072 Obj

											if True == tmp2108 {
												tmp2096 := PrimHead(V385)

												tmp2097 := PrimTail(tmp2096)

												tmp2098 := PrimHead(tmp2097)

												tmp2099 := PrimTail(tmp2098)

												tmp2100 := PrimTail(tmp2099)

												tmp2101 := PrimTail(tmp2100)

												tmp2102 := PrimEqual(Nil, tmp2101)

												var ifres2074 Obj

												if True == tmp2102 {
													tmp2092 := PrimHead(V385)

													tmp2093 := PrimTail(tmp2092)

													tmp2094 := PrimTail(tmp2093)

													tmp2095 := PrimIsPair(tmp2094)

													var ifres2076 Obj

													if True == tmp2095 {
														tmp2087 := PrimHead(V385)

														tmp2088 := PrimTail(tmp2087)

														tmp2089 := PrimTail(tmp2088)

														tmp2090 := PrimTail(tmp2089)

														tmp2091 := PrimEqual(Nil, tmp2090)

														var ifres2078 Obj

														if True == tmp2091 {
															tmp2085 := PrimTail(V385)

															tmp2086 := PrimIsPair(tmp2085)

															var ifres2080 Obj

															if True == tmp2086 {
																tmp2082 := PrimTail(V385)

																tmp2083 := PrimTail(tmp2082)

																tmp2084 := PrimEqual(Nil, tmp2083)

																var ifres2081 Obj

																if True == tmp2084 {
																	ifres2081 = True

																} else {
																	ifres2081 = False

																}

																ifres2080 = ifres2081

															} else {
																ifres2080 = False

															}

															var ifres2079 Obj

															if True == ifres2080 {
																ifres2079 = True

															} else {
																ifres2079 = False

															}

															ifres2078 = ifres2079

														} else {
															ifres2078 = False

														}

														var ifres2077 Obj

														if True == ifres2078 {
															ifres2077 = True

														} else {
															ifres2077 = False

														}

														ifres2076 = ifres2077

													} else {
														ifres2076 = False

													}

													var ifres2075 Obj

													if True == ifres2076 {
														ifres2075 = True

													} else {
														ifres2075 = False

													}

													ifres2074 = ifres2075

												} else {
													ifres2074 = False

												}

												var ifres2073 Obj

												if True == ifres2074 {
													ifres2073 = True

												} else {
													ifres2073 = False

												}

												ifres2072 = ifres2073

											} else {
												ifres2072 = False

											}

											var ifres2071 Obj

											if True == ifres2072 {
												ifres2071 = True

											} else {
												ifres2071 = False

											}

											ifres2070 = ifres2071

										} else {
											ifres2070 = False

										}

										var ifres2069 Obj

										if True == ifres2070 {
											ifres2069 = True

										} else {
											ifres2069 = False

										}

										ifres2068 = ifres2069

									} else {
										ifres2068 = False

									}

									var ifres2067 Obj

									if True == ifres2068 {
										ifres2067 = True

									} else {
										ifres2067 = False

									}

									ifres2066 = ifres2067

								} else {
									ifres2066 = False

								}

								var ifres2065 Obj

								if True == ifres2066 {
									ifres2065 = True

								} else {
									ifres2065 = False

								}

								ifres2064 = ifres2065

							} else {
								ifres2064 = False

							}

							var ifres2063 Obj

							if True == ifres2064 {
								ifres2063 = True

							} else {
								ifres2063 = False

							}

							ifres2062 = ifres2063

						} else {
							ifres2062 = False

						}

						var ifres2061 Obj

						if True == ifres2062 {
							ifres2061 = True

						} else {
							ifres2061 = False

						}

						ifres2060 = ifres2061

					} else {
						ifres2060 = False

					}

					var ifres2059 Obj

					if True == ifres2060 {
						ifres2059 = True

					} else {
						ifres2059 = False

					}

					ifres2058 = ifres2059

				} else {
					ifres2058 = False

				}

				if True == ifres2058 {
					tmp2018 := PrimTail(V385)

					tmp2019 := PrimCons(symshen_4_7vector_2, tmp2018)

					tmp2020 := Call(__e, PrimNS2Value(symshen_4add__test), tmp2019)

					_ = tmp2020

					tmp2021 := MakeNative(func(__e *ControlFlow) {
						Abstraction := __e.Get(1)
						_ = Abstraction
						tmp2022 := MakeNative(func(__e *ControlFlow) {
							Application := __e.Get(1)
							_ = Application
							__e.TailApply(PrimNS2Value(symshen_4reduce__help), Application)
							return
						}, 1)

						tmp2023 := PrimTail(V385)

						tmp2024 := PrimCons(symhdv, tmp2023)

						tmp2025 := PrimCons(tmp2024, Nil)

						tmp2026 := PrimCons(Abstraction, tmp2025)

						tmp2027 := PrimTail(V385)

						tmp2028 := PrimCons(symtlv, tmp2027)

						tmp2029 := PrimCons(tmp2028, Nil)

						tmp2030 := PrimCons(tmp2026, tmp2029)

						__e.TailApply(tmp2022, tmp2030)
						return

					}, 1)

					tmp2031 := PrimHead(V385)

					tmp2032 := PrimTail(tmp2031)

					tmp2033 := PrimHead(tmp2032)

					tmp2034 := PrimTail(tmp2033)

					tmp2035 := PrimHead(tmp2034)

					tmp2036 := PrimHead(V385)

					tmp2037 := PrimTail(tmp2036)

					tmp2038 := PrimHead(tmp2037)

					tmp2039 := PrimTail(tmp2038)

					tmp2040 := PrimTail(tmp2039)

					tmp2041 := PrimHead(tmp2040)

					tmp2042 := PrimTail(V385)

					tmp2043 := PrimHead(tmp2042)

					tmp2044 := PrimHead(V385)

					tmp2045 := PrimTail(tmp2044)

					tmp2046 := PrimHead(tmp2045)

					tmp2047 := PrimHead(V385)

					tmp2048 := PrimTail(tmp2047)

					tmp2049 := PrimTail(tmp2048)

					tmp2050 := PrimHead(tmp2049)

					tmp2051 := Call(__e, PrimNS2Value(symshen_4ebr), tmp2043, tmp2046, tmp2050)

					tmp2052 := PrimCons(tmp2051, Nil)

					tmp2053 := PrimCons(tmp2041, tmp2052)

					tmp2054 := PrimCons(sym_c_4, tmp2053)

					tmp2055 := PrimCons(tmp2054, Nil)

					tmp2056 := PrimCons(tmp2035, tmp2055)

					tmp2057 := PrimCons(sym_c_4, tmp2056)

					__e.TailApply(tmp2021, tmp2057)
					return

				} else {
					tmp2017 := PrimIsPair(V385)

					var ifres1944 Obj

					if True == tmp2017 {
						tmp2015 := PrimHead(V385)

						tmp2016 := PrimIsPair(tmp2015)

						var ifres1946 Obj

						if True == tmp2016 {
							tmp2012 := PrimHead(V385)

							tmp2013 := PrimHead(tmp2012)

							tmp2014 := PrimEqual(sym_c_4, tmp2013)

							var ifres1948 Obj

							if True == tmp2014 {
								tmp2009 := PrimHead(V385)

								tmp2010 := PrimTail(tmp2009)

								tmp2011 := PrimIsPair(tmp2010)

								var ifres1950 Obj

								if True == tmp2011 {
									tmp2005 := PrimHead(V385)

									tmp2006 := PrimTail(tmp2005)

									tmp2007 := PrimHead(tmp2006)

									tmp2008 := PrimIsPair(tmp2007)

									var ifres1952 Obj

									if True == tmp2008 {
										tmp2000 := PrimHead(V385)

										tmp2001 := PrimTail(tmp2000)

										tmp2002 := PrimHead(tmp2001)

										tmp2003 := PrimHead(tmp2002)

										tmp2004 := PrimEqual(sym_8s, tmp2003)

										var ifres1954 Obj

										if True == tmp2004 {
											tmp1995 := PrimHead(V385)

											tmp1996 := PrimTail(tmp1995)

											tmp1997 := PrimHead(tmp1996)

											tmp1998 := PrimTail(tmp1997)

											tmp1999 := PrimIsPair(tmp1998)

											var ifres1956 Obj

											if True == tmp1999 {
												tmp1989 := PrimHead(V385)

												tmp1990 := PrimTail(tmp1989)

												tmp1991 := PrimHead(tmp1990)

												tmp1992 := PrimTail(tmp1991)

												tmp1993 := PrimTail(tmp1992)

												tmp1994 := PrimIsPair(tmp1993)

												var ifres1958 Obj

												if True == tmp1994 {
													tmp1982 := PrimHead(V385)

													tmp1983 := PrimTail(tmp1982)

													tmp1984 := PrimHead(tmp1983)

													tmp1985 := PrimTail(tmp1984)

													tmp1986 := PrimTail(tmp1985)

													tmp1987 := PrimTail(tmp1986)

													tmp1988 := PrimEqual(Nil, tmp1987)

													var ifres1960 Obj

													if True == tmp1988 {
														tmp1978 := PrimHead(V385)

														tmp1979 := PrimTail(tmp1978)

														tmp1980 := PrimTail(tmp1979)

														tmp1981 := PrimIsPair(tmp1980)

														var ifres1962 Obj

														if True == tmp1981 {
															tmp1973 := PrimHead(V385)

															tmp1974 := PrimTail(tmp1973)

															tmp1975 := PrimTail(tmp1974)

															tmp1976 := PrimTail(tmp1975)

															tmp1977 := PrimEqual(Nil, tmp1976)

															var ifres1964 Obj

															if True == tmp1977 {
																tmp1971 := PrimTail(V385)

																tmp1972 := PrimIsPair(tmp1971)

																var ifres1966 Obj

																if True == tmp1972 {
																	tmp1968 := PrimTail(V385)

																	tmp1969 := PrimTail(tmp1968)

																	tmp1970 := PrimEqual(Nil, tmp1969)

																	var ifres1967 Obj

																	if True == tmp1970 {
																		ifres1967 = True

																	} else {
																		ifres1967 = False

																	}

																	ifres1966 = ifres1967

																} else {
																	ifres1966 = False

																}

																var ifres1965 Obj

																if True == ifres1966 {
																	ifres1965 = True

																} else {
																	ifres1965 = False

																}

																ifres1964 = ifres1965

															} else {
																ifres1964 = False

															}

															var ifres1963 Obj

															if True == ifres1964 {
																ifres1963 = True

															} else {
																ifres1963 = False

															}

															ifres1962 = ifres1963

														} else {
															ifres1962 = False

														}

														var ifres1961 Obj

														if True == ifres1962 {
															ifres1961 = True

														} else {
															ifres1961 = False

														}

														ifres1960 = ifres1961

													} else {
														ifres1960 = False

													}

													var ifres1959 Obj

													if True == ifres1960 {
														ifres1959 = True

													} else {
														ifres1959 = False

													}

													ifres1958 = ifres1959

												} else {
													ifres1958 = False

												}

												var ifres1957 Obj

												if True == ifres1958 {
													ifres1957 = True

												} else {
													ifres1957 = False

												}

												ifres1956 = ifres1957

											} else {
												ifres1956 = False

											}

											var ifres1955 Obj

											if True == ifres1956 {
												ifres1955 = True

											} else {
												ifres1955 = False

											}

											ifres1954 = ifres1955

										} else {
											ifres1954 = False

										}

										var ifres1953 Obj

										if True == ifres1954 {
											ifres1953 = True

										} else {
											ifres1953 = False

										}

										ifres1952 = ifres1953

									} else {
										ifres1952 = False

									}

									var ifres1951 Obj

									if True == ifres1952 {
										ifres1951 = True

									} else {
										ifres1951 = False

									}

									ifres1950 = ifres1951

								} else {
									ifres1950 = False

								}

								var ifres1949 Obj

								if True == ifres1950 {
									ifres1949 = True

								} else {
									ifres1949 = False

								}

								ifres1948 = ifres1949

							} else {
								ifres1948 = False

							}

							var ifres1947 Obj

							if True == ifres1948 {
								ifres1947 = True

							} else {
								ifres1947 = False

							}

							ifres1946 = ifres1947

						} else {
							ifres1946 = False

						}

						var ifres1945 Obj

						if True == ifres1946 {
							ifres1945 = True

						} else {
							ifres1945 = False

						}

						ifres1944 = ifres1945

					} else {
						ifres1944 = False

					}

					if True == ifres1944 {
						tmp1901 := PrimTail(V385)

						tmp1902 := PrimCons(symshen_4_7string_2, tmp1901)

						tmp1903 := Call(__e, PrimNS2Value(symshen_4add__test), tmp1902)

						_ = tmp1903

						tmp1904 := MakeNative(func(__e *ControlFlow) {
							Abstraction := __e.Get(1)
							_ = Abstraction
							tmp1905 := MakeNative(func(__e *ControlFlow) {
								Application := __e.Get(1)
								_ = Application
								__e.TailApply(PrimNS2Value(symshen_4reduce__help), Application)
								return
							}, 1)

							tmp1906 := PrimTail(V385)

							tmp1907 := PrimHead(tmp1906)

							tmp1908 := PrimCons(MakeNumber(0), Nil)

							tmp1909 := PrimCons(tmp1907, tmp1908)

							tmp1910 := PrimCons(sympos, tmp1909)

							tmp1911 := PrimCons(tmp1910, Nil)

							tmp1912 := PrimCons(Abstraction, tmp1911)

							tmp1913 := PrimTail(V385)

							tmp1914 := PrimCons(symtlstr, tmp1913)

							tmp1915 := PrimCons(tmp1914, Nil)

							tmp1916 := PrimCons(tmp1912, tmp1915)

							__e.TailApply(tmp1905, tmp1916)
							return

						}, 1)

						tmp1917 := PrimHead(V385)

						tmp1918 := PrimTail(tmp1917)

						tmp1919 := PrimHead(tmp1918)

						tmp1920 := PrimTail(tmp1919)

						tmp1921 := PrimHead(tmp1920)

						tmp1922 := PrimHead(V385)

						tmp1923 := PrimTail(tmp1922)

						tmp1924 := PrimHead(tmp1923)

						tmp1925 := PrimTail(tmp1924)

						tmp1926 := PrimTail(tmp1925)

						tmp1927 := PrimHead(tmp1926)

						tmp1928 := PrimTail(V385)

						tmp1929 := PrimHead(tmp1928)

						tmp1930 := PrimHead(V385)

						tmp1931 := PrimTail(tmp1930)

						tmp1932 := PrimHead(tmp1931)

						tmp1933 := PrimHead(V385)

						tmp1934 := PrimTail(tmp1933)

						tmp1935 := PrimTail(tmp1934)

						tmp1936 := PrimHead(tmp1935)

						tmp1937 := Call(__e, PrimNS2Value(symshen_4ebr), tmp1929, tmp1932, tmp1936)

						tmp1938 := PrimCons(tmp1937, Nil)

						tmp1939 := PrimCons(tmp1927, tmp1938)

						tmp1940 := PrimCons(sym_c_4, tmp1939)

						tmp1941 := PrimCons(tmp1940, Nil)

						tmp1942 := PrimCons(tmp1921, tmp1941)

						tmp1943 := PrimCons(sym_c_4, tmp1942)

						__e.TailApply(tmp1904, tmp1943)
						return

					} else {
						tmp1900 := PrimIsPair(V385)

						var ifres1858 Obj

						if True == tmp1900 {
							tmp1898 := PrimHead(V385)

							tmp1899 := PrimIsPair(tmp1898)

							var ifres1860 Obj

							if True == tmp1899 {
								tmp1895 := PrimHead(V385)

								tmp1896 := PrimHead(tmp1895)

								tmp1897 := PrimEqual(sym_c_4, tmp1896)

								var ifres1862 Obj

								if True == tmp1897 {
									tmp1892 := PrimHead(V385)

									tmp1893 := PrimTail(tmp1892)

									tmp1894 := PrimIsPair(tmp1893)

									var ifres1864 Obj

									if True == tmp1894 {
										tmp1888 := PrimHead(V385)

										tmp1889 := PrimTail(tmp1888)

										tmp1890 := PrimHead(tmp1889)

										tmp1891 := PrimIsPair(tmp1890)

										var ifres1866 Obj

										if True == tmp1891 {
											tmp1884 := PrimHead(V385)

											tmp1885 := PrimTail(tmp1884)

											tmp1886 := PrimTail(tmp1885)

											tmp1887 := PrimIsPair(tmp1886)

											var ifres1868 Obj

											if True == tmp1887 {
												tmp1879 := PrimHead(V385)

												tmp1880 := PrimTail(tmp1879)

												tmp1881 := PrimTail(tmp1880)

												tmp1882 := PrimTail(tmp1881)

												tmp1883 := PrimEqual(Nil, tmp1882)

												var ifres1870 Obj

												if True == tmp1883 {
													tmp1877 := PrimTail(V385)

													tmp1878 := PrimIsPair(tmp1877)

													var ifres1872 Obj

													if True == tmp1878 {
														tmp1874 := PrimTail(V385)

														tmp1875 := PrimTail(tmp1874)

														tmp1876 := PrimEqual(Nil, tmp1875)

														var ifres1873 Obj

														if True == tmp1876 {
															ifres1873 = True

														} else {
															ifres1873 = False

														}

														ifres1872 = ifres1873

													} else {
														ifres1872 = False

													}

													var ifres1871 Obj

													if True == ifres1872 {
														ifres1871 = True

													} else {
														ifres1871 = False

													}

													ifres1870 = ifres1871

												} else {
													ifres1870 = False

												}

												var ifres1869 Obj

												if True == ifres1870 {
													ifres1869 = True

												} else {
													ifres1869 = False

												}

												ifres1868 = ifres1869

											} else {
												ifres1868 = False

											}

											var ifres1867 Obj

											if True == ifres1868 {
												ifres1867 = True

											} else {
												ifres1867 = False

											}

											ifres1866 = ifres1867

										} else {
											ifres1866 = False

										}

										var ifres1865 Obj

										if True == ifres1866 {
											ifres1865 = True

										} else {
											ifres1865 = False

										}

										ifres1864 = ifres1865

									} else {
										ifres1864 = False

									}

									var ifres1863 Obj

									if True == ifres1864 {
										ifres1863 = True

									} else {
										ifres1863 = False

									}

									ifres1862 = ifres1863

								} else {
									ifres1862 = False

								}

								var ifres1861 Obj

								if True == ifres1862 {
									ifres1861 = True

								} else {
									ifres1861 = False

								}

								ifres1860 = ifres1861

							} else {
								ifres1860 = False

							}

							var ifres1859 Obj

							if True == ifres1860 {
								ifres1859 = True

							} else {
								ifres1859 = False

							}

							ifres1858 = ifres1859

						} else {
							ifres1858 = False

						}

						if True == ifres1858 {
							__e.TailApply(PrimNS2Value(symshen_4custom_1pattern_1reducer), V385)
							return
						} else {
							tmp1857 := PrimIsPair(V385)

							var ifres1814 Obj

							if True == tmp1857 {
								tmp1855 := PrimHead(V385)

								tmp1856 := PrimIsPair(tmp1855)

								var ifres1816 Obj

								if True == tmp1856 {
									tmp1852 := PrimHead(V385)

									tmp1853 := PrimHead(tmp1852)

									tmp1854 := PrimEqual(sym_c_4, tmp1853)

									var ifres1818 Obj

									if True == tmp1854 {
										tmp1849 := PrimHead(V385)

										tmp1850 := PrimTail(tmp1849)

										tmp1851 := PrimIsPair(tmp1850)

										var ifres1820 Obj

										if True == tmp1851 {
											tmp1845 := PrimHead(V385)

											tmp1846 := PrimTail(tmp1845)

											tmp1847 := PrimTail(tmp1846)

											tmp1848 := PrimIsPair(tmp1847)

											var ifres1822 Obj

											if True == tmp1848 {
												tmp1840 := PrimHead(V385)

												tmp1841 := PrimTail(tmp1840)

												tmp1842 := PrimTail(tmp1841)

												tmp1843 := PrimTail(tmp1842)

												tmp1844 := PrimEqual(Nil, tmp1843)

												var ifres1824 Obj

												if True == tmp1844 {
													tmp1838 := PrimTail(V385)

													tmp1839 := PrimIsPair(tmp1838)

													var ifres1826 Obj

													if True == tmp1839 {
														tmp1835 := PrimTail(V385)

														tmp1836 := PrimTail(tmp1835)

														tmp1837 := PrimEqual(Nil, tmp1836)

														var ifres1828 Obj

														if True == tmp1837 {
															tmp1830 := PrimHead(V385)

															tmp1831 := PrimTail(tmp1830)

															tmp1832 := PrimHead(tmp1831)

															tmp1833 := PrimIsVariable(tmp1832)

															tmp1834 := PrimNot(tmp1833)

															var ifres1829 Obj

															if True == tmp1834 {
																ifres1829 = True

															} else {
																ifres1829 = False

															}

															ifres1828 = ifres1829

														} else {
															ifres1828 = False

														}

														var ifres1827 Obj

														if True == ifres1828 {
															ifres1827 = True

														} else {
															ifres1827 = False

														}

														ifres1826 = ifres1827

													} else {
														ifres1826 = False

													}

													var ifres1825 Obj

													if True == ifres1826 {
														ifres1825 = True

													} else {
														ifres1825 = False

													}

													ifres1824 = ifres1825

												} else {
													ifres1824 = False

												}

												var ifres1823 Obj

												if True == ifres1824 {
													ifres1823 = True

												} else {
													ifres1823 = False

												}

												ifres1822 = ifres1823

											} else {
												ifres1822 = False

											}

											var ifres1821 Obj

											if True == ifres1822 {
												ifres1821 = True

											} else {
												ifres1821 = False

											}

											ifres1820 = ifres1821

										} else {
											ifres1820 = False

										}

										var ifres1819 Obj

										if True == ifres1820 {
											ifres1819 = True

										} else {
											ifres1819 = False

										}

										ifres1818 = ifres1819

									} else {
										ifres1818 = False

									}

									var ifres1817 Obj

									if True == ifres1818 {
										ifres1817 = True

									} else {
										ifres1817 = False

									}

									ifres1816 = ifres1817

								} else {
									ifres1816 = False

								}

								var ifres1815 Obj

								if True == ifres1816 {
									ifres1815 = True

								} else {
									ifres1815 = False

								}

								ifres1814 = ifres1815

							} else {
								ifres1814 = False

							}

							if True == ifres1814 {
								tmp1803 := PrimHead(V385)

								tmp1804 := PrimTail(tmp1803)

								tmp1805 := PrimHead(tmp1804)

								tmp1806 := PrimTail(V385)

								tmp1807 := PrimCons(tmp1805, tmp1806)

								tmp1808 := PrimCons(sym_a, tmp1807)

								tmp1809 := Call(__e, PrimNS2Value(symshen_4add__test), tmp1808)

								_ = tmp1809

								tmp1810 := PrimHead(V385)

								tmp1811 := PrimTail(tmp1810)

								tmp1812 := PrimTail(tmp1811)

								tmp1813 := PrimHead(tmp1812)

								__e.TailApply(PrimNS2Value(symshen_4reduce__help), tmp1813)
								return

							} else {
								tmp1802 := PrimIsPair(V385)

								var ifres1766 Obj

								if True == tmp1802 {
									tmp1800 := PrimHead(V385)

									tmp1801 := PrimIsPair(tmp1800)

									var ifres1768 Obj

									if True == tmp1801 {
										tmp1797 := PrimHead(V385)

										tmp1798 := PrimHead(tmp1797)

										tmp1799 := PrimEqual(sym_c_4, tmp1798)

										var ifres1770 Obj

										if True == tmp1799 {
											tmp1794 := PrimHead(V385)

											tmp1795 := PrimTail(tmp1794)

											tmp1796 := PrimIsPair(tmp1795)

											var ifres1772 Obj

											if True == tmp1796 {
												tmp1790 := PrimHead(V385)

												tmp1791 := PrimTail(tmp1790)

												tmp1792 := PrimTail(tmp1791)

												tmp1793 := PrimIsPair(tmp1792)

												var ifres1774 Obj

												if True == tmp1793 {
													tmp1785 := PrimHead(V385)

													tmp1786 := PrimTail(tmp1785)

													tmp1787 := PrimTail(tmp1786)

													tmp1788 := PrimTail(tmp1787)

													tmp1789 := PrimEqual(Nil, tmp1788)

													var ifres1776 Obj

													if True == tmp1789 {
														tmp1783 := PrimTail(V385)

														tmp1784 := PrimIsPair(tmp1783)

														var ifres1778 Obj

														if True == tmp1784 {
															tmp1780 := PrimTail(V385)

															tmp1781 := PrimTail(tmp1780)

															tmp1782 := PrimEqual(Nil, tmp1781)

															var ifres1779 Obj

															if True == tmp1782 {
																ifres1779 = True

															} else {
																ifres1779 = False

															}

															ifres1778 = ifres1779

														} else {
															ifres1778 = False

														}

														var ifres1777 Obj

														if True == ifres1778 {
															ifres1777 = True

														} else {
															ifres1777 = False

														}

														ifres1776 = ifres1777

													} else {
														ifres1776 = False

													}

													var ifres1775 Obj

													if True == ifres1776 {
														ifres1775 = True

													} else {
														ifres1775 = False

													}

													ifres1774 = ifres1775

												} else {
													ifres1774 = False

												}

												var ifres1773 Obj

												if True == ifres1774 {
													ifres1773 = True

												} else {
													ifres1773 = False

												}

												ifres1772 = ifres1773

											} else {
												ifres1772 = False

											}

											var ifres1771 Obj

											if True == ifres1772 {
												ifres1771 = True

											} else {
												ifres1771 = False

											}

											ifres1770 = ifres1771

										} else {
											ifres1770 = False

										}

										var ifres1769 Obj

										if True == ifres1770 {
											ifres1769 = True

										} else {
											ifres1769 = False

										}

										ifres1768 = ifres1769

									} else {
										ifres1768 = False

									}

									var ifres1767 Obj

									if True == ifres1768 {
										ifres1767 = True

									} else {
										ifres1767 = False

									}

									ifres1766 = ifres1767

								} else {
									ifres1766 = False

								}

								if True == ifres1766 {
									tmp1756 := PrimTail(V385)

									tmp1757 := PrimHead(tmp1756)

									tmp1758 := PrimHead(V385)

									tmp1759 := PrimTail(tmp1758)

									tmp1760 := PrimHead(tmp1759)

									tmp1761 := PrimHead(V385)

									tmp1762 := PrimTail(tmp1761)

									tmp1763 := PrimTail(tmp1762)

									tmp1764 := PrimHead(tmp1763)

									tmp1765 := Call(__e, PrimNS2Value(symshen_4ebr), tmp1757, tmp1760, tmp1764)

									__e.TailApply(PrimNS2Value(symshen_4reduce__help), tmp1765)
									return

								} else {
									tmp1755 := PrimIsPair(V385)

									var ifres1736 Obj

									if True == tmp1755 {
										tmp1753 := PrimHead(V385)

										tmp1754 := PrimEqual(symwhere, tmp1753)

										var ifres1738 Obj

										if True == tmp1754 {
											tmp1751 := PrimTail(V385)

											tmp1752 := PrimIsPair(tmp1751)

											var ifres1740 Obj

											if True == tmp1752 {
												tmp1748 := PrimTail(V385)

												tmp1749 := PrimTail(tmp1748)

												tmp1750 := PrimIsPair(tmp1749)

												var ifres1742 Obj

												if True == tmp1750 {
													tmp1744 := PrimTail(V385)

													tmp1745 := PrimTail(tmp1744)

													tmp1746 := PrimTail(tmp1745)

													tmp1747 := PrimEqual(Nil, tmp1746)

													var ifres1743 Obj

													if True == tmp1747 {
														ifres1743 = True

													} else {
														ifres1743 = False

													}

													ifres1742 = ifres1743

												} else {
													ifres1742 = False

												}

												var ifres1741 Obj

												if True == ifres1742 {
													ifres1741 = True

												} else {
													ifres1741 = False

												}

												ifres1740 = ifres1741

											} else {
												ifres1740 = False

											}

											var ifres1739 Obj

											if True == ifres1740 {
												ifres1739 = True

											} else {
												ifres1739 = False

											}

											ifres1738 = ifres1739

										} else {
											ifres1738 = False

										}

										var ifres1737 Obj

										if True == ifres1738 {
											ifres1737 = True

										} else {
											ifres1737 = False

										}

										ifres1736 = ifres1737

									} else {
										ifres1736 = False

									}

									if True == ifres1736 {
										tmp1730 := PrimTail(V385)

										tmp1731 := PrimHead(tmp1730)

										tmp1732 := Call(__e, PrimNS2Value(symshen_4add__test), tmp1731)

										_ = tmp1732

										tmp1733 := PrimTail(V385)

										tmp1734 := PrimTail(tmp1733)

										tmp1735 := PrimHead(tmp1734)

										__e.TailApply(PrimNS2Value(symshen_4reduce__help), tmp1735)
										return

									} else {
										tmp1729 := PrimIsPair(V385)

										var ifres1720 Obj

										if True == tmp1729 {
											tmp1727 := PrimTail(V385)

											tmp1728 := PrimIsPair(tmp1727)

											var ifres1722 Obj

											if True == tmp1728 {
												tmp1724 := PrimTail(V385)

												tmp1725 := PrimTail(tmp1724)

												tmp1726 := PrimEqual(Nil, tmp1725)

												var ifres1723 Obj

												if True == tmp1726 {
													ifres1723 = True

												} else {
													ifres1723 = False

												}

												ifres1722 = ifres1723

											} else {
												ifres1722 = False

											}

											var ifres1721 Obj

											if True == ifres1722 {
												ifres1721 = True

											} else {
												ifres1721 = False

											}

											ifres1720 = ifres1721

										} else {
											ifres1720 = False

										}

										if True == ifres1720 {
											tmp1712 := MakeNative(func(__e *ControlFlow) {
												Z := __e.Get(1)
												_ = Z
												tmp1716 := PrimHead(V385)

												tmp1717 := PrimEqual(tmp1716, Z)

												if True == tmp1717 {
													__e.Return(V385)
													return
												} else {
													tmp1714 := PrimTail(V385)

													tmp1715 := PrimCons(Z, tmp1714)

													__e.TailApply(PrimNS2Value(symshen_4reduce__help), tmp1715)
													return

												}

											}, 1)

											tmp1718 := PrimHead(V385)

											tmp1719 := Call(__e, PrimNS2Value(symshen_4reduce__help), tmp1718)

											__e.TailApply(tmp1712, tmp1719)
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

	tmp2360 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reduce__help, tmp1702)

	_ = tmp2360

	tmp2361 := MakeNative(func(__e *ControlFlow) {
		V387 := __e.Get(1)
		_ = V387
		tmp2363 := PrimEqual(MakeString(""), V387)

		if True == tmp2363 {
			__e.Return(False)
			return
		} else {
			__e.Return(PrimIsString(V387))
			return
		}

	}, 1)

	tmp2364 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_7string_2, tmp2361)

	_ = tmp2364

	tmp2365 := MakeNative(func(__e *ControlFlow) {
		V389 := __e.Get(1)
		_ = V389
		tmp2370 := PrimIsVector(V389)

		if True == tmp2370 {
			tmp2368 := PrimVectorGet(V389, MakeNumber(0))

			tmp2369 := PrimGreatThan(tmp2368, MakeNumber(0))

			if True == tmp2369 {
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

	tmp2371 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_7vector_2, tmp2365)

	_ = tmp2371

	tmp2372 := MakeNative(func(__e *ControlFlow) {
		V402 := __e.Get(1)
		_ = V402
		V403 := __e.Get(2)
		_ = V403
		V404 := __e.Get(3)
		_ = V404
		tmp2450 := PrimEqual(V404, V403)

		if True == tmp2450 {
			__e.Return(V402)
			return
		} else {
			tmp2449 := PrimIsPair(V404)

			var ifres2425 Obj

			if True == tmp2449 {
				tmp2447 := PrimHead(V404)

				tmp2448 := PrimEqual(symlambda, tmp2447)

				var ifres2427 Obj

				if True == tmp2448 {
					tmp2445 := PrimTail(V404)

					tmp2446 := PrimIsPair(tmp2445)

					var ifres2429 Obj

					if True == tmp2446 {
						tmp2442 := PrimTail(V404)

						tmp2443 := PrimTail(tmp2442)

						tmp2444 := PrimIsPair(tmp2443)

						var ifres2431 Obj

						if True == tmp2444 {
							tmp2438 := PrimTail(V404)

							tmp2439 := PrimTail(tmp2438)

							tmp2440 := PrimTail(tmp2439)

							tmp2441 := PrimEqual(Nil, tmp2440)

							var ifres2433 Obj

							if True == tmp2441 {
								tmp2435 := PrimTail(V404)

								tmp2436 := PrimHead(tmp2435)

								tmp2437 := Call(__e, PrimNS2Value(symshen_4clash_2), tmp2436, V403)

								var ifres2434 Obj

								if True == tmp2437 {
									ifres2434 = True

								} else {
									ifres2434 = False

								}

								ifres2433 = ifres2434

							} else {
								ifres2433 = False

							}

							var ifres2432 Obj

							if True == ifres2433 {
								ifres2432 = True

							} else {
								ifres2432 = False

							}

							ifres2431 = ifres2432

						} else {
							ifres2431 = False

						}

						var ifres2430 Obj

						if True == ifres2431 {
							ifres2430 = True

						} else {
							ifres2430 = False

						}

						ifres2429 = ifres2430

					} else {
						ifres2429 = False

					}

					var ifres2428 Obj

					if True == ifres2429 {
						ifres2428 = True

					} else {
						ifres2428 = False

					}

					ifres2427 = ifres2428

				} else {
					ifres2427 = False

				}

				var ifres2426 Obj

				if True == ifres2427 {
					ifres2426 = True

				} else {
					ifres2426 = False

				}

				ifres2425 = ifres2426

			} else {
				ifres2425 = False

			}

			if True == ifres2425 {
				__e.Return(V404)
				return
			} else {
				tmp2424 := PrimIsPair(V404)

				var ifres2393 Obj

				if True == tmp2424 {
					tmp2422 := PrimHead(V404)

					tmp2423 := PrimEqual(symlet, tmp2422)

					var ifres2395 Obj

					if True == tmp2423 {
						tmp2420 := PrimTail(V404)

						tmp2421 := PrimIsPair(tmp2420)

						var ifres2397 Obj

						if True == tmp2421 {
							tmp2417 := PrimTail(V404)

							tmp2418 := PrimTail(tmp2417)

							tmp2419 := PrimIsPair(tmp2418)

							var ifres2399 Obj

							if True == tmp2419 {
								tmp2413 := PrimTail(V404)

								tmp2414 := PrimTail(tmp2413)

								tmp2415 := PrimTail(tmp2414)

								tmp2416 := PrimIsPair(tmp2415)

								var ifres2401 Obj

								if True == tmp2416 {
									tmp2408 := PrimTail(V404)

									tmp2409 := PrimTail(tmp2408)

									tmp2410 := PrimTail(tmp2409)

									tmp2411 := PrimTail(tmp2410)

									tmp2412 := PrimEqual(Nil, tmp2411)

									var ifres2403 Obj

									if True == tmp2412 {
										tmp2405 := PrimTail(V404)

										tmp2406 := PrimHead(tmp2405)

										tmp2407 := Call(__e, PrimNS2Value(symshen_4clash_2), tmp2406, V403)

										var ifres2404 Obj

										if True == tmp2407 {
											ifres2404 = True

										} else {
											ifres2404 = False

										}

										ifres2403 = ifres2404

									} else {
										ifres2403 = False

									}

									var ifres2402 Obj

									if True == ifres2403 {
										ifres2402 = True

									} else {
										ifres2402 = False

									}

									ifres2401 = ifres2402

								} else {
									ifres2401 = False

								}

								var ifres2400 Obj

								if True == ifres2401 {
									ifres2400 = True

								} else {
									ifres2400 = False

								}

								ifres2399 = ifres2400

							} else {
								ifres2399 = False

							}

							var ifres2398 Obj

							if True == ifres2399 {
								ifres2398 = True

							} else {
								ifres2398 = False

							}

							ifres2397 = ifres2398

						} else {
							ifres2397 = False

						}

						var ifres2396 Obj

						if True == ifres2397 {
							ifres2396 = True

						} else {
							ifres2396 = False

						}

						ifres2395 = ifres2396

					} else {
						ifres2395 = False

					}

					var ifres2394 Obj

					if True == ifres2395 {
						ifres2394 = True

					} else {
						ifres2394 = False

					}

					ifres2393 = ifres2394

				} else {
					ifres2393 = False

				}

				if True == ifres2393 {
					tmp2382 := PrimTail(V404)

					tmp2383 := PrimHead(tmp2382)

					tmp2384 := PrimTail(V404)

					tmp2385 := PrimTail(tmp2384)

					tmp2386 := PrimHead(tmp2385)

					tmp2387 := Call(__e, PrimNS2Value(symshen_4ebr), V402, V403, tmp2386)

					tmp2388 := PrimTail(V404)

					tmp2389 := PrimTail(tmp2388)

					tmp2390 := PrimTail(tmp2389)

					tmp2391 := PrimCons(tmp2387, tmp2390)

					tmp2392 := PrimCons(tmp2383, tmp2391)

					__e.Return(PrimCons(symlet, tmp2392))
					return

				} else {
					tmp2381 := PrimIsPair(V404)

					if True == tmp2381 {
						tmp2377 := PrimHead(V404)

						tmp2378 := Call(__e, PrimNS2Value(symshen_4ebr), V402, V403, tmp2377)

						tmp2379 := PrimTail(V404)

						tmp2380 := Call(__e, PrimNS2Value(symshen_4ebr), V402, V403, tmp2379)

						__e.Return(PrimCons(tmp2378, tmp2380))
						return

					} else {
						__e.Return(V404)
						return
					}

				}

			}

		}

	}, 3)

	tmp2451 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ebr, tmp2372)

	_ = tmp2451

	tmp2452 := MakeNative(func(__e *ControlFlow) {
		V416 := __e.Get(1)
		_ = V416
		V417 := __e.Get(2)
		_ = V417
		tmp2462 := PrimEqual(V417, V416)

		if True == tmp2462 {
			__e.Return(True)
			return
		} else {
			tmp2461 := PrimIsPair(V417)

			if True == tmp2461 {
				tmp2459 := PrimHead(V417)

				tmp2460 := Call(__e, PrimNS2Value(symshen_4clash_2), V416, tmp2459)

				if True == tmp2460 {
					__e.Return(True)
					return
				} else {
					tmp2457 := PrimTail(V417)

					tmp2458 := Call(__e, PrimNS2Value(symshen_4clash_2), V416, tmp2457)

					if True == tmp2458 {
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

	tmp2463 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clash_2, tmp2452)

	_ = tmp2463

	tmp2464 := MakeNative(func(__e *ControlFlow) {
		V419 := __e.Get(1)
		_ = V419
		tmp2465 := PrimNS3Value(symshen_4_dteststack_d)

		tmp2466 := PrimCons(V419, tmp2465)

		__e.Return(PrimNS3Set(symshen_4_dteststack_d, tmp2466))
		return

	}, 1)

	tmp2467 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add__test, tmp2464)

	_ = tmp2467

	tmp2468 := MakeNative(func(__e *ControlFlow) {
		V423 := __e.Get(1)
		_ = V423
		V424 := __e.Get(2)
		_ = V424
		V425 := __e.Get(3)
		_ = V425
		tmp2469 := MakeNative(func(__e *ControlFlow) {
			Err := __e.Get(1)
			_ = Err
			tmp2470 := MakeNative(func(__e *ControlFlow) {
				Cases := __e.Get(1)
				_ = Cases
				tmp2471 := MakeNative(func(__e *ControlFlow) {
					EncodeChoices := __e.Get(1)
					_ = EncodeChoices
					__e.TailApply(PrimNS2Value(symshen_4cond_1form), EncodeChoices)
					return
				}, 1)

				tmp2472 := Call(__e, PrimNS2Value(symshen_4encode_1choices), Cases, V423)

				__e.TailApply(tmp2471, tmp2472)
				return

			}, 1)

			tmp2473 := Call(__e, PrimNS2Value(symshen_4case_1form), V425, Err)

			__e.TailApply(tmp2470, tmp2473)
			return

		}, 1)

		tmp2474 := Call(__e, PrimNS2Value(symshen_4err_1condition), V423)

		__e.TailApply(tmp2469, tmp2474)
		return

	}, 3)

	tmp2475 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cond_1expression, tmp2468)

	_ = tmp2475

	tmp2476 := MakeNative(func(__e *ControlFlow) {
		V429 := __e.Get(1)
		_ = V429
		tmp2500 := PrimIsPair(V429)

		var ifres2480 Obj

		if True == tmp2500 {
			tmp2498 := PrimHead(V429)

			tmp2499 := PrimIsPair(tmp2498)

			var ifres2482 Obj

			if True == tmp2499 {
				tmp2495 := PrimHead(V429)

				tmp2496 := PrimHead(tmp2495)

				tmp2497 := PrimEqual(True, tmp2496)

				var ifres2484 Obj

				if True == tmp2497 {
					tmp2492 := PrimHead(V429)

					tmp2493 := PrimTail(tmp2492)

					tmp2494 := PrimIsPair(tmp2493)

					var ifres2486 Obj

					if True == tmp2494 {
						tmp2488 := PrimHead(V429)

						tmp2489 := PrimTail(tmp2488)

						tmp2490 := PrimTail(tmp2489)

						tmp2491 := PrimEqual(Nil, tmp2490)

						var ifres2487 Obj

						if True == tmp2491 {
							ifres2487 = True

						} else {
							ifres2487 = False

						}

						ifres2486 = ifres2487

					} else {
						ifres2486 = False

					}

					var ifres2485 Obj

					if True == ifres2486 {
						ifres2485 = True

					} else {
						ifres2485 = False

					}

					ifres2484 = ifres2485

				} else {
					ifres2484 = False

				}

				var ifres2483 Obj

				if True == ifres2484 {
					ifres2483 = True

				} else {
					ifres2483 = False

				}

				ifres2482 = ifres2483

			} else {
				ifres2482 = False

			}

			var ifres2481 Obj

			if True == ifres2482 {
				ifres2481 = True

			} else {
				ifres2481 = False

			}

			ifres2480 = ifres2481

		} else {
			ifres2480 = False

		}

		if True == ifres2480 {
			tmp2478 := PrimHead(V429)

			tmp2479 := PrimTail(tmp2478)

			__e.Return(PrimHead(tmp2479))
			return

		} else {
			__e.Return(PrimCons(symcond, V429))
			return
		}

	}, 1)

	tmp2501 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cond_1form, tmp2476)

	_ = tmp2501

	tmp2502 := MakeNative(func(__e *ControlFlow) {
		V434 := __e.Get(1)
		_ = V434
		V435 := __e.Get(2)
		_ = V435
		tmp2758 := PrimEqual(Nil, V434)

		if True == tmp2758 {
			__e.Return(Nil)
			return
		} else {
			tmp2757 := PrimIsPair(V434)

			var ifres2705 Obj

			if True == tmp2757 {
				tmp2755 := PrimHead(V434)

				tmp2756 := PrimIsPair(tmp2755)

				var ifres2707 Obj

				if True == tmp2756 {
					tmp2752 := PrimHead(V434)

					tmp2753 := PrimHead(tmp2752)

					tmp2754 := PrimEqual(True, tmp2753)

					var ifres2709 Obj

					if True == tmp2754 {
						tmp2749 := PrimHead(V434)

						tmp2750 := PrimTail(tmp2749)

						tmp2751 := PrimIsPair(tmp2750)

						var ifres2711 Obj

						if True == tmp2751 {
							tmp2745 := PrimHead(V434)

							tmp2746 := PrimTail(tmp2745)

							tmp2747 := PrimHead(tmp2746)

							tmp2748 := PrimIsPair(tmp2747)

							var ifres2713 Obj

							if True == tmp2748 {
								tmp2740 := PrimHead(V434)

								tmp2741 := PrimTail(tmp2740)

								tmp2742 := PrimHead(tmp2741)

								tmp2743 := PrimHead(tmp2742)

								tmp2744 := PrimEqual(symshen_4choicepoint_b, tmp2743)

								var ifres2715 Obj

								if True == tmp2744 {
									tmp2735 := PrimHead(V434)

									tmp2736 := PrimTail(tmp2735)

									tmp2737 := PrimHead(tmp2736)

									tmp2738 := PrimTail(tmp2737)

									tmp2739 := PrimIsPair(tmp2738)

									var ifres2717 Obj

									if True == tmp2739 {
										tmp2729 := PrimHead(V434)

										tmp2730 := PrimTail(tmp2729)

										tmp2731 := PrimHead(tmp2730)

										tmp2732 := PrimTail(tmp2731)

										tmp2733 := PrimTail(tmp2732)

										tmp2734 := PrimEqual(Nil, tmp2733)

										var ifres2719 Obj

										if True == tmp2734 {
											tmp2725 := PrimHead(V434)

											tmp2726 := PrimTail(tmp2725)

											tmp2727 := PrimTail(tmp2726)

											tmp2728 := PrimEqual(Nil, tmp2727)

											var ifres2721 Obj

											if True == tmp2728 {
												tmp2723 := PrimTail(V434)

												tmp2724 := PrimEqual(Nil, tmp2723)

												var ifres2722 Obj

												if True == tmp2724 {
													ifres2722 = True

												} else {
													ifres2722 = False

												}

												ifres2721 = ifres2722

											} else {
												ifres2721 = False

											}

											var ifres2720 Obj

											if True == ifres2721 {
												ifres2720 = True

											} else {
												ifres2720 = False

											}

											ifres2719 = ifres2720

										} else {
											ifres2719 = False

										}

										var ifres2718 Obj

										if True == ifres2719 {
											ifres2718 = True

										} else {
											ifres2718 = False

										}

										ifres2717 = ifres2718

									} else {
										ifres2717 = False

									}

									var ifres2716 Obj

									if True == ifres2717 {
										ifres2716 = True

									} else {
										ifres2716 = False

									}

									ifres2715 = ifres2716

								} else {
									ifres2715 = False

								}

								var ifres2714 Obj

								if True == ifres2715 {
									ifres2714 = True

								} else {
									ifres2714 = False

								}

								ifres2713 = ifres2714

							} else {
								ifres2713 = False

							}

							var ifres2712 Obj

							if True == ifres2713 {
								ifres2712 = True

							} else {
								ifres2712 = False

							}

							ifres2711 = ifres2712

						} else {
							ifres2711 = False

						}

						var ifres2710 Obj

						if True == ifres2711 {
							ifres2710 = True

						} else {
							ifres2710 = False

						}

						ifres2709 = ifres2710

					} else {
						ifres2709 = False

					}

					var ifres2708 Obj

					if True == ifres2709 {
						ifres2708 = True

					} else {
						ifres2708 = False

					}

					ifres2707 = ifres2708

				} else {
					ifres2707 = False

				}

				var ifres2706 Obj

				if True == ifres2707 {
					ifres2706 = True

				} else {
					ifres2706 = False

				}

				ifres2705 = ifres2706

			} else {
				ifres2705 = False

			}

			if True == ifres2705 {
				tmp2680 := PrimHead(V434)

				tmp2681 := PrimTail(tmp2680)

				tmp2682 := PrimHead(tmp2681)

				tmp2683 := PrimTail(tmp2682)

				tmp2684 := PrimHead(tmp2683)

				tmp2685 := PrimCons(symfail, Nil)

				tmp2686 := PrimCons(tmp2685, Nil)

				tmp2687 := PrimCons(symResult, tmp2686)

				tmp2688 := PrimCons(sym_a, tmp2687)

				tmp2694 := PrimNS3Value(symshen_4_dinstalling_1kl_d)

				var ifres2689 Obj

				if True == tmp2694 {
					tmp2692 := PrimCons(V435, Nil)

					tmp2693 := PrimCons(symshen_4sys_1error, tmp2692)

					ifres2689 = tmp2693

				} else {
					tmp2690 := PrimCons(V435, Nil)

					tmp2691 := PrimCons(symshen_4f__error, tmp2690)

					ifres2689 = tmp2691

				}

				tmp2695 := PrimCons(symResult, Nil)

				tmp2696 := PrimCons(ifres2689, tmp2695)

				tmp2697 := PrimCons(tmp2688, tmp2696)

				tmp2698 := PrimCons(symif, tmp2697)

				tmp2699 := PrimCons(tmp2698, Nil)

				tmp2700 := PrimCons(tmp2684, tmp2699)

				tmp2701 := PrimCons(symResult, tmp2700)

				tmp2702 := PrimCons(symlet, tmp2701)

				tmp2703 := PrimCons(tmp2702, Nil)

				tmp2704 := PrimCons(True, tmp2703)

				__e.Return(PrimCons(tmp2704, Nil))
				return

			} else {
				tmp2679 := PrimIsPair(V434)

				var ifres2631 Obj

				if True == tmp2679 {
					tmp2677 := PrimHead(V434)

					tmp2678 := PrimIsPair(tmp2677)

					var ifres2633 Obj

					if True == tmp2678 {
						tmp2674 := PrimHead(V434)

						tmp2675 := PrimHead(tmp2674)

						tmp2676 := PrimEqual(True, tmp2675)

						var ifres2635 Obj

						if True == tmp2676 {
							tmp2671 := PrimHead(V434)

							tmp2672 := PrimTail(tmp2671)

							tmp2673 := PrimIsPair(tmp2672)

							var ifres2637 Obj

							if True == tmp2673 {
								tmp2667 := PrimHead(V434)

								tmp2668 := PrimTail(tmp2667)

								tmp2669 := PrimHead(tmp2668)

								tmp2670 := PrimIsPair(tmp2669)

								var ifres2639 Obj

								if True == tmp2670 {
									tmp2662 := PrimHead(V434)

									tmp2663 := PrimTail(tmp2662)

									tmp2664 := PrimHead(tmp2663)

									tmp2665 := PrimHead(tmp2664)

									tmp2666 := PrimEqual(symshen_4choicepoint_b, tmp2665)

									var ifres2641 Obj

									if True == tmp2666 {
										tmp2657 := PrimHead(V434)

										tmp2658 := PrimTail(tmp2657)

										tmp2659 := PrimHead(tmp2658)

										tmp2660 := PrimTail(tmp2659)

										tmp2661 := PrimIsPair(tmp2660)

										var ifres2643 Obj

										if True == tmp2661 {
											tmp2651 := PrimHead(V434)

											tmp2652 := PrimTail(tmp2651)

											tmp2653 := PrimHead(tmp2652)

											tmp2654 := PrimTail(tmp2653)

											tmp2655 := PrimTail(tmp2654)

											tmp2656 := PrimEqual(Nil, tmp2655)

											var ifres2645 Obj

											if True == tmp2656 {
												tmp2647 := PrimHead(V434)

												tmp2648 := PrimTail(tmp2647)

												tmp2649 := PrimTail(tmp2648)

												tmp2650 := PrimEqual(Nil, tmp2649)

												var ifres2646 Obj

												if True == tmp2650 {
													ifres2646 = True

												} else {
													ifres2646 = False

												}

												ifres2645 = ifres2646

											} else {
												ifres2645 = False

											}

											var ifres2644 Obj

											if True == ifres2645 {
												ifres2644 = True

											} else {
												ifres2644 = False

											}

											ifres2643 = ifres2644

										} else {
											ifres2643 = False

										}

										var ifres2642 Obj

										if True == ifres2643 {
											ifres2642 = True

										} else {
											ifres2642 = False

										}

										ifres2641 = ifres2642

									} else {
										ifres2641 = False

									}

									var ifres2640 Obj

									if True == ifres2641 {
										ifres2640 = True

									} else {
										ifres2640 = False

									}

									ifres2639 = ifres2640

								} else {
									ifres2639 = False

								}

								var ifres2638 Obj

								if True == ifres2639 {
									ifres2638 = True

								} else {
									ifres2638 = False

								}

								ifres2637 = ifres2638

							} else {
								ifres2637 = False

							}

							var ifres2636 Obj

							if True == ifres2637 {
								ifres2636 = True

							} else {
								ifres2636 = False

							}

							ifres2635 = ifres2636

						} else {
							ifres2635 = False

						}

						var ifres2634 Obj

						if True == ifres2635 {
							ifres2634 = True

						} else {
							ifres2634 = False

						}

						ifres2633 = ifres2634

					} else {
						ifres2633 = False

					}

					var ifres2632 Obj

					if True == ifres2633 {
						ifres2632 = True

					} else {
						ifres2632 = False

					}

					ifres2631 = ifres2632

				} else {
					ifres2631 = False

				}

				if True == ifres2631 {
					tmp2609 := PrimHead(V434)

					tmp2610 := PrimTail(tmp2609)

					tmp2611 := PrimHead(tmp2610)

					tmp2612 := PrimTail(tmp2611)

					tmp2613 := PrimHead(tmp2612)

					tmp2614 := PrimCons(symfail, Nil)

					tmp2615 := PrimCons(tmp2614, Nil)

					tmp2616 := PrimCons(symResult, tmp2615)

					tmp2617 := PrimCons(sym_a, tmp2616)

					tmp2618 := PrimTail(V434)

					tmp2619 := Call(__e, PrimNS2Value(symshen_4encode_1choices), tmp2618, V435)

					tmp2620 := Call(__e, PrimNS2Value(symshen_4cond_1form), tmp2619)

					tmp2621 := PrimCons(symResult, Nil)

					tmp2622 := PrimCons(tmp2620, tmp2621)

					tmp2623 := PrimCons(tmp2617, tmp2622)

					tmp2624 := PrimCons(symif, tmp2623)

					tmp2625 := PrimCons(tmp2624, Nil)

					tmp2626 := PrimCons(tmp2613, tmp2625)

					tmp2627 := PrimCons(symResult, tmp2626)

					tmp2628 := PrimCons(symlet, tmp2627)

					tmp2629 := PrimCons(tmp2628, Nil)

					tmp2630 := PrimCons(True, tmp2629)

					__e.Return(PrimCons(tmp2630, Nil))
					return

				} else {
					tmp2608 := PrimIsPair(V434)

					var ifres2565 Obj

					if True == tmp2608 {
						tmp2606 := PrimHead(V434)

						tmp2607 := PrimIsPair(tmp2606)

						var ifres2567 Obj

						if True == tmp2607 {
							tmp2603 := PrimHead(V434)

							tmp2604 := PrimTail(tmp2603)

							tmp2605 := PrimIsPair(tmp2604)

							var ifres2569 Obj

							if True == tmp2605 {
								tmp2599 := PrimHead(V434)

								tmp2600 := PrimTail(tmp2599)

								tmp2601 := PrimHead(tmp2600)

								tmp2602 := PrimIsPair(tmp2601)

								var ifres2571 Obj

								if True == tmp2602 {
									tmp2594 := PrimHead(V434)

									tmp2595 := PrimTail(tmp2594)

									tmp2596 := PrimHead(tmp2595)

									tmp2597 := PrimHead(tmp2596)

									tmp2598 := PrimEqual(symshen_4choicepoint_b, tmp2597)

									var ifres2573 Obj

									if True == tmp2598 {
										tmp2589 := PrimHead(V434)

										tmp2590 := PrimTail(tmp2589)

										tmp2591 := PrimHead(tmp2590)

										tmp2592 := PrimTail(tmp2591)

										tmp2593 := PrimIsPair(tmp2592)

										var ifres2575 Obj

										if True == tmp2593 {
											tmp2583 := PrimHead(V434)

											tmp2584 := PrimTail(tmp2583)

											tmp2585 := PrimHead(tmp2584)

											tmp2586 := PrimTail(tmp2585)

											tmp2587 := PrimTail(tmp2586)

											tmp2588 := PrimEqual(Nil, tmp2587)

											var ifres2577 Obj

											if True == tmp2588 {
												tmp2579 := PrimHead(V434)

												tmp2580 := PrimTail(tmp2579)

												tmp2581 := PrimTail(tmp2580)

												tmp2582 := PrimEqual(Nil, tmp2581)

												var ifres2578 Obj

												if True == tmp2582 {
													ifres2578 = True

												} else {
													ifres2578 = False

												}

												ifres2577 = ifres2578

											} else {
												ifres2577 = False

											}

											var ifres2576 Obj

											if True == ifres2577 {
												ifres2576 = True

											} else {
												ifres2576 = False

											}

											ifres2575 = ifres2576

										} else {
											ifres2575 = False

										}

										var ifres2574 Obj

										if True == ifres2575 {
											ifres2574 = True

										} else {
											ifres2574 = False

										}

										ifres2573 = ifres2574

									} else {
										ifres2573 = False

									}

									var ifres2572 Obj

									if True == ifres2573 {
										ifres2572 = True

									} else {
										ifres2572 = False

									}

									ifres2571 = ifres2572

								} else {
									ifres2571 = False

								}

								var ifres2570 Obj

								if True == ifres2571 {
									ifres2570 = True

								} else {
									ifres2570 = False

								}

								ifres2569 = ifres2570

							} else {
								ifres2569 = False

							}

							var ifres2568 Obj

							if True == ifres2569 {
								ifres2568 = True

							} else {
								ifres2568 = False

							}

							ifres2567 = ifres2568

						} else {
							ifres2567 = False

						}

						var ifres2566 Obj

						if True == ifres2567 {
							ifres2566 = True

						} else {
							ifres2566 = False

						}

						ifres2565 = ifres2566

					} else {
						ifres2565 = False

					}

					if True == ifres2565 {
						tmp2527 := PrimTail(V434)

						tmp2528 := Call(__e, PrimNS2Value(symshen_4encode_1choices), tmp2527, V435)

						tmp2529 := Call(__e, PrimNS2Value(symshen_4cond_1form), tmp2528)

						tmp2530 := PrimCons(tmp2529, Nil)

						tmp2531 := PrimCons(symfreeze, tmp2530)

						tmp2532 := PrimHead(V434)

						tmp2533 := PrimHead(tmp2532)

						tmp2534 := PrimHead(V434)

						tmp2535 := PrimTail(tmp2534)

						tmp2536 := PrimHead(tmp2535)

						tmp2537 := PrimTail(tmp2536)

						tmp2538 := PrimHead(tmp2537)

						tmp2539 := PrimCons(symfail, Nil)

						tmp2540 := PrimCons(tmp2539, Nil)

						tmp2541 := PrimCons(symResult, tmp2540)

						tmp2542 := PrimCons(sym_a, tmp2541)

						tmp2543 := PrimCons(symFreeze, Nil)

						tmp2544 := PrimCons(symthaw, tmp2543)

						tmp2545 := PrimCons(symResult, Nil)

						tmp2546 := PrimCons(tmp2544, tmp2545)

						tmp2547 := PrimCons(tmp2542, tmp2546)

						tmp2548 := PrimCons(symif, tmp2547)

						tmp2549 := PrimCons(tmp2548, Nil)

						tmp2550 := PrimCons(tmp2538, tmp2549)

						tmp2551 := PrimCons(symResult, tmp2550)

						tmp2552 := PrimCons(symlet, tmp2551)

						tmp2553 := PrimCons(symFreeze, Nil)

						tmp2554 := PrimCons(symthaw, tmp2553)

						tmp2555 := PrimCons(tmp2554, Nil)

						tmp2556 := PrimCons(tmp2552, tmp2555)

						tmp2557 := PrimCons(tmp2533, tmp2556)

						tmp2558 := PrimCons(symif, tmp2557)

						tmp2559 := PrimCons(tmp2558, Nil)

						tmp2560 := PrimCons(tmp2531, tmp2559)

						tmp2561 := PrimCons(symFreeze, tmp2560)

						tmp2562 := PrimCons(symlet, tmp2561)

						tmp2563 := PrimCons(tmp2562, Nil)

						tmp2564 := PrimCons(True, tmp2563)

						__e.Return(PrimCons(tmp2564, Nil))
						return

					} else {
						tmp2526 := PrimIsPair(V434)

						var ifres2511 Obj

						if True == tmp2526 {
							tmp2524 := PrimHead(V434)

							tmp2525 := PrimIsPair(tmp2524)

							var ifres2513 Obj

							if True == tmp2525 {
								tmp2521 := PrimHead(V434)

								tmp2522 := PrimTail(tmp2521)

								tmp2523 := PrimIsPair(tmp2522)

								var ifres2515 Obj

								if True == tmp2523 {
									tmp2517 := PrimHead(V434)

									tmp2518 := PrimTail(tmp2517)

									tmp2519 := PrimTail(tmp2518)

									tmp2520 := PrimEqual(Nil, tmp2519)

									var ifres2516 Obj

									if True == tmp2520 {
										ifres2516 = True

									} else {
										ifres2516 = False

									}

									ifres2515 = ifres2516

								} else {
									ifres2515 = False

								}

								var ifres2514 Obj

								if True == ifres2515 {
									ifres2514 = True

								} else {
									ifres2514 = False

								}

								ifres2513 = ifres2514

							} else {
								ifres2513 = False

							}

							var ifres2512 Obj

							if True == ifres2513 {
								ifres2512 = True

							} else {
								ifres2512 = False

							}

							ifres2511 = ifres2512

						} else {
							ifres2511 = False

						}

						if True == ifres2511 {
							tmp2508 := PrimHead(V434)

							tmp2509 := PrimTail(V434)

							tmp2510 := Call(__e, PrimNS2Value(symshen_4encode_1choices), tmp2509, V435)

							__e.Return(PrimCons(tmp2508, tmp2510))
							return

						} else {
							__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4encode_1choices)
							return
						}

					}

				}

			}

		}

	}, 2)

	tmp2759 := Call(__e, PrimNS1Value(symns2_1set), symshen_4encode_1choices, tmp2502)

	_ = tmp2759

	tmp2760 := MakeNative(func(__e *ControlFlow) {
		V442 := __e.Get(1)
		_ = V442
		V443 := __e.Get(2)
		_ = V443
		tmp2945 := PrimEqual(Nil, V442)

		if True == tmp2945 {
			__e.Return(PrimCons(V443, Nil))
			return
		} else {
			tmp2944 := PrimIsPair(V442)

			var ifres2870 Obj

			if True == tmp2944 {
				tmp2942 := PrimHead(V442)

				tmp2943 := PrimIsPair(tmp2942)

				var ifres2872 Obj

				if True == tmp2943 {
					tmp2939 := PrimHead(V442)

					tmp2940 := PrimHead(tmp2939)

					tmp2941 := PrimIsPair(tmp2940)

					var ifres2874 Obj

					if True == tmp2941 {
						tmp2935 := PrimHead(V442)

						tmp2936 := PrimHead(tmp2935)

						tmp2937 := PrimHead(tmp2936)

						tmp2938 := PrimEqual(sym_h, tmp2937)

						var ifres2876 Obj

						if True == tmp2938 {
							tmp2931 := PrimHead(V442)

							tmp2932 := PrimHead(tmp2931)

							tmp2933 := PrimTail(tmp2932)

							tmp2934 := PrimIsPair(tmp2933)

							var ifres2878 Obj

							if True == tmp2934 {
								tmp2926 := PrimHead(V442)

								tmp2927 := PrimHead(tmp2926)

								tmp2928 := PrimTail(tmp2927)

								tmp2929 := PrimHead(tmp2928)

								tmp2930 := PrimEqual(symshen_4tests, tmp2929)

								var ifres2880 Obj

								if True == tmp2930 {
									tmp2921 := PrimHead(V442)

									tmp2922 := PrimHead(tmp2921)

									tmp2923 := PrimTail(tmp2922)

									tmp2924 := PrimTail(tmp2923)

									tmp2925 := PrimEqual(Nil, tmp2924)

									var ifres2882 Obj

									if True == tmp2925 {
										tmp2918 := PrimHead(V442)

										tmp2919 := PrimTail(tmp2918)

										tmp2920 := PrimIsPair(tmp2919)

										var ifres2884 Obj

										if True == tmp2920 {
											tmp2914 := PrimHead(V442)

											tmp2915 := PrimTail(tmp2914)

											tmp2916 := PrimHead(tmp2915)

											tmp2917 := PrimIsPair(tmp2916)

											var ifres2886 Obj

											if True == tmp2917 {
												tmp2909 := PrimHead(V442)

												tmp2910 := PrimTail(tmp2909)

												tmp2911 := PrimHead(tmp2910)

												tmp2912 := PrimHead(tmp2911)

												tmp2913 := PrimEqual(symshen_4choicepoint_b, tmp2912)

												var ifres2888 Obj

												if True == tmp2913 {
													tmp2904 := PrimHead(V442)

													tmp2905 := PrimTail(tmp2904)

													tmp2906 := PrimHead(tmp2905)

													tmp2907 := PrimTail(tmp2906)

													tmp2908 := PrimIsPair(tmp2907)

													var ifres2890 Obj

													if True == tmp2908 {
														tmp2898 := PrimHead(V442)

														tmp2899 := PrimTail(tmp2898)

														tmp2900 := PrimHead(tmp2899)

														tmp2901 := PrimTail(tmp2900)

														tmp2902 := PrimTail(tmp2901)

														tmp2903 := PrimEqual(Nil, tmp2902)

														var ifres2892 Obj

														if True == tmp2903 {
															tmp2894 := PrimHead(V442)

															tmp2895 := PrimTail(tmp2894)

															tmp2896 := PrimTail(tmp2895)

															tmp2897 := PrimEqual(Nil, tmp2896)

															var ifres2893 Obj

															if True == tmp2897 {
																ifres2893 = True

															} else {
																ifres2893 = False

															}

															ifres2892 = ifres2893

														} else {
															ifres2892 = False

														}

														var ifres2891 Obj

														if True == ifres2892 {
															ifres2891 = True

														} else {
															ifres2891 = False

														}

														ifres2890 = ifres2891

													} else {
														ifres2890 = False

													}

													var ifres2889 Obj

													if True == ifres2890 {
														ifres2889 = True

													} else {
														ifres2889 = False

													}

													ifres2888 = ifres2889

												} else {
													ifres2888 = False

												}

												var ifres2887 Obj

												if True == ifres2888 {
													ifres2887 = True

												} else {
													ifres2887 = False

												}

												ifres2886 = ifres2887

											} else {
												ifres2886 = False

											}

											var ifres2885 Obj

											if True == ifres2886 {
												ifres2885 = True

											} else {
												ifres2885 = False

											}

											ifres2884 = ifres2885

										} else {
											ifres2884 = False

										}

										var ifres2883 Obj

										if True == ifres2884 {
											ifres2883 = True

										} else {
											ifres2883 = False

										}

										ifres2882 = ifres2883

									} else {
										ifres2882 = False

									}

									var ifres2881 Obj

									if True == ifres2882 {
										ifres2881 = True

									} else {
										ifres2881 = False

									}

									ifres2880 = ifres2881

								} else {
									ifres2880 = False

								}

								var ifres2879 Obj

								if True == ifres2880 {
									ifres2879 = True

								} else {
									ifres2879 = False

								}

								ifres2878 = ifres2879

							} else {
								ifres2878 = False

							}

							var ifres2877 Obj

							if True == ifres2878 {
								ifres2877 = True

							} else {
								ifres2877 = False

							}

							ifres2876 = ifres2877

						} else {
							ifres2876 = False

						}

						var ifres2875 Obj

						if True == ifres2876 {
							ifres2875 = True

						} else {
							ifres2875 = False

						}

						ifres2874 = ifres2875

					} else {
						ifres2874 = False

					}

					var ifres2873 Obj

					if True == ifres2874 {
						ifres2873 = True

					} else {
						ifres2873 = False

					}

					ifres2872 = ifres2873

				} else {
					ifres2872 = False

				}

				var ifres2871 Obj

				if True == ifres2872 {
					ifres2871 = True

				} else {
					ifres2871 = False

				}

				ifres2870 = ifres2871

			} else {
				ifres2870 = False

			}

			if True == ifres2870 {
				tmp2865 := PrimHead(V442)

				tmp2866 := PrimTail(tmp2865)

				tmp2867 := PrimCons(True, tmp2866)

				tmp2868 := PrimTail(V442)

				tmp2869 := Call(__e, PrimNS2Value(symshen_4case_1form), tmp2868, V443)

				__e.Return(PrimCons(tmp2867, tmp2869))
				return

			} else {
				tmp2864 := PrimIsPair(V442)

				var ifres2818 Obj

				if True == tmp2864 {
					tmp2862 := PrimHead(V442)

					tmp2863 := PrimIsPair(tmp2862)

					var ifres2820 Obj

					if True == tmp2863 {
						tmp2859 := PrimHead(V442)

						tmp2860 := PrimHead(tmp2859)

						tmp2861 := PrimIsPair(tmp2860)

						var ifres2822 Obj

						if True == tmp2861 {
							tmp2855 := PrimHead(V442)

							tmp2856 := PrimHead(tmp2855)

							tmp2857 := PrimHead(tmp2856)

							tmp2858 := PrimEqual(sym_h, tmp2857)

							var ifres2824 Obj

							if True == tmp2858 {
								tmp2851 := PrimHead(V442)

								tmp2852 := PrimHead(tmp2851)

								tmp2853 := PrimTail(tmp2852)

								tmp2854 := PrimIsPair(tmp2853)

								var ifres2826 Obj

								if True == tmp2854 {
									tmp2846 := PrimHead(V442)

									tmp2847 := PrimHead(tmp2846)

									tmp2848 := PrimTail(tmp2847)

									tmp2849 := PrimHead(tmp2848)

									tmp2850 := PrimEqual(symshen_4tests, tmp2849)

									var ifres2828 Obj

									if True == tmp2850 {
										tmp2841 := PrimHead(V442)

										tmp2842 := PrimHead(tmp2841)

										tmp2843 := PrimTail(tmp2842)

										tmp2844 := PrimTail(tmp2843)

										tmp2845 := PrimEqual(Nil, tmp2844)

										var ifres2830 Obj

										if True == tmp2845 {
											tmp2838 := PrimHead(V442)

											tmp2839 := PrimTail(tmp2838)

											tmp2840 := PrimIsPair(tmp2839)

											var ifres2832 Obj

											if True == tmp2840 {
												tmp2834 := PrimHead(V442)

												tmp2835 := PrimTail(tmp2834)

												tmp2836 := PrimTail(tmp2835)

												tmp2837 := PrimEqual(Nil, tmp2836)

												var ifres2833 Obj

												if True == tmp2837 {
													ifres2833 = True

												} else {
													ifres2833 = False

												}

												ifres2832 = ifres2833

											} else {
												ifres2832 = False

											}

											var ifres2831 Obj

											if True == ifres2832 {
												ifres2831 = True

											} else {
												ifres2831 = False

											}

											ifres2830 = ifres2831

										} else {
											ifres2830 = False

										}

										var ifres2829 Obj

										if True == ifres2830 {
											ifres2829 = True

										} else {
											ifres2829 = False

										}

										ifres2828 = ifres2829

									} else {
										ifres2828 = False

									}

									var ifres2827 Obj

									if True == ifres2828 {
										ifres2827 = True

									} else {
										ifres2827 = False

									}

									ifres2826 = ifres2827

								} else {
									ifres2826 = False

								}

								var ifres2825 Obj

								if True == ifres2826 {
									ifres2825 = True

								} else {
									ifres2825 = False

								}

								ifres2824 = ifres2825

							} else {
								ifres2824 = False

							}

							var ifres2823 Obj

							if True == ifres2824 {
								ifres2823 = True

							} else {
								ifres2823 = False

							}

							ifres2822 = ifres2823

						} else {
							ifres2822 = False

						}

						var ifres2821 Obj

						if True == ifres2822 {
							ifres2821 = True

						} else {
							ifres2821 = False

						}

						ifres2820 = ifres2821

					} else {
						ifres2820 = False

					}

					var ifres2819 Obj

					if True == ifres2820 {
						ifres2819 = True

					} else {
						ifres2819 = False

					}

					ifres2818 = ifres2819

				} else {
					ifres2818 = False

				}

				if True == ifres2818 {
					tmp2815 := PrimHead(V442)

					tmp2816 := PrimTail(tmp2815)

					tmp2817 := PrimCons(True, tmp2816)

					__e.Return(PrimCons(tmp2817, Nil))
					return

				} else {
					tmp2814 := PrimIsPair(V442)

					var ifres2775 Obj

					if True == tmp2814 {
						tmp2812 := PrimHead(V442)

						tmp2813 := PrimIsPair(tmp2812)

						var ifres2777 Obj

						if True == tmp2813 {
							tmp2809 := PrimHead(V442)

							tmp2810 := PrimHead(tmp2809)

							tmp2811 := PrimIsPair(tmp2810)

							var ifres2779 Obj

							if True == tmp2811 {
								tmp2805 := PrimHead(V442)

								tmp2806 := PrimHead(tmp2805)

								tmp2807 := PrimHead(tmp2806)

								tmp2808 := PrimEqual(sym_h, tmp2807)

								var ifres2781 Obj

								if True == tmp2808 {
									tmp2801 := PrimHead(V442)

									tmp2802 := PrimHead(tmp2801)

									tmp2803 := PrimTail(tmp2802)

									tmp2804 := PrimIsPair(tmp2803)

									var ifres2783 Obj

									if True == tmp2804 {
										tmp2796 := PrimHead(V442)

										tmp2797 := PrimHead(tmp2796)

										tmp2798 := PrimTail(tmp2797)

										tmp2799 := PrimHead(tmp2798)

										tmp2800 := PrimEqual(symshen_4tests, tmp2799)

										var ifres2785 Obj

										if True == tmp2800 {
											tmp2793 := PrimHead(V442)

											tmp2794 := PrimTail(tmp2793)

											tmp2795 := PrimIsPair(tmp2794)

											var ifres2787 Obj

											if True == tmp2795 {
												tmp2789 := PrimHead(V442)

												tmp2790 := PrimTail(tmp2789)

												tmp2791 := PrimTail(tmp2790)

												tmp2792 := PrimEqual(Nil, tmp2791)

												var ifres2788 Obj

												if True == tmp2792 {
													ifres2788 = True

												} else {
													ifres2788 = False

												}

												ifres2787 = ifres2788

											} else {
												ifres2787 = False

											}

											var ifres2786 Obj

											if True == ifres2787 {
												ifres2786 = True

											} else {
												ifres2786 = False

											}

											ifres2785 = ifres2786

										} else {
											ifres2785 = False

										}

										var ifres2784 Obj

										if True == ifres2785 {
											ifres2784 = True

										} else {
											ifres2784 = False

										}

										ifres2783 = ifres2784

									} else {
										ifres2783 = False

									}

									var ifres2782 Obj

									if True == ifres2783 {
										ifres2782 = True

									} else {
										ifres2782 = False

									}

									ifres2781 = ifres2782

								} else {
									ifres2781 = False

								}

								var ifres2780 Obj

								if True == ifres2781 {
									ifres2780 = True

								} else {
									ifres2780 = False

								}

								ifres2779 = ifres2780

							} else {
								ifres2779 = False

							}

							var ifres2778 Obj

							if True == ifres2779 {
								ifres2778 = True

							} else {
								ifres2778 = False

							}

							ifres2777 = ifres2778

						} else {
							ifres2777 = False

						}

						var ifres2776 Obj

						if True == ifres2777 {
							ifres2776 = True

						} else {
							ifres2776 = False

						}

						ifres2775 = ifres2776

					} else {
						ifres2775 = False

					}

					if True == ifres2775 {
						tmp2765 := PrimHead(V442)

						tmp2766 := PrimHead(tmp2765)

						tmp2767 := PrimTail(tmp2766)

						tmp2768 := PrimTail(tmp2767)

						tmp2769 := Call(__e, PrimNS2Value(symshen_4embed_1and), tmp2768)

						tmp2770 := PrimHead(V442)

						tmp2771 := PrimTail(tmp2770)

						tmp2772 := PrimCons(tmp2769, tmp2771)

						tmp2773 := PrimTail(V442)

						tmp2774 := Call(__e, PrimNS2Value(symshen_4case_1form), tmp2773, V443)

						__e.Return(PrimCons(tmp2772, tmp2774))
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4case_1form)
						return
					}

				}

			}

		}

	}, 2)

	tmp2946 := Call(__e, PrimNS1Value(symns2_1set), symshen_4case_1form, tmp2760)

	_ = tmp2946

	tmp2947 := MakeNative(func(__e *ControlFlow) {
		V445 := __e.Get(1)
		_ = V445
		tmp2960 := PrimIsPair(V445)

		var ifres2956 Obj

		if True == tmp2960 {
			tmp2958 := PrimTail(V445)

			tmp2959 := PrimEqual(Nil, tmp2958)

			var ifres2957 Obj

			if True == tmp2959 {
				ifres2957 = True

			} else {
				ifres2957 = False

			}

			ifres2956 = ifres2957

		} else {
			ifres2956 = False

		}

		if True == ifres2956 {
			__e.Return(PrimHead(V445))
			return
		} else {
			tmp2955 := PrimIsPair(V445)

			if True == tmp2955 {
				tmp2950 := PrimHead(V445)

				tmp2951 := PrimTail(V445)

				tmp2952 := Call(__e, PrimNS2Value(symshen_4embed_1and), tmp2951)

				tmp2953 := PrimCons(tmp2952, Nil)

				tmp2954 := PrimCons(tmp2950, tmp2953)

				__e.Return(PrimCons(symand, tmp2954))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4embed_1and)
				return
			}

		}

	}, 1)

	tmp2961 := Call(__e, PrimNS1Value(symns2_1set), symshen_4embed_1and, tmp2947)

	_ = tmp2961

	tmp2962 := MakeNative(func(__e *ControlFlow) {
		V447 := __e.Get(1)
		_ = V447
		tmp2963 := PrimCons(V447, Nil)

		tmp2964 := PrimCons(symshen_4f__error, tmp2963)

		tmp2965 := PrimCons(tmp2964, Nil)

		__e.Return(PrimCons(True, tmp2965))
		return

	}, 1)

	tmp2966 := Call(__e, PrimNS1Value(symns2_1set), symshen_4err_1condition, tmp2962)

	_ = tmp2966

	tmp2967 := MakeNative(func(__e *ControlFlow) {
		V449 := __e.Get(1)
		_ = V449
		tmp2968 := Call(__e, PrimNS2Value(symshen_4app), V449, MakeString(": unexpected argument\n"), symshen_4a)

		tmp2969 := PrimStringConcat(MakeString("system function "), tmp2968)

		__e.Return(PrimSimpleError(tmp2969))
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4sys_1error, tmp2967)
	return

}, 0)
