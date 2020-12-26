package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen16215 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2731 := __args[0]
			_ = V2731
			V2732 := __args[1]
			_ = V2732
			gen16206 := Call(__e, ShenFunc(symshen_4curry), V2731)

			Curry := gen16206
			gen16207 := Call(__e, ShenFunc(symshen_4start_1new_1prolog_1process))

			ProcessN := gen16207
			gen16208 := Call(__e, ShenFunc(symshen_4curry_1type), V2732)

			gen16209 := Call(__e, ShenFunc(symshen_4demodulate), gen16208)

			gen16210 := Call(__e, ShenFunc(symshen_4insert_1prolog_1variables), gen16209, ProcessN)

			Type := gen16210
			gen16211 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symreturn), Type, ProcessN, MakeSymbol("shen.void"))

				return
			}, 0)
			Continuation := gen16211
			gen16212 := Call(__e, ShenFunc(symcons), Type, Nil)

			gen16213 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16212)

			gen16214 := Call(__e, ShenFunc(symcons), Curry, gen16213)

			__e.TailApply(ShenFunc(symshen_4t_d), gen16214, Nil, ProcessN, Continuation)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typecheck"), gen16215)

		gen16283 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2734 := __args[0]
			_ = V2734
			gen16281 := Call(__e, ShenFunc(symcons_2), V2734)

			var gen16282 Obj
			if True == gen16281 {
				gen16279 := Call(__e, ShenFunc(symhd), V2734)

				gen16280 := Call(__e, ShenFunc(symshen_4special_2), gen16279)

				if True == gen16280 {
					gen16282 = True
				} else {
					gen16282 = False
				}

			} else {
				gen16282 = False
			}
			if True == gen16282 {
				gen16275 := Call(__e, ShenFunc(symhd), V2734)

				gen16276 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Y := __args[0]
					_ = Y
					__e.TailApply(ShenFunc(symshen_4curry), Y)

					return
				}, 1)
				gen16277 := Call(__e, ShenFunc(symtl), V2734)

				gen16278 := Call(__e, ShenFunc(symmap), gen16276, gen16277)

				__e.TailApply(ShenFunc(symcons), gen16275, gen16278)

				return

			} else {
				gen16273 := Call(__e, ShenFunc(symcons_2), V2734)

				var gen16274 Obj
				if True == gen16273 {
					gen16270 := Call(__e, ShenFunc(symtl), V2734)

					gen16271 := Call(__e, ShenFunc(symcons_2), gen16270)

					var gen16272 Obj
					if True == gen16271 {
						gen16268 := Call(__e, ShenFunc(symhd), V2734)

						gen16269 := Call(__e, ShenFunc(symshen_4extraspecial_2), gen16268)

						if True == gen16269 {
							gen16272 = True
						} else {
							gen16272 = False
						}

					} else {
						gen16272 = False
					}
					if True == gen16272 {
						gen16274 = True
					} else {
						gen16274 = False
					}

				} else {
					gen16274 = False
				}
				if True == gen16274 {
					__e.Return(V2734)
					return
				} else {
					gen16266 := Call(__e, ShenFunc(symcons_2), V2734)

					var gen16267 Obj
					if True == gen16266 {
						gen16263 := Call(__e, ShenFunc(symhd), V2734)

						gen16264 := Call(__e, ShenFunc(sym_a), MakeSymbol("type"), gen16263)

						var gen16265 Obj
						if True == gen16264 {
							gen16260 := Call(__e, ShenFunc(symtl), V2734)

							gen16261 := Call(__e, ShenFunc(symcons_2), gen16260)

							var gen16262 Obj
							if True == gen16261 {
								gen16256 := Call(__e, ShenFunc(symtl), V2734)

								gen16257 := Call(__e, ShenFunc(symtl), gen16256)

								gen16258 := Call(__e, ShenFunc(symcons_2), gen16257)

								var gen16259 Obj
								if True == gen16258 {
									gen16252 := Call(__e, ShenFunc(symtl), V2734)

									gen16253 := Call(__e, ShenFunc(symtl), gen16252)

									gen16254 := Call(__e, ShenFunc(symtl), gen16253)

									gen16255 := Call(__e, ShenFunc(sym_a), Nil, gen16254)

									if True == gen16255 {
										gen16259 = True
									} else {
										gen16259 = False
									}

								} else {
									gen16259 = False
								}
								if True == gen16259 {
									gen16262 = True
								} else {
									gen16262 = False
								}

							} else {
								gen16262 = False
							}
							if True == gen16262 {
								gen16265 = True
							} else {
								gen16265 = False
							}

						} else {
							gen16265 = False
						}
						if True == gen16265 {
							gen16267 = True
						} else {
							gen16267 = False
						}

					} else {
						gen16267 = False
					}
					if True == gen16267 {
						gen16246 := Call(__e, ShenFunc(symtl), V2734)

						gen16247 := Call(__e, ShenFunc(symhd), gen16246)

						gen16248 := Call(__e, ShenFunc(symshen_4curry), gen16247)

						gen16249 := Call(__e, ShenFunc(symtl), V2734)

						gen16250 := Call(__e, ShenFunc(symtl), gen16249)

						gen16251 := Call(__e, ShenFunc(symcons), gen16248, gen16250)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen16251)

						return

					} else {
						gen16244 := Call(__e, ShenFunc(symcons_2), V2734)

						var gen16245 Obj
						if True == gen16244 {
							gen16241 := Call(__e, ShenFunc(symtl), V2734)

							gen16242 := Call(__e, ShenFunc(symcons_2), gen16241)

							var gen16243 Obj
							if True == gen16242 {
								gen16238 := Call(__e, ShenFunc(symtl), V2734)

								gen16239 := Call(__e, ShenFunc(symtl), gen16238)

								gen16240 := Call(__e, ShenFunc(symcons_2), gen16239)

								if True == gen16240 {
									gen16243 = True
								} else {
									gen16243 = False
								}

							} else {
								gen16243 = False
							}
							if True == gen16243 {
								gen16245 = True
							} else {
								gen16245 = False
							}

						} else {
							gen16245 = False
						}
						if True == gen16245 {
							gen16230 := Call(__e, ShenFunc(symhd), V2734)

							gen16231 := Call(__e, ShenFunc(symtl), V2734)

							gen16232 := Call(__e, ShenFunc(symhd), gen16231)

							gen16233 := Call(__e, ShenFunc(symcons), gen16232, Nil)

							gen16234 := Call(__e, ShenFunc(symcons), gen16230, gen16233)

							gen16235 := Call(__e, ShenFunc(symtl), V2734)

							gen16236 := Call(__e, ShenFunc(symtl), gen16235)

							gen16237 := Call(__e, ShenFunc(symcons), gen16234, gen16236)

							__e.TailApply(ShenFunc(symshen_4curry), gen16237)

							return

						} else {
							gen16228 := Call(__e, ShenFunc(symcons_2), V2734)

							var gen16229 Obj
							if True == gen16228 {
								gen16225 := Call(__e, ShenFunc(symtl), V2734)

								gen16226 := Call(__e, ShenFunc(symcons_2), gen16225)

								var gen16227 Obj
								if True == gen16226 {
									gen16222 := Call(__e, ShenFunc(symtl), V2734)

									gen16223 := Call(__e, ShenFunc(symtl), gen16222)

									gen16224 := Call(__e, ShenFunc(sym_a), Nil, gen16223)

									if True == gen16224 {
										gen16227 = True
									} else {
										gen16227 = False
									}

								} else {
									gen16227 = False
								}
								if True == gen16227 {
									gen16229 = True
								} else {
									gen16229 = False
								}

							} else {
								gen16229 = False
							}
							if True == gen16229 {
								gen16216 := Call(__e, ShenFunc(symhd), V2734)

								gen16217 := Call(__e, ShenFunc(symshen_4curry), gen16216)

								gen16218 := Call(__e, ShenFunc(symtl), V2734)

								gen16219 := Call(__e, ShenFunc(symhd), gen16218)

								gen16220 := Call(__e, ShenFunc(symshen_4curry), gen16219)

								gen16221 := Call(__e, ShenFunc(symcons), gen16220, Nil)

								__e.TailApply(ShenFunc(symcons), gen16217, gen16221)

								return

							} else {
								__e.Return(V2734)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry"), gen16283)

		gen16285 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2736 := __args[0]
			_ = V2736
			gen16284 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*special*"))

			__e.TailApply(ShenFunc(symelement_2), V2736, gen16284)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.special?"), gen16285)

		gen16287 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2738 := __args[0]
			_ = V2738
			gen16286 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*extraspecial*"))

			__e.TailApply(ShenFunc(symelement_2), V2738, gen16286)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extraspecial?"), gen16287)

		gen16326 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2743 := __args[0]
			_ = V2743
			V2744 := __args[1]
			_ = V2744
			V2745 := __args[2]
			_ = V2745
			V2746 := __args[3]
			_ = V2746
			gen16288 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen16288
			gen16289 := Call(__e, ShenFunc(symshen_4newpv), V2745)

			Error := gen16289
			Call(__e, ShenFunc(symshen_4incinfs))
			gen16290 := Call(__e, ShenFunc(symshen_4maxinfexceeded_2))

			gen16292 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen16291 := Call(__e, ShenFunc(symshen_4errormaxinfs))

				__e.TailApply(ShenFunc(symbind), Error, gen16291, V2745, V2746)

				return

			}, 0)
			gen16293 := Call(__e, ShenFunc(symfwhen), gen16290, V2745, gen16292)

			Case := gen16293
			gen16324 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen16325 Obj
			if True == gen16324 {
				gen16294 := Call(__e, ShenFunc(symshen_4lazyderef), V2743, V2745)

				V2723 := gen16294
				gen16296 := Call(__e, ShenFunc(sym_a), MakeSymbol("fail"), V2723)

				var gen16297 Obj
				if True == gen16296 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen16295 := MakeNative(func(__e Evaluator, __args ...Obj) {
						__e.TailApply(ShenFunc(symshen_4prolog_1failure), V2745, V2746)

						return
					}, 0)
					gen16297 = Call(__e, ShenFunc(symcut), Throwcontrol, V2745, gen16295)

				} else {
					gen16297 = False
				}
				Case := gen16297
				gen16323 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen16323 {
					gen16298 := Call(__e, ShenFunc(symshen_4lazyderef), V2743, V2745)

					V2724 := gen16298
					gen16316 := Call(__e, ShenFunc(symcons_2), V2724)

					var gen16317 Obj
					if True == gen16316 {
						gen16299 := Call(__e, ShenFunc(symhd), V2724)

						X := gen16299
						gen16300 := Call(__e, ShenFunc(symtl), V2724)

						gen16301 := Call(__e, ShenFunc(symshen_4lazyderef), gen16300, V2745)

						V2725 := gen16301
						gen16315 := Call(__e, ShenFunc(symcons_2), V2725)

						if True == gen16315 {
							gen16302 := Call(__e, ShenFunc(symhd), V2725)

							gen16303 := Call(__e, ShenFunc(symshen_4lazyderef), gen16302, V2745)

							V2726 := gen16303
							gen16314 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2726)

							if True == gen16314 {
								gen16304 := Call(__e, ShenFunc(symtl), V2725)

								gen16305 := Call(__e, ShenFunc(symshen_4lazyderef), gen16304, V2745)

								V2727 := gen16305
								gen16313 := Call(__e, ShenFunc(symcons_2), V2727)

								if True == gen16313 {
									gen16306 := Call(__e, ShenFunc(symhd), V2727)

									A := gen16306
									gen16307 := Call(__e, ShenFunc(symtl), V2727)

									gen16308 := Call(__e, ShenFunc(symshen_4lazyderef), gen16307, V2745)

									V2728 := gen16308
									gen16312 := Call(__e, ShenFunc(sym_a), Nil, V2728)

									if True == gen16312 {
										Call(__e, ShenFunc(symshen_4incinfs))
										gen16309 := Call(__e, ShenFunc(symshen_4type_1theory_1enabled_2))

										gen16311 := MakeNative(func(__e Evaluator, __args ...Obj) {
											gen16310 := MakeNative(func(__e Evaluator, __args ...Obj) {
												__e.TailApply(ShenFunc(symshen_4th_d), X, A, V2744, V2745, V2746)

												return
											}, 0)
											__e.TailApply(ShenFunc(symcut), Throwcontrol, V2745, gen16310)

											return

										}, 0)
										gen16317 = Call(__e, ShenFunc(symfwhen), gen16309, V2745, gen16311)

									} else {
										gen16317 = False
									}

								} else {
									gen16317 = False
								}

							} else {
								gen16317 = False
							}

						} else {
							gen16317 = False
						}

					} else {
						gen16317 = False
					}
					Case := gen16317
					gen16322 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen16322 {
						gen16318 := Call(__e, ShenFunc(symshen_4newpv), V2745)

						Datatypes := gen16318
						Call(__e, ShenFunc(symshen_4incinfs))
						gen16321 := MakeNative(func(__e Evaluator, __args ...Obj) {
							gen16319 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

							gen16320 := MakeNative(func(__e Evaluator, __args ...Obj) {
								__e.TailApply(ShenFunc(symshen_4udefs_d), V2743, V2744, Datatypes, V2745, V2746)

								return
							}, 0)
							__e.TailApply(ShenFunc(symbind), Datatypes, gen16319, V2745, gen16320)

							return

						}, 0)
						gen16325 = Call(__e, ShenFunc(symshen_4show), V2743, V2744, V2745, gen16321)

					} else {
						gen16325 = Case
					}

				} else {
					gen16325 = Case
				}

			} else {
				gen16325 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen16325)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*"), gen16326)

		gen16327 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*shen-type-theory-enabled?*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-theory-enabled?"), gen16327)

		gen16330 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2752 := __args[0]
			_ = V2752
			gen16329 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V2752)

			if True == gen16329 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), True)

				return
			} else {
				gen16328 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V2752)

				if True == gen16328 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("enable-type-theory expects a + or a -\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("enable-type-theory"), gen16330)

		gen16331 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2763 := __args[0]
			_ = V2763
			V2764 := __args[1]
			_ = V2764
			__e.Return(False)
			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-failure"), gen16331)

		gen16334 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen16332 := Call(__e, ShenFunc(syminferences))

			gen16333 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*maxinferences*"))

			__e.TailApply(ShenFunc(sym_6), gen16332, gen16333)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.maxinfexceeded?"), gen16334)

		gen16335 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString("maximum inferences exceeded~%"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.errormaxinfs"), gen16335)

		gen16347 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2770 := __args[0]
			_ = V2770
			V2771 := __args[1]
			_ = V2771
			V2772 := __args[2]
			_ = V2772
			V2773 := __args[3]
			_ = V2773
			V2774 := __args[4]
			_ = V2774
			gen16336 := Call(__e, ShenFunc(symshen_4lazyderef), V2772, V2773)

			V2719 := gen16336
			gen16341 := Call(__e, ShenFunc(symcons_2), V2719)

			var gen16342 Obj
			if True == gen16341 {
				gen16337 := Call(__e, ShenFunc(symhd), V2719)

				D := gen16337
				Call(__e, ShenFunc(symshen_4incinfs))
				gen16338 := Call(__e, ShenFunc(symcons), V2771, Nil)

				gen16339 := Call(__e, ShenFunc(symcons), V2770, gen16338)

				gen16340 := Call(__e, ShenFunc(symcons), D, gen16339)

				gen16342 = Call(__e, ShenFunc(symcall), gen16340, V2773, V2774)

			} else {
				gen16342 = False
			}
			Case := gen16342
			gen16346 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen16346 {
				gen16343 := Call(__e, ShenFunc(symshen_4lazyderef), V2772, V2773)

				V2720 := gen16343
				gen16345 := Call(__e, ShenFunc(symcons_2), V2720)

				if True == gen16345 {
					gen16344 := Call(__e, ShenFunc(symtl), V2720)

					Ds := gen16344
					Call(__e, ShenFunc(symshen_4incinfs))
					__e.TailApply(ShenFunc(symshen_4udefs_d), V2770, V2771, Ds, V2773, V2774)

					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.udefs*"), gen16347)

		gen17091 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2780 := __args[0]
			_ = V2780
			V2781 := __args[1]
			_ = V2781
			V2782 := __args[2]
			_ = V2782
			V2783 := __args[3]
			_ = V2783
			V2784 := __args[4]
			_ = V2784
			gen16348 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen16348
			Call(__e, ShenFunc(symshen_4incinfs))
			gen16349 := Call(__e, ShenFunc(symcons), V2781, Nil)

			gen16350 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16349)

			gen16351 := Call(__e, ShenFunc(symcons), V2780, gen16350)

			gen16352 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symfwhen), False, V2783, V2784)

				return
			}, 0)
			gen16353 := Call(__e, ShenFunc(symshen_4show), gen16351, V2782, V2783, gen16352)

			Case := gen16353
			gen17089 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen17090 Obj
			if True == gen17089 {
				gen16354 := Call(__e, ShenFunc(symshen_4newpv), V2783)

				F := gen16354
				Call(__e, ShenFunc(symshen_4incinfs))
				gen16355 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

				gen16356 := Call(__e, ShenFunc(symshen_4typedf_2), gen16355)

				gen16362 := MakeNative(func(__e Evaluator, __args ...Obj) {
					gen16357 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

					gen16358 := Call(__e, ShenFunc(symshen_4sigf), gen16357)

					gen16361 := MakeNative(func(__e Evaluator, __args ...Obj) {
						gen16359 := Call(__e, ShenFunc(symcons), V2781, Nil)

						gen16360 := Call(__e, ShenFunc(symcons), F, gen16359)

						__e.TailApply(ShenFunc(symcall), gen16360, V2783, V2784)

						return

					}, 0)
					__e.TailApply(ShenFunc(symbind), F, gen16358, V2783, gen16361)

					return

				}, 0)
				gen16363 := Call(__e, ShenFunc(symfwhen), gen16356, V2783, gen16362)

				Case := gen16363
				gen17088 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen17088 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen16364 := Call(__e, ShenFunc(symshen_4base), V2780, V2781, V2783, V2784)

					Case := gen16364
					gen17087 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen17087 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen16365 := Call(__e, ShenFunc(symshen_4by__hypothesis), V2780, V2781, V2782, V2783, V2784)

						Case := gen16365
						gen17086 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen17086 {
							gen16366 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

							V2615 := gen16366
							gen16373 := Call(__e, ShenFunc(symcons_2), V2615)

							var gen16374 Obj
							if True == gen16373 {
								gen16367 := Call(__e, ShenFunc(symhd), V2615)

								F := gen16367
								gen16368 := Call(__e, ShenFunc(symtl), V2615)

								gen16369 := Call(__e, ShenFunc(symshen_4lazyderef), gen16368, V2783)

								V2616 := gen16369
								gen16372 := Call(__e, ShenFunc(sym_a), Nil, V2616)

								if True == gen16372 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen16370 := Call(__e, ShenFunc(symcons), V2781, Nil)

									gen16371 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16370)

									gen16374 = Call(__e, ShenFunc(symshen_4th_d), F, gen16371, V2782, V2783, V2784)

								} else {
									gen16374 = False
								}

							} else {
								gen16374 = False
							}
							Case := gen16374
							gen17085 := Call(__e, ShenFunc(sym_a), Case, False)

							if True == gen17085 {
								gen16375 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

								V2617 := gen16375
								gen16389 := Call(__e, ShenFunc(symcons_2), V2617)

								var gen16390 Obj
								if True == gen16389 {
									gen16376 := Call(__e, ShenFunc(symhd), V2617)

									F := gen16376
									gen16377 := Call(__e, ShenFunc(symtl), V2617)

									gen16378 := Call(__e, ShenFunc(symshen_4lazyderef), gen16377, V2783)

									V2618 := gen16378
									gen16388 := Call(__e, ShenFunc(symcons_2), V2618)

									if True == gen16388 {
										gen16379 := Call(__e, ShenFunc(symhd), V2618)

										X := gen16379
										gen16380 := Call(__e, ShenFunc(symtl), V2618)

										gen16381 := Call(__e, ShenFunc(symshen_4lazyderef), gen16380, V2783)

										V2619 := gen16381
										gen16387 := Call(__e, ShenFunc(sym_a), Nil, V2619)

										if True == gen16387 {
											gen16382 := Call(__e, ShenFunc(symshen_4newpv), V2783)

											B := gen16382
											Call(__e, ShenFunc(symshen_4incinfs))
											gen16383 := Call(__e, ShenFunc(symcons), V2781, Nil)

											gen16384 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16383)

											gen16385 := Call(__e, ShenFunc(symcons), B, gen16384)

											gen16386 := MakeNative(func(__e Evaluator, __args ...Obj) {
												__e.TailApply(ShenFunc(symshen_4th_d), X, B, V2782, V2783, V2784)

												return
											}, 0)
											gen16390 = Call(__e, ShenFunc(symshen_4th_d), F, gen16385, V2782, V2783, gen16386)

										} else {
											gen16390 = False
										}

									} else {
										gen16390 = False
									}

								} else {
									gen16390 = False
								}
								Case := gen16390
								gen17084 := Call(__e, ShenFunc(sym_a), Case, False)

								if True == gen17084 {
									gen16391 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

									V2620 := gen16391
									gen16465 := Call(__e, ShenFunc(symcons_2), V2620)

									var gen16466 Obj
									if True == gen16465 {
										gen16392 := Call(__e, ShenFunc(symhd), V2620)

										gen16393 := Call(__e, ShenFunc(symshen_4lazyderef), gen16392, V2783)

										V2621 := gen16393
										gen16464 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), V2621)

										if True == gen16464 {
											gen16394 := Call(__e, ShenFunc(symtl), V2620)

											gen16395 := Call(__e, ShenFunc(symshen_4lazyderef), gen16394, V2783)

											V2622 := gen16395
											gen16463 := Call(__e, ShenFunc(symcons_2), V2622)

											if True == gen16463 {
												gen16396 := Call(__e, ShenFunc(symhd), V2622)

												X := gen16396
												gen16397 := Call(__e, ShenFunc(symtl), V2622)

												gen16398 := Call(__e, ShenFunc(symshen_4lazyderef), gen16397, V2783)

												V2623 := gen16398
												gen16462 := Call(__e, ShenFunc(symcons_2), V2623)

												if True == gen16462 {
													gen16399 := Call(__e, ShenFunc(symhd), V2623)

													Y := gen16399
													gen16400 := Call(__e, ShenFunc(symtl), V2623)

													gen16401 := Call(__e, ShenFunc(symshen_4lazyderef), gen16400, V2783)

													V2624 := gen16401
													gen16461 := Call(__e, ShenFunc(sym_a), Nil, V2624)

													if True == gen16461 {
														gen16402 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

														V2625 := gen16402
														gen16460 := Call(__e, ShenFunc(symcons_2), V2625)

														if True == gen16460 {
															gen16411 := Call(__e, ShenFunc(symhd), V2625)

															gen16412 := Call(__e, ShenFunc(symshen_4lazyderef), gen16411, V2783)

															V2626 := gen16412
															gen16459 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2626)

															if True == gen16459 {
																gen16437 := Call(__e, ShenFunc(symtl), V2625)

																gen16438 := Call(__e, ShenFunc(symshen_4lazyderef), gen16437, V2783)

																V2627 := gen16438
																gen16458 := Call(__e, ShenFunc(symcons_2), V2627)

																if True == gen16458 {
																	gen16446 := Call(__e, ShenFunc(symhd), V2627)

																	A := gen16446
																	gen16447 := Call(__e, ShenFunc(symtl), V2627)

																	gen16448 := Call(__e, ShenFunc(symshen_4lazyderef), gen16447, V2783)

																	V2628 := gen16448
																	gen16457 := Call(__e, ShenFunc(sym_a), Nil, V2628)

																	if True == gen16457 {
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen16456 := MakeNative(func(__e Evaluator, __args ...Obj) {
																			gen16454 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen16455 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16454)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16455, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen16466 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16456)

																	} else {
																		gen16453 := Call(__e, ShenFunc(symshen_4pvar_2), V2628)

																		if True == gen16453 {
																			Call(__e, ShenFunc(symshen_4bindv), V2628, Nil, V2783)
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen16451 := MakeNative(func(__e Evaluator, __args ...Obj) {
																				gen16449 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen16450 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16449)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16450, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen16452 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16451)

																			Result := gen16452
																			Call(__e, ShenFunc(symshen_4unbindv), V2628, V2783)
																			gen16466 = Result

																		} else {
																			gen16466 = False
																		}

																	}

																} else {
																	gen16445 := Call(__e, ShenFunc(symshen_4pvar_2), V2627)

																	if True == gen16445 {
																		gen16439 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		A := gen16439
																		gen16440 := Call(__e, ShenFunc(symcons), A, Nil)

																		Call(__e, ShenFunc(symshen_4bindv), V2627, gen16440, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen16443 := MakeNative(func(__e Evaluator, __args ...Obj) {
																			gen16441 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen16442 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16441)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16442, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen16444 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16443)

																		Result := gen16444
																		Call(__e, ShenFunc(symshen_4unbindv), V2627, V2783)
																		gen16466 = Result

																	} else {
																		gen16466 = False
																	}

																}

															} else {
																gen16436 := Call(__e, ShenFunc(symshen_4pvar_2), V2626)

																if True == gen16436 {
																	Call(__e, ShenFunc(symshen_4bindv), V2626, MakeSymbol("list"), V2783)
																	gen16413 := Call(__e, ShenFunc(symtl), V2625)

																	gen16414 := Call(__e, ShenFunc(symshen_4lazyderef), gen16413, V2783)

																	V2629 := gen16414
																	gen16434 := Call(__e, ShenFunc(symcons_2), V2629)

																	var gen16435 Obj
																	if True == gen16434 {
																		gen16422 := Call(__e, ShenFunc(symhd), V2629)

																		A := gen16422
																		gen16423 := Call(__e, ShenFunc(symtl), V2629)

																		gen16424 := Call(__e, ShenFunc(symshen_4lazyderef), gen16423, V2783)

																		V2630 := gen16424
																		gen16433 := Call(__e, ShenFunc(sym_a), Nil, V2630)

																		if True == gen16433 {
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen16432 := MakeNative(func(__e Evaluator, __args ...Obj) {
																				gen16430 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen16431 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16430)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16431, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen16435 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16432)

																		} else {
																			gen16429 := Call(__e, ShenFunc(symshen_4pvar_2), V2630)

																			if True == gen16429 {
																				Call(__e, ShenFunc(symshen_4bindv), V2630, Nil, V2783)
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16427 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					gen16425 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen16426 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16425)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16426, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen16428 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16427)

																				Result := gen16428
																				Call(__e, ShenFunc(symshen_4unbindv), V2630, V2783)
																				gen16435 = Result

																			} else {
																				gen16435 = False
																			}

																		}

																	} else {
																		gen16421 := Call(__e, ShenFunc(symshen_4pvar_2), V2629)

																		if True == gen16421 {
																			gen16415 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																			A := gen16415
																			gen16416 := Call(__e, ShenFunc(symcons), A, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2629, gen16416, V2783)

																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen16419 := MakeNative(func(__e Evaluator, __args ...Obj) {
																				gen16417 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen16418 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16417)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16418, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen16420 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16419)

																			Result := gen16420
																			Call(__e, ShenFunc(symshen_4unbindv), V2629, V2783)
																			gen16435 = Result

																		} else {
																			gen16435 = False
																		}

																	}
																	Result := gen16435
																	Call(__e, ShenFunc(symshen_4unbindv), V2626, V2783)
																	gen16466 = Result

																} else {
																	gen16466 = False
																}

															}

														} else {
															gen16410 := Call(__e, ShenFunc(symshen_4pvar_2), V2625)

															if True == gen16410 {
																gen16403 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																A := gen16403
																gen16404 := Call(__e, ShenFunc(symcons), A, Nil)

																gen16405 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16404)

																Call(__e, ShenFunc(symshen_4bindv), V2625, gen16405, V2783)

																Call(__e, ShenFunc(symshen_4incinfs))
																gen16408 := MakeNative(func(__e Evaluator, __args ...Obj) {
																	gen16406 := Call(__e, ShenFunc(symcons), A, Nil)

																	gen16407 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16406)

																	__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16407, V2782, V2783, V2784)

																	return

																}, 0)
																gen16409 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16408)

																Result := gen16409
																Call(__e, ShenFunc(symshen_4unbindv), V2625, V2783)
																gen16466 = Result

															} else {
																gen16466 = False
															}

														}

													} else {
														gen16466 = False
													}

												} else {
													gen16466 = False
												}

											} else {
												gen16466 = False
											}

										} else {
											gen16466 = False
										}

									} else {
										gen16466 = False
									}
									Case := gen16466
									gen17083 := Call(__e, ShenFunc(sym_a), Case, False)

									if True == gen17083 {
										gen16467 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

										V2631 := gen16467
										gen16539 := Call(__e, ShenFunc(symcons_2), V2631)

										var gen16540 Obj
										if True == gen16539 {
											gen16468 := Call(__e, ShenFunc(symhd), V2631)

											gen16469 := Call(__e, ShenFunc(symshen_4lazyderef), gen16468, V2783)

											V2632 := gen16469
											gen16538 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), V2632)

											if True == gen16538 {
												gen16470 := Call(__e, ShenFunc(symtl), V2631)

												gen16471 := Call(__e, ShenFunc(symshen_4lazyderef), gen16470, V2783)

												V2633 := gen16471
												gen16537 := Call(__e, ShenFunc(symcons_2), V2633)

												if True == gen16537 {
													gen16472 := Call(__e, ShenFunc(symhd), V2633)

													X := gen16472
													gen16473 := Call(__e, ShenFunc(symtl), V2633)

													gen16474 := Call(__e, ShenFunc(symshen_4lazyderef), gen16473, V2783)

													V2634 := gen16474
													gen16536 := Call(__e, ShenFunc(symcons_2), V2634)

													if True == gen16536 {
														gen16475 := Call(__e, ShenFunc(symhd), V2634)

														Y := gen16475
														gen16476 := Call(__e, ShenFunc(symtl), V2634)

														gen16477 := Call(__e, ShenFunc(symshen_4lazyderef), gen16476, V2783)

														V2635 := gen16477
														gen16535 := Call(__e, ShenFunc(sym_a), Nil, V2635)

														if True == gen16535 {
															gen16478 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

															V2636 := gen16478
															gen16534 := Call(__e, ShenFunc(symcons_2), V2636)

															if True == gen16534 {
																gen16487 := Call(__e, ShenFunc(symhd), V2636)

																A := gen16487
																gen16488 := Call(__e, ShenFunc(symtl), V2636)

																gen16489 := Call(__e, ShenFunc(symshen_4lazyderef), gen16488, V2783)

																V2637 := gen16489
																gen16533 := Call(__e, ShenFunc(symcons_2), V2637)

																if True == gen16533 {
																	gen16496 := Call(__e, ShenFunc(symhd), V2637)

																	gen16497 := Call(__e, ShenFunc(symshen_4lazyderef), gen16496, V2783)

																	V2638 := gen16497
																	gen16532 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), V2638)

																	if True == gen16532 {
																		gen16516 := Call(__e, ShenFunc(symtl), V2637)

																		gen16517 := Call(__e, ShenFunc(symshen_4lazyderef), gen16516, V2783)

																		V2639 := gen16517
																		gen16531 := Call(__e, ShenFunc(symcons_2), V2639)

																		if True == gen16531 {
																			gen16523 := Call(__e, ShenFunc(symhd), V2639)

																			B := gen16523
																			gen16524 := Call(__e, ShenFunc(symtl), V2639)

																			gen16525 := Call(__e, ShenFunc(symshen_4lazyderef), gen16524, V2783)

																			V2640 := gen16525
																			gen16530 := Call(__e, ShenFunc(sym_a), Nil, V2640)

																			if True == gen16530 {
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16529 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																					return
																				}, 0)
																				gen16540 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16529)

																			} else {
																				gen16528 := Call(__e, ShenFunc(symshen_4pvar_2), V2640)

																				if True == gen16528 {
																					Call(__e, ShenFunc(symshen_4bindv), V2640, Nil, V2783)
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16526 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen16527 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16526)

																					Result := gen16527
																					Call(__e, ShenFunc(symshen_4unbindv), V2640, V2783)
																					gen16540 = Result

																				} else {
																					gen16540 = False
																				}

																			}

																		} else {
																			gen16522 := Call(__e, ShenFunc(symshen_4pvar_2), V2639)

																			if True == gen16522 {
																				gen16518 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen16518
																				gen16519 := Call(__e, ShenFunc(symcons), B, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2639, gen16519, V2783)

																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16520 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																					return
																				}, 0)
																				gen16521 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16520)

																				Result := gen16521
																				Call(__e, ShenFunc(symshen_4unbindv), V2639, V2783)
																				gen16540 = Result

																			} else {
																				gen16540 = False
																			}

																		}

																	} else {
																		gen16515 := Call(__e, ShenFunc(symshen_4pvar_2), V2638)

																		if True == gen16515 {
																			Call(__e, ShenFunc(symshen_4bindv), V2638, MakeSymbol("*"), V2783)
																			gen16498 := Call(__e, ShenFunc(symtl), V2637)

																			gen16499 := Call(__e, ShenFunc(symshen_4lazyderef), gen16498, V2783)

																			V2641 := gen16499
																			gen16513 := Call(__e, ShenFunc(symcons_2), V2641)

																			var gen16514 Obj
																			if True == gen16513 {
																				gen16505 := Call(__e, ShenFunc(symhd), V2641)

																				B := gen16505
																				gen16506 := Call(__e, ShenFunc(symtl), V2641)

																				gen16507 := Call(__e, ShenFunc(symshen_4lazyderef), gen16506, V2783)

																				V2642 := gen16507
																				gen16512 := Call(__e, ShenFunc(sym_a), Nil, V2642)

																				if True == gen16512 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16511 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen16514 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16511)

																				} else {
																					gen16510 := Call(__e, ShenFunc(symshen_4pvar_2), V2642)

																					if True == gen16510 {
																						Call(__e, ShenFunc(symshen_4bindv), V2642, Nil, V2783)
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen16508 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																							return
																						}, 0)
																						gen16509 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16508)

																						Result := gen16509
																						Call(__e, ShenFunc(symshen_4unbindv), V2642, V2783)
																						gen16514 = Result

																					} else {
																						gen16514 = False
																					}

																				}

																			} else {
																				gen16504 := Call(__e, ShenFunc(symshen_4pvar_2), V2641)

																				if True == gen16504 {
																					gen16500 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					B := gen16500
																					gen16501 := Call(__e, ShenFunc(symcons), B, Nil)

																					Call(__e, ShenFunc(symshen_4bindv), V2641, gen16501, V2783)

																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16502 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen16503 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16502)

																					Result := gen16503
																					Call(__e, ShenFunc(symshen_4unbindv), V2641, V2783)
																					gen16514 = Result

																				} else {
																					gen16514 = False
																				}

																			}
																			Result := gen16514
																			Call(__e, ShenFunc(symshen_4unbindv), V2638, V2783)
																			gen16540 = Result

																		} else {
																			gen16540 = False
																		}

																	}

																} else {
																	gen16495 := Call(__e, ShenFunc(symshen_4pvar_2), V2637)

																	if True == gen16495 {
																		gen16490 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		B := gen16490
																		gen16491 := Call(__e, ShenFunc(symcons), B, Nil)

																		gen16492 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16491)

																		Call(__e, ShenFunc(symshen_4bindv), V2637, gen16492, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen16493 := MakeNative(func(__e Evaluator, __args ...Obj) {
																			__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																			return
																		}, 0)
																		gen16494 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16493)

																		Result := gen16494
																		Call(__e, ShenFunc(symshen_4unbindv), V2637, V2783)
																		gen16540 = Result

																	} else {
																		gen16540 = False
																	}

																}

															} else {
																gen16486 := Call(__e, ShenFunc(symshen_4pvar_2), V2636)

																if True == gen16486 {
																	gen16479 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																	A := gen16479
																	gen16480 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																	B := gen16480
																	gen16481 := Call(__e, ShenFunc(symcons), B, Nil)

																	gen16482 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16481)

																	gen16483 := Call(__e, ShenFunc(symcons), A, gen16482)

																	Call(__e, ShenFunc(symshen_4bindv), V2636, gen16483, V2783)

																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen16484 := MakeNative(func(__e Evaluator, __args ...Obj) {
																		__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																		return
																	}, 0)
																	gen16485 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16484)

																	Result := gen16485
																	Call(__e, ShenFunc(symshen_4unbindv), V2636, V2783)
																	gen16540 = Result

																} else {
																	gen16540 = False
																}

															}

														} else {
															gen16540 = False
														}

													} else {
														gen16540 = False
													}

												} else {
													gen16540 = False
												}

											} else {
												gen16540 = False
											}

										} else {
											gen16540 = False
										}
										Case := gen16540
										gen17082 := Call(__e, ShenFunc(sym_a), Case, False)

										if True == gen17082 {
											gen16541 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

											V2643 := gen16541
											gen16615 := Call(__e, ShenFunc(symcons_2), V2643)

											var gen16616 Obj
											if True == gen16615 {
												gen16542 := Call(__e, ShenFunc(symhd), V2643)

												gen16543 := Call(__e, ShenFunc(symshen_4lazyderef), gen16542, V2783)

												V2644 := gen16543
												gen16614 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), V2644)

												if True == gen16614 {
													gen16544 := Call(__e, ShenFunc(symtl), V2643)

													gen16545 := Call(__e, ShenFunc(symshen_4lazyderef), gen16544, V2783)

													V2645 := gen16545
													gen16613 := Call(__e, ShenFunc(symcons_2), V2645)

													if True == gen16613 {
														gen16546 := Call(__e, ShenFunc(symhd), V2645)

														X := gen16546
														gen16547 := Call(__e, ShenFunc(symtl), V2645)

														gen16548 := Call(__e, ShenFunc(symshen_4lazyderef), gen16547, V2783)

														V2646 := gen16548
														gen16612 := Call(__e, ShenFunc(symcons_2), V2646)

														if True == gen16612 {
															gen16549 := Call(__e, ShenFunc(symhd), V2646)

															Y := gen16549
															gen16550 := Call(__e, ShenFunc(symtl), V2646)

															gen16551 := Call(__e, ShenFunc(symshen_4lazyderef), gen16550, V2783)

															V2647 := gen16551
															gen16611 := Call(__e, ShenFunc(sym_a), Nil, V2647)

															if True == gen16611 {
																gen16552 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																V2648 := gen16552
																gen16610 := Call(__e, ShenFunc(symcons_2), V2648)

																if True == gen16610 {
																	gen16561 := Call(__e, ShenFunc(symhd), V2648)

																	gen16562 := Call(__e, ShenFunc(symshen_4lazyderef), gen16561, V2783)

																	V2649 := gen16562
																	gen16609 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), V2649)

																	if True == gen16609 {
																		gen16587 := Call(__e, ShenFunc(symtl), V2648)

																		gen16588 := Call(__e, ShenFunc(symshen_4lazyderef), gen16587, V2783)

																		V2650 := gen16588
																		gen16608 := Call(__e, ShenFunc(symcons_2), V2650)

																		if True == gen16608 {
																			gen16596 := Call(__e, ShenFunc(symhd), V2650)

																			A := gen16596
																			gen16597 := Call(__e, ShenFunc(symtl), V2650)

																			gen16598 := Call(__e, ShenFunc(symshen_4lazyderef), gen16597, V2783)

																			V2651 := gen16598
																			gen16607 := Call(__e, ShenFunc(sym_a), Nil, V2651)

																			if True == gen16607 {
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16606 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					gen16604 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen16605 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16604)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16605, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen16616 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16606)

																			} else {
																				gen16603 := Call(__e, ShenFunc(symshen_4pvar_2), V2651)

																				if True == gen16603 {
																					Call(__e, ShenFunc(symshen_4bindv), V2651, Nil, V2783)
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16601 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16599 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen16600 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16599)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16600, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen16602 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16601)

																					Result := gen16602
																					Call(__e, ShenFunc(symshen_4unbindv), V2651, V2783)
																					gen16616 = Result

																				} else {
																					gen16616 = False
																				}

																			}

																		} else {
																			gen16595 := Call(__e, ShenFunc(symshen_4pvar_2), V2650)

																			if True == gen16595 {
																				gen16589 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				A := gen16589
																				gen16590 := Call(__e, ShenFunc(symcons), A, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2650, gen16590, V2783)

																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16593 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					gen16591 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen16592 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16591)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16592, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen16594 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16593)

																				Result := gen16594
																				Call(__e, ShenFunc(symshen_4unbindv), V2650, V2783)
																				gen16616 = Result

																			} else {
																				gen16616 = False
																			}

																		}

																	} else {
																		gen16586 := Call(__e, ShenFunc(symshen_4pvar_2), V2649)

																		if True == gen16586 {
																			Call(__e, ShenFunc(symshen_4bindv), V2649, MakeSymbol("vector"), V2783)
																			gen16563 := Call(__e, ShenFunc(symtl), V2648)

																			gen16564 := Call(__e, ShenFunc(symshen_4lazyderef), gen16563, V2783)

																			V2652 := gen16564
																			gen16584 := Call(__e, ShenFunc(symcons_2), V2652)

																			var gen16585 Obj
																			if True == gen16584 {
																				gen16572 := Call(__e, ShenFunc(symhd), V2652)

																				A := gen16572
																				gen16573 := Call(__e, ShenFunc(symtl), V2652)

																				gen16574 := Call(__e, ShenFunc(symshen_4lazyderef), gen16573, V2783)

																				V2653 := gen16574
																				gen16583 := Call(__e, ShenFunc(sym_a), Nil, V2653)

																				if True == gen16583 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16582 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16580 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen16581 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16580)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16581, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen16585 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16582)

																				} else {
																					gen16579 := Call(__e, ShenFunc(symshen_4pvar_2), V2653)

																					if True == gen16579 {
																						Call(__e, ShenFunc(symshen_4bindv), V2653, Nil, V2783)
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen16577 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							gen16575 := Call(__e, ShenFunc(symcons), A, Nil)

																							gen16576 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16575)

																							__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16576, V2782, V2783, V2784)

																							return

																						}, 0)
																						gen16578 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16577)

																						Result := gen16578
																						Call(__e, ShenFunc(symshen_4unbindv), V2653, V2783)
																						gen16585 = Result

																					} else {
																						gen16585 = False
																					}

																				}

																			} else {
																				gen16571 := Call(__e, ShenFunc(symshen_4pvar_2), V2652)

																				if True == gen16571 {
																					gen16565 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					A := gen16565
																					gen16566 := Call(__e, ShenFunc(symcons), A, Nil)

																					Call(__e, ShenFunc(symshen_4bindv), V2652, gen16566, V2783)

																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16569 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16567 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen16568 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16567)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16568, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen16570 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16569)

																					Result := gen16570
																					Call(__e, ShenFunc(symshen_4unbindv), V2652, V2783)
																					gen16585 = Result

																				} else {
																					gen16585 = False
																				}

																			}
																			Result := gen16585
																			Call(__e, ShenFunc(symshen_4unbindv), V2649, V2783)
																			gen16616 = Result

																		} else {
																			gen16616 = False
																		}

																	}

																} else {
																	gen16560 := Call(__e, ShenFunc(symshen_4pvar_2), V2648)

																	if True == gen16560 {
																		gen16553 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		A := gen16553
																		gen16554 := Call(__e, ShenFunc(symcons), A, Nil)

																		gen16555 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16554)

																		Call(__e, ShenFunc(symshen_4bindv), V2648, gen16555, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen16558 := MakeNative(func(__e Evaluator, __args ...Obj) {
																			gen16556 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen16557 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16556)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen16557, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen16559 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen16558)

																		Result := gen16559
																		Call(__e, ShenFunc(symshen_4unbindv), V2648, V2783)
																		gen16616 = Result

																	} else {
																		gen16616 = False
																	}

																}

															} else {
																gen16616 = False
															}

														} else {
															gen16616 = False
														}

													} else {
														gen16616 = False
													}

												} else {
													gen16616 = False
												}

											} else {
												gen16616 = False
											}
											Case := gen16616
											gen17081 := Call(__e, ShenFunc(sym_a), Case, False)

											if True == gen17081 {
												gen16617 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

												V2654 := gen16617
												gen16638 := Call(__e, ShenFunc(symcons_2), V2654)

												var gen16639 Obj
												if True == gen16638 {
													gen16618 := Call(__e, ShenFunc(symhd), V2654)

													gen16619 := Call(__e, ShenFunc(symshen_4lazyderef), gen16618, V2783)

													V2655 := gen16619
													gen16637 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), V2655)

													if True == gen16637 {
														gen16620 := Call(__e, ShenFunc(symtl), V2654)

														gen16621 := Call(__e, ShenFunc(symshen_4lazyderef), gen16620, V2783)

														V2656 := gen16621
														gen16636 := Call(__e, ShenFunc(symcons_2), V2656)

														if True == gen16636 {
															gen16622 := Call(__e, ShenFunc(symhd), V2656)

															X := gen16622
															gen16623 := Call(__e, ShenFunc(symtl), V2656)

															gen16624 := Call(__e, ShenFunc(symshen_4lazyderef), gen16623, V2783)

															V2657 := gen16624
															gen16635 := Call(__e, ShenFunc(symcons_2), V2657)

															if True == gen16635 {
																gen16625 := Call(__e, ShenFunc(symhd), V2657)

																Y := gen16625
																gen16626 := Call(__e, ShenFunc(symtl), V2657)

																gen16627 := Call(__e, ShenFunc(symshen_4lazyderef), gen16626, V2783)

																V2658 := gen16627
																gen16634 := Call(__e, ShenFunc(sym_a), Nil, V2658)

																if True == gen16634 {
																	gen16628 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																	V2659 := gen16628
																	gen16633 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2659)

																	if True == gen16633 {
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen16632 := MakeNative(func(__e Evaluator, __args ...Obj) {
																			__e.TailApply(ShenFunc(symshen_4th_d), Y, MakeSymbol("string"), V2782, V2783, V2784)

																			return
																		}, 0)
																		gen16639 = Call(__e, ShenFunc(symshen_4th_d), X, MakeSymbol("string"), V2782, V2783, gen16632)

																	} else {
																		gen16631 := Call(__e, ShenFunc(symshen_4pvar_2), V2659)

																		if True == gen16631 {
																			Call(__e, ShenFunc(symshen_4bindv), V2659, MakeSymbol("string"), V2783)
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen16629 := MakeNative(func(__e Evaluator, __args ...Obj) {
																				__e.TailApply(ShenFunc(symshen_4th_d), Y, MakeSymbol("string"), V2782, V2783, V2784)

																				return
																			}, 0)
																			gen16630 := Call(__e, ShenFunc(symshen_4th_d), X, MakeSymbol("string"), V2782, V2783, gen16629)

																			Result := gen16630
																			Call(__e, ShenFunc(symshen_4unbindv), V2659, V2783)
																			gen16639 = Result

																		} else {
																			gen16639 = False
																		}

																	}

																} else {
																	gen16639 = False
																}

															} else {
																gen16639 = False
															}

														} else {
															gen16639 = False
														}

													} else {
														gen16639 = False
													}

												} else {
													gen16639 = False
												}
												Case := gen16639
												gen17080 := Call(__e, ShenFunc(sym_a), Case, False)

												if True == gen17080 {
													gen16640 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

													V2660 := gen16640
													gen16808 := Call(__e, ShenFunc(symcons_2), V2660)

													var gen16809 Obj
													if True == gen16808 {
														gen16641 := Call(__e, ShenFunc(symhd), V2660)

														gen16642 := Call(__e, ShenFunc(symshen_4lazyderef), gen16641, V2783)

														V2661 := gen16642
														gen16807 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), V2661)

														if True == gen16807 {
															gen16643 := Call(__e, ShenFunc(symtl), V2660)

															gen16644 := Call(__e, ShenFunc(symshen_4lazyderef), gen16643, V2783)

															V2662 := gen16644
															gen16806 := Call(__e, ShenFunc(symcons_2), V2662)

															if True == gen16806 {
																gen16645 := Call(__e, ShenFunc(symhd), V2662)

																X := gen16645
																gen16646 := Call(__e, ShenFunc(symtl), V2662)

																gen16647 := Call(__e, ShenFunc(symshen_4lazyderef), gen16646, V2783)

																V2663 := gen16647
																gen16805 := Call(__e, ShenFunc(symcons_2), V2663)

																if True == gen16805 {
																	gen16648 := Call(__e, ShenFunc(symhd), V2663)

																	Y := gen16648
																	gen16649 := Call(__e, ShenFunc(symtl), V2663)

																	gen16650 := Call(__e, ShenFunc(symshen_4lazyderef), gen16649, V2783)

																	V2664 := gen16650
																	gen16804 := Call(__e, ShenFunc(sym_a), Nil, V2664)

																	if True == gen16804 {
																		gen16651 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																		V2665 := gen16651
																		gen16803 := Call(__e, ShenFunc(symcons_2), V2665)

																		if True == gen16803 {
																			gen16672 := Call(__e, ShenFunc(symhd), V2665)

																			A := gen16672
																			gen16673 := Call(__e, ShenFunc(symtl), V2665)

																			gen16674 := Call(__e, ShenFunc(symshen_4lazyderef), gen16673, V2783)

																			V2666 := gen16674
																			gen16802 := Call(__e, ShenFunc(symcons_2), V2666)

																			if True == gen16802 {
																				gen16693 := Call(__e, ShenFunc(symhd), V2666)

																				gen16694 := Call(__e, ShenFunc(symshen_4lazyderef), gen16693, V2783)

																				V2667 := gen16694
																				gen16801 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), V2667)

																				if True == gen16801 {
																					gen16749 := Call(__e, ShenFunc(symtl), V2666)

																					gen16750 := Call(__e, ShenFunc(symshen_4lazyderef), gen16749, V2783)

																					V2668 := gen16750
																					gen16800 := Call(__e, ShenFunc(symcons_2), V2668)

																					if True == gen16800 {
																						gen16768 := Call(__e, ShenFunc(symhd), V2668)

																						B := gen16768
																						gen16769 := Call(__e, ShenFunc(symtl), V2668)

																						gen16770 := Call(__e, ShenFunc(symshen_4lazyderef), gen16769, V2783)

																						V2669 := gen16770
																						gen16799 := Call(__e, ShenFunc(sym_a), Nil, V2669)

																						if True == gen16799 {
																							gen16786 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							Z := gen16786
																							gen16787 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							X_e_e := gen16787
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen16788 := Call(__e, ShenFunc(symshen_4placeholder))

																							gen16798 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen16789 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																								gen16790 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																								gen16791 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																								gen16792 := Call(__e, ShenFunc(symshen_4ebr), gen16789, gen16790, gen16791)

																								gen16797 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16793 := Call(__e, ShenFunc(symcons), A, Nil)

																									gen16794 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16793)

																									gen16795 := Call(__e, ShenFunc(symcons), X_e_e, gen16794)

																									gen16796 := Call(__e, ShenFunc(symcons), gen16795, V2782)

																									__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16796, V2783, V2784)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symbind), Z, gen16792, V2783, gen16797)

																								return

																							}, 0)
																							gen16809 = Call(__e, ShenFunc(symbind), X_e_e, gen16788, V2783, gen16798)

																						} else {
																							gen16785 := Call(__e, ShenFunc(symshen_4pvar_2), V2669)

																							if True == gen16785 {
																								Call(__e, ShenFunc(symshen_4bindv), V2669, Nil, V2783)
																								gen16771 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen16771
																								gen16772 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen16772
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen16773 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen16783 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16774 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen16775 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen16776 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen16777 := Call(__e, ShenFunc(symshen_4ebr), gen16774, gen16775, gen16776)

																									gen16782 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16778 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen16779 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16778)

																										gen16780 := Call(__e, ShenFunc(symcons), X_e_e, gen16779)

																										gen16781 := Call(__e, ShenFunc(symcons), gen16780, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16781, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen16777, V2783, gen16782)

																									return

																								}, 0)
																								gen16784 := Call(__e, ShenFunc(symbind), X_e_e, gen16773, V2783, gen16783)

																								Result := gen16784
																								Call(__e, ShenFunc(symshen_4unbindv), V2669, V2783)
																								gen16809 = Result

																							} else {
																								gen16809 = False
																							}

																						}

																					} else {
																						gen16767 := Call(__e, ShenFunc(symshen_4pvar_2), V2668)

																						if True == gen16767 {
																							gen16751 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							B := gen16751
																							gen16752 := Call(__e, ShenFunc(symcons), B, Nil)

																							Call(__e, ShenFunc(symshen_4bindv), V2668, gen16752, V2783)

																							gen16753 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							Z := gen16753
																							gen16754 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							X_e_e := gen16754
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen16755 := Call(__e, ShenFunc(symshen_4placeholder))

																							gen16765 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen16756 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																								gen16757 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																								gen16758 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																								gen16759 := Call(__e, ShenFunc(symshen_4ebr), gen16756, gen16757, gen16758)

																								gen16764 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16760 := Call(__e, ShenFunc(symcons), A, Nil)

																									gen16761 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16760)

																									gen16762 := Call(__e, ShenFunc(symcons), X_e_e, gen16761)

																									gen16763 := Call(__e, ShenFunc(symcons), gen16762, V2782)

																									__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16763, V2783, V2784)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symbind), Z, gen16759, V2783, gen16764)

																								return

																							}, 0)
																							gen16766 := Call(__e, ShenFunc(symbind), X_e_e, gen16755, V2783, gen16765)

																							Result := gen16766
																							Call(__e, ShenFunc(symshen_4unbindv), V2668, V2783)
																							gen16809 = Result

																						} else {
																							gen16809 = False
																						}

																					}

																				} else {
																					gen16748 := Call(__e, ShenFunc(symshen_4pvar_2), V2667)

																					if True == gen16748 {
																						Call(__e, ShenFunc(symshen_4bindv), V2667, MakeSymbol("-->"), V2783)
																						gen16695 := Call(__e, ShenFunc(symtl), V2666)

																						gen16696 := Call(__e, ShenFunc(symshen_4lazyderef), gen16695, V2783)

																						V2670 := gen16696
																						gen16746 := Call(__e, ShenFunc(symcons_2), V2670)

																						var gen16747 Obj
																						if True == gen16746 {
																							gen16714 := Call(__e, ShenFunc(symhd), V2670)

																							B := gen16714
																							gen16715 := Call(__e, ShenFunc(symtl), V2670)

																							gen16716 := Call(__e, ShenFunc(symshen_4lazyderef), gen16715, V2783)

																							V2671 := gen16716
																							gen16745 := Call(__e, ShenFunc(sym_a), Nil, V2671)

																							if True == gen16745 {
																								gen16732 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen16732
																								gen16733 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen16733
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen16734 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen16744 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16735 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen16736 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen16737 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen16738 := Call(__e, ShenFunc(symshen_4ebr), gen16735, gen16736, gen16737)

																									gen16743 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16739 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen16740 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16739)

																										gen16741 := Call(__e, ShenFunc(symcons), X_e_e, gen16740)

																										gen16742 := Call(__e, ShenFunc(symcons), gen16741, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16742, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen16738, V2783, gen16743)

																									return

																								}, 0)
																								gen16747 = Call(__e, ShenFunc(symbind), X_e_e, gen16734, V2783, gen16744)

																							} else {
																								gen16731 := Call(__e, ShenFunc(symshen_4pvar_2), V2671)

																								if True == gen16731 {
																									Call(__e, ShenFunc(symshen_4bindv), V2671, Nil, V2783)
																									gen16717 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									Z := gen16717
																									gen16718 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									X_e_e := gen16718
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen16719 := Call(__e, ShenFunc(symshen_4placeholder))

																									gen16729 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16720 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																										gen16721 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																										gen16722 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																										gen16723 := Call(__e, ShenFunc(symshen_4ebr), gen16720, gen16721, gen16722)

																										gen16728 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											gen16724 := Call(__e, ShenFunc(symcons), A, Nil)

																											gen16725 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16724)

																											gen16726 := Call(__e, ShenFunc(symcons), X_e_e, gen16725)

																											gen16727 := Call(__e, ShenFunc(symcons), gen16726, V2782)

																											__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16727, V2783, V2784)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symbind), Z, gen16723, V2783, gen16728)

																										return

																									}, 0)
																									gen16730 := Call(__e, ShenFunc(symbind), X_e_e, gen16719, V2783, gen16729)

																									Result := gen16730
																									Call(__e, ShenFunc(symshen_4unbindv), V2671, V2783)
																									gen16747 = Result

																								} else {
																									gen16747 = False
																								}

																							}

																						} else {
																							gen16713 := Call(__e, ShenFunc(symshen_4pvar_2), V2670)

																							if True == gen16713 {
																								gen16697 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								B := gen16697
																								gen16698 := Call(__e, ShenFunc(symcons), B, Nil)

																								Call(__e, ShenFunc(symshen_4bindv), V2670, gen16698, V2783)

																								gen16699 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen16699
																								gen16700 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen16700
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen16701 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen16711 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16702 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen16703 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen16704 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen16705 := Call(__e, ShenFunc(symshen_4ebr), gen16702, gen16703, gen16704)

																									gen16710 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16706 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen16707 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16706)

																										gen16708 := Call(__e, ShenFunc(symcons), X_e_e, gen16707)

																										gen16709 := Call(__e, ShenFunc(symcons), gen16708, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16709, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen16705, V2783, gen16710)

																									return

																								}, 0)
																								gen16712 := Call(__e, ShenFunc(symbind), X_e_e, gen16701, V2783, gen16711)

																								Result := gen16712
																								Call(__e, ShenFunc(symshen_4unbindv), V2670, V2783)
																								gen16747 = Result

																							} else {
																								gen16747 = False
																							}

																						}
																						Result := gen16747
																						Call(__e, ShenFunc(symshen_4unbindv), V2667, V2783)
																						gen16809 = Result

																					} else {
																						gen16809 = False
																					}

																				}

																			} else {
																				gen16692 := Call(__e, ShenFunc(symshen_4pvar_2), V2666)

																				if True == gen16692 {
																					gen16675 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					B := gen16675
																					gen16676 := Call(__e, ShenFunc(symcons), B, Nil)

																					gen16677 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16676)

																					Call(__e, ShenFunc(symshen_4bindv), V2666, gen16677, V2783)

																					gen16678 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					Z := gen16678
																					gen16679 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					X_e_e := gen16679
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16680 := Call(__e, ShenFunc(symshen_4placeholder))

																					gen16690 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16681 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																						gen16682 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																						gen16683 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																						gen16684 := Call(__e, ShenFunc(symshen_4ebr), gen16681, gen16682, gen16683)

																						gen16689 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							gen16685 := Call(__e, ShenFunc(symcons), A, Nil)

																							gen16686 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16685)

																							gen16687 := Call(__e, ShenFunc(symcons), X_e_e, gen16686)

																							gen16688 := Call(__e, ShenFunc(symcons), gen16687, V2782)

																							__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16688, V2783, V2784)

																							return

																						}, 0)
																						__e.TailApply(ShenFunc(symbind), Z, gen16684, V2783, gen16689)

																						return

																					}, 0)
																					gen16691 := Call(__e, ShenFunc(symbind), X_e_e, gen16680, V2783, gen16690)

																					Result := gen16691
																					Call(__e, ShenFunc(symshen_4unbindv), V2666, V2783)
																					gen16809 = Result

																				} else {
																					gen16809 = False
																				}

																			}

																		} else {
																			gen16671 := Call(__e, ShenFunc(symshen_4pvar_2), V2665)

																			if True == gen16671 {
																				gen16652 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				A := gen16652
																				gen16653 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen16653
																				gen16654 := Call(__e, ShenFunc(symcons), B, Nil)

																				gen16655 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16654)

																				gen16656 := Call(__e, ShenFunc(symcons), A, gen16655)

																				Call(__e, ShenFunc(symshen_4bindv), V2665, gen16656, V2783)

																				gen16657 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				Z := gen16657
																				gen16658 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				X_e_e := gen16658
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16659 := Call(__e, ShenFunc(symshen_4placeholder))

																				gen16669 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					gen16660 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																					gen16661 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																					gen16662 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																					gen16663 := Call(__e, ShenFunc(symshen_4ebr), gen16660, gen16661, gen16662)

																					gen16668 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16664 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen16665 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16664)

																						gen16666 := Call(__e, ShenFunc(symcons), X_e_e, gen16665)

																						gen16667 := Call(__e, ShenFunc(symcons), gen16666, V2782)

																						__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen16667, V2783, V2784)

																						return

																					}, 0)
																					__e.TailApply(ShenFunc(symbind), Z, gen16663, V2783, gen16668)

																					return

																				}, 0)
																				gen16670 := Call(__e, ShenFunc(symbind), X_e_e, gen16659, V2783, gen16669)

																				Result := gen16670
																				Call(__e, ShenFunc(symshen_4unbindv), V2665, V2783)
																				gen16809 = Result

																			} else {
																				gen16809 = False
																			}

																		}

																	} else {
																		gen16809 = False
																	}

																} else {
																	gen16809 = False
																}

															} else {
																gen16809 = False
															}

														} else {
															gen16809 = False
														}

													} else {
														gen16809 = False
													}
													Case := gen16809
													gen17079 := Call(__e, ShenFunc(sym_a), Case, False)

													if True == gen17079 {
														gen16810 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

														V2672 := gen16810
														gen16844 := Call(__e, ShenFunc(symcons_2), V2672)

														var gen16845 Obj
														if True == gen16844 {
															gen16811 := Call(__e, ShenFunc(symhd), V2672)

															gen16812 := Call(__e, ShenFunc(symshen_4lazyderef), gen16811, V2783)

															V2673 := gen16812
															gen16843 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), V2673)

															if True == gen16843 {
																gen16813 := Call(__e, ShenFunc(symtl), V2672)

																gen16814 := Call(__e, ShenFunc(symshen_4lazyderef), gen16813, V2783)

																V2674 := gen16814
																gen16842 := Call(__e, ShenFunc(symcons_2), V2674)

																if True == gen16842 {
																	gen16815 := Call(__e, ShenFunc(symhd), V2674)

																	X := gen16815
																	gen16816 := Call(__e, ShenFunc(symtl), V2674)

																	gen16817 := Call(__e, ShenFunc(symshen_4lazyderef), gen16816, V2783)

																	V2675 := gen16817
																	gen16841 := Call(__e, ShenFunc(symcons_2), V2675)

																	if True == gen16841 {
																		gen16818 := Call(__e, ShenFunc(symhd), V2675)

																		Y := gen16818
																		gen16819 := Call(__e, ShenFunc(symtl), V2675)

																		gen16820 := Call(__e, ShenFunc(symshen_4lazyderef), gen16819, V2783)

																		V2676 := gen16820
																		gen16840 := Call(__e, ShenFunc(symcons_2), V2676)

																		if True == gen16840 {
																			gen16821 := Call(__e, ShenFunc(symhd), V2676)

																			Z := gen16821
																			gen16822 := Call(__e, ShenFunc(symtl), V2676)

																			gen16823 := Call(__e, ShenFunc(symshen_4lazyderef), gen16822, V2783)

																			V2677 := gen16823
																			gen16839 := Call(__e, ShenFunc(sym_a), Nil, V2677)

																			if True == gen16839 {
																				gen16824 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				W := gen16824
																				gen16825 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				X_e_e := gen16825
																				gen16826 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen16826
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen16838 := MakeNative(func(__e Evaluator, __args ...Obj) {
																					gen16827 := Call(__e, ShenFunc(symshen_4placeholder))

																					gen16837 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16828 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																						gen16829 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																						gen16830 := Call(__e, ShenFunc(symshen_4lazyderef), Z, V2783)

																						gen16831 := Call(__e, ShenFunc(symshen_4ebr), gen16828, gen16829, gen16830)

																						gen16836 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							gen16832 := Call(__e, ShenFunc(symcons), B, Nil)

																							gen16833 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen16832)

																							gen16834 := Call(__e, ShenFunc(symcons), X_e_e, gen16833)

																							gen16835 := Call(__e, ShenFunc(symcons), gen16834, V2782)

																							__e.TailApply(ShenFunc(symshen_4th_d), W, V2781, gen16835, V2783, V2784)

																							return

																						}, 0)
																						__e.TailApply(ShenFunc(symbind), W, gen16831, V2783, gen16836)

																						return

																					}, 0)
																					__e.TailApply(ShenFunc(symbind), X_e_e, gen16827, V2783, gen16837)

																					return

																				}, 0)
																				gen16845 = Call(__e, ShenFunc(symshen_4th_d), Y, B, V2782, V2783, gen16838)

																			} else {
																				gen16845 = False
																			}

																		} else {
																			gen16845 = False
																		}

																	} else {
																		gen16845 = False
																	}

																} else {
																	gen16845 = False
																}

															} else {
																gen16845 = False
															}

														} else {
															gen16845 = False
														}
														Case := gen16845
														gen17078 := Call(__e, ShenFunc(sym_a), Case, False)

														if True == gen17078 {
															gen16846 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

															V2678 := gen16846
															gen16948 := Call(__e, ShenFunc(symcons_2), V2678)

															var gen16949 Obj
															if True == gen16948 {
																gen16847 := Call(__e, ShenFunc(symhd), V2678)

																gen16848 := Call(__e, ShenFunc(symshen_4lazyderef), gen16847, V2783)

																V2679 := gen16848
																gen16947 := Call(__e, ShenFunc(sym_a), MakeSymbol("open"), V2679)

																if True == gen16947 {
																	gen16849 := Call(__e, ShenFunc(symtl), V2678)

																	gen16850 := Call(__e, ShenFunc(symshen_4lazyderef), gen16849, V2783)

																	V2680 := gen16850
																	gen16946 := Call(__e, ShenFunc(symcons_2), V2680)

																	if True == gen16946 {
																		gen16851 := Call(__e, ShenFunc(symhd), V2680)

																		FileName := gen16851
																		gen16852 := Call(__e, ShenFunc(symtl), V2680)

																		gen16853 := Call(__e, ShenFunc(symshen_4lazyderef), gen16852, V2783)

																		V2681 := gen16853
																		gen16945 := Call(__e, ShenFunc(symcons_2), V2681)

																		if True == gen16945 {
																			gen16854 := Call(__e, ShenFunc(symhd), V2681)

																			Direction2611 := gen16854
																			gen16855 := Call(__e, ShenFunc(symtl), V2681)

																			gen16856 := Call(__e, ShenFunc(symshen_4lazyderef), gen16855, V2783)

																			V2682 := gen16856
																			gen16944 := Call(__e, ShenFunc(sym_a), Nil, V2682)

																			if True == gen16944 {
																				gen16857 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																				V2683 := gen16857
																				gen16943 := Call(__e, ShenFunc(symcons_2), V2683)

																				if True == gen16943 {
																					gen16870 := Call(__e, ShenFunc(symhd), V2683)

																					gen16871 := Call(__e, ShenFunc(symshen_4lazyderef), gen16870, V2783)

																					V2684 := gen16871
																					gen16942 := Call(__e, ShenFunc(sym_a), MakeSymbol("stream"), V2684)

																					if True == gen16942 {
																						gen16908 := Call(__e, ShenFunc(symtl), V2683)

																						gen16909 := Call(__e, ShenFunc(symshen_4lazyderef), gen16908, V2783)

																						V2685 := gen16909
																						gen16941 := Call(__e, ShenFunc(symcons_2), V2685)

																						if True == gen16941 {
																							gen16921 := Call(__e, ShenFunc(symhd), V2685)

																							Direction := gen16921
																							gen16922 := Call(__e, ShenFunc(symtl), V2685)

																							gen16923 := Call(__e, ShenFunc(symshen_4lazyderef), gen16922, V2783)

																							V2686 := gen16923
																							gen16940 := Call(__e, ShenFunc(sym_a), Nil, V2686)

																							if True == gen16940 {
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen16939 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16938 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16933 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																										gen16934 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																										gen16935 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16934)

																										gen16936 := Call(__e, ShenFunc(symelement_2), gen16933, gen16935)

																										gen16937 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symfwhen), gen16936, V2783, gen16937)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16938)

																									return

																								}, 0)
																								gen16949 = Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16939)

																							} else {
																								gen16932 := Call(__e, ShenFunc(symshen_4pvar_2), V2686)

																								if True == gen16932 {
																									Call(__e, ShenFunc(symshen_4bindv), V2686, Nil, V2783)
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen16930 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16929 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											gen16924 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen16925 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen16926 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16925)

																											gen16927 := Call(__e, ShenFunc(symelement_2), gen16924, gen16926)

																											gen16928 := MakeNative(func(__e Evaluator, __args ...Obj) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen16927, V2783, gen16928)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16929)

																										return

																									}, 0)
																									gen16931 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16930)

																									Result := gen16931
																									Call(__e, ShenFunc(symshen_4unbindv), V2686, V2783)
																									gen16949 = Result

																								} else {
																									gen16949 = False
																								}

																							}

																						} else {
																							gen16920 := Call(__e, ShenFunc(symshen_4pvar_2), V2685)

																							if True == gen16920 {
																								gen16910 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Direction := gen16910
																								gen16911 := Call(__e, ShenFunc(symcons), Direction, Nil)

																								Call(__e, ShenFunc(symshen_4bindv), V2685, gen16911, V2783)

																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen16918 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen16917 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16912 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																										gen16913 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																										gen16914 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16913)

																										gen16915 := Call(__e, ShenFunc(symelement_2), gen16912, gen16914)

																										gen16916 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symfwhen), gen16915, V2783, gen16916)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16917)

																									return

																								}, 0)
																								gen16919 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16918)

																								Result := gen16919
																								Call(__e, ShenFunc(symshen_4unbindv), V2685, V2783)
																								gen16949 = Result

																							} else {
																								gen16949 = False
																							}

																						}

																					} else {
																						gen16907 := Call(__e, ShenFunc(symshen_4pvar_2), V2684)

																						if True == gen16907 {
																							Call(__e, ShenFunc(symshen_4bindv), V2684, MakeSymbol("stream"), V2783)
																							gen16872 := Call(__e, ShenFunc(symtl), V2683)

																							gen16873 := Call(__e, ShenFunc(symshen_4lazyderef), gen16872, V2783)

																							V2687 := gen16873
																							gen16905 := Call(__e, ShenFunc(symcons_2), V2687)

																							var gen16906 Obj
																							if True == gen16905 {
																								gen16885 := Call(__e, ShenFunc(symhd), V2687)

																								Direction := gen16885
																								gen16886 := Call(__e, ShenFunc(symtl), V2687)

																								gen16887 := Call(__e, ShenFunc(symshen_4lazyderef), gen16886, V2783)

																								V2688 := gen16887
																								gen16904 := Call(__e, ShenFunc(sym_a), Nil, V2688)

																								if True == gen16904 {
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen16903 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16902 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											gen16897 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen16898 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen16899 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16898)

																											gen16900 := Call(__e, ShenFunc(symelement_2), gen16897, gen16899)

																											gen16901 := MakeNative(func(__e Evaluator, __args ...Obj) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen16900, V2783, gen16901)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16902)

																										return

																									}, 0)
																									gen16906 = Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16903)

																								} else {
																									gen16896 := Call(__e, ShenFunc(symshen_4pvar_2), V2688)

																									if True == gen16896 {
																										Call(__e, ShenFunc(symshen_4bindv), V2688, Nil, V2783)
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen16894 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											gen16893 := MakeNative(func(__e Evaluator, __args ...Obj) {
																												gen16888 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																												gen16889 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																												gen16890 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16889)

																												gen16891 := Call(__e, ShenFunc(symelement_2), gen16888, gen16890)

																												gen16892 := MakeNative(func(__e Evaluator, __args ...Obj) {
																													__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																													return
																												}, 0)
																												__e.TailApply(ShenFunc(symfwhen), gen16891, V2783, gen16892)

																												return

																											}, 0)
																											__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16893)

																											return

																										}, 0)
																										gen16895 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16894)

																										Result := gen16895
																										Call(__e, ShenFunc(symshen_4unbindv), V2688, V2783)
																										gen16906 = Result

																									} else {
																										gen16906 = False
																									}

																								}

																							} else {
																								gen16884 := Call(__e, ShenFunc(symshen_4pvar_2), V2687)

																								if True == gen16884 {
																									gen16874 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									Direction := gen16874
																									gen16875 := Call(__e, ShenFunc(symcons), Direction, Nil)

																									Call(__e, ShenFunc(symshen_4bindv), V2687, gen16875, V2783)

																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen16882 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen16881 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											gen16876 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen16877 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen16878 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16877)

																											gen16879 := Call(__e, ShenFunc(symelement_2), gen16876, gen16878)

																											gen16880 := MakeNative(func(__e Evaluator, __args ...Obj) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen16879, V2783, gen16880)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16881)

																										return

																									}, 0)
																									gen16883 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16882)

																									Result := gen16883
																									Call(__e, ShenFunc(symshen_4unbindv), V2687, V2783)
																									gen16906 = Result

																								} else {
																									gen16906 = False
																								}

																							}
																							Result := gen16906
																							Call(__e, ShenFunc(symshen_4unbindv), V2684, V2783)
																							gen16949 = Result

																						} else {
																							gen16949 = False
																						}

																					}

																				} else {
																					gen16869 := Call(__e, ShenFunc(symshen_4pvar_2), V2683)

																					if True == gen16869 {
																						gen16858 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																						Direction := gen16858
																						gen16859 := Call(__e, ShenFunc(symcons), Direction, Nil)

																						gen16860 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16859)

																						Call(__e, ShenFunc(symshen_4bindv), V2683, gen16860, V2783)

																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen16867 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							gen16866 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen16861 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																								gen16862 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																								gen16863 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen16862)

																								gen16864 := Call(__e, ShenFunc(symelement_2), gen16861, gen16863)

																								gen16865 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																									return
																								}, 0)
																								__e.TailApply(ShenFunc(symfwhen), gen16864, V2783, gen16865)

																								return

																							}, 0)
																							__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen16866)

																							return

																						}, 0)
																						gen16868 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen16867)

																						Result := gen16868
																						Call(__e, ShenFunc(symshen_4unbindv), V2683, V2783)
																						gen16949 = Result

																					} else {
																						gen16949 = False
																					}

																				}

																			} else {
																				gen16949 = False
																			}

																		} else {
																			gen16949 = False
																		}

																	} else {
																		gen16949 = False
																	}

																} else {
																	gen16949 = False
																}

															} else {
																gen16949 = False
															}
															Case := gen16949
															gen17077 := Call(__e, ShenFunc(sym_a), Case, False)

															if True == gen17077 {
																gen16950 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																V2689 := gen16950
																gen16967 := Call(__e, ShenFunc(symcons_2), V2689)

																var gen16968 Obj
																if True == gen16967 {
																	gen16951 := Call(__e, ShenFunc(symhd), V2689)

																	gen16952 := Call(__e, ShenFunc(symshen_4lazyderef), gen16951, V2783)

																	V2690 := gen16952
																	gen16966 := Call(__e, ShenFunc(sym_a), MakeSymbol("type"), V2690)

																	if True == gen16966 {
																		gen16953 := Call(__e, ShenFunc(symtl), V2689)

																		gen16954 := Call(__e, ShenFunc(symshen_4lazyderef), gen16953, V2783)

																		V2691 := gen16954
																		gen16965 := Call(__e, ShenFunc(symcons_2), V2691)

																		if True == gen16965 {
																			gen16955 := Call(__e, ShenFunc(symhd), V2691)

																			X := gen16955
																			gen16956 := Call(__e, ShenFunc(symtl), V2691)

																			gen16957 := Call(__e, ShenFunc(symshen_4lazyderef), gen16956, V2783)

																			V2692 := gen16957
																			gen16964 := Call(__e, ShenFunc(symcons_2), V2692)

																			if True == gen16964 {
																				gen16958 := Call(__e, ShenFunc(symhd), V2692)

																				A := gen16958
																				gen16959 := Call(__e, ShenFunc(symtl), V2692)

																				gen16960 := Call(__e, ShenFunc(symshen_4lazyderef), gen16959, V2783)

																				V2693 := gen16960
																				gen16963 := Call(__e, ShenFunc(sym_a), Nil, V2693)

																				if True == gen16963 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen16962 := MakeNative(func(__e Evaluator, __args ...Obj) {
																						gen16961 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							__e.TailApply(ShenFunc(symshen_4th_d), X, A, V2782, V2783, V2784)

																							return
																						}, 0)
																						__e.TailApply(ShenFunc(symunify), A, V2781, V2783, gen16961)

																						return

																					}, 0)
																					gen16968 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen16962)

																				} else {
																					gen16968 = False
																				}

																			} else {
																				gen16968 = False
																			}

																		} else {
																			gen16968 = False
																		}

																	} else {
																		gen16968 = False
																	}

																} else {
																	gen16968 = False
																}
																Case := gen16968
																gen17076 := Call(__e, ShenFunc(sym_a), Case, False)

																if True == gen17076 {
																	gen16969 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																	V2694 := gen16969
																	gen16991 := Call(__e, ShenFunc(symcons_2), V2694)

																	var gen16992 Obj
																	if True == gen16991 {
																		gen16970 := Call(__e, ShenFunc(symhd), V2694)

																		gen16971 := Call(__e, ShenFunc(symshen_4lazyderef), gen16970, V2783)

																		V2695 := gen16971
																		gen16990 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), V2695)

																		if True == gen16990 {
																			gen16972 := Call(__e, ShenFunc(symtl), V2694)

																			gen16973 := Call(__e, ShenFunc(symshen_4lazyderef), gen16972, V2783)

																			V2696 := gen16973
																			gen16989 := Call(__e, ShenFunc(symcons_2), V2696)

																			if True == gen16989 {
																				gen16974 := Call(__e, ShenFunc(symhd), V2696)

																				A := gen16974
																				gen16975 := Call(__e, ShenFunc(symtl), V2696)

																				gen16976 := Call(__e, ShenFunc(symshen_4lazyderef), gen16975, V2783)

																				V2697 := gen16976
																				gen16988 := Call(__e, ShenFunc(symcons_2), V2697)

																				if True == gen16988 {
																					gen16977 := Call(__e, ShenFunc(symhd), V2697)

																					Stream := gen16977
																					gen16978 := Call(__e, ShenFunc(symtl), V2697)

																					gen16979 := Call(__e, ShenFunc(symshen_4lazyderef), gen16978, V2783)

																					V2698 := gen16979
																					gen16987 := Call(__e, ShenFunc(sym_a), Nil, V2698)

																					if True == gen16987 {
																						gen16980 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																						C := gen16980
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen16981 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2783)

																						gen16982 := Call(__e, ShenFunc(symshen_4demodulate), gen16981)

																						gen16986 := MakeNative(func(__e Evaluator, __args ...Obj) {
																							gen16985 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen16983 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

																								gen16984 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16983)

																								__e.TailApply(ShenFunc(symshen_4th_d), Stream, gen16984, V2782, V2783, V2784)

																								return

																							}, 0)
																							__e.TailApply(ShenFunc(symunify), V2781, C, V2783, gen16985)

																							return

																						}, 0)
																						gen16992 = Call(__e, ShenFunc(symbind), C, gen16982, V2783, gen16986)

																					} else {
																						gen16992 = False
																					}

																				} else {
																					gen16992 = False
																				}

																			} else {
																				gen16992 = False
																			}

																		} else {
																			gen16992 = False
																		}

																	} else {
																		gen16992 = False
																	}
																	Case := gen16992
																	gen17075 := Call(__e, ShenFunc(sym_a), Case, False)

																	if True == gen17075 {
																		gen16993 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																		V2699 := gen16993
																		gen17014 := Call(__e, ShenFunc(symcons_2), V2699)

																		var gen17015 Obj
																		if True == gen17014 {
																			gen16994 := Call(__e, ShenFunc(symhd), V2699)

																			gen16995 := Call(__e, ShenFunc(symshen_4lazyderef), gen16994, V2783)

																			V2700 := gen16995
																			gen17013 := Call(__e, ShenFunc(sym_a), MakeSymbol("set"), V2700)

																			if True == gen17013 {
																				gen16996 := Call(__e, ShenFunc(symtl), V2699)

																				gen16997 := Call(__e, ShenFunc(symshen_4lazyderef), gen16996, V2783)

																				V2701 := gen16997
																				gen17012 := Call(__e, ShenFunc(symcons_2), V2701)

																				if True == gen17012 {
																					gen16998 := Call(__e, ShenFunc(symhd), V2701)

																					Var := gen16998
																					gen16999 := Call(__e, ShenFunc(symtl), V2701)

																					gen17000 := Call(__e, ShenFunc(symshen_4lazyderef), gen16999, V2783)

																					V2702 := gen17000
																					gen17011 := Call(__e, ShenFunc(symcons_2), V2702)

																					if True == gen17011 {
																						gen17001 := Call(__e, ShenFunc(symhd), V2702)

																						Val := gen17001
																						gen17002 := Call(__e, ShenFunc(symtl), V2702)

																						gen17003 := Call(__e, ShenFunc(symshen_4lazyderef), gen17002, V2783)

																						V2703 := gen17003
																						gen17010 := Call(__e, ShenFunc(sym_a), Nil, V2703)

																						if True == gen17010 {
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen17009 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen17008 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen17007 := MakeNative(func(__e Evaluator, __args ...Obj) {
																										gen17004 := Call(__e, ShenFunc(symcons), Var, Nil)

																										gen17005 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen17004)

																										gen17006 := MakeNative(func(__e Evaluator, __args ...Obj) {
																											__e.TailApply(ShenFunc(symshen_4th_d), Val, V2781, V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symshen_4th_d), gen17005, V2781, V2782, V2783, gen17006)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen17007)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symshen_4th_d), Var, MakeSymbol("symbol"), V2782, V2783, gen17008)

																								return

																							}, 0)
																							gen17015 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen17009)

																						} else {
																							gen17015 = False
																						}

																					} else {
																						gen17015 = False
																					}

																				} else {
																					gen17015 = False
																				}

																			} else {
																				gen17015 = False
																			}

																		} else {
																			gen17015 = False
																		}
																		Case := gen17015
																		gen17074 := Call(__e, ShenFunc(sym_a), Case, False)

																		if True == gen17074 {
																			gen17016 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																			NewHyp := gen17016
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17017 := MakeNative(func(__e Evaluator, __args ...Obj) {
																				__e.TailApply(ShenFunc(symshen_4th_d), V2780, V2781, NewHyp, V2783, V2784)

																				return
																			}, 0)
																			gen17018 := Call(__e, ShenFunc(symshen_4t_d_1hyps), V2782, NewHyp, V2783, gen17017)

																			Case := gen17018
																			gen17073 := Call(__e, ShenFunc(sym_a), Case, False)

																			if True == gen17073 {
																				gen17019 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																				V2704 := gen17019
																				gen17031 := Call(__e, ShenFunc(symcons_2), V2704)

																				var gen17032 Obj
																				if True == gen17031 {
																					gen17020 := Call(__e, ShenFunc(symhd), V2704)

																					gen17021 := Call(__e, ShenFunc(symshen_4lazyderef), gen17020, V2783)

																					V2705 := gen17021
																					gen17030 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), V2705)

																					if True == gen17030 {
																						gen17022 := Call(__e, ShenFunc(symtl), V2704)

																						gen17023 := Call(__e, ShenFunc(symshen_4lazyderef), gen17022, V2783)

																						V2706 := gen17023
																						gen17029 := Call(__e, ShenFunc(symcons_2), V2706)

																						if True == gen17029 {
																							gen17024 := Call(__e, ShenFunc(symhd), V2706)

																							F := gen17024
																							gen17025 := Call(__e, ShenFunc(symtl), V2706)

																							X := gen17025
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen17028 := MakeNative(func(__e Evaluator, __args ...Obj) {
																								gen17026 := Call(__e, ShenFunc(symcons), F, X)

																								gen17027 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen17026)

																								__e.TailApply(ShenFunc(symshen_4t_d_1def), gen17027, V2781, V2782, V2783, V2784)

																								return

																							}, 0)
																							gen17032 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen17028)

																						} else {
																							gen17032 = False
																						}

																					} else {
																						gen17032 = False
																					}

																				} else {
																					gen17032 = False
																				}
																				Case := gen17032
																				gen17072 := Call(__e, ShenFunc(sym_a), Case, False)

																				if True == gen17072 {
																					gen17033 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																					V2707 := gen17033
																					gen17041 := Call(__e, ShenFunc(symcons_2), V2707)

																					var gen17042 Obj
																					if True == gen17041 {
																						gen17034 := Call(__e, ShenFunc(symhd), V2707)

																						gen17035 := Call(__e, ShenFunc(symshen_4lazyderef), gen17034, V2783)

																						V2708 := gen17035
																						gen17040 := Call(__e, ShenFunc(sym_a), MakeSymbol("defmacro"), V2708)

																						if True == gen17040 {
																							gen17036 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																							V2709 := gen17036
																							gen17039 := Call(__e, ShenFunc(sym_a), MakeSymbol("unit"), V2709)

																							if True == gen17039 {
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen17042 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, V2784)

																							} else {
																								gen17038 := Call(__e, ShenFunc(symshen_4pvar_2), V2709)

																								if True == gen17038 {
																									Call(__e, ShenFunc(symshen_4bindv), V2709, MakeSymbol("unit"), V2783)
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen17037 := Call(__e, ShenFunc(symcut), Throwcontrol, V2783, V2784)

																									Result := gen17037
																									Call(__e, ShenFunc(symshen_4unbindv), V2709, V2783)
																									gen17042 = Result

																								} else {
																									gen17042 = False
																								}

																							}

																						} else {
																							gen17042 = False
																						}

																					} else {
																						gen17042 = False
																					}
																					Case := gen17042
																					gen17071 := Call(__e, ShenFunc(sym_a), Case, False)

																					if True == gen17071 {
																						gen17043 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																						V2710 := gen17043
																						gen17051 := Call(__e, ShenFunc(symcons_2), V2710)

																						var gen17052 Obj
																						if True == gen17051 {
																							gen17044 := Call(__e, ShenFunc(symhd), V2710)

																							gen17045 := Call(__e, ShenFunc(symshen_4lazyderef), gen17044, V2783)

																							V2711 := gen17045
																							gen17050 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.process-datatype"), V2711)

																							if True == gen17050 {
																								gen17046 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																								V2712 := gen17046
																								gen17049 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2712)

																								if True == gen17049 {
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen17052 = Call(__e, ShenFunc(symthaw), V2784)

																								} else {
																									gen17048 := Call(__e, ShenFunc(symshen_4pvar_2), V2712)

																									if True == gen17048 {
																										Call(__e, ShenFunc(symshen_4bindv), V2712, MakeSymbol("symbol"), V2783)
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen17047 := Call(__e, ShenFunc(symthaw), V2784)

																										Result := gen17047
																										Call(__e, ShenFunc(symshen_4unbindv), V2712, V2783)
																										gen17052 = Result

																									} else {
																										gen17052 = False
																									}

																								}

																							} else {
																								gen17052 = False
																							}

																						} else {
																							gen17052 = False
																						}
																						Case := gen17052
																						gen17070 := Call(__e, ShenFunc(sym_a), Case, False)

																						if True == gen17070 {
																							gen17053 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																							V2713 := gen17053
																							gen17061 := Call(__e, ShenFunc(symcons_2), V2713)

																							var gen17062 Obj
																							if True == gen17061 {
																								gen17054 := Call(__e, ShenFunc(symhd), V2713)

																								gen17055 := Call(__e, ShenFunc(symshen_4lazyderef), gen17054, V2783)

																								V2714 := gen17055
																								gen17060 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.synonyms-help"), V2714)

																								if True == gen17060 {
																									gen17056 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																									V2715 := gen17056
																									gen17059 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2715)

																									if True == gen17059 {
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen17062 = Call(__e, ShenFunc(symthaw), V2784)

																									} else {
																										gen17058 := Call(__e, ShenFunc(symshen_4pvar_2), V2715)

																										if True == gen17058 {
																											Call(__e, ShenFunc(symshen_4bindv), V2715, MakeSymbol("symbol"), V2783)
																											Call(__e, ShenFunc(symshen_4incinfs))
																											gen17057 := Call(__e, ShenFunc(symthaw), V2784)

																											Result := gen17057
																											Call(__e, ShenFunc(symshen_4unbindv), V2715, V2783)
																											gen17062 = Result

																										} else {
																											gen17062 = False
																										}

																									}

																								} else {
																									gen17062 = False
																								}

																							} else {
																								gen17062 = False
																							}
																							Case := gen17062
																							gen17069 := Call(__e, ShenFunc(sym_a), Case, False)

																							if True == gen17069 {
																								gen17063 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Datatypes := gen17063
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen17064 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

																								gen17068 := MakeNative(func(__e Evaluator, __args ...Obj) {
																									gen17065 := Call(__e, ShenFunc(symcons), V2781, Nil)

																									gen17066 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17065)

																									gen17067 := Call(__e, ShenFunc(symcons), V2780, gen17066)

																									__e.TailApply(ShenFunc(symshen_4udefs_d), gen17067, V2782, Datatypes, V2783, V2784)

																									return

																								}, 0)
																								gen17090 = Call(__e, ShenFunc(symbind), Datatypes, gen17064, V2783, gen17068)

																							} else {
																								gen17090 = Case
																							}

																						} else {
																							gen17090 = Case
																						}

																					} else {
																						gen17090 = Case
																					}

																				} else {
																					gen17090 = Case
																				}

																			} else {
																				gen17090 = Case
																			}

																		} else {
																			gen17090 = Case
																		}

																	} else {
																		gen17090 = Case
																	}

																} else {
																	gen17090 = Case
																}

															} else {
																gen17090 = Case
															}

														} else {
															gen17090 = Case
														}

													} else {
														gen17090 = Case
													}

												} else {
													gen17090 = Case
												}

											} else {
												gen17090 = Case
											}

										} else {
											gen17090 = Case
										}

									} else {
										gen17090 = Case
									}

								} else {
									gen17090 = Case
								}

							} else {
								gen17090 = Case
							}

						} else {
							gen17090 = Case
						}

					} else {
						gen17090 = Case
					}

				} else {
					gen17090 = Case
				}

			} else {
				gen17090 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen17090)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.th*"), gen17091)

		gen18206 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2789 := __args[0]
			_ = V2789
			V2790 := __args[1]
			_ = V2790
			V2791 := __args[2]
			_ = V2791
			V2792 := __args[3]
			_ = V2792
			gen17092 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

			V2526 := gen17092
			gen17420 := Call(__e, ShenFunc(symcons_2), V2526)

			var gen17421 Obj
			if True == gen17420 {
				gen17093 := Call(__e, ShenFunc(symhd), V2526)

				gen17094 := Call(__e, ShenFunc(symshen_4lazyderef), gen17093, V2791)

				V2527 := gen17094
				gen17419 := Call(__e, ShenFunc(symcons_2), V2527)

				if True == gen17419 {
					gen17095 := Call(__e, ShenFunc(symhd), V2527)

					gen17096 := Call(__e, ShenFunc(symshen_4lazyderef), gen17095, V2791)

					V2528 := gen17096
					gen17418 := Call(__e, ShenFunc(symcons_2), V2528)

					if True == gen17418 {
						gen17097 := Call(__e, ShenFunc(symhd), V2528)

						gen17098 := Call(__e, ShenFunc(symshen_4lazyderef), gen17097, V2791)

						V2529 := gen17098
						gen17417 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), V2529)

						if True == gen17417 {
							gen17099 := Call(__e, ShenFunc(symtl), V2528)

							gen17100 := Call(__e, ShenFunc(symshen_4lazyderef), gen17099, V2791)

							V2530 := gen17100
							gen17416 := Call(__e, ShenFunc(symcons_2), V2530)

							if True == gen17416 {
								gen17101 := Call(__e, ShenFunc(symhd), V2530)

								X := gen17101
								gen17102 := Call(__e, ShenFunc(symtl), V2530)

								gen17103 := Call(__e, ShenFunc(symshen_4lazyderef), gen17102, V2791)

								V2531 := gen17103
								gen17415 := Call(__e, ShenFunc(symcons_2), V2531)

								if True == gen17415 {
									gen17104 := Call(__e, ShenFunc(symhd), V2531)

									Y := gen17104
									gen17105 := Call(__e, ShenFunc(symtl), V2531)

									gen17106 := Call(__e, ShenFunc(symshen_4lazyderef), gen17105, V2791)

									V2532 := gen17106
									gen17414 := Call(__e, ShenFunc(sym_a), Nil, V2532)

									if True == gen17414 {
										gen17107 := Call(__e, ShenFunc(symtl), V2527)

										gen17108 := Call(__e, ShenFunc(symshen_4lazyderef), gen17107, V2791)

										V2533 := gen17108
										gen17413 := Call(__e, ShenFunc(symcons_2), V2533)

										if True == gen17413 {
											gen17109 := Call(__e, ShenFunc(symhd), V2533)

											gen17110 := Call(__e, ShenFunc(symshen_4lazyderef), gen17109, V2791)

											V2534 := gen17110
											gen17412 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2534)

											if True == gen17412 {
												gen17111 := Call(__e, ShenFunc(symtl), V2533)

												gen17112 := Call(__e, ShenFunc(symshen_4lazyderef), gen17111, V2791)

												V2535 := gen17112
												gen17411 := Call(__e, ShenFunc(symcons_2), V2535)

												if True == gen17411 {
													gen17113 := Call(__e, ShenFunc(symhd), V2535)

													gen17114 := Call(__e, ShenFunc(symshen_4lazyderef), gen17113, V2791)

													V2536 := gen17114
													gen17410 := Call(__e, ShenFunc(symcons_2), V2536)

													if True == gen17410 {
														gen17157 := Call(__e, ShenFunc(symhd), V2536)

														gen17158 := Call(__e, ShenFunc(symshen_4lazyderef), gen17157, V2791)

														V2537 := gen17158
														gen17409 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2537)

														if True == gen17409 {
															gen17285 := Call(__e, ShenFunc(symtl), V2536)

															gen17286 := Call(__e, ShenFunc(symshen_4lazyderef), gen17285, V2791)

															V2538 := gen17286
															gen17408 := Call(__e, ShenFunc(symcons_2), V2538)

															if True == gen17408 {
																gen17328 := Call(__e, ShenFunc(symhd), V2538)

																A := gen17328
																gen17329 := Call(__e, ShenFunc(symtl), V2538)

																gen17330 := Call(__e, ShenFunc(symshen_4lazyderef), gen17329, V2791)

																V2539 := gen17330
																gen17407 := Call(__e, ShenFunc(sym_a), Nil, V2539)

																if True == gen17407 {
																	gen17370 := Call(__e, ShenFunc(symtl), V2535)

																	gen17371 := Call(__e, ShenFunc(symshen_4lazyderef), gen17370, V2791)

																	V2540 := gen17371
																	gen17406 := Call(__e, ShenFunc(sym_a), Nil, V2540)

																	if True == gen17406 {
																		gen17390 := Call(__e, ShenFunc(symtl), V2526)

																		Hyp := gen17390
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen17391 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen17392 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17393 := Call(__e, ShenFunc(symcons), gen17392, Nil)

																		gen17394 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17393)

																		gen17395 := Call(__e, ShenFunc(symcons), gen17391, gen17394)

																		gen17396 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen17397 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17398 := Call(__e, ShenFunc(symcons), gen17397, Nil)

																		gen17399 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17398)

																		gen17400 := Call(__e, ShenFunc(symcons), gen17399, Nil)

																		gen17401 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17400)

																		gen17402 := Call(__e, ShenFunc(symcons), gen17396, gen17401)

																		gen17403 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen17404 := Call(__e, ShenFunc(symcons), gen17402, gen17403)

																		gen17405 := Call(__e, ShenFunc(symcons), gen17395, gen17404)

																		gen17421 = Call(__e, ShenFunc(symbind), V2790, gen17405, V2791, V2792)

																	} else {
																		gen17389 := Call(__e, ShenFunc(symshen_4pvar_2), V2540)

																		if True == gen17389 {
																			Call(__e, ShenFunc(symshen_4bindv), V2540, Nil, V2791)
																			gen17372 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen17372
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17373 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17374 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17375 := Call(__e, ShenFunc(symcons), gen17374, Nil)

																			gen17376 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17375)

																			gen17377 := Call(__e, ShenFunc(symcons), gen17373, gen17376)

																			gen17378 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17379 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17380 := Call(__e, ShenFunc(symcons), gen17379, Nil)

																			gen17381 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17380)

																			gen17382 := Call(__e, ShenFunc(symcons), gen17381, Nil)

																			gen17383 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17382)

																			gen17384 := Call(__e, ShenFunc(symcons), gen17378, gen17383)

																			gen17385 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17386 := Call(__e, ShenFunc(symcons), gen17384, gen17385)

																			gen17387 := Call(__e, ShenFunc(symcons), gen17377, gen17386)

																			gen17388 := Call(__e, ShenFunc(symbind), V2790, gen17387, V2791, V2792)

																			Result := gen17388
																			Call(__e, ShenFunc(symshen_4unbindv), V2540, V2791)
																			gen17421 = Result

																		} else {
																			gen17421 = False
																		}

																	}

																} else {
																	gen17369 := Call(__e, ShenFunc(symshen_4pvar_2), V2539)

																	if True == gen17369 {
																		Call(__e, ShenFunc(symshen_4bindv), V2539, Nil, V2791)
																		gen17331 := Call(__e, ShenFunc(symtl), V2535)

																		gen17332 := Call(__e, ShenFunc(symshen_4lazyderef), gen17331, V2791)

																		V2541 := gen17332
																		gen17367 := Call(__e, ShenFunc(sym_a), Nil, V2541)

																		var gen17368 Obj
																		if True == gen17367 {
																			gen17351 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen17351
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17352 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17353 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17354 := Call(__e, ShenFunc(symcons), gen17353, Nil)

																			gen17355 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17354)

																			gen17356 := Call(__e, ShenFunc(symcons), gen17352, gen17355)

																			gen17357 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17358 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17359 := Call(__e, ShenFunc(symcons), gen17358, Nil)

																			gen17360 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17359)

																			gen17361 := Call(__e, ShenFunc(symcons), gen17360, Nil)

																			gen17362 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17361)

																			gen17363 := Call(__e, ShenFunc(symcons), gen17357, gen17362)

																			gen17364 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17365 := Call(__e, ShenFunc(symcons), gen17363, gen17364)

																			gen17366 := Call(__e, ShenFunc(symcons), gen17356, gen17365)

																			gen17368 = Call(__e, ShenFunc(symbind), V2790, gen17366, V2791, V2792)

																		} else {
																			gen17350 := Call(__e, ShenFunc(symshen_4pvar_2), V2541)

																			if True == gen17350 {
																				Call(__e, ShenFunc(symshen_4bindv), V2541, Nil, V2791)
																				gen17333 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen17333
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17334 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17335 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17336 := Call(__e, ShenFunc(symcons), gen17335, Nil)

																				gen17337 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17336)

																				gen17338 := Call(__e, ShenFunc(symcons), gen17334, gen17337)

																				gen17339 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17340 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17341 := Call(__e, ShenFunc(symcons), gen17340, Nil)

																				gen17342 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17341)

																				gen17343 := Call(__e, ShenFunc(symcons), gen17342, Nil)

																				gen17344 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17343)

																				gen17345 := Call(__e, ShenFunc(symcons), gen17339, gen17344)

																				gen17346 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17347 := Call(__e, ShenFunc(symcons), gen17345, gen17346)

																				gen17348 := Call(__e, ShenFunc(symcons), gen17338, gen17347)

																				gen17349 := Call(__e, ShenFunc(symbind), V2790, gen17348, V2791, V2792)

																				Result := gen17349
																				Call(__e, ShenFunc(symshen_4unbindv), V2541, V2791)
																				gen17368 = Result

																			} else {
																				gen17368 = False
																			}

																		}
																		Result := gen17368
																		Call(__e, ShenFunc(symshen_4unbindv), V2539, V2791)
																		gen17421 = Result

																	} else {
																		gen17421 = False
																	}

																}

															} else {
																gen17327 := Call(__e, ShenFunc(symshen_4pvar_2), V2538)

																if True == gen17327 {
																	gen17287 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	A := gen17287
																	gen17288 := Call(__e, ShenFunc(symcons), A, Nil)

																	Call(__e, ShenFunc(symshen_4bindv), V2538, gen17288, V2791)

																	gen17289 := Call(__e, ShenFunc(symtl), V2535)

																	gen17290 := Call(__e, ShenFunc(symshen_4lazyderef), gen17289, V2791)

																	V2542 := gen17290
																	gen17325 := Call(__e, ShenFunc(sym_a), Nil, V2542)

																	var gen17326 Obj
																	if True == gen17325 {
																		gen17309 := Call(__e, ShenFunc(symtl), V2526)

																		Hyp := gen17309
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen17310 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen17311 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17312 := Call(__e, ShenFunc(symcons), gen17311, Nil)

																		gen17313 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17312)

																		gen17314 := Call(__e, ShenFunc(symcons), gen17310, gen17313)

																		gen17315 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen17316 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17317 := Call(__e, ShenFunc(symcons), gen17316, Nil)

																		gen17318 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17317)

																		gen17319 := Call(__e, ShenFunc(symcons), gen17318, Nil)

																		gen17320 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17319)

																		gen17321 := Call(__e, ShenFunc(symcons), gen17315, gen17320)

																		gen17322 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen17323 := Call(__e, ShenFunc(symcons), gen17321, gen17322)

																		gen17324 := Call(__e, ShenFunc(symcons), gen17314, gen17323)

																		gen17326 = Call(__e, ShenFunc(symbind), V2790, gen17324, V2791, V2792)

																	} else {
																		gen17308 := Call(__e, ShenFunc(symshen_4pvar_2), V2542)

																		if True == gen17308 {
																			Call(__e, ShenFunc(symshen_4bindv), V2542, Nil, V2791)
																			gen17291 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen17291
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17292 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17293 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17294 := Call(__e, ShenFunc(symcons), gen17293, Nil)

																			gen17295 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17294)

																			gen17296 := Call(__e, ShenFunc(symcons), gen17292, gen17295)

																			gen17297 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17298 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17299 := Call(__e, ShenFunc(symcons), gen17298, Nil)

																			gen17300 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17299)

																			gen17301 := Call(__e, ShenFunc(symcons), gen17300, Nil)

																			gen17302 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17301)

																			gen17303 := Call(__e, ShenFunc(symcons), gen17297, gen17302)

																			gen17304 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17305 := Call(__e, ShenFunc(symcons), gen17303, gen17304)

																			gen17306 := Call(__e, ShenFunc(symcons), gen17296, gen17305)

																			gen17307 := Call(__e, ShenFunc(symbind), V2790, gen17306, V2791, V2792)

																			Result := gen17307
																			Call(__e, ShenFunc(symshen_4unbindv), V2542, V2791)
																			gen17326 = Result

																		} else {
																			gen17326 = False
																		}

																	}
																	Result := gen17326
																	Call(__e, ShenFunc(symshen_4unbindv), V2538, V2791)
																	gen17421 = Result

																} else {
																	gen17421 = False
																}

															}

														} else {
															gen17284 := Call(__e, ShenFunc(symshen_4pvar_2), V2537)

															if True == gen17284 {
																Call(__e, ShenFunc(symshen_4bindv), V2537, MakeSymbol("list"), V2791)
																gen17159 := Call(__e, ShenFunc(symtl), V2536)

																gen17160 := Call(__e, ShenFunc(symshen_4lazyderef), gen17159, V2791)

																V2543 := gen17160
																gen17282 := Call(__e, ShenFunc(symcons_2), V2543)

																var gen17283 Obj
																if True == gen17282 {
																	gen17202 := Call(__e, ShenFunc(symhd), V2543)

																	A := gen17202
																	gen17203 := Call(__e, ShenFunc(symtl), V2543)

																	gen17204 := Call(__e, ShenFunc(symshen_4lazyderef), gen17203, V2791)

																	V2544 := gen17204
																	gen17281 := Call(__e, ShenFunc(sym_a), Nil, V2544)

																	if True == gen17281 {
																		gen17244 := Call(__e, ShenFunc(symtl), V2535)

																		gen17245 := Call(__e, ShenFunc(symshen_4lazyderef), gen17244, V2791)

																		V2545 := gen17245
																		gen17280 := Call(__e, ShenFunc(sym_a), Nil, V2545)

																		if True == gen17280 {
																			gen17264 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen17264
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17265 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17266 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17267 := Call(__e, ShenFunc(symcons), gen17266, Nil)

																			gen17268 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17267)

																			gen17269 := Call(__e, ShenFunc(symcons), gen17265, gen17268)

																			gen17270 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17271 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17272 := Call(__e, ShenFunc(symcons), gen17271, Nil)

																			gen17273 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17272)

																			gen17274 := Call(__e, ShenFunc(symcons), gen17273, Nil)

																			gen17275 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17274)

																			gen17276 := Call(__e, ShenFunc(symcons), gen17270, gen17275)

																			gen17277 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17278 := Call(__e, ShenFunc(symcons), gen17276, gen17277)

																			gen17279 := Call(__e, ShenFunc(symcons), gen17269, gen17278)

																			gen17283 = Call(__e, ShenFunc(symbind), V2790, gen17279, V2791, V2792)

																		} else {
																			gen17263 := Call(__e, ShenFunc(symshen_4pvar_2), V2545)

																			if True == gen17263 {
																				Call(__e, ShenFunc(symshen_4bindv), V2545, Nil, V2791)
																				gen17246 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen17246
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17247 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17248 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17249 := Call(__e, ShenFunc(symcons), gen17248, Nil)

																				gen17250 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17249)

																				gen17251 := Call(__e, ShenFunc(symcons), gen17247, gen17250)

																				gen17252 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17253 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17254 := Call(__e, ShenFunc(symcons), gen17253, Nil)

																				gen17255 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17254)

																				gen17256 := Call(__e, ShenFunc(symcons), gen17255, Nil)

																				gen17257 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17256)

																				gen17258 := Call(__e, ShenFunc(symcons), gen17252, gen17257)

																				gen17259 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17260 := Call(__e, ShenFunc(symcons), gen17258, gen17259)

																				gen17261 := Call(__e, ShenFunc(symcons), gen17251, gen17260)

																				gen17262 := Call(__e, ShenFunc(symbind), V2790, gen17261, V2791, V2792)

																				Result := gen17262
																				Call(__e, ShenFunc(symshen_4unbindv), V2545, V2791)
																				gen17283 = Result

																			} else {
																				gen17283 = False
																			}

																		}

																	} else {
																		gen17243 := Call(__e, ShenFunc(symshen_4pvar_2), V2544)

																		if True == gen17243 {
																			Call(__e, ShenFunc(symshen_4bindv), V2544, Nil, V2791)
																			gen17205 := Call(__e, ShenFunc(symtl), V2535)

																			gen17206 := Call(__e, ShenFunc(symshen_4lazyderef), gen17205, V2791)

																			V2546 := gen17206
																			gen17241 := Call(__e, ShenFunc(sym_a), Nil, V2546)

																			var gen17242 Obj
																			if True == gen17241 {
																				gen17225 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen17225
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17226 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17227 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17228 := Call(__e, ShenFunc(symcons), gen17227, Nil)

																				gen17229 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17228)

																				gen17230 := Call(__e, ShenFunc(symcons), gen17226, gen17229)

																				gen17231 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17232 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17233 := Call(__e, ShenFunc(symcons), gen17232, Nil)

																				gen17234 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17233)

																				gen17235 := Call(__e, ShenFunc(symcons), gen17234, Nil)

																				gen17236 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17235)

																				gen17237 := Call(__e, ShenFunc(symcons), gen17231, gen17236)

																				gen17238 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17239 := Call(__e, ShenFunc(symcons), gen17237, gen17238)

																				gen17240 := Call(__e, ShenFunc(symcons), gen17230, gen17239)

																				gen17242 = Call(__e, ShenFunc(symbind), V2790, gen17240, V2791, V2792)

																			} else {
																				gen17224 := Call(__e, ShenFunc(symshen_4pvar_2), V2546)

																				if True == gen17224 {
																					Call(__e, ShenFunc(symshen_4bindv), V2546, Nil, V2791)
																					gen17207 := Call(__e, ShenFunc(symtl), V2526)

																					Hyp := gen17207
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17208 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17209 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17210 := Call(__e, ShenFunc(symcons), gen17209, Nil)

																					gen17211 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17210)

																					gen17212 := Call(__e, ShenFunc(symcons), gen17208, gen17211)

																					gen17213 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17214 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17215 := Call(__e, ShenFunc(symcons), gen17214, Nil)

																					gen17216 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17215)

																					gen17217 := Call(__e, ShenFunc(symcons), gen17216, Nil)

																					gen17218 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17217)

																					gen17219 := Call(__e, ShenFunc(symcons), gen17213, gen17218)

																					gen17220 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17221 := Call(__e, ShenFunc(symcons), gen17219, gen17220)

																					gen17222 := Call(__e, ShenFunc(symcons), gen17212, gen17221)

																					gen17223 := Call(__e, ShenFunc(symbind), V2790, gen17222, V2791, V2792)

																					Result := gen17223
																					Call(__e, ShenFunc(symshen_4unbindv), V2546, V2791)
																					gen17242 = Result

																				} else {
																					gen17242 = False
																				}

																			}
																			Result := gen17242
																			Call(__e, ShenFunc(symshen_4unbindv), V2544, V2791)
																			gen17283 = Result

																		} else {
																			gen17283 = False
																		}

																	}

																} else {
																	gen17201 := Call(__e, ShenFunc(symshen_4pvar_2), V2543)

																	if True == gen17201 {
																		gen17161 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																		A := gen17161
																		gen17162 := Call(__e, ShenFunc(symcons), A, Nil)

																		Call(__e, ShenFunc(symshen_4bindv), V2543, gen17162, V2791)

																		gen17163 := Call(__e, ShenFunc(symtl), V2535)

																		gen17164 := Call(__e, ShenFunc(symshen_4lazyderef), gen17163, V2791)

																		V2547 := gen17164
																		gen17199 := Call(__e, ShenFunc(sym_a), Nil, V2547)

																		var gen17200 Obj
																		if True == gen17199 {
																			gen17183 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen17183
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17184 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17185 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17186 := Call(__e, ShenFunc(symcons), gen17185, Nil)

																			gen17187 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17186)

																			gen17188 := Call(__e, ShenFunc(symcons), gen17184, gen17187)

																			gen17189 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17190 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17191 := Call(__e, ShenFunc(symcons), gen17190, Nil)

																			gen17192 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17191)

																			gen17193 := Call(__e, ShenFunc(symcons), gen17192, Nil)

																			gen17194 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17193)

																			gen17195 := Call(__e, ShenFunc(symcons), gen17189, gen17194)

																			gen17196 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17197 := Call(__e, ShenFunc(symcons), gen17195, gen17196)

																			gen17198 := Call(__e, ShenFunc(symcons), gen17188, gen17197)

																			gen17200 = Call(__e, ShenFunc(symbind), V2790, gen17198, V2791, V2792)

																		} else {
																			gen17182 := Call(__e, ShenFunc(symshen_4pvar_2), V2547)

																			if True == gen17182 {
																				Call(__e, ShenFunc(symshen_4bindv), V2547, Nil, V2791)
																				gen17165 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen17165
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17166 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17167 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17168 := Call(__e, ShenFunc(symcons), gen17167, Nil)

																				gen17169 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17168)

																				gen17170 := Call(__e, ShenFunc(symcons), gen17166, gen17169)

																				gen17171 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17172 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17173 := Call(__e, ShenFunc(symcons), gen17172, Nil)

																				gen17174 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17173)

																				gen17175 := Call(__e, ShenFunc(symcons), gen17174, Nil)

																				gen17176 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17175)

																				gen17177 := Call(__e, ShenFunc(symcons), gen17171, gen17176)

																				gen17178 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17179 := Call(__e, ShenFunc(symcons), gen17177, gen17178)

																				gen17180 := Call(__e, ShenFunc(symcons), gen17170, gen17179)

																				gen17181 := Call(__e, ShenFunc(symbind), V2790, gen17180, V2791, V2792)

																				Result := gen17181
																				Call(__e, ShenFunc(symshen_4unbindv), V2547, V2791)
																				gen17200 = Result

																			} else {
																				gen17200 = False
																			}

																		}
																		Result := gen17200
																		Call(__e, ShenFunc(symshen_4unbindv), V2543, V2791)
																		gen17283 = Result

																	} else {
																		gen17283 = False
																	}

																}
																Result := gen17283
																Call(__e, ShenFunc(symshen_4unbindv), V2537, V2791)
																gen17421 = Result

															} else {
																gen17421 = False
															}

														}

													} else {
														gen17156 := Call(__e, ShenFunc(symshen_4pvar_2), V2536)

														if True == gen17156 {
															gen17115 := Call(__e, ShenFunc(symshen_4newpv), V2791)

															A := gen17115
															gen17116 := Call(__e, ShenFunc(symcons), A, Nil)

															gen17117 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17116)

															Call(__e, ShenFunc(symshen_4bindv), V2536, gen17117, V2791)

															gen17118 := Call(__e, ShenFunc(symtl), V2535)

															gen17119 := Call(__e, ShenFunc(symshen_4lazyderef), gen17118, V2791)

															V2548 := gen17119
															gen17154 := Call(__e, ShenFunc(sym_a), Nil, V2548)

															var gen17155 Obj
															if True == gen17154 {
																gen17138 := Call(__e, ShenFunc(symtl), V2526)

																Hyp := gen17138
																Call(__e, ShenFunc(symshen_4incinfs))
																gen17139 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																gen17140 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																gen17141 := Call(__e, ShenFunc(symcons), gen17140, Nil)

																gen17142 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17141)

																gen17143 := Call(__e, ShenFunc(symcons), gen17139, gen17142)

																gen17144 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																gen17145 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																gen17146 := Call(__e, ShenFunc(symcons), gen17145, Nil)

																gen17147 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17146)

																gen17148 := Call(__e, ShenFunc(symcons), gen17147, Nil)

																gen17149 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17148)

																gen17150 := Call(__e, ShenFunc(symcons), gen17144, gen17149)

																gen17151 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																gen17152 := Call(__e, ShenFunc(symcons), gen17150, gen17151)

																gen17153 := Call(__e, ShenFunc(symcons), gen17143, gen17152)

																gen17155 = Call(__e, ShenFunc(symbind), V2790, gen17153, V2791, V2792)

															} else {
																gen17137 := Call(__e, ShenFunc(symshen_4pvar_2), V2548)

																if True == gen17137 {
																	Call(__e, ShenFunc(symshen_4bindv), V2548, Nil, V2791)
																	gen17120 := Call(__e, ShenFunc(symtl), V2526)

																	Hyp := gen17120
																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen17121 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																	gen17122 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen17123 := Call(__e, ShenFunc(symcons), gen17122, Nil)

																	gen17124 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17123)

																	gen17125 := Call(__e, ShenFunc(symcons), gen17121, gen17124)

																	gen17126 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																	gen17127 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen17128 := Call(__e, ShenFunc(symcons), gen17127, Nil)

																	gen17129 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen17128)

																	gen17130 := Call(__e, ShenFunc(symcons), gen17129, Nil)

																	gen17131 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17130)

																	gen17132 := Call(__e, ShenFunc(symcons), gen17126, gen17131)

																	gen17133 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																	gen17134 := Call(__e, ShenFunc(symcons), gen17132, gen17133)

																	gen17135 := Call(__e, ShenFunc(symcons), gen17125, gen17134)

																	gen17136 := Call(__e, ShenFunc(symbind), V2790, gen17135, V2791, V2792)

																	Result := gen17136
																	Call(__e, ShenFunc(symshen_4unbindv), V2548, V2791)
																	gen17155 = Result

																} else {
																	gen17155 = False
																}

															}
															Result := gen17155
															Call(__e, ShenFunc(symshen_4unbindv), V2536, V2791)
															gen17421 = Result

														} else {
															gen17421 = False
														}

													}

												} else {
													gen17421 = False
												}

											} else {
												gen17421 = False
											}

										} else {
											gen17421 = False
										}

									} else {
										gen17421 = False
									}

								} else {
									gen17421 = False
								}

							} else {
								gen17421 = False
							}

						} else {
							gen17421 = False
						}

					} else {
						gen17421 = False
					}

				} else {
					gen17421 = False
				}

			} else {
				gen17421 = False
			}
			Case := gen17421
			gen18205 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18205 {
				gen17422 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

				V2549 := gen17422
				gen17766 := Call(__e, ShenFunc(symcons_2), V2549)

				var gen17767 Obj
				if True == gen17766 {
					gen17423 := Call(__e, ShenFunc(symhd), V2549)

					gen17424 := Call(__e, ShenFunc(symshen_4lazyderef), gen17423, V2791)

					V2550 := gen17424
					gen17765 := Call(__e, ShenFunc(symcons_2), V2550)

					if True == gen17765 {
						gen17425 := Call(__e, ShenFunc(symhd), V2550)

						gen17426 := Call(__e, ShenFunc(symshen_4lazyderef), gen17425, V2791)

						V2551 := gen17426
						gen17764 := Call(__e, ShenFunc(symcons_2), V2551)

						if True == gen17764 {
							gen17427 := Call(__e, ShenFunc(symhd), V2551)

							gen17428 := Call(__e, ShenFunc(symshen_4lazyderef), gen17427, V2791)

							V2552 := gen17428
							gen17763 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), V2552)

							if True == gen17763 {
								gen17429 := Call(__e, ShenFunc(symtl), V2551)

								gen17430 := Call(__e, ShenFunc(symshen_4lazyderef), gen17429, V2791)

								V2553 := gen17430
								gen17762 := Call(__e, ShenFunc(symcons_2), V2553)

								if True == gen17762 {
									gen17431 := Call(__e, ShenFunc(symhd), V2553)

									X := gen17431
									gen17432 := Call(__e, ShenFunc(symtl), V2553)

									gen17433 := Call(__e, ShenFunc(symshen_4lazyderef), gen17432, V2791)

									V2554 := gen17433
									gen17761 := Call(__e, ShenFunc(symcons_2), V2554)

									if True == gen17761 {
										gen17434 := Call(__e, ShenFunc(symhd), V2554)

										Y := gen17434
										gen17435 := Call(__e, ShenFunc(symtl), V2554)

										gen17436 := Call(__e, ShenFunc(symshen_4lazyderef), gen17435, V2791)

										V2555 := gen17436
										gen17760 := Call(__e, ShenFunc(sym_a), Nil, V2555)

										if True == gen17760 {
											gen17437 := Call(__e, ShenFunc(symtl), V2550)

											gen17438 := Call(__e, ShenFunc(symshen_4lazyderef), gen17437, V2791)

											V2556 := gen17438
											gen17759 := Call(__e, ShenFunc(symcons_2), V2556)

											if True == gen17759 {
												gen17439 := Call(__e, ShenFunc(symhd), V2556)

												gen17440 := Call(__e, ShenFunc(symshen_4lazyderef), gen17439, V2791)

												V2557 := gen17440
												gen17758 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2557)

												if True == gen17758 {
													gen17441 := Call(__e, ShenFunc(symtl), V2556)

													gen17442 := Call(__e, ShenFunc(symshen_4lazyderef), gen17441, V2791)

													V2558 := gen17442
													gen17757 := Call(__e, ShenFunc(symcons_2), V2558)

													if True == gen17757 {
														gen17443 := Call(__e, ShenFunc(symhd), V2558)

														gen17444 := Call(__e, ShenFunc(symshen_4lazyderef), gen17443, V2791)

														V2559 := gen17444
														gen17756 := Call(__e, ShenFunc(symcons_2), V2559)

														if True == gen17756 {
															gen17485 := Call(__e, ShenFunc(symhd), V2559)

															A := gen17485
															gen17486 := Call(__e, ShenFunc(symtl), V2559)

															gen17487 := Call(__e, ShenFunc(symshen_4lazyderef), gen17486, V2791)

															V2560 := gen17487
															gen17755 := Call(__e, ShenFunc(symcons_2), V2560)

															if True == gen17755 {
																gen17526 := Call(__e, ShenFunc(symhd), V2560)

																gen17527 := Call(__e, ShenFunc(symshen_4lazyderef), gen17526, V2791)

																V2561 := gen17527
																gen17754 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), V2561)

																if True == gen17754 {
																	gen17642 := Call(__e, ShenFunc(symtl), V2560)

																	gen17643 := Call(__e, ShenFunc(symshen_4lazyderef), gen17642, V2791)

																	V2562 := gen17643
																	gen17753 := Call(__e, ShenFunc(symcons_2), V2562)

																	if True == gen17753 {
																		gen17681 := Call(__e, ShenFunc(symhd), V2562)

																		B := gen17681
																		gen17682 := Call(__e, ShenFunc(symtl), V2562)

																		gen17683 := Call(__e, ShenFunc(symshen_4lazyderef), gen17682, V2791)

																		V2563 := gen17683
																		gen17752 := Call(__e, ShenFunc(sym_a), Nil, V2563)

																		if True == gen17752 {
																			gen17719 := Call(__e, ShenFunc(symtl), V2558)

																			gen17720 := Call(__e, ShenFunc(symshen_4lazyderef), gen17719, V2791)

																			V2564 := gen17720
																			gen17751 := Call(__e, ShenFunc(sym_a), Nil, V2564)

																			if True == gen17751 {
																				gen17737 := Call(__e, ShenFunc(symtl), V2549)

																				Hyp := gen17737
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17738 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17739 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17740 := Call(__e, ShenFunc(symcons), gen17739, Nil)

																				gen17741 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17740)

																				gen17742 := Call(__e, ShenFunc(symcons), gen17738, gen17741)

																				gen17743 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17744 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																				gen17745 := Call(__e, ShenFunc(symcons), gen17744, Nil)

																				gen17746 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17745)

																				gen17747 := Call(__e, ShenFunc(symcons), gen17743, gen17746)

																				gen17748 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17749 := Call(__e, ShenFunc(symcons), gen17747, gen17748)

																				gen17750 := Call(__e, ShenFunc(symcons), gen17742, gen17749)

																				gen17767 = Call(__e, ShenFunc(symbind), V2790, gen17750, V2791, V2792)

																			} else {
																				gen17736 := Call(__e, ShenFunc(symshen_4pvar_2), V2564)

																				if True == gen17736 {
																					Call(__e, ShenFunc(symshen_4bindv), V2564, Nil, V2791)
																					gen17721 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen17721
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17722 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17723 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17724 := Call(__e, ShenFunc(symcons), gen17723, Nil)

																					gen17725 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17724)

																					gen17726 := Call(__e, ShenFunc(symcons), gen17722, gen17725)

																					gen17727 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17728 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen17729 := Call(__e, ShenFunc(symcons), gen17728, Nil)

																					gen17730 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17729)

																					gen17731 := Call(__e, ShenFunc(symcons), gen17727, gen17730)

																					gen17732 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17733 := Call(__e, ShenFunc(symcons), gen17731, gen17732)

																					gen17734 := Call(__e, ShenFunc(symcons), gen17726, gen17733)

																					gen17735 := Call(__e, ShenFunc(symbind), V2790, gen17734, V2791, V2792)

																					Result := gen17735
																					Call(__e, ShenFunc(symshen_4unbindv), V2564, V2791)
																					gen17767 = Result

																				} else {
																					gen17767 = False
																				}

																			}

																		} else {
																			gen17718 := Call(__e, ShenFunc(symshen_4pvar_2), V2563)

																			if True == gen17718 {
																				Call(__e, ShenFunc(symshen_4bindv), V2563, Nil, V2791)
																				gen17684 := Call(__e, ShenFunc(symtl), V2558)

																				gen17685 := Call(__e, ShenFunc(symshen_4lazyderef), gen17684, V2791)

																				V2565 := gen17685
																				gen17716 := Call(__e, ShenFunc(sym_a), Nil, V2565)

																				var gen17717 Obj
																				if True == gen17716 {
																					gen17702 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen17702
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17703 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17704 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17705 := Call(__e, ShenFunc(symcons), gen17704, Nil)

																					gen17706 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17705)

																					gen17707 := Call(__e, ShenFunc(symcons), gen17703, gen17706)

																					gen17708 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17709 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen17710 := Call(__e, ShenFunc(symcons), gen17709, Nil)

																					gen17711 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17710)

																					gen17712 := Call(__e, ShenFunc(symcons), gen17708, gen17711)

																					gen17713 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17714 := Call(__e, ShenFunc(symcons), gen17712, gen17713)

																					gen17715 := Call(__e, ShenFunc(symcons), gen17707, gen17714)

																					gen17717 = Call(__e, ShenFunc(symbind), V2790, gen17715, V2791, V2792)

																				} else {
																					gen17701 := Call(__e, ShenFunc(symshen_4pvar_2), V2565)

																					if True == gen17701 {
																						Call(__e, ShenFunc(symshen_4bindv), V2565, Nil, V2791)
																						gen17686 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen17686
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17687 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17688 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17689 := Call(__e, ShenFunc(symcons), gen17688, Nil)

																						gen17690 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17689)

																						gen17691 := Call(__e, ShenFunc(symcons), gen17687, gen17690)

																						gen17692 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17693 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen17694 := Call(__e, ShenFunc(symcons), gen17693, Nil)

																						gen17695 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17694)

																						gen17696 := Call(__e, ShenFunc(symcons), gen17692, gen17695)

																						gen17697 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17698 := Call(__e, ShenFunc(symcons), gen17696, gen17697)

																						gen17699 := Call(__e, ShenFunc(symcons), gen17691, gen17698)

																						gen17700 := Call(__e, ShenFunc(symbind), V2790, gen17699, V2791, V2792)

																						Result := gen17700
																						Call(__e, ShenFunc(symshen_4unbindv), V2565, V2791)
																						gen17717 = Result

																					} else {
																						gen17717 = False
																					}

																				}
																				Result := gen17717
																				Call(__e, ShenFunc(symshen_4unbindv), V2563, V2791)
																				gen17767 = Result

																			} else {
																				gen17767 = False
																			}

																		}

																	} else {
																		gen17680 := Call(__e, ShenFunc(symshen_4pvar_2), V2562)

																		if True == gen17680 {
																			gen17644 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																			B := gen17644
																			gen17645 := Call(__e, ShenFunc(symcons), B, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2562, gen17645, V2791)

																			gen17646 := Call(__e, ShenFunc(symtl), V2558)

																			gen17647 := Call(__e, ShenFunc(symshen_4lazyderef), gen17646, V2791)

																			V2566 := gen17647
																			gen17678 := Call(__e, ShenFunc(sym_a), Nil, V2566)

																			var gen17679 Obj
																			if True == gen17678 {
																				gen17664 := Call(__e, ShenFunc(symtl), V2549)

																				Hyp := gen17664
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17665 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17666 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17667 := Call(__e, ShenFunc(symcons), gen17666, Nil)

																				gen17668 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17667)

																				gen17669 := Call(__e, ShenFunc(symcons), gen17665, gen17668)

																				gen17670 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17671 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																				gen17672 := Call(__e, ShenFunc(symcons), gen17671, Nil)

																				gen17673 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17672)

																				gen17674 := Call(__e, ShenFunc(symcons), gen17670, gen17673)

																				gen17675 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17676 := Call(__e, ShenFunc(symcons), gen17674, gen17675)

																				gen17677 := Call(__e, ShenFunc(symcons), gen17669, gen17676)

																				gen17679 = Call(__e, ShenFunc(symbind), V2790, gen17677, V2791, V2792)

																			} else {
																				gen17663 := Call(__e, ShenFunc(symshen_4pvar_2), V2566)

																				if True == gen17663 {
																					Call(__e, ShenFunc(symshen_4bindv), V2566, Nil, V2791)
																					gen17648 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen17648
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17649 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17650 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17651 := Call(__e, ShenFunc(symcons), gen17650, Nil)

																					gen17652 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17651)

																					gen17653 := Call(__e, ShenFunc(symcons), gen17649, gen17652)

																					gen17654 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17655 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen17656 := Call(__e, ShenFunc(symcons), gen17655, Nil)

																					gen17657 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17656)

																					gen17658 := Call(__e, ShenFunc(symcons), gen17654, gen17657)

																					gen17659 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17660 := Call(__e, ShenFunc(symcons), gen17658, gen17659)

																					gen17661 := Call(__e, ShenFunc(symcons), gen17653, gen17660)

																					gen17662 := Call(__e, ShenFunc(symbind), V2790, gen17661, V2791, V2792)

																					Result := gen17662
																					Call(__e, ShenFunc(symshen_4unbindv), V2566, V2791)
																					gen17679 = Result

																				} else {
																					gen17679 = False
																				}

																			}
																			Result := gen17679
																			Call(__e, ShenFunc(symshen_4unbindv), V2562, V2791)
																			gen17767 = Result

																		} else {
																			gen17767 = False
																		}

																	}

																} else {
																	gen17641 := Call(__e, ShenFunc(symshen_4pvar_2), V2561)

																	if True == gen17641 {
																		Call(__e, ShenFunc(symshen_4bindv), V2561, MakeSymbol("*"), V2791)
																		gen17528 := Call(__e, ShenFunc(symtl), V2560)

																		gen17529 := Call(__e, ShenFunc(symshen_4lazyderef), gen17528, V2791)

																		V2567 := gen17529
																		gen17639 := Call(__e, ShenFunc(symcons_2), V2567)

																		var gen17640 Obj
																		if True == gen17639 {
																			gen17567 := Call(__e, ShenFunc(symhd), V2567)

																			B := gen17567
																			gen17568 := Call(__e, ShenFunc(symtl), V2567)

																			gen17569 := Call(__e, ShenFunc(symshen_4lazyderef), gen17568, V2791)

																			V2568 := gen17569
																			gen17638 := Call(__e, ShenFunc(sym_a), Nil, V2568)

																			if True == gen17638 {
																				gen17605 := Call(__e, ShenFunc(symtl), V2558)

																				gen17606 := Call(__e, ShenFunc(symshen_4lazyderef), gen17605, V2791)

																				V2569 := gen17606
																				gen17637 := Call(__e, ShenFunc(sym_a), Nil, V2569)

																				if True == gen17637 {
																					gen17623 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen17623
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17624 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17625 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17626 := Call(__e, ShenFunc(symcons), gen17625, Nil)

																					gen17627 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17626)

																					gen17628 := Call(__e, ShenFunc(symcons), gen17624, gen17627)

																					gen17629 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17630 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen17631 := Call(__e, ShenFunc(symcons), gen17630, Nil)

																					gen17632 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17631)

																					gen17633 := Call(__e, ShenFunc(symcons), gen17629, gen17632)

																					gen17634 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17635 := Call(__e, ShenFunc(symcons), gen17633, gen17634)

																					gen17636 := Call(__e, ShenFunc(symcons), gen17628, gen17635)

																					gen17640 = Call(__e, ShenFunc(symbind), V2790, gen17636, V2791, V2792)

																				} else {
																					gen17622 := Call(__e, ShenFunc(symshen_4pvar_2), V2569)

																					if True == gen17622 {
																						Call(__e, ShenFunc(symshen_4bindv), V2569, Nil, V2791)
																						gen17607 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen17607
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17608 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17609 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17610 := Call(__e, ShenFunc(symcons), gen17609, Nil)

																						gen17611 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17610)

																						gen17612 := Call(__e, ShenFunc(symcons), gen17608, gen17611)

																						gen17613 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17614 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen17615 := Call(__e, ShenFunc(symcons), gen17614, Nil)

																						gen17616 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17615)

																						gen17617 := Call(__e, ShenFunc(symcons), gen17613, gen17616)

																						gen17618 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17619 := Call(__e, ShenFunc(symcons), gen17617, gen17618)

																						gen17620 := Call(__e, ShenFunc(symcons), gen17612, gen17619)

																						gen17621 := Call(__e, ShenFunc(symbind), V2790, gen17620, V2791, V2792)

																						Result := gen17621
																						Call(__e, ShenFunc(symshen_4unbindv), V2569, V2791)
																						gen17640 = Result

																					} else {
																						gen17640 = False
																					}

																				}

																			} else {
																				gen17604 := Call(__e, ShenFunc(symshen_4pvar_2), V2568)

																				if True == gen17604 {
																					Call(__e, ShenFunc(symshen_4bindv), V2568, Nil, V2791)
																					gen17570 := Call(__e, ShenFunc(symtl), V2558)

																					gen17571 := Call(__e, ShenFunc(symshen_4lazyderef), gen17570, V2791)

																					V2570 := gen17571
																					gen17602 := Call(__e, ShenFunc(sym_a), Nil, V2570)

																					var gen17603 Obj
																					if True == gen17602 {
																						gen17588 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen17588
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17589 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17590 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17591 := Call(__e, ShenFunc(symcons), gen17590, Nil)

																						gen17592 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17591)

																						gen17593 := Call(__e, ShenFunc(symcons), gen17589, gen17592)

																						gen17594 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17595 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen17596 := Call(__e, ShenFunc(symcons), gen17595, Nil)

																						gen17597 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17596)

																						gen17598 := Call(__e, ShenFunc(symcons), gen17594, gen17597)

																						gen17599 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17600 := Call(__e, ShenFunc(symcons), gen17598, gen17599)

																						gen17601 := Call(__e, ShenFunc(symcons), gen17593, gen17600)

																						gen17603 = Call(__e, ShenFunc(symbind), V2790, gen17601, V2791, V2792)

																					} else {
																						gen17587 := Call(__e, ShenFunc(symshen_4pvar_2), V2570)

																						if True == gen17587 {
																							Call(__e, ShenFunc(symshen_4bindv), V2570, Nil, V2791)
																							gen17572 := Call(__e, ShenFunc(symtl), V2549)

																							Hyp := gen17572
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen17573 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																							gen17574 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen17575 := Call(__e, ShenFunc(symcons), gen17574, Nil)

																							gen17576 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17575)

																							gen17577 := Call(__e, ShenFunc(symcons), gen17573, gen17576)

																							gen17578 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																							gen17579 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																							gen17580 := Call(__e, ShenFunc(symcons), gen17579, Nil)

																							gen17581 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17580)

																							gen17582 := Call(__e, ShenFunc(symcons), gen17578, gen17581)

																							gen17583 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																							gen17584 := Call(__e, ShenFunc(symcons), gen17582, gen17583)

																							gen17585 := Call(__e, ShenFunc(symcons), gen17577, gen17584)

																							gen17586 := Call(__e, ShenFunc(symbind), V2790, gen17585, V2791, V2792)

																							Result := gen17586
																							Call(__e, ShenFunc(symshen_4unbindv), V2570, V2791)
																							gen17603 = Result

																						} else {
																							gen17603 = False
																						}

																					}
																					Result := gen17603
																					Call(__e, ShenFunc(symshen_4unbindv), V2568, V2791)
																					gen17640 = Result

																				} else {
																					gen17640 = False
																				}

																			}

																		} else {
																			gen17566 := Call(__e, ShenFunc(symshen_4pvar_2), V2567)

																			if True == gen17566 {
																				gen17530 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																				B := gen17530
																				gen17531 := Call(__e, ShenFunc(symcons), B, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2567, gen17531, V2791)

																				gen17532 := Call(__e, ShenFunc(symtl), V2558)

																				gen17533 := Call(__e, ShenFunc(symshen_4lazyderef), gen17532, V2791)

																				V2571 := gen17533
																				gen17564 := Call(__e, ShenFunc(sym_a), Nil, V2571)

																				var gen17565 Obj
																				if True == gen17564 {
																					gen17550 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen17550
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17551 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17552 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17553 := Call(__e, ShenFunc(symcons), gen17552, Nil)

																					gen17554 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17553)

																					gen17555 := Call(__e, ShenFunc(symcons), gen17551, gen17554)

																					gen17556 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17557 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen17558 := Call(__e, ShenFunc(symcons), gen17557, Nil)

																					gen17559 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17558)

																					gen17560 := Call(__e, ShenFunc(symcons), gen17556, gen17559)

																					gen17561 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17562 := Call(__e, ShenFunc(symcons), gen17560, gen17561)

																					gen17563 := Call(__e, ShenFunc(symcons), gen17555, gen17562)

																					gen17565 = Call(__e, ShenFunc(symbind), V2790, gen17563, V2791, V2792)

																				} else {
																					gen17549 := Call(__e, ShenFunc(symshen_4pvar_2), V2571)

																					if True == gen17549 {
																						Call(__e, ShenFunc(symshen_4bindv), V2571, Nil, V2791)
																						gen17534 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen17534
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17535 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17536 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17537 := Call(__e, ShenFunc(symcons), gen17536, Nil)

																						gen17538 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17537)

																						gen17539 := Call(__e, ShenFunc(symcons), gen17535, gen17538)

																						gen17540 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17541 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen17542 := Call(__e, ShenFunc(symcons), gen17541, Nil)

																						gen17543 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17542)

																						gen17544 := Call(__e, ShenFunc(symcons), gen17540, gen17543)

																						gen17545 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17546 := Call(__e, ShenFunc(symcons), gen17544, gen17545)

																						gen17547 := Call(__e, ShenFunc(symcons), gen17539, gen17546)

																						gen17548 := Call(__e, ShenFunc(symbind), V2790, gen17547, V2791, V2792)

																						Result := gen17548
																						Call(__e, ShenFunc(symshen_4unbindv), V2571, V2791)
																						gen17565 = Result

																					} else {
																						gen17565 = False
																					}

																				}
																				Result := gen17565
																				Call(__e, ShenFunc(symshen_4unbindv), V2567, V2791)
																				gen17640 = Result

																			} else {
																				gen17640 = False
																			}

																		}
																		Result := gen17640
																		Call(__e, ShenFunc(symshen_4unbindv), V2561, V2791)
																		gen17767 = Result

																	} else {
																		gen17767 = False
																	}

																}

															} else {
																gen17525 := Call(__e, ShenFunc(symshen_4pvar_2), V2560)

																if True == gen17525 {
																	gen17488 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	B := gen17488
																	gen17489 := Call(__e, ShenFunc(symcons), B, Nil)

																	gen17490 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen17489)

																	Call(__e, ShenFunc(symshen_4bindv), V2560, gen17490, V2791)

																	gen17491 := Call(__e, ShenFunc(symtl), V2558)

																	gen17492 := Call(__e, ShenFunc(symshen_4lazyderef), gen17491, V2791)

																	V2572 := gen17492
																	gen17523 := Call(__e, ShenFunc(sym_a), Nil, V2572)

																	var gen17524 Obj
																	if True == gen17523 {
																		gen17509 := Call(__e, ShenFunc(symtl), V2549)

																		Hyp := gen17509
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen17510 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen17511 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17512 := Call(__e, ShenFunc(symcons), gen17511, Nil)

																		gen17513 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17512)

																		gen17514 := Call(__e, ShenFunc(symcons), gen17510, gen17513)

																		gen17515 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen17516 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																		gen17517 := Call(__e, ShenFunc(symcons), gen17516, Nil)

																		gen17518 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17517)

																		gen17519 := Call(__e, ShenFunc(symcons), gen17515, gen17518)

																		gen17520 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen17521 := Call(__e, ShenFunc(symcons), gen17519, gen17520)

																		gen17522 := Call(__e, ShenFunc(symcons), gen17514, gen17521)

																		gen17524 = Call(__e, ShenFunc(symbind), V2790, gen17522, V2791, V2792)

																	} else {
																		gen17508 := Call(__e, ShenFunc(symshen_4pvar_2), V2572)

																		if True == gen17508 {
																			Call(__e, ShenFunc(symshen_4bindv), V2572, Nil, V2791)
																			gen17493 := Call(__e, ShenFunc(symtl), V2549)

																			Hyp := gen17493
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17494 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17495 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17496 := Call(__e, ShenFunc(symcons), gen17495, Nil)

																			gen17497 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17496)

																			gen17498 := Call(__e, ShenFunc(symcons), gen17494, gen17497)

																			gen17499 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17500 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																			gen17501 := Call(__e, ShenFunc(symcons), gen17500, Nil)

																			gen17502 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17501)

																			gen17503 := Call(__e, ShenFunc(symcons), gen17499, gen17502)

																			gen17504 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17505 := Call(__e, ShenFunc(symcons), gen17503, gen17504)

																			gen17506 := Call(__e, ShenFunc(symcons), gen17498, gen17505)

																			gen17507 := Call(__e, ShenFunc(symbind), V2790, gen17506, V2791, V2792)

																			Result := gen17507
																			Call(__e, ShenFunc(symshen_4unbindv), V2572, V2791)
																			gen17524 = Result

																		} else {
																			gen17524 = False
																		}

																	}
																	Result := gen17524
																	Call(__e, ShenFunc(symshen_4unbindv), V2560, V2791)
																	gen17767 = Result

																} else {
																	gen17767 = False
																}

															}

														} else {
															gen17484 := Call(__e, ShenFunc(symshen_4pvar_2), V2559)

															if True == gen17484 {
																gen17445 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																A := gen17445
																gen17446 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																B := gen17446
																gen17447 := Call(__e, ShenFunc(symcons), B, Nil)

																gen17448 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen17447)

																gen17449 := Call(__e, ShenFunc(symcons), A, gen17448)

																Call(__e, ShenFunc(symshen_4bindv), V2559, gen17449, V2791)

																gen17450 := Call(__e, ShenFunc(symtl), V2558)

																gen17451 := Call(__e, ShenFunc(symshen_4lazyderef), gen17450, V2791)

																V2573 := gen17451
																gen17482 := Call(__e, ShenFunc(sym_a), Nil, V2573)

																var gen17483 Obj
																if True == gen17482 {
																	gen17468 := Call(__e, ShenFunc(symtl), V2549)

																	Hyp := gen17468
																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen17469 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																	gen17470 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen17471 := Call(__e, ShenFunc(symcons), gen17470, Nil)

																	gen17472 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17471)

																	gen17473 := Call(__e, ShenFunc(symcons), gen17469, gen17472)

																	gen17474 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																	gen17475 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																	gen17476 := Call(__e, ShenFunc(symcons), gen17475, Nil)

																	gen17477 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17476)

																	gen17478 := Call(__e, ShenFunc(symcons), gen17474, gen17477)

																	gen17479 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																	gen17480 := Call(__e, ShenFunc(symcons), gen17478, gen17479)

																	gen17481 := Call(__e, ShenFunc(symcons), gen17473, gen17480)

																	gen17483 = Call(__e, ShenFunc(symbind), V2790, gen17481, V2791, V2792)

																} else {
																	gen17467 := Call(__e, ShenFunc(symshen_4pvar_2), V2573)

																	if True == gen17467 {
																		Call(__e, ShenFunc(symshen_4bindv), V2573, Nil, V2791)
																		gen17452 := Call(__e, ShenFunc(symtl), V2549)

																		Hyp := gen17452
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen17453 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen17454 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17455 := Call(__e, ShenFunc(symcons), gen17454, Nil)

																		gen17456 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17455)

																		gen17457 := Call(__e, ShenFunc(symcons), gen17453, gen17456)

																		gen17458 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen17459 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																		gen17460 := Call(__e, ShenFunc(symcons), gen17459, Nil)

																		gen17461 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17460)

																		gen17462 := Call(__e, ShenFunc(symcons), gen17458, gen17461)

																		gen17463 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen17464 := Call(__e, ShenFunc(symcons), gen17462, gen17463)

																		gen17465 := Call(__e, ShenFunc(symcons), gen17457, gen17464)

																		gen17466 := Call(__e, ShenFunc(symbind), V2790, gen17465, V2791, V2792)

																		Result := gen17466
																		Call(__e, ShenFunc(symshen_4unbindv), V2573, V2791)
																		gen17483 = Result

																	} else {
																		gen17483 = False
																	}

																}
																Result := gen17483
																Call(__e, ShenFunc(symshen_4unbindv), V2559, V2791)
																gen17767 = Result

															} else {
																gen17767 = False
															}

														}

													} else {
														gen17767 = False
													}

												} else {
													gen17767 = False
												}

											} else {
												gen17767 = False
											}

										} else {
											gen17767 = False
										}

									} else {
										gen17767 = False
									}

								} else {
									gen17767 = False
								}

							} else {
								gen17767 = False
							}

						} else {
							gen17767 = False
						}

					} else {
						gen17767 = False
					}

				} else {
					gen17767 = False
				}
				Case := gen17767
				gen18204 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen18204 {
					gen17768 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

					V2574 := gen17768
					gen18096 := Call(__e, ShenFunc(symcons_2), V2574)

					var gen18097 Obj
					if True == gen18096 {
						gen17769 := Call(__e, ShenFunc(symhd), V2574)

						gen17770 := Call(__e, ShenFunc(symshen_4lazyderef), gen17769, V2791)

						V2575 := gen17770
						gen18095 := Call(__e, ShenFunc(symcons_2), V2575)

						if True == gen18095 {
							gen17771 := Call(__e, ShenFunc(symhd), V2575)

							gen17772 := Call(__e, ShenFunc(symshen_4lazyderef), gen17771, V2791)

							V2576 := gen17772
							gen18094 := Call(__e, ShenFunc(symcons_2), V2576)

							if True == gen18094 {
								gen17773 := Call(__e, ShenFunc(symhd), V2576)

								gen17774 := Call(__e, ShenFunc(symshen_4lazyderef), gen17773, V2791)

								V2577 := gen17774
								gen18093 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), V2577)

								if True == gen18093 {
									gen17775 := Call(__e, ShenFunc(symtl), V2576)

									gen17776 := Call(__e, ShenFunc(symshen_4lazyderef), gen17775, V2791)

									V2578 := gen17776
									gen18092 := Call(__e, ShenFunc(symcons_2), V2578)

									if True == gen18092 {
										gen17777 := Call(__e, ShenFunc(symhd), V2578)

										X := gen17777
										gen17778 := Call(__e, ShenFunc(symtl), V2578)

										gen17779 := Call(__e, ShenFunc(symshen_4lazyderef), gen17778, V2791)

										V2579 := gen17779
										gen18091 := Call(__e, ShenFunc(symcons_2), V2579)

										if True == gen18091 {
											gen17780 := Call(__e, ShenFunc(symhd), V2579)

											Y := gen17780
											gen17781 := Call(__e, ShenFunc(symtl), V2579)

											gen17782 := Call(__e, ShenFunc(symshen_4lazyderef), gen17781, V2791)

											V2580 := gen17782
											gen18090 := Call(__e, ShenFunc(sym_a), Nil, V2580)

											if True == gen18090 {
												gen17783 := Call(__e, ShenFunc(symtl), V2575)

												gen17784 := Call(__e, ShenFunc(symshen_4lazyderef), gen17783, V2791)

												V2581 := gen17784
												gen18089 := Call(__e, ShenFunc(symcons_2), V2581)

												if True == gen18089 {
													gen17785 := Call(__e, ShenFunc(symhd), V2581)

													gen17786 := Call(__e, ShenFunc(symshen_4lazyderef), gen17785, V2791)

													V2582 := gen17786
													gen18088 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2582)

													if True == gen18088 {
														gen17787 := Call(__e, ShenFunc(symtl), V2581)

														gen17788 := Call(__e, ShenFunc(symshen_4lazyderef), gen17787, V2791)

														V2583 := gen17788
														gen18087 := Call(__e, ShenFunc(symcons_2), V2583)

														if True == gen18087 {
															gen17789 := Call(__e, ShenFunc(symhd), V2583)

															gen17790 := Call(__e, ShenFunc(symshen_4lazyderef), gen17789, V2791)

															V2584 := gen17790
															gen18086 := Call(__e, ShenFunc(symcons_2), V2584)

															if True == gen18086 {
																gen17833 := Call(__e, ShenFunc(symhd), V2584)

																gen17834 := Call(__e, ShenFunc(symshen_4lazyderef), gen17833, V2791)

																V2585 := gen17834
																gen18085 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), V2585)

																if True == gen18085 {
																	gen17961 := Call(__e, ShenFunc(symtl), V2584)

																	gen17962 := Call(__e, ShenFunc(symshen_4lazyderef), gen17961, V2791)

																	V2586 := gen17962
																	gen18084 := Call(__e, ShenFunc(symcons_2), V2586)

																	if True == gen18084 {
																		gen18004 := Call(__e, ShenFunc(symhd), V2586)

																		A := gen18004
																		gen18005 := Call(__e, ShenFunc(symtl), V2586)

																		gen18006 := Call(__e, ShenFunc(symshen_4lazyderef), gen18005, V2791)

																		V2587 := gen18006
																		gen18083 := Call(__e, ShenFunc(sym_a), Nil, V2587)

																		if True == gen18083 {
																			gen18046 := Call(__e, ShenFunc(symtl), V2583)

																			gen18047 := Call(__e, ShenFunc(symshen_4lazyderef), gen18046, V2791)

																			V2588 := gen18047
																			gen18082 := Call(__e, ShenFunc(sym_a), Nil, V2588)

																			if True == gen18082 {
																				gen18066 := Call(__e, ShenFunc(symtl), V2574)

																				Hyp := gen18066
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen18067 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen18068 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen18069 := Call(__e, ShenFunc(symcons), gen18068, Nil)

																				gen18070 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18069)

																				gen18071 := Call(__e, ShenFunc(symcons), gen18067, gen18070)

																				gen18072 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen18073 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen18074 := Call(__e, ShenFunc(symcons), gen18073, Nil)

																				gen18075 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen18074)

																				gen18076 := Call(__e, ShenFunc(symcons), gen18075, Nil)

																				gen18077 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18076)

																				gen18078 := Call(__e, ShenFunc(symcons), gen18072, gen18077)

																				gen18079 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen18080 := Call(__e, ShenFunc(symcons), gen18078, gen18079)

																				gen18081 := Call(__e, ShenFunc(symcons), gen18071, gen18080)

																				gen18097 = Call(__e, ShenFunc(symbind), V2790, gen18081, V2791, V2792)

																			} else {
																				gen18065 := Call(__e, ShenFunc(symshen_4pvar_2), V2588)

																				if True == gen18065 {
																					Call(__e, ShenFunc(symshen_4bindv), V2588, Nil, V2791)
																					gen18048 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen18048
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen18049 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen18050 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen18051 := Call(__e, ShenFunc(symcons), gen18050, Nil)

																					gen18052 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18051)

																					gen18053 := Call(__e, ShenFunc(symcons), gen18049, gen18052)

																					gen18054 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen18055 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen18056 := Call(__e, ShenFunc(symcons), gen18055, Nil)

																					gen18057 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen18056)

																					gen18058 := Call(__e, ShenFunc(symcons), gen18057, Nil)

																					gen18059 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18058)

																					gen18060 := Call(__e, ShenFunc(symcons), gen18054, gen18059)

																					gen18061 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen18062 := Call(__e, ShenFunc(symcons), gen18060, gen18061)

																					gen18063 := Call(__e, ShenFunc(symcons), gen18053, gen18062)

																					gen18064 := Call(__e, ShenFunc(symbind), V2790, gen18063, V2791, V2792)

																					Result := gen18064
																					Call(__e, ShenFunc(symshen_4unbindv), V2588, V2791)
																					gen18097 = Result

																				} else {
																					gen18097 = False
																				}

																			}

																		} else {
																			gen18045 := Call(__e, ShenFunc(symshen_4pvar_2), V2587)

																			if True == gen18045 {
																				Call(__e, ShenFunc(symshen_4bindv), V2587, Nil, V2791)
																				gen18007 := Call(__e, ShenFunc(symtl), V2583)

																				gen18008 := Call(__e, ShenFunc(symshen_4lazyderef), gen18007, V2791)

																				V2589 := gen18008
																				gen18043 := Call(__e, ShenFunc(sym_a), Nil, V2589)

																				var gen18044 Obj
																				if True == gen18043 {
																					gen18027 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen18027
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen18028 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen18029 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen18030 := Call(__e, ShenFunc(symcons), gen18029, Nil)

																					gen18031 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18030)

																					gen18032 := Call(__e, ShenFunc(symcons), gen18028, gen18031)

																					gen18033 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen18034 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen18035 := Call(__e, ShenFunc(symcons), gen18034, Nil)

																					gen18036 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen18035)

																					gen18037 := Call(__e, ShenFunc(symcons), gen18036, Nil)

																					gen18038 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18037)

																					gen18039 := Call(__e, ShenFunc(symcons), gen18033, gen18038)

																					gen18040 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen18041 := Call(__e, ShenFunc(symcons), gen18039, gen18040)

																					gen18042 := Call(__e, ShenFunc(symcons), gen18032, gen18041)

																					gen18044 = Call(__e, ShenFunc(symbind), V2790, gen18042, V2791, V2792)

																				} else {
																					gen18026 := Call(__e, ShenFunc(symshen_4pvar_2), V2589)

																					if True == gen18026 {
																						Call(__e, ShenFunc(symshen_4bindv), V2589, Nil, V2791)
																						gen18009 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen18009
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen18010 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen18011 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen18012 := Call(__e, ShenFunc(symcons), gen18011, Nil)

																						gen18013 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18012)

																						gen18014 := Call(__e, ShenFunc(symcons), gen18010, gen18013)

																						gen18015 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen18016 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen18017 := Call(__e, ShenFunc(symcons), gen18016, Nil)

																						gen18018 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen18017)

																						gen18019 := Call(__e, ShenFunc(symcons), gen18018, Nil)

																						gen18020 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18019)

																						gen18021 := Call(__e, ShenFunc(symcons), gen18015, gen18020)

																						gen18022 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen18023 := Call(__e, ShenFunc(symcons), gen18021, gen18022)

																						gen18024 := Call(__e, ShenFunc(symcons), gen18014, gen18023)

																						gen18025 := Call(__e, ShenFunc(symbind), V2790, gen18024, V2791, V2792)

																						Result := gen18025
																						Call(__e, ShenFunc(symshen_4unbindv), V2589, V2791)
																						gen18044 = Result

																					} else {
																						gen18044 = False
																					}

																				}
																				Result := gen18044
																				Call(__e, ShenFunc(symshen_4unbindv), V2587, V2791)
																				gen18097 = Result

																			} else {
																				gen18097 = False
																			}

																		}

																	} else {
																		gen18003 := Call(__e, ShenFunc(symshen_4pvar_2), V2586)

																		if True == gen18003 {
																			gen17963 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																			A := gen17963
																			gen17964 := Call(__e, ShenFunc(symcons), A, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2586, gen17964, V2791)

																			gen17965 := Call(__e, ShenFunc(symtl), V2583)

																			gen17966 := Call(__e, ShenFunc(symshen_4lazyderef), gen17965, V2791)

																			V2590 := gen17966
																			gen18001 := Call(__e, ShenFunc(sym_a), Nil, V2590)

																			var gen18002 Obj
																			if True == gen18001 {
																				gen17985 := Call(__e, ShenFunc(symtl), V2574)

																				Hyp := gen17985
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen17986 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen17987 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17988 := Call(__e, ShenFunc(symcons), gen17987, Nil)

																				gen17989 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17988)

																				gen17990 := Call(__e, ShenFunc(symcons), gen17986, gen17989)

																				gen17991 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen17992 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen17993 := Call(__e, ShenFunc(symcons), gen17992, Nil)

																				gen17994 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17993)

																				gen17995 := Call(__e, ShenFunc(symcons), gen17994, Nil)

																				gen17996 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17995)

																				gen17997 := Call(__e, ShenFunc(symcons), gen17991, gen17996)

																				gen17998 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen17999 := Call(__e, ShenFunc(symcons), gen17997, gen17998)

																				gen18000 := Call(__e, ShenFunc(symcons), gen17990, gen17999)

																				gen18002 = Call(__e, ShenFunc(symbind), V2790, gen18000, V2791, V2792)

																			} else {
																				gen17984 := Call(__e, ShenFunc(symshen_4pvar_2), V2590)

																				if True == gen17984 {
																					Call(__e, ShenFunc(symshen_4bindv), V2590, Nil, V2791)
																					gen17967 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen17967
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17968 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17969 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17970 := Call(__e, ShenFunc(symcons), gen17969, Nil)

																					gen17971 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17970)

																					gen17972 := Call(__e, ShenFunc(symcons), gen17968, gen17971)

																					gen17973 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17974 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17975 := Call(__e, ShenFunc(symcons), gen17974, Nil)

																					gen17976 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17975)

																					gen17977 := Call(__e, ShenFunc(symcons), gen17976, Nil)

																					gen17978 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17977)

																					gen17979 := Call(__e, ShenFunc(symcons), gen17973, gen17978)

																					gen17980 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17981 := Call(__e, ShenFunc(symcons), gen17979, gen17980)

																					gen17982 := Call(__e, ShenFunc(symcons), gen17972, gen17981)

																					gen17983 := Call(__e, ShenFunc(symbind), V2790, gen17982, V2791, V2792)

																					Result := gen17983
																					Call(__e, ShenFunc(symshen_4unbindv), V2590, V2791)
																					gen18002 = Result

																				} else {
																					gen18002 = False
																				}

																			}
																			Result := gen18002
																			Call(__e, ShenFunc(symshen_4unbindv), V2586, V2791)
																			gen18097 = Result

																		} else {
																			gen18097 = False
																		}

																	}

																} else {
																	gen17960 := Call(__e, ShenFunc(symshen_4pvar_2), V2585)

																	if True == gen17960 {
																		Call(__e, ShenFunc(symshen_4bindv), V2585, MakeSymbol("vector"), V2791)
																		gen17835 := Call(__e, ShenFunc(symtl), V2584)

																		gen17836 := Call(__e, ShenFunc(symshen_4lazyderef), gen17835, V2791)

																		V2591 := gen17836
																		gen17958 := Call(__e, ShenFunc(symcons_2), V2591)

																		var gen17959 Obj
																		if True == gen17958 {
																			gen17878 := Call(__e, ShenFunc(symhd), V2591)

																			A := gen17878
																			gen17879 := Call(__e, ShenFunc(symtl), V2591)

																			gen17880 := Call(__e, ShenFunc(symshen_4lazyderef), gen17879, V2791)

																			V2592 := gen17880
																			gen17957 := Call(__e, ShenFunc(sym_a), Nil, V2592)

																			if True == gen17957 {
																				gen17920 := Call(__e, ShenFunc(symtl), V2583)

																				gen17921 := Call(__e, ShenFunc(symshen_4lazyderef), gen17920, V2791)

																				V2593 := gen17921
																				gen17956 := Call(__e, ShenFunc(sym_a), Nil, V2593)

																				if True == gen17956 {
																					gen17940 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen17940
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17941 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17942 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17943 := Call(__e, ShenFunc(symcons), gen17942, Nil)

																					gen17944 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17943)

																					gen17945 := Call(__e, ShenFunc(symcons), gen17941, gen17944)

																					gen17946 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17947 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17948 := Call(__e, ShenFunc(symcons), gen17947, Nil)

																					gen17949 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17948)

																					gen17950 := Call(__e, ShenFunc(symcons), gen17949, Nil)

																					gen17951 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17950)

																					gen17952 := Call(__e, ShenFunc(symcons), gen17946, gen17951)

																					gen17953 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17954 := Call(__e, ShenFunc(symcons), gen17952, gen17953)

																					gen17955 := Call(__e, ShenFunc(symcons), gen17945, gen17954)

																					gen17959 = Call(__e, ShenFunc(symbind), V2790, gen17955, V2791, V2792)

																				} else {
																					gen17939 := Call(__e, ShenFunc(symshen_4pvar_2), V2593)

																					if True == gen17939 {
																						Call(__e, ShenFunc(symshen_4bindv), V2593, Nil, V2791)
																						gen17922 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen17922
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17923 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17924 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17925 := Call(__e, ShenFunc(symcons), gen17924, Nil)

																						gen17926 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17925)

																						gen17927 := Call(__e, ShenFunc(symcons), gen17923, gen17926)

																						gen17928 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17929 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17930 := Call(__e, ShenFunc(symcons), gen17929, Nil)

																						gen17931 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17930)

																						gen17932 := Call(__e, ShenFunc(symcons), gen17931, Nil)

																						gen17933 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17932)

																						gen17934 := Call(__e, ShenFunc(symcons), gen17928, gen17933)

																						gen17935 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17936 := Call(__e, ShenFunc(symcons), gen17934, gen17935)

																						gen17937 := Call(__e, ShenFunc(symcons), gen17927, gen17936)

																						gen17938 := Call(__e, ShenFunc(symbind), V2790, gen17937, V2791, V2792)

																						Result := gen17938
																						Call(__e, ShenFunc(symshen_4unbindv), V2593, V2791)
																						gen17959 = Result

																					} else {
																						gen17959 = False
																					}

																				}

																			} else {
																				gen17919 := Call(__e, ShenFunc(symshen_4pvar_2), V2592)

																				if True == gen17919 {
																					Call(__e, ShenFunc(symshen_4bindv), V2592, Nil, V2791)
																					gen17881 := Call(__e, ShenFunc(symtl), V2583)

																					gen17882 := Call(__e, ShenFunc(symshen_4lazyderef), gen17881, V2791)

																					V2594 := gen17882
																					gen17917 := Call(__e, ShenFunc(sym_a), Nil, V2594)

																					var gen17918 Obj
																					if True == gen17917 {
																						gen17901 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen17901
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17902 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17903 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17904 := Call(__e, ShenFunc(symcons), gen17903, Nil)

																						gen17905 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17904)

																						gen17906 := Call(__e, ShenFunc(symcons), gen17902, gen17905)

																						gen17907 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17908 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17909 := Call(__e, ShenFunc(symcons), gen17908, Nil)

																						gen17910 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17909)

																						gen17911 := Call(__e, ShenFunc(symcons), gen17910, Nil)

																						gen17912 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17911)

																						gen17913 := Call(__e, ShenFunc(symcons), gen17907, gen17912)

																						gen17914 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17915 := Call(__e, ShenFunc(symcons), gen17913, gen17914)

																						gen17916 := Call(__e, ShenFunc(symcons), gen17906, gen17915)

																						gen17918 = Call(__e, ShenFunc(symbind), V2790, gen17916, V2791, V2792)

																					} else {
																						gen17900 := Call(__e, ShenFunc(symshen_4pvar_2), V2594)

																						if True == gen17900 {
																							Call(__e, ShenFunc(symshen_4bindv), V2594, Nil, V2791)
																							gen17883 := Call(__e, ShenFunc(symtl), V2574)

																							Hyp := gen17883
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen17884 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																							gen17885 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen17886 := Call(__e, ShenFunc(symcons), gen17885, Nil)

																							gen17887 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17886)

																							gen17888 := Call(__e, ShenFunc(symcons), gen17884, gen17887)

																							gen17889 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																							gen17890 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen17891 := Call(__e, ShenFunc(symcons), gen17890, Nil)

																							gen17892 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17891)

																							gen17893 := Call(__e, ShenFunc(symcons), gen17892, Nil)

																							gen17894 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17893)

																							gen17895 := Call(__e, ShenFunc(symcons), gen17889, gen17894)

																							gen17896 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																							gen17897 := Call(__e, ShenFunc(symcons), gen17895, gen17896)

																							gen17898 := Call(__e, ShenFunc(symcons), gen17888, gen17897)

																							gen17899 := Call(__e, ShenFunc(symbind), V2790, gen17898, V2791, V2792)

																							Result := gen17899
																							Call(__e, ShenFunc(symshen_4unbindv), V2594, V2791)
																							gen17918 = Result

																						} else {
																							gen17918 = False
																						}

																					}
																					Result := gen17918
																					Call(__e, ShenFunc(symshen_4unbindv), V2592, V2791)
																					gen17959 = Result

																				} else {
																					gen17959 = False
																				}

																			}

																		} else {
																			gen17877 := Call(__e, ShenFunc(symshen_4pvar_2), V2591)

																			if True == gen17877 {
																				gen17837 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																				A := gen17837
																				gen17838 := Call(__e, ShenFunc(symcons), A, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2591, gen17838, V2791)

																				gen17839 := Call(__e, ShenFunc(symtl), V2583)

																				gen17840 := Call(__e, ShenFunc(symshen_4lazyderef), gen17839, V2791)

																				V2595 := gen17840
																				gen17875 := Call(__e, ShenFunc(sym_a), Nil, V2595)

																				var gen17876 Obj
																				if True == gen17875 {
																					gen17859 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen17859
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen17860 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen17861 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17862 := Call(__e, ShenFunc(symcons), gen17861, Nil)

																					gen17863 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17862)

																					gen17864 := Call(__e, ShenFunc(symcons), gen17860, gen17863)

																					gen17865 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen17866 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen17867 := Call(__e, ShenFunc(symcons), gen17866, Nil)

																					gen17868 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17867)

																					gen17869 := Call(__e, ShenFunc(symcons), gen17868, Nil)

																					gen17870 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17869)

																					gen17871 := Call(__e, ShenFunc(symcons), gen17865, gen17870)

																					gen17872 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen17873 := Call(__e, ShenFunc(symcons), gen17871, gen17872)

																					gen17874 := Call(__e, ShenFunc(symcons), gen17864, gen17873)

																					gen17876 = Call(__e, ShenFunc(symbind), V2790, gen17874, V2791, V2792)

																				} else {
																					gen17858 := Call(__e, ShenFunc(symshen_4pvar_2), V2595)

																					if True == gen17858 {
																						Call(__e, ShenFunc(symshen_4bindv), V2595, Nil, V2791)
																						gen17841 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen17841
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen17842 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen17843 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17844 := Call(__e, ShenFunc(symcons), gen17843, Nil)

																						gen17845 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17844)

																						gen17846 := Call(__e, ShenFunc(symcons), gen17842, gen17845)

																						gen17847 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen17848 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen17849 := Call(__e, ShenFunc(symcons), gen17848, Nil)

																						gen17850 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17849)

																						gen17851 := Call(__e, ShenFunc(symcons), gen17850, Nil)

																						gen17852 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17851)

																						gen17853 := Call(__e, ShenFunc(symcons), gen17847, gen17852)

																						gen17854 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen17855 := Call(__e, ShenFunc(symcons), gen17853, gen17854)

																						gen17856 := Call(__e, ShenFunc(symcons), gen17846, gen17855)

																						gen17857 := Call(__e, ShenFunc(symbind), V2790, gen17856, V2791, V2792)

																						Result := gen17857
																						Call(__e, ShenFunc(symshen_4unbindv), V2595, V2791)
																						gen17876 = Result

																					} else {
																						gen17876 = False
																					}

																				}
																				Result := gen17876
																				Call(__e, ShenFunc(symshen_4unbindv), V2591, V2791)
																				gen17959 = Result

																			} else {
																				gen17959 = False
																			}

																		}
																		Result := gen17959
																		Call(__e, ShenFunc(symshen_4unbindv), V2585, V2791)
																		gen18097 = Result

																	} else {
																		gen18097 = False
																	}

																}

															} else {
																gen17832 := Call(__e, ShenFunc(symshen_4pvar_2), V2584)

																if True == gen17832 {
																	gen17791 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	A := gen17791
																	gen17792 := Call(__e, ShenFunc(symcons), A, Nil)

																	gen17793 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17792)

																	Call(__e, ShenFunc(symshen_4bindv), V2584, gen17793, V2791)

																	gen17794 := Call(__e, ShenFunc(symtl), V2583)

																	gen17795 := Call(__e, ShenFunc(symshen_4lazyderef), gen17794, V2791)

																	V2596 := gen17795
																	gen17830 := Call(__e, ShenFunc(sym_a), Nil, V2596)

																	var gen17831 Obj
																	if True == gen17830 {
																		gen17814 := Call(__e, ShenFunc(symtl), V2574)

																		Hyp := gen17814
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen17815 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen17816 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17817 := Call(__e, ShenFunc(symcons), gen17816, Nil)

																		gen17818 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17817)

																		gen17819 := Call(__e, ShenFunc(symcons), gen17815, gen17818)

																		gen17820 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen17821 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen17822 := Call(__e, ShenFunc(symcons), gen17821, Nil)

																		gen17823 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17822)

																		gen17824 := Call(__e, ShenFunc(symcons), gen17823, Nil)

																		gen17825 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17824)

																		gen17826 := Call(__e, ShenFunc(symcons), gen17820, gen17825)

																		gen17827 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen17828 := Call(__e, ShenFunc(symcons), gen17826, gen17827)

																		gen17829 := Call(__e, ShenFunc(symcons), gen17819, gen17828)

																		gen17831 = Call(__e, ShenFunc(symbind), V2790, gen17829, V2791, V2792)

																	} else {
																		gen17813 := Call(__e, ShenFunc(symshen_4pvar_2), V2596)

																		if True == gen17813 {
																			Call(__e, ShenFunc(symshen_4bindv), V2596, Nil, V2791)
																			gen17796 := Call(__e, ShenFunc(symtl), V2574)

																			Hyp := gen17796
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen17797 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen17798 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17799 := Call(__e, ShenFunc(symcons), gen17798, Nil)

																			gen17800 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17799)

																			gen17801 := Call(__e, ShenFunc(symcons), gen17797, gen17800)

																			gen17802 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen17803 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen17804 := Call(__e, ShenFunc(symcons), gen17803, Nil)

																			gen17805 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen17804)

																			gen17806 := Call(__e, ShenFunc(symcons), gen17805, Nil)

																			gen17807 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen17806)

																			gen17808 := Call(__e, ShenFunc(symcons), gen17802, gen17807)

																			gen17809 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen17810 := Call(__e, ShenFunc(symcons), gen17808, gen17809)

																			gen17811 := Call(__e, ShenFunc(symcons), gen17801, gen17810)

																			gen17812 := Call(__e, ShenFunc(symbind), V2790, gen17811, V2791, V2792)

																			Result := gen17812
																			Call(__e, ShenFunc(symshen_4unbindv), V2596, V2791)
																			gen17831 = Result

																		} else {
																			gen17831 = False
																		}

																	}
																	Result := gen17831
																	Call(__e, ShenFunc(symshen_4unbindv), V2584, V2791)
																	gen18097 = Result

																} else {
																	gen18097 = False
																}

															}

														} else {
															gen18097 = False
														}

													} else {
														gen18097 = False
													}

												} else {
													gen18097 = False
												}

											} else {
												gen18097 = False
											}

										} else {
											gen18097 = False
										}

									} else {
										gen18097 = False
									}

								} else {
									gen18097 = False
								}

							} else {
								gen18097 = False
							}

						} else {
							gen18097 = False
						}

					} else {
						gen18097 = False
					}
					Case := gen18097
					gen18203 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen18203 {
						gen18098 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

						V2597 := gen18098
						gen18191 := Call(__e, ShenFunc(symcons_2), V2597)

						var gen18192 Obj
						if True == gen18191 {
							gen18099 := Call(__e, ShenFunc(symhd), V2597)

							gen18100 := Call(__e, ShenFunc(symshen_4lazyderef), gen18099, V2791)

							V2598 := gen18100
							gen18190 := Call(__e, ShenFunc(symcons_2), V2598)

							if True == gen18190 {
								gen18101 := Call(__e, ShenFunc(symhd), V2598)

								gen18102 := Call(__e, ShenFunc(symshen_4lazyderef), gen18101, V2791)

								V2599 := gen18102
								gen18189 := Call(__e, ShenFunc(symcons_2), V2599)

								if True == gen18189 {
									gen18103 := Call(__e, ShenFunc(symhd), V2599)

									gen18104 := Call(__e, ShenFunc(symshen_4lazyderef), gen18103, V2791)

									V2600 := gen18104
									gen18188 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), V2600)

									if True == gen18188 {
										gen18105 := Call(__e, ShenFunc(symtl), V2599)

										gen18106 := Call(__e, ShenFunc(symshen_4lazyderef), gen18105, V2791)

										V2601 := gen18106
										gen18187 := Call(__e, ShenFunc(symcons_2), V2601)

										if True == gen18187 {
											gen18107 := Call(__e, ShenFunc(symhd), V2601)

											X := gen18107
											gen18108 := Call(__e, ShenFunc(symtl), V2601)

											gen18109 := Call(__e, ShenFunc(symshen_4lazyderef), gen18108, V2791)

											V2602 := gen18109
											gen18186 := Call(__e, ShenFunc(symcons_2), V2602)

											if True == gen18186 {
												gen18110 := Call(__e, ShenFunc(symhd), V2602)

												Y := gen18110
												gen18111 := Call(__e, ShenFunc(symtl), V2602)

												gen18112 := Call(__e, ShenFunc(symshen_4lazyderef), gen18111, V2791)

												V2603 := gen18112
												gen18185 := Call(__e, ShenFunc(sym_a), Nil, V2603)

												if True == gen18185 {
													gen18113 := Call(__e, ShenFunc(symtl), V2598)

													gen18114 := Call(__e, ShenFunc(symshen_4lazyderef), gen18113, V2791)

													V2604 := gen18114
													gen18184 := Call(__e, ShenFunc(symcons_2), V2604)

													if True == gen18184 {
														gen18115 := Call(__e, ShenFunc(symhd), V2604)

														gen18116 := Call(__e, ShenFunc(symshen_4lazyderef), gen18115, V2791)

														V2605 := gen18116
														gen18183 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2605)

														if True == gen18183 {
															gen18117 := Call(__e, ShenFunc(symtl), V2604)

															gen18118 := Call(__e, ShenFunc(symshen_4lazyderef), gen18117, V2791)

															V2606 := gen18118
															gen18182 := Call(__e, ShenFunc(symcons_2), V2606)

															if True == gen18182 {
																gen18119 := Call(__e, ShenFunc(symhd), V2606)

																gen18120 := Call(__e, ShenFunc(symshen_4lazyderef), gen18119, V2791)

																V2607 := gen18120
																gen18181 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2607)

																if True == gen18181 {
																	gen18152 := Call(__e, ShenFunc(symtl), V2606)

																	gen18153 := Call(__e, ShenFunc(symshen_4lazyderef), gen18152, V2791)

																	V2608 := gen18153
																	gen18180 := Call(__e, ShenFunc(sym_a), Nil, V2608)

																	if True == gen18180 {
																		gen18168 := Call(__e, ShenFunc(symtl), V2597)

																		Hyp := gen18168
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen18169 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen18170 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																		gen18171 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18170)

																		gen18172 := Call(__e, ShenFunc(symcons), gen18169, gen18171)

																		gen18173 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen18174 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																		gen18175 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18174)

																		gen18176 := Call(__e, ShenFunc(symcons), gen18173, gen18175)

																		gen18177 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen18178 := Call(__e, ShenFunc(symcons), gen18176, gen18177)

																		gen18179 := Call(__e, ShenFunc(symcons), gen18172, gen18178)

																		gen18192 = Call(__e, ShenFunc(symbind), V2790, gen18179, V2791, V2792)

																	} else {
																		gen18167 := Call(__e, ShenFunc(symshen_4pvar_2), V2608)

																		if True == gen18167 {
																			Call(__e, ShenFunc(symshen_4bindv), V2608, Nil, V2791)
																			gen18154 := Call(__e, ShenFunc(symtl), V2597)

																			Hyp := gen18154
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen18155 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen18156 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen18157 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18156)

																			gen18158 := Call(__e, ShenFunc(symcons), gen18155, gen18157)

																			gen18159 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen18160 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen18161 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18160)

																			gen18162 := Call(__e, ShenFunc(symcons), gen18159, gen18161)

																			gen18163 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen18164 := Call(__e, ShenFunc(symcons), gen18162, gen18163)

																			gen18165 := Call(__e, ShenFunc(symcons), gen18158, gen18164)

																			gen18166 := Call(__e, ShenFunc(symbind), V2790, gen18165, V2791, V2792)

																			Result := gen18166
																			Call(__e, ShenFunc(symshen_4unbindv), V2608, V2791)
																			gen18192 = Result

																		} else {
																			gen18192 = False
																		}

																	}

																} else {
																	gen18151 := Call(__e, ShenFunc(symshen_4pvar_2), V2607)

																	if True == gen18151 {
																		Call(__e, ShenFunc(symshen_4bindv), V2607, MakeSymbol("string"), V2791)
																		gen18121 := Call(__e, ShenFunc(symtl), V2606)

																		gen18122 := Call(__e, ShenFunc(symshen_4lazyderef), gen18121, V2791)

																		V2609 := gen18122
																		gen18149 := Call(__e, ShenFunc(sym_a), Nil, V2609)

																		var gen18150 Obj
																		if True == gen18149 {
																			gen18137 := Call(__e, ShenFunc(symtl), V2597)

																			Hyp := gen18137
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen18138 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen18139 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen18140 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18139)

																			gen18141 := Call(__e, ShenFunc(symcons), gen18138, gen18140)

																			gen18142 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen18143 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen18144 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18143)

																			gen18145 := Call(__e, ShenFunc(symcons), gen18142, gen18144)

																			gen18146 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen18147 := Call(__e, ShenFunc(symcons), gen18145, gen18146)

																			gen18148 := Call(__e, ShenFunc(symcons), gen18141, gen18147)

																			gen18150 = Call(__e, ShenFunc(symbind), V2790, gen18148, V2791, V2792)

																		} else {
																			gen18136 := Call(__e, ShenFunc(symshen_4pvar_2), V2609)

																			if True == gen18136 {
																				Call(__e, ShenFunc(symshen_4bindv), V2609, Nil, V2791)
																				gen18123 := Call(__e, ShenFunc(symtl), V2597)

																				Hyp := gen18123
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen18124 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen18125 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																				gen18126 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18125)

																				gen18127 := Call(__e, ShenFunc(symcons), gen18124, gen18126)

																				gen18128 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen18129 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																				gen18130 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18129)

																				gen18131 := Call(__e, ShenFunc(symcons), gen18128, gen18130)

																				gen18132 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen18133 := Call(__e, ShenFunc(symcons), gen18131, gen18132)

																				gen18134 := Call(__e, ShenFunc(symcons), gen18127, gen18133)

																				gen18135 := Call(__e, ShenFunc(symbind), V2790, gen18134, V2791, V2792)

																				Result := gen18135
																				Call(__e, ShenFunc(symshen_4unbindv), V2609, V2791)
																				gen18150 = Result

																			} else {
																				gen18150 = False
																			}

																		}
																		Result := gen18150
																		Call(__e, ShenFunc(symshen_4unbindv), V2607, V2791)
																		gen18192 = Result

																	} else {
																		gen18192 = False
																	}

																}

															} else {
																gen18192 = False
															}

														} else {
															gen18192 = False
														}

													} else {
														gen18192 = False
													}

												} else {
													gen18192 = False
												}

											} else {
												gen18192 = False
											}

										} else {
											gen18192 = False
										}

									} else {
										gen18192 = False
									}

								} else {
									gen18192 = False
								}

							} else {
								gen18192 = False
							}

						} else {
							gen18192 = False
						}
						Case := gen18192
						gen18202 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen18202 {
							gen18193 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

							V2610 := gen18193
							gen18201 := Call(__e, ShenFunc(symcons_2), V2610)

							if True == gen18201 {
								gen18194 := Call(__e, ShenFunc(symhd), V2610)

								X := gen18194
								gen18195 := Call(__e, ShenFunc(symtl), V2610)

								Hyp := gen18195
								gen18196 := Call(__e, ShenFunc(symshen_4newpv), V2791)

								NewHyps := gen18196
								Call(__e, ShenFunc(symshen_4incinfs))
								gen18197 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

								gen18198 := Call(__e, ShenFunc(symshen_4lazyderef), NewHyps, V2791)

								gen18199 := Call(__e, ShenFunc(symcons), gen18197, gen18198)

								gen18200 := MakeNative(func(__e Evaluator, __args ...Obj) {
									__e.TailApply(ShenFunc(symshen_4t_d_1hyps), Hyp, NewHyps, V2791, V2792)

									return
								}, 0)
								__e.TailApply(ShenFunc(symbind), V2790, gen18199, V2791, gen18200)

								return

							} else {
								__e.Return(False)
								return
							}

						} else {
							__e.Return(Case)
							return
						}

					} else {
						__e.Return(Case)
						return
					}

				} else {
					__e.Return(Case)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-hyps"), gen18206)

		gen18211 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2809 := __args[0]
			_ = V2809
			V2810 := __args[1]
			_ = V2810
			V2811 := __args[2]
			_ = V2811
			V2812 := __args[3]
			_ = V2812
			gen18210 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*spy*"))

			if True == gen18210 {
				Call(__e, ShenFunc(symshen_4line))
				gen18207 := Call(__e, ShenFunc(symshen_4deref), V2809, V2811)

				Call(__e, ShenFunc(symshen_4show_1p), gen18207)

				Call(__e, ShenFunc(symnl), MakeNumber(1))
				Call(__e, ShenFunc(symnl), MakeNumber(1))
				gen18208 := Call(__e, ShenFunc(symshen_4deref), V2810, V2811)

				Call(__e, ShenFunc(symshen_4show_1assumptions), gen18208, MakeNumber(1))

				gen18209 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("\n> "), gen18209)

				Call(__e, ShenFunc(symshen_4pause_1for_1user))
				__e.TailApply(ShenFunc(symthaw), V2812)

				return

			} else {
				__e.TailApply(ShenFunc(symthaw), V2812)

				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show"), gen18211)

		gen18220 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen18212 := Call(__e, ShenFunc(syminferences))

			Infs := gen18212
			gen18213 := Call(__e, ShenFunc(sym_a), MakeNumber(1), Infs)

			var gen18214 Obj
			if True == gen18213 {
				gen18214 = MakeString("")
			} else {
				gen18214 = MakeString("s")
			}
			gen18215 := Call(__e, ShenFunc(symshen_4app), gen18214, MakeString(" \n?- "), MakeSymbol("shen.a"))

			gen18216 := Call(__e, ShenFunc(symcn), MakeString(" inference"), gen18215)

			gen18217 := Call(__e, ShenFunc(symshen_4app), Infs, gen18216, MakeSymbol("shen.a"))

			gen18218 := Call(__e, ShenFunc(symcn), MakeString("____________________________________________________________ "), gen18217)

			gen18219 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(symshen_4prhush), gen18218, gen18219)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.line"), gen18220)

		gen18248 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2814 := __args[0]
			_ = V2814
			gen18246 := Call(__e, ShenFunc(symcons_2), V2814)

			var gen18247 Obj
			if True == gen18246 {
				gen18243 := Call(__e, ShenFunc(symtl), V2814)

				gen18244 := Call(__e, ShenFunc(symcons_2), gen18243)

				var gen18245 Obj
				if True == gen18244 {
					gen18239 := Call(__e, ShenFunc(symtl), V2814)

					gen18240 := Call(__e, ShenFunc(symhd), gen18239)

					gen18241 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen18240)

					var gen18242 Obj
					if True == gen18241 {
						gen18235 := Call(__e, ShenFunc(symtl), V2814)

						gen18236 := Call(__e, ShenFunc(symtl), gen18235)

						gen18237 := Call(__e, ShenFunc(symcons_2), gen18236)

						var gen18238 Obj
						if True == gen18237 {
							gen18231 := Call(__e, ShenFunc(symtl), V2814)

							gen18232 := Call(__e, ShenFunc(symtl), gen18231)

							gen18233 := Call(__e, ShenFunc(symtl), gen18232)

							gen18234 := Call(__e, ShenFunc(sym_a), Nil, gen18233)

							if True == gen18234 {
								gen18238 = True
							} else {
								gen18238 = False
							}

						} else {
							gen18238 = False
						}
						if True == gen18238 {
							gen18242 = True
						} else {
							gen18242 = False
						}

					} else {
						gen18242 = False
					}
					if True == gen18242 {
						gen18245 = True
					} else {
						gen18245 = False
					}

				} else {
					gen18245 = False
				}
				if True == gen18245 {
					gen18247 = True
				} else {
					gen18247 = False
				}

			} else {
				gen18247 = False
			}
			if True == gen18247 {
				gen18223 := Call(__e, ShenFunc(symhd), V2814)

				gen18224 := Call(__e, ShenFunc(symtl), V2814)

				gen18225 := Call(__e, ShenFunc(symtl), gen18224)

				gen18226 := Call(__e, ShenFunc(symhd), gen18225)

				gen18227 := Call(__e, ShenFunc(symshen_4app), gen18226, MakeString(""), MakeSymbol("shen.r"))

				gen18228 := Call(__e, ShenFunc(symcn), MakeString(" : "), gen18227)

				gen18229 := Call(__e, ShenFunc(symshen_4app), gen18223, gen18228, MakeSymbol("shen.r"))

				gen18230 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen18229, gen18230)

				return

			} else {
				gen18221 := Call(__e, ShenFunc(symshen_4app), V2814, MakeString(""), MakeSymbol("shen.r"))

				gen18222 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen18221, gen18222)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show-p"), gen18248)

		gen18256 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2819 := __args[0]
			_ = V2819
			V2820 := __args[1]
			_ = V2820
			gen18255 := Call(__e, ShenFunc(sym_a), Nil, V2819)

			if True == gen18255 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen18254 := Call(__e, ShenFunc(symcons_2), V2819)

				if True == gen18254 {
					gen18249 := Call(__e, ShenFunc(symshen_4app), V2820, MakeString(". "), MakeSymbol("shen.a"))

					gen18250 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), gen18249, gen18250)

					gen18251 := Call(__e, ShenFunc(symhd), V2819)

					Call(__e, ShenFunc(symshen_4show_1p), gen18251)

					Call(__e, ShenFunc(symnl), MakeNumber(1))
					gen18252 := Call(__e, ShenFunc(symtl), V2819)

					gen18253 := Call(__e, ShenFunc(sym_7), V2820, MakeNumber(1))

					__e.TailApply(ShenFunc(symshen_4show_1assumptions), gen18252, gen18253)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.show-assumptions"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show-assumptions"), gen18256)

		gen18260 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen18257 := Call(__e, ShenFunc(symstinput))

			gen18258 := Call(__e, ShenFunc(symread_1byte), gen18257)

			Byte := gen18258
			gen18259 := Call(__e, ShenFunc(sym_a), Byte, MakeNumber(94))

			if True == gen18259 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("input aborted\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symnl), MakeNumber(1))

				return
			}

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pause-for-user"), gen18260)

		gen18263 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2822 := __args[0]
			_ = V2822
			gen18261 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen18262 := Call(__e, ShenFunc(symassoc), V2822, gen18261)

			__e.TailApply(ShenFunc(symcons_2), gen18262)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typedf?"), gen18263)

		gen18264 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2824 := __args[0]
			_ = V2824
			__e.TailApply(ShenFunc(symconcat), MakeSymbol("shen.type-signature-of-"), V2824)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sigf"), gen18264)

		gen18265 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symgensym), MakeSymbol("&&"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.placeholder"), gen18265)

		gen18354 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2829 := __args[0]
			_ = V2829
			V2830 := __args[1]
			_ = V2830
			V2831 := __args[2]
			_ = V2831
			V2832 := __args[3]
			_ = V2832
			gen18266 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

			V2513 := gen18266
			gen18273 := Call(__e, ShenFunc(sym_a), MakeSymbol("number"), V2513)

			var gen18274 Obj
			if True == gen18273 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18271 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

				gen18272 := Call(__e, ShenFunc(symnumber_2), gen18271)

				gen18274 = Call(__e, ShenFunc(symfwhen), gen18272, V2831, V2832)

			} else {
				gen18270 := Call(__e, ShenFunc(symshen_4pvar_2), V2513)

				if True == gen18270 {
					Call(__e, ShenFunc(symshen_4bindv), V2513, MakeSymbol("number"), V2831)
					Call(__e, ShenFunc(symshen_4incinfs))
					gen18267 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

					gen18268 := Call(__e, ShenFunc(symnumber_2), gen18267)

					gen18269 := Call(__e, ShenFunc(symfwhen), gen18268, V2831, V2832)

					Result := gen18269
					Call(__e, ShenFunc(symshen_4unbindv), V2513, V2831)
					gen18274 = Result

				} else {
					gen18274 = False
				}

			}
			Case := gen18274
			gen18353 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18353 {
				gen18275 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

				V2514 := gen18275
				gen18282 := Call(__e, ShenFunc(sym_a), MakeSymbol("boolean"), V2514)

				var gen18283 Obj
				if True == gen18282 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen18280 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

					gen18281 := Call(__e, ShenFunc(symboolean_2), gen18280)

					gen18283 = Call(__e, ShenFunc(symfwhen), gen18281, V2831, V2832)

				} else {
					gen18279 := Call(__e, ShenFunc(symshen_4pvar_2), V2514)

					if True == gen18279 {
						Call(__e, ShenFunc(symshen_4bindv), V2514, MakeSymbol("boolean"), V2831)
						Call(__e, ShenFunc(symshen_4incinfs))
						gen18276 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

						gen18277 := Call(__e, ShenFunc(symboolean_2), gen18276)

						gen18278 := Call(__e, ShenFunc(symfwhen), gen18277, V2831, V2832)

						Result := gen18278
						Call(__e, ShenFunc(symshen_4unbindv), V2514, V2831)
						gen18283 = Result

					} else {
						gen18283 = False
					}

				}
				Case := gen18283
				gen18352 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen18352 {
					gen18284 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

					V2515 := gen18284
					gen18291 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2515)

					var gen18292 Obj
					if True == gen18291 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen18289 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

						gen18290 := Call(__e, ShenFunc(symstring_2), gen18289)

						gen18292 = Call(__e, ShenFunc(symfwhen), gen18290, V2831, V2832)

					} else {
						gen18288 := Call(__e, ShenFunc(symshen_4pvar_2), V2515)

						if True == gen18288 {
							Call(__e, ShenFunc(symshen_4bindv), V2515, MakeSymbol("string"), V2831)
							Call(__e, ShenFunc(symshen_4incinfs))
							gen18285 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							gen18286 := Call(__e, ShenFunc(symstring_2), gen18285)

							gen18287 := Call(__e, ShenFunc(symfwhen), gen18286, V2831, V2832)

							Result := gen18287
							Call(__e, ShenFunc(symshen_4unbindv), V2515, V2831)
							gen18292 = Result

						} else {
							gen18292 = False
						}

					}
					Case := gen18292
					gen18351 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen18351 {
						gen18293 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

						V2516 := gen18293
						gen18308 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2516)

						var gen18309 Obj
						if True == gen18308 {
							Call(__e, ShenFunc(symshen_4incinfs))
							gen18302 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							gen18303 := Call(__e, ShenFunc(symsymbol_2), gen18302)

							gen18307 := MakeNative(func(__e Evaluator, __args ...Obj) {
								gen18304 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

								gen18305 := Call(__e, ShenFunc(symshen_4ue_2), gen18304)

								gen18306 := Call(__e, ShenFunc(symnot), gen18305)

								__e.TailApply(ShenFunc(symfwhen), gen18306, V2831, V2832)

								return

							}, 0)
							gen18309 = Call(__e, ShenFunc(symfwhen), gen18303, V2831, gen18307)

						} else {
							gen18301 := Call(__e, ShenFunc(symshen_4pvar_2), V2516)

							if True == gen18301 {
								Call(__e, ShenFunc(symshen_4bindv), V2516, MakeSymbol("symbol"), V2831)
								Call(__e, ShenFunc(symshen_4incinfs))
								gen18294 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

								gen18295 := Call(__e, ShenFunc(symsymbol_2), gen18294)

								gen18299 := MakeNative(func(__e Evaluator, __args ...Obj) {
									gen18296 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

									gen18297 := Call(__e, ShenFunc(symshen_4ue_2), gen18296)

									gen18298 := Call(__e, ShenFunc(symnot), gen18297)

									__e.TailApply(ShenFunc(symfwhen), gen18298, V2831, V2832)

									return

								}, 0)
								gen18300 := Call(__e, ShenFunc(symfwhen), gen18295, V2831, gen18299)

								Result := gen18300
								Call(__e, ShenFunc(symshen_4unbindv), V2516, V2831)
								gen18309 = Result

							} else {
								gen18309 = False
							}

						}
						Case := gen18309
						gen18350 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen18350 {
							gen18310 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							V2517 := gen18310
							gen18349 := Call(__e, ShenFunc(sym_a), Nil, V2517)

							if True == gen18349 {
								gen18311 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

								V2518 := gen18311
								gen18348 := Call(__e, ShenFunc(symcons_2), V2518)

								if True == gen18348 {
									gen18317 := Call(__e, ShenFunc(symhd), V2518)

									gen18318 := Call(__e, ShenFunc(symshen_4lazyderef), gen18317, V2831)

									V2519 := gen18318
									gen18347 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2519)

									if True == gen18347 {
										gen18334 := Call(__e, ShenFunc(symtl), V2518)

										gen18335 := Call(__e, ShenFunc(symshen_4lazyderef), gen18334, V2831)

										V2520 := gen18335
										gen18346 := Call(__e, ShenFunc(symcons_2), V2520)

										if True == gen18346 {
											gen18340 := Call(__e, ShenFunc(symhd), V2520)

											A := gen18340
											_ = A
											gen18341 := Call(__e, ShenFunc(symtl), V2520)

											gen18342 := Call(__e, ShenFunc(symshen_4lazyderef), gen18341, V2831)

											V2521 := gen18342
											gen18345 := Call(__e, ShenFunc(sym_a), Nil, V2521)

											if True == gen18345 {
												Call(__e, ShenFunc(symshen_4incinfs))
												__e.TailApply(ShenFunc(symthaw), V2832)

												return

											} else {
												gen18344 := Call(__e, ShenFunc(symshen_4pvar_2), V2521)

												if True == gen18344 {
													Call(__e, ShenFunc(symshen_4bindv), V2521, Nil, V2831)
													Call(__e, ShenFunc(symshen_4incinfs))
													gen18343 := Call(__e, ShenFunc(symthaw), V2832)

													Result := gen18343
													Call(__e, ShenFunc(symshen_4unbindv), V2521, V2831)
													__e.Return(Result)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										} else {
											gen18339 := Call(__e, ShenFunc(symshen_4pvar_2), V2520)

											if True == gen18339 {
												gen18336 := Call(__e, ShenFunc(symshen_4newpv), V2831)

												A := gen18336
												gen18337 := Call(__e, ShenFunc(symcons), A, Nil)

												Call(__e, ShenFunc(symshen_4bindv), V2520, gen18337, V2831)

												Call(__e, ShenFunc(symshen_4incinfs))
												gen18338 := Call(__e, ShenFunc(symthaw), V2832)

												Result := gen18338
												Call(__e, ShenFunc(symshen_4unbindv), V2520, V2831)
												__e.Return(Result)
												return

											} else {
												__e.Return(False)
												return
											}

										}

									} else {
										gen18333 := Call(__e, ShenFunc(symshen_4pvar_2), V2519)

										if True == gen18333 {
											Call(__e, ShenFunc(symshen_4bindv), V2519, MakeSymbol("list"), V2831)
											gen18319 := Call(__e, ShenFunc(symtl), V2518)

											gen18320 := Call(__e, ShenFunc(symshen_4lazyderef), gen18319, V2831)

											V2522 := gen18320
											gen18331 := Call(__e, ShenFunc(symcons_2), V2522)

											var gen18332 Obj
											if True == gen18331 {
												gen18325 := Call(__e, ShenFunc(symhd), V2522)

												A := gen18325
												_ = A
												gen18326 := Call(__e, ShenFunc(symtl), V2522)

												gen18327 := Call(__e, ShenFunc(symshen_4lazyderef), gen18326, V2831)

												V2523 := gen18327
												gen18330 := Call(__e, ShenFunc(sym_a), Nil, V2523)

												if True == gen18330 {
													Call(__e, ShenFunc(symshen_4incinfs))
													gen18332 = Call(__e, ShenFunc(symthaw), V2832)

												} else {
													gen18329 := Call(__e, ShenFunc(symshen_4pvar_2), V2523)

													if True == gen18329 {
														Call(__e, ShenFunc(symshen_4bindv), V2523, Nil, V2831)
														Call(__e, ShenFunc(symshen_4incinfs))
														gen18328 := Call(__e, ShenFunc(symthaw), V2832)

														Result := gen18328
														Call(__e, ShenFunc(symshen_4unbindv), V2523, V2831)
														gen18332 = Result

													} else {
														gen18332 = False
													}

												}

											} else {
												gen18324 := Call(__e, ShenFunc(symshen_4pvar_2), V2522)

												if True == gen18324 {
													gen18321 := Call(__e, ShenFunc(symshen_4newpv), V2831)

													A := gen18321
													gen18322 := Call(__e, ShenFunc(symcons), A, Nil)

													Call(__e, ShenFunc(symshen_4bindv), V2522, gen18322, V2831)

													Call(__e, ShenFunc(symshen_4incinfs))
													gen18323 := Call(__e, ShenFunc(symthaw), V2832)

													Result := gen18323
													Call(__e, ShenFunc(symshen_4unbindv), V2522, V2831)
													gen18332 = Result

												} else {
													gen18332 = False
												}

											}
											Result := gen18332
											Call(__e, ShenFunc(symshen_4unbindv), V2519, V2831)
											__e.Return(Result)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								} else {
									gen18316 := Call(__e, ShenFunc(symshen_4pvar_2), V2518)

									if True == gen18316 {
										gen18312 := Call(__e, ShenFunc(symshen_4newpv), V2831)

										A := gen18312
										gen18313 := Call(__e, ShenFunc(symcons), A, Nil)

										gen18314 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen18313)

										Call(__e, ShenFunc(symshen_4bindv), V2518, gen18314, V2831)

										Call(__e, ShenFunc(symshen_4incinfs))
										gen18315 := Call(__e, ShenFunc(symthaw), V2832)

										Result := gen18315
										Call(__e, ShenFunc(symshen_4unbindv), V2518, V2831)
										__e.Return(Result)
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

						} else {
							__e.Return(Case)
							return
						}

					} else {
						__e.Return(Case)
						return
					}

				} else {
					__e.Return(Case)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.base"), gen18354)

		gen18380 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2838 := __args[0]
			_ = V2838
			V2839 := __args[1]
			_ = V2839
			V2840 := __args[2]
			_ = V2840
			V2841 := __args[3]
			_ = V2841
			V2842 := __args[4]
			_ = V2842
			gen18355 := Call(__e, ShenFunc(symshen_4lazyderef), V2840, V2841)

			V2504 := gen18355
			gen18374 := Call(__e, ShenFunc(symcons_2), V2504)

			var gen18375 Obj
			if True == gen18374 {
				gen18356 := Call(__e, ShenFunc(symhd), V2504)

				gen18357 := Call(__e, ShenFunc(symshen_4lazyderef), gen18356, V2841)

				V2505 := gen18357
				gen18373 := Call(__e, ShenFunc(symcons_2), V2505)

				if True == gen18373 {
					gen18358 := Call(__e, ShenFunc(symhd), V2505)

					Y := gen18358
					gen18359 := Call(__e, ShenFunc(symtl), V2505)

					gen18360 := Call(__e, ShenFunc(symshen_4lazyderef), gen18359, V2841)

					V2506 := gen18360
					gen18372 := Call(__e, ShenFunc(symcons_2), V2506)

					if True == gen18372 {
						gen18361 := Call(__e, ShenFunc(symhd), V2506)

						gen18362 := Call(__e, ShenFunc(symshen_4lazyderef), gen18361, V2841)

						V2507 := gen18362
						gen18371 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2507)

						if True == gen18371 {
							gen18363 := Call(__e, ShenFunc(symtl), V2506)

							gen18364 := Call(__e, ShenFunc(symshen_4lazyderef), gen18363, V2841)

							V2508 := gen18364
							gen18370 := Call(__e, ShenFunc(symcons_2), V2508)

							if True == gen18370 {
								gen18365 := Call(__e, ShenFunc(symhd), V2508)

								B := gen18365
								gen18366 := Call(__e, ShenFunc(symtl), V2508)

								gen18367 := Call(__e, ShenFunc(symshen_4lazyderef), gen18366, V2841)

								V2509 := gen18367
								gen18369 := Call(__e, ShenFunc(sym_a), Nil, V2509)

								if True == gen18369 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen18368 := MakeNative(func(__e Evaluator, __args ...Obj) {
										__e.TailApply(ShenFunc(symunify_b), V2839, B, V2841, V2842)

										return
									}, 0)
									gen18375 = Call(__e, ShenFunc(symidentical), V2838, Y, V2841, gen18368)

								} else {
									gen18375 = False
								}

							} else {
								gen18375 = False
							}

						} else {
							gen18375 = False
						}

					} else {
						gen18375 = False
					}

				} else {
					gen18375 = False
				}

			} else {
				gen18375 = False
			}
			Case := gen18375
			gen18379 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18379 {
				gen18376 := Call(__e, ShenFunc(symshen_4lazyderef), V2840, V2841)

				V2510 := gen18376
				gen18378 := Call(__e, ShenFunc(symcons_2), V2510)

				if True == gen18378 {
					gen18377 := Call(__e, ShenFunc(symtl), V2510)

					Hyp := gen18377
					Call(__e, ShenFunc(symshen_4incinfs))
					__e.TailApply(ShenFunc(symshen_4by__hypothesis), V2838, V2839, Hyp, V2841, V2842)

					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.by_hypothesis"), gen18380)

		gen18399 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2848 := __args[0]
			_ = V2848
			V2849 := __args[1]
			_ = V2849
			V2850 := __args[2]
			_ = V2850
			V2851 := __args[3]
			_ = V2851
			V2852 := __args[4]
			_ = V2852
			gen18381 := Call(__e, ShenFunc(symshen_4lazyderef), V2848, V2851)

			V2498 := gen18381
			gen18398 := Call(__e, ShenFunc(symcons_2), V2498)

			if True == gen18398 {
				gen18382 := Call(__e, ShenFunc(symhd), V2498)

				gen18383 := Call(__e, ShenFunc(symshen_4lazyderef), gen18382, V2851)

				V2499 := gen18383
				gen18397 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), V2499)

				if True == gen18397 {
					gen18384 := Call(__e, ShenFunc(symtl), V2498)

					gen18385 := Call(__e, ShenFunc(symshen_4lazyderef), gen18384, V2851)

					V2500 := gen18385
					gen18396 := Call(__e, ShenFunc(symcons_2), V2500)

					if True == gen18396 {
						gen18386 := Call(__e, ShenFunc(symhd), V2500)

						F := gen18386
						gen18387 := Call(__e, ShenFunc(symtl), V2500)

						X := gen18387
						gen18388 := Call(__e, ShenFunc(symshen_4newpv), V2851)

						Y := gen18388
						_ = Y
						gen18389 := Call(__e, ShenFunc(symshen_4newpv), V2851)

						E := gen18389
						_ = E
						Call(__e, ShenFunc(symshen_4incinfs))
						gen18390 := MakeNative(func(__e Evaluator, __args ...Obj) {
							Y := __args[0]
							_ = Y
							__e.TailApply(ShenFunc(symshen_4_5sig_7rules_6), Y)

							return
						}, 1)
						gen18394 := MakeNative(func(__e Evaluator, __args ...Obj) {
							E := __args[0]
							_ = E
							gen18393 := Call(__e, ShenFunc(symcons_2), E)

							if True == gen18393 {
								gen18391 := Call(__e, ShenFunc(symshen_4app), E, MakeString("\n"), MakeSymbol("shen.s"))

								gen18392 := Call(__e, ShenFunc(symcn), MakeString("parse error here: "), gen18391)

								__e.TailApply(ShenFunc(symsimple_1error), gen18392)

								return

							} else {
								__e.TailApply(ShenFunc(symsimple_1error), MakeString("parse error\n"))

								return
							}

						}, 1)
						gen18395 := Call(__e, ShenFunc(symcompile), gen18390, X, gen18394)

						__e.TailApply(ShenFunc(symshen_4t_d_1defh), gen18395, F, V2849, V2850, V2851, V2852)

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
				__e.Return(False)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-def"), gen18399)

		gen18405 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2859 := __args[0]
			_ = V2859
			V2860 := __args[1]
			_ = V2860
			V2861 := __args[2]
			_ = V2861
			V2862 := __args[3]
			_ = V2862
			V2863 := __args[4]
			_ = V2863
			V2864 := __args[5]
			_ = V2864
			gen18400 := Call(__e, ShenFunc(symshen_4lazyderef), V2859, V2863)

			V2494 := gen18400
			gen18404 := Call(__e, ShenFunc(symcons_2), V2494)

			if True == gen18404 {
				gen18401 := Call(__e, ShenFunc(symhd), V2494)

				Sig := gen18401
				gen18402 := Call(__e, ShenFunc(symtl), V2494)

				Rules := gen18402
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18403 := Call(__e, ShenFunc(symshen_4ue_1sig), Sig)

				__e.TailApply(ShenFunc(symshen_4t_d_1defhh), Sig, gen18403, V2860, V2861, V2862, Rules, V2863, V2864)

				return

			} else {
				__e.Return(False)
				return
			}

		}, 6)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-defh"), gen18405)

		gen18411 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2873 := __args[0]
			_ = V2873
			V2874 := __args[1]
			_ = V2874
			V2875 := __args[2]
			_ = V2875
			V2876 := __args[3]
			_ = V2876
			V2877 := __args[4]
			_ = V2877
			V2878 := __args[5]
			_ = V2878
			V2879 := __args[6]
			_ = V2879
			V2880 := __args[7]
			_ = V2880
			Call(__e, ShenFunc(symshen_4incinfs))
			gen18406 := Call(__e, ShenFunc(symcons), V2874, Nil)

			gen18407 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18406)

			gen18408 := Call(__e, ShenFunc(symcons), V2875, gen18407)

			gen18409 := Call(__e, ShenFunc(symcons), gen18408, V2877)

			gen18410 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.TailApply(ShenFunc(symshen_4memo), V2875, V2873, V2876, V2879, V2880)

				return
			}, 0)
			__e.TailApply(ShenFunc(symshen_4t_d_1rules), V2878, V2874, MakeNumber(1), V2875, gen18409, V2879, gen18410)

			return

		}, 8)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-defhh"), gen18411)

		gen18417 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2886 := __args[0]
			_ = V2886
			V2887 := __args[1]
			_ = V2887
			V2888 := __args[2]
			_ = V2888
			V2889 := __args[3]
			_ = V2889
			V2890 := __args[4]
			_ = V2890
			gen18412 := Call(__e, ShenFunc(symshen_4newpv), V2889)

			Jnk := gen18412
			Call(__e, ShenFunc(symshen_4incinfs))
			gen18416 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen18413 := Call(__e, ShenFunc(symshen_4lazyderef), V2886, V2889)

				gen18414 := Call(__e, ShenFunc(symshen_4lazyderef), V2888, V2889)

				gen18415 := Call(__e, ShenFunc(symdeclare), gen18413, gen18414)

				__e.TailApply(ShenFunc(symbind), Jnk, gen18415, V2889, V2890)

				return

			}, 0)
			__e.TailApply(ShenFunc(symunify_b), V2888, V2887, V2889, gen18416)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.memo"), gen18417)

		gen18430 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2892 := __args[0]
			_ = V2892
			gen18418 := Call(__e, ShenFunc(symshen_4_5signature_6), V2892)

			Parse__shen_4_5signature_6 := gen18418
			gen18427 := Call(__e, ShenFunc(symfail))

			gen18428 := Call(__e, ShenFunc(sym_a), gen18427, Parse__shen_4_5signature_6)

			gen18429 := Call(__e, ShenFunc(symnot), gen18428)

			if True == gen18429 {
				gen18419 := Call(__e, ShenFunc(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5signature_6)

				Parse__shen_4_5non_1ll_1rules_6 := gen18419
				gen18424 := Call(__e, ShenFunc(symfail))

				gen18425 := Call(__e, ShenFunc(sym_a), gen18424, Parse__shen_4_5non_1ll_1rules_6)

				gen18426 := Call(__e, ShenFunc(symnot), gen18425)

				if True == gen18426 {
					gen18420 := Call(__e, ShenFunc(symhd), Parse__shen_4_5non_1ll_1rules_6)

					gen18421 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_6)

					gen18422 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

					gen18423 := Call(__e, ShenFunc(symcons), gen18421, gen18422)

					__e.TailApply(ShenFunc(symshen_4pair), gen18420, gen18423)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<sig+rules>"), gen18430)

		gen18453 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2894 := __args[0]
			_ = V2894
			gen18431 := Call(__e, ShenFunc(symshen_4_5rule_6), V2894)

			Parse__shen_4_5rule_6 := gen18431
			gen18440 := Call(__e, ShenFunc(symfail))

			gen18441 := Call(__e, ShenFunc(sym_a), gen18440, Parse__shen_4_5rule_6)

			gen18442 := Call(__e, ShenFunc(symnot), gen18441)

			var gen18443 Obj
			if True == gen18442 {
				gen18432 := Call(__e, ShenFunc(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5rule_6)

				Parse__shen_4_5non_1ll_1rules_6 := gen18432
				gen18437 := Call(__e, ShenFunc(symfail))

				gen18438 := Call(__e, ShenFunc(sym_a), gen18437, Parse__shen_4_5non_1ll_1rules_6)

				gen18439 := Call(__e, ShenFunc(symnot), gen18438)

				if True == gen18439 {
					gen18433 := Call(__e, ShenFunc(symhd), Parse__shen_4_5non_1ll_1rules_6)

					gen18434 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen18435 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

					gen18436 := Call(__e, ShenFunc(symcons), gen18434, gen18435)

					gen18443 = Call(__e, ShenFunc(symshen_4pair), gen18433, gen18436)

				} else {
					gen18443 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen18443 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen18443
			gen18451 := Call(__e, ShenFunc(symfail))

			gen18452 := Call(__e, ShenFunc(sym_a), YaccParse, gen18451)

			if True == gen18452 {
				gen18444 := Call(__e, ShenFunc(symshen_4_5rule_6), V2894)

				Parse__shen_4_5rule_6 := gen18444
				gen18448 := Call(__e, ShenFunc(symfail))

				gen18449 := Call(__e, ShenFunc(sym_a), gen18448, Parse__shen_4_5rule_6)

				gen18450 := Call(__e, ShenFunc(symnot), gen18449)

				if True == gen18450 {
					gen18445 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rule_6)

					gen18446 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen18447 := Call(__e, ShenFunc(symcons), gen18446, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen18445, gen18447)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<non-ll-rules>"), gen18453)

		gen18468 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2896 := __args[0]
			_ = V2896
			gen18466 := Call(__e, ShenFunc(symcons_2), V2896)

			var gen18467 Obj
			if True == gen18466 {
				gen18463 := Call(__e, ShenFunc(symtl), V2896)

				gen18464 := Call(__e, ShenFunc(symcons_2), gen18463)

				var gen18465 Obj
				if True == gen18464 {
					gen18459 := Call(__e, ShenFunc(symtl), V2896)

					gen18460 := Call(__e, ShenFunc(symtl), gen18459)

					gen18461 := Call(__e, ShenFunc(sym_a), Nil, gen18460)

					var gen18462 Obj
					if True == gen18461 {
						gen18457 := Call(__e, ShenFunc(symhd), V2896)

						gen18458 := Call(__e, ShenFunc(sym_a), gen18457, MakeSymbol("protect"))

						if True == gen18458 {
							gen18462 = True
						} else {
							gen18462 = False
						}

					} else {
						gen18462 = False
					}
					if True == gen18462 {
						gen18465 = True
					} else {
						gen18465 = False
					}

				} else {
					gen18465 = False
				}
				if True == gen18465 {
					gen18467 = True
				} else {
					gen18467 = False
				}

			} else {
				gen18467 = False
			}
			if True == gen18467 {
				__e.Return(V2896)
				return
			} else {
				gen18456 := Call(__e, ShenFunc(symcons_2), V2896)

				if True == gen18456 {
					gen18455 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Z := __args[0]
						_ = Z
						__e.TailApply(ShenFunc(symshen_4ue), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen18455, V2896)

					return

				} else {
					gen18454 := Call(__e, ShenFunc(symvariable_2), V2896)

					if True == gen18454 {
						__e.TailApply(ShenFunc(symconcat), MakeSymbol("&&"), V2896)

						return
					} else {
						__e.Return(V2896)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue"), gen18468)

		gen18472 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2898 := __args[0]
			_ = V2898
			gen18471 := Call(__e, ShenFunc(symcons_2), V2898)

			if True == gen18471 {
				gen18470 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__e.TailApply(ShenFunc(symshen_4ue_1sig), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symmap), gen18470, V2898)

				return

			} else {
				gen18469 := Call(__e, ShenFunc(symvariable_2), V2898)

				if True == gen18469 {
					__e.TailApply(ShenFunc(symconcat), MakeSymbol("&&&"), V2898)

					return
				} else {
					__e.Return(V2898)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue-sig"), gen18472)

		gen18479 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2904 := __args[0]
			_ = V2904
			gen18478 := Call(__e, ShenFunc(symshen_4ue_2), V2904)

			if True == gen18478 {
				__e.TailApply(ShenFunc(symcons), V2904, Nil)

				return
			} else {
				gen18477 := Call(__e, ShenFunc(symcons_2), V2904)

				if True == gen18477 {
					gen18473 := Call(__e, ShenFunc(symhd), V2904)

					gen18474 := Call(__e, ShenFunc(symshen_4ues), gen18473)

					gen18475 := Call(__e, ShenFunc(symtl), V2904)

					gen18476 := Call(__e, ShenFunc(symshen_4ues), gen18475)

					__e.TailApply(ShenFunc(symunion), gen18474, gen18476)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ues"), gen18479)

		gen18483 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2906 := __args[0]
			_ = V2906
			gen18482 := Call(__e, ShenFunc(symsymbol_2), V2906)

			if True == gen18482 {
				gen18480 := Call(__e, ShenFunc(symstr), V2906)

				gen18481 := Call(__e, ShenFunc(symshen_4ue_1h_2), gen18480)

				if True == gen18481 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue?"), gen18483)

		gen18495 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2914 := __args[0]
			_ = V2914
			gen18493 := Call(__e, ShenFunc(symshen_4_7string_2), V2914)

			var gen18494 Obj
			if True == gen18493 {
				gen18490 := Call(__e, ShenFunc(sympos), V2914, MakeNumber(0))

				gen18491 := Call(__e, ShenFunc(sym_a), MakeString("&"), gen18490)

				var gen18492 Obj
				if True == gen18491 {
					gen18487 := Call(__e, ShenFunc(symtlstr), V2914)

					gen18488 := Call(__e, ShenFunc(symshen_4_7string_2), gen18487)

					var gen18489 Obj
					if True == gen18488 {
						gen18484 := Call(__e, ShenFunc(symtlstr), V2914)

						gen18485 := Call(__e, ShenFunc(sympos), gen18484, MakeNumber(0))

						gen18486 := Call(__e, ShenFunc(sym_a), MakeString("&"), gen18485)

						if True == gen18486 {
							gen18489 = True
						} else {
							gen18489 = False
						}

					} else {
						gen18489 = False
					}
					if True == gen18489 {
						gen18492 = True
					} else {
						gen18492 = False
					}

				} else {
					gen18492 = False
				}
				if True == gen18492 {
					gen18494 = True
				} else {
					gen18494 = False
				}

			} else {
				gen18494 = False
			}
			if True == gen18494 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue-h?"), gen18495)

		gen18520 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2922 := __args[0]
			_ = V2922
			V2923 := __args[1]
			_ = V2923
			V2924 := __args[2]
			_ = V2924
			V2925 := __args[3]
			_ = V2925
			V2926 := __args[4]
			_ = V2926
			V2927 := __args[5]
			_ = V2927
			V2928 := __args[6]
			_ = V2928
			gen18496 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen18496
			gen18497 := Call(__e, ShenFunc(symshen_4lazyderef), V2922, V2927)

			V2478 := gen18497
			gen18498 := Call(__e, ShenFunc(sym_a), Nil, V2478)

			var gen18499 Obj
			if True == gen18498 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18499 = Call(__e, ShenFunc(symthaw), V2928)

			} else {
				gen18499 = False
			}
			Case := gen18499
			gen18518 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen18519 Obj
			if True == gen18518 {
				gen18500 := Call(__e, ShenFunc(symshen_4lazyderef), V2922, V2927)

				V2479 := gen18500
				gen18507 := Call(__e, ShenFunc(symcons_2), V2479)

				var gen18508 Obj
				if True == gen18507 {
					gen18501 := Call(__e, ShenFunc(symhd), V2479)

					Rule := gen18501
					gen18502 := Call(__e, ShenFunc(symtl), V2479)

					Rules := gen18502
					Call(__e, ShenFunc(symshen_4incinfs))
					gen18503 := Call(__e, ShenFunc(symshen_4ue), Rule)

					gen18506 := MakeNative(func(__e Evaluator, __args ...Obj) {
						gen18505 := MakeNative(func(__e Evaluator, __args ...Obj) {
							gen18504 := Call(__e, ShenFunc(sym_7), V2924, MakeNumber(1))

							__e.TailApply(ShenFunc(symshen_4t_d_1rules), Rules, V2923, gen18504, V2925, V2926, V2927, V2928)

							return

						}, 0)
						__e.TailApply(ShenFunc(symcut), Throwcontrol, V2927, gen18505)

						return

					}, 0)
					gen18508 = Call(__e, ShenFunc(symshen_4t_d_1rule), gen18503, V2923, V2926, V2927, gen18506)

				} else {
					gen18508 = False
				}
				Case := gen18508
				gen18517 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen18517 {
					gen18509 := Call(__e, ShenFunc(symshen_4newpv), V2927)

					Err := gen18509
					Call(__e, ShenFunc(symshen_4incinfs))
					gen18510 := Call(__e, ShenFunc(symshen_4lazyderef), V2924, V2927)

					gen18511 := Call(__e, ShenFunc(symshen_4lazyderef), V2925, V2927)

					gen18512 := Call(__e, ShenFunc(symshen_4app), gen18511, MakeString(""), MakeSymbol("shen.a"))

					gen18513 := Call(__e, ShenFunc(symcn), MakeString(" of "), gen18512)

					gen18514 := Call(__e, ShenFunc(symshen_4app), gen18510, gen18513, MakeSymbol("shen.a"))

					gen18515 := Call(__e, ShenFunc(symcn), MakeString("type error in rule "), gen18514)

					gen18516 := Call(__e, ShenFunc(symsimple_1error), gen18515)

					gen18519 = Call(__e, ShenFunc(symbind), Err, gen18516, V2927, V2928)

				} else {
					gen18519 = Case
				}

			} else {
				gen18519 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen18519)

			return

		}, 7)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-rules"), gen18520)

		gen18542 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2934 := __args[0]
			_ = V2934
			V2935 := __args[1]
			_ = V2935
			V2936 := __args[2]
			_ = V2936
			V2937 := __args[3]
			_ = V2937
			V2938 := __args[4]
			_ = V2938
			gen18521 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen18521
			gen18522 := Call(__e, ShenFunc(symshen_4lazyderef), V2934, V2937)

			V2470 := gen18522
			gen18540 := Call(__e, ShenFunc(symcons_2), V2470)

			var gen18541 Obj
			if True == gen18540 {
				gen18523 := Call(__e, ShenFunc(symhd), V2470)

				Patterns := gen18523
				gen18524 := Call(__e, ShenFunc(symtl), V2470)

				gen18525 := Call(__e, ShenFunc(symshen_4lazyderef), gen18524, V2937)

				V2471 := gen18525
				gen18539 := Call(__e, ShenFunc(symcons_2), V2471)

				if True == gen18539 {
					gen18526 := Call(__e, ShenFunc(symhd), V2471)

					Action := gen18526
					gen18527 := Call(__e, ShenFunc(symtl), V2471)

					gen18528 := Call(__e, ShenFunc(symshen_4lazyderef), gen18527, V2937)

					V2472 := gen18528
					gen18538 := Call(__e, ShenFunc(sym_a), Nil, V2472)

					if True == gen18538 {
						gen18529 := Call(__e, ShenFunc(symshen_4newpv), V2937)

						NewHyps := gen18529
						Call(__e, ShenFunc(symshen_4incinfs))
						gen18530 := Call(__e, ShenFunc(symshen_4placeholders), Patterns)

						gen18537 := MakeNative(func(__e Evaluator, __args ...Obj) {
							gen18536 := MakeNative(func(__e Evaluator, __args ...Obj) {
								gen18535 := MakeNative(func(__e Evaluator, __args ...Obj) {
									gen18531 := Call(__e, ShenFunc(symshen_4ue), Action)

									gen18532 := Call(__e, ShenFunc(symshen_4curry), gen18531)

									gen18533 := Call(__e, ShenFunc(symshen_4result_1type), Patterns, V2935)

									gen18534 := Call(__e, ShenFunc(symshen_4patthyps), Patterns, V2935, V2936)

									__e.TailApply(ShenFunc(symshen_4t_d_1action), gen18532, gen18533, gen18534, V2937, V2938)

									return

								}, 0)
								__e.TailApply(ShenFunc(symcut), Throwcontrol, V2937, gen18535)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4t_d_1patterns), Patterns, V2935, NewHyps, V2937, gen18536)

							return

						}, 0)
						gen18541 = Call(__e, ShenFunc(symshen_4newhyps), gen18530, V2936, NewHyps, V2937, gen18537)

					} else {
						gen18541 = False
					}

				} else {
					gen18541 = False
				}

			} else {
				gen18541 = False
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen18541)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-rule"), gen18542)

		gen18549 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2944 := __args[0]
			_ = V2944
			gen18548 := Call(__e, ShenFunc(symshen_4ue_2), V2944)

			if True == gen18548 {
				__e.TailApply(ShenFunc(symcons), V2944, Nil)

				return
			} else {
				gen18547 := Call(__e, ShenFunc(symcons_2), V2944)

				if True == gen18547 {
					gen18543 := Call(__e, ShenFunc(symhd), V2944)

					gen18544 := Call(__e, ShenFunc(symshen_4placeholders), gen18543)

					gen18545 := Call(__e, ShenFunc(symtl), V2944)

					gen18546 := Call(__e, ShenFunc(symshen_4placeholders), gen18545)

					__e.TailApply(ShenFunc(symunion), gen18544, gen18546)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.placeholders"), gen18549)

		gen18636 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2950 := __args[0]
			_ = V2950
			V2951 := __args[1]
			_ = V2951
			V2952 := __args[2]
			_ = V2952
			V2953 := __args[3]
			_ = V2953
			V2954 := __args[4]
			_ = V2954
			gen18550 := Call(__e, ShenFunc(symshen_4lazyderef), V2950, V2953)

			V2457 := gen18550
			gen18551 := Call(__e, ShenFunc(sym_a), Nil, V2457)

			var gen18552 Obj
			if True == gen18551 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18552 = Call(__e, ShenFunc(symunify_b), V2952, V2951, V2953, V2954)

			} else {
				gen18552 = False
			}
			Case := gen18552
			gen18635 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18635 {
				gen18553 := Call(__e, ShenFunc(symshen_4lazyderef), V2950, V2953)

				V2458 := gen18553
				gen18634 := Call(__e, ShenFunc(symcons_2), V2458)

				if True == gen18634 {
					gen18554 := Call(__e, ShenFunc(symhd), V2458)

					V2453 := gen18554
					gen18555 := Call(__e, ShenFunc(symtl), V2458)

					Vs := gen18555
					gen18556 := Call(__e, ShenFunc(symshen_4lazyderef), V2952, V2953)

					V2459 := gen18556
					gen18633 := Call(__e, ShenFunc(symcons_2), V2459)

					if True == gen18633 {
						gen18567 := Call(__e, ShenFunc(symhd), V2459)

						gen18568 := Call(__e, ShenFunc(symshen_4lazyderef), gen18567, V2953)

						V2460 := gen18568
						gen18632 := Call(__e, ShenFunc(symcons_2), V2460)

						if True == gen18632 {
							gen18578 := Call(__e, ShenFunc(symhd), V2460)

							V := gen18578
							gen18579 := Call(__e, ShenFunc(symtl), V2460)

							gen18580 := Call(__e, ShenFunc(symshen_4lazyderef), gen18579, V2953)

							V2461 := gen18580
							gen18631 := Call(__e, ShenFunc(symcons_2), V2461)

							if True == gen18631 {
								gen18588 := Call(__e, ShenFunc(symhd), V2461)

								gen18589 := Call(__e, ShenFunc(symshen_4lazyderef), gen18588, V2953)

								V2462 := gen18589
								gen18630 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2462)

								if True == gen18630 {
									gen18611 := Call(__e, ShenFunc(symtl), V2461)

									gen18612 := Call(__e, ShenFunc(symshen_4lazyderef), gen18611, V2953)

									V2463 := gen18612
									gen18629 := Call(__e, ShenFunc(symcons_2), V2463)

									if True == gen18629 {
										gen18619 := Call(__e, ShenFunc(symhd), V2463)

										A := gen18619
										_ = A
										gen18620 := Call(__e, ShenFunc(symtl), V2463)

										gen18621 := Call(__e, ShenFunc(symshen_4lazyderef), gen18620, V2953)

										V2464 := gen18621
										gen18628 := Call(__e, ShenFunc(sym_a), Nil, V2464)

										if True == gen18628 {
											gen18626 := Call(__e, ShenFunc(symtl), V2459)

											NewHyp := gen18626
											Call(__e, ShenFunc(symshen_4incinfs))
											gen18627 := MakeNative(func(__e Evaluator, __args ...Obj) {
												__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

												return
											}, 0)
											__e.TailApply(ShenFunc(symunify_b), V, V2453, V2953, gen18627)

											return

										} else {
											gen18625 := Call(__e, ShenFunc(symshen_4pvar_2), V2464)

											if True == gen18625 {
												Call(__e, ShenFunc(symshen_4bindv), V2464, Nil, V2953)
												gen18622 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen18622
												Call(__e, ShenFunc(symshen_4incinfs))
												gen18623 := MakeNative(func(__e Evaluator, __args ...Obj) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen18624 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18623)

												Result := gen18624
												Call(__e, ShenFunc(symshen_4unbindv), V2464, V2953)
												__e.Return(Result)
												return

											} else {
												__e.Return(False)
												return
											}

										}

									} else {
										gen18618 := Call(__e, ShenFunc(symshen_4pvar_2), V2463)

										if True == gen18618 {
											gen18613 := Call(__e, ShenFunc(symshen_4newpv), V2953)

											A := gen18613
											gen18614 := Call(__e, ShenFunc(symcons), A, Nil)

											Call(__e, ShenFunc(symshen_4bindv), V2463, gen18614, V2953)

											gen18615 := Call(__e, ShenFunc(symtl), V2459)

											NewHyp := gen18615
											Call(__e, ShenFunc(symshen_4incinfs))
											gen18616 := MakeNative(func(__e Evaluator, __args ...Obj) {
												__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

												return
											}, 0)
											gen18617 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18616)

											Result := gen18617
											Call(__e, ShenFunc(symshen_4unbindv), V2463, V2953)
											__e.Return(Result)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								} else {
									gen18610 := Call(__e, ShenFunc(symshen_4pvar_2), V2462)

									if True == gen18610 {
										Call(__e, ShenFunc(symshen_4bindv), V2462, MakeSymbol(":"), V2953)
										gen18590 := Call(__e, ShenFunc(symtl), V2461)

										gen18591 := Call(__e, ShenFunc(symshen_4lazyderef), gen18590, V2953)

										V2465 := gen18591
										gen18608 := Call(__e, ShenFunc(symcons_2), V2465)

										var gen18609 Obj
										if True == gen18608 {
											gen18598 := Call(__e, ShenFunc(symhd), V2465)

											A := gen18598
											_ = A
											gen18599 := Call(__e, ShenFunc(symtl), V2465)

											gen18600 := Call(__e, ShenFunc(symshen_4lazyderef), gen18599, V2953)

											V2466 := gen18600
											gen18607 := Call(__e, ShenFunc(sym_a), Nil, V2466)

											if True == gen18607 {
												gen18605 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen18605
												Call(__e, ShenFunc(symshen_4incinfs))
												gen18606 := MakeNative(func(__e Evaluator, __args ...Obj) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen18609 = Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18606)

											} else {
												gen18604 := Call(__e, ShenFunc(symshen_4pvar_2), V2466)

												if True == gen18604 {
													Call(__e, ShenFunc(symshen_4bindv), V2466, Nil, V2953)
													gen18601 := Call(__e, ShenFunc(symtl), V2459)

													NewHyp := gen18601
													Call(__e, ShenFunc(symshen_4incinfs))
													gen18602 := MakeNative(func(__e Evaluator, __args ...Obj) {
														__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

														return
													}, 0)
													gen18603 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18602)

													Result := gen18603
													Call(__e, ShenFunc(symshen_4unbindv), V2466, V2953)
													gen18609 = Result

												} else {
													gen18609 = False
												}

											}

										} else {
											gen18597 := Call(__e, ShenFunc(symshen_4pvar_2), V2465)

											if True == gen18597 {
												gen18592 := Call(__e, ShenFunc(symshen_4newpv), V2953)

												A := gen18592
												gen18593 := Call(__e, ShenFunc(symcons), A, Nil)

												Call(__e, ShenFunc(symshen_4bindv), V2465, gen18593, V2953)

												gen18594 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen18594
												Call(__e, ShenFunc(symshen_4incinfs))
												gen18595 := MakeNative(func(__e Evaluator, __args ...Obj) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen18596 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18595)

												Result := gen18596
												Call(__e, ShenFunc(symshen_4unbindv), V2465, V2953)
												gen18609 = Result

											} else {
												gen18609 = False
											}

										}
										Result := gen18609
										Call(__e, ShenFunc(symshen_4unbindv), V2462, V2953)
										__e.Return(Result)
										return

									} else {
										__e.Return(False)
										return
									}

								}

							} else {
								gen18587 := Call(__e, ShenFunc(symshen_4pvar_2), V2461)

								if True == gen18587 {
									gen18581 := Call(__e, ShenFunc(symshen_4newpv), V2953)

									A := gen18581
									gen18582 := Call(__e, ShenFunc(symcons), A, Nil)

									gen18583 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18582)

									Call(__e, ShenFunc(symshen_4bindv), V2461, gen18583, V2953)

									gen18584 := Call(__e, ShenFunc(symtl), V2459)

									NewHyp := gen18584
									Call(__e, ShenFunc(symshen_4incinfs))
									gen18585 := MakeNative(func(__e Evaluator, __args ...Obj) {
										__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

										return
									}, 0)
									gen18586 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18585)

									Result := gen18586
									Call(__e, ShenFunc(symshen_4unbindv), V2461, V2953)
									__e.Return(Result)
									return

								} else {
									__e.Return(False)
									return
								}

							}

						} else {
							gen18577 := Call(__e, ShenFunc(symshen_4pvar_2), V2460)

							if True == gen18577 {
								gen18569 := Call(__e, ShenFunc(symshen_4newpv), V2953)

								V := gen18569
								gen18570 := Call(__e, ShenFunc(symshen_4newpv), V2953)

								A := gen18570
								gen18571 := Call(__e, ShenFunc(symcons), A, Nil)

								gen18572 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18571)

								gen18573 := Call(__e, ShenFunc(symcons), V, gen18572)

								Call(__e, ShenFunc(symshen_4bindv), V2460, gen18573, V2953)

								gen18574 := Call(__e, ShenFunc(symtl), V2459)

								NewHyp := gen18574
								Call(__e, ShenFunc(symshen_4incinfs))
								gen18575 := MakeNative(func(__e Evaluator, __args ...Obj) {
									__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

									return
								}, 0)
								gen18576 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18575)

								Result := gen18576
								Call(__e, ShenFunc(symshen_4unbindv), V2460, V2953)
								__e.Return(Result)
								return

							} else {
								__e.Return(False)
								return
							}

						}

					} else {
						gen18566 := Call(__e, ShenFunc(symshen_4pvar_2), V2459)

						if True == gen18566 {
							gen18557 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							V := gen18557
							gen18558 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							A := gen18558
							gen18559 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							NewHyp := gen18559
							gen18560 := Call(__e, ShenFunc(symcons), A, Nil)

							gen18561 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18560)

							gen18562 := Call(__e, ShenFunc(symcons), V, gen18561)

							gen18563 := Call(__e, ShenFunc(symcons), gen18562, NewHyp)

							Call(__e, ShenFunc(symshen_4bindv), V2459, gen18563, V2953)

							Call(__e, ShenFunc(symshen_4incinfs))
							gen18564 := MakeNative(func(__e Evaluator, __args ...Obj) {
								__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

								return
							}, 0)
							gen18565 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen18564)

							Result := gen18565
							Call(__e, ShenFunc(symshen_4unbindv), V2459, V2953)
							__e.Return(Result)
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

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newhyps"), gen18636)

		gen18668 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2960 := __args[0]
			_ = V2960
			V2961 := __args[1]
			_ = V2961
			V2962 := __args[2]
			_ = V2962
			gen18667 := Call(__e, ShenFunc(sym_a), Nil, V2960)

			if True == gen18667 {
				__e.Return(V2962)
				return
			} else {
				gen18665 := Call(__e, ShenFunc(symcons_2), V2960)

				var gen18666 Obj
				if True == gen18665 {
					gen18663 := Call(__e, ShenFunc(symcons_2), V2961)

					var gen18664 Obj
					if True == gen18663 {
						gen18660 := Call(__e, ShenFunc(symtl), V2961)

						gen18661 := Call(__e, ShenFunc(symcons_2), gen18660)

						var gen18662 Obj
						if True == gen18661 {
							gen18656 := Call(__e, ShenFunc(symtl), V2961)

							gen18657 := Call(__e, ShenFunc(symhd), gen18656)

							gen18658 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen18657)

							var gen18659 Obj
							if True == gen18658 {
								gen18652 := Call(__e, ShenFunc(symtl), V2961)

								gen18653 := Call(__e, ShenFunc(symtl), gen18652)

								gen18654 := Call(__e, ShenFunc(symcons_2), gen18653)

								var gen18655 Obj
								if True == gen18654 {
									gen18648 := Call(__e, ShenFunc(symtl), V2961)

									gen18649 := Call(__e, ShenFunc(symtl), gen18648)

									gen18650 := Call(__e, ShenFunc(symtl), gen18649)

									gen18651 := Call(__e, ShenFunc(sym_a), Nil, gen18650)

									if True == gen18651 {
										gen18655 = True
									} else {
										gen18655 = False
									}

								} else {
									gen18655 = False
								}
								if True == gen18655 {
									gen18659 = True
								} else {
									gen18659 = False
								}

							} else {
								gen18659 = False
							}
							if True == gen18659 {
								gen18662 = True
							} else {
								gen18662 = False
							}

						} else {
							gen18662 = False
						}
						if True == gen18662 {
							gen18664 = True
						} else {
							gen18664 = False
						}

					} else {
						gen18664 = False
					}
					if True == gen18664 {
						gen18666 = True
					} else {
						gen18666 = False
					}

				} else {
					gen18666 = False
				}
				if True == gen18666 {
					gen18637 := Call(__e, ShenFunc(symhd), V2960)

					gen18638 := Call(__e, ShenFunc(symshen_4curry), gen18637)

					gen18639 := Call(__e, ShenFunc(symhd), V2961)

					gen18640 := Call(__e, ShenFunc(symcons), gen18639, Nil)

					gen18641 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18640)

					gen18642 := Call(__e, ShenFunc(symcons), gen18638, gen18641)

					gen18643 := Call(__e, ShenFunc(symtl), V2960)

					gen18644 := Call(__e, ShenFunc(symtl), V2961)

					gen18645 := Call(__e, ShenFunc(symtl), gen18644)

					gen18646 := Call(__e, ShenFunc(symhd), gen18645)

					gen18647 := Call(__e, ShenFunc(symshen_4patthyps), gen18643, gen18646, V2962)

					__e.TailApply(ShenFunc(symadjoin), gen18642, gen18647)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.patthyps"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.patthyps"), gen18668)

		gen18707 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2969 := __args[0]
			_ = V2969
			V2970 := __args[1]
			_ = V2970
			gen18705 := Call(__e, ShenFunc(sym_a), Nil, V2969)

			var gen18706 Obj
			if True == gen18705 {
				gen18703 := Call(__e, ShenFunc(symcons_2), V2970)

				var gen18704 Obj
				if True == gen18703 {
					gen18700 := Call(__e, ShenFunc(symhd), V2970)

					gen18701 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen18700)

					var gen18702 Obj
					if True == gen18701 {
						gen18697 := Call(__e, ShenFunc(symtl), V2970)

						gen18698 := Call(__e, ShenFunc(symcons_2), gen18697)

						var gen18699 Obj
						if True == gen18698 {
							gen18694 := Call(__e, ShenFunc(symtl), V2970)

							gen18695 := Call(__e, ShenFunc(symtl), gen18694)

							gen18696 := Call(__e, ShenFunc(sym_a), Nil, gen18695)

							if True == gen18696 {
								gen18699 = True
							} else {
								gen18699 = False
							}

						} else {
							gen18699 = False
						}
						if True == gen18699 {
							gen18702 = True
						} else {
							gen18702 = False
						}

					} else {
						gen18702 = False
					}
					if True == gen18702 {
						gen18704 = True
					} else {
						gen18704 = False
					}

				} else {
					gen18704 = False
				}
				if True == gen18704 {
					gen18706 = True
				} else {
					gen18706 = False
				}

			} else {
				gen18706 = False
			}
			if True == gen18706 {
				gen18693 := Call(__e, ShenFunc(symtl), V2970)

				__e.TailApply(ShenFunc(symhd), gen18693)

				return

			} else {
				gen18692 := Call(__e, ShenFunc(sym_a), Nil, V2969)

				if True == gen18692 {
					__e.Return(V2970)
					return
				} else {
					gen18690 := Call(__e, ShenFunc(symcons_2), V2969)

					var gen18691 Obj
					if True == gen18690 {
						gen18688 := Call(__e, ShenFunc(symcons_2), V2970)

						var gen18689 Obj
						if True == gen18688 {
							gen18685 := Call(__e, ShenFunc(symtl), V2970)

							gen18686 := Call(__e, ShenFunc(symcons_2), gen18685)

							var gen18687 Obj
							if True == gen18686 {
								gen18681 := Call(__e, ShenFunc(symtl), V2970)

								gen18682 := Call(__e, ShenFunc(symhd), gen18681)

								gen18683 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen18682)

								var gen18684 Obj
								if True == gen18683 {
									gen18677 := Call(__e, ShenFunc(symtl), V2970)

									gen18678 := Call(__e, ShenFunc(symtl), gen18677)

									gen18679 := Call(__e, ShenFunc(symcons_2), gen18678)

									var gen18680 Obj
									if True == gen18679 {
										gen18673 := Call(__e, ShenFunc(symtl), V2970)

										gen18674 := Call(__e, ShenFunc(symtl), gen18673)

										gen18675 := Call(__e, ShenFunc(symtl), gen18674)

										gen18676 := Call(__e, ShenFunc(sym_a), Nil, gen18675)

										if True == gen18676 {
											gen18680 = True
										} else {
											gen18680 = False
										}

									} else {
										gen18680 = False
									}
									if True == gen18680 {
										gen18684 = True
									} else {
										gen18684 = False
									}

								} else {
									gen18684 = False
								}
								if True == gen18684 {
									gen18687 = True
								} else {
									gen18687 = False
								}

							} else {
								gen18687 = False
							}
							if True == gen18687 {
								gen18689 = True
							} else {
								gen18689 = False
							}

						} else {
							gen18689 = False
						}
						if True == gen18689 {
							gen18691 = True
						} else {
							gen18691 = False
						}

					} else {
						gen18691 = False
					}
					if True == gen18691 {
						gen18669 := Call(__e, ShenFunc(symtl), V2969)

						gen18670 := Call(__e, ShenFunc(symtl), V2970)

						gen18671 := Call(__e, ShenFunc(symtl), gen18670)

						gen18672 := Call(__e, ShenFunc(symhd), gen18671)

						__e.TailApply(ShenFunc(symshen_4result_1type), gen18669, gen18672)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.result-type"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.result-type"), gen18707)

		gen18737 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2976 := __args[0]
			_ = V2976
			V2977 := __args[1]
			_ = V2977
			V2978 := __args[2]
			_ = V2978
			V2979 := __args[3]
			_ = V2979
			V2980 := __args[4]
			_ = V2980
			gen18708 := Call(__e, ShenFunc(symshen_4lazyderef), V2976, V2979)

			V2445 := gen18708
			gen18709 := Call(__e, ShenFunc(sym_a), Nil, V2445)

			var gen18710 Obj
			if True == gen18709 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18710 = Call(__e, ShenFunc(symthaw), V2980)

			} else {
				gen18710 = False
			}
			Case := gen18710
			gen18736 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18736 {
				gen18711 := Call(__e, ShenFunc(symshen_4lazyderef), V2976, V2979)

				V2446 := gen18711
				gen18735 := Call(__e, ShenFunc(symcons_2), V2446)

				if True == gen18735 {
					gen18712 := Call(__e, ShenFunc(symhd), V2446)

					Pattern := gen18712
					gen18713 := Call(__e, ShenFunc(symtl), V2446)

					Patterns := gen18713
					gen18714 := Call(__e, ShenFunc(symshen_4lazyderef), V2977, V2979)

					V2447 := gen18714
					gen18734 := Call(__e, ShenFunc(symcons_2), V2447)

					if True == gen18734 {
						gen18715 := Call(__e, ShenFunc(symhd), V2447)

						A := gen18715
						gen18716 := Call(__e, ShenFunc(symtl), V2447)

						gen18717 := Call(__e, ShenFunc(symshen_4lazyderef), gen18716, V2979)

						V2448 := gen18717
						gen18733 := Call(__e, ShenFunc(symcons_2), V2448)

						if True == gen18733 {
							gen18718 := Call(__e, ShenFunc(symhd), V2448)

							gen18719 := Call(__e, ShenFunc(symshen_4lazyderef), gen18718, V2979)

							V2449 := gen18719
							gen18732 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), V2449)

							if True == gen18732 {
								gen18720 := Call(__e, ShenFunc(symtl), V2448)

								gen18721 := Call(__e, ShenFunc(symshen_4lazyderef), gen18720, V2979)

								V2450 := gen18721
								gen18731 := Call(__e, ShenFunc(symcons_2), V2450)

								if True == gen18731 {
									gen18722 := Call(__e, ShenFunc(symhd), V2450)

									B := gen18722
									gen18723 := Call(__e, ShenFunc(symtl), V2450)

									gen18724 := Call(__e, ShenFunc(symshen_4lazyderef), gen18723, V2979)

									V2451 := gen18724
									gen18730 := Call(__e, ShenFunc(sym_a), Nil, V2451)

									if True == gen18730 {
										Call(__e, ShenFunc(symshen_4incinfs))
										gen18725 := Call(__e, ShenFunc(symshen_4curry), Pattern)

										gen18726 := Call(__e, ShenFunc(symcons), A, Nil)

										gen18727 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18726)

										gen18728 := Call(__e, ShenFunc(symcons), gen18725, gen18727)

										gen18729 := MakeNative(func(__e Evaluator, __args ...Obj) {
											__e.TailApply(ShenFunc(symshen_4t_d_1patterns), Patterns, B, V2978, V2979, V2980)

											return
										}, 0)
										__e.TailApply(ShenFunc(symshen_4t_d), gen18728, V2978, V2979, gen18729)

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
								__e.Return(False)
								return
							}

						} else {
							__e.Return(False)
							return
						}

					} else {
						__e.Return(False)
						return
					}

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-patterns"), gen18737)

		gen18840 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2986 := __args[0]
			_ = V2986
			V2987 := __args[1]
			_ = V2987
			V2988 := __args[2]
			_ = V2988
			V2989 := __args[3]
			_ = V2989
			V2990 := __args[4]
			_ = V2990
			gen18738 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen18738
			gen18739 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

			V2422 := gen18739
			gen18764 := Call(__e, ShenFunc(symcons_2), V2422)

			var gen18765 Obj
			if True == gen18764 {
				gen18740 := Call(__e, ShenFunc(symhd), V2422)

				gen18741 := Call(__e, ShenFunc(symshen_4lazyderef), gen18740, V2989)

				V2423 := gen18741
				gen18763 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), V2423)

				if True == gen18763 {
					gen18742 := Call(__e, ShenFunc(symtl), V2422)

					gen18743 := Call(__e, ShenFunc(symshen_4lazyderef), gen18742, V2989)

					V2424 := gen18743
					gen18762 := Call(__e, ShenFunc(symcons_2), V2424)

					if True == gen18762 {
						gen18744 := Call(__e, ShenFunc(symhd), V2424)

						P := gen18744
						gen18745 := Call(__e, ShenFunc(symtl), V2424)

						gen18746 := Call(__e, ShenFunc(symshen_4lazyderef), gen18745, V2989)

						V2425 := gen18746
						gen18761 := Call(__e, ShenFunc(symcons_2), V2425)

						if True == gen18761 {
							gen18747 := Call(__e, ShenFunc(symhd), V2425)

							Action := gen18747
							gen18748 := Call(__e, ShenFunc(symtl), V2425)

							gen18749 := Call(__e, ShenFunc(symshen_4lazyderef), gen18748, V2989)

							V2426 := gen18749
							gen18760 := Call(__e, ShenFunc(sym_a), Nil, V2426)

							if True == gen18760 {
								Call(__e, ShenFunc(symshen_4incinfs))
								gen18759 := MakeNative(func(__e Evaluator, __args ...Obj) {
									gen18750 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

									gen18751 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18750)

									gen18752 := Call(__e, ShenFunc(symcons), P, gen18751)

									gen18758 := MakeNative(func(__e Evaluator, __args ...Obj) {
										gen18757 := MakeNative(func(__e Evaluator, __args ...Obj) {
											gen18753 := Call(__e, ShenFunc(symcons), MakeSymbol("verified"), Nil)

											gen18754 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18753)

											gen18755 := Call(__e, ShenFunc(symcons), P, gen18754)

											gen18756 := Call(__e, ShenFunc(symcons), gen18755, V2988)

											__e.TailApply(ShenFunc(symshen_4t_d_1action), Action, V2987, gen18756, V2989, V2990)

											return

										}, 0)
										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2989, gen18757)

										return

									}, 0)
									__e.TailApply(ShenFunc(symshen_4t_d), gen18752, V2988, V2989, gen18758)

									return

								}, 0)
								gen18765 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen18759)

							} else {
								gen18765 = False
							}

						} else {
							gen18765 = False
						}

					} else {
						gen18765 = False
					}

				} else {
					gen18765 = False
				}

			} else {
				gen18765 = False
			}
			Case := gen18765
			gen18838 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen18839 Obj
			if True == gen18838 {
				gen18766 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

				V2427 := gen18766
				gen18807 := Call(__e, ShenFunc(symcons_2), V2427)

				var gen18808 Obj
				if True == gen18807 {
					gen18767 := Call(__e, ShenFunc(symhd), V2427)

					gen18768 := Call(__e, ShenFunc(symshen_4lazyderef), gen18767, V2989)

					V2428 := gen18768
					gen18806 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), V2428)

					if True == gen18806 {
						gen18769 := Call(__e, ShenFunc(symtl), V2427)

						gen18770 := Call(__e, ShenFunc(symshen_4lazyderef), gen18769, V2989)

						V2429 := gen18770
						gen18805 := Call(__e, ShenFunc(symcons_2), V2429)

						if True == gen18805 {
							gen18771 := Call(__e, ShenFunc(symhd), V2429)

							gen18772 := Call(__e, ShenFunc(symshen_4lazyderef), gen18771, V2989)

							V2430 := gen18772
							gen18804 := Call(__e, ShenFunc(symcons_2), V2430)

							if True == gen18804 {
								gen18773 := Call(__e, ShenFunc(symhd), V2430)

								gen18774 := Call(__e, ShenFunc(symshen_4lazyderef), gen18773, V2989)

								V2431 := gen18774
								gen18803 := Call(__e, ShenFunc(symcons_2), V2431)

								if True == gen18803 {
									gen18775 := Call(__e, ShenFunc(symhd), V2431)

									gen18776 := Call(__e, ShenFunc(symshen_4lazyderef), gen18775, V2989)

									V2432 := gen18776
									gen18802 := Call(__e, ShenFunc(sym_a), MakeSymbol("fail-if"), V2432)

									if True == gen18802 {
										gen18777 := Call(__e, ShenFunc(symtl), V2431)

										gen18778 := Call(__e, ShenFunc(symshen_4lazyderef), gen18777, V2989)

										V2433 := gen18778
										gen18801 := Call(__e, ShenFunc(symcons_2), V2433)

										if True == gen18801 {
											gen18779 := Call(__e, ShenFunc(symhd), V2433)

											F := gen18779
											gen18780 := Call(__e, ShenFunc(symtl), V2433)

											gen18781 := Call(__e, ShenFunc(symshen_4lazyderef), gen18780, V2989)

											V2434 := gen18781
											gen18800 := Call(__e, ShenFunc(sym_a), Nil, V2434)

											if True == gen18800 {
												gen18782 := Call(__e, ShenFunc(symtl), V2430)

												gen18783 := Call(__e, ShenFunc(symshen_4lazyderef), gen18782, V2989)

												V2435 := gen18783
												gen18799 := Call(__e, ShenFunc(symcons_2), V2435)

												if True == gen18799 {
													gen18784 := Call(__e, ShenFunc(symhd), V2435)

													Action := gen18784
													gen18785 := Call(__e, ShenFunc(symtl), V2435)

													gen18786 := Call(__e, ShenFunc(symshen_4lazyderef), gen18785, V2989)

													V2436 := gen18786
													gen18798 := Call(__e, ShenFunc(sym_a), Nil, V2436)

													if True == gen18798 {
														gen18787 := Call(__e, ShenFunc(symtl), V2429)

														gen18788 := Call(__e, ShenFunc(symshen_4lazyderef), gen18787, V2989)

														V2437 := gen18788
														gen18797 := Call(__e, ShenFunc(sym_a), Nil, V2437)

														if True == gen18797 {
															Call(__e, ShenFunc(symshen_4incinfs))
															gen18796 := MakeNative(func(__e Evaluator, __args ...Obj) {
																gen18789 := Call(__e, ShenFunc(symcons), Action, Nil)

																gen18790 := Call(__e, ShenFunc(symcons), F, gen18789)

																gen18791 := Call(__e, ShenFunc(symcons), gen18790, Nil)

																gen18792 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen18791)

																gen18793 := Call(__e, ShenFunc(symcons), Action, Nil)

																gen18794 := Call(__e, ShenFunc(symcons), gen18792, gen18793)

																gen18795 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen18794)

																__e.TailApply(ShenFunc(symshen_4t_d_1action), gen18795, V2987, V2988, V2989, V2990)

																return

															}, 0)
															gen18808 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen18796)

														} else {
															gen18808 = False
														}

													} else {
														gen18808 = False
													}

												} else {
													gen18808 = False
												}

											} else {
												gen18808 = False
											}

										} else {
											gen18808 = False
										}

									} else {
										gen18808 = False
									}

								} else {
									gen18808 = False
								}

							} else {
								gen18808 = False
							}

						} else {
							gen18808 = False
						}

					} else {
						gen18808 = False
					}

				} else {
					gen18808 = False
				}
				Case := gen18808
				gen18837 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen18837 {
					gen18809 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

					V2438 := gen18809
					gen18831 := Call(__e, ShenFunc(symcons_2), V2438)

					var gen18832 Obj
					if True == gen18831 {
						gen18810 := Call(__e, ShenFunc(symhd), V2438)

						gen18811 := Call(__e, ShenFunc(symshen_4lazyderef), gen18810, V2989)

						V2439 := gen18811
						gen18830 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), V2439)

						if True == gen18830 {
							gen18812 := Call(__e, ShenFunc(symtl), V2438)

							gen18813 := Call(__e, ShenFunc(symshen_4lazyderef), gen18812, V2989)

							V2440 := gen18813
							gen18829 := Call(__e, ShenFunc(symcons_2), V2440)

							if True == gen18829 {
								gen18814 := Call(__e, ShenFunc(symhd), V2440)

								Action := gen18814
								gen18815 := Call(__e, ShenFunc(symtl), V2440)

								gen18816 := Call(__e, ShenFunc(symshen_4lazyderef), gen18815, V2989)

								V2441 := gen18816
								gen18828 := Call(__e, ShenFunc(sym_a), Nil, V2441)

								if True == gen18828 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen18827 := MakeNative(func(__e Evaluator, __args ...Obj) {
										gen18817 := Call(__e, ShenFunc(symcons), Action, Nil)

										gen18818 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen18817)

										gen18819 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

										gen18820 := Call(__e, ShenFunc(symcons), gen18819, Nil)

										gen18821 := Call(__e, ShenFunc(symcons), gen18818, gen18820)

										gen18822 := Call(__e, ShenFunc(symcons), gen18821, Nil)

										gen18823 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen18822)

										gen18824 := Call(__e, ShenFunc(symcons), Action, Nil)

										gen18825 := Call(__e, ShenFunc(symcons), gen18823, gen18824)

										gen18826 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen18825)

										__e.TailApply(ShenFunc(symshen_4t_d_1action), gen18826, V2987, V2988, V2989, V2990)

										return

									}, 0)
									gen18832 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen18827)

								} else {
									gen18832 = False
								}

							} else {
								gen18832 = False
							}

						} else {
							gen18832 = False
						}

					} else {
						gen18832 = False
					}
					Case := gen18832
					gen18836 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen18836 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen18833 := Call(__e, ShenFunc(symcons), V2987, Nil)

						gen18834 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen18833)

						gen18835 := Call(__e, ShenFunc(symcons), V2986, gen18834)

						gen18839 = Call(__e, ShenFunc(symshen_4t_d), gen18835, V2988, V2989, V2990)

					} else {
						gen18839 = Case
					}

				} else {
					gen18839 = Case
				}

			} else {
				gen18839 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen18839)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-action"), gen18840)

		gen18848 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V2996 := __args[0]
			_ = V2996
			V2997 := __args[1]
			_ = V2997
			V2998 := __args[2]
			_ = V2998
			V2999 := __args[3]
			_ = V2999
			V3000 := __args[4]
			_ = V3000
			gen18841 := Call(__e, ShenFunc(symshen_4newpv), V2999)

			B := gen18841
			gen18842 := Call(__e, ShenFunc(symshen_4newpv), V2999)

			A := gen18842
			Call(__e, ShenFunc(symshen_4incinfs))
			gen18843 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.a"))

			gen18847 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen18844 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2999)

				gen18845 := Call(__e, ShenFunc(symset), gen18844, Nil)

				gen18846 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(symshen_4findallhelp), V2996, V2997, V2998, A, V2999, V3000)

					return
				}, 0)
				__e.TailApply(ShenFunc(symbind), B, gen18845, V2999, gen18846)

				return

			}, 0)
			__e.TailApply(ShenFunc(symbind), A, gen18843, V2999, gen18847)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("findall"), gen18848)

		gen18855 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3007 := __args[0]
			_ = V3007
			V3008 := __args[1]
			_ = V3008
			V3009 := __args[2]
			_ = V3009
			V3010 := __args[3]
			_ = V3010
			V3011 := __args[4]
			_ = V3011
			V3012 := __args[5]
			_ = V3012
			Call(__e, ShenFunc(symshen_4incinfs))
			gen18850 := MakeNative(func(__e Evaluator, __args ...Obj) {
				gen18849 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(symfwhen), False, V3011, V3012)

					return
				}, 0)
				__e.TailApply(ShenFunc(symshen_4remember), V3010, V3007, V3011, gen18849)

				return

			}, 0)
			gen18851 := Call(__e, ShenFunc(symcall), V3008, V3011, gen18850)

			Case := gen18851
			gen18854 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen18854 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen18852 := Call(__e, ShenFunc(symshen_4lazyderef), V3010, V3011)

				gen18853 := Call(__e, ShenFunc(symvalue), gen18852)

				__e.TailApply(ShenFunc(symbind), V3009, gen18853, V3011, V3012)

				return

			} else {
				__e.Return(Case)
				return
			}

		}, 6)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.findallhelp"), gen18855)

		gen18863 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3017 := __args[0]
			_ = V3017
			V3018 := __args[1]
			_ = V3018
			V3019 := __args[2]
			_ = V3019
			V3020 := __args[3]
			_ = V3020
			gen18856 := Call(__e, ShenFunc(symshen_4newpv), V3019)

			B := gen18856
			Call(__e, ShenFunc(symshen_4incinfs))
			gen18857 := Call(__e, ShenFunc(symshen_4deref), V3017, V3019)

			gen18858 := Call(__e, ShenFunc(symshen_4deref), V3018, V3019)

			gen18859 := Call(__e, ShenFunc(symshen_4deref), V3017, V3019)

			gen18860 := Call(__e, ShenFunc(symvalue), gen18859)

			gen18861 := Call(__e, ShenFunc(symcons), gen18858, gen18860)

			gen18862 := Call(__e, ShenFunc(symset), gen18857, gen18861)

			__e.TailApply(ShenFunc(symbind), B, gen18862, V3019, V3020)

			return

		}, 4)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.remember"), gen18863)

		return

	}, 0))
}
