package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen9272 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3210 := __args[0]
			_ = V3210
			V3211 := __args[1]
			_ = V3211
			gen9241 := Call(__e, ShenFunc(symcons), V3210, V3211)

			gen9242 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen9243 := Call(__e, ShenFunc(symcons), gen9241, gen9242)

			gen9244 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen9243)

			Record := gen9244
			_ = Record
			gen9245 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.Return(MakeSymbol("shen.skip"))
				return
			}, 1)
			gen9246 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4variancy_1test), V3210, V3211)

				return
			}, 0)
			gen9247 := Try(__e, gen9246).Catch(gen9245)
			Variancy := gen9247
			_ = Variancy
			gen9248 := Call(__e, ShenFunc(symshen_4demodulate), V3211)

			gen9249 := Call(__e, ShenFunc(symshen_4rcons__form), gen9248)

			Type := gen9249
			gen9250 := Call(__e, ShenFunc(symconcat), MakeSymbol("shen.type-signature-of-"), V3210)

			F_d := gen9250
			gen9251 := Call(__e, ShenFunc(symshen_4parameters), MakeNumber(1))

			Parameters := gen9251
			gen9252 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

			gen9253 := Call(__e, ShenFunc(symcons), F_d, gen9252)

			gen9254 := Call(__e, ShenFunc(symcons), Type, Nil)

			gen9255 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen9254)

			gen9256 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen9255)

			gen9257 := Call(__e, ShenFunc(symcons), gen9256, Nil)

			gen9258 := Call(__e, ShenFunc(symcons), gen9257, Nil)

			gen9259 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen9258)

			gen9260 := Call(__e, ShenFunc(symcons), gen9253, gen9259)

			Clause := gen9260
			gen9261 := Call(__e, ShenFunc(symshen_4aum), Clause, Parameters)

			AUM__instruction := gen9261
			gen9262 := Call(__e, ShenFunc(symshen_4aum__to__shen), AUM__instruction)

			Code := gen9262
			gen9263 := Call(__e, ShenFunc(symcons), MakeSymbol("Continuation"), Nil)

			gen9264 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), gen9263)

			gen9265 := Call(__e, ShenFunc(symcons), Code, Nil)

			gen9266 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen9265)

			gen9267 := Call(__e, ShenFunc(symappend), gen9264, gen9266)

			gen9268 := Call(__e, ShenFunc(symappend), Parameters, gen9267)

			gen9269 := Call(__e, ShenFunc(symcons), F_d, gen9268)

			gen9270 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen9269)

			ShenDef := gen9270
			gen9271 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), ShenDef)

			Eval := gen9271
			_ = Eval
			__e.Return(V3210)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("declare"), gen9272)

		gen9276 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3213 := __args[0]
			_ = V3213
			gen9273 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*demodulation-function*"))

			gen9274 := Call(__e, ShenFunc(symshen_4walk), gen9273, V3213)

			Demod := gen9274
			gen9275 := Call(__e, ShenFunc(sym_a), Demod, V3213)

			if True == gen9275 {
				__e.Return(V3213)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4demodulate), Demod)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.demodulate"), gen9276)

		gen9284 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3216 := __args[0]
			_ = V3216
			V3217 := __args[1]
			_ = V3217
			gen9277 := Call(__e, ShenFunc(symshen_4typecheck), V3216, MakeSymbol("B"))

			TypeF := gen9277
			gen9282 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), TypeF)

			var gen9283 Obj
			if True == gen9282 {
				gen9283 = MakeSymbol("shen.skip")
			} else {
				gen9281 := Call(__e, ShenFunc(symshen_4variant_2), TypeF, V3217)

				if True == gen9281 {
					gen9283 = MakeSymbol("shen.skip")
				} else {
					gen9278 := Call(__e, ShenFunc(symshen_4app), V3216, MakeString(" may create errors\n"), MakeSymbol("shen.a"))

					gen9279 := Call(__e, ShenFunc(symcn), MakeString("warning: changing the type of "), gen9278)

					gen9280 := Call(__e, ShenFunc(symstoutput))

					gen9283 = Call(__e, ShenFunc(symshen_4prhush), gen9279, gen9280)

				}

			}
			Check := gen9283
			_ = Check
			__e.Return(MakeSymbol("shen.skip"))
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.variancy-test"), gen9284)

		gen9325 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3230 := __args[0]
			_ = V3230
			V3231 := __args[1]
			_ = V3231
			gen9324 := Call(__e, ShenFunc(sym_a), V3231, V3230)

			if True == gen9324 {
				__e.Return(True)
				return
			} else {
				gen9322 := Call(__e, ShenFunc(symcons_2), V3230)

				var gen9323 Obj
				if True == gen9322 {
					gen9320 := Call(__e, ShenFunc(symcons_2), V3231)

					var gen9321 Obj
					if True == gen9320 {
						gen9317 := Call(__e, ShenFunc(symhd), V3231)

						gen9318 := Call(__e, ShenFunc(symhd), V3230)

						gen9319 := Call(__e, ShenFunc(sym_a), gen9317, gen9318)

						if True == gen9319 {
							gen9321 = True
						} else {
							gen9321 = False
						}

					} else {
						gen9321 = False
					}
					if True == gen9321 {
						gen9323 = True
					} else {
						gen9323 = False
					}

				} else {
					gen9323 = False
				}
				if True == gen9323 {
					gen9315 := Call(__e, ShenFunc(symtl), V3230)

					gen9316 := Call(__e, ShenFunc(symtl), V3231)

					__e.TailApply(ShenFunc(symshen_4variant_2), gen9315, gen9316)

					return

				} else {
					gen9313 := Call(__e, ShenFunc(symcons_2), V3230)

					var gen9314 Obj
					if True == gen9313 {
						gen9311 := Call(__e, ShenFunc(symcons_2), V3231)

						var gen9312 Obj
						if True == gen9311 {
							gen9308 := Call(__e, ShenFunc(symhd), V3230)

							gen9309 := Call(__e, ShenFunc(symshen_4pvar_2), gen9308)

							var gen9310 Obj
							if True == gen9309 {
								gen9306 := Call(__e, ShenFunc(symhd), V3231)

								gen9307 := Call(__e, ShenFunc(symvariable_2), gen9306)

								if True == gen9307 {
									gen9310 = True
								} else {
									gen9310 = False
								}

							} else {
								gen9310 = False
							}
							if True == gen9310 {
								gen9312 = True
							} else {
								gen9312 = False
							}

						} else {
							gen9312 = False
						}
						if True == gen9312 {
							gen9314 = True
						} else {
							gen9314 = False
						}

					} else {
						gen9314 = False
					}
					if True == gen9314 {
						gen9300 := Call(__e, ShenFunc(symhd), V3230)

						gen9301 := Call(__e, ShenFunc(symtl), V3230)

						gen9302 := Call(__e, ShenFunc(symsubst), MakeSymbol("shen.a"), gen9300, gen9301)

						gen9303 := Call(__e, ShenFunc(symhd), V3231)

						gen9304 := Call(__e, ShenFunc(symtl), V3231)

						gen9305 := Call(__e, ShenFunc(symsubst), MakeSymbol("shen.a"), gen9303, gen9304)

						__e.TailApply(ShenFunc(symshen_4variant_2), gen9302, gen9305)

						return

					} else {
						gen9298 := Call(__e, ShenFunc(symcons_2), V3230)

						var gen9299 Obj
						if True == gen9298 {
							gen9295 := Call(__e, ShenFunc(symhd), V3230)

							gen9296 := Call(__e, ShenFunc(symcons_2), gen9295)

							var gen9297 Obj
							if True == gen9296 {
								gen9293 := Call(__e, ShenFunc(symcons_2), V3231)

								var gen9294 Obj
								if True == gen9293 {
									gen9291 := Call(__e, ShenFunc(symhd), V3231)

									gen9292 := Call(__e, ShenFunc(symcons_2), gen9291)

									if True == gen9292 {
										gen9294 = True
									} else {
										gen9294 = False
									}

								} else {
									gen9294 = False
								}
								if True == gen9294 {
									gen9297 = True
								} else {
									gen9297 = False
								}

							} else {
								gen9297 = False
							}
							if True == gen9297 {
								gen9299 = True
							} else {
								gen9299 = False
							}

						} else {
							gen9299 = False
						}
						if True == gen9299 {
							gen9285 := Call(__e, ShenFunc(symhd), V3230)

							gen9286 := Call(__e, ShenFunc(symtl), V3230)

							gen9287 := Call(__e, ShenFunc(symappend), gen9285, gen9286)

							gen9288 := Call(__e, ShenFunc(symhd), V3231)

							gen9289 := Call(__e, ShenFunc(symtl), V3231)

							gen9290 := Call(__e, ShenFunc(symappend), gen9288, gen9289)

							__e.TailApply(ShenFunc(symshen_4variant_2), gen9287, gen9290)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.variant?"), gen9325)

		gen9330 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3236 := __args[0]
			_ = V3236
			V3237 := __args[1]
			_ = V3237
			V3238 := __args[2]
			_ = V3238
			gen9326 := Call(__e, ShenFunc(symshen_4newpv), V3237)

			A := gen9326
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9327 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9328 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9327)

			gen9329 := Call(__e, ShenFunc(symcons), A, gen9328)

			__e.TailApply(ShenFunc(symunify_b), V3236, gen9329, V3237, V3238)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-absvector?"), gen9330)

		gen9342 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3246 := __args[0]
			_ = V3246
			V3247 := __args[1]
			_ = V3247
			V3248 := __args[2]
			_ = V3248
			gen9331 := Call(__e, ShenFunc(symshen_4newpv), V3247)

			A := gen9331
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9332 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9333 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9332)

			gen9334 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9335 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9334)

			gen9336 := Call(__e, ShenFunc(symcons), gen9335, Nil)

			gen9337 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9336)

			gen9338 := Call(__e, ShenFunc(symcons), gen9333, gen9337)

			gen9339 := Call(__e, ShenFunc(symcons), gen9338, Nil)

			gen9340 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9339)

			gen9341 := Call(__e, ShenFunc(symcons), A, gen9340)

			__e.TailApply(ShenFunc(symunify_b), V3246, gen9341, V3247, V3248)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-adjoin"), gen9342)

		gen9349 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3256 := __args[0]
			_ = V3256
			V3257 := __args[1]
			_ = V3257
			V3258 := __args[2]
			_ = V3258
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9343 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9344 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9343)

			gen9345 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9344)

			gen9346 := Call(__e, ShenFunc(symcons), gen9345, Nil)

			gen9347 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9346)

			gen9348 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9347)

			__e.TailApply(ShenFunc(symunify_b), V3256, gen9348, V3257, V3258)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-and"), gen9349)

		gen9360 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3266 := __args[0]
			_ = V3266
			V3267 := __args[1]
			_ = V3267
			V3268 := __args[2]
			_ = V3268
			gen9350 := Call(__e, ShenFunc(symshen_4newpv), V3267)

			A := gen9350
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9351 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9352 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9351)

			gen9353 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9352)

			gen9354 := Call(__e, ShenFunc(symcons), gen9353, Nil)

			gen9355 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9354)

			gen9356 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9355)

			gen9357 := Call(__e, ShenFunc(symcons), gen9356, Nil)

			gen9358 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9357)

			gen9359 := Call(__e, ShenFunc(symcons), A, gen9358)

			__e.TailApply(ShenFunc(symunify_b), V3266, gen9359, V3267, V3268)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-shen.app"), gen9360)

		gen9374 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3276 := __args[0]
			_ = V3276
			V3277 := __args[1]
			_ = V3277
			V3278 := __args[2]
			_ = V3278
			gen9361 := Call(__e, ShenFunc(symshen_4newpv), V3277)

			A := gen9361
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9362 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9363 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9362)

			gen9364 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9365 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9364)

			gen9366 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9367 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9366)

			gen9368 := Call(__e, ShenFunc(symcons), gen9367, Nil)

			gen9369 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9368)

			gen9370 := Call(__e, ShenFunc(symcons), gen9365, gen9369)

			gen9371 := Call(__e, ShenFunc(symcons), gen9370, Nil)

			gen9372 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9371)

			gen9373 := Call(__e, ShenFunc(symcons), gen9363, gen9372)

			__e.TailApply(ShenFunc(symunify_b), V3276, gen9373, V3277, V3278)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-append"), gen9374)

		gen9379 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3286 := __args[0]
			_ = V3286
			V3287 := __args[1]
			_ = V3287
			V3288 := __args[2]
			_ = V3288
			gen9375 := Call(__e, ShenFunc(symshen_4newpv), V3287)

			A := gen9375
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9376 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9377 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9376)

			gen9378 := Call(__e, ShenFunc(symcons), A, gen9377)

			__e.TailApply(ShenFunc(symunify_b), V3286, gen9378, V3287, V3288)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-arity"), gen9379)

		gen9393 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3296 := __args[0]
			_ = V3296
			V3297 := __args[1]
			_ = V3297
			V3298 := __args[2]
			_ = V3298
			gen9380 := Call(__e, ShenFunc(symshen_4newpv), V3297)

			A := gen9380
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9381 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9382 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9381)

			gen9383 := Call(__e, ShenFunc(symcons), gen9382, Nil)

			gen9384 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9383)

			gen9385 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9386 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9385)

			gen9387 := Call(__e, ShenFunc(symcons), gen9386, Nil)

			gen9388 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9387)

			gen9389 := Call(__e, ShenFunc(symcons), gen9384, gen9388)

			gen9390 := Call(__e, ShenFunc(symcons), gen9389, Nil)

			gen9391 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9390)

			gen9392 := Call(__e, ShenFunc(symcons), A, gen9391)

			__e.TailApply(ShenFunc(symunify_b), V3296, gen9392, V3297, V3298)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-assoc"), gen9393)

		gen9398 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3306 := __args[0]
			_ = V3306
			V3307 := __args[1]
			_ = V3307
			V3308 := __args[2]
			_ = V3308
			gen9394 := Call(__e, ShenFunc(symshen_4newpv), V3307)

			A := gen9394
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9395 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9396 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9395)

			gen9397 := Call(__e, ShenFunc(symcons), A, gen9396)

			__e.TailApply(ShenFunc(symunify_b), V3306, gen9397, V3307, V3308)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-boolean?"), gen9398)

		gen9402 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3316 := __args[0]
			_ = V3316
			V3317 := __args[1]
			_ = V3317
			V3318 := __args[2]
			_ = V3318
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9399 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9400 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9399)

			gen9401 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9400)

			__e.TailApply(ShenFunc(symunify_b), V3316, gen9401, V3317, V3318)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-bound?"), gen9402)

		gen9406 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3326 := __args[0]
			_ = V3326
			V3327 := __args[1]
			_ = V3327
			V3328 := __args[2]
			_ = V3328
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9403 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9404 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9403)

			gen9405 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9404)

			__e.TailApply(ShenFunc(symunify_b), V3326, gen9405, V3327, V3328)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-cd"), gen9406)

		gen9416 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3336 := __args[0]
			_ = V3336
			V3337 := __args[1]
			_ = V3337
			V3338 := __args[2]
			_ = V3338
			gen9407 := Call(__e, ShenFunc(symshen_4newpv), V3337)

			A := gen9407
			gen9408 := Call(__e, ShenFunc(symshen_4newpv), V3337)

			B := gen9408
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9409 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9410 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen9409)

			gen9411 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9412 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9411)

			gen9413 := Call(__e, ShenFunc(symcons), gen9412, Nil)

			gen9414 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9413)

			gen9415 := Call(__e, ShenFunc(symcons), gen9410, gen9414)

			__e.TailApply(ShenFunc(symunify_b), V3336, gen9415, V3337, V3338)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-close"), gen9416)

		gen9423 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3346 := __args[0]
			_ = V3346
			V3347 := __args[1]
			_ = V3347
			V3348 := __args[2]
			_ = V3348
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9417 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9418 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9417)

			gen9419 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9418)

			gen9420 := Call(__e, ShenFunc(symcons), gen9419, Nil)

			gen9421 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9420)

			gen9422 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9421)

			__e.TailApply(ShenFunc(symunify_b), V3346, gen9422, V3347, V3348)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-cn"), gen9423)

		gen9441 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3356 := __args[0]
			_ = V3356
			V3357 := __args[1]
			_ = V3357
			V3358 := __args[2]
			_ = V3358
			gen9424 := Call(__e, ShenFunc(symshen_4newpv), V3357)

			A := gen9424
			gen9425 := Call(__e, ShenFunc(symshen_4newpv), V3357)

			B := gen9425
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9426 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9427 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen9426)

			gen9428 := Call(__e, ShenFunc(symcons), A, gen9427)

			gen9429 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9430 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9429)

			gen9431 := Call(__e, ShenFunc(symcons), A, gen9430)

			gen9432 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9433 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9432)

			gen9434 := Call(__e, ShenFunc(symcons), gen9431, gen9433)

			gen9435 := Call(__e, ShenFunc(symcons), gen9434, Nil)

			gen9436 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9435)

			gen9437 := Call(__e, ShenFunc(symcons), A, gen9436)

			gen9438 := Call(__e, ShenFunc(symcons), gen9437, Nil)

			gen9439 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9438)

			gen9440 := Call(__e, ShenFunc(symcons), gen9428, gen9439)

			__e.TailApply(ShenFunc(symunify_b), V3356, gen9440, V3357, V3358)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-compile"), gen9441)

		gen9446 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3366 := __args[0]
			_ = V3366
			V3367 := __args[1]
			_ = V3367
			V3368 := __args[2]
			_ = V3368
			gen9442 := Call(__e, ShenFunc(symshen_4newpv), V3367)

			A := gen9442
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9443 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9444 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9443)

			gen9445 := Call(__e, ShenFunc(symcons), A, gen9444)

			__e.TailApply(ShenFunc(symunify_b), V3366, gen9445, V3367, V3368)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-cons?"), gen9446)

		gen9455 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3376 := __args[0]
			_ = V3376
			V3377 := __args[1]
			_ = V3377
			V3378 := __args[2]
			_ = V3378
			gen9447 := Call(__e, ShenFunc(symshen_4newpv), V3377)

			A := gen9447
			gen9448 := Call(__e, ShenFunc(symshen_4newpv), V3377)

			B := gen9448
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9449 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9450 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9449)

			gen9451 := Call(__e, ShenFunc(symcons), A, gen9450)

			gen9452 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9453 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9452)

			gen9454 := Call(__e, ShenFunc(symcons), gen9451, gen9453)

			__e.TailApply(ShenFunc(symunify_b), V3376, gen9454, V3377, V3378)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-destroy"), gen9455)

		gen9469 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3386 := __args[0]
			_ = V3386
			V3387 := __args[1]
			_ = V3387
			V3388 := __args[2]
			_ = V3388
			gen9456 := Call(__e, ShenFunc(symshen_4newpv), V3387)

			A := gen9456
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9457 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9458 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9457)

			gen9459 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9460 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9459)

			gen9461 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9462 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9461)

			gen9463 := Call(__e, ShenFunc(symcons), gen9462, Nil)

			gen9464 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9463)

			gen9465 := Call(__e, ShenFunc(symcons), gen9460, gen9464)

			gen9466 := Call(__e, ShenFunc(symcons), gen9465, Nil)

			gen9467 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9466)

			gen9468 := Call(__e, ShenFunc(symcons), gen9458, gen9467)

			__e.TailApply(ShenFunc(symunify_b), V3386, gen9468, V3387, V3388)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-difference"), gen9469)

		gen9478 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3396 := __args[0]
			_ = V3396
			V3397 := __args[1]
			_ = V3397
			V3398 := __args[2]
			_ = V3398
			gen9470 := Call(__e, ShenFunc(symshen_4newpv), V3397)

			A := gen9470
			gen9471 := Call(__e, ShenFunc(symshen_4newpv), V3397)

			B := gen9471
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9472 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9473 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9472)

			gen9474 := Call(__e, ShenFunc(symcons), B, gen9473)

			gen9475 := Call(__e, ShenFunc(symcons), gen9474, Nil)

			gen9476 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9475)

			gen9477 := Call(__e, ShenFunc(symcons), A, gen9476)

			__e.TailApply(ShenFunc(symunify_b), V3396, gen9477, V3397, V3398)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-do"), gen9478)

		gen9488 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3406 := __args[0]
			_ = V3406
			V3407 := __args[1]
			_ = V3407
			V3408 := __args[2]
			_ = V3408
			gen9479 := Call(__e, ShenFunc(symshen_4newpv), V3407)

			A := gen9479
			gen9480 := Call(__e, ShenFunc(symshen_4newpv), V3407)

			B := gen9480
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9481 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9482 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9481)

			gen9483 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9484 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9483)

			gen9485 := Call(__e, ShenFunc(symcons), gen9484, Nil)

			gen9486 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen9485)

			gen9487 := Call(__e, ShenFunc(symcons), gen9482, gen9486)

			__e.TailApply(ShenFunc(symunify_b), V3406, gen9487, V3407, V3408)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-<e>"), gen9488)

		gen9497 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3416 := __args[0]
			_ = V3416
			V3417 := __args[1]
			_ = V3417
			V3418 := __args[2]
			_ = V3418
			gen9489 := Call(__e, ShenFunc(symshen_4newpv), V3417)

			A := gen9489
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9490 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9491 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9490)

			gen9492 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9493 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9492)

			gen9494 := Call(__e, ShenFunc(symcons), gen9493, Nil)

			gen9495 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen9494)

			gen9496 := Call(__e, ShenFunc(symcons), gen9491, gen9495)

			__e.TailApply(ShenFunc(symunify_b), V3416, gen9496, V3417, V3418)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-<!>"), gen9497)

		gen9507 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3426 := __args[0]
			_ = V3426
			V3427 := __args[1]
			_ = V3427
			V3428 := __args[2]
			_ = V3428
			gen9498 := Call(__e, ShenFunc(symshen_4newpv), V3427)

			A := gen9498
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9499 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9500 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9499)

			gen9501 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9502 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9501)

			gen9503 := Call(__e, ShenFunc(symcons), gen9500, gen9502)

			gen9504 := Call(__e, ShenFunc(symcons), gen9503, Nil)

			gen9505 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9504)

			gen9506 := Call(__e, ShenFunc(symcons), A, gen9505)

			__e.TailApply(ShenFunc(symunify_b), V3426, gen9506, V3427, V3428)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-element?"), gen9507)

		gen9512 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3436 := __args[0]
			_ = V3436
			V3437 := __args[1]
			_ = V3437
			V3438 := __args[2]
			_ = V3438
			gen9508 := Call(__e, ShenFunc(symshen_4newpv), V3437)

			A := gen9508
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9509 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9510 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9509)

			gen9511 := Call(__e, ShenFunc(symcons), A, gen9510)

			__e.TailApply(ShenFunc(symunify_b), V3436, gen9511, V3437, V3438)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-empty?"), gen9512)

		gen9516 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3446 := __args[0]
			_ = V3446
			V3447 := __args[1]
			_ = V3447
			V3448 := __args[2]
			_ = V3448
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9513 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9514 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9513)

			gen9515 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9514)

			__e.TailApply(ShenFunc(symunify_b), V3446, gen9515, V3447, V3448)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-enable-type-theory"), gen9516)

		gen9522 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3456 := __args[0]
			_ = V3456
			V3457 := __args[1]
			_ = V3457
			V3458 := __args[2]
			_ = V3458
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9517 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9518 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9517)

			gen9519 := Call(__e, ShenFunc(symcons), gen9518, Nil)

			gen9520 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9519)

			gen9521 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9520)

			__e.TailApply(ShenFunc(symunify_b), V3456, gen9521, V3457, V3458)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-external"), gen9522)

		gen9526 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3466 := __args[0]
			_ = V3466
			V3467 := __args[1]
			_ = V3467
			V3468 := __args[2]
			_ = V3468
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9523 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9524 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9523)

			gen9525 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen9524)

			__e.TailApply(ShenFunc(symunify_b), V3466, gen9525, V3467, V3468)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-error-to-string"), gen9526)

		gen9533 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3476 := __args[0]
			_ = V3476
			V3477 := __args[1]
			_ = V3477
			V3478 := __args[2]
			_ = V3478
			gen9527 := Call(__e, ShenFunc(symshen_4newpv), V3477)

			A := gen9527
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9528 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9529 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9528)

			gen9530 := Call(__e, ShenFunc(symcons), gen9529, Nil)

			gen9531 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9530)

			gen9532 := Call(__e, ShenFunc(symcons), A, gen9531)

			__e.TailApply(ShenFunc(symunify_b), V3476, gen9532, V3477, V3478)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-explode"), gen9533)

		gen9536 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3486 := __args[0]
			_ = V3486
			V3487 := __args[1]
			_ = V3487
			V3488 := __args[2]
			_ = V3488
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9534 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9535 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9534)

			__e.TailApply(ShenFunc(symunify_b), V3486, gen9535, V3487, V3488)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-fail"), gen9536)

		gen9546 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3496 := __args[0]
			_ = V3496
			V3497 := __args[1]
			_ = V3497
			V3498 := __args[2]
			_ = V3498
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9537 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9538 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9537)

			gen9539 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9538)

			gen9540 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9541 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9540)

			gen9542 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9541)

			gen9543 := Call(__e, ShenFunc(symcons), gen9542, Nil)

			gen9544 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9543)

			gen9545 := Call(__e, ShenFunc(symcons), gen9539, gen9544)

			__e.TailApply(ShenFunc(symunify_b), V3496, gen9545, V3497, V3498)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-fail-if"), gen9546)

		gen9557 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3506 := __args[0]
			_ = V3506
			V3507 := __args[1]
			_ = V3507
			V3508 := __args[2]
			_ = V3508
			gen9547 := Call(__e, ShenFunc(symshen_4newpv), V3507)

			A := gen9547
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9548 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9549 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9548)

			gen9550 := Call(__e, ShenFunc(symcons), A, gen9549)

			gen9551 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9552 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9551)

			gen9553 := Call(__e, ShenFunc(symcons), A, gen9552)

			gen9554 := Call(__e, ShenFunc(symcons), gen9553, Nil)

			gen9555 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9554)

			gen9556 := Call(__e, ShenFunc(symcons), gen9550, gen9555)

			__e.TailApply(ShenFunc(symunify_b), V3506, gen9556, V3507, V3508)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-fix"), gen9557)

		gen9564 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3516 := __args[0]
			_ = V3516
			V3517 := __args[1]
			_ = V3517
			V3518 := __args[2]
			_ = V3518
			gen9558 := Call(__e, ShenFunc(symshen_4newpv), V3517)

			A := gen9558
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9559 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9560 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen9559)

			gen9561 := Call(__e, ShenFunc(symcons), gen9560, Nil)

			gen9562 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9561)

			gen9563 := Call(__e, ShenFunc(symcons), A, gen9562)

			__e.TailApply(ShenFunc(symunify_b), V3516, gen9563, V3517, V3518)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-freeze"), gen9564)

		gen9573 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3526 := __args[0]
			_ = V3526
			V3527 := __args[1]
			_ = V3527
			V3528 := __args[2]
			_ = V3528
			gen9565 := Call(__e, ShenFunc(symshen_4newpv), V3527)

			B := gen9565
			gen9566 := Call(__e, ShenFunc(symshen_4newpv), V3527)

			A := gen9566
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9567 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9568 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen9567)

			gen9569 := Call(__e, ShenFunc(symcons), A, gen9568)

			gen9570 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9571 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9570)

			gen9572 := Call(__e, ShenFunc(symcons), gen9569, gen9571)

			__e.TailApply(ShenFunc(symunify_b), V3526, gen9572, V3527, V3528)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-fst"), gen9573)

		gen9585 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3536 := __args[0]
			_ = V3536
			V3537 := __args[1]
			_ = V3537
			V3538 := __args[2]
			_ = V3538
			gen9574 := Call(__e, ShenFunc(symshen_4newpv), V3537)

			A := gen9574
			gen9575 := Call(__e, ShenFunc(symshen_4newpv), V3537)

			B := gen9575
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9576 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9577 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9576)

			gen9578 := Call(__e, ShenFunc(symcons), A, gen9577)

			gen9579 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9580 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9579)

			gen9581 := Call(__e, ShenFunc(symcons), A, gen9580)

			gen9582 := Call(__e, ShenFunc(symcons), gen9581, Nil)

			gen9583 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9582)

			gen9584 := Call(__e, ShenFunc(symcons), gen9578, gen9583)

			__e.TailApply(ShenFunc(symunify_b), V3536, gen9584, V3537, V3538)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-function"), gen9585)

		gen9589 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3546 := __args[0]
			_ = V3546
			V3547 := __args[1]
			_ = V3547
			V3548 := __args[2]
			_ = V3548
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9586 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9587 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9586)

			gen9588 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9587)

			__e.TailApply(ShenFunc(symunify_b), V3546, gen9588, V3547, V3548)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-gensym"), gen9589)

		gen9599 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3556 := __args[0]
			_ = V3556
			V3557 := __args[1]
			_ = V3557
			V3558 := __args[2]
			_ = V3558
			gen9590 := Call(__e, ShenFunc(symshen_4newpv), V3557)

			A := gen9590
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9591 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9592 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9591)

			gen9593 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9594 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9593)

			gen9595 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9594)

			gen9596 := Call(__e, ShenFunc(symcons), gen9595, Nil)

			gen9597 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9596)

			gen9598 := Call(__e, ShenFunc(symcons), gen9592, gen9597)

			__e.TailApply(ShenFunc(symunify_b), V3556, gen9598, V3557, V3558)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-<-vector"), gen9599)

		gen9614 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3566 := __args[0]
			_ = V3566
			V3567 := __args[1]
			_ = V3567
			V3568 := __args[2]
			_ = V3568
			gen9600 := Call(__e, ShenFunc(symshen_4newpv), V3567)

			A := gen9600
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9601 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9602 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9601)

			gen9603 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9604 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9603)

			gen9605 := Call(__e, ShenFunc(symcons), gen9604, Nil)

			gen9606 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9605)

			gen9607 := Call(__e, ShenFunc(symcons), A, gen9606)

			gen9608 := Call(__e, ShenFunc(symcons), gen9607, Nil)

			gen9609 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9608)

			gen9610 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9609)

			gen9611 := Call(__e, ShenFunc(symcons), gen9610, Nil)

			gen9612 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9611)

			gen9613 := Call(__e, ShenFunc(symcons), gen9602, gen9612)

			__e.TailApply(ShenFunc(symunify_b), V3566, gen9613, V3567, V3568)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-vector->"), gen9614)

		gen9621 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3576 := __args[0]
			_ = V3576
			V3577 := __args[1]
			_ = V3577
			V3578 := __args[2]
			_ = V3578
			gen9615 := Call(__e, ShenFunc(symshen_4newpv), V3577)

			A := gen9615
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9616 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9617 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9616)

			gen9618 := Call(__e, ShenFunc(symcons), gen9617, Nil)

			gen9619 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9618)

			gen9620 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9619)

			__e.TailApply(ShenFunc(symunify_b), V3576, gen9620, V3577, V3578)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-vector"), gen9621)

		gen9625 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3586 := __args[0]
			_ = V3586
			V3587 := __args[1]
			_ = V3587
			V3588 := __args[2]
			_ = V3588
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9622 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9623 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9622)

			gen9624 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9623)

			__e.TailApply(ShenFunc(symunify_b), V3586, gen9624, V3587, V3588)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-get-time"), gen9625)

		gen9633 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3596 := __args[0]
			_ = V3596
			V3597 := __args[1]
			_ = V3597
			V3598 := __args[2]
			_ = V3598
			gen9626 := Call(__e, ShenFunc(symshen_4newpv), V3597)

			A := gen9626
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9627 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9628 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9627)

			gen9629 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9628)

			gen9630 := Call(__e, ShenFunc(symcons), gen9629, Nil)

			gen9631 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9630)

			gen9632 := Call(__e, ShenFunc(symcons), A, gen9631)

			__e.TailApply(ShenFunc(symunify_b), V3596, gen9632, V3597, V3598)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-hash"), gen9633)

		gen9640 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3606 := __args[0]
			_ = V3606
			V3607 := __args[1]
			_ = V3607
			V3608 := __args[2]
			_ = V3608
			gen9634 := Call(__e, ShenFunc(symshen_4newpv), V3607)

			A := gen9634
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9635 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9636 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9635)

			gen9637 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9638 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9637)

			gen9639 := Call(__e, ShenFunc(symcons), gen9636, gen9638)

			__e.TailApply(ShenFunc(symunify_b), V3606, gen9639, V3607, V3608)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-head"), gen9640)

		gen9647 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3616 := __args[0]
			_ = V3616
			V3617 := __args[1]
			_ = V3617
			V3618 := __args[2]
			_ = V3618
			gen9641 := Call(__e, ShenFunc(symshen_4newpv), V3617)

			A := gen9641
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9642 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9643 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9642)

			gen9644 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9645 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9644)

			gen9646 := Call(__e, ShenFunc(symcons), gen9643, gen9645)

			__e.TailApply(ShenFunc(symunify_b), V3616, gen9646, V3617, V3618)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-hdv"), gen9647)

		gen9651 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3626 := __args[0]
			_ = V3626
			V3627 := __args[1]
			_ = V3627
			V3628 := __args[2]
			_ = V3628
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9648 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9649 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9648)

			gen9650 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9649)

			__e.TailApply(ShenFunc(symunify_b), V3626, gen9650, V3627, V3628)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-hdstr"), gen9651)

		gen9662 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3636 := __args[0]
			_ = V3636
			V3637 := __args[1]
			_ = V3637
			V3638 := __args[2]
			_ = V3638
			gen9652 := Call(__e, ShenFunc(symshen_4newpv), V3637)

			A := gen9652
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9653 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9654 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9653)

			gen9655 := Call(__e, ShenFunc(symcons), A, gen9654)

			gen9656 := Call(__e, ShenFunc(symcons), gen9655, Nil)

			gen9657 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9656)

			gen9658 := Call(__e, ShenFunc(symcons), A, gen9657)

			gen9659 := Call(__e, ShenFunc(symcons), gen9658, Nil)

			gen9660 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9659)

			gen9661 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9660)

			__e.TailApply(ShenFunc(symunify_b), V3636, gen9661, V3637, V3638)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-if"), gen9662)

		gen9665 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3646 := __args[0]
			_ = V3646
			V3647 := __args[1]
			_ = V3647
			V3648 := __args[2]
			_ = V3648
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9663 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9664 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9663)

			__e.TailApply(ShenFunc(symunify_b), V3646, gen9664, V3647, V3648)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-it"), gen9665)

		gen9668 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3656 := __args[0]
			_ = V3656
			V3657 := __args[1]
			_ = V3657
			V3658 := __args[2]
			_ = V3658
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9666 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9667 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9666)

			__e.TailApply(ShenFunc(symunify_b), V3656, gen9667, V3657, V3658)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-implementation"), gen9668)

		gen9676 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3666 := __args[0]
			_ = V3666
			V3667 := __args[1]
			_ = V3667
			V3668 := __args[2]
			_ = V3668
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9669 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9670 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9669)

			gen9671 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9672 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9671)

			gen9673 := Call(__e, ShenFunc(symcons), gen9672, Nil)

			gen9674 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9673)

			gen9675 := Call(__e, ShenFunc(symcons), gen9670, gen9674)

			__e.TailApply(ShenFunc(symunify_b), V3666, gen9675, V3667, V3668)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-include"), gen9676)

		gen9684 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3676 := __args[0]
			_ = V3676
			V3677 := __args[1]
			_ = V3677
			V3678 := __args[2]
			_ = V3678
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9677 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9678 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9677)

			gen9679 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9680 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9679)

			gen9681 := Call(__e, ShenFunc(symcons), gen9680, Nil)

			gen9682 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9681)

			gen9683 := Call(__e, ShenFunc(symcons), gen9678, gen9682)

			__e.TailApply(ShenFunc(symunify_b), V3676, gen9683, V3677, V3678)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-include-all-but"), gen9684)

		gen9687 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3686 := __args[0]
			_ = V3686
			V3687 := __args[1]
			_ = V3687
			V3688 := __args[2]
			_ = V3688
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9685 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9686 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9685)

			__e.TailApply(ShenFunc(symunify_b), V3686, gen9686, V3687, V3688)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-inferences"), gen9687)

		gen9695 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3696 := __args[0]
			_ = V3696
			V3697 := __args[1]
			_ = V3697
			V3698 := __args[2]
			_ = V3698
			gen9688 := Call(__e, ShenFunc(symshen_4newpv), V3697)

			A := gen9688
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9689 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9690 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9689)

			gen9691 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9690)

			gen9692 := Call(__e, ShenFunc(symcons), gen9691, Nil)

			gen9693 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9692)

			gen9694 := Call(__e, ShenFunc(symcons), A, gen9693)

			__e.TailApply(ShenFunc(symunify_b), V3696, gen9694, V3697, V3698)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-shen.insert"), gen9695)

		gen9700 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3706 := __args[0]
			_ = V3706
			V3707 := __args[1]
			_ = V3707
			V3708 := __args[2]
			_ = V3708
			gen9696 := Call(__e, ShenFunc(symshen_4newpv), V3707)

			A := gen9696
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9697 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9698 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9697)

			gen9699 := Call(__e, ShenFunc(symcons), A, gen9698)

			__e.TailApply(ShenFunc(symunify_b), V3706, gen9699, V3707, V3708)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-integer?"), gen9700)

		gen9706 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3716 := __args[0]
			_ = V3716
			V3717 := __args[1]
			_ = V3717
			V3718 := __args[2]
			_ = V3718
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9701 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9702 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9701)

			gen9703 := Call(__e, ShenFunc(symcons), gen9702, Nil)

			gen9704 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9703)

			gen9705 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9704)

			__e.TailApply(ShenFunc(symunify_b), V3716, gen9705, V3717, V3718)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-internal"), gen9706)

		gen9720 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3726 := __args[0]
			_ = V3726
			V3727 := __args[1]
			_ = V3727
			V3728 := __args[2]
			_ = V3728
			gen9707 := Call(__e, ShenFunc(symshen_4newpv), V3727)

			A := gen9707
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9708 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9709 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9708)

			gen9710 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9711 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9710)

			gen9712 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9713 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9712)

			gen9714 := Call(__e, ShenFunc(symcons), gen9713, Nil)

			gen9715 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9714)

			gen9716 := Call(__e, ShenFunc(symcons), gen9711, gen9715)

			gen9717 := Call(__e, ShenFunc(symcons), gen9716, Nil)

			gen9718 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9717)

			gen9719 := Call(__e, ShenFunc(symcons), gen9709, gen9718)

			__e.TailApply(ShenFunc(symunify_b), V3726, gen9719, V3727, V3728)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-intersection"), gen9720)

		gen9724 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3736 := __args[0]
			_ = V3736
			V3737 := __args[1]
			_ = V3737
			V3738 := __args[2]
			_ = V3738
			gen9721 := Call(__e, ShenFunc(symshen_4newpv), V3737)

			A := gen9721
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9722 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9723 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9722)

			__e.TailApply(ShenFunc(symunify_b), V3736, gen9723, V3737, V3738)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-kill"), gen9724)

		gen9727 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3746 := __args[0]
			_ = V3746
			V3747 := __args[1]
			_ = V3747
			V3748 := __args[2]
			_ = V3748
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9725 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9726 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9725)

			__e.TailApply(ShenFunc(symunify_b), V3746, gen9726, V3747, V3748)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-language"), gen9727)

		gen9734 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3756 := __args[0]
			_ = V3756
			V3757 := __args[1]
			_ = V3757
			V3758 := __args[2]
			_ = V3758
			gen9728 := Call(__e, ShenFunc(symshen_4newpv), V3757)

			A := gen9728
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9729 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9730 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9729)

			gen9731 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9732 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9731)

			gen9733 := Call(__e, ShenFunc(symcons), gen9730, gen9732)

			__e.TailApply(ShenFunc(symunify_b), V3756, gen9733, V3757, V3758)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-length"), gen9734)

		gen9741 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3766 := __args[0]
			_ = V3766
			V3767 := __args[1]
			_ = V3767
			V3768 := __args[2]
			_ = V3768
			gen9735 := Call(__e, ShenFunc(symshen_4newpv), V3767)

			A := gen9735
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9736 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9737 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen9736)

			gen9738 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9739 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9738)

			gen9740 := Call(__e, ShenFunc(symcons), gen9737, gen9739)

			__e.TailApply(ShenFunc(symunify_b), V3766, gen9740, V3767, V3768)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-limit"), gen9741)

		gen9745 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3776 := __args[0]
			_ = V3776
			V3777 := __args[1]
			_ = V3777
			V3778 := __args[2]
			_ = V3778
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9742 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9743 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9742)

			gen9744 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9743)

			__e.TailApply(ShenFunc(symunify_b), V3776, gen9744, V3777, V3778)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-load"), gen9745)

		gen9761 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3786 := __args[0]
			_ = V3786
			V3787 := __args[1]
			_ = V3787
			V3788 := __args[2]
			_ = V3788
			gen9746 := Call(__e, ShenFunc(symshen_4newpv), V3787)

			A := gen9746
			gen9747 := Call(__e, ShenFunc(symshen_4newpv), V3787)

			B := gen9747
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9748 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9749 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9748)

			gen9750 := Call(__e, ShenFunc(symcons), A, gen9749)

			gen9751 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9752 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9751)

			gen9753 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9754 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9753)

			gen9755 := Call(__e, ShenFunc(symcons), gen9754, Nil)

			gen9756 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9755)

			gen9757 := Call(__e, ShenFunc(symcons), gen9752, gen9756)

			gen9758 := Call(__e, ShenFunc(symcons), gen9757, Nil)

			gen9759 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9758)

			gen9760 := Call(__e, ShenFunc(symcons), gen9750, gen9759)

			__e.TailApply(ShenFunc(symunify_b), V3786, gen9760, V3787, V3788)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-map"), gen9761)

		gen9779 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3796 := __args[0]
			_ = V3796
			V3797 := __args[1]
			_ = V3797
			V3798 := __args[2]
			_ = V3798
			gen9762 := Call(__e, ShenFunc(symshen_4newpv), V3797)

			A := gen9762
			gen9763 := Call(__e, ShenFunc(symshen_4newpv), V3797)

			B := gen9763
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9764 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9765 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9764)

			gen9766 := Call(__e, ShenFunc(symcons), gen9765, Nil)

			gen9767 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9766)

			gen9768 := Call(__e, ShenFunc(symcons), A, gen9767)

			gen9769 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9770 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9769)

			gen9771 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9772 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9771)

			gen9773 := Call(__e, ShenFunc(symcons), gen9772, Nil)

			gen9774 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9773)

			gen9775 := Call(__e, ShenFunc(symcons), gen9770, gen9774)

			gen9776 := Call(__e, ShenFunc(symcons), gen9775, Nil)

			gen9777 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9776)

			gen9778 := Call(__e, ShenFunc(symcons), gen9768, gen9777)

			__e.TailApply(ShenFunc(symunify_b), V3796, gen9778, V3797, V3798)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-mapcan"), gen9779)

		gen9783 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3806 := __args[0]
			_ = V3806
			V3807 := __args[1]
			_ = V3807
			V3808 := __args[2]
			_ = V3808
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9780 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9781 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9780)

			gen9782 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9781)

			__e.TailApply(ShenFunc(symunify_b), V3806, gen9782, V3807, V3808)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-maxinferences"), gen9783)

		gen9787 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3816 := __args[0]
			_ = V3816
			V3817 := __args[1]
			_ = V3817
			V3818 := __args[2]
			_ = V3818
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9784 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9785 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9784)

			gen9786 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9785)

			__e.TailApply(ShenFunc(symunify_b), V3816, gen9786, V3817, V3818)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-n->string"), gen9787)

		gen9791 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3826 := __args[0]
			_ = V3826
			V3827 := __args[1]
			_ = V3827
			V3828 := __args[2]
			_ = V3828
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9788 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9789 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9788)

			gen9790 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9789)

			__e.TailApply(ShenFunc(symunify_b), V3826, gen9790, V3827, V3828)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-nl"), gen9791)

		gen9795 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3836 := __args[0]
			_ = V3836
			V3837 := __args[1]
			_ = V3837
			V3838 := __args[2]
			_ = V3838
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9792 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9793 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9792)

			gen9794 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9793)

			__e.TailApply(ShenFunc(symunify_b), V3836, gen9794, V3837, V3838)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-not"), gen9795)

		gen9805 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3846 := __args[0]
			_ = V3846
			V3847 := __args[1]
			_ = V3847
			V3848 := __args[2]
			_ = V3848
			gen9796 := Call(__e, ShenFunc(symshen_4newpv), V3847)

			A := gen9796
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9797 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9798 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9797)

			gen9799 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9800 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9799)

			gen9801 := Call(__e, ShenFunc(symcons), gen9798, gen9800)

			gen9802 := Call(__e, ShenFunc(symcons), gen9801, Nil)

			gen9803 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9802)

			gen9804 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9803)

			__e.TailApply(ShenFunc(symunify_b), V3846, gen9804, V3847, V3848)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-nth"), gen9805)

		gen9810 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3856 := __args[0]
			_ = V3856
			V3857 := __args[1]
			_ = V3857
			V3858 := __args[2]
			_ = V3858
			gen9806 := Call(__e, ShenFunc(symshen_4newpv), V3857)

			A := gen9806
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9807 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9808 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9807)

			gen9809 := Call(__e, ShenFunc(symcons), A, gen9808)

			__e.TailApply(ShenFunc(symunify_b), V3856, gen9809, V3857, V3858)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-number?"), gen9810)

		gen9819 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3866 := __args[0]
			_ = V3866
			V3867 := __args[1]
			_ = V3867
			V3868 := __args[2]
			_ = V3868
			gen9811 := Call(__e, ShenFunc(symshen_4newpv), V3867)

			A := gen9811
			gen9812 := Call(__e, ShenFunc(symshen_4newpv), V3867)

			B := gen9812
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9813 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9814 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9813)

			gen9815 := Call(__e, ShenFunc(symcons), B, gen9814)

			gen9816 := Call(__e, ShenFunc(symcons), gen9815, Nil)

			gen9817 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9816)

			gen9818 := Call(__e, ShenFunc(symcons), A, gen9817)

			__e.TailApply(ShenFunc(symunify_b), V3866, gen9818, V3867, V3868)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-occurrences"), gen9819)

		gen9823 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3876 := __args[0]
			_ = V3876
			V3877 := __args[1]
			_ = V3877
			V3878 := __args[2]
			_ = V3878
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9820 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9821 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9820)

			gen9822 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9821)

			__e.TailApply(ShenFunc(symunify_b), V3876, gen9822, V3877, V3878)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-occurs-check"), gen9823)

		gen9827 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3886 := __args[0]
			_ = V3886
			V3887 := __args[1]
			_ = V3887
			V3888 := __args[2]
			_ = V3888
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9824 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9825 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9824)

			gen9826 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9825)

			__e.TailApply(ShenFunc(symunify_b), V3886, gen9826, V3887, V3888)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-optimise"), gen9827)

		gen9834 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3896 := __args[0]
			_ = V3896
			V3897 := __args[1]
			_ = V3897
			V3898 := __args[2]
			_ = V3898
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9828 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9829 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9828)

			gen9830 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9829)

			gen9831 := Call(__e, ShenFunc(symcons), gen9830, Nil)

			gen9832 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9831)

			gen9833 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen9832)

			__e.TailApply(ShenFunc(symunify_b), V3896, gen9833, V3897, V3898)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-or"), gen9834)

		gen9837 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3906 := __args[0]
			_ = V3906
			V3907 := __args[1]
			_ = V3907
			V3908 := __args[2]
			_ = V3908
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9835 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9836 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9835)

			__e.TailApply(ShenFunc(symunify_b), V3906, gen9836, V3907, V3908)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-os"), gen9837)

		gen9841 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3916 := __args[0]
			_ = V3916
			V3917 := __args[1]
			_ = V3917
			V3918 := __args[2]
			_ = V3918
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9838 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen9839 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9838)

			gen9840 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9839)

			__e.TailApply(ShenFunc(symunify_b), V3916, gen9840, V3917, V3918)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-package?"), gen9841)

		gen9844 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3926 := __args[0]
			_ = V3926
			V3927 := __args[1]
			_ = V3927
			V3928 := __args[2]
			_ = V3928
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9842 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9843 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9842)

			__e.TailApply(ShenFunc(symunify_b), V3926, gen9843, V3927, V3928)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-port"), gen9844)

		gen9847 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3936 := __args[0]
			_ = V3936
			V3937 := __args[1]
			_ = V3937
			V3938 := __args[2]
			_ = V3938
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9845 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9846 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9845)

			__e.TailApply(ShenFunc(symunify_b), V3936, gen9846, V3937, V3938)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-porters"), gen9847)

		gen9854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3946 := __args[0]
			_ = V3946
			V3947 := __args[1]
			_ = V3947
			V3948 := __args[2]
			_ = V3948
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9848 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9849 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9848)

			gen9850 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen9849)

			gen9851 := Call(__e, ShenFunc(symcons), gen9850, Nil)

			gen9852 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9851)

			gen9853 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9852)

			__e.TailApply(ShenFunc(symunify_b), V3946, gen9853, V3947, V3948)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-pos"), gen9854)

		gen9863 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3956 := __args[0]
			_ = V3956
			V3957 := __args[1]
			_ = V3957
			V3958 := __args[2]
			_ = V3958
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9855 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen9856 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen9855)

			gen9857 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9858 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9857)

			gen9859 := Call(__e, ShenFunc(symcons), gen9856, gen9858)

			gen9860 := Call(__e, ShenFunc(symcons), gen9859, Nil)

			gen9861 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9860)

			gen9862 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9861)

			__e.TailApply(ShenFunc(symunify_b), V3956, gen9862, V3957, V3958)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-pr"), gen9863)

		gen9868 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3966 := __args[0]
			_ = V3966
			V3967 := __args[1]
			_ = V3967
			V3968 := __args[2]
			_ = V3968
			gen9864 := Call(__e, ShenFunc(symshen_4newpv), V3967)

			A := gen9864
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9865 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9866 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9865)

			gen9867 := Call(__e, ShenFunc(symcons), A, gen9866)

			__e.TailApply(ShenFunc(symunify_b), V3966, gen9867, V3967, V3968)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-print"), gen9868)

		gen9880 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3976 := __args[0]
			_ = V3976
			V3977 := __args[1]
			_ = V3977
			V3978 := __args[2]
			_ = V3978
			gen9869 := Call(__e, ShenFunc(symshen_4newpv), V3977)

			A := gen9869
			gen9870 := Call(__e, ShenFunc(symshen_4newpv), V3977)

			B := gen9870
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9871 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9872 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9871)

			gen9873 := Call(__e, ShenFunc(symcons), A, gen9872)

			gen9874 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9875 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9874)

			gen9876 := Call(__e, ShenFunc(symcons), A, gen9875)

			gen9877 := Call(__e, ShenFunc(symcons), gen9876, Nil)

			gen9878 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9877)

			gen9879 := Call(__e, ShenFunc(symcons), gen9873, gen9878)

			__e.TailApply(ShenFunc(symunify_b), V3976, gen9879, V3977, V3978)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-profile"), gen9880)

		gen9888 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3986 := __args[0]
			_ = V3986
			V3987 := __args[1]
			_ = V3987
			V3988 := __args[2]
			_ = V3988
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9881 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9882 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9881)

			gen9883 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9884 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9883)

			gen9885 := Call(__e, ShenFunc(symcons), gen9884, Nil)

			gen9886 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9885)

			gen9887 := Call(__e, ShenFunc(symcons), gen9882, gen9886)

			__e.TailApply(ShenFunc(symunify_b), V3986, gen9887, V3987, V3988)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-preclude"), gen9888)

		gen9892 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3996 := __args[0]
			_ = V3996
			V3997 := __args[1]
			_ = V3997
			V3998 := __args[2]
			_ = V3998
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9889 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9890 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9889)

			gen9891 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9890)

			__e.TailApply(ShenFunc(symunify_b), V3996, gen9891, V3997, V3998)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-shen.proc-nl"), gen9892)

		gen9907 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4006 := __args[0]
			_ = V4006
			V4007 := __args[1]
			_ = V4007
			V4008 := __args[2]
			_ = V4008
			gen9893 := Call(__e, ShenFunc(symshen_4newpv), V4007)

			A := gen9893
			gen9894 := Call(__e, ShenFunc(symshen_4newpv), V4007)

			B := gen9894
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9895 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9896 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9895)

			gen9897 := Call(__e, ShenFunc(symcons), A, gen9896)

			gen9898 := Call(__e, ShenFunc(symcons), B, Nil)

			gen9899 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9898)

			gen9900 := Call(__e, ShenFunc(symcons), A, gen9899)

			gen9901 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9902 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen9901)

			gen9903 := Call(__e, ShenFunc(symcons), gen9900, gen9902)

			gen9904 := Call(__e, ShenFunc(symcons), gen9903, Nil)

			gen9905 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9904)

			gen9906 := Call(__e, ShenFunc(symcons), gen9897, gen9905)

			__e.TailApply(ShenFunc(symunify_b), V4006, gen9906, V4007, V4008)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-profile-results"), gen9907)

		gen9911 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4016 := __args[0]
			_ = V4016
			V4017 := __args[1]
			_ = V4017
			V4018 := __args[2]
			_ = V4018
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9908 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9909 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9908)

			gen9910 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9909)

			__e.TailApply(ShenFunc(symunify_b), V4016, gen9910, V4017, V4018)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-protect"), gen9911)

		gen9919 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4026 := __args[0]
			_ = V4026
			V4027 := __args[1]
			_ = V4027
			V4028 := __args[2]
			_ = V4028
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9912 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9913 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9912)

			gen9914 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen9915 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9914)

			gen9916 := Call(__e, ShenFunc(symcons), gen9915, Nil)

			gen9917 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9916)

			gen9918 := Call(__e, ShenFunc(symcons), gen9913, gen9917)

			__e.TailApply(ShenFunc(symunify_b), V4026, gen9918, V4027, V4028)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-preclude-all-but"), gen9919)

		gen9928 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4036 := __args[0]
			_ = V4036
			V4037 := __args[1]
			_ = V4037
			V4038 := __args[2]
			_ = V4038
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9920 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen9921 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen9920)

			gen9922 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9923 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9922)

			gen9924 := Call(__e, ShenFunc(symcons), gen9921, gen9923)

			gen9925 := Call(__e, ShenFunc(symcons), gen9924, Nil)

			gen9926 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9925)

			gen9927 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9926)

			__e.TailApply(ShenFunc(symunify_b), V4036, gen9927, V4037, V4038)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-shen.prhush"), gen9928)

		gen9934 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4046 := __args[0]
			_ = V4046
			V4047 := __args[1]
			_ = V4047
			V4048 := __args[2]
			_ = V4048
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9929 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen9930 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9929)

			gen9931 := Call(__e, ShenFunc(symcons), gen9930, Nil)

			gen9932 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9931)

			gen9933 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen9932)

			__e.TailApply(ShenFunc(symunify_b), V4046, gen9933, V4047, V4048)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-ps"), gen9934)

		gen9940 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4056 := __args[0]
			_ = V4056
			V4057 := __args[1]
			_ = V4057
			V4058 := __args[2]
			_ = V4058
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9935 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen9936 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen9935)

			gen9937 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen9938 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9937)

			gen9939 := Call(__e, ShenFunc(symcons), gen9936, gen9938)

			__e.TailApply(ShenFunc(symunify_b), V4056, gen9939, V4057, V4058)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read"), gen9940)

		gen9946 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4066 := __args[0]
			_ = V4066
			V4067 := __args[1]
			_ = V4067
			V4068 := __args[2]
			_ = V4068
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9941 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen9942 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen9941)

			gen9943 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9944 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9943)

			gen9945 := Call(__e, ShenFunc(symcons), gen9942, gen9944)

			__e.TailApply(ShenFunc(symunify_b), V4066, gen9945, V4067, V4068)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read-byte"), gen9946)

		gen9952 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4076 := __args[0]
			_ = V4076
			V4077 := __args[1]
			_ = V4077
			V4078 := __args[2]
			_ = V4078
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9947 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen9948 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9947)

			gen9949 := Call(__e, ShenFunc(symcons), gen9948, Nil)

			gen9950 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9949)

			gen9951 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9950)

			__e.TailApply(ShenFunc(symunify_b), V4076, gen9951, V4077, V4078)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read-file-as-bytelist"), gen9952)

		gen9956 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4086 := __args[0]
			_ = V4086
			V4087 := __args[1]
			_ = V4087
			V4088 := __args[2]
			_ = V4088
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9953 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9954 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9953)

			gen9955 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9954)

			__e.TailApply(ShenFunc(symunify_b), V4086, gen9955, V4087, V4088)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read-file-as-string"), gen9956)

		gen9962 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4096 := __args[0]
			_ = V4096
			V4097 := __args[1]
			_ = V4097
			V4098 := __args[2]
			_ = V4098
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9957 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen9958 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9957)

			gen9959 := Call(__e, ShenFunc(symcons), gen9958, Nil)

			gen9960 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9959)

			gen9961 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9960)

			__e.TailApply(ShenFunc(symunify_b), V4096, gen9961, V4097, V4098)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read-file"), gen9962)

		gen9968 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4106 := __args[0]
			_ = V4106
			V4107 := __args[1]
			_ = V4107
			V4108 := __args[2]
			_ = V4108
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9963 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen9964 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9963)

			gen9965 := Call(__e, ShenFunc(symcons), gen9964, Nil)

			gen9966 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9965)

			gen9967 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9966)

			__e.TailApply(ShenFunc(symunify_b), V4106, gen9967, V4107, V4108)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-read-from-string"), gen9968)

		gen9971 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4116 := __args[0]
			_ = V4116
			V4117 := __args[1]
			_ = V4117
			V4118 := __args[2]
			_ = V4118
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9969 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen9970 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9969)

			__e.TailApply(ShenFunc(symunify_b), V4116, gen9970, V4117, V4118)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-release"), gen9971)

		gen9983 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4126 := __args[0]
			_ = V4126
			V4127 := __args[1]
			_ = V4127
			V4128 := __args[2]
			_ = V4128
			gen9972 := Call(__e, ShenFunc(symshen_4newpv), V4127)

			A := gen9972
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9973 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9974 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9973)

			gen9975 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9976 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9975)

			gen9977 := Call(__e, ShenFunc(symcons), gen9976, Nil)

			gen9978 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9977)

			gen9979 := Call(__e, ShenFunc(symcons), gen9974, gen9978)

			gen9980 := Call(__e, ShenFunc(symcons), gen9979, Nil)

			gen9981 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9980)

			gen9982 := Call(__e, ShenFunc(symcons), A, gen9981)

			__e.TailApply(ShenFunc(symunify_b), V4126, gen9982, V4127, V4128)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-remove"), gen9983)

		gen9992 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4136 := __args[0]
			_ = V4136
			V4137 := __args[1]
			_ = V4137
			V4138 := __args[2]
			_ = V4138
			gen9984 := Call(__e, ShenFunc(symshen_4newpv), V4137)

			A := gen9984
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9985 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9986 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9985)

			gen9987 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9988 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen9987)

			gen9989 := Call(__e, ShenFunc(symcons), gen9988, Nil)

			gen9990 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9989)

			gen9991 := Call(__e, ShenFunc(symcons), gen9986, gen9990)

			__e.TailApply(ShenFunc(symunify_b), V4136, gen9991, V4137, V4138)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-reverse"), gen9992)

		gen9997 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4146 := __args[0]
			_ = V4146
			V4147 := __args[1]
			_ = V4147
			V4148 := __args[2]
			_ = V4148
			gen9993 := Call(__e, ShenFunc(symshen_4newpv), V4147)

			A := gen9993
			Call(__e, ShenFunc(symshen_4incinfs))
			gen9994 := Call(__e, ShenFunc(symcons), A, Nil)

			gen9995 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen9994)

			gen9996 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen9995)

			__e.TailApply(ShenFunc(symunify_b), V4146, gen9996, V4147, V4148)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-simple-error"), gen9997)

		gen10006 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4156 := __args[0]
			_ = V4156
			V4157 := __args[1]
			_ = V4157
			V4158 := __args[2]
			_ = V4158
			gen9998 := Call(__e, ShenFunc(symshen_4newpv), V4157)

			A := gen9998
			gen9999 := Call(__e, ShenFunc(symshen_4newpv), V4157)

			B := gen9999
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10000 := Call(__e, ShenFunc(symcons), B, Nil)

			gen10001 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen10000)

			gen10002 := Call(__e, ShenFunc(symcons), A, gen10001)

			gen10003 := Call(__e, ShenFunc(symcons), B, Nil)

			gen10004 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10003)

			gen10005 := Call(__e, ShenFunc(symcons), gen10002, gen10004)

			__e.TailApply(ShenFunc(symunify_b), V4156, gen10005, V4157, V4158)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-snd"), gen10006)

		gen10010 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4166 := __args[0]
			_ = V4166
			V4167 := __args[1]
			_ = V4167
			V4168 := __args[2]
			_ = V4168
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10007 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10008 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10007)

			gen10009 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10008)

			__e.TailApply(ShenFunc(symunify_b), V4166, gen10009, V4167, V4168)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-specialise"), gen10010)

		gen10014 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4176 := __args[0]
			_ = V4176
			V4177 := __args[1]
			_ = V4177
			V4178 := __args[2]
			_ = V4178
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10011 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10012 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10011)

			gen10013 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10012)

			__e.TailApply(ShenFunc(symunify_b), V4176, gen10013, V4177, V4178)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-spy"), gen10014)

		gen10018 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4186 := __args[0]
			_ = V4186
			V4187 := __args[1]
			_ = V4187
			V4188 := __args[2]
			_ = V4188
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10015 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10016 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10015)

			gen10017 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10016)

			__e.TailApply(ShenFunc(symunify_b), V4186, gen10017, V4187, V4188)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-step"), gen10018)

		gen10023 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4196 := __args[0]
			_ = V4196
			V4197 := __args[1]
			_ = V4197
			V4198 := __args[2]
			_ = V4198
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10019 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen10020 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen10019)

			gen10021 := Call(__e, ShenFunc(symcons), gen10020, Nil)

			gen10022 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10021)

			__e.TailApply(ShenFunc(symunify_b), V4196, gen10022, V4197, V4198)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-stinput"), gen10023)

		gen10028 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4206 := __args[0]
			_ = V4206
			V4207 := __args[1]
			_ = V4207
			V4208 := __args[2]
			_ = V4208
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10024 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen10025 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen10024)

			gen10026 := Call(__e, ShenFunc(symcons), gen10025, Nil)

			gen10027 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10026)

			__e.TailApply(ShenFunc(symunify_b), V4206, gen10027, V4207, V4208)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-sterror"), gen10028)

		gen10033 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4216 := __args[0]
			_ = V4216
			V4217 := __args[1]
			_ = V4217
			V4218 := __args[2]
			_ = V4218
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10029 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen10030 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen10029)

			gen10031 := Call(__e, ShenFunc(symcons), gen10030, Nil)

			gen10032 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10031)

			__e.TailApply(ShenFunc(symunify_b), V4216, gen10032, V4217, V4218)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-stoutput"), gen10033)

		gen10038 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4226 := __args[0]
			_ = V4226
			V4227 := __args[1]
			_ = V4227
			V4228 := __args[2]
			_ = V4228
			gen10034 := Call(__e, ShenFunc(symshen_4newpv), V4227)

			A := gen10034
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10035 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10036 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10035)

			gen10037 := Call(__e, ShenFunc(symcons), A, gen10036)

			__e.TailApply(ShenFunc(symunify_b), V4226, gen10037, V4227, V4228)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-string?"), gen10038)

		gen10043 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4236 := __args[0]
			_ = V4236
			V4237 := __args[1]
			_ = V4237
			V4238 := __args[2]
			_ = V4238
			gen10039 := Call(__e, ShenFunc(symshen_4newpv), V4237)

			A := gen10039
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10040 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen10041 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10040)

			gen10042 := Call(__e, ShenFunc(symcons), A, gen10041)

			__e.TailApply(ShenFunc(symunify_b), V4236, gen10042, V4237, V4238)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-str"), gen10043)

		gen10047 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4246 := __args[0]
			_ = V4246
			V4247 := __args[1]
			_ = V4247
			V4248 := __args[2]
			_ = V4248
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10044 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10045 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10044)

			gen10046 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen10045)

			__e.TailApply(ShenFunc(symunify_b), V4246, gen10046, V4247, V4248)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-string->n"), gen10047)

		gen10051 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4256 := __args[0]
			_ = V4256
			V4257 := __args[1]
			_ = V4257
			V4258 := __args[2]
			_ = V4258
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10048 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10049 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10048)

			gen10050 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen10049)

			__e.TailApply(ShenFunc(symunify_b), V4256, gen10050, V4257, V4258)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-string->symbol"), gen10051)

		gen10057 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4266 := __args[0]
			_ = V4266
			V4267 := __args[1]
			_ = V4267
			V4268 := __args[2]
			_ = V4268
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10052 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10053 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10052)

			gen10054 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10055 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10054)

			gen10056 := Call(__e, ShenFunc(symcons), gen10053, gen10055)

			__e.TailApply(ShenFunc(symunify_b), V4266, gen10056, V4267, V4268)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-sum"), gen10057)

		gen10062 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4276 := __args[0]
			_ = V4276
			V4277 := __args[1]
			_ = V4277
			V4278 := __args[2]
			_ = V4278
			gen10058 := Call(__e, ShenFunc(symshen_4newpv), V4277)

			A := gen10058
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10059 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10060 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10059)

			gen10061 := Call(__e, ShenFunc(symcons), A, gen10060)

			__e.TailApply(ShenFunc(symunify_b), V4276, gen10061, V4277, V4278)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-symbol?"), gen10062)

		gen10066 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4286 := __args[0]
			_ = V4286
			V4287 := __args[1]
			_ = V4287
			V4288 := __args[2]
			_ = V4288
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10063 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10064 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10063)

			gen10065 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10064)

			__e.TailApply(ShenFunc(symunify_b), V4286, gen10065, V4287, V4288)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-systemf"), gen10066)

		gen10075 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4296 := __args[0]
			_ = V4296
			V4297 := __args[1]
			_ = V4297
			V4298 := __args[2]
			_ = V4298
			gen10067 := Call(__e, ShenFunc(symshen_4newpv), V4297)

			A := gen10067
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10068 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10069 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10068)

			gen10070 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10071 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10070)

			gen10072 := Call(__e, ShenFunc(symcons), gen10071, Nil)

			gen10073 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10072)

			gen10074 := Call(__e, ShenFunc(symcons), gen10069, gen10073)

			__e.TailApply(ShenFunc(symunify_b), V4296, gen10074, V4297, V4298)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tail"), gen10075)

		gen10079 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4306 := __args[0]
			_ = V4306
			V4307 := __args[1]
			_ = V4307
			V4308 := __args[2]
			_ = V4308
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10076 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen10077 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10076)

			gen10078 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen10077)

			__e.TailApply(ShenFunc(symunify_b), V4306, gen10078, V4307, V4308)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tlstr"), gen10079)

		gen10088 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4316 := __args[0]
			_ = V4316
			V4317 := __args[1]
			_ = V4317
			V4318 := __args[2]
			_ = V4318
			gen10080 := Call(__e, ShenFunc(symshen_4newpv), V4317)

			A := gen10080
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10081 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10082 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen10081)

			gen10083 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10084 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen10083)

			gen10085 := Call(__e, ShenFunc(symcons), gen10084, Nil)

			gen10086 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10085)

			gen10087 := Call(__e, ShenFunc(symcons), gen10082, gen10086)

			__e.TailApply(ShenFunc(symunify_b), V4316, gen10087, V4317, V4318)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tlv"), gen10088)

		gen10092 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4326 := __args[0]
			_ = V4326
			V4327 := __args[1]
			_ = V4327
			V4328 := __args[2]
			_ = V4328
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10089 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10090 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10089)

			gen10091 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10090)

			__e.TailApply(ShenFunc(symunify_b), V4326, gen10091, V4327, V4328)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tc"), gen10092)

		gen10095 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4336 := __args[0]
			_ = V4336
			V4337 := __args[1]
			_ = V4337
			V4338 := __args[2]
			_ = V4338
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10093 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10094 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10093)

			__e.TailApply(ShenFunc(symunify_b), V4336, gen10094, V4337, V4338)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tc?"), gen10095)

		gen10102 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4346 := __args[0]
			_ = V4346
			V4347 := __args[1]
			_ = V4347
			V4348 := __args[2]
			_ = V4348
			gen10096 := Call(__e, ShenFunc(symshen_4newpv), V4347)

			A := gen10096
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10097 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10098 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen10097)

			gen10099 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10100 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10099)

			gen10101 := Call(__e, ShenFunc(symcons), gen10098, gen10100)

			__e.TailApply(ShenFunc(symunify_b), V4346, gen10101, V4347, V4348)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-thaw"), gen10102)

		gen10106 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4356 := __args[0]
			_ = V4356
			V4357 := __args[1]
			_ = V4357
			V4358 := __args[2]
			_ = V4358
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10103 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10104 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10103)

			gen10105 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10104)

			__e.TailApply(ShenFunc(symunify_b), V4356, gen10105, V4357, V4358)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-track"), gen10106)

		gen10117 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4366 := __args[0]
			_ = V4366
			V4367 := __args[1]
			_ = V4367
			V4368 := __args[2]
			_ = V4368
			gen10107 := Call(__e, ShenFunc(symshen_4newpv), V4367)

			A := gen10107
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10108 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10109 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10108)

			gen10110 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen10109)

			gen10111 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10112 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10111)

			gen10113 := Call(__e, ShenFunc(symcons), gen10110, gen10112)

			gen10114 := Call(__e, ShenFunc(symcons), gen10113, Nil)

			gen10115 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10114)

			gen10116 := Call(__e, ShenFunc(symcons), A, gen10115)

			__e.TailApply(ShenFunc(symunify_b), V4366, gen10116, V4367, V4368)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-trap-error"), gen10117)

		gen10122 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4376 := __args[0]
			_ = V4376
			V4377 := __args[1]
			_ = V4377
			V4378 := __args[2]
			_ = V4378
			gen10118 := Call(__e, ShenFunc(symshen_4newpv), V4377)

			A := gen10118
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10119 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10120 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10119)

			gen10121 := Call(__e, ShenFunc(symcons), A, gen10120)

			__e.TailApply(ShenFunc(symunify_b), V4376, gen10121, V4377, V4378)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-tuple?"), gen10122)

		gen10126 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4386 := __args[0]
			_ = V4386
			V4387 := __args[1]
			_ = V4387
			V4388 := __args[2]
			_ = V4388
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10123 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10124 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10123)

			gen10125 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10124)

			__e.TailApply(ShenFunc(symunify_b), V4386, gen10125, V4387, V4388)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-undefmacro"), gen10126)

		gen10140 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4396 := __args[0]
			_ = V4396
			V4397 := __args[1]
			_ = V4397
			V4398 := __args[2]
			_ = V4398
			gen10127 := Call(__e, ShenFunc(symshen_4newpv), V4397)

			A := gen10127
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10128 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10129 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10128)

			gen10130 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10131 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10130)

			gen10132 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10133 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen10132)

			gen10134 := Call(__e, ShenFunc(symcons), gen10133, Nil)

			gen10135 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10134)

			gen10136 := Call(__e, ShenFunc(symcons), gen10131, gen10135)

			gen10137 := Call(__e, ShenFunc(symcons), gen10136, Nil)

			gen10138 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10137)

			gen10139 := Call(__e, ShenFunc(symcons), gen10129, gen10138)

			__e.TailApply(ShenFunc(symunify_b), V4396, gen10139, V4397, V4398)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-union"), gen10140)

		gen10152 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4406 := __args[0]
			_ = V4406
			V4407 := __args[1]
			_ = V4407
			V4408 := __args[2]
			_ = V4408
			gen10141 := Call(__e, ShenFunc(symshen_4newpv), V4407)

			A := gen10141
			gen10142 := Call(__e, ShenFunc(symshen_4newpv), V4407)

			B := gen10142
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10143 := Call(__e, ShenFunc(symcons), B, Nil)

			gen10144 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10143)

			gen10145 := Call(__e, ShenFunc(symcons), A, gen10144)

			gen10146 := Call(__e, ShenFunc(symcons), B, Nil)

			gen10147 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10146)

			gen10148 := Call(__e, ShenFunc(symcons), A, gen10147)

			gen10149 := Call(__e, ShenFunc(symcons), gen10148, Nil)

			gen10150 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10149)

			gen10151 := Call(__e, ShenFunc(symcons), gen10145, gen10150)

			__e.TailApply(ShenFunc(symunify_b), V4406, gen10151, V4407, V4408)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-unprofile"), gen10152)

		gen10156 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4416 := __args[0]
			_ = V4416
			V4417 := __args[1]
			_ = V4417
			V4418 := __args[2]
			_ = V4418
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10153 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10154 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10153)

			gen10155 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10154)

			__e.TailApply(ShenFunc(symunify_b), V4416, gen10155, V4417, V4418)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-untrack"), gen10156)

		gen10160 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4426 := __args[0]
			_ = V4426
			V4427 := __args[1]
			_ = V4427
			V4428 := __args[2]
			_ = V4428
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10157 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen10158 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10157)

			gen10159 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen10158)

			__e.TailApply(ShenFunc(symunify_b), V4426, gen10159, V4427, V4428)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-unspecialise"), gen10160)

		gen10165 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4436 := __args[0]
			_ = V4436
			V4437 := __args[1]
			_ = V4437
			V4438 := __args[2]
			_ = V4438
			gen10161 := Call(__e, ShenFunc(symshen_4newpv), V4437)

			A := gen10161
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10162 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10163 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10162)

			gen10164 := Call(__e, ShenFunc(symcons), A, gen10163)

			__e.TailApply(ShenFunc(symunify_b), V4436, gen10164, V4437, V4438)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-variable?"), gen10165)

		gen10170 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4446 := __args[0]
			_ = V4446
			V4447 := __args[1]
			_ = V4447
			V4448 := __args[2]
			_ = V4448
			gen10166 := Call(__e, ShenFunc(symshen_4newpv), V4447)

			A := gen10166
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10167 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10168 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10167)

			gen10169 := Call(__e, ShenFunc(symcons), A, gen10168)

			__e.TailApply(ShenFunc(symunify_b), V4446, gen10169, V4447, V4448)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-vector?"), gen10170)

		gen10173 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4456 := __args[0]
			_ = V4456
			V4457 := __args[1]
			_ = V4457
			V4458 := __args[2]
			_ = V4458
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10171 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen10172 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10171)

			__e.TailApply(ShenFunc(symunify_b), V4456, gen10172, V4457, V4458)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-version"), gen10173)

		gen10181 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4466 := __args[0]
			_ = V4466
			V4467 := __args[1]
			_ = V4467
			V4468 := __args[2]
			_ = V4468
			gen10174 := Call(__e, ShenFunc(symshen_4newpv), V4467)

			A := gen10174
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10175 := Call(__e, ShenFunc(symcons), A, Nil)

			gen10176 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10175)

			gen10177 := Call(__e, ShenFunc(symcons), A, gen10176)

			gen10178 := Call(__e, ShenFunc(symcons), gen10177, Nil)

			gen10179 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10178)

			gen10180 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen10179)

			__e.TailApply(ShenFunc(symunify_b), V4466, gen10180, V4467, V4468)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-write-to-file"), gen10181)

		gen10190 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4476 := __args[0]
			_ = V4476
			V4477 := __args[1]
			_ = V4477
			V4478 := __args[2]
			_ = V4478
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10182 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen10183 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen10182)

			gen10184 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10185 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10184)

			gen10186 := Call(__e, ShenFunc(symcons), gen10183, gen10185)

			gen10187 := Call(__e, ShenFunc(symcons), gen10186, Nil)

			gen10188 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10187)

			gen10189 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10188)

			__e.TailApply(ShenFunc(symunify_b), V4476, gen10189, V4477, V4478)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-write-byte"), gen10190)

		gen10194 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4486 := __args[0]
			_ = V4486
			V4487 := __args[1]
			_ = V4487
			V4488 := __args[2]
			_ = V4488
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10191 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10192 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10191)

			gen10193 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen10192)

			__e.TailApply(ShenFunc(symunify_b), V4486, gen10193, V4487, V4488)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-y-or-n?"), gen10194)

		gen10201 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4496 := __args[0]
			_ = V4496
			V4497 := __args[1]
			_ = V4497
			V4498 := __args[2]
			_ = V4498
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10195 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10196 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10195)

			gen10197 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10196)

			gen10198 := Call(__e, ShenFunc(symcons), gen10197, Nil)

			gen10199 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10198)

			gen10200 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10199)

			__e.TailApply(ShenFunc(symunify_b), V4496, gen10200, V4497, V4498)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of->"), gen10201)

		gen10208 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4506 := __args[0]
			_ = V4506
			V4507 := __args[1]
			_ = V4507
			V4508 := __args[2]
			_ = V4508
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10202 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10203 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10202)

			gen10204 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10203)

			gen10205 := Call(__e, ShenFunc(symcons), gen10204, Nil)

			gen10206 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10205)

			gen10207 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10206)

			__e.TailApply(ShenFunc(symunify_b), V4506, gen10207, V4507, V4508)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-<"), gen10208)

		gen10215 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4516 := __args[0]
			_ = V4516
			V4517 := __args[1]
			_ = V4517
			V4518 := __args[2]
			_ = V4518
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10209 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10210 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10209)

			gen10211 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10210)

			gen10212 := Call(__e, ShenFunc(symcons), gen10211, Nil)

			gen10213 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10212)

			gen10214 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10213)

			__e.TailApply(ShenFunc(symunify_b), V4516, gen10214, V4517, V4518)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of->="), gen10215)

		gen10222 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4526 := __args[0]
			_ = V4526
			V4527 := __args[1]
			_ = V4527
			V4528 := __args[2]
			_ = V4528
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10216 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10217 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10216)

			gen10218 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10217)

			gen10219 := Call(__e, ShenFunc(symcons), gen10218, Nil)

			gen10220 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10219)

			gen10221 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10220)

			__e.TailApply(ShenFunc(symunify_b), V4526, gen10221, V4527, V4528)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-<="), gen10222)

		gen10230 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4536 := __args[0]
			_ = V4536
			V4537 := __args[1]
			_ = V4537
			V4538 := __args[2]
			_ = V4538
			gen10223 := Call(__e, ShenFunc(symshen_4newpv), V4537)

			A := gen10223
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10224 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10225 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10224)

			gen10226 := Call(__e, ShenFunc(symcons), A, gen10225)

			gen10227 := Call(__e, ShenFunc(symcons), gen10226, Nil)

			gen10228 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10227)

			gen10229 := Call(__e, ShenFunc(symcons), A, gen10228)

			__e.TailApply(ShenFunc(symunify_b), V4536, gen10229, V4537, V4538)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-="), gen10230)

		gen10237 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4546 := __args[0]
			_ = V4546
			V4547 := __args[1]
			_ = V4547
			V4548 := __args[2]
			_ = V4548
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10231 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10232 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10231)

			gen10233 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10232)

			gen10234 := Call(__e, ShenFunc(symcons), gen10233, Nil)

			gen10235 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10234)

			gen10236 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10235)

			__e.TailApply(ShenFunc(symunify_b), V4546, gen10236, V4547, V4548)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-+"), gen10237)

		gen10244 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4556 := __args[0]
			_ = V4556
			V4557 := __args[1]
			_ = V4557
			V4558 := __args[2]
			_ = V4558
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10238 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10239 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10238)

			gen10240 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10239)

			gen10241 := Call(__e, ShenFunc(symcons), gen10240, Nil)

			gen10242 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10241)

			gen10243 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10242)

			__e.TailApply(ShenFunc(symunify_b), V4556, gen10243, V4557, V4558)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-/"), gen10244)

		gen10251 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4566 := __args[0]
			_ = V4566
			V4567 := __args[1]
			_ = V4567
			V4568 := __args[2]
			_ = V4568
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10245 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10246 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10245)

			gen10247 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10246)

			gen10248 := Call(__e, ShenFunc(symcons), gen10247, Nil)

			gen10249 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10248)

			gen10250 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10249)

			__e.TailApply(ShenFunc(symunify_b), V4566, gen10250, V4567, V4568)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of--"), gen10251)

		gen10258 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4576 := __args[0]
			_ = V4576
			V4577 := __args[1]
			_ = V4577
			V4578 := __args[2]
			_ = V4578
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10252 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen10253 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10252)

			gen10254 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10253)

			gen10255 := Call(__e, ShenFunc(symcons), gen10254, Nil)

			gen10256 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10255)

			gen10257 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen10256)

			__e.TailApply(ShenFunc(symunify_b), V4576, gen10257, V4577, V4578)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-*"), gen10258)

		gen10267 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4586 := __args[0]
			_ = V4586
			V4587 := __args[1]
			_ = V4587
			V4588 := __args[2]
			_ = V4588
			gen10259 := Call(__e, ShenFunc(symshen_4newpv), V4587)

			A := gen10259
			gen10260 := Call(__e, ShenFunc(symshen_4newpv), V4587)

			B := gen10260
			Call(__e, ShenFunc(symshen_4incinfs))
			gen10261 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen10262 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10261)

			gen10263 := Call(__e, ShenFunc(symcons), B, gen10262)

			gen10264 := Call(__e, ShenFunc(symcons), gen10263, Nil)

			gen10265 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen10264)

			gen10266 := Call(__e, ShenFunc(symcons), A, gen10265)

			__e.TailApply(ShenFunc(symunify_b), V4586, gen10266, V4587, V4588)

			return

		}, 3)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.type-signature-of-=="), gen10267)

		return

	}, 0))
}
