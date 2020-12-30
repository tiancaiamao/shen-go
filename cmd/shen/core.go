package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen448 := MakeNative(func(__e Evaluator) {
			V152 := __e.Get(1)
			_ = V152
			V153 := __e.Get(2)
			_ = V153
			gen445 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5define_6), X)

				return
			}, 1)
			gen446 := Call(__e, ShenFunc(symcons), V152, V153)

			gen447 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4shen_1syntax_1error), V152, X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen445, gen446, gen447)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.shen->kl"), gen448)

		gen458 := MakeNative(func(__e Evaluator) {
			V160 := __e.Get(1)
			_ = V160
			V161 := __e.Get(2)
			_ = V161
			gen457 := Call(__e, ShenFunc(symcons_2), V161)

			if True == gen457 {
				gen451 := Call(__e, ShenFunc(symhd), V161)

				gen452 := Call(__e, ShenFunc(symshen_4next_150), MakeNumber(50), gen451)

				gen453 := Call(__e, ShenFunc(symshen_4app), gen452, MakeString("\n"), MakeSymbol("shen.a"))

				gen454 := Call(__e, ShenFunc(symcn), MakeString(" here:\n\n "), gen453)

				gen455 := Call(__e, ShenFunc(symshen_4app), V160, gen454, MakeSymbol("shen.a"))

				gen456 := Call(__e, ShenFunc(symcn), MakeString("syntax error in "), gen455)

				__e.TailApply(ShenFunc(symsimple_1error), gen456)

				return

			} else {
				gen449 := Call(__e, ShenFunc(symshen_4app), V160, MakeString("\n"), MakeSymbol("shen.a"))

				gen450 := Call(__e, ShenFunc(symcn), MakeString("syntax error in "), gen449)

				__e.TailApply(ShenFunc(symsimple_1error), gen450)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.shen-syntax-error"), gen458)

		gen490 := MakeNative(func(__e Evaluator) {
			V163 := __e.Get(1)
			_ = V163
			gen459 := Call(__e, ShenFunc(symshen_4_5name_6), V163)

			Parse__shen_4_5name_6 := gen459
			gen472 := Call(__e, ShenFunc(symfail))

			gen473 := Call(__e, ShenFunc(sym_a), gen472, Parse__shen_4_5name_6)

			gen474 := Call(__e, ShenFunc(symnot), gen473)

			var gen475 Obj
			if True == gen474 {
				gen460 := Call(__e, ShenFunc(symshen_4_5signature_6), Parse__shen_4_5name_6)

				Parse__shen_4_5signature_6 := gen460
				gen469 := Call(__e, ShenFunc(symfail))

				gen470 := Call(__e, ShenFunc(sym_a), gen469, Parse__shen_4_5signature_6)

				gen471 := Call(__e, ShenFunc(symnot), gen470)

				if True == gen471 {
					gen461 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5signature_6)

					Parse__shen_4_5rules_6 := gen461
					gen466 := Call(__e, ShenFunc(symfail))

					gen467 := Call(__e, ShenFunc(sym_a), gen466, Parse__shen_4_5rules_6)

					gen468 := Call(__e, ShenFunc(symnot), gen467)

					if True == gen468 {
						gen462 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

						gen463 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5name_6)

						gen464 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

						gen465 := Call(__e, ShenFunc(symshen_4compile__to__machine__code), gen463, gen464)

						gen475 = Call(__e, ShenFunc(symshen_4pair), gen462, gen465)

					} else {
						gen475 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen475 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen475 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen475
			gen488 := Call(__e, ShenFunc(symfail))

			gen489 := Call(__e, ShenFunc(sym_a), YaccParse, gen488)

			if True == gen489 {
				gen476 := Call(__e, ShenFunc(symshen_4_5name_6), V163)

				Parse__shen_4_5name_6 := gen476
				gen485 := Call(__e, ShenFunc(symfail))

				gen486 := Call(__e, ShenFunc(sym_a), gen485, Parse__shen_4_5name_6)

				gen487 := Call(__e, ShenFunc(symnot), gen486)

				if True == gen487 {
					gen477 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5name_6)

					Parse__shen_4_5rules_6 := gen477
					gen482 := Call(__e, ShenFunc(symfail))

					gen483 := Call(__e, ShenFunc(sym_a), gen482, Parse__shen_4_5rules_6)

					gen484 := Call(__e, ShenFunc(symnot), gen483)

					if True == gen484 {
						gen478 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

						gen479 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5name_6)

						gen480 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

						gen481 := Call(__e, ShenFunc(symshen_4compile__to__machine__code), gen479, gen480)

						__e.TailApply(ShenFunc(symshen_4pair), gen478, gen481)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<define>"), gen490)

		gen504 := MakeNative(func(__e Evaluator) {
			V165 := __e.Get(1)
			_ = V165
			gen502 := Call(__e, ShenFunc(symhd), V165)

			gen503 := Call(__e, ShenFunc(symcons_2), gen502)

			if True == gen503 {
				gen491 := Call(__e, ShenFunc(symshen_4hdhd), V165)

				Parse__X := gen491
				gen492 := Call(__e, ShenFunc(symshen_4tlhd), V165)

				gen493 := Call(__e, ShenFunc(symshen_4hdtl), V165)

				gen494 := Call(__e, ShenFunc(symshen_4pair), gen492, gen493)

				gen495 := Call(__e, ShenFunc(symhd), gen494)

				gen499 := Call(__e, ShenFunc(symsymbol_2), Parse__X)

				var gen500 Obj
				if True == gen499 {
					gen497 := Call(__e, ShenFunc(symshen_4sysfunc_2), Parse__X)

					gen498 := Call(__e, ShenFunc(symnot), gen497)

					if True == gen498 {
						gen500 = True
					} else {
						gen500 = False
					}

				} else {
					gen500 = False
				}
				var gen501 Obj
				if True == gen500 {
					gen501 = Parse__X
				} else {
					gen496 := Call(__e, ShenFunc(symshen_4app), Parse__X, MakeString(" is not a legitimate function name.\n"), MakeSymbol("shen.a"))

					gen501 = Call(__e, ShenFunc(symsimple_1error), gen496)

				}
				__e.TailApply(ShenFunc(symshen_4pair), gen495, gen501)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<name>"), gen504)

		gen508 := MakeNative(func(__e Evaluator) {
			V167 := __e.Get(1)
			_ = V167
			gen505 := Call(__e, ShenFunc(symintern), MakeString("shen"))

			gen506 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			gen507 := Call(__e, ShenFunc(symget), gen505, MakeSymbol("shen.external-symbols"), gen506)

			__e.TailApply(ShenFunc(symelement_2), V167, gen507)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sysfunc?"), gen508)

		gen533 := MakeNative(func(__e Evaluator) {
			V171 := __e.Get(1)
			_ = V171
			gen530 := Call(__e, ShenFunc(symhd), V171)

			gen531 := Call(__e, ShenFunc(symcons_2), gen530)

			var gen532 Obj
			if True == gen531 {
				gen528 := Call(__e, ShenFunc(symshen_4hdhd), V171)

				gen529 := Call(__e, ShenFunc(sym_a), MakeSymbol("{"), gen528)

				if True == gen529 {
					gen532 = True
				} else {
					gen532 = False
				}

			} else {
				gen532 = False
			}
			if True == gen532 {
				gen509 := Call(__e, ShenFunc(symshen_4tlhd), V171)

				gen510 := Call(__e, ShenFunc(symshen_4hdtl), V171)

				gen511 := Call(__e, ShenFunc(symshen_4pair), gen509, gen510)

				NewStream168 := gen511
				gen512 := Call(__e, ShenFunc(symshen_4_5signature_1help_6), NewStream168)

				Parse__shen_4_5signature_1help_6 := gen512
				gen525 := Call(__e, ShenFunc(symfail))

				gen526 := Call(__e, ShenFunc(sym_a), gen525, Parse__shen_4_5signature_1help_6)

				gen527 := Call(__e, ShenFunc(symnot), gen526)

				if True == gen527 {
					gen522 := Call(__e, ShenFunc(symhd), Parse__shen_4_5signature_1help_6)

					gen523 := Call(__e, ShenFunc(symcons_2), gen522)

					var gen524 Obj
					if True == gen523 {
						gen520 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5signature_1help_6)

						gen521 := Call(__e, ShenFunc(sym_a), MakeSymbol("}"), gen520)

						if True == gen521 {
							gen524 = True
						} else {
							gen524 = False
						}

					} else {
						gen524 = False
					}
					if True == gen524 {
						gen513 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5signature_1help_6)

						gen514 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen515 := Call(__e, ShenFunc(symshen_4pair), gen513, gen514)

						NewStream169 := gen515
						gen516 := Call(__e, ShenFunc(symhd), NewStream169)

						gen517 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen518 := Call(__e, ShenFunc(symshen_4curry_1type), gen517)

						gen519 := Call(__e, ShenFunc(symshen_4demodulate), gen518)

						__e.TailApply(ShenFunc(symshen_4pair), gen516, gen519)

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

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<signature>"), gen533)

		gen535 := MakeNative(func(__e Evaluator) {
			V173 := __e.Get(1)
			_ = V173
			gen534 := Call(__e, ShenFunc(symshen_4curry_1type_1h), V173)

			__e.TailApply(ShenFunc(symshen_4active_1cons), gen534)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-type"), gen535)

		gen564 := MakeNative(func(__e Evaluator) {
			V175 := __e.Get(1)
			_ = V175
			gen562 := Call(__e, ShenFunc(symcons_2), V175)

			var gen563 Obj
			if True == gen562 {
				gen559 := Call(__e, ShenFunc(symtl), V175)

				gen560 := Call(__e, ShenFunc(symcons_2), gen559)

				var gen561 Obj
				if True == gen560 {
					gen555 := Call(__e, ShenFunc(symtl), V175)

					gen556 := Call(__e, ShenFunc(symtl), gen555)

					gen557 := Call(__e, ShenFunc(symcons_2), gen556)

					var gen558 Obj
					if True == gen557 {
						gen550 := Call(__e, ShenFunc(symtl), V175)

						gen551 := Call(__e, ShenFunc(symtl), gen550)

						gen552 := Call(__e, ShenFunc(symtl), gen551)

						gen553 := Call(__e, ShenFunc(sym_a), Nil, gen552)

						var gen554 Obj
						if True == gen553 {
							gen547 := Call(__e, ShenFunc(symtl), V175)

							gen548 := Call(__e, ShenFunc(symhd), gen547)

							gen549 := Call(__e, ShenFunc(sym_a), gen548, MakeSymbol("bar!"))

							if True == gen549 {
								gen554 = True
							} else {
								gen554 = False
							}

						} else {
							gen554 = False
						}
						if True == gen554 {
							gen558 = True
						} else {
							gen558 = False
						}

					} else {
						gen558 = False
					}
					if True == gen558 {
						gen561 = True
					} else {
						gen561 = False
					}

				} else {
					gen561 = False
				}
				if True == gen561 {
					gen563 = True
				} else {
					gen563 = False
				}

			} else {
				gen563 = False
			}
			if True == gen563 {
				gen541 := Call(__e, ShenFunc(symhd), V175)

				gen542 := Call(__e, ShenFunc(symshen_4active_1cons), gen541)

				gen543 := Call(__e, ShenFunc(symtl), V175)

				gen544 := Call(__e, ShenFunc(symtl), gen543)

				gen545 := Call(__e, ShenFunc(symhd), gen544)

				gen546 := Call(__e, ShenFunc(symshen_4active_1cons), gen545)

				__e.TailApply(ShenFunc(symcons), gen542, gen546)

				return

			} else {
				gen540 := Call(__e, ShenFunc(symcons_2), V175)

				if True == gen540 {
					gen536 := Call(__e, ShenFunc(symhd), V175)

					gen537 := Call(__e, ShenFunc(symshen_4active_1cons), gen536)

					gen538 := Call(__e, ShenFunc(symtl), V175)

					gen539 := Call(__e, ShenFunc(symshen_4active_1cons), gen538)

					__e.TailApply(ShenFunc(symcons), gen537, gen539)

					return

				} else {
					__e.Return(V175)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.active-cons"), gen564)

		gen625 := MakeNative(func(__e Evaluator) {
			V177 := __e.Get(1)
			_ = V177
			gen623 := Call(__e, ShenFunc(symcons_2), V177)

			var gen624 Obj
			if True == gen623 {
				gen620 := Call(__e, ShenFunc(symtl), V177)

				gen621 := Call(__e, ShenFunc(symcons_2), gen620)

				var gen622 Obj
				if True == gen621 {
					gen616 := Call(__e, ShenFunc(symtl), V177)

					gen617 := Call(__e, ShenFunc(symhd), gen616)

					gen618 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen617)

					var gen619 Obj
					if True == gen618 {
						gen612 := Call(__e, ShenFunc(symtl), V177)

						gen613 := Call(__e, ShenFunc(symtl), gen612)

						gen614 := Call(__e, ShenFunc(symcons_2), gen613)

						var gen615 Obj
						if True == gen614 {
							gen607 := Call(__e, ShenFunc(symtl), V177)

							gen608 := Call(__e, ShenFunc(symtl), gen607)

							gen609 := Call(__e, ShenFunc(symtl), gen608)

							gen610 := Call(__e, ShenFunc(symcons_2), gen609)

							var gen611 Obj
							if True == gen610 {
								gen602 := Call(__e, ShenFunc(symtl), V177)

								gen603 := Call(__e, ShenFunc(symtl), gen602)

								gen604 := Call(__e, ShenFunc(symtl), gen603)

								gen605 := Call(__e, ShenFunc(symhd), gen604)

								gen606 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen605)

								if True == gen606 {
									gen611 = True
								} else {
									gen611 = False
								}

							} else {
								gen611 = False
							}
							if True == gen611 {
								gen615 = True
							} else {
								gen615 = False
							}

						} else {
							gen615 = False
						}
						if True == gen615 {
							gen619 = True
						} else {
							gen619 = False
						}

					} else {
						gen619 = False
					}
					if True == gen619 {
						gen622 = True
					} else {
						gen622 = False
					}

				} else {
					gen622 = False
				}
				if True == gen622 {
					gen624 = True
				} else {
					gen624 = False
				}

			} else {
				gen624 = False
			}
			if True == gen624 {
				gen596 := Call(__e, ShenFunc(symhd), V177)

				gen597 := Call(__e, ShenFunc(symtl), V177)

				gen598 := Call(__e, ShenFunc(symtl), gen597)

				gen599 := Call(__e, ShenFunc(symcons), gen598, Nil)

				gen600 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen599)

				gen601 := Call(__e, ShenFunc(symcons), gen596, gen600)

				__e.TailApply(ShenFunc(symshen_4curry_1type_1h), gen601)

				return

			} else {
				gen594 := Call(__e, ShenFunc(symcons_2), V177)

				var gen595 Obj
				if True == gen594 {
					gen591 := Call(__e, ShenFunc(symtl), V177)

					gen592 := Call(__e, ShenFunc(symcons_2), gen591)

					var gen593 Obj
					if True == gen592 {
						gen587 := Call(__e, ShenFunc(symtl), V177)

						gen588 := Call(__e, ShenFunc(symhd), gen587)

						gen589 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), gen588)

						var gen590 Obj
						if True == gen589 {
							gen583 := Call(__e, ShenFunc(symtl), V177)

							gen584 := Call(__e, ShenFunc(symtl), gen583)

							gen585 := Call(__e, ShenFunc(symcons_2), gen584)

							var gen586 Obj
							if True == gen585 {
								gen578 := Call(__e, ShenFunc(symtl), V177)

								gen579 := Call(__e, ShenFunc(symtl), gen578)

								gen580 := Call(__e, ShenFunc(symtl), gen579)

								gen581 := Call(__e, ShenFunc(symcons_2), gen580)

								var gen582 Obj
								if True == gen581 {
									gen573 := Call(__e, ShenFunc(symtl), V177)

									gen574 := Call(__e, ShenFunc(symtl), gen573)

									gen575 := Call(__e, ShenFunc(symtl), gen574)

									gen576 := Call(__e, ShenFunc(symhd), gen575)

									gen577 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), gen576)

									if True == gen577 {
										gen582 = True
									} else {
										gen582 = False
									}

								} else {
									gen582 = False
								}
								if True == gen582 {
									gen586 = True
								} else {
									gen586 = False
								}

							} else {
								gen586 = False
							}
							if True == gen586 {
								gen590 = True
							} else {
								gen590 = False
							}

						} else {
							gen590 = False
						}
						if True == gen590 {
							gen593 = True
						} else {
							gen593 = False
						}

					} else {
						gen593 = False
					}
					if True == gen593 {
						gen595 = True
					} else {
						gen595 = False
					}

				} else {
					gen595 = False
				}
				if True == gen595 {
					gen567 := Call(__e, ShenFunc(symhd), V177)

					gen568 := Call(__e, ShenFunc(symtl), V177)

					gen569 := Call(__e, ShenFunc(symtl), gen568)

					gen570 := Call(__e, ShenFunc(symcons), gen569, Nil)

					gen571 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen570)

					gen572 := Call(__e, ShenFunc(symcons), gen567, gen571)

					__e.TailApply(ShenFunc(symshen_4curry_1type_1h), gen572)

					return

				} else {
					gen566 := Call(__e, ShenFunc(symcons_2), V177)

					if True == gen566 {
						gen565 := MakeNative(func(__e Evaluator) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(ShenFunc(symshen_4curry_1type_1h), Z)

							return
						}, 1)
						__e.TailApply(ShenFunc(symmap), gen565, V177)

						return

					} else {
						__e.Return(V177)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-type-h"), gen625)

		gen651 := MakeNative(func(__e Evaluator) {
			V179 := __e.Get(1)
			_ = V179
			gen641 := Call(__e, ShenFunc(symhd), V179)

			gen642 := Call(__e, ShenFunc(symcons_2), gen641)

			var gen643 Obj
			if True == gen642 {
				gen626 := Call(__e, ShenFunc(symshen_4hdhd), V179)

				Parse__X := gen626
				gen627 := Call(__e, ShenFunc(symshen_4tlhd), V179)

				gen628 := Call(__e, ShenFunc(symshen_4hdtl), V179)

				gen629 := Call(__e, ShenFunc(symshen_4pair), gen627, gen628)

				gen630 := Call(__e, ShenFunc(symshen_4_5signature_1help_6), gen629)

				Parse__shen_4_5signature_1help_6 := gen630
				gen638 := Call(__e, ShenFunc(symfail))

				gen639 := Call(__e, ShenFunc(sym_a), gen638, Parse__shen_4_5signature_1help_6)

				gen640 := Call(__e, ShenFunc(symnot), gen639)

				if True == gen640 {
					gen634 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), Nil)

					gen635 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen634)

					gen636 := Call(__e, ShenFunc(symelement_2), Parse__X, gen635)

					gen637 := Call(__e, ShenFunc(symnot), gen636)

					if True == gen637 {
						gen631 := Call(__e, ShenFunc(symhd), Parse__shen_4_5signature_1help_6)

						gen632 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen633 := Call(__e, ShenFunc(symcons), Parse__X, gen632)

						gen643 = Call(__e, ShenFunc(symshen_4pair), gen631, gen633)

					} else {
						gen643 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen643 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen643 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen643
			gen649 := Call(__e, ShenFunc(symfail))

			gen650 := Call(__e, ShenFunc(sym_a), YaccParse, gen649)

			if True == gen650 {
				gen644 := Call(__e, ShenFunc(sym_5e_6), V179)

				Parse___5e_6 := gen644
				gen646 := Call(__e, ShenFunc(symfail))

				gen647 := Call(__e, ShenFunc(sym_a), gen646, Parse___5e_6)

				gen648 := Call(__e, ShenFunc(symnot), gen647)

				if True == gen648 {
					gen645 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen645, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<signature-help>"), gen651)

		gen676 := MakeNative(func(__e Evaluator) {
			V181 := __e.Get(1)
			_ = V181
			gen652 := Call(__e, ShenFunc(symshen_4_5rule_6), V181)

			Parse__shen_4_5rule_6 := gen652
			gen662 := Call(__e, ShenFunc(symfail))

			gen663 := Call(__e, ShenFunc(sym_a), gen662, Parse__shen_4_5rule_6)

			gen664 := Call(__e, ShenFunc(symnot), gen663)

			var gen665 Obj
			if True == gen664 {
				gen653 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5rule_6)

				Parse__shen_4_5rules_6 := gen653
				gen659 := Call(__e, ShenFunc(symfail))

				gen660 := Call(__e, ShenFunc(sym_a), gen659, Parse__shen_4_5rules_6)

				gen661 := Call(__e, ShenFunc(symnot), gen660)

				if True == gen661 {
					gen654 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

					gen655 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen656 := Call(__e, ShenFunc(symshen_4linearise), gen655)

					gen657 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

					gen658 := Call(__e, ShenFunc(symcons), gen656, gen657)

					gen665 = Call(__e, ShenFunc(symshen_4pair), gen654, gen658)

				} else {
					gen665 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen665 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen665
			gen674 := Call(__e, ShenFunc(symfail))

			gen675 := Call(__e, ShenFunc(sym_a), YaccParse, gen674)

			if True == gen675 {
				gen666 := Call(__e, ShenFunc(symshen_4_5rule_6), V181)

				Parse__shen_4_5rule_6 := gen666
				gen671 := Call(__e, ShenFunc(symfail))

				gen672 := Call(__e, ShenFunc(sym_a), gen671, Parse__shen_4_5rule_6)

				gen673 := Call(__e, ShenFunc(symnot), gen672)

				if True == gen673 {
					gen667 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rule_6)

					gen668 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen669 := Call(__e, ShenFunc(symshen_4linearise), gen668)

					gen670 := Call(__e, ShenFunc(symcons), gen669, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen667, gen670)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rules>"), gen676)

		gen806 := MakeNative(func(__e Evaluator) {
			V189 := __e.Get(1)
			_ = V189
			gen677 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

			Parse__shen_4_5patterns_6 := gen677
			gen711 := Call(__e, ShenFunc(symfail))

			gen712 := Call(__e, ShenFunc(sym_a), gen711, Parse__shen_4_5patterns_6)

			gen713 := Call(__e, ShenFunc(symnot), gen712)

			var gen714 Obj
			if True == gen713 {
				gen708 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

				gen709 := Call(__e, ShenFunc(symcons_2), gen708)

				var gen710 Obj
				if True == gen709 {
					gen706 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

					gen707 := Call(__e, ShenFunc(sym_a), MakeSymbol("->"), gen706)

					if True == gen707 {
						gen710 = True
					} else {
						gen710 = False
					}

				} else {
					gen710 = False
				}
				if True == gen710 {
					gen678 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

					gen679 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

					gen680 := Call(__e, ShenFunc(symshen_4pair), gen678, gen679)

					NewStream182 := gen680
					gen681 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream182)

					Parse__shen_4_5action_6 := gen681
					gen703 := Call(__e, ShenFunc(symfail))

					gen704 := Call(__e, ShenFunc(sym_a), gen703, Parse__shen_4_5action_6)

					gen705 := Call(__e, ShenFunc(symnot), gen704)

					if True == gen705 {
						gen700 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

						gen701 := Call(__e, ShenFunc(symcons_2), gen700)

						var gen702 Obj
						if True == gen701 {
							gen698 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5action_6)

							gen699 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen698)

							if True == gen699 {
								gen702 = True
							} else {
								gen702 = False
							}

						} else {
							gen702 = False
						}
						if True == gen702 {
							gen682 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5action_6)

							gen683 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

							gen684 := Call(__e, ShenFunc(symshen_4pair), gen682, gen683)

							NewStream183 := gen684
							gen685 := Call(__e, ShenFunc(symshen_4_5guard_6), NewStream183)

							Parse__shen_4_5guard_6 := gen685
							gen695 := Call(__e, ShenFunc(symfail))

							gen696 := Call(__e, ShenFunc(sym_a), gen695, Parse__shen_4_5guard_6)

							gen697 := Call(__e, ShenFunc(symnot), gen696)

							if True == gen697 {
								gen686 := Call(__e, ShenFunc(symhd), Parse__shen_4_5guard_6)

								gen687 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

								gen688 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5guard_6)

								gen689 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

								gen690 := Call(__e, ShenFunc(symcons), gen689, Nil)

								gen691 := Call(__e, ShenFunc(symcons), gen688, gen690)

								gen692 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen691)

								gen693 := Call(__e, ShenFunc(symcons), gen692, Nil)

								gen694 := Call(__e, ShenFunc(symcons), gen687, gen693)

								gen714 = Call(__e, ShenFunc(symshen_4pair), gen686, gen694)

							} else {
								gen714 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen714 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen714 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen714 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen714 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen714
			gen804 := Call(__e, ShenFunc(symfail))

			gen805 := Call(__e, ShenFunc(sym_a), YaccParse, gen804)

			if True == gen805 {
				gen715 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

				Parse__shen_4_5patterns_6 := gen715
				gen733 := Call(__e, ShenFunc(symfail))

				gen734 := Call(__e, ShenFunc(sym_a), gen733, Parse__shen_4_5patterns_6)

				gen735 := Call(__e, ShenFunc(symnot), gen734)

				var gen736 Obj
				if True == gen735 {
					gen730 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

					gen731 := Call(__e, ShenFunc(symcons_2), gen730)

					var gen732 Obj
					if True == gen731 {
						gen728 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

						gen729 := Call(__e, ShenFunc(sym_a), MakeSymbol("->"), gen728)

						if True == gen729 {
							gen732 = True
						} else {
							gen732 = False
						}

					} else {
						gen732 = False
					}
					if True == gen732 {
						gen716 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

						gen717 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

						gen718 := Call(__e, ShenFunc(symshen_4pair), gen716, gen717)

						NewStream184 := gen718
						gen719 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream184)

						Parse__shen_4_5action_6 := gen719
						gen725 := Call(__e, ShenFunc(symfail))

						gen726 := Call(__e, ShenFunc(sym_a), gen725, Parse__shen_4_5action_6)

						gen727 := Call(__e, ShenFunc(symnot), gen726)

						if True == gen727 {
							gen720 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

							gen721 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

							gen722 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

							gen723 := Call(__e, ShenFunc(symcons), gen722, Nil)

							gen724 := Call(__e, ShenFunc(symcons), gen721, gen723)

							gen736 = Call(__e, ShenFunc(symshen_4pair), gen720, gen724)

						} else {
							gen736 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen736 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen736 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen736
				gen802 := Call(__e, ShenFunc(symfail))

				gen803 := Call(__e, ShenFunc(sym_a), YaccParse, gen802)

				if True == gen803 {
					gen737 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

					Parse__shen_4_5patterns_6 := gen737
					gen773 := Call(__e, ShenFunc(symfail))

					gen774 := Call(__e, ShenFunc(sym_a), gen773, Parse__shen_4_5patterns_6)

					gen775 := Call(__e, ShenFunc(symnot), gen774)

					var gen776 Obj
					if True == gen775 {
						gen770 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

						gen771 := Call(__e, ShenFunc(symcons_2), gen770)

						var gen772 Obj
						if True == gen771 {
							gen768 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

							gen769 := Call(__e, ShenFunc(sym_a), MakeSymbol("<-"), gen768)

							if True == gen769 {
								gen772 = True
							} else {
								gen772 = False
							}

						} else {
							gen772 = False
						}
						if True == gen772 {
							gen738 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

							gen739 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

							gen740 := Call(__e, ShenFunc(symshen_4pair), gen738, gen739)

							NewStream185 := gen740
							gen741 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream185)

							Parse__shen_4_5action_6 := gen741
							gen765 := Call(__e, ShenFunc(symfail))

							gen766 := Call(__e, ShenFunc(sym_a), gen765, Parse__shen_4_5action_6)

							gen767 := Call(__e, ShenFunc(symnot), gen766)

							if True == gen767 {
								gen762 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

								gen763 := Call(__e, ShenFunc(symcons_2), gen762)

								var gen764 Obj
								if True == gen763 {
									gen760 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5action_6)

									gen761 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen760)

									if True == gen761 {
										gen764 = True
									} else {
										gen764 = False
									}

								} else {
									gen764 = False
								}
								if True == gen764 {
									gen742 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5action_6)

									gen743 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

									gen744 := Call(__e, ShenFunc(symshen_4pair), gen742, gen743)

									NewStream186 := gen744
									gen745 := Call(__e, ShenFunc(symshen_4_5guard_6), NewStream186)

									Parse__shen_4_5guard_6 := gen745
									gen757 := Call(__e, ShenFunc(symfail))

									gen758 := Call(__e, ShenFunc(sym_a), gen757, Parse__shen_4_5guard_6)

									gen759 := Call(__e, ShenFunc(symnot), gen758)

									if True == gen759 {
										gen746 := Call(__e, ShenFunc(symhd), Parse__shen_4_5guard_6)

										gen747 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

										gen748 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5guard_6)

										gen749 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

										gen750 := Call(__e, ShenFunc(symcons), gen749, Nil)

										gen751 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.choicepoint!"), gen750)

										gen752 := Call(__e, ShenFunc(symcons), gen751, Nil)

										gen753 := Call(__e, ShenFunc(symcons), gen748, gen752)

										gen754 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen753)

										gen755 := Call(__e, ShenFunc(symcons), gen754, Nil)

										gen756 := Call(__e, ShenFunc(symcons), gen747, gen755)

										gen776 = Call(__e, ShenFunc(symshen_4pair), gen746, gen756)

									} else {
										gen776 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen776 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen776 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen776 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen776 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen776
					gen800 := Call(__e, ShenFunc(symfail))

					gen801 := Call(__e, ShenFunc(sym_a), YaccParse, gen800)

					if True == gen801 {
						gen777 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

						Parse__shen_4_5patterns_6 := gen777
						gen797 := Call(__e, ShenFunc(symfail))

						gen798 := Call(__e, ShenFunc(sym_a), gen797, Parse__shen_4_5patterns_6)

						gen799 := Call(__e, ShenFunc(symnot), gen798)

						if True == gen799 {
							gen794 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

							gen795 := Call(__e, ShenFunc(symcons_2), gen794)

							var gen796 Obj
							if True == gen795 {
								gen792 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

								gen793 := Call(__e, ShenFunc(sym_a), MakeSymbol("<-"), gen792)

								if True == gen793 {
									gen796 = True
								} else {
									gen796 = False
								}

							} else {
								gen796 = False
							}
							if True == gen796 {
								gen778 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

								gen779 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

								gen780 := Call(__e, ShenFunc(symshen_4pair), gen778, gen779)

								NewStream187 := gen780
								gen781 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream187)

								Parse__shen_4_5action_6 := gen781
								gen789 := Call(__e, ShenFunc(symfail))

								gen790 := Call(__e, ShenFunc(sym_a), gen789, Parse__shen_4_5action_6)

								gen791 := Call(__e, ShenFunc(symnot), gen790)

								if True == gen791 {
									gen782 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

									gen783 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

									gen784 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

									gen785 := Call(__e, ShenFunc(symcons), gen784, Nil)

									gen786 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.choicepoint!"), gen785)

									gen787 := Call(__e, ShenFunc(symcons), gen786, Nil)

									gen788 := Call(__e, ShenFunc(symcons), gen783, gen787)

									__e.TailApply(ShenFunc(symshen_4pair), gen782, gen788)

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
						__e.Return(YaccParse)
						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rule>"), gen806)

		gen808 := MakeNative(func(__e Evaluator) {
			V192 := __e.Get(1)
			_ = V192
			V193 := __e.Get(2)
			_ = V193
			gen807 := Call(__e, V192, V193)

			if True == gen807 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V193)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fail_if"), gen808)

		gen811 := MakeNative(func(__e Evaluator) {
			V199 := __e.Get(1)
			_ = V199
			gen809 := Call(__e, ShenFunc(symfail))

			gen810 := Call(__e, ShenFunc(sym_a), V199, gen809)

			if True == gen810 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.succeeds?"), gen811)

		gen813 := MakeNative(func(__e Evaluator) {
			V202 := __e.Get(1)
			_ = V202
			V203 := __e.Get(2)
			_ = V203
			gen812 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*custom-pattern-compiler*"))

			f28 := gen812
			__e.TailApply(f28, V202, V203)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.custom-pattern-compiler"), gen813)

		gen815 := MakeNative(func(__e Evaluator) {
			V205 := __e.Get(1)
			_ = V205
			gen814 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*custom-pattern-reducer*"))

			f29 := gen814
			__e.TailApply(f29, V205)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.custom-pattern-reducer"), gen815)

		gen836 := MakeNative(func(__e Evaluator) {
			V207 := __e.Get(1)
			_ = V207
			gen816 := Call(__e, ShenFunc(symshen_4_5pattern_6), V207)

			Parse__shen_4_5pattern_6 := gen816
			gen825 := Call(__e, ShenFunc(symfail))

			gen826 := Call(__e, ShenFunc(sym_a), gen825, Parse__shen_4_5pattern_6)

			gen827 := Call(__e, ShenFunc(symnot), gen826)

			var gen828 Obj
			if True == gen827 {
				gen817 := Call(__e, ShenFunc(symshen_4_5patterns_6), Parse__shen_4_5pattern_6)

				Parse__shen_4_5patterns_6 := gen817
				gen822 := Call(__e, ShenFunc(symfail))

				gen823 := Call(__e, ShenFunc(sym_a), gen822, Parse__shen_4_5patterns_6)

				gen824 := Call(__e, ShenFunc(symnot), gen823)

				if True == gen824 {
					gen818 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

					gen819 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

					gen820 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

					gen821 := Call(__e, ShenFunc(symcons), gen819, gen820)

					gen828 = Call(__e, ShenFunc(symshen_4pair), gen818, gen821)

				} else {
					gen828 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen828 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen828
			gen834 := Call(__e, ShenFunc(symfail))

			gen835 := Call(__e, ShenFunc(sym_a), YaccParse, gen834)

			if True == gen835 {
				gen829 := Call(__e, ShenFunc(sym_5e_6), V207)

				Parse___5e_6 := gen829
				gen831 := Call(__e, ShenFunc(symfail))

				gen832 := Call(__e, ShenFunc(sym_a), gen831, Parse___5e_6)

				gen833 := Call(__e, ShenFunc(symnot), gen832)

				if True == gen833 {
					gen830 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen830, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<patterns>"), gen836)

		gen1078 := MakeNative(func(__e Evaluator) {
			V220 := __e.Get(1)
			_ = V220
			gen876 := Call(__e, ShenFunc(symhd), V220)

			gen877 := Call(__e, ShenFunc(symcons_2), gen876)

			var gen878 Obj
			if True == gen877 {
				gen874 := Call(__e, ShenFunc(symshen_4hdhd), V220)

				gen875 := Call(__e, ShenFunc(symcons_2), gen874)

				if True == gen875 {
					gen878 = True
				} else {
					gen878 = False
				}

			} else {
				gen878 = False
			}
			var gen879 Obj
			if True == gen878 {
				gen868 := Call(__e, ShenFunc(symshen_4hdhd), V220)

				gen869 := Call(__e, ShenFunc(symshen_4hdtl), V220)

				gen870 := Call(__e, ShenFunc(symshen_4pair), gen868, gen869)

				gen871 := Call(__e, ShenFunc(symhd), gen870)

				gen872 := Call(__e, ShenFunc(symcons_2), gen871)

				var gen873 Obj
				if True == gen872 {
					gen863 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen864 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen865 := Call(__e, ShenFunc(symshen_4pair), gen863, gen864)

					gen866 := Call(__e, ShenFunc(symshen_4hdhd), gen865)

					gen867 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), gen866)

					if True == gen867 {
						gen873 = True
					} else {
						gen873 = False
					}

				} else {
					gen873 = False
				}
				if True == gen873 {
					gen837 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen838 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen839 := Call(__e, ShenFunc(symshen_4pair), gen837, gen838)

					gen840 := Call(__e, ShenFunc(symshen_4tlhd), gen839)

					gen841 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen842 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen843 := Call(__e, ShenFunc(symshen_4pair), gen841, gen842)

					gen844 := Call(__e, ShenFunc(symshen_4hdtl), gen843)

					gen845 := Call(__e, ShenFunc(symshen_4pair), gen840, gen844)

					NewStream209 := gen845
					gen846 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream209)

					Parse__shen_4_5pattern1_6 := gen846
					gen860 := Call(__e, ShenFunc(symfail))

					gen861 := Call(__e, ShenFunc(sym_a), gen860, Parse__shen_4_5pattern1_6)

					gen862 := Call(__e, ShenFunc(symnot), gen861)

					if True == gen862 {
						gen847 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

						Parse__shen_4_5pattern2_6 := gen847
						gen857 := Call(__e, ShenFunc(symfail))

						gen858 := Call(__e, ShenFunc(sym_a), gen857, Parse__shen_4_5pattern2_6)

						gen859 := Call(__e, ShenFunc(symnot), gen858)

						if True == gen859 {
							gen848 := Call(__e, ShenFunc(symshen_4tlhd), V220)

							gen849 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen850 := Call(__e, ShenFunc(symshen_4pair), gen848, gen849)

							gen851 := Call(__e, ShenFunc(symhd), gen850)

							gen852 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

							gen853 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

							gen854 := Call(__e, ShenFunc(symcons), gen853, Nil)

							gen855 := Call(__e, ShenFunc(symcons), gen852, gen854)

							gen856 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen855)

							gen879 = Call(__e, ShenFunc(symshen_4pair), gen851, gen856)

						} else {
							gen879 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen879 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen879 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen879 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen879
			gen1076 := Call(__e, ShenFunc(symfail))

			gen1077 := Call(__e, ShenFunc(sym_a), YaccParse, gen1076)

			if True == gen1077 {
				gen919 := Call(__e, ShenFunc(symhd), V220)

				gen920 := Call(__e, ShenFunc(symcons_2), gen919)

				var gen921 Obj
				if True == gen920 {
					gen917 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen918 := Call(__e, ShenFunc(symcons_2), gen917)

					if True == gen918 {
						gen921 = True
					} else {
						gen921 = False
					}

				} else {
					gen921 = False
				}
				var gen922 Obj
				if True == gen921 {
					gen911 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen912 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen913 := Call(__e, ShenFunc(symshen_4pair), gen911, gen912)

					gen914 := Call(__e, ShenFunc(symhd), gen913)

					gen915 := Call(__e, ShenFunc(symcons_2), gen914)

					var gen916 Obj
					if True == gen915 {
						gen906 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen907 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen908 := Call(__e, ShenFunc(symshen_4pair), gen906, gen907)

						gen909 := Call(__e, ShenFunc(symshen_4hdhd), gen908)

						gen910 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen909)

						if True == gen910 {
							gen916 = True
						} else {
							gen916 = False
						}

					} else {
						gen916 = False
					}
					if True == gen916 {
						gen880 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen881 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen882 := Call(__e, ShenFunc(symshen_4pair), gen880, gen881)

						gen883 := Call(__e, ShenFunc(symshen_4tlhd), gen882)

						gen884 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen885 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen886 := Call(__e, ShenFunc(symshen_4pair), gen884, gen885)

						gen887 := Call(__e, ShenFunc(symshen_4hdtl), gen886)

						gen888 := Call(__e, ShenFunc(symshen_4pair), gen883, gen887)

						NewStream211 := gen888
						gen889 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream211)

						Parse__shen_4_5pattern1_6 := gen889
						gen903 := Call(__e, ShenFunc(symfail))

						gen904 := Call(__e, ShenFunc(sym_a), gen903, Parse__shen_4_5pattern1_6)

						gen905 := Call(__e, ShenFunc(symnot), gen904)

						if True == gen905 {
							gen890 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

							Parse__shen_4_5pattern2_6 := gen890
							gen900 := Call(__e, ShenFunc(symfail))

							gen901 := Call(__e, ShenFunc(sym_a), gen900, Parse__shen_4_5pattern2_6)

							gen902 := Call(__e, ShenFunc(symnot), gen901)

							if True == gen902 {
								gen891 := Call(__e, ShenFunc(symshen_4tlhd), V220)

								gen892 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen893 := Call(__e, ShenFunc(symshen_4pair), gen891, gen892)

								gen894 := Call(__e, ShenFunc(symhd), gen893)

								gen895 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

								gen896 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

								gen897 := Call(__e, ShenFunc(symcons), gen896, Nil)

								gen898 := Call(__e, ShenFunc(symcons), gen895, gen897)

								gen899 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen898)

								gen922 = Call(__e, ShenFunc(symshen_4pair), gen894, gen899)

							} else {
								gen922 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen922 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen922 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen922 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen922
				gen1074 := Call(__e, ShenFunc(symfail))

				gen1075 := Call(__e, ShenFunc(sym_a), YaccParse, gen1074)

				if True == gen1075 {
					gen962 := Call(__e, ShenFunc(symhd), V220)

					gen963 := Call(__e, ShenFunc(symcons_2), gen962)

					var gen964 Obj
					if True == gen963 {
						gen960 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen961 := Call(__e, ShenFunc(symcons_2), gen960)

						if True == gen961 {
							gen964 = True
						} else {
							gen964 = False
						}

					} else {
						gen964 = False
					}
					var gen965 Obj
					if True == gen964 {
						gen954 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen955 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen956 := Call(__e, ShenFunc(symshen_4pair), gen954, gen955)

						gen957 := Call(__e, ShenFunc(symhd), gen956)

						gen958 := Call(__e, ShenFunc(symcons_2), gen957)

						var gen959 Obj
						if True == gen958 {
							gen949 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen950 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen951 := Call(__e, ShenFunc(symshen_4pair), gen949, gen950)

							gen952 := Call(__e, ShenFunc(symshen_4hdhd), gen951)

							gen953 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), gen952)

							if True == gen953 {
								gen959 = True
							} else {
								gen959 = False
							}

						} else {
							gen959 = False
						}
						if True == gen959 {
							gen923 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen924 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen925 := Call(__e, ShenFunc(symshen_4pair), gen923, gen924)

							gen926 := Call(__e, ShenFunc(symshen_4tlhd), gen925)

							gen927 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen928 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen929 := Call(__e, ShenFunc(symshen_4pair), gen927, gen928)

							gen930 := Call(__e, ShenFunc(symshen_4hdtl), gen929)

							gen931 := Call(__e, ShenFunc(symshen_4pair), gen926, gen930)

							NewStream213 := gen931
							gen932 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream213)

							Parse__shen_4_5pattern1_6 := gen932
							gen946 := Call(__e, ShenFunc(symfail))

							gen947 := Call(__e, ShenFunc(sym_a), gen946, Parse__shen_4_5pattern1_6)

							gen948 := Call(__e, ShenFunc(symnot), gen947)

							if True == gen948 {
								gen933 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

								Parse__shen_4_5pattern2_6 := gen933
								gen943 := Call(__e, ShenFunc(symfail))

								gen944 := Call(__e, ShenFunc(sym_a), gen943, Parse__shen_4_5pattern2_6)

								gen945 := Call(__e, ShenFunc(symnot), gen944)

								if True == gen945 {
									gen934 := Call(__e, ShenFunc(symshen_4tlhd), V220)

									gen935 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen936 := Call(__e, ShenFunc(symshen_4pair), gen934, gen935)

									gen937 := Call(__e, ShenFunc(symhd), gen936)

									gen938 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

									gen939 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

									gen940 := Call(__e, ShenFunc(symcons), gen939, Nil)

									gen941 := Call(__e, ShenFunc(symcons), gen938, gen940)

									gen942 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen941)

									gen965 = Call(__e, ShenFunc(symshen_4pair), gen937, gen942)

								} else {
									gen965 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen965 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen965 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen965 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen965
					gen1072 := Call(__e, ShenFunc(symfail))

					gen1073 := Call(__e, ShenFunc(sym_a), YaccParse, gen1072)

					if True == gen1073 {
						gen1005 := Call(__e, ShenFunc(symhd), V220)

						gen1006 := Call(__e, ShenFunc(symcons_2), gen1005)

						var gen1007 Obj
						if True == gen1006 {
							gen1003 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen1004 := Call(__e, ShenFunc(symcons_2), gen1003)

							if True == gen1004 {
								gen1007 = True
							} else {
								gen1007 = False
							}

						} else {
							gen1007 = False
						}
						var gen1008 Obj
						if True == gen1007 {
							gen997 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen998 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen999 := Call(__e, ShenFunc(symshen_4pair), gen997, gen998)

							gen1000 := Call(__e, ShenFunc(symhd), gen999)

							gen1001 := Call(__e, ShenFunc(symcons_2), gen1000)

							var gen1002 Obj
							if True == gen1001 {
								gen992 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen993 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen994 := Call(__e, ShenFunc(symshen_4pair), gen992, gen993)

								gen995 := Call(__e, ShenFunc(symshen_4hdhd), gen994)

								gen996 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen995)

								if True == gen996 {
									gen1002 = True
								} else {
									gen1002 = False
								}

							} else {
								gen1002 = False
							}
							if True == gen1002 {
								gen966 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen967 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen968 := Call(__e, ShenFunc(symshen_4pair), gen966, gen967)

								gen969 := Call(__e, ShenFunc(symshen_4tlhd), gen968)

								gen970 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen971 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen972 := Call(__e, ShenFunc(symshen_4pair), gen970, gen971)

								gen973 := Call(__e, ShenFunc(symshen_4hdtl), gen972)

								gen974 := Call(__e, ShenFunc(symshen_4pair), gen969, gen973)

								NewStream215 := gen974
								gen975 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream215)

								Parse__shen_4_5pattern1_6 := gen975
								gen989 := Call(__e, ShenFunc(symfail))

								gen990 := Call(__e, ShenFunc(sym_a), gen989, Parse__shen_4_5pattern1_6)

								gen991 := Call(__e, ShenFunc(symnot), gen990)

								if True == gen991 {
									gen976 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

									Parse__shen_4_5pattern2_6 := gen976
									gen986 := Call(__e, ShenFunc(symfail))

									gen987 := Call(__e, ShenFunc(sym_a), gen986, Parse__shen_4_5pattern2_6)

									gen988 := Call(__e, ShenFunc(symnot), gen987)

									if True == gen988 {
										gen977 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen978 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen979 := Call(__e, ShenFunc(symshen_4pair), gen977, gen978)

										gen980 := Call(__e, ShenFunc(symhd), gen979)

										gen981 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

										gen982 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

										gen983 := Call(__e, ShenFunc(symcons), gen982, Nil)

										gen984 := Call(__e, ShenFunc(symcons), gen981, gen983)

										gen985 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen984)

										gen1008 = Call(__e, ShenFunc(symshen_4pair), gen980, gen985)

									} else {
										gen1008 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen1008 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen1008 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen1008 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen1008
						gen1070 := Call(__e, ShenFunc(symfail))

						gen1071 := Call(__e, ShenFunc(sym_a), YaccParse, gen1070)

						if True == gen1071 {
							gen1045 := Call(__e, ShenFunc(symhd), V220)

							gen1046 := Call(__e, ShenFunc(symcons_2), gen1045)

							var gen1047 Obj
							if True == gen1046 {
								gen1043 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen1044 := Call(__e, ShenFunc(symcons_2), gen1043)

								if True == gen1044 {
									gen1047 = True
								} else {
									gen1047 = False
								}

							} else {
								gen1047 = False
							}
							var gen1048 Obj
							if True == gen1047 {
								gen1037 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen1038 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen1039 := Call(__e, ShenFunc(symshen_4pair), gen1037, gen1038)

								gen1040 := Call(__e, ShenFunc(symhd), gen1039)

								gen1041 := Call(__e, ShenFunc(symcons_2), gen1040)

								var gen1042 Obj
								if True == gen1041 {
									gen1032 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen1033 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen1034 := Call(__e, ShenFunc(symshen_4pair), gen1032, gen1033)

									gen1035 := Call(__e, ShenFunc(symshen_4hdhd), gen1034)

									gen1036 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), gen1035)

									if True == gen1036 {
										gen1042 = True
									} else {
										gen1042 = False
									}

								} else {
									gen1042 = False
								}
								if True == gen1042 {
									gen1009 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen1010 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen1011 := Call(__e, ShenFunc(symshen_4pair), gen1009, gen1010)

									gen1012 := Call(__e, ShenFunc(symshen_4tlhd), gen1011)

									gen1013 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen1014 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen1015 := Call(__e, ShenFunc(symshen_4pair), gen1013, gen1014)

									gen1016 := Call(__e, ShenFunc(symshen_4hdtl), gen1015)

									gen1017 := Call(__e, ShenFunc(symshen_4pair), gen1012, gen1016)

									NewStream217 := gen1017
									gen1029 := Call(__e, ShenFunc(symhd), NewStream217)

									gen1030 := Call(__e, ShenFunc(symcons_2), gen1029)

									var gen1031 Obj
									if True == gen1030 {
										gen1027 := Call(__e, ShenFunc(symshen_4hdhd), NewStream217)

										gen1028 := Call(__e, ShenFunc(sym_a), MakeNumber(0), gen1027)

										if True == gen1028 {
											gen1031 = True
										} else {
											gen1031 = False
										}

									} else {
										gen1031 = False
									}
									if True == gen1031 {
										gen1018 := Call(__e, ShenFunc(symshen_4tlhd), NewStream217)

										gen1019 := Call(__e, ShenFunc(symshen_4hdtl), NewStream217)

										gen1020 := Call(__e, ShenFunc(symshen_4pair), gen1018, gen1019)

										NewStream218 := gen1020
										_ = NewStream218
										gen1021 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen1022 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen1023 := Call(__e, ShenFunc(symshen_4pair), gen1021, gen1022)

										gen1024 := Call(__e, ShenFunc(symhd), gen1023)

										gen1025 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

										gen1026 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen1025)

										gen1048 = Call(__e, ShenFunc(symshen_4pair), gen1024, gen1026)

									} else {
										gen1048 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen1048 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen1048 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen1048
							gen1068 := Call(__e, ShenFunc(symfail))

							gen1069 := Call(__e, ShenFunc(sym_a), YaccParse, gen1068)

							if True == gen1069 {
								gen1057 := Call(__e, ShenFunc(symhd), V220)

								gen1058 := Call(__e, ShenFunc(symcons_2), gen1057)

								var gen1059 Obj
								if True == gen1058 {
									gen1049 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									Parse__X := gen1049
									gen1056 := Call(__e, ShenFunc(symcons_2), Parse__X)

									if True == gen1056 {
										gen1050 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen1051 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen1052 := Call(__e, ShenFunc(symshen_4pair), gen1050, gen1051)

										gen1053 := Call(__e, ShenFunc(symhd), gen1052)

										gen1054 := MakeNative(func(__e Evaluator) {
											__e.TailApply(ShenFunc(symshen_4constructor_1error), Parse__X)

											return
										}, 0)
										gen1055 := Call(__e, ShenFunc(symshen_4custom_1pattern_1compiler), Parse__X, gen1054)

										gen1059 = Call(__e, ShenFunc(symshen_4pair), gen1053, gen1055)

									} else {
										gen1059 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen1059 = Call(__e, ShenFunc(symfail))

								}
								YaccParse := gen1059
								gen1066 := Call(__e, ShenFunc(symfail))

								gen1067 := Call(__e, ShenFunc(sym_a), YaccParse, gen1066)

								if True == gen1067 {
									gen1060 := Call(__e, ShenFunc(symshen_4_5simple__pattern_6), V220)

									Parse__shen_4_5simple__pattern_6 := gen1060
									gen1063 := Call(__e, ShenFunc(symfail))

									gen1064 := Call(__e, ShenFunc(sym_a), gen1063, Parse__shen_4_5simple__pattern_6)

									gen1065 := Call(__e, ShenFunc(symnot), gen1064)

									if True == gen1065 {
										gen1061 := Call(__e, ShenFunc(symhd), Parse__shen_4_5simple__pattern_6)

										gen1062 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5simple__pattern_6)

										__e.TailApply(ShenFunc(symshen_4pair), gen1061, gen1062)

										return

									} else {
										__e.TailApply(ShenFunc(symfail))

										return
									}

								} else {
									__e.Return(YaccParse)
									return
								}

							} else {
								__e.Return(YaccParse)
								return
							}

						} else {
							__e.Return(YaccParse)
							return
						}

					} else {
						__e.Return(YaccParse)
						return
					}

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern>"), gen1078)

		gen1080 := MakeNative(func(__e Evaluator) {
			V222 := __e.Get(1)
			_ = V222
			gen1079 := Call(__e, ShenFunc(symshen_4app), V222, MakeString(" is not a legitimate constructor\n"), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symsimple_1error), gen1079)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.constructor-error"), gen1080)

		gen1104 := MakeNative(func(__e Evaluator) {
			V224 := __e.Get(1)
			_ = V224
			gen1088 := Call(__e, ShenFunc(symhd), V224)

			gen1089 := Call(__e, ShenFunc(symcons_2), gen1088)

			var gen1090 Obj
			if True == gen1089 {
				gen1081 := Call(__e, ShenFunc(symshen_4hdhd), V224)

				Parse__X := gen1081
				gen1087 := Call(__e, ShenFunc(sym_a), Parse__X, MakeSymbol("_"))

				if True == gen1087 {
					gen1082 := Call(__e, ShenFunc(symshen_4tlhd), V224)

					gen1083 := Call(__e, ShenFunc(symshen_4hdtl), V224)

					gen1084 := Call(__e, ShenFunc(symshen_4pair), gen1082, gen1083)

					gen1085 := Call(__e, ShenFunc(symhd), gen1084)

					gen1086 := Call(__e, ShenFunc(symgensym), MakeSymbol("Parse_Y"))

					gen1090 = Call(__e, ShenFunc(symshen_4pair), gen1085, gen1086)

				} else {
					gen1090 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen1090 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen1090
			gen1102 := Call(__e, ShenFunc(symfail))

			gen1103 := Call(__e, ShenFunc(sym_a), YaccParse, gen1102)

			if True == gen1103 {
				gen1100 := Call(__e, ShenFunc(symhd), V224)

				gen1101 := Call(__e, ShenFunc(symcons_2), gen1100)

				if True == gen1101 {
					gen1091 := Call(__e, ShenFunc(symshen_4hdhd), V224)

					Parse__X := gen1091
					gen1096 := Call(__e, ShenFunc(symcons), MakeSymbol("<-"), Nil)

					gen1097 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen1096)

					gen1098 := Call(__e, ShenFunc(symelement_2), Parse__X, gen1097)

					gen1099 := Call(__e, ShenFunc(symnot), gen1098)

					if True == gen1099 {
						gen1092 := Call(__e, ShenFunc(symshen_4tlhd), V224)

						gen1093 := Call(__e, ShenFunc(symshen_4hdtl), V224)

						gen1094 := Call(__e, ShenFunc(symshen_4pair), gen1092, gen1093)

						gen1095 := Call(__e, ShenFunc(symhd), gen1094)

						__e.TailApply(ShenFunc(symshen_4pair), gen1095, Parse__X)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<simple_pattern>"), gen1104)

		gen1111 := MakeNative(func(__e Evaluator) {
			V226 := __e.Get(1)
			_ = V226
			gen1105 := Call(__e, ShenFunc(symshen_4_5pattern_6), V226)

			Parse__shen_4_5pattern_6 := gen1105
			gen1108 := Call(__e, ShenFunc(symfail))

			gen1109 := Call(__e, ShenFunc(sym_a), gen1108, Parse__shen_4_5pattern_6)

			gen1110 := Call(__e, ShenFunc(symnot), gen1109)

			if True == gen1110 {
				gen1106 := Call(__e, ShenFunc(symhd), Parse__shen_4_5pattern_6)

				gen1107 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen1106, gen1107)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern1>"), gen1111)

		gen1118 := MakeNative(func(__e Evaluator) {
			V228 := __e.Get(1)
			_ = V228
			gen1112 := Call(__e, ShenFunc(symshen_4_5pattern_6), V228)

			Parse__shen_4_5pattern_6 := gen1112
			gen1115 := Call(__e, ShenFunc(symfail))

			gen1116 := Call(__e, ShenFunc(sym_a), gen1115, Parse__shen_4_5pattern_6)

			gen1117 := Call(__e, ShenFunc(symnot), gen1116)

			if True == gen1117 {
				gen1113 := Call(__e, ShenFunc(symhd), Parse__shen_4_5pattern_6)

				gen1114 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen1113, gen1114)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern2>"), gen1118)

		gen1126 := MakeNative(func(__e Evaluator) {
			V230 := __e.Get(1)
			_ = V230
			gen1124 := Call(__e, ShenFunc(symhd), V230)

			gen1125 := Call(__e, ShenFunc(symcons_2), gen1124)

			if True == gen1125 {
				gen1119 := Call(__e, ShenFunc(symshen_4hdhd), V230)

				Parse__X := gen1119
				gen1120 := Call(__e, ShenFunc(symshen_4tlhd), V230)

				gen1121 := Call(__e, ShenFunc(symshen_4hdtl), V230)

				gen1122 := Call(__e, ShenFunc(symshen_4pair), gen1120, gen1121)

				gen1123 := Call(__e, ShenFunc(symhd), gen1122)

				__e.TailApply(ShenFunc(symshen_4pair), gen1123, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<action>"), gen1126)

		gen1134 := MakeNative(func(__e Evaluator) {
			V232 := __e.Get(1)
			_ = V232
			gen1132 := Call(__e, ShenFunc(symhd), V232)

			gen1133 := Call(__e, ShenFunc(symcons_2), gen1132)

			if True == gen1133 {
				gen1127 := Call(__e, ShenFunc(symshen_4hdhd), V232)

				Parse__X := gen1127
				gen1128 := Call(__e, ShenFunc(symshen_4tlhd), V232)

				gen1129 := Call(__e, ShenFunc(symshen_4hdtl), V232)

				gen1130 := Call(__e, ShenFunc(symshen_4pair), gen1128, gen1129)

				gen1131 := Call(__e, ShenFunc(symhd), gen1130)

				__e.TailApply(ShenFunc(symshen_4pair), gen1131, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<guard>"), gen1134)

		gen1138 := MakeNative(func(__e Evaluator) {
			V235 := __e.Get(1)
			_ = V235
			V236 := __e.Get(2)
			_ = V236
			gen1135 := Call(__e, ShenFunc(symshen_4compile__to__lambda_7), V235, V236)

			Lambda_7 := gen1135
			gen1136 := Call(__e, ShenFunc(symshen_4compile__to__kl), V235, Lambda_7)

			KL := gen1136
			gen1137 := Call(__e, ShenFunc(symshen_4record_1source), V235, KL)

			Record := gen1137
			_ = Record
			__e.Return(KL)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_machine_code"), gen1138)

		gen1141 := MakeNative(func(__e Evaluator) {
			V241 := __e.Get(1)
			_ = V241
			V242 := __e.Get(2)
			_ = V242
			gen1140 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

			if True == gen1140 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1139 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V241, MakeSymbol("shen.source"), V242, gen1139)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-source"), gen1141)

		gen1154 := MakeNative(func(__e Evaluator) {
			V245 := __e.Get(1)
			_ = V245
			V246 := __e.Get(2)
			_ = V246
			gen1142 := Call(__e, ShenFunc(symshen_4aritycheck), V245, V246)

			Arity := gen1142
			gen1143 := Call(__e, ShenFunc(symshen_4update_1symbol_1table), V245, Arity)

			UpDateSymbolTable := gen1143
			_ = UpDateSymbolTable
			gen1144 := MakeNative(func(__e Evaluator) {
				Rule := __e.Get(1)
				_ = Rule
				__e.TailApply(ShenFunc(symshen_4free__variable__check), V245, Rule)

				return
			}, 1)
			gen1145 := Call(__e, ShenFunc(symshen_4for_1each), gen1144, V246)

			Free := gen1145
			_ = Free
			gen1146 := Call(__e, ShenFunc(symshen_4parameters), Arity)

			Variables := gen1146
			gen1147 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4strip_1protect), X)

				return
			}, 1)
			gen1148 := Call(__e, ShenFunc(symmap), gen1147, V246)

			Strip := gen1148
			gen1149 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4abstract__rule), X)

				return
			}, 1)
			gen1150 := Call(__e, ShenFunc(symmap), gen1149, Strip)

			Abstractions := gen1150
			gen1151 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4application__build), Variables, X)

				return
			}, 1)
			gen1152 := Call(__e, ShenFunc(symmap), gen1151, Abstractions)

			Applications := gen1152
			gen1153 := Call(__e, ShenFunc(symcons), Applications, Nil)

			__e.TailApply(ShenFunc(symcons), Variables, gen1153)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_lambda+"), gen1154)

		gen1159 := MakeNative(func(__e Evaluator) {
			V249 := __e.Get(1)
			_ = V249
			V250 := __e.Get(2)
			_ = V250
			gen1158 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V250)

			if True == gen1158 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1155 := Call(__e, ShenFunc(symshen_4lambda_1form), V249, V250)

				gen1156 := Call(__e, ShenFunc(symeval_1kl), gen1155)

				gen1157 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V249, MakeSymbol("shen.lambda-form"), gen1156, gen1157)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.update-symbol-table"), gen1159)

		gen1173 := MakeNative(func(__e Evaluator) {
			V253 := __e.Get(1)
			_ = V253
			V254 := __e.Get(2)
			_ = V254
			gen1171 := Call(__e, ShenFunc(symcons_2), V254)

			var gen1172 Obj
			if True == gen1171 {
				gen1168 := Call(__e, ShenFunc(symtl), V254)

				gen1169 := Call(__e, ShenFunc(symcons_2), gen1168)

				var gen1170 Obj
				if True == gen1169 {
					gen1165 := Call(__e, ShenFunc(symtl), V254)

					gen1166 := Call(__e, ShenFunc(symtl), gen1165)

					gen1167 := Call(__e, ShenFunc(sym_a), Nil, gen1166)

					if True == gen1167 {
						gen1170 = True
					} else {
						gen1170 = False
					}

				} else {
					gen1170 = False
				}
				if True == gen1170 {
					gen1172 = True
				} else {
					gen1172 = False
				}

			} else {
				gen1172 = False
			}
			if True == gen1172 {
				gen1160 := Call(__e, ShenFunc(symhd), V254)

				gen1161 := Call(__e, ShenFunc(symshen_4extract__vars), gen1160)

				Bound := gen1161
				gen1162 := Call(__e, ShenFunc(symtl), V254)

				gen1163 := Call(__e, ShenFunc(symhd), gen1162)

				gen1164 := Call(__e, ShenFunc(symshen_4extract__free__vars), Bound, gen1163)

				Free := gen1164
				__e.TailApply(ShenFunc(symshen_4free__variable__warnings), V253, Free)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.free_variable_check"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.free_variable_check"), gen1173)

		gen1180 := MakeNative(func(__e Evaluator) {
			V256 := __e.Get(1)
			_ = V256
			gen1179 := Call(__e, ShenFunc(symvariable_2), V256)

			if True == gen1179 {
				__e.TailApply(ShenFunc(symcons), V256, Nil)

				return
			} else {
				gen1178 := Call(__e, ShenFunc(symcons_2), V256)

				if True == gen1178 {
					gen1174 := Call(__e, ShenFunc(symhd), V256)

					gen1175 := Call(__e, ShenFunc(symshen_4extract__vars), gen1174)

					gen1176 := Call(__e, ShenFunc(symtl), V256)

					gen1177 := Call(__e, ShenFunc(symshen_4extract__vars), gen1176)

					__e.TailApply(ShenFunc(symunion), gen1175, gen1177)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extract_vars"), gen1180)

		gen1257 := MakeNative(func(__e Evaluator) {
			V268 := __e.Get(1)
			_ = V268
			V269 := __e.Get(2)
			_ = V269
			gen1255 := Call(__e, ShenFunc(symcons_2), V269)

			var gen1256 Obj
			if True == gen1255 {
				gen1252 := Call(__e, ShenFunc(symtl), V269)

				gen1253 := Call(__e, ShenFunc(symcons_2), gen1252)

				var gen1254 Obj
				if True == gen1253 {
					gen1248 := Call(__e, ShenFunc(symtl), V269)

					gen1249 := Call(__e, ShenFunc(symtl), gen1248)

					gen1250 := Call(__e, ShenFunc(sym_a), Nil, gen1249)

					var gen1251 Obj
					if True == gen1250 {
						gen1246 := Call(__e, ShenFunc(symhd), V269)

						gen1247 := Call(__e, ShenFunc(sym_a), gen1246, MakeSymbol("protect"))

						if True == gen1247 {
							gen1251 = True
						} else {
							gen1251 = False
						}

					} else {
						gen1251 = False
					}
					if True == gen1251 {
						gen1254 = True
					} else {
						gen1254 = False
					}

				} else {
					gen1254 = False
				}
				if True == gen1254 {
					gen1256 = True
				} else {
					gen1256 = False
				}

			} else {
				gen1256 = False
			}
			if True == gen1256 {
				__e.Return(Nil)
				return
			} else {
				gen1244 := Call(__e, ShenFunc(symvariable_2), V269)

				var gen1245 Obj
				if True == gen1244 {
					gen1242 := Call(__e, ShenFunc(symelement_2), V269, V268)

					gen1243 := Call(__e, ShenFunc(symnot), gen1242)

					if True == gen1243 {
						gen1245 = True
					} else {
						gen1245 = False
					}

				} else {
					gen1245 = False
				}
				if True == gen1245 {
					__e.TailApply(ShenFunc(symcons), V269, Nil)

					return
				} else {
					gen1240 := Call(__e, ShenFunc(symcons_2), V269)

					var gen1241 Obj
					if True == gen1240 {
						gen1237 := Call(__e, ShenFunc(symhd), V269)

						gen1238 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen1237)

						var gen1239 Obj
						if True == gen1238 {
							gen1234 := Call(__e, ShenFunc(symtl), V269)

							gen1235 := Call(__e, ShenFunc(symcons_2), gen1234)

							var gen1236 Obj
							if True == gen1235 {
								gen1230 := Call(__e, ShenFunc(symtl), V269)

								gen1231 := Call(__e, ShenFunc(symtl), gen1230)

								gen1232 := Call(__e, ShenFunc(symcons_2), gen1231)

								var gen1233 Obj
								if True == gen1232 {
									gen1226 := Call(__e, ShenFunc(symtl), V269)

									gen1227 := Call(__e, ShenFunc(symtl), gen1226)

									gen1228 := Call(__e, ShenFunc(symtl), gen1227)

									gen1229 := Call(__e, ShenFunc(sym_a), Nil, gen1228)

									if True == gen1229 {
										gen1233 = True
									} else {
										gen1233 = False
									}

								} else {
									gen1233 = False
								}
								if True == gen1233 {
									gen1236 = True
								} else {
									gen1236 = False
								}

							} else {
								gen1236 = False
							}
							if True == gen1236 {
								gen1239 = True
							} else {
								gen1239 = False
							}

						} else {
							gen1239 = False
						}
						if True == gen1239 {
							gen1241 = True
						} else {
							gen1241 = False
						}

					} else {
						gen1241 = False
					}
					if True == gen1241 {
						gen1220 := Call(__e, ShenFunc(symtl), V269)

						gen1221 := Call(__e, ShenFunc(symhd), gen1220)

						gen1222 := Call(__e, ShenFunc(symcons), gen1221, V268)

						gen1223 := Call(__e, ShenFunc(symtl), V269)

						gen1224 := Call(__e, ShenFunc(symtl), gen1223)

						gen1225 := Call(__e, ShenFunc(symhd), gen1224)

						__e.TailApply(ShenFunc(symshen_4extract__free__vars), gen1222, gen1225)

						return

					} else {
						gen1218 := Call(__e, ShenFunc(symcons_2), V269)

						var gen1219 Obj
						if True == gen1218 {
							gen1215 := Call(__e, ShenFunc(symhd), V269)

							gen1216 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen1215)

							var gen1217 Obj
							if True == gen1216 {
								gen1212 := Call(__e, ShenFunc(symtl), V269)

								gen1213 := Call(__e, ShenFunc(symcons_2), gen1212)

								var gen1214 Obj
								if True == gen1213 {
									gen1208 := Call(__e, ShenFunc(symtl), V269)

									gen1209 := Call(__e, ShenFunc(symtl), gen1208)

									gen1210 := Call(__e, ShenFunc(symcons_2), gen1209)

									var gen1211 Obj
									if True == gen1210 {
										gen1203 := Call(__e, ShenFunc(symtl), V269)

										gen1204 := Call(__e, ShenFunc(symtl), gen1203)

										gen1205 := Call(__e, ShenFunc(symtl), gen1204)

										gen1206 := Call(__e, ShenFunc(symcons_2), gen1205)

										var gen1207 Obj
										if True == gen1206 {
											gen1198 := Call(__e, ShenFunc(symtl), V269)

											gen1199 := Call(__e, ShenFunc(symtl), gen1198)

											gen1200 := Call(__e, ShenFunc(symtl), gen1199)

											gen1201 := Call(__e, ShenFunc(symtl), gen1200)

											gen1202 := Call(__e, ShenFunc(sym_a), Nil, gen1201)

											if True == gen1202 {
												gen1207 = True
											} else {
												gen1207 = False
											}

										} else {
											gen1207 = False
										}
										if True == gen1207 {
											gen1211 = True
										} else {
											gen1211 = False
										}

									} else {
										gen1211 = False
									}
									if True == gen1211 {
										gen1214 = True
									} else {
										gen1214 = False
									}

								} else {
									gen1214 = False
								}
								if True == gen1214 {
									gen1217 = True
								} else {
									gen1217 = False
								}

							} else {
								gen1217 = False
							}
							if True == gen1217 {
								gen1219 = True
							} else {
								gen1219 = False
							}

						} else {
							gen1219 = False
						}
						if True == gen1219 {
							gen1186 := Call(__e, ShenFunc(symtl), V269)

							gen1187 := Call(__e, ShenFunc(symtl), gen1186)

							gen1188 := Call(__e, ShenFunc(symhd), gen1187)

							gen1189 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen1188)

							gen1190 := Call(__e, ShenFunc(symtl), V269)

							gen1191 := Call(__e, ShenFunc(symhd), gen1190)

							gen1192 := Call(__e, ShenFunc(symcons), gen1191, V268)

							gen1193 := Call(__e, ShenFunc(symtl), V269)

							gen1194 := Call(__e, ShenFunc(symtl), gen1193)

							gen1195 := Call(__e, ShenFunc(symtl), gen1194)

							gen1196 := Call(__e, ShenFunc(symhd), gen1195)

							gen1197 := Call(__e, ShenFunc(symshen_4extract__free__vars), gen1192, gen1196)

							__e.TailApply(ShenFunc(symunion), gen1189, gen1197)

							return

						} else {
							gen1185 := Call(__e, ShenFunc(symcons_2), V269)

							if True == gen1185 {
								gen1181 := Call(__e, ShenFunc(symhd), V269)

								gen1182 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen1181)

								gen1183 := Call(__e, ShenFunc(symtl), V269)

								gen1184 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen1183)

								__e.TailApply(ShenFunc(symunion), gen1182, gen1184)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extract_free_vars"), gen1257)

		gen1264 := MakeNative(func(__e Evaluator) {
			V274 := __e.Get(1)
			_ = V274
			V275 := __e.Get(2)
			_ = V275
			gen1263 := Call(__e, ShenFunc(sym_a), Nil, V275)

			if True == gen1263 {
				__e.Return(MakeSymbol("_"))
				return
			} else {
				gen1258 := Call(__e, ShenFunc(symshen_4list__variables), V275)

				gen1259 := Call(__e, ShenFunc(symshen_4app), gen1258, MakeString(""), MakeSymbol("shen.a"))

				gen1260 := Call(__e, ShenFunc(symcn), MakeString(": "), gen1259)

				gen1261 := Call(__e, ShenFunc(symshen_4app), V274, gen1260, MakeSymbol("shen.a"))

				gen1262 := Call(__e, ShenFunc(symcn), MakeString("error: the following variables are free in "), gen1261)

				__e.TailApply(ShenFunc(symsimple_1error), gen1262)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.free_variable_warnings"), gen1264)

		gen1277 := MakeNative(func(__e Evaluator) {
			V277 := __e.Get(1)
			_ = V277
			gen1275 := Call(__e, ShenFunc(symcons_2), V277)

			var gen1276 Obj
			if True == gen1275 {
				gen1273 := Call(__e, ShenFunc(symtl), V277)

				gen1274 := Call(__e, ShenFunc(sym_a), Nil, gen1273)

				if True == gen1274 {
					gen1276 = True
				} else {
					gen1276 = False
				}

			} else {
				gen1276 = False
			}
			if True == gen1276 {
				gen1271 := Call(__e, ShenFunc(symhd), V277)

				gen1272 := Call(__e, ShenFunc(symstr), gen1271)

				__e.TailApply(ShenFunc(symcn), gen1272, MakeString("."))

				return

			} else {
				gen1270 := Call(__e, ShenFunc(symcons_2), V277)

				if True == gen1270 {
					gen1265 := Call(__e, ShenFunc(symhd), V277)

					gen1266 := Call(__e, ShenFunc(symstr), gen1265)

					gen1267 := Call(__e, ShenFunc(symtl), V277)

					gen1268 := Call(__e, ShenFunc(symshen_4list__variables), gen1267)

					gen1269 := Call(__e, ShenFunc(symcn), MakeString(", "), gen1268)

					__e.TailApply(ShenFunc(symcn), gen1266, gen1269)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.list_variables"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.list_variables"), gen1277)

		gen1293 := MakeNative(func(__e Evaluator) {
			V279 := __e.Get(1)
			_ = V279
			gen1291 := Call(__e, ShenFunc(symcons_2), V279)

			var gen1292 Obj
			if True == gen1291 {
				gen1288 := Call(__e, ShenFunc(symtl), V279)

				gen1289 := Call(__e, ShenFunc(symcons_2), gen1288)

				var gen1290 Obj
				if True == gen1289 {
					gen1284 := Call(__e, ShenFunc(symtl), V279)

					gen1285 := Call(__e, ShenFunc(symtl), gen1284)

					gen1286 := Call(__e, ShenFunc(sym_a), Nil, gen1285)

					var gen1287 Obj
					if True == gen1286 {
						gen1282 := Call(__e, ShenFunc(symhd), V279)

						gen1283 := Call(__e, ShenFunc(sym_a), gen1282, MakeSymbol("protect"))

						if True == gen1283 {
							gen1287 = True
						} else {
							gen1287 = False
						}

					} else {
						gen1287 = False
					}
					if True == gen1287 {
						gen1290 = True
					} else {
						gen1290 = False
					}

				} else {
					gen1290 = False
				}
				if True == gen1290 {
					gen1292 = True
				} else {
					gen1292 = False
				}

			} else {
				gen1292 = False
			}
			if True == gen1292 {
				gen1280 := Call(__e, ShenFunc(symtl), V279)

				gen1281 := Call(__e, ShenFunc(symhd), gen1280)

				__e.TailApply(ShenFunc(symshen_4strip_1protect), gen1281)

				return

			} else {
				gen1279 := Call(__e, ShenFunc(symcons_2), V279)

				if True == gen1279 {
					gen1278 := MakeNative(func(__e Evaluator) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(ShenFunc(symshen_4strip_1protect), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen1278, V279)

					return

				} else {
					__e.Return(V279)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.strip-protect"), gen1293)

		gen1307 := MakeNative(func(__e Evaluator) {
			V281 := __e.Get(1)
			_ = V281
			gen1305 := Call(__e, ShenFunc(symcons_2), V281)

			var gen1306 Obj
			if True == gen1305 {
				gen1302 := Call(__e, ShenFunc(symtl), V281)

				gen1303 := Call(__e, ShenFunc(symcons_2), gen1302)

				var gen1304 Obj
				if True == gen1303 {
					gen1299 := Call(__e, ShenFunc(symtl), V281)

					gen1300 := Call(__e, ShenFunc(symtl), gen1299)

					gen1301 := Call(__e, ShenFunc(sym_a), Nil, gen1300)

					if True == gen1301 {
						gen1304 = True
					} else {
						gen1304 = False
					}

				} else {
					gen1304 = False
				}
				if True == gen1304 {
					gen1306 = True
				} else {
					gen1306 = False
				}

			} else {
				gen1306 = False
			}
			if True == gen1306 {
				gen1294 := Call(__e, ShenFunc(symhd), V281)

				gen1295 := Call(__e, ShenFunc(symshen_4flatten), gen1294)

				gen1296 := Call(__e, ShenFunc(symhd), V281)

				gen1297 := Call(__e, ShenFunc(symtl), V281)

				gen1298 := Call(__e, ShenFunc(symhd), gen1297)

				__e.TailApply(ShenFunc(symshen_4linearise__help), gen1295, gen1296, gen1298)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise"), gen1307)

		gen1314 := MakeNative(func(__e Evaluator) {
			V283 := __e.Get(1)
			_ = V283
			gen1313 := Call(__e, ShenFunc(sym_a), Nil, V283)

			if True == gen1313 {
				__e.Return(Nil)
				return
			} else {
				gen1312 := Call(__e, ShenFunc(symcons_2), V283)

				if True == gen1312 {
					gen1308 := Call(__e, ShenFunc(symhd), V283)

					gen1309 := Call(__e, ShenFunc(symshen_4flatten), gen1308)

					gen1310 := Call(__e, ShenFunc(symtl), V283)

					gen1311 := Call(__e, ShenFunc(symshen_4flatten), gen1310)

					__e.TailApply(ShenFunc(symappend), gen1309, gen1311)

					return

				} else {
					__e.TailApply(ShenFunc(symcons), V283, Nil)

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.flatten"), gen1314)

		gen1337 := MakeNative(func(__e Evaluator) {
			V287 := __e.Get(1)
			_ = V287
			V288 := __e.Get(2)
			_ = V288
			V289 := __e.Get(3)
			_ = V289
			gen1336 := Call(__e, ShenFunc(sym_a), Nil, V287)

			if True == gen1336 {
				gen1335 := Call(__e, ShenFunc(symcons), V289, Nil)

				__e.TailApply(ShenFunc(symcons), V288, gen1335)

				return

			} else {
				gen1334 := Call(__e, ShenFunc(symcons_2), V287)

				if True == gen1334 {
					gen1331 := Call(__e, ShenFunc(symhd), V287)

					gen1332 := Call(__e, ShenFunc(symvariable_2), gen1331)

					var gen1333 Obj
					if True == gen1332 {
						gen1328 := Call(__e, ShenFunc(symhd), V287)

						gen1329 := Call(__e, ShenFunc(symtl), V287)

						gen1330 := Call(__e, ShenFunc(symelement_2), gen1328, gen1329)

						if True == gen1330 {
							gen1333 = True
						} else {
							gen1333 = False
						}

					} else {
						gen1333 = False
					}
					if True == gen1333 {
						gen1316 := Call(__e, ShenFunc(symhd), V287)

						gen1317 := Call(__e, ShenFunc(symgensym), gen1316)

						Var := gen1317
						gen1318 := Call(__e, ShenFunc(symhd), V287)

						gen1319 := Call(__e, ShenFunc(symcons), Var, Nil)

						gen1320 := Call(__e, ShenFunc(symcons), gen1318, gen1319)

						gen1321 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen1320)

						gen1322 := Call(__e, ShenFunc(symcons), V289, Nil)

						gen1323 := Call(__e, ShenFunc(symcons), gen1321, gen1322)

						gen1324 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen1323)

						NewAction := gen1324
						gen1325 := Call(__e, ShenFunc(symhd), V287)

						gen1326 := Call(__e, ShenFunc(symshen_4linearise__X), gen1325, Var, V288)

						NewPatts := gen1326
						gen1327 := Call(__e, ShenFunc(symtl), V287)

						__e.TailApply(ShenFunc(symshen_4linearise__help), gen1327, NewPatts, NewAction)

						return

					} else {
						gen1315 := Call(__e, ShenFunc(symtl), V287)

						__e.TailApply(ShenFunc(symshen_4linearise__help), gen1315, V288, V289)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise_help"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise_help"), gen1337)

		gen1348 := MakeNative(func(__e Evaluator) {
			V302 := __e.Get(1)
			_ = V302
			V303 := __e.Get(2)
			_ = V303
			V304 := __e.Get(3)
			_ = V304
			gen1347 := Call(__e, ShenFunc(sym_a), V304, V302)

			if True == gen1347 {
				__e.Return(V303)
				return
			} else {
				gen1346 := Call(__e, ShenFunc(symcons_2), V304)

				if True == gen1346 {
					gen1338 := Call(__e, ShenFunc(symhd), V304)

					gen1339 := Call(__e, ShenFunc(symshen_4linearise__X), V302, V303, gen1338)

					L := gen1339
					gen1344 := Call(__e, ShenFunc(symhd), V304)

					gen1345 := Call(__e, ShenFunc(sym_a), L, gen1344)

					if True == gen1345 {
						gen1341 := Call(__e, ShenFunc(symhd), V304)

						gen1342 := Call(__e, ShenFunc(symtl), V304)

						gen1343 := Call(__e, ShenFunc(symshen_4linearise__X), V302, V303, gen1342)

						__e.TailApply(ShenFunc(symcons), gen1341, gen1343)

						return

					} else {
						gen1340 := Call(__e, ShenFunc(symtl), V304)

						__e.TailApply(ShenFunc(symcons), L, gen1340)

						return

					}

				} else {
					__e.Return(V304)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise_X"), gen1348)

		gen1417 := MakeNative(func(__e Evaluator) {
			V307 := __e.Get(1)
			_ = V307
			V308 := __e.Get(2)
			_ = V308
			gen1415 := Call(__e, ShenFunc(symcons_2), V308)

			var gen1416 Obj
			if True == gen1415 {
				gen1412 := Call(__e, ShenFunc(symhd), V308)

				gen1413 := Call(__e, ShenFunc(symcons_2), gen1412)

				var gen1414 Obj
				if True == gen1413 {
					gen1408 := Call(__e, ShenFunc(symhd), V308)

					gen1409 := Call(__e, ShenFunc(symtl), gen1408)

					gen1410 := Call(__e, ShenFunc(symcons_2), gen1409)

					var gen1411 Obj
					if True == gen1410 {
						gen1403 := Call(__e, ShenFunc(symhd), V308)

						gen1404 := Call(__e, ShenFunc(symtl), gen1403)

						gen1405 := Call(__e, ShenFunc(symtl), gen1404)

						gen1406 := Call(__e, ShenFunc(sym_a), Nil, gen1405)

						var gen1407 Obj
						if True == gen1406 {
							gen1401 := Call(__e, ShenFunc(symtl), V308)

							gen1402 := Call(__e, ShenFunc(sym_a), Nil, gen1401)

							if True == gen1402 {
								gen1407 = True
							} else {
								gen1407 = False
							}

						} else {
							gen1407 = False
						}
						if True == gen1407 {
							gen1411 = True
						} else {
							gen1411 = False
						}

					} else {
						gen1411 = False
					}
					if True == gen1411 {
						gen1414 = True
					} else {
						gen1414 = False
					}

				} else {
					gen1414 = False
				}
				if True == gen1414 {
					gen1416 = True
				} else {
					gen1416 = False
				}

			} else {
				gen1416 = False
			}
			if True == gen1416 {
				gen1394 := Call(__e, ShenFunc(symhd), V308)

				gen1395 := Call(__e, ShenFunc(symtl), gen1394)

				gen1396 := Call(__e, ShenFunc(symhd), gen1395)

				Call(__e, ShenFunc(symshen_4aritycheck_1action), gen1396)

				gen1397 := Call(__e, ShenFunc(symarity), V307)

				gen1398 := Call(__e, ShenFunc(symhd), V308)

				gen1399 := Call(__e, ShenFunc(symhd), gen1398)

				gen1400 := Call(__e, ShenFunc(symlength), gen1399)

				__e.TailApply(ShenFunc(symshen_4aritycheck_1name), V307, gen1397, gen1400)

				return

			} else {
				gen1392 := Call(__e, ShenFunc(symcons_2), V308)

				var gen1393 Obj
				if True == gen1392 {
					gen1389 := Call(__e, ShenFunc(symhd), V308)

					gen1390 := Call(__e, ShenFunc(symcons_2), gen1389)

					var gen1391 Obj
					if True == gen1390 {
						gen1385 := Call(__e, ShenFunc(symhd), V308)

						gen1386 := Call(__e, ShenFunc(symtl), gen1385)

						gen1387 := Call(__e, ShenFunc(symcons_2), gen1386)

						var gen1388 Obj
						if True == gen1387 {
							gen1380 := Call(__e, ShenFunc(symhd), V308)

							gen1381 := Call(__e, ShenFunc(symtl), gen1380)

							gen1382 := Call(__e, ShenFunc(symtl), gen1381)

							gen1383 := Call(__e, ShenFunc(sym_a), Nil, gen1382)

							var gen1384 Obj
							if True == gen1383 {
								gen1377 := Call(__e, ShenFunc(symtl), V308)

								gen1378 := Call(__e, ShenFunc(symcons_2), gen1377)

								var gen1379 Obj
								if True == gen1378 {
									gen1373 := Call(__e, ShenFunc(symtl), V308)

									gen1374 := Call(__e, ShenFunc(symhd), gen1373)

									gen1375 := Call(__e, ShenFunc(symcons_2), gen1374)

									var gen1376 Obj
									if True == gen1375 {
										gen1368 := Call(__e, ShenFunc(symtl), V308)

										gen1369 := Call(__e, ShenFunc(symhd), gen1368)

										gen1370 := Call(__e, ShenFunc(symtl), gen1369)

										gen1371 := Call(__e, ShenFunc(symcons_2), gen1370)

										var gen1372 Obj
										if True == gen1371 {
											gen1363 := Call(__e, ShenFunc(symtl), V308)

											gen1364 := Call(__e, ShenFunc(symhd), gen1363)

											gen1365 := Call(__e, ShenFunc(symtl), gen1364)

											gen1366 := Call(__e, ShenFunc(symtl), gen1365)

											gen1367 := Call(__e, ShenFunc(sym_a), Nil, gen1366)

											if True == gen1367 {
												gen1372 = True
											} else {
												gen1372 = False
											}

										} else {
											gen1372 = False
										}
										if True == gen1372 {
											gen1376 = True
										} else {
											gen1376 = False
										}

									} else {
										gen1376 = False
									}
									if True == gen1376 {
										gen1379 = True
									} else {
										gen1379 = False
									}

								} else {
									gen1379 = False
								}
								if True == gen1379 {
									gen1384 = True
								} else {
									gen1384 = False
								}

							} else {
								gen1384 = False
							}
							if True == gen1384 {
								gen1388 = True
							} else {
								gen1388 = False
							}

						} else {
							gen1388 = False
						}
						if True == gen1388 {
							gen1391 = True
						} else {
							gen1391 = False
						}

					} else {
						gen1391 = False
					}
					if True == gen1391 {
						gen1393 = True
					} else {
						gen1393 = False
					}

				} else {
					gen1393 = False
				}
				if True == gen1393 {
					gen1355 := Call(__e, ShenFunc(symhd), V308)

					gen1356 := Call(__e, ShenFunc(symhd), gen1355)

					gen1357 := Call(__e, ShenFunc(symlength), gen1356)

					gen1358 := Call(__e, ShenFunc(symtl), V308)

					gen1359 := Call(__e, ShenFunc(symhd), gen1358)

					gen1360 := Call(__e, ShenFunc(symhd), gen1359)

					gen1361 := Call(__e, ShenFunc(symlength), gen1360)

					gen1362 := Call(__e, ShenFunc(sym_a), gen1357, gen1361)

					if True == gen1362 {
						gen1351 := Call(__e, ShenFunc(symhd), V308)

						gen1352 := Call(__e, ShenFunc(symtl), gen1351)

						gen1353 := Call(__e, ShenFunc(symhd), gen1352)

						Call(__e, ShenFunc(symshen_4aritycheck_1action), gen1353)

						gen1354 := Call(__e, ShenFunc(symtl), V308)

						__e.TailApply(ShenFunc(symshen_4aritycheck), V307, gen1354)

						return

					} else {
						gen1349 := Call(__e, ShenFunc(symshen_4app), V307, MakeString("\n"), MakeSymbol("shen.a"))

						gen1350 := Call(__e, ShenFunc(symcn), MakeString("arity error in "), gen1349)

						__e.TailApply(ShenFunc(symsimple_1error), gen1350)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.aritycheck"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck"), gen1417)

		gen1423 := MakeNative(func(__e Evaluator) {
			V321 := __e.Get(1)
			_ = V321
			V322 := __e.Get(2)
			_ = V322
			V323 := __e.Get(3)
			_ = V323
			gen1422 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V322)

			if True == gen1422 {
				__e.Return(V323)
				return
			} else {
				gen1421 := Call(__e, ShenFunc(sym_a), V323, V322)

				if True == gen1421 {
					__e.Return(V323)
					return
				} else {
					gen1418 := Call(__e, ShenFunc(symshen_4app), V321, MakeString(" can cause errors.\n"), MakeSymbol("shen.a"))

					gen1419 := Call(__e, ShenFunc(symcn), MakeString("\nwarning: changing the arity of "), gen1418)

					gen1420 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), gen1419, gen1420)

					__e.Return(V323)
					return

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck-name"), gen1423)

		gen1428 := MakeNative(func(__e Evaluator) {
			V329 := __e.Get(1)
			_ = V329
			gen1427 := Call(__e, ShenFunc(symcons_2), V329)

			if True == gen1427 {
				gen1424 := Call(__e, ShenFunc(symhd), V329)

				gen1425 := Call(__e, ShenFunc(symtl), V329)

				Call(__e, ShenFunc(symshen_4aah), gen1424, gen1425)

				gen1426 := MakeNative(func(__e Evaluator) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(ShenFunc(symshen_4aritycheck_1action), Y)

					return
				}, 1)
				__e.TailApply(ShenFunc(symshen_4for_1each), gen1426, V329)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck-action"), gen1428)

		gen1443 := MakeNative(func(__e Evaluator) {
			V332 := __e.Get(1)
			_ = V332
			V333 := __e.Get(2)
			_ = V333
			gen1429 := Call(__e, ShenFunc(symarity), V332)

			Arity := gen1429
			gen1430 := Call(__e, ShenFunc(symlength), V333)

			Len := gen1430
			gen1441 := Call(__e, ShenFunc(sym_6), Arity, MakeNumber(-1))

			var gen1442 Obj
			if True == gen1441 {
				gen1440 := Call(__e, ShenFunc(sym_6), Len, Arity)

				if True == gen1440 {
					gen1442 = True
				} else {
					gen1442 = False
				}

			} else {
				gen1442 = False
			}
			if True == gen1442 {
				gen1431 := Call(__e, ShenFunc(sym_6), Len, MakeNumber(1))

				var gen1432 Obj
				if True == gen1431 {
					gen1432 = MakeString("s")
				} else {
					gen1432 = MakeString("")
				}
				gen1433 := Call(__e, ShenFunc(symshen_4app), gen1432, MakeString(".\n"), MakeSymbol("shen.a"))

				gen1434 := Call(__e, ShenFunc(symcn), MakeString(" argument"), gen1433)

				gen1435 := Call(__e, ShenFunc(symshen_4app), Len, gen1434, MakeSymbol("shen.a"))

				gen1436 := Call(__e, ShenFunc(symcn), MakeString(" might not like "), gen1435)

				gen1437 := Call(__e, ShenFunc(symshen_4app), V332, gen1436, MakeSymbol("shen.a"))

				gen1438 := Call(__e, ShenFunc(symcn), MakeString("warning: "), gen1437)

				gen1439 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen1438, gen1439)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aah"), gen1443)

		gen1455 := MakeNative(func(__e Evaluator) {
			V335 := __e.Get(1)
			_ = V335
			gen1453 := Call(__e, ShenFunc(symcons_2), V335)

			var gen1454 Obj
			if True == gen1453 {
				gen1450 := Call(__e, ShenFunc(symtl), V335)

				gen1451 := Call(__e, ShenFunc(symcons_2), gen1450)

				var gen1452 Obj
				if True == gen1451 {
					gen1447 := Call(__e, ShenFunc(symtl), V335)

					gen1448 := Call(__e, ShenFunc(symtl), gen1447)

					gen1449 := Call(__e, ShenFunc(sym_a), Nil, gen1448)

					if True == gen1449 {
						gen1452 = True
					} else {
						gen1452 = False
					}

				} else {
					gen1452 = False
				}
				if True == gen1452 {
					gen1454 = True
				} else {
					gen1454 = False
				}

			} else {
				gen1454 = False
			}
			if True == gen1454 {
				gen1444 := Call(__e, ShenFunc(symhd), V335)

				gen1445 := Call(__e, ShenFunc(symtl), V335)

				gen1446 := Call(__e, ShenFunc(symhd), gen1445)

				__e.TailApply(ShenFunc(symshen_4abstraction__build), gen1444, gen1446)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.abstract_rule"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abstract_rule"), gen1455)

		gen1463 := MakeNative(func(__e Evaluator) {
			V338 := __e.Get(1)
			_ = V338
			V339 := __e.Get(2)
			_ = V339
			gen1462 := Call(__e, ShenFunc(sym_a), Nil, V338)

			if True == gen1462 {
				__e.Return(V339)
				return
			} else {
				gen1461 := Call(__e, ShenFunc(symcons_2), V338)

				if True == gen1461 {
					gen1456 := Call(__e, ShenFunc(symhd), V338)

					gen1457 := Call(__e, ShenFunc(symtl), V338)

					gen1458 := Call(__e, ShenFunc(symshen_4abstraction__build), gen1457, V339)

					gen1459 := Call(__e, ShenFunc(symcons), gen1458, Nil)

					gen1460 := Call(__e, ShenFunc(symcons), gen1456, gen1459)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("/."), gen1460)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.abstraction_build"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abstraction_build"), gen1463)

		gen1468 := MakeNative(func(__e Evaluator) {
			V341 := __e.Get(1)
			_ = V341
			gen1467 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V341)

			if True == gen1467 {
				__e.Return(Nil)
				return
			} else {
				gen1464 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

				gen1465 := Call(__e, ShenFunc(sym_1), V341, MakeNumber(1))

				gen1466 := Call(__e, ShenFunc(symshen_4parameters), gen1465)

				__e.TailApply(ShenFunc(symcons), gen1464, gen1466)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.parameters"), gen1468)

		gen1475 := MakeNative(func(__e Evaluator) {
			V344 := __e.Get(1)
			_ = V344
			V345 := __e.Get(2)
			_ = V345
			gen1474 := Call(__e, ShenFunc(sym_a), Nil, V344)

			if True == gen1474 {
				__e.Return(V345)
				return
			} else {
				gen1473 := Call(__e, ShenFunc(symcons_2), V344)

				if True == gen1473 {
					gen1469 := Call(__e, ShenFunc(symtl), V344)

					gen1470 := Call(__e, ShenFunc(symhd), V344)

					gen1471 := Call(__e, ShenFunc(symcons), gen1470, Nil)

					gen1472 := Call(__e, ShenFunc(symcons), V345, gen1471)

					__e.TailApply(ShenFunc(symshen_4application__build), gen1469, gen1472)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.application_build"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.application_build"), gen1475)

		gen1504 := MakeNative(func(__e Evaluator) {
			V348 := __e.Get(1)
			_ = V348
			V349 := __e.Get(2)
			_ = V349
			gen1502 := Call(__e, ShenFunc(symcons_2), V349)

			var gen1503 Obj
			if True == gen1502 {
				gen1499 := Call(__e, ShenFunc(symtl), V349)

				gen1500 := Call(__e, ShenFunc(symcons_2), gen1499)

				var gen1501 Obj
				if True == gen1500 {
					gen1496 := Call(__e, ShenFunc(symtl), V349)

					gen1497 := Call(__e, ShenFunc(symtl), gen1496)

					gen1498 := Call(__e, ShenFunc(sym_a), Nil, gen1497)

					if True == gen1498 {
						gen1501 = True
					} else {
						gen1501 = False
					}

				} else {
					gen1501 = False
				}
				if True == gen1501 {
					gen1503 = True
				} else {
					gen1503 = False
				}

			} else {
				gen1503 = False
			}
			if True == gen1503 {
				gen1476 := Call(__e, ShenFunc(symhd), V349)

				gen1477 := Call(__e, ShenFunc(symlength), gen1476)

				gen1478 := Call(__e, ShenFunc(symshen_4store_1arity), V348, gen1477)

				Arity := gen1478
				_ = Arity
				gen1479 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(ShenFunc(symshen_4reduce), X)

					return
				}, 1)
				gen1480 := Call(__e, ShenFunc(symtl), V349)

				gen1481 := Call(__e, ShenFunc(symhd), gen1480)

				gen1482 := Call(__e, ShenFunc(symmap), gen1479, gen1481)

				Reduce := gen1482
				gen1483 := Call(__e, ShenFunc(symhd), V349)

				gen1484 := Call(__e, ShenFunc(symshen_4cond_1expression), V348, gen1483, Reduce)

				CondExpression := gen1484
				gen1487 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*optimise*"))

				var gen1488 Obj
				if True == gen1487 {
					gen1485 := Call(__e, ShenFunc(symshen_4get_1type), V348)

					gen1486 := Call(__e, ShenFunc(symhd), V349)

					gen1488 = Call(__e, ShenFunc(symshen_4typextable), gen1485, gen1486)

				} else {
					gen1488 = MakeSymbol("shen.skip")
				}
				TypeTable := gen1488
				gen1490 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*optimise*"))

				var gen1491 Obj
				if True == gen1490 {
					gen1489 := Call(__e, ShenFunc(symhd), V349)

					gen1491 = Call(__e, ShenFunc(symshen_4assign_1types), gen1489, TypeTable, CondExpression)

				} else {
					gen1491 = CondExpression
				}
				TypedCondExpression := gen1491
				gen1492 := Call(__e, ShenFunc(symhd), V349)

				gen1493 := Call(__e, ShenFunc(symcons), TypedCondExpression, Nil)

				gen1494 := Call(__e, ShenFunc(symcons), gen1492, gen1493)

				gen1495 := Call(__e, ShenFunc(symcons), V348, gen1494)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("defun"), gen1495)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compile_to_kl"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_kl"), gen1504)

		gen1509 := MakeNative(func(__e Evaluator) {
			V355 := __e.Get(1)
			_ = V355
			gen1508 := Call(__e, ShenFunc(symcons_2), V355)

			if True == gen1508 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1505 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

				gen1506 := Call(__e, ShenFunc(symassoc), V355, gen1505)

				FType := gen1506
				gen1507 := Call(__e, ShenFunc(symempty_2), FType)

				if True == gen1507 {
					__e.Return(MakeSymbol("shen.skip"))
					return
				} else {
					__e.TailApply(ShenFunc(symtl), FType)

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.get-type"), gen1509)

		gen1543 := MakeNative(func(__e Evaluator) {
			V366 := __e.Get(1)
			_ = V366
			V367 := __e.Get(2)
			_ = V367
			gen1541 := Call(__e, ShenFunc(symcons_2), V366)

			var gen1542 Obj
			if True == gen1541 {
				gen1538 := Call(__e, ShenFunc(symtl), V366)

				gen1539 := Call(__e, ShenFunc(symcons_2), gen1538)

				var gen1540 Obj
				if True == gen1539 {
					gen1534 := Call(__e, ShenFunc(symtl), V366)

					gen1535 := Call(__e, ShenFunc(symhd), gen1534)

					gen1536 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen1535)

					var gen1537 Obj
					if True == gen1536 {
						gen1530 := Call(__e, ShenFunc(symtl), V366)

						gen1531 := Call(__e, ShenFunc(symtl), gen1530)

						gen1532 := Call(__e, ShenFunc(symcons_2), gen1531)

						var gen1533 Obj
						if True == gen1532 {
							gen1525 := Call(__e, ShenFunc(symtl), V366)

							gen1526 := Call(__e, ShenFunc(symtl), gen1525)

							gen1527 := Call(__e, ShenFunc(symtl), gen1526)

							gen1528 := Call(__e, ShenFunc(sym_a), Nil, gen1527)

							var gen1529 Obj
							if True == gen1528 {
								gen1524 := Call(__e, ShenFunc(symcons_2), V367)

								if True == gen1524 {
									gen1529 = True
								} else {
									gen1529 = False
								}

							} else {
								gen1529 = False
							}
							if True == gen1529 {
								gen1533 = True
							} else {
								gen1533 = False
							}

						} else {
							gen1533 = False
						}
						if True == gen1533 {
							gen1537 = True
						} else {
							gen1537 = False
						}

					} else {
						gen1537 = False
					}
					if True == gen1537 {
						gen1540 = True
					} else {
						gen1540 = False
					}

				} else {
					gen1540 = False
				}
				if True == gen1540 {
					gen1542 = True
				} else {
					gen1542 = False
				}

			} else {
				gen1542 = False
			}
			if True == gen1542 {
				gen1522 := Call(__e, ShenFunc(symhd), V366)

				gen1523 := Call(__e, ShenFunc(symvariable_2), gen1522)

				if True == gen1523 {
					gen1518 := Call(__e, ShenFunc(symtl), V366)

					gen1519 := Call(__e, ShenFunc(symtl), gen1518)

					gen1520 := Call(__e, ShenFunc(symhd), gen1519)

					gen1521 := Call(__e, ShenFunc(symtl), V367)

					__e.TailApply(ShenFunc(symshen_4typextable), gen1520, gen1521)

					return

				} else {
					gen1510 := Call(__e, ShenFunc(symhd), V367)

					gen1511 := Call(__e, ShenFunc(symhd), V366)

					gen1512 := Call(__e, ShenFunc(symcons), gen1510, gen1511)

					gen1513 := Call(__e, ShenFunc(symtl), V366)

					gen1514 := Call(__e, ShenFunc(symtl), gen1513)

					gen1515 := Call(__e, ShenFunc(symhd), gen1514)

					gen1516 := Call(__e, ShenFunc(symtl), V367)

					gen1517 := Call(__e, ShenFunc(symshen_4typextable), gen1515, gen1516)

					__e.TailApply(ShenFunc(symcons), gen1512, gen1517)

					return

				}

			} else {
				__e.Return(Nil)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typextable"), gen1543)

		gen1639 := MakeNative(func(__e Evaluator) {
			V371 := __e.Get(1)
			_ = V371
			V372 := __e.Get(2)
			_ = V372
			V373 := __e.Get(3)
			_ = V373
			gen1637 := Call(__e, ShenFunc(symcons_2), V373)

			var gen1638 Obj
			if True == gen1637 {
				gen1634 := Call(__e, ShenFunc(symhd), V373)

				gen1635 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen1634)

				var gen1636 Obj
				if True == gen1635 {
					gen1631 := Call(__e, ShenFunc(symtl), V373)

					gen1632 := Call(__e, ShenFunc(symcons_2), gen1631)

					var gen1633 Obj
					if True == gen1632 {
						gen1627 := Call(__e, ShenFunc(symtl), V373)

						gen1628 := Call(__e, ShenFunc(symtl), gen1627)

						gen1629 := Call(__e, ShenFunc(symcons_2), gen1628)

						var gen1630 Obj
						if True == gen1629 {
							gen1622 := Call(__e, ShenFunc(symtl), V373)

							gen1623 := Call(__e, ShenFunc(symtl), gen1622)

							gen1624 := Call(__e, ShenFunc(symtl), gen1623)

							gen1625 := Call(__e, ShenFunc(symcons_2), gen1624)

							var gen1626 Obj
							if True == gen1625 {
								gen1617 := Call(__e, ShenFunc(symtl), V373)

								gen1618 := Call(__e, ShenFunc(symtl), gen1617)

								gen1619 := Call(__e, ShenFunc(symtl), gen1618)

								gen1620 := Call(__e, ShenFunc(symtl), gen1619)

								gen1621 := Call(__e, ShenFunc(sym_a), Nil, gen1620)

								if True == gen1621 {
									gen1626 = True
								} else {
									gen1626 = False
								}

							} else {
								gen1626 = False
							}
							if True == gen1626 {
								gen1630 = True
							} else {
								gen1630 = False
							}

						} else {
							gen1630 = False
						}
						if True == gen1630 {
							gen1633 = True
						} else {
							gen1633 = False
						}

					} else {
						gen1633 = False
					}
					if True == gen1633 {
						gen1636 = True
					} else {
						gen1636 = False
					}

				} else {
					gen1636 = False
				}
				if True == gen1636 {
					gen1638 = True
				} else {
					gen1638 = False
				}

			} else {
				gen1638 = False
			}
			if True == gen1638 {
				gen1600 := Call(__e, ShenFunc(symtl), V373)

				gen1601 := Call(__e, ShenFunc(symhd), gen1600)

				gen1602 := Call(__e, ShenFunc(symtl), V373)

				gen1603 := Call(__e, ShenFunc(symtl), gen1602)

				gen1604 := Call(__e, ShenFunc(symhd), gen1603)

				gen1605 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1604)

				gen1606 := Call(__e, ShenFunc(symtl), V373)

				gen1607 := Call(__e, ShenFunc(symhd), gen1606)

				gen1608 := Call(__e, ShenFunc(symcons), gen1607, V371)

				gen1609 := Call(__e, ShenFunc(symtl), V373)

				gen1610 := Call(__e, ShenFunc(symtl), gen1609)

				gen1611 := Call(__e, ShenFunc(symtl), gen1610)

				gen1612 := Call(__e, ShenFunc(symhd), gen1611)

				gen1613 := Call(__e, ShenFunc(symshen_4assign_1types), gen1608, V372, gen1612)

				gen1614 := Call(__e, ShenFunc(symcons), gen1613, Nil)

				gen1615 := Call(__e, ShenFunc(symcons), gen1605, gen1614)

				gen1616 := Call(__e, ShenFunc(symcons), gen1601, gen1615)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen1616)

				return

			} else {
				gen1598 := Call(__e, ShenFunc(symcons_2), V373)

				var gen1599 Obj
				if True == gen1598 {
					gen1595 := Call(__e, ShenFunc(symhd), V373)

					gen1596 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen1595)

					var gen1597 Obj
					if True == gen1596 {
						gen1592 := Call(__e, ShenFunc(symtl), V373)

						gen1593 := Call(__e, ShenFunc(symcons_2), gen1592)

						var gen1594 Obj
						if True == gen1593 {
							gen1588 := Call(__e, ShenFunc(symtl), V373)

							gen1589 := Call(__e, ShenFunc(symtl), gen1588)

							gen1590 := Call(__e, ShenFunc(symcons_2), gen1589)

							var gen1591 Obj
							if True == gen1590 {
								gen1584 := Call(__e, ShenFunc(symtl), V373)

								gen1585 := Call(__e, ShenFunc(symtl), gen1584)

								gen1586 := Call(__e, ShenFunc(symtl), gen1585)

								gen1587 := Call(__e, ShenFunc(sym_a), Nil, gen1586)

								if True == gen1587 {
									gen1591 = True
								} else {
									gen1591 = False
								}

							} else {
								gen1591 = False
							}
							if True == gen1591 {
								gen1594 = True
							} else {
								gen1594 = False
							}

						} else {
							gen1594 = False
						}
						if True == gen1594 {
							gen1597 = True
						} else {
							gen1597 = False
						}

					} else {
						gen1597 = False
					}
					if True == gen1597 {
						gen1599 = True
					} else {
						gen1599 = False
					}

				} else {
					gen1599 = False
				}
				if True == gen1599 {
					gen1573 := Call(__e, ShenFunc(symtl), V373)

					gen1574 := Call(__e, ShenFunc(symhd), gen1573)

					gen1575 := Call(__e, ShenFunc(symtl), V373)

					gen1576 := Call(__e, ShenFunc(symhd), gen1575)

					gen1577 := Call(__e, ShenFunc(symcons), gen1576, V371)

					gen1578 := Call(__e, ShenFunc(symtl), V373)

					gen1579 := Call(__e, ShenFunc(symtl), gen1578)

					gen1580 := Call(__e, ShenFunc(symhd), gen1579)

					gen1581 := Call(__e, ShenFunc(symshen_4assign_1types), gen1577, V372, gen1580)

					gen1582 := Call(__e, ShenFunc(symcons), gen1581, Nil)

					gen1583 := Call(__e, ShenFunc(symcons), gen1574, gen1582)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen1583)

					return

				} else {
					gen1571 := Call(__e, ShenFunc(symcons_2), V373)

					var gen1572 Obj
					if True == gen1571 {
						gen1569 := Call(__e, ShenFunc(symhd), V373)

						gen1570 := Call(__e, ShenFunc(sym_a), MakeSymbol("cond"), gen1569)

						if True == gen1570 {
							gen1572 = True
						} else {
							gen1572 = False
						}

					} else {
						gen1572 = False
					}
					if True == gen1572 {
						gen1566 := MakeNative(func(__e Evaluator) {
							Y := __e.Get(1)
							_ = Y
							gen1560 := Call(__e, ShenFunc(symhd), Y)

							gen1561 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1560)

							gen1562 := Call(__e, ShenFunc(symtl), Y)

							gen1563 := Call(__e, ShenFunc(symhd), gen1562)

							gen1564 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1563)

							gen1565 := Call(__e, ShenFunc(symcons), gen1564, Nil)

							__e.TailApply(ShenFunc(symcons), gen1561, gen1565)

							return

						}, 1)
						gen1567 := Call(__e, ShenFunc(symtl), V373)

						gen1568 := Call(__e, ShenFunc(symmap), gen1566, gen1567)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("cond"), gen1568)

						return

					} else {
						gen1559 := Call(__e, ShenFunc(symcons_2), V373)

						if True == gen1559 {
							gen1550 := Call(__e, ShenFunc(symhd), V373)

							gen1551 := Call(__e, ShenFunc(symshen_4get_1type), gen1550)

							gen1552 := Call(__e, ShenFunc(symtl), V373)

							gen1553 := Call(__e, ShenFunc(symshen_4typextable), gen1551, gen1552)

							NewTable := gen1553
							gen1554 := Call(__e, ShenFunc(symhd), V373)

							gen1556 := MakeNative(func(__e Evaluator) {
								Y := __e.Get(1)
								_ = Y
								gen1555 := Call(__e, ShenFunc(symappend), V372, NewTable)

								__e.TailApply(ShenFunc(symshen_4assign_1types), V371, gen1555, Y)

								return

							}, 1)
							gen1557 := Call(__e, ShenFunc(symtl), V373)

							gen1558 := Call(__e, ShenFunc(symmap), gen1556, gen1557)

							__e.TailApply(ShenFunc(symcons), gen1554, gen1558)

							return

						} else {
							gen1544 := Call(__e, ShenFunc(symassoc), V373, V372)

							AtomType := gen1544
							gen1549 := Call(__e, ShenFunc(symcons_2), AtomType)

							if True == gen1549 {
								gen1546 := Call(__e, ShenFunc(symtl), AtomType)

								gen1547 := Call(__e, ShenFunc(symcons), gen1546, Nil)

								gen1548 := Call(__e, ShenFunc(symcons), V373, gen1547)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1548)

								return

							} else {
								gen1545 := Call(__e, ShenFunc(symelement_2), V373, V371)

								if True == gen1545 {
									__e.Return(V373)
									return
								} else {
									__e.TailApply(ShenFunc(symshen_4atom_1type), V373)

									return
								}

							}

						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assign-types"), gen1639)

		gen1652 := MakeNative(func(__e Evaluator) {
			V375 := __e.Get(1)
			_ = V375
			gen1651 := Call(__e, ShenFunc(symstring_2), V375)

			if True == gen1651 {
				gen1649 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

				gen1650 := Call(__e, ShenFunc(symcons), V375, gen1649)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1650)

				return

			} else {
				gen1648 := Call(__e, ShenFunc(symnumber_2), V375)

				if True == gen1648 {
					gen1646 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

					gen1647 := Call(__e, ShenFunc(symcons), V375, gen1646)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1647)

					return

				} else {
					gen1645 := Call(__e, ShenFunc(symboolean_2), V375)

					if True == gen1645 {
						gen1643 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

						gen1644 := Call(__e, ShenFunc(symcons), V375, gen1643)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1644)

						return

					} else {
						gen1642 := Call(__e, ShenFunc(symsymbol_2), V375)

						if True == gen1642 {
							gen1640 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

							gen1641 := Call(__e, ShenFunc(symcons), V375, gen1640)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1641)

							return

						} else {
							__e.Return(V375)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.atom-type"), gen1652)

		gen1655 := MakeNative(func(__e Evaluator) {
			V380 := __e.Get(1)
			_ = V380
			V381 := __e.Get(2)
			_ = V381
			gen1654 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

			if True == gen1654 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1653 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V380, MakeSymbol("arity"), V381, gen1653)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.store-arity"), gen1655)

		gen1662 := MakeNative(func(__e Evaluator) {
			V383 := __e.Get(1)
			_ = V383
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*teststack*"), Nil)
			gen1656 := Call(__e, ShenFunc(symshen_4reduce__help), V383)

			Result := gen1656
			gen1657 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*teststack*"))

			gen1658 := Call(__e, ShenFunc(symreverse), gen1657)

			gen1659 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tests"), gen1658)

			gen1660 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen1659)

			gen1661 := Call(__e, ShenFunc(symcons), Result, Nil)

			__e.TailApply(ShenFunc(symcons), gen1660, gen1661)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reduce"), gen1662)

		gen2218 := MakeNative(func(__e Evaluator) {
			V385 := __e.Get(1)
			_ = V385
			gen2216 := Call(__e, ShenFunc(symcons_2), V385)

			var gen2217 Obj
			if True == gen2216 {
				gen2213 := Call(__e, ShenFunc(symhd), V385)

				gen2214 := Call(__e, ShenFunc(symcons_2), gen2213)

				var gen2215 Obj
				if True == gen2214 {
					gen2209 := Call(__e, ShenFunc(symhd), V385)

					gen2210 := Call(__e, ShenFunc(symhd), gen2209)

					gen2211 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen2210)

					var gen2212 Obj
					if True == gen2211 {
						gen2205 := Call(__e, ShenFunc(symhd), V385)

						gen2206 := Call(__e, ShenFunc(symtl), gen2205)

						gen2207 := Call(__e, ShenFunc(symcons_2), gen2206)

						var gen2208 Obj
						if True == gen2207 {
							gen2200 := Call(__e, ShenFunc(symhd), V385)

							gen2201 := Call(__e, ShenFunc(symtl), gen2200)

							gen2202 := Call(__e, ShenFunc(symhd), gen2201)

							gen2203 := Call(__e, ShenFunc(symcons_2), gen2202)

							var gen2204 Obj
							if True == gen2203 {
								gen2194 := Call(__e, ShenFunc(symhd), V385)

								gen2195 := Call(__e, ShenFunc(symtl), gen2194)

								gen2196 := Call(__e, ShenFunc(symhd), gen2195)

								gen2197 := Call(__e, ShenFunc(symhd), gen2196)

								gen2198 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen2197)

								var gen2199 Obj
								if True == gen2198 {
									gen2188 := Call(__e, ShenFunc(symhd), V385)

									gen2189 := Call(__e, ShenFunc(symtl), gen2188)

									gen2190 := Call(__e, ShenFunc(symhd), gen2189)

									gen2191 := Call(__e, ShenFunc(symtl), gen2190)

									gen2192 := Call(__e, ShenFunc(symcons_2), gen2191)

									var gen2193 Obj
									if True == gen2192 {
										gen2181 := Call(__e, ShenFunc(symhd), V385)

										gen2182 := Call(__e, ShenFunc(symtl), gen2181)

										gen2183 := Call(__e, ShenFunc(symhd), gen2182)

										gen2184 := Call(__e, ShenFunc(symtl), gen2183)

										gen2185 := Call(__e, ShenFunc(symtl), gen2184)

										gen2186 := Call(__e, ShenFunc(symcons_2), gen2185)

										var gen2187 Obj
										if True == gen2186 {
											gen2173 := Call(__e, ShenFunc(symhd), V385)

											gen2174 := Call(__e, ShenFunc(symtl), gen2173)

											gen2175 := Call(__e, ShenFunc(symhd), gen2174)

											gen2176 := Call(__e, ShenFunc(symtl), gen2175)

											gen2177 := Call(__e, ShenFunc(symtl), gen2176)

											gen2178 := Call(__e, ShenFunc(symtl), gen2177)

											gen2179 := Call(__e, ShenFunc(sym_a), Nil, gen2178)

											var gen2180 Obj
											if True == gen2179 {
												gen2168 := Call(__e, ShenFunc(symhd), V385)

												gen2169 := Call(__e, ShenFunc(symtl), gen2168)

												gen2170 := Call(__e, ShenFunc(symtl), gen2169)

												gen2171 := Call(__e, ShenFunc(symcons_2), gen2170)

												var gen2172 Obj
												if True == gen2171 {
													gen2162 := Call(__e, ShenFunc(symhd), V385)

													gen2163 := Call(__e, ShenFunc(symtl), gen2162)

													gen2164 := Call(__e, ShenFunc(symtl), gen2163)

													gen2165 := Call(__e, ShenFunc(symtl), gen2164)

													gen2166 := Call(__e, ShenFunc(sym_a), Nil, gen2165)

													var gen2167 Obj
													if True == gen2166 {
														gen2159 := Call(__e, ShenFunc(symtl), V385)

														gen2160 := Call(__e, ShenFunc(symcons_2), gen2159)

														var gen2161 Obj
														if True == gen2160 {
															gen2156 := Call(__e, ShenFunc(symtl), V385)

															gen2157 := Call(__e, ShenFunc(symtl), gen2156)

															gen2158 := Call(__e, ShenFunc(sym_a), Nil, gen2157)

															if True == gen2158 {
																gen2161 = True
															} else {
																gen2161 = False
															}

														} else {
															gen2161 = False
														}
														if True == gen2161 {
															gen2167 = True
														} else {
															gen2167 = False
														}

													} else {
														gen2167 = False
													}
													if True == gen2167 {
														gen2172 = True
													} else {
														gen2172 = False
													}

												} else {
													gen2172 = False
												}
												if True == gen2172 {
													gen2180 = True
												} else {
													gen2180 = False
												}

											} else {
												gen2180 = False
											}
											if True == gen2180 {
												gen2187 = True
											} else {
												gen2187 = False
											}

										} else {
											gen2187 = False
										}
										if True == gen2187 {
											gen2193 = True
										} else {
											gen2193 = False
										}

									} else {
										gen2193 = False
									}
									if True == gen2193 {
										gen2199 = True
									} else {
										gen2199 = False
									}

								} else {
									gen2199 = False
								}
								if True == gen2199 {
									gen2204 = True
								} else {
									gen2204 = False
								}

							} else {
								gen2204 = False
							}
							if True == gen2204 {
								gen2208 = True
							} else {
								gen2208 = False
							}

						} else {
							gen2208 = False
						}
						if True == gen2208 {
							gen2212 = True
						} else {
							gen2212 = False
						}

					} else {
						gen2212 = False
					}
					if True == gen2212 {
						gen2215 = True
					} else {
						gen2215 = False
					}

				} else {
					gen2215 = False
				}
				if True == gen2215 {
					gen2217 = True
				} else {
					gen2217 = False
				}

			} else {
				gen2217 = False
			}
			if True == gen2217 {
				gen2119 := Call(__e, ShenFunc(symtl), V385)

				gen2120 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen2119)

				Call(__e, ShenFunc(symshen_4add__test), gen2120)

				gen2121 := Call(__e, ShenFunc(symhd), V385)

				gen2122 := Call(__e, ShenFunc(symtl), gen2121)

				gen2123 := Call(__e, ShenFunc(symhd), gen2122)

				gen2124 := Call(__e, ShenFunc(symtl), gen2123)

				gen2125 := Call(__e, ShenFunc(symhd), gen2124)

				gen2126 := Call(__e, ShenFunc(symhd), V385)

				gen2127 := Call(__e, ShenFunc(symtl), gen2126)

				gen2128 := Call(__e, ShenFunc(symhd), gen2127)

				gen2129 := Call(__e, ShenFunc(symtl), gen2128)

				gen2130 := Call(__e, ShenFunc(symtl), gen2129)

				gen2131 := Call(__e, ShenFunc(symhd), gen2130)

				gen2132 := Call(__e, ShenFunc(symtl), V385)

				gen2133 := Call(__e, ShenFunc(symhd), gen2132)

				gen2134 := Call(__e, ShenFunc(symhd), V385)

				gen2135 := Call(__e, ShenFunc(symtl), gen2134)

				gen2136 := Call(__e, ShenFunc(symhd), gen2135)

				gen2137 := Call(__e, ShenFunc(symhd), V385)

				gen2138 := Call(__e, ShenFunc(symtl), gen2137)

				gen2139 := Call(__e, ShenFunc(symtl), gen2138)

				gen2140 := Call(__e, ShenFunc(symhd), gen2139)

				gen2141 := Call(__e, ShenFunc(symshen_4ebr), gen2133, gen2136, gen2140)

				gen2142 := Call(__e, ShenFunc(symcons), gen2141, Nil)

				gen2143 := Call(__e, ShenFunc(symcons), gen2131, gen2142)

				gen2144 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen2143)

				gen2145 := Call(__e, ShenFunc(symcons), gen2144, Nil)

				gen2146 := Call(__e, ShenFunc(symcons), gen2125, gen2145)

				gen2147 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen2146)

				Abstraction := gen2147
				gen2148 := Call(__e, ShenFunc(symtl), V385)

				gen2149 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen2148)

				gen2150 := Call(__e, ShenFunc(symcons), gen2149, Nil)

				gen2151 := Call(__e, ShenFunc(symcons), Abstraction, gen2150)

				gen2152 := Call(__e, ShenFunc(symtl), V385)

				gen2153 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen2152)

				gen2154 := Call(__e, ShenFunc(symcons), gen2153, Nil)

				gen2155 := Call(__e, ShenFunc(symcons), gen2151, gen2154)

				Application := gen2155
				__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

				return

			} else {
				gen2117 := Call(__e, ShenFunc(symcons_2), V385)

				var gen2118 Obj
				if True == gen2117 {
					gen2114 := Call(__e, ShenFunc(symhd), V385)

					gen2115 := Call(__e, ShenFunc(symcons_2), gen2114)

					var gen2116 Obj
					if True == gen2115 {
						gen2110 := Call(__e, ShenFunc(symhd), V385)

						gen2111 := Call(__e, ShenFunc(symhd), gen2110)

						gen2112 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen2111)

						var gen2113 Obj
						if True == gen2112 {
							gen2106 := Call(__e, ShenFunc(symhd), V385)

							gen2107 := Call(__e, ShenFunc(symtl), gen2106)

							gen2108 := Call(__e, ShenFunc(symcons_2), gen2107)

							var gen2109 Obj
							if True == gen2108 {
								gen2101 := Call(__e, ShenFunc(symhd), V385)

								gen2102 := Call(__e, ShenFunc(symtl), gen2101)

								gen2103 := Call(__e, ShenFunc(symhd), gen2102)

								gen2104 := Call(__e, ShenFunc(symcons_2), gen2103)

								var gen2105 Obj
								if True == gen2104 {
									gen2095 := Call(__e, ShenFunc(symhd), V385)

									gen2096 := Call(__e, ShenFunc(symtl), gen2095)

									gen2097 := Call(__e, ShenFunc(symhd), gen2096)

									gen2098 := Call(__e, ShenFunc(symhd), gen2097)

									gen2099 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), gen2098)

									var gen2100 Obj
									if True == gen2099 {
										gen2089 := Call(__e, ShenFunc(symhd), V385)

										gen2090 := Call(__e, ShenFunc(symtl), gen2089)

										gen2091 := Call(__e, ShenFunc(symhd), gen2090)

										gen2092 := Call(__e, ShenFunc(symtl), gen2091)

										gen2093 := Call(__e, ShenFunc(symcons_2), gen2092)

										var gen2094 Obj
										if True == gen2093 {
											gen2082 := Call(__e, ShenFunc(symhd), V385)

											gen2083 := Call(__e, ShenFunc(symtl), gen2082)

											gen2084 := Call(__e, ShenFunc(symhd), gen2083)

											gen2085 := Call(__e, ShenFunc(symtl), gen2084)

											gen2086 := Call(__e, ShenFunc(symtl), gen2085)

											gen2087 := Call(__e, ShenFunc(symcons_2), gen2086)

											var gen2088 Obj
											if True == gen2087 {
												gen2074 := Call(__e, ShenFunc(symhd), V385)

												gen2075 := Call(__e, ShenFunc(symtl), gen2074)

												gen2076 := Call(__e, ShenFunc(symhd), gen2075)

												gen2077 := Call(__e, ShenFunc(symtl), gen2076)

												gen2078 := Call(__e, ShenFunc(symtl), gen2077)

												gen2079 := Call(__e, ShenFunc(symtl), gen2078)

												gen2080 := Call(__e, ShenFunc(sym_a), Nil, gen2079)

												var gen2081 Obj
												if True == gen2080 {
													gen2069 := Call(__e, ShenFunc(symhd), V385)

													gen2070 := Call(__e, ShenFunc(symtl), gen2069)

													gen2071 := Call(__e, ShenFunc(symtl), gen2070)

													gen2072 := Call(__e, ShenFunc(symcons_2), gen2071)

													var gen2073 Obj
													if True == gen2072 {
														gen2063 := Call(__e, ShenFunc(symhd), V385)

														gen2064 := Call(__e, ShenFunc(symtl), gen2063)

														gen2065 := Call(__e, ShenFunc(symtl), gen2064)

														gen2066 := Call(__e, ShenFunc(symtl), gen2065)

														gen2067 := Call(__e, ShenFunc(sym_a), Nil, gen2066)

														var gen2068 Obj
														if True == gen2067 {
															gen2060 := Call(__e, ShenFunc(symtl), V385)

															gen2061 := Call(__e, ShenFunc(symcons_2), gen2060)

															var gen2062 Obj
															if True == gen2061 {
																gen2057 := Call(__e, ShenFunc(symtl), V385)

																gen2058 := Call(__e, ShenFunc(symtl), gen2057)

																gen2059 := Call(__e, ShenFunc(sym_a), Nil, gen2058)

																if True == gen2059 {
																	gen2062 = True
																} else {
																	gen2062 = False
																}

															} else {
																gen2062 = False
															}
															if True == gen2062 {
																gen2068 = True
															} else {
																gen2068 = False
															}

														} else {
															gen2068 = False
														}
														if True == gen2068 {
															gen2073 = True
														} else {
															gen2073 = False
														}

													} else {
														gen2073 = False
													}
													if True == gen2073 {
														gen2081 = True
													} else {
														gen2081 = False
													}

												} else {
													gen2081 = False
												}
												if True == gen2081 {
													gen2088 = True
												} else {
													gen2088 = False
												}

											} else {
												gen2088 = False
											}
											if True == gen2088 {
												gen2094 = True
											} else {
												gen2094 = False
											}

										} else {
											gen2094 = False
										}
										if True == gen2094 {
											gen2100 = True
										} else {
											gen2100 = False
										}

									} else {
										gen2100 = False
									}
									if True == gen2100 {
										gen2105 = True
									} else {
										gen2105 = False
									}

								} else {
									gen2105 = False
								}
								if True == gen2105 {
									gen2109 = True
								} else {
									gen2109 = False
								}

							} else {
								gen2109 = False
							}
							if True == gen2109 {
								gen2113 = True
							} else {
								gen2113 = False
							}

						} else {
							gen2113 = False
						}
						if True == gen2113 {
							gen2116 = True
						} else {
							gen2116 = False
						}

					} else {
						gen2116 = False
					}
					if True == gen2116 {
						gen2118 = True
					} else {
						gen2118 = False
					}

				} else {
					gen2118 = False
				}
				if True == gen2118 {
					gen2020 := Call(__e, ShenFunc(symtl), V385)

					gen2021 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen2020)

					Call(__e, ShenFunc(symshen_4add__test), gen2021)

					gen2022 := Call(__e, ShenFunc(symhd), V385)

					gen2023 := Call(__e, ShenFunc(symtl), gen2022)

					gen2024 := Call(__e, ShenFunc(symhd), gen2023)

					gen2025 := Call(__e, ShenFunc(symtl), gen2024)

					gen2026 := Call(__e, ShenFunc(symhd), gen2025)

					gen2027 := Call(__e, ShenFunc(symhd), V385)

					gen2028 := Call(__e, ShenFunc(symtl), gen2027)

					gen2029 := Call(__e, ShenFunc(symhd), gen2028)

					gen2030 := Call(__e, ShenFunc(symtl), gen2029)

					gen2031 := Call(__e, ShenFunc(symtl), gen2030)

					gen2032 := Call(__e, ShenFunc(symhd), gen2031)

					gen2033 := Call(__e, ShenFunc(symtl), V385)

					gen2034 := Call(__e, ShenFunc(symhd), gen2033)

					gen2035 := Call(__e, ShenFunc(symhd), V385)

					gen2036 := Call(__e, ShenFunc(symtl), gen2035)

					gen2037 := Call(__e, ShenFunc(symhd), gen2036)

					gen2038 := Call(__e, ShenFunc(symhd), V385)

					gen2039 := Call(__e, ShenFunc(symtl), gen2038)

					gen2040 := Call(__e, ShenFunc(symtl), gen2039)

					gen2041 := Call(__e, ShenFunc(symhd), gen2040)

					gen2042 := Call(__e, ShenFunc(symshen_4ebr), gen2034, gen2037, gen2041)

					gen2043 := Call(__e, ShenFunc(symcons), gen2042, Nil)

					gen2044 := Call(__e, ShenFunc(symcons), gen2032, gen2043)

					gen2045 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen2044)

					gen2046 := Call(__e, ShenFunc(symcons), gen2045, Nil)

					gen2047 := Call(__e, ShenFunc(symcons), gen2026, gen2046)

					gen2048 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen2047)

					Abstraction := gen2048
					gen2049 := Call(__e, ShenFunc(symtl), V385)

					gen2050 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen2049)

					gen2051 := Call(__e, ShenFunc(symcons), gen2050, Nil)

					gen2052 := Call(__e, ShenFunc(symcons), Abstraction, gen2051)

					gen2053 := Call(__e, ShenFunc(symtl), V385)

					gen2054 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen2053)

					gen2055 := Call(__e, ShenFunc(symcons), gen2054, Nil)

					gen2056 := Call(__e, ShenFunc(symcons), gen2052, gen2055)

					Application := gen2056
					__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

					return

				} else {
					gen2018 := Call(__e, ShenFunc(symcons_2), V385)

					var gen2019 Obj
					if True == gen2018 {
						gen2015 := Call(__e, ShenFunc(symhd), V385)

						gen2016 := Call(__e, ShenFunc(symcons_2), gen2015)

						var gen2017 Obj
						if True == gen2016 {
							gen2011 := Call(__e, ShenFunc(symhd), V385)

							gen2012 := Call(__e, ShenFunc(symhd), gen2011)

							gen2013 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen2012)

							var gen2014 Obj
							if True == gen2013 {
								gen2007 := Call(__e, ShenFunc(symhd), V385)

								gen2008 := Call(__e, ShenFunc(symtl), gen2007)

								gen2009 := Call(__e, ShenFunc(symcons_2), gen2008)

								var gen2010 Obj
								if True == gen2009 {
									gen2002 := Call(__e, ShenFunc(symhd), V385)

									gen2003 := Call(__e, ShenFunc(symtl), gen2002)

									gen2004 := Call(__e, ShenFunc(symhd), gen2003)

									gen2005 := Call(__e, ShenFunc(symcons_2), gen2004)

									var gen2006 Obj
									if True == gen2005 {
										gen1996 := Call(__e, ShenFunc(symhd), V385)

										gen1997 := Call(__e, ShenFunc(symtl), gen1996)

										gen1998 := Call(__e, ShenFunc(symhd), gen1997)

										gen1999 := Call(__e, ShenFunc(symhd), gen1998)

										gen2000 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), gen1999)

										var gen2001 Obj
										if True == gen2000 {
											gen1990 := Call(__e, ShenFunc(symhd), V385)

											gen1991 := Call(__e, ShenFunc(symtl), gen1990)

											gen1992 := Call(__e, ShenFunc(symhd), gen1991)

											gen1993 := Call(__e, ShenFunc(symtl), gen1992)

											gen1994 := Call(__e, ShenFunc(symcons_2), gen1993)

											var gen1995 Obj
											if True == gen1994 {
												gen1983 := Call(__e, ShenFunc(symhd), V385)

												gen1984 := Call(__e, ShenFunc(symtl), gen1983)

												gen1985 := Call(__e, ShenFunc(symhd), gen1984)

												gen1986 := Call(__e, ShenFunc(symtl), gen1985)

												gen1987 := Call(__e, ShenFunc(symtl), gen1986)

												gen1988 := Call(__e, ShenFunc(symcons_2), gen1987)

												var gen1989 Obj
												if True == gen1988 {
													gen1975 := Call(__e, ShenFunc(symhd), V385)

													gen1976 := Call(__e, ShenFunc(symtl), gen1975)

													gen1977 := Call(__e, ShenFunc(symhd), gen1976)

													gen1978 := Call(__e, ShenFunc(symtl), gen1977)

													gen1979 := Call(__e, ShenFunc(symtl), gen1978)

													gen1980 := Call(__e, ShenFunc(symtl), gen1979)

													gen1981 := Call(__e, ShenFunc(sym_a), Nil, gen1980)

													var gen1982 Obj
													if True == gen1981 {
														gen1970 := Call(__e, ShenFunc(symhd), V385)

														gen1971 := Call(__e, ShenFunc(symtl), gen1970)

														gen1972 := Call(__e, ShenFunc(symtl), gen1971)

														gen1973 := Call(__e, ShenFunc(symcons_2), gen1972)

														var gen1974 Obj
														if True == gen1973 {
															gen1964 := Call(__e, ShenFunc(symhd), V385)

															gen1965 := Call(__e, ShenFunc(symtl), gen1964)

															gen1966 := Call(__e, ShenFunc(symtl), gen1965)

															gen1967 := Call(__e, ShenFunc(symtl), gen1966)

															gen1968 := Call(__e, ShenFunc(sym_a), Nil, gen1967)

															var gen1969 Obj
															if True == gen1968 {
																gen1961 := Call(__e, ShenFunc(symtl), V385)

																gen1962 := Call(__e, ShenFunc(symcons_2), gen1961)

																var gen1963 Obj
																if True == gen1962 {
																	gen1958 := Call(__e, ShenFunc(symtl), V385)

																	gen1959 := Call(__e, ShenFunc(symtl), gen1958)

																	gen1960 := Call(__e, ShenFunc(sym_a), Nil, gen1959)

																	if True == gen1960 {
																		gen1963 = True
																	} else {
																		gen1963 = False
																	}

																} else {
																	gen1963 = False
																}
																if True == gen1963 {
																	gen1969 = True
																} else {
																	gen1969 = False
																}

															} else {
																gen1969 = False
															}
															if True == gen1969 {
																gen1974 = True
															} else {
																gen1974 = False
															}

														} else {
															gen1974 = False
														}
														if True == gen1974 {
															gen1982 = True
														} else {
															gen1982 = False
														}

													} else {
														gen1982 = False
													}
													if True == gen1982 {
														gen1989 = True
													} else {
														gen1989 = False
													}

												} else {
													gen1989 = False
												}
												if True == gen1989 {
													gen1995 = True
												} else {
													gen1995 = False
												}

											} else {
												gen1995 = False
											}
											if True == gen1995 {
												gen2001 = True
											} else {
												gen2001 = False
											}

										} else {
											gen2001 = False
										}
										if True == gen2001 {
											gen2006 = True
										} else {
											gen2006 = False
										}

									} else {
										gen2006 = False
									}
									if True == gen2006 {
										gen2010 = True
									} else {
										gen2010 = False
									}

								} else {
									gen2010 = False
								}
								if True == gen2010 {
									gen2014 = True
								} else {
									gen2014 = False
								}

							} else {
								gen2014 = False
							}
							if True == gen2014 {
								gen2017 = True
							} else {
								gen2017 = False
							}

						} else {
							gen2017 = False
						}
						if True == gen2017 {
							gen2019 = True
						} else {
							gen2019 = False
						}

					} else {
						gen2019 = False
					}
					if True == gen2019 {
						gen1921 := Call(__e, ShenFunc(symtl), V385)

						gen1922 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.+vector?"), gen1921)

						Call(__e, ShenFunc(symshen_4add__test), gen1922)

						gen1923 := Call(__e, ShenFunc(symhd), V385)

						gen1924 := Call(__e, ShenFunc(symtl), gen1923)

						gen1925 := Call(__e, ShenFunc(symhd), gen1924)

						gen1926 := Call(__e, ShenFunc(symtl), gen1925)

						gen1927 := Call(__e, ShenFunc(symhd), gen1926)

						gen1928 := Call(__e, ShenFunc(symhd), V385)

						gen1929 := Call(__e, ShenFunc(symtl), gen1928)

						gen1930 := Call(__e, ShenFunc(symhd), gen1929)

						gen1931 := Call(__e, ShenFunc(symtl), gen1930)

						gen1932 := Call(__e, ShenFunc(symtl), gen1931)

						gen1933 := Call(__e, ShenFunc(symhd), gen1932)

						gen1934 := Call(__e, ShenFunc(symtl), V385)

						gen1935 := Call(__e, ShenFunc(symhd), gen1934)

						gen1936 := Call(__e, ShenFunc(symhd), V385)

						gen1937 := Call(__e, ShenFunc(symtl), gen1936)

						gen1938 := Call(__e, ShenFunc(symhd), gen1937)

						gen1939 := Call(__e, ShenFunc(symhd), V385)

						gen1940 := Call(__e, ShenFunc(symtl), gen1939)

						gen1941 := Call(__e, ShenFunc(symtl), gen1940)

						gen1942 := Call(__e, ShenFunc(symhd), gen1941)

						gen1943 := Call(__e, ShenFunc(symshen_4ebr), gen1935, gen1938, gen1942)

						gen1944 := Call(__e, ShenFunc(symcons), gen1943, Nil)

						gen1945 := Call(__e, ShenFunc(symcons), gen1933, gen1944)

						gen1946 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1945)

						gen1947 := Call(__e, ShenFunc(symcons), gen1946, Nil)

						gen1948 := Call(__e, ShenFunc(symcons), gen1927, gen1947)

						gen1949 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1948)

						Abstraction := gen1949
						gen1950 := Call(__e, ShenFunc(symtl), V385)

						gen1951 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen1950)

						gen1952 := Call(__e, ShenFunc(symcons), gen1951, Nil)

						gen1953 := Call(__e, ShenFunc(symcons), Abstraction, gen1952)

						gen1954 := Call(__e, ShenFunc(symtl), V385)

						gen1955 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen1954)

						gen1956 := Call(__e, ShenFunc(symcons), gen1955, Nil)

						gen1957 := Call(__e, ShenFunc(symcons), gen1953, gen1956)

						Application := gen1957
						__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

						return

					} else {
						gen1919 := Call(__e, ShenFunc(symcons_2), V385)

						var gen1920 Obj
						if True == gen1919 {
							gen1916 := Call(__e, ShenFunc(symhd), V385)

							gen1917 := Call(__e, ShenFunc(symcons_2), gen1916)

							var gen1918 Obj
							if True == gen1917 {
								gen1912 := Call(__e, ShenFunc(symhd), V385)

								gen1913 := Call(__e, ShenFunc(symhd), gen1912)

								gen1914 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1913)

								var gen1915 Obj
								if True == gen1914 {
									gen1908 := Call(__e, ShenFunc(symhd), V385)

									gen1909 := Call(__e, ShenFunc(symtl), gen1908)

									gen1910 := Call(__e, ShenFunc(symcons_2), gen1909)

									var gen1911 Obj
									if True == gen1910 {
										gen1903 := Call(__e, ShenFunc(symhd), V385)

										gen1904 := Call(__e, ShenFunc(symtl), gen1903)

										gen1905 := Call(__e, ShenFunc(symhd), gen1904)

										gen1906 := Call(__e, ShenFunc(symcons_2), gen1905)

										var gen1907 Obj
										if True == gen1906 {
											gen1897 := Call(__e, ShenFunc(symhd), V385)

											gen1898 := Call(__e, ShenFunc(symtl), gen1897)

											gen1899 := Call(__e, ShenFunc(symhd), gen1898)

											gen1900 := Call(__e, ShenFunc(symhd), gen1899)

											gen1901 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen1900)

											var gen1902 Obj
											if True == gen1901 {
												gen1891 := Call(__e, ShenFunc(symhd), V385)

												gen1892 := Call(__e, ShenFunc(symtl), gen1891)

												gen1893 := Call(__e, ShenFunc(symhd), gen1892)

												gen1894 := Call(__e, ShenFunc(symtl), gen1893)

												gen1895 := Call(__e, ShenFunc(symcons_2), gen1894)

												var gen1896 Obj
												if True == gen1895 {
													gen1884 := Call(__e, ShenFunc(symhd), V385)

													gen1885 := Call(__e, ShenFunc(symtl), gen1884)

													gen1886 := Call(__e, ShenFunc(symhd), gen1885)

													gen1887 := Call(__e, ShenFunc(symtl), gen1886)

													gen1888 := Call(__e, ShenFunc(symtl), gen1887)

													gen1889 := Call(__e, ShenFunc(symcons_2), gen1888)

													var gen1890 Obj
													if True == gen1889 {
														gen1876 := Call(__e, ShenFunc(symhd), V385)

														gen1877 := Call(__e, ShenFunc(symtl), gen1876)

														gen1878 := Call(__e, ShenFunc(symhd), gen1877)

														gen1879 := Call(__e, ShenFunc(symtl), gen1878)

														gen1880 := Call(__e, ShenFunc(symtl), gen1879)

														gen1881 := Call(__e, ShenFunc(symtl), gen1880)

														gen1882 := Call(__e, ShenFunc(sym_a), Nil, gen1881)

														var gen1883 Obj
														if True == gen1882 {
															gen1871 := Call(__e, ShenFunc(symhd), V385)

															gen1872 := Call(__e, ShenFunc(symtl), gen1871)

															gen1873 := Call(__e, ShenFunc(symtl), gen1872)

															gen1874 := Call(__e, ShenFunc(symcons_2), gen1873)

															var gen1875 Obj
															if True == gen1874 {
																gen1865 := Call(__e, ShenFunc(symhd), V385)

																gen1866 := Call(__e, ShenFunc(symtl), gen1865)

																gen1867 := Call(__e, ShenFunc(symtl), gen1866)

																gen1868 := Call(__e, ShenFunc(symtl), gen1867)

																gen1869 := Call(__e, ShenFunc(sym_a), Nil, gen1868)

																var gen1870 Obj
																if True == gen1869 {
																	gen1862 := Call(__e, ShenFunc(symtl), V385)

																	gen1863 := Call(__e, ShenFunc(symcons_2), gen1862)

																	var gen1864 Obj
																	if True == gen1863 {
																		gen1859 := Call(__e, ShenFunc(symtl), V385)

																		gen1860 := Call(__e, ShenFunc(symtl), gen1859)

																		gen1861 := Call(__e, ShenFunc(sym_a), Nil, gen1860)

																		if True == gen1861 {
																			gen1864 = True
																		} else {
																			gen1864 = False
																		}

																	} else {
																		gen1864 = False
																	}
																	if True == gen1864 {
																		gen1870 = True
																	} else {
																		gen1870 = False
																	}

																} else {
																	gen1870 = False
																}
																if True == gen1870 {
																	gen1875 = True
																} else {
																	gen1875 = False
																}

															} else {
																gen1875 = False
															}
															if True == gen1875 {
																gen1883 = True
															} else {
																gen1883 = False
															}

														} else {
															gen1883 = False
														}
														if True == gen1883 {
															gen1890 = True
														} else {
															gen1890 = False
														}

													} else {
														gen1890 = False
													}
													if True == gen1890 {
														gen1896 = True
													} else {
														gen1896 = False
													}

												} else {
													gen1896 = False
												}
												if True == gen1896 {
													gen1902 = True
												} else {
													gen1902 = False
												}

											} else {
												gen1902 = False
											}
											if True == gen1902 {
												gen1907 = True
											} else {
												gen1907 = False
											}

										} else {
											gen1907 = False
										}
										if True == gen1907 {
											gen1911 = True
										} else {
											gen1911 = False
										}

									} else {
										gen1911 = False
									}
									if True == gen1911 {
										gen1915 = True
									} else {
										gen1915 = False
									}

								} else {
									gen1915 = False
								}
								if True == gen1915 {
									gen1918 = True
								} else {
									gen1918 = False
								}

							} else {
								gen1918 = False
							}
							if True == gen1918 {
								gen1920 = True
							} else {
								gen1920 = False
							}

						} else {
							gen1920 = False
						}
						if True == gen1920 {
							gen1819 := Call(__e, ShenFunc(symtl), V385)

							gen1820 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.+string?"), gen1819)

							Call(__e, ShenFunc(symshen_4add__test), gen1820)

							gen1821 := Call(__e, ShenFunc(symhd), V385)

							gen1822 := Call(__e, ShenFunc(symtl), gen1821)

							gen1823 := Call(__e, ShenFunc(symhd), gen1822)

							gen1824 := Call(__e, ShenFunc(symtl), gen1823)

							gen1825 := Call(__e, ShenFunc(symhd), gen1824)

							gen1826 := Call(__e, ShenFunc(symhd), V385)

							gen1827 := Call(__e, ShenFunc(symtl), gen1826)

							gen1828 := Call(__e, ShenFunc(symhd), gen1827)

							gen1829 := Call(__e, ShenFunc(symtl), gen1828)

							gen1830 := Call(__e, ShenFunc(symtl), gen1829)

							gen1831 := Call(__e, ShenFunc(symhd), gen1830)

							gen1832 := Call(__e, ShenFunc(symtl), V385)

							gen1833 := Call(__e, ShenFunc(symhd), gen1832)

							gen1834 := Call(__e, ShenFunc(symhd), V385)

							gen1835 := Call(__e, ShenFunc(symtl), gen1834)

							gen1836 := Call(__e, ShenFunc(symhd), gen1835)

							gen1837 := Call(__e, ShenFunc(symhd), V385)

							gen1838 := Call(__e, ShenFunc(symtl), gen1837)

							gen1839 := Call(__e, ShenFunc(symtl), gen1838)

							gen1840 := Call(__e, ShenFunc(symhd), gen1839)

							gen1841 := Call(__e, ShenFunc(symshen_4ebr), gen1833, gen1836, gen1840)

							gen1842 := Call(__e, ShenFunc(symcons), gen1841, Nil)

							gen1843 := Call(__e, ShenFunc(symcons), gen1831, gen1842)

							gen1844 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1843)

							gen1845 := Call(__e, ShenFunc(symcons), gen1844, Nil)

							gen1846 := Call(__e, ShenFunc(symcons), gen1825, gen1845)

							gen1847 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1846)

							Abstraction := gen1847
							gen1848 := Call(__e, ShenFunc(symtl), V385)

							gen1849 := Call(__e, ShenFunc(symhd), gen1848)

							gen1850 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

							gen1851 := Call(__e, ShenFunc(symcons), gen1849, gen1850)

							gen1852 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen1851)

							gen1853 := Call(__e, ShenFunc(symcons), gen1852, Nil)

							gen1854 := Call(__e, ShenFunc(symcons), Abstraction, gen1853)

							gen1855 := Call(__e, ShenFunc(symtl), V385)

							gen1856 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen1855)

							gen1857 := Call(__e, ShenFunc(symcons), gen1856, Nil)

							gen1858 := Call(__e, ShenFunc(symcons), gen1854, gen1857)

							Application := gen1858
							__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

							return

						} else {
							gen1817 := Call(__e, ShenFunc(symcons_2), V385)

							var gen1818 Obj
							if True == gen1817 {
								gen1814 := Call(__e, ShenFunc(symhd), V385)

								gen1815 := Call(__e, ShenFunc(symcons_2), gen1814)

								var gen1816 Obj
								if True == gen1815 {
									gen1810 := Call(__e, ShenFunc(symhd), V385)

									gen1811 := Call(__e, ShenFunc(symhd), gen1810)

									gen1812 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1811)

									var gen1813 Obj
									if True == gen1812 {
										gen1806 := Call(__e, ShenFunc(symhd), V385)

										gen1807 := Call(__e, ShenFunc(symtl), gen1806)

										gen1808 := Call(__e, ShenFunc(symcons_2), gen1807)

										var gen1809 Obj
										if True == gen1808 {
											gen1801 := Call(__e, ShenFunc(symhd), V385)

											gen1802 := Call(__e, ShenFunc(symtl), gen1801)

											gen1803 := Call(__e, ShenFunc(symhd), gen1802)

											gen1804 := Call(__e, ShenFunc(symcons_2), gen1803)

											var gen1805 Obj
											if True == gen1804 {
												gen1796 := Call(__e, ShenFunc(symhd), V385)

												gen1797 := Call(__e, ShenFunc(symtl), gen1796)

												gen1798 := Call(__e, ShenFunc(symtl), gen1797)

												gen1799 := Call(__e, ShenFunc(symcons_2), gen1798)

												var gen1800 Obj
												if True == gen1799 {
													gen1790 := Call(__e, ShenFunc(symhd), V385)

													gen1791 := Call(__e, ShenFunc(symtl), gen1790)

													gen1792 := Call(__e, ShenFunc(symtl), gen1791)

													gen1793 := Call(__e, ShenFunc(symtl), gen1792)

													gen1794 := Call(__e, ShenFunc(sym_a), Nil, gen1793)

													var gen1795 Obj
													if True == gen1794 {
														gen1787 := Call(__e, ShenFunc(symtl), V385)

														gen1788 := Call(__e, ShenFunc(symcons_2), gen1787)

														var gen1789 Obj
														if True == gen1788 {
															gen1784 := Call(__e, ShenFunc(symtl), V385)

															gen1785 := Call(__e, ShenFunc(symtl), gen1784)

															gen1786 := Call(__e, ShenFunc(sym_a), Nil, gen1785)

															if True == gen1786 {
																gen1789 = True
															} else {
																gen1789 = False
															}

														} else {
															gen1789 = False
														}
														if True == gen1789 {
															gen1795 = True
														} else {
															gen1795 = False
														}

													} else {
														gen1795 = False
													}
													if True == gen1795 {
														gen1800 = True
													} else {
														gen1800 = False
													}

												} else {
													gen1800 = False
												}
												if True == gen1800 {
													gen1805 = True
												} else {
													gen1805 = False
												}

											} else {
												gen1805 = False
											}
											if True == gen1805 {
												gen1809 = True
											} else {
												gen1809 = False
											}

										} else {
											gen1809 = False
										}
										if True == gen1809 {
											gen1813 = True
										} else {
											gen1813 = False
										}

									} else {
										gen1813 = False
									}
									if True == gen1813 {
										gen1816 = True
									} else {
										gen1816 = False
									}

								} else {
									gen1816 = False
								}
								if True == gen1816 {
									gen1818 = True
								} else {
									gen1818 = False
								}

							} else {
								gen1818 = False
							}
							if True == gen1818 {
								__e.TailApply(ShenFunc(symshen_4custom_1pattern_1reducer), V385)

								return
							} else {
								gen1782 := Call(__e, ShenFunc(symcons_2), V385)

								var gen1783 Obj
								if True == gen1782 {
									gen1779 := Call(__e, ShenFunc(symhd), V385)

									gen1780 := Call(__e, ShenFunc(symcons_2), gen1779)

									var gen1781 Obj
									if True == gen1780 {
										gen1775 := Call(__e, ShenFunc(symhd), V385)

										gen1776 := Call(__e, ShenFunc(symhd), gen1775)

										gen1777 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1776)

										var gen1778 Obj
										if True == gen1777 {
											gen1771 := Call(__e, ShenFunc(symhd), V385)

											gen1772 := Call(__e, ShenFunc(symtl), gen1771)

											gen1773 := Call(__e, ShenFunc(symcons_2), gen1772)

											var gen1774 Obj
											if True == gen1773 {
												gen1766 := Call(__e, ShenFunc(symhd), V385)

												gen1767 := Call(__e, ShenFunc(symtl), gen1766)

												gen1768 := Call(__e, ShenFunc(symtl), gen1767)

												gen1769 := Call(__e, ShenFunc(symcons_2), gen1768)

												var gen1770 Obj
												if True == gen1769 {
													gen1760 := Call(__e, ShenFunc(symhd), V385)

													gen1761 := Call(__e, ShenFunc(symtl), gen1760)

													gen1762 := Call(__e, ShenFunc(symtl), gen1761)

													gen1763 := Call(__e, ShenFunc(symtl), gen1762)

													gen1764 := Call(__e, ShenFunc(sym_a), Nil, gen1763)

													var gen1765 Obj
													if True == gen1764 {
														gen1757 := Call(__e, ShenFunc(symtl), V385)

														gen1758 := Call(__e, ShenFunc(symcons_2), gen1757)

														var gen1759 Obj
														if True == gen1758 {
															gen1753 := Call(__e, ShenFunc(symtl), V385)

															gen1754 := Call(__e, ShenFunc(symtl), gen1753)

															gen1755 := Call(__e, ShenFunc(sym_a), Nil, gen1754)

															var gen1756 Obj
															if True == gen1755 {
																gen1748 := Call(__e, ShenFunc(symhd), V385)

																gen1749 := Call(__e, ShenFunc(symtl), gen1748)

																gen1750 := Call(__e, ShenFunc(symhd), gen1749)

																gen1751 := Call(__e, ShenFunc(symvariable_2), gen1750)

																gen1752 := Call(__e, ShenFunc(symnot), gen1751)

																if True == gen1752 {
																	gen1756 = True
																} else {
																	gen1756 = False
																}

															} else {
																gen1756 = False
															}
															if True == gen1756 {
																gen1759 = True
															} else {
																gen1759 = False
															}

														} else {
															gen1759 = False
														}
														if True == gen1759 {
															gen1765 = True
														} else {
															gen1765 = False
														}

													} else {
														gen1765 = False
													}
													if True == gen1765 {
														gen1770 = True
													} else {
														gen1770 = False
													}

												} else {
													gen1770 = False
												}
												if True == gen1770 {
													gen1774 = True
												} else {
													gen1774 = False
												}

											} else {
												gen1774 = False
											}
											if True == gen1774 {
												gen1778 = True
											} else {
												gen1778 = False
											}

										} else {
											gen1778 = False
										}
										if True == gen1778 {
											gen1781 = True
										} else {
											gen1781 = False
										}

									} else {
										gen1781 = False
									}
									if True == gen1781 {
										gen1783 = True
									} else {
										gen1783 = False
									}

								} else {
									gen1783 = False
								}
								if True == gen1783 {
									gen1738 := Call(__e, ShenFunc(symhd), V385)

									gen1739 := Call(__e, ShenFunc(symtl), gen1738)

									gen1740 := Call(__e, ShenFunc(symhd), gen1739)

									gen1741 := Call(__e, ShenFunc(symtl), V385)

									gen1742 := Call(__e, ShenFunc(symcons), gen1740, gen1741)

									gen1743 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen1742)

									Call(__e, ShenFunc(symshen_4add__test), gen1743)

									gen1744 := Call(__e, ShenFunc(symhd), V385)

									gen1745 := Call(__e, ShenFunc(symtl), gen1744)

									gen1746 := Call(__e, ShenFunc(symtl), gen1745)

									gen1747 := Call(__e, ShenFunc(symhd), gen1746)

									__e.TailApply(ShenFunc(symshen_4reduce__help), gen1747)

									return

								} else {
									gen1736 := Call(__e, ShenFunc(symcons_2), V385)

									var gen1737 Obj
									if True == gen1736 {
										gen1733 := Call(__e, ShenFunc(symhd), V385)

										gen1734 := Call(__e, ShenFunc(symcons_2), gen1733)

										var gen1735 Obj
										if True == gen1734 {
											gen1729 := Call(__e, ShenFunc(symhd), V385)

											gen1730 := Call(__e, ShenFunc(symhd), gen1729)

											gen1731 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1730)

											var gen1732 Obj
											if True == gen1731 {
												gen1725 := Call(__e, ShenFunc(symhd), V385)

												gen1726 := Call(__e, ShenFunc(symtl), gen1725)

												gen1727 := Call(__e, ShenFunc(symcons_2), gen1726)

												var gen1728 Obj
												if True == gen1727 {
													gen1720 := Call(__e, ShenFunc(symhd), V385)

													gen1721 := Call(__e, ShenFunc(symtl), gen1720)

													gen1722 := Call(__e, ShenFunc(symtl), gen1721)

													gen1723 := Call(__e, ShenFunc(symcons_2), gen1722)

													var gen1724 Obj
													if True == gen1723 {
														gen1714 := Call(__e, ShenFunc(symhd), V385)

														gen1715 := Call(__e, ShenFunc(symtl), gen1714)

														gen1716 := Call(__e, ShenFunc(symtl), gen1715)

														gen1717 := Call(__e, ShenFunc(symtl), gen1716)

														gen1718 := Call(__e, ShenFunc(sym_a), Nil, gen1717)

														var gen1719 Obj
														if True == gen1718 {
															gen1711 := Call(__e, ShenFunc(symtl), V385)

															gen1712 := Call(__e, ShenFunc(symcons_2), gen1711)

															var gen1713 Obj
															if True == gen1712 {
																gen1708 := Call(__e, ShenFunc(symtl), V385)

																gen1709 := Call(__e, ShenFunc(symtl), gen1708)

																gen1710 := Call(__e, ShenFunc(sym_a), Nil, gen1709)

																if True == gen1710 {
																	gen1713 = True
																} else {
																	gen1713 = False
																}

															} else {
																gen1713 = False
															}
															if True == gen1713 {
																gen1719 = True
															} else {
																gen1719 = False
															}

														} else {
															gen1719 = False
														}
														if True == gen1719 {
															gen1724 = True
														} else {
															gen1724 = False
														}

													} else {
														gen1724 = False
													}
													if True == gen1724 {
														gen1728 = True
													} else {
														gen1728 = False
													}

												} else {
													gen1728 = False
												}
												if True == gen1728 {
													gen1732 = True
												} else {
													gen1732 = False
												}

											} else {
												gen1732 = False
											}
											if True == gen1732 {
												gen1735 = True
											} else {
												gen1735 = False
											}

										} else {
											gen1735 = False
										}
										if True == gen1735 {
											gen1737 = True
										} else {
											gen1737 = False
										}

									} else {
										gen1737 = False
									}
									if True == gen1737 {
										gen1698 := Call(__e, ShenFunc(symtl), V385)

										gen1699 := Call(__e, ShenFunc(symhd), gen1698)

										gen1700 := Call(__e, ShenFunc(symhd), V385)

										gen1701 := Call(__e, ShenFunc(symtl), gen1700)

										gen1702 := Call(__e, ShenFunc(symhd), gen1701)

										gen1703 := Call(__e, ShenFunc(symhd), V385)

										gen1704 := Call(__e, ShenFunc(symtl), gen1703)

										gen1705 := Call(__e, ShenFunc(symtl), gen1704)

										gen1706 := Call(__e, ShenFunc(symhd), gen1705)

										gen1707 := Call(__e, ShenFunc(symshen_4ebr), gen1699, gen1702, gen1706)

										__e.TailApply(ShenFunc(symshen_4reduce__help), gen1707)

										return

									} else {
										gen1696 := Call(__e, ShenFunc(symcons_2), V385)

										var gen1697 Obj
										if True == gen1696 {
											gen1693 := Call(__e, ShenFunc(symhd), V385)

											gen1694 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen1693)

											var gen1695 Obj
											if True == gen1694 {
												gen1690 := Call(__e, ShenFunc(symtl), V385)

												gen1691 := Call(__e, ShenFunc(symcons_2), gen1690)

												var gen1692 Obj
												if True == gen1691 {
													gen1686 := Call(__e, ShenFunc(symtl), V385)

													gen1687 := Call(__e, ShenFunc(symtl), gen1686)

													gen1688 := Call(__e, ShenFunc(symcons_2), gen1687)

													var gen1689 Obj
													if True == gen1688 {
														gen1682 := Call(__e, ShenFunc(symtl), V385)

														gen1683 := Call(__e, ShenFunc(symtl), gen1682)

														gen1684 := Call(__e, ShenFunc(symtl), gen1683)

														gen1685 := Call(__e, ShenFunc(sym_a), Nil, gen1684)

														if True == gen1685 {
															gen1689 = True
														} else {
															gen1689 = False
														}

													} else {
														gen1689 = False
													}
													if True == gen1689 {
														gen1692 = True
													} else {
														gen1692 = False
													}

												} else {
													gen1692 = False
												}
												if True == gen1692 {
													gen1695 = True
												} else {
													gen1695 = False
												}

											} else {
												gen1695 = False
											}
											if True == gen1695 {
												gen1697 = True
											} else {
												gen1697 = False
											}

										} else {
											gen1697 = False
										}
										if True == gen1697 {
											gen1677 := Call(__e, ShenFunc(symtl), V385)

											gen1678 := Call(__e, ShenFunc(symhd), gen1677)

											Call(__e, ShenFunc(symshen_4add__test), gen1678)

											gen1679 := Call(__e, ShenFunc(symtl), V385)

											gen1680 := Call(__e, ShenFunc(symtl), gen1679)

											gen1681 := Call(__e, ShenFunc(symhd), gen1680)

											__e.TailApply(ShenFunc(symshen_4reduce__help), gen1681)

											return

										} else {
											gen1675 := Call(__e, ShenFunc(symcons_2), V385)

											var gen1676 Obj
											if True == gen1675 {
												gen1672 := Call(__e, ShenFunc(symtl), V385)

												gen1673 := Call(__e, ShenFunc(symcons_2), gen1672)

												var gen1674 Obj
												if True == gen1673 {
													gen1669 := Call(__e, ShenFunc(symtl), V385)

													gen1670 := Call(__e, ShenFunc(symtl), gen1669)

													gen1671 := Call(__e, ShenFunc(sym_a), Nil, gen1670)

													if True == gen1671 {
														gen1674 = True
													} else {
														gen1674 = False
													}

												} else {
													gen1674 = False
												}
												if True == gen1674 {
													gen1676 = True
												} else {
													gen1676 = False
												}

											} else {
												gen1676 = False
											}
											if True == gen1676 {
												gen1663 := Call(__e, ShenFunc(symhd), V385)

												gen1664 := Call(__e, ShenFunc(symshen_4reduce__help), gen1663)

												Z := gen1664
												gen1667 := Call(__e, ShenFunc(symhd), V385)

												gen1668 := Call(__e, ShenFunc(sym_a), gen1667, Z)

												if True == gen1668 {
													__e.Return(V385)
													return
												} else {
													gen1665 := Call(__e, ShenFunc(symtl), V385)

													gen1666 := Call(__e, ShenFunc(symcons), Z, gen1665)

													__e.TailApply(ShenFunc(symshen_4reduce__help), gen1666)

													return

												}

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reduce_help"), gen2218)

		gen2220 := MakeNative(func(__e Evaluator) {
			V387 := __e.Get(1)
			_ = V387
			gen2219 := Call(__e, ShenFunc(sym_a), MakeString(""), V387)

			if True == gen2219 {
				__e.Return(False)
				return
			} else {
				__e.TailApply(ShenFunc(symstring_2), V387)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.+string?"), gen2220)

		gen2224 := MakeNative(func(__e Evaluator) {
			V389 := __e.Get(1)
			_ = V389
			gen2223 := Call(__e, ShenFunc(symabsvector_2), V389)

			if True == gen2223 {
				gen2221 := Call(__e, ShenFunc(sym_5_1address), V389, MakeNumber(0))

				gen2222 := Call(__e, ShenFunc(sym_6), gen2221, MakeNumber(0))

				if True == gen2222 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.+vector?"), gen2224)

		gen2288 := MakeNative(func(__e Evaluator) {
			V402 := __e.Get(1)
			_ = V402
			V403 := __e.Get(2)
			_ = V403
			V404 := __e.Get(3)
			_ = V404
			gen2287 := Call(__e, ShenFunc(sym_a), V404, V403)

			if True == gen2287 {
				__e.Return(V402)
				return
			} else {
				gen2285 := Call(__e, ShenFunc(symcons_2), V404)

				var gen2286 Obj
				if True == gen2285 {
					gen2282 := Call(__e, ShenFunc(symhd), V404)

					gen2283 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen2282)

					var gen2284 Obj
					if True == gen2283 {
						gen2279 := Call(__e, ShenFunc(symtl), V404)

						gen2280 := Call(__e, ShenFunc(symcons_2), gen2279)

						var gen2281 Obj
						if True == gen2280 {
							gen2275 := Call(__e, ShenFunc(symtl), V404)

							gen2276 := Call(__e, ShenFunc(symtl), gen2275)

							gen2277 := Call(__e, ShenFunc(symcons_2), gen2276)

							var gen2278 Obj
							if True == gen2277 {
								gen2270 := Call(__e, ShenFunc(symtl), V404)

								gen2271 := Call(__e, ShenFunc(symtl), gen2270)

								gen2272 := Call(__e, ShenFunc(symtl), gen2271)

								gen2273 := Call(__e, ShenFunc(sym_a), Nil, gen2272)

								var gen2274 Obj
								if True == gen2273 {
									gen2267 := Call(__e, ShenFunc(symtl), V404)

									gen2268 := Call(__e, ShenFunc(symhd), gen2267)

									gen2269 := Call(__e, ShenFunc(symshen_4clash_2), gen2268, V403)

									if True == gen2269 {
										gen2274 = True
									} else {
										gen2274 = False
									}

								} else {
									gen2274 = False
								}
								if True == gen2274 {
									gen2278 = True
								} else {
									gen2278 = False
								}

							} else {
								gen2278 = False
							}
							if True == gen2278 {
								gen2281 = True
							} else {
								gen2281 = False
							}

						} else {
							gen2281 = False
						}
						if True == gen2281 {
							gen2284 = True
						} else {
							gen2284 = False
						}

					} else {
						gen2284 = False
					}
					if True == gen2284 {
						gen2286 = True
					} else {
						gen2286 = False
					}

				} else {
					gen2286 = False
				}
				if True == gen2286 {
					__e.Return(V404)
					return
				} else {
					gen2265 := Call(__e, ShenFunc(symcons_2), V404)

					var gen2266 Obj
					if True == gen2265 {
						gen2262 := Call(__e, ShenFunc(symhd), V404)

						gen2263 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen2262)

						var gen2264 Obj
						if True == gen2263 {
							gen2259 := Call(__e, ShenFunc(symtl), V404)

							gen2260 := Call(__e, ShenFunc(symcons_2), gen2259)

							var gen2261 Obj
							if True == gen2260 {
								gen2255 := Call(__e, ShenFunc(symtl), V404)

								gen2256 := Call(__e, ShenFunc(symtl), gen2255)

								gen2257 := Call(__e, ShenFunc(symcons_2), gen2256)

								var gen2258 Obj
								if True == gen2257 {
									gen2250 := Call(__e, ShenFunc(symtl), V404)

									gen2251 := Call(__e, ShenFunc(symtl), gen2250)

									gen2252 := Call(__e, ShenFunc(symtl), gen2251)

									gen2253 := Call(__e, ShenFunc(symcons_2), gen2252)

									var gen2254 Obj
									if True == gen2253 {
										gen2244 := Call(__e, ShenFunc(symtl), V404)

										gen2245 := Call(__e, ShenFunc(symtl), gen2244)

										gen2246 := Call(__e, ShenFunc(symtl), gen2245)

										gen2247 := Call(__e, ShenFunc(symtl), gen2246)

										gen2248 := Call(__e, ShenFunc(sym_a), Nil, gen2247)

										var gen2249 Obj
										if True == gen2248 {
											gen2241 := Call(__e, ShenFunc(symtl), V404)

											gen2242 := Call(__e, ShenFunc(symhd), gen2241)

											gen2243 := Call(__e, ShenFunc(symshen_4clash_2), gen2242, V403)

											if True == gen2243 {
												gen2249 = True
											} else {
												gen2249 = False
											}

										} else {
											gen2249 = False
										}
										if True == gen2249 {
											gen2254 = True
										} else {
											gen2254 = False
										}

									} else {
										gen2254 = False
									}
									if True == gen2254 {
										gen2258 = True
									} else {
										gen2258 = False
									}

								} else {
									gen2258 = False
								}
								if True == gen2258 {
									gen2261 = True
								} else {
									gen2261 = False
								}

							} else {
								gen2261 = False
							}
							if True == gen2261 {
								gen2264 = True
							} else {
								gen2264 = False
							}

						} else {
							gen2264 = False
						}
						if True == gen2264 {
							gen2266 = True
						} else {
							gen2266 = False
						}

					} else {
						gen2266 = False
					}
					if True == gen2266 {
						gen2230 := Call(__e, ShenFunc(symtl), V404)

						gen2231 := Call(__e, ShenFunc(symhd), gen2230)

						gen2232 := Call(__e, ShenFunc(symtl), V404)

						gen2233 := Call(__e, ShenFunc(symtl), gen2232)

						gen2234 := Call(__e, ShenFunc(symhd), gen2233)

						gen2235 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen2234)

						gen2236 := Call(__e, ShenFunc(symtl), V404)

						gen2237 := Call(__e, ShenFunc(symtl), gen2236)

						gen2238 := Call(__e, ShenFunc(symtl), gen2237)

						gen2239 := Call(__e, ShenFunc(symcons), gen2235, gen2238)

						gen2240 := Call(__e, ShenFunc(symcons), gen2231, gen2239)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen2240)

						return

					} else {
						gen2229 := Call(__e, ShenFunc(symcons_2), V404)

						if True == gen2229 {
							gen2225 := Call(__e, ShenFunc(symhd), V404)

							gen2226 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen2225)

							gen2227 := Call(__e, ShenFunc(symtl), V404)

							gen2228 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen2227)

							__e.TailApply(ShenFunc(symcons), gen2226, gen2228)

							return

						} else {
							__e.Return(V404)
							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ebr"), gen2288)

		gen2295 := MakeNative(func(__e Evaluator) {
			V416 := __e.Get(1)
			_ = V416
			V417 := __e.Get(2)
			_ = V417
			gen2294 := Call(__e, ShenFunc(sym_a), V417, V416)

			if True == gen2294 {
				__e.Return(True)
				return
			} else {
				gen2293 := Call(__e, ShenFunc(symcons_2), V417)

				if True == gen2293 {
					gen2291 := Call(__e, ShenFunc(symhd), V417)

					gen2292 := Call(__e, ShenFunc(symshen_4clash_2), V416, gen2291)

					if True == gen2292 {
						__e.Return(True)
						return
					} else {
						gen2289 := Call(__e, ShenFunc(symtl), V417)

						gen2290 := Call(__e, ShenFunc(symshen_4clash_2), V416, gen2289)

						if True == gen2290 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clash?"), gen2295)

		gen2298 := MakeNative(func(__e Evaluator) {
			V419 := __e.Get(1)
			_ = V419
			gen2296 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*teststack*"))

			gen2297 := Call(__e, ShenFunc(symcons), V419, gen2296)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*teststack*"), gen2297)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.add_test"), gen2298)

		gen2302 := MakeNative(func(__e Evaluator) {
			V423 := __e.Get(1)
			_ = V423
			V424 := __e.Get(2)
			_ = V424
			V425 := __e.Get(3)
			_ = V425
			gen2299 := Call(__e, ShenFunc(symshen_4err_1condition), V423)

			Err := gen2299
			gen2300 := Call(__e, ShenFunc(symshen_4case_1form), V425, Err)

			Cases := gen2300
			gen2301 := Call(__e, ShenFunc(symshen_4encode_1choices), Cases, V423)

			EncodeChoices := gen2301
			__e.TailApply(ShenFunc(symshen_4cond_1form), EncodeChoices)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cond-expression"), gen2302)

		gen2322 := MakeNative(func(__e Evaluator) {
			V429 := __e.Get(1)
			_ = V429
			gen2320 := Call(__e, ShenFunc(symcons_2), V429)

			var gen2321 Obj
			if True == gen2320 {
				gen2317 := Call(__e, ShenFunc(symhd), V429)

				gen2318 := Call(__e, ShenFunc(symcons_2), gen2317)

				var gen2319 Obj
				if True == gen2318 {
					gen2313 := Call(__e, ShenFunc(symhd), V429)

					gen2314 := Call(__e, ShenFunc(symhd), gen2313)

					gen2315 := Call(__e, ShenFunc(sym_a), True, gen2314)

					var gen2316 Obj
					if True == gen2315 {
						gen2309 := Call(__e, ShenFunc(symhd), V429)

						gen2310 := Call(__e, ShenFunc(symtl), gen2309)

						gen2311 := Call(__e, ShenFunc(symcons_2), gen2310)

						var gen2312 Obj
						if True == gen2311 {
							gen2305 := Call(__e, ShenFunc(symhd), V429)

							gen2306 := Call(__e, ShenFunc(symtl), gen2305)

							gen2307 := Call(__e, ShenFunc(symtl), gen2306)

							gen2308 := Call(__e, ShenFunc(sym_a), Nil, gen2307)

							if True == gen2308 {
								gen2312 = True
							} else {
								gen2312 = False
							}

						} else {
							gen2312 = False
						}
						if True == gen2312 {
							gen2316 = True
						} else {
							gen2316 = False
						}

					} else {
						gen2316 = False
					}
					if True == gen2316 {
						gen2319 = True
					} else {
						gen2319 = False
					}

				} else {
					gen2319 = False
				}
				if True == gen2319 {
					gen2321 = True
				} else {
					gen2321 = False
				}

			} else {
				gen2321 = False
			}
			if True == gen2321 {
				gen2303 := Call(__e, ShenFunc(symhd), V429)

				gen2304 := Call(__e, ShenFunc(symtl), gen2303)

				__e.TailApply(ShenFunc(symhd), gen2304)

				return

			} else {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("cond"), V429)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cond-form"), gen2322)

		gen2545 := MakeNative(func(__e Evaluator) {
			V434 := __e.Get(1)
			_ = V434
			V435 := __e.Get(2)
			_ = V435
			gen2544 := Call(__e, ShenFunc(sym_a), Nil, V434)

			if True == gen2544 {
				__e.Return(Nil)
				return
			} else {
				gen2542 := Call(__e, ShenFunc(symcons_2), V434)

				var gen2543 Obj
				if True == gen2542 {
					gen2539 := Call(__e, ShenFunc(symhd), V434)

					gen2540 := Call(__e, ShenFunc(symcons_2), gen2539)

					var gen2541 Obj
					if True == gen2540 {
						gen2535 := Call(__e, ShenFunc(symhd), V434)

						gen2536 := Call(__e, ShenFunc(symhd), gen2535)

						gen2537 := Call(__e, ShenFunc(sym_a), True, gen2536)

						var gen2538 Obj
						if True == gen2537 {
							gen2531 := Call(__e, ShenFunc(symhd), V434)

							gen2532 := Call(__e, ShenFunc(symtl), gen2531)

							gen2533 := Call(__e, ShenFunc(symcons_2), gen2532)

							var gen2534 Obj
							if True == gen2533 {
								gen2526 := Call(__e, ShenFunc(symhd), V434)

								gen2527 := Call(__e, ShenFunc(symtl), gen2526)

								gen2528 := Call(__e, ShenFunc(symhd), gen2527)

								gen2529 := Call(__e, ShenFunc(symcons_2), gen2528)

								var gen2530 Obj
								if True == gen2529 {
									gen2520 := Call(__e, ShenFunc(symhd), V434)

									gen2521 := Call(__e, ShenFunc(symtl), gen2520)

									gen2522 := Call(__e, ShenFunc(symhd), gen2521)

									gen2523 := Call(__e, ShenFunc(symhd), gen2522)

									gen2524 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2523)

									var gen2525 Obj
									if True == gen2524 {
										gen2514 := Call(__e, ShenFunc(symhd), V434)

										gen2515 := Call(__e, ShenFunc(symtl), gen2514)

										gen2516 := Call(__e, ShenFunc(symhd), gen2515)

										gen2517 := Call(__e, ShenFunc(symtl), gen2516)

										gen2518 := Call(__e, ShenFunc(symcons_2), gen2517)

										var gen2519 Obj
										if True == gen2518 {
											gen2507 := Call(__e, ShenFunc(symhd), V434)

											gen2508 := Call(__e, ShenFunc(symtl), gen2507)

											gen2509 := Call(__e, ShenFunc(symhd), gen2508)

											gen2510 := Call(__e, ShenFunc(symtl), gen2509)

											gen2511 := Call(__e, ShenFunc(symtl), gen2510)

											gen2512 := Call(__e, ShenFunc(sym_a), Nil, gen2511)

											var gen2513 Obj
											if True == gen2512 {
												gen2502 := Call(__e, ShenFunc(symhd), V434)

												gen2503 := Call(__e, ShenFunc(symtl), gen2502)

												gen2504 := Call(__e, ShenFunc(symtl), gen2503)

												gen2505 := Call(__e, ShenFunc(sym_a), Nil, gen2504)

												var gen2506 Obj
												if True == gen2505 {
													gen2500 := Call(__e, ShenFunc(symtl), V434)

													gen2501 := Call(__e, ShenFunc(sym_a), Nil, gen2500)

													if True == gen2501 {
														gen2506 = True
													} else {
														gen2506 = False
													}

												} else {
													gen2506 = False
												}
												if True == gen2506 {
													gen2513 = True
												} else {
													gen2513 = False
												}

											} else {
												gen2513 = False
											}
											if True == gen2513 {
												gen2519 = True
											} else {
												gen2519 = False
											}

										} else {
											gen2519 = False
										}
										if True == gen2519 {
											gen2525 = True
										} else {
											gen2525 = False
										}

									} else {
										gen2525 = False
									}
									if True == gen2525 {
										gen2530 = True
									} else {
										gen2530 = False
									}

								} else {
									gen2530 = False
								}
								if True == gen2530 {
									gen2534 = True
								} else {
									gen2534 = False
								}

							} else {
								gen2534 = False
							}
							if True == gen2534 {
								gen2538 = True
							} else {
								gen2538 = False
							}

						} else {
							gen2538 = False
						}
						if True == gen2538 {
							gen2541 = True
						} else {
							gen2541 = False
						}

					} else {
						gen2541 = False
					}
					if True == gen2541 {
						gen2543 = True
					} else {
						gen2543 = False
					}

				} else {
					gen2543 = False
				}
				if True == gen2543 {
					gen2477 := Call(__e, ShenFunc(symhd), V434)

					gen2478 := Call(__e, ShenFunc(symtl), gen2477)

					gen2479 := Call(__e, ShenFunc(symhd), gen2478)

					gen2480 := Call(__e, ShenFunc(symtl), gen2479)

					gen2481 := Call(__e, ShenFunc(symhd), gen2480)

					gen2482 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

					gen2483 := Call(__e, ShenFunc(symcons), gen2482, Nil)

					gen2484 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2483)

					gen2485 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen2484)

					gen2488 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

					var gen2489 Obj
					if True == gen2488 {
						gen2487 := Call(__e, ShenFunc(symcons), V435, Nil)

						gen2489 = Call(__e, ShenFunc(symcons), MakeSymbol("shen.sys-error"), gen2487)

					} else {
						gen2486 := Call(__e, ShenFunc(symcons), V435, Nil)

						gen2489 = Call(__e, ShenFunc(symcons), MakeSymbol("shen.f_error"), gen2486)

					}
					gen2490 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

					gen2491 := Call(__e, ShenFunc(symcons), gen2489, gen2490)

					gen2492 := Call(__e, ShenFunc(symcons), gen2485, gen2491)

					gen2493 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2492)

					gen2494 := Call(__e, ShenFunc(symcons), gen2493, Nil)

					gen2495 := Call(__e, ShenFunc(symcons), gen2481, gen2494)

					gen2496 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2495)

					gen2497 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen2496)

					gen2498 := Call(__e, ShenFunc(symcons), gen2497, Nil)

					gen2499 := Call(__e, ShenFunc(symcons), True, gen2498)

					__e.TailApply(ShenFunc(symcons), gen2499, Nil)

					return

				} else {
					gen2475 := Call(__e, ShenFunc(symcons_2), V434)

					var gen2476 Obj
					if True == gen2475 {
						gen2472 := Call(__e, ShenFunc(symhd), V434)

						gen2473 := Call(__e, ShenFunc(symcons_2), gen2472)

						var gen2474 Obj
						if True == gen2473 {
							gen2468 := Call(__e, ShenFunc(symhd), V434)

							gen2469 := Call(__e, ShenFunc(symhd), gen2468)

							gen2470 := Call(__e, ShenFunc(sym_a), True, gen2469)

							var gen2471 Obj
							if True == gen2470 {
								gen2464 := Call(__e, ShenFunc(symhd), V434)

								gen2465 := Call(__e, ShenFunc(symtl), gen2464)

								gen2466 := Call(__e, ShenFunc(symcons_2), gen2465)

								var gen2467 Obj
								if True == gen2466 {
									gen2459 := Call(__e, ShenFunc(symhd), V434)

									gen2460 := Call(__e, ShenFunc(symtl), gen2459)

									gen2461 := Call(__e, ShenFunc(symhd), gen2460)

									gen2462 := Call(__e, ShenFunc(symcons_2), gen2461)

									var gen2463 Obj
									if True == gen2462 {
										gen2453 := Call(__e, ShenFunc(symhd), V434)

										gen2454 := Call(__e, ShenFunc(symtl), gen2453)

										gen2455 := Call(__e, ShenFunc(symhd), gen2454)

										gen2456 := Call(__e, ShenFunc(symhd), gen2455)

										gen2457 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2456)

										var gen2458 Obj
										if True == gen2457 {
											gen2447 := Call(__e, ShenFunc(symhd), V434)

											gen2448 := Call(__e, ShenFunc(symtl), gen2447)

											gen2449 := Call(__e, ShenFunc(symhd), gen2448)

											gen2450 := Call(__e, ShenFunc(symtl), gen2449)

											gen2451 := Call(__e, ShenFunc(symcons_2), gen2450)

											var gen2452 Obj
											if True == gen2451 {
												gen2440 := Call(__e, ShenFunc(symhd), V434)

												gen2441 := Call(__e, ShenFunc(symtl), gen2440)

												gen2442 := Call(__e, ShenFunc(symhd), gen2441)

												gen2443 := Call(__e, ShenFunc(symtl), gen2442)

												gen2444 := Call(__e, ShenFunc(symtl), gen2443)

												gen2445 := Call(__e, ShenFunc(sym_a), Nil, gen2444)

												var gen2446 Obj
												if True == gen2445 {
													gen2436 := Call(__e, ShenFunc(symhd), V434)

													gen2437 := Call(__e, ShenFunc(symtl), gen2436)

													gen2438 := Call(__e, ShenFunc(symtl), gen2437)

													gen2439 := Call(__e, ShenFunc(sym_a), Nil, gen2438)

													if True == gen2439 {
														gen2446 = True
													} else {
														gen2446 = False
													}

												} else {
													gen2446 = False
												}
												if True == gen2446 {
													gen2452 = True
												} else {
													gen2452 = False
												}

											} else {
												gen2452 = False
											}
											if True == gen2452 {
												gen2458 = True
											} else {
												gen2458 = False
											}

										} else {
											gen2458 = False
										}
										if True == gen2458 {
											gen2463 = True
										} else {
											gen2463 = False
										}

									} else {
										gen2463 = False
									}
									if True == gen2463 {
										gen2467 = True
									} else {
										gen2467 = False
									}

								} else {
									gen2467 = False
								}
								if True == gen2467 {
									gen2471 = True
								} else {
									gen2471 = False
								}

							} else {
								gen2471 = False
							}
							if True == gen2471 {
								gen2474 = True
							} else {
								gen2474 = False
							}

						} else {
							gen2474 = False
						}
						if True == gen2474 {
							gen2476 = True
						} else {
							gen2476 = False
						}

					} else {
						gen2476 = False
					}
					if True == gen2476 {
						gen2414 := Call(__e, ShenFunc(symhd), V434)

						gen2415 := Call(__e, ShenFunc(symtl), gen2414)

						gen2416 := Call(__e, ShenFunc(symhd), gen2415)

						gen2417 := Call(__e, ShenFunc(symtl), gen2416)

						gen2418 := Call(__e, ShenFunc(symhd), gen2417)

						gen2419 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

						gen2420 := Call(__e, ShenFunc(symcons), gen2419, Nil)

						gen2421 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2420)

						gen2422 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen2421)

						gen2423 := Call(__e, ShenFunc(symtl), V434)

						gen2424 := Call(__e, ShenFunc(symshen_4encode_1choices), gen2423, V435)

						gen2425 := Call(__e, ShenFunc(symshen_4cond_1form), gen2424)

						gen2426 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

						gen2427 := Call(__e, ShenFunc(symcons), gen2425, gen2426)

						gen2428 := Call(__e, ShenFunc(symcons), gen2422, gen2427)

						gen2429 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2428)

						gen2430 := Call(__e, ShenFunc(symcons), gen2429, Nil)

						gen2431 := Call(__e, ShenFunc(symcons), gen2418, gen2430)

						gen2432 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2431)

						gen2433 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen2432)

						gen2434 := Call(__e, ShenFunc(symcons), gen2433, Nil)

						gen2435 := Call(__e, ShenFunc(symcons), True, gen2434)

						__e.TailApply(ShenFunc(symcons), gen2435, Nil)

						return

					} else {
						gen2412 := Call(__e, ShenFunc(symcons_2), V434)

						var gen2413 Obj
						if True == gen2412 {
							gen2409 := Call(__e, ShenFunc(symhd), V434)

							gen2410 := Call(__e, ShenFunc(symcons_2), gen2409)

							var gen2411 Obj
							if True == gen2410 {
								gen2405 := Call(__e, ShenFunc(symhd), V434)

								gen2406 := Call(__e, ShenFunc(symtl), gen2405)

								gen2407 := Call(__e, ShenFunc(symcons_2), gen2406)

								var gen2408 Obj
								if True == gen2407 {
									gen2400 := Call(__e, ShenFunc(symhd), V434)

									gen2401 := Call(__e, ShenFunc(symtl), gen2400)

									gen2402 := Call(__e, ShenFunc(symhd), gen2401)

									gen2403 := Call(__e, ShenFunc(symcons_2), gen2402)

									var gen2404 Obj
									if True == gen2403 {
										gen2394 := Call(__e, ShenFunc(symhd), V434)

										gen2395 := Call(__e, ShenFunc(symtl), gen2394)

										gen2396 := Call(__e, ShenFunc(symhd), gen2395)

										gen2397 := Call(__e, ShenFunc(symhd), gen2396)

										gen2398 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2397)

										var gen2399 Obj
										if True == gen2398 {
											gen2388 := Call(__e, ShenFunc(symhd), V434)

											gen2389 := Call(__e, ShenFunc(symtl), gen2388)

											gen2390 := Call(__e, ShenFunc(symhd), gen2389)

											gen2391 := Call(__e, ShenFunc(symtl), gen2390)

											gen2392 := Call(__e, ShenFunc(symcons_2), gen2391)

											var gen2393 Obj
											if True == gen2392 {
												gen2381 := Call(__e, ShenFunc(symhd), V434)

												gen2382 := Call(__e, ShenFunc(symtl), gen2381)

												gen2383 := Call(__e, ShenFunc(symhd), gen2382)

												gen2384 := Call(__e, ShenFunc(symtl), gen2383)

												gen2385 := Call(__e, ShenFunc(symtl), gen2384)

												gen2386 := Call(__e, ShenFunc(sym_a), Nil, gen2385)

												var gen2387 Obj
												if True == gen2386 {
													gen2377 := Call(__e, ShenFunc(symhd), V434)

													gen2378 := Call(__e, ShenFunc(symtl), gen2377)

													gen2379 := Call(__e, ShenFunc(symtl), gen2378)

													gen2380 := Call(__e, ShenFunc(sym_a), Nil, gen2379)

													if True == gen2380 {
														gen2387 = True
													} else {
														gen2387 = False
													}

												} else {
													gen2387 = False
												}
												if True == gen2387 {
													gen2393 = True
												} else {
													gen2393 = False
												}

											} else {
												gen2393 = False
											}
											if True == gen2393 {
												gen2399 = True
											} else {
												gen2399 = False
											}

										} else {
											gen2399 = False
										}
										if True == gen2399 {
											gen2404 = True
										} else {
											gen2404 = False
										}

									} else {
										gen2404 = False
									}
									if True == gen2404 {
										gen2408 = True
									} else {
										gen2408 = False
									}

								} else {
									gen2408 = False
								}
								if True == gen2408 {
									gen2411 = True
								} else {
									gen2411 = False
								}

							} else {
								gen2411 = False
							}
							if True == gen2411 {
								gen2413 = True
							} else {
								gen2413 = False
							}

						} else {
							gen2413 = False
						}
						if True == gen2413 {
							gen2339 := Call(__e, ShenFunc(symtl), V434)

							gen2340 := Call(__e, ShenFunc(symshen_4encode_1choices), gen2339, V435)

							gen2341 := Call(__e, ShenFunc(symshen_4cond_1form), gen2340)

							gen2342 := Call(__e, ShenFunc(symcons), gen2341, Nil)

							gen2343 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen2342)

							gen2344 := Call(__e, ShenFunc(symhd), V434)

							gen2345 := Call(__e, ShenFunc(symhd), gen2344)

							gen2346 := Call(__e, ShenFunc(symhd), V434)

							gen2347 := Call(__e, ShenFunc(symtl), gen2346)

							gen2348 := Call(__e, ShenFunc(symhd), gen2347)

							gen2349 := Call(__e, ShenFunc(symtl), gen2348)

							gen2350 := Call(__e, ShenFunc(symhd), gen2349)

							gen2351 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

							gen2352 := Call(__e, ShenFunc(symcons), gen2351, Nil)

							gen2353 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2352)

							gen2354 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen2353)

							gen2355 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), Nil)

							gen2356 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen2355)

							gen2357 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

							gen2358 := Call(__e, ShenFunc(symcons), gen2356, gen2357)

							gen2359 := Call(__e, ShenFunc(symcons), gen2354, gen2358)

							gen2360 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2359)

							gen2361 := Call(__e, ShenFunc(symcons), gen2360, Nil)

							gen2362 := Call(__e, ShenFunc(symcons), gen2350, gen2361)

							gen2363 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2362)

							gen2364 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen2363)

							gen2365 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), Nil)

							gen2366 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen2365)

							gen2367 := Call(__e, ShenFunc(symcons), gen2366, Nil)

							gen2368 := Call(__e, ShenFunc(symcons), gen2364, gen2367)

							gen2369 := Call(__e, ShenFunc(symcons), gen2345, gen2368)

							gen2370 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2369)

							gen2371 := Call(__e, ShenFunc(symcons), gen2370, Nil)

							gen2372 := Call(__e, ShenFunc(symcons), gen2343, gen2371)

							gen2373 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), gen2372)

							gen2374 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen2373)

							gen2375 := Call(__e, ShenFunc(symcons), gen2374, Nil)

							gen2376 := Call(__e, ShenFunc(symcons), True, gen2375)

							__e.TailApply(ShenFunc(symcons), gen2376, Nil)

							return

						} else {
							gen2337 := Call(__e, ShenFunc(symcons_2), V434)

							var gen2338 Obj
							if True == gen2337 {
								gen2334 := Call(__e, ShenFunc(symhd), V434)

								gen2335 := Call(__e, ShenFunc(symcons_2), gen2334)

								var gen2336 Obj
								if True == gen2335 {
									gen2330 := Call(__e, ShenFunc(symhd), V434)

									gen2331 := Call(__e, ShenFunc(symtl), gen2330)

									gen2332 := Call(__e, ShenFunc(symcons_2), gen2331)

									var gen2333 Obj
									if True == gen2332 {
										gen2326 := Call(__e, ShenFunc(symhd), V434)

										gen2327 := Call(__e, ShenFunc(symtl), gen2326)

										gen2328 := Call(__e, ShenFunc(symtl), gen2327)

										gen2329 := Call(__e, ShenFunc(sym_a), Nil, gen2328)

										if True == gen2329 {
											gen2333 = True
										} else {
											gen2333 = False
										}

									} else {
										gen2333 = False
									}
									if True == gen2333 {
										gen2336 = True
									} else {
										gen2336 = False
									}

								} else {
									gen2336 = False
								}
								if True == gen2336 {
									gen2338 = True
								} else {
									gen2338 = False
								}

							} else {
								gen2338 = False
							}
							if True == gen2338 {
								gen2323 := Call(__e, ShenFunc(symhd), V434)

								gen2324 := Call(__e, ShenFunc(symtl), V434)

								gen2325 := Call(__e, ShenFunc(symshen_4encode_1choices), gen2324, V435)

								__e.TailApply(ShenFunc(symcons), gen2323, gen2325)

								return

							} else {
								__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.encode-choices"))

								return
							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.encode-choices"), gen2545)

		gen2700 := MakeNative(func(__e Evaluator) {
			V442 := __e.Get(1)
			_ = V442
			V443 := __e.Get(2)
			_ = V443
			gen2699 := Call(__e, ShenFunc(sym_a), Nil, V442)

			if True == gen2699 {
				__e.TailApply(ShenFunc(symcons), V443, Nil)

				return
			} else {
				gen2697 := Call(__e, ShenFunc(symcons_2), V442)

				var gen2698 Obj
				if True == gen2697 {
					gen2694 := Call(__e, ShenFunc(symhd), V442)

					gen2695 := Call(__e, ShenFunc(symcons_2), gen2694)

					var gen2696 Obj
					if True == gen2695 {
						gen2690 := Call(__e, ShenFunc(symhd), V442)

						gen2691 := Call(__e, ShenFunc(symhd), gen2690)

						gen2692 := Call(__e, ShenFunc(symcons_2), gen2691)

						var gen2693 Obj
						if True == gen2692 {
							gen2685 := Call(__e, ShenFunc(symhd), V442)

							gen2686 := Call(__e, ShenFunc(symhd), gen2685)

							gen2687 := Call(__e, ShenFunc(symhd), gen2686)

							gen2688 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2687)

							var gen2689 Obj
							if True == gen2688 {
								gen2680 := Call(__e, ShenFunc(symhd), V442)

								gen2681 := Call(__e, ShenFunc(symhd), gen2680)

								gen2682 := Call(__e, ShenFunc(symtl), gen2681)

								gen2683 := Call(__e, ShenFunc(symcons_2), gen2682)

								var gen2684 Obj
								if True == gen2683 {
									gen2674 := Call(__e, ShenFunc(symhd), V442)

									gen2675 := Call(__e, ShenFunc(symhd), gen2674)

									gen2676 := Call(__e, ShenFunc(symtl), gen2675)

									gen2677 := Call(__e, ShenFunc(symhd), gen2676)

									gen2678 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2677)

									var gen2679 Obj
									if True == gen2678 {
										gen2668 := Call(__e, ShenFunc(symhd), V442)

										gen2669 := Call(__e, ShenFunc(symhd), gen2668)

										gen2670 := Call(__e, ShenFunc(symtl), gen2669)

										gen2671 := Call(__e, ShenFunc(symtl), gen2670)

										gen2672 := Call(__e, ShenFunc(sym_a), Nil, gen2671)

										var gen2673 Obj
										if True == gen2672 {
											gen2664 := Call(__e, ShenFunc(symhd), V442)

											gen2665 := Call(__e, ShenFunc(symtl), gen2664)

											gen2666 := Call(__e, ShenFunc(symcons_2), gen2665)

											var gen2667 Obj
											if True == gen2666 {
												gen2659 := Call(__e, ShenFunc(symhd), V442)

												gen2660 := Call(__e, ShenFunc(symtl), gen2659)

												gen2661 := Call(__e, ShenFunc(symhd), gen2660)

												gen2662 := Call(__e, ShenFunc(symcons_2), gen2661)

												var gen2663 Obj
												if True == gen2662 {
													gen2653 := Call(__e, ShenFunc(symhd), V442)

													gen2654 := Call(__e, ShenFunc(symtl), gen2653)

													gen2655 := Call(__e, ShenFunc(symhd), gen2654)

													gen2656 := Call(__e, ShenFunc(symhd), gen2655)

													gen2657 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2656)

													var gen2658 Obj
													if True == gen2657 {
														gen2647 := Call(__e, ShenFunc(symhd), V442)

														gen2648 := Call(__e, ShenFunc(symtl), gen2647)

														gen2649 := Call(__e, ShenFunc(symhd), gen2648)

														gen2650 := Call(__e, ShenFunc(symtl), gen2649)

														gen2651 := Call(__e, ShenFunc(symcons_2), gen2650)

														var gen2652 Obj
														if True == gen2651 {
															gen2640 := Call(__e, ShenFunc(symhd), V442)

															gen2641 := Call(__e, ShenFunc(symtl), gen2640)

															gen2642 := Call(__e, ShenFunc(symhd), gen2641)

															gen2643 := Call(__e, ShenFunc(symtl), gen2642)

															gen2644 := Call(__e, ShenFunc(symtl), gen2643)

															gen2645 := Call(__e, ShenFunc(sym_a), Nil, gen2644)

															var gen2646 Obj
															if True == gen2645 {
																gen2636 := Call(__e, ShenFunc(symhd), V442)

																gen2637 := Call(__e, ShenFunc(symtl), gen2636)

																gen2638 := Call(__e, ShenFunc(symtl), gen2637)

																gen2639 := Call(__e, ShenFunc(sym_a), Nil, gen2638)

																if True == gen2639 {
																	gen2646 = True
																} else {
																	gen2646 = False
																}

															} else {
																gen2646 = False
															}
															if True == gen2646 {
																gen2652 = True
															} else {
																gen2652 = False
															}

														} else {
															gen2652 = False
														}
														if True == gen2652 {
															gen2658 = True
														} else {
															gen2658 = False
														}

													} else {
														gen2658 = False
													}
													if True == gen2658 {
														gen2663 = True
													} else {
														gen2663 = False
													}

												} else {
													gen2663 = False
												}
												if True == gen2663 {
													gen2667 = True
												} else {
													gen2667 = False
												}

											} else {
												gen2667 = False
											}
											if True == gen2667 {
												gen2673 = True
											} else {
												gen2673 = False
											}

										} else {
											gen2673 = False
										}
										if True == gen2673 {
											gen2679 = True
										} else {
											gen2679 = False
										}

									} else {
										gen2679 = False
									}
									if True == gen2679 {
										gen2684 = True
									} else {
										gen2684 = False
									}

								} else {
									gen2684 = False
								}
								if True == gen2684 {
									gen2689 = True
								} else {
									gen2689 = False
								}

							} else {
								gen2689 = False
							}
							if True == gen2689 {
								gen2693 = True
							} else {
								gen2693 = False
							}

						} else {
							gen2693 = False
						}
						if True == gen2693 {
							gen2696 = True
						} else {
							gen2696 = False
						}

					} else {
						gen2696 = False
					}
					if True == gen2696 {
						gen2698 = True
					} else {
						gen2698 = False
					}

				} else {
					gen2698 = False
				}
				if True == gen2698 {
					gen2631 := Call(__e, ShenFunc(symhd), V442)

					gen2632 := Call(__e, ShenFunc(symtl), gen2631)

					gen2633 := Call(__e, ShenFunc(symcons), True, gen2632)

					gen2634 := Call(__e, ShenFunc(symtl), V442)

					gen2635 := Call(__e, ShenFunc(symshen_4case_1form), gen2634, V443)

					__e.TailApply(ShenFunc(symcons), gen2633, gen2635)

					return

				} else {
					gen2629 := Call(__e, ShenFunc(symcons_2), V442)

					var gen2630 Obj
					if True == gen2629 {
						gen2626 := Call(__e, ShenFunc(symhd), V442)

						gen2627 := Call(__e, ShenFunc(symcons_2), gen2626)

						var gen2628 Obj
						if True == gen2627 {
							gen2622 := Call(__e, ShenFunc(symhd), V442)

							gen2623 := Call(__e, ShenFunc(symhd), gen2622)

							gen2624 := Call(__e, ShenFunc(symcons_2), gen2623)

							var gen2625 Obj
							if True == gen2624 {
								gen2617 := Call(__e, ShenFunc(symhd), V442)

								gen2618 := Call(__e, ShenFunc(symhd), gen2617)

								gen2619 := Call(__e, ShenFunc(symhd), gen2618)

								gen2620 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2619)

								var gen2621 Obj
								if True == gen2620 {
									gen2612 := Call(__e, ShenFunc(symhd), V442)

									gen2613 := Call(__e, ShenFunc(symhd), gen2612)

									gen2614 := Call(__e, ShenFunc(symtl), gen2613)

									gen2615 := Call(__e, ShenFunc(symcons_2), gen2614)

									var gen2616 Obj
									if True == gen2615 {
										gen2606 := Call(__e, ShenFunc(symhd), V442)

										gen2607 := Call(__e, ShenFunc(symhd), gen2606)

										gen2608 := Call(__e, ShenFunc(symtl), gen2607)

										gen2609 := Call(__e, ShenFunc(symhd), gen2608)

										gen2610 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2609)

										var gen2611 Obj
										if True == gen2610 {
											gen2600 := Call(__e, ShenFunc(symhd), V442)

											gen2601 := Call(__e, ShenFunc(symhd), gen2600)

											gen2602 := Call(__e, ShenFunc(symtl), gen2601)

											gen2603 := Call(__e, ShenFunc(symtl), gen2602)

											gen2604 := Call(__e, ShenFunc(sym_a), Nil, gen2603)

											var gen2605 Obj
											if True == gen2604 {
												gen2596 := Call(__e, ShenFunc(symhd), V442)

												gen2597 := Call(__e, ShenFunc(symtl), gen2596)

												gen2598 := Call(__e, ShenFunc(symcons_2), gen2597)

												var gen2599 Obj
												if True == gen2598 {
													gen2592 := Call(__e, ShenFunc(symhd), V442)

													gen2593 := Call(__e, ShenFunc(symtl), gen2592)

													gen2594 := Call(__e, ShenFunc(symtl), gen2593)

													gen2595 := Call(__e, ShenFunc(sym_a), Nil, gen2594)

													if True == gen2595 {
														gen2599 = True
													} else {
														gen2599 = False
													}

												} else {
													gen2599 = False
												}
												if True == gen2599 {
													gen2605 = True
												} else {
													gen2605 = False
												}

											} else {
												gen2605 = False
											}
											if True == gen2605 {
												gen2611 = True
											} else {
												gen2611 = False
											}

										} else {
											gen2611 = False
										}
										if True == gen2611 {
											gen2616 = True
										} else {
											gen2616 = False
										}

									} else {
										gen2616 = False
									}
									if True == gen2616 {
										gen2621 = True
									} else {
										gen2621 = False
									}

								} else {
									gen2621 = False
								}
								if True == gen2621 {
									gen2625 = True
								} else {
									gen2625 = False
								}

							} else {
								gen2625 = False
							}
							if True == gen2625 {
								gen2628 = True
							} else {
								gen2628 = False
							}

						} else {
							gen2628 = False
						}
						if True == gen2628 {
							gen2630 = True
						} else {
							gen2630 = False
						}

					} else {
						gen2630 = False
					}
					if True == gen2630 {
						gen2589 := Call(__e, ShenFunc(symhd), V442)

						gen2590 := Call(__e, ShenFunc(symtl), gen2589)

						gen2591 := Call(__e, ShenFunc(symcons), True, gen2590)

						__e.TailApply(ShenFunc(symcons), gen2591, Nil)

						return

					} else {
						gen2587 := Call(__e, ShenFunc(symcons_2), V442)

						var gen2588 Obj
						if True == gen2587 {
							gen2584 := Call(__e, ShenFunc(symhd), V442)

							gen2585 := Call(__e, ShenFunc(symcons_2), gen2584)

							var gen2586 Obj
							if True == gen2585 {
								gen2580 := Call(__e, ShenFunc(symhd), V442)

								gen2581 := Call(__e, ShenFunc(symhd), gen2580)

								gen2582 := Call(__e, ShenFunc(symcons_2), gen2581)

								var gen2583 Obj
								if True == gen2582 {
									gen2575 := Call(__e, ShenFunc(symhd), V442)

									gen2576 := Call(__e, ShenFunc(symhd), gen2575)

									gen2577 := Call(__e, ShenFunc(symhd), gen2576)

									gen2578 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2577)

									var gen2579 Obj
									if True == gen2578 {
										gen2570 := Call(__e, ShenFunc(symhd), V442)

										gen2571 := Call(__e, ShenFunc(symhd), gen2570)

										gen2572 := Call(__e, ShenFunc(symtl), gen2571)

										gen2573 := Call(__e, ShenFunc(symcons_2), gen2572)

										var gen2574 Obj
										if True == gen2573 {
											gen2564 := Call(__e, ShenFunc(symhd), V442)

											gen2565 := Call(__e, ShenFunc(symhd), gen2564)

											gen2566 := Call(__e, ShenFunc(symtl), gen2565)

											gen2567 := Call(__e, ShenFunc(symhd), gen2566)

											gen2568 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2567)

											var gen2569 Obj
											if True == gen2568 {
												gen2560 := Call(__e, ShenFunc(symhd), V442)

												gen2561 := Call(__e, ShenFunc(symtl), gen2560)

												gen2562 := Call(__e, ShenFunc(symcons_2), gen2561)

												var gen2563 Obj
												if True == gen2562 {
													gen2556 := Call(__e, ShenFunc(symhd), V442)

													gen2557 := Call(__e, ShenFunc(symtl), gen2556)

													gen2558 := Call(__e, ShenFunc(symtl), gen2557)

													gen2559 := Call(__e, ShenFunc(sym_a), Nil, gen2558)

													if True == gen2559 {
														gen2563 = True
													} else {
														gen2563 = False
													}

												} else {
													gen2563 = False
												}
												if True == gen2563 {
													gen2569 = True
												} else {
													gen2569 = False
												}

											} else {
												gen2569 = False
											}
											if True == gen2569 {
												gen2574 = True
											} else {
												gen2574 = False
											}

										} else {
											gen2574 = False
										}
										if True == gen2574 {
											gen2579 = True
										} else {
											gen2579 = False
										}

									} else {
										gen2579 = False
									}
									if True == gen2579 {
										gen2583 = True
									} else {
										gen2583 = False
									}

								} else {
									gen2583 = False
								}
								if True == gen2583 {
									gen2586 = True
								} else {
									gen2586 = False
								}

							} else {
								gen2586 = False
							}
							if True == gen2586 {
								gen2588 = True
							} else {
								gen2588 = False
							}

						} else {
							gen2588 = False
						}
						if True == gen2588 {
							gen2546 := Call(__e, ShenFunc(symhd), V442)

							gen2547 := Call(__e, ShenFunc(symhd), gen2546)

							gen2548 := Call(__e, ShenFunc(symtl), gen2547)

							gen2549 := Call(__e, ShenFunc(symtl), gen2548)

							gen2550 := Call(__e, ShenFunc(symshen_4embed_1and), gen2549)

							gen2551 := Call(__e, ShenFunc(symhd), V442)

							gen2552 := Call(__e, ShenFunc(symtl), gen2551)

							gen2553 := Call(__e, ShenFunc(symcons), gen2550, gen2552)

							gen2554 := Call(__e, ShenFunc(symtl), V442)

							gen2555 := Call(__e, ShenFunc(symshen_4case_1form), gen2554, V443)

							__e.TailApply(ShenFunc(symcons), gen2553, gen2555)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.case-form"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.case-form"), gen2700)

		gen2711 := MakeNative(func(__e Evaluator) {
			V445 := __e.Get(1)
			_ = V445
			gen2709 := Call(__e, ShenFunc(symcons_2), V445)

			var gen2710 Obj
			if True == gen2709 {
				gen2707 := Call(__e, ShenFunc(symtl), V445)

				gen2708 := Call(__e, ShenFunc(sym_a), Nil, gen2707)

				if True == gen2708 {
					gen2710 = True
				} else {
					gen2710 = False
				}

			} else {
				gen2710 = False
			}
			if True == gen2710 {
				__e.TailApply(ShenFunc(symhd), V445)

				return
			} else {
				gen2706 := Call(__e, ShenFunc(symcons_2), V445)

				if True == gen2706 {
					gen2701 := Call(__e, ShenFunc(symhd), V445)

					gen2702 := Call(__e, ShenFunc(symtl), V445)

					gen2703 := Call(__e, ShenFunc(symshen_4embed_1and), gen2702)

					gen2704 := Call(__e, ShenFunc(symcons), gen2703, Nil)

					gen2705 := Call(__e, ShenFunc(symcons), gen2701, gen2704)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("and"), gen2705)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.embed-and"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.embed-and"), gen2711)

		gen2715 := MakeNative(func(__e Evaluator) {
			V447 := __e.Get(1)
			_ = V447
			gen2712 := Call(__e, ShenFunc(symcons), V447, Nil)

			gen2713 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.f_error"), gen2712)

			gen2714 := Call(__e, ShenFunc(symcons), gen2713, Nil)

			__e.TailApply(ShenFunc(symcons), True, gen2714)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.err-condition"), gen2715)

		gen2718 := MakeNative(func(__e Evaluator) {
			V449 := __e.Get(1)
			_ = V449
			gen2716 := Call(__e, ShenFunc(symshen_4app), V449, MakeString(": unexpected argument\n"), MakeSymbol("shen.a"))

			gen2717 := Call(__e, ShenFunc(symcn), MakeString("system function "), gen2716)

			__e.TailApply(ShenFunc(symsimple_1error), gen2717)

			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.sys-error"), gen2718)

		return

	}, 0))
}
