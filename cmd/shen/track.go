package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp17628 := MakeNative(func(__e *ControlFlow) {
		V3135 := __e.Get(1)
		_ = V3135
		tmp17629 := Call(__e, PrimNS2Value(symshen_4app), V3135, MakeString(";\n"), symshen_4a)

		tmp17630 := PrimStringConcat(MakeString("partial function "), tmp17629)

		tmp17631 := Call(__e, PrimNS2Value(symstoutput))

		tmp17632 := Call(__e, PrimNS2Value(symshen_4prhush), tmp17630, tmp17631)

		_ = tmp17632

		tmp17641 := Call(__e, PrimNS2Value(symshen_4tracked_2), V3135)

		tmp17642 := PrimNot(tmp17641)

		var ifres17636 Obj

		if True == tmp17642 {
			tmp17638 := Call(__e, PrimNS2Value(symshen_4app), V3135, MakeString("? "), symshen_4a)

			tmp17639 := PrimStringConcat(MakeString("track "), tmp17638)

			tmp17640 := Call(__e, PrimNS2Value(symy_1or_1n_2), tmp17639)

			var ifres17637 Obj

			if True == tmp17640 {
				ifres17637 = True

			} else {
				ifres17637 = False

			}

			ifres17636 = ifres17637

		} else {
			ifres17636 = False

		}

		var ifres17633 Obj

		if True == ifres17636 {
			tmp17634 := Call(__e, PrimNS2Value(symps), V3135)

			tmp17635 := Call(__e, PrimNS2Value(symshen_4track_1function), tmp17634)

			ifres17633 = tmp17635

		} else {
			ifres17633 = symshen_4ok

		}

		_ = ifres17633

		__e.Return(PrimSimpleError(MakeString("aborted")))
		return

	}, 1)

	tmp17643 := Call(__e, PrimNS1Value(symns2_1set), symshen_4f__error, tmp17628)

	_ = tmp17643

	tmp17644 := MakeNative(func(__e *ControlFlow) {
		V3137 := __e.Get(1)
		_ = V3137
		tmp17645 := PrimNS3Value(symshen_4_dtracking_d)

		__e.TailApply(PrimNS2Value(symelement_2), V3137, tmp17645)
		return

	}, 1)

	tmp17646 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tracked_2, tmp17644)

	_ = tmp17646

	tmp17647 := MakeNative(func(__e *ControlFlow) {
		V3139 := __e.Get(1)
		_ = V3139
		tmp17648 := MakeNative(func(__e *ControlFlow) {
			Source := __e.Get(1)
			_ = Source
			__e.TailApply(PrimNS2Value(symshen_4track_1function), Source)
			return
		}, 1)

		tmp17649 := Call(__e, PrimNS2Value(symps), V3139)

		__e.TailApply(tmp17648, tmp17649)
		return

	}, 1)

	tmp17650 := Call(__e, PrimNS1Value(symns2_1set), symtrack, tmp17647)

	_ = tmp17650

	tmp17651 := MakeNative(func(__e *ControlFlow) {
		V3141 := __e.Get(1)
		_ = V3141
		tmp17705 := PrimIsPair(V3141)

		var ifres17679 Obj

		if True == tmp17705 {
			tmp17703 := PrimHead(V3141)

			tmp17704 := PrimEqual(symdefun, tmp17703)

			var ifres17681 Obj

			if True == tmp17704 {
				tmp17701 := PrimTail(V3141)

				tmp17702 := PrimIsPair(tmp17701)

				var ifres17683 Obj

				if True == tmp17702 {
					tmp17698 := PrimTail(V3141)

					tmp17699 := PrimTail(tmp17698)

					tmp17700 := PrimIsPair(tmp17699)

					var ifres17685 Obj

					if True == tmp17700 {
						tmp17694 := PrimTail(V3141)

						tmp17695 := PrimTail(tmp17694)

						tmp17696 := PrimTail(tmp17695)

						tmp17697 := PrimIsPair(tmp17696)

						var ifres17687 Obj

						if True == tmp17697 {
							tmp17689 := PrimTail(V3141)

							tmp17690 := PrimTail(tmp17689)

							tmp17691 := PrimTail(tmp17690)

							tmp17692 := PrimTail(tmp17691)

							tmp17693 := PrimEqual(Nil, tmp17692)

							var ifres17688 Obj

							if True == tmp17693 {
								ifres17688 = True

							} else {
								ifres17688 = False

							}

							ifres17687 = ifres17688

						} else {
							ifres17687 = False

						}

						var ifres17686 Obj

						if True == ifres17687 {
							ifres17686 = True

						} else {
							ifres17686 = False

						}

						ifres17685 = ifres17686

					} else {
						ifres17685 = False

					}

					var ifres17684 Obj

					if True == ifres17685 {
						ifres17684 = True

					} else {
						ifres17684 = False

					}

					ifres17683 = ifres17684

				} else {
					ifres17683 = False

				}

				var ifres17682 Obj

				if True == ifres17683 {
					ifres17682 = True

				} else {
					ifres17682 = False

				}

				ifres17681 = ifres17682

			} else {
				ifres17681 = False

			}

			var ifres17680 Obj

			if True == ifres17681 {
				ifres17680 = True

			} else {
				ifres17680 = False

			}

			ifres17679 = ifres17680

		} else {
			ifres17679 = False

		}

		if True == ifres17679 {
			tmp17653 := MakeNative(func(__e *ControlFlow) {
				KL := __e.Get(1)
				_ = KL
				tmp17654 := MakeNative(func(__e *ControlFlow) {
					Ob := __e.Get(1)
					_ = Ob
					tmp17655 := MakeNative(func(__e *ControlFlow) {
						Tr := __e.Get(1)
						_ = Tr
						__e.Return(Ob)
						return
					}, 1)

					tmp17656 := PrimNS3Value(symshen_4_dtracking_d)

					tmp17657 := PrimCons(Ob, tmp17656)

					tmp17658 := PrimNS3Set(symshen_4_dtracking_d, tmp17657)

					__e.TailApply(tmp17655, tmp17658)
					return

				}, 1)

				tmp17659 := Call(__e, PrimNS2Value(symeval_1kl), KL)

				__e.TailApply(tmp17654, tmp17659)
				return

			}, 1)

			tmp17660 := PrimTail(V3141)

			tmp17661 := PrimHead(tmp17660)

			tmp17662 := PrimTail(V3141)

			tmp17663 := PrimTail(tmp17662)

			tmp17664 := PrimHead(tmp17663)

			tmp17665 := PrimTail(V3141)

			tmp17666 := PrimHead(tmp17665)

			tmp17667 := PrimTail(V3141)

			tmp17668 := PrimTail(tmp17667)

			tmp17669 := PrimHead(tmp17668)

			tmp17670 := PrimTail(V3141)

			tmp17671 := PrimTail(tmp17670)

			tmp17672 := PrimTail(tmp17671)

			tmp17673 := PrimHead(tmp17672)

			tmp17674 := Call(__e, PrimNS2Value(symshen_4insert_1tracking_1code), tmp17666, tmp17669, tmp17673)

			tmp17675 := PrimCons(tmp17674, Nil)

			tmp17676 := PrimCons(tmp17664, tmp17675)

			tmp17677 := PrimCons(tmp17661, tmp17676)

			tmp17678 := PrimCons(symdefun, tmp17677)

			__e.TailApply(tmp17653, tmp17678)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4track_1function)
			return
		}

	}, 1)

	tmp17706 := Call(__e, PrimNS1Value(symns2_1set), symshen_4track_1function, tmp17651)

	_ = tmp17706

	tmp17707 := MakeNative(func(__e *ControlFlow) {
		V3145 := __e.Get(1)
		_ = V3145
		V3146 := __e.Get(2)
		_ = V3146
		V3147 := __e.Get(3)
		_ = V3147
		tmp17708 := PrimCons(symshen_4_dcall_d, Nil)

		tmp17709 := PrimCons(symvalue, tmp17708)

		tmp17710 := PrimCons(MakeNumber(1), Nil)

		tmp17711 := PrimCons(tmp17709, tmp17710)

		tmp17712 := PrimCons(sym_7, tmp17711)

		tmp17713 := PrimCons(tmp17712, Nil)

		tmp17714 := PrimCons(symshen_4_dcall_d, tmp17713)

		tmp17715 := PrimCons(symset, tmp17714)

		tmp17716 := PrimCons(symshen_4_dcall_d, Nil)

		tmp17717 := PrimCons(symvalue, tmp17716)

		tmp17718 := Call(__e, PrimNS2Value(symshen_4cons__form), V3146)

		tmp17719 := PrimCons(tmp17718, Nil)

		tmp17720 := PrimCons(V3145, tmp17719)

		tmp17721 := PrimCons(tmp17717, tmp17720)

		tmp17722 := PrimCons(symshen_4input_1track, tmp17721)

		tmp17723 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp17724 := PrimCons(symshen_4_dcall_d, Nil)

		tmp17725 := PrimCons(symvalue, tmp17724)

		tmp17726 := PrimCons(symResult, Nil)

		tmp17727 := PrimCons(V3145, tmp17726)

		tmp17728 := PrimCons(tmp17725, tmp17727)

		tmp17729 := PrimCons(symshen_4output_1track, tmp17728)

		tmp17730 := PrimCons(symshen_4_dcall_d, Nil)

		tmp17731 := PrimCons(symvalue, tmp17730)

		tmp17732 := PrimCons(MakeNumber(1), Nil)

		tmp17733 := PrimCons(tmp17731, tmp17732)

		tmp17734 := PrimCons(sym_1, tmp17733)

		tmp17735 := PrimCons(tmp17734, Nil)

		tmp17736 := PrimCons(symshen_4_dcall_d, tmp17735)

		tmp17737 := PrimCons(symset, tmp17736)

		tmp17738 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp17739 := PrimCons(symResult, Nil)

		tmp17740 := PrimCons(tmp17738, tmp17739)

		tmp17741 := PrimCons(symdo, tmp17740)

		tmp17742 := PrimCons(tmp17741, Nil)

		tmp17743 := PrimCons(tmp17737, tmp17742)

		tmp17744 := PrimCons(symdo, tmp17743)

		tmp17745 := PrimCons(tmp17744, Nil)

		tmp17746 := PrimCons(tmp17729, tmp17745)

		tmp17747 := PrimCons(symdo, tmp17746)

		tmp17748 := PrimCons(tmp17747, Nil)

		tmp17749 := PrimCons(V3147, tmp17748)

		tmp17750 := PrimCons(symResult, tmp17749)

		tmp17751 := PrimCons(symlet, tmp17750)

		tmp17752 := PrimCons(tmp17751, Nil)

		tmp17753 := PrimCons(tmp17723, tmp17752)

		tmp17754 := PrimCons(symdo, tmp17753)

		tmp17755 := PrimCons(tmp17754, Nil)

		tmp17756 := PrimCons(tmp17722, tmp17755)

		tmp17757 := PrimCons(symdo, tmp17756)

		tmp17758 := PrimCons(tmp17757, Nil)

		tmp17759 := PrimCons(tmp17715, tmp17758)

		__e.Return(PrimCons(symdo, tmp17759))
		return

	}, 3)

	tmp17760 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1tracking_1code, tmp17707)

	_ = tmp17760

	tmp17761 := MakeNative(func(__e *ControlFlow) {
		V3153 := __e.Get(1)
		_ = V3153
		tmp17765 := PrimEqual(sym_7, V3153)

		if True == tmp17765 {
			__e.Return(PrimNS3Set(symshen_4_dstep_d, True))
			return
		} else {
			tmp17764 := PrimEqual(sym_1, V3153)

			if True == tmp17764 {
				__e.Return(PrimNS3Set(symshen_4_dstep_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp17766 := Call(__e, PrimNS1Value(symns2_1set), symstep, tmp17761)

	_ = tmp17766

	tmp17767 := MakeNative(func(__e *ControlFlow) {
		V3159 := __e.Get(1)
		_ = V3159
		tmp17771 := PrimEqual(sym_7, V3159)

		if True == tmp17771 {
			__e.Return(PrimNS3Set(symshen_4_dspy_d, True))
			return
		} else {
			tmp17770 := PrimEqual(sym_1, V3159)

			if True == tmp17770 {
				__e.Return(PrimNS3Set(symshen_4_dspy_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp17772 := Call(__e, PrimNS1Value(symns2_1set), symspy, tmp17767)

	_ = tmp17772

	tmp17773 := MakeNative(func(__e *ControlFlow) {
		tmp17777 := PrimNS3Value(symshen_4_dstep_d)

		if True == tmp17777 {
			tmp17775 := PrimNS3Value(sym_dstinput_d)

			tmp17776 := PrimReadByte(tmp17775)

			__e.TailApply(PrimNS2Value(symshen_4check_1byte), tmp17776)
			return

		} else {
			__e.TailApply(PrimNS2Value(symnl), MakeNumber(1))
			return
		}

	}, 0)

	tmp17778 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terpri_1or_1read_1char, tmp17773)

	_ = tmp17778

	tmp17779 := MakeNative(func(__e *ControlFlow) {
		V3165 := __e.Get(1)
		_ = V3165
		tmp17781 := Call(__e, PrimNS2Value(symshen_4hat))

		tmp17782 := PrimEqual(V3165, tmp17781)

		if True == tmp17782 {
			__e.Return(PrimSimpleError(MakeString("aborted")))
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp17783 := Call(__e, PrimNS1Value(symns2_1set), symshen_4check_1byte, tmp17779)

	_ = tmp17783

	tmp17784 := MakeNative(func(__e *ControlFlow) {
		V3169 := __e.Get(1)
		_ = V3169
		V3170 := __e.Get(2)
		_ = V3170
		V3171 := __e.Get(3)
		_ = V3171
		tmp17785 := Call(__e, PrimNS2Value(symshen_4spaces), V3169)

		tmp17786 := Call(__e, PrimNS2Value(symshen_4spaces), V3169)

		tmp17787 := Call(__e, PrimNS2Value(symshen_4app), tmp17786, MakeString(""), symshen_4a)

		tmp17788 := PrimStringConcat(MakeString(" \n"), tmp17787)

		tmp17789 := Call(__e, PrimNS2Value(symshen_4app), V3170, tmp17788, symshen_4a)

		tmp17790 := PrimStringConcat(MakeString("> Inputs to "), tmp17789)

		tmp17791 := Call(__e, PrimNS2Value(symshen_4app), V3169, tmp17790, symshen_4a)

		tmp17792 := PrimStringConcat(MakeString("<"), tmp17791)

		tmp17793 := Call(__e, PrimNS2Value(symshen_4app), tmp17785, tmp17792, symshen_4a)

		tmp17794 := PrimStringConcat(MakeString("\n"), tmp17793)

		tmp17795 := Call(__e, PrimNS2Value(symstoutput))

		tmp17796 := Call(__e, PrimNS2Value(symshen_4prhush), tmp17794, tmp17795)

		_ = tmp17796

		__e.TailApply(PrimNS2Value(symshen_4recursively_1print), V3171)
		return

	}, 3)

	tmp17797 := Call(__e, PrimNS1Value(symns2_1set), symshen_4input_1track, tmp17784)

	_ = tmp17797

	tmp17798 := MakeNative(func(__e *ControlFlow) {
		V3173 := __e.Get(1)
		_ = V3173
		tmp17808 := PrimEqual(Nil, V3173)

		if True == tmp17808 {
			tmp17807 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), MakeString(" ==>"), tmp17807)
			return

		} else {
			tmp17806 := PrimIsPair(V3173)

			if True == tmp17806 {
				tmp17801 := PrimHead(V3173)

				tmp17802 := Call(__e, PrimNS2Value(symprint), tmp17801)

				_ = tmp17802

				tmp17803 := Call(__e, PrimNS2Value(symstoutput))

				tmp17804 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString(", "), tmp17803)

				_ = tmp17804

				tmp17805 := PrimTail(V3173)

				__e.TailApply(PrimNS2Value(symshen_4recursively_1print), tmp17805)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4recursively_1print)
				return
			}

		}

	}, 1)

	tmp17809 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursively_1print, tmp17798)

	_ = tmp17809

	tmp17810 := MakeNative(func(__e *ControlFlow) {
		V3175 := __e.Get(1)
		_ = V3175
		tmp17814 := PrimEqual(MakeNumber(0), V3175)

		if True == tmp17814 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp17812 := PrimNumberSubtract(V3175, MakeNumber(1))

			tmp17813 := Call(__e, PrimNS2Value(symshen_4spaces), tmp17812)

			__e.Return(PrimStringConcat(MakeString(" "), tmp17813))
			return

		}

	}, 1)

	tmp17815 := Call(__e, PrimNS1Value(symns2_1set), symshen_4spaces, tmp17810)

	_ = tmp17815

	tmp17816 := MakeNative(func(__e *ControlFlow) {
		V3179 := __e.Get(1)
		_ = V3179
		V3180 := __e.Get(2)
		_ = V3180
		V3181 := __e.Get(3)
		_ = V3181
		tmp17817 := Call(__e, PrimNS2Value(symshen_4spaces), V3179)

		tmp17818 := Call(__e, PrimNS2Value(symshen_4spaces), V3179)

		tmp17819 := Call(__e, PrimNS2Value(symshen_4app), V3181, MakeString(""), symshen_4s)

		tmp17820 := PrimStringConcat(MakeString("==> "), tmp17819)

		tmp17821 := Call(__e, PrimNS2Value(symshen_4app), tmp17818, tmp17820, symshen_4a)

		tmp17822 := PrimStringConcat(MakeString(" \n"), tmp17821)

		tmp17823 := Call(__e, PrimNS2Value(symshen_4app), V3180, tmp17822, symshen_4a)

		tmp17824 := PrimStringConcat(MakeString("> Output from "), tmp17823)

		tmp17825 := Call(__e, PrimNS2Value(symshen_4app), V3179, tmp17824, symshen_4a)

		tmp17826 := PrimStringConcat(MakeString("<"), tmp17825)

		tmp17827 := Call(__e, PrimNS2Value(symshen_4app), tmp17817, tmp17826, symshen_4a)

		tmp17828 := PrimStringConcat(MakeString("\n"), tmp17827)

		tmp17829 := Call(__e, PrimNS2Value(symstoutput))

		__e.TailApply(PrimNS2Value(symshen_4prhush), tmp17828, tmp17829)
		return

	}, 3)

	tmp17830 := Call(__e, PrimNS1Value(symns2_1set), symshen_4output_1track, tmp17816)

	_ = tmp17830

	tmp17831 := MakeNative(func(__e *ControlFlow) {
		V3183 := __e.Get(1)
		_ = V3183
		tmp17832 := MakeNative(func(__e *ControlFlow) {
			Tracking := __e.Get(1)
			_ = Tracking
			tmp17833 := MakeNative(func(__e *ControlFlow) {
				Tracking := __e.Get(1)
				_ = Tracking
				tmp17834 := Call(__e, PrimNS2Value(symps), V3183)

				__e.TailApply(PrimNS2Value(symeval), tmp17834)
				return

			}, 1)

			tmp17835 := Call(__e, PrimNS2Value(symremove), V3183, Tracking)

			tmp17836 := PrimNS3Set(symshen_4_dtracking_d, tmp17835)

			__e.TailApply(tmp17833, tmp17836)
			return

		}, 1)

		tmp17837 := PrimNS3Value(symshen_4_dtracking_d)

		__e.TailApply(tmp17832, tmp17837)
		return

	}, 1)

	tmp17838 := Call(__e, PrimNS1Value(symns2_1set), symuntrack, tmp17831)

	_ = tmp17838

	tmp17839 := MakeNative(func(__e *ControlFlow) {
		V3185 := __e.Get(1)
		_ = V3185
		tmp17840 := Call(__e, PrimNS2Value(symps), V3185)

		__e.TailApply(PrimNS2Value(symshen_4profile_1help), tmp17840)
		return

	}, 1)

	tmp17841 := Call(__e, PrimNS1Value(symns2_1set), symprofile, tmp17839)

	_ = tmp17841

	tmp17842 := MakeNative(func(__e *ControlFlow) {
		V3191 := __e.Get(1)
		_ = V3191
		tmp17912 := PrimIsPair(V3191)

		var ifres17886 Obj

		if True == tmp17912 {
			tmp17910 := PrimHead(V3191)

			tmp17911 := PrimEqual(symdefun, tmp17910)

			var ifres17888 Obj

			if True == tmp17911 {
				tmp17908 := PrimTail(V3191)

				tmp17909 := PrimIsPair(tmp17908)

				var ifres17890 Obj

				if True == tmp17909 {
					tmp17905 := PrimTail(V3191)

					tmp17906 := PrimTail(tmp17905)

					tmp17907 := PrimIsPair(tmp17906)

					var ifres17892 Obj

					if True == tmp17907 {
						tmp17901 := PrimTail(V3191)

						tmp17902 := PrimTail(tmp17901)

						tmp17903 := PrimTail(tmp17902)

						tmp17904 := PrimIsPair(tmp17903)

						var ifres17894 Obj

						if True == tmp17904 {
							tmp17896 := PrimTail(V3191)

							tmp17897 := PrimTail(tmp17896)

							tmp17898 := PrimTail(tmp17897)

							tmp17899 := PrimTail(tmp17898)

							tmp17900 := PrimEqual(Nil, tmp17899)

							var ifres17895 Obj

							if True == tmp17900 {
								ifres17895 = True

							} else {
								ifres17895 = False

							}

							ifres17894 = ifres17895

						} else {
							ifres17894 = False

						}

						var ifres17893 Obj

						if True == ifres17894 {
							ifres17893 = True

						} else {
							ifres17893 = False

						}

						ifres17892 = ifres17893

					} else {
						ifres17892 = False

					}

					var ifres17891 Obj

					if True == ifres17892 {
						ifres17891 = True

					} else {
						ifres17891 = False

					}

					ifres17890 = ifres17891

				} else {
					ifres17890 = False

				}

				var ifres17889 Obj

				if True == ifres17890 {
					ifres17889 = True

				} else {
					ifres17889 = False

				}

				ifres17888 = ifres17889

			} else {
				ifres17888 = False

			}

			var ifres17887 Obj

			if True == ifres17888 {
				ifres17887 = True

			} else {
				ifres17887 = False

			}

			ifres17886 = ifres17887

		} else {
			ifres17886 = False

		}

		if True == ifres17886 {
			tmp17844 := MakeNative(func(__e *ControlFlow) {
				G := __e.Get(1)
				_ = G
				tmp17845 := MakeNative(func(__e *ControlFlow) {
					Profile := __e.Get(1)
					_ = Profile
					tmp17846 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						tmp17847 := MakeNative(func(__e *ControlFlow) {
							CompileProfile := __e.Get(1)
							_ = CompileProfile
							tmp17848 := MakeNative(func(__e *ControlFlow) {
								CompileG := __e.Get(1)
								_ = CompileG
								tmp17849 := PrimTail(V3191)

								__e.Return(PrimHead(tmp17849))
								return

							}, 1)

							tmp17850 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), Def)

							__e.TailApply(tmp17848, tmp17850)
							return

						}, 1)

						tmp17851 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), Profile)

						__e.TailApply(tmp17847, tmp17851)
						return

					}, 1)

					tmp17852 := PrimTail(V3191)

					tmp17853 := PrimTail(tmp17852)

					tmp17854 := PrimHead(tmp17853)

					tmp17855 := PrimTail(V3191)

					tmp17856 := PrimHead(tmp17855)

					tmp17857 := PrimTail(V3191)

					tmp17858 := PrimTail(tmp17857)

					tmp17859 := PrimTail(tmp17858)

					tmp17860 := PrimHead(tmp17859)

					tmp17861 := Call(__e, PrimNS2Value(symsubst), G, tmp17856, tmp17860)

					tmp17862 := PrimCons(tmp17861, Nil)

					tmp17863 := PrimCons(tmp17854, tmp17862)

					tmp17864 := PrimCons(G, tmp17863)

					tmp17865 := PrimCons(symdefun, tmp17864)

					__e.TailApply(tmp17846, tmp17865)
					return

				}, 1)

				tmp17866 := PrimTail(V3191)

				tmp17867 := PrimHead(tmp17866)

				tmp17868 := PrimTail(V3191)

				tmp17869 := PrimTail(tmp17868)

				tmp17870 := PrimHead(tmp17869)

				tmp17871 := PrimTail(V3191)

				tmp17872 := PrimHead(tmp17871)

				tmp17873 := PrimTail(V3191)

				tmp17874 := PrimTail(tmp17873)

				tmp17875 := PrimHead(tmp17874)

				tmp17876 := PrimTail(V3191)

				tmp17877 := PrimTail(tmp17876)

				tmp17878 := PrimHead(tmp17877)

				tmp17879 := PrimCons(G, tmp17878)

				tmp17880 := Call(__e, PrimNS2Value(symshen_4profile_1func), tmp17872, tmp17875, tmp17879)

				tmp17881 := PrimCons(tmp17880, Nil)

				tmp17882 := PrimCons(tmp17870, tmp17881)

				tmp17883 := PrimCons(tmp17867, tmp17882)

				tmp17884 := PrimCons(symdefun, tmp17883)

				__e.TailApply(tmp17845, tmp17884)
				return

			}, 1)

			tmp17885 := Call(__e, PrimNS2Value(symgensym), symshen_4f)

			__e.TailApply(tmp17844, tmp17885)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
			return
		}

	}, 1)

	tmp17913 := Call(__e, PrimNS1Value(symns2_1set), symshen_4profile_1help, tmp17842)

	_ = tmp17913

	tmp17914 := MakeNative(func(__e *ControlFlow) {
		V3193 := __e.Get(1)
		_ = V3193
		__e.TailApply(PrimNS2Value(symuntrack), V3193)
		return
	}, 1)

	tmp17915 := Call(__e, PrimNS1Value(symns2_1set), symunprofile, tmp17914)

	_ = tmp17915

	tmp17916 := MakeNative(func(__e *ControlFlow) {
		V3197 := __e.Get(1)
		_ = V3197
		V3198 := __e.Get(2)
		_ = V3198
		V3199 := __e.Get(3)
		_ = V3199
		tmp17917 := PrimCons(symrun, Nil)

		tmp17918 := PrimCons(symget_1time, tmp17917)

		tmp17919 := PrimCons(symrun, Nil)

		tmp17920 := PrimCons(symget_1time, tmp17919)

		tmp17921 := PrimCons(symStart, Nil)

		tmp17922 := PrimCons(tmp17920, tmp17921)

		tmp17923 := PrimCons(sym_1, tmp17922)

		tmp17924 := PrimCons(V3197, Nil)

		tmp17925 := PrimCons(symshen_4get_1profile, tmp17924)

		tmp17926 := PrimCons(symFinish, Nil)

		tmp17927 := PrimCons(tmp17925, tmp17926)

		tmp17928 := PrimCons(sym_7, tmp17927)

		tmp17929 := PrimCons(tmp17928, Nil)

		tmp17930 := PrimCons(V3197, tmp17929)

		tmp17931 := PrimCons(symshen_4put_1profile, tmp17930)

		tmp17932 := PrimCons(symResult, Nil)

		tmp17933 := PrimCons(tmp17931, tmp17932)

		tmp17934 := PrimCons(symRecord, tmp17933)

		tmp17935 := PrimCons(symlet, tmp17934)

		tmp17936 := PrimCons(tmp17935, Nil)

		tmp17937 := PrimCons(tmp17923, tmp17936)

		tmp17938 := PrimCons(symFinish, tmp17937)

		tmp17939 := PrimCons(symlet, tmp17938)

		tmp17940 := PrimCons(tmp17939, Nil)

		tmp17941 := PrimCons(V3199, tmp17940)

		tmp17942 := PrimCons(symResult, tmp17941)

		tmp17943 := PrimCons(symlet, tmp17942)

		tmp17944 := PrimCons(tmp17943, Nil)

		tmp17945 := PrimCons(tmp17918, tmp17944)

		tmp17946 := PrimCons(symStart, tmp17945)

		__e.Return(PrimCons(symlet, tmp17946))
		return

	}, 3)

	tmp17947 := Call(__e, PrimNS1Value(symns2_1set), symshen_4profile_1func, tmp17916)

	_ = tmp17947

	tmp17948 := MakeNative(func(__e *ControlFlow) {
		V3201 := __e.Get(1)
		_ = V3201
		tmp17949 := MakeNative(func(__e *ControlFlow) {
			Results := __e.Get(1)
			_ = Results
			tmp17950 := MakeNative(func(__e *ControlFlow) {
				Initialise := __e.Get(1)
				_ = Initialise
				__e.TailApply(PrimNS2Value(sym_8p), V3201, Results)
				return
			}, 1)

			tmp17951 := Call(__e, PrimNS2Value(symshen_4put_1profile), V3201, MakeNumber(0))

			__e.TailApply(tmp17950, tmp17951)
			return

		}, 1)

		tmp17952 := Call(__e, PrimNS2Value(symshen_4get_1profile), V3201)

		__e.TailApply(tmp17949, tmp17952)
		return

	}, 1)

	tmp17953 := Call(__e, PrimNS1Value(symns2_1set), symprofile_1results, tmp17948)

	_ = tmp17953

	tmp17954 := MakeNative(func(__e *ControlFlow) {
		V3203 := __e.Get(1)
		_ = V3203
		tmp17955 := MakeNative(func(__e *ControlFlow) {
			tmp17956 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V3203, symprofile, tmp17956)
			return

		}, 0)

		tmp17957 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(0))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp17955, tmp17957)
		return

	}, 1)

	tmp17958 := Call(__e, PrimNS1Value(symns2_1set), symshen_4get_1profile, tmp17954)

	_ = tmp17958

	tmp17959 := MakeNative(func(__e *ControlFlow) {
		V3206 := __e.Get(1)
		_ = V3206
		V3207 := __e.Get(2)
		_ = V3207
		tmp17960 := PrimNS3Value(sym_dproperty_1vector_d)

		__e.TailApply(PrimNS2Value(symput), V3206, symprofile, V3207, tmp17960)
		return

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4put_1profile, tmp17959)
	return

}, 0)
