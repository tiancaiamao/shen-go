package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen20684 := MakeNative(func(__e *ControlFlow) {
		V748 := __e.Get(1)
		_ = V748
		gen20679 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			gen20677 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20678 := Call(__e, gen20677, V748, Y)

			if True == gen20678 {
				__e.Return(V748)
				return
			} else {
				gen20674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

				gen20676 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					gen20675 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

					__e.TailApply(gen20675, Z)

					return

				}, 1)

				__e.TailApply(gen20674, gen20676, Y)

				return

			}

		}, 1)

		gen20680 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compose)

		gen20681 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen20682 := Call(__e, gen20681, sym_dmacros_d)

		gen20683 := Call(__e, gen20680, gen20682, V748)

		__e.TailApply(gen20679, gen20683)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symmacroexpand, gen20684)

	gen20711 := MakeNative(func(__e *ControlFlow) {
		V750 := __e.Get(1)
		_ = V750
		gen20708 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20709 := Call(__e, gen20708, V750)

		var gen20710 Obj

		if True == gen20709 {
			gen20703 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20704 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20705 := Call(__e, gen20704, V750)

			gen20706 := Call(__e, gen20703, symerror, gen20705)

			var gen20707 Obj

			if True == gen20706 {
				gen20699 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20701 := Call(__e, gen20700, V750)

				gen20702 := Call(__e, gen20699, gen20701)

				if True == gen20702 {
					gen20707 = True
				} else {
					gen20707 = False
				}

			} else {
				gen20707 = False
			}

			if True == gen20707 {
				gen20710 = True
			} else {
				gen20710 = False
			}

		} else {
			gen20710 = False
		}

		if True == gen20710 {
			gen20686 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20688 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			gen20689 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20690 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20691 := Call(__e, gen20690, V750)

			gen20692 := Call(__e, gen20689, gen20691)

			gen20693 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20694 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20695 := Call(__e, gen20694, V750)

			gen20696 := Call(__e, gen20693, gen20695)

			gen20697 := Call(__e, gen20688, gen20692, gen20696)

			gen20698 := Call(__e, gen20687, gen20697, Nil)

			__e.TailApply(gen20686, symsimple_1error, gen20698)

			return

		} else {
			if True == True {
				__e.Return(V750)
				return
			} else {
				gen20685 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20685, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4error_1macro, gen20711)

	gen20772 := MakeNative(func(__e *ControlFlow) {
		V752 := __e.Get(1)
		_ = V752
		gen20769 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20770 := Call(__e, gen20769, V752)

		var gen20771 Obj

		if True == gen20770 {
			gen20764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20766 := Call(__e, gen20765, V752)

			gen20767 := Call(__e, gen20764, symoutput, gen20766)

			var gen20768 Obj

			if True == gen20767 {
				gen20760 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20762 := Call(__e, gen20761, V752)

				gen20763 := Call(__e, gen20760, gen20762)

				if True == gen20763 {
					gen20768 = True
				} else {
					gen20768 = False
				}

			} else {
				gen20768 = False
			}

			if True == gen20768 {
				gen20771 = True
			} else {
				gen20771 = False
			}

		} else {
			gen20771 = False
		}

		if True == gen20771 {
			gen20743 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20744 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			gen20746 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20748 := Call(__e, gen20747, V752)

			gen20749 := Call(__e, gen20746, gen20748)

			gen20750 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20751 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20752 := Call(__e, gen20751, V752)

			gen20753 := Call(__e, gen20750, gen20752)

			gen20754 := Call(__e, gen20745, gen20749, gen20753)

			gen20755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20757 := Call(__e, gen20756, symstoutput, Nil)

			gen20758 := Call(__e, gen20755, gen20757, Nil)

			gen20759 := Call(__e, gen20744, gen20754, gen20758)

			__e.TailApply(gen20743, symshen_4prhush, gen20759)

			return

		} else {
			gen20740 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen20741 := Call(__e, gen20740, V752)

			var gen20742 Obj

			if True == gen20741 {
				gen20735 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20736 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen20737 := Call(__e, gen20736, V752)

				gen20738 := Call(__e, gen20735, sympr, gen20737)

				var gen20739 Obj

				if True == gen20738 {
					gen20730 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen20731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20732 := Call(__e, gen20731, V752)

					gen20733 := Call(__e, gen20730, gen20732)

					var gen20734 Obj

					if True == gen20733 {
						gen20724 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20725 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20726 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20727 := Call(__e, gen20726, V752)

						gen20728 := Call(__e, gen20725, gen20727)

						gen20729 := Call(__e, gen20724, Nil, gen20728)

						if True == gen20729 {
							gen20734 = True
						} else {
							gen20734 = False
						}

					} else {
						gen20734 = False
					}

					if True == gen20734 {
						gen20739 = True
					} else {
						gen20739 = False
					}

				} else {
					gen20739 = False
				}

				if True == gen20739 {
					gen20742 = True
				} else {
					gen20742 = False
				}

			} else {
				gen20742 = False
			}

			if True == gen20742 {
				gen20713 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20714 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20715 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen20716 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20717 := Call(__e, gen20716, V752)

				gen20718 := Call(__e, gen20715, gen20717)

				gen20719 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20720 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20721 := Call(__e, gen20720, symstoutput, Nil)

				gen20722 := Call(__e, gen20719, gen20721, Nil)

				gen20723 := Call(__e, gen20714, gen20718, gen20722)

				__e.TailApply(gen20713, sympr, gen20723)

				return

			} else {
				if True == True {
					__e.Return(V752)
					return
				} else {
					gen20712 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen20712, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4output_1macro, gen20772)

	gen20795 := MakeNative(func(__e *ControlFlow) {
		V754 := __e.Get(1)
		_ = V754
		gen20792 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20793 := Call(__e, gen20792, V754)

		var gen20794 Obj

		if True == gen20793 {
			gen20787 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20788 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20789 := Call(__e, gen20788, V754)

			gen20790 := Call(__e, gen20787, symmake_1string, gen20789)

			var gen20791 Obj

			if True == gen20790 {
				gen20783 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20785 := Call(__e, gen20784, V754)

				gen20786 := Call(__e, gen20783, gen20785)

				if True == gen20786 {
					gen20791 = True
				} else {
					gen20791 = False
				}

			} else {
				gen20791 = False
			}

			if True == gen20791 {
				gen20794 = True
			} else {
				gen20794 = False
			}

		} else {
			gen20794 = False
		}

		if True == gen20794 {
			gen20774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mkstr)

			gen20775 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20777 := Call(__e, gen20776, V754)

			gen20778 := Call(__e, gen20775, gen20777)

			gen20779 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20781 := Call(__e, gen20780, V754)

			gen20782 := Call(__e, gen20779, gen20781)

			__e.TailApply(gen20774, gen20778, gen20782)

			return

		} else {
			if True == True {
				__e.Return(V754)
				return
			} else {
				gen20773 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20773, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1string_1macro, gen20795)

	gen20895 := MakeNative(func(__e *ControlFlow) {
		V756 := __e.Get(1)
		_ = V756
		gen20892 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20893 := Call(__e, gen20892, V756)

		var gen20894 Obj

		if True == gen20893 {
			gen20887 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20888 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20889 := Call(__e, gen20888, V756)

			gen20890 := Call(__e, gen20887, symlineread, gen20889)

			var gen20891 Obj

			if True == gen20890 {
				gen20883 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20884 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20885 := Call(__e, gen20884, V756)

				gen20886 := Call(__e, gen20883, Nil, gen20885)

				if True == gen20886 {
					gen20891 = True
				} else {
					gen20891 = False
				}

			} else {
				gen20891 = False
			}

			if True == gen20891 {
				gen20894 = True
			} else {
				gen20894 = False
			}

		} else {
			gen20894 = False
		}

		if True == gen20894 {
			gen20878 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20879 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20880 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20881 := Call(__e, gen20880, symstinput, Nil)

			gen20882 := Call(__e, gen20879, gen20881, Nil)

			__e.TailApply(gen20878, symlineread, gen20882)

			return

		} else {
			gen20875 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen20876 := Call(__e, gen20875, V756)

			var gen20877 Obj

			if True == gen20876 {
				gen20870 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen20871 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen20872 := Call(__e, gen20871, V756)

				gen20873 := Call(__e, gen20870, syminput, gen20872)

				var gen20874 Obj

				if True == gen20873 {
					gen20866 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20868 := Call(__e, gen20867, V756)

					gen20869 := Call(__e, gen20866, Nil, gen20868)

					if True == gen20869 {
						gen20874 = True
					} else {
						gen20874 = False
					}

				} else {
					gen20874 = False
				}

				if True == gen20874 {
					gen20877 = True
				} else {
					gen20877 = False
				}

			} else {
				gen20877 = False
			}

			if True == gen20877 {
				gen20861 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen20864 := Call(__e, gen20863, symstinput, Nil)

				gen20865 := Call(__e, gen20862, gen20864, Nil)

				__e.TailApply(gen20861, syminput, gen20865)

				return

			} else {
				gen20858 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20859 := Call(__e, gen20858, V756)

				var gen20860 Obj

				if True == gen20859 {
					gen20853 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen20854 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen20855 := Call(__e, gen20854, V756)

					gen20856 := Call(__e, gen20853, symread, gen20855)

					var gen20857 Obj

					if True == gen20856 {
						gen20849 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20850 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20851 := Call(__e, gen20850, V756)

						gen20852 := Call(__e, gen20849, Nil, gen20851)

						if True == gen20852 {
							gen20857 = True
						} else {
							gen20857 = False
						}

					} else {
						gen20857 = False
					}

					if True == gen20857 {
						gen20860 = True
					} else {
						gen20860 = False
					}

				} else {
					gen20860 = False
				}

				if True == gen20860 {
					gen20844 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20845 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20846 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen20847 := Call(__e, gen20846, symstinput, Nil)

					gen20848 := Call(__e, gen20845, gen20847, Nil)

					__e.TailApply(gen20844, symread, gen20848)

					return

				} else {
					gen20841 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen20842 := Call(__e, gen20841, V756)

					var gen20843 Obj

					if True == gen20842 {
						gen20836 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20837 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen20838 := Call(__e, gen20837, V756)

						gen20839 := Call(__e, gen20836, syminput_7, gen20838)

						var gen20840 Obj

						if True == gen20839 {
							gen20831 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen20832 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen20833 := Call(__e, gen20832, V756)

							gen20834 := Call(__e, gen20831, gen20833)

							var gen20835 Obj

							if True == gen20834 {
								gen20825 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen20826 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20827 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20828 := Call(__e, gen20827, V756)

								gen20829 := Call(__e, gen20826, gen20828)

								gen20830 := Call(__e, gen20825, Nil, gen20829)

								if True == gen20830 {
									gen20835 = True
								} else {
									gen20835 = False
								}

							} else {
								gen20835 = False
							}

							if True == gen20835 {
								gen20840 = True
							} else {
								gen20840 = False
							}

						} else {
							gen20840 = False
						}

						if True == gen20840 {
							gen20843 = True
						} else {
							gen20843 = False
						}

					} else {
						gen20843 = False
					}

					if True == gen20843 {
						gen20814 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20816 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen20817 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20818 := Call(__e, gen20817, V756)

						gen20819 := Call(__e, gen20816, gen20818)

						gen20820 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20821 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen20822 := Call(__e, gen20821, symstinput, Nil)

						gen20823 := Call(__e, gen20820, gen20822, Nil)

						gen20824 := Call(__e, gen20815, gen20819, gen20823)

						__e.TailApply(gen20814, syminput_7, gen20824)

						return

					} else {
						gen20811 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen20812 := Call(__e, gen20811, V756)

						var gen20813 Obj

						if True == gen20812 {
							gen20806 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen20807 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen20808 := Call(__e, gen20807, V756)

							gen20809 := Call(__e, gen20806, symread_1byte, gen20808)

							var gen20810 Obj

							if True == gen20809 {
								gen20802 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen20803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen20804 := Call(__e, gen20803, V756)

								gen20805 := Call(__e, gen20802, Nil, gen20804)

								if True == gen20805 {
									gen20810 = True
								} else {
									gen20810 = False
								}

							} else {
								gen20810 = False
							}

							if True == gen20810 {
								gen20813 = True
							} else {
								gen20813 = False
							}

						} else {
							gen20813 = False
						}

						if True == gen20813 {
							gen20797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20799 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen20800 := Call(__e, gen20799, symstinput, Nil)

							gen20801 := Call(__e, gen20798, gen20800, Nil)

							__e.TailApply(gen20797, symread_1byte, gen20801)

							return

						} else {
							if True == True {
								__e.Return(V756)
								return
							} else {
								gen20796 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								__e.TailApply(gen20796, MakeString("no cond match"))

								return

							}
						}

					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4input_1macro, gen20895)

	gen20908 := MakeNative(func(__e *ControlFlow) {
		V759 := __e.Get(1)
		_ = V759
		V760 := __e.Get(2)
		_ = V760
		gen20906 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen20907 := Call(__e, gen20906, Nil, V759)

		if True == gen20907 {
			__e.Return(V760)
			return
		} else {
			gen20904 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen20905 := Call(__e, gen20904, V759)

			if True == gen20905 {
				gen20898 := Call(__e, PrimNS1Value(symns2_1value), symshen_4compose)

				gen20899 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20900 := Call(__e, gen20899, V759)

				gen20901 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen20902 := Call(__e, gen20901, V759)

				gen20903 := Call(__e, gen20902, V760)

				__e.TailApply(gen20898, gen20900, gen20903)

				return

			} else {
				if True == True {
					gen20897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen20897, symshen_4compose)

					return

				} else {
					gen20896 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen20896, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4compose, gen20908)

	gen20983 := MakeNative(func(__e *ControlFlow) {
		V762 := __e.Get(1)
		_ = V762
		gen20980 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen20981 := Call(__e, gen20980, V762)

		var gen20982 Obj

		if True == gen20981 {
			gen20975 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen20976 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20977 := Call(__e, gen20976, V762)

			gen20978 := Call(__e, gen20975, symcompile, gen20977)

			var gen20979 Obj

			if True == gen20978 {
				gen20970 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen20971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen20972 := Call(__e, gen20971, V762)

				gen20973 := Call(__e, gen20970, gen20972)

				var gen20974 Obj

				if True == gen20973 {
					gen20963 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen20964 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20966 := Call(__e, gen20965, V762)

					gen20967 := Call(__e, gen20964, gen20966)

					gen20968 := Call(__e, gen20963, gen20967)

					var gen20969 Obj

					if True == gen20968 {
						gen20955 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen20956 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20958 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen20959 := Call(__e, gen20958, V762)

						gen20960 := Call(__e, gen20957, gen20959)

						gen20961 := Call(__e, gen20956, gen20960)

						gen20962 := Call(__e, gen20955, Nil, gen20961)

						if True == gen20962 {
							gen20969 = True
						} else {
							gen20969 = False
						}

					} else {
						gen20969 = False
					}

					if True == gen20969 {
						gen20974 = True
					} else {
						gen20974 = False
					}

				} else {
					gen20974 = False
				}

				if True == gen20974 {
					gen20979 = True
				} else {
					gen20979 = False
				}

			} else {
				gen20979 = False
			}

			if True == gen20979 {
				gen20982 = True
			} else {
				gen20982 = False
			}

		} else {
			gen20982 = False
		}

		if True == gen20982 {
			gen20910 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20911 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20912 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20913 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20914 := Call(__e, gen20913, V762)

			gen20915 := Call(__e, gen20912, gen20914)

			gen20916 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20917 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen20918 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20919 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen20920 := Call(__e, gen20919, V762)

			gen20921 := Call(__e, gen20918, gen20920)

			gen20922 := Call(__e, gen20917, gen20921)

			gen20923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20927 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20931 := Call(__e, gen20930, symE, Nil)

			gen20932 := Call(__e, gen20929, symcons_2, gen20931)

			gen20933 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20934 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20937 := Call(__e, gen20936, symE, Nil)

			gen20938 := Call(__e, gen20935, MakeString("parse error here: ~S~%"), gen20937)

			gen20939 := Call(__e, gen20934, symerror, gen20938)

			gen20940 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20941 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20942 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20943 := Call(__e, gen20942, MakeString("parse error~%"), Nil)

			gen20944 := Call(__e, gen20941, symerror, gen20943)

			gen20945 := Call(__e, gen20940, gen20944, Nil)

			gen20946 := Call(__e, gen20933, gen20939, gen20945)

			gen20947 := Call(__e, gen20928, gen20932, gen20946)

			gen20948 := Call(__e, gen20927, symif, gen20947)

			gen20949 := Call(__e, gen20926, gen20948, Nil)

			gen20950 := Call(__e, gen20925, symE, gen20949)

			gen20951 := Call(__e, gen20924, symlambda, gen20950)

			gen20952 := Call(__e, gen20923, gen20951, Nil)

			gen20953 := Call(__e, gen20916, gen20922, gen20952)

			gen20954 := Call(__e, gen20911, gen20915, gen20953)

			__e.TailApply(gen20910, symcompile, gen20954)

			return

		} else {
			if True == True {
				__e.Return(V762)
				return
			} else {
				gen20909 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20909, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4compile_1macro, gen20983)

	gen21021 := MakeNative(func(__e *ControlFlow) {
		V764 := __e.Get(1)
		_ = V764
		gen21018 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21019 := Call(__e, gen21018, V764)

		var gen21020 Obj

		if True == gen21019 {
			gen21014 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21015 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21016 := Call(__e, gen21015, V764)

			gen21017 := Call(__e, gen21014, symprolog_2, gen21016)

			if True == gen21017 {
				gen21020 = True
			} else {
				gen21020 = False
			}

		} else {
			gen21020 = False
		}

		if True == gen21020 {
			gen20985 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20986 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20988 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen20989 := Call(__e, gen20988, symshen_4start_1new_1prolog_1process, Nil)

			gen20990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21005 := MakeNative(func(__e *ControlFlow) {
				Calls := __e.Get(1)
				_ = Calls
				gen21000 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					gen20995 := MakeNative(func(__e *ControlFlow) {
						External := __e.Get(1)
						_ = External
						gen20992 := MakeNative(func(__e *ControlFlow) {
							PrologVs := __e.Get(1)
							_ = PrologVs
							gen20991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4locally_1bind_1prolog_1vs)

							__e.TailApply(gen20991, symNPP, PrologVs, Calls)

							return

						}, 1)

						gen20993 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

						gen20994 := Call(__e, gen20993, Vs, External)

						__e.TailApply(gen20992, gen20994)

						return

					}, 1)

					gen20996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

					gen20997 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen20998 := Call(__e, gen20997, V764)

					gen20999 := Call(__e, gen20996, gen20998)

					__e.TailApply(gen20995, gen20999)

					return

				}, 1)

				gen21001 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				gen21002 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21003 := Call(__e, gen21002, V764)

				gen21004 := Call(__e, gen21001, gen21003)

				__e.TailApply(gen21000, gen21004)

				return

			}, 1)

			gen21006 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

			gen21007 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21008 := Call(__e, gen21007, V764)

			gen21009 := Call(__e, gen21006, symNPP, gen21008)

			gen21010 := Call(__e, gen21005, gen21009)

			gen21011 := Call(__e, gen20990, gen21010, Nil)

			gen21012 := Call(__e, gen20987, gen20989, gen21011)

			gen21013 := Call(__e, gen20986, symNPP, gen21012)

			__e.TailApply(gen20985, symlet, gen21013)

			return

		} else {
			if True == True {
				__e.Return(V764)
				return
			} else {
				gen20984 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen20984, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1macro, gen21021)

	gen21054 := MakeNative(func(__e *ControlFlow) {
		V770 := __e.Get(1)
		_ = V770
		gen21051 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21052 := Call(__e, gen21051, V770)

		var gen21053 Obj

		if True == gen21052 {
			gen21046 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21047 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21048 := Call(__e, gen21047, V770)

			gen21049 := Call(__e, gen21046, symreceive, gen21048)

			var gen21050 Obj

			if True == gen21049 {
				gen21041 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21043 := Call(__e, gen21042, V770)

				gen21044 := Call(__e, gen21041, gen21043)

				var gen21045 Obj

				if True == gen21044 {
					gen21035 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen21036 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21038 := Call(__e, gen21037, V770)

					gen21039 := Call(__e, gen21036, gen21038)

					gen21040 := Call(__e, gen21035, Nil, gen21039)

					if True == gen21040 {
						gen21045 = True
					} else {
						gen21045 = False
					}

				} else {
					gen21045 = False
				}

				if True == gen21045 {
					gen21050 = True
				} else {
					gen21050 = False
				}

			} else {
				gen21050 = False
			}

			if True == gen21050 {
				gen21053 = True
			} else {
				gen21053 = False
			}

		} else {
			gen21053 = False
		}

		if True == gen21053 {
			gen21034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(gen21034, V770)

			return

		} else {
			gen21032 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21033 := Call(__e, gen21032, V770)

			if True == gen21033 {
				gen21023 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				gen21024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

				gen21025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21026 := Call(__e, gen21025, V770)

				gen21027 := Call(__e, gen21024, gen21026)

				gen21028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4externally_1bound)

				gen21029 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21030 := Call(__e, gen21029, V770)

				gen21031 := Call(__e, gen21028, gen21030)

				__e.TailApply(gen21023, gen21027, gen21031)

				return

			} else {
				if True == True {
					__e.Return(Nil)
					return
				} else {
					gen21022 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen21022, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4externally_1bound, gen21054)

	gen21078 := MakeNative(func(__e *ControlFlow) {
		V788 := __e.Get(1)
		_ = V788
		V789 := __e.Get(2)
		_ = V789
		V790 := __e.Get(3)
		_ = V790
		gen21076 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen21077 := Call(__e, gen21076, Nil, V789)

		if True == gen21077 {
			__e.Return(V790)
			return
		} else {
			gen21074 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21075 := Call(__e, gen21074, V789)

			if True == gen21075 {
				gen21057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21059 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21060 := Call(__e, gen21059, V789)

				gen21061 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21062 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21063 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21064 := Call(__e, gen21063, V788, Nil)

				gen21065 := Call(__e, gen21062, symshen_4newpv, gen21064)

				gen21066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4locally_1bind_1prolog_1vs)

				gen21068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21069 := Call(__e, gen21068, V789)

				gen21070 := Call(__e, gen21067, V788, gen21069, V790)

				gen21071 := Call(__e, gen21066, gen21070, Nil)

				gen21072 := Call(__e, gen21061, gen21065, gen21071)

				gen21073 := Call(__e, gen21058, gen21060, gen21072)

				__e.TailApply(gen21057, symlet, gen21073)

				return

			} else {
				if True == True {
					gen21056 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen21056, MakeString("implementation error inp locally-bind-prolog-vs"))

					return

				} else {
					gen21055 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen21055, MakeString("no cond match"))

					return

				}
			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4locally_1bind_1prolog_1vs, gen21078)

	gen21409 := MakeNative(func(__e *ControlFlow) {
		V803 := __e.Get(1)
		_ = V803
		V804 := __e.Get(2)
		_ = V804
		gen21407 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen21408 := Call(__e, gen21407, Nil, V804)

		if True == gen21408 {
			__e.Return(True)
			return
		} else {
			gen21404 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21405 := Call(__e, gen21404, V804)

			var gen21406 Obj

			if True == gen21405 {
				gen21400 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen21401 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21402 := Call(__e, gen21401, V804)

				gen21403 := Call(__e, gen21400, sym_b, gen21402)

				if True == gen21403 {
					gen21406 = True
				} else {
					gen21406 = False
				}

			} else {
				gen21406 = False
			}

			if True == gen21406 {
				gen21385 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21386 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21387 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21388 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21389 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

				gen21392 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21393 := Call(__e, gen21392, V804)

				gen21394 := Call(__e, gen21391, V803, gen21393)

				gen21395 := Call(__e, gen21390, gen21394, Nil)

				gen21396 := Call(__e, gen21389, symfreeze, gen21395)

				gen21397 := Call(__e, gen21388, gen21396, Nil)

				gen21398 := Call(__e, gen21387, V803, gen21397)

				gen21399 := Call(__e, gen21386, False, gen21398)

				__e.TailApply(gen21385, symcut, gen21399)

				return

			} else {
				gen21382 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21383 := Call(__e, gen21382, V804)

				var gen21384 Obj

				if True == gen21383 {
					gen21377 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21378 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21379 := Call(__e, gen21378, V804)

					gen21380 := Call(__e, gen21377, gen21379)

					var gen21381 Obj

					if True == gen21380 {
						gen21370 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen21371 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21372 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21373 := Call(__e, gen21372, V804)

						gen21374 := Call(__e, gen21371, gen21373)

						gen21375 := Call(__e, gen21370, symwhen, gen21374)

						var gen21376 Obj

						if True == gen21375 {
							gen21363 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21364 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen21366 := Call(__e, gen21365, V804)

							gen21367 := Call(__e, gen21364, gen21366)

							gen21368 := Call(__e, gen21363, gen21367)

							var gen21369 Obj

							if True == gen21368 {
								gen21355 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen21356 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21357 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21359 := Call(__e, gen21358, V804)

								gen21360 := Call(__e, gen21357, gen21359)

								gen21361 := Call(__e, gen21356, gen21360)

								gen21362 := Call(__e, gen21355, Nil, gen21361)

								if True == gen21362 {
									gen21369 = True
								} else {
									gen21369 = False
								}

							} else {
								gen21369 = False
							}

							if True == gen21369 {
								gen21376 = True
							} else {
								gen21376 = False
							}

						} else {
							gen21376 = False
						}

						if True == gen21376 {
							gen21381 = True
						} else {
							gen21381 = False
						}

					} else {
						gen21381 = False
					}

					if True == gen21381 {
						gen21384 = True
					} else {
						gen21384 = False
					}

				} else {
					gen21384 = False
				}

				if True == gen21384 {
					gen21332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21334 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

					gen21335 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21336 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21337 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21338 := Call(__e, gen21337, V804)

					gen21339 := Call(__e, gen21336, gen21338)

					gen21340 := Call(__e, gen21335, gen21339)

					gen21341 := Call(__e, gen21334, gen21340, V803)

					gen21342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21346 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

					gen21347 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21348 := Call(__e, gen21347, V804)

					gen21349 := Call(__e, gen21346, V803, gen21348)

					gen21350 := Call(__e, gen21345, gen21349, Nil)

					gen21351 := Call(__e, gen21344, symfreeze, gen21350)

					gen21352 := Call(__e, gen21343, gen21351, Nil)

					gen21353 := Call(__e, gen21342, V803, gen21352)

					gen21354 := Call(__e, gen21333, gen21341, gen21353)

					__e.TailApply(gen21332, symfwhen, gen21354)

					return

				} else {
					gen21329 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21330 := Call(__e, gen21329, V804)

					var gen21331 Obj

					if True == gen21330 {
						gen21324 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21325 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21326 := Call(__e, gen21325, V804)

						gen21327 := Call(__e, gen21324, gen21326)

						var gen21328 Obj

						if True == gen21327 {
							gen21317 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen21318 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen21319 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen21320 := Call(__e, gen21319, V804)

							gen21321 := Call(__e, gen21318, gen21320)

							gen21322 := Call(__e, gen21317, symis, gen21321)

							var gen21323 Obj

							if True == gen21322 {
								gen21310 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen21311 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21312 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21313 := Call(__e, gen21312, V804)

								gen21314 := Call(__e, gen21311, gen21313)

								gen21315 := Call(__e, gen21310, gen21314)

								var gen21316 Obj

								if True == gen21315 {
									gen21301 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen21302 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen21303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen21304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21305 := Call(__e, gen21304, V804)

									gen21306 := Call(__e, gen21303, gen21305)

									gen21307 := Call(__e, gen21302, gen21306)

									gen21308 := Call(__e, gen21301, gen21307)

									var gen21309 Obj

									if True == gen21308 {
										gen21291 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen21292 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21293 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21294 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21295 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21296 := Call(__e, gen21295, V804)

										gen21297 := Call(__e, gen21294, gen21296)

										gen21298 := Call(__e, gen21293, gen21297)

										gen21299 := Call(__e, gen21292, gen21298)

										gen21300 := Call(__e, gen21291, Nil, gen21299)

										if True == gen21300 {
											gen21309 = True
										} else {
											gen21309 = False
										}

									} else {
										gen21309 = False
									}

									if True == gen21309 {
										gen21316 = True
									} else {
										gen21316 = False
									}

								} else {
									gen21316 = False
								}

								if True == gen21316 {
									gen21323 = True
								} else {
									gen21323 = False
								}

							} else {
								gen21323 = False
							}

							if True == gen21323 {
								gen21328 = True
							} else {
								gen21328 = False
							}

						} else {
							gen21328 = False
						}

						if True == gen21328 {
							gen21331 = True
						} else {
							gen21331 = False
						}

					} else {
						gen21331 = False
					}

					if True == gen21331 {
						gen21258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21260 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21261 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21262 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21263 := Call(__e, gen21262, V804)

						gen21264 := Call(__e, gen21261, gen21263)

						gen21265 := Call(__e, gen21260, gen21264)

						gen21266 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1deref)

						gen21268 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21269 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21270 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21271 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21272 := Call(__e, gen21271, V804)

						gen21273 := Call(__e, gen21270, gen21272)

						gen21274 := Call(__e, gen21269, gen21273)

						gen21275 := Call(__e, gen21268, gen21274)

						gen21276 := Call(__e, gen21267, gen21275, V803)

						gen21277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

						gen21282 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21283 := Call(__e, gen21282, V804)

						gen21284 := Call(__e, gen21281, V803, gen21283)

						gen21285 := Call(__e, gen21280, gen21284, Nil)

						gen21286 := Call(__e, gen21279, symfreeze, gen21285)

						gen21287 := Call(__e, gen21278, gen21286, Nil)

						gen21288 := Call(__e, gen21277, V803, gen21287)

						gen21289 := Call(__e, gen21266, gen21276, gen21288)

						gen21290 := Call(__e, gen21259, gen21265, gen21289)

						__e.TailApply(gen21258, symbind, gen21290)

						return

					} else {
						gen21255 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21256 := Call(__e, gen21255, V804)

						var gen21257 Obj

						if True == gen21256 {
							gen21250 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21251 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen21252 := Call(__e, gen21251, V804)

							gen21253 := Call(__e, gen21250, gen21252)

							var gen21254 Obj

							if True == gen21253 {
								gen21243 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen21244 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21245 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21246 := Call(__e, gen21245, V804)

								gen21247 := Call(__e, gen21244, gen21246)

								gen21248 := Call(__e, gen21243, symreceive, gen21247)

								var gen21249 Obj

								if True == gen21248 {
									gen21236 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen21237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen21238 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21239 := Call(__e, gen21238, V804)

									gen21240 := Call(__e, gen21237, gen21239)

									gen21241 := Call(__e, gen21236, gen21240)

									var gen21242 Obj

									if True == gen21241 {
										gen21228 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen21229 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21230 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21232 := Call(__e, gen21231, V804)

										gen21233 := Call(__e, gen21230, gen21232)

										gen21234 := Call(__e, gen21229, gen21233)

										gen21235 := Call(__e, gen21228, Nil, gen21234)

										if True == gen21235 {
											gen21242 = True
										} else {
											gen21242 = False
										}

									} else {
										gen21242 = False
									}

									if True == gen21242 {
										gen21249 = True
									} else {
										gen21249 = False
									}

								} else {
									gen21249 = False
								}

								if True == gen21249 {
									gen21254 = True
								} else {
									gen21254 = False
								}

							} else {
								gen21254 = False
							}

							if True == gen21254 {
								gen21257 = True
							} else {
								gen21257 = False
							}

						} else {
							gen21257 = False
						}

						if True == gen21257 {
							gen21225 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

							gen21226 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21227 := Call(__e, gen21226, V804)

							__e.TailApply(gen21225, V803, gen21227)

							return

						} else {
							gen21222 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21223 := Call(__e, gen21222, V804)

							var gen21224 Obj

							if True == gen21223 {
								gen21217 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen21218 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21219 := Call(__e, gen21218, V804)

								gen21220 := Call(__e, gen21217, gen21219)

								var gen21221 Obj

								if True == gen21220 {
									gen21210 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen21211 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21212 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21213 := Call(__e, gen21212, V804)

									gen21214 := Call(__e, gen21211, gen21213)

									gen21215 := Call(__e, gen21210, symbind, gen21214)

									var gen21216 Obj

									if True == gen21215 {
										gen21203 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										gen21204 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21206 := Call(__e, gen21205, V804)

										gen21207 := Call(__e, gen21204, gen21206)

										gen21208 := Call(__e, gen21203, gen21207)

										var gen21209 Obj

										if True == gen21208 {
											gen21194 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen21195 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen21196 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen21197 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen21198 := Call(__e, gen21197, V804)

											gen21199 := Call(__e, gen21196, gen21198)

											gen21200 := Call(__e, gen21195, gen21199)

											gen21201 := Call(__e, gen21194, gen21200)

											var gen21202 Obj

											if True == gen21201 {
												gen21184 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen21185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen21186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen21187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen21188 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen21189 := Call(__e, gen21188, V804)

												gen21190 := Call(__e, gen21187, gen21189)

												gen21191 := Call(__e, gen21186, gen21190)

												gen21192 := Call(__e, gen21185, gen21191)

												gen21193 := Call(__e, gen21184, Nil, gen21192)

												if True == gen21193 {
													gen21202 = True
												} else {
													gen21202 = False
												}

											} else {
												gen21202 = False
											}

											if True == gen21202 {
												gen21209 = True
											} else {
												gen21209 = False
											}

										} else {
											gen21209 = False
										}

										if True == gen21209 {
											gen21216 = True
										} else {
											gen21216 = False
										}

									} else {
										gen21216 = False
									}

									if True == gen21216 {
										gen21221 = True
									} else {
										gen21221 = False
									}

								} else {
									gen21221 = False
								}

								if True == gen21221 {
									gen21224 = True
								} else {
									gen21224 = False
								}

							} else {
								gen21224 = False
							}

							if True == gen21224 {
								gen21151 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21152 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21153 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21154 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21155 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21156 := Call(__e, gen21155, V804)

								gen21157 := Call(__e, gen21154, gen21156)

								gen21158 := Call(__e, gen21153, gen21157)

								gen21159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21160 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

								gen21161 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21162 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21163 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21164 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21165 := Call(__e, gen21164, V804)

								gen21166 := Call(__e, gen21163, gen21165)

								gen21167 := Call(__e, gen21162, gen21166)

								gen21168 := Call(__e, gen21161, gen21167)

								gen21169 := Call(__e, gen21160, gen21168, V803)

								gen21170 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen21174 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

								gen21175 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21176 := Call(__e, gen21175, V804)

								gen21177 := Call(__e, gen21174, V803, gen21176)

								gen21178 := Call(__e, gen21173, gen21177, Nil)

								gen21179 := Call(__e, gen21172, symfreeze, gen21178)

								gen21180 := Call(__e, gen21171, gen21179, Nil)

								gen21181 := Call(__e, gen21170, V803, gen21180)

								gen21182 := Call(__e, gen21159, gen21169, gen21181)

								gen21183 := Call(__e, gen21152, gen21158, gen21182)

								__e.TailApply(gen21151, symbind, gen21183)

								return

							} else {
								gen21148 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen21149 := Call(__e, gen21148, V804)

								var gen21150 Obj

								if True == gen21149 {
									gen21143 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen21144 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21145 := Call(__e, gen21144, V804)

									gen21146 := Call(__e, gen21143, gen21145)

									var gen21147 Obj

									if True == gen21146 {
										gen21136 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										gen21137 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21138 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21139 := Call(__e, gen21138, V804)

										gen21140 := Call(__e, gen21137, gen21139)

										gen21141 := Call(__e, gen21136, symfwhen, gen21140)

										var gen21142 Obj

										if True == gen21141 {
											gen21129 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											gen21130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											gen21131 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											gen21132 := Call(__e, gen21131, V804)

											gen21133 := Call(__e, gen21130, gen21132)

											gen21134 := Call(__e, gen21129, gen21133)

											var gen21135 Obj

											if True == gen21134 {
												gen21121 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												gen21122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen21123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												gen21124 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen21125 := Call(__e, gen21124, V804)

												gen21126 := Call(__e, gen21123, gen21125)

												gen21127 := Call(__e, gen21122, gen21126)

												gen21128 := Call(__e, gen21121, Nil, gen21127)

												if True == gen21128 {
													gen21135 = True
												} else {
													gen21135 = False
												}

											} else {
												gen21135 = False
											}

											if True == gen21135 {
												gen21142 = True
											} else {
												gen21142 = False
											}

										} else {
											gen21142 = False
										}

										if True == gen21142 {
											gen21147 = True
										} else {
											gen21147 = False
										}

									} else {
										gen21147 = False
									}

									if True == gen21147 {
										gen21150 = True
									} else {
										gen21150 = False
									}

								} else {
									gen21150 = False
								}

								if True == gen21150 {
									gen21098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1lazyderef)

									gen21101 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen21103 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen21104 := Call(__e, gen21103, V804)

									gen21105 := Call(__e, gen21102, gen21104)

									gen21106 := Call(__e, gen21101, gen21105)

									gen21107 := Call(__e, gen21100, gen21106, V803)

									gen21108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21111 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									gen21112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

									gen21113 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen21114 := Call(__e, gen21113, V804)

									gen21115 := Call(__e, gen21112, V803, gen21114)

									gen21116 := Call(__e, gen21111, gen21115, Nil)

									gen21117 := Call(__e, gen21110, symfreeze, gen21116)

									gen21118 := Call(__e, gen21109, gen21117, Nil)

									gen21119 := Call(__e, gen21108, V803, gen21118)

									gen21120 := Call(__e, gen21099, gen21107, gen21119)

									__e.TailApply(gen21098, symfwhen, gen21120)

									return

								} else {
									gen21096 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen21097 := Call(__e, gen21096, V804)

									if True == gen21097 {
										gen21081 := Call(__e, PrimNS1Value(symns2_1value), symappend)

										gen21082 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen21083 := Call(__e, gen21082, V804)

										gen21084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen21085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen21086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen21087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen21088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bld_1prolog_1call)

										gen21089 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen21090 := Call(__e, gen21089, V804)

										gen21091 := Call(__e, gen21088, V803, gen21090)

										gen21092 := Call(__e, gen21087, gen21091, Nil)

										gen21093 := Call(__e, gen21086, symfreeze, gen21092)

										gen21094 := Call(__e, gen21085, gen21093, Nil)

										gen21095 := Call(__e, gen21084, V803, gen21094)

										__e.TailApply(gen21081, gen21083, gen21095)

										return

									} else {
										if True == True {
											gen21080 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

											__e.TailApply(gen21080, MakeString("implementation error in bld-prolog-call"))

											return

										} else {
											gen21079 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

											__e.TailApply(gen21079, MakeString("no cond match"))

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

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4bld_1prolog_1call, gen21409)

	gen21434 := MakeNative(func(__e *ControlFlow) {
		V806 := __e.Get(1)
		_ = V806
		gen21431 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21432 := Call(__e, gen21431, V806)

		var gen21433 Obj

		if True == gen21432 {
			gen21426 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21427 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21428 := Call(__e, gen21427, V806)

			gen21429 := Call(__e, gen21426, symdefprolog, gen21428)

			var gen21430 Obj

			if True == gen21429 {
				gen21422 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21423 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21424 := Call(__e, gen21423, V806)

				gen21425 := Call(__e, gen21422, gen21424)

				if True == gen21425 {
					gen21430 = True
				} else {
					gen21430 = False
				}

			} else {
				gen21430 = False
			}

			if True == gen21430 {
				gen21433 = True
			} else {
				gen21433 = False
			}

		} else {
			gen21433 = False
		}

		if True == gen21433 {
			gen21411 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			gen21413 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				gen21412 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5defprolog_6)

				__e.TailApply(gen21412, Y)

				return

			}, 1)

			gen21414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21415 := Call(__e, gen21414, V806)

			gen21421 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				gen21416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1error)

				gen21417 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21418 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21419 := Call(__e, gen21418, V806)

				gen21420 := Call(__e, gen21417, gen21419)

				__e.TailApply(gen21416, gen21420, Y)

				return

			}, 1)

			__e.TailApply(gen21411, gen21413, gen21415, gen21421)

			return

		} else {
			if True == True {
				__e.Return(V806)
				return
			} else {
				gen21410 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21410, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4defprolog_1macro, gen21434)

	gen21487 := MakeNative(func(__e *ControlFlow) {
		V808 := __e.Get(1)
		_ = V808
		gen21484 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21485 := Call(__e, gen21484, V808)

		var gen21486 Obj

		if True == gen21485 {
			gen21479 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21480 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21481 := Call(__e, gen21480, V808)

			gen21482 := Call(__e, gen21479, symdatatype, gen21481)

			var gen21483 Obj

			if True == gen21482 {
				gen21475 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21477 := Call(__e, gen21476, V808)

				gen21478 := Call(__e, gen21475, gen21477)

				if True == gen21478 {
					gen21483 = True
				} else {
					gen21483 = False
				}

			} else {
				gen21483 = False
			}

			if True == gen21483 {
				gen21486 = True
			} else {
				gen21486 = False
			}

		} else {
			gen21486 = False
		}

		if True == gen21486 {
			gen21436 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21437 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			gen21439 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21440 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21441 := Call(__e, gen21440, V808)

			gen21442 := Call(__e, gen21439, gen21441)

			gen21443 := Call(__e, gen21438, gen21442)

			gen21444 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21445 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21446 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21447 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21448 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21449 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21450 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21451 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21452 := Call(__e, gen21451, symX, Nil)

			gen21453 := Call(__e, gen21450, symshen_4_5datatype_1rules_6, gen21452)

			gen21454 := Call(__e, gen21449, gen21453, Nil)

			gen21455 := Call(__e, gen21448, symX, gen21454)

			gen21456 := Call(__e, gen21447, symlambda, gen21455)

			gen21457 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			gen21459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21460 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21461 := Call(__e, gen21460, V808)

			gen21462 := Call(__e, gen21459, gen21461)

			gen21463 := Call(__e, gen21458, gen21462)

			gen21464 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21465 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21466 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21467 := Call(__e, gen21466, symshen_4datatype_1error, Nil)

			gen21468 := Call(__e, gen21465, symfunction, gen21467)

			gen21469 := Call(__e, gen21464, gen21468, Nil)

			gen21470 := Call(__e, gen21457, gen21463, gen21469)

			gen21471 := Call(__e, gen21446, gen21456, gen21470)

			gen21472 := Call(__e, gen21445, symcompile, gen21471)

			gen21473 := Call(__e, gen21444, gen21472, Nil)

			gen21474 := Call(__e, gen21437, gen21443, gen21473)

			__e.TailApply(gen21436, symshen_4process_1datatype, gen21474)

			return

		} else {
			if True == True {
				__e.Return(V808)
				return
			} else {
				gen21435 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21435, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1macro, gen21487)

	gen21493 := MakeNative(func(__e *ControlFlow) {
		V810 := __e.Get(1)
		_ = V810
		gen21488 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		gen21489 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen21490 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		gen21491 := Call(__e, gen21490, V810)

		gen21492 := Call(__e, gen21489, gen21491, MakeString("#type"))

		__e.TailApply(gen21488, gen21492)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4intern_1type, gen21493)

	gen21595 := MakeNative(func(__e *ControlFlow) {
		V812 := __e.Get(1)
		_ = V812
		gen21592 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21593 := Call(__e, gen21592, V812)

		var gen21594 Obj

		if True == gen21593 {
			gen21587 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21588 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21589 := Call(__e, gen21588, V812)

			gen21590 := Call(__e, gen21587, sym_8s, gen21589)

			var gen21591 Obj

			if True == gen21590 {
				gen21582 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21584 := Call(__e, gen21583, V812)

				gen21585 := Call(__e, gen21582, gen21584)

				var gen21586 Obj

				if True == gen21585 {
					gen21575 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21577 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21578 := Call(__e, gen21577, V812)

					gen21579 := Call(__e, gen21576, gen21578)

					gen21580 := Call(__e, gen21575, gen21579)

					var gen21581 Obj

					if True == gen21580 {
						gen21567 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21570 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21571 := Call(__e, gen21570, V812)

						gen21572 := Call(__e, gen21569, gen21571)

						gen21573 := Call(__e, gen21568, gen21572)

						gen21574 := Call(__e, gen21567, gen21573)

						if True == gen21574 {
							gen21581 = True
						} else {
							gen21581 = False
						}

					} else {
						gen21581 = False
					}

					if True == gen21581 {
						gen21586 = True
					} else {
						gen21586 = False
					}

				} else {
					gen21586 = False
				}

				if True == gen21586 {
					gen21591 = True
				} else {
					gen21591 = False
				}

			} else {
				gen21591 = False
			}

			if True == gen21591 {
				gen21594 = True
			} else {
				gen21594 = False
			}

		} else {
			gen21594 = False
		}

		if True == gen21594 {
			gen21550 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21551 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21552 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21553 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21554 := Call(__e, gen21553, V812)

			gen21555 := Call(__e, gen21552, gen21554)

			gen21556 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8s_1macro)

			gen21558 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21559 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21560 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21561 := Call(__e, gen21560, V812)

			gen21562 := Call(__e, gen21559, gen21561)

			gen21563 := Call(__e, gen21558, sym_8s, gen21562)

			gen21564 := Call(__e, gen21557, gen21563)

			gen21565 := Call(__e, gen21556, gen21564, Nil)

			gen21566 := Call(__e, gen21551, gen21555, gen21565)

			__e.TailApply(gen21550, sym_8s, gen21566)

			return

		} else {
			gen21547 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21548 := Call(__e, gen21547, V812)

			var gen21549 Obj

			if True == gen21548 {
				gen21542 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen21543 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21544 := Call(__e, gen21543, V812)

				gen21545 := Call(__e, gen21542, sym_8s, gen21544)

				var gen21546 Obj

				if True == gen21545 {
					gen21537 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21538 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21539 := Call(__e, gen21538, V812)

					gen21540 := Call(__e, gen21537, gen21539)

					var gen21541 Obj

					if True == gen21540 {
						gen21530 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21531 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21532 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21533 := Call(__e, gen21532, V812)

						gen21534 := Call(__e, gen21531, gen21533)

						gen21535 := Call(__e, gen21530, gen21534)

						var gen21536 Obj

						if True == gen21535 {
							gen21521 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen21522 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21523 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21524 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21525 := Call(__e, gen21524, V812)

							gen21526 := Call(__e, gen21523, gen21525)

							gen21527 := Call(__e, gen21522, gen21526)

							gen21528 := Call(__e, gen21521, Nil, gen21527)

							var gen21529 Obj

							if True == gen21528 {
								gen21515 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

								gen21516 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen21517 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21518 := Call(__e, gen21517, V812)

								gen21519 := Call(__e, gen21516, gen21518)

								gen21520 := Call(__e, gen21515, gen21519)

								if True == gen21520 {
									gen21529 = True
								} else {
									gen21529 = False
								}

							} else {
								gen21529 = False
							}

							if True == gen21529 {
								gen21536 = True
							} else {
								gen21536 = False
							}

						} else {
							gen21536 = False
						}

						if True == gen21536 {
							gen21541 = True
						} else {
							gen21541 = False
						}

					} else {
						gen21541 = False
					}

					if True == gen21541 {
						gen21546 = True
					} else {
						gen21546 = False
					}

				} else {
					gen21546 = False
				}

				if True == gen21546 {
					gen21549 = True
				} else {
					gen21549 = False
				}

			} else {
				gen21549 = False
			}

			if True == gen21549 {
				gen21508 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					gen21504 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					gen21505 := Call(__e, PrimNS1Value(symns2_1value), symlength)

					gen21506 := Call(__e, gen21505, E)

					gen21507 := Call(__e, gen21504, gen21506, MakeNumber(1))

					if True == gen21507 {
						gen21495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8s_1macro)

						gen21496 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21497 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen21498 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21499 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21500 := Call(__e, gen21499, V812)

						gen21501 := Call(__e, gen21498, gen21500)

						gen21502 := Call(__e, gen21497, E, gen21501)

						gen21503 := Call(__e, gen21496, sym_8s, gen21502)

						__e.TailApply(gen21495, gen21503)

						return

					} else {
						__e.Return(V812)
						return
					}

				}, 1)

				gen21509 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

				gen21510 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21511 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21512 := Call(__e, gen21511, V812)

				gen21513 := Call(__e, gen21510, gen21512)

				gen21514 := Call(__e, gen21509, gen21513)

				__e.TailApply(gen21508, gen21514)

				return

			} else {
				if True == True {
					__e.Return(V812)
					return
				} else {
					gen21494 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen21494, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_8s_1macro, gen21595)

	gen21613 := MakeNative(func(__e *ControlFlow) {
		V814 := __e.Get(1)
		_ = V814
		gen21610 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21611 := Call(__e, gen21610, V814)

		var gen21612 Obj

		if True == gen21611 {
			gen21606 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21607 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21608 := Call(__e, gen21607, V814)

			gen21609 := Call(__e, gen21606, symsynonyms, gen21608)

			if True == gen21609 {
				gen21612 = True
			} else {
				gen21612 = False
			}

		} else {
			gen21612 = False
		}

		if True == gen21612 {
			gen21597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21599 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			gen21600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1synonyms)

			gen21601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21602 := Call(__e, gen21601, V814)

			gen21603 := Call(__e, gen21600, gen21602)

			gen21604 := Call(__e, gen21599, gen21603)

			gen21605 := Call(__e, gen21598, gen21604, Nil)

			__e.TailApply(gen21597, symshen_4synonyms_1help, gen21605)

			return

		} else {
			if True == True {
				__e.Return(V814)
				return
			} else {
				gen21596 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21596, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1macro, gen21613)

	gen21617 := MakeNative(func(__e *ControlFlow) {
		V816 := __e.Get(1)
		_ = V816
		gen21614 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen21616 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen21615 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

			__e.TailApply(gen21615, X)

			return

		}, 1)

		__e.TailApply(gen21614, gen21616, V816)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4curry_1synonyms, gen21617)

	gen21634 := MakeNative(func(__e *ControlFlow) {
		V818 := __e.Get(1)
		_ = V818
		gen21631 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21632 := Call(__e, gen21631, V818)

		var gen21633 Obj

		if True == gen21632 {
			gen21626 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21627 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21628 := Call(__e, gen21627, V818)

			gen21629 := Call(__e, gen21626, symnl, gen21628)

			var gen21630 Obj

			if True == gen21629 {
				gen21622 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen21623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21624 := Call(__e, gen21623, V818)

				gen21625 := Call(__e, gen21622, Nil, gen21624)

				if True == gen21625 {
					gen21630 = True
				} else {
					gen21630 = False
				}

			} else {
				gen21630 = False
			}

			if True == gen21630 {
				gen21633 = True
			} else {
				gen21633 = False
			}

		} else {
			gen21633 = False
		}

		if True == gen21633 {
			gen21619 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21620 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21621 := Call(__e, gen21620, MakeNumber(1), Nil)

			__e.TailApply(gen21619, symnl, gen21621)

			return

		} else {
			if True == True {
				__e.Return(V818)
				return
			} else {
				gen21618 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21618, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4nl_1macro, gen21634)

	gen21701 := MakeNative(func(__e *ControlFlow) {
		V820 := __e.Get(1)
		_ = V820
		gen21698 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21699 := Call(__e, gen21698, V820)

		var gen21700 Obj

		if True == gen21699 {
			gen21693 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21694 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21695 := Call(__e, gen21694, V820)

			gen21696 := Call(__e, gen21693, gen21695)

			var gen21697 Obj

			if True == gen21696 {
				gen21686 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21689 := Call(__e, gen21688, V820)

				gen21690 := Call(__e, gen21687, gen21689)

				gen21691 := Call(__e, gen21686, gen21690)

				var gen21692 Obj

				if True == gen21691 {
					gen21677 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21678 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21679 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21681 := Call(__e, gen21680, V820)

					gen21682 := Call(__e, gen21679, gen21681)

					gen21683 := Call(__e, gen21678, gen21682)

					gen21684 := Call(__e, gen21677, gen21683)

					var gen21685 Obj

					if True == gen21684 {
						gen21657 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

						gen21658 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21659 := Call(__e, gen21658, V820)

						gen21660 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21664 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen21668 := Call(__e, gen21667, symdo, Nil)

						gen21669 := Call(__e, gen21666, sym_d, gen21668)

						gen21670 := Call(__e, gen21665, sym_7, gen21669)

						gen21671 := Call(__e, gen21664, symor, gen21670)

						gen21672 := Call(__e, gen21663, symand, gen21671)

						gen21673 := Call(__e, gen21662, symappend, gen21672)

						gen21674 := Call(__e, gen21661, sym_8v, gen21673)

						gen21675 := Call(__e, gen21660, sym_8p, gen21674)

						gen21676 := Call(__e, gen21657, gen21659, gen21675)

						if True == gen21676 {
							gen21685 = True
						} else {
							gen21685 = False
						}

					} else {
						gen21685 = False
					}

					if True == gen21685 {
						gen21692 = True
					} else {
						gen21692 = False
					}

				} else {
					gen21692 = False
				}

				if True == gen21692 {
					gen21697 = True
				} else {
					gen21697 = False
				}

			} else {
				gen21697 = False
			}

			if True == gen21697 {
				gen21700 = True
			} else {
				gen21700 = False
			}

		} else {
			gen21700 = False
		}

		if True == gen21700 {
			gen21636 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21638 := Call(__e, gen21637, V820)

			gen21639 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21640 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21642 := Call(__e, gen21641, V820)

			gen21643 := Call(__e, gen21640, gen21642)

			gen21644 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21645 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1macro)

			gen21646 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21647 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21648 := Call(__e, gen21647, V820)

			gen21649 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21651 := Call(__e, gen21650, V820)

			gen21652 := Call(__e, gen21649, gen21651)

			gen21653 := Call(__e, gen21646, gen21648, gen21652)

			gen21654 := Call(__e, gen21645, gen21653)

			gen21655 := Call(__e, gen21644, gen21654, Nil)

			gen21656 := Call(__e, gen21639, gen21643, gen21655)

			__e.TailApply(gen21636, gen21638, gen21656)

			return

		} else {
			if True == True {
				__e.Return(V820)
				return
			} else {
				gen21635 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21635, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1macro, gen21701)

	gen21769 := MakeNative(func(__e *ControlFlow) {
		V822 := __e.Get(1)
		_ = V822
		gen21766 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21767 := Call(__e, gen21766, V822)

		var gen21768 Obj

		if True == gen21767 {
			gen21761 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21762 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21763 := Call(__e, gen21762, V822)

			gen21764 := Call(__e, gen21761, symlet, gen21763)

			var gen21765 Obj

			if True == gen21764 {
				gen21756 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21757 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21758 := Call(__e, gen21757, V822)

				gen21759 := Call(__e, gen21756, gen21758)

				var gen21760 Obj

				if True == gen21759 {
					gen21749 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21750 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21751 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21752 := Call(__e, gen21751, V822)

					gen21753 := Call(__e, gen21750, gen21752)

					gen21754 := Call(__e, gen21749, gen21753)

					var gen21755 Obj

					if True == gen21754 {
						gen21740 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21741 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21742 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21744 := Call(__e, gen21743, V822)

						gen21745 := Call(__e, gen21742, gen21744)

						gen21746 := Call(__e, gen21741, gen21745)

						gen21747 := Call(__e, gen21740, gen21746)

						var gen21748 Obj

						if True == gen21747 {
							gen21730 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21731 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21732 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21733 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21734 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21735 := Call(__e, gen21734, V822)

							gen21736 := Call(__e, gen21733, gen21735)

							gen21737 := Call(__e, gen21732, gen21736)

							gen21738 := Call(__e, gen21731, gen21737)

							gen21739 := Call(__e, gen21730, gen21738)

							if True == gen21739 {
								gen21748 = True
							} else {
								gen21748 = False
							}

						} else {
							gen21748 = False
						}

						if True == gen21748 {
							gen21755 = True
						} else {
							gen21755 = False
						}

					} else {
						gen21755 = False
					}

					if True == gen21755 {
						gen21760 = True
					} else {
						gen21760 = False
					}

				} else {
					gen21760 = False
				}

				if True == gen21760 {
					gen21765 = True
				} else {
					gen21765 = False
				}

			} else {
				gen21765 = False
			}

			if True == gen21765 {
				gen21768 = True
			} else {
				gen21768 = False
			}

		} else {
			gen21768 = False
		}

		if True == gen21768 {
			gen21703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21705 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21706 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21707 := Call(__e, gen21706, V822)

			gen21708 := Call(__e, gen21705, gen21707)

			gen21709 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21710 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21711 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21712 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21713 := Call(__e, gen21712, V822)

			gen21714 := Call(__e, gen21711, gen21713)

			gen21715 := Call(__e, gen21710, gen21714)

			gen21716 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21717 := Call(__e, PrimNS1Value(symns2_1value), symshen_4let_1macro)

			gen21718 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21719 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21722 := Call(__e, gen21721, V822)

			gen21723 := Call(__e, gen21720, gen21722)

			gen21724 := Call(__e, gen21719, gen21723)

			gen21725 := Call(__e, gen21718, symlet, gen21724)

			gen21726 := Call(__e, gen21717, gen21725)

			gen21727 := Call(__e, gen21716, gen21726, Nil)

			gen21728 := Call(__e, gen21709, gen21715, gen21727)

			gen21729 := Call(__e, gen21704, gen21708, gen21728)

			__e.TailApply(gen21703, symlet, gen21729)

			return

		} else {
			if True == True {
				__e.Return(V822)
				return
			} else {
				gen21702 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21702, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4let_1macro, gen21769)

	gen21847 := MakeNative(func(__e *ControlFlow) {
		V824 := __e.Get(1)
		_ = V824
		gen21844 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21845 := Call(__e, gen21844, V824)

		var gen21846 Obj

		if True == gen21845 {
			gen21839 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21840 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21841 := Call(__e, gen21840, V824)

			gen21842 := Call(__e, gen21839, sym_c_4, gen21841)

			var gen21843 Obj

			if True == gen21842 {
				gen21834 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21836 := Call(__e, gen21835, V824)

				gen21837 := Call(__e, gen21834, gen21836)

				var gen21838 Obj

				if True == gen21837 {
					gen21827 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21828 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21829 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21830 := Call(__e, gen21829, V824)

					gen21831 := Call(__e, gen21828, gen21830)

					gen21832 := Call(__e, gen21827, gen21831)

					var gen21833 Obj

					if True == gen21832 {
						gen21819 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21820 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21822 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21823 := Call(__e, gen21822, V824)

						gen21824 := Call(__e, gen21821, gen21823)

						gen21825 := Call(__e, gen21820, gen21824)

						gen21826 := Call(__e, gen21819, gen21825)

						if True == gen21826 {
							gen21833 = True
						} else {
							gen21833 = False
						}

					} else {
						gen21833 = False
					}

					if True == gen21833 {
						gen21838 = True
					} else {
						gen21838 = False
					}

				} else {
					gen21838 = False
				}

				if True == gen21838 {
					gen21843 = True
				} else {
					gen21843 = False
				}

			} else {
				gen21843 = False
			}

			if True == gen21843 {
				gen21846 = True
			} else {
				gen21846 = False
			}

		} else {
			gen21846 = False
		}

		if True == gen21846 {
			gen21802 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21803 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21804 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21805 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21806 := Call(__e, gen21805, V824)

			gen21807 := Call(__e, gen21804, gen21806)

			gen21808 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abs_1macro)

			gen21810 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21811 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21812 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21813 := Call(__e, gen21812, V824)

			gen21814 := Call(__e, gen21811, gen21813)

			gen21815 := Call(__e, gen21810, sym_c_4, gen21814)

			gen21816 := Call(__e, gen21809, gen21815)

			gen21817 := Call(__e, gen21808, gen21816, Nil)

			gen21818 := Call(__e, gen21803, gen21807, gen21817)

			__e.TailApply(gen21802, symlambda, gen21818)

			return

		} else {
			gen21799 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21800 := Call(__e, gen21799, V824)

			var gen21801 Obj

			if True == gen21800 {
				gen21794 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen21795 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21796 := Call(__e, gen21795, V824)

				gen21797 := Call(__e, gen21794, sym_c_4, gen21796)

				var gen21798 Obj

				if True == gen21797 {
					gen21789 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21790 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21791 := Call(__e, gen21790, V824)

					gen21792 := Call(__e, gen21789, gen21791)

					var gen21793 Obj

					if True == gen21792 {
						gen21782 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21783 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21785 := Call(__e, gen21784, V824)

						gen21786 := Call(__e, gen21783, gen21785)

						gen21787 := Call(__e, gen21782, gen21786)

						var gen21788 Obj

						if True == gen21787 {
							gen21774 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen21775 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21777 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21778 := Call(__e, gen21777, V824)

							gen21779 := Call(__e, gen21776, gen21778)

							gen21780 := Call(__e, gen21775, gen21779)

							gen21781 := Call(__e, gen21774, Nil, gen21780)

							if True == gen21781 {
								gen21788 = True
							} else {
								gen21788 = False
							}

						} else {
							gen21788 = False
						}

						if True == gen21788 {
							gen21793 = True
						} else {
							gen21793 = False
						}

					} else {
						gen21793 = False
					}

					if True == gen21793 {
						gen21798 = True
					} else {
						gen21798 = False
					}

				} else {
					gen21798 = False
				}

				if True == gen21798 {
					gen21801 = True
				} else {
					gen21801 = False
				}

			} else {
				gen21801 = False
			}

			if True == gen21801 {
				gen21771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21772 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21773 := Call(__e, gen21772, V824)

				__e.TailApply(gen21771, symlambda, gen21773)

				return

			} else {
				if True == True {
					__e.Return(V824)
					return
				} else {
					gen21770 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen21770, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4abs_1macro, gen21847)

	gen21995 := MakeNative(func(__e *ControlFlow) {
		V828 := __e.Get(1)
		_ = V828
		gen21992 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen21993 := Call(__e, gen21992, V828)

		var gen21994 Obj

		if True == gen21993 {
			gen21987 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen21988 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21989 := Call(__e, gen21988, V828)

			gen21990 := Call(__e, gen21987, symcases, gen21989)

			var gen21991 Obj

			if True == gen21990 {
				gen21982 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21983 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21984 := Call(__e, gen21983, V828)

				gen21985 := Call(__e, gen21982, gen21984)

				var gen21986 Obj

				if True == gen21985 {
					gen21975 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen21976 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21977 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21978 := Call(__e, gen21977, V828)

					gen21979 := Call(__e, gen21976, gen21978)

					gen21980 := Call(__e, gen21975, True, gen21979)

					var gen21981 Obj

					if True == gen21980 {
						gen21969 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21970 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21972 := Call(__e, gen21971, V828)

						gen21973 := Call(__e, gen21970, gen21972)

						gen21974 := Call(__e, gen21969, gen21973)

						if True == gen21974 {
							gen21981 = True
						} else {
							gen21981 = False
						}

					} else {
						gen21981 = False
					}

					if True == gen21981 {
						gen21986 = True
					} else {
						gen21986 = False
					}

				} else {
					gen21986 = False
				}

				if True == gen21986 {
					gen21991 = True
				} else {
					gen21991 = False
				}

			} else {
				gen21991 = False
			}

			if True == gen21991 {
				gen21994 = True
			} else {
				gen21994 = False
			}

		} else {
			gen21994 = False
		}

		if True == gen21994 {
			gen21964 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen21965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21966 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen21967 := Call(__e, gen21966, V828)

			gen21968 := Call(__e, gen21965, gen21967)

			__e.TailApply(gen21964, gen21968)

			return

		} else {
			gen21961 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen21962 := Call(__e, gen21961, V828)

			var gen21963 Obj

			if True == gen21962 {
				gen21956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen21957 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21958 := Call(__e, gen21957, V828)

				gen21959 := Call(__e, gen21956, symcases, gen21958)

				var gen21960 Obj

				if True == gen21959 {
					gen21951 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21953 := Call(__e, gen21952, V828)

					gen21954 := Call(__e, gen21951, gen21953)

					var gen21955 Obj

					if True == gen21954 {
						gen21944 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21945 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21946 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21947 := Call(__e, gen21946, V828)

						gen21948 := Call(__e, gen21945, gen21947)

						gen21949 := Call(__e, gen21944, gen21948)

						var gen21950 Obj

						if True == gen21949 {
							gen21936 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen21937 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21938 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21940 := Call(__e, gen21939, V828)

							gen21941 := Call(__e, gen21938, gen21940)

							gen21942 := Call(__e, gen21937, gen21941)

							gen21943 := Call(__e, gen21936, Nil, gen21942)

							if True == gen21943 {
								gen21950 = True
							} else {
								gen21950 = False
							}

						} else {
							gen21950 = False
						}

						if True == gen21950 {
							gen21955 = True
						} else {
							gen21955 = False
						}

					} else {
						gen21955 = False
					}

					if True == gen21955 {
						gen21960 = True
					} else {
						gen21960 = False
					}

				} else {
					gen21960 = False
				}

				if True == gen21960 {
					gen21963 = True
				} else {
					gen21963 = False
				}

			} else {
				gen21963 = False
			}

			if True == gen21963 {
				gen21915 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21916 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21917 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21918 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21919 := Call(__e, gen21918, V828)

				gen21920 := Call(__e, gen21917, gen21919)

				gen21921 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21922 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen21923 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21924 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen21925 := Call(__e, gen21924, V828)

				gen21926 := Call(__e, gen21923, gen21925)

				gen21927 := Call(__e, gen21922, gen21926)

				gen21928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen21931 := Call(__e, gen21930, MakeString("error: cases exhausted"), Nil)

				gen21932 := Call(__e, gen21929, symsimple_1error, gen21931)

				gen21933 := Call(__e, gen21928, gen21932, Nil)

				gen21934 := Call(__e, gen21921, gen21927, gen21933)

				gen21935 := Call(__e, gen21916, gen21920, gen21934)

				__e.TailApply(gen21915, symif, gen21935)

				return

			} else {
				gen21912 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen21913 := Call(__e, gen21912, V828)

				var gen21914 Obj

				if True == gen21913 {
					gen21907 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen21908 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21909 := Call(__e, gen21908, V828)

					gen21910 := Call(__e, gen21907, symcases, gen21909)

					var gen21911 Obj

					if True == gen21910 {
						gen21902 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen21903 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen21904 := Call(__e, gen21903, V828)

						gen21905 := Call(__e, gen21902, gen21904)

						var gen21906 Obj

						if True == gen21905 {
							gen21896 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21898 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21899 := Call(__e, gen21898, V828)

							gen21900 := Call(__e, gen21897, gen21899)

							gen21901 := Call(__e, gen21896, gen21900)

							if True == gen21901 {
								gen21906 = True
							} else {
								gen21906 = False
							}

						} else {
							gen21906 = False
						}

						if True == gen21906 {
							gen21911 = True
						} else {
							gen21911 = False
						}

					} else {
						gen21911 = False
					}

					if True == gen21911 {
						gen21914 = True
					} else {
						gen21914 = False
					}

				} else {
					gen21914 = False
				}

				if True == gen21914 {
					gen21869 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21870 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21871 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21873 := Call(__e, gen21872, V828)

					gen21874 := Call(__e, gen21871, gen21873)

					gen21875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen21877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21879 := Call(__e, gen21878, V828)

					gen21880 := Call(__e, gen21877, gen21879)

					gen21881 := Call(__e, gen21876, gen21880)

					gen21882 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cases_1macro)

					gen21884 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen21885 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21886 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21887 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen21888 := Call(__e, gen21887, V828)

					gen21889 := Call(__e, gen21886, gen21888)

					gen21890 := Call(__e, gen21885, gen21889)

					gen21891 := Call(__e, gen21884, symcases, gen21890)

					gen21892 := Call(__e, gen21883, gen21891)

					gen21893 := Call(__e, gen21882, gen21892, Nil)

					gen21894 := Call(__e, gen21875, gen21881, gen21893)

					gen21895 := Call(__e, gen21870, gen21874, gen21894)

					__e.TailApply(gen21869, symif, gen21895)

					return

				} else {
					gen21866 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen21867 := Call(__e, gen21866, V828)

					var gen21868 Obj

					if True == gen21867 {
						gen21861 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen21862 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen21863 := Call(__e, gen21862, V828)

						gen21864 := Call(__e, gen21861, symcases, gen21863)

						var gen21865 Obj

						if True == gen21864 {
							gen21856 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen21857 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen21858 := Call(__e, gen21857, V828)

							gen21859 := Call(__e, gen21856, gen21858)

							var gen21860 Obj

							if True == gen21859 {
								gen21850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen21851 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen21853 := Call(__e, gen21852, V828)

								gen21854 := Call(__e, gen21851, gen21853)

								gen21855 := Call(__e, gen21850, Nil, gen21854)

								if True == gen21855 {
									gen21860 = True
								} else {
									gen21860 = False
								}

							} else {
								gen21860 = False
							}

							if True == gen21860 {
								gen21865 = True
							} else {
								gen21865 = False
							}

						} else {
							gen21865 = False
						}

						if True == gen21865 {
							gen21868 = True
						} else {
							gen21868 = False
						}

					} else {
						gen21868 = False
					}

					if True == gen21868 {
						gen21849 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen21849, MakeString("error: odd number of case elements\n"))

						return

					} else {
						if True == True {
							__e.Return(V828)
							return
						} else {
							gen21848 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen21848, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4cases_1macro, gen21995)

	gen22083 := MakeNative(func(__e *ControlFlow) {
		V830 := __e.Get(1)
		_ = V830
		gen22080 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22081 := Call(__e, gen22080, V830)

		var gen22082 Obj

		if True == gen22081 {
			gen22075 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen22076 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22077 := Call(__e, gen22076, V830)

			gen22078 := Call(__e, gen22075, symtime, gen22077)

			var gen22079 Obj

			if True == gen22078 {
				gen22070 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22072 := Call(__e, gen22071, V830)

				gen22073 := Call(__e, gen22070, gen22072)

				var gen22074 Obj

				if True == gen22073 {
					gen22064 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen22065 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22066 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22067 := Call(__e, gen22066, V830)

					gen22068 := Call(__e, gen22065, gen22067)

					gen22069 := Call(__e, gen22064, Nil, gen22068)

					if True == gen22069 {
						gen22074 = True
					} else {
						gen22074 = False
					}

				} else {
					gen22074 = False
				}

				if True == gen22074 {
					gen22079 = True
				} else {
					gen22079 = False
				}

			} else {
				gen22079 = False
			}

			if True == gen22079 {
				gen22082 = True
			} else {
				gen22082 = False
			}

		} else {
			gen22082 = False
		}

		if True == gen22082 {
			gen21997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4let_1macro)

			gen21998 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen21999 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22000 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22001 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22003 := Call(__e, gen22002, symrun, Nil)

			gen22004 := Call(__e, gen22001, symget_1time, gen22003)

			gen22005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22007 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22008 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22009 := Call(__e, gen22008, V830)

			gen22010 := Call(__e, gen22007, gen22009)

			gen22011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22015 := Call(__e, gen22014, symrun, Nil)

			gen22016 := Call(__e, gen22013, symget_1time, gen22015)

			gen22017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22021 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22022 := Call(__e, gen22021, symStart, Nil)

			gen22023 := Call(__e, gen22020, symFinish, gen22022)

			gen22024 := Call(__e, gen22019, sym_1, gen22023)

			gen22025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22028 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22029 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22030 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22031 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22032 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22033 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22034 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22035 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22036 := Call(__e, gen22035, symTime, Nil)

			gen22037 := Call(__e, gen22034, symstr, gen22036)

			gen22038 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22039 := Call(__e, gen22038, MakeString(" secs\n"), Nil)

			gen22040 := Call(__e, gen22033, gen22037, gen22039)

			gen22041 := Call(__e, gen22032, symcn, gen22040)

			gen22042 := Call(__e, gen22031, gen22041, Nil)

			gen22043 := Call(__e, gen22030, MakeString("\nrun time: "), gen22042)

			gen22044 := Call(__e, gen22029, symcn, gen22043)

			gen22045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22047 := Call(__e, gen22046, symstoutput, Nil)

			gen22048 := Call(__e, gen22045, gen22047, Nil)

			gen22049 := Call(__e, gen22028, gen22044, gen22048)

			gen22050 := Call(__e, gen22027, symshen_4prhush, gen22049)

			gen22051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22052 := Call(__e, gen22051, symResult, Nil)

			gen22053 := Call(__e, gen22026, gen22050, gen22052)

			gen22054 := Call(__e, gen22025, symMessage, gen22053)

			gen22055 := Call(__e, gen22018, gen22024, gen22054)

			gen22056 := Call(__e, gen22017, symTime, gen22055)

			gen22057 := Call(__e, gen22012, gen22016, gen22056)

			gen22058 := Call(__e, gen22011, symFinish, gen22057)

			gen22059 := Call(__e, gen22006, gen22010, gen22058)

			gen22060 := Call(__e, gen22005, symResult, gen22059)

			gen22061 := Call(__e, gen22000, gen22004, gen22060)

			gen22062 := Call(__e, gen21999, symStart, gen22061)

			gen22063 := Call(__e, gen21998, symlet, gen22062)

			__e.TailApply(gen21997, gen22063)

			return

		} else {
			if True == True {
				__e.Return(V830)
				return
			} else {
				gen21996 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen21996, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4timer_1macro, gen22083)

	gen22098 := MakeNative(func(__e *ControlFlow) {
		V832 := __e.Get(1)
		_ = V832
		gen22096 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22097 := Call(__e, gen22096, V832)

		if True == gen22097 {
			gen22085 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22086 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22087 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22088 := Call(__e, gen22087, V832)

			gen22089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22090 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tuple_1up)

			gen22091 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22092 := Call(__e, gen22091, V832)

			gen22093 := Call(__e, gen22090, gen22092)

			gen22094 := Call(__e, gen22089, gen22093, Nil)

			gen22095 := Call(__e, gen22086, gen22088, gen22094)

			__e.TailApply(gen22085, sym_8p, gen22095)

			return

		} else {
			if True == True {
				__e.Return(V832)
				return
			} else {
				gen22084 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22084, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4tuple_1up, gen22098)

	gen22268 := MakeNative(func(__e *ControlFlow) {
		V834 := __e.Get(1)
		_ = V834
		gen22265 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22266 := Call(__e, gen22265, V834)

		var gen22267 Obj

		if True == gen22266 {
			gen22260 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen22261 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22262 := Call(__e, gen22261, V834)

			gen22263 := Call(__e, gen22260, symput, gen22262)

			var gen22264 Obj

			if True == gen22263 {
				gen22255 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22256 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22257 := Call(__e, gen22256, V834)

				gen22258 := Call(__e, gen22255, gen22257)

				var gen22259 Obj

				if True == gen22258 {
					gen22248 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen22249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22251 := Call(__e, gen22250, V834)

					gen22252 := Call(__e, gen22249, gen22251)

					gen22253 := Call(__e, gen22248, gen22252)

					var gen22254 Obj

					if True == gen22253 {
						gen22239 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen22240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22241 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22242 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22243 := Call(__e, gen22242, V834)

						gen22244 := Call(__e, gen22241, gen22243)

						gen22245 := Call(__e, gen22240, gen22244)

						gen22246 := Call(__e, gen22239, gen22245)

						var gen22247 Obj

						if True == gen22246 {
							gen22229 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen22230 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22231 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22232 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22233 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22234 := Call(__e, gen22233, V834)

							gen22235 := Call(__e, gen22232, gen22234)

							gen22236 := Call(__e, gen22231, gen22235)

							gen22237 := Call(__e, gen22230, gen22236)

							gen22238 := Call(__e, gen22229, Nil, gen22237)

							if True == gen22238 {
								gen22247 = True
							} else {
								gen22247 = False
							}

						} else {
							gen22247 = False
						}

						if True == gen22247 {
							gen22254 = True
						} else {
							gen22254 = False
						}

					} else {
						gen22254 = False
					}

					if True == gen22254 {
						gen22259 = True
					} else {
						gen22259 = False
					}

				} else {
					gen22259 = False
				}

				if True == gen22259 {
					gen22264 = True
				} else {
					gen22264 = False
				}

			} else {
				gen22264 = False
			}

			if True == gen22264 {
				gen22267 = True
			} else {
				gen22267 = False
			}

		} else {
			gen22267 = False
		}

		if True == gen22267 {
			gen22198 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22199 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22200 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22201 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22202 := Call(__e, gen22201, V834)

			gen22203 := Call(__e, gen22200, gen22202)

			gen22204 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22206 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22208 := Call(__e, gen22207, V834)

			gen22209 := Call(__e, gen22206, gen22208)

			gen22210 := Call(__e, gen22205, gen22209)

			gen22211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22212 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22213 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22214 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22215 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22216 := Call(__e, gen22215, V834)

			gen22217 := Call(__e, gen22214, gen22216)

			gen22218 := Call(__e, gen22213, gen22217)

			gen22219 := Call(__e, gen22212, gen22218)

			gen22220 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22221 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22222 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen22223 := Call(__e, gen22222, sym_dproperty_1vector_d, Nil)

			gen22224 := Call(__e, gen22221, symvalue, gen22223)

			gen22225 := Call(__e, gen22220, gen22224, Nil)

			gen22226 := Call(__e, gen22211, gen22219, gen22225)

			gen22227 := Call(__e, gen22204, gen22210, gen22226)

			gen22228 := Call(__e, gen22199, gen22203, gen22227)

			__e.TailApply(gen22198, symput, gen22228)

			return

		} else {
			gen22195 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen22196 := Call(__e, gen22195, V834)

			var gen22197 Obj

			if True == gen22196 {
				gen22190 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen22191 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22192 := Call(__e, gen22191, V834)

				gen22193 := Call(__e, gen22190, symget, gen22192)

				var gen22194 Obj

				if True == gen22193 {
					gen22185 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen22186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22187 := Call(__e, gen22186, V834)

					gen22188 := Call(__e, gen22185, gen22187)

					var gen22189 Obj

					if True == gen22188 {
						gen22178 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen22179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22180 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22181 := Call(__e, gen22180, V834)

						gen22182 := Call(__e, gen22179, gen22181)

						gen22183 := Call(__e, gen22178, gen22182)

						var gen22184 Obj

						if True == gen22183 {
							gen22170 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen22171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22172 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22173 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22174 := Call(__e, gen22173, V834)

							gen22175 := Call(__e, gen22172, gen22174)

							gen22176 := Call(__e, gen22171, gen22175)

							gen22177 := Call(__e, gen22170, Nil, gen22176)

							if True == gen22177 {
								gen22184 = True
							} else {
								gen22184 = False
							}

						} else {
							gen22184 = False
						}

						if True == gen22184 {
							gen22189 = True
						} else {
							gen22189 = False
						}

					} else {
						gen22189 = False
					}

					if True == gen22189 {
						gen22194 = True
					} else {
						gen22194 = False
					}

				} else {
					gen22194 = False
				}

				if True == gen22194 {
					gen22197 = True
				} else {
					gen22197 = False
				}

			} else {
				gen22197 = False
			}

			if True == gen22197 {
				gen22149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22151 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22152 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22153 := Call(__e, gen22152, V834)

				gen22154 := Call(__e, gen22151, gen22153)

				gen22155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22156 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22157 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22158 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22159 := Call(__e, gen22158, V834)

				gen22160 := Call(__e, gen22157, gen22159)

				gen22161 := Call(__e, gen22156, gen22160)

				gen22162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22163 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22164 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22165 := Call(__e, gen22164, sym_dproperty_1vector_d, Nil)

				gen22166 := Call(__e, gen22163, symvalue, gen22165)

				gen22167 := Call(__e, gen22162, gen22166, Nil)

				gen22168 := Call(__e, gen22155, gen22161, gen22167)

				gen22169 := Call(__e, gen22150, gen22154, gen22168)

				__e.TailApply(gen22149, symget, gen22169)

				return

			} else {
				gen22146 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22147 := Call(__e, gen22146, V834)

				var gen22148 Obj

				if True == gen22147 {
					gen22141 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen22142 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen22143 := Call(__e, gen22142, V834)

					gen22144 := Call(__e, gen22141, symunput, gen22143)

					var gen22145 Obj

					if True == gen22144 {
						gen22136 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen22137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen22138 := Call(__e, gen22137, V834)

						gen22139 := Call(__e, gen22136, gen22138)

						var gen22140 Obj

						if True == gen22139 {
							gen22129 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen22130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen22132 := Call(__e, gen22131, V834)

							gen22133 := Call(__e, gen22130, gen22132)

							gen22134 := Call(__e, gen22129, gen22133)

							var gen22135 Obj

							if True == gen22134 {
								gen22121 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen22122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen22123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen22124 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen22125 := Call(__e, gen22124, V834)

								gen22126 := Call(__e, gen22123, gen22125)

								gen22127 := Call(__e, gen22122, gen22126)

								gen22128 := Call(__e, gen22121, Nil, gen22127)

								if True == gen22128 {
									gen22135 = True
								} else {
									gen22135 = False
								}

							} else {
								gen22135 = False
							}

							if True == gen22135 {
								gen22140 = True
							} else {
								gen22140 = False
							}

						} else {
							gen22140 = False
						}

						if True == gen22140 {
							gen22145 = True
						} else {
							gen22145 = False
						}

					} else {
						gen22145 = False
					}

					if True == gen22145 {
						gen22148 = True
					} else {
						gen22148 = False
					}

				} else {
					gen22148 = False
				}

				if True == gen22148 {
					gen22100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22101 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22102 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen22103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22104 := Call(__e, gen22103, V834)

					gen22105 := Call(__e, gen22102, gen22104)

					gen22106 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22107 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen22108 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22109 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22110 := Call(__e, gen22109, V834)

					gen22111 := Call(__e, gen22108, gen22110)

					gen22112 := Call(__e, gen22107, gen22111)

					gen22113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22116 := Call(__e, gen22115, sym_dproperty_1vector_d, Nil)

					gen22117 := Call(__e, gen22114, symvalue, gen22116)

					gen22118 := Call(__e, gen22113, gen22117, Nil)

					gen22119 := Call(__e, gen22106, gen22112, gen22118)

					gen22120 := Call(__e, gen22101, gen22105, gen22119)

					__e.TailApply(gen22100, symunput, gen22120)

					return

				} else {
					if True == True {
						__e.Return(V834)
						return
					} else {
						gen22099 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen22099, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4put_cget_1macro, gen22268)

	gen22300 := MakeNative(func(__e *ControlFlow) {
		V836 := __e.Get(1)
		_ = V836
		gen22297 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen22298 := Call(__e, gen22297, V836)

		var gen22299 Obj

		if True == gen22298 {
			gen22292 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen22293 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22294 := Call(__e, gen22293, V836)

			gen22295 := Call(__e, gen22292, symfunction, gen22294)

			var gen22296 Obj

			if True == gen22295 {
				gen22287 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22288 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22289 := Call(__e, gen22288, V836)

				gen22290 := Call(__e, gen22287, gen22289)

				var gen22291 Obj

				if True == gen22290 {
					gen22281 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen22282 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22283 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22284 := Call(__e, gen22283, V836)

					gen22285 := Call(__e, gen22282, gen22284)

					gen22286 := Call(__e, gen22281, Nil, gen22285)

					if True == gen22286 {
						gen22291 = True
					} else {
						gen22291 = False
					}

				} else {
					gen22291 = False
				}

				if True == gen22291 {
					gen22296 = True
				} else {
					gen22296 = False
				}

			} else {
				gen22296 = False
			}

			if True == gen22296 {
				gen22299 = True
			} else {
				gen22299 = False
			}

		} else {
			gen22299 = False
		}

		if True == gen22299 {
			gen22270 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction)

			gen22271 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22272 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22273 := Call(__e, gen22272, V836)

			gen22274 := Call(__e, gen22271, gen22273)

			gen22275 := Call(__e, PrimNS1Value(symns2_1value), symarity)

			gen22276 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen22277 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen22278 := Call(__e, gen22277, V836)

			gen22279 := Call(__e, gen22276, gen22278)

			gen22280 := Call(__e, gen22275, gen22279)

			__e.TailApply(gen22270, gen22274, gen22280)

			return

		} else {
			if True == True {
				__e.Return(V836)
				return
			} else {
				gen22269 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22269, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1macro, gen22300)

	gen22313 := MakeNative(func(__e *ControlFlow) {
		V839 := __e.Get(1)
		_ = V839
		V840 := __e.Get(2)
		_ = V840
		gen22311 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22312 := Call(__e, gen22311, MakeNumber(0), V840)

		if True == gen22312 {
			gen22308 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen22309 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen22310 := Call(__e, gen22309, V839, MakeString(" has no lambda form\n"), symshen_4a)

			__e.TailApply(gen22308, gen22310)

			return

		} else {
			gen22306 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen22307 := Call(__e, gen22306, MakeNumber(-1), V840)

			if True == gen22307 {
				gen22303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22305 := Call(__e, gen22304, V839, Nil)

				__e.TailApply(gen22303, symfunction, gen22305)

				return

			} else {
				if True == True {
					gen22302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction_1help)

					__e.TailApply(gen22302, V839, V840, Nil)

					return

				} else {
					gen22301 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen22301, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction, gen22313)

	gen22334 := MakeNative(func(__e *ControlFlow) {
		V844 := __e.Get(1)
		_ = V844
		V845 := __e.Get(2)
		_ = V845
		V846 := __e.Get(3)
		_ = V846
		gen22332 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22333 := Call(__e, gen22332, MakeNumber(0), V845)

		if True == gen22333 {
			gen22331 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen22331, V844, V846)

			return

		} else {
			if True == True {
				gen22328 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen22315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4function_1abstraction_1help)

					gen22319 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen22320 := Call(__e, gen22319, V845, MakeNumber(1))

					gen22321 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					gen22322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen22323 := Call(__e, gen22322, X, Nil)

					gen22324 := Call(__e, gen22321, V846, gen22323)

					gen22325 := Call(__e, gen22318, V844, gen22320, gen22324)

					gen22326 := Call(__e, gen22317, gen22325, Nil)

					gen22327 := Call(__e, gen22316, X, gen22326)

					__e.TailApply(gen22315, sym_c_4, gen22327)

					return

				}, 1)

				gen22329 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				gen22330 := Call(__e, gen22329, symV)

				__e.TailApply(gen22328, gen22330)

				return

			} else {
				gen22314 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen22314, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4function_1abstraction_1help, gen22334)

	gen22353 := MakeNative(func(__e *ControlFlow) {
		V848 := __e.Get(1)
		_ = V848
		gen22350 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			gen22347 := MakeNative(func(__e *ControlFlow) {
				Pos := __e.Get(1)
				_ = Pos
				gen22342 := MakeNative(func(__e *ControlFlow) {
					Remove1 := __e.Get(1)
					_ = Remove1
					gen22335 := MakeNative(func(__e *ControlFlow) {
						Remove2 := __e.Get(1)
						_ = Remove2
						__e.Return(V848)
						return
					}, 1)

					gen22336 := Call(__e, PrimNS1Value(symns2_1value), symset)

					gen22337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

					gen22338 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					gen22339 := Call(__e, gen22338, sym_dmacros_d)

					gen22340 := Call(__e, gen22337, Pos, gen22339)

					gen22341 := Call(__e, gen22336, sym_dmacros_d, gen22340)

					__e.TailApply(gen22335, gen22341)

					return

				}, 1)

				gen22343 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen22344 := Call(__e, PrimNS1Value(symns2_1value), symremove)

				gen22345 := Call(__e, gen22344, V848, MacroReg)

				gen22346 := Call(__e, gen22343, symshen_4_dmacroreg_d, gen22345)

				__e.TailApply(gen22342, gen22346)

				return

			}, 1)

			gen22348 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

			gen22349 := Call(__e, gen22348, V848, MacroReg)

			__e.TailApply(gen22347, gen22349)

			return

		}, 1)

		gen22351 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen22352 := Call(__e, gen22351, symshen_4_dmacroreg_d)

		__e.TailApply(gen22350, gen22352)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symundefmacro, gen22353)

	gen22375 := MakeNative(func(__e *ControlFlow) {
		V858 := __e.Get(1)
		_ = V858
		V859 := __e.Get(2)
		_ = V859
		gen22373 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22374 := Call(__e, gen22373, Nil, V859)

		if True == gen22374 {
			gen22370 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen22371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen22372 := Call(__e, gen22371, V858, MakeString(" is not a macro\n"), symshen_4a)

			__e.TailApply(gen22370, gen22372)

			return

		} else {
			gen22367 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen22368 := Call(__e, gen22367, V859)

			var gen22369 Obj

			if True == gen22368 {
				gen22363 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen22364 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22365 := Call(__e, gen22364, V859)

				gen22366 := Call(__e, gen22363, gen22365, V858)

				if True == gen22366 {
					gen22369 = True
				} else {
					gen22369 = False
				}

			} else {
				gen22369 = False
			}

			if True == gen22369 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen22361 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen22362 := Call(__e, gen22361, V859)

				if True == gen22362 {
					gen22356 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					gen22357 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findpos)

					gen22358 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen22359 := Call(__e, gen22358, V859)

					gen22360 := Call(__e, gen22357, V858, gen22359)

					__e.TailApply(gen22356, MakeNumber(1), gen22360)

					return

				} else {
					if True == True {
						gen22355 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen22355, symshen_4findpos)

						return

					} else {
						gen22354 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen22354, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4findpos, gen22375)

	gen22395 := MakeNative(func(__e *ControlFlow) {
		V864 := __e.Get(1)
		_ = V864
		V865 := __e.Get(2)
		_ = V865
		gen22392 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen22393 := Call(__e, gen22392, MakeNumber(1), V864)

		var gen22394 Obj

		if True == gen22393 {
			gen22390 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen22391 := Call(__e, gen22390, V865)

			if True == gen22391 {
				gen22394 = True
			} else {
				gen22394 = False
			}

		} else {
			gen22394 = False
		}

		if True == gen22394 {
			gen22389 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(gen22389, V865)

			return

		} else {
			gen22387 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen22388 := Call(__e, gen22387, V865)

			if True == gen22388 {
				gen22378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen22379 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen22380 := Call(__e, gen22379, V865)

				gen22381 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1nth)

				gen22382 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen22383 := Call(__e, gen22382, V864, MakeNumber(1))

				gen22384 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen22385 := Call(__e, gen22384, V865)

				gen22386 := Call(__e, gen22381, gen22383, gen22385)

				__e.TailApply(gen22378, gen22380, gen22386)

				return

			} else {
				if True == True {
					gen22377 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen22377, symshen_4remove_1nth)

					return

				} else {
					gen22376 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen22376, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4remove_1nth, gen22395)

	return

}, 0)
