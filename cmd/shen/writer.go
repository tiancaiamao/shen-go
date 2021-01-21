package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp104476 := MakeNative(func(__e *ControlFlow) {
		V4594 := __e.Get(1)
		_ = V4594
		V4595 := __e.Get(2)
		_ = V4595
		tmp104477 := MakeNative(func(__e *ControlFlow) {
			tmp104478 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prh)

			__e.TailApply(tmp104478, V4594, V4595, MakeNumber(0))
			return

		}, 0)

		tmp104479 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4594)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp104477, tmp104479)
		return

	}, 2)

	tmp104480 := Call(__e, PrimNS1Value(symns2_1set), sympr, tmp104476)

	_ = tmp104480

	tmp104481 := MakeNative(func(__e *ControlFlow) {
		V4599 := __e.Get(1)
		_ = V4599
		V4600 := __e.Get(2)
		_ = V4600
		V4601 := __e.Get(3)
		_ = V4601
		tmp104482 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prh)

		tmp104483 := Call(__e, PrimNS1Value(symns2_1value), symshen_4write_1char_1and_1inc)

		tmp104484 := Call(__e, tmp104483, V4599, V4600, V4601)

		__e.TailApply(tmp104482, V4599, V4600, tmp104484)
		return

	}, 3)

	tmp104485 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prh, tmp104481)

	_ = tmp104485

	tmp104486 := MakeNative(func(__e *ControlFlow) {
		V4605 := __e.Get(1)
		_ = V4605
		V4606 := __e.Get(2)
		_ = V4606
		V4607 := __e.Get(3)
		_ = V4607
		tmp104487 := Call(__e, PrimNS1Value(symns2_1value), symwrite_1byte)

		tmp104488 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

		tmp104489 := Call(__e, PrimNS1Value(symns2_1value), sympos)

		tmp104490 := Call(__e, tmp104489, V4605, V4607)

		tmp104491 := Call(__e, tmp104488, tmp104490)

		tmp104492 := Call(__e, tmp104487, tmp104491, V4606)

		_ = tmp104492

		tmp104493 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		__e.TailApply(tmp104493, V4607, MakeNumber(1))
		return

	}, 3)

	tmp104494 := Call(__e, PrimNS1Value(symns2_1set), symshen_4write_1char_1and_1inc, tmp104486)

	_ = tmp104494

	tmp104495 := MakeNative(func(__e *ControlFlow) {
		V4609 := __e.Get(1)
		_ = V4609
		tmp104496 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp104497 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				__e.Return(V4609)
				return
			}, 1)

			tmp104498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp104499 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp104500 := Call(__e, tmp104499)

			tmp104501 := Call(__e, tmp104498, String, tmp104500)

			__e.TailApply(tmp104497, tmp104501)
			return

		}, 1)

		tmp104502 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert)

		tmp104503 := Call(__e, tmp104502, V4609, MakeString("~S"))

		__e.TailApply(tmp104496, tmp104503)
		return

	}, 1)

	tmp104504 := Call(__e, PrimNS1Value(symns2_1set), symprint, tmp104495)

	_ = tmp104504

	tmp104505 := MakeNative(func(__e *ControlFlow) {
		V4612 := __e.Get(1)
		_ = V4612
		V4613 := __e.Get(2)
		_ = V4613
		tmp104508 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp104509 := Call(__e, tmp104508, sym_dhush_d)

		if True == tmp104509 {
			__e.Return(V4612)
			return
		} else {
			tmp104507 := Call(__e, PrimNS1Value(symns2_1value), sympr)

			__e.TailApply(tmp104507, V4612, V4613)
			return

		}

	}, 2)

	tmp104510 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prhush, tmp104505)

	_ = tmp104510

	tmp104511 := MakeNative(func(__e *ControlFlow) {
		V4616 := __e.Get(1)
		_ = V4616
		V4617 := __e.Get(2)
		_ = V4617
		tmp104521 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

		tmp104522 := Call(__e, tmp104521, V4616)

		if True == tmp104522 {
			tmp104518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1l)

			tmp104519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

			tmp104520 := Call(__e, tmp104519, V4616)

			__e.TailApply(tmp104518, tmp104520, V4617)
			return

		} else {
			tmp104513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1r)

			tmp104514 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104515 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104516 := Call(__e, tmp104515, V4616, Nil)

			tmp104517 := Call(__e, tmp104514, symshen_4proc_1nl, tmp104516)

			__e.TailApply(tmp104513, tmp104517, V4617)
			return

		}

	}, 2)

	tmp104523 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr, tmp104511)

	_ = tmp104523

	tmp104524 := MakeNative(func(__e *ControlFlow) {
		V4620 := __e.Get(1)
		_ = V4620
		V4621 := __e.Get(2)
		_ = V4621
		tmp104537 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp104538 := Call(__e, tmp104537, Nil, V4621)

		if True == tmp104538 {
			__e.Return(V4620)
			return
		} else {
			tmp104535 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp104536 := Call(__e, tmp104535, V4621)

			if True == tmp104536 {
				tmp104528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1l)

				tmp104529 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

				tmp104530 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp104531 := Call(__e, tmp104530, V4621)

				tmp104532 := Call(__e, tmp104529, tmp104531, V4620)

				tmp104533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp104534 := Call(__e, tmp104533, V4621)

				__e.TailApply(tmp104528, tmp104532, tmp104534)
				return

			} else {
				tmp104527 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp104527, symshen_4mkstr_1l)
				return

			}

		}

	}, 2)

	tmp104539 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1l, tmp104524)

	_ = tmp104539

	tmp104540 := MakeNative(func(__e *ControlFlow) {
		V4626 := __e.Get(1)
		_ = V4626
		V4627 := __e.Get(2)
		_ = V4627
		tmp104779 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp104780 := Call(__e, tmp104779, MakeString(""), V4627)

		if True == tmp104780 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp104777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp104778 := Call(__e, tmp104777, V4627)

			var ifres104757 Obj

			if True == tmp104778 {
				tmp104773 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp104774 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp104775 := Call(__e, tmp104774, V4627, MakeNumber(0))

				tmp104776 := Call(__e, tmp104773, MakeString("~"), tmp104775)

				var ifres104759 Obj

				if True == tmp104776 {
					tmp104769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					tmp104770 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp104771 := Call(__e, tmp104770, V4627)

					tmp104772 := Call(__e, tmp104769, tmp104771)

					var ifres104761 Obj

					if True == tmp104772 {
						tmp104763 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp104764 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						tmp104765 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp104766 := Call(__e, tmp104765, V4627)

						tmp104767 := Call(__e, tmp104764, tmp104766, MakeNumber(0))

						tmp104768 := Call(__e, tmp104763, MakeString("A"), tmp104767)

						var ifres104762 Obj

						if True == tmp104768 {
							ifres104762 = True

						} else {
							ifres104762 = False

						}

						ifres104761 = ifres104762

					} else {
						ifres104761 = False

					}

					var ifres104760 Obj

					if True == ifres104761 {
						ifres104760 = True

					} else {
						ifres104760 = False

					}

					ifres104759 = ifres104760

				} else {
					ifres104759 = False

				}

				var ifres104758 Obj

				if True == ifres104759 {
					ifres104758 = True

				} else {
					ifres104758 = False

				}

				ifres104757 = ifres104758

			} else {
				ifres104757 = False

			}

			if True == ifres104757 {
				tmp104746 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104747 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104748 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104749 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp104750 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp104751 := Call(__e, tmp104750, V4627)

				tmp104752 := Call(__e, tmp104749, tmp104751)

				tmp104753 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104754 := Call(__e, tmp104753, symshen_4a, Nil)

				tmp104755 := Call(__e, tmp104748, tmp104752, tmp104754)

				tmp104756 := Call(__e, tmp104747, V4626, tmp104755)

				__e.TailApply(tmp104746, symshen_4app, tmp104756)
				return

			} else {
				tmp104744 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				tmp104745 := Call(__e, tmp104744, V4627)

				var ifres104724 Obj

				if True == tmp104745 {
					tmp104740 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp104741 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					tmp104742 := Call(__e, tmp104741, V4627, MakeNumber(0))

					tmp104743 := Call(__e, tmp104740, MakeString("~"), tmp104742)

					var ifres104726 Obj

					if True == tmp104743 {
						tmp104736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						tmp104737 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp104738 := Call(__e, tmp104737, V4627)

						tmp104739 := Call(__e, tmp104736, tmp104738)

						var ifres104728 Obj

						if True == tmp104739 {
							tmp104730 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp104731 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							tmp104732 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp104733 := Call(__e, tmp104732, V4627)

							tmp104734 := Call(__e, tmp104731, tmp104733, MakeNumber(0))

							tmp104735 := Call(__e, tmp104730, MakeString("R"), tmp104734)

							var ifres104729 Obj

							if True == tmp104735 {
								ifres104729 = True

							} else {
								ifres104729 = False

							}

							ifres104728 = ifres104729

						} else {
							ifres104728 = False

						}

						var ifres104727 Obj

						if True == ifres104728 {
							ifres104727 = True

						} else {
							ifres104727 = False

						}

						ifres104726 = ifres104727

					} else {
						ifres104726 = False

					}

					var ifres104725 Obj

					if True == ifres104726 {
						ifres104725 = True

					} else {
						ifres104725 = False

					}

					ifres104724 = ifres104725

				} else {
					ifres104724 = False

				}

				if True == ifres104724 {
					tmp104713 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp104714 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp104715 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp104716 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp104717 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp104718 := Call(__e, tmp104717, V4627)

					tmp104719 := Call(__e, tmp104716, tmp104718)

					tmp104720 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp104721 := Call(__e, tmp104720, symshen_4r, Nil)

					tmp104722 := Call(__e, tmp104715, tmp104719, tmp104721)

					tmp104723 := Call(__e, tmp104714, V4626, tmp104722)

					__e.TailApply(tmp104713, symshen_4app, tmp104723)
					return

				} else {
					tmp104711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					tmp104712 := Call(__e, tmp104711, V4627)

					var ifres104691 Obj

					if True == tmp104712 {
						tmp104707 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp104708 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						tmp104709 := Call(__e, tmp104708, V4627, MakeNumber(0))

						tmp104710 := Call(__e, tmp104707, MakeString("~"), tmp104709)

						var ifres104693 Obj

						if True == tmp104710 {
							tmp104703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

							tmp104704 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp104705 := Call(__e, tmp104704, V4627)

							tmp104706 := Call(__e, tmp104703, tmp104705)

							var ifres104695 Obj

							if True == tmp104706 {
								tmp104697 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp104698 := Call(__e, PrimNS1Value(symns2_1value), sympos)

								tmp104699 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

								tmp104700 := Call(__e, tmp104699, V4627)

								tmp104701 := Call(__e, tmp104698, tmp104700, MakeNumber(0))

								tmp104702 := Call(__e, tmp104697, MakeString("S"), tmp104701)

								var ifres104696 Obj

								if True == tmp104702 {
									ifres104696 = True

								} else {
									ifres104696 = False

								}

								ifres104695 = ifres104696

							} else {
								ifres104695 = False

							}

							var ifres104694 Obj

							if True == ifres104695 {
								ifres104694 = True

							} else {
								ifres104694 = False

							}

							ifres104693 = ifres104694

						} else {
							ifres104693 = False

						}

						var ifres104692 Obj

						if True == ifres104693 {
							ifres104692 = True

						} else {
							ifres104692 = False

						}

						ifres104691 = ifres104692

					} else {
						ifres104691 = False

					}

					if True == ifres104691 {
						tmp104680 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp104681 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp104682 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp104683 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp104684 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp104685 := Call(__e, tmp104684, V4627)

						tmp104686 := Call(__e, tmp104683, tmp104685)

						tmp104687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp104688 := Call(__e, tmp104687, symshen_4s, Nil)

						tmp104689 := Call(__e, tmp104682, tmp104686, tmp104688)

						tmp104690 := Call(__e, tmp104681, V4626, tmp104689)

						__e.TailApply(tmp104680, symshen_4app, tmp104690)
						return

					} else {
						tmp104678 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						tmp104679 := Call(__e, tmp104678, V4627)

						if True == tmp104679 {
							tmp104665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4factor_1cn)

							tmp104666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp104667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp104668 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							tmp104669 := Call(__e, tmp104668, V4627, MakeNumber(0))

							tmp104670 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp104671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

							tmp104672 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp104673 := Call(__e, tmp104672, V4627)

							tmp104674 := Call(__e, tmp104671, V4626, tmp104673)

							tmp104675 := Call(__e, tmp104670, tmp104674, Nil)

							tmp104676 := Call(__e, tmp104667, tmp104669, tmp104675)

							tmp104677 := Call(__e, tmp104666, symcn, tmp104676)

							__e.TailApply(tmp104665, tmp104677)
							return

						} else {
							tmp104663 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp104664 := Call(__e, tmp104663, V4627)

							var ifres104633 Obj

							if True == tmp104664 {
								tmp104659 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp104660 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp104661 := Call(__e, tmp104660, V4627)

								tmp104662 := Call(__e, tmp104659, symcn, tmp104661)

								var ifres104635 Obj

								if True == tmp104662 {
									tmp104655 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp104656 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104657 := Call(__e, tmp104656, V4627)

									tmp104658 := Call(__e, tmp104655, tmp104657)

									var ifres104637 Obj

									if True == tmp104658 {
										tmp104649 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp104650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104651 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104652 := Call(__e, tmp104651, V4627)

										tmp104653 := Call(__e, tmp104650, tmp104652)

										tmp104654 := Call(__e, tmp104649, tmp104653)

										var ifres104639 Obj

										if True == tmp104654 {
											tmp104641 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp104642 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104643 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104644 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104645 := Call(__e, tmp104644, V4627)

											tmp104646 := Call(__e, tmp104643, tmp104645)

											tmp104647 := Call(__e, tmp104642, tmp104646)

											tmp104648 := Call(__e, tmp104641, Nil, tmp104647)

											var ifres104640 Obj

											if True == tmp104648 {
												ifres104640 = True

											} else {
												ifres104640 = False

											}

											ifres104639 = ifres104640

										} else {
											ifres104639 = False

										}

										var ifres104638 Obj

										if True == ifres104639 {
											ifres104638 = True

										} else {
											ifres104638 = False

										}

										ifres104637 = ifres104638

									} else {
										ifres104637 = False

									}

									var ifres104636 Obj

									if True == ifres104637 {
										ifres104636 = True

									} else {
										ifres104636 = False

									}

									ifres104635 = ifres104636

								} else {
									ifres104635 = False

								}

								var ifres104634 Obj

								if True == ifres104635 {
									ifres104634 = True

								} else {
									ifres104634 = False

								}

								ifres104633 = ifres104634

							} else {
								ifres104633 = False

							}

							if True == ifres104633 {
								tmp104616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp104617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp104618 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp104619 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104620 := Call(__e, tmp104619, V4627)

								tmp104621 := Call(__e, tmp104618, tmp104620)

								tmp104622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp104623 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

								tmp104624 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp104625 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104626 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104627 := Call(__e, tmp104626, V4627)

								tmp104628 := Call(__e, tmp104625, tmp104627)

								tmp104629 := Call(__e, tmp104624, tmp104628)

								tmp104630 := Call(__e, tmp104623, V4626, tmp104629)

								tmp104631 := Call(__e, tmp104622, tmp104630, Nil)

								tmp104632 := Call(__e, tmp104617, tmp104621, tmp104631)

								__e.TailApply(tmp104616, symcn, tmp104632)
								return

							} else {
								tmp104614 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp104615 := Call(__e, tmp104614, V4627)

								var ifres104572 Obj

								if True == tmp104615 {
									tmp104610 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp104611 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp104612 := Call(__e, tmp104611, V4627)

									tmp104613 := Call(__e, tmp104610, symshen_4app, tmp104612)

									var ifres104574 Obj

									if True == tmp104613 {
										tmp104606 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp104607 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104608 := Call(__e, tmp104607, V4627)

										tmp104609 := Call(__e, tmp104606, tmp104608)

										var ifres104576 Obj

										if True == tmp104609 {
											tmp104600 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp104601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104603 := Call(__e, tmp104602, V4627)

											tmp104604 := Call(__e, tmp104601, tmp104603)

											tmp104605 := Call(__e, tmp104600, tmp104604)

											var ifres104578 Obj

											if True == tmp104605 {
												tmp104592 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp104593 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp104594 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp104595 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp104596 := Call(__e, tmp104595, V4627)

												tmp104597 := Call(__e, tmp104594, tmp104596)

												tmp104598 := Call(__e, tmp104593, tmp104597)

												tmp104599 := Call(__e, tmp104592, tmp104598)

												var ifres104580 Obj

												if True == tmp104599 {
													tmp104582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp104583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104584 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104585 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104586 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104587 := Call(__e, tmp104586, V4627)

													tmp104588 := Call(__e, tmp104585, tmp104587)

													tmp104589 := Call(__e, tmp104584, tmp104588)

													tmp104590 := Call(__e, tmp104583, tmp104589)

													tmp104591 := Call(__e, tmp104582, Nil, tmp104590)

													var ifres104581 Obj

													if True == tmp104591 {
														ifres104581 = True

													} else {
														ifres104581 = False

													}

													ifres104580 = ifres104581

												} else {
													ifres104580 = False

												}

												var ifres104579 Obj

												if True == ifres104580 {
													ifres104579 = True

												} else {
													ifres104579 = False

												}

												ifres104578 = ifres104579

											} else {
												ifres104578 = False

											}

											var ifres104577 Obj

											if True == ifres104578 {
												ifres104577 = True

											} else {
												ifres104577 = False

											}

											ifres104576 = ifres104577

										} else {
											ifres104576 = False

										}

										var ifres104575 Obj

										if True == ifres104576 {
											ifres104575 = True

										} else {
											ifres104575 = False

										}

										ifres104574 = ifres104575

									} else {
										ifres104574 = False

									}

									var ifres104573 Obj

									if True == ifres104574 {
										ifres104573 = True

									} else {
										ifres104573 = False

									}

									ifres104572 = ifres104573

								} else {
									ifres104572 = False

								}

								if True == ifres104572 {
									tmp104549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp104550 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp104551 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp104552 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104553 := Call(__e, tmp104552, V4627)

									tmp104554 := Call(__e, tmp104551, tmp104553)

									tmp104555 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp104556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1l)

									tmp104557 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp104558 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104559 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104560 := Call(__e, tmp104559, V4627)

									tmp104561 := Call(__e, tmp104558, tmp104560)

									tmp104562 := Call(__e, tmp104557, tmp104561)

									tmp104563 := Call(__e, tmp104556, V4626, tmp104562)

									tmp104564 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104565 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104566 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104567 := Call(__e, tmp104566, V4627)

									tmp104568 := Call(__e, tmp104565, tmp104567)

									tmp104569 := Call(__e, tmp104564, tmp104568)

									tmp104570 := Call(__e, tmp104555, tmp104563, tmp104569)

									tmp104571 := Call(__e, tmp104550, tmp104554, tmp104570)

									__e.TailApply(tmp104549, symshen_4app, tmp104571)
									return

								} else {
									tmp104548 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

									__e.TailApply(tmp104548, symshen_4insert_1l)
									return

								}

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp104781 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1l, tmp104540)

	_ = tmp104781

	tmp104782 := MakeNative(func(__e *ControlFlow) {
		V4629 := __e.Get(1)
		_ = V4629
		tmp104929 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp104930 := Call(__e, tmp104929, V4629)

		var ifres104813 Obj

		if True == tmp104930 {
			tmp104925 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp104926 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp104927 := Call(__e, tmp104926, V4629)

			tmp104928 := Call(__e, tmp104925, symcn, tmp104927)

			var ifres104815 Obj

			if True == tmp104928 {
				tmp104921 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp104922 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp104923 := Call(__e, tmp104922, V4629)

				tmp104924 := Call(__e, tmp104921, tmp104923)

				var ifres104817 Obj

				if True == tmp104924 {
					tmp104915 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp104916 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp104917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp104918 := Call(__e, tmp104917, V4629)

					tmp104919 := Call(__e, tmp104916, tmp104918)

					tmp104920 := Call(__e, tmp104915, tmp104919)

					var ifres104819 Obj

					if True == tmp104920 {
						tmp104907 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp104908 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp104909 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp104910 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp104911 := Call(__e, tmp104910, V4629)

						tmp104912 := Call(__e, tmp104909, tmp104911)

						tmp104913 := Call(__e, tmp104908, tmp104912)

						tmp104914 := Call(__e, tmp104907, tmp104913)

						var ifres104821 Obj

						if True == tmp104914 {
							tmp104897 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp104898 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp104899 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp104900 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp104901 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp104902 := Call(__e, tmp104901, V4629)

							tmp104903 := Call(__e, tmp104900, tmp104902)

							tmp104904 := Call(__e, tmp104899, tmp104903)

							tmp104905 := Call(__e, tmp104898, tmp104904)

							tmp104906 := Call(__e, tmp104897, symcn, tmp104905)

							var ifres104823 Obj

							if True == tmp104906 {
								tmp104887 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp104888 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104889 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp104890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104891 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp104892 := Call(__e, tmp104891, V4629)

								tmp104893 := Call(__e, tmp104890, tmp104892)

								tmp104894 := Call(__e, tmp104889, tmp104893)

								tmp104895 := Call(__e, tmp104888, tmp104894)

								tmp104896 := Call(__e, tmp104887, tmp104895)

								var ifres104825 Obj

								if True == tmp104896 {
									tmp104875 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp104876 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104878 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp104879 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104880 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp104881 := Call(__e, tmp104880, V4629)

									tmp104882 := Call(__e, tmp104879, tmp104881)

									tmp104883 := Call(__e, tmp104878, tmp104882)

									tmp104884 := Call(__e, tmp104877, tmp104883)

									tmp104885 := Call(__e, tmp104876, tmp104884)

									tmp104886 := Call(__e, tmp104875, tmp104885)

									var ifres104827 Obj

									if True == tmp104886 {
										tmp104861 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp104862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104863 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104864 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104865 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp104866 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp104868 := Call(__e, tmp104867, V4629)

										tmp104869 := Call(__e, tmp104866, tmp104868)

										tmp104870 := Call(__e, tmp104865, tmp104869)

										tmp104871 := Call(__e, tmp104864, tmp104870)

										tmp104872 := Call(__e, tmp104863, tmp104871)

										tmp104873 := Call(__e, tmp104862, tmp104872)

										tmp104874 := Call(__e, tmp104861, Nil, tmp104873)

										var ifres104829 Obj

										if True == tmp104874 {
											tmp104853 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp104854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104856 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp104857 := Call(__e, tmp104856, V4629)

											tmp104858 := Call(__e, tmp104855, tmp104857)

											tmp104859 := Call(__e, tmp104854, tmp104858)

											tmp104860 := Call(__e, tmp104853, Nil, tmp104859)

											var ifres104831 Obj

											if True == tmp104860 {
												tmp104847 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

												tmp104848 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp104849 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp104850 := Call(__e, tmp104849, V4629)

												tmp104851 := Call(__e, tmp104848, tmp104850)

												tmp104852 := Call(__e, tmp104847, tmp104851)

												var ifres104833 Obj

												if True == tmp104852 {
													tmp104835 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

													tmp104836 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp104837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104838 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp104839 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp104841 := Call(__e, tmp104840, V4629)

													tmp104842 := Call(__e, tmp104839, tmp104841)

													tmp104843 := Call(__e, tmp104838, tmp104842)

													tmp104844 := Call(__e, tmp104837, tmp104843)

													tmp104845 := Call(__e, tmp104836, tmp104844)

													tmp104846 := Call(__e, tmp104835, tmp104845)

													var ifres104834 Obj

													if True == tmp104846 {
														ifres104834 = True

													} else {
														ifres104834 = False

													}

													ifres104833 = ifres104834

												} else {
													ifres104833 = False

												}

												var ifres104832 Obj

												if True == ifres104833 {
													ifres104832 = True

												} else {
													ifres104832 = False

												}

												ifres104831 = ifres104832

											} else {
												ifres104831 = False

											}

											var ifres104830 Obj

											if True == ifres104831 {
												ifres104830 = True

											} else {
												ifres104830 = False

											}

											ifres104829 = ifres104830

										} else {
											ifres104829 = False

										}

										var ifres104828 Obj

										if True == ifres104829 {
											ifres104828 = True

										} else {
											ifres104828 = False

										}

										ifres104827 = ifres104828

									} else {
										ifres104827 = False

									}

									var ifres104826 Obj

									if True == ifres104827 {
										ifres104826 = True

									} else {
										ifres104826 = False

									}

									ifres104825 = ifres104826

								} else {
									ifres104825 = False

								}

								var ifres104824 Obj

								if True == ifres104825 {
									ifres104824 = True

								} else {
									ifres104824 = False

								}

								ifres104823 = ifres104824

							} else {
								ifres104823 = False

							}

							var ifres104822 Obj

							if True == ifres104823 {
								ifres104822 = True

							} else {
								ifres104822 = False

							}

							ifres104821 = ifres104822

						} else {
							ifres104821 = False

						}

						var ifres104820 Obj

						if True == ifres104821 {
							ifres104820 = True

						} else {
							ifres104820 = False

						}

						ifres104819 = ifres104820

					} else {
						ifres104819 = False

					}

					var ifres104818 Obj

					if True == ifres104819 {
						ifres104818 = True

					} else {
						ifres104818 = False

					}

					ifres104817 = ifres104818

				} else {
					ifres104817 = False

				}

				var ifres104816 Obj

				if True == ifres104817 {
					ifres104816 = True

				} else {
					ifres104816 = False

				}

				ifres104815 = ifres104816

			} else {
				ifres104815 = False

			}

			var ifres104814 Obj

			if True == ifres104815 {
				ifres104814 = True

			} else {
				ifres104814 = False

			}

			ifres104813 = ifres104814

		} else {
			ifres104813 = False

		}

		if True == ifres104813 {
			tmp104784 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104785 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp104786 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp104787 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp104788 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104789 := Call(__e, tmp104788, V4629)

			tmp104790 := Call(__e, tmp104787, tmp104789)

			tmp104791 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp104792 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104793 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp104794 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104796 := Call(__e, tmp104795, V4629)

			tmp104797 := Call(__e, tmp104794, tmp104796)

			tmp104798 := Call(__e, tmp104793, tmp104797)

			tmp104799 := Call(__e, tmp104792, tmp104798)

			tmp104800 := Call(__e, tmp104791, tmp104799)

			tmp104801 := Call(__e, tmp104786, tmp104790, tmp104800)

			tmp104802 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104804 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp104805 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104806 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp104807 := Call(__e, tmp104806, V4629)

			tmp104808 := Call(__e, tmp104805, tmp104807)

			tmp104809 := Call(__e, tmp104804, tmp104808)

			tmp104810 := Call(__e, tmp104803, tmp104809)

			tmp104811 := Call(__e, tmp104802, tmp104810)

			tmp104812 := Call(__e, tmp104785, tmp104801, tmp104811)

			__e.TailApply(tmp104784, symcn, tmp104812)
			return

		} else {
			__e.Return(V4629)
			return
		}

	}, 1)

	tmp104931 := Call(__e, PrimNS1Value(symns2_1set), symshen_4factor_1cn, tmp104782)

	_ = tmp104931

	tmp104932 := MakeNative(func(__e *ControlFlow) {
		V4631 := __e.Get(1)
		_ = V4631
		tmp104977 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp104978 := Call(__e, tmp104977, MakeString(""), V4631)

		if True == tmp104978 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp104975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp104976 := Call(__e, tmp104975, V4631)

			var ifres104955 Obj

			if True == tmp104976 {
				tmp104971 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp104972 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp104973 := Call(__e, tmp104972, V4631, MakeNumber(0))

				tmp104974 := Call(__e, tmp104971, MakeString("~"), tmp104973)

				var ifres104957 Obj

				if True == tmp104974 {
					tmp104967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					tmp104968 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp104969 := Call(__e, tmp104968, V4631)

					tmp104970 := Call(__e, tmp104967, tmp104969)

					var ifres104959 Obj

					if True == tmp104970 {
						tmp104961 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp104962 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						tmp104963 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp104964 := Call(__e, tmp104963, V4631)

						tmp104965 := Call(__e, tmp104962, tmp104964, MakeNumber(0))

						tmp104966 := Call(__e, tmp104961, MakeString("%"), tmp104965)

						var ifres104960 Obj

						if True == tmp104966 {
							ifres104960 = True

						} else {
							ifres104960 = False

						}

						ifres104959 = ifres104960

					} else {
						ifres104959 = False

					}

					var ifres104958 Obj

					if True == ifres104959 {
						ifres104958 = True

					} else {
						ifres104958 = False

					}

					ifres104957 = ifres104958

				} else {
					ifres104957 = False

				}

				var ifres104956 Obj

				if True == ifres104957 {
					ifres104956 = True

				} else {
					ifres104956 = False

				}

				ifres104955 = ifres104956

			} else {
				ifres104955 = False

			}

			if True == ifres104955 {
				tmp104946 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp104947 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

				tmp104948 := Call(__e, tmp104947, MakeNumber(10))

				tmp104949 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

				tmp104950 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp104951 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp104952 := Call(__e, tmp104951, V4631)

				tmp104953 := Call(__e, tmp104950, tmp104952)

				tmp104954 := Call(__e, tmp104949, tmp104953)

				__e.TailApply(tmp104946, tmp104948, tmp104954)
				return

			} else {
				tmp104944 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				tmp104945 := Call(__e, tmp104944, V4631)

				if True == tmp104945 {
					tmp104937 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp104938 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					tmp104939 := Call(__e, tmp104938, V4631, MakeNumber(0))

					tmp104940 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

					tmp104941 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp104942 := Call(__e, tmp104941, V4631)

					tmp104943 := Call(__e, tmp104940, tmp104942)

					__e.TailApply(tmp104937, tmp104939, tmp104943)
					return

				} else {
					tmp104936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp104936, symshen_4proc_1nl)
					return

				}

			}

		}

	}, 1)

	tmp104979 := Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1nl, tmp104932)

	_ = tmp104979

	tmp104980 := MakeNative(func(__e *ControlFlow) {
		V4634 := __e.Get(1)
		_ = V4634
		V4635 := __e.Get(2)
		_ = V4635
		tmp104997 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp104998 := Call(__e, tmp104997, Nil, V4635)

		if True == tmp104998 {
			__e.Return(V4634)
			return
		} else {
			tmp104995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp104996 := Call(__e, tmp104995, V4635)

			if True == tmp104996 {
				tmp104984 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr_1r)

				tmp104985 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104986 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104987 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp104988 := Call(__e, tmp104987, V4635)

				tmp104989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp104990 := Call(__e, tmp104989, V4634, Nil)

				tmp104991 := Call(__e, tmp104986, tmp104988, tmp104990)

				tmp104992 := Call(__e, tmp104985, symshen_4insert, tmp104991)

				tmp104993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp104994 := Call(__e, tmp104993, V4635)

				__e.TailApply(tmp104984, tmp104992, tmp104994)
				return

			} else {
				tmp104983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp104983, symshen_4mkstr_1r)
				return

			}

		}

	}, 2)

	tmp104999 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mkstr_1r, tmp104980)

	_ = tmp104999

	tmp105000 := MakeNative(func(__e *ControlFlow) {
		V4638 := __e.Get(1)
		_ = V4638
		V4639 := __e.Get(2)
		_ = V4639
		tmp105001 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1h)

		__e.TailApply(tmp105001, V4638, V4639, MakeString(""))
		return

	}, 2)

	tmp105002 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert, tmp105000)

	_ = tmp105002

	tmp105003 := MakeNative(func(__e *ControlFlow) {
		V4645 := __e.Get(1)
		_ = V4645
		V4646 := __e.Get(2)
		_ = V4646
		V4647 := __e.Get(3)
		_ = V4647
		tmp105106 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105107 := Call(__e, tmp105106, MakeString(""), V4646)

		if True == tmp105107 {
			__e.Return(V4647)
			return
		} else {
			tmp105104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp105105 := Call(__e, tmp105104, V4646)

			var ifres105084 Obj

			if True == tmp105105 {
				tmp105100 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105101 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp105102 := Call(__e, tmp105101, V4646, MakeNumber(0))

				tmp105103 := Call(__e, tmp105100, MakeString("~"), tmp105102)

				var ifres105086 Obj

				if True == tmp105103 {
					tmp105096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					tmp105097 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp105098 := Call(__e, tmp105097, V4646)

					tmp105099 := Call(__e, tmp105096, tmp105098)

					var ifres105088 Obj

					if True == tmp105099 {
						tmp105090 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp105091 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						tmp105092 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp105093 := Call(__e, tmp105092, V4646)

						tmp105094 := Call(__e, tmp105091, tmp105093, MakeNumber(0))

						tmp105095 := Call(__e, tmp105090, MakeString("A"), tmp105094)

						var ifres105089 Obj

						if True == tmp105095 {
							ifres105089 = True

						} else {
							ifres105089 = False

						}

						ifres105088 = ifres105089

					} else {
						ifres105088 = False

					}

					var ifres105087 Obj

					if True == ifres105088 {
						ifres105087 = True

					} else {
						ifres105087 = False

					}

					ifres105086 = ifres105087

				} else {
					ifres105086 = False

				}

				var ifres105085 Obj

				if True == ifres105086 {
					ifres105085 = True

				} else {
					ifres105085 = False

				}

				ifres105084 = ifres105085

			} else {
				ifres105084 = False

			}

			if True == ifres105084 {
				tmp105077 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp105078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp105079 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp105080 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp105081 := Call(__e, tmp105080, V4646)

				tmp105082 := Call(__e, tmp105079, tmp105081)

				tmp105083 := Call(__e, tmp105078, V4645, tmp105082, symshen_4a)

				__e.TailApply(tmp105077, V4647, tmp105083)
				return

			} else {
				tmp105075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				tmp105076 := Call(__e, tmp105075, V4646)

				var ifres105055 Obj

				if True == tmp105076 {
					tmp105071 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105072 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					tmp105073 := Call(__e, tmp105072, V4646, MakeNumber(0))

					tmp105074 := Call(__e, tmp105071, MakeString("~"), tmp105073)

					var ifres105057 Obj

					if True == tmp105074 {
						tmp105067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						tmp105068 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp105069 := Call(__e, tmp105068, V4646)

						tmp105070 := Call(__e, tmp105067, tmp105069)

						var ifres105059 Obj

						if True == tmp105070 {
							tmp105061 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp105062 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							tmp105063 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp105064 := Call(__e, tmp105063, V4646)

							tmp105065 := Call(__e, tmp105062, tmp105064, MakeNumber(0))

							tmp105066 := Call(__e, tmp105061, MakeString("R"), tmp105065)

							var ifres105060 Obj

							if True == tmp105066 {
								ifres105060 = True

							} else {
								ifres105060 = False

							}

							ifres105059 = ifres105060

						} else {
							ifres105059 = False

						}

						var ifres105058 Obj

						if True == ifres105059 {
							ifres105058 = True

						} else {
							ifres105058 = False

						}

						ifres105057 = ifres105058

					} else {
						ifres105057 = False

					}

					var ifres105056 Obj

					if True == ifres105057 {
						ifres105056 = True

					} else {
						ifres105056 = False

					}

					ifres105055 = ifres105056

				} else {
					ifres105055 = False

				}

				if True == ifres105055 {
					tmp105048 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					tmp105049 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp105050 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp105051 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp105052 := Call(__e, tmp105051, V4646)

					tmp105053 := Call(__e, tmp105050, tmp105052)

					tmp105054 := Call(__e, tmp105049, V4645, tmp105053, symshen_4r)

					__e.TailApply(tmp105048, V4647, tmp105054)
					return

				} else {
					tmp105046 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

					tmp105047 := Call(__e, tmp105046, V4646)

					var ifres105026 Obj

					if True == tmp105047 {
						tmp105042 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp105043 := Call(__e, PrimNS1Value(symns2_1value), sympos)

						tmp105044 := Call(__e, tmp105043, V4646, MakeNumber(0))

						tmp105045 := Call(__e, tmp105042, MakeString("~"), tmp105044)

						var ifres105028 Obj

						if True == tmp105045 {
							tmp105038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

							tmp105039 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp105040 := Call(__e, tmp105039, V4646)

							tmp105041 := Call(__e, tmp105038, tmp105040)

							var ifres105030 Obj

							if True == tmp105041 {
								tmp105032 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp105033 := Call(__e, PrimNS1Value(symns2_1value), sympos)

								tmp105034 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

								tmp105035 := Call(__e, tmp105034, V4646)

								tmp105036 := Call(__e, tmp105033, tmp105035, MakeNumber(0))

								tmp105037 := Call(__e, tmp105032, MakeString("S"), tmp105036)

								var ifres105031 Obj

								if True == tmp105037 {
									ifres105031 = True

								} else {
									ifres105031 = False

								}

								ifres105030 = ifres105031

							} else {
								ifres105030 = False

							}

							var ifres105029 Obj

							if True == ifres105030 {
								ifres105029 = True

							} else {
								ifres105029 = False

							}

							ifres105028 = ifres105029

						} else {
							ifres105028 = False

						}

						var ifres105027 Obj

						if True == ifres105028 {
							ifres105027 = True

						} else {
							ifres105027 = False

						}

						ifres105026 = ifres105027

					} else {
						ifres105026 = False

					}

					if True == ifres105026 {
						tmp105019 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						tmp105020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						tmp105021 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp105022 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

						tmp105023 := Call(__e, tmp105022, V4646)

						tmp105024 := Call(__e, tmp105021, tmp105023)

						tmp105025 := Call(__e, tmp105020, V4645, tmp105024, symshen_4s)

						__e.TailApply(tmp105019, V4647, tmp105025)
						return

					} else {
						tmp105017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

						tmp105018 := Call(__e, tmp105017, V4646)

						if True == tmp105018 {
							tmp105010 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1h)

							tmp105011 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

							tmp105012 := Call(__e, tmp105011, V4646)

							tmp105013 := Call(__e, PrimNS1Value(symns2_1value), symcn)

							tmp105014 := Call(__e, PrimNS1Value(symns2_1value), sympos)

							tmp105015 := Call(__e, tmp105014, V4646, MakeNumber(0))

							tmp105016 := Call(__e, tmp105013, V4647, tmp105015)

							__e.TailApply(tmp105010, V4645, tmp105012, tmp105016)
							return

						} else {
							tmp105009 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(tmp105009, symshen_4insert_1h)
							return

						}

					}

				}

			}

		}

	}, 3)

	tmp105108 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1h, tmp105003)

	_ = tmp105108

	tmp105109 := MakeNative(func(__e *ControlFlow) {
		V4651 := __e.Get(1)
		_ = V4651
		V4652 := __e.Get(2)
		_ = V4652
		V4653 := __e.Get(3)
		_ = V4653
		tmp105110 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp105111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

		tmp105112 := Call(__e, tmp105111, V4651, V4653)

		__e.TailApply(tmp105110, tmp105112, V4652)
		return

	}, 3)

	tmp105113 := Call(__e, PrimNS1Value(symns2_1set), symshen_4app, tmp105109)

	_ = tmp105113

	tmp105114 := MakeNative(func(__e *ControlFlow) {
		V4661 := __e.Get(1)
		_ = V4661
		V4662 := __e.Get(2)
		_ = V4662
		tmp105129 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105130 := Call(__e, PrimNS1Value(symns2_1value), symfail)

		tmp105131 := Call(__e, tmp105130)

		tmp105132 := Call(__e, tmp105129, V4661, tmp105131)

		if True == tmp105132 {
			__e.Return(MakeString("..."))
			return
		} else {
			tmp105127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_2)

			tmp105128 := Call(__e, tmp105127, V4661)

			if True == tmp105128 {
				tmp105126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_1_6str)

				__e.TailApply(tmp105126, V4661, V4662)
				return

			} else {
				tmp105124 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				tmp105125 := Call(__e, tmp105124, V4661)

				if True == tmp105125 {
					tmp105123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4str_1_6str)

					__e.TailApply(tmp105123, V4661, V4662)
					return

				} else {
					tmp105121 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

					tmp105122 := Call(__e, tmp105121, V4661)

					if True == tmp105122 {
						tmp105120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4vector_1_6str)

						__e.TailApply(tmp105120, V4661, V4662)
						return

					} else {
						tmp105119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4atom_1_6str)

						__e.TailApply(tmp105119, V4661)
						return

					}

				}

			}

		}

	}, 2)

	tmp105133 := Call(__e, PrimNS1Value(symns2_1set), symshen_4arg_1_6str, tmp105114)

	_ = tmp105133

	tmp105134 := MakeNative(func(__e *ControlFlow) {
		V4665 := __e.Get(1)
		_ = V4665
		V4666 := __e.Get(2)
		_ = V4666
		tmp105150 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105151 := Call(__e, tmp105150, symshen_4r, V4666)

		if True == tmp105151 {
			tmp105143 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105144 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105145 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

			tmp105146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

			tmp105147 := Call(__e, tmp105146)

			tmp105148 := Call(__e, tmp105145, V4665, symshen_4r, tmp105147)

			tmp105149 := Call(__e, tmp105144, tmp105148, MakeString(")"))

			__e.TailApply(tmp105143, MakeString("("), tmp105149)
			return

		} else {
			tmp105136 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105137 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105138 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

			tmp105139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

			tmp105140 := Call(__e, tmp105139)

			tmp105141 := Call(__e, tmp105138, V4665, V4666, tmp105140)

			tmp105142 := Call(__e, tmp105137, tmp105141, MakeString("]"))

			__e.TailApply(tmp105136, MakeString("["), tmp105142)
			return

		}

	}, 2)

	tmp105152 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1_6str, tmp105134)

	_ = tmp105152

	tmp105153 := MakeNative(func(__e *ControlFlow) {
		tmp105154 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp105154, sym_dmaximum_1print_1sequence_1size_d)
		return

	}, 0)

	tmp105155 := Call(__e, PrimNS1Value(symns2_1set), symshen_4maxseq, tmp105153)

	_ = tmp105155

	tmp105156 := MakeNative(func(__e *ControlFlow) {
		V4680 := __e.Get(1)
		_ = V4680
		V4681 := __e.Get(2)
		_ = V4681
		V4682 := __e.Get(3)
		_ = V4682
		tmp105194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105195 := Call(__e, tmp105194, Nil, V4680)

		if True == tmp105195 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp105192 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105193 := Call(__e, tmp105192, MakeNumber(0), V4682)

			if True == tmp105193 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				tmp105190 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105191 := Call(__e, tmp105190, V4680)

				var ifres105184 Obj

				if True == tmp105191 {
					tmp105186 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105188 := Call(__e, tmp105187, V4680)

					tmp105189 := Call(__e, tmp105186, Nil, tmp105188)

					var ifres105185 Obj

					if True == tmp105189 {
						ifres105185 = True

					} else {
						ifres105185 = False

					}

					ifres105184 = ifres105185

				} else {
					ifres105184 = False

				}

				if True == ifres105184 {
					tmp105181 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

					tmp105182 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105183 := Call(__e, tmp105182, V4680)

					__e.TailApply(tmp105181, tmp105183, V4681)
					return

				} else {
					tmp105179 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105180 := Call(__e, tmp105179, V4680)

					if True == tmp105180 {
						tmp105166 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp105167 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

						tmp105168 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp105169 := Call(__e, tmp105168, V4680)

						tmp105170 := Call(__e, tmp105167, tmp105169, V4681)

						tmp105171 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp105172 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1list)

						tmp105173 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105174 := Call(__e, tmp105173, V4680)

						tmp105175 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						tmp105176 := Call(__e, tmp105175, V4682, MakeNumber(1))

						tmp105177 := Call(__e, tmp105172, tmp105174, V4681, tmp105176)

						tmp105178 := Call(__e, tmp105171, MakeString(" "), tmp105177)

						__e.TailApply(tmp105166, tmp105170, tmp105178)
						return

					} else {
						tmp105161 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp105162 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

						tmp105163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

						tmp105164 := Call(__e, tmp105163, V4680, V4681)

						tmp105165 := Call(__e, tmp105162, MakeString(" "), tmp105164)

						__e.TailApply(tmp105161, MakeString("|"), tmp105165)
						return

					}

				}

			}

		}

	}, 3)

	tmp105196 := Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1list, tmp105156)

	_ = tmp105196

	tmp105197 := MakeNative(func(__e *ControlFlow) {
		V4689 := __e.Get(1)
		_ = V4689
		V4690 := __e.Get(2)
		_ = V4690
		tmp105206 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105207 := Call(__e, tmp105206, symshen_4a, V4690)

		if True == tmp105207 {
			__e.Return(V4689)
			return
		} else {
			tmp105199 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105200 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			tmp105201 := Call(__e, tmp105200, MakeNumber(34))

			tmp105202 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

			tmp105203 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			tmp105204 := Call(__e, tmp105203, MakeNumber(34))

			tmp105205 := Call(__e, tmp105202, V4689, tmp105204)

			__e.TailApply(tmp105199, tmp105201, tmp105205)
			return

		}

	}, 2)

	tmp105208 := Call(__e, PrimNS1Value(symns2_1set), symshen_4str_1_6str, tmp105197)

	_ = tmp105208

	tmp105209 := MakeNative(func(__e *ControlFlow) {
		V4693 := __e.Get(1)
		_ = V4693
		V4694 := __e.Get(2)
		_ = V4694
		tmp105234 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1vector_2)

		tmp105235 := Call(__e, tmp105234, V4693)

		if True == tmp105235 {
			tmp105230 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

			tmp105231 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp105232 := Call(__e, tmp105231, V4693, MakeNumber(0))

			tmp105233 := Call(__e, tmp105230, tmp105232)

			__e.TailApply(tmp105233, V4693)
			return

		} else {
			tmp105228 := Call(__e, PrimNS1Value(symns2_1value), symvector_2)

			tmp105229 := Call(__e, tmp105228, V4693)

			if True == tmp105229 {
				tmp105221 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				tmp105222 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				tmp105223 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

				tmp105224 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

				tmp105225 := Call(__e, tmp105224)

				tmp105226 := Call(__e, tmp105223, V4693, MakeNumber(1), V4694, tmp105225)

				tmp105227 := Call(__e, tmp105222, tmp105226, MakeString(">"))

				__e.TailApply(tmp105221, MakeString("<"), tmp105227)
				return

			} else {
				tmp105212 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				tmp105213 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				tmp105214 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

				tmp105215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

				tmp105216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxseq)

				tmp105217 := Call(__e, tmp105216)

				tmp105218 := Call(__e, tmp105215, V4693, MakeNumber(0), V4694, tmp105217)

				tmp105219 := Call(__e, tmp105214, tmp105218, MakeString(">>"))

				tmp105220 := Call(__e, tmp105213, MakeString("<"), tmp105219)

				__e.TailApply(tmp105212, MakeString("<"), tmp105220)
				return

			}

		}

	}, 2)

	tmp105236 := Call(__e, PrimNS1Value(symns2_1set), symshen_4vector_1_6str, tmp105209)

	_ = tmp105236

	tmp105237 := MakeNative(func(__e *ControlFlow) {
		V4696 := __e.Get(1)
		_ = V4696
		tmp105238 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105239 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp105240 := Call(__e, tmp105239, symshen_4_dempty_1absvector_d)

		__e.TailApply(tmp105238, V4696, tmp105240)
		return

	}, 1)

	tmp105241 := Call(__e, PrimNS1Value(symns2_1set), symshen_4empty_1absvector_2, tmp105237)

	_ = tmp105241

	tmp105242 := MakeNative(func(__e *ControlFlow) {
		V4698 := __e.Get(1)
		_ = V4698
		tmp105269 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		tmp105270 := Call(__e, PrimNS1Value(symns2_1value), symshen_4empty_1absvector_2)

		tmp105271 := Call(__e, tmp105270, V4698)

		tmp105272 := Call(__e, tmp105269, tmp105271)

		if True == tmp105272 {
			tmp105245 := MakeNative(func(__e *ControlFlow) {
				First := __e.Get(1)
				_ = First
				tmp105264 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105265 := Call(__e, tmp105264, First, symshen_4tuple)

				if True == tmp105265 {
					__e.Return(True)
					return
				} else {
					tmp105262 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105263 := Call(__e, tmp105262, First, symshen_4pvar)

					var ifres105248 Obj

					if True == tmp105263 {
						ifres105248 = True

					} else {
						tmp105260 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp105261 := Call(__e, tmp105260, First, symshen_4dictionary)

						var ifres105250 Obj

						if True == tmp105261 {
							ifres105250 = True

						} else {
							tmp105256 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp105257 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

							tmp105258 := Call(__e, tmp105257, First)

							tmp105259 := Call(__e, tmp105256, tmp105258)

							var ifres105252 Obj

							if True == tmp105259 {
								tmp105254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fbound_2)

								tmp105255 := Call(__e, tmp105254, First)

								var ifres105253 Obj

								if True == tmp105255 {
									ifres105253 = True

								} else {
									ifres105253 = False

								}

								ifres105252 = ifres105253

							} else {
								ifres105252 = False

							}

							var ifres105251 Obj

							if True == ifres105252 {
								ifres105251 = True

							} else {
								ifres105251 = False

							}

							ifres105250 = ifres105251

						}

						var ifres105249 Obj

						if True == ifres105250 {
							ifres105249 = True

						} else {
							ifres105249 = False

						}

						ifres105248 = ifres105249

					}

					if True == ifres105248 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}, 1)

			tmp105266 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp105267 := Call(__e, tmp105266, V4698, MakeNumber(0))

			tmp105268 := Call(__e, tmp105245, tmp105267)

			if True == tmp105268 {
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

	tmp105273 := Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1vector_2, tmp105242)

	_ = tmp105273

	tmp105274 := MakeNative(func(__e *ControlFlow) {
		V4700 := __e.Get(1)
		_ = V4700
		tmp105275 := MakeNative(func(__e *ControlFlow) {
			tmp105276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lookup_1func)

			tmp105277 := Call(__e, tmp105276, V4700)

			_ = tmp105277

			__e.Return(True)
			return

		}, 0)

		tmp105278 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp105275, tmp105278)
		return

	}, 1)

	tmp105279 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fbound_2, tmp105274)

	_ = tmp105279

	tmp105280 := MakeNative(func(__e *ControlFlow) {
		V4702 := __e.Get(1)
		_ = V4702
		tmp105281 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp105282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp105283 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp105284 := Call(__e, tmp105283, V4702, MakeNumber(1))

		tmp105285 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp105286 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp105287 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		tmp105288 := Call(__e, tmp105287, V4702, MakeNumber(2))

		tmp105289 := Call(__e, tmp105286, tmp105288, MakeString(")"), symshen_4s)

		tmp105290 := Call(__e, tmp105285, MakeString(" "), tmp105289)

		tmp105291 := Call(__e, tmp105282, tmp105284, tmp105290, symshen_4s)

		__e.TailApply(tmp105281, MakeString("(@p "), tmp105291)
		return

	}, 1)

	tmp105292 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple, tmp105280)

	_ = tmp105292

	tmp105293 := MakeNative(func(__e *ControlFlow) {
		V4704 := __e.Get(1)
		_ = V4704
		__e.Return(MakeString("(dict ...)"))
		return
	}, 1)

	tmp105294 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dictionary, tmp105293)

	_ = tmp105294

	tmp105295 := MakeNative(func(__e *ControlFlow) {
		V4715 := __e.Get(1)
		_ = V4715
		V4716 := __e.Get(2)
		_ = V4716
		V4717 := __e.Get(3)
		_ = V4717
		V4718 := __e.Get(4)
		_ = V4718
		tmp105327 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105328 := Call(__e, tmp105327, MakeNumber(0), V4718)

		if True == tmp105328 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			tmp105297 := MakeNative(func(__e *ControlFlow) {
				Item := __e.Get(1)
				_ = Item
				tmp105298 := MakeNative(func(__e *ControlFlow) {
					Next := __e.Get(1)
					_ = Next
					tmp105315 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105316 := Call(__e, tmp105315, Item, symshen_4out_1of_1bounds)

					if True == tmp105316 {
						__e.Return(MakeString(""))
						return
					} else {
						tmp105313 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp105314 := Call(__e, tmp105313, Next, symshen_4out_1of_1bounds)

						if True == tmp105314 {
							tmp105312 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

							__e.TailApply(tmp105312, Item, V4717)
							return

						} else {
							tmp105301 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

							tmp105302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

							tmp105303 := Call(__e, tmp105302, Item, V4717)

							tmp105304 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

							tmp105305 := Call(__e, PrimNS1Value(symns2_1value), symshen_4iter_1vector)

							tmp105306 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

							tmp105307 := Call(__e, tmp105306, V4716, MakeNumber(1))

							tmp105308 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

							tmp105309 := Call(__e, tmp105308, V4718, MakeNumber(1))

							tmp105310 := Call(__e, tmp105305, V4715, tmp105307, V4717, tmp105309)

							tmp105311 := Call(__e, tmp105304, MakeString(" "), tmp105310)

							__e.TailApply(tmp105301, tmp105303, tmp105311)
							return

						}

					}

				}, 1)

				tmp105317 := MakeNative(func(__e *ControlFlow) {
					tmp105318 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

					tmp105319 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					tmp105320 := Call(__e, tmp105319, V4716, MakeNumber(1))

					__e.TailApply(tmp105318, V4715, tmp105320)
					return

				}, 0)

				tmp105321 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				tmp105322 := Call(__e, PrimNS1Value(symtry_1catch), tmp105317, tmp105321)

				__e.TailApply(tmp105298, tmp105322)
				return

			}, 1)

			tmp105323 := MakeNative(func(__e *ControlFlow) {
				tmp105324 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(tmp105324, V4715, V4716)
				return

			}, 0)

			tmp105325 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4out_1of_1bounds)
				return
			}, 1)

			tmp105326 := Call(__e, PrimNS1Value(symtry_1catch), tmp105323, tmp105325)

			__e.TailApply(tmp105297, tmp105326)
			return

		}

	}, 4)

	tmp105329 := Call(__e, PrimNS1Value(symns2_1set), symshen_4iter_1vector, tmp105295)

	_ = tmp105329

	tmp105330 := MakeNative(func(__e *ControlFlow) {
		V4720 := __e.Get(1)
		_ = V4720
		tmp105331 := MakeNative(func(__e *ControlFlow) {
			tmp105332 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			__e.TailApply(tmp105332, V4720)
			return

		}, 0)

		tmp105333 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp105334 := Call(__e, PrimNS1Value(symns2_1value), symshen_4funexstring)

			__e.TailApply(tmp105334)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp105331, tmp105333)
		return

	}, 1)

	tmp105335 := Call(__e, PrimNS1Value(symns2_1set), symshen_4atom_1_6str, tmp105330)

	_ = tmp105335

	tmp105336 := MakeNative(func(__e *ControlFlow) {
		tmp105337 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105338 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105339 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105340 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105341 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105342 := Call(__e, PrimNS1Value(symns2_1value), sym_8s)

		tmp105343 := Call(__e, PrimNS1Value(symns2_1value), symshen_4arg_1_6str)

		tmp105344 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		tmp105345 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp105346 := Call(__e, tmp105345, MakeString("x"))

		tmp105347 := Call(__e, tmp105344, tmp105346)

		tmp105348 := Call(__e, tmp105343, tmp105347, symshen_4a)

		tmp105349 := Call(__e, tmp105342, tmp105348, MakeString("\x11"))

		tmp105350 := Call(__e, tmp105341, MakeString("e"), tmp105349)

		tmp105351 := Call(__e, tmp105340, MakeString("n"), tmp105350)

		tmp105352 := Call(__e, tmp105339, MakeString("u"), tmp105351)

		tmp105353 := Call(__e, tmp105338, MakeString("f"), tmp105352)

		__e.TailApply(tmp105337, MakeString("\x10"), tmp105353)
		return

	}, 0)

	tmp105354 := Call(__e, PrimNS1Value(symns2_1set), symshen_4funexstring, tmp105336)

	_ = tmp105354

	tmp105355 := MakeNative(func(__e *ControlFlow) {
		V4722 := __e.Get(1)
		_ = V4722
		tmp105360 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

		tmp105361 := Call(__e, tmp105360, V4722)

		if True == tmp105361 {
			__e.Return(True)
			return
		} else {
			tmp105358 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105359 := Call(__e, tmp105358, V4722)

			if True == tmp105359 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4list_2, tmp105355)
	return

}, 0)
