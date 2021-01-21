package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp94743 := MakeNative(func(__e *ControlFlow) {
		V2731 := __e.Get(1)
		_ = V2731
		V2732 := __e.Get(2)
		_ = V2732
		tmp94744 := MakeNative(func(__e *ControlFlow) {
			Curry := __e.Get(1)
			_ = Curry
			tmp94745 := MakeNative(func(__e *ControlFlow) {
				ProcessN := __e.Get(1)
				_ = ProcessN
				tmp94746 := MakeNative(func(__e *ControlFlow) {
					Type := __e.Get(1)
					_ = Type
					tmp94747 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp94748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d)

						tmp94749 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94750 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94751 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94752 := Call(__e, tmp94751, Type, Nil)

						tmp94753 := Call(__e, tmp94750, sym_h, tmp94752)

						tmp94754 := Call(__e, tmp94749, Curry, tmp94753)

						__e.TailApply(tmp94748, tmp94754, Nil, ProcessN, Continuation)
						return

					}, 1)

					tmp94755 := MakeNative(func(__e *ControlFlow) {
						tmp94756 := Call(__e, PrimNS1Value(symns2_1value), symreturn)

						__e.TailApply(tmp94756, Type, ProcessN, symshen_4void)
						return

					}, 0)

					__e.TailApply(tmp94747, tmp94755)
					return

				}, 1)

				tmp94757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1prolog_1variables)

				tmp94758 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

				tmp94759 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

				tmp94760 := Call(__e, tmp94759, V2732)

				tmp94761 := Call(__e, tmp94758, tmp94760)

				tmp94762 := Call(__e, tmp94757, tmp94761, ProcessN)

				__e.TailApply(tmp94746, tmp94762)
				return

			}, 1)

			tmp94763 := Call(__e, PrimNS1Value(symns2_1value), symshen_4start_1new_1prolog_1process)

			tmp94764 := Call(__e, tmp94763)

			__e.TailApply(tmp94745, tmp94764)
			return

		}, 1)

		tmp94765 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

		tmp94766 := Call(__e, tmp94765, V2731)

		__e.TailApply(tmp94744, tmp94766)
		return

	}, 2)

	tmp94767 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck, tmp94743)

	_ = tmp94767

	tmp94768 := MakeNative(func(__e *ControlFlow) {
		V2734 := __e.Get(1)
		_ = V2734
		tmp94910 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp94911 := Call(__e, tmp94910, V2734)

		var ifres94904 Obj

		if True == tmp94911 {
			tmp94906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4special_2)

			tmp94907 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94908 := Call(__e, tmp94907, V2734)

			tmp94909 := Call(__e, tmp94906, tmp94908)

			var ifres94905 Obj

			if True == tmp94909 {
				ifres94905 = True

			} else {
				ifres94905 = False

			}

			ifres94904 = ifres94905

		} else {
			ifres94904 = False

		}

		if True == ifres94904 {
			tmp94895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp94896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94897 := Call(__e, tmp94896, V2734)

			tmp94898 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp94899 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp94900 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

				__e.TailApply(tmp94900, Y)
				return

			}, 1)

			tmp94901 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94902 := Call(__e, tmp94901, V2734)

			tmp94903 := Call(__e, tmp94898, tmp94899, tmp94902)

			__e.TailApply(tmp94895, tmp94897, tmp94903)
			return

		} else {
			tmp94893 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp94894 := Call(__e, tmp94893, V2734)

			var ifres94881 Obj

			if True == tmp94894 {
				tmp94889 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94891 := Call(__e, tmp94890, V2734)

				tmp94892 := Call(__e, tmp94889, tmp94891)

				var ifres94883 Obj

				if True == tmp94892 {
					tmp94885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extraspecial_2)

					tmp94886 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94887 := Call(__e, tmp94886, V2734)

					tmp94888 := Call(__e, tmp94885, tmp94887)

					var ifres94884 Obj

					if True == tmp94888 {
						ifres94884 = True

					} else {
						ifres94884 = False

					}

					ifres94883 = ifres94884

				} else {
					ifres94883 = False

				}

				var ifres94882 Obj

				if True == ifres94883 {
					ifres94882 = True

				} else {
					ifres94882 = False

				}

				ifres94881 = ifres94882

			} else {
				ifres94881 = False

			}

			if True == ifres94881 {
				__e.Return(V2734)
				return
			} else {
				tmp94879 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94880 := Call(__e, tmp94879, V2734)

				var ifres94849 Obj

				if True == tmp94880 {
					tmp94875 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp94876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94877 := Call(__e, tmp94876, V2734)

					tmp94878 := Call(__e, tmp94875, symtype, tmp94877)

					var ifres94851 Obj

					if True == tmp94878 {
						tmp94871 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp94872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94873 := Call(__e, tmp94872, V2734)

						tmp94874 := Call(__e, tmp94871, tmp94873)

						var ifres94853 Obj

						if True == tmp94874 {
							tmp94865 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp94866 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94868 := Call(__e, tmp94867, V2734)

							tmp94869 := Call(__e, tmp94866, tmp94868)

							tmp94870 := Call(__e, tmp94865, tmp94869)

							var ifres94855 Obj

							if True == tmp94870 {
								tmp94857 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp94858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94859 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94860 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94861 := Call(__e, tmp94860, V2734)

								tmp94862 := Call(__e, tmp94859, tmp94861)

								tmp94863 := Call(__e, tmp94858, tmp94862)

								tmp94864 := Call(__e, tmp94857, Nil, tmp94863)

								var ifres94856 Obj

								if True == tmp94864 {
									ifres94856 = True

								} else {
									ifres94856 = False

								}

								ifres94855 = ifres94856

							} else {
								ifres94855 = False

							}

							var ifres94854 Obj

							if True == ifres94855 {
								ifres94854 = True

							} else {
								ifres94854 = False

							}

							ifres94853 = ifres94854

						} else {
							ifres94853 = False

						}

						var ifres94852 Obj

						if True == ifres94853 {
							ifres94852 = True

						} else {
							ifres94852 = False

						}

						ifres94851 = ifres94852

					} else {
						ifres94851 = False

					}

					var ifres94850 Obj

					if True == ifres94851 {
						ifres94850 = True

					} else {
						ifres94850 = False

					}

					ifres94849 = ifres94850

				} else {
					ifres94849 = False

				}

				if True == ifres94849 {
					tmp94836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

					tmp94839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94841 := Call(__e, tmp94840, V2734)

					tmp94842 := Call(__e, tmp94839, tmp94841)

					tmp94843 := Call(__e, tmp94838, tmp94842)

					tmp94844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94846 := Call(__e, tmp94845, V2734)

					tmp94847 := Call(__e, tmp94844, tmp94846)

					tmp94848 := Call(__e, tmp94837, tmp94843, tmp94847)

					__e.TailApply(tmp94836, symtype, tmp94848)
					return

				} else {
					tmp94834 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp94835 := Call(__e, tmp94834, V2734)

					var ifres94820 Obj

					if True == tmp94835 {
						tmp94830 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp94831 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94832 := Call(__e, tmp94831, V2734)

						tmp94833 := Call(__e, tmp94830, tmp94832)

						var ifres94822 Obj

						if True == tmp94833 {
							tmp94824 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp94825 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94826 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94827 := Call(__e, tmp94826, V2734)

							tmp94828 := Call(__e, tmp94825, tmp94827)

							tmp94829 := Call(__e, tmp94824, tmp94828)

							var ifres94823 Obj

							if True == tmp94829 {
								ifres94823 = True

							} else {
								ifres94823 = False

							}

							ifres94822 = ifres94823

						} else {
							ifres94822 = False

						}

						var ifres94821 Obj

						if True == ifres94822 {
							ifres94821 = True

						} else {
							ifres94821 = False

						}

						ifres94820 = ifres94821

					} else {
						ifres94820 = False

					}

					if True == ifres94820 {
						tmp94803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

						tmp94804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94806 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp94807 := Call(__e, tmp94806, V2734)

						tmp94808 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp94809 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp94810 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94811 := Call(__e, tmp94810, V2734)

						tmp94812 := Call(__e, tmp94809, tmp94811)

						tmp94813 := Call(__e, tmp94808, tmp94812, Nil)

						tmp94814 := Call(__e, tmp94805, tmp94807, tmp94813)

						tmp94815 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94816 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94817 := Call(__e, tmp94816, V2734)

						tmp94818 := Call(__e, tmp94815, tmp94817)

						tmp94819 := Call(__e, tmp94804, tmp94814, tmp94818)

						__e.TailApply(tmp94803, tmp94819)
						return

					} else {
						tmp94801 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp94802 := Call(__e, tmp94801, V2734)

						var ifres94787 Obj

						if True == tmp94802 {
							tmp94797 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp94798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94799 := Call(__e, tmp94798, V2734)

							tmp94800 := Call(__e, tmp94797, tmp94799)

							var ifres94789 Obj

							if True == tmp94800 {
								tmp94791 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp94792 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94793 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94794 := Call(__e, tmp94793, V2734)

								tmp94795 := Call(__e, tmp94792, tmp94794)

								tmp94796 := Call(__e, tmp94791, Nil, tmp94795)

								var ifres94790 Obj

								if True == tmp94796 {
									ifres94790 = True

								} else {
									ifres94790 = False

								}

								ifres94789 = ifres94790

							} else {
								ifres94789 = False

							}

							var ifres94788 Obj

							if True == ifres94789 {
								ifres94788 = True

							} else {
								ifres94788 = False

							}

							ifres94787 = ifres94788

						} else {
							ifres94787 = False

						}

						if True == ifres94787 {
							tmp94774 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp94775 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

							tmp94776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp94777 := Call(__e, tmp94776, V2734)

							tmp94778 := Call(__e, tmp94775, tmp94777)

							tmp94779 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp94780 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

							tmp94781 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp94782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94783 := Call(__e, tmp94782, V2734)

							tmp94784 := Call(__e, tmp94781, tmp94783)

							tmp94785 := Call(__e, tmp94780, tmp94784)

							tmp94786 := Call(__e, tmp94779, tmp94785, Nil)

							__e.TailApply(tmp94774, tmp94778, tmp94786)
							return

						} else {
							__e.Return(V2734)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp94912 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry, tmp94768)

	_ = tmp94912

	tmp94913 := MakeNative(func(__e *ControlFlow) {
		V2736 := __e.Get(1)
		_ = V2736
		tmp94914 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp94915 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94916 := Call(__e, tmp94915, symshen_4_dspecial_d)

		__e.TailApply(tmp94914, V2736, tmp94916)
		return

	}, 1)

	tmp94917 := Call(__e, PrimNS1Value(symns2_1set), symshen_4special_2, tmp94913)

	_ = tmp94917

	tmp94918 := MakeNative(func(__e *ControlFlow) {
		V2738 := __e.Get(1)
		_ = V2738
		tmp94919 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp94920 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94921 := Call(__e, tmp94920, symshen_4_dextraspecial_d)

		__e.TailApply(tmp94919, V2738, tmp94921)
		return

	}, 1)

	tmp94922 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extraspecial_2, tmp94918)

	_ = tmp94922

	tmp94923 := MakeNative(func(__e *ControlFlow) {
		V2743 := __e.Get(1)
		_ = V2743
		V2744 := __e.Get(2)
		_ = V2744
		V2745 := __e.Get(3)
		_ = V2745
		V2746 := __e.Get(4)
		_ = V2746
		tmp94924 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp94925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cutpoint)

			tmp94926 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp95014 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp95015 := Call(__e, tmp95014, Case, False)

				if True == tmp95015 {
					tmp94928 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp95000 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp95001 := Call(__e, tmp95000, Case, False)

						if True == tmp95001 {
							tmp94930 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp94944 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp94945 := Call(__e, tmp94944, Case, False)

								if True == tmp94945 {
									tmp94932 := MakeNative(func(__e *ControlFlow) {
										Datatypes := __e.Get(1)
										_ = Datatypes
										tmp94933 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

										tmp94934 := Call(__e, tmp94933)

										_ = tmp94934

										tmp94935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show)

										tmp94936 := MakeNative(func(__e *ControlFlow) {
											tmp94937 := Call(__e, PrimNS1Value(symns2_1value), symbind)

											tmp94938 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

											tmp94939 := Call(__e, tmp94938, symshen_4_ddatatypes_d)

											tmp94940 := MakeNative(func(__e *ControlFlow) {
												tmp94941 := Call(__e, PrimNS1Value(symns2_1value), symshen_4udefs_d)

												__e.TailApply(tmp94941, V2743, V2744, Datatypes, V2745, V2746)
												return

											}, 0)

											__e.TailApply(tmp94937, Datatypes, tmp94939, V2745, tmp94940)
											return

										}, 0)

										__e.TailApply(tmp94935, V2743, V2744, V2745, tmp94936)
										return

									}, 1)

									tmp94942 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

									tmp94943 := Call(__e, tmp94942, V2745)

									__e.TailApply(tmp94932, tmp94943)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp94946 := MakeNative(func(__e *ControlFlow) {
								V2724 := __e.Get(1)
								_ = V2724
								tmp94995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp94996 := Call(__e, tmp94995, V2724)

								if True == tmp94996 {
									tmp94948 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp94949 := MakeNative(func(__e *ControlFlow) {
											V2725 := __e.Get(1)
											_ = V2725
											tmp94987 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp94988 := Call(__e, tmp94987, V2725)

											if True == tmp94988 {
												tmp94951 := MakeNative(func(__e *ControlFlow) {
													V2726 := __e.Get(1)
													_ = V2726
													tmp94981 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp94982 := Call(__e, tmp94981, sym_h, V2726)

													if True == tmp94982 {
														tmp94953 := MakeNative(func(__e *ControlFlow) {
															V2727 := __e.Get(1)
															_ = V2727
															tmp94975 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp94976 := Call(__e, tmp94975, V2727)

															if True == tmp94976 {
																tmp94955 := MakeNative(func(__e *ControlFlow) {
																	A := __e.Get(1)
																	_ = A
																	tmp94956 := MakeNative(func(__e *ControlFlow) {
																		V2728 := __e.Get(1)
																		_ = V2728
																		tmp94967 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp94968 := Call(__e, tmp94967, Nil, V2728)

																		if True == tmp94968 {
																			tmp94958 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																			tmp94959 := Call(__e, tmp94958)

																			_ = tmp94959

																			tmp94960 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																			tmp94961 := Call(__e, PrimNS1Value(symns2_1value), symshen_4type_1theory_1enabled_2)

																			tmp94962 := Call(__e, tmp94961)

																			tmp94963 := MakeNative(func(__e *ControlFlow) {
																				tmp94964 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																				tmp94965 := MakeNative(func(__e *ControlFlow) {
																					tmp94966 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																					__e.TailApply(tmp94966, X, A, V2744, V2745, V2746)
																					return

																				}, 0)

																				__e.TailApply(tmp94964, Throwcontrol, V2745, tmp94965)
																				return

																			}, 0)

																			__e.TailApply(tmp94960, tmp94962, V2745, tmp94963)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp94969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp94970 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp94971 := Call(__e, tmp94970, V2727)

																	tmp94972 := Call(__e, tmp94969, tmp94971, V2745)

																	__e.TailApply(tmp94956, tmp94972)
																	return

																}, 1)

																tmp94973 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp94974 := Call(__e, tmp94973, V2727)

																__e.TailApply(tmp94955, tmp94974)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp94977 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp94978 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp94979 := Call(__e, tmp94978, V2725)

														tmp94980 := Call(__e, tmp94977, tmp94979, V2745)

														__e.TailApply(tmp94953, tmp94980)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp94983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp94984 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp94985 := Call(__e, tmp94984, V2725)

												tmp94986 := Call(__e, tmp94983, tmp94985, V2745)

												__e.TailApply(tmp94951, tmp94986)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp94989 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp94990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp94991 := Call(__e, tmp94990, V2724)

										tmp94992 := Call(__e, tmp94989, tmp94991, V2745)

										__e.TailApply(tmp94949, tmp94992)
										return

									}, 1)

									tmp94993 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp94994 := Call(__e, tmp94993, V2724)

									__e.TailApply(tmp94948, tmp94994)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp94997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp94998 := Call(__e, tmp94997, V2743, V2745)

							tmp94999 := Call(__e, tmp94946, tmp94998)

							__e.TailApply(tmp94930, tmp94999)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp95002 := MakeNative(func(__e *ControlFlow) {
						V2723 := __e.Get(1)
						_ = V2723
						tmp95009 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp95010 := Call(__e, tmp95009, symfail, V2723)

						if True == tmp95010 {
							tmp95004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

							tmp95005 := Call(__e, tmp95004)

							_ = tmp95005

							tmp95006 := Call(__e, PrimNS1Value(symns2_1value), symcut)

							tmp95007 := MakeNative(func(__e *ControlFlow) {
								tmp95008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prolog_1failure)

								__e.TailApply(tmp95008, V2745, V2746)
								return

							}, 0)

							__e.TailApply(tmp95006, Throwcontrol, V2745, tmp95007)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp95011 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp95012 := Call(__e, tmp95011, V2743, V2745)

					tmp95013 := Call(__e, tmp95002, tmp95012)

					__e.TailApply(tmp94928, tmp95013)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp95016 := MakeNative(func(__e *ControlFlow) {
				Error := __e.Get(1)
				_ = Error
				tmp95017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp95018 := Call(__e, tmp95017)

				_ = tmp95018

				tmp95019 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

				tmp95020 := Call(__e, PrimNS1Value(symns2_1value), symshen_4maxinfexceeded_2)

				tmp95021 := Call(__e, tmp95020)

				tmp95022 := MakeNative(func(__e *ControlFlow) {
					tmp95023 := Call(__e, PrimNS1Value(symns2_1value), symbind)

					tmp95024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4errormaxinfs)

					tmp95025 := Call(__e, tmp95024)

					__e.TailApply(tmp95023, Error, tmp95025, V2745, V2746)
					return

				}, 0)

				__e.TailApply(tmp95019, tmp95021, V2745, tmp95022)
				return

			}, 1)

			tmp95026 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp95027 := Call(__e, tmp95026, V2745)

			tmp95028 := Call(__e, tmp95016, tmp95027)

			tmp95029 := Call(__e, tmp94926, tmp95028)

			__e.TailApply(tmp94925, Throwcontrol, tmp95029)
			return

		}, 1)

		tmp95030 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catchpoint)

		tmp95031 := Call(__e, tmp95030)

		__e.TailApply(tmp94924, tmp95031)
		return

	}, 4)

	tmp95032 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d, tmp94923)

	_ = tmp95032

	tmp95033 := MakeNative(func(__e *ControlFlow) {
		tmp95034 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(tmp95034, symshen_4_dshen_1type_1theory_1enabled_2_d)
		return

	}, 0)

	tmp95035 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1theory_1enabled_2, tmp95033)

	_ = tmp95035

	tmp95036 := MakeNative(func(__e *ControlFlow) {
		V2752 := __e.Get(1)
		_ = V2752
		tmp95044 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp95045 := Call(__e, tmp95044, sym_7, V2752)

		if True == tmp95045 {
			tmp95043 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp95043, symshen_4_dshen_1type_1theory_1enabled_2_d, True)
			return

		} else {
			tmp95041 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp95042 := Call(__e, tmp95041, sym_1, V2752)

			if True == tmp95042 {
				tmp95040 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp95040, symshen_4_dshen_1type_1theory_1enabled_2_d, False)
				return

			} else {
				tmp95039 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp95039, MakeString("enable-type-theory expects a + or a -\n"))
				return

			}

		}

	}, 1)

	tmp95046 := Call(__e, PrimNS1Value(symns2_1set), symenable_1type_1theory, tmp95036)

	_ = tmp95046

	tmp95047 := MakeNative(func(__e *ControlFlow) {
		V2763 := __e.Get(1)
		_ = V2763
		V2764 := __e.Get(2)
		_ = V2764
		__e.Return(False)
		return
	}, 2)

	tmp95048 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1failure, tmp95047)

	_ = tmp95048

	tmp95049 := MakeNative(func(__e *ControlFlow) {
		tmp95050 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		tmp95051 := Call(__e, PrimNS1Value(symns2_1value), syminferences)

		tmp95052 := Call(__e, tmp95051)

		tmp95053 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp95054 := Call(__e, tmp95053, symshen_4_dmaxinferences_d)

		__e.TailApply(tmp95050, tmp95052, tmp95054)
		return

	}, 0)

	tmp95055 := Call(__e, PrimNS1Value(symns2_1set), symshen_4maxinfexceeded_2, tmp95049)

	_ = tmp95055

	tmp95056 := MakeNative(func(__e *ControlFlow) {
		tmp95057 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(tmp95057, MakeString("maximum inferences exceeded~%"))
		return

	}, 0)

	tmp95058 := Call(__e, PrimNS1Value(symns2_1set), symshen_4errormaxinfs, tmp95056)

	_ = tmp95058

	tmp95059 := MakeNative(func(__e *ControlFlow) {
		V2770 := __e.Get(1)
		_ = V2770
		V2771 := __e.Get(2)
		_ = V2771
		V2772 := __e.Get(3)
		_ = V2772
		V2773 := __e.Get(4)
		_ = V2773
		V2774 := __e.Get(5)
		_ = V2774
		tmp95060 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp95074 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp95075 := Call(__e, tmp95074, Case, False)

			if True == tmp95075 {
				tmp95062 := MakeNative(func(__e *ControlFlow) {
					V2720 := __e.Get(1)
					_ = V2720
					tmp95070 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp95071 := Call(__e, tmp95070, V2720)

					if True == tmp95071 {
						tmp95064 := MakeNative(func(__e *ControlFlow) {
							Ds := __e.Get(1)
							_ = Ds
							tmp95065 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

							tmp95066 := Call(__e, tmp95065)

							_ = tmp95066

							tmp95067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4udefs_d)

							__e.TailApply(tmp95067, V2770, V2771, Ds, V2773, V2774)
							return

						}, 1)

						tmp95068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp95069 := Call(__e, tmp95068, V2720)

						__e.TailApply(tmp95064, tmp95069)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp95072 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp95073 := Call(__e, tmp95072, V2772, V2773)

				__e.TailApply(tmp95062, tmp95073)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp95076 := MakeNative(func(__e *ControlFlow) {
			V2719 := __e.Get(1)
			_ = V2719
			tmp95090 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp95091 := Call(__e, tmp95090, V2719)

			if True == tmp95091 {
				tmp95078 := MakeNative(func(__e *ControlFlow) {
					D := __e.Get(1)
					_ = D
					tmp95079 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

					tmp95080 := Call(__e, tmp95079)

					_ = tmp95080

					tmp95081 := Call(__e, PrimNS1Value(symns2_1value), symcall)

					tmp95082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp95083 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp95084 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp95085 := Call(__e, tmp95084, V2771, Nil)

					tmp95086 := Call(__e, tmp95083, V2770, tmp95085)

					tmp95087 := Call(__e, tmp95082, D, tmp95086)

					__e.TailApply(tmp95081, tmp95087, V2773, V2774)
					return

				}, 1)

				tmp95088 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp95089 := Call(__e, tmp95088, V2719)

				__e.TailApply(tmp95078, tmp95089)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp95092 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp95093 := Call(__e, tmp95092, V2772, V2773)

		tmp95094 := Call(__e, tmp95076, tmp95093)

		__e.TailApply(tmp95060, tmp95094)
		return

	}, 5)

	tmp95095 := Call(__e, PrimNS1Value(symns2_1set), symshen_4udefs_d, tmp95059)

	_ = tmp95095

	tmp95096 := MakeNative(func(__e *ControlFlow) {
		V2780 := __e.Get(1)
		_ = V2780
		V2781 := __e.Get(2)
		_ = V2781
		V2782 := __e.Get(3)
		_ = V2782
		V2783 := __e.Get(4)
		_ = V2783
		V2784 := __e.Get(5)
		_ = V2784
		tmp95097 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp95098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cutpoint)

			tmp95099 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp97225 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp97226 := Call(__e, tmp97225, Case, False)

				if True == tmp97226 {
					tmp95101 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp97200 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp97201 := Call(__e, tmp97200, Case, False)

						if True == tmp97201 {
							tmp95103 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp97194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp97195 := Call(__e, tmp97194, Case, False)

								if True == tmp97195 {
									tmp95105 := MakeNative(func(__e *ControlFlow) {
										Case := __e.Get(1)
										_ = Case
										tmp97188 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp97189 := Call(__e, tmp97188, Case, False)

										if True == tmp97189 {
											tmp95107 := MakeNative(func(__e *ControlFlow) {
												Case := __e.Get(1)
												_ = Case
												tmp97161 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

												tmp97162 := Call(__e, tmp97161, Case, False)

												if True == tmp97162 {
													tmp95109 := MakeNative(func(__e *ControlFlow) {
														Case := __e.Get(1)
														_ = Case
														tmp97116 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp97117 := Call(__e, tmp97116, Case, False)

														if True == tmp97117 {
															tmp95111 := MakeNative(func(__e *ControlFlow) {
																Case := __e.Get(1)
																_ = Case
																tmp96883 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp96884 := Call(__e, tmp96883, Case, False)

																if True == tmp96884 {
																	tmp95113 := MakeNative(func(__e *ControlFlow) {
																		Case := __e.Get(1)
																		_ = Case
																		tmp96641 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp96642 := Call(__e, tmp96641, Case, False)

																		if True == tmp96642 {
																			tmp95115 := MakeNative(func(__e *ControlFlow) {
																				Case := __e.Get(1)
																				_ = Case
																				tmp96408 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp96409 := Call(__e, tmp96408, Case, False)

																				if True == tmp96409 {
																					tmp95117 := MakeNative(func(__e *ControlFlow) {
																						Case := __e.Get(1)
																						_ = Case
																						tmp96336 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																						tmp96337 := Call(__e, tmp96336, Case, False)

																						if True == tmp96337 {
																							tmp95119 := MakeNative(func(__e *ControlFlow) {
																								Case := __e.Get(1)
																								_ = Case
																								tmp95886 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																								tmp95887 := Call(__e, tmp95886, Case, False)

																								if True == tmp95887 {
																									tmp95121 := MakeNative(func(__e *ControlFlow) {
																										Case := __e.Get(1)
																										_ = Case
																										tmp95792 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																										tmp95793 := Call(__e, tmp95792, Case, False)

																										if True == tmp95793 {
																											tmp95123 := MakeNative(func(__e *ControlFlow) {
																												Case := __e.Get(1)
																												_ = Case
																												tmp95503 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																												tmp95504 := Call(__e, tmp95503, Case, False)

																												if True == tmp95504 {
																													tmp95125 := MakeNative(func(__e *ControlFlow) {
																														Case := __e.Get(1)
																														_ = Case
																														tmp95449 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																														tmp95450 := Call(__e, tmp95449, Case, False)

																														if True == tmp95450 {
																															tmp95127 := MakeNative(func(__e *ControlFlow) {
																																Case := __e.Get(1)
																																_ = Case
																																tmp95384 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp95385 := Call(__e, tmp95384, Case, False)

																																if True == tmp95385 {
																																	tmp95129 := MakeNative(func(__e *ControlFlow) {
																																		Case := __e.Get(1)
																																		_ = Case
																																		tmp95322 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp95323 := Call(__e, tmp95322, Case, False)

																																		if True == tmp95323 {
																																			tmp95131 := MakeNative(func(__e *ControlFlow) {
																																				Case := __e.Get(1)
																																				_ = Case
																																				tmp95311 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																				tmp95312 := Call(__e, tmp95311, Case, False)

																																				if True == tmp95312 {
																																					tmp95133 := MakeNative(func(__e *ControlFlow) {
																																						Case := __e.Get(1)
																																						_ = Case
																																						tmp95271 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp95272 := Call(__e, tmp95271, Case, False)

																																						if True == tmp95272 {
																																							tmp95135 := MakeNative(func(__e *ControlFlow) {
																																								Case := __e.Get(1)
																																								_ = Case
																																								tmp95233 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp95234 := Call(__e, tmp95233, Case, False)

																																								if True == tmp95234 {
																																									tmp95137 := MakeNative(func(__e *ControlFlow) {
																																										Case := __e.Get(1)
																																										_ = Case
																																										tmp95195 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																										tmp95196 := Call(__e, tmp95195, Case, False)

																																										if True == tmp95196 {
																																											tmp95139 := MakeNative(func(__e *ControlFlow) {
																																												Case := __e.Get(1)
																																												_ = Case
																																												tmp95157 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																												tmp95158 := Call(__e, tmp95157, Case, False)

																																												if True == tmp95158 {
																																													tmp95141 := MakeNative(func(__e *ControlFlow) {
																																														Datatypes := __e.Get(1)
																																														_ = Datatypes
																																														tmp95142 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																														tmp95143 := Call(__e, tmp95142)

																																														_ = tmp95143

																																														tmp95144 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																														tmp95145 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

																																														tmp95146 := Call(__e, tmp95145, symshen_4_ddatatypes_d)

																																														tmp95147 := MakeNative(func(__e *ControlFlow) {
																																															tmp95148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4udefs_d)

																																															tmp95149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																															tmp95150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																															tmp95151 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																															tmp95152 := Call(__e, tmp95151, V2781, Nil)

																																															tmp95153 := Call(__e, tmp95150, sym_h, tmp95152)

																																															tmp95154 := Call(__e, tmp95149, V2780, tmp95153)

																																															__e.TailApply(tmp95148, tmp95154, V2782, Datatypes, V2783, V2784)
																																															return

																																														}, 0)

																																														__e.TailApply(tmp95144, Datatypes, tmp95146, V2783, tmp95147)
																																														return

																																													}, 1)

																																													tmp95155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																													tmp95156 := Call(__e, tmp95155, V2783)

																																													__e.TailApply(tmp95141, tmp95156)
																																													return

																																												} else {
																																													__e.Return(Case)
																																													return
																																												}

																																											}, 1)

																																											tmp95159 := MakeNative(func(__e *ControlFlow) {
																																												V2713 := __e.Get(1)
																																												_ = V2713
																																												tmp95190 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																												tmp95191 := Call(__e, tmp95190, V2713)

																																												if True == tmp95191 {
																																													tmp95161 := MakeNative(func(__e *ControlFlow) {
																																														V2714 := __e.Get(1)
																																														_ = V2714
																																														tmp95184 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																														tmp95185 := Call(__e, tmp95184, symshen_4synonyms_1help, V2714)

																																														if True == tmp95185 {
																																															tmp95163 := MakeNative(func(__e *ControlFlow) {
																																																V2715 := __e.Get(1)
																																																_ = V2715
																																																tmp95180 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																																tmp95181 := Call(__e, tmp95180, symsymbol, V2715)

																																																if True == tmp95181 {
																																																	tmp95177 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																	tmp95178 := Call(__e, tmp95177)

																																																	_ = tmp95178

																																																	tmp95179 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																																																	__e.TailApply(tmp95179, V2784)
																																																	return

																																																} else {
																																																	tmp95175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																																	tmp95176 := Call(__e, tmp95175, V2715)

																																																	if True == tmp95176 {
																																																		tmp95166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																		tmp95167 := Call(__e, tmp95166, V2715, symsymbol, V2783)

																																																		_ = tmp95167

																																																		tmp95168 := MakeNative(func(__e *ControlFlow) {
																																																			Result := __e.Get(1)
																																																			_ = Result
																																																			tmp95169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																			tmp95170 := Call(__e, tmp95169, V2715, V2783)

																																																			_ = tmp95170

																																																			__e.Return(Result)
																																																			return

																																																		}, 1)

																																																		tmp95171 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																		tmp95172 := Call(__e, tmp95171)

																																																		_ = tmp95172

																																																		tmp95173 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																																																		tmp95174 := Call(__e, tmp95173, V2784)

																																																		__e.TailApply(tmp95168, tmp95174)
																																																		return

																																																	} else {
																																																		__e.Return(False)
																																																		return
																																																	}

																																																}

																																															}, 1)

																																															tmp95182 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																															tmp95183 := Call(__e, tmp95182, V2781, V2783)

																																															__e.TailApply(tmp95163, tmp95183)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}, 1)

																																													tmp95186 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95187 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																													tmp95188 := Call(__e, tmp95187, V2713)

																																													tmp95189 := Call(__e, tmp95186, tmp95188, V2783)

																																													__e.TailApply(tmp95161, tmp95189)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp95192 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95193 := Call(__e, tmp95192, V2780, V2783)

																																											tmp95194 := Call(__e, tmp95159, tmp95193)

																																											__e.TailApply(tmp95139, tmp95194)
																																											return

																																										} else {
																																											__e.Return(Case)
																																											return
																																										}

																																									}, 1)

																																									tmp95197 := MakeNative(func(__e *ControlFlow) {
																																										V2710 := __e.Get(1)
																																										_ = V2710
																																										tmp95228 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																										tmp95229 := Call(__e, tmp95228, V2710)

																																										if True == tmp95229 {
																																											tmp95199 := MakeNative(func(__e *ControlFlow) {
																																												V2711 := __e.Get(1)
																																												_ = V2711
																																												tmp95222 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																												tmp95223 := Call(__e, tmp95222, symshen_4process_1datatype, V2711)

																																												if True == tmp95223 {
																																													tmp95201 := MakeNative(func(__e *ControlFlow) {
																																														V2712 := __e.Get(1)
																																														_ = V2712
																																														tmp95218 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																														tmp95219 := Call(__e, tmp95218, symsymbol, V2712)

																																														if True == tmp95219 {
																																															tmp95215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																															tmp95216 := Call(__e, tmp95215)

																																															_ = tmp95216

																																															tmp95217 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																																															__e.TailApply(tmp95217, V2784)
																																															return

																																														} else {
																																															tmp95213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																															tmp95214 := Call(__e, tmp95213, V2712)

																																															if True == tmp95214 {
																																																tmp95204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																tmp95205 := Call(__e, tmp95204, V2712, symsymbol, V2783)

																																																_ = tmp95205

																																																tmp95206 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp95207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																	tmp95208 := Call(__e, tmp95207, V2712, V2783)

																																																	_ = tmp95208

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp95209 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																tmp95210 := Call(__e, tmp95209)

																																																_ = tmp95210

																																																tmp95211 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																																																tmp95212 := Call(__e, tmp95211, V2784)

																																																__e.TailApply(tmp95206, tmp95212)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp95220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95221 := Call(__e, tmp95220, V2781, V2783)

																																													__e.TailApply(tmp95201, tmp95221)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp95224 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95225 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																											tmp95226 := Call(__e, tmp95225, V2710)

																																											tmp95227 := Call(__e, tmp95224, tmp95226, V2783)

																																											__e.TailApply(tmp95199, tmp95227)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp95230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp95231 := Call(__e, tmp95230, V2780, V2783)

																																									tmp95232 := Call(__e, tmp95197, tmp95231)

																																									__e.TailApply(tmp95137, tmp95232)
																																									return

																																								} else {
																																									__e.Return(Case)
																																									return
																																								}

																																							}, 1)

																																							tmp95235 := MakeNative(func(__e *ControlFlow) {
																																								V2707 := __e.Get(1)
																																								_ = V2707
																																								tmp95266 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																								tmp95267 := Call(__e, tmp95266, V2707)

																																								if True == tmp95267 {
																																									tmp95237 := MakeNative(func(__e *ControlFlow) {
																																										V2708 := __e.Get(1)
																																										_ = V2708
																																										tmp95260 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																										tmp95261 := Call(__e, tmp95260, symdefmacro, V2708)

																																										if True == tmp95261 {
																																											tmp95239 := MakeNative(func(__e *ControlFlow) {
																																												V2709 := __e.Get(1)
																																												_ = V2709
																																												tmp95256 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																												tmp95257 := Call(__e, tmp95256, symunit, V2709)

																																												if True == tmp95257 {
																																													tmp95253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																													tmp95254 := Call(__e, tmp95253)

																																													_ = tmp95254

																																													tmp95255 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																													__e.TailApply(tmp95255, Throwcontrol, V2783, V2784)
																																													return

																																												} else {
																																													tmp95251 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																													tmp95252 := Call(__e, tmp95251, V2709)

																																													if True == tmp95252 {
																																														tmp95242 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																														tmp95243 := Call(__e, tmp95242, V2709, symunit, V2783)

																																														_ = tmp95243

																																														tmp95244 := MakeNative(func(__e *ControlFlow) {
																																															Result := __e.Get(1)
																																															_ = Result
																																															tmp95245 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																															tmp95246 := Call(__e, tmp95245, V2709, V2783)

																																															_ = tmp95246

																																															__e.Return(Result)
																																															return

																																														}, 1)

																																														tmp95247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																														tmp95248 := Call(__e, tmp95247)

																																														_ = tmp95248

																																														tmp95249 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																														tmp95250 := Call(__e, tmp95249, Throwcontrol, V2783, V2784)

																																														__e.TailApply(tmp95244, tmp95250)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp95258 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95259 := Call(__e, tmp95258, V2781, V2783)

																																											__e.TailApply(tmp95239, tmp95259)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp95262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp95263 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																									tmp95264 := Call(__e, tmp95263, V2707)

																																									tmp95265 := Call(__e, tmp95262, tmp95264, V2783)

																																									__e.TailApply(tmp95237, tmp95265)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp95268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp95269 := Call(__e, tmp95268, V2780, V2783)

																																							tmp95270 := Call(__e, tmp95235, tmp95269)

																																							__e.TailApply(tmp95135, tmp95270)
																																							return

																																						} else {
																																							__e.Return(Case)
																																							return
																																						}

																																					}, 1)

																																					tmp95273 := MakeNative(func(__e *ControlFlow) {
																																						V2704 := __e.Get(1)
																																						_ = V2704
																																						tmp95306 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																						tmp95307 := Call(__e, tmp95306, V2704)

																																						if True == tmp95307 {
																																							tmp95275 := MakeNative(func(__e *ControlFlow) {
																																								V2705 := __e.Get(1)
																																								_ = V2705
																																								tmp95300 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp95301 := Call(__e, tmp95300, symdefine, V2705)

																																								if True == tmp95301 {
																																									tmp95277 := MakeNative(func(__e *ControlFlow) {
																																										V2706 := __e.Get(1)
																																										_ = V2706
																																										tmp95294 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																										tmp95295 := Call(__e, tmp95294, V2706)

																																										if True == tmp95295 {
																																											tmp95279 := MakeNative(func(__e *ControlFlow) {
																																												F := __e.Get(1)
																																												_ = F
																																												tmp95280 := MakeNative(func(__e *ControlFlow) {
																																													X := __e.Get(1)
																																													_ = X
																																													tmp95281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																													tmp95282 := Call(__e, tmp95281)

																																													_ = tmp95282

																																													tmp95283 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																													tmp95284 := MakeNative(func(__e *ControlFlow) {
																																														tmp95285 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1def)

																																														tmp95286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95287 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95288 := Call(__e, tmp95287, F, X)

																																														tmp95289 := Call(__e, tmp95286, symdefine, tmp95288)

																																														__e.TailApply(tmp95285, tmp95289, V2781, V2782, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95283, Throwcontrol, V2783, tmp95284)
																																													return

																																												}, 1)

																																												tmp95290 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																												tmp95291 := Call(__e, tmp95290, V2706)

																																												__e.TailApply(tmp95280, tmp95291)
																																												return

																																											}, 1)

																																											tmp95292 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																											tmp95293 := Call(__e, tmp95292, V2706)

																																											__e.TailApply(tmp95279, tmp95293)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp95296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp95297 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp95298 := Call(__e, tmp95297, V2704)

																																									tmp95299 := Call(__e, tmp95296, tmp95298, V2783)

																																									__e.TailApply(tmp95277, tmp95299)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp95302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp95303 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																							tmp95304 := Call(__e, tmp95303, V2704)

																																							tmp95305 := Call(__e, tmp95302, tmp95304, V2783)

																																							__e.TailApply(tmp95275, tmp95305)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp95308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp95309 := Call(__e, tmp95308, V2780, V2783)

																																					tmp95310 := Call(__e, tmp95273, tmp95309)

																																					__e.TailApply(tmp95133, tmp95310)
																																					return

																																				} else {
																																					__e.Return(Case)
																																					return
																																				}

																																			}, 1)

																																			tmp95313 := MakeNative(func(__e *ControlFlow) {
																																				NewHyp := __e.Get(1)
																																				_ = NewHyp
																																				tmp95314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp95315 := Call(__e, tmp95314)

																																				_ = tmp95315

																																				tmp95316 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1hyps)

																																				tmp95317 := MakeNative(func(__e *ControlFlow) {
																																					tmp95318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					__e.TailApply(tmp95318, V2780, V2781, NewHyp, V2783, V2784)
																																					return

																																				}, 0)

																																				__e.TailApply(tmp95316, V2782, NewHyp, V2783, tmp95317)
																																				return

																																			}, 1)

																																			tmp95319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																			tmp95320 := Call(__e, tmp95319, V2783)

																																			tmp95321 := Call(__e, tmp95313, tmp95320)

																																			__e.TailApply(tmp95131, tmp95321)
																																			return

																																		} else {
																																			__e.Return(Case)
																																			return
																																		}

																																	}, 1)

																																	tmp95324 := MakeNative(func(__e *ControlFlow) {
																																		V2699 := __e.Get(1)
																																		_ = V2699
																																		tmp95379 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																		tmp95380 := Call(__e, tmp95379, V2699)

																																		if True == tmp95380 {
																																			tmp95326 := MakeNative(func(__e *ControlFlow) {
																																				V2700 := __e.Get(1)
																																				_ = V2700
																																				tmp95373 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																				tmp95374 := Call(__e, tmp95373, symset, V2700)

																																				if True == tmp95374 {
																																					tmp95328 := MakeNative(func(__e *ControlFlow) {
																																						V2701 := __e.Get(1)
																																						_ = V2701
																																						tmp95367 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																						tmp95368 := Call(__e, tmp95367, V2701)

																																						if True == tmp95368 {
																																							tmp95330 := MakeNative(func(__e *ControlFlow) {
																																								Var := __e.Get(1)
																																								_ = Var
																																								tmp95331 := MakeNative(func(__e *ControlFlow) {
																																									V2702 := __e.Get(1)
																																									_ = V2702
																																									tmp95359 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																									tmp95360 := Call(__e, tmp95359, V2702)

																																									if True == tmp95360 {
																																										tmp95333 := MakeNative(func(__e *ControlFlow) {
																																											Val := __e.Get(1)
																																											_ = Val
																																											tmp95334 := MakeNative(func(__e *ControlFlow) {
																																												V2703 := __e.Get(1)
																																												_ = V2703
																																												tmp95351 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																												tmp95352 := Call(__e, tmp95351, Nil, V2703)

																																												if True == tmp95352 {
																																													tmp95336 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																													tmp95337 := Call(__e, tmp95336)

																																													_ = tmp95337

																																													tmp95338 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																													tmp95339 := MakeNative(func(__e *ControlFlow) {
																																														tmp95340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																														tmp95341 := MakeNative(func(__e *ControlFlow) {
																																															tmp95342 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																															tmp95343 := MakeNative(func(__e *ControlFlow) {
																																																tmp95344 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																tmp95345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																tmp95346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																tmp95347 := Call(__e, tmp95346, Var, Nil)

																																																tmp95348 := Call(__e, tmp95345, symvalue, tmp95347)

																																																tmp95349 := MakeNative(func(__e *ControlFlow) {
																																																	tmp95350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																	__e.TailApply(tmp95350, Val, V2781, V2782, V2783, V2784)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp95344, tmp95348, V2781, V2782, V2783, tmp95349)
																																																return

																																															}, 0)

																																															__e.TailApply(tmp95342, Throwcontrol, V2783, tmp95343)
																																															return

																																														}, 0)

																																														__e.TailApply(tmp95340, Var, symsymbol, V2782, V2783, tmp95341)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95338, Throwcontrol, V2783, tmp95339)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp95353 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95354 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																											tmp95355 := Call(__e, tmp95354, V2702)

																																											tmp95356 := Call(__e, tmp95353, tmp95355, V2783)

																																											__e.TailApply(tmp95334, tmp95356)
																																											return

																																										}, 1)

																																										tmp95357 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																										tmp95358 := Call(__e, tmp95357, V2702)

																																										__e.TailApply(tmp95333, tmp95358)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp95361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp95362 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp95363 := Call(__e, tmp95362, V2701)

																																								tmp95364 := Call(__e, tmp95361, tmp95363, V2783)

																																								__e.TailApply(tmp95331, tmp95364)
																																								return

																																							}, 1)

																																							tmp95365 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																							tmp95366 := Call(__e, tmp95365, V2701)

																																							__e.TailApply(tmp95330, tmp95366)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp95369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp95370 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp95371 := Call(__e, tmp95370, V2699)

																																					tmp95372 := Call(__e, tmp95369, tmp95371, V2783)

																																					__e.TailApply(tmp95328, tmp95372)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp95375 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp95376 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																			tmp95377 := Call(__e, tmp95376, V2699)

																																			tmp95378 := Call(__e, tmp95375, tmp95377, V2783)

																																			__e.TailApply(tmp95326, tmp95378)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp95381 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp95382 := Call(__e, tmp95381, V2780, V2783)

																																	tmp95383 := Call(__e, tmp95324, tmp95382)

																																	__e.TailApply(tmp95129, tmp95383)
																																	return

																																} else {
																																	__e.Return(Case)
																																	return
																																}

																															}, 1)

																															tmp95386 := MakeNative(func(__e *ControlFlow) {
																																V2694 := __e.Get(1)
																																_ = V2694
																																tmp95444 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																tmp95445 := Call(__e, tmp95444, V2694)

																																if True == tmp95445 {
																																	tmp95388 := MakeNative(func(__e *ControlFlow) {
																																		V2695 := __e.Get(1)
																																		_ = V2695
																																		tmp95438 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp95439 := Call(__e, tmp95438, syminput_7, V2695)

																																		if True == tmp95439 {
																																			tmp95390 := MakeNative(func(__e *ControlFlow) {
																																				V2696 := __e.Get(1)
																																				_ = V2696
																																				tmp95432 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																				tmp95433 := Call(__e, tmp95432, V2696)

																																				if True == tmp95433 {
																																					tmp95392 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp95393 := MakeNative(func(__e *ControlFlow) {
																																							V2697 := __e.Get(1)
																																							_ = V2697
																																							tmp95424 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																							tmp95425 := Call(__e, tmp95424, V2697)

																																							if True == tmp95425 {
																																								tmp95395 := MakeNative(func(__e *ControlFlow) {
																																									Stream := __e.Get(1)
																																									_ = Stream
																																									tmp95396 := MakeNative(func(__e *ControlFlow) {
																																										V2698 := __e.Get(1)
																																										_ = V2698
																																										tmp95416 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																										tmp95417 := Call(__e, tmp95416, Nil, V2698)

																																										if True == tmp95417 {
																																											tmp95398 := MakeNative(func(__e *ControlFlow) {
																																												C := __e.Get(1)
																																												_ = C
																																												tmp95399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																												tmp95400 := Call(__e, tmp95399)

																																												_ = tmp95400

																																												tmp95401 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																												tmp95402 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

																																												tmp95403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp95404 := Call(__e, tmp95403, A, V2783)

																																												tmp95405 := Call(__e, tmp95402, tmp95404)

																																												tmp95406 := MakeNative(func(__e *ControlFlow) {
																																													tmp95407 := Call(__e, PrimNS1Value(symns2_1value), symunify)

																																													tmp95408 := MakeNative(func(__e *ControlFlow) {
																																														tmp95409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																														tmp95410 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95412 := Call(__e, tmp95411, symin, Nil)

																																														tmp95413 := Call(__e, tmp95410, symstream, tmp95412)

																																														__e.TailApply(tmp95409, Stream, tmp95413, V2782, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95407, V2781, C, V2783, tmp95408)
																																													return

																																												}, 0)

																																												__e.TailApply(tmp95401, C, tmp95405, V2783, tmp95406)
																																												return

																																											}, 1)

																																											tmp95414 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																											tmp95415 := Call(__e, tmp95414, V2783)

																																											__e.TailApply(tmp95398, tmp95415)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp95418 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp95419 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp95420 := Call(__e, tmp95419, V2697)

																																									tmp95421 := Call(__e, tmp95418, tmp95420, V2783)

																																									__e.TailApply(tmp95396, tmp95421)
																																									return

																																								}, 1)

																																								tmp95422 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																								tmp95423 := Call(__e, tmp95422, V2697)

																																								__e.TailApply(tmp95395, tmp95423)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp95426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp95427 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp95428 := Call(__e, tmp95427, V2696)

																																						tmp95429 := Call(__e, tmp95426, tmp95428, V2783)

																																						__e.TailApply(tmp95393, tmp95429)
																																						return

																																					}, 1)

																																					tmp95430 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																					tmp95431 := Call(__e, tmp95430, V2696)

																																					__e.TailApply(tmp95392, tmp95431)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp95434 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp95435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp95436 := Call(__e, tmp95435, V2694)

																																			tmp95437 := Call(__e, tmp95434, tmp95436, V2783)

																																			__e.TailApply(tmp95390, tmp95437)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp95440 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp95441 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																	tmp95442 := Call(__e, tmp95441, V2694)

																																	tmp95443 := Call(__e, tmp95440, tmp95442, V2783)

																																	__e.TailApply(tmp95388, tmp95443)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp95446 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp95447 := Call(__e, tmp95446, V2780, V2783)

																															tmp95448 := Call(__e, tmp95386, tmp95447)

																															__e.TailApply(tmp95127, tmp95448)
																															return

																														} else {
																															__e.Return(Case)
																															return
																														}

																													}, 1)

																													tmp95451 := MakeNative(func(__e *ControlFlow) {
																														V2689 := __e.Get(1)
																														_ = V2689
																														tmp95498 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																														tmp95499 := Call(__e, tmp95498, V2689)

																														if True == tmp95499 {
																															tmp95453 := MakeNative(func(__e *ControlFlow) {
																																V2690 := __e.Get(1)
																																_ = V2690
																																tmp95492 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp95493 := Call(__e, tmp95492, symtype, V2690)

																																if True == tmp95493 {
																																	tmp95455 := MakeNative(func(__e *ControlFlow) {
																																		V2691 := __e.Get(1)
																																		_ = V2691
																																		tmp95486 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																		tmp95487 := Call(__e, tmp95486, V2691)

																																		if True == tmp95487 {
																																			tmp95457 := MakeNative(func(__e *ControlFlow) {
																																				X := __e.Get(1)
																																				_ = X
																																				tmp95458 := MakeNative(func(__e *ControlFlow) {
																																					V2692 := __e.Get(1)
																																					_ = V2692
																																					tmp95478 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																					tmp95479 := Call(__e, tmp95478, V2692)

																																					if True == tmp95479 {
																																						tmp95460 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp95461 := MakeNative(func(__e *ControlFlow) {
																																								V2693 := __e.Get(1)
																																								_ = V2693
																																								tmp95470 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp95471 := Call(__e, tmp95470, Nil, V2693)

																																								if True == tmp95471 {
																																									tmp95463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp95464 := Call(__e, tmp95463)

																																									_ = tmp95464

																																									tmp95465 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																									tmp95466 := MakeNative(func(__e *ControlFlow) {
																																										tmp95467 := Call(__e, PrimNS1Value(symns2_1value), symunify)

																																										tmp95468 := MakeNative(func(__e *ControlFlow) {
																																											tmp95469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											__e.TailApply(tmp95469, X, A, V2782, V2783, V2784)
																																											return

																																										}, 0)

																																										__e.TailApply(tmp95467, A, V2781, V2783, tmp95468)
																																										return

																																									}, 0)

																																									__e.TailApply(tmp95465, Throwcontrol, V2783, tmp95466)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp95472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp95473 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp95474 := Call(__e, tmp95473, V2692)

																																							tmp95475 := Call(__e, tmp95472, tmp95474, V2783)

																																							__e.TailApply(tmp95461, tmp95475)
																																							return

																																						}, 1)

																																						tmp95476 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																						tmp95477 := Call(__e, tmp95476, V2692)

																																						__e.TailApply(tmp95460, tmp95477)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp95480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp95481 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp95482 := Call(__e, tmp95481, V2691)

																																				tmp95483 := Call(__e, tmp95480, tmp95482, V2783)

																																				__e.TailApply(tmp95458, tmp95483)
																																				return

																																			}, 1)

																																			tmp95484 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																			tmp95485 := Call(__e, tmp95484, V2691)

																																			__e.TailApply(tmp95457, tmp95485)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp95488 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp95489 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp95490 := Call(__e, tmp95489, V2689)

																																	tmp95491 := Call(__e, tmp95488, tmp95490, V2783)

																																	__e.TailApply(tmp95455, tmp95491)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp95494 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp95495 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																															tmp95496 := Call(__e, tmp95495, V2689)

																															tmp95497 := Call(__e, tmp95494, tmp95496, V2783)

																															__e.TailApply(tmp95453, tmp95497)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp95500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp95501 := Call(__e, tmp95500, V2780, V2783)

																													tmp95502 := Call(__e, tmp95451, tmp95501)

																													__e.TailApply(tmp95125, tmp95502)
																													return

																												} else {
																													__e.Return(Case)
																													return
																												}

																											}, 1)

																											tmp95505 := MakeNative(func(__e *ControlFlow) {
																												V2678 := __e.Get(1)
																												_ = V2678
																												tmp95787 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																												tmp95788 := Call(__e, tmp95787, V2678)

																												if True == tmp95788 {
																													tmp95507 := MakeNative(func(__e *ControlFlow) {
																														V2679 := __e.Get(1)
																														_ = V2679
																														tmp95781 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																														tmp95782 := Call(__e, tmp95781, symopen, V2679)

																														if True == tmp95782 {
																															tmp95509 := MakeNative(func(__e *ControlFlow) {
																																V2680 := __e.Get(1)
																																_ = V2680
																																tmp95775 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																tmp95776 := Call(__e, tmp95775, V2680)

																																if True == tmp95776 {
																																	tmp95511 := MakeNative(func(__e *ControlFlow) {
																																		FileName := __e.Get(1)
																																		_ = FileName
																																		tmp95512 := MakeNative(func(__e *ControlFlow) {
																																			V2681 := __e.Get(1)
																																			_ = V2681
																																			tmp95767 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																			tmp95768 := Call(__e, tmp95767, V2681)

																																			if True == tmp95768 {
																																				tmp95514 := MakeNative(func(__e *ControlFlow) {
																																					Direction2611 := __e.Get(1)
																																					_ = Direction2611
																																					tmp95515 := MakeNative(func(__e *ControlFlow) {
																																						V2682 := __e.Get(1)
																																						_ = V2682
																																						tmp95759 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp95760 := Call(__e, tmp95759, Nil, V2682)

																																						if True == tmp95760 {
																																							tmp95517 := MakeNative(func(__e *ControlFlow) {
																																								V2683 := __e.Get(1)
																																								_ = V2683
																																								tmp95755 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																								tmp95756 := Call(__e, tmp95755, V2683)

																																								if True == tmp95756 {
																																									tmp95552 := MakeNative(func(__e *ControlFlow) {
																																										V2684 := __e.Get(1)
																																										_ = V2684
																																										tmp95749 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																										tmp95750 := Call(__e, tmp95749, symstream, V2684)

																																										if True == tmp95750 {
																																											tmp95656 := MakeNative(func(__e *ControlFlow) {
																																												V2685 := __e.Get(1)
																																												_ = V2685
																																												tmp95743 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																												tmp95744 := Call(__e, tmp95743, V2685)

																																												if True == tmp95744 {
																																													tmp95689 := MakeNative(func(__e *ControlFlow) {
																																														Direction := __e.Get(1)
																																														_ = Direction
																																														tmp95690 := MakeNative(func(__e *ControlFlow) {
																																															V2686 := __e.Get(1)
																																															_ = V2686
																																															tmp95735 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																															tmp95736 := Call(__e, tmp95735, Nil, V2686)

																																															if True == tmp95736 {
																																																tmp95718 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																tmp95719 := Call(__e, tmp95718)

																																																_ = tmp95719

																																																tmp95720 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																																tmp95721 := MakeNative(func(__e *ControlFlow) {
																																																	tmp95722 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																	tmp95723 := MakeNative(func(__e *ControlFlow) {
																																																		tmp95724 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																		tmp95725 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																		tmp95726 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp95727 := Call(__e, tmp95726, Direction, V2783)

																																																		tmp95728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp95729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp95730 := Call(__e, tmp95729, symout, Nil)

																																																		tmp95731 := Call(__e, tmp95728, symin, tmp95730)

																																																		tmp95732 := Call(__e, tmp95725, tmp95727, tmp95731)

																																																		tmp95733 := MakeNative(func(__e *ControlFlow) {
																																																			tmp95734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																			__e.TailApply(tmp95734, FileName, symstring, V2782, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp95724, tmp95732, V2783, tmp95733)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp95722, Throwcontrol, V2783, tmp95723)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp95720, Direction, Direction2611, V2783, tmp95721)
																																																return

																																															} else {
																																																tmp95716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																																tmp95717 := Call(__e, tmp95716, V2686)

																																																if True == tmp95717 {
																																																	tmp95693 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																	tmp95694 := Call(__e, tmp95693, V2686, Nil, V2783)

																																																	_ = tmp95694

																																																	tmp95695 := MakeNative(func(__e *ControlFlow) {
																																																		Result := __e.Get(1)
																																																		_ = Result
																																																		tmp95696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																		tmp95697 := Call(__e, tmp95696, V2686, V2783)

																																																		_ = tmp95697

																																																		__e.Return(Result)
																																																		return

																																																	}, 1)

																																																	tmp95698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																	tmp95699 := Call(__e, tmp95698)

																																																	_ = tmp95699

																																																	tmp95700 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																																	tmp95701 := MakeNative(func(__e *ControlFlow) {
																																																		tmp95702 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																		tmp95703 := MakeNative(func(__e *ControlFlow) {
																																																			tmp95704 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																			tmp95705 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																			tmp95706 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp95707 := Call(__e, tmp95706, Direction, V2783)

																																																			tmp95708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp95709 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp95710 := Call(__e, tmp95709, symout, Nil)

																																																			tmp95711 := Call(__e, tmp95708, symin, tmp95710)

																																																			tmp95712 := Call(__e, tmp95705, tmp95707, tmp95711)

																																																			tmp95713 := MakeNative(func(__e *ControlFlow) {
																																																				tmp95714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																				__e.TailApply(tmp95714, FileName, symstring, V2782, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp95704, tmp95712, V2783, tmp95713)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp95702, Throwcontrol, V2783, tmp95703)
																																																		return

																																																	}, 0)

																																																	tmp95715 := Call(__e, tmp95700, Direction, Direction2611, V2783, tmp95701)

																																																	__e.TailApply(tmp95695, tmp95715)
																																																	return

																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp95737 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																														tmp95738 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																														tmp95739 := Call(__e, tmp95738, V2685)

																																														tmp95740 := Call(__e, tmp95737, tmp95739, V2783)

																																														__e.TailApply(tmp95690, tmp95740)
																																														return

																																													}, 1)

																																													tmp95741 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																													tmp95742 := Call(__e, tmp95741, V2685)

																																													__e.TailApply(tmp95689, tmp95742)
																																													return

																																												} else {
																																													tmp95687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																													tmp95688 := Call(__e, tmp95687, V2685)

																																													if True == tmp95688 {
																																														tmp95659 := MakeNative(func(__e *ControlFlow) {
																																															Direction := __e.Get(1)
																																															_ = Direction
																																															tmp95660 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																															tmp95661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																															tmp95662 := Call(__e, tmp95661, Direction, Nil)

																																															tmp95663 := Call(__e, tmp95660, V2685, tmp95662, V2783)

																																															_ = tmp95663

																																															tmp95664 := MakeNative(func(__e *ControlFlow) {
																																																Result := __e.Get(1)
																																																_ = Result
																																																tmp95665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																tmp95666 := Call(__e, tmp95665, V2685, V2783)

																																																_ = tmp95666

																																																__e.Return(Result)
																																																return

																																															}, 1)

																																															tmp95667 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																															tmp95668 := Call(__e, tmp95667)

																																															_ = tmp95668

																																															tmp95669 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																															tmp95670 := MakeNative(func(__e *ControlFlow) {
																																																tmp95671 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																tmp95672 := MakeNative(func(__e *ControlFlow) {
																																																	tmp95673 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																	tmp95674 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																	tmp95675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																	tmp95676 := Call(__e, tmp95675, Direction, V2783)

																																																	tmp95677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																	tmp95678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																	tmp95679 := Call(__e, tmp95678, symout, Nil)

																																																	tmp95680 := Call(__e, tmp95677, symin, tmp95679)

																																																	tmp95681 := Call(__e, tmp95674, tmp95676, tmp95680)

																																																	tmp95682 := MakeNative(func(__e *ControlFlow) {
																																																		tmp95683 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																		__e.TailApply(tmp95683, FileName, symstring, V2782, V2783, V2784)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp95673, tmp95681, V2783, tmp95682)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp95671, Throwcontrol, V2783, tmp95672)
																																																return

																																															}, 0)

																																															tmp95684 := Call(__e, tmp95669, Direction, Direction2611, V2783, tmp95670)

																																															__e.TailApply(tmp95664, tmp95684)
																																															return

																																														}, 1)

																																														tmp95685 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																														tmp95686 := Call(__e, tmp95685, V2783)

																																														__e.TailApply(tmp95659, tmp95686)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp95745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95746 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																											tmp95747 := Call(__e, tmp95746, V2683)

																																											tmp95748 := Call(__e, tmp95745, tmp95747, V2783)

																																											__e.TailApply(tmp95656, tmp95748)
																																											return

																																										} else {
																																											tmp95654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																											tmp95655 := Call(__e, tmp95654, V2684)

																																											if True == tmp95655 {
																																												tmp95555 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																												tmp95556 := Call(__e, tmp95555, V2684, symstream, V2783)

																																												_ = tmp95556

																																												tmp95557 := MakeNative(func(__e *ControlFlow) {
																																													Result := __e.Get(1)
																																													_ = Result
																																													tmp95558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																													tmp95559 := Call(__e, tmp95558, V2684, V2783)

																																													_ = tmp95559

																																													__e.Return(Result)
																																													return

																																												}, 1)

																																												tmp95560 := MakeNative(func(__e *ControlFlow) {
																																													V2687 := __e.Get(1)
																																													_ = V2687
																																													tmp95647 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																													tmp95648 := Call(__e, tmp95647, V2687)

																																													if True == tmp95648 {
																																														tmp95593 := MakeNative(func(__e *ControlFlow) {
																																															Direction := __e.Get(1)
																																															_ = Direction
																																															tmp95594 := MakeNative(func(__e *ControlFlow) {
																																																V2688 := __e.Get(1)
																																																_ = V2688
																																																tmp95639 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																																tmp95640 := Call(__e, tmp95639, Nil, V2688)

																																																if True == tmp95640 {
																																																	tmp95622 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																	tmp95623 := Call(__e, tmp95622)

																																																	_ = tmp95623

																																																	tmp95624 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																																	tmp95625 := MakeNative(func(__e *ControlFlow) {
																																																		tmp95626 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																		tmp95627 := MakeNative(func(__e *ControlFlow) {
																																																			tmp95628 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																			tmp95629 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																			tmp95630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp95631 := Call(__e, tmp95630, Direction, V2783)

																																																			tmp95632 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp95633 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp95634 := Call(__e, tmp95633, symout, Nil)

																																																			tmp95635 := Call(__e, tmp95632, symin, tmp95634)

																																																			tmp95636 := Call(__e, tmp95629, tmp95631, tmp95635)

																																																			tmp95637 := MakeNative(func(__e *ControlFlow) {
																																																				tmp95638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																				__e.TailApply(tmp95638, FileName, symstring, V2782, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp95628, tmp95636, V2783, tmp95637)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp95626, Throwcontrol, V2783, tmp95627)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp95624, Direction, Direction2611, V2783, tmp95625)
																																																	return

																																																} else {
																																																	tmp95620 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																																	tmp95621 := Call(__e, tmp95620, V2688)

																																																	if True == tmp95621 {
																																																		tmp95597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																		tmp95598 := Call(__e, tmp95597, V2688, Nil, V2783)

																																																		_ = tmp95598

																																																		tmp95599 := MakeNative(func(__e *ControlFlow) {
																																																			Result := __e.Get(1)
																																																			_ = Result
																																																			tmp95600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																			tmp95601 := Call(__e, tmp95600, V2688, V2783)

																																																			_ = tmp95601

																																																			__e.Return(Result)
																																																			return

																																																		}, 1)

																																																		tmp95602 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																		tmp95603 := Call(__e, tmp95602)

																																																		_ = tmp95603

																																																		tmp95604 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																																		tmp95605 := MakeNative(func(__e *ControlFlow) {
																																																			tmp95606 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																			tmp95607 := MakeNative(func(__e *ControlFlow) {
																																																				tmp95608 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																				tmp95609 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																				tmp95610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																				tmp95611 := Call(__e, tmp95610, Direction, V2783)

																																																				tmp95612 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp95613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp95614 := Call(__e, tmp95613, symout, Nil)

																																																				tmp95615 := Call(__e, tmp95612, symin, tmp95614)

																																																				tmp95616 := Call(__e, tmp95609, tmp95611, tmp95615)

																																																				tmp95617 := MakeNative(func(__e *ControlFlow) {
																																																					tmp95618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																					__e.TailApply(tmp95618, FileName, symstring, V2782, V2783, V2784)
																																																					return

																																																				}, 0)

																																																				__e.TailApply(tmp95608, tmp95616, V2783, tmp95617)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp95606, Throwcontrol, V2783, tmp95607)
																																																			return

																																																		}, 0)

																																																		tmp95619 := Call(__e, tmp95604, Direction, Direction2611, V2783, tmp95605)

																																																		__e.TailApply(tmp95599, tmp95619)
																																																		return

																																																	} else {
																																																		__e.Return(False)
																																																		return
																																																	}

																																																}

																																															}, 1)

																																															tmp95641 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																															tmp95642 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																															tmp95643 := Call(__e, tmp95642, V2687)

																																															tmp95644 := Call(__e, tmp95641, tmp95643, V2783)

																																															__e.TailApply(tmp95594, tmp95644)
																																															return

																																														}, 1)

																																														tmp95645 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																														tmp95646 := Call(__e, tmp95645, V2687)

																																														__e.TailApply(tmp95593, tmp95646)
																																														return

																																													} else {
																																														tmp95591 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																														tmp95592 := Call(__e, tmp95591, V2687)

																																														if True == tmp95592 {
																																															tmp95563 := MakeNative(func(__e *ControlFlow) {
																																																Direction := __e.Get(1)
																																																_ = Direction
																																																tmp95564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																tmp95565 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																tmp95566 := Call(__e, tmp95565, Direction, Nil)

																																																tmp95567 := Call(__e, tmp95564, V2687, tmp95566, V2783)

																																																_ = tmp95567

																																																tmp95568 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp95569 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																	tmp95570 := Call(__e, tmp95569, V2687, V2783)

																																																	_ = tmp95570

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp95571 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																tmp95572 := Call(__e, tmp95571)

																																																_ = tmp95572

																																																tmp95573 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																																tmp95574 := MakeNative(func(__e *ControlFlow) {
																																																	tmp95575 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																																	tmp95576 := MakeNative(func(__e *ControlFlow) {
																																																		tmp95577 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																																		tmp95578 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																																		tmp95579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp95580 := Call(__e, tmp95579, Direction, V2783)

																																																		tmp95581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp95582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp95583 := Call(__e, tmp95582, symout, Nil)

																																																		tmp95584 := Call(__e, tmp95581, symin, tmp95583)

																																																		tmp95585 := Call(__e, tmp95578, tmp95580, tmp95584)

																																																		tmp95586 := MakeNative(func(__e *ControlFlow) {
																																																			tmp95587 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																			__e.TailApply(tmp95587, FileName, symstring, V2782, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp95577, tmp95585, V2783, tmp95586)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp95575, Throwcontrol, V2783, tmp95576)
																																																	return

																																																}, 0)

																																																tmp95588 := Call(__e, tmp95573, Direction, Direction2611, V2783, tmp95574)

																																																__e.TailApply(tmp95568, tmp95588)
																																																return

																																															}, 1)

																																															tmp95589 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																															tmp95590 := Call(__e, tmp95589, V2783)

																																															__e.TailApply(tmp95563, tmp95590)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}

																																												}, 1)

																																												tmp95649 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp95650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																												tmp95651 := Call(__e, tmp95650, V2683)

																																												tmp95652 := Call(__e, tmp95649, tmp95651, V2783)

																																												tmp95653 := Call(__e, tmp95560, tmp95652)

																																												__e.TailApply(tmp95557, tmp95653)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}

																																									}, 1)

																																									tmp95751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp95752 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																									tmp95753 := Call(__e, tmp95752, V2683)

																																									tmp95754 := Call(__e, tmp95751, tmp95753, V2783)

																																									__e.TailApply(tmp95552, tmp95754)
																																									return

																																								} else {
																																									tmp95550 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp95551 := Call(__e, tmp95550, V2683)

																																									if True == tmp95551 {
																																										tmp95520 := MakeNative(func(__e *ControlFlow) {
																																											Direction := __e.Get(1)
																																											_ = Direction
																																											tmp95521 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																											tmp95522 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp95523 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp95524 := Call(__e, tmp95523, Direction, Nil)

																																											tmp95525 := Call(__e, tmp95522, symstream, tmp95524)

																																											tmp95526 := Call(__e, tmp95521, V2683, tmp95525, V2783)

																																											_ = tmp95526

																																											tmp95527 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp95528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																												tmp95529 := Call(__e, tmp95528, V2683, V2783)

																																												_ = tmp95529

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp95530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp95531 := Call(__e, tmp95530)

																																											_ = tmp95531

																																											tmp95532 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																																											tmp95533 := MakeNative(func(__e *ControlFlow) {
																																												tmp95534 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																																												tmp95535 := MakeNative(func(__e *ControlFlow) {
																																													tmp95536 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

																																													tmp95537 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

																																													tmp95538 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95539 := Call(__e, tmp95538, Direction, V2783)

																																													tmp95540 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																													tmp95541 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																													tmp95542 := Call(__e, tmp95541, symout, Nil)

																																													tmp95543 := Call(__e, tmp95540, symin, tmp95542)

																																													tmp95544 := Call(__e, tmp95537, tmp95539, tmp95543)

																																													tmp95545 := MakeNative(func(__e *ControlFlow) {
																																														tmp95546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																														__e.TailApply(tmp95546, FileName, symstring, V2782, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95536, tmp95544, V2783, tmp95545)
																																													return

																																												}, 0)

																																												__e.TailApply(tmp95534, Throwcontrol, V2783, tmp95535)
																																												return

																																											}, 0)

																																											tmp95547 := Call(__e, tmp95532, Direction, Direction2611, V2783, tmp95533)

																																											__e.TailApply(tmp95527, tmp95547)
																																											return

																																										}, 1)

																																										tmp95548 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																										tmp95549 := Call(__e, tmp95548, V2783)

																																										__e.TailApply(tmp95520, tmp95549)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp95757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp95758 := Call(__e, tmp95757, V2781, V2783)

																																							__e.TailApply(tmp95517, tmp95758)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp95761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp95762 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp95763 := Call(__e, tmp95762, V2681)

																																					tmp95764 := Call(__e, tmp95761, tmp95763, V2783)

																																					__e.TailApply(tmp95515, tmp95764)
																																					return

																																				}, 1)

																																				tmp95765 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																				tmp95766 := Call(__e, tmp95765, V2681)

																																				__e.TailApply(tmp95514, tmp95766)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp95769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp95770 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp95771 := Call(__e, tmp95770, V2680)

																																		tmp95772 := Call(__e, tmp95769, tmp95771, V2783)

																																		__e.TailApply(tmp95512, tmp95772)
																																		return

																																	}, 1)

																																	tmp95773 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																	tmp95774 := Call(__e, tmp95773, V2680)

																																	__e.TailApply(tmp95511, tmp95774)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp95777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp95778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp95779 := Call(__e, tmp95778, V2678)

																															tmp95780 := Call(__e, tmp95777, tmp95779, V2783)

																															__e.TailApply(tmp95509, tmp95780)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp95783 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp95784 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																													tmp95785 := Call(__e, tmp95784, V2678)

																													tmp95786 := Call(__e, tmp95783, tmp95785, V2783)

																													__e.TailApply(tmp95507, tmp95786)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp95789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp95790 := Call(__e, tmp95789, V2780, V2783)

																											tmp95791 := Call(__e, tmp95505, tmp95790)

																											__e.TailApply(tmp95123, tmp95791)
																											return

																										} else {
																											__e.Return(Case)
																											return
																										}

																									}, 1)

																									tmp95794 := MakeNative(func(__e *ControlFlow) {
																										V2672 := __e.Get(1)
																										_ = V2672
																										tmp95881 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																										tmp95882 := Call(__e, tmp95881, V2672)

																										if True == tmp95882 {
																											tmp95796 := MakeNative(func(__e *ControlFlow) {
																												V2673 := __e.Get(1)
																												_ = V2673
																												tmp95875 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																												tmp95876 := Call(__e, tmp95875, symlet, V2673)

																												if True == tmp95876 {
																													tmp95798 := MakeNative(func(__e *ControlFlow) {
																														V2674 := __e.Get(1)
																														_ = V2674
																														tmp95869 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																														tmp95870 := Call(__e, tmp95869, V2674)

																														if True == tmp95870 {
																															tmp95800 := MakeNative(func(__e *ControlFlow) {
																																X := __e.Get(1)
																																_ = X
																																tmp95801 := MakeNative(func(__e *ControlFlow) {
																																	V2675 := __e.Get(1)
																																	_ = V2675
																																	tmp95861 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																	tmp95862 := Call(__e, tmp95861, V2675)

																																	if True == tmp95862 {
																																		tmp95803 := MakeNative(func(__e *ControlFlow) {
																																			Y := __e.Get(1)
																																			_ = Y
																																			tmp95804 := MakeNative(func(__e *ControlFlow) {
																																				V2676 := __e.Get(1)
																																				_ = V2676
																																				tmp95853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																				tmp95854 := Call(__e, tmp95853, V2676)

																																				if True == tmp95854 {
																																					tmp95806 := MakeNative(func(__e *ControlFlow) {
																																						Z := __e.Get(1)
																																						_ = Z
																																						tmp95807 := MakeNative(func(__e *ControlFlow) {
																																							V2677 := __e.Get(1)
																																							_ = V2677
																																							tmp95845 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp95846 := Call(__e, tmp95845, Nil, V2677)

																																							if True == tmp95846 {
																																								tmp95809 := MakeNative(func(__e *ControlFlow) {
																																									W := __e.Get(1)
																																									_ = W
																																									tmp95810 := MakeNative(func(__e *ControlFlow) {
																																										X_e_e := __e.Get(1)
																																										_ = X_e_e
																																										tmp95811 := MakeNative(func(__e *ControlFlow) {
																																											B := __e.Get(1)
																																											_ = B
																																											tmp95812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp95813 := Call(__e, tmp95812)

																																											_ = tmp95813

																																											tmp95814 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											tmp95815 := MakeNative(func(__e *ControlFlow) {
																																												tmp95816 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																												tmp95817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																												tmp95818 := Call(__e, tmp95817)

																																												tmp95819 := MakeNative(func(__e *ControlFlow) {
																																													tmp95820 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																													tmp95821 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																													tmp95822 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95823 := Call(__e, tmp95822, X_e_e, V2783)

																																													tmp95824 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95825 := Call(__e, tmp95824, X, V2783)

																																													tmp95826 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95827 := Call(__e, tmp95826, Z, V2783)

																																													tmp95828 := Call(__e, tmp95821, tmp95823, tmp95825, tmp95827)

																																													tmp95829 := MakeNative(func(__e *ControlFlow) {
																																														tmp95830 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																														tmp95831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95834 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95835 := Call(__e, tmp95834, B, Nil)

																																														tmp95836 := Call(__e, tmp95833, sym_h, tmp95835)

																																														tmp95837 := Call(__e, tmp95832, X_e_e, tmp95836)

																																														tmp95838 := Call(__e, tmp95831, tmp95837, V2782)

																																														__e.TailApply(tmp95830, W, V2781, tmp95838, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95820, W, tmp95828, V2783, tmp95829)
																																													return

																																												}, 0)

																																												__e.TailApply(tmp95816, X_e_e, tmp95818, V2783, tmp95819)
																																												return

																																											}, 0)

																																											__e.TailApply(tmp95814, Y, B, V2782, V2783, tmp95815)
																																											return

																																										}, 1)

																																										tmp95839 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																										tmp95840 := Call(__e, tmp95839, V2783)

																																										__e.TailApply(tmp95811, tmp95840)
																																										return

																																									}, 1)

																																									tmp95841 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																									tmp95842 := Call(__e, tmp95841, V2783)

																																									__e.TailApply(tmp95810, tmp95842)
																																									return

																																								}, 1)

																																								tmp95843 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																								tmp95844 := Call(__e, tmp95843, V2783)

																																								__e.TailApply(tmp95809, tmp95844)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp95847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp95848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp95849 := Call(__e, tmp95848, V2676)

																																						tmp95850 := Call(__e, tmp95847, tmp95849, V2783)

																																						__e.TailApply(tmp95807, tmp95850)
																																						return

																																					}, 1)

																																					tmp95851 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																					tmp95852 := Call(__e, tmp95851, V2676)

																																					__e.TailApply(tmp95806, tmp95852)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp95855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp95856 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp95857 := Call(__e, tmp95856, V2675)

																																			tmp95858 := Call(__e, tmp95855, tmp95857, V2783)

																																			__e.TailApply(tmp95804, tmp95858)
																																			return

																																		}, 1)

																																		tmp95859 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																		tmp95860 := Call(__e, tmp95859, V2675)

																																		__e.TailApply(tmp95803, tmp95860)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp95863 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp95864 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp95865 := Call(__e, tmp95864, V2674)

																																tmp95866 := Call(__e, tmp95863, tmp95865, V2783)

																																__e.TailApply(tmp95801, tmp95866)
																																return

																															}, 1)

																															tmp95867 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																															tmp95868 := Call(__e, tmp95867, V2674)

																															__e.TailApply(tmp95800, tmp95868)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp95871 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp95872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																													tmp95873 := Call(__e, tmp95872, V2672)

																													tmp95874 := Call(__e, tmp95871, tmp95873, V2783)

																													__e.TailApply(tmp95798, tmp95874)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp95877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp95878 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																											tmp95879 := Call(__e, tmp95878, V2672)

																											tmp95880 := Call(__e, tmp95877, tmp95879, V2783)

																											__e.TailApply(tmp95796, tmp95880)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp95883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																									tmp95884 := Call(__e, tmp95883, V2780, V2783)

																									tmp95885 := Call(__e, tmp95794, tmp95884)

																									__e.TailApply(tmp95121, tmp95885)
																									return

																								} else {
																									__e.Return(Case)
																									return
																								}

																							}, 1)

																							tmp95888 := MakeNative(func(__e *ControlFlow) {
																								V2660 := __e.Get(1)
																								_ = V2660
																								tmp96331 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																								tmp96332 := Call(__e, tmp96331, V2660)

																								if True == tmp96332 {
																									tmp95890 := MakeNative(func(__e *ControlFlow) {
																										V2661 := __e.Get(1)
																										_ = V2661
																										tmp96325 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																										tmp96326 := Call(__e, tmp96325, symlambda, V2661)

																										if True == tmp96326 {
																											tmp95892 := MakeNative(func(__e *ControlFlow) {
																												V2662 := __e.Get(1)
																												_ = V2662
																												tmp96319 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																												tmp96320 := Call(__e, tmp96319, V2662)

																												if True == tmp96320 {
																													tmp95894 := MakeNative(func(__e *ControlFlow) {
																														X := __e.Get(1)
																														_ = X
																														tmp95895 := MakeNative(func(__e *ControlFlow) {
																															V2663 := __e.Get(1)
																															_ = V2663
																															tmp96311 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																															tmp96312 := Call(__e, tmp96311, V2663)

																															if True == tmp96312 {
																																tmp95897 := MakeNative(func(__e *ControlFlow) {
																																	Y := __e.Get(1)
																																	_ = Y
																																	tmp95898 := MakeNative(func(__e *ControlFlow) {
																																		V2664 := __e.Get(1)
																																		_ = V2664
																																		tmp96303 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp96304 := Call(__e, tmp96303, Nil, V2664)

																																		if True == tmp96304 {
																																			tmp95900 := MakeNative(func(__e *ControlFlow) {
																																				V2665 := __e.Get(1)
																																				_ = V2665
																																				tmp96299 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																				tmp96300 := Call(__e, tmp96299, V2665)

																																				if True == tmp96300 {
																																					tmp95954 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp95955 := MakeNative(func(__e *ControlFlow) {
																																							V2666 := __e.Get(1)
																																							_ = V2666
																																							tmp96291 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																							tmp96292 := Call(__e, tmp96291, V2666)

																																							if True == tmp96292 {
																																								tmp96004 := MakeNative(func(__e *ControlFlow) {
																																									V2667 := __e.Get(1)
																																									_ = V2667
																																									tmp96285 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																									tmp96286 := Call(__e, tmp96285, sym_1_1_6, V2667)

																																									if True == tmp96286 {
																																										tmp96150 := MakeNative(func(__e *ControlFlow) {
																																											V2668 := __e.Get(1)
																																											_ = V2668
																																											tmp96279 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																											tmp96280 := Call(__e, tmp96279, V2668)

																																											if True == tmp96280 {
																																												tmp96197 := MakeNative(func(__e *ControlFlow) {
																																													B := __e.Get(1)
																																													_ = B
																																													tmp96198 := MakeNative(func(__e *ControlFlow) {
																																														V2669 := __e.Get(1)
																																														_ = V2669
																																														tmp96271 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																														tmp96272 := Call(__e, tmp96271, Nil, V2669)

																																														if True == tmp96272 {
																																															tmp96240 := MakeNative(func(__e *ControlFlow) {
																																																Z := __e.Get(1)
																																																_ = Z
																																																tmp96241 := MakeNative(func(__e *ControlFlow) {
																																																	X_e_e := __e.Get(1)
																																																	_ = X_e_e
																																																	tmp96242 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																	tmp96243 := Call(__e, tmp96242)

																																																	_ = tmp96243

																																																	tmp96244 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																	tmp96245 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																	tmp96246 := Call(__e, tmp96245)

																																																	tmp96247 := MakeNative(func(__e *ControlFlow) {
																																																		tmp96248 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																		tmp96249 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																		tmp96250 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96251 := Call(__e, tmp96250, X_e_e, V2783)

																																																		tmp96252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96253 := Call(__e, tmp96252, X, V2783)

																																																		tmp96254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96255 := Call(__e, tmp96254, Y, V2783)

																																																		tmp96256 := Call(__e, tmp96249, tmp96251, tmp96253, tmp96255)

																																																		tmp96257 := MakeNative(func(__e *ControlFlow) {
																																																			tmp96258 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																			tmp96259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96263 := Call(__e, tmp96262, A, Nil)

																																																			tmp96264 := Call(__e, tmp96261, sym_h, tmp96263)

																																																			tmp96265 := Call(__e, tmp96260, X_e_e, tmp96264)

																																																			tmp96266 := Call(__e, tmp96259, tmp96265, V2782)

																																																			__e.TailApply(tmp96258, Z, B, tmp96266, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp96248, Z, tmp96256, V2783, tmp96257)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp96244, X_e_e, tmp96246, V2783, tmp96247)
																																																	return

																																																}, 1)

																																																tmp96267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																tmp96268 := Call(__e, tmp96267, V2783)

																																																__e.TailApply(tmp96241, tmp96268)
																																																return

																																															}, 1)

																																															tmp96269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																															tmp96270 := Call(__e, tmp96269, V2783)

																																															__e.TailApply(tmp96240, tmp96270)
																																															return

																																														} else {
																																															tmp96238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																															tmp96239 := Call(__e, tmp96238, V2669)

																																															if True == tmp96239 {
																																																tmp96201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																tmp96202 := Call(__e, tmp96201, V2669, Nil, V2783)

																																																_ = tmp96202

																																																tmp96203 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp96204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																	tmp96205 := Call(__e, tmp96204, V2669, V2783)

																																																	_ = tmp96205

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp96206 := MakeNative(func(__e *ControlFlow) {
																																																	Z := __e.Get(1)
																																																	_ = Z
																																																	tmp96207 := MakeNative(func(__e *ControlFlow) {
																																																		X_e_e := __e.Get(1)
																																																		_ = X_e_e
																																																		tmp96208 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																		tmp96209 := Call(__e, tmp96208)

																																																		_ = tmp96209

																																																		tmp96210 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																		tmp96211 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																		tmp96212 := Call(__e, tmp96211)

																																																		tmp96213 := MakeNative(func(__e *ControlFlow) {
																																																			tmp96214 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																			tmp96215 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																			tmp96216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96217 := Call(__e, tmp96216, X_e_e, V2783)

																																																			tmp96218 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96219 := Call(__e, tmp96218, X, V2783)

																																																			tmp96220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96221 := Call(__e, tmp96220, Y, V2783)

																																																			tmp96222 := Call(__e, tmp96215, tmp96217, tmp96219, tmp96221)

																																																			tmp96223 := MakeNative(func(__e *ControlFlow) {
																																																				tmp96224 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																				tmp96225 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96226 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96227 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96228 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96229 := Call(__e, tmp96228, A, Nil)

																																																				tmp96230 := Call(__e, tmp96227, sym_h, tmp96229)

																																																				tmp96231 := Call(__e, tmp96226, X_e_e, tmp96230)

																																																				tmp96232 := Call(__e, tmp96225, tmp96231, V2782)

																																																				__e.TailApply(tmp96224, Z, B, tmp96232, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp96214, Z, tmp96222, V2783, tmp96223)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp96210, X_e_e, tmp96212, V2783, tmp96213)
																																																		return

																																																	}, 1)

																																																	tmp96233 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																	tmp96234 := Call(__e, tmp96233, V2783)

																																																	__e.TailApply(tmp96207, tmp96234)
																																																	return

																																																}, 1)

																																																tmp96235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																tmp96236 := Call(__e, tmp96235, V2783)

																																																tmp96237 := Call(__e, tmp96206, tmp96236)

																																																__e.TailApply(tmp96203, tmp96237)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp96273 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp96274 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																													tmp96275 := Call(__e, tmp96274, V2668)

																																													tmp96276 := Call(__e, tmp96273, tmp96275, V2783)

																																													__e.TailApply(tmp96198, tmp96276)
																																													return

																																												}, 1)

																																												tmp96277 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																												tmp96278 := Call(__e, tmp96277, V2668)

																																												__e.TailApply(tmp96197, tmp96278)
																																												return

																																											} else {
																																												tmp96195 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																												tmp96196 := Call(__e, tmp96195, V2668)

																																												if True == tmp96196 {
																																													tmp96153 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp96154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																														tmp96155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp96156 := Call(__e, tmp96155, B, Nil)

																																														tmp96157 := Call(__e, tmp96154, V2668, tmp96156, V2783)

																																														_ = tmp96157

																																														tmp96158 := MakeNative(func(__e *ControlFlow) {
																																															Result := __e.Get(1)
																																															_ = Result
																																															tmp96159 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																															tmp96160 := Call(__e, tmp96159, V2668, V2783)

																																															_ = tmp96160

																																															__e.Return(Result)
																																															return

																																														}, 1)

																																														tmp96161 := MakeNative(func(__e *ControlFlow) {
																																															Z := __e.Get(1)
																																															_ = Z
																																															tmp96162 := MakeNative(func(__e *ControlFlow) {
																																																X_e_e := __e.Get(1)
																																																_ = X_e_e
																																																tmp96163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																tmp96164 := Call(__e, tmp96163)

																																																_ = tmp96164

																																																tmp96165 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																tmp96166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																tmp96167 := Call(__e, tmp96166)

																																																tmp96168 := MakeNative(func(__e *ControlFlow) {
																																																	tmp96169 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																	tmp96170 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																	tmp96171 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																	tmp96172 := Call(__e, tmp96171, X_e_e, V2783)

																																																	tmp96173 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																	tmp96174 := Call(__e, tmp96173, X, V2783)

																																																	tmp96175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																	tmp96176 := Call(__e, tmp96175, Y, V2783)

																																																	tmp96177 := Call(__e, tmp96170, tmp96172, tmp96174, tmp96176)

																																																	tmp96178 := MakeNative(func(__e *ControlFlow) {
																																																		tmp96179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																		tmp96180 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp96181 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp96182 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp96183 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																		tmp96184 := Call(__e, tmp96183, A, Nil)

																																																		tmp96185 := Call(__e, tmp96182, sym_h, tmp96184)

																																																		tmp96186 := Call(__e, tmp96181, X_e_e, tmp96185)

																																																		tmp96187 := Call(__e, tmp96180, tmp96186, V2782)

																																																		__e.TailApply(tmp96179, Z, B, tmp96187, V2783, V2784)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp96169, Z, tmp96177, V2783, tmp96178)
																																																	return

																																																}, 0)

																																																__e.TailApply(tmp96165, X_e_e, tmp96167, V2783, tmp96168)
																																																return

																																															}, 1)

																																															tmp96188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																															tmp96189 := Call(__e, tmp96188, V2783)

																																															__e.TailApply(tmp96162, tmp96189)
																																															return

																																														}, 1)

																																														tmp96190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																														tmp96191 := Call(__e, tmp96190, V2783)

																																														tmp96192 := Call(__e, tmp96161, tmp96191)

																																														__e.TailApply(tmp96158, tmp96192)
																																														return

																																													}, 1)

																																													tmp96193 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																													tmp96194 := Call(__e, tmp96193, V2783)

																																													__e.TailApply(tmp96153, tmp96194)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp96281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp96282 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																										tmp96283 := Call(__e, tmp96282, V2666)

																																										tmp96284 := Call(__e, tmp96281, tmp96283, V2783)

																																										__e.TailApply(tmp96150, tmp96284)
																																										return

																																									} else {
																																										tmp96148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																										tmp96149 := Call(__e, tmp96148, V2667)

																																										if True == tmp96149 {
																																											tmp96007 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																											tmp96008 := Call(__e, tmp96007, V2667, sym_1_1_6, V2783)

																																											_ = tmp96008

																																											tmp96009 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp96010 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																												tmp96011 := Call(__e, tmp96010, V2667, V2783)

																																												_ = tmp96011

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp96012 := MakeNative(func(__e *ControlFlow) {
																																												V2670 := __e.Get(1)
																																												_ = V2670
																																												tmp96141 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																												tmp96142 := Call(__e, tmp96141, V2670)

																																												if True == tmp96142 {
																																													tmp96059 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp96060 := MakeNative(func(__e *ControlFlow) {
																																															V2671 := __e.Get(1)
																																															_ = V2671
																																															tmp96133 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																															tmp96134 := Call(__e, tmp96133, Nil, V2671)

																																															if True == tmp96134 {
																																																tmp96102 := MakeNative(func(__e *ControlFlow) {
																																																	Z := __e.Get(1)
																																																	_ = Z
																																																	tmp96103 := MakeNative(func(__e *ControlFlow) {
																																																		X_e_e := __e.Get(1)
																																																		_ = X_e_e
																																																		tmp96104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																		tmp96105 := Call(__e, tmp96104)

																																																		_ = tmp96105

																																																		tmp96106 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																		tmp96107 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																		tmp96108 := Call(__e, tmp96107)

																																																		tmp96109 := MakeNative(func(__e *ControlFlow) {
																																																			tmp96110 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																			tmp96111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																			tmp96112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96113 := Call(__e, tmp96112, X_e_e, V2783)

																																																			tmp96114 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96115 := Call(__e, tmp96114, X, V2783)

																																																			tmp96116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																			tmp96117 := Call(__e, tmp96116, Y, V2783)

																																																			tmp96118 := Call(__e, tmp96111, tmp96113, tmp96115, tmp96117)

																																																			tmp96119 := MakeNative(func(__e *ControlFlow) {
																																																				tmp96120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																				tmp96121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96123 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96124 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																				tmp96125 := Call(__e, tmp96124, A, Nil)

																																																				tmp96126 := Call(__e, tmp96123, sym_h, tmp96125)

																																																				tmp96127 := Call(__e, tmp96122, X_e_e, tmp96126)

																																																				tmp96128 := Call(__e, tmp96121, tmp96127, V2782)

																																																				__e.TailApply(tmp96120, Z, B, tmp96128, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp96110, Z, tmp96118, V2783, tmp96119)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp96106, X_e_e, tmp96108, V2783, tmp96109)
																																																		return

																																																	}, 1)

																																																	tmp96129 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																	tmp96130 := Call(__e, tmp96129, V2783)

																																																	__e.TailApply(tmp96103, tmp96130)
																																																	return

																																																}, 1)

																																																tmp96131 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																tmp96132 := Call(__e, tmp96131, V2783)

																																																__e.TailApply(tmp96102, tmp96132)
																																																return

																																															} else {
																																																tmp96100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																																tmp96101 := Call(__e, tmp96100, V2671)

																																																if True == tmp96101 {
																																																	tmp96063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																																	tmp96064 := Call(__e, tmp96063, V2671, Nil, V2783)

																																																	_ = tmp96064

																																																	tmp96065 := MakeNative(func(__e *ControlFlow) {
																																																		Result := __e.Get(1)
																																																		_ = Result
																																																		tmp96066 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																		tmp96067 := Call(__e, tmp96066, V2671, V2783)

																																																		_ = tmp96067

																																																		__e.Return(Result)
																																																		return

																																																	}, 1)

																																																	tmp96068 := MakeNative(func(__e *ControlFlow) {
																																																		Z := __e.Get(1)
																																																		_ = Z
																																																		tmp96069 := MakeNative(func(__e *ControlFlow) {
																																																			X_e_e := __e.Get(1)
																																																			_ = X_e_e
																																																			tmp96070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																			tmp96071 := Call(__e, tmp96070)

																																																			_ = tmp96071

																																																			tmp96072 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																			tmp96073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																			tmp96074 := Call(__e, tmp96073)

																																																			tmp96075 := MakeNative(func(__e *ControlFlow) {
																																																				tmp96076 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																				tmp96077 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																				tmp96078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																				tmp96079 := Call(__e, tmp96078, X_e_e, V2783)

																																																				tmp96080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																				tmp96081 := Call(__e, tmp96080, X, V2783)

																																																				tmp96082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																				tmp96083 := Call(__e, tmp96082, Y, V2783)

																																																				tmp96084 := Call(__e, tmp96077, tmp96079, tmp96081, tmp96083)

																																																				tmp96085 := MakeNative(func(__e *ControlFlow) {
																																																					tmp96086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																					tmp96087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																					tmp96088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																					tmp96089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																					tmp96090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																					tmp96091 := Call(__e, tmp96090, A, Nil)

																																																					tmp96092 := Call(__e, tmp96089, sym_h, tmp96091)

																																																					tmp96093 := Call(__e, tmp96088, X_e_e, tmp96092)

																																																					tmp96094 := Call(__e, tmp96087, tmp96093, V2782)

																																																					__e.TailApply(tmp96086, Z, B, tmp96094, V2783, V2784)
																																																					return

																																																				}, 0)

																																																				__e.TailApply(tmp96076, Z, tmp96084, V2783, tmp96085)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(tmp96072, X_e_e, tmp96074, V2783, tmp96075)
																																																			return

																																																		}, 1)

																																																		tmp96095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																		tmp96096 := Call(__e, tmp96095, V2783)

																																																		__e.TailApply(tmp96069, tmp96096)
																																																		return

																																																	}, 1)

																																																	tmp96097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																	tmp96098 := Call(__e, tmp96097, V2783)

																																																	tmp96099 := Call(__e, tmp96068, tmp96098)

																																																	__e.TailApply(tmp96065, tmp96099)
																																																	return

																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp96135 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																														tmp96136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																														tmp96137 := Call(__e, tmp96136, V2670)

																																														tmp96138 := Call(__e, tmp96135, tmp96137, V2783)

																																														__e.TailApply(tmp96060, tmp96138)
																																														return

																																													}, 1)

																																													tmp96139 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																													tmp96140 := Call(__e, tmp96139, V2670)

																																													__e.TailApply(tmp96059, tmp96140)
																																													return

																																												} else {
																																													tmp96057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																													tmp96058 := Call(__e, tmp96057, V2670)

																																													if True == tmp96058 {
																																														tmp96015 := MakeNative(func(__e *ControlFlow) {
																																															B := __e.Get(1)
																																															_ = B
																																															tmp96016 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																															tmp96017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																															tmp96018 := Call(__e, tmp96017, B, Nil)

																																															tmp96019 := Call(__e, tmp96016, V2670, tmp96018, V2783)

																																															_ = tmp96019

																																															tmp96020 := MakeNative(func(__e *ControlFlow) {
																																																Result := __e.Get(1)
																																																_ = Result
																																																tmp96021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																																tmp96022 := Call(__e, tmp96021, V2670, V2783)

																																																_ = tmp96022

																																																__e.Return(Result)
																																																return

																																															}, 1)

																																															tmp96023 := MakeNative(func(__e *ControlFlow) {
																																																Z := __e.Get(1)
																																																_ = Z
																																																tmp96024 := MakeNative(func(__e *ControlFlow) {
																																																	X_e_e := __e.Get(1)
																																																	_ = X_e_e
																																																	tmp96025 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																																	tmp96026 := Call(__e, tmp96025)

																																																	_ = tmp96026

																																																	tmp96027 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																	tmp96028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																																	tmp96029 := Call(__e, tmp96028)

																																																	tmp96030 := MakeNative(func(__e *ControlFlow) {
																																																		tmp96031 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																																		tmp96032 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																																		tmp96033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96034 := Call(__e, tmp96033, X_e_e, V2783)

																																																		tmp96035 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96036 := Call(__e, tmp96035, X, V2783)

																																																		tmp96037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																																		tmp96038 := Call(__e, tmp96037, Y, V2783)

																																																		tmp96039 := Call(__e, tmp96032, tmp96034, tmp96036, tmp96038)

																																																		tmp96040 := MakeNative(func(__e *ControlFlow) {
																																																			tmp96041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																																			tmp96042 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96043 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																																			tmp96046 := Call(__e, tmp96045, A, Nil)

																																																			tmp96047 := Call(__e, tmp96044, sym_h, tmp96046)

																																																			tmp96048 := Call(__e, tmp96043, X_e_e, tmp96047)

																																																			tmp96049 := Call(__e, tmp96042, tmp96048, V2782)

																																																			__e.TailApply(tmp96041, Z, B, tmp96049, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(tmp96031, Z, tmp96039, V2783, tmp96040)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(tmp96027, X_e_e, tmp96029, V2783, tmp96030)
																																																	return

																																																}, 1)

																																																tmp96050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																																tmp96051 := Call(__e, tmp96050, V2783)

																																																__e.TailApply(tmp96024, tmp96051)
																																																return

																																															}, 1)

																																															tmp96052 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																															tmp96053 := Call(__e, tmp96052, V2783)

																																															tmp96054 := Call(__e, tmp96023, tmp96053)

																																															__e.TailApply(tmp96020, tmp96054)
																																															return

																																														}, 1)

																																														tmp96055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																														tmp96056 := Call(__e, tmp96055, V2783)

																																														__e.TailApply(tmp96015, tmp96056)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp96143 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp96144 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																											tmp96145 := Call(__e, tmp96144, V2666)

																																											tmp96146 := Call(__e, tmp96143, tmp96145, V2783)

																																											tmp96147 := Call(__e, tmp96012, tmp96146)

																																											__e.TailApply(tmp96009, tmp96147)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp96287 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp96288 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																								tmp96289 := Call(__e, tmp96288, V2666)

																																								tmp96290 := Call(__e, tmp96287, tmp96289, V2783)

																																								__e.TailApply(tmp96004, tmp96290)
																																								return

																																							} else {
																																								tmp96002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp96003 := Call(__e, tmp96002, V2666)

																																								if True == tmp96003 {
																																									tmp95958 := MakeNative(func(__e *ControlFlow) {
																																										B := __e.Get(1)
																																										_ = B
																																										tmp95959 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp95960 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp95961 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp95962 := Call(__e, tmp95961, B, Nil)

																																										tmp95963 := Call(__e, tmp95960, sym_1_1_6, tmp95962)

																																										tmp95964 := Call(__e, tmp95959, V2666, tmp95963, V2783)

																																										_ = tmp95964

																																										tmp95965 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp95966 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp95967 := Call(__e, tmp95966, V2666, V2783)

																																											_ = tmp95967

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp95968 := MakeNative(func(__e *ControlFlow) {
																																											Z := __e.Get(1)
																																											_ = Z
																																											tmp95969 := MakeNative(func(__e *ControlFlow) {
																																												X_e_e := __e.Get(1)
																																												_ = X_e_e
																																												tmp95970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																												tmp95971 := Call(__e, tmp95970)

																																												_ = tmp95971

																																												tmp95972 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																												tmp95973 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																												tmp95974 := Call(__e, tmp95973)

																																												tmp95975 := MakeNative(func(__e *ControlFlow) {
																																													tmp95976 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																													tmp95977 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																													tmp95978 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95979 := Call(__e, tmp95978, X_e_e, V2783)

																																													tmp95980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95981 := Call(__e, tmp95980, X, V2783)

																																													tmp95982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																													tmp95983 := Call(__e, tmp95982, Y, V2783)

																																													tmp95984 := Call(__e, tmp95977, tmp95979, tmp95981, tmp95983)

																																													tmp95985 := MakeNative(func(__e *ControlFlow) {
																																														tmp95986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																														tmp95987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95988 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95989 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95990 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																														tmp95991 := Call(__e, tmp95990, A, Nil)

																																														tmp95992 := Call(__e, tmp95989, sym_h, tmp95991)

																																														tmp95993 := Call(__e, tmp95988, X_e_e, tmp95992)

																																														tmp95994 := Call(__e, tmp95987, tmp95993, V2782)

																																														__e.TailApply(tmp95986, Z, B, tmp95994, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(tmp95976, Z, tmp95984, V2783, tmp95985)
																																													return

																																												}, 0)

																																												__e.TailApply(tmp95972, X_e_e, tmp95974, V2783, tmp95975)
																																												return

																																											}, 1)

																																											tmp95995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																											tmp95996 := Call(__e, tmp95995, V2783)

																																											__e.TailApply(tmp95969, tmp95996)
																																											return

																																										}, 1)

																																										tmp95997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																										tmp95998 := Call(__e, tmp95997, V2783)

																																										tmp95999 := Call(__e, tmp95968, tmp95998)

																																										__e.TailApply(tmp95965, tmp95999)
																																										return

																																									}, 1)

																																									tmp96000 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																									tmp96001 := Call(__e, tmp96000, V2783)

																																									__e.TailApply(tmp95958, tmp96001)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp96293 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp96294 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp96295 := Call(__e, tmp96294, V2665)

																																						tmp96296 := Call(__e, tmp96293, tmp96295, V2783)

																																						__e.TailApply(tmp95955, tmp96296)
																																						return

																																					}, 1)

																																					tmp96297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																					tmp96298 := Call(__e, tmp96297, V2665)

																																					__e.TailApply(tmp95954, tmp96298)
																																					return

																																				} else {
																																					tmp95952 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																					tmp95953 := Call(__e, tmp95952, V2665)

																																					if True == tmp95953 {
																																						tmp95903 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp95904 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp95905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp95906 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp95907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp95908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp95909 := Call(__e, tmp95908, B, Nil)

																																								tmp95910 := Call(__e, tmp95907, sym_1_1_6, tmp95909)

																																								tmp95911 := Call(__e, tmp95906, A, tmp95910)

																																								tmp95912 := Call(__e, tmp95905, V2665, tmp95911, V2783)

																																								_ = tmp95912

																																								tmp95913 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp95914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp95915 := Call(__e, tmp95914, V2665, V2783)

																																									_ = tmp95915

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp95916 := MakeNative(func(__e *ControlFlow) {
																																									Z := __e.Get(1)
																																									_ = Z
																																									tmp95917 := MakeNative(func(__e *ControlFlow) {
																																										X_e_e := __e.Get(1)
																																										_ = X_e_e
																																										tmp95918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp95919 := Call(__e, tmp95918)

																																										_ = tmp95919

																																										tmp95920 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp95921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholder)

																																										tmp95922 := Call(__e, tmp95921)

																																										tmp95923 := MakeNative(func(__e *ControlFlow) {
																																											tmp95924 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																											tmp95925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ebr)

																																											tmp95926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95927 := Call(__e, tmp95926, X_e_e, V2783)

																																											tmp95928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95929 := Call(__e, tmp95928, X, V2783)

																																											tmp95930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp95931 := Call(__e, tmp95930, Y, V2783)

																																											tmp95932 := Call(__e, tmp95925, tmp95927, tmp95929, tmp95931)

																																											tmp95933 := MakeNative(func(__e *ControlFlow) {
																																												tmp95934 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																												tmp95935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp95936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp95937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp95938 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp95939 := Call(__e, tmp95938, A, Nil)

																																												tmp95940 := Call(__e, tmp95937, sym_h, tmp95939)

																																												tmp95941 := Call(__e, tmp95936, X_e_e, tmp95940)

																																												tmp95942 := Call(__e, tmp95935, tmp95941, V2782)

																																												__e.TailApply(tmp95934, Z, B, tmp95942, V2783, V2784)
																																												return

																																											}, 0)

																																											__e.TailApply(tmp95924, Z, tmp95932, V2783, tmp95933)
																																											return

																																										}, 0)

																																										__e.TailApply(tmp95920, X_e_e, tmp95922, V2783, tmp95923)
																																										return

																																									}, 1)

																																									tmp95943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																									tmp95944 := Call(__e, tmp95943, V2783)

																																									__e.TailApply(tmp95917, tmp95944)
																																									return

																																								}, 1)

																																								tmp95945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																								tmp95946 := Call(__e, tmp95945, V2783)

																																								tmp95947 := Call(__e, tmp95916, tmp95946)

																																								__e.TailApply(tmp95913, tmp95947)
																																								return

																																							}, 1)

																																							tmp95948 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																							tmp95949 := Call(__e, tmp95948, V2783)

																																							__e.TailApply(tmp95904, tmp95949)
																																							return

																																						}, 1)

																																						tmp95950 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																						tmp95951 := Call(__e, tmp95950, V2783)

																																						__e.TailApply(tmp95903, tmp95951)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp96301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp96302 := Call(__e, tmp96301, V2781, V2783)

																																			__e.TailApply(tmp95900, tmp96302)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp96305 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp96306 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp96307 := Call(__e, tmp96306, V2663)

																																	tmp96308 := Call(__e, tmp96305, tmp96307, V2783)

																																	__e.TailApply(tmp95898, tmp96308)
																																	return

																																}, 1)

																																tmp96309 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																tmp96310 := Call(__e, tmp96309, V2663)

																																__e.TailApply(tmp95897, tmp96310)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp96313 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																														tmp96314 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																														tmp96315 := Call(__e, tmp96314, V2662)

																														tmp96316 := Call(__e, tmp96313, tmp96315, V2783)

																														__e.TailApply(tmp95895, tmp96316)
																														return

																													}, 1)

																													tmp96317 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																													tmp96318 := Call(__e, tmp96317, V2662)

																													__e.TailApply(tmp95894, tmp96318)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp96321 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp96322 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																											tmp96323 := Call(__e, tmp96322, V2660)

																											tmp96324 := Call(__e, tmp96321, tmp96323, V2783)

																											__e.TailApply(tmp95892, tmp96324)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp96327 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																									tmp96328 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																									tmp96329 := Call(__e, tmp96328, V2660)

																									tmp96330 := Call(__e, tmp96327, tmp96329, V2783)

																									__e.TailApply(tmp95890, tmp96330)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp96333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																							tmp96334 := Call(__e, tmp96333, V2780, V2783)

																							tmp96335 := Call(__e, tmp95888, tmp96334)

																							__e.TailApply(tmp95119, tmp96335)
																							return

																						} else {
																							__e.Return(Case)
																							return
																						}

																					}, 1)

																					tmp96338 := MakeNative(func(__e *ControlFlow) {
																						V2654 := __e.Get(1)
																						_ = V2654
																						tmp96403 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																						tmp96404 := Call(__e, tmp96403, V2654)

																						if True == tmp96404 {
																							tmp96340 := MakeNative(func(__e *ControlFlow) {
																								V2655 := __e.Get(1)
																								_ = V2655
																								tmp96397 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																								tmp96398 := Call(__e, tmp96397, sym_8s, V2655)

																								if True == tmp96398 {
																									tmp96342 := MakeNative(func(__e *ControlFlow) {
																										V2656 := __e.Get(1)
																										_ = V2656
																										tmp96391 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																										tmp96392 := Call(__e, tmp96391, V2656)

																										if True == tmp96392 {
																											tmp96344 := MakeNative(func(__e *ControlFlow) {
																												X := __e.Get(1)
																												_ = X
																												tmp96345 := MakeNative(func(__e *ControlFlow) {
																													V2657 := __e.Get(1)
																													_ = V2657
																													tmp96383 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																													tmp96384 := Call(__e, tmp96383, V2657)

																													if True == tmp96384 {
																														tmp96347 := MakeNative(func(__e *ControlFlow) {
																															Y := __e.Get(1)
																															_ = Y
																															tmp96348 := MakeNative(func(__e *ControlFlow) {
																																V2658 := __e.Get(1)
																																_ = V2658
																																tmp96375 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp96376 := Call(__e, tmp96375, Nil, V2658)

																																if True == tmp96376 {
																																	tmp96350 := MakeNative(func(__e *ControlFlow) {
																																		V2659 := __e.Get(1)
																																		_ = V2659
																																		tmp96371 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp96372 := Call(__e, tmp96371, symstring, V2659)

																																		if True == tmp96372 {
																																			tmp96366 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp96367 := Call(__e, tmp96366)

																																			_ = tmp96367

																																			tmp96368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																			tmp96369 := MakeNative(func(__e *ControlFlow) {
																																				tmp96370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				__e.TailApply(tmp96370, Y, symstring, V2782, V2783, V2784)
																																				return

																																			}, 0)

																																			__e.TailApply(tmp96368, X, symstring, V2782, V2783, tmp96369)
																																			return

																																		} else {
																																			tmp96364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp96365 := Call(__e, tmp96364, V2659)

																																			if True == tmp96365 {
																																				tmp96353 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp96354 := Call(__e, tmp96353, V2659, symstring, V2783)

																																				_ = tmp96354

																																				tmp96355 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp96356 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp96357 := Call(__e, tmp96356, V2659, V2783)

																																					_ = tmp96357

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp96358 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp96359 := Call(__e, tmp96358)

																																				_ = tmp96359

																																				tmp96360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp96361 := MakeNative(func(__e *ControlFlow) {
																																					tmp96362 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					__e.TailApply(tmp96362, Y, symstring, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				tmp96363 := Call(__e, tmp96360, X, symstring, V2782, V2783, tmp96361)

																																				__e.TailApply(tmp96355, tmp96363)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp96373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp96374 := Call(__e, tmp96373, V2781, V2783)

																																	__e.TailApply(tmp96350, tmp96374)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp96377 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp96378 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp96379 := Call(__e, tmp96378, V2657)

																															tmp96380 := Call(__e, tmp96377, tmp96379, V2783)

																															__e.TailApply(tmp96348, tmp96380)
																															return

																														}, 1)

																														tmp96381 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp96382 := Call(__e, tmp96381, V2657)

																														__e.TailApply(tmp96347, tmp96382)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp96385 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																												tmp96386 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																												tmp96387 := Call(__e, tmp96386, V2656)

																												tmp96388 := Call(__e, tmp96385, tmp96387, V2783)

																												__e.TailApply(tmp96345, tmp96388)
																												return

																											}, 1)

																											tmp96389 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																											tmp96390 := Call(__e, tmp96389, V2656)

																											__e.TailApply(tmp96344, tmp96390)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp96393 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																									tmp96394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																									tmp96395 := Call(__e, tmp96394, V2654)

																									tmp96396 := Call(__e, tmp96393, tmp96395, V2783)

																									__e.TailApply(tmp96342, tmp96396)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp96399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																							tmp96400 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																							tmp96401 := Call(__e, tmp96400, V2654)

																							tmp96402 := Call(__e, tmp96399, tmp96401, V2783)

																							__e.TailApply(tmp96340, tmp96402)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp96405 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																					tmp96406 := Call(__e, tmp96405, V2780, V2783)

																					tmp96407 := Call(__e, tmp96338, tmp96406)

																					__e.TailApply(tmp95117, tmp96407)
																					return

																				} else {
																					__e.Return(Case)
																					return
																				}

																			}, 1)

																			tmp96410 := MakeNative(func(__e *ControlFlow) {
																				V2643 := __e.Get(1)
																				_ = V2643
																				tmp96636 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																				tmp96637 := Call(__e, tmp96636, V2643)

																				if True == tmp96637 {
																					tmp96412 := MakeNative(func(__e *ControlFlow) {
																						V2644 := __e.Get(1)
																						_ = V2644
																						tmp96630 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																						tmp96631 := Call(__e, tmp96630, sym_8v, V2644)

																						if True == tmp96631 {
																							tmp96414 := MakeNative(func(__e *ControlFlow) {
																								V2645 := __e.Get(1)
																								_ = V2645
																								tmp96624 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																								tmp96625 := Call(__e, tmp96624, V2645)

																								if True == tmp96625 {
																									tmp96416 := MakeNative(func(__e *ControlFlow) {
																										X := __e.Get(1)
																										_ = X
																										tmp96417 := MakeNative(func(__e *ControlFlow) {
																											V2646 := __e.Get(1)
																											_ = V2646
																											tmp96616 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																											tmp96617 := Call(__e, tmp96616, V2646)

																											if True == tmp96617 {
																												tmp96419 := MakeNative(func(__e *ControlFlow) {
																													Y := __e.Get(1)
																													_ = Y
																													tmp96420 := MakeNative(func(__e *ControlFlow) {
																														V2647 := __e.Get(1)
																														_ = V2647
																														tmp96608 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																														tmp96609 := Call(__e, tmp96608, Nil, V2647)

																														if True == tmp96609 {
																															tmp96422 := MakeNative(func(__e *ControlFlow) {
																																V2648 := __e.Get(1)
																																_ = V2648
																																tmp96604 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																tmp96605 := Call(__e, tmp96604, V2648)

																																if True == tmp96605 {
																																	tmp96449 := MakeNative(func(__e *ControlFlow) {
																																		V2649 := __e.Get(1)
																																		_ = V2649
																																		tmp96598 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp96599 := Call(__e, tmp96598, symvector, V2649)

																																		if True == tmp96599 {
																																			tmp96529 := MakeNative(func(__e *ControlFlow) {
																																				V2650 := __e.Get(1)
																																				_ = V2650
																																				tmp96592 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																				tmp96593 := Call(__e, tmp96592, V2650)

																																				if True == tmp96593 {
																																					tmp96554 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp96555 := MakeNative(func(__e *ControlFlow) {
																																							V2651 := __e.Get(1)
																																							_ = V2651
																																							tmp96584 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp96585 := Call(__e, tmp96584, Nil, V2651)

																																							if True == tmp96585 {
																																								tmp96575 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp96576 := Call(__e, tmp96575)

																																								_ = tmp96576

																																								tmp96577 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																								tmp96578 := MakeNative(func(__e *ControlFlow) {
																																									tmp96579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96580 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp96581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp96582 := Call(__e, tmp96581, A, Nil)

																																									tmp96583 := Call(__e, tmp96580, symvector, tmp96582)

																																									__e.TailApply(tmp96579, Y, tmp96583, V2782, V2783, V2784)
																																									return

																																								}, 0)

																																								__e.TailApply(tmp96577, X, A, V2782, V2783, tmp96578)
																																								return

																																							} else {
																																								tmp96573 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp96574 := Call(__e, tmp96573, V2651)

																																								if True == tmp96574 {
																																									tmp96558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp96559 := Call(__e, tmp96558, V2651, Nil, V2783)

																																									_ = tmp96559

																																									tmp96560 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp96561 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp96562 := Call(__e, tmp96561, V2651, V2783)

																																										_ = tmp96562

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp96563 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp96564 := Call(__e, tmp96563)

																																									_ = tmp96564

																																									tmp96565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96566 := MakeNative(func(__e *ControlFlow) {
																																										tmp96567 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										tmp96568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp96569 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp96570 := Call(__e, tmp96569, A, Nil)

																																										tmp96571 := Call(__e, tmp96568, symvector, tmp96570)

																																										__e.TailApply(tmp96567, Y, tmp96571, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									tmp96572 := Call(__e, tmp96565, X, A, V2782, V2783, tmp96566)

																																									__e.TailApply(tmp96560, tmp96572)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp96586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp96587 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp96588 := Call(__e, tmp96587, V2650)

																																						tmp96589 := Call(__e, tmp96586, tmp96588, V2783)

																																						__e.TailApply(tmp96555, tmp96589)
																																						return

																																					}, 1)

																																					tmp96590 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																					tmp96591 := Call(__e, tmp96590, V2650)

																																					__e.TailApply(tmp96554, tmp96591)
																																					return

																																				} else {
																																					tmp96552 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																					tmp96553 := Call(__e, tmp96552, V2650)

																																					if True == tmp96553 {
																																						tmp96532 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp96533 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																							tmp96534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp96535 := Call(__e, tmp96534, A, Nil)

																																							tmp96536 := Call(__e, tmp96533, V2650, tmp96535, V2783)

																																							_ = tmp96536

																																							tmp96537 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp96538 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																								tmp96539 := Call(__e, tmp96538, V2650, V2783)

																																								_ = tmp96539

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp96540 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																							tmp96541 := Call(__e, tmp96540)

																																							_ = tmp96541

																																							tmp96542 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																							tmp96543 := MakeNative(func(__e *ControlFlow) {
																																								tmp96544 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																								tmp96545 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp96546 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp96547 := Call(__e, tmp96546, A, Nil)

																																								tmp96548 := Call(__e, tmp96545, symvector, tmp96547)

																																								__e.TailApply(tmp96544, Y, tmp96548, V2782, V2783, V2784)
																																								return

																																							}, 0)

																																							tmp96549 := Call(__e, tmp96542, X, A, V2782, V2783, tmp96543)

																																							__e.TailApply(tmp96537, tmp96549)
																																							return

																																						}, 1)

																																						tmp96550 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																						tmp96551 := Call(__e, tmp96550, V2783)

																																						__e.TailApply(tmp96532, tmp96551)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp96594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp96595 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp96596 := Call(__e, tmp96595, V2648)

																																			tmp96597 := Call(__e, tmp96594, tmp96596, V2783)

																																			__e.TailApply(tmp96529, tmp96597)
																																			return

																																		} else {
																																			tmp96527 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp96528 := Call(__e, tmp96527, V2649)

																																			if True == tmp96528 {
																																				tmp96452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp96453 := Call(__e, tmp96452, V2649, symvector, V2783)

																																				_ = tmp96453

																																				tmp96454 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp96455 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp96456 := Call(__e, tmp96455, V2649, V2783)

																																					_ = tmp96456

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp96457 := MakeNative(func(__e *ControlFlow) {
																																					V2652 := __e.Get(1)
																																					_ = V2652
																																					tmp96520 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																					tmp96521 := Call(__e, tmp96520, V2652)

																																					if True == tmp96521 {
																																						tmp96482 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp96483 := MakeNative(func(__e *ControlFlow) {
																																								V2653 := __e.Get(1)
																																								_ = V2653
																																								tmp96512 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp96513 := Call(__e, tmp96512, Nil, V2653)

																																								if True == tmp96513 {
																																									tmp96503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp96504 := Call(__e, tmp96503)

																																									_ = tmp96504

																																									tmp96505 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96506 := MakeNative(func(__e *ControlFlow) {
																																										tmp96507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										tmp96508 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp96509 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp96510 := Call(__e, tmp96509, A, Nil)

																																										tmp96511 := Call(__e, tmp96508, symvector, tmp96510)

																																										__e.TailApply(tmp96507, Y, tmp96511, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									__e.TailApply(tmp96505, X, A, V2782, V2783, tmp96506)
																																									return

																																								} else {
																																									tmp96501 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp96502 := Call(__e, tmp96501, V2653)

																																									if True == tmp96502 {
																																										tmp96486 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp96487 := Call(__e, tmp96486, V2653, Nil, V2783)

																																										_ = tmp96487

																																										tmp96488 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp96489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp96490 := Call(__e, tmp96489, V2653, V2783)

																																											_ = tmp96490

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp96491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp96492 := Call(__e, tmp96491)

																																										_ = tmp96492

																																										tmp96493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										tmp96494 := MakeNative(func(__e *ControlFlow) {
																																											tmp96495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											tmp96496 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp96497 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp96498 := Call(__e, tmp96497, A, Nil)

																																											tmp96499 := Call(__e, tmp96496, symvector, tmp96498)

																																											__e.TailApply(tmp96495, Y, tmp96499, V2782, V2783, V2784)
																																											return

																																										}, 0)

																																										tmp96500 := Call(__e, tmp96493, X, A, V2782, V2783, tmp96494)

																																										__e.TailApply(tmp96488, tmp96500)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp96514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp96515 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp96516 := Call(__e, tmp96515, V2652)

																																							tmp96517 := Call(__e, tmp96514, tmp96516, V2783)

																																							__e.TailApply(tmp96483, tmp96517)
																																							return

																																						}, 1)

																																						tmp96518 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																						tmp96519 := Call(__e, tmp96518, V2652)

																																						__e.TailApply(tmp96482, tmp96519)
																																						return

																																					} else {
																																						tmp96480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																						tmp96481 := Call(__e, tmp96480, V2652)

																																						if True == tmp96481 {
																																							tmp96460 := MakeNative(func(__e *ControlFlow) {
																																								A := __e.Get(1)
																																								_ = A
																																								tmp96461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp96462 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp96463 := Call(__e, tmp96462, A, Nil)

																																								tmp96464 := Call(__e, tmp96461, V2652, tmp96463, V2783)

																																								_ = tmp96464

																																								tmp96465 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp96466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp96467 := Call(__e, tmp96466, V2652, V2783)

																																									_ = tmp96467

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp96468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp96469 := Call(__e, tmp96468)

																																								_ = tmp96469

																																								tmp96470 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																								tmp96471 := MakeNative(func(__e *ControlFlow) {
																																									tmp96472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp96474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp96475 := Call(__e, tmp96474, A, Nil)

																																									tmp96476 := Call(__e, tmp96473, symvector, tmp96475)

																																									__e.TailApply(tmp96472, Y, tmp96476, V2782, V2783, V2784)
																																									return

																																								}, 0)

																																								tmp96477 := Call(__e, tmp96470, X, A, V2782, V2783, tmp96471)

																																								__e.TailApply(tmp96465, tmp96477)
																																								return

																																							}, 1)

																																							tmp96478 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																							tmp96479 := Call(__e, tmp96478, V2783)

																																							__e.TailApply(tmp96460, tmp96479)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp96522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp96523 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp96524 := Call(__e, tmp96523, V2648)

																																				tmp96525 := Call(__e, tmp96522, tmp96524, V2783)

																																				tmp96526 := Call(__e, tmp96457, tmp96525)

																																				__e.TailApply(tmp96454, tmp96526)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp96600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp96601 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																	tmp96602 := Call(__e, tmp96601, V2648)

																																	tmp96603 := Call(__e, tmp96600, tmp96602, V2783)

																																	__e.TailApply(tmp96449, tmp96603)
																																	return

																																} else {
																																	tmp96447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																	tmp96448 := Call(__e, tmp96447, V2648)

																																	if True == tmp96448 {
																																		tmp96425 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp96426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp96427 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp96428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp96429 := Call(__e, tmp96428, A, Nil)

																																			tmp96430 := Call(__e, tmp96427, symvector, tmp96429)

																																			tmp96431 := Call(__e, tmp96426, V2648, tmp96430, V2783)

																																			_ = tmp96431

																																			tmp96432 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp96433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp96434 := Call(__e, tmp96433, V2648, V2783)

																																				_ = tmp96434

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp96435 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp96436 := Call(__e, tmp96435)

																																			_ = tmp96436

																																			tmp96437 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																			tmp96438 := MakeNative(func(__e *ControlFlow) {
																																				tmp96439 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp96440 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp96441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp96442 := Call(__e, tmp96441, A, Nil)

																																				tmp96443 := Call(__e, tmp96440, symvector, tmp96442)

																																				__e.TailApply(tmp96439, Y, tmp96443, V2782, V2783, V2784)
																																				return

																																			}, 0)

																																			tmp96444 := Call(__e, tmp96437, X, A, V2782, V2783, tmp96438)

																																			__e.TailApply(tmp96432, tmp96444)
																																			return

																																		}, 1)

																																		tmp96445 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																		tmp96446 := Call(__e, tmp96445, V2783)

																																		__e.TailApply(tmp96425, tmp96446)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp96606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp96607 := Call(__e, tmp96606, V2781, V2783)

																															__e.TailApply(tmp96422, tmp96607)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp96610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp96611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																													tmp96612 := Call(__e, tmp96611, V2646)

																													tmp96613 := Call(__e, tmp96610, tmp96612, V2783)

																													__e.TailApply(tmp96420, tmp96613)
																													return

																												}, 1)

																												tmp96614 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																												tmp96615 := Call(__e, tmp96614, V2646)

																												__e.TailApply(tmp96419, tmp96615)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp96618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																										tmp96619 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																										tmp96620 := Call(__e, tmp96619, V2645)

																										tmp96621 := Call(__e, tmp96618, tmp96620, V2783)

																										__e.TailApply(tmp96417, tmp96621)
																										return

																									}, 1)

																									tmp96622 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																									tmp96623 := Call(__e, tmp96622, V2645)

																									__e.TailApply(tmp96416, tmp96623)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp96626 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																							tmp96627 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp96628 := Call(__e, tmp96627, V2643)

																							tmp96629 := Call(__e, tmp96626, tmp96628, V2783)

																							__e.TailApply(tmp96414, tmp96629)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp96632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																					tmp96633 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																					tmp96634 := Call(__e, tmp96633, V2643)

																					tmp96635 := Call(__e, tmp96632, tmp96634, V2783)

																					__e.TailApply(tmp96412, tmp96635)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp96638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp96639 := Call(__e, tmp96638, V2780, V2783)

																			tmp96640 := Call(__e, tmp96410, tmp96639)

																			__e.TailApply(tmp95115, tmp96640)
																			return

																		} else {
																			__e.Return(Case)
																			return
																		}

																	}, 1)

																	tmp96643 := MakeNative(func(__e *ControlFlow) {
																		V2631 := __e.Get(1)
																		_ = V2631
																		tmp96878 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp96879 := Call(__e, tmp96878, V2631)

																		if True == tmp96879 {
																			tmp96645 := MakeNative(func(__e *ControlFlow) {
																				V2632 := __e.Get(1)
																				_ = V2632
																				tmp96872 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp96873 := Call(__e, tmp96872, sym_8p, V2632)

																				if True == tmp96873 {
																					tmp96647 := MakeNative(func(__e *ControlFlow) {
																						V2633 := __e.Get(1)
																						_ = V2633
																						tmp96866 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																						tmp96867 := Call(__e, tmp96866, V2633)

																						if True == tmp96867 {
																							tmp96649 := MakeNative(func(__e *ControlFlow) {
																								X := __e.Get(1)
																								_ = X
																								tmp96650 := MakeNative(func(__e *ControlFlow) {
																									V2634 := __e.Get(1)
																									_ = V2634
																									tmp96858 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																									tmp96859 := Call(__e, tmp96858, V2634)

																									if True == tmp96859 {
																										tmp96652 := MakeNative(func(__e *ControlFlow) {
																											Y := __e.Get(1)
																											_ = Y
																											tmp96653 := MakeNative(func(__e *ControlFlow) {
																												V2635 := __e.Get(1)
																												_ = V2635
																												tmp96850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																												tmp96851 := Call(__e, tmp96850, Nil, V2635)

																												if True == tmp96851 {
																													tmp96655 := MakeNative(func(__e *ControlFlow) {
																														V2636 := __e.Get(1)
																														_ = V2636
																														tmp96846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																														tmp96847 := Call(__e, tmp96846, V2636)

																														if True == tmp96847 {
																															tmp96683 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp96684 := MakeNative(func(__e *ControlFlow) {
																																	V2637 := __e.Get(1)
																																	_ = V2637
																																	tmp96838 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																	tmp96839 := Call(__e, tmp96838, V2637)

																																	if True == tmp96839 {
																																		tmp96707 := MakeNative(func(__e *ControlFlow) {
																																			V2638 := __e.Get(1)
																																			_ = V2638
																																			tmp96832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																			tmp96833 := Call(__e, tmp96832, sym_d, V2638)

																																			if True == tmp96833 {
																																				tmp96775 := MakeNative(func(__e *ControlFlow) {
																																					V2639 := __e.Get(1)
																																					_ = V2639
																																					tmp96826 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																					tmp96827 := Call(__e, tmp96826, V2639)

																																					if True == tmp96827 {
																																						tmp96796 := MakeNative(func(__e *ControlFlow) {
																																							B := __e.Get(1)
																																							_ = B
																																							tmp96797 := MakeNative(func(__e *ControlFlow) {
																																								V2640 := __e.Get(1)
																																								_ = V2640
																																								tmp96818 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp96819 := Call(__e, tmp96818, Nil, V2640)

																																								if True == tmp96819 {
																																									tmp96813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp96814 := Call(__e, tmp96813)

																																									_ = tmp96814

																																									tmp96815 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96816 := MakeNative(func(__e *ControlFlow) {
																																										tmp96817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										__e.TailApply(tmp96817, Y, B, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									__e.TailApply(tmp96815, X, A, V2782, V2783, tmp96816)
																																									return

																																								} else {
																																									tmp96811 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp96812 := Call(__e, tmp96811, V2640)

																																									if True == tmp96812 {
																																										tmp96800 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp96801 := Call(__e, tmp96800, V2640, Nil, V2783)

																																										_ = tmp96801

																																										tmp96802 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp96803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp96804 := Call(__e, tmp96803, V2640, V2783)

																																											_ = tmp96804

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp96805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp96806 := Call(__e, tmp96805)

																																										_ = tmp96806

																																										tmp96807 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										tmp96808 := MakeNative(func(__e *ControlFlow) {
																																											tmp96809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											__e.TailApply(tmp96809, Y, B, V2782, V2783, V2784)
																																											return

																																										}, 0)

																																										tmp96810 := Call(__e, tmp96807, X, A, V2782, V2783, tmp96808)

																																										__e.TailApply(tmp96802, tmp96810)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp96820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp96821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp96822 := Call(__e, tmp96821, V2639)

																																							tmp96823 := Call(__e, tmp96820, tmp96822, V2783)

																																							__e.TailApply(tmp96797, tmp96823)
																																							return

																																						}, 1)

																																						tmp96824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																						tmp96825 := Call(__e, tmp96824, V2639)

																																						__e.TailApply(tmp96796, tmp96825)
																																						return

																																					} else {
																																						tmp96794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																						tmp96795 := Call(__e, tmp96794, V2639)

																																						if True == tmp96795 {
																																							tmp96778 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp96779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp96780 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp96781 := Call(__e, tmp96780, B, Nil)

																																								tmp96782 := Call(__e, tmp96779, V2639, tmp96781, V2783)

																																								_ = tmp96782

																																								tmp96783 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp96784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp96785 := Call(__e, tmp96784, V2639, V2783)

																																									_ = tmp96785

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp96786 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp96787 := Call(__e, tmp96786)

																																								_ = tmp96787

																																								tmp96788 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																								tmp96789 := MakeNative(func(__e *ControlFlow) {
																																									tmp96790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									__e.TailApply(tmp96790, Y, B, V2782, V2783, V2784)
																																									return

																																								}, 0)

																																								tmp96791 := Call(__e, tmp96788, X, A, V2782, V2783, tmp96789)

																																								__e.TailApply(tmp96783, tmp96791)
																																								return

																																							}, 1)

																																							tmp96792 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																							tmp96793 := Call(__e, tmp96792, V2783)

																																							__e.TailApply(tmp96778, tmp96793)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp96828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp96829 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp96830 := Call(__e, tmp96829, V2637)

																																				tmp96831 := Call(__e, tmp96828, tmp96830, V2783)

																																				__e.TailApply(tmp96775, tmp96831)
																																				return

																																			} else {
																																				tmp96773 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																				tmp96774 := Call(__e, tmp96773, V2638)

																																				if True == tmp96774 {
																																					tmp96710 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp96711 := Call(__e, tmp96710, V2638, sym_d, V2783)

																																					_ = tmp96711

																																					tmp96712 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp96713 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp96714 := Call(__e, tmp96713, V2638, V2783)

																																						_ = tmp96714

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp96715 := MakeNative(func(__e *ControlFlow) {
																																						V2641 := __e.Get(1)
																																						_ = V2641
																																						tmp96766 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																						tmp96767 := Call(__e, tmp96766, V2641)

																																						if True == tmp96767 {
																																							tmp96736 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp96737 := MakeNative(func(__e *ControlFlow) {
																																									V2642 := __e.Get(1)
																																									_ = V2642
																																									tmp96758 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																									tmp96759 := Call(__e, tmp96758, Nil, V2642)

																																									if True == tmp96759 {
																																										tmp96753 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp96754 := Call(__e, tmp96753)

																																										_ = tmp96754

																																										tmp96755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										tmp96756 := MakeNative(func(__e *ControlFlow) {
																																											tmp96757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											__e.TailApply(tmp96757, Y, B, V2782, V2783, V2784)
																																											return

																																										}, 0)

																																										__e.TailApply(tmp96755, X, A, V2782, V2783, tmp96756)
																																										return

																																									} else {
																																										tmp96751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																										tmp96752 := Call(__e, tmp96751, V2642)

																																										if True == tmp96752 {
																																											tmp96740 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																											tmp96741 := Call(__e, tmp96740, V2642, Nil, V2783)

																																											_ = tmp96741

																																											tmp96742 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp96743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																												tmp96744 := Call(__e, tmp96743, V2642, V2783)

																																												_ = tmp96744

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp96745 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp96746 := Call(__e, tmp96745)

																																											_ = tmp96746

																																											tmp96747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																											tmp96748 := MakeNative(func(__e *ControlFlow) {
																																												tmp96749 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																												__e.TailApply(tmp96749, Y, B, V2782, V2783, V2784)
																																												return

																																											}, 0)

																																											tmp96750 := Call(__e, tmp96747, X, A, V2782, V2783, tmp96748)

																																											__e.TailApply(tmp96742, tmp96750)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp96760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp96761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp96762 := Call(__e, tmp96761, V2641)

																																								tmp96763 := Call(__e, tmp96760, tmp96762, V2783)

																																								__e.TailApply(tmp96737, tmp96763)
																																								return

																																							}, 1)

																																							tmp96764 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																							tmp96765 := Call(__e, tmp96764, V2641)

																																							__e.TailApply(tmp96736, tmp96765)
																																							return

																																						} else {
																																							tmp96734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																							tmp96735 := Call(__e, tmp96734, V2641)

																																							if True == tmp96735 {
																																								tmp96718 := MakeNative(func(__e *ControlFlow) {
																																									B := __e.Get(1)
																																									_ = B
																																									tmp96719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp96720 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp96721 := Call(__e, tmp96720, B, Nil)

																																									tmp96722 := Call(__e, tmp96719, V2641, tmp96721, V2783)

																																									_ = tmp96722

																																									tmp96723 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp96724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp96725 := Call(__e, tmp96724, V2641, V2783)

																																										_ = tmp96725

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp96726 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp96727 := Call(__e, tmp96726)

																																									_ = tmp96727

																																									tmp96728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																									tmp96729 := MakeNative(func(__e *ControlFlow) {
																																										tmp96730 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																										__e.TailApply(tmp96730, Y, B, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									tmp96731 := Call(__e, tmp96728, X, A, V2782, V2783, tmp96729)

																																									__e.TailApply(tmp96723, tmp96731)
																																									return

																																								}, 1)

																																								tmp96732 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																								tmp96733 := Call(__e, tmp96732, V2783)

																																								__e.TailApply(tmp96718, tmp96733)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp96768 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp96769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp96770 := Call(__e, tmp96769, V2637)

																																					tmp96771 := Call(__e, tmp96768, tmp96770, V2783)

																																					tmp96772 := Call(__e, tmp96715, tmp96771)

																																					__e.TailApply(tmp96712, tmp96772)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp96834 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp96835 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																		tmp96836 := Call(__e, tmp96835, V2637)

																																		tmp96837 := Call(__e, tmp96834, tmp96836, V2783)

																																		__e.TailApply(tmp96707, tmp96837)
																																		return

																																	} else {
																																		tmp96705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp96706 := Call(__e, tmp96705, V2637)

																																		if True == tmp96706 {
																																			tmp96687 := MakeNative(func(__e *ControlFlow) {
																																				B := __e.Get(1)
																																				_ = B
																																				tmp96688 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp96689 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp96690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp96691 := Call(__e, tmp96690, B, Nil)

																																				tmp96692 := Call(__e, tmp96689, sym_d, tmp96691)

																																				tmp96693 := Call(__e, tmp96688, V2637, tmp96692, V2783)

																																				_ = tmp96693

																																				tmp96694 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp96695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp96696 := Call(__e, tmp96695, V2637, V2783)

																																					_ = tmp96696

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp96697 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp96698 := Call(__e, tmp96697)

																																				_ = tmp96698

																																				tmp96699 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp96700 := MakeNative(func(__e *ControlFlow) {
																																					tmp96701 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					__e.TailApply(tmp96701, Y, B, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				tmp96702 := Call(__e, tmp96699, X, A, V2782, V2783, tmp96700)

																																				__e.TailApply(tmp96694, tmp96702)
																																				return

																																			}, 1)

																																			tmp96703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																			tmp96704 := Call(__e, tmp96703, V2783)

																																			__e.TailApply(tmp96687, tmp96704)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp96840 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp96841 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp96842 := Call(__e, tmp96841, V2636)

																																tmp96843 := Call(__e, tmp96840, tmp96842, V2783)

																																__e.TailApply(tmp96684, tmp96843)
																																return

																															}, 1)

																															tmp96844 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																															tmp96845 := Call(__e, tmp96844, V2636)

																															__e.TailApply(tmp96683, tmp96845)
																															return

																														} else {
																															tmp96681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																															tmp96682 := Call(__e, tmp96681, V2636)

																															if True == tmp96682 {
																																tmp96658 := MakeNative(func(__e *ControlFlow) {
																																	A := __e.Get(1)
																																	_ = A
																																	tmp96659 := MakeNative(func(__e *ControlFlow) {
																																		B := __e.Get(1)
																																		_ = B
																																		tmp96660 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																		tmp96661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp96662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp96663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp96664 := Call(__e, tmp96663, B, Nil)

																																		tmp96665 := Call(__e, tmp96662, sym_d, tmp96664)

																																		tmp96666 := Call(__e, tmp96661, A, tmp96665)

																																		tmp96667 := Call(__e, tmp96660, V2636, tmp96666, V2783)

																																		_ = tmp96667

																																		tmp96668 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp96669 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																			tmp96670 := Call(__e, tmp96669, V2636, V2783)

																																			_ = tmp96670

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp96671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																		tmp96672 := Call(__e, tmp96671)

																																		_ = tmp96672

																																		tmp96673 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																		tmp96674 := MakeNative(func(__e *ControlFlow) {
																																			tmp96675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																			__e.TailApply(tmp96675, Y, B, V2782, V2783, V2784)
																																			return

																																		}, 0)

																																		tmp96676 := Call(__e, tmp96673, X, A, V2782, V2783, tmp96674)

																																		__e.TailApply(tmp96668, tmp96676)
																																		return

																																	}, 1)

																																	tmp96677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																	tmp96678 := Call(__e, tmp96677, V2783)

																																	__e.TailApply(tmp96659, tmp96678)
																																	return

																																}, 1)

																																tmp96679 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																tmp96680 := Call(__e, tmp96679, V2783)

																																__e.TailApply(tmp96658, tmp96680)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp96848 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp96849 := Call(__e, tmp96848, V2781, V2783)

																													__e.TailApply(tmp96655, tmp96849)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp96852 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp96853 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																											tmp96854 := Call(__e, tmp96853, V2634)

																											tmp96855 := Call(__e, tmp96852, tmp96854, V2783)

																											__e.TailApply(tmp96653, tmp96855)
																											return

																										}, 1)

																										tmp96856 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										tmp96857 := Call(__e, tmp96856, V2634)

																										__e.TailApply(tmp96652, tmp96857)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp96860 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																								tmp96861 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																								tmp96862 := Call(__e, tmp96861, V2633)

																								tmp96863 := Call(__e, tmp96860, tmp96862, V2783)

																								__e.TailApply(tmp96650, tmp96863)
																								return

																							}, 1)

																							tmp96864 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																							tmp96865 := Call(__e, tmp96864, V2633)

																							__e.TailApply(tmp96649, tmp96865)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp96868 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																					tmp96869 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp96870 := Call(__e, tmp96869, V2631)

																					tmp96871 := Call(__e, tmp96868, tmp96870, V2783)

																					__e.TailApply(tmp96647, tmp96871)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp96874 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp96875 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp96876 := Call(__e, tmp96875, V2631)

																			tmp96877 := Call(__e, tmp96874, tmp96876, V2783)

																			__e.TailApply(tmp96645, tmp96877)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp96880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp96881 := Call(__e, tmp96880, V2780, V2783)

																	tmp96882 := Call(__e, tmp96643, tmp96881)

																	__e.TailApply(tmp95113, tmp96882)
																	return

																} else {
																	__e.Return(Case)
																	return
																}

															}, 1)

															tmp96885 := MakeNative(func(__e *ControlFlow) {
																V2620 := __e.Get(1)
																_ = V2620
																tmp97111 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp97112 := Call(__e, tmp97111, V2620)

																if True == tmp97112 {
																	tmp96887 := MakeNative(func(__e *ControlFlow) {
																		V2621 := __e.Get(1)
																		_ = V2621
																		tmp97105 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																		tmp97106 := Call(__e, tmp97105, symcons, V2621)

																		if True == tmp97106 {
																			tmp96889 := MakeNative(func(__e *ControlFlow) {
																				V2622 := __e.Get(1)
																				_ = V2622
																				tmp97099 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																				tmp97100 := Call(__e, tmp97099, V2622)

																				if True == tmp97100 {
																					tmp96891 := MakeNative(func(__e *ControlFlow) {
																						X := __e.Get(1)
																						_ = X
																						tmp96892 := MakeNative(func(__e *ControlFlow) {
																							V2623 := __e.Get(1)
																							_ = V2623
																							tmp97091 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																							tmp97092 := Call(__e, tmp97091, V2623)

																							if True == tmp97092 {
																								tmp96894 := MakeNative(func(__e *ControlFlow) {
																									Y := __e.Get(1)
																									_ = Y
																									tmp96895 := MakeNative(func(__e *ControlFlow) {
																										V2624 := __e.Get(1)
																										_ = V2624
																										tmp97083 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																										tmp97084 := Call(__e, tmp97083, Nil, V2624)

																										if True == tmp97084 {
																											tmp96897 := MakeNative(func(__e *ControlFlow) {
																												V2625 := __e.Get(1)
																												_ = V2625
																												tmp97079 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																												tmp97080 := Call(__e, tmp97079, V2625)

																												if True == tmp97080 {
																													tmp96924 := MakeNative(func(__e *ControlFlow) {
																														V2626 := __e.Get(1)
																														_ = V2626
																														tmp97073 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																														tmp97074 := Call(__e, tmp97073, symlist, V2626)

																														if True == tmp97074 {
																															tmp97004 := MakeNative(func(__e *ControlFlow) {
																																V2627 := __e.Get(1)
																																_ = V2627
																																tmp97067 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																tmp97068 := Call(__e, tmp97067, V2627)

																																if True == tmp97068 {
																																	tmp97029 := MakeNative(func(__e *ControlFlow) {
																																		A := __e.Get(1)
																																		_ = A
																																		tmp97030 := MakeNative(func(__e *ControlFlow) {
																																			V2628 := __e.Get(1)
																																			_ = V2628
																																			tmp97059 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																			tmp97060 := Call(__e, tmp97059, Nil, V2628)

																																			if True == tmp97060 {
																																				tmp97050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp97051 := Call(__e, tmp97050)

																																				_ = tmp97051

																																				tmp97052 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp97053 := MakeNative(func(__e *ControlFlow) {
																																					tmp97054 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					tmp97055 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97057 := Call(__e, tmp97056, A, Nil)

																																					tmp97058 := Call(__e, tmp97055, symlist, tmp97057)

																																					__e.TailApply(tmp97054, Y, tmp97058, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				__e.TailApply(tmp97052, X, A, V2782, V2783, tmp97053)
																																				return

																																			} else {
																																				tmp97048 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																				tmp97049 := Call(__e, tmp97048, V2628)

																																				if True == tmp97049 {
																																					tmp97033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp97034 := Call(__e, tmp97033, V2628, Nil, V2783)

																																					_ = tmp97034

																																					tmp97035 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp97036 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp97037 := Call(__e, tmp97036, V2628, V2783)

																																						_ = tmp97037

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp97038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp97039 := Call(__e, tmp97038)

																																					_ = tmp97039

																																					tmp97040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					tmp97041 := MakeNative(func(__e *ControlFlow) {
																																						tmp97042 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																						tmp97043 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp97044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp97045 := Call(__e, tmp97044, A, Nil)

																																						tmp97046 := Call(__e, tmp97043, symlist, tmp97045)

																																						__e.TailApply(tmp97042, Y, tmp97046, V2782, V2783, V2784)
																																						return

																																					}, 0)

																																					tmp97047 := Call(__e, tmp97040, X, A, V2782, V2783, tmp97041)

																																					__e.TailApply(tmp97035, tmp97047)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp97061 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp97062 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp97063 := Call(__e, tmp97062, V2627)

																																		tmp97064 := Call(__e, tmp97061, tmp97063, V2783)

																																		__e.TailApply(tmp97030, tmp97064)
																																		return

																																	}, 1)

																																	tmp97065 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																	tmp97066 := Call(__e, tmp97065, V2627)

																																	__e.TailApply(tmp97029, tmp97066)
																																	return

																																} else {
																																	tmp97027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																	tmp97028 := Call(__e, tmp97027, V2627)

																																	if True == tmp97028 {
																																		tmp97007 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp97008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp97009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97010 := Call(__e, tmp97009, A, Nil)

																																			tmp97011 := Call(__e, tmp97008, V2627, tmp97010, V2783)

																																			_ = tmp97011

																																			tmp97012 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp97013 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp97014 := Call(__e, tmp97013, V2627, V2783)

																																				_ = tmp97014

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp97015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp97016 := Call(__e, tmp97015)

																																			_ = tmp97016

																																			tmp97017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																			tmp97018 := MakeNative(func(__e *ControlFlow) {
																																				tmp97019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp97020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97021 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97022 := Call(__e, tmp97021, A, Nil)

																																				tmp97023 := Call(__e, tmp97020, symlist, tmp97022)

																																				__e.TailApply(tmp97019, Y, tmp97023, V2782, V2783, V2784)
																																				return

																																			}, 0)

																																			tmp97024 := Call(__e, tmp97017, X, A, V2782, V2783, tmp97018)

																																			__e.TailApply(tmp97012, tmp97024)
																																			return

																																		}, 1)

																																		tmp97025 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																		tmp97026 := Call(__e, tmp97025, V2783)

																																		__e.TailApply(tmp97007, tmp97026)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp97069 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp97070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp97071 := Call(__e, tmp97070, V2625)

																															tmp97072 := Call(__e, tmp97069, tmp97071, V2783)

																															__e.TailApply(tmp97004, tmp97072)
																															return

																														} else {
																															tmp97002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																															tmp97003 := Call(__e, tmp97002, V2626)

																															if True == tmp97003 {
																																tmp96927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																tmp96928 := Call(__e, tmp96927, V2626, symlist, V2783)

																																_ = tmp96928

																																tmp96929 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp96930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																	tmp96931 := Call(__e, tmp96930, V2626, V2783)

																																	_ = tmp96931

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp96932 := MakeNative(func(__e *ControlFlow) {
																																	V2629 := __e.Get(1)
																																	_ = V2629
																																	tmp96995 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																	tmp96996 := Call(__e, tmp96995, V2629)

																																	if True == tmp96996 {
																																		tmp96957 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp96958 := MakeNative(func(__e *ControlFlow) {
																																				V2630 := __e.Get(1)
																																				_ = V2630
																																				tmp96987 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																				tmp96988 := Call(__e, tmp96987, Nil, V2630)

																																				if True == tmp96988 {
																																					tmp96978 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp96979 := Call(__e, tmp96978)

																																					_ = tmp96979

																																					tmp96980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					tmp96981 := MakeNative(func(__e *ControlFlow) {
																																						tmp96982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																						tmp96983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp96984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp96985 := Call(__e, tmp96984, A, Nil)

																																						tmp96986 := Call(__e, tmp96983, symlist, tmp96985)

																																						__e.TailApply(tmp96982, Y, tmp96986, V2782, V2783, V2784)
																																						return

																																					}, 0)

																																					__e.TailApply(tmp96980, X, A, V2782, V2783, tmp96981)
																																					return

																																				} else {
																																					tmp96976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																					tmp96977 := Call(__e, tmp96976, V2630)

																																					if True == tmp96977 {
																																						tmp96961 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																						tmp96962 := Call(__e, tmp96961, V2630, Nil, V2783)

																																						_ = tmp96962

																																						tmp96963 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp96964 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																							tmp96965 := Call(__e, tmp96964, V2630, V2783)

																																							_ = tmp96965

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp96966 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																						tmp96967 := Call(__e, tmp96966)

																																						_ = tmp96967

																																						tmp96968 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																						tmp96969 := MakeNative(func(__e *ControlFlow) {
																																							tmp96970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																							tmp96971 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp96972 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp96973 := Call(__e, tmp96972, A, Nil)

																																							tmp96974 := Call(__e, tmp96971, symlist, tmp96973)

																																							__e.TailApply(tmp96970, Y, tmp96974, V2782, V2783, V2784)
																																							return

																																						}, 0)

																																						tmp96975 := Call(__e, tmp96968, X, A, V2782, V2783, tmp96969)

																																						__e.TailApply(tmp96963, tmp96975)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp96989 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp96990 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp96991 := Call(__e, tmp96990, V2629)

																																			tmp96992 := Call(__e, tmp96989, tmp96991, V2783)

																																			__e.TailApply(tmp96958, tmp96992)
																																			return

																																		}, 1)

																																		tmp96993 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																		tmp96994 := Call(__e, tmp96993, V2629)

																																		__e.TailApply(tmp96957, tmp96994)
																																		return

																																	} else {
																																		tmp96955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp96956 := Call(__e, tmp96955, V2629)

																																		if True == tmp96956 {
																																			tmp96935 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp96936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp96937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp96938 := Call(__e, tmp96937, A, Nil)

																																				tmp96939 := Call(__e, tmp96936, V2629, tmp96938, V2783)

																																				_ = tmp96939

																																				tmp96940 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp96941 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp96942 := Call(__e, tmp96941, V2629, V2783)

																																					_ = tmp96942

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp96943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp96944 := Call(__e, tmp96943)

																																				_ = tmp96944

																																				tmp96945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																				tmp96946 := MakeNative(func(__e *ControlFlow) {
																																					tmp96947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																					tmp96948 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp96949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp96950 := Call(__e, tmp96949, A, Nil)

																																					tmp96951 := Call(__e, tmp96948, symlist, tmp96950)

																																					__e.TailApply(tmp96947, Y, tmp96951, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				tmp96952 := Call(__e, tmp96945, X, A, V2782, V2783, tmp96946)

																																				__e.TailApply(tmp96940, tmp96952)
																																				return

																																			}, 1)

																																			tmp96953 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																			tmp96954 := Call(__e, tmp96953, V2783)

																																			__e.TailApply(tmp96935, tmp96954)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp96997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp96998 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp96999 := Call(__e, tmp96998, V2625)

																																tmp97000 := Call(__e, tmp96997, tmp96999, V2783)

																																tmp97001 := Call(__e, tmp96932, tmp97000)

																																__e.TailApply(tmp96929, tmp97001)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp97075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp97076 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																													tmp97077 := Call(__e, tmp97076, V2625)

																													tmp97078 := Call(__e, tmp97075, tmp97077, V2783)

																													__e.TailApply(tmp96924, tmp97078)
																													return

																												} else {
																													tmp96922 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																													tmp96923 := Call(__e, tmp96922, V2625)

																													if True == tmp96923 {
																														tmp96900 := MakeNative(func(__e *ControlFlow) {
																															A := __e.Get(1)
																															_ = A
																															tmp96901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																															tmp96902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp96903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp96904 := Call(__e, tmp96903, A, Nil)

																															tmp96905 := Call(__e, tmp96902, symlist, tmp96904)

																															tmp96906 := Call(__e, tmp96901, V2625, tmp96905, V2783)

																															_ = tmp96906

																															tmp96907 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp96908 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																tmp96909 := Call(__e, tmp96908, V2625, V2783)

																																_ = tmp96909

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp96910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																															tmp96911 := Call(__e, tmp96910)

																															_ = tmp96911

																															tmp96912 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																															tmp96913 := MakeNative(func(__e *ControlFlow) {
																																tmp96914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																																tmp96915 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp96916 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp96917 := Call(__e, tmp96916, A, Nil)

																																tmp96918 := Call(__e, tmp96915, symlist, tmp96917)

																																__e.TailApply(tmp96914, Y, tmp96918, V2782, V2783, V2784)
																																return

																															}, 0)

																															tmp96919 := Call(__e, tmp96912, X, A, V2782, V2783, tmp96913)

																															__e.TailApply(tmp96907, tmp96919)
																															return

																														}, 1)

																														tmp96920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																														tmp96921 := Call(__e, tmp96920, V2783)

																														__e.TailApply(tmp96900, tmp96921)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}

																											}, 1)

																											tmp97081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp97082 := Call(__e, tmp97081, V2781, V2783)

																											__e.TailApply(tmp96897, tmp97082)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp97085 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																									tmp97086 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																									tmp97087 := Call(__e, tmp97086, V2623)

																									tmp97088 := Call(__e, tmp97085, tmp97087, V2783)

																									__e.TailApply(tmp96895, tmp97088)
																									return

																								}, 1)

																								tmp97089 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								tmp97090 := Call(__e, tmp97089, V2623)

																								__e.TailApply(tmp96894, tmp97090)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp97093 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp97094 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp97095 := Call(__e, tmp97094, V2622)

																						tmp97096 := Call(__e, tmp97093, tmp97095, V2783)

																						__e.TailApply(tmp96892, tmp97096)
																						return

																					}, 1)

																					tmp97097 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																					tmp97098 := Call(__e, tmp97097, V2622)

																					__e.TailApply(tmp96891, tmp97098)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp97101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp97102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp97103 := Call(__e, tmp97102, V2620)

																			tmp97104 := Call(__e, tmp97101, tmp97103, V2783)

																			__e.TailApply(tmp96889, tmp97104)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp97107 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp97108 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp97109 := Call(__e, tmp97108, V2620)

																	tmp97110 := Call(__e, tmp97107, tmp97109, V2783)

																	__e.TailApply(tmp96887, tmp97110)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp97113 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp97114 := Call(__e, tmp97113, V2780, V2783)

															tmp97115 := Call(__e, tmp96885, tmp97114)

															__e.TailApply(tmp95111, tmp97115)
															return

														} else {
															__e.Return(Case)
															return
														}

													}, 1)

													tmp97118 := MakeNative(func(__e *ControlFlow) {
														V2617 := __e.Get(1)
														_ = V2617
														tmp97156 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp97157 := Call(__e, tmp97156, V2617)

														if True == tmp97157 {
															tmp97120 := MakeNative(func(__e *ControlFlow) {
																F := __e.Get(1)
																_ = F
																tmp97121 := MakeNative(func(__e *ControlFlow) {
																	V2618 := __e.Get(1)
																	_ = V2618
																	tmp97148 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp97149 := Call(__e, tmp97148, V2618)

																	if True == tmp97149 {
																		tmp97123 := MakeNative(func(__e *ControlFlow) {
																			X := __e.Get(1)
																			_ = X
																			tmp97124 := MakeNative(func(__e *ControlFlow) {
																				V2619 := __e.Get(1)
																				_ = V2619
																				tmp97140 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp97141 := Call(__e, tmp97140, Nil, V2619)

																				if True == tmp97141 {
																					tmp97126 := MakeNative(func(__e *ControlFlow) {
																						B := __e.Get(1)
																						_ = B
																						tmp97127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																						tmp97128 := Call(__e, tmp97127)

																						_ = tmp97128

																						tmp97129 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																						tmp97130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp97131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp97132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp97133 := Call(__e, tmp97132, V2781, Nil)

																						tmp97134 := Call(__e, tmp97131, sym_1_1_6, tmp97133)

																						tmp97135 := Call(__e, tmp97130, B, tmp97134)

																						tmp97136 := MakeNative(func(__e *ControlFlow) {
																							tmp97137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																							__e.TailApply(tmp97137, X, B, V2782, V2783, V2784)
																							return

																						}, 0)

																						__e.TailApply(tmp97129, F, tmp97135, V2782, V2783, tmp97136)
																						return

																					}, 1)

																					tmp97138 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																					tmp97139 := Call(__e, tmp97138, V2783)

																					__e.TailApply(tmp97126, tmp97139)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp97142 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp97143 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp97144 := Call(__e, tmp97143, V2618)

																			tmp97145 := Call(__e, tmp97142, tmp97144, V2783)

																			__e.TailApply(tmp97124, tmp97145)
																			return

																		}, 1)

																		tmp97146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp97147 := Call(__e, tmp97146, V2618)

																		__e.TailApply(tmp97123, tmp97147)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp97150 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																tmp97151 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp97152 := Call(__e, tmp97151, V2617)

																tmp97153 := Call(__e, tmp97150, tmp97152, V2783)

																__e.TailApply(tmp97121, tmp97153)
																return

															}, 1)

															tmp97154 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp97155 := Call(__e, tmp97154, V2617)

															__e.TailApply(tmp97120, tmp97155)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp97158 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp97159 := Call(__e, tmp97158, V2780, V2783)

													tmp97160 := Call(__e, tmp97118, tmp97159)

													__e.TailApply(tmp95109, tmp97160)
													return

												} else {
													__e.Return(Case)
													return
												}

											}, 1)

											tmp97163 := MakeNative(func(__e *ControlFlow) {
												V2615 := __e.Get(1)
												_ = V2615
												tmp97183 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp97184 := Call(__e, tmp97183, V2615)

												if True == tmp97184 {
													tmp97165 := MakeNative(func(__e *ControlFlow) {
														F := __e.Get(1)
														_ = F
														tmp97166 := MakeNative(func(__e *ControlFlow) {
															V2616 := __e.Get(1)
															_ = V2616
															tmp97175 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp97176 := Call(__e, tmp97175, Nil, V2616)

															if True == tmp97176 {
																tmp97168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																tmp97169 := Call(__e, tmp97168)

																_ = tmp97169

																tmp97170 := Call(__e, PrimNS1Value(symns2_1value), symshen_4th_d)

																tmp97171 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp97172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp97173 := Call(__e, tmp97172, V2781, Nil)

																tmp97174 := Call(__e, tmp97171, sym_1_1_6, tmp97173)

																__e.TailApply(tmp97170, F, tmp97174, V2782, V2783, V2784)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp97177 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp97178 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp97179 := Call(__e, tmp97178, V2615)

														tmp97180 := Call(__e, tmp97177, tmp97179, V2783)

														__e.TailApply(tmp97166, tmp97180)
														return

													}, 1)

													tmp97181 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp97182 := Call(__e, tmp97181, V2615)

													__e.TailApply(tmp97165, tmp97182)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp97185 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp97186 := Call(__e, tmp97185, V2780, V2783)

											tmp97187 := Call(__e, tmp97163, tmp97186)

											__e.TailApply(tmp95107, tmp97187)
											return

										} else {
											__e.Return(Case)
											return
										}

									}, 1)

									tmp97190 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

									tmp97191 := Call(__e, tmp97190)

									_ = tmp97191

									tmp97192 := Call(__e, PrimNS1Value(symns2_1value), symshen_4by__hypothesis)

									tmp97193 := Call(__e, tmp97192, V2780, V2781, V2782, V2783, V2784)

									__e.TailApply(tmp95105, tmp97193)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp97196 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

							tmp97197 := Call(__e, tmp97196)

							_ = tmp97197

							tmp97198 := Call(__e, PrimNS1Value(symns2_1value), symshen_4base)

							tmp97199 := Call(__e, tmp97198, V2780, V2781, V2783, V2784)

							__e.TailApply(tmp95103, tmp97199)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp97202 := MakeNative(func(__e *ControlFlow) {
						F := __e.Get(1)
						_ = F
						tmp97203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

						tmp97204 := Call(__e, tmp97203)

						_ = tmp97204

						tmp97205 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

						tmp97206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typedf_2)

						tmp97207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp97208 := Call(__e, tmp97207, V2780, V2783)

						tmp97209 := Call(__e, tmp97206, tmp97208)

						tmp97210 := MakeNative(func(__e *ControlFlow) {
							tmp97211 := Call(__e, PrimNS1Value(symns2_1value), symbind)

							tmp97212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sigf)

							tmp97213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp97214 := Call(__e, tmp97213, V2780, V2783)

							tmp97215 := Call(__e, tmp97212, tmp97214)

							tmp97216 := MakeNative(func(__e *ControlFlow) {
								tmp97217 := Call(__e, PrimNS1Value(symns2_1value), symcall)

								tmp97218 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp97219 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp97220 := Call(__e, tmp97219, V2781, Nil)

								tmp97221 := Call(__e, tmp97218, F, tmp97220)

								__e.TailApply(tmp97217, tmp97221, V2783, V2784)
								return

							}, 0)

							__e.TailApply(tmp97211, F, tmp97215, V2783, tmp97216)
							return

						}, 0)

						__e.TailApply(tmp97205, tmp97209, V2783, tmp97210)
						return

					}, 1)

					tmp97222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

					tmp97223 := Call(__e, tmp97222, V2783)

					tmp97224 := Call(__e, tmp97202, tmp97223)

					__e.TailApply(tmp95101, tmp97224)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp97227 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp97228 := Call(__e, tmp97227)

			_ = tmp97228

			tmp97229 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show)

			tmp97230 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp97231 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp97232 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp97233 := Call(__e, tmp97232, V2781, Nil)

			tmp97234 := Call(__e, tmp97231, sym_h, tmp97233)

			tmp97235 := Call(__e, tmp97230, V2780, tmp97234)

			tmp97236 := MakeNative(func(__e *ControlFlow) {
				tmp97237 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

				__e.TailApply(tmp97237, False, V2783, V2784)
				return

			}, 0)

			tmp97238 := Call(__e, tmp97229, tmp97235, V2782, V2783, tmp97236)

			tmp97239 := Call(__e, tmp95099, tmp97238)

			__e.TailApply(tmp95098, Throwcontrol, tmp97239)
			return

		}, 1)

		tmp97240 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catchpoint)

		tmp97241 := Call(__e, tmp97240)

		__e.TailApply(tmp95097, tmp97241)
		return

	}, 5)

	tmp97242 := Call(__e, PrimNS1Value(symns2_1set), symshen_4th_d, tmp95096)

	_ = tmp97242

	tmp97243 := MakeNative(func(__e *ControlFlow) {
		V2789 := __e.Get(1)
		_ = V2789
		V2790 := __e.Get(2)
		_ = V2790
		V2791 := __e.Get(3)
		_ = V2791
		V2792 := __e.Get(4)
		_ = V2792
		tmp97244 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp99255 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp99256 := Call(__e, tmp99255, Case, False)

			if True == tmp99256 {
				tmp97246 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp98363 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp98364 := Call(__e, tmp98363, Case, False)

					if True == tmp98364 {
						tmp97248 := MakeNative(func(__e *ControlFlow) {
							Case := __e.Get(1)
							_ = Case
							tmp97528 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp97529 := Call(__e, tmp97528, Case, False)

							if True == tmp97529 {
								tmp97250 := MakeNative(func(__e *ControlFlow) {
									Case := __e.Get(1)
									_ = Case
									tmp97278 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp97279 := Call(__e, tmp97278, Case, False)

									if True == tmp97279 {
										tmp97252 := MakeNative(func(__e *ControlFlow) {
											V2610 := __e.Get(1)
											_ = V2610
											tmp97274 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp97275 := Call(__e, tmp97274, V2610)

											if True == tmp97275 {
												tmp97254 := MakeNative(func(__e *ControlFlow) {
													X := __e.Get(1)
													_ = X
													tmp97255 := MakeNative(func(__e *ControlFlow) {
														Hyp := __e.Get(1)
														_ = Hyp
														tmp97256 := MakeNative(func(__e *ControlFlow) {
															NewHyps := __e.Get(1)
															_ = NewHyps
															tmp97257 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

															tmp97258 := Call(__e, tmp97257)

															_ = tmp97258

															tmp97259 := Call(__e, PrimNS1Value(symns2_1value), symbind)

															tmp97260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp97261 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp97262 := Call(__e, tmp97261, X, V2791)

															tmp97263 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp97264 := Call(__e, tmp97263, NewHyps, V2791)

															tmp97265 := Call(__e, tmp97260, tmp97262, tmp97264)

															tmp97266 := MakeNative(func(__e *ControlFlow) {
																tmp97267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1hyps)

																__e.TailApply(tmp97267, Hyp, NewHyps, V2791, V2792)
																return

															}, 0)

															__e.TailApply(tmp97259, V2790, tmp97265, V2791, tmp97266)
															return

														}, 1)

														tmp97268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

														tmp97269 := Call(__e, tmp97268, V2791)

														__e.TailApply(tmp97256, tmp97269)
														return

													}, 1)

													tmp97270 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp97271 := Call(__e, tmp97270, V2610)

													__e.TailApply(tmp97255, tmp97271)
													return

												}, 1)

												tmp97272 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp97273 := Call(__e, tmp97272, V2610)

												__e.TailApply(tmp97254, tmp97273)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp97276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp97277 := Call(__e, tmp97276, V2789, V2791)

										__e.TailApply(tmp97252, tmp97277)
										return

									} else {
										__e.Return(Case)
										return
									}

								}, 1)

								tmp97280 := MakeNative(func(__e *ControlFlow) {
									V2597 := __e.Get(1)
									_ = V2597
									tmp97523 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp97524 := Call(__e, tmp97523, V2597)

									if True == tmp97524 {
										tmp97282 := MakeNative(func(__e *ControlFlow) {
											V2598 := __e.Get(1)
											_ = V2598
											tmp97517 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp97518 := Call(__e, tmp97517, V2598)

											if True == tmp97518 {
												tmp97284 := MakeNative(func(__e *ControlFlow) {
													V2599 := __e.Get(1)
													_ = V2599
													tmp97511 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp97512 := Call(__e, tmp97511, V2599)

													if True == tmp97512 {
														tmp97286 := MakeNative(func(__e *ControlFlow) {
															V2600 := __e.Get(1)
															_ = V2600
															tmp97505 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp97506 := Call(__e, tmp97505, sym_8s, V2600)

															if True == tmp97506 {
																tmp97288 := MakeNative(func(__e *ControlFlow) {
																	V2601 := __e.Get(1)
																	_ = V2601
																	tmp97499 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp97500 := Call(__e, tmp97499, V2601)

																	if True == tmp97500 {
																		tmp97290 := MakeNative(func(__e *ControlFlow) {
																			X := __e.Get(1)
																			_ = X
																			tmp97291 := MakeNative(func(__e *ControlFlow) {
																				V2602 := __e.Get(1)
																				_ = V2602
																				tmp97491 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																				tmp97492 := Call(__e, tmp97491, V2602)

																				if True == tmp97492 {
																					tmp97293 := MakeNative(func(__e *ControlFlow) {
																						Y := __e.Get(1)
																						_ = Y
																						tmp97294 := MakeNative(func(__e *ControlFlow) {
																							V2603 := __e.Get(1)
																							_ = V2603
																							tmp97483 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp97484 := Call(__e, tmp97483, Nil, V2603)

																							if True == tmp97484 {
																								tmp97296 := MakeNative(func(__e *ControlFlow) {
																									V2604 := __e.Get(1)
																									_ = V2604
																									tmp97477 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																									tmp97478 := Call(__e, tmp97477, V2604)

																									if True == tmp97478 {
																										tmp97298 := MakeNative(func(__e *ControlFlow) {
																											V2605 := __e.Get(1)
																											_ = V2605
																											tmp97471 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											tmp97472 := Call(__e, tmp97471, sym_h, V2605)

																											if True == tmp97472 {
																												tmp97300 := MakeNative(func(__e *ControlFlow) {
																													V2606 := __e.Get(1)
																													_ = V2606
																													tmp97465 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																													tmp97466 := Call(__e, tmp97465, V2606)

																													if True == tmp97466 {
																														tmp97302 := MakeNative(func(__e *ControlFlow) {
																															V2607 := __e.Get(1)
																															_ = V2607
																															tmp97459 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																															tmp97460 := Call(__e, tmp97459, symstring, V2607)

																															if True == tmp97460 {
																																tmp97386 := MakeNative(func(__e *ControlFlow) {
																																	V2608 := __e.Get(1)
																																	_ = V2608
																																	tmp97453 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																	tmp97454 := Call(__e, tmp97453, Nil, V2608)

																																	if True == tmp97454 {
																																		tmp97425 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp97426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp97427 := Call(__e, tmp97426)

																																			_ = tmp97427

																																			tmp97428 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																			tmp97429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97430 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97431 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97432 := Call(__e, tmp97431, X, V2791)

																																			tmp97433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97434 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97435 := Call(__e, tmp97434, symstring, Nil)

																																			tmp97436 := Call(__e, tmp97433, sym_h, tmp97435)

																																			tmp97437 := Call(__e, tmp97430, tmp97432, tmp97436)

																																			tmp97438 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97439 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97440 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97441 := Call(__e, tmp97440, Y, V2791)

																																			tmp97442 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97443 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97444 := Call(__e, tmp97443, symstring, Nil)

																																			tmp97445 := Call(__e, tmp97442, sym_h, tmp97444)

																																			tmp97446 := Call(__e, tmp97439, tmp97441, tmp97445)

																																			tmp97447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97448 := Call(__e, tmp97447, Hyp, V2791)

																																			tmp97449 := Call(__e, tmp97438, tmp97446, tmp97448)

																																			tmp97450 := Call(__e, tmp97429, tmp97437, tmp97449)

																																			__e.TailApply(tmp97428, V2790, tmp97450, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp97451 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp97452 := Call(__e, tmp97451, V2597)

																																		__e.TailApply(tmp97425, tmp97452)
																																		return

																																	} else {
																																		tmp97423 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp97424 := Call(__e, tmp97423, V2608)

																																		if True == tmp97424 {
																																			tmp97389 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp97390 := Call(__e, tmp97389, V2608, Nil, V2791)

																																			_ = tmp97390

																																			tmp97391 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp97392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp97393 := Call(__e, tmp97392, V2608, V2791)

																																				_ = tmp97393

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp97394 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp97395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp97396 := Call(__e, tmp97395)

																																				_ = tmp97396

																																				tmp97397 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp97398 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97399 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97401 := Call(__e, tmp97400, X, V2791)

																																				tmp97402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97403 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97404 := Call(__e, tmp97403, symstring, Nil)

																																				tmp97405 := Call(__e, tmp97402, sym_h, tmp97404)

																																				tmp97406 := Call(__e, tmp97399, tmp97401, tmp97405)

																																				tmp97407 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97408 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97410 := Call(__e, tmp97409, Y, V2791)

																																				tmp97411 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97412 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97413 := Call(__e, tmp97412, symstring, Nil)

																																				tmp97414 := Call(__e, tmp97411, sym_h, tmp97413)

																																				tmp97415 := Call(__e, tmp97408, tmp97410, tmp97414)

																																				tmp97416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97417 := Call(__e, tmp97416, Hyp, V2791)

																																				tmp97418 := Call(__e, tmp97407, tmp97415, tmp97417)

																																				tmp97419 := Call(__e, tmp97398, tmp97406, tmp97418)

																																				__e.TailApply(tmp97397, V2790, tmp97419, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp97420 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp97421 := Call(__e, tmp97420, V2597)

																																			tmp97422 := Call(__e, tmp97394, tmp97421)

																																			__e.TailApply(tmp97391, tmp97422)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp97455 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp97456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp97457 := Call(__e, tmp97456, V2606)

																																tmp97458 := Call(__e, tmp97455, tmp97457, V2791)

																																__e.TailApply(tmp97386, tmp97458)
																																return

																															} else {
																																tmp97384 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																tmp97385 := Call(__e, tmp97384, V2607)

																																if True == tmp97385 {
																																	tmp97305 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																	tmp97306 := Call(__e, tmp97305, V2607, symstring, V2791)

																																	_ = tmp97306

																																	tmp97307 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp97308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																		tmp97309 := Call(__e, tmp97308, V2607, V2791)

																																		_ = tmp97309

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp97310 := MakeNative(func(__e *ControlFlow) {
																																		V2609 := __e.Get(1)
																																		_ = V2609
																																		tmp97377 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp97378 := Call(__e, tmp97377, Nil, V2609)

																																		if True == tmp97378 {
																																			tmp97349 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp97350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp97351 := Call(__e, tmp97350)

																																				_ = tmp97351

																																				tmp97352 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp97353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97355 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97356 := Call(__e, tmp97355, X, V2791)

																																				tmp97357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97359 := Call(__e, tmp97358, symstring, Nil)

																																				tmp97360 := Call(__e, tmp97357, sym_h, tmp97359)

																																				tmp97361 := Call(__e, tmp97354, tmp97356, tmp97360)

																																				tmp97362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97365 := Call(__e, tmp97364, Y, V2791)

																																				tmp97366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97367 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97368 := Call(__e, tmp97367, symstring, Nil)

																																				tmp97369 := Call(__e, tmp97366, sym_h, tmp97368)

																																				tmp97370 := Call(__e, tmp97363, tmp97365, tmp97369)

																																				tmp97371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97372 := Call(__e, tmp97371, Hyp, V2791)

																																				tmp97373 := Call(__e, tmp97362, tmp97370, tmp97372)

																																				tmp97374 := Call(__e, tmp97353, tmp97361, tmp97373)

																																				__e.TailApply(tmp97352, V2790, tmp97374, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp97375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp97376 := Call(__e, tmp97375, V2597)

																																			__e.TailApply(tmp97349, tmp97376)
																																			return

																																		} else {
																																			tmp97347 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp97348 := Call(__e, tmp97347, V2609)

																																			if True == tmp97348 {
																																				tmp97313 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp97314 := Call(__e, tmp97313, V2609, Nil, V2791)

																																				_ = tmp97314

																																				tmp97315 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp97316 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp97317 := Call(__e, tmp97316, V2609, V2791)

																																					_ = tmp97317

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp97318 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp97319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp97320 := Call(__e, tmp97319)

																																					_ = tmp97320

																																					tmp97321 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp97322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97324 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp97325 := Call(__e, tmp97324, X, V2791)

																																					tmp97326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97328 := Call(__e, tmp97327, symstring, Nil)

																																					tmp97329 := Call(__e, tmp97326, sym_h, tmp97328)

																																					tmp97330 := Call(__e, tmp97323, tmp97325, tmp97329)

																																					tmp97331 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp97334 := Call(__e, tmp97333, Y, V2791)

																																					tmp97335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97337 := Call(__e, tmp97336, symstring, Nil)

																																					tmp97338 := Call(__e, tmp97335, sym_h, tmp97337)

																																					tmp97339 := Call(__e, tmp97332, tmp97334, tmp97338)

																																					tmp97340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp97341 := Call(__e, tmp97340, Hyp, V2791)

																																					tmp97342 := Call(__e, tmp97331, tmp97339, tmp97341)

																																					tmp97343 := Call(__e, tmp97322, tmp97330, tmp97342)

																																					__e.TailApply(tmp97321, V2790, tmp97343, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp97344 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp97345 := Call(__e, tmp97344, V2597)

																																				tmp97346 := Call(__e, tmp97318, tmp97345)

																																				__e.TailApply(tmp97315, tmp97346)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp97379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp97380 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp97381 := Call(__e, tmp97380, V2606)

																																	tmp97382 := Call(__e, tmp97379, tmp97381, V2791)

																																	tmp97383 := Call(__e, tmp97310, tmp97382)

																																	__e.TailApply(tmp97307, tmp97383)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}

																														}, 1)

																														tmp97461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																														tmp97462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp97463 := Call(__e, tmp97462, V2606)

																														tmp97464 := Call(__e, tmp97461, tmp97463, V2791)

																														__e.TailApply(tmp97302, tmp97464)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp97467 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																												tmp97468 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																												tmp97469 := Call(__e, tmp97468, V2604)

																												tmp97470 := Call(__e, tmp97467, tmp97469, V2791)

																												__e.TailApply(tmp97300, tmp97470)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp97473 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																										tmp97474 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										tmp97475 := Call(__e, tmp97474, V2604)

																										tmp97476 := Call(__e, tmp97473, tmp97475, V2791)

																										__e.TailApply(tmp97298, tmp97476)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp97479 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																								tmp97480 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																								tmp97481 := Call(__e, tmp97480, V2598)

																								tmp97482 := Call(__e, tmp97479, tmp97481, V2791)

																								__e.TailApply(tmp97296, tmp97482)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp97485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp97486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp97487 := Call(__e, tmp97486, V2602)

																						tmp97488 := Call(__e, tmp97485, tmp97487, V2791)

																						__e.TailApply(tmp97294, tmp97488)
																						return

																					}, 1)

																					tmp97489 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																					tmp97490 := Call(__e, tmp97489, V2602)

																					__e.TailApply(tmp97293, tmp97490)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp97493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp97494 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp97495 := Call(__e, tmp97494, V2601)

																			tmp97496 := Call(__e, tmp97493, tmp97495, V2791)

																			__e.TailApply(tmp97291, tmp97496)
																			return

																		}, 1)

																		tmp97497 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp97498 := Call(__e, tmp97497, V2601)

																		__e.TailApply(tmp97290, tmp97498)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp97501 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																tmp97502 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp97503 := Call(__e, tmp97502, V2599)

																tmp97504 := Call(__e, tmp97501, tmp97503, V2791)

																__e.TailApply(tmp97288, tmp97504)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp97507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp97508 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp97509 := Call(__e, tmp97508, V2599)

														tmp97510 := Call(__e, tmp97507, tmp97509, V2791)

														__e.TailApply(tmp97286, tmp97510)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp97513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp97514 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp97515 := Call(__e, tmp97514, V2598)

												tmp97516 := Call(__e, tmp97513, tmp97515, V2791)

												__e.TailApply(tmp97284, tmp97516)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp97519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp97520 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp97521 := Call(__e, tmp97520, V2597)

										tmp97522 := Call(__e, tmp97519, tmp97521, V2791)

										__e.TailApply(tmp97282, tmp97522)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp97525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp97526 := Call(__e, tmp97525, V2789, V2791)

								tmp97527 := Call(__e, tmp97280, tmp97526)

								__e.TailApply(tmp97250, tmp97527)
								return

							} else {
								__e.Return(Case)
								return
							}

						}, 1)

						tmp97530 := MakeNative(func(__e *ControlFlow) {
							V2574 := __e.Get(1)
							_ = V2574
							tmp98358 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp98359 := Call(__e, tmp98358, V2574)

							if True == tmp98359 {
								tmp97532 := MakeNative(func(__e *ControlFlow) {
									V2575 := __e.Get(1)
									_ = V2575
									tmp98352 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp98353 := Call(__e, tmp98352, V2575)

									if True == tmp98353 {
										tmp97534 := MakeNative(func(__e *ControlFlow) {
											V2576 := __e.Get(1)
											_ = V2576
											tmp98346 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp98347 := Call(__e, tmp98346, V2576)

											if True == tmp98347 {
												tmp97536 := MakeNative(func(__e *ControlFlow) {
													V2577 := __e.Get(1)
													_ = V2577
													tmp98340 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

													tmp98341 := Call(__e, tmp98340, sym_8v, V2577)

													if True == tmp98341 {
														tmp97538 := MakeNative(func(__e *ControlFlow) {
															V2578 := __e.Get(1)
															_ = V2578
															tmp98334 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

															tmp98335 := Call(__e, tmp98334, V2578)

															if True == tmp98335 {
																tmp97540 := MakeNative(func(__e *ControlFlow) {
																	X := __e.Get(1)
																	_ = X
																	tmp97541 := MakeNative(func(__e *ControlFlow) {
																		V2579 := __e.Get(1)
																		_ = V2579
																		tmp98326 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp98327 := Call(__e, tmp98326, V2579)

																		if True == tmp98327 {
																			tmp97543 := MakeNative(func(__e *ControlFlow) {
																				Y := __e.Get(1)
																				_ = Y
																				tmp97544 := MakeNative(func(__e *ControlFlow) {
																					V2580 := __e.Get(1)
																					_ = V2580
																					tmp98318 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp98319 := Call(__e, tmp98318, Nil, V2580)

																					if True == tmp98319 {
																						tmp97546 := MakeNative(func(__e *ControlFlow) {
																							V2581 := __e.Get(1)
																							_ = V2581
																							tmp98312 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																							tmp98313 := Call(__e, tmp98312, V2581)

																							if True == tmp98313 {
																								tmp97548 := MakeNative(func(__e *ControlFlow) {
																									V2582 := __e.Get(1)
																									_ = V2582
																									tmp98306 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																									tmp98307 := Call(__e, tmp98306, sym_h, V2582)

																									if True == tmp98307 {
																										tmp97550 := MakeNative(func(__e *ControlFlow) {
																											V2583 := __e.Get(1)
																											_ = V2583
																											tmp98300 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																											tmp98301 := Call(__e, tmp98300, V2583)

																											if True == tmp98301 {
																												tmp97552 := MakeNative(func(__e *ControlFlow) {
																													V2584 := __e.Get(1)
																													_ = V2584
																													tmp98294 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																													tmp98295 := Call(__e, tmp98294, V2584)

																													if True == tmp98295 {
																														tmp97659 := MakeNative(func(__e *ControlFlow) {
																															V2585 := __e.Get(1)
																															_ = V2585
																															tmp98288 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																															tmp98289 := Call(__e, tmp98288, symvector, V2585)

																															if True == tmp98289 {
																																tmp97979 := MakeNative(func(__e *ControlFlow) {
																																	V2586 := __e.Get(1)
																																	_ = V2586
																																	tmp98282 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																	tmp98283 := Call(__e, tmp98282, V2586)

																																	if True == tmp98283 {
																																		tmp98084 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp98085 := MakeNative(func(__e *ControlFlow) {
																																				V2587 := __e.Get(1)
																																				_ = V2587
																																				tmp98274 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																				tmp98275 := Call(__e, tmp98274, Nil, V2587)

																																				if True == tmp98275 {
																																					tmp98185 := MakeNative(func(__e *ControlFlow) {
																																						V2588 := __e.Get(1)
																																						_ = V2588
																																						tmp98268 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp98269 := Call(__e, tmp98268, Nil, V2588)

																																						if True == tmp98269 {
																																							tmp98232 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp98233 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp98234 := Call(__e, tmp98233)

																																								_ = tmp98234

																																								tmp98235 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																								tmp98236 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98237 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98239 := Call(__e, tmp98238, X, V2791)

																																								tmp98240 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98241 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98242 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98243 := Call(__e, tmp98242, A, V2791)

																																								tmp98244 := Call(__e, tmp98241, tmp98243, Nil)

																																								tmp98245 := Call(__e, tmp98240, sym_h, tmp98244)

																																								tmp98246 := Call(__e, tmp98237, tmp98239, tmp98245)

																																								tmp98247 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98248 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98249 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98250 := Call(__e, tmp98249, Y, V2791)

																																								tmp98251 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98252 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98253 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98254 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98255 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98256 := Call(__e, tmp98255, A, V2791)

																																								tmp98257 := Call(__e, tmp98254, tmp98256, Nil)

																																								tmp98258 := Call(__e, tmp98253, symvector, tmp98257)

																																								tmp98259 := Call(__e, tmp98252, tmp98258, Nil)

																																								tmp98260 := Call(__e, tmp98251, sym_h, tmp98259)

																																								tmp98261 := Call(__e, tmp98248, tmp98250, tmp98260)

																																								tmp98262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98263 := Call(__e, tmp98262, Hyp, V2791)

																																								tmp98264 := Call(__e, tmp98247, tmp98261, tmp98263)

																																								tmp98265 := Call(__e, tmp98236, tmp98246, tmp98264)

																																								__e.TailApply(tmp98235, V2790, tmp98265, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp98266 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp98267 := Call(__e, tmp98266, V2574)

																																							__e.TailApply(tmp98232, tmp98267)
																																							return

																																						} else {
																																							tmp98230 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																							tmp98231 := Call(__e, tmp98230, V2588)

																																							if True == tmp98231 {
																																								tmp98188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp98189 := Call(__e, tmp98188, V2588, Nil, V2791)

																																								_ = tmp98189

																																								tmp98190 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp98191 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp98192 := Call(__e, tmp98191, V2588, V2791)

																																									_ = tmp98192

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp98193 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp98194 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp98195 := Call(__e, tmp98194)

																																									_ = tmp98195

																																									tmp98196 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp98197 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98198 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98200 := Call(__e, tmp98199, X, V2791)

																																									tmp98201 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98202 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98203 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98204 := Call(__e, tmp98203, A, V2791)

																																									tmp98205 := Call(__e, tmp98202, tmp98204, Nil)

																																									tmp98206 := Call(__e, tmp98201, sym_h, tmp98205)

																																									tmp98207 := Call(__e, tmp98198, tmp98200, tmp98206)

																																									tmp98208 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98209 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98211 := Call(__e, tmp98210, Y, V2791)

																																									tmp98212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98213 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98214 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98215 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98217 := Call(__e, tmp98216, A, V2791)

																																									tmp98218 := Call(__e, tmp98215, tmp98217, Nil)

																																									tmp98219 := Call(__e, tmp98214, symvector, tmp98218)

																																									tmp98220 := Call(__e, tmp98213, tmp98219, Nil)

																																									tmp98221 := Call(__e, tmp98212, sym_h, tmp98220)

																																									tmp98222 := Call(__e, tmp98209, tmp98211, tmp98221)

																																									tmp98223 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98224 := Call(__e, tmp98223, Hyp, V2791)

																																									tmp98225 := Call(__e, tmp98208, tmp98222, tmp98224)

																																									tmp98226 := Call(__e, tmp98197, tmp98207, tmp98225)

																																									__e.TailApply(tmp98196, V2790, tmp98226, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp98227 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp98228 := Call(__e, tmp98227, V2574)

																																								tmp98229 := Call(__e, tmp98193, tmp98228)

																																								__e.TailApply(tmp98190, tmp98229)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp98270 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98271 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp98272 := Call(__e, tmp98271, V2583)

																																					tmp98273 := Call(__e, tmp98270, tmp98272, V2791)

																																					__e.TailApply(tmp98185, tmp98273)
																																					return

																																				} else {
																																					tmp98183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																					tmp98184 := Call(__e, tmp98183, V2587)

																																					if True == tmp98184 {
																																						tmp98088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																						tmp98089 := Call(__e, tmp98088, V2587, Nil, V2791)

																																						_ = tmp98089

																																						tmp98090 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp98091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																							tmp98092 := Call(__e, tmp98091, V2587, V2791)

																																							_ = tmp98092

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp98093 := MakeNative(func(__e *ControlFlow) {
																																							V2589 := __e.Get(1)
																																							_ = V2589
																																							tmp98176 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp98177 := Call(__e, tmp98176, Nil, V2589)

																																							if True == tmp98177 {
																																								tmp98140 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp98141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp98142 := Call(__e, tmp98141)

																																									_ = tmp98142

																																									tmp98143 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp98144 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98147 := Call(__e, tmp98146, X, V2791)

																																									tmp98148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98150 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98151 := Call(__e, tmp98150, A, V2791)

																																									tmp98152 := Call(__e, tmp98149, tmp98151, Nil)

																																									tmp98153 := Call(__e, tmp98148, sym_h, tmp98152)

																																									tmp98154 := Call(__e, tmp98145, tmp98147, tmp98153)

																																									tmp98155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98157 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98158 := Call(__e, tmp98157, Y, V2791)

																																									tmp98159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98164 := Call(__e, tmp98163, A, V2791)

																																									tmp98165 := Call(__e, tmp98162, tmp98164, Nil)

																																									tmp98166 := Call(__e, tmp98161, symvector, tmp98165)

																																									tmp98167 := Call(__e, tmp98160, tmp98166, Nil)

																																									tmp98168 := Call(__e, tmp98159, sym_h, tmp98167)

																																									tmp98169 := Call(__e, tmp98156, tmp98158, tmp98168)

																																									tmp98170 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98171 := Call(__e, tmp98170, Hyp, V2791)

																																									tmp98172 := Call(__e, tmp98155, tmp98169, tmp98171)

																																									tmp98173 := Call(__e, tmp98144, tmp98154, tmp98172)

																																									__e.TailApply(tmp98143, V2790, tmp98173, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp98174 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp98175 := Call(__e, tmp98174, V2574)

																																								__e.TailApply(tmp98140, tmp98175)
																																								return

																																							} else {
																																								tmp98138 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp98139 := Call(__e, tmp98138, V2589)

																																								if True == tmp98139 {
																																									tmp98096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp98097 := Call(__e, tmp98096, V2589, Nil, V2791)

																																									_ = tmp98097

																																									tmp98098 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp98099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp98100 := Call(__e, tmp98099, V2589, V2791)

																																										_ = tmp98100

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp98101 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp98102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp98103 := Call(__e, tmp98102)

																																										_ = tmp98103

																																										tmp98104 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp98105 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98106 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98107 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98108 := Call(__e, tmp98107, X, V2791)

																																										tmp98109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98112 := Call(__e, tmp98111, A, V2791)

																																										tmp98113 := Call(__e, tmp98110, tmp98112, Nil)

																																										tmp98114 := Call(__e, tmp98109, sym_h, tmp98113)

																																										tmp98115 := Call(__e, tmp98106, tmp98108, tmp98114)

																																										tmp98116 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98117 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98119 := Call(__e, tmp98118, Y, V2791)

																																										tmp98120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98123 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98125 := Call(__e, tmp98124, A, V2791)

																																										tmp98126 := Call(__e, tmp98123, tmp98125, Nil)

																																										tmp98127 := Call(__e, tmp98122, symvector, tmp98126)

																																										tmp98128 := Call(__e, tmp98121, tmp98127, Nil)

																																										tmp98129 := Call(__e, tmp98120, sym_h, tmp98128)

																																										tmp98130 := Call(__e, tmp98117, tmp98119, tmp98129)

																																										tmp98131 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98132 := Call(__e, tmp98131, Hyp, V2791)

																																										tmp98133 := Call(__e, tmp98116, tmp98130, tmp98132)

																																										tmp98134 := Call(__e, tmp98105, tmp98115, tmp98133)

																																										__e.TailApply(tmp98104, V2790, tmp98134, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp98135 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp98136 := Call(__e, tmp98135, V2574)

																																									tmp98137 := Call(__e, tmp98101, tmp98136)

																																									__e.TailApply(tmp98098, tmp98137)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp98178 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp98179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp98180 := Call(__e, tmp98179, V2583)

																																						tmp98181 := Call(__e, tmp98178, tmp98180, V2791)

																																						tmp98182 := Call(__e, tmp98093, tmp98181)

																																						__e.TailApply(tmp98090, tmp98182)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp98276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98277 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp98278 := Call(__e, tmp98277, V2586)

																																			tmp98279 := Call(__e, tmp98276, tmp98278, V2791)

																																			__e.TailApply(tmp98085, tmp98279)
																																			return

																																		}, 1)

																																		tmp98280 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																		tmp98281 := Call(__e, tmp98280, V2586)

																																		__e.TailApply(tmp98084, tmp98281)
																																		return

																																	} else {
																																		tmp98082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp98083 := Call(__e, tmp98082, V2586)

																																		if True == tmp98083 {
																																			tmp97982 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp97983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp97984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97985 := Call(__e, tmp97984, A, Nil)

																																				tmp97986 := Call(__e, tmp97983, V2586, tmp97985, V2791)

																																				_ = tmp97986

																																				tmp97987 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp97988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp97989 := Call(__e, tmp97988, V2586, V2791)

																																					_ = tmp97989

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp97990 := MakeNative(func(__e *ControlFlow) {
																																					V2590 := __e.Get(1)
																																					_ = V2590
																																					tmp98073 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																					tmp98074 := Call(__e, tmp98073, Nil, V2590)

																																					if True == tmp98074 {
																																						tmp98037 := MakeNative(func(__e *ControlFlow) {
																																							Hyp := __e.Get(1)
																																							_ = Hyp
																																							tmp98038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																							tmp98039 := Call(__e, tmp98038)

																																							_ = tmp98039

																																							tmp98040 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																							tmp98041 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98042 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98043 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98044 := Call(__e, tmp98043, X, V2791)

																																							tmp98045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98047 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98048 := Call(__e, tmp98047, A, V2791)

																																							tmp98049 := Call(__e, tmp98046, tmp98048, Nil)

																																							tmp98050 := Call(__e, tmp98045, sym_h, tmp98049)

																																							tmp98051 := Call(__e, tmp98042, tmp98044, tmp98050)

																																							tmp98052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98054 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98055 := Call(__e, tmp98054, Y, V2791)

																																							tmp98056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98057 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98058 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98059 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp98060 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98061 := Call(__e, tmp98060, A, V2791)

																																							tmp98062 := Call(__e, tmp98059, tmp98061, Nil)

																																							tmp98063 := Call(__e, tmp98058, symvector, tmp98062)

																																							tmp98064 := Call(__e, tmp98057, tmp98063, Nil)

																																							tmp98065 := Call(__e, tmp98056, sym_h, tmp98064)

																																							tmp98066 := Call(__e, tmp98053, tmp98055, tmp98065)

																																							tmp98067 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98068 := Call(__e, tmp98067, Hyp, V2791)

																																							tmp98069 := Call(__e, tmp98052, tmp98066, tmp98068)

																																							tmp98070 := Call(__e, tmp98041, tmp98051, tmp98069)

																																							__e.TailApply(tmp98040, V2790, tmp98070, V2791, V2792)
																																							return

																																						}, 1)

																																						tmp98071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp98072 := Call(__e, tmp98071, V2574)

																																						__e.TailApply(tmp98037, tmp98072)
																																						return

																																					} else {
																																						tmp98035 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																						tmp98036 := Call(__e, tmp98035, V2590)

																																						if True == tmp98036 {
																																							tmp97993 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																							tmp97994 := Call(__e, tmp97993, V2590, Nil, V2791)

																																							_ = tmp97994

																																							tmp97995 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp97996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																								tmp97997 := Call(__e, tmp97996, V2590, V2791)

																																								_ = tmp97997

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp97998 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp97999 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp98000 := Call(__e, tmp97999)

																																								_ = tmp98000

																																								tmp98001 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																								tmp98002 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98003 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98005 := Call(__e, tmp98004, X, V2791)

																																								tmp98006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98007 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98009 := Call(__e, tmp98008, A, V2791)

																																								tmp98010 := Call(__e, tmp98007, tmp98009, Nil)

																																								tmp98011 := Call(__e, tmp98006, sym_h, tmp98010)

																																								tmp98012 := Call(__e, tmp98003, tmp98005, tmp98011)

																																								tmp98013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98016 := Call(__e, tmp98015, Y, V2791)

																																								tmp98017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98019 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98022 := Call(__e, tmp98021, A, V2791)

																																								tmp98023 := Call(__e, tmp98020, tmp98022, Nil)

																																								tmp98024 := Call(__e, tmp98019, symvector, tmp98023)

																																								tmp98025 := Call(__e, tmp98018, tmp98024, Nil)

																																								tmp98026 := Call(__e, tmp98017, sym_h, tmp98025)

																																								tmp98027 := Call(__e, tmp98014, tmp98016, tmp98026)

																																								tmp98028 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98029 := Call(__e, tmp98028, Hyp, V2791)

																																								tmp98030 := Call(__e, tmp98013, tmp98027, tmp98029)

																																								tmp98031 := Call(__e, tmp98002, tmp98012, tmp98030)

																																								__e.TailApply(tmp98001, V2790, tmp98031, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp98032 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp98033 := Call(__e, tmp98032, V2574)

																																							tmp98034 := Call(__e, tmp97998, tmp98033)

																																							__e.TailApply(tmp97995, tmp98034)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp98075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98076 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp98077 := Call(__e, tmp98076, V2583)

																																				tmp98078 := Call(__e, tmp98075, tmp98077, V2791)

																																				tmp98079 := Call(__e, tmp97990, tmp98078)

																																				__e.TailApply(tmp97987, tmp98079)
																																				return

																																			}, 1)

																																			tmp98080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																			tmp98081 := Call(__e, tmp98080, V2791)

																																			__e.TailApply(tmp97982, tmp98081)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp98284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp98285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp98286 := Call(__e, tmp98285, V2584)

																																tmp98287 := Call(__e, tmp98284, tmp98286, V2791)

																																__e.TailApply(tmp97979, tmp98287)
																																return

																															} else {
																																tmp97977 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																tmp97978 := Call(__e, tmp97977, V2585)

																																if True == tmp97978 {
																																	tmp97662 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																	tmp97663 := Call(__e, tmp97662, V2585, symvector, V2791)

																																	_ = tmp97663

																																	tmp97664 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp97665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																		tmp97666 := Call(__e, tmp97665, V2585, V2791)

																																		_ = tmp97666

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp97667 := MakeNative(func(__e *ControlFlow) {
																																		V2591 := __e.Get(1)
																																		_ = V2591
																																		tmp97970 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																		tmp97971 := Call(__e, tmp97970, V2591)

																																		if True == tmp97971 {
																																			tmp97772 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp97773 := MakeNative(func(__e *ControlFlow) {
																																					V2592 := __e.Get(1)
																																					_ = V2592
																																					tmp97962 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																					tmp97963 := Call(__e, tmp97962, Nil, V2592)

																																					if True == tmp97963 {
																																						tmp97873 := MakeNative(func(__e *ControlFlow) {
																																							V2593 := __e.Get(1)
																																							_ = V2593
																																							tmp97956 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp97957 := Call(__e, tmp97956, Nil, V2593)

																																							if True == tmp97957 {
																																								tmp97920 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp97921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp97922 := Call(__e, tmp97921)

																																									_ = tmp97922

																																									tmp97923 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp97924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97927 := Call(__e, tmp97926, X, V2791)

																																									tmp97928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97931 := Call(__e, tmp97930, A, V2791)

																																									tmp97932 := Call(__e, tmp97929, tmp97931, Nil)

																																									tmp97933 := Call(__e, tmp97928, sym_h, tmp97932)

																																									tmp97934 := Call(__e, tmp97925, tmp97927, tmp97933)

																																									tmp97935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97937 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97938 := Call(__e, tmp97937, Y, V2791)

																																									tmp97939 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97940 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97941 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97942 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97944 := Call(__e, tmp97943, A, V2791)

																																									tmp97945 := Call(__e, tmp97942, tmp97944, Nil)

																																									tmp97946 := Call(__e, tmp97941, symvector, tmp97945)

																																									tmp97947 := Call(__e, tmp97940, tmp97946, Nil)

																																									tmp97948 := Call(__e, tmp97939, sym_h, tmp97947)

																																									tmp97949 := Call(__e, tmp97936, tmp97938, tmp97948)

																																									tmp97950 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97951 := Call(__e, tmp97950, Hyp, V2791)

																																									tmp97952 := Call(__e, tmp97935, tmp97949, tmp97951)

																																									tmp97953 := Call(__e, tmp97924, tmp97934, tmp97952)

																																									__e.TailApply(tmp97923, V2790, tmp97953, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp97954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp97955 := Call(__e, tmp97954, V2574)

																																								__e.TailApply(tmp97920, tmp97955)
																																								return

																																							} else {
																																								tmp97918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp97919 := Call(__e, tmp97918, V2593)

																																								if True == tmp97919 {
																																									tmp97876 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp97877 := Call(__e, tmp97876, V2593, Nil, V2791)

																																									_ = tmp97877

																																									tmp97878 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp97879 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp97880 := Call(__e, tmp97879, V2593, V2791)

																																										_ = tmp97880

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp97881 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp97882 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp97883 := Call(__e, tmp97882)

																																										_ = tmp97883

																																										tmp97884 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp97885 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97886 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97887 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97888 := Call(__e, tmp97887, X, V2791)

																																										tmp97889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97892 := Call(__e, tmp97891, A, V2791)

																																										tmp97893 := Call(__e, tmp97890, tmp97892, Nil)

																																										tmp97894 := Call(__e, tmp97889, sym_h, tmp97893)

																																										tmp97895 := Call(__e, tmp97886, tmp97888, tmp97894)

																																										tmp97896 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97897 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97898 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97899 := Call(__e, tmp97898, Y, V2791)

																																										tmp97900 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97901 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97902 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97905 := Call(__e, tmp97904, A, V2791)

																																										tmp97906 := Call(__e, tmp97903, tmp97905, Nil)

																																										tmp97907 := Call(__e, tmp97902, symvector, tmp97906)

																																										tmp97908 := Call(__e, tmp97901, tmp97907, Nil)

																																										tmp97909 := Call(__e, tmp97900, sym_h, tmp97908)

																																										tmp97910 := Call(__e, tmp97897, tmp97899, tmp97909)

																																										tmp97911 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97912 := Call(__e, tmp97911, Hyp, V2791)

																																										tmp97913 := Call(__e, tmp97896, tmp97910, tmp97912)

																																										tmp97914 := Call(__e, tmp97885, tmp97895, tmp97913)

																																										__e.TailApply(tmp97884, V2790, tmp97914, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp97915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp97916 := Call(__e, tmp97915, V2574)

																																									tmp97917 := Call(__e, tmp97881, tmp97916)

																																									__e.TailApply(tmp97878, tmp97917)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp97958 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp97959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp97960 := Call(__e, tmp97959, V2583)

																																						tmp97961 := Call(__e, tmp97958, tmp97960, V2791)

																																						__e.TailApply(tmp97873, tmp97961)
																																						return

																																					} else {
																																						tmp97871 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																						tmp97872 := Call(__e, tmp97871, V2592)

																																						if True == tmp97872 {
																																							tmp97776 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																							tmp97777 := Call(__e, tmp97776, V2592, Nil, V2791)

																																							_ = tmp97777

																																							tmp97778 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp97779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																								tmp97780 := Call(__e, tmp97779, V2592, V2791)

																																								_ = tmp97780

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp97781 := MakeNative(func(__e *ControlFlow) {
																																								V2594 := __e.Get(1)
																																								_ = V2594
																																								tmp97864 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp97865 := Call(__e, tmp97864, Nil, V2594)

																																								if True == tmp97865 {
																																									tmp97828 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp97829 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp97830 := Call(__e, tmp97829)

																																										_ = tmp97830

																																										tmp97831 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp97832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97834 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97835 := Call(__e, tmp97834, X, V2791)

																																										tmp97836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97839 := Call(__e, tmp97838, A, V2791)

																																										tmp97840 := Call(__e, tmp97837, tmp97839, Nil)

																																										tmp97841 := Call(__e, tmp97836, sym_h, tmp97840)

																																										tmp97842 := Call(__e, tmp97833, tmp97835, tmp97841)

																																										tmp97843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97844 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97845 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97846 := Call(__e, tmp97845, Y, V2791)

																																										tmp97847 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp97851 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97852 := Call(__e, tmp97851, A, V2791)

																																										tmp97853 := Call(__e, tmp97850, tmp97852, Nil)

																																										tmp97854 := Call(__e, tmp97849, symvector, tmp97853)

																																										tmp97855 := Call(__e, tmp97848, tmp97854, Nil)

																																										tmp97856 := Call(__e, tmp97847, sym_h, tmp97855)

																																										tmp97857 := Call(__e, tmp97844, tmp97846, tmp97856)

																																										tmp97858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp97859 := Call(__e, tmp97858, Hyp, V2791)

																																										tmp97860 := Call(__e, tmp97843, tmp97857, tmp97859)

																																										tmp97861 := Call(__e, tmp97832, tmp97842, tmp97860)

																																										__e.TailApply(tmp97831, V2790, tmp97861, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp97862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp97863 := Call(__e, tmp97862, V2574)

																																									__e.TailApply(tmp97828, tmp97863)
																																									return

																																								} else {
																																									tmp97826 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp97827 := Call(__e, tmp97826, V2594)

																																									if True == tmp97827 {
																																										tmp97784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp97785 := Call(__e, tmp97784, V2594, Nil, V2791)

																																										_ = tmp97785

																																										tmp97786 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp97787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp97788 := Call(__e, tmp97787, V2594, V2791)

																																											_ = tmp97788

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp97789 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp97790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp97791 := Call(__e, tmp97790)

																																											_ = tmp97791

																																											tmp97792 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																											tmp97793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97794 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97795 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp97796 := Call(__e, tmp97795, X, V2791)

																																											tmp97797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97799 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp97800 := Call(__e, tmp97799, A, V2791)

																																											tmp97801 := Call(__e, tmp97798, tmp97800, Nil)

																																											tmp97802 := Call(__e, tmp97797, sym_h, tmp97801)

																																											tmp97803 := Call(__e, tmp97794, tmp97796, tmp97802)

																																											tmp97804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp97807 := Call(__e, tmp97806, Y, V2791)

																																											tmp97808 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97809 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97810 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97811 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp97812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp97813 := Call(__e, tmp97812, A, V2791)

																																											tmp97814 := Call(__e, tmp97811, tmp97813, Nil)

																																											tmp97815 := Call(__e, tmp97810, symvector, tmp97814)

																																											tmp97816 := Call(__e, tmp97809, tmp97815, Nil)

																																											tmp97817 := Call(__e, tmp97808, sym_h, tmp97816)

																																											tmp97818 := Call(__e, tmp97805, tmp97807, tmp97817)

																																											tmp97819 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp97820 := Call(__e, tmp97819, Hyp, V2791)

																																											tmp97821 := Call(__e, tmp97804, tmp97818, tmp97820)

																																											tmp97822 := Call(__e, tmp97793, tmp97803, tmp97821)

																																											__e.TailApply(tmp97792, V2790, tmp97822, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp97823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																										tmp97824 := Call(__e, tmp97823, V2574)

																																										tmp97825 := Call(__e, tmp97789, tmp97824)

																																										__e.TailApply(tmp97786, tmp97825)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp97866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp97867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp97868 := Call(__e, tmp97867, V2583)

																																							tmp97869 := Call(__e, tmp97866, tmp97868, V2791)

																																							tmp97870 := Call(__e, tmp97781, tmp97869)

																																							__e.TailApply(tmp97778, tmp97870)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp97964 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp97966 := Call(__e, tmp97965, V2591)

																																				tmp97967 := Call(__e, tmp97964, tmp97966, V2791)

																																				__e.TailApply(tmp97773, tmp97967)
																																				return

																																			}, 1)

																																			tmp97968 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																			tmp97969 := Call(__e, tmp97968, V2591)

																																			__e.TailApply(tmp97772, tmp97969)
																																			return

																																		} else {
																																			tmp97770 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp97771 := Call(__e, tmp97770, V2591)

																																			if True == tmp97771 {
																																				tmp97670 := MakeNative(func(__e *ControlFlow) {
																																					A := __e.Get(1)
																																					_ = A
																																					tmp97671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp97672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp97673 := Call(__e, tmp97672, A, Nil)

																																					tmp97674 := Call(__e, tmp97671, V2591, tmp97673, V2791)

																																					_ = tmp97674

																																					tmp97675 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp97676 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp97677 := Call(__e, tmp97676, V2591, V2791)

																																						_ = tmp97677

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp97678 := MakeNative(func(__e *ControlFlow) {
																																						V2595 := __e.Get(1)
																																						_ = V2595
																																						tmp97761 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp97762 := Call(__e, tmp97761, Nil, V2595)

																																						if True == tmp97762 {
																																							tmp97725 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp97726 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp97727 := Call(__e, tmp97726)

																																								_ = tmp97727

																																								tmp97728 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																								tmp97729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp97732 := Call(__e, tmp97731, X, V2791)

																																								tmp97733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97735 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp97736 := Call(__e, tmp97735, A, V2791)

																																								tmp97737 := Call(__e, tmp97734, tmp97736, Nil)

																																								tmp97738 := Call(__e, tmp97733, sym_h, tmp97737)

																																								tmp97739 := Call(__e, tmp97730, tmp97732, tmp97738)

																																								tmp97740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp97743 := Call(__e, tmp97742, Y, V2791)

																																								tmp97744 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97745 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97746 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97747 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp97748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp97749 := Call(__e, tmp97748, A, V2791)

																																								tmp97750 := Call(__e, tmp97747, tmp97749, Nil)

																																								tmp97751 := Call(__e, tmp97746, symvector, tmp97750)

																																								tmp97752 := Call(__e, tmp97745, tmp97751, Nil)

																																								tmp97753 := Call(__e, tmp97744, sym_h, tmp97752)

																																								tmp97754 := Call(__e, tmp97741, tmp97743, tmp97753)

																																								tmp97755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp97756 := Call(__e, tmp97755, Hyp, V2791)

																																								tmp97757 := Call(__e, tmp97740, tmp97754, tmp97756)

																																								tmp97758 := Call(__e, tmp97729, tmp97739, tmp97757)

																																								__e.TailApply(tmp97728, V2790, tmp97758, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp97759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp97760 := Call(__e, tmp97759, V2574)

																																							__e.TailApply(tmp97725, tmp97760)
																																							return

																																						} else {
																																							tmp97723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																							tmp97724 := Call(__e, tmp97723, V2595)

																																							if True == tmp97724 {
																																								tmp97681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp97682 := Call(__e, tmp97681, V2595, Nil, V2791)

																																								_ = tmp97682

																																								tmp97683 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp97684 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp97685 := Call(__e, tmp97684, V2595, V2791)

																																									_ = tmp97685

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp97686 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp97687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp97688 := Call(__e, tmp97687)

																																									_ = tmp97688

																																									tmp97689 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp97690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97693 := Call(__e, tmp97692, X, V2791)

																																									tmp97694 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97695 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97697 := Call(__e, tmp97696, A, V2791)

																																									tmp97698 := Call(__e, tmp97695, tmp97697, Nil)

																																									tmp97699 := Call(__e, tmp97694, sym_h, tmp97698)

																																									tmp97700 := Call(__e, tmp97691, tmp97693, tmp97699)

																																									tmp97701 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97702 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97704 := Call(__e, tmp97703, Y, V2791)

																																									tmp97705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97707 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp97709 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97710 := Call(__e, tmp97709, A, V2791)

																																									tmp97711 := Call(__e, tmp97708, tmp97710, Nil)

																																									tmp97712 := Call(__e, tmp97707, symvector, tmp97711)

																																									tmp97713 := Call(__e, tmp97706, tmp97712, Nil)

																																									tmp97714 := Call(__e, tmp97705, sym_h, tmp97713)

																																									tmp97715 := Call(__e, tmp97702, tmp97704, tmp97714)

																																									tmp97716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp97717 := Call(__e, tmp97716, Hyp, V2791)

																																									tmp97718 := Call(__e, tmp97701, tmp97715, tmp97717)

																																									tmp97719 := Call(__e, tmp97690, tmp97700, tmp97718)

																																									__e.TailApply(tmp97689, V2790, tmp97719, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp97720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp97721 := Call(__e, tmp97720, V2574)

																																								tmp97722 := Call(__e, tmp97686, tmp97721)

																																								__e.TailApply(tmp97683, tmp97722)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp97763 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp97764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp97765 := Call(__e, tmp97764, V2583)

																																					tmp97766 := Call(__e, tmp97763, tmp97765, V2791)

																																					tmp97767 := Call(__e, tmp97678, tmp97766)

																																					__e.TailApply(tmp97675, tmp97767)
																																					return

																																				}, 1)

																																				tmp97768 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																				tmp97769 := Call(__e, tmp97768, V2791)

																																				__e.TailApply(tmp97670, tmp97769)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp97972 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp97973 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp97974 := Call(__e, tmp97973, V2584)

																																	tmp97975 := Call(__e, tmp97972, tmp97974, V2791)

																																	tmp97976 := Call(__e, tmp97667, tmp97975)

																																	__e.TailApply(tmp97664, tmp97976)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}

																														}, 1)

																														tmp98290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																														tmp98291 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp98292 := Call(__e, tmp98291, V2584)

																														tmp98293 := Call(__e, tmp98290, tmp98292, V2791)

																														__e.TailApply(tmp97659, tmp98293)
																														return

																													} else {
																														tmp97657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																														tmp97658 := Call(__e, tmp97657, V2584)

																														if True == tmp97658 {
																															tmp97555 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp97556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																tmp97557 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp97558 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp97559 := Call(__e, tmp97558, A, Nil)

																																tmp97560 := Call(__e, tmp97557, symvector, tmp97559)

																																tmp97561 := Call(__e, tmp97556, V2584, tmp97560, V2791)

																																_ = tmp97561

																																tmp97562 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp97563 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																	tmp97564 := Call(__e, tmp97563, V2584, V2791)

																																	_ = tmp97564

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp97565 := MakeNative(func(__e *ControlFlow) {
																																	V2596 := __e.Get(1)
																																	_ = V2596
																																	tmp97648 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																	tmp97649 := Call(__e, tmp97648, Nil, V2596)

																																	if True == tmp97649 {
																																		tmp97612 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp97613 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp97614 := Call(__e, tmp97613)

																																			_ = tmp97614

																																			tmp97615 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																			tmp97616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97619 := Call(__e, tmp97618, X, V2791)

																																			tmp97620 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97621 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97622 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97623 := Call(__e, tmp97622, A, V2791)

																																			tmp97624 := Call(__e, tmp97621, tmp97623, Nil)

																																			tmp97625 := Call(__e, tmp97620, sym_h, tmp97624)

																																			tmp97626 := Call(__e, tmp97617, tmp97619, tmp97625)

																																			tmp97627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97629 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97630 := Call(__e, tmp97629, Y, V2791)

																																			tmp97631 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97632 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97633 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97634 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp97635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97636 := Call(__e, tmp97635, A, V2791)

																																			tmp97637 := Call(__e, tmp97634, tmp97636, Nil)

																																			tmp97638 := Call(__e, tmp97633, symvector, tmp97637)

																																			tmp97639 := Call(__e, tmp97632, tmp97638, Nil)

																																			tmp97640 := Call(__e, tmp97631, sym_h, tmp97639)

																																			tmp97641 := Call(__e, tmp97628, tmp97630, tmp97640)

																																			tmp97642 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp97643 := Call(__e, tmp97642, Hyp, V2791)

																																			tmp97644 := Call(__e, tmp97627, tmp97641, tmp97643)

																																			tmp97645 := Call(__e, tmp97616, tmp97626, tmp97644)

																																			__e.TailApply(tmp97615, V2790, tmp97645, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp97646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp97647 := Call(__e, tmp97646, V2574)

																																		__e.TailApply(tmp97612, tmp97647)
																																		return

																																	} else {
																																		tmp97610 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp97611 := Call(__e, tmp97610, V2596)

																																		if True == tmp97611 {
																																			tmp97568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp97569 := Call(__e, tmp97568, V2596, Nil, V2791)

																																			_ = tmp97569

																																			tmp97570 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp97571 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp97572 := Call(__e, tmp97571, V2596, V2791)

																																				_ = tmp97572

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp97573 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp97574 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp97575 := Call(__e, tmp97574)

																																				_ = tmp97575

																																				tmp97576 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp97577 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97578 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97580 := Call(__e, tmp97579, X, V2791)

																																				tmp97581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97583 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97584 := Call(__e, tmp97583, A, V2791)

																																				tmp97585 := Call(__e, tmp97582, tmp97584, Nil)

																																				tmp97586 := Call(__e, tmp97581, sym_h, tmp97585)

																																				tmp97587 := Call(__e, tmp97578, tmp97580, tmp97586)

																																				tmp97588 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97589 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97591 := Call(__e, tmp97590, Y, V2791)

																																				tmp97592 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97594 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97595 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp97596 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97597 := Call(__e, tmp97596, A, V2791)

																																				tmp97598 := Call(__e, tmp97595, tmp97597, Nil)

																																				tmp97599 := Call(__e, tmp97594, symvector, tmp97598)

																																				tmp97600 := Call(__e, tmp97593, tmp97599, Nil)

																																				tmp97601 := Call(__e, tmp97592, sym_h, tmp97600)

																																				tmp97602 := Call(__e, tmp97589, tmp97591, tmp97601)

																																				tmp97603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp97604 := Call(__e, tmp97603, Hyp, V2791)

																																				tmp97605 := Call(__e, tmp97588, tmp97602, tmp97604)

																																				tmp97606 := Call(__e, tmp97577, tmp97587, tmp97605)

																																				__e.TailApply(tmp97576, V2790, tmp97606, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp97607 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp97608 := Call(__e, tmp97607, V2574)

																																			tmp97609 := Call(__e, tmp97573, tmp97608)

																																			__e.TailApply(tmp97570, tmp97609)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp97650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp97651 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp97652 := Call(__e, tmp97651, V2583)

																																tmp97653 := Call(__e, tmp97650, tmp97652, V2791)

																																tmp97654 := Call(__e, tmp97565, tmp97653)

																																__e.TailApply(tmp97562, tmp97654)
																																return

																															}, 1)

																															tmp97655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																															tmp97656 := Call(__e, tmp97655, V2791)

																															__e.TailApply(tmp97555, tmp97656)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp98296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																												tmp98297 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																												tmp98298 := Call(__e, tmp98297, V2583)

																												tmp98299 := Call(__e, tmp98296, tmp98298, V2791)

																												__e.TailApply(tmp97552, tmp98299)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp98302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																										tmp98303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																										tmp98304 := Call(__e, tmp98303, V2581)

																										tmp98305 := Call(__e, tmp98302, tmp98304, V2791)

																										__e.TailApply(tmp97550, tmp98305)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp98308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																								tmp98309 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								tmp98310 := Call(__e, tmp98309, V2581)

																								tmp98311 := Call(__e, tmp98308, tmp98310, V2791)

																								__e.TailApply(tmp97548, tmp98311)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp98314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp98315 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp98316 := Call(__e, tmp98315, V2575)

																						tmp98317 := Call(__e, tmp98314, tmp98316, V2791)

																						__e.TailApply(tmp97546, tmp98317)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp98320 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp98321 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp98322 := Call(__e, tmp98321, V2579)

																				tmp98323 := Call(__e, tmp98320, tmp98322, V2791)

																				__e.TailApply(tmp97544, tmp98323)
																				return

																			}, 1)

																			tmp98324 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp98325 := Call(__e, tmp98324, V2579)

																			__e.TailApply(tmp97543, tmp98325)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp98328 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp98329 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp98330 := Call(__e, tmp98329, V2578)

																	tmp98331 := Call(__e, tmp98328, tmp98330, V2791)

																	__e.TailApply(tmp97541, tmp98331)
																	return

																}, 1)

																tmp98332 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																tmp98333 := Call(__e, tmp98332, V2578)

																__e.TailApply(tmp97540, tmp98333)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp98336 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp98337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp98338 := Call(__e, tmp98337, V2576)

														tmp98339 := Call(__e, tmp98336, tmp98338, V2791)

														__e.TailApply(tmp97538, tmp98339)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp98342 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp98343 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp98344 := Call(__e, tmp98343, V2576)

												tmp98345 := Call(__e, tmp98342, tmp98344, V2791)

												__e.TailApply(tmp97536, tmp98345)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp98348 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp98349 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp98350 := Call(__e, tmp98349, V2575)

										tmp98351 := Call(__e, tmp98348, tmp98350, V2791)

										__e.TailApply(tmp97534, tmp98351)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp98354 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp98355 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp98356 := Call(__e, tmp98355, V2574)

								tmp98357 := Call(__e, tmp98354, tmp98356, V2791)

								__e.TailApply(tmp97532, tmp98357)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp98360 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp98361 := Call(__e, tmp98360, V2789, V2791)

						tmp98362 := Call(__e, tmp97530, tmp98361)

						__e.TailApply(tmp97248, tmp98362)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp98365 := MakeNative(func(__e *ControlFlow) {
					V2549 := __e.Get(1)
					_ = V2549
					tmp99250 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp99251 := Call(__e, tmp99250, V2549)

					if True == tmp99251 {
						tmp98367 := MakeNative(func(__e *ControlFlow) {
							V2550 := __e.Get(1)
							_ = V2550
							tmp99244 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp99245 := Call(__e, tmp99244, V2550)

							if True == tmp99245 {
								tmp98369 := MakeNative(func(__e *ControlFlow) {
									V2551 := __e.Get(1)
									_ = V2551
									tmp99238 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp99239 := Call(__e, tmp99238, V2551)

									if True == tmp99239 {
										tmp98371 := MakeNative(func(__e *ControlFlow) {
											V2552 := __e.Get(1)
											_ = V2552
											tmp99232 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp99233 := Call(__e, tmp99232, sym_8p, V2552)

											if True == tmp99233 {
												tmp98373 := MakeNative(func(__e *ControlFlow) {
													V2553 := __e.Get(1)
													_ = V2553
													tmp99226 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp99227 := Call(__e, tmp99226, V2553)

													if True == tmp99227 {
														tmp98375 := MakeNative(func(__e *ControlFlow) {
															X := __e.Get(1)
															_ = X
															tmp98376 := MakeNative(func(__e *ControlFlow) {
																V2554 := __e.Get(1)
																_ = V2554
																tmp99218 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp99219 := Call(__e, tmp99218, V2554)

																if True == tmp99219 {
																	tmp98378 := MakeNative(func(__e *ControlFlow) {
																		Y := __e.Get(1)
																		_ = Y
																		tmp98379 := MakeNative(func(__e *ControlFlow) {
																			V2555 := __e.Get(1)
																			_ = V2555
																			tmp99210 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp99211 := Call(__e, tmp99210, Nil, V2555)

																			if True == tmp99211 {
																				tmp98381 := MakeNative(func(__e *ControlFlow) {
																					V2556 := __e.Get(1)
																					_ = V2556
																					tmp99204 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																					tmp99205 := Call(__e, tmp99204, V2556)

																					if True == tmp99205 {
																						tmp98383 := MakeNative(func(__e *ControlFlow) {
																							V2557 := __e.Get(1)
																							_ = V2557
																							tmp99198 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																							tmp99199 := Call(__e, tmp99198, sym_h, V2557)

																							if True == tmp99199 {
																								tmp98385 := MakeNative(func(__e *ControlFlow) {
																									V2558 := __e.Get(1)
																									_ = V2558
																									tmp99192 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																									tmp99193 := Call(__e, tmp99192, V2558)

																									if True == tmp99193 {
																										tmp98387 := MakeNative(func(__e *ControlFlow) {
																											V2559 := __e.Get(1)
																											_ = V2559
																											tmp99186 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																											tmp99187 := Call(__e, tmp99186, V2559)

																											if True == tmp99187 {
																												tmp98491 := MakeNative(func(__e *ControlFlow) {
																													A := __e.Get(1)
																													_ = A
																													tmp98492 := MakeNative(func(__e *ControlFlow) {
																														V2560 := __e.Get(1)
																														_ = V2560
																														tmp99178 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																														tmp99179 := Call(__e, tmp99178, V2560)

																														if True == tmp99179 {
																															tmp98591 := MakeNative(func(__e *ControlFlow) {
																																V2561 := __e.Get(1)
																																_ = V2561
																																tmp99172 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp99173 := Call(__e, tmp99172, sym_d, V2561)

																																if True == tmp99173 {
																																	tmp98887 := MakeNative(func(__e *ControlFlow) {
																																		V2562 := __e.Get(1)
																																		_ = V2562
																																		tmp99166 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																		tmp99167 := Call(__e, tmp99166, V2562)

																																		if True == tmp99167 {
																																			tmp98984 := MakeNative(func(__e *ControlFlow) {
																																				B := __e.Get(1)
																																				_ = B
																																				tmp98985 := MakeNative(func(__e *ControlFlow) {
																																					V2563 := __e.Get(1)
																																					_ = V2563
																																					tmp99158 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																					tmp99159 := Call(__e, tmp99158, Nil, V2563)

																																					if True == tmp99159 {
																																						tmp99077 := MakeNative(func(__e *ControlFlow) {
																																							V2564 := __e.Get(1)
																																							_ = V2564
																																							tmp99152 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp99153 := Call(__e, tmp99152, Nil, V2564)

																																							if True == tmp99153 {
																																								tmp99120 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp99121 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp99122 := Call(__e, tmp99121)

																																									_ = tmp99122

																																									tmp99123 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp99124 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99125 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp99127 := Call(__e, tmp99126, X, V2791)

																																									tmp99128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp99131 := Call(__e, tmp99130, A, V2791)

																																									tmp99132 := Call(__e, tmp99129, tmp99131, Nil)

																																									tmp99133 := Call(__e, tmp99128, sym_h, tmp99132)

																																									tmp99134 := Call(__e, tmp99125, tmp99127, tmp99133)

																																									tmp99135 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99137 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp99138 := Call(__e, tmp99137, Y, V2791)

																																									tmp99139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp99141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp99142 := Call(__e, tmp99141, B, V2791)

																																									tmp99143 := Call(__e, tmp99140, tmp99142, Nil)

																																									tmp99144 := Call(__e, tmp99139, sym_h, tmp99143)

																																									tmp99145 := Call(__e, tmp99136, tmp99138, tmp99144)

																																									tmp99146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp99147 := Call(__e, tmp99146, Hyp, V2791)

																																									tmp99148 := Call(__e, tmp99135, tmp99145, tmp99147)

																																									tmp99149 := Call(__e, tmp99124, tmp99134, tmp99148)

																																									__e.TailApply(tmp99123, V2790, tmp99149, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp99150 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp99151 := Call(__e, tmp99150, V2549)

																																								__e.TailApply(tmp99120, tmp99151)
																																								return

																																							} else {
																																								tmp99118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp99119 := Call(__e, tmp99118, V2564)

																																								if True == tmp99119 {
																																									tmp99080 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp99081 := Call(__e, tmp99080, V2564, Nil, V2791)

																																									_ = tmp99081

																																									tmp99082 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp99083 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp99084 := Call(__e, tmp99083, V2564, V2791)

																																										_ = tmp99084

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp99085 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp99086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp99087 := Call(__e, tmp99086)

																																										_ = tmp99087

																																										tmp99088 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp99089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99092 := Call(__e, tmp99091, X, V2791)

																																										tmp99093 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99094 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99096 := Call(__e, tmp99095, A, V2791)

																																										tmp99097 := Call(__e, tmp99094, tmp99096, Nil)

																																										tmp99098 := Call(__e, tmp99093, sym_h, tmp99097)

																																										tmp99099 := Call(__e, tmp99090, tmp99092, tmp99098)

																																										tmp99100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99101 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99103 := Call(__e, tmp99102, Y, V2791)

																																										tmp99104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99105 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99107 := Call(__e, tmp99106, B, V2791)

																																										tmp99108 := Call(__e, tmp99105, tmp99107, Nil)

																																										tmp99109 := Call(__e, tmp99104, sym_h, tmp99108)

																																										tmp99110 := Call(__e, tmp99101, tmp99103, tmp99109)

																																										tmp99111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99112 := Call(__e, tmp99111, Hyp, V2791)

																																										tmp99113 := Call(__e, tmp99100, tmp99110, tmp99112)

																																										tmp99114 := Call(__e, tmp99089, tmp99099, tmp99113)

																																										__e.TailApply(tmp99088, V2790, tmp99114, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp99115 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp99116 := Call(__e, tmp99115, V2549)

																																									tmp99117 := Call(__e, tmp99085, tmp99116)

																																									__e.TailApply(tmp99082, tmp99117)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp99154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99155 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp99156 := Call(__e, tmp99155, V2558)

																																						tmp99157 := Call(__e, tmp99154, tmp99156, V2791)

																																						__e.TailApply(tmp99077, tmp99157)
																																						return

																																					} else {
																																						tmp99075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																						tmp99076 := Call(__e, tmp99075, V2563)

																																						if True == tmp99076 {
																																							tmp98988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																							tmp98989 := Call(__e, tmp98988, V2563, Nil, V2791)

																																							_ = tmp98989

																																							tmp98990 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp98991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																								tmp98992 := Call(__e, tmp98991, V2563, V2791)

																																								_ = tmp98992

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp98993 := MakeNative(func(__e *ControlFlow) {
																																								V2565 := __e.Get(1)
																																								_ = V2565
																																								tmp99068 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp99069 := Call(__e, tmp99068, Nil, V2565)

																																								if True == tmp99069 {
																																									tmp99036 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp99037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp99038 := Call(__e, tmp99037)

																																										_ = tmp99038

																																										tmp99039 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp99040 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99041 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99042 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99043 := Call(__e, tmp99042, X, V2791)

																																										tmp99044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99046 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99047 := Call(__e, tmp99046, A, V2791)

																																										tmp99048 := Call(__e, tmp99045, tmp99047, Nil)

																																										tmp99049 := Call(__e, tmp99044, sym_h, tmp99048)

																																										tmp99050 := Call(__e, tmp99041, tmp99043, tmp99049)

																																										tmp99051 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99052 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99053 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99054 := Call(__e, tmp99053, Y, V2791)

																																										tmp99055 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99056 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp99057 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99058 := Call(__e, tmp99057, B, V2791)

																																										tmp99059 := Call(__e, tmp99056, tmp99058, Nil)

																																										tmp99060 := Call(__e, tmp99055, sym_h, tmp99059)

																																										tmp99061 := Call(__e, tmp99052, tmp99054, tmp99060)

																																										tmp99062 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp99063 := Call(__e, tmp99062, Hyp, V2791)

																																										tmp99064 := Call(__e, tmp99051, tmp99061, tmp99063)

																																										tmp99065 := Call(__e, tmp99040, tmp99050, tmp99064)

																																										__e.TailApply(tmp99039, V2790, tmp99065, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp99066 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp99067 := Call(__e, tmp99066, V2549)

																																									__e.TailApply(tmp99036, tmp99067)
																																									return

																																								} else {
																																									tmp99034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp99035 := Call(__e, tmp99034, V2565)

																																									if True == tmp99035 {
																																										tmp98996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp98997 := Call(__e, tmp98996, V2565, Nil, V2791)

																																										_ = tmp98997

																																										tmp98998 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp98999 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp99000 := Call(__e, tmp98999, V2565, V2791)

																																											_ = tmp99000

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp99001 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp99002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp99003 := Call(__e, tmp99002)

																																											_ = tmp99003

																																											tmp99004 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																											tmp99005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99007 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp99008 := Call(__e, tmp99007, X, V2791)

																																											tmp99009 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99010 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99011 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp99012 := Call(__e, tmp99011, A, V2791)

																																											tmp99013 := Call(__e, tmp99010, tmp99012, Nil)

																																											tmp99014 := Call(__e, tmp99009, sym_h, tmp99013)

																																											tmp99015 := Call(__e, tmp99006, tmp99008, tmp99014)

																																											tmp99016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99018 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp99019 := Call(__e, tmp99018, Y, V2791)

																																											tmp99020 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99021 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp99022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp99023 := Call(__e, tmp99022, B, V2791)

																																											tmp99024 := Call(__e, tmp99021, tmp99023, Nil)

																																											tmp99025 := Call(__e, tmp99020, sym_h, tmp99024)

																																											tmp99026 := Call(__e, tmp99017, tmp99019, tmp99025)

																																											tmp99027 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp99028 := Call(__e, tmp99027, Hyp, V2791)

																																											tmp99029 := Call(__e, tmp99016, tmp99026, tmp99028)

																																											tmp99030 := Call(__e, tmp99005, tmp99015, tmp99029)

																																											__e.TailApply(tmp99004, V2790, tmp99030, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp99031 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																										tmp99032 := Call(__e, tmp99031, V2549)

																																										tmp99033 := Call(__e, tmp99001, tmp99032)

																																										__e.TailApply(tmp98998, tmp99033)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp99070 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp99072 := Call(__e, tmp99071, V2558)

																																							tmp99073 := Call(__e, tmp99070, tmp99072, V2791)

																																							tmp99074 := Call(__e, tmp98993, tmp99073)

																																							__e.TailApply(tmp98990, tmp99074)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp99160 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99161 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp99162 := Call(__e, tmp99161, V2562)

																																				tmp99163 := Call(__e, tmp99160, tmp99162, V2791)

																																				__e.TailApply(tmp98985, tmp99163)
																																				return

																																			}, 1)

																																			tmp99164 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																			tmp99165 := Call(__e, tmp99164, V2562)

																																			__e.TailApply(tmp98984, tmp99165)
																																			return

																																		} else {
																																			tmp98982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp98983 := Call(__e, tmp98982, V2562)

																																			if True == tmp98983 {
																																				tmp98890 := MakeNative(func(__e *ControlFlow) {
																																					B := __e.Get(1)
																																					_ = B
																																					tmp98891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp98892 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98893 := Call(__e, tmp98892, B, Nil)

																																					tmp98894 := Call(__e, tmp98891, V2562, tmp98893, V2791)

																																					_ = tmp98894

																																					tmp98895 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp98896 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp98897 := Call(__e, tmp98896, V2562, V2791)

																																						_ = tmp98897

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp98898 := MakeNative(func(__e *ControlFlow) {
																																						V2566 := __e.Get(1)
																																						_ = V2566
																																						tmp98973 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp98974 := Call(__e, tmp98973, Nil, V2566)

																																						if True == tmp98974 {
																																							tmp98941 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp98942 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																								tmp98943 := Call(__e, tmp98942)

																																								_ = tmp98943

																																								tmp98944 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																								tmp98945 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98946 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98948 := Call(__e, tmp98947, X, V2791)

																																								tmp98949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98951 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98952 := Call(__e, tmp98951, A, V2791)

																																								tmp98953 := Call(__e, tmp98950, tmp98952, Nil)

																																								tmp98954 := Call(__e, tmp98949, sym_h, tmp98953)

																																								tmp98955 := Call(__e, tmp98946, tmp98948, tmp98954)

																																								tmp98956 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98957 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98958 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98959 := Call(__e, tmp98958, Y, V2791)

																																								tmp98960 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98961 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																								tmp98962 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98963 := Call(__e, tmp98962, B, V2791)

																																								tmp98964 := Call(__e, tmp98961, tmp98963, Nil)

																																								tmp98965 := Call(__e, tmp98960, sym_h, tmp98964)

																																								tmp98966 := Call(__e, tmp98957, tmp98959, tmp98965)

																																								tmp98967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98968 := Call(__e, tmp98967, Hyp, V2791)

																																								tmp98969 := Call(__e, tmp98956, tmp98966, tmp98968)

																																								tmp98970 := Call(__e, tmp98945, tmp98955, tmp98969)

																																								__e.TailApply(tmp98944, V2790, tmp98970, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp98971 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp98972 := Call(__e, tmp98971, V2549)

																																							__e.TailApply(tmp98941, tmp98972)
																																							return

																																						} else {
																																							tmp98939 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																							tmp98940 := Call(__e, tmp98939, V2566)

																																							if True == tmp98940 {
																																								tmp98901 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp98902 := Call(__e, tmp98901, V2566, Nil, V2791)

																																								_ = tmp98902

																																								tmp98903 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp98904 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp98905 := Call(__e, tmp98904, V2566, V2791)

																																									_ = tmp98905

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp98906 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp98907 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp98908 := Call(__e, tmp98907)

																																									_ = tmp98908

																																									tmp98909 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp98910 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98911 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98912 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98913 := Call(__e, tmp98912, X, V2791)

																																									tmp98914 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98915 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98916 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98917 := Call(__e, tmp98916, A, V2791)

																																									tmp98918 := Call(__e, tmp98915, tmp98917, Nil)

																																									tmp98919 := Call(__e, tmp98914, sym_h, tmp98918)

																																									tmp98920 := Call(__e, tmp98911, tmp98913, tmp98919)

																																									tmp98921 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98923 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98924 := Call(__e, tmp98923, Y, V2791)

																																									tmp98925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98927 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98928 := Call(__e, tmp98927, B, V2791)

																																									tmp98929 := Call(__e, tmp98926, tmp98928, Nil)

																																									tmp98930 := Call(__e, tmp98925, sym_h, tmp98929)

																																									tmp98931 := Call(__e, tmp98922, tmp98924, tmp98930)

																																									tmp98932 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98933 := Call(__e, tmp98932, Hyp, V2791)

																																									tmp98934 := Call(__e, tmp98921, tmp98931, tmp98933)

																																									tmp98935 := Call(__e, tmp98910, tmp98920, tmp98934)

																																									__e.TailApply(tmp98909, V2790, tmp98935, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp98936 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp98937 := Call(__e, tmp98936, V2549)

																																								tmp98938 := Call(__e, tmp98906, tmp98937)

																																								__e.TailApply(tmp98903, tmp98938)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp98975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98976 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp98977 := Call(__e, tmp98976, V2558)

																																					tmp98978 := Call(__e, tmp98975, tmp98977, V2791)

																																					tmp98979 := Call(__e, tmp98898, tmp98978)

																																					__e.TailApply(tmp98895, tmp98979)
																																					return

																																				}, 1)

																																				tmp98980 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																				tmp98981 := Call(__e, tmp98980, V2791)

																																				__e.TailApply(tmp98890, tmp98981)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp99168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp99169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp99170 := Call(__e, tmp99169, V2560)

																																	tmp99171 := Call(__e, tmp99168, tmp99170, V2791)

																																	__e.TailApply(tmp98887, tmp99171)
																																	return

																																} else {
																																	tmp98885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																	tmp98886 := Call(__e, tmp98885, V2561)

																																	if True == tmp98886 {
																																		tmp98594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																		tmp98595 := Call(__e, tmp98594, V2561, sym_d, V2791)

																																		_ = tmp98595

																																		tmp98596 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp98597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																			tmp98598 := Call(__e, tmp98597, V2561, V2791)

																																			_ = tmp98598

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp98599 := MakeNative(func(__e *ControlFlow) {
																																			V2567 := __e.Get(1)
																																			_ = V2567
																																			tmp98878 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																																			tmp98879 := Call(__e, tmp98878, V2567)

																																			if True == tmp98879 {
																																				tmp98696 := MakeNative(func(__e *ControlFlow) {
																																					B := __e.Get(1)
																																					_ = B
																																					tmp98697 := MakeNative(func(__e *ControlFlow) {
																																						V2568 := __e.Get(1)
																																						_ = V2568
																																						tmp98870 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																						tmp98871 := Call(__e, tmp98870, Nil, V2568)

																																						if True == tmp98871 {
																																							tmp98789 := MakeNative(func(__e *ControlFlow) {
																																								V2569 := __e.Get(1)
																																								_ = V2569
																																								tmp98864 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																								tmp98865 := Call(__e, tmp98864, Nil, V2569)

																																								if True == tmp98865 {
																																									tmp98832 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp98833 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp98834 := Call(__e, tmp98833)

																																										_ = tmp98834

																																										tmp98835 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp98836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98839 := Call(__e, tmp98838, X, V2791)

																																										tmp98840 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98841 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98843 := Call(__e, tmp98842, A, V2791)

																																										tmp98844 := Call(__e, tmp98841, tmp98843, Nil)

																																										tmp98845 := Call(__e, tmp98840, sym_h, tmp98844)

																																										tmp98846 := Call(__e, tmp98837, tmp98839, tmp98845)

																																										tmp98847 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98849 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98850 := Call(__e, tmp98849, Y, V2791)

																																										tmp98851 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98852 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98853 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98854 := Call(__e, tmp98853, B, V2791)

																																										tmp98855 := Call(__e, tmp98852, tmp98854, Nil)

																																										tmp98856 := Call(__e, tmp98851, sym_h, tmp98855)

																																										tmp98857 := Call(__e, tmp98848, tmp98850, tmp98856)

																																										tmp98858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98859 := Call(__e, tmp98858, Hyp, V2791)

																																										tmp98860 := Call(__e, tmp98847, tmp98857, tmp98859)

																																										tmp98861 := Call(__e, tmp98836, tmp98846, tmp98860)

																																										__e.TailApply(tmp98835, V2790, tmp98861, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp98862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp98863 := Call(__e, tmp98862, V2549)

																																									__e.TailApply(tmp98832, tmp98863)
																																									return

																																								} else {
																																									tmp98830 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																									tmp98831 := Call(__e, tmp98830, V2569)

																																									if True == tmp98831 {
																																										tmp98792 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																										tmp98793 := Call(__e, tmp98792, V2569, Nil, V2791)

																																										_ = tmp98793

																																										tmp98794 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp98795 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																											tmp98796 := Call(__e, tmp98795, V2569, V2791)

																																											_ = tmp98796

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp98797 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp98798 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp98799 := Call(__e, tmp98798)

																																											_ = tmp98799

																																											tmp98800 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																											tmp98801 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98802 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98804 := Call(__e, tmp98803, X, V2791)

																																											tmp98805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98807 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98808 := Call(__e, tmp98807, A, V2791)

																																											tmp98809 := Call(__e, tmp98806, tmp98808, Nil)

																																											tmp98810 := Call(__e, tmp98805, sym_h, tmp98809)

																																											tmp98811 := Call(__e, tmp98802, tmp98804, tmp98810)

																																											tmp98812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98814 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98815 := Call(__e, tmp98814, Y, V2791)

																																											tmp98816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98817 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98819 := Call(__e, tmp98818, B, V2791)

																																											tmp98820 := Call(__e, tmp98817, tmp98819, Nil)

																																											tmp98821 := Call(__e, tmp98816, sym_h, tmp98820)

																																											tmp98822 := Call(__e, tmp98813, tmp98815, tmp98821)

																																											tmp98823 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98824 := Call(__e, tmp98823, Hyp, V2791)

																																											tmp98825 := Call(__e, tmp98812, tmp98822, tmp98824)

																																											tmp98826 := Call(__e, tmp98801, tmp98811, tmp98825)

																																											__e.TailApply(tmp98800, V2790, tmp98826, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp98827 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																										tmp98828 := Call(__e, tmp98827, V2549)

																																										tmp98829 := Call(__e, tmp98797, tmp98828)

																																										__e.TailApply(tmp98794, tmp98829)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp98866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp98867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																							tmp98868 := Call(__e, tmp98867, V2558)

																																							tmp98869 := Call(__e, tmp98866, tmp98868, V2791)

																																							__e.TailApply(tmp98789, tmp98869)
																																							return

																																						} else {
																																							tmp98787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																							tmp98788 := Call(__e, tmp98787, V2568)

																																							if True == tmp98788 {
																																								tmp98700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																								tmp98701 := Call(__e, tmp98700, V2568, Nil, V2791)

																																								_ = tmp98701

																																								tmp98702 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp98703 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																									tmp98704 := Call(__e, tmp98703, V2568, V2791)

																																									_ = tmp98704

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp98705 := MakeNative(func(__e *ControlFlow) {
																																									V2570 := __e.Get(1)
																																									_ = V2570
																																									tmp98780 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																									tmp98781 := Call(__e, tmp98780, Nil, V2570)

																																									if True == tmp98781 {
																																										tmp98748 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp98749 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																											tmp98750 := Call(__e, tmp98749)

																																											_ = tmp98750

																																											tmp98751 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																											tmp98752 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98753 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98754 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98755 := Call(__e, tmp98754, X, V2791)

																																											tmp98756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98758 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98759 := Call(__e, tmp98758, A, V2791)

																																											tmp98760 := Call(__e, tmp98757, tmp98759, Nil)

																																											tmp98761 := Call(__e, tmp98756, sym_h, tmp98760)

																																											tmp98762 := Call(__e, tmp98753, tmp98755, tmp98761)

																																											tmp98763 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98764 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98765 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98766 := Call(__e, tmp98765, Y, V2791)

																																											tmp98767 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98768 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																											tmp98769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98770 := Call(__e, tmp98769, B, V2791)

																																											tmp98771 := Call(__e, tmp98768, tmp98770, Nil)

																																											tmp98772 := Call(__e, tmp98767, sym_h, tmp98771)

																																											tmp98773 := Call(__e, tmp98764, tmp98766, tmp98772)

																																											tmp98774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																											tmp98775 := Call(__e, tmp98774, Hyp, V2791)

																																											tmp98776 := Call(__e, tmp98763, tmp98773, tmp98775)

																																											tmp98777 := Call(__e, tmp98752, tmp98762, tmp98776)

																																											__e.TailApply(tmp98751, V2790, tmp98777, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp98778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																										tmp98779 := Call(__e, tmp98778, V2549)

																																										__e.TailApply(tmp98748, tmp98779)
																																										return

																																									} else {
																																										tmp98746 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																										tmp98747 := Call(__e, tmp98746, V2570)

																																										if True == tmp98747 {
																																											tmp98708 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																											tmp98709 := Call(__e, tmp98708, V2570, Nil, V2791)

																																											_ = tmp98709

																																											tmp98710 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp98711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																												tmp98712 := Call(__e, tmp98711, V2570, V2791)

																																												_ = tmp98712

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp98713 := MakeNative(func(__e *ControlFlow) {
																																												Hyp := __e.Get(1)
																																												_ = Hyp
																																												tmp98714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																												tmp98715 := Call(__e, tmp98714)

																																												_ = tmp98715

																																												tmp98716 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																												tmp98717 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98718 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp98720 := Call(__e, tmp98719, X, V2791)

																																												tmp98721 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98722 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp98724 := Call(__e, tmp98723, A, V2791)

																																												tmp98725 := Call(__e, tmp98722, tmp98724, Nil)

																																												tmp98726 := Call(__e, tmp98721, sym_h, tmp98725)

																																												tmp98727 := Call(__e, tmp98718, tmp98720, tmp98726)

																																												tmp98728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98730 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp98731 := Call(__e, tmp98730, Y, V2791)

																																												tmp98732 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																												tmp98734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp98735 := Call(__e, tmp98734, B, V2791)

																																												tmp98736 := Call(__e, tmp98733, tmp98735, Nil)

																																												tmp98737 := Call(__e, tmp98732, sym_h, tmp98736)

																																												tmp98738 := Call(__e, tmp98729, tmp98731, tmp98737)

																																												tmp98739 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																												tmp98740 := Call(__e, tmp98739, Hyp, V2791)

																																												tmp98741 := Call(__e, tmp98728, tmp98738, tmp98740)

																																												tmp98742 := Call(__e, tmp98717, tmp98727, tmp98741)

																																												__e.TailApply(tmp98716, V2790, tmp98742, V2791, V2792)
																																												return

																																											}, 1)

																																											tmp98743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																											tmp98744 := Call(__e, tmp98743, V2549)

																																											tmp98745 := Call(__e, tmp98713, tmp98744)

																																											__e.TailApply(tmp98710, tmp98745)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp98782 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																								tmp98783 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp98784 := Call(__e, tmp98783, V2558)

																																								tmp98785 := Call(__e, tmp98782, tmp98784, V2791)

																																								tmp98786 := Call(__e, tmp98705, tmp98785)

																																								__e.TailApply(tmp98702, tmp98786)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp98872 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98873 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp98874 := Call(__e, tmp98873, V2567)

																																					tmp98875 := Call(__e, tmp98872, tmp98874, V2791)

																																					__e.TailApply(tmp98697, tmp98875)
																																					return

																																				}, 1)

																																				tmp98876 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																																				tmp98877 := Call(__e, tmp98876, V2567)

																																				__e.TailApply(tmp98696, tmp98877)
																																				return

																																			} else {
																																				tmp98694 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																				tmp98695 := Call(__e, tmp98694, V2567)

																																				if True == tmp98695 {
																																					tmp98602 := MakeNative(func(__e *ControlFlow) {
																																						B := __e.Get(1)
																																						_ = B
																																						tmp98603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																						tmp98604 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp98605 := Call(__e, tmp98604, B, Nil)

																																						tmp98606 := Call(__e, tmp98603, V2567, tmp98605, V2791)

																																						_ = tmp98606

																																						tmp98607 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp98608 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																							tmp98609 := Call(__e, tmp98608, V2567, V2791)

																																							_ = tmp98609

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp98610 := MakeNative(func(__e *ControlFlow) {
																																							V2571 := __e.Get(1)
																																							_ = V2571
																																							tmp98685 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																							tmp98686 := Call(__e, tmp98685, Nil, V2571)

																																							if True == tmp98686 {
																																								tmp98653 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp98654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																									tmp98655 := Call(__e, tmp98654)

																																									_ = tmp98655

																																									tmp98656 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																									tmp98657 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98658 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98659 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98660 := Call(__e, tmp98659, X, V2791)

																																									tmp98661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98664 := Call(__e, tmp98663, A, V2791)

																																									tmp98665 := Call(__e, tmp98662, tmp98664, Nil)

																																									tmp98666 := Call(__e, tmp98661, sym_h, tmp98665)

																																									tmp98667 := Call(__e, tmp98658, tmp98660, tmp98666)

																																									tmp98668 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98669 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98670 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98671 := Call(__e, tmp98670, Y, V2791)

																																									tmp98672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98673 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																									tmp98674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98675 := Call(__e, tmp98674, B, V2791)

																																									tmp98676 := Call(__e, tmp98673, tmp98675, Nil)

																																									tmp98677 := Call(__e, tmp98672, sym_h, tmp98676)

																																									tmp98678 := Call(__e, tmp98669, tmp98671, tmp98677)

																																									tmp98679 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																									tmp98680 := Call(__e, tmp98679, Hyp, V2791)

																																									tmp98681 := Call(__e, tmp98668, tmp98678, tmp98680)

																																									tmp98682 := Call(__e, tmp98657, tmp98667, tmp98681)

																																									__e.TailApply(tmp98656, V2790, tmp98682, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp98683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																								tmp98684 := Call(__e, tmp98683, V2549)

																																								__e.TailApply(tmp98653, tmp98684)
																																								return

																																							} else {
																																								tmp98651 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																								tmp98652 := Call(__e, tmp98651, V2571)

																																								if True == tmp98652 {
																																									tmp98613 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																									tmp98614 := Call(__e, tmp98613, V2571, Nil, V2791)

																																									_ = tmp98614

																																									tmp98615 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp98616 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																										tmp98617 := Call(__e, tmp98616, V2571, V2791)

																																										_ = tmp98617

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp98618 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp98619 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																										tmp98620 := Call(__e, tmp98619)

																																										_ = tmp98620

																																										tmp98621 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																										tmp98622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98624 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98625 := Call(__e, tmp98624, X, V2791)

																																										tmp98626 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98628 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98629 := Call(__e, tmp98628, A, V2791)

																																										tmp98630 := Call(__e, tmp98627, tmp98629, Nil)

																																										tmp98631 := Call(__e, tmp98626, sym_h, tmp98630)

																																										tmp98632 := Call(__e, tmp98623, tmp98625, tmp98631)

																																										tmp98633 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98634 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98636 := Call(__e, tmp98635, Y, V2791)

																																										tmp98637 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																										tmp98639 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98640 := Call(__e, tmp98639, B, V2791)

																																										tmp98641 := Call(__e, tmp98638, tmp98640, Nil)

																																										tmp98642 := Call(__e, tmp98637, sym_h, tmp98641)

																																										tmp98643 := Call(__e, tmp98634, tmp98636, tmp98642)

																																										tmp98644 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																										tmp98645 := Call(__e, tmp98644, Hyp, V2791)

																																										tmp98646 := Call(__e, tmp98633, tmp98643, tmp98645)

																																										tmp98647 := Call(__e, tmp98622, tmp98632, tmp98646)

																																										__e.TailApply(tmp98621, V2790, tmp98647, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp98648 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																									tmp98649 := Call(__e, tmp98648, V2549)

																																									tmp98650 := Call(__e, tmp98618, tmp98649)

																																									__e.TailApply(tmp98615, tmp98650)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp98687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp98688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp98689 := Call(__e, tmp98688, V2558)

																																						tmp98690 := Call(__e, tmp98687, tmp98689, V2791)

																																						tmp98691 := Call(__e, tmp98610, tmp98690)

																																						__e.TailApply(tmp98607, tmp98691)
																																						return

																																					}, 1)

																																					tmp98692 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																					tmp98693 := Call(__e, tmp98692, V2791)

																																					__e.TailApply(tmp98602, tmp98693)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp98880 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98881 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp98882 := Call(__e, tmp98881, V2560)

																																		tmp98883 := Call(__e, tmp98880, tmp98882, V2791)

																																		tmp98884 := Call(__e, tmp98599, tmp98883)

																																		__e.TailApply(tmp98596, tmp98884)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp99174 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99175 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																															tmp99176 := Call(__e, tmp99175, V2560)

																															tmp99177 := Call(__e, tmp99174, tmp99176, V2791)

																															__e.TailApply(tmp98591, tmp99177)
																															return

																														} else {
																															tmp98589 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																															tmp98590 := Call(__e, tmp98589, V2560)

																															if True == tmp98590 {
																																tmp98495 := MakeNative(func(__e *ControlFlow) {
																																	B := __e.Get(1)
																																	_ = B
																																	tmp98496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																	tmp98497 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																	tmp98498 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																	tmp98499 := Call(__e, tmp98498, B, Nil)

																																	tmp98500 := Call(__e, tmp98497, sym_d, tmp98499)

																																	tmp98501 := Call(__e, tmp98496, V2560, tmp98500, V2791)

																																	_ = tmp98501

																																	tmp98502 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp98503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																		tmp98504 := Call(__e, tmp98503, V2560, V2791)

																																		_ = tmp98504

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp98505 := MakeNative(func(__e *ControlFlow) {
																																		V2572 := __e.Get(1)
																																		_ = V2572
																																		tmp98580 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp98581 := Call(__e, tmp98580, Nil, V2572)

																																		if True == tmp98581 {
																																			tmp98548 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp98549 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp98550 := Call(__e, tmp98549)

																																				_ = tmp98550

																																				tmp98551 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp98552 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98553 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98554 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98555 := Call(__e, tmp98554, X, V2791)

																																				tmp98556 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98557 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98559 := Call(__e, tmp98558, A, V2791)

																																				tmp98560 := Call(__e, tmp98557, tmp98559, Nil)

																																				tmp98561 := Call(__e, tmp98556, sym_h, tmp98560)

																																				tmp98562 := Call(__e, tmp98553, tmp98555, tmp98561)

																																				tmp98563 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98564 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98566 := Call(__e, tmp98565, Y, V2791)

																																				tmp98567 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp98569 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98570 := Call(__e, tmp98569, B, V2791)

																																				tmp98571 := Call(__e, tmp98568, tmp98570, Nil)

																																				tmp98572 := Call(__e, tmp98567, sym_h, tmp98571)

																																				tmp98573 := Call(__e, tmp98564, tmp98566, tmp98572)

																																				tmp98574 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp98575 := Call(__e, tmp98574, Hyp, V2791)

																																				tmp98576 := Call(__e, tmp98563, tmp98573, tmp98575)

																																				tmp98577 := Call(__e, tmp98552, tmp98562, tmp98576)

																																				__e.TailApply(tmp98551, V2790, tmp98577, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp98578 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp98579 := Call(__e, tmp98578, V2549)

																																			__e.TailApply(tmp98548, tmp98579)
																																			return

																																		} else {
																																			tmp98546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp98547 := Call(__e, tmp98546, V2572)

																																			if True == tmp98547 {
																																				tmp98508 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp98509 := Call(__e, tmp98508, V2572, Nil, V2791)

																																				_ = tmp98509

																																				tmp98510 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp98511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp98512 := Call(__e, tmp98511, V2572, V2791)

																																					_ = tmp98512

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp98513 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp98514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp98515 := Call(__e, tmp98514)

																																					_ = tmp98515

																																					tmp98516 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp98517 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98518 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98520 := Call(__e, tmp98519, X, V2791)

																																					tmp98521 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98522 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98524 := Call(__e, tmp98523, A, V2791)

																																					tmp98525 := Call(__e, tmp98522, tmp98524, Nil)

																																					tmp98526 := Call(__e, tmp98521, sym_h, tmp98525)

																																					tmp98527 := Call(__e, tmp98518, tmp98520, tmp98526)

																																					tmp98528 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98529 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98531 := Call(__e, tmp98530, Y, V2791)

																																					tmp98532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98533 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp98534 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98535 := Call(__e, tmp98534, B, V2791)

																																					tmp98536 := Call(__e, tmp98533, tmp98535, Nil)

																																					tmp98537 := Call(__e, tmp98532, sym_h, tmp98536)

																																					tmp98538 := Call(__e, tmp98529, tmp98531, tmp98537)

																																					tmp98539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp98540 := Call(__e, tmp98539, Hyp, V2791)

																																					tmp98541 := Call(__e, tmp98528, tmp98538, tmp98540)

																																					tmp98542 := Call(__e, tmp98517, tmp98527, tmp98541)

																																					__e.TailApply(tmp98516, V2790, tmp98542, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp98543 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp98544 := Call(__e, tmp98543, V2549)

																																				tmp98545 := Call(__e, tmp98513, tmp98544)

																																				__e.TailApply(tmp98510, tmp98545)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp98582 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp98583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp98584 := Call(__e, tmp98583, V2558)

																																	tmp98585 := Call(__e, tmp98582, tmp98584, V2791)

																																	tmp98586 := Call(__e, tmp98505, tmp98585)

																																	__e.TailApply(tmp98502, tmp98586)
																																	return

																																}, 1)

																																tmp98587 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																tmp98588 := Call(__e, tmp98587, V2791)

																																__e.TailApply(tmp98495, tmp98588)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp99180 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp99181 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																													tmp99182 := Call(__e, tmp99181, V2559)

																													tmp99183 := Call(__e, tmp99180, tmp99182, V2791)

																													__e.TailApply(tmp98492, tmp99183)
																													return

																												}, 1)

																												tmp99184 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																												tmp99185 := Call(__e, tmp99184, V2559)

																												__e.TailApply(tmp98491, tmp99185)
																												return

																											} else {
																												tmp98489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																												tmp98490 := Call(__e, tmp98489, V2559)

																												if True == tmp98490 {
																													tmp98390 := MakeNative(func(__e *ControlFlow) {
																														A := __e.Get(1)
																														_ = A
																														tmp98391 := MakeNative(func(__e *ControlFlow) {
																															B := __e.Get(1)
																															_ = B
																															tmp98392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																															tmp98393 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp98394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp98395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp98396 := Call(__e, tmp98395, B, Nil)

																															tmp98397 := Call(__e, tmp98394, sym_d, tmp98396)

																															tmp98398 := Call(__e, tmp98393, A, tmp98397)

																															tmp98399 := Call(__e, tmp98392, V2559, tmp98398, V2791)

																															_ = tmp98399

																															tmp98400 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp98401 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																tmp98402 := Call(__e, tmp98401, V2559, V2791)

																																_ = tmp98402

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp98403 := MakeNative(func(__e *ControlFlow) {
																																V2573 := __e.Get(1)
																																_ = V2573
																																tmp98478 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp98479 := Call(__e, tmp98478, Nil, V2573)

																																if True == tmp98479 {
																																	tmp98446 := MakeNative(func(__e *ControlFlow) {
																																		Hyp := __e.Get(1)
																																		_ = Hyp
																																		tmp98447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																		tmp98448 := Call(__e, tmp98447)

																																		_ = tmp98448

																																		tmp98449 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																		tmp98450 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98451 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98453 := Call(__e, tmp98452, X, V2791)

																																		tmp98454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98456 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98457 := Call(__e, tmp98456, A, V2791)

																																		tmp98458 := Call(__e, tmp98455, tmp98457, Nil)

																																		tmp98459 := Call(__e, tmp98454, sym_h, tmp98458)

																																		tmp98460 := Call(__e, tmp98451, tmp98453, tmp98459)

																																		tmp98461 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98462 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98464 := Call(__e, tmp98463, Y, V2791)

																																		tmp98465 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98466 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																		tmp98467 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98468 := Call(__e, tmp98467, B, V2791)

																																		tmp98469 := Call(__e, tmp98466, tmp98468, Nil)

																																		tmp98470 := Call(__e, tmp98465, sym_h, tmp98469)

																																		tmp98471 := Call(__e, tmp98462, tmp98464, tmp98470)

																																		tmp98472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp98473 := Call(__e, tmp98472, Hyp, V2791)

																																		tmp98474 := Call(__e, tmp98461, tmp98471, tmp98473)

																																		tmp98475 := Call(__e, tmp98450, tmp98460, tmp98474)

																																		__e.TailApply(tmp98449, V2790, tmp98475, V2791, V2792)
																																		return

																																	}, 1)

																																	tmp98476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp98477 := Call(__e, tmp98476, V2549)

																																	__e.TailApply(tmp98446, tmp98477)
																																	return

																																} else {
																																	tmp98444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																	tmp98445 := Call(__e, tmp98444, V2573)

																																	if True == tmp98445 {
																																		tmp98406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																		tmp98407 := Call(__e, tmp98406, V2573, Nil, V2791)

																																		_ = tmp98407

																																		tmp98408 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp98409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																			tmp98410 := Call(__e, tmp98409, V2573, V2791)

																																			_ = tmp98410

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp98411 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp98412 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp98413 := Call(__e, tmp98412)

																																			_ = tmp98413

																																			tmp98414 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																			tmp98415 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98416 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98417 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98418 := Call(__e, tmp98417, X, V2791)

																																			tmp98419 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98420 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98421 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98422 := Call(__e, tmp98421, A, V2791)

																																			tmp98423 := Call(__e, tmp98420, tmp98422, Nil)

																																			tmp98424 := Call(__e, tmp98419, sym_h, tmp98423)

																																			tmp98425 := Call(__e, tmp98416, tmp98418, tmp98424)

																																			tmp98426 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98427 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98429 := Call(__e, tmp98428, Y, V2791)

																																			tmp98430 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98431 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp98432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98433 := Call(__e, tmp98432, B, V2791)

																																			tmp98434 := Call(__e, tmp98431, tmp98433, Nil)

																																			tmp98435 := Call(__e, tmp98430, sym_h, tmp98434)

																																			tmp98436 := Call(__e, tmp98427, tmp98429, tmp98435)

																																			tmp98437 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp98438 := Call(__e, tmp98437, Hyp, V2791)

																																			tmp98439 := Call(__e, tmp98426, tmp98436, tmp98438)

																																			tmp98440 := Call(__e, tmp98415, tmp98425, tmp98439)

																																			__e.TailApply(tmp98414, V2790, tmp98440, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp98441 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp98442 := Call(__e, tmp98441, V2549)

																																		tmp98443 := Call(__e, tmp98411, tmp98442)

																																		__e.TailApply(tmp98408, tmp98443)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp98480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp98481 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp98482 := Call(__e, tmp98481, V2558)

																															tmp98483 := Call(__e, tmp98480, tmp98482, V2791)

																															tmp98484 := Call(__e, tmp98403, tmp98483)

																															__e.TailApply(tmp98400, tmp98484)
																															return

																														}, 1)

																														tmp98485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																														tmp98486 := Call(__e, tmp98485, V2791)

																														__e.TailApply(tmp98391, tmp98486)
																														return

																													}, 1)

																													tmp98487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																													tmp98488 := Call(__e, tmp98487, V2791)

																													__e.TailApply(tmp98390, tmp98488)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}

																										}, 1)

																										tmp99188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																										tmp99189 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										tmp99190 := Call(__e, tmp99189, V2558)

																										tmp99191 := Call(__e, tmp99188, tmp99190, V2791)

																										__e.TailApply(tmp98387, tmp99191)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp99194 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																								tmp99195 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																								tmp99196 := Call(__e, tmp99195, V2556)

																								tmp99197 := Call(__e, tmp99194, tmp99196, V2791)

																								__e.TailApply(tmp98385, tmp99197)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp99200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp99201 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																						tmp99202 := Call(__e, tmp99201, V2556)

																						tmp99203 := Call(__e, tmp99200, tmp99202, V2791)

																						__e.TailApply(tmp98383, tmp99203)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp99206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp99207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp99208 := Call(__e, tmp99207, V2550)

																				tmp99209 := Call(__e, tmp99206, tmp99208, V2791)

																				__e.TailApply(tmp98381, tmp99209)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp99212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																		tmp99213 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp99214 := Call(__e, tmp99213, V2554)

																		tmp99215 := Call(__e, tmp99212, tmp99214, V2791)

																		__e.TailApply(tmp98379, tmp99215)
																		return

																	}, 1)

																	tmp99216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp99217 := Call(__e, tmp99216, V2554)

																	__e.TailApply(tmp98378, tmp99217)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp99220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp99221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp99222 := Call(__e, tmp99221, V2553)

															tmp99223 := Call(__e, tmp99220, tmp99222, V2791)

															__e.TailApply(tmp98376, tmp99223)
															return

														}, 1)

														tmp99224 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp99225 := Call(__e, tmp99224, V2553)

														__e.TailApply(tmp98375, tmp99225)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp99228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp99229 := Call(__e, PrimNS1Value(symns2_1value), symtl)

												tmp99230 := Call(__e, tmp99229, V2551)

												tmp99231 := Call(__e, tmp99228, tmp99230, V2791)

												__e.TailApply(tmp98373, tmp99231)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp99234 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp99235 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp99236 := Call(__e, tmp99235, V2551)

										tmp99237 := Call(__e, tmp99234, tmp99236, V2791)

										__e.TailApply(tmp98371, tmp99237)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp99240 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp99241 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp99242 := Call(__e, tmp99241, V2550)

								tmp99243 := Call(__e, tmp99240, tmp99242, V2791)

								__e.TailApply(tmp98369, tmp99243)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp99246 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp99247 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp99248 := Call(__e, tmp99247, V2549)

						tmp99249 := Call(__e, tmp99246, tmp99248, V2791)

						__e.TailApply(tmp98367, tmp99249)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp99252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp99253 := Call(__e, tmp99252, V2789, V2791)

				tmp99254 := Call(__e, tmp98365, tmp99253)

				__e.TailApply(tmp97246, tmp99254)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp99257 := MakeNative(func(__e *ControlFlow) {
			V2526 := __e.Get(1)
			_ = V2526
			tmp100085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100086 := Call(__e, tmp100085, V2526)

			if True == tmp100086 {
				tmp99259 := MakeNative(func(__e *ControlFlow) {
					V2527 := __e.Get(1)
					_ = V2527
					tmp100079 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp100080 := Call(__e, tmp100079, V2527)

					if True == tmp100080 {
						tmp99261 := MakeNative(func(__e *ControlFlow) {
							V2528 := __e.Get(1)
							_ = V2528
							tmp100073 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp100074 := Call(__e, tmp100073, V2528)

							if True == tmp100074 {
								tmp99263 := MakeNative(func(__e *ControlFlow) {
									V2529 := __e.Get(1)
									_ = V2529
									tmp100067 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp100068 := Call(__e, tmp100067, symcons, V2529)

									if True == tmp100068 {
										tmp99265 := MakeNative(func(__e *ControlFlow) {
											V2530 := __e.Get(1)
											_ = V2530
											tmp100061 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp100062 := Call(__e, tmp100061, V2530)

											if True == tmp100062 {
												tmp99267 := MakeNative(func(__e *ControlFlow) {
													X := __e.Get(1)
													_ = X
													tmp99268 := MakeNative(func(__e *ControlFlow) {
														V2531 := __e.Get(1)
														_ = V2531
														tmp100053 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp100054 := Call(__e, tmp100053, V2531)

														if True == tmp100054 {
															tmp99270 := MakeNative(func(__e *ControlFlow) {
																Y := __e.Get(1)
																_ = Y
																tmp99271 := MakeNative(func(__e *ControlFlow) {
																	V2532 := __e.Get(1)
																	_ = V2532
																	tmp100045 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																	tmp100046 := Call(__e, tmp100045, Nil, V2532)

																	if True == tmp100046 {
																		tmp99273 := MakeNative(func(__e *ControlFlow) {
																			V2533 := __e.Get(1)
																			_ = V2533
																			tmp100039 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																			tmp100040 := Call(__e, tmp100039, V2533)

																			if True == tmp100040 {
																				tmp99275 := MakeNative(func(__e *ControlFlow) {
																					V2534 := __e.Get(1)
																					_ = V2534
																					tmp100033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp100034 := Call(__e, tmp100033, sym_h, V2534)

																					if True == tmp100034 {
																						tmp99277 := MakeNative(func(__e *ControlFlow) {
																							V2535 := __e.Get(1)
																							_ = V2535
																							tmp100027 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																							tmp100028 := Call(__e, tmp100027, V2535)

																							if True == tmp100028 {
																								tmp99279 := MakeNative(func(__e *ControlFlow) {
																									V2536 := __e.Get(1)
																									_ = V2536
																									tmp100021 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																									tmp100022 := Call(__e, tmp100021, V2536)

																									if True == tmp100022 {
																										tmp99386 := MakeNative(func(__e *ControlFlow) {
																											V2537 := __e.Get(1)
																											_ = V2537
																											tmp100015 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																											tmp100016 := Call(__e, tmp100015, symlist, V2537)

																											if True == tmp100016 {
																												tmp99706 := MakeNative(func(__e *ControlFlow) {
																													V2538 := __e.Get(1)
																													_ = V2538
																													tmp100009 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																													tmp100010 := Call(__e, tmp100009, V2538)

																													if True == tmp100010 {
																														tmp99811 := MakeNative(func(__e *ControlFlow) {
																															A := __e.Get(1)
																															_ = A
																															tmp99812 := MakeNative(func(__e *ControlFlow) {
																																V2539 := __e.Get(1)
																																_ = V2539
																																tmp100001 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																tmp100002 := Call(__e, tmp100001, Nil, V2539)

																																if True == tmp100002 {
																																	tmp99912 := MakeNative(func(__e *ControlFlow) {
																																		V2540 := __e.Get(1)
																																		_ = V2540
																																		tmp99995 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp99996 := Call(__e, tmp99995, Nil, V2540)

																																		if True == tmp99996 {
																																			tmp99959 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp99960 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp99961 := Call(__e, tmp99960)

																																				_ = tmp99961

																																				tmp99962 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp99963 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99964 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99966 := Call(__e, tmp99965, X, V2791)

																																				tmp99967 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99970 := Call(__e, tmp99969, A, V2791)

																																				tmp99971 := Call(__e, tmp99968, tmp99970, Nil)

																																				tmp99972 := Call(__e, tmp99967, sym_h, tmp99971)

																																				tmp99973 := Call(__e, tmp99964, tmp99966, tmp99972)

																																				tmp99974 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99975 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99977 := Call(__e, tmp99976, Y, V2791)

																																				tmp99978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99979 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99980 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99981 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99982 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99983 := Call(__e, tmp99982, A, V2791)

																																				tmp99984 := Call(__e, tmp99981, tmp99983, Nil)

																																				tmp99985 := Call(__e, tmp99980, symlist, tmp99984)

																																				tmp99986 := Call(__e, tmp99979, tmp99985, Nil)

																																				tmp99987 := Call(__e, tmp99978, sym_h, tmp99986)

																																				tmp99988 := Call(__e, tmp99975, tmp99977, tmp99987)

																																				tmp99989 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99990 := Call(__e, tmp99989, Hyp, V2791)

																																				tmp99991 := Call(__e, tmp99974, tmp99988, tmp99990)

																																				tmp99992 := Call(__e, tmp99963, tmp99973, tmp99991)

																																				__e.TailApply(tmp99962, V2790, tmp99992, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp99993 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp99994 := Call(__e, tmp99993, V2526)

																																			__e.TailApply(tmp99959, tmp99994)
																																			return

																																		} else {
																																			tmp99957 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp99958 := Call(__e, tmp99957, V2540)

																																			if True == tmp99958 {
																																				tmp99915 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp99916 := Call(__e, tmp99915, V2540, Nil, V2791)

																																				_ = tmp99916

																																				tmp99917 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp99918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp99919 := Call(__e, tmp99918, V2540, V2791)

																																					_ = tmp99919

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp99920 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp99921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp99922 := Call(__e, tmp99921)

																																					_ = tmp99922

																																					tmp99923 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp99924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99927 := Call(__e, tmp99926, X, V2791)

																																					tmp99928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99931 := Call(__e, tmp99930, A, V2791)

																																					tmp99932 := Call(__e, tmp99929, tmp99931, Nil)

																																					tmp99933 := Call(__e, tmp99928, sym_h, tmp99932)

																																					tmp99934 := Call(__e, tmp99925, tmp99927, tmp99933)

																																					tmp99935 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99936 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99937 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99938 := Call(__e, tmp99937, Y, V2791)

																																					tmp99939 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99940 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99941 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99942 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99944 := Call(__e, tmp99943, A, V2791)

																																					tmp99945 := Call(__e, tmp99942, tmp99944, Nil)

																																					tmp99946 := Call(__e, tmp99941, symlist, tmp99945)

																																					tmp99947 := Call(__e, tmp99940, tmp99946, Nil)

																																					tmp99948 := Call(__e, tmp99939, sym_h, tmp99947)

																																					tmp99949 := Call(__e, tmp99936, tmp99938, tmp99948)

																																					tmp99950 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99951 := Call(__e, tmp99950, Hyp, V2791)

																																					tmp99952 := Call(__e, tmp99935, tmp99949, tmp99951)

																																					tmp99953 := Call(__e, tmp99924, tmp99934, tmp99952)

																																					__e.TailApply(tmp99923, V2790, tmp99953, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp99954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp99955 := Call(__e, tmp99954, V2526)

																																				tmp99956 := Call(__e, tmp99920, tmp99955)

																																				__e.TailApply(tmp99917, tmp99956)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp99997 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp99998 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp99999 := Call(__e, tmp99998, V2535)

																																	tmp100000 := Call(__e, tmp99997, tmp99999, V2791)

																																	__e.TailApply(tmp99912, tmp100000)
																																	return

																																} else {
																																	tmp99910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																	tmp99911 := Call(__e, tmp99910, V2539)

																																	if True == tmp99911 {
																																		tmp99815 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																		tmp99816 := Call(__e, tmp99815, V2539, Nil, V2791)

																																		_ = tmp99816

																																		tmp99817 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp99818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																			tmp99819 := Call(__e, tmp99818, V2539, V2791)

																																			_ = tmp99819

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp99820 := MakeNative(func(__e *ControlFlow) {
																																			V2541 := __e.Get(1)
																																			_ = V2541
																																			tmp99903 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																			tmp99904 := Call(__e, tmp99903, Nil, V2541)

																																			if True == tmp99904 {
																																				tmp99867 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp99868 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp99869 := Call(__e, tmp99868)

																																					_ = tmp99869

																																					tmp99870 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp99871 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99872 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99873 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99874 := Call(__e, tmp99873, X, V2791)

																																					tmp99875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99876 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99877 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99878 := Call(__e, tmp99877, A, V2791)

																																					tmp99879 := Call(__e, tmp99876, tmp99878, Nil)

																																					tmp99880 := Call(__e, tmp99875, sym_h, tmp99879)

																																					tmp99881 := Call(__e, tmp99872, tmp99874, tmp99880)

																																					tmp99882 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99883 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99884 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99885 := Call(__e, tmp99884, Y, V2791)

																																					tmp99886 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99887 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99888 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99890 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99891 := Call(__e, tmp99890, A, V2791)

																																					tmp99892 := Call(__e, tmp99889, tmp99891, Nil)

																																					tmp99893 := Call(__e, tmp99888, symlist, tmp99892)

																																					tmp99894 := Call(__e, tmp99887, tmp99893, Nil)

																																					tmp99895 := Call(__e, tmp99886, sym_h, tmp99894)

																																					tmp99896 := Call(__e, tmp99883, tmp99885, tmp99895)

																																					tmp99897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99898 := Call(__e, tmp99897, Hyp, V2791)

																																					tmp99899 := Call(__e, tmp99882, tmp99896, tmp99898)

																																					tmp99900 := Call(__e, tmp99871, tmp99881, tmp99899)

																																					__e.TailApply(tmp99870, V2790, tmp99900, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp99901 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp99902 := Call(__e, tmp99901, V2526)

																																				__e.TailApply(tmp99867, tmp99902)
																																				return

																																			} else {
																																				tmp99865 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																				tmp99866 := Call(__e, tmp99865, V2541)

																																				if True == tmp99866 {
																																					tmp99823 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp99824 := Call(__e, tmp99823, V2541, Nil, V2791)

																																					_ = tmp99824

																																					tmp99825 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp99826 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp99827 := Call(__e, tmp99826, V2541, V2791)

																																						_ = tmp99827

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp99828 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp99829 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																						tmp99830 := Call(__e, tmp99829)

																																						_ = tmp99830

																																						tmp99831 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																						tmp99832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99834 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99835 := Call(__e, tmp99834, X, V2791)

																																						tmp99836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99839 := Call(__e, tmp99838, A, V2791)

																																						tmp99840 := Call(__e, tmp99837, tmp99839, Nil)

																																						tmp99841 := Call(__e, tmp99836, sym_h, tmp99840)

																																						tmp99842 := Call(__e, tmp99833, tmp99835, tmp99841)

																																						tmp99843 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99844 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99845 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99846 := Call(__e, tmp99845, Y, V2791)

																																						tmp99847 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99848 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99851 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99852 := Call(__e, tmp99851, A, V2791)

																																						tmp99853 := Call(__e, tmp99850, tmp99852, Nil)

																																						tmp99854 := Call(__e, tmp99849, symlist, tmp99853)

																																						tmp99855 := Call(__e, tmp99848, tmp99854, Nil)

																																						tmp99856 := Call(__e, tmp99847, sym_h, tmp99855)

																																						tmp99857 := Call(__e, tmp99844, tmp99846, tmp99856)

																																						tmp99858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99859 := Call(__e, tmp99858, Hyp, V2791)

																																						tmp99860 := Call(__e, tmp99843, tmp99857, tmp99859)

																																						tmp99861 := Call(__e, tmp99832, tmp99842, tmp99860)

																																						__e.TailApply(tmp99831, V2790, tmp99861, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp99862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp99863 := Call(__e, tmp99862, V2526)

																																					tmp99864 := Call(__e, tmp99828, tmp99863)

																																					__e.TailApply(tmp99825, tmp99864)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp99905 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp99906 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp99907 := Call(__e, tmp99906, V2535)

																																		tmp99908 := Call(__e, tmp99905, tmp99907, V2791)

																																		tmp99909 := Call(__e, tmp99820, tmp99908)

																																		__e.TailApply(tmp99817, tmp99909)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp100003 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp100004 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp100005 := Call(__e, tmp100004, V2538)

																															tmp100006 := Call(__e, tmp100003, tmp100005, V2791)

																															__e.TailApply(tmp99812, tmp100006)
																															return

																														}, 1)

																														tmp100007 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																														tmp100008 := Call(__e, tmp100007, V2538)

																														__e.TailApply(tmp99811, tmp100008)
																														return

																													} else {
																														tmp99809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																														tmp99810 := Call(__e, tmp99809, V2538)

																														if True == tmp99810 {
																															tmp99709 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp99710 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																tmp99711 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99712 := Call(__e, tmp99711, A, Nil)

																																tmp99713 := Call(__e, tmp99710, V2538, tmp99712, V2791)

																																_ = tmp99713

																																tmp99714 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp99715 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																	tmp99716 := Call(__e, tmp99715, V2538, V2791)

																																	_ = tmp99716

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp99717 := MakeNative(func(__e *ControlFlow) {
																																	V2542 := __e.Get(1)
																																	_ = V2542
																																	tmp99800 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																	tmp99801 := Call(__e, tmp99800, Nil, V2542)

																																	if True == tmp99801 {
																																		tmp99764 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp99765 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																			tmp99766 := Call(__e, tmp99765)

																																			_ = tmp99766

																																			tmp99767 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																			tmp99768 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99769 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99770 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99771 := Call(__e, tmp99770, X, V2791)

																																			tmp99772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99773 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99774 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99775 := Call(__e, tmp99774, A, V2791)

																																			tmp99776 := Call(__e, tmp99773, tmp99775, Nil)

																																			tmp99777 := Call(__e, tmp99772, sym_h, tmp99776)

																																			tmp99778 := Call(__e, tmp99769, tmp99771, tmp99777)

																																			tmp99779 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99780 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99782 := Call(__e, tmp99781, Y, V2791)

																																			tmp99783 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99784 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99785 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99786 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																			tmp99787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99788 := Call(__e, tmp99787, A, V2791)

																																			tmp99789 := Call(__e, tmp99786, tmp99788, Nil)

																																			tmp99790 := Call(__e, tmp99785, symlist, tmp99789)

																																			tmp99791 := Call(__e, tmp99784, tmp99790, Nil)

																																			tmp99792 := Call(__e, tmp99783, sym_h, tmp99791)

																																			tmp99793 := Call(__e, tmp99780, tmp99782, tmp99792)

																																			tmp99794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99795 := Call(__e, tmp99794, Hyp, V2791)

																																			tmp99796 := Call(__e, tmp99779, tmp99793, tmp99795)

																																			tmp99797 := Call(__e, tmp99768, tmp99778, tmp99796)

																																			__e.TailApply(tmp99767, V2790, tmp99797, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp99798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp99799 := Call(__e, tmp99798, V2526)

																																		__e.TailApply(tmp99764, tmp99799)
																																		return

																																	} else {
																																		tmp99762 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp99763 := Call(__e, tmp99762, V2542)

																																		if True == tmp99763 {
																																			tmp99720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp99721 := Call(__e, tmp99720, V2542, Nil, V2791)

																																			_ = tmp99721

																																			tmp99722 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp99723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp99724 := Call(__e, tmp99723, V2542, V2791)

																																				_ = tmp99724

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp99725 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp99726 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp99727 := Call(__e, tmp99726)

																																				_ = tmp99727

																																				tmp99728 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp99729 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99732 := Call(__e, tmp99731, X, V2791)

																																				tmp99733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99735 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99736 := Call(__e, tmp99735, A, V2791)

																																				tmp99737 := Call(__e, tmp99734, tmp99736, Nil)

																																				tmp99738 := Call(__e, tmp99733, sym_h, tmp99737)

																																				tmp99739 := Call(__e, tmp99730, tmp99732, tmp99738)

																																				tmp99740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99743 := Call(__e, tmp99742, Y, V2791)

																																				tmp99744 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99745 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99746 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99747 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99749 := Call(__e, tmp99748, A, V2791)

																																				tmp99750 := Call(__e, tmp99747, tmp99749, Nil)

																																				tmp99751 := Call(__e, tmp99746, symlist, tmp99750)

																																				tmp99752 := Call(__e, tmp99745, tmp99751, Nil)

																																				tmp99753 := Call(__e, tmp99744, sym_h, tmp99752)

																																				tmp99754 := Call(__e, tmp99741, tmp99743, tmp99753)

																																				tmp99755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99756 := Call(__e, tmp99755, Hyp, V2791)

																																				tmp99757 := Call(__e, tmp99740, tmp99754, tmp99756)

																																				tmp99758 := Call(__e, tmp99729, tmp99739, tmp99757)

																																				__e.TailApply(tmp99728, V2790, tmp99758, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp99759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp99760 := Call(__e, tmp99759, V2526)

																																			tmp99761 := Call(__e, tmp99725, tmp99760)

																																			__e.TailApply(tmp99722, tmp99761)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp99802 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99803 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp99804 := Call(__e, tmp99803, V2535)

																																tmp99805 := Call(__e, tmp99802, tmp99804, V2791)

																																tmp99806 := Call(__e, tmp99717, tmp99805)

																																__e.TailApply(tmp99714, tmp99806)
																																return

																															}, 1)

																															tmp99807 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																															tmp99808 := Call(__e, tmp99807, V2791)

																															__e.TailApply(tmp99709, tmp99808)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp100011 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																												tmp100012 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																												tmp100013 := Call(__e, tmp100012, V2536)

																												tmp100014 := Call(__e, tmp100011, tmp100013, V2791)

																												__e.TailApply(tmp99706, tmp100014)
																												return

																											} else {
																												tmp99704 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																												tmp99705 := Call(__e, tmp99704, V2537)

																												if True == tmp99705 {
																													tmp99389 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																													tmp99390 := Call(__e, tmp99389, V2537, symlist, V2791)

																													_ = tmp99390

																													tmp99391 := MakeNative(func(__e *ControlFlow) {
																														Result := __e.Get(1)
																														_ = Result
																														tmp99392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																														tmp99393 := Call(__e, tmp99392, V2537, V2791)

																														_ = tmp99393

																														__e.Return(Result)
																														return

																													}, 1)

																													tmp99394 := MakeNative(func(__e *ControlFlow) {
																														V2543 := __e.Get(1)
																														_ = V2543
																														tmp99697 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																														tmp99698 := Call(__e, tmp99697, V2543)

																														if True == tmp99698 {
																															tmp99499 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp99500 := MakeNative(func(__e *ControlFlow) {
																																	V2544 := __e.Get(1)
																																	_ = V2544
																																	tmp99689 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																	tmp99690 := Call(__e, tmp99689, Nil, V2544)

																																	if True == tmp99690 {
																																		tmp99600 := MakeNative(func(__e *ControlFlow) {
																																			V2545 := __e.Get(1)
																																			_ = V2545
																																			tmp99683 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																			tmp99684 := Call(__e, tmp99683, Nil, V2545)

																																			if True == tmp99684 {
																																				tmp99647 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp99648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp99649 := Call(__e, tmp99648)

																																					_ = tmp99649

																																					tmp99650 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp99651 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99653 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99654 := Call(__e, tmp99653, X, V2791)

																																					tmp99655 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99656 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99658 := Call(__e, tmp99657, A, V2791)

																																					tmp99659 := Call(__e, tmp99656, tmp99658, Nil)

																																					tmp99660 := Call(__e, tmp99655, sym_h, tmp99659)

																																					tmp99661 := Call(__e, tmp99652, tmp99654, tmp99660)

																																					tmp99662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99664 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99665 := Call(__e, tmp99664, Y, V2791)

																																					tmp99666 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99668 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99669 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99670 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99671 := Call(__e, tmp99670, A, V2791)

																																					tmp99672 := Call(__e, tmp99669, tmp99671, Nil)

																																					tmp99673 := Call(__e, tmp99668, symlist, tmp99672)

																																					tmp99674 := Call(__e, tmp99667, tmp99673, Nil)

																																					tmp99675 := Call(__e, tmp99666, sym_h, tmp99674)

																																					tmp99676 := Call(__e, tmp99663, tmp99665, tmp99675)

																																					tmp99677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99678 := Call(__e, tmp99677, Hyp, V2791)

																																					tmp99679 := Call(__e, tmp99662, tmp99676, tmp99678)

																																					tmp99680 := Call(__e, tmp99651, tmp99661, tmp99679)

																																					__e.TailApply(tmp99650, V2790, tmp99680, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp99681 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp99682 := Call(__e, tmp99681, V2526)

																																				__e.TailApply(tmp99647, tmp99682)
																																				return

																																			} else {
																																				tmp99645 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																				tmp99646 := Call(__e, tmp99645, V2545)

																																				if True == tmp99646 {
																																					tmp99603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																					tmp99604 := Call(__e, tmp99603, V2545, Nil, V2791)

																																					_ = tmp99604

																																					tmp99605 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp99606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																						tmp99607 := Call(__e, tmp99606, V2545, V2791)

																																						_ = tmp99607

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp99608 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp99609 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																						tmp99610 := Call(__e, tmp99609)

																																						_ = tmp99610

																																						tmp99611 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																						tmp99612 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99613 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99614 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99615 := Call(__e, tmp99614, X, V2791)

																																						tmp99616 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99618 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99619 := Call(__e, tmp99618, A, V2791)

																																						tmp99620 := Call(__e, tmp99617, tmp99619, Nil)

																																						tmp99621 := Call(__e, tmp99616, sym_h, tmp99620)

																																						tmp99622 := Call(__e, tmp99613, tmp99615, tmp99621)

																																						tmp99623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99625 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99626 := Call(__e, tmp99625, Y, V2791)

																																						tmp99627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99629 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99630 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99631 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99632 := Call(__e, tmp99631, A, V2791)

																																						tmp99633 := Call(__e, tmp99630, tmp99632, Nil)

																																						tmp99634 := Call(__e, tmp99629, symlist, tmp99633)

																																						tmp99635 := Call(__e, tmp99628, tmp99634, Nil)

																																						tmp99636 := Call(__e, tmp99627, sym_h, tmp99635)

																																						tmp99637 := Call(__e, tmp99624, tmp99626, tmp99636)

																																						tmp99638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99639 := Call(__e, tmp99638, Hyp, V2791)

																																						tmp99640 := Call(__e, tmp99623, tmp99637, tmp99639)

																																						tmp99641 := Call(__e, tmp99612, tmp99622, tmp99640)

																																						__e.TailApply(tmp99611, V2790, tmp99641, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp99642 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp99643 := Call(__e, tmp99642, V2526)

																																					tmp99644 := Call(__e, tmp99608, tmp99643)

																																					__e.TailApply(tmp99605, tmp99644)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp99685 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																		tmp99686 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																		tmp99687 := Call(__e, tmp99686, V2535)

																																		tmp99688 := Call(__e, tmp99685, tmp99687, V2791)

																																		__e.TailApply(tmp99600, tmp99688)
																																		return

																																	} else {
																																		tmp99598 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																		tmp99599 := Call(__e, tmp99598, V2544)

																																		if True == tmp99599 {
																																			tmp99503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																			tmp99504 := Call(__e, tmp99503, V2544, Nil, V2791)

																																			_ = tmp99504

																																			tmp99505 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp99506 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																				tmp99507 := Call(__e, tmp99506, V2544, V2791)

																																				_ = tmp99507

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp99508 := MakeNative(func(__e *ControlFlow) {
																																				V2546 := __e.Get(1)
																																				_ = V2546
																																				tmp99591 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																				tmp99592 := Call(__e, tmp99591, Nil, V2546)

																																				if True == tmp99592 {
																																					tmp99555 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp99556 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																						tmp99557 := Call(__e, tmp99556)

																																						_ = tmp99557

																																						tmp99558 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																						tmp99559 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99561 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99562 := Call(__e, tmp99561, X, V2791)

																																						tmp99563 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99564 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99566 := Call(__e, tmp99565, A, V2791)

																																						tmp99567 := Call(__e, tmp99564, tmp99566, Nil)

																																						tmp99568 := Call(__e, tmp99563, sym_h, tmp99567)

																																						tmp99569 := Call(__e, tmp99560, tmp99562, tmp99568)

																																						tmp99570 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99573 := Call(__e, tmp99572, Y, V2791)

																																						tmp99574 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99575 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99576 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99577 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																						tmp99578 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99579 := Call(__e, tmp99578, A, V2791)

																																						tmp99580 := Call(__e, tmp99577, tmp99579, Nil)

																																						tmp99581 := Call(__e, tmp99576, symlist, tmp99580)

																																						tmp99582 := Call(__e, tmp99575, tmp99581, Nil)

																																						tmp99583 := Call(__e, tmp99574, sym_h, tmp99582)

																																						tmp99584 := Call(__e, tmp99571, tmp99573, tmp99583)

																																						tmp99585 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																						tmp99586 := Call(__e, tmp99585, Hyp, V2791)

																																						tmp99587 := Call(__e, tmp99570, tmp99584, tmp99586)

																																						tmp99588 := Call(__e, tmp99559, tmp99569, tmp99587)

																																						__e.TailApply(tmp99558, V2790, tmp99588, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp99589 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																					tmp99590 := Call(__e, tmp99589, V2526)

																																					__e.TailApply(tmp99555, tmp99590)
																																					return

																																				} else {
																																					tmp99553 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																					tmp99554 := Call(__e, tmp99553, V2546)

																																					if True == tmp99554 {
																																						tmp99511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																						tmp99512 := Call(__e, tmp99511, V2546, Nil, V2791)

																																						_ = tmp99512

																																						tmp99513 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp99514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																							tmp99515 := Call(__e, tmp99514, V2546, V2791)

																																							_ = tmp99515

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp99516 := MakeNative(func(__e *ControlFlow) {
																																							Hyp := __e.Get(1)
																																							_ = Hyp
																																							tmp99517 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																							tmp99518 := Call(__e, tmp99517)

																																							_ = tmp99518

																																							tmp99519 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																							tmp99520 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99521 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99523 := Call(__e, tmp99522, X, V2791)

																																							tmp99524 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99525 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99527 := Call(__e, tmp99526, A, V2791)

																																							tmp99528 := Call(__e, tmp99525, tmp99527, Nil)

																																							tmp99529 := Call(__e, tmp99524, sym_h, tmp99528)

																																							tmp99530 := Call(__e, tmp99521, tmp99523, tmp99529)

																																							tmp99531 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99533 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99534 := Call(__e, tmp99533, Y, V2791)

																																							tmp99535 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99536 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99537 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99538 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																							tmp99539 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99540 := Call(__e, tmp99539, A, V2791)

																																							tmp99541 := Call(__e, tmp99538, tmp99540, Nil)

																																							tmp99542 := Call(__e, tmp99537, symlist, tmp99541)

																																							tmp99543 := Call(__e, tmp99536, tmp99542, Nil)

																																							tmp99544 := Call(__e, tmp99535, sym_h, tmp99543)

																																							tmp99545 := Call(__e, tmp99532, tmp99534, tmp99544)

																																							tmp99546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																							tmp99547 := Call(__e, tmp99546, Hyp, V2791)

																																							tmp99548 := Call(__e, tmp99531, tmp99545, tmp99547)

																																							tmp99549 := Call(__e, tmp99520, tmp99530, tmp99548)

																																							__e.TailApply(tmp99519, V2790, tmp99549, V2791, V2792)
																																							return

																																						}, 1)

																																						tmp99550 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																						tmp99551 := Call(__e, tmp99550, V2526)

																																						tmp99552 := Call(__e, tmp99516, tmp99551)

																																						__e.TailApply(tmp99513, tmp99552)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp99593 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																			tmp99594 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp99595 := Call(__e, tmp99594, V2535)

																																			tmp99596 := Call(__e, tmp99593, tmp99595, V2791)

																																			tmp99597 := Call(__e, tmp99508, tmp99596)

																																			__e.TailApply(tmp99505, tmp99597)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp99691 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99692 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																tmp99693 := Call(__e, tmp99692, V2543)

																																tmp99694 := Call(__e, tmp99691, tmp99693, V2791)

																																__e.TailApply(tmp99500, tmp99694)
																																return

																															}, 1)

																															tmp99695 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																															tmp99696 := Call(__e, tmp99695, V2543)

																															__e.TailApply(tmp99499, tmp99696)
																															return

																														} else {
																															tmp99497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																															tmp99498 := Call(__e, tmp99497, V2543)

																															if True == tmp99498 {
																																tmp99397 := MakeNative(func(__e *ControlFlow) {
																																	A := __e.Get(1)
																																	_ = A
																																	tmp99398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																	tmp99399 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																	tmp99400 := Call(__e, tmp99399, A, Nil)

																																	tmp99401 := Call(__e, tmp99398, V2543, tmp99400, V2791)

																																	_ = tmp99401

																																	tmp99402 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp99403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																		tmp99404 := Call(__e, tmp99403, V2543, V2791)

																																		_ = tmp99404

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp99405 := MakeNative(func(__e *ControlFlow) {
																																		V2547 := __e.Get(1)
																																		_ = V2547
																																		tmp99488 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																																		tmp99489 := Call(__e, tmp99488, Nil, V2547)

																																		if True == tmp99489 {
																																			tmp99452 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp99453 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																				tmp99454 := Call(__e, tmp99453)

																																				_ = tmp99454

																																				tmp99455 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																				tmp99456 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99457 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99459 := Call(__e, tmp99458, X, V2791)

																																				tmp99460 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99461 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99463 := Call(__e, tmp99462, A, V2791)

																																				tmp99464 := Call(__e, tmp99461, tmp99463, Nil)

																																				tmp99465 := Call(__e, tmp99460, sym_h, tmp99464)

																																				tmp99466 := Call(__e, tmp99457, tmp99459, tmp99465)

																																				tmp99467 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99468 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99469 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99470 := Call(__e, tmp99469, Y, V2791)

																																				tmp99471 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99472 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																				tmp99475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99476 := Call(__e, tmp99475, A, V2791)

																																				tmp99477 := Call(__e, tmp99474, tmp99476, Nil)

																																				tmp99478 := Call(__e, tmp99473, symlist, tmp99477)

																																				tmp99479 := Call(__e, tmp99472, tmp99478, Nil)

																																				tmp99480 := Call(__e, tmp99471, sym_h, tmp99479)

																																				tmp99481 := Call(__e, tmp99468, tmp99470, tmp99480)

																																				tmp99482 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																				tmp99483 := Call(__e, tmp99482, Hyp, V2791)

																																				tmp99484 := Call(__e, tmp99467, tmp99481, tmp99483)

																																				tmp99485 := Call(__e, tmp99456, tmp99466, tmp99484)

																																				__e.TailApply(tmp99455, V2790, tmp99485, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp99486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																			tmp99487 := Call(__e, tmp99486, V2526)

																																			__e.TailApply(tmp99452, tmp99487)
																																			return

																																		} else {
																																			tmp99450 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																																			tmp99451 := Call(__e, tmp99450, V2547)

																																			if True == tmp99451 {
																																				tmp99408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																																				tmp99409 := Call(__e, tmp99408, V2547, Nil, V2791)

																																				_ = tmp99409

																																				tmp99410 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp99411 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																					tmp99412 := Call(__e, tmp99411, V2547, V2791)

																																					_ = tmp99412

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp99413 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp99414 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																					tmp99415 := Call(__e, tmp99414)

																																					_ = tmp99415

																																					tmp99416 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																					tmp99417 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99418 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99419 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99420 := Call(__e, tmp99419, X, V2791)

																																					tmp99421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99422 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99423 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99424 := Call(__e, tmp99423, A, V2791)

																																					tmp99425 := Call(__e, tmp99422, tmp99424, Nil)

																																					tmp99426 := Call(__e, tmp99421, sym_h, tmp99425)

																																					tmp99427 := Call(__e, tmp99418, tmp99420, tmp99426)

																																					tmp99428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99430 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99431 := Call(__e, tmp99430, Y, V2791)

																																					tmp99432 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99434 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99435 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																					tmp99436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99437 := Call(__e, tmp99436, A, V2791)

																																					tmp99438 := Call(__e, tmp99435, tmp99437, Nil)

																																					tmp99439 := Call(__e, tmp99434, symlist, tmp99438)

																																					tmp99440 := Call(__e, tmp99433, tmp99439, Nil)

																																					tmp99441 := Call(__e, tmp99432, sym_h, tmp99440)

																																					tmp99442 := Call(__e, tmp99429, tmp99431, tmp99441)

																																					tmp99443 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																					tmp99444 := Call(__e, tmp99443, Hyp, V2791)

																																					tmp99445 := Call(__e, tmp99428, tmp99442, tmp99444)

																																					tmp99446 := Call(__e, tmp99417, tmp99427, tmp99445)

																																					__e.TailApply(tmp99416, V2790, tmp99446, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp99447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																				tmp99448 := Call(__e, tmp99447, V2526)

																																				tmp99449 := Call(__e, tmp99413, tmp99448)

																																				__e.TailApply(tmp99410, tmp99449)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp99490 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																	tmp99491 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																																	tmp99492 := Call(__e, tmp99491, V2535)

																																	tmp99493 := Call(__e, tmp99490, tmp99492, V2791)

																																	tmp99494 := Call(__e, tmp99405, tmp99493)

																																	__e.TailApply(tmp99402, tmp99494)
																																	return

																																}, 1)

																																tmp99495 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																																tmp99496 := Call(__e, tmp99495, V2791)

																																__e.TailApply(tmp99397, tmp99496)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp99699 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																													tmp99700 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																													tmp99701 := Call(__e, tmp99700, V2536)

																													tmp99702 := Call(__e, tmp99699, tmp99701, V2791)

																													tmp99703 := Call(__e, tmp99394, tmp99702)

																													__e.TailApply(tmp99391, tmp99703)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}

																										}, 1)

																										tmp100017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																										tmp100018 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																										tmp100019 := Call(__e, tmp100018, V2536)

																										tmp100020 := Call(__e, tmp100017, tmp100019, V2791)

																										__e.TailApply(tmp99386, tmp100020)
																										return

																									} else {
																										tmp99384 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																										tmp99385 := Call(__e, tmp99384, V2536)

																										if True == tmp99385 {
																											tmp99282 := MakeNative(func(__e *ControlFlow) {
																												A := __e.Get(1)
																												_ = A
																												tmp99283 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																												tmp99284 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																												tmp99285 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																												tmp99286 := Call(__e, tmp99285, A, Nil)

																												tmp99287 := Call(__e, tmp99284, symlist, tmp99286)

																												tmp99288 := Call(__e, tmp99283, V2536, tmp99287, V2791)

																												_ = tmp99288

																												tmp99289 := MakeNative(func(__e *ControlFlow) {
																													Result := __e.Get(1)
																													_ = Result
																													tmp99290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																													tmp99291 := Call(__e, tmp99290, V2536, V2791)

																													_ = tmp99291

																													__e.Return(Result)
																													return

																												}, 1)

																												tmp99292 := MakeNative(func(__e *ControlFlow) {
																													V2548 := __e.Get(1)
																													_ = V2548
																													tmp99375 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																													tmp99376 := Call(__e, tmp99375, Nil, V2548)

																													if True == tmp99376 {
																														tmp99339 := MakeNative(func(__e *ControlFlow) {
																															Hyp := __e.Get(1)
																															_ = Hyp
																															tmp99340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																															tmp99341 := Call(__e, tmp99340)

																															_ = tmp99341

																															tmp99342 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																															tmp99343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99346 := Call(__e, tmp99345, X, V2791)

																															tmp99347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99349 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99350 := Call(__e, tmp99349, A, V2791)

																															tmp99351 := Call(__e, tmp99348, tmp99350, Nil)

																															tmp99352 := Call(__e, tmp99347, sym_h, tmp99351)

																															tmp99353 := Call(__e, tmp99344, tmp99346, tmp99352)

																															tmp99354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99356 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99357 := Call(__e, tmp99356, Y, V2791)

																															tmp99358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99360 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99361 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																															tmp99362 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99363 := Call(__e, tmp99362, A, V2791)

																															tmp99364 := Call(__e, tmp99361, tmp99363, Nil)

																															tmp99365 := Call(__e, tmp99360, symlist, tmp99364)

																															tmp99366 := Call(__e, tmp99359, tmp99365, Nil)

																															tmp99367 := Call(__e, tmp99358, sym_h, tmp99366)

																															tmp99368 := Call(__e, tmp99355, tmp99357, tmp99367)

																															tmp99369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																															tmp99370 := Call(__e, tmp99369, Hyp, V2791)

																															tmp99371 := Call(__e, tmp99354, tmp99368, tmp99370)

																															tmp99372 := Call(__e, tmp99343, tmp99353, tmp99371)

																															__e.TailApply(tmp99342, V2790, tmp99372, V2791, V2792)
																															return

																														}, 1)

																														tmp99373 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																														tmp99374 := Call(__e, tmp99373, V2526)

																														__e.TailApply(tmp99339, tmp99374)
																														return

																													} else {
																														tmp99337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																														tmp99338 := Call(__e, tmp99337, V2548)

																														if True == tmp99338 {
																															tmp99295 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																															tmp99296 := Call(__e, tmp99295, V2548, Nil, V2791)

																															_ = tmp99296

																															tmp99297 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp99298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																																tmp99299 := Call(__e, tmp99298, V2548, V2791)

																																_ = tmp99299

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp99300 := MakeNative(func(__e *ControlFlow) {
																																Hyp := __e.Get(1)
																																_ = Hyp
																																tmp99301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																																tmp99302 := Call(__e, tmp99301)

																																_ = tmp99302

																																tmp99303 := Call(__e, PrimNS1Value(symns2_1value), symbind)

																																tmp99304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99305 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99307 := Call(__e, tmp99306, X, V2791)

																																tmp99308 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99309 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99311 := Call(__e, tmp99310, A, V2791)

																																tmp99312 := Call(__e, tmp99309, tmp99311, Nil)

																																tmp99313 := Call(__e, tmp99308, sym_h, tmp99312)

																																tmp99314 := Call(__e, tmp99305, tmp99307, tmp99313)

																																tmp99315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99317 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99318 := Call(__e, tmp99317, Y, V2791)

																																tmp99319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																																tmp99323 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99324 := Call(__e, tmp99323, A, V2791)

																																tmp99325 := Call(__e, tmp99322, tmp99324, Nil)

																																tmp99326 := Call(__e, tmp99321, symlist, tmp99325)

																																tmp99327 := Call(__e, tmp99320, tmp99326, Nil)

																																tmp99328 := Call(__e, tmp99319, sym_h, tmp99327)

																																tmp99329 := Call(__e, tmp99316, tmp99318, tmp99328)

																																tmp99330 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																																tmp99331 := Call(__e, tmp99330, Hyp, V2791)

																																tmp99332 := Call(__e, tmp99315, tmp99329, tmp99331)

																																tmp99333 := Call(__e, tmp99304, tmp99314, tmp99332)

																																__e.TailApply(tmp99303, V2790, tmp99333, V2791, V2792)
																																return

																															}, 1)

																															tmp99334 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																															tmp99335 := Call(__e, tmp99334, V2526)

																															tmp99336 := Call(__e, tmp99300, tmp99335)

																															__e.TailApply(tmp99297, tmp99336)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp99377 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																												tmp99378 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																												tmp99379 := Call(__e, tmp99378, V2535)

																												tmp99380 := Call(__e, tmp99377, tmp99379, V2791)

																												tmp99381 := Call(__e, tmp99292, tmp99380)

																												__e.TailApply(tmp99289, tmp99381)
																												return

																											}, 1)

																											tmp99382 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																											tmp99383 := Call(__e, tmp99382, V2791)

																											__e.TailApply(tmp99282, tmp99383)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}

																								}, 1)

																								tmp100023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																								tmp100024 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								tmp100025 := Call(__e, tmp100024, V2535)

																								tmp100026 := Call(__e, tmp100023, tmp100025, V2791)

																								__e.TailApply(tmp99279, tmp100026)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp100029 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp100030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp100031 := Call(__e, tmp100030, V2533)

																						tmp100032 := Call(__e, tmp100029, tmp100031, V2791)

																						__e.TailApply(tmp99277, tmp100032)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp100035 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp100036 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				tmp100037 := Call(__e, tmp100036, V2533)

																				tmp100038 := Call(__e, tmp100035, tmp100037, V2791)

																				__e.TailApply(tmp99275, tmp100038)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp100041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																		tmp100042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp100043 := Call(__e, tmp100042, V2527)

																		tmp100044 := Call(__e, tmp100041, tmp100043, V2791)

																		__e.TailApply(tmp99273, tmp100044)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp100047 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																tmp100048 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp100049 := Call(__e, tmp100048, V2531)

																tmp100050 := Call(__e, tmp100047, tmp100049, V2791)

																__e.TailApply(tmp99271, tmp100050)
																return

															}, 1)

															tmp100051 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp100052 := Call(__e, tmp100051, V2531)

															__e.TailApply(tmp99270, tmp100052)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp100055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp100056 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp100057 := Call(__e, tmp100056, V2530)

													tmp100058 := Call(__e, tmp100055, tmp100057, V2791)

													__e.TailApply(tmp99268, tmp100058)
													return

												}, 1)

												tmp100059 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp100060 := Call(__e, tmp100059, V2530)

												__e.TailApply(tmp99267, tmp100060)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp100063 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp100064 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp100065 := Call(__e, tmp100064, V2528)

										tmp100066 := Call(__e, tmp100063, tmp100065, V2791)

										__e.TailApply(tmp99265, tmp100066)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp100069 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp100070 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp100071 := Call(__e, tmp100070, V2528)

								tmp100072 := Call(__e, tmp100069, tmp100071, V2791)

								__e.TailApply(tmp99263, tmp100072)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp100075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp100076 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp100077 := Call(__e, tmp100076, V2527)

						tmp100078 := Call(__e, tmp100075, tmp100077, V2791)

						__e.TailApply(tmp99261, tmp100078)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp100081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100082 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100083 := Call(__e, tmp100082, V2526)

				tmp100084 := Call(__e, tmp100081, tmp100083, V2791)

				__e.TailApply(tmp99259, tmp100084)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp100087 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp100088 := Call(__e, tmp100087, V2789, V2791)

		tmp100089 := Call(__e, tmp99257, tmp100088)

		__e.TailApply(tmp97244, tmp100089)
		return

	}, 4)

	tmp100090 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1hyps, tmp97243)

	_ = tmp100090

	tmp100091 := MakeNative(func(__e *ControlFlow) {
		V2809 := __e.Get(1)
		_ = V2809
		V2810 := __e.Get(2)
		_ = V2810
		V2811 := __e.Get(3)
		_ = V2811
		V2812 := __e.Get(4)
		_ = V2812
		tmp100115 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp100116 := Call(__e, tmp100115, symshen_4_dspy_d)

		if True == tmp100116 {
			tmp100094 := Call(__e, PrimNS1Value(symns2_1value), symshen_4line)

			tmp100095 := Call(__e, tmp100094)

			_ = tmp100095

			tmp100096 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show_1p)

			tmp100097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp100098 := Call(__e, tmp100097, V2809, V2811)

			tmp100099 := Call(__e, tmp100096, tmp100098)

			_ = tmp100099

			tmp100100 := Call(__e, PrimNS1Value(symns2_1value), symnl)

			tmp100101 := Call(__e, tmp100100, MakeNumber(1))

			_ = tmp100101

			tmp100102 := Call(__e, PrimNS1Value(symns2_1value), symnl)

			tmp100103 := Call(__e, tmp100102, MakeNumber(1))

			_ = tmp100103

			tmp100104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show_1assumptions)

			tmp100105 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp100106 := Call(__e, tmp100105, V2810, V2811)

			tmp100107 := Call(__e, tmp100104, tmp100106, MakeNumber(1))

			_ = tmp100107

			tmp100108 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp100109 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp100110 := Call(__e, tmp100109)

			tmp100111 := Call(__e, tmp100108, MakeString("\n> "), tmp100110)

			_ = tmp100111

			tmp100112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pause_1for_1user)

			tmp100113 := Call(__e, tmp100112)

			_ = tmp100113

			tmp100114 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp100114, V2812)
			return

		} else {
			tmp100093 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

			__e.TailApply(tmp100093, V2812)
			return

		}

	}, 4)

	tmp100117 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show, tmp100091)

	_ = tmp100117

	tmp100118 := MakeNative(func(__e *ControlFlow) {
		tmp100119 := MakeNative(func(__e *ControlFlow) {
			Infs := __e.Get(1)
			_ = Infs
			tmp100120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp100121 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp100122 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp100123 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp100124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp100126 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100127 := Call(__e, tmp100126, MakeNumber(1), Infs)

			var ifres100125 Obj

			if True == tmp100127 {
				ifres100125 = MakeString("")

			} else {
				ifres100125 = MakeString("s")

			}

			tmp100128 := Call(__e, tmp100124, ifres100125, MakeString(" \n?- "), symshen_4a)

			tmp100129 := Call(__e, tmp100123, MakeString(" inference"), tmp100128)

			tmp100130 := Call(__e, tmp100122, Infs, tmp100129, symshen_4a)

			tmp100131 := Call(__e, tmp100121, MakeString("____________________________________________________________ "), tmp100130)

			tmp100132 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp100133 := Call(__e, tmp100132)

			__e.TailApply(tmp100120, tmp100131, tmp100133)
			return

		}, 1)

		tmp100134 := Call(__e, PrimNS1Value(symns2_1value), syminferences)

		tmp100135 := Call(__e, tmp100134)

		__e.TailApply(tmp100119, tmp100135)
		return

	}, 0)

	tmp100136 := Call(__e, PrimNS1Value(symns2_1set), symshen_4line, tmp100118)

	_ = tmp100136

	tmp100137 := MakeNative(func(__e *ControlFlow) {
		V2814 := __e.Get(1)
		_ = V2814
		tmp100193 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp100194 := Call(__e, tmp100193, V2814)

		var ifres100161 Obj

		if True == tmp100194 {
			tmp100189 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100190 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp100191 := Call(__e, tmp100190, V2814)

			tmp100192 := Call(__e, tmp100189, tmp100191)

			var ifres100163 Obj

			if True == tmp100192 {
				tmp100183 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp100184 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100185 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp100186 := Call(__e, tmp100185, V2814)

				tmp100187 := Call(__e, tmp100184, tmp100186)

				tmp100188 := Call(__e, tmp100183, sym_h, tmp100187)

				var ifres100165 Obj

				if True == tmp100188 {
					tmp100177 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp100178 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp100179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp100180 := Call(__e, tmp100179, V2814)

					tmp100181 := Call(__e, tmp100178, tmp100180)

					tmp100182 := Call(__e, tmp100177, tmp100181)

					var ifres100167 Obj

					if True == tmp100182 {
						tmp100169 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp100170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp100171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp100172 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp100173 := Call(__e, tmp100172, V2814)

						tmp100174 := Call(__e, tmp100171, tmp100173)

						tmp100175 := Call(__e, tmp100170, tmp100174)

						tmp100176 := Call(__e, tmp100169, Nil, tmp100175)

						var ifres100168 Obj

						if True == tmp100176 {
							ifres100168 = True

						} else {
							ifres100168 = False

						}

						ifres100167 = ifres100168

					} else {
						ifres100167 = False

					}

					var ifres100166 Obj

					if True == ifres100167 {
						ifres100166 = True

					} else {
						ifres100166 = False

					}

					ifres100165 = ifres100166

				} else {
					ifres100165 = False

				}

				var ifres100164 Obj

				if True == ifres100165 {
					ifres100164 = True

				} else {
					ifres100164 = False

				}

				ifres100163 = ifres100164

			} else {
				ifres100163 = False

			}

			var ifres100162 Obj

			if True == ifres100163 {
				ifres100162 = True

			} else {
				ifres100162 = False

			}

			ifres100161 = ifres100162

		} else {
			ifres100161 = False

		}

		if True == ifres100161 {
			tmp100144 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp100145 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp100146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp100147 := Call(__e, tmp100146, V2814)

			tmp100148 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp100149 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp100150 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp100151 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp100152 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp100153 := Call(__e, tmp100152, V2814)

			tmp100154 := Call(__e, tmp100151, tmp100153)

			tmp100155 := Call(__e, tmp100150, tmp100154)

			tmp100156 := Call(__e, tmp100149, tmp100155, MakeString(""), symshen_4r)

			tmp100157 := Call(__e, tmp100148, MakeString(" : "), tmp100156)

			tmp100158 := Call(__e, tmp100145, tmp100147, tmp100157, symshen_4r)

			tmp100159 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp100160 := Call(__e, tmp100159)

			__e.TailApply(tmp100144, tmp100158, tmp100160)
			return

		} else {
			tmp100139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp100140 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp100141 := Call(__e, tmp100140, V2814, MakeString(""), symshen_4r)

			tmp100142 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp100143 := Call(__e, tmp100142)

			__e.TailApply(tmp100139, tmp100141, tmp100143)
			return

		}

	}, 1)

	tmp100195 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show_1p, tmp100137)

	_ = tmp100195

	tmp100196 := MakeNative(func(__e *ControlFlow) {
		V2819 := __e.Get(1)
		_ = V2819
		V2820 := __e.Get(2)
		_ = V2820
		tmp100219 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp100220 := Call(__e, tmp100219, Nil, V2819)

		if True == tmp100220 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp100217 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100218 := Call(__e, tmp100217, V2819)

			if True == tmp100218 {
				tmp100200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp100201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp100202 := Call(__e, tmp100201, V2820, MakeString(". "), symshen_4a)

				tmp100203 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp100204 := Call(__e, tmp100203)

				tmp100205 := Call(__e, tmp100200, tmp100202, tmp100204)

				_ = tmp100205

				tmp100206 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show_1p)

				tmp100207 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100208 := Call(__e, tmp100207, V2819)

				tmp100209 := Call(__e, tmp100206, tmp100208)

				_ = tmp100209

				tmp100210 := Call(__e, PrimNS1Value(symns2_1value), symnl)

				tmp100211 := Call(__e, tmp100210, MakeNumber(1))

				_ = tmp100211

				tmp100212 := Call(__e, PrimNS1Value(symns2_1value), symshen_4show_1assumptions)

				tmp100213 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp100214 := Call(__e, tmp100213, V2819)

				tmp100215 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp100216 := Call(__e, tmp100215, V2820, MakeNumber(1))

				__e.TailApply(tmp100212, tmp100214, tmp100216)
				return

			} else {
				tmp100199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp100199, symshen_4show_1assumptions)
				return

			}

		}

	}, 2)

	tmp100221 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show_1assumptions, tmp100196)

	_ = tmp100221

	tmp100222 := MakeNative(func(__e *ControlFlow) {
		tmp100223 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			tmp100227 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100228 := Call(__e, tmp100227, Byte, MakeNumber(94))

			if True == tmp100228 {
				tmp100226 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp100226, MakeString("input aborted\n"))
				return

			} else {
				tmp100225 := Call(__e, PrimNS1Value(symns2_1value), symnl)

				__e.TailApply(tmp100225, MakeNumber(1))
				return

			}

		}, 1)

		tmp100229 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

		tmp100230 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

		tmp100231 := Call(__e, tmp100230)

		tmp100232 := Call(__e, tmp100229, tmp100231)

		__e.TailApply(tmp100223, tmp100232)
		return

	}, 0)

	tmp100233 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pause_1for_1user, tmp100222)

	_ = tmp100233

	tmp100234 := MakeNative(func(__e *ControlFlow) {
		V2822 := __e.Get(1)
		_ = V2822
		tmp100235 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp100236 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

		tmp100237 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp100238 := Call(__e, tmp100237, symshen_4_dsignedfuncs_d)

		tmp100239 := Call(__e, tmp100236, V2822, tmp100238)

		__e.TailApply(tmp100235, tmp100239)
		return

	}, 1)

	tmp100240 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typedf_2, tmp100234)

	_ = tmp100240

	tmp100241 := MakeNative(func(__e *ControlFlow) {
		V2824 := __e.Get(1)
		_ = V2824
		tmp100242 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		__e.TailApply(tmp100242, symshen_4type_1signature_1of_1, V2824)
		return

	}, 1)

	tmp100243 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sigf, tmp100241)

	_ = tmp100243

	tmp100244 := MakeNative(func(__e *ControlFlow) {
		tmp100245 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		__e.TailApply(tmp100245, sym_e_e)
		return

	}, 0)

	tmp100246 := Call(__e, PrimNS1Value(symns2_1set), symshen_4placeholder, tmp100244)

	_ = tmp100246

	tmp100247 := MakeNative(func(__e *ControlFlow) {
		V2829 := __e.Get(1)
		_ = V2829
		V2830 := __e.Get(2)
		_ = V2830
		V2831 := __e.Get(3)
		_ = V2831
		V2832 := __e.Get(4)
		_ = V2832
		tmp100248 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp100518 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100519 := Call(__e, tmp100518, Case, False)

			if True == tmp100519 {
				tmp100250 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp100486 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100487 := Call(__e, tmp100486, Case, False)

					if True == tmp100487 {
						tmp100252 := MakeNative(func(__e *ControlFlow) {
							Case := __e.Get(1)
							_ = Case
							tmp100454 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp100455 := Call(__e, tmp100454, Case, False)

							if True == tmp100455 {
								tmp100254 := MakeNative(func(__e *ControlFlow) {
									Case := __e.Get(1)
									_ = Case
									tmp100406 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp100407 := Call(__e, tmp100406, Case, False)

									if True == tmp100407 {
										tmp100256 := MakeNative(func(__e *ControlFlow) {
											V2517 := __e.Get(1)
											_ = V2517
											tmp100402 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp100403 := Call(__e, tmp100402, Nil, V2517)

											if True == tmp100403 {
												tmp100258 := MakeNative(func(__e *ControlFlow) {
													V2518 := __e.Get(1)
													_ = V2518
													tmp100398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp100399 := Call(__e, tmp100398, V2518)

													if True == tmp100399 {
														tmp100279 := MakeNative(func(__e *ControlFlow) {
															V2519 := __e.Get(1)
															_ = V2519
															tmp100392 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp100393 := Call(__e, tmp100392, symlist, V2519)

															if True == tmp100393 {
																tmp100341 := MakeNative(func(__e *ControlFlow) {
																	V2520 := __e.Get(1)
																	_ = V2520
																	tmp100386 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																	tmp100387 := Call(__e, tmp100386, V2520)

																	if True == tmp100387 {
																		tmp100360 := MakeNative(func(__e *ControlFlow) {
																			A := __e.Get(1)
																			_ = A
																			tmp100361 := MakeNative(func(__e *ControlFlow) {
																				V2521 := __e.Get(1)
																				_ = V2521
																				tmp100378 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																				tmp100379 := Call(__e, tmp100378, Nil, V2521)

																				if True == tmp100379 {
																					tmp100375 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																					tmp100376 := Call(__e, tmp100375)

																					_ = tmp100376

																					tmp100377 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																					__e.TailApply(tmp100377, V2832)
																					return

																				} else {
																					tmp100373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																					tmp100374 := Call(__e, tmp100373, V2521)

																					if True == tmp100374 {
																						tmp100364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																						tmp100365 := Call(__e, tmp100364, V2521, Nil, V2831)

																						_ = tmp100365

																						tmp100366 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp100367 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																							tmp100368 := Call(__e, tmp100367, V2521, V2831)

																							_ = tmp100368

																							__e.Return(Result)
																							return

																						}, 1)

																						tmp100369 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																						tmp100370 := Call(__e, tmp100369)

																						_ = tmp100370

																						tmp100371 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																						tmp100372 := Call(__e, tmp100371, V2832)

																						__e.TailApply(tmp100366, tmp100372)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp100380 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																			tmp100381 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																			tmp100382 := Call(__e, tmp100381, V2520)

																			tmp100383 := Call(__e, tmp100380, tmp100382, V2831)

																			__e.TailApply(tmp100361, tmp100383)
																			return

																		}, 1)

																		tmp100384 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																		tmp100385 := Call(__e, tmp100384, V2520)

																		__e.TailApply(tmp100360, tmp100385)
																		return

																	} else {
																		tmp100358 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																		tmp100359 := Call(__e, tmp100358, V2520)

																		if True == tmp100359 {
																			tmp100344 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp100345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																				tmp100346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				tmp100347 := Call(__e, tmp100346, A, Nil)

																				tmp100348 := Call(__e, tmp100345, V2520, tmp100347, V2831)

																				_ = tmp100348

																				tmp100349 := MakeNative(func(__e *ControlFlow) {
																					Result := __e.Get(1)
																					_ = Result
																					tmp100350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																					tmp100351 := Call(__e, tmp100350, V2520, V2831)

																					_ = tmp100351

																					__e.Return(Result)
																					return

																				}, 1)

																				tmp100352 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																				tmp100353 := Call(__e, tmp100352)

																				_ = tmp100353

																				tmp100354 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																				tmp100355 := Call(__e, tmp100354, V2832)

																				__e.TailApply(tmp100349, tmp100355)
																				return

																			}, 1)

																			tmp100356 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																			tmp100357 := Call(__e, tmp100356, V2831)

																			__e.TailApply(tmp100344, tmp100357)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}

																}, 1)

																tmp100388 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																tmp100389 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																tmp100390 := Call(__e, tmp100389, V2518)

																tmp100391 := Call(__e, tmp100388, tmp100390, V2831)

																__e.TailApply(tmp100341, tmp100391)
																return

															} else {
																tmp100339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																tmp100340 := Call(__e, tmp100339, V2519)

																if True == tmp100340 {
																	tmp100282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																	tmp100283 := Call(__e, tmp100282, V2519, symlist, V2831)

																	_ = tmp100283

																	tmp100284 := MakeNative(func(__e *ControlFlow) {
																		Result := __e.Get(1)
																		_ = Result
																		tmp100285 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																		tmp100286 := Call(__e, tmp100285, V2519, V2831)

																		_ = tmp100286

																		__e.Return(Result)
																		return

																	}, 1)

																	tmp100287 := MakeNative(func(__e *ControlFlow) {
																		V2522 := __e.Get(1)
																		_ = V2522
																		tmp100332 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp100333 := Call(__e, tmp100332, V2522)

																		if True == tmp100333 {
																			tmp100306 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp100307 := MakeNative(func(__e *ControlFlow) {
																					V2523 := __e.Get(1)
																					_ = V2523
																					tmp100324 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp100325 := Call(__e, tmp100324, Nil, V2523)

																					if True == tmp100325 {
																						tmp100321 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																						tmp100322 := Call(__e, tmp100321)

																						_ = tmp100322

																						tmp100323 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																						__e.TailApply(tmp100323, V2832)
																						return

																					} else {
																						tmp100319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																						tmp100320 := Call(__e, tmp100319, V2523)

																						if True == tmp100320 {
																							tmp100310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																							tmp100311 := Call(__e, tmp100310, V2523, Nil, V2831)

																							_ = tmp100311

																							tmp100312 := MakeNative(func(__e *ControlFlow) {
																								Result := __e.Get(1)
																								_ = Result
																								tmp100313 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																								tmp100314 := Call(__e, tmp100313, V2523, V2831)

																								_ = tmp100314

																								__e.Return(Result)
																								return

																							}, 1)

																							tmp100315 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																							tmp100316 := Call(__e, tmp100315)

																							_ = tmp100316

																							tmp100317 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																							tmp100318 := Call(__e, tmp100317, V2832)

																							__e.TailApply(tmp100312, tmp100318)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}

																				}, 1)

																				tmp100326 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp100327 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp100328 := Call(__e, tmp100327, V2522)

																				tmp100329 := Call(__e, tmp100326, tmp100328, V2831)

																				__e.TailApply(tmp100307, tmp100329)
																				return

																			}, 1)

																			tmp100330 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp100331 := Call(__e, tmp100330, V2522)

																			__e.TailApply(tmp100306, tmp100331)
																			return

																		} else {
																			tmp100304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																			tmp100305 := Call(__e, tmp100304, V2522)

																			if True == tmp100305 {
																				tmp100290 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp100291 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																					tmp100292 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																					tmp100293 := Call(__e, tmp100292, A, Nil)

																					tmp100294 := Call(__e, tmp100291, V2522, tmp100293, V2831)

																					_ = tmp100294

																					tmp100295 := MakeNative(func(__e *ControlFlow) {
																						Result := __e.Get(1)
																						_ = Result
																						tmp100296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																						tmp100297 := Call(__e, tmp100296, V2522, V2831)

																						_ = tmp100297

																						__e.Return(Result)
																						return

																					}, 1)

																					tmp100298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																					tmp100299 := Call(__e, tmp100298)

																					_ = tmp100299

																					tmp100300 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																					tmp100301 := Call(__e, tmp100300, V2832)

																					__e.TailApply(tmp100295, tmp100301)
																					return

																				}, 1)

																				tmp100302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																				tmp100303 := Call(__e, tmp100302, V2831)

																				__e.TailApply(tmp100290, tmp100303)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp100334 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp100335 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp100336 := Call(__e, tmp100335, V2518)

																	tmp100337 := Call(__e, tmp100334, tmp100336, V2831)

																	tmp100338 := Call(__e, tmp100287, tmp100337)

																	__e.TailApply(tmp100284, tmp100338)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}

														}, 1)

														tmp100394 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp100395 := Call(__e, PrimNS1Value(symns2_1value), symhd)

														tmp100396 := Call(__e, tmp100395, V2518)

														tmp100397 := Call(__e, tmp100394, tmp100396, V2831)

														__e.TailApply(tmp100279, tmp100397)
														return

													} else {
														tmp100277 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

														tmp100278 := Call(__e, tmp100277, V2518)

														if True == tmp100278 {
															tmp100261 := MakeNative(func(__e *ControlFlow) {
																A := __e.Get(1)
																_ = A
																tmp100262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																tmp100263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp100264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp100265 := Call(__e, tmp100264, A, Nil)

																tmp100266 := Call(__e, tmp100263, symlist, tmp100265)

																tmp100267 := Call(__e, tmp100262, V2518, tmp100266, V2831)

																_ = tmp100267

																tmp100268 := MakeNative(func(__e *ControlFlow) {
																	Result := __e.Get(1)
																	_ = Result
																	tmp100269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																	tmp100270 := Call(__e, tmp100269, V2518, V2831)

																	_ = tmp100270

																	__e.Return(Result)
																	return

																}, 1)

																tmp100271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																tmp100272 := Call(__e, tmp100271)

																_ = tmp100272

																tmp100273 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

																tmp100274 := Call(__e, tmp100273, V2832)

																__e.TailApply(tmp100268, tmp100274)
																return

															}, 1)

															tmp100275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

															tmp100276 := Call(__e, tmp100275, V2831)

															__e.TailApply(tmp100261, tmp100276)
															return

														} else {
															__e.Return(False)
															return
														}

													}

												}, 1)

												tmp100400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp100401 := Call(__e, tmp100400, V2830, V2831)

												__e.TailApply(tmp100258, tmp100401)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp100404 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp100405 := Call(__e, tmp100404, V2829, V2831)

										__e.TailApply(tmp100256, tmp100405)
										return

									} else {
										__e.Return(Case)
										return
									}

								}, 1)

								tmp100408 := MakeNative(func(__e *ControlFlow) {
									V2516 := __e.Get(1)
									_ = V2516
									tmp100449 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp100450 := Call(__e, tmp100449, symsymbol, V2516)

									if True == tmp100450 {
										tmp100434 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

										tmp100435 := Call(__e, tmp100434)

										_ = tmp100435

										tmp100436 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

										tmp100437 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

										tmp100438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp100439 := Call(__e, tmp100438, V2829, V2831)

										tmp100440 := Call(__e, tmp100437, tmp100439)

										tmp100441 := MakeNative(func(__e *ControlFlow) {
											tmp100442 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

											tmp100443 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp100444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_2)

											tmp100445 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp100446 := Call(__e, tmp100445, V2829, V2831)

											tmp100447 := Call(__e, tmp100444, tmp100446)

											tmp100448 := Call(__e, tmp100443, tmp100447)

											__e.TailApply(tmp100442, tmp100448, V2831, V2832)
											return

										}, 0)

										__e.TailApply(tmp100436, tmp100440, V2831, tmp100441)
										return

									} else {
										tmp100432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

										tmp100433 := Call(__e, tmp100432, V2516)

										if True == tmp100433 {
											tmp100411 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

											tmp100412 := Call(__e, tmp100411, V2516, symsymbol, V2831)

											_ = tmp100412

											tmp100413 := MakeNative(func(__e *ControlFlow) {
												Result := __e.Get(1)
												_ = Result
												tmp100414 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

												tmp100415 := Call(__e, tmp100414, V2516, V2831)

												_ = tmp100415

												__e.Return(Result)
												return

											}, 1)

											tmp100416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

											tmp100417 := Call(__e, tmp100416)

											_ = tmp100417

											tmp100418 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

											tmp100419 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

											tmp100420 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp100421 := Call(__e, tmp100420, V2829, V2831)

											tmp100422 := Call(__e, tmp100419, tmp100421)

											tmp100423 := MakeNative(func(__e *ControlFlow) {
												tmp100424 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

												tmp100425 := Call(__e, PrimNS1Value(symns2_1value), symnot)

												tmp100426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_2)

												tmp100427 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

												tmp100428 := Call(__e, tmp100427, V2829, V2831)

												tmp100429 := Call(__e, tmp100426, tmp100428)

												tmp100430 := Call(__e, tmp100425, tmp100429)

												__e.TailApply(tmp100424, tmp100430, V2831, V2832)
												return

											}, 0)

											tmp100431 := Call(__e, tmp100418, tmp100422, V2831, tmp100423)

											__e.TailApply(tmp100413, tmp100431)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp100451 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp100452 := Call(__e, tmp100451, V2830, V2831)

								tmp100453 := Call(__e, tmp100408, tmp100452)

								__e.TailApply(tmp100254, tmp100453)
								return

							} else {
								__e.Return(Case)
								return
							}

						}, 1)

						tmp100456 := MakeNative(func(__e *ControlFlow) {
							V2515 := __e.Get(1)
							_ = V2515
							tmp100481 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp100482 := Call(__e, tmp100481, symstring, V2515)

							if True == tmp100482 {
								tmp100474 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

								tmp100475 := Call(__e, tmp100474)

								_ = tmp100475

								tmp100476 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

								tmp100477 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

								tmp100478 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp100479 := Call(__e, tmp100478, V2829, V2831)

								tmp100480 := Call(__e, tmp100477, tmp100479)

								__e.TailApply(tmp100476, tmp100480, V2831, V2832)
								return

							} else {
								tmp100472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

								tmp100473 := Call(__e, tmp100472, V2515)

								if True == tmp100473 {
									tmp100459 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

									tmp100460 := Call(__e, tmp100459, V2515, symstring, V2831)

									_ = tmp100460

									tmp100461 := MakeNative(func(__e *ControlFlow) {
										Result := __e.Get(1)
										_ = Result
										tmp100462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

										tmp100463 := Call(__e, tmp100462, V2515, V2831)

										_ = tmp100463

										__e.Return(Result)
										return

									}, 1)

									tmp100464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

									tmp100465 := Call(__e, tmp100464)

									_ = tmp100465

									tmp100466 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

									tmp100467 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

									tmp100468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

									tmp100469 := Call(__e, tmp100468, V2829, V2831)

									tmp100470 := Call(__e, tmp100467, tmp100469)

									tmp100471 := Call(__e, tmp100466, tmp100470, V2831, V2832)

									__e.TailApply(tmp100461, tmp100471)
									return

								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp100483 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp100484 := Call(__e, tmp100483, V2830, V2831)

						tmp100485 := Call(__e, tmp100456, tmp100484)

						__e.TailApply(tmp100252, tmp100485)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp100488 := MakeNative(func(__e *ControlFlow) {
					V2514 := __e.Get(1)
					_ = V2514
					tmp100513 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100514 := Call(__e, tmp100513, symboolean, V2514)

					if True == tmp100514 {
						tmp100506 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

						tmp100507 := Call(__e, tmp100506)

						_ = tmp100507

						tmp100508 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

						tmp100509 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

						tmp100510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp100511 := Call(__e, tmp100510, V2829, V2831)

						tmp100512 := Call(__e, tmp100509, tmp100511)

						__e.TailApply(tmp100508, tmp100512, V2831, V2832)
						return

					} else {
						tmp100504 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

						tmp100505 := Call(__e, tmp100504, V2514)

						if True == tmp100505 {
							tmp100491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

							tmp100492 := Call(__e, tmp100491, V2514, symboolean, V2831)

							_ = tmp100492

							tmp100493 := MakeNative(func(__e *ControlFlow) {
								Result := __e.Get(1)
								_ = Result
								tmp100494 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

								tmp100495 := Call(__e, tmp100494, V2514, V2831)

								_ = tmp100495

								__e.Return(Result)
								return

							}, 1)

							tmp100496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

							tmp100497 := Call(__e, tmp100496)

							_ = tmp100497

							tmp100498 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

							tmp100499 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

							tmp100500 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp100501 := Call(__e, tmp100500, V2829, V2831)

							tmp100502 := Call(__e, tmp100499, tmp100501)

							tmp100503 := Call(__e, tmp100498, tmp100502, V2831, V2832)

							__e.TailApply(tmp100493, tmp100503)
							return

						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp100515 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100516 := Call(__e, tmp100515, V2830, V2831)

				tmp100517 := Call(__e, tmp100488, tmp100516)

				__e.TailApply(tmp100250, tmp100517)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp100520 := MakeNative(func(__e *ControlFlow) {
			V2513 := __e.Get(1)
			_ = V2513
			tmp100545 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100546 := Call(__e, tmp100545, symnumber, V2513)

			if True == tmp100546 {
				tmp100538 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp100539 := Call(__e, tmp100538)

				_ = tmp100539

				tmp100540 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

				tmp100541 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

				tmp100542 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100543 := Call(__e, tmp100542, V2829, V2831)

				tmp100544 := Call(__e, tmp100541, tmp100543)

				__e.TailApply(tmp100540, tmp100544, V2831, V2832)
				return

			} else {
				tmp100536 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

				tmp100537 := Call(__e, tmp100536, V2513)

				if True == tmp100537 {
					tmp100523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

					tmp100524 := Call(__e, tmp100523, V2513, symnumber, V2831)

					_ = tmp100524

					tmp100525 := MakeNative(func(__e *ControlFlow) {
						Result := __e.Get(1)
						_ = Result
						tmp100526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

						tmp100527 := Call(__e, tmp100526, V2513, V2831)

						_ = tmp100527

						__e.Return(Result)
						return

					}, 1)

					tmp100528 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

					tmp100529 := Call(__e, tmp100528)

					_ = tmp100529

					tmp100530 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

					tmp100531 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

					tmp100532 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp100533 := Call(__e, tmp100532, V2829, V2831)

					tmp100534 := Call(__e, tmp100531, tmp100533)

					tmp100535 := Call(__e, tmp100530, tmp100534, V2831, V2832)

					__e.TailApply(tmp100525, tmp100535)
					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)

		tmp100547 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp100548 := Call(__e, tmp100547, V2830, V2831)

		tmp100549 := Call(__e, tmp100520, tmp100548)

		__e.TailApply(tmp100248, tmp100549)
		return

	}, 4)

	tmp100550 := Call(__e, PrimNS1Value(symns2_1set), symshen_4base, tmp100247)

	_ = tmp100550

	tmp100551 := MakeNative(func(__e *ControlFlow) {
		V2838 := __e.Get(1)
		_ = V2838
		V2839 := __e.Get(2)
		_ = V2839
		V2840 := __e.Get(3)
		_ = V2840
		V2841 := __e.Get(4)
		_ = V2841
		V2842 := __e.Get(5)
		_ = V2842
		tmp100552 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp100566 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100567 := Call(__e, tmp100566, Case, False)

			if True == tmp100567 {
				tmp100554 := MakeNative(func(__e *ControlFlow) {
					V2510 := __e.Get(1)
					_ = V2510
					tmp100562 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp100563 := Call(__e, tmp100562, V2510)

					if True == tmp100563 {
						tmp100556 := MakeNative(func(__e *ControlFlow) {
							Hyp := __e.Get(1)
							_ = Hyp
							tmp100557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

							tmp100558 := Call(__e, tmp100557)

							_ = tmp100558

							tmp100559 := Call(__e, PrimNS1Value(symns2_1value), symshen_4by__hypothesis)

							__e.TailApply(tmp100559, V2838, V2839, Hyp, V2841, V2842)
							return

						}, 1)

						tmp100560 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp100561 := Call(__e, tmp100560, V2510)

						__e.TailApply(tmp100556, tmp100561)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp100564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100565 := Call(__e, tmp100564, V2840, V2841)

				__e.TailApply(tmp100554, tmp100565)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp100568 := MakeNative(func(__e *ControlFlow) {
			V2504 := __e.Get(1)
			_ = V2504
			tmp100621 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100622 := Call(__e, tmp100621, V2504)

			if True == tmp100622 {
				tmp100570 := MakeNative(func(__e *ControlFlow) {
					V2505 := __e.Get(1)
					_ = V2505
					tmp100615 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp100616 := Call(__e, tmp100615, V2505)

					if True == tmp100616 {
						tmp100572 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							tmp100573 := MakeNative(func(__e *ControlFlow) {
								V2506 := __e.Get(1)
								_ = V2506
								tmp100607 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp100608 := Call(__e, tmp100607, V2506)

								if True == tmp100608 {
									tmp100575 := MakeNative(func(__e *ControlFlow) {
										V2507 := __e.Get(1)
										_ = V2507
										tmp100601 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp100602 := Call(__e, tmp100601, sym_h, V2507)

										if True == tmp100602 {
											tmp100577 := MakeNative(func(__e *ControlFlow) {
												V2508 := __e.Get(1)
												_ = V2508
												tmp100595 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp100596 := Call(__e, tmp100595, V2508)

												if True == tmp100596 {
													tmp100579 := MakeNative(func(__e *ControlFlow) {
														B := __e.Get(1)
														_ = B
														tmp100580 := MakeNative(func(__e *ControlFlow) {
															V2509 := __e.Get(1)
															_ = V2509
															tmp100587 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp100588 := Call(__e, tmp100587, Nil, V2509)

															if True == tmp100588 {
																tmp100582 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																tmp100583 := Call(__e, tmp100582)

																_ = tmp100583

																tmp100584 := Call(__e, PrimNS1Value(symns2_1value), symidentical)

																tmp100585 := MakeNative(func(__e *ControlFlow) {
																	tmp100586 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																	__e.TailApply(tmp100586, V2839, B, V2841, V2842)
																	return

																}, 0)

																__e.TailApply(tmp100584, V2838, Y, V2841, tmp100585)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp100589 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp100590 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp100591 := Call(__e, tmp100590, V2508)

														tmp100592 := Call(__e, tmp100589, tmp100591, V2841)

														__e.TailApply(tmp100580, tmp100592)
														return

													}, 1)

													tmp100593 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp100594 := Call(__e, tmp100593, V2508)

													__e.TailApply(tmp100579, tmp100594)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp100597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp100598 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp100599 := Call(__e, tmp100598, V2506)

											tmp100600 := Call(__e, tmp100597, tmp100599, V2841)

											__e.TailApply(tmp100577, tmp100600)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp100603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

									tmp100604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp100605 := Call(__e, tmp100604, V2506)

									tmp100606 := Call(__e, tmp100603, tmp100605, V2841)

									__e.TailApply(tmp100575, tmp100606)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp100609 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp100610 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp100611 := Call(__e, tmp100610, V2505)

							tmp100612 := Call(__e, tmp100609, tmp100611, V2841)

							__e.TailApply(tmp100573, tmp100612)
							return

						}, 1)

						tmp100613 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp100614 := Call(__e, tmp100613, V2505)

						__e.TailApply(tmp100572, tmp100614)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp100617 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100618 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100619 := Call(__e, tmp100618, V2504)

				tmp100620 := Call(__e, tmp100617, tmp100619, V2841)

				__e.TailApply(tmp100570, tmp100620)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp100623 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp100624 := Call(__e, tmp100623, V2840, V2841)

		tmp100625 := Call(__e, tmp100568, tmp100624)

		__e.TailApply(tmp100552, tmp100625)
		return

	}, 5)

	tmp100626 := Call(__e, PrimNS1Value(symns2_1set), symshen_4by__hypothesis, tmp100551)

	_ = tmp100626

	tmp100627 := MakeNative(func(__e *ControlFlow) {
		V2848 := __e.Get(1)
		_ = V2848
		V2849 := __e.Get(2)
		_ = V2849
		V2850 := __e.Get(3)
		_ = V2850
		V2851 := __e.Get(4)
		_ = V2851
		V2852 := __e.Get(5)
		_ = V2852
		tmp100628 := MakeNative(func(__e *ControlFlow) {
			V2498 := __e.Get(1)
			_ = V2498
			tmp100675 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100676 := Call(__e, tmp100675, V2498)

			if True == tmp100676 {
				tmp100630 := MakeNative(func(__e *ControlFlow) {
					V2499 := __e.Get(1)
					_ = V2499
					tmp100669 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100670 := Call(__e, tmp100669, symdefine, V2499)

					if True == tmp100670 {
						tmp100632 := MakeNative(func(__e *ControlFlow) {
							V2500 := __e.Get(1)
							_ = V2500
							tmp100663 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp100664 := Call(__e, tmp100663, V2500)

							if True == tmp100664 {
								tmp100634 := MakeNative(func(__e *ControlFlow) {
									F := __e.Get(1)
									_ = F
									tmp100635 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp100636 := MakeNative(func(__e *ControlFlow) {
											Y := __e.Get(1)
											_ = Y
											tmp100637 := MakeNative(func(__e *ControlFlow) {
												E := __e.Get(1)
												_ = E
												tmp100638 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

												tmp100639 := Call(__e, tmp100638)

												_ = tmp100639

												tmp100640 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1defh)

												tmp100641 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

												tmp100642 := MakeNative(func(__e *ControlFlow) {
													Y := __e.Get(1)
													_ = Y
													tmp100643 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5sig_7rules_6)

													__e.TailApply(tmp100643, Y)
													return

												}, 1)

												tmp100644 := MakeNative(func(__e *ControlFlow) {
													E := __e.Get(1)
													_ = E
													tmp100652 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

													tmp100653 := Call(__e, tmp100652, E)

													if True == tmp100653 {
														tmp100647 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

														tmp100648 := Call(__e, PrimNS1Value(symns2_1value), symcn)

														tmp100649 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

														tmp100650 := Call(__e, tmp100649, E, MakeString("\n"), symshen_4s)

														tmp100651 := Call(__e, tmp100648, MakeString("parse error here: "), tmp100650)

														__e.TailApply(tmp100647, tmp100651)
														return

													} else {
														tmp100646 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

														__e.TailApply(tmp100646, MakeString("parse error\n"))
														return

													}

												}, 1)

												tmp100654 := Call(__e, tmp100641, tmp100642, X, tmp100644)

												__e.TailApply(tmp100640, tmp100654, F, V2849, V2850, V2851, V2852)
												return

											}, 1)

											tmp100655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

											tmp100656 := Call(__e, tmp100655, V2851)

											__e.TailApply(tmp100637, tmp100656)
											return

										}, 1)

										tmp100657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

										tmp100658 := Call(__e, tmp100657, V2851)

										__e.TailApply(tmp100636, tmp100658)
										return

									}, 1)

									tmp100659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp100660 := Call(__e, tmp100659, V2500)

									__e.TailApply(tmp100635, tmp100660)
									return

								}, 1)

								tmp100661 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp100662 := Call(__e, tmp100661, V2500)

								__e.TailApply(tmp100634, tmp100662)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp100665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp100666 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp100667 := Call(__e, tmp100666, V2498)

						tmp100668 := Call(__e, tmp100665, tmp100667, V2851)

						__e.TailApply(tmp100632, tmp100668)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp100671 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100672 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100673 := Call(__e, tmp100672, V2498)

				tmp100674 := Call(__e, tmp100671, tmp100673, V2851)

				__e.TailApply(tmp100630, tmp100674)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp100677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp100678 := Call(__e, tmp100677, V2848, V2851)

		__e.TailApply(tmp100628, tmp100678)
		return

	}, 5)

	tmp100679 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1def, tmp100627)

	_ = tmp100679

	tmp100680 := MakeNative(func(__e *ControlFlow) {
		V2859 := __e.Get(1)
		_ = V2859
		V2860 := __e.Get(2)
		_ = V2860
		V2861 := __e.Get(3)
		_ = V2861
		V2862 := __e.Get(4)
		_ = V2862
		V2863 := __e.Get(5)
		_ = V2863
		V2864 := __e.Get(6)
		_ = V2864
		tmp100681 := MakeNative(func(__e *ControlFlow) {
			V2494 := __e.Get(1)
			_ = V2494
			tmp100694 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100695 := Call(__e, tmp100694, V2494)

			if True == tmp100695 {
				tmp100683 := MakeNative(func(__e *ControlFlow) {
					Sig := __e.Get(1)
					_ = Sig
					tmp100684 := MakeNative(func(__e *ControlFlow) {
						Rules := __e.Get(1)
						_ = Rules
						tmp100685 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

						tmp100686 := Call(__e, tmp100685)

						_ = tmp100686

						tmp100687 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1defhh)

						tmp100688 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_1sig)

						tmp100689 := Call(__e, tmp100688, Sig)

						__e.TailApply(tmp100687, Sig, tmp100689, V2860, V2861, V2862, Rules, V2863, V2864)
						return

					}, 1)

					tmp100690 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp100691 := Call(__e, tmp100690, V2494)

					__e.TailApply(tmp100684, tmp100691)
					return

				}, 1)

				tmp100692 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100693 := Call(__e, tmp100692, V2494)

				__e.TailApply(tmp100683, tmp100693)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp100696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp100697 := Call(__e, tmp100696, V2859, V2863)

		__e.TailApply(tmp100681, tmp100697)
		return

	}, 6)

	tmp100698 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1defh, tmp100680)

	_ = tmp100698

	tmp100699 := MakeNative(func(__e *ControlFlow) {
		V2873 := __e.Get(1)
		_ = V2873
		V2874 := __e.Get(2)
		_ = V2874
		V2875 := __e.Get(3)
		_ = V2875
		V2876 := __e.Get(4)
		_ = V2876
		V2877 := __e.Get(5)
		_ = V2877
		V2878 := __e.Get(6)
		_ = V2878
		V2879 := __e.Get(7)
		_ = V2879
		V2880 := __e.Get(8)
		_ = V2880
		tmp100700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp100701 := Call(__e, tmp100700)

		_ = tmp100701

		tmp100702 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1rules)

		tmp100703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp100704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp100705 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp100706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp100707 := Call(__e, tmp100706, V2874, Nil)

		tmp100708 := Call(__e, tmp100705, sym_h, tmp100707)

		tmp100709 := Call(__e, tmp100704, V2875, tmp100708)

		tmp100710 := Call(__e, tmp100703, tmp100709, V2877)

		tmp100711 := MakeNative(func(__e *ControlFlow) {
			tmp100712 := Call(__e, PrimNS1Value(symns2_1value), symshen_4memo)

			__e.TailApply(tmp100712, V2875, V2873, V2876, V2879, V2880)
			return

		}, 0)

		__e.TailApply(tmp100702, V2878, V2874, MakeNumber(1), V2875, tmp100710, V2879, tmp100711)
		return

	}, 8)

	tmp100713 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1defhh, tmp100699)

	_ = tmp100713

	tmp100714 := MakeNative(func(__e *ControlFlow) {
		V2886 := __e.Get(1)
		_ = V2886
		V2887 := __e.Get(2)
		_ = V2887
		V2888 := __e.Get(3)
		_ = V2888
		V2889 := __e.Get(4)
		_ = V2889
		V2890 := __e.Get(5)
		_ = V2890
		tmp100715 := MakeNative(func(__e *ControlFlow) {
			Jnk := __e.Get(1)
			_ = Jnk
			tmp100716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp100717 := Call(__e, tmp100716)

			_ = tmp100717

			tmp100718 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

			tmp100719 := MakeNative(func(__e *ControlFlow) {
				tmp100720 := Call(__e, PrimNS1Value(symns2_1value), symbind)

				tmp100721 := Call(__e, PrimNS1Value(symns2_1value), symdeclare)

				tmp100722 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100723 := Call(__e, tmp100722, V2886, V2889)

				tmp100724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp100725 := Call(__e, tmp100724, V2888, V2889)

				tmp100726 := Call(__e, tmp100721, tmp100723, tmp100725)

				__e.TailApply(tmp100720, Jnk, tmp100726, V2889, V2890)
				return

			}, 0)

			__e.TailApply(tmp100718, V2888, V2887, V2889, tmp100719)
			return

		}, 1)

		tmp100727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp100728 := Call(__e, tmp100727, V2889)

		__e.TailApply(tmp100715, tmp100728)
		return

	}, 5)

	tmp100729 := Call(__e, PrimNS1Value(symns2_1set), symshen_4memo, tmp100714)

	_ = tmp100729

	tmp100730 := MakeNative(func(__e *ControlFlow) {
		V2892 := __e.Get(1)
		_ = V2892
		tmp100731 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5signature_6 := __e.Get(1)
			_ = Parse__shen_4_5signature_6
			tmp100754 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp100755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100756 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp100757 := Call(__e, tmp100756)

			tmp100758 := Call(__e, tmp100755, tmp100757, Parse__shen_4_5signature_6)

			tmp100759 := Call(__e, tmp100754, tmp100758)

			if True == tmp100759 {
				tmp100734 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5non_1ll_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5non_1ll_1rules_6
					tmp100746 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp100747 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100748 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp100749 := Call(__e, tmp100748)

					tmp100750 := Call(__e, tmp100747, tmp100749, Parse__shen_4_5non_1ll_1rules_6)

					tmp100751 := Call(__e, tmp100746, tmp100750)

					if True == tmp100751 {
						tmp100737 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp100738 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp100739 := Call(__e, tmp100738, Parse__shen_4_5non_1ll_1rules_6)

						tmp100740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp100741 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp100742 := Call(__e, tmp100741, Parse__shen_4_5signature_6)

						tmp100743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp100744 := Call(__e, tmp100743, Parse__shen_4_5non_1ll_1rules_6)

						tmp100745 := Call(__e, tmp100740, tmp100742, tmp100744)

						__e.TailApply(tmp100737, tmp100739, tmp100745)
						return

					} else {
						tmp100736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp100736)
						return

					}

				}, 1)

				tmp100752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5non_1ll_1rules_6)

				tmp100753 := Call(__e, tmp100752, Parse__shen_4_5signature_6)

				__e.TailApply(tmp100734, tmp100753)
				return

			} else {
				tmp100733 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp100733)
				return

			}

		}, 1)

		tmp100760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_6)

		tmp100761 := Call(__e, tmp100760, V2892)

		__e.TailApply(tmp100731, tmp100761)
		return

	}, 1)

	tmp100762 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sig_7rules_6, tmp100730)

	_ = tmp100762

	tmp100763 := MakeNative(func(__e *ControlFlow) {
		V2894 := __e.Get(1)
		_ = V2894
		tmp100764 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp100784 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100785 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp100786 := Call(__e, tmp100785)

			tmp100787 := Call(__e, tmp100784, YaccParse, tmp100786)

			if True == tmp100787 {
				tmp100766 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rule_6 := __e.Get(1)
					_ = Parse__shen_4_5rule_6
					tmp100776 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp100777 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100778 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp100779 := Call(__e, tmp100778)

					tmp100780 := Call(__e, tmp100777, tmp100779, Parse__shen_4_5rule_6)

					tmp100781 := Call(__e, tmp100776, tmp100780)

					if True == tmp100781 {
						tmp100769 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp100770 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp100771 := Call(__e, tmp100770, Parse__shen_4_5rule_6)

						tmp100772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp100773 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp100774 := Call(__e, tmp100773, Parse__shen_4_5rule_6)

						tmp100775 := Call(__e, tmp100772, tmp100774, Nil)

						__e.TailApply(tmp100769, tmp100771, tmp100775)
						return

					} else {
						tmp100768 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp100768)
						return

					}

				}, 1)

				tmp100782 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rule_6)

				tmp100783 := Call(__e, tmp100782, V2894)

				__e.TailApply(tmp100766, tmp100783)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp100788 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5rule_6 := __e.Get(1)
			_ = Parse__shen_4_5rule_6
			tmp100811 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp100812 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100813 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp100814 := Call(__e, tmp100813)

			tmp100815 := Call(__e, tmp100812, tmp100814, Parse__shen_4_5rule_6)

			tmp100816 := Call(__e, tmp100811, tmp100815)

			if True == tmp100816 {
				tmp100791 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5non_1ll_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5non_1ll_1rules_6
					tmp100803 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp100804 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100805 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp100806 := Call(__e, tmp100805)

					tmp100807 := Call(__e, tmp100804, tmp100806, Parse__shen_4_5non_1ll_1rules_6)

					tmp100808 := Call(__e, tmp100803, tmp100807)

					if True == tmp100808 {
						tmp100794 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp100795 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp100796 := Call(__e, tmp100795, Parse__shen_4_5non_1ll_1rules_6)

						tmp100797 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp100798 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp100799 := Call(__e, tmp100798, Parse__shen_4_5rule_6)

						tmp100800 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp100801 := Call(__e, tmp100800, Parse__shen_4_5non_1ll_1rules_6)

						tmp100802 := Call(__e, tmp100797, tmp100799, tmp100801)

						__e.TailApply(tmp100794, tmp100796, tmp100802)
						return

					} else {
						tmp100793 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp100793)
						return

					}

				}, 1)

				tmp100809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5non_1ll_1rules_6)

				tmp100810 := Call(__e, tmp100809, Parse__shen_4_5rule_6)

				__e.TailApply(tmp100791, tmp100810)
				return

			} else {
				tmp100790 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp100790)
				return

			}

		}, 1)

		tmp100817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5rule_6)

		tmp100818 := Call(__e, tmp100817, V2894)

		tmp100819 := Call(__e, tmp100788, tmp100818)

		__e.TailApply(tmp100764, tmp100819)
		return

	}, 1)

	tmp100820 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5non_1ll_1rules_6, tmp100763)

	_ = tmp100820

	tmp100821 := MakeNative(func(__e *ControlFlow) {
		V2896 := __e.Get(1)
		_ = V2896
		tmp100853 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp100854 := Call(__e, tmp100853, V2896)

		var ifres100833 Obj

		if True == tmp100854 {
			tmp100849 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100850 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp100851 := Call(__e, tmp100850, V2896)

			tmp100852 := Call(__e, tmp100849, tmp100851)

			var ifres100835 Obj

			if True == tmp100852 {
				tmp100843 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp100844 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp100845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp100846 := Call(__e, tmp100845, V2896)

				tmp100847 := Call(__e, tmp100844, tmp100846)

				tmp100848 := Call(__e, tmp100843, Nil, tmp100847)

				var ifres100837 Obj

				if True == tmp100848 {
					tmp100839 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100840 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp100841 := Call(__e, tmp100840, V2896)

					tmp100842 := Call(__e, tmp100839, tmp100841, symprotect)

					var ifres100838 Obj

					if True == tmp100842 {
						ifres100838 = True

					} else {
						ifres100838 = False

					}

					ifres100837 = ifres100838

				} else {
					ifres100837 = False

				}

				var ifres100836 Obj

				if True == ifres100837 {
					ifres100836 = True

				} else {
					ifres100836 = False

				}

				ifres100835 = ifres100836

			} else {
				ifres100835 = False

			}

			var ifres100834 Obj

			if True == ifres100835 {
				ifres100834 = True

			} else {
				ifres100834 = False

			}

			ifres100833 = ifres100834

		} else {
			ifres100833 = False

		}

		if True == ifres100833 {
			__e.Return(V2896)
			return
		} else {
			tmp100831 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100832 := Call(__e, tmp100831, V2896)

			if True == tmp100832 {
				tmp100828 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp100829 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp100830 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue)

					__e.TailApply(tmp100830, Z)
					return

				}, 1)

				__e.TailApply(tmp100828, tmp100829, V2896)
				return

			} else {
				tmp100826 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp100827 := Call(__e, tmp100826, V2896)

				if True == tmp100827 {
					tmp100825 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					__e.TailApply(tmp100825, sym_e_e, V2896)
					return

				} else {
					__e.Return(V2896)
					return
				}

			}

		}

	}, 1)

	tmp100855 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue, tmp100821)

	_ = tmp100855

	tmp100856 := MakeNative(func(__e *ControlFlow) {
		V2898 := __e.Get(1)
		_ = V2898
		tmp100865 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp100866 := Call(__e, tmp100865, V2898)

		if True == tmp100866 {
			tmp100862 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp100863 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				tmp100864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_1sig)

				__e.TailApply(tmp100864, Z)
				return

			}, 1)

			__e.TailApply(tmp100862, tmp100863, V2898)
			return

		} else {
			tmp100860 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp100861 := Call(__e, tmp100860, V2898)

			if True == tmp100861 {
				tmp100859 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				__e.TailApply(tmp100859, sym_e_e_e, V2898)
				return

			} else {
				__e.Return(V2898)
				return
			}

		}

	}, 1)

	tmp100867 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_1sig, tmp100856)

	_ = tmp100867

	tmp100868 := MakeNative(func(__e *ControlFlow) {
		V2904 := __e.Get(1)
		_ = V2904
		tmp100883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_2)

		tmp100884 := Call(__e, tmp100883, V2904)

		if True == tmp100884 {
			tmp100882 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp100882, V2904, Nil)
			return

		} else {
			tmp100880 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp100881 := Call(__e, tmp100880, V2904)

			if True == tmp100881 {
				tmp100871 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp100872 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ues)

				tmp100873 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp100874 := Call(__e, tmp100873, V2904)

				tmp100875 := Call(__e, tmp100872, tmp100874)

				tmp100876 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ues)

				tmp100877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp100878 := Call(__e, tmp100877, V2904)

				tmp100879 := Call(__e, tmp100876, tmp100878)

				__e.TailApply(tmp100871, tmp100875, tmp100879)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp100885 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ues, tmp100868)

	_ = tmp100885

	tmp100886 := MakeNative(func(__e *ControlFlow) {
		V2906 := __e.Get(1)
		_ = V2906
		tmp100893 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp100894 := Call(__e, tmp100893, V2906)

		if True == tmp100894 {
			tmp100889 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_1h_2)

			tmp100890 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			tmp100891 := Call(__e, tmp100890, V2906)

			tmp100892 := Call(__e, tmp100889, tmp100891)

			if True == tmp100892 {
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

	tmp100895 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_2, tmp100886)

	_ = tmp100895

	tmp100896 := MakeNative(func(__e *ControlFlow) {
		V2914 := __e.Get(1)
		_ = V2914
		tmp100918 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

		tmp100919 := Call(__e, tmp100918, V2914)

		var ifres100898 Obj

		if True == tmp100919 {
			tmp100914 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp100915 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			tmp100916 := Call(__e, tmp100915, V2914, MakeNumber(0))

			tmp100917 := Call(__e, tmp100914, MakeString("&"), tmp100916)

			var ifres100900 Obj

			if True == tmp100917 {
				tmp100910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

				tmp100911 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp100912 := Call(__e, tmp100911, V2914)

				tmp100913 := Call(__e, tmp100910, tmp100912)

				var ifres100902 Obj

				if True == tmp100913 {
					tmp100904 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp100905 := Call(__e, PrimNS1Value(symns2_1value), sympos)

					tmp100906 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					tmp100907 := Call(__e, tmp100906, V2914)

					tmp100908 := Call(__e, tmp100905, tmp100907, MakeNumber(0))

					tmp100909 := Call(__e, tmp100904, MakeString("&"), tmp100908)

					var ifres100903 Obj

					if True == tmp100909 {
						ifres100903 = True

					} else {
						ifres100903 = False

					}

					ifres100902 = ifres100903

				} else {
					ifres100902 = False

				}

				var ifres100901 Obj

				if True == ifres100902 {
					ifres100901 = True

				} else {
					ifres100901 = False

				}

				ifres100900 = ifres100901

			} else {
				ifres100900 = False

			}

			var ifres100899 Obj

			if True == ifres100900 {
				ifres100899 = True

			} else {
				ifres100899 = False

			}

			ifres100898 = ifres100899

		} else {
			ifres100898 = False

		}

		if True == ifres100898 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp100920 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_1h_2, tmp100896)

	_ = tmp100920

	tmp100921 := MakeNative(func(__e *ControlFlow) {
		V2922 := __e.Get(1)
		_ = V2922
		V2923 := __e.Get(2)
		_ = V2923
		V2924 := __e.Get(3)
		_ = V2924
		V2925 := __e.Get(4)
		_ = V2925
		V2926 := __e.Get(5)
		_ = V2926
		V2927 := __e.Get(6)
		_ = V2927
		V2928 := __e.Get(7)
		_ = V2928
		tmp100922 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp100923 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cutpoint)

			tmp100924 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp100974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp100975 := Call(__e, tmp100974, Case, False)

				if True == tmp100975 {
					tmp100926 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp100948 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp100949 := Call(__e, tmp100948, Case, False)

						if True == tmp100949 {
							tmp100928 := MakeNative(func(__e *ControlFlow) {
								Err := __e.Get(1)
								_ = Err
								tmp100929 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

								tmp100930 := Call(__e, tmp100929)

								_ = tmp100930

								tmp100931 := Call(__e, PrimNS1Value(symns2_1value), symbind)

								tmp100932 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								tmp100933 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								tmp100934 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								tmp100935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp100936 := Call(__e, tmp100935, V2924, V2927)

								tmp100937 := Call(__e, PrimNS1Value(symns2_1value), symcn)

								tmp100938 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

								tmp100939 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp100940 := Call(__e, tmp100939, V2925, V2927)

								tmp100941 := Call(__e, tmp100938, tmp100940, MakeString(""), symshen_4a)

								tmp100942 := Call(__e, tmp100937, MakeString(" of "), tmp100941)

								tmp100943 := Call(__e, tmp100934, tmp100936, tmp100942, symshen_4a)

								tmp100944 := Call(__e, tmp100933, MakeString("type error in rule "), tmp100943)

								tmp100945 := Call(__e, tmp100932, tmp100944)

								__e.TailApply(tmp100931, Err, tmp100945, V2927, V2928)
								return

							}, 1)

							tmp100946 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

							tmp100947 := Call(__e, tmp100946, V2927)

							__e.TailApply(tmp100928, tmp100947)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp100950 := MakeNative(func(__e *ControlFlow) {
						V2479 := __e.Get(1)
						_ = V2479
						tmp100969 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp100970 := Call(__e, tmp100969, V2479)

						if True == tmp100970 {
							tmp100952 := MakeNative(func(__e *ControlFlow) {
								Rule := __e.Get(1)
								_ = Rule
								tmp100953 := MakeNative(func(__e *ControlFlow) {
									Rules := __e.Get(1)
									_ = Rules
									tmp100954 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

									tmp100955 := Call(__e, tmp100954)

									_ = tmp100955

									tmp100956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1rule)

									tmp100957 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue)

									tmp100958 := Call(__e, tmp100957, Rule)

									tmp100959 := MakeNative(func(__e *ControlFlow) {
										tmp100960 := Call(__e, PrimNS1Value(symns2_1value), symcut)

										tmp100961 := MakeNative(func(__e *ControlFlow) {
											tmp100962 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1rules)

											tmp100963 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

											tmp100964 := Call(__e, tmp100963, V2924, MakeNumber(1))

											__e.TailApply(tmp100962, Rules, V2923, tmp100964, V2925, V2926, V2927, V2928)
											return

										}, 0)

										__e.TailApply(tmp100960, Throwcontrol, V2927, tmp100961)
										return

									}, 0)

									__e.TailApply(tmp100956, tmp100958, V2923, V2926, V2927, tmp100959)
									return

								}, 1)

								tmp100965 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp100966 := Call(__e, tmp100965, V2479)

								__e.TailApply(tmp100953, tmp100966)
								return

							}, 1)

							tmp100967 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp100968 := Call(__e, tmp100967, V2479)

							__e.TailApply(tmp100952, tmp100968)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp100971 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp100972 := Call(__e, tmp100971, V2922, V2927)

					tmp100973 := Call(__e, tmp100950, tmp100972)

					__e.TailApply(tmp100926, tmp100973)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp100976 := MakeNative(func(__e *ControlFlow) {
				V2478 := __e.Get(1)
				_ = V2478
				tmp100981 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp100982 := Call(__e, tmp100981, Nil, V2478)

				if True == tmp100982 {
					tmp100978 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

					tmp100979 := Call(__e, tmp100978)

					_ = tmp100979

					tmp100980 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

					__e.TailApply(tmp100980, V2928)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp100983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

			tmp100984 := Call(__e, tmp100983, V2922, V2927)

			tmp100985 := Call(__e, tmp100976, tmp100984)

			tmp100986 := Call(__e, tmp100924, tmp100985)

			__e.TailApply(tmp100923, Throwcontrol, tmp100986)
			return

		}, 1)

		tmp100987 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catchpoint)

		tmp100988 := Call(__e, tmp100987)

		__e.TailApply(tmp100922, tmp100988)
		return

	}, 7)

	tmp100989 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1rules, tmp100921)

	_ = tmp100989

	tmp100990 := MakeNative(func(__e *ControlFlow) {
		V2934 := __e.Get(1)
		_ = V2934
		V2935 := __e.Get(2)
		_ = V2935
		V2936 := __e.Get(3)
		_ = V2936
		V2937 := __e.Get(4)
		_ = V2937
		V2938 := __e.Get(5)
		_ = V2938
		tmp100991 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp100992 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cutpoint)

			tmp100993 := MakeNative(func(__e *ControlFlow) {
				V2470 := __e.Get(1)
				_ = V2470
				tmp101039 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp101040 := Call(__e, tmp101039, V2470)

				if True == tmp101040 {
					tmp100995 := MakeNative(func(__e *ControlFlow) {
						Patterns := __e.Get(1)
						_ = Patterns
						tmp100996 := MakeNative(func(__e *ControlFlow) {
							V2471 := __e.Get(1)
							_ = V2471
							tmp101031 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp101032 := Call(__e, tmp101031, V2471)

							if True == tmp101032 {
								tmp100998 := MakeNative(func(__e *ControlFlow) {
									Action := __e.Get(1)
									_ = Action
									tmp100999 := MakeNative(func(__e *ControlFlow) {
										V2472 := __e.Get(1)
										_ = V2472
										tmp101023 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp101024 := Call(__e, tmp101023, Nil, V2472)

										if True == tmp101024 {
											tmp101001 := MakeNative(func(__e *ControlFlow) {
												NewHyps := __e.Get(1)
												_ = NewHyps
												tmp101002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

												tmp101003 := Call(__e, tmp101002)

												_ = tmp101003

												tmp101004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

												tmp101005 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholders)

												tmp101006 := Call(__e, tmp101005, Patterns)

												tmp101007 := MakeNative(func(__e *ControlFlow) {
													tmp101008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1patterns)

													tmp101009 := MakeNative(func(__e *ControlFlow) {
														tmp101010 := Call(__e, PrimNS1Value(symns2_1value), symcut)

														tmp101011 := MakeNative(func(__e *ControlFlow) {
															tmp101012 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1action)

															tmp101013 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

															tmp101014 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue)

															tmp101015 := Call(__e, tmp101014, Action)

															tmp101016 := Call(__e, tmp101013, tmp101015)

															tmp101017 := Call(__e, PrimNS1Value(symns2_1value), symshen_4result_1type)

															tmp101018 := Call(__e, tmp101017, Patterns, V2935)

															tmp101019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4patthyps)

															tmp101020 := Call(__e, tmp101019, Patterns, V2935, V2936)

															__e.TailApply(tmp101012, tmp101016, tmp101018, tmp101020, V2937, V2938)
															return

														}, 0)

														__e.TailApply(tmp101010, Throwcontrol, V2937, tmp101011)
														return

													}, 0)

													__e.TailApply(tmp101008, Patterns, V2935, NewHyps, V2937, tmp101009)
													return

												}, 0)

												__e.TailApply(tmp101004, tmp101006, V2936, NewHyps, V2937, tmp101007)
												return

											}, 1)

											tmp101021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

											tmp101022 := Call(__e, tmp101021, V2937)

											__e.TailApply(tmp101001, tmp101022)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp101025 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

									tmp101026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp101027 := Call(__e, tmp101026, V2471)

									tmp101028 := Call(__e, tmp101025, tmp101027, V2937)

									__e.TailApply(tmp100999, tmp101028)
									return

								}, 1)

								tmp101029 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp101030 := Call(__e, tmp101029, V2471)

								__e.TailApply(tmp100998, tmp101030)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp101033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

						tmp101034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp101035 := Call(__e, tmp101034, V2470)

						tmp101036 := Call(__e, tmp101033, tmp101035, V2937)

						__e.TailApply(tmp100996, tmp101036)
						return

					}, 1)

					tmp101037 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp101038 := Call(__e, tmp101037, V2470)

					__e.TailApply(tmp100995, tmp101038)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp101041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

			tmp101042 := Call(__e, tmp101041, V2934, V2937)

			tmp101043 := Call(__e, tmp100993, tmp101042)

			__e.TailApply(tmp100992, Throwcontrol, tmp101043)
			return

		}, 1)

		tmp101044 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catchpoint)

		tmp101045 := Call(__e, tmp101044)

		__e.TailApply(tmp100991, tmp101045)
		return

	}, 5)

	tmp101046 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1rule, tmp100990)

	_ = tmp101046

	tmp101047 := MakeNative(func(__e *ControlFlow) {
		V2944 := __e.Get(1)
		_ = V2944
		tmp101062 := Call(__e, PrimNS1Value(symns2_1value), symshen_4ue_2)

		tmp101063 := Call(__e, tmp101062, V2944)

		if True == tmp101063 {
			tmp101061 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp101061, V2944, Nil)
			return

		} else {
			tmp101059 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp101060 := Call(__e, tmp101059, V2944)

			if True == tmp101060 {
				tmp101050 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp101051 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholders)

				tmp101052 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp101053 := Call(__e, tmp101052, V2944)

				tmp101054 := Call(__e, tmp101051, tmp101053)

				tmp101055 := Call(__e, PrimNS1Value(symns2_1value), symshen_4placeholders)

				tmp101056 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp101057 := Call(__e, tmp101056, V2944)

				tmp101058 := Call(__e, tmp101055, tmp101057)

				__e.TailApply(tmp101050, tmp101054, tmp101058)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp101064 := Call(__e, PrimNS1Value(symns2_1set), symshen_4placeholders, tmp101047)

	_ = tmp101064

	tmp101065 := MakeNative(func(__e *ControlFlow) {
		V2950 := __e.Get(1)
		_ = V2950
		V2951 := __e.Get(2)
		_ = V2951
		V2952 := __e.Get(3)
		_ = V2952
		V2953 := __e.Get(4)
		_ = V2953
		V2954 := __e.Get(5)
		_ = V2954
		tmp101066 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp101338 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101339 := Call(__e, tmp101338, Case, False)

			if True == tmp101339 {
				tmp101068 := MakeNative(func(__e *ControlFlow) {
					V2458 := __e.Get(1)
					_ = V2458
					tmp101334 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp101335 := Call(__e, tmp101334, V2458)

					if True == tmp101335 {
						tmp101070 := MakeNative(func(__e *ControlFlow) {
							V2453 := __e.Get(1)
							_ = V2453
							tmp101071 := MakeNative(func(__e *ControlFlow) {
								Vs := __e.Get(1)
								_ = Vs
								tmp101072 := MakeNative(func(__e *ControlFlow) {
									V2459 := __e.Get(1)
									_ = V2459
									tmp101326 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp101327 := Call(__e, tmp101326, V2459)

									if True == tmp101327 {
										tmp101105 := MakeNative(func(__e *ControlFlow) {
											V2460 := __e.Get(1)
											_ = V2460
											tmp101320 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp101321 := Call(__e, tmp101320, V2460)

											if True == tmp101321 {
												tmp101136 := MakeNative(func(__e *ControlFlow) {
													V := __e.Get(1)
													_ = V
													tmp101137 := MakeNative(func(__e *ControlFlow) {
														V2461 := __e.Get(1)
														_ = V2461
														tmp101312 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp101313 := Call(__e, tmp101312, V2461)

														if True == tmp101313 {
															tmp101163 := MakeNative(func(__e *ControlFlow) {
																V2462 := __e.Get(1)
																_ = V2462
																tmp101306 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp101307 := Call(__e, tmp101306, sym_h, V2462)

																if True == tmp101307 {
																	tmp101240 := MakeNative(func(__e *ControlFlow) {
																		V2463 := __e.Get(1)
																		_ = V2463
																		tmp101300 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp101301 := Call(__e, tmp101300, V2463)

																		if True == tmp101301 {
																			tmp101264 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp101265 := MakeNative(func(__e *ControlFlow) {
																					V2464 := __e.Get(1)
																					_ = V2464
																					tmp101292 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp101293 := Call(__e, tmp101292, Nil, V2464)

																					if True == tmp101293 {
																						tmp101284 := MakeNative(func(__e *ControlFlow) {
																							NewHyp := __e.Get(1)
																							_ = NewHyp
																							tmp101285 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																							tmp101286 := Call(__e, tmp101285)

																							_ = tmp101286

																							tmp101287 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																							tmp101288 := MakeNative(func(__e *ControlFlow) {
																								tmp101289 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																								__e.TailApply(tmp101289, Vs, V2951, NewHyp, V2953, V2954)
																								return

																							}, 0)

																							__e.TailApply(tmp101287, V, V2453, V2953, tmp101288)
																							return

																						}, 1)

																						tmp101290 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp101291 := Call(__e, tmp101290, V2459)

																						__e.TailApply(tmp101284, tmp101291)
																						return

																					} else {
																						tmp101282 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																						tmp101283 := Call(__e, tmp101282, V2464)

																						if True == tmp101283 {
																							tmp101268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																							tmp101269 := Call(__e, tmp101268, V2464, Nil, V2953)

																							_ = tmp101269

																							tmp101270 := MakeNative(func(__e *ControlFlow) {
																								Result := __e.Get(1)
																								_ = Result
																								tmp101271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																								tmp101272 := Call(__e, tmp101271, V2464, V2953)

																								_ = tmp101272

																								__e.Return(Result)
																								return

																							}, 1)

																							tmp101273 := MakeNative(func(__e *ControlFlow) {
																								NewHyp := __e.Get(1)
																								_ = NewHyp
																								tmp101274 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																								tmp101275 := Call(__e, tmp101274)

																								_ = tmp101275

																								tmp101276 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																								tmp101277 := MakeNative(func(__e *ControlFlow) {
																									tmp101278 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																									__e.TailApply(tmp101278, Vs, V2951, NewHyp, V2953, V2954)
																									return

																								}, 0)

																								__e.TailApply(tmp101276, V, V2453, V2953, tmp101277)
																								return

																							}, 1)

																							tmp101279 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp101280 := Call(__e, tmp101279, V2459)

																							tmp101281 := Call(__e, tmp101273, tmp101280)

																							__e.TailApply(tmp101270, tmp101281)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}

																				}, 1)

																				tmp101294 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp101295 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp101296 := Call(__e, tmp101295, V2463)

																				tmp101297 := Call(__e, tmp101294, tmp101296, V2953)

																				__e.TailApply(tmp101265, tmp101297)
																				return

																			}, 1)

																			tmp101298 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp101299 := Call(__e, tmp101298, V2463)

																			__e.TailApply(tmp101264, tmp101299)
																			return

																		} else {
																			tmp101262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																			tmp101263 := Call(__e, tmp101262, V2463)

																			if True == tmp101263 {
																				tmp101243 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp101244 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																					tmp101245 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																					tmp101246 := Call(__e, tmp101245, A, Nil)

																					tmp101247 := Call(__e, tmp101244, V2463, tmp101246, V2953)

																					_ = tmp101247

																					tmp101248 := MakeNative(func(__e *ControlFlow) {
																						Result := __e.Get(1)
																						_ = Result
																						tmp101249 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																						tmp101250 := Call(__e, tmp101249, V2463, V2953)

																						_ = tmp101250

																						__e.Return(Result)
																						return

																					}, 1)

																					tmp101251 := MakeNative(func(__e *ControlFlow) {
																						NewHyp := __e.Get(1)
																						_ = NewHyp
																						tmp101252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																						tmp101253 := Call(__e, tmp101252)

																						_ = tmp101253

																						tmp101254 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																						tmp101255 := MakeNative(func(__e *ControlFlow) {
																							tmp101256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																							__e.TailApply(tmp101256, Vs, V2951, NewHyp, V2953, V2954)
																							return

																						}, 0)

																						__e.TailApply(tmp101254, V, V2453, V2953, tmp101255)
																						return

																					}, 1)

																					tmp101257 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp101258 := Call(__e, tmp101257, V2459)

																					tmp101259 := Call(__e, tmp101251, tmp101258)

																					__e.TailApply(tmp101248, tmp101259)
																					return

																				}, 1)

																				tmp101260 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																				tmp101261 := Call(__e, tmp101260, V2953)

																				__e.TailApply(tmp101243, tmp101261)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp101302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp101303 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp101304 := Call(__e, tmp101303, V2461)

																	tmp101305 := Call(__e, tmp101302, tmp101304, V2953)

																	__e.TailApply(tmp101240, tmp101305)
																	return

																} else {
																	tmp101238 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																	tmp101239 := Call(__e, tmp101238, V2462)

																	if True == tmp101239 {
																		tmp101166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																		tmp101167 := Call(__e, tmp101166, V2462, sym_h, V2953)

																		_ = tmp101167

																		tmp101168 := MakeNative(func(__e *ControlFlow) {
																			Result := __e.Get(1)
																			_ = Result
																			tmp101169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																			tmp101170 := Call(__e, tmp101169, V2462, V2953)

																			_ = tmp101170

																			__e.Return(Result)
																			return

																		}, 1)

																		tmp101171 := MakeNative(func(__e *ControlFlow) {
																			V2465 := __e.Get(1)
																			_ = V2465
																			tmp101231 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																			tmp101232 := Call(__e, tmp101231, V2465)

																			if True == tmp101232 {
																				tmp101195 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp101196 := MakeNative(func(__e *ControlFlow) {
																						V2466 := __e.Get(1)
																						_ = V2466
																						tmp101223 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																						tmp101224 := Call(__e, tmp101223, Nil, V2466)

																						if True == tmp101224 {
																							tmp101215 := MakeNative(func(__e *ControlFlow) {
																								NewHyp := __e.Get(1)
																								_ = NewHyp
																								tmp101216 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																								tmp101217 := Call(__e, tmp101216)

																								_ = tmp101217

																								tmp101218 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																								tmp101219 := MakeNative(func(__e *ControlFlow) {
																									tmp101220 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																									__e.TailApply(tmp101220, Vs, V2951, NewHyp, V2953, V2954)
																									return

																								}, 0)

																								__e.TailApply(tmp101218, V, V2453, V2953, tmp101219)
																								return

																							}, 1)

																							tmp101221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																							tmp101222 := Call(__e, tmp101221, V2459)

																							__e.TailApply(tmp101215, tmp101222)
																							return

																						} else {
																							tmp101213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																							tmp101214 := Call(__e, tmp101213, V2466)

																							if True == tmp101214 {
																								tmp101199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																								tmp101200 := Call(__e, tmp101199, V2466, Nil, V2953)

																								_ = tmp101200

																								tmp101201 := MakeNative(func(__e *ControlFlow) {
																									Result := __e.Get(1)
																									_ = Result
																									tmp101202 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																									tmp101203 := Call(__e, tmp101202, V2466, V2953)

																									_ = tmp101203

																									__e.Return(Result)
																									return

																								}, 1)

																								tmp101204 := MakeNative(func(__e *ControlFlow) {
																									NewHyp := __e.Get(1)
																									_ = NewHyp
																									tmp101205 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																									tmp101206 := Call(__e, tmp101205)

																									_ = tmp101206

																									tmp101207 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																									tmp101208 := MakeNative(func(__e *ControlFlow) {
																										tmp101209 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																										__e.TailApply(tmp101209, Vs, V2951, NewHyp, V2953, V2954)
																										return

																									}, 0)

																									__e.TailApply(tmp101207, V, V2453, V2953, tmp101208)
																									return

																								}, 1)

																								tmp101210 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																								tmp101211 := Call(__e, tmp101210, V2459)

																								tmp101212 := Call(__e, tmp101204, tmp101211)

																								__e.TailApply(tmp101201, tmp101212)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}

																					}, 1)

																					tmp101225 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																					tmp101226 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																					tmp101227 := Call(__e, tmp101226, V2465)

																					tmp101228 := Call(__e, tmp101225, tmp101227, V2953)

																					__e.TailApply(tmp101196, tmp101228)
																					return

																				}, 1)

																				tmp101229 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																				tmp101230 := Call(__e, tmp101229, V2465)

																				__e.TailApply(tmp101195, tmp101230)
																				return

																			} else {
																				tmp101193 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

																				tmp101194 := Call(__e, tmp101193, V2465)

																				if True == tmp101194 {
																					tmp101174 := MakeNative(func(__e *ControlFlow) {
																						A := __e.Get(1)
																						_ = A
																						tmp101175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																						tmp101176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																						tmp101177 := Call(__e, tmp101176, A, Nil)

																						tmp101178 := Call(__e, tmp101175, V2465, tmp101177, V2953)

																						_ = tmp101178

																						tmp101179 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp101180 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																							tmp101181 := Call(__e, tmp101180, V2465, V2953)

																							_ = tmp101181

																							__e.Return(Result)
																							return

																						}, 1)

																						tmp101182 := MakeNative(func(__e *ControlFlow) {
																							NewHyp := __e.Get(1)
																							_ = NewHyp
																							tmp101183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																							tmp101184 := Call(__e, tmp101183)

																							_ = tmp101184

																							tmp101185 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																							tmp101186 := MakeNative(func(__e *ControlFlow) {
																								tmp101187 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																								__e.TailApply(tmp101187, Vs, V2951, NewHyp, V2953, V2954)
																								return

																							}, 0)

																							__e.TailApply(tmp101185, V, V2453, V2953, tmp101186)
																							return

																						}, 1)

																						tmp101188 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp101189 := Call(__e, tmp101188, V2459)

																						tmp101190 := Call(__e, tmp101182, tmp101189)

																						__e.TailApply(tmp101179, tmp101190)
																						return

																					}, 1)

																					tmp101191 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																					tmp101192 := Call(__e, tmp101191, V2953)

																					__e.TailApply(tmp101174, tmp101192)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}

																		}, 1)

																		tmp101233 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																		tmp101234 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp101235 := Call(__e, tmp101234, V2461)

																		tmp101236 := Call(__e, tmp101233, tmp101235, V2953)

																		tmp101237 := Call(__e, tmp101171, tmp101236)

																		__e.TailApply(tmp101168, tmp101237)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}

															}, 1)

															tmp101308 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp101309 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp101310 := Call(__e, tmp101309, V2461)

															tmp101311 := Call(__e, tmp101308, tmp101310, V2953)

															__e.TailApply(tmp101163, tmp101311)
															return

														} else {
															tmp101161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

															tmp101162 := Call(__e, tmp101161, V2461)

															if True == tmp101162 {
																tmp101140 := MakeNative(func(__e *ControlFlow) {
																	A := __e.Get(1)
																	_ = A
																	tmp101141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

																	tmp101142 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101143 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101144 := Call(__e, tmp101143, A, Nil)

																	tmp101145 := Call(__e, tmp101142, sym_h, tmp101144)

																	tmp101146 := Call(__e, tmp101141, V2461, tmp101145, V2953)

																	_ = tmp101146

																	tmp101147 := MakeNative(func(__e *ControlFlow) {
																		Result := __e.Get(1)
																		_ = Result
																		tmp101148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																		tmp101149 := Call(__e, tmp101148, V2461, V2953)

																		_ = tmp101149

																		__e.Return(Result)
																		return

																	}, 1)

																	tmp101150 := MakeNative(func(__e *ControlFlow) {
																		NewHyp := __e.Get(1)
																		_ = NewHyp
																		tmp101151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																		tmp101152 := Call(__e, tmp101151)

																		_ = tmp101152

																		tmp101153 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																		tmp101154 := MakeNative(func(__e *ControlFlow) {
																			tmp101155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																			__e.TailApply(tmp101155, Vs, V2951, NewHyp, V2953, V2954)
																			return

																		}, 0)

																		__e.TailApply(tmp101153, V, V2453, V2953, tmp101154)
																		return

																	}, 1)

																	tmp101156 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp101157 := Call(__e, tmp101156, V2459)

																	tmp101158 := Call(__e, tmp101150, tmp101157)

																	__e.TailApply(tmp101147, tmp101158)
																	return

																}, 1)

																tmp101159 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

																tmp101160 := Call(__e, tmp101159, V2953)

																__e.TailApply(tmp101140, tmp101160)
																return

															} else {
																__e.Return(False)
																return
															}

														}

													}, 1)

													tmp101314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp101315 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp101316 := Call(__e, tmp101315, V2460)

													tmp101317 := Call(__e, tmp101314, tmp101316, V2953)

													__e.TailApply(tmp101137, tmp101317)
													return

												}, 1)

												tmp101318 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp101319 := Call(__e, tmp101318, V2460)

												__e.TailApply(tmp101136, tmp101319)
												return

											} else {
												tmp101134 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

												tmp101135 := Call(__e, tmp101134, V2460)

												if True == tmp101135 {
													tmp101108 := MakeNative(func(__e *ControlFlow) {
														V := __e.Get(1)
														_ = V
														tmp101109 := MakeNative(func(__e *ControlFlow) {
															A := __e.Get(1)
															_ = A
															tmp101110 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

															tmp101111 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp101112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp101113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

															tmp101114 := Call(__e, tmp101113, A, Nil)

															tmp101115 := Call(__e, tmp101112, sym_h, tmp101114)

															tmp101116 := Call(__e, tmp101111, V, tmp101115)

															tmp101117 := Call(__e, tmp101110, V2460, tmp101116, V2953)

															_ = tmp101117

															tmp101118 := MakeNative(func(__e *ControlFlow) {
																Result := __e.Get(1)
																_ = Result
																tmp101119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

																tmp101120 := Call(__e, tmp101119, V2460, V2953)

																_ = tmp101120

																__e.Return(Result)
																return

															}, 1)

															tmp101121 := MakeNative(func(__e *ControlFlow) {
																NewHyp := __e.Get(1)
																_ = NewHyp
																tmp101122 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																tmp101123 := Call(__e, tmp101122)

																_ = tmp101123

																tmp101124 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

																tmp101125 := MakeNative(func(__e *ControlFlow) {
																	tmp101126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

																	__e.TailApply(tmp101126, Vs, V2951, NewHyp, V2953, V2954)
																	return

																}, 0)

																__e.TailApply(tmp101124, V, V2453, V2953, tmp101125)
																return

															}, 1)

															tmp101127 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp101128 := Call(__e, tmp101127, V2459)

															tmp101129 := Call(__e, tmp101121, tmp101128)

															__e.TailApply(tmp101118, tmp101129)
															return

														}, 1)

														tmp101130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

														tmp101131 := Call(__e, tmp101130, V2953)

														__e.TailApply(tmp101109, tmp101131)
														return

													}, 1)

													tmp101132 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

													tmp101133 := Call(__e, tmp101132, V2953)

													__e.TailApply(tmp101108, tmp101133)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp101322 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp101323 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp101324 := Call(__e, tmp101323, V2459)

										tmp101325 := Call(__e, tmp101322, tmp101324, V2953)

										__e.TailApply(tmp101105, tmp101325)
										return

									} else {
										tmp101103 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

										tmp101104 := Call(__e, tmp101103, V2459)

										if True == tmp101104 {
											tmp101075 := MakeNative(func(__e *ControlFlow) {
												V := __e.Get(1)
												_ = V
												tmp101076 := MakeNative(func(__e *ControlFlow) {
													A := __e.Get(1)
													_ = A
													tmp101077 := MakeNative(func(__e *ControlFlow) {
														NewHyp := __e.Get(1)
														_ = NewHyp
														tmp101078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4bindv)

														tmp101079 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp101080 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp101081 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp101082 := Call(__e, PrimNS1Value(symns2_1value), symcons)

														tmp101083 := Call(__e, tmp101082, A, Nil)

														tmp101084 := Call(__e, tmp101081, sym_h, tmp101083)

														tmp101085 := Call(__e, tmp101080, V, tmp101084)

														tmp101086 := Call(__e, tmp101079, tmp101085, NewHyp)

														tmp101087 := Call(__e, tmp101078, V2459, tmp101086, V2953)

														_ = tmp101087

														tmp101088 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp101089 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unbindv)

															tmp101090 := Call(__e, tmp101089, V2459, V2953)

															_ = tmp101090

															__e.Return(Result)
															return

														}, 1)

														tmp101091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

														tmp101092 := Call(__e, tmp101091)

														_ = tmp101092

														tmp101093 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

														tmp101094 := MakeNative(func(__e *ControlFlow) {
															tmp101095 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newhyps)

															__e.TailApply(tmp101095, Vs, V2951, NewHyp, V2953, V2954)
															return

														}, 0)

														tmp101096 := Call(__e, tmp101093, V, V2453, V2953, tmp101094)

														__e.TailApply(tmp101088, tmp101096)
														return

													}, 1)

													tmp101097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

													tmp101098 := Call(__e, tmp101097, V2953)

													__e.TailApply(tmp101077, tmp101098)
													return

												}, 1)

												tmp101099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

												tmp101100 := Call(__e, tmp101099, V2953)

												__e.TailApply(tmp101076, tmp101100)
												return

											}, 1)

											tmp101101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

											tmp101102 := Call(__e, tmp101101, V2953)

											__e.TailApply(tmp101075, tmp101102)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp101328 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp101329 := Call(__e, tmp101328, V2952, V2953)

								__e.TailApply(tmp101072, tmp101329)
								return

							}, 1)

							tmp101330 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101331 := Call(__e, tmp101330, V2458)

							__e.TailApply(tmp101071, tmp101331)
							return

						}, 1)

						tmp101332 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp101333 := Call(__e, tmp101332, V2458)

						__e.TailApply(tmp101070, tmp101333)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp101336 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp101337 := Call(__e, tmp101336, V2950, V2953)

				__e.TailApply(tmp101068, tmp101337)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp101340 := MakeNative(func(__e *ControlFlow) {
			V2457 := __e.Get(1)
			_ = V2457
			tmp101345 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101346 := Call(__e, tmp101345, Nil, V2457)

			if True == tmp101346 {
				tmp101342 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp101343 := Call(__e, tmp101342)

				_ = tmp101343

				tmp101344 := Call(__e, PrimNS1Value(symns2_1value), symunify_b)

				__e.TailApply(tmp101344, V2952, V2951, V2953, V2954)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp101347 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp101348 := Call(__e, tmp101347, V2950, V2953)

		tmp101349 := Call(__e, tmp101340, tmp101348)

		__e.TailApply(tmp101066, tmp101349)
		return

	}, 5)

	tmp101350 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newhyps, tmp101065)

	_ = tmp101350

	tmp101351 := MakeNative(func(__e *ControlFlow) {
		V2960 := __e.Get(1)
		_ = V2960
		V2961 := __e.Get(2)
		_ = V2961
		V2962 := __e.Get(3)
		_ = V2962
		tmp101416 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp101417 := Call(__e, tmp101416, Nil, V2960)

		if True == tmp101417 {
			__e.Return(V2962)
			return
		} else {
			tmp101414 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp101415 := Call(__e, tmp101414, V2960)

			var ifres101378 Obj

			if True == tmp101415 {
				tmp101412 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp101413 := Call(__e, tmp101412, V2961)

				var ifres101380 Obj

				if True == tmp101413 {
					tmp101408 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp101409 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp101410 := Call(__e, tmp101409, V2961)

					tmp101411 := Call(__e, tmp101408, tmp101410)

					var ifres101382 Obj

					if True == tmp101411 {
						tmp101402 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp101403 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp101404 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp101405 := Call(__e, tmp101404, V2961)

						tmp101406 := Call(__e, tmp101403, tmp101405)

						tmp101407 := Call(__e, tmp101402, sym_1_1_6, tmp101406)

						var ifres101384 Obj

						if True == tmp101407 {
							tmp101396 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp101397 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101398 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101399 := Call(__e, tmp101398, V2961)

							tmp101400 := Call(__e, tmp101397, tmp101399)

							tmp101401 := Call(__e, tmp101396, tmp101400)

							var ifres101386 Obj

							if True == tmp101401 {
								tmp101388 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp101389 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp101390 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp101391 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp101392 := Call(__e, tmp101391, V2961)

								tmp101393 := Call(__e, tmp101390, tmp101392)

								tmp101394 := Call(__e, tmp101389, tmp101393)

								tmp101395 := Call(__e, tmp101388, Nil, tmp101394)

								var ifres101387 Obj

								if True == tmp101395 {
									ifres101387 = True

								} else {
									ifres101387 = False

								}

								ifres101386 = ifres101387

							} else {
								ifres101386 = False

							}

							var ifres101385 Obj

							if True == ifres101386 {
								ifres101385 = True

							} else {
								ifres101385 = False

							}

							ifres101384 = ifres101385

						} else {
							ifres101384 = False

						}

						var ifres101383 Obj

						if True == ifres101384 {
							ifres101383 = True

						} else {
							ifres101383 = False

						}

						ifres101382 = ifres101383

					} else {
						ifres101382 = False

					}

					var ifres101381 Obj

					if True == ifres101382 {
						ifres101381 = True

					} else {
						ifres101381 = False

					}

					ifres101380 = ifres101381

				} else {
					ifres101380 = False

				}

				var ifres101379 Obj

				if True == ifres101380 {
					ifres101379 = True

				} else {
					ifres101379 = False

				}

				ifres101378 = ifres101379

			} else {
				ifres101378 = False

			}

			if True == ifres101378 {
				tmp101355 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

				tmp101356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp101357 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

				tmp101358 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp101359 := Call(__e, tmp101358, V2960)

				tmp101360 := Call(__e, tmp101357, tmp101359)

				tmp101361 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp101362 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp101363 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp101364 := Call(__e, tmp101363, V2961)

				tmp101365 := Call(__e, tmp101362, tmp101364, Nil)

				tmp101366 := Call(__e, tmp101361, sym_h, tmp101365)

				tmp101367 := Call(__e, tmp101356, tmp101360, tmp101366)

				tmp101368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4patthyps)

				tmp101369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp101370 := Call(__e, tmp101369, V2960)

				tmp101371 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp101372 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp101373 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp101374 := Call(__e, tmp101373, V2961)

				tmp101375 := Call(__e, tmp101372, tmp101374)

				tmp101376 := Call(__e, tmp101371, tmp101375)

				tmp101377 := Call(__e, tmp101368, tmp101370, tmp101376, V2962)

				__e.TailApply(tmp101355, tmp101367, tmp101377)
				return

			} else {
				tmp101354 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp101354, symshen_4patthyps)
				return

			}

		}

	}, 3)

	tmp101418 := Call(__e, PrimNS1Value(symns2_1set), symshen_4patthyps, tmp101351)

	_ = tmp101418

	tmp101419 := MakeNative(func(__e *ControlFlow) {
		V2969 := __e.Get(1)
		_ = V2969
		V2970 := __e.Get(2)
		_ = V2970
		tmp101500 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp101501 := Call(__e, tmp101500, Nil, V2969)

		var ifres101476 Obj

		if True == tmp101501 {
			tmp101498 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp101499 := Call(__e, tmp101498, V2970)

			var ifres101478 Obj

			if True == tmp101499 {
				tmp101494 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp101495 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp101496 := Call(__e, tmp101495, V2970)

				tmp101497 := Call(__e, tmp101494, sym_1_1_6, tmp101496)

				var ifres101480 Obj

				if True == tmp101497 {
					tmp101490 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp101491 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp101492 := Call(__e, tmp101491, V2970)

					tmp101493 := Call(__e, tmp101490, tmp101492)

					var ifres101482 Obj

					if True == tmp101493 {
						tmp101484 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp101485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp101486 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp101487 := Call(__e, tmp101486, V2970)

						tmp101488 := Call(__e, tmp101485, tmp101487)

						tmp101489 := Call(__e, tmp101484, Nil, tmp101488)

						var ifres101483 Obj

						if True == tmp101489 {
							ifres101483 = True

						} else {
							ifres101483 = False

						}

						ifres101482 = ifres101483

					} else {
						ifres101482 = False

					}

					var ifres101481 Obj

					if True == ifres101482 {
						ifres101481 = True

					} else {
						ifres101481 = False

					}

					ifres101480 = ifres101481

				} else {
					ifres101480 = False

				}

				var ifres101479 Obj

				if True == ifres101480 {
					ifres101479 = True

				} else {
					ifres101479 = False

				}

				ifres101478 = ifres101479

			} else {
				ifres101478 = False

			}

			var ifres101477 Obj

			if True == ifres101478 {
				ifres101477 = True

			} else {
				ifres101477 = False

			}

			ifres101476 = ifres101477

		} else {
			ifres101476 = False

		}

		if True == ifres101476 {
			tmp101473 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp101474 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp101475 := Call(__e, tmp101474, V2970)

			__e.TailApply(tmp101473, tmp101475)
			return

		} else {
			tmp101471 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101472 := Call(__e, tmp101471, Nil, V2969)

			if True == tmp101472 {
				__e.Return(V2970)
				return
			} else {
				tmp101469 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp101470 := Call(__e, tmp101469, V2969)

				var ifres101433 Obj

				if True == tmp101470 {
					tmp101467 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp101468 := Call(__e, tmp101467, V2970)

					var ifres101435 Obj

					if True == tmp101468 {
						tmp101463 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp101464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp101465 := Call(__e, tmp101464, V2970)

						tmp101466 := Call(__e, tmp101463, tmp101465)

						var ifres101437 Obj

						if True == tmp101466 {
							tmp101457 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp101458 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp101459 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101460 := Call(__e, tmp101459, V2970)

							tmp101461 := Call(__e, tmp101458, tmp101460)

							tmp101462 := Call(__e, tmp101457, sym_1_1_6, tmp101461)

							var ifres101439 Obj

							if True == tmp101462 {
								tmp101451 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp101452 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp101453 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp101454 := Call(__e, tmp101453, V2970)

								tmp101455 := Call(__e, tmp101452, tmp101454)

								tmp101456 := Call(__e, tmp101451, tmp101455)

								var ifres101441 Obj

								if True == tmp101456 {
									tmp101443 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp101444 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp101445 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp101446 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp101447 := Call(__e, tmp101446, V2970)

									tmp101448 := Call(__e, tmp101445, tmp101447)

									tmp101449 := Call(__e, tmp101444, tmp101448)

									tmp101450 := Call(__e, tmp101443, Nil, tmp101449)

									var ifres101442 Obj

									if True == tmp101450 {
										ifres101442 = True

									} else {
										ifres101442 = False

									}

									ifres101441 = ifres101442

								} else {
									ifres101441 = False

								}

								var ifres101440 Obj

								if True == ifres101441 {
									ifres101440 = True

								} else {
									ifres101440 = False

								}

								ifres101439 = ifres101440

							} else {
								ifres101439 = False

							}

							var ifres101438 Obj

							if True == ifres101439 {
								ifres101438 = True

							} else {
								ifres101438 = False

							}

							ifres101437 = ifres101438

						} else {
							ifres101437 = False

						}

						var ifres101436 Obj

						if True == ifres101437 {
							ifres101436 = True

						} else {
							ifres101436 = False

						}

						ifres101435 = ifres101436

					} else {
						ifres101435 = False

					}

					var ifres101434 Obj

					if True == ifres101435 {
						ifres101434 = True

					} else {
						ifres101434 = False

					}

					ifres101433 = ifres101434

				} else {
					ifres101433 = False

				}

				if True == ifres101433 {
					tmp101424 := Call(__e, PrimNS1Value(symns2_1value), symshen_4result_1type)

					tmp101425 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp101426 := Call(__e, tmp101425, V2969)

					tmp101427 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp101428 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp101429 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp101430 := Call(__e, tmp101429, V2970)

					tmp101431 := Call(__e, tmp101428, tmp101430)

					tmp101432 := Call(__e, tmp101427, tmp101431)

					__e.TailApply(tmp101424, tmp101426, tmp101432)
					return

				} else {
					tmp101423 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp101423, symshen_4result_1type)
					return

				}

			}

		}

	}, 2)

	tmp101502 := Call(__e, PrimNS1Value(symns2_1set), symshen_4result_1type, tmp101419)

	_ = tmp101502

	tmp101503 := MakeNative(func(__e *ControlFlow) {
		V2976 := __e.Get(1)
		_ = V2976
		V2977 := __e.Get(2)
		_ = V2977
		V2978 := __e.Get(3)
		_ = V2978
		V2979 := __e.Get(4)
		_ = V2979
		V2980 := __e.Get(5)
		_ = V2980
		tmp101504 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp101575 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101576 := Call(__e, tmp101575, Case, False)

			if True == tmp101576 {
				tmp101506 := MakeNative(func(__e *ControlFlow) {
					V2446 := __e.Get(1)
					_ = V2446
					tmp101571 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp101572 := Call(__e, tmp101571, V2446)

					if True == tmp101572 {
						tmp101508 := MakeNative(func(__e *ControlFlow) {
							Pattern := __e.Get(1)
							_ = Pattern
							tmp101509 := MakeNative(func(__e *ControlFlow) {
								Patterns := __e.Get(1)
								_ = Patterns
								tmp101510 := MakeNative(func(__e *ControlFlow) {
									V2447 := __e.Get(1)
									_ = V2447
									tmp101563 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp101564 := Call(__e, tmp101563, V2447)

									if True == tmp101564 {
										tmp101512 := MakeNative(func(__e *ControlFlow) {
											A := __e.Get(1)
											_ = A
											tmp101513 := MakeNative(func(__e *ControlFlow) {
												V2448 := __e.Get(1)
												_ = V2448
												tmp101555 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp101556 := Call(__e, tmp101555, V2448)

												if True == tmp101556 {
													tmp101515 := MakeNative(func(__e *ControlFlow) {
														V2449 := __e.Get(1)
														_ = V2449
														tmp101549 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp101550 := Call(__e, tmp101549, sym_1_1_6, V2449)

														if True == tmp101550 {
															tmp101517 := MakeNative(func(__e *ControlFlow) {
																V2450 := __e.Get(1)
																_ = V2450
																tmp101543 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																tmp101544 := Call(__e, tmp101543, V2450)

																if True == tmp101544 {
																	tmp101519 := MakeNative(func(__e *ControlFlow) {
																		B := __e.Get(1)
																		_ = B
																		tmp101520 := MakeNative(func(__e *ControlFlow) {
																			V2451 := __e.Get(1)
																			_ = V2451
																			tmp101535 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																			tmp101536 := Call(__e, tmp101535, Nil, V2451)

																			if True == tmp101536 {
																				tmp101522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																				tmp101523 := Call(__e, tmp101522)

																				_ = tmp101523

																				tmp101524 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d)

																				tmp101525 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				tmp101526 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

																				tmp101527 := Call(__e, tmp101526, Pattern)

																				tmp101528 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				tmp101529 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																				tmp101530 := Call(__e, tmp101529, A, Nil)

																				tmp101531 := Call(__e, tmp101528, sym_h, tmp101530)

																				tmp101532 := Call(__e, tmp101525, tmp101527, tmp101531)

																				tmp101533 := MakeNative(func(__e *ControlFlow) {
																					tmp101534 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1patterns)

																					__e.TailApply(tmp101534, Patterns, B, V2978, V2979, V2980)
																					return

																				}, 0)

																				__e.TailApply(tmp101524, tmp101532, V2978, V2979, tmp101533)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp101537 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																		tmp101538 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																		tmp101539 := Call(__e, tmp101538, V2450)

																		tmp101540 := Call(__e, tmp101537, tmp101539, V2979)

																		__e.TailApply(tmp101520, tmp101540)
																		return

																	}, 1)

																	tmp101541 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																	tmp101542 := Call(__e, tmp101541, V2450)

																	__e.TailApply(tmp101519, tmp101542)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp101545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp101546 := Call(__e, PrimNS1Value(symns2_1value), symtl)

															tmp101547 := Call(__e, tmp101546, V2448)

															tmp101548 := Call(__e, tmp101545, tmp101547, V2979)

															__e.TailApply(tmp101517, tmp101548)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp101551 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp101552 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp101553 := Call(__e, tmp101552, V2448)

													tmp101554 := Call(__e, tmp101551, tmp101553, V2979)

													__e.TailApply(tmp101515, tmp101554)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp101557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp101558 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp101559 := Call(__e, tmp101558, V2447)

											tmp101560 := Call(__e, tmp101557, tmp101559, V2979)

											__e.TailApply(tmp101513, tmp101560)
											return

										}, 1)

										tmp101561 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp101562 := Call(__e, tmp101561, V2447)

										__e.TailApply(tmp101512, tmp101562)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp101565 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

								tmp101566 := Call(__e, tmp101565, V2977, V2979)

								__e.TailApply(tmp101510, tmp101566)
								return

							}, 1)

							tmp101567 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101568 := Call(__e, tmp101567, V2446)

							__e.TailApply(tmp101509, tmp101568)
							return

						}, 1)

						tmp101569 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp101570 := Call(__e, tmp101569, V2446)

						__e.TailApply(tmp101508, tmp101570)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp101573 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp101574 := Call(__e, tmp101573, V2976, V2979)

				__e.TailApply(tmp101506, tmp101574)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp101577 := MakeNative(func(__e *ControlFlow) {
			V2445 := __e.Get(1)
			_ = V2445
			tmp101582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101583 := Call(__e, tmp101582, Nil, V2445)

			if True == tmp101583 {
				tmp101579 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp101580 := Call(__e, tmp101579)

				_ = tmp101580

				tmp101581 := Call(__e, PrimNS1Value(symns2_1value), symthaw)

				__e.TailApply(tmp101581, V2980)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp101584 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

		tmp101585 := Call(__e, tmp101584, V2976, V2979)

		tmp101586 := Call(__e, tmp101577, tmp101585)

		__e.TailApply(tmp101504, tmp101586)
		return

	}, 5)

	tmp101587 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1patterns, tmp101503)

	_ = tmp101587

	tmp101588 := MakeNative(func(__e *ControlFlow) {
		V2986 := __e.Get(1)
		_ = V2986
		V2987 := __e.Get(2)
		_ = V2987
		V2988 := __e.Get(3)
		_ = V2988
		V2989 := __e.Get(4)
		_ = V2989
		V2990 := __e.Get(5)
		_ = V2990
		tmp101589 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp101590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cutpoint)

			tmp101591 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp101781 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp101782 := Call(__e, tmp101781, Case, False)

				if True == tmp101782 {
					tmp101593 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp101667 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp101668 := Call(__e, tmp101667, Case, False)

						if True == tmp101668 {
							tmp101595 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp101606 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp101607 := Call(__e, tmp101606, Case, False)

								if True == tmp101607 {
									tmp101597 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

									tmp101598 := Call(__e, tmp101597)

									_ = tmp101598

									tmp101599 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d)

									tmp101600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp101601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp101602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

									tmp101603 := Call(__e, tmp101602, V2987, Nil)

									tmp101604 := Call(__e, tmp101601, sym_h, tmp101603)

									tmp101605 := Call(__e, tmp101600, V2986, tmp101604)

									__e.TailApply(tmp101599, tmp101605, V2988, V2989, V2990)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp101608 := MakeNative(func(__e *ControlFlow) {
								V2438 := __e.Get(1)
								_ = V2438
								tmp101662 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp101663 := Call(__e, tmp101662, V2438)

								if True == tmp101663 {
									tmp101610 := MakeNative(func(__e *ControlFlow) {
										V2439 := __e.Get(1)
										_ = V2439
										tmp101656 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

										tmp101657 := Call(__e, tmp101656, symshen_4choicepoint_b, V2439)

										if True == tmp101657 {
											tmp101612 := MakeNative(func(__e *ControlFlow) {
												V2440 := __e.Get(1)
												_ = V2440
												tmp101650 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp101651 := Call(__e, tmp101650, V2440)

												if True == tmp101651 {
													tmp101614 := MakeNative(func(__e *ControlFlow) {
														Action := __e.Get(1)
														_ = Action
														tmp101615 := MakeNative(func(__e *ControlFlow) {
															V2441 := __e.Get(1)
															_ = V2441
															tmp101642 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

															tmp101643 := Call(__e, tmp101642, Nil, V2441)

															if True == tmp101643 {
																tmp101617 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																tmp101618 := Call(__e, tmp101617)

																_ = tmp101618

																tmp101619 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																tmp101620 := MakeNative(func(__e *ControlFlow) {
																	tmp101621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1action)

																	tmp101622 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101625 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101626 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101627 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101629 := Call(__e, tmp101628, Action, Nil)

																	tmp101630 := Call(__e, tmp101627, sym_a, tmp101629)

																	tmp101631 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101632 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101633 := Call(__e, tmp101632, symfail, Nil)

																	tmp101634 := Call(__e, tmp101631, tmp101633, Nil)

																	tmp101635 := Call(__e, tmp101626, tmp101630, tmp101634)

																	tmp101636 := Call(__e, tmp101625, tmp101635, Nil)

																	tmp101637 := Call(__e, tmp101624, symnot, tmp101636)

																	tmp101638 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																	tmp101639 := Call(__e, tmp101638, Action, Nil)

																	tmp101640 := Call(__e, tmp101623, tmp101637, tmp101639)

																	tmp101641 := Call(__e, tmp101622, symwhere, tmp101640)

																	__e.TailApply(tmp101621, tmp101641, V2987, V2988, V2989, V2990)
																	return

																}, 0)

																__e.TailApply(tmp101619, Throwcontrol, V2989, tmp101620)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp101644 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

														tmp101645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

														tmp101646 := Call(__e, tmp101645, V2440)

														tmp101647 := Call(__e, tmp101644, tmp101646, V2989)

														__e.TailApply(tmp101615, tmp101647)
														return

													}, 1)

													tmp101648 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp101649 := Call(__e, tmp101648, V2440)

													__e.TailApply(tmp101614, tmp101649)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp101652 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp101653 := Call(__e, PrimNS1Value(symns2_1value), symtl)

											tmp101654 := Call(__e, tmp101653, V2438)

											tmp101655 := Call(__e, tmp101652, tmp101654, V2989)

											__e.TailApply(tmp101612, tmp101655)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp101658 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

									tmp101659 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp101660 := Call(__e, tmp101659, V2438)

									tmp101661 := Call(__e, tmp101658, tmp101660, V2989)

									__e.TailApply(tmp101610, tmp101661)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp101664 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp101665 := Call(__e, tmp101664, V2986, V2989)

							tmp101666 := Call(__e, tmp101608, tmp101665)

							__e.TailApply(tmp101595, tmp101666)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp101669 := MakeNative(func(__e *ControlFlow) {
						V2427 := __e.Get(1)
						_ = V2427
						tmp101776 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp101777 := Call(__e, tmp101776, V2427)

						if True == tmp101777 {
							tmp101671 := MakeNative(func(__e *ControlFlow) {
								V2428 := __e.Get(1)
								_ = V2428
								tmp101770 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp101771 := Call(__e, tmp101770, symshen_4choicepoint_b, V2428)

								if True == tmp101771 {
									tmp101673 := MakeNative(func(__e *ControlFlow) {
										V2429 := __e.Get(1)
										_ = V2429
										tmp101764 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

										tmp101765 := Call(__e, tmp101764, V2429)

										if True == tmp101765 {
											tmp101675 := MakeNative(func(__e *ControlFlow) {
												V2430 := __e.Get(1)
												_ = V2430
												tmp101758 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

												tmp101759 := Call(__e, tmp101758, V2430)

												if True == tmp101759 {
													tmp101677 := MakeNative(func(__e *ControlFlow) {
														V2431 := __e.Get(1)
														_ = V2431
														tmp101752 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

														tmp101753 := Call(__e, tmp101752, V2431)

														if True == tmp101753 {
															tmp101679 := MakeNative(func(__e *ControlFlow) {
																V2432 := __e.Get(1)
																_ = V2432
																tmp101746 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																tmp101747 := Call(__e, tmp101746, symfail_1if, V2432)

																if True == tmp101747 {
																	tmp101681 := MakeNative(func(__e *ControlFlow) {
																		V2433 := __e.Get(1)
																		_ = V2433
																		tmp101740 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																		tmp101741 := Call(__e, tmp101740, V2433)

																		if True == tmp101741 {
																			tmp101683 := MakeNative(func(__e *ControlFlow) {
																				F := __e.Get(1)
																				_ = F
																				tmp101684 := MakeNative(func(__e *ControlFlow) {
																					V2434 := __e.Get(1)
																					_ = V2434
																					tmp101732 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																					tmp101733 := Call(__e, tmp101732, Nil, V2434)

																					if True == tmp101733 {
																						tmp101686 := MakeNative(func(__e *ControlFlow) {
																							V2435 := __e.Get(1)
																							_ = V2435
																							tmp101726 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

																							tmp101727 := Call(__e, tmp101726, V2435)

																							if True == tmp101727 {
																								tmp101688 := MakeNative(func(__e *ControlFlow) {
																									Action := __e.Get(1)
																									_ = Action
																									tmp101689 := MakeNative(func(__e *ControlFlow) {
																										V2436 := __e.Get(1)
																										_ = V2436
																										tmp101718 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																										tmp101719 := Call(__e, tmp101718, Nil, V2436)

																										if True == tmp101719 {
																											tmp101691 := MakeNative(func(__e *ControlFlow) {
																												V2437 := __e.Get(1)
																												_ = V2437
																												tmp101712 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

																												tmp101713 := Call(__e, tmp101712, Nil, V2437)

																												if True == tmp101713 {
																													tmp101693 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

																													tmp101694 := Call(__e, tmp101693)

																													_ = tmp101694

																													tmp101695 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																													tmp101696 := MakeNative(func(__e *ControlFlow) {
																														tmp101697 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1action)

																														tmp101698 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101699 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101700 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101701 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101702 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101704 := Call(__e, tmp101703, Action, Nil)

																														tmp101705 := Call(__e, tmp101702, F, tmp101704)

																														tmp101706 := Call(__e, tmp101701, tmp101705, Nil)

																														tmp101707 := Call(__e, tmp101700, symnot, tmp101706)

																														tmp101708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																														tmp101709 := Call(__e, tmp101708, Action, Nil)

																														tmp101710 := Call(__e, tmp101699, tmp101707, tmp101709)

																														tmp101711 := Call(__e, tmp101698, symwhere, tmp101710)

																														__e.TailApply(tmp101697, tmp101711, V2987, V2988, V2989, V2990)
																														return

																													}, 0)

																													__e.TailApply(tmp101695, Throwcontrol, V2989, tmp101696)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp101714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																											tmp101715 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																											tmp101716 := Call(__e, tmp101715, V2429)

																											tmp101717 := Call(__e, tmp101714, tmp101716, V2989)

																											__e.TailApply(tmp101691, tmp101717)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp101720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																									tmp101721 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																									tmp101722 := Call(__e, tmp101721, V2435)

																									tmp101723 := Call(__e, tmp101720, tmp101722, V2989)

																									__e.TailApply(tmp101689, tmp101723)
																									return

																								}, 1)

																								tmp101724 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																								tmp101725 := Call(__e, tmp101724, V2435)

																								__e.TailApply(tmp101688, tmp101725)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp101728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																						tmp101729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																						tmp101730 := Call(__e, tmp101729, V2430)

																						tmp101731 := Call(__e, tmp101728, tmp101730, V2989)

																						__e.TailApply(tmp101686, tmp101731)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp101734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																				tmp101735 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																				tmp101736 := Call(__e, tmp101735, V2433)

																				tmp101737 := Call(__e, tmp101734, tmp101736, V2989)

																				__e.TailApply(tmp101684, tmp101737)
																				return

																			}, 1)

																			tmp101738 := Call(__e, PrimNS1Value(symns2_1value), symhd)

																			tmp101739 := Call(__e, tmp101738, V2433)

																			__e.TailApply(tmp101683, tmp101739)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp101742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

																	tmp101743 := Call(__e, PrimNS1Value(symns2_1value), symtl)

																	tmp101744 := Call(__e, tmp101743, V2431)

																	tmp101745 := Call(__e, tmp101742, tmp101744, V2989)

																	__e.TailApply(tmp101681, tmp101745)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp101748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

															tmp101749 := Call(__e, PrimNS1Value(symns2_1value), symhd)

															tmp101750 := Call(__e, tmp101749, V2431)

															tmp101751 := Call(__e, tmp101748, tmp101750, V2989)

															__e.TailApply(tmp101679, tmp101751)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp101754 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp101755 := Call(__e, PrimNS1Value(symns2_1value), symhd)

													tmp101756 := Call(__e, tmp101755, V2430)

													tmp101757 := Call(__e, tmp101754, tmp101756, V2989)

													__e.TailApply(tmp101677, tmp101757)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp101760 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

											tmp101761 := Call(__e, PrimNS1Value(symns2_1value), symhd)

											tmp101762 := Call(__e, tmp101761, V2429)

											tmp101763 := Call(__e, tmp101760, tmp101762, V2989)

											__e.TailApply(tmp101675, tmp101763)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp101766 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

									tmp101767 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp101768 := Call(__e, tmp101767, V2427)

									tmp101769 := Call(__e, tmp101766, tmp101768, V2989)

									__e.TailApply(tmp101673, tmp101769)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp101772 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp101773 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp101774 := Call(__e, tmp101773, V2427)

							tmp101775 := Call(__e, tmp101772, tmp101774, V2989)

							__e.TailApply(tmp101671, tmp101775)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp101778 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp101779 := Call(__e, tmp101778, V2986, V2989)

					tmp101780 := Call(__e, tmp101669, tmp101779)

					__e.TailApply(tmp101593, tmp101780)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp101783 := MakeNative(func(__e *ControlFlow) {
				V2422 := __e.Get(1)
				_ = V2422
				tmp101846 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp101847 := Call(__e, tmp101846, V2422)

				if True == tmp101847 {
					tmp101785 := MakeNative(func(__e *ControlFlow) {
						V2423 := __e.Get(1)
						_ = V2423
						tmp101840 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp101841 := Call(__e, tmp101840, symwhere, V2423)

						if True == tmp101841 {
							tmp101787 := MakeNative(func(__e *ControlFlow) {
								V2424 := __e.Get(1)
								_ = V2424
								tmp101834 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp101835 := Call(__e, tmp101834, V2424)

								if True == tmp101835 {
									tmp101789 := MakeNative(func(__e *ControlFlow) {
										P := __e.Get(1)
										_ = P
										tmp101790 := MakeNative(func(__e *ControlFlow) {
											V2425 := __e.Get(1)
											_ = V2425
											tmp101826 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

											tmp101827 := Call(__e, tmp101826, V2425)

											if True == tmp101827 {
												tmp101792 := MakeNative(func(__e *ControlFlow) {
													Action := __e.Get(1)
													_ = Action
													tmp101793 := MakeNative(func(__e *ControlFlow) {
														V2426 := __e.Get(1)
														_ = V2426
														tmp101818 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

														tmp101819 := Call(__e, tmp101818, Nil, V2426)

														if True == tmp101819 {
															tmp101795 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

															tmp101796 := Call(__e, tmp101795)

															_ = tmp101796

															tmp101797 := Call(__e, PrimNS1Value(symns2_1value), symcut)

															tmp101798 := MakeNative(func(__e *ControlFlow) {
																tmp101799 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d)

																tmp101800 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp101801 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp101802 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																tmp101803 := Call(__e, tmp101802, symboolean, Nil)

																tmp101804 := Call(__e, tmp101801, sym_h, tmp101803)

																tmp101805 := Call(__e, tmp101800, P, tmp101804)

																tmp101806 := MakeNative(func(__e *ControlFlow) {
																	tmp101807 := Call(__e, PrimNS1Value(symns2_1value), symcut)

																	tmp101808 := MakeNative(func(__e *ControlFlow) {
																		tmp101809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4t_d_1action)

																		tmp101810 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																		tmp101811 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																		tmp101812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																		tmp101813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

																		tmp101814 := Call(__e, tmp101813, symverified, Nil)

																		tmp101815 := Call(__e, tmp101812, sym_h, tmp101814)

																		tmp101816 := Call(__e, tmp101811, P, tmp101815)

																		tmp101817 := Call(__e, tmp101810, tmp101816, V2988)

																		__e.TailApply(tmp101809, Action, V2987, tmp101817, V2989, V2990)
																		return

																	}, 0)

																	__e.TailApply(tmp101807, Throwcontrol, V2989, tmp101808)
																	return

																}, 0)

																__e.TailApply(tmp101799, tmp101805, V2988, V2989, tmp101806)
																return

															}, 0)

															__e.TailApply(tmp101797, Throwcontrol, V2989, tmp101798)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp101820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

													tmp101821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

													tmp101822 := Call(__e, tmp101821, V2425)

													tmp101823 := Call(__e, tmp101820, tmp101822, V2989)

													__e.TailApply(tmp101793, tmp101823)
													return

												}, 1)

												tmp101824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp101825 := Call(__e, tmp101824, V2425)

												__e.TailApply(tmp101792, tmp101825)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp101828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

										tmp101829 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp101830 := Call(__e, tmp101829, V2424)

										tmp101831 := Call(__e, tmp101828, tmp101830, V2989)

										__e.TailApply(tmp101790, tmp101831)
										return

									}, 1)

									tmp101832 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp101833 := Call(__e, tmp101832, V2424)

									__e.TailApply(tmp101789, tmp101833)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp101836 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

							tmp101837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp101838 := Call(__e, tmp101837, V2422)

							tmp101839 := Call(__e, tmp101836, tmp101838, V2989)

							__e.TailApply(tmp101787, tmp101839)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp101842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp101843 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp101844 := Call(__e, tmp101843, V2422)

					tmp101845 := Call(__e, tmp101842, tmp101844, V2989)

					__e.TailApply(tmp101785, tmp101845)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp101848 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

			tmp101849 := Call(__e, tmp101848, V2986, V2989)

			tmp101850 := Call(__e, tmp101783, tmp101849)

			tmp101851 := Call(__e, tmp101591, tmp101850)

			__e.TailApply(tmp101590, Throwcontrol, tmp101851)
			return

		}, 1)

		tmp101852 := Call(__e, PrimNS1Value(symns2_1value), symshen_4catchpoint)

		tmp101853 := Call(__e, tmp101852)

		__e.TailApply(tmp101589, tmp101853)
		return

	}, 5)

	tmp101854 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1action, tmp101588)

	_ = tmp101854

	tmp101855 := MakeNative(func(__e *ControlFlow) {
		V2996 := __e.Get(1)
		_ = V2996
		V2997 := __e.Get(2)
		_ = V2997
		V2998 := __e.Get(3)
		_ = V2998
		V2999 := __e.Get(4)
		_ = V2999
		V3000 := __e.Get(5)
		_ = V3000
		tmp101856 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp101857 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				tmp101858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp101859 := Call(__e, tmp101858)

				_ = tmp101859

				tmp101860 := Call(__e, PrimNS1Value(symns2_1value), symbind)

				tmp101861 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				tmp101862 := Call(__e, tmp101861, symshen_4a)

				tmp101863 := MakeNative(func(__e *ControlFlow) {
					tmp101864 := Call(__e, PrimNS1Value(symns2_1value), symbind)

					tmp101865 := Call(__e, PrimNS1Value(symns2_1value), symset)

					tmp101866 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

					tmp101867 := Call(__e, tmp101866, A, V2999)

					tmp101868 := Call(__e, tmp101865, tmp101867, Nil)

					tmp101869 := MakeNative(func(__e *ControlFlow) {
						tmp101870 := Call(__e, PrimNS1Value(symns2_1value), symshen_4findallhelp)

						__e.TailApply(tmp101870, V2996, V2997, V2998, A, V2999, V3000)
						return

					}, 0)

					__e.TailApply(tmp101864, B, tmp101868, V2999, tmp101869)
					return

				}, 0)

				__e.TailApply(tmp101860, A, tmp101862, V2999, tmp101863)
				return

			}, 1)

			tmp101871 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

			tmp101872 := Call(__e, tmp101871, V2999)

			__e.TailApply(tmp101857, tmp101872)
			return

		}, 1)

		tmp101873 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp101874 := Call(__e, tmp101873, V2999)

		__e.TailApply(tmp101856, tmp101874)
		return

	}, 5)

	tmp101875 := Call(__e, PrimNS1Value(symns2_1set), symfindall, tmp101855)

	_ = tmp101875

	tmp101876 := MakeNative(func(__e *ControlFlow) {
		V3007 := __e.Get(1)
		_ = V3007
		V3008 := __e.Get(2)
		_ = V3008
		V3009 := __e.Get(3)
		_ = V3009
		V3010 := __e.Get(4)
		_ = V3010
		V3011 := __e.Get(5)
		_ = V3011
		V3012 := __e.Get(6)
		_ = V3012
		tmp101877 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp101886 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp101887 := Call(__e, tmp101886, Case, False)

			if True == tmp101887 {
				tmp101879 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

				tmp101880 := Call(__e, tmp101879)

				_ = tmp101880

				tmp101881 := Call(__e, PrimNS1Value(symns2_1value), symbind)

				tmp101882 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				tmp101883 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lazyderef)

				tmp101884 := Call(__e, tmp101883, V3010, V3011)

				tmp101885 := Call(__e, tmp101882, tmp101884)

				__e.TailApply(tmp101881, V3009, tmp101885, V3011, V3012)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp101888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

		tmp101889 := Call(__e, tmp101888)

		_ = tmp101889

		tmp101890 := Call(__e, PrimNS1Value(symns2_1value), symcall)

		tmp101891 := MakeNative(func(__e *ControlFlow) {
			tmp101892 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remember)

			tmp101893 := MakeNative(func(__e *ControlFlow) {
				tmp101894 := Call(__e, PrimNS1Value(symns2_1value), symfwhen)

				__e.TailApply(tmp101894, False, V3011, V3012)
				return

			}, 0)

			__e.TailApply(tmp101892, V3010, V3007, V3011, tmp101893)
			return

		}, 0)

		tmp101895 := Call(__e, tmp101890, V3008, V3011, tmp101891)

		__e.TailApply(tmp101877, tmp101895)
		return

	}, 6)

	tmp101896 := Call(__e, PrimNS1Value(symns2_1set), symshen_4findallhelp, tmp101876)

	_ = tmp101896

	tmp101897 := MakeNative(func(__e *ControlFlow) {
		V3017 := __e.Get(1)
		_ = V3017
		V3018 := __e.Get(2)
		_ = V3018
		V3019 := __e.Get(3)
		_ = V3019
		V3020 := __e.Get(4)
		_ = V3020
		tmp101898 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp101899 := Call(__e, PrimNS1Value(symns2_1value), symshen_4incinfs)

			tmp101900 := Call(__e, tmp101899)

			_ = tmp101900

			tmp101901 := Call(__e, PrimNS1Value(symns2_1value), symbind)

			tmp101902 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp101903 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp101904 := Call(__e, tmp101903, V3017, V3019)

			tmp101905 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp101906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp101907 := Call(__e, tmp101906, V3018, V3019)

			tmp101908 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp101909 := Call(__e, PrimNS1Value(symns2_1value), symshen_4deref)

			tmp101910 := Call(__e, tmp101909, V3017, V3019)

			tmp101911 := Call(__e, tmp101908, tmp101910)

			tmp101912 := Call(__e, tmp101905, tmp101907, tmp101911)

			tmp101913 := Call(__e, tmp101902, tmp101904, tmp101912)

			__e.TailApply(tmp101901, B, tmp101913, V3019, V3020)
			return

		}, 1)

		tmp101914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newpv)

		tmp101915 := Call(__e, tmp101914, V3019)

		__e.TailApply(tmp101898, tmp101915)
		return

	}, 4)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4remember, tmp101897)
	return

}, 0)
