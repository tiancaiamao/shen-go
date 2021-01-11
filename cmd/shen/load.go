package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen19689 := MakeNative(func(__e *ControlFlow) {
		V701 := __e.Get(1)
		_ = V701
		gen19660 := MakeNative(func(__e *ControlFlow) {
			Load := __e.Get(1)
			_ = Load
			gen19647 := MakeNative(func(__e *ControlFlow) {
				Infs := __e.Get(1)
				_ = Infs
				__e.Return(symloaded)
				return
			}, 1)

			gen19657 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen19658 := Call(__e, gen19657, symshen_4_dtc_d)

			var gen19659 Obj

			if True == gen19658 {
				gen19648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				gen19649 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen19650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen19651 := Call(__e, PrimNS1Value(symns2_1value), syminferences)

				gen19652 := Call(__e, gen19651)

				gen19653 := Call(__e, gen19650, gen19652, MakeString(" inferences\n"), symshen_4a)

				gen19654 := Call(__e, gen19649, MakeString("\ntypechecked in "), gen19653)

				gen19655 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				gen19656 := Call(__e, gen19655)

				gen19659 = Call(__e, gen19648, gen19654, gen19656)

			} else {
				gen19659 = symshen_4skip
			}

			__e.TailApply(gen19647, gen19659)

			return

		}, 1)

		gen19685 := MakeNative(func(__e *ControlFlow) {
			Start := __e.Get(1)
			_ = Start
			gen19678 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				gen19675 := MakeNative(func(__e *ControlFlow) {
					Finish := __e.Get(1)
					_ = Finish
					gen19672 := MakeNative(func(__e *ControlFlow) {
						Time := __e.Get(1)
						_ = Time
						gen19661 := MakeNative(func(__e *ControlFlow) {
							Message := __e.Get(1)
							_ = Message
							__e.Return(Result)
							return
						}, 1)

						gen19662 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

						gen19663 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						gen19664 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						gen19665 := Call(__e, PrimNS1Value(symns2_1value), symstr)

						gen19666 := Call(__e, gen19665, Time)

						gen19667 := Call(__e, gen19664, gen19666, MakeString(" secs\n"))

						gen19668 := Call(__e, gen19663, MakeString("\nrun time: "), gen19667)

						gen19669 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

						gen19670 := Call(__e, gen19669)

						gen19671 := Call(__e, gen19662, gen19668, gen19670)

						__e.TailApply(gen19661, gen19671)

						return

					}, 1)

					gen19673 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen19674 := Call(__e, gen19673, Finish, Start)

					__e.TailApply(gen19672, gen19674)

					return

				}, 1)

				gen19676 := Call(__e, PrimNS1Value(symns2_1value), symget_1time)

				gen19677 := Call(__e, gen19676, symrun)

				__e.TailApply(gen19675, gen19677)

				return

			}, 1)

			gen19679 := Call(__e, PrimNS1Value(symns2_1value), symshen_4load_1help)

			gen19680 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen19681 := Call(__e, gen19680, symshen_4_dtc_d)

			gen19682 := Call(__e, PrimNS1Value(symns2_1value), symread_1file)

			gen19683 := Call(__e, gen19682, V701)

			gen19684 := Call(__e, gen19679, gen19681, gen19683)

			__e.TailApply(gen19678, gen19684)

			return

		}, 1)

		gen19686 := Call(__e, PrimNS1Value(symns2_1value), symget_1time)

		gen19687 := Call(__e, gen19686, symrun)

		gen19688 := Call(__e, gen19685, gen19687)

		__e.TailApply(gen19660, gen19688)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symload, gen19689)

	gen19723 := MakeNative(func(__e *ControlFlow) {
		V708 := __e.Get(1)
		_ = V708
		V709 := __e.Get(2)
		_ = V709
		gen19721 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen19722 := Call(__e, gen19721, False, V708)

		if True == gen19722 {
			gen19712 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

			gen19720 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen19713 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				gen19714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen19715 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				gen19716 := Call(__e, gen19715, X)

				gen19717 := Call(__e, gen19714, gen19716, MakeString("\n"), symshen_4s)

				gen19718 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				gen19719 := Call(__e, gen19718)

				__e.TailApply(gen19713, gen19717, gen19719)

				return

			}, 1)

			__e.TailApply(gen19712, gen19720, V709)

			return

		} else {
			if True == True {
				gen19707 := MakeNative(func(__e *ControlFlow) {
					RemoveSynonyms := __e.Get(1)
					_ = RemoveSynonyms
					gen19702 := MakeNative(func(__e *ControlFlow) {
						Table := __e.Get(1)
						_ = Table
						gen19697 := MakeNative(func(__e *ControlFlow) {
							Assume := __e.Get(1)
							_ = Assume
							gen19694 := MakeNative(func(__e *ControlFlow) {
								gen19691 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

								gen19693 := MakeNative(func(__e *ControlFlow) {
									X := __e.Get(1)
									_ = X
									gen19692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1load)

									__e.TailApply(gen19692, X)

									return

								}, 1)

								__e.TailApply(gen19691, gen19693, RemoveSynonyms)

								return

							}, 0)

							gen19696 := MakeNative(func(__e *ControlFlow) {
								E := __e.Get(1)
								_ = E
								gen19695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unwind_1types)

								__e.TailApply(gen19695, E, Table)

								return

							}, 1)

							__e.TailApply(PrimNS1Value(symtry_1catch), gen19694, gen19696)

							return

						}, 1)

						gen19698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

						gen19700 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							gen19699 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assumetype)

							__e.TailApply(gen19699, X)

							return

						}, 1)

						gen19701 := Call(__e, gen19698, gen19700, Table)

						__e.TailApply(gen19697, gen19701)

						return

					}, 1)

					gen19703 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

					gen19705 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						gen19704 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typetable)

						__e.TailApply(gen19704, X)

						return

					}, 1)

					gen19706 := Call(__e, gen19703, gen19705, RemoveSynonyms)

					__e.TailApply(gen19702, gen19706)

					return

				}, 1)

				gen19708 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

				gen19710 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen19709 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1synonyms)

					__e.TailApply(gen19709, X)

					return

				}, 1)

				gen19711 := Call(__e, gen19708, gen19710, V709)

				__e.TailApply(gen19707, gen19711)

				return

			} else {
				gen19690 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen19690, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4load_1help, gen19723)

	gen19734 := MakeNative(func(__e *ControlFlow) {
		V711 := __e.Get(1)
		_ = V711
		gen19731 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen19732 := Call(__e, gen19731, V711)

		var gen19733 Obj

		if True == gen19732 {
			gen19727 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen19728 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen19729 := Call(__e, gen19728, V711)

			gen19730 := Call(__e, gen19727, symshen_4synonyms_1help, gen19729)

			if True == gen19730 {
				gen19733 = True
			} else {
				gen19733 = False
			}

		} else {
			gen19733 = False
		}

		if True == gen19733 {
			gen19726 := Call(__e, PrimNS1Value(symns2_1value), symeval)

			Call(__e, gen19726, V711)

			__e.Return(Nil)
			return

		} else {
			if True == True {
				gen19725 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				__e.TailApply(gen19725, V711, Nil)

				return

			} else {
				gen19724 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen19724, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1synonyms, gen19734)

	gen19739 := MakeNative(func(__e *ControlFlow) {
		V713 := __e.Get(1)
		_ = V713
		gen19735 := Call(__e, PrimNS1Value(symns2_1value), symnl)

		Call(__e, gen19735, MakeNumber(1))

		gen19736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

		gen19737 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		gen19738 := Call(__e, gen19737, symA)

		__e.TailApply(gen19736, V713, gen19738)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1load, gen19739)

	gen19777 := MakeNative(func(__e *ControlFlow) {
		V719 := __e.Get(1)
		_ = V719
		gen19774 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen19775 := Call(__e, gen19774, V719)

		var gen19776 Obj

		if True == gen19775 {
			gen19769 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen19770 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen19771 := Call(__e, gen19770, V719)

			gen19772 := Call(__e, gen19769, symdefine, gen19771)

			var gen19773 Obj

			if True == gen19772 {
				gen19765 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen19766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19767 := Call(__e, gen19766, V719)

				gen19768 := Call(__e, gen19765, gen19767)

				if True == gen19768 {
					gen19773 = True
				} else {
					gen19773 = False
				}

			} else {
				gen19773 = False
			}

			if True == gen19773 {
				gen19776 = True
			} else {
				gen19776 = False
			}

		} else {
			gen19776 = False
		}

		if True == gen19776 {
			gen19748 := MakeNative(func(__e *ControlFlow) {
				Sig := __e.Get(1)
				_ = Sig
				gen19741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen19742 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen19743 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19744 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19745 := Call(__e, gen19744, V719)

				gen19746 := Call(__e, gen19743, gen19745)

				gen19747 := Call(__e, gen19742, gen19746, Sig)

				__e.TailApply(gen19741, gen19747, Nil)

				return

			}, 1)

			gen19749 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			gen19751 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				gen19750 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5sig_7rest_6)

				__e.TailApply(gen19750, Y)

				return

			}, 1)

			gen19752 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen19753 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen19754 := Call(__e, gen19753, V719)

			gen19755 := Call(__e, gen19752, gen19754)

			gen19763 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				gen19756 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				gen19757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen19758 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19760 := Call(__e, gen19759, V719)

				gen19761 := Call(__e, gen19758, gen19760)

				gen19762 := Call(__e, gen19757, gen19761, MakeString(" lacks a proper signature.\n"), symshen_4a)

				__e.TailApply(gen19756, gen19762)

				return

			}, 1)

			gen19764 := Call(__e, gen19749, gen19751, gen19755, gen19763)

			__e.TailApply(gen19748, gen19764)

			return

		} else {
			if True == True {
				__e.Return(Nil)
				return
			} else {
				gen19740 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen19740, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4typetable, gen19777)

	gen19787 := MakeNative(func(__e *ControlFlow) {
		V721 := __e.Get(1)
		_ = V721
		gen19785 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen19786 := Call(__e, gen19785, V721)

		if True == gen19786 {
			gen19780 := Call(__e, PrimNS1Value(symns2_1value), symdeclare)

			gen19781 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen19782 := Call(__e, gen19781, V721)

			gen19783 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen19784 := Call(__e, gen19783, V721)

			__e.TailApply(gen19780, gen19782, gen19784)

			return

		} else {
			if True == True {
				gen19779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen19779, symshen_4assumetype)

				return

			} else {
				gen19778 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen19778, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4assumetype, gen19787)

	gen19810 := MakeNative(func(__e *ControlFlow) {
		V728 := __e.Get(1)
		_ = V728
		V729 := __e.Get(2)
		_ = V729
		gen19808 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen19809 := Call(__e, gen19808, Nil, V729)

		if True == gen19809 {
			gen19805 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen19806 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

			gen19807 := Call(__e, gen19806, V728)

			__e.TailApply(gen19805, gen19807)

			return

		} else {
			gen19802 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen19803 := Call(__e, gen19802, V729)

			var gen19804 Obj

			if True == gen19803 {
				gen19798 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen19799 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19800 := Call(__e, gen19799, V729)

				gen19801 := Call(__e, gen19798, gen19800)

				if True == gen19801 {
					gen19804 = True
				} else {
					gen19804 = False
				}

			} else {
				gen19804 = False
			}

			if True == gen19804 {
				gen19790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remtype)

				gen19791 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19792 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19793 := Call(__e, gen19792, V729)

				gen19794 := Call(__e, gen19791, gen19793)

				Call(__e, gen19790, gen19794)

				gen19795 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unwind_1types)

				gen19796 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19797 := Call(__e, gen19796, V729)

				__e.TailApply(gen19795, V728, gen19797)

				return

			} else {
				if True == True {
					gen19789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen19789, symshen_4unwind_1types)

					return

				} else {
					gen19788 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen19788, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4unwind_1types, gen19810)

	gen19816 := MakeNative(func(__e *ControlFlow) {
		V731 := __e.Get(1)
		_ = V731
		gen19811 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen19812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

		gen19813 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen19814 := Call(__e, gen19813, symshen_4_dsignedfuncs_d)

		gen19815 := Call(__e, gen19812, V731, gen19814)

		__e.TailApply(gen19811, symshen_4_dsignedfuncs_d, gen19815)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4remtype, gen19816)

	gen19851 := MakeNative(func(__e *ControlFlow) {
		V739 := __e.Get(1)
		_ = V739
		V740 := __e.Get(2)
		_ = V740
		gen19849 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen19850 := Call(__e, gen19849, Nil, V740)

		if True == gen19850 {
			__e.Return(Nil)
			return
		} else {
			gen19846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen19847 := Call(__e, gen19846, V740)

			var gen19848 Obj

			if True == gen19847 {
				gen19841 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen19842 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19843 := Call(__e, gen19842, V740)

				gen19844 := Call(__e, gen19841, gen19843)

				var gen19845 Obj

				if True == gen19844 {
					gen19835 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen19836 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen19837 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen19838 := Call(__e, gen19837, V740)

					gen19839 := Call(__e, gen19836, gen19838)

					gen19840 := Call(__e, gen19835, gen19839, V739)

					if True == gen19840 {
						gen19845 = True
					} else {
						gen19845 = False
					}

				} else {
					gen19845 = False
				}

				if True == gen19845 {
					gen19848 = True
				} else {
					gen19848 = False
				}

			} else {
				gen19848 = False
			}

			if True == gen19848 {
				gen19828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

				gen19829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19830 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen19831 := Call(__e, gen19830, V740)

				gen19832 := Call(__e, gen19829, gen19831)

				gen19833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen19834 := Call(__e, gen19833, V740)

				__e.TailApply(gen19828, gen19832, gen19834)

				return

			} else {
				gen19826 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen19827 := Call(__e, gen19826, V740)

				if True == gen19827 {
					gen19819 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen19820 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen19821 := Call(__e, gen19820, V740)

					gen19822 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

					gen19823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen19824 := Call(__e, gen19823, V740)

					gen19825 := Call(__e, gen19822, V739, gen19824)

					__e.TailApply(gen19819, gen19821, gen19825)

					return

				} else {
					if True == True {
						gen19818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen19818, symshen_4removetype)

						return

					} else {
						gen19817 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen19817, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4removetype, gen19851)

	gen19877 := MakeNative(func(__e *ControlFlow) {
		V742 := __e.Get(1)
		_ = V742
		gen19874 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5signature_6 := __e.Get(1)
			_ = Parse__shen_4_5signature_6
			gen19868 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen19869 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen19870 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen19871 := Call(__e, gen19870)

			gen19872 := Call(__e, gen19869, gen19871, Parse__shen_4_5signature_6)

			gen19873 := Call(__e, gen19868, gen19872)

			if True == gen19873 {
				gen19865 := MakeNative(func(__e *ControlFlow) {
					Parse___5_b_6 := __e.Get(1)
					_ = Parse___5_b_6
					gen19859 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen19860 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen19861 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen19862 := Call(__e, gen19861)

					gen19863 := Call(__e, gen19860, gen19862, Parse___5_b_6)

					gen19864 := Call(__e, gen19859, gen19863)

					if True == gen19864 {
						gen19854 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen19855 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen19856 := Call(__e, gen19855, Parse___5_b_6)

						gen19857 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen19858 := Call(__e, gen19857, Parse__shen_4_5signature_6)

						__e.TailApply(gen19854, gen19856, gen19858)

						return

					} else {
						gen19853 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen19853)

						return

					}

				}, 1)

				gen19866 := Call(__e, PrimNS1Value(symns2_1value), sym_5_b_6)

				gen19867 := Call(__e, gen19866, Parse__shen_4_5signature_6)

				__e.TailApply(gen19865, gen19867)

				return

			} else {
				gen19852 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen19852)

				return

			}

		}, 1)

		gen19875 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_6)

		gen19876 := Call(__e, gen19875, V742)

		__e.TailApply(gen19874, gen19876)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sig_7rest_6, gen19877)

	gen19893 := MakeNative(func(__e *ControlFlow) {
		V745 := __e.Get(1)
		_ = V745
		V746 := __e.Get(2)
		_ = V746
		gen19890 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			gen19884 := MakeNative(func(__e *ControlFlow) {
				String := __e.Get(1)
				_ = String
				gen19881 := MakeNative(func(__e *ControlFlow) {
					Write := __e.Get(1)
					_ = Write
					gen19878 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.Return(V746)
						return
					}, 1)

					gen19879 := Call(__e, PrimNS1Value(symns2_1value), symclose)

					gen19880 := Call(__e, gen19879, Stream)

					__e.TailApply(gen19878, gen19880)

					return

				}, 1)

				gen19882 := Call(__e, PrimNS1Value(symns2_1value), sympr)

				gen19883 := Call(__e, gen19882, String, Stream)

				__e.TailApply(gen19881, gen19883)

				return

			}, 1)

			gen19887 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

			gen19888 := Call(__e, gen19887, V746)

			var gen19889 Obj

			if True == gen19888 {
				gen19886 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen19889 = Call(__e, gen19886, V746, MakeString("\n\n"), symshen_4a)

			} else {
				gen19885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen19889 = Call(__e, gen19885, V746, MakeString("\n\n"), symshen_4s)

			}

			__e.TailApply(gen19884, gen19889)

			return

		}, 1)

		gen19891 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		gen19892 := Call(__e, gen19891, V745, symout)

		__e.TailApply(gen19890, gen19892)

		return

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symwrite_1to_1file, gen19893)

	return

}, 0)
