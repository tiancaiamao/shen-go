package main

import . "github.com/tiancaiamao/shen-go/cora"

var TypesMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp22129 := MakeNative(func(__e *ControlFlow) {
		V3210 := __e.Get(1)
		_ = V3210
		V3211 := __e.Get(2)
		_ = V3211
		tmp22130 := MakeNative(func(__e *ControlFlow) {
			Record := __e.Get(1)
			_ = Record
			tmp22131 := MakeNative(func(__e *ControlFlow) {
				Variancy := __e.Get(1)
				_ = Variancy
				tmp22132 := MakeNative(func(__e *ControlFlow) {
					Type := __e.Get(1)
					_ = Type
					tmp22133 := MakeNative(func(__e *ControlFlow) {
						F_d := __e.Get(1)
						_ = F_d
						tmp22134 := MakeNative(func(__e *ControlFlow) {
							Parameters := __e.Get(1)
							_ = Parameters
							tmp22135 := MakeNative(func(__e *ControlFlow) {
								Clause := __e.Get(1)
								_ = Clause
								tmp22136 := MakeNative(func(__e *ControlFlow) {
									AUM__instruction := __e.Get(1)
									_ = AUM__instruction
									tmp22137 := MakeNative(func(__e *ControlFlow) {
										Code := __e.Get(1)
										_ = Code
										tmp22138 := MakeNative(func(__e *ControlFlow) {
											ShenDef := __e.Get(1)
											_ = ShenDef
											tmp22139 := MakeNative(func(__e *ControlFlow) {
												Eval := __e.Get(1)
												_ = Eval
												__e.Return(V3210)
												return
											}, 1)

											tmp22140 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), ShenDef)

											__e.TailApply(tmp22139, tmp22140)
											return

										}, 1)

										tmp22141 := PrimCons(symContinuation, Nil)

										tmp22142 := PrimCons(symProcessN, tmp22141)

										tmp22143 := PrimCons(Code, Nil)

										tmp22144 := PrimCons(sym_1_6, tmp22143)

										tmp22145 := Call(__e, PrimNS2Value(symappend), tmp22142, tmp22144)

										tmp22146 := Call(__e, PrimNS2Value(symappend), Parameters, tmp22145)

										tmp22147 := PrimCons(F_d, tmp22146)

										tmp22148 := PrimCons(symdefine, tmp22147)

										__e.TailApply(tmp22138, tmp22148)
										return

									}, 1)

									tmp22149 := Call(__e, PrimNS2Value(symshen_4aum__to__shen), AUM__instruction)

									__e.TailApply(tmp22137, tmp22149)
									return

								}, 1)

								tmp22150 := Call(__e, PrimNS2Value(symshen_4aum), Clause, Parameters)

								__e.TailApply(tmp22136, tmp22150)
								return

							}, 1)

							tmp22151 := PrimCons(symX, Nil)

							tmp22152 := PrimCons(F_d, tmp22151)

							tmp22153 := PrimCons(Type, Nil)

							tmp22154 := PrimCons(symX, tmp22153)

							tmp22155 := PrimCons(symunify_b, tmp22154)

							tmp22156 := PrimCons(tmp22155, Nil)

							tmp22157 := PrimCons(tmp22156, Nil)

							tmp22158 := PrimCons(sym_h_1, tmp22157)

							tmp22159 := PrimCons(tmp22152, tmp22158)

							__e.TailApply(tmp22135, tmp22159)
							return

						}, 1)

						tmp22160 := Call(__e, PrimNS2Value(symshen_4parameters), MakeNumber(1))

						__e.TailApply(tmp22134, tmp22160)
						return

					}, 1)

					tmp22161 := Call(__e, PrimNS2Value(symconcat), symshen_4type_1signature_1of_1, V3210)

					__e.TailApply(tmp22133, tmp22161)
					return

				}, 1)

				tmp22162 := Call(__e, PrimNS2Value(symshen_4demodulate), V3211)

				tmp22163 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp22162)

				__e.TailApply(tmp22132, tmp22163)
				return

			}, 1)

			tmp22164 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(symshen_4variancy_1test), V3210, V3211)
				return
			}, 0)

			tmp22165 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4skip)
				return
			}, 1)

			tmp22166 := Call(__e, PrimNS1Value(symtry_1catch), tmp22164, tmp22165)

			__e.TailApply(tmp22131, tmp22166)
			return

		}, 1)

		tmp22167 := PrimCons(V3210, V3211)

		tmp22168 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp22169 := PrimCons(tmp22167, tmp22168)

		tmp22170 := PrimNS3Set(symshen_4_dsignedfuncs_d, tmp22169)

		__e.TailApply(tmp22130, tmp22170)
		return

	}, 2)

	tmp22171 := Call(__e, PrimNS1Value(symns2_1set), symdeclare, tmp22129)

	_ = tmp22171

	tmp22172 := MakeNative(func(__e *ControlFlow) {
		V3213 := __e.Get(1)
		_ = V3213
		tmp22173 := MakeNative(func(__e *ControlFlow) {
			Demod := __e.Get(1)
			_ = Demod
			tmp22175 := PrimEqual(Demod, V3213)

			if True == tmp22175 {
				__e.Return(V3213)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4demodulate), Demod)
				return
			}

		}, 1)

		tmp22176 := PrimNS3Value(symshen_4_ddemodulation_1function_d)

		tmp22177 := Call(__e, PrimNS2Value(symshen_4walk), tmp22176, V3213)

		__e.TailApply(tmp22173, tmp22177)
		return

	}, 1)

	tmp22178 := Call(__e, PrimNS1Value(symns2_1set), symshen_4demodulate, tmp22172)

	_ = tmp22178

	tmp22179 := MakeNative(func(__e *ControlFlow) {
		V3216 := __e.Get(1)
		_ = V3216
		V3217 := __e.Get(2)
		_ = V3217
		tmp22180 := MakeNative(func(__e *ControlFlow) {
			TypeF := __e.Get(1)
			_ = TypeF
			tmp22181 := MakeNative(func(__e *ControlFlow) {
				Check := __e.Get(1)
				_ = Check
				__e.Return(symshen_4skip)
				return
			}, 1)

			tmp22189 := PrimEqual(symsymbol, TypeF)

			var ifres22182 Obj

			if True == tmp22189 {
				ifres22182 = symshen_4skip

			} else {
				tmp22188 := Call(__e, PrimNS2Value(symshen_4variant_2), TypeF, V3217)

				var ifres22183 Obj

				if True == tmp22188 {
					ifres22183 = symshen_4skip

				} else {
					tmp22184 := Call(__e, PrimNS2Value(symshen_4app), V3216, MakeString(" may create errors\n"), symshen_4a)

					tmp22185 := PrimStringConcat(MakeString("warning: changing the type of "), tmp22184)

					tmp22186 := Call(__e, PrimNS2Value(symstoutput))

					tmp22187 := Call(__e, PrimNS2Value(symshen_4prhush), tmp22185, tmp22186)

					ifres22183 = tmp22187

				}

				ifres22182 = ifres22183

			}

			__e.TailApply(tmp22181, ifres22182)
			return

		}, 1)

		tmp22190 := Call(__e, PrimNS2Value(symshen_4typecheck), V3216, symB)

		__e.TailApply(tmp22180, tmp22190)
		return

	}, 2)

	tmp22191 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variancy_1test, tmp22179)

	_ = tmp22191

	tmp22192 := MakeNative(func(__e *ControlFlow) {
		V3230 := __e.Get(1)
		_ = V3230
		V3231 := __e.Get(2)
		_ = V3231
		tmp22244 := PrimEqual(V3231, V3230)

		if True == tmp22244 {
			__e.Return(True)
			return
		} else {
			tmp22243 := PrimIsPair(V3230)

			var ifres22235 Obj

			if True == tmp22243 {
				tmp22242 := PrimIsPair(V3231)

				var ifres22237 Obj

				if True == tmp22242 {
					tmp22239 := PrimHead(V3231)

					tmp22240 := PrimHead(V3230)

					tmp22241 := PrimEqual(tmp22239, tmp22240)

					var ifres22238 Obj

					if True == tmp22241 {
						ifres22238 = True

					} else {
						ifres22238 = False

					}

					ifres22237 = ifres22238

				} else {
					ifres22237 = False

				}

				var ifres22236 Obj

				if True == ifres22237 {
					ifres22236 = True

				} else {
					ifres22236 = False

				}

				ifres22235 = ifres22236

			} else {
				ifres22235 = False

			}

			if True == ifres22235 {
				tmp22233 := PrimTail(V3230)

				tmp22234 := PrimTail(V3231)

				__e.TailApply(PrimNS2Value(symshen_4variant_2), tmp22233, tmp22234)
				return

			} else {
				tmp22232 := PrimIsPair(V3230)

				var ifres22221 Obj

				if True == tmp22232 {
					tmp22231 := PrimIsPair(V3231)

					var ifres22223 Obj

					if True == tmp22231 {
						tmp22229 := PrimHead(V3230)

						tmp22230 := Call(__e, PrimNS2Value(symshen_4pvar_2), tmp22229)

						var ifres22225 Obj

						if True == tmp22230 {
							tmp22227 := PrimHead(V3231)

							tmp22228 := PrimIsVariable(tmp22227)

							var ifres22226 Obj

							if True == tmp22228 {
								ifres22226 = True

							} else {
								ifres22226 = False

							}

							ifres22225 = ifres22226

						} else {
							ifres22225 = False

						}

						var ifres22224 Obj

						if True == ifres22225 {
							ifres22224 = True

						} else {
							ifres22224 = False

						}

						ifres22223 = ifres22224

					} else {
						ifres22223 = False

					}

					var ifres22222 Obj

					if True == ifres22223 {
						ifres22222 = True

					} else {
						ifres22222 = False

					}

					ifres22221 = ifres22222

				} else {
					ifres22221 = False

				}

				if True == ifres22221 {
					tmp22215 := PrimHead(V3230)

					tmp22216 := PrimTail(V3230)

					tmp22217 := Call(__e, PrimNS2Value(symsubst), symshen_4a, tmp22215, tmp22216)

					tmp22218 := PrimHead(V3231)

					tmp22219 := PrimTail(V3231)

					tmp22220 := Call(__e, PrimNS2Value(symsubst), symshen_4a, tmp22218, tmp22219)

					__e.TailApply(PrimNS2Value(symshen_4variant_2), tmp22217, tmp22220)
					return

				} else {
					tmp22214 := PrimIsPair(V3230)

					var ifres22203 Obj

					if True == tmp22214 {
						tmp22212 := PrimHead(V3230)

						tmp22213 := PrimIsPair(tmp22212)

						var ifres22205 Obj

						if True == tmp22213 {
							tmp22211 := PrimIsPair(V3231)

							var ifres22207 Obj

							if True == tmp22211 {
								tmp22209 := PrimHead(V3231)

								tmp22210 := PrimIsPair(tmp22209)

								var ifres22208 Obj

								if True == tmp22210 {
									ifres22208 = True

								} else {
									ifres22208 = False

								}

								ifres22207 = ifres22208

							} else {
								ifres22207 = False

							}

							var ifres22206 Obj

							if True == ifres22207 {
								ifres22206 = True

							} else {
								ifres22206 = False

							}

							ifres22205 = ifres22206

						} else {
							ifres22205 = False

						}

						var ifres22204 Obj

						if True == ifres22205 {
							ifres22204 = True

						} else {
							ifres22204 = False

						}

						ifres22203 = ifres22204

					} else {
						ifres22203 = False

					}

					if True == ifres22203 {
						tmp22197 := PrimHead(V3230)

						tmp22198 := PrimTail(V3230)

						tmp22199 := Call(__e, PrimNS2Value(symappend), tmp22197, tmp22198)

						tmp22200 := PrimHead(V3231)

						tmp22201 := PrimTail(V3231)

						tmp22202 := Call(__e, PrimNS2Value(symappend), tmp22200, tmp22201)

						__e.TailApply(PrimNS2Value(symshen_4variant_2), tmp22199, tmp22202)
						return

					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 2)

	tmp22245 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variant_2, tmp22192)

	_ = tmp22245

	tmp22246 := MakeNative(func(__e *ControlFlow) {
		V3236 := __e.Get(1)
		_ = V3236
		V3237 := __e.Get(2)
		_ = V3237
		V3238 := __e.Get(3)
		_ = V3238
		tmp22247 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22248 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22248

			tmp22249 := PrimCons(symboolean, Nil)

			tmp22250 := PrimCons(sym_1_1_6, tmp22249)

			tmp22251 := PrimCons(A, tmp22250)

			__e.TailApply(PrimNS2Value(symunify_b), V3236, tmp22251, V3237, V3238)
			return

		}, 1)

		tmp22252 := Call(__e, PrimNS2Value(symshen_4newpv), V3237)

		__e.TailApply(tmp22247, tmp22252)
		return

	}, 3)

	tmp22253 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1absvector_2, tmp22246)

	_ = tmp22253

	tmp22254 := MakeNative(func(__e *ControlFlow) {
		V3246 := __e.Get(1)
		_ = V3246
		V3247 := __e.Get(2)
		_ = V3247
		V3248 := __e.Get(3)
		_ = V3248
		tmp22255 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22256 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22256

			tmp22257 := PrimCons(A, Nil)

			tmp22258 := PrimCons(symlist, tmp22257)

			tmp22259 := PrimCons(A, Nil)

			tmp22260 := PrimCons(symlist, tmp22259)

			tmp22261 := PrimCons(tmp22260, Nil)

			tmp22262 := PrimCons(sym_1_1_6, tmp22261)

			tmp22263 := PrimCons(tmp22258, tmp22262)

			tmp22264 := PrimCons(tmp22263, Nil)

			tmp22265 := PrimCons(sym_1_1_6, tmp22264)

			tmp22266 := PrimCons(A, tmp22265)

			__e.TailApply(PrimNS2Value(symunify_b), V3246, tmp22266, V3247, V3248)
			return

		}, 1)

		tmp22267 := Call(__e, PrimNS2Value(symshen_4newpv), V3247)

		__e.TailApply(tmp22255, tmp22267)
		return

	}, 3)

	tmp22268 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1adjoin, tmp22254)

	_ = tmp22268

	tmp22269 := MakeNative(func(__e *ControlFlow) {
		V3256 := __e.Get(1)
		_ = V3256
		V3257 := __e.Get(2)
		_ = V3257
		V3258 := __e.Get(3)
		_ = V3258
		tmp22270 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22270

		tmp22271 := PrimCons(symboolean, Nil)

		tmp22272 := PrimCons(sym_1_1_6, tmp22271)

		tmp22273 := PrimCons(symboolean, tmp22272)

		tmp22274 := PrimCons(tmp22273, Nil)

		tmp22275 := PrimCons(sym_1_1_6, tmp22274)

		tmp22276 := PrimCons(symboolean, tmp22275)

		__e.TailApply(PrimNS2Value(symunify_b), V3256, tmp22276, V3257, V3258)
		return

	}, 3)

	tmp22277 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1and, tmp22269)

	_ = tmp22277

	tmp22278 := MakeNative(func(__e *ControlFlow) {
		V3266 := __e.Get(1)
		_ = V3266
		V3267 := __e.Get(2)
		_ = V3267
		V3268 := __e.Get(3)
		_ = V3268
		tmp22279 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22280 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22280

			tmp22281 := PrimCons(symstring, Nil)

			tmp22282 := PrimCons(sym_1_1_6, tmp22281)

			tmp22283 := PrimCons(symsymbol, tmp22282)

			tmp22284 := PrimCons(tmp22283, Nil)

			tmp22285 := PrimCons(sym_1_1_6, tmp22284)

			tmp22286 := PrimCons(symstring, tmp22285)

			tmp22287 := PrimCons(tmp22286, Nil)

			tmp22288 := PrimCons(sym_1_1_6, tmp22287)

			tmp22289 := PrimCons(A, tmp22288)

			__e.TailApply(PrimNS2Value(symunify_b), V3266, tmp22289, V3267, V3268)
			return

		}, 1)

		tmp22290 := Call(__e, PrimNS2Value(symshen_4newpv), V3267)

		__e.TailApply(tmp22279, tmp22290)
		return

	}, 3)

	tmp22291 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4app, tmp22278)

	_ = tmp22291

	tmp22292 := MakeNative(func(__e *ControlFlow) {
		V3276 := __e.Get(1)
		_ = V3276
		V3277 := __e.Get(2)
		_ = V3277
		V3278 := __e.Get(3)
		_ = V3278
		tmp22293 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22294 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22294

			tmp22295 := PrimCons(A, Nil)

			tmp22296 := PrimCons(symlist, tmp22295)

			tmp22297 := PrimCons(A, Nil)

			tmp22298 := PrimCons(symlist, tmp22297)

			tmp22299 := PrimCons(A, Nil)

			tmp22300 := PrimCons(symlist, tmp22299)

			tmp22301 := PrimCons(tmp22300, Nil)

			tmp22302 := PrimCons(sym_1_1_6, tmp22301)

			tmp22303 := PrimCons(tmp22298, tmp22302)

			tmp22304 := PrimCons(tmp22303, Nil)

			tmp22305 := PrimCons(sym_1_1_6, tmp22304)

			tmp22306 := PrimCons(tmp22296, tmp22305)

			__e.TailApply(PrimNS2Value(symunify_b), V3276, tmp22306, V3277, V3278)
			return

		}, 1)

		tmp22307 := Call(__e, PrimNS2Value(symshen_4newpv), V3277)

		__e.TailApply(tmp22293, tmp22307)
		return

	}, 3)

	tmp22308 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1append, tmp22292)

	_ = tmp22308

	tmp22309 := MakeNative(func(__e *ControlFlow) {
		V3286 := __e.Get(1)
		_ = V3286
		V3287 := __e.Get(2)
		_ = V3287
		V3288 := __e.Get(3)
		_ = V3288
		tmp22310 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22311 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22311

			tmp22312 := PrimCons(symnumber, Nil)

			tmp22313 := PrimCons(sym_1_1_6, tmp22312)

			tmp22314 := PrimCons(A, tmp22313)

			__e.TailApply(PrimNS2Value(symunify_b), V3286, tmp22314, V3287, V3288)
			return

		}, 1)

		tmp22315 := Call(__e, PrimNS2Value(symshen_4newpv), V3287)

		__e.TailApply(tmp22310, tmp22315)
		return

	}, 3)

	tmp22316 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1arity, tmp22309)

	_ = tmp22316

	tmp22317 := MakeNative(func(__e *ControlFlow) {
		V3296 := __e.Get(1)
		_ = V3296
		V3297 := __e.Get(2)
		_ = V3297
		V3298 := __e.Get(3)
		_ = V3298
		tmp22318 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22319 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22319

			tmp22320 := PrimCons(A, Nil)

			tmp22321 := PrimCons(symlist, tmp22320)

			tmp22322 := PrimCons(tmp22321, Nil)

			tmp22323 := PrimCons(symlist, tmp22322)

			tmp22324 := PrimCons(A, Nil)

			tmp22325 := PrimCons(symlist, tmp22324)

			tmp22326 := PrimCons(tmp22325, Nil)

			tmp22327 := PrimCons(sym_1_1_6, tmp22326)

			tmp22328 := PrimCons(tmp22323, tmp22327)

			tmp22329 := PrimCons(tmp22328, Nil)

			tmp22330 := PrimCons(sym_1_1_6, tmp22329)

			tmp22331 := PrimCons(A, tmp22330)

			__e.TailApply(PrimNS2Value(symunify_b), V3296, tmp22331, V3297, V3298)
			return

		}, 1)

		tmp22332 := Call(__e, PrimNS2Value(symshen_4newpv), V3297)

		__e.TailApply(tmp22318, tmp22332)
		return

	}, 3)

	tmp22333 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1assoc, tmp22317)

	_ = tmp22333

	tmp22334 := MakeNative(func(__e *ControlFlow) {
		V3306 := __e.Get(1)
		_ = V3306
		V3307 := __e.Get(2)
		_ = V3307
		V3308 := __e.Get(3)
		_ = V3308
		tmp22335 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22336 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22336

			tmp22337 := PrimCons(symboolean, Nil)

			tmp22338 := PrimCons(sym_1_1_6, tmp22337)

			tmp22339 := PrimCons(A, tmp22338)

			__e.TailApply(PrimNS2Value(symunify_b), V3306, tmp22339, V3307, V3308)
			return

		}, 1)

		tmp22340 := Call(__e, PrimNS2Value(symshen_4newpv), V3307)

		__e.TailApply(tmp22335, tmp22340)
		return

	}, 3)

	tmp22341 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1boolean_2, tmp22334)

	_ = tmp22341

	tmp22342 := MakeNative(func(__e *ControlFlow) {
		V3316 := __e.Get(1)
		_ = V3316
		V3317 := __e.Get(2)
		_ = V3317
		V3318 := __e.Get(3)
		_ = V3318
		tmp22343 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22343

		tmp22344 := PrimCons(symboolean, Nil)

		tmp22345 := PrimCons(sym_1_1_6, tmp22344)

		tmp22346 := PrimCons(symsymbol, tmp22345)

		__e.TailApply(PrimNS2Value(symunify_b), V3316, tmp22346, V3317, V3318)
		return

	}, 3)

	tmp22347 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1bound_2, tmp22342)

	_ = tmp22347

	tmp22348 := MakeNative(func(__e *ControlFlow) {
		V3326 := __e.Get(1)
		_ = V3326
		V3327 := __e.Get(2)
		_ = V3327
		V3328 := __e.Get(3)
		_ = V3328
		tmp22349 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22349

		tmp22350 := PrimCons(symstring, Nil)

		tmp22351 := PrimCons(sym_1_1_6, tmp22350)

		tmp22352 := PrimCons(symstring, tmp22351)

		__e.TailApply(PrimNS2Value(symunify_b), V3326, tmp22352, V3327, V3328)
		return

	}, 3)

	tmp22353 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cd, tmp22348)

	_ = tmp22353

	tmp22354 := MakeNative(func(__e *ControlFlow) {
		V3336 := __e.Get(1)
		_ = V3336
		V3337 := __e.Get(2)
		_ = V3337
		V3338 := __e.Get(3)
		_ = V3338
		tmp22355 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22356 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22357 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22357

				tmp22358 := PrimCons(A, Nil)

				tmp22359 := PrimCons(symstream, tmp22358)

				tmp22360 := PrimCons(B, Nil)

				tmp22361 := PrimCons(symlist, tmp22360)

				tmp22362 := PrimCons(tmp22361, Nil)

				tmp22363 := PrimCons(sym_1_1_6, tmp22362)

				tmp22364 := PrimCons(tmp22359, tmp22363)

				__e.TailApply(PrimNS2Value(symunify_b), V3336, tmp22364, V3337, V3338)
				return

			}, 1)

			tmp22365 := Call(__e, PrimNS2Value(symshen_4newpv), V3337)

			__e.TailApply(tmp22356, tmp22365)
			return

		}, 1)

		tmp22366 := Call(__e, PrimNS2Value(symshen_4newpv), V3337)

		__e.TailApply(tmp22355, tmp22366)
		return

	}, 3)

	tmp22367 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1close, tmp22354)

	_ = tmp22367

	tmp22368 := MakeNative(func(__e *ControlFlow) {
		V3346 := __e.Get(1)
		_ = V3346
		V3347 := __e.Get(2)
		_ = V3347
		V3348 := __e.Get(3)
		_ = V3348
		tmp22369 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22369

		tmp22370 := PrimCons(symstring, Nil)

		tmp22371 := PrimCons(sym_1_1_6, tmp22370)

		tmp22372 := PrimCons(symstring, tmp22371)

		tmp22373 := PrimCons(tmp22372, Nil)

		tmp22374 := PrimCons(sym_1_1_6, tmp22373)

		tmp22375 := PrimCons(symstring, tmp22374)

		__e.TailApply(PrimNS2Value(symunify_b), V3346, tmp22375, V3347, V3348)
		return

	}, 3)

	tmp22376 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cn, tmp22368)

	_ = tmp22376

	tmp22377 := MakeNative(func(__e *ControlFlow) {
		V3356 := __e.Get(1)
		_ = V3356
		V3357 := __e.Get(2)
		_ = V3357
		V3358 := __e.Get(3)
		_ = V3358
		tmp22378 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22379 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22380 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22380

				tmp22381 := PrimCons(B, Nil)

				tmp22382 := PrimCons(symshen_4_a_a_6, tmp22381)

				tmp22383 := PrimCons(A, tmp22382)

				tmp22384 := PrimCons(B, Nil)

				tmp22385 := PrimCons(sym_1_1_6, tmp22384)

				tmp22386 := PrimCons(A, tmp22385)

				tmp22387 := PrimCons(B, Nil)

				tmp22388 := PrimCons(sym_1_1_6, tmp22387)

				tmp22389 := PrimCons(tmp22386, tmp22388)

				tmp22390 := PrimCons(tmp22389, Nil)

				tmp22391 := PrimCons(sym_1_1_6, tmp22390)

				tmp22392 := PrimCons(A, tmp22391)

				tmp22393 := PrimCons(tmp22392, Nil)

				tmp22394 := PrimCons(sym_1_1_6, tmp22393)

				tmp22395 := PrimCons(tmp22383, tmp22394)

				__e.TailApply(PrimNS2Value(symunify_b), V3356, tmp22395, V3357, V3358)
				return

			}, 1)

			tmp22396 := Call(__e, PrimNS2Value(symshen_4newpv), V3357)

			__e.TailApply(tmp22379, tmp22396)
			return

		}, 1)

		tmp22397 := Call(__e, PrimNS2Value(symshen_4newpv), V3357)

		__e.TailApply(tmp22378, tmp22397)
		return

	}, 3)

	tmp22398 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1compile, tmp22377)

	_ = tmp22398

	tmp22399 := MakeNative(func(__e *ControlFlow) {
		V3366 := __e.Get(1)
		_ = V3366
		V3367 := __e.Get(2)
		_ = V3367
		V3368 := __e.Get(3)
		_ = V3368
		tmp22400 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22401 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22401

			tmp22402 := PrimCons(symboolean, Nil)

			tmp22403 := PrimCons(sym_1_1_6, tmp22402)

			tmp22404 := PrimCons(A, tmp22403)

			__e.TailApply(PrimNS2Value(symunify_b), V3366, tmp22404, V3367, V3368)
			return

		}, 1)

		tmp22405 := Call(__e, PrimNS2Value(symshen_4newpv), V3367)

		__e.TailApply(tmp22400, tmp22405)
		return

	}, 3)

	tmp22406 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1cons_2, tmp22399)

	_ = tmp22406

	tmp22407 := MakeNative(func(__e *ControlFlow) {
		V3376 := __e.Get(1)
		_ = V3376
		V3377 := __e.Get(2)
		_ = V3377
		V3378 := __e.Get(3)
		_ = V3378
		tmp22408 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22409 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22410 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22410

				tmp22411 := PrimCons(B, Nil)

				tmp22412 := PrimCons(sym_1_1_6, tmp22411)

				tmp22413 := PrimCons(A, tmp22412)

				tmp22414 := PrimCons(symsymbol, Nil)

				tmp22415 := PrimCons(sym_1_1_6, tmp22414)

				tmp22416 := PrimCons(tmp22413, tmp22415)

				__e.TailApply(PrimNS2Value(symunify_b), V3376, tmp22416, V3377, V3378)
				return

			}, 1)

			tmp22417 := Call(__e, PrimNS2Value(symshen_4newpv), V3377)

			__e.TailApply(tmp22409, tmp22417)
			return

		}, 1)

		tmp22418 := Call(__e, PrimNS2Value(symshen_4newpv), V3377)

		__e.TailApply(tmp22408, tmp22418)
		return

	}, 3)

	tmp22419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1destroy, tmp22407)

	_ = tmp22419

	tmp22420 := MakeNative(func(__e *ControlFlow) {
		V3386 := __e.Get(1)
		_ = V3386
		V3387 := __e.Get(2)
		_ = V3387
		V3388 := __e.Get(3)
		_ = V3388
		tmp22421 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22422 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22422

			tmp22423 := PrimCons(A, Nil)

			tmp22424 := PrimCons(symlist, tmp22423)

			tmp22425 := PrimCons(A, Nil)

			tmp22426 := PrimCons(symlist, tmp22425)

			tmp22427 := PrimCons(A, Nil)

			tmp22428 := PrimCons(symlist, tmp22427)

			tmp22429 := PrimCons(tmp22428, Nil)

			tmp22430 := PrimCons(sym_1_1_6, tmp22429)

			tmp22431 := PrimCons(tmp22426, tmp22430)

			tmp22432 := PrimCons(tmp22431, Nil)

			tmp22433 := PrimCons(sym_1_1_6, tmp22432)

			tmp22434 := PrimCons(tmp22424, tmp22433)

			__e.TailApply(PrimNS2Value(symunify_b), V3386, tmp22434, V3387, V3388)
			return

		}, 1)

		tmp22435 := Call(__e, PrimNS2Value(symshen_4newpv), V3387)

		__e.TailApply(tmp22421, tmp22435)
		return

	}, 3)

	tmp22436 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1difference, tmp22420)

	_ = tmp22436

	tmp22437 := MakeNative(func(__e *ControlFlow) {
		V3396 := __e.Get(1)
		_ = V3396
		V3397 := __e.Get(2)
		_ = V3397
		V3398 := __e.Get(3)
		_ = V3398
		tmp22438 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22439 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22440 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22440

				tmp22441 := PrimCons(B, Nil)

				tmp22442 := PrimCons(sym_1_1_6, tmp22441)

				tmp22443 := PrimCons(B, tmp22442)

				tmp22444 := PrimCons(tmp22443, Nil)

				tmp22445 := PrimCons(sym_1_1_6, tmp22444)

				tmp22446 := PrimCons(A, tmp22445)

				__e.TailApply(PrimNS2Value(symunify_b), V3396, tmp22446, V3397, V3398)
				return

			}, 1)

			tmp22447 := Call(__e, PrimNS2Value(symshen_4newpv), V3397)

			__e.TailApply(tmp22439, tmp22447)
			return

		}, 1)

		tmp22448 := Call(__e, PrimNS2Value(symshen_4newpv), V3397)

		__e.TailApply(tmp22438, tmp22448)
		return

	}, 3)

	tmp22449 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1do, tmp22437)

	_ = tmp22449

	tmp22450 := MakeNative(func(__e *ControlFlow) {
		V3406 := __e.Get(1)
		_ = V3406
		V3407 := __e.Get(2)
		_ = V3407
		V3408 := __e.Get(3)
		_ = V3408
		tmp22451 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22452 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22453 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22453

				tmp22454 := PrimCons(A, Nil)

				tmp22455 := PrimCons(symlist, tmp22454)

				tmp22456 := PrimCons(B, Nil)

				tmp22457 := PrimCons(symlist, tmp22456)

				tmp22458 := PrimCons(tmp22457, Nil)

				tmp22459 := PrimCons(symshen_4_a_a_6, tmp22458)

				tmp22460 := PrimCons(tmp22455, tmp22459)

				__e.TailApply(PrimNS2Value(symunify_b), V3406, tmp22460, V3407, V3408)
				return

			}, 1)

			tmp22461 := Call(__e, PrimNS2Value(symshen_4newpv), V3407)

			__e.TailApply(tmp22452, tmp22461)
			return

		}, 1)

		tmp22462 := Call(__e, PrimNS2Value(symshen_4newpv), V3407)

		__e.TailApply(tmp22451, tmp22462)
		return

	}, 3)

	tmp22463 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5e_6, tmp22450)

	_ = tmp22463

	tmp22464 := MakeNative(func(__e *ControlFlow) {
		V3416 := __e.Get(1)
		_ = V3416
		V3417 := __e.Get(2)
		_ = V3417
		V3418 := __e.Get(3)
		_ = V3418
		tmp22465 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22466 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22466

			tmp22467 := PrimCons(A, Nil)

			tmp22468 := PrimCons(symlist, tmp22467)

			tmp22469 := PrimCons(A, Nil)

			tmp22470 := PrimCons(symlist, tmp22469)

			tmp22471 := PrimCons(tmp22470, Nil)

			tmp22472 := PrimCons(symshen_4_a_a_6, tmp22471)

			tmp22473 := PrimCons(tmp22468, tmp22472)

			__e.TailApply(PrimNS2Value(symunify_b), V3416, tmp22473, V3417, V3418)
			return

		}, 1)

		tmp22474 := Call(__e, PrimNS2Value(symshen_4newpv), V3417)

		__e.TailApply(tmp22465, tmp22474)
		return

	}, 3)

	tmp22475 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_b_6, tmp22464)

	_ = tmp22475

	tmp22476 := MakeNative(func(__e *ControlFlow) {
		V3426 := __e.Get(1)
		_ = V3426
		V3427 := __e.Get(2)
		_ = V3427
		V3428 := __e.Get(3)
		_ = V3428
		tmp22477 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22478 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22478

			tmp22479 := PrimCons(A, Nil)

			tmp22480 := PrimCons(symlist, tmp22479)

			tmp22481 := PrimCons(symboolean, Nil)

			tmp22482 := PrimCons(sym_1_1_6, tmp22481)

			tmp22483 := PrimCons(tmp22480, tmp22482)

			tmp22484 := PrimCons(tmp22483, Nil)

			tmp22485 := PrimCons(sym_1_1_6, tmp22484)

			tmp22486 := PrimCons(A, tmp22485)

			__e.TailApply(PrimNS2Value(symunify_b), V3426, tmp22486, V3427, V3428)
			return

		}, 1)

		tmp22487 := Call(__e, PrimNS2Value(symshen_4newpv), V3427)

		__e.TailApply(tmp22477, tmp22487)
		return

	}, 3)

	tmp22488 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1element_2, tmp22476)

	_ = tmp22488

	tmp22489 := MakeNative(func(__e *ControlFlow) {
		V3436 := __e.Get(1)
		_ = V3436
		V3437 := __e.Get(2)
		_ = V3437
		V3438 := __e.Get(3)
		_ = V3438
		tmp22490 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22491 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22491

			tmp22492 := PrimCons(symboolean, Nil)

			tmp22493 := PrimCons(sym_1_1_6, tmp22492)

			tmp22494 := PrimCons(A, tmp22493)

			__e.TailApply(PrimNS2Value(symunify_b), V3436, tmp22494, V3437, V3438)
			return

		}, 1)

		tmp22495 := Call(__e, PrimNS2Value(symshen_4newpv), V3437)

		__e.TailApply(tmp22490, tmp22495)
		return

	}, 3)

	tmp22496 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1empty_2, tmp22489)

	_ = tmp22496

	tmp22497 := MakeNative(func(__e *ControlFlow) {
		V3446 := __e.Get(1)
		_ = V3446
		V3447 := __e.Get(2)
		_ = V3447
		V3448 := __e.Get(3)
		_ = V3448
		tmp22498 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22498

		tmp22499 := PrimCons(symboolean, Nil)

		tmp22500 := PrimCons(sym_1_1_6, tmp22499)

		tmp22501 := PrimCons(symsymbol, tmp22500)

		__e.TailApply(PrimNS2Value(symunify_b), V3446, tmp22501, V3447, V3448)
		return

	}, 3)

	tmp22502 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1enable_1type_1theory, tmp22497)

	_ = tmp22502

	tmp22503 := MakeNative(func(__e *ControlFlow) {
		V3456 := __e.Get(1)
		_ = V3456
		V3457 := __e.Get(2)
		_ = V3457
		V3458 := __e.Get(3)
		_ = V3458
		tmp22504 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22504

		tmp22505 := PrimCons(symsymbol, Nil)

		tmp22506 := PrimCons(symlist, tmp22505)

		tmp22507 := PrimCons(tmp22506, Nil)

		tmp22508 := PrimCons(sym_1_1_6, tmp22507)

		tmp22509 := PrimCons(symsymbol, tmp22508)

		__e.TailApply(PrimNS2Value(symunify_b), V3456, tmp22509, V3457, V3458)
		return

	}, 3)

	tmp22510 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1external, tmp22503)

	_ = tmp22510

	tmp22511 := MakeNative(func(__e *ControlFlow) {
		V3466 := __e.Get(1)
		_ = V3466
		V3467 := __e.Get(2)
		_ = V3467
		V3468 := __e.Get(3)
		_ = V3468
		tmp22512 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22512

		tmp22513 := PrimCons(symstring, Nil)

		tmp22514 := PrimCons(sym_1_1_6, tmp22513)

		tmp22515 := PrimCons(symexception, tmp22514)

		__e.TailApply(PrimNS2Value(symunify_b), V3466, tmp22515, V3467, V3468)
		return

	}, 3)

	tmp22516 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1error_1to_1string, tmp22511)

	_ = tmp22516

	tmp22517 := MakeNative(func(__e *ControlFlow) {
		V3476 := __e.Get(1)
		_ = V3476
		V3477 := __e.Get(2)
		_ = V3477
		V3478 := __e.Get(3)
		_ = V3478
		tmp22518 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22519 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22519

			tmp22520 := PrimCons(symstring, Nil)

			tmp22521 := PrimCons(symlist, tmp22520)

			tmp22522 := PrimCons(tmp22521, Nil)

			tmp22523 := PrimCons(sym_1_1_6, tmp22522)

			tmp22524 := PrimCons(A, tmp22523)

			__e.TailApply(PrimNS2Value(symunify_b), V3476, tmp22524, V3477, V3478)
			return

		}, 1)

		tmp22525 := Call(__e, PrimNS2Value(symshen_4newpv), V3477)

		__e.TailApply(tmp22518, tmp22525)
		return

	}, 3)

	tmp22526 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1explode, tmp22517)

	_ = tmp22526

	tmp22527 := MakeNative(func(__e *ControlFlow) {
		V3486 := __e.Get(1)
		_ = V3486
		V3487 := __e.Get(2)
		_ = V3487
		V3488 := __e.Get(3)
		_ = V3488
		tmp22528 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22528

		tmp22529 := PrimCons(symsymbol, Nil)

		tmp22530 := PrimCons(sym_1_1_6, tmp22529)

		__e.TailApply(PrimNS2Value(symunify_b), V3486, tmp22530, V3487, V3488)
		return

	}, 3)

	tmp22531 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail, tmp22527)

	_ = tmp22531

	tmp22532 := MakeNative(func(__e *ControlFlow) {
		V3496 := __e.Get(1)
		_ = V3496
		V3497 := __e.Get(2)
		_ = V3497
		V3498 := __e.Get(3)
		_ = V3498
		tmp22533 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22533

		tmp22534 := PrimCons(symboolean, Nil)

		tmp22535 := PrimCons(sym_1_1_6, tmp22534)

		tmp22536 := PrimCons(symsymbol, tmp22535)

		tmp22537 := PrimCons(symsymbol, Nil)

		tmp22538 := PrimCons(sym_1_1_6, tmp22537)

		tmp22539 := PrimCons(symsymbol, tmp22538)

		tmp22540 := PrimCons(tmp22539, Nil)

		tmp22541 := PrimCons(sym_1_1_6, tmp22540)

		tmp22542 := PrimCons(tmp22536, tmp22541)

		__e.TailApply(PrimNS2Value(symunify_b), V3496, tmp22542, V3497, V3498)
		return

	}, 3)

	tmp22543 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fail_1if, tmp22532)

	_ = tmp22543

	tmp22544 := MakeNative(func(__e *ControlFlow) {
		V3506 := __e.Get(1)
		_ = V3506
		V3507 := __e.Get(2)
		_ = V3507
		V3508 := __e.Get(3)
		_ = V3508
		tmp22545 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22546 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22546

			tmp22547 := PrimCons(A, Nil)

			tmp22548 := PrimCons(sym_1_1_6, tmp22547)

			tmp22549 := PrimCons(A, tmp22548)

			tmp22550 := PrimCons(A, Nil)

			tmp22551 := PrimCons(sym_1_1_6, tmp22550)

			tmp22552 := PrimCons(A, tmp22551)

			tmp22553 := PrimCons(tmp22552, Nil)

			tmp22554 := PrimCons(sym_1_1_6, tmp22553)

			tmp22555 := PrimCons(tmp22549, tmp22554)

			__e.TailApply(PrimNS2Value(symunify_b), V3506, tmp22555, V3507, V3508)
			return

		}, 1)

		tmp22556 := Call(__e, PrimNS2Value(symshen_4newpv), V3507)

		__e.TailApply(tmp22545, tmp22556)
		return

	}, 3)

	tmp22557 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fix, tmp22544)

	_ = tmp22557

	tmp22558 := MakeNative(func(__e *ControlFlow) {
		V3516 := __e.Get(1)
		_ = V3516
		V3517 := __e.Get(2)
		_ = V3517
		V3518 := __e.Get(3)
		_ = V3518
		tmp22559 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22560 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22560

			tmp22561 := PrimCons(A, Nil)

			tmp22562 := PrimCons(symlazy, tmp22561)

			tmp22563 := PrimCons(tmp22562, Nil)

			tmp22564 := PrimCons(sym_1_1_6, tmp22563)

			tmp22565 := PrimCons(A, tmp22564)

			__e.TailApply(PrimNS2Value(symunify_b), V3516, tmp22565, V3517, V3518)
			return

		}, 1)

		tmp22566 := Call(__e, PrimNS2Value(symshen_4newpv), V3517)

		__e.TailApply(tmp22559, tmp22566)
		return

	}, 3)

	tmp22567 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1freeze, tmp22558)

	_ = tmp22567

	tmp22568 := MakeNative(func(__e *ControlFlow) {
		V3526 := __e.Get(1)
		_ = V3526
		V3527 := __e.Get(2)
		_ = V3527
		V3528 := __e.Get(3)
		_ = V3528
		tmp22569 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp22570 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				tmp22571 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22571

				tmp22572 := PrimCons(B, Nil)

				tmp22573 := PrimCons(sym_d, tmp22572)

				tmp22574 := PrimCons(A, tmp22573)

				tmp22575 := PrimCons(A, Nil)

				tmp22576 := PrimCons(sym_1_1_6, tmp22575)

				tmp22577 := PrimCons(tmp22574, tmp22576)

				__e.TailApply(PrimNS2Value(symunify_b), V3526, tmp22577, V3527, V3528)
				return

			}, 1)

			tmp22578 := Call(__e, PrimNS2Value(symshen_4newpv), V3527)

			__e.TailApply(tmp22570, tmp22578)
			return

		}, 1)

		tmp22579 := Call(__e, PrimNS2Value(symshen_4newpv), V3527)

		__e.TailApply(tmp22569, tmp22579)
		return

	}, 3)

	tmp22580 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1fst, tmp22568)

	_ = tmp22580

	tmp22581 := MakeNative(func(__e *ControlFlow) {
		V3536 := __e.Get(1)
		_ = V3536
		V3537 := __e.Get(2)
		_ = V3537
		V3538 := __e.Get(3)
		_ = V3538
		tmp22582 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22583 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22584 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22584

				tmp22585 := PrimCons(B, Nil)

				tmp22586 := PrimCons(sym_1_1_6, tmp22585)

				tmp22587 := PrimCons(A, tmp22586)

				tmp22588 := PrimCons(B, Nil)

				tmp22589 := PrimCons(sym_1_1_6, tmp22588)

				tmp22590 := PrimCons(A, tmp22589)

				tmp22591 := PrimCons(tmp22590, Nil)

				tmp22592 := PrimCons(sym_1_1_6, tmp22591)

				tmp22593 := PrimCons(tmp22587, tmp22592)

				__e.TailApply(PrimNS2Value(symunify_b), V3536, tmp22593, V3537, V3538)
				return

			}, 1)

			tmp22594 := Call(__e, PrimNS2Value(symshen_4newpv), V3537)

			__e.TailApply(tmp22583, tmp22594)
			return

		}, 1)

		tmp22595 := Call(__e, PrimNS2Value(symshen_4newpv), V3537)

		__e.TailApply(tmp22582, tmp22595)
		return

	}, 3)

	tmp22596 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1function, tmp22581)

	_ = tmp22596

	tmp22597 := MakeNative(func(__e *ControlFlow) {
		V3546 := __e.Get(1)
		_ = V3546
		V3547 := __e.Get(2)
		_ = V3547
		V3548 := __e.Get(3)
		_ = V3548
		tmp22598 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22598

		tmp22599 := PrimCons(symsymbol, Nil)

		tmp22600 := PrimCons(sym_1_1_6, tmp22599)

		tmp22601 := PrimCons(symsymbol, tmp22600)

		__e.TailApply(PrimNS2Value(symunify_b), V3546, tmp22601, V3547, V3548)
		return

	}, 3)

	tmp22602 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1gensym, tmp22597)

	_ = tmp22602

	tmp22603 := MakeNative(func(__e *ControlFlow) {
		V3556 := __e.Get(1)
		_ = V3556
		V3557 := __e.Get(2)
		_ = V3557
		V3558 := __e.Get(3)
		_ = V3558
		tmp22604 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22605 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22605

			tmp22606 := PrimCons(A, Nil)

			tmp22607 := PrimCons(symvector, tmp22606)

			tmp22608 := PrimCons(A, Nil)

			tmp22609 := PrimCons(sym_1_1_6, tmp22608)

			tmp22610 := PrimCons(symnumber, tmp22609)

			tmp22611 := PrimCons(tmp22610, Nil)

			tmp22612 := PrimCons(sym_1_1_6, tmp22611)

			tmp22613 := PrimCons(tmp22607, tmp22612)

			__e.TailApply(PrimNS2Value(symunify_b), V3556, tmp22613, V3557, V3558)
			return

		}, 1)

		tmp22614 := Call(__e, PrimNS2Value(symshen_4newpv), V3557)

		__e.TailApply(tmp22604, tmp22614)
		return

	}, 3)

	tmp22615 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_1vector, tmp22603)

	_ = tmp22615

	tmp22616 := MakeNative(func(__e *ControlFlow) {
		V3566 := __e.Get(1)
		_ = V3566
		V3567 := __e.Get(2)
		_ = V3567
		V3568 := __e.Get(3)
		_ = V3568
		tmp22617 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22618 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22618

			tmp22619 := PrimCons(A, Nil)

			tmp22620 := PrimCons(symvector, tmp22619)

			tmp22621 := PrimCons(A, Nil)

			tmp22622 := PrimCons(symvector, tmp22621)

			tmp22623 := PrimCons(tmp22622, Nil)

			tmp22624 := PrimCons(sym_1_1_6, tmp22623)

			tmp22625 := PrimCons(A, tmp22624)

			tmp22626 := PrimCons(tmp22625, Nil)

			tmp22627 := PrimCons(sym_1_1_6, tmp22626)

			tmp22628 := PrimCons(symnumber, tmp22627)

			tmp22629 := PrimCons(tmp22628, Nil)

			tmp22630 := PrimCons(sym_1_1_6, tmp22629)

			tmp22631 := PrimCons(tmp22620, tmp22630)

			__e.TailApply(PrimNS2Value(symunify_b), V3566, tmp22631, V3567, V3568)
			return

		}, 1)

		tmp22632 := Call(__e, PrimNS2Value(symshen_4newpv), V3567)

		__e.TailApply(tmp22617, tmp22632)
		return

	}, 3)

	tmp22633 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_1_6, tmp22616)

	_ = tmp22633

	tmp22634 := MakeNative(func(__e *ControlFlow) {
		V3576 := __e.Get(1)
		_ = V3576
		V3577 := __e.Get(2)
		_ = V3577
		V3578 := __e.Get(3)
		_ = V3578
		tmp22635 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22636 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22636

			tmp22637 := PrimCons(A, Nil)

			tmp22638 := PrimCons(symvector, tmp22637)

			tmp22639 := PrimCons(tmp22638, Nil)

			tmp22640 := PrimCons(sym_1_1_6, tmp22639)

			tmp22641 := PrimCons(symnumber, tmp22640)

			__e.TailApply(PrimNS2Value(symunify_b), V3576, tmp22641, V3577, V3578)
			return

		}, 1)

		tmp22642 := Call(__e, PrimNS2Value(symshen_4newpv), V3577)

		__e.TailApply(tmp22635, tmp22642)
		return

	}, 3)

	tmp22643 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector, tmp22634)

	_ = tmp22643

	tmp22644 := MakeNative(func(__e *ControlFlow) {
		V3586 := __e.Get(1)
		_ = V3586
		V3587 := __e.Get(2)
		_ = V3587
		V3588 := __e.Get(3)
		_ = V3588
		tmp22645 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22645

		tmp22646 := PrimCons(symnumber, Nil)

		tmp22647 := PrimCons(sym_1_1_6, tmp22646)

		tmp22648 := PrimCons(symsymbol, tmp22647)

		__e.TailApply(PrimNS2Value(symunify_b), V3586, tmp22648, V3587, V3588)
		return

	}, 3)

	tmp22649 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1get_1time, tmp22644)

	_ = tmp22649

	tmp22650 := MakeNative(func(__e *ControlFlow) {
		V3596 := __e.Get(1)
		_ = V3596
		V3597 := __e.Get(2)
		_ = V3597
		V3598 := __e.Get(3)
		_ = V3598
		tmp22651 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22652 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22652

			tmp22653 := PrimCons(symnumber, Nil)

			tmp22654 := PrimCons(sym_1_1_6, tmp22653)

			tmp22655 := PrimCons(symnumber, tmp22654)

			tmp22656 := PrimCons(tmp22655, Nil)

			tmp22657 := PrimCons(sym_1_1_6, tmp22656)

			tmp22658 := PrimCons(A, tmp22657)

			__e.TailApply(PrimNS2Value(symunify_b), V3596, tmp22658, V3597, V3598)
			return

		}, 1)

		tmp22659 := Call(__e, PrimNS2Value(symshen_4newpv), V3597)

		__e.TailApply(tmp22651, tmp22659)
		return

	}, 3)

	tmp22660 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hash, tmp22650)

	_ = tmp22660

	tmp22661 := MakeNative(func(__e *ControlFlow) {
		V3606 := __e.Get(1)
		_ = V3606
		V3607 := __e.Get(2)
		_ = V3607
		V3608 := __e.Get(3)
		_ = V3608
		tmp22662 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22663 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22663

			tmp22664 := PrimCons(A, Nil)

			tmp22665 := PrimCons(symlist, tmp22664)

			tmp22666 := PrimCons(A, Nil)

			tmp22667 := PrimCons(sym_1_1_6, tmp22666)

			tmp22668 := PrimCons(tmp22665, tmp22667)

			__e.TailApply(PrimNS2Value(symunify_b), V3606, tmp22668, V3607, V3608)
			return

		}, 1)

		tmp22669 := Call(__e, PrimNS2Value(symshen_4newpv), V3607)

		__e.TailApply(tmp22662, tmp22669)
		return

	}, 3)

	tmp22670 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1head, tmp22661)

	_ = tmp22670

	tmp22671 := MakeNative(func(__e *ControlFlow) {
		V3616 := __e.Get(1)
		_ = V3616
		V3617 := __e.Get(2)
		_ = V3617
		V3618 := __e.Get(3)
		_ = V3618
		tmp22672 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22673 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22673

			tmp22674 := PrimCons(A, Nil)

			tmp22675 := PrimCons(symvector, tmp22674)

			tmp22676 := PrimCons(A, Nil)

			tmp22677 := PrimCons(sym_1_1_6, tmp22676)

			tmp22678 := PrimCons(tmp22675, tmp22677)

			__e.TailApply(PrimNS2Value(symunify_b), V3616, tmp22678, V3617, V3618)
			return

		}, 1)

		tmp22679 := Call(__e, PrimNS2Value(symshen_4newpv), V3617)

		__e.TailApply(tmp22672, tmp22679)
		return

	}, 3)

	tmp22680 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdv, tmp22671)

	_ = tmp22680

	tmp22681 := MakeNative(func(__e *ControlFlow) {
		V3626 := __e.Get(1)
		_ = V3626
		V3627 := __e.Get(2)
		_ = V3627
		V3628 := __e.Get(3)
		_ = V3628
		tmp22682 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22682

		tmp22683 := PrimCons(symstring, Nil)

		tmp22684 := PrimCons(sym_1_1_6, tmp22683)

		tmp22685 := PrimCons(symstring, tmp22684)

		__e.TailApply(PrimNS2Value(symunify_b), V3626, tmp22685, V3627, V3628)
		return

	}, 3)

	tmp22686 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1hdstr, tmp22681)

	_ = tmp22686

	tmp22687 := MakeNative(func(__e *ControlFlow) {
		V3636 := __e.Get(1)
		_ = V3636
		V3637 := __e.Get(2)
		_ = V3637
		V3638 := __e.Get(3)
		_ = V3638
		tmp22688 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22689 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22689

			tmp22690 := PrimCons(A, Nil)

			tmp22691 := PrimCons(sym_1_1_6, tmp22690)

			tmp22692 := PrimCons(A, tmp22691)

			tmp22693 := PrimCons(tmp22692, Nil)

			tmp22694 := PrimCons(sym_1_1_6, tmp22693)

			tmp22695 := PrimCons(A, tmp22694)

			tmp22696 := PrimCons(tmp22695, Nil)

			tmp22697 := PrimCons(sym_1_1_6, tmp22696)

			tmp22698 := PrimCons(symboolean, tmp22697)

			__e.TailApply(PrimNS2Value(symunify_b), V3636, tmp22698, V3637, V3638)
			return

		}, 1)

		tmp22699 := Call(__e, PrimNS2Value(symshen_4newpv), V3637)

		__e.TailApply(tmp22688, tmp22699)
		return

	}, 3)

	tmp22700 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1if, tmp22687)

	_ = tmp22700

	tmp22701 := MakeNative(func(__e *ControlFlow) {
		V3646 := __e.Get(1)
		_ = V3646
		V3647 := __e.Get(2)
		_ = V3647
		V3648 := __e.Get(3)
		_ = V3648
		tmp22702 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22702

		tmp22703 := PrimCons(symstring, Nil)

		tmp22704 := PrimCons(sym_1_1_6, tmp22703)

		__e.TailApply(PrimNS2Value(symunify_b), V3646, tmp22704, V3647, V3648)
		return

	}, 3)

	tmp22705 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1it, tmp22701)

	_ = tmp22705

	tmp22706 := MakeNative(func(__e *ControlFlow) {
		V3656 := __e.Get(1)
		_ = V3656
		V3657 := __e.Get(2)
		_ = V3657
		V3658 := __e.Get(3)
		_ = V3658
		tmp22707 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22707

		tmp22708 := PrimCons(symstring, Nil)

		tmp22709 := PrimCons(sym_1_1_6, tmp22708)

		__e.TailApply(PrimNS2Value(symunify_b), V3656, tmp22709, V3657, V3658)
		return

	}, 3)

	tmp22710 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1implementation, tmp22706)

	_ = tmp22710

	tmp22711 := MakeNative(func(__e *ControlFlow) {
		V3666 := __e.Get(1)
		_ = V3666
		V3667 := __e.Get(2)
		_ = V3667
		V3668 := __e.Get(3)
		_ = V3668
		tmp22712 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22712

		tmp22713 := PrimCons(symsymbol, Nil)

		tmp22714 := PrimCons(symlist, tmp22713)

		tmp22715 := PrimCons(symsymbol, Nil)

		tmp22716 := PrimCons(symlist, tmp22715)

		tmp22717 := PrimCons(tmp22716, Nil)

		tmp22718 := PrimCons(sym_1_1_6, tmp22717)

		tmp22719 := PrimCons(tmp22714, tmp22718)

		__e.TailApply(PrimNS2Value(symunify_b), V3666, tmp22719, V3667, V3668)
		return

	}, 3)

	tmp22720 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include, tmp22711)

	_ = tmp22720

	tmp22721 := MakeNative(func(__e *ControlFlow) {
		V3676 := __e.Get(1)
		_ = V3676
		V3677 := __e.Get(2)
		_ = V3677
		V3678 := __e.Get(3)
		_ = V3678
		tmp22722 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22722

		tmp22723 := PrimCons(symsymbol, Nil)

		tmp22724 := PrimCons(symlist, tmp22723)

		tmp22725 := PrimCons(symsymbol, Nil)

		tmp22726 := PrimCons(symlist, tmp22725)

		tmp22727 := PrimCons(tmp22726, Nil)

		tmp22728 := PrimCons(sym_1_1_6, tmp22727)

		tmp22729 := PrimCons(tmp22724, tmp22728)

		__e.TailApply(PrimNS2Value(symunify_b), V3676, tmp22729, V3677, V3678)
		return

	}, 3)

	tmp22730 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1include_1all_1but, tmp22721)

	_ = tmp22730

	tmp22731 := MakeNative(func(__e *ControlFlow) {
		V3686 := __e.Get(1)
		_ = V3686
		V3687 := __e.Get(2)
		_ = V3687
		V3688 := __e.Get(3)
		_ = V3688
		tmp22732 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22732

		tmp22733 := PrimCons(symnumber, Nil)

		tmp22734 := PrimCons(sym_1_1_6, tmp22733)

		__e.TailApply(PrimNS2Value(symunify_b), V3686, tmp22734, V3687, V3688)
		return

	}, 3)

	tmp22735 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1inferences, tmp22731)

	_ = tmp22735

	tmp22736 := MakeNative(func(__e *ControlFlow) {
		V3696 := __e.Get(1)
		_ = V3696
		V3697 := __e.Get(2)
		_ = V3697
		V3698 := __e.Get(3)
		_ = V3698
		tmp22737 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22738 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22738

			tmp22739 := PrimCons(symstring, Nil)

			tmp22740 := PrimCons(sym_1_1_6, tmp22739)

			tmp22741 := PrimCons(symstring, tmp22740)

			tmp22742 := PrimCons(tmp22741, Nil)

			tmp22743 := PrimCons(sym_1_1_6, tmp22742)

			tmp22744 := PrimCons(A, tmp22743)

			__e.TailApply(PrimNS2Value(symunify_b), V3696, tmp22744, V3697, V3698)
			return

		}, 1)

		tmp22745 := Call(__e, PrimNS2Value(symshen_4newpv), V3697)

		__e.TailApply(tmp22737, tmp22745)
		return

	}, 3)

	tmp22746 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4insert, tmp22736)

	_ = tmp22746

	tmp22747 := MakeNative(func(__e *ControlFlow) {
		V3706 := __e.Get(1)
		_ = V3706
		V3707 := __e.Get(2)
		_ = V3707
		V3708 := __e.Get(3)
		_ = V3708
		tmp22748 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22749 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22749

			tmp22750 := PrimCons(symboolean, Nil)

			tmp22751 := PrimCons(sym_1_1_6, tmp22750)

			tmp22752 := PrimCons(A, tmp22751)

			__e.TailApply(PrimNS2Value(symunify_b), V3706, tmp22752, V3707, V3708)
			return

		}, 1)

		tmp22753 := Call(__e, PrimNS2Value(symshen_4newpv), V3707)

		__e.TailApply(tmp22748, tmp22753)
		return

	}, 3)

	tmp22754 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1integer_2, tmp22747)

	_ = tmp22754

	tmp22755 := MakeNative(func(__e *ControlFlow) {
		V3716 := __e.Get(1)
		_ = V3716
		V3717 := __e.Get(2)
		_ = V3717
		V3718 := __e.Get(3)
		_ = V3718
		tmp22756 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22756

		tmp22757 := PrimCons(symsymbol, Nil)

		tmp22758 := PrimCons(symlist, tmp22757)

		tmp22759 := PrimCons(tmp22758, Nil)

		tmp22760 := PrimCons(sym_1_1_6, tmp22759)

		tmp22761 := PrimCons(symsymbol, tmp22760)

		__e.TailApply(PrimNS2Value(symunify_b), V3716, tmp22761, V3717, V3718)
		return

	}, 3)

	tmp22762 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1internal, tmp22755)

	_ = tmp22762

	tmp22763 := MakeNative(func(__e *ControlFlow) {
		V3726 := __e.Get(1)
		_ = V3726
		V3727 := __e.Get(2)
		_ = V3727
		V3728 := __e.Get(3)
		_ = V3728
		tmp22764 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22765 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22765

			tmp22766 := PrimCons(A, Nil)

			tmp22767 := PrimCons(symlist, tmp22766)

			tmp22768 := PrimCons(A, Nil)

			tmp22769 := PrimCons(symlist, tmp22768)

			tmp22770 := PrimCons(A, Nil)

			tmp22771 := PrimCons(symlist, tmp22770)

			tmp22772 := PrimCons(tmp22771, Nil)

			tmp22773 := PrimCons(sym_1_1_6, tmp22772)

			tmp22774 := PrimCons(tmp22769, tmp22773)

			tmp22775 := PrimCons(tmp22774, Nil)

			tmp22776 := PrimCons(sym_1_1_6, tmp22775)

			tmp22777 := PrimCons(tmp22767, tmp22776)

			__e.TailApply(PrimNS2Value(symunify_b), V3726, tmp22777, V3727, V3728)
			return

		}, 1)

		tmp22778 := Call(__e, PrimNS2Value(symshen_4newpv), V3727)

		__e.TailApply(tmp22764, tmp22778)
		return

	}, 3)

	tmp22779 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1intersection, tmp22763)

	_ = tmp22779

	tmp22780 := MakeNative(func(__e *ControlFlow) {
		V3736 := __e.Get(1)
		_ = V3736
		V3737 := __e.Get(2)
		_ = V3737
		V3738 := __e.Get(3)
		_ = V3738
		tmp22781 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22782 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22782

			tmp22783 := PrimCons(A, Nil)

			tmp22784 := PrimCons(sym_1_1_6, tmp22783)

			__e.TailApply(PrimNS2Value(symunify_b), V3736, tmp22784, V3737, V3738)
			return

		}, 1)

		tmp22785 := Call(__e, PrimNS2Value(symshen_4newpv), V3737)

		__e.TailApply(tmp22781, tmp22785)
		return

	}, 3)

	tmp22786 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1kill, tmp22780)

	_ = tmp22786

	tmp22787 := MakeNative(func(__e *ControlFlow) {
		V3746 := __e.Get(1)
		_ = V3746
		V3747 := __e.Get(2)
		_ = V3747
		V3748 := __e.Get(3)
		_ = V3748
		tmp22788 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22788

		tmp22789 := PrimCons(symstring, Nil)

		tmp22790 := PrimCons(sym_1_1_6, tmp22789)

		__e.TailApply(PrimNS2Value(symunify_b), V3746, tmp22790, V3747, V3748)
		return

	}, 3)

	tmp22791 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1language, tmp22787)

	_ = tmp22791

	tmp22792 := MakeNative(func(__e *ControlFlow) {
		V3756 := __e.Get(1)
		_ = V3756
		V3757 := __e.Get(2)
		_ = V3757
		V3758 := __e.Get(3)
		_ = V3758
		tmp22793 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22794 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22794

			tmp22795 := PrimCons(A, Nil)

			tmp22796 := PrimCons(symlist, tmp22795)

			tmp22797 := PrimCons(symnumber, Nil)

			tmp22798 := PrimCons(sym_1_1_6, tmp22797)

			tmp22799 := PrimCons(tmp22796, tmp22798)

			__e.TailApply(PrimNS2Value(symunify_b), V3756, tmp22799, V3757, V3758)
			return

		}, 1)

		tmp22800 := Call(__e, PrimNS2Value(symshen_4newpv), V3757)

		__e.TailApply(tmp22793, tmp22800)
		return

	}, 3)

	tmp22801 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1length, tmp22792)

	_ = tmp22801

	tmp22802 := MakeNative(func(__e *ControlFlow) {
		V3766 := __e.Get(1)
		_ = V3766
		V3767 := __e.Get(2)
		_ = V3767
		V3768 := __e.Get(3)
		_ = V3768
		tmp22803 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22804 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22804

			tmp22805 := PrimCons(A, Nil)

			tmp22806 := PrimCons(symvector, tmp22805)

			tmp22807 := PrimCons(symnumber, Nil)

			tmp22808 := PrimCons(sym_1_1_6, tmp22807)

			tmp22809 := PrimCons(tmp22806, tmp22808)

			__e.TailApply(PrimNS2Value(symunify_b), V3766, tmp22809, V3767, V3768)
			return

		}, 1)

		tmp22810 := Call(__e, PrimNS2Value(symshen_4newpv), V3767)

		__e.TailApply(tmp22803, tmp22810)
		return

	}, 3)

	tmp22811 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1limit, tmp22802)

	_ = tmp22811

	tmp22812 := MakeNative(func(__e *ControlFlow) {
		V3776 := __e.Get(1)
		_ = V3776
		V3777 := __e.Get(2)
		_ = V3777
		V3778 := __e.Get(3)
		_ = V3778
		tmp22813 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22813

		tmp22814 := PrimCons(symsymbol, Nil)

		tmp22815 := PrimCons(sym_1_1_6, tmp22814)

		tmp22816 := PrimCons(symstring, tmp22815)

		__e.TailApply(PrimNS2Value(symunify_b), V3776, tmp22816, V3777, V3778)
		return

	}, 3)

	tmp22817 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1load, tmp22812)

	_ = tmp22817

	tmp22818 := MakeNative(func(__e *ControlFlow) {
		V3786 := __e.Get(1)
		_ = V3786
		V3787 := __e.Get(2)
		_ = V3787
		V3788 := __e.Get(3)
		_ = V3788
		tmp22819 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22820 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22821 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22821

				tmp22822 := PrimCons(B, Nil)

				tmp22823 := PrimCons(sym_1_1_6, tmp22822)

				tmp22824 := PrimCons(A, tmp22823)

				tmp22825 := PrimCons(A, Nil)

				tmp22826 := PrimCons(symlist, tmp22825)

				tmp22827 := PrimCons(B, Nil)

				tmp22828 := PrimCons(symlist, tmp22827)

				tmp22829 := PrimCons(tmp22828, Nil)

				tmp22830 := PrimCons(sym_1_1_6, tmp22829)

				tmp22831 := PrimCons(tmp22826, tmp22830)

				tmp22832 := PrimCons(tmp22831, Nil)

				tmp22833 := PrimCons(sym_1_1_6, tmp22832)

				tmp22834 := PrimCons(tmp22824, tmp22833)

				__e.TailApply(PrimNS2Value(symunify_b), V3786, tmp22834, V3787, V3788)
				return

			}, 1)

			tmp22835 := Call(__e, PrimNS2Value(symshen_4newpv), V3787)

			__e.TailApply(tmp22820, tmp22835)
			return

		}, 1)

		tmp22836 := Call(__e, PrimNS2Value(symshen_4newpv), V3787)

		__e.TailApply(tmp22819, tmp22836)
		return

	}, 3)

	tmp22837 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1map, tmp22818)

	_ = tmp22837

	tmp22838 := MakeNative(func(__e *ControlFlow) {
		V3796 := __e.Get(1)
		_ = V3796
		V3797 := __e.Get(2)
		_ = V3797
		V3798 := __e.Get(3)
		_ = V3798
		tmp22839 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22840 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22841 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22841

				tmp22842 := PrimCons(B, Nil)

				tmp22843 := PrimCons(symlist, tmp22842)

				tmp22844 := PrimCons(tmp22843, Nil)

				tmp22845 := PrimCons(sym_1_1_6, tmp22844)

				tmp22846 := PrimCons(A, tmp22845)

				tmp22847 := PrimCons(A, Nil)

				tmp22848 := PrimCons(symlist, tmp22847)

				tmp22849 := PrimCons(B, Nil)

				tmp22850 := PrimCons(symlist, tmp22849)

				tmp22851 := PrimCons(tmp22850, Nil)

				tmp22852 := PrimCons(sym_1_1_6, tmp22851)

				tmp22853 := PrimCons(tmp22848, tmp22852)

				tmp22854 := PrimCons(tmp22853, Nil)

				tmp22855 := PrimCons(sym_1_1_6, tmp22854)

				tmp22856 := PrimCons(tmp22846, tmp22855)

				__e.TailApply(PrimNS2Value(symunify_b), V3796, tmp22856, V3797, V3798)
				return

			}, 1)

			tmp22857 := Call(__e, PrimNS2Value(symshen_4newpv), V3797)

			__e.TailApply(tmp22840, tmp22857)
			return

		}, 1)

		tmp22858 := Call(__e, PrimNS2Value(symshen_4newpv), V3797)

		__e.TailApply(tmp22839, tmp22858)
		return

	}, 3)

	tmp22859 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1mapcan, tmp22838)

	_ = tmp22859

	tmp22860 := MakeNative(func(__e *ControlFlow) {
		V3806 := __e.Get(1)
		_ = V3806
		V3807 := __e.Get(2)
		_ = V3807
		V3808 := __e.Get(3)
		_ = V3808
		tmp22861 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22861

		tmp22862 := PrimCons(symnumber, Nil)

		tmp22863 := PrimCons(sym_1_1_6, tmp22862)

		tmp22864 := PrimCons(symnumber, tmp22863)

		__e.TailApply(PrimNS2Value(symunify_b), V3806, tmp22864, V3807, V3808)
		return

	}, 3)

	tmp22865 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1maxinferences, tmp22860)

	_ = tmp22865

	tmp22866 := MakeNative(func(__e *ControlFlow) {
		V3816 := __e.Get(1)
		_ = V3816
		V3817 := __e.Get(2)
		_ = V3817
		V3818 := __e.Get(3)
		_ = V3818
		tmp22867 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22867

		tmp22868 := PrimCons(symstring, Nil)

		tmp22869 := PrimCons(sym_1_1_6, tmp22868)

		tmp22870 := PrimCons(symnumber, tmp22869)

		__e.TailApply(PrimNS2Value(symunify_b), V3816, tmp22870, V3817, V3818)
		return

	}, 3)

	tmp22871 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1n_1_6string, tmp22866)

	_ = tmp22871

	tmp22872 := MakeNative(func(__e *ControlFlow) {
		V3826 := __e.Get(1)
		_ = V3826
		V3827 := __e.Get(2)
		_ = V3827
		V3828 := __e.Get(3)
		_ = V3828
		tmp22873 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22873

		tmp22874 := PrimCons(symnumber, Nil)

		tmp22875 := PrimCons(sym_1_1_6, tmp22874)

		tmp22876 := PrimCons(symnumber, tmp22875)

		__e.TailApply(PrimNS2Value(symunify_b), V3826, tmp22876, V3827, V3828)
		return

	}, 3)

	tmp22877 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nl, tmp22872)

	_ = tmp22877

	tmp22878 := MakeNative(func(__e *ControlFlow) {
		V3836 := __e.Get(1)
		_ = V3836
		V3837 := __e.Get(2)
		_ = V3837
		V3838 := __e.Get(3)
		_ = V3838
		tmp22879 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22879

		tmp22880 := PrimCons(symboolean, Nil)

		tmp22881 := PrimCons(sym_1_1_6, tmp22880)

		tmp22882 := PrimCons(symboolean, tmp22881)

		__e.TailApply(PrimNS2Value(symunify_b), V3836, tmp22882, V3837, V3838)
		return

	}, 3)

	tmp22883 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1not, tmp22878)

	_ = tmp22883

	tmp22884 := MakeNative(func(__e *ControlFlow) {
		V3846 := __e.Get(1)
		_ = V3846
		V3847 := __e.Get(2)
		_ = V3847
		V3848 := __e.Get(3)
		_ = V3848
		tmp22885 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22886 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22886

			tmp22887 := PrimCons(A, Nil)

			tmp22888 := PrimCons(symlist, tmp22887)

			tmp22889 := PrimCons(A, Nil)

			tmp22890 := PrimCons(sym_1_1_6, tmp22889)

			tmp22891 := PrimCons(tmp22888, tmp22890)

			tmp22892 := PrimCons(tmp22891, Nil)

			tmp22893 := PrimCons(sym_1_1_6, tmp22892)

			tmp22894 := PrimCons(symnumber, tmp22893)

			__e.TailApply(PrimNS2Value(symunify_b), V3846, tmp22894, V3847, V3848)
			return

		}, 1)

		tmp22895 := Call(__e, PrimNS2Value(symshen_4newpv), V3847)

		__e.TailApply(tmp22885, tmp22895)
		return

	}, 3)

	tmp22896 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1nth, tmp22884)

	_ = tmp22896

	tmp22897 := MakeNative(func(__e *ControlFlow) {
		V3856 := __e.Get(1)
		_ = V3856
		V3857 := __e.Get(2)
		_ = V3857
		V3858 := __e.Get(3)
		_ = V3858
		tmp22898 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22899 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22899

			tmp22900 := PrimCons(symboolean, Nil)

			tmp22901 := PrimCons(sym_1_1_6, tmp22900)

			tmp22902 := PrimCons(A, tmp22901)

			__e.TailApply(PrimNS2Value(symunify_b), V3856, tmp22902, V3857, V3858)
			return

		}, 1)

		tmp22903 := Call(__e, PrimNS2Value(symshen_4newpv), V3857)

		__e.TailApply(tmp22898, tmp22903)
		return

	}, 3)

	tmp22904 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1number_2, tmp22897)

	_ = tmp22904

	tmp22905 := MakeNative(func(__e *ControlFlow) {
		V3866 := __e.Get(1)
		_ = V3866
		V3867 := __e.Get(2)
		_ = V3867
		V3868 := __e.Get(3)
		_ = V3868
		tmp22906 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22907 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22908 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22908

				tmp22909 := PrimCons(symnumber, Nil)

				tmp22910 := PrimCons(sym_1_1_6, tmp22909)

				tmp22911 := PrimCons(B, tmp22910)

				tmp22912 := PrimCons(tmp22911, Nil)

				tmp22913 := PrimCons(sym_1_1_6, tmp22912)

				tmp22914 := PrimCons(A, tmp22913)

				__e.TailApply(PrimNS2Value(symunify_b), V3866, tmp22914, V3867, V3868)
				return

			}, 1)

			tmp22915 := Call(__e, PrimNS2Value(symshen_4newpv), V3867)

			__e.TailApply(tmp22907, tmp22915)
			return

		}, 1)

		tmp22916 := Call(__e, PrimNS2Value(symshen_4newpv), V3867)

		__e.TailApply(tmp22906, tmp22916)
		return

	}, 3)

	tmp22917 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurrences, tmp22905)

	_ = tmp22917

	tmp22918 := MakeNative(func(__e *ControlFlow) {
		V3876 := __e.Get(1)
		_ = V3876
		V3877 := __e.Get(2)
		_ = V3877
		V3878 := __e.Get(3)
		_ = V3878
		tmp22919 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22919

		tmp22920 := PrimCons(symboolean, Nil)

		tmp22921 := PrimCons(sym_1_1_6, tmp22920)

		tmp22922 := PrimCons(symsymbol, tmp22921)

		__e.TailApply(PrimNS2Value(symunify_b), V3876, tmp22922, V3877, V3878)
		return

	}, 3)

	tmp22923 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1occurs_1check, tmp22918)

	_ = tmp22923

	tmp22924 := MakeNative(func(__e *ControlFlow) {
		V3886 := __e.Get(1)
		_ = V3886
		V3887 := __e.Get(2)
		_ = V3887
		V3888 := __e.Get(3)
		_ = V3888
		tmp22925 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22925

		tmp22926 := PrimCons(symboolean, Nil)

		tmp22927 := PrimCons(sym_1_1_6, tmp22926)

		tmp22928 := PrimCons(symsymbol, tmp22927)

		__e.TailApply(PrimNS2Value(symunify_b), V3886, tmp22928, V3887, V3888)
		return

	}, 3)

	tmp22929 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1optimise, tmp22924)

	_ = tmp22929

	tmp22930 := MakeNative(func(__e *ControlFlow) {
		V3896 := __e.Get(1)
		_ = V3896
		V3897 := __e.Get(2)
		_ = V3897
		V3898 := __e.Get(3)
		_ = V3898
		tmp22931 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22931

		tmp22932 := PrimCons(symboolean, Nil)

		tmp22933 := PrimCons(sym_1_1_6, tmp22932)

		tmp22934 := PrimCons(symboolean, tmp22933)

		tmp22935 := PrimCons(tmp22934, Nil)

		tmp22936 := PrimCons(sym_1_1_6, tmp22935)

		tmp22937 := PrimCons(symboolean, tmp22936)

		__e.TailApply(PrimNS2Value(symunify_b), V3896, tmp22937, V3897, V3898)
		return

	}, 3)

	tmp22938 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1or, tmp22930)

	_ = tmp22938

	tmp22939 := MakeNative(func(__e *ControlFlow) {
		V3906 := __e.Get(1)
		_ = V3906
		V3907 := __e.Get(2)
		_ = V3907
		V3908 := __e.Get(3)
		_ = V3908
		tmp22940 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22940

		tmp22941 := PrimCons(symstring, Nil)

		tmp22942 := PrimCons(sym_1_1_6, tmp22941)

		__e.TailApply(PrimNS2Value(symunify_b), V3906, tmp22942, V3907, V3908)
		return

	}, 3)

	tmp22943 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1os, tmp22939)

	_ = tmp22943

	tmp22944 := MakeNative(func(__e *ControlFlow) {
		V3916 := __e.Get(1)
		_ = V3916
		V3917 := __e.Get(2)
		_ = V3917
		V3918 := __e.Get(3)
		_ = V3918
		tmp22945 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22945

		tmp22946 := PrimCons(symboolean, Nil)

		tmp22947 := PrimCons(sym_1_1_6, tmp22946)

		tmp22948 := PrimCons(symsymbol, tmp22947)

		__e.TailApply(PrimNS2Value(symunify_b), V3916, tmp22948, V3917, V3918)
		return

	}, 3)

	tmp22949 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1package_2, tmp22944)

	_ = tmp22949

	tmp22950 := MakeNative(func(__e *ControlFlow) {
		V3926 := __e.Get(1)
		_ = V3926
		V3927 := __e.Get(2)
		_ = V3927
		V3928 := __e.Get(3)
		_ = V3928
		tmp22951 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22951

		tmp22952 := PrimCons(symstring, Nil)

		tmp22953 := PrimCons(sym_1_1_6, tmp22952)

		__e.TailApply(PrimNS2Value(symunify_b), V3926, tmp22953, V3927, V3928)
		return

	}, 3)

	tmp22954 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1port, tmp22950)

	_ = tmp22954

	tmp22955 := MakeNative(func(__e *ControlFlow) {
		V3936 := __e.Get(1)
		_ = V3936
		V3937 := __e.Get(2)
		_ = V3937
		V3938 := __e.Get(3)
		_ = V3938
		tmp22956 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22956

		tmp22957 := PrimCons(symstring, Nil)

		tmp22958 := PrimCons(sym_1_1_6, tmp22957)

		__e.TailApply(PrimNS2Value(symunify_b), V3936, tmp22958, V3937, V3938)
		return

	}, 3)

	tmp22959 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1porters, tmp22955)

	_ = tmp22959

	tmp22960 := MakeNative(func(__e *ControlFlow) {
		V3946 := __e.Get(1)
		_ = V3946
		V3947 := __e.Get(2)
		_ = V3947
		V3948 := __e.Get(3)
		_ = V3948
		tmp22961 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22961

		tmp22962 := PrimCons(symstring, Nil)

		tmp22963 := PrimCons(sym_1_1_6, tmp22962)

		tmp22964 := PrimCons(symnumber, tmp22963)

		tmp22965 := PrimCons(tmp22964, Nil)

		tmp22966 := PrimCons(sym_1_1_6, tmp22965)

		tmp22967 := PrimCons(symstring, tmp22966)

		__e.TailApply(PrimNS2Value(symunify_b), V3946, tmp22967, V3947, V3948)
		return

	}, 3)

	tmp22968 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pos, tmp22960)

	_ = tmp22968

	tmp22969 := MakeNative(func(__e *ControlFlow) {
		V3956 := __e.Get(1)
		_ = V3956
		V3957 := __e.Get(2)
		_ = V3957
		V3958 := __e.Get(3)
		_ = V3958
		tmp22970 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22970

		tmp22971 := PrimCons(symout, Nil)

		tmp22972 := PrimCons(symstream, tmp22971)

		tmp22973 := PrimCons(symstring, Nil)

		tmp22974 := PrimCons(sym_1_1_6, tmp22973)

		tmp22975 := PrimCons(tmp22972, tmp22974)

		tmp22976 := PrimCons(tmp22975, Nil)

		tmp22977 := PrimCons(sym_1_1_6, tmp22976)

		tmp22978 := PrimCons(symstring, tmp22977)

		__e.TailApply(PrimNS2Value(symunify_b), V3956, tmp22978, V3957, V3958)
		return

	}, 3)

	tmp22979 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1pr, tmp22969)

	_ = tmp22979

	tmp22980 := MakeNative(func(__e *ControlFlow) {
		V3966 := __e.Get(1)
		_ = V3966
		V3967 := __e.Get(2)
		_ = V3967
		V3968 := __e.Get(3)
		_ = V3968
		tmp22981 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22982 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22982

			tmp22983 := PrimCons(A, Nil)

			tmp22984 := PrimCons(sym_1_1_6, tmp22983)

			tmp22985 := PrimCons(A, tmp22984)

			__e.TailApply(PrimNS2Value(symunify_b), V3966, tmp22985, V3967, V3968)
			return

		}, 1)

		tmp22986 := Call(__e, PrimNS2Value(symshen_4newpv), V3967)

		__e.TailApply(tmp22981, tmp22986)
		return

	}, 3)

	tmp22987 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1print, tmp22980)

	_ = tmp22987

	tmp22988 := MakeNative(func(__e *ControlFlow) {
		V3976 := __e.Get(1)
		_ = V3976
		V3977 := __e.Get(2)
		_ = V3977
		V3978 := __e.Get(3)
		_ = V3978
		tmp22989 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp22990 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp22991 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22991

				tmp22992 := PrimCons(B, Nil)

				tmp22993 := PrimCons(sym_1_1_6, tmp22992)

				tmp22994 := PrimCons(A, tmp22993)

				tmp22995 := PrimCons(B, Nil)

				tmp22996 := PrimCons(sym_1_1_6, tmp22995)

				tmp22997 := PrimCons(A, tmp22996)

				tmp22998 := PrimCons(tmp22997, Nil)

				tmp22999 := PrimCons(sym_1_1_6, tmp22998)

				tmp23000 := PrimCons(tmp22994, tmp22999)

				__e.TailApply(PrimNS2Value(symunify_b), V3976, tmp23000, V3977, V3978)
				return

			}, 1)

			tmp23001 := Call(__e, PrimNS2Value(symshen_4newpv), V3977)

			__e.TailApply(tmp22990, tmp23001)
			return

		}, 1)

		tmp23002 := Call(__e, PrimNS2Value(symshen_4newpv), V3977)

		__e.TailApply(tmp22989, tmp23002)
		return

	}, 3)

	tmp23003 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile, tmp22988)

	_ = tmp23003

	tmp23004 := MakeNative(func(__e *ControlFlow) {
		V3986 := __e.Get(1)
		_ = V3986
		V3987 := __e.Get(2)
		_ = V3987
		V3988 := __e.Get(3)
		_ = V3988
		tmp23005 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23005

		tmp23006 := PrimCons(symsymbol, Nil)

		tmp23007 := PrimCons(symlist, tmp23006)

		tmp23008 := PrimCons(symsymbol, Nil)

		tmp23009 := PrimCons(symlist, tmp23008)

		tmp23010 := PrimCons(tmp23009, Nil)

		tmp23011 := PrimCons(sym_1_1_6, tmp23010)

		tmp23012 := PrimCons(tmp23007, tmp23011)

		__e.TailApply(PrimNS2Value(symunify_b), V3986, tmp23012, V3987, V3988)
		return

	}, 3)

	tmp23013 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude, tmp23004)

	_ = tmp23013

	tmp23014 := MakeNative(func(__e *ControlFlow) {
		V3996 := __e.Get(1)
		_ = V3996
		V3997 := __e.Get(2)
		_ = V3997
		V3998 := __e.Get(3)
		_ = V3998
		tmp23015 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23015

		tmp23016 := PrimCons(symstring, Nil)

		tmp23017 := PrimCons(sym_1_1_6, tmp23016)

		tmp23018 := PrimCons(symstring, tmp23017)

		__e.TailApply(PrimNS2Value(symunify_b), V3996, tmp23018, V3997, V3998)
		return

	}, 3)

	tmp23019 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4proc_1nl, tmp23014)

	_ = tmp23019

	tmp23020 := MakeNative(func(__e *ControlFlow) {
		V4006 := __e.Get(1)
		_ = V4006
		V4007 := __e.Get(2)
		_ = V4007
		V4008 := __e.Get(3)
		_ = V4008
		tmp23021 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23022 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp23023 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp23023

				tmp23024 := PrimCons(B, Nil)

				tmp23025 := PrimCons(sym_1_1_6, tmp23024)

				tmp23026 := PrimCons(A, tmp23025)

				tmp23027 := PrimCons(B, Nil)

				tmp23028 := PrimCons(sym_1_1_6, tmp23027)

				tmp23029 := PrimCons(A, tmp23028)

				tmp23030 := PrimCons(symnumber, Nil)

				tmp23031 := PrimCons(sym_d, tmp23030)

				tmp23032 := PrimCons(tmp23029, tmp23031)

				tmp23033 := PrimCons(tmp23032, Nil)

				tmp23034 := PrimCons(sym_1_1_6, tmp23033)

				tmp23035 := PrimCons(tmp23026, tmp23034)

				__e.TailApply(PrimNS2Value(symunify_b), V4006, tmp23035, V4007, V4008)
				return

			}, 1)

			tmp23036 := Call(__e, PrimNS2Value(symshen_4newpv), V4007)

			__e.TailApply(tmp23022, tmp23036)
			return

		}, 1)

		tmp23037 := Call(__e, PrimNS2Value(symshen_4newpv), V4007)

		__e.TailApply(tmp23021, tmp23037)
		return

	}, 3)

	tmp23038 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1profile_1results, tmp23020)

	_ = tmp23038

	tmp23039 := MakeNative(func(__e *ControlFlow) {
		V4016 := __e.Get(1)
		_ = V4016
		V4017 := __e.Get(2)
		_ = V4017
		V4018 := __e.Get(3)
		_ = V4018
		tmp23040 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23040

		tmp23041 := PrimCons(symsymbol, Nil)

		tmp23042 := PrimCons(sym_1_1_6, tmp23041)

		tmp23043 := PrimCons(symsymbol, tmp23042)

		__e.TailApply(PrimNS2Value(symunify_b), V4016, tmp23043, V4017, V4018)
		return

	}, 3)

	tmp23044 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1protect, tmp23039)

	_ = tmp23044

	tmp23045 := MakeNative(func(__e *ControlFlow) {
		V4026 := __e.Get(1)
		_ = V4026
		V4027 := __e.Get(2)
		_ = V4027
		V4028 := __e.Get(3)
		_ = V4028
		tmp23046 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23046

		tmp23047 := PrimCons(symsymbol, Nil)

		tmp23048 := PrimCons(symlist, tmp23047)

		tmp23049 := PrimCons(symsymbol, Nil)

		tmp23050 := PrimCons(symlist, tmp23049)

		tmp23051 := PrimCons(tmp23050, Nil)

		tmp23052 := PrimCons(sym_1_1_6, tmp23051)

		tmp23053 := PrimCons(tmp23048, tmp23052)

		__e.TailApply(PrimNS2Value(symunify_b), V4026, tmp23053, V4027, V4028)
		return

	}, 3)

	tmp23054 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1preclude_1all_1but, tmp23045)

	_ = tmp23054

	tmp23055 := MakeNative(func(__e *ControlFlow) {
		V4036 := __e.Get(1)
		_ = V4036
		V4037 := __e.Get(2)
		_ = V4037
		V4038 := __e.Get(3)
		_ = V4038
		tmp23056 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23056

		tmp23057 := PrimCons(symout, Nil)

		tmp23058 := PrimCons(symstream, tmp23057)

		tmp23059 := PrimCons(symstring, Nil)

		tmp23060 := PrimCons(sym_1_1_6, tmp23059)

		tmp23061 := PrimCons(tmp23058, tmp23060)

		tmp23062 := PrimCons(tmp23061, Nil)

		tmp23063 := PrimCons(sym_1_1_6, tmp23062)

		tmp23064 := PrimCons(symstring, tmp23063)

		__e.TailApply(PrimNS2Value(symunify_b), V4036, tmp23064, V4037, V4038)
		return

	}, 3)

	tmp23065 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1shen_4prhush, tmp23055)

	_ = tmp23065

	tmp23066 := MakeNative(func(__e *ControlFlow) {
		V4046 := __e.Get(1)
		_ = V4046
		V4047 := __e.Get(2)
		_ = V4047
		V4048 := __e.Get(3)
		_ = V4048
		tmp23067 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23067

		tmp23068 := PrimCons(symunit, Nil)

		tmp23069 := PrimCons(symlist, tmp23068)

		tmp23070 := PrimCons(tmp23069, Nil)

		tmp23071 := PrimCons(sym_1_1_6, tmp23070)

		tmp23072 := PrimCons(symsymbol, tmp23071)

		__e.TailApply(PrimNS2Value(symunify_b), V4046, tmp23072, V4047, V4048)
		return

	}, 3)

	tmp23073 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1ps, tmp23066)

	_ = tmp23073

	tmp23074 := MakeNative(func(__e *ControlFlow) {
		V4056 := __e.Get(1)
		_ = V4056
		V4057 := __e.Get(2)
		_ = V4057
		V4058 := __e.Get(3)
		_ = V4058
		tmp23075 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23075

		tmp23076 := PrimCons(symin, Nil)

		tmp23077 := PrimCons(symstream, tmp23076)

		tmp23078 := PrimCons(symunit, Nil)

		tmp23079 := PrimCons(sym_1_1_6, tmp23078)

		tmp23080 := PrimCons(tmp23077, tmp23079)

		__e.TailApply(PrimNS2Value(symunify_b), V4056, tmp23080, V4057, V4058)
		return

	}, 3)

	tmp23081 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read, tmp23074)

	_ = tmp23081

	tmp23082 := MakeNative(func(__e *ControlFlow) {
		V4066 := __e.Get(1)
		_ = V4066
		V4067 := __e.Get(2)
		_ = V4067
		V4068 := __e.Get(3)
		_ = V4068
		tmp23083 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23083

		tmp23084 := PrimCons(symin, Nil)

		tmp23085 := PrimCons(symstream, tmp23084)

		tmp23086 := PrimCons(symnumber, Nil)

		tmp23087 := PrimCons(sym_1_1_6, tmp23086)

		tmp23088 := PrimCons(tmp23085, tmp23087)

		__e.TailApply(PrimNS2Value(symunify_b), V4066, tmp23088, V4067, V4068)
		return

	}, 3)

	tmp23089 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1byte, tmp23082)

	_ = tmp23089

	tmp23090 := MakeNative(func(__e *ControlFlow) {
		V4076 := __e.Get(1)
		_ = V4076
		V4077 := __e.Get(2)
		_ = V4077
		V4078 := __e.Get(3)
		_ = V4078
		tmp23091 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23091

		tmp23092 := PrimCons(symnumber, Nil)

		tmp23093 := PrimCons(symlist, tmp23092)

		tmp23094 := PrimCons(tmp23093, Nil)

		tmp23095 := PrimCons(sym_1_1_6, tmp23094)

		tmp23096 := PrimCons(symstring, tmp23095)

		__e.TailApply(PrimNS2Value(symunify_b), V4076, tmp23096, V4077, V4078)
		return

	}, 3)

	tmp23097 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1bytelist, tmp23090)

	_ = tmp23097

	tmp23098 := MakeNative(func(__e *ControlFlow) {
		V4086 := __e.Get(1)
		_ = V4086
		V4087 := __e.Get(2)
		_ = V4087
		V4088 := __e.Get(3)
		_ = V4088
		tmp23099 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23099

		tmp23100 := PrimCons(symstring, Nil)

		tmp23101 := PrimCons(sym_1_1_6, tmp23100)

		tmp23102 := PrimCons(symstring, tmp23101)

		__e.TailApply(PrimNS2Value(symunify_b), V4086, tmp23102, V4087, V4088)
		return

	}, 3)

	tmp23103 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file_1as_1string, tmp23098)

	_ = tmp23103

	tmp23104 := MakeNative(func(__e *ControlFlow) {
		V4096 := __e.Get(1)
		_ = V4096
		V4097 := __e.Get(2)
		_ = V4097
		V4098 := __e.Get(3)
		_ = V4098
		tmp23105 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23105

		tmp23106 := PrimCons(symunit, Nil)

		tmp23107 := PrimCons(symlist, tmp23106)

		tmp23108 := PrimCons(tmp23107, Nil)

		tmp23109 := PrimCons(sym_1_1_6, tmp23108)

		tmp23110 := PrimCons(symstring, tmp23109)

		__e.TailApply(PrimNS2Value(symunify_b), V4096, tmp23110, V4097, V4098)
		return

	}, 3)

	tmp23111 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1file, tmp23104)

	_ = tmp23111

	tmp23112 := MakeNative(func(__e *ControlFlow) {
		V4106 := __e.Get(1)
		_ = V4106
		V4107 := __e.Get(2)
		_ = V4107
		V4108 := __e.Get(3)
		_ = V4108
		tmp23113 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23113

		tmp23114 := PrimCons(symunit, Nil)

		tmp23115 := PrimCons(symlist, tmp23114)

		tmp23116 := PrimCons(tmp23115, Nil)

		tmp23117 := PrimCons(sym_1_1_6, tmp23116)

		tmp23118 := PrimCons(symstring, tmp23117)

		__e.TailApply(PrimNS2Value(symunify_b), V4106, tmp23118, V4107, V4108)
		return

	}, 3)

	tmp23119 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1read_1from_1string, tmp23112)

	_ = tmp23119

	tmp23120 := MakeNative(func(__e *ControlFlow) {
		V4116 := __e.Get(1)
		_ = V4116
		V4117 := __e.Get(2)
		_ = V4117
		V4118 := __e.Get(3)
		_ = V4118
		tmp23121 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23121

		tmp23122 := PrimCons(symstring, Nil)

		tmp23123 := PrimCons(sym_1_1_6, tmp23122)

		__e.TailApply(PrimNS2Value(symunify_b), V4116, tmp23123, V4117, V4118)
		return

	}, 3)

	tmp23124 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1release, tmp23120)

	_ = tmp23124

	tmp23125 := MakeNative(func(__e *ControlFlow) {
		V4126 := __e.Get(1)
		_ = V4126
		V4127 := __e.Get(2)
		_ = V4127
		V4128 := __e.Get(3)
		_ = V4128
		tmp23126 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23127 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23127

			tmp23128 := PrimCons(A, Nil)

			tmp23129 := PrimCons(symlist, tmp23128)

			tmp23130 := PrimCons(A, Nil)

			tmp23131 := PrimCons(symlist, tmp23130)

			tmp23132 := PrimCons(tmp23131, Nil)

			tmp23133 := PrimCons(sym_1_1_6, tmp23132)

			tmp23134 := PrimCons(tmp23129, tmp23133)

			tmp23135 := PrimCons(tmp23134, Nil)

			tmp23136 := PrimCons(sym_1_1_6, tmp23135)

			tmp23137 := PrimCons(A, tmp23136)

			__e.TailApply(PrimNS2Value(symunify_b), V4126, tmp23137, V4127, V4128)
			return

		}, 1)

		tmp23138 := Call(__e, PrimNS2Value(symshen_4newpv), V4127)

		__e.TailApply(tmp23126, tmp23138)
		return

	}, 3)

	tmp23139 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1remove, tmp23125)

	_ = tmp23139

	tmp23140 := MakeNative(func(__e *ControlFlow) {
		V4136 := __e.Get(1)
		_ = V4136
		V4137 := __e.Get(2)
		_ = V4137
		V4138 := __e.Get(3)
		_ = V4138
		tmp23141 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23142 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23142

			tmp23143 := PrimCons(A, Nil)

			tmp23144 := PrimCons(symlist, tmp23143)

			tmp23145 := PrimCons(A, Nil)

			tmp23146 := PrimCons(symlist, tmp23145)

			tmp23147 := PrimCons(tmp23146, Nil)

			tmp23148 := PrimCons(sym_1_1_6, tmp23147)

			tmp23149 := PrimCons(tmp23144, tmp23148)

			__e.TailApply(PrimNS2Value(symunify_b), V4136, tmp23149, V4137, V4138)
			return

		}, 1)

		tmp23150 := Call(__e, PrimNS2Value(symshen_4newpv), V4137)

		__e.TailApply(tmp23141, tmp23150)
		return

	}, 3)

	tmp23151 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1reverse, tmp23140)

	_ = tmp23151

	tmp23152 := MakeNative(func(__e *ControlFlow) {
		V4146 := __e.Get(1)
		_ = V4146
		V4147 := __e.Get(2)
		_ = V4147
		V4148 := __e.Get(3)
		_ = V4148
		tmp23153 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23154 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23154

			tmp23155 := PrimCons(A, Nil)

			tmp23156 := PrimCons(sym_1_1_6, tmp23155)

			tmp23157 := PrimCons(symstring, tmp23156)

			__e.TailApply(PrimNS2Value(symunify_b), V4146, tmp23157, V4147, V4148)
			return

		}, 1)

		tmp23158 := Call(__e, PrimNS2Value(symshen_4newpv), V4147)

		__e.TailApply(tmp23153, tmp23158)
		return

	}, 3)

	tmp23159 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1simple_1error, tmp23152)

	_ = tmp23159

	tmp23160 := MakeNative(func(__e *ControlFlow) {
		V4156 := __e.Get(1)
		_ = V4156
		V4157 := __e.Get(2)
		_ = V4157
		V4158 := __e.Get(3)
		_ = V4158
		tmp23161 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23162 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp23163 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp23163

				tmp23164 := PrimCons(B, Nil)

				tmp23165 := PrimCons(sym_d, tmp23164)

				tmp23166 := PrimCons(A, tmp23165)

				tmp23167 := PrimCons(B, Nil)

				tmp23168 := PrimCons(sym_1_1_6, tmp23167)

				tmp23169 := PrimCons(tmp23166, tmp23168)

				__e.TailApply(PrimNS2Value(symunify_b), V4156, tmp23169, V4157, V4158)
				return

			}, 1)

			tmp23170 := Call(__e, PrimNS2Value(symshen_4newpv), V4157)

			__e.TailApply(tmp23162, tmp23170)
			return

		}, 1)

		tmp23171 := Call(__e, PrimNS2Value(symshen_4newpv), V4157)

		__e.TailApply(tmp23161, tmp23171)
		return

	}, 3)

	tmp23172 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1snd, tmp23160)

	_ = tmp23172

	tmp23173 := MakeNative(func(__e *ControlFlow) {
		V4166 := __e.Get(1)
		_ = V4166
		V4167 := __e.Get(2)
		_ = V4167
		V4168 := __e.Get(3)
		_ = V4168
		tmp23174 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23174

		tmp23175 := PrimCons(symsymbol, Nil)

		tmp23176 := PrimCons(sym_1_1_6, tmp23175)

		tmp23177 := PrimCons(symsymbol, tmp23176)

		__e.TailApply(PrimNS2Value(symunify_b), V4166, tmp23177, V4167, V4168)
		return

	}, 3)

	tmp23178 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1specialise, tmp23173)

	_ = tmp23178

	tmp23179 := MakeNative(func(__e *ControlFlow) {
		V4176 := __e.Get(1)
		_ = V4176
		V4177 := __e.Get(2)
		_ = V4177
		V4178 := __e.Get(3)
		_ = V4178
		tmp23180 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23180

		tmp23181 := PrimCons(symboolean, Nil)

		tmp23182 := PrimCons(sym_1_1_6, tmp23181)

		tmp23183 := PrimCons(symsymbol, tmp23182)

		__e.TailApply(PrimNS2Value(symunify_b), V4176, tmp23183, V4177, V4178)
		return

	}, 3)

	tmp23184 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1spy, tmp23179)

	_ = tmp23184

	tmp23185 := MakeNative(func(__e *ControlFlow) {
		V4186 := __e.Get(1)
		_ = V4186
		V4187 := __e.Get(2)
		_ = V4187
		V4188 := __e.Get(3)
		_ = V4188
		tmp23186 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23186

		tmp23187 := PrimCons(symboolean, Nil)

		tmp23188 := PrimCons(sym_1_1_6, tmp23187)

		tmp23189 := PrimCons(symsymbol, tmp23188)

		__e.TailApply(PrimNS2Value(symunify_b), V4186, tmp23189, V4187, V4188)
		return

	}, 3)

	tmp23190 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1step, tmp23185)

	_ = tmp23190

	tmp23191 := MakeNative(func(__e *ControlFlow) {
		V4196 := __e.Get(1)
		_ = V4196
		V4197 := __e.Get(2)
		_ = V4197
		V4198 := __e.Get(3)
		_ = V4198
		tmp23192 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23192

		tmp23193 := PrimCons(symin, Nil)

		tmp23194 := PrimCons(symstream, tmp23193)

		tmp23195 := PrimCons(tmp23194, Nil)

		tmp23196 := PrimCons(sym_1_1_6, tmp23195)

		__e.TailApply(PrimNS2Value(symunify_b), V4196, tmp23196, V4197, V4198)
		return

	}, 3)

	tmp23197 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stinput, tmp23191)

	_ = tmp23197

	tmp23198 := MakeNative(func(__e *ControlFlow) {
		V4206 := __e.Get(1)
		_ = V4206
		V4207 := __e.Get(2)
		_ = V4207
		V4208 := __e.Get(3)
		_ = V4208
		tmp23199 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23199

		tmp23200 := PrimCons(symout, Nil)

		tmp23201 := PrimCons(symstream, tmp23200)

		tmp23202 := PrimCons(tmp23201, Nil)

		tmp23203 := PrimCons(sym_1_1_6, tmp23202)

		__e.TailApply(PrimNS2Value(symunify_b), V4206, tmp23203, V4207, V4208)
		return

	}, 3)

	tmp23204 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sterror, tmp23198)

	_ = tmp23204

	tmp23205 := MakeNative(func(__e *ControlFlow) {
		V4216 := __e.Get(1)
		_ = V4216
		V4217 := __e.Get(2)
		_ = V4217
		V4218 := __e.Get(3)
		_ = V4218
		tmp23206 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23206

		tmp23207 := PrimCons(symout, Nil)

		tmp23208 := PrimCons(symstream, tmp23207)

		tmp23209 := PrimCons(tmp23208, Nil)

		tmp23210 := PrimCons(sym_1_1_6, tmp23209)

		__e.TailApply(PrimNS2Value(symunify_b), V4216, tmp23210, V4217, V4218)
		return

	}, 3)

	tmp23211 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1stoutput, tmp23205)

	_ = tmp23211

	tmp23212 := MakeNative(func(__e *ControlFlow) {
		V4226 := __e.Get(1)
		_ = V4226
		V4227 := __e.Get(2)
		_ = V4227
		V4228 := __e.Get(3)
		_ = V4228
		tmp23213 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23214 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23214

			tmp23215 := PrimCons(symboolean, Nil)

			tmp23216 := PrimCons(sym_1_1_6, tmp23215)

			tmp23217 := PrimCons(A, tmp23216)

			__e.TailApply(PrimNS2Value(symunify_b), V4226, tmp23217, V4227, V4228)
			return

		}, 1)

		tmp23218 := Call(__e, PrimNS2Value(symshen_4newpv), V4227)

		__e.TailApply(tmp23213, tmp23218)
		return

	}, 3)

	tmp23219 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_2, tmp23212)

	_ = tmp23219

	tmp23220 := MakeNative(func(__e *ControlFlow) {
		V4236 := __e.Get(1)
		_ = V4236
		V4237 := __e.Get(2)
		_ = V4237
		V4238 := __e.Get(3)
		_ = V4238
		tmp23221 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23222 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23222

			tmp23223 := PrimCons(symstring, Nil)

			tmp23224 := PrimCons(sym_1_1_6, tmp23223)

			tmp23225 := PrimCons(A, tmp23224)

			__e.TailApply(PrimNS2Value(symunify_b), V4236, tmp23225, V4237, V4238)
			return

		}, 1)

		tmp23226 := Call(__e, PrimNS2Value(symshen_4newpv), V4237)

		__e.TailApply(tmp23221, tmp23226)
		return

	}, 3)

	tmp23227 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1str, tmp23220)

	_ = tmp23227

	tmp23228 := MakeNative(func(__e *ControlFlow) {
		V4246 := __e.Get(1)
		_ = V4246
		V4247 := __e.Get(2)
		_ = V4247
		V4248 := __e.Get(3)
		_ = V4248
		tmp23229 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23229

		tmp23230 := PrimCons(symnumber, Nil)

		tmp23231 := PrimCons(sym_1_1_6, tmp23230)

		tmp23232 := PrimCons(symstring, tmp23231)

		__e.TailApply(PrimNS2Value(symunify_b), V4246, tmp23232, V4247, V4248)
		return

	}, 3)

	tmp23233 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6n, tmp23228)

	_ = tmp23233

	tmp23234 := MakeNative(func(__e *ControlFlow) {
		V4256 := __e.Get(1)
		_ = V4256
		V4257 := __e.Get(2)
		_ = V4257
		V4258 := __e.Get(3)
		_ = V4258
		tmp23235 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23235

		tmp23236 := PrimCons(symsymbol, Nil)

		tmp23237 := PrimCons(sym_1_1_6, tmp23236)

		tmp23238 := PrimCons(symstring, tmp23237)

		__e.TailApply(PrimNS2Value(symunify_b), V4256, tmp23238, V4257, V4258)
		return

	}, 3)

	tmp23239 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1string_1_6symbol, tmp23234)

	_ = tmp23239

	tmp23240 := MakeNative(func(__e *ControlFlow) {
		V4266 := __e.Get(1)
		_ = V4266
		V4267 := __e.Get(2)
		_ = V4267
		V4268 := __e.Get(3)
		_ = V4268
		tmp23241 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23241

		tmp23242 := PrimCons(symnumber, Nil)

		tmp23243 := PrimCons(symlist, tmp23242)

		tmp23244 := PrimCons(symnumber, Nil)

		tmp23245 := PrimCons(sym_1_1_6, tmp23244)

		tmp23246 := PrimCons(tmp23243, tmp23245)

		__e.TailApply(PrimNS2Value(symunify_b), V4266, tmp23246, V4267, V4268)
		return

	}, 3)

	tmp23247 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1sum, tmp23240)

	_ = tmp23247

	tmp23248 := MakeNative(func(__e *ControlFlow) {
		V4276 := __e.Get(1)
		_ = V4276
		V4277 := __e.Get(2)
		_ = V4277
		V4278 := __e.Get(3)
		_ = V4278
		tmp23249 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23250 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23250

			tmp23251 := PrimCons(symboolean, Nil)

			tmp23252 := PrimCons(sym_1_1_6, tmp23251)

			tmp23253 := PrimCons(A, tmp23252)

			__e.TailApply(PrimNS2Value(symunify_b), V4276, tmp23253, V4277, V4278)
			return

		}, 1)

		tmp23254 := Call(__e, PrimNS2Value(symshen_4newpv), V4277)

		__e.TailApply(tmp23249, tmp23254)
		return

	}, 3)

	tmp23255 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1symbol_2, tmp23248)

	_ = tmp23255

	tmp23256 := MakeNative(func(__e *ControlFlow) {
		V4286 := __e.Get(1)
		_ = V4286
		V4287 := __e.Get(2)
		_ = V4287
		V4288 := __e.Get(3)
		_ = V4288
		tmp23257 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23257

		tmp23258 := PrimCons(symsymbol, Nil)

		tmp23259 := PrimCons(sym_1_1_6, tmp23258)

		tmp23260 := PrimCons(symsymbol, tmp23259)

		__e.TailApply(PrimNS2Value(symunify_b), V4286, tmp23260, V4287, V4288)
		return

	}, 3)

	tmp23261 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1systemf, tmp23256)

	_ = tmp23261

	tmp23262 := MakeNative(func(__e *ControlFlow) {
		V4296 := __e.Get(1)
		_ = V4296
		V4297 := __e.Get(2)
		_ = V4297
		V4298 := __e.Get(3)
		_ = V4298
		tmp23263 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23264 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23264

			tmp23265 := PrimCons(A, Nil)

			tmp23266 := PrimCons(symlist, tmp23265)

			tmp23267 := PrimCons(A, Nil)

			tmp23268 := PrimCons(symlist, tmp23267)

			tmp23269 := PrimCons(tmp23268, Nil)

			tmp23270 := PrimCons(sym_1_1_6, tmp23269)

			tmp23271 := PrimCons(tmp23266, tmp23270)

			__e.TailApply(PrimNS2Value(symunify_b), V4296, tmp23271, V4297, V4298)
			return

		}, 1)

		tmp23272 := Call(__e, PrimNS2Value(symshen_4newpv), V4297)

		__e.TailApply(tmp23263, tmp23272)
		return

	}, 3)

	tmp23273 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tail, tmp23262)

	_ = tmp23273

	tmp23274 := MakeNative(func(__e *ControlFlow) {
		V4306 := __e.Get(1)
		_ = V4306
		V4307 := __e.Get(2)
		_ = V4307
		V4308 := __e.Get(3)
		_ = V4308
		tmp23275 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23275

		tmp23276 := PrimCons(symstring, Nil)

		tmp23277 := PrimCons(sym_1_1_6, tmp23276)

		tmp23278 := PrimCons(symstring, tmp23277)

		__e.TailApply(PrimNS2Value(symunify_b), V4306, tmp23278, V4307, V4308)
		return

	}, 3)

	tmp23279 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlstr, tmp23274)

	_ = tmp23279

	tmp23280 := MakeNative(func(__e *ControlFlow) {
		V4316 := __e.Get(1)
		_ = V4316
		V4317 := __e.Get(2)
		_ = V4317
		V4318 := __e.Get(3)
		_ = V4318
		tmp23281 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23282 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23282

			tmp23283 := PrimCons(A, Nil)

			tmp23284 := PrimCons(symvector, tmp23283)

			tmp23285 := PrimCons(A, Nil)

			tmp23286 := PrimCons(symvector, tmp23285)

			tmp23287 := PrimCons(tmp23286, Nil)

			tmp23288 := PrimCons(sym_1_1_6, tmp23287)

			tmp23289 := PrimCons(tmp23284, tmp23288)

			__e.TailApply(PrimNS2Value(symunify_b), V4316, tmp23289, V4317, V4318)
			return

		}, 1)

		tmp23290 := Call(__e, PrimNS2Value(symshen_4newpv), V4317)

		__e.TailApply(tmp23281, tmp23290)
		return

	}, 3)

	tmp23291 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tlv, tmp23280)

	_ = tmp23291

	tmp23292 := MakeNative(func(__e *ControlFlow) {
		V4326 := __e.Get(1)
		_ = V4326
		V4327 := __e.Get(2)
		_ = V4327
		V4328 := __e.Get(3)
		_ = V4328
		tmp23293 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23293

		tmp23294 := PrimCons(symboolean, Nil)

		tmp23295 := PrimCons(sym_1_1_6, tmp23294)

		tmp23296 := PrimCons(symsymbol, tmp23295)

		__e.TailApply(PrimNS2Value(symunify_b), V4326, tmp23296, V4327, V4328)
		return

	}, 3)

	tmp23297 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc, tmp23292)

	_ = tmp23297

	tmp23298 := MakeNative(func(__e *ControlFlow) {
		V4336 := __e.Get(1)
		_ = V4336
		V4337 := __e.Get(2)
		_ = V4337
		V4338 := __e.Get(3)
		_ = V4338
		tmp23299 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23299

		tmp23300 := PrimCons(symboolean, Nil)

		tmp23301 := PrimCons(sym_1_1_6, tmp23300)

		__e.TailApply(PrimNS2Value(symunify_b), V4336, tmp23301, V4337, V4338)
		return

	}, 3)

	tmp23302 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tc_2, tmp23298)

	_ = tmp23302

	tmp23303 := MakeNative(func(__e *ControlFlow) {
		V4346 := __e.Get(1)
		_ = V4346
		V4347 := __e.Get(2)
		_ = V4347
		V4348 := __e.Get(3)
		_ = V4348
		tmp23304 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23305 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23305

			tmp23306 := PrimCons(A, Nil)

			tmp23307 := PrimCons(symlazy, tmp23306)

			tmp23308 := PrimCons(A, Nil)

			tmp23309 := PrimCons(sym_1_1_6, tmp23308)

			tmp23310 := PrimCons(tmp23307, tmp23309)

			__e.TailApply(PrimNS2Value(symunify_b), V4346, tmp23310, V4347, V4348)
			return

		}, 1)

		tmp23311 := Call(__e, PrimNS2Value(symshen_4newpv), V4347)

		__e.TailApply(tmp23304, tmp23311)
		return

	}, 3)

	tmp23312 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1thaw, tmp23303)

	_ = tmp23312

	tmp23313 := MakeNative(func(__e *ControlFlow) {
		V4356 := __e.Get(1)
		_ = V4356
		V4357 := __e.Get(2)
		_ = V4357
		V4358 := __e.Get(3)
		_ = V4358
		tmp23314 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23314

		tmp23315 := PrimCons(symsymbol, Nil)

		tmp23316 := PrimCons(sym_1_1_6, tmp23315)

		tmp23317 := PrimCons(symsymbol, tmp23316)

		__e.TailApply(PrimNS2Value(symunify_b), V4356, tmp23317, V4357, V4358)
		return

	}, 3)

	tmp23318 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1track, tmp23313)

	_ = tmp23318

	tmp23319 := MakeNative(func(__e *ControlFlow) {
		V4366 := __e.Get(1)
		_ = V4366
		V4367 := __e.Get(2)
		_ = V4367
		V4368 := __e.Get(3)
		_ = V4368
		tmp23320 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23321 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23321

			tmp23322 := PrimCons(A, Nil)

			tmp23323 := PrimCons(sym_1_1_6, tmp23322)

			tmp23324 := PrimCons(symexception, tmp23323)

			tmp23325 := PrimCons(A, Nil)

			tmp23326 := PrimCons(sym_1_1_6, tmp23325)

			tmp23327 := PrimCons(tmp23324, tmp23326)

			tmp23328 := PrimCons(tmp23327, Nil)

			tmp23329 := PrimCons(sym_1_1_6, tmp23328)

			tmp23330 := PrimCons(A, tmp23329)

			__e.TailApply(PrimNS2Value(symunify_b), V4366, tmp23330, V4367, V4368)
			return

		}, 1)

		tmp23331 := Call(__e, PrimNS2Value(symshen_4newpv), V4367)

		__e.TailApply(tmp23320, tmp23331)
		return

	}, 3)

	tmp23332 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1trap_1error, tmp23319)

	_ = tmp23332

	tmp23333 := MakeNative(func(__e *ControlFlow) {
		V4376 := __e.Get(1)
		_ = V4376
		V4377 := __e.Get(2)
		_ = V4377
		V4378 := __e.Get(3)
		_ = V4378
		tmp23334 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23335 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23335

			tmp23336 := PrimCons(symboolean, Nil)

			tmp23337 := PrimCons(sym_1_1_6, tmp23336)

			tmp23338 := PrimCons(A, tmp23337)

			__e.TailApply(PrimNS2Value(symunify_b), V4376, tmp23338, V4377, V4378)
			return

		}, 1)

		tmp23339 := Call(__e, PrimNS2Value(symshen_4newpv), V4377)

		__e.TailApply(tmp23334, tmp23339)
		return

	}, 3)

	tmp23340 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1tuple_2, tmp23333)

	_ = tmp23340

	tmp23341 := MakeNative(func(__e *ControlFlow) {
		V4386 := __e.Get(1)
		_ = V4386
		V4387 := __e.Get(2)
		_ = V4387
		V4388 := __e.Get(3)
		_ = V4388
		tmp23342 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23342

		tmp23343 := PrimCons(symsymbol, Nil)

		tmp23344 := PrimCons(sym_1_1_6, tmp23343)

		tmp23345 := PrimCons(symsymbol, tmp23344)

		__e.TailApply(PrimNS2Value(symunify_b), V4386, tmp23345, V4387, V4388)
		return

	}, 3)

	tmp23346 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1undefmacro, tmp23341)

	_ = tmp23346

	tmp23347 := MakeNative(func(__e *ControlFlow) {
		V4396 := __e.Get(1)
		_ = V4396
		V4397 := __e.Get(2)
		_ = V4397
		V4398 := __e.Get(3)
		_ = V4398
		tmp23348 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23349 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23349

			tmp23350 := PrimCons(A, Nil)

			tmp23351 := PrimCons(symlist, tmp23350)

			tmp23352 := PrimCons(A, Nil)

			tmp23353 := PrimCons(symlist, tmp23352)

			tmp23354 := PrimCons(A, Nil)

			tmp23355 := PrimCons(symlist, tmp23354)

			tmp23356 := PrimCons(tmp23355, Nil)

			tmp23357 := PrimCons(sym_1_1_6, tmp23356)

			tmp23358 := PrimCons(tmp23353, tmp23357)

			tmp23359 := PrimCons(tmp23358, Nil)

			tmp23360 := PrimCons(sym_1_1_6, tmp23359)

			tmp23361 := PrimCons(tmp23351, tmp23360)

			__e.TailApply(PrimNS2Value(symunify_b), V4396, tmp23361, V4397, V4398)
			return

		}, 1)

		tmp23362 := Call(__e, PrimNS2Value(symshen_4newpv), V4397)

		__e.TailApply(tmp23348, tmp23362)
		return

	}, 3)

	tmp23363 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1union, tmp23347)

	_ = tmp23363

	tmp23364 := MakeNative(func(__e *ControlFlow) {
		V4406 := __e.Get(1)
		_ = V4406
		V4407 := __e.Get(2)
		_ = V4407
		V4408 := __e.Get(3)
		_ = V4408
		tmp23365 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23366 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp23367 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp23367

				tmp23368 := PrimCons(B, Nil)

				tmp23369 := PrimCons(sym_1_1_6, tmp23368)

				tmp23370 := PrimCons(A, tmp23369)

				tmp23371 := PrimCons(B, Nil)

				tmp23372 := PrimCons(sym_1_1_6, tmp23371)

				tmp23373 := PrimCons(A, tmp23372)

				tmp23374 := PrimCons(tmp23373, Nil)

				tmp23375 := PrimCons(sym_1_1_6, tmp23374)

				tmp23376 := PrimCons(tmp23370, tmp23375)

				__e.TailApply(PrimNS2Value(symunify_b), V4406, tmp23376, V4407, V4408)
				return

			}, 1)

			tmp23377 := Call(__e, PrimNS2Value(symshen_4newpv), V4407)

			__e.TailApply(tmp23366, tmp23377)
			return

		}, 1)

		tmp23378 := Call(__e, PrimNS2Value(symshen_4newpv), V4407)

		__e.TailApply(tmp23365, tmp23378)
		return

	}, 3)

	tmp23379 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unprofile, tmp23364)

	_ = tmp23379

	tmp23380 := MakeNative(func(__e *ControlFlow) {
		V4416 := __e.Get(1)
		_ = V4416
		V4417 := __e.Get(2)
		_ = V4417
		V4418 := __e.Get(3)
		_ = V4418
		tmp23381 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23381

		tmp23382 := PrimCons(symsymbol, Nil)

		tmp23383 := PrimCons(sym_1_1_6, tmp23382)

		tmp23384 := PrimCons(symsymbol, tmp23383)

		__e.TailApply(PrimNS2Value(symunify_b), V4416, tmp23384, V4417, V4418)
		return

	}, 3)

	tmp23385 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1untrack, tmp23380)

	_ = tmp23385

	tmp23386 := MakeNative(func(__e *ControlFlow) {
		V4426 := __e.Get(1)
		_ = V4426
		V4427 := __e.Get(2)
		_ = V4427
		V4428 := __e.Get(3)
		_ = V4428
		tmp23387 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23387

		tmp23388 := PrimCons(symsymbol, Nil)

		tmp23389 := PrimCons(sym_1_1_6, tmp23388)

		tmp23390 := PrimCons(symsymbol, tmp23389)

		__e.TailApply(PrimNS2Value(symunify_b), V4426, tmp23390, V4427, V4428)
		return

	}, 3)

	tmp23391 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1unspecialise, tmp23386)

	_ = tmp23391

	tmp23392 := MakeNative(func(__e *ControlFlow) {
		V4436 := __e.Get(1)
		_ = V4436
		V4437 := __e.Get(2)
		_ = V4437
		V4438 := __e.Get(3)
		_ = V4438
		tmp23393 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23394 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23394

			tmp23395 := PrimCons(symboolean, Nil)

			tmp23396 := PrimCons(sym_1_1_6, tmp23395)

			tmp23397 := PrimCons(A, tmp23396)

			__e.TailApply(PrimNS2Value(symunify_b), V4436, tmp23397, V4437, V4438)
			return

		}, 1)

		tmp23398 := Call(__e, PrimNS2Value(symshen_4newpv), V4437)

		__e.TailApply(tmp23393, tmp23398)
		return

	}, 3)

	tmp23399 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1variable_2, tmp23392)

	_ = tmp23399

	tmp23400 := MakeNative(func(__e *ControlFlow) {
		V4446 := __e.Get(1)
		_ = V4446
		V4447 := __e.Get(2)
		_ = V4447
		V4448 := __e.Get(3)
		_ = V4448
		tmp23401 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23402 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23402

			tmp23403 := PrimCons(symboolean, Nil)

			tmp23404 := PrimCons(sym_1_1_6, tmp23403)

			tmp23405 := PrimCons(A, tmp23404)

			__e.TailApply(PrimNS2Value(symunify_b), V4446, tmp23405, V4447, V4448)
			return

		}, 1)

		tmp23406 := Call(__e, PrimNS2Value(symshen_4newpv), V4447)

		__e.TailApply(tmp23401, tmp23406)
		return

	}, 3)

	tmp23407 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1vector_2, tmp23400)

	_ = tmp23407

	tmp23408 := MakeNative(func(__e *ControlFlow) {
		V4456 := __e.Get(1)
		_ = V4456
		V4457 := __e.Get(2)
		_ = V4457
		V4458 := __e.Get(3)
		_ = V4458
		tmp23409 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23409

		tmp23410 := PrimCons(symstring, Nil)

		tmp23411 := PrimCons(sym_1_1_6, tmp23410)

		__e.TailApply(PrimNS2Value(symunify_b), V4456, tmp23411, V4457, V4458)
		return

	}, 3)

	tmp23412 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1version, tmp23408)

	_ = tmp23412

	tmp23413 := MakeNative(func(__e *ControlFlow) {
		V4466 := __e.Get(1)
		_ = V4466
		V4467 := __e.Get(2)
		_ = V4467
		V4468 := __e.Get(3)
		_ = V4468
		tmp23414 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23415 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23415

			tmp23416 := PrimCons(A, Nil)

			tmp23417 := PrimCons(sym_1_1_6, tmp23416)

			tmp23418 := PrimCons(A, tmp23417)

			tmp23419 := PrimCons(tmp23418, Nil)

			tmp23420 := PrimCons(sym_1_1_6, tmp23419)

			tmp23421 := PrimCons(symstring, tmp23420)

			__e.TailApply(PrimNS2Value(symunify_b), V4466, tmp23421, V4467, V4468)
			return

		}, 1)

		tmp23422 := Call(__e, PrimNS2Value(symshen_4newpv), V4467)

		__e.TailApply(tmp23414, tmp23422)
		return

	}, 3)

	tmp23423 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1to_1file, tmp23413)

	_ = tmp23423

	tmp23424 := MakeNative(func(__e *ControlFlow) {
		V4476 := __e.Get(1)
		_ = V4476
		V4477 := __e.Get(2)
		_ = V4477
		V4478 := __e.Get(3)
		_ = V4478
		tmp23425 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23425

		tmp23426 := PrimCons(symout, Nil)

		tmp23427 := PrimCons(symstream, tmp23426)

		tmp23428 := PrimCons(symnumber, Nil)

		tmp23429 := PrimCons(sym_1_1_6, tmp23428)

		tmp23430 := PrimCons(tmp23427, tmp23429)

		tmp23431 := PrimCons(tmp23430, Nil)

		tmp23432 := PrimCons(sym_1_1_6, tmp23431)

		tmp23433 := PrimCons(symnumber, tmp23432)

		__e.TailApply(PrimNS2Value(symunify_b), V4476, tmp23433, V4477, V4478)
		return

	}, 3)

	tmp23434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1write_1byte, tmp23424)

	_ = tmp23434

	tmp23435 := MakeNative(func(__e *ControlFlow) {
		V4486 := __e.Get(1)
		_ = V4486
		V4487 := __e.Get(2)
		_ = V4487
		V4488 := __e.Get(3)
		_ = V4488
		tmp23436 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23436

		tmp23437 := PrimCons(symboolean, Nil)

		tmp23438 := PrimCons(sym_1_1_6, tmp23437)

		tmp23439 := PrimCons(symstring, tmp23438)

		__e.TailApply(PrimNS2Value(symunify_b), V4486, tmp23439, V4487, V4488)
		return

	}, 3)

	tmp23440 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1y_1or_1n_2, tmp23435)

	_ = tmp23440

	tmp23441 := MakeNative(func(__e *ControlFlow) {
		V4496 := __e.Get(1)
		_ = V4496
		V4497 := __e.Get(2)
		_ = V4497
		V4498 := __e.Get(3)
		_ = V4498
		tmp23442 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23442

		tmp23443 := PrimCons(symboolean, Nil)

		tmp23444 := PrimCons(sym_1_1_6, tmp23443)

		tmp23445 := PrimCons(symnumber, tmp23444)

		tmp23446 := PrimCons(tmp23445, Nil)

		tmp23447 := PrimCons(sym_1_1_6, tmp23446)

		tmp23448 := PrimCons(symnumber, tmp23447)

		__e.TailApply(PrimNS2Value(symunify_b), V4496, tmp23448, V4497, V4498)
		return

	}, 3)

	tmp23449 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6, tmp23441)

	_ = tmp23449

	tmp23450 := MakeNative(func(__e *ControlFlow) {
		V4506 := __e.Get(1)
		_ = V4506
		V4507 := __e.Get(2)
		_ = V4507
		V4508 := __e.Get(3)
		_ = V4508
		tmp23451 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23451

		tmp23452 := PrimCons(symboolean, Nil)

		tmp23453 := PrimCons(sym_1_1_6, tmp23452)

		tmp23454 := PrimCons(symnumber, tmp23453)

		tmp23455 := PrimCons(tmp23454, Nil)

		tmp23456 := PrimCons(sym_1_1_6, tmp23455)

		tmp23457 := PrimCons(symnumber, tmp23456)

		__e.TailApply(PrimNS2Value(symunify_b), V4506, tmp23457, V4507, V4508)
		return

	}, 3)

	tmp23458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5, tmp23450)

	_ = tmp23458

	tmp23459 := MakeNative(func(__e *ControlFlow) {
		V4516 := __e.Get(1)
		_ = V4516
		V4517 := __e.Get(2)
		_ = V4517
		V4518 := __e.Get(3)
		_ = V4518
		tmp23460 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23460

		tmp23461 := PrimCons(symboolean, Nil)

		tmp23462 := PrimCons(sym_1_1_6, tmp23461)

		tmp23463 := PrimCons(symnumber, tmp23462)

		tmp23464 := PrimCons(tmp23463, Nil)

		tmp23465 := PrimCons(sym_1_1_6, tmp23464)

		tmp23466 := PrimCons(symnumber, tmp23465)

		__e.TailApply(PrimNS2Value(symunify_b), V4516, tmp23466, V4517, V4518)
		return

	}, 3)

	tmp23467 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_6_a, tmp23459)

	_ = tmp23467

	tmp23468 := MakeNative(func(__e *ControlFlow) {
		V4526 := __e.Get(1)
		_ = V4526
		V4527 := __e.Get(2)
		_ = V4527
		V4528 := __e.Get(3)
		_ = V4528
		tmp23469 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23469

		tmp23470 := PrimCons(symboolean, Nil)

		tmp23471 := PrimCons(sym_1_1_6, tmp23470)

		tmp23472 := PrimCons(symnumber, tmp23471)

		tmp23473 := PrimCons(tmp23472, Nil)

		tmp23474 := PrimCons(sym_1_1_6, tmp23473)

		tmp23475 := PrimCons(symnumber, tmp23474)

		__e.TailApply(PrimNS2Value(symunify_b), V4526, tmp23475, V4527, V4528)
		return

	}, 3)

	tmp23476 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_5_a, tmp23468)

	_ = tmp23476

	tmp23477 := MakeNative(func(__e *ControlFlow) {
		V4536 := __e.Get(1)
		_ = V4536
		V4537 := __e.Get(2)
		_ = V4537
		V4538 := __e.Get(3)
		_ = V4538
		tmp23478 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23479 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp23479

			tmp23480 := PrimCons(symboolean, Nil)

			tmp23481 := PrimCons(sym_1_1_6, tmp23480)

			tmp23482 := PrimCons(A, tmp23481)

			tmp23483 := PrimCons(tmp23482, Nil)

			tmp23484 := PrimCons(sym_1_1_6, tmp23483)

			tmp23485 := PrimCons(A, tmp23484)

			__e.TailApply(PrimNS2Value(symunify_b), V4536, tmp23485, V4537, V4538)
			return

		}, 1)

		tmp23486 := Call(__e, PrimNS2Value(symshen_4newpv), V4537)

		__e.TailApply(tmp23478, tmp23486)
		return

	}, 3)

	tmp23487 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a, tmp23477)

	_ = tmp23487

	tmp23488 := MakeNative(func(__e *ControlFlow) {
		V4546 := __e.Get(1)
		_ = V4546
		V4547 := __e.Get(2)
		_ = V4547
		V4548 := __e.Get(3)
		_ = V4548
		tmp23489 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23489

		tmp23490 := PrimCons(symnumber, Nil)

		tmp23491 := PrimCons(sym_1_1_6, tmp23490)

		tmp23492 := PrimCons(symnumber, tmp23491)

		tmp23493 := PrimCons(tmp23492, Nil)

		tmp23494 := PrimCons(sym_1_1_6, tmp23493)

		tmp23495 := PrimCons(symnumber, tmp23494)

		__e.TailApply(PrimNS2Value(symunify_b), V4546, tmp23495, V4547, V4548)
		return

	}, 3)

	tmp23496 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_7, tmp23488)

	_ = tmp23496

	tmp23497 := MakeNative(func(__e *ControlFlow) {
		V4556 := __e.Get(1)
		_ = V4556
		V4557 := __e.Get(2)
		_ = V4557
		V4558 := __e.Get(3)
		_ = V4558
		tmp23498 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23498

		tmp23499 := PrimCons(symnumber, Nil)

		tmp23500 := PrimCons(sym_1_1_6, tmp23499)

		tmp23501 := PrimCons(symnumber, tmp23500)

		tmp23502 := PrimCons(tmp23501, Nil)

		tmp23503 := PrimCons(sym_1_1_6, tmp23502)

		tmp23504 := PrimCons(symnumber, tmp23503)

		__e.TailApply(PrimNS2Value(symunify_b), V4556, tmp23504, V4557, V4558)
		return

	}, 3)

	tmp23505 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_c, tmp23497)

	_ = tmp23505

	tmp23506 := MakeNative(func(__e *ControlFlow) {
		V4566 := __e.Get(1)
		_ = V4566
		V4567 := __e.Get(2)
		_ = V4567
		V4568 := __e.Get(3)
		_ = V4568
		tmp23507 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23507

		tmp23508 := PrimCons(symnumber, Nil)

		tmp23509 := PrimCons(sym_1_1_6, tmp23508)

		tmp23510 := PrimCons(symnumber, tmp23509)

		tmp23511 := PrimCons(tmp23510, Nil)

		tmp23512 := PrimCons(sym_1_1_6, tmp23511)

		tmp23513 := PrimCons(symnumber, tmp23512)

		__e.TailApply(PrimNS2Value(symunify_b), V4566, tmp23513, V4567, V4568)
		return

	}, 3)

	tmp23514 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_1, tmp23506)

	_ = tmp23514

	tmp23515 := MakeNative(func(__e *ControlFlow) {
		V4576 := __e.Get(1)
		_ = V4576
		V4577 := __e.Get(2)
		_ = V4577
		V4578 := __e.Get(3)
		_ = V4578
		tmp23516 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp23516

		tmp23517 := PrimCons(symnumber, Nil)

		tmp23518 := PrimCons(sym_1_1_6, tmp23517)

		tmp23519 := PrimCons(symnumber, tmp23518)

		tmp23520 := PrimCons(tmp23519, Nil)

		tmp23521 := PrimCons(sym_1_1_6, tmp23520)

		tmp23522 := PrimCons(symnumber, tmp23521)

		__e.TailApply(PrimNS2Value(symunify_b), V4576, tmp23522, V4577, V4578)
		return

	}, 3)

	tmp23523 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_d, tmp23515)

	_ = tmp23523

	tmp23524 := MakeNative(func(__e *ControlFlow) {
		V4586 := __e.Get(1)
		_ = V4586
		V4587 := __e.Get(2)
		_ = V4587
		V4588 := __e.Get(3)
		_ = V4588
		tmp23525 := MakeNative(func(__e *ControlFlow) {
			A := __e.Get(1)
			_ = A
			tmp23526 := MakeNative(func(__e *ControlFlow) {
				B := __e.Get(1)
				_ = B
				tmp23527 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp23527

				tmp23528 := PrimCons(symboolean, Nil)

				tmp23529 := PrimCons(sym_1_1_6, tmp23528)

				tmp23530 := PrimCons(B, tmp23529)

				tmp23531 := PrimCons(tmp23530, Nil)

				tmp23532 := PrimCons(sym_1_1_6, tmp23531)

				tmp23533 := PrimCons(A, tmp23532)

				__e.TailApply(PrimNS2Value(symunify_b), V4586, tmp23533, V4587, V4588)
				return

			}, 1)

			tmp23534 := Call(__e, PrimNS2Value(symshen_4newpv), V4587)

			__e.TailApply(tmp23526, tmp23534)
			return

		}, 1)

		tmp23535 := Call(__e, PrimNS2Value(symshen_4newpv), V4587)

		__e.TailApply(tmp23525, tmp23535)
		return

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4type_1signature_1of_1_a_a, tmp23524)
	return

}, 0)
