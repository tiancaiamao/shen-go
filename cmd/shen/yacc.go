package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen4361 := MakeNative(func(__e Evaluator) {
			V4724 := __e.Get(1)
			_ = V4724
			gen4359 := Call(__e, ShenFunc(symcons_2), V4724)

			var gen4360 Obj
			if True == gen4359 {
				gen4356 := Call(__e, ShenFunc(symhd), V4724)

				gen4357 := Call(__e, ShenFunc(sym_a), MakeSymbol("defcc"), gen4356)

				var gen4358 Obj
				if True == gen4357 {
					gen4354 := Call(__e, ShenFunc(symtl), V4724)

					gen4355 := Call(__e, ShenFunc(symcons_2), gen4354)

					if True == gen4355 {
						gen4358 = True
					} else {
						gen4358 = False
					}

				} else {
					gen4358 = False
				}
				if True == gen4358 {
					gen4360 = True
				} else {
					gen4360 = False
				}

			} else {
				gen4360 = False
			}
			if True == gen4360 {
				gen4350 := Call(__e, ShenFunc(symtl), V4724)

				gen4351 := Call(__e, ShenFunc(symhd), gen4350)

				gen4352 := Call(__e, ShenFunc(symtl), V4724)

				gen4353 := Call(__e, ShenFunc(symtl), gen4352)

				__e.TailApply(ShenFunc(symshen_4yacc_1_6shen), gen4351, gen4353)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.yacc"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc"), gen4361)

		gen4371 := MakeNative(func(__e Evaluator) {
			V4727 := __e.Get(1)
			_ = V4727
			V4728 := __e.Get(2)
			_ = V4728
			gen4362 := Call(__e, ShenFunc(symshen_4split__cc__rules), True, V4728, Nil)

			CCRules := gen4362
			gen4363 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4cc__body), X)

				return
			}, 1)
			gen4364 := Call(__e, ShenFunc(symmap), gen4363, CCRules)

			CCBody := gen4364
			gen4365 := Call(__e, ShenFunc(symshen_4yacc__cases), CCBody)

			YaccCases := gen4365
			gen4366 := Call(__e, ShenFunc(symshen_4kill_1code), YaccCases)

			gen4367 := Call(__e, ShenFunc(symcons), gen4366, Nil)

			gen4368 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen4367)

			gen4369 := Call(__e, ShenFunc(symcons), MakeSymbol("Stream"), gen4368)

			gen4370 := Call(__e, ShenFunc(symcons), V4727, gen4369)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("define"), gen4370)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc->shen"), gen4371)

		gen4381 := MakeNative(func(__e Evaluator) {
			V4730 := __e.Get(1)
			_ = V4730
			gen4379 := Call(__e, ShenFunc(symoccurrences), MakeSymbol("kill"), V4730)

			gen4380 := Call(__e, ShenFunc(sym_6), gen4379, MakeNumber(0))

			if True == gen4380 {
				gen4372 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen4373 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.analyse-kill"), gen4372)

				gen4374 := Call(__e, ShenFunc(symcons), gen4373, Nil)

				gen4375 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen4374)

				gen4376 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen4375)

				gen4377 := Call(__e, ShenFunc(symcons), gen4376, Nil)

				gen4378 := Call(__e, ShenFunc(symcons), V4730, gen4377)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("trap-error"), gen4378)

				return

			} else {
				__e.Return(V4730)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.kill-code"), gen4381)

		gen4382 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString("yacc kill"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("kill"), gen4382)

		gen4385 := MakeNative(func(__e Evaluator) {
			V4732 := __e.Get(1)
			_ = V4732
			gen4383 := Call(__e, ShenFunc(symerror_1to_1string), V4732)

			String := gen4383
			gen4384 := Call(__e, ShenFunc(sym_a), String, MakeString("yacc kill"))

			if True == gen4384 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V4732)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.analyse-kill"), gen4385)

		gen4404 := MakeNative(func(__e Evaluator) {
			V4738 := __e.Get(1)
			_ = V4738
			V4739 := __e.Get(2)
			_ = V4739
			V4740 := __e.Get(3)
			_ = V4740
			gen4402 := Call(__e, ShenFunc(sym_a), Nil, V4739)

			var gen4403 Obj
			if True == gen4402 {
				gen4401 := Call(__e, ShenFunc(sym_a), Nil, V4740)

				if True == gen4401 {
					gen4403 = True
				} else {
					gen4403 = False
				}

			} else {
				gen4403 = False
			}
			if True == gen4403 {
				__e.Return(Nil)
				return
			} else {
				gen4400 := Call(__e, ShenFunc(sym_a), Nil, V4739)

				if True == gen4400 {
					gen4398 := Call(__e, ShenFunc(symreverse), V4740)

					gen4399 := Call(__e, ShenFunc(symshen_4split__cc__rule), V4738, gen4398, Nil)

					__e.TailApply(ShenFunc(symcons), gen4399, Nil)

					return

				} else {
					gen4396 := Call(__e, ShenFunc(symcons_2), V4739)

					var gen4397 Obj
					if True == gen4396 {
						gen4394 := Call(__e, ShenFunc(symhd), V4739)

						gen4395 := Call(__e, ShenFunc(sym_a), MakeSymbol(";"), gen4394)

						if True == gen4395 {
							gen4397 = True
						} else {
							gen4397 = False
						}

					} else {
						gen4397 = False
					}
					if True == gen4397 {
						gen4390 := Call(__e, ShenFunc(symreverse), V4740)

						gen4391 := Call(__e, ShenFunc(symshen_4split__cc__rule), V4738, gen4390, Nil)

						gen4392 := Call(__e, ShenFunc(symtl), V4739)

						gen4393 := Call(__e, ShenFunc(symshen_4split__cc__rules), V4738, gen4392, Nil)

						__e.TailApply(ShenFunc(symcons), gen4391, gen4393)

						return

					} else {
						gen4389 := Call(__e, ShenFunc(symcons_2), V4739)

						if True == gen4389 {
							gen4386 := Call(__e, ShenFunc(symtl), V4739)

							gen4387 := Call(__e, ShenFunc(symhd), V4739)

							gen4388 := Call(__e, ShenFunc(symcons), gen4387, V4740)

							__e.TailApply(ShenFunc(symshen_4split__cc__rules), V4738, gen4386, gen4388)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.split_cc_rules"))

							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.split_cc_rules"), gen4404)

		gen4465 := MakeNative(func(__e Evaluator) {
			V4748 := __e.Get(1)
			_ = V4748
			V4749 := __e.Get(2)
			_ = V4749
			V4750 := __e.Get(3)
			_ = V4750
			gen4463 := Call(__e, ShenFunc(symcons_2), V4749)

			var gen4464 Obj
			if True == gen4463 {
				gen4460 := Call(__e, ShenFunc(symhd), V4749)

				gen4461 := Call(__e, ShenFunc(sym_a), MakeSymbol(":="), gen4460)

				var gen4462 Obj
				if True == gen4461 {
					gen4457 := Call(__e, ShenFunc(symtl), V4749)

					gen4458 := Call(__e, ShenFunc(symcons_2), gen4457)

					var gen4459 Obj
					if True == gen4458 {
						gen4454 := Call(__e, ShenFunc(symtl), V4749)

						gen4455 := Call(__e, ShenFunc(symtl), gen4454)

						gen4456 := Call(__e, ShenFunc(sym_a), Nil, gen4455)

						if True == gen4456 {
							gen4459 = True
						} else {
							gen4459 = False
						}

					} else {
						gen4459 = False
					}
					if True == gen4459 {
						gen4462 = True
					} else {
						gen4462 = False
					}

				} else {
					gen4462 = False
				}
				if True == gen4462 {
					gen4464 = True
				} else {
					gen4464 = False
				}

			} else {
				gen4464 = False
			}
			if True == gen4464 {
				gen4452 := Call(__e, ShenFunc(symreverse), V4750)

				gen4453 := Call(__e, ShenFunc(symtl), V4749)

				__e.TailApply(ShenFunc(symcons), gen4452, gen4453)

				return

			} else {
				gen4450 := Call(__e, ShenFunc(symcons_2), V4749)

				var gen4451 Obj
				if True == gen4450 {
					gen4447 := Call(__e, ShenFunc(symhd), V4749)

					gen4448 := Call(__e, ShenFunc(sym_a), MakeSymbol(":="), gen4447)

					var gen4449 Obj
					if True == gen4448 {
						gen4444 := Call(__e, ShenFunc(symtl), V4749)

						gen4445 := Call(__e, ShenFunc(symcons_2), gen4444)

						var gen4446 Obj
						if True == gen4445 {
							gen4440 := Call(__e, ShenFunc(symtl), V4749)

							gen4441 := Call(__e, ShenFunc(symtl), gen4440)

							gen4442 := Call(__e, ShenFunc(symcons_2), gen4441)

							var gen4443 Obj
							if True == gen4442 {
								gen4435 := Call(__e, ShenFunc(symtl), V4749)

								gen4436 := Call(__e, ShenFunc(symtl), gen4435)

								gen4437 := Call(__e, ShenFunc(symhd), gen4436)

								gen4438 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen4437)

								var gen4439 Obj
								if True == gen4438 {
									gen4430 := Call(__e, ShenFunc(symtl), V4749)

									gen4431 := Call(__e, ShenFunc(symtl), gen4430)

									gen4432 := Call(__e, ShenFunc(symtl), gen4431)

									gen4433 := Call(__e, ShenFunc(symcons_2), gen4432)

									var gen4434 Obj
									if True == gen4433 {
										gen4425 := Call(__e, ShenFunc(symtl), V4749)

										gen4426 := Call(__e, ShenFunc(symtl), gen4425)

										gen4427 := Call(__e, ShenFunc(symtl), gen4426)

										gen4428 := Call(__e, ShenFunc(symtl), gen4427)

										gen4429 := Call(__e, ShenFunc(sym_a), Nil, gen4428)

										if True == gen4429 {
											gen4434 = True
										} else {
											gen4434 = False
										}

									} else {
										gen4434 = False
									}
									if True == gen4434 {
										gen4439 = True
									} else {
										gen4439 = False
									}

								} else {
									gen4439 = False
								}
								if True == gen4439 {
									gen4443 = True
								} else {
									gen4443 = False
								}

							} else {
								gen4443 = False
							}
							if True == gen4443 {
								gen4446 = True
							} else {
								gen4446 = False
							}

						} else {
							gen4446 = False
						}
						if True == gen4446 {
							gen4449 = True
						} else {
							gen4449 = False
						}

					} else {
						gen4449 = False
					}
					if True == gen4449 {
						gen4451 = True
					} else {
						gen4451 = False
					}

				} else {
					gen4451 = False
				}
				if True == gen4451 {
					gen4414 := Call(__e, ShenFunc(symreverse), V4750)

					gen4415 := Call(__e, ShenFunc(symtl), V4749)

					gen4416 := Call(__e, ShenFunc(symtl), gen4415)

					gen4417 := Call(__e, ShenFunc(symtl), gen4416)

					gen4418 := Call(__e, ShenFunc(symhd), gen4417)

					gen4419 := Call(__e, ShenFunc(symtl), V4749)

					gen4420 := Call(__e, ShenFunc(symhd), gen4419)

					gen4421 := Call(__e, ShenFunc(symcons), gen4420, Nil)

					gen4422 := Call(__e, ShenFunc(symcons), gen4418, gen4421)

					gen4423 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen4422)

					gen4424 := Call(__e, ShenFunc(symcons), gen4423, Nil)

					__e.TailApply(ShenFunc(symcons), gen4414, gen4424)

					return

				} else {
					gen4413 := Call(__e, ShenFunc(sym_a), Nil, V4749)

					if True == gen4413 {
						Call(__e, ShenFunc(symshen_4semantic_1completion_1warning), V4748, V4750)
						gen4409 := Call(__e, ShenFunc(symreverse), V4750)

						gen4410 := Call(__e, ShenFunc(symshen_4default__semantics), gen4409)

						gen4411 := Call(__e, ShenFunc(symcons), gen4410, Nil)

						gen4412 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen4411)

						__e.TailApply(ShenFunc(symshen_4split__cc__rule), V4748, gen4412, V4750)

						return

					} else {
						gen4408 := Call(__e, ShenFunc(symcons_2), V4749)

						if True == gen4408 {
							gen4405 := Call(__e, ShenFunc(symtl), V4749)

							gen4406 := Call(__e, ShenFunc(symhd), V4749)

							gen4407 := Call(__e, ShenFunc(symcons), gen4406, V4750)

							__e.TailApply(ShenFunc(symshen_4split__cc__rule), V4748, gen4405, gen4407)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.split_cc_rule"))

							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.split_cc_rule"), gen4465)

		gen4473 := MakeNative(func(__e Evaluator) {
			V4761 := __e.Get(1)
			_ = V4761
			V4762 := __e.Get(2)
			_ = V4762
			gen4472 := Call(__e, ShenFunc(sym_a), True, V4761)

			if True == gen4472 {
				gen4466 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("warning: "), gen4466)

				gen4469 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					gen4467 := Call(__e, ShenFunc(symshen_4app), X, MakeString(" "), MakeSymbol("shen.a"))

					gen4468 := Call(__e, ShenFunc(symstoutput))

					__e.TailApply(ShenFunc(symshen_4prhush), gen4467, gen4468)

					return

				}, 1)
				gen4470 := Call(__e, ShenFunc(symreverse), V4762)

				Call(__e, ShenFunc(symshen_4for_1each), gen4469, gen4470)

				gen4471 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), MakeString("has no semantics.\n"), gen4471)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.semantic-completion-warning"), gen4473)

		gen4497 := MakeNative(func(__e Evaluator) {
			V4764 := __e.Get(1)
			_ = V4764
			gen4496 := Call(__e, ShenFunc(sym_a), Nil, V4764)

			if True == gen4496 {
				__e.Return(Nil)
				return
			} else {
				gen4494 := Call(__e, ShenFunc(symcons_2), V4764)

				var gen4495 Obj
				if True == gen4494 {
					gen4491 := Call(__e, ShenFunc(symtl), V4764)

					gen4492 := Call(__e, ShenFunc(sym_a), Nil, gen4491)

					var gen4493 Obj
					if True == gen4492 {
						gen4489 := Call(__e, ShenFunc(symhd), V4764)

						gen4490 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen4489)

						if True == gen4490 {
							gen4493 = True
						} else {
							gen4493 = False
						}

					} else {
						gen4493 = False
					}
					if True == gen4493 {
						gen4495 = True
					} else {
						gen4495 = False
					}

				} else {
					gen4495 = False
				}
				if True == gen4495 {
					__e.TailApply(ShenFunc(symhd), V4764)

					return
				} else {
					gen4487 := Call(__e, ShenFunc(symcons_2), V4764)

					var gen4488 Obj
					if True == gen4487 {
						gen4485 := Call(__e, ShenFunc(symhd), V4764)

						gen4486 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen4485)

						if True == gen4486 {
							gen4488 = True
						} else {
							gen4488 = False
						}

					} else {
						gen4488 = False
					}
					if True == gen4488 {
						gen4480 := Call(__e, ShenFunc(symhd), V4764)

						gen4481 := Call(__e, ShenFunc(symtl), V4764)

						gen4482 := Call(__e, ShenFunc(symshen_4default__semantics), gen4481)

						gen4483 := Call(__e, ShenFunc(symcons), gen4482, Nil)

						gen4484 := Call(__e, ShenFunc(symcons), gen4480, gen4483)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("append"), gen4484)

						return

					} else {
						gen4479 := Call(__e, ShenFunc(symcons_2), V4764)

						if True == gen4479 {
							gen4474 := Call(__e, ShenFunc(symhd), V4764)

							gen4475 := Call(__e, ShenFunc(symtl), V4764)

							gen4476 := Call(__e, ShenFunc(symshen_4default__semantics), gen4475)

							gen4477 := Call(__e, ShenFunc(symcons), gen4476, Nil)

							gen4478 := Call(__e, ShenFunc(symcons), gen4474, gen4477)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen4478)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.default_semantics"))

							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.default_semantics"), gen4497)

		gen4507 := MakeNative(func(__e Evaluator) {
			V4766 := __e.Get(1)
			_ = V4766
			gen4506 := Call(__e, ShenFunc(symsymbol_2), V4766)

			if True == gen4506 {
				gen4498 := Call(__e, ShenFunc(symexplode), V4766)

				gen4499 := Call(__e, ShenFunc(symshen_4strip_1pathname), gen4498)

				Cs := gen4499
				gen4503 := Call(__e, ShenFunc(symhd), Cs)

				gen4504 := Call(__e, ShenFunc(sym_a), gen4503, MakeString("<"))

				var gen4505 Obj
				if True == gen4504 {
					gen4500 := Call(__e, ShenFunc(symreverse), Cs)

					gen4501 := Call(__e, ShenFunc(symhd), gen4500)

					gen4502 := Call(__e, ShenFunc(sym_a), gen4501, MakeString(">"))

					if True == gen4502 {
						gen4505 = True
					} else {
						gen4505 = False
					}

				} else {
					gen4505 = False
				}
				if True == gen4505 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.grammar_symbol?"), gen4507)

		gen4527 := MakeNative(func(__e Evaluator) {
			V4768 := __e.Get(1)
			_ = V4768
			gen4525 := Call(__e, ShenFunc(symcons_2), V4768)

			var gen4526 Obj
			if True == gen4525 {
				gen4523 := Call(__e, ShenFunc(symtl), V4768)

				gen4524 := Call(__e, ShenFunc(sym_a), Nil, gen4523)

				if True == gen4524 {
					gen4526 = True
				} else {
					gen4526 = False
				}

			} else {
				gen4526 = False
			}
			if True == gen4526 {
				__e.TailApply(ShenFunc(symhd), V4768)

				return
			} else {
				gen4522 := Call(__e, ShenFunc(symcons_2), V4768)

				if True == gen4522 {
					P := MakeSymbol("YaccParse")
					gen4508 := Call(__e, ShenFunc(symhd), V4768)

					gen4509 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

					gen4510 := Call(__e, ShenFunc(symcons), gen4509, Nil)

					gen4511 := Call(__e, ShenFunc(symcons), P, gen4510)

					gen4512 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen4511)

					gen4513 := Call(__e, ShenFunc(symtl), V4768)

					gen4514 := Call(__e, ShenFunc(symshen_4yacc__cases), gen4513)

					gen4515 := Call(__e, ShenFunc(symcons), P, Nil)

					gen4516 := Call(__e, ShenFunc(symcons), gen4514, gen4515)

					gen4517 := Call(__e, ShenFunc(symcons), gen4512, gen4516)

					gen4518 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen4517)

					gen4519 := Call(__e, ShenFunc(symcons), gen4518, Nil)

					gen4520 := Call(__e, ShenFunc(symcons), gen4508, gen4519)

					gen4521 := Call(__e, ShenFunc(symcons), P, gen4520)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen4521)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.yacc_cases"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.yacc_cases"), gen4527)

		gen4539 := MakeNative(func(__e Evaluator) {
			V4770 := __e.Get(1)
			_ = V4770
			gen4537 := Call(__e, ShenFunc(symcons_2), V4770)

			var gen4538 Obj
			if True == gen4537 {
				gen4534 := Call(__e, ShenFunc(symtl), V4770)

				gen4535 := Call(__e, ShenFunc(symcons_2), gen4534)

				var gen4536 Obj
				if True == gen4535 {
					gen4531 := Call(__e, ShenFunc(symtl), V4770)

					gen4532 := Call(__e, ShenFunc(symtl), gen4531)

					gen4533 := Call(__e, ShenFunc(sym_a), Nil, gen4532)

					if True == gen4533 {
						gen4536 = True
					} else {
						gen4536 = False
					}

				} else {
					gen4536 = False
				}
				if True == gen4536 {
					gen4538 = True
				} else {
					gen4538 = False
				}

			} else {
				gen4538 = False
			}
			if True == gen4538 {
				gen4528 := Call(__e, ShenFunc(symhd), V4770)

				gen4529 := Call(__e, ShenFunc(symtl), V4770)

				gen4530 := Call(__e, ShenFunc(symhd), gen4529)

				__e.TailApply(ShenFunc(symshen_4syntax), gen4528, MakeSymbol("Stream"), gen4530)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.cc_body"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cc_body"), gen4539)

		gen4596 := MakeNative(func(__e Evaluator) {
			V4774 := __e.Get(1)
			_ = V4774
			V4775 := __e.Get(2)
			_ = V4775
			V4776 := __e.Get(3)
			_ = V4776
			gen4594 := Call(__e, ShenFunc(sym_a), Nil, V4774)

			var gen4595 Obj
			if True == gen4594 {
				gen4592 := Call(__e, ShenFunc(symcons_2), V4776)

				var gen4593 Obj
				if True == gen4592 {
					gen4589 := Call(__e, ShenFunc(symhd), V4776)

					gen4590 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen4589)

					var gen4591 Obj
					if True == gen4590 {
						gen4586 := Call(__e, ShenFunc(symtl), V4776)

						gen4587 := Call(__e, ShenFunc(symcons_2), gen4586)

						var gen4588 Obj
						if True == gen4587 {
							gen4582 := Call(__e, ShenFunc(symtl), V4776)

							gen4583 := Call(__e, ShenFunc(symtl), gen4582)

							gen4584 := Call(__e, ShenFunc(symcons_2), gen4583)

							var gen4585 Obj
							if True == gen4584 {
								gen4578 := Call(__e, ShenFunc(symtl), V4776)

								gen4579 := Call(__e, ShenFunc(symtl), gen4578)

								gen4580 := Call(__e, ShenFunc(symtl), gen4579)

								gen4581 := Call(__e, ShenFunc(sym_a), Nil, gen4580)

								if True == gen4581 {
									gen4585 = True
								} else {
									gen4585 = False
								}

							} else {
								gen4585 = False
							}
							if True == gen4585 {
								gen4588 = True
							} else {
								gen4588 = False
							}

						} else {
							gen4588 = False
						}
						if True == gen4588 {
							gen4591 = True
						} else {
							gen4591 = False
						}

					} else {
						gen4591 = False
					}
					if True == gen4591 {
						gen4593 = True
					} else {
						gen4593 = False
					}

				} else {
					gen4593 = False
				}
				if True == gen4593 {
					gen4595 = True
				} else {
					gen4595 = False
				}

			} else {
				gen4595 = False
			}
			if True == gen4595 {
				gen4562 := Call(__e, ShenFunc(symtl), V4776)

				gen4563 := Call(__e, ShenFunc(symhd), gen4562)

				gen4564 := Call(__e, ShenFunc(symshen_4semantics), gen4563)

				gen4565 := Call(__e, ShenFunc(symcons), V4775, Nil)

				gen4566 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4565)

				gen4567 := Call(__e, ShenFunc(symtl), V4776)

				gen4568 := Call(__e, ShenFunc(symtl), gen4567)

				gen4569 := Call(__e, ShenFunc(symhd), gen4568)

				gen4570 := Call(__e, ShenFunc(symshen_4semantics), gen4569)

				gen4571 := Call(__e, ShenFunc(symcons), gen4570, Nil)

				gen4572 := Call(__e, ShenFunc(symcons), gen4566, gen4571)

				gen4573 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4572)

				gen4574 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				gen4575 := Call(__e, ShenFunc(symcons), gen4574, Nil)

				gen4576 := Call(__e, ShenFunc(symcons), gen4573, gen4575)

				gen4577 := Call(__e, ShenFunc(symcons), gen4564, gen4576)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen4577)

				return

			} else {
				gen4561 := Call(__e, ShenFunc(sym_a), Nil, V4774)

				if True == gen4561 {
					gen4556 := Call(__e, ShenFunc(symcons), V4775, Nil)

					gen4557 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4556)

					gen4558 := Call(__e, ShenFunc(symshen_4semantics), V4776)

					gen4559 := Call(__e, ShenFunc(symcons), gen4558, Nil)

					gen4560 := Call(__e, ShenFunc(symcons), gen4557, gen4559)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pair"), gen4560)

					return

				} else {
					gen4555 := Call(__e, ShenFunc(symcons_2), V4774)

					if True == gen4555 {
						gen4553 := Call(__e, ShenFunc(symhd), V4774)

						gen4554 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), gen4553)

						if True == gen4554 {
							__e.TailApply(ShenFunc(symshen_4recursive__descent), V4774, V4775, V4776)

							return
						} else {
							gen4551 := Call(__e, ShenFunc(symhd), V4774)

							gen4552 := Call(__e, ShenFunc(symvariable_2), gen4551)

							if True == gen4552 {
								__e.TailApply(ShenFunc(symshen_4variable_1match), V4774, V4775, V4776)

								return
							} else {
								gen4549 := Call(__e, ShenFunc(symhd), V4774)

								gen4550 := Call(__e, ShenFunc(symshen_4jump__stream_2), gen4549)

								if True == gen4550 {
									__e.TailApply(ShenFunc(symshen_4jump__stream), V4774, V4775, V4776)

									return
								} else {
									gen4547 := Call(__e, ShenFunc(symhd), V4774)

									gen4548 := Call(__e, ShenFunc(symshen_4terminal_2), gen4547)

									if True == gen4548 {
										__e.TailApply(ShenFunc(symshen_4check__stream), V4774, V4775, V4776)

										return
									} else {
										gen4545 := Call(__e, ShenFunc(symhd), V4774)

										gen4546 := Call(__e, ShenFunc(symcons_2), gen4545)

										if True == gen4546 {
											gen4542 := Call(__e, ShenFunc(symhd), V4774)

											gen4543 := Call(__e, ShenFunc(symshen_4decons), gen4542)

											gen4544 := Call(__e, ShenFunc(symtl), V4774)

											__e.TailApply(ShenFunc(symshen_4list_1stream), gen4543, gen4544, V4775, V4776)

											return

										} else {
											gen4540 := Call(__e, ShenFunc(symhd), V4774)

											gen4541 := Call(__e, ShenFunc(symshen_4app), gen4540, MakeString(" is not legal syntax\n"), MakeSymbol("shen.a"))

											__e.TailApply(ShenFunc(symsimple_1error), gen4541)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.syntax"), gen4596)

		gen4630 := MakeNative(func(__e Evaluator) {
			V4781 := __e.Get(1)
			_ = V4781
			V4782 := __e.Get(2)
			_ = V4782
			V4783 := __e.Get(3)
			_ = V4783
			V4784 := __e.Get(4)
			_ = V4784
			gen4597 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4598 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4597)

			gen4599 := Call(__e, ShenFunc(symcons), gen4598, Nil)

			gen4600 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4599)

			gen4601 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4602 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen4601)

			gen4603 := Call(__e, ShenFunc(symcons), gen4602, Nil)

			gen4604 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4603)

			gen4605 := Call(__e, ShenFunc(symcons), gen4604, Nil)

			gen4606 := Call(__e, ShenFunc(symcons), gen4600, gen4605)

			gen4607 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen4606)

			Test := gen4607
			gen4608 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.place"))

			Placeholder := gen4608
			gen4609 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4610 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen4609)

			gen4611 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4612 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4611)

			gen4613 := Call(__e, ShenFunc(symcons), gen4612, Nil)

			gen4614 := Call(__e, ShenFunc(symcons), gen4610, gen4613)

			gen4615 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4614)

			gen4616 := Call(__e, ShenFunc(symshen_4syntax), V4782, gen4615, V4784)

			RunOn := gen4616
			gen4617 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4618 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen4617)

			gen4619 := Call(__e, ShenFunc(symcons), V4783, Nil)

			gen4620 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4619)

			gen4621 := Call(__e, ShenFunc(symcons), gen4620, Nil)

			gen4622 := Call(__e, ShenFunc(symcons), gen4618, gen4621)

			gen4623 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4622)

			gen4624 := Call(__e, ShenFunc(symshen_4syntax), V4781, gen4623, Placeholder)

			gen4625 := Call(__e, ShenFunc(symshen_4insert_1runon), RunOn, Placeholder, gen4624)

			Action := gen4625
			gen4626 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

			gen4627 := Call(__e, ShenFunc(symcons), gen4626, Nil)

			gen4628 := Call(__e, ShenFunc(symcons), Action, gen4627)

			gen4629 := Call(__e, ShenFunc(symcons), Test, gen4628)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen4629)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.list-stream"), gen4630)

		gen4676 := MakeNative(func(__e Evaluator) {
			V4786 := __e.Get(1)
			_ = V4786
			gen4674 := Call(__e, ShenFunc(symcons_2), V4786)

			var gen4675 Obj
			if True == gen4674 {
				gen4671 := Call(__e, ShenFunc(symhd), V4786)

				gen4672 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen4671)

				var gen4673 Obj
				if True == gen4672 {
					gen4668 := Call(__e, ShenFunc(symtl), V4786)

					gen4669 := Call(__e, ShenFunc(symcons_2), gen4668)

					var gen4670 Obj
					if True == gen4669 {
						gen4664 := Call(__e, ShenFunc(symtl), V4786)

						gen4665 := Call(__e, ShenFunc(symtl), gen4664)

						gen4666 := Call(__e, ShenFunc(symcons_2), gen4665)

						var gen4667 Obj
						if True == gen4666 {
							gen4659 := Call(__e, ShenFunc(symtl), V4786)

							gen4660 := Call(__e, ShenFunc(symtl), gen4659)

							gen4661 := Call(__e, ShenFunc(symhd), gen4660)

							gen4662 := Call(__e, ShenFunc(sym_a), Nil, gen4661)

							var gen4663 Obj
							if True == gen4662 {
								gen4655 := Call(__e, ShenFunc(symtl), V4786)

								gen4656 := Call(__e, ShenFunc(symtl), gen4655)

								gen4657 := Call(__e, ShenFunc(symtl), gen4656)

								gen4658 := Call(__e, ShenFunc(sym_a), Nil, gen4657)

								if True == gen4658 {
									gen4663 = True
								} else {
									gen4663 = False
								}

							} else {
								gen4663 = False
							}
							if True == gen4663 {
								gen4667 = True
							} else {
								gen4667 = False
							}

						} else {
							gen4667 = False
						}
						if True == gen4667 {
							gen4670 = True
						} else {
							gen4670 = False
						}

					} else {
						gen4670 = False
					}
					if True == gen4670 {
						gen4673 = True
					} else {
						gen4673 = False
					}

				} else {
					gen4673 = False
				}
				if True == gen4673 {
					gen4675 = True
				} else {
					gen4675 = False
				}

			} else {
				gen4675 = False
			}
			if True == gen4675 {
				gen4653 := Call(__e, ShenFunc(symtl), V4786)

				gen4654 := Call(__e, ShenFunc(symhd), gen4653)

				__e.TailApply(ShenFunc(symcons), gen4654, Nil)

				return

			} else {
				gen4651 := Call(__e, ShenFunc(symcons_2), V4786)

				var gen4652 Obj
				if True == gen4651 {
					gen4648 := Call(__e, ShenFunc(symhd), V4786)

					gen4649 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen4648)

					var gen4650 Obj
					if True == gen4649 {
						gen4645 := Call(__e, ShenFunc(symtl), V4786)

						gen4646 := Call(__e, ShenFunc(symcons_2), gen4645)

						var gen4647 Obj
						if True == gen4646 {
							gen4641 := Call(__e, ShenFunc(symtl), V4786)

							gen4642 := Call(__e, ShenFunc(symtl), gen4641)

							gen4643 := Call(__e, ShenFunc(symcons_2), gen4642)

							var gen4644 Obj
							if True == gen4643 {
								gen4637 := Call(__e, ShenFunc(symtl), V4786)

								gen4638 := Call(__e, ShenFunc(symtl), gen4637)

								gen4639 := Call(__e, ShenFunc(symtl), gen4638)

								gen4640 := Call(__e, ShenFunc(sym_a), Nil, gen4639)

								if True == gen4640 {
									gen4644 = True
								} else {
									gen4644 = False
								}

							} else {
								gen4644 = False
							}
							if True == gen4644 {
								gen4647 = True
							} else {
								gen4647 = False
							}

						} else {
							gen4647 = False
						}
						if True == gen4647 {
							gen4650 = True
						} else {
							gen4650 = False
						}

					} else {
						gen4650 = False
					}
					if True == gen4650 {
						gen4652 = True
					} else {
						gen4652 = False
					}

				} else {
					gen4652 = False
				}
				if True == gen4652 {
					gen4631 := Call(__e, ShenFunc(symtl), V4786)

					gen4632 := Call(__e, ShenFunc(symhd), gen4631)

					gen4633 := Call(__e, ShenFunc(symtl), V4786)

					gen4634 := Call(__e, ShenFunc(symtl), gen4633)

					gen4635 := Call(__e, ShenFunc(symhd), gen4634)

					gen4636 := Call(__e, ShenFunc(symshen_4decons), gen4635)

					__e.TailApply(ShenFunc(symcons), gen4632, gen4636)

					return

				} else {
					__e.Return(V4786)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decons"), gen4676)

		gen4700 := MakeNative(func(__e Evaluator) {
			V4801 := __e.Get(1)
			_ = V4801
			V4802 := __e.Get(2)
			_ = V4802
			V4803 := __e.Get(3)
			_ = V4803
			gen4698 := Call(__e, ShenFunc(symcons_2), V4803)

			var gen4699 Obj
			if True == gen4698 {
				gen4695 := Call(__e, ShenFunc(symhd), V4803)

				gen4696 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.pair"), gen4695)

				var gen4697 Obj
				if True == gen4696 {
					gen4692 := Call(__e, ShenFunc(symtl), V4803)

					gen4693 := Call(__e, ShenFunc(symcons_2), gen4692)

					var gen4694 Obj
					if True == gen4693 {
						gen4688 := Call(__e, ShenFunc(symtl), V4803)

						gen4689 := Call(__e, ShenFunc(symtl), gen4688)

						gen4690 := Call(__e, ShenFunc(symcons_2), gen4689)

						var gen4691 Obj
						if True == gen4690 {
							gen4683 := Call(__e, ShenFunc(symtl), V4803)

							gen4684 := Call(__e, ShenFunc(symtl), gen4683)

							gen4685 := Call(__e, ShenFunc(symtl), gen4684)

							gen4686 := Call(__e, ShenFunc(sym_a), Nil, gen4685)

							var gen4687 Obj
							if True == gen4686 {
								gen4679 := Call(__e, ShenFunc(symtl), V4803)

								gen4680 := Call(__e, ShenFunc(symtl), gen4679)

								gen4681 := Call(__e, ShenFunc(symhd), gen4680)

								gen4682 := Call(__e, ShenFunc(sym_a), gen4681, V4802)

								if True == gen4682 {
									gen4687 = True
								} else {
									gen4687 = False
								}

							} else {
								gen4687 = False
							}
							if True == gen4687 {
								gen4691 = True
							} else {
								gen4691 = False
							}

						} else {
							gen4691 = False
						}
						if True == gen4691 {
							gen4694 = True
						} else {
							gen4694 = False
						}

					} else {
						gen4694 = False
					}
					if True == gen4694 {
						gen4697 = True
					} else {
						gen4697 = False
					}

				} else {
					gen4697 = False
				}
				if True == gen4697 {
					gen4699 = True
				} else {
					gen4699 = False
				}

			} else {
				gen4699 = False
			}
			if True == gen4699 {
				__e.Return(V4801)
				return
			} else {
				gen4678 := Call(__e, ShenFunc(symcons_2), V4803)

				if True == gen4678 {
					gen4677 := MakeNative(func(__e Evaluator) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(ShenFunc(symshen_4insert_1runon), V4801, V4802, Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen4677, V4803)

					return

				} else {
					__e.Return(V4803)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-runon"), gen4700)

		gen4705 := MakeNative(func(__e Evaluator) {
			V4809 := __e.Get(1)
			_ = V4809
			gen4703 := Call(__e, ShenFunc(symelement_2), MakeString("."), V4809)

			gen4704 := Call(__e, ShenFunc(symnot), gen4703)

			if True == gen4704 {
				__e.Return(V4809)
				return
			} else {
				gen4702 := Call(__e, ShenFunc(symcons_2), V4809)

				if True == gen4702 {
					gen4701 := Call(__e, ShenFunc(symtl), V4809)

					__e.TailApply(ShenFunc(symshen_4strip_1pathname), gen4701)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.strip-pathname"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.strip-pathname"), gen4705)

		gen4732 := MakeNative(func(__e Evaluator) {
			V4813 := __e.Get(1)
			_ = V4813
			V4814 := __e.Get(2)
			_ = V4814
			V4815 := __e.Get(3)
			_ = V4815
			gen4731 := Call(__e, ShenFunc(symcons_2), V4813)

			if True == gen4731 {
				gen4706 := Call(__e, ShenFunc(symhd), V4813)

				gen4707 := Call(__e, ShenFunc(symcons), V4814, Nil)

				gen4708 := Call(__e, ShenFunc(symcons), gen4706, gen4707)

				Test := gen4708
				gen4709 := Call(__e, ShenFunc(symtl), V4813)

				gen4710 := Call(__e, ShenFunc(symhd), V4813)

				gen4711 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen4710)

				gen4712 := Call(__e, ShenFunc(symshen_4syntax), gen4709, gen4711, V4815)

				Action := gen4712
				gen4713 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen4713
				gen4714 := Call(__e, ShenFunc(symhd), V4813)

				gen4715 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen4714)

				gen4716 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				gen4717 := Call(__e, ShenFunc(symhd), V4813)

				gen4718 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen4717)

				gen4719 := Call(__e, ShenFunc(symcons), gen4718, Nil)

				gen4720 := Call(__e, ShenFunc(symcons), gen4716, gen4719)

				gen4721 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen4720)

				gen4722 := Call(__e, ShenFunc(symcons), gen4721, Nil)

				gen4723 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen4722)

				gen4724 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen4725 := Call(__e, ShenFunc(symcons), Action, gen4724)

				gen4726 := Call(__e, ShenFunc(symcons), gen4723, gen4725)

				gen4727 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen4726)

				gen4728 := Call(__e, ShenFunc(symcons), gen4727, Nil)

				gen4729 := Call(__e, ShenFunc(symcons), Test, gen4728)

				gen4730 := Call(__e, ShenFunc(symcons), gen4715, gen4729)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen4730)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.recursive_descent"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.recursive_descent"), gen4732)

		gen4759 := MakeNative(func(__e Evaluator) {
			V4819 := __e.Get(1)
			_ = V4819
			V4820 := __e.Get(2)
			_ = V4820
			V4821 := __e.Get(3)
			_ = V4821
			gen4758 := Call(__e, ShenFunc(symcons_2), V4819)

			if True == gen4758 {
				gen4733 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen4734 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4733)

				gen4735 := Call(__e, ShenFunc(symcons), gen4734, Nil)

				gen4736 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4735)

				Test := gen4736
				gen4737 := Call(__e, ShenFunc(symhd), V4819)

				gen4738 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), gen4737)

				gen4739 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen4740 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen4739)

				gen4741 := Call(__e, ShenFunc(symtl), V4819)

				gen4742 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen4743 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen4742)

				gen4744 := Call(__e, ShenFunc(symcons), V4820, Nil)

				gen4745 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4744)

				gen4746 := Call(__e, ShenFunc(symcons), gen4745, Nil)

				gen4747 := Call(__e, ShenFunc(symcons), gen4743, gen4746)

				gen4748 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4747)

				gen4749 := Call(__e, ShenFunc(symshen_4syntax), gen4741, gen4748, V4821)

				gen4750 := Call(__e, ShenFunc(symcons), gen4749, Nil)

				gen4751 := Call(__e, ShenFunc(symcons), gen4740, gen4750)

				gen4752 := Call(__e, ShenFunc(symcons), gen4738, gen4751)

				gen4753 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen4752)

				Action := gen4753
				gen4754 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen4754
				gen4755 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen4756 := Call(__e, ShenFunc(symcons), Action, gen4755)

				gen4757 := Call(__e, ShenFunc(symcons), Test, gen4756)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen4757)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.variable-match"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.variable-match"), gen4759)

		gen4762 := MakeNative(func(__e Evaluator) {
			V4831 := __e.Get(1)
			_ = V4831
			gen4761 := Call(__e, ShenFunc(symcons_2), V4831)

			if True == gen4761 {
				__e.Return(False)
				return
			} else {
				gen4760 := Call(__e, ShenFunc(symvariable_2), V4831)

				if True == gen4760 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.terminal?"), gen4762)

		gen4764 := MakeNative(func(__e Evaluator) {
			V4837 := __e.Get(1)
			_ = V4837
			gen4763 := Call(__e, ShenFunc(sym_a), V4837, MakeSymbol("_"))

			if True == gen4763 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.jump_stream?"), gen4764)

		gen4797 := MakeNative(func(__e Evaluator) {
			V4841 := __e.Get(1)
			_ = V4841
			V4842 := __e.Get(2)
			_ = V4842
			V4843 := __e.Get(3)
			_ = V4843
			gen4796 := Call(__e, ShenFunc(symcons_2), V4841)

			if True == gen4796 {
				gen4765 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen4766 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4765)

				gen4767 := Call(__e, ShenFunc(symcons), gen4766, Nil)

				gen4768 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4767)

				gen4769 := Call(__e, ShenFunc(symhd), V4841)

				gen4770 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen4771 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdhd"), gen4770)

				gen4772 := Call(__e, ShenFunc(symcons), gen4771, Nil)

				gen4773 := Call(__e, ShenFunc(symcons), gen4769, gen4772)

				gen4774 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen4773)

				gen4775 := Call(__e, ShenFunc(symcons), gen4774, Nil)

				gen4776 := Call(__e, ShenFunc(symcons), gen4768, gen4775)

				gen4777 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen4776)

				Test := gen4777
				gen4778 := Call(__e, ShenFunc(symgensym), MakeSymbol("NewStream"))

				NewStr := gen4778
				gen4779 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen4780 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen4779)

				gen4781 := Call(__e, ShenFunc(symcons), V4842, Nil)

				gen4782 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4781)

				gen4783 := Call(__e, ShenFunc(symcons), gen4782, Nil)

				gen4784 := Call(__e, ShenFunc(symcons), gen4780, gen4783)

				gen4785 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4784)

				gen4786 := Call(__e, ShenFunc(symtl), V4841)

				gen4787 := Call(__e, ShenFunc(symshen_4syntax), gen4786, NewStr, V4843)

				gen4788 := Call(__e, ShenFunc(symcons), gen4787, Nil)

				gen4789 := Call(__e, ShenFunc(symcons), gen4785, gen4788)

				gen4790 := Call(__e, ShenFunc(symcons), NewStr, gen4789)

				gen4791 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen4790)

				Action := gen4791
				gen4792 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen4792
				gen4793 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen4794 := Call(__e, ShenFunc(symcons), Action, gen4793)

				gen4795 := Call(__e, ShenFunc(symcons), Test, gen4794)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen4795)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.check_stream"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.check_stream"), gen4797)

		gen4816 := MakeNative(func(__e Evaluator) {
			V4847 := __e.Get(1)
			_ = V4847
			V4848 := __e.Get(2)
			_ = V4848
			V4849 := __e.Get(3)
			_ = V4849
			gen4815 := Call(__e, ShenFunc(symcons_2), V4847)

			if True == gen4815 {
				gen4798 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen4799 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4798)

				gen4800 := Call(__e, ShenFunc(symcons), gen4799, Nil)

				gen4801 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4800)

				Test := gen4801
				gen4802 := Call(__e, ShenFunc(symtl), V4847)

				gen4803 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen4804 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tlhd"), gen4803)

				gen4805 := Call(__e, ShenFunc(symcons), V4848, Nil)

				gen4806 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4805)

				gen4807 := Call(__e, ShenFunc(symcons), gen4806, Nil)

				gen4808 := Call(__e, ShenFunc(symcons), gen4804, gen4807)

				gen4809 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pair"), gen4808)

				gen4810 := Call(__e, ShenFunc(symshen_4syntax), gen4802, gen4809, V4849)

				Action := gen4810
				gen4811 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

				Else := gen4811
				gen4812 := Call(__e, ShenFunc(symcons), Else, Nil)

				gen4813 := Call(__e, ShenFunc(symcons), Action, gen4812)

				gen4814 := Call(__e, ShenFunc(symcons), Test, gen4813)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen4814)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.jump_stream"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.jump_stream"), gen4816)

		gen4824 := MakeNative(func(__e Evaluator) {
			V4851 := __e.Get(1)
			_ = V4851
			gen4823 := Call(__e, ShenFunc(sym_a), Nil, V4851)

			if True == gen4823 {
				__e.Return(Nil)
				return
			} else {
				gen4822 := Call(__e, ShenFunc(symshen_4grammar__symbol_2), V4851)

				if True == gen4822 {
					gen4820 := Call(__e, ShenFunc(symconcat), MakeSymbol("Parse_"), V4851)

					gen4821 := Call(__e, ShenFunc(symcons), gen4820, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.hdtl"), gen4821)

					return

				} else {
					gen4819 := Call(__e, ShenFunc(symvariable_2), V4851)

					if True == gen4819 {
						__e.TailApply(ShenFunc(symconcat), MakeSymbol("Parse_"), V4851)

						return
					} else {
						gen4818 := Call(__e, ShenFunc(symcons_2), V4851)

						if True == gen4818 {
							gen4817 := MakeNative(func(__e Evaluator) {
								Z := __e.Get(1)
								_ = Z
								__e.TailApply(ShenFunc(symshen_4semantics), Z)

								return
							}, 1)
							__e.TailApply(ShenFunc(symmap), gen4817, V4851)

							return

						} else {
							__e.Return(V4851)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.semantics"), gen4824)

		gen4826 := MakeNative(func(__e Evaluator) {
			V4854 := __e.Get(1)
			_ = V4854
			V4855 := __e.Get(2)
			_ = V4855
			gen4825 := Call(__e, ShenFunc(symcons), V4855, Nil)

			__e.TailApply(ShenFunc(symcons), V4854, gen4825)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pair"), gen4826)

		gen4828 := MakeNative(func(__e Evaluator) {
			V4857 := __e.Get(1)
			_ = V4857
			gen4827 := Call(__e, ShenFunc(symtl), V4857)

			__e.TailApply(ShenFunc(symhd), gen4827)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.hdtl"), gen4828)

		gen4830 := MakeNative(func(__e Evaluator) {
			V4859 := __e.Get(1)
			_ = V4859
			gen4829 := Call(__e, ShenFunc(symhd), V4859)

			__e.TailApply(ShenFunc(symhd), gen4829)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.hdhd"), gen4830)

		gen4832 := MakeNative(func(__e Evaluator) {
			V4861 := __e.Get(1)
			_ = V4861
			gen4831 := Call(__e, ShenFunc(symhd), V4861)

			__e.TailApply(ShenFunc(symtl), gen4831)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tlhd"), gen4832)

		gen4842 := MakeNative(func(__e Evaluator) {
			V4869 := __e.Get(1)
			_ = V4869
			gen4840 := Call(__e, ShenFunc(symcons_2), V4869)

			var gen4841 Obj
			if True == gen4840 {
				gen4837 := Call(__e, ShenFunc(symtl), V4869)

				gen4838 := Call(__e, ShenFunc(symcons_2), gen4837)

				var gen4839 Obj
				if True == gen4838 {
					gen4834 := Call(__e, ShenFunc(symtl), V4869)

					gen4835 := Call(__e, ShenFunc(symtl), gen4834)

					gen4836 := Call(__e, ShenFunc(sym_a), Nil, gen4835)

					if True == gen4836 {
						gen4839 = True
					} else {
						gen4839 = False
					}

				} else {
					gen4839 = False
				}
				if True == gen4839 {
					gen4841 = True
				} else {
					gen4841 = False
				}

			} else {
				gen4841 = False
			}
			if True == gen4841 {
				gen4833 := Call(__e, ShenFunc(symtl), V4869)

				__e.TailApply(ShenFunc(symhd), gen4833)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.snd-or-fail"), gen4842)

		gen4843 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeSymbol("shen.fail!"))
			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fail"), gen4843)

		gen4854 := MakeNative(func(__e Evaluator) {
			V4877 := __e.Get(1)
			_ = V4877
			gen4852 := Call(__e, ShenFunc(symcons_2), V4877)

			var gen4853 Obj
			if True == gen4852 {
				gen4849 := Call(__e, ShenFunc(symtl), V4877)

				gen4850 := Call(__e, ShenFunc(symcons_2), gen4849)

				var gen4851 Obj
				if True == gen4850 {
					gen4846 := Call(__e, ShenFunc(symtl), V4877)

					gen4847 := Call(__e, ShenFunc(symtl), gen4846)

					gen4848 := Call(__e, ShenFunc(sym_a), Nil, gen4847)

					if True == gen4848 {
						gen4851 = True
					} else {
						gen4851 = False
					}

				} else {
					gen4851 = False
				}
				if True == gen4851 {
					gen4853 = True
				} else {
					gen4853 = False
				}

			} else {
				gen4853 = False
			}
			if True == gen4853 {
				gen4844 := Call(__e, ShenFunc(symhd), V4877)

				gen4845 := Call(__e, ShenFunc(symcons), gen4844, Nil)

				__e.TailApply(ShenFunc(symcons), Nil, gen4845)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("<!>"), gen4854)

		gen4865 := MakeNative(func(__e Evaluator) {
			V4883 := __e.Get(1)
			_ = V4883
			gen4863 := Call(__e, ShenFunc(symcons_2), V4883)

			var gen4864 Obj
			if True == gen4863 {
				gen4860 := Call(__e, ShenFunc(symtl), V4883)

				gen4861 := Call(__e, ShenFunc(symcons_2), gen4860)

				var gen4862 Obj
				if True == gen4861 {
					gen4857 := Call(__e, ShenFunc(symtl), V4883)

					gen4858 := Call(__e, ShenFunc(symtl), gen4857)

					gen4859 := Call(__e, ShenFunc(sym_a), Nil, gen4858)

					if True == gen4859 {
						gen4862 = True
					} else {
						gen4862 = False
					}

				} else {
					gen4862 = False
				}
				if True == gen4862 {
					gen4864 = True
				} else {
					gen4864 = False
				}

			} else {
				gen4864 = False
			}
			if True == gen4864 {
				gen4855 := Call(__e, ShenFunc(symhd), V4883)

				gen4856 := Call(__e, ShenFunc(symcons), Nil, Nil)

				__e.TailApply(ShenFunc(symcons), gen4855, gen4856)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("<e>"))

				return
			}

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("<e>"), gen4865)

		return

	}, 0))
}
