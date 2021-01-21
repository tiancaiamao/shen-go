package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp77522 := MakeNative(func(__e *ControlFlow) {
		V748 := __e.Get(1)
		_ = V748
		tmp77523 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			tmp77528 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77529 := Call(__e, tmp77528, V748, Y)

			if True == tmp77529 {
				__e.Return(V748)
				return
			} else {
				tmp77525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

				tmp77526 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp77527 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

					__e.TailApply(tmp77527, Z)
					return

				}, 1)

				__e.TailApply(tmp77525, tmp77526, Y)
				return

			}

		}, 1)

		tmp77530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compose)

		tmp77531 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp77532 := Call(__e, tmp77531, sym_dmacros_d)

		tmp77533 := Call(__e, tmp77530, tmp77532, V748)

		__e.TailApply(tmp77523, tmp77533)
		return

	}, 1)

	tmp77534 := Call(__e, PrimNS1Value(symns2_1set), symmacroexpand, tmp77522)

	_ = tmp77534

	tmp77535 := MakeNative(func(__e *ControlFlow) {
		V750 := __e.Get(1)
		_ = V750
		tmp77562 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77563 := Call(__e, tmp77562, V750)

		var ifres77550 Obj

		if True == tmp77563 {
			tmp77558 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77559 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77560 := Call(__e, tmp77559, V750)

			tmp77561 := Call(__e, tmp77558, symerror, tmp77560)

			var ifres77552 Obj

			if True == tmp77561 {
				tmp77554 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77556 := Call(__e, tmp77555, V750)

				tmp77557 := Call(__e, tmp77554, tmp77556)

				var ifres77553 Obj

				if True == tmp77557 {
					ifres77553 = True

				} else {
					ifres77553 = False

				}

				ifres77552 = ifres77553

			} else {
				ifres77552 = False

			}

			var ifres77551 Obj

			if True == ifres77552 {
				ifres77551 = True

			} else {
				ifres77551 = False

			}

			ifres77550 = ifres77551

		} else {
			ifres77550 = False

		}

		if True == ifres77550 {
			tmp77537 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77538 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			tmp77540 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77541 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77542 := Call(__e, tmp77541, V750)

			tmp77543 := Call(__e, tmp77540, tmp77542)

			tmp77544 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77545 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77546 := Call(__e, tmp77545, V750)

			tmp77547 := Call(__e, tmp77544, tmp77546)

			tmp77548 := Call(__e, tmp77539, tmp77543, tmp77547)

			tmp77549 := Call(__e, tmp77538, tmp77548, Nil)

			__e.TailApply(tmp77537, symsimple_1error, tmp77549)
			return

		} else {
			__e.Return(V750)
			return
		}

	}, 1)

	tmp77564 := Call(__e, PrimNS1Value(symns2_1set), symshen_4error_1macro, tmp77535)

	_ = tmp77564

	tmp77565 := MakeNative(func(__e *ControlFlow) {
		V752 := __e.Get(1)
		_ = V752
		tmp77630 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77631 := Call(__e, tmp77630, V752)

		var ifres77618 Obj

		if True == tmp77631 {
			tmp77626 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77627 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77628 := Call(__e, tmp77627, V752)

			tmp77629 := Call(__e, tmp77626, symoutput, tmp77628)

			var ifres77620 Obj

			if True == tmp77629 {
				tmp77622 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77624 := Call(__e, tmp77623, V752)

				tmp77625 := Call(__e, tmp77622, tmp77624)

				var ifres77621 Obj

				if True == tmp77625 {
					ifres77621 = True

				} else {
					ifres77621 = False

				}

				ifres77620 = ifres77621

			} else {
				ifres77620 = False

			}

			var ifres77619 Obj

			if True == ifres77620 {
				ifres77619 = True

			} else {
				ifres77619 = False

			}

			ifres77618 = ifres77619

		} else {
			ifres77618 = False

		}

		if True == ifres77618 {
			tmp77601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			tmp77604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77605 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77606 := Call(__e, tmp77605, V752)

			tmp77607 := Call(__e, tmp77604, tmp77606)

			tmp77608 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77609 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77610 := Call(__e, tmp77609, V752)

			tmp77611 := Call(__e, tmp77608, tmp77610)

			tmp77612 := Call(__e, tmp77603, tmp77607, tmp77611)

			tmp77613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77614 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77615 := Call(__e, tmp77614, symstoutput, Nil)

			tmp77616 := Call(__e, tmp77613, tmp77615, Nil)

			tmp77617 := Call(__e, tmp77602, tmp77612, tmp77616)

			__e.TailApply(tmp77601, symshen_4prhush, tmp77617)
			return

		} else {
			tmp77599 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77600 := Call(__e, tmp77599, V752)

			var ifres77579 Obj

			if True == tmp77600 {
				tmp77595 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp77596 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77597 := Call(__e, tmp77596, V752)

				tmp77598 := Call(__e, tmp77595, sympr, tmp77597)

				var ifres77581 Obj

				if True == tmp77598 {
					tmp77591 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp77592 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77593 := Call(__e, tmp77592, V752)

					tmp77594 := Call(__e, tmp77591, tmp77593)

					var ifres77583 Obj

					if True == tmp77594 {
						tmp77585 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp77586 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77587 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77588 := Call(__e, tmp77587, V752)

						tmp77589 := Call(__e, tmp77586, tmp77588)

						tmp77590 := Call(__e, tmp77585, Nil, tmp77589)

						var ifres77584 Obj

						if True == tmp77590 {
							ifres77584 = True

						} else {
							ifres77584 = False

						}

						ifres77583 = ifres77584

					} else {
						ifres77583 = False

					}

					var ifres77582 Obj

					if True == ifres77583 {
						ifres77582 = True

					} else {
						ifres77582 = False

					}

					ifres77581 = ifres77582

				} else {
					ifres77581 = False

				}

				var ifres77580 Obj

				if True == ifres77581 {
					ifres77580 = True

				} else {
					ifres77580 = False

				}

				ifres77579 = ifres77580

			} else {
				ifres77579 = False

			}

			if True == ifres77579 {
				tmp77568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77569 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77570 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77571 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77572 := Call(__e, tmp77571, V752)

				tmp77573 := Call(__e, tmp77570, tmp77572)

				tmp77574 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77575 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77576 := Call(__e, tmp77575, symstoutput, Nil)

				tmp77577 := Call(__e, tmp77574, tmp77576, Nil)

				tmp77578 := Call(__e, tmp77569, tmp77573, tmp77577)

				__e.TailApply(tmp77568, sympr, tmp77578)
				return

			} else {
				__e.Return(V752)
				return
			}

		}

	}, 1)

	tmp77632 := Call(__e, PrimNS1Value(symns2_1set), symshen_4output_1macro, tmp77565)

	_ = tmp77632

	tmp77633 := MakeNative(func(__e *ControlFlow) {
		V754 := __e.Get(1)
		_ = V754
		tmp77656 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77657 := Call(__e, tmp77656, V754)

		var ifres77644 Obj

		if True == tmp77657 {
			tmp77652 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77653 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77654 := Call(__e, tmp77653, V754)

			tmp77655 := Call(__e, tmp77652, symmake_1string, tmp77654)

			var ifres77646 Obj

			if True == tmp77655 {
				tmp77648 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77649 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77650 := Call(__e, tmp77649, V754)

				tmp77651 := Call(__e, tmp77648, tmp77650)

				var ifres77647 Obj

				if True == tmp77651 {
					ifres77647 = True

				} else {
					ifres77647 = False

				}

				ifres77646 = ifres77647

			} else {
				ifres77646 = False

			}

			var ifres77645 Obj

			if True == ifres77646 {
				ifres77645 = True

			} else {
				ifres77645 = False

			}

			ifres77644 = ifres77645

		} else {
			ifres77644 = False

		}

		if True == ifres77644 {
			tmp77635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			tmp77636 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77637 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77638 := Call(__e, tmp77637, V754)

			tmp77639 := Call(__e, tmp77636, tmp77638)

			tmp77640 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77642 := Call(__e, tmp77641, V754)

			tmp77643 := Call(__e, tmp77640, tmp77642)

			__e.TailApply(tmp77635, tmp77639, tmp77643)
			return

		} else {
			__e.Return(V754)
			return
		}

	}, 1)

	tmp77658 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1string_1macro, tmp77633)

	_ = tmp77658

	tmp77659 := MakeNative(func(__e *ControlFlow) {
		V756 := __e.Get(1)
		_ = V756
		tmp77772 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77773 := Call(__e, tmp77772, V756)

		var ifres77760 Obj

		if True == tmp77773 {
			tmp77768 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77769 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77770 := Call(__e, tmp77769, V756)

			tmp77771 := Call(__e, tmp77768, symlineread, tmp77770)

			var ifres77762 Obj

			if True == tmp77771 {
				tmp77764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp77765 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77766 := Call(__e, tmp77765, V756)

				tmp77767 := Call(__e, tmp77764, Nil, tmp77766)

				var ifres77763 Obj

				if True == tmp77767 {
					ifres77763 = True

				} else {
					ifres77763 = False

				}

				ifres77762 = ifres77763

			} else {
				ifres77762 = False

			}

			var ifres77761 Obj

			if True == ifres77762 {
				ifres77761 = True

			} else {
				ifres77761 = False

			}

			ifres77760 = ifres77761

		} else {
			ifres77760 = False

		}

		if True == ifres77760 {
			tmp77755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77758 := Call(__e, tmp77757, symstinput, Nil)

			tmp77759 := Call(__e, tmp77756, tmp77758, Nil)

			__e.TailApply(tmp77755, symlineread, tmp77759)
			return

		} else {
			tmp77753 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77754 := Call(__e, tmp77753, V756)

			var ifres77741 Obj

			if True == tmp77754 {
				tmp77749 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp77750 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77751 := Call(__e, tmp77750, V756)

				tmp77752 := Call(__e, tmp77749, syminput, tmp77751)

				var ifres77743 Obj

				if True == tmp77752 {
					tmp77745 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp77746 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77747 := Call(__e, tmp77746, V756)

					tmp77748 := Call(__e, tmp77745, Nil, tmp77747)

					var ifres77744 Obj

					if True == tmp77748 {
						ifres77744 = True

					} else {
						ifres77744 = False

					}

					ifres77743 = ifres77744

				} else {
					ifres77743 = False

				}

				var ifres77742 Obj

				if True == ifres77743 {
					ifres77742 = True

				} else {
					ifres77742 = False

				}

				ifres77741 = ifres77742

			} else {
				ifres77741 = False

			}

			if True == ifres77741 {
				tmp77736 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77737 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77739 := Call(__e, tmp77738, symstinput, Nil)

				tmp77740 := Call(__e, tmp77737, tmp77739, Nil)

				__e.TailApply(tmp77736, syminput, tmp77740)
				return

			} else {
				tmp77734 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77735 := Call(__e, tmp77734, V756)

				var ifres77722 Obj

				if True == tmp77735 {
					tmp77730 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp77731 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp77732 := Call(__e, tmp77731, V756)

					tmp77733 := Call(__e, tmp77730, symread, tmp77732)

					var ifres77724 Obj

					if True == tmp77733 {
						tmp77726 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp77727 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77728 := Call(__e, tmp77727, V756)

						tmp77729 := Call(__e, tmp77726, Nil, tmp77728)

						var ifres77725 Obj

						if True == tmp77729 {
							ifres77725 = True

						} else {
							ifres77725 = False

						}

						ifres77724 = ifres77725

					} else {
						ifres77724 = False

					}

					var ifres77723 Obj

					if True == ifres77724 {
						ifres77723 = True

					} else {
						ifres77723 = False

					}

					ifres77722 = ifres77723

				} else {
					ifres77722 = False

				}

				if True == ifres77722 {
					tmp77717 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp77718 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp77719 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp77720 := Call(__e, tmp77719, symstinput, Nil)

					tmp77721 := Call(__e, tmp77718, tmp77720, Nil)

					__e.TailApply(tmp77717, symread, tmp77721)
					return

				} else {
					tmp77715 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp77716 := Call(__e, tmp77715, V756)

					var ifres77695 Obj

					if True == tmp77716 {
						tmp77711 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp77712 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp77713 := Call(__e, tmp77712, V756)

						tmp77714 := Call(__e, tmp77711, syminput_7, tmp77713)

						var ifres77697 Obj

						if True == tmp77714 {
							tmp77707 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp77708 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp77709 := Call(__e, tmp77708, V756)

							tmp77710 := Call(__e, tmp77707, tmp77709)

							var ifres77699 Obj

							if True == tmp77710 {
								tmp77701 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp77702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp77703 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp77704 := Call(__e, tmp77703, V756)

								tmp77705 := Call(__e, tmp77702, tmp77704)

								tmp77706 := Call(__e, tmp77701, Nil, tmp77705)

								var ifres77700 Obj

								if True == tmp77706 {
									ifres77700 = True

								} else {
									ifres77700 = False

								}

								ifres77699 = ifres77700

							} else {
								ifres77699 = False

							}

							var ifres77698 Obj

							if True == ifres77699 {
								ifres77698 = True

							} else {
								ifres77698 = False

							}

							ifres77697 = ifres77698

						} else {
							ifres77697 = False

						}

						var ifres77696 Obj

						if True == ifres77697 {
							ifres77696 = True

						} else {
							ifres77696 = False

						}

						ifres77695 = ifres77696

					} else {
						ifres77695 = False

					}

					if True == ifres77695 {
						tmp77684 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp77685 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp77686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp77687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77688 := Call(__e, tmp77687, V756)

						tmp77689 := Call(__e, tmp77686, tmp77688)

						tmp77690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp77691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp77692 := Call(__e, tmp77691, symstinput, Nil)

						tmp77693 := Call(__e, tmp77690, tmp77692, Nil)

						tmp77694 := Call(__e, tmp77685, tmp77689, tmp77693)

						__e.TailApply(tmp77684, syminput_7, tmp77694)
						return

					} else {
						tmp77682 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp77683 := Call(__e, tmp77682, V756)

						var ifres77670 Obj

						if True == tmp77683 {
							tmp77678 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp77679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp77680 := Call(__e, tmp77679, V756)

							tmp77681 := Call(__e, tmp77678, symread_1byte, tmp77680)

							var ifres77672 Obj

							if True == tmp77681 {
								tmp77674 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp77675 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp77676 := Call(__e, tmp77675, V756)

								tmp77677 := Call(__e, tmp77674, Nil, tmp77676)

								var ifres77673 Obj

								if True == tmp77677 {
									ifres77673 = True

								} else {
									ifres77673 = False

								}

								ifres77672 = ifres77673

							} else {
								ifres77672 = False

							}

							var ifres77671 Obj

							if True == ifres77672 {
								ifres77671 = True

							} else {
								ifres77671 = False

							}

							ifres77670 = ifres77671

						} else {
							ifres77670 = False

						}

						if True == ifres77670 {
							tmp77665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp77666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp77667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp77668 := Call(__e, tmp77667, symstinput, Nil)

							tmp77669 := Call(__e, tmp77666, tmp77668, Nil)

							__e.TailApply(tmp77665, symread_1byte, tmp77669)
							return

						} else {
							__e.Return(V756)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp77774 := Call(__e, PrimNS1Value(symns2_1set), symshen_4input_1macro, tmp77659)

	_ = tmp77774

	tmp77775 := MakeNative(func(__e *ControlFlow) {
		V759 := __e.Get(1)
		_ = V759
		V760 := __e.Get(2)
		_ = V760
		tmp77787 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp77788 := Call(__e, tmp77787, Nil, V759)

		if True == tmp77788 {
			__e.Return(V760)
			return
		} else {
			tmp77785 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77786 := Call(__e, tmp77785, V759)

			if True == tmp77786 {
				tmp77779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compose)

				tmp77780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77781 := Call(__e, tmp77780, V759)

				tmp77782 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77783 := Call(__e, tmp77782, V759)

				tmp77784 := Call(__e, tmp77783, V760)

				__e.TailApply(tmp77779, tmp77781, tmp77784)
				return

			} else {
				tmp77778 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp77778, symshen_4compose)
				return

			}

		}

	}, 2)

	tmp77789 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compose, tmp77775)

	_ = tmp77789

	tmp77790 := MakeNative(func(__e *ControlFlow) {
		V762 := __e.Get(1)
		_ = V762
		tmp77867 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77868 := Call(__e, tmp77867, V762)

		var ifres77837 Obj

		if True == tmp77868 {
			tmp77863 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77864 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77865 := Call(__e, tmp77864, V762)

			tmp77866 := Call(__e, tmp77863, symcompile, tmp77865)

			var ifres77839 Obj

			if True == tmp77866 {
				tmp77859 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77861 := Call(__e, tmp77860, V762)

				tmp77862 := Call(__e, tmp77859, tmp77861)

				var ifres77841 Obj

				if True == tmp77862 {
					tmp77853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp77854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77856 := Call(__e, tmp77855, V762)

					tmp77857 := Call(__e, tmp77854, tmp77856)

					tmp77858 := Call(__e, tmp77853, tmp77857)

					var ifres77843 Obj

					if True == tmp77858 {
						tmp77845 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp77846 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77847 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp77849 := Call(__e, tmp77848, V762)

						tmp77850 := Call(__e, tmp77847, tmp77849)

						tmp77851 := Call(__e, tmp77846, tmp77850)

						tmp77852 := Call(__e, tmp77845, Nil, tmp77851)

						var ifres77844 Obj

						if True == tmp77852 {
							ifres77844 = True

						} else {
							ifres77844 = False

						}

						ifres77843 = ifres77844

					} else {
						ifres77843 = False

					}

					var ifres77842 Obj

					if True == ifres77843 {
						ifres77842 = True

					} else {
						ifres77842 = False

					}

					ifres77841 = ifres77842

				} else {
					ifres77841 = False

				}

				var ifres77840 Obj

				if True == ifres77841 {
					ifres77840 = True

				} else {
					ifres77840 = False

				}

				ifres77839 = ifres77840

			} else {
				ifres77839 = False

			}

			var ifres77838 Obj

			if True == ifres77839 {
				ifres77838 = True

			} else {
				ifres77838 = False

			}

			ifres77837 = ifres77838

		} else {
			ifres77837 = False

		}

		if True == ifres77837 {
			tmp77792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77796 := Call(__e, tmp77795, V762)

			tmp77797 := Call(__e, tmp77794, tmp77796)

			tmp77798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77799 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77800 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77801 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77802 := Call(__e, tmp77801, V762)

			tmp77803 := Call(__e, tmp77800, tmp77802)

			tmp77804 := Call(__e, tmp77799, tmp77803)

			tmp77805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77807 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77808 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77809 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77810 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77811 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77813 := Call(__e, tmp77812, symE, Nil)

			tmp77814 := Call(__e, tmp77811, symcons_2, tmp77813)

			tmp77815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77817 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77818 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77819 := Call(__e, tmp77818, symE, Nil)

			tmp77820 := Call(__e, tmp77817, MakeString("parse error here: ~S~%"), tmp77819)

			tmp77821 := Call(__e, tmp77816, symerror, tmp77820)

			tmp77822 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77823 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77825 := Call(__e, tmp77824, MakeString("parse error~%"), Nil)

			tmp77826 := Call(__e, tmp77823, symerror, tmp77825)

			tmp77827 := Call(__e, tmp77822, tmp77826, Nil)

			tmp77828 := Call(__e, tmp77815, tmp77821, tmp77827)

			tmp77829 := Call(__e, tmp77810, tmp77814, tmp77828)

			tmp77830 := Call(__e, tmp77809, symif, tmp77829)

			tmp77831 := Call(__e, tmp77808, tmp77830, Nil)

			tmp77832 := Call(__e, tmp77807, symE, tmp77831)

			tmp77833 := Call(__e, tmp77806, symlambda, tmp77832)

			tmp77834 := Call(__e, tmp77805, tmp77833, Nil)

			tmp77835 := Call(__e, tmp77798, tmp77804, tmp77834)

			tmp77836 := Call(__e, tmp77793, tmp77797, tmp77835)

			__e.TailApply(tmp77792, symcompile, tmp77836)
			return

		} else {
			__e.Return(V762)
			return
		}

	}, 1)

	tmp77869 := Call(__e, PrimNS1Value(symns2_1set), symshen_4compile_1macro, tmp77790)

	_ = tmp77869

	tmp77870 := MakeNative(func(__e *ControlFlow) {
		V764 := __e.Get(1)
		_ = V764
		tmp77907 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77908 := Call(__e, tmp77907, V764)

		var ifres77901 Obj

		if True == tmp77908 {
			tmp77903 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77904 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77905 := Call(__e, tmp77904, V764)

			tmp77906 := Call(__e, tmp77903, symprolog_2, tmp77905)

			var ifres77902 Obj

			if True == tmp77906 {
				ifres77902 = True

			} else {
				ifres77902 = False

			}

			ifres77901 = ifres77902

		} else {
			ifres77901 = False

		}

		if True == ifres77901 {
			tmp77872 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77876 := Call(__e, tmp77875, symshen_4start_1new_1prolog_1process, Nil)

			tmp77877 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp77878 := MakeNative(func(__e *ControlFlow) {
				Calls := __e.Get(1)
				_ = Calls
				tmp77879 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					tmp77880 := MakeNative(func(__e *ControlFlow) {
						External := __e.Get(1)
						_ = External
						tmp77881 := MakeNative(func(__e *ControlFlow) {
							PrologVs := __e.Get(1)
							_ = PrologVs
							tmp77882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4locally_1bind_1prolog_1vs)

							__e.TailApply(tmp77882, symNPP, PrologVs, Calls)
							return

						}, 1)

						tmp77883 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

						tmp77884 := Call(__e, tmp77883, Vs, External)

						__e.TailApply(tmp77881, tmp77884)
						return

					}, 1)

					tmp77885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

					tmp77886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77887 := Call(__e, tmp77886, V764)

					tmp77888 := Call(__e, tmp77885, tmp77887)

					__e.TailApply(tmp77880, tmp77888)
					return

				}, 1)

				tmp77889 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				tmp77890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77891 := Call(__e, tmp77890, V764)

				tmp77892 := Call(__e, tmp77889, tmp77891)

				__e.TailApply(tmp77879, tmp77892)
				return

			}, 1)

			tmp77893 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

			tmp77894 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77895 := Call(__e, tmp77894, V764)

			tmp77896 := Call(__e, tmp77893, symNPP, tmp77895)

			tmp77897 := Call(__e, tmp77878, tmp77896)

			tmp77898 := Call(__e, tmp77877, tmp77897, Nil)

			tmp77899 := Call(__e, tmp77874, tmp77876, tmp77898)

			tmp77900 := Call(__e, tmp77873, symNPP, tmp77899)

			__e.TailApply(tmp77872, symlet, tmp77900)
			return

		} else {
			__e.Return(V764)
			return
		}

	}, 1)

	tmp77909 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1macro, tmp77870)

	_ = tmp77909

	tmp77910 := MakeNative(func(__e *ControlFlow) {
		V770 := __e.Get(1)
		_ = V770
		tmp77945 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77946 := Call(__e, tmp77945, V770)

		var ifres77925 Obj

		if True == tmp77946 {
			tmp77941 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77942 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77943 := Call(__e, tmp77942, V770)

			tmp77944 := Call(__e, tmp77941, symreceive, tmp77943)

			var ifres77927 Obj

			if True == tmp77944 {
				tmp77937 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77938 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77939 := Call(__e, tmp77938, V770)

				tmp77940 := Call(__e, tmp77937, tmp77939)

				var ifres77929 Obj

				if True == tmp77940 {
					tmp77931 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp77932 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77933 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77934 := Call(__e, tmp77933, V770)

					tmp77935 := Call(__e, tmp77932, tmp77934)

					tmp77936 := Call(__e, tmp77931, Nil, tmp77935)

					var ifres77930 Obj

					if True == tmp77936 {
						ifres77930 = True

					} else {
						ifres77930 = False

					}

					ifres77929 = ifres77930

				} else {
					ifres77929 = False

				}

				var ifres77928 Obj

				if True == ifres77929 {
					ifres77928 = True

				} else {
					ifres77928 = False

				}

				ifres77927 = ifres77928

			} else {
				ifres77927 = False

			}

			var ifres77926 Obj

			if True == ifres77927 {
				ifres77926 = True

			} else {
				ifres77926 = False

			}

			ifres77925 = ifres77926

		} else {
			ifres77925 = False

		}

		if True == ifres77925 {
			tmp77924 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(tmp77924, V770)
			return

		} else {
			tmp77922 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77923 := Call(__e, tmp77922, V770)

			if True == tmp77923 {
				tmp77913 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp77914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

				tmp77915 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77916 := Call(__e, tmp77915, V770)

				tmp77917 := Call(__e, tmp77914, tmp77916)

				tmp77918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

				tmp77919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77920 := Call(__e, tmp77919, V770)

				tmp77921 := Call(__e, tmp77918, tmp77920)

				__e.TailApply(tmp77913, tmp77917, tmp77921)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp77947 := Call(__e, PrimNS1Value(symns2_1set), symshen_4externally_1bound, tmp77910)

	_ = tmp77947

	tmp77948 := MakeNative(func(__e *ControlFlow) {
		V788 := __e.Get(1)
		_ = V788
		V789 := __e.Get(2)
		_ = V789
		V790 := __e.Get(3)
		_ = V790
		tmp77971 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp77972 := Call(__e, tmp77971, Nil, V789)

		if True == tmp77972 {
			__e.Return(V790)
			return
		} else {
			tmp77969 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77970 := Call(__e, tmp77969, V789)

			if True == tmp77970 {
				tmp77952 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77953 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77954 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77955 := Call(__e, tmp77954, V789)

				tmp77956 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77957 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77958 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77959 := Call(__e, tmp77958, V788, Nil)

				tmp77960 := Call(__e, tmp77957, symshen_4newpv, tmp77959)

				tmp77961 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77962 := Call(__e, PrimNS1Value(symns2_1value), symshen_4locally_1bind_1prolog_1vs)

				tmp77963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77964 := Call(__e, tmp77963, V789)

				tmp77965 := Call(__e, tmp77962, V788, tmp77964, V790)

				tmp77966 := Call(__e, tmp77961, tmp77965, Nil)

				tmp77967 := Call(__e, tmp77956, tmp77960, tmp77966)

				tmp77968 := Call(__e, tmp77953, tmp77955, tmp77967)

				__e.TailApply(tmp77952, symlet, tmp77968)
				return

			} else {
				tmp77951 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp77951, MakeString("implementation error inp locally-bind-prolog-vs"))
				return

			}

		}

	}, 3)

	tmp77973 := Call(__e, PrimNS1Value(symns2_1set), symshen_4locally_1bind_1prolog_1vs, tmp77948)

	_ = tmp77973

	tmp77974 := MakeNative(func(__e *ControlFlow) {
		V803 := __e.Get(1)
		_ = V803
		V804 := __e.Get(2)
		_ = V804
		tmp78333 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp78334 := Call(__e, tmp78333, Nil, V804)

		if True == tmp78334 {
			__e.Return(True)
			return
		} else {
			tmp78331 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp78332 := Call(__e, tmp78331, V804)

			var ifres78325 Obj

			if True == tmp78332 {
				tmp78327 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp78328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78329 := Call(__e, tmp78328, V804)

				tmp78330 := Call(__e, tmp78327, sym_b, tmp78329)

				var ifres78326 Obj

				if True == tmp78330 {
					ifres78326 = True

				} else {
					ifres78326 = False

				}

				ifres78325 = ifres78326

			} else {
				ifres78325 = False

			}

			if True == ifres78325 {
				tmp78310 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78311 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78312 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78314 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78316 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

				tmp78317 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78318 := Call(__e, tmp78317, V804)

				tmp78319 := Call(__e, tmp78316, V803, tmp78318)

				tmp78320 := Call(__e, tmp78315, tmp78319, Nil)

				tmp78321 := Call(__e, tmp78314, symfreeze, tmp78320)

				tmp78322 := Call(__e, tmp78313, tmp78321, Nil)

				tmp78323 := Call(__e, tmp78312, V803, tmp78322)

				tmp78324 := Call(__e, tmp78311, False, tmp78323)

				__e.TailApply(tmp78310, symcut, tmp78324)
				return

			} else {
				tmp78308 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78309 := Call(__e, tmp78308, V804)

				var ifres78276 Obj

				if True == tmp78309 {
					tmp78304 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78305 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78306 := Call(__e, tmp78305, V804)

					tmp78307 := Call(__e, tmp78304, tmp78306)

					var ifres78278 Obj

					if True == tmp78307 {
						tmp78298 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp78299 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78300 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78301 := Call(__e, tmp78300, V804)

						tmp78302 := Call(__e, tmp78299, tmp78301)

						tmp78303 := Call(__e, tmp78298, symwhen, tmp78302)

						var ifres78280 Obj

						if True == tmp78303 {
							tmp78292 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78293 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78294 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp78295 := Call(__e, tmp78294, V804)

							tmp78296 := Call(__e, tmp78293, tmp78295)

							tmp78297 := Call(__e, tmp78292, tmp78296)

							var ifres78282 Obj

							if True == tmp78297 {
								tmp78284 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp78285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78286 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78287 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78288 := Call(__e, tmp78287, V804)

								tmp78289 := Call(__e, tmp78286, tmp78288)

								tmp78290 := Call(__e, tmp78285, tmp78289)

								tmp78291 := Call(__e, tmp78284, Nil, tmp78290)

								var ifres78283 Obj

								if True == tmp78291 {
									ifres78283 = True

								} else {
									ifres78283 = False

								}

								ifres78282 = ifres78283

							} else {
								ifres78282 = False

							}

							var ifres78281 Obj

							if True == ifres78282 {
								ifres78281 = True

							} else {
								ifres78281 = False

							}

							ifres78280 = ifres78281

						} else {
							ifres78280 = False

						}

						var ifres78279 Obj

						if True == ifres78280 {
							ifres78279 = True

						} else {
							ifres78279 = False

						}

						ifres78278 = ifres78279

					} else {
						ifres78278 = False

					}

					var ifres78277 Obj

					if True == ifres78278 {
						ifres78277 = True

					} else {
						ifres78277 = False

					}

					ifres78276 = ifres78277

				} else {
					ifres78276 = False

				}

				if True == ifres78276 {
					tmp78253 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78254 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78255 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

					tmp78256 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78257 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78258 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78259 := Call(__e, tmp78258, V804)

					tmp78260 := Call(__e, tmp78257, tmp78259)

					tmp78261 := Call(__e, tmp78256, tmp78260)

					tmp78262 := Call(__e, tmp78255, tmp78261, V803)

					tmp78263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78265 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78266 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

					tmp78268 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78269 := Call(__e, tmp78268, V804)

					tmp78270 := Call(__e, tmp78267, V803, tmp78269)

					tmp78271 := Call(__e, tmp78266, tmp78270, Nil)

					tmp78272 := Call(__e, tmp78265, symfreeze, tmp78271)

					tmp78273 := Call(__e, tmp78264, tmp78272, Nil)

					tmp78274 := Call(__e, tmp78263, V803, tmp78273)

					tmp78275 := Call(__e, tmp78254, tmp78262, tmp78274)

					__e.TailApply(tmp78253, symfwhen, tmp78275)
					return

				} else {
					tmp78251 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78252 := Call(__e, tmp78251, V804)

					var ifres78207 Obj

					if True == tmp78252 {
						tmp78247 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78248 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78249 := Call(__e, tmp78248, V804)

						tmp78250 := Call(__e, tmp78247, tmp78249)

						var ifres78209 Obj

						if True == tmp78250 {
							tmp78241 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp78242 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp78243 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp78244 := Call(__e, tmp78243, V804)

							tmp78245 := Call(__e, tmp78242, tmp78244)

							tmp78246 := Call(__e, tmp78241, symis, tmp78245)

							var ifres78211 Obj

							if True == tmp78246 {
								tmp78235 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp78236 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78237 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78238 := Call(__e, tmp78237, V804)

								tmp78239 := Call(__e, tmp78236, tmp78238)

								tmp78240 := Call(__e, tmp78235, tmp78239)

								var ifres78213 Obj

								if True == tmp78240 {
									tmp78227 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp78228 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp78229 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp78230 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78231 := Call(__e, tmp78230, V804)

									tmp78232 := Call(__e, tmp78229, tmp78231)

									tmp78233 := Call(__e, tmp78228, tmp78232)

									tmp78234 := Call(__e, tmp78227, tmp78233)

									var ifres78215 Obj

									if True == tmp78234 {
										tmp78217 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp78218 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78219 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78220 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78221 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp78222 := Call(__e, tmp78221, V804)

										tmp78223 := Call(__e, tmp78220, tmp78222)

										tmp78224 := Call(__e, tmp78219, tmp78223)

										tmp78225 := Call(__e, tmp78218, tmp78224)

										tmp78226 := Call(__e, tmp78217, Nil, tmp78225)

										var ifres78216 Obj

										if True == tmp78226 {
											ifres78216 = True

										} else {
											ifres78216 = False

										}

										ifres78215 = ifres78216

									} else {
										ifres78215 = False

									}

									var ifres78214 Obj

									if True == ifres78215 {
										ifres78214 = True

									} else {
										ifres78214 = False

									}

									ifres78213 = ifres78214

								} else {
									ifres78213 = False

								}

								var ifres78212 Obj

								if True == ifres78213 {
									ifres78212 = True

								} else {
									ifres78212 = False

								}

								ifres78211 = ifres78212

							} else {
								ifres78211 = False

							}

							var ifres78210 Obj

							if True == ifres78211 {
								ifres78210 = True

							} else {
								ifres78210 = False

							}

							ifres78209 = ifres78210

						} else {
							ifres78209 = False

						}

						var ifres78208 Obj

						if True == ifres78209 {
							ifres78208 = True

						} else {
							ifres78208 = False

						}

						ifres78207 = ifres78208

					} else {
						ifres78207 = False

					}

					if True == ifres78207 {
						tmp78174 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78175 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78176 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78178 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78179 := Call(__e, tmp78178, V804)

						tmp78180 := Call(__e, tmp78177, tmp78179)

						tmp78181 := Call(__e, tmp78176, tmp78180)

						tmp78182 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

						tmp78184 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78187 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78188 := Call(__e, tmp78187, V804)

						tmp78189 := Call(__e, tmp78186, tmp78188)

						tmp78190 := Call(__e, tmp78185, tmp78189)

						tmp78191 := Call(__e, tmp78184, tmp78190)

						tmp78192 := Call(__e, tmp78183, tmp78191, V803)

						tmp78193 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78194 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78195 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78196 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78197 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

						tmp78198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78199 := Call(__e, tmp78198, V804)

						tmp78200 := Call(__e, tmp78197, V803, tmp78199)

						tmp78201 := Call(__e, tmp78196, tmp78200, Nil)

						tmp78202 := Call(__e, tmp78195, symfreeze, tmp78201)

						tmp78203 := Call(__e, tmp78194, tmp78202, Nil)

						tmp78204 := Call(__e, tmp78193, V803, tmp78203)

						tmp78205 := Call(__e, tmp78182, tmp78192, tmp78204)

						tmp78206 := Call(__e, tmp78175, tmp78181, tmp78205)

						__e.TailApply(tmp78174, symbind, tmp78206)
						return

					} else {
						tmp78172 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78173 := Call(__e, tmp78172, V804)

						var ifres78140 Obj

						if True == tmp78173 {
							tmp78168 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78169 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp78170 := Call(__e, tmp78169, V804)

							tmp78171 := Call(__e, tmp78168, tmp78170)

							var ifres78142 Obj

							if True == tmp78171 {
								tmp78162 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp78163 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78164 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78165 := Call(__e, tmp78164, V804)

								tmp78166 := Call(__e, tmp78163, tmp78165)

								tmp78167 := Call(__e, tmp78162, symreceive, tmp78166)

								var ifres78144 Obj

								if True == tmp78167 {
									tmp78156 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp78157 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp78158 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78159 := Call(__e, tmp78158, V804)

									tmp78160 := Call(__e, tmp78157, tmp78159)

									tmp78161 := Call(__e, tmp78156, tmp78160)

									var ifres78146 Obj

									if True == tmp78161 {
										tmp78148 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp78149 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78150 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78151 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp78152 := Call(__e, tmp78151, V804)

										tmp78153 := Call(__e, tmp78150, tmp78152)

										tmp78154 := Call(__e, tmp78149, tmp78153)

										tmp78155 := Call(__e, tmp78148, Nil, tmp78154)

										var ifres78147 Obj

										if True == tmp78155 {
											ifres78147 = True

										} else {
											ifres78147 = False

										}

										ifres78146 = ifres78147

									} else {
										ifres78146 = False

									}

									var ifres78145 Obj

									if True == ifres78146 {
										ifres78145 = True

									} else {
										ifres78145 = False

									}

									ifres78144 = ifres78145

								} else {
									ifres78144 = False

								}

								var ifres78143 Obj

								if True == ifres78144 {
									ifres78143 = True

								} else {
									ifres78143 = False

								}

								ifres78142 = ifres78143

							} else {
								ifres78142 = False

							}

							var ifres78141 Obj

							if True == ifres78142 {
								ifres78141 = True

							} else {
								ifres78141 = False

							}

							ifres78140 = ifres78141

						} else {
							ifres78140 = False

						}

						if True == ifres78140 {
							tmp78137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

							tmp78138 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78139 := Call(__e, tmp78138, V804)

							__e.TailApply(tmp78137, V803, tmp78139)
							return

						} else {
							tmp78135 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78136 := Call(__e, tmp78135, V804)

							var ifres78091 Obj

							if True == tmp78136 {
								tmp78131 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp78132 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78133 := Call(__e, tmp78132, V804)

								tmp78134 := Call(__e, tmp78131, tmp78133)

								var ifres78093 Obj

								if True == tmp78134 {
									tmp78125 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp78126 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78127 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78128 := Call(__e, tmp78127, V804)

									tmp78129 := Call(__e, tmp78126, tmp78128)

									tmp78130 := Call(__e, tmp78125, symbind, tmp78129)

									var ifres78095 Obj

									if True == tmp78130 {
										tmp78119 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp78120 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp78121 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp78122 := Call(__e, tmp78121, V804)

										tmp78123 := Call(__e, tmp78120, tmp78122)

										tmp78124 := Call(__e, tmp78119, tmp78123)

										var ifres78097 Obj

										if True == tmp78124 {
											tmp78111 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp78112 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp78113 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp78114 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp78115 := Call(__e, tmp78114, V804)

											tmp78116 := Call(__e, tmp78113, tmp78115)

											tmp78117 := Call(__e, tmp78112, tmp78116)

											tmp78118 := Call(__e, tmp78111, tmp78117)

											var ifres78099 Obj

											if True == tmp78118 {
												tmp78101 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp78102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp78103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp78104 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp78105 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp78106 := Call(__e, tmp78105, V804)

												tmp78107 := Call(__e, tmp78104, tmp78106)

												tmp78108 := Call(__e, tmp78103, tmp78107)

												tmp78109 := Call(__e, tmp78102, tmp78108)

												tmp78110 := Call(__e, tmp78101, Nil, tmp78109)

												var ifres78100 Obj

												if True == tmp78110 {
													ifres78100 = True

												} else {
													ifres78100 = False

												}

												ifres78099 = ifres78100

											} else {
												ifres78099 = False

											}

											var ifres78098 Obj

											if True == ifres78099 {
												ifres78098 = True

											} else {
												ifres78098 = False

											}

											ifres78097 = ifres78098

										} else {
											ifres78097 = False

										}

										var ifres78096 Obj

										if True == ifres78097 {
											ifres78096 = True

										} else {
											ifres78096 = False

										}

										ifres78095 = ifres78096

									} else {
										ifres78095 = False

									}

									var ifres78094 Obj

									if True == ifres78095 {
										ifres78094 = True

									} else {
										ifres78094 = False

									}

									ifres78093 = ifres78094

								} else {
									ifres78093 = False

								}

								var ifres78092 Obj

								if True == ifres78093 {
									ifres78092 = True

								} else {
									ifres78092 = False

								}

								ifres78091 = ifres78092

							} else {
								ifres78091 = False

							}

							if True == ifres78091 {
								tmp78058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78059 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78060 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78061 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78062 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78063 := Call(__e, tmp78062, V804)

								tmp78064 := Call(__e, tmp78061, tmp78063)

								tmp78065 := Call(__e, tmp78060, tmp78064)

								tmp78066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

								tmp78068 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78069 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78071 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78072 := Call(__e, tmp78071, V804)

								tmp78073 := Call(__e, tmp78070, tmp78072)

								tmp78074 := Call(__e, tmp78069, tmp78073)

								tmp78075 := Call(__e, tmp78068, tmp78074)

								tmp78076 := Call(__e, tmp78067, tmp78075, V803)

								tmp78077 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78078 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp78081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

								tmp78082 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78083 := Call(__e, tmp78082, V804)

								tmp78084 := Call(__e, tmp78081, V803, tmp78083)

								tmp78085 := Call(__e, tmp78080, tmp78084, Nil)

								tmp78086 := Call(__e, tmp78079, symfreeze, tmp78085)

								tmp78087 := Call(__e, tmp78078, tmp78086, Nil)

								tmp78088 := Call(__e, tmp78077, V803, tmp78087)

								tmp78089 := Call(__e, tmp78066, tmp78076, tmp78088)

								tmp78090 := Call(__e, tmp78059, tmp78065, tmp78089)

								__e.TailApply(tmp78058, symbind, tmp78090)
								return

							} else {
								tmp78056 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp78057 := Call(__e, tmp78056, V804)

								var ifres78024 Obj

								if True == tmp78057 {
									tmp78052 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp78053 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78054 := Call(__e, tmp78053, V804)

									tmp78055 := Call(__e, tmp78052, tmp78054)

									var ifres78026 Obj

									if True == tmp78055 {
										tmp78046 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp78047 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp78048 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp78049 := Call(__e, tmp78048, V804)

										tmp78050 := Call(__e, tmp78047, tmp78049)

										tmp78051 := Call(__e, tmp78046, symfwhen, tmp78050)

										var ifres78028 Obj

										if True == tmp78051 {
											tmp78040 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp78041 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp78042 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp78043 := Call(__e, tmp78042, V804)

											tmp78044 := Call(__e, tmp78041, tmp78043)

											tmp78045 := Call(__e, tmp78040, tmp78044)

											var ifres78030 Obj

											if True == tmp78045 {
												tmp78032 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp78033 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp78034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp78035 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp78036 := Call(__e, tmp78035, V804)

												tmp78037 := Call(__e, tmp78034, tmp78036)

												tmp78038 := Call(__e, tmp78033, tmp78037)

												tmp78039 := Call(__e, tmp78032, Nil, tmp78038)

												var ifres78031 Obj

												if True == tmp78039 {
													ifres78031 = True

												} else {
													ifres78031 = False

												}

												ifres78030 = ifres78031

											} else {
												ifres78030 = False

											}

											var ifres78029 Obj

											if True == ifres78030 {
												ifres78029 = True

											} else {
												ifres78029 = False

											}

											ifres78028 = ifres78029

										} else {
											ifres78028 = False

										}

										var ifres78027 Obj

										if True == ifres78028 {
											ifres78027 = True

										} else {
											ifres78027 = False

										}

										ifres78026 = ifres78027

									} else {
										ifres78026 = False

									}

									var ifres78025 Obj

									if True == ifres78026 {
										ifres78025 = True

									} else {
										ifres78025 = False

									}

									ifres78024 = ifres78025

								} else {
									ifres78024 = False

								}

								if True == ifres78024 {
									tmp78001 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78003 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

									tmp78004 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78005 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp78006 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp78007 := Call(__e, tmp78006, V804)

									tmp78008 := Call(__e, tmp78005, tmp78007)

									tmp78009 := Call(__e, tmp78004, tmp78008)

									tmp78010 := Call(__e, tmp78003, tmp78009, V803)

									tmp78011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp78015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

									tmp78016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp78017 := Call(__e, tmp78016, V804)

									tmp78018 := Call(__e, tmp78015, V803, tmp78017)

									tmp78019 := Call(__e, tmp78014, tmp78018, Nil)

									tmp78020 := Call(__e, tmp78013, symfreeze, tmp78019)

									tmp78021 := Call(__e, tmp78012, tmp78020, Nil)

									tmp78022 := Call(__e, tmp78011, V803, tmp78021)

									tmp78023 := Call(__e, tmp78002, tmp78010, tmp78022)

									__e.TailApply(tmp78001, symfwhen, tmp78023)
									return

								} else {
									tmp77999 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp78000 := Call(__e, tmp77999, V804)

									if True == tmp78000 {
										tmp77984 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										tmp77985 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp77986 := Call(__e, tmp77985, V804)

										tmp77987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp77988 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp77989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp77990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp77991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

										tmp77992 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp77993 := Call(__e, tmp77992, V804)

										tmp77994 := Call(__e, tmp77991, V803, tmp77993)

										tmp77995 := Call(__e, tmp77990, tmp77994, Nil)

										tmp77996 := Call(__e, tmp77989, symfreeze, tmp77995)

										tmp77997 := Call(__e, tmp77988, tmp77996, Nil)

										tmp77998 := Call(__e, tmp77987, V803, tmp77997)

										__e.TailApply(tmp77984, tmp77986, tmp77998)
										return

									} else {
										tmp77983 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

										__e.TailApply(tmp77983, MakeString("implementation error in bld-prolog-call"))
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

	tmp78335 := Call(__e, PrimNS1Value(symns2_1set), symshen_4bld_1prolog_1call, tmp77974)

	_ = tmp78335

	tmp78336 := MakeNative(func(__e *ControlFlow) {
		V806 := __e.Get(1)
		_ = V806
		tmp78361 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78362 := Call(__e, tmp78361, V806)

		var ifres78349 Obj

		if True == tmp78362 {
			tmp78357 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78359 := Call(__e, tmp78358, V806)

			tmp78360 := Call(__e, tmp78357, symdefprolog, tmp78359)

			var ifres78351 Obj

			if True == tmp78360 {
				tmp78353 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78354 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78355 := Call(__e, tmp78354, V806)

				tmp78356 := Call(__e, tmp78353, tmp78355)

				var ifres78352 Obj

				if True == tmp78356 {
					ifres78352 = True

				} else {
					ifres78352 = False

				}

				ifres78351 = ifres78352

			} else {
				ifres78351 = False

			}

			var ifres78350 Obj

			if True == ifres78351 {
				ifres78350 = True

			} else {
				ifres78350 = False

			}

			ifres78349 = ifres78350

		} else {
			ifres78349 = False

		}

		if True == ifres78349 {
			tmp78338 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			tmp78339 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp78340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5defprolog_6)

				__e.TailApply(tmp78340, Y)
				return

			}, 1)

			tmp78341 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78342 := Call(__e, tmp78341, V806)

			tmp78343 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp78344 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1error)

				tmp78345 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78347 := Call(__e, tmp78346, V806)

				tmp78348 := Call(__e, tmp78345, tmp78347)

				__e.TailApply(tmp78344, tmp78348, Y)
				return

			}, 1)

			__e.TailApply(tmp78338, tmp78339, tmp78342, tmp78343)
			return

		} else {
			__e.Return(V806)
			return
		}

	}, 1)

	tmp78363 := Call(__e, PrimNS1Value(symns2_1set), symshen_4defprolog_1macro, tmp78336)

	_ = tmp78363

	tmp78364 := MakeNative(func(__e *ControlFlow) {
		V808 := __e.Get(1)
		_ = V808
		tmp78417 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78418 := Call(__e, tmp78417, V808)

		var ifres78405 Obj

		if True == tmp78418 {
			tmp78413 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78414 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78415 := Call(__e, tmp78414, V808)

			tmp78416 := Call(__e, tmp78413, symdatatype, tmp78415)

			var ifres78407 Obj

			if True == tmp78416 {
				tmp78409 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78410 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78411 := Call(__e, tmp78410, V808)

				tmp78412 := Call(__e, tmp78409, tmp78411)

				var ifres78408 Obj

				if True == tmp78412 {
					ifres78408 = True

				} else {
					ifres78408 = False

				}

				ifres78407 = ifres78408

			} else {
				ifres78407 = False

			}

			var ifres78406 Obj

			if True == ifres78407 {
				ifres78406 = True

			} else {
				ifres78406 = False

			}

			ifres78405 = ifres78406

		} else {
			ifres78405 = False

		}

		if True == ifres78405 {
			tmp78366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			tmp78369 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78370 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78371 := Call(__e, tmp78370, V808)

			tmp78372 := Call(__e, tmp78369, tmp78371)

			tmp78373 := Call(__e, tmp78368, tmp78372)

			tmp78374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78376 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78377 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78379 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78380 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78381 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78382 := Call(__e, tmp78381, symX, Nil)

			tmp78383 := Call(__e, tmp78380, symshen_4_5datatype_1rules_6, tmp78382)

			tmp78384 := Call(__e, tmp78379, tmp78383, Nil)

			tmp78385 := Call(__e, tmp78378, symX, tmp78384)

			tmp78386 := Call(__e, tmp78377, symlambda, tmp78385)

			tmp78387 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp78389 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78390 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78391 := Call(__e, tmp78390, V808)

			tmp78392 := Call(__e, tmp78389, tmp78391)

			tmp78393 := Call(__e, tmp78388, tmp78392)

			tmp78394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78396 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78397 := Call(__e, tmp78396, symshen_4datatype_1error, Nil)

			tmp78398 := Call(__e, tmp78395, symfunction, tmp78397)

			tmp78399 := Call(__e, tmp78394, tmp78398, Nil)

			tmp78400 := Call(__e, tmp78387, tmp78393, tmp78399)

			tmp78401 := Call(__e, tmp78376, tmp78386, tmp78400)

			tmp78402 := Call(__e, tmp78375, symcompile, tmp78401)

			tmp78403 := Call(__e, tmp78374, tmp78402, Nil)

			tmp78404 := Call(__e, tmp78367, tmp78373, tmp78403)

			__e.TailApply(tmp78366, symshen_4process_1datatype, tmp78404)
			return

		} else {
			__e.Return(V808)
			return
		}

	}, 1)

	tmp78419 := Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1macro, tmp78364)

	_ = tmp78419

	tmp78420 := MakeNative(func(__e *ControlFlow) {
		V810 := __e.Get(1)
		_ = V810
		tmp78421 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		tmp78422 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp78423 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		tmp78424 := Call(__e, tmp78423, V810)

		tmp78425 := Call(__e, tmp78422, tmp78424, MakeString("#type"))

		__e.TailApply(tmp78421, tmp78425)
		return

	}, 1)

	tmp78426 := Call(__e, PrimNS1Value(symns2_1set), symshen_4intern_1type, tmp78420)

	_ = tmp78426

	tmp78427 := MakeNative(func(__e *ControlFlow) {
		V812 := __e.Get(1)
		_ = V812
		tmp78538 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78539 := Call(__e, tmp78538, V812)

		var ifres78508 Obj

		if True == tmp78539 {
			tmp78534 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78536 := Call(__e, tmp78535, V812)

			tmp78537 := Call(__e, tmp78534, sym_8s, tmp78536)

			var ifres78510 Obj

			if True == tmp78537 {
				tmp78530 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78531 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78532 := Call(__e, tmp78531, V812)

				tmp78533 := Call(__e, tmp78530, tmp78532)

				var ifres78512 Obj

				if True == tmp78533 {
					tmp78524 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78525 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78526 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78527 := Call(__e, tmp78526, V812)

					tmp78528 := Call(__e, tmp78525, tmp78527)

					tmp78529 := Call(__e, tmp78524, tmp78528)

					var ifres78514 Obj

					if True == tmp78529 {
						tmp78516 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78517 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78518 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78519 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78520 := Call(__e, tmp78519, V812)

						tmp78521 := Call(__e, tmp78518, tmp78520)

						tmp78522 := Call(__e, tmp78517, tmp78521)

						tmp78523 := Call(__e, tmp78516, tmp78522)

						var ifres78515 Obj

						if True == tmp78523 {
							ifres78515 = True

						} else {
							ifres78515 = False

						}

						ifres78514 = ifres78515

					} else {
						ifres78514 = False

					}

					var ifres78513 Obj

					if True == ifres78514 {
						ifres78513 = True

					} else {
						ifres78513 = False

					}

					ifres78512 = ifres78513

				} else {
					ifres78512 = False

				}

				var ifres78511 Obj

				if True == ifres78512 {
					ifres78511 = True

				} else {
					ifres78511 = False

				}

				ifres78510 = ifres78511

			} else {
				ifres78510 = False

			}

			var ifres78509 Obj

			if True == ifres78510 {
				ifres78509 = True

			} else {
				ifres78509 = False

			}

			ifres78508 = ifres78509

		} else {
			ifres78508 = False

		}

		if True == ifres78508 {
			tmp78491 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78492 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78493 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78494 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78495 := Call(__e, tmp78494, V812)

			tmp78496 := Call(__e, tmp78493, tmp78495)

			tmp78497 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8s_1macro)

			tmp78499 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78500 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78501 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78502 := Call(__e, tmp78501, V812)

			tmp78503 := Call(__e, tmp78500, tmp78502)

			tmp78504 := Call(__e, tmp78499, sym_8s, tmp78503)

			tmp78505 := Call(__e, tmp78498, tmp78504)

			tmp78506 := Call(__e, tmp78497, tmp78505, Nil)

			tmp78507 := Call(__e, tmp78492, tmp78496, tmp78506)

			__e.TailApply(tmp78491, sym_8s, tmp78507)
			return

		} else {
			tmp78489 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp78490 := Call(__e, tmp78489, V812)

			var ifres78451 Obj

			if True == tmp78490 {
				tmp78485 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp78486 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78487 := Call(__e, tmp78486, V812)

				tmp78488 := Call(__e, tmp78485, sym_8s, tmp78487)

				var ifres78453 Obj

				if True == tmp78488 {
					tmp78481 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78482 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78483 := Call(__e, tmp78482, V812)

					tmp78484 := Call(__e, tmp78481, tmp78483)

					var ifres78455 Obj

					if True == tmp78484 {
						tmp78475 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78478 := Call(__e, tmp78477, V812)

						tmp78479 := Call(__e, tmp78476, tmp78478)

						tmp78480 := Call(__e, tmp78475, tmp78479)

						var ifres78457 Obj

						if True == tmp78480 {
							tmp78467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp78468 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78469 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78470 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78471 := Call(__e, tmp78470, V812)

							tmp78472 := Call(__e, tmp78469, tmp78471)

							tmp78473 := Call(__e, tmp78468, tmp78472)

							tmp78474 := Call(__e, tmp78467, Nil, tmp78473)

							var ifres78459 Obj

							if True == tmp78474 {
								tmp78461 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

								tmp78462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp78463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78464 := Call(__e, tmp78463, V812)

								tmp78465 := Call(__e, tmp78462, tmp78464)

								tmp78466 := Call(__e, tmp78461, tmp78465)

								var ifres78460 Obj

								if True == tmp78466 {
									ifres78460 = True

								} else {
									ifres78460 = False

								}

								ifres78459 = ifres78460

							} else {
								ifres78459 = False

							}

							var ifres78458 Obj

							if True == ifres78459 {
								ifres78458 = True

							} else {
								ifres78458 = False

							}

							ifres78457 = ifres78458

						} else {
							ifres78457 = False

						}

						var ifres78456 Obj

						if True == ifres78457 {
							ifres78456 = True

						} else {
							ifres78456 = False

						}

						ifres78455 = ifres78456

					} else {
						ifres78455 = False

					}

					var ifres78454 Obj

					if True == ifres78455 {
						ifres78454 = True

					} else {
						ifres78454 = False

					}

					ifres78453 = ifres78454

				} else {
					ifres78453 = False

				}

				var ifres78452 Obj

				if True == ifres78453 {
					ifres78452 = True

				} else {
					ifres78452 = False

				}

				ifres78451 = ifres78452

			} else {
				ifres78451 = False

			}

			if True == ifres78451 {
				tmp78430 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					tmp78441 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					tmp78442 := Call(__e, PrimNS1Value(symns2_1value), symlength)

					tmp78443 := Call(__e, tmp78442, E)

					tmp78444 := Call(__e, tmp78441, tmp78443, MakeNumber(1))

					if True == tmp78444 {
						tmp78432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8s_1macro)

						tmp78433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78434 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						tmp78435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78436 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78437 := Call(__e, tmp78436, V812)

						tmp78438 := Call(__e, tmp78435, tmp78437)

						tmp78439 := Call(__e, tmp78434, E, tmp78438)

						tmp78440 := Call(__e, tmp78433, sym_8s, tmp78439)

						__e.TailApply(tmp78432, tmp78440)
						return

					} else {
						__e.Return(V812)
						return
					}

				}, 1)

				tmp78445 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

				tmp78446 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78448 := Call(__e, tmp78447, V812)

				tmp78449 := Call(__e, tmp78446, tmp78448)

				tmp78450 := Call(__e, tmp78445, tmp78449)

				__e.TailApply(tmp78430, tmp78450)
				return

			} else {
				__e.Return(V812)
				return
			}

		}

	}, 1)

	tmp78540 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_8s_1macro, tmp78427)

	_ = tmp78540

	tmp78541 := MakeNative(func(__e *ControlFlow) {
		V814 := __e.Get(1)
		_ = V814
		tmp78558 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78559 := Call(__e, tmp78558, V814)

		var ifres78552 Obj

		if True == tmp78559 {
			tmp78554 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78555 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78556 := Call(__e, tmp78555, V814)

			tmp78557 := Call(__e, tmp78554, symsynonyms, tmp78556)

			var ifres78553 Obj

			if True == tmp78557 {
				ifres78553 = True

			} else {
				ifres78553 = False

			}

			ifres78552 = ifres78553

		} else {
			ifres78552 = False

		}

		if True == ifres78552 {
			tmp78543 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp78546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1synonyms)

			tmp78547 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78548 := Call(__e, tmp78547, V814)

			tmp78549 := Call(__e, tmp78546, tmp78548)

			tmp78550 := Call(__e, tmp78545, tmp78549)

			tmp78551 := Call(__e, tmp78544, tmp78550, Nil)

			__e.TailApply(tmp78543, symshen_4synonyms_1help, tmp78551)
			return

		} else {
			__e.Return(V814)
			return
		}

	}, 1)

	tmp78560 := Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1macro, tmp78541)

	_ = tmp78560

	tmp78561 := MakeNative(func(__e *ControlFlow) {
		V816 := __e.Get(1)
		_ = V816
		tmp78562 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp78563 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp78564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

			__e.TailApply(tmp78564, X)
			return

		}, 1)

		__e.TailApply(tmp78562, tmp78563, V816)
		return

	}, 1)

	tmp78565 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1synonyms, tmp78561)

	_ = tmp78565

	tmp78566 := MakeNative(func(__e *ControlFlow) {
		V818 := __e.Get(1)
		_ = V818
		tmp78583 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78584 := Call(__e, tmp78583, V818)

		var ifres78571 Obj

		if True == tmp78584 {
			tmp78579 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78580 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78581 := Call(__e, tmp78580, V818)

			tmp78582 := Call(__e, tmp78579, symnl, tmp78581)

			var ifres78573 Obj

			if True == tmp78582 {
				tmp78575 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp78576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78577 := Call(__e, tmp78576, V818)

				tmp78578 := Call(__e, tmp78575, Nil, tmp78577)

				var ifres78574 Obj

				if True == tmp78578 {
					ifres78574 = True

				} else {
					ifres78574 = False

				}

				ifres78573 = ifres78574

			} else {
				ifres78573 = False

			}

			var ifres78572 Obj

			if True == ifres78573 {
				ifres78572 = True

			} else {
				ifres78572 = False

			}

			ifres78571 = ifres78572

		} else {
			ifres78571 = False

		}

		if True == ifres78571 {
			tmp78568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78569 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78570 := Call(__e, tmp78569, MakeNumber(1), Nil)

			__e.TailApply(tmp78568, symnl, tmp78570)
			return

		} else {
			__e.Return(V818)
			return
		}

	}, 1)

	tmp78585 := Call(__e, PrimNS1Value(symns2_1set), symshen_4nl_1macro, tmp78566)

	_ = tmp78585

	tmp78586 := MakeNative(func(__e *ControlFlow) {
		V820 := __e.Get(1)
		_ = V820
		tmp78655 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78656 := Call(__e, tmp78655, V820)

		var ifres78609 Obj

		if True == tmp78656 {
			tmp78651 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp78652 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78653 := Call(__e, tmp78652, V820)

			tmp78654 := Call(__e, tmp78651, tmp78653)

			var ifres78611 Obj

			if True == tmp78654 {
				tmp78645 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78648 := Call(__e, tmp78647, V820)

				tmp78649 := Call(__e, tmp78646, tmp78648)

				tmp78650 := Call(__e, tmp78645, tmp78649)

				var ifres78613 Obj

				if True == tmp78650 {
					tmp78637 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78638 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78639 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78640 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78641 := Call(__e, tmp78640, V820)

					tmp78642 := Call(__e, tmp78639, tmp78641)

					tmp78643 := Call(__e, tmp78638, tmp78642)

					tmp78644 := Call(__e, tmp78637, tmp78643)

					var ifres78615 Obj

					if True == tmp78644 {
						tmp78617 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

						tmp78618 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78619 := Call(__e, tmp78618, V820)

						tmp78620 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78621 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78625 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78626 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp78628 := Call(__e, tmp78627, symdo, Nil)

						tmp78629 := Call(__e, tmp78626, sym_d, tmp78628)

						tmp78630 := Call(__e, tmp78625, sym_7, tmp78629)

						tmp78631 := Call(__e, tmp78624, symor, tmp78630)

						tmp78632 := Call(__e, tmp78623, symand, tmp78631)

						tmp78633 := Call(__e, tmp78622, symappend, tmp78632)

						tmp78634 := Call(__e, tmp78621, sym_8v, tmp78633)

						tmp78635 := Call(__e, tmp78620, sym_8p, tmp78634)

						tmp78636 := Call(__e, tmp78617, tmp78619, tmp78635)

						var ifres78616 Obj

						if True == tmp78636 {
							ifres78616 = True

						} else {
							ifres78616 = False

						}

						ifres78615 = ifres78616

					} else {
						ifres78615 = False

					}

					var ifres78614 Obj

					if True == ifres78615 {
						ifres78614 = True

					} else {
						ifres78614 = False

					}

					ifres78613 = ifres78614

				} else {
					ifres78613 = False

				}

				var ifres78612 Obj

				if True == ifres78613 {
					ifres78612 = True

				} else {
					ifres78612 = False

				}

				ifres78611 = ifres78612

			} else {
				ifres78611 = False

			}

			var ifres78610 Obj

			if True == ifres78611 {
				ifres78610 = True

			} else {
				ifres78610 = False

			}

			ifres78609 = ifres78610

		} else {
			ifres78609 = False

		}

		if True == ifres78609 {
			tmp78588 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78589 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78590 := Call(__e, tmp78589, V820)

			tmp78591 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78592 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78593 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78594 := Call(__e, tmp78593, V820)

			tmp78595 := Call(__e, tmp78592, tmp78594)

			tmp78596 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1macro)

			tmp78598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78599 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78600 := Call(__e, tmp78599, V820)

			tmp78601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78603 := Call(__e, tmp78602, V820)

			tmp78604 := Call(__e, tmp78601, tmp78603)

			tmp78605 := Call(__e, tmp78598, tmp78600, tmp78604)

			tmp78606 := Call(__e, tmp78597, tmp78605)

			tmp78607 := Call(__e, tmp78596, tmp78606, Nil)

			tmp78608 := Call(__e, tmp78591, tmp78595, tmp78607)

			__e.TailApply(tmp78588, tmp78590, tmp78608)
			return

		} else {
			__e.Return(V820)
			return
		}

	}, 1)

	tmp78657 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1macro, tmp78586)

	_ = tmp78657

	tmp78658 := MakeNative(func(__e *ControlFlow) {
		V822 := __e.Get(1)
		_ = V822
		tmp78729 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78730 := Call(__e, tmp78729, V822)

		var ifres78687 Obj

		if True == tmp78730 {
			tmp78725 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78726 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78727 := Call(__e, tmp78726, V822)

			tmp78728 := Call(__e, tmp78725, symlet, tmp78727)

			var ifres78689 Obj

			if True == tmp78728 {
				tmp78721 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78722 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78723 := Call(__e, tmp78722, V822)

				tmp78724 := Call(__e, tmp78721, tmp78723)

				var ifres78691 Obj

				if True == tmp78724 {
					tmp78715 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78716 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78717 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78718 := Call(__e, tmp78717, V822)

					tmp78719 := Call(__e, tmp78716, tmp78718)

					tmp78720 := Call(__e, tmp78715, tmp78719)

					var ifres78693 Obj

					if True == tmp78720 {
						tmp78707 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78708 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78709 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78710 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78711 := Call(__e, tmp78710, V822)

						tmp78712 := Call(__e, tmp78709, tmp78711)

						tmp78713 := Call(__e, tmp78708, tmp78712)

						tmp78714 := Call(__e, tmp78707, tmp78713)

						var ifres78695 Obj

						if True == tmp78714 {
							tmp78697 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78698 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78699 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78701 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78702 := Call(__e, tmp78701, V822)

							tmp78703 := Call(__e, tmp78700, tmp78702)

							tmp78704 := Call(__e, tmp78699, tmp78703)

							tmp78705 := Call(__e, tmp78698, tmp78704)

							tmp78706 := Call(__e, tmp78697, tmp78705)

							var ifres78696 Obj

							if True == tmp78706 {
								ifres78696 = True

							} else {
								ifres78696 = False

							}

							ifres78695 = ifres78696

						} else {
							ifres78695 = False

						}

						var ifres78694 Obj

						if True == ifres78695 {
							ifres78694 = True

						} else {
							ifres78694 = False

						}

						ifres78693 = ifres78694

					} else {
						ifres78693 = False

					}

					var ifres78692 Obj

					if True == ifres78693 {
						ifres78692 = True

					} else {
						ifres78692 = False

					}

					ifres78691 = ifres78692

				} else {
					ifres78691 = False

				}

				var ifres78690 Obj

				if True == ifres78691 {
					ifres78690 = True

				} else {
					ifres78690 = False

				}

				ifres78689 = ifres78690

			} else {
				ifres78689 = False

			}

			var ifres78688 Obj

			if True == ifres78689 {
				ifres78688 = True

			} else {
				ifres78688 = False

			}

			ifres78687 = ifres78688

		} else {
			ifres78687 = False

		}

		if True == ifres78687 {
			tmp78660 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78662 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78663 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78664 := Call(__e, tmp78663, V822)

			tmp78665 := Call(__e, tmp78662, tmp78664)

			tmp78666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78667 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78669 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78670 := Call(__e, tmp78669, V822)

			tmp78671 := Call(__e, tmp78668, tmp78670)

			tmp78672 := Call(__e, tmp78667, tmp78671)

			tmp78673 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4let_1macro)

			tmp78675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78676 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78677 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78678 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78679 := Call(__e, tmp78678, V822)

			tmp78680 := Call(__e, tmp78677, tmp78679)

			tmp78681 := Call(__e, tmp78676, tmp78680)

			tmp78682 := Call(__e, tmp78675, symlet, tmp78681)

			tmp78683 := Call(__e, tmp78674, tmp78682)

			tmp78684 := Call(__e, tmp78673, tmp78683, Nil)

			tmp78685 := Call(__e, tmp78666, tmp78672, tmp78684)

			tmp78686 := Call(__e, tmp78661, tmp78665, tmp78685)

			__e.TailApply(tmp78660, symlet, tmp78686)
			return

		} else {
			__e.Return(V822)
			return
		}

	}, 1)

	tmp78731 := Call(__e, PrimNS1Value(symns2_1set), symshen_4let_1macro, tmp78658)

	_ = tmp78731

	tmp78732 := MakeNative(func(__e *ControlFlow) {
		V824 := __e.Get(1)
		_ = V824
		tmp78817 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78818 := Call(__e, tmp78817, V824)

		var ifres78787 Obj

		if True == tmp78818 {
			tmp78813 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78814 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78815 := Call(__e, tmp78814, V824)

			tmp78816 := Call(__e, tmp78813, sym_c_4, tmp78815)

			var ifres78789 Obj

			if True == tmp78816 {
				tmp78809 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78811 := Call(__e, tmp78810, V824)

				tmp78812 := Call(__e, tmp78809, tmp78811)

				var ifres78791 Obj

				if True == tmp78812 {
					tmp78803 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78804 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78805 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78806 := Call(__e, tmp78805, V824)

					tmp78807 := Call(__e, tmp78804, tmp78806)

					tmp78808 := Call(__e, tmp78803, tmp78807)

					var ifres78793 Obj

					if True == tmp78808 {
						tmp78795 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78799 := Call(__e, tmp78798, V824)

						tmp78800 := Call(__e, tmp78797, tmp78799)

						tmp78801 := Call(__e, tmp78796, tmp78800)

						tmp78802 := Call(__e, tmp78795, tmp78801)

						var ifres78794 Obj

						if True == tmp78802 {
							ifres78794 = True

						} else {
							ifres78794 = False

						}

						ifres78793 = ifres78794

					} else {
						ifres78793 = False

					}

					var ifres78792 Obj

					if True == ifres78793 {
						ifres78792 = True

					} else {
						ifres78792 = False

					}

					ifres78791 = ifres78792

				} else {
					ifres78791 = False

				}

				var ifres78790 Obj

				if True == ifres78791 {
					ifres78790 = True

				} else {
					ifres78790 = False

				}

				ifres78789 = ifres78790

			} else {
				ifres78789 = False

			}

			var ifres78788 Obj

			if True == ifres78789 {
				ifres78788 = True

			} else {
				ifres78788 = False

			}

			ifres78787 = ifres78788

		} else {
			ifres78787 = False

		}

		if True == ifres78787 {
			tmp78770 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78772 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78773 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78774 := Call(__e, tmp78773, V824)

			tmp78775 := Call(__e, tmp78772, tmp78774)

			tmp78776 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abs_1macro)

			tmp78778 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78779 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78781 := Call(__e, tmp78780, V824)

			tmp78782 := Call(__e, tmp78779, tmp78781)

			tmp78783 := Call(__e, tmp78778, sym_c_4, tmp78782)

			tmp78784 := Call(__e, tmp78777, tmp78783)

			tmp78785 := Call(__e, tmp78776, tmp78784, Nil)

			tmp78786 := Call(__e, tmp78771, tmp78775, tmp78785)

			__e.TailApply(tmp78770, symlambda, tmp78786)
			return

		} else {
			tmp78768 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp78769 := Call(__e, tmp78768, V824)

			var ifres78738 Obj

			if True == tmp78769 {
				tmp78764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp78765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78766 := Call(__e, tmp78765, V824)

				tmp78767 := Call(__e, tmp78764, sym_c_4, tmp78766)

				var ifres78740 Obj

				if True == tmp78767 {
					tmp78760 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78762 := Call(__e, tmp78761, V824)

					tmp78763 := Call(__e, tmp78760, tmp78762)

					var ifres78742 Obj

					if True == tmp78763 {
						tmp78754 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78756 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78757 := Call(__e, tmp78756, V824)

						tmp78758 := Call(__e, tmp78755, tmp78757)

						tmp78759 := Call(__e, tmp78754, tmp78758)

						var ifres78744 Obj

						if True == tmp78759 {
							tmp78746 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp78747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78748 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78749 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78750 := Call(__e, tmp78749, V824)

							tmp78751 := Call(__e, tmp78748, tmp78750)

							tmp78752 := Call(__e, tmp78747, tmp78751)

							tmp78753 := Call(__e, tmp78746, Nil, tmp78752)

							var ifres78745 Obj

							if True == tmp78753 {
								ifres78745 = True

							} else {
								ifres78745 = False

							}

							ifres78744 = ifres78745

						} else {
							ifres78744 = False

						}

						var ifres78743 Obj

						if True == ifres78744 {
							ifres78743 = True

						} else {
							ifres78743 = False

						}

						ifres78742 = ifres78743

					} else {
						ifres78742 = False

					}

					var ifres78741 Obj

					if True == ifres78742 {
						ifres78741 = True

					} else {
						ifres78741 = False

					}

					ifres78740 = ifres78741

				} else {
					ifres78740 = False

				}

				var ifres78739 Obj

				if True == ifres78740 {
					ifres78739 = True

				} else {
					ifres78739 = False

				}

				ifres78738 = ifres78739

			} else {
				ifres78738 = False

			}

			if True == ifres78738 {
				tmp78735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78736 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78737 := Call(__e, tmp78736, V824)

				__e.TailApply(tmp78735, symlambda, tmp78737)
				return

			} else {
				__e.Return(V824)
				return
			}

		}

	}, 1)

	tmp78819 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abs_1macro, tmp78732)

	_ = tmp78819

	tmp78820 := MakeNative(func(__e *ControlFlow) {
		V828 := __e.Get(1)
		_ = V828
		tmp78983 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp78984 := Call(__e, tmp78983, V828)

		var ifres78955 Obj

		if True == tmp78984 {
			tmp78979 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp78980 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78981 := Call(__e, tmp78980, V828)

			tmp78982 := Call(__e, tmp78979, symcases, tmp78981)

			var ifres78957 Obj

			if True == tmp78982 {
				tmp78975 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78976 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78977 := Call(__e, tmp78976, V828)

				tmp78978 := Call(__e, tmp78975, tmp78977)

				var ifres78959 Obj

				if True == tmp78978 {
					tmp78969 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp78970 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78972 := Call(__e, tmp78971, V828)

					tmp78973 := Call(__e, tmp78970, tmp78972)

					tmp78974 := Call(__e, tmp78969, True, tmp78973)

					var ifres78961 Obj

					if True == tmp78974 {
						tmp78963 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78964 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78966 := Call(__e, tmp78965, V828)

						tmp78967 := Call(__e, tmp78964, tmp78966)

						tmp78968 := Call(__e, tmp78963, tmp78967)

						var ifres78962 Obj

						if True == tmp78968 {
							ifres78962 = True

						} else {
							ifres78962 = False

						}

						ifres78961 = ifres78962

					} else {
						ifres78961 = False

					}

					var ifres78960 Obj

					if True == ifres78961 {
						ifres78960 = True

					} else {
						ifres78960 = False

					}

					ifres78959 = ifres78960

				} else {
					ifres78959 = False

				}

				var ifres78958 Obj

				if True == ifres78959 {
					ifres78958 = True

				} else {
					ifres78958 = False

				}

				ifres78957 = ifres78958

			} else {
				ifres78957 = False

			}

			var ifres78956 Obj

			if True == ifres78957 {
				ifres78956 = True

			} else {
				ifres78956 = False

			}

			ifres78955 = ifres78956

		} else {
			ifres78955 = False

		}

		if True == ifres78955 {
			tmp78950 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78951 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp78953 := Call(__e, tmp78952, V828)

			tmp78954 := Call(__e, tmp78951, tmp78953)

			__e.TailApply(tmp78950, tmp78954)
			return

		} else {
			tmp78948 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp78949 := Call(__e, tmp78948, V828)

			var ifres78918 Obj

			if True == tmp78949 {
				tmp78944 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp78945 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78946 := Call(__e, tmp78945, V828)

				tmp78947 := Call(__e, tmp78944, symcases, tmp78946)

				var ifres78920 Obj

				if True == tmp78947 {
					tmp78940 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78942 := Call(__e, tmp78941, V828)

					tmp78943 := Call(__e, tmp78940, tmp78942)

					var ifres78922 Obj

					if True == tmp78943 {
						tmp78934 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78935 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78936 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78937 := Call(__e, tmp78936, V828)

						tmp78938 := Call(__e, tmp78935, tmp78937)

						tmp78939 := Call(__e, tmp78934, tmp78938)

						var ifres78924 Obj

						if True == tmp78939 {
							tmp78926 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp78927 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78929 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78930 := Call(__e, tmp78929, V828)

							tmp78931 := Call(__e, tmp78928, tmp78930)

							tmp78932 := Call(__e, tmp78927, tmp78931)

							tmp78933 := Call(__e, tmp78926, Nil, tmp78932)

							var ifres78925 Obj

							if True == tmp78933 {
								ifres78925 = True

							} else {
								ifres78925 = False

							}

							ifres78924 = ifres78925

						} else {
							ifres78924 = False

						}

						var ifres78923 Obj

						if True == ifres78924 {
							ifres78923 = True

						} else {
							ifres78923 = False

						}

						ifres78922 = ifres78923

					} else {
						ifres78922 = False

					}

					var ifres78921 Obj

					if True == ifres78922 {
						ifres78921 = True

					} else {
						ifres78921 = False

					}

					ifres78920 = ifres78921

				} else {
					ifres78920 = False

				}

				var ifres78919 Obj

				if True == ifres78920 {
					ifres78919 = True

				} else {
					ifres78919 = False

				}

				ifres78918 = ifres78919

			} else {
				ifres78918 = False

			}

			if True == ifres78918 {
				tmp78897 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78898 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78899 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78900 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78901 := Call(__e, tmp78900, V828)

				tmp78902 := Call(__e, tmp78899, tmp78901)

				tmp78903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78904 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp78905 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78906 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp78907 := Call(__e, tmp78906, V828)

				tmp78908 := Call(__e, tmp78905, tmp78907)

				tmp78909 := Call(__e, tmp78904, tmp78908)

				tmp78910 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78911 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78912 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp78913 := Call(__e, tmp78912, MakeString("error: cases exhausted"), Nil)

				tmp78914 := Call(__e, tmp78911, symsimple_1error, tmp78913)

				tmp78915 := Call(__e, tmp78910, tmp78914, Nil)

				tmp78916 := Call(__e, tmp78903, tmp78909, tmp78915)

				tmp78917 := Call(__e, tmp78898, tmp78902, tmp78916)

				__e.TailApply(tmp78897, symif, tmp78917)
				return

			} else {
				tmp78895 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp78896 := Call(__e, tmp78895, V828)

				var ifres78875 Obj

				if True == tmp78896 {
					tmp78891 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp78892 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78893 := Call(__e, tmp78892, V828)

					tmp78894 := Call(__e, tmp78891, symcases, tmp78893)

					var ifres78877 Obj

					if True == tmp78894 {
						tmp78887 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp78888 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp78889 := Call(__e, tmp78888, V828)

						tmp78890 := Call(__e, tmp78887, tmp78889)

						var ifres78879 Obj

						if True == tmp78890 {
							tmp78881 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78882 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78884 := Call(__e, tmp78883, V828)

							tmp78885 := Call(__e, tmp78882, tmp78884)

							tmp78886 := Call(__e, tmp78881, tmp78885)

							var ifres78880 Obj

							if True == tmp78886 {
								ifres78880 = True

							} else {
								ifres78880 = False

							}

							ifres78879 = ifres78880

						} else {
							ifres78879 = False

						}

						var ifres78878 Obj

						if True == ifres78879 {
							ifres78878 = True

						} else {
							ifres78878 = False

						}

						ifres78877 = ifres78878

					} else {
						ifres78877 = False

					}

					var ifres78876 Obj

					if True == ifres78877 {
						ifres78876 = True

					} else {
						ifres78876 = False

					}

					ifres78875 = ifres78876

				} else {
					ifres78875 = False

				}

				if True == ifres78875 {
					tmp78848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78850 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78851 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78852 := Call(__e, tmp78851, V828)

					tmp78853 := Call(__e, tmp78850, tmp78852)

					tmp78854 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78855 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp78856 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78857 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78858 := Call(__e, tmp78857, V828)

					tmp78859 := Call(__e, tmp78856, tmp78858)

					tmp78860 := Call(__e, tmp78855, tmp78859)

					tmp78861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78862 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cases_1macro)

					tmp78863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp78864 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78865 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78866 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp78867 := Call(__e, tmp78866, V828)

					tmp78868 := Call(__e, tmp78865, tmp78867)

					tmp78869 := Call(__e, tmp78864, tmp78868)

					tmp78870 := Call(__e, tmp78863, symcases, tmp78869)

					tmp78871 := Call(__e, tmp78862, tmp78870)

					tmp78872 := Call(__e, tmp78861, tmp78871, Nil)

					tmp78873 := Call(__e, tmp78854, tmp78860, tmp78872)

					tmp78874 := Call(__e, tmp78849, tmp78853, tmp78873)

					__e.TailApply(tmp78848, symif, tmp78874)
					return

				} else {
					tmp78846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp78847 := Call(__e, tmp78846, V828)

					var ifres78826 Obj

					if True == tmp78847 {
						tmp78842 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp78843 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp78844 := Call(__e, tmp78843, V828)

						tmp78845 := Call(__e, tmp78842, symcases, tmp78844)

						var ifres78828 Obj

						if True == tmp78845 {
							tmp78838 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp78839 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp78840 := Call(__e, tmp78839, V828)

							tmp78841 := Call(__e, tmp78838, tmp78840)

							var ifres78830 Obj

							if True == tmp78841 {
								tmp78832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp78833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp78835 := Call(__e, tmp78834, V828)

								tmp78836 := Call(__e, tmp78833, tmp78835)

								tmp78837 := Call(__e, tmp78832, Nil, tmp78836)

								var ifres78831 Obj

								if True == tmp78837 {
									ifres78831 = True

								} else {
									ifres78831 = False

								}

								ifres78830 = ifres78831

							} else {
								ifres78830 = False

							}

							var ifres78829 Obj

							if True == ifres78830 {
								ifres78829 = True

							} else {
								ifres78829 = False

							}

							ifres78828 = ifres78829

						} else {
							ifres78828 = False

						}

						var ifres78827 Obj

						if True == ifres78828 {
							ifres78827 = True

						} else {
							ifres78827 = False

						}

						ifres78826 = ifres78827

					} else {
						ifres78826 = False

					}

					if True == ifres78826 {
						tmp78825 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(tmp78825, MakeString("error: odd number of case elements\n"))
						return

					} else {
						__e.Return(V828)
						return
					}

				}

			}

		}

	}, 1)

	tmp78985 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cases_1macro, tmp78820)

	_ = tmp78985

	tmp78986 := MakeNative(func(__e *ControlFlow) {
		V830 := __e.Get(1)
		_ = V830
		tmp79075 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79076 := Call(__e, tmp79075, V830)

		var ifres79055 Obj

		if True == tmp79076 {
			tmp79071 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79072 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79073 := Call(__e, tmp79072, V830)

			tmp79074 := Call(__e, tmp79071, symtime, tmp79073)

			var ifres79057 Obj

			if True == tmp79074 {
				tmp79067 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79069 := Call(__e, tmp79068, V830)

				tmp79070 := Call(__e, tmp79067, tmp79069)

				var ifres79059 Obj

				if True == tmp79070 {
					tmp79061 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79062 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79063 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79064 := Call(__e, tmp79063, V830)

					tmp79065 := Call(__e, tmp79062, tmp79064)

					tmp79066 := Call(__e, tmp79061, Nil, tmp79065)

					var ifres79060 Obj

					if True == tmp79066 {
						ifres79060 = True

					} else {
						ifres79060 = False

					}

					ifres79059 = ifres79060

				} else {
					ifres79059 = False

				}

				var ifres79058 Obj

				if True == ifres79059 {
					ifres79058 = True

				} else {
					ifres79058 = False

				}

				ifres79057 = ifres79058

			} else {
				ifres79057 = False

			}

			var ifres79056 Obj

			if True == ifres79057 {
				ifres79056 = True

			} else {
				ifres79056 = False

			}

			ifres79055 = ifres79056

		} else {
			ifres79055 = False

		}

		if True == ifres79055 {
			tmp78988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4let_1macro)

			tmp78989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78991 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78992 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78993 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78994 := Call(__e, tmp78993, symrun, Nil)

			tmp78995 := Call(__e, tmp78992, symget_1time, tmp78994)

			tmp78996 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78997 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp78998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp78999 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79000 := Call(__e, tmp78999, V830)

			tmp79001 := Call(__e, tmp78998, tmp79000)

			tmp79002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79003 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79004 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79006 := Call(__e, tmp79005, symrun, Nil)

			tmp79007 := Call(__e, tmp79004, symget_1time, tmp79006)

			tmp79008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79010 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79013 := Call(__e, tmp79012, symStart, Nil)

			tmp79014 := Call(__e, tmp79011, symFinish, tmp79013)

			tmp79015 := Call(__e, tmp79010, sym_1, tmp79014)

			tmp79016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79021 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79022 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79023 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79024 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79027 := Call(__e, tmp79026, symTime, Nil)

			tmp79028 := Call(__e, tmp79025, symstr, tmp79027)

			tmp79029 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79030 := Call(__e, tmp79029, MakeString(" secs\n"), Nil)

			tmp79031 := Call(__e, tmp79024, tmp79028, tmp79030)

			tmp79032 := Call(__e, tmp79023, symcn, tmp79031)

			tmp79033 := Call(__e, tmp79022, tmp79032, Nil)

			tmp79034 := Call(__e, tmp79021, MakeString("\nrun time: "), tmp79033)

			tmp79035 := Call(__e, tmp79020, symcn, tmp79034)

			tmp79036 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79037 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79038 := Call(__e, tmp79037, symstoutput, Nil)

			tmp79039 := Call(__e, tmp79036, tmp79038, Nil)

			tmp79040 := Call(__e, tmp79019, tmp79035, tmp79039)

			tmp79041 := Call(__e, tmp79018, symshen_4prhush, tmp79040)

			tmp79042 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79043 := Call(__e, tmp79042, symResult, Nil)

			tmp79044 := Call(__e, tmp79017, tmp79041, tmp79043)

			tmp79045 := Call(__e, tmp79016, symMessage, tmp79044)

			tmp79046 := Call(__e, tmp79009, tmp79015, tmp79045)

			tmp79047 := Call(__e, tmp79008, symTime, tmp79046)

			tmp79048 := Call(__e, tmp79003, tmp79007, tmp79047)

			tmp79049 := Call(__e, tmp79002, symFinish, tmp79048)

			tmp79050 := Call(__e, tmp78997, tmp79001, tmp79049)

			tmp79051 := Call(__e, tmp78996, symResult, tmp79050)

			tmp79052 := Call(__e, tmp78991, tmp78995, tmp79051)

			tmp79053 := Call(__e, tmp78990, symStart, tmp79052)

			tmp79054 := Call(__e, tmp78989, symlet, tmp79053)

			__e.TailApply(tmp78988, tmp79054)
			return

		} else {
			__e.Return(V830)
			return
		}

	}, 1)

	tmp79077 := Call(__e, PrimNS1Value(symns2_1set), symshen_4timer_1macro, tmp78986)

	_ = tmp79077

	tmp79078 := MakeNative(func(__e *ControlFlow) {
		V832 := __e.Get(1)
		_ = V832
		tmp79091 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79092 := Call(__e, tmp79091, V832)

		if True == tmp79092 {
			tmp79080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79081 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79082 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79083 := Call(__e, tmp79082, V832)

			tmp79084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tuple_1up)

			tmp79086 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79087 := Call(__e, tmp79086, V832)

			tmp79088 := Call(__e, tmp79085, tmp79087)

			tmp79089 := Call(__e, tmp79084, tmp79088, Nil)

			tmp79090 := Call(__e, tmp79081, tmp79083, tmp79089)

			__e.TailApply(tmp79080, sym_8p, tmp79090)
			return

		} else {
			__e.Return(V832)
			return
		}

	}, 1)

	tmp79093 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple_1up, tmp79078)

	_ = tmp79093

	tmp79094 := MakeNative(func(__e *ControlFlow) {
		V834 := __e.Get(1)
		_ = V834
		tmp79277 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79278 := Call(__e, tmp79277, V834)

		var ifres79235 Obj

		if True == tmp79278 {
			tmp79273 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79274 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79275 := Call(__e, tmp79274, V834)

			tmp79276 := Call(__e, tmp79273, symput, tmp79275)

			var ifres79237 Obj

			if True == tmp79276 {
				tmp79269 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79270 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79271 := Call(__e, tmp79270, V834)

				tmp79272 := Call(__e, tmp79269, tmp79271)

				var ifres79239 Obj

				if True == tmp79272 {
					tmp79263 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79264 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79265 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79266 := Call(__e, tmp79265, V834)

					tmp79267 := Call(__e, tmp79264, tmp79266)

					tmp79268 := Call(__e, tmp79263, tmp79267)

					var ifres79241 Obj

					if True == tmp79268 {
						tmp79255 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79256 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79257 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79259 := Call(__e, tmp79258, V834)

						tmp79260 := Call(__e, tmp79257, tmp79259)

						tmp79261 := Call(__e, tmp79256, tmp79260)

						tmp79262 := Call(__e, tmp79255, tmp79261)

						var ifres79243 Obj

						if True == tmp79262 {
							tmp79245 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp79246 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79247 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79250 := Call(__e, tmp79249, V834)

							tmp79251 := Call(__e, tmp79248, tmp79250)

							tmp79252 := Call(__e, tmp79247, tmp79251)

							tmp79253 := Call(__e, tmp79246, tmp79252)

							tmp79254 := Call(__e, tmp79245, Nil, tmp79253)

							var ifres79244 Obj

							if True == tmp79254 {
								ifres79244 = True

							} else {
								ifres79244 = False

							}

							ifres79243 = ifres79244

						} else {
							ifres79243 = False

						}

						var ifres79242 Obj

						if True == ifres79243 {
							ifres79242 = True

						} else {
							ifres79242 = False

						}

						ifres79241 = ifres79242

					} else {
						ifres79241 = False

					}

					var ifres79240 Obj

					if True == ifres79241 {
						ifres79240 = True

					} else {
						ifres79240 = False

					}

					ifres79239 = ifres79240

				} else {
					ifres79239 = False

				}

				var ifres79238 Obj

				if True == ifres79239 {
					ifres79238 = True

				} else {
					ifres79238 = False

				}

				ifres79237 = ifres79238

			} else {
				ifres79237 = False

			}

			var ifres79236 Obj

			if True == ifres79237 {
				ifres79236 = True

			} else {
				ifres79236 = False

			}

			ifres79235 = ifres79236

		} else {
			ifres79235 = False

		}

		if True == ifres79235 {
			tmp79204 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79205 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79206 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79208 := Call(__e, tmp79207, V834)

			tmp79209 := Call(__e, tmp79206, tmp79208)

			tmp79210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79211 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79212 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79213 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79214 := Call(__e, tmp79213, V834)

			tmp79215 := Call(__e, tmp79212, tmp79214)

			tmp79216 := Call(__e, tmp79211, tmp79215)

			tmp79217 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79218 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79219 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79220 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79222 := Call(__e, tmp79221, V834)

			tmp79223 := Call(__e, tmp79220, tmp79222)

			tmp79224 := Call(__e, tmp79219, tmp79223)

			tmp79225 := Call(__e, tmp79218, tmp79224)

			tmp79226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79227 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp79229 := Call(__e, tmp79228, sym_dproperty_1vector_d, Nil)

			tmp79230 := Call(__e, tmp79227, symvalue, tmp79229)

			tmp79231 := Call(__e, tmp79226, tmp79230, Nil)

			tmp79232 := Call(__e, tmp79217, tmp79225, tmp79231)

			tmp79233 := Call(__e, tmp79210, tmp79216, tmp79232)

			tmp79234 := Call(__e, tmp79205, tmp79209, tmp79233)

			__e.TailApply(tmp79204, symput, tmp79234)
			return

		} else {
			tmp79202 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79203 := Call(__e, tmp79202, V834)

			var ifres79172 Obj

			if True == tmp79203 {
				tmp79198 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79199 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79200 := Call(__e, tmp79199, V834)

				tmp79201 := Call(__e, tmp79198, symget, tmp79200)

				var ifres79174 Obj

				if True == tmp79201 {
					tmp79194 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp79195 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79196 := Call(__e, tmp79195, V834)

					tmp79197 := Call(__e, tmp79194, tmp79196)

					var ifres79176 Obj

					if True == tmp79197 {
						tmp79188 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79189 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79190 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79191 := Call(__e, tmp79190, V834)

						tmp79192 := Call(__e, tmp79189, tmp79191)

						tmp79193 := Call(__e, tmp79188, tmp79192)

						var ifres79178 Obj

						if True == tmp79193 {
							tmp79180 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp79181 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79182 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79183 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79184 := Call(__e, tmp79183, V834)

							tmp79185 := Call(__e, tmp79182, tmp79184)

							tmp79186 := Call(__e, tmp79181, tmp79185)

							tmp79187 := Call(__e, tmp79180, Nil, tmp79186)

							var ifres79179 Obj

							if True == tmp79187 {
								ifres79179 = True

							} else {
								ifres79179 = False

							}

							ifres79178 = ifres79179

						} else {
							ifres79178 = False

						}

						var ifres79177 Obj

						if True == ifres79178 {
							ifres79177 = True

						} else {
							ifres79177 = False

						}

						ifres79176 = ifres79177

					} else {
						ifres79176 = False

					}

					var ifres79175 Obj

					if True == ifres79176 {
						ifres79175 = True

					} else {
						ifres79175 = False

					}

					ifres79174 = ifres79175

				} else {
					ifres79174 = False

				}

				var ifres79173 Obj

				if True == ifres79174 {
					ifres79173 = True

				} else {
					ifres79173 = False

				}

				ifres79172 = ifres79173

			} else {
				ifres79172 = False

			}

			if True == ifres79172 {
				tmp79151 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79152 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79153 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79154 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79155 := Call(__e, tmp79154, V834)

				tmp79156 := Call(__e, tmp79153, tmp79155)

				tmp79157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79158 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79159 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79160 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79161 := Call(__e, tmp79160, V834)

				tmp79162 := Call(__e, tmp79159, tmp79161)

				tmp79163 := Call(__e, tmp79158, tmp79162)

				tmp79164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79167 := Call(__e, tmp79166, sym_dproperty_1vector_d, Nil)

				tmp79168 := Call(__e, tmp79165, symvalue, tmp79167)

				tmp79169 := Call(__e, tmp79164, tmp79168, Nil)

				tmp79170 := Call(__e, tmp79157, tmp79163, tmp79169)

				tmp79171 := Call(__e, tmp79152, tmp79156, tmp79170)

				__e.TailApply(tmp79151, symget, tmp79171)
				return

			} else {
				tmp79149 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79150 := Call(__e, tmp79149, V834)

				var ifres79119 Obj

				if True == tmp79150 {
					tmp79145 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79147 := Call(__e, tmp79146, V834)

					tmp79148 := Call(__e, tmp79145, symunput, tmp79147)

					var ifres79121 Obj

					if True == tmp79148 {
						tmp79141 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp79142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp79143 := Call(__e, tmp79142, V834)

						tmp79144 := Call(__e, tmp79141, tmp79143)

						var ifres79123 Obj

						if True == tmp79144 {
							tmp79135 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp79136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp79138 := Call(__e, tmp79137, V834)

							tmp79139 := Call(__e, tmp79136, tmp79138)

							tmp79140 := Call(__e, tmp79135, tmp79139)

							var ifres79125 Obj

							if True == tmp79140 {
								tmp79127 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp79128 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79129 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp79131 := Call(__e, tmp79130, V834)

								tmp79132 := Call(__e, tmp79129, tmp79131)

								tmp79133 := Call(__e, tmp79128, tmp79132)

								tmp79134 := Call(__e, tmp79127, Nil, tmp79133)

								var ifres79126 Obj

								if True == tmp79134 {
									ifres79126 = True

								} else {
									ifres79126 = False

								}

								ifres79125 = ifres79126

							} else {
								ifres79125 = False

							}

							var ifres79124 Obj

							if True == ifres79125 {
								ifres79124 = True

							} else {
								ifres79124 = False

							}

							ifres79123 = ifres79124

						} else {
							ifres79123 = False

						}

						var ifres79122 Obj

						if True == ifres79123 {
							ifres79122 = True

						} else {
							ifres79122 = False

						}

						ifres79121 = ifres79122

					} else {
						ifres79121 = False

					}

					var ifres79120 Obj

					if True == ifres79121 {
						ifres79120 = True

					} else {
						ifres79120 = False

					}

					ifres79119 = ifres79120

				} else {
					ifres79119 = False

				}

				if True == ifres79119 {
					tmp79098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79100 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79101 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79102 := Call(__e, tmp79101, V834)

					tmp79103 := Call(__e, tmp79100, tmp79102)

					tmp79104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79105 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp79106 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79107 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79108 := Call(__e, tmp79107, V834)

					tmp79109 := Call(__e, tmp79106, tmp79108)

					tmp79110 := Call(__e, tmp79105, tmp79109)

					tmp79111 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp79114 := Call(__e, tmp79113, sym_dproperty_1vector_d, Nil)

					tmp79115 := Call(__e, tmp79112, symvalue, tmp79114)

					tmp79116 := Call(__e, tmp79111, tmp79115, Nil)

					tmp79117 := Call(__e, tmp79104, tmp79110, tmp79116)

					tmp79118 := Call(__e, tmp79099, tmp79103, tmp79117)

					__e.TailApply(tmp79098, symunput, tmp79118)
					return

				} else {
					__e.Return(V834)
					return
				}

			}

		}

	}, 1)

	tmp79279 := Call(__e, PrimNS1Value(symns2_1set), symshen_4put_cget_1macro, tmp79094)

	_ = tmp79279

	tmp79280 := MakeNative(func(__e *ControlFlow) {
		V836 := __e.Get(1)
		_ = V836
		tmp79313 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp79314 := Call(__e, tmp79313, V836)

		var ifres79293 Obj

		if True == tmp79314 {
			tmp79309 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79310 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79311 := Call(__e, tmp79310, V836)

			tmp79312 := Call(__e, tmp79309, symfunction, tmp79311)

			var ifres79295 Obj

			if True == tmp79312 {
				tmp79305 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79306 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79307 := Call(__e, tmp79306, V836)

				tmp79308 := Call(__e, tmp79305, tmp79307)

				var ifres79297 Obj

				if True == tmp79308 {
					tmp79299 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp79300 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79301 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79302 := Call(__e, tmp79301, V836)

					tmp79303 := Call(__e, tmp79300, tmp79302)

					tmp79304 := Call(__e, tmp79299, Nil, tmp79303)

					var ifres79298 Obj

					if True == tmp79304 {
						ifres79298 = True

					} else {
						ifres79298 = False

					}

					ifres79297 = ifres79298

				} else {
					ifres79297 = False

				}

				var ifres79296 Obj

				if True == ifres79297 {
					ifres79296 = True

				} else {
					ifres79296 = False

				}

				ifres79295 = ifres79296

			} else {
				ifres79295 = False

			}

			var ifres79294 Obj

			if True == ifres79295 {
				ifres79294 = True

			} else {
				ifres79294 = False

			}

			ifres79293 = ifres79294

		} else {
			ifres79293 = False

		}

		if True == ifres79293 {
			tmp79282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction)

			tmp79283 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79284 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79285 := Call(__e, tmp79284, V836)

			tmp79286 := Call(__e, tmp79283, tmp79285)

			tmp79287 := Call(__e, PrimNS1Value(symns2_1value), symarity)

			tmp79288 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp79289 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp79290 := Call(__e, tmp79289, V836)

			tmp79291 := Call(__e, tmp79288, tmp79290)

			tmp79292 := Call(__e, tmp79287, tmp79291)

			__e.TailApply(tmp79282, tmp79286, tmp79292)
			return

		} else {
			__e.Return(V836)
			return
		}

	}, 1)

	tmp79315 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1macro, tmp79280)

	_ = tmp79315

	tmp79316 := MakeNative(func(__e *ControlFlow) {
		V839 := __e.Get(1)
		_ = V839
		V840 := __e.Get(2)
		_ = V840
		tmp79328 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp79329 := Call(__e, tmp79328, MakeNumber(0), V840)

		if True == tmp79329 {
			tmp79325 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp79326 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79327 := Call(__e, tmp79326, V839, MakeString(" has no lambda form\n"), symshen_4a)

			__e.TailApply(tmp79325, tmp79327)
			return

		} else {
			tmp79323 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp79324 := Call(__e, tmp79323, MakeNumber(-1), V840)

			if True == tmp79324 {
				tmp79320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79322 := Call(__e, tmp79321, V839, Nil)

				__e.TailApply(tmp79320, symfunction, tmp79322)
				return

			} else {
				tmp79319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction_1help)

				__e.TailApply(tmp79319, V839, V840, Nil)
				return

			}

		}

	}, 2)

	tmp79330 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction, tmp79316)

	_ = tmp79330

	tmp79331 := MakeNative(func(__e *ControlFlow) {
		V844 := __e.Get(1)
		_ = V844
		V845 := __e.Get(2)
		_ = V845
		V846 := __e.Get(3)
		_ = V846
		tmp79350 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp79351 := Call(__e, tmp79350, MakeNumber(0), V845)

		if True == tmp79351 {
			tmp79349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp79349, V844, V846)
			return

		} else {
			tmp79333 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp79334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction_1help)

				tmp79338 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp79339 := Call(__e, tmp79338, V845, MakeNumber(1))

				tmp79340 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				tmp79341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79342 := Call(__e, tmp79341, X, Nil)

				tmp79343 := Call(__e, tmp79340, V846, tmp79342)

				tmp79344 := Call(__e, tmp79337, V844, tmp79339, tmp79343)

				tmp79345 := Call(__e, tmp79336, tmp79344, Nil)

				tmp79346 := Call(__e, tmp79335, X, tmp79345)

				__e.TailApply(tmp79334, sym_c_4, tmp79346)
				return

			}, 1)

			tmp79347 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp79348 := Call(__e, tmp79347, symV)

			__e.TailApply(tmp79333, tmp79348)
			return

		}

	}, 3)

	tmp79352 := Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction_1help, tmp79331)

	_ = tmp79352

	tmp79353 := MakeNative(func(__e *ControlFlow) {
		V848 := __e.Get(1)
		_ = V848
		tmp79354 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			tmp79355 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				tmp79356 := MakeNative(func(__e *ControlFlow) {
					Remove1 := __e.Get(1)
					_ = Remove1
					tmp79357 := MakeNative(func(__e *ControlFlow) {
						Remove2 := __e.Get(1)
						_ = Remove2
						__e.Return(V848)
						return
					}, 1)

					tmp79358 := Call(__e, PrimNS1Value(symns2_1value), symset)

					tmp79359 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

					tmp79360 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					tmp79361 := Call(__e, tmp79360, sym_dmacros_d)

					tmp79362 := Call(__e, tmp79359, Pos, tmp79361)

					tmp79363 := Call(__e, tmp79358, sym_dmacros_d, tmp79362)

					__e.TailApply(tmp79357, tmp79363)
					return

				}, 1)

				tmp79364 := Call(__e, PrimNS1Value(symns2_1value), symset)

				tmp79365 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				tmp79366 := Call(__e, tmp79365, V848, MacroReg)

				tmp79367 := Call(__e, tmp79364, symshen_4_dmacroreg_d, tmp79366)

				__e.TailApply(tmp79356, tmp79367)
				return

			}, 1)

			tmp79368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

			tmp79369 := Call(__e, tmp79368, V848, MacroReg)

			__e.TailApply(tmp79355, tmp79369)
			return

		}, 1)

		tmp79370 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp79371 := Call(__e, tmp79370, symshen_4_dmacroreg_d)

		__e.TailApply(tmp79354, tmp79371)
		return

	}, 1)

	tmp79372 := Call(__e, PrimNS1Value(symns2_1set), symundefmacro, tmp79353)

	_ = tmp79372

	tmp79373 := MakeNative(func(__e *ControlFlow) {
		V858 := __e.Get(1)
		_ = V858
		V859 := __e.Get(2)
		_ = V859
		tmp79396 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp79397 := Call(__e, tmp79396, Nil, V859)

		if True == tmp79397 {
			tmp79393 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp79394 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp79395 := Call(__e, tmp79394, V858, MakeString(" is not a macro\n"), symshen_4a)

			__e.TailApply(tmp79393, tmp79395)
			return

		} else {
			tmp79391 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79392 := Call(__e, tmp79391, V859)

			var ifres79385 Obj

			if True == tmp79392 {
				tmp79387 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp79388 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79389 := Call(__e, tmp79388, V859)

				tmp79390 := Call(__e, tmp79387, tmp79389, V858)

				var ifres79386 Obj

				if True == tmp79390 {
					ifres79386 = True

				} else {
					ifres79386 = False

				}

				ifres79385 = ifres79386

			} else {
				ifres79385 = False

			}

			if True == ifres79385 {
				__e.Return(MakeNumber(1))
				return
			} else {
				tmp79383 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp79384 := Call(__e, tmp79383, V859)

				if True == tmp79384 {
					tmp79378 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					tmp79379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

					tmp79380 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp79381 := Call(__e, tmp79380, V859)

					tmp79382 := Call(__e, tmp79379, V858, tmp79381)

					__e.TailApply(tmp79378, MakeNumber(1), tmp79382)
					return

				} else {
					tmp79377 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp79377, symshen_4findpos)
					return

				}

			}

		}

	}, 2)

	tmp79398 := Call(__e, PrimNS1Value(symns2_1set), symshen_4findpos, tmp79373)

	_ = tmp79398

	tmp79399 := MakeNative(func(__e *ControlFlow) {
		V864 := __e.Get(1)
		_ = V864
		V865 := __e.Get(2)
		_ = V865
		tmp79419 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp79420 := Call(__e, tmp79419, MakeNumber(1), V864)

		var ifres79415 Obj

		if True == tmp79420 {
			tmp79417 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79418 := Call(__e, tmp79417, V865)

			var ifres79416 Obj

			if True == tmp79418 {
				ifres79416 = True

			} else {
				ifres79416 = False

			}

			ifres79415 = ifres79416

		} else {
			ifres79415 = False

		}

		if True == ifres79415 {
			tmp79414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(tmp79414, V865)
			return

		} else {
			tmp79412 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp79413 := Call(__e, tmp79412, V865)

			if True == tmp79413 {
				tmp79403 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp79404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp79405 := Call(__e, tmp79404, V865)

				tmp79406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

				tmp79407 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				tmp79408 := Call(__e, tmp79407, V864, MakeNumber(1))

				tmp79409 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp79410 := Call(__e, tmp79409, V865)

				tmp79411 := Call(__e, tmp79406, tmp79408, tmp79410)

				__e.TailApply(tmp79403, tmp79405, tmp79411)
				return

			} else {
				tmp79402 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp79402, symshen_4remove_1nth)
				return

			}

		}

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4remove_1nth, tmp79399)
	return

}, 0)
