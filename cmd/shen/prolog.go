package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp9115 := MakeNative(func(__e *ControlFlow) {
		V867 := __e.Get(1)
		_ = V867
		tmp9116 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5predicate_d_6 := __e.Get(1)
			_ = Parse__shen_4_5predicate_d_6
			tmp9131 := Call(__e, PrimNS2Value(symfail))

			tmp9132 := PrimEqual(tmp9131, Parse__shen_4_5predicate_d_6)

			tmp9133 := PrimNot(tmp9132)

			if True == tmp9133 {
				tmp9118 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5clauses_d_6 := __e.Get(1)
					_ = Parse__shen_4_5clauses_d_6
					tmp9127 := Call(__e, PrimNS2Value(symfail))

					tmp9128 := PrimEqual(tmp9127, Parse__shen_4_5clauses_d_6)

					tmp9129 := PrimNot(tmp9128)

					if True == tmp9129 {
						tmp9120 := PrimHead(Parse__shen_4_5clauses_d_6)

						tmp9121 := MakeNative(func(__e *ControlFlow) {
							Parse__X := __e.Get(1)
							_ = Parse__X
							tmp9122 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5predicate_d_6)

							__e.TailApply(PrimNS2Value(symshen_4insert_1predicate), tmp9122, Parse__X)
							return

						}, 1)

						tmp9123 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

						tmp9124 := Call(__e, PrimNS2Value(symmap), tmp9121, tmp9123)

						tmp9125 := Call(__e, PrimNS2Value(symshen_4prolog_1_6shen), tmp9124)

						tmp9126 := PrimHead(tmp9125)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9120, tmp9126)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9130 := Call(__e, PrimNS2Value(symshen_4_5clauses_d_6), Parse__shen_4_5predicate_d_6)

				__e.TailApply(tmp9118, tmp9130)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp9134 := Call(__e, PrimNS2Value(symshen_4_5predicate_d_6), V867)

		__e.TailApply(tmp9116, tmp9134)
		return

	}, 1)

	tmp9135 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5defprolog_6, tmp9115)

	_ = tmp9135

	tmp9136 := MakeNative(func(__e *ControlFlow) {
		V876 := __e.Get(1)
		_ = V876
		V877 := __e.Get(2)
		_ = V877
		tmp9155 := PrimIsPair(V877)

		var ifres9146 Obj

		if True == tmp9155 {
			tmp9153 := PrimTail(V877)

			tmp9154 := PrimIsPair(tmp9153)

			var ifres9148 Obj

			if True == tmp9154 {
				tmp9150 := PrimTail(V877)

				tmp9151 := PrimTail(tmp9150)

				tmp9152 := PrimEqual(Nil, tmp9151)

				var ifres9149 Obj

				if True == tmp9152 {
					ifres9149 = True

				} else {
					ifres9149 = False

				}

				ifres9148 = ifres9149

			} else {
				ifres9148 = False

			}

			var ifres9147 Obj

			if True == ifres9148 {
				ifres9147 = True

			} else {
				ifres9147 = False

			}

			ifres9146 = ifres9147

		} else {
			ifres9146 = False

		}

		if True == ifres9146 {
			tmp9140 := PrimHead(V877)

			tmp9141 := Call(__e, PrimNS2Value(symshen_4next_150), MakeNumber(50), tmp9140)

			tmp9142 := Call(__e, PrimNS2Value(symshen_4app), tmp9141, MakeString("\n"), symshen_4a)

			tmp9143 := PrimStringConcat(MakeString(" here:\n\n "), tmp9142)

			tmp9144 := Call(__e, PrimNS2Value(symshen_4app), V876, tmp9143, symshen_4a)

			tmp9145 := PrimStringConcat(MakeString("prolog syntax error in "), tmp9144)

			__e.Return(PrimSimpleError(tmp9145))
			return

		} else {
			tmp9138 := Call(__e, PrimNS2Value(symshen_4app), V876, MakeString("\n"), symshen_4a)

			tmp9139 := PrimStringConcat(MakeString("prolog syntax error in "), tmp9138)

			__e.Return(PrimSimpleError(tmp9139))
			return

		}

	}, 2)

	tmp9156 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1error, tmp9136)

	_ = tmp9156

	tmp9157 := MakeNative(func(__e *ControlFlow) {
		V884 := __e.Get(1)
		_ = V884
		V885 := __e.Get(2)
		_ = V885
		tmp9168 := PrimEqual(Nil, V885)

		if True == tmp9168 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp9167 := PrimEqual(MakeNumber(0), V884)

			if True == tmp9167 {
				__e.Return(MakeString(""))
				return
			} else {
				tmp9166 := PrimIsPair(V885)

				if True == tmp9166 {
					tmp9161 := PrimHead(V885)

					tmp9162 := Call(__e, PrimNS2Value(symshen_4decons_1string), tmp9161)

					tmp9163 := PrimNumberSubtract(V884, MakeNumber(1))

					tmp9164 := PrimTail(V885)

					tmp9165 := Call(__e, PrimNS2Value(symshen_4next_150), tmp9163, tmp9164)

					__e.Return(PrimStringConcat(tmp9162, tmp9165))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4next_150)
					return
				}

			}

		}

	}, 2)

	tmp9169 := Call(__e, PrimNS1Value(symns2_1set), symshen_4next_150, tmp9157)

	_ = tmp9169

	tmp9170 := MakeNative(func(__e *ControlFlow) {
		V887 := __e.Get(1)
		_ = V887
		tmp9192 := PrimIsPair(V887)

		var ifres9173 Obj

		if True == tmp9192 {
			tmp9190 := PrimHead(V887)

			tmp9191 := PrimEqual(symcons, tmp9190)

			var ifres9175 Obj

			if True == tmp9191 {
				tmp9188 := PrimTail(V887)

				tmp9189 := PrimIsPair(tmp9188)

				var ifres9177 Obj

				if True == tmp9189 {
					tmp9185 := PrimTail(V887)

					tmp9186 := PrimTail(tmp9185)

					tmp9187 := PrimIsPair(tmp9186)

					var ifres9179 Obj

					if True == tmp9187 {
						tmp9181 := PrimTail(V887)

						tmp9182 := PrimTail(tmp9181)

						tmp9183 := PrimTail(tmp9182)

						tmp9184 := PrimEqual(Nil, tmp9183)

						var ifres9180 Obj

						if True == tmp9184 {
							ifres9180 = True

						} else {
							ifres9180 = False

						}

						ifres9179 = ifres9180

					} else {
						ifres9179 = False

					}

					var ifres9178 Obj

					if True == ifres9179 {
						ifres9178 = True

					} else {
						ifres9178 = False

					}

					ifres9177 = ifres9178

				} else {
					ifres9177 = False

				}

				var ifres9176 Obj

				if True == ifres9177 {
					ifres9176 = True

				} else {
					ifres9176 = False

				}

				ifres9175 = ifres9176

			} else {
				ifres9175 = False

			}

			var ifres9174 Obj

			if True == ifres9175 {
				ifres9174 = True

			} else {
				ifres9174 = False

			}

			ifres9173 = ifres9174

		} else {
			ifres9173 = False

		}

		if True == ifres9173 {
			tmp9172 := Call(__e, PrimNS2Value(symshen_4eval_1cons), V887)

			__e.TailApply(PrimNS2Value(symshen_4app), tmp9172, MakeString(" "), symshen_4s)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4app), V887, MakeString(" "), symshen_4r)
			return
		}

	}, 1)

	tmp9193 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decons_1string, tmp9170)

	_ = tmp9193

	tmp9194 := MakeNative(func(__e *ControlFlow) {
		V890 := __e.Get(1)
		_ = V890
		V891 := __e.Get(2)
		_ = V891
		tmp9209 := PrimIsPair(V891)

		var ifres9200 Obj

		if True == tmp9209 {
			tmp9207 := PrimTail(V891)

			tmp9208 := PrimIsPair(tmp9207)

			var ifres9202 Obj

			if True == tmp9208 {
				tmp9204 := PrimTail(V891)

				tmp9205 := PrimTail(tmp9204)

				tmp9206 := PrimEqual(Nil, tmp9205)

				var ifres9203 Obj

				if True == tmp9206 {
					ifres9203 = True

				} else {
					ifres9203 = False

				}

				ifres9202 = ifres9203

			} else {
				ifres9202 = False

			}

			var ifres9201 Obj

			if True == ifres9202 {
				ifres9201 = True

			} else {
				ifres9201 = False

			}

			ifres9200 = ifres9201

		} else {
			ifres9200 = False

		}

		if True == ifres9200 {
			tmp9196 := PrimHead(V891)

			tmp9197 := PrimCons(V890, tmp9196)

			tmp9198 := PrimTail(V891)

			tmp9199 := PrimCons(sym_h_1, tmp9198)

			__e.Return(PrimCons(tmp9197, tmp9199))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4insert_1predicate)
			return
		}

	}, 2)

	tmp9210 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1predicate, tmp9194)

	_ = tmp9210

	tmp9211 := MakeNative(func(__e *ControlFlow) {
		V893 := __e.Get(1)
		_ = V893
		tmp9219 := PrimHead(V893)

		tmp9220 := PrimIsPair(tmp9219)

		if True == tmp9220 {
			tmp9213 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp9214 := Call(__e, PrimNS2Value(symshen_4tlhd), V893)

				tmp9215 := Call(__e, PrimNS2Value(symshen_4hdtl), V893)

				tmp9216 := Call(__e, PrimNS2Value(symshen_4pair), tmp9214, tmp9215)

				tmp9217 := PrimHead(tmp9216)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp9217, Parse__X)
				return

			}, 1)

			tmp9218 := Call(__e, PrimNS2Value(symshen_4hdhd), V893)

			__e.TailApply(tmp9213, tmp9218)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp9221 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5predicate_d_6, tmp9211)

	_ = tmp9221

	tmp9222 := MakeNative(func(__e *ControlFlow) {
		V895 := __e.Get(1)
		_ = V895
		tmp9223 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp9232 := Call(__e, PrimNS2Value(symfail))

			tmp9233 := PrimEqual(YaccParse, tmp9232)

			if True == tmp9233 {
				tmp9225 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp9228 := Call(__e, PrimNS2Value(symfail))

					tmp9229 := PrimEqual(tmp9228, Parse___5e_6)

					tmp9230 := PrimNot(tmp9229)

					if True == tmp9230 {
						tmp9227 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9227, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9231 := Call(__e, PrimNS2Value(sym_5e_6), V895)

				__e.TailApply(tmp9225, tmp9231)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp9234 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5clause_d_6 := __e.Get(1)
			_ = Parse__shen_4_5clause_d_6
			tmp9246 := Call(__e, PrimNS2Value(symfail))

			tmp9247 := PrimEqual(tmp9246, Parse__shen_4_5clause_d_6)

			tmp9248 := PrimNot(tmp9247)

			if True == tmp9248 {
				tmp9236 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5clauses_d_6 := __e.Get(1)
					_ = Parse__shen_4_5clauses_d_6
					tmp9242 := Call(__e, PrimNS2Value(symfail))

					tmp9243 := PrimEqual(tmp9242, Parse__shen_4_5clauses_d_6)

					tmp9244 := PrimNot(tmp9243)

					if True == tmp9244 {
						tmp9238 := PrimHead(Parse__shen_4_5clauses_d_6)

						tmp9239 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5clause_d_6)

						tmp9240 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

						tmp9241 := PrimCons(tmp9239, tmp9240)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9238, tmp9241)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9245 := Call(__e, PrimNS2Value(symshen_4_5clauses_d_6), Parse__shen_4_5clause_d_6)

				__e.TailApply(tmp9236, tmp9245)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp9249 := Call(__e, PrimNS2Value(symshen_4_5clause_d_6), V895)

		tmp9250 := Call(__e, tmp9234, tmp9249)

		__e.TailApply(tmp9223, tmp9250)
		return

	}, 1)

	tmp9251 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5clauses_d_6, tmp9222)

	_ = tmp9251

	tmp9252 := MakeNative(func(__e *ControlFlow) {
		V898 := __e.Get(1)
		_ = V898
		tmp9253 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5head_d_6 := __e.Get(1)
			_ = Parse__shen_4_5head_d_6
			tmp9283 := Call(__e, PrimNS2Value(symfail))

			tmp9284 := PrimEqual(tmp9283, Parse__shen_4_5head_d_6)

			tmp9285 := PrimNot(tmp9284)

			if True == tmp9285 {
				tmp9281 := PrimHead(Parse__shen_4_5head_d_6)

				tmp9282 := PrimIsPair(tmp9281)

				var ifres9277 Obj

				if True == tmp9282 {
					tmp9279 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5head_d_6)

					tmp9280 := PrimEqual(sym_5_1_1, tmp9279)

					var ifres9278 Obj

					if True == tmp9280 {
						ifres9278 = True

					} else {
						ifres9278 = False

					}

					ifres9277 = ifres9278

				} else {
					ifres9277 = False

				}

				if True == ifres9277 {
					tmp9256 := MakeNative(func(__e *ControlFlow) {
						NewStream896 := __e.Get(1)
						_ = NewStream896
						tmp9257 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5body_d_6 := __e.Get(1)
							_ = Parse__shen_4_5body_d_6
							tmp9270 := Call(__e, PrimNS2Value(symfail))

							tmp9271 := PrimEqual(tmp9270, Parse__shen_4_5body_d_6)

							tmp9272 := PrimNot(tmp9271)

							if True == tmp9272 {
								tmp9259 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5end_d_6 := __e.Get(1)
									_ = Parse__shen_4_5end_d_6
									tmp9266 := Call(__e, PrimNS2Value(symfail))

									tmp9267 := PrimEqual(tmp9266, Parse__shen_4_5end_d_6)

									tmp9268 := PrimNot(tmp9267)

									if True == tmp9268 {
										tmp9261 := PrimHead(Parse__shen_4_5end_d_6)

										tmp9262 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5head_d_6)

										tmp9263 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5body_d_6)

										tmp9264 := PrimCons(tmp9263, Nil)

										tmp9265 := PrimCons(tmp9262, tmp9264)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp9261, tmp9265)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp9269 := Call(__e, PrimNS2Value(symshen_4_5end_d_6), Parse__shen_4_5body_d_6)

								__e.TailApply(tmp9259, tmp9269)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp9273 := Call(__e, PrimNS2Value(symshen_4_5body_d_6), NewStream896)

						__e.TailApply(tmp9257, tmp9273)
						return

					}, 1)

					tmp9274 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5head_d_6)

					tmp9275 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5head_d_6)

					tmp9276 := Call(__e, PrimNS2Value(symshen_4pair), tmp9274, tmp9275)

					__e.TailApply(tmp9256, tmp9276)
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

		tmp9286 := Call(__e, PrimNS2Value(symshen_4_5head_d_6), V898)

		__e.TailApply(tmp9253, tmp9286)
		return

	}, 1)

	tmp9287 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5clause_d_6, tmp9252)

	_ = tmp9287

	tmp9288 := MakeNative(func(__e *ControlFlow) {
		V900 := __e.Get(1)
		_ = V900
		tmp9289 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp9298 := Call(__e, PrimNS2Value(symfail))

			tmp9299 := PrimEqual(YaccParse, tmp9298)

			if True == tmp9299 {
				tmp9291 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp9294 := Call(__e, PrimNS2Value(symfail))

					tmp9295 := PrimEqual(tmp9294, Parse___5e_6)

					tmp9296 := PrimNot(tmp9295)

					if True == tmp9296 {
						tmp9293 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9293, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9297 := Call(__e, PrimNS2Value(sym_5e_6), V900)

				__e.TailApply(tmp9291, tmp9297)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp9300 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5term_d_6 := __e.Get(1)
			_ = Parse__shen_4_5term_d_6
			tmp9312 := Call(__e, PrimNS2Value(symfail))

			tmp9313 := PrimEqual(tmp9312, Parse__shen_4_5term_d_6)

			tmp9314 := PrimNot(tmp9313)

			if True == tmp9314 {
				tmp9302 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5head_d_6 := __e.Get(1)
					_ = Parse__shen_4_5head_d_6
					tmp9308 := Call(__e, PrimNS2Value(symfail))

					tmp9309 := PrimEqual(tmp9308, Parse__shen_4_5head_d_6)

					tmp9310 := PrimNot(tmp9309)

					if True == tmp9310 {
						tmp9304 := PrimHead(Parse__shen_4_5head_d_6)

						tmp9305 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5term_d_6)

						tmp9306 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5head_d_6)

						tmp9307 := PrimCons(tmp9305, tmp9306)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9304, tmp9307)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9311 := Call(__e, PrimNS2Value(symshen_4_5head_d_6), Parse__shen_4_5term_d_6)

				__e.TailApply(tmp9302, tmp9311)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp9315 := Call(__e, PrimNS2Value(symshen_4_5term_d_6), V900)

		tmp9316 := Call(__e, tmp9300, tmp9315)

		__e.TailApply(tmp9289, tmp9316)
		return

	}, 1)

	tmp9317 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5head_d_6, tmp9288)

	_ = tmp9317

	tmp9318 := MakeNative(func(__e *ControlFlow) {
		V902 := __e.Get(1)
		_ = V902
		tmp9333 := PrimHead(V902)

		tmp9334 := PrimIsPair(tmp9333)

		if True == tmp9334 {
			tmp9320 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp9330 := PrimEqual(sym_5_1_1, Parse__X)

				tmp9331 := PrimNot(tmp9330)

				var ifres9327 Obj

				if True == tmp9331 {
					tmp9329 := Call(__e, PrimNS2Value(symshen_4legitimate_1term_2), Parse__X)

					var ifres9328 Obj

					if True == tmp9329 {
						ifres9328 = True

					} else {
						ifres9328 = False

					}

					ifres9327 = ifres9328

				} else {
					ifres9327 = False

				}

				if True == ifres9327 {
					tmp9322 := Call(__e, PrimNS2Value(symshen_4tlhd), V902)

					tmp9323 := Call(__e, PrimNS2Value(symshen_4hdtl), V902)

					tmp9324 := Call(__e, PrimNS2Value(symshen_4pair), tmp9322, tmp9323)

					tmp9325 := PrimHead(tmp9324)

					tmp9326 := Call(__e, PrimNS2Value(symshen_4eval_1cons), Parse__X)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp9325, tmp9326)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp9332 := Call(__e, PrimNS2Value(symshen_4hdhd), V902)

			__e.TailApply(tmp9320, tmp9332)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp9335 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5term_d_6, tmp9318)

	_ = tmp9335

	tmp9336 := MakeNative(func(__e *ControlFlow) {
		V908 := __e.Get(1)
		_ = V908
		tmp9426 := PrimIsPair(V908)

		var ifres9407 Obj

		if True == tmp9426 {
			tmp9424 := PrimHead(V908)

			tmp9425 := PrimEqual(symcons, tmp9424)

			var ifres9409 Obj

			if True == tmp9425 {
				tmp9422 := PrimTail(V908)

				tmp9423 := PrimIsPair(tmp9422)

				var ifres9411 Obj

				if True == tmp9423 {
					tmp9419 := PrimTail(V908)

					tmp9420 := PrimTail(tmp9419)

					tmp9421 := PrimIsPair(tmp9420)

					var ifres9413 Obj

					if True == tmp9421 {
						tmp9415 := PrimTail(V908)

						tmp9416 := PrimTail(tmp9415)

						tmp9417 := PrimTail(tmp9416)

						tmp9418 := PrimEqual(Nil, tmp9417)

						var ifres9414 Obj

						if True == tmp9418 {
							ifres9414 = True

						} else {
							ifres9414 = False

						}

						ifres9413 = ifres9414

					} else {
						ifres9413 = False

					}

					var ifres9412 Obj

					if True == ifres9413 {
						ifres9412 = True

					} else {
						ifres9412 = False

					}

					ifres9411 = ifres9412

				} else {
					ifres9411 = False

				}

				var ifres9410 Obj

				if True == ifres9411 {
					ifres9410 = True

				} else {
					ifres9410 = False

				}

				ifres9409 = ifres9410

			} else {
				ifres9409 = False

			}

			var ifres9408 Obj

			if True == ifres9409 {
				ifres9408 = True

			} else {
				ifres9408 = False

			}

			ifres9407 = ifres9408

		} else {
			ifres9407 = False

		}

		if True == ifres9407 {
			tmp9404 := PrimTail(V908)

			tmp9405 := PrimHead(tmp9404)

			tmp9406 := Call(__e, PrimNS2Value(symshen_4legitimate_1term_2), tmp9405)

			if True == tmp9406 {
				tmp9400 := PrimTail(V908)

				tmp9401 := PrimTail(tmp9400)

				tmp9402 := PrimHead(tmp9401)

				tmp9403 := Call(__e, PrimNS2Value(symshen_4legitimate_1term_2), tmp9402)

				if True == tmp9403 {
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

		} else {
			tmp9397 := PrimIsPair(V908)

			var ifres9372 Obj

			if True == tmp9397 {
				tmp9395 := PrimHead(V908)

				tmp9396 := PrimEqual(symmode, tmp9395)

				var ifres9374 Obj

				if True == tmp9396 {
					tmp9393 := PrimTail(V908)

					tmp9394 := PrimIsPair(tmp9393)

					var ifres9376 Obj

					if True == tmp9394 {
						tmp9390 := PrimTail(V908)

						tmp9391 := PrimTail(tmp9390)

						tmp9392 := PrimIsPair(tmp9391)

						var ifres9378 Obj

						if True == tmp9392 {
							tmp9386 := PrimTail(V908)

							tmp9387 := PrimTail(tmp9386)

							tmp9388 := PrimHead(tmp9387)

							tmp9389 := PrimEqual(sym_7, tmp9388)

							var ifres9380 Obj

							if True == tmp9389 {
								tmp9382 := PrimTail(V908)

								tmp9383 := PrimTail(tmp9382)

								tmp9384 := PrimTail(tmp9383)

								tmp9385 := PrimEqual(Nil, tmp9384)

								var ifres9381 Obj

								if True == tmp9385 {
									ifres9381 = True

								} else {
									ifres9381 = False

								}

								ifres9380 = ifres9381

							} else {
								ifres9380 = False

							}

							var ifres9379 Obj

							if True == ifres9380 {
								ifres9379 = True

							} else {
								ifres9379 = False

							}

							ifres9378 = ifres9379

						} else {
							ifres9378 = False

						}

						var ifres9377 Obj

						if True == ifres9378 {
							ifres9377 = True

						} else {
							ifres9377 = False

						}

						ifres9376 = ifres9377

					} else {
						ifres9376 = False

					}

					var ifres9375 Obj

					if True == ifres9376 {
						ifres9375 = True

					} else {
						ifres9375 = False

					}

					ifres9374 = ifres9375

				} else {
					ifres9374 = False

				}

				var ifres9373 Obj

				if True == ifres9374 {
					ifres9373 = True

				} else {
					ifres9373 = False

				}

				ifres9372 = ifres9373

			} else {
				ifres9372 = False

			}

			if True == ifres9372 {
				tmp9370 := PrimTail(V908)

				tmp9371 := PrimHead(tmp9370)

				__e.TailApply(PrimNS2Value(symshen_4legitimate_1term_2), tmp9371)
				return

			} else {
				tmp9369 := PrimIsPair(V908)

				var ifres9344 Obj

				if True == tmp9369 {
					tmp9367 := PrimHead(V908)

					tmp9368 := PrimEqual(symmode, tmp9367)

					var ifres9346 Obj

					if True == tmp9368 {
						tmp9365 := PrimTail(V908)

						tmp9366 := PrimIsPair(tmp9365)

						var ifres9348 Obj

						if True == tmp9366 {
							tmp9362 := PrimTail(V908)

							tmp9363 := PrimTail(tmp9362)

							tmp9364 := PrimIsPair(tmp9363)

							var ifres9350 Obj

							if True == tmp9364 {
								tmp9358 := PrimTail(V908)

								tmp9359 := PrimTail(tmp9358)

								tmp9360 := PrimHead(tmp9359)

								tmp9361 := PrimEqual(sym_1, tmp9360)

								var ifres9352 Obj

								if True == tmp9361 {
									tmp9354 := PrimTail(V908)

									tmp9355 := PrimTail(tmp9354)

									tmp9356 := PrimTail(tmp9355)

									tmp9357 := PrimEqual(Nil, tmp9356)

									var ifres9353 Obj

									if True == tmp9357 {
										ifres9353 = True

									} else {
										ifres9353 = False

									}

									ifres9352 = ifres9353

								} else {
									ifres9352 = False

								}

								var ifres9351 Obj

								if True == ifres9352 {
									ifres9351 = True

								} else {
									ifres9351 = False

								}

								ifres9350 = ifres9351

							} else {
								ifres9350 = False

							}

							var ifres9349 Obj

							if True == ifres9350 {
								ifres9349 = True

							} else {
								ifres9349 = False

							}

							ifres9348 = ifres9349

						} else {
							ifres9348 = False

						}

						var ifres9347 Obj

						if True == ifres9348 {
							ifres9347 = True

						} else {
							ifres9347 = False

						}

						ifres9346 = ifres9347

					} else {
						ifres9346 = False

					}

					var ifres9345 Obj

					if True == ifres9346 {
						ifres9345 = True

					} else {
						ifres9345 = False

					}

					ifres9344 = ifres9345

				} else {
					ifres9344 = False

				}

				if True == ifres9344 {
					tmp9342 := PrimTail(V908)

					tmp9343 := PrimHead(tmp9342)

					__e.TailApply(PrimNS2Value(symshen_4legitimate_1term_2), tmp9343)
					return

				} else {
					tmp9341 := PrimIsPair(V908)

					if True == tmp9341 {
						__e.Return(False)
						return
					} else {
						__e.Return(True)
						return
					}

				}

			}

		}

	}, 1)

	tmp9427 := Call(__e, PrimNS1Value(symns2_1set), symshen_4legitimate_1term_2, tmp9336)

	_ = tmp9427

	tmp9428 := MakeNative(func(__e *ControlFlow) {
		V910 := __e.Get(1)
		_ = V910
		tmp9483 := PrimIsPair(V910)

		var ifres9464 Obj

		if True == tmp9483 {
			tmp9481 := PrimHead(V910)

			tmp9482 := PrimEqual(symcons, tmp9481)

			var ifres9466 Obj

			if True == tmp9482 {
				tmp9479 := PrimTail(V910)

				tmp9480 := PrimIsPair(tmp9479)

				var ifres9468 Obj

				if True == tmp9480 {
					tmp9476 := PrimTail(V910)

					tmp9477 := PrimTail(tmp9476)

					tmp9478 := PrimIsPair(tmp9477)

					var ifres9470 Obj

					if True == tmp9478 {
						tmp9472 := PrimTail(V910)

						tmp9473 := PrimTail(tmp9472)

						tmp9474 := PrimTail(tmp9473)

						tmp9475 := PrimEqual(Nil, tmp9474)

						var ifres9471 Obj

						if True == tmp9475 {
							ifres9471 = True

						} else {
							ifres9471 = False

						}

						ifres9470 = ifres9471

					} else {
						ifres9470 = False

					}

					var ifres9469 Obj

					if True == ifres9470 {
						ifres9469 = True

					} else {
						ifres9469 = False

					}

					ifres9468 = ifres9469

				} else {
					ifres9468 = False

				}

				var ifres9467 Obj

				if True == ifres9468 {
					ifres9467 = True

				} else {
					ifres9467 = False

				}

				ifres9466 = ifres9467

			} else {
				ifres9466 = False

			}

			var ifres9465 Obj

			if True == ifres9466 {
				ifres9465 = True

			} else {
				ifres9465 = False

			}

			ifres9464 = ifres9465

		} else {
			ifres9464 = False

		}

		if True == ifres9464 {
			tmp9457 := PrimTail(V910)

			tmp9458 := PrimHead(tmp9457)

			tmp9459 := Call(__e, PrimNS2Value(symshen_4eval_1cons), tmp9458)

			tmp9460 := PrimTail(V910)

			tmp9461 := PrimTail(tmp9460)

			tmp9462 := PrimHead(tmp9461)

			tmp9463 := Call(__e, PrimNS2Value(symshen_4eval_1cons), tmp9462)

			__e.Return(PrimCons(tmp9459, tmp9463))
			return

		} else {
			tmp9456 := PrimIsPair(V910)

			var ifres9437 Obj

			if True == tmp9456 {
				tmp9454 := PrimHead(V910)

				tmp9455 := PrimEqual(symmode, tmp9454)

				var ifres9439 Obj

				if True == tmp9455 {
					tmp9452 := PrimTail(V910)

					tmp9453 := PrimIsPair(tmp9452)

					var ifres9441 Obj

					if True == tmp9453 {
						tmp9449 := PrimTail(V910)

						tmp9450 := PrimTail(tmp9449)

						tmp9451 := PrimIsPair(tmp9450)

						var ifres9443 Obj

						if True == tmp9451 {
							tmp9445 := PrimTail(V910)

							tmp9446 := PrimTail(tmp9445)

							tmp9447 := PrimTail(tmp9446)

							tmp9448 := PrimEqual(Nil, tmp9447)

							var ifres9444 Obj

							if True == tmp9448 {
								ifres9444 = True

							} else {
								ifres9444 = False

							}

							ifres9443 = ifres9444

						} else {
							ifres9443 = False

						}

						var ifres9442 Obj

						if True == ifres9443 {
							ifres9442 = True

						} else {
							ifres9442 = False

						}

						ifres9441 = ifres9442

					} else {
						ifres9441 = False

					}

					var ifres9440 Obj

					if True == ifres9441 {
						ifres9440 = True

					} else {
						ifres9440 = False

					}

					ifres9439 = ifres9440

				} else {
					ifres9439 = False

				}

				var ifres9438 Obj

				if True == ifres9439 {
					ifres9438 = True

				} else {
					ifres9438 = False

				}

				ifres9437 = ifres9438

			} else {
				ifres9437 = False

			}

			if True == ifres9437 {
				tmp9431 := PrimTail(V910)

				tmp9432 := PrimHead(tmp9431)

				tmp9433 := Call(__e, PrimNS2Value(symshen_4eval_1cons), tmp9432)

				tmp9434 := PrimTail(V910)

				tmp9435 := PrimTail(tmp9434)

				tmp9436 := PrimCons(tmp9433, tmp9435)

				__e.Return(PrimCons(symmode, tmp9436))
				return

			} else {
				__e.Return(V910)
				return
			}

		}

	}, 1)

	tmp9484 := Call(__e, PrimNS1Value(symns2_1set), symshen_4eval_1cons, tmp9428)

	_ = tmp9484

	tmp9485 := MakeNative(func(__e *ControlFlow) {
		V912 := __e.Get(1)
		_ = V912
		tmp9486 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp9495 := Call(__e, PrimNS2Value(symfail))

			tmp9496 := PrimEqual(YaccParse, tmp9495)

			if True == tmp9496 {
				tmp9488 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp9491 := Call(__e, PrimNS2Value(symfail))

					tmp9492 := PrimEqual(tmp9491, Parse___5e_6)

					tmp9493 := PrimNot(tmp9492)

					if True == tmp9493 {
						tmp9490 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9490, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9494 := Call(__e, PrimNS2Value(sym_5e_6), V912)

				__e.TailApply(tmp9488, tmp9494)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp9497 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5literal_d_6 := __e.Get(1)
			_ = Parse__shen_4_5literal_d_6
			tmp9509 := Call(__e, PrimNS2Value(symfail))

			tmp9510 := PrimEqual(tmp9509, Parse__shen_4_5literal_d_6)

			tmp9511 := PrimNot(tmp9510)

			if True == tmp9511 {
				tmp9499 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5body_d_6 := __e.Get(1)
					_ = Parse__shen_4_5body_d_6
					tmp9505 := Call(__e, PrimNS2Value(symfail))

					tmp9506 := PrimEqual(tmp9505, Parse__shen_4_5body_d_6)

					tmp9507 := PrimNot(tmp9506)

					if True == tmp9507 {
						tmp9501 := PrimHead(Parse__shen_4_5body_d_6)

						tmp9502 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5literal_d_6)

						tmp9503 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5body_d_6)

						tmp9504 := PrimCons(tmp9502, tmp9503)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp9501, tmp9504)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp9508 := Call(__e, PrimNS2Value(symshen_4_5body_d_6), Parse__shen_4_5literal_d_6)

				__e.TailApply(tmp9499, tmp9508)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp9512 := Call(__e, PrimNS2Value(symshen_4_5literal_d_6), V912)

		tmp9513 := Call(__e, tmp9497, tmp9512)

		__e.TailApply(tmp9486, tmp9513)
		return

	}, 1)

	tmp9514 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5body_d_6, tmp9485)

	_ = tmp9514

	tmp9515 := MakeNative(func(__e *ControlFlow) {
		V915 := __e.Get(1)
		_ = V915
		tmp9516 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp9529 := Call(__e, PrimNS2Value(symfail))

			tmp9530 := PrimEqual(YaccParse, tmp9529)

			if True == tmp9530 {
				tmp9527 := PrimHead(V915)

				tmp9528 := PrimIsPair(tmp9527)

				if True == tmp9528 {
					tmp9519 := MakeNative(func(__e *ControlFlow) {
						Parse__X := __e.Get(1)
						_ = Parse__X
						tmp9525 := PrimIsPair(Parse__X)

						if True == tmp9525 {
							tmp9521 := Call(__e, PrimNS2Value(symshen_4tlhd), V915)

							tmp9522 := Call(__e, PrimNS2Value(symshen_4hdtl), V915)

							tmp9523 := Call(__e, PrimNS2Value(symshen_4pair), tmp9521, tmp9522)

							tmp9524 := PrimHead(tmp9523)

							__e.TailApply(PrimNS2Value(symshen_4pair), tmp9524, Parse__X)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					}, 1)

					tmp9526 := Call(__e, PrimNS2Value(symshen_4hdhd), V915)

					__e.TailApply(tmp9519, tmp9526)
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

		tmp9546 := PrimHead(V915)

		tmp9547 := PrimIsPair(tmp9546)

		var ifres9542 Obj

		if True == tmp9547 {
			tmp9544 := Call(__e, PrimNS2Value(symshen_4hdhd), V915)

			tmp9545 := PrimEqual(sym_b, tmp9544)

			var ifres9543 Obj

			if True == tmp9545 {
				ifres9543 = True

			} else {
				ifres9543 = False

			}

			ifres9542 = ifres9543

		} else {
			ifres9542 = False

		}

		var ifres9531 Obj

		if True == ifres9542 {
			tmp9533 := MakeNative(func(__e *ControlFlow) {
				NewStream913 := __e.Get(1)
				_ = NewStream913
				tmp9534 := PrimHead(NewStream913)

				tmp9535 := PrimIntern(MakeString("Throwcontrol"))

				tmp9536 := PrimCons(tmp9535, Nil)

				tmp9537 := PrimCons(symcut, tmp9536)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp9534, tmp9537)
				return

			}, 1)

			tmp9538 := Call(__e, PrimNS2Value(symshen_4tlhd), V915)

			tmp9539 := Call(__e, PrimNS2Value(symshen_4hdtl), V915)

			tmp9540 := Call(__e, PrimNS2Value(symshen_4pair), tmp9538, tmp9539)

			tmp9541 := Call(__e, tmp9533, tmp9540)

			ifres9531 = tmp9541

		} else {
			tmp9532 := Call(__e, PrimNS2Value(symfail))

			ifres9531 = tmp9532

		}

		__e.TailApply(tmp9516, ifres9531)
		return

	}, 1)

	tmp9548 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5literal_d_6, tmp9515)

	_ = tmp9548

	tmp9549 := MakeNative(func(__e *ControlFlow) {
		V917 := __e.Get(1)
		_ = V917
		tmp9559 := PrimHead(V917)

		tmp9560 := PrimIsPair(tmp9559)

		if True == tmp9560 {
			tmp9551 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp9557 := PrimEqual(Parse__X, sym_k)

				if True == tmp9557 {
					tmp9553 := Call(__e, PrimNS2Value(symshen_4tlhd), V917)

					tmp9554 := Call(__e, PrimNS2Value(symshen_4hdtl), V917)

					tmp9555 := Call(__e, PrimNS2Value(symshen_4pair), tmp9553, tmp9554)

					tmp9556 := PrimHead(tmp9555)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp9556, Parse__X)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp9558 := Call(__e, PrimNS2Value(symshen_4hdhd), V917)

			__e.TailApply(tmp9551, tmp9558)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp9561 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5end_d_6, tmp9549)

	_ = tmp9561

	tmp9562 := MakeNative(func(__e *ControlFlow) {
		V921 := __e.Get(1)
		_ = V921
		V922 := __e.Get(2)
		_ = V922
		V923 := __e.Get(3)
		_ = V923
		tmp9563 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp9565 := PrimEqual(Result, False)

			if True == tmp9565 {
				__e.Return(V921)
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp9566 := Call(__e, PrimNS2Value(symthaw), V923)

		__e.TailApply(tmp9563, tmp9566)
		return

	}, 3)

	tmp9567 := Call(__e, PrimNS1Value(symns2_1set), symcut, tmp9562)

	_ = tmp9567

	tmp9568 := MakeNative(func(__e *ControlFlow) {
		V925 := __e.Get(1)
		_ = V925
		tmp9602 := PrimIsPair(V925)

		var ifres9583 Obj

		if True == tmp9602 {
			tmp9600 := PrimHead(V925)

			tmp9601 := PrimEqual(symmode, tmp9600)

			var ifres9585 Obj

			if True == tmp9601 {
				tmp9598 := PrimTail(V925)

				tmp9599 := PrimIsPair(tmp9598)

				var ifres9587 Obj

				if True == tmp9599 {
					tmp9595 := PrimTail(V925)

					tmp9596 := PrimTail(tmp9595)

					tmp9597 := PrimIsPair(tmp9596)

					var ifres9589 Obj

					if True == tmp9597 {
						tmp9591 := PrimTail(V925)

						tmp9592 := PrimTail(tmp9591)

						tmp9593 := PrimTail(tmp9592)

						tmp9594 := PrimEqual(Nil, tmp9593)

						var ifres9590 Obj

						if True == tmp9594 {
							ifres9590 = True

						} else {
							ifres9590 = False

						}

						ifres9589 = ifres9590

					} else {
						ifres9589 = False

					}

					var ifres9588 Obj

					if True == ifres9589 {
						ifres9588 = True

					} else {
						ifres9588 = False

					}

					ifres9587 = ifres9588

				} else {
					ifres9587 = False

				}

				var ifres9586 Obj

				if True == ifres9587 {
					ifres9586 = True

				} else {
					ifres9586 = False

				}

				ifres9585 = ifres9586

			} else {
				ifres9585 = False

			}

			var ifres9584 Obj

			if True == ifres9585 {
				ifres9584 = True

			} else {
				ifres9584 = False

			}

			ifres9583 = ifres9584

		} else {
			ifres9583 = False

		}

		if True == ifres9583 {
			__e.Return(V925)
			return
		} else {
			tmp9582 := PrimEqual(Nil, V925)

			if True == tmp9582 {
				__e.Return(Nil)
				return
			} else {
				tmp9581 := PrimIsPair(V925)

				if True == tmp9581 {
					tmp9572 := PrimHead(V925)

					tmp9573 := PrimCons(sym_7, Nil)

					tmp9574 := PrimCons(tmp9572, tmp9573)

					tmp9575 := PrimCons(symmode, tmp9574)

					tmp9576 := PrimTail(V925)

					tmp9577 := Call(__e, PrimNS2Value(symshen_4insert__modes), tmp9576)

					tmp9578 := PrimCons(sym_1, Nil)

					tmp9579 := PrimCons(tmp9577, tmp9578)

					tmp9580 := PrimCons(symmode, tmp9579)

					__e.Return(PrimCons(tmp9575, tmp9580))
					return

				} else {
					__e.Return(V925)
					return
				}

			}

		}

	}, 1)

	tmp9603 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert__modes, tmp9568)

	_ = tmp9603

	tmp9604 := MakeNative(func(__e *ControlFlow) {
		V927 := __e.Get(1)
		_ = V927
		tmp9605 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symeval), X)
			return
		}, 1)

		tmp9606 := Call(__e, PrimNS2Value(symshen_4prolog_1_6shen), V927)

		__e.TailApply(PrimNS2Value(symmap), tmp9605, tmp9606)
		return

	}, 1)

	tmp9607 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog, tmp9604)

	_ = tmp9607

	tmp9608 := MakeNative(func(__e *ControlFlow) {
		V929 := __e.Get(1)
		_ = V929
		tmp9609 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4compile__prolog__procedure), X)
			return
		}, 1)

		tmp9610 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4s_1prolog__clause), X)
			return
		}, 1)

		tmp9611 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4head__abstraction), X)
			return
		}, 1)

		tmp9612 := Call(__e, PrimNS2Value(symmapcan), tmp9611, V929)

		tmp9613 := Call(__e, PrimNS2Value(symmap), tmp9610, tmp9612)

		tmp9614 := Call(__e, PrimNS2Value(symshen_4group__clauses), tmp9613)

		__e.TailApply(PrimNS2Value(symmap), tmp9609, tmp9614)
		return

	}, 1)

	tmp9615 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1_6shen, tmp9608)

	_ = tmp9615

	tmp9616 := MakeNative(func(__e *ControlFlow) {
		V931 := __e.Get(1)
		_ = V931
		tmp9646 := PrimIsPair(V931)

		var ifres9626 Obj

		if True == tmp9646 {
			tmp9644 := PrimTail(V931)

			tmp9645 := PrimIsPair(tmp9644)

			var ifres9628 Obj

			if True == tmp9645 {
				tmp9641 := PrimTail(V931)

				tmp9642 := PrimHead(tmp9641)

				tmp9643 := PrimEqual(sym_h_1, tmp9642)

				var ifres9630 Obj

				if True == tmp9643 {
					tmp9638 := PrimTail(V931)

					tmp9639 := PrimTail(tmp9638)

					tmp9640 := PrimIsPair(tmp9639)

					var ifres9632 Obj

					if True == tmp9640 {
						tmp9634 := PrimTail(V931)

						tmp9635 := PrimTail(tmp9634)

						tmp9636 := PrimTail(tmp9635)

						tmp9637 := PrimEqual(Nil, tmp9636)

						var ifres9633 Obj

						if True == tmp9637 {
							ifres9633 = True

						} else {
							ifres9633 = False

						}

						ifres9632 = ifres9633

					} else {
						ifres9632 = False

					}

					var ifres9631 Obj

					if True == ifres9632 {
						ifres9631 = True

					} else {
						ifres9631 = False

					}

					ifres9630 = ifres9631

				} else {
					ifres9630 = False

				}

				var ifres9629 Obj

				if True == ifres9630 {
					ifres9629 = True

				} else {
					ifres9629 = False

				}

				ifres9628 = ifres9629

			} else {
				ifres9628 = False

			}

			var ifres9627 Obj

			if True == ifres9628 {
				ifres9627 = True

			} else {
				ifres9627 = False

			}

			ifres9626 = ifres9627

		} else {
			ifres9626 = False

		}

		if True == ifres9626 {
			tmp9618 := PrimHead(V931)

			tmp9619 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4s_1prolog__literal), X)
				return
			}, 1)

			tmp9620 := PrimTail(V931)

			tmp9621 := PrimTail(tmp9620)

			tmp9622 := PrimHead(tmp9621)

			tmp9623 := Call(__e, PrimNS2Value(symmap), tmp9619, tmp9622)

			tmp9624 := PrimCons(tmp9623, Nil)

			tmp9625 := PrimCons(sym_h_1, tmp9624)

			__e.Return(PrimCons(tmp9618, tmp9625))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4s_1prolog__clause)
			return
		}

	}, 1)

	tmp9647 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog__clause, tmp9616)

	_ = tmp9647

	tmp9648 := MakeNative(func(__e *ControlFlow) {
		V933 := __e.Get(1)
		_ = V933
		tmp9730 := PrimIsPair(V933)

		var ifres9702 Obj

		if True == tmp9730 {
			tmp9728 := PrimTail(V933)

			tmp9729 := PrimIsPair(tmp9728)

			var ifres9704 Obj

			if True == tmp9729 {
				tmp9725 := PrimTail(V933)

				tmp9726 := PrimHead(tmp9725)

				tmp9727 := PrimEqual(sym_h_1, tmp9726)

				var ifres9706 Obj

				if True == tmp9727 {
					tmp9722 := PrimTail(V933)

					tmp9723 := PrimTail(tmp9722)

					tmp9724 := PrimIsPair(tmp9723)

					var ifres9708 Obj

					if True == tmp9724 {
						tmp9718 := PrimTail(V933)

						tmp9719 := PrimTail(tmp9718)

						tmp9720 := PrimTail(tmp9719)

						tmp9721 := PrimEqual(Nil, tmp9720)

						var ifres9710 Obj

						if True == tmp9721 {
							tmp9712 := MakeNative(func(__e *ControlFlow) {
								tmp9713 := PrimHead(V933)

								tmp9714 := Call(__e, PrimNS2Value(symshen_4complexity__head), tmp9713)

								tmp9715 := PrimNS3Value(symshen_4_dmaxcomplexity_d)

								__e.Return(PrimLessThan(tmp9714, tmp9715))
								return

							}, 0)

							tmp9716 := MakeNative(func(__e *ControlFlow) {
								__ := __e.Get(1)
								_ = __
								__e.Return(False)
								return
							}, 1)

							tmp9717 := Call(__e, PrimNS1Value(symtry_1catch), tmp9712, tmp9716)

							var ifres9711 Obj

							if True == tmp9717 {
								ifres9711 = True

							} else {
								ifres9711 = False

							}

							ifres9710 = ifres9711

						} else {
							ifres9710 = False

						}

						var ifres9709 Obj

						if True == ifres9710 {
							ifres9709 = True

						} else {
							ifres9709 = False

						}

						ifres9708 = ifres9709

					} else {
						ifres9708 = False

					}

					var ifres9707 Obj

					if True == ifres9708 {
						ifres9707 = True

					} else {
						ifres9707 = False

					}

					ifres9706 = ifres9707

				} else {
					ifres9706 = False

				}

				var ifres9705 Obj

				if True == ifres9706 {
					ifres9705 = True

				} else {
					ifres9705 = False

				}

				ifres9704 = ifres9705

			} else {
				ifres9704 = False

			}

			var ifres9703 Obj

			if True == ifres9704 {
				ifres9703 = True

			} else {
				ifres9703 = False

			}

			ifres9702 = ifres9703

		} else {
			ifres9702 = False

		}

		if True == ifres9702 {
			__e.Return(PrimCons(V933, Nil))
			return
		} else {
			tmp9701 := PrimIsPair(V933)

			var ifres9677 Obj

			if True == tmp9701 {
				tmp9699 := PrimHead(V933)

				tmp9700 := PrimIsPair(tmp9699)

				var ifres9679 Obj

				if True == tmp9700 {
					tmp9697 := PrimTail(V933)

					tmp9698 := PrimIsPair(tmp9697)

					var ifres9681 Obj

					if True == tmp9698 {
						tmp9694 := PrimTail(V933)

						tmp9695 := PrimHead(tmp9694)

						tmp9696 := PrimEqual(sym_h_1, tmp9695)

						var ifres9683 Obj

						if True == tmp9696 {
							tmp9691 := PrimTail(V933)

							tmp9692 := PrimTail(tmp9691)

							tmp9693 := PrimIsPair(tmp9692)

							var ifres9685 Obj

							if True == tmp9693 {
								tmp9687 := PrimTail(V933)

								tmp9688 := PrimTail(tmp9687)

								tmp9689 := PrimTail(tmp9688)

								tmp9690 := PrimEqual(Nil, tmp9689)

								var ifres9686 Obj

								if True == tmp9690 {
									ifres9686 = True

								} else {
									ifres9686 = False

								}

								ifres9685 = ifres9686

							} else {
								ifres9685 = False

							}

							var ifres9684 Obj

							if True == ifres9685 {
								ifres9684 = True

							} else {
								ifres9684 = False

							}

							ifres9683 = ifres9684

						} else {
							ifres9683 = False

						}

						var ifres9682 Obj

						if True == ifres9683 {
							ifres9682 = True

						} else {
							ifres9682 = False

						}

						ifres9681 = ifres9682

					} else {
						ifres9681 = False

					}

					var ifres9680 Obj

					if True == ifres9681 {
						ifres9680 = True

					} else {
						ifres9680 = False

					}

					ifres9679 = ifres9680

				} else {
					ifres9679 = False

				}

				var ifres9678 Obj

				if True == ifres9679 {
					ifres9678 = True

				} else {
					ifres9678 = False

				}

				ifres9677 = ifres9678

			} else {
				ifres9677 = False

			}

			if True == ifres9677 {
				tmp9651 := MakeNative(func(__e *ControlFlow) {
					Terms := __e.Get(1)
					_ = Terms
					tmp9652 := MakeNative(func(__e *ControlFlow) {
						XTerms := __e.Get(1)
						_ = XTerms
						tmp9653 := MakeNative(func(__e *ControlFlow) {
							Literal := __e.Get(1)
							_ = Literal
							tmp9654 := MakeNative(func(__e *ControlFlow) {
								Clause := __e.Get(1)
								_ = Clause
								__e.Return(PrimCons(Clause, Nil))
								return
							}, 1)

							tmp9655 := PrimHead(V933)

							tmp9656 := PrimHead(tmp9655)

							tmp9657 := PrimCons(tmp9656, Terms)

							tmp9658 := PrimTail(V933)

							tmp9659 := PrimTail(tmp9658)

							tmp9660 := PrimHead(tmp9659)

							tmp9661 := PrimCons(Literal, tmp9660)

							tmp9662 := PrimCons(tmp9661, Nil)

							tmp9663 := PrimCons(sym_h_1, tmp9662)

							tmp9664 := PrimCons(tmp9657, tmp9663)

							__e.TailApply(tmp9654, tmp9664)
							return

						}, 1)

						tmp9665 := Call(__e, PrimNS2Value(symshen_4cons__form), Terms)

						tmp9666 := PrimCons(XTerms, Nil)

						tmp9667 := PrimCons(tmp9665, tmp9666)

						tmp9668 := PrimCons(symunify, tmp9667)

						__e.TailApply(tmp9653, tmp9668)
						return

					}, 1)

					tmp9669 := PrimHead(V933)

					tmp9670 := PrimTail(tmp9669)

					tmp9671 := Call(__e, PrimNS2Value(symshen_4remove__modes), tmp9670)

					tmp9672 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp9671)

					__e.TailApply(tmp9652, tmp9672)
					return

				}, 1)

				tmp9673 := MakeNative(func(__e *ControlFlow) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(PrimNS2Value(symgensym), symV)
					return
				}, 1)

				tmp9674 := PrimHead(V933)

				tmp9675 := PrimTail(tmp9674)

				tmp9676 := Call(__e, PrimNS2Value(symmap), tmp9673, tmp9675)

				__e.TailApply(tmp9651, tmp9676)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4head__abstraction)
				return
			}

		}

	}, 1)

	tmp9731 := Call(__e, PrimNS1Value(symns2_1set), symshen_4head__abstraction, tmp9648)

	_ = tmp9731

	tmp9732 := MakeNative(func(__e *ControlFlow) {
		V939 := __e.Get(1)
		_ = V939
		tmp9737 := PrimIsPair(V939)

		if True == tmp9737 {
			tmp9734 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4complexity), X)
				return
			}, 1)

			tmp9735 := PrimTail(V939)

			tmp9736 := Call(__e, PrimNS2Value(symmap), tmp9734, tmp9735)

			__e.TailApply(PrimNS2Value(symshen_4safe_1product), tmp9736)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4complexity__head)
			return
		}

	}, 1)

	tmp9738 := Call(__e, PrimNS1Value(symns2_1set), symshen_4complexity__head, tmp9732)

	_ = tmp9738

	tmp9739 := MakeNative(func(__e *ControlFlow) {
		V942 := __e.Get(1)
		_ = V942
		V943 := __e.Get(2)
		_ = V943
		__e.Return(PrimNumberMultiply(V942, V943))
		return
	}, 2)

	tmp9740 := Call(__e, PrimNS1Value(symns2_1set), symshen_4safe_1multiply, tmp9739)

	_ = tmp9740

	tmp9741 := MakeNative(func(__e *ControlFlow) {
		V952 := __e.Get(1)
		_ = V952
		tmp9976 := PrimIsPair(V952)

		var ifres9925 Obj

		if True == tmp9976 {
			tmp9974 := PrimHead(V952)

			tmp9975 := PrimEqual(symmode, tmp9974)

			var ifres9927 Obj

			if True == tmp9975 {
				tmp9972 := PrimTail(V952)

				tmp9973 := PrimIsPair(tmp9972)

				var ifres9929 Obj

				if True == tmp9973 {
					tmp9969 := PrimTail(V952)

					tmp9970 := PrimHead(tmp9969)

					tmp9971 := PrimIsPair(tmp9970)

					var ifres9931 Obj

					if True == tmp9971 {
						tmp9965 := PrimTail(V952)

						tmp9966 := PrimHead(tmp9965)

						tmp9967 := PrimHead(tmp9966)

						tmp9968 := PrimEqual(symmode, tmp9967)

						var ifres9933 Obj

						if True == tmp9968 {
							tmp9961 := PrimTail(V952)

							tmp9962 := PrimHead(tmp9961)

							tmp9963 := PrimTail(tmp9962)

							tmp9964 := PrimIsPair(tmp9963)

							var ifres9935 Obj

							if True == tmp9964 {
								tmp9956 := PrimTail(V952)

								tmp9957 := PrimHead(tmp9956)

								tmp9958 := PrimTail(tmp9957)

								tmp9959 := PrimTail(tmp9958)

								tmp9960 := PrimIsPair(tmp9959)

								var ifres9937 Obj

								if True == tmp9960 {
									tmp9950 := PrimTail(V952)

									tmp9951 := PrimHead(tmp9950)

									tmp9952 := PrimTail(tmp9951)

									tmp9953 := PrimTail(tmp9952)

									tmp9954 := PrimTail(tmp9953)

									tmp9955 := PrimEqual(Nil, tmp9954)

									var ifres9939 Obj

									if True == tmp9955 {
										tmp9947 := PrimTail(V952)

										tmp9948 := PrimTail(tmp9947)

										tmp9949 := PrimIsPair(tmp9948)

										var ifres9941 Obj

										if True == tmp9949 {
											tmp9943 := PrimTail(V952)

											tmp9944 := PrimTail(tmp9943)

											tmp9945 := PrimTail(tmp9944)

											tmp9946 := PrimEqual(Nil, tmp9945)

											var ifres9942 Obj

											if True == tmp9946 {
												ifres9942 = True

											} else {
												ifres9942 = False

											}

											ifres9941 = ifres9942

										} else {
											ifres9941 = False

										}

										var ifres9940 Obj

										if True == ifres9941 {
											ifres9940 = True

										} else {
											ifres9940 = False

										}

										ifres9939 = ifres9940

									} else {
										ifres9939 = False

									}

									var ifres9938 Obj

									if True == ifres9939 {
										ifres9938 = True

									} else {
										ifres9938 = False

									}

									ifres9937 = ifres9938

								} else {
									ifres9937 = False

								}

								var ifres9936 Obj

								if True == ifres9937 {
									ifres9936 = True

								} else {
									ifres9936 = False

								}

								ifres9935 = ifres9936

							} else {
								ifres9935 = False

							}

							var ifres9934 Obj

							if True == ifres9935 {
								ifres9934 = True

							} else {
								ifres9934 = False

							}

							ifres9933 = ifres9934

						} else {
							ifres9933 = False

						}

						var ifres9932 Obj

						if True == ifres9933 {
							ifres9932 = True

						} else {
							ifres9932 = False

						}

						ifres9931 = ifres9932

					} else {
						ifres9931 = False

					}

					var ifres9930 Obj

					if True == ifres9931 {
						ifres9930 = True

					} else {
						ifres9930 = False

					}

					ifres9929 = ifres9930

				} else {
					ifres9929 = False

				}

				var ifres9928 Obj

				if True == ifres9929 {
					ifres9928 = True

				} else {
					ifres9928 = False

				}

				ifres9927 = ifres9928

			} else {
				ifres9927 = False

			}

			var ifres9926 Obj

			if True == ifres9927 {
				ifres9926 = True

			} else {
				ifres9926 = False

			}

			ifres9925 = ifres9926

		} else {
			ifres9925 = False

		}

		if True == ifres9925 {
			tmp9923 := PrimTail(V952)

			tmp9924 := PrimHead(tmp9923)

			__e.TailApply(PrimNS2Value(symshen_4complexity), tmp9924)
			return

		} else {
			tmp9922 := PrimIsPair(V952)

			var ifres9892 Obj

			if True == tmp9922 {
				tmp9920 := PrimHead(V952)

				tmp9921 := PrimEqual(symmode, tmp9920)

				var ifres9894 Obj

				if True == tmp9921 {
					tmp9918 := PrimTail(V952)

					tmp9919 := PrimIsPair(tmp9918)

					var ifres9896 Obj

					if True == tmp9919 {
						tmp9915 := PrimTail(V952)

						tmp9916 := PrimHead(tmp9915)

						tmp9917 := PrimIsPair(tmp9916)

						var ifres9898 Obj

						if True == tmp9917 {
							tmp9912 := PrimTail(V952)

							tmp9913 := PrimTail(tmp9912)

							tmp9914 := PrimIsPair(tmp9913)

							var ifres9900 Obj

							if True == tmp9914 {
								tmp9908 := PrimTail(V952)

								tmp9909 := PrimTail(tmp9908)

								tmp9910 := PrimHead(tmp9909)

								tmp9911 := PrimEqual(sym_7, tmp9910)

								var ifres9902 Obj

								if True == tmp9911 {
									tmp9904 := PrimTail(V952)

									tmp9905 := PrimTail(tmp9904)

									tmp9906 := PrimTail(tmp9905)

									tmp9907 := PrimEqual(Nil, tmp9906)

									var ifres9903 Obj

									if True == tmp9907 {
										ifres9903 = True

									} else {
										ifres9903 = False

									}

									ifres9902 = ifres9903

								} else {
									ifres9902 = False

								}

								var ifres9901 Obj

								if True == ifres9902 {
									ifres9901 = True

								} else {
									ifres9901 = False

								}

								ifres9900 = ifres9901

							} else {
								ifres9900 = False

							}

							var ifres9899 Obj

							if True == ifres9900 {
								ifres9899 = True

							} else {
								ifres9899 = False

							}

							ifres9898 = ifres9899

						} else {
							ifres9898 = False

						}

						var ifres9897 Obj

						if True == ifres9898 {
							ifres9897 = True

						} else {
							ifres9897 = False

						}

						ifres9896 = ifres9897

					} else {
						ifres9896 = False

					}

					var ifres9895 Obj

					if True == ifres9896 {
						ifres9895 = True

					} else {
						ifres9895 = False

					}

					ifres9894 = ifres9895

				} else {
					ifres9894 = False

				}

				var ifres9893 Obj

				if True == ifres9894 {
					ifres9893 = True

				} else {
					ifres9893 = False

				}

				ifres9892 = ifres9893

			} else {
				ifres9892 = False

			}

			if True == ifres9892 {
				tmp9875 := PrimTail(V952)

				tmp9876 := PrimHead(tmp9875)

				tmp9877 := PrimHead(tmp9876)

				tmp9878 := PrimTail(V952)

				tmp9879 := PrimTail(tmp9878)

				tmp9880 := PrimCons(tmp9877, tmp9879)

				tmp9881 := PrimCons(symmode, tmp9880)

				tmp9882 := Call(__e, PrimNS2Value(symshen_4complexity), tmp9881)

				tmp9883 := PrimTail(V952)

				tmp9884 := PrimHead(tmp9883)

				tmp9885 := PrimTail(tmp9884)

				tmp9886 := PrimTail(V952)

				tmp9887 := PrimTail(tmp9886)

				tmp9888 := PrimCons(tmp9885, tmp9887)

				tmp9889 := PrimCons(symmode, tmp9888)

				tmp9890 := Call(__e, PrimNS2Value(symshen_4complexity), tmp9889)

				tmp9891 := Call(__e, PrimNS2Value(symshen_4safe_1multiply), tmp9882, tmp9890)

				__e.TailApply(PrimNS2Value(symshen_4safe_1multiply), MakeNumber(2), tmp9891)
				return

			} else {
				tmp9874 := PrimIsPair(V952)

				var ifres9844 Obj

				if True == tmp9874 {
					tmp9872 := PrimHead(V952)

					tmp9873 := PrimEqual(symmode, tmp9872)

					var ifres9846 Obj

					if True == tmp9873 {
						tmp9870 := PrimTail(V952)

						tmp9871 := PrimIsPair(tmp9870)

						var ifres9848 Obj

						if True == tmp9871 {
							tmp9867 := PrimTail(V952)

							tmp9868 := PrimHead(tmp9867)

							tmp9869 := PrimIsPair(tmp9868)

							var ifres9850 Obj

							if True == tmp9869 {
								tmp9864 := PrimTail(V952)

								tmp9865 := PrimTail(tmp9864)

								tmp9866 := PrimIsPair(tmp9865)

								var ifres9852 Obj

								if True == tmp9866 {
									tmp9860 := PrimTail(V952)

									tmp9861 := PrimTail(tmp9860)

									tmp9862 := PrimHead(tmp9861)

									tmp9863 := PrimEqual(sym_1, tmp9862)

									var ifres9854 Obj

									if True == tmp9863 {
										tmp9856 := PrimTail(V952)

										tmp9857 := PrimTail(tmp9856)

										tmp9858 := PrimTail(tmp9857)

										tmp9859 := PrimEqual(Nil, tmp9858)

										var ifres9855 Obj

										if True == tmp9859 {
											ifres9855 = True

										} else {
											ifres9855 = False

										}

										ifres9854 = ifres9855

									} else {
										ifres9854 = False

									}

									var ifres9853 Obj

									if True == ifres9854 {
										ifres9853 = True

									} else {
										ifres9853 = False

									}

									ifres9852 = ifres9853

								} else {
									ifres9852 = False

								}

								var ifres9851 Obj

								if True == ifres9852 {
									ifres9851 = True

								} else {
									ifres9851 = False

								}

								ifres9850 = ifres9851

							} else {
								ifres9850 = False

							}

							var ifres9849 Obj

							if True == ifres9850 {
								ifres9849 = True

							} else {
								ifres9849 = False

							}

							ifres9848 = ifres9849

						} else {
							ifres9848 = False

						}

						var ifres9847 Obj

						if True == ifres9848 {
							ifres9847 = True

						} else {
							ifres9847 = False

						}

						ifres9846 = ifres9847

					} else {
						ifres9846 = False

					}

					var ifres9845 Obj

					if True == ifres9846 {
						ifres9845 = True

					} else {
						ifres9845 = False

					}

					ifres9844 = ifres9845

				} else {
					ifres9844 = False

				}

				if True == ifres9844 {
					tmp9828 := PrimTail(V952)

					tmp9829 := PrimHead(tmp9828)

					tmp9830 := PrimHead(tmp9829)

					tmp9831 := PrimTail(V952)

					tmp9832 := PrimTail(tmp9831)

					tmp9833 := PrimCons(tmp9830, tmp9832)

					tmp9834 := PrimCons(symmode, tmp9833)

					tmp9835 := Call(__e, PrimNS2Value(symshen_4complexity), tmp9834)

					tmp9836 := PrimTail(V952)

					tmp9837 := PrimHead(tmp9836)

					tmp9838 := PrimTail(tmp9837)

					tmp9839 := PrimTail(V952)

					tmp9840 := PrimTail(tmp9839)

					tmp9841 := PrimCons(tmp9838, tmp9840)

					tmp9842 := PrimCons(symmode, tmp9841)

					tmp9843 := Call(__e, PrimNS2Value(symshen_4complexity), tmp9842)

					__e.TailApply(PrimNS2Value(symshen_4safe_1multiply), tmp9835, tmp9843)
					return

				} else {
					tmp9827 := PrimIsPair(V952)

					var ifres9803 Obj

					if True == tmp9827 {
						tmp9825 := PrimHead(V952)

						tmp9826 := PrimEqual(symmode, tmp9825)

						var ifres9805 Obj

						if True == tmp9826 {
							tmp9823 := PrimTail(V952)

							tmp9824 := PrimIsPair(tmp9823)

							var ifres9807 Obj

							if True == tmp9824 {
								tmp9820 := PrimTail(V952)

								tmp9821 := PrimTail(tmp9820)

								tmp9822 := PrimIsPair(tmp9821)

								var ifres9809 Obj

								if True == tmp9822 {
									tmp9816 := PrimTail(V952)

									tmp9817 := PrimTail(tmp9816)

									tmp9818 := PrimTail(tmp9817)

									tmp9819 := PrimEqual(Nil, tmp9818)

									var ifres9811 Obj

									if True == tmp9819 {
										tmp9813 := PrimTail(V952)

										tmp9814 := PrimHead(tmp9813)

										tmp9815 := PrimIsVariable(tmp9814)

										var ifres9812 Obj

										if True == tmp9815 {
											ifres9812 = True

										} else {
											ifres9812 = False

										}

										ifres9811 = ifres9812

									} else {
										ifres9811 = False

									}

									var ifres9810 Obj

									if True == ifres9811 {
										ifres9810 = True

									} else {
										ifres9810 = False

									}

									ifres9809 = ifres9810

								} else {
									ifres9809 = False

								}

								var ifres9808 Obj

								if True == ifres9809 {
									ifres9808 = True

								} else {
									ifres9808 = False

								}

								ifres9807 = ifres9808

							} else {
								ifres9807 = False

							}

							var ifres9806 Obj

							if True == ifres9807 {
								ifres9806 = True

							} else {
								ifres9806 = False

							}

							ifres9805 = ifres9806

						} else {
							ifres9805 = False

						}

						var ifres9804 Obj

						if True == ifres9805 {
							ifres9804 = True

						} else {
							ifres9804 = False

						}

						ifres9803 = ifres9804

					} else {
						ifres9803 = False

					}

					if True == ifres9803 {
						__e.Return(MakeNumber(1))
						return
					} else {
						tmp9802 := PrimIsPair(V952)

						var ifres9777 Obj

						if True == tmp9802 {
							tmp9800 := PrimHead(V952)

							tmp9801 := PrimEqual(symmode, tmp9800)

							var ifres9779 Obj

							if True == tmp9801 {
								tmp9798 := PrimTail(V952)

								tmp9799 := PrimIsPair(tmp9798)

								var ifres9781 Obj

								if True == tmp9799 {
									tmp9795 := PrimTail(V952)

									tmp9796 := PrimTail(tmp9795)

									tmp9797 := PrimIsPair(tmp9796)

									var ifres9783 Obj

									if True == tmp9797 {
										tmp9791 := PrimTail(V952)

										tmp9792 := PrimTail(tmp9791)

										tmp9793 := PrimHead(tmp9792)

										tmp9794 := PrimEqual(sym_7, tmp9793)

										var ifres9785 Obj

										if True == tmp9794 {
											tmp9787 := PrimTail(V952)

											tmp9788 := PrimTail(tmp9787)

											tmp9789 := PrimTail(tmp9788)

											tmp9790 := PrimEqual(Nil, tmp9789)

											var ifres9786 Obj

											if True == tmp9790 {
												ifres9786 = True

											} else {
												ifres9786 = False

											}

											ifres9785 = ifres9786

										} else {
											ifres9785 = False

										}

										var ifres9784 Obj

										if True == ifres9785 {
											ifres9784 = True

										} else {
											ifres9784 = False

										}

										ifres9783 = ifres9784

									} else {
										ifres9783 = False

									}

									var ifres9782 Obj

									if True == ifres9783 {
										ifres9782 = True

									} else {
										ifres9782 = False

									}

									ifres9781 = ifres9782

								} else {
									ifres9781 = False

								}

								var ifres9780 Obj

								if True == ifres9781 {
									ifres9780 = True

								} else {
									ifres9780 = False

								}

								ifres9779 = ifres9780

							} else {
								ifres9779 = False

							}

							var ifres9778 Obj

							if True == ifres9779 {
								ifres9778 = True

							} else {
								ifres9778 = False

							}

							ifres9777 = ifres9778

						} else {
							ifres9777 = False

						}

						if True == ifres9777 {
							__e.Return(MakeNumber(2))
							return
						} else {
							tmp9776 := PrimIsPair(V952)

							var ifres9751 Obj

							if True == tmp9776 {
								tmp9774 := PrimHead(V952)

								tmp9775 := PrimEqual(symmode, tmp9774)

								var ifres9753 Obj

								if True == tmp9775 {
									tmp9772 := PrimTail(V952)

									tmp9773 := PrimIsPair(tmp9772)

									var ifres9755 Obj

									if True == tmp9773 {
										tmp9769 := PrimTail(V952)

										tmp9770 := PrimTail(tmp9769)

										tmp9771 := PrimIsPair(tmp9770)

										var ifres9757 Obj

										if True == tmp9771 {
											tmp9765 := PrimTail(V952)

											tmp9766 := PrimTail(tmp9765)

											tmp9767 := PrimHead(tmp9766)

											tmp9768 := PrimEqual(sym_1, tmp9767)

											var ifres9759 Obj

											if True == tmp9768 {
												tmp9761 := PrimTail(V952)

												tmp9762 := PrimTail(tmp9761)

												tmp9763 := PrimTail(tmp9762)

												tmp9764 := PrimEqual(Nil, tmp9763)

												var ifres9760 Obj

												if True == tmp9764 {
													ifres9760 = True

												} else {
													ifres9760 = False

												}

												ifres9759 = ifres9760

											} else {
												ifres9759 = False

											}

											var ifres9758 Obj

											if True == ifres9759 {
												ifres9758 = True

											} else {
												ifres9758 = False

											}

											ifres9757 = ifres9758

										} else {
											ifres9757 = False

										}

										var ifres9756 Obj

										if True == ifres9757 {
											ifres9756 = True

										} else {
											ifres9756 = False

										}

										ifres9755 = ifres9756

									} else {
										ifres9755 = False

									}

									var ifres9754 Obj

									if True == ifres9755 {
										ifres9754 = True

									} else {
										ifres9754 = False

									}

									ifres9753 = ifres9754

								} else {
									ifres9753 = False

								}

								var ifres9752 Obj

								if True == ifres9753 {
									ifres9752 = True

								} else {
									ifres9752 = False

								}

								ifres9751 = ifres9752

							} else {
								ifres9751 = False

							}

							if True == ifres9751 {
								__e.Return(MakeNumber(1))
								return
							} else {
								tmp9748 := PrimCons(sym_7, Nil)

								tmp9749 := PrimCons(V952, tmp9748)

								tmp9750 := PrimCons(symmode, tmp9749)

								__e.TailApply(PrimNS2Value(symshen_4complexity), tmp9750)
								return

							}

						}

					}

				}

			}

		}

	}, 1)

	tmp9977 := Call(__e, PrimNS1Value(symns2_1set), symshen_4complexity, tmp9741)

	_ = tmp9977

	tmp9978 := MakeNative(func(__e *ControlFlow) {
		V954 := __e.Get(1)
		_ = V954
		tmp9985 := PrimEqual(Nil, V954)

		if True == tmp9985 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp9984 := PrimIsPair(V954)

			if True == tmp9984 {
				tmp9981 := PrimHead(V954)

				tmp9982 := PrimTail(V954)

				tmp9983 := Call(__e, PrimNS2Value(symshen_4safe_1product), tmp9982)

				__e.TailApply(PrimNS2Value(symshen_4safe_1multiply), tmp9981, tmp9983)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4safe_1product)
				return
			}

		}

	}, 1)

	tmp9986 := Call(__e, PrimNS1Value(symns2_1set), symshen_4safe_1product, tmp9978)

	_ = tmp9986

	tmp9987 := MakeNative(func(__e *ControlFlow) {
		V956 := __e.Get(1)
		_ = V956
		tmp10085 := PrimIsPair(V956)

		var ifres10066 Obj

		if True == tmp10085 {
			tmp10083 := PrimHead(V956)

			tmp10084 := PrimEqual(symis, tmp10083)

			var ifres10068 Obj

			if True == tmp10084 {
				tmp10081 := PrimTail(V956)

				tmp10082 := PrimIsPair(tmp10081)

				var ifres10070 Obj

				if True == tmp10082 {
					tmp10078 := PrimTail(V956)

					tmp10079 := PrimTail(tmp10078)

					tmp10080 := PrimIsPair(tmp10079)

					var ifres10072 Obj

					if True == tmp10080 {
						tmp10074 := PrimTail(V956)

						tmp10075 := PrimTail(tmp10074)

						tmp10076 := PrimTail(tmp10075)

						tmp10077 := PrimEqual(Nil, tmp10076)

						var ifres10073 Obj

						if True == tmp10077 {
							ifres10073 = True

						} else {
							ifres10073 = False

						}

						ifres10072 = ifres10073

					} else {
						ifres10072 = False

					}

					var ifres10071 Obj

					if True == ifres10072 {
						ifres10071 = True

					} else {
						ifres10071 = False

					}

					ifres10070 = ifres10071

				} else {
					ifres10070 = False

				}

				var ifres10069 Obj

				if True == ifres10070 {
					ifres10069 = True

				} else {
					ifres10069 = False

				}

				ifres10068 = ifres10069

			} else {
				ifres10068 = False

			}

			var ifres10067 Obj

			if True == ifres10068 {
				ifres10067 = True

			} else {
				ifres10067 = False

			}

			ifres10066 = ifres10067

		} else {
			ifres10066 = False

		}

		if True == ifres10066 {
			tmp10058 := PrimTail(V956)

			tmp10059 := PrimHead(tmp10058)

			tmp10060 := PrimTail(V956)

			tmp10061 := PrimTail(tmp10060)

			tmp10062 := PrimHead(tmp10061)

			tmp10063 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10062, symProcessN)

			tmp10064 := PrimCons(tmp10063, Nil)

			tmp10065 := PrimCons(tmp10059, tmp10064)

			__e.Return(PrimCons(symbind, tmp10065))
			return

		} else {
			tmp10057 := PrimIsPair(V956)

			var ifres10044 Obj

			if True == tmp10057 {
				tmp10055 := PrimHead(V956)

				tmp10056 := PrimEqual(symwhen, tmp10055)

				var ifres10046 Obj

				if True == tmp10056 {
					tmp10053 := PrimTail(V956)

					tmp10054 := PrimIsPair(tmp10053)

					var ifres10048 Obj

					if True == tmp10054 {
						tmp10050 := PrimTail(V956)

						tmp10051 := PrimTail(tmp10050)

						tmp10052 := PrimEqual(Nil, tmp10051)

						var ifres10049 Obj

						if True == tmp10052 {
							ifres10049 = True

						} else {
							ifres10049 = False

						}

						ifres10048 = ifres10049

					} else {
						ifres10048 = False

					}

					var ifres10047 Obj

					if True == ifres10048 {
						ifres10047 = True

					} else {
						ifres10047 = False

					}

					ifres10046 = ifres10047

				} else {
					ifres10046 = False

				}

				var ifres10045 Obj

				if True == ifres10046 {
					ifres10045 = True

				} else {
					ifres10045 = False

				}

				ifres10044 = ifres10045

			} else {
				ifres10044 = False

			}

			if True == ifres10044 {
				tmp10040 := PrimTail(V956)

				tmp10041 := PrimHead(tmp10040)

				tmp10042 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10041, symProcessN)

				tmp10043 := PrimCons(tmp10042, Nil)

				__e.Return(PrimCons(symfwhen, tmp10043))
				return

			} else {
				tmp10039 := PrimIsPair(V956)

				var ifres10020 Obj

				if True == tmp10039 {
					tmp10037 := PrimHead(V956)

					tmp10038 := PrimEqual(symbind, tmp10037)

					var ifres10022 Obj

					if True == tmp10038 {
						tmp10035 := PrimTail(V956)

						tmp10036 := PrimIsPair(tmp10035)

						var ifres10024 Obj

						if True == tmp10036 {
							tmp10032 := PrimTail(V956)

							tmp10033 := PrimTail(tmp10032)

							tmp10034 := PrimIsPair(tmp10033)

							var ifres10026 Obj

							if True == tmp10034 {
								tmp10028 := PrimTail(V956)

								tmp10029 := PrimTail(tmp10028)

								tmp10030 := PrimTail(tmp10029)

								tmp10031 := PrimEqual(Nil, tmp10030)

								var ifres10027 Obj

								if True == tmp10031 {
									ifres10027 = True

								} else {
									ifres10027 = False

								}

								ifres10026 = ifres10027

							} else {
								ifres10026 = False

							}

							var ifres10025 Obj

							if True == ifres10026 {
								ifres10025 = True

							} else {
								ifres10025 = False

							}

							ifres10024 = ifres10025

						} else {
							ifres10024 = False

						}

						var ifres10023 Obj

						if True == ifres10024 {
							ifres10023 = True

						} else {
							ifres10023 = False

						}

						ifres10022 = ifres10023

					} else {
						ifres10022 = False

					}

					var ifres10021 Obj

					if True == ifres10022 {
						ifres10021 = True

					} else {
						ifres10021 = False

					}

					ifres10020 = ifres10021

				} else {
					ifres10020 = False

				}

				if True == ifres10020 {
					tmp10012 := PrimTail(V956)

					tmp10013 := PrimHead(tmp10012)

					tmp10014 := PrimTail(V956)

					tmp10015 := PrimTail(tmp10014)

					tmp10016 := PrimHead(tmp10015)

					tmp10017 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10016, symProcessN)

					tmp10018 := PrimCons(tmp10017, Nil)

					tmp10019 := PrimCons(tmp10013, tmp10018)

					__e.Return(PrimCons(symbind, tmp10019))
					return

				} else {
					tmp10011 := PrimIsPair(V956)

					var ifres9998 Obj

					if True == tmp10011 {
						tmp10009 := PrimHead(V956)

						tmp10010 := PrimEqual(symfwhen, tmp10009)

						var ifres10000 Obj

						if True == tmp10010 {
							tmp10007 := PrimTail(V956)

							tmp10008 := PrimIsPair(tmp10007)

							var ifres10002 Obj

							if True == tmp10008 {
								tmp10004 := PrimTail(V956)

								tmp10005 := PrimTail(tmp10004)

								tmp10006 := PrimEqual(Nil, tmp10005)

								var ifres10003 Obj

								if True == tmp10006 {
									ifres10003 = True

								} else {
									ifres10003 = False

								}

								ifres10002 = ifres10003

							} else {
								ifres10002 = False

							}

							var ifres10001 Obj

							if True == ifres10002 {
								ifres10001 = True

							} else {
								ifres10001 = False

							}

							ifres10000 = ifres10001

						} else {
							ifres10000 = False

						}

						var ifres9999 Obj

						if True == ifres10000 {
							ifres9999 = True

						} else {
							ifres9999 = False

						}

						ifres9998 = ifres9999

					} else {
						ifres9998 = False

					}

					if True == ifres9998 {
						tmp9994 := PrimTail(V956)

						tmp9995 := PrimHead(tmp9994)

						tmp9996 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp9995, symProcessN)

						tmp9997 := PrimCons(tmp9996, Nil)

						__e.Return(PrimCons(symfwhen, tmp9997))
						return

					} else {
						tmp9993 := PrimIsPair(V956)

						if True == tmp9993 {
							__e.Return(V956)
							return
						} else {
							__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4s_1prolog__literal)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp10086 := Call(__e, PrimNS1Value(symns2_1set), symshen_4s_1prolog__literal, tmp9987)

	_ = tmp10086

	tmp10087 := MakeNative(func(__e *ControlFlow) {
		V963 := __e.Get(1)
		_ = V963
		V964 := __e.Get(2)
		_ = V964
		tmp10168 := PrimIsVariable(V963)

		if True == tmp10168 {
			tmp10166 := PrimCons(V964, Nil)

			tmp10167 := PrimCons(V963, tmp10166)

			__e.Return(PrimCons(symshen_4deref, tmp10167))
			return

		} else {
			tmp10165 := PrimIsPair(V963)

			var ifres10146 Obj

			if True == tmp10165 {
				tmp10163 := PrimHead(V963)

				tmp10164 := PrimEqual(symlambda, tmp10163)

				var ifres10148 Obj

				if True == tmp10164 {
					tmp10161 := PrimTail(V963)

					tmp10162 := PrimIsPair(tmp10161)

					var ifres10150 Obj

					if True == tmp10162 {
						tmp10158 := PrimTail(V963)

						tmp10159 := PrimTail(tmp10158)

						tmp10160 := PrimIsPair(tmp10159)

						var ifres10152 Obj

						if True == tmp10160 {
							tmp10154 := PrimTail(V963)

							tmp10155 := PrimTail(tmp10154)

							tmp10156 := PrimTail(tmp10155)

							tmp10157 := PrimEqual(Nil, tmp10156)

							var ifres10153 Obj

							if True == tmp10157 {
								ifres10153 = True

							} else {
								ifres10153 = False

							}

							ifres10152 = ifres10153

						} else {
							ifres10152 = False

						}

						var ifres10151 Obj

						if True == ifres10152 {
							ifres10151 = True

						} else {
							ifres10151 = False

						}

						ifres10150 = ifres10151

					} else {
						ifres10150 = False

					}

					var ifres10149 Obj

					if True == ifres10150 {
						ifres10149 = True

					} else {
						ifres10149 = False

					}

					ifres10148 = ifres10149

				} else {
					ifres10148 = False

				}

				var ifres10147 Obj

				if True == ifres10148 {
					ifres10147 = True

				} else {
					ifres10147 = False

				}

				ifres10146 = ifres10147

			} else {
				ifres10146 = False

			}

			if True == ifres10146 {
				tmp10138 := PrimTail(V963)

				tmp10139 := PrimHead(tmp10138)

				tmp10140 := PrimTail(V963)

				tmp10141 := PrimTail(tmp10140)

				tmp10142 := PrimHead(tmp10141)

				tmp10143 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10142, V964)

				tmp10144 := PrimCons(tmp10143, Nil)

				tmp10145 := PrimCons(tmp10139, tmp10144)

				__e.Return(PrimCons(symlambda, tmp10145))
				return

			} else {
				tmp10137 := PrimIsPair(V963)

				var ifres10111 Obj

				if True == tmp10137 {
					tmp10135 := PrimHead(V963)

					tmp10136 := PrimEqual(symlet, tmp10135)

					var ifres10113 Obj

					if True == tmp10136 {
						tmp10133 := PrimTail(V963)

						tmp10134 := PrimIsPair(tmp10133)

						var ifres10115 Obj

						if True == tmp10134 {
							tmp10130 := PrimTail(V963)

							tmp10131 := PrimTail(tmp10130)

							tmp10132 := PrimIsPair(tmp10131)

							var ifres10117 Obj

							if True == tmp10132 {
								tmp10126 := PrimTail(V963)

								tmp10127 := PrimTail(tmp10126)

								tmp10128 := PrimTail(tmp10127)

								tmp10129 := PrimIsPair(tmp10128)

								var ifres10119 Obj

								if True == tmp10129 {
									tmp10121 := PrimTail(V963)

									tmp10122 := PrimTail(tmp10121)

									tmp10123 := PrimTail(tmp10122)

									tmp10124 := PrimTail(tmp10123)

									tmp10125 := PrimEqual(Nil, tmp10124)

									var ifres10120 Obj

									if True == tmp10125 {
										ifres10120 = True

									} else {
										ifres10120 = False

									}

									ifres10119 = ifres10120

								} else {
									ifres10119 = False

								}

								var ifres10118 Obj

								if True == ifres10119 {
									ifres10118 = True

								} else {
									ifres10118 = False

								}

								ifres10117 = ifres10118

							} else {
								ifres10117 = False

							}

							var ifres10116 Obj

							if True == ifres10117 {
								ifres10116 = True

							} else {
								ifres10116 = False

							}

							ifres10115 = ifres10116

						} else {
							ifres10115 = False

						}

						var ifres10114 Obj

						if True == ifres10115 {
							ifres10114 = True

						} else {
							ifres10114 = False

						}

						ifres10113 = ifres10114

					} else {
						ifres10113 = False

					}

					var ifres10112 Obj

					if True == ifres10113 {
						ifres10112 = True

					} else {
						ifres10112 = False

					}

					ifres10111 = ifres10112

				} else {
					ifres10111 = False

				}

				if True == ifres10111 {
					tmp10097 := PrimTail(V963)

					tmp10098 := PrimHead(tmp10097)

					tmp10099 := PrimTail(V963)

					tmp10100 := PrimTail(tmp10099)

					tmp10101 := PrimHead(tmp10100)

					tmp10102 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10101, V964)

					tmp10103 := PrimTail(V963)

					tmp10104 := PrimTail(tmp10103)

					tmp10105 := PrimTail(tmp10104)

					tmp10106 := PrimHead(tmp10105)

					tmp10107 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10106, V964)

					tmp10108 := PrimCons(tmp10107, Nil)

					tmp10109 := PrimCons(tmp10102, tmp10108)

					tmp10110 := PrimCons(tmp10098, tmp10109)

					__e.Return(PrimCons(symlet, tmp10110))
					return

				} else {
					tmp10096 := PrimIsPair(V963)

					if True == tmp10096 {
						tmp10092 := PrimHead(V963)

						tmp10093 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10092, V964)

						tmp10094 := PrimTail(V963)

						tmp10095 := Call(__e, PrimNS2Value(symshen_4insert_1deref), tmp10094, V964)

						__e.Return(PrimCons(tmp10093, tmp10095))
						return

					} else {
						__e.Return(V963)
						return
					}

				}

			}

		}

	}, 2)

	tmp10169 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1deref, tmp10087)

	_ = tmp10169

	tmp10170 := MakeNative(func(__e *ControlFlow) {
		V971 := __e.Get(1)
		_ = V971
		V972 := __e.Get(2)
		_ = V972
		tmp10251 := PrimIsVariable(V971)

		if True == tmp10251 {
			tmp10249 := PrimCons(V972, Nil)

			tmp10250 := PrimCons(V971, tmp10249)

			__e.Return(PrimCons(symshen_4lazyderef, tmp10250))
			return

		} else {
			tmp10248 := PrimIsPair(V971)

			var ifres10229 Obj

			if True == tmp10248 {
				tmp10246 := PrimHead(V971)

				tmp10247 := PrimEqual(symlambda, tmp10246)

				var ifres10231 Obj

				if True == tmp10247 {
					tmp10244 := PrimTail(V971)

					tmp10245 := PrimIsPair(tmp10244)

					var ifres10233 Obj

					if True == tmp10245 {
						tmp10241 := PrimTail(V971)

						tmp10242 := PrimTail(tmp10241)

						tmp10243 := PrimIsPair(tmp10242)

						var ifres10235 Obj

						if True == tmp10243 {
							tmp10237 := PrimTail(V971)

							tmp10238 := PrimTail(tmp10237)

							tmp10239 := PrimTail(tmp10238)

							tmp10240 := PrimEqual(Nil, tmp10239)

							var ifres10236 Obj

							if True == tmp10240 {
								ifres10236 = True

							} else {
								ifres10236 = False

							}

							ifres10235 = ifres10236

						} else {
							ifres10235 = False

						}

						var ifres10234 Obj

						if True == ifres10235 {
							ifres10234 = True

						} else {
							ifres10234 = False

						}

						ifres10233 = ifres10234

					} else {
						ifres10233 = False

					}

					var ifres10232 Obj

					if True == ifres10233 {
						ifres10232 = True

					} else {
						ifres10232 = False

					}

					ifres10231 = ifres10232

				} else {
					ifres10231 = False

				}

				var ifres10230 Obj

				if True == ifres10231 {
					ifres10230 = True

				} else {
					ifres10230 = False

				}

				ifres10229 = ifres10230

			} else {
				ifres10229 = False

			}

			if True == ifres10229 {
				tmp10221 := PrimTail(V971)

				tmp10222 := PrimHead(tmp10221)

				tmp10223 := PrimTail(V971)

				tmp10224 := PrimTail(tmp10223)

				tmp10225 := PrimHead(tmp10224)

				tmp10226 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10225, V972)

				tmp10227 := PrimCons(tmp10226, Nil)

				tmp10228 := PrimCons(tmp10222, tmp10227)

				__e.Return(PrimCons(symlambda, tmp10228))
				return

			} else {
				tmp10220 := PrimIsPair(V971)

				var ifres10194 Obj

				if True == tmp10220 {
					tmp10218 := PrimHead(V971)

					tmp10219 := PrimEqual(symlet, tmp10218)

					var ifres10196 Obj

					if True == tmp10219 {
						tmp10216 := PrimTail(V971)

						tmp10217 := PrimIsPair(tmp10216)

						var ifres10198 Obj

						if True == tmp10217 {
							tmp10213 := PrimTail(V971)

							tmp10214 := PrimTail(tmp10213)

							tmp10215 := PrimIsPair(tmp10214)

							var ifres10200 Obj

							if True == tmp10215 {
								tmp10209 := PrimTail(V971)

								tmp10210 := PrimTail(tmp10209)

								tmp10211 := PrimTail(tmp10210)

								tmp10212 := PrimIsPair(tmp10211)

								var ifres10202 Obj

								if True == tmp10212 {
									tmp10204 := PrimTail(V971)

									tmp10205 := PrimTail(tmp10204)

									tmp10206 := PrimTail(tmp10205)

									tmp10207 := PrimTail(tmp10206)

									tmp10208 := PrimEqual(Nil, tmp10207)

									var ifres10203 Obj

									if True == tmp10208 {
										ifres10203 = True

									} else {
										ifres10203 = False

									}

									ifres10202 = ifres10203

								} else {
									ifres10202 = False

								}

								var ifres10201 Obj

								if True == ifres10202 {
									ifres10201 = True

								} else {
									ifres10201 = False

								}

								ifres10200 = ifres10201

							} else {
								ifres10200 = False

							}

							var ifres10199 Obj

							if True == ifres10200 {
								ifres10199 = True

							} else {
								ifres10199 = False

							}

							ifres10198 = ifres10199

						} else {
							ifres10198 = False

						}

						var ifres10197 Obj

						if True == ifres10198 {
							ifres10197 = True

						} else {
							ifres10197 = False

						}

						ifres10196 = ifres10197

					} else {
						ifres10196 = False

					}

					var ifres10195 Obj

					if True == ifres10196 {
						ifres10195 = True

					} else {
						ifres10195 = False

					}

					ifres10194 = ifres10195

				} else {
					ifres10194 = False

				}

				if True == ifres10194 {
					tmp10180 := PrimTail(V971)

					tmp10181 := PrimHead(tmp10180)

					tmp10182 := PrimTail(V971)

					tmp10183 := PrimTail(tmp10182)

					tmp10184 := PrimHead(tmp10183)

					tmp10185 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10184, V972)

					tmp10186 := PrimTail(V971)

					tmp10187 := PrimTail(tmp10186)

					tmp10188 := PrimTail(tmp10187)

					tmp10189 := PrimHead(tmp10188)

					tmp10190 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10189, V972)

					tmp10191 := PrimCons(tmp10190, Nil)

					tmp10192 := PrimCons(tmp10185, tmp10191)

					tmp10193 := PrimCons(tmp10181, tmp10192)

					__e.Return(PrimCons(symlet, tmp10193))
					return

				} else {
					tmp10179 := PrimIsPair(V971)

					if True == tmp10179 {
						tmp10175 := PrimHead(V971)

						tmp10176 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10175, V972)

						tmp10177 := PrimTail(V971)

						tmp10178 := Call(__e, PrimNS2Value(symshen_4insert_1lazyderef), tmp10177, V972)

						__e.Return(PrimCons(tmp10176, tmp10178))
						return

					} else {
						__e.Return(V971)
						return
					}

				}

			}

		}

	}, 2)

	tmp10252 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1lazyderef, tmp10170)

	_ = tmp10252

	tmp10253 := MakeNative(func(__e *ControlFlow) {
		V974 := __e.Get(1)
		_ = V974
		tmp10264 := PrimEqual(Nil, V974)

		if True == tmp10264 {
			__e.Return(Nil)
			return
		} else {
			tmp10263 := PrimIsPair(V974)

			if True == tmp10263 {
				tmp10256 := MakeNative(func(__e *ControlFlow) {
					Group := __e.Get(1)
					_ = Group
					tmp10257 := MakeNative(func(__e *ControlFlow) {
						Rest := __e.Get(1)
						_ = Rest
						tmp10258 := Call(__e, PrimNS2Value(symshen_4group__clauses), Rest)

						__e.Return(PrimCons(Group, tmp10258))
						return

					}, 1)

					tmp10259 := Call(__e, PrimNS2Value(symdifference), V974, Group)

					__e.TailApply(tmp10257, tmp10259)
					return

				}, 1)

				tmp10260 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp10261 := PrimHead(V974)

					__e.TailApply(PrimNS2Value(symshen_4same__predicate_2), tmp10261, X)
					return

				}, 1)

				tmp10262 := Call(__e, PrimNS2Value(symshen_4collect), tmp10260, V974)

				__e.TailApply(tmp10256, tmp10262)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4group__clauses)
				return
			}

		}

	}, 1)

	tmp10265 := Call(__e, PrimNS1Value(symns2_1set), symshen_4group__clauses, tmp10253)

	_ = tmp10265

	tmp10266 := MakeNative(func(__e *ControlFlow) {
		V979 := __e.Get(1)
		_ = V979
		V980 := __e.Get(2)
		_ = V980
		tmp10277 := PrimEqual(Nil, V980)

		if True == tmp10277 {
			__e.Return(Nil)
			return
		} else {
			tmp10276 := PrimIsPair(V980)

			if True == tmp10276 {
				tmp10274 := PrimHead(V980)

				tmp10275 := Call(__e, V979, tmp10274)

				if True == tmp10275 {
					tmp10271 := PrimHead(V980)

					tmp10272 := PrimTail(V980)

					tmp10273 := Call(__e, PrimNS2Value(symshen_4collect), V979, tmp10272)

					__e.Return(PrimCons(tmp10271, tmp10273))
					return

				} else {
					tmp10270 := PrimTail(V980)

					__e.TailApply(PrimNS2Value(symshen_4collect), V979, tmp10270)
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4collect)
				return
			}

		}

	}, 2)

	tmp10278 := Call(__e, PrimNS1Value(symns2_1set), symshen_4collect, tmp10266)

	_ = tmp10278

	tmp10279 := MakeNative(func(__e *ControlFlow) {
		V999 := __e.Get(1)
		_ = V999
		V1000 := __e.Get(2)
		_ = V1000
		tmp10296 := PrimIsPair(V999)

		var ifres10285 Obj

		if True == tmp10296 {
			tmp10294 := PrimHead(V999)

			tmp10295 := PrimIsPair(tmp10294)

			var ifres10287 Obj

			if True == tmp10295 {
				tmp10293 := PrimIsPair(V1000)

				var ifres10289 Obj

				if True == tmp10293 {
					tmp10291 := PrimHead(V1000)

					tmp10292 := PrimIsPair(tmp10291)

					var ifres10290 Obj

					if True == tmp10292 {
						ifres10290 = True

					} else {
						ifres10290 = False

					}

					ifres10289 = ifres10290

				} else {
					ifres10289 = False

				}

				var ifres10288 Obj

				if True == ifres10289 {
					ifres10288 = True

				} else {
					ifres10288 = False

				}

				ifres10287 = ifres10288

			} else {
				ifres10287 = False

			}

			var ifres10286 Obj

			if True == ifres10287 {
				ifres10286 = True

			} else {
				ifres10286 = False

			}

			ifres10285 = ifres10286

		} else {
			ifres10285 = False

		}

		if True == ifres10285 {
			tmp10281 := PrimHead(V999)

			tmp10282 := PrimHead(tmp10281)

			tmp10283 := PrimHead(V1000)

			tmp10284 := PrimHead(tmp10283)

			__e.Return(PrimEqual(tmp10282, tmp10284))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4same__predicate_2)
			return
		}

	}, 2)

	tmp10297 := Call(__e, PrimNS1Value(symns2_1set), symshen_4same__predicate_2, tmp10279)

	_ = tmp10297

	tmp10298 := MakeNative(func(__e *ControlFlow) {
		V1002 := __e.Get(1)
		_ = V1002
		tmp10299 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp10300 := MakeNative(func(__e *ControlFlow) {
				Shen := __e.Get(1)
				_ = Shen
				__e.Return(Shen)
				return
			}, 1)

			tmp10301 := Call(__e, PrimNS2Value(symshen_4clauses_1to_1shen), F, V1002)

			__e.TailApply(tmp10300, tmp10301)
			return

		}, 1)

		tmp10302 := Call(__e, PrimNS2Value(symshen_4procedure__name), V1002)

		__e.TailApply(tmp10299, tmp10302)
		return

	}, 1)

	tmp10303 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile__prolog__procedure, tmp10298)

	_ = tmp10303

	tmp10304 := MakeNative(func(__e *ControlFlow) {
		V1016 := __e.Get(1)
		_ = V1016
		tmp10317 := PrimIsPair(V1016)

		var ifres10308 Obj

		if True == tmp10317 {
			tmp10315 := PrimHead(V1016)

			tmp10316 := PrimIsPair(tmp10315)

			var ifres10310 Obj

			if True == tmp10316 {
				tmp10312 := PrimHead(V1016)

				tmp10313 := PrimHead(tmp10312)

				tmp10314 := PrimIsPair(tmp10313)

				var ifres10311 Obj

				if True == tmp10314 {
					ifres10311 = True

				} else {
					ifres10311 = False

				}

				ifres10310 = ifres10311

			} else {
				ifres10310 = False

			}

			var ifres10309 Obj

			if True == ifres10310 {
				ifres10309 = True

			} else {
				ifres10309 = False

			}

			ifres10308 = ifres10309

		} else {
			ifres10308 = False

		}

		if True == ifres10308 {
			tmp10306 := PrimHead(V1016)

			tmp10307 := PrimHead(tmp10306)

			__e.Return(PrimHead(tmp10307))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4procedure__name)
			return
		}

	}, 1)

	tmp10318 := Call(__e, PrimNS1Value(symns2_1set), symshen_4procedure__name, tmp10304)

	_ = tmp10318

	tmp10319 := MakeNative(func(__e *ControlFlow) {
		V1019 := __e.Get(1)
		_ = V1019
		V1020 := __e.Get(2)
		_ = V1020
		tmp10320 := MakeNative(func(__e *ControlFlow) {
			Linear := __e.Get(1)
			_ = Linear
			tmp10321 := MakeNative(func(__e *ControlFlow) {
				Arity := __e.Get(1)
				_ = Arity
				tmp10322 := MakeNative(func(__e *ControlFlow) {
					Parameters := __e.Get(1)
					_ = Parameters
					tmp10323 := MakeNative(func(__e *ControlFlow) {
						AUM__instructions := __e.Get(1)
						_ = AUM__instructions
						tmp10324 := MakeNative(func(__e *ControlFlow) {
							Code := __e.Get(1)
							_ = Code
							tmp10325 := MakeNative(func(__e *ControlFlow) {
								ShenDef := __e.Get(1)
								_ = ShenDef
								__e.Return(ShenDef)
								return
							}, 1)

							tmp10326 := PrimCons(symContinuation, Nil)

							tmp10327 := PrimCons(symProcessN, tmp10326)

							tmp10328 := PrimCons(Code, Nil)

							tmp10329 := PrimCons(sym_1_6, tmp10328)

							tmp10330 := Call(__e, PrimNS2Value(symappend), tmp10327, tmp10329)

							tmp10331 := Call(__e, PrimNS2Value(symappend), Parameters, tmp10330)

							tmp10332 := PrimCons(V1019, tmp10331)

							tmp10333 := PrimCons(symdefine, tmp10332)

							__e.TailApply(tmp10325, tmp10333)
							return

						}, 1)

						tmp10334 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.TailApply(PrimNS2Value(symshen_4aum__to__shen), X)
							return
						}, 1)

						tmp10335 := Call(__e, PrimNS2Value(symmap), tmp10334, AUM__instructions)

						tmp10336 := Call(__e, PrimNS2Value(symshen_4nest_1disjunct), tmp10335)

						tmp10337 := Call(__e, PrimNS2Value(symshen_4catch_1cut), tmp10336)

						__e.TailApply(tmp10324, tmp10337)
						return

					}, 1)

					tmp10338 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4aum), X, Parameters)
						return
					}, 1)

					tmp10339 := Call(__e, PrimNS2Value(symmap), tmp10338, Linear)

					__e.TailApply(tmp10323, tmp10339)
					return

				}, 1)

				tmp10340 := Call(__e, PrimNS2Value(symshen_4parameters), Arity)

				__e.TailApply(tmp10322, tmp10340)
				return

			}, 1)

			tmp10341 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symhead), X)
				return
			}, 1)

			tmp10342 := Call(__e, PrimNS2Value(symmap), tmp10341, V1020)

			tmp10343 := Call(__e, PrimNS2Value(symshen_4prolog_1aritycheck), V1019, tmp10342)

			__e.TailApply(tmp10321, tmp10343)
			return

		}, 1)

		tmp10344 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4linearise_1clause), X)
			return
		}, 1)

		tmp10345 := Call(__e, PrimNS2Value(symmap), tmp10344, V1020)

		__e.TailApply(tmp10320, tmp10345)
		return

	}, 2)

	tmp10346 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clauses_1to_1shen, tmp10319)

	_ = tmp10346

	tmp10347 := MakeNative(func(__e *ControlFlow) {
		V1022 := __e.Get(1)
		_ = V1022
		tmp10356 := Call(__e, PrimNS2Value(symshen_4occurs_2), symcut, V1022)

		tmp10357 := PrimNot(tmp10356)

		if True == tmp10357 {
			__e.Return(V1022)
			return
		} else {
			tmp10349 := PrimCons(symshen_4catchpoint, Nil)

			tmp10350 := PrimCons(V1022, Nil)

			tmp10351 := PrimCons(symThrowcontrol, tmp10350)

			tmp10352 := PrimCons(symshen_4cutpoint, tmp10351)

			tmp10353 := PrimCons(tmp10352, Nil)

			tmp10354 := PrimCons(tmp10349, tmp10353)

			tmp10355 := PrimCons(symThrowcontrol, tmp10354)

			__e.Return(PrimCons(symlet, tmp10355))
			return

		}

	}, 1)

	tmp10358 := Call(__e, PrimNS1Value(symns2_1set), symshen_4catch_1cut, tmp10347)

	_ = tmp10358

	tmp10359 := MakeNative(func(__e *ControlFlow) {
		tmp10360 := PrimNS3Value(symshen_4_dcatch_d)

		tmp10361 := PrimNumberAdd(MakeNumber(1), tmp10360)

		tmp10362 := PrimNS3Set(symshen_4_dcatch_d, tmp10361)

		__e.Return(PrimCons(symshen_4catchpoint_b, tmp10362))
		return

	}, 0)

	tmp10363 := Call(__e, PrimNS1Value(symns2_1set), symshen_4catchpoint, tmp10359)

	_ = tmp10363

	tmp10364 := MakeNative(func(__e *ControlFlow) {
		V1030 := __e.Get(1)
		_ = V1030
		V1031 := __e.Get(2)
		_ = V1031
		tmp10366 := PrimEqual(V1031, V1030)

		if True == tmp10366 {
			__e.Return(False)
			return
		} else {
			__e.Return(V1031)
			return
		}

	}, 2)

	tmp10367 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cutpoint, tmp10364)

	_ = tmp10367

	tmp10368 := MakeNative(func(__e *ControlFlow) {
		V1033 := __e.Get(1)
		_ = V1033
		tmp10379 := PrimIsPair(V1033)

		var ifres10375 Obj

		if True == tmp10379 {
			tmp10377 := PrimTail(V1033)

			tmp10378 := PrimEqual(Nil, tmp10377)

			var ifres10376 Obj

			if True == tmp10378 {
				ifres10376 = True

			} else {
				ifres10376 = False

			}

			ifres10375 = ifres10376

		} else {
			ifres10375 = False

		}

		if True == ifres10375 {
			__e.Return(PrimHead(V1033))
			return
		} else {
			tmp10374 := PrimIsPair(V1033)

			if True == tmp10374 {
				tmp10371 := PrimHead(V1033)

				tmp10372 := PrimTail(V1033)

				tmp10373 := Call(__e, PrimNS2Value(symshen_4nest_1disjunct), tmp10372)

				__e.TailApply(PrimNS2Value(symshen_4lisp_1or), tmp10371, tmp10373)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4nest_1disjunct)
				return
			}

		}

	}, 1)

	tmp10380 := Call(__e, PrimNS1Value(symns2_1set), symshen_4nest_1disjunct, tmp10368)

	_ = tmp10380

	tmp10381 := MakeNative(func(__e *ControlFlow) {
		V1036 := __e.Get(1)
		_ = V1036
		V1037 := __e.Get(2)
		_ = V1037
		tmp10382 := PrimCons(False, Nil)

		tmp10383 := PrimCons(symCase, tmp10382)

		tmp10384 := PrimCons(sym_a, tmp10383)

		tmp10385 := PrimCons(symCase, Nil)

		tmp10386 := PrimCons(V1037, tmp10385)

		tmp10387 := PrimCons(tmp10384, tmp10386)

		tmp10388 := PrimCons(symif, tmp10387)

		tmp10389 := PrimCons(tmp10388, Nil)

		tmp10390 := PrimCons(V1036, tmp10389)

		tmp10391 := PrimCons(symCase, tmp10390)

		__e.Return(PrimCons(symlet, tmp10391))
		return

	}, 2)

	tmp10392 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lisp_1or, tmp10381)

	_ = tmp10392

	tmp10393 := MakeNative(func(__e *ControlFlow) {
		V1042 := __e.Get(1)
		_ = V1042
		V1043 := __e.Get(2)
		_ = V1043
		tmp10418 := PrimIsPair(V1043)

		var ifres10414 Obj

		if True == tmp10418 {
			tmp10416 := PrimTail(V1043)

			tmp10417 := PrimEqual(Nil, tmp10416)

			var ifres10415 Obj

			if True == tmp10417 {
				ifres10415 = True

			} else {
				ifres10415 = False

			}

			ifres10414 = ifres10415

		} else {
			ifres10414 = False

		}

		if True == ifres10414 {
			tmp10412 := PrimHead(V1043)

			tmp10413 := Call(__e, PrimNS2Value(symlength), tmp10412)

			__e.Return(PrimNumberSubtract(tmp10413, MakeNumber(1)))
			return

		} else {
			tmp10411 := PrimIsPair(V1043)

			var ifres10407 Obj

			if True == tmp10411 {
				tmp10409 := PrimTail(V1043)

				tmp10410 := PrimIsPair(tmp10409)

				var ifres10408 Obj

				if True == tmp10410 {
					ifres10408 = True

				} else {
					ifres10408 = False

				}

				ifres10407 = ifres10408

			} else {
				ifres10407 = False

			}

			if True == ifres10407 {
				tmp10401 := PrimHead(V1043)

				tmp10402 := Call(__e, PrimNS2Value(symlength), tmp10401)

				tmp10403 := PrimTail(V1043)

				tmp10404 := PrimHead(tmp10403)

				tmp10405 := Call(__e, PrimNS2Value(symlength), tmp10404)

				tmp10406 := PrimEqual(tmp10402, tmp10405)

				if True == tmp10406 {
					tmp10400 := PrimTail(V1043)

					__e.TailApply(PrimNS2Value(symshen_4prolog_1aritycheck), V1042, tmp10400)
					return

				} else {
					tmp10397 := PrimCons(V1042, Nil)

					tmp10398 := Call(__e, PrimNS2Value(symshen_4app), tmp10397, MakeString("\n"), symshen_4a)

					tmp10399 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp10398)

					__e.Return(PrimSimpleError(tmp10399))
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4prolog_1aritycheck)
				return
			}

		}

	}, 2)

	tmp10419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1aritycheck, tmp10393)

	_ = tmp10419

	tmp10420 := MakeNative(func(__e *ControlFlow) {
		V1045 := __e.Get(1)
		_ = V1045
		tmp10448 := PrimIsPair(V1045)

		var ifres10428 Obj

		if True == tmp10448 {
			tmp10446 := PrimTail(V1045)

			tmp10447 := PrimIsPair(tmp10446)

			var ifres10430 Obj

			if True == tmp10447 {
				tmp10443 := PrimTail(V1045)

				tmp10444 := PrimHead(tmp10443)

				tmp10445 := PrimEqual(sym_h_1, tmp10444)

				var ifres10432 Obj

				if True == tmp10445 {
					tmp10440 := PrimTail(V1045)

					tmp10441 := PrimTail(tmp10440)

					tmp10442 := PrimIsPair(tmp10441)

					var ifres10434 Obj

					if True == tmp10442 {
						tmp10436 := PrimTail(V1045)

						tmp10437 := PrimTail(tmp10436)

						tmp10438 := PrimTail(tmp10437)

						tmp10439 := PrimEqual(Nil, tmp10438)

						var ifres10435 Obj

						if True == tmp10439 {
							ifres10435 = True

						} else {
							ifres10435 = False

						}

						ifres10434 = ifres10435

					} else {
						ifres10434 = False

					}

					var ifres10433 Obj

					if True == ifres10434 {
						ifres10433 = True

					} else {
						ifres10433 = False

					}

					ifres10432 = ifres10433

				} else {
					ifres10432 = False

				}

				var ifres10431 Obj

				if True == ifres10432 {
					ifres10431 = True

				} else {
					ifres10431 = False

				}

				ifres10430 = ifres10431

			} else {
				ifres10430 = False

			}

			var ifres10429 Obj

			if True == ifres10430 {
				ifres10429 = True

			} else {
				ifres10429 = False

			}

			ifres10428 = ifres10429

		} else {
			ifres10428 = False

		}

		if True == ifres10428 {
			tmp10422 := MakeNative(func(__e *ControlFlow) {
				Linear := __e.Get(1)
				_ = Linear
				__e.TailApply(PrimNS2Value(symshen_4clause__form), Linear)
				return
			}, 1)

			tmp10423 := PrimHead(V1045)

			tmp10424 := PrimTail(V1045)

			tmp10425 := PrimTail(tmp10424)

			tmp10426 := PrimCons(tmp10423, tmp10425)

			tmp10427 := Call(__e, PrimNS2Value(symshen_4linearise), tmp10426)

			__e.TailApply(tmp10422, tmp10427)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4linearise_1clause)
			return
		}

	}, 1)

	tmp10449 := Call(__e, PrimNS1Value(symns2_1set), symshen_4linearise_1clause, tmp10420)

	_ = tmp10449

	tmp10450 := MakeNative(func(__e *ControlFlow) {
		V1047 := __e.Get(1)
		_ = V1047
		tmp10468 := PrimIsPair(V1047)

		var ifres10459 Obj

		if True == tmp10468 {
			tmp10466 := PrimTail(V1047)

			tmp10467 := PrimIsPair(tmp10466)

			var ifres10461 Obj

			if True == tmp10467 {
				tmp10463 := PrimTail(V1047)

				tmp10464 := PrimTail(tmp10463)

				tmp10465 := PrimEqual(Nil, tmp10464)

				var ifres10462 Obj

				if True == tmp10465 {
					ifres10462 = True

				} else {
					ifres10462 = False

				}

				ifres10461 = ifres10462

			} else {
				ifres10461 = False

			}

			var ifres10460 Obj

			if True == ifres10461 {
				ifres10460 = True

			} else {
				ifres10460 = False

			}

			ifres10459 = ifres10460

		} else {
			ifres10459 = False

		}

		if True == ifres10459 {
			tmp10452 := PrimHead(V1047)

			tmp10453 := Call(__e, PrimNS2Value(symshen_4explicit__modes), tmp10452)

			tmp10454 := PrimTail(V1047)

			tmp10455 := PrimHead(tmp10454)

			tmp10456 := Call(__e, PrimNS2Value(symshen_4cf__help), tmp10455)

			tmp10457 := PrimCons(tmp10456, Nil)

			tmp10458 := PrimCons(sym_h_1, tmp10457)

			__e.Return(PrimCons(tmp10453, tmp10458))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4clause__form)
			return
		}

	}, 1)

	tmp10469 := Call(__e, PrimNS1Value(symns2_1set), symshen_4clause__form, tmp10450)

	_ = tmp10469

	tmp10470 := MakeNative(func(__e *ControlFlow) {
		V1049 := __e.Get(1)
		_ = V1049
		tmp10476 := PrimIsPair(V1049)

		if True == tmp10476 {
			tmp10472 := PrimHead(V1049)

			tmp10473 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4em__help), X)
				return
			}, 1)

			tmp10474 := PrimTail(V1049)

			tmp10475 := Call(__e, PrimNS2Value(symmap), tmp10473, tmp10474)

			__e.Return(PrimCons(tmp10472, tmp10475))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4explicit__modes)
			return
		}

	}, 1)

	tmp10477 := Call(__e, PrimNS1Value(symns2_1set), symshen_4explicit__modes, tmp10470)

	_ = tmp10477

	tmp10478 := MakeNative(func(__e *ControlFlow) {
		V1051 := __e.Get(1)
		_ = V1051
		tmp10501 := PrimIsPair(V1051)

		var ifres10482 Obj

		if True == tmp10501 {
			tmp10499 := PrimHead(V1051)

			tmp10500 := PrimEqual(symmode, tmp10499)

			var ifres10484 Obj

			if True == tmp10500 {
				tmp10497 := PrimTail(V1051)

				tmp10498 := PrimIsPair(tmp10497)

				var ifres10486 Obj

				if True == tmp10498 {
					tmp10494 := PrimTail(V1051)

					tmp10495 := PrimTail(tmp10494)

					tmp10496 := PrimIsPair(tmp10495)

					var ifres10488 Obj

					if True == tmp10496 {
						tmp10490 := PrimTail(V1051)

						tmp10491 := PrimTail(tmp10490)

						tmp10492 := PrimTail(tmp10491)

						tmp10493 := PrimEqual(Nil, tmp10492)

						var ifres10489 Obj

						if True == tmp10493 {
							ifres10489 = True

						} else {
							ifres10489 = False

						}

						ifres10488 = ifres10489

					} else {
						ifres10488 = False

					}

					var ifres10487 Obj

					if True == ifres10488 {
						ifres10487 = True

					} else {
						ifres10487 = False

					}

					ifres10486 = ifres10487

				} else {
					ifres10486 = False

				}

				var ifres10485 Obj

				if True == ifres10486 {
					ifres10485 = True

				} else {
					ifres10485 = False

				}

				ifres10484 = ifres10485

			} else {
				ifres10484 = False

			}

			var ifres10483 Obj

			if True == ifres10484 {
				ifres10483 = True

			} else {
				ifres10483 = False

			}

			ifres10482 = ifres10483

		} else {
			ifres10482 = False

		}

		if True == ifres10482 {
			__e.Return(V1051)
			return
		} else {
			tmp10480 := PrimCons(sym_7, Nil)

			tmp10481 := PrimCons(V1051, tmp10480)

			__e.Return(PrimCons(symmode, tmp10481))
			return

		}

	}, 1)

	tmp10502 := Call(__e, PrimNS1Value(symns2_1set), symshen_4em__help, tmp10478)

	_ = tmp10502

	tmp10503 := MakeNative(func(__e *ControlFlow) {
		V1053 := __e.Get(1)
		_ = V1053
		tmp10566 := PrimIsPair(V1053)

		var ifres10515 Obj

		if True == tmp10566 {
			tmp10564 := PrimHead(V1053)

			tmp10565 := PrimEqual(symwhere, tmp10564)

			var ifres10517 Obj

			if True == tmp10565 {
				tmp10562 := PrimTail(V1053)

				tmp10563 := PrimIsPair(tmp10562)

				var ifres10519 Obj

				if True == tmp10563 {
					tmp10559 := PrimTail(V1053)

					tmp10560 := PrimHead(tmp10559)

					tmp10561 := PrimIsPair(tmp10560)

					var ifres10521 Obj

					if True == tmp10561 {
						tmp10555 := PrimTail(V1053)

						tmp10556 := PrimHead(tmp10555)

						tmp10557 := PrimHead(tmp10556)

						tmp10558 := PrimEqual(sym_a, tmp10557)

						var ifres10523 Obj

						if True == tmp10558 {
							tmp10551 := PrimTail(V1053)

							tmp10552 := PrimHead(tmp10551)

							tmp10553 := PrimTail(tmp10552)

							tmp10554 := PrimIsPair(tmp10553)

							var ifres10525 Obj

							if True == tmp10554 {
								tmp10546 := PrimTail(V1053)

								tmp10547 := PrimHead(tmp10546)

								tmp10548 := PrimTail(tmp10547)

								tmp10549 := PrimTail(tmp10548)

								tmp10550 := PrimIsPair(tmp10549)

								var ifres10527 Obj

								if True == tmp10550 {
									tmp10540 := PrimTail(V1053)

									tmp10541 := PrimHead(tmp10540)

									tmp10542 := PrimTail(tmp10541)

									tmp10543 := PrimTail(tmp10542)

									tmp10544 := PrimTail(tmp10543)

									tmp10545 := PrimEqual(Nil, tmp10544)

									var ifres10529 Obj

									if True == tmp10545 {
										tmp10537 := PrimTail(V1053)

										tmp10538 := PrimTail(tmp10537)

										tmp10539 := PrimIsPair(tmp10538)

										var ifres10531 Obj

										if True == tmp10539 {
											tmp10533 := PrimTail(V1053)

											tmp10534 := PrimTail(tmp10533)

											tmp10535 := PrimTail(tmp10534)

											tmp10536 := PrimEqual(Nil, tmp10535)

											var ifres10532 Obj

											if True == tmp10536 {
												ifres10532 = True

											} else {
												ifres10532 = False

											}

											ifres10531 = ifres10532

										} else {
											ifres10531 = False

										}

										var ifres10530 Obj

										if True == ifres10531 {
											ifres10530 = True

										} else {
											ifres10530 = False

										}

										ifres10529 = ifres10530

									} else {
										ifres10529 = False

									}

									var ifres10528 Obj

									if True == ifres10529 {
										ifres10528 = True

									} else {
										ifres10528 = False

									}

									ifres10527 = ifres10528

								} else {
									ifres10527 = False

								}

								var ifres10526 Obj

								if True == ifres10527 {
									ifres10526 = True

								} else {
									ifres10526 = False

								}

								ifres10525 = ifres10526

							} else {
								ifres10525 = False

							}

							var ifres10524 Obj

							if True == ifres10525 {
								ifres10524 = True

							} else {
								ifres10524 = False

							}

							ifres10523 = ifres10524

						} else {
							ifres10523 = False

						}

						var ifres10522 Obj

						if True == ifres10523 {
							ifres10522 = True

						} else {
							ifres10522 = False

						}

						ifres10521 = ifres10522

					} else {
						ifres10521 = False

					}

					var ifres10520 Obj

					if True == ifres10521 {
						ifres10520 = True

					} else {
						ifres10520 = False

					}

					ifres10519 = ifres10520

				} else {
					ifres10519 = False

				}

				var ifres10518 Obj

				if True == ifres10519 {
					ifres10518 = True

				} else {
					ifres10518 = False

				}

				ifres10517 = ifres10518

			} else {
				ifres10517 = False

			}

			var ifres10516 Obj

			if True == ifres10517 {
				ifres10516 = True

			} else {
				ifres10516 = False

			}

			ifres10515 = ifres10516

		} else {
			ifres10515 = False

		}

		if True == ifres10515 {
			tmp10506 := PrimNS3Value(symshen_4_doccurs_d)

			var ifres10505 Obj

			if True == tmp10506 {
				ifres10505 = symunify_b

			} else {
				ifres10505 = symunify

			}

			tmp10507 := PrimTail(V1053)

			tmp10508 := PrimHead(tmp10507)

			tmp10509 := PrimTail(tmp10508)

			tmp10510 := PrimCons(ifres10505, tmp10509)

			tmp10511 := PrimTail(V1053)

			tmp10512 := PrimTail(tmp10511)

			tmp10513 := PrimHead(tmp10512)

			tmp10514 := Call(__e, PrimNS2Value(symshen_4cf__help), tmp10513)

			__e.Return(PrimCons(tmp10510, tmp10514))
			return

		} else {
			__e.Return(V1053)
			return
		}

	}, 1)

	tmp10567 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cf__help, tmp10503)

	_ = tmp10567

	tmp10568 := MakeNative(func(__e *ControlFlow) {
		V1059 := __e.Get(1)
		_ = V1059
		tmp10572 := PrimEqual(sym_7, V1059)

		if True == tmp10572 {
			__e.Return(PrimNS3Set(symshen_4_doccurs_d, True))
			return
		} else {
			tmp10571 := PrimEqual(sym_1, V1059)

			if True == tmp10571 {
				__e.Return(PrimNS3Set(symshen_4_doccurs_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("occurs-check expects + or -\n")))
				return
			}

		}

	}, 1)

	tmp10573 := Call(__e, PrimNS1Value(symns2_1set), symoccurs_1check, tmp10568)

	_ = tmp10573

	tmp10574 := MakeNative(func(__e *ControlFlow) {
		V1062 := __e.Get(1)
		_ = V1062
		V1063 := __e.Get(2)
		_ = V1063
		tmp10613 := PrimIsPair(V1062)

		var ifres10589 Obj

		if True == tmp10613 {
			tmp10611 := PrimHead(V1062)

			tmp10612 := PrimIsPair(tmp10611)

			var ifres10591 Obj

			if True == tmp10612 {
				tmp10609 := PrimTail(V1062)

				tmp10610 := PrimIsPair(tmp10609)

				var ifres10593 Obj

				if True == tmp10610 {
					tmp10606 := PrimTail(V1062)

					tmp10607 := PrimHead(tmp10606)

					tmp10608 := PrimEqual(sym_h_1, tmp10607)

					var ifres10595 Obj

					if True == tmp10608 {
						tmp10603 := PrimTail(V1062)

						tmp10604 := PrimTail(tmp10603)

						tmp10605 := PrimIsPair(tmp10604)

						var ifres10597 Obj

						if True == tmp10605 {
							tmp10599 := PrimTail(V1062)

							tmp10600 := PrimTail(tmp10599)

							tmp10601 := PrimTail(tmp10600)

							tmp10602 := PrimEqual(Nil, tmp10601)

							var ifres10598 Obj

							if True == tmp10602 {
								ifres10598 = True

							} else {
								ifres10598 = False

							}

							ifres10597 = ifres10598

						} else {
							ifres10597 = False

						}

						var ifres10596 Obj

						if True == ifres10597 {
							ifres10596 = True

						} else {
							ifres10596 = False

						}

						ifres10595 = ifres10596

					} else {
						ifres10595 = False

					}

					var ifres10594 Obj

					if True == ifres10595 {
						ifres10594 = True

					} else {
						ifres10594 = False

					}

					ifres10593 = ifres10594

				} else {
					ifres10593 = False

				}

				var ifres10592 Obj

				if True == ifres10593 {
					ifres10592 = True

				} else {
					ifres10592 = False

				}

				ifres10591 = ifres10592

			} else {
				ifres10591 = False

			}

			var ifres10590 Obj

			if True == ifres10591 {
				ifres10590 = True

			} else {
				ifres10590 = False

			}

			ifres10589 = ifres10590

		} else {
			ifres10589 = False

		}

		if True == ifres10589 {
			tmp10576 := MakeNative(func(__e *ControlFlow) {
				MuApplication := __e.Get(1)
				_ = MuApplication
				__e.TailApply(PrimNS2Value(symshen_4mu__reduction), MuApplication, sym_7)
				return
			}, 1)

			tmp10577 := PrimHead(V1062)

			tmp10578 := PrimTail(tmp10577)

			tmp10579 := PrimHead(V1062)

			tmp10580 := PrimTail(tmp10579)

			tmp10581 := PrimTail(V1062)

			tmp10582 := PrimTail(tmp10581)

			tmp10583 := PrimHead(tmp10582)

			tmp10584 := Call(__e, PrimNS2Value(symshen_4continuation__call), tmp10580, tmp10583)

			tmp10585 := PrimCons(tmp10584, Nil)

			tmp10586 := PrimCons(tmp10578, tmp10585)

			tmp10587 := PrimCons(symshen_4mu, tmp10586)

			tmp10588 := Call(__e, PrimNS2Value(symshen_4make__mu__application), tmp10587, V1063)

			__e.TailApply(tmp10576, tmp10588)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4aum)
			return
		}

	}, 2)

	tmp10614 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aum, tmp10574)

	_ = tmp10614

	tmp10615 := MakeNative(func(__e *ControlFlow) {
		V1066 := __e.Get(1)
		_ = V1066
		V1067 := __e.Get(2)
		_ = V1067
		tmp10616 := MakeNative(func(__e *ControlFlow) {
			VTerms := __e.Get(1)
			_ = VTerms
			tmp10617 := MakeNative(func(__e *ControlFlow) {
				VBody := __e.Get(1)
				_ = VBody
				tmp10618 := MakeNative(func(__e *ControlFlow) {
					Free := __e.Get(1)
					_ = Free
					__e.TailApply(PrimNS2Value(symshen_4cc__help), Free, V1067)
					return
				}, 1)

				tmp10619 := Call(__e, PrimNS2Value(symdifference), VBody, VTerms)

				tmp10620 := Call(__e, PrimNS2Value(symremove), symThrowcontrol, tmp10619)

				__e.TailApply(tmp10618, tmp10620)
				return

			}, 1)

			tmp10621 := Call(__e, PrimNS2Value(symshen_4extract__vars), V1067)

			__e.TailApply(tmp10617, tmp10621)
			return

		}, 1)

		tmp10622 := Call(__e, PrimNS2Value(symshen_4extract__vars), V1066)

		tmp10623 := PrimCons(symProcessN, tmp10622)

		__e.TailApply(tmp10616, tmp10623)
		return

	}, 2)

	tmp10624 := Call(__e, PrimNS1Value(symns2_1set), symshen_4continuation__call, tmp10615)

	_ = tmp10624

	tmp10625 := MakeNative(func(__e *ControlFlow) {
		V1070 := __e.Get(1)
		_ = V1070
		V1071 := __e.Get(2)
		_ = V1071
		__e.TailApply(PrimNS2Value(symshen_4remove_1h), V1070, V1071, Nil)
		return
	}, 2)

	tmp10626 := Call(__e, PrimNS1Value(symns2_1set), symremove, tmp10625)

	_ = tmp10626

	tmp10627 := MakeNative(func(__e *ControlFlow) {
		V1078 := __e.Get(1)
		_ = V1078
		V1079 := __e.Get(2)
		_ = V1079
		V1080 := __e.Get(3)
		_ = V1080
		tmp10642 := PrimEqual(Nil, V1079)

		if True == tmp10642 {
			__e.TailApply(PrimNS2Value(symreverse), V1080)
			return
		} else {
			tmp10641 := PrimIsPair(V1079)

			var ifres10637 Obj

			if True == tmp10641 {
				tmp10639 := PrimHead(V1079)

				tmp10640 := PrimEqual(tmp10639, V1078)

				var ifres10638 Obj

				if True == tmp10640 {
					ifres10638 = True

				} else {
					ifres10638 = False

				}

				ifres10637 = ifres10638

			} else {
				ifres10637 = False

			}

			if True == ifres10637 {
				tmp10635 := PrimHead(V1079)

				tmp10636 := PrimTail(V1079)

				__e.TailApply(PrimNS2Value(symshen_4remove_1h), tmp10635, tmp10636, V1080)
				return

			} else {
				tmp10634 := PrimIsPair(V1079)

				if True == tmp10634 {
					tmp10631 := PrimTail(V1079)

					tmp10632 := PrimHead(V1079)

					tmp10633 := PrimCons(tmp10632, V1080)

					__e.TailApply(PrimNS2Value(symshen_4remove_1h), V1078, tmp10631, tmp10633)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4remove_1h)
					return
				}

			}

		}

	}, 3)

	tmp10643 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1h, tmp10627)

	_ = tmp10643

	tmp10644 := MakeNative(func(__e *ControlFlow) {
		V1083 := __e.Get(1)
		_ = V1083
		V1084 := __e.Get(2)
		_ = V1084
		tmp10679 := PrimEqual(Nil, V1083)

		var ifres10676 Obj

		if True == tmp10679 {
			tmp10678 := PrimEqual(Nil, V1084)

			var ifres10677 Obj

			if True == tmp10678 {
				ifres10677 = True

			} else {
				ifres10677 = False

			}

			ifres10676 = ifres10677

		} else {
			ifres10676 = False

		}

		if True == ifres10676 {
			tmp10674 := PrimCons(symshen_4stack, Nil)

			tmp10675 := PrimCons(symshen_4the, tmp10674)

			__e.Return(PrimCons(symshen_4pop, tmp10675))
			return

		} else {
			tmp10673 := PrimEqual(Nil, V1084)

			if True == tmp10673 {
				tmp10663 := PrimCons(symshen_4stack, Nil)

				tmp10664 := PrimCons(symshen_4the, tmp10663)

				tmp10665 := PrimCons(symshen_4pop, tmp10664)

				tmp10666 := PrimCons(tmp10665, Nil)

				tmp10667 := PrimCons(symshen_4then, tmp10666)

				tmp10668 := PrimCons(symand, tmp10667)

				tmp10669 := PrimCons(V1083, tmp10668)

				tmp10670 := PrimCons(symin, tmp10669)

				tmp10671 := PrimCons(symshen_4variables, tmp10670)

				tmp10672 := PrimCons(symshen_4the, tmp10671)

				__e.Return(PrimCons(symshen_4rename, tmp10672))
				return

			} else {
				tmp10662 := PrimEqual(Nil, V1083)

				if True == tmp10662 {
					tmp10659 := PrimCons(V1084, Nil)

					tmp10660 := PrimCons(symshen_4continuation, tmp10659)

					tmp10661 := PrimCons(symshen_4the, tmp10660)

					__e.Return(PrimCons(symcall, tmp10661))
					return

				} else {
					tmp10648 := PrimCons(V1084, Nil)

					tmp10649 := PrimCons(symshen_4continuation, tmp10648)

					tmp10650 := PrimCons(symshen_4the, tmp10649)

					tmp10651 := PrimCons(symcall, tmp10650)

					tmp10652 := PrimCons(tmp10651, Nil)

					tmp10653 := PrimCons(symshen_4then, tmp10652)

					tmp10654 := PrimCons(symand, tmp10653)

					tmp10655 := PrimCons(V1083, tmp10654)

					tmp10656 := PrimCons(symin, tmp10655)

					tmp10657 := PrimCons(symshen_4variables, tmp10656)

					tmp10658 := PrimCons(symshen_4the, tmp10657)

					__e.Return(PrimCons(symshen_4rename, tmp10658))
					return

				}

			}

		}

	}, 2)

	tmp10680 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cc__help, tmp10644)

	_ = tmp10680

	tmp10681 := MakeNative(func(__e *ControlFlow) {
		V1087 := __e.Get(1)
		_ = V1087
		V1088 := __e.Get(2)
		_ = V1088
		tmp10758 := PrimIsPair(V1087)

		var ifres10731 Obj

		if True == tmp10758 {
			tmp10756 := PrimHead(V1087)

			tmp10757 := PrimEqual(symshen_4mu, tmp10756)

			var ifres10733 Obj

			if True == tmp10757 {
				tmp10754 := PrimTail(V1087)

				tmp10755 := PrimIsPair(tmp10754)

				var ifres10735 Obj

				if True == tmp10755 {
					tmp10751 := PrimTail(V1087)

					tmp10752 := PrimHead(tmp10751)

					tmp10753 := PrimEqual(Nil, tmp10752)

					var ifres10737 Obj

					if True == tmp10753 {
						tmp10748 := PrimTail(V1087)

						tmp10749 := PrimTail(tmp10748)

						tmp10750 := PrimIsPair(tmp10749)

						var ifres10739 Obj

						if True == tmp10750 {
							tmp10744 := PrimTail(V1087)

							tmp10745 := PrimTail(tmp10744)

							tmp10746 := PrimTail(tmp10745)

							tmp10747 := PrimEqual(Nil, tmp10746)

							var ifres10741 Obj

							if True == tmp10747 {
								tmp10743 := PrimEqual(Nil, V1088)

								var ifres10742 Obj

								if True == tmp10743 {
									ifres10742 = True

								} else {
									ifres10742 = False

								}

								ifres10741 = ifres10742

							} else {
								ifres10741 = False

							}

							var ifres10740 Obj

							if True == ifres10741 {
								ifres10740 = True

							} else {
								ifres10740 = False

							}

							ifres10739 = ifres10740

						} else {
							ifres10739 = False

						}

						var ifres10738 Obj

						if True == ifres10739 {
							ifres10738 = True

						} else {
							ifres10738 = False

						}

						ifres10737 = ifres10738

					} else {
						ifres10737 = False

					}

					var ifres10736 Obj

					if True == ifres10737 {
						ifres10736 = True

					} else {
						ifres10736 = False

					}

					ifres10735 = ifres10736

				} else {
					ifres10735 = False

				}

				var ifres10734 Obj

				if True == ifres10735 {
					ifres10734 = True

				} else {
					ifres10734 = False

				}

				ifres10733 = ifres10734

			} else {
				ifres10733 = False

			}

			var ifres10732 Obj

			if True == ifres10733 {
				ifres10732 = True

			} else {
				ifres10732 = False

			}

			ifres10731 = ifres10732

		} else {
			ifres10731 = False

		}

		if True == ifres10731 {
			tmp10729 := PrimTail(V1087)

			tmp10730 := PrimTail(tmp10729)

			__e.Return(PrimHead(tmp10730))
			return

		} else {
			tmp10728 := PrimIsPair(V1087)

			var ifres10701 Obj

			if True == tmp10728 {
				tmp10726 := PrimHead(V1087)

				tmp10727 := PrimEqual(symshen_4mu, tmp10726)

				var ifres10703 Obj

				if True == tmp10727 {
					tmp10724 := PrimTail(V1087)

					tmp10725 := PrimIsPair(tmp10724)

					var ifres10705 Obj

					if True == tmp10725 {
						tmp10721 := PrimTail(V1087)

						tmp10722 := PrimHead(tmp10721)

						tmp10723 := PrimIsPair(tmp10722)

						var ifres10707 Obj

						if True == tmp10723 {
							tmp10718 := PrimTail(V1087)

							tmp10719 := PrimTail(tmp10718)

							tmp10720 := PrimIsPair(tmp10719)

							var ifres10709 Obj

							if True == tmp10720 {
								tmp10714 := PrimTail(V1087)

								tmp10715 := PrimTail(tmp10714)

								tmp10716 := PrimTail(tmp10715)

								tmp10717 := PrimEqual(Nil, tmp10716)

								var ifres10711 Obj

								if True == tmp10717 {
									tmp10713 := PrimIsPair(V1088)

									var ifres10712 Obj

									if True == tmp10713 {
										ifres10712 = True

									} else {
										ifres10712 = False

									}

									ifres10711 = ifres10712

								} else {
									ifres10711 = False

								}

								var ifres10710 Obj

								if True == ifres10711 {
									ifres10710 = True

								} else {
									ifres10710 = False

								}

								ifres10709 = ifres10710

							} else {
								ifres10709 = False

							}

							var ifres10708 Obj

							if True == ifres10709 {
								ifres10708 = True

							} else {
								ifres10708 = False

							}

							ifres10707 = ifres10708

						} else {
							ifres10707 = False

						}

						var ifres10706 Obj

						if True == ifres10707 {
							ifres10706 = True

						} else {
							ifres10706 = False

						}

						ifres10705 = ifres10706

					} else {
						ifres10705 = False

					}

					var ifres10704 Obj

					if True == ifres10705 {
						ifres10704 = True

					} else {
						ifres10704 = False

					}

					ifres10703 = ifres10704

				} else {
					ifres10703 = False

				}

				var ifres10702 Obj

				if True == ifres10703 {
					ifres10702 = True

				} else {
					ifres10702 = False

				}

				ifres10701 = ifres10702

			} else {
				ifres10701 = False

			}

			if True == ifres10701 {
				tmp10684 := PrimTail(V1087)

				tmp10685 := PrimHead(tmp10684)

				tmp10686 := PrimHead(tmp10685)

				tmp10687 := PrimTail(V1087)

				tmp10688 := PrimHead(tmp10687)

				tmp10689 := PrimTail(tmp10688)

				tmp10690 := PrimTail(V1087)

				tmp10691 := PrimTail(tmp10690)

				tmp10692 := PrimCons(tmp10689, tmp10691)

				tmp10693 := PrimCons(symshen_4mu, tmp10692)

				tmp10694 := PrimTail(V1088)

				tmp10695 := Call(__e, PrimNS2Value(symshen_4make__mu__application), tmp10693, tmp10694)

				tmp10696 := PrimCons(tmp10695, Nil)

				tmp10697 := PrimCons(tmp10686, tmp10696)

				tmp10698 := PrimCons(symshen_4mu, tmp10697)

				tmp10699 := PrimHead(V1088)

				tmp10700 := PrimCons(tmp10699, Nil)

				__e.Return(PrimCons(tmp10698, tmp10700))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4make__mu__application)
				return
			}

		}

	}, 2)

	tmp10759 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make__mu__application, tmp10681)

	_ = tmp10759

	tmp10760 := MakeNative(func(__e *ControlFlow) {
		V1097 := __e.Get(1)
		_ = V1097
		V1098 := __e.Get(2)
		_ = V1098
		tmp11432 := PrimIsPair(V1097)

		var ifres11359 Obj

		if True == tmp11432 {
			tmp11430 := PrimHead(V1097)

			tmp11431 := PrimIsPair(tmp11430)

			var ifres11361 Obj

			if True == tmp11431 {
				tmp11427 := PrimHead(V1097)

				tmp11428 := PrimHead(tmp11427)

				tmp11429 := PrimEqual(symshen_4mu, tmp11428)

				var ifres11363 Obj

				if True == tmp11429 {
					tmp11424 := PrimHead(V1097)

					tmp11425 := PrimTail(tmp11424)

					tmp11426 := PrimIsPair(tmp11425)

					var ifres11365 Obj

					if True == tmp11426 {
						tmp11420 := PrimHead(V1097)

						tmp11421 := PrimTail(tmp11420)

						tmp11422 := PrimHead(tmp11421)

						tmp11423 := PrimIsPair(tmp11422)

						var ifres11367 Obj

						if True == tmp11423 {
							tmp11415 := PrimHead(V1097)

							tmp11416 := PrimTail(tmp11415)

							tmp11417 := PrimHead(tmp11416)

							tmp11418 := PrimHead(tmp11417)

							tmp11419 := PrimEqual(symmode, tmp11418)

							var ifres11369 Obj

							if True == tmp11419 {
								tmp11410 := PrimHead(V1097)

								tmp11411 := PrimTail(tmp11410)

								tmp11412 := PrimHead(tmp11411)

								tmp11413 := PrimTail(tmp11412)

								tmp11414 := PrimIsPair(tmp11413)

								var ifres11371 Obj

								if True == tmp11414 {
									tmp11404 := PrimHead(V1097)

									tmp11405 := PrimTail(tmp11404)

									tmp11406 := PrimHead(tmp11405)

									tmp11407 := PrimTail(tmp11406)

									tmp11408 := PrimTail(tmp11407)

									tmp11409 := PrimIsPair(tmp11408)

									var ifres11373 Obj

									if True == tmp11409 {
										tmp11397 := PrimHead(V1097)

										tmp11398 := PrimTail(tmp11397)

										tmp11399 := PrimHead(tmp11398)

										tmp11400 := PrimTail(tmp11399)

										tmp11401 := PrimTail(tmp11400)

										tmp11402 := PrimTail(tmp11401)

										tmp11403 := PrimEqual(Nil, tmp11402)

										var ifres11375 Obj

										if True == tmp11403 {
											tmp11393 := PrimHead(V1097)

											tmp11394 := PrimTail(tmp11393)

											tmp11395 := PrimTail(tmp11394)

											tmp11396 := PrimIsPair(tmp11395)

											var ifres11377 Obj

											if True == tmp11396 {
												tmp11388 := PrimHead(V1097)

												tmp11389 := PrimTail(tmp11388)

												tmp11390 := PrimTail(tmp11389)

												tmp11391 := PrimTail(tmp11390)

												tmp11392 := PrimEqual(Nil, tmp11391)

												var ifres11379 Obj

												if True == tmp11392 {
													tmp11386 := PrimTail(V1097)

													tmp11387 := PrimIsPair(tmp11386)

													var ifres11381 Obj

													if True == tmp11387 {
														tmp11383 := PrimTail(V1097)

														tmp11384 := PrimTail(tmp11383)

														tmp11385 := PrimEqual(Nil, tmp11384)

														var ifres11382 Obj

														if True == tmp11385 {
															ifres11382 = True

														} else {
															ifres11382 = False

														}

														ifres11381 = ifres11382

													} else {
														ifres11381 = False

													}

													var ifres11380 Obj

													if True == ifres11381 {
														ifres11380 = True

													} else {
														ifres11380 = False

													}

													ifres11379 = ifres11380

												} else {
													ifres11379 = False

												}

												var ifres11378 Obj

												if True == ifres11379 {
													ifres11378 = True

												} else {
													ifres11378 = False

												}

												ifres11377 = ifres11378

											} else {
												ifres11377 = False

											}

											var ifres11376 Obj

											if True == ifres11377 {
												ifres11376 = True

											} else {
												ifres11376 = False

											}

											ifres11375 = ifres11376

										} else {
											ifres11375 = False

										}

										var ifres11374 Obj

										if True == ifres11375 {
											ifres11374 = True

										} else {
											ifres11374 = False

										}

										ifres11373 = ifres11374

									} else {
										ifres11373 = False

									}

									var ifres11372 Obj

									if True == ifres11373 {
										ifres11372 = True

									} else {
										ifres11372 = False

									}

									ifres11371 = ifres11372

								} else {
									ifres11371 = False

								}

								var ifres11370 Obj

								if True == ifres11371 {
									ifres11370 = True

								} else {
									ifres11370 = False

								}

								ifres11369 = ifres11370

							} else {
								ifres11369 = False

							}

							var ifres11368 Obj

							if True == ifres11369 {
								ifres11368 = True

							} else {
								ifres11368 = False

							}

							ifres11367 = ifres11368

						} else {
							ifres11367 = False

						}

						var ifres11366 Obj

						if True == ifres11367 {
							ifres11366 = True

						} else {
							ifres11366 = False

						}

						ifres11365 = ifres11366

					} else {
						ifres11365 = False

					}

					var ifres11364 Obj

					if True == ifres11365 {
						ifres11364 = True

					} else {
						ifres11364 = False

					}

					ifres11363 = ifres11364

				} else {
					ifres11363 = False

				}

				var ifres11362 Obj

				if True == ifres11363 {
					ifres11362 = True

				} else {
					ifres11362 = False

				}

				ifres11361 = ifres11362

			} else {
				ifres11361 = False

			}

			var ifres11360 Obj

			if True == ifres11361 {
				ifres11360 = True

			} else {
				ifres11360 = False

			}

			ifres11359 = ifres11360

		} else {
			ifres11359 = False

		}

		if True == ifres11359 {
			tmp11341 := PrimHead(V1097)

			tmp11342 := PrimTail(tmp11341)

			tmp11343 := PrimHead(tmp11342)

			tmp11344 := PrimTail(tmp11343)

			tmp11345 := PrimHead(tmp11344)

			tmp11346 := PrimHead(V1097)

			tmp11347 := PrimTail(tmp11346)

			tmp11348 := PrimTail(tmp11347)

			tmp11349 := PrimCons(tmp11345, tmp11348)

			tmp11350 := PrimCons(symshen_4mu, tmp11349)

			tmp11351 := PrimTail(V1097)

			tmp11352 := PrimCons(tmp11350, tmp11351)

			tmp11353 := PrimHead(V1097)

			tmp11354 := PrimTail(tmp11353)

			tmp11355 := PrimHead(tmp11354)

			tmp11356 := PrimTail(tmp11355)

			tmp11357 := PrimTail(tmp11356)

			tmp11358 := PrimHead(tmp11357)

			__e.TailApply(PrimNS2Value(symshen_4mu__reduction), tmp11352, tmp11358)
			return

		} else {
			tmp11340 := PrimIsPair(V1097)

			var ifres11298 Obj

			if True == tmp11340 {
				tmp11338 := PrimHead(V1097)

				tmp11339 := PrimIsPair(tmp11338)

				var ifres11300 Obj

				if True == tmp11339 {
					tmp11335 := PrimHead(V1097)

					tmp11336 := PrimHead(tmp11335)

					tmp11337 := PrimEqual(symshen_4mu, tmp11336)

					var ifres11302 Obj

					if True == tmp11337 {
						tmp11332 := PrimHead(V1097)

						tmp11333 := PrimTail(tmp11332)

						tmp11334 := PrimIsPair(tmp11333)

						var ifres11304 Obj

						if True == tmp11334 {
							tmp11328 := PrimHead(V1097)

							tmp11329 := PrimTail(tmp11328)

							tmp11330 := PrimTail(tmp11329)

							tmp11331 := PrimIsPair(tmp11330)

							var ifres11306 Obj

							if True == tmp11331 {
								tmp11323 := PrimHead(V1097)

								tmp11324 := PrimTail(tmp11323)

								tmp11325 := PrimTail(tmp11324)

								tmp11326 := PrimTail(tmp11325)

								tmp11327 := PrimEqual(Nil, tmp11326)

								var ifres11308 Obj

								if True == tmp11327 {
									tmp11321 := PrimTail(V1097)

									tmp11322 := PrimIsPair(tmp11321)

									var ifres11310 Obj

									if True == tmp11322 {
										tmp11318 := PrimTail(V1097)

										tmp11319 := PrimTail(tmp11318)

										tmp11320 := PrimEqual(Nil, tmp11319)

										var ifres11312 Obj

										if True == tmp11320 {
											tmp11314 := PrimHead(V1097)

											tmp11315 := PrimTail(tmp11314)

											tmp11316 := PrimHead(tmp11315)

											tmp11317 := PrimEqual(sym__, tmp11316)

											var ifres11313 Obj

											if True == tmp11317 {
												ifres11313 = True

											} else {
												ifres11313 = False

											}

											ifres11312 = ifres11313

										} else {
											ifres11312 = False

										}

										var ifres11311 Obj

										if True == ifres11312 {
											ifres11311 = True

										} else {
											ifres11311 = False

										}

										ifres11310 = ifres11311

									} else {
										ifres11310 = False

									}

									var ifres11309 Obj

									if True == ifres11310 {
										ifres11309 = True

									} else {
										ifres11309 = False

									}

									ifres11308 = ifres11309

								} else {
									ifres11308 = False

								}

								var ifres11307 Obj

								if True == ifres11308 {
									ifres11307 = True

								} else {
									ifres11307 = False

								}

								ifres11306 = ifres11307

							} else {
								ifres11306 = False

							}

							var ifres11305 Obj

							if True == ifres11306 {
								ifres11305 = True

							} else {
								ifres11305 = False

							}

							ifres11304 = ifres11305

						} else {
							ifres11304 = False

						}

						var ifres11303 Obj

						if True == ifres11304 {
							ifres11303 = True

						} else {
							ifres11303 = False

						}

						ifres11302 = ifres11303

					} else {
						ifres11302 = False

					}

					var ifres11301 Obj

					if True == ifres11302 {
						ifres11301 = True

					} else {
						ifres11301 = False

					}

					ifres11300 = ifres11301

				} else {
					ifres11300 = False

				}

				var ifres11299 Obj

				if True == ifres11300 {
					ifres11299 = True

				} else {
					ifres11299 = False

				}

				ifres11298 = ifres11299

			} else {
				ifres11298 = False

			}

			if True == ifres11298 {
				tmp11294 := PrimHead(V1097)

				tmp11295 := PrimTail(tmp11294)

				tmp11296 := PrimTail(tmp11295)

				tmp11297 := PrimHead(tmp11296)

				__e.TailApply(PrimNS2Value(symshen_4mu__reduction), tmp11297, V1098)
				return

			} else {
				tmp11293 := PrimIsPair(V1097)

				var ifres11249 Obj

				if True == tmp11293 {
					tmp11291 := PrimHead(V1097)

					tmp11292 := PrimIsPair(tmp11291)

					var ifres11251 Obj

					if True == tmp11292 {
						tmp11288 := PrimHead(V1097)

						tmp11289 := PrimHead(tmp11288)

						tmp11290 := PrimEqual(symshen_4mu, tmp11289)

						var ifres11253 Obj

						if True == tmp11290 {
							tmp11285 := PrimHead(V1097)

							tmp11286 := PrimTail(tmp11285)

							tmp11287 := PrimIsPair(tmp11286)

							var ifres11255 Obj

							if True == tmp11287 {
								tmp11281 := PrimHead(V1097)

								tmp11282 := PrimTail(tmp11281)

								tmp11283 := PrimTail(tmp11282)

								tmp11284 := PrimIsPair(tmp11283)

								var ifres11257 Obj

								if True == tmp11284 {
									tmp11276 := PrimHead(V1097)

									tmp11277 := PrimTail(tmp11276)

									tmp11278 := PrimTail(tmp11277)

									tmp11279 := PrimTail(tmp11278)

									tmp11280 := PrimEqual(Nil, tmp11279)

									var ifres11259 Obj

									if True == tmp11280 {
										tmp11274 := PrimTail(V1097)

										tmp11275 := PrimIsPair(tmp11274)

										var ifres11261 Obj

										if True == tmp11275 {
											tmp11271 := PrimTail(V1097)

											tmp11272 := PrimTail(tmp11271)

											tmp11273 := PrimEqual(Nil, tmp11272)

											var ifres11263 Obj

											if True == tmp11273 {
												tmp11265 := PrimHead(V1097)

												tmp11266 := PrimTail(tmp11265)

												tmp11267 := PrimHead(tmp11266)

												tmp11268 := PrimTail(V1097)

												tmp11269 := PrimHead(tmp11268)

												tmp11270 := Call(__e, PrimNS2Value(symshen_4ephemeral__variable_2), tmp11267, tmp11269)

												var ifres11264 Obj

												if True == tmp11270 {
													ifres11264 = True

												} else {
													ifres11264 = False

												}

												ifres11263 = ifres11264

											} else {
												ifres11263 = False

											}

											var ifres11262 Obj

											if True == ifres11263 {
												ifres11262 = True

											} else {
												ifres11262 = False

											}

											ifres11261 = ifres11262

										} else {
											ifres11261 = False

										}

										var ifres11260 Obj

										if True == ifres11261 {
											ifres11260 = True

										} else {
											ifres11260 = False

										}

										ifres11259 = ifres11260

									} else {
										ifres11259 = False

									}

									var ifres11258 Obj

									if True == ifres11259 {
										ifres11258 = True

									} else {
										ifres11258 = False

									}

									ifres11257 = ifres11258

								} else {
									ifres11257 = False

								}

								var ifres11256 Obj

								if True == ifres11257 {
									ifres11256 = True

								} else {
									ifres11256 = False

								}

								ifres11255 = ifres11256

							} else {
								ifres11255 = False

							}

							var ifres11254 Obj

							if True == ifres11255 {
								ifres11254 = True

							} else {
								ifres11254 = False

							}

							ifres11253 = ifres11254

						} else {
							ifres11253 = False

						}

						var ifres11252 Obj

						if True == ifres11253 {
							ifres11252 = True

						} else {
							ifres11252 = False

						}

						ifres11251 = ifres11252

					} else {
						ifres11251 = False

					}

					var ifres11250 Obj

					if True == ifres11251 {
						ifres11250 = True

					} else {
						ifres11250 = False

					}

					ifres11249 = ifres11250

				} else {
					ifres11249 = False

				}

				if True == ifres11249 {
					tmp11239 := PrimTail(V1097)

					tmp11240 := PrimHead(tmp11239)

					tmp11241 := PrimHead(V1097)

					tmp11242 := PrimTail(tmp11241)

					tmp11243 := PrimHead(tmp11242)

					tmp11244 := PrimHead(V1097)

					tmp11245 := PrimTail(tmp11244)

					tmp11246 := PrimTail(tmp11245)

					tmp11247 := PrimHead(tmp11246)

					tmp11248 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp11247, V1098)

					__e.TailApply(PrimNS2Value(symsubst), tmp11240, tmp11243, tmp11248)
					return

				} else {
					tmp11238 := PrimIsPair(V1097)

					var ifres11196 Obj

					if True == tmp11238 {
						tmp11236 := PrimHead(V1097)

						tmp11237 := PrimIsPair(tmp11236)

						var ifres11198 Obj

						if True == tmp11237 {
							tmp11233 := PrimHead(V1097)

							tmp11234 := PrimHead(tmp11233)

							tmp11235 := PrimEqual(symshen_4mu, tmp11234)

							var ifres11200 Obj

							if True == tmp11235 {
								tmp11230 := PrimHead(V1097)

								tmp11231 := PrimTail(tmp11230)

								tmp11232 := PrimIsPair(tmp11231)

								var ifres11202 Obj

								if True == tmp11232 {
									tmp11226 := PrimHead(V1097)

									tmp11227 := PrimTail(tmp11226)

									tmp11228 := PrimTail(tmp11227)

									tmp11229 := PrimIsPair(tmp11228)

									var ifres11204 Obj

									if True == tmp11229 {
										tmp11221 := PrimHead(V1097)

										tmp11222 := PrimTail(tmp11221)

										tmp11223 := PrimTail(tmp11222)

										tmp11224 := PrimTail(tmp11223)

										tmp11225 := PrimEqual(Nil, tmp11224)

										var ifres11206 Obj

										if True == tmp11225 {
											tmp11219 := PrimTail(V1097)

											tmp11220 := PrimIsPair(tmp11219)

											var ifres11208 Obj

											if True == tmp11220 {
												tmp11216 := PrimTail(V1097)

												tmp11217 := PrimTail(tmp11216)

												tmp11218 := PrimEqual(Nil, tmp11217)

												var ifres11210 Obj

												if True == tmp11218 {
													tmp11212 := PrimHead(V1097)

													tmp11213 := PrimTail(tmp11212)

													tmp11214 := PrimHead(tmp11213)

													tmp11215 := PrimIsVariable(tmp11214)

													var ifres11211 Obj

													if True == tmp11215 {
														ifres11211 = True

													} else {
														ifres11211 = False

													}

													ifres11210 = ifres11211

												} else {
													ifres11210 = False

												}

												var ifres11209 Obj

												if True == ifres11210 {
													ifres11209 = True

												} else {
													ifres11209 = False

												}

												ifres11208 = ifres11209

											} else {
												ifres11208 = False

											}

											var ifres11207 Obj

											if True == ifres11208 {
												ifres11207 = True

											} else {
												ifres11207 = False

											}

											ifres11206 = ifres11207

										} else {
											ifres11206 = False

										}

										var ifres11205 Obj

										if True == ifres11206 {
											ifres11205 = True

										} else {
											ifres11205 = False

										}

										ifres11204 = ifres11205

									} else {
										ifres11204 = False

									}

									var ifres11203 Obj

									if True == ifres11204 {
										ifres11203 = True

									} else {
										ifres11203 = False

									}

									ifres11202 = ifres11203

								} else {
									ifres11202 = False

								}

								var ifres11201 Obj

								if True == ifres11202 {
									ifres11201 = True

								} else {
									ifres11201 = False

								}

								ifres11200 = ifres11201

							} else {
								ifres11200 = False

							}

							var ifres11199 Obj

							if True == ifres11200 {
								ifres11199 = True

							} else {
								ifres11199 = False

							}

							ifres11198 = ifres11199

						} else {
							ifres11198 = False

						}

						var ifres11197 Obj

						if True == ifres11198 {
							ifres11197 = True

						} else {
							ifres11197 = False

						}

						ifres11196 = ifres11197

					} else {
						ifres11196 = False

					}

					if True == ifres11196 {
						tmp11181 := PrimHead(V1097)

						tmp11182 := PrimTail(tmp11181)

						tmp11183 := PrimHead(tmp11182)

						tmp11184 := PrimTail(V1097)

						tmp11185 := PrimHead(tmp11184)

						tmp11186 := PrimHead(V1097)

						tmp11187 := PrimTail(tmp11186)

						tmp11188 := PrimTail(tmp11187)

						tmp11189 := PrimHead(tmp11188)

						tmp11190 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp11189, V1098)

						tmp11191 := PrimCons(tmp11190, Nil)

						tmp11192 := PrimCons(symin, tmp11191)

						tmp11193 := PrimCons(tmp11185, tmp11192)

						tmp11194 := PrimCons(symshen_4be, tmp11193)

						tmp11195 := PrimCons(tmp11183, tmp11194)

						__e.Return(PrimCons(symlet, tmp11195))
						return

					} else {
						tmp11180 := PrimIsPair(V1097)

						var ifres11135 Obj

						if True == tmp11180 {
							tmp11178 := PrimHead(V1097)

							tmp11179 := PrimIsPair(tmp11178)

							var ifres11137 Obj

							if True == tmp11179 {
								tmp11175 := PrimHead(V1097)

								tmp11176 := PrimHead(tmp11175)

								tmp11177 := PrimEqual(symshen_4mu, tmp11176)

								var ifres11139 Obj

								if True == tmp11177 {
									tmp11172 := PrimHead(V1097)

									tmp11173 := PrimTail(tmp11172)

									tmp11174 := PrimIsPair(tmp11173)

									var ifres11141 Obj

									if True == tmp11174 {
										tmp11168 := PrimHead(V1097)

										tmp11169 := PrimTail(tmp11168)

										tmp11170 := PrimTail(tmp11169)

										tmp11171 := PrimIsPair(tmp11170)

										var ifres11143 Obj

										if True == tmp11171 {
											tmp11163 := PrimHead(V1097)

											tmp11164 := PrimTail(tmp11163)

											tmp11165 := PrimTail(tmp11164)

											tmp11166 := PrimTail(tmp11165)

											tmp11167 := PrimEqual(Nil, tmp11166)

											var ifres11145 Obj

											if True == tmp11167 {
												tmp11161 := PrimTail(V1097)

												tmp11162 := PrimIsPair(tmp11161)

												var ifres11147 Obj

												if True == tmp11162 {
													tmp11158 := PrimTail(V1097)

													tmp11159 := PrimTail(tmp11158)

													tmp11160 := PrimEqual(Nil, tmp11159)

													var ifres11149 Obj

													if True == tmp11160 {
														tmp11157 := PrimEqual(sym_1, V1098)

														var ifres11151 Obj

														if True == tmp11157 {
															tmp11153 := PrimHead(V1097)

															tmp11154 := PrimTail(tmp11153)

															tmp11155 := PrimHead(tmp11154)

															tmp11156 := Call(__e, PrimNS2Value(symshen_4prolog__constant_2), tmp11155)

															var ifres11152 Obj

															if True == tmp11156 {
																ifres11152 = True

															} else {
																ifres11152 = False

															}

															ifres11151 = ifres11152

														} else {
															ifres11151 = False

														}

														var ifres11150 Obj

														if True == ifres11151 {
															ifres11150 = True

														} else {
															ifres11150 = False

														}

														ifres11149 = ifres11150

													} else {
														ifres11149 = False

													}

													var ifres11148 Obj

													if True == ifres11149 {
														ifres11148 = True

													} else {
														ifres11148 = False

													}

													ifres11147 = ifres11148

												} else {
													ifres11147 = False

												}

												var ifres11146 Obj

												if True == ifres11147 {
													ifres11146 = True

												} else {
													ifres11146 = False

												}

												ifres11145 = ifres11146

											} else {
												ifres11145 = False

											}

											var ifres11144 Obj

											if True == ifres11145 {
												ifres11144 = True

											} else {
												ifres11144 = False

											}

											ifres11143 = ifres11144

										} else {
											ifres11143 = False

										}

										var ifres11142 Obj

										if True == ifres11143 {
											ifres11142 = True

										} else {
											ifres11142 = False

										}

										ifres11141 = ifres11142

									} else {
										ifres11141 = False

									}

									var ifres11140 Obj

									if True == ifres11141 {
										ifres11140 = True

									} else {
										ifres11140 = False

									}

									ifres11139 = ifres11140

								} else {
									ifres11139 = False

								}

								var ifres11138 Obj

								if True == ifres11139 {
									ifres11138 = True

								} else {
									ifres11138 = False

								}

								ifres11137 = ifres11138

							} else {
								ifres11137 = False

							}

							var ifres11136 Obj

							if True == ifres11137 {
								ifres11136 = True

							} else {
								ifres11136 = False

							}

							ifres11135 = ifres11136

						} else {
							ifres11135 = False

						}

						if True == ifres11135 {
							tmp11104 := MakeNative(func(__e *ControlFlow) {
								Z := __e.Get(1)
								_ = Z
								tmp11105 := PrimTail(V1097)

								tmp11106 := PrimCons(symshen_4dereferencing, tmp11105)

								tmp11107 := PrimCons(symshen_4of, tmp11106)

								tmp11108 := PrimCons(symshen_4result, tmp11107)

								tmp11109 := PrimCons(symshen_4the, tmp11108)

								tmp11110 := PrimHead(V1097)

								tmp11111 := PrimTail(tmp11110)

								tmp11112 := PrimHead(tmp11111)

								tmp11113 := PrimCons(tmp11112, Nil)

								tmp11114 := PrimCons(symshen_4to, tmp11113)

								tmp11115 := PrimCons(symidentical, tmp11114)

								tmp11116 := PrimCons(symis, tmp11115)

								tmp11117 := PrimCons(Z, tmp11116)

								tmp11118 := PrimHead(V1097)

								tmp11119 := PrimTail(tmp11118)

								tmp11120 := PrimTail(tmp11119)

								tmp11121 := PrimHead(tmp11120)

								tmp11122 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp11121, sym_1)

								tmp11123 := PrimCons(symshen_4failed_b, Nil)

								tmp11124 := PrimCons(symshen_4else, tmp11123)

								tmp11125 := PrimCons(tmp11122, tmp11124)

								tmp11126 := PrimCons(symshen_4then, tmp11125)

								tmp11127 := PrimCons(tmp11117, tmp11126)

								tmp11128 := PrimCons(symif, tmp11127)

								tmp11129 := PrimCons(tmp11128, Nil)

								tmp11130 := PrimCons(symin, tmp11129)

								tmp11131 := PrimCons(tmp11109, tmp11130)

								tmp11132 := PrimCons(symshen_4be, tmp11131)

								tmp11133 := PrimCons(Z, tmp11132)

								__e.Return(PrimCons(symlet, tmp11133))
								return

							}, 1)

							tmp11134 := Call(__e, PrimNS2Value(symgensym), symV)

							__e.TailApply(tmp11104, tmp11134)
							return

						} else {
							tmp11103 := PrimIsPair(V1097)

							var ifres11058 Obj

							if True == tmp11103 {
								tmp11101 := PrimHead(V1097)

								tmp11102 := PrimIsPair(tmp11101)

								var ifres11060 Obj

								if True == tmp11102 {
									tmp11098 := PrimHead(V1097)

									tmp11099 := PrimHead(tmp11098)

									tmp11100 := PrimEqual(symshen_4mu, tmp11099)

									var ifres11062 Obj

									if True == tmp11100 {
										tmp11095 := PrimHead(V1097)

										tmp11096 := PrimTail(tmp11095)

										tmp11097 := PrimIsPair(tmp11096)

										var ifres11064 Obj

										if True == tmp11097 {
											tmp11091 := PrimHead(V1097)

											tmp11092 := PrimTail(tmp11091)

											tmp11093 := PrimTail(tmp11092)

											tmp11094 := PrimIsPair(tmp11093)

											var ifres11066 Obj

											if True == tmp11094 {
												tmp11086 := PrimHead(V1097)

												tmp11087 := PrimTail(tmp11086)

												tmp11088 := PrimTail(tmp11087)

												tmp11089 := PrimTail(tmp11088)

												tmp11090 := PrimEqual(Nil, tmp11089)

												var ifres11068 Obj

												if True == tmp11090 {
													tmp11084 := PrimTail(V1097)

													tmp11085 := PrimIsPair(tmp11084)

													var ifres11070 Obj

													if True == tmp11085 {
														tmp11081 := PrimTail(V1097)

														tmp11082 := PrimTail(tmp11081)

														tmp11083 := PrimEqual(Nil, tmp11082)

														var ifres11072 Obj

														if True == tmp11083 {
															tmp11080 := PrimEqual(sym_7, V1098)

															var ifres11074 Obj

															if True == tmp11080 {
																tmp11076 := PrimHead(V1097)

																tmp11077 := PrimTail(tmp11076)

																tmp11078 := PrimHead(tmp11077)

																tmp11079 := Call(__e, PrimNS2Value(symshen_4prolog__constant_2), tmp11078)

																var ifres11075 Obj

																if True == tmp11079 {
																	ifres11075 = True

																} else {
																	ifres11075 = False

																}

																ifres11074 = ifres11075

															} else {
																ifres11074 = False

															}

															var ifres11073 Obj

															if True == ifres11074 {
																ifres11073 = True

															} else {
																ifres11073 = False

															}

															ifres11072 = ifres11073

														} else {
															ifres11072 = False

														}

														var ifres11071 Obj

														if True == ifres11072 {
															ifres11071 = True

														} else {
															ifres11071 = False

														}

														ifres11070 = ifres11071

													} else {
														ifres11070 = False

													}

													var ifres11069 Obj

													if True == ifres11070 {
														ifres11069 = True

													} else {
														ifres11069 = False

													}

													ifres11068 = ifres11069

												} else {
													ifres11068 = False

												}

												var ifres11067 Obj

												if True == ifres11068 {
													ifres11067 = True

												} else {
													ifres11067 = False

												}

												ifres11066 = ifres11067

											} else {
												ifres11066 = False

											}

											var ifres11065 Obj

											if True == ifres11066 {
												ifres11065 = True

											} else {
												ifres11065 = False

											}

											ifres11064 = ifres11065

										} else {
											ifres11064 = False

										}

										var ifres11063 Obj

										if True == ifres11064 {
											ifres11063 = True

										} else {
											ifres11063 = False

										}

										ifres11062 = ifres11063

									} else {
										ifres11062 = False

									}

									var ifres11061 Obj

									if True == ifres11062 {
										ifres11061 = True

									} else {
										ifres11061 = False

									}

									ifres11060 = ifres11061

								} else {
									ifres11060 = False

								}

								var ifres11059 Obj

								if True == ifres11060 {
									ifres11059 = True

								} else {
									ifres11059 = False

								}

								ifres11058 = ifres11059

							} else {
								ifres11058 = False

							}

							if True == ifres11058 {
								tmp11003 := MakeNative(func(__e *ControlFlow) {
									Z := __e.Get(1)
									_ = Z
									tmp11004 := PrimTail(V1097)

									tmp11005 := PrimCons(symshen_4dereferencing, tmp11004)

									tmp11006 := PrimCons(symshen_4of, tmp11005)

									tmp11007 := PrimCons(symshen_4result, tmp11006)

									tmp11008 := PrimCons(symshen_4the, tmp11007)

									tmp11009 := PrimHead(V1097)

									tmp11010 := PrimTail(tmp11009)

									tmp11011 := PrimHead(tmp11010)

									tmp11012 := PrimCons(tmp11011, Nil)

									tmp11013 := PrimCons(symshen_4to, tmp11012)

									tmp11014 := PrimCons(symidentical, tmp11013)

									tmp11015 := PrimCons(symis, tmp11014)

									tmp11016 := PrimCons(Z, tmp11015)

									tmp11017 := PrimHead(V1097)

									tmp11018 := PrimTail(tmp11017)

									tmp11019 := PrimTail(tmp11018)

									tmp11020 := PrimHead(tmp11019)

									tmp11021 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp11020, sym_7)

									tmp11022 := PrimCons(symshen_4variable, Nil)

									tmp11023 := PrimCons(symshen_4a, tmp11022)

									tmp11024 := PrimCons(symis, tmp11023)

									tmp11025 := PrimCons(Z, tmp11024)

									tmp11026 := PrimHead(V1097)

									tmp11027 := PrimTail(tmp11026)

									tmp11028 := PrimHead(tmp11027)

									tmp11029 := PrimHead(V1097)

									tmp11030 := PrimTail(tmp11029)

									tmp11031 := PrimTail(tmp11030)

									tmp11032 := PrimHead(tmp11031)

									tmp11033 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp11032, sym_7)

									tmp11034 := PrimCons(tmp11033, Nil)

									tmp11035 := PrimCons(symin, tmp11034)

									tmp11036 := PrimCons(tmp11028, tmp11035)

									tmp11037 := PrimCons(symshen_4to, tmp11036)

									tmp11038 := PrimCons(Z, tmp11037)

									tmp11039 := PrimCons(symbind, tmp11038)

									tmp11040 := PrimCons(symshen_4failed_b, Nil)

									tmp11041 := PrimCons(symshen_4else, tmp11040)

									tmp11042 := PrimCons(tmp11039, tmp11041)

									tmp11043 := PrimCons(symshen_4then, tmp11042)

									tmp11044 := PrimCons(tmp11025, tmp11043)

									tmp11045 := PrimCons(symif, tmp11044)

									tmp11046 := PrimCons(tmp11045, Nil)

									tmp11047 := PrimCons(symshen_4else, tmp11046)

									tmp11048 := PrimCons(tmp11021, tmp11047)

									tmp11049 := PrimCons(symshen_4then, tmp11048)

									tmp11050 := PrimCons(tmp11016, tmp11049)

									tmp11051 := PrimCons(symif, tmp11050)

									tmp11052 := PrimCons(tmp11051, Nil)

									tmp11053 := PrimCons(symin, tmp11052)

									tmp11054 := PrimCons(tmp11008, tmp11053)

									tmp11055 := PrimCons(symshen_4be, tmp11054)

									tmp11056 := PrimCons(Z, tmp11055)

									__e.Return(PrimCons(symlet, tmp11056))
									return

								}, 1)

								tmp11057 := Call(__e, PrimNS2Value(symgensym), symV)

								__e.TailApply(tmp11003, tmp11057)
								return

							} else {
								tmp11002 := PrimIsPair(V1097)

								var ifres10957 Obj

								if True == tmp11002 {
									tmp11000 := PrimHead(V1097)

									tmp11001 := PrimIsPair(tmp11000)

									var ifres10959 Obj

									if True == tmp11001 {
										tmp10997 := PrimHead(V1097)

										tmp10998 := PrimHead(tmp10997)

										tmp10999 := PrimEqual(symshen_4mu, tmp10998)

										var ifres10961 Obj

										if True == tmp10999 {
											tmp10994 := PrimHead(V1097)

											tmp10995 := PrimTail(tmp10994)

											tmp10996 := PrimIsPair(tmp10995)

											var ifres10963 Obj

											if True == tmp10996 {
												tmp10990 := PrimHead(V1097)

												tmp10991 := PrimTail(tmp10990)

												tmp10992 := PrimHead(tmp10991)

												tmp10993 := PrimIsPair(tmp10992)

												var ifres10965 Obj

												if True == tmp10993 {
													tmp10986 := PrimHead(V1097)

													tmp10987 := PrimTail(tmp10986)

													tmp10988 := PrimTail(tmp10987)

													tmp10989 := PrimIsPair(tmp10988)

													var ifres10967 Obj

													if True == tmp10989 {
														tmp10981 := PrimHead(V1097)

														tmp10982 := PrimTail(tmp10981)

														tmp10983 := PrimTail(tmp10982)

														tmp10984 := PrimTail(tmp10983)

														tmp10985 := PrimEqual(Nil, tmp10984)

														var ifres10969 Obj

														if True == tmp10985 {
															tmp10979 := PrimTail(V1097)

															tmp10980 := PrimIsPair(tmp10979)

															var ifres10971 Obj

															if True == tmp10980 {
																tmp10976 := PrimTail(V1097)

																tmp10977 := PrimTail(tmp10976)

																tmp10978 := PrimEqual(Nil, tmp10977)

																var ifres10973 Obj

																if True == tmp10978 {
																	tmp10975 := PrimEqual(sym_1, V1098)

																	var ifres10974 Obj

																	if True == tmp10975 {
																		ifres10974 = True

																	} else {
																		ifres10974 = False

																	}

																	ifres10973 = ifres10974

																} else {
																	ifres10973 = False

																}

																var ifres10972 Obj

																if True == ifres10973 {
																	ifres10972 = True

																} else {
																	ifres10972 = False

																}

																ifres10971 = ifres10972

															} else {
																ifres10971 = False

															}

															var ifres10970 Obj

															if True == ifres10971 {
																ifres10970 = True

															} else {
																ifres10970 = False

															}

															ifres10969 = ifres10970

														} else {
															ifres10969 = False

														}

														var ifres10968 Obj

														if True == ifres10969 {
															ifres10968 = True

														} else {
															ifres10968 = False

														}

														ifres10967 = ifres10968

													} else {
														ifres10967 = False

													}

													var ifres10966 Obj

													if True == ifres10967 {
														ifres10966 = True

													} else {
														ifres10966 = False

													}

													ifres10965 = ifres10966

												} else {
													ifres10965 = False

												}

												var ifres10964 Obj

												if True == ifres10965 {
													ifres10964 = True

												} else {
													ifres10964 = False

												}

												ifres10963 = ifres10964

											} else {
												ifres10963 = False

											}

											var ifres10962 Obj

											if True == ifres10963 {
												ifres10962 = True

											} else {
												ifres10962 = False

											}

											ifres10961 = ifres10962

										} else {
											ifres10961 = False

										}

										var ifres10960 Obj

										if True == ifres10961 {
											ifres10960 = True

										} else {
											ifres10960 = False

										}

										ifres10959 = ifres10960

									} else {
										ifres10959 = False

									}

									var ifres10958 Obj

									if True == ifres10959 {
										ifres10958 = True

									} else {
										ifres10958 = False

									}

									ifres10957 = ifres10958

								} else {
									ifres10957 = False

								}

								if True == ifres10957 {
									tmp10905 := MakeNative(func(__e *ControlFlow) {
										Z := __e.Get(1)
										_ = Z
										tmp10906 := PrimTail(V1097)

										tmp10907 := PrimCons(symshen_4dereferencing, tmp10906)

										tmp10908 := PrimCons(symshen_4of, tmp10907)

										tmp10909 := PrimCons(symshen_4result, tmp10908)

										tmp10910 := PrimCons(symshen_4the, tmp10909)

										tmp10911 := PrimCons(symlist, Nil)

										tmp10912 := PrimCons(symshen_4non_1empty, tmp10911)

										tmp10913 := PrimCons(symshen_4a, tmp10912)

										tmp10914 := PrimCons(symis, tmp10913)

										tmp10915 := PrimCons(Z, tmp10914)

										tmp10916 := PrimHead(V1097)

										tmp10917 := PrimTail(tmp10916)

										tmp10918 := PrimHead(tmp10917)

										tmp10919 := PrimHead(tmp10918)

										tmp10920 := PrimHead(V1097)

										tmp10921 := PrimTail(tmp10920)

										tmp10922 := PrimHead(tmp10921)

										tmp10923 := PrimTail(tmp10922)

										tmp10924 := PrimHead(V1097)

										tmp10925 := PrimTail(tmp10924)

										tmp10926 := PrimTail(tmp10925)

										tmp10927 := PrimCons(tmp10923, tmp10926)

										tmp10928 := PrimCons(symshen_4mu, tmp10927)

										tmp10929 := PrimCons(Z, Nil)

										tmp10930 := PrimCons(symshen_4of, tmp10929)

										tmp10931 := PrimCons(symtail, tmp10930)

										tmp10932 := PrimCons(symshen_4the, tmp10931)

										tmp10933 := PrimCons(tmp10932, Nil)

										tmp10934 := PrimCons(tmp10928, tmp10933)

										tmp10935 := PrimCons(tmp10934, Nil)

										tmp10936 := PrimCons(tmp10919, tmp10935)

										tmp10937 := PrimCons(symshen_4mu, tmp10936)

										tmp10938 := PrimCons(Z, Nil)

										tmp10939 := PrimCons(symshen_4of, tmp10938)

										tmp10940 := PrimCons(symhead, tmp10939)

										tmp10941 := PrimCons(symshen_4the, tmp10940)

										tmp10942 := PrimCons(tmp10941, Nil)

										tmp10943 := PrimCons(tmp10937, tmp10942)

										tmp10944 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp10943, sym_1)

										tmp10945 := PrimCons(symshen_4failed_b, Nil)

										tmp10946 := PrimCons(symshen_4else, tmp10945)

										tmp10947 := PrimCons(tmp10944, tmp10946)

										tmp10948 := PrimCons(symshen_4then, tmp10947)

										tmp10949 := PrimCons(tmp10915, tmp10948)

										tmp10950 := PrimCons(symif, tmp10949)

										tmp10951 := PrimCons(tmp10950, Nil)

										tmp10952 := PrimCons(symin, tmp10951)

										tmp10953 := PrimCons(tmp10910, tmp10952)

										tmp10954 := PrimCons(symshen_4be, tmp10953)

										tmp10955 := PrimCons(Z, tmp10954)

										__e.Return(PrimCons(symlet, tmp10955))
										return

									}, 1)

									tmp10956 := Call(__e, PrimNS2Value(symgensym), symV)

									__e.TailApply(tmp10905, tmp10956)
									return

								} else {
									tmp10904 := PrimIsPair(V1097)

									var ifres10859 Obj

									if True == tmp10904 {
										tmp10902 := PrimHead(V1097)

										tmp10903 := PrimIsPair(tmp10902)

										var ifres10861 Obj

										if True == tmp10903 {
											tmp10899 := PrimHead(V1097)

											tmp10900 := PrimHead(tmp10899)

											tmp10901 := PrimEqual(symshen_4mu, tmp10900)

											var ifres10863 Obj

											if True == tmp10901 {
												tmp10896 := PrimHead(V1097)

												tmp10897 := PrimTail(tmp10896)

												tmp10898 := PrimIsPair(tmp10897)

												var ifres10865 Obj

												if True == tmp10898 {
													tmp10892 := PrimHead(V1097)

													tmp10893 := PrimTail(tmp10892)

													tmp10894 := PrimHead(tmp10893)

													tmp10895 := PrimIsPair(tmp10894)

													var ifres10867 Obj

													if True == tmp10895 {
														tmp10888 := PrimHead(V1097)

														tmp10889 := PrimTail(tmp10888)

														tmp10890 := PrimTail(tmp10889)

														tmp10891 := PrimIsPair(tmp10890)

														var ifres10869 Obj

														if True == tmp10891 {
															tmp10883 := PrimHead(V1097)

															tmp10884 := PrimTail(tmp10883)

															tmp10885 := PrimTail(tmp10884)

															tmp10886 := PrimTail(tmp10885)

															tmp10887 := PrimEqual(Nil, tmp10886)

															var ifres10871 Obj

															if True == tmp10887 {
																tmp10881 := PrimTail(V1097)

																tmp10882 := PrimIsPair(tmp10881)

																var ifres10873 Obj

																if True == tmp10882 {
																	tmp10878 := PrimTail(V1097)

																	tmp10879 := PrimTail(tmp10878)

																	tmp10880 := PrimEqual(Nil, tmp10879)

																	var ifres10875 Obj

																	if True == tmp10880 {
																		tmp10877 := PrimEqual(sym_7, V1098)

																		var ifres10876 Obj

																		if True == tmp10877 {
																			ifres10876 = True

																		} else {
																			ifres10876 = False

																		}

																		ifres10875 = ifres10876

																	} else {
																		ifres10875 = False

																	}

																	var ifres10874 Obj

																	if True == ifres10875 {
																		ifres10874 = True

																	} else {
																		ifres10874 = False

																	}

																	ifres10873 = ifres10874

																} else {
																	ifres10873 = False

																}

																var ifres10872 Obj

																if True == ifres10873 {
																	ifres10872 = True

																} else {
																	ifres10872 = False

																}

																ifres10871 = ifres10872

															} else {
																ifres10871 = False

															}

															var ifres10870 Obj

															if True == ifres10871 {
																ifres10870 = True

															} else {
																ifres10870 = False

															}

															ifres10869 = ifres10870

														} else {
															ifres10869 = False

														}

														var ifres10868 Obj

														if True == ifres10869 {
															ifres10868 = True

														} else {
															ifres10868 = False

														}

														ifres10867 = ifres10868

													} else {
														ifres10867 = False

													}

													var ifres10866 Obj

													if True == ifres10867 {
														ifres10866 = True

													} else {
														ifres10866 = False

													}

													ifres10865 = ifres10866

												} else {
													ifres10865 = False

												}

												var ifres10864 Obj

												if True == ifres10865 {
													ifres10864 = True

												} else {
													ifres10864 = False

												}

												ifres10863 = ifres10864

											} else {
												ifres10863 = False

											}

											var ifres10862 Obj

											if True == ifres10863 {
												ifres10862 = True

											} else {
												ifres10862 = False

											}

											ifres10861 = ifres10862

										} else {
											ifres10861 = False

										}

										var ifres10860 Obj

										if True == ifres10861 {
											ifres10860 = True

										} else {
											ifres10860 = False

										}

										ifres10859 = ifres10860

									} else {
										ifres10859 = False

									}

									if True == ifres10859 {
										tmp10769 := MakeNative(func(__e *ControlFlow) {
											Z := __e.Get(1)
											_ = Z
											tmp10770 := PrimTail(V1097)

											tmp10771 := PrimCons(symshen_4dereferencing, tmp10770)

											tmp10772 := PrimCons(symshen_4of, tmp10771)

											tmp10773 := PrimCons(symshen_4result, tmp10772)

											tmp10774 := PrimCons(symshen_4the, tmp10773)

											tmp10775 := PrimCons(symlist, Nil)

											tmp10776 := PrimCons(symshen_4non_1empty, tmp10775)

											tmp10777 := PrimCons(symshen_4a, tmp10776)

											tmp10778 := PrimCons(symis, tmp10777)

											tmp10779 := PrimCons(Z, tmp10778)

											tmp10780 := PrimHead(V1097)

											tmp10781 := PrimTail(tmp10780)

											tmp10782 := PrimHead(tmp10781)

											tmp10783 := PrimHead(tmp10782)

											tmp10784 := PrimHead(V1097)

											tmp10785 := PrimTail(tmp10784)

											tmp10786 := PrimHead(tmp10785)

											tmp10787 := PrimTail(tmp10786)

											tmp10788 := PrimHead(V1097)

											tmp10789 := PrimTail(tmp10788)

											tmp10790 := PrimTail(tmp10789)

											tmp10791 := PrimCons(tmp10787, tmp10790)

											tmp10792 := PrimCons(symshen_4mu, tmp10791)

											tmp10793 := PrimCons(Z, Nil)

											tmp10794 := PrimCons(symshen_4of, tmp10793)

											tmp10795 := PrimCons(symtail, tmp10794)

											tmp10796 := PrimCons(symshen_4the, tmp10795)

											tmp10797 := PrimCons(tmp10796, Nil)

											tmp10798 := PrimCons(tmp10792, tmp10797)

											tmp10799 := PrimCons(tmp10798, Nil)

											tmp10800 := PrimCons(tmp10783, tmp10799)

											tmp10801 := PrimCons(symshen_4mu, tmp10800)

											tmp10802 := PrimCons(Z, Nil)

											tmp10803 := PrimCons(symshen_4of, tmp10802)

											tmp10804 := PrimCons(symhead, tmp10803)

											tmp10805 := PrimCons(symshen_4the, tmp10804)

											tmp10806 := PrimCons(tmp10805, Nil)

											tmp10807 := PrimCons(tmp10801, tmp10806)

											tmp10808 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp10807, sym_7)

											tmp10809 := PrimCons(symshen_4variable, Nil)

											tmp10810 := PrimCons(symshen_4a, tmp10809)

											tmp10811 := PrimCons(symis, tmp10810)

											tmp10812 := PrimCons(Z, tmp10811)

											tmp10813 := PrimHead(V1097)

											tmp10814 := PrimTail(tmp10813)

											tmp10815 := PrimHead(tmp10814)

											tmp10816 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp10815)

											tmp10817 := PrimHead(V1097)

											tmp10818 := PrimTail(tmp10817)

											tmp10819 := PrimHead(tmp10818)

											tmp10820 := Call(__e, PrimNS2Value(symshen_4remove__modes), tmp10819)

											tmp10821 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp10820)

											tmp10822 := PrimHead(V1097)

											tmp10823 := PrimTail(tmp10822)

											tmp10824 := PrimTail(tmp10823)

											tmp10825 := PrimHead(tmp10824)

											tmp10826 := Call(__e, PrimNS2Value(symshen_4mu__reduction), tmp10825, sym_7)

											tmp10827 := PrimCons(tmp10826, Nil)

											tmp10828 := PrimCons(symin, tmp10827)

											tmp10829 := PrimCons(tmp10821, tmp10828)

											tmp10830 := PrimCons(symshen_4to, tmp10829)

											tmp10831 := PrimCons(Z, tmp10830)

											tmp10832 := PrimCons(symbind, tmp10831)

											tmp10833 := PrimCons(tmp10832, Nil)

											tmp10834 := PrimCons(symshen_4then, tmp10833)

											tmp10835 := PrimCons(symand, tmp10834)

											tmp10836 := PrimCons(tmp10816, tmp10835)

											tmp10837 := PrimCons(symin, tmp10836)

											tmp10838 := PrimCons(symshen_4variables, tmp10837)

											tmp10839 := PrimCons(symshen_4the, tmp10838)

											tmp10840 := PrimCons(symshen_4rename, tmp10839)

											tmp10841 := PrimCons(symshen_4failed_b, Nil)

											tmp10842 := PrimCons(symshen_4else, tmp10841)

											tmp10843 := PrimCons(tmp10840, tmp10842)

											tmp10844 := PrimCons(symshen_4then, tmp10843)

											tmp10845 := PrimCons(tmp10812, tmp10844)

											tmp10846 := PrimCons(symif, tmp10845)

											tmp10847 := PrimCons(tmp10846, Nil)

											tmp10848 := PrimCons(symshen_4else, tmp10847)

											tmp10849 := PrimCons(tmp10808, tmp10848)

											tmp10850 := PrimCons(symshen_4then, tmp10849)

											tmp10851 := PrimCons(tmp10779, tmp10850)

											tmp10852 := PrimCons(symif, tmp10851)

											tmp10853 := PrimCons(tmp10852, Nil)

											tmp10854 := PrimCons(symin, tmp10853)

											tmp10855 := PrimCons(tmp10774, tmp10854)

											tmp10856 := PrimCons(symshen_4be, tmp10855)

											tmp10857 := PrimCons(Z, tmp10856)

											__e.Return(PrimCons(symlet, tmp10857))
											return

										}, 1)

										tmp10858 := Call(__e, PrimNS2Value(symgensym), symV)

										__e.TailApply(tmp10769, tmp10858)
										return

									} else {
										__e.Return(V1097)
										return
									}

								}

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp11433 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mu__reduction, tmp10760)

	_ = tmp11433

	tmp11434 := MakeNative(func(__e *ControlFlow) {
		V1100 := __e.Get(1)
		_ = V1100
		tmp11442 := PrimIsPair(V1100)

		if True == tmp11442 {
			tmp11436 := PrimHead(V1100)

			tmp11437 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp11436)

			tmp11438 := PrimTail(V1100)

			tmp11439 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp11438)

			tmp11440 := PrimCons(tmp11439, Nil)

			tmp11441 := PrimCons(tmp11437, tmp11440)

			__e.Return(PrimCons(symcons, tmp11441))
			return

		} else {
			__e.Return(V1100)
			return
		}

	}, 1)

	tmp11443 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rcons__form, tmp11434)

	_ = tmp11443

	tmp11444 := MakeNative(func(__e *ControlFlow) {
		V1102 := __e.Get(1)
		_ = V1102
		tmp11508 := PrimIsPair(V1102)

		var ifres11483 Obj

		if True == tmp11508 {
			tmp11506 := PrimHead(V1102)

			tmp11507 := PrimEqual(symmode, tmp11506)

			var ifres11485 Obj

			if True == tmp11507 {
				tmp11504 := PrimTail(V1102)

				tmp11505 := PrimIsPair(tmp11504)

				var ifres11487 Obj

				if True == tmp11505 {
					tmp11501 := PrimTail(V1102)

					tmp11502 := PrimTail(tmp11501)

					tmp11503 := PrimIsPair(tmp11502)

					var ifres11489 Obj

					if True == tmp11503 {
						tmp11497 := PrimTail(V1102)

						tmp11498 := PrimTail(tmp11497)

						tmp11499 := PrimHead(tmp11498)

						tmp11500 := PrimEqual(sym_7, tmp11499)

						var ifres11491 Obj

						if True == tmp11500 {
							tmp11493 := PrimTail(V1102)

							tmp11494 := PrimTail(tmp11493)

							tmp11495 := PrimTail(tmp11494)

							tmp11496 := PrimEqual(Nil, tmp11495)

							var ifres11492 Obj

							if True == tmp11496 {
								ifres11492 = True

							} else {
								ifres11492 = False

							}

							ifres11491 = ifres11492

						} else {
							ifres11491 = False

						}

						var ifres11490 Obj

						if True == ifres11491 {
							ifres11490 = True

						} else {
							ifres11490 = False

						}

						ifres11489 = ifres11490

					} else {
						ifres11489 = False

					}

					var ifres11488 Obj

					if True == ifres11489 {
						ifres11488 = True

					} else {
						ifres11488 = False

					}

					ifres11487 = ifres11488

				} else {
					ifres11487 = False

				}

				var ifres11486 Obj

				if True == ifres11487 {
					ifres11486 = True

				} else {
					ifres11486 = False

				}

				ifres11485 = ifres11486

			} else {
				ifres11485 = False

			}

			var ifres11484 Obj

			if True == ifres11485 {
				ifres11484 = True

			} else {
				ifres11484 = False

			}

			ifres11483 = ifres11484

		} else {
			ifres11483 = False

		}

		if True == ifres11483 {
			tmp11481 := PrimTail(V1102)

			tmp11482 := PrimHead(tmp11481)

			__e.TailApply(PrimNS2Value(symshen_4remove__modes), tmp11482)
			return

		} else {
			tmp11480 := PrimIsPair(V1102)

			var ifres11455 Obj

			if True == tmp11480 {
				tmp11478 := PrimHead(V1102)

				tmp11479 := PrimEqual(symmode, tmp11478)

				var ifres11457 Obj

				if True == tmp11479 {
					tmp11476 := PrimTail(V1102)

					tmp11477 := PrimIsPair(tmp11476)

					var ifres11459 Obj

					if True == tmp11477 {
						tmp11473 := PrimTail(V1102)

						tmp11474 := PrimTail(tmp11473)

						tmp11475 := PrimIsPair(tmp11474)

						var ifres11461 Obj

						if True == tmp11475 {
							tmp11469 := PrimTail(V1102)

							tmp11470 := PrimTail(tmp11469)

							tmp11471 := PrimHead(tmp11470)

							tmp11472 := PrimEqual(sym_1, tmp11471)

							var ifres11463 Obj

							if True == tmp11472 {
								tmp11465 := PrimTail(V1102)

								tmp11466 := PrimTail(tmp11465)

								tmp11467 := PrimTail(tmp11466)

								tmp11468 := PrimEqual(Nil, tmp11467)

								var ifres11464 Obj

								if True == tmp11468 {
									ifres11464 = True

								} else {
									ifres11464 = False

								}

								ifres11463 = ifres11464

							} else {
								ifres11463 = False

							}

							var ifres11462 Obj

							if True == ifres11463 {
								ifres11462 = True

							} else {
								ifres11462 = False

							}

							ifres11461 = ifres11462

						} else {
							ifres11461 = False

						}

						var ifres11460 Obj

						if True == ifres11461 {
							ifres11460 = True

						} else {
							ifres11460 = False

						}

						ifres11459 = ifres11460

					} else {
						ifres11459 = False

					}

					var ifres11458 Obj

					if True == ifres11459 {
						ifres11458 = True

					} else {
						ifres11458 = False

					}

					ifres11457 = ifres11458

				} else {
					ifres11457 = False

				}

				var ifres11456 Obj

				if True == ifres11457 {
					ifres11456 = True

				} else {
					ifres11456 = False

				}

				ifres11455 = ifres11456

			} else {
				ifres11455 = False

			}

			if True == ifres11455 {
				tmp11453 := PrimTail(V1102)

				tmp11454 := PrimHead(tmp11453)

				__e.TailApply(PrimNS2Value(symshen_4remove__modes), tmp11454)
				return

			} else {
				tmp11452 := PrimIsPair(V1102)

				if True == tmp11452 {
					tmp11448 := PrimHead(V1102)

					tmp11449 := Call(__e, PrimNS2Value(symshen_4remove__modes), tmp11448)

					tmp11450 := PrimTail(V1102)

					tmp11451 := Call(__e, PrimNS2Value(symshen_4remove__modes), tmp11450)

					__e.Return(PrimCons(tmp11449, tmp11451))
					return

				} else {
					__e.Return(V1102)
					return
				}

			}

		}

	}, 1)

	tmp11509 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove__modes, tmp11444)

	_ = tmp11509

	tmp11510 := MakeNative(func(__e *ControlFlow) {
		V1105 := __e.Get(1)
		_ = V1105
		V1106 := __e.Get(2)
		_ = V1106
		tmp11514 := PrimIsVariable(V1105)

		if True == tmp11514 {
			tmp11513 := PrimIsVariable(V1106)

			if True == tmp11513 {
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

	}, 2)

	tmp11515 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ephemeral__variable_2, tmp11510)

	_ = tmp11515

	tmp11516 := MakeNative(func(__e *ControlFlow) {
		V1116 := __e.Get(1)
		_ = V1116
		tmp11518 := PrimIsPair(V1116)

		if True == tmp11518 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp11519 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog__constant_2, tmp11516)

	_ = tmp11519

	tmp11520 := MakeNative(func(__e *ControlFlow) {
		V1118 := __e.Get(1)
		_ = V1118
		tmp12417 := PrimIsPair(V1118)

		var ifres12360 Obj

		if True == tmp12417 {
			tmp12415 := PrimHead(V1118)

			tmp12416 := PrimEqual(symlet, tmp12415)

			var ifres12362 Obj

			if True == tmp12416 {
				tmp12413 := PrimTail(V1118)

				tmp12414 := PrimIsPair(tmp12413)

				var ifres12364 Obj

				if True == tmp12414 {
					tmp12410 := PrimTail(V1118)

					tmp12411 := PrimTail(tmp12410)

					tmp12412 := PrimIsPair(tmp12411)

					var ifres12366 Obj

					if True == tmp12412 {
						tmp12406 := PrimTail(V1118)

						tmp12407 := PrimTail(tmp12406)

						tmp12408 := PrimHead(tmp12407)

						tmp12409 := PrimEqual(symshen_4be, tmp12408)

						var ifres12368 Obj

						if True == tmp12409 {
							tmp12402 := PrimTail(V1118)

							tmp12403 := PrimTail(tmp12402)

							tmp12404 := PrimTail(tmp12403)

							tmp12405 := PrimIsPair(tmp12404)

							var ifres12370 Obj

							if True == tmp12405 {
								tmp12397 := PrimTail(V1118)

								tmp12398 := PrimTail(tmp12397)

								tmp12399 := PrimTail(tmp12398)

								tmp12400 := PrimTail(tmp12399)

								tmp12401 := PrimIsPair(tmp12400)

								var ifres12372 Obj

								if True == tmp12401 {
									tmp12391 := PrimTail(V1118)

									tmp12392 := PrimTail(tmp12391)

									tmp12393 := PrimTail(tmp12392)

									tmp12394 := PrimTail(tmp12393)

									tmp12395 := PrimHead(tmp12394)

									tmp12396 := PrimEqual(symin, tmp12395)

									var ifres12374 Obj

									if True == tmp12396 {
										tmp12385 := PrimTail(V1118)

										tmp12386 := PrimTail(tmp12385)

										tmp12387 := PrimTail(tmp12386)

										tmp12388 := PrimTail(tmp12387)

										tmp12389 := PrimTail(tmp12388)

										tmp12390 := PrimIsPair(tmp12389)

										var ifres12376 Obj

										if True == tmp12390 {
											tmp12378 := PrimTail(V1118)

											tmp12379 := PrimTail(tmp12378)

											tmp12380 := PrimTail(tmp12379)

											tmp12381 := PrimTail(tmp12380)

											tmp12382 := PrimTail(tmp12381)

											tmp12383 := PrimTail(tmp12382)

											tmp12384 := PrimEqual(Nil, tmp12383)

											var ifres12377 Obj

											if True == tmp12384 {
												ifres12377 = True

											} else {
												ifres12377 = False

											}

											ifres12376 = ifres12377

										} else {
											ifres12376 = False

										}

										var ifres12375 Obj

										if True == ifres12376 {
											ifres12375 = True

										} else {
											ifres12375 = False

										}

										ifres12374 = ifres12375

									} else {
										ifres12374 = False

									}

									var ifres12373 Obj

									if True == ifres12374 {
										ifres12373 = True

									} else {
										ifres12373 = False

									}

									ifres12372 = ifres12373

								} else {
									ifres12372 = False

								}

								var ifres12371 Obj

								if True == ifres12372 {
									ifres12371 = True

								} else {
									ifres12371 = False

								}

								ifres12370 = ifres12371

							} else {
								ifres12370 = False

							}

							var ifres12369 Obj

							if True == ifres12370 {
								ifres12369 = True

							} else {
								ifres12369 = False

							}

							ifres12368 = ifres12369

						} else {
							ifres12368 = False

						}

						var ifres12367 Obj

						if True == ifres12368 {
							ifres12367 = True

						} else {
							ifres12367 = False

						}

						ifres12366 = ifres12367

					} else {
						ifres12366 = False

					}

					var ifres12365 Obj

					if True == ifres12366 {
						ifres12365 = True

					} else {
						ifres12365 = False

					}

					ifres12364 = ifres12365

				} else {
					ifres12364 = False

				}

				var ifres12363 Obj

				if True == ifres12364 {
					ifres12363 = True

				} else {
					ifres12363 = False

				}

				ifres12362 = ifres12363

			} else {
				ifres12362 = False

			}

			var ifres12361 Obj

			if True == ifres12362 {
				ifres12361 = True

			} else {
				ifres12361 = False

			}

			ifres12360 = ifres12361

		} else {
			ifres12360 = False

		}

		if True == ifres12360 {
			tmp12343 := PrimTail(V1118)

			tmp12344 := PrimHead(tmp12343)

			tmp12345 := PrimTail(V1118)

			tmp12346 := PrimTail(tmp12345)

			tmp12347 := PrimTail(tmp12346)

			tmp12348 := PrimHead(tmp12347)

			tmp12349 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12348)

			tmp12350 := PrimTail(V1118)

			tmp12351 := PrimTail(tmp12350)

			tmp12352 := PrimTail(tmp12351)

			tmp12353 := PrimTail(tmp12352)

			tmp12354 := PrimTail(tmp12353)

			tmp12355 := PrimHead(tmp12354)

			tmp12356 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12355)

			tmp12357 := PrimCons(tmp12356, Nil)

			tmp12358 := PrimCons(tmp12349, tmp12357)

			tmp12359 := PrimCons(tmp12344, tmp12358)

			__e.Return(PrimCons(symlet, tmp12359))
			return

		} else {
			tmp12342 := PrimIsPair(V1118)

			var ifres12290 Obj

			if True == tmp12342 {
				tmp12340 := PrimHead(V1118)

				tmp12341 := PrimEqual(symshen_4the, tmp12340)

				var ifres12292 Obj

				if True == tmp12341 {
					tmp12338 := PrimTail(V1118)

					tmp12339 := PrimIsPair(tmp12338)

					var ifres12294 Obj

					if True == tmp12339 {
						tmp12335 := PrimTail(V1118)

						tmp12336 := PrimHead(tmp12335)

						tmp12337 := PrimEqual(symshen_4result, tmp12336)

						var ifres12296 Obj

						if True == tmp12337 {
							tmp12332 := PrimTail(V1118)

							tmp12333 := PrimTail(tmp12332)

							tmp12334 := PrimIsPair(tmp12333)

							var ifres12298 Obj

							if True == tmp12334 {
								tmp12328 := PrimTail(V1118)

								tmp12329 := PrimTail(tmp12328)

								tmp12330 := PrimHead(tmp12329)

								tmp12331 := PrimEqual(symshen_4of, tmp12330)

								var ifres12300 Obj

								if True == tmp12331 {
									tmp12324 := PrimTail(V1118)

									tmp12325 := PrimTail(tmp12324)

									tmp12326 := PrimTail(tmp12325)

									tmp12327 := PrimIsPair(tmp12326)

									var ifres12302 Obj

									if True == tmp12327 {
										tmp12319 := PrimTail(V1118)

										tmp12320 := PrimTail(tmp12319)

										tmp12321 := PrimTail(tmp12320)

										tmp12322 := PrimHead(tmp12321)

										tmp12323 := PrimEqual(symshen_4dereferencing, tmp12322)

										var ifres12304 Obj

										if True == tmp12323 {
											tmp12314 := PrimTail(V1118)

											tmp12315 := PrimTail(tmp12314)

											tmp12316 := PrimTail(tmp12315)

											tmp12317 := PrimTail(tmp12316)

											tmp12318 := PrimIsPair(tmp12317)

											var ifres12306 Obj

											if True == tmp12318 {
												tmp12308 := PrimTail(V1118)

												tmp12309 := PrimTail(tmp12308)

												tmp12310 := PrimTail(tmp12309)

												tmp12311 := PrimTail(tmp12310)

												tmp12312 := PrimTail(tmp12311)

												tmp12313 := PrimEqual(Nil, tmp12312)

												var ifres12307 Obj

												if True == tmp12313 {
													ifres12307 = True

												} else {
													ifres12307 = False

												}

												ifres12306 = ifres12307

											} else {
												ifres12306 = False

											}

											var ifres12305 Obj

											if True == ifres12306 {
												ifres12305 = True

											} else {
												ifres12305 = False

											}

											ifres12304 = ifres12305

										} else {
											ifres12304 = False

										}

										var ifres12303 Obj

										if True == ifres12304 {
											ifres12303 = True

										} else {
											ifres12303 = False

										}

										ifres12302 = ifres12303

									} else {
										ifres12302 = False

									}

									var ifres12301 Obj

									if True == ifres12302 {
										ifres12301 = True

									} else {
										ifres12301 = False

									}

									ifres12300 = ifres12301

								} else {
									ifres12300 = False

								}

								var ifres12299 Obj

								if True == ifres12300 {
									ifres12299 = True

								} else {
									ifres12299 = False

								}

								ifres12298 = ifres12299

							} else {
								ifres12298 = False

							}

							var ifres12297 Obj

							if True == ifres12298 {
								ifres12297 = True

							} else {
								ifres12297 = False

							}

							ifres12296 = ifres12297

						} else {
							ifres12296 = False

						}

						var ifres12295 Obj

						if True == ifres12296 {
							ifres12295 = True

						} else {
							ifres12295 = False

						}

						ifres12294 = ifres12295

					} else {
						ifres12294 = False

					}

					var ifres12293 Obj

					if True == ifres12294 {
						ifres12293 = True

					} else {
						ifres12293 = False

					}

					ifres12292 = ifres12293

				} else {
					ifres12292 = False

				}

				var ifres12291 Obj

				if True == ifres12292 {
					ifres12291 = True

				} else {
					ifres12291 = False

				}

				ifres12290 = ifres12291

			} else {
				ifres12290 = False

			}

			if True == ifres12290 {
				tmp12282 := PrimTail(V1118)

				tmp12283 := PrimTail(tmp12282)

				tmp12284 := PrimTail(tmp12283)

				tmp12285 := PrimTail(tmp12284)

				tmp12286 := PrimHead(tmp12285)

				tmp12287 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12286)

				tmp12288 := PrimCons(symProcessN, Nil)

				tmp12289 := PrimCons(tmp12287, tmp12288)

				__e.Return(PrimCons(symshen_4lazyderef, tmp12289))
				return

			} else {
				tmp12281 := PrimIsPair(V1118)

				var ifres12224 Obj

				if True == tmp12281 {
					tmp12279 := PrimHead(V1118)

					tmp12280 := PrimEqual(symif, tmp12279)

					var ifres12226 Obj

					if True == tmp12280 {
						tmp12277 := PrimTail(V1118)

						tmp12278 := PrimIsPair(tmp12277)

						var ifres12228 Obj

						if True == tmp12278 {
							tmp12274 := PrimTail(V1118)

							tmp12275 := PrimTail(tmp12274)

							tmp12276 := PrimIsPair(tmp12275)

							var ifres12230 Obj

							if True == tmp12276 {
								tmp12270 := PrimTail(V1118)

								tmp12271 := PrimTail(tmp12270)

								tmp12272 := PrimHead(tmp12271)

								tmp12273 := PrimEqual(symshen_4then, tmp12272)

								var ifres12232 Obj

								if True == tmp12273 {
									tmp12266 := PrimTail(V1118)

									tmp12267 := PrimTail(tmp12266)

									tmp12268 := PrimTail(tmp12267)

									tmp12269 := PrimIsPair(tmp12268)

									var ifres12234 Obj

									if True == tmp12269 {
										tmp12261 := PrimTail(V1118)

										tmp12262 := PrimTail(tmp12261)

										tmp12263 := PrimTail(tmp12262)

										tmp12264 := PrimTail(tmp12263)

										tmp12265 := PrimIsPair(tmp12264)

										var ifres12236 Obj

										if True == tmp12265 {
											tmp12255 := PrimTail(V1118)

											tmp12256 := PrimTail(tmp12255)

											tmp12257 := PrimTail(tmp12256)

											tmp12258 := PrimTail(tmp12257)

											tmp12259 := PrimHead(tmp12258)

											tmp12260 := PrimEqual(symshen_4else, tmp12259)

											var ifres12238 Obj

											if True == tmp12260 {
												tmp12249 := PrimTail(V1118)

												tmp12250 := PrimTail(tmp12249)

												tmp12251 := PrimTail(tmp12250)

												tmp12252 := PrimTail(tmp12251)

												tmp12253 := PrimTail(tmp12252)

												tmp12254 := PrimIsPair(tmp12253)

												var ifres12240 Obj

												if True == tmp12254 {
													tmp12242 := PrimTail(V1118)

													tmp12243 := PrimTail(tmp12242)

													tmp12244 := PrimTail(tmp12243)

													tmp12245 := PrimTail(tmp12244)

													tmp12246 := PrimTail(tmp12245)

													tmp12247 := PrimTail(tmp12246)

													tmp12248 := PrimEqual(Nil, tmp12247)

													var ifres12241 Obj

													if True == tmp12248 {
														ifres12241 = True

													} else {
														ifres12241 = False

													}

													ifres12240 = ifres12241

												} else {
													ifres12240 = False

												}

												var ifres12239 Obj

												if True == ifres12240 {
													ifres12239 = True

												} else {
													ifres12239 = False

												}

												ifres12238 = ifres12239

											} else {
												ifres12238 = False

											}

											var ifres12237 Obj

											if True == ifres12238 {
												ifres12237 = True

											} else {
												ifres12237 = False

											}

											ifres12236 = ifres12237

										} else {
											ifres12236 = False

										}

										var ifres12235 Obj

										if True == ifres12236 {
											ifres12235 = True

										} else {
											ifres12235 = False

										}

										ifres12234 = ifres12235

									} else {
										ifres12234 = False

									}

									var ifres12233 Obj

									if True == ifres12234 {
										ifres12233 = True

									} else {
										ifres12233 = False

									}

									ifres12232 = ifres12233

								} else {
									ifres12232 = False

								}

								var ifres12231 Obj

								if True == ifres12232 {
									ifres12231 = True

								} else {
									ifres12231 = False

								}

								ifres12230 = ifres12231

							} else {
								ifres12230 = False

							}

							var ifres12229 Obj

							if True == ifres12230 {
								ifres12229 = True

							} else {
								ifres12229 = False

							}

							ifres12228 = ifres12229

						} else {
							ifres12228 = False

						}

						var ifres12227 Obj

						if True == ifres12228 {
							ifres12227 = True

						} else {
							ifres12227 = False

						}

						ifres12226 = ifres12227

					} else {
						ifres12226 = False

					}

					var ifres12225 Obj

					if True == ifres12226 {
						ifres12225 = True

					} else {
						ifres12225 = False

					}

					ifres12224 = ifres12225

				} else {
					ifres12224 = False

				}

				if True == ifres12224 {
					tmp12206 := PrimTail(V1118)

					tmp12207 := PrimHead(tmp12206)

					tmp12208 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12207)

					tmp12209 := PrimTail(V1118)

					tmp12210 := PrimTail(tmp12209)

					tmp12211 := PrimTail(tmp12210)

					tmp12212 := PrimHead(tmp12211)

					tmp12213 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12212)

					tmp12214 := PrimTail(V1118)

					tmp12215 := PrimTail(tmp12214)

					tmp12216 := PrimTail(tmp12215)

					tmp12217 := PrimTail(tmp12216)

					tmp12218 := PrimTail(tmp12217)

					tmp12219 := PrimHead(tmp12218)

					tmp12220 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp12219)

					tmp12221 := PrimCons(tmp12220, Nil)

					tmp12222 := PrimCons(tmp12213, tmp12221)

					tmp12223 := PrimCons(tmp12208, tmp12222)

					__e.Return(PrimCons(symif, tmp12223))
					return

				} else {
					tmp12205 := PrimIsPair(V1118)

					var ifres12165 Obj

					if True == tmp12205 {
						tmp12203 := PrimTail(V1118)

						tmp12204 := PrimIsPair(tmp12203)

						var ifres12167 Obj

						if True == tmp12204 {
							tmp12200 := PrimTail(V1118)

							tmp12201 := PrimHead(tmp12200)

							tmp12202 := PrimEqual(symis, tmp12201)

							var ifres12169 Obj

							if True == tmp12202 {
								tmp12197 := PrimTail(V1118)

								tmp12198 := PrimTail(tmp12197)

								tmp12199 := PrimIsPair(tmp12198)

								var ifres12171 Obj

								if True == tmp12199 {
									tmp12193 := PrimTail(V1118)

									tmp12194 := PrimTail(tmp12193)

									tmp12195 := PrimHead(tmp12194)

									tmp12196 := PrimEqual(symshen_4a, tmp12195)

									var ifres12173 Obj

									if True == tmp12196 {
										tmp12189 := PrimTail(V1118)

										tmp12190 := PrimTail(tmp12189)

										tmp12191 := PrimTail(tmp12190)

										tmp12192 := PrimIsPair(tmp12191)

										var ifres12175 Obj

										if True == tmp12192 {
											tmp12184 := PrimTail(V1118)

											tmp12185 := PrimTail(tmp12184)

											tmp12186 := PrimTail(tmp12185)

											tmp12187 := PrimHead(tmp12186)

											tmp12188 := PrimEqual(symshen_4variable, tmp12187)

											var ifres12177 Obj

											if True == tmp12188 {
												tmp12179 := PrimTail(V1118)

												tmp12180 := PrimTail(tmp12179)

												tmp12181 := PrimTail(tmp12180)

												tmp12182 := PrimTail(tmp12181)

												tmp12183 := PrimEqual(Nil, tmp12182)

												var ifres12178 Obj

												if True == tmp12183 {
													ifres12178 = True

												} else {
													ifres12178 = False

												}

												ifres12177 = ifres12178

											} else {
												ifres12177 = False

											}

											var ifres12176 Obj

											if True == ifres12177 {
												ifres12176 = True

											} else {
												ifres12176 = False

											}

											ifres12175 = ifres12176

										} else {
											ifres12175 = False

										}

										var ifres12174 Obj

										if True == ifres12175 {
											ifres12174 = True

										} else {
											ifres12174 = False

										}

										ifres12173 = ifres12174

									} else {
										ifres12173 = False

									}

									var ifres12172 Obj

									if True == ifres12173 {
										ifres12172 = True

									} else {
										ifres12172 = False

									}

									ifres12171 = ifres12172

								} else {
									ifres12171 = False

								}

								var ifres12170 Obj

								if True == ifres12171 {
									ifres12170 = True

								} else {
									ifres12170 = False

								}

								ifres12169 = ifres12170

							} else {
								ifres12169 = False

							}

							var ifres12168 Obj

							if True == ifres12169 {
								ifres12168 = True

							} else {
								ifres12168 = False

							}

							ifres12167 = ifres12168

						} else {
							ifres12167 = False

						}

						var ifres12166 Obj

						if True == ifres12167 {
							ifres12166 = True

						} else {
							ifres12166 = False

						}

						ifres12165 = ifres12166

					} else {
						ifres12165 = False

					}

					if True == ifres12165 {
						tmp12163 := PrimHead(V1118)

						tmp12164 := PrimCons(tmp12163, Nil)

						__e.Return(PrimCons(symshen_4pvar_2, tmp12164))
						return

					} else {
						tmp12162 := PrimIsPair(V1118)

						var ifres12106 Obj

						if True == tmp12162 {
							tmp12160 := PrimTail(V1118)

							tmp12161 := PrimIsPair(tmp12160)

							var ifres12108 Obj

							if True == tmp12161 {
								tmp12157 := PrimTail(V1118)

								tmp12158 := PrimHead(tmp12157)

								tmp12159 := PrimEqual(symis, tmp12158)

								var ifres12110 Obj

								if True == tmp12159 {
									tmp12154 := PrimTail(V1118)

									tmp12155 := PrimTail(tmp12154)

									tmp12156 := PrimIsPair(tmp12155)

									var ifres12112 Obj

									if True == tmp12156 {
										tmp12150 := PrimTail(V1118)

										tmp12151 := PrimTail(tmp12150)

										tmp12152 := PrimHead(tmp12151)

										tmp12153 := PrimEqual(symshen_4a, tmp12152)

										var ifres12114 Obj

										if True == tmp12153 {
											tmp12146 := PrimTail(V1118)

											tmp12147 := PrimTail(tmp12146)

											tmp12148 := PrimTail(tmp12147)

											tmp12149 := PrimIsPair(tmp12148)

											var ifres12116 Obj

											if True == tmp12149 {
												tmp12141 := PrimTail(V1118)

												tmp12142 := PrimTail(tmp12141)

												tmp12143 := PrimTail(tmp12142)

												tmp12144 := PrimHead(tmp12143)

												tmp12145 := PrimEqual(symshen_4non_1empty, tmp12144)

												var ifres12118 Obj

												if True == tmp12145 {
													tmp12136 := PrimTail(V1118)

													tmp12137 := PrimTail(tmp12136)

													tmp12138 := PrimTail(tmp12137)

													tmp12139 := PrimTail(tmp12138)

													tmp12140 := PrimIsPair(tmp12139)

													var ifres12120 Obj

													if True == tmp12140 {
														tmp12130 := PrimTail(V1118)

														tmp12131 := PrimTail(tmp12130)

														tmp12132 := PrimTail(tmp12131)

														tmp12133 := PrimTail(tmp12132)

														tmp12134 := PrimHead(tmp12133)

														tmp12135 := PrimEqual(symlist, tmp12134)

														var ifres12122 Obj

														if True == tmp12135 {
															tmp12124 := PrimTail(V1118)

															tmp12125 := PrimTail(tmp12124)

															tmp12126 := PrimTail(tmp12125)

															tmp12127 := PrimTail(tmp12126)

															tmp12128 := PrimTail(tmp12127)

															tmp12129 := PrimEqual(Nil, tmp12128)

															var ifres12123 Obj

															if True == tmp12129 {
																ifres12123 = True

															} else {
																ifres12123 = False

															}

															ifres12122 = ifres12123

														} else {
															ifres12122 = False

														}

														var ifres12121 Obj

														if True == ifres12122 {
															ifres12121 = True

														} else {
															ifres12121 = False

														}

														ifres12120 = ifres12121

													} else {
														ifres12120 = False

													}

													var ifres12119 Obj

													if True == ifres12120 {
														ifres12119 = True

													} else {
														ifres12119 = False

													}

													ifres12118 = ifres12119

												} else {
													ifres12118 = False

												}

												var ifres12117 Obj

												if True == ifres12118 {
													ifres12117 = True

												} else {
													ifres12117 = False

												}

												ifres12116 = ifres12117

											} else {
												ifres12116 = False

											}

											var ifres12115 Obj

											if True == ifres12116 {
												ifres12115 = True

											} else {
												ifres12115 = False

											}

											ifres12114 = ifres12115

										} else {
											ifres12114 = False

										}

										var ifres12113 Obj

										if True == ifres12114 {
											ifres12113 = True

										} else {
											ifres12113 = False

										}

										ifres12112 = ifres12113

									} else {
										ifres12112 = False

									}

									var ifres12111 Obj

									if True == ifres12112 {
										ifres12111 = True

									} else {
										ifres12111 = False

									}

									ifres12110 = ifres12111

								} else {
									ifres12110 = False

								}

								var ifres12109 Obj

								if True == ifres12110 {
									ifres12109 = True

								} else {
									ifres12109 = False

								}

								ifres12108 = ifres12109

							} else {
								ifres12108 = False

							}

							var ifres12107 Obj

							if True == ifres12108 {
								ifres12107 = True

							} else {
								ifres12107 = False

							}

							ifres12106 = ifres12107

						} else {
							ifres12106 = False

						}

						if True == ifres12106 {
							tmp12104 := PrimHead(V1118)

							tmp12105 := PrimCons(tmp12104, Nil)

							__e.Return(PrimCons(symcons_2, tmp12105))
							return

						} else {
							tmp12103 := PrimIsPair(V1118)

							var ifres11994 Obj

							if True == tmp12103 {
								tmp12101 := PrimHead(V1118)

								tmp12102 := PrimEqual(symshen_4rename, tmp12101)

								var ifres11996 Obj

								if True == tmp12102 {
									tmp12099 := PrimTail(V1118)

									tmp12100 := PrimIsPair(tmp12099)

									var ifres11998 Obj

									if True == tmp12100 {
										tmp12096 := PrimTail(V1118)

										tmp12097 := PrimHead(tmp12096)

										tmp12098 := PrimEqual(symshen_4the, tmp12097)

										var ifres12000 Obj

										if True == tmp12098 {
											tmp12093 := PrimTail(V1118)

											tmp12094 := PrimTail(tmp12093)

											tmp12095 := PrimIsPair(tmp12094)

											var ifres12002 Obj

											if True == tmp12095 {
												tmp12089 := PrimTail(V1118)

												tmp12090 := PrimTail(tmp12089)

												tmp12091 := PrimHead(tmp12090)

												tmp12092 := PrimEqual(symshen_4variables, tmp12091)

												var ifres12004 Obj

												if True == tmp12092 {
													tmp12085 := PrimTail(V1118)

													tmp12086 := PrimTail(tmp12085)

													tmp12087 := PrimTail(tmp12086)

													tmp12088 := PrimIsPair(tmp12087)

													var ifres12006 Obj

													if True == tmp12088 {
														tmp12080 := PrimTail(V1118)

														tmp12081 := PrimTail(tmp12080)

														tmp12082 := PrimTail(tmp12081)

														tmp12083 := PrimHead(tmp12082)

														tmp12084 := PrimEqual(symin, tmp12083)

														var ifres12008 Obj

														if True == tmp12084 {
															tmp12075 := PrimTail(V1118)

															tmp12076 := PrimTail(tmp12075)

															tmp12077 := PrimTail(tmp12076)

															tmp12078 := PrimTail(tmp12077)

															tmp12079 := PrimIsPair(tmp12078)

															var ifres12010 Obj

															if True == tmp12079 {
																tmp12069 := PrimTail(V1118)

																tmp12070 := PrimTail(tmp12069)

																tmp12071 := PrimTail(tmp12070)

																tmp12072 := PrimTail(tmp12071)

																tmp12073 := PrimHead(tmp12072)

																tmp12074 := PrimEqual(Nil, tmp12073)

																var ifres12012 Obj

																if True == tmp12074 {
																	tmp12063 := PrimTail(V1118)

																	tmp12064 := PrimTail(tmp12063)

																	tmp12065 := PrimTail(tmp12064)

																	tmp12066 := PrimTail(tmp12065)

																	tmp12067 := PrimTail(tmp12066)

																	tmp12068 := PrimIsPair(tmp12067)

																	var ifres12014 Obj

																	if True == tmp12068 {
																		tmp12056 := PrimTail(V1118)

																		tmp12057 := PrimTail(tmp12056)

																		tmp12058 := PrimTail(tmp12057)

																		tmp12059 := PrimTail(tmp12058)

																		tmp12060 := PrimTail(tmp12059)

																		tmp12061 := PrimHead(tmp12060)

																		tmp12062 := PrimEqual(symand, tmp12061)

																		var ifres12016 Obj

																		if True == tmp12062 {
																			tmp12049 := PrimTail(V1118)

																			tmp12050 := PrimTail(tmp12049)

																			tmp12051 := PrimTail(tmp12050)

																			tmp12052 := PrimTail(tmp12051)

																			tmp12053 := PrimTail(tmp12052)

																			tmp12054 := PrimTail(tmp12053)

																			tmp12055 := PrimIsPair(tmp12054)

																			var ifres12018 Obj

																			if True == tmp12055 {
																				tmp12041 := PrimTail(V1118)

																				tmp12042 := PrimTail(tmp12041)

																				tmp12043 := PrimTail(tmp12042)

																				tmp12044 := PrimTail(tmp12043)

																				tmp12045 := PrimTail(tmp12044)

																				tmp12046 := PrimTail(tmp12045)

																				tmp12047 := PrimHead(tmp12046)

																				tmp12048 := PrimEqual(symshen_4then, tmp12047)

																				var ifres12020 Obj

																				if True == tmp12048 {
																					tmp12033 := PrimTail(V1118)

																					tmp12034 := PrimTail(tmp12033)

																					tmp12035 := PrimTail(tmp12034)

																					tmp12036 := PrimTail(tmp12035)

																					tmp12037 := PrimTail(tmp12036)

																					tmp12038 := PrimTail(tmp12037)

																					tmp12039 := PrimTail(tmp12038)

																					tmp12040 := PrimIsPair(tmp12039)

																					var ifres12022 Obj

																					if True == tmp12040 {
																						tmp12024 := PrimTail(V1118)

																						tmp12025 := PrimTail(tmp12024)

																						tmp12026 := PrimTail(tmp12025)

																						tmp12027 := PrimTail(tmp12026)

																						tmp12028 := PrimTail(tmp12027)

																						tmp12029 := PrimTail(tmp12028)

																						tmp12030 := PrimTail(tmp12029)

																						tmp12031 := PrimTail(tmp12030)

																						tmp12032 := PrimEqual(Nil, tmp12031)

																						var ifres12023 Obj

																						if True == tmp12032 {
																							ifres12023 = True

																						} else {
																							ifres12023 = False

																						}

																						ifres12022 = ifres12023

																					} else {
																						ifres12022 = False

																					}

																					var ifres12021 Obj

																					if True == ifres12022 {
																						ifres12021 = True

																					} else {
																						ifres12021 = False

																					}

																					ifres12020 = ifres12021

																				} else {
																					ifres12020 = False

																				}

																				var ifres12019 Obj

																				if True == ifres12020 {
																					ifres12019 = True

																				} else {
																					ifres12019 = False

																				}

																				ifres12018 = ifres12019

																			} else {
																				ifres12018 = False

																			}

																			var ifres12017 Obj

																			if True == ifres12018 {
																				ifres12017 = True

																			} else {
																				ifres12017 = False

																			}

																			ifres12016 = ifres12017

																		} else {
																			ifres12016 = False

																		}

																		var ifres12015 Obj

																		if True == ifres12016 {
																			ifres12015 = True

																		} else {
																			ifres12015 = False

																		}

																		ifres12014 = ifres12015

																	} else {
																		ifres12014 = False

																	}

																	var ifres12013 Obj

																	if True == ifres12014 {
																		ifres12013 = True

																	} else {
																		ifres12013 = False

																	}

																	ifres12012 = ifres12013

																} else {
																	ifres12012 = False

																}

																var ifres12011 Obj

																if True == ifres12012 {
																	ifres12011 = True

																} else {
																	ifres12011 = False

																}

																ifres12010 = ifres12011

															} else {
																ifres12010 = False

															}

															var ifres12009 Obj

															if True == ifres12010 {
																ifres12009 = True

															} else {
																ifres12009 = False

															}

															ifres12008 = ifres12009

														} else {
															ifres12008 = False

														}

														var ifres12007 Obj

														if True == ifres12008 {
															ifres12007 = True

														} else {
															ifres12007 = False

														}

														ifres12006 = ifres12007

													} else {
														ifres12006 = False

													}

													var ifres12005 Obj

													if True == ifres12006 {
														ifres12005 = True

													} else {
														ifres12005 = False

													}

													ifres12004 = ifres12005

												} else {
													ifres12004 = False

												}

												var ifres12003 Obj

												if True == ifres12004 {
													ifres12003 = True

												} else {
													ifres12003 = False

												}

												ifres12002 = ifres12003

											} else {
												ifres12002 = False

											}

											var ifres12001 Obj

											if True == ifres12002 {
												ifres12001 = True

											} else {
												ifres12001 = False

											}

											ifres12000 = ifres12001

										} else {
											ifres12000 = False

										}

										var ifres11999 Obj

										if True == ifres12000 {
											ifres11999 = True

										} else {
											ifres11999 = False

										}

										ifres11998 = ifres11999

									} else {
										ifres11998 = False

									}

									var ifres11997 Obj

									if True == ifres11998 {
										ifres11997 = True

									} else {
										ifres11997 = False

									}

									ifres11996 = ifres11997

								} else {
									ifres11996 = False

								}

								var ifres11995 Obj

								if True == ifres11996 {
									ifres11995 = True

								} else {
									ifres11995 = False

								}

								ifres11994 = ifres11995

							} else {
								ifres11994 = False

							}

							if True == ifres11994 {
								tmp11986 := PrimTail(V1118)

								tmp11987 := PrimTail(tmp11986)

								tmp11988 := PrimTail(tmp11987)

								tmp11989 := PrimTail(tmp11988)

								tmp11990 := PrimTail(tmp11989)

								tmp11991 := PrimTail(tmp11990)

								tmp11992 := PrimTail(tmp11991)

								tmp11993 := PrimHead(tmp11992)

								__e.TailApply(PrimNS2Value(symshen_4aum__to__shen), tmp11993)
								return

							} else {
								tmp11985 := PrimIsPair(V1118)

								var ifres11876 Obj

								if True == tmp11985 {
									tmp11983 := PrimHead(V1118)

									tmp11984 := PrimEqual(symshen_4rename, tmp11983)

									var ifres11878 Obj

									if True == tmp11984 {
										tmp11981 := PrimTail(V1118)

										tmp11982 := PrimIsPair(tmp11981)

										var ifres11880 Obj

										if True == tmp11982 {
											tmp11978 := PrimTail(V1118)

											tmp11979 := PrimHead(tmp11978)

											tmp11980 := PrimEqual(symshen_4the, tmp11979)

											var ifres11882 Obj

											if True == tmp11980 {
												tmp11975 := PrimTail(V1118)

												tmp11976 := PrimTail(tmp11975)

												tmp11977 := PrimIsPair(tmp11976)

												var ifres11884 Obj

												if True == tmp11977 {
													tmp11971 := PrimTail(V1118)

													tmp11972 := PrimTail(tmp11971)

													tmp11973 := PrimHead(tmp11972)

													tmp11974 := PrimEqual(symshen_4variables, tmp11973)

													var ifres11886 Obj

													if True == tmp11974 {
														tmp11967 := PrimTail(V1118)

														tmp11968 := PrimTail(tmp11967)

														tmp11969 := PrimTail(tmp11968)

														tmp11970 := PrimIsPair(tmp11969)

														var ifres11888 Obj

														if True == tmp11970 {
															tmp11962 := PrimTail(V1118)

															tmp11963 := PrimTail(tmp11962)

															tmp11964 := PrimTail(tmp11963)

															tmp11965 := PrimHead(tmp11964)

															tmp11966 := PrimEqual(symin, tmp11965)

															var ifres11890 Obj

															if True == tmp11966 {
																tmp11957 := PrimTail(V1118)

																tmp11958 := PrimTail(tmp11957)

																tmp11959 := PrimTail(tmp11958)

																tmp11960 := PrimTail(tmp11959)

																tmp11961 := PrimIsPair(tmp11960)

																var ifres11892 Obj

																if True == tmp11961 {
																	tmp11951 := PrimTail(V1118)

																	tmp11952 := PrimTail(tmp11951)

																	tmp11953 := PrimTail(tmp11952)

																	tmp11954 := PrimTail(tmp11953)

																	tmp11955 := PrimHead(tmp11954)

																	tmp11956 := PrimIsPair(tmp11955)

																	var ifres11894 Obj

																	if True == tmp11956 {
																		tmp11945 := PrimTail(V1118)

																		tmp11946 := PrimTail(tmp11945)

																		tmp11947 := PrimTail(tmp11946)

																		tmp11948 := PrimTail(tmp11947)

																		tmp11949 := PrimTail(tmp11948)

																		tmp11950 := PrimIsPair(tmp11949)

																		var ifres11896 Obj

																		if True == tmp11950 {
																			tmp11938 := PrimTail(V1118)

																			tmp11939 := PrimTail(tmp11938)

																			tmp11940 := PrimTail(tmp11939)

																			tmp11941 := PrimTail(tmp11940)

																			tmp11942 := PrimTail(tmp11941)

																			tmp11943 := PrimHead(tmp11942)

																			tmp11944 := PrimEqual(symand, tmp11943)

																			var ifres11898 Obj

																			if True == tmp11944 {
																				tmp11931 := PrimTail(V1118)

																				tmp11932 := PrimTail(tmp11931)

																				tmp11933 := PrimTail(tmp11932)

																				tmp11934 := PrimTail(tmp11933)

																				tmp11935 := PrimTail(tmp11934)

																				tmp11936 := PrimTail(tmp11935)

																				tmp11937 := PrimIsPair(tmp11936)

																				var ifres11900 Obj

																				if True == tmp11937 {
																					tmp11923 := PrimTail(V1118)

																					tmp11924 := PrimTail(tmp11923)

																					tmp11925 := PrimTail(tmp11924)

																					tmp11926 := PrimTail(tmp11925)

																					tmp11927 := PrimTail(tmp11926)

																					tmp11928 := PrimTail(tmp11927)

																					tmp11929 := PrimHead(tmp11928)

																					tmp11930 := PrimEqual(symshen_4then, tmp11929)

																					var ifres11902 Obj

																					if True == tmp11930 {
																						tmp11915 := PrimTail(V1118)

																						tmp11916 := PrimTail(tmp11915)

																						tmp11917 := PrimTail(tmp11916)

																						tmp11918 := PrimTail(tmp11917)

																						tmp11919 := PrimTail(tmp11918)

																						tmp11920 := PrimTail(tmp11919)

																						tmp11921 := PrimTail(tmp11920)

																						tmp11922 := PrimIsPair(tmp11921)

																						var ifres11904 Obj

																						if True == tmp11922 {
																							tmp11906 := PrimTail(V1118)

																							tmp11907 := PrimTail(tmp11906)

																							tmp11908 := PrimTail(tmp11907)

																							tmp11909 := PrimTail(tmp11908)

																							tmp11910 := PrimTail(tmp11909)

																							tmp11911 := PrimTail(tmp11910)

																							tmp11912 := PrimTail(tmp11911)

																							tmp11913 := PrimTail(tmp11912)

																							tmp11914 := PrimEqual(Nil, tmp11913)

																							var ifres11905 Obj

																							if True == tmp11914 {
																								ifres11905 = True

																							} else {
																								ifres11905 = False

																							}

																							ifres11904 = ifres11905

																						} else {
																							ifres11904 = False

																						}

																						var ifres11903 Obj

																						if True == ifres11904 {
																							ifres11903 = True

																						} else {
																							ifres11903 = False

																						}

																						ifres11902 = ifres11903

																					} else {
																						ifres11902 = False

																					}

																					var ifres11901 Obj

																					if True == ifres11902 {
																						ifres11901 = True

																					} else {
																						ifres11901 = False

																					}

																					ifres11900 = ifres11901

																				} else {
																					ifres11900 = False

																				}

																				var ifres11899 Obj

																				if True == ifres11900 {
																					ifres11899 = True

																				} else {
																					ifres11899 = False

																				}

																				ifres11898 = ifres11899

																			} else {
																				ifres11898 = False

																			}

																			var ifres11897 Obj

																			if True == ifres11898 {
																				ifres11897 = True

																			} else {
																				ifres11897 = False

																			}

																			ifres11896 = ifres11897

																		} else {
																			ifres11896 = False

																		}

																		var ifres11895 Obj

																		if True == ifres11896 {
																			ifres11895 = True

																		} else {
																			ifres11895 = False

																		}

																		ifres11894 = ifres11895

																	} else {
																		ifres11894 = False

																	}

																	var ifres11893 Obj

																	if True == ifres11894 {
																		ifres11893 = True

																	} else {
																		ifres11893 = False

																	}

																	ifres11892 = ifres11893

																} else {
																	ifres11892 = False

																}

																var ifres11891 Obj

																if True == ifres11892 {
																	ifres11891 = True

																} else {
																	ifres11891 = False

																}

																ifres11890 = ifres11891

															} else {
																ifres11890 = False

															}

															var ifres11889 Obj

															if True == ifres11890 {
																ifres11889 = True

															} else {
																ifres11889 = False

															}

															ifres11888 = ifres11889

														} else {
															ifres11888 = False

														}

														var ifres11887 Obj

														if True == ifres11888 {
															ifres11887 = True

														} else {
															ifres11887 = False

														}

														ifres11886 = ifres11887

													} else {
														ifres11886 = False

													}

													var ifres11885 Obj

													if True == ifres11886 {
														ifres11885 = True

													} else {
														ifres11885 = False

													}

													ifres11884 = ifres11885

												} else {
													ifres11884 = False

												}

												var ifres11883 Obj

												if True == ifres11884 {
													ifres11883 = True

												} else {
													ifres11883 = False

												}

												ifres11882 = ifres11883

											} else {
												ifres11882 = False

											}

											var ifres11881 Obj

											if True == ifres11882 {
												ifres11881 = True

											} else {
												ifres11881 = False

											}

											ifres11880 = ifres11881

										} else {
											ifres11880 = False

										}

										var ifres11879 Obj

										if True == ifres11880 {
											ifres11879 = True

										} else {
											ifres11879 = False

										}

										ifres11878 = ifres11879

									} else {
										ifres11878 = False

									}

									var ifres11877 Obj

									if True == ifres11878 {
										ifres11877 = True

									} else {
										ifres11877 = False

									}

									ifres11876 = ifres11877

								} else {
									ifres11876 = False

								}

								if True == ifres11876 {
									tmp11848 := PrimTail(V1118)

									tmp11849 := PrimTail(tmp11848)

									tmp11850 := PrimTail(tmp11849)

									tmp11851 := PrimTail(tmp11850)

									tmp11852 := PrimHead(tmp11851)

									tmp11853 := PrimHead(tmp11852)

									tmp11854 := PrimCons(symProcessN, Nil)

									tmp11855 := PrimCons(symshen_4newpv, tmp11854)

									tmp11856 := PrimTail(V1118)

									tmp11857 := PrimTail(tmp11856)

									tmp11858 := PrimTail(tmp11857)

									tmp11859 := PrimTail(tmp11858)

									tmp11860 := PrimHead(tmp11859)

									tmp11861 := PrimTail(tmp11860)

									tmp11862 := PrimTail(V1118)

									tmp11863 := PrimTail(tmp11862)

									tmp11864 := PrimTail(tmp11863)

									tmp11865 := PrimTail(tmp11864)

									tmp11866 := PrimTail(tmp11865)

									tmp11867 := PrimCons(tmp11861, tmp11866)

									tmp11868 := PrimCons(symin, tmp11867)

									tmp11869 := PrimCons(symshen_4variables, tmp11868)

									tmp11870 := PrimCons(symshen_4the, tmp11869)

									tmp11871 := PrimCons(symshen_4rename, tmp11870)

									tmp11872 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp11871)

									tmp11873 := PrimCons(tmp11872, Nil)

									tmp11874 := PrimCons(tmp11855, tmp11873)

									tmp11875 := PrimCons(tmp11853, tmp11874)

									__e.Return(PrimCons(symlet, tmp11875))
									return

								} else {
									tmp11847 := PrimIsPair(V1118)

									var ifres11790 Obj

									if True == tmp11847 {
										tmp11845 := PrimHead(V1118)

										tmp11846 := PrimEqual(symbind, tmp11845)

										var ifres11792 Obj

										if True == tmp11846 {
											tmp11843 := PrimTail(V1118)

											tmp11844 := PrimIsPair(tmp11843)

											var ifres11794 Obj

											if True == tmp11844 {
												tmp11840 := PrimTail(V1118)

												tmp11841 := PrimTail(tmp11840)

												tmp11842 := PrimIsPair(tmp11841)

												var ifres11796 Obj

												if True == tmp11842 {
													tmp11836 := PrimTail(V1118)

													tmp11837 := PrimTail(tmp11836)

													tmp11838 := PrimHead(tmp11837)

													tmp11839 := PrimEqual(symshen_4to, tmp11838)

													var ifres11798 Obj

													if True == tmp11839 {
														tmp11832 := PrimTail(V1118)

														tmp11833 := PrimTail(tmp11832)

														tmp11834 := PrimTail(tmp11833)

														tmp11835 := PrimIsPair(tmp11834)

														var ifres11800 Obj

														if True == tmp11835 {
															tmp11827 := PrimTail(V1118)

															tmp11828 := PrimTail(tmp11827)

															tmp11829 := PrimTail(tmp11828)

															tmp11830 := PrimTail(tmp11829)

															tmp11831 := PrimIsPair(tmp11830)

															var ifres11802 Obj

															if True == tmp11831 {
																tmp11821 := PrimTail(V1118)

																tmp11822 := PrimTail(tmp11821)

																tmp11823 := PrimTail(tmp11822)

																tmp11824 := PrimTail(tmp11823)

																tmp11825 := PrimHead(tmp11824)

																tmp11826 := PrimEqual(symin, tmp11825)

																var ifres11804 Obj

																if True == tmp11826 {
																	tmp11815 := PrimTail(V1118)

																	tmp11816 := PrimTail(tmp11815)

																	tmp11817 := PrimTail(tmp11816)

																	tmp11818 := PrimTail(tmp11817)

																	tmp11819 := PrimTail(tmp11818)

																	tmp11820 := PrimIsPair(tmp11819)

																	var ifres11806 Obj

																	if True == tmp11820 {
																		tmp11808 := PrimTail(V1118)

																		tmp11809 := PrimTail(tmp11808)

																		tmp11810 := PrimTail(tmp11809)

																		tmp11811 := PrimTail(tmp11810)

																		tmp11812 := PrimTail(tmp11811)

																		tmp11813 := PrimTail(tmp11812)

																		tmp11814 := PrimEqual(Nil, tmp11813)

																		var ifres11807 Obj

																		if True == tmp11814 {
																			ifres11807 = True

																		} else {
																			ifres11807 = False

																		}

																		ifres11806 = ifres11807

																	} else {
																		ifres11806 = False

																	}

																	var ifres11805 Obj

																	if True == ifres11806 {
																		ifres11805 = True

																	} else {
																		ifres11805 = False

																	}

																	ifres11804 = ifres11805

																} else {
																	ifres11804 = False

																}

																var ifres11803 Obj

																if True == ifres11804 {
																	ifres11803 = True

																} else {
																	ifres11803 = False

																}

																ifres11802 = ifres11803

															} else {
																ifres11802 = False

															}

															var ifres11801 Obj

															if True == ifres11802 {
																ifres11801 = True

															} else {
																ifres11801 = False

															}

															ifres11800 = ifres11801

														} else {
															ifres11800 = False

														}

														var ifres11799 Obj

														if True == ifres11800 {
															ifres11799 = True

														} else {
															ifres11799 = False

														}

														ifres11798 = ifres11799

													} else {
														ifres11798 = False

													}

													var ifres11797 Obj

													if True == ifres11798 {
														ifres11797 = True

													} else {
														ifres11797 = False

													}

													ifres11796 = ifres11797

												} else {
													ifres11796 = False

												}

												var ifres11795 Obj

												if True == ifres11796 {
													ifres11795 = True

												} else {
													ifres11795 = False

												}

												ifres11794 = ifres11795

											} else {
												ifres11794 = False

											}

											var ifres11793 Obj

											if True == ifres11794 {
												ifres11793 = True

											} else {
												ifres11793 = False

											}

											ifres11792 = ifres11793

										} else {
											ifres11792 = False

										}

										var ifres11791 Obj

										if True == ifres11792 {
											ifres11791 = True

										} else {
											ifres11791 = False

										}

										ifres11790 = ifres11791

									} else {
										ifres11790 = False

									}

									if True == ifres11790 {
										tmp11758 := PrimTail(V1118)

										tmp11759 := PrimHead(tmp11758)

										tmp11760 := PrimTail(V1118)

										tmp11761 := PrimTail(tmp11760)

										tmp11762 := PrimTail(tmp11761)

										tmp11763 := PrimHead(tmp11762)

										tmp11764 := Call(__e, PrimNS2Value(symshen_4chwild), tmp11763)

										tmp11765 := PrimCons(symProcessN, Nil)

										tmp11766 := PrimCons(tmp11764, tmp11765)

										tmp11767 := PrimCons(tmp11759, tmp11766)

										tmp11768 := PrimCons(symshen_4bindv, tmp11767)

										tmp11769 := PrimTail(V1118)

										tmp11770 := PrimTail(tmp11769)

										tmp11771 := PrimTail(tmp11770)

										tmp11772 := PrimTail(tmp11771)

										tmp11773 := PrimTail(tmp11772)

										tmp11774 := PrimHead(tmp11773)

										tmp11775 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), tmp11774)

										tmp11776 := PrimTail(V1118)

										tmp11777 := PrimHead(tmp11776)

										tmp11778 := PrimCons(symProcessN, Nil)

										tmp11779 := PrimCons(tmp11777, tmp11778)

										tmp11780 := PrimCons(symshen_4unbindv, tmp11779)

										tmp11781 := PrimCons(symResult, Nil)

										tmp11782 := PrimCons(tmp11780, tmp11781)

										tmp11783 := PrimCons(symdo, tmp11782)

										tmp11784 := PrimCons(tmp11783, Nil)

										tmp11785 := PrimCons(tmp11775, tmp11784)

										tmp11786 := PrimCons(symResult, tmp11785)

										tmp11787 := PrimCons(symlet, tmp11786)

										tmp11788 := PrimCons(tmp11787, Nil)

										tmp11789 := PrimCons(tmp11768, tmp11788)

										__e.Return(PrimCons(symdo, tmp11789))
										return

									} else {
										tmp11757 := PrimIsPair(V1118)

										var ifres11709 Obj

										if True == tmp11757 {
											tmp11755 := PrimTail(V1118)

											tmp11756 := PrimIsPair(tmp11755)

											var ifres11711 Obj

											if True == tmp11756 {
												tmp11752 := PrimTail(V1118)

												tmp11753 := PrimHead(tmp11752)

												tmp11754 := PrimEqual(symis, tmp11753)

												var ifres11713 Obj

												if True == tmp11754 {
													tmp11749 := PrimTail(V1118)

													tmp11750 := PrimTail(tmp11749)

													tmp11751 := PrimIsPair(tmp11750)

													var ifres11715 Obj

													if True == tmp11751 {
														tmp11745 := PrimTail(V1118)

														tmp11746 := PrimTail(tmp11745)

														tmp11747 := PrimHead(tmp11746)

														tmp11748 := PrimEqual(symidentical, tmp11747)

														var ifres11717 Obj

														if True == tmp11748 {
															tmp11741 := PrimTail(V1118)

															tmp11742 := PrimTail(tmp11741)

															tmp11743 := PrimTail(tmp11742)

															tmp11744 := PrimIsPair(tmp11743)

															var ifres11719 Obj

															if True == tmp11744 {
																tmp11736 := PrimTail(V1118)

																tmp11737 := PrimTail(tmp11736)

																tmp11738 := PrimTail(tmp11737)

																tmp11739 := PrimHead(tmp11738)

																tmp11740 := PrimEqual(symshen_4to, tmp11739)

																var ifres11721 Obj

																if True == tmp11740 {
																	tmp11731 := PrimTail(V1118)

																	tmp11732 := PrimTail(tmp11731)

																	tmp11733 := PrimTail(tmp11732)

																	tmp11734 := PrimTail(tmp11733)

																	tmp11735 := PrimIsPair(tmp11734)

																	var ifres11723 Obj

																	if True == tmp11735 {
																		tmp11725 := PrimTail(V1118)

																		tmp11726 := PrimTail(tmp11725)

																		tmp11727 := PrimTail(tmp11726)

																		tmp11728 := PrimTail(tmp11727)

																		tmp11729 := PrimTail(tmp11728)

																		tmp11730 := PrimEqual(Nil, tmp11729)

																		var ifres11724 Obj

																		if True == tmp11730 {
																			ifres11724 = True

																		} else {
																			ifres11724 = False

																		}

																		ifres11723 = ifres11724

																	} else {
																		ifres11723 = False

																	}

																	var ifres11722 Obj

																	if True == ifres11723 {
																		ifres11722 = True

																	} else {
																		ifres11722 = False

																	}

																	ifres11721 = ifres11722

																} else {
																	ifres11721 = False

																}

																var ifres11720 Obj

																if True == ifres11721 {
																	ifres11720 = True

																} else {
																	ifres11720 = False

																}

																ifres11719 = ifres11720

															} else {
																ifres11719 = False

															}

															var ifres11718 Obj

															if True == ifres11719 {
																ifres11718 = True

															} else {
																ifres11718 = False

															}

															ifres11717 = ifres11718

														} else {
															ifres11717 = False

														}

														var ifres11716 Obj

														if True == ifres11717 {
															ifres11716 = True

														} else {
															ifres11716 = False

														}

														ifres11715 = ifres11716

													} else {
														ifres11715 = False

													}

													var ifres11714 Obj

													if True == ifres11715 {
														ifres11714 = True

													} else {
														ifres11714 = False

													}

													ifres11713 = ifres11714

												} else {
													ifres11713 = False

												}

												var ifres11712 Obj

												if True == ifres11713 {
													ifres11712 = True

												} else {
													ifres11712 = False

												}

												ifres11711 = ifres11712

											} else {
												ifres11711 = False

											}

											var ifres11710 Obj

											if True == ifres11711 {
												ifres11710 = True

											} else {
												ifres11710 = False

											}

											ifres11709 = ifres11710

										} else {
											ifres11709 = False

										}

										if True == ifres11709 {
											tmp11701 := PrimTail(V1118)

											tmp11702 := PrimTail(tmp11701)

											tmp11703 := PrimTail(tmp11702)

											tmp11704 := PrimTail(tmp11703)

											tmp11705 := PrimHead(tmp11704)

											tmp11706 := PrimHead(V1118)

											tmp11707 := PrimCons(tmp11706, Nil)

											tmp11708 := PrimCons(tmp11705, tmp11707)

											__e.Return(PrimCons(sym_a, tmp11708))
											return

										} else {
											tmp11700 := PrimEqual(symshen_4failed_b, V1118)

											if True == tmp11700 {
												__e.Return(False)
												return
											} else {
												tmp11699 := PrimIsPair(V1118)

												var ifres11662 Obj

												if True == tmp11699 {
													tmp11697 := PrimHead(V1118)

													tmp11698 := PrimEqual(symshen_4the, tmp11697)

													var ifres11664 Obj

													if True == tmp11698 {
														tmp11695 := PrimTail(V1118)

														tmp11696 := PrimIsPair(tmp11695)

														var ifres11666 Obj

														if True == tmp11696 {
															tmp11692 := PrimTail(V1118)

															tmp11693 := PrimHead(tmp11692)

															tmp11694 := PrimEqual(symhead, tmp11693)

															var ifres11668 Obj

															if True == tmp11694 {
																tmp11689 := PrimTail(V1118)

																tmp11690 := PrimTail(tmp11689)

																tmp11691 := PrimIsPair(tmp11690)

																var ifres11670 Obj

																if True == tmp11691 {
																	tmp11685 := PrimTail(V1118)

																	tmp11686 := PrimTail(tmp11685)

																	tmp11687 := PrimHead(tmp11686)

																	tmp11688 := PrimEqual(symshen_4of, tmp11687)

																	var ifres11672 Obj

																	if True == tmp11688 {
																		tmp11681 := PrimTail(V1118)

																		tmp11682 := PrimTail(tmp11681)

																		tmp11683 := PrimTail(tmp11682)

																		tmp11684 := PrimIsPair(tmp11683)

																		var ifres11674 Obj

																		if True == tmp11684 {
																			tmp11676 := PrimTail(V1118)

																			tmp11677 := PrimTail(tmp11676)

																			tmp11678 := PrimTail(tmp11677)

																			tmp11679 := PrimTail(tmp11678)

																			tmp11680 := PrimEqual(Nil, tmp11679)

																			var ifres11675 Obj

																			if True == tmp11680 {
																				ifres11675 = True

																			} else {
																				ifres11675 = False

																			}

																			ifres11674 = ifres11675

																		} else {
																			ifres11674 = False

																		}

																		var ifres11673 Obj

																		if True == ifres11674 {
																			ifres11673 = True

																		} else {
																			ifres11673 = False

																		}

																		ifres11672 = ifres11673

																	} else {
																		ifres11672 = False

																	}

																	var ifres11671 Obj

																	if True == ifres11672 {
																		ifres11671 = True

																	} else {
																		ifres11671 = False

																	}

																	ifres11670 = ifres11671

																} else {
																	ifres11670 = False

																}

																var ifres11669 Obj

																if True == ifres11670 {
																	ifres11669 = True

																} else {
																	ifres11669 = False

																}

																ifres11668 = ifres11669

															} else {
																ifres11668 = False

															}

															var ifres11667 Obj

															if True == ifres11668 {
																ifres11667 = True

															} else {
																ifres11667 = False

															}

															ifres11666 = ifres11667

														} else {
															ifres11666 = False

														}

														var ifres11665 Obj

														if True == ifres11666 {
															ifres11665 = True

														} else {
															ifres11665 = False

														}

														ifres11664 = ifres11665

													} else {
														ifres11664 = False

													}

													var ifres11663 Obj

													if True == ifres11664 {
														ifres11663 = True

													} else {
														ifres11663 = False

													}

													ifres11662 = ifres11663

												} else {
													ifres11662 = False

												}

												if True == ifres11662 {
													tmp11659 := PrimTail(V1118)

													tmp11660 := PrimTail(tmp11659)

													tmp11661 := PrimTail(tmp11660)

													__e.Return(PrimCons(symhd, tmp11661))
													return

												} else {
													tmp11658 := PrimIsPair(V1118)

													var ifres11621 Obj

													if True == tmp11658 {
														tmp11656 := PrimHead(V1118)

														tmp11657 := PrimEqual(symshen_4the, tmp11656)

														var ifres11623 Obj

														if True == tmp11657 {
															tmp11654 := PrimTail(V1118)

															tmp11655 := PrimIsPair(tmp11654)

															var ifres11625 Obj

															if True == tmp11655 {
																tmp11651 := PrimTail(V1118)

																tmp11652 := PrimHead(tmp11651)

																tmp11653 := PrimEqual(symtail, tmp11652)

																var ifres11627 Obj

																if True == tmp11653 {
																	tmp11648 := PrimTail(V1118)

																	tmp11649 := PrimTail(tmp11648)

																	tmp11650 := PrimIsPair(tmp11649)

																	var ifres11629 Obj

																	if True == tmp11650 {
																		tmp11644 := PrimTail(V1118)

																		tmp11645 := PrimTail(tmp11644)

																		tmp11646 := PrimHead(tmp11645)

																		tmp11647 := PrimEqual(symshen_4of, tmp11646)

																		var ifres11631 Obj

																		if True == tmp11647 {
																			tmp11640 := PrimTail(V1118)

																			tmp11641 := PrimTail(tmp11640)

																			tmp11642 := PrimTail(tmp11641)

																			tmp11643 := PrimIsPair(tmp11642)

																			var ifres11633 Obj

																			if True == tmp11643 {
																				tmp11635 := PrimTail(V1118)

																				tmp11636 := PrimTail(tmp11635)

																				tmp11637 := PrimTail(tmp11636)

																				tmp11638 := PrimTail(tmp11637)

																				tmp11639 := PrimEqual(Nil, tmp11638)

																				var ifres11634 Obj

																				if True == tmp11639 {
																					ifres11634 = True

																				} else {
																					ifres11634 = False

																				}

																				ifres11633 = ifres11634

																			} else {
																				ifres11633 = False

																			}

																			var ifres11632 Obj

																			if True == ifres11633 {
																				ifres11632 = True

																			} else {
																				ifres11632 = False

																			}

																			ifres11631 = ifres11632

																		} else {
																			ifres11631 = False

																		}

																		var ifres11630 Obj

																		if True == ifres11631 {
																			ifres11630 = True

																		} else {
																			ifres11630 = False

																		}

																		ifres11629 = ifres11630

																	} else {
																		ifres11629 = False

																	}

																	var ifres11628 Obj

																	if True == ifres11629 {
																		ifres11628 = True

																	} else {
																		ifres11628 = False

																	}

																	ifres11627 = ifres11628

																} else {
																	ifres11627 = False

																}

																var ifres11626 Obj

																if True == ifres11627 {
																	ifres11626 = True

																} else {
																	ifres11626 = False

																}

																ifres11625 = ifres11626

															} else {
																ifres11625 = False

															}

															var ifres11624 Obj

															if True == ifres11625 {
																ifres11624 = True

															} else {
																ifres11624 = False

															}

															ifres11623 = ifres11624

														} else {
															ifres11623 = False

														}

														var ifres11622 Obj

														if True == ifres11623 {
															ifres11622 = True

														} else {
															ifres11622 = False

														}

														ifres11621 = ifres11622

													} else {
														ifres11621 = False

													}

													if True == ifres11621 {
														tmp11618 := PrimTail(V1118)

														tmp11619 := PrimTail(tmp11618)

														tmp11620 := PrimTail(tmp11619)

														__e.Return(PrimCons(symtl, tmp11620))
														return

													} else {
														tmp11617 := PrimIsPair(V1118)

														var ifres11587 Obj

														if True == tmp11617 {
															tmp11615 := PrimHead(V1118)

															tmp11616 := PrimEqual(symshen_4pop, tmp11615)

															var ifres11589 Obj

															if True == tmp11616 {
																tmp11613 := PrimTail(V1118)

																tmp11614 := PrimIsPair(tmp11613)

																var ifres11591 Obj

																if True == tmp11614 {
																	tmp11610 := PrimTail(V1118)

																	tmp11611 := PrimHead(tmp11610)

																	tmp11612 := PrimEqual(symshen_4the, tmp11611)

																	var ifres11593 Obj

																	if True == tmp11612 {
																		tmp11607 := PrimTail(V1118)

																		tmp11608 := PrimTail(tmp11607)

																		tmp11609 := PrimIsPair(tmp11608)

																		var ifres11595 Obj

																		if True == tmp11609 {
																			tmp11603 := PrimTail(V1118)

																			tmp11604 := PrimTail(tmp11603)

																			tmp11605 := PrimHead(tmp11604)

																			tmp11606 := PrimEqual(symshen_4stack, tmp11605)

																			var ifres11597 Obj

																			if True == tmp11606 {
																				tmp11599 := PrimTail(V1118)

																				tmp11600 := PrimTail(tmp11599)

																				tmp11601 := PrimTail(tmp11600)

																				tmp11602 := PrimEqual(Nil, tmp11601)

																				var ifres11598 Obj

																				if True == tmp11602 {
																					ifres11598 = True

																				} else {
																					ifres11598 = False

																				}

																				ifres11597 = ifres11598

																			} else {
																				ifres11597 = False

																			}

																			var ifres11596 Obj

																			if True == ifres11597 {
																				ifres11596 = True

																			} else {
																				ifres11596 = False

																			}

																			ifres11595 = ifres11596

																		} else {
																			ifres11595 = False

																		}

																		var ifres11594 Obj

																		if True == ifres11595 {
																			ifres11594 = True

																		} else {
																			ifres11594 = False

																		}

																		ifres11593 = ifres11594

																	} else {
																		ifres11593 = False

																	}

																	var ifres11592 Obj

																	if True == ifres11593 {
																		ifres11592 = True

																	} else {
																		ifres11592 = False

																	}

																	ifres11591 = ifres11592

																} else {
																	ifres11591 = False

																}

																var ifres11590 Obj

																if True == ifres11591 {
																	ifres11590 = True

																} else {
																	ifres11590 = False

																}

																ifres11589 = ifres11590

															} else {
																ifres11589 = False

															}

															var ifres11588 Obj

															if True == ifres11589 {
																ifres11588 = True

															} else {
																ifres11588 = False

															}

															ifres11587 = ifres11588

														} else {
															ifres11587 = False

														}

														if True == ifres11587 {
															tmp11582 := PrimCons(symshen_4incinfs, Nil)

															tmp11583 := PrimCons(symContinuation, Nil)

															tmp11584 := PrimCons(symthaw, tmp11583)

															tmp11585 := PrimCons(tmp11584, Nil)

															tmp11586 := PrimCons(tmp11582, tmp11585)

															__e.Return(PrimCons(symdo, tmp11586))
															return

														} else {
															tmp11581 := PrimIsPair(V1118)

															var ifres11544 Obj

															if True == tmp11581 {
																tmp11579 := PrimHead(V1118)

																tmp11580 := PrimEqual(symcall, tmp11579)

																var ifres11546 Obj

																if True == tmp11580 {
																	tmp11577 := PrimTail(V1118)

																	tmp11578 := PrimIsPair(tmp11577)

																	var ifres11548 Obj

																	if True == tmp11578 {
																		tmp11574 := PrimTail(V1118)

																		tmp11575 := PrimHead(tmp11574)

																		tmp11576 := PrimEqual(symshen_4the, tmp11575)

																		var ifres11550 Obj

																		if True == tmp11576 {
																			tmp11571 := PrimTail(V1118)

																			tmp11572 := PrimTail(tmp11571)

																			tmp11573 := PrimIsPair(tmp11572)

																			var ifres11552 Obj

																			if True == tmp11573 {
																				tmp11567 := PrimTail(V1118)

																				tmp11568 := PrimTail(tmp11567)

																				tmp11569 := PrimHead(tmp11568)

																				tmp11570 := PrimEqual(symshen_4continuation, tmp11569)

																				var ifres11554 Obj

																				if True == tmp11570 {
																					tmp11563 := PrimTail(V1118)

																					tmp11564 := PrimTail(tmp11563)

																					tmp11565 := PrimTail(tmp11564)

																					tmp11566 := PrimIsPair(tmp11565)

																					var ifres11556 Obj

																					if True == tmp11566 {
																						tmp11558 := PrimTail(V1118)

																						tmp11559 := PrimTail(tmp11558)

																						tmp11560 := PrimTail(tmp11559)

																						tmp11561 := PrimTail(tmp11560)

																						tmp11562 := PrimEqual(Nil, tmp11561)

																						var ifres11557 Obj

																						if True == tmp11562 {
																							ifres11557 = True

																						} else {
																							ifres11557 = False

																						}

																						ifres11556 = ifres11557

																					} else {
																						ifres11556 = False

																					}

																					var ifres11555 Obj

																					if True == ifres11556 {
																						ifres11555 = True

																					} else {
																						ifres11555 = False

																					}

																					ifres11554 = ifres11555

																				} else {
																					ifres11554 = False

																				}

																				var ifres11553 Obj

																				if True == ifres11554 {
																					ifres11553 = True

																				} else {
																					ifres11553 = False

																				}

																				ifres11552 = ifres11553

																			} else {
																				ifres11552 = False

																			}

																			var ifres11551 Obj

																			if True == ifres11552 {
																				ifres11551 = True

																			} else {
																				ifres11551 = False

																			}

																			ifres11550 = ifres11551

																		} else {
																			ifres11550 = False

																		}

																		var ifres11549 Obj

																		if True == ifres11550 {
																			ifres11549 = True

																		} else {
																			ifres11549 = False

																		}

																		ifres11548 = ifres11549

																	} else {
																		ifres11548 = False

																	}

																	var ifres11547 Obj

																	if True == ifres11548 {
																		ifres11547 = True

																	} else {
																		ifres11547 = False

																	}

																	ifres11546 = ifres11547

																} else {
																	ifres11546 = False

																}

																var ifres11545 Obj

																if True == ifres11546 {
																	ifres11545 = True

																} else {
																	ifres11545 = False

																}

																ifres11544 = ifres11545

															} else {
																ifres11544 = False

															}

															if True == ifres11544 {
																tmp11535 := PrimCons(symshen_4incinfs, Nil)

																tmp11536 := PrimTail(V1118)

																tmp11537 := PrimTail(tmp11536)

																tmp11538 := PrimTail(tmp11537)

																tmp11539 := PrimHead(tmp11538)

																tmp11540 := Call(__e, PrimNS2Value(symshen_4chwild), tmp11539)

																tmp11541 := Call(__e, PrimNS2Value(symshen_4call__the__continuation), tmp11540, symProcessN, symContinuation)

																tmp11542 := PrimCons(tmp11541, Nil)

																tmp11543 := PrimCons(tmp11535, tmp11542)

																__e.Return(PrimCons(symdo, tmp11543))
																return

															} else {
																__e.Return(V1118)
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

						}

					}

				}

			}

		}

	}, 1)

	tmp12418 := Call(__e, PrimNS1Value(symns2_1set), symshen_4aum__to__shen, tmp11520)

	_ = tmp12418

	tmp12419 := MakeNative(func(__e *ControlFlow) {
		V1120 := __e.Get(1)
		_ = V1120
		tmp12425 := PrimEqual(V1120, sym__)

		if True == tmp12425 {
			tmp12424 := PrimCons(symProcessN, Nil)

			__e.Return(PrimCons(symshen_4newpv, tmp12424))
			return

		} else {
			tmp12423 := PrimIsPair(V1120)

			if True == tmp12423 {
				tmp12422 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4chwild), Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp12422, V1120)
				return

			} else {
				__e.Return(V1120)
				return
			}

		}

	}, 1)

	tmp12426 := Call(__e, PrimNS1Value(symns2_1set), symshen_4chwild, tmp12419)

	_ = tmp12426

	tmp12427 := MakeNative(func(__e *ControlFlow) {
		V1122 := __e.Get(1)
		_ = V1122
		tmp12428 := MakeNative(func(__e *ControlFlow) {
			Count_71 := __e.Get(1)
			_ = Count_71
			tmp12429 := MakeNative(func(__e *ControlFlow) {
				IncVar := __e.Get(1)
				_ = IncVar
				tmp12430 := MakeNative(func(__e *ControlFlow) {
					Vector := __e.Get(1)
					_ = Vector
					tmp12431 := MakeNative(func(__e *ControlFlow) {
						ResizeVectorIfNeeded := __e.Get(1)
						_ = ResizeVectorIfNeeded
						__e.TailApply(PrimNS2Value(symshen_4mk_1pvar), Count_71)
						return
					}, 1)

					tmp12434 := Call(__e, PrimNS2Value(symlimit), Vector)

					tmp12435 := PrimEqual(Count_71, tmp12434)

					var ifres12432 Obj

					if True == tmp12435 {
						tmp12433 := Call(__e, PrimNS2Value(symshen_4resizeprocessvector), V1122, Count_71)

						ifres12432 = tmp12433

					} else {
						ifres12432 = symshen_4skip

					}

					__e.TailApply(tmp12431, ifres12432)
					return

				}, 1)

				tmp12436 := PrimNS3Value(symshen_4_dprologvectors_d)

				tmp12437 := PrimVectorGet(tmp12436, V1122)

				__e.TailApply(tmp12430, tmp12437)
				return

			}, 1)

			tmp12438 := PrimNS3Value(symshen_4_dvarcounter_d)

			tmp12439 := PrimVectorSet(tmp12438, V1122, Count_71)

			__e.TailApply(tmp12429, tmp12439)
			return

		}, 1)

		tmp12440 := PrimNS3Value(symshen_4_dvarcounter_d)

		tmp12441 := PrimVectorGet(tmp12440, V1122)

		tmp12442 := PrimNumberAdd(tmp12441, MakeNumber(1))

		__e.TailApply(tmp12428, tmp12442)
		return

	}, 1)

	tmp12443 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newpv, tmp12427)

	_ = tmp12443

	tmp12444 := MakeNative(func(__e *ControlFlow) {
		V1125 := __e.Get(1)
		_ = V1125
		V1126 := __e.Get(2)
		_ = V1126
		tmp12445 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp12446 := MakeNative(func(__e *ControlFlow) {
				BigVector := __e.Get(1)
				_ = BigVector
				tmp12447 := PrimNS3Value(symshen_4_dprologvectors_d)

				__e.Return(PrimVectorSet(tmp12447, V1125, BigVector))
				return

			}, 1)

			tmp12448 := PrimNumberAdd(V1126, V1126)

			tmp12449 := Call(__e, PrimNS2Value(symshen_4resize_1vector), Vector, tmp12448, symshen_4_1null_1)

			__e.TailApply(tmp12446, tmp12449)
			return

		}, 1)

		tmp12450 := PrimNS3Value(symshen_4_dprologvectors_d)

		tmp12451 := PrimVectorGet(tmp12450, V1125)

		__e.TailApply(tmp12445, tmp12451)
		return

	}, 2)

	tmp12452 := Call(__e, PrimNS1Value(symns2_1set), symshen_4resizeprocessvector, tmp12444)

	_ = tmp12452

	tmp12453 := MakeNative(func(__e *ControlFlow) {
		V1130 := __e.Get(1)
		_ = V1130
		V1131 := __e.Get(2)
		_ = V1131
		V1132 := __e.Get(3)
		_ = V1132
		tmp12454 := MakeNative(func(__e *ControlFlow) {
			BigVector := __e.Get(1)
			_ = BigVector
			tmp12455 := Call(__e, PrimNS2Value(symlimit), V1130)

			__e.TailApply(PrimNS2Value(symshen_4copy_1vector), V1130, BigVector, tmp12455, V1131, V1132)
			return

		}, 1)

		tmp12456 := PrimNumberAdd(MakeNumber(1), V1131)

		tmp12457 := PrimAbsvector(tmp12456)

		tmp12458 := PrimVectorSet(tmp12457, MakeNumber(0), V1131)

		__e.TailApply(tmp12454, tmp12458)
		return

	}, 3)

	tmp12459 := Call(__e, PrimNS1Value(symns2_1set), symshen_4resize_1vector, tmp12453)

	_ = tmp12459

	tmp12460 := MakeNative(func(__e *ControlFlow) {
		V1138 := __e.Get(1)
		_ = V1138
		V1139 := __e.Get(2)
		_ = V1139
		V1140 := __e.Get(3)
		_ = V1140
		V1141 := __e.Get(4)
		_ = V1141
		V1142 := __e.Get(5)
		_ = V1142
		tmp12461 := PrimNumberAdd(MakeNumber(1), V1140)

		tmp12462 := PrimNumberAdd(V1141, MakeNumber(1))

		tmp12463 := PrimNumberAdd(MakeNumber(1), V1140)

		tmp12464 := Call(__e, PrimNS2Value(symshen_4copy_1vector_1stage_11), MakeNumber(1), V1138, V1139, tmp12463)

		__e.TailApply(PrimNS2Value(symshen_4copy_1vector_1stage_12), tmp12461, tmp12462, V1142, tmp12464)
		return

	}, 5)

	tmp12465 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector, tmp12460)

	_ = tmp12465

	tmp12466 := MakeNative(func(__e *ControlFlow) {
		V1150 := __e.Get(1)
		_ = V1150
		V1151 := __e.Get(2)
		_ = V1151
		V1152 := __e.Get(3)
		_ = V1152
		V1153 := __e.Get(4)
		_ = V1153
		tmp12471 := PrimEqual(V1153, V1150)

		if True == tmp12471 {
			__e.Return(V1152)
			return
		} else {
			tmp12468 := PrimNumberAdd(MakeNumber(1), V1150)

			tmp12469 := PrimVectorGet(V1151, V1150)

			tmp12470 := PrimVectorSet(V1152, V1150, tmp12469)

			__e.TailApply(PrimNS2Value(symshen_4copy_1vector_1stage_11), tmp12468, V1151, tmp12470, V1153)
			return

		}

	}, 4)

	tmp12472 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector_1stage_11, tmp12466)

	_ = tmp12472

	tmp12473 := MakeNative(func(__e *ControlFlow) {
		V1161 := __e.Get(1)
		_ = V1161
		V1162 := __e.Get(2)
		_ = V1162
		V1163 := __e.Get(3)
		_ = V1163
		V1164 := __e.Get(4)
		_ = V1164
		tmp12477 := PrimEqual(V1162, V1161)

		if True == tmp12477 {
			__e.Return(V1164)
			return
		} else {
			tmp12475 := PrimNumberAdd(V1161, MakeNumber(1))

			tmp12476 := PrimVectorSet(V1164, V1161, V1163)

			__e.TailApply(PrimNS2Value(symshen_4copy_1vector_1stage_12), tmp12475, V1162, V1163, tmp12476)
			return

		}

	}, 4)

	tmp12478 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copy_1vector_1stage_12, tmp12473)

	_ = tmp12478

	tmp12479 := MakeNative(func(__e *ControlFlow) {
		V1166 := __e.Get(1)
		_ = V1166
		tmp12480 := PrimAbsvector(MakeNumber(2))

		tmp12481 := PrimVectorSet(tmp12480, MakeNumber(0), symshen_4pvar)

		__e.Return(PrimVectorSet(tmp12481, MakeNumber(1), V1166))
		return

	}, 1)

	tmp12482 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mk_1pvar, tmp12479)

	_ = tmp12482

	tmp12483 := MakeNative(func(__e *ControlFlow) {
		V1168 := __e.Get(1)
		_ = V1168
		tmp12490 := PrimIsVector(V1168)

		if True == tmp12490 {
			tmp12486 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V1168, MakeNumber(0)))
				return
			}, 0)

			tmp12487 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1pvar)
				return
			}, 1)

			tmp12488 := Call(__e, PrimNS1Value(symtry_1catch), tmp12486, tmp12487)

			tmp12489 := PrimEqual(tmp12488, symshen_4pvar)

			if True == tmp12489 {
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

	tmp12491 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pvar_2, tmp12483)

	_ = tmp12491

	tmp12492 := MakeNative(func(__e *ControlFlow) {
		V1172 := __e.Get(1)
		_ = V1172
		V1173 := __e.Get(2)
		_ = V1173
		V1174 := __e.Get(3)
		_ = V1174
		tmp12493 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp12494 := PrimVectorGet(V1172, MakeNumber(1))

			__e.Return(PrimVectorSet(Vector, tmp12494, V1173))
			return

		}, 1)

		tmp12495 := PrimNS3Value(symshen_4_dprologvectors_d)

		tmp12496 := PrimVectorGet(tmp12495, V1174)

		__e.TailApply(tmp12493, tmp12496)
		return

	}, 3)

	tmp12497 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bindv, tmp12492)

	_ = tmp12497

	tmp12498 := MakeNative(func(__e *ControlFlow) {
		V1177 := __e.Get(1)
		_ = V1177
		V1178 := __e.Get(2)
		_ = V1178
		tmp12499 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp12500 := PrimVectorGet(V1177, MakeNumber(1))

			__e.Return(PrimVectorSet(Vector, tmp12500, symshen_4_1null_1))
			return

		}, 1)

		tmp12501 := PrimNS3Value(symshen_4_dprologvectors_d)

		tmp12502 := PrimVectorGet(tmp12501, V1178)

		__e.TailApply(tmp12499, tmp12502)
		return

	}, 2)

	tmp12503 := Call(__e, PrimNS1Value(symns2_1set), symshen_4unbindv, tmp12498)

	_ = tmp12503

	tmp12504 := MakeNative(func(__e *ControlFlow) {
		tmp12505 := PrimNS3Value(symshen_4_dinfs_d)

		tmp12506 := PrimNumberAdd(MakeNumber(1), tmp12505)

		__e.Return(PrimNS3Set(symshen_4_dinfs_d, tmp12506))
		return

	}, 0)

	tmp12507 := Call(__e, PrimNS1Value(symns2_1set), symshen_4incinfs, tmp12504)

	_ = tmp12507

	tmp12508 := MakeNative(func(__e *ControlFlow) {
		V1182 := __e.Get(1)
		_ = V1182
		V1183 := __e.Get(2)
		_ = V1183
		V1184 := __e.Get(3)
		_ = V1184
		tmp12541 := PrimIsPair(V1182)

		var ifres12533 Obj

		if True == tmp12541 {
			tmp12539 := PrimHead(V1182)

			tmp12540 := PrimIsPair(tmp12539)

			var ifres12535 Obj

			if True == tmp12540 {
				tmp12537 := PrimTail(V1182)

				tmp12538 := PrimEqual(Nil, tmp12537)

				var ifres12536 Obj

				if True == tmp12538 {
					ifres12536 = True

				} else {
					ifres12536 = False

				}

				ifres12535 = ifres12536

			} else {
				ifres12535 = False

			}

			var ifres12534 Obj

			if True == ifres12535 {
				ifres12534 = True

			} else {
				ifres12534 = False

			}

			ifres12533 = ifres12534

		} else {
			ifres12533 = False

		}

		if True == ifres12533 {
			tmp12526 := PrimHead(V1182)

			tmp12527 := PrimHead(tmp12526)

			tmp12528 := PrimHead(V1182)

			tmp12529 := PrimTail(tmp12528)

			tmp12530 := PrimCons(V1184, Nil)

			tmp12531 := PrimCons(V1183, tmp12530)

			tmp12532 := Call(__e, PrimNS2Value(symappend), tmp12529, tmp12531)

			__e.Return(PrimCons(tmp12527, tmp12532))
			return

		} else {
			tmp12525 := PrimIsPair(V1182)

			var ifres12521 Obj

			if True == tmp12525 {
				tmp12523 := PrimHead(V1182)

				tmp12524 := PrimIsPair(tmp12523)

				var ifres12522 Obj

				if True == tmp12524 {
					ifres12522 = True

				} else {
					ifres12522 = False

				}

				ifres12521 = ifres12522

			} else {
				ifres12521 = False

			}

			if True == ifres12521 {
				tmp12511 := MakeNative(func(__e *ControlFlow) {
					NewContinuation := __e.Get(1)
					_ = NewContinuation
					tmp12512 := PrimHead(V1182)

					tmp12513 := PrimHead(tmp12512)

					tmp12514 := PrimHead(V1182)

					tmp12515 := PrimTail(tmp12514)

					tmp12516 := PrimCons(NewContinuation, Nil)

					tmp12517 := PrimCons(V1183, tmp12516)

					tmp12518 := Call(__e, PrimNS2Value(symappend), tmp12515, tmp12517)

					__e.Return(PrimCons(tmp12513, tmp12518))
					return

				}, 1)

				tmp12519 := PrimTail(V1182)

				tmp12520 := Call(__e, PrimNS2Value(symshen_4newcontinuation), tmp12519, V1183, V1184)

				__e.TailApply(tmp12511, tmp12520)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4call__the__continuation)
				return
			}

		}

	}, 3)

	tmp12542 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call__the__continuation, tmp12508)

	_ = tmp12542

	tmp12543 := MakeNative(func(__e *ControlFlow) {
		V1188 := __e.Get(1)
		_ = V1188
		V1189 := __e.Get(2)
		_ = V1189
		V1190 := __e.Get(3)
		_ = V1190
		tmp12562 := PrimEqual(Nil, V1188)

		if True == tmp12562 {
			__e.Return(V1190)
			return
		} else {
			tmp12561 := PrimIsPair(V1188)

			var ifres12557 Obj

			if True == tmp12561 {
				tmp12559 := PrimHead(V1188)

				tmp12560 := PrimIsPair(tmp12559)

				var ifres12558 Obj

				if True == tmp12560 {
					ifres12558 = True

				} else {
					ifres12558 = False

				}

				ifres12557 = ifres12558

			} else {
				ifres12557 = False

			}

			if True == ifres12557 {
				tmp12546 := PrimHead(V1188)

				tmp12547 := PrimHead(tmp12546)

				tmp12548 := PrimHead(V1188)

				tmp12549 := PrimTail(tmp12548)

				tmp12550 := PrimTail(V1188)

				tmp12551 := Call(__e, PrimNS2Value(symshen_4newcontinuation), tmp12550, V1189, V1190)

				tmp12552 := PrimCons(tmp12551, Nil)

				tmp12553 := PrimCons(V1189, tmp12552)

				tmp12554 := Call(__e, PrimNS2Value(symappend), tmp12549, tmp12553)

				tmp12555 := PrimCons(tmp12547, tmp12554)

				tmp12556 := PrimCons(tmp12555, Nil)

				__e.Return(PrimCons(symfreeze, tmp12556))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4newcontinuation)
				return
			}

		}

	}, 3)

	tmp12563 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newcontinuation, tmp12543)

	_ = tmp12563

	tmp12564 := MakeNative(func(__e *ControlFlow) {
		V1198 := __e.Get(1)
		_ = V1198
		V1199 := __e.Get(2)
		_ = V1199
		V1200 := __e.Get(3)
		_ = V1200
		__e.TailApply(PrimNS2Value(symshen_4deref), V1198, V1199)
		return
	}, 3)

	tmp12565 := Call(__e, PrimNS1Value(symns2_1set), symreturn, tmp12564)

	_ = tmp12565

	tmp12566 := MakeNative(func(__e *ControlFlow) {
		V1208 := __e.Get(1)
		_ = V1208
		V1209 := __e.Get(2)
		_ = V1209
		V1210 := __e.Get(3)
		_ = V1210
		tmp12567 := PrimNS3Value(symshen_4_dinfs_d)

		tmp12568 := Call(__e, PrimNS2Value(symshen_4app), tmp12567, MakeString(" inferences\n"), symshen_4a)

		tmp12569 := Call(__e, PrimNS2Value(symstoutput))

		tmp12570 := Call(__e, PrimNS2Value(symshen_4prhush), tmp12568, tmp12569)

		_ = tmp12570

		__e.TailApply(PrimNS2Value(symshen_4deref), V1208, V1209)
		return

	}, 3)

	tmp12571 := Call(__e, PrimNS1Value(symns2_1set), symshen_4measure_ereturn, tmp12566)

	_ = tmp12571

	tmp12572 := MakeNative(func(__e *ControlFlow) {
		V1215 := __e.Get(1)
		_ = V1215
		V1216 := __e.Get(2)
		_ = V1216
		V1217 := __e.Get(3)
		_ = V1217
		V1218 := __e.Get(4)
		_ = V1218
		tmp12573 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1215, V1217)

		tmp12574 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1216, V1217)

		__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp12573, tmp12574, V1217, V1218)
		return

	}, 4)

	tmp12575 := Call(__e, PrimNS1Value(symns2_1set), symunify, tmp12572)

	_ = tmp12575

	tmp12576 := MakeNative(func(__e *ControlFlow) {
		V1240 := __e.Get(1)
		_ = V1240
		V1241 := __e.Get(2)
		_ = V1241
		V1242 := __e.Get(3)
		_ = V1242
		V1243 := __e.Get(4)
		_ = V1243
		tmp12596 := PrimEqual(V1241, V1240)

		if True == tmp12596 {
			__e.TailApply(PrimNS2Value(symthaw), V1243)
			return
		} else {
			tmp12595 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1240)

			if True == tmp12595 {
				__e.TailApply(PrimNS2Value(symbind), V1240, V1241, V1242, V1243)
				return
			} else {
				tmp12594 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1241)

				if True == tmp12594 {
					__e.TailApply(PrimNS2Value(symbind), V1241, V1240, V1242, V1243)
					return
				} else {
					tmp12593 := PrimIsPair(V1240)

					var ifres12590 Obj

					if True == tmp12593 {
						tmp12592 := PrimIsPair(V1241)

						var ifres12591 Obj

						if True == tmp12592 {
							ifres12591 = True

						} else {
							ifres12591 = False

						}

						ifres12590 = ifres12591

					} else {
						ifres12590 = False

					}

					if True == ifres12590 {
						tmp12581 := PrimHead(V1240)

						tmp12582 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12581, V1242)

						tmp12583 := PrimHead(V1241)

						tmp12584 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12583, V1242)

						tmp12585 := MakeNative(func(__e *ControlFlow) {
							tmp12586 := PrimTail(V1240)

							tmp12587 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12586, V1242)

							tmp12588 := PrimTail(V1241)

							tmp12589 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12588, V1242)

							__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp12587, tmp12589, V1242, V1243)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symshen_4lzy_a), tmp12582, tmp12584, V1242, tmp12585)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp12597 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a, tmp12576)

	_ = tmp12597

	tmp12598 := MakeNative(func(__e *ControlFlow) {
		V1246 := __e.Get(1)
		_ = V1246
		V1247 := __e.Get(2)
		_ = V1247
		tmp12610 := PrimIsPair(V1246)

		if True == tmp12610 {
			tmp12606 := PrimHead(V1246)

			tmp12607 := Call(__e, PrimNS2Value(symshen_4deref), tmp12606, V1247)

			tmp12608 := PrimTail(V1246)

			tmp12609 := Call(__e, PrimNS2Value(symshen_4deref), tmp12608, V1247)

			__e.Return(PrimCons(tmp12607, tmp12609))
			return

		} else {
			tmp12605 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1246)

			if True == tmp12605 {
				tmp12601 := MakeNative(func(__e *ControlFlow) {
					Value := __e.Get(1)
					_ = Value
					tmp12603 := PrimEqual(Value, symshen_4_1null_1)

					if True == tmp12603 {
						__e.Return(V1246)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4deref), Value, V1247)
						return
					}

				}, 1)

				tmp12604 := Call(__e, PrimNS2Value(symshen_4valvector), V1246, V1247)

				__e.TailApply(tmp12601, tmp12604)
				return

			} else {
				__e.Return(V1246)
				return
			}

		}

	}, 2)

	tmp12611 := Call(__e, PrimNS1Value(symns2_1set), symshen_4deref, tmp12598)

	_ = tmp12611

	tmp12612 := MakeNative(func(__e *ControlFlow) {
		V1250 := __e.Get(1)
		_ = V1250
		V1251 := __e.Get(2)
		_ = V1251
		tmp12618 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1250)

		if True == tmp12618 {
			tmp12614 := MakeNative(func(__e *ControlFlow) {
				Value := __e.Get(1)
				_ = Value
				tmp12616 := PrimEqual(Value, symshen_4_1null_1)

				if True == tmp12616 {
					__e.Return(V1250)
					return
				} else {
					__e.TailApply(PrimNS2Value(symshen_4lazyderef), Value, V1251)
					return
				}

			}, 1)

			tmp12617 := Call(__e, PrimNS2Value(symshen_4valvector), V1250, V1251)

			__e.TailApply(tmp12614, tmp12617)
			return

		} else {
			__e.Return(V1250)
			return
		}

	}, 2)

	tmp12619 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lazyderef, tmp12612)

	_ = tmp12619

	tmp12620 := MakeNative(func(__e *ControlFlow) {
		V1254 := __e.Get(1)
		_ = V1254
		V1255 := __e.Get(2)
		_ = V1255
		tmp12621 := PrimNS3Value(symshen_4_dprologvectors_d)

		tmp12622 := PrimVectorGet(tmp12621, V1255)

		tmp12623 := PrimVectorGet(V1254, MakeNumber(1))

		__e.Return(PrimVectorGet(tmp12622, tmp12623))
		return

	}, 2)

	tmp12624 := Call(__e, PrimNS1Value(symns2_1set), symshen_4valvector, tmp12620)

	_ = tmp12624

	tmp12625 := MakeNative(func(__e *ControlFlow) {
		V1260 := __e.Get(1)
		_ = V1260
		V1261 := __e.Get(2)
		_ = V1261
		V1262 := __e.Get(3)
		_ = V1262
		V1263 := __e.Get(4)
		_ = V1263
		tmp12626 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1260, V1262)

		tmp12627 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1261, V1262)

		__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp12626, tmp12627, V1262, V1263)
		return

	}, 4)

	tmp12628 := Call(__e, PrimNS1Value(symns2_1set), symunify_b, tmp12625)

	_ = tmp12628

	tmp12629 := MakeNative(func(__e *ControlFlow) {
		V1285 := __e.Get(1)
		_ = V1285
		V1286 := __e.Get(2)
		_ = V1286
		V1287 := __e.Get(3)
		_ = V1287
		V1288 := __e.Get(4)
		_ = V1288
		tmp12659 := PrimEqual(V1286, V1285)

		if True == tmp12659 {
			__e.TailApply(PrimNS2Value(symthaw), V1288)
			return
		} else {
			tmp12658 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1285)

			var ifres12653 Obj

			if True == tmp12658 {
				tmp12655 := Call(__e, PrimNS2Value(symshen_4deref), V1286, V1287)

				tmp12656 := Call(__e, PrimNS2Value(symshen_4occurs_2), V1285, tmp12655)

				tmp12657 := PrimNot(tmp12656)

				var ifres12654 Obj

				if True == tmp12657 {
					ifres12654 = True

				} else {
					ifres12654 = False

				}

				ifres12653 = ifres12654

			} else {
				ifres12653 = False

			}

			if True == ifres12653 {
				__e.TailApply(PrimNS2Value(symbind), V1285, V1286, V1287, V1288)
				return
			} else {
				tmp12652 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1286)

				var ifres12647 Obj

				if True == tmp12652 {
					tmp12649 := Call(__e, PrimNS2Value(symshen_4deref), V1285, V1287)

					tmp12650 := Call(__e, PrimNS2Value(symshen_4occurs_2), V1286, tmp12649)

					tmp12651 := PrimNot(tmp12650)

					var ifres12648 Obj

					if True == tmp12651 {
						ifres12648 = True

					} else {
						ifres12648 = False

					}

					ifres12647 = ifres12648

				} else {
					ifres12647 = False

				}

				if True == ifres12647 {
					__e.TailApply(PrimNS2Value(symbind), V1286, V1285, V1287, V1288)
					return
				} else {
					tmp12646 := PrimIsPair(V1285)

					var ifres12643 Obj

					if True == tmp12646 {
						tmp12645 := PrimIsPair(V1286)

						var ifres12644 Obj

						if True == tmp12645 {
							ifres12644 = True

						} else {
							ifres12644 = False

						}

						ifres12643 = ifres12644

					} else {
						ifres12643 = False

					}

					if True == ifres12643 {
						tmp12634 := PrimHead(V1285)

						tmp12635 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12634, V1287)

						tmp12636 := PrimHead(V1286)

						tmp12637 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12636, V1287)

						tmp12638 := MakeNative(func(__e *ControlFlow) {
							tmp12639 := PrimTail(V1285)

							tmp12640 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12639, V1287)

							tmp12641 := PrimTail(V1286)

							tmp12642 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12641, V1287)

							__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp12640, tmp12642, V1287, V1288)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symshen_4lzy_a_b), tmp12635, tmp12637, V1287, tmp12638)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 4)

	tmp12660 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a_b, tmp12629)

	_ = tmp12660

	tmp12661 := MakeNative(func(__e *ControlFlow) {
		V1300 := __e.Get(1)
		_ = V1300
		V1301 := __e.Get(2)
		_ = V1301
		tmp12671 := PrimEqual(V1301, V1300)

		if True == tmp12671 {
			__e.Return(True)
			return
		} else {
			tmp12670 := PrimIsPair(V1301)

			if True == tmp12670 {
				tmp12668 := PrimHead(V1301)

				tmp12669 := Call(__e, PrimNS2Value(symshen_4occurs_2), V1300, tmp12668)

				if True == tmp12669 {
					__e.Return(True)
					return
				} else {
					tmp12666 := PrimTail(V1301)

					tmp12667 := Call(__e, PrimNS2Value(symshen_4occurs_2), V1300, tmp12666)

					if True == tmp12667 {
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

	tmp12672 := Call(__e, PrimNS1Value(symns2_1set), symshen_4occurs_2, tmp12661)

	_ = tmp12672

	tmp12673 := MakeNative(func(__e *ControlFlow) {
		V1306 := __e.Get(1)
		_ = V1306
		V1307 := __e.Get(2)
		_ = V1307
		V1308 := __e.Get(3)
		_ = V1308
		V1309 := __e.Get(4)
		_ = V1309
		tmp12674 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1306, V1308)

		tmp12675 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1307, V1308)

		__e.TailApply(PrimNS2Value(symshen_4lzy_a_a), tmp12674, tmp12675, V1308, V1309)
		return

	}, 4)

	tmp12676 := Call(__e, PrimNS1Value(symns2_1set), symidentical, tmp12673)

	_ = tmp12676

	tmp12677 := MakeNative(func(__e *ControlFlow) {
		V1331 := __e.Get(1)
		_ = V1331
		V1332 := __e.Get(2)
		_ = V1332
		V1333 := __e.Get(3)
		_ = V1333
		V1334 := __e.Get(4)
		_ = V1334
		tmp12691 := PrimEqual(V1332, V1331)

		if True == tmp12691 {
			__e.TailApply(PrimNS2Value(symthaw), V1334)
			return
		} else {
			tmp12690 := PrimIsPair(V1331)

			var ifres12687 Obj

			if True == tmp12690 {
				tmp12689 := PrimIsPair(V1332)

				var ifres12688 Obj

				if True == tmp12689 {
					ifres12688 = True

				} else {
					ifres12688 = False

				}

				ifres12687 = ifres12688

			} else {
				ifres12687 = False

			}

			if True == ifres12687 {
				tmp12680 := PrimHead(V1331)

				tmp12681 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12680, V1333)

				tmp12682 := PrimHead(V1332)

				tmp12683 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12682, V1333)

				tmp12684 := MakeNative(func(__e *ControlFlow) {
					tmp12685 := PrimTail(V1331)

					tmp12686 := PrimTail(V1332)

					__e.TailApply(PrimNS2Value(symshen_4lzy_a_a), tmp12685, tmp12686, V1333, V1334)
					return

				}, 0)

				__e.TailApply(PrimNS2Value(symshen_4lzy_a_a), tmp12681, tmp12683, V1333, tmp12684)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 4)

	tmp12692 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lzy_a_a, tmp12677)

	_ = tmp12692

	tmp12693 := MakeNative(func(__e *ControlFlow) {
		V1336 := __e.Get(1)
		_ = V1336
		tmp12694 := PrimVectorGet(V1336, MakeNumber(1))

		tmp12695 := Call(__e, PrimNS2Value(symshen_4app), tmp12694, MakeString(""), symshen_4a)

		__e.Return(PrimStringConcat(MakeString("Var"), tmp12695))
		return

	}, 1)

	tmp12696 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pvar, tmp12693)

	_ = tmp12696

	tmp12697 := MakeNative(func(__e *ControlFlow) {
		V1341 := __e.Get(1)
		_ = V1341
		V1342 := __e.Get(2)
		_ = V1342
		V1343 := __e.Get(3)
		_ = V1343
		V1344 := __e.Get(4)
		_ = V1344
		tmp12698 := Call(__e, PrimNS2Value(symshen_4bindv), V1341, V1342, V1343)

		_ = tmp12698

		tmp12699 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp12700 := Call(__e, PrimNS2Value(symshen_4unbindv), V1341, V1343)

			_ = tmp12700

			__e.Return(Result)
			return

		}, 1)

		tmp12701 := Call(__e, PrimNS2Value(symthaw), V1344)

		__e.TailApply(tmp12699, tmp12701)
		return

	}, 4)

	tmp12702 := Call(__e, PrimNS1Value(symns2_1set), symbind, tmp12697)

	_ = tmp12702

	tmp12703 := MakeNative(func(__e *ControlFlow) {
		V1362 := __e.Get(1)
		_ = V1362
		V1363 := __e.Get(2)
		_ = V1363
		V1364 := __e.Get(3)
		_ = V1364
		tmp12709 := PrimEqual(True, V1362)

		if True == tmp12709 {
			__e.TailApply(PrimNS2Value(symthaw), V1364)
			return
		} else {
			tmp12708 := PrimEqual(False, V1362)

			if True == tmp12708 {
				__e.Return(False)
				return
			} else {
				tmp12706 := Call(__e, PrimNS2Value(symshen_4app), V1362, MakeString("%"), symshen_4s)

				tmp12707 := PrimStringConcat(MakeString("fwhen expects a boolean: not "), tmp12706)

				__e.Return(PrimSimpleError(tmp12707))
				return

			}

		}

	}, 3)

	tmp12710 := Call(__e, PrimNS1Value(symns2_1set), symfwhen, tmp12703)

	_ = tmp12710

	tmp12711 := MakeNative(func(__e *ControlFlow) {
		V1380 := __e.Get(1)
		_ = V1380
		V1381 := __e.Get(2)
		_ = V1381
		V1382 := __e.Get(3)
		_ = V1382
		tmp12720 := PrimIsPair(V1380)

		if True == tmp12720 {
			tmp12716 := PrimHead(V1380)

			tmp12717 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp12716, V1381)

			tmp12718 := Call(__e, PrimNS2Value(symfunction), tmp12717)

			tmp12719 := PrimTail(V1380)

			__e.TailApply(PrimNS2Value(symshen_4call_1help), tmp12718, tmp12719, V1381, V1382)
			return

		} else {
			tmp12715 := Call(__e, PrimNS2Value(symshen_4pvar_2), V1380)

			if True == tmp12715 {
				tmp12714 := Call(__e, PrimNS2Value(symshen_4lazyderef), V1380, V1381)

				__e.TailApply(PrimNS2Value(symcall), tmp12714, V1381, V1382)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 3)

	tmp12721 := Call(__e, PrimNS1Value(symns2_1set), symcall, tmp12711)

	_ = tmp12721

	tmp12722 := MakeNative(func(__e *ControlFlow) {
		V1387 := __e.Get(1)
		_ = V1387
		V1388 := __e.Get(2)
		_ = V1388
		V1389 := __e.Get(3)
		_ = V1389
		V1390 := __e.Get(4)
		_ = V1390
		tmp12729 := PrimEqual(Nil, V1388)

		if True == tmp12729 {
			__e.TailApply(V1387, V1389, V1390)
			return
		} else {
			tmp12728 := PrimIsPair(V1388)

			if True == tmp12728 {
				tmp12725 := PrimHead(V1388)

				tmp12726 := Call(__e, V1387, tmp12725)

				tmp12727 := PrimTail(V1388)

				__e.TailApply(PrimNS2Value(symshen_4call_1help), tmp12726, tmp12727, V1389, V1390)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4call_1help)
				return
			}

		}

	}, 4)

	tmp12730 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call_1help, tmp12722)

	_ = tmp12730

	tmp12731 := MakeNative(func(__e *ControlFlow) {
		V1392 := __e.Get(1)
		_ = V1392
		tmp12747 := PrimIsPair(V1392)

		var ifres12743 Obj

		if True == tmp12747 {
			tmp12745 := PrimHead(V1392)

			tmp12746 := PrimIsPair(tmp12745)

			var ifres12744 Obj

			if True == tmp12746 {
				ifres12744 = True

			} else {
				ifres12744 = False

			}

			ifres12743 = ifres12744

		} else {
			ifres12743 = False

		}

		if True == ifres12743 {
			tmp12733 := MakeNative(func(__e *ControlFlow) {
				ProcessN := __e.Get(1)
				_ = ProcessN
				tmp12734 := PrimHead(V1392)

				tmp12735 := PrimHead(tmp12734)

				tmp12736 := PrimHead(V1392)

				tmp12737 := PrimTail(tmp12736)

				tmp12738 := PrimTail(V1392)

				tmp12739 := PrimCons(tmp12738, Nil)

				tmp12740 := PrimCons(tmp12737, tmp12739)

				tmp12741 := Call(__e, PrimNS2Value(symshen_4insert_1prolog_1variables), tmp12740, ProcessN)

				__e.TailApply(PrimNS2Value(symshen_4intprolog_1help), tmp12735, tmp12741, ProcessN)
				return

			}, 1)

			tmp12742 := Call(__e, PrimNS2Value(symshen_4start_1new_1prolog_1process))

			__e.TailApply(tmp12733, tmp12742)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4intprolog)
			return
		}

	}, 1)

	tmp12748 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog, tmp12731)

	_ = tmp12748

	tmp12749 := MakeNative(func(__e *ControlFlow) {
		V1396 := __e.Get(1)
		_ = V1396
		V1397 := __e.Get(2)
		_ = V1397
		V1398 := __e.Get(3)
		_ = V1398
		tmp12763 := PrimIsPair(V1397)

		var ifres12754 Obj

		if True == tmp12763 {
			tmp12761 := PrimTail(V1397)

			tmp12762 := PrimIsPair(tmp12761)

			var ifres12756 Obj

			if True == tmp12762 {
				tmp12758 := PrimTail(V1397)

				tmp12759 := PrimTail(tmp12758)

				tmp12760 := PrimEqual(Nil, tmp12759)

				var ifres12757 Obj

				if True == tmp12760 {
					ifres12757 = True

				} else {
					ifres12757 = False

				}

				ifres12756 = ifres12757

			} else {
				ifres12756 = False

			}

			var ifres12755 Obj

			if True == ifres12756 {
				ifres12755 = True

			} else {
				ifres12755 = False

			}

			ifres12754 = ifres12755

		} else {
			ifres12754 = False

		}

		if True == ifres12754 {
			tmp12751 := PrimHead(V1397)

			tmp12752 := PrimTail(V1397)

			tmp12753 := PrimHead(tmp12752)

			__e.TailApply(PrimNS2Value(symshen_4intprolog_1help_1help), V1396, tmp12751, tmp12753, V1398)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4intprolog_1help)
			return
		}

	}, 3)

	tmp12764 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog_1help, tmp12749)

	_ = tmp12764

	tmp12765 := MakeNative(func(__e *ControlFlow) {
		V1403 := __e.Get(1)
		_ = V1403
		V1404 := __e.Get(2)
		_ = V1404
		V1405 := __e.Get(3)
		_ = V1405
		V1406 := __e.Get(4)
		_ = V1406
		tmp12773 := PrimEqual(Nil, V1404)

		if True == tmp12773 {
			tmp12772 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(symshen_4call_1rest), V1405, V1406)
				return
			}, 0)

			__e.TailApply(V1403, V1406, tmp12772)
			return

		} else {
			tmp12771 := PrimIsPair(V1404)

			if True == tmp12771 {
				tmp12768 := PrimHead(V1404)

				tmp12769 := Call(__e, V1403, tmp12768)

				tmp12770 := PrimTail(V1404)

				__e.TailApply(PrimNS2Value(symshen_4intprolog_1help_1help), tmp12769, tmp12770, V1405, V1406)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4intprolog_1help_1help)
				return
			}

		}

	}, 4)

	tmp12774 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intprolog_1help_1help, tmp12765)

	_ = tmp12774

	tmp12775 := MakeNative(func(__e *ControlFlow) {
		V1411 := __e.Get(1)
		_ = V1411
		V1412 := __e.Get(2)
		_ = V1412
		tmp12815 := PrimEqual(Nil, V1411)

		if True == tmp12815 {
			__e.Return(True)
			return
		} else {
			tmp12814 := PrimIsPair(V1411)

			var ifres12805 Obj

			if True == tmp12814 {
				tmp12812 := PrimHead(V1411)

				tmp12813 := PrimIsPair(tmp12812)

				var ifres12807 Obj

				if True == tmp12813 {
					tmp12809 := PrimHead(V1411)

					tmp12810 := PrimTail(tmp12809)

					tmp12811 := PrimIsPair(tmp12810)

					var ifres12808 Obj

					if True == tmp12811 {
						ifres12808 = True

					} else {
						ifres12808 = False

					}

					ifres12807 = ifres12808

				} else {
					ifres12807 = False

				}

				var ifres12806 Obj

				if True == ifres12807 {
					ifres12806 = True

				} else {
					ifres12806 = False

				}

				ifres12805 = ifres12806

			} else {
				ifres12805 = False

			}

			if True == ifres12805 {
				tmp12793 := PrimHead(V1411)

				tmp12794 := PrimHead(tmp12793)

				tmp12795 := PrimHead(V1411)

				tmp12796 := PrimTail(tmp12795)

				tmp12797 := PrimHead(tmp12796)

				tmp12798 := Call(__e, tmp12794, tmp12797)

				tmp12799 := PrimHead(V1411)

				tmp12800 := PrimTail(tmp12799)

				tmp12801 := PrimTail(tmp12800)

				tmp12802 := PrimCons(tmp12798, tmp12801)

				tmp12803 := PrimTail(V1411)

				tmp12804 := PrimCons(tmp12802, tmp12803)

				__e.TailApply(PrimNS2Value(symshen_4call_1rest), tmp12804, V1412)
				return

			} else {
				tmp12792 := PrimIsPair(V1411)

				var ifres12783 Obj

				if True == tmp12792 {
					tmp12790 := PrimHead(V1411)

					tmp12791 := PrimIsPair(tmp12790)

					var ifres12785 Obj

					if True == tmp12791 {
						tmp12787 := PrimHead(V1411)

						tmp12788 := PrimTail(tmp12787)

						tmp12789 := PrimEqual(Nil, tmp12788)

						var ifres12786 Obj

						if True == tmp12789 {
							ifres12786 = True

						} else {
							ifres12786 = False

						}

						ifres12785 = ifres12786

					} else {
						ifres12785 = False

					}

					var ifres12784 Obj

					if True == ifres12785 {
						ifres12784 = True

					} else {
						ifres12784 = False

					}

					ifres12783 = ifres12784

				} else {
					ifres12783 = False

				}

				if True == ifres12783 {
					tmp12779 := PrimHead(V1411)

					tmp12780 := PrimHead(tmp12779)

					tmp12781 := MakeNative(func(__e *ControlFlow) {
						tmp12782 := PrimTail(V1411)

						__e.TailApply(PrimNS2Value(symshen_4call_1rest), tmp12782, V1412)
						return

					}, 0)

					__e.TailApply(tmp12780, V1412, tmp12781)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4call_1rest)
					return
				}

			}

		}

	}, 2)

	tmp12816 := Call(__e, PrimNS1Value(symns2_1set), symshen_4call_1rest, tmp12775)

	_ = tmp12816

	tmp12817 := MakeNative(func(__e *ControlFlow) {
		tmp12818 := MakeNative(func(__e *ControlFlow) {
			IncrementProcessCounter := __e.Get(1)
			_ = IncrementProcessCounter
			__e.TailApply(PrimNS2Value(symshen_4initialise_1prolog), IncrementProcessCounter)
			return
		}, 1)

		tmp12819 := PrimNS3Value(symshen_4_dprocess_1counter_d)

		tmp12820 := PrimNumberAdd(MakeNumber(1), tmp12819)

		tmp12821 := PrimNS3Set(symshen_4_dprocess_1counter_d, tmp12820)

		__e.TailApply(tmp12818, tmp12821)
		return

	}, 0)

	tmp12822 := Call(__e, PrimNS1Value(symns2_1set), symshen_4start_1new_1prolog_1process, tmp12817)

	_ = tmp12822

	tmp12823 := MakeNative(func(__e *ControlFlow) {
		V1415 := __e.Get(1)
		_ = V1415
		V1416 := __e.Get(2)
		_ = V1416
		tmp12824 := Call(__e, PrimNS2Value(symshen_4flatten), V1415)

		__e.TailApply(PrimNS2Value(symshen_4insert_1prolog_1variables_1help), V1415, tmp12824, V1416)
		return

	}, 2)

	tmp12825 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1prolog_1variables, tmp12823)

	_ = tmp12825

	tmp12826 := MakeNative(func(__e *ControlFlow) {
		V1424 := __e.Get(1)
		_ = V1424
		V1425 := __e.Get(2)
		_ = V1425
		V1426 := __e.Get(3)
		_ = V1426
		tmp12846 := PrimEqual(Nil, V1425)

		if True == tmp12846 {
			__e.Return(V1424)
			return
		} else {
			tmp12845 := PrimIsPair(V1425)

			var ifres12841 Obj

			if True == tmp12845 {
				tmp12843 := PrimHead(V1425)

				tmp12844 := PrimIsVariable(tmp12843)

				var ifres12842 Obj

				if True == tmp12844 {
					ifres12842 = True

				} else {
					ifres12842 = False

				}

				ifres12841 = ifres12842

			} else {
				ifres12841 = False

			}

			if True == ifres12841 {
				tmp12832 := MakeNative(func(__e *ControlFlow) {
					V := __e.Get(1)
					_ = V
					tmp12833 := MakeNative(func(__e *ControlFlow) {
						XV_cY := __e.Get(1)
						_ = XV_cY
						tmp12834 := MakeNative(func(__e *ControlFlow) {
							Z_1Y := __e.Get(1)
							_ = Z_1Y
							__e.TailApply(PrimNS2Value(symshen_4insert_1prolog_1variables_1help), XV_cY, Z_1Y, V1426)
							return
						}, 1)

						tmp12835 := PrimHead(V1425)

						tmp12836 := PrimTail(V1425)

						tmp12837 := Call(__e, PrimNS2Value(symremove), tmp12835, tmp12836)

						__e.TailApply(tmp12834, tmp12837)
						return

					}, 1)

					tmp12838 := PrimHead(V1425)

					tmp12839 := Call(__e, PrimNS2Value(symsubst), V, tmp12838, V1424)

					__e.TailApply(tmp12833, tmp12839)
					return

				}, 1)

				tmp12840 := Call(__e, PrimNS2Value(symshen_4newpv), V1426)

				__e.TailApply(tmp12832, tmp12840)
				return

			} else {
				tmp12831 := PrimIsPair(V1425)

				if True == tmp12831 {
					tmp12830 := PrimTail(V1425)

					__e.TailApply(PrimNS2Value(symshen_4insert_1prolog_1variables_1help), V1424, tmp12830, V1426)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4insert_1prolog_1variables_1help)
					return
				}

			}

		}

	}, 3)

	tmp12847 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1prolog_1variables_1help, tmp12826)

	_ = tmp12847

	tmp12848 := MakeNative(func(__e *ControlFlow) {
		V1428 := __e.Get(1)
		_ = V1428
		tmp12849 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp12850 := MakeNative(func(__e *ControlFlow) {
				Counter := __e.Get(1)
				_ = Counter
				__e.Return(V1428)
				return
			}, 1)

			tmp12851 := PrimNS3Value(symshen_4_dvarcounter_d)

			tmp12852 := PrimVectorSet(tmp12851, V1428, MakeNumber(1))

			__e.TailApply(tmp12850, tmp12852)
			return

		}, 1)

		tmp12853 := PrimNS3Value(symshen_4_dprologvectors_d)

		tmp12854 := Call(__e, PrimNS2Value(symvector), MakeNumber(10))

		tmp12855 := Call(__e, PrimNS2Value(symshen_4fillvector), tmp12854, MakeNumber(1), MakeNumber(10), symshen_4_1null_1)

		tmp12856 := PrimVectorSet(tmp12853, V1428, tmp12855)

		__e.TailApply(tmp12849, tmp12856)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4initialise_1prolog, tmp12848)
	return

}, 0)
