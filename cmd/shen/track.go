package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen9307 := MakeNative(func(__e Evaluator) {
			V3135 := __e.Get(1)
			_ = V3135
			gen9297 := Call(__e, ShenFunc(symshen_4app), V3135, MakeString(";\n"), MakeSymbol("shen.a"))

			gen9298 := Call(__e, ShenFunc(symcn), MakeString("partial function "), gen9297)

			gen9299 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen9298, gen9299)

			gen9304 := Call(__e, ShenFunc(symshen_4tracked_2), V3135)

			gen9305 := Call(__e, ShenFunc(symnot), gen9304)

			var gen9306 Obj
			if True == gen9305 {
				gen9301 := Call(__e, ShenFunc(symshen_4app), V3135, MakeString("? "), MakeSymbol("shen.a"))

				gen9302 := Call(__e, ShenFunc(symcn), MakeString("track "), gen9301)

				gen9303 := Call(__e, ShenFunc(symy_1or_1n_2), gen9302)

				if True == gen9303 {
					gen9306 = True
				} else {
					gen9306 = False
				}

			} else {
				gen9306 = False
			}
			if True == gen9306 {
				gen9300 := Call(__e, ShenFunc(symps), V3135)

				Call(__e, ShenFunc(symshen_4track_1function), gen9300)

			} else {
				MakeSymbol("shen.ok")
			}

			__e.TailApply(ShenFunc(symsimple_1error), MakeString("aborted"))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.f_error"), gen9307)

		gen9309 := MakeNative(func(__e Evaluator) {
			V3137 := __e.Get(1)
			_ = V3137
			gen9308 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

			__e.TailApply(ShenFunc(symelement_2), V3137, gen9308)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tracked?"), gen9309)

		gen9311 := MakeNative(func(__e Evaluator) {
			V3139 := __e.Get(1)
			_ = V3139
			gen9310 := Call(__e, ShenFunc(symps), V3139)

			Source := gen9310
			__e.TailApply(ShenFunc(symshen_4track_1function), Source)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("track"), gen9311)

		gen9357 := MakeNative(func(__e Evaluator) {
			V3141 := __e.Get(1)
			_ = V3141
			gen9355 := Call(__e, ShenFunc(symcons_2), V3141)

			var gen9356 Obj
			if True == gen9355 {
				gen9352 := Call(__e, ShenFunc(symhd), V3141)

				gen9353 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen9352)

				var gen9354 Obj
				if True == gen9353 {
					gen9349 := Call(__e, ShenFunc(symtl), V3141)

					gen9350 := Call(__e, ShenFunc(symcons_2), gen9349)

					var gen9351 Obj
					if True == gen9350 {
						gen9345 := Call(__e, ShenFunc(symtl), V3141)

						gen9346 := Call(__e, ShenFunc(symtl), gen9345)

						gen9347 := Call(__e, ShenFunc(symcons_2), gen9346)

						var gen9348 Obj
						if True == gen9347 {
							gen9340 := Call(__e, ShenFunc(symtl), V3141)

							gen9341 := Call(__e, ShenFunc(symtl), gen9340)

							gen9342 := Call(__e, ShenFunc(symtl), gen9341)

							gen9343 := Call(__e, ShenFunc(symcons_2), gen9342)

							var gen9344 Obj
							if True == gen9343 {
								gen9335 := Call(__e, ShenFunc(symtl), V3141)

								gen9336 := Call(__e, ShenFunc(symtl), gen9335)

								gen9337 := Call(__e, ShenFunc(symtl), gen9336)

								gen9338 := Call(__e, ShenFunc(symtl), gen9337)

								gen9339 := Call(__e, ShenFunc(sym_a), Nil, gen9338)

								if True == gen9339 {
									gen9344 = True
								} else {
									gen9344 = False
								}

							} else {
								gen9344 = False
							}
							if True == gen9344 {
								gen9348 = True
							} else {
								gen9348 = False
							}

						} else {
							gen9348 = False
						}
						if True == gen9348 {
							gen9351 = True
						} else {
							gen9351 = False
						}

					} else {
						gen9351 = False
					}
					if True == gen9351 {
						gen9354 = True
					} else {
						gen9354 = False
					}

				} else {
					gen9354 = False
				}
				if True == gen9354 {
					gen9356 = True
				} else {
					gen9356 = False
				}

			} else {
				gen9356 = False
			}
			if True == gen9356 {
				gen9312 := Call(__e, ShenFunc(symtl), V3141)

				gen9313 := Call(__e, ShenFunc(symhd), gen9312)

				gen9314 := Call(__e, ShenFunc(symtl), V3141)

				gen9315 := Call(__e, ShenFunc(symtl), gen9314)

				gen9316 := Call(__e, ShenFunc(symhd), gen9315)

				gen9317 := Call(__e, ShenFunc(symtl), V3141)

				gen9318 := Call(__e, ShenFunc(symhd), gen9317)

				gen9319 := Call(__e, ShenFunc(symtl), V3141)

				gen9320 := Call(__e, ShenFunc(symtl), gen9319)

				gen9321 := Call(__e, ShenFunc(symhd), gen9320)

				gen9322 := Call(__e, ShenFunc(symtl), V3141)

				gen9323 := Call(__e, ShenFunc(symtl), gen9322)

				gen9324 := Call(__e, ShenFunc(symtl), gen9323)

				gen9325 := Call(__e, ShenFunc(symhd), gen9324)

				gen9326 := Call(__e, ShenFunc(symshen_4insert_1tracking_1code), gen9318, gen9321, gen9325)

				gen9327 := Call(__e, ShenFunc(symcons), gen9326, Nil)

				gen9328 := Call(__e, ShenFunc(symcons), gen9316, gen9327)

				gen9329 := Call(__e, ShenFunc(symcons), gen9313, gen9328)

				gen9330 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9329)

				KL := gen9330
				gen9331 := Call(__e, ShenFunc(symeval_1kl), KL)

				Ob := gen9331
				gen9332 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

				gen9333 := Call(__e, ShenFunc(symcons), Ob, gen9332)

				gen9334 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), gen9333)

				Tr := gen9334
				_ = Tr
				__e.Return(Ob)
				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.track-function"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.track-function"), gen9357)

		gen9410 := MakeNative(func(__e Evaluator) {
			V3145 := __e.Get(1)
			_ = V3145
			V3146 := __e.Get(2)
			_ = V3146
			V3147 := __e.Get(3)
			_ = V3147
			gen9358 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9359 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9358)

			gen9360 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen9361 := Call(__e, ShenFunc(symcons), gen9359, gen9360)

			gen9362 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen9361)

			gen9363 := Call(__e, ShenFunc(symcons), gen9362, Nil)

			gen9364 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), gen9363)

			gen9365 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen9364)

			gen9366 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9367 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9366)

			gen9368 := Call(__e, ShenFunc(symshen_4cons__form), V3146)

			gen9369 := Call(__e, ShenFunc(symcons), gen9368, Nil)

			gen9370 := Call(__e, ShenFunc(symcons), V3145, gen9369)

			gen9371 := Call(__e, ShenFunc(symcons), gen9367, gen9370)

			gen9372 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.input-track"), gen9371)

			gen9373 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.terpri-or-read-char"), Nil)

			gen9374 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9375 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9374)

			gen9376 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9377 := Call(__e, ShenFunc(symcons), V3145, gen9376)

			gen9378 := Call(__e, ShenFunc(symcons), gen9375, gen9377)

			gen9379 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.output-track"), gen9378)

			gen9380 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), Nil)

			gen9381 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen9380)

			gen9382 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen9383 := Call(__e, ShenFunc(symcons), gen9381, gen9382)

			gen9384 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen9383)

			gen9385 := Call(__e, ShenFunc(symcons), gen9384, Nil)

			gen9386 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.*call*"), gen9385)

			gen9387 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen9386)

			gen9388 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.terpri-or-read-char"), Nil)

			gen9389 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9390 := Call(__e, ShenFunc(symcons), gen9388, gen9389)

			gen9391 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9390)

			gen9392 := Call(__e, ShenFunc(symcons), gen9391, Nil)

			gen9393 := Call(__e, ShenFunc(symcons), gen9387, gen9392)

			gen9394 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9393)

			gen9395 := Call(__e, ShenFunc(symcons), gen9394, Nil)

			gen9396 := Call(__e, ShenFunc(symcons), gen9379, gen9395)

			gen9397 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9396)

			gen9398 := Call(__e, ShenFunc(symcons), gen9397, Nil)

			gen9399 := Call(__e, ShenFunc(symcons), V3147, gen9398)

			gen9400 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen9399)

			gen9401 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9400)

			gen9402 := Call(__e, ShenFunc(symcons), gen9401, Nil)

			gen9403 := Call(__e, ShenFunc(symcons), gen9373, gen9402)

			gen9404 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9403)

			gen9405 := Call(__e, ShenFunc(symcons), gen9404, Nil)

			gen9406 := Call(__e, ShenFunc(symcons), gen9372, gen9405)

			gen9407 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen9406)

			gen9408 := Call(__e, ShenFunc(symcons), gen9407, Nil)

			gen9409 := Call(__e, ShenFunc(symcons), gen9365, gen9408)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen9409)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-tracking-code"), gen9410)

		gen9413 := MakeNative(func(__e Evaluator) {
			V3153 := __e.Get(1)
			_ = V3153
			gen9412 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V3153)

			if True == gen9412 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*step*"), True)

				return
			} else {
				gen9411 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V3153)

				if True == gen9411 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*step*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("step expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("step"), gen9413)

		gen9416 := MakeNative(func(__e Evaluator) {
			V3159 := __e.Get(1)
			_ = V3159
			gen9415 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V3159)

			if True == gen9415 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*spy*"), True)

				return
			} else {
				gen9414 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V3159)

				if True == gen9414 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*spy*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("spy expects a + or a -.\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("spy"), gen9416)

		gen9420 := MakeNative(func(__e Evaluator) {
			gen9419 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*step*"))

			if True == gen9419 {
				gen9417 := Call(__e, ShenFunc(symvalue), MakeSymbol("*stinput*"))

				gen9418 := Call(__e, ShenFunc(symread_1byte), gen9417)

				__e.TailApply(ShenFunc(symshen_4check_1byte), gen9418)

				return

			} else {
				__e.TailApply(ShenFunc(symnl), MakeNumber(1))

				return
			}

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terpri-or-read-char"), gen9420)

		gen9423 := MakeNative(func(__e Evaluator) {
			V3165 := __e.Get(1)
			_ = V3165
			gen9421 := Call(__e, ShenFunc(symshen_4hat))

			gen9422 := Call(__e, ShenFunc(sym_a), V3165, gen9421)

			if True == gen9422 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("aborted"))

				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.check-byte"), gen9423)

		gen9435 := MakeNative(func(__e Evaluator) {
			V3169 := __e.Get(1)
			_ = V3169
			V3170 := __e.Get(2)
			_ = V3170
			V3171 := __e.Get(3)
			_ = V3171
			gen9424 := Call(__e, ShenFunc(symshen_4spaces), V3169)

			gen9425 := Call(__e, ShenFunc(symshen_4spaces), V3169)

			gen9426 := Call(__e, ShenFunc(symshen_4app), gen9425, MakeString(""), MakeSymbol("shen.a"))

			gen9427 := Call(__e, ShenFunc(symcn), MakeString(" \n"), gen9426)

			gen9428 := Call(__e, ShenFunc(symshen_4app), V3170, gen9427, MakeSymbol("shen.a"))

			gen9429 := Call(__e, ShenFunc(symcn), MakeString("> Inputs to "), gen9428)

			gen9430 := Call(__e, ShenFunc(symshen_4app), V3169, gen9429, MakeSymbol("shen.a"))

			gen9431 := Call(__e, ShenFunc(symcn), MakeString("<"), gen9430)

			gen9432 := Call(__e, ShenFunc(symshen_4app), gen9424, gen9431, MakeSymbol("shen.a"))

			gen9433 := Call(__e, ShenFunc(symcn), MakeString("\n"), gen9432)

			gen9434 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen9433, gen9434)

			__e.TailApply(ShenFunc(symshen_4recursively_1print), V3171)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.input-track"), gen9435)

		gen9442 := MakeNative(func(__e Evaluator) {
			V3173 := __e.Get(1)
			_ = V3173
			gen9441 := Call(__e, ShenFunc(sym_a), Nil, V3173)

			if True == gen9441 {
				gen9440 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), MakeString(" ==>"), gen9440)

				return

			} else {
				gen9439 := Call(__e, ShenFunc(symcons_2), V3173)

				if True == gen9439 {
					gen9436 := Call(__e, ShenFunc(symhd), V3173)

					Call(__e, ShenFunc(symprint), gen9436)

					gen9437 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), MakeString(", "), gen9437)

					gen9438 := Call(__e, ShenFunc(symtl), V3173)

					__e.TailApply(ShenFunc(symshen_4recursively_1print), gen9438)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.recursively-print"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.recursively-print"), gen9442)

		gen9446 := MakeNative(func(__e Evaluator) {
			V3175 := __e.Get(1)
			_ = V3175
			gen9445 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V3175)

			if True == gen9445 {
				__e.Return(MakeString(""))
				return
			} else {
				gen9443 := Call(__e, ShenFunc(sym_1), V3175, MakeNumber(1))

				gen9444 := Call(__e, ShenFunc(symshen_4spaces), gen9443)

				__e.TailApply(ShenFunc(symcn), MakeString(" "), gen9444)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.spaces"), gen9446)

		gen9460 := MakeNative(func(__e Evaluator) {
			V3179 := __e.Get(1)
			_ = V3179
			V3180 := __e.Get(2)
			_ = V3180
			V3181 := __e.Get(3)
			_ = V3181
			gen9447 := Call(__e, ShenFunc(symshen_4spaces), V3179)

			gen9448 := Call(__e, ShenFunc(symshen_4spaces), V3179)

			gen9449 := Call(__e, ShenFunc(symshen_4app), V3181, MakeString(""), MakeSymbol("shen.s"))

			gen9450 := Call(__e, ShenFunc(symcn), MakeString("==> "), gen9449)

			gen9451 := Call(__e, ShenFunc(symshen_4app), gen9448, gen9450, MakeSymbol("shen.a"))

			gen9452 := Call(__e, ShenFunc(symcn), MakeString(" \n"), gen9451)

			gen9453 := Call(__e, ShenFunc(symshen_4app), V3180, gen9452, MakeSymbol("shen.a"))

			gen9454 := Call(__e, ShenFunc(symcn), MakeString("> Output from "), gen9453)

			gen9455 := Call(__e, ShenFunc(symshen_4app), V3179, gen9454, MakeSymbol("shen.a"))

			gen9456 := Call(__e, ShenFunc(symcn), MakeString("<"), gen9455)

			gen9457 := Call(__e, ShenFunc(symshen_4app), gen9447, gen9456, MakeSymbol("shen.a"))

			gen9458 := Call(__e, ShenFunc(symcn), MakeString("\n"), gen9457)

			gen9459 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(symshen_4prhush), gen9458, gen9459)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.output-track"), gen9460)

		gen9465 := MakeNative(func(__e Evaluator) {
			V3183 := __e.Get(1)
			_ = V3183
			gen9461 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tracking*"))

			Tracking := gen9461
			gen9462 := Call(__e, ShenFunc(symremove), V3183, Tracking)

			gen9463 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), gen9462)
			_ = gen9463

			gen9464 := Call(__e, ShenFunc(symps), V3183)

			__e.TailApply(ShenFunc(symeval), gen9464)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("untrack"), gen9465)

		gen9467 := MakeNative(func(__e Evaluator) {
			V3185 := __e.Get(1)
			_ = V3185
			gen9466 := Call(__e, ShenFunc(symps), V3185)

			__e.TailApply(ShenFunc(symshen_4profile_1help), gen9466)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("profile"), gen9467)

		gen9527 := MakeNative(func(__e Evaluator) {
			V3191 := __e.Get(1)
			_ = V3191
			gen9525 := Call(__e, ShenFunc(symcons_2), V3191)

			var gen9526 Obj
			if True == gen9525 {
				gen9522 := Call(__e, ShenFunc(symhd), V3191)

				gen9523 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen9522)

				var gen9524 Obj
				if True == gen9523 {
					gen9519 := Call(__e, ShenFunc(symtl), V3191)

					gen9520 := Call(__e, ShenFunc(symcons_2), gen9519)

					var gen9521 Obj
					if True == gen9520 {
						gen9515 := Call(__e, ShenFunc(symtl), V3191)

						gen9516 := Call(__e, ShenFunc(symtl), gen9515)

						gen9517 := Call(__e, ShenFunc(symcons_2), gen9516)

						var gen9518 Obj
						if True == gen9517 {
							gen9510 := Call(__e, ShenFunc(symtl), V3191)

							gen9511 := Call(__e, ShenFunc(symtl), gen9510)

							gen9512 := Call(__e, ShenFunc(symtl), gen9511)

							gen9513 := Call(__e, ShenFunc(symcons_2), gen9512)

							var gen9514 Obj
							if True == gen9513 {
								gen9505 := Call(__e, ShenFunc(symtl), V3191)

								gen9506 := Call(__e, ShenFunc(symtl), gen9505)

								gen9507 := Call(__e, ShenFunc(symtl), gen9506)

								gen9508 := Call(__e, ShenFunc(symtl), gen9507)

								gen9509 := Call(__e, ShenFunc(sym_a), Nil, gen9508)

								if True == gen9509 {
									gen9514 = True
								} else {
									gen9514 = False
								}

							} else {
								gen9514 = False
							}
							if True == gen9514 {
								gen9518 = True
							} else {
								gen9518 = False
							}

						} else {
							gen9518 = False
						}
						if True == gen9518 {
							gen9521 = True
						} else {
							gen9521 = False
						}

					} else {
						gen9521 = False
					}
					if True == gen9521 {
						gen9524 = True
					} else {
						gen9524 = False
					}

				} else {
					gen9524 = False
				}
				if True == gen9524 {
					gen9526 = True
				} else {
					gen9526 = False
				}

			} else {
				gen9526 = False
			}
			if True == gen9526 {
				gen9468 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.f"))

				G := gen9468
				gen9469 := Call(__e, ShenFunc(symtl), V3191)

				gen9470 := Call(__e, ShenFunc(symhd), gen9469)

				gen9471 := Call(__e, ShenFunc(symtl), V3191)

				gen9472 := Call(__e, ShenFunc(symtl), gen9471)

				gen9473 := Call(__e, ShenFunc(symhd), gen9472)

				gen9474 := Call(__e, ShenFunc(symtl), V3191)

				gen9475 := Call(__e, ShenFunc(symhd), gen9474)

				gen9476 := Call(__e, ShenFunc(symtl), V3191)

				gen9477 := Call(__e, ShenFunc(symtl), gen9476)

				gen9478 := Call(__e, ShenFunc(symhd), gen9477)

				gen9479 := Call(__e, ShenFunc(symtl), V3191)

				gen9480 := Call(__e, ShenFunc(symtl), gen9479)

				gen9481 := Call(__e, ShenFunc(symhd), gen9480)

				gen9482 := Call(__e, ShenFunc(symcons), G, gen9481)

				gen9483 := Call(__e, ShenFunc(symshen_4profile_1func), gen9475, gen9478, gen9482)

				gen9484 := Call(__e, ShenFunc(symcons), gen9483, Nil)

				gen9485 := Call(__e, ShenFunc(symcons), gen9473, gen9484)

				gen9486 := Call(__e, ShenFunc(symcons), gen9470, gen9485)

				gen9487 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9486)

				Profile := gen9487
				gen9488 := Call(__e, ShenFunc(symtl), V3191)

				gen9489 := Call(__e, ShenFunc(symtl), gen9488)

				gen9490 := Call(__e, ShenFunc(symhd), gen9489)

				gen9491 := Call(__e, ShenFunc(symtl), V3191)

				gen9492 := Call(__e, ShenFunc(symhd), gen9491)

				gen9493 := Call(__e, ShenFunc(symtl), V3191)

				gen9494 := Call(__e, ShenFunc(symtl), gen9493)

				gen9495 := Call(__e, ShenFunc(symtl), gen9494)

				gen9496 := Call(__e, ShenFunc(symhd), gen9495)

				gen9497 := Call(__e, ShenFunc(symsubst), G, gen9492, gen9496)

				gen9498 := Call(__e, ShenFunc(symcons), gen9497, Nil)

				gen9499 := Call(__e, ShenFunc(symcons), gen9490, gen9498)

				gen9500 := Call(__e, ShenFunc(symcons), G, gen9499)

				gen9501 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen9500)

				Def := gen9501
				gen9502 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), Profile)

				CompileProfile := gen9502
				_ = CompileProfile
				gen9503 := Call(__e, ShenFunc(symshen_4eval_1without_1macros), Def)

				CompileG := gen9503
				_ = CompileG
				gen9504 := Call(__e, ShenFunc(symtl), V3191)

				__e.TailApply(ShenFunc(symhd), gen9504)

				return

			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("Cannot profile.\n"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.profile-help"), gen9527)

		gen9528 := MakeNative(func(__e Evaluator) {
			V3193 := __e.Get(1)
			_ = V3193
			__e.TailApply(ShenFunc(symuntrack), V3193)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unprofile"), gen9528)

		gen9559 := MakeNative(func(__e Evaluator) {
			V3197 := __e.Get(1)
			_ = V3197
			V3198 := __e.Get(2)
			_ = V3198
			V3199 := __e.Get(3)
			_ = V3199
			gen9529 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

			gen9530 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen9529)

			gen9531 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

			gen9532 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen9531)

			gen9533 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), Nil)

			gen9534 := Call(__e, ShenFunc(symcons), gen9532, gen9533)

			gen9535 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen9534)

			gen9536 := Call(__e, ShenFunc(symcons), V3197, Nil)

			gen9537 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.get-profile"), gen9536)

			gen9538 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), Nil)

			gen9539 := Call(__e, ShenFunc(symcons), gen9537, gen9538)

			gen9540 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen9539)

			gen9541 := Call(__e, ShenFunc(symcons), gen9540, Nil)

			gen9542 := Call(__e, ShenFunc(symcons), V3197, gen9541)

			gen9543 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.put-profile"), gen9542)

			gen9544 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

			gen9545 := Call(__e, ShenFunc(symcons), gen9543, gen9544)

			gen9546 := Call(__e, ShenFunc(symcons), MakeSymbol("Record"), gen9545)

			gen9547 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9546)

			gen9548 := Call(__e, ShenFunc(symcons), gen9547, Nil)

			gen9549 := Call(__e, ShenFunc(symcons), gen9535, gen9548)

			gen9550 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen9549)

			gen9551 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9550)

			gen9552 := Call(__e, ShenFunc(symcons), gen9551, Nil)

			gen9553 := Call(__e, ShenFunc(symcons), V3199, gen9552)

			gen9554 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen9553)

			gen9555 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen9554)

			gen9556 := Call(__e, ShenFunc(symcons), gen9555, Nil)

			gen9557 := Call(__e, ShenFunc(symcons), gen9530, gen9556)

			gen9558 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), gen9557)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen9558)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.profile-func"), gen9559)

		gen9562 := MakeNative(func(__e Evaluator) {
			V3201 := __e.Get(1)
			_ = V3201
			gen9560 := Call(__e, ShenFunc(symshen_4get_1profile), V3201)

			Results := gen9560
			gen9561 := Call(__e, ShenFunc(symshen_4put_1profile), V3201, MakeNumber(0))

			Initialise := gen9561
			_ = Initialise
			__e.TailApply(ShenFunc(sym_8p), V3201, Results)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("profile-results"), gen9562)

		gen9566 := MakeNative(func(__e Evaluator) {
			V3203 := __e.Get(1)
			_ = V3203
			gen9563 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeNumber(0))
				return
			}, 1)
			gen9565 := MakeNative(func(__e Evaluator) {
				gen9564 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symget), V3203, MakeSymbol("profile"), gen9564)

				return

			}, 0)
			__e.Return(Try(__e, gen9565).Catch(gen9563))
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.get-profile"), gen9566)

		gen9568 := MakeNative(func(__e Evaluator) {
			V3206 := __e.Get(1)
			_ = V3206
			V3207 := __e.Get(2)
			_ = V3207
			gen9567 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			__e.TailApply(ShenFunc(symput), V3206, MakeSymbol("profile"), V3207, gen9567)

			return

		}, 2)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.put-profile"), gen9568)

		return

	}, 0))
}
