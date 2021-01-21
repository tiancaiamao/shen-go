package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp91514 := MakeNative(func(__e *ControlFlow) {
		V1920 := __e.Get(1)
		_ = V1920
		__e.TailApply(V1920)
		return
	}, 1)

	tmp91515 := Call(__e, PrimNS1Value(symns2_1set), symthaw, tmp91514)

	_ = tmp91515

	tmp91516 := MakeNative(func(__e *ControlFlow) {
		V1922 := __e.Get(1)
		_ = V1922
		tmp91517 := MakeNative(func(__e *ControlFlow) {
			Macroexpand := __e.Get(1)
			_ = Macroexpand
			tmp91525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packaged_2)

			tmp91526 := Call(__e, tmp91525, Macroexpand)

			if True == tmp91526 {
				tmp91520 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp91521 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp91522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

					__e.TailApply(tmp91522, Z)
					return

				}, 1)

				tmp91523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4package_1contents)

				tmp91524 := Call(__e, tmp91523, Macroexpand)

				__e.TailApply(tmp91520, tmp91521, tmp91524)
				return

			} else {
				tmp91519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				__e.TailApply(tmp91519, Macroexpand)
				return

			}

		}, 1)

		tmp91527 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

		tmp91528 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			tmp91529 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

			__e.TailApply(tmp91529, Y)
			return

		}, 1)

		tmp91530 := Call(__e, tmp91527, tmp91528, V1922)

		__e.TailApply(tmp91517, tmp91530)
		return

	}, 1)

	tmp91531 := Call(__e, PrimNS1Value(symns2_1set), symeval, tmp91516)

	_ = tmp91531

	tmp91532 := MakeNative(func(__e *ControlFlow) {
		V1924 := __e.Get(1)
		_ = V1924
		tmp91533 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

		tmp91534 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

		tmp91535 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1input_7)

		tmp91536 := Call(__e, tmp91535, V1924)

		tmp91537 := Call(__e, tmp91534, tmp91536)

		__e.TailApply(tmp91533, tmp91537)
		return

	}, 1)

	tmp91538 := Call(__e, PrimNS1Value(symns2_1set), symshen_4eval_1without_1macros, tmp91532)

	_ = tmp91538

	tmp91539 := MakeNative(func(__e *ControlFlow) {
		V1926 := __e.Get(1)
		_ = V1926
		tmp91636 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91637 := Call(__e, tmp91636, V1926)

		var ifres91606 Obj

		if True == tmp91637 {
			tmp91632 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91633 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91634 := Call(__e, tmp91633, V1926)

			tmp91635 := Call(__e, tmp91632, syminput_7, tmp91634)

			var ifres91608 Obj

			if True == tmp91635 {
				tmp91628 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91629 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91630 := Call(__e, tmp91629, V1926)

				tmp91631 := Call(__e, tmp91628, tmp91630)

				var ifres91610 Obj

				if True == tmp91631 {
					tmp91622 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91624 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91625 := Call(__e, tmp91624, V1926)

					tmp91626 := Call(__e, tmp91623, tmp91625)

					tmp91627 := Call(__e, tmp91622, tmp91626)

					var ifres91612 Obj

					if True == tmp91627 {
						tmp91614 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp91615 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91616 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91617 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91618 := Call(__e, tmp91617, V1926)

						tmp91619 := Call(__e, tmp91616, tmp91618)

						tmp91620 := Call(__e, tmp91615, tmp91619)

						tmp91621 := Call(__e, tmp91614, Nil, tmp91620)

						var ifres91613 Obj

						if True == tmp91621 {
							ifres91613 = True

						} else {
							ifres91613 = False

						}

						ifres91612 = ifres91613

					} else {
						ifres91612 = False

					}

					var ifres91611 Obj

					if True == ifres91612 {
						ifres91611 = True

					} else {
						ifres91611 = False

					}

					ifres91610 = ifres91611

				} else {
					ifres91610 = False

				}

				var ifres91609 Obj

				if True == ifres91610 {
					ifres91609 = True

				} else {
					ifres91609 = False

				}

				ifres91608 = ifres91609

			} else {
				ifres91608 = False

			}

			var ifres91607 Obj

			if True == ifres91608 {
				ifres91607 = True

			} else {
				ifres91607 = False

			}

			ifres91606 = ifres91607

		} else {
			ifres91606 = False

		}

		if True == ifres91606 {
			tmp91593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91594 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp91596 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91597 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91598 := Call(__e, tmp91597, V1926)

			tmp91599 := Call(__e, tmp91596, tmp91598)

			tmp91600 := Call(__e, tmp91595, tmp91599)

			tmp91601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91603 := Call(__e, tmp91602, V1926)

			tmp91604 := Call(__e, tmp91601, tmp91603)

			tmp91605 := Call(__e, tmp91594, tmp91600, tmp91604)

			__e.TailApply(tmp91593, syminput_7, tmp91605)
			return

		} else {
			tmp91591 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91592 := Call(__e, tmp91591, V1926)

			var ifres91561 Obj

			if True == tmp91592 {
				tmp91587 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91588 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91589 := Call(__e, tmp91588, V1926)

				tmp91590 := Call(__e, tmp91587, symshen_4read_7, tmp91589)

				var ifres91563 Obj

				if True == tmp91590 {
					tmp91583 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91584 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91585 := Call(__e, tmp91584, V1926)

					tmp91586 := Call(__e, tmp91583, tmp91585)

					var ifres91565 Obj

					if True == tmp91586 {
						tmp91577 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91578 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91579 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91580 := Call(__e, tmp91579, V1926)

						tmp91581 := Call(__e, tmp91578, tmp91580)

						tmp91582 := Call(__e, tmp91577, tmp91581)

						var ifres91567 Obj

						if True == tmp91582 {
							tmp91569 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp91570 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91571 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91572 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91573 := Call(__e, tmp91572, V1926)

							tmp91574 := Call(__e, tmp91571, tmp91573)

							tmp91575 := Call(__e, tmp91570, tmp91574)

							tmp91576 := Call(__e, tmp91569, Nil, tmp91575)

							var ifres91568 Obj

							if True == tmp91576 {
								ifres91568 = True

							} else {
								ifres91568 = False

							}

							ifres91567 = ifres91568

						} else {
							ifres91567 = False

						}

						var ifres91566 Obj

						if True == ifres91567 {
							ifres91566 = True

						} else {
							ifres91566 = False

						}

						ifres91565 = ifres91566

					} else {
						ifres91565 = False

					}

					var ifres91564 Obj

					if True == ifres91565 {
						ifres91564 = True

					} else {
						ifres91564 = False

					}

					ifres91563 = ifres91564

				} else {
					ifres91563 = False

				}

				var ifres91562 Obj

				if True == ifres91563 {
					ifres91562 = True

				} else {
					ifres91562 = False

				}

				ifres91561 = ifres91562

			} else {
				ifres91561 = False

			}

			if True == ifres91561 {
				tmp91548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91550 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

				tmp91551 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91552 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91553 := Call(__e, tmp91552, V1926)

				tmp91554 := Call(__e, tmp91551, tmp91553)

				tmp91555 := Call(__e, tmp91550, tmp91554)

				tmp91556 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91558 := Call(__e, tmp91557, V1926)

				tmp91559 := Call(__e, tmp91556, tmp91558)

				tmp91560 := Call(__e, tmp91549, tmp91555, tmp91559)

				__e.TailApply(tmp91548, symshen_4read_7, tmp91560)
				return

			} else {
				tmp91546 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91547 := Call(__e, tmp91546, V1926)

				if True == tmp91547 {
					tmp91543 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					tmp91544 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						tmp91545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1input_7)

						__e.TailApply(tmp91545, Z)
						return

					}, 1)

					__e.TailApply(tmp91543, tmp91544, V1926)
					return

				} else {
					__e.Return(V1926)
					return
				}

			}

		}

	}, 1)

	tmp91638 := Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1input_7, tmp91539)

	_ = tmp91638

	tmp91639 := MakeNative(func(__e *ControlFlow) {
		V1928 := __e.Get(1)
		_ = V1928
		tmp91732 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91733 := Call(__e, tmp91732, V1928)

		var ifres91720 Obj

		if True == tmp91733 {
			tmp91728 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91729 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91730 := Call(__e, tmp91729, V1928)

			tmp91731 := Call(__e, tmp91728, symdefine, tmp91730)

			var ifres91722 Obj

			if True == tmp91731 {
				tmp91724 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91725 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91726 := Call(__e, tmp91725, V1928)

				tmp91727 := Call(__e, tmp91724, tmp91726)

				var ifres91723 Obj

				if True == tmp91727 {
					ifres91723 = True

				} else {
					ifres91723 = False

				}

				ifres91722 = ifres91723

			} else {
				ifres91722 = False

			}

			var ifres91721 Obj

			if True == ifres91722 {
				ifres91721 = True

			} else {
				ifres91721 = False

			}

			ifres91720 = ifres91721

		} else {
			ifres91720 = False

		}

		if True == ifres91720 {
			tmp91711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4shen_1_6kl)

			tmp91712 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91713 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91714 := Call(__e, tmp91713, V1928)

			tmp91715 := Call(__e, tmp91712, tmp91714)

			tmp91716 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91717 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91718 := Call(__e, tmp91717, V1928)

			tmp91719 := Call(__e, tmp91716, tmp91718)

			__e.TailApply(tmp91711, tmp91715, tmp91719)
			return

		} else {
			tmp91709 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91710 := Call(__e, tmp91709, V1928)

			var ifres91697 Obj

			if True == tmp91710 {
				tmp91705 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91706 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91707 := Call(__e, tmp91706, V1928)

				tmp91708 := Call(__e, tmp91705, symdefmacro, tmp91707)

				var ifres91699 Obj

				if True == tmp91708 {
					tmp91701 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91703 := Call(__e, tmp91702, V1928)

					tmp91704 := Call(__e, tmp91701, tmp91703)

					var ifres91700 Obj

					if True == tmp91704 {
						ifres91700 = True

					} else {
						ifres91700 = False

					}

					ifres91699 = ifres91700

				} else {
					ifres91699 = False

				}

				var ifres91698 Obj

				if True == ifres91699 {
					ifres91698 = True

				} else {
					ifres91698 = False

				}

				ifres91697 = ifres91698

			} else {
				ifres91697 = False

			}

			if True == ifres91697 {
				tmp91666 := MakeNative(func(__e *ControlFlow) {
					Default := __e.Get(1)
					_ = Default
					tmp91667 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						tmp91668 := MakeNative(func(__e *ControlFlow) {
							MacroAdd := __e.Get(1)
							_ = MacroAdd
							__e.Return(Def)
							return
						}, 1)

						tmp91669 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1macro)

						tmp91670 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91671 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91672 := Call(__e, tmp91671, V1928)

						tmp91673 := Call(__e, tmp91670, tmp91672)

						tmp91674 := Call(__e, tmp91669, tmp91673)

						__e.TailApply(tmp91668, tmp91674)
						return

					}, 1)

					tmp91675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

					tmp91676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91678 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91679 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91680 := Call(__e, tmp91679, V1928)

					tmp91681 := Call(__e, tmp91678, tmp91680)

					tmp91682 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp91683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91684 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91685 := Call(__e, tmp91684, V1928)

					tmp91686 := Call(__e, tmp91683, tmp91685)

					tmp91687 := Call(__e, tmp91682, tmp91686, Default)

					tmp91688 := Call(__e, tmp91677, tmp91681, tmp91687)

					tmp91689 := Call(__e, tmp91676, symdefine, tmp91688)

					tmp91690 := Call(__e, tmp91675, tmp91689)

					__e.TailApply(tmp91667, tmp91690)
					return

				}, 1)

				tmp91691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91694 := Call(__e, tmp91693, symX, Nil)

				tmp91695 := Call(__e, tmp91692, sym_1_6, tmp91694)

				tmp91696 := Call(__e, tmp91691, symX, tmp91695)

				__e.TailApply(tmp91666, tmp91696)
				return

			} else {
				tmp91664 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91665 := Call(__e, tmp91664, V1928)

				var ifres91652 Obj

				if True == tmp91665 {
					tmp91660 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp91661 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91662 := Call(__e, tmp91661, V1928)

					tmp91663 := Call(__e, tmp91660, symdefcc, tmp91662)

					var ifres91654 Obj

					if True == tmp91663 {
						tmp91656 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91657 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91658 := Call(__e, tmp91657, V1928)

						tmp91659 := Call(__e, tmp91656, tmp91658)

						var ifres91655 Obj

						if True == tmp91659 {
							ifres91655 = True

						} else {
							ifres91655 = False

						}

						ifres91654 = ifres91655

					} else {
						ifres91654 = False

					}

					var ifres91653 Obj

					if True == ifres91654 {
						ifres91653 = True

					} else {
						ifres91653 = False

					}

					ifres91652 = ifres91653

				} else {
					ifres91652 = False

				}

				if True == ifres91652 {
					tmp91649 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

					tmp91650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc)

					tmp91651 := Call(__e, tmp91650, V1928)

					__e.TailApply(tmp91649, tmp91651)
					return

				} else {
					tmp91647 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91648 := Call(__e, tmp91647, V1928)

					if True == tmp91648 {
						tmp91644 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						tmp91645 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							tmp91646 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

							__e.TailApply(tmp91646, Z)
							return

						}, 1)

						__e.TailApply(tmp91644, tmp91645, V1928)
						return

					} else {
						__e.Return(V1928)
						return
					}

				}

			}

		}

	}, 1)

	tmp91734 := Call(__e, PrimNS1Value(symns2_1set), symshen_4elim_1def, tmp91639)

	_ = tmp91734

	tmp91735 := MakeNative(func(__e *ControlFlow) {
		V1930 := __e.Get(1)
		_ = V1930
		tmp91736 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			tmp91737 := MakeNative(func(__e *ControlFlow) {
				NewMacroReg := __e.Get(1)
				_ = NewMacroReg
				tmp91746 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91747 := Call(__e, tmp91746, MacroReg, NewMacroReg)

				if True == tmp91747 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp91739 := Call(__e, PrimNS1Value(symns2_1value), symset)

					tmp91740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91741 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

					tmp91742 := Call(__e, tmp91741, V1930)

					tmp91743 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					tmp91744 := Call(__e, tmp91743, sym_dmacros_d)

					tmp91745 := Call(__e, tmp91740, tmp91742, tmp91744)

					__e.TailApply(tmp91739, sym_dmacros_d, tmp91745)
					return

				}

			}, 1)

			tmp91748 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp91749 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			tmp91750 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91751 := Call(__e, tmp91750, symshen_4_dmacroreg_d)

			tmp91752 := Call(__e, tmp91749, V1930, tmp91751)

			tmp91753 := Call(__e, tmp91748, symshen_4_dmacroreg_d, tmp91752)

			__e.TailApply(tmp91737, tmp91753)
			return

		}, 1)

		tmp91754 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91755 := Call(__e, tmp91754, symshen_4_dmacroreg_d)

		__e.TailApply(tmp91736, tmp91755)
		return

	}, 1)

	tmp91756 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1macro, tmp91735)

	_ = tmp91756

	tmp91757 := MakeNative(func(__e *ControlFlow) {
		V1938 := __e.Get(1)
		_ = V1938
		tmp91779 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91780 := Call(__e, tmp91779, V1938)

		var ifres91759 Obj

		if True == tmp91780 {
			tmp91775 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91777 := Call(__e, tmp91776, V1938)

			tmp91778 := Call(__e, tmp91775, sympackage, tmp91777)

			var ifres91761 Obj

			if True == tmp91778 {
				tmp91771 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91772 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91773 := Call(__e, tmp91772, V1938)

				tmp91774 := Call(__e, tmp91771, tmp91773)

				var ifres91763 Obj

				if True == tmp91774 {
					tmp91765 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91767 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91768 := Call(__e, tmp91767, V1938)

					tmp91769 := Call(__e, tmp91766, tmp91768)

					tmp91770 := Call(__e, tmp91765, tmp91769)

					var ifres91764 Obj

					if True == tmp91770 {
						ifres91764 = True

					} else {
						ifres91764 = False

					}

					ifres91763 = ifres91764

				} else {
					ifres91763 = False

				}

				var ifres91762 Obj

				if True == ifres91763 {
					ifres91762 = True

				} else {
					ifres91762 = False

				}

				ifres91761 = ifres91762

			} else {
				ifres91761 = False

			}

			var ifres91760 Obj

			if True == ifres91761 {
				ifres91760 = True

			} else {
				ifres91760 = False

			}

			ifres91759 = ifres91760

		} else {
			ifres91759 = False

		}

		if True == ifres91759 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp91781 := Call(__e, PrimNS1Value(symns2_1set), symshen_4packaged_2, tmp91757)

	_ = tmp91781

	tmp91782 := MakeNative(func(__e *ControlFlow) {
		V1940 := __e.Get(1)
		_ = V1940
		tmp91783 := MakeNative(func(__e *ControlFlow) {
			tmp91784 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp91785 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91786 := Call(__e, tmp91785, sym_dproperty_1vector_d)

			__e.TailApply(tmp91784, V1940, symshen_4external_1symbols, tmp91786)
			return

		}, 0)

		tmp91787 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp91788 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp91789 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp91790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp91791 := Call(__e, tmp91790, V1940, MakeString(" has not been used.\n"), symshen_4a)

			tmp91792 := Call(__e, tmp91789, MakeString("package "), tmp91791)

			__e.TailApply(tmp91788, tmp91792)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp91783, tmp91787)
		return

	}, 1)

	tmp91793 := Call(__e, PrimNS1Value(symns2_1set), symexternal, tmp91782)

	_ = tmp91793

	tmp91794 := MakeNative(func(__e *ControlFlow) {
		V1942 := __e.Get(1)
		_ = V1942
		tmp91795 := MakeNative(func(__e *ControlFlow) {
			tmp91796 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp91797 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91798 := Call(__e, tmp91797, sym_dproperty_1vector_d)

			__e.TailApply(tmp91796, V1942, symshen_4internal_1symbols, tmp91798)
			return

		}, 0)

		tmp91799 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp91800 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp91801 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp91802 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp91803 := Call(__e, tmp91802, V1942, MakeString(" has not been used.\n"), symshen_4a)

			tmp91804 := Call(__e, tmp91801, MakeString("package "), tmp91803)

			__e.TailApply(tmp91800, tmp91804)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp91795, tmp91799)
		return

	}, 1)

	tmp91805 := Call(__e, PrimNS1Value(symns2_1set), syminternal, tmp91794)

	_ = tmp91805

	tmp91806 := MakeNative(func(__e *ControlFlow) {
		V1946 := __e.Get(1)
		_ = V1946
		tmp91896 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91897 := Call(__e, tmp91896, V1946)

		var ifres91868 Obj

		if True == tmp91897 {
			tmp91892 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91893 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91894 := Call(__e, tmp91893, V1946)

			tmp91895 := Call(__e, tmp91892, sympackage, tmp91894)

			var ifres91870 Obj

			if True == tmp91895 {
				tmp91888 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91889 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91890 := Call(__e, tmp91889, V1946)

				tmp91891 := Call(__e, tmp91888, tmp91890)

				var ifres91872 Obj

				if True == tmp91891 {
					tmp91882 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp91883 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91885 := Call(__e, tmp91884, V1946)

					tmp91886 := Call(__e, tmp91883, tmp91885)

					tmp91887 := Call(__e, tmp91882, symnull, tmp91886)

					var ifres91874 Obj

					if True == tmp91887 {
						tmp91876 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91879 := Call(__e, tmp91878, V1946)

						tmp91880 := Call(__e, tmp91877, tmp91879)

						tmp91881 := Call(__e, tmp91876, tmp91880)

						var ifres91875 Obj

						if True == tmp91881 {
							ifres91875 = True

						} else {
							ifres91875 = False

						}

						ifres91874 = ifres91875

					} else {
						ifres91874 = False

					}

					var ifres91873 Obj

					if True == ifres91874 {
						ifres91873 = True

					} else {
						ifres91873 = False

					}

					ifres91872 = ifres91873

				} else {
					ifres91872 = False

				}

				var ifres91871 Obj

				if True == ifres91872 {
					ifres91871 = True

				} else {
					ifres91871 = False

				}

				ifres91870 = ifres91871

			} else {
				ifres91870 = False

			}

			var ifres91869 Obj

			if True == ifres91870 {
				ifres91869 = True

			} else {
				ifres91869 = False

			}

			ifres91868 = ifres91869

		} else {
			ifres91868 = False

		}

		if True == ifres91868 {
			tmp91863 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91864 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91865 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91866 := Call(__e, tmp91865, V1946)

			tmp91867 := Call(__e, tmp91864, tmp91866)

			__e.TailApply(tmp91863, tmp91867)
			return

		} else {
			tmp91861 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91862 := Call(__e, tmp91861, V1946)

			var ifres91841 Obj

			if True == tmp91862 {
				tmp91857 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91858 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91859 := Call(__e, tmp91858, V1946)

				tmp91860 := Call(__e, tmp91857, sympackage, tmp91859)

				var ifres91843 Obj

				if True == tmp91860 {
					tmp91853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91855 := Call(__e, tmp91854, V1946)

					tmp91856 := Call(__e, tmp91853, tmp91855)

					var ifres91845 Obj

					if True == tmp91856 {
						tmp91847 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91849 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91850 := Call(__e, tmp91849, V1946)

						tmp91851 := Call(__e, tmp91848, tmp91850)

						tmp91852 := Call(__e, tmp91847, tmp91851)

						var ifres91846 Obj

						if True == tmp91852 {
							ifres91846 = True

						} else {
							ifres91846 = False

						}

						ifres91845 = ifres91846

					} else {
						ifres91845 = False

					}

					var ifres91844 Obj

					if True == ifres91845 {
						ifres91844 = True

					} else {
						ifres91844 = False

					}

					ifres91843 = ifres91844

				} else {
					ifres91843 = False

				}

				var ifres91842 Obj

				if True == ifres91843 {
					ifres91842 = True

				} else {
					ifres91842 = False

				}

				ifres91841 = ifres91842

			} else {
				ifres91841 = False

			}

			if True == ifres91841 {
				tmp91810 := MakeNative(func(__e *ControlFlow) {
					PackageNameDot := __e.Get(1)
					_ = PackageNameDot
					tmp91811 := MakeNative(func(__e *ControlFlow) {
						ExpPackageNameDot := __e.Get(1)
						_ = ExpPackageNameDot
						tmp91812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

						tmp91813 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91814 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91815 := Call(__e, tmp91814, V1946)

						tmp91816 := Call(__e, tmp91813, tmp91815)

						tmp91817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91818 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91819 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91820 := Call(__e, tmp91819, V1946)

						tmp91821 := Call(__e, tmp91818, tmp91820)

						tmp91822 := Call(__e, tmp91817, tmp91821)

						tmp91823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91824 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91825 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91826 := Call(__e, tmp91825, V1946)

						tmp91827 := Call(__e, tmp91824, tmp91826)

						tmp91828 := Call(__e, tmp91823, tmp91827)

						__e.TailApply(tmp91812, tmp91816, tmp91822, tmp91828, ExpPackageNameDot)
						return

					}, 1)

					tmp91829 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

					tmp91830 := Call(__e, tmp91829, PackageNameDot)

					__e.TailApply(tmp91811, tmp91830)
					return

				}, 1)

				tmp91831 := Call(__e, PrimNS1Value(symns2_1value), symintern)

				tmp91832 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp91833 := Call(__e, PrimNS1Value(symns2_1value), symstr)

				tmp91834 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91836 := Call(__e, tmp91835, V1946)

				tmp91837 := Call(__e, tmp91834, tmp91836)

				tmp91838 := Call(__e, tmp91833, tmp91837)

				tmp91839 := Call(__e, tmp91832, tmp91838, MakeString("."))

				tmp91840 := Call(__e, tmp91831, tmp91839)

				__e.TailApply(tmp91810, tmp91840)
				return

			} else {
				tmp91809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp91809, symshen_4package_1contents)
				return

			}

		}

	}, 1)

	tmp91898 := Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1contents, tmp91806)

	_ = tmp91898

	tmp91899 := MakeNative(func(__e *ControlFlow) {
		V1949 := __e.Get(1)
		_ = V1949
		V1950 := __e.Get(2)
		_ = V1950
		tmp91905 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91906 := Call(__e, tmp91905, V1950)

		if True == tmp91906 {
			tmp91901 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp91902 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				tmp91903 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

				__e.TailApply(tmp91903, V1949, Z)
				return

			}, 1)

			tmp91904 := Call(__e, tmp91901, tmp91902, V1950)

			__e.TailApply(V1949, tmp91904)
			return

		} else {
			__e.TailApply(V1949, V1950)
			return
		}

	}, 2)

	tmp91907 := Call(__e, PrimNS1Value(symns2_1set), symshen_4walk, tmp91899)

	_ = tmp91907

	tmp91908 := MakeNative(func(__e *ControlFlow) {
		V1954 := __e.Get(1)
		_ = V1954
		V1955 := __e.Get(2)
		_ = V1955
		V1956 := __e.Get(3)
		_ = V1956
		tmp91909 := MakeNative(func(__e *ControlFlow) {
			O := __e.Get(1)
			_ = O
			tmp91920 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91921 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp91922 := Call(__e, tmp91921)

			tmp91923 := Call(__e, tmp91920, tmp91922, O)

			var ifres91912 Obj

			if True == tmp91923 {
				ifres91912 = True

			} else {
				tmp91914 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp91915 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				tmp91916 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91917 := Call(__e, tmp91916, O)

				tmp91918 := Call(__e, tmp91915, tmp91917)

				tmp91919 := Call(__e, tmp91914, tmp91918)

				var ifres91913 Obj

				if True == tmp91919 {
					ifres91913 = True

				} else {
					ifres91913 = False

				}

				ifres91912 = ifres91913

			}

			if True == ifres91912 {
				__e.TailApply(V1956, O)
				return
			} else {
				tmp91911 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				__e.TailApply(tmp91911, O)
				return

			}

		}, 1)

		tmp91924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91926 := Call(__e, tmp91925, Nil, Nil)

		tmp91927 := Call(__e, tmp91924, V1955, tmp91926)

		tmp91928 := Call(__e, V1954, tmp91927)

		__e.TailApply(tmp91909, tmp91928)
		return

	}, 3)

	tmp91929 := Call(__e, PrimNS1Value(symns2_1set), symcompile, tmp91908)

	_ = tmp91929

	tmp91930 := MakeNative(func(__e *ControlFlow) {
		V1959 := __e.Get(1)
		_ = V1959
		V1960 := __e.Get(2)
		_ = V1960
		tmp91933 := Call(__e, V1959, V1960)

		if True == tmp91933 {
			tmp91932 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp91932)
			return

		} else {
			__e.Return(V1960)
			return
		}

	}, 2)

	tmp91934 := Call(__e, PrimNS1Value(symns2_1set), symfail_1if, tmp91930)

	_ = tmp91934

	tmp91935 := MakeNative(func(__e *ControlFlow) {
		V1963 := __e.Get(1)
		_ = V1963
		V1964 := __e.Get(2)
		_ = V1964
		tmp91936 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		__e.TailApply(tmp91936, V1963, V1964)
		return

	}, 2)

	tmp91937 := Call(__e, PrimNS1Value(symns2_1set), sym_8s, tmp91935)

	_ = tmp91937

	tmp91938 := MakeNative(func(__e *ControlFlow) {
		tmp91939 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp91939, symshen_4_dtc_d)
		return

	}, 0)

	tmp91940 := Call(__e, PrimNS1Value(symns2_1set), symtc_2, tmp91938)

	_ = tmp91940

	tmp91941 := MakeNative(func(__e *ControlFlow) {
		V1966 := __e.Get(1)
		_ = V1966
		tmp91942 := MakeNative(func(__e *ControlFlow) {
			tmp91943 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp91944 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91945 := Call(__e, tmp91944, sym_dproperty_1vector_d)

			__e.TailApply(tmp91943, V1966, symshen_4source, tmp91945)
			return

		}, 0)

		tmp91946 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp91947 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp91948 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp91949 := Call(__e, tmp91948, V1966, MakeString(" not found.\n"), symshen_4a)

			__e.TailApply(tmp91947, tmp91949)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp91942, tmp91946)
		return

	}, 1)

	tmp91950 := Call(__e, PrimNS1Value(symns2_1set), symps, tmp91941)

	_ = tmp91950

	tmp91951 := MakeNative(func(__e *ControlFlow) {
		tmp91952 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp91952, sym_dstinput_d)
		return

	}, 0)

	tmp91953 := Call(__e, PrimNS1Value(symns2_1set), symstinput, tmp91951)

	_ = tmp91953

	tmp91954 := MakeNative(func(__e *ControlFlow) {
		V1968 := __e.Get(1)
		_ = V1968
		tmp91955 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp91956 := MakeNative(func(__e *ControlFlow) {
				ZeroStamp := __e.Get(1)
				_ = ZeroStamp
				tmp91957 := MakeNative(func(__e *ControlFlow) {
					Standard := __e.Get(1)
					_ = Standard
					__e.Return(Standard)
					return
				}, 1)

				tmp91963 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91964 := Call(__e, tmp91963, V1968, MakeNumber(0))

				var ifres91958 Obj

				if True == tmp91964 {
					ifres91958 = ZeroStamp

				} else {
					tmp91959 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

					tmp91960 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp91961 := Call(__e, tmp91960)

					tmp91962 := Call(__e, tmp91959, ZeroStamp, MakeNumber(1), V1968, tmp91961)

					ifres91958 = tmp91962

				}

				__e.TailApply(tmp91957, ifres91958)
				return

			}, 1)

			tmp91965 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp91966 := Call(__e, tmp91965, Vector, MakeNumber(0), V1968)

			__e.TailApply(tmp91956, tmp91966)
			return

		}, 1)

		tmp91967 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		tmp91968 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp91969 := Call(__e, tmp91968, V1968, MakeNumber(1))

		tmp91970 := Call(__e, tmp91967, tmp91969)

		__e.TailApply(tmp91955, tmp91970)
		return

	}, 1)

	tmp91971 := Call(__e, PrimNS1Value(symns2_1set), symvector, tmp91954)

	_ = tmp91971

	tmp91972 := MakeNative(func(__e *ControlFlow) {
		V1974 := __e.Get(1)
		_ = V1974
		V1975 := __e.Get(2)
		_ = V1975
		V1976 := __e.Get(3)
		_ = V1976
		V1977 := __e.Get(4)
		_ = V1977
		tmp91980 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp91981 := Call(__e, tmp91980, V1976, V1975)

		if True == tmp91981 {
			tmp91979 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			__e.TailApply(tmp91979, V1974, V1976, V1977)
			return

		} else {
			tmp91974 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

			tmp91975 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp91976 := Call(__e, tmp91975, V1974, V1975, V1977)

			tmp91977 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp91978 := Call(__e, tmp91977, MakeNumber(1), V1975)

			__e.TailApply(tmp91974, tmp91976, tmp91978, V1976, V1977)
			return

		}

	}, 4)

	tmp91982 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fillvector, tmp91972)

	_ = tmp91982

	tmp91983 := MakeNative(func(__e *ControlFlow) {
		V1979 := __e.Get(1)
		_ = V1979
		tmp91998 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		tmp91999 := Call(__e, tmp91998, V1979)

		if True == tmp91999 {
			tmp91986 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp91991 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

				tmp91992 := Call(__e, tmp91991, X)

				if True == tmp91992 {
					tmp91989 := Call(__e, PrimNS1Value(symns2_1value), sym_6_a)

					tmp91990 := Call(__e, tmp91989, X, MakeNumber(0))

					if True == tmp91990 {
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

			tmp91993 := MakeNative(func(__e *ControlFlow) {
				tmp91994 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(tmp91994, V1979, MakeNumber(0))
				return

			}, 0)

			tmp91995 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeNumber(-1))
				return
			}, 1)

			tmp91996 := Call(__e, PrimNS1Value(symtry_1catch), tmp91993, tmp91995)

			tmp91997 := Call(__e, tmp91986, tmp91996)

			if True == tmp91997 {
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

	tmp92000 := Call(__e, PrimNS1Value(symns2_1set), symvector_2, tmp91983)

	_ = tmp92000

	tmp92001 := MakeNative(func(__e *ControlFlow) {
		V1983 := __e.Get(1)
		_ = V1983
		V1984 := __e.Get(2)
		_ = V1984
		V1985 := __e.Get(3)
		_ = V1985
		tmp92005 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92006 := Call(__e, tmp92005, V1984, MakeNumber(0))

		if True == tmp92006 {
			tmp92004 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp92004, MakeString("cannot access 0th element of a vector\n"))
			return

		} else {
			tmp92003 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			__e.TailApply(tmp92003, V1983, V1984, V1985)
			return

		}

	}, 3)

	tmp92007 := Call(__e, PrimNS1Value(symns2_1set), symvector_1_6, tmp92001)

	_ = tmp92007

	tmp92008 := MakeNative(func(__e *ControlFlow) {
		V1988 := __e.Get(1)
		_ = V1988
		V1989 := __e.Get(2)
		_ = V1989
		tmp92020 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92021 := Call(__e, tmp92020, V1989, MakeNumber(0))

		if True == tmp92021 {
			tmp92019 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp92019, MakeString("cannot access 0th element of a vector\n"))
			return

		} else {
			tmp92010 := MakeNative(func(__e *ControlFlow) {
				VectorElement := __e.Get(1)
				_ = VectorElement
				tmp92013 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp92014 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				tmp92015 := Call(__e, tmp92014)

				tmp92016 := Call(__e, tmp92013, VectorElement, tmp92015)

				if True == tmp92016 {
					tmp92012 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(tmp92012, MakeString("vector element not found\n"))
					return

				} else {
					__e.Return(VectorElement)
					return
				}

			}, 1)

			tmp92017 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			tmp92018 := Call(__e, tmp92017, V1988, V1989)

			__e.TailApply(tmp92010, tmp92018)
			return

		}

	}, 2)

	tmp92022 := Call(__e, PrimNS1Value(symns2_1set), sym_5_1vector, tmp92008)

	_ = tmp92022

	tmp92023 := MakeNative(func(__e *ControlFlow) {
		V1991 := __e.Get(1)
		_ = V1991
		tmp92028 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

		tmp92029 := Call(__e, tmp92028, V1991)

		if True == tmp92029 {
			tmp92026 := Call(__e, PrimNS1Value(symns2_1value), sym_6_a)

			tmp92027 := Call(__e, tmp92026, V1991, MakeNumber(0))

			if True == tmp92027 {
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

	tmp92030 := Call(__e, PrimNS1Value(symns2_1set), symshen_4posint_2, tmp92023)

	_ = tmp92030

	tmp92031 := MakeNative(func(__e *ControlFlow) {
		V1993 := __e.Get(1)
		_ = V1993
		tmp92032 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(tmp92032, V1993, MakeNumber(0))
		return

	}, 1)

	tmp92033 := Call(__e, PrimNS1Value(symns2_1set), symlimit, tmp92031)

	_ = tmp92033

	tmp92034 := MakeNative(func(__e *ControlFlow) {
		V1995 := __e.Get(1)
		_ = V1995
		tmp92050 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

		tmp92051 := Call(__e, tmp92050, V1995)

		var ifres92042 Obj

		if True == tmp92051 {
			ifres92042 = True

		} else {
			tmp92048 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

			tmp92049 := Call(__e, tmp92048, V1995)

			var ifres92044 Obj

			if True == tmp92049 {
				ifres92044 = True

			} else {
				tmp92046 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				tmp92047 := Call(__e, tmp92046, V1995)

				var ifres92045 Obj

				if True == tmp92047 {
					ifres92045 = True

				} else {
					ifres92045 = False

				}

				ifres92044 = ifres92045

			}

			var ifres92043 Obj

			if True == ifres92044 {
				ifres92043 = True

			} else {
				ifres92043 = False

			}

			ifres92042 = ifres92043

		}

		if True == ifres92042 {
			__e.Return(False)
			return
		} else {
			tmp92036 := MakeNative(func(__e *ControlFlow) {
				tmp92037 := MakeNative(func(__e *ControlFlow) {
					String := __e.Get(1)
					_ = String
					tmp92038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4analyse_1symbol_2)

					__e.TailApply(tmp92038, String)
					return

				}, 1)

				tmp92039 := Call(__e, PrimNS1Value(symns2_1value), symstr)

				tmp92040 := Call(__e, tmp92039, V1995)

				__e.TailApply(tmp92037, tmp92040)
				return

			}, 0)

			tmp92041 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp92036, tmp92041)
			return

		}

	}, 1)

	tmp92052 := Call(__e, PrimNS1Value(symns2_1set), symsymbol_2, tmp92034)

	_ = tmp92052

	tmp92053 := MakeNative(func(__e *ControlFlow) {
		V1997 := __e.Get(1)
		_ = V1997
		tmp92069 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92070 := Call(__e, tmp92069, MakeString(""), V1997)

		if True == tmp92070 {
			__e.Return(False)
			return
		} else {
			tmp92067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp92068 := Call(__e, tmp92067, V1997)

			if True == tmp92068 {
				tmp92063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alpha_2)

				tmp92064 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp92065 := Call(__e, tmp92064, V1997, MakeNumber(0))

				tmp92066 := Call(__e, tmp92063, tmp92065)

				if True == tmp92066 {
					tmp92059 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

					tmp92060 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp92061 := Call(__e, tmp92060, V1997)

					tmp92062 := Call(__e, tmp92059, tmp92061)

					if True == tmp92062 {
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
				tmp92056 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92056, symshen_4analyse_1symbol_2)
				return

			}

		}

	}, 1)

	tmp92071 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1symbol_2, tmp92053)

	_ = tmp92071

	tmp92072 := MakeNative(func(__e *ControlFlow) {
		V1999 := __e.Get(1)
		_ = V1999
		tmp92073 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp92074 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92075 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92076 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92077 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92078 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92081 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92091 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92093 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92094 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92095 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92096 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92097 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92101 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92102 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92105 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92106 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92107 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92111 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92116 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92117 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92118 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92123 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92124 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92125 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92126 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92134 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92135 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92138 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92141 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92142 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92143 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92144 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92146 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92147 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92149 := Call(__e, tmp92148, MakeString("."), Nil)

		tmp92150 := Call(__e, tmp92147, MakeString("'"), tmp92149)

		tmp92151 := Call(__e, tmp92146, MakeString("#"), tmp92150)

		tmp92152 := Call(__e, tmp92145, MakeString("`"), tmp92151)

		tmp92153 := Call(__e, tmp92144, MakeString(";"), tmp92152)

		tmp92154 := Call(__e, tmp92143, MakeString(":"), tmp92153)

		tmp92155 := Call(__e, tmp92142, MakeString("}"), tmp92154)

		tmp92156 := Call(__e, tmp92141, MakeString("{"), tmp92155)

		tmp92157 := Call(__e, tmp92140, MakeString("%"), tmp92156)

		tmp92158 := Call(__e, tmp92139, MakeString("&"), tmp92157)

		tmp92159 := Call(__e, tmp92138, MakeString("<"), tmp92158)

		tmp92160 := Call(__e, tmp92137, MakeString(">"), tmp92159)

		tmp92161 := Call(__e, tmp92136, MakeString("~"), tmp92160)

		tmp92162 := Call(__e, tmp92135, MakeString("@"), tmp92161)

		tmp92163 := Call(__e, tmp92134, MakeString("!"), tmp92162)

		tmp92164 := Call(__e, tmp92133, MakeString("$"), tmp92163)

		tmp92165 := Call(__e, tmp92132, MakeString("?"), tmp92164)

		tmp92166 := Call(__e, tmp92131, MakeString("_"), tmp92165)

		tmp92167 := Call(__e, tmp92130, MakeString("-"), tmp92166)

		tmp92168 := Call(__e, tmp92129, MakeString("+"), tmp92167)

		tmp92169 := Call(__e, tmp92128, MakeString("/"), tmp92168)

		tmp92170 := Call(__e, tmp92127, MakeString("*"), tmp92169)

		tmp92171 := Call(__e, tmp92126, MakeString("="), tmp92170)

		tmp92172 := Call(__e, tmp92125, MakeString("z"), tmp92171)

		tmp92173 := Call(__e, tmp92124, MakeString("y"), tmp92172)

		tmp92174 := Call(__e, tmp92123, MakeString("x"), tmp92173)

		tmp92175 := Call(__e, tmp92122, MakeString("w"), tmp92174)

		tmp92176 := Call(__e, tmp92121, MakeString("v"), tmp92175)

		tmp92177 := Call(__e, tmp92120, MakeString("u"), tmp92176)

		tmp92178 := Call(__e, tmp92119, MakeString("t"), tmp92177)

		tmp92179 := Call(__e, tmp92118, MakeString("s"), tmp92178)

		tmp92180 := Call(__e, tmp92117, MakeString("r"), tmp92179)

		tmp92181 := Call(__e, tmp92116, MakeString("q"), tmp92180)

		tmp92182 := Call(__e, tmp92115, MakeString("p"), tmp92181)

		tmp92183 := Call(__e, tmp92114, MakeString("o"), tmp92182)

		tmp92184 := Call(__e, tmp92113, MakeString("n"), tmp92183)

		tmp92185 := Call(__e, tmp92112, MakeString("m"), tmp92184)

		tmp92186 := Call(__e, tmp92111, MakeString("l"), tmp92185)

		tmp92187 := Call(__e, tmp92110, MakeString("k"), tmp92186)

		tmp92188 := Call(__e, tmp92109, MakeString("j"), tmp92187)

		tmp92189 := Call(__e, tmp92108, MakeString("i"), tmp92188)

		tmp92190 := Call(__e, tmp92107, MakeString("h"), tmp92189)

		tmp92191 := Call(__e, tmp92106, MakeString("g"), tmp92190)

		tmp92192 := Call(__e, tmp92105, MakeString("f"), tmp92191)

		tmp92193 := Call(__e, tmp92104, MakeString("e"), tmp92192)

		tmp92194 := Call(__e, tmp92103, MakeString("d"), tmp92193)

		tmp92195 := Call(__e, tmp92102, MakeString("c"), tmp92194)

		tmp92196 := Call(__e, tmp92101, MakeString("b"), tmp92195)

		tmp92197 := Call(__e, tmp92100, MakeString("a"), tmp92196)

		tmp92198 := Call(__e, tmp92099, MakeString("Z"), tmp92197)

		tmp92199 := Call(__e, tmp92098, MakeString("Y"), tmp92198)

		tmp92200 := Call(__e, tmp92097, MakeString("X"), tmp92199)

		tmp92201 := Call(__e, tmp92096, MakeString("W"), tmp92200)

		tmp92202 := Call(__e, tmp92095, MakeString("V"), tmp92201)

		tmp92203 := Call(__e, tmp92094, MakeString("U"), tmp92202)

		tmp92204 := Call(__e, tmp92093, MakeString("T"), tmp92203)

		tmp92205 := Call(__e, tmp92092, MakeString("S"), tmp92204)

		tmp92206 := Call(__e, tmp92091, MakeString("R"), tmp92205)

		tmp92207 := Call(__e, tmp92090, MakeString("Q"), tmp92206)

		tmp92208 := Call(__e, tmp92089, MakeString("P"), tmp92207)

		tmp92209 := Call(__e, tmp92088, MakeString("O"), tmp92208)

		tmp92210 := Call(__e, tmp92087, MakeString("N"), tmp92209)

		tmp92211 := Call(__e, tmp92086, MakeString("M"), tmp92210)

		tmp92212 := Call(__e, tmp92085, MakeString("L"), tmp92211)

		tmp92213 := Call(__e, tmp92084, MakeString("K"), tmp92212)

		tmp92214 := Call(__e, tmp92083, MakeString("J"), tmp92213)

		tmp92215 := Call(__e, tmp92082, MakeString("I"), tmp92214)

		tmp92216 := Call(__e, tmp92081, MakeString("H"), tmp92215)

		tmp92217 := Call(__e, tmp92080, MakeString("G"), tmp92216)

		tmp92218 := Call(__e, tmp92079, MakeString("F"), tmp92217)

		tmp92219 := Call(__e, tmp92078, MakeString("E"), tmp92218)

		tmp92220 := Call(__e, tmp92077, MakeString("D"), tmp92219)

		tmp92221 := Call(__e, tmp92076, MakeString("C"), tmp92220)

		tmp92222 := Call(__e, tmp92075, MakeString("B"), tmp92221)

		tmp92223 := Call(__e, tmp92074, MakeString("A"), tmp92222)

		__e.TailApply(tmp92073, V1999, tmp92223)
		return

	}, 1)

	tmp92224 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alpha_2, tmp92072)

	_ = tmp92224

	tmp92225 := MakeNative(func(__e *ControlFlow) {
		V2001 := __e.Get(1)
		_ = V2001
		tmp92241 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92242 := Call(__e, tmp92241, MakeString(""), V2001)

		if True == tmp92242 {
			__e.Return(True)
			return
		} else {
			tmp92239 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp92240 := Call(__e, tmp92239, V2001)

			if True == tmp92240 {
				tmp92235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanum_2)

				tmp92236 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp92237 := Call(__e, tmp92236, V2001, MakeNumber(0))

				tmp92238 := Call(__e, tmp92235, tmp92237)

				if True == tmp92238 {
					tmp92231 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

					tmp92232 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp92233 := Call(__e, tmp92232, V2001)

					tmp92234 := Call(__e, tmp92231, tmp92233)

					if True == tmp92234 {
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
				tmp92228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92228, symshen_4alphanums_2)
				return

			}

		}

	}, 1)

	tmp92243 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanums_2, tmp92225)

	_ = tmp92243

	tmp92244 := MakeNative(func(__e *ControlFlow) {
		V2003 := __e.Get(1)
		_ = V2003
		tmp92249 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alpha_2)

		tmp92250 := Call(__e, tmp92249, V2003)

		if True == tmp92250 {
			__e.Return(True)
			return
		} else {
			tmp92247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digit_2)

			tmp92248 := Call(__e, tmp92247, V2003)

			if True == tmp92248 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp92251 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanum_2, tmp92244)

	_ = tmp92251

	tmp92252 := MakeNative(func(__e *ControlFlow) {
		V2005 := __e.Get(1)
		_ = V2005
		tmp92253 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp92254 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92255 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92256 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92257 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92264 := Call(__e, tmp92263, MakeString("0"), Nil)

		tmp92265 := Call(__e, tmp92262, MakeString("9"), tmp92264)

		tmp92266 := Call(__e, tmp92261, MakeString("8"), tmp92265)

		tmp92267 := Call(__e, tmp92260, MakeString("7"), tmp92266)

		tmp92268 := Call(__e, tmp92259, MakeString("6"), tmp92267)

		tmp92269 := Call(__e, tmp92258, MakeString("5"), tmp92268)

		tmp92270 := Call(__e, tmp92257, MakeString("4"), tmp92269)

		tmp92271 := Call(__e, tmp92256, MakeString("3"), tmp92270)

		tmp92272 := Call(__e, tmp92255, MakeString("2"), tmp92271)

		tmp92273 := Call(__e, tmp92254, MakeString("1"), tmp92272)

		__e.TailApply(tmp92253, V2005, tmp92273)
		return

	}, 1)

	tmp92274 := Call(__e, PrimNS1Value(symns2_1set), symshen_4digit_2, tmp92252)

	_ = tmp92274

	tmp92275 := MakeNative(func(__e *ControlFlow) {
		V2007 := __e.Get(1)
		_ = V2007
		tmp92291 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

		tmp92292 := Call(__e, tmp92291, V2007)

		var ifres92283 Obj

		if True == tmp92292 {
			ifres92283 = True

		} else {
			tmp92289 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

			tmp92290 := Call(__e, tmp92289, V2007)

			var ifres92285 Obj

			if True == tmp92290 {
				ifres92285 = True

			} else {
				tmp92287 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				tmp92288 := Call(__e, tmp92287, V2007)

				var ifres92286 Obj

				if True == tmp92288 {
					ifres92286 = True

				} else {
					ifres92286 = False

				}

				ifres92285 = ifres92286

			}

			var ifres92284 Obj

			if True == ifres92285 {
				ifres92284 = True

			} else {
				ifres92284 = False

			}

			ifres92283 = ifres92284

		}

		if True == ifres92283 {
			__e.Return(False)
			return
		} else {
			tmp92277 := MakeNative(func(__e *ControlFlow) {
				tmp92278 := MakeNative(func(__e *ControlFlow) {
					String := __e.Get(1)
					_ = String
					tmp92279 := Call(__e, PrimNS1Value(symns2_1value), symshen_4analyse_1variable_2)

					__e.TailApply(tmp92279, String)
					return

				}, 1)

				tmp92280 := Call(__e, PrimNS1Value(symns2_1value), symstr)

				tmp92281 := Call(__e, tmp92280, V2007)

				__e.TailApply(tmp92278, tmp92281)
				return

			}, 0)

			tmp92282 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp92277, tmp92282)
			return

		}

	}, 1)

	tmp92293 := Call(__e, PrimNS1Value(symns2_1set), symvariable_2, tmp92275)

	_ = tmp92293

	tmp92294 := MakeNative(func(__e *ControlFlow) {
		V2009 := __e.Get(1)
		_ = V2009
		tmp92307 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

		tmp92308 := Call(__e, tmp92307, V2009)

		if True == tmp92308 {
			tmp92303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4uppercase_2)

			tmp92304 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			tmp92305 := Call(__e, tmp92304, V2009, MakeNumber(0))

			tmp92306 := Call(__e, tmp92303, tmp92305)

			if True == tmp92306 {
				tmp92299 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

				tmp92300 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp92301 := Call(__e, tmp92300, V2009)

				tmp92302 := Call(__e, tmp92299, tmp92301)

				if True == tmp92302 {
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
			tmp92296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp92296, symshen_4analyse_1variable_2)
			return

		}

	}, 1)

	tmp92309 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1variable_2, tmp92294)

	_ = tmp92309

	tmp92310 := MakeNative(func(__e *ControlFlow) {
		V2011 := __e.Get(1)
		_ = V2011
		tmp92311 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp92312 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92314 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92328 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92330 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92331 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92338 := Call(__e, tmp92337, MakeString("Z"), Nil)

		tmp92339 := Call(__e, tmp92336, MakeString("Y"), tmp92338)

		tmp92340 := Call(__e, tmp92335, MakeString("X"), tmp92339)

		tmp92341 := Call(__e, tmp92334, MakeString("W"), tmp92340)

		tmp92342 := Call(__e, tmp92333, MakeString("V"), tmp92341)

		tmp92343 := Call(__e, tmp92332, MakeString("U"), tmp92342)

		tmp92344 := Call(__e, tmp92331, MakeString("T"), tmp92343)

		tmp92345 := Call(__e, tmp92330, MakeString("S"), tmp92344)

		tmp92346 := Call(__e, tmp92329, MakeString("R"), tmp92345)

		tmp92347 := Call(__e, tmp92328, MakeString("Q"), tmp92346)

		tmp92348 := Call(__e, tmp92327, MakeString("P"), tmp92347)

		tmp92349 := Call(__e, tmp92326, MakeString("O"), tmp92348)

		tmp92350 := Call(__e, tmp92325, MakeString("N"), tmp92349)

		tmp92351 := Call(__e, tmp92324, MakeString("M"), tmp92350)

		tmp92352 := Call(__e, tmp92323, MakeString("L"), tmp92351)

		tmp92353 := Call(__e, tmp92322, MakeString("K"), tmp92352)

		tmp92354 := Call(__e, tmp92321, MakeString("J"), tmp92353)

		tmp92355 := Call(__e, tmp92320, MakeString("I"), tmp92354)

		tmp92356 := Call(__e, tmp92319, MakeString("H"), tmp92355)

		tmp92357 := Call(__e, tmp92318, MakeString("G"), tmp92356)

		tmp92358 := Call(__e, tmp92317, MakeString("F"), tmp92357)

		tmp92359 := Call(__e, tmp92316, MakeString("E"), tmp92358)

		tmp92360 := Call(__e, tmp92315, MakeString("D"), tmp92359)

		tmp92361 := Call(__e, tmp92314, MakeString("C"), tmp92360)

		tmp92362 := Call(__e, tmp92313, MakeString("B"), tmp92361)

		tmp92363 := Call(__e, tmp92312, MakeString("A"), tmp92362)

		__e.TailApply(tmp92311, V2011, tmp92363)
		return

	}, 1)

	tmp92364 := Call(__e, PrimNS1Value(symns2_1set), symshen_4uppercase_2, tmp92310)

	_ = tmp92364

	tmp92365 := MakeNative(func(__e *ControlFlow) {
		V2013 := __e.Get(1)
		_ = V2013
		tmp92366 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		tmp92367 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp92368 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		tmp92369 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp92370 := Call(__e, tmp92369, symshen_4_dgensym_d)

		tmp92371 := Call(__e, tmp92368, MakeNumber(1), tmp92370)

		tmp92372 := Call(__e, tmp92367, symshen_4_dgensym_d, tmp92371)

		__e.TailApply(tmp92366, V2013, tmp92372)
		return

	}, 1)

	tmp92373 := Call(__e, PrimNS1Value(symns2_1set), symgensym, tmp92365)

	_ = tmp92373

	tmp92374 := MakeNative(func(__e *ControlFlow) {
		V2016 := __e.Get(1)
		_ = V2016
		V2017 := __e.Get(2)
		_ = V2017
		tmp92375 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp92376 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp92377 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		tmp92378 := Call(__e, tmp92377, V2016)

		tmp92379 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		tmp92380 := Call(__e, tmp92379, V2017)

		tmp92381 := Call(__e, tmp92376, tmp92378, tmp92380)

		__e.TailApply(tmp92375, tmp92381)
		return

	}, 2)

	tmp92382 := Call(__e, PrimNS1Value(symns2_1set), symconcat, tmp92374)

	_ = tmp92382

	tmp92383 := MakeNative(func(__e *ControlFlow) {
		V2020 := __e.Get(1)
		_ = V2020
		V2021 := __e.Get(2)
		_ = V2021
		tmp92384 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp92385 := MakeNative(func(__e *ControlFlow) {
				Tag := __e.Get(1)
				_ = Tag
				tmp92386 := MakeNative(func(__e *ControlFlow) {
					Fst := __e.Get(1)
					_ = Fst
					tmp92387 := MakeNative(func(__e *ControlFlow) {
						Snd := __e.Get(1)
						_ = Snd
						__e.Return(Vector)
						return
					}, 1)

					tmp92388 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

					tmp92389 := Call(__e, tmp92388, Vector, MakeNumber(2), V2021)

					__e.TailApply(tmp92387, tmp92389)
					return

				}, 1)

				tmp92390 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

				tmp92391 := Call(__e, tmp92390, Vector, MakeNumber(1), V2020)

				__e.TailApply(tmp92386, tmp92391)
				return

			}, 1)

			tmp92392 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			tmp92393 := Call(__e, tmp92392, Vector, MakeNumber(0), symshen_4tuple)

			__e.TailApply(tmp92385, tmp92393)
			return

		}, 1)

		tmp92394 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		tmp92395 := Call(__e, tmp92394, MakeNumber(3))

		__e.TailApply(tmp92384, tmp92395)
		return

	}, 2)

	tmp92396 := Call(__e, PrimNS1Value(symns2_1set), sym_8p, tmp92383)

	_ = tmp92396

	tmp92397 := MakeNative(func(__e *ControlFlow) {
		V2023 := __e.Get(1)
		_ = V2023
		tmp92398 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(tmp92398, V2023, MakeNumber(1))
		return

	}, 1)

	tmp92399 := Call(__e, PrimNS1Value(symns2_1set), symfst, tmp92397)

	_ = tmp92399

	tmp92400 := MakeNative(func(__e *ControlFlow) {
		V2025 := __e.Get(1)
		_ = V2025
		tmp92401 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(tmp92401, V2025, MakeNumber(2))
		return

	}, 1)

	tmp92402 := Call(__e, PrimNS1Value(symns2_1set), symsnd, tmp92400)

	_ = tmp92402

	tmp92403 := MakeNative(func(__e *ControlFlow) {
		V2027 := __e.Get(1)
		_ = V2027
		tmp92412 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		tmp92413 := Call(__e, tmp92412, V2027)

		if True == tmp92413 {
			tmp92406 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp92407 := MakeNative(func(__e *ControlFlow) {
				tmp92408 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(tmp92408, V2027, MakeNumber(0))
				return

			}, 0)

			tmp92409 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1tuple)
				return
			}, 1)

			tmp92410 := Call(__e, PrimNS1Value(symtry_1catch), tmp92407, tmp92409)

			tmp92411 := Call(__e, tmp92406, symshen_4tuple, tmp92410)

			if True == tmp92411 {
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

	tmp92414 := Call(__e, PrimNS1Value(symns2_1set), symtuple_2, tmp92403)

	_ = tmp92414

	tmp92415 := MakeNative(func(__e *ControlFlow) {
		V2030 := __e.Get(1)
		_ = V2030
		V2031 := __e.Get(2)
		_ = V2031
		tmp92428 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92429 := Call(__e, tmp92428, Nil, V2030)

		if True == tmp92429 {
			__e.Return(V2031)
			return
		} else {
			tmp92426 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92427 := Call(__e, tmp92426, V2030)

			if True == tmp92427 {
				tmp92419 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92420 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92421 := Call(__e, tmp92420, V2030)

				tmp92422 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp92423 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp92424 := Call(__e, tmp92423, V2030)

				tmp92425 := Call(__e, tmp92422, tmp92424, V2031)

				__e.TailApply(tmp92419, tmp92421, tmp92425)
				return

			} else {
				tmp92418 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92418, symappend)
				return

			}

		}

	}, 2)

	tmp92430 := Call(__e, PrimNS1Value(symns2_1set), symappend, tmp92415)

	_ = tmp92430

	tmp92431 := MakeNative(func(__e *ControlFlow) {
		V2034 := __e.Get(1)
		_ = V2034
		V2035 := __e.Get(2)
		_ = V2035
		tmp92432 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp92433 := MakeNative(func(__e *ControlFlow) {
				NewVector := __e.Get(1)
				_ = NewVector
				tmp92434 := MakeNative(func(__e *ControlFlow) {
					X_7NewVector := __e.Get(1)
					_ = X_7NewVector
					tmp92437 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp92438 := Call(__e, tmp92437, Limit, MakeNumber(0))

					if True == tmp92438 {
						__e.Return(X_7NewVector)
						return
					} else {
						tmp92436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8v_1help)

						__e.TailApply(tmp92436, V2035, MakeNumber(1), Limit, X_7NewVector)
						return

					}

				}, 1)

				tmp92439 := Call(__e, PrimNS1Value(symns2_1value), symvector_1_6)

				tmp92440 := Call(__e, tmp92439, NewVector, MakeNumber(1), V2034)

				__e.TailApply(tmp92434, tmp92440)
				return

			}, 1)

			tmp92441 := Call(__e, PrimNS1Value(symns2_1value), symvector)

			tmp92442 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp92443 := Call(__e, tmp92442, Limit, MakeNumber(1))

			tmp92444 := Call(__e, tmp92441, tmp92443)

			__e.TailApply(tmp92433, tmp92444)
			return

		}, 1)

		tmp92445 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

		tmp92446 := Call(__e, tmp92445, V2035)

		__e.TailApply(tmp92432, tmp92446)
		return

	}, 2)

	tmp92447 := Call(__e, PrimNS1Value(symns2_1set), sym_8v, tmp92431)

	_ = tmp92447

	tmp92448 := MakeNative(func(__e *ControlFlow) {
		V2041 := __e.Get(1)
		_ = V2041
		V2042 := __e.Get(2)
		_ = V2042
		V2043 := __e.Get(3)
		_ = V2043
		V2044 := __e.Get(4)
		_ = V2044
		tmp92460 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92461 := Call(__e, tmp92460, V2043, V2042)

		if True == tmp92461 {
			tmp92457 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			tmp92458 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp92459 := Call(__e, tmp92458, V2043, MakeNumber(1))

			__e.TailApply(tmp92457, V2041, V2044, V2043, tmp92459)
			return

		} else {
			tmp92450 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8v_1help)

			tmp92451 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp92452 := Call(__e, tmp92451, V2042, MakeNumber(1))

			tmp92453 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			tmp92454 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp92455 := Call(__e, tmp92454, V2042, MakeNumber(1))

			tmp92456 := Call(__e, tmp92453, V2041, V2044, V2042, tmp92455)

			__e.TailApply(tmp92450, V2041, tmp92452, V2043, tmp92456)
			return

		}

	}, 4)

	tmp92462 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_8v_1help, tmp92448)

	_ = tmp92462

	tmp92463 := MakeNative(func(__e *ControlFlow) {
		V2049 := __e.Get(1)
		_ = V2049
		V2050 := __e.Get(2)
		_ = V2050
		V2051 := __e.Get(3)
		_ = V2051
		V2052 := __e.Get(4)
		_ = V2052
		tmp92464 := MakeNative(func(__e *ControlFlow) {
			tmp92465 := Call(__e, PrimNS1Value(symns2_1value), symvector_1_6)

			tmp92466 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1vector)

			tmp92467 := Call(__e, tmp92466, V2049, V2051)

			__e.TailApply(tmp92465, V2050, V2052, tmp92467)
			return

		}, 0)

		tmp92468 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V2050)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp92464, tmp92468)
		return

	}, 4)

	tmp92469 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copyfromvector, tmp92463)

	_ = tmp92469

	tmp92470 := MakeNative(func(__e *ControlFlow) {
		V2054 := __e.Get(1)
		_ = V2054
		tmp92471 := MakeNative(func(__e *ControlFlow) {
			tmp92472 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1vector)

			__e.TailApply(tmp92472, V2054, MakeNumber(1))
			return

		}, 0)

		tmp92473 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp92474 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp92475 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp92476 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp92477 := Call(__e, tmp92476, V2054, MakeString("\n"), symshen_4s)

			tmp92478 := Call(__e, tmp92475, MakeString("hdv needs a non-empty vector as an argument; not "), tmp92477)

			__e.TailApply(tmp92474, tmp92478)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp92471, tmp92473)
		return

	}, 1)

	tmp92479 := Call(__e, PrimNS1Value(symns2_1set), symhdv, tmp92470)

	_ = tmp92479

	tmp92480 := MakeNative(func(__e *ControlFlow) {
		V2056 := __e.Get(1)
		_ = V2056
		tmp92481 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp92498 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp92499 := Call(__e, tmp92498, Limit, MakeNumber(0))

			if True == tmp92499 {
				tmp92497 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp92497, MakeString("cannot take the tail of the empty vector\n"))
				return

			} else {
				tmp92495 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp92496 := Call(__e, tmp92495, Limit, MakeNumber(1))

				if True == tmp92496 {
					tmp92494 := Call(__e, PrimNS1Value(symns2_1value), symvector)

					__e.TailApply(tmp92494, MakeNumber(0))
					return

				} else {
					tmp92484 := MakeNative(func(__e *ControlFlow) {
						NewVector := __e.Get(1)
						_ = NewVector
						tmp92485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlv_1help)

						tmp92486 := Call(__e, PrimNS1Value(symns2_1value), symvector)

						tmp92487 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						tmp92488 := Call(__e, tmp92487, Limit, MakeNumber(1))

						tmp92489 := Call(__e, tmp92486, tmp92488)

						__e.TailApply(tmp92485, V2056, MakeNumber(2), Limit, tmp92489)
						return

					}, 1)

					tmp92490 := Call(__e, PrimNS1Value(symns2_1value), symvector)

					tmp92491 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					tmp92492 := Call(__e, tmp92491, Limit, MakeNumber(1))

					tmp92493 := Call(__e, tmp92490, tmp92492)

					__e.TailApply(tmp92484, tmp92493)
					return

				}

			}

		}, 1)

		tmp92500 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

		tmp92501 := Call(__e, tmp92500, V2056)

		__e.TailApply(tmp92481, tmp92501)
		return

	}, 1)

	tmp92502 := Call(__e, PrimNS1Value(symns2_1set), symtlv, tmp92480)

	_ = tmp92502

	tmp92503 := MakeNative(func(__e *ControlFlow) {
		V2062 := __e.Get(1)
		_ = V2062
		V2063 := __e.Get(2)
		_ = V2063
		V2064 := __e.Get(3)
		_ = V2064
		V2065 := __e.Get(4)
		_ = V2065
		tmp92515 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92516 := Call(__e, tmp92515, V2064, V2063)

		if True == tmp92516 {
			tmp92512 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			tmp92513 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp92514 := Call(__e, tmp92513, V2064, MakeNumber(1))

			__e.TailApply(tmp92512, V2062, V2065, V2064, tmp92514)
			return

		} else {
			tmp92505 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlv_1help)

			tmp92506 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp92507 := Call(__e, tmp92506, V2063, MakeNumber(1))

			tmp92508 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			tmp92509 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp92510 := Call(__e, tmp92509, V2063, MakeNumber(1))

			tmp92511 := Call(__e, tmp92508, V2062, V2065, V2063, tmp92510)

			__e.TailApply(tmp92505, V2062, tmp92507, V2064, tmp92511)
			return

		}

	}, 4)

	tmp92517 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tlv_1help, tmp92503)

	_ = tmp92517

	tmp92518 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		V2078 := __e.Get(2)
		_ = V2078
		tmp92545 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92546 := Call(__e, tmp92545, Nil, V2078)

		if True == tmp92546 {
			__e.Return(Nil)
			return
		} else {
			tmp92543 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92544 := Call(__e, tmp92543, V2078)

			var ifres92529 Obj

			if True == tmp92544 {
				tmp92539 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92540 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92541 := Call(__e, tmp92540, V2078)

				tmp92542 := Call(__e, tmp92539, tmp92541)

				var ifres92531 Obj

				if True == tmp92542 {
					tmp92533 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp92534 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92536 := Call(__e, tmp92535, V2078)

					tmp92537 := Call(__e, tmp92534, tmp92536)

					tmp92538 := Call(__e, tmp92533, tmp92537, V2077)

					var ifres92532 Obj

					if True == tmp92538 {
						ifres92532 = True

					} else {
						ifres92532 = False

					}

					ifres92531 = ifres92532

				} else {
					ifres92531 = False

				}

				var ifres92530 Obj

				if True == ifres92531 {
					ifres92530 = True

				} else {
					ifres92530 = False

				}

				ifres92529 = ifres92530

			} else {
				ifres92529 = False

			}

			if True == ifres92529 {
				tmp92528 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				__e.TailApply(tmp92528, V2078)
				return

			} else {
				tmp92526 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92527 := Call(__e, tmp92526, V2078)

				if True == tmp92527 {
					tmp92523 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

					tmp92524 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92525 := Call(__e, tmp92524, V2078)

					__e.TailApply(tmp92523, V2077, tmp92525)
					return

				} else {
					tmp92522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp92522, symassoc)
					return

				}

			}

		}

	}, 2)

	tmp92547 := Call(__e, PrimNS1Value(symns2_1set), symassoc, tmp92518)

	_ = tmp92547

	tmp92548 := MakeNative(func(__e *ControlFlow) {
		V2085 := __e.Get(1)
		_ = V2085
		V2086 := __e.Get(2)
		_ = V2086
		V2087 := __e.Get(3)
		_ = V2087
		tmp92590 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92591 := Call(__e, tmp92590, Nil, V2087)

		if True == tmp92591 {
			tmp92587 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp92588 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp92589 := Call(__e, tmp92588, V2085, V2086)

			__e.TailApply(tmp92587, tmp92589, Nil)
			return

		} else {
			tmp92585 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92586 := Call(__e, tmp92585, V2087)

			var ifres92571 Obj

			if True == tmp92586 {
				tmp92581 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92582 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92583 := Call(__e, tmp92582, V2087)

				tmp92584 := Call(__e, tmp92581, tmp92583)

				var ifres92573 Obj

				if True == tmp92584 {
					tmp92575 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp92576 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92577 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92578 := Call(__e, tmp92577, V2087)

					tmp92579 := Call(__e, tmp92576, tmp92578)

					tmp92580 := Call(__e, tmp92575, tmp92579, V2085)

					var ifres92574 Obj

					if True == tmp92580 {
						ifres92574 = True

					} else {
						ifres92574 = False

					}

					ifres92573 = ifres92574

				} else {
					ifres92573 = False

				}

				var ifres92572 Obj

				if True == ifres92573 {
					ifres92572 = True

				} else {
					ifres92572 = False

				}

				ifres92571 = ifres92572

			} else {
				ifres92571 = False

			}

			if True == ifres92571 {
				tmp92562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92563 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92564 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92565 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92566 := Call(__e, tmp92565, V2087)

				tmp92567 := Call(__e, tmp92564, tmp92566)

				tmp92568 := Call(__e, tmp92563, tmp92567, V2086)

				tmp92569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp92570 := Call(__e, tmp92569, V2087)

				__e.TailApply(tmp92562, tmp92568, tmp92570)
				return

			} else {
				tmp92560 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92561 := Call(__e, tmp92560, V2087)

				if True == tmp92561 {
					tmp92553 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp92554 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92555 := Call(__e, tmp92554, V2087)

					tmp92556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

					tmp92557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92558 := Call(__e, tmp92557, V2087)

					tmp92559 := Call(__e, tmp92556, V2085, V2086, tmp92558)

					__e.TailApply(tmp92553, tmp92555, tmp92559)
					return

				} else {
					tmp92552 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp92552, symshen_4assoc_1set)
					return

				}

			}

		}

	}, 3)

	tmp92592 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1set, tmp92548)

	_ = tmp92592

	tmp92593 := MakeNative(func(__e *ControlFlow) {
		V2093 := __e.Get(1)
		_ = V2093
		V2094 := __e.Get(2)
		_ = V2094
		tmp92624 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92625 := Call(__e, tmp92624, Nil, V2094)

		if True == tmp92625 {
			__e.Return(Nil)
			return
		} else {
			tmp92622 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92623 := Call(__e, tmp92622, V2094)

			var ifres92608 Obj

			if True == tmp92623 {
				tmp92618 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92619 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92620 := Call(__e, tmp92619, V2094)

				tmp92621 := Call(__e, tmp92618, tmp92620)

				var ifres92610 Obj

				if True == tmp92621 {
					tmp92612 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp92613 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92614 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92615 := Call(__e, tmp92614, V2094)

					tmp92616 := Call(__e, tmp92613, tmp92615)

					tmp92617 := Call(__e, tmp92612, tmp92616, V2093)

					var ifres92611 Obj

					if True == tmp92617 {
						ifres92611 = True

					} else {
						ifres92611 = False

					}

					ifres92610 = ifres92611

				} else {
					ifres92610 = False

				}

				var ifres92609 Obj

				if True == ifres92610 {
					ifres92609 = True

				} else {
					ifres92609 = False

				}

				ifres92608 = ifres92609

			} else {
				ifres92608 = False

			}

			if True == ifres92608 {
				tmp92607 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				__e.TailApply(tmp92607, V2094)
				return

			} else {
				tmp92605 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92606 := Call(__e, tmp92605, V2094)

				if True == tmp92606 {
					tmp92598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp92599 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92600 := Call(__e, tmp92599, V2094)

					tmp92601 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

					tmp92602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92603 := Call(__e, tmp92602, V2094)

					tmp92604 := Call(__e, tmp92601, V2093, tmp92603)

					__e.TailApply(tmp92598, tmp92600, tmp92604)
					return

				} else {
					tmp92597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp92597, symshen_4assoc_1rm)
					return

				}

			}

		}

	}, 2)

	tmp92626 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1rm, tmp92593)

	_ = tmp92626

	tmp92627 := MakeNative(func(__e *ControlFlow) {
		V2100 := __e.Get(1)
		_ = V2100
		tmp92632 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92633 := Call(__e, tmp92632, True, V2100)

		if True == tmp92633 {
			__e.Return(True)
			return
		} else {
			tmp92630 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp92631 := Call(__e, tmp92630, False, V2100)

			if True == tmp92631 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp92634 := Call(__e, PrimNS1Value(symns2_1set), symboolean_2, tmp92627)

	_ = tmp92634

	tmp92635 := MakeNative(func(__e *ControlFlow) {
		V2102 := __e.Get(1)
		_ = V2102
		tmp92644 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92645 := Call(__e, tmp92644, MakeNumber(0), V2102)

		if True == tmp92645 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp92637 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp92638 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp92639 := Call(__e, tmp92638)

			tmp92640 := Call(__e, tmp92637, MakeString("\n"), tmp92639)

			_ = tmp92640

			tmp92641 := Call(__e, PrimNS1Value(symns2_1value), symnl)

			tmp92642 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp92643 := Call(__e, tmp92642, V2102, MakeNumber(1))

			__e.TailApply(tmp92641, tmp92643)
			return

		}

	}, 1)

	tmp92646 := Call(__e, PrimNS1Value(symns2_1set), symnl, tmp92635)

	_ = tmp92646

	tmp92647 := MakeNative(func(__e *ControlFlow) {
		V2107 := __e.Get(1)
		_ = V2107
		V2108 := __e.Get(2)
		_ = V2108
		tmp92668 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92669 := Call(__e, tmp92668, Nil, V2107)

		if True == tmp92669 {
			__e.Return(Nil)
			return
		} else {
			tmp92666 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92667 := Call(__e, tmp92666, V2107)

			if True == tmp92667 {
				tmp92662 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp92663 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92664 := Call(__e, tmp92663, V2107)

				tmp92665 := Call(__e, tmp92662, tmp92664, V2108)

				if True == tmp92665 {
					tmp92659 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

					tmp92660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92661 := Call(__e, tmp92660, V2107)

					__e.TailApply(tmp92659, tmp92661, V2108)
					return

				} else {
					tmp92652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp92653 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92654 := Call(__e, tmp92653, V2107)

					tmp92655 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

					tmp92656 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92657 := Call(__e, tmp92656, V2107)

					tmp92658 := Call(__e, tmp92655, tmp92657, V2108)

					__e.TailApply(tmp92652, tmp92654, tmp92658)
					return

				}

			} else {
				tmp92650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92650, symdifference)
				return

			}

		}

	}, 2)

	tmp92670 := Call(__e, PrimNS1Value(symns2_1set), symdifference, tmp92647)

	_ = tmp92670

	tmp92671 := MakeNative(func(__e *ControlFlow) {
		V2111 := __e.Get(1)
		_ = V2111
		V2112 := __e.Get(2)
		_ = V2112
		__e.Return(V2112)
		return
	}, 2)

	tmp92672 := Call(__e, PrimNS1Value(symns2_1set), symdo, tmp92671)

	_ = tmp92672

	tmp92673 := MakeNative(func(__e *ControlFlow) {
		V2124 := __e.Get(1)
		_ = V2124
		V2125 := __e.Get(2)
		_ = V2125
		tmp92691 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92692 := Call(__e, tmp92691, Nil, V2125)

		if True == tmp92692 {
			__e.Return(False)
			return
		} else {
			tmp92689 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92690 := Call(__e, tmp92689, V2125)

			var ifres92683 Obj

			if True == tmp92690 {
				tmp92685 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp92686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92687 := Call(__e, tmp92686, V2125)

				tmp92688 := Call(__e, tmp92685, tmp92687, V2124)

				var ifres92684 Obj

				if True == tmp92688 {
					ifres92684 = True

				} else {
					ifres92684 = False

				}

				ifres92683 = ifres92684

			} else {
				ifres92683 = False

			}

			if True == ifres92683 {
				__e.Return(True)
				return
			} else {
				tmp92681 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92682 := Call(__e, tmp92681, V2125)

				if True == tmp92682 {
					tmp92678 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					tmp92679 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92680 := Call(__e, tmp92679, V2125)

					__e.TailApply(tmp92678, V2124, tmp92680)
					return

				} else {
					tmp92677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp92677, symelement_2)
					return

				}

			}

		}

	}, 2)

	tmp92693 := Call(__e, PrimNS1Value(symns2_1set), symelement_2, tmp92673)

	_ = tmp92693

	tmp92694 := MakeNative(func(__e *ControlFlow) {
		V2131 := __e.Get(1)
		_ = V2131
		tmp92696 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92697 := Call(__e, tmp92696, Nil, V2131)

		if True == tmp92697 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp92698 := Call(__e, PrimNS1Value(symns2_1set), symempty_2, tmp92694)

	_ = tmp92698

	tmp92699 := MakeNative(func(__e *ControlFlow) {
		V2134 := __e.Get(1)
		_ = V2134
		V2135 := __e.Get(2)
		_ = V2135
		tmp92700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fix_1help)

		tmp92701 := Call(__e, V2134, V2135)

		__e.TailApply(tmp92700, V2134, V2135, tmp92701)
		return

	}, 2)

	tmp92702 := Call(__e, PrimNS1Value(symns2_1set), symfix, tmp92699)

	_ = tmp92702

	tmp92703 := MakeNative(func(__e *ControlFlow) {
		V2146 := __e.Get(1)
		_ = V2146
		V2147 := __e.Get(2)
		_ = V2147
		V2148 := __e.Get(3)
		_ = V2148
		tmp92707 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92708 := Call(__e, tmp92707, V2148, V2147)

		if True == tmp92708 {
			__e.Return(V2148)
			return
		} else {
			tmp92705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fix_1help)

			tmp92706 := Call(__e, V2146, V2148)

			__e.TailApply(tmp92705, V2146, V2148, tmp92706)
			return

		}

	}, 3)

	tmp92709 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fix_1help, tmp92703)

	_ = tmp92709

	tmp92710 := MakeNative(func(__e *ControlFlow) {
		V2153 := __e.Get(1)
		_ = V2153
		V2154 := __e.Get(2)
		_ = V2154
		V2155 := __e.Get(3)
		_ = V2155
		V2156 := __e.Get(4)
		_ = V2156
		tmp92711 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			tmp92712 := MakeNative(func(__e *ControlFlow) {
				Added := __e.Get(1)
				_ = Added
				tmp92713 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2155)
					return
				}, 1)

				tmp92714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1_6)

				tmp92715 := Call(__e, tmp92714, V2156, V2153, Added)

				__e.TailApply(tmp92713, tmp92715)
				return

			}, 1)

			tmp92716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

			tmp92717 := Call(__e, tmp92716, V2154, V2155, Curr)

			__e.TailApply(tmp92712, tmp92717)
			return

		}, 1)

		tmp92718 := MakeNative(func(__e *ControlFlow) {
			tmp92719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(tmp92719, V2156, V2153)
			return

		}, 0)

		tmp92720 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp92721 := Call(__e, PrimNS1Value(symtry_1catch), tmp92718, tmp92720)

		__e.TailApply(tmp92711, tmp92721)
		return

	}, 4)

	tmp92722 := Call(__e, PrimNS1Value(symns2_1set), symput, tmp92710)

	_ = tmp92722

	tmp92723 := MakeNative(func(__e *ControlFlow) {
		V2160 := __e.Get(1)
		_ = V2160
		V2161 := __e.Get(2)
		_ = V2161
		V2162 := __e.Get(3)
		_ = V2162
		tmp92724 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			tmp92725 := MakeNative(func(__e *ControlFlow) {
				Removed := __e.Get(1)
				_ = Removed
				tmp92726 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2160)
					return
				}, 1)

				tmp92727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1_6)

				tmp92728 := Call(__e, tmp92727, V2162, V2160, Removed)

				__e.TailApply(tmp92726, tmp92728)
				return

			}, 1)

			tmp92729 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

			tmp92730 := Call(__e, tmp92729, V2161, Curr)

			__e.TailApply(tmp92725, tmp92730)
			return

		}, 1)

		tmp92731 := MakeNative(func(__e *ControlFlow) {
			tmp92732 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(tmp92732, V2162, V2160)
			return

		}, 0)

		tmp92733 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp92734 := Call(__e, PrimNS1Value(symtry_1catch), tmp92731, tmp92733)

		__e.TailApply(tmp92724, tmp92734)
		return

	}, 3)

	tmp92735 := Call(__e, PrimNS1Value(symns2_1set), symunput, tmp92723)

	_ = tmp92735

	tmp92736 := MakeNative(func(__e *ControlFlow) {
		V2166 := __e.Get(1)
		_ = V2166
		V2167 := __e.Get(2)
		_ = V2167
		V2168 := __e.Get(3)
		_ = V2168
		tmp92737 := MakeNative(func(__e *ControlFlow) {
			Entry := __e.Get(1)
			_ = Entry
			tmp92738 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				tmp92742 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				tmp92743 := Call(__e, tmp92742, Result)

				if True == tmp92743 {
					tmp92741 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(tmp92741, MakeString("value not found\n"))
					return

				} else {
					tmp92740 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					__e.TailApply(tmp92740, Result)
					return

				}

			}, 1)

			tmp92744 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

			tmp92745 := Call(__e, tmp92744, V2167, Entry)

			__e.TailApply(tmp92738, tmp92745)
			return

		}, 1)

		tmp92746 := MakeNative(func(__e *ControlFlow) {
			tmp92747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(tmp92747, V2168, V2166)
			return

		}, 0)

		tmp92748 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp92749 := Call(__e, PrimNS1Value(symtry_1catch), tmp92746, tmp92748)

		__e.TailApply(tmp92737, tmp92749)
		return

	}, 3)

	tmp92750 := Call(__e, PrimNS1Value(symns2_1set), symget, tmp92736)

	_ = tmp92750

	tmp92751 := MakeNative(func(__e *ControlFlow) {
		V2171 := __e.Get(1)
		_ = V2171
		V2172 := __e.Get(2)
		_ = V2172
		tmp92752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mod)

		tmp92753 := Call(__e, PrimNS1Value(symns2_1value), symsum)

		tmp92754 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp92755 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp92756 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

			__e.TailApply(tmp92756, X)
			return

		}, 1)

		tmp92757 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

		tmp92758 := Call(__e, tmp92757, V2171)

		tmp92759 := Call(__e, tmp92754, tmp92755, tmp92758)

		tmp92760 := Call(__e, tmp92753, tmp92759)

		__e.TailApply(tmp92752, tmp92760, V2172)
		return

	}, 2)

	tmp92761 := Call(__e, PrimNS1Value(symns2_1set), symhash, tmp92751)

	_ = tmp92761

	tmp92762 := MakeNative(func(__e *ControlFlow) {
		V2175 := __e.Get(1)
		_ = V2175
		V2176 := __e.Get(2)
		_ = V2176
		tmp92763 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

		tmp92764 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiples)

		tmp92765 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp92766 := Call(__e, tmp92765, V2176, Nil)

		tmp92767 := Call(__e, tmp92764, V2175, tmp92766)

		__e.TailApply(tmp92763, V2175, tmp92767)
		return

	}, 2)

	tmp92768 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mod, tmp92762)

	_ = tmp92768

	tmp92769 := MakeNative(func(__e *ControlFlow) {
		V2179 := __e.Get(1)
		_ = V2179
		V2180 := __e.Get(2)
		_ = V2180
		tmp92789 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp92790 := Call(__e, tmp92789, V2180)

		var ifres92783 Obj

		if True == tmp92790 {
			tmp92785 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp92786 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp92787 := Call(__e, tmp92786, V2180)

			tmp92788 := Call(__e, tmp92785, tmp92787, V2179)

			var ifres92784 Obj

			if True == tmp92788 {
				ifres92784 = True

			} else {
				ifres92784 = False

			}

			ifres92783 = ifres92784

		} else {
			ifres92783 = False

		}

		if True == ifres92783 {
			tmp92782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(tmp92782, V2180)
			return

		} else {
			tmp92780 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92781 := Call(__e, tmp92780, V2180)

			if True == tmp92781 {
				tmp92773 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiples)

				tmp92774 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92775 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				tmp92776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92777 := Call(__e, tmp92776, V2180)

				tmp92778 := Call(__e, tmp92775, MakeNumber(2), tmp92777)

				tmp92779 := Call(__e, tmp92774, tmp92778, V2180)

				__e.TailApply(tmp92773, V2179, tmp92779)
				return

			} else {
				tmp92772 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92772, symshen_4multiples)
				return

			}

		}

	}, 2)

	tmp92791 := Call(__e, PrimNS1Value(symns2_1set), symshen_4multiples, tmp92769)

	_ = tmp92791

	tmp92792 := MakeNative(func(__e *ControlFlow) {
		V2185 := __e.Get(1)
		_ = V2185
		V2186 := __e.Get(2)
		_ = V2186
		tmp92823 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92824 := Call(__e, tmp92823, MakeNumber(0), V2185)

		if True == tmp92824 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp92821 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp92822 := Call(__e, tmp92821, Nil, V2186)

			if True == tmp92822 {
				__e.Return(V2185)
				return
			} else {
				tmp92819 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp92820 := Call(__e, tmp92819, V2186)

				var ifres92813 Obj

				if True == tmp92820 {
					tmp92815 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp92816 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92817 := Call(__e, tmp92816, V2186)

					tmp92818 := Call(__e, tmp92815, tmp92817, V2185)

					var ifres92814 Obj

					if True == tmp92818 {
						ifres92814 = True

					} else {
						ifres92814 = False

					}

					ifres92813 = ifres92814

				} else {
					ifres92813 = False

				}

				if True == ifres92813 {
					tmp92809 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					tmp92810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92811 := Call(__e, tmp92810, V2186)

					tmp92812 := Call(__e, tmp92809, tmp92811)

					if True == tmp92812 {
						__e.Return(V2185)
						return
					} else {
						tmp92806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

						tmp92807 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp92808 := Call(__e, tmp92807, V2186)

						__e.TailApply(tmp92806, V2185, tmp92808)
						return

					}

				} else {
					tmp92803 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp92804 := Call(__e, tmp92803, V2186)

					if True == tmp92804 {
						tmp92798 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

						tmp92799 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						tmp92800 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp92801 := Call(__e, tmp92800, V2186)

						tmp92802 := Call(__e, tmp92799, V2185, tmp92801)

						__e.TailApply(tmp92798, tmp92802, V2186)
						return

					} else {
						tmp92797 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp92797, symshen_4modh)
						return

					}

				}

			}

		}

	}, 2)

	tmp92825 := Call(__e, PrimNS1Value(symns2_1set), symshen_4modh, tmp92792)

	_ = tmp92825

	tmp92826 := MakeNative(func(__e *ControlFlow) {
		V2188 := __e.Get(1)
		_ = V2188
		tmp92839 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92840 := Call(__e, tmp92839, Nil, V2188)

		if True == tmp92840 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp92837 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92838 := Call(__e, tmp92837, V2188)

			if True == tmp92838 {
				tmp92830 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp92831 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92832 := Call(__e, tmp92831, V2188)

				tmp92833 := Call(__e, PrimNS1Value(symns2_1value), symsum)

				tmp92834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp92835 := Call(__e, tmp92834, V2188)

				tmp92836 := Call(__e, tmp92833, tmp92835)

				__e.TailApply(tmp92830, tmp92832, tmp92836)
				return

			} else {
				tmp92829 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92829, symsum)
				return

			}

		}

	}, 1)

	tmp92841 := Call(__e, PrimNS1Value(symns2_1set), symsum, tmp92826)

	_ = tmp92841

	tmp92842 := MakeNative(func(__e *ControlFlow) {
		V2196 := __e.Get(1)
		_ = V2196
		tmp92846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp92847 := Call(__e, tmp92846, V2196)

		if True == tmp92847 {
			tmp92845 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp92845, V2196)
			return

		} else {
			tmp92844 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp92844, MakeString("head expects a non-empty list"))
			return

		}

	}, 1)

	tmp92848 := Call(__e, PrimNS1Value(symns2_1set), symhead, tmp92842)

	_ = tmp92848

	tmp92849 := MakeNative(func(__e *ControlFlow) {
		V2204 := __e.Get(1)
		_ = V2204
		tmp92853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp92854 := Call(__e, tmp92853, V2204)

		if True == tmp92854 {
			tmp92852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(tmp92852, V2204)
			return

		} else {
			tmp92851 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp92851, MakeString("tail expects a non-empty list"))
			return

		}

	}, 1)

	tmp92855 := Call(__e, PrimNS1Value(symns2_1set), symtail, tmp92849)

	_ = tmp92855

	tmp92856 := MakeNative(func(__e *ControlFlow) {
		V2206 := __e.Get(1)
		_ = V2206
		tmp92857 := Call(__e, PrimNS1Value(symns2_1value), sympos)

		__e.TailApply(tmp92857, V2206, MakeNumber(0))
		return

	}, 1)

	tmp92858 := Call(__e, PrimNS1Value(symns2_1set), symhdstr, tmp92856)

	_ = tmp92858

	tmp92859 := MakeNative(func(__e *ControlFlow) {
		V2211 := __e.Get(1)
		_ = V2211
		V2212 := __e.Get(2)
		_ = V2212
		tmp92880 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92881 := Call(__e, tmp92880, Nil, V2211)

		if True == tmp92881 {
			__e.Return(Nil)
			return
		} else {
			tmp92878 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92879 := Call(__e, tmp92878, V2211)

			if True == tmp92879 {
				tmp92874 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp92875 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92876 := Call(__e, tmp92875, V2211)

				tmp92877 := Call(__e, tmp92874, tmp92876, V2212)

				if True == tmp92877 {
					tmp92867 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp92868 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92869 := Call(__e, tmp92868, V2211)

					tmp92870 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

					tmp92871 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92872 := Call(__e, tmp92871, V2211)

					tmp92873 := Call(__e, tmp92870, tmp92872, V2212)

					__e.TailApply(tmp92867, tmp92869, tmp92873)
					return

				} else {
					tmp92864 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

					tmp92865 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92866 := Call(__e, tmp92865, V2211)

					__e.TailApply(tmp92864, tmp92866, V2212)
					return

				}

			} else {
				tmp92862 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92862, symintersection)
				return

			}

		}

	}, 2)

	tmp92882 := Call(__e, PrimNS1Value(symns2_1set), symintersection, tmp92859)

	_ = tmp92882

	tmp92883 := MakeNative(func(__e *ControlFlow) {
		V2214 := __e.Get(1)
		_ = V2214
		tmp92884 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reverse__help)

		__e.TailApply(tmp92884, V2214, Nil)
		return

	}, 1)

	tmp92885 := Call(__e, PrimNS1Value(symns2_1set), symreverse, tmp92883)

	_ = tmp92885

	tmp92886 := MakeNative(func(__e *ControlFlow) {
		V2217 := __e.Get(1)
		_ = V2217
		V2218 := __e.Get(2)
		_ = V2218
		tmp92899 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92900 := Call(__e, tmp92899, Nil, V2217)

		if True == tmp92900 {
			__e.Return(V2218)
			return
		} else {
			tmp92897 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92898 := Call(__e, tmp92897, V2217)

			if True == tmp92898 {
				tmp92890 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reverse__help)

				tmp92891 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp92892 := Call(__e, tmp92891, V2217)

				tmp92893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92894 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92895 := Call(__e, tmp92894, V2217)

				tmp92896 := Call(__e, tmp92893, tmp92895, V2218)

				__e.TailApply(tmp92890, tmp92892, tmp92896)
				return

			} else {
				tmp92889 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92889, symshen_4reverse__help)
				return

			}

		}

	}, 2)

	tmp92901 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reverse__help, tmp92886)

	_ = tmp92901

	tmp92902 := MakeNative(func(__e *ControlFlow) {
		V2221 := __e.Get(1)
		_ = V2221
		V2222 := __e.Get(2)
		_ = V2222
		tmp92923 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92924 := Call(__e, tmp92923, Nil, V2221)

		if True == tmp92924 {
			__e.Return(V2222)
			return
		} else {
			tmp92921 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92922 := Call(__e, tmp92921, V2221)

			if True == tmp92922 {
				tmp92917 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp92918 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp92919 := Call(__e, tmp92918, V2221)

				tmp92920 := Call(__e, tmp92917, tmp92919, V2222)

				if True == tmp92920 {
					tmp92914 := Call(__e, PrimNS1Value(symns2_1value), symunion)

					tmp92915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92916 := Call(__e, tmp92915, V2221)

					__e.TailApply(tmp92914, tmp92916, V2222)
					return

				} else {
					tmp92907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp92908 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp92909 := Call(__e, tmp92908, V2221)

					tmp92910 := Call(__e, PrimNS1Value(symns2_1value), symunion)

					tmp92911 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp92912 := Call(__e, tmp92911, V2221)

					tmp92913 := Call(__e, tmp92910, tmp92912, V2222)

					__e.TailApply(tmp92907, tmp92909, tmp92913)
					return

				}

			} else {
				tmp92905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92905, symunion)
				return

			}

		}

	}, 2)

	tmp92925 := Call(__e, PrimNS1Value(symns2_1set), symunion, tmp92902)

	_ = tmp92925

	tmp92926 := MakeNative(func(__e *ControlFlow) {
		V2224 := __e.Get(1)
		_ = V2224
		tmp92927 := MakeNative(func(__e *ControlFlow) {
			Message := __e.Get(1)
			_ = Message
			tmp92928 := MakeNative(func(__e *ControlFlow) {
				Y_1or_1N := __e.Get(1)
				_ = Y_1or_1N
				tmp92929 := MakeNative(func(__e *ControlFlow) {
					Input := __e.Get(1)
					_ = Input
					tmp92939 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp92940 := Call(__e, tmp92939, MakeString("y"), Input)

					if True == tmp92940 {
						__e.Return(True)
						return
					} else {
						tmp92937 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp92938 := Call(__e, tmp92937, MakeString("n"), Input)

						if True == tmp92938 {
							__e.Return(False)
							return
						} else {
							tmp92932 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

							tmp92933 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

							tmp92934 := Call(__e, tmp92933)

							tmp92935 := Call(__e, tmp92932, MakeString("please answer y or n\n"), tmp92934)

							_ = tmp92935

							tmp92936 := Call(__e, PrimNS1Value(symns2_1value), symy_1or_1n_2)

							__e.TailApply(tmp92936, V2224)
							return

						}

					}

				}, 1)

				tmp92941 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp92942 := Call(__e, PrimNS1Value(symns2_1value), symread)

				tmp92943 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

				tmp92944 := Call(__e, tmp92943)

				tmp92945 := Call(__e, tmp92942, tmp92944)

				tmp92946 := Call(__e, tmp92941, tmp92945, MakeString(""), symshen_4s)

				__e.TailApply(tmp92929, tmp92946)
				return

			}, 1)

			tmp92947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp92948 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp92949 := Call(__e, tmp92948)

			tmp92950 := Call(__e, tmp92947, MakeString(" (y/n) "), tmp92949)

			__e.TailApply(tmp92928, tmp92950)
			return

		}, 1)

		tmp92951 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp92952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

		tmp92953 := Call(__e, tmp92952, V2224)

		tmp92954 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp92955 := Call(__e, tmp92954)

		tmp92956 := Call(__e, tmp92951, tmp92953, tmp92955)

		__e.TailApply(tmp92927, tmp92956)
		return

	}, 1)

	tmp92957 := Call(__e, PrimNS1Value(symns2_1set), symy_1or_1n_2, tmp92926)

	_ = tmp92957

	tmp92958 := MakeNative(func(__e *ControlFlow) {
		V2226 := __e.Get(1)
		_ = V2226
		if True == V2226 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}
	}, 1)

	tmp92960 := Call(__e, PrimNS1Value(symns2_1set), symnot, tmp92958)

	_ = tmp92960

	tmp92961 := MakeNative(func(__e *ControlFlow) {
		V2239 := __e.Get(1)
		_ = V2239
		V2240 := __e.Get(2)
		_ = V2240
		V2241 := __e.Get(3)
		_ = V2241
		tmp92969 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92970 := Call(__e, tmp92969, V2241, V2240)

		if True == tmp92970 {
			__e.Return(V2239)
			return
		} else {
			tmp92967 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp92968 := Call(__e, tmp92967, V2241)

			if True == tmp92968 {
				tmp92964 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp92965 := MakeNative(func(__e *ControlFlow) {
					W := __e.Get(1)
					_ = W
					tmp92966 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					__e.TailApply(tmp92966, V2239, V2240, W)
					return

				}, 1)

				__e.TailApply(tmp92964, tmp92965, V2241)
				return

			} else {
				__e.Return(V2241)
				return
			}

		}

	}, 3)

	tmp92971 := Call(__e, PrimNS1Value(symns2_1set), symsubst, tmp92961)

	_ = tmp92971

	tmp92972 := MakeNative(func(__e *ControlFlow) {
		V2243 := __e.Get(1)
		_ = V2243
		tmp92973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4explode_1h)

		tmp92974 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp92975 := Call(__e, tmp92974, V2243, MakeString(""), symshen_4a)

		__e.TailApply(tmp92973, tmp92975)
		return

	}, 1)

	tmp92976 := Call(__e, PrimNS1Value(symns2_1set), symexplode, tmp92972)

	_ = tmp92976

	tmp92977 := MakeNative(func(__e *ControlFlow) {
		V2245 := __e.Get(1)
		_ = V2245
		tmp92990 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92991 := Call(__e, tmp92990, MakeString(""), V2245)

		if True == tmp92991 {
			__e.Return(Nil)
			return
		} else {
			tmp92988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			tmp92989 := Call(__e, tmp92988, V2245)

			if True == tmp92989 {
				tmp92981 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp92982 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				tmp92983 := Call(__e, tmp92982, V2245, MakeNumber(0))

				tmp92984 := Call(__e, PrimNS1Value(symns2_1value), symshen_4explode_1h)

				tmp92985 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp92986 := Call(__e, tmp92985, V2245)

				tmp92987 := Call(__e, tmp92984, tmp92986)

				__e.TailApply(tmp92981, tmp92983, tmp92987)
				return

			} else {
				tmp92980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp92980, symshen_4explode_1h)
				return

			}

		}

	}, 1)

	tmp92992 := Call(__e, PrimNS1Value(symns2_1set), symshen_4explode_1h, tmp92977)

	_ = tmp92992

	tmp92993 := MakeNative(func(__e *ControlFlow) {
		V2247 := __e.Get(1)
		_ = V2247
		tmp92994 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp92998 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp92999 := Call(__e, tmp92998, V2247, MakeString(""))

		var ifres92995 Obj

		if True == tmp92999 {
			ifres92995 = MakeString("")

		} else {
			tmp92996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp92997 := Call(__e, tmp92996, V2247, MakeString("/"), symshen_4a)

			ifres92995 = tmp92997

		}

		__e.TailApply(tmp92994, sym_dhome_1directory_d, ifres92995)
		return

	}, 1)

	tmp93000 := Call(__e, PrimNS1Value(symns2_1set), symcd, tmp92993)

	_ = tmp93000

	tmp93001 := MakeNative(func(__e *ControlFlow) {
		V2250 := __e.Get(1)
		_ = V2250
		V2251 := __e.Get(2)
		_ = V2251
		tmp93014 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93015 := Call(__e, tmp93014, Nil, V2251)

		if True == tmp93015 {
			__e.Return(True)
			return
		} else {
			tmp93012 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93013 := Call(__e, tmp93012, V2251)

			if True == tmp93013 {
				tmp93005 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					tmp93006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

					tmp93007 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93008 := Call(__e, tmp93007, V2251)

					__e.TailApply(tmp93006, V2250, tmp93008)
					return

				}, 1)

				tmp93009 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93010 := Call(__e, tmp93009, V2251)

				tmp93011 := Call(__e, V2250, tmp93010)

				__e.TailApply(tmp93005, tmp93011)
				return

			} else {
				tmp93004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp93004, symshen_4for_1each)
				return

			}

		}

	}, 2)

	tmp93016 := Call(__e, PrimNS1Value(symns2_1set), symshen_4for_1each, tmp93001)

	_ = tmp93016

	tmp93017 := MakeNative(func(__e *ControlFlow) {
		V2256 := __e.Get(1)
		_ = V2256
		V2257 := __e.Get(2)
		_ = V2257
		tmp93030 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93031 := Call(__e, tmp93030, Nil, V2257)

		if True == tmp93031 {
			__e.Return(Nil)
			return
		} else {
			tmp93028 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93029 := Call(__e, tmp93028, V2257)

			if True == tmp93029 {
				tmp93020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp93021 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93022 := Call(__e, tmp93021, V2257)

				tmp93023 := Call(__e, V2256, tmp93022)

				tmp93024 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp93025 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93026 := Call(__e, tmp93025, V2257)

				tmp93027 := Call(__e, tmp93024, V2256, tmp93026)

				__e.TailApply(tmp93020, tmp93023, tmp93027)
				return

			} else {
				__e.TailApply(V2256, V2257)
				return
			}

		}

	}, 2)

	tmp93032 := Call(__e, PrimNS1Value(symns2_1set), symmap, tmp93017)

	_ = tmp93032

	tmp93033 := MakeNative(func(__e *ControlFlow) {
		V2259 := __e.Get(1)
		_ = V2259
		tmp93034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4length_1h)

		__e.TailApply(tmp93034, V2259, MakeNumber(0))
		return

	}, 1)

	tmp93035 := Call(__e, PrimNS1Value(symns2_1set), symlength, tmp93033)

	_ = tmp93035

	tmp93036 := MakeNative(func(__e *ControlFlow) {
		V2262 := __e.Get(1)
		_ = V2262
		V2263 := __e.Get(2)
		_ = V2263
		tmp93043 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93044 := Call(__e, tmp93043, Nil, V2262)

		if True == tmp93044 {
			__e.Return(V2263)
			return
		} else {
			tmp93038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4length_1h)

			tmp93039 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp93040 := Call(__e, tmp93039, V2262)

			tmp93041 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			tmp93042 := Call(__e, tmp93041, V2263, MakeNumber(1))

			__e.TailApply(tmp93038, tmp93040, tmp93042)
			return

		}

	}, 2)

	tmp93045 := Call(__e, PrimNS1Value(symns2_1set), symshen_4length_1h, tmp93036)

	_ = tmp93045

	tmp93046 := MakeNative(func(__e *ControlFlow) {
		V2275 := __e.Get(1)
		_ = V2275
		V2276 := __e.Get(2)
		_ = V2276
		tmp93060 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93061 := Call(__e, tmp93060, V2276, V2275)

		if True == tmp93061 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp93058 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93059 := Call(__e, tmp93058, V2276)

			if True == tmp93059 {
				tmp93049 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp93050 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

				tmp93051 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93052 := Call(__e, tmp93051, V2276)

				tmp93053 := Call(__e, tmp93050, V2275, tmp93052)

				tmp93054 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

				tmp93055 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93056 := Call(__e, tmp93055, V2276)

				tmp93057 := Call(__e, tmp93054, V2275, tmp93056)

				__e.TailApply(tmp93049, tmp93053, tmp93057)
				return

			} else {
				__e.Return(MakeNumber(0))
				return
			}

		}

	}, 2)

	tmp93062 := Call(__e, PrimNS1Value(symns2_1set), symoccurrences, tmp93046)

	_ = tmp93062

	tmp93063 := MakeNative(func(__e *ControlFlow) {
		V2283 := __e.Get(1)
		_ = V2283
		V2284 := __e.Get(2)
		_ = V2284
		tmp93087 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93088 := Call(__e, tmp93087, MakeNumber(1), V2283)

		var ifres93083 Obj

		if True == tmp93088 {
			tmp93085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93086 := Call(__e, tmp93085, V2284)

			var ifres93084 Obj

			if True == tmp93086 {
				ifres93084 = True

			} else {
				ifres93084 = False

			}

			ifres93083 = ifres93084

		} else {
			ifres93083 = False

		}

		if True == ifres93083 {
			tmp93082 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp93082, V2284)
			return

		} else {
			tmp93080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93081 := Call(__e, tmp93080, V2284)

			if True == tmp93081 {
				tmp93075 := Call(__e, PrimNS1Value(symns2_1value), symnth)

				tmp93076 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp93077 := Call(__e, tmp93076, V2283, MakeNumber(1))

				tmp93078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93079 := Call(__e, tmp93078, V2284)

				__e.TailApply(tmp93075, tmp93077, tmp93079)
				return

			} else {
				tmp93066 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp93067 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp93068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp93069 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp93070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp93071 := Call(__e, tmp93070, V2284, MakeString("\n"), symshen_4a)

				tmp93072 := Call(__e, tmp93069, MakeString(", "), tmp93071)

				tmp93073 := Call(__e, tmp93068, V2283, tmp93072, symshen_4a)

				tmp93074 := Call(__e, tmp93067, MakeString("nth applied to "), tmp93073)

				__e.TailApply(tmp93066, tmp93074)
				return

			}

		}

	}, 2)

	tmp93089 := Call(__e, PrimNS1Value(symns2_1set), symnth, tmp93063)

	_ = tmp93089

	tmp93090 := MakeNative(func(__e *ControlFlow) {
		V2286 := __e.Get(1)
		_ = V2286
		tmp93100 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

		tmp93101 := Call(__e, tmp93100, V2286)

		if True == tmp93101 {
			tmp93093 := MakeNative(func(__e *ControlFlow) {
				Abs := __e.Get(1)
				_ = Abs
				tmp93094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4integer_1test_2)

				tmp93095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4magless)

				tmp93096 := Call(__e, tmp93095, Abs, MakeNumber(1))

				__e.TailApply(tmp93094, Abs, tmp93096)
				return

			}, 1)

			tmp93097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abs)

			tmp93098 := Call(__e, tmp93097, V2286)

			tmp93099 := Call(__e, tmp93093, tmp93098)

			if True == tmp93099 {
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

	tmp93102 := Call(__e, PrimNS1Value(symns2_1set), syminteger_2, tmp93090)

	_ = tmp93102

	tmp93103 := MakeNative(func(__e *ControlFlow) {
		V2288 := __e.Get(1)
		_ = V2288
		tmp93106 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		tmp93107 := Call(__e, tmp93106, V2288, MakeNumber(0))

		if True == tmp93107 {
			__e.Return(V2288)
			return
		} else {
			tmp93105 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			__e.TailApply(tmp93105, MakeNumber(0), V2288)
			return

		}

	}, 1)

	tmp93108 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abs, tmp93103)

	_ = tmp93108

	tmp93109 := MakeNative(func(__e *ControlFlow) {
		V2291 := __e.Get(1)
		_ = V2291
		V2292 := __e.Get(2)
		_ = V2292
		tmp93110 := MakeNative(func(__e *ControlFlow) {
			Nx2 := __e.Get(1)
			_ = Nx2
			tmp93113 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp93114 := Call(__e, tmp93113, Nx2, V2291)

			if True == tmp93114 {
				__e.Return(V2292)
				return
			} else {
				tmp93112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4magless)

				__e.TailApply(tmp93112, V2291, Nx2)
				return

			}

		}, 1)

		tmp93115 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

		tmp93116 := Call(__e, tmp93115, V2292, MakeNumber(2))

		__e.TailApply(tmp93110, tmp93116)
		return

	}, 2)

	tmp93117 := Call(__e, PrimNS1Value(symns2_1set), symshen_4magless, tmp93109)

	_ = tmp93117

	tmp93118 := MakeNative(func(__e *ControlFlow) {
		V2298 := __e.Get(1)
		_ = V2298
		V2299 := __e.Get(2)
		_ = V2299
		tmp93131 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93132 := Call(__e, tmp93131, MakeNumber(0), V2298)

		if True == tmp93132 {
			__e.Return(True)
			return
		} else {
			tmp93129 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			tmp93130 := Call(__e, tmp93129, MakeNumber(1), V2298)

			if True == tmp93130 {
				__e.Return(False)
				return
			} else {
				tmp93121 := MakeNative(func(__e *ControlFlow) {
					Abs_1N := __e.Get(1)
					_ = Abs_1N
					tmp93125 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp93126 := Call(__e, tmp93125, MakeNumber(0), Abs_1N)

					if True == tmp93126 {
						tmp93124 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

						__e.TailApply(tmp93124, V2298)
						return

					} else {
						tmp93123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4integer_1test_2)

						__e.TailApply(tmp93123, Abs_1N, V2299)
						return

					}

				}, 1)

				tmp93127 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp93128 := Call(__e, tmp93127, V2298, V2299)

				__e.TailApply(tmp93121, tmp93128)
				return

			}

		}

	}, 2)

	tmp93133 := Call(__e, PrimNS1Value(symns2_1set), symshen_4integer_1test_2, tmp93118)

	_ = tmp93133

	tmp93134 := MakeNative(func(__e *ControlFlow) {
		V2304 := __e.Get(1)
		_ = V2304
		V2305 := __e.Get(2)
		_ = V2305
		tmp93148 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93149 := Call(__e, tmp93148, Nil, V2305)

		if True == tmp93149 {
			__e.Return(Nil)
			return
		} else {
			tmp93146 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93147 := Call(__e, tmp93146, V2305)

			if True == tmp93147 {
				tmp93138 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp93139 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93140 := Call(__e, tmp93139, V2305)

				tmp93141 := Call(__e, V2304, tmp93140)

				tmp93142 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

				tmp93143 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93144 := Call(__e, tmp93143, V2305)

				tmp93145 := Call(__e, tmp93142, V2304, tmp93144)

				__e.TailApply(tmp93138, tmp93141, tmp93145)
				return

			} else {
				tmp93137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp93137, symmapcan)
				return

			}

		}

	}, 2)

	tmp93150 := Call(__e, PrimNS1Value(symns2_1set), symmapcan, tmp93134)

	_ = tmp93150

	tmp93151 := MakeNative(func(__e *ControlFlow) {
		V2317 := __e.Get(1)
		_ = V2317
		V2318 := __e.Get(2)
		_ = V2318
		tmp93153 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93154 := Call(__e, tmp93153, V2318, V2317)

		if True == tmp93154 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp93155 := Call(__e, PrimNS1Value(symns2_1set), sym_a_a, tmp93151)

	_ = tmp93155

	tmp93156 := MakeNative(func(__e *ControlFlow) {
		tmp93157 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(tmp93157, MakeString(""))
		return

	}, 0)

	tmp93158 := Call(__e, PrimNS1Value(symns2_1set), symabort, tmp93156)

	_ = tmp93158

	tmp93159 := MakeNative(func(__e *ControlFlow) {
		V2320 := __e.Get(1)
		_ = V2320
		tmp93171 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp93172 := Call(__e, tmp93171, V2320)

		if True == tmp93172 {
			tmp93162 := MakeNative(func(__e *ControlFlow) {
				Val := __e.Get(1)
				_ = Val
				tmp93164 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp93165 := Call(__e, tmp93164, Val, symshen_4this_1symbol_1is_1unbound)

				if True == tmp93165 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}, 1)

			tmp93166 := MakeNative(func(__e *ControlFlow) {
				tmp93167 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				__e.TailApply(tmp93167, V2320)
				return

			}, 0)

			tmp93168 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4this_1symbol_1is_1unbound)
				return
			}, 1)

			tmp93169 := Call(__e, PrimNS1Value(symtry_1catch), tmp93166, tmp93168)

			tmp93170 := Call(__e, tmp93162, tmp93169)

			if True == tmp93170 {
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

	tmp93173 := Call(__e, PrimNS1Value(symns2_1set), symbound_2, tmp93159)

	_ = tmp93173

	tmp93174 := MakeNative(func(__e *ControlFlow) {
		V2322 := __e.Get(1)
		_ = V2322
		tmp93185 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93186 := Call(__e, tmp93185, MakeString(""), V2322)

		if True == tmp93186 {
			__e.Return(Nil)
			return
		} else {
			tmp93176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp93177 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

			tmp93178 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			tmp93179 := Call(__e, tmp93178, V2322, MakeNumber(0))

			tmp93180 := Call(__e, tmp93177, tmp93179)

			tmp93181 := Call(__e, PrimNS1Value(symns2_1value), symshen_4string_1_6bytes)

			tmp93182 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

			tmp93183 := Call(__e, tmp93182, V2322)

			tmp93184 := Call(__e, tmp93181, tmp93183)

			__e.TailApply(tmp93176, tmp93180, tmp93184)
			return

		}

	}, 1)

	tmp93187 := Call(__e, PrimNS1Value(symns2_1set), symshen_4string_1_6bytes, tmp93174)

	_ = tmp93187

	tmp93188 := MakeNative(func(__e *ControlFlow) {
		V2324 := __e.Get(1)
		_ = V2324
		tmp93189 := Call(__e, PrimNS1Value(symns2_1value), symset)

		__e.TailApply(tmp93189, symshen_4_dmaxinferences_d, V2324)
		return

	}, 1)

	tmp93190 := Call(__e, PrimNS1Value(symns2_1set), symmaxinferences, tmp93188)

	_ = tmp93190

	tmp93191 := MakeNative(func(__e *ControlFlow) {
		tmp93192 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93192, symshen_4_dinfs_d)
		return

	}, 0)

	tmp93193 := Call(__e, PrimNS1Value(symns2_1set), syminferences, tmp93191)

	_ = tmp93193

	tmp93194 := MakeNative(func(__e *ControlFlow) {
		V2326 := __e.Get(1)
		_ = V2326
		__e.Return(V2326)
		return
	}, 1)

	tmp93195 := Call(__e, PrimNS1Value(symns2_1set), symprotect, tmp93194)

	_ = tmp93195

	tmp93196 := MakeNative(func(__e *ControlFlow) {
		tmp93197 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93197, sym_dstoutput_d)
		return

	}, 0)

	tmp93198 := Call(__e, PrimNS1Value(symns2_1set), symstoutput, tmp93196)

	_ = tmp93198

	tmp93199 := MakeNative(func(__e *ControlFlow) {
		tmp93200 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93200, sym_dsterror_d)
		return

	}, 0)

	tmp93201 := Call(__e, PrimNS1Value(symns2_1set), symsterror, tmp93199)

	_ = tmp93201

	tmp93202 := MakeNative(func(__e *ControlFlow) {
		V2328 := __e.Get(1)
		_ = V2328
		tmp93203 := MakeNative(func(__e *ControlFlow) {
			Symbol := __e.Get(1)
			_ = Symbol
			tmp93210 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

			tmp93211 := Call(__e, tmp93210, Symbol)

			if True == tmp93211 {
				__e.Return(Symbol)
				return
			} else {
				tmp93205 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp93206 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp93207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp93208 := Call(__e, tmp93207, V2328, MakeString(" to a symbol"), symshen_4s)

				tmp93209 := Call(__e, tmp93206, MakeString("cannot intern "), tmp93208)

				__e.TailApply(tmp93205, tmp93209)
				return

			}

		}, 1)

		tmp93212 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp93213 := Call(__e, tmp93212, V2328)

		__e.TailApply(tmp93203, tmp93213)
		return

	}, 1)

	tmp93214 := Call(__e, PrimNS1Value(symns2_1set), symstring_1_6symbol, tmp93202)

	_ = tmp93214

	tmp93215 := MakeNative(func(__e *ControlFlow) {
		V2334 := __e.Get(1)
		_ = V2334
		tmp93223 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93224 := Call(__e, tmp93223, sym_7, V2334)

		if True == tmp93224 {
			tmp93222 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp93222, symshen_4_doptimise_d, True)
			return

		} else {
			tmp93220 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp93221 := Call(__e, tmp93220, sym_1, V2334)

			if True == tmp93221 {
				tmp93219 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp93219, symshen_4_doptimise_d, False)
				return

			} else {
				tmp93218 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp93218, MakeString("optimise expects a + or a -.\n"))
				return

			}

		}

	}, 1)

	tmp93225 := Call(__e, PrimNS1Value(symns2_1set), symoptimise, tmp93215)

	_ = tmp93225

	tmp93226 := MakeNative(func(__e *ControlFlow) {
		tmp93227 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93227, sym_dos_d)
		return

	}, 0)

	tmp93228 := Call(__e, PrimNS1Value(symns2_1set), symos, tmp93226)

	_ = tmp93228

	tmp93229 := MakeNative(func(__e *ControlFlow) {
		tmp93230 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93230, sym_dlanguage_d)
		return

	}, 0)

	tmp93231 := Call(__e, PrimNS1Value(symns2_1set), symlanguage, tmp93229)

	_ = tmp93231

	tmp93232 := MakeNative(func(__e *ControlFlow) {
		tmp93233 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93233, sym_dversion_d)
		return

	}, 0)

	tmp93234 := Call(__e, PrimNS1Value(symns2_1set), symversion, tmp93232)

	_ = tmp93234

	tmp93235 := MakeNative(func(__e *ControlFlow) {
		tmp93236 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93236, sym_dport_d)
		return

	}, 0)

	tmp93237 := Call(__e, PrimNS1Value(symns2_1set), symport, tmp93235)

	_ = tmp93237

	tmp93238 := MakeNative(func(__e *ControlFlow) {
		tmp93239 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93239, sym_dporters_d)
		return

	}, 0)

	tmp93240 := Call(__e, PrimNS1Value(symns2_1set), symporters, tmp93238)

	_ = tmp93240

	tmp93241 := MakeNative(func(__e *ControlFlow) {
		tmp93242 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93242, sym_dimplementation_d)
		return

	}, 0)

	tmp93243 := Call(__e, PrimNS1Value(symns2_1set), symimplementation, tmp93241)

	_ = tmp93243

	tmp93244 := MakeNative(func(__e *ControlFlow) {
		tmp93245 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp93245, sym_drelease_d)
		return

	}, 0)

	tmp93246 := Call(__e, PrimNS1Value(symns2_1set), symrelease, tmp93244)

	_ = tmp93246

	tmp93247 := MakeNative(func(__e *ControlFlow) {
		V2336 := __e.Get(1)
		_ = V2336
		tmp93248 := MakeNative(func(__e *ControlFlow) {
			tmp93249 := Call(__e, PrimNS1Value(symns2_1value), symexternal)

			tmp93250 := Call(__e, tmp93249, V2336)

			_ = tmp93250

			__e.Return(True)
			return

		}, 0)

		tmp93251 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp93248, tmp93251)
		return

	}, 1)

	tmp93252 := Call(__e, PrimNS1Value(symns2_1set), sympackage_2, tmp93247)

	_ = tmp93252

	tmp93253 := MakeNative(func(__e *ControlFlow) {
		V2338 := __e.Get(1)
		_ = V2338
		tmp93254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lookup_1func)

		__e.TailApply(tmp93254, V2338)
		return

	}, 1)

	tmp93255 := Call(__e, PrimNS1Value(symns2_1set), symfunction, tmp93253)

	_ = tmp93255

	tmp93256 := MakeNative(func(__e *ControlFlow) {
		V2340 := __e.Get(1)
		_ = V2340
		tmp93257 := MakeNative(func(__e *ControlFlow) {
			tmp93258 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp93259 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp93260 := Call(__e, tmp93259, sym_dproperty_1vector_d)

			__e.TailApply(tmp93258, V2340, symshen_4lambda_1form, tmp93260)
			return

		}, 0)

		tmp93261 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp93262 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp93263 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp93264 := Call(__e, tmp93263, V2340, MakeString(" has no lambda expansion\n"), symshen_4a)

			__e.TailApply(tmp93262, tmp93264)
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp93257, tmp93261)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4lookup_1func, tmp93256)
	return

}, 0)
