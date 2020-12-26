package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen10279 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4724 := __args[0]
			_ = V4724
			gen10277 := Call(__e, ShenFunc(symcons_2), V4724)

			var gen10278 Obj
			if True == gen10277 {
				gen10274 := Call(__e, ShenFunc(symhd), V4724)

				gen10275 := Call(__e, ShenFunc(sym_a), MakeSymbol("defcc"), gen10274)

				var gen10276 Obj
				if True == gen10275 {
					gen10272 := Call(__e, ShenFunc(symtl), V4724)

					gen10273 := Call(__e, ShenFunc(symcons_2), gen10272)

					if True == gen10273 {
						gen10276 = True
					} else {
						gen10276 = False
					}

				} else {
					gen10276 = False
				}
				if True == gen10276 {
					gen10278 = True
				} else {
					gen10278 = False
				}

			} else {
				gen10278 = False
			}
			if True == gen10278 {
				gen10268 := Call(__e, ShenFunc(symtl), V4724)

				gen10269 := Call(__e, ShenFunc(symhd), gen10268)

				gen10270 := Call(__e, ShenFunc(symtl), V4724)

				gen10271 := Call(__e, ShenFunc(symtl), gen10270)

				__e.TailApply(ShenFunc(symshen_4yacc_1_6shen), gen10269, gen10271)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.yacc"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc"), gen10279)

		gen10289 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4727 := __args[0]
			_ = V4727
			V4728 := __args[1]
			_ = V4728
			gen10280 := Call(__e, ShenFunc(symshen_4split__cc__rules), True, V4728, Nil)

			CCRules := gen10280
			gen10281 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4cc__body), X)

				return
			}, 1)
			gen10282 := Call(__e, ShenFunc(symmap), gen10281, CCRules)

			CCBody := gen10282
			gen10283 := Call(__e, ShenFunc(symshen_4yacc__cases), CCBody)

			YaccCases := gen10283
			gen10284 := Call(__e, ShenFunc(symshen_4kill_1code), YaccCases)

			gen10285 := Call(__e, ShenFunc(symcons), gen10284, Nil)

			gen10286 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen10285)

			gen10287 := Call(__e, ShenFunc(symcons), MakeSymbol("Stream"), gen10286)

			gen10288 := Call(__e, ShenFunc(symcons), V4727, gen10287)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("define"), gen10288)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc->shen"), gen10289)

		gen10299 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4730 := __args[0]
			_ = V4730
			gen10297 := Call(__e, ShenFunc(symoccurrences), MakeSymbol("kill"), V4730)

			gen10298 := Call(__e, ShenFunc(sym_6), gen10297, MakeNumber(0))

			if True == gen10298 {
				gen10290 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen10291 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.analyse-kill"), gen10290)

				gen10292 := Call(__e, ShenFunc(symcons), gen10291, Nil)

				gen10293 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen10292)

				gen10294 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen10293)

				gen10295 := Call(__e, ShenFunc(symcons), gen10294, Nil)

				gen10296 := Call(__e, ShenFunc(symcons), V4730, gen10295)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("trap-error"), gen10296)

				return

			} else {
				__e.Return(V4730)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.kill-code"), gen10299)

		gen10300 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString("yacc kill"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("kill"), gen10300)

		gen10303 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4732 := __args[0]
			_ = V4732
			gen10301 := Call(__e, ShenFunc(symerror_1to_1string), V4732)

			String := gen10301
			gen10302 := Call(__e, ShenFunc(sym_a), String, MakeString("yacc kill"))

			if True == gen10302 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V4732)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-kill"), gen10303)

		gen10322 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4738 := __args[0]
			_ = V4738
			V4739 := __args[1]
			_ = V4739
			V4740 := __args[2]
			_ = V4740
			gen10320 := Call(__e, ShenFunc(sym_a), Nil, V4739)

			var gen10321 Obj
			if True == gen10320 {
				gen10319 := Call(__e, ShenFunc(sym_a), Nil, V4740)

				if True == gen10319 {
					gen10321 = True
				} else {
					gen10321 = False
				}

			} else {
				gen10321 = False
			}
			if True == gen10321 {
				__e.Return(Nil)
				return
			} else {
				gen10318 := Call(__e, ShenFunc(sym_a), Nil, V4739)

				if True == gen10318 {
					gen10316 := Call(__e, ShenFunc(symreverse), V4740)

					gen10317 := Call(__e, ShenFunc(symshen_4split__cc__rule), V4738, gen10316, Nil)

					__e.TailApply(ShenFunc(symcons), gen10317, Nil)

					return

				} else {
					gen10314 := Call(__e, ShenFunc(symcons_2), V4739)

					var gen10315 Obj
					if True == gen10314 {
						gen10312 := Call(__e, ShenFunc(symhd), V4739)

						gen10313 := Call(__e, ShenFunc(sym_a), MakeSymbol(";"), gen10312)

						if True == gen10313 {
							gen10315 = True
						} else {
							gen10315 = False
						}

					} else {
						gen10315 = False
					}
					if True == gen10315 {
						gen10308 := Call(__e, ShenFunc(symreverse), V4740)

						gen10309 := Call(__e, ShenFunc(symshen_4split__cc__rule), V4738, gen10308, Nil)

						gen10310 := Call(__e, ShenFunc(symtl), V4739)

						gen10311 := Call(__e, ShenFunc(symshen_4split__cc__rules), V4738, gen10310, Nil)

						__e.TailApply(ShenFunc(symcons), gen10309, gen10311)

						return

					} else {
						gen10307 := Call(__e, ShenFunc(symcons_2), V4739)

						if True == gen10307 {
							gen10304 := Call(__e, ShenFunc(symtl), V4739)

							gen10305 := Call(__e, ShenFunc(symhd), V4739)

							gen10306 := Call(__e, ShenFunc(symcons), gen10305, V4740)

							__e.TailApply(ShenFunc(symshen_4split__cc__rules), V4738, gen10304, gen10306)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.split_cc_rules"))

							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.split_cc_rules"), gen10322)

		gen10383 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4748 := __args[0]
			_ = V4748
			V4749 := __args[1]
			_ = V4749
			V4750 := __args[2]
			_ = V4750
			gen10381 := Call(__e, ShenFunc(symcons_2), V4749)

			var gen10382 Obj
			if True == gen10381 {
				gen10378 := Call(__e, ShenFunc(symhd), V4749)

				gen10379 := Call(__e, ShenFunc(sym_a), MakeSymbol(":="), gen10378)

				var gen10380 Obj
				if True == gen10379 {
					gen10375 := Call(__e, ShenFunc(symtl), V4749)

					gen10376 := Call(__e, ShenFunc(symcons_2), gen10375)

					var gen10377 Obj
					if True == gen10376 {
						gen10372 := Call(__e, ShenFunc(symtl), V4749)

						gen10373 := Call(__e, ShenFunc(symtl), gen10372)

						gen10374 := Call(__e, ShenFunc(sym_a), Nil, gen10373)

						if True == gen10374 {
							gen10377 = True
						} else {
							gen10377 = False
						}

					} else {
						gen10377 = False
					}
					if True == gen10377 {
						gen10380 = True
					} else {
						gen10380 = False
					}

				} else {
					gen10380 = False
				}
				if True == gen10380 {
					gen10382 = True
				} else {
					gen10382 = False
				}

			} else {
				gen10382 = False
			}
			if True == gen10382 {
				gen10370 := Call(__e, ShenFunc(symreverse), V4750)

				gen10371 := Call(__e, ShenFunc(symtl), V4749)

				__e.TailApply(ShenFunc(symcons), gen10370, gen10371)

				return

			} else {
				gen10368 := Call(__e, ShenFunc(symcons_2), V4749)

				var gen10369 Obj
				if True == gen10368 {
					gen10365 := Call(__e, ShenFunc(symhd), V4749)

					gen10366 := Call(__e, ShenFunc(sym_a), MakeSymbol(":="), gen10365)

					var gen10367 Obj
					if True == gen10366 {
						gen10362 := Call(__e, ShenFunc(symtl), V4749)

						gen10363 := Call(__e, ShenFunc(symcons_2), gen10362)

						var gen10364 Obj
						if True == gen10363 {
							gen10358 := Call(__e, ShenFunc(symtl), V4749)

							gen10359 := Call(__e, ShenFunc(symtl), gen10358)

							gen10360 := Call(__e, ShenFunc(symcons_2), gen10359)

							var gen10361 Obj
							if True == gen10360 {
								gen10353 := Call(__e, ShenFunc(symtl), V4749)

								gen10354 := Call(__e, ShenFunc(symtl), gen10353)

								gen10355 := Call(__e, ShenFunc(symhd), gen10354)

								gen10356 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen10355)

								var gen10357 Obj
								if True == gen10356 {
									gen10348 := Call(__e, ShenFunc(symtl), V4749)

									gen10349 := Call(__e, ShenFunc(symtl), gen10348)

									gen10350 := Call(__e, ShenFunc(symtl), gen10349)

									gen10351 := Call(__e, ShenFunc(symcons_2), gen10350)

									var gen10352 Obj
									if True == gen10351 {
										gen10343 := Call(__e, ShenFunc(symtl), V4749)

										gen10344 := Call(__e, ShenFunc(symtl), gen10343)

										gen10345 := Call(__e, ShenFunc(symtl), gen10344)

										gen10346 := Call(__e, ShenFunc(symtl), gen10345)

										gen10347 := Call(__e, ShenFunc(sym_a), Nil, gen10346)

										if True == gen10347 {
											gen10352 = True
										} else {
											gen10352 = False
										}

									} else {
										gen10352 = False
									}
									if True == gen10352 {
										gen10357 = True
									} else {
										gen10357 = False
									}

								} else {
									gen10357 = False
								}
								if True == gen10357 {
									gen10361 = True
								} else {
									gen10361 = False
								}

							} else {
								gen10361 = False
							}
							if True == gen10361 {
								gen10364 = True
							} else {
								gen10364 = False
							}

						} else {
							gen10364 = False
						}
						if True == gen10364 {
							gen10367 = True
						} else {
							gen10367 = False
						}

					} else {
						gen10367 = False
					}
					if True == gen10367 {
						gen10369 = True
					} else {
						gen10369 = False
					}

				} else {
					gen10369 = False
				}
				if True == gen10369 {
					gen10332 := Call(__e, ShenFunc(symreverse), V4750)

					gen10333 := Call(__e, ShenFunc(symtl), V4749)

					gen10334 := Call(__e, ShenFunc(symtl), gen10333)

					gen10335 := Call(__e, ShenFunc(symtl), gen10334)

					gen10336 := Call(__e, ShenFunc(symhd), gen10335)

					gen10337 := Call(__e, ShenFunc(symtl), V4749)

					gen10338 := Call(__e, ShenFunc(symhd), gen10337)

					gen10339 := Call(__e, ShenFunc(symcons), gen10338, Nil)

					gen10340 := Call(__e, ShenFunc(symcons), gen10336, gen10339)

					gen10341 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen10340)

					gen10342 := Call(__e, ShenFunc(symcons), gen10341, Nil)

					__e.TailApply(ShenFunc(symcons), gen10332, gen10342)

					return

				} else {
					gen10331 := Call(__e, ShenFunc(sym_a), Nil, V4749)

					if True == gen10331 {
						Call(__e, ShenFunc(symshen_4semantic_1completion_1warning), V4748, V4750)
						gen10327 := Call(__e, ShenFunc(symreverse), V4750)

						gen10328 := Call(__e, ShenFunc(symshen_4default__semantics), gen10327)

						gen10329 := Call(__e, ShenFunc(symcons), gen10328, Nil)

						gen10330 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen10329)

						__e.TailApply(ShenFunc(symshen_4split__cc__rule), V4748, gen10330, V4750)

						return

					} else {
						gen10326 := Call(__e, ShenFunc(symcons_2), V4749)

						if True == gen10326 {
							gen10323 := Call(__e, ShenFunc(symtl), V4749)

							gen10324 := Call(__e, ShenFunc(symhd), V4749)

							gen10325 := Call(__e, ShenFunc(symcons), gen10324, V4750)

							__e.TailApply(ShenFunc(symshen_4split__cc__rule), V4748, gen10323, gen10325)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.split_cc_rule"))

							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.split_cc_rule"), gen10383)

		gen10391 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4761 := __args[0]
			_ = V4761
			V4762 := __args[1]
			_ = V4762
			gen10390 := Call(__e, ShenFunc(sym_a), True, V4761)

			if True == gen10390 {
				gen10384 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("warning: "), gen10384)

				gen10387 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					gen10385 := Call(__e, ShenFunc(symshen_4app), X, MakeString(" "), MakeSymbol("shen.a"))

					gen10386 := Call(__e, ShenFunc(symstoutput))

					__e.TailApply(ShenFunc(symshen_4prhush), gen10385, gen10386)

					return

				}, 1)
				gen10388 := Call(__e, ShenFunc(symreverse), V4762)

				Call(__e, ShenFunc(symshen_4for_1each), gen10387, gen10388)

				gen10389 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), MakeString("has no semantics.\n"), gen10389)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.semantic-completion-warning"), gen10391)

		gen10415 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4764 := __args[0]
			_ = V4764
			gen10414 := Call(__e, ShenFunc(sym_a), Nil, V4764)

			if True == gen10414 {
				__e.Return(Nil)
				return
			} else {
				gen10412 := Call(__e, ShenFunc(symcons_2), V4764)

				var gen10413 Obj
				if True == gen10412 {
					gen10409 := Call(__e, ShenFunc(symtl), V4764)

					gen10410 := Call(__e, ShenFunc(sym_a), Nil, gen10409)

					var gen10411 Obj
					if True == gen10410 {
						gen10407 := Call(__e, ShenFunc(symhd), V4764)

						gen10408 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen10407)

						if True == gen10408 {
							gen10411 = True
						} else {
							gen10411 = False
						}

					} else {
						gen10411 = False
					}
					if True == gen10411 {
						gen10413 = True
					} else {
						gen10413 = False
					}

				} else {
					gen10413 = False
				}
				if True == gen10413 {
					__e.TailApply(ShenFunc(symhd), V4764)

					return
				} else {
					gen10405 := Call(__e, ShenFunc(symcons_2), V4764)

					var gen10406 Obj
					if True == gen10405 {
						gen10403 := Call(__e, ShenFunc(symhd), V4764)

						gen10404 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen10403)

						if True == gen10404 {
							gen10406 = True
						} else {
							gen10406 = False
						}

					} else {
						gen10406 = False
					}
					if True == gen10406 {
						gen10398 := Call(__e, ShenFunc(symhd), V4764)

						gen10399 := Call(__e, ShenFunc(symtl), V4764)

						gen10400 := Call(__e, ShenFunc(symshen_4default__semantics), gen10399)

						gen10401 := Call(__e, ShenFunc(symcons), gen10400, Nil)

						gen10402 := Call(__e, ShenFunc(symcons), gen10398, gen10401)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("append"), gen10402)

						return

					} else {
						gen10397 := Call(__e, ShenFunc(symcons_2), V4764)

						if True == gen10397 {
							gen10392 := Call(__e, ShenFunc(symhd), V4764)

							gen10393 := Call(__e, ShenFunc(symtl), V4764)

							gen10394 := Call(__e, ShenFunc(symshen_4default__semantics), gen10393)

							gen10395 := Call(__e, ShenFunc(symcons), gen10394, Nil)

							gen10396 := Call(__e, ShenFunc(symcons), gen10392, gen10395)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen10396)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.default_semantics"))

							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.default_semantics"), gen10415)

		gen10425 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4766 := __args[0]
			_ = V4766
			gen10424 := Call(__e, ShenFunc(symsymbol_2), V4766)

			if True == gen10424 {
				gen10416 := Call(__e, ShenFunc(symexplode), V4766)

				gen10417 := Call(__e, ShenFunc(symshen_4strip_1pathname), gen10416)

				Cs := gen10417
				gen10421 := Call(__e, ShenFunc(symhd), Cs)

				gen10422 := Call(__e, ShenFunc(sym_a), gen10421, MakeString("<"))

				var gen10423 Obj
				if True == gen10422 {
					gen10418 := Call(__e, ShenFunc(symreverse), Cs)

					gen10419 := Call(__e, ShenFunc(symhd), gen10418)

					gen10420 := Call(__e, ShenFunc(sym_a), gen10419, MakeString(">"))

					if True == gen10420 {
						gen10423 = True
					} else {
						gen10423 = False
					}

				} else {
					gen10423 = False
				}
				if True == gen10423 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.grammar_symbol?"), gen10425)

		gen10445 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4768 := __args[0]
			_ = V4768
			gen10443 := Call(__e, ShenFunc(symcons_2), V4768)

			var gen10444 Obj
			if True == gen10443 {
				gen10441 := Call(__e, ShenFunc(symtl), V4768)

				gen10442 := Call(__e, ShenFunc(sym_a), Nil, gen10441)

				if True == gen10442 {
					gen10444 = True
				} else {
					gen10444 = False
				}

			} else {
				gen10444 = False
			}
			if True == gen10444 {
				__e.TailApply(ShenFunc(symhd), V4768)

				return
			} else {
				gen10440 := Call(__e, ShenFunc(symcons_2), V4768)

				if True == gen10440 {
					P := MakeSymbol("YaccParse")
					gen10426 := Call(__e, ShenFunc(symhd), V4768)

					gen10427 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

					gen10428 := Call(__e, ShenFunc(symcons), gen10427, Nil)

					gen10429 := Call(__e, ShenFunc(symcons), P, gen10428)

					gen10430 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen10429)

					gen10431 := Call(__e, ShenFunc(symtl), V4768)

					gen10432 := Call(__e, ShenFunc(symshen_4yacc__cases), gen10431)

					gen10433 := Call(__e, ShenFunc(symcons), P, Nil)

					gen10434 := Call(__e, ShenFunc(symcons), gen10432, gen10433)

					gen10435 := Call(__e, ShenFunc(symcons), gen10430, gen10434)

					gen10436 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen10435)

					gen10437 := Call(__e, ShenFunc(symcons), gen10436, Nil)

					gen10438 := Call(__e, ShenFunc(symcons), gen10426, gen10437)

					gen10439 := Call(__e, ShenFunc(symcons), P, gen10438)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen10439)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.yacc_cases"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc_cases"), gen10445)

		gen10457 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4770 := __args[0]
			_ = V4770
			gen10455 := Call(__e, ShenFunc(symcons_2), V4770)

			var gen10456 Obj
			if True == gen10455 {
				gen10452 := Call(__e, ShenFunc(symtl), V4770)

				gen10453 := Call(__e, ShenFunc(symcons_2), gen10452)

				var gen10454 Obj
				if True == gen10453 {
					gen10449 := Call(__e, ShenFunc(symtl), V4770)

					gen10450 := Call(__e, ShenFunc(symtl), gen10449)

					gen10451 := Call(__e, ShenFunc(sym_a), Nil, gen10450)

					if True == gen10451 {
						gen10454 = True
					} else {
						gen10454 = False
					}

				} else {
					gen10454 = False
				}
				if True == gen10454 {
					gen10456 = True
				} else {
					gen10456 = False
				}

			} else {
				gen10456 = False
			}
			if True == gen10456 {
				gen10446 := Call(__e, ShenFunc(symhd), V4770)

				gen10447 := Call(__e, ShenFunc(symtl), V4770)

				gen10448 := Call(__e, ShenFunc(symhd), gen10447)

				__e.TailApply(ShenFunc(symshen_4syntax), gen10446, MakeSymbol("Stream"), gen10448)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cc_body"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cc_body"), gen10457)

		gen10514 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4774 := __args[0]
			_ = V4774
			V4775 := __args[1]
			_ = V4775
			V4776 := __args[2]
			_ = V4776
			gen10512 := Call(__e, ShenFunc(sym_a), Nil, V4774)

			var gen10513 Obj
			if True == gen10512 {
				gen10510 := Call(__e, ShenFunc(symcons_2), V4776)

				var gen10511 Obj
				if True == gen10510 {
					gen10507 := Call(__e, ShenFunc(symhd), V4776)

					gen10508 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen10507)

					var gen10509 Obj
					if True == gen10508 {
						gen10504 := Call(__e, ShenFunc(symtl), V4776)

						gen10505 := Call(__e, ShenFunc(symcons_2), gen10504)

						var gen10506 Obj
						if True == gen10505 {
							gen10500 := Call(__e, ShenFunc(symtl), V4776)

							gen10501 := Call(__e, ShenFunc(symtl), gen10500)

							gen10502 := Call(__e, ShenFunc(symcons_2), gen10501)

							var gen10503 Obj
							if True == gen10502 {
								gen10496 := Call(__e, ShenFunc(symtl), V4776)

								gen10497 := Call(__e, ShenFunc(symtl), gen10496)

								gen10498 := Call(__e, ShenFunc(symtl), gen10497)

								gen10499 := Call(__e, ShenFunc(sym_a), Nil, gen10498)

								if True == gen10499 {
									gen10503 = True
								} else {
									gen10503 = False
								}

							} else {
								gen10503 = False
							}
							if True == gen10503 {
								gen10506 = True
							} else {
								gen10506 = False
							}

						} else {
							gen10506 = False
						}
						if True == gen10506 {
							gen10509 = True
						} else {
							gen10509 = False
						}

					} else {
						gen10509 = False
					}
					if True == gen10509 {
						gen10511 = True
					} else {
						gen10511 = False
					}

				} else {
					gen10511 = False
				}
				if True == gen10511 {
					gen10513 = True
				} else {
					gen10513 = False
				}

			} else {
				gen10513 = False
			}
			if True == gen10513 {
				gen10480 := Call(__e, ShenFunc(symtl), V4776)

				gen10481 := Call(__e, ShenFunc(symhd), gen10480)

				gen10482 := Call(__e, ShenFunc(symshen_4semantics), gen10481)

				gen10483 := Call(__e, ShenFunc(symcons), V4775, Nil)

				gen10484 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10483)

				gen10485 := Call(__e, ShenFunc(symtl), V4776)

				gen10486 := Call(__e, ShenFunc(symtl), gen10485)

				gen10487 := Call(__e, ShenFunc(symhd), gen10486)

				gen10488 := Call(__e, ShenFunc(symshen_4semantics), gen10487)

				gen10489 := Call(__e, ShenFunc(symcons), gen10488, Nil)

				gen10490 := Call(__e, ShenFunc(symcons), gen10484, gen10489)

				gen10491 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10490)

				gen10492 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				gen10493 := Call(__e, ShenFunc(symcons), gen10492, Nil)

				gen10494 := Call(__e, ShenFunc(symcons), gen10491, gen10493)

				gen10495 := Call(__e, ShenFunc(symcons), gen10482, gen10494)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10495)

				return

			} else {
				gen10479 := Call(__e, ShenFunc(sym_a), Nil, V4774)

				if True == gen10479 {
					gen10474 := Call(__e, ShenFunc(symcons), V4775, Nil)

					gen10475 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10474)

					gen10476 := Call(__e, ShenFunc(symshen_4semantics), V4776)

					gen10477 := Call(__e, ShenFunc(symcons), gen10476, Nil)

					gen10478 := Call(__e, ShenFunc(symcons), gen10475, gen10477)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pair"), gen10478)

					return

				} else {
					gen10473 := Call(__e, ShenFunc(symcons_2), V4774)

					if True == gen10473 {
						gen10471 := Call(__e, ShenFunc(symhd), V4774)

						gen10472 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen10471)

						if True == gen10472 {
							__e.TailApply(ShenFunc(symshen_4recursive__descent), V4774, V4775, V4776)

							return
						} else {
							gen10469 := Call(__e, ShenFunc(symhd), V4774)

							gen10470 := Call(__e, ShenFunc(symvariable_2), gen10469)

							if True == gen10470 {
								__e.TailApply(ShenFunc(symshen_4variable_1match), V4774, V4775, V4776)

								return
							} else {
								gen10467 := Call(__e, ShenFunc(symhd), V4774)

								gen10468 := Call(__e, ShenFunc(symshen_4jump__stream_2), gen10467)

								if True == gen10468 {
									__e.TailApply(ShenFunc(symshen_4jump__stream), V4774, V4775, V4776)

									return
								} else {
									gen10465 := Call(__e, ShenFunc(symhd), V4774)

									gen10466 := Call(__e, ShenFunc(symshen_4terminal_2), gen10465)

									if True == gen10466 {
										__e.TailApply(ShenFunc(symshen_4check__stream), V4774, V4775, V4776)

										return
									} else {
										gen10463 := Call(__e, ShenFunc(symhd), V4774)

										gen10464 := Call(__e, ShenFunc(symcons_2), gen10463)

										if True == gen10464 {
											gen10460 := Call(__e, ShenFunc(symhd), V4774)

											gen10461 := Call(__e, ShenFunc(symshen_4decons), gen10460)

											gen10462 := Call(__e, ShenFunc(symtl), V4774)

											__e.TailApply(ShenFunc(symshen_4list_1stream), gen10461, gen10462, V4775, V4776)

											return

										} else {
											gen10458 := Call(__e, ShenFunc(symhd), V4774)

											gen10459 := Call(__e, ShenFunc(symshen_4app), gen10458, MakeString(" is not legal syntax\n"), MakeSymbol("shen.a"))

											__e.TailApply(ShenFunc(symsimple_1error), gen10459)

											return

										}

									}

								}

							}

						}

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.syntax"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.syntax"), gen10514)

		gen10548 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4781 := __args[0]
			_ = V4781
			V4782 := __args[1]
			_ = V4782
			V4783 := __args[2]
			_ = V4783
			V4784 := __args[3]
			_ = V4784
			gen10515 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10516 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10515)

			gen10517 := Call(__e, ShenFunc(symcons), gen10516, Nil)

			gen10518 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10517)

			gen10519 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10520 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen10519)

			gen10521 := Call(__e, ShenFunc(symcons), gen10520, Nil)

			gen10522 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10521)

			gen10523 := Call(__e, ShenFunc(symcons), gen10522, Nil)

			gen10524 := Call(__e, ShenFunc(symcons), gen10518, gen10523)

			gen10525 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen10524)

			Test := gen10525
			gen10526 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.place"))

			Placeholder := gen10526
			gen10527 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10528 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen10527)

			gen10529 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10530 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10529)

			gen10531 := Call(__e, ShenFunc(symcons), gen10530, Nil)

			gen10532 := Call(__e, ShenFunc(symcons), gen10528, gen10531)

			gen10533 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10532)

			gen10534 := Call(__e, ShenFunc(symshen_4syntax), V4782, gen10533, V4784)

			RunOn := gen10534
			gen10535 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10536 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen10535)

			gen10537 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen10538 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10537)

			gen10539 := Call(__e, ShenFunc(symcons), gen10538, Nil)

			gen10540 := Call(__e, ShenFunc(symcons), gen10536, gen10539)

			gen10541 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10540)

			gen10542 := Call(__e, ShenFunc(symshen_4syntax), V4781, gen10541, Placeholder)

			gen10543 := Call(__e, ShenFunc(symshen_4insert_1runon), RunOn, Placeholder, gen10542)

			Action := gen10543
			gen10544 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

			gen10545 := Call(__e, ShenFunc(symcons), gen10544, Nil)

			gen10546 := Call(__e, ShenFunc(symcons), Action, gen10545)

			gen10547 := Call(__e, ShenFunc(symcons), Test, gen10546)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10547)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.list-stream"), gen10548)

		gen10594 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4786 := __args[0]
			_ = V4786
			gen10592 := Call(__e, ShenFunc(symcons_2), V4786)

			var gen10593 Obj
			if True == gen10592 {
				gen10589 := Call(__e, ShenFunc(symhd), V4786)

				gen10590 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen10589)

				var gen10591 Obj
				if True == gen10590 {
					gen10586 := Call(__e, ShenFunc(symtl), V4786)

					gen10587 := Call(__e, ShenFunc(symcons_2), gen10586)

					var gen10588 Obj
					if True == gen10587 {
						gen10582 := Call(__e, ShenFunc(symtl), V4786)

						gen10583 := Call(__e, ShenFunc(symtl), gen10582)

						gen10584 := Call(__e, ShenFunc(symcons_2), gen10583)

						var gen10585 Obj
						if True == gen10584 {
							gen10577 := Call(__e, ShenFunc(symtl), V4786)

							gen10578 := Call(__e, ShenFunc(symtl), gen10577)

							gen10579 := Call(__e, ShenFunc(symhd), gen10578)

							gen10580 := Call(__e, ShenFunc(sym_a), Nil, gen10579)

							var gen10581 Obj
							if True == gen10580 {
								gen10573 := Call(__e, ShenFunc(symtl), V4786)

								gen10574 := Call(__e, ShenFunc(symtl), gen10573)

								gen10575 := Call(__e, ShenFunc(symtl), gen10574)

								gen10576 := Call(__e, ShenFunc(sym_a), Nil, gen10575)

								if True == gen10576 {
									gen10581 = True
								} else {
									gen10581 = False
								}

							} else {
								gen10581 = False
							}
							if True == gen10581 {
								gen10585 = True
							} else {
								gen10585 = False
							}

						} else {
							gen10585 = False
						}
						if True == gen10585 {
							gen10588 = True
						} else {
							gen10588 = False
						}

					} else {
						gen10588 = False
					}
					if True == gen10588 {
						gen10591 = True
					} else {
						gen10591 = False
					}

				} else {
					gen10591 = False
				}
				if True == gen10591 {
					gen10593 = True
				} else {
					gen10593 = False
				}

			} else {
				gen10593 = False
			}
			if True == gen10593 {
				gen10571 := Call(__e, ShenFunc(symtl), V4786)

				gen10572 := Call(__e, ShenFunc(symhd), gen10571)

				__e.TailApply(ShenFunc(symcons), gen10572, Nil)

				return

			} else {
				gen10569 := Call(__e, ShenFunc(symcons_2), V4786)

				var gen10570 Obj
				if True == gen10569 {
					gen10566 := Call(__e, ShenFunc(symhd), V4786)

					gen10567 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen10566)

					var gen10568 Obj
					if True == gen10567 {
						gen10563 := Call(__e, ShenFunc(symtl), V4786)

						gen10564 := Call(__e, ShenFunc(symcons_2), gen10563)

						var gen10565 Obj
						if True == gen10564 {
							gen10559 := Call(__e, ShenFunc(symtl), V4786)

							gen10560 := Call(__e, ShenFunc(symtl), gen10559)

							gen10561 := Call(__e, ShenFunc(symcons_2), gen10560)

							var gen10562 Obj
							if True == gen10561 {
								gen10555 := Call(__e, ShenFunc(symtl), V4786)

								gen10556 := Call(__e, ShenFunc(symtl), gen10555)

								gen10557 := Call(__e, ShenFunc(symtl), gen10556)

								gen10558 := Call(__e, ShenFunc(sym_a), Nil, gen10557)

								if True == gen10558 {
									gen10562 = True
								} else {
									gen10562 = False
								}

							} else {
								gen10562 = False
							}
							if True == gen10562 {
								gen10565 = True
							} else {
								gen10565 = False
							}

						} else {
							gen10565 = False
						}
						if True == gen10565 {
							gen10568 = True
						} else {
							gen10568 = False
						}

					} else {
						gen10568 = False
					}
					if True == gen10568 {
						gen10570 = True
					} else {
						gen10570 = False
					}

				} else {
					gen10570 = False
				}
				if True == gen10570 {
					gen10549 := Call(__e, ShenFunc(symtl), V4786)

					gen10550 := Call(__e, ShenFunc(symhd), gen10549)

					gen10551 := Call(__e, ShenFunc(symtl), V4786)

					gen10552 := Call(__e, ShenFunc(symtl), gen10551)

					gen10553 := Call(__e, ShenFunc(symhd), gen10552)

					gen10554 := Call(__e, ShenFunc(symshen_4decons), gen10553)

					__e.TailApply(ShenFunc(symcons), gen10550, gen10554)

					return

				} else {
					__e.Return(V4786)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decons"), gen10594)

		gen10618 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4801 := __args[0]
			_ = V4801
			V4802 := __args[1]
			_ = V4802
			V4803 := __args[2]
			_ = V4803
			gen10616 := Call(__e, ShenFunc(symcons_2), V4803)

			var gen10617 Obj
			if True == gen10616 {
				gen10613 := Call(__e, ShenFunc(symhd), V4803)

				gen10614 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.pair"), gen10613)

				var gen10615 Obj
				if True == gen10614 {
					gen10610 := Call(__e, ShenFunc(symtl), V4803)

					gen10611 := Call(__e, ShenFunc(symcons_2), gen10610)

					var gen10612 Obj
					if True == gen10611 {
						gen10606 := Call(__e, ShenFunc(symtl), V4803)

						gen10607 := Call(__e, ShenFunc(symtl), gen10606)

						gen10608 := Call(__e, ShenFunc(symcons_2), gen10607)

						var gen10609 Obj
						if True == gen10608 {
							gen10601 := Call(__e, ShenFunc(symtl), V4803)

							gen10602 := Call(__e, ShenFunc(symtl), gen10601)

							gen10603 := Call(__e, ShenFunc(symtl), gen10602)

							gen10604 := Call(__e, ShenFunc(sym_a), Nil, gen10603)

							var gen10605 Obj
							if True == gen10604 {
								gen10597 := Call(__e, ShenFunc(symtl), V4803)

								gen10598 := Call(__e, ShenFunc(symtl), gen10597)

								gen10599 := Call(__e, ShenFunc(symhd), gen10598)

								gen10600 := Call(__e, ShenFunc(sym_a), gen10599, V4802)

								if True == gen10600 {
									gen10605 = True
								} else {
									gen10605 = False
								}

							} else {
								gen10605 = False
							}
							if True == gen10605 {
								gen10609 = True
							} else {
								gen10609 = False
							}

						} else {
							gen10609 = False
						}
						if True == gen10609 {
							gen10612 = True
						} else {
							gen10612 = False
						}

					} else {
						gen10612 = False
					}
					if True == gen10612 {
						gen10615 = True
					} else {
						gen10615 = False
					}

				} else {
					gen10615 = False
				}
				if True == gen10615 {
					gen10617 = True
				} else {
					gen10617 = False
				}

			} else {
				gen10617 = False
			}
			if True == gen10617 {
				__e.Return(V4801)
				return
			} else {
				gen10596 := Call(__e, ShenFunc(symcons_2), V4803)

				if True == gen10596 {
					gen10595 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Z := __args[0]
						_ = Z
						__e.TailApply(ShenFunc(symshen_4insert_1runon), V4801, V4802, Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen10595, V4803)

					return

				} else {
					__e.Return(V4803)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-runon"), gen10618)

		gen10623 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4809 := __args[0]
			_ = V4809
			gen10621 := Call(__e, ShenFunc(symelement_2), MakeString("."), V4809)

			gen10622 := Call(__e, ShenFunc(symnot), gen10621)

			if True == gen10622 {
				__e.Return(V4809)
				return
			} else {
				gen10620 := Call(__e, ShenFunc(symcons_2), V4809)

				if True == gen10620 {
					gen10619 := Call(__e, ShenFunc(symtl), V4809)

					__e.TailApply(ShenFunc(symshen_4strip_1pathname), gen10619)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.strip-pathname"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.strip-pathname"), gen10623)

		gen10650 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4813 := __args[0]
			_ = V4813
			V4814 := __args[1]
			_ = V4814
			V4815 := __args[2]
			_ = V4815
			gen10649 := Call(__e, ShenFunc(symcons_2), V4813)

			if True == gen10649 {
				gen10624 := Call(__e, ShenFunc(symhd), V4813)

				gen10625 := Call(__e, ShenFunc(symcons), V4814, Nil)

				gen10626 := Call(__e, ShenFunc(symcons), gen10624, gen10625)

				Test := gen10626
				gen10627 := Call(__e, ShenFunc(symtl), V4813)

				gen10628 := Call(__e, ShenFunc(symhd), V4813)

				gen10629 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen10628)

				gen10630 := Call(__e, ShenFunc(symshen_4syntax), gen10627, gen10629, V4815)

				Action := gen10630
				gen10631 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen10631
				gen10632 := Call(__e, ShenFunc(symhd), V4813)

				gen10633 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen10632)

				gen10634 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				gen10635 := Call(__e, ShenFunc(symhd), V4813)

				gen10636 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen10635)

				gen10637 := Call(__e, ShenFunc(symcons), gen10636, Nil)

				gen10638 := Call(__e, ShenFunc(symcons), gen10634, gen10637)

				gen10639 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen10638)

				gen10640 := Call(__e, ShenFunc(symcons), gen10639, Nil)

				gen10641 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen10640)

				gen10642 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen10643 := Call(__e, ShenFunc(symcons), Action, gen10642)

				gen10644 := Call(__e, ShenFunc(symcons), gen10641, gen10643)

				gen10645 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen10644)

				gen10646 := Call(__e, ShenFunc(symcons), gen10645, Nil)

				gen10647 := Call(__e, ShenFunc(symcons), Test, gen10646)

				gen10648 := Call(__e, ShenFunc(symcons), gen10633, gen10647)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen10648)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.recursive_descent"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.recursive_descent"), gen10650)

		gen10677 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4819 := __args[0]
			_ = V4819
			V4820 := __args[1]
			_ = V4820
			V4821 := __args[2]
			_ = V4821
			gen10676 := Call(__e, ShenFunc(symcons_2), V4819)

			if True == gen10676 {
				gen10651 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen10652 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10651)

				gen10653 := Call(__e, ShenFunc(symcons), gen10652, Nil)

				gen10654 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10653)

				Test := gen10654
				gen10655 := Call(__e, ShenFunc(symhd), V4819)

				gen10656 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen10655)

				gen10657 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen10658 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen10657)

				gen10659 := Call(__e, ShenFunc(symtl), V4819)

				gen10660 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen10661 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen10660)

				gen10662 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen10663 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10662)

				gen10664 := Call(__e, ShenFunc(symcons), gen10663, Nil)

				gen10665 := Call(__e, ShenFunc(symcons), gen10661, gen10664)

				gen10666 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10665)

				gen10667 := Call(__e, ShenFunc(symshen_4syntax), gen10659, gen10666, V4821)

				gen10668 := Call(__e, ShenFunc(symcons), gen10667, Nil)

				gen10669 := Call(__e, ShenFunc(symcons), gen10658, gen10668)

				gen10670 := Call(__e, ShenFunc(symcons), gen10656, gen10669)

				gen10671 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen10670)

				Action := gen10671
				gen10672 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen10672
				gen10673 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen10674 := Call(__e, ShenFunc(symcons), Action, gen10673)

				gen10675 := Call(__e, ShenFunc(symcons), Test, gen10674)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10675)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.variable-match"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.variable-match"), gen10677)

		gen10680 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4831 := __args[0]
			_ = V4831
			gen10679 := Call(__e, ShenFunc(symcons_2), V4831)

			if True == gen10679 {
				__e.Return(False)
				return
			} else {
				gen10678 := Call(__e, ShenFunc(symvariable_2), V4831)

				if True == gen10678 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terminal?"), gen10680)

		gen10682 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4837 := __args[0]
			_ = V4837
			gen10681 := Call(__e, ShenFunc(sym_a), V4837, MakeSymbol("_"))

			if True == gen10681 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.jump_stream?"), gen10682)

		gen10715 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4841 := __args[0]
			_ = V4841
			V4842 := __args[1]
			_ = V4842
			V4843 := __args[2]
			_ = V4843
			gen10714 := Call(__e, ShenFunc(symcons_2), V4841)

			if True == gen10714 {
				gen10683 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen10684 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10683)

				gen10685 := Call(__e, ShenFunc(symcons), gen10684, Nil)

				gen10686 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10685)

				gen10687 := Call(__e, ShenFunc(symhd), V4841)

				gen10688 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen10689 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen10688)

				gen10690 := Call(__e, ShenFunc(symcons), gen10689, Nil)

				gen10691 := Call(__e, ShenFunc(symcons), gen10687, gen10690)

				gen10692 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen10691)

				gen10693 := Call(__e, ShenFunc(symcons), gen10692, Nil)

				gen10694 := Call(__e, ShenFunc(symcons), gen10686, gen10693)

				gen10695 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen10694)

				Test := gen10695
				gen10696 := Call(__e, ShenFunc(symgensym), MakeSymbol("NewStream"))

				NewStr := gen10696
				gen10697 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen10698 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen10697)

				gen10699 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen10700 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10699)

				gen10701 := Call(__e, ShenFunc(symcons), gen10700, Nil)

				gen10702 := Call(__e, ShenFunc(symcons), gen10698, gen10701)

				gen10703 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10702)

				gen10704 := Call(__e, ShenFunc(symtl), V4841)

				gen10705 := Call(__e, ShenFunc(symshen_4syntax), gen10704, NewStr, V4843)

				gen10706 := Call(__e, ShenFunc(symcons), gen10705, Nil)

				gen10707 := Call(__e, ShenFunc(symcons), gen10703, gen10706)

				gen10708 := Call(__e, ShenFunc(symcons), NewStr, gen10707)

				gen10709 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen10708)

				Action := gen10709
				gen10710 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen10710
				gen10711 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen10712 := Call(__e, ShenFunc(symcons), Action, gen10711)

				gen10713 := Call(__e, ShenFunc(symcons), Test, gen10712)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10713)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.check_stream"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.check_stream"), gen10715)

		gen10734 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4847 := __args[0]
			_ = V4847
			V4848 := __args[1]
			_ = V4848
			V4849 := __args[2]
			_ = V4849
			gen10733 := Call(__e, ShenFunc(symcons_2), V4847)

			if True == gen10733 {
				gen10716 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen10717 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen10716)

				gen10718 := Call(__e, ShenFunc(symcons), gen10717, Nil)

				gen10719 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10718)

				Test := gen10719
				gen10720 := Call(__e, ShenFunc(symtl), V4847)

				gen10721 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen10722 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen10721)

				gen10723 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen10724 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10723)

				gen10725 := Call(__e, ShenFunc(symcons), gen10724, Nil)

				gen10726 := Call(__e, ShenFunc(symcons), gen10722, gen10725)

				gen10727 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen10726)

				gen10728 := Call(__e, ShenFunc(symshen_4syntax), gen10720, gen10727, V4849)

				Action := gen10728
				gen10729 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen10729
				gen10730 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen10731 := Call(__e, ShenFunc(symcons), Action, gen10730)

				gen10732 := Call(__e, ShenFunc(symcons), Test, gen10731)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10732)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.jump_stream"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.jump_stream"), gen10734)

		gen10742 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4851 := __args[0]
			_ = V4851
			gen10741 := Call(__e, ShenFunc(sym_a), Nil, V4851)

			if True == gen10741 {
				__e.Return(Nil)
				return
			} else {
				gen10740 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), V4851)

				if True == gen10740 {
					gen10738 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), V4851)

					gen10739 := Call(__e, ShenFunc(symcons), gen10738, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen10739)

					return

				} else {
					gen10737 := Call(__e, ShenFunc(symvariable_2), V4851)

					if True == gen10737 {
						__e.TailApply(ShenFunc(symconcat), MakeSymbol("Parse_"), V4851)

						return
					} else {
						gen10736 := Call(__e, ShenFunc(symcons_2), V4851)

						if True == gen10736 {
							gen10735 := MakeNative(func(__e Evaluator, __args ...Obj) {
								Z := __args[0]
								_ = Z
								__e.TailApply(ShenFunc(symshen_4semantics), Z)

								return
							}, 1)
							__e.TailApply(ShenFunc(symmap), gen10735, V4851)

							return

						} else {
							__e.Return(V4851)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.semantics"), gen10742)

		gen10744 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4854 := __args[0]
			_ = V4854
			V4855 := __args[1]
			_ = V4855
			gen10743 := Call(__e, ShenFunc(symcons), V4855, Nil)

			__e.TailApply(ShenFunc(symcons), V4854, gen10743)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pair"), gen10744)

		gen10746 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4857 := __args[0]
			_ = V4857
			gen10745 := Call(__e, ShenFunc(symtl), V4857)

			__e.TailApply(ShenFunc(symhd), gen10745)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.hdtl"), gen10746)

		gen10748 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4859 := __args[0]
			_ = V4859
			gen10747 := Call(__e, ShenFunc(symhd), V4859)

			__e.TailApply(ShenFunc(symhd), gen10747)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.hdhd"), gen10748)

		gen10750 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4861 := __args[0]
			_ = V4861
			gen10749 := Call(__e, ShenFunc(symhd), V4861)

			__e.TailApply(ShenFunc(symtl), gen10749)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tlhd"), gen10750)

		gen10760 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4869 := __args[0]
			_ = V4869
			gen10758 := Call(__e, ShenFunc(symcons_2), V4869)

			var gen10759 Obj
			if True == gen10758 {
				gen10755 := Call(__e, ShenFunc(symtl), V4869)

				gen10756 := Call(__e, ShenFunc(symcons_2), gen10755)

				var gen10757 Obj
				if True == gen10756 {
					gen10752 := Call(__e, ShenFunc(symtl), V4869)

					gen10753 := Call(__e, ShenFunc(symtl), gen10752)

					gen10754 := Call(__e, ShenFunc(sym_a), Nil, gen10753)

					if True == gen10754 {
						gen10757 = True
					} else {
						gen10757 = False
					}

				} else {
					gen10757 = False
				}
				if True == gen10757 {
					gen10759 = True
				} else {
					gen10759 = False
				}

			} else {
				gen10759 = False
			}
			if True == gen10759 {
				gen10751 := Call(__e, ShenFunc(symtl), V4869)

				__e.TailApply(ShenFunc(symhd), gen10751)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.snd-or-fail"), gen10760)

		gen10761 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__e.Return(MakeSymbol("shen.fail!"))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fail"), gen10761)

		gen10772 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4877 := __args[0]
			_ = V4877
			gen10770 := Call(__e, ShenFunc(symcons_2), V4877)

			var gen10771 Obj
			if True == gen10770 {
				gen10767 := Call(__e, ShenFunc(symtl), V4877)

				gen10768 := Call(__e, ShenFunc(symcons_2), gen10767)

				var gen10769 Obj
				if True == gen10768 {
					gen10764 := Call(__e, ShenFunc(symtl), V4877)

					gen10765 := Call(__e, ShenFunc(symtl), gen10764)

					gen10766 := Call(__e, ShenFunc(sym_a), Nil, gen10765)

					if True == gen10766 {
						gen10769 = True
					} else {
						gen10769 = False
					}

				} else {
					gen10769 = False
				}
				if True == gen10769 {
					gen10771 = True
				} else {
					gen10771 = False
				}

			} else {
				gen10771 = False
			}
			if True == gen10771 {
				gen10762 := Call(__e, ShenFunc(symhd), V4877)

				gen10763 := Call(__e, ShenFunc(symcons), gen10762, Nil)

				__e.TailApply(ShenFunc(symcons), Nil, gen10763)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("<!>"), gen10772)

		gen10783 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4883 := __args[0]
			_ = V4883
			gen10781 := Call(__e, ShenFunc(symcons_2), V4883)

			var gen10782 Obj
			if True == gen10781 {
				gen10778 := Call(__e, ShenFunc(symtl), V4883)

				gen10779 := Call(__e, ShenFunc(symcons_2), gen10778)

				var gen10780 Obj
				if True == gen10779 {
					gen10775 := Call(__e, ShenFunc(symtl), V4883)

					gen10776 := Call(__e, ShenFunc(symtl), gen10775)

					gen10777 := Call(__e, ShenFunc(sym_a), Nil, gen10776)

					if True == gen10777 {
						gen10780 = True
					} else {
						gen10780 = False
					}

				} else {
					gen10780 = False
				}
				if True == gen10780 {
					gen10782 = True
				} else {
					gen10782 = False
				}

			} else {
				gen10782 = False
			}
			if True == gen10782 {
				gen10773 := Call(__e, ShenFunc(symhd), V4883)

				gen10774 := Call(__e, ShenFunc(symcons), Nil, Nil)

				__e.TailApply(ShenFunc(symcons), gen10773, gen10774)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("<e>"))

				return
			}

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("<e>"), gen10783)

		return

	}, 0))
}
