package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen3470 := MakeNative(func(__e Evaluator) {
			V1722 := __e.Get(1)
			_ = V1722
			gen3468 := Call(__e, __e.Global(symcons_2), V1722)

			var gen3469 Obj
			if True == gen3468 {
				gen3465 := Call(__e, __e.Global(symtl), V1722)

				gen3466 := Call(__e, __e.Global(symcons_2), gen3465)

				var gen3467 Obj
				if True == gen3466 {
					gen3462 := Call(__e, __e.Global(symtl), V1722)

					gen3463 := Call(__e, __e.Global(symtl), gen3462)

					gen3464 := Call(__e, __e.Global(sym_a), Nil, gen3463)

					if True == gen3464 {
						gen3467 = True
					} else {
						gen3467 = False
					}

				} else {
					gen3467 = False
				}
				if True == gen3467 {
					gen3469 = True
				} else {
					gen3469 = False
				}

			} else {
				gen3469 = False
			}
			if True == gen3469 {
				gen3458 := Call(__e, __e.Global(symhd), V1722)

				gen3459 := Call(__e, __e.Global(symshen_4next_150), MakeNumber(50), gen3458)

				gen3460 := Call(__e, __e.Global(symshen_4app), gen3459, MakeString("\n"), MakeSymbol("shen.a"))

				gen3461 := Call(__e, __e.Global(symcn), MakeString("datatype syntax error here:\n\n "), gen3460)

				__e.TailApply(__e.Global(symsimple_1error), gen3461)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.datatype-error"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.datatype-error"), gen3470)

		gen3491 := MakeNative(func(__e Evaluator) {
			V1724 := __e.Get(1)
			_ = V1724
			gen3471 := Call(__e, __e.Global(symshen_4_5datatype_1rule_6), V1724)

			Parse__shen_4_5datatype_1rule_6 := gen3471
			gen3480 := Call(__e, __e.Global(symfail))

			gen3481 := Call(__e, __e.Global(sym_a), gen3480, Parse__shen_4_5datatype_1rule_6)

			gen3482 := Call(__e, __e.Global(symnot), gen3481)

			var gen3483 Obj
			if True == gen3482 {
				gen3472 := Call(__e, __e.Global(symshen_4_5datatype_1rules_6), Parse__shen_4_5datatype_1rule_6)

				Parse__shen_4_5datatype_1rules_6 := gen3472
				gen3477 := Call(__e, __e.Global(symfail))

				gen3478 := Call(__e, __e.Global(sym_a), gen3477, Parse__shen_4_5datatype_1rules_6)

				gen3479 := Call(__e, __e.Global(symnot), gen3478)

				if True == gen3479 {
					gen3473 := Call(__e, __e.Global(symhd), Parse__shen_4_5datatype_1rules_6)

					gen3474 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5datatype_1rule_6)

					gen3475 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5datatype_1rules_6)

					gen3476 := Call(__e, __e.Global(symcons), gen3474, gen3475)

					gen3483 = Call(__e, __e.Global(symshen_4pair), gen3473, gen3476)

				} else {
					gen3483 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3483 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3483
			gen3489 := Call(__e, __e.Global(symfail))

			gen3490 := Call(__e, __e.Global(sym_a), YaccParse, gen3489)

			if True == gen3490 {
				gen3484 := Call(__e, __e.Global(sym_5e_6), V1724)

				Parse___5e_6 := gen3484
				gen3486 := Call(__e, __e.Global(symfail))

				gen3487 := Call(__e, __e.Global(sym_a), gen3486, Parse___5e_6)

				gen3488 := Call(__e, __e.Global(symnot), gen3487)

				if True == gen3488 {
					gen3485 := Call(__e, __e.Global(symhd), Parse___5e_6)

					__e.TailApply(__e.Global(symshen_4pair), gen3485, Nil)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<datatype-rules>"), gen3491)

		gen3543 := MakeNative(func(__e Evaluator) {
			V1726 := __e.Get(1)
			_ = V1726
			gen3492 := Call(__e, __e.Global(symshen_4_5side_1conditions_6), V1726)

			Parse__shen_4_5side_1conditions_6 := gen3492
			gen3513 := Call(__e, __e.Global(symfail))

			gen3514 := Call(__e, __e.Global(sym_a), gen3513, Parse__shen_4_5side_1conditions_6)

			gen3515 := Call(__e, __e.Global(symnot), gen3514)

			var gen3516 Obj
			if True == gen3515 {
				gen3493 := Call(__e, __e.Global(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

				Parse__shen_4_5premises_6 := gen3493
				gen3510 := Call(__e, __e.Global(symfail))

				gen3511 := Call(__e, __e.Global(sym_a), gen3510, Parse__shen_4_5premises_6)

				gen3512 := Call(__e, __e.Global(symnot), gen3511)

				if True == gen3512 {
					gen3494 := Call(__e, __e.Global(symshen_4_5singleunderline_6), Parse__shen_4_5premises_6)

					Parse__shen_4_5singleunderline_6 := gen3494
					gen3507 := Call(__e, __e.Global(symfail))

					gen3508 := Call(__e, __e.Global(sym_a), gen3507, Parse__shen_4_5singleunderline_6)

					gen3509 := Call(__e, __e.Global(symnot), gen3508)

					if True == gen3509 {
						gen3495 := Call(__e, __e.Global(symshen_4_5conclusion_6), Parse__shen_4_5singleunderline_6)

						Parse__shen_4_5conclusion_6 := gen3495
						gen3504 := Call(__e, __e.Global(symfail))

						gen3505 := Call(__e, __e.Global(sym_a), gen3504, Parse__shen_4_5conclusion_6)

						gen3506 := Call(__e, __e.Global(symnot), gen3505)

						if True == gen3506 {
							gen3496 := Call(__e, __e.Global(symhd), Parse__shen_4_5conclusion_6)

							gen3497 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

							gen3498 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5premises_6)

							gen3499 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5conclusion_6)

							gen3500 := Call(__e, __e.Global(symcons), gen3499, Nil)

							gen3501 := Call(__e, __e.Global(symcons), gen3498, gen3500)

							gen3502 := Call(__e, __e.Global(symcons), gen3497, gen3501)

							gen3503 := Call(__e, __e.Global(symshen_4sequent), MakeSymbol("shen.single"), gen3502)

							gen3516 = Call(__e, __e.Global(symshen_4pair), gen3496, gen3503)

						} else {
							gen3516 = Call(__e, __e.Global(symfail))

						}

					} else {
						gen3516 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3516 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3516 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3516
			gen3541 := Call(__e, __e.Global(symfail))

			gen3542 := Call(__e, __e.Global(sym_a), YaccParse, gen3541)

			if True == gen3542 {
				gen3517 := Call(__e, __e.Global(symshen_4_5side_1conditions_6), V1726)

				Parse__shen_4_5side_1conditions_6 := gen3517
				gen3538 := Call(__e, __e.Global(symfail))

				gen3539 := Call(__e, __e.Global(sym_a), gen3538, Parse__shen_4_5side_1conditions_6)

				gen3540 := Call(__e, __e.Global(symnot), gen3539)

				if True == gen3540 {
					gen3518 := Call(__e, __e.Global(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

					Parse__shen_4_5premises_6 := gen3518
					gen3535 := Call(__e, __e.Global(symfail))

					gen3536 := Call(__e, __e.Global(sym_a), gen3535, Parse__shen_4_5premises_6)

					gen3537 := Call(__e, __e.Global(symnot), gen3536)

					if True == gen3537 {
						gen3519 := Call(__e, __e.Global(symshen_4_5doubleunderline_6), Parse__shen_4_5premises_6)

						Parse__shen_4_5doubleunderline_6 := gen3519
						gen3532 := Call(__e, __e.Global(symfail))

						gen3533 := Call(__e, __e.Global(sym_a), gen3532, Parse__shen_4_5doubleunderline_6)

						gen3534 := Call(__e, __e.Global(symnot), gen3533)

						if True == gen3534 {
							gen3520 := Call(__e, __e.Global(symshen_4_5conclusion_6), Parse__shen_4_5doubleunderline_6)

							Parse__shen_4_5conclusion_6 := gen3520
							gen3529 := Call(__e, __e.Global(symfail))

							gen3530 := Call(__e, __e.Global(sym_a), gen3529, Parse__shen_4_5conclusion_6)

							gen3531 := Call(__e, __e.Global(symnot), gen3530)

							if True == gen3531 {
								gen3521 := Call(__e, __e.Global(symhd), Parse__shen_4_5conclusion_6)

								gen3522 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

								gen3523 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5premises_6)

								gen3524 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5conclusion_6)

								gen3525 := Call(__e, __e.Global(symcons), gen3524, Nil)

								gen3526 := Call(__e, __e.Global(symcons), gen3523, gen3525)

								gen3527 := Call(__e, __e.Global(symcons), gen3522, gen3526)

								gen3528 := Call(__e, __e.Global(symshen_4sequent), MakeSymbol("shen.double"), gen3527)

								__e.TailApply(__e.Global(symshen_4pair), gen3521, gen3528)

								return

							} else {
								__e.TailApply(__e.Global(symfail))

								return
							}

						} else {
							__e.TailApply(__e.Global(symfail))

							return
						}

					} else {
						__e.TailApply(__e.Global(symfail))

						return
					}

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<datatype-rule>"), gen3543)

		gen3564 := MakeNative(func(__e Evaluator) {
			V1728 := __e.Get(1)
			_ = V1728
			gen3544 := Call(__e, __e.Global(symshen_4_5side_1condition_6), V1728)

			Parse__shen_4_5side_1condition_6 := gen3544
			gen3553 := Call(__e, __e.Global(symfail))

			gen3554 := Call(__e, __e.Global(sym_a), gen3553, Parse__shen_4_5side_1condition_6)

			gen3555 := Call(__e, __e.Global(symnot), gen3554)

			var gen3556 Obj
			if True == gen3555 {
				gen3545 := Call(__e, __e.Global(symshen_4_5side_1conditions_6), Parse__shen_4_5side_1condition_6)

				Parse__shen_4_5side_1conditions_6 := gen3545
				gen3550 := Call(__e, __e.Global(symfail))

				gen3551 := Call(__e, __e.Global(sym_a), gen3550, Parse__shen_4_5side_1conditions_6)

				gen3552 := Call(__e, __e.Global(symnot), gen3551)

				if True == gen3552 {
					gen3546 := Call(__e, __e.Global(symhd), Parse__shen_4_5side_1conditions_6)

					gen3547 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5side_1condition_6)

					gen3548 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

					gen3549 := Call(__e, __e.Global(symcons), gen3547, gen3548)

					gen3556 = Call(__e, __e.Global(symshen_4pair), gen3546, gen3549)

				} else {
					gen3556 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3556 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3556
			gen3562 := Call(__e, __e.Global(symfail))

			gen3563 := Call(__e, __e.Global(sym_a), YaccParse, gen3562)

			if True == gen3563 {
				gen3557 := Call(__e, __e.Global(sym_5e_6), V1728)

				Parse___5e_6 := gen3557
				gen3559 := Call(__e, __e.Global(symfail))

				gen3560 := Call(__e, __e.Global(sym_a), gen3559, Parse___5e_6)

				gen3561 := Call(__e, __e.Global(symnot), gen3560)

				if True == gen3561 {
					gen3558 := Call(__e, __e.Global(symhd), Parse___5e_6)

					__e.TailApply(__e.Global(symshen_4pair), gen3558, Nil)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<side-conditions>"), gen3564)

		gen3606 := MakeNative(func(__e Evaluator) {
			V1732 := __e.Get(1)
			_ = V1732
			gen3578 := Call(__e, __e.Global(symhd), V1732)

			gen3579 := Call(__e, __e.Global(symcons_2), gen3578)

			var gen3580 Obj
			if True == gen3579 {
				gen3576 := Call(__e, __e.Global(symshen_4hdhd), V1732)

				gen3577 := Call(__e, __e.Global(sym_a), MakeSymbol("if"), gen3576)

				if True == gen3577 {
					gen3580 = True
				} else {
					gen3580 = False
				}

			} else {
				gen3580 = False
			}
			var gen3581 Obj
			if True == gen3580 {
				gen3565 := Call(__e, __e.Global(symshen_4tlhd), V1732)

				gen3566 := Call(__e, __e.Global(symshen_4hdtl), V1732)

				gen3567 := Call(__e, __e.Global(symshen_4pair), gen3565, gen3566)

				NewStream1729 := gen3567
				gen3568 := Call(__e, __e.Global(symshen_4_5expr_6), NewStream1729)

				Parse__shen_4_5expr_6 := gen3568
				gen3573 := Call(__e, __e.Global(symfail))

				gen3574 := Call(__e, __e.Global(sym_a), gen3573, Parse__shen_4_5expr_6)

				gen3575 := Call(__e, __e.Global(symnot), gen3574)

				if True == gen3575 {
					gen3569 := Call(__e, __e.Global(symhd), Parse__shen_4_5expr_6)

					gen3570 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

					gen3571 := Call(__e, __e.Global(symcons), gen3570, Nil)

					gen3572 := Call(__e, __e.Global(symcons), MakeSymbol("if"), gen3571)

					gen3581 = Call(__e, __e.Global(symshen_4pair), gen3569, gen3572)

				} else {
					gen3581 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3581 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3581
			gen3604 := Call(__e, __e.Global(symfail))

			gen3605 := Call(__e, __e.Global(sym_a), YaccParse, gen3604)

			if True == gen3605 {
				gen3601 := Call(__e, __e.Global(symhd), V1732)

				gen3602 := Call(__e, __e.Global(symcons_2), gen3601)

				var gen3603 Obj
				if True == gen3602 {
					gen3599 := Call(__e, __e.Global(symshen_4hdhd), V1732)

					gen3600 := Call(__e, __e.Global(sym_a), MakeSymbol("let"), gen3599)

					if True == gen3600 {
						gen3603 = True
					} else {
						gen3603 = False
					}

				} else {
					gen3603 = False
				}
				if True == gen3603 {
					gen3582 := Call(__e, __e.Global(symshen_4tlhd), V1732)

					gen3583 := Call(__e, __e.Global(symshen_4hdtl), V1732)

					gen3584 := Call(__e, __e.Global(symshen_4pair), gen3582, gen3583)

					NewStream1730 := gen3584
					gen3585 := Call(__e, __e.Global(symshen_4_5variable_2_6), NewStream1730)

					Parse__shen_4_5variable_2_6 := gen3585
					gen3596 := Call(__e, __e.Global(symfail))

					gen3597 := Call(__e, __e.Global(sym_a), gen3596, Parse__shen_4_5variable_2_6)

					gen3598 := Call(__e, __e.Global(symnot), gen3597)

					if True == gen3598 {
						gen3586 := Call(__e, __e.Global(symshen_4_5expr_6), Parse__shen_4_5variable_2_6)

						Parse__shen_4_5expr_6 := gen3586
						gen3593 := Call(__e, __e.Global(symfail))

						gen3594 := Call(__e, __e.Global(sym_a), gen3593, Parse__shen_4_5expr_6)

						gen3595 := Call(__e, __e.Global(symnot), gen3594)

						if True == gen3595 {
							gen3587 := Call(__e, __e.Global(symhd), Parse__shen_4_5expr_6)

							gen3588 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5variable_2_6)

							gen3589 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

							gen3590 := Call(__e, __e.Global(symcons), gen3589, Nil)

							gen3591 := Call(__e, __e.Global(symcons), gen3588, gen3590)

							gen3592 := Call(__e, __e.Global(symcons), MakeSymbol("let"), gen3591)

							__e.TailApply(__e.Global(symshen_4pair), gen3587, gen3592)

							return

						} else {
							__e.TailApply(__e.Global(symfail))

							return
						}

					} else {
						__e.TailApply(__e.Global(symfail))

						return
					}

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<side-condition>"), gen3606)

		gen3615 := MakeNative(func(__e Evaluator) {
			V1734 := __e.Get(1)
			_ = V1734
			gen3613 := Call(__e, __e.Global(symhd), V1734)

			gen3614 := Call(__e, __e.Global(symcons_2), gen3613)

			if True == gen3614 {
				gen3607 := Call(__e, __e.Global(symshen_4hdhd), V1734)

				Parse__X := gen3607
				gen3612 := Call(__e, __e.Global(symvariable_2), Parse__X)

				if True == gen3612 {
					gen3608 := Call(__e, __e.Global(symshen_4tlhd), V1734)

					gen3609 := Call(__e, __e.Global(symshen_4hdtl), V1734)

					gen3610 := Call(__e, __e.Global(symshen_4pair), gen3608, gen3609)

					gen3611 := Call(__e, __e.Global(symhd), gen3610)

					__e.TailApply(__e.Global(symshen_4pair), gen3611, Parse__X)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<variable?>"), gen3615)

		gen3632 := MakeNative(func(__e Evaluator) {
			V1736 := __e.Get(1)
			_ = V1736
			gen3630 := Call(__e, __e.Global(symhd), V1736)

			gen3631 := Call(__e, __e.Global(symcons_2), gen3630)

			if True == gen3631 {
				gen3616 := Call(__e, __e.Global(symshen_4hdhd), V1736)

				Parse__X := gen3616
				gen3625 := Call(__e, __e.Global(symcons), MakeSymbol(";"), Nil)

				gen3626 := Call(__e, __e.Global(symcons), MakeSymbol(">>"), gen3625)

				gen3627 := Call(__e, __e.Global(symelement_2), Parse__X, gen3626)

				var gen3628 Obj
				if True == gen3627 {
					gen3628 = True
				} else {
					gen3623 := Call(__e, __e.Global(symshen_4singleunderline_2), Parse__X)

					var gen3624 Obj
					if True == gen3623 {
						gen3624 = True
					} else {
						gen3622 := Call(__e, __e.Global(symshen_4doubleunderline_2), Parse__X)

						if True == gen3622 {
							gen3624 = True
						} else {
							gen3624 = False
						}

					}
					if True == gen3624 {
						gen3628 = True
					} else {
						gen3628 = False
					}

				}
				gen3629 := Call(__e, __e.Global(symnot), gen3628)

				if True == gen3629 {
					gen3617 := Call(__e, __e.Global(symshen_4tlhd), V1736)

					gen3618 := Call(__e, __e.Global(symshen_4hdtl), V1736)

					gen3619 := Call(__e, __e.Global(symshen_4pair), gen3617, gen3618)

					gen3620 := Call(__e, __e.Global(symhd), gen3619)

					gen3621 := Call(__e, __e.Global(symshen_4remove_1bar), Parse__X)

					__e.TailApply(__e.Global(symshen_4pair), gen3620, gen3621)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<expr>"), gen3632)

		gen3659 := MakeNative(func(__e Evaluator) {
			V1738 := __e.Get(1)
			_ = V1738
			gen3657 := Call(__e, __e.Global(symcons_2), V1738)

			var gen3658 Obj
			if True == gen3657 {
				gen3654 := Call(__e, __e.Global(symtl), V1738)

				gen3655 := Call(__e, __e.Global(symcons_2), gen3654)

				var gen3656 Obj
				if True == gen3655 {
					gen3650 := Call(__e, __e.Global(symtl), V1738)

					gen3651 := Call(__e, __e.Global(symtl), gen3650)

					gen3652 := Call(__e, __e.Global(symcons_2), gen3651)

					var gen3653 Obj
					if True == gen3652 {
						gen3645 := Call(__e, __e.Global(symtl), V1738)

						gen3646 := Call(__e, __e.Global(symtl), gen3645)

						gen3647 := Call(__e, __e.Global(symtl), gen3646)

						gen3648 := Call(__e, __e.Global(sym_a), Nil, gen3647)

						var gen3649 Obj
						if True == gen3648 {
							gen3642 := Call(__e, __e.Global(symtl), V1738)

							gen3643 := Call(__e, __e.Global(symhd), gen3642)

							gen3644 := Call(__e, __e.Global(sym_a), gen3643, MakeSymbol("bar!"))

							if True == gen3644 {
								gen3649 = True
							} else {
								gen3649 = False
							}

						} else {
							gen3649 = False
						}
						if True == gen3649 {
							gen3653 = True
						} else {
							gen3653 = False
						}

					} else {
						gen3653 = False
					}
					if True == gen3653 {
						gen3656 = True
					} else {
						gen3656 = False
					}

				} else {
					gen3656 = False
				}
				if True == gen3656 {
					gen3658 = True
				} else {
					gen3658 = False
				}

			} else {
				gen3658 = False
			}
			if True == gen3658 {
				gen3638 := Call(__e, __e.Global(symhd), V1738)

				gen3639 := Call(__e, __e.Global(symtl), V1738)

				gen3640 := Call(__e, __e.Global(symtl), gen3639)

				gen3641 := Call(__e, __e.Global(symhd), gen3640)

				__e.TailApply(__e.Global(symcons), gen3638, gen3641)

				return

			} else {
				gen3637 := Call(__e, __e.Global(symcons_2), V1738)

				if True == gen3637 {
					gen3633 := Call(__e, __e.Global(symhd), V1738)

					gen3634 := Call(__e, __e.Global(symshen_4remove_1bar), gen3633)

					gen3635 := Call(__e, __e.Global(symtl), V1738)

					gen3636 := Call(__e, __e.Global(symshen_4remove_1bar), gen3635)

					__e.TailApply(__e.Global(symcons), gen3634, gen3636)

					return

				} else {
					__e.Return(V1738)
					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.remove-bar"), gen3659)

		gen3684 := MakeNative(func(__e Evaluator) {
			V1740 := __e.Get(1)
			_ = V1740
			gen3660 := Call(__e, __e.Global(symshen_4_5premise_6), V1740)

			Parse__shen_4_5premise_6 := gen3660
			gen3673 := Call(__e, __e.Global(symfail))

			gen3674 := Call(__e, __e.Global(sym_a), gen3673, Parse__shen_4_5premise_6)

			gen3675 := Call(__e, __e.Global(symnot), gen3674)

			var gen3676 Obj
			if True == gen3675 {
				gen3661 := Call(__e, __e.Global(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5premise_6)

				Parse__shen_4_5semicolon_1symbol_6 := gen3661
				gen3670 := Call(__e, __e.Global(symfail))

				gen3671 := Call(__e, __e.Global(sym_a), gen3670, Parse__shen_4_5semicolon_1symbol_6)

				gen3672 := Call(__e, __e.Global(symnot), gen3671)

				if True == gen3672 {
					gen3662 := Call(__e, __e.Global(symshen_4_5premises_6), Parse__shen_4_5semicolon_1symbol_6)

					Parse__shen_4_5premises_6 := gen3662
					gen3667 := Call(__e, __e.Global(symfail))

					gen3668 := Call(__e, __e.Global(sym_a), gen3667, Parse__shen_4_5premises_6)

					gen3669 := Call(__e, __e.Global(symnot), gen3668)

					if True == gen3669 {
						gen3663 := Call(__e, __e.Global(symhd), Parse__shen_4_5premises_6)

						gen3664 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5premise_6)

						gen3665 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5premises_6)

						gen3666 := Call(__e, __e.Global(symcons), gen3664, gen3665)

						gen3676 = Call(__e, __e.Global(symshen_4pair), gen3663, gen3666)

					} else {
						gen3676 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3676 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3676 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3676
			gen3682 := Call(__e, __e.Global(symfail))

			gen3683 := Call(__e, __e.Global(sym_a), YaccParse, gen3682)

			if True == gen3683 {
				gen3677 := Call(__e, __e.Global(sym_5e_6), V1740)

				Parse___5e_6 := gen3677
				gen3679 := Call(__e, __e.Global(symfail))

				gen3680 := Call(__e, __e.Global(sym_a), gen3679, Parse___5e_6)

				gen3681 := Call(__e, __e.Global(symnot), gen3680)

				if True == gen3681 {
					gen3678 := Call(__e, __e.Global(symhd), Parse___5e_6)

					__e.TailApply(__e.Global(symshen_4pair), gen3678, Nil)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<premises>"), gen3684)

		gen3693 := MakeNative(func(__e Evaluator) {
			V1742 := __e.Get(1)
			_ = V1742
			gen3691 := Call(__e, __e.Global(symhd), V1742)

			gen3692 := Call(__e, __e.Global(symcons_2), gen3691)

			if True == gen3692 {
				gen3685 := Call(__e, __e.Global(symshen_4hdhd), V1742)

				Parse__X := gen3685
				gen3690 := Call(__e, __e.Global(sym_a), Parse__X, MakeSymbol(";"))

				if True == gen3690 {
					gen3686 := Call(__e, __e.Global(symshen_4tlhd), V1742)

					gen3687 := Call(__e, __e.Global(symshen_4hdtl), V1742)

					gen3688 := Call(__e, __e.Global(symshen_4pair), gen3686, gen3687)

					gen3689 := Call(__e, __e.Global(symhd), gen3688)

					__e.TailApply(__e.Global(symshen_4pair), gen3689, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<semicolon-symbol>"), gen3693)

		gen3736 := MakeNative(func(__e Evaluator) {
			V1746 := __e.Get(1)
			_ = V1746
			gen3700 := Call(__e, __e.Global(symhd), V1746)

			gen3701 := Call(__e, __e.Global(symcons_2), gen3700)

			var gen3702 Obj
			if True == gen3701 {
				gen3698 := Call(__e, __e.Global(symshen_4hdhd), V1746)

				gen3699 := Call(__e, __e.Global(sym_a), MakeSymbol("!"), gen3698)

				if True == gen3699 {
					gen3702 = True
				} else {
					gen3702 = False
				}

			} else {
				gen3702 = False
			}
			var gen3703 Obj
			if True == gen3702 {
				gen3694 := Call(__e, __e.Global(symshen_4tlhd), V1746)

				gen3695 := Call(__e, __e.Global(symshen_4hdtl), V1746)

				gen3696 := Call(__e, __e.Global(symshen_4pair), gen3694, gen3695)

				NewStream1743 := gen3696
				gen3697 := Call(__e, __e.Global(symhd), NewStream1743)

				gen3703 = Call(__e, __e.Global(symshen_4pair), gen3697, MakeSymbol("!"))

			} else {
				gen3703 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3703
			gen3734 := Call(__e, __e.Global(symfail))

			gen3735 := Call(__e, __e.Global(sym_a), YaccParse, gen3734)

			if True == gen3735 {
				gen3704 := Call(__e, __e.Global(symshen_4_5formulae_6), V1746)

				Parse__shen_4_5formulae_6 := gen3704
				gen3721 := Call(__e, __e.Global(symfail))

				gen3722 := Call(__e, __e.Global(sym_a), gen3721, Parse__shen_4_5formulae_6)

				gen3723 := Call(__e, __e.Global(symnot), gen3722)

				var gen3724 Obj
				if True == gen3723 {
					gen3718 := Call(__e, __e.Global(symhd), Parse__shen_4_5formulae_6)

					gen3719 := Call(__e, __e.Global(symcons_2), gen3718)

					var gen3720 Obj
					if True == gen3719 {
						gen3716 := Call(__e, __e.Global(symshen_4hdhd), Parse__shen_4_5formulae_6)

						gen3717 := Call(__e, __e.Global(sym_a), MakeSymbol(">>"), gen3716)

						if True == gen3717 {
							gen3720 = True
						} else {
							gen3720 = False
						}

					} else {
						gen3720 = False
					}
					if True == gen3720 {
						gen3705 := Call(__e, __e.Global(symshen_4tlhd), Parse__shen_4_5formulae_6)

						gen3706 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formulae_6)

						gen3707 := Call(__e, __e.Global(symshen_4pair), gen3705, gen3706)

						NewStream1744 := gen3707
						gen3708 := Call(__e, __e.Global(symshen_4_5formula_6), NewStream1744)

						Parse__shen_4_5formula_6 := gen3708
						gen3713 := Call(__e, __e.Global(symfail))

						gen3714 := Call(__e, __e.Global(sym_a), gen3713, Parse__shen_4_5formula_6)

						gen3715 := Call(__e, __e.Global(symnot), gen3714)

						if True == gen3715 {
							gen3709 := Call(__e, __e.Global(symhd), Parse__shen_4_5formula_6)

							gen3710 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formulae_6)

							gen3711 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

							gen3712 := Call(__e, __e.Global(symshen_4sequent), gen3710, gen3711)

							gen3724 = Call(__e, __e.Global(symshen_4pair), gen3709, gen3712)

						} else {
							gen3724 = Call(__e, __e.Global(symfail))

						}

					} else {
						gen3724 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3724 = Call(__e, __e.Global(symfail))

				}
				YaccParse := gen3724
				gen3732 := Call(__e, __e.Global(symfail))

				gen3733 := Call(__e, __e.Global(sym_a), YaccParse, gen3732)

				if True == gen3733 {
					gen3725 := Call(__e, __e.Global(symshen_4_5formula_6), V1746)

					Parse__shen_4_5formula_6 := gen3725
					gen3729 := Call(__e, __e.Global(symfail))

					gen3730 := Call(__e, __e.Global(sym_a), gen3729, Parse__shen_4_5formula_6)

					gen3731 := Call(__e, __e.Global(symnot), gen3730)

					if True == gen3731 {
						gen3726 := Call(__e, __e.Global(symhd), Parse__shen_4_5formula_6)

						gen3727 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen3728 := Call(__e, __e.Global(symshen_4sequent), Nil, gen3727)

						__e.TailApply(__e.Global(symshen_4pair), gen3726, gen3728)

						return

					} else {
						__e.TailApply(__e.Global(symfail))

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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<premise>"), gen3736)

		gen3775 := MakeNative(func(__e Evaluator) {
			V1749 := __e.Get(1)
			_ = V1749
			gen3737 := Call(__e, __e.Global(symshen_4_5formulae_6), V1749)

			Parse__shen_4_5formulae_6 := gen3737
			gen3758 := Call(__e, __e.Global(symfail))

			gen3759 := Call(__e, __e.Global(sym_a), gen3758, Parse__shen_4_5formulae_6)

			gen3760 := Call(__e, __e.Global(symnot), gen3759)

			var gen3761 Obj
			if True == gen3760 {
				gen3755 := Call(__e, __e.Global(symhd), Parse__shen_4_5formulae_6)

				gen3756 := Call(__e, __e.Global(symcons_2), gen3755)

				var gen3757 Obj
				if True == gen3756 {
					gen3753 := Call(__e, __e.Global(symshen_4hdhd), Parse__shen_4_5formulae_6)

					gen3754 := Call(__e, __e.Global(sym_a), MakeSymbol(">>"), gen3753)

					if True == gen3754 {
						gen3757 = True
					} else {
						gen3757 = False
					}

				} else {
					gen3757 = False
				}
				if True == gen3757 {
					gen3738 := Call(__e, __e.Global(symshen_4tlhd), Parse__shen_4_5formulae_6)

					gen3739 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formulae_6)

					gen3740 := Call(__e, __e.Global(symshen_4pair), gen3738, gen3739)

					NewStream1747 := gen3740
					gen3741 := Call(__e, __e.Global(symshen_4_5formula_6), NewStream1747)

					Parse__shen_4_5formula_6 := gen3741
					gen3750 := Call(__e, __e.Global(symfail))

					gen3751 := Call(__e, __e.Global(sym_a), gen3750, Parse__shen_4_5formula_6)

					gen3752 := Call(__e, __e.Global(symnot), gen3751)

					if True == gen3752 {
						gen3742 := Call(__e, __e.Global(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

						Parse__shen_4_5semicolon_1symbol_6 := gen3742
						gen3747 := Call(__e, __e.Global(symfail))

						gen3748 := Call(__e, __e.Global(sym_a), gen3747, Parse__shen_4_5semicolon_1symbol_6)

						gen3749 := Call(__e, __e.Global(symnot), gen3748)

						if True == gen3749 {
							gen3743 := Call(__e, __e.Global(symhd), Parse__shen_4_5semicolon_1symbol_6)

							gen3744 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formulae_6)

							gen3745 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

							gen3746 := Call(__e, __e.Global(symshen_4sequent), gen3744, gen3745)

							gen3761 = Call(__e, __e.Global(symshen_4pair), gen3743, gen3746)

						} else {
							gen3761 = Call(__e, __e.Global(symfail))

						}

					} else {
						gen3761 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3761 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3761 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3761
			gen3773 := Call(__e, __e.Global(symfail))

			gen3774 := Call(__e, __e.Global(sym_a), YaccParse, gen3773)

			if True == gen3774 {
				gen3762 := Call(__e, __e.Global(symshen_4_5formula_6), V1749)

				Parse__shen_4_5formula_6 := gen3762
				gen3770 := Call(__e, __e.Global(symfail))

				gen3771 := Call(__e, __e.Global(sym_a), gen3770, Parse__shen_4_5formula_6)

				gen3772 := Call(__e, __e.Global(symnot), gen3771)

				if True == gen3772 {
					gen3763 := Call(__e, __e.Global(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

					Parse__shen_4_5semicolon_1symbol_6 := gen3763
					gen3767 := Call(__e, __e.Global(symfail))

					gen3768 := Call(__e, __e.Global(sym_a), gen3767, Parse__shen_4_5semicolon_1symbol_6)

					gen3769 := Call(__e, __e.Global(symnot), gen3768)

					if True == gen3769 {
						gen3764 := Call(__e, __e.Global(symhd), Parse__shen_4_5semicolon_1symbol_6)

						gen3765 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen3766 := Call(__e, __e.Global(symshen_4sequent), Nil, gen3765)

						__e.TailApply(__e.Global(symshen_4pair), gen3764, gen3766)

						return

					} else {
						__e.TailApply(__e.Global(symfail))

						return
					}

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<conclusion>"), gen3775)

		gen3776 := MakeNative(func(__e Evaluator) {
			V1752 := __e.Get(1)
			_ = V1752
			V1753 := __e.Get(2)
			_ = V1753
			__e.TailApply(__e.Global(sym_8p), V1752, V1753)

			return
		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.sequent"), gen3776)

		gen3811 := MakeNative(func(__e Evaluator) {
			V1755 := __e.Get(1)
			_ = V1755
			gen3777 := Call(__e, __e.Global(symshen_4_5formula_6), V1755)

			Parse__shen_4_5formula_6 := gen3777
			gen3790 := Call(__e, __e.Global(symfail))

			gen3791 := Call(__e, __e.Global(sym_a), gen3790, Parse__shen_4_5formula_6)

			gen3792 := Call(__e, __e.Global(symnot), gen3791)

			var gen3793 Obj
			if True == gen3792 {
				gen3778 := Call(__e, __e.Global(symshen_4_5comma_1symbol_6), Parse__shen_4_5formula_6)

				Parse__shen_4_5comma_1symbol_6 := gen3778
				gen3787 := Call(__e, __e.Global(symfail))

				gen3788 := Call(__e, __e.Global(sym_a), gen3787, Parse__shen_4_5comma_1symbol_6)

				gen3789 := Call(__e, __e.Global(symnot), gen3788)

				if True == gen3789 {
					gen3779 := Call(__e, __e.Global(symshen_4_5formulae_6), Parse__shen_4_5comma_1symbol_6)

					Parse__shen_4_5formulae_6 := gen3779
					gen3784 := Call(__e, __e.Global(symfail))

					gen3785 := Call(__e, __e.Global(sym_a), gen3784, Parse__shen_4_5formulae_6)

					gen3786 := Call(__e, __e.Global(symnot), gen3785)

					if True == gen3786 {
						gen3780 := Call(__e, __e.Global(symhd), Parse__shen_4_5formulae_6)

						gen3781 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen3782 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formulae_6)

						gen3783 := Call(__e, __e.Global(symcons), gen3781, gen3782)

						gen3793 = Call(__e, __e.Global(symshen_4pair), gen3780, gen3783)

					} else {
						gen3793 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3793 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3793 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3793
			gen3809 := Call(__e, __e.Global(symfail))

			gen3810 := Call(__e, __e.Global(sym_a), YaccParse, gen3809)

			if True == gen3810 {
				gen3794 := Call(__e, __e.Global(symshen_4_5formula_6), V1755)

				Parse__shen_4_5formula_6 := gen3794
				gen3798 := Call(__e, __e.Global(symfail))

				gen3799 := Call(__e, __e.Global(sym_a), gen3798, Parse__shen_4_5formula_6)

				gen3800 := Call(__e, __e.Global(symnot), gen3799)

				var gen3801 Obj
				if True == gen3800 {
					gen3795 := Call(__e, __e.Global(symhd), Parse__shen_4_5formula_6)

					gen3796 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5formula_6)

					gen3797 := Call(__e, __e.Global(symcons), gen3796, Nil)

					gen3801 = Call(__e, __e.Global(symshen_4pair), gen3795, gen3797)

				} else {
					gen3801 = Call(__e, __e.Global(symfail))

				}
				YaccParse := gen3801
				gen3807 := Call(__e, __e.Global(symfail))

				gen3808 := Call(__e, __e.Global(sym_a), YaccParse, gen3807)

				if True == gen3808 {
					gen3802 := Call(__e, __e.Global(sym_5e_6), V1755)

					Parse___5e_6 := gen3802
					gen3804 := Call(__e, __e.Global(symfail))

					gen3805 := Call(__e, __e.Global(sym_a), gen3804, Parse___5e_6)

					gen3806 := Call(__e, __e.Global(symnot), gen3805)

					if True == gen3806 {
						gen3803 := Call(__e, __e.Global(symhd), Parse___5e_6)

						__e.TailApply(__e.Global(symshen_4pair), gen3803, Nil)

						return

					} else {
						__e.TailApply(__e.Global(symfail))

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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<formulae>"), gen3811)

		gen3821 := MakeNative(func(__e Evaluator) {
			V1757 := __e.Get(1)
			_ = V1757
			gen3819 := Call(__e, __e.Global(symhd), V1757)

			gen3820 := Call(__e, __e.Global(symcons_2), gen3819)

			if True == gen3820 {
				gen3812 := Call(__e, __e.Global(symshen_4hdhd), V1757)

				Parse__X := gen3812
				gen3817 := Call(__e, __e.Global(symintern), MakeString(","))

				gen3818 := Call(__e, __e.Global(sym_a), Parse__X, gen3817)

				if True == gen3818 {
					gen3813 := Call(__e, __e.Global(symshen_4tlhd), V1757)

					gen3814 := Call(__e, __e.Global(symshen_4hdtl), V1757)

					gen3815 := Call(__e, __e.Global(symshen_4pair), gen3813, gen3814)

					gen3816 := Call(__e, __e.Global(symhd), gen3815)

					__e.TailApply(__e.Global(symshen_4pair), gen3816, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<comma-symbol>"), gen3821)

		gen3855 := MakeNative(func(__e Evaluator) {
			V1760 := __e.Get(1)
			_ = V1760
			gen3822 := Call(__e, __e.Global(symshen_4_5expr_6), V1760)

			Parse__shen_4_5expr_6 := gen3822
			gen3843 := Call(__e, __e.Global(symfail))

			gen3844 := Call(__e, __e.Global(sym_a), gen3843, Parse__shen_4_5expr_6)

			gen3845 := Call(__e, __e.Global(symnot), gen3844)

			var gen3846 Obj
			if True == gen3845 {
				gen3840 := Call(__e, __e.Global(symhd), Parse__shen_4_5expr_6)

				gen3841 := Call(__e, __e.Global(symcons_2), gen3840)

				var gen3842 Obj
				if True == gen3841 {
					gen3838 := Call(__e, __e.Global(symshen_4hdhd), Parse__shen_4_5expr_6)

					gen3839 := Call(__e, __e.Global(sym_a), MakeSymbol(":"), gen3838)

					if True == gen3839 {
						gen3842 = True
					} else {
						gen3842 = False
					}

				} else {
					gen3842 = False
				}
				if True == gen3842 {
					gen3823 := Call(__e, __e.Global(symshen_4tlhd), Parse__shen_4_5expr_6)

					gen3824 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

					gen3825 := Call(__e, __e.Global(symshen_4pair), gen3823, gen3824)

					NewStream1758 := gen3825
					gen3826 := Call(__e, __e.Global(symshen_4_5type_6), NewStream1758)

					Parse__shen_4_5type_6 := gen3826
					gen3835 := Call(__e, __e.Global(symfail))

					gen3836 := Call(__e, __e.Global(sym_a), gen3835, Parse__shen_4_5type_6)

					gen3837 := Call(__e, __e.Global(symnot), gen3836)

					if True == gen3837 {
						gen3827 := Call(__e, __e.Global(symhd), Parse__shen_4_5type_6)

						gen3828 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

						gen3829 := Call(__e, __e.Global(symshen_4curry), gen3828)

						gen3830 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5type_6)

						gen3831 := Call(__e, __e.Global(symshen_4demodulate), gen3830)

						gen3832 := Call(__e, __e.Global(symcons), gen3831, Nil)

						gen3833 := Call(__e, __e.Global(symcons), MakeSymbol(":"), gen3832)

						gen3834 := Call(__e, __e.Global(symcons), gen3829, gen3833)

						gen3846 = Call(__e, __e.Global(symshen_4pair), gen3827, gen3834)

					} else {
						gen3846 = Call(__e, __e.Global(symfail))

					}

				} else {
					gen3846 = Call(__e, __e.Global(symfail))

				}

			} else {
				gen3846 = Call(__e, __e.Global(symfail))

			}
			YaccParse := gen3846
			gen3853 := Call(__e, __e.Global(symfail))

			gen3854 := Call(__e, __e.Global(sym_a), YaccParse, gen3853)

			if True == gen3854 {
				gen3847 := Call(__e, __e.Global(symshen_4_5expr_6), V1760)

				Parse__shen_4_5expr_6 := gen3847
				gen3850 := Call(__e, __e.Global(symfail))

				gen3851 := Call(__e, __e.Global(sym_a), gen3850, Parse__shen_4_5expr_6)

				gen3852 := Call(__e, __e.Global(symnot), gen3851)

				if True == gen3852 {
					gen3848 := Call(__e, __e.Global(symhd), Parse__shen_4_5expr_6)

					gen3849 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

					__e.TailApply(__e.Global(symshen_4pair), gen3848, gen3849)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<formula>"), gen3855)

		gen3863 := MakeNative(func(__e Evaluator) {
			V1762 := __e.Get(1)
			_ = V1762
			gen3856 := Call(__e, __e.Global(symshen_4_5expr_6), V1762)

			Parse__shen_4_5expr_6 := gen3856
			gen3860 := Call(__e, __e.Global(symfail))

			gen3861 := Call(__e, __e.Global(sym_a), gen3860, Parse__shen_4_5expr_6)

			gen3862 := Call(__e, __e.Global(symnot), gen3861)

			if True == gen3862 {
				gen3857 := Call(__e, __e.Global(symhd), Parse__shen_4_5expr_6)

				gen3858 := Call(__e, __e.Global(symshen_4hdtl), Parse__shen_4_5expr_6)

				gen3859 := Call(__e, __e.Global(symshen_4curry_1type), gen3858)

				__e.TailApply(__e.Global(symshen_4pair), gen3857, gen3859)

				return

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<type>"), gen3863)

		gen3872 := MakeNative(func(__e Evaluator) {
			V1764 := __e.Get(1)
			_ = V1764
			gen3870 := Call(__e, __e.Global(symhd), V1764)

			gen3871 := Call(__e, __e.Global(symcons_2), gen3870)

			if True == gen3871 {
				gen3864 := Call(__e, __e.Global(symshen_4hdhd), V1764)

				Parse__X := gen3864
				gen3869 := Call(__e, __e.Global(symshen_4doubleunderline_2), Parse__X)

				if True == gen3869 {
					gen3865 := Call(__e, __e.Global(symshen_4tlhd), V1764)

					gen3866 := Call(__e, __e.Global(symshen_4hdtl), V1764)

					gen3867 := Call(__e, __e.Global(symshen_4pair), gen3865, gen3866)

					gen3868 := Call(__e, __e.Global(symhd), gen3867)

					__e.TailApply(__e.Global(symshen_4pair), gen3868, Parse__X)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<doubleunderline>"), gen3872)

		gen3881 := MakeNative(func(__e Evaluator) {
			V1766 := __e.Get(1)
			_ = V1766
			gen3879 := Call(__e, __e.Global(symhd), V1766)

			gen3880 := Call(__e, __e.Global(symcons_2), gen3879)

			if True == gen3880 {
				gen3873 := Call(__e, __e.Global(symshen_4hdhd), V1766)

				Parse__X := gen3873
				gen3878 := Call(__e, __e.Global(symshen_4singleunderline_2), Parse__X)

				if True == gen3878 {
					gen3874 := Call(__e, __e.Global(symshen_4tlhd), V1766)

					gen3875 := Call(__e, __e.Global(symshen_4hdtl), V1766)

					gen3876 := Call(__e, __e.Global(symshen_4pair), gen3874, gen3875)

					gen3877 := Call(__e, __e.Global(symhd), gen3876)

					__e.TailApply(__e.Global(symshen_4pair), gen3877, Parse__X)

					return

				} else {
					__e.TailApply(__e.Global(symfail))

					return
				}

			} else {
				__e.TailApply(__e.Global(symfail))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.<singleunderline>"), gen3881)

		gen3885 := MakeNative(func(__e Evaluator) {
			V1768 := __e.Get(1)
			_ = V1768
			gen3884 := Call(__e, __e.Global(symsymbol_2), V1768)

			if True == gen3884 {
				gen3882 := Call(__e, __e.Global(symstr), V1768)

				gen3883 := Call(__e, __e.Global(symshen_4sh_2), gen3882)

				if True == gen3883 {
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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.singleunderline?"), gen3885)

		gen3891 := MakeNative(func(__e Evaluator) {
			V1770 := __e.Get(1)
			_ = V1770
			gen3890 := Call(__e, __e.Global(sym_a), MakeString("_"), V1770)

			if True == gen3890 {
				__e.Return(True)
				return
			} else {
				gen3888 := Call(__e, __e.Global(sympos), V1770, MakeNumber(0))

				gen3889 := Call(__e, __e.Global(sym_a), gen3888, MakeString("_"))

				if True == gen3889 {
					gen3886 := Call(__e, __e.Global(symtlstr), V1770)

					gen3887 := Call(__e, __e.Global(symshen_4sh_2), gen3886)

					if True == gen3887 {
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

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.sh?"), gen3891)

		gen3895 := MakeNative(func(__e Evaluator) {
			V1772 := __e.Get(1)
			_ = V1772
			gen3894 := Call(__e, __e.Global(symsymbol_2), V1772)

			if True == gen3894 {
				gen3892 := Call(__e, __e.Global(symstr), V1772)

				gen3893 := Call(__e, __e.Global(symshen_4dh_2), gen3892)

				if True == gen3893 {
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
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.doubleunderline?"), gen3895)

		gen3901 := MakeNative(func(__e Evaluator) {
			V1774 := __e.Get(1)
			_ = V1774
			gen3900 := Call(__e, __e.Global(sym_a), MakeString("="), V1774)

			if True == gen3900 {
				__e.Return(True)
				return
			} else {
				gen3898 := Call(__e, __e.Global(sympos), V1774, MakeNumber(0))

				gen3899 := Call(__e, __e.Global(sym_a), gen3898, MakeString("="))

				if True == gen3899 {
					gen3896 := Call(__e, __e.Global(symtlstr), V1774)

					gen3897 := Call(__e, __e.Global(symshen_4dh_2), gen3896)

					if True == gen3897 {
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

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.dh?"), gen3901)

		gen3904 := MakeNative(func(__e Evaluator) {
			V1777 := __e.Get(1)
			_ = V1777
			V1778 := __e.Get(2)
			_ = V1778
			gen3902 := Call(__e, __e.Global(symshen_4rules_1_6horn_1clauses), V1777, V1778)

			gen3903 := Call(__e, __e.Global(symshen_4s_1prolog), gen3902)

			__e.TailApply(__e.Global(symshen_4remember_1datatype), gen3903)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.process-datatype"), gen3904)

		gen3912 := MakeNative(func(__e Evaluator) {
			V1784 := __e.Get(1)
			_ = V1784
			gen3911 := Call(__e, __e.Global(symcons_2), V1784)

			if True == gen3911 {
				gen3905 := Call(__e, __e.Global(symhd), V1784)

				gen3906 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*datatypes*"))

				gen3907 := Call(__e, __e.Global(symadjoin), gen3905, gen3906)

				Call(__e, __e.Global(symset), MakeSymbol("shen.*datatypes*"), gen3907)

				gen3908 := Call(__e, __e.Global(symhd), V1784)

				gen3909 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*alldatatypes*"))

				gen3910 := Call(__e, __e.Global(symadjoin), gen3908, gen3909)

				Call(__e, __e.Global(symset), MakeSymbol("shen.*alldatatypes*"), gen3910)

				__e.TailApply(__e.Global(symhd), V1784)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.remember-datatype"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.remember-datatype"), gen3912)

		gen3940 := MakeNative(func(__e Evaluator) {
			V1789 := __e.Get(1)
			_ = V1789
			V1790 := __e.Get(2)
			_ = V1790
			gen3939 := Call(__e, __e.Global(sym_a), Nil, V1790)

			if True == gen3939 {
				__e.Return(Nil)
				return
			} else {
				gen3937 := Call(__e, __e.Global(symcons_2), V1790)

				var gen3938 Obj
				if True == gen3937 {
					gen3934 := Call(__e, __e.Global(symhd), V1790)

					gen3935 := Call(__e, __e.Global(symtuple_2), gen3934)

					var gen3936 Obj
					if True == gen3935 {
						gen3931 := Call(__e, __e.Global(symhd), V1790)

						gen3932 := Call(__e, __e.Global(symfst), gen3931)

						gen3933 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.single"), gen3932)

						if True == gen3933 {
							gen3936 = True
						} else {
							gen3936 = False
						}

					} else {
						gen3936 = False
					}
					if True == gen3936 {
						gen3938 = True
					} else {
						gen3938 = False
					}

				} else {
					gen3938 = False
				}
				if True == gen3938 {
					gen3926 := Call(__e, __e.Global(symhd), V1790)

					gen3927 := Call(__e, __e.Global(symsnd), gen3926)

					gen3928 := Call(__e, __e.Global(symshen_4rule_1_6horn_1clause), V1789, gen3927)

					gen3929 := Call(__e, __e.Global(symtl), V1790)

					gen3930 := Call(__e, __e.Global(symshen_4rules_1_6horn_1clauses), V1789, gen3929)

					__e.TailApply(__e.Global(symcons), gen3928, gen3930)

					return

				} else {
					gen3924 := Call(__e, __e.Global(symcons_2), V1790)

					var gen3925 Obj
					if True == gen3924 {
						gen3921 := Call(__e, __e.Global(symhd), V1790)

						gen3922 := Call(__e, __e.Global(symtuple_2), gen3921)

						var gen3923 Obj
						if True == gen3922 {
							gen3918 := Call(__e, __e.Global(symhd), V1790)

							gen3919 := Call(__e, __e.Global(symfst), gen3918)

							gen3920 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.double"), gen3919)

							if True == gen3920 {
								gen3923 = True
							} else {
								gen3923 = False
							}

						} else {
							gen3923 = False
						}
						if True == gen3923 {
							gen3925 = True
						} else {
							gen3925 = False
						}

					} else {
						gen3925 = False
					}
					if True == gen3925 {
						gen3913 := Call(__e, __e.Global(symhd), V1790)

						gen3914 := Call(__e, __e.Global(symsnd), gen3913)

						gen3915 := Call(__e, __e.Global(symshen_4double_1_6singles), gen3914)

						gen3916 := Call(__e, __e.Global(symtl), V1790)

						gen3917 := Call(__e, __e.Global(symappend), gen3915, gen3916)

						__e.TailApply(__e.Global(symshen_4rules_1_6horn_1clauses), V1789, gen3917)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.rules->horn-clauses"))

						return
					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.rules->horn-clauses"), gen3940)

		gen3944 := MakeNative(func(__e Evaluator) {
			V1792 := __e.Get(1)
			_ = V1792
			gen3941 := Call(__e, __e.Global(symshen_4right_1rule), V1792)

			gen3942 := Call(__e, __e.Global(symshen_4left_1rule), V1792)

			gen3943 := Call(__e, __e.Global(symcons), gen3942, Nil)

			__e.TailApply(__e.Global(symcons), gen3941, gen3943)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.double->singles"), gen3944)

		gen3945 := MakeNative(func(__e Evaluator) {
			V1794 := __e.Get(1)
			_ = V1794
			__e.TailApply(__e.Global(sym_8p), MakeSymbol("shen.single"), V1794)

			return
		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.right-rule"), gen3945)

		gen3987 := MakeNative(func(__e Evaluator) {
			V1796 := __e.Get(1)
			_ = V1796
			gen3985 := Call(__e, __e.Global(symcons_2), V1796)

			var gen3986 Obj
			if True == gen3985 {
				gen3982 := Call(__e, __e.Global(symtl), V1796)

				gen3983 := Call(__e, __e.Global(symcons_2), gen3982)

				var gen3984 Obj
				if True == gen3983 {
					gen3978 := Call(__e, __e.Global(symtl), V1796)

					gen3979 := Call(__e, __e.Global(symtl), gen3978)

					gen3980 := Call(__e, __e.Global(symcons_2), gen3979)

					var gen3981 Obj
					if True == gen3980 {
						gen3973 := Call(__e, __e.Global(symtl), V1796)

						gen3974 := Call(__e, __e.Global(symtl), gen3973)

						gen3975 := Call(__e, __e.Global(symhd), gen3974)

						gen3976 := Call(__e, __e.Global(symtuple_2), gen3975)

						var gen3977 Obj
						if True == gen3976 {
							gen3967 := Call(__e, __e.Global(symtl), V1796)

							gen3968 := Call(__e, __e.Global(symtl), gen3967)

							gen3969 := Call(__e, __e.Global(symhd), gen3968)

							gen3970 := Call(__e, __e.Global(symfst), gen3969)

							gen3971 := Call(__e, __e.Global(sym_a), Nil, gen3970)

							var gen3972 Obj
							if True == gen3971 {
								gen3963 := Call(__e, __e.Global(symtl), V1796)

								gen3964 := Call(__e, __e.Global(symtl), gen3963)

								gen3965 := Call(__e, __e.Global(symtl), gen3964)

								gen3966 := Call(__e, __e.Global(sym_a), Nil, gen3965)

								if True == gen3966 {
									gen3972 = True
								} else {
									gen3972 = False
								}

							} else {
								gen3972 = False
							}
							if True == gen3972 {
								gen3977 = True
							} else {
								gen3977 = False
							}

						} else {
							gen3977 = False
						}
						if True == gen3977 {
							gen3981 = True
						} else {
							gen3981 = False
						}

					} else {
						gen3981 = False
					}
					if True == gen3981 {
						gen3984 = True
					} else {
						gen3984 = False
					}

				} else {
					gen3984 = False
				}
				if True == gen3984 {
					gen3986 = True
				} else {
					gen3986 = False
				}

			} else {
				gen3986 = False
			}
			if True == gen3986 {
				gen3946 := Call(__e, __e.Global(symgensym), MakeSymbol("Qv"))

				Q := gen3946
				gen3947 := Call(__e, __e.Global(symtl), V1796)

				gen3948 := Call(__e, __e.Global(symtl), gen3947)

				gen3949 := Call(__e, __e.Global(symhd), gen3948)

				gen3950 := Call(__e, __e.Global(symsnd), gen3949)

				gen3951 := Call(__e, __e.Global(symcons), gen3950, Nil)

				gen3952 := Call(__e, __e.Global(sym_8p), gen3951, Q)

				NewConclusion := gen3952
				gen3953 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4right_1_6left), X)

					return
				}, 1)
				gen3954 := Call(__e, __e.Global(symtl), V1796)

				gen3955 := Call(__e, __e.Global(symhd), gen3954)

				gen3956 := Call(__e, __e.Global(symmap), gen3953, gen3955)

				gen3957 := Call(__e, __e.Global(sym_8p), gen3956, Q)

				gen3958 := Call(__e, __e.Global(symcons), gen3957, Nil)

				NewPremises := gen3958
				gen3959 := Call(__e, __e.Global(symhd), V1796)

				gen3960 := Call(__e, __e.Global(symcons), NewConclusion, Nil)

				gen3961 := Call(__e, __e.Global(symcons), NewPremises, gen3960)

				gen3962 := Call(__e, __e.Global(symcons), gen3959, gen3961)

				__e.TailApply(__e.Global(sym_8p), MakeSymbol("shen.single"), gen3962)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.left-rule"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.left-rule"), gen3987)

		gen3992 := MakeNative(func(__e Evaluator) {
			V1802 := __e.Get(1)
			_ = V1802
			gen3990 := Call(__e, __e.Global(symtuple_2), V1802)

			var gen3991 Obj
			if True == gen3990 {
				gen3988 := Call(__e, __e.Global(symfst), V1802)

				gen3989 := Call(__e, __e.Global(sym_a), Nil, gen3988)

				if True == gen3989 {
					gen3991 = True
				} else {
					gen3991 = False
				}

			} else {
				gen3991 = False
			}
			if True == gen3991 {
				__e.TailApply(__e.Global(symsnd), V1802)

				return
			} else {
				__e.TailApply(__e.Global(symsimple_1error), MakeString("syntax error with ==========\n"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.right->left"), gen3992)

		gen4026 := MakeNative(func(__e Evaluator) {
			V1805 := __e.Get(1)
			_ = V1805
			V1806 := __e.Get(2)
			_ = V1806
			gen4024 := Call(__e, __e.Global(symcons_2), V1806)

			var gen4025 Obj
			if True == gen4024 {
				gen4021 := Call(__e, __e.Global(symtl), V1806)

				gen4022 := Call(__e, __e.Global(symcons_2), gen4021)

				var gen4023 Obj
				if True == gen4022 {
					gen4017 := Call(__e, __e.Global(symtl), V1806)

					gen4018 := Call(__e, __e.Global(symtl), gen4017)

					gen4019 := Call(__e, __e.Global(symcons_2), gen4018)

					var gen4020 Obj
					if True == gen4019 {
						gen4012 := Call(__e, __e.Global(symtl), V1806)

						gen4013 := Call(__e, __e.Global(symtl), gen4012)

						gen4014 := Call(__e, __e.Global(symhd), gen4013)

						gen4015 := Call(__e, __e.Global(symtuple_2), gen4014)

						var gen4016 Obj
						if True == gen4015 {
							gen4008 := Call(__e, __e.Global(symtl), V1806)

							gen4009 := Call(__e, __e.Global(symtl), gen4008)

							gen4010 := Call(__e, __e.Global(symtl), gen4009)

							gen4011 := Call(__e, __e.Global(sym_a), Nil, gen4010)

							if True == gen4011 {
								gen4016 = True
							} else {
								gen4016 = False
							}

						} else {
							gen4016 = False
						}
						if True == gen4016 {
							gen4020 = True
						} else {
							gen4020 = False
						}

					} else {
						gen4020 = False
					}
					if True == gen4020 {
						gen4023 = True
					} else {
						gen4023 = False
					}

				} else {
					gen4023 = False
				}
				if True == gen4023 {
					gen4025 = True
				} else {
					gen4025 = False
				}

			} else {
				gen4025 = False
			}
			if True == gen4025 {
				gen3993 := Call(__e, __e.Global(symtl), V1806)

				gen3994 := Call(__e, __e.Global(symtl), gen3993)

				gen3995 := Call(__e, __e.Global(symhd), gen3994)

				gen3996 := Call(__e, __e.Global(symsnd), gen3995)

				gen3997 := Call(__e, __e.Global(symshen_4rule_1_6horn_1clause_1head), V1805, gen3996)

				gen3998 := Call(__e, __e.Global(symhd), V1806)

				gen3999 := Call(__e, __e.Global(symtl), V1806)

				gen4000 := Call(__e, __e.Global(symhd), gen3999)

				gen4001 := Call(__e, __e.Global(symtl), V1806)

				gen4002 := Call(__e, __e.Global(symtl), gen4001)

				gen4003 := Call(__e, __e.Global(symhd), gen4002)

				gen4004 := Call(__e, __e.Global(symfst), gen4003)

				gen4005 := Call(__e, __e.Global(symshen_4rule_1_6horn_1clause_1body), gen3998, gen4000, gen4004)

				gen4006 := Call(__e, __e.Global(symcons), gen4005, Nil)

				gen4007 := Call(__e, __e.Global(symcons), MakeSymbol(":-"), gen4006)

				__e.TailApply(__e.Global(symcons), gen3997, gen4007)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.rule->horn-clause"))

				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.rule->horn-clause"), gen4026)

		gen4030 := MakeNative(func(__e Evaluator) {
			V1809 := __e.Get(1)
			_ = V1809
			V1810 := __e.Get(2)
			_ = V1810
			gen4027 := Call(__e, __e.Global(symshen_4mode_1ify), V1810)

			gen4028 := Call(__e, __e.Global(symcons), MakeSymbol("Context_1957"), Nil)

			gen4029 := Call(__e, __e.Global(symcons), gen4027, gen4028)

			__e.TailApply(__e.Global(symcons), V1809, gen4029)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.rule->horn-clause-head"), gen4030)

		gen4060 := MakeNative(func(__e Evaluator) {
			V1812 := __e.Get(1)
			_ = V1812
			gen4058 := Call(__e, __e.Global(symcons_2), V1812)

			var gen4059 Obj
			if True == gen4058 {
				gen4055 := Call(__e, __e.Global(symtl), V1812)

				gen4056 := Call(__e, __e.Global(symcons_2), gen4055)

				var gen4057 Obj
				if True == gen4056 {
					gen4051 := Call(__e, __e.Global(symtl), V1812)

					gen4052 := Call(__e, __e.Global(symhd), gen4051)

					gen4053 := Call(__e, __e.Global(sym_a), MakeSymbol(":"), gen4052)

					var gen4054 Obj
					if True == gen4053 {
						gen4047 := Call(__e, __e.Global(symtl), V1812)

						gen4048 := Call(__e, __e.Global(symtl), gen4047)

						gen4049 := Call(__e, __e.Global(symcons_2), gen4048)

						var gen4050 Obj
						if True == gen4049 {
							gen4043 := Call(__e, __e.Global(symtl), V1812)

							gen4044 := Call(__e, __e.Global(symtl), gen4043)

							gen4045 := Call(__e, __e.Global(symtl), gen4044)

							gen4046 := Call(__e, __e.Global(sym_a), Nil, gen4045)

							if True == gen4046 {
								gen4050 = True
							} else {
								gen4050 = False
							}

						} else {
							gen4050 = False
						}
						if True == gen4050 {
							gen4054 = True
						} else {
							gen4054 = False
						}

					} else {
						gen4054 = False
					}
					if True == gen4054 {
						gen4057 = True
					} else {
						gen4057 = False
					}

				} else {
					gen4057 = False
				}
				if True == gen4057 {
					gen4059 = True
				} else {
					gen4059 = False
				}

			} else {
				gen4059 = False
			}
			if True == gen4059 {
				gen4031 := Call(__e, __e.Global(symhd), V1812)

				gen4032 := Call(__e, __e.Global(symtl), V1812)

				gen4033 := Call(__e, __e.Global(symtl), gen4032)

				gen4034 := Call(__e, __e.Global(symhd), gen4033)

				gen4035 := Call(__e, __e.Global(symcons), MakeSymbol("+"), Nil)

				gen4036 := Call(__e, __e.Global(symcons), gen4034, gen4035)

				gen4037 := Call(__e, __e.Global(symcons), MakeSymbol("mode"), gen4036)

				gen4038 := Call(__e, __e.Global(symcons), gen4037, Nil)

				gen4039 := Call(__e, __e.Global(symcons), MakeSymbol(":"), gen4038)

				gen4040 := Call(__e, __e.Global(symcons), gen4031, gen4039)

				gen4041 := Call(__e, __e.Global(symcons), MakeSymbol("-"), Nil)

				gen4042 := Call(__e, __e.Global(symcons), gen4040, gen4041)

				__e.TailApply(__e.Global(symcons), MakeSymbol("mode"), gen4042)

				return

			} else {
				__e.Return(V1812)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.mode-ify"), gen4060)

		gen4072 := MakeNative(func(__e Evaluator) {
			V1816 := __e.Get(1)
			_ = V1816
			V1817 := __e.Get(2)
			_ = V1817
			V1818 := __e.Get(3)
			_ = V1818
			gen4061 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4extract__vars), X)

				return
			}, 1)
			gen4062 := Call(__e, __e.Global(symmap), gen4061, V1818)

			Variables := gen4062
			gen4063 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symgensym), MakeSymbol("shen.cl"))

				return
			}, 1)
			gen4064 := Call(__e, __e.Global(symmap), gen4063, V1818)

			Predicates := gen4064
			gen4065 := Call(__e, __e.Global(symshen_4construct_1search_1literals), Predicates, Variables, MakeSymbol("Context_1957"), MakeSymbol("Context1_1957"))

			SearchLiterals := gen4065
			gen4066 := Call(__e, __e.Global(symshen_4construct_1search_1clauses), Predicates, V1818, Variables)

			SearchClauses := gen4066
			_ = SearchClauses
			gen4067 := Call(__e, __e.Global(symshen_4construct_1side_1literals), V1816)

			SideLiterals := gen4067
			gen4069 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				gen4068 := Call(__e, __e.Global(symempty_2), V1818)

				__e.TailApply(__e.Global(symshen_4construct_1premiss_1literal), X, gen4068)

				return

			}, 1)
			gen4070 := Call(__e, __e.Global(symmap), gen4069, V1817)

			PremissLiterals := gen4070
			gen4071 := Call(__e, __e.Global(symappend), SideLiterals, PremissLiterals)

			__e.TailApply(__e.Global(symappend), SearchLiterals, gen4071)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.rule->horn-clause-body"), gen4072)

		gen4076 := MakeNative(func(__e Evaluator) {
			V1827 := __e.Get(1)
			_ = V1827
			V1828 := __e.Get(2)
			_ = V1828
			V1829 := __e.Get(3)
			_ = V1829
			V1830 := __e.Get(4)
			_ = V1830
			gen4074 := Call(__e, __e.Global(sym_a), Nil, V1827)

			var gen4075 Obj
			if True == gen4074 {
				gen4073 := Call(__e, __e.Global(sym_a), Nil, V1828)

				if True == gen4073 {
					gen4075 = True
				} else {
					gen4075 = False
				}

			} else {
				gen4075 = False
			}
			if True == gen4075 {
				__e.Return(Nil)
				return
			} else {
				__e.TailApply(__e.Global(symshen_4csl_1help), V1827, V1828, V1829, V1830)

				return
			}

		}, 4)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-search-literals"), gen4076)

		gen4095 := MakeNative(func(__e Evaluator) {
			V1837 := __e.Get(1)
			_ = V1837
			V1838 := __e.Get(2)
			_ = V1838
			V1839 := __e.Get(3)
			_ = V1839
			V1840 := __e.Get(4)
			_ = V1840
			gen4093 := Call(__e, __e.Global(sym_a), Nil, V1837)

			var gen4094 Obj
			if True == gen4093 {
				gen4092 := Call(__e, __e.Global(sym_a), Nil, V1838)

				if True == gen4092 {
					gen4094 = True
				} else {
					gen4094 = False
				}

			} else {
				gen4094 = False
			}
			if True == gen4094 {
				gen4089 := Call(__e, __e.Global(symcons), V1839, Nil)

				gen4090 := Call(__e, __e.Global(symcons), MakeSymbol("ContextOut_1957"), gen4089)

				gen4091 := Call(__e, __e.Global(symcons), MakeSymbol("bind"), gen4090)

				__e.TailApply(__e.Global(symcons), gen4091, Nil)

				return

			} else {
				gen4087 := Call(__e, __e.Global(symcons_2), V1837)

				var gen4088 Obj
				if True == gen4087 {
					gen4086 := Call(__e, __e.Global(symcons_2), V1838)

					if True == gen4086 {
						gen4088 = True
					} else {
						gen4088 = False
					}

				} else {
					gen4088 = False
				}
				if True == gen4088 {
					gen4077 := Call(__e, __e.Global(symhd), V1837)

					gen4078 := Call(__e, __e.Global(symhd), V1838)

					gen4079 := Call(__e, __e.Global(symcons), V1840, gen4078)

					gen4080 := Call(__e, __e.Global(symcons), V1839, gen4079)

					gen4081 := Call(__e, __e.Global(symcons), gen4077, gen4080)

					gen4082 := Call(__e, __e.Global(symtl), V1837)

					gen4083 := Call(__e, __e.Global(symtl), V1838)

					gen4084 := Call(__e, __e.Global(symgensym), MakeSymbol("Context"))

					gen4085 := Call(__e, __e.Global(symshen_4csl_1help), gen4082, gen4083, V1840, gen4084)

					__e.TailApply(__e.Global(symcons), gen4081, gen4085)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.csl-help"))

					return
				}

			}

		}, 4)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.csl-help"), gen4095)

		gen4112 := MakeNative(func(__e Evaluator) {
			V1844 := __e.Get(1)
			_ = V1844
			V1845 := __e.Get(2)
			_ = V1845
			V1846 := __e.Get(3)
			_ = V1846
			gen4110 := Call(__e, __e.Global(sym_a), Nil, V1844)

			var gen4111 Obj
			if True == gen4110 {
				gen4108 := Call(__e, __e.Global(sym_a), Nil, V1845)

				var gen4109 Obj
				if True == gen4108 {
					gen4107 := Call(__e, __e.Global(sym_a), Nil, V1846)

					if True == gen4107 {
						gen4109 = True
					} else {
						gen4109 = False
					}

				} else {
					gen4109 = False
				}
				if True == gen4109 {
					gen4111 = True
				} else {
					gen4111 = False
				}

			} else {
				gen4111 = False
			}
			if True == gen4111 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen4105 := Call(__e, __e.Global(symcons_2), V1844)

				var gen4106 Obj
				if True == gen4105 {
					gen4103 := Call(__e, __e.Global(symcons_2), V1845)

					var gen4104 Obj
					if True == gen4103 {
						gen4102 := Call(__e, __e.Global(symcons_2), V1846)

						if True == gen4102 {
							gen4104 = True
						} else {
							gen4104 = False
						}

					} else {
						gen4104 = False
					}
					if True == gen4104 {
						gen4106 = True
					} else {
						gen4106 = False
					}

				} else {
					gen4106 = False
				}
				if True == gen4106 {
					gen4096 := Call(__e, __e.Global(symhd), V1844)

					gen4097 := Call(__e, __e.Global(symhd), V1845)

					gen4098 := Call(__e, __e.Global(symhd), V1846)

					Call(__e, __e.Global(symshen_4construct_1search_1clause), gen4096, gen4097, gen4098)

					gen4099 := Call(__e, __e.Global(symtl), V1844)

					gen4100 := Call(__e, __e.Global(symtl), V1845)

					gen4101 := Call(__e, __e.Global(symtl), V1846)

					__e.TailApply(__e.Global(symshen_4construct_1search_1clauses), gen4099, gen4100, gen4101)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.construct-search-clauses"))

					return
				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-search-clauses"), gen4112)

		gen4117 := MakeNative(func(__e Evaluator) {
			V1850 := __e.Get(1)
			_ = V1850
			V1851 := __e.Get(2)
			_ = V1851
			V1852 := __e.Get(3)
			_ = V1852
			gen4113 := Call(__e, __e.Global(symshen_4construct_1base_1search_1clause), V1850, V1851, V1852)

			gen4114 := Call(__e, __e.Global(symshen_4construct_1recursive_1search_1clause), V1850, V1851, V1852)

			gen4115 := Call(__e, __e.Global(symcons), gen4114, Nil)

			gen4116 := Call(__e, __e.Global(symcons), gen4113, gen4115)

			__e.TailApply(__e.Global(symshen_4s_1prolog), gen4116)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-search-clause"), gen4117)

		gen4125 := MakeNative(func(__e Evaluator) {
			V1856 := __e.Get(1)
			_ = V1856
			V1857 := __e.Get(2)
			_ = V1857
			V1858 := __e.Get(3)
			_ = V1858
			gen4118 := Call(__e, __e.Global(symshen_4mode_1ify), V1857)

			gen4119 := Call(__e, __e.Global(symcons), gen4118, MakeSymbol("In_1957"))

			gen4120 := Call(__e, __e.Global(symcons), MakeSymbol("In_1957"), V1858)

			gen4121 := Call(__e, __e.Global(symcons), gen4119, gen4120)

			gen4122 := Call(__e, __e.Global(symcons), V1856, gen4121)

			gen4123 := Call(__e, __e.Global(symcons), Nil, Nil)

			gen4124 := Call(__e, __e.Global(symcons), MakeSymbol(":-"), gen4123)

			__e.TailApply(__e.Global(symcons), gen4122, gen4124)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-base-search-clause"), gen4125)

		gen4137 := MakeNative(func(__e Evaluator) {
			V1862 := __e.Get(1)
			_ = V1862
			V1863 := __e.Get(2)
			_ = V1863
			V1864 := __e.Get(3)
			_ = V1864
			gen4126 := Call(__e, __e.Global(symcons), MakeSymbol("Assumption_1957"), MakeSymbol("Assumptions_1957"))

			gen4127 := Call(__e, __e.Global(symcons), MakeSymbol("Assumption_1957"), MakeSymbol("Out_1957"))

			gen4128 := Call(__e, __e.Global(symcons), gen4127, V1864)

			gen4129 := Call(__e, __e.Global(symcons), gen4126, gen4128)

			gen4130 := Call(__e, __e.Global(symcons), V1862, gen4129)

			gen4131 := Call(__e, __e.Global(symcons), MakeSymbol("Out_1957"), V1864)

			gen4132 := Call(__e, __e.Global(symcons), MakeSymbol("Assumptions_1957"), gen4131)

			gen4133 := Call(__e, __e.Global(symcons), V1862, gen4132)

			gen4134 := Call(__e, __e.Global(symcons), gen4133, Nil)

			gen4135 := Call(__e, __e.Global(symcons), gen4134, Nil)

			gen4136 := Call(__e, __e.Global(symcons), MakeSymbol(":-"), gen4135)

			__e.TailApply(__e.Global(symcons), gen4130, gen4136)

			return

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-recursive-search-clause"), gen4137)

		gen4191 := MakeNative(func(__e Evaluator) {
			V1870 := __e.Get(1)
			_ = V1870
			gen4190 := Call(__e, __e.Global(sym_a), Nil, V1870)

			if True == gen4190 {
				__e.Return(Nil)
				return
			} else {
				gen4188 := Call(__e, __e.Global(symcons_2), V1870)

				var gen4189 Obj
				if True == gen4188 {
					gen4185 := Call(__e, __e.Global(symhd), V1870)

					gen4186 := Call(__e, __e.Global(symcons_2), gen4185)

					var gen4187 Obj
					if True == gen4186 {
						gen4181 := Call(__e, __e.Global(symhd), V1870)

						gen4182 := Call(__e, __e.Global(symhd), gen4181)

						gen4183 := Call(__e, __e.Global(sym_a), MakeSymbol("if"), gen4182)

						var gen4184 Obj
						if True == gen4183 {
							gen4177 := Call(__e, __e.Global(symhd), V1870)

							gen4178 := Call(__e, __e.Global(symtl), gen4177)

							gen4179 := Call(__e, __e.Global(symcons_2), gen4178)

							var gen4180 Obj
							if True == gen4179 {
								gen4173 := Call(__e, __e.Global(symhd), V1870)

								gen4174 := Call(__e, __e.Global(symtl), gen4173)

								gen4175 := Call(__e, __e.Global(symtl), gen4174)

								gen4176 := Call(__e, __e.Global(sym_a), Nil, gen4175)

								if True == gen4176 {
									gen4180 = True
								} else {
									gen4180 = False
								}

							} else {
								gen4180 = False
							}
							if True == gen4180 {
								gen4184 = True
							} else {
								gen4184 = False
							}

						} else {
							gen4184 = False
						}
						if True == gen4184 {
							gen4187 = True
						} else {
							gen4187 = False
						}

					} else {
						gen4187 = False
					}
					if True == gen4187 {
						gen4189 = True
					} else {
						gen4189 = False
					}

				} else {
					gen4189 = False
				}
				if True == gen4189 {
					gen4168 := Call(__e, __e.Global(symhd), V1870)

					gen4169 := Call(__e, __e.Global(symtl), gen4168)

					gen4170 := Call(__e, __e.Global(symcons), MakeSymbol("when"), gen4169)

					gen4171 := Call(__e, __e.Global(symtl), V1870)

					gen4172 := Call(__e, __e.Global(symshen_4construct_1side_1literals), gen4171)

					__e.TailApply(__e.Global(symcons), gen4170, gen4172)

					return

				} else {
					gen4166 := Call(__e, __e.Global(symcons_2), V1870)

					var gen4167 Obj
					if True == gen4166 {
						gen4163 := Call(__e, __e.Global(symhd), V1870)

						gen4164 := Call(__e, __e.Global(symcons_2), gen4163)

						var gen4165 Obj
						if True == gen4164 {
							gen4159 := Call(__e, __e.Global(symhd), V1870)

							gen4160 := Call(__e, __e.Global(symhd), gen4159)

							gen4161 := Call(__e, __e.Global(sym_a), MakeSymbol("let"), gen4160)

							var gen4162 Obj
							if True == gen4161 {
								gen4155 := Call(__e, __e.Global(symhd), V1870)

								gen4156 := Call(__e, __e.Global(symtl), gen4155)

								gen4157 := Call(__e, __e.Global(symcons_2), gen4156)

								var gen4158 Obj
								if True == gen4157 {
									gen4150 := Call(__e, __e.Global(symhd), V1870)

									gen4151 := Call(__e, __e.Global(symtl), gen4150)

									gen4152 := Call(__e, __e.Global(symtl), gen4151)

									gen4153 := Call(__e, __e.Global(symcons_2), gen4152)

									var gen4154 Obj
									if True == gen4153 {
										gen4145 := Call(__e, __e.Global(symhd), V1870)

										gen4146 := Call(__e, __e.Global(symtl), gen4145)

										gen4147 := Call(__e, __e.Global(symtl), gen4146)

										gen4148 := Call(__e, __e.Global(symtl), gen4147)

										gen4149 := Call(__e, __e.Global(sym_a), Nil, gen4148)

										if True == gen4149 {
											gen4154 = True
										} else {
											gen4154 = False
										}

									} else {
										gen4154 = False
									}
									if True == gen4154 {
										gen4158 = True
									} else {
										gen4158 = False
									}

								} else {
									gen4158 = False
								}
								if True == gen4158 {
									gen4162 = True
								} else {
									gen4162 = False
								}

							} else {
								gen4162 = False
							}
							if True == gen4162 {
								gen4165 = True
							} else {
								gen4165 = False
							}

						} else {
							gen4165 = False
						}
						if True == gen4165 {
							gen4167 = True
						} else {
							gen4167 = False
						}

					} else {
						gen4167 = False
					}
					if True == gen4167 {
						gen4140 := Call(__e, __e.Global(symhd), V1870)

						gen4141 := Call(__e, __e.Global(symtl), gen4140)

						gen4142 := Call(__e, __e.Global(symcons), MakeSymbol("is"), gen4141)

						gen4143 := Call(__e, __e.Global(symtl), V1870)

						gen4144 := Call(__e, __e.Global(symshen_4construct_1side_1literals), gen4143)

						__e.TailApply(__e.Global(symcons), gen4142, gen4144)

						return

					} else {
						gen4139 := Call(__e, __e.Global(symcons_2), V1870)

						if True == gen4139 {
							gen4138 := Call(__e, __e.Global(symtl), V1870)

							__e.TailApply(__e.Global(symshen_4construct_1side_1literals), gen4138)

							return

						} else {
							__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.construct-side-literals"))

							return
						}

					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-side-literals"), gen4191)

		gen4201 := MakeNative(func(__e Evaluator) {
			V1877 := __e.Get(1)
			_ = V1877
			V1878 := __e.Get(2)
			_ = V1878
			gen4200 := Call(__e, __e.Global(symtuple_2), V1877)

			if True == gen4200 {
				gen4194 := Call(__e, __e.Global(symsnd), V1877)

				gen4195 := Call(__e, __e.Global(symshen_4recursive__cons__form), gen4194)

				gen4196 := Call(__e, __e.Global(symfst), V1877)

				gen4197 := Call(__e, __e.Global(symshen_4construct_1context), V1878, gen4196)

				gen4198 := Call(__e, __e.Global(symcons), gen4197, Nil)

				gen4199 := Call(__e, __e.Global(symcons), gen4195, gen4198)

				__e.TailApply(__e.Global(symcons), MakeSymbol("shen.t*"), gen4199)

				return

			} else {
				gen4193 := Call(__e, __e.Global(sym_a), MakeSymbol("!"), V1877)

				if True == gen4193 {
					gen4192 := Call(__e, __e.Global(symcons), MakeSymbol("Throwcontrol"), Nil)

					__e.TailApply(__e.Global(symcons), MakeSymbol("cut"), gen4192)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.construct-premiss-literal"))

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-premiss-literal"), gen4201)

		gen4215 := MakeNative(func(__e Evaluator) {
			V1881 := __e.Get(1)
			_ = V1881
			V1882 := __e.Get(2)
			_ = V1882
			gen4213 := Call(__e, __e.Global(sym_a), True, V1881)

			var gen4214 Obj
			if True == gen4213 {
				gen4212 := Call(__e, __e.Global(sym_a), Nil, V1882)

				if True == gen4212 {
					gen4214 = True
				} else {
					gen4214 = False
				}

			} else {
				gen4214 = False
			}
			if True == gen4214 {
				__e.Return(MakeSymbol("Context_1957"))
				return
			} else {
				gen4210 := Call(__e, __e.Global(sym_a), False, V1881)

				var gen4211 Obj
				if True == gen4210 {
					gen4209 := Call(__e, __e.Global(sym_a), Nil, V1882)

					if True == gen4209 {
						gen4211 = True
					} else {
						gen4211 = False
					}

				} else {
					gen4211 = False
				}
				if True == gen4211 {
					__e.Return(MakeSymbol("ContextOut_1957"))
					return
				} else {
					gen4208 := Call(__e, __e.Global(symcons_2), V1882)

					if True == gen4208 {
						gen4202 := Call(__e, __e.Global(symhd), V1882)

						gen4203 := Call(__e, __e.Global(symshen_4recursive__cons__form), gen4202)

						gen4204 := Call(__e, __e.Global(symtl), V1882)

						gen4205 := Call(__e, __e.Global(symshen_4construct_1context), V1881, gen4204)

						gen4206 := Call(__e, __e.Global(symcons), gen4205, Nil)

						gen4207 := Call(__e, __e.Global(symcons), gen4203, gen4206)

						__e.TailApply(__e.Global(symcons), MakeSymbol("cons"), gen4207)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.construct-context"))

						return
					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.construct-context"), gen4215)

		gen4223 := MakeNative(func(__e Evaluator) {
			V1884 := __e.Get(1)
			_ = V1884
			gen4222 := Call(__e, __e.Global(symcons_2), V1884)

			if True == gen4222 {
				gen4216 := Call(__e, __e.Global(symhd), V1884)

				gen4217 := Call(__e, __e.Global(symshen_4recursive__cons__form), gen4216)

				gen4218 := Call(__e, __e.Global(symtl), V1884)

				gen4219 := Call(__e, __e.Global(symshen_4recursive__cons__form), gen4218)

				gen4220 := Call(__e, __e.Global(symcons), gen4219, Nil)

				gen4221 := Call(__e, __e.Global(symcons), gen4217, gen4220)

				__e.TailApply(__e.Global(symcons), MakeSymbol("cons"), gen4221)

				return

			} else {
				__e.Return(V1884)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.recursive_cons_form"), gen4223)

		gen4226 := MakeNative(func(__e Evaluator) {
			V1886 := __e.Get(1)
			_ = V1886
			gen4224 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4intern_1type), X)

				return
			}, 1)
			gen4225 := Call(__e, __e.Global(symmap), gen4224, V1886)

			__e.TailApply(__e.Global(symshen_4preclude_1h), gen4225)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("preclude"), gen4226)

		gen4230 := MakeNative(func(__e Evaluator) {
			V1888 := __e.Get(1)
			_ = V1888
			gen4227 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*datatypes*"))

			gen4228 := Call(__e, __e.Global(symdifference), gen4227, V1888)

			gen4229 := Call(__e, __e.Global(symset), MakeSymbol("shen.*datatypes*"), gen4228)

			FilterDatatypes := gen4229
			_ = FilterDatatypes
			__e.TailApply(__e.Global(symvalue), MakeSymbol("shen.*datatypes*"))

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.preclude-h"), gen4230)

		gen4233 := MakeNative(func(__e Evaluator) {
			V1890 := __e.Get(1)
			_ = V1890
			gen4231 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4intern_1type), X)

				return
			}, 1)
			gen4232 := Call(__e, __e.Global(symmap), gen4231, V1890)

			__e.TailApply(__e.Global(symshen_4include_1h), gen4232)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("include"), gen4233)

		gen4239 := MakeNative(func(__e Evaluator) {
			V1892 := __e.Get(1)
			_ = V1892
			gen4234 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen4235 := Call(__e, __e.Global(symintersection), V1892, gen4234)

			ValidTypes := gen4235
			gen4236 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*datatypes*"))

			gen4237 := Call(__e, __e.Global(symunion), ValidTypes, gen4236)

			gen4238 := Call(__e, __e.Global(symset), MakeSymbol("shen.*datatypes*"), gen4237)

			NewDatatypes := gen4238
			_ = NewDatatypes
			__e.TailApply(__e.Global(symvalue), MakeSymbol("shen.*datatypes*"))

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.include-h"), gen4239)

		gen4244 := MakeNative(func(__e Evaluator) {
			V1894 := __e.Get(1)
			_ = V1894
			gen4240 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen4241 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4intern_1type), X)

				return
			}, 1)
			gen4242 := Call(__e, __e.Global(symmap), gen4241, V1894)

			gen4243 := Call(__e, __e.Global(symdifference), gen4240, gen4242)

			__e.TailApply(__e.Global(symshen_4preclude_1h), gen4243)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("preclude-all-but"), gen4244)

		gen4249 := MakeNative(func(__e Evaluator) {
			V1896 := __e.Get(1)
			_ = V1896
			gen4245 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen4246 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4intern_1type), X)

				return
			}, 1)
			gen4247 := Call(__e, __e.Global(symmap), gen4246, V1896)

			gen4248 := Call(__e, __e.Global(symdifference), gen4245, gen4247)

			__e.TailApply(__e.Global(symshen_4include_1h), gen4248)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("include-all-but"), gen4249)

		gen4275 := MakeNative(func(__e Evaluator) {
			V1902 := __e.Get(1)
			_ = V1902
			gen4274 := Call(__e, __e.Global(sym_a), Nil, V1902)

			if True == gen4274 {
				gen4270 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*tc*"))

				gen4271 := MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(__e.Global(symshen_4demod_1rule), X)

					return
				}, 1)
				gen4272 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*synonyms*"))

				gen4273 := Call(__e, __e.Global(symmapcan), gen4271, gen4272)

				__e.TailApply(__e.Global(symshen_4update_1demodulation_1function), gen4270, gen4273)

				return

			} else {
				gen4268 := Call(__e, __e.Global(symcons_2), V1902)

				var gen4269 Obj
				if True == gen4268 {
					gen4266 := Call(__e, __e.Global(symtl), V1902)

					gen4267 := Call(__e, __e.Global(symcons_2), gen4266)

					if True == gen4267 {
						gen4269 = True
					} else {
						gen4269 = False
					}

				} else {
					gen4269 = False
				}
				if True == gen4269 {
					gen4250 := Call(__e, __e.Global(symtl), V1902)

					gen4251 := Call(__e, __e.Global(symhd), gen4250)

					gen4252 := Call(__e, __e.Global(symshen_4extract__vars), gen4251)

					gen4253 := Call(__e, __e.Global(symhd), V1902)

					gen4254 := Call(__e, __e.Global(symshen_4extract__vars), gen4253)

					gen4255 := Call(__e, __e.Global(symdifference), gen4252, gen4254)

					Vs := gen4255
					gen4265 := Call(__e, __e.Global(symempty_2), Vs)

					if True == gen4265 {
						gen4258 := Call(__e, __e.Global(symhd), V1902)

						gen4259 := Call(__e, __e.Global(symtl), V1902)

						gen4260 := Call(__e, __e.Global(symhd), gen4259)

						gen4261 := Call(__e, __e.Global(symcons), gen4260, Nil)

						gen4262 := Call(__e, __e.Global(symcons), gen4258, gen4261)

						Call(__e, __e.Global(symshen_4pushnew), gen4262, MakeSymbol("shen.*synonyms*"))

						gen4263 := Call(__e, __e.Global(symtl), V1902)

						gen4264 := Call(__e, __e.Global(symtl), gen4263)

						__e.TailApply(__e.Global(symshen_4synonyms_1help), gen4264)

						return

					} else {
						gen4256 := Call(__e, __e.Global(symtl), V1902)

						gen4257 := Call(__e, __e.Global(symhd), gen4256)

						__e.TailApply(__e.Global(symshen_4free__variable__warnings), gen4257, Vs)

						return

					}

				} else {
					__e.TailApply(__e.Global(symsimple_1error), MakeString("odd number of synonyms\n"))

					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.synonyms-help"), gen4275)

		gen4280 := MakeNative(func(__e Evaluator) {
			V1905 := __e.Get(1)
			_ = V1905
			V1906 := __e.Get(2)
			_ = V1906
			gen4278 := Call(__e, __e.Global(symvalue), V1906)

			gen4279 := Call(__e, __e.Global(symelement_2), V1905, gen4278)

			if True == gen4279 {
				__e.TailApply(__e.Global(symvalue), V1906)

				return
			} else {
				gen4276 := Call(__e, __e.Global(symvalue), V1906)

				gen4277 := Call(__e, __e.Global(symcons), V1905, gen4276)

				__e.TailApply(__e.Global(symset), V1906, gen4277)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.pushnew"), gen4280)

		gen4296 := MakeNative(func(__e Evaluator) {
			V1908 := __e.Get(1)
			_ = V1908
			gen4294 := Call(__e, __e.Global(symcons_2), V1908)

			var gen4295 Obj
			if True == gen4294 {
				gen4291 := Call(__e, __e.Global(symtl), V1908)

				gen4292 := Call(__e, __e.Global(symcons_2), gen4291)

				var gen4293 Obj
				if True == gen4292 {
					gen4288 := Call(__e, __e.Global(symtl), V1908)

					gen4289 := Call(__e, __e.Global(symtl), gen4288)

					gen4290 := Call(__e, __e.Global(sym_a), Nil, gen4289)

					if True == gen4290 {
						gen4293 = True
					} else {
						gen4293 = False
					}

				} else {
					gen4293 = False
				}
				if True == gen4293 {
					gen4295 = True
				} else {
					gen4295 = False
				}

			} else {
				gen4295 = False
			}
			if True == gen4295 {
				gen4281 := Call(__e, __e.Global(symhd), V1908)

				gen4282 := Call(__e, __e.Global(symshen_4rcons__form), gen4281)

				gen4283 := Call(__e, __e.Global(symtl), V1908)

				gen4284 := Call(__e, __e.Global(symhd), gen4283)

				gen4285 := Call(__e, __e.Global(symshen_4rcons__form), gen4284)

				gen4286 := Call(__e, __e.Global(symcons), gen4285, Nil)

				gen4287 := Call(__e, __e.Global(symcons), MakeSymbol("->"), gen4286)

				__e.TailApply(__e.Global(symcons), gen4282, gen4287)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.demod-rule"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.demod-rule"), gen4296)

		gen4339 := MakeNative(func(__e Evaluator) {
			V1914 := __e.Get(1)
			_ = V1914
			gen4337 := Call(__e, __e.Global(symcons_2), V1914)

			var gen4338 Obj
			if True == gen4337 {
				gen4334 := Call(__e, __e.Global(symhd), V1914)

				gen4335 := Call(__e, __e.Global(sym_a), MakeSymbol("defun"), gen4334)

				var gen4336 Obj
				if True == gen4335 {
					gen4331 := Call(__e, __e.Global(symtl), V1914)

					gen4332 := Call(__e, __e.Global(symcons_2), gen4331)

					var gen4333 Obj
					if True == gen4332 {
						gen4327 := Call(__e, __e.Global(symtl), V1914)

						gen4328 := Call(__e, __e.Global(symtl), gen4327)

						gen4329 := Call(__e, __e.Global(symcons_2), gen4328)

						var gen4330 Obj
						if True == gen4329 {
							gen4322 := Call(__e, __e.Global(symtl), V1914)

							gen4323 := Call(__e, __e.Global(symtl), gen4322)

							gen4324 := Call(__e, __e.Global(symhd), gen4323)

							gen4325 := Call(__e, __e.Global(symcons_2), gen4324)

							var gen4326 Obj
							if True == gen4325 {
								gen4316 := Call(__e, __e.Global(symtl), V1914)

								gen4317 := Call(__e, __e.Global(symtl), gen4316)

								gen4318 := Call(__e, __e.Global(symhd), gen4317)

								gen4319 := Call(__e, __e.Global(symtl), gen4318)

								gen4320 := Call(__e, __e.Global(sym_a), Nil, gen4319)

								var gen4321 Obj
								if True == gen4320 {
									gen4311 := Call(__e, __e.Global(symtl), V1914)

									gen4312 := Call(__e, __e.Global(symtl), gen4311)

									gen4313 := Call(__e, __e.Global(symtl), gen4312)

									gen4314 := Call(__e, __e.Global(symcons_2), gen4313)

									var gen4315 Obj
									if True == gen4314 {
										gen4306 := Call(__e, __e.Global(symtl), V1914)

										gen4307 := Call(__e, __e.Global(symtl), gen4306)

										gen4308 := Call(__e, __e.Global(symtl), gen4307)

										gen4309 := Call(__e, __e.Global(symtl), gen4308)

										gen4310 := Call(__e, __e.Global(sym_a), Nil, gen4309)

										if True == gen4310 {
											gen4315 = True
										} else {
											gen4315 = False
										}

									} else {
										gen4315 = False
									}
									if True == gen4315 {
										gen4321 = True
									} else {
										gen4321 = False
									}

								} else {
									gen4321 = False
								}
								if True == gen4321 {
									gen4326 = True
								} else {
									gen4326 = False
								}

							} else {
								gen4326 = False
							}
							if True == gen4326 {
								gen4330 = True
							} else {
								gen4330 = False
							}

						} else {
							gen4330 = False
						}
						if True == gen4330 {
							gen4333 = True
						} else {
							gen4333 = False
						}

					} else {
						gen4333 = False
					}
					if True == gen4333 {
						gen4336 = True
					} else {
						gen4336 = False
					}

				} else {
					gen4336 = False
				}
				if True == gen4336 {
					gen4338 = True
				} else {
					gen4338 = False
				}

			} else {
				gen4338 = False
			}
			if True == gen4338 {
				gen4297 := Call(__e, __e.Global(symtl), V1914)

				gen4298 := Call(__e, __e.Global(symtl), gen4297)

				gen4299 := Call(__e, __e.Global(symhd), gen4298)

				gen4300 := Call(__e, __e.Global(symhd), gen4299)

				gen4301 := Call(__e, __e.Global(symtl), V1914)

				gen4302 := Call(__e, __e.Global(symtl), gen4301)

				gen4303 := Call(__e, __e.Global(symtl), gen4302)

				gen4304 := Call(__e, __e.Global(symcons), gen4300, gen4303)

				gen4305 := Call(__e, __e.Global(symcons), MakeSymbol("/."), gen4304)

				__e.TailApply(__e.Global(symeval), gen4305)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.lambda-of-defun"))

				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.lambda-of-defun"), gen4339)

		gen4346 := MakeNative(func(__e Evaluator) {
			V1917 := __e.Get(1)
			_ = V1917
			V1918 := __e.Get(2)
			_ = V1918
			Call(__e, __e.Global(symtc), MakeSymbol("-"))
			gen4340 := Call(__e, __e.Global(symshen_4default_1rule))

			gen4341 := Call(__e, __e.Global(symappend), V1918, gen4340)

			gen4342 := Call(__e, __e.Global(symcons), MakeSymbol("shen.demod"), gen4341)

			gen4343 := Call(__e, __e.Global(symcons), MakeSymbol("define"), gen4342)

			gen4344 := Call(__e, __e.Global(symshen_4elim_1def), gen4343)

			gen4345 := Call(__e, __e.Global(symshen_4lambda_1of_1defun), gen4344)

			Call(__e, __e.Global(symset), MakeSymbol("shen.*demodulation-function*"), gen4345)

			if True == V1917 {
				Call(__e, __e.Global(symtc), MakeSymbol("+"))
			} else {
				MakeSymbol("shen.skip")
			}
			__e.Return(MakeSymbol("synonyms"))
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.update-demodulation-function"), gen4346)

		gen4349 := MakeNative(func(__e Evaluator) {
			gen4347 := Call(__e, __e.Global(symcons), MakeSymbol("X"), Nil)

			gen4348 := Call(__e, __e.Global(symcons), MakeSymbol("->"), gen4347)

			__e.TailApply(__e.Global(symcons), MakeSymbol("X"), gen4348)

			return

		}, 0)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.default-rule"), gen4349)

		return

	}, 0))
}
