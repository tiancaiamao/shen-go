package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen6294 := MakeNative(func(__e Evaluator) {
			V867 := __e.Get(1)
			_ = V867
			gen6279 := Call(__e, ShenFunc(symshen_4_5predicate_d_6), V867)

			Parse__shen_4_5predicate_d_6 := gen6279
			gen6291 := Call(__e, ShenFunc(symfail))

			gen6292 := Call(__e, ShenFunc(sym_a), gen6291, Parse__shen_4_5predicate_d_6)

			gen6293 := Call(__e, ShenFunc(symnot), gen6292)

			if True == gen6293 {
				gen6280 := Call(__e, ShenFunc(symshen_4_5clauses_d_6), Parse__shen_4_5predicate_d_6)

				Parse__shen_4_5clauses_d_6 := gen6280
				gen6288 := Call(__e, ShenFunc(symfail))

				gen6289 := Call(__e, ShenFunc(sym_a), gen6288, Parse__shen_4_5clauses_d_6)

				gen6290 := Call(__e, ShenFunc(symnot), gen6289)

				if True == gen6290 {
					gen6281 := Call(__e, ShenFunc(symhd), Parse__shen_4_5clauses_d_6)

					gen6283 := MakeNative(func(__e Evaluator) {
						Parse__X := __e.Get(1)
						_ = Parse__X
						gen6282 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predicate_d_6)

						__e.TailApply(ShenFunc(symshen_4insert_1predicate), gen6282, Parse__X)

						return

					}, 1)
					gen6284 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

					gen6285 := Call(__e, ShenFunc(symmap), gen6283, gen6284)

					gen6286 := Call(__e, ShenFunc(symshen_4prolog_1_6shen), gen6285)

					gen6287 := Call(__e, ShenFunc(symhd), gen6286)

					__e.TailApply(ShenFunc(symshen_4pair), gen6281, gen6287)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<defprolog>"), gen6294)

		gen6311 := MakeNative(func(__e Evaluator) {
			V876 := __e.Get(1)
			_ = V876
			V877 := __e.Get(2)
			_ = V877
			gen6309 := Call(__e, ShenFunc(symcons_2), V877)

			var gen6310 Obj
			if True == gen6309 {
				gen6306 := Call(__e, ShenFunc(symtl), V877)

				gen6307 := Call(__e, ShenFunc(symcons_2), gen6306)

				var gen6308 Obj
				if True == gen6307 {
					gen6303 := Call(__e, ShenFunc(symtl), V877)

					gen6304 := Call(__e, ShenFunc(symtl), gen6303)

					gen6305 := Call(__e, ShenFunc(sym_a), Nil, gen6304)

					if True == gen6305 {
						gen6308 = True
					} else {
						gen6308 = False
					}

				} else {
					gen6308 = False
				}
				if True == gen6308 {
					gen6310 = True
				} else {
					gen6310 = False
				}

			} else {
				gen6310 = False
			}
			if True == gen6310 {
				gen6297 := Call(__e, ShenFunc(symhd), V877)

				gen6298 := Call(__e, ShenFunc(symshen_4next_150), MakeNumber(50), gen6297)

				gen6299 := Call(__e, ShenFunc(symshen_4app), gen6298, MakeString("\n"), MakeSymbol("shen.a"))

				gen6300 := Call(__e, ShenFunc(symcn), MakeString(" here:\n\n "), gen6299)

				gen6301 := Call(__e, ShenFunc(symshen_4app), V876, gen6300, MakeSymbol("shen.a"))

				gen6302 := Call(__e, ShenFunc(symcn), MakeString("prolog syntax error in "), gen6301)

				__e.TailApply(ShenFunc(symsimple_1error), gen6302)

				return

			} else {
				gen6295 := Call(__e, ShenFunc(symshen_4app), V876, MakeString("\n"), MakeSymbol("shen.a"))

				gen6296 := Call(__e, ShenFunc(symcn), MakeString("prolog syntax error in "), gen6295)

				__e.TailApply(ShenFunc(symsimple_1error), gen6296)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-error"), gen6311)

		gen6320 := MakeNative(func(__e Evaluator) {
			V884 := __e.Get(1)
			_ = V884
			V885 := __e.Get(2)
			_ = V885
			gen6319 := Call(__e, ShenFunc(sym_a), Nil, V885)

			if True == gen6319 {
				__e.Return(MakeString(""))
				return
			} else {
				gen6318 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V884)

				if True == gen6318 {
					__e.Return(MakeString(""))
					return
				} else {
					gen6317 := Call(__e, ShenFunc(symcons_2), V885)

					if True == gen6317 {
						gen6312 := Call(__e, ShenFunc(symhd), V885)

						gen6313 := Call(__e, ShenFunc(symshen_4decons_1string), gen6312)

						gen6314 := Call(__e, ShenFunc(sym_1), V884, MakeNumber(1))

						gen6315 := Call(__e, ShenFunc(symtl), V885)

						gen6316 := Call(__e, ShenFunc(symshen_4next_150), gen6314, gen6315)

						__e.TailApply(ShenFunc(symcn), gen6313, gen6316)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.next-50"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.next-50"), gen6320)

		gen6338 := MakeNative(func(__e Evaluator) {
			V887 := __e.Get(1)
			_ = V887
			gen6336 := Call(__e, ShenFunc(symcons_2), V887)

			var gen6337 Obj
			if True == gen6336 {
				gen6333 := Call(__e, ShenFunc(symhd), V887)

				gen6334 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen6333)

				var gen6335 Obj
				if True == gen6334 {
					gen6330 := Call(__e, ShenFunc(symtl), V887)

					gen6331 := Call(__e, ShenFunc(symcons_2), gen6330)

					var gen6332 Obj
					if True == gen6331 {
						gen6326 := Call(__e, ShenFunc(symtl), V887)

						gen6327 := Call(__e, ShenFunc(symtl), gen6326)

						gen6328 := Call(__e, ShenFunc(symcons_2), gen6327)

						var gen6329 Obj
						if True == gen6328 {
							gen6322 := Call(__e, ShenFunc(symtl), V887)

							gen6323 := Call(__e, ShenFunc(symtl), gen6322)

							gen6324 := Call(__e, ShenFunc(symtl), gen6323)

							gen6325 := Call(__e, ShenFunc(sym_a), Nil, gen6324)

							if True == gen6325 {
								gen6329 = True
							} else {
								gen6329 = False
							}

						} else {
							gen6329 = False
						}
						if True == gen6329 {
							gen6332 = True
						} else {
							gen6332 = False
						}

					} else {
						gen6332 = False
					}
					if True == gen6332 {
						gen6335 = True
					} else {
						gen6335 = False
					}

				} else {
					gen6335 = False
				}
				if True == gen6335 {
					gen6337 = True
				} else {
					gen6337 = False
				}

			} else {
				gen6337 = False
			}
			if True == gen6337 {
				gen6321 := Call(__e, ShenFunc(symshen_4eval_1cons), V887)

				__e.TailApply(ShenFunc(symshen_4app), gen6321, MakeString(" "), MakeSymbol("shen.s"))

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4app), V887, MakeString(" "), MakeSymbol("shen.r"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decons-string"), gen6338)

		gen6351 := MakeNative(func(__e Evaluator) {
			V890 := __e.Get(1)
			_ = V890
			V891 := __e.Get(2)
			_ = V891
			gen6349 := Call(__e, ShenFunc(symcons_2), V891)

			var gen6350 Obj
			if True == gen6349 {
				gen6346 := Call(__e, ShenFunc(symtl), V891)

				gen6347 := Call(__e, ShenFunc(symcons_2), gen6346)

				var gen6348 Obj
				if True == gen6347 {
					gen6343 := Call(__e, ShenFunc(symtl), V891)

					gen6344 := Call(__e, ShenFunc(symtl), gen6343)

					gen6345 := Call(__e, ShenFunc(sym_a), Nil, gen6344)

					if True == gen6345 {
						gen6348 = True
					} else {
						gen6348 = False
					}

				} else {
					gen6348 = False
				}
				if True == gen6348 {
					gen6350 = True
				} else {
					gen6350 = False
				}

			} else {
				gen6350 = False
			}
			if True == gen6350 {
				gen6339 := Call(__e, ShenFunc(symhd), V891)

				gen6340 := Call(__e, ShenFunc(symcons), V890, gen6339)

				gen6341 := Call(__e, ShenFunc(symtl), V891)

				gen6342 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen6341)

				__e.TailApply(ShenFunc(symcons), gen6340, gen6342)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-predicate"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-predicate"), gen6351)

		gen6359 := MakeNative(func(__e Evaluator) {
			V893 := __e.Get(1)
			_ = V893
			gen6357 := Call(__e, ShenFunc(symhd), V893)

			gen6358 := Call(__e, ShenFunc(symcons_2), gen6357)

			if True == gen6358 {
				gen6352 := Call(__e, ShenFunc(symshen_4hdhd), V893)

				Parse__X := gen6352
				gen6353 := Call(__e, ShenFunc(symshen_4tlhd), V893)

				gen6354 := Call(__e, ShenFunc(symshen_4hdtl), V893)

				gen6355 := Call(__e, ShenFunc(symshen_4pair), gen6353, gen6354)

				gen6356 := Call(__e, ShenFunc(symhd), gen6355)

				__e.TailApply(ShenFunc(symshen_4pair), gen6356, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<predicate*>"), gen6359)

		gen6380 := MakeNative(func(__e Evaluator) {
			V895 := __e.Get(1)
			_ = V895
			gen6360 := Call(__e, ShenFunc(symshen_4_5clause_d_6), V895)

			Parse__shen_4_5clause_d_6 := gen6360
			gen6369 := Call(__e, ShenFunc(symfail))

			gen6370 := Call(__e, ShenFunc(sym_a), gen6369, Parse__shen_4_5clause_d_6)

			gen6371 := Call(__e, ShenFunc(symnot), gen6370)

			var gen6372 Obj
			if True == gen6371 {
				gen6361 := Call(__e, ShenFunc(symshen_4_5clauses_d_6), Parse__shen_4_5clause_d_6)

				Parse__shen_4_5clauses_d_6 := gen6361
				gen6366 := Call(__e, ShenFunc(symfail))

				gen6367 := Call(__e, ShenFunc(sym_a), gen6366, Parse__shen_4_5clauses_d_6)

				gen6368 := Call(__e, ShenFunc(symnot), gen6367)

				if True == gen6368 {
					gen6362 := Call(__e, ShenFunc(symhd), Parse__shen_4_5clauses_d_6)

					gen6363 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clause_d_6)

					gen6364 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

					gen6365 := Call(__e, ShenFunc(symcons), gen6363, gen6364)

					gen6372 = Call(__e, ShenFunc(symshen_4pair), gen6362, gen6365)

				} else {
					gen6372 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6372 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6372
			gen6378 := Call(__e, ShenFunc(symfail))

			gen6379 := Call(__e, ShenFunc(sym_a), YaccParse, gen6378)

			if True == gen6379 {
				gen6373 := Call(__e, ShenFunc(sym_5e_6), V895)

				Parse___5e_6 := gen6373
				gen6375 := Call(__e, ShenFunc(symfail))

				gen6376 := Call(__e, ShenFunc(sym_a), gen6375, Parse___5e_6)

				gen6377 := Call(__e, ShenFunc(symnot), gen6376)

				if True == gen6377 {
					gen6374 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen6374, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<clauses*>"), gen6380)

		gen6406 := MakeNative(func(__e Evaluator) {
			V898 := __e.Get(1)
			_ = V898
			gen6381 := Call(__e, ShenFunc(symshen_4_5head_d_6), V898)

			Parse__shen_4_5head_d_6 := gen6381
			gen6403 := Call(__e, ShenFunc(symfail))

			gen6404 := Call(__e, ShenFunc(sym_a), gen6403, Parse__shen_4_5head_d_6)

			gen6405 := Call(__e, ShenFunc(symnot), gen6404)

			if True == gen6405 {
				gen6400 := Call(__e, ShenFunc(symhd), Parse__shen_4_5head_d_6)

				gen6401 := Call(__e, ShenFunc(symcons_2), gen6400)

				var gen6402 Obj
				if True == gen6401 {
					gen6398 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5head_d_6)

					gen6399 := Call(__e, ShenFunc(sym_a), MakeSymbol("<--"), gen6398)

					if True == gen6399 {
						gen6402 = True
					} else {
						gen6402 = False
					}

				} else {
					gen6402 = False
				}
				if True == gen6402 {
					gen6382 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5head_d_6)

					gen6383 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

					gen6384 := Call(__e, ShenFunc(symshen_4pair), gen6382, gen6383)

					NewStream896 := gen6384
					gen6385 := Call(__e, ShenFunc(symshen_4_5body_d_6), NewStream896)

					Parse__shen_4_5body_d_6 := gen6385
					gen6395 := Call(__e, ShenFunc(symfail))

					gen6396 := Call(__e, ShenFunc(sym_a), gen6395, Parse__shen_4_5body_d_6)

					gen6397 := Call(__e, ShenFunc(symnot), gen6396)

					if True == gen6397 {
						gen6386 := Call(__e, ShenFunc(symshen_4_5end_d_6), Parse__shen_4_5body_d_6)

						Parse__shen_4_5end_d_6 := gen6386
						gen6392 := Call(__e, ShenFunc(symfail))

						gen6393 := Call(__e, ShenFunc(sym_a), gen6392, Parse__shen_4_5end_d_6)

						gen6394 := Call(__e, ShenFunc(symnot), gen6393)

						if True == gen6394 {
							gen6387 := Call(__e, ShenFunc(symhd), Parse__shen_4_5end_d_6)

							gen6388 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

							gen6389 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5body_d_6)

							gen6390 := Call(__e, ShenFunc(symcons), gen6389, Nil)

							gen6391 := Call(__e, ShenFunc(symcons), gen6388, gen6390)

							__e.TailApply(ShenFunc(symshen_4pair), gen6387, gen6391)

							return

						} else {
							__e.TailApply(ShenFunc(symfail))

							return
						}

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<clause*>"), gen6406)

		gen6427 := MakeNative(func(__e Evaluator) {
			V900 := __e.Get(1)
			_ = V900
			gen6407 := Call(__e, ShenFunc(symshen_4_5term_d_6), V900)

			Parse__shen_4_5term_d_6 := gen6407
			gen6416 := Call(__e, ShenFunc(symfail))

			gen6417 := Call(__e, ShenFunc(sym_a), gen6416, Parse__shen_4_5term_d_6)

			gen6418 := Call(__e, ShenFunc(symnot), gen6417)

			var gen6419 Obj
			if True == gen6418 {
				gen6408 := Call(__e, ShenFunc(symshen_4_5head_d_6), Parse__shen_4_5term_d_6)

				Parse__shen_4_5head_d_6 := gen6408
				gen6413 := Call(__e, ShenFunc(symfail))

				gen6414 := Call(__e, ShenFunc(sym_a), gen6413, Parse__shen_4_5head_d_6)

				gen6415 := Call(__e, ShenFunc(symnot), gen6414)

				if True == gen6415 {
					gen6409 := Call(__e, ShenFunc(symhd), Parse__shen_4_5head_d_6)

					gen6410 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5term_d_6)

					gen6411 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

					gen6412 := Call(__e, ShenFunc(symcons), gen6410, gen6411)

					gen6419 = Call(__e, ShenFunc(symshen_4pair), gen6409, gen6412)

				} else {
					gen6419 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6419 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6419
			gen6425 := Call(__e, ShenFunc(symfail))

			gen6426 := Call(__e, ShenFunc(sym_a), YaccParse, gen6425)

			if True == gen6426 {
				gen6420 := Call(__e, ShenFunc(sym_5e_6), V900)

				Parse___5e_6 := gen6420
				gen6422 := Call(__e, ShenFunc(symfail))

				gen6423 := Call(__e, ShenFunc(sym_a), gen6422, Parse___5e_6)

				gen6424 := Call(__e, ShenFunc(symnot), gen6423)

				if True == gen6424 {
					gen6421 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen6421, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<head*>"), gen6427)

		gen6440 := MakeNative(func(__e Evaluator) {
			V902 := __e.Get(1)
			_ = V902
			gen6438 := Call(__e, ShenFunc(symhd), V902)

			gen6439 := Call(__e, ShenFunc(symcons_2), gen6438)

			if True == gen6439 {
				gen6428 := Call(__e, ShenFunc(symshen_4hdhd), V902)

				Parse__X := gen6428
				gen6435 := Call(__e, ShenFunc(sym_a), MakeSymbol("<--"), Parse__X)

				gen6436 := Call(__e, ShenFunc(symnot), gen6435)

				var gen6437 Obj
				if True == gen6436 {
					gen6434 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), Parse__X)

					if True == gen6434 {
						gen6437 = True
					} else {
						gen6437 = False
					}

				} else {
					gen6437 = False
				}
				if True == gen6437 {
					gen6429 := Call(__e, ShenFunc(symshen_4tlhd), V902)

					gen6430 := Call(__e, ShenFunc(symshen_4hdtl), V902)

					gen6431 := Call(__e, ShenFunc(symshen_4pair), gen6429, gen6430)

					gen6432 := Call(__e, ShenFunc(symhd), gen6431)

					gen6433 := Call(__e, ShenFunc(symshen_4eval_1cons), Parse__X)

					__e.TailApply(ShenFunc(symshen_4pair), gen6432, gen6433)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<term*>"), gen6440)

		gen6511 := MakeNative(func(__e Evaluator) {
			V908 := __e.Get(1)
			_ = V908
			gen6509 := Call(__e, ShenFunc(symcons_2), V908)

			var gen6510 Obj
			if True == gen6509 {
				gen6506 := Call(__e, ShenFunc(symhd), V908)

				gen6507 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen6506)

				var gen6508 Obj
				if True == gen6507 {
					gen6503 := Call(__e, ShenFunc(symtl), V908)

					gen6504 := Call(__e, ShenFunc(symcons_2), gen6503)

					var gen6505 Obj
					if True == gen6504 {
						gen6499 := Call(__e, ShenFunc(symtl), V908)

						gen6500 := Call(__e, ShenFunc(symtl), gen6499)

						gen6501 := Call(__e, ShenFunc(symcons_2), gen6500)

						var gen6502 Obj
						if True == gen6501 {
							gen6495 := Call(__e, ShenFunc(symtl), V908)

							gen6496 := Call(__e, ShenFunc(symtl), gen6495)

							gen6497 := Call(__e, ShenFunc(symtl), gen6496)

							gen6498 := Call(__e, ShenFunc(sym_a), Nil, gen6497)

							if True == gen6498 {
								gen6502 = True
							} else {
								gen6502 = False
							}

						} else {
							gen6502 = False
						}
						if True == gen6502 {
							gen6505 = True
						} else {
							gen6505 = False
						}

					} else {
						gen6505 = False
					}
					if True == gen6505 {
						gen6508 = True
					} else {
						gen6508 = False
					}

				} else {
					gen6508 = False
				}
				if True == gen6508 {
					gen6510 = True
				} else {
					gen6510 = False
				}

			} else {
				gen6510 = False
			}
			if True == gen6510 {
				gen6492 := Call(__e, ShenFunc(symtl), V908)

				gen6493 := Call(__e, ShenFunc(symhd), gen6492)

				gen6494 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), gen6493)

				if True == gen6494 {
					gen6488 := Call(__e, ShenFunc(symtl), V908)

					gen6489 := Call(__e, ShenFunc(symtl), gen6488)

					gen6490 := Call(__e, ShenFunc(symhd), gen6489)

					gen6491 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), gen6490)

					if True == gen6491 {
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
				gen6486 := Call(__e, ShenFunc(symcons_2), V908)

				var gen6487 Obj
				if True == gen6486 {
					gen6483 := Call(__e, ShenFunc(symhd), V908)

					gen6484 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6483)

					var gen6485 Obj
					if True == gen6484 {
						gen6480 := Call(__e, ShenFunc(symtl), V908)

						gen6481 := Call(__e, ShenFunc(symcons_2), gen6480)

						var gen6482 Obj
						if True == gen6481 {
							gen6476 := Call(__e, ShenFunc(symtl), V908)

							gen6477 := Call(__e, ShenFunc(symtl), gen6476)

							gen6478 := Call(__e, ShenFunc(symcons_2), gen6477)

							var gen6479 Obj
							if True == gen6478 {
								gen6471 := Call(__e, ShenFunc(symtl), V908)

								gen6472 := Call(__e, ShenFunc(symtl), gen6471)

								gen6473 := Call(__e, ShenFunc(symhd), gen6472)

								gen6474 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen6473)

								var gen6475 Obj
								if True == gen6474 {
									gen6467 := Call(__e, ShenFunc(symtl), V908)

									gen6468 := Call(__e, ShenFunc(symtl), gen6467)

									gen6469 := Call(__e, ShenFunc(symtl), gen6468)

									gen6470 := Call(__e, ShenFunc(sym_a), Nil, gen6469)

									if True == gen6470 {
										gen6475 = True
									} else {
										gen6475 = False
									}

								} else {
									gen6475 = False
								}
								if True == gen6475 {
									gen6479 = True
								} else {
									gen6479 = False
								}

							} else {
								gen6479 = False
							}
							if True == gen6479 {
								gen6482 = True
							} else {
								gen6482 = False
							}

						} else {
							gen6482 = False
						}
						if True == gen6482 {
							gen6485 = True
						} else {
							gen6485 = False
						}

					} else {
						gen6485 = False
					}
					if True == gen6485 {
						gen6487 = True
					} else {
						gen6487 = False
					}

				} else {
					gen6487 = False
				}
				if True == gen6487 {
					gen6465 := Call(__e, ShenFunc(symtl), V908)

					gen6466 := Call(__e, ShenFunc(symhd), gen6465)

					__e.TailApply(ShenFunc(symshen_4legitimate_1term_2), gen6466)

					return

				} else {
					gen6463 := Call(__e, ShenFunc(symcons_2), V908)

					var gen6464 Obj
					if True == gen6463 {
						gen6460 := Call(__e, ShenFunc(symhd), V908)

						gen6461 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6460)

						var gen6462 Obj
						if True == gen6461 {
							gen6457 := Call(__e, ShenFunc(symtl), V908)

							gen6458 := Call(__e, ShenFunc(symcons_2), gen6457)

							var gen6459 Obj
							if True == gen6458 {
								gen6453 := Call(__e, ShenFunc(symtl), V908)

								gen6454 := Call(__e, ShenFunc(symtl), gen6453)

								gen6455 := Call(__e, ShenFunc(symcons_2), gen6454)

								var gen6456 Obj
								if True == gen6455 {
									gen6448 := Call(__e, ShenFunc(symtl), V908)

									gen6449 := Call(__e, ShenFunc(symtl), gen6448)

									gen6450 := Call(__e, ShenFunc(symhd), gen6449)

									gen6451 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen6450)

									var gen6452 Obj
									if True == gen6451 {
										gen6444 := Call(__e, ShenFunc(symtl), V908)

										gen6445 := Call(__e, ShenFunc(symtl), gen6444)

										gen6446 := Call(__e, ShenFunc(symtl), gen6445)

										gen6447 := Call(__e, ShenFunc(sym_a), Nil, gen6446)

										if True == gen6447 {
											gen6452 = True
										} else {
											gen6452 = False
										}

									} else {
										gen6452 = False
									}
									if True == gen6452 {
										gen6456 = True
									} else {
										gen6456 = False
									}

								} else {
									gen6456 = False
								}
								if True == gen6456 {
									gen6459 = True
								} else {
									gen6459 = False
								}

							} else {
								gen6459 = False
							}
							if True == gen6459 {
								gen6462 = True
							} else {
								gen6462 = False
							}

						} else {
							gen6462 = False
						}
						if True == gen6462 {
							gen6464 = True
						} else {
							gen6464 = False
						}

					} else {
						gen6464 = False
					}
					if True == gen6464 {
						gen6442 := Call(__e, ShenFunc(symtl), V908)

						gen6443 := Call(__e, ShenFunc(symhd), gen6442)

						__e.TailApply(ShenFunc(symshen_4legitimate_1term_2), gen6443)

						return

					} else {
						gen6441 := Call(__e, ShenFunc(symcons_2), V908)

						if True == gen6441 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.legitimate-term?"), gen6511)

		gen6557 := MakeNative(func(__e Evaluator) {
			V910 := __e.Get(1)
			_ = V910
			gen6555 := Call(__e, ShenFunc(symcons_2), V910)

			var gen6556 Obj
			if True == gen6555 {
				gen6552 := Call(__e, ShenFunc(symhd), V910)

				gen6553 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen6552)

				var gen6554 Obj
				if True == gen6553 {
					gen6549 := Call(__e, ShenFunc(symtl), V910)

					gen6550 := Call(__e, ShenFunc(symcons_2), gen6549)

					var gen6551 Obj
					if True == gen6550 {
						gen6545 := Call(__e, ShenFunc(symtl), V910)

						gen6546 := Call(__e, ShenFunc(symtl), gen6545)

						gen6547 := Call(__e, ShenFunc(symcons_2), gen6546)

						var gen6548 Obj
						if True == gen6547 {
							gen6541 := Call(__e, ShenFunc(symtl), V910)

							gen6542 := Call(__e, ShenFunc(symtl), gen6541)

							gen6543 := Call(__e, ShenFunc(symtl), gen6542)

							gen6544 := Call(__e, ShenFunc(sym_a), Nil, gen6543)

							if True == gen6544 {
								gen6548 = True
							} else {
								gen6548 = False
							}

						} else {
							gen6548 = False
						}
						if True == gen6548 {
							gen6551 = True
						} else {
							gen6551 = False
						}

					} else {
						gen6551 = False
					}
					if True == gen6551 {
						gen6554 = True
					} else {
						gen6554 = False
					}

				} else {
					gen6554 = False
				}
				if True == gen6554 {
					gen6556 = True
				} else {
					gen6556 = False
				}

			} else {
				gen6556 = False
			}
			if True == gen6556 {
				gen6534 := Call(__e, ShenFunc(symtl), V910)

				gen6535 := Call(__e, ShenFunc(symhd), gen6534)

				gen6536 := Call(__e, ShenFunc(symshen_4eval_1cons), gen6535)

				gen6537 := Call(__e, ShenFunc(symtl), V910)

				gen6538 := Call(__e, ShenFunc(symtl), gen6537)

				gen6539 := Call(__e, ShenFunc(symhd), gen6538)

				gen6540 := Call(__e, ShenFunc(symshen_4eval_1cons), gen6539)

				__e.TailApply(ShenFunc(symcons), gen6536, gen6540)

				return

			} else {
				gen6532 := Call(__e, ShenFunc(symcons_2), V910)

				var gen6533 Obj
				if True == gen6532 {
					gen6529 := Call(__e, ShenFunc(symhd), V910)

					gen6530 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6529)

					var gen6531 Obj
					if True == gen6530 {
						gen6526 := Call(__e, ShenFunc(symtl), V910)

						gen6527 := Call(__e, ShenFunc(symcons_2), gen6526)

						var gen6528 Obj
						if True == gen6527 {
							gen6522 := Call(__e, ShenFunc(symtl), V910)

							gen6523 := Call(__e, ShenFunc(symtl), gen6522)

							gen6524 := Call(__e, ShenFunc(symcons_2), gen6523)

							var gen6525 Obj
							if True == gen6524 {
								gen6518 := Call(__e, ShenFunc(symtl), V910)

								gen6519 := Call(__e, ShenFunc(symtl), gen6518)

								gen6520 := Call(__e, ShenFunc(symtl), gen6519)

								gen6521 := Call(__e, ShenFunc(sym_a), Nil, gen6520)

								if True == gen6521 {
									gen6525 = True
								} else {
									gen6525 = False
								}

							} else {
								gen6525 = False
							}
							if True == gen6525 {
								gen6528 = True
							} else {
								gen6528 = False
							}

						} else {
							gen6528 = False
						}
						if True == gen6528 {
							gen6531 = True
						} else {
							gen6531 = False
						}

					} else {
						gen6531 = False
					}
					if True == gen6531 {
						gen6533 = True
					} else {
						gen6533 = False
					}

				} else {
					gen6533 = False
				}
				if True == gen6533 {
					gen6512 := Call(__e, ShenFunc(symtl), V910)

					gen6513 := Call(__e, ShenFunc(symhd), gen6512)

					gen6514 := Call(__e, ShenFunc(symshen_4eval_1cons), gen6513)

					gen6515 := Call(__e, ShenFunc(symtl), V910)

					gen6516 := Call(__e, ShenFunc(symtl), gen6515)

					gen6517 := Call(__e, ShenFunc(symcons), gen6514, gen6516)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("mode"), gen6517)

					return

				} else {
					__e.Return(V910)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.eval-cons"), gen6557)

		gen6578 := MakeNative(func(__e Evaluator) {
			V912 := __e.Get(1)
			_ = V912
			gen6558 := Call(__e, ShenFunc(symshen_4_5literal_d_6), V912)

			Parse__shen_4_5literal_d_6 := gen6558
			gen6567 := Call(__e, ShenFunc(symfail))

			gen6568 := Call(__e, ShenFunc(sym_a), gen6567, Parse__shen_4_5literal_d_6)

			gen6569 := Call(__e, ShenFunc(symnot), gen6568)

			var gen6570 Obj
			if True == gen6569 {
				gen6559 := Call(__e, ShenFunc(symshen_4_5body_d_6), Parse__shen_4_5literal_d_6)

				Parse__shen_4_5body_d_6 := gen6559
				gen6564 := Call(__e, ShenFunc(symfail))

				gen6565 := Call(__e, ShenFunc(sym_a), gen6564, Parse__shen_4_5body_d_6)

				gen6566 := Call(__e, ShenFunc(symnot), gen6565)

				if True == gen6566 {
					gen6560 := Call(__e, ShenFunc(symhd), Parse__shen_4_5body_d_6)

					gen6561 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5literal_d_6)

					gen6562 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5body_d_6)

					gen6563 := Call(__e, ShenFunc(symcons), gen6561, gen6562)

					gen6570 = Call(__e, ShenFunc(symshen_4pair), gen6560, gen6563)

				} else {
					gen6570 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen6570 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6570
			gen6576 := Call(__e, ShenFunc(symfail))

			gen6577 := Call(__e, ShenFunc(sym_a), YaccParse, gen6576)

			if True == gen6577 {
				gen6571 := Call(__e, ShenFunc(sym_5e_6), V912)

				Parse___5e_6 := gen6571
				gen6573 := Call(__e, ShenFunc(symfail))

				gen6574 := Call(__e, ShenFunc(sym_a), gen6573, Parse___5e_6)

				gen6575 := Call(__e, ShenFunc(symnot), gen6574)

				if True == gen6575 {
					gen6572 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen6572, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<body*>"), gen6578)

		gen6602 := MakeNative(func(__e Evaluator) {
			V915 := __e.Get(1)
			_ = V915
			gen6588 := Call(__e, ShenFunc(symhd), V915)

			gen6589 := Call(__e, ShenFunc(symcons_2), gen6588)

			var gen6590 Obj
			if True == gen6589 {
				gen6586 := Call(__e, ShenFunc(symshen_4hdhd), V915)

				gen6587 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), gen6586)

				if True == gen6587 {
					gen6590 = True
				} else {
					gen6590 = False
				}

			} else {
				gen6590 = False
			}
			var gen6591 Obj
			if True == gen6590 {
				gen6579 := Call(__e, ShenFunc(symshen_4tlhd), V915)

				gen6580 := Call(__e, ShenFunc(symshen_4hdtl), V915)

				gen6581 := Call(__e, ShenFunc(symshen_4pair), gen6579, gen6580)

				NewStream913 := gen6581
				gen6582 := Call(__e, ShenFunc(symhd), NewStream913)

				gen6583 := Call(__e, ShenFunc(symintern), MakeString("Throwcontrol"))

				gen6584 := Call(__e, ShenFunc(symcons), gen6583, Nil)

				gen6585 := Call(__e, ShenFunc(symcons), MakeSymbol("cut"), gen6584)

				gen6591 = Call(__e, ShenFunc(symshen_4pair), gen6582, gen6585)

			} else {
				gen6591 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen6591
			gen6600 := Call(__e, ShenFunc(symfail))

			gen6601 := Call(__e, ShenFunc(sym_a), YaccParse, gen6600)

			if True == gen6601 {
				gen6598 := Call(__e, ShenFunc(symhd), V915)

				gen6599 := Call(__e, ShenFunc(symcons_2), gen6598)

				if True == gen6599 {
					gen6592 := Call(__e, ShenFunc(symshen_4hdhd), V915)

					Parse__X := gen6592
					gen6597 := Call(__e, ShenFunc(symcons_2), Parse__X)

					if True == gen6597 {
						gen6593 := Call(__e, ShenFunc(symshen_4tlhd), V915)

						gen6594 := Call(__e, ShenFunc(symshen_4hdtl), V915)

						gen6595 := Call(__e, ShenFunc(symshen_4pair), gen6593, gen6594)

						gen6596 := Call(__e, ShenFunc(symhd), gen6595)

						__e.TailApply(ShenFunc(symshen_4pair), gen6596, Parse__X)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<literal*>"), gen6602)

		gen6611 := MakeNative(func(__e Evaluator) {
			V917 := __e.Get(1)
			_ = V917
			gen6609 := Call(__e, ShenFunc(symhd), V917)

			gen6610 := Call(__e, ShenFunc(symcons_2), gen6609)

			if True == gen6610 {
				gen6603 := Call(__e, ShenFunc(symshen_4hdhd), V917)

				Parse__X := gen6603
				gen6608 := Call(__e, ShenFunc(sym_a), Parse__X, MakeSymbol(";"))

				if True == gen6608 {
					gen6604 := Call(__e, ShenFunc(symshen_4tlhd), V917)

					gen6605 := Call(__e, ShenFunc(symshen_4hdtl), V917)

					gen6606 := Call(__e, ShenFunc(symshen_4pair), gen6604, gen6605)

					gen6607 := Call(__e, ShenFunc(symhd), gen6606)

					__e.TailApply(ShenFunc(symshen_4pair), gen6607, Parse__X)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<end*>"), gen6611)

		gen6614 := MakeNative(func(__e Evaluator) {
			V921 := __e.Get(1)
			_ = V921
			V922 := __e.Get(2)
			_ = V922
			V923 := __e.Get(3)
			_ = V923
			gen6612 := Call(__e, ShenFunc(symthaw), V923)

			Result := gen6612
			gen6613 := Call(__e, ShenFunc(sym_a), Result, False)

			if True == gen6613 {
				__e.Return(V921)
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("cut"), gen6614)

		gen6642 := MakeNative(func(__e Evaluator) {
			V925 := __e.Get(1)
			_ = V925
			gen6640 := Call(__e, ShenFunc(symcons_2), V925)

			var gen6641 Obj
			if True == gen6640 {
				gen6637 := Call(__e, ShenFunc(symhd), V925)

				gen6638 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6637)

				var gen6639 Obj
				if True == gen6638 {
					gen6634 := Call(__e, ShenFunc(symtl), V925)

					gen6635 := Call(__e, ShenFunc(symcons_2), gen6634)

					var gen6636 Obj
					if True == gen6635 {
						gen6630 := Call(__e, ShenFunc(symtl), V925)

						gen6631 := Call(__e, ShenFunc(symtl), gen6630)

						gen6632 := Call(__e, ShenFunc(symcons_2), gen6631)

						var gen6633 Obj
						if True == gen6632 {
							gen6626 := Call(__e, ShenFunc(symtl), V925)

							gen6627 := Call(__e, ShenFunc(symtl), gen6626)

							gen6628 := Call(__e, ShenFunc(symtl), gen6627)

							gen6629 := Call(__e, ShenFunc(sym_a), Nil, gen6628)

							if True == gen6629 {
								gen6633 = True
							} else {
								gen6633 = False
							}

						} else {
							gen6633 = False
						}
						if True == gen6633 {
							gen6636 = True
						} else {
							gen6636 = False
						}

					} else {
						gen6636 = False
					}
					if True == gen6636 {
						gen6639 = True
					} else {
						gen6639 = False
					}

				} else {
					gen6639 = False
				}
				if True == gen6639 {
					gen6641 = True
				} else {
					gen6641 = False
				}

			} else {
				gen6641 = False
			}
			if True == gen6641 {
				__e.Return(V925)
				return
			} else {
				gen6625 := Call(__e, ShenFunc(sym_a), Nil, V925)

				if True == gen6625 {
					__e.Return(Nil)
					return
				} else {
					gen6624 := Call(__e, ShenFunc(symcons_2), V925)

					if True == gen6624 {
						gen6615 := Call(__e, ShenFunc(symhd), V925)

						gen6616 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

						gen6617 := Call(__e, ShenFunc(symcons), gen6615, gen6616)

						gen6618 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6617)

						gen6619 := Call(__e, ShenFunc(symtl), V925)

						gen6620 := Call(__e, ShenFunc(symshen_4insert__modes), gen6619)

						gen6621 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), Nil)

						gen6622 := Call(__e, ShenFunc(symcons), gen6620, gen6621)

						gen6623 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6622)

						__e.TailApply(ShenFunc(symcons), gen6618, gen6623)

						return

					} else {
						__e.Return(V925)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert_modes"), gen6642)

		gen6645 := MakeNative(func(__e Evaluator) {
			V927 := __e.Get(1)
			_ = V927
			gen6643 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symeval), X)

				return
			}, 1)
			gen6644 := Call(__e, ShenFunc(symshen_4prolog_1_6shen), V927)

			__e.TailApply(ShenFunc(symmap), gen6643, gen6644)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog"), gen6645)

		gen6652 := MakeNative(func(__e Evaluator) {
			V929 := __e.Get(1)
			_ = V929
			gen6646 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4compile__prolog__procedure), X)

				return
			}, 1)
			gen6647 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4s_1prolog__clause), X)

				return
			}, 1)
			gen6648 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4head__abstraction), X)

				return
			}, 1)
			gen6649 := Call(__e, ShenFunc(symmapcan), gen6648, V929)

			gen6650 := Call(__e, ShenFunc(symmap), gen6647, gen6649)

			gen6651 := Call(__e, ShenFunc(symshen_4group__clauses), gen6650)

			__e.TailApply(ShenFunc(symmap), gen6646, gen6651)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog->shen"), gen6652)

		gen6678 := MakeNative(func(__e Evaluator) {
			V931 := __e.Get(1)
			_ = V931
			gen6676 := Call(__e, ShenFunc(symcons_2), V931)

			var gen6677 Obj
			if True == gen6676 {
				gen6673 := Call(__e, ShenFunc(symtl), V931)

				gen6674 := Call(__e, ShenFunc(symcons_2), gen6673)

				var gen6675 Obj
				if True == gen6674 {
					gen6669 := Call(__e, ShenFunc(symtl), V931)

					gen6670 := Call(__e, ShenFunc(symhd), gen6669)

					gen6671 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen6670)

					var gen6672 Obj
					if True == gen6671 {
						gen6665 := Call(__e, ShenFunc(symtl), V931)

						gen6666 := Call(__e, ShenFunc(symtl), gen6665)

						gen6667 := Call(__e, ShenFunc(symcons_2), gen6666)

						var gen6668 Obj
						if True == gen6667 {
							gen6661 := Call(__e, ShenFunc(symtl), V931)

							gen6662 := Call(__e, ShenFunc(symtl), gen6661)

							gen6663 := Call(__e, ShenFunc(symtl), gen6662)

							gen6664 := Call(__e, ShenFunc(sym_a), Nil, gen6663)

							if True == gen6664 {
								gen6668 = True
							} else {
								gen6668 = False
							}

						} else {
							gen6668 = False
						}
						if True == gen6668 {
							gen6672 = True
						} else {
							gen6672 = False
						}

					} else {
						gen6672 = False
					}
					if True == gen6672 {
						gen6675 = True
					} else {
						gen6675 = False
					}

				} else {
					gen6675 = False
				}
				if True == gen6675 {
					gen6677 = True
				} else {
					gen6677 = False
				}

			} else {
				gen6677 = False
			}
			if True == gen6677 {
				gen6653 := Call(__e, ShenFunc(symhd), V931)

				gen6654 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(ShenFunc(symshen_4s_1prolog__literal), X)

					return
				}, 1)
				gen6655 := Call(__e, ShenFunc(symtl), V931)

				gen6656 := Call(__e, ShenFunc(symtl), gen6655)

				gen6657 := Call(__e, ShenFunc(symhd), gen6656)

				gen6658 := Call(__e, ShenFunc(symmap), gen6654, gen6657)

				gen6659 := Call(__e, ShenFunc(symcons), gen6658, Nil)

				gen6660 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen6659)

				__e.TailApply(ShenFunc(symcons), gen6653, gen6660)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.s-prolog_clause"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog_clause"), gen6678)

		gen6745 := MakeNative(func(__e Evaluator) {
			V933 := __e.Get(1)
			_ = V933
			gen6743 := Call(__e, ShenFunc(symcons_2), V933)

			var gen6744 Obj
			if True == gen6743 {
				gen6740 := Call(__e, ShenFunc(symtl), V933)

				gen6741 := Call(__e, ShenFunc(symcons_2), gen6740)

				var gen6742 Obj
				if True == gen6741 {
					gen6736 := Call(__e, ShenFunc(symtl), V933)

					gen6737 := Call(__e, ShenFunc(symhd), gen6736)

					gen6738 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen6737)

					var gen6739 Obj
					if True == gen6738 {
						gen6732 := Call(__e, ShenFunc(symtl), V933)

						gen6733 := Call(__e, ShenFunc(symtl), gen6732)

						gen6734 := Call(__e, ShenFunc(symcons_2), gen6733)

						var gen6735 Obj
						if True == gen6734 {
							gen6727 := Call(__e, ShenFunc(symtl), V933)

							gen6728 := Call(__e, ShenFunc(symtl), gen6727)

							gen6729 := Call(__e, ShenFunc(symtl), gen6728)

							gen6730 := Call(__e, ShenFunc(sym_a), Nil, gen6729)

							var gen6731 Obj
							if True == gen6730 {
								gen6721 := MakeNative(func(__e Evaluator) {
									__ := __e.Get(1)
									_ = __
									__e.Return(False)
									return
								}, 1)
								gen6725 := MakeNative(func(__e Evaluator) {
									gen6722 := Call(__e, ShenFunc(symhd), V933)

									gen6723 := Call(__e, ShenFunc(symshen_4complexity__head), gen6722)

									gen6724 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*maxcomplexity*"))

									__e.TailApply(ShenFunc(sym_5), gen6723, gen6724)

									return

								}, 0)
								gen6726 := Try(__e, gen6725).Catch(gen6721)
								if True == gen6726 {
									gen6731 = True
								} else {
									gen6731 = False
								}

							} else {
								gen6731 = False
							}
							if True == gen6731 {
								gen6735 = True
							} else {
								gen6735 = False
							}

						} else {
							gen6735 = False
						}
						if True == gen6735 {
							gen6739 = True
						} else {
							gen6739 = False
						}

					} else {
						gen6739 = False
					}
					if True == gen6739 {
						gen6742 = True
					} else {
						gen6742 = False
					}

				} else {
					gen6742 = False
				}
				if True == gen6742 {
					gen6744 = True
				} else {
					gen6744 = False
				}

			} else {
				gen6744 = False
			}
			if True == gen6744 {
				__e.TailApply(ShenFunc(symcons), V933, Nil)

				return
			} else {
				gen6719 := Call(__e, ShenFunc(symcons_2), V933)

				var gen6720 Obj
				if True == gen6719 {
					gen6716 := Call(__e, ShenFunc(symhd), V933)

					gen6717 := Call(__e, ShenFunc(symcons_2), gen6716)

					var gen6718 Obj
					if True == gen6717 {
						gen6713 := Call(__e, ShenFunc(symtl), V933)

						gen6714 := Call(__e, ShenFunc(symcons_2), gen6713)

						var gen6715 Obj
						if True == gen6714 {
							gen6709 := Call(__e, ShenFunc(symtl), V933)

							gen6710 := Call(__e, ShenFunc(symhd), gen6709)

							gen6711 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen6710)

							var gen6712 Obj
							if True == gen6711 {
								gen6705 := Call(__e, ShenFunc(symtl), V933)

								gen6706 := Call(__e, ShenFunc(symtl), gen6705)

								gen6707 := Call(__e, ShenFunc(symcons_2), gen6706)

								var gen6708 Obj
								if True == gen6707 {
									gen6701 := Call(__e, ShenFunc(symtl), V933)

									gen6702 := Call(__e, ShenFunc(symtl), gen6701)

									gen6703 := Call(__e, ShenFunc(symtl), gen6702)

									gen6704 := Call(__e, ShenFunc(sym_a), Nil, gen6703)

									if True == gen6704 {
										gen6708 = True
									} else {
										gen6708 = False
									}

								} else {
									gen6708 = False
								}
								if True == gen6708 {
									gen6712 = True
								} else {
									gen6712 = False
								}

							} else {
								gen6712 = False
							}
							if True == gen6712 {
								gen6715 = True
							} else {
								gen6715 = False
							}

						} else {
							gen6715 = False
						}
						if True == gen6715 {
							gen6718 = True
						} else {
							gen6718 = False
						}

					} else {
						gen6718 = False
					}
					if True == gen6718 {
						gen6720 = True
					} else {
						gen6720 = False
					}

				} else {
					gen6720 = False
				}
				if True == gen6720 {
					gen6679 := MakeNative(func(__e Evaluator) {
						Y := __e.Get(1)
						_ = Y
						__e.TailApply(ShenFunc(symgensym), MakeSymbol("V"))

						return
					}, 1)
					gen6680 := Call(__e, ShenFunc(symhd), V933)

					gen6681 := Call(__e, ShenFunc(symtl), gen6680)

					gen6682 := Call(__e, ShenFunc(symmap), gen6679, gen6681)

					Terms := gen6682
					gen6683 := Call(__e, ShenFunc(symhd), V933)

					gen6684 := Call(__e, ShenFunc(symtl), gen6683)

					gen6685 := Call(__e, ShenFunc(symshen_4remove__modes), gen6684)

					gen6686 := Call(__e, ShenFunc(symshen_4rcons__form), gen6685)

					XTerms := gen6686
					gen6687 := Call(__e, ShenFunc(symshen_4cons__form), Terms)

					gen6688 := Call(__e, ShenFunc(symcons), XTerms, Nil)

					gen6689 := Call(__e, ShenFunc(symcons), gen6687, gen6688)

					gen6690 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen6689)

					Literal := gen6690
					gen6691 := Call(__e, ShenFunc(symhd), V933)

					gen6692 := Call(__e, ShenFunc(symhd), gen6691)

					gen6693 := Call(__e, ShenFunc(symcons), gen6692, Terms)

					gen6694 := Call(__e, ShenFunc(symtl), V933)

					gen6695 := Call(__e, ShenFunc(symtl), gen6694)

					gen6696 := Call(__e, ShenFunc(symhd), gen6695)

					gen6697 := Call(__e, ShenFunc(symcons), Literal, gen6696)

					gen6698 := Call(__e, ShenFunc(symcons), gen6697, Nil)

					gen6699 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen6698)

					gen6700 := Call(__e, ShenFunc(symcons), gen6693, gen6699)

					Clause := gen6700
					__e.TailApply(ShenFunc(symcons), Clause, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.head_abstraction"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.head_abstraction"), gen6745)

		gen6750 := MakeNative(func(__e Evaluator) {
			V939 := __e.Get(1)
			_ = V939
			gen6749 := Call(__e, ShenFunc(symcons_2), V939)

			if True == gen6749 {
				gen6746 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(ShenFunc(symshen_4complexity), X)

					return
				}, 1)
				gen6747 := Call(__e, ShenFunc(symtl), V939)

				gen6748 := Call(__e, ShenFunc(symmap), gen6746, gen6747)

				__e.TailApply(ShenFunc(symshen_4safe_1product), gen6748)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.complexity_head"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.complexity_head"), gen6750)

		gen6751 := MakeNative(func(__e Evaluator) {
			V942 := __e.Get(1)
			_ = V942
			V943 := __e.Get(2)
			_ = V943
			__e.TailApply(ShenFunc(sym_d), V942, V943)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.safe-multiply"), gen6751)

		gen6945 := MakeNative(func(__e Evaluator) {
			V952 := __e.Get(1)
			_ = V952
			gen6943 := Call(__e, ShenFunc(symcons_2), V952)

			var gen6944 Obj
			if True == gen6943 {
				gen6940 := Call(__e, ShenFunc(symhd), V952)

				gen6941 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6940)

				var gen6942 Obj
				if True == gen6941 {
					gen6937 := Call(__e, ShenFunc(symtl), V952)

					gen6938 := Call(__e, ShenFunc(symcons_2), gen6937)

					var gen6939 Obj
					if True == gen6938 {
						gen6933 := Call(__e, ShenFunc(symtl), V952)

						gen6934 := Call(__e, ShenFunc(symhd), gen6933)

						gen6935 := Call(__e, ShenFunc(symcons_2), gen6934)

						var gen6936 Obj
						if True == gen6935 {
							gen6928 := Call(__e, ShenFunc(symtl), V952)

							gen6929 := Call(__e, ShenFunc(symhd), gen6928)

							gen6930 := Call(__e, ShenFunc(symhd), gen6929)

							gen6931 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6930)

							var gen6932 Obj
							if True == gen6931 {
								gen6923 := Call(__e, ShenFunc(symtl), V952)

								gen6924 := Call(__e, ShenFunc(symhd), gen6923)

								gen6925 := Call(__e, ShenFunc(symtl), gen6924)

								gen6926 := Call(__e, ShenFunc(symcons_2), gen6925)

								var gen6927 Obj
								if True == gen6926 {
									gen6917 := Call(__e, ShenFunc(symtl), V952)

									gen6918 := Call(__e, ShenFunc(symhd), gen6917)

									gen6919 := Call(__e, ShenFunc(symtl), gen6918)

									gen6920 := Call(__e, ShenFunc(symtl), gen6919)

									gen6921 := Call(__e, ShenFunc(symcons_2), gen6920)

									var gen6922 Obj
									if True == gen6921 {
										gen6910 := Call(__e, ShenFunc(symtl), V952)

										gen6911 := Call(__e, ShenFunc(symhd), gen6910)

										gen6912 := Call(__e, ShenFunc(symtl), gen6911)

										gen6913 := Call(__e, ShenFunc(symtl), gen6912)

										gen6914 := Call(__e, ShenFunc(symtl), gen6913)

										gen6915 := Call(__e, ShenFunc(sym_a), Nil, gen6914)

										var gen6916 Obj
										if True == gen6915 {
											gen6906 := Call(__e, ShenFunc(symtl), V952)

											gen6907 := Call(__e, ShenFunc(symtl), gen6906)

											gen6908 := Call(__e, ShenFunc(symcons_2), gen6907)

											var gen6909 Obj
											if True == gen6908 {
												gen6902 := Call(__e, ShenFunc(symtl), V952)

												gen6903 := Call(__e, ShenFunc(symtl), gen6902)

												gen6904 := Call(__e, ShenFunc(symtl), gen6903)

												gen6905 := Call(__e, ShenFunc(sym_a), Nil, gen6904)

												if True == gen6905 {
													gen6909 = True
												} else {
													gen6909 = False
												}

											} else {
												gen6909 = False
											}
											if True == gen6909 {
												gen6916 = True
											} else {
												gen6916 = False
											}

										} else {
											gen6916 = False
										}
										if True == gen6916 {
											gen6922 = True
										} else {
											gen6922 = False
										}

									} else {
										gen6922 = False
									}
									if True == gen6922 {
										gen6927 = True
									} else {
										gen6927 = False
									}

								} else {
									gen6927 = False
								}
								if True == gen6927 {
									gen6932 = True
								} else {
									gen6932 = False
								}

							} else {
								gen6932 = False
							}
							if True == gen6932 {
								gen6936 = True
							} else {
								gen6936 = False
							}

						} else {
							gen6936 = False
						}
						if True == gen6936 {
							gen6939 = True
						} else {
							gen6939 = False
						}

					} else {
						gen6939 = False
					}
					if True == gen6939 {
						gen6942 = True
					} else {
						gen6942 = False
					}

				} else {
					gen6942 = False
				}
				if True == gen6942 {
					gen6944 = True
				} else {
					gen6944 = False
				}

			} else {
				gen6944 = False
			}
			if True == gen6944 {
				gen6900 := Call(__e, ShenFunc(symtl), V952)

				gen6901 := Call(__e, ShenFunc(symhd), gen6900)

				__e.TailApply(ShenFunc(symshen_4complexity), gen6901)

				return

			} else {
				gen6898 := Call(__e, ShenFunc(symcons_2), V952)

				var gen6899 Obj
				if True == gen6898 {
					gen6895 := Call(__e, ShenFunc(symhd), V952)

					gen6896 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6895)

					var gen6897 Obj
					if True == gen6896 {
						gen6892 := Call(__e, ShenFunc(symtl), V952)

						gen6893 := Call(__e, ShenFunc(symcons_2), gen6892)

						var gen6894 Obj
						if True == gen6893 {
							gen6888 := Call(__e, ShenFunc(symtl), V952)

							gen6889 := Call(__e, ShenFunc(symhd), gen6888)

							gen6890 := Call(__e, ShenFunc(symcons_2), gen6889)

							var gen6891 Obj
							if True == gen6890 {
								gen6884 := Call(__e, ShenFunc(symtl), V952)

								gen6885 := Call(__e, ShenFunc(symtl), gen6884)

								gen6886 := Call(__e, ShenFunc(symcons_2), gen6885)

								var gen6887 Obj
								if True == gen6886 {
									gen6879 := Call(__e, ShenFunc(symtl), V952)

									gen6880 := Call(__e, ShenFunc(symtl), gen6879)

									gen6881 := Call(__e, ShenFunc(symhd), gen6880)

									gen6882 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen6881)

									var gen6883 Obj
									if True == gen6882 {
										gen6875 := Call(__e, ShenFunc(symtl), V952)

										gen6876 := Call(__e, ShenFunc(symtl), gen6875)

										gen6877 := Call(__e, ShenFunc(symtl), gen6876)

										gen6878 := Call(__e, ShenFunc(sym_a), Nil, gen6877)

										if True == gen6878 {
											gen6883 = True
										} else {
											gen6883 = False
										}

									} else {
										gen6883 = False
									}
									if True == gen6883 {
										gen6887 = True
									} else {
										gen6887 = False
									}

								} else {
									gen6887 = False
								}
								if True == gen6887 {
									gen6891 = True
								} else {
									gen6891 = False
								}

							} else {
								gen6891 = False
							}
							if True == gen6891 {
								gen6894 = True
							} else {
								gen6894 = False
							}

						} else {
							gen6894 = False
						}
						if True == gen6894 {
							gen6897 = True
						} else {
							gen6897 = False
						}

					} else {
						gen6897 = False
					}
					if True == gen6897 {
						gen6899 = True
					} else {
						gen6899 = False
					}

				} else {
					gen6899 = False
				}
				if True == gen6899 {
					gen6858 := Call(__e, ShenFunc(symtl), V952)

					gen6859 := Call(__e, ShenFunc(symhd), gen6858)

					gen6860 := Call(__e, ShenFunc(symhd), gen6859)

					gen6861 := Call(__e, ShenFunc(symtl), V952)

					gen6862 := Call(__e, ShenFunc(symtl), gen6861)

					gen6863 := Call(__e, ShenFunc(symcons), gen6860, gen6862)

					gen6864 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6863)

					gen6865 := Call(__e, ShenFunc(symshen_4complexity), gen6864)

					gen6866 := Call(__e, ShenFunc(symtl), V952)

					gen6867 := Call(__e, ShenFunc(symhd), gen6866)

					gen6868 := Call(__e, ShenFunc(symtl), gen6867)

					gen6869 := Call(__e, ShenFunc(symtl), V952)

					gen6870 := Call(__e, ShenFunc(symtl), gen6869)

					gen6871 := Call(__e, ShenFunc(symcons), gen6868, gen6870)

					gen6872 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6871)

					gen6873 := Call(__e, ShenFunc(symshen_4complexity), gen6872)

					gen6874 := Call(__e, ShenFunc(symshen_4safe_1multiply), gen6865, gen6873)

					__e.TailApply(ShenFunc(symshen_4safe_1multiply), MakeNumber(2), gen6874)

					return

				} else {
					gen6856 := Call(__e, ShenFunc(symcons_2), V952)

					var gen6857 Obj
					if True == gen6856 {
						gen6853 := Call(__e, ShenFunc(symhd), V952)

						gen6854 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6853)

						var gen6855 Obj
						if True == gen6854 {
							gen6850 := Call(__e, ShenFunc(symtl), V952)

							gen6851 := Call(__e, ShenFunc(symcons_2), gen6850)

							var gen6852 Obj
							if True == gen6851 {
								gen6846 := Call(__e, ShenFunc(symtl), V952)

								gen6847 := Call(__e, ShenFunc(symhd), gen6846)

								gen6848 := Call(__e, ShenFunc(symcons_2), gen6847)

								var gen6849 Obj
								if True == gen6848 {
									gen6842 := Call(__e, ShenFunc(symtl), V952)

									gen6843 := Call(__e, ShenFunc(symtl), gen6842)

									gen6844 := Call(__e, ShenFunc(symcons_2), gen6843)

									var gen6845 Obj
									if True == gen6844 {
										gen6837 := Call(__e, ShenFunc(symtl), V952)

										gen6838 := Call(__e, ShenFunc(symtl), gen6837)

										gen6839 := Call(__e, ShenFunc(symhd), gen6838)

										gen6840 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen6839)

										var gen6841 Obj
										if True == gen6840 {
											gen6833 := Call(__e, ShenFunc(symtl), V952)

											gen6834 := Call(__e, ShenFunc(symtl), gen6833)

											gen6835 := Call(__e, ShenFunc(symtl), gen6834)

											gen6836 := Call(__e, ShenFunc(sym_a), Nil, gen6835)

											if True == gen6836 {
												gen6841 = True
											} else {
												gen6841 = False
											}

										} else {
											gen6841 = False
										}
										if True == gen6841 {
											gen6845 = True
										} else {
											gen6845 = False
										}

									} else {
										gen6845 = False
									}
									if True == gen6845 {
										gen6849 = True
									} else {
										gen6849 = False
									}

								} else {
									gen6849 = False
								}
								if True == gen6849 {
									gen6852 = True
								} else {
									gen6852 = False
								}

							} else {
								gen6852 = False
							}
							if True == gen6852 {
								gen6855 = True
							} else {
								gen6855 = False
							}

						} else {
							gen6855 = False
						}
						if True == gen6855 {
							gen6857 = True
						} else {
							gen6857 = False
						}

					} else {
						gen6857 = False
					}
					if True == gen6857 {
						gen6817 := Call(__e, ShenFunc(symtl), V952)

						gen6818 := Call(__e, ShenFunc(symhd), gen6817)

						gen6819 := Call(__e, ShenFunc(symhd), gen6818)

						gen6820 := Call(__e, ShenFunc(symtl), V952)

						gen6821 := Call(__e, ShenFunc(symtl), gen6820)

						gen6822 := Call(__e, ShenFunc(symcons), gen6819, gen6821)

						gen6823 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6822)

						gen6824 := Call(__e, ShenFunc(symshen_4complexity), gen6823)

						gen6825 := Call(__e, ShenFunc(symtl), V952)

						gen6826 := Call(__e, ShenFunc(symhd), gen6825)

						gen6827 := Call(__e, ShenFunc(symtl), gen6826)

						gen6828 := Call(__e, ShenFunc(symtl), V952)

						gen6829 := Call(__e, ShenFunc(symtl), gen6828)

						gen6830 := Call(__e, ShenFunc(symcons), gen6827, gen6829)

						gen6831 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6830)

						gen6832 := Call(__e, ShenFunc(symshen_4complexity), gen6831)

						__e.TailApply(ShenFunc(symshen_4safe_1multiply), gen6824, gen6832)

						return

					} else {
						gen6815 := Call(__e, ShenFunc(symcons_2), V952)

						var gen6816 Obj
						if True == gen6815 {
							gen6812 := Call(__e, ShenFunc(symhd), V952)

							gen6813 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6812)

							var gen6814 Obj
							if True == gen6813 {
								gen6809 := Call(__e, ShenFunc(symtl), V952)

								gen6810 := Call(__e, ShenFunc(symcons_2), gen6809)

								var gen6811 Obj
								if True == gen6810 {
									gen6805 := Call(__e, ShenFunc(symtl), V952)

									gen6806 := Call(__e, ShenFunc(symtl), gen6805)

									gen6807 := Call(__e, ShenFunc(symcons_2), gen6806)

									var gen6808 Obj
									if True == gen6807 {
										gen6800 := Call(__e, ShenFunc(symtl), V952)

										gen6801 := Call(__e, ShenFunc(symtl), gen6800)

										gen6802 := Call(__e, ShenFunc(symtl), gen6801)

										gen6803 := Call(__e, ShenFunc(sym_a), Nil, gen6802)

										var gen6804 Obj
										if True == gen6803 {
											gen6797 := Call(__e, ShenFunc(symtl), V952)

											gen6798 := Call(__e, ShenFunc(symhd), gen6797)

											gen6799 := Call(__e, ShenFunc(symvariable_2), gen6798)

											if True == gen6799 {
												gen6804 = True
											} else {
												gen6804 = False
											}

										} else {
											gen6804 = False
										}
										if True == gen6804 {
											gen6808 = True
										} else {
											gen6808 = False
										}

									} else {
										gen6808 = False
									}
									if True == gen6808 {
										gen6811 = True
									} else {
										gen6811 = False
									}

								} else {
									gen6811 = False
								}
								if True == gen6811 {
									gen6814 = True
								} else {
									gen6814 = False
								}

							} else {
								gen6814 = False
							}
							if True == gen6814 {
								gen6816 = True
							} else {
								gen6816 = False
							}

						} else {
							gen6816 = False
						}
						if True == gen6816 {
							__e.Return(MakeNumber(1))
							return
						} else {
							gen6795 := Call(__e, ShenFunc(symcons_2), V952)

							var gen6796 Obj
							if True == gen6795 {
								gen6792 := Call(__e, ShenFunc(symhd), V952)

								gen6793 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6792)

								var gen6794 Obj
								if True == gen6793 {
									gen6789 := Call(__e, ShenFunc(symtl), V952)

									gen6790 := Call(__e, ShenFunc(symcons_2), gen6789)

									var gen6791 Obj
									if True == gen6790 {
										gen6785 := Call(__e, ShenFunc(symtl), V952)

										gen6786 := Call(__e, ShenFunc(symtl), gen6785)

										gen6787 := Call(__e, ShenFunc(symcons_2), gen6786)

										var gen6788 Obj
										if True == gen6787 {
											gen6780 := Call(__e, ShenFunc(symtl), V952)

											gen6781 := Call(__e, ShenFunc(symtl), gen6780)

											gen6782 := Call(__e, ShenFunc(symhd), gen6781)

											gen6783 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen6782)

											var gen6784 Obj
											if True == gen6783 {
												gen6776 := Call(__e, ShenFunc(symtl), V952)

												gen6777 := Call(__e, ShenFunc(symtl), gen6776)

												gen6778 := Call(__e, ShenFunc(symtl), gen6777)

												gen6779 := Call(__e, ShenFunc(sym_a), Nil, gen6778)

												if True == gen6779 {
													gen6784 = True
												} else {
													gen6784 = False
												}

											} else {
												gen6784 = False
											}
											if True == gen6784 {
												gen6788 = True
											} else {
												gen6788 = False
											}

										} else {
											gen6788 = False
										}
										if True == gen6788 {
											gen6791 = True
										} else {
											gen6791 = False
										}

									} else {
										gen6791 = False
									}
									if True == gen6791 {
										gen6794 = True
									} else {
										gen6794 = False
									}

								} else {
									gen6794 = False
								}
								if True == gen6794 {
									gen6796 = True
								} else {
									gen6796 = False
								}

							} else {
								gen6796 = False
							}
							if True == gen6796 {
								__e.Return(MakeNumber(2))
								return
							} else {
								gen6774 := Call(__e, ShenFunc(symcons_2), V952)

								var gen6775 Obj
								if True == gen6774 {
									gen6771 := Call(__e, ShenFunc(symhd), V952)

									gen6772 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen6771)

									var gen6773 Obj
									if True == gen6772 {
										gen6768 := Call(__e, ShenFunc(symtl), V952)

										gen6769 := Call(__e, ShenFunc(symcons_2), gen6768)

										var gen6770 Obj
										if True == gen6769 {
											gen6764 := Call(__e, ShenFunc(symtl), V952)

											gen6765 := Call(__e, ShenFunc(symtl), gen6764)

											gen6766 := Call(__e, ShenFunc(symcons_2), gen6765)

											var gen6767 Obj
											if True == gen6766 {
												gen6759 := Call(__e, ShenFunc(symtl), V952)

												gen6760 := Call(__e, ShenFunc(symtl), gen6759)

												gen6761 := Call(__e, ShenFunc(symhd), gen6760)

												gen6762 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen6761)

												var gen6763 Obj
												if True == gen6762 {
													gen6755 := Call(__e, ShenFunc(symtl), V952)

													gen6756 := Call(__e, ShenFunc(symtl), gen6755)

													gen6757 := Call(__e, ShenFunc(symtl), gen6756)

													gen6758 := Call(__e, ShenFunc(sym_a), Nil, gen6757)

													if True == gen6758 {
														gen6763 = True
													} else {
														gen6763 = False
													}

												} else {
													gen6763 = False
												}
												if True == gen6763 {
													gen6767 = True
												} else {
													gen6767 = False
												}

											} else {
												gen6767 = False
											}
											if True == gen6767 {
												gen6770 = True
											} else {
												gen6770 = False
											}

										} else {
											gen6770 = False
										}
										if True == gen6770 {
											gen6773 = True
										} else {
											gen6773 = False
										}

									} else {
										gen6773 = False
									}
									if True == gen6773 {
										gen6775 = True
									} else {
										gen6775 = False
									}

								} else {
									gen6775 = False
								}
								if True == gen6775 {
									__e.Return(MakeNumber(1))
									return
								} else {
									gen6752 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

									gen6753 := Call(__e, ShenFunc(symcons), V952, gen6752)

									gen6754 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen6753)

									__e.TailApply(ShenFunc(symshen_4complexity), gen6754)

									return

								}

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.complexity"), gen6945)

		gen6951 := MakeNative(func(__e Evaluator) {
			V954 := __e.Get(1)
			_ = V954
			gen6950 := Call(__e, ShenFunc(sym_a), Nil, V954)

			if True == gen6950 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen6949 := Call(__e, ShenFunc(symcons_2), V954)

				if True == gen6949 {
					gen6946 := Call(__e, ShenFunc(symhd), V954)

					gen6947 := Call(__e, ShenFunc(symtl), V954)

					gen6948 := Call(__e, ShenFunc(symshen_4safe_1product), gen6947)

					__e.TailApply(ShenFunc(symshen_4safe_1multiply), gen6946, gen6948)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.safe-product"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.safe-product"), gen6951)

		gen7031 := MakeNative(func(__e Evaluator) {
			V956 := __e.Get(1)
			_ = V956
			gen7029 := Call(__e, ShenFunc(symcons_2), V956)

			var gen7030 Obj
			if True == gen7029 {
				gen7026 := Call(__e, ShenFunc(symhd), V956)

				gen7027 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen7026)

				var gen7028 Obj
				if True == gen7027 {
					gen7023 := Call(__e, ShenFunc(symtl), V956)

					gen7024 := Call(__e, ShenFunc(symcons_2), gen7023)

					var gen7025 Obj
					if True == gen7024 {
						gen7019 := Call(__e, ShenFunc(symtl), V956)

						gen7020 := Call(__e, ShenFunc(symtl), gen7019)

						gen7021 := Call(__e, ShenFunc(symcons_2), gen7020)

						var gen7022 Obj
						if True == gen7021 {
							gen7015 := Call(__e, ShenFunc(symtl), V956)

							gen7016 := Call(__e, ShenFunc(symtl), gen7015)

							gen7017 := Call(__e, ShenFunc(symtl), gen7016)

							gen7018 := Call(__e, ShenFunc(sym_a), Nil, gen7017)

							if True == gen7018 {
								gen7022 = True
							} else {
								gen7022 = False
							}

						} else {
							gen7022 = False
						}
						if True == gen7022 {
							gen7025 = True
						} else {
							gen7025 = False
						}

					} else {
						gen7025 = False
					}
					if True == gen7025 {
						gen7028 = True
					} else {
						gen7028 = False
					}

				} else {
					gen7028 = False
				}
				if True == gen7028 {
					gen7030 = True
				} else {
					gen7030 = False
				}

			} else {
				gen7030 = False
			}
			if True == gen7030 {
				gen7007 := Call(__e, ShenFunc(symtl), V956)

				gen7008 := Call(__e, ShenFunc(symhd), gen7007)

				gen7009 := Call(__e, ShenFunc(symtl), V956)

				gen7010 := Call(__e, ShenFunc(symtl), gen7009)

				gen7011 := Call(__e, ShenFunc(symhd), gen7010)

				gen7012 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7011, MakeSymbol("ProcessN"))

				gen7013 := Call(__e, ShenFunc(symcons), gen7012, Nil)

				gen7014 := Call(__e, ShenFunc(symcons), gen7008, gen7013)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen7014)

				return

			} else {
				gen7005 := Call(__e, ShenFunc(symcons_2), V956)

				var gen7006 Obj
				if True == gen7005 {
					gen7002 := Call(__e, ShenFunc(symhd), V956)

					gen7003 := Call(__e, ShenFunc(sym_a), MakeSymbol("when"), gen7002)

					var gen7004 Obj
					if True == gen7003 {
						gen6999 := Call(__e, ShenFunc(symtl), V956)

						gen7000 := Call(__e, ShenFunc(symcons_2), gen6999)

						var gen7001 Obj
						if True == gen7000 {
							gen6996 := Call(__e, ShenFunc(symtl), V956)

							gen6997 := Call(__e, ShenFunc(symtl), gen6996)

							gen6998 := Call(__e, ShenFunc(sym_a), Nil, gen6997)

							if True == gen6998 {
								gen7001 = True
							} else {
								gen7001 = False
							}

						} else {
							gen7001 = False
						}
						if True == gen7001 {
							gen7004 = True
						} else {
							gen7004 = False
						}

					} else {
						gen7004 = False
					}
					if True == gen7004 {
						gen7006 = True
					} else {
						gen7006 = False
					}

				} else {
					gen7006 = False
				}
				if True == gen7006 {
					gen6992 := Call(__e, ShenFunc(symtl), V956)

					gen6993 := Call(__e, ShenFunc(symhd), gen6992)

					gen6994 := Call(__e, ShenFunc(symshen_4insert_1deref), gen6993, MakeSymbol("ProcessN"))

					gen6995 := Call(__e, ShenFunc(symcons), gen6994, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen6995)

					return

				} else {
					gen6990 := Call(__e, ShenFunc(symcons_2), V956)

					var gen6991 Obj
					if True == gen6990 {
						gen6987 := Call(__e, ShenFunc(symhd), V956)

						gen6988 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen6987)

						var gen6989 Obj
						if True == gen6988 {
							gen6984 := Call(__e, ShenFunc(symtl), V956)

							gen6985 := Call(__e, ShenFunc(symcons_2), gen6984)

							var gen6986 Obj
							if True == gen6985 {
								gen6980 := Call(__e, ShenFunc(symtl), V956)

								gen6981 := Call(__e, ShenFunc(symtl), gen6980)

								gen6982 := Call(__e, ShenFunc(symcons_2), gen6981)

								var gen6983 Obj
								if True == gen6982 {
									gen6976 := Call(__e, ShenFunc(symtl), V956)

									gen6977 := Call(__e, ShenFunc(symtl), gen6976)

									gen6978 := Call(__e, ShenFunc(symtl), gen6977)

									gen6979 := Call(__e, ShenFunc(sym_a), Nil, gen6978)

									if True == gen6979 {
										gen6983 = True
									} else {
										gen6983 = False
									}

								} else {
									gen6983 = False
								}
								if True == gen6983 {
									gen6986 = True
								} else {
									gen6986 = False
								}

							} else {
								gen6986 = False
							}
							if True == gen6986 {
								gen6989 = True
							} else {
								gen6989 = False
							}

						} else {
							gen6989 = False
						}
						if True == gen6989 {
							gen6991 = True
						} else {
							gen6991 = False
						}

					} else {
						gen6991 = False
					}
					if True == gen6991 {
						gen6968 := Call(__e, ShenFunc(symtl), V956)

						gen6969 := Call(__e, ShenFunc(symhd), gen6968)

						gen6970 := Call(__e, ShenFunc(symtl), V956)

						gen6971 := Call(__e, ShenFunc(symtl), gen6970)

						gen6972 := Call(__e, ShenFunc(symhd), gen6971)

						gen6973 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen6972, MakeSymbol("ProcessN"))

						gen6974 := Call(__e, ShenFunc(symcons), gen6973, Nil)

						gen6975 := Call(__e, ShenFunc(symcons), gen6969, gen6974)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen6975)

						return

					} else {
						gen6966 := Call(__e, ShenFunc(symcons_2), V956)

						var gen6967 Obj
						if True == gen6966 {
							gen6963 := Call(__e, ShenFunc(symhd), V956)

							gen6964 := Call(__e, ShenFunc(sym_a), MakeSymbol("fwhen"), gen6963)

							var gen6965 Obj
							if True == gen6964 {
								gen6960 := Call(__e, ShenFunc(symtl), V956)

								gen6961 := Call(__e, ShenFunc(symcons_2), gen6960)

								var gen6962 Obj
								if True == gen6961 {
									gen6957 := Call(__e, ShenFunc(symtl), V956)

									gen6958 := Call(__e, ShenFunc(symtl), gen6957)

									gen6959 := Call(__e, ShenFunc(sym_a), Nil, gen6958)

									if True == gen6959 {
										gen6962 = True
									} else {
										gen6962 = False
									}

								} else {
									gen6962 = False
								}
								if True == gen6962 {
									gen6965 = True
								} else {
									gen6965 = False
								}

							} else {
								gen6965 = False
							}
							if True == gen6965 {
								gen6967 = True
							} else {
								gen6967 = False
							}

						} else {
							gen6967 = False
						}
						if True == gen6967 {
							gen6953 := Call(__e, ShenFunc(symtl), V956)

							gen6954 := Call(__e, ShenFunc(symhd), gen6953)

							gen6955 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen6954, MakeSymbol("ProcessN"))

							gen6956 := Call(__e, ShenFunc(symcons), gen6955, Nil)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen6956)

							return

						} else {
							gen6952 := Call(__e, ShenFunc(symcons_2), V956)

							if True == gen6952 {
								__e.Return(V956)
								return
							} else {
								__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.s-prolog_literal"))

								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog_literal"), gen7031)

		gen7100 := MakeNative(func(__e Evaluator) {
			V963 := __e.Get(1)
			_ = V963
			V964 := __e.Get(2)
			_ = V964
			gen7099 := Call(__e, ShenFunc(symvariable_2), V963)

			if True == gen7099 {
				gen7097 := Call(__e, ShenFunc(symcons), V964, Nil)

				gen7098 := Call(__e, ShenFunc(symcons), V963, gen7097)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.deref"), gen7098)

				return

			} else {
				gen7095 := Call(__e, ShenFunc(symcons_2), V963)

				var gen7096 Obj
				if True == gen7095 {
					gen7092 := Call(__e, ShenFunc(symhd), V963)

					gen7093 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen7092)

					var gen7094 Obj
					if True == gen7093 {
						gen7089 := Call(__e, ShenFunc(symtl), V963)

						gen7090 := Call(__e, ShenFunc(symcons_2), gen7089)

						var gen7091 Obj
						if True == gen7090 {
							gen7085 := Call(__e, ShenFunc(symtl), V963)

							gen7086 := Call(__e, ShenFunc(symtl), gen7085)

							gen7087 := Call(__e, ShenFunc(symcons_2), gen7086)

							var gen7088 Obj
							if True == gen7087 {
								gen7081 := Call(__e, ShenFunc(symtl), V963)

								gen7082 := Call(__e, ShenFunc(symtl), gen7081)

								gen7083 := Call(__e, ShenFunc(symtl), gen7082)

								gen7084 := Call(__e, ShenFunc(sym_a), Nil, gen7083)

								if True == gen7084 {
									gen7088 = True
								} else {
									gen7088 = False
								}

							} else {
								gen7088 = False
							}
							if True == gen7088 {
								gen7091 = True
							} else {
								gen7091 = False
							}

						} else {
							gen7091 = False
						}
						if True == gen7091 {
							gen7094 = True
						} else {
							gen7094 = False
						}

					} else {
						gen7094 = False
					}
					if True == gen7094 {
						gen7096 = True
					} else {
						gen7096 = False
					}

				} else {
					gen7096 = False
				}
				if True == gen7096 {
					gen7073 := Call(__e, ShenFunc(symtl), V963)

					gen7074 := Call(__e, ShenFunc(symhd), gen7073)

					gen7075 := Call(__e, ShenFunc(symtl), V963)

					gen7076 := Call(__e, ShenFunc(symtl), gen7075)

					gen7077 := Call(__e, ShenFunc(symhd), gen7076)

					gen7078 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7077, V964)

					gen7079 := Call(__e, ShenFunc(symcons), gen7078, Nil)

					gen7080 := Call(__e, ShenFunc(symcons), gen7074, gen7079)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen7080)

					return

				} else {
					gen7071 := Call(__e, ShenFunc(symcons_2), V963)

					var gen7072 Obj
					if True == gen7071 {
						gen7068 := Call(__e, ShenFunc(symhd), V963)

						gen7069 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen7068)

						var gen7070 Obj
						if True == gen7069 {
							gen7065 := Call(__e, ShenFunc(symtl), V963)

							gen7066 := Call(__e, ShenFunc(symcons_2), gen7065)

							var gen7067 Obj
							if True == gen7066 {
								gen7061 := Call(__e, ShenFunc(symtl), V963)

								gen7062 := Call(__e, ShenFunc(symtl), gen7061)

								gen7063 := Call(__e, ShenFunc(symcons_2), gen7062)

								var gen7064 Obj
								if True == gen7063 {
									gen7056 := Call(__e, ShenFunc(symtl), V963)

									gen7057 := Call(__e, ShenFunc(symtl), gen7056)

									gen7058 := Call(__e, ShenFunc(symtl), gen7057)

									gen7059 := Call(__e, ShenFunc(symcons_2), gen7058)

									var gen7060 Obj
									if True == gen7059 {
										gen7051 := Call(__e, ShenFunc(symtl), V963)

										gen7052 := Call(__e, ShenFunc(symtl), gen7051)

										gen7053 := Call(__e, ShenFunc(symtl), gen7052)

										gen7054 := Call(__e, ShenFunc(symtl), gen7053)

										gen7055 := Call(__e, ShenFunc(sym_a), Nil, gen7054)

										if True == gen7055 {
											gen7060 = True
										} else {
											gen7060 = False
										}

									} else {
										gen7060 = False
									}
									if True == gen7060 {
										gen7064 = True
									} else {
										gen7064 = False
									}

								} else {
									gen7064 = False
								}
								if True == gen7064 {
									gen7067 = True
								} else {
									gen7067 = False
								}

							} else {
								gen7067 = False
							}
							if True == gen7067 {
								gen7070 = True
							} else {
								gen7070 = False
							}

						} else {
							gen7070 = False
						}
						if True == gen7070 {
							gen7072 = True
						} else {
							gen7072 = False
						}

					} else {
						gen7072 = False
					}
					if True == gen7072 {
						gen7037 := Call(__e, ShenFunc(symtl), V963)

						gen7038 := Call(__e, ShenFunc(symhd), gen7037)

						gen7039 := Call(__e, ShenFunc(symtl), V963)

						gen7040 := Call(__e, ShenFunc(symtl), gen7039)

						gen7041 := Call(__e, ShenFunc(symhd), gen7040)

						gen7042 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7041, V964)

						gen7043 := Call(__e, ShenFunc(symtl), V963)

						gen7044 := Call(__e, ShenFunc(symtl), gen7043)

						gen7045 := Call(__e, ShenFunc(symtl), gen7044)

						gen7046 := Call(__e, ShenFunc(symhd), gen7045)

						gen7047 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7046, V964)

						gen7048 := Call(__e, ShenFunc(symcons), gen7047, Nil)

						gen7049 := Call(__e, ShenFunc(symcons), gen7042, gen7048)

						gen7050 := Call(__e, ShenFunc(symcons), gen7038, gen7049)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7050)

						return

					} else {
						gen7036 := Call(__e, ShenFunc(symcons_2), V963)

						if True == gen7036 {
							gen7032 := Call(__e, ShenFunc(symhd), V963)

							gen7033 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7032, V964)

							gen7034 := Call(__e, ShenFunc(symtl), V963)

							gen7035 := Call(__e, ShenFunc(symshen_4insert_1deref), gen7034, V964)

							__e.TailApply(ShenFunc(symcons), gen7033, gen7035)

							return

						} else {
							__e.Return(V963)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-deref"), gen7100)

		gen7169 := MakeNative(func(__e Evaluator) {
			V971 := __e.Get(1)
			_ = V971
			V972 := __e.Get(2)
			_ = V972
			gen7168 := Call(__e, ShenFunc(symvariable_2), V971)

			if True == gen7168 {
				gen7166 := Call(__e, ShenFunc(symcons), V972, Nil)

				gen7167 := Call(__e, ShenFunc(symcons), V971, gen7166)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.lazyderef"), gen7167)

				return

			} else {
				gen7164 := Call(__e, ShenFunc(symcons_2), V971)

				var gen7165 Obj
				if True == gen7164 {
					gen7161 := Call(__e, ShenFunc(symhd), V971)

					gen7162 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen7161)

					var gen7163 Obj
					if True == gen7162 {
						gen7158 := Call(__e, ShenFunc(symtl), V971)

						gen7159 := Call(__e, ShenFunc(symcons_2), gen7158)

						var gen7160 Obj
						if True == gen7159 {
							gen7154 := Call(__e, ShenFunc(symtl), V971)

							gen7155 := Call(__e, ShenFunc(symtl), gen7154)

							gen7156 := Call(__e, ShenFunc(symcons_2), gen7155)

							var gen7157 Obj
							if True == gen7156 {
								gen7150 := Call(__e, ShenFunc(symtl), V971)

								gen7151 := Call(__e, ShenFunc(symtl), gen7150)

								gen7152 := Call(__e, ShenFunc(symtl), gen7151)

								gen7153 := Call(__e, ShenFunc(sym_a), Nil, gen7152)

								if True == gen7153 {
									gen7157 = True
								} else {
									gen7157 = False
								}

							} else {
								gen7157 = False
							}
							if True == gen7157 {
								gen7160 = True
							} else {
								gen7160 = False
							}

						} else {
							gen7160 = False
						}
						if True == gen7160 {
							gen7163 = True
						} else {
							gen7163 = False
						}

					} else {
						gen7163 = False
					}
					if True == gen7163 {
						gen7165 = True
					} else {
						gen7165 = False
					}

				} else {
					gen7165 = False
				}
				if True == gen7165 {
					gen7142 := Call(__e, ShenFunc(symtl), V971)

					gen7143 := Call(__e, ShenFunc(symhd), gen7142)

					gen7144 := Call(__e, ShenFunc(symtl), V971)

					gen7145 := Call(__e, ShenFunc(symtl), gen7144)

					gen7146 := Call(__e, ShenFunc(symhd), gen7145)

					gen7147 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen7146, V972)

					gen7148 := Call(__e, ShenFunc(symcons), gen7147, Nil)

					gen7149 := Call(__e, ShenFunc(symcons), gen7143, gen7148)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen7149)

					return

				} else {
					gen7140 := Call(__e, ShenFunc(symcons_2), V971)

					var gen7141 Obj
					if True == gen7140 {
						gen7137 := Call(__e, ShenFunc(symhd), V971)

						gen7138 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen7137)

						var gen7139 Obj
						if True == gen7138 {
							gen7134 := Call(__e, ShenFunc(symtl), V971)

							gen7135 := Call(__e, ShenFunc(symcons_2), gen7134)

							var gen7136 Obj
							if True == gen7135 {
								gen7130 := Call(__e, ShenFunc(symtl), V971)

								gen7131 := Call(__e, ShenFunc(symtl), gen7130)

								gen7132 := Call(__e, ShenFunc(symcons_2), gen7131)

								var gen7133 Obj
								if True == gen7132 {
									gen7125 := Call(__e, ShenFunc(symtl), V971)

									gen7126 := Call(__e, ShenFunc(symtl), gen7125)

									gen7127 := Call(__e, ShenFunc(symtl), gen7126)

									gen7128 := Call(__e, ShenFunc(symcons_2), gen7127)

									var gen7129 Obj
									if True == gen7128 {
										gen7120 := Call(__e, ShenFunc(symtl), V971)

										gen7121 := Call(__e, ShenFunc(symtl), gen7120)

										gen7122 := Call(__e, ShenFunc(symtl), gen7121)

										gen7123 := Call(__e, ShenFunc(symtl), gen7122)

										gen7124 := Call(__e, ShenFunc(sym_a), Nil, gen7123)

										if True == gen7124 {
											gen7129 = True
										} else {
											gen7129 = False
										}

									} else {
										gen7129 = False
									}
									if True == gen7129 {
										gen7133 = True
									} else {
										gen7133 = False
									}

								} else {
									gen7133 = False
								}
								if True == gen7133 {
									gen7136 = True
								} else {
									gen7136 = False
								}

							} else {
								gen7136 = False
							}
							if True == gen7136 {
								gen7139 = True
							} else {
								gen7139 = False
							}

						} else {
							gen7139 = False
						}
						if True == gen7139 {
							gen7141 = True
						} else {
							gen7141 = False
						}

					} else {
						gen7141 = False
					}
					if True == gen7141 {
						gen7106 := Call(__e, ShenFunc(symtl), V971)

						gen7107 := Call(__e, ShenFunc(symhd), gen7106)

						gen7108 := Call(__e, ShenFunc(symtl), V971)

						gen7109 := Call(__e, ShenFunc(symtl), gen7108)

						gen7110 := Call(__e, ShenFunc(symhd), gen7109)

						gen7111 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen7110, V972)

						gen7112 := Call(__e, ShenFunc(symtl), V971)

						gen7113 := Call(__e, ShenFunc(symtl), gen7112)

						gen7114 := Call(__e, ShenFunc(symtl), gen7113)

						gen7115 := Call(__e, ShenFunc(symhd), gen7114)

						gen7116 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen7115, V972)

						gen7117 := Call(__e, ShenFunc(symcons), gen7116, Nil)

						gen7118 := Call(__e, ShenFunc(symcons), gen7111, gen7117)

						gen7119 := Call(__e, ShenFunc(symcons), gen7107, gen7118)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7119)

						return

					} else {
						gen7105 := Call(__e, ShenFunc(symcons_2), V971)

						if True == gen7105 {
							gen7101 := Call(__e, ShenFunc(symhd), V971)

							gen7102 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen7101, V972)

							gen7103 := Call(__e, ShenFunc(symtl), V971)

							gen7104 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen7103, V972)

							__e.TailApply(ShenFunc(symcons), gen7102, gen7104)

							return

						} else {
							__e.Return(V971)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-lazyderef"), gen7169)

		gen7177 := MakeNative(func(__e Evaluator) {
			V974 := __e.Get(1)
			_ = V974
			gen7176 := Call(__e, ShenFunc(sym_a), Nil, V974)

			if True == gen7176 {
				__e.Return(Nil)
				return
			} else {
				gen7175 := Call(__e, ShenFunc(symcons_2), V974)

				if True == gen7175 {
					gen7171 := MakeNative(func(__e Evaluator) {
						X := __e.Get(1)
						_ = X
						gen7170 := Call(__e, ShenFunc(symhd), V974)

						__e.TailApply(ShenFunc(symshen_4same__predicate_2), gen7170, X)

						return

					}, 1)
					gen7172 := Call(__e, ShenFunc(symshen_4collect), gen7171, V974)

					Group := gen7172
					gen7173 := Call(__e, ShenFunc(symdifference), V974, Group)

					Rest := gen7173
					gen7174 := Call(__e, ShenFunc(symshen_4group__clauses), Rest)

					__e.TailApply(ShenFunc(symcons), Group, gen7174)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.group_clauses"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.group_clauses"), gen7177)

		gen7186 := MakeNative(func(__e Evaluator) {
			V979 := __e.Get(1)
			_ = V979
			V980 := __e.Get(2)
			_ = V980
			gen7185 := Call(__e, ShenFunc(sym_a), Nil, V980)

			if True == gen7185 {
				__e.Return(Nil)
				return
			} else {
				gen7184 := Call(__e, ShenFunc(symcons_2), V980)

				if True == gen7184 {
					gen7182 := Call(__e, ShenFunc(symhd), V980)

					gen7183 := Call(__e, V979, gen7182)

					if True == gen7183 {
						gen7179 := Call(__e, ShenFunc(symhd), V980)

						gen7180 := Call(__e, ShenFunc(symtl), V980)

						gen7181 := Call(__e, ShenFunc(symshen_4collect), V979, gen7180)

						__e.TailApply(ShenFunc(symcons), gen7179, gen7181)

						return

					} else {
						gen7178 := Call(__e, ShenFunc(symtl), V980)

						__e.TailApply(ShenFunc(symshen_4collect), V979, gen7178)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.collect"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.collect"), gen7186)

		gen7200 := MakeNative(func(__e Evaluator) {
			V999 := __e.Get(1)
			_ = V999
			V1000 := __e.Get(2)
			_ = V1000
			gen7198 := Call(__e, ShenFunc(symcons_2), V999)

			var gen7199 Obj
			if True == gen7198 {
				gen7195 := Call(__e, ShenFunc(symhd), V999)

				gen7196 := Call(__e, ShenFunc(symcons_2), gen7195)

				var gen7197 Obj
				if True == gen7196 {
					gen7193 := Call(__e, ShenFunc(symcons_2), V1000)

					var gen7194 Obj
					if True == gen7193 {
						gen7191 := Call(__e, ShenFunc(symhd), V1000)

						gen7192 := Call(__e, ShenFunc(symcons_2), gen7191)

						if True == gen7192 {
							gen7194 = True
						} else {
							gen7194 = False
						}

					} else {
						gen7194 = False
					}
					if True == gen7194 {
						gen7197 = True
					} else {
						gen7197 = False
					}

				} else {
					gen7197 = False
				}
				if True == gen7197 {
					gen7199 = True
				} else {
					gen7199 = False
				}

			} else {
				gen7199 = False
			}
			if True == gen7199 {
				gen7187 := Call(__e, ShenFunc(symhd), V999)

				gen7188 := Call(__e, ShenFunc(symhd), gen7187)

				gen7189 := Call(__e, ShenFunc(symhd), V1000)

				gen7190 := Call(__e, ShenFunc(symhd), gen7189)

				__e.TailApply(ShenFunc(sym_a), gen7188, gen7190)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.same_predicate?"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.same_predicate?"), gen7200)

		gen7203 := MakeNative(func(__e Evaluator) {
			V1002 := __e.Get(1)
			_ = V1002
			gen7201 := Call(__e, ShenFunc(symshen_4procedure__name), V1002)

			F := gen7201
			gen7202 := Call(__e, ShenFunc(symshen_4clauses_1to_1shen), F, V1002)

			Shen := gen7202
			__e.Return(Shen)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_prolog_procedure"), gen7203)

		gen7214 := MakeNative(func(__e Evaluator) {
			V1016 := __e.Get(1)
			_ = V1016
			gen7212 := Call(__e, ShenFunc(symcons_2), V1016)

			var gen7213 Obj
			if True == gen7212 {
				gen7209 := Call(__e, ShenFunc(symhd), V1016)

				gen7210 := Call(__e, ShenFunc(symcons_2), gen7209)

				var gen7211 Obj
				if True == gen7210 {
					gen7206 := Call(__e, ShenFunc(symhd), V1016)

					gen7207 := Call(__e, ShenFunc(symhd), gen7206)

					gen7208 := Call(__e, ShenFunc(symcons_2), gen7207)

					if True == gen7208 {
						gen7211 = True
					} else {
						gen7211 = False
					}

				} else {
					gen7211 = False
				}
				if True == gen7211 {
					gen7213 = True
				} else {
					gen7213 = False
				}

			} else {
				gen7213 = False
			}
			if True == gen7213 {
				gen7204 := Call(__e, ShenFunc(symhd), V1016)

				gen7205 := Call(__e, ShenFunc(symhd), gen7204)

				__e.TailApply(ShenFunc(symhd), gen7205)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.procedure_name"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.procedure_name"), gen7214)

		gen7235 := MakeNative(func(__e Evaluator) {
			V1019 := __e.Get(1)
			_ = V1019
			V1020 := __e.Get(2)
			_ = V1020
			gen7215 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4linearise_1clause), X)

				return
			}, 1)
			gen7216 := Call(__e, ShenFunc(symmap), gen7215, V1020)

			Linear := gen7216
			gen7217 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symhead), X)

				return
			}, 1)
			gen7218 := Call(__e, ShenFunc(symmap), gen7217, V1020)

			gen7219 := Call(__e, ShenFunc(symshen_4prolog_1aritycheck), V1019, gen7218)

			Arity := gen7219
			gen7220 := Call(__e, ShenFunc(symshen_4parameters), Arity)

			Parameters := gen7220
			gen7221 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4aum), X, Parameters)

				return
			}, 1)
			gen7222 := Call(__e, ShenFunc(symmap), gen7221, Linear)

			AUM__instructions := gen7222
			gen7223 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4aum__to__shen), X)

				return
			}, 1)
			gen7224 := Call(__e, ShenFunc(symmap), gen7223, AUM__instructions)

			gen7225 := Call(__e, ShenFunc(symshen_4nest_1disjunct), gen7224)

			gen7226 := Call(__e, ShenFunc(symshen_4catch_1cut), gen7225)

			Code := gen7226
			gen7227 := Call(__e, ShenFunc(symcons), MakeSymbol("Continuation"), Nil)

			gen7228 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), gen7227)

			gen7229 := Call(__e, ShenFunc(symcons), Code, Nil)

			gen7230 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen7229)

			gen7231 := Call(__e, ShenFunc(symappend), gen7228, gen7230)

			gen7232 := Call(__e, ShenFunc(symappend), Parameters, gen7231)

			gen7233 := Call(__e, ShenFunc(symcons), V1019, gen7232)

			gen7234 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen7233)

			ShenDef := gen7234
			__e.Return(ShenDef)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clauses-to-shen"), gen7235)

		gen7245 := MakeNative(func(__e Evaluator) {
			V1022 := __e.Get(1)
			_ = V1022
			gen7243 := Call(__e, ShenFunc(symshen_4occurs_2), MakeSymbol("cut"), V1022)

			gen7244 := Call(__e, ShenFunc(symnot), gen7243)

			if True == gen7244 {
				__e.Return(V1022)
				return
			} else {
				gen7236 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.catchpoint"), Nil)

				gen7237 := Call(__e, ShenFunc(symcons), V1022, Nil)

				gen7238 := Call(__e, ShenFunc(symcons), MakeSymbol("Throwcontrol"), gen7237)

				gen7239 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.cutpoint"), gen7238)

				gen7240 := Call(__e, ShenFunc(symcons), gen7239, Nil)

				gen7241 := Call(__e, ShenFunc(symcons), gen7236, gen7240)

				gen7242 := Call(__e, ShenFunc(symcons), MakeSymbol("Throwcontrol"), gen7241)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7242)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.catch-cut"), gen7245)

		gen7249 := MakeNative(func(__e Evaluator) {
			gen7246 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*catch*"))

			gen7247 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen7246)

			gen7248 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*catch*"), gen7247)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.catchpoint!"), gen7248)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.catchpoint"), gen7249)

		gen7251 := MakeNative(func(__e Evaluator) {
			V1030 := __e.Get(1)
			_ = V1030
			V1031 := __e.Get(2)
			_ = V1031
			gen7250 := Call(__e, ShenFunc(sym_a), V1031, V1030)

			if True == gen7250 {
				__e.Return(False)
				return
			} else {
				__e.Return(V1031)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cutpoint"), gen7251)

		gen7260 := MakeNative(func(__e Evaluator) {
			V1033 := __e.Get(1)
			_ = V1033
			gen7258 := Call(__e, ShenFunc(symcons_2), V1033)

			var gen7259 Obj
			if True == gen7258 {
				gen7256 := Call(__e, ShenFunc(symtl), V1033)

				gen7257 := Call(__e, ShenFunc(sym_a), Nil, gen7256)

				if True == gen7257 {
					gen7259 = True
				} else {
					gen7259 = False
				}

			} else {
				gen7259 = False
			}
			if True == gen7259 {
				__e.TailApply(ShenFunc(symhd), V1033)

				return
			} else {
				gen7255 := Call(__e, ShenFunc(symcons_2), V1033)

				if True == gen7255 {
					gen7252 := Call(__e, ShenFunc(symhd), V1033)

					gen7253 := Call(__e, ShenFunc(symtl), V1033)

					gen7254 := Call(__e, ShenFunc(symshen_4nest_1disjunct), gen7253)

					__e.TailApply(ShenFunc(symshen_4lisp_1or), gen7252, gen7254)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.nest-disjunct"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.nest-disjunct"), gen7260)

		gen7271 := MakeNative(func(__e Evaluator) {
			V1036 := __e.Get(1)
			_ = V1036
			V1037 := __e.Get(2)
			_ = V1037
			gen7261 := Call(__e, ShenFunc(symcons), False, Nil)

			gen7262 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), gen7261)

			gen7263 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen7262)

			gen7264 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), Nil)

			gen7265 := Call(__e, ShenFunc(symcons), V1037, gen7264)

			gen7266 := Call(__e, ShenFunc(symcons), gen7263, gen7265)

			gen7267 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7266)

			gen7268 := Call(__e, ShenFunc(symcons), gen7267, Nil)

			gen7269 := Call(__e, ShenFunc(symcons), V1036, gen7268)

			gen7270 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), gen7269)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7270)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lisp-or"), gen7271)

		gen7292 := MakeNative(func(__e Evaluator) {
			V1042 := __e.Get(1)
			_ = V1042
			V1043 := __e.Get(2)
			_ = V1043
			gen7290 := Call(__e, ShenFunc(symcons_2), V1043)

			var gen7291 Obj
			if True == gen7290 {
				gen7288 := Call(__e, ShenFunc(symtl), V1043)

				gen7289 := Call(__e, ShenFunc(sym_a), Nil, gen7288)

				if True == gen7289 {
					gen7291 = True
				} else {
					gen7291 = False
				}

			} else {
				gen7291 = False
			}
			if True == gen7291 {
				gen7286 := Call(__e, ShenFunc(symhd), V1043)

				gen7287 := Call(__e, ShenFunc(symlength), gen7286)

				__e.TailApply(ShenFunc(sym_1), gen7287, MakeNumber(1))

				return

			} else {
				gen7284 := Call(__e, ShenFunc(symcons_2), V1043)

				var gen7285 Obj
				if True == gen7284 {
					gen7282 := Call(__e, ShenFunc(symtl), V1043)

					gen7283 := Call(__e, ShenFunc(symcons_2), gen7282)

					if True == gen7283 {
						gen7285 = True
					} else {
						gen7285 = False
					}

				} else {
					gen7285 = False
				}
				if True == gen7285 {
					gen7276 := Call(__e, ShenFunc(symhd), V1043)

					gen7277 := Call(__e, ShenFunc(symlength), gen7276)

					gen7278 := Call(__e, ShenFunc(symtl), V1043)

					gen7279 := Call(__e, ShenFunc(symhd), gen7278)

					gen7280 := Call(__e, ShenFunc(symlength), gen7279)

					gen7281 := Call(__e, ShenFunc(sym_a), gen7277, gen7280)

					if True == gen7281 {
						gen7275 := Call(__e, ShenFunc(symtl), V1043)

						__e.TailApply(ShenFunc(symshen_4prolog_1aritycheck), V1042, gen7275)

						return

					} else {
						gen7272 := Call(__e, ShenFunc(symcons), V1042, Nil)

						gen7273 := Call(__e, ShenFunc(symshen_4app), gen7272, MakeString("\n"), MakeSymbol("shen.a"))

						gen7274 := Call(__e, ShenFunc(symcn), MakeString("arity error in prolog procedure "), gen7273)

						__e.TailApply(ShenFunc(symsimple_1error), gen7274)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.prolog-aritycheck"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-aritycheck"), gen7292)

		gen7315 := MakeNative(func(__e Evaluator) {
			V1045 := __e.Get(1)
			_ = V1045
			gen7313 := Call(__e, ShenFunc(symcons_2), V1045)

			var gen7314 Obj
			if True == gen7313 {
				gen7310 := Call(__e, ShenFunc(symtl), V1045)

				gen7311 := Call(__e, ShenFunc(symcons_2), gen7310)

				var gen7312 Obj
				if True == gen7311 {
					gen7306 := Call(__e, ShenFunc(symtl), V1045)

					gen7307 := Call(__e, ShenFunc(symhd), gen7306)

					gen7308 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen7307)

					var gen7309 Obj
					if True == gen7308 {
						gen7302 := Call(__e, ShenFunc(symtl), V1045)

						gen7303 := Call(__e, ShenFunc(symtl), gen7302)

						gen7304 := Call(__e, ShenFunc(symcons_2), gen7303)

						var gen7305 Obj
						if True == gen7304 {
							gen7298 := Call(__e, ShenFunc(symtl), V1045)

							gen7299 := Call(__e, ShenFunc(symtl), gen7298)

							gen7300 := Call(__e, ShenFunc(symtl), gen7299)

							gen7301 := Call(__e, ShenFunc(sym_a), Nil, gen7300)

							if True == gen7301 {
								gen7305 = True
							} else {
								gen7305 = False
							}

						} else {
							gen7305 = False
						}
						if True == gen7305 {
							gen7309 = True
						} else {
							gen7309 = False
						}

					} else {
						gen7309 = False
					}
					if True == gen7309 {
						gen7312 = True
					} else {
						gen7312 = False
					}

				} else {
					gen7312 = False
				}
				if True == gen7312 {
					gen7314 = True
				} else {
					gen7314 = False
				}

			} else {
				gen7314 = False
			}
			if True == gen7314 {
				gen7293 := Call(__e, ShenFunc(symhd), V1045)

				gen7294 := Call(__e, ShenFunc(symtl), V1045)

				gen7295 := Call(__e, ShenFunc(symtl), gen7294)

				gen7296 := Call(__e, ShenFunc(symcons), gen7293, gen7295)

				gen7297 := Call(__e, ShenFunc(symshen_4linearise), gen7296)

				Linear := gen7297
				__e.TailApply(ShenFunc(symshen_4clause__form), Linear)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise-clause"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise-clause"), gen7315)

		gen7331 := MakeNative(func(__e Evaluator) {
			V1047 := __e.Get(1)
			_ = V1047
			gen7329 := Call(__e, ShenFunc(symcons_2), V1047)

			var gen7330 Obj
			if True == gen7329 {
				gen7326 := Call(__e, ShenFunc(symtl), V1047)

				gen7327 := Call(__e, ShenFunc(symcons_2), gen7326)

				var gen7328 Obj
				if True == gen7327 {
					gen7323 := Call(__e, ShenFunc(symtl), V1047)

					gen7324 := Call(__e, ShenFunc(symtl), gen7323)

					gen7325 := Call(__e, ShenFunc(sym_a), Nil, gen7324)

					if True == gen7325 {
						gen7328 = True
					} else {
						gen7328 = False
					}

				} else {
					gen7328 = False
				}
				if True == gen7328 {
					gen7330 = True
				} else {
					gen7330 = False
				}

			} else {
				gen7330 = False
			}
			if True == gen7330 {
				gen7316 := Call(__e, ShenFunc(symhd), V1047)

				gen7317 := Call(__e, ShenFunc(symshen_4explicit__modes), gen7316)

				gen7318 := Call(__e, ShenFunc(symtl), V1047)

				gen7319 := Call(__e, ShenFunc(symhd), gen7318)

				gen7320 := Call(__e, ShenFunc(symshen_4cf__help), gen7319)

				gen7321 := Call(__e, ShenFunc(symcons), gen7320, Nil)

				gen7322 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen7321)

				__e.TailApply(ShenFunc(symcons), gen7317, gen7322)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.clause_form"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clause_form"), gen7331)

		gen7337 := MakeNative(func(__e Evaluator) {
			V1049 := __e.Get(1)
			_ = V1049
			gen7336 := Call(__e, ShenFunc(symcons_2), V1049)

			if True == gen7336 {
				gen7332 := Call(__e, ShenFunc(symhd), V1049)

				gen7333 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(ShenFunc(symshen_4em__help), X)

					return
				}, 1)
				gen7334 := Call(__e, ShenFunc(symtl), V1049)

				gen7335 := Call(__e, ShenFunc(symmap), gen7333, gen7334)

				__e.TailApply(ShenFunc(symcons), gen7332, gen7335)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.explicit_modes"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.explicit_modes"), gen7337)

		gen7356 := MakeNative(func(__e Evaluator) {
			V1051 := __e.Get(1)
			_ = V1051
			gen7354 := Call(__e, ShenFunc(symcons_2), V1051)

			var gen7355 Obj
			if True == gen7354 {
				gen7351 := Call(__e, ShenFunc(symhd), V1051)

				gen7352 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen7351)

				var gen7353 Obj
				if True == gen7352 {
					gen7348 := Call(__e, ShenFunc(symtl), V1051)

					gen7349 := Call(__e, ShenFunc(symcons_2), gen7348)

					var gen7350 Obj
					if True == gen7349 {
						gen7344 := Call(__e, ShenFunc(symtl), V1051)

						gen7345 := Call(__e, ShenFunc(symtl), gen7344)

						gen7346 := Call(__e, ShenFunc(symcons_2), gen7345)

						var gen7347 Obj
						if True == gen7346 {
							gen7340 := Call(__e, ShenFunc(symtl), V1051)

							gen7341 := Call(__e, ShenFunc(symtl), gen7340)

							gen7342 := Call(__e, ShenFunc(symtl), gen7341)

							gen7343 := Call(__e, ShenFunc(sym_a), Nil, gen7342)

							if True == gen7343 {
								gen7347 = True
							} else {
								gen7347 = False
							}

						} else {
							gen7347 = False
						}
						if True == gen7347 {
							gen7350 = True
						} else {
							gen7350 = False
						}

					} else {
						gen7350 = False
					}
					if True == gen7350 {
						gen7353 = True
					} else {
						gen7353 = False
					}

				} else {
					gen7353 = False
				}
				if True == gen7353 {
					gen7355 = True
				} else {
					gen7355 = False
				}

			} else {
				gen7355 = False
			}
			if True == gen7355 {
				__e.Return(V1051)
				return
			} else {
				gen7338 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

				gen7339 := Call(__e, ShenFunc(symcons), V1051, gen7338)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("mode"), gen7339)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.em_help"), gen7356)

		gen7410 := MakeNative(func(__e Evaluator) {
			V1053 := __e.Get(1)
			_ = V1053
			gen7408 := Call(__e, ShenFunc(symcons_2), V1053)

			var gen7409 Obj
			if True == gen7408 {
				gen7405 := Call(__e, ShenFunc(symhd), V1053)

				gen7406 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen7405)

				var gen7407 Obj
				if True == gen7406 {
					gen7402 := Call(__e, ShenFunc(symtl), V1053)

					gen7403 := Call(__e, ShenFunc(symcons_2), gen7402)

					var gen7404 Obj
					if True == gen7403 {
						gen7398 := Call(__e, ShenFunc(symtl), V1053)

						gen7399 := Call(__e, ShenFunc(symhd), gen7398)

						gen7400 := Call(__e, ShenFunc(symcons_2), gen7399)

						var gen7401 Obj
						if True == gen7400 {
							gen7393 := Call(__e, ShenFunc(symtl), V1053)

							gen7394 := Call(__e, ShenFunc(symhd), gen7393)

							gen7395 := Call(__e, ShenFunc(symhd), gen7394)

							gen7396 := Call(__e, ShenFunc(sym_a), MakeSymbol("="), gen7395)

							var gen7397 Obj
							if True == gen7396 {
								gen7388 := Call(__e, ShenFunc(symtl), V1053)

								gen7389 := Call(__e, ShenFunc(symhd), gen7388)

								gen7390 := Call(__e, ShenFunc(symtl), gen7389)

								gen7391 := Call(__e, ShenFunc(symcons_2), gen7390)

								var gen7392 Obj
								if True == gen7391 {
									gen7382 := Call(__e, ShenFunc(symtl), V1053)

									gen7383 := Call(__e, ShenFunc(symhd), gen7382)

									gen7384 := Call(__e, ShenFunc(symtl), gen7383)

									gen7385 := Call(__e, ShenFunc(symtl), gen7384)

									gen7386 := Call(__e, ShenFunc(symcons_2), gen7385)

									var gen7387 Obj
									if True == gen7386 {
										gen7375 := Call(__e, ShenFunc(symtl), V1053)

										gen7376 := Call(__e, ShenFunc(symhd), gen7375)

										gen7377 := Call(__e, ShenFunc(symtl), gen7376)

										gen7378 := Call(__e, ShenFunc(symtl), gen7377)

										gen7379 := Call(__e, ShenFunc(symtl), gen7378)

										gen7380 := Call(__e, ShenFunc(sym_a), Nil, gen7379)

										var gen7381 Obj
										if True == gen7380 {
											gen7371 := Call(__e, ShenFunc(symtl), V1053)

											gen7372 := Call(__e, ShenFunc(symtl), gen7371)

											gen7373 := Call(__e, ShenFunc(symcons_2), gen7372)

											var gen7374 Obj
											if True == gen7373 {
												gen7367 := Call(__e, ShenFunc(symtl), V1053)

												gen7368 := Call(__e, ShenFunc(symtl), gen7367)

												gen7369 := Call(__e, ShenFunc(symtl), gen7368)

												gen7370 := Call(__e, ShenFunc(sym_a), Nil, gen7369)

												if True == gen7370 {
													gen7374 = True
												} else {
													gen7374 = False
												}

											} else {
												gen7374 = False
											}
											if True == gen7374 {
												gen7381 = True
											} else {
												gen7381 = False
											}

										} else {
											gen7381 = False
										}
										if True == gen7381 {
											gen7387 = True
										} else {
											gen7387 = False
										}

									} else {
										gen7387 = False
									}
									if True == gen7387 {
										gen7392 = True
									} else {
										gen7392 = False
									}

								} else {
									gen7392 = False
								}
								if True == gen7392 {
									gen7397 = True
								} else {
									gen7397 = False
								}

							} else {
								gen7397 = False
							}
							if True == gen7397 {
								gen7401 = True
							} else {
								gen7401 = False
							}

						} else {
							gen7401 = False
						}
						if True == gen7401 {
							gen7404 = True
						} else {
							gen7404 = False
						}

					} else {
						gen7404 = False
					}
					if True == gen7404 {
						gen7407 = True
					} else {
						gen7407 = False
					}

				} else {
					gen7407 = False
				}
				if True == gen7407 {
					gen7409 = True
				} else {
					gen7409 = False
				}

			} else {
				gen7409 = False
			}
			if True == gen7409 {
				gen7357 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*occurs*"))

				var gen7358 Obj
				if True == gen7357 {
					gen7358 = MakeSymbol("unify!")
				} else {
					gen7358 = MakeSymbol("unify")
				}
				gen7359 := Call(__e, ShenFunc(symtl), V1053)

				gen7360 := Call(__e, ShenFunc(symhd), gen7359)

				gen7361 := Call(__e, ShenFunc(symtl), gen7360)

				gen7362 := Call(__e, ShenFunc(symcons), gen7358, gen7361)

				gen7363 := Call(__e, ShenFunc(symtl), V1053)

				gen7364 := Call(__e, ShenFunc(symtl), gen7363)

				gen7365 := Call(__e, ShenFunc(symhd), gen7364)

				gen7366 := Call(__e, ShenFunc(symshen_4cf__help), gen7365)

				__e.TailApply(ShenFunc(symcons), gen7362, gen7366)

				return

			} else {
				__e.Return(V1053)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cf_help"), gen7410)

		gen7413 := MakeNative(func(__e Evaluator) {
			V1059 := __e.Get(1)
			_ = V1059
			gen7412 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1059)

			if True == gen7412 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*occurs*"), True)

				return
			} else {
				gen7411 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1059)

				if True == gen7411 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*occurs*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("occurs-check expects + or -\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("occurs-check"), gen7413)

		gen7446 := MakeNative(func(__e Evaluator) {
			V1062 := __e.Get(1)
			_ = V1062
			V1063 := __e.Get(2)
			_ = V1063
			gen7444 := Call(__e, ShenFunc(symcons_2), V1062)

			var gen7445 Obj
			if True == gen7444 {
				gen7441 := Call(__e, ShenFunc(symhd), V1062)

				gen7442 := Call(__e, ShenFunc(symcons_2), gen7441)

				var gen7443 Obj
				if True == gen7442 {
					gen7438 := Call(__e, ShenFunc(symtl), V1062)

					gen7439 := Call(__e, ShenFunc(symcons_2), gen7438)

					var gen7440 Obj
					if True == gen7439 {
						gen7434 := Call(__e, ShenFunc(symtl), V1062)

						gen7435 := Call(__e, ShenFunc(symhd), gen7434)

						gen7436 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen7435)

						var gen7437 Obj
						if True == gen7436 {
							gen7430 := Call(__e, ShenFunc(symtl), V1062)

							gen7431 := Call(__e, ShenFunc(symtl), gen7430)

							gen7432 := Call(__e, ShenFunc(symcons_2), gen7431)

							var gen7433 Obj
							if True == gen7432 {
								gen7426 := Call(__e, ShenFunc(symtl), V1062)

								gen7427 := Call(__e, ShenFunc(symtl), gen7426)

								gen7428 := Call(__e, ShenFunc(symtl), gen7427)

								gen7429 := Call(__e, ShenFunc(sym_a), Nil, gen7428)

								if True == gen7429 {
									gen7433 = True
								} else {
									gen7433 = False
								}

							} else {
								gen7433 = False
							}
							if True == gen7433 {
								gen7437 = True
							} else {
								gen7437 = False
							}

						} else {
							gen7437 = False
						}
						if True == gen7437 {
							gen7440 = True
						} else {
							gen7440 = False
						}

					} else {
						gen7440 = False
					}
					if True == gen7440 {
						gen7443 = True
					} else {
						gen7443 = False
					}

				} else {
					gen7443 = False
				}
				if True == gen7443 {
					gen7445 = True
				} else {
					gen7445 = False
				}

			} else {
				gen7445 = False
			}
			if True == gen7445 {
				gen7414 := Call(__e, ShenFunc(symhd), V1062)

				gen7415 := Call(__e, ShenFunc(symtl), gen7414)

				gen7416 := Call(__e, ShenFunc(symhd), V1062)

				gen7417 := Call(__e, ShenFunc(symtl), gen7416)

				gen7418 := Call(__e, ShenFunc(symtl), V1062)

				gen7419 := Call(__e, ShenFunc(symtl), gen7418)

				gen7420 := Call(__e, ShenFunc(symhd), gen7419)

				gen7421 := Call(__e, ShenFunc(symshen_4continuation__call), gen7417, gen7420)

				gen7422 := Call(__e, ShenFunc(symcons), gen7421, Nil)

				gen7423 := Call(__e, ShenFunc(symcons), gen7415, gen7422)

				gen7424 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7423)

				gen7425 := Call(__e, ShenFunc(symshen_4make__mu__application), gen7424, V1063)

				MuApplication := gen7425
				__e.TailApply(ShenFunc(symshen_4mu__reduction), MuApplication, MakeSymbol("+"))

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.aum"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aum"), gen7446)

		gen7452 := MakeNative(func(__e Evaluator) {
			V1066 := __e.Get(1)
			_ = V1066
			V1067 := __e.Get(2)
			_ = V1067
			gen7447 := Call(__e, ShenFunc(symshen_4extract__vars), V1066)

			gen7448 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), gen7447)

			VTerms := gen7448
			gen7449 := Call(__e, ShenFunc(symshen_4extract__vars), V1067)

			VBody := gen7449
			gen7450 := Call(__e, ShenFunc(symdifference), VBody, VTerms)

			gen7451 := Call(__e, ShenFunc(symremove), MakeSymbol("Throwcontrol"), gen7450)

			Free := gen7451
			__e.TailApply(ShenFunc(symshen_4cc__help), Free, V1067)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.continuation_call"), gen7452)

		gen7453 := MakeNative(func(__e Evaluator) {
			V1070 := __e.Get(1)
			_ = V1070
			V1071 := __e.Get(2)
			_ = V1071
			__e.TailApply(ShenFunc(symshen_4remove_1h), V1070, V1071, Nil)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("remove"), gen7453)

		gen7465 := MakeNative(func(__e Evaluator) {
			V1078 := __e.Get(1)
			_ = V1078
			V1079 := __e.Get(2)
			_ = V1079
			V1080 := __e.Get(3)
			_ = V1080
			gen7464 := Call(__e, ShenFunc(sym_a), Nil, V1079)

			if True == gen7464 {
				__e.TailApply(ShenFunc(symreverse), V1080)

				return
			} else {
				gen7462 := Call(__e, ShenFunc(symcons_2), V1079)

				var gen7463 Obj
				if True == gen7462 {
					gen7460 := Call(__e, ShenFunc(symhd), V1079)

					gen7461 := Call(__e, ShenFunc(sym_a), gen7460, V1078)

					if True == gen7461 {
						gen7463 = True
					} else {
						gen7463 = False
					}

				} else {
					gen7463 = False
				}
				if True == gen7463 {
					gen7458 := Call(__e, ShenFunc(symhd), V1079)

					gen7459 := Call(__e, ShenFunc(symtl), V1079)

					__e.TailApply(ShenFunc(symshen_4remove_1h), gen7458, gen7459, V1080)

					return

				} else {
					gen7457 := Call(__e, ShenFunc(symcons_2), V1079)

					if True == gen7457 {
						gen7454 := Call(__e, ShenFunc(symtl), V1079)

						gen7455 := Call(__e, ShenFunc(symhd), V1079)

						gen7456 := Call(__e, ShenFunc(symcons), gen7455, V1080)

						__e.TailApply(ShenFunc(symshen_4remove_1h), V1078, gen7454, gen7456)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.remove-h"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove-h"), gen7465)

		gen7497 := MakeNative(func(__e Evaluator) {
			V1083 := __e.Get(1)
			_ = V1083
			V1084 := __e.Get(2)
			_ = V1084
			gen7495 := Call(__e, ShenFunc(sym_a), Nil, V1083)

			var gen7496 Obj
			if True == gen7495 {
				gen7494 := Call(__e, ShenFunc(sym_a), Nil, V1084)

				if True == gen7494 {
					gen7496 = True
				} else {
					gen7496 = False
				}

			} else {
				gen7496 = False
			}
			if True == gen7496 {
				gen7492 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.stack"), Nil)

				gen7493 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7492)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pop"), gen7493)

				return

			} else {
				gen7491 := Call(__e, ShenFunc(sym_a), Nil, V1084)

				if True == gen7491 {
					gen7481 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.stack"), Nil)

					gen7482 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7481)

					gen7483 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pop"), gen7482)

					gen7484 := Call(__e, ShenFunc(symcons), gen7483, Nil)

					gen7485 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7484)

					gen7486 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen7485)

					gen7487 := Call(__e, ShenFunc(symcons), V1083, gen7486)

					gen7488 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7487)

					gen7489 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen7488)

					gen7490 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7489)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.rename"), gen7490)

					return

				} else {
					gen7480 := Call(__e, ShenFunc(sym_a), Nil, V1083)

					if True == gen7480 {
						gen7477 := Call(__e, ShenFunc(symcons), V1084, Nil)

						gen7478 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.continuation"), gen7477)

						gen7479 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7478)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("call"), gen7479)

						return

					} else {
						gen7466 := Call(__e, ShenFunc(symcons), V1084, Nil)

						gen7467 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.continuation"), gen7466)

						gen7468 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7467)

						gen7469 := Call(__e, ShenFunc(symcons), MakeSymbol("call"), gen7468)

						gen7470 := Call(__e, ShenFunc(symcons), gen7469, Nil)

						gen7471 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7470)

						gen7472 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen7471)

						gen7473 := Call(__e, ShenFunc(symcons), V1083, gen7472)

						gen7474 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7473)

						gen7475 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen7474)

						gen7476 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7475)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.rename"), gen7476)

						return

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cc_help"), gen7497)

		gen7561 := MakeNative(func(__e Evaluator) {
			V1087 := __e.Get(1)
			_ = V1087
			V1088 := __e.Get(2)
			_ = V1088
			gen7559 := Call(__e, ShenFunc(symcons_2), V1087)

			var gen7560 Obj
			if True == gen7559 {
				gen7556 := Call(__e, ShenFunc(symhd), V1087)

				gen7557 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7556)

				var gen7558 Obj
				if True == gen7557 {
					gen7553 := Call(__e, ShenFunc(symtl), V1087)

					gen7554 := Call(__e, ShenFunc(symcons_2), gen7553)

					var gen7555 Obj
					if True == gen7554 {
						gen7549 := Call(__e, ShenFunc(symtl), V1087)

						gen7550 := Call(__e, ShenFunc(symhd), gen7549)

						gen7551 := Call(__e, ShenFunc(sym_a), Nil, gen7550)

						var gen7552 Obj
						if True == gen7551 {
							gen7545 := Call(__e, ShenFunc(symtl), V1087)

							gen7546 := Call(__e, ShenFunc(symtl), gen7545)

							gen7547 := Call(__e, ShenFunc(symcons_2), gen7546)

							var gen7548 Obj
							if True == gen7547 {
								gen7540 := Call(__e, ShenFunc(symtl), V1087)

								gen7541 := Call(__e, ShenFunc(symtl), gen7540)

								gen7542 := Call(__e, ShenFunc(symtl), gen7541)

								gen7543 := Call(__e, ShenFunc(sym_a), Nil, gen7542)

								var gen7544 Obj
								if True == gen7543 {
									gen7539 := Call(__e, ShenFunc(sym_a), Nil, V1088)

									if True == gen7539 {
										gen7544 = True
									} else {
										gen7544 = False
									}

								} else {
									gen7544 = False
								}
								if True == gen7544 {
									gen7548 = True
								} else {
									gen7548 = False
								}

							} else {
								gen7548 = False
							}
							if True == gen7548 {
								gen7552 = True
							} else {
								gen7552 = False
							}

						} else {
							gen7552 = False
						}
						if True == gen7552 {
							gen7555 = True
						} else {
							gen7555 = False
						}

					} else {
						gen7555 = False
					}
					if True == gen7555 {
						gen7558 = True
					} else {
						gen7558 = False
					}

				} else {
					gen7558 = False
				}
				if True == gen7558 {
					gen7560 = True
				} else {
					gen7560 = False
				}

			} else {
				gen7560 = False
			}
			if True == gen7560 {
				gen7537 := Call(__e, ShenFunc(symtl), V1087)

				gen7538 := Call(__e, ShenFunc(symtl), gen7537)

				__e.TailApply(ShenFunc(symhd), gen7538)

				return

			} else {
				gen7535 := Call(__e, ShenFunc(symcons_2), V1087)

				var gen7536 Obj
				if True == gen7535 {
					gen7532 := Call(__e, ShenFunc(symhd), V1087)

					gen7533 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7532)

					var gen7534 Obj
					if True == gen7533 {
						gen7529 := Call(__e, ShenFunc(symtl), V1087)

						gen7530 := Call(__e, ShenFunc(symcons_2), gen7529)

						var gen7531 Obj
						if True == gen7530 {
							gen7525 := Call(__e, ShenFunc(symtl), V1087)

							gen7526 := Call(__e, ShenFunc(symhd), gen7525)

							gen7527 := Call(__e, ShenFunc(symcons_2), gen7526)

							var gen7528 Obj
							if True == gen7527 {
								gen7521 := Call(__e, ShenFunc(symtl), V1087)

								gen7522 := Call(__e, ShenFunc(symtl), gen7521)

								gen7523 := Call(__e, ShenFunc(symcons_2), gen7522)

								var gen7524 Obj
								if True == gen7523 {
									gen7516 := Call(__e, ShenFunc(symtl), V1087)

									gen7517 := Call(__e, ShenFunc(symtl), gen7516)

									gen7518 := Call(__e, ShenFunc(symtl), gen7517)

									gen7519 := Call(__e, ShenFunc(sym_a), Nil, gen7518)

									var gen7520 Obj
									if True == gen7519 {
										gen7515 := Call(__e, ShenFunc(symcons_2), V1088)

										if True == gen7515 {
											gen7520 = True
										} else {
											gen7520 = False
										}

									} else {
										gen7520 = False
									}
									if True == gen7520 {
										gen7524 = True
									} else {
										gen7524 = False
									}

								} else {
									gen7524 = False
								}
								if True == gen7524 {
									gen7528 = True
								} else {
									gen7528 = False
								}

							} else {
								gen7528 = False
							}
							if True == gen7528 {
								gen7531 = True
							} else {
								gen7531 = False
							}

						} else {
							gen7531 = False
						}
						if True == gen7531 {
							gen7534 = True
						} else {
							gen7534 = False
						}

					} else {
						gen7534 = False
					}
					if True == gen7534 {
						gen7536 = True
					} else {
						gen7536 = False
					}

				} else {
					gen7536 = False
				}
				if True == gen7536 {
					gen7498 := Call(__e, ShenFunc(symtl), V1087)

					gen7499 := Call(__e, ShenFunc(symhd), gen7498)

					gen7500 := Call(__e, ShenFunc(symhd), gen7499)

					gen7501 := Call(__e, ShenFunc(symtl), V1087)

					gen7502 := Call(__e, ShenFunc(symhd), gen7501)

					gen7503 := Call(__e, ShenFunc(symtl), gen7502)

					gen7504 := Call(__e, ShenFunc(symtl), V1087)

					gen7505 := Call(__e, ShenFunc(symtl), gen7504)

					gen7506 := Call(__e, ShenFunc(symcons), gen7503, gen7505)

					gen7507 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7506)

					gen7508 := Call(__e, ShenFunc(symtl), V1088)

					gen7509 := Call(__e, ShenFunc(symshen_4make__mu__application), gen7507, gen7508)

					gen7510 := Call(__e, ShenFunc(symcons), gen7509, Nil)

					gen7511 := Call(__e, ShenFunc(symcons), gen7500, gen7510)

					gen7512 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7511)

					gen7513 := Call(__e, ShenFunc(symhd), V1088)

					gen7514 := Call(__e, ShenFunc(symcons), gen7513, Nil)

					__e.TailApply(ShenFunc(symcons), gen7512, gen7514)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.make_mu_application"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.make_mu_application"), gen7561)

		gen8150 := MakeNative(func(__e Evaluator) {
			V1097 := __e.Get(1)
			_ = V1097
			V1098 := __e.Get(2)
			_ = V1098
			gen8148 := Call(__e, ShenFunc(symcons_2), V1097)

			var gen8149 Obj
			if True == gen8148 {
				gen8145 := Call(__e, ShenFunc(symhd), V1097)

				gen8146 := Call(__e, ShenFunc(symcons_2), gen8145)

				var gen8147 Obj
				if True == gen8146 {
					gen8141 := Call(__e, ShenFunc(symhd), V1097)

					gen8142 := Call(__e, ShenFunc(symhd), gen8141)

					gen8143 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen8142)

					var gen8144 Obj
					if True == gen8143 {
						gen8137 := Call(__e, ShenFunc(symhd), V1097)

						gen8138 := Call(__e, ShenFunc(symtl), gen8137)

						gen8139 := Call(__e, ShenFunc(symcons_2), gen8138)

						var gen8140 Obj
						if True == gen8139 {
							gen8132 := Call(__e, ShenFunc(symhd), V1097)

							gen8133 := Call(__e, ShenFunc(symtl), gen8132)

							gen8134 := Call(__e, ShenFunc(symhd), gen8133)

							gen8135 := Call(__e, ShenFunc(symcons_2), gen8134)

							var gen8136 Obj
							if True == gen8135 {
								gen8126 := Call(__e, ShenFunc(symhd), V1097)

								gen8127 := Call(__e, ShenFunc(symtl), gen8126)

								gen8128 := Call(__e, ShenFunc(symhd), gen8127)

								gen8129 := Call(__e, ShenFunc(symhd), gen8128)

								gen8130 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen8129)

								var gen8131 Obj
								if True == gen8130 {
									gen8120 := Call(__e, ShenFunc(symhd), V1097)

									gen8121 := Call(__e, ShenFunc(symtl), gen8120)

									gen8122 := Call(__e, ShenFunc(symhd), gen8121)

									gen8123 := Call(__e, ShenFunc(symtl), gen8122)

									gen8124 := Call(__e, ShenFunc(symcons_2), gen8123)

									var gen8125 Obj
									if True == gen8124 {
										gen8113 := Call(__e, ShenFunc(symhd), V1097)

										gen8114 := Call(__e, ShenFunc(symtl), gen8113)

										gen8115 := Call(__e, ShenFunc(symhd), gen8114)

										gen8116 := Call(__e, ShenFunc(symtl), gen8115)

										gen8117 := Call(__e, ShenFunc(symtl), gen8116)

										gen8118 := Call(__e, ShenFunc(symcons_2), gen8117)

										var gen8119 Obj
										if True == gen8118 {
											gen8105 := Call(__e, ShenFunc(symhd), V1097)

											gen8106 := Call(__e, ShenFunc(symtl), gen8105)

											gen8107 := Call(__e, ShenFunc(symhd), gen8106)

											gen8108 := Call(__e, ShenFunc(symtl), gen8107)

											gen8109 := Call(__e, ShenFunc(symtl), gen8108)

											gen8110 := Call(__e, ShenFunc(symtl), gen8109)

											gen8111 := Call(__e, ShenFunc(sym_a), Nil, gen8110)

											var gen8112 Obj
											if True == gen8111 {
												gen8100 := Call(__e, ShenFunc(symhd), V1097)

												gen8101 := Call(__e, ShenFunc(symtl), gen8100)

												gen8102 := Call(__e, ShenFunc(symtl), gen8101)

												gen8103 := Call(__e, ShenFunc(symcons_2), gen8102)

												var gen8104 Obj
												if True == gen8103 {
													gen8094 := Call(__e, ShenFunc(symhd), V1097)

													gen8095 := Call(__e, ShenFunc(symtl), gen8094)

													gen8096 := Call(__e, ShenFunc(symtl), gen8095)

													gen8097 := Call(__e, ShenFunc(symtl), gen8096)

													gen8098 := Call(__e, ShenFunc(sym_a), Nil, gen8097)

													var gen8099 Obj
													if True == gen8098 {
														gen8091 := Call(__e, ShenFunc(symtl), V1097)

														gen8092 := Call(__e, ShenFunc(symcons_2), gen8091)

														var gen8093 Obj
														if True == gen8092 {
															gen8088 := Call(__e, ShenFunc(symtl), V1097)

															gen8089 := Call(__e, ShenFunc(symtl), gen8088)

															gen8090 := Call(__e, ShenFunc(sym_a), Nil, gen8089)

															if True == gen8090 {
																gen8093 = True
															} else {
																gen8093 = False
															}

														} else {
															gen8093 = False
														}
														if True == gen8093 {
															gen8099 = True
														} else {
															gen8099 = False
														}

													} else {
														gen8099 = False
													}
													if True == gen8099 {
														gen8104 = True
													} else {
														gen8104 = False
													}

												} else {
													gen8104 = False
												}
												if True == gen8104 {
													gen8112 = True
												} else {
													gen8112 = False
												}

											} else {
												gen8112 = False
											}
											if True == gen8112 {
												gen8119 = True
											} else {
												gen8119 = False
											}

										} else {
											gen8119 = False
										}
										if True == gen8119 {
											gen8125 = True
										} else {
											gen8125 = False
										}

									} else {
										gen8125 = False
									}
									if True == gen8125 {
										gen8131 = True
									} else {
										gen8131 = False
									}

								} else {
									gen8131 = False
								}
								if True == gen8131 {
									gen8136 = True
								} else {
									gen8136 = False
								}

							} else {
								gen8136 = False
							}
							if True == gen8136 {
								gen8140 = True
							} else {
								gen8140 = False
							}

						} else {
							gen8140 = False
						}
						if True == gen8140 {
							gen8144 = True
						} else {
							gen8144 = False
						}

					} else {
						gen8144 = False
					}
					if True == gen8144 {
						gen8147 = True
					} else {
						gen8147 = False
					}

				} else {
					gen8147 = False
				}
				if True == gen8147 {
					gen8149 = True
				} else {
					gen8149 = False
				}

			} else {
				gen8149 = False
			}
			if True == gen8149 {
				gen8070 := Call(__e, ShenFunc(symhd), V1097)

				gen8071 := Call(__e, ShenFunc(symtl), gen8070)

				gen8072 := Call(__e, ShenFunc(symhd), gen8071)

				gen8073 := Call(__e, ShenFunc(symtl), gen8072)

				gen8074 := Call(__e, ShenFunc(symhd), gen8073)

				gen8075 := Call(__e, ShenFunc(symhd), V1097)

				gen8076 := Call(__e, ShenFunc(symtl), gen8075)

				gen8077 := Call(__e, ShenFunc(symtl), gen8076)

				gen8078 := Call(__e, ShenFunc(symcons), gen8074, gen8077)

				gen8079 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen8078)

				gen8080 := Call(__e, ShenFunc(symtl), V1097)

				gen8081 := Call(__e, ShenFunc(symcons), gen8079, gen8080)

				gen8082 := Call(__e, ShenFunc(symhd), V1097)

				gen8083 := Call(__e, ShenFunc(symtl), gen8082)

				gen8084 := Call(__e, ShenFunc(symhd), gen8083)

				gen8085 := Call(__e, ShenFunc(symtl), gen8084)

				gen8086 := Call(__e, ShenFunc(symtl), gen8085)

				gen8087 := Call(__e, ShenFunc(symhd), gen8086)

				__e.TailApply(ShenFunc(symshen_4mu__reduction), gen8081, gen8087)

				return

			} else {
				gen8068 := Call(__e, ShenFunc(symcons_2), V1097)

				var gen8069 Obj
				if True == gen8068 {
					gen8065 := Call(__e, ShenFunc(symhd), V1097)

					gen8066 := Call(__e, ShenFunc(symcons_2), gen8065)

					var gen8067 Obj
					if True == gen8066 {
						gen8061 := Call(__e, ShenFunc(symhd), V1097)

						gen8062 := Call(__e, ShenFunc(symhd), gen8061)

						gen8063 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen8062)

						var gen8064 Obj
						if True == gen8063 {
							gen8057 := Call(__e, ShenFunc(symhd), V1097)

							gen8058 := Call(__e, ShenFunc(symtl), gen8057)

							gen8059 := Call(__e, ShenFunc(symcons_2), gen8058)

							var gen8060 Obj
							if True == gen8059 {
								gen8052 := Call(__e, ShenFunc(symhd), V1097)

								gen8053 := Call(__e, ShenFunc(symtl), gen8052)

								gen8054 := Call(__e, ShenFunc(symtl), gen8053)

								gen8055 := Call(__e, ShenFunc(symcons_2), gen8054)

								var gen8056 Obj
								if True == gen8055 {
									gen8046 := Call(__e, ShenFunc(symhd), V1097)

									gen8047 := Call(__e, ShenFunc(symtl), gen8046)

									gen8048 := Call(__e, ShenFunc(symtl), gen8047)

									gen8049 := Call(__e, ShenFunc(symtl), gen8048)

									gen8050 := Call(__e, ShenFunc(sym_a), Nil, gen8049)

									var gen8051 Obj
									if True == gen8050 {
										gen8043 := Call(__e, ShenFunc(symtl), V1097)

										gen8044 := Call(__e, ShenFunc(symcons_2), gen8043)

										var gen8045 Obj
										if True == gen8044 {
											gen8039 := Call(__e, ShenFunc(symtl), V1097)

											gen8040 := Call(__e, ShenFunc(symtl), gen8039)

											gen8041 := Call(__e, ShenFunc(sym_a), Nil, gen8040)

											var gen8042 Obj
											if True == gen8041 {
												gen8035 := Call(__e, ShenFunc(symhd), V1097)

												gen8036 := Call(__e, ShenFunc(symtl), gen8035)

												gen8037 := Call(__e, ShenFunc(symhd), gen8036)

												gen8038 := Call(__e, ShenFunc(sym_a), MakeSymbol("_"), gen8037)

												if True == gen8038 {
													gen8042 = True
												} else {
													gen8042 = False
												}

											} else {
												gen8042 = False
											}
											if True == gen8042 {
												gen8045 = True
											} else {
												gen8045 = False
											}

										} else {
											gen8045 = False
										}
										if True == gen8045 {
											gen8051 = True
										} else {
											gen8051 = False
										}

									} else {
										gen8051 = False
									}
									if True == gen8051 {
										gen8056 = True
									} else {
										gen8056 = False
									}

								} else {
									gen8056 = False
								}
								if True == gen8056 {
									gen8060 = True
								} else {
									gen8060 = False
								}

							} else {
								gen8060 = False
							}
							if True == gen8060 {
								gen8064 = True
							} else {
								gen8064 = False
							}

						} else {
							gen8064 = False
						}
						if True == gen8064 {
							gen8067 = True
						} else {
							gen8067 = False
						}

					} else {
						gen8067 = False
					}
					if True == gen8067 {
						gen8069 = True
					} else {
						gen8069 = False
					}

				} else {
					gen8069 = False
				}
				if True == gen8069 {
					gen8031 := Call(__e, ShenFunc(symhd), V1097)

					gen8032 := Call(__e, ShenFunc(symtl), gen8031)

					gen8033 := Call(__e, ShenFunc(symtl), gen8032)

					gen8034 := Call(__e, ShenFunc(symhd), gen8033)

					__e.TailApply(ShenFunc(symshen_4mu__reduction), gen8034, V1098)

					return

				} else {
					gen8029 := Call(__e, ShenFunc(symcons_2), V1097)

					var gen8030 Obj
					if True == gen8029 {
						gen8026 := Call(__e, ShenFunc(symhd), V1097)

						gen8027 := Call(__e, ShenFunc(symcons_2), gen8026)

						var gen8028 Obj
						if True == gen8027 {
							gen8022 := Call(__e, ShenFunc(symhd), V1097)

							gen8023 := Call(__e, ShenFunc(symhd), gen8022)

							gen8024 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen8023)

							var gen8025 Obj
							if True == gen8024 {
								gen8018 := Call(__e, ShenFunc(symhd), V1097)

								gen8019 := Call(__e, ShenFunc(symtl), gen8018)

								gen8020 := Call(__e, ShenFunc(symcons_2), gen8019)

								var gen8021 Obj
								if True == gen8020 {
									gen8013 := Call(__e, ShenFunc(symhd), V1097)

									gen8014 := Call(__e, ShenFunc(symtl), gen8013)

									gen8015 := Call(__e, ShenFunc(symtl), gen8014)

									gen8016 := Call(__e, ShenFunc(symcons_2), gen8015)

									var gen8017 Obj
									if True == gen8016 {
										gen8007 := Call(__e, ShenFunc(symhd), V1097)

										gen8008 := Call(__e, ShenFunc(symtl), gen8007)

										gen8009 := Call(__e, ShenFunc(symtl), gen8008)

										gen8010 := Call(__e, ShenFunc(symtl), gen8009)

										gen8011 := Call(__e, ShenFunc(sym_a), Nil, gen8010)

										var gen8012 Obj
										if True == gen8011 {
											gen8004 := Call(__e, ShenFunc(symtl), V1097)

											gen8005 := Call(__e, ShenFunc(symcons_2), gen8004)

											var gen8006 Obj
											if True == gen8005 {
												gen8000 := Call(__e, ShenFunc(symtl), V1097)

												gen8001 := Call(__e, ShenFunc(symtl), gen8000)

												gen8002 := Call(__e, ShenFunc(sym_a), Nil, gen8001)

												var gen8003 Obj
												if True == gen8002 {
													gen7994 := Call(__e, ShenFunc(symhd), V1097)

													gen7995 := Call(__e, ShenFunc(symtl), gen7994)

													gen7996 := Call(__e, ShenFunc(symhd), gen7995)

													gen7997 := Call(__e, ShenFunc(symtl), V1097)

													gen7998 := Call(__e, ShenFunc(symhd), gen7997)

													gen7999 := Call(__e, ShenFunc(symshen_4ephemeral__variable_2), gen7996, gen7998)

													if True == gen7999 {
														gen8003 = True
													} else {
														gen8003 = False
													}

												} else {
													gen8003 = False
												}
												if True == gen8003 {
													gen8006 = True
												} else {
													gen8006 = False
												}

											} else {
												gen8006 = False
											}
											if True == gen8006 {
												gen8012 = True
											} else {
												gen8012 = False
											}

										} else {
											gen8012 = False
										}
										if True == gen8012 {
											gen8017 = True
										} else {
											gen8017 = False
										}

									} else {
										gen8017 = False
									}
									if True == gen8017 {
										gen8021 = True
									} else {
										gen8021 = False
									}

								} else {
									gen8021 = False
								}
								if True == gen8021 {
									gen8025 = True
								} else {
									gen8025 = False
								}

							} else {
								gen8025 = False
							}
							if True == gen8025 {
								gen8028 = True
							} else {
								gen8028 = False
							}

						} else {
							gen8028 = False
						}
						if True == gen8028 {
							gen8030 = True
						} else {
							gen8030 = False
						}

					} else {
						gen8030 = False
					}
					if True == gen8030 {
						gen7984 := Call(__e, ShenFunc(symtl), V1097)

						gen7985 := Call(__e, ShenFunc(symhd), gen7984)

						gen7986 := Call(__e, ShenFunc(symhd), V1097)

						gen7987 := Call(__e, ShenFunc(symtl), gen7986)

						gen7988 := Call(__e, ShenFunc(symhd), gen7987)

						gen7989 := Call(__e, ShenFunc(symhd), V1097)

						gen7990 := Call(__e, ShenFunc(symtl), gen7989)

						gen7991 := Call(__e, ShenFunc(symtl), gen7990)

						gen7992 := Call(__e, ShenFunc(symhd), gen7991)

						gen7993 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7992, V1098)

						__e.TailApply(ShenFunc(symsubst), gen7985, gen7988, gen7993)

						return

					} else {
						gen7982 := Call(__e, ShenFunc(symcons_2), V1097)

						var gen7983 Obj
						if True == gen7982 {
							gen7979 := Call(__e, ShenFunc(symhd), V1097)

							gen7980 := Call(__e, ShenFunc(symcons_2), gen7979)

							var gen7981 Obj
							if True == gen7980 {
								gen7975 := Call(__e, ShenFunc(symhd), V1097)

								gen7976 := Call(__e, ShenFunc(symhd), gen7975)

								gen7977 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7976)

								var gen7978 Obj
								if True == gen7977 {
									gen7971 := Call(__e, ShenFunc(symhd), V1097)

									gen7972 := Call(__e, ShenFunc(symtl), gen7971)

									gen7973 := Call(__e, ShenFunc(symcons_2), gen7972)

									var gen7974 Obj
									if True == gen7973 {
										gen7966 := Call(__e, ShenFunc(symhd), V1097)

										gen7967 := Call(__e, ShenFunc(symtl), gen7966)

										gen7968 := Call(__e, ShenFunc(symtl), gen7967)

										gen7969 := Call(__e, ShenFunc(symcons_2), gen7968)

										var gen7970 Obj
										if True == gen7969 {
											gen7960 := Call(__e, ShenFunc(symhd), V1097)

											gen7961 := Call(__e, ShenFunc(symtl), gen7960)

											gen7962 := Call(__e, ShenFunc(symtl), gen7961)

											gen7963 := Call(__e, ShenFunc(symtl), gen7962)

											gen7964 := Call(__e, ShenFunc(sym_a), Nil, gen7963)

											var gen7965 Obj
											if True == gen7964 {
												gen7957 := Call(__e, ShenFunc(symtl), V1097)

												gen7958 := Call(__e, ShenFunc(symcons_2), gen7957)

												var gen7959 Obj
												if True == gen7958 {
													gen7953 := Call(__e, ShenFunc(symtl), V1097)

													gen7954 := Call(__e, ShenFunc(symtl), gen7953)

													gen7955 := Call(__e, ShenFunc(sym_a), Nil, gen7954)

													var gen7956 Obj
													if True == gen7955 {
														gen7949 := Call(__e, ShenFunc(symhd), V1097)

														gen7950 := Call(__e, ShenFunc(symtl), gen7949)

														gen7951 := Call(__e, ShenFunc(symhd), gen7950)

														gen7952 := Call(__e, ShenFunc(symvariable_2), gen7951)

														if True == gen7952 {
															gen7956 = True
														} else {
															gen7956 = False
														}

													} else {
														gen7956 = False
													}
													if True == gen7956 {
														gen7959 = True
													} else {
														gen7959 = False
													}

												} else {
													gen7959 = False
												}
												if True == gen7959 {
													gen7965 = True
												} else {
													gen7965 = False
												}

											} else {
												gen7965 = False
											}
											if True == gen7965 {
												gen7970 = True
											} else {
												gen7970 = False
											}

										} else {
											gen7970 = False
										}
										if True == gen7970 {
											gen7974 = True
										} else {
											gen7974 = False
										}

									} else {
										gen7974 = False
									}
									if True == gen7974 {
										gen7978 = True
									} else {
										gen7978 = False
									}

								} else {
									gen7978 = False
								}
								if True == gen7978 {
									gen7981 = True
								} else {
									gen7981 = False
								}

							} else {
								gen7981 = False
							}
							if True == gen7981 {
								gen7983 = True
							} else {
								gen7983 = False
							}

						} else {
							gen7983 = False
						}
						if True == gen7983 {
							gen7934 := Call(__e, ShenFunc(symhd), V1097)

							gen7935 := Call(__e, ShenFunc(symtl), gen7934)

							gen7936 := Call(__e, ShenFunc(symhd), gen7935)

							gen7937 := Call(__e, ShenFunc(symtl), V1097)

							gen7938 := Call(__e, ShenFunc(symhd), gen7937)

							gen7939 := Call(__e, ShenFunc(symhd), V1097)

							gen7940 := Call(__e, ShenFunc(symtl), gen7939)

							gen7941 := Call(__e, ShenFunc(symtl), gen7940)

							gen7942 := Call(__e, ShenFunc(symhd), gen7941)

							gen7943 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7942, V1098)

							gen7944 := Call(__e, ShenFunc(symcons), gen7943, Nil)

							gen7945 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7944)

							gen7946 := Call(__e, ShenFunc(symcons), gen7938, gen7945)

							gen7947 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen7946)

							gen7948 := Call(__e, ShenFunc(symcons), gen7936, gen7947)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7948)

							return

						} else {
							gen7932 := Call(__e, ShenFunc(symcons_2), V1097)

							var gen7933 Obj
							if True == gen7932 {
								gen7929 := Call(__e, ShenFunc(symhd), V1097)

								gen7930 := Call(__e, ShenFunc(symcons_2), gen7929)

								var gen7931 Obj
								if True == gen7930 {
									gen7925 := Call(__e, ShenFunc(symhd), V1097)

									gen7926 := Call(__e, ShenFunc(symhd), gen7925)

									gen7927 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7926)

									var gen7928 Obj
									if True == gen7927 {
										gen7921 := Call(__e, ShenFunc(symhd), V1097)

										gen7922 := Call(__e, ShenFunc(symtl), gen7921)

										gen7923 := Call(__e, ShenFunc(symcons_2), gen7922)

										var gen7924 Obj
										if True == gen7923 {
											gen7916 := Call(__e, ShenFunc(symhd), V1097)

											gen7917 := Call(__e, ShenFunc(symtl), gen7916)

											gen7918 := Call(__e, ShenFunc(symtl), gen7917)

											gen7919 := Call(__e, ShenFunc(symcons_2), gen7918)

											var gen7920 Obj
											if True == gen7919 {
												gen7910 := Call(__e, ShenFunc(symhd), V1097)

												gen7911 := Call(__e, ShenFunc(symtl), gen7910)

												gen7912 := Call(__e, ShenFunc(symtl), gen7911)

												gen7913 := Call(__e, ShenFunc(symtl), gen7912)

												gen7914 := Call(__e, ShenFunc(sym_a), Nil, gen7913)

												var gen7915 Obj
												if True == gen7914 {
													gen7907 := Call(__e, ShenFunc(symtl), V1097)

													gen7908 := Call(__e, ShenFunc(symcons_2), gen7907)

													var gen7909 Obj
													if True == gen7908 {
														gen7903 := Call(__e, ShenFunc(symtl), V1097)

														gen7904 := Call(__e, ShenFunc(symtl), gen7903)

														gen7905 := Call(__e, ShenFunc(sym_a), Nil, gen7904)

														var gen7906 Obj
														if True == gen7905 {
															gen7901 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1098)

															var gen7902 Obj
															if True == gen7901 {
																gen7897 := Call(__e, ShenFunc(symhd), V1097)

																gen7898 := Call(__e, ShenFunc(symtl), gen7897)

																gen7899 := Call(__e, ShenFunc(symhd), gen7898)

																gen7900 := Call(__e, ShenFunc(symshen_4prolog__constant_2), gen7899)

																if True == gen7900 {
																	gen7902 = True
																} else {
																	gen7902 = False
																}

															} else {
																gen7902 = False
															}
															if True == gen7902 {
																gen7906 = True
															} else {
																gen7906 = False
															}

														} else {
															gen7906 = False
														}
														if True == gen7906 {
															gen7909 = True
														} else {
															gen7909 = False
														}

													} else {
														gen7909 = False
													}
													if True == gen7909 {
														gen7915 = True
													} else {
														gen7915 = False
													}

												} else {
													gen7915 = False
												}
												if True == gen7915 {
													gen7920 = True
												} else {
													gen7920 = False
												}

											} else {
												gen7920 = False
											}
											if True == gen7920 {
												gen7924 = True
											} else {
												gen7924 = False
											}

										} else {
											gen7924 = False
										}
										if True == gen7924 {
											gen7928 = True
										} else {
											gen7928 = False
										}

									} else {
										gen7928 = False
									}
									if True == gen7928 {
										gen7931 = True
									} else {
										gen7931 = False
									}

								} else {
									gen7931 = False
								}
								if True == gen7931 {
									gen7933 = True
								} else {
									gen7933 = False
								}

							} else {
								gen7933 = False
							}
							if True == gen7933 {
								gen7867 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

								Z := gen7867
								gen7868 := Call(__e, ShenFunc(symtl), V1097)

								gen7869 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen7868)

								gen7870 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7869)

								gen7871 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen7870)

								gen7872 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7871)

								gen7873 := Call(__e, ShenFunc(symhd), V1097)

								gen7874 := Call(__e, ShenFunc(symtl), gen7873)

								gen7875 := Call(__e, ShenFunc(symhd), gen7874)

								gen7876 := Call(__e, ShenFunc(symcons), gen7875, Nil)

								gen7877 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen7876)

								gen7878 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen7877)

								gen7879 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7878)

								gen7880 := Call(__e, ShenFunc(symcons), Z, gen7879)

								gen7881 := Call(__e, ShenFunc(symhd), V1097)

								gen7882 := Call(__e, ShenFunc(symtl), gen7881)

								gen7883 := Call(__e, ShenFunc(symtl), gen7882)

								gen7884 := Call(__e, ShenFunc(symhd), gen7883)

								gen7885 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7884, MakeSymbol("-"))

								gen7886 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

								gen7887 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7886)

								gen7888 := Call(__e, ShenFunc(symcons), gen7885, gen7887)

								gen7889 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7888)

								gen7890 := Call(__e, ShenFunc(symcons), gen7880, gen7889)

								gen7891 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7890)

								gen7892 := Call(__e, ShenFunc(symcons), gen7891, Nil)

								gen7893 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7892)

								gen7894 := Call(__e, ShenFunc(symcons), gen7872, gen7893)

								gen7895 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen7894)

								gen7896 := Call(__e, ShenFunc(symcons), Z, gen7895)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7896)

								return

							} else {
								gen7865 := Call(__e, ShenFunc(symcons_2), V1097)

								var gen7866 Obj
								if True == gen7865 {
									gen7862 := Call(__e, ShenFunc(symhd), V1097)

									gen7863 := Call(__e, ShenFunc(symcons_2), gen7862)

									var gen7864 Obj
									if True == gen7863 {
										gen7858 := Call(__e, ShenFunc(symhd), V1097)

										gen7859 := Call(__e, ShenFunc(symhd), gen7858)

										gen7860 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7859)

										var gen7861 Obj
										if True == gen7860 {
											gen7854 := Call(__e, ShenFunc(symhd), V1097)

											gen7855 := Call(__e, ShenFunc(symtl), gen7854)

											gen7856 := Call(__e, ShenFunc(symcons_2), gen7855)

											var gen7857 Obj
											if True == gen7856 {
												gen7849 := Call(__e, ShenFunc(symhd), V1097)

												gen7850 := Call(__e, ShenFunc(symtl), gen7849)

												gen7851 := Call(__e, ShenFunc(symtl), gen7850)

												gen7852 := Call(__e, ShenFunc(symcons_2), gen7851)

												var gen7853 Obj
												if True == gen7852 {
													gen7843 := Call(__e, ShenFunc(symhd), V1097)

													gen7844 := Call(__e, ShenFunc(symtl), gen7843)

													gen7845 := Call(__e, ShenFunc(symtl), gen7844)

													gen7846 := Call(__e, ShenFunc(symtl), gen7845)

													gen7847 := Call(__e, ShenFunc(sym_a), Nil, gen7846)

													var gen7848 Obj
													if True == gen7847 {
														gen7840 := Call(__e, ShenFunc(symtl), V1097)

														gen7841 := Call(__e, ShenFunc(symcons_2), gen7840)

														var gen7842 Obj
														if True == gen7841 {
															gen7836 := Call(__e, ShenFunc(symtl), V1097)

															gen7837 := Call(__e, ShenFunc(symtl), gen7836)

															gen7838 := Call(__e, ShenFunc(sym_a), Nil, gen7837)

															var gen7839 Obj
															if True == gen7838 {
																gen7834 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1098)

																var gen7835 Obj
																if True == gen7834 {
																	gen7830 := Call(__e, ShenFunc(symhd), V1097)

																	gen7831 := Call(__e, ShenFunc(symtl), gen7830)

																	gen7832 := Call(__e, ShenFunc(symhd), gen7831)

																	gen7833 := Call(__e, ShenFunc(symshen_4prolog__constant_2), gen7832)

																	if True == gen7833 {
																		gen7835 = True
																	} else {
																		gen7835 = False
																	}

																} else {
																	gen7835 = False
																}
																if True == gen7835 {
																	gen7839 = True
																} else {
																	gen7839 = False
																}

															} else {
																gen7839 = False
															}
															if True == gen7839 {
																gen7842 = True
															} else {
																gen7842 = False
															}

														} else {
															gen7842 = False
														}
														if True == gen7842 {
															gen7848 = True
														} else {
															gen7848 = False
														}

													} else {
														gen7848 = False
													}
													if True == gen7848 {
														gen7853 = True
													} else {
														gen7853 = False
													}

												} else {
													gen7853 = False
												}
												if True == gen7853 {
													gen7857 = True
												} else {
													gen7857 = False
												}

											} else {
												gen7857 = False
											}
											if True == gen7857 {
												gen7861 = True
											} else {
												gen7861 = False
											}

										} else {
											gen7861 = False
										}
										if True == gen7861 {
											gen7864 = True
										} else {
											gen7864 = False
										}

									} else {
										gen7864 = False
									}
									if True == gen7864 {
										gen7866 = True
									} else {
										gen7866 = False
									}

								} else {
									gen7866 = False
								}
								if True == gen7866 {
									gen7776 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

									Z := gen7776
									gen7777 := Call(__e, ShenFunc(symtl), V1097)

									gen7778 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen7777)

									gen7779 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7778)

									gen7780 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen7779)

									gen7781 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7780)

									gen7782 := Call(__e, ShenFunc(symhd), V1097)

									gen7783 := Call(__e, ShenFunc(symtl), gen7782)

									gen7784 := Call(__e, ShenFunc(symhd), gen7783)

									gen7785 := Call(__e, ShenFunc(symcons), gen7784, Nil)

									gen7786 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen7785)

									gen7787 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen7786)

									gen7788 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7787)

									gen7789 := Call(__e, ShenFunc(symcons), Z, gen7788)

									gen7790 := Call(__e, ShenFunc(symhd), V1097)

									gen7791 := Call(__e, ShenFunc(symtl), gen7790)

									gen7792 := Call(__e, ShenFunc(symtl), gen7791)

									gen7793 := Call(__e, ShenFunc(symhd), gen7792)

									gen7794 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7793, MakeSymbol("+"))

									gen7795 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variable"), Nil)

									gen7796 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen7795)

									gen7797 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7796)

									gen7798 := Call(__e, ShenFunc(symcons), Z, gen7797)

									gen7799 := Call(__e, ShenFunc(symhd), V1097)

									gen7800 := Call(__e, ShenFunc(symtl), gen7799)

									gen7801 := Call(__e, ShenFunc(symhd), gen7800)

									gen7802 := Call(__e, ShenFunc(symhd), V1097)

									gen7803 := Call(__e, ShenFunc(symtl), gen7802)

									gen7804 := Call(__e, ShenFunc(symtl), gen7803)

									gen7805 := Call(__e, ShenFunc(symhd), gen7804)

									gen7806 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7805, MakeSymbol("+"))

									gen7807 := Call(__e, ShenFunc(symcons), gen7806, Nil)

									gen7808 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7807)

									gen7809 := Call(__e, ShenFunc(symcons), gen7801, gen7808)

									gen7810 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen7809)

									gen7811 := Call(__e, ShenFunc(symcons), Z, gen7810)

									gen7812 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen7811)

									gen7813 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

									gen7814 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7813)

									gen7815 := Call(__e, ShenFunc(symcons), gen7812, gen7814)

									gen7816 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7815)

									gen7817 := Call(__e, ShenFunc(symcons), gen7798, gen7816)

									gen7818 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7817)

									gen7819 := Call(__e, ShenFunc(symcons), gen7818, Nil)

									gen7820 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7819)

									gen7821 := Call(__e, ShenFunc(symcons), gen7794, gen7820)

									gen7822 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7821)

									gen7823 := Call(__e, ShenFunc(symcons), gen7789, gen7822)

									gen7824 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7823)

									gen7825 := Call(__e, ShenFunc(symcons), gen7824, Nil)

									gen7826 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7825)

									gen7827 := Call(__e, ShenFunc(symcons), gen7781, gen7826)

									gen7828 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen7827)

									gen7829 := Call(__e, ShenFunc(symcons), Z, gen7828)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7829)

									return

								} else {
									gen7774 := Call(__e, ShenFunc(symcons_2), V1097)

									var gen7775 Obj
									if True == gen7774 {
										gen7771 := Call(__e, ShenFunc(symhd), V1097)

										gen7772 := Call(__e, ShenFunc(symcons_2), gen7771)

										var gen7773 Obj
										if True == gen7772 {
											gen7767 := Call(__e, ShenFunc(symhd), V1097)

											gen7768 := Call(__e, ShenFunc(symhd), gen7767)

											gen7769 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7768)

											var gen7770 Obj
											if True == gen7769 {
												gen7763 := Call(__e, ShenFunc(symhd), V1097)

												gen7764 := Call(__e, ShenFunc(symtl), gen7763)

												gen7765 := Call(__e, ShenFunc(symcons_2), gen7764)

												var gen7766 Obj
												if True == gen7765 {
													gen7758 := Call(__e, ShenFunc(symhd), V1097)

													gen7759 := Call(__e, ShenFunc(symtl), gen7758)

													gen7760 := Call(__e, ShenFunc(symhd), gen7759)

													gen7761 := Call(__e, ShenFunc(symcons_2), gen7760)

													var gen7762 Obj
													if True == gen7761 {
														gen7753 := Call(__e, ShenFunc(symhd), V1097)

														gen7754 := Call(__e, ShenFunc(symtl), gen7753)

														gen7755 := Call(__e, ShenFunc(symtl), gen7754)

														gen7756 := Call(__e, ShenFunc(symcons_2), gen7755)

														var gen7757 Obj
														if True == gen7756 {
															gen7747 := Call(__e, ShenFunc(symhd), V1097)

															gen7748 := Call(__e, ShenFunc(symtl), gen7747)

															gen7749 := Call(__e, ShenFunc(symtl), gen7748)

															gen7750 := Call(__e, ShenFunc(symtl), gen7749)

															gen7751 := Call(__e, ShenFunc(sym_a), Nil, gen7750)

															var gen7752 Obj
															if True == gen7751 {
																gen7744 := Call(__e, ShenFunc(symtl), V1097)

																gen7745 := Call(__e, ShenFunc(symcons_2), gen7744)

																var gen7746 Obj
																if True == gen7745 {
																	gen7740 := Call(__e, ShenFunc(symtl), V1097)

																	gen7741 := Call(__e, ShenFunc(symtl), gen7740)

																	gen7742 := Call(__e, ShenFunc(sym_a), Nil, gen7741)

																	var gen7743 Obj
																	if True == gen7742 {
																		gen7739 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1098)

																		if True == gen7739 {
																			gen7743 = True
																		} else {
																			gen7743 = False
																		}

																	} else {
																		gen7743 = False
																	}
																	if True == gen7743 {
																		gen7746 = True
																	} else {
																		gen7746 = False
																	}

																} else {
																	gen7746 = False
																}
																if True == gen7746 {
																	gen7752 = True
																} else {
																	gen7752 = False
																}

															} else {
																gen7752 = False
															}
															if True == gen7752 {
																gen7757 = True
															} else {
																gen7757 = False
															}

														} else {
															gen7757 = False
														}
														if True == gen7757 {
															gen7762 = True
														} else {
															gen7762 = False
														}

													} else {
														gen7762 = False
													}
													if True == gen7762 {
														gen7766 = True
													} else {
														gen7766 = False
													}

												} else {
													gen7766 = False
												}
												if True == gen7766 {
													gen7770 = True
												} else {
													gen7770 = False
												}

											} else {
												gen7770 = False
											}
											if True == gen7770 {
												gen7773 = True
											} else {
												gen7773 = False
											}

										} else {
											gen7773 = False
										}
										if True == gen7773 {
											gen7775 = True
										} else {
											gen7775 = False
										}

									} else {
										gen7775 = False
									}
									if True == gen7775 {
										gen7688 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

										Z := gen7688
										gen7689 := Call(__e, ShenFunc(symtl), V1097)

										gen7690 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen7689)

										gen7691 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7690)

										gen7692 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen7691)

										gen7693 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7692)

										gen7694 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), Nil)

										gen7695 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.non-empty"), gen7694)

										gen7696 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen7695)

										gen7697 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7696)

										gen7698 := Call(__e, ShenFunc(symcons), Z, gen7697)

										gen7699 := Call(__e, ShenFunc(symhd), V1097)

										gen7700 := Call(__e, ShenFunc(symtl), gen7699)

										gen7701 := Call(__e, ShenFunc(symhd), gen7700)

										gen7702 := Call(__e, ShenFunc(symhd), gen7701)

										gen7703 := Call(__e, ShenFunc(symhd), V1097)

										gen7704 := Call(__e, ShenFunc(symtl), gen7703)

										gen7705 := Call(__e, ShenFunc(symhd), gen7704)

										gen7706 := Call(__e, ShenFunc(symtl), gen7705)

										gen7707 := Call(__e, ShenFunc(symhd), V1097)

										gen7708 := Call(__e, ShenFunc(symtl), gen7707)

										gen7709 := Call(__e, ShenFunc(symtl), gen7708)

										gen7710 := Call(__e, ShenFunc(symcons), gen7706, gen7709)

										gen7711 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7710)

										gen7712 := Call(__e, ShenFunc(symcons), Z, Nil)

										gen7713 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7712)

										gen7714 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen7713)

										gen7715 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7714)

										gen7716 := Call(__e, ShenFunc(symcons), gen7715, Nil)

										gen7717 := Call(__e, ShenFunc(symcons), gen7711, gen7716)

										gen7718 := Call(__e, ShenFunc(symcons), gen7717, Nil)

										gen7719 := Call(__e, ShenFunc(symcons), gen7702, gen7718)

										gen7720 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7719)

										gen7721 := Call(__e, ShenFunc(symcons), Z, Nil)

										gen7722 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7721)

										gen7723 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen7722)

										gen7724 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7723)

										gen7725 := Call(__e, ShenFunc(symcons), gen7724, Nil)

										gen7726 := Call(__e, ShenFunc(symcons), gen7720, gen7725)

										gen7727 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7726, MakeSymbol("-"))

										gen7728 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

										gen7729 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7728)

										gen7730 := Call(__e, ShenFunc(symcons), gen7727, gen7729)

										gen7731 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7730)

										gen7732 := Call(__e, ShenFunc(symcons), gen7698, gen7731)

										gen7733 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7732)

										gen7734 := Call(__e, ShenFunc(symcons), gen7733, Nil)

										gen7735 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7734)

										gen7736 := Call(__e, ShenFunc(symcons), gen7693, gen7735)

										gen7737 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen7736)

										gen7738 := Call(__e, ShenFunc(symcons), Z, gen7737)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7738)

										return

									} else {
										gen7686 := Call(__e, ShenFunc(symcons_2), V1097)

										var gen7687 Obj
										if True == gen7686 {
											gen7683 := Call(__e, ShenFunc(symhd), V1097)

											gen7684 := Call(__e, ShenFunc(symcons_2), gen7683)

											var gen7685 Obj
											if True == gen7684 {
												gen7679 := Call(__e, ShenFunc(symhd), V1097)

												gen7680 := Call(__e, ShenFunc(symhd), gen7679)

												gen7681 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen7680)

												var gen7682 Obj
												if True == gen7681 {
													gen7675 := Call(__e, ShenFunc(symhd), V1097)

													gen7676 := Call(__e, ShenFunc(symtl), gen7675)

													gen7677 := Call(__e, ShenFunc(symcons_2), gen7676)

													var gen7678 Obj
													if True == gen7677 {
														gen7670 := Call(__e, ShenFunc(symhd), V1097)

														gen7671 := Call(__e, ShenFunc(symtl), gen7670)

														gen7672 := Call(__e, ShenFunc(symhd), gen7671)

														gen7673 := Call(__e, ShenFunc(symcons_2), gen7672)

														var gen7674 Obj
														if True == gen7673 {
															gen7665 := Call(__e, ShenFunc(symhd), V1097)

															gen7666 := Call(__e, ShenFunc(symtl), gen7665)

															gen7667 := Call(__e, ShenFunc(symtl), gen7666)

															gen7668 := Call(__e, ShenFunc(symcons_2), gen7667)

															var gen7669 Obj
															if True == gen7668 {
																gen7659 := Call(__e, ShenFunc(symhd), V1097)

																gen7660 := Call(__e, ShenFunc(symtl), gen7659)

																gen7661 := Call(__e, ShenFunc(symtl), gen7660)

																gen7662 := Call(__e, ShenFunc(symtl), gen7661)

																gen7663 := Call(__e, ShenFunc(sym_a), Nil, gen7662)

																var gen7664 Obj
																if True == gen7663 {
																	gen7656 := Call(__e, ShenFunc(symtl), V1097)

																	gen7657 := Call(__e, ShenFunc(symcons_2), gen7656)

																	var gen7658 Obj
																	if True == gen7657 {
																		gen7652 := Call(__e, ShenFunc(symtl), V1097)

																		gen7653 := Call(__e, ShenFunc(symtl), gen7652)

																		gen7654 := Call(__e, ShenFunc(sym_a), Nil, gen7653)

																		var gen7655 Obj
																		if True == gen7654 {
																			gen7651 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1098)

																			if True == gen7651 {
																				gen7655 = True
																			} else {
																				gen7655 = False
																			}

																		} else {
																			gen7655 = False
																		}
																		if True == gen7655 {
																			gen7658 = True
																		} else {
																			gen7658 = False
																		}

																	} else {
																		gen7658 = False
																	}
																	if True == gen7658 {
																		gen7664 = True
																	} else {
																		gen7664 = False
																	}

																} else {
																	gen7664 = False
																}
																if True == gen7664 {
																	gen7669 = True
																} else {
																	gen7669 = False
																}

															} else {
																gen7669 = False
															}
															if True == gen7669 {
																gen7674 = True
															} else {
																gen7674 = False
															}

														} else {
															gen7674 = False
														}
														if True == gen7674 {
															gen7678 = True
														} else {
															gen7678 = False
														}

													} else {
														gen7678 = False
													}
													if True == gen7678 {
														gen7682 = True
													} else {
														gen7682 = False
													}

												} else {
													gen7682 = False
												}
												if True == gen7682 {
													gen7685 = True
												} else {
													gen7685 = False
												}

											} else {
												gen7685 = False
											}
											if True == gen7685 {
												gen7687 = True
											} else {
												gen7687 = False
											}

										} else {
											gen7687 = False
										}
										if True == gen7687 {
											gen7562 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

											Z := gen7562
											gen7563 := Call(__e, ShenFunc(symtl), V1097)

											gen7564 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen7563)

											gen7565 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7564)

											gen7566 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen7565)

											gen7567 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7566)

											gen7568 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), Nil)

											gen7569 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.non-empty"), gen7568)

											gen7570 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen7569)

											gen7571 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7570)

											gen7572 := Call(__e, ShenFunc(symcons), Z, gen7571)

											gen7573 := Call(__e, ShenFunc(symhd), V1097)

											gen7574 := Call(__e, ShenFunc(symtl), gen7573)

											gen7575 := Call(__e, ShenFunc(symhd), gen7574)

											gen7576 := Call(__e, ShenFunc(symhd), gen7575)

											gen7577 := Call(__e, ShenFunc(symhd), V1097)

											gen7578 := Call(__e, ShenFunc(symtl), gen7577)

											gen7579 := Call(__e, ShenFunc(symhd), gen7578)

											gen7580 := Call(__e, ShenFunc(symtl), gen7579)

											gen7581 := Call(__e, ShenFunc(symhd), V1097)

											gen7582 := Call(__e, ShenFunc(symtl), gen7581)

											gen7583 := Call(__e, ShenFunc(symtl), gen7582)

											gen7584 := Call(__e, ShenFunc(symcons), gen7580, gen7583)

											gen7585 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7584)

											gen7586 := Call(__e, ShenFunc(symcons), Z, Nil)

											gen7587 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7586)

											gen7588 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen7587)

											gen7589 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7588)

											gen7590 := Call(__e, ShenFunc(symcons), gen7589, Nil)

											gen7591 := Call(__e, ShenFunc(symcons), gen7585, gen7590)

											gen7592 := Call(__e, ShenFunc(symcons), gen7591, Nil)

											gen7593 := Call(__e, ShenFunc(symcons), gen7576, gen7592)

											gen7594 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen7593)

											gen7595 := Call(__e, ShenFunc(symcons), Z, Nil)

											gen7596 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen7595)

											gen7597 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen7596)

											gen7598 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7597)

											gen7599 := Call(__e, ShenFunc(symcons), gen7598, Nil)

											gen7600 := Call(__e, ShenFunc(symcons), gen7594, gen7599)

											gen7601 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7600, MakeSymbol("+"))

											gen7602 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variable"), Nil)

											gen7603 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen7602)

											gen7604 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen7603)

											gen7605 := Call(__e, ShenFunc(symcons), Z, gen7604)

											gen7606 := Call(__e, ShenFunc(symhd), V1097)

											gen7607 := Call(__e, ShenFunc(symtl), gen7606)

											gen7608 := Call(__e, ShenFunc(symhd), gen7607)

											gen7609 := Call(__e, ShenFunc(symshen_4extract__vars), gen7608)

											gen7610 := Call(__e, ShenFunc(symhd), V1097)

											gen7611 := Call(__e, ShenFunc(symtl), gen7610)

											gen7612 := Call(__e, ShenFunc(symhd), gen7611)

											gen7613 := Call(__e, ShenFunc(symshen_4remove__modes), gen7612)

											gen7614 := Call(__e, ShenFunc(symshen_4rcons__form), gen7613)

											gen7615 := Call(__e, ShenFunc(symhd), V1097)

											gen7616 := Call(__e, ShenFunc(symtl), gen7615)

											gen7617 := Call(__e, ShenFunc(symtl), gen7616)

											gen7618 := Call(__e, ShenFunc(symhd), gen7617)

											gen7619 := Call(__e, ShenFunc(symshen_4mu__reduction), gen7618, MakeSymbol("+"))

											gen7620 := Call(__e, ShenFunc(symcons), gen7619, Nil)

											gen7621 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7620)

											gen7622 := Call(__e, ShenFunc(symcons), gen7614, gen7621)

											gen7623 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen7622)

											gen7624 := Call(__e, ShenFunc(symcons), Z, gen7623)

											gen7625 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen7624)

											gen7626 := Call(__e, ShenFunc(symcons), gen7625, Nil)

											gen7627 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7626)

											gen7628 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen7627)

											gen7629 := Call(__e, ShenFunc(symcons), gen7609, gen7628)

											gen7630 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7629)

											gen7631 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen7630)

											gen7632 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen7631)

											gen7633 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.rename"), gen7632)

											gen7634 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

											gen7635 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7634)

											gen7636 := Call(__e, ShenFunc(symcons), gen7633, gen7635)

											gen7637 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7636)

											gen7638 := Call(__e, ShenFunc(symcons), gen7605, gen7637)

											gen7639 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7638)

											gen7640 := Call(__e, ShenFunc(symcons), gen7639, Nil)

											gen7641 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen7640)

											gen7642 := Call(__e, ShenFunc(symcons), gen7601, gen7641)

											gen7643 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen7642)

											gen7644 := Call(__e, ShenFunc(symcons), gen7572, gen7643)

											gen7645 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen7644)

											gen7646 := Call(__e, ShenFunc(symcons), gen7645, Nil)

											gen7647 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen7646)

											gen7648 := Call(__e, ShenFunc(symcons), gen7567, gen7647)

											gen7649 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen7648)

											gen7650 := Call(__e, ShenFunc(symcons), Z, gen7649)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen7650)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mu_reduction"), gen8150)

		gen8158 := MakeNative(func(__e Evaluator) {
			V1100 := __e.Get(1)
			_ = V1100
			gen8157 := Call(__e, ShenFunc(symcons_2), V1100)

			if True == gen8157 {
				gen8151 := Call(__e, ShenFunc(symhd), V1100)

				gen8152 := Call(__e, ShenFunc(symshen_4rcons__form), gen8151)

				gen8153 := Call(__e, ShenFunc(symtl), V1100)

				gen8154 := Call(__e, ShenFunc(symshen_4rcons__form), gen8153)

				gen8155 := Call(__e, ShenFunc(symcons), gen8154, Nil)

				gen8156 := Call(__e, ShenFunc(symcons), gen8152, gen8155)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen8156)

				return

			} else {
				__e.Return(V1100)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rcons_form"), gen8158)

		gen8210 := MakeNative(func(__e Evaluator) {
			V1102 := __e.Get(1)
			_ = V1102
			gen8208 := Call(__e, ShenFunc(symcons_2), V1102)

			var gen8209 Obj
			if True == gen8208 {
				gen8205 := Call(__e, ShenFunc(symhd), V1102)

				gen8206 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen8205)

				var gen8207 Obj
				if True == gen8206 {
					gen8202 := Call(__e, ShenFunc(symtl), V1102)

					gen8203 := Call(__e, ShenFunc(symcons_2), gen8202)

					var gen8204 Obj
					if True == gen8203 {
						gen8198 := Call(__e, ShenFunc(symtl), V1102)

						gen8199 := Call(__e, ShenFunc(symtl), gen8198)

						gen8200 := Call(__e, ShenFunc(symcons_2), gen8199)

						var gen8201 Obj
						if True == gen8200 {
							gen8193 := Call(__e, ShenFunc(symtl), V1102)

							gen8194 := Call(__e, ShenFunc(symtl), gen8193)

							gen8195 := Call(__e, ShenFunc(symhd), gen8194)

							gen8196 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen8195)

							var gen8197 Obj
							if True == gen8196 {
								gen8189 := Call(__e, ShenFunc(symtl), V1102)

								gen8190 := Call(__e, ShenFunc(symtl), gen8189)

								gen8191 := Call(__e, ShenFunc(symtl), gen8190)

								gen8192 := Call(__e, ShenFunc(sym_a), Nil, gen8191)

								if True == gen8192 {
									gen8197 = True
								} else {
									gen8197 = False
								}

							} else {
								gen8197 = False
							}
							if True == gen8197 {
								gen8201 = True
							} else {
								gen8201 = False
							}

						} else {
							gen8201 = False
						}
						if True == gen8201 {
							gen8204 = True
						} else {
							gen8204 = False
						}

					} else {
						gen8204 = False
					}
					if True == gen8204 {
						gen8207 = True
					} else {
						gen8207 = False
					}

				} else {
					gen8207 = False
				}
				if True == gen8207 {
					gen8209 = True
				} else {
					gen8209 = False
				}

			} else {
				gen8209 = False
			}
			if True == gen8209 {
				gen8187 := Call(__e, ShenFunc(symtl), V1102)

				gen8188 := Call(__e, ShenFunc(symhd), gen8187)

				__e.TailApply(ShenFunc(symshen_4remove__modes), gen8188)

				return

			} else {
				gen8185 := Call(__e, ShenFunc(symcons_2), V1102)

				var gen8186 Obj
				if True == gen8185 {
					gen8182 := Call(__e, ShenFunc(symhd), V1102)

					gen8183 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen8182)

					var gen8184 Obj
					if True == gen8183 {
						gen8179 := Call(__e, ShenFunc(symtl), V1102)

						gen8180 := Call(__e, ShenFunc(symcons_2), gen8179)

						var gen8181 Obj
						if True == gen8180 {
							gen8175 := Call(__e, ShenFunc(symtl), V1102)

							gen8176 := Call(__e, ShenFunc(symtl), gen8175)

							gen8177 := Call(__e, ShenFunc(symcons_2), gen8176)

							var gen8178 Obj
							if True == gen8177 {
								gen8170 := Call(__e, ShenFunc(symtl), V1102)

								gen8171 := Call(__e, ShenFunc(symtl), gen8170)

								gen8172 := Call(__e, ShenFunc(symhd), gen8171)

								gen8173 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen8172)

								var gen8174 Obj
								if True == gen8173 {
									gen8166 := Call(__e, ShenFunc(symtl), V1102)

									gen8167 := Call(__e, ShenFunc(symtl), gen8166)

									gen8168 := Call(__e, ShenFunc(symtl), gen8167)

									gen8169 := Call(__e, ShenFunc(sym_a), Nil, gen8168)

									if True == gen8169 {
										gen8174 = True
									} else {
										gen8174 = False
									}

								} else {
									gen8174 = False
								}
								if True == gen8174 {
									gen8178 = True
								} else {
									gen8178 = False
								}

							} else {
								gen8178 = False
							}
							if True == gen8178 {
								gen8181 = True
							} else {
								gen8181 = False
							}

						} else {
							gen8181 = False
						}
						if True == gen8181 {
							gen8184 = True
						} else {
							gen8184 = False
						}

					} else {
						gen8184 = False
					}
					if True == gen8184 {
						gen8186 = True
					} else {
						gen8186 = False
					}

				} else {
					gen8186 = False
				}
				if True == gen8186 {
					gen8164 := Call(__e, ShenFunc(symtl), V1102)

					gen8165 := Call(__e, ShenFunc(symhd), gen8164)

					__e.TailApply(ShenFunc(symshen_4remove__modes), gen8165)

					return

				} else {
					gen8163 := Call(__e, ShenFunc(symcons_2), V1102)

					if True == gen8163 {
						gen8159 := Call(__e, ShenFunc(symhd), V1102)

						gen8160 := Call(__e, ShenFunc(symshen_4remove__modes), gen8159)

						gen8161 := Call(__e, ShenFunc(symtl), V1102)

						gen8162 := Call(__e, ShenFunc(symshen_4remove__modes), gen8161)

						__e.TailApply(ShenFunc(symcons), gen8160, gen8162)

						return

					} else {
						__e.Return(V1102)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove_modes"), gen8210)

		gen8213 := MakeNative(func(__e Evaluator) {
			V1105 := __e.Get(1)
			_ = V1105
			V1106 := __e.Get(2)
			_ = V1106
			gen8212 := Call(__e, ShenFunc(symvariable_2), V1105)

			if True == gen8212 {
				gen8211 := Call(__e, ShenFunc(symvariable_2), V1106)

				if True == gen8211 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ephemeral_variable?"), gen8213)

		gen8215 := MakeNative(func(__e Evaluator) {
			V1116 := __e.Get(1)
			_ = V1116
			gen8214 := Call(__e, ShenFunc(symcons_2), V1116)

			if True == gen8214 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog_constant?"), gen8215)

		gen8982 := MakeNative(func(__e Evaluator) {
			V1118 := __e.Get(1)
			_ = V1118
			gen8980 := Call(__e, ShenFunc(symcons_2), V1118)

			var gen8981 Obj
			if True == gen8980 {
				gen8977 := Call(__e, ShenFunc(symhd), V1118)

				gen8978 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen8977)

				var gen8979 Obj
				if True == gen8978 {
					gen8974 := Call(__e, ShenFunc(symtl), V1118)

					gen8975 := Call(__e, ShenFunc(symcons_2), gen8974)

					var gen8976 Obj
					if True == gen8975 {
						gen8970 := Call(__e, ShenFunc(symtl), V1118)

						gen8971 := Call(__e, ShenFunc(symtl), gen8970)

						gen8972 := Call(__e, ShenFunc(symcons_2), gen8971)

						var gen8973 Obj
						if True == gen8972 {
							gen8965 := Call(__e, ShenFunc(symtl), V1118)

							gen8966 := Call(__e, ShenFunc(symtl), gen8965)

							gen8967 := Call(__e, ShenFunc(symhd), gen8966)

							gen8968 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.be"), gen8967)

							var gen8969 Obj
							if True == gen8968 {
								gen8960 := Call(__e, ShenFunc(symtl), V1118)

								gen8961 := Call(__e, ShenFunc(symtl), gen8960)

								gen8962 := Call(__e, ShenFunc(symtl), gen8961)

								gen8963 := Call(__e, ShenFunc(symcons_2), gen8962)

								var gen8964 Obj
								if True == gen8963 {
									gen8954 := Call(__e, ShenFunc(symtl), V1118)

									gen8955 := Call(__e, ShenFunc(symtl), gen8954)

									gen8956 := Call(__e, ShenFunc(symtl), gen8955)

									gen8957 := Call(__e, ShenFunc(symtl), gen8956)

									gen8958 := Call(__e, ShenFunc(symcons_2), gen8957)

									var gen8959 Obj
									if True == gen8958 {
										gen8947 := Call(__e, ShenFunc(symtl), V1118)

										gen8948 := Call(__e, ShenFunc(symtl), gen8947)

										gen8949 := Call(__e, ShenFunc(symtl), gen8948)

										gen8950 := Call(__e, ShenFunc(symtl), gen8949)

										gen8951 := Call(__e, ShenFunc(symhd), gen8950)

										gen8952 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen8951)

										var gen8953 Obj
										if True == gen8952 {
											gen8940 := Call(__e, ShenFunc(symtl), V1118)

											gen8941 := Call(__e, ShenFunc(symtl), gen8940)

											gen8942 := Call(__e, ShenFunc(symtl), gen8941)

											gen8943 := Call(__e, ShenFunc(symtl), gen8942)

											gen8944 := Call(__e, ShenFunc(symtl), gen8943)

											gen8945 := Call(__e, ShenFunc(symcons_2), gen8944)

											var gen8946 Obj
											if True == gen8945 {
												gen8933 := Call(__e, ShenFunc(symtl), V1118)

												gen8934 := Call(__e, ShenFunc(symtl), gen8933)

												gen8935 := Call(__e, ShenFunc(symtl), gen8934)

												gen8936 := Call(__e, ShenFunc(symtl), gen8935)

												gen8937 := Call(__e, ShenFunc(symtl), gen8936)

												gen8938 := Call(__e, ShenFunc(symtl), gen8937)

												gen8939 := Call(__e, ShenFunc(sym_a), Nil, gen8938)

												if True == gen8939 {
													gen8946 = True
												} else {
													gen8946 = False
												}

											} else {
												gen8946 = False
											}
											if True == gen8946 {
												gen8953 = True
											} else {
												gen8953 = False
											}

										} else {
											gen8953 = False
										}
										if True == gen8953 {
											gen8959 = True
										} else {
											gen8959 = False
										}

									} else {
										gen8959 = False
									}
									if True == gen8959 {
										gen8964 = True
									} else {
										gen8964 = False
									}

								} else {
									gen8964 = False
								}
								if True == gen8964 {
									gen8969 = True
								} else {
									gen8969 = False
								}

							} else {
								gen8969 = False
							}
							if True == gen8969 {
								gen8973 = True
							} else {
								gen8973 = False
							}

						} else {
							gen8973 = False
						}
						if True == gen8973 {
							gen8976 = True
						} else {
							gen8976 = False
						}

					} else {
						gen8976 = False
					}
					if True == gen8976 {
						gen8979 = True
					} else {
						gen8979 = False
					}

				} else {
					gen8979 = False
				}
				if True == gen8979 {
					gen8981 = True
				} else {
					gen8981 = False
				}

			} else {
				gen8981 = False
			}
			if True == gen8981 {
				gen8916 := Call(__e, ShenFunc(symtl), V1118)

				gen8917 := Call(__e, ShenFunc(symhd), gen8916)

				gen8918 := Call(__e, ShenFunc(symtl), V1118)

				gen8919 := Call(__e, ShenFunc(symtl), gen8918)

				gen8920 := Call(__e, ShenFunc(symtl), gen8919)

				gen8921 := Call(__e, ShenFunc(symhd), gen8920)

				gen8922 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8921)

				gen8923 := Call(__e, ShenFunc(symtl), V1118)

				gen8924 := Call(__e, ShenFunc(symtl), gen8923)

				gen8925 := Call(__e, ShenFunc(symtl), gen8924)

				gen8926 := Call(__e, ShenFunc(symtl), gen8925)

				gen8927 := Call(__e, ShenFunc(symtl), gen8926)

				gen8928 := Call(__e, ShenFunc(symhd), gen8927)

				gen8929 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8928)

				gen8930 := Call(__e, ShenFunc(symcons), gen8929, Nil)

				gen8931 := Call(__e, ShenFunc(symcons), gen8922, gen8930)

				gen8932 := Call(__e, ShenFunc(symcons), gen8917, gen8931)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen8932)

				return

			} else {
				gen8914 := Call(__e, ShenFunc(symcons_2), V1118)

				var gen8915 Obj
				if True == gen8914 {
					gen8911 := Call(__e, ShenFunc(symhd), V1118)

					gen8912 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8911)

					var gen8913 Obj
					if True == gen8912 {
						gen8908 := Call(__e, ShenFunc(symtl), V1118)

						gen8909 := Call(__e, ShenFunc(symcons_2), gen8908)

						var gen8910 Obj
						if True == gen8909 {
							gen8904 := Call(__e, ShenFunc(symtl), V1118)

							gen8905 := Call(__e, ShenFunc(symhd), gen8904)

							gen8906 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.result"), gen8905)

							var gen8907 Obj
							if True == gen8906 {
								gen8900 := Call(__e, ShenFunc(symtl), V1118)

								gen8901 := Call(__e, ShenFunc(symtl), gen8900)

								gen8902 := Call(__e, ShenFunc(symcons_2), gen8901)

								var gen8903 Obj
								if True == gen8902 {
									gen8895 := Call(__e, ShenFunc(symtl), V1118)

									gen8896 := Call(__e, ShenFunc(symtl), gen8895)

									gen8897 := Call(__e, ShenFunc(symhd), gen8896)

									gen8898 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen8897)

									var gen8899 Obj
									if True == gen8898 {
										gen8890 := Call(__e, ShenFunc(symtl), V1118)

										gen8891 := Call(__e, ShenFunc(symtl), gen8890)

										gen8892 := Call(__e, ShenFunc(symtl), gen8891)

										gen8893 := Call(__e, ShenFunc(symcons_2), gen8892)

										var gen8894 Obj
										if True == gen8893 {
											gen8884 := Call(__e, ShenFunc(symtl), V1118)

											gen8885 := Call(__e, ShenFunc(symtl), gen8884)

											gen8886 := Call(__e, ShenFunc(symtl), gen8885)

											gen8887 := Call(__e, ShenFunc(symhd), gen8886)

											gen8888 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.dereferencing"), gen8887)

											var gen8889 Obj
											if True == gen8888 {
												gen8878 := Call(__e, ShenFunc(symtl), V1118)

												gen8879 := Call(__e, ShenFunc(symtl), gen8878)

												gen8880 := Call(__e, ShenFunc(symtl), gen8879)

												gen8881 := Call(__e, ShenFunc(symtl), gen8880)

												gen8882 := Call(__e, ShenFunc(symcons_2), gen8881)

												var gen8883 Obj
												if True == gen8882 {
													gen8872 := Call(__e, ShenFunc(symtl), V1118)

													gen8873 := Call(__e, ShenFunc(symtl), gen8872)

													gen8874 := Call(__e, ShenFunc(symtl), gen8873)

													gen8875 := Call(__e, ShenFunc(symtl), gen8874)

													gen8876 := Call(__e, ShenFunc(symtl), gen8875)

													gen8877 := Call(__e, ShenFunc(sym_a), Nil, gen8876)

													if True == gen8877 {
														gen8883 = True
													} else {
														gen8883 = False
													}

												} else {
													gen8883 = False
												}
												if True == gen8883 {
													gen8889 = True
												} else {
													gen8889 = False
												}

											} else {
												gen8889 = False
											}
											if True == gen8889 {
												gen8894 = True
											} else {
												gen8894 = False
											}

										} else {
											gen8894 = False
										}
										if True == gen8894 {
											gen8899 = True
										} else {
											gen8899 = False
										}

									} else {
										gen8899 = False
									}
									if True == gen8899 {
										gen8903 = True
									} else {
										gen8903 = False
									}

								} else {
									gen8903 = False
								}
								if True == gen8903 {
									gen8907 = True
								} else {
									gen8907 = False
								}

							} else {
								gen8907 = False
							}
							if True == gen8907 {
								gen8910 = True
							} else {
								gen8910 = False
							}

						} else {
							gen8910 = False
						}
						if True == gen8910 {
							gen8913 = True
						} else {
							gen8913 = False
						}

					} else {
						gen8913 = False
					}
					if True == gen8913 {
						gen8915 = True
					} else {
						gen8915 = False
					}

				} else {
					gen8915 = False
				}
				if True == gen8915 {
					gen8864 := Call(__e, ShenFunc(symtl), V1118)

					gen8865 := Call(__e, ShenFunc(symtl), gen8864)

					gen8866 := Call(__e, ShenFunc(symtl), gen8865)

					gen8867 := Call(__e, ShenFunc(symtl), gen8866)

					gen8868 := Call(__e, ShenFunc(symhd), gen8867)

					gen8869 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8868)

					gen8870 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

					gen8871 := Call(__e, ShenFunc(symcons), gen8869, gen8870)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.lazyderef"), gen8871)

					return

				} else {
					gen8862 := Call(__e, ShenFunc(symcons_2), V1118)

					var gen8863 Obj
					if True == gen8862 {
						gen8859 := Call(__e, ShenFunc(symhd), V1118)

						gen8860 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen8859)

						var gen8861 Obj
						if True == gen8860 {
							gen8856 := Call(__e, ShenFunc(symtl), V1118)

							gen8857 := Call(__e, ShenFunc(symcons_2), gen8856)

							var gen8858 Obj
							if True == gen8857 {
								gen8852 := Call(__e, ShenFunc(symtl), V1118)

								gen8853 := Call(__e, ShenFunc(symtl), gen8852)

								gen8854 := Call(__e, ShenFunc(symcons_2), gen8853)

								var gen8855 Obj
								if True == gen8854 {
									gen8847 := Call(__e, ShenFunc(symtl), V1118)

									gen8848 := Call(__e, ShenFunc(symtl), gen8847)

									gen8849 := Call(__e, ShenFunc(symhd), gen8848)

									gen8850 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen8849)

									var gen8851 Obj
									if True == gen8850 {
										gen8842 := Call(__e, ShenFunc(symtl), V1118)

										gen8843 := Call(__e, ShenFunc(symtl), gen8842)

										gen8844 := Call(__e, ShenFunc(symtl), gen8843)

										gen8845 := Call(__e, ShenFunc(symcons_2), gen8844)

										var gen8846 Obj
										if True == gen8845 {
											gen8836 := Call(__e, ShenFunc(symtl), V1118)

											gen8837 := Call(__e, ShenFunc(symtl), gen8836)

											gen8838 := Call(__e, ShenFunc(symtl), gen8837)

											gen8839 := Call(__e, ShenFunc(symtl), gen8838)

											gen8840 := Call(__e, ShenFunc(symcons_2), gen8839)

											var gen8841 Obj
											if True == gen8840 {
												gen8829 := Call(__e, ShenFunc(symtl), V1118)

												gen8830 := Call(__e, ShenFunc(symtl), gen8829)

												gen8831 := Call(__e, ShenFunc(symtl), gen8830)

												gen8832 := Call(__e, ShenFunc(symtl), gen8831)

												gen8833 := Call(__e, ShenFunc(symhd), gen8832)

												gen8834 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.else"), gen8833)

												var gen8835 Obj
												if True == gen8834 {
													gen8822 := Call(__e, ShenFunc(symtl), V1118)

													gen8823 := Call(__e, ShenFunc(symtl), gen8822)

													gen8824 := Call(__e, ShenFunc(symtl), gen8823)

													gen8825 := Call(__e, ShenFunc(symtl), gen8824)

													gen8826 := Call(__e, ShenFunc(symtl), gen8825)

													gen8827 := Call(__e, ShenFunc(symcons_2), gen8826)

													var gen8828 Obj
													if True == gen8827 {
														gen8815 := Call(__e, ShenFunc(symtl), V1118)

														gen8816 := Call(__e, ShenFunc(symtl), gen8815)

														gen8817 := Call(__e, ShenFunc(symtl), gen8816)

														gen8818 := Call(__e, ShenFunc(symtl), gen8817)

														gen8819 := Call(__e, ShenFunc(symtl), gen8818)

														gen8820 := Call(__e, ShenFunc(symtl), gen8819)

														gen8821 := Call(__e, ShenFunc(sym_a), Nil, gen8820)

														if True == gen8821 {
															gen8828 = True
														} else {
															gen8828 = False
														}

													} else {
														gen8828 = False
													}
													if True == gen8828 {
														gen8835 = True
													} else {
														gen8835 = False
													}

												} else {
													gen8835 = False
												}
												if True == gen8835 {
													gen8841 = True
												} else {
													gen8841 = False
												}

											} else {
												gen8841 = False
											}
											if True == gen8841 {
												gen8846 = True
											} else {
												gen8846 = False
											}

										} else {
											gen8846 = False
										}
										if True == gen8846 {
											gen8851 = True
										} else {
											gen8851 = False
										}

									} else {
										gen8851 = False
									}
									if True == gen8851 {
										gen8855 = True
									} else {
										gen8855 = False
									}

								} else {
									gen8855 = False
								}
								if True == gen8855 {
									gen8858 = True
								} else {
									gen8858 = False
								}

							} else {
								gen8858 = False
							}
							if True == gen8858 {
								gen8861 = True
							} else {
								gen8861 = False
							}

						} else {
							gen8861 = False
						}
						if True == gen8861 {
							gen8863 = True
						} else {
							gen8863 = False
						}

					} else {
						gen8863 = False
					}
					if True == gen8863 {
						gen8797 := Call(__e, ShenFunc(symtl), V1118)

						gen8798 := Call(__e, ShenFunc(symhd), gen8797)

						gen8799 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8798)

						gen8800 := Call(__e, ShenFunc(symtl), V1118)

						gen8801 := Call(__e, ShenFunc(symtl), gen8800)

						gen8802 := Call(__e, ShenFunc(symtl), gen8801)

						gen8803 := Call(__e, ShenFunc(symhd), gen8802)

						gen8804 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8803)

						gen8805 := Call(__e, ShenFunc(symtl), V1118)

						gen8806 := Call(__e, ShenFunc(symtl), gen8805)

						gen8807 := Call(__e, ShenFunc(symtl), gen8806)

						gen8808 := Call(__e, ShenFunc(symtl), gen8807)

						gen8809 := Call(__e, ShenFunc(symtl), gen8808)

						gen8810 := Call(__e, ShenFunc(symhd), gen8809)

						gen8811 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8810)

						gen8812 := Call(__e, ShenFunc(symcons), gen8811, Nil)

						gen8813 := Call(__e, ShenFunc(symcons), gen8804, gen8812)

						gen8814 := Call(__e, ShenFunc(symcons), gen8799, gen8813)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen8814)

						return

					} else {
						gen8795 := Call(__e, ShenFunc(symcons_2), V1118)

						var gen8796 Obj
						if True == gen8795 {
							gen8792 := Call(__e, ShenFunc(symtl), V1118)

							gen8793 := Call(__e, ShenFunc(symcons_2), gen8792)

							var gen8794 Obj
							if True == gen8793 {
								gen8788 := Call(__e, ShenFunc(symtl), V1118)

								gen8789 := Call(__e, ShenFunc(symhd), gen8788)

								gen8790 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen8789)

								var gen8791 Obj
								if True == gen8790 {
									gen8784 := Call(__e, ShenFunc(symtl), V1118)

									gen8785 := Call(__e, ShenFunc(symtl), gen8784)

									gen8786 := Call(__e, ShenFunc(symcons_2), gen8785)

									var gen8787 Obj
									if True == gen8786 {
										gen8779 := Call(__e, ShenFunc(symtl), V1118)

										gen8780 := Call(__e, ShenFunc(symtl), gen8779)

										gen8781 := Call(__e, ShenFunc(symhd), gen8780)

										gen8782 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.a"), gen8781)

										var gen8783 Obj
										if True == gen8782 {
											gen8774 := Call(__e, ShenFunc(symtl), V1118)

											gen8775 := Call(__e, ShenFunc(symtl), gen8774)

											gen8776 := Call(__e, ShenFunc(symtl), gen8775)

											gen8777 := Call(__e, ShenFunc(symcons_2), gen8776)

											var gen8778 Obj
											if True == gen8777 {
												gen8768 := Call(__e, ShenFunc(symtl), V1118)

												gen8769 := Call(__e, ShenFunc(symtl), gen8768)

												gen8770 := Call(__e, ShenFunc(symtl), gen8769)

												gen8771 := Call(__e, ShenFunc(symhd), gen8770)

												gen8772 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variable"), gen8771)

												var gen8773 Obj
												if True == gen8772 {
													gen8763 := Call(__e, ShenFunc(symtl), V1118)

													gen8764 := Call(__e, ShenFunc(symtl), gen8763)

													gen8765 := Call(__e, ShenFunc(symtl), gen8764)

													gen8766 := Call(__e, ShenFunc(symtl), gen8765)

													gen8767 := Call(__e, ShenFunc(sym_a), Nil, gen8766)

													if True == gen8767 {
														gen8773 = True
													} else {
														gen8773 = False
													}

												} else {
													gen8773 = False
												}
												if True == gen8773 {
													gen8778 = True
												} else {
													gen8778 = False
												}

											} else {
												gen8778 = False
											}
											if True == gen8778 {
												gen8783 = True
											} else {
												gen8783 = False
											}

										} else {
											gen8783 = False
										}
										if True == gen8783 {
											gen8787 = True
										} else {
											gen8787 = False
										}

									} else {
										gen8787 = False
									}
									if True == gen8787 {
										gen8791 = True
									} else {
										gen8791 = False
									}

								} else {
									gen8791 = False
								}
								if True == gen8791 {
									gen8794 = True
								} else {
									gen8794 = False
								}

							} else {
								gen8794 = False
							}
							if True == gen8794 {
								gen8796 = True
							} else {
								gen8796 = False
							}

						} else {
							gen8796 = False
						}
						if True == gen8796 {
							gen8761 := Call(__e, ShenFunc(symhd), V1118)

							gen8762 := Call(__e, ShenFunc(symcons), gen8761, Nil)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pvar?"), gen8762)

							return

						} else {
							gen8759 := Call(__e, ShenFunc(symcons_2), V1118)

							var gen8760 Obj
							if True == gen8759 {
								gen8756 := Call(__e, ShenFunc(symtl), V1118)

								gen8757 := Call(__e, ShenFunc(symcons_2), gen8756)

								var gen8758 Obj
								if True == gen8757 {
									gen8752 := Call(__e, ShenFunc(symtl), V1118)

									gen8753 := Call(__e, ShenFunc(symhd), gen8752)

									gen8754 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen8753)

									var gen8755 Obj
									if True == gen8754 {
										gen8748 := Call(__e, ShenFunc(symtl), V1118)

										gen8749 := Call(__e, ShenFunc(symtl), gen8748)

										gen8750 := Call(__e, ShenFunc(symcons_2), gen8749)

										var gen8751 Obj
										if True == gen8750 {
											gen8743 := Call(__e, ShenFunc(symtl), V1118)

											gen8744 := Call(__e, ShenFunc(symtl), gen8743)

											gen8745 := Call(__e, ShenFunc(symhd), gen8744)

											gen8746 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.a"), gen8745)

											var gen8747 Obj
											if True == gen8746 {
												gen8738 := Call(__e, ShenFunc(symtl), V1118)

												gen8739 := Call(__e, ShenFunc(symtl), gen8738)

												gen8740 := Call(__e, ShenFunc(symtl), gen8739)

												gen8741 := Call(__e, ShenFunc(symcons_2), gen8740)

												var gen8742 Obj
												if True == gen8741 {
													gen8732 := Call(__e, ShenFunc(symtl), V1118)

													gen8733 := Call(__e, ShenFunc(symtl), gen8732)

													gen8734 := Call(__e, ShenFunc(symtl), gen8733)

													gen8735 := Call(__e, ShenFunc(symhd), gen8734)

													gen8736 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.non-empty"), gen8735)

													var gen8737 Obj
													if True == gen8736 {
														gen8726 := Call(__e, ShenFunc(symtl), V1118)

														gen8727 := Call(__e, ShenFunc(symtl), gen8726)

														gen8728 := Call(__e, ShenFunc(symtl), gen8727)

														gen8729 := Call(__e, ShenFunc(symtl), gen8728)

														gen8730 := Call(__e, ShenFunc(symcons_2), gen8729)

														var gen8731 Obj
														if True == gen8730 {
															gen8719 := Call(__e, ShenFunc(symtl), V1118)

															gen8720 := Call(__e, ShenFunc(symtl), gen8719)

															gen8721 := Call(__e, ShenFunc(symtl), gen8720)

															gen8722 := Call(__e, ShenFunc(symtl), gen8721)

															gen8723 := Call(__e, ShenFunc(symhd), gen8722)

															gen8724 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), gen8723)

															var gen8725 Obj
															if True == gen8724 {
																gen8713 := Call(__e, ShenFunc(symtl), V1118)

																gen8714 := Call(__e, ShenFunc(symtl), gen8713)

																gen8715 := Call(__e, ShenFunc(symtl), gen8714)

																gen8716 := Call(__e, ShenFunc(symtl), gen8715)

																gen8717 := Call(__e, ShenFunc(symtl), gen8716)

																gen8718 := Call(__e, ShenFunc(sym_a), Nil, gen8717)

																if True == gen8718 {
																	gen8725 = True
																} else {
																	gen8725 = False
																}

															} else {
																gen8725 = False
															}
															if True == gen8725 {
																gen8731 = True
															} else {
																gen8731 = False
															}

														} else {
															gen8731 = False
														}
														if True == gen8731 {
															gen8737 = True
														} else {
															gen8737 = False
														}

													} else {
														gen8737 = False
													}
													if True == gen8737 {
														gen8742 = True
													} else {
														gen8742 = False
													}

												} else {
													gen8742 = False
												}
												if True == gen8742 {
													gen8747 = True
												} else {
													gen8747 = False
												}

											} else {
												gen8747 = False
											}
											if True == gen8747 {
												gen8751 = True
											} else {
												gen8751 = False
											}

										} else {
											gen8751 = False
										}
										if True == gen8751 {
											gen8755 = True
										} else {
											gen8755 = False
										}

									} else {
										gen8755 = False
									}
									if True == gen8755 {
										gen8758 = True
									} else {
										gen8758 = False
									}

								} else {
									gen8758 = False
								}
								if True == gen8758 {
									gen8760 = True
								} else {
									gen8760 = False
								}

							} else {
								gen8760 = False
							}
							if True == gen8760 {
								gen8711 := Call(__e, ShenFunc(symhd), V1118)

								gen8712 := Call(__e, ShenFunc(symcons), gen8711, Nil)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("cons?"), gen8712)

								return

							} else {
								gen8709 := Call(__e, ShenFunc(symcons_2), V1118)

								var gen8710 Obj
								if True == gen8709 {
									gen8706 := Call(__e, ShenFunc(symhd), V1118)

									gen8707 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.rename"), gen8706)

									var gen8708 Obj
									if True == gen8707 {
										gen8703 := Call(__e, ShenFunc(symtl), V1118)

										gen8704 := Call(__e, ShenFunc(symcons_2), gen8703)

										var gen8705 Obj
										if True == gen8704 {
											gen8699 := Call(__e, ShenFunc(symtl), V1118)

											gen8700 := Call(__e, ShenFunc(symhd), gen8699)

											gen8701 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8700)

											var gen8702 Obj
											if True == gen8701 {
												gen8695 := Call(__e, ShenFunc(symtl), V1118)

												gen8696 := Call(__e, ShenFunc(symtl), gen8695)

												gen8697 := Call(__e, ShenFunc(symcons_2), gen8696)

												var gen8698 Obj
												if True == gen8697 {
													gen8690 := Call(__e, ShenFunc(symtl), V1118)

													gen8691 := Call(__e, ShenFunc(symtl), gen8690)

													gen8692 := Call(__e, ShenFunc(symhd), gen8691)

													gen8693 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variables"), gen8692)

													var gen8694 Obj
													if True == gen8693 {
														gen8685 := Call(__e, ShenFunc(symtl), V1118)

														gen8686 := Call(__e, ShenFunc(symtl), gen8685)

														gen8687 := Call(__e, ShenFunc(symtl), gen8686)

														gen8688 := Call(__e, ShenFunc(symcons_2), gen8687)

														var gen8689 Obj
														if True == gen8688 {
															gen8679 := Call(__e, ShenFunc(symtl), V1118)

															gen8680 := Call(__e, ShenFunc(symtl), gen8679)

															gen8681 := Call(__e, ShenFunc(symtl), gen8680)

															gen8682 := Call(__e, ShenFunc(symhd), gen8681)

															gen8683 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen8682)

															var gen8684 Obj
															if True == gen8683 {
																gen8673 := Call(__e, ShenFunc(symtl), V1118)

																gen8674 := Call(__e, ShenFunc(symtl), gen8673)

																gen8675 := Call(__e, ShenFunc(symtl), gen8674)

																gen8676 := Call(__e, ShenFunc(symtl), gen8675)

																gen8677 := Call(__e, ShenFunc(symcons_2), gen8676)

																var gen8678 Obj
																if True == gen8677 {
																	gen8666 := Call(__e, ShenFunc(symtl), V1118)

																	gen8667 := Call(__e, ShenFunc(symtl), gen8666)

																	gen8668 := Call(__e, ShenFunc(symtl), gen8667)

																	gen8669 := Call(__e, ShenFunc(symtl), gen8668)

																	gen8670 := Call(__e, ShenFunc(symhd), gen8669)

																	gen8671 := Call(__e, ShenFunc(sym_a), Nil, gen8670)

																	var gen8672 Obj
																	if True == gen8671 {
																		gen8659 := Call(__e, ShenFunc(symtl), V1118)

																		gen8660 := Call(__e, ShenFunc(symtl), gen8659)

																		gen8661 := Call(__e, ShenFunc(symtl), gen8660)

																		gen8662 := Call(__e, ShenFunc(symtl), gen8661)

																		gen8663 := Call(__e, ShenFunc(symtl), gen8662)

																		gen8664 := Call(__e, ShenFunc(symcons_2), gen8663)

																		var gen8665 Obj
																		if True == gen8664 {
																			gen8651 := Call(__e, ShenFunc(symtl), V1118)

																			gen8652 := Call(__e, ShenFunc(symtl), gen8651)

																			gen8653 := Call(__e, ShenFunc(symtl), gen8652)

																			gen8654 := Call(__e, ShenFunc(symtl), gen8653)

																			gen8655 := Call(__e, ShenFunc(symtl), gen8654)

																			gen8656 := Call(__e, ShenFunc(symhd), gen8655)

																			gen8657 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen8656)

																			var gen8658 Obj
																			if True == gen8657 {
																				gen8643 := Call(__e, ShenFunc(symtl), V1118)

																				gen8644 := Call(__e, ShenFunc(symtl), gen8643)

																				gen8645 := Call(__e, ShenFunc(symtl), gen8644)

																				gen8646 := Call(__e, ShenFunc(symtl), gen8645)

																				gen8647 := Call(__e, ShenFunc(symtl), gen8646)

																				gen8648 := Call(__e, ShenFunc(symtl), gen8647)

																				gen8649 := Call(__e, ShenFunc(symcons_2), gen8648)

																				var gen8650 Obj
																				if True == gen8649 {
																					gen8634 := Call(__e, ShenFunc(symtl), V1118)

																					gen8635 := Call(__e, ShenFunc(symtl), gen8634)

																					gen8636 := Call(__e, ShenFunc(symtl), gen8635)

																					gen8637 := Call(__e, ShenFunc(symtl), gen8636)

																					gen8638 := Call(__e, ShenFunc(symtl), gen8637)

																					gen8639 := Call(__e, ShenFunc(symtl), gen8638)

																					gen8640 := Call(__e, ShenFunc(symhd), gen8639)

																					gen8641 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen8640)

																					var gen8642 Obj
																					if True == gen8641 {
																						gen8625 := Call(__e, ShenFunc(symtl), V1118)

																						gen8626 := Call(__e, ShenFunc(symtl), gen8625)

																						gen8627 := Call(__e, ShenFunc(symtl), gen8626)

																						gen8628 := Call(__e, ShenFunc(symtl), gen8627)

																						gen8629 := Call(__e, ShenFunc(symtl), gen8628)

																						gen8630 := Call(__e, ShenFunc(symtl), gen8629)

																						gen8631 := Call(__e, ShenFunc(symtl), gen8630)

																						gen8632 := Call(__e, ShenFunc(symcons_2), gen8631)

																						var gen8633 Obj
																						if True == gen8632 {
																							gen8616 := Call(__e, ShenFunc(symtl), V1118)

																							gen8617 := Call(__e, ShenFunc(symtl), gen8616)

																							gen8618 := Call(__e, ShenFunc(symtl), gen8617)

																							gen8619 := Call(__e, ShenFunc(symtl), gen8618)

																							gen8620 := Call(__e, ShenFunc(symtl), gen8619)

																							gen8621 := Call(__e, ShenFunc(symtl), gen8620)

																							gen8622 := Call(__e, ShenFunc(symtl), gen8621)

																							gen8623 := Call(__e, ShenFunc(symtl), gen8622)

																							gen8624 := Call(__e, ShenFunc(sym_a), Nil, gen8623)

																							if True == gen8624 {
																								gen8633 = True
																							} else {
																								gen8633 = False
																							}

																						} else {
																							gen8633 = False
																						}
																						if True == gen8633 {
																							gen8642 = True
																						} else {
																							gen8642 = False
																						}

																					} else {
																						gen8642 = False
																					}
																					if True == gen8642 {
																						gen8650 = True
																					} else {
																						gen8650 = False
																					}

																				} else {
																					gen8650 = False
																				}
																				if True == gen8650 {
																					gen8658 = True
																				} else {
																					gen8658 = False
																				}

																			} else {
																				gen8658 = False
																			}
																			if True == gen8658 {
																				gen8665 = True
																			} else {
																				gen8665 = False
																			}

																		} else {
																			gen8665 = False
																		}
																		if True == gen8665 {
																			gen8672 = True
																		} else {
																			gen8672 = False
																		}

																	} else {
																		gen8672 = False
																	}
																	if True == gen8672 {
																		gen8678 = True
																	} else {
																		gen8678 = False
																	}

																} else {
																	gen8678 = False
																}
																if True == gen8678 {
																	gen8684 = True
																} else {
																	gen8684 = False
																}

															} else {
																gen8684 = False
															}
															if True == gen8684 {
																gen8689 = True
															} else {
																gen8689 = False
															}

														} else {
															gen8689 = False
														}
														if True == gen8689 {
															gen8694 = True
														} else {
															gen8694 = False
														}

													} else {
														gen8694 = False
													}
													if True == gen8694 {
														gen8698 = True
													} else {
														gen8698 = False
													}

												} else {
													gen8698 = False
												}
												if True == gen8698 {
													gen8702 = True
												} else {
													gen8702 = False
												}

											} else {
												gen8702 = False
											}
											if True == gen8702 {
												gen8705 = True
											} else {
												gen8705 = False
											}

										} else {
											gen8705 = False
										}
										if True == gen8705 {
											gen8708 = True
										} else {
											gen8708 = False
										}

									} else {
										gen8708 = False
									}
									if True == gen8708 {
										gen8710 = True
									} else {
										gen8710 = False
									}

								} else {
									gen8710 = False
								}
								if True == gen8710 {
									gen8608 := Call(__e, ShenFunc(symtl), V1118)

									gen8609 := Call(__e, ShenFunc(symtl), gen8608)

									gen8610 := Call(__e, ShenFunc(symtl), gen8609)

									gen8611 := Call(__e, ShenFunc(symtl), gen8610)

									gen8612 := Call(__e, ShenFunc(symtl), gen8611)

									gen8613 := Call(__e, ShenFunc(symtl), gen8612)

									gen8614 := Call(__e, ShenFunc(symtl), gen8613)

									gen8615 := Call(__e, ShenFunc(symhd), gen8614)

									__e.TailApply(ShenFunc(symshen_4aum__to__shen), gen8615)

									return

								} else {
									gen8606 := Call(__e, ShenFunc(symcons_2), V1118)

									var gen8607 Obj
									if True == gen8606 {
										gen8603 := Call(__e, ShenFunc(symhd), V1118)

										gen8604 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.rename"), gen8603)

										var gen8605 Obj
										if True == gen8604 {
											gen8600 := Call(__e, ShenFunc(symtl), V1118)

											gen8601 := Call(__e, ShenFunc(symcons_2), gen8600)

											var gen8602 Obj
											if True == gen8601 {
												gen8596 := Call(__e, ShenFunc(symtl), V1118)

												gen8597 := Call(__e, ShenFunc(symhd), gen8596)

												gen8598 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8597)

												var gen8599 Obj
												if True == gen8598 {
													gen8592 := Call(__e, ShenFunc(symtl), V1118)

													gen8593 := Call(__e, ShenFunc(symtl), gen8592)

													gen8594 := Call(__e, ShenFunc(symcons_2), gen8593)

													var gen8595 Obj
													if True == gen8594 {
														gen8587 := Call(__e, ShenFunc(symtl), V1118)

														gen8588 := Call(__e, ShenFunc(symtl), gen8587)

														gen8589 := Call(__e, ShenFunc(symhd), gen8588)

														gen8590 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variables"), gen8589)

														var gen8591 Obj
														if True == gen8590 {
															gen8582 := Call(__e, ShenFunc(symtl), V1118)

															gen8583 := Call(__e, ShenFunc(symtl), gen8582)

															gen8584 := Call(__e, ShenFunc(symtl), gen8583)

															gen8585 := Call(__e, ShenFunc(symcons_2), gen8584)

															var gen8586 Obj
															if True == gen8585 {
																gen8576 := Call(__e, ShenFunc(symtl), V1118)

																gen8577 := Call(__e, ShenFunc(symtl), gen8576)

																gen8578 := Call(__e, ShenFunc(symtl), gen8577)

																gen8579 := Call(__e, ShenFunc(symhd), gen8578)

																gen8580 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen8579)

																var gen8581 Obj
																if True == gen8580 {
																	gen8570 := Call(__e, ShenFunc(symtl), V1118)

																	gen8571 := Call(__e, ShenFunc(symtl), gen8570)

																	gen8572 := Call(__e, ShenFunc(symtl), gen8571)

																	gen8573 := Call(__e, ShenFunc(symtl), gen8572)

																	gen8574 := Call(__e, ShenFunc(symcons_2), gen8573)

																	var gen8575 Obj
																	if True == gen8574 {
																		gen8563 := Call(__e, ShenFunc(symtl), V1118)

																		gen8564 := Call(__e, ShenFunc(symtl), gen8563)

																		gen8565 := Call(__e, ShenFunc(symtl), gen8564)

																		gen8566 := Call(__e, ShenFunc(symtl), gen8565)

																		gen8567 := Call(__e, ShenFunc(symhd), gen8566)

																		gen8568 := Call(__e, ShenFunc(symcons_2), gen8567)

																		var gen8569 Obj
																		if True == gen8568 {
																			gen8556 := Call(__e, ShenFunc(symtl), V1118)

																			gen8557 := Call(__e, ShenFunc(symtl), gen8556)

																			gen8558 := Call(__e, ShenFunc(symtl), gen8557)

																			gen8559 := Call(__e, ShenFunc(symtl), gen8558)

																			gen8560 := Call(__e, ShenFunc(symtl), gen8559)

																			gen8561 := Call(__e, ShenFunc(symcons_2), gen8560)

																			var gen8562 Obj
																			if True == gen8561 {
																				gen8548 := Call(__e, ShenFunc(symtl), V1118)

																				gen8549 := Call(__e, ShenFunc(symtl), gen8548)

																				gen8550 := Call(__e, ShenFunc(symtl), gen8549)

																				gen8551 := Call(__e, ShenFunc(symtl), gen8550)

																				gen8552 := Call(__e, ShenFunc(symtl), gen8551)

																				gen8553 := Call(__e, ShenFunc(symhd), gen8552)

																				gen8554 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen8553)

																				var gen8555 Obj
																				if True == gen8554 {
																					gen8540 := Call(__e, ShenFunc(symtl), V1118)

																					gen8541 := Call(__e, ShenFunc(symtl), gen8540)

																					gen8542 := Call(__e, ShenFunc(symtl), gen8541)

																					gen8543 := Call(__e, ShenFunc(symtl), gen8542)

																					gen8544 := Call(__e, ShenFunc(symtl), gen8543)

																					gen8545 := Call(__e, ShenFunc(symtl), gen8544)

																					gen8546 := Call(__e, ShenFunc(symcons_2), gen8545)

																					var gen8547 Obj
																					if True == gen8546 {
																						gen8531 := Call(__e, ShenFunc(symtl), V1118)

																						gen8532 := Call(__e, ShenFunc(symtl), gen8531)

																						gen8533 := Call(__e, ShenFunc(symtl), gen8532)

																						gen8534 := Call(__e, ShenFunc(symtl), gen8533)

																						gen8535 := Call(__e, ShenFunc(symtl), gen8534)

																						gen8536 := Call(__e, ShenFunc(symtl), gen8535)

																						gen8537 := Call(__e, ShenFunc(symhd), gen8536)

																						gen8538 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen8537)

																						var gen8539 Obj
																						if True == gen8538 {
																							gen8522 := Call(__e, ShenFunc(symtl), V1118)

																							gen8523 := Call(__e, ShenFunc(symtl), gen8522)

																							gen8524 := Call(__e, ShenFunc(symtl), gen8523)

																							gen8525 := Call(__e, ShenFunc(symtl), gen8524)

																							gen8526 := Call(__e, ShenFunc(symtl), gen8525)

																							gen8527 := Call(__e, ShenFunc(symtl), gen8526)

																							gen8528 := Call(__e, ShenFunc(symtl), gen8527)

																							gen8529 := Call(__e, ShenFunc(symcons_2), gen8528)

																							var gen8530 Obj
																							if True == gen8529 {
																								gen8513 := Call(__e, ShenFunc(symtl), V1118)

																								gen8514 := Call(__e, ShenFunc(symtl), gen8513)

																								gen8515 := Call(__e, ShenFunc(symtl), gen8514)

																								gen8516 := Call(__e, ShenFunc(symtl), gen8515)

																								gen8517 := Call(__e, ShenFunc(symtl), gen8516)

																								gen8518 := Call(__e, ShenFunc(symtl), gen8517)

																								gen8519 := Call(__e, ShenFunc(symtl), gen8518)

																								gen8520 := Call(__e, ShenFunc(symtl), gen8519)

																								gen8521 := Call(__e, ShenFunc(sym_a), Nil, gen8520)

																								if True == gen8521 {
																									gen8530 = True
																								} else {
																									gen8530 = False
																								}

																							} else {
																								gen8530 = False
																							}
																							if True == gen8530 {
																								gen8539 = True
																							} else {
																								gen8539 = False
																							}

																						} else {
																							gen8539 = False
																						}
																						if True == gen8539 {
																							gen8547 = True
																						} else {
																							gen8547 = False
																						}

																					} else {
																						gen8547 = False
																					}
																					if True == gen8547 {
																						gen8555 = True
																					} else {
																						gen8555 = False
																					}

																				} else {
																					gen8555 = False
																				}
																				if True == gen8555 {
																					gen8562 = True
																				} else {
																					gen8562 = False
																				}

																			} else {
																				gen8562 = False
																			}
																			if True == gen8562 {
																				gen8569 = True
																			} else {
																				gen8569 = False
																			}

																		} else {
																			gen8569 = False
																		}
																		if True == gen8569 {
																			gen8575 = True
																		} else {
																			gen8575 = False
																		}

																	} else {
																		gen8575 = False
																	}
																	if True == gen8575 {
																		gen8581 = True
																	} else {
																		gen8581 = False
																	}

																} else {
																	gen8581 = False
																}
																if True == gen8581 {
																	gen8586 = True
																} else {
																	gen8586 = False
																}

															} else {
																gen8586 = False
															}
															if True == gen8586 {
																gen8591 = True
															} else {
																gen8591 = False
															}

														} else {
															gen8591 = False
														}
														if True == gen8591 {
															gen8595 = True
														} else {
															gen8595 = False
														}

													} else {
														gen8595 = False
													}
													if True == gen8595 {
														gen8599 = True
													} else {
														gen8599 = False
													}

												} else {
													gen8599 = False
												}
												if True == gen8599 {
													gen8602 = True
												} else {
													gen8602 = False
												}

											} else {
												gen8602 = False
											}
											if True == gen8602 {
												gen8605 = True
											} else {
												gen8605 = False
											}

										} else {
											gen8605 = False
										}
										if True == gen8605 {
											gen8607 = True
										} else {
											gen8607 = False
										}

									} else {
										gen8607 = False
									}
									if True == gen8607 {
										gen8485 := Call(__e, ShenFunc(symtl), V1118)

										gen8486 := Call(__e, ShenFunc(symtl), gen8485)

										gen8487 := Call(__e, ShenFunc(symtl), gen8486)

										gen8488 := Call(__e, ShenFunc(symtl), gen8487)

										gen8489 := Call(__e, ShenFunc(symhd), gen8488)

										gen8490 := Call(__e, ShenFunc(symhd), gen8489)

										gen8491 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

										gen8492 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.newpv"), gen8491)

										gen8493 := Call(__e, ShenFunc(symtl), V1118)

										gen8494 := Call(__e, ShenFunc(symtl), gen8493)

										gen8495 := Call(__e, ShenFunc(symtl), gen8494)

										gen8496 := Call(__e, ShenFunc(symtl), gen8495)

										gen8497 := Call(__e, ShenFunc(symhd), gen8496)

										gen8498 := Call(__e, ShenFunc(symtl), gen8497)

										gen8499 := Call(__e, ShenFunc(symtl), V1118)

										gen8500 := Call(__e, ShenFunc(symtl), gen8499)

										gen8501 := Call(__e, ShenFunc(symtl), gen8500)

										gen8502 := Call(__e, ShenFunc(symtl), gen8501)

										gen8503 := Call(__e, ShenFunc(symtl), gen8502)

										gen8504 := Call(__e, ShenFunc(symcons), gen8498, gen8503)

										gen8505 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen8504)

										gen8506 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen8505)

										gen8507 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen8506)

										gen8508 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.rename"), gen8507)

										gen8509 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8508)

										gen8510 := Call(__e, ShenFunc(symcons), gen8509, Nil)

										gen8511 := Call(__e, ShenFunc(symcons), gen8492, gen8510)

										gen8512 := Call(__e, ShenFunc(symcons), gen8490, gen8511)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen8512)

										return

									} else {
										gen8483 := Call(__e, ShenFunc(symcons_2), V1118)

										var gen8484 Obj
										if True == gen8483 {
											gen8480 := Call(__e, ShenFunc(symhd), V1118)

											gen8481 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen8480)

											var gen8482 Obj
											if True == gen8481 {
												gen8477 := Call(__e, ShenFunc(symtl), V1118)

												gen8478 := Call(__e, ShenFunc(symcons_2), gen8477)

												var gen8479 Obj
												if True == gen8478 {
													gen8473 := Call(__e, ShenFunc(symtl), V1118)

													gen8474 := Call(__e, ShenFunc(symtl), gen8473)

													gen8475 := Call(__e, ShenFunc(symcons_2), gen8474)

													var gen8476 Obj
													if True == gen8475 {
														gen8468 := Call(__e, ShenFunc(symtl), V1118)

														gen8469 := Call(__e, ShenFunc(symtl), gen8468)

														gen8470 := Call(__e, ShenFunc(symhd), gen8469)

														gen8471 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.to"), gen8470)

														var gen8472 Obj
														if True == gen8471 {
															gen8463 := Call(__e, ShenFunc(symtl), V1118)

															gen8464 := Call(__e, ShenFunc(symtl), gen8463)

															gen8465 := Call(__e, ShenFunc(symtl), gen8464)

															gen8466 := Call(__e, ShenFunc(symcons_2), gen8465)

															var gen8467 Obj
															if True == gen8466 {
																gen8457 := Call(__e, ShenFunc(symtl), V1118)

																gen8458 := Call(__e, ShenFunc(symtl), gen8457)

																gen8459 := Call(__e, ShenFunc(symtl), gen8458)

																gen8460 := Call(__e, ShenFunc(symtl), gen8459)

																gen8461 := Call(__e, ShenFunc(symcons_2), gen8460)

																var gen8462 Obj
																if True == gen8461 {
																	gen8450 := Call(__e, ShenFunc(symtl), V1118)

																	gen8451 := Call(__e, ShenFunc(symtl), gen8450)

																	gen8452 := Call(__e, ShenFunc(symtl), gen8451)

																	gen8453 := Call(__e, ShenFunc(symtl), gen8452)

																	gen8454 := Call(__e, ShenFunc(symhd), gen8453)

																	gen8455 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen8454)

																	var gen8456 Obj
																	if True == gen8455 {
																		gen8443 := Call(__e, ShenFunc(symtl), V1118)

																		gen8444 := Call(__e, ShenFunc(symtl), gen8443)

																		gen8445 := Call(__e, ShenFunc(symtl), gen8444)

																		gen8446 := Call(__e, ShenFunc(symtl), gen8445)

																		gen8447 := Call(__e, ShenFunc(symtl), gen8446)

																		gen8448 := Call(__e, ShenFunc(symcons_2), gen8447)

																		var gen8449 Obj
																		if True == gen8448 {
																			gen8436 := Call(__e, ShenFunc(symtl), V1118)

																			gen8437 := Call(__e, ShenFunc(symtl), gen8436)

																			gen8438 := Call(__e, ShenFunc(symtl), gen8437)

																			gen8439 := Call(__e, ShenFunc(symtl), gen8438)

																			gen8440 := Call(__e, ShenFunc(symtl), gen8439)

																			gen8441 := Call(__e, ShenFunc(symtl), gen8440)

																			gen8442 := Call(__e, ShenFunc(sym_a), Nil, gen8441)

																			if True == gen8442 {
																				gen8449 = True
																			} else {
																				gen8449 = False
																			}

																		} else {
																			gen8449 = False
																		}
																		if True == gen8449 {
																			gen8456 = True
																		} else {
																			gen8456 = False
																		}

																	} else {
																		gen8456 = False
																	}
																	if True == gen8456 {
																		gen8462 = True
																	} else {
																		gen8462 = False
																	}

																} else {
																	gen8462 = False
																}
																if True == gen8462 {
																	gen8467 = True
																} else {
																	gen8467 = False
																}

															} else {
																gen8467 = False
															}
															if True == gen8467 {
																gen8472 = True
															} else {
																gen8472 = False
															}

														} else {
															gen8472 = False
														}
														if True == gen8472 {
															gen8476 = True
														} else {
															gen8476 = False
														}

													} else {
														gen8476 = False
													}
													if True == gen8476 {
														gen8479 = True
													} else {
														gen8479 = False
													}

												} else {
													gen8479 = False
												}
												if True == gen8479 {
													gen8482 = True
												} else {
													gen8482 = False
												}

											} else {
												gen8482 = False
											}
											if True == gen8482 {
												gen8484 = True
											} else {
												gen8484 = False
											}

										} else {
											gen8484 = False
										}
										if True == gen8484 {
											gen8404 := Call(__e, ShenFunc(symtl), V1118)

											gen8405 := Call(__e, ShenFunc(symhd), gen8404)

											gen8406 := Call(__e, ShenFunc(symtl), V1118)

											gen8407 := Call(__e, ShenFunc(symtl), gen8406)

											gen8408 := Call(__e, ShenFunc(symtl), gen8407)

											gen8409 := Call(__e, ShenFunc(symhd), gen8408)

											gen8410 := Call(__e, ShenFunc(symshen_4chwild), gen8409)

											gen8411 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

											gen8412 := Call(__e, ShenFunc(symcons), gen8410, gen8411)

											gen8413 := Call(__e, ShenFunc(symcons), gen8405, gen8412)

											gen8414 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.bindv"), gen8413)

											gen8415 := Call(__e, ShenFunc(symtl), V1118)

											gen8416 := Call(__e, ShenFunc(symtl), gen8415)

											gen8417 := Call(__e, ShenFunc(symtl), gen8416)

											gen8418 := Call(__e, ShenFunc(symtl), gen8417)

											gen8419 := Call(__e, ShenFunc(symtl), gen8418)

											gen8420 := Call(__e, ShenFunc(symhd), gen8419)

											gen8421 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen8420)

											gen8422 := Call(__e, ShenFunc(symtl), V1118)

											gen8423 := Call(__e, ShenFunc(symhd), gen8422)

											gen8424 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

											gen8425 := Call(__e, ShenFunc(symcons), gen8423, gen8424)

											gen8426 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.unbindv"), gen8425)

											gen8427 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

											gen8428 := Call(__e, ShenFunc(symcons), gen8426, gen8427)

											gen8429 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen8428)

											gen8430 := Call(__e, ShenFunc(symcons), gen8429, Nil)

											gen8431 := Call(__e, ShenFunc(symcons), gen8421, gen8430)

											gen8432 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen8431)

											gen8433 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen8432)

											gen8434 := Call(__e, ShenFunc(symcons), gen8433, Nil)

											gen8435 := Call(__e, ShenFunc(symcons), gen8414, gen8434)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen8435)

											return

										} else {
											gen8402 := Call(__e, ShenFunc(symcons_2), V1118)

											var gen8403 Obj
											if True == gen8402 {
												gen8399 := Call(__e, ShenFunc(symtl), V1118)

												gen8400 := Call(__e, ShenFunc(symcons_2), gen8399)

												var gen8401 Obj
												if True == gen8400 {
													gen8395 := Call(__e, ShenFunc(symtl), V1118)

													gen8396 := Call(__e, ShenFunc(symhd), gen8395)

													gen8397 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen8396)

													var gen8398 Obj
													if True == gen8397 {
														gen8391 := Call(__e, ShenFunc(symtl), V1118)

														gen8392 := Call(__e, ShenFunc(symtl), gen8391)

														gen8393 := Call(__e, ShenFunc(symcons_2), gen8392)

														var gen8394 Obj
														if True == gen8393 {
															gen8386 := Call(__e, ShenFunc(symtl), V1118)

															gen8387 := Call(__e, ShenFunc(symtl), gen8386)

															gen8388 := Call(__e, ShenFunc(symhd), gen8387)

															gen8389 := Call(__e, ShenFunc(sym_a), MakeSymbol("identical"), gen8388)

															var gen8390 Obj
															if True == gen8389 {
																gen8381 := Call(__e, ShenFunc(symtl), V1118)

																gen8382 := Call(__e, ShenFunc(symtl), gen8381)

																gen8383 := Call(__e, ShenFunc(symtl), gen8382)

																gen8384 := Call(__e, ShenFunc(symcons_2), gen8383)

																var gen8385 Obj
																if True == gen8384 {
																	gen8375 := Call(__e, ShenFunc(symtl), V1118)

																	gen8376 := Call(__e, ShenFunc(symtl), gen8375)

																	gen8377 := Call(__e, ShenFunc(symtl), gen8376)

																	gen8378 := Call(__e, ShenFunc(symhd), gen8377)

																	gen8379 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.to"), gen8378)

																	var gen8380 Obj
																	if True == gen8379 {
																		gen8369 := Call(__e, ShenFunc(symtl), V1118)

																		gen8370 := Call(__e, ShenFunc(symtl), gen8369)

																		gen8371 := Call(__e, ShenFunc(symtl), gen8370)

																		gen8372 := Call(__e, ShenFunc(symtl), gen8371)

																		gen8373 := Call(__e, ShenFunc(symcons_2), gen8372)

																		var gen8374 Obj
																		if True == gen8373 {
																			gen8363 := Call(__e, ShenFunc(symtl), V1118)

																			gen8364 := Call(__e, ShenFunc(symtl), gen8363)

																			gen8365 := Call(__e, ShenFunc(symtl), gen8364)

																			gen8366 := Call(__e, ShenFunc(symtl), gen8365)

																			gen8367 := Call(__e, ShenFunc(symtl), gen8366)

																			gen8368 := Call(__e, ShenFunc(sym_a), Nil, gen8367)

																			if True == gen8368 {
																				gen8374 = True
																			} else {
																				gen8374 = False
																			}

																		} else {
																			gen8374 = False
																		}
																		if True == gen8374 {
																			gen8380 = True
																		} else {
																			gen8380 = False
																		}

																	} else {
																		gen8380 = False
																	}
																	if True == gen8380 {
																		gen8385 = True
																	} else {
																		gen8385 = False
																	}

																} else {
																	gen8385 = False
																}
																if True == gen8385 {
																	gen8390 = True
																} else {
																	gen8390 = False
																}

															} else {
																gen8390 = False
															}
															if True == gen8390 {
																gen8394 = True
															} else {
																gen8394 = False
															}

														} else {
															gen8394 = False
														}
														if True == gen8394 {
															gen8398 = True
														} else {
															gen8398 = False
														}

													} else {
														gen8398 = False
													}
													if True == gen8398 {
														gen8401 = True
													} else {
														gen8401 = False
													}

												} else {
													gen8401 = False
												}
												if True == gen8401 {
													gen8403 = True
												} else {
													gen8403 = False
												}

											} else {
												gen8403 = False
											}
											if True == gen8403 {
												gen8355 := Call(__e, ShenFunc(symtl), V1118)

												gen8356 := Call(__e, ShenFunc(symtl), gen8355)

												gen8357 := Call(__e, ShenFunc(symtl), gen8356)

												gen8358 := Call(__e, ShenFunc(symtl), gen8357)

												gen8359 := Call(__e, ShenFunc(symhd), gen8358)

												gen8360 := Call(__e, ShenFunc(symhd), V1118)

												gen8361 := Call(__e, ShenFunc(symcons), gen8360, Nil)

												gen8362 := Call(__e, ShenFunc(symcons), gen8359, gen8361)

												__e.TailApply(ShenFunc(symcons), MakeSymbol("="), gen8362)

												return

											} else {
												gen8354 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.failed!"), V1118)

												if True == gen8354 {
													__e.Return(False)
													return
												} else {
													gen8352 := Call(__e, ShenFunc(symcons_2), V1118)

													var gen8353 Obj
													if True == gen8352 {
														gen8349 := Call(__e, ShenFunc(symhd), V1118)

														gen8350 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8349)

														var gen8351 Obj
														if True == gen8350 {
															gen8346 := Call(__e, ShenFunc(symtl), V1118)

															gen8347 := Call(__e, ShenFunc(symcons_2), gen8346)

															var gen8348 Obj
															if True == gen8347 {
																gen8342 := Call(__e, ShenFunc(symtl), V1118)

																gen8343 := Call(__e, ShenFunc(symhd), gen8342)

																gen8344 := Call(__e, ShenFunc(sym_a), MakeSymbol("head"), gen8343)

																var gen8345 Obj
																if True == gen8344 {
																	gen8338 := Call(__e, ShenFunc(symtl), V1118)

																	gen8339 := Call(__e, ShenFunc(symtl), gen8338)

																	gen8340 := Call(__e, ShenFunc(symcons_2), gen8339)

																	var gen8341 Obj
																	if True == gen8340 {
																		gen8333 := Call(__e, ShenFunc(symtl), V1118)

																		gen8334 := Call(__e, ShenFunc(symtl), gen8333)

																		gen8335 := Call(__e, ShenFunc(symhd), gen8334)

																		gen8336 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen8335)

																		var gen8337 Obj
																		if True == gen8336 {
																			gen8328 := Call(__e, ShenFunc(symtl), V1118)

																			gen8329 := Call(__e, ShenFunc(symtl), gen8328)

																			gen8330 := Call(__e, ShenFunc(symtl), gen8329)

																			gen8331 := Call(__e, ShenFunc(symcons_2), gen8330)

																			var gen8332 Obj
																			if True == gen8331 {
																				gen8323 := Call(__e, ShenFunc(symtl), V1118)

																				gen8324 := Call(__e, ShenFunc(symtl), gen8323)

																				gen8325 := Call(__e, ShenFunc(symtl), gen8324)

																				gen8326 := Call(__e, ShenFunc(symtl), gen8325)

																				gen8327 := Call(__e, ShenFunc(sym_a), Nil, gen8326)

																				if True == gen8327 {
																					gen8332 = True
																				} else {
																					gen8332 = False
																				}

																			} else {
																				gen8332 = False
																			}
																			if True == gen8332 {
																				gen8337 = True
																			} else {
																				gen8337 = False
																			}

																		} else {
																			gen8337 = False
																		}
																		if True == gen8337 {
																			gen8341 = True
																		} else {
																			gen8341 = False
																		}

																	} else {
																		gen8341 = False
																	}
																	if True == gen8341 {
																		gen8345 = True
																	} else {
																		gen8345 = False
																	}

																} else {
																	gen8345 = False
																}
																if True == gen8345 {
																	gen8348 = True
																} else {
																	gen8348 = False
																}

															} else {
																gen8348 = False
															}
															if True == gen8348 {
																gen8351 = True
															} else {
																gen8351 = False
															}

														} else {
															gen8351 = False
														}
														if True == gen8351 {
															gen8353 = True
														} else {
															gen8353 = False
														}

													} else {
														gen8353 = False
													}
													if True == gen8353 {
														gen8320 := Call(__e, ShenFunc(symtl), V1118)

														gen8321 := Call(__e, ShenFunc(symtl), gen8320)

														gen8322 := Call(__e, ShenFunc(symtl), gen8321)

														__e.TailApply(ShenFunc(symcons), MakeSymbol("hd"), gen8322)

														return

													} else {
														gen8318 := Call(__e, ShenFunc(symcons_2), V1118)

														var gen8319 Obj
														if True == gen8318 {
															gen8315 := Call(__e, ShenFunc(symhd), V1118)

															gen8316 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8315)

															var gen8317 Obj
															if True == gen8316 {
																gen8312 := Call(__e, ShenFunc(symtl), V1118)

																gen8313 := Call(__e, ShenFunc(symcons_2), gen8312)

																var gen8314 Obj
																if True == gen8313 {
																	gen8308 := Call(__e, ShenFunc(symtl), V1118)

																	gen8309 := Call(__e, ShenFunc(symhd), gen8308)

																	gen8310 := Call(__e, ShenFunc(sym_a), MakeSymbol("tail"), gen8309)

																	var gen8311 Obj
																	if True == gen8310 {
																		gen8304 := Call(__e, ShenFunc(symtl), V1118)

																		gen8305 := Call(__e, ShenFunc(symtl), gen8304)

																		gen8306 := Call(__e, ShenFunc(symcons_2), gen8305)

																		var gen8307 Obj
																		if True == gen8306 {
																			gen8299 := Call(__e, ShenFunc(symtl), V1118)

																			gen8300 := Call(__e, ShenFunc(symtl), gen8299)

																			gen8301 := Call(__e, ShenFunc(symhd), gen8300)

																			gen8302 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen8301)

																			var gen8303 Obj
																			if True == gen8302 {
																				gen8294 := Call(__e, ShenFunc(symtl), V1118)

																				gen8295 := Call(__e, ShenFunc(symtl), gen8294)

																				gen8296 := Call(__e, ShenFunc(symtl), gen8295)

																				gen8297 := Call(__e, ShenFunc(symcons_2), gen8296)

																				var gen8298 Obj
																				if True == gen8297 {
																					gen8289 := Call(__e, ShenFunc(symtl), V1118)

																					gen8290 := Call(__e, ShenFunc(symtl), gen8289)

																					gen8291 := Call(__e, ShenFunc(symtl), gen8290)

																					gen8292 := Call(__e, ShenFunc(symtl), gen8291)

																					gen8293 := Call(__e, ShenFunc(sym_a), Nil, gen8292)

																					if True == gen8293 {
																						gen8298 = True
																					} else {
																						gen8298 = False
																					}

																				} else {
																					gen8298 = False
																				}
																				if True == gen8298 {
																					gen8303 = True
																				} else {
																					gen8303 = False
																				}

																			} else {
																				gen8303 = False
																			}
																			if True == gen8303 {
																				gen8307 = True
																			} else {
																				gen8307 = False
																			}

																		} else {
																			gen8307 = False
																		}
																		if True == gen8307 {
																			gen8311 = True
																		} else {
																			gen8311 = False
																		}

																	} else {
																		gen8311 = False
																	}
																	if True == gen8311 {
																		gen8314 = True
																	} else {
																		gen8314 = False
																	}

																} else {
																	gen8314 = False
																}
																if True == gen8314 {
																	gen8317 = True
																} else {
																	gen8317 = False
																}

															} else {
																gen8317 = False
															}
															if True == gen8317 {
																gen8319 = True
															} else {
																gen8319 = False
															}

														} else {
															gen8319 = False
														}
														if True == gen8319 {
															gen8286 := Call(__e, ShenFunc(symtl), V1118)

															gen8287 := Call(__e, ShenFunc(symtl), gen8286)

															gen8288 := Call(__e, ShenFunc(symtl), gen8287)

															__e.TailApply(ShenFunc(symcons), MakeSymbol("tl"), gen8288)

															return

														} else {
															gen8284 := Call(__e, ShenFunc(symcons_2), V1118)

															var gen8285 Obj
															if True == gen8284 {
																gen8281 := Call(__e, ShenFunc(symhd), V1118)

																gen8282 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.pop"), gen8281)

																var gen8283 Obj
																if True == gen8282 {
																	gen8278 := Call(__e, ShenFunc(symtl), V1118)

																	gen8279 := Call(__e, ShenFunc(symcons_2), gen8278)

																	var gen8280 Obj
																	if True == gen8279 {
																		gen8274 := Call(__e, ShenFunc(symtl), V1118)

																		gen8275 := Call(__e, ShenFunc(symhd), gen8274)

																		gen8276 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8275)

																		var gen8277 Obj
																		if True == gen8276 {
																			gen8270 := Call(__e, ShenFunc(symtl), V1118)

																			gen8271 := Call(__e, ShenFunc(symtl), gen8270)

																			gen8272 := Call(__e, ShenFunc(symcons_2), gen8271)

																			var gen8273 Obj
																			if True == gen8272 {
																				gen8265 := Call(__e, ShenFunc(symtl), V1118)

																				gen8266 := Call(__e, ShenFunc(symtl), gen8265)

																				gen8267 := Call(__e, ShenFunc(symhd), gen8266)

																				gen8268 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.stack"), gen8267)

																				var gen8269 Obj
																				if True == gen8268 {
																					gen8261 := Call(__e, ShenFunc(symtl), V1118)

																					gen8262 := Call(__e, ShenFunc(symtl), gen8261)

																					gen8263 := Call(__e, ShenFunc(symtl), gen8262)

																					gen8264 := Call(__e, ShenFunc(sym_a), Nil, gen8263)

																					if True == gen8264 {
																						gen8269 = True
																					} else {
																						gen8269 = False
																					}

																				} else {
																					gen8269 = False
																				}
																				if True == gen8269 {
																					gen8273 = True
																				} else {
																					gen8273 = False
																				}

																			} else {
																				gen8273 = False
																			}
																			if True == gen8273 {
																				gen8277 = True
																			} else {
																				gen8277 = False
																			}

																		} else {
																			gen8277 = False
																		}
																		if True == gen8277 {
																			gen8280 = True
																		} else {
																			gen8280 = False
																		}

																	} else {
																		gen8280 = False
																	}
																	if True == gen8280 {
																		gen8283 = True
																	} else {
																		gen8283 = False
																	}

																} else {
																	gen8283 = False
																}
																if True == gen8283 {
																	gen8285 = True
																} else {
																	gen8285 = False
																}

															} else {
																gen8285 = False
															}
															if True == gen8285 {
																gen8256 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.incinfs"), Nil)

																gen8257 := Call(__e, ShenFunc(symcons), MakeSymbol("Continuation"), Nil)

																gen8258 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen8257)

																gen8259 := Call(__e, ShenFunc(symcons), gen8258, Nil)

																gen8260 := Call(__e, ShenFunc(symcons), gen8256, gen8259)

																__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen8260)

																return

															} else {
																gen8254 := Call(__e, ShenFunc(symcons_2), V1118)

																var gen8255 Obj
																if True == gen8254 {
																	gen8251 := Call(__e, ShenFunc(symhd), V1118)

																	gen8252 := Call(__e, ShenFunc(sym_a), MakeSymbol("call"), gen8251)

																	var gen8253 Obj
																	if True == gen8252 {
																		gen8248 := Call(__e, ShenFunc(symtl), V1118)

																		gen8249 := Call(__e, ShenFunc(symcons_2), gen8248)

																		var gen8250 Obj
																		if True == gen8249 {
																			gen8244 := Call(__e, ShenFunc(symtl), V1118)

																			gen8245 := Call(__e, ShenFunc(symhd), gen8244)

																			gen8246 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen8245)

																			var gen8247 Obj
																			if True == gen8246 {
																				gen8240 := Call(__e, ShenFunc(symtl), V1118)

																				gen8241 := Call(__e, ShenFunc(symtl), gen8240)

																				gen8242 := Call(__e, ShenFunc(symcons_2), gen8241)

																				var gen8243 Obj
																				if True == gen8242 {
																					gen8235 := Call(__e, ShenFunc(symtl), V1118)

																					gen8236 := Call(__e, ShenFunc(symtl), gen8235)

																					gen8237 := Call(__e, ShenFunc(symhd), gen8236)

																					gen8238 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.continuation"), gen8237)

																					var gen8239 Obj
																					if True == gen8238 {
																						gen8230 := Call(__e, ShenFunc(symtl), V1118)

																						gen8231 := Call(__e, ShenFunc(symtl), gen8230)

																						gen8232 := Call(__e, ShenFunc(symtl), gen8231)

																						gen8233 := Call(__e, ShenFunc(symcons_2), gen8232)

																						var gen8234 Obj
																						if True == gen8233 {
																							gen8225 := Call(__e, ShenFunc(symtl), V1118)

																							gen8226 := Call(__e, ShenFunc(symtl), gen8225)

																							gen8227 := Call(__e, ShenFunc(symtl), gen8226)

																							gen8228 := Call(__e, ShenFunc(symtl), gen8227)

																							gen8229 := Call(__e, ShenFunc(sym_a), Nil, gen8228)

																							if True == gen8229 {
																								gen8234 = True
																							} else {
																								gen8234 = False
																							}

																						} else {
																							gen8234 = False
																						}
																						if True == gen8234 {
																							gen8239 = True
																						} else {
																							gen8239 = False
																						}

																					} else {
																						gen8239 = False
																					}
																					if True == gen8239 {
																						gen8243 = True
																					} else {
																						gen8243 = False
																					}

																				} else {
																					gen8243 = False
																				}
																				if True == gen8243 {
																					gen8247 = True
																				} else {
																					gen8247 = False
																				}

																			} else {
																				gen8247 = False
																			}
																			if True == gen8247 {
																				gen8250 = True
																			} else {
																				gen8250 = False
																			}

																		} else {
																			gen8250 = False
																		}
																		if True == gen8250 {
																			gen8253 = True
																		} else {
																			gen8253 = False
																		}

																	} else {
																		gen8253 = False
																	}
																	if True == gen8253 {
																		gen8255 = True
																	} else {
																		gen8255 = False
																	}

																} else {
																	gen8255 = False
																}
																if True == gen8255 {
																	gen8216 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.incinfs"), Nil)

																	gen8217 := Call(__e, ShenFunc(symtl), V1118)

																	gen8218 := Call(__e, ShenFunc(symtl), gen8217)

																	gen8219 := Call(__e, ShenFunc(symtl), gen8218)

																	gen8220 := Call(__e, ShenFunc(symhd), gen8219)

																	gen8221 := Call(__e, ShenFunc(symshen_4chwild), gen8220)

																	gen8222 := Call(__e, ShenFunc(symshen_4call__the__continuation), gen8221, MakeSymbol("ProcessN"), MakeSymbol("Continuation"))

																	gen8223 := Call(__e, ShenFunc(symcons), gen8222, Nil)

																	gen8224 := Call(__e, ShenFunc(symcons), gen8216, gen8223)

																	__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen8224)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aum_to_shen"), gen8982)

		gen8987 := MakeNative(func(__e Evaluator) {
			V1120 := __e.Get(1)
			_ = V1120
			gen8986 := Call(__e, ShenFunc(sym_a), V1120, MakeSymbol("_"))

			if True == gen8986 {
				gen8985 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.newpv"), gen8985)

				return

			} else {
				gen8984 := Call(__e, ShenFunc(symcons_2), V1120)

				if True == gen8984 {
					gen8983 := MakeNative(func(__e Evaluator) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(ShenFunc(symshen_4chwild), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen8983, V1120)

					return

				} else {
					__e.Return(V1120)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.chwild"), gen8987)

		gen8998 := MakeNative(func(__e Evaluator) {
			V1122 := __e.Get(1)
			_ = V1122
			gen8988 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen8989 := Call(__e, ShenFunc(sym_5_1address), gen8988, V1122)

			gen8990 := Call(__e, ShenFunc(sym_7), gen8989, MakeNumber(1))

			Count_71 := gen8990
			gen8991 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen8992 := Call(__e, ShenFunc(symaddress_1_6), gen8991, V1122, Count_71)

			IncVar := gen8992
			_ = IncVar
			gen8993 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen8994 := Call(__e, ShenFunc(sym_5_1address), gen8993, V1122)

			Vector := gen8994
			gen8995 := Call(__e, ShenFunc(symlimit), Vector)

			gen8996 := Call(__e, ShenFunc(sym_a), Count_71, gen8995)

			var gen8997 Obj
			if True == gen8996 {
				gen8997 = Call(__e, ShenFunc(symshen_4resizeprocessvector), V1122, Count_71)

			} else {
				gen8997 = MakeSymbol("shen.skip")
			}
			ResizeVectorIfNeeded := gen8997
			_ = ResizeVectorIfNeeded
			__e.TailApply(ShenFunc(symshen_4mk_1pvar), Count_71)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newpv"), gen8998)

		gen9004 := MakeNative(func(__e Evaluator) {
			V1125 := __e.Get(1)
			_ = V1125
			V1126 := __e.Get(2)
			_ = V1126
			gen8999 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen9000 := Call(__e, ShenFunc(sym_5_1address), gen8999, V1125)

			Vector := gen9000
			gen9001 := Call(__e, ShenFunc(sym_7), V1126, V1126)

			gen9002 := Call(__e, ShenFunc(symshen_4resize_1vector), Vector, gen9001, MakeSymbol("shen.-null-"))

			BigVector := gen9002
			gen9003 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			__e.TailApply(ShenFunc(symaddress_1_6), gen9003, V1125, BigVector)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.resizeprocessvector"), gen9004)

		gen9009 := MakeNative(func(__e Evaluator) {
			V1130 := __e.Get(1)
			_ = V1130
			V1131 := __e.Get(2)
			_ = V1131
			V1132 := __e.Get(3)
			_ = V1132
			gen9005 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1131)

			gen9006 := Call(__e, ShenFunc(symabsvector), gen9005)

			gen9007 := Call(__e, ShenFunc(symaddress_1_6), gen9006, MakeNumber(0), V1131)

			BigVector := gen9007
			gen9008 := Call(__e, ShenFunc(symlimit), V1130)

			__e.TailApply(ShenFunc(symshen_4copy_1vector), V1130, BigVector, gen9008, V1131, V1132)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.resize-vector"), gen9009)

		gen9014 := MakeNative(func(__e Evaluator) {
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
			gen9010 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1140)

			gen9011 := Call(__e, ShenFunc(sym_7), V1141, MakeNumber(1))

			gen9012 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1140)

			gen9013 := Call(__e, ShenFunc(symshen_4copy_1vector_1stage_11), MakeNumber(1), V1138, V1139, gen9012)

			__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_12), gen9010, gen9011, V1142, gen9013)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector"), gen9014)

		gen9019 := MakeNative(func(__e Evaluator) {
			V1150 := __e.Get(1)
			_ = V1150
			V1151 := __e.Get(2)
			_ = V1151
			V1152 := __e.Get(3)
			_ = V1152
			V1153 := __e.Get(4)
			_ = V1153
			gen9018 := Call(__e, ShenFunc(sym_a), V1153, V1150)

			if True == gen9018 {
				__e.Return(V1152)
				return
			} else {
				gen9015 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1150)

				gen9016 := Call(__e, ShenFunc(sym_5_1address), V1151, V1150)

				gen9017 := Call(__e, ShenFunc(symaddress_1_6), V1152, V1150, gen9016)

				__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_11), gen9015, V1151, gen9017, V1153)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector-stage-1"), gen9019)

		gen9023 := MakeNative(func(__e Evaluator) {
			V1161 := __e.Get(1)
			_ = V1161
			V1162 := __e.Get(2)
			_ = V1162
			V1163 := __e.Get(3)
			_ = V1163
			V1164 := __e.Get(4)
			_ = V1164
			gen9022 := Call(__e, ShenFunc(sym_a), V1162, V1161)

			if True == gen9022 {
				__e.Return(V1164)
				return
			} else {
				gen9020 := Call(__e, ShenFunc(sym_7), V1161, MakeNumber(1))

				gen9021 := Call(__e, ShenFunc(symaddress_1_6), V1164, V1161, V1163)

				__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_12), gen9020, V1162, V1163, gen9021)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector-stage-2"), gen9023)

		gen9026 := MakeNative(func(__e Evaluator) {
			V1166 := __e.Get(1)
			_ = V1166
			gen9024 := Call(__e, ShenFunc(symabsvector), MakeNumber(2))

			gen9025 := Call(__e, ShenFunc(symaddress_1_6), gen9024, MakeNumber(0), MakeSymbol("shen.pvar"))

			__e.TailApply(ShenFunc(symaddress_1_6), gen9025, MakeNumber(1), V1166)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mk-pvar"), gen9026)

		gen9032 := MakeNative(func(__e Evaluator) {
			V1168 := __e.Get(1)
			_ = V1168
			gen9031 := Call(__e, ShenFunc(symabsvector_2), V1168)

			if True == gen9031 {
				gen9027 := MakeNative(func(__e Evaluator) {
					E := __e.Get(1)
					_ = E
					__e.Return(MakeSymbol("shen.not-pvar"))
					return
				}, 1)
				gen9028 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(sym_5_1address), V1168, MakeNumber(0))

					return
				}, 0)
				gen9029 := Try(__e, gen9028).Catch(gen9027)
				gen9030 := Call(__e, ShenFunc(sym_a), gen9029, MakeSymbol("shen.pvar"))

				if True == gen9030 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pvar?"), gen9032)

		gen9036 := MakeNative(func(__e Evaluator) {
			V1172 := __e.Get(1)
			_ = V1172
			V1173 := __e.Get(2)
			_ = V1173
			V1174 := __e.Get(3)
			_ = V1174
			gen9033 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen9034 := Call(__e, ShenFunc(sym_5_1address), gen9033, V1174)

			Vector := gen9034
			gen9035 := Call(__e, ShenFunc(sym_5_1address), V1172, MakeNumber(1))

			__e.TailApply(ShenFunc(symaddress_1_6), Vector, gen9035, V1173)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.bindv"), gen9036)

		gen9040 := MakeNative(func(__e Evaluator) {
			V1177 := __e.Get(1)
			_ = V1177
			V1178 := __e.Get(2)
			_ = V1178
			gen9037 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen9038 := Call(__e, ShenFunc(sym_5_1address), gen9037, V1178)

			Vector := gen9038
			gen9039 := Call(__e, ShenFunc(sym_5_1address), V1177, MakeNumber(1))

			__e.TailApply(ShenFunc(symaddress_1_6), Vector, gen9039, MakeSymbol("shen.-null-"))

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.unbindv"), gen9040)

		gen9043 := MakeNative(func(__e Evaluator) {
			gen9041 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			gen9042 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen9041)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*infs*"), gen9042)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.incinfs"), gen9043)

		gen9071 := MakeNative(func(__e Evaluator) {
			V1182 := __e.Get(1)
			_ = V1182
			V1183 := __e.Get(2)
			_ = V1183
			V1184 := __e.Get(3)
			_ = V1184
			gen9069 := Call(__e, ShenFunc(symcons_2), V1182)

			var gen9070 Obj
			if True == gen9069 {
				gen9066 := Call(__e, ShenFunc(symhd), V1182)

				gen9067 := Call(__e, ShenFunc(symcons_2), gen9066)

				var gen9068 Obj
				if True == gen9067 {
					gen9064 := Call(__e, ShenFunc(symtl), V1182)

					gen9065 := Call(__e, ShenFunc(sym_a), Nil, gen9064)

					if True == gen9065 {
						gen9068 = True
					} else {
						gen9068 = False
					}

				} else {
					gen9068 = False
				}
				if True == gen9068 {
					gen9070 = True
				} else {
					gen9070 = False
				}

			} else {
				gen9070 = False
			}
			if True == gen9070 {
				gen9057 := Call(__e, ShenFunc(symhd), V1182)

				gen9058 := Call(__e, ShenFunc(symhd), gen9057)

				gen9059 := Call(__e, ShenFunc(symhd), V1182)

				gen9060 := Call(__e, ShenFunc(symtl), gen9059)

				gen9061 := Call(__e, ShenFunc(symcons), V1184, Nil)

				gen9062 := Call(__e, ShenFunc(symcons), V1183, gen9061)

				gen9063 := Call(__e, ShenFunc(symappend), gen9060, gen9062)

				__e.TailApply(ShenFunc(symcons), gen9058, gen9063)

				return

			} else {
				gen9055 := Call(__e, ShenFunc(symcons_2), V1182)

				var gen9056 Obj
				if True == gen9055 {
					gen9053 := Call(__e, ShenFunc(symhd), V1182)

					gen9054 := Call(__e, ShenFunc(symcons_2), gen9053)

					if True == gen9054 {
						gen9056 = True
					} else {
						gen9056 = False
					}

				} else {
					gen9056 = False
				}
				if True == gen9056 {
					gen9044 := Call(__e, ShenFunc(symtl), V1182)

					gen9045 := Call(__e, ShenFunc(symshen_4newcontinuation), gen9044, V1183, V1184)

					NewContinuation := gen9045
					gen9046 := Call(__e, ShenFunc(symhd), V1182)

					gen9047 := Call(__e, ShenFunc(symhd), gen9046)

					gen9048 := Call(__e, ShenFunc(symhd), V1182)

					gen9049 := Call(__e, ShenFunc(symtl), gen9048)

					gen9050 := Call(__e, ShenFunc(symcons), NewContinuation, Nil)

					gen9051 := Call(__e, ShenFunc(symcons), V1183, gen9050)

					gen9052 := Call(__e, ShenFunc(symappend), gen9049, gen9051)

					__e.TailApply(ShenFunc(symcons), gen9047, gen9052)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call_the_continuation"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call_the_continuation"), gen9071)

		gen9088 := MakeNative(func(__e Evaluator) {
			V1188 := __e.Get(1)
			_ = V1188
			V1189 := __e.Get(2)
			_ = V1189
			V1190 := __e.Get(3)
			_ = V1190
			gen9087 := Call(__e, ShenFunc(sym_a), Nil, V1188)

			if True == gen9087 {
				__e.Return(V1190)
				return
			} else {
				gen9085 := Call(__e, ShenFunc(symcons_2), V1188)

				var gen9086 Obj
				if True == gen9085 {
					gen9083 := Call(__e, ShenFunc(symhd), V1188)

					gen9084 := Call(__e, ShenFunc(symcons_2), gen9083)

					if True == gen9084 {
						gen9086 = True
					} else {
						gen9086 = False
					}

				} else {
					gen9086 = False
				}
				if True == gen9086 {
					gen9072 := Call(__e, ShenFunc(symhd), V1188)

					gen9073 := Call(__e, ShenFunc(symhd), gen9072)

					gen9074 := Call(__e, ShenFunc(symhd), V1188)

					gen9075 := Call(__e, ShenFunc(symtl), gen9074)

					gen9076 := Call(__e, ShenFunc(symtl), V1188)

					gen9077 := Call(__e, ShenFunc(symshen_4newcontinuation), gen9076, V1189, V1190)

					gen9078 := Call(__e, ShenFunc(symcons), gen9077, Nil)

					gen9079 := Call(__e, ShenFunc(symcons), V1189, gen9078)

					gen9080 := Call(__e, ShenFunc(symappend), gen9075, gen9079)

					gen9081 := Call(__e, ShenFunc(symcons), gen9073, gen9080)

					gen9082 := Call(__e, ShenFunc(symcons), gen9081, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("freeze"), gen9082)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.newcontinuation"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newcontinuation"), gen9088)

		gen9089 := MakeNative(func(__e Evaluator) {
			V1198 := __e.Get(1)
			_ = V1198
			V1199 := __e.Get(2)
			_ = V1199
			V1200 := __e.Get(3)
			_ = V1200
			__e.TailApply(ShenFunc(symshen_4deref), V1198, V1199)

			return
		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("return"), gen9089)

		gen9093 := MakeNative(func(__e Evaluator) {
			V1208 := __e.Get(1)
			_ = V1208
			V1209 := __e.Get(2)
			_ = V1209
			V1210 := __e.Get(3)
			_ = V1210
			gen9090 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			gen9091 := Call(__e, ShenFunc(symshen_4app), gen9090, MakeString(" inferences\n"), MakeSymbol("shen.a"))

			gen9092 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen9091, gen9092)

			__e.TailApply(ShenFunc(symshen_4deref), V1208, V1209)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.measure&return"), gen9093)

		gen9096 := MakeNative(func(__e Evaluator) {
			V1215 := __e.Get(1)
			_ = V1215
			V1216 := __e.Get(2)
			_ = V1216
			V1217 := __e.Get(3)
			_ = V1217
			V1218 := __e.Get(4)
			_ = V1218
			gen9094 := Call(__e, ShenFunc(symshen_4lazyderef), V1215, V1217)

			gen9095 := Call(__e, ShenFunc(symshen_4lazyderef), V1216, V1217)

			__e.TailApply(ShenFunc(symshen_4lzy_a), gen9094, gen9095, V1217, V1218)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unify"), gen9096)

		gen9112 := MakeNative(func(__e Evaluator) {
			V1240 := __e.Get(1)
			_ = V1240
			V1241 := __e.Get(2)
			_ = V1241
			V1242 := __e.Get(3)
			_ = V1242
			V1243 := __e.Get(4)
			_ = V1243
			gen9111 := Call(__e, ShenFunc(sym_a), V1241, V1240)

			if True == gen9111 {
				__e.TailApply(ShenFunc(symthaw), V1243)

				return
			} else {
				gen9110 := Call(__e, ShenFunc(symshen_4pvar_2), V1240)

				if True == gen9110 {
					__e.TailApply(ShenFunc(symbind), V1240, V1241, V1242, V1243)

					return
				} else {
					gen9109 := Call(__e, ShenFunc(symshen_4pvar_2), V1241)

					if True == gen9109 {
						__e.TailApply(ShenFunc(symbind), V1241, V1240, V1242, V1243)

						return
					} else {
						gen9107 := Call(__e, ShenFunc(symcons_2), V1240)

						var gen9108 Obj
						if True == gen9107 {
							gen9106 := Call(__e, ShenFunc(symcons_2), V1241)

							if True == gen9106 {
								gen9108 = True
							} else {
								gen9108 = False
							}

						} else {
							gen9108 = False
						}
						if True == gen9108 {
							gen9097 := Call(__e, ShenFunc(symhd), V1240)

							gen9098 := Call(__e, ShenFunc(symshen_4lazyderef), gen9097, V1242)

							gen9099 := Call(__e, ShenFunc(symhd), V1241)

							gen9100 := Call(__e, ShenFunc(symshen_4lazyderef), gen9099, V1242)

							gen9105 := MakeNative(func(__e Evaluator) {
								gen9101 := Call(__e, ShenFunc(symtl), V1240)

								gen9102 := Call(__e, ShenFunc(symshen_4lazyderef), gen9101, V1242)

								gen9103 := Call(__e, ShenFunc(symtl), V1241)

								gen9104 := Call(__e, ShenFunc(symshen_4lazyderef), gen9103, V1242)

								__e.TailApply(ShenFunc(symshen_4lzy_a), gen9102, gen9104, V1242, V1243)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4lzy_a), gen9098, gen9100, V1242, gen9105)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy="), gen9112)

		gen9121 := MakeNative(func(__e Evaluator) {
			V1246 := __e.Get(1)
			_ = V1246
			V1247 := __e.Get(2)
			_ = V1247
			gen9120 := Call(__e, ShenFunc(symcons_2), V1246)

			if True == gen9120 {
				gen9116 := Call(__e, ShenFunc(symhd), V1246)

				gen9117 := Call(__e, ShenFunc(symshen_4deref), gen9116, V1247)

				gen9118 := Call(__e, ShenFunc(symtl), V1246)

				gen9119 := Call(__e, ShenFunc(symshen_4deref), gen9118, V1247)

				__e.TailApply(ShenFunc(symcons), gen9117, gen9119)

				return

			} else {
				gen9115 := Call(__e, ShenFunc(symshen_4pvar_2), V1246)

				if True == gen9115 {
					gen9113 := Call(__e, ShenFunc(symshen_4valvector), V1246, V1247)

					Value := gen9113
					gen9114 := Call(__e, ShenFunc(sym_a), Value, MakeSymbol("shen.-null-"))

					if True == gen9114 {
						__e.Return(V1246)
						return
					} else {
						__e.TailApply(ShenFunc(symshen_4deref), Value, V1247)

						return
					}

				} else {
					__e.Return(V1246)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.deref"), gen9121)

		gen9125 := MakeNative(func(__e Evaluator) {
			V1250 := __e.Get(1)
			_ = V1250
			V1251 := __e.Get(2)
			_ = V1251
			gen9124 := Call(__e, ShenFunc(symshen_4pvar_2), V1250)

			if True == gen9124 {
				gen9122 := Call(__e, ShenFunc(symshen_4valvector), V1250, V1251)

				Value := gen9122
				gen9123 := Call(__e, ShenFunc(sym_a), Value, MakeSymbol("shen.-null-"))

				if True == gen9123 {
					__e.Return(V1250)
					return
				} else {
					__e.TailApply(ShenFunc(symshen_4lazyderef), Value, V1251)

					return
				}

			} else {
				__e.Return(V1250)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lazyderef"), gen9125)

		gen9129 := MakeNative(func(__e Evaluator) {
			V1254 := __e.Get(1)
			_ = V1254
			V1255 := __e.Get(2)
			_ = V1255
			gen9126 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen9127 := Call(__e, ShenFunc(sym_5_1address), gen9126, V1255)

			gen9128 := Call(__e, ShenFunc(sym_5_1address), V1254, MakeNumber(1))

			__e.TailApply(ShenFunc(sym_5_1address), gen9127, gen9128)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.valvector"), gen9129)

		gen9132 := MakeNative(func(__e Evaluator) {
			V1260 := __e.Get(1)
			_ = V1260
			V1261 := __e.Get(2)
			_ = V1261
			V1262 := __e.Get(3)
			_ = V1262
			V1263 := __e.Get(4)
			_ = V1263
			gen9130 := Call(__e, ShenFunc(symshen_4lazyderef), V1260, V1262)

			gen9131 := Call(__e, ShenFunc(symshen_4lazyderef), V1261, V1262)

			__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen9130, gen9131, V1262, V1263)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unify!"), gen9132)

		gen9156 := MakeNative(func(__e Evaluator) {
			V1285 := __e.Get(1)
			_ = V1285
			V1286 := __e.Get(2)
			_ = V1286
			V1287 := __e.Get(3)
			_ = V1287
			V1288 := __e.Get(4)
			_ = V1288
			gen9155 := Call(__e, ShenFunc(sym_a), V1286, V1285)

			if True == gen9155 {
				__e.TailApply(ShenFunc(symthaw), V1288)

				return
			} else {
				gen9153 := Call(__e, ShenFunc(symshen_4pvar_2), V1285)

				var gen9154 Obj
				if True == gen9153 {
					gen9150 := Call(__e, ShenFunc(symshen_4deref), V1286, V1287)

					gen9151 := Call(__e, ShenFunc(symshen_4occurs_2), V1285, gen9150)

					gen9152 := Call(__e, ShenFunc(symnot), gen9151)

					if True == gen9152 {
						gen9154 = True
					} else {
						gen9154 = False
					}

				} else {
					gen9154 = False
				}
				if True == gen9154 {
					__e.TailApply(ShenFunc(symbind), V1285, V1286, V1287, V1288)

					return
				} else {
					gen9148 := Call(__e, ShenFunc(symshen_4pvar_2), V1286)

					var gen9149 Obj
					if True == gen9148 {
						gen9145 := Call(__e, ShenFunc(symshen_4deref), V1285, V1287)

						gen9146 := Call(__e, ShenFunc(symshen_4occurs_2), V1286, gen9145)

						gen9147 := Call(__e, ShenFunc(symnot), gen9146)

						if True == gen9147 {
							gen9149 = True
						} else {
							gen9149 = False
						}

					} else {
						gen9149 = False
					}
					if True == gen9149 {
						__e.TailApply(ShenFunc(symbind), V1286, V1285, V1287, V1288)

						return
					} else {
						gen9143 := Call(__e, ShenFunc(symcons_2), V1285)

						var gen9144 Obj
						if True == gen9143 {
							gen9142 := Call(__e, ShenFunc(symcons_2), V1286)

							if True == gen9142 {
								gen9144 = True
							} else {
								gen9144 = False
							}

						} else {
							gen9144 = False
						}
						if True == gen9144 {
							gen9133 := Call(__e, ShenFunc(symhd), V1285)

							gen9134 := Call(__e, ShenFunc(symshen_4lazyderef), gen9133, V1287)

							gen9135 := Call(__e, ShenFunc(symhd), V1286)

							gen9136 := Call(__e, ShenFunc(symshen_4lazyderef), gen9135, V1287)

							gen9141 := MakeNative(func(__e Evaluator) {
								gen9137 := Call(__e, ShenFunc(symtl), V1285)

								gen9138 := Call(__e, ShenFunc(symshen_4lazyderef), gen9137, V1287)

								gen9139 := Call(__e, ShenFunc(symtl), V1286)

								gen9140 := Call(__e, ShenFunc(symshen_4lazyderef), gen9139, V1287)

								__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen9138, gen9140, V1287, V1288)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen9134, gen9136, V1287, gen9141)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy=!"), gen9156)

		gen9163 := MakeNative(func(__e Evaluator) {
			V1300 := __e.Get(1)
			_ = V1300
			V1301 := __e.Get(2)
			_ = V1301
			gen9162 := Call(__e, ShenFunc(sym_a), V1301, V1300)

			if True == gen9162 {
				__e.Return(True)
				return
			} else {
				gen9161 := Call(__e, ShenFunc(symcons_2), V1301)

				if True == gen9161 {
					gen9159 := Call(__e, ShenFunc(symhd), V1301)

					gen9160 := Call(__e, ShenFunc(symshen_4occurs_2), V1300, gen9159)

					if True == gen9160 {
						__e.Return(True)
						return
					} else {
						gen9157 := Call(__e, ShenFunc(symtl), V1301)

						gen9158 := Call(__e, ShenFunc(symshen_4occurs_2), V1300, gen9157)

						if True == gen9158 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.occurs?"), gen9163)

		gen9166 := MakeNative(func(__e Evaluator) {
			V1306 := __e.Get(1)
			_ = V1306
			V1307 := __e.Get(2)
			_ = V1307
			V1308 := __e.Get(3)
			_ = V1308
			V1309 := __e.Get(4)
			_ = V1309
			gen9164 := Call(__e, ShenFunc(symshen_4lazyderef), V1306, V1308)

			gen9165 := Call(__e, ShenFunc(symshen_4lazyderef), V1307, V1308)

			__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen9164, gen9165, V1308, V1309)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("identical"), gen9166)

		gen9178 := MakeNative(func(__e Evaluator) {
			V1331 := __e.Get(1)
			_ = V1331
			V1332 := __e.Get(2)
			_ = V1332
			V1333 := __e.Get(3)
			_ = V1333
			V1334 := __e.Get(4)
			_ = V1334
			gen9177 := Call(__e, ShenFunc(sym_a), V1332, V1331)

			if True == gen9177 {
				__e.TailApply(ShenFunc(symthaw), V1334)

				return
			} else {
				gen9175 := Call(__e, ShenFunc(symcons_2), V1331)

				var gen9176 Obj
				if True == gen9175 {
					gen9174 := Call(__e, ShenFunc(symcons_2), V1332)

					if True == gen9174 {
						gen9176 = True
					} else {
						gen9176 = False
					}

				} else {
					gen9176 = False
				}
				if True == gen9176 {
					gen9167 := Call(__e, ShenFunc(symhd), V1331)

					gen9168 := Call(__e, ShenFunc(symshen_4lazyderef), gen9167, V1333)

					gen9169 := Call(__e, ShenFunc(symhd), V1332)

					gen9170 := Call(__e, ShenFunc(symshen_4lazyderef), gen9169, V1333)

					gen9173 := MakeNative(func(__e Evaluator) {
						gen9171 := Call(__e, ShenFunc(symtl), V1331)

						gen9172 := Call(__e, ShenFunc(symtl), V1332)

						__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen9171, gen9172, V1333, V1334)

						return

					}, 0)
					__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen9168, gen9170, V1333, gen9173)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy=="), gen9178)

		gen9181 := MakeNative(func(__e Evaluator) {
			V1336 := __e.Get(1)
			_ = V1336
			gen9179 := Call(__e, ShenFunc(sym_5_1address), V1336, MakeNumber(1))

			gen9180 := Call(__e, ShenFunc(symshen_4app), gen9179, MakeString(""), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symcn), MakeString("Var"), gen9180)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pvar"), gen9181)

		gen9183 := MakeNative(func(__e Evaluator) {
			V1341 := __e.Get(1)
			_ = V1341
			V1342 := __e.Get(2)
			_ = V1342
			V1343 := __e.Get(3)
			_ = V1343
			V1344 := __e.Get(4)
			_ = V1344
			Call(__e, ShenFunc(symshen_4bindv), V1341, V1342, V1343)
			gen9182 := Call(__e, ShenFunc(symthaw), V1344)

			Result := gen9182
			Call(__e, ShenFunc(symshen_4unbindv), V1341, V1343)
			__e.Return(Result)
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("bind"), gen9183)

		gen9188 := MakeNative(func(__e Evaluator) {
			V1362 := __e.Get(1)
			_ = V1362
			V1363 := __e.Get(2)
			_ = V1363
			V1364 := __e.Get(3)
			_ = V1364
			gen9187 := Call(__e, ShenFunc(sym_a), True, V1362)

			if True == gen9187 {
				__e.TailApply(ShenFunc(symthaw), V1364)

				return
			} else {
				gen9186 := Call(__e, ShenFunc(sym_a), False, V1362)

				if True == gen9186 {
					__e.Return(False)
					return
				} else {
					gen9184 := Call(__e, ShenFunc(symshen_4app), V1362, MakeString("%"), MakeSymbol("shen.s"))

					gen9185 := Call(__e, ShenFunc(symcn), MakeString("fwhen expects a boolean: not "), gen9184)

					__e.TailApply(ShenFunc(symsimple_1error), gen9185)

					return

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fwhen"), gen9188)

		gen9196 := MakeNative(func(__e Evaluator) {
			V1380 := __e.Get(1)
			_ = V1380
			V1381 := __e.Get(2)
			_ = V1381
			V1382 := __e.Get(3)
			_ = V1382
			gen9195 := Call(__e, ShenFunc(symcons_2), V1380)

			if True == gen9195 {
				gen9191 := Call(__e, ShenFunc(symhd), V1380)

				gen9192 := Call(__e, ShenFunc(symshen_4lazyderef), gen9191, V1381)

				gen9193 := Call(__e, ShenFunc(symfunction), gen9192)

				gen9194 := Call(__e, ShenFunc(symtl), V1380)

				__e.TailApply(ShenFunc(symshen_4call_1help), gen9193, gen9194, V1381, V1382)

				return

			} else {
				gen9190 := Call(__e, ShenFunc(symshen_4pvar_2), V1380)

				if True == gen9190 {
					gen9189 := Call(__e, ShenFunc(symshen_4lazyderef), V1380, V1381)

					__e.TailApply(ShenFunc(symcall), gen9189, V1381, V1382)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("call"), gen9196)

		gen9202 := MakeNative(func(__e Evaluator) {
			V1387 := __e.Get(1)
			_ = V1387
			V1388 := __e.Get(2)
			_ = V1388
			V1389 := __e.Get(3)
			_ = V1389
			V1390 := __e.Get(4)
			_ = V1390
			gen9201 := Call(__e, ShenFunc(sym_a), Nil, V1388)

			if True == gen9201 {
				__e.TailApply(V1387, V1389, V1390)

				return
			} else {
				gen9200 := Call(__e, ShenFunc(symcons_2), V1388)

				if True == gen9200 {
					gen9197 := Call(__e, ShenFunc(symhd), V1388)

					gen9198 := Call(__e, V1387, gen9197)

					gen9199 := Call(__e, ShenFunc(symtl), V1388)

					__e.TailApply(ShenFunc(symshen_4call_1help), gen9198, gen9199, V1389, V1390)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call-help"))

					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call-help"), gen9202)

		gen9216 := MakeNative(func(__e Evaluator) {
			V1392 := __e.Get(1)
			_ = V1392
			gen9214 := Call(__e, ShenFunc(symcons_2), V1392)

			var gen9215 Obj
			if True == gen9214 {
				gen9212 := Call(__e, ShenFunc(symhd), V1392)

				gen9213 := Call(__e, ShenFunc(symcons_2), gen9212)

				if True == gen9213 {
					gen9215 = True
				} else {
					gen9215 = False
				}

			} else {
				gen9215 = False
			}
			if True == gen9215 {
				gen9203 := Call(__e, ShenFunc(symshen_4start_1new_1prolog_1process))

				ProcessN := gen9203
				gen9204 := Call(__e, ShenFunc(symhd), V1392)

				gen9205 := Call(__e, ShenFunc(symhd), gen9204)

				gen9206 := Call(__e, ShenFunc(symhd), V1392)

				gen9207 := Call(__e, ShenFunc(symtl), gen9206)

				gen9208 := Call(__e, ShenFunc(symtl), V1392)

				gen9209 := Call(__e, ShenFunc(symcons), gen9208, Nil)

				gen9210 := Call(__e, ShenFunc(symcons), gen9207, gen9209)

				gen9211 := Call(__e, ShenFunc(symshen_4insert_1prolog_1variables), gen9210, ProcessN)

				__e.TailApply(ShenFunc(symshen_4intprolog_1help), gen9205, gen9211, ProcessN)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog"), gen9216)

		gen9228 := MakeNative(func(__e Evaluator) {
			V1396 := __e.Get(1)
			_ = V1396
			V1397 := __e.Get(2)
			_ = V1397
			V1398 := __e.Get(3)
			_ = V1398
			gen9226 := Call(__e, ShenFunc(symcons_2), V1397)

			var gen9227 Obj
			if True == gen9226 {
				gen9223 := Call(__e, ShenFunc(symtl), V1397)

				gen9224 := Call(__e, ShenFunc(symcons_2), gen9223)

				var gen9225 Obj
				if True == gen9224 {
					gen9220 := Call(__e, ShenFunc(symtl), V1397)

					gen9221 := Call(__e, ShenFunc(symtl), gen9220)

					gen9222 := Call(__e, ShenFunc(sym_a), Nil, gen9221)

					if True == gen9222 {
						gen9225 = True
					} else {
						gen9225 = False
					}

				} else {
					gen9225 = False
				}
				if True == gen9225 {
					gen9227 = True
				} else {
					gen9227 = False
				}

			} else {
				gen9227 = False
			}
			if True == gen9227 {
				gen9217 := Call(__e, ShenFunc(symhd), V1397)

				gen9218 := Call(__e, ShenFunc(symtl), V1397)

				gen9219 := Call(__e, ShenFunc(symhd), gen9218)

				__e.TailApply(ShenFunc(symshen_4intprolog_1help_1help), V1396, gen9217, gen9219, V1398)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog-help"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog-help"), gen9228)

		gen9235 := MakeNative(func(__e Evaluator) {
			V1403 := __e.Get(1)
			_ = V1403
			V1404 := __e.Get(2)
			_ = V1404
			V1405 := __e.Get(3)
			_ = V1405
			V1406 := __e.Get(4)
			_ = V1406
			gen9234 := Call(__e, ShenFunc(sym_a), Nil, V1404)

			if True == gen9234 {
				gen9233 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(symshen_4call_1rest), V1405, V1406)

					return
				}, 0)
				__e.TailApply(V1403, V1406, gen9233)

				return

			} else {
				gen9232 := Call(__e, ShenFunc(symcons_2), V1404)

				if True == gen9232 {
					gen9229 := Call(__e, ShenFunc(symhd), V1404)

					gen9230 := Call(__e, V1403, gen9229)

					gen9231 := Call(__e, ShenFunc(symtl), V1404)

					__e.TailApply(ShenFunc(symshen_4intprolog_1help_1help), gen9230, gen9231, V1405, V1406)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog-help-help"))

					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog-help-help"), gen9235)

		gen9269 := MakeNative(func(__e Evaluator) {
			V1411 := __e.Get(1)
			_ = V1411
			V1412 := __e.Get(2)
			_ = V1412
			gen9268 := Call(__e, ShenFunc(sym_a), Nil, V1411)

			if True == gen9268 {
				__e.Return(True)
				return
			} else {
				gen9266 := Call(__e, ShenFunc(symcons_2), V1411)

				var gen9267 Obj
				if True == gen9266 {
					gen9263 := Call(__e, ShenFunc(symhd), V1411)

					gen9264 := Call(__e, ShenFunc(symcons_2), gen9263)

					var gen9265 Obj
					if True == gen9264 {
						gen9260 := Call(__e, ShenFunc(symhd), V1411)

						gen9261 := Call(__e, ShenFunc(symtl), gen9260)

						gen9262 := Call(__e, ShenFunc(symcons_2), gen9261)

						if True == gen9262 {
							gen9265 = True
						} else {
							gen9265 = False
						}

					} else {
						gen9265 = False
					}
					if True == gen9265 {
						gen9267 = True
					} else {
						gen9267 = False
					}

				} else {
					gen9267 = False
				}
				if True == gen9267 {
					gen9248 := Call(__e, ShenFunc(symhd), V1411)

					gen9249 := Call(__e, ShenFunc(symhd), gen9248)

					f30 := gen9249
					gen9250 := Call(__e, ShenFunc(symhd), V1411)

					gen9251 := Call(__e, ShenFunc(symtl), gen9250)

					gen9252 := Call(__e, ShenFunc(symhd), gen9251)

					gen9253 := Call(__e, f30, gen9252)

					gen9254 := Call(__e, ShenFunc(symhd), V1411)

					gen9255 := Call(__e, ShenFunc(symtl), gen9254)

					gen9256 := Call(__e, ShenFunc(symtl), gen9255)

					gen9257 := Call(__e, ShenFunc(symcons), gen9253, gen9256)

					gen9258 := Call(__e, ShenFunc(symtl), V1411)

					gen9259 := Call(__e, ShenFunc(symcons), gen9257, gen9258)

					__e.TailApply(ShenFunc(symshen_4call_1rest), gen9259, V1412)

					return

				} else {
					gen9246 := Call(__e, ShenFunc(symcons_2), V1411)

					var gen9247 Obj
					if True == gen9246 {
						gen9243 := Call(__e, ShenFunc(symhd), V1411)

						gen9244 := Call(__e, ShenFunc(symcons_2), gen9243)

						var gen9245 Obj
						if True == gen9244 {
							gen9240 := Call(__e, ShenFunc(symhd), V1411)

							gen9241 := Call(__e, ShenFunc(symtl), gen9240)

							gen9242 := Call(__e, ShenFunc(sym_a), Nil, gen9241)

							if True == gen9242 {
								gen9245 = True
							} else {
								gen9245 = False
							}

						} else {
							gen9245 = False
						}
						if True == gen9245 {
							gen9247 = True
						} else {
							gen9247 = False
						}

					} else {
						gen9247 = False
					}
					if True == gen9247 {
						gen9236 := Call(__e, ShenFunc(symhd), V1411)

						gen9237 := Call(__e, ShenFunc(symhd), gen9236)

						f31 := gen9237
						gen9239 := MakeNative(func(__e Evaluator) {
							gen9238 := Call(__e, ShenFunc(symtl), V1411)

							__e.TailApply(ShenFunc(symshen_4call_1rest), gen9238, V1412)

							return

						}, 0)
						__e.TailApply(f31, V1412, gen9239)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call-rest"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call-rest"), gen9269)

		gen9273 := MakeNative(func(__e Evaluator) {
			gen9270 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*process-counter*"))

			gen9271 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen9270)

			gen9272 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*process-counter*"), gen9271)

			IncrementProcessCounter := gen9272
			__e.TailApply(ShenFunc(symshen_4initialise_1prolog), IncrementProcessCounter)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.start-new-prolog-process"), gen9273)

		gen9275 := MakeNative(func(__e Evaluator) {
			V1415 := __e.Get(1)
			_ = V1415
			V1416 := __e.Get(2)
			_ = V1416
			gen9274 := Call(__e, ShenFunc(symshen_4flatten), V1415)

			__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), V1415, gen9274, V1416)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-prolog-variables"), gen9275)

		gen9289 := MakeNative(func(__e Evaluator) {
			V1424 := __e.Get(1)
			_ = V1424
			V1425 := __e.Get(2)
			_ = V1425
			V1426 := __e.Get(3)
			_ = V1426
			gen9288 := Call(__e, ShenFunc(sym_a), Nil, V1425)

			if True == gen9288 {
				__e.Return(V1424)
				return
			} else {
				gen9286 := Call(__e, ShenFunc(symcons_2), V1425)

				var gen9287 Obj
				if True == gen9286 {
					gen9284 := Call(__e, ShenFunc(symhd), V1425)

					gen9285 := Call(__e, ShenFunc(symvariable_2), gen9284)

					if True == gen9285 {
						gen9287 = True
					} else {
						gen9287 = False
					}

				} else {
					gen9287 = False
				}
				if True == gen9287 {
					gen9278 := Call(__e, ShenFunc(symshen_4newpv), V1426)

					V := gen9278
					gen9279 := Call(__e, ShenFunc(symhd), V1425)

					gen9280 := Call(__e, ShenFunc(symsubst), V, gen9279, V1424)

					XV_cY := gen9280
					gen9281 := Call(__e, ShenFunc(symhd), V1425)

					gen9282 := Call(__e, ShenFunc(symtl), V1425)

					gen9283 := Call(__e, ShenFunc(symremove), gen9281, gen9282)

					Z_1Y := gen9283
					__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), XV_cY, Z_1Y, V1426)

					return

				} else {
					gen9277 := Call(__e, ShenFunc(symcons_2), V1425)

					if True == gen9277 {
						gen9276 := Call(__e, ShenFunc(symtl), V1425)

						__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), V1424, gen9276, V1426)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-prolog-variables-help"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-prolog-variables-help"), gen9289)

		gen9296 := MakeNative(func(__e Evaluator) {
			V1428 := __e.Get(1)
			_ = V1428
			gen9290 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen9291 := Call(__e, ShenFunc(symvector), MakeNumber(10))

			gen9292 := Call(__e, ShenFunc(symshen_4fillvector), gen9291, MakeNumber(1), MakeNumber(10), MakeSymbol("shen.-null-"))

			gen9293 := Call(__e, ShenFunc(symaddress_1_6), gen9290, V1428, gen9292)

			Vector := gen9293
			_ = Vector
			gen9294 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen9295 := Call(__e, ShenFunc(symaddress_1_6), gen9294, V1428, MakeNumber(1))

			Counter := gen9295
			_ = Counter
			__e.Return(V1428)
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.initialise-prolog"), gen9296)

		return

	}, 0))
}
