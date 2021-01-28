package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp17961 := MakeNative(func(__e *ControlFlow) {
		V2731 := __e.Get(1)
		_ = V2731
		V2732 := __e.Get(2)
		_ = V2732
		tmp17962 := MakeNative(func(__e *ControlFlow) {
			Curry := __e.Get(1)
			_ = Curry
			tmp17963 := MakeNative(func(__e *ControlFlow) {
				ProcessN := __e.Get(1)
				_ = ProcessN
				tmp17964 := MakeNative(func(__e *ControlFlow) {
					Type := __e.Get(1)
					_ = Type
					tmp17965 := MakeNative(func(__e *ControlFlow) {
						Continuation := __e.Get(1)
						_ = Continuation
						tmp17966 := PrimCons(Type, Nil)

						tmp17967 := PrimCons(sym_h, tmp17966)

						tmp17968 := PrimCons(Curry, tmp17967)

						__e.TailApply(PrimNS2Value(symshen_4t_d), tmp17968, Nil, ProcessN, Continuation)
						return

					}, 1)

					tmp17969 := MakeNative(func(__e *ControlFlow) {
						__e.TailApply(PrimNS2Value(symreturn), Type, ProcessN, symshen_4void)
						return
					}, 0)

					__e.TailApply(tmp17965, tmp17969)
					return

				}, 1)

				tmp17970 := Call(__e, PrimNS2Value(symshen_4curry_1type), V2732)

				tmp17971 := Call(__e, PrimNS2Value(symshen_4demodulate), tmp17970)

				tmp17972 := Call(__e, PrimNS2Value(symshen_4insert_1prolog_1variables), tmp17971, ProcessN)

				__e.TailApply(tmp17964, tmp17972)
				return

			}, 1)

			tmp17973 := Call(__e, PrimNS2Value(symshen_4start_1new_1prolog_1process))

			__e.TailApply(tmp17963, tmp17973)
			return

		}, 1)

		tmp17974 := Call(__e, PrimNS2Value(symshen_4curry), V2731)

		__e.TailApply(tmp17962, tmp17974)
		return

	}, 2)

	tmp17975 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck, tmp17961)

	_ = tmp17975

	tmp17976 := MakeNative(func(__e *ControlFlow) {
		V2734 := __e.Get(1)
		_ = V2734
		tmp18059 := PrimIsPair(V2734)

		var ifres18055 Obj

		if True == tmp18059 {
			tmp18057 := PrimHead(V2734)

			tmp18058 := Call(__e, PrimNS2Value(symshen_4special_2), tmp18057)

			var ifres18056 Obj

			if True == tmp18058 {
				ifres18056 = True

			} else {
				ifres18056 = False

			}

			ifres18055 = ifres18056

		} else {
			ifres18055 = False

		}

		if True == ifres18055 {
			tmp18051 := PrimHead(V2734)

			tmp18052 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4curry), Y)
				return
			}, 1)

			tmp18053 := PrimTail(V2734)

			tmp18054 := Call(__e, PrimNS2Value(symmap), tmp18052, tmp18053)

			__e.Return(PrimCons(tmp18051, tmp18054))
			return

		} else {
			tmp18050 := PrimIsPair(V2734)

			var ifres18042 Obj

			if True == tmp18050 {
				tmp18048 := PrimTail(V2734)

				tmp18049 := PrimIsPair(tmp18048)

				var ifres18044 Obj

				if True == tmp18049 {
					tmp18046 := PrimHead(V2734)

					tmp18047 := Call(__e, PrimNS2Value(symshen_4extraspecial_2), tmp18046)

					var ifres18045 Obj

					if True == tmp18047 {
						ifres18045 = True

					} else {
						ifres18045 = False

					}

					ifres18044 = ifres18045

				} else {
					ifres18044 = False

				}

				var ifres18043 Obj

				if True == ifres18044 {
					ifres18043 = True

				} else {
					ifres18043 = False

				}

				ifres18042 = ifres18043

			} else {
				ifres18042 = False

			}

			if True == ifres18042 {
				__e.Return(V2734)
				return
			} else {
				tmp18041 := PrimIsPair(V2734)

				var ifres18022 Obj

				if True == tmp18041 {
					tmp18039 := PrimHead(V2734)

					tmp18040 := PrimEqual(symtype, tmp18039)

					var ifres18024 Obj

					if True == tmp18040 {
						tmp18037 := PrimTail(V2734)

						tmp18038 := PrimIsPair(tmp18037)

						var ifres18026 Obj

						if True == tmp18038 {
							tmp18034 := PrimTail(V2734)

							tmp18035 := PrimTail(tmp18034)

							tmp18036 := PrimIsPair(tmp18035)

							var ifres18028 Obj

							if True == tmp18036 {
								tmp18030 := PrimTail(V2734)

								tmp18031 := PrimTail(tmp18030)

								tmp18032 := PrimTail(tmp18031)

								tmp18033 := PrimEqual(Nil, tmp18032)

								var ifres18029 Obj

								if True == tmp18033 {
									ifres18029 = True

								} else {
									ifres18029 = False

								}

								ifres18028 = ifres18029

							} else {
								ifres18028 = False

							}

							var ifres18027 Obj

							if True == ifres18028 {
								ifres18027 = True

							} else {
								ifres18027 = False

							}

							ifres18026 = ifres18027

						} else {
							ifres18026 = False

						}

						var ifres18025 Obj

						if True == ifres18026 {
							ifres18025 = True

						} else {
							ifres18025 = False

						}

						ifres18024 = ifres18025

					} else {
						ifres18024 = False

					}

					var ifres18023 Obj

					if True == ifres18024 {
						ifres18023 = True

					} else {
						ifres18023 = False

					}

					ifres18022 = ifres18023

				} else {
					ifres18022 = False

				}

				if True == ifres18022 {
					tmp18016 := PrimTail(V2734)

					tmp18017 := PrimHead(tmp18016)

					tmp18018 := Call(__e, PrimNS2Value(symshen_4curry), tmp18017)

					tmp18019 := PrimTail(V2734)

					tmp18020 := PrimTail(tmp18019)

					tmp18021 := PrimCons(tmp18018, tmp18020)

					__e.Return(PrimCons(symtype, tmp18021))
					return

				} else {
					tmp18015 := PrimIsPair(V2734)

					var ifres18006 Obj

					if True == tmp18015 {
						tmp18013 := PrimTail(V2734)

						tmp18014 := PrimIsPair(tmp18013)

						var ifres18008 Obj

						if True == tmp18014 {
							tmp18010 := PrimTail(V2734)

							tmp18011 := PrimTail(tmp18010)

							tmp18012 := PrimIsPair(tmp18011)

							var ifres18009 Obj

							if True == tmp18012 {
								ifres18009 = True

							} else {
								ifres18009 = False

							}

							ifres18008 = ifres18009

						} else {
							ifres18008 = False

						}

						var ifres18007 Obj

						if True == ifres18008 {
							ifres18007 = True

						} else {
							ifres18007 = False

						}

						ifres18006 = ifres18007

					} else {
						ifres18006 = False

					}

					if True == ifres18006 {
						tmp17998 := PrimHead(V2734)

						tmp17999 := PrimTail(V2734)

						tmp18000 := PrimHead(tmp17999)

						tmp18001 := PrimCons(tmp18000, Nil)

						tmp18002 := PrimCons(tmp17998, tmp18001)

						tmp18003 := PrimTail(V2734)

						tmp18004 := PrimTail(tmp18003)

						tmp18005 := PrimCons(tmp18002, tmp18004)

						__e.TailApply(PrimNS2Value(symshen_4curry), tmp18005)
						return

					} else {
						tmp17997 := PrimIsPair(V2734)

						var ifres17988 Obj

						if True == tmp17997 {
							tmp17995 := PrimTail(V2734)

							tmp17996 := PrimIsPair(tmp17995)

							var ifres17990 Obj

							if True == tmp17996 {
								tmp17992 := PrimTail(V2734)

								tmp17993 := PrimTail(tmp17992)

								tmp17994 := PrimEqual(Nil, tmp17993)

								var ifres17991 Obj

								if True == tmp17994 {
									ifres17991 = True

								} else {
									ifres17991 = False

								}

								ifres17990 = ifres17991

							} else {
								ifres17990 = False

							}

							var ifres17989 Obj

							if True == ifres17990 {
								ifres17989 = True

							} else {
								ifres17989 = False

							}

							ifres17988 = ifres17989

						} else {
							ifres17988 = False

						}

						if True == ifres17988 {
							tmp17982 := PrimHead(V2734)

							tmp17983 := Call(__e, PrimNS2Value(symshen_4curry), tmp17982)

							tmp17984 := PrimTail(V2734)

							tmp17985 := PrimHead(tmp17984)

							tmp17986 := Call(__e, PrimNS2Value(symshen_4curry), tmp17985)

							tmp17987 := PrimCons(tmp17986, Nil)

							__e.Return(PrimCons(tmp17983, tmp17987))
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

	tmp18060 := Call(__e, PrimNS1Value(symns2_1set), symshen_4curry, tmp17976)

	_ = tmp18060

	tmp18061 := MakeNative(func(__e *ControlFlow) {
		V2736 := __e.Get(1)
		_ = V2736
		tmp18062 := PrimNS3Value(symshen_4_dspecial_d)

		__e.TailApply(PrimNS2Value(symelement_2), V2736, tmp18062)
		return

	}, 1)

	tmp18063 := Call(__e, PrimNS1Value(symns2_1set), symshen_4special_2, tmp18061)

	_ = tmp18063

	tmp18064 := MakeNative(func(__e *ControlFlow) {
		V2738 := __e.Get(1)
		_ = V2738
		tmp18065 := PrimNS3Value(symshen_4_dextraspecial_d)

		__e.TailApply(PrimNS2Value(symelement_2), V2738, tmp18065)
		return

	}, 1)

	tmp18066 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extraspecial_2, tmp18064)

	_ = tmp18066

	tmp18067 := MakeNative(func(__e *ControlFlow) {
		V2743 := __e.Get(1)
		_ = V2743
		V2744 := __e.Get(2)
		_ = V2744
		V2745 := __e.Get(3)
		_ = V2745
		V2746 := __e.Get(4)
		_ = V2746
		tmp18068 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp18069 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp18123 := PrimEqual(Case, False)

				if True == tmp18123 {
					tmp18071 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp18115 := PrimEqual(Case, False)

						if True == tmp18115 {
							tmp18073 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp18081 := PrimEqual(Case, False)

								if True == tmp18081 {
									tmp18075 := MakeNative(func(__e *ControlFlow) {
										Datatypes := __e.Get(1)
										_ = Datatypes
										tmp18076 := Call(__e, PrimNS2Value(symshen_4incinfs))

										_ = tmp18076

										tmp18077 := MakeNative(func(__e *ControlFlow) {
											tmp18078 := PrimNS3Value(symshen_4_ddatatypes_d)

											tmp18079 := MakeNative(func(__e *ControlFlow) {
												__e.TailApply(PrimNS2Value(symshen_4udefs_d), V2743, V2744, Datatypes, V2745, V2746)
												return
											}, 0)

											__e.TailApply(PrimNS2Value(symbind), Datatypes, tmp18078, V2745, tmp18079)
											return

										}, 0)

										__e.TailApply(PrimNS2Value(symshen_4show), V2743, V2744, V2745, tmp18077)
										return

									}, 1)

									tmp18080 := Call(__e, PrimNS2Value(symshen_4newpv), V2745)

									__e.TailApply(tmp18075, tmp18080)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp18082 := MakeNative(func(__e *ControlFlow) {
								V2724 := __e.Get(1)
								_ = V2724
								tmp18112 := PrimIsPair(V2724)

								if True == tmp18112 {
									tmp18084 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp18085 := MakeNative(func(__e *ControlFlow) {
											V2725 := __e.Get(1)
											_ = V2725
											tmp18108 := PrimIsPair(V2725)

											if True == tmp18108 {
												tmp18087 := MakeNative(func(__e *ControlFlow) {
													V2726 := __e.Get(1)
													_ = V2726
													tmp18105 := PrimEqual(sym_h, V2726)

													if True == tmp18105 {
														tmp18089 := MakeNative(func(__e *ControlFlow) {
															V2727 := __e.Get(1)
															_ = V2727
															tmp18102 := PrimIsPair(V2727)

															if True == tmp18102 {
																tmp18091 := MakeNative(func(__e *ControlFlow) {
																	A := __e.Get(1)
																	_ = A
																	tmp18092 := MakeNative(func(__e *ControlFlow) {
																		V2728 := __e.Get(1)
																		_ = V2728
																		tmp18098 := PrimEqual(Nil, V2728)

																		if True == tmp18098 {
																			tmp18094 := Call(__e, PrimNS2Value(symshen_4incinfs))

																			_ = tmp18094

																			tmp18095 := Call(__e, PrimNS2Value(symshen_4type_1theory_1enabled_2))

																			tmp18096 := MakeNative(func(__e *ControlFlow) {
																				tmp18097 := MakeNative(func(__e *ControlFlow) {
																					__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2744, V2745, V2746)
																					return
																				}, 0)

																				__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2745, tmp18097)
																				return

																			}, 0)

																			__e.TailApply(PrimNS2Value(symfwhen), tmp18095, V2745, tmp18096)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp18099 := PrimTail(V2727)

																	tmp18100 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18099, V2745)

																	__e.TailApply(tmp18092, tmp18100)
																	return

																}, 1)

																tmp18101 := PrimHead(V2727)

																__e.TailApply(tmp18091, tmp18101)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp18103 := PrimTail(V2725)

														tmp18104 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18103, V2745)

														__e.TailApply(tmp18089, tmp18104)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp18106 := PrimHead(V2725)

												tmp18107 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18106, V2745)

												__e.TailApply(tmp18087, tmp18107)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp18109 := PrimTail(V2724)

										tmp18110 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18109, V2745)

										__e.TailApply(tmp18085, tmp18110)
										return

									}, 1)

									tmp18111 := PrimHead(V2724)

									__e.TailApply(tmp18084, tmp18111)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp18113 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2743, V2745)

							tmp18114 := Call(__e, tmp18082, tmp18113)

							__e.TailApply(tmp18073, tmp18114)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp18116 := MakeNative(func(__e *ControlFlow) {
						V2723 := __e.Get(1)
						_ = V2723
						tmp18120 := PrimEqual(symfail, V2723)

						if True == tmp18120 {
							tmp18118 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp18118

							tmp18119 := MakeNative(func(__e *ControlFlow) {
								__e.TailApply(PrimNS2Value(symshen_4prolog_1failure), V2745, V2746)
								return
							}, 0)

							__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2745, tmp18119)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp18121 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2743, V2745)

					tmp18122 := Call(__e, tmp18116, tmp18121)

					__e.TailApply(tmp18071, tmp18122)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp18124 := MakeNative(func(__e *ControlFlow) {
				Error := __e.Get(1)
				_ = Error
				tmp18125 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp18125

				tmp18126 := Call(__e, PrimNS2Value(symshen_4maxinfexceeded_2))

				tmp18127 := MakeNative(func(__e *ControlFlow) {
					tmp18128 := Call(__e, PrimNS2Value(symshen_4errormaxinfs))

					__e.TailApply(PrimNS2Value(symbind), Error, tmp18128, V2745, V2746)
					return

				}, 0)

				__e.TailApply(PrimNS2Value(symfwhen), tmp18126, V2745, tmp18127)
				return

			}, 1)

			tmp18129 := Call(__e, PrimNS2Value(symshen_4newpv), V2745)

			tmp18130 := Call(__e, tmp18124, tmp18129)

			tmp18131 := Call(__e, tmp18069, tmp18130)

			__e.TailApply(PrimNS2Value(symshen_4cutpoint), Throwcontrol, tmp18131)
			return

		}, 1)

		tmp18132 := Call(__e, PrimNS2Value(symshen_4catchpoint))

		__e.TailApply(tmp18068, tmp18132)
		return

	}, 4)

	tmp18133 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d, tmp18067)

	_ = tmp18133

	tmp18134 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dshen_1type_1theory_1enabled_2_d))
		return
	}, 0)

	tmp18135 := Call(__e, PrimNS1Value(symns2_1set), symshen_4type_1theory_1enabled_2, tmp18134)

	_ = tmp18135

	tmp18136 := MakeNative(func(__e *ControlFlow) {
		V2752 := __e.Get(1)
		_ = V2752
		tmp18140 := PrimEqual(sym_7, V2752)

		if True == tmp18140 {
			__e.Return(PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
			return
		} else {
			tmp18139 := PrimEqual(sym_1, V2752)

			if True == tmp18139 {
				__e.Return(PrimNS3Set(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
				return
			}

		}

	}, 1)

	tmp18141 := Call(__e, PrimNS1Value(symns2_1set), symenable_1type_1theory, tmp18136)

	_ = tmp18141

	tmp18142 := MakeNative(func(__e *ControlFlow) {
		V2763 := __e.Get(1)
		_ = V2763
		V2764 := __e.Get(2)
		_ = V2764
		__e.Return(False)
		return
	}, 2)

	tmp18143 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prolog_1failure, tmp18142)

	_ = tmp18143

	tmp18144 := MakeNative(func(__e *ControlFlow) {
		tmp18145 := Call(__e, PrimNS2Value(syminferences))

		tmp18146 := PrimNS3Value(symshen_4_dmaxinferences_d)

		__e.Return(PrimGreatThan(tmp18145, tmp18146))
		return

	}, 0)

	tmp18147 := Call(__e, PrimNS1Value(symns2_1set), symshen_4maxinfexceeded_2, tmp18144)

	_ = tmp18147

	tmp18148 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded~%")))
		return
	}, 0)

	tmp18149 := Call(__e, PrimNS1Value(symns2_1set), symshen_4errormaxinfs, tmp18148)

	_ = tmp18149

	tmp18150 := MakeNative(func(__e *ControlFlow) {
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
		tmp18151 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp18160 := PrimEqual(Case, False)

			if True == tmp18160 {
				tmp18153 := MakeNative(func(__e *ControlFlow) {
					V2720 := __e.Get(1)
					_ = V2720
					tmp18158 := PrimIsPair(V2720)

					if True == tmp18158 {
						tmp18155 := MakeNative(func(__e *ControlFlow) {
							Ds := __e.Get(1)
							_ = Ds
							tmp18156 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp18156

							__e.TailApply(PrimNS2Value(symshen_4udefs_d), V2770, V2771, Ds, V2773, V2774)
							return

						}, 1)

						tmp18157 := PrimTail(V2720)

						__e.TailApply(tmp18155, tmp18157)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp18159 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2772, V2773)

				__e.TailApply(tmp18153, tmp18159)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp18161 := MakeNative(func(__e *ControlFlow) {
			V2719 := __e.Get(1)
			_ = V2719
			tmp18169 := PrimIsPair(V2719)

			if True == tmp18169 {
				tmp18163 := MakeNative(func(__e *ControlFlow) {
					D := __e.Get(1)
					_ = D
					tmp18164 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp18164

					tmp18165 := PrimCons(V2771, Nil)

					tmp18166 := PrimCons(V2770, tmp18165)

					tmp18167 := PrimCons(D, tmp18166)

					__e.TailApply(PrimNS2Value(symcall), tmp18167, V2773, V2774)
					return

				}, 1)

				tmp18168 := PrimHead(V2719)

				__e.TailApply(tmp18163, tmp18168)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp18170 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2772, V2773)

		tmp18171 := Call(__e, tmp18161, tmp18170)

		__e.TailApply(tmp18151, tmp18171)
		return

	}, 5)

	tmp18172 := Call(__e, PrimNS1Value(symns2_1set), symshen_4udefs_d, tmp18150)

	_ = tmp18172

	tmp18173 := MakeNative(func(__e *ControlFlow) {
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
		tmp18174 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp18175 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp19436 := PrimEqual(Case, False)

				if True == tmp19436 {
					tmp18177 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp19423 := PrimEqual(Case, False)

						if True == tmp19423 {
							tmp18179 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp19420 := PrimEqual(Case, False)

								if True == tmp19420 {
									tmp18181 := MakeNative(func(__e *ControlFlow) {
										Case := __e.Get(1)
										_ = Case
										tmp19417 := PrimEqual(Case, False)

										if True == tmp19417 {
											tmp18183 := MakeNative(func(__e *ControlFlow) {
												Case := __e.Get(1)
												_ = Case
												tmp19401 := PrimEqual(Case, False)

												if True == tmp19401 {
													tmp18185 := MakeNative(func(__e *ControlFlow) {
														Case := __e.Get(1)
														_ = Case
														tmp19374 := PrimEqual(Case, False)

														if True == tmp19374 {
															tmp18187 := MakeNative(func(__e *ControlFlow) {
																Case := __e.Get(1)
																_ = Case
																tmp19237 := PrimEqual(Case, False)

																if True == tmp19237 {
																	tmp18189 := MakeNative(func(__e *ControlFlow) {
																		Case := __e.Get(1)
																		_ = Case
																		tmp19092 := PrimEqual(Case, False)

																		if True == tmp19092 {
																			tmp18191 := MakeNative(func(__e *ControlFlow) {
																				Case := __e.Get(1)
																				_ = Case
																				tmp18955 := PrimEqual(Case, False)

																				if True == tmp18955 {
																					tmp18193 := MakeNative(func(__e *ControlFlow) {
																						Case := __e.Get(1)
																						_ = Case
																						tmp18911 := PrimEqual(Case, False)

																						if True == tmp18911 {
																							tmp18195 := MakeNative(func(__e *ControlFlow) {
																								Case := __e.Get(1)
																								_ = Case
																								tmp18654 := PrimEqual(Case, False)

																								if True == tmp18654 {
																									tmp18197 := MakeNative(func(__e *ControlFlow) {
																										Case := __e.Get(1)
																										_ = Case
																										tmp18598 := PrimEqual(Case, False)

																										if True == tmp18598 {
																											tmp18199 := MakeNative(func(__e *ControlFlow) {
																												Case := __e.Get(1)
																												_ = Case
																												tmp18433 := PrimEqual(Case, False)

																												if True == tmp18433 {
																													tmp18201 := MakeNative(func(__e *ControlFlow) {
																														Case := __e.Get(1)
																														_ = Case
																														tmp18400 := PrimEqual(Case, False)

																														if True == tmp18400 {
																															tmp18203 := MakeNative(func(__e *ControlFlow) {
																																Case := __e.Get(1)
																																_ = Case
																																tmp18361 := PrimEqual(Case, False)

																																if True == tmp18361 {
																																	tmp18205 := MakeNative(func(__e *ControlFlow) {
																																		Case := __e.Get(1)
																																		_ = Case
																																		tmp18324 := PrimEqual(Case, False)

																																		if True == tmp18324 {
																																			tmp18207 := MakeNative(func(__e *ControlFlow) {
																																				Case := __e.Get(1)
																																				_ = Case
																																				tmp18318 := PrimEqual(Case, False)

																																				if True == tmp18318 {
																																					tmp18209 := MakeNative(func(__e *ControlFlow) {
																																						Case := __e.Get(1)
																																						_ = Case
																																						tmp18294 := PrimEqual(Case, False)

																																						if True == tmp18294 {
																																							tmp18211 := MakeNative(func(__e *ControlFlow) {
																																								Case := __e.Get(1)
																																								_ = Case
																																								tmp18271 := PrimEqual(Case, False)

																																								if True == tmp18271 {
																																									tmp18213 := MakeNative(func(__e *ControlFlow) {
																																										Case := __e.Get(1)
																																										_ = Case
																																										tmp18248 := PrimEqual(Case, False)

																																										if True == tmp18248 {
																																											tmp18215 := MakeNative(func(__e *ControlFlow) {
																																												Case := __e.Get(1)
																																												_ = Case
																																												tmp18225 := PrimEqual(Case, False)

																																												if True == tmp18225 {
																																													tmp18217 := MakeNative(func(__e *ControlFlow) {
																																														Datatypes := __e.Get(1)
																																														_ = Datatypes
																																														tmp18218 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																														_ = tmp18218

																																														tmp18219 := PrimNS3Value(symshen_4_ddatatypes_d)

																																														tmp18220 := MakeNative(func(__e *ControlFlow) {
																																															tmp18221 := PrimCons(V2781, Nil)

																																															tmp18222 := PrimCons(sym_h, tmp18221)

																																															tmp18223 := PrimCons(V2780, tmp18222)

																																															__e.TailApply(PrimNS2Value(symshen_4udefs_d), tmp18223, V2782, Datatypes, V2783, V2784)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimNS2Value(symbind), Datatypes, tmp18219, V2783, tmp18220)
																																														return

																																													}, 1)

																																													tmp18224 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																													__e.TailApply(tmp18217, tmp18224)
																																													return

																																												} else {
																																													__e.Return(Case)
																																													return
																																												}

																																											}, 1)

																																											tmp18226 := MakeNative(func(__e *ControlFlow) {
																																												V2713 := __e.Get(1)
																																												_ = V2713
																																												tmp18245 := PrimIsPair(V2713)

																																												if True == tmp18245 {
																																													tmp18228 := MakeNative(func(__e *ControlFlow) {
																																														V2714 := __e.Get(1)
																																														_ = V2714
																																														tmp18242 := PrimEqual(symshen_4synonyms_1help, V2714)

																																														if True == tmp18242 {
																																															tmp18230 := MakeNative(func(__e *ControlFlow) {
																																																V2715 := __e.Get(1)
																																																_ = V2715
																																																tmp18240 := PrimEqual(symsymbol, V2715)

																																																if True == tmp18240 {
																																																	tmp18239 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																	_ = tmp18239

																																																	__e.TailApply(PrimNS2Value(symthaw), V2784)
																																																	return

																																																} else {
																																																	tmp18238 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2715)

																																																	if True == tmp18238 {
																																																		tmp18233 := Call(__e, PrimNS2Value(symshen_4bindv), V2715, symsymbol, V2783)

																																																		_ = tmp18233

																																																		tmp18234 := MakeNative(func(__e *ControlFlow) {
																																																			Result := __e.Get(1)
																																																			_ = Result
																																																			tmp18235 := Call(__e, PrimNS2Value(symshen_4unbindv), V2715, V2783)

																																																			_ = tmp18235

																																																			__e.Return(Result)
																																																			return

																																																		}, 1)

																																																		tmp18236 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																		_ = tmp18236

																																																		tmp18237 := Call(__e, PrimNS2Value(symthaw), V2784)

																																																		__e.TailApply(tmp18234, tmp18237)
																																																		return

																																																	} else {
																																																		__e.Return(False)
																																																		return
																																																	}

																																																}

																																															}, 1)

																																															tmp18241 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																															__e.TailApply(tmp18230, tmp18241)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}, 1)

																																													tmp18243 := PrimHead(V2713)

																																													tmp18244 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18243, V2783)

																																													__e.TailApply(tmp18228, tmp18244)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp18246 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																																											tmp18247 := Call(__e, tmp18226, tmp18246)

																																											__e.TailApply(tmp18215, tmp18247)
																																											return

																																										} else {
																																											__e.Return(Case)
																																											return
																																										}

																																									}, 1)

																																									tmp18249 := MakeNative(func(__e *ControlFlow) {
																																										V2710 := __e.Get(1)
																																										_ = V2710
																																										tmp18268 := PrimIsPair(V2710)

																																										if True == tmp18268 {
																																											tmp18251 := MakeNative(func(__e *ControlFlow) {
																																												V2711 := __e.Get(1)
																																												_ = V2711
																																												tmp18265 := PrimEqual(symshen_4process_1datatype, V2711)

																																												if True == tmp18265 {
																																													tmp18253 := MakeNative(func(__e *ControlFlow) {
																																														V2712 := __e.Get(1)
																																														_ = V2712
																																														tmp18263 := PrimEqual(symsymbol, V2712)

																																														if True == tmp18263 {
																																															tmp18262 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																															_ = tmp18262

																																															__e.TailApply(PrimNS2Value(symthaw), V2784)
																																															return

																																														} else {
																																															tmp18261 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2712)

																																															if True == tmp18261 {
																																																tmp18256 := Call(__e, PrimNS2Value(symshen_4bindv), V2712, symsymbol, V2783)

																																																_ = tmp18256

																																																tmp18257 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp18258 := Call(__e, PrimNS2Value(symshen_4unbindv), V2712, V2783)

																																																	_ = tmp18258

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp18259 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																_ = tmp18259

																																																tmp18260 := Call(__e, PrimNS2Value(symthaw), V2784)

																																																__e.TailApply(tmp18257, tmp18260)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp18264 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																													__e.TailApply(tmp18253, tmp18264)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp18266 := PrimHead(V2710)

																																											tmp18267 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18266, V2783)

																																											__e.TailApply(tmp18251, tmp18267)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp18269 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																																									tmp18270 := Call(__e, tmp18249, tmp18269)

																																									__e.TailApply(tmp18213, tmp18270)
																																									return

																																								} else {
																																									__e.Return(Case)
																																									return
																																								}

																																							}, 1)

																																							tmp18272 := MakeNative(func(__e *ControlFlow) {
																																								V2707 := __e.Get(1)
																																								_ = V2707
																																								tmp18291 := PrimIsPair(V2707)

																																								if True == tmp18291 {
																																									tmp18274 := MakeNative(func(__e *ControlFlow) {
																																										V2708 := __e.Get(1)
																																										_ = V2708
																																										tmp18288 := PrimEqual(symdefmacro, V2708)

																																										if True == tmp18288 {
																																											tmp18276 := MakeNative(func(__e *ControlFlow) {
																																												V2709 := __e.Get(1)
																																												_ = V2709
																																												tmp18286 := PrimEqual(symunit, V2709)

																																												if True == tmp18286 {
																																													tmp18285 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																													_ = tmp18285

																																													__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, V2784)
																																													return

																																												} else {
																																													tmp18284 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2709)

																																													if True == tmp18284 {
																																														tmp18279 := Call(__e, PrimNS2Value(symshen_4bindv), V2709, symunit, V2783)

																																														_ = tmp18279

																																														tmp18280 := MakeNative(func(__e *ControlFlow) {
																																															Result := __e.Get(1)
																																															_ = Result
																																															tmp18281 := Call(__e, PrimNS2Value(symshen_4unbindv), V2709, V2783)

																																															_ = tmp18281

																																															__e.Return(Result)
																																															return

																																														}, 1)

																																														tmp18282 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																														_ = tmp18282

																																														tmp18283 := Call(__e, PrimNS2Value(symcut), Throwcontrol, V2783, V2784)

																																														__e.TailApply(tmp18280, tmp18283)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp18287 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																											__e.TailApply(tmp18276, tmp18287)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp18289 := PrimHead(V2707)

																																									tmp18290 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18289, V2783)

																																									__e.TailApply(tmp18274, tmp18290)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp18292 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																																							tmp18293 := Call(__e, tmp18272, tmp18292)

																																							__e.TailApply(tmp18211, tmp18293)
																																							return

																																						} else {
																																							__e.Return(Case)
																																							return
																																						}

																																					}, 1)

																																					tmp18295 := MakeNative(func(__e *ControlFlow) {
																																						V2704 := __e.Get(1)
																																						_ = V2704
																																						tmp18315 := PrimIsPair(V2704)

																																						if True == tmp18315 {
																																							tmp18297 := MakeNative(func(__e *ControlFlow) {
																																								V2705 := __e.Get(1)
																																								_ = V2705
																																								tmp18312 := PrimEqual(symdefine, V2705)

																																								if True == tmp18312 {
																																									tmp18299 := MakeNative(func(__e *ControlFlow) {
																																										V2706 := __e.Get(1)
																																										_ = V2706
																																										tmp18309 := PrimIsPair(V2706)

																																										if True == tmp18309 {
																																											tmp18301 := MakeNative(func(__e *ControlFlow) {
																																												F := __e.Get(1)
																																												_ = F
																																												tmp18302 := MakeNative(func(__e *ControlFlow) {
																																													X := __e.Get(1)
																																													_ = X
																																													tmp18303 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																													_ = tmp18303

																																													tmp18304 := MakeNative(func(__e *ControlFlow) {
																																														tmp18305 := PrimCons(F, X)

																																														tmp18306 := PrimCons(symdefine, tmp18305)

																																														__e.TailApply(PrimNS2Value(symshen_4t_d_1def), tmp18306, V2781, V2782, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18304)
																																													return

																																												}, 1)

																																												tmp18307 := PrimTail(V2706)

																																												__e.TailApply(tmp18302, tmp18307)
																																												return

																																											}, 1)

																																											tmp18308 := PrimHead(V2706)

																																											__e.TailApply(tmp18301, tmp18308)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp18310 := PrimTail(V2704)

																																									tmp18311 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18310, V2783)

																																									__e.TailApply(tmp18299, tmp18311)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp18313 := PrimHead(V2704)

																																							tmp18314 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18313, V2783)

																																							__e.TailApply(tmp18297, tmp18314)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp18316 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																																					tmp18317 := Call(__e, tmp18295, tmp18316)

																																					__e.TailApply(tmp18209, tmp18317)
																																					return

																																				} else {
																																					__e.Return(Case)
																																					return
																																				}

																																			}, 1)

																																			tmp18319 := MakeNative(func(__e *ControlFlow) {
																																				NewHyp := __e.Get(1)
																																				_ = NewHyp
																																				tmp18320 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp18320

																																				tmp18321 := MakeNative(func(__e *ControlFlow) {
																																					__e.TailApply(PrimNS2Value(symshen_4th_d), V2780, V2781, NewHyp, V2783, V2784)
																																					return
																																				}, 0)

																																				__e.TailApply(PrimNS2Value(symshen_4t_d_1hyps), V2782, NewHyp, V2783, tmp18321)
																																				return

																																			}, 1)

																																			tmp18322 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																			tmp18323 := Call(__e, tmp18319, tmp18322)

																																			__e.TailApply(tmp18207, tmp18323)
																																			return

																																		} else {
																																			__e.Return(Case)
																																			return
																																		}

																																	}, 1)

																																	tmp18325 := MakeNative(func(__e *ControlFlow) {
																																		V2699 := __e.Get(1)
																																		_ = V2699
																																		tmp18358 := PrimIsPair(V2699)

																																		if True == tmp18358 {
																																			tmp18327 := MakeNative(func(__e *ControlFlow) {
																																				V2700 := __e.Get(1)
																																				_ = V2700
																																				tmp18355 := PrimEqual(symset, V2700)

																																				if True == tmp18355 {
																																					tmp18329 := MakeNative(func(__e *ControlFlow) {
																																						V2701 := __e.Get(1)
																																						_ = V2701
																																						tmp18352 := PrimIsPair(V2701)

																																						if True == tmp18352 {
																																							tmp18331 := MakeNative(func(__e *ControlFlow) {
																																								Var := __e.Get(1)
																																								_ = Var
																																								tmp18332 := MakeNative(func(__e *ControlFlow) {
																																									V2702 := __e.Get(1)
																																									_ = V2702
																																									tmp18348 := PrimIsPair(V2702)

																																									if True == tmp18348 {
																																										tmp18334 := MakeNative(func(__e *ControlFlow) {
																																											Val := __e.Get(1)
																																											_ = Val
																																											tmp18335 := MakeNative(func(__e *ControlFlow) {
																																												V2703 := __e.Get(1)
																																												_ = V2703
																																												tmp18344 := PrimEqual(Nil, V2703)

																																												if True == tmp18344 {
																																													tmp18337 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																													_ = tmp18337

																																													tmp18338 := MakeNative(func(__e *ControlFlow) {
																																														tmp18339 := MakeNative(func(__e *ControlFlow) {
																																															tmp18340 := MakeNative(func(__e *ControlFlow) {
																																																tmp18341 := PrimCons(Var, Nil)

																																																tmp18342 := PrimCons(symvalue, tmp18341)

																																																tmp18343 := MakeNative(func(__e *ControlFlow) {
																																																	__e.TailApply(PrimNS2Value(symshen_4th_d), Val, V2781, V2782, V2783, V2784)
																																																	return
																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symshen_4th_d), tmp18342, V2781, V2782, V2783, tmp18343)
																																																return

																																															}, 0)

																																															__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18340)
																																															return

																																														}, 0)

																																														__e.TailApply(PrimNS2Value(symshen_4th_d), Var, symsymbol, V2782, V2783, tmp18339)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18338)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}, 1)

																																											tmp18345 := PrimTail(V2702)

																																											tmp18346 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18345, V2783)

																																											__e.TailApply(tmp18335, tmp18346)
																																											return

																																										}, 1)

																																										tmp18347 := PrimHead(V2702)

																																										__e.TailApply(tmp18334, tmp18347)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}, 1)

																																								tmp18349 := PrimTail(V2701)

																																								tmp18350 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18349, V2783)

																																								__e.TailApply(tmp18332, tmp18350)
																																								return

																																							}, 1)

																																							tmp18351 := PrimHead(V2701)

																																							__e.TailApply(tmp18331, tmp18351)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp18353 := PrimTail(V2699)

																																					tmp18354 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18353, V2783)

																																					__e.TailApply(tmp18329, tmp18354)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp18356 := PrimHead(V2699)

																																			tmp18357 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18356, V2783)

																																			__e.TailApply(tmp18327, tmp18357)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp18359 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																																	tmp18360 := Call(__e, tmp18325, tmp18359)

																																	__e.TailApply(tmp18205, tmp18360)
																																	return

																																} else {
																																	__e.Return(Case)
																																	return
																																}

																															}, 1)

																															tmp18362 := MakeNative(func(__e *ControlFlow) {
																																V2694 := __e.Get(1)
																																_ = V2694
																																tmp18397 := PrimIsPair(V2694)

																																if True == tmp18397 {
																																	tmp18364 := MakeNative(func(__e *ControlFlow) {
																																		V2695 := __e.Get(1)
																																		_ = V2695
																																		tmp18394 := PrimEqual(syminput_7, V2695)

																																		if True == tmp18394 {
																																			tmp18366 := MakeNative(func(__e *ControlFlow) {
																																				V2696 := __e.Get(1)
																																				_ = V2696
																																				tmp18391 := PrimIsPair(V2696)

																																				if True == tmp18391 {
																																					tmp18368 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp18369 := MakeNative(func(__e *ControlFlow) {
																																							V2697 := __e.Get(1)
																																							_ = V2697
																																							tmp18387 := PrimIsPair(V2697)

																																							if True == tmp18387 {
																																								tmp18371 := MakeNative(func(__e *ControlFlow) {
																																									Stream := __e.Get(1)
																																									_ = Stream
																																									tmp18372 := MakeNative(func(__e *ControlFlow) {
																																										V2698 := __e.Get(1)
																																										_ = V2698
																																										tmp18383 := PrimEqual(Nil, V2698)

																																										if True == tmp18383 {
																																											tmp18374 := MakeNative(func(__e *ControlFlow) {
																																												C := __e.Get(1)
																																												_ = C
																																												tmp18375 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp18375

																																												tmp18376 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2783)

																																												tmp18377 := Call(__e, PrimNS2Value(symshen_4demodulate), tmp18376)

																																												tmp18378 := MakeNative(func(__e *ControlFlow) {
																																													tmp18379 := MakeNative(func(__e *ControlFlow) {
																																														tmp18380 := PrimCons(symin, Nil)

																																														tmp18381 := PrimCons(symstream, tmp18380)

																																														__e.TailApply(PrimNS2Value(symshen_4th_d), Stream, tmp18381, V2782, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symunify), V2781, C, V2783, tmp18379)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symbind), C, tmp18377, V2783, tmp18378)
																																												return

																																											}, 1)

																																											tmp18382 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																											__e.TailApply(tmp18374, tmp18382)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}, 1)

																																									tmp18384 := PrimTail(V2697)

																																									tmp18385 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18384, V2783)

																																									__e.TailApply(tmp18372, tmp18385)
																																									return

																																								}, 1)

																																								tmp18386 := PrimHead(V2697)

																																								__e.TailApply(tmp18371, tmp18386)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp18388 := PrimTail(V2696)

																																						tmp18389 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18388, V2783)

																																						__e.TailApply(tmp18369, tmp18389)
																																						return

																																					}, 1)

																																					tmp18390 := PrimHead(V2696)

																																					__e.TailApply(tmp18368, tmp18390)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp18392 := PrimTail(V2694)

																																			tmp18393 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18392, V2783)

																																			__e.TailApply(tmp18366, tmp18393)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp18395 := PrimHead(V2694)

																																	tmp18396 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18395, V2783)

																																	__e.TailApply(tmp18364, tmp18396)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp18398 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																															tmp18399 := Call(__e, tmp18362, tmp18398)

																															__e.TailApply(tmp18203, tmp18399)
																															return

																														} else {
																															__e.Return(Case)
																															return
																														}

																													}, 1)

																													tmp18401 := MakeNative(func(__e *ControlFlow) {
																														V2689 := __e.Get(1)
																														_ = V2689
																														tmp18430 := PrimIsPair(V2689)

																														if True == tmp18430 {
																															tmp18403 := MakeNative(func(__e *ControlFlow) {
																																V2690 := __e.Get(1)
																																_ = V2690
																																tmp18427 := PrimEqual(symtype, V2690)

																																if True == tmp18427 {
																																	tmp18405 := MakeNative(func(__e *ControlFlow) {
																																		V2691 := __e.Get(1)
																																		_ = V2691
																																		tmp18424 := PrimIsPair(V2691)

																																		if True == tmp18424 {
																																			tmp18407 := MakeNative(func(__e *ControlFlow) {
																																				X := __e.Get(1)
																																				_ = X
																																				tmp18408 := MakeNative(func(__e *ControlFlow) {
																																					V2692 := __e.Get(1)
																																					_ = V2692
																																					tmp18420 := PrimIsPair(V2692)

																																					if True == tmp18420 {
																																						tmp18410 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp18411 := MakeNative(func(__e *ControlFlow) {
																																								V2693 := __e.Get(1)
																																								_ = V2693
																																								tmp18416 := PrimEqual(Nil, V2693)

																																								if True == tmp18416 {
																																									tmp18413 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp18413

																																									tmp18414 := MakeNative(func(__e *ControlFlow) {
																																										tmp18415 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, V2784)
																																											return
																																										}, 0)

																																										__e.TailApply(PrimNS2Value(symunify), A, V2781, V2783, tmp18415)
																																										return

																																									}, 0)

																																									__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18414)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}, 1)

																																							tmp18417 := PrimTail(V2692)

																																							tmp18418 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18417, V2783)

																																							__e.TailApply(tmp18411, tmp18418)
																																							return

																																						}, 1)

																																						tmp18419 := PrimHead(V2692)

																																						__e.TailApply(tmp18410, tmp18419)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}, 1)

																																				tmp18421 := PrimTail(V2691)

																																				tmp18422 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18421, V2783)

																																				__e.TailApply(tmp18408, tmp18422)
																																				return

																																			}, 1)

																																			tmp18423 := PrimHead(V2691)

																																			__e.TailApply(tmp18407, tmp18423)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp18425 := PrimTail(V2689)

																																	tmp18426 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18425, V2783)

																																	__e.TailApply(tmp18405, tmp18426)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp18428 := PrimHead(V2689)

																															tmp18429 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18428, V2783)

																															__e.TailApply(tmp18403, tmp18429)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp18431 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																													tmp18432 := Call(__e, tmp18401, tmp18431)

																													__e.TailApply(tmp18201, tmp18432)
																													return

																												} else {
																													__e.Return(Case)
																													return
																												}

																											}, 1)

																											tmp18434 := MakeNative(func(__e *ControlFlow) {
																												V2678 := __e.Get(1)
																												_ = V2678
																												tmp18595 := PrimIsPair(V2678)

																												if True == tmp18595 {
																													tmp18436 := MakeNative(func(__e *ControlFlow) {
																														V2679 := __e.Get(1)
																														_ = V2679
																														tmp18592 := PrimEqual(symopen, V2679)

																														if True == tmp18592 {
																															tmp18438 := MakeNative(func(__e *ControlFlow) {
																																V2680 := __e.Get(1)
																																_ = V2680
																																tmp18589 := PrimIsPair(V2680)

																																if True == tmp18589 {
																																	tmp18440 := MakeNative(func(__e *ControlFlow) {
																																		FileName := __e.Get(1)
																																		_ = FileName
																																		tmp18441 := MakeNative(func(__e *ControlFlow) {
																																			V2681 := __e.Get(1)
																																			_ = V2681
																																			tmp18585 := PrimIsPair(V2681)

																																			if True == tmp18585 {
																																				tmp18443 := MakeNative(func(__e *ControlFlow) {
																																					Direction2611 := __e.Get(1)
																																					_ = Direction2611
																																					tmp18444 := MakeNative(func(__e *ControlFlow) {
																																						V2682 := __e.Get(1)
																																						_ = V2682
																																						tmp18581 := PrimEqual(Nil, V2682)

																																						if True == tmp18581 {
																																							tmp18446 := MakeNative(func(__e *ControlFlow) {
																																								V2683 := __e.Get(1)
																																								_ = V2683
																																								tmp18579 := PrimIsPair(V2683)

																																								if True == tmp18579 {
																																									tmp18466 := MakeNative(func(__e *ControlFlow) {
																																										V2684 := __e.Get(1)
																																										_ = V2684
																																										tmp18576 := PrimEqual(symstream, V2684)

																																										if True == tmp18576 {
																																											tmp18525 := MakeNative(func(__e *ControlFlow) {
																																												V2685 := __e.Get(1)
																																												_ = V2685
																																												tmp18573 := PrimIsPair(V2685)

																																												if True == tmp18573 {
																																													tmp18544 := MakeNative(func(__e *ControlFlow) {
																																														Direction := __e.Get(1)
																																														_ = Direction
																																														tmp18545 := MakeNative(func(__e *ControlFlow) {
																																															V2686 := __e.Get(1)
																																															_ = V2686
																																															tmp18569 := PrimEqual(Nil, V2686)

																																															if True == tmp18569 {
																																																tmp18561 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																_ = tmp18561

																																																tmp18562 := MakeNative(func(__e *ControlFlow) {
																																																	tmp18563 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18564 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																		tmp18565 := PrimCons(symout, Nil)

																																																		tmp18566 := PrimCons(symin, tmp18565)

																																																		tmp18567 := Call(__e, PrimNS2Value(symelement_2), tmp18564, tmp18566)

																																																		tmp18568 := MakeNative(func(__e *ControlFlow) {
																																																			__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																			return
																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symfwhen), tmp18567, V2783, tmp18568)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18563)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18562)
																																																return

																																															} else {
																																																tmp18560 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2686)

																																																if True == tmp18560 {
																																																	tmp18548 := Call(__e, PrimNS2Value(symshen_4bindv), V2686, Nil, V2783)

																																																	_ = tmp18548

																																																	tmp18549 := MakeNative(func(__e *ControlFlow) {
																																																		Result := __e.Get(1)
																																																		_ = Result
																																																		tmp18550 := Call(__e, PrimNS2Value(symshen_4unbindv), V2686, V2783)

																																																		_ = tmp18550

																																																		__e.Return(Result)
																																																		return

																																																	}, 1)

																																																	tmp18551 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																	_ = tmp18551

																																																	tmp18552 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18553 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18554 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																			tmp18555 := PrimCons(symout, Nil)

																																																			tmp18556 := PrimCons(symin, tmp18555)

																																																			tmp18557 := Call(__e, PrimNS2Value(symelement_2), tmp18554, tmp18556)

																																																			tmp18558 := MakeNative(func(__e *ControlFlow) {
																																																				__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																				return
																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symfwhen), tmp18557, V2783, tmp18558)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18553)
																																																		return

																																																	}, 0)

																																																	tmp18559 := Call(__e, PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18552)

																																																	__e.TailApply(tmp18549, tmp18559)
																																																	return

																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp18570 := PrimTail(V2685)

																																														tmp18571 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18570, V2783)

																																														__e.TailApply(tmp18545, tmp18571)
																																														return

																																													}, 1)

																																													tmp18572 := PrimHead(V2685)

																																													__e.TailApply(tmp18544, tmp18572)
																																													return

																																												} else {
																																													tmp18543 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2685)

																																													if True == tmp18543 {
																																														tmp18528 := MakeNative(func(__e *ControlFlow) {
																																															Direction := __e.Get(1)
																																															_ = Direction
																																															tmp18529 := PrimCons(Direction, Nil)

																																															tmp18530 := Call(__e, PrimNS2Value(symshen_4bindv), V2685, tmp18529, V2783)

																																															_ = tmp18530

																																															tmp18531 := MakeNative(func(__e *ControlFlow) {
																																																Result := __e.Get(1)
																																																_ = Result
																																																tmp18532 := Call(__e, PrimNS2Value(symshen_4unbindv), V2685, V2783)

																																																_ = tmp18532

																																																__e.Return(Result)
																																																return

																																															}, 1)

																																															tmp18533 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																															_ = tmp18533

																																															tmp18534 := MakeNative(func(__e *ControlFlow) {
																																																tmp18535 := MakeNative(func(__e *ControlFlow) {
																																																	tmp18536 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																	tmp18537 := PrimCons(symout, Nil)

																																																	tmp18538 := PrimCons(symin, tmp18537)

																																																	tmp18539 := Call(__e, PrimNS2Value(symelement_2), tmp18536, tmp18538)

																																																	tmp18540 := MakeNative(func(__e *ControlFlow) {
																																																		__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																		return
																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symfwhen), tmp18539, V2783, tmp18540)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18535)
																																																return

																																															}, 0)

																																															tmp18541 := Call(__e, PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18534)

																																															__e.TailApply(tmp18531, tmp18541)
																																															return

																																														}, 1)

																																														tmp18542 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																														__e.TailApply(tmp18528, tmp18542)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp18574 := PrimTail(V2683)

																																											tmp18575 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18574, V2783)

																																											__e.TailApply(tmp18525, tmp18575)
																																											return

																																										} else {
																																											tmp18524 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2684)

																																											if True == tmp18524 {
																																												tmp18469 := Call(__e, PrimNS2Value(symshen_4bindv), V2684, symstream, V2783)

																																												_ = tmp18469

																																												tmp18470 := MakeNative(func(__e *ControlFlow) {
																																													Result := __e.Get(1)
																																													_ = Result
																																													tmp18471 := Call(__e, PrimNS2Value(symshen_4unbindv), V2684, V2783)

																																													_ = tmp18471

																																													__e.Return(Result)
																																													return

																																												}, 1)

																																												tmp18472 := MakeNative(func(__e *ControlFlow) {
																																													V2687 := __e.Get(1)
																																													_ = V2687
																																													tmp18520 := PrimIsPair(V2687)

																																													if True == tmp18520 {
																																														tmp18491 := MakeNative(func(__e *ControlFlow) {
																																															Direction := __e.Get(1)
																																															_ = Direction
																																															tmp18492 := MakeNative(func(__e *ControlFlow) {
																																																V2688 := __e.Get(1)
																																																_ = V2688
																																																tmp18516 := PrimEqual(Nil, V2688)

																																																if True == tmp18516 {
																																																	tmp18508 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																	_ = tmp18508

																																																	tmp18509 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18510 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18511 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																			tmp18512 := PrimCons(symout, Nil)

																																																			tmp18513 := PrimCons(symin, tmp18512)

																																																			tmp18514 := Call(__e, PrimNS2Value(symelement_2), tmp18511, tmp18513)

																																																			tmp18515 := MakeNative(func(__e *ControlFlow) {
																																																				__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																				return
																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symfwhen), tmp18514, V2783, tmp18515)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18510)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18509)
																																																	return

																																																} else {
																																																	tmp18507 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2688)

																																																	if True == tmp18507 {
																																																		tmp18495 := Call(__e, PrimNS2Value(symshen_4bindv), V2688, Nil, V2783)

																																																		_ = tmp18495

																																																		tmp18496 := MakeNative(func(__e *ControlFlow) {
																																																			Result := __e.Get(1)
																																																			_ = Result
																																																			tmp18497 := Call(__e, PrimNS2Value(symshen_4unbindv), V2688, V2783)

																																																			_ = tmp18497

																																																			__e.Return(Result)
																																																			return

																																																		}, 1)

																																																		tmp18498 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																		_ = tmp18498

																																																		tmp18499 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18500 := MakeNative(func(__e *ControlFlow) {
																																																				tmp18501 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																				tmp18502 := PrimCons(symout, Nil)

																																																				tmp18503 := PrimCons(symin, tmp18502)

																																																				tmp18504 := Call(__e, PrimNS2Value(symelement_2), tmp18501, tmp18503)

																																																				tmp18505 := MakeNative(func(__e *ControlFlow) {
																																																					__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																					return
																																																				}, 0)

																																																				__e.TailApply(PrimNS2Value(symfwhen), tmp18504, V2783, tmp18505)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18500)
																																																			return

																																																		}, 0)

																																																		tmp18506 := Call(__e, PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18499)

																																																		__e.TailApply(tmp18496, tmp18506)
																																																		return

																																																	} else {
																																																		__e.Return(False)
																																																		return
																																																	}

																																																}

																																															}, 1)

																																															tmp18517 := PrimTail(V2687)

																																															tmp18518 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18517, V2783)

																																															__e.TailApply(tmp18492, tmp18518)
																																															return

																																														}, 1)

																																														tmp18519 := PrimHead(V2687)

																																														__e.TailApply(tmp18491, tmp18519)
																																														return

																																													} else {
																																														tmp18490 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2687)

																																														if True == tmp18490 {
																																															tmp18475 := MakeNative(func(__e *ControlFlow) {
																																																Direction := __e.Get(1)
																																																_ = Direction
																																																tmp18476 := PrimCons(Direction, Nil)

																																																tmp18477 := Call(__e, PrimNS2Value(symshen_4bindv), V2687, tmp18476, V2783)

																																																_ = tmp18477

																																																tmp18478 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp18479 := Call(__e, PrimNS2Value(symshen_4unbindv), V2687, V2783)

																																																	_ = tmp18479

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp18480 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																_ = tmp18480

																																																tmp18481 := MakeNative(func(__e *ControlFlow) {
																																																	tmp18482 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18483 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																																		tmp18484 := PrimCons(symout, Nil)

																																																		tmp18485 := PrimCons(symin, tmp18484)

																																																		tmp18486 := Call(__e, PrimNS2Value(symelement_2), tmp18483, tmp18485)

																																																		tmp18487 := MakeNative(func(__e *ControlFlow) {
																																																			__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																																			return
																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symfwhen), tmp18486, V2783, tmp18487)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18482)
																																																	return

																																																}, 0)

																																																tmp18488 := Call(__e, PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18481)

																																																__e.TailApply(tmp18478, tmp18488)
																																																return

																																															}, 1)

																																															tmp18489 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																															__e.TailApply(tmp18475, tmp18489)
																																															return

																																														} else {
																																															__e.Return(False)
																																															return
																																														}

																																													}

																																												}, 1)

																																												tmp18521 := PrimTail(V2683)

																																												tmp18522 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18521, V2783)

																																												tmp18523 := Call(__e, tmp18472, tmp18522)

																																												__e.TailApply(tmp18470, tmp18523)
																																												return

																																											} else {
																																												__e.Return(False)
																																												return
																																											}

																																										}

																																									}, 1)

																																									tmp18577 := PrimHead(V2683)

																																									tmp18578 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18577, V2783)

																																									__e.TailApply(tmp18466, tmp18578)
																																									return

																																								} else {
																																									tmp18465 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2683)

																																									if True == tmp18465 {
																																										tmp18449 := MakeNative(func(__e *ControlFlow) {
																																											Direction := __e.Get(1)
																																											_ = Direction
																																											tmp18450 := PrimCons(Direction, Nil)

																																											tmp18451 := PrimCons(symstream, tmp18450)

																																											tmp18452 := Call(__e, PrimNS2Value(symshen_4bindv), V2683, tmp18451, V2783)

																																											_ = tmp18452

																																											tmp18453 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp18454 := Call(__e, PrimNS2Value(symshen_4unbindv), V2683, V2783)

																																												_ = tmp18454

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp18455 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp18455

																																											tmp18456 := MakeNative(func(__e *ControlFlow) {
																																												tmp18457 := MakeNative(func(__e *ControlFlow) {
																																													tmp18458 := Call(__e, PrimNS2Value(symshen_4lazyderef), Direction, V2783)

																																													tmp18459 := PrimCons(symout, Nil)

																																													tmp18460 := PrimCons(symin, tmp18459)

																																													tmp18461 := Call(__e, PrimNS2Value(symelement_2), tmp18458, tmp18460)

																																													tmp18462 := MakeNative(func(__e *ControlFlow) {
																																														__e.TailApply(PrimNS2Value(symshen_4th_d), FileName, symstring, V2782, V2783, V2784)
																																														return
																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symfwhen), tmp18461, V2783, tmp18462)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2783, tmp18457)
																																												return

																																											}, 0)

																																											tmp18463 := Call(__e, PrimNS2Value(symunify_b), Direction, Direction2611, V2783, tmp18456)

																																											__e.TailApply(tmp18453, tmp18463)
																																											return

																																										}, 1)

																																										tmp18464 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																										__e.TailApply(tmp18449, tmp18464)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp18580 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																							__e.TailApply(tmp18446, tmp18580)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}, 1)

																																					tmp18582 := PrimTail(V2681)

																																					tmp18583 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18582, V2783)

																																					__e.TailApply(tmp18444, tmp18583)
																																					return

																																				}, 1)

																																				tmp18584 := PrimHead(V2681)

																																				__e.TailApply(tmp18443, tmp18584)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}, 1)

																																		tmp18586 := PrimTail(V2680)

																																		tmp18587 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18586, V2783)

																																		__e.TailApply(tmp18441, tmp18587)
																																		return

																																	}, 1)

																																	tmp18588 := PrimHead(V2680)

																																	__e.TailApply(tmp18440, tmp18588)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp18590 := PrimTail(V2678)

																															tmp18591 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18590, V2783)

																															__e.TailApply(tmp18438, tmp18591)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp18593 := PrimHead(V2678)

																													tmp18594 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18593, V2783)

																													__e.TailApply(tmp18436, tmp18594)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp18596 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																											tmp18597 := Call(__e, tmp18434, tmp18596)

																											__e.TailApply(tmp18199, tmp18597)
																											return

																										} else {
																											__e.Return(Case)
																											return
																										}

																									}, 1)

																									tmp18599 := MakeNative(func(__e *ControlFlow) {
																										V2672 := __e.Get(1)
																										_ = V2672
																										tmp18651 := PrimIsPair(V2672)

																										if True == tmp18651 {
																											tmp18601 := MakeNative(func(__e *ControlFlow) {
																												V2673 := __e.Get(1)
																												_ = V2673
																												tmp18648 := PrimEqual(symlet, V2673)

																												if True == tmp18648 {
																													tmp18603 := MakeNative(func(__e *ControlFlow) {
																														V2674 := __e.Get(1)
																														_ = V2674
																														tmp18645 := PrimIsPair(V2674)

																														if True == tmp18645 {
																															tmp18605 := MakeNative(func(__e *ControlFlow) {
																																X := __e.Get(1)
																																_ = X
																																tmp18606 := MakeNative(func(__e *ControlFlow) {
																																	V2675 := __e.Get(1)
																																	_ = V2675
																																	tmp18641 := PrimIsPair(V2675)

																																	if True == tmp18641 {
																																		tmp18608 := MakeNative(func(__e *ControlFlow) {
																																			Y := __e.Get(1)
																																			_ = Y
																																			tmp18609 := MakeNative(func(__e *ControlFlow) {
																																				V2676 := __e.Get(1)
																																				_ = V2676
																																				tmp18637 := PrimIsPair(V2676)

																																				if True == tmp18637 {
																																					tmp18611 := MakeNative(func(__e *ControlFlow) {
																																						Z := __e.Get(1)
																																						_ = Z
																																						tmp18612 := MakeNative(func(__e *ControlFlow) {
																																							V2677 := __e.Get(1)
																																							_ = V2677
																																							tmp18633 := PrimEqual(Nil, V2677)

																																							if True == tmp18633 {
																																								tmp18614 := MakeNative(func(__e *ControlFlow) {
																																									W := __e.Get(1)
																																									_ = W
																																									tmp18615 := MakeNative(func(__e *ControlFlow) {
																																										X_e_e := __e.Get(1)
																																										_ = X_e_e
																																										tmp18616 := MakeNative(func(__e *ControlFlow) {
																																											B := __e.Get(1)
																																											_ = B
																																											tmp18617 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp18617

																																											tmp18618 := MakeNative(func(__e *ControlFlow) {
																																												tmp18619 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																												tmp18620 := MakeNative(func(__e *ControlFlow) {
																																													tmp18621 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																													tmp18622 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																													tmp18623 := Call(__e, PrimNS2Value(symshen_4lazyderef), Z, V2783)

																																													tmp18624 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18621, tmp18622, tmp18623)

																																													tmp18625 := MakeNative(func(__e *ControlFlow) {
																																														tmp18626 := PrimCons(B, Nil)

																																														tmp18627 := PrimCons(sym_h, tmp18626)

																																														tmp18628 := PrimCons(X_e_e, tmp18627)

																																														tmp18629 := PrimCons(tmp18628, V2782)

																																														__e.TailApply(PrimNS2Value(symshen_4th_d), W, V2781, tmp18629, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symbind), W, tmp18624, V2783, tmp18625)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18619, V2783, tmp18620)
																																												return

																																											}, 0)

																																											__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, tmp18618)
																																											return

																																										}, 1)

																																										tmp18630 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																										__e.TailApply(tmp18616, tmp18630)
																																										return

																																									}, 1)

																																									tmp18631 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																									__e.TailApply(tmp18615, tmp18631)
																																									return

																																								}, 1)

																																								tmp18632 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																								__e.TailApply(tmp18614, tmp18632)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}, 1)

																																						tmp18634 := PrimTail(V2676)

																																						tmp18635 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18634, V2783)

																																						__e.TailApply(tmp18612, tmp18635)
																																						return

																																					}, 1)

																																					tmp18636 := PrimHead(V2676)

																																					__e.TailApply(tmp18611, tmp18636)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}, 1)

																																			tmp18638 := PrimTail(V2675)

																																			tmp18639 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18638, V2783)

																																			__e.TailApply(tmp18609, tmp18639)
																																			return

																																		}, 1)

																																		tmp18640 := PrimHead(V2675)

																																		__e.TailApply(tmp18608, tmp18640)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}, 1)

																																tmp18642 := PrimTail(V2674)

																																tmp18643 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18642, V2783)

																																__e.TailApply(tmp18606, tmp18643)
																																return

																															}, 1)

																															tmp18644 := PrimHead(V2674)

																															__e.TailApply(tmp18605, tmp18644)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp18646 := PrimTail(V2672)

																													tmp18647 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18646, V2783)

																													__e.TailApply(tmp18603, tmp18647)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp18649 := PrimHead(V2672)

																											tmp18650 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18649, V2783)

																											__e.TailApply(tmp18601, tmp18650)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp18652 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																									tmp18653 := Call(__e, tmp18599, tmp18652)

																									__e.TailApply(tmp18197, tmp18653)
																									return

																								} else {
																									__e.Return(Case)
																									return
																								}

																							}, 1)

																							tmp18655 := MakeNative(func(__e *ControlFlow) {
																								V2660 := __e.Get(1)
																								_ = V2660
																								tmp18908 := PrimIsPair(V2660)

																								if True == tmp18908 {
																									tmp18657 := MakeNative(func(__e *ControlFlow) {
																										V2661 := __e.Get(1)
																										_ = V2661
																										tmp18905 := PrimEqual(symlambda, V2661)

																										if True == tmp18905 {
																											tmp18659 := MakeNative(func(__e *ControlFlow) {
																												V2662 := __e.Get(1)
																												_ = V2662
																												tmp18902 := PrimIsPair(V2662)

																												if True == tmp18902 {
																													tmp18661 := MakeNative(func(__e *ControlFlow) {
																														X := __e.Get(1)
																														_ = X
																														tmp18662 := MakeNative(func(__e *ControlFlow) {
																															V2663 := __e.Get(1)
																															_ = V2663
																															tmp18898 := PrimIsPair(V2663)

																															if True == tmp18898 {
																																tmp18664 := MakeNative(func(__e *ControlFlow) {
																																	Y := __e.Get(1)
																																	_ = Y
																																	tmp18665 := MakeNative(func(__e *ControlFlow) {
																																		V2664 := __e.Get(1)
																																		_ = V2664
																																		tmp18894 := PrimEqual(Nil, V2664)

																																		if True == tmp18894 {
																																			tmp18667 := MakeNative(func(__e *ControlFlow) {
																																				V2665 := __e.Get(1)
																																				_ = V2665
																																				tmp18892 := PrimIsPair(V2665)

																																				if True == tmp18892 {
																																					tmp18698 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp18699 := MakeNative(func(__e *ControlFlow) {
																																							V2666 := __e.Get(1)
																																							_ = V2666
																																							tmp18888 := PrimIsPair(V2666)

																																							if True == tmp18888 {
																																								tmp18727 := MakeNative(func(__e *ControlFlow) {
																																									V2667 := __e.Get(1)
																																									_ = V2667
																																									tmp18885 := PrimEqual(sym_1_1_6, V2667)

																																									if True == tmp18885 {
																																										tmp18810 := MakeNative(func(__e *ControlFlow) {
																																											V2668 := __e.Get(1)
																																											_ = V2668
																																											tmp18882 := PrimIsPair(V2668)

																																											if True == tmp18882 {
																																												tmp18837 := MakeNative(func(__e *ControlFlow) {
																																													B := __e.Get(1)
																																													_ = B
																																													tmp18838 := MakeNative(func(__e *ControlFlow) {
																																														V2669 := __e.Get(1)
																																														_ = V2669
																																														tmp18878 := PrimEqual(Nil, V2669)

																																														if True == tmp18878 {
																																															tmp18862 := MakeNative(func(__e *ControlFlow) {
																																																Z := __e.Get(1)
																																																_ = Z
																																																tmp18863 := MakeNative(func(__e *ControlFlow) {
																																																	X_e_e := __e.Get(1)
																																																	_ = X_e_e
																																																	tmp18864 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																	_ = tmp18864

																																																	tmp18865 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																	tmp18866 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18867 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																		tmp18868 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																		tmp18869 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																		tmp18870 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18867, tmp18868, tmp18869)

																																																		tmp18871 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18872 := PrimCons(A, Nil)

																																																			tmp18873 := PrimCons(sym_h, tmp18872)

																																																			tmp18874 := PrimCons(X_e_e, tmp18873)

																																																			tmp18875 := PrimCons(tmp18874, V2782)

																																																			__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18875, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symbind), Z, tmp18870, V2783, tmp18871)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18865, V2783, tmp18866)
																																																	return

																																																}, 1)

																																																tmp18876 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																__e.TailApply(tmp18863, tmp18876)
																																																return

																																															}, 1)

																																															tmp18877 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																															__e.TailApply(tmp18862, tmp18877)
																																															return

																																														} else {
																																															tmp18861 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2669)

																																															if True == tmp18861 {
																																																tmp18841 := Call(__e, PrimNS2Value(symshen_4bindv), V2669, Nil, V2783)

																																																_ = tmp18841

																																																tmp18842 := MakeNative(func(__e *ControlFlow) {
																																																	Result := __e.Get(1)
																																																	_ = Result
																																																	tmp18843 := Call(__e, PrimNS2Value(symshen_4unbindv), V2669, V2783)

																																																	_ = tmp18843

																																																	__e.Return(Result)
																																																	return

																																																}, 1)

																																																tmp18844 := MakeNative(func(__e *ControlFlow) {
																																																	Z := __e.Get(1)
																																																	_ = Z
																																																	tmp18845 := MakeNative(func(__e *ControlFlow) {
																																																		X_e_e := __e.Get(1)
																																																		_ = X_e_e
																																																		tmp18846 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																		_ = tmp18846

																																																		tmp18847 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																		tmp18848 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18849 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																			tmp18850 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																			tmp18851 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																			tmp18852 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18849, tmp18850, tmp18851)

																																																			tmp18853 := MakeNative(func(__e *ControlFlow) {
																																																				tmp18854 := PrimCons(A, Nil)

																																																				tmp18855 := PrimCons(sym_h, tmp18854)

																																																				tmp18856 := PrimCons(X_e_e, tmp18855)

																																																				tmp18857 := PrimCons(tmp18856, V2782)

																																																				__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18857, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symbind), Z, tmp18852, V2783, tmp18853)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18847, V2783, tmp18848)
																																																		return

																																																	}, 1)

																																																	tmp18858 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																	__e.TailApply(tmp18845, tmp18858)
																																																	return

																																																}, 1)

																																																tmp18859 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																tmp18860 := Call(__e, tmp18844, tmp18859)

																																																__e.TailApply(tmp18842, tmp18860)
																																																return

																																															} else {
																																																__e.Return(False)
																																																return
																																															}

																																														}

																																													}, 1)

																																													tmp18879 := PrimTail(V2668)

																																													tmp18880 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18879, V2783)

																																													__e.TailApply(tmp18838, tmp18880)
																																													return

																																												}, 1)

																																												tmp18881 := PrimHead(V2668)

																																												__e.TailApply(tmp18837, tmp18881)
																																												return

																																											} else {
																																												tmp18836 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2668)

																																												if True == tmp18836 {
																																													tmp18813 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp18814 := PrimCons(B, Nil)

																																														tmp18815 := Call(__e, PrimNS2Value(symshen_4bindv), V2668, tmp18814, V2783)

																																														_ = tmp18815

																																														tmp18816 := MakeNative(func(__e *ControlFlow) {
																																															Result := __e.Get(1)
																																															_ = Result
																																															tmp18817 := Call(__e, PrimNS2Value(symshen_4unbindv), V2668, V2783)

																																															_ = tmp18817

																																															__e.Return(Result)
																																															return

																																														}, 1)

																																														tmp18818 := MakeNative(func(__e *ControlFlow) {
																																															Z := __e.Get(1)
																																															_ = Z
																																															tmp18819 := MakeNative(func(__e *ControlFlow) {
																																																X_e_e := __e.Get(1)
																																																_ = X_e_e
																																																tmp18820 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																_ = tmp18820

																																																tmp18821 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																tmp18822 := MakeNative(func(__e *ControlFlow) {
																																																	tmp18823 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																	tmp18824 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																	tmp18825 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																	tmp18826 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18823, tmp18824, tmp18825)

																																																	tmp18827 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18828 := PrimCons(A, Nil)

																																																		tmp18829 := PrimCons(sym_h, tmp18828)

																																																		tmp18830 := PrimCons(X_e_e, tmp18829)

																																																		tmp18831 := PrimCons(tmp18830, V2782)

																																																		__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18831, V2783, V2784)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symbind), Z, tmp18826, V2783, tmp18827)
																																																	return

																																																}, 0)

																																																__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18821, V2783, tmp18822)
																																																return

																																															}, 1)

																																															tmp18832 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																															__e.TailApply(tmp18819, tmp18832)
																																															return

																																														}, 1)

																																														tmp18833 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																														tmp18834 := Call(__e, tmp18818, tmp18833)

																																														__e.TailApply(tmp18816, tmp18834)
																																														return

																																													}, 1)

																																													tmp18835 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																													__e.TailApply(tmp18813, tmp18835)
																																													return

																																												} else {
																																													__e.Return(False)
																																													return
																																												}

																																											}

																																										}, 1)

																																										tmp18883 := PrimTail(V2666)

																																										tmp18884 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18883, V2783)

																																										__e.TailApply(tmp18810, tmp18884)
																																										return

																																									} else {
																																										tmp18809 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2667)

																																										if True == tmp18809 {
																																											tmp18730 := Call(__e, PrimNS2Value(symshen_4bindv), V2667, sym_1_1_6, V2783)

																																											_ = tmp18730

																																											tmp18731 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp18732 := Call(__e, PrimNS2Value(symshen_4unbindv), V2667, V2783)

																																												_ = tmp18732

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp18733 := MakeNative(func(__e *ControlFlow) {
																																												V2670 := __e.Get(1)
																																												_ = V2670
																																												tmp18805 := PrimIsPair(V2670)

																																												if True == tmp18805 {
																																													tmp18760 := MakeNative(func(__e *ControlFlow) {
																																														B := __e.Get(1)
																																														_ = B
																																														tmp18761 := MakeNative(func(__e *ControlFlow) {
																																															V2671 := __e.Get(1)
																																															_ = V2671
																																															tmp18801 := PrimEqual(Nil, V2671)

																																															if True == tmp18801 {
																																																tmp18785 := MakeNative(func(__e *ControlFlow) {
																																																	Z := __e.Get(1)
																																																	_ = Z
																																																	tmp18786 := MakeNative(func(__e *ControlFlow) {
																																																		X_e_e := __e.Get(1)
																																																		_ = X_e_e
																																																		tmp18787 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																		_ = tmp18787

																																																		tmp18788 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																		tmp18789 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18790 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																			tmp18791 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																			tmp18792 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																			tmp18793 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18790, tmp18791, tmp18792)

																																																			tmp18794 := MakeNative(func(__e *ControlFlow) {
																																																				tmp18795 := PrimCons(A, Nil)

																																																				tmp18796 := PrimCons(sym_h, tmp18795)

																																																				tmp18797 := PrimCons(X_e_e, tmp18796)

																																																				tmp18798 := PrimCons(tmp18797, V2782)

																																																				__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18798, V2783, V2784)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symbind), Z, tmp18793, V2783, tmp18794)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18788, V2783, tmp18789)
																																																		return

																																																	}, 1)

																																																	tmp18799 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																	__e.TailApply(tmp18786, tmp18799)
																																																	return

																																																}, 1)

																																																tmp18800 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																__e.TailApply(tmp18785, tmp18800)
																																																return

																																															} else {
																																																tmp18784 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2671)

																																																if True == tmp18784 {
																																																	tmp18764 := Call(__e, PrimNS2Value(symshen_4bindv), V2671, Nil, V2783)

																																																	_ = tmp18764

																																																	tmp18765 := MakeNative(func(__e *ControlFlow) {
																																																		Result := __e.Get(1)
																																																		_ = Result
																																																		tmp18766 := Call(__e, PrimNS2Value(symshen_4unbindv), V2671, V2783)

																																																		_ = tmp18766

																																																		__e.Return(Result)
																																																		return

																																																	}, 1)

																																																	tmp18767 := MakeNative(func(__e *ControlFlow) {
																																																		Z := __e.Get(1)
																																																		_ = Z
																																																		tmp18768 := MakeNative(func(__e *ControlFlow) {
																																																			X_e_e := __e.Get(1)
																																																			_ = X_e_e
																																																			tmp18769 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																			_ = tmp18769

																																																			tmp18770 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																			tmp18771 := MakeNative(func(__e *ControlFlow) {
																																																				tmp18772 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																				tmp18773 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																				tmp18774 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																				tmp18775 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18772, tmp18773, tmp18774)

																																																				tmp18776 := MakeNative(func(__e *ControlFlow) {
																																																					tmp18777 := PrimCons(A, Nil)

																																																					tmp18778 := PrimCons(sym_h, tmp18777)

																																																					tmp18779 := PrimCons(X_e_e, tmp18778)

																																																					tmp18780 := PrimCons(tmp18779, V2782)

																																																					__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18780, V2783, V2784)
																																																					return

																																																				}, 0)

																																																				__e.TailApply(PrimNS2Value(symbind), Z, tmp18775, V2783, tmp18776)
																																																				return

																																																			}, 0)

																																																			__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18770, V2783, tmp18771)
																																																			return

																																																		}, 1)

																																																		tmp18781 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																		__e.TailApply(tmp18768, tmp18781)
																																																		return

																																																	}, 1)

																																																	tmp18782 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																	tmp18783 := Call(__e, tmp18767, tmp18782)

																																																	__e.TailApply(tmp18765, tmp18783)
																																																	return

																																																} else {
																																																	__e.Return(False)
																																																	return
																																																}

																																															}

																																														}, 1)

																																														tmp18802 := PrimTail(V2670)

																																														tmp18803 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18802, V2783)

																																														__e.TailApply(tmp18761, tmp18803)
																																														return

																																													}, 1)

																																													tmp18804 := PrimHead(V2670)

																																													__e.TailApply(tmp18760, tmp18804)
																																													return

																																												} else {
																																													tmp18759 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2670)

																																													if True == tmp18759 {
																																														tmp18736 := MakeNative(func(__e *ControlFlow) {
																																															B := __e.Get(1)
																																															_ = B
																																															tmp18737 := PrimCons(B, Nil)

																																															tmp18738 := Call(__e, PrimNS2Value(symshen_4bindv), V2670, tmp18737, V2783)

																																															_ = tmp18738

																																															tmp18739 := MakeNative(func(__e *ControlFlow) {
																																																Result := __e.Get(1)
																																																_ = Result
																																																tmp18740 := Call(__e, PrimNS2Value(symshen_4unbindv), V2670, V2783)

																																																_ = tmp18740

																																																__e.Return(Result)
																																																return

																																															}, 1)

																																															tmp18741 := MakeNative(func(__e *ControlFlow) {
																																																Z := __e.Get(1)
																																																_ = Z
																																																tmp18742 := MakeNative(func(__e *ControlFlow) {
																																																	X_e_e := __e.Get(1)
																																																	_ = X_e_e
																																																	tmp18743 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																																	_ = tmp18743

																																																	tmp18744 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																																	tmp18745 := MakeNative(func(__e *ControlFlow) {
																																																		tmp18746 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																																		tmp18747 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																																		tmp18748 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																																		tmp18749 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18746, tmp18747, tmp18748)

																																																		tmp18750 := MakeNative(func(__e *ControlFlow) {
																																																			tmp18751 := PrimCons(A, Nil)

																																																			tmp18752 := PrimCons(sym_h, tmp18751)

																																																			tmp18753 := PrimCons(X_e_e, tmp18752)

																																																			tmp18754 := PrimCons(tmp18753, V2782)

																																																			__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18754, V2783, V2784)
																																																			return

																																																		}, 0)

																																																		__e.TailApply(PrimNS2Value(symbind), Z, tmp18749, V2783, tmp18750)
																																																		return

																																																	}, 0)

																																																	__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18744, V2783, tmp18745)
																																																	return

																																																}, 1)

																																																tmp18755 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																																__e.TailApply(tmp18742, tmp18755)
																																																return

																																															}, 1)

																																															tmp18756 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																															tmp18757 := Call(__e, tmp18741, tmp18756)

																																															__e.TailApply(tmp18739, tmp18757)
																																															return

																																														}, 1)

																																														tmp18758 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																														__e.TailApply(tmp18736, tmp18758)
																																														return

																																													} else {
																																														__e.Return(False)
																																														return
																																													}

																																												}

																																											}, 1)

																																											tmp18806 := PrimTail(V2666)

																																											tmp18807 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18806, V2783)

																																											tmp18808 := Call(__e, tmp18733, tmp18807)

																																											__e.TailApply(tmp18731, tmp18808)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp18886 := PrimHead(V2666)

																																								tmp18887 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18886, V2783)

																																								__e.TailApply(tmp18727, tmp18887)
																																								return

																																							} else {
																																								tmp18726 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2666)

																																								if True == tmp18726 {
																																									tmp18702 := MakeNative(func(__e *ControlFlow) {
																																										B := __e.Get(1)
																																										_ = B
																																										tmp18703 := PrimCons(B, Nil)

																																										tmp18704 := PrimCons(sym_1_1_6, tmp18703)

																																										tmp18705 := Call(__e, PrimNS2Value(symshen_4bindv), V2666, tmp18704, V2783)

																																										_ = tmp18705

																																										tmp18706 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp18707 := Call(__e, PrimNS2Value(symshen_4unbindv), V2666, V2783)

																																											_ = tmp18707

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp18708 := MakeNative(func(__e *ControlFlow) {
																																											Z := __e.Get(1)
																																											_ = Z
																																											tmp18709 := MakeNative(func(__e *ControlFlow) {
																																												X_e_e := __e.Get(1)
																																												_ = X_e_e
																																												tmp18710 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp18710

																																												tmp18711 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																												tmp18712 := MakeNative(func(__e *ControlFlow) {
																																													tmp18713 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																													tmp18714 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																													tmp18715 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																													tmp18716 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18713, tmp18714, tmp18715)

																																													tmp18717 := MakeNative(func(__e *ControlFlow) {
																																														tmp18718 := PrimCons(A, Nil)

																																														tmp18719 := PrimCons(sym_h, tmp18718)

																																														tmp18720 := PrimCons(X_e_e, tmp18719)

																																														tmp18721 := PrimCons(tmp18720, V2782)

																																														__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18721, V2783, V2784)
																																														return

																																													}, 0)

																																													__e.TailApply(PrimNS2Value(symbind), Z, tmp18716, V2783, tmp18717)
																																													return

																																												}, 0)

																																												__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18711, V2783, tmp18712)
																																												return

																																											}, 1)

																																											tmp18722 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																											__e.TailApply(tmp18709, tmp18722)
																																											return

																																										}, 1)

																																										tmp18723 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																										tmp18724 := Call(__e, tmp18708, tmp18723)

																																										__e.TailApply(tmp18706, tmp18724)
																																										return

																																									}, 1)

																																									tmp18725 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																									__e.TailApply(tmp18702, tmp18725)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp18889 := PrimTail(V2665)

																																						tmp18890 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18889, V2783)

																																						__e.TailApply(tmp18699, tmp18890)
																																						return

																																					}, 1)

																																					tmp18891 := PrimHead(V2665)

																																					__e.TailApply(tmp18698, tmp18891)
																																					return

																																				} else {
																																					tmp18697 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2665)

																																					if True == tmp18697 {
																																						tmp18670 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp18671 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp18672 := PrimCons(B, Nil)

																																								tmp18673 := PrimCons(sym_1_1_6, tmp18672)

																																								tmp18674 := PrimCons(A, tmp18673)

																																								tmp18675 := Call(__e, PrimNS2Value(symshen_4bindv), V2665, tmp18674, V2783)

																																								_ = tmp18675

																																								tmp18676 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp18677 := Call(__e, PrimNS2Value(symshen_4unbindv), V2665, V2783)

																																									_ = tmp18677

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp18678 := MakeNative(func(__e *ControlFlow) {
																																									Z := __e.Get(1)
																																									_ = Z
																																									tmp18679 := MakeNative(func(__e *ControlFlow) {
																																										X_e_e := __e.Get(1)
																																										_ = X_e_e
																																										tmp18680 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp18680

																																										tmp18681 := Call(__e, PrimNS2Value(symshen_4placeholder))

																																										tmp18682 := MakeNative(func(__e *ControlFlow) {
																																											tmp18683 := Call(__e, PrimNS2Value(symshen_4lazyderef), X_e_e, V2783)

																																											tmp18684 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2783)

																																											tmp18685 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2783)

																																											tmp18686 := Call(__e, PrimNS2Value(symshen_4ebr), tmp18683, tmp18684, tmp18685)

																																											tmp18687 := MakeNative(func(__e *ControlFlow) {
																																												tmp18688 := PrimCons(A, Nil)

																																												tmp18689 := PrimCons(sym_h, tmp18688)

																																												tmp18690 := PrimCons(X_e_e, tmp18689)

																																												tmp18691 := PrimCons(tmp18690, V2782)

																																												__e.TailApply(PrimNS2Value(symshen_4th_d), Z, B, tmp18691, V2783, V2784)
																																												return

																																											}, 0)

																																											__e.TailApply(PrimNS2Value(symbind), Z, tmp18686, V2783, tmp18687)
																																											return

																																										}, 0)

																																										__e.TailApply(PrimNS2Value(symbind), X_e_e, tmp18681, V2783, tmp18682)
																																										return

																																									}, 1)

																																									tmp18692 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																									__e.TailApply(tmp18679, tmp18692)
																																									return

																																								}, 1)

																																								tmp18693 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																								tmp18694 := Call(__e, tmp18678, tmp18693)

																																								__e.TailApply(tmp18676, tmp18694)
																																								return

																																							}, 1)

																																							tmp18695 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																							__e.TailApply(tmp18671, tmp18695)
																																							return

																																						}, 1)

																																						tmp18696 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																						__e.TailApply(tmp18670, tmp18696)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp18893 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																			__e.TailApply(tmp18667, tmp18893)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}, 1)

																																	tmp18895 := PrimTail(V2663)

																																	tmp18896 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18895, V2783)

																																	__e.TailApply(tmp18665, tmp18896)
																																	return

																																}, 1)

																																tmp18897 := PrimHead(V2663)

																																__e.TailApply(tmp18664, tmp18897)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}, 1)

																														tmp18899 := PrimTail(V2662)

																														tmp18900 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18899, V2783)

																														__e.TailApply(tmp18662, tmp18900)
																														return

																													}, 1)

																													tmp18901 := PrimHead(V2662)

																													__e.TailApply(tmp18661, tmp18901)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp18903 := PrimTail(V2660)

																											tmp18904 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18903, V2783)

																											__e.TailApply(tmp18659, tmp18904)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp18906 := PrimHead(V2660)

																									tmp18907 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18906, V2783)

																									__e.TailApply(tmp18657, tmp18907)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp18909 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																							tmp18910 := Call(__e, tmp18655, tmp18909)

																							__e.TailApply(tmp18195, tmp18910)
																							return

																						} else {
																							__e.Return(Case)
																							return
																						}

																					}, 1)

																					tmp18912 := MakeNative(func(__e *ControlFlow) {
																						V2654 := __e.Get(1)
																						_ = V2654
																						tmp18952 := PrimIsPair(V2654)

																						if True == tmp18952 {
																							tmp18914 := MakeNative(func(__e *ControlFlow) {
																								V2655 := __e.Get(1)
																								_ = V2655
																								tmp18949 := PrimEqual(sym_8s, V2655)

																								if True == tmp18949 {
																									tmp18916 := MakeNative(func(__e *ControlFlow) {
																										V2656 := __e.Get(1)
																										_ = V2656
																										tmp18946 := PrimIsPair(V2656)

																										if True == tmp18946 {
																											tmp18918 := MakeNative(func(__e *ControlFlow) {
																												X := __e.Get(1)
																												_ = X
																												tmp18919 := MakeNative(func(__e *ControlFlow) {
																													V2657 := __e.Get(1)
																													_ = V2657
																													tmp18942 := PrimIsPair(V2657)

																													if True == tmp18942 {
																														tmp18921 := MakeNative(func(__e *ControlFlow) {
																															Y := __e.Get(1)
																															_ = Y
																															tmp18922 := MakeNative(func(__e *ControlFlow) {
																																V2658 := __e.Get(1)
																																_ = V2658
																																tmp18938 := PrimEqual(Nil, V2658)

																																if True == tmp18938 {
																																	tmp18924 := MakeNative(func(__e *ControlFlow) {
																																		V2659 := __e.Get(1)
																																		_ = V2659
																																		tmp18936 := PrimEqual(symstring, V2659)

																																		if True == tmp18936 {
																																			tmp18934 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp18934

																																			tmp18935 := MakeNative(func(__e *ControlFlow) {
																																				__e.TailApply(PrimNS2Value(symshen_4th_d), Y, symstring, V2782, V2783, V2784)
																																				return
																																			}, 0)

																																			__e.TailApply(PrimNS2Value(symshen_4th_d), X, symstring, V2782, V2783, tmp18935)
																																			return

																																		} else {
																																			tmp18933 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2659)

																																			if True == tmp18933 {
																																				tmp18927 := Call(__e, PrimNS2Value(symshen_4bindv), V2659, symstring, V2783)

																																				_ = tmp18927

																																				tmp18928 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp18929 := Call(__e, PrimNS2Value(symshen_4unbindv), V2659, V2783)

																																					_ = tmp18929

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp18930 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp18930

																																				tmp18931 := MakeNative(func(__e *ControlFlow) {
																																					__e.TailApply(PrimNS2Value(symshen_4th_d), Y, symstring, V2782, V2783, V2784)
																																					return
																																				}, 0)

																																				tmp18932 := Call(__e, PrimNS2Value(symshen_4th_d), X, symstring, V2782, V2783, tmp18931)

																																				__e.TailApply(tmp18928, tmp18932)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp18937 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																																	__e.TailApply(tmp18924, tmp18937)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}, 1)

																															tmp18939 := PrimTail(V2657)

																															tmp18940 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18939, V2783)

																															__e.TailApply(tmp18922, tmp18940)
																															return

																														}, 1)

																														tmp18941 := PrimHead(V2657)

																														__e.TailApply(tmp18921, tmp18941)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp18943 := PrimTail(V2656)

																												tmp18944 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18943, V2783)

																												__e.TailApply(tmp18919, tmp18944)
																												return

																											}, 1)

																											tmp18945 := PrimHead(V2656)

																											__e.TailApply(tmp18918, tmp18945)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp18947 := PrimTail(V2654)

																									tmp18948 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18947, V2783)

																									__e.TailApply(tmp18916, tmp18948)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp18950 := PrimHead(V2654)

																							tmp18951 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp18950, V2783)

																							__e.TailApply(tmp18914, tmp18951)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp18953 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																					tmp18954 := Call(__e, tmp18912, tmp18953)

																					__e.TailApply(tmp18193, tmp18954)
																					return

																				} else {
																					__e.Return(Case)
																					return
																				}

																			}, 1)

																			tmp18956 := MakeNative(func(__e *ControlFlow) {
																				V2643 := __e.Get(1)
																				_ = V2643
																				tmp19089 := PrimIsPair(V2643)

																				if True == tmp19089 {
																					tmp18958 := MakeNative(func(__e *ControlFlow) {
																						V2644 := __e.Get(1)
																						_ = V2644
																						tmp19086 := PrimEqual(sym_8v, V2644)

																						if True == tmp19086 {
																							tmp18960 := MakeNative(func(__e *ControlFlow) {
																								V2645 := __e.Get(1)
																								_ = V2645
																								tmp19083 := PrimIsPair(V2645)

																								if True == tmp19083 {
																									tmp18962 := MakeNative(func(__e *ControlFlow) {
																										X := __e.Get(1)
																										_ = X
																										tmp18963 := MakeNative(func(__e *ControlFlow) {
																											V2646 := __e.Get(1)
																											_ = V2646
																											tmp19079 := PrimIsPair(V2646)

																											if True == tmp19079 {
																												tmp18965 := MakeNative(func(__e *ControlFlow) {
																													Y := __e.Get(1)
																													_ = Y
																													tmp18966 := MakeNative(func(__e *ControlFlow) {
																														V2647 := __e.Get(1)
																														_ = V2647
																														tmp19075 := PrimEqual(Nil, V2647)

																														if True == tmp19075 {
																															tmp18968 := MakeNative(func(__e *ControlFlow) {
																																V2648 := __e.Get(1)
																																_ = V2648
																																tmp19073 := PrimIsPair(V2648)

																																if True == tmp19073 {
																																	tmp18984 := MakeNative(func(__e *ControlFlow) {
																																		V2649 := __e.Get(1)
																																		_ = V2649
																																		tmp19070 := PrimEqual(symvector, V2649)

																																		if True == tmp19070 {
																																			tmp19031 := MakeNative(func(__e *ControlFlow) {
																																				V2650 := __e.Get(1)
																																				_ = V2650
																																				tmp19067 := PrimIsPair(V2650)

																																				if True == tmp19067 {
																																					tmp19046 := MakeNative(func(__e *ControlFlow) {
																																						A := __e.Get(1)
																																						_ = A
																																						tmp19047 := MakeNative(func(__e *ControlFlow) {
																																							V2651 := __e.Get(1)
																																							_ = V2651
																																							tmp19063 := PrimEqual(Nil, V2651)

																																							if True == tmp19063 {
																																								tmp19059 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp19059

																																								tmp19060 := MakeNative(func(__e *ControlFlow) {
																																									tmp19061 := PrimCons(A, Nil)

																																									tmp19062 := PrimCons(symvector, tmp19061)

																																									__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19062, V2782, V2783, V2784)
																																									return

																																								}, 0)

																																								__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19060)
																																								return

																																							} else {
																																								tmp19058 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2651)

																																								if True == tmp19058 {
																																									tmp19050 := Call(__e, PrimNS2Value(symshen_4bindv), V2651, Nil, V2783)

																																									_ = tmp19050

																																									tmp19051 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp19052 := Call(__e, PrimNS2Value(symshen_4unbindv), V2651, V2783)

																																										_ = tmp19052

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp19053 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19053

																																									tmp19054 := MakeNative(func(__e *ControlFlow) {
																																										tmp19055 := PrimCons(A, Nil)

																																										tmp19056 := PrimCons(symvector, tmp19055)

																																										__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19056, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									tmp19057 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19054)

																																									__e.TailApply(tmp19051, tmp19057)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp19064 := PrimTail(V2650)

																																						tmp19065 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19064, V2783)

																																						__e.TailApply(tmp19047, tmp19065)
																																						return

																																					}, 1)

																																					tmp19066 := PrimHead(V2650)

																																					__e.TailApply(tmp19046, tmp19066)
																																					return

																																				} else {
																																					tmp19045 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2650)

																																					if True == tmp19045 {
																																						tmp19034 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp19035 := PrimCons(A, Nil)

																																							tmp19036 := Call(__e, PrimNS2Value(symshen_4bindv), V2650, tmp19035, V2783)

																																							_ = tmp19036

																																							tmp19037 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp19038 := Call(__e, PrimNS2Value(symshen_4unbindv), V2650, V2783)

																																								_ = tmp19038

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp19039 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																							_ = tmp19039

																																							tmp19040 := MakeNative(func(__e *ControlFlow) {
																																								tmp19041 := PrimCons(A, Nil)

																																								tmp19042 := PrimCons(symvector, tmp19041)

																																								__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19042, V2782, V2783, V2784)
																																								return

																																							}, 0)

																																							tmp19043 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19040)

																																							__e.TailApply(tmp19037, tmp19043)
																																							return

																																						}, 1)

																																						tmp19044 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																						__e.TailApply(tmp19034, tmp19044)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp19068 := PrimTail(V2648)

																																			tmp19069 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19068, V2783)

																																			__e.TailApply(tmp19031, tmp19069)
																																			return

																																		} else {
																																			tmp19030 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2649)

																																			if True == tmp19030 {
																																				tmp18987 := Call(__e, PrimNS2Value(symshen_4bindv), V2649, symvector, V2783)

																																				_ = tmp18987

																																				tmp18988 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp18989 := Call(__e, PrimNS2Value(symshen_4unbindv), V2649, V2783)

																																					_ = tmp18989

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp18990 := MakeNative(func(__e *ControlFlow) {
																																					V2652 := __e.Get(1)
																																					_ = V2652
																																					tmp19026 := PrimIsPair(V2652)

																																					if True == tmp19026 {
																																						tmp19005 := MakeNative(func(__e *ControlFlow) {
																																							A := __e.Get(1)
																																							_ = A
																																							tmp19006 := MakeNative(func(__e *ControlFlow) {
																																								V2653 := __e.Get(1)
																																								_ = V2653
																																								tmp19022 := PrimEqual(Nil, V2653)

																																								if True == tmp19022 {
																																									tmp19018 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19018

																																									tmp19019 := MakeNative(func(__e *ControlFlow) {
																																										tmp19020 := PrimCons(A, Nil)

																																										tmp19021 := PrimCons(symvector, tmp19020)

																																										__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19021, V2782, V2783, V2784)
																																										return

																																									}, 0)

																																									__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19019)
																																									return

																																								} else {
																																									tmp19017 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2653)

																																									if True == tmp19017 {
																																										tmp19009 := Call(__e, PrimNS2Value(symshen_4bindv), V2653, Nil, V2783)

																																										_ = tmp19009

																																										tmp19010 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp19011 := Call(__e, PrimNS2Value(symshen_4unbindv), V2653, V2783)

																																											_ = tmp19011

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp19012 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19012

																																										tmp19013 := MakeNative(func(__e *ControlFlow) {
																																											tmp19014 := PrimCons(A, Nil)

																																											tmp19015 := PrimCons(symvector, tmp19014)

																																											__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19015, V2782, V2783, V2784)
																																											return

																																										}, 0)

																																										tmp19016 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19013)

																																										__e.TailApply(tmp19010, tmp19016)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp19023 := PrimTail(V2652)

																																							tmp19024 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19023, V2783)

																																							__e.TailApply(tmp19006, tmp19024)
																																							return

																																						}, 1)

																																						tmp19025 := PrimHead(V2652)

																																						__e.TailApply(tmp19005, tmp19025)
																																						return

																																					} else {
																																						tmp19004 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2652)

																																						if True == tmp19004 {
																																							tmp18993 := MakeNative(func(__e *ControlFlow) {
																																								A := __e.Get(1)
																																								_ = A
																																								tmp18994 := PrimCons(A, Nil)

																																								tmp18995 := Call(__e, PrimNS2Value(symshen_4bindv), V2652, tmp18994, V2783)

																																								_ = tmp18995

																																								tmp18996 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp18997 := Call(__e, PrimNS2Value(symshen_4unbindv), V2652, V2783)

																																									_ = tmp18997

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp18998 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp18998

																																								tmp18999 := MakeNative(func(__e *ControlFlow) {
																																									tmp19000 := PrimCons(A, Nil)

																																									tmp19001 := PrimCons(symvector, tmp19000)

																																									__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19001, V2782, V2783, V2784)
																																									return

																																								}, 0)

																																								tmp19002 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp18999)

																																								__e.TailApply(tmp18996, tmp19002)
																																								return

																																							}, 1)

																																							tmp19003 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																							__e.TailApply(tmp18993, tmp19003)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp19027 := PrimTail(V2648)

																																				tmp19028 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19027, V2783)

																																				tmp19029 := Call(__e, tmp18990, tmp19028)

																																				__e.TailApply(tmp18988, tmp19029)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp19071 := PrimHead(V2648)

																																	tmp19072 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19071, V2783)

																																	__e.TailApply(tmp18984, tmp19072)
																																	return

																																} else {
																																	tmp18983 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2648)

																																	if True == tmp18983 {
																																		tmp18971 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp18972 := PrimCons(A, Nil)

																																			tmp18973 := PrimCons(symvector, tmp18972)

																																			tmp18974 := Call(__e, PrimNS2Value(symshen_4bindv), V2648, tmp18973, V2783)

																																			_ = tmp18974

																																			tmp18975 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp18976 := Call(__e, PrimNS2Value(symshen_4unbindv), V2648, V2783)

																																				_ = tmp18976

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp18977 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp18977

																																			tmp18978 := MakeNative(func(__e *ControlFlow) {
																																				tmp18979 := PrimCons(A, Nil)

																																				tmp18980 := PrimCons(symvector, tmp18979)

																																				__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp18980, V2782, V2783, V2784)
																																				return

																																			}, 0)

																																			tmp18981 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp18978)

																																			__e.TailApply(tmp18975, tmp18981)
																																			return

																																		}, 1)

																																		tmp18982 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																		__e.TailApply(tmp18971, tmp18982)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp19074 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																															__e.TailApply(tmp18968, tmp19074)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}, 1)

																													tmp19076 := PrimTail(V2646)

																													tmp19077 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19076, V2783)

																													__e.TailApply(tmp18966, tmp19077)
																													return

																												}, 1)

																												tmp19078 := PrimHead(V2646)

																												__e.TailApply(tmp18965, tmp19078)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp19080 := PrimTail(V2645)

																										tmp19081 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19080, V2783)

																										__e.TailApply(tmp18963, tmp19081)
																										return

																									}, 1)

																									tmp19082 := PrimHead(V2645)

																									__e.TailApply(tmp18962, tmp19082)
																									return

																								} else {
																									__e.Return(False)
																									return
																								}

																							}, 1)

																							tmp19084 := PrimTail(V2643)

																							tmp19085 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19084, V2783)

																							__e.TailApply(tmp18960, tmp19085)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp19087 := PrimHead(V2643)

																					tmp19088 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19087, V2783)

																					__e.TailApply(tmp18958, tmp19088)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp19090 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																			tmp19091 := Call(__e, tmp18956, tmp19090)

																			__e.TailApply(tmp18191, tmp19091)
																			return

																		} else {
																			__e.Return(Case)
																			return
																		}

																	}, 1)

																	tmp19093 := MakeNative(func(__e *ControlFlow) {
																		V2631 := __e.Get(1)
																		_ = V2631
																		tmp19234 := PrimIsPair(V2631)

																		if True == tmp19234 {
																			tmp19095 := MakeNative(func(__e *ControlFlow) {
																				V2632 := __e.Get(1)
																				_ = V2632
																				tmp19231 := PrimEqual(sym_8p, V2632)

																				if True == tmp19231 {
																					tmp19097 := MakeNative(func(__e *ControlFlow) {
																						V2633 := __e.Get(1)
																						_ = V2633
																						tmp19228 := PrimIsPair(V2633)

																						if True == tmp19228 {
																							tmp19099 := MakeNative(func(__e *ControlFlow) {
																								X := __e.Get(1)
																								_ = X
																								tmp19100 := MakeNative(func(__e *ControlFlow) {
																									V2634 := __e.Get(1)
																									_ = V2634
																									tmp19224 := PrimIsPair(V2634)

																									if True == tmp19224 {
																										tmp19102 := MakeNative(func(__e *ControlFlow) {
																											Y := __e.Get(1)
																											_ = Y
																											tmp19103 := MakeNative(func(__e *ControlFlow) {
																												V2635 := __e.Get(1)
																												_ = V2635
																												tmp19220 := PrimEqual(Nil, V2635)

																												if True == tmp19220 {
																													tmp19105 := MakeNative(func(__e *ControlFlow) {
																														V2636 := __e.Get(1)
																														_ = V2636
																														tmp19218 := PrimIsPair(V2636)

																														if True == tmp19218 {
																															tmp19122 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp19123 := MakeNative(func(__e *ControlFlow) {
																																	V2637 := __e.Get(1)
																																	_ = V2637
																																	tmp19214 := PrimIsPair(V2637)

																																	if True == tmp19214 {
																																		tmp19137 := MakeNative(func(__e *ControlFlow) {
																																			V2638 := __e.Get(1)
																																			_ = V2638
																																			tmp19211 := PrimEqual(sym_d, V2638)

																																			if True == tmp19211 {
																																				tmp19178 := MakeNative(func(__e *ControlFlow) {
																																					V2639 := __e.Get(1)
																																					_ = V2639
																																					tmp19208 := PrimIsPair(V2639)

																																					if True == tmp19208 {
																																						tmp19191 := MakeNative(func(__e *ControlFlow) {
																																							B := __e.Get(1)
																																							_ = B
																																							tmp19192 := MakeNative(func(__e *ControlFlow) {
																																								V2640 := __e.Get(1)
																																								_ = V2640
																																								tmp19204 := PrimEqual(Nil, V2640)

																																								if True == tmp19204 {
																																									tmp19202 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19202

																																									tmp19203 := MakeNative(func(__e *ControlFlow) {
																																										__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																										return
																																									}, 0)

																																									__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19203)
																																									return

																																								} else {
																																									tmp19201 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2640)

																																									if True == tmp19201 {
																																										tmp19195 := Call(__e, PrimNS2Value(symshen_4bindv), V2640, Nil, V2783)

																																										_ = tmp19195

																																										tmp19196 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp19197 := Call(__e, PrimNS2Value(symshen_4unbindv), V2640, V2783)

																																											_ = tmp19197

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp19198 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19198

																																										tmp19199 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																											return
																																										}, 0)

																																										tmp19200 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19199)

																																										__e.TailApply(tmp19196, tmp19200)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp19205 := PrimTail(V2639)

																																							tmp19206 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19205, V2783)

																																							__e.TailApply(tmp19192, tmp19206)
																																							return

																																						}, 1)

																																						tmp19207 := PrimHead(V2639)

																																						__e.TailApply(tmp19191, tmp19207)
																																						return

																																					} else {
																																						tmp19190 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2639)

																																						if True == tmp19190 {
																																							tmp19181 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp19182 := PrimCons(B, Nil)

																																								tmp19183 := Call(__e, PrimNS2Value(symshen_4bindv), V2639, tmp19182, V2783)

																																								_ = tmp19183

																																								tmp19184 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp19185 := Call(__e, PrimNS2Value(symshen_4unbindv), V2639, V2783)

																																									_ = tmp19185

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp19186 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp19186

																																								tmp19187 := MakeNative(func(__e *ControlFlow) {
																																									__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																									return
																																								}, 0)

																																								tmp19188 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19187)

																																								__e.TailApply(tmp19184, tmp19188)
																																								return

																																							}, 1)

																																							tmp19189 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																							__e.TailApply(tmp19181, tmp19189)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp19209 := PrimTail(V2637)

																																				tmp19210 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19209, V2783)

																																				__e.TailApply(tmp19178, tmp19210)
																																				return

																																			} else {
																																				tmp19177 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2638)

																																				if True == tmp19177 {
																																					tmp19140 := Call(__e, PrimNS2Value(symshen_4bindv), V2638, sym_d, V2783)

																																					_ = tmp19140

																																					tmp19141 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp19142 := Call(__e, PrimNS2Value(symshen_4unbindv), V2638, V2783)

																																						_ = tmp19142

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp19143 := MakeNative(func(__e *ControlFlow) {
																																						V2641 := __e.Get(1)
																																						_ = V2641
																																						tmp19173 := PrimIsPair(V2641)

																																						if True == tmp19173 {
																																							tmp19156 := MakeNative(func(__e *ControlFlow) {
																																								B := __e.Get(1)
																																								_ = B
																																								tmp19157 := MakeNative(func(__e *ControlFlow) {
																																									V2642 := __e.Get(1)
																																									_ = V2642
																																									tmp19169 := PrimEqual(Nil, V2642)

																																									if True == tmp19169 {
																																										tmp19167 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19167

																																										tmp19168 := MakeNative(func(__e *ControlFlow) {
																																											__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																											return
																																										}, 0)

																																										__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19168)
																																										return

																																									} else {
																																										tmp19166 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2642)

																																										if True == tmp19166 {
																																											tmp19160 := Call(__e, PrimNS2Value(symshen_4bindv), V2642, Nil, V2783)

																																											_ = tmp19160

																																											tmp19161 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp19162 := Call(__e, PrimNS2Value(symshen_4unbindv), V2642, V2783)

																																												_ = tmp19162

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp19163 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp19163

																																											tmp19164 := MakeNative(func(__e *ControlFlow) {
																																												__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																												return
																																											}, 0)

																																											tmp19165 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19164)

																																											__e.TailApply(tmp19161, tmp19165)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp19170 := PrimTail(V2641)

																																								tmp19171 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19170, V2783)

																																								__e.TailApply(tmp19157, tmp19171)
																																								return

																																							}, 1)

																																							tmp19172 := PrimHead(V2641)

																																							__e.TailApply(tmp19156, tmp19172)
																																							return

																																						} else {
																																							tmp19155 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2641)

																																							if True == tmp19155 {
																																								tmp19146 := MakeNative(func(__e *ControlFlow) {
																																									B := __e.Get(1)
																																									_ = B
																																									tmp19147 := PrimCons(B, Nil)

																																									tmp19148 := Call(__e, PrimNS2Value(symshen_4bindv), V2641, tmp19147, V2783)

																																									_ = tmp19148

																																									tmp19149 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp19150 := Call(__e, PrimNS2Value(symshen_4unbindv), V2641, V2783)

																																										_ = tmp19150

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp19151 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19151

																																									tmp19152 := MakeNative(func(__e *ControlFlow) {
																																										__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																										return
																																									}, 0)

																																									tmp19153 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19152)

																																									__e.TailApply(tmp19149, tmp19153)
																																									return

																																								}, 1)

																																								tmp19154 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																								__e.TailApply(tmp19146, tmp19154)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp19174 := PrimTail(V2637)

																																					tmp19175 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19174, V2783)

																																					tmp19176 := Call(__e, tmp19143, tmp19175)

																																					__e.TailApply(tmp19141, tmp19176)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp19212 := PrimHead(V2637)

																																		tmp19213 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19212, V2783)

																																		__e.TailApply(tmp19137, tmp19213)
																																		return

																																	} else {
																																		tmp19136 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2637)

																																		if True == tmp19136 {
																																			tmp19126 := MakeNative(func(__e *ControlFlow) {
																																				B := __e.Get(1)
																																				_ = B
																																				tmp19127 := PrimCons(B, Nil)

																																				tmp19128 := PrimCons(sym_d, tmp19127)

																																				tmp19129 := Call(__e, PrimNS2Value(symshen_4bindv), V2637, tmp19128, V2783)

																																				_ = tmp19129

																																				tmp19130 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp19131 := Call(__e, PrimNS2Value(symshen_4unbindv), V2637, V2783)

																																					_ = tmp19131

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp19132 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19132

																																				tmp19133 := MakeNative(func(__e *ControlFlow) {
																																					__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																					return
																																				}, 0)

																																				tmp19134 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19133)

																																				__e.TailApply(tmp19130, tmp19134)
																																				return

																																			}, 1)

																																			tmp19135 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																			__e.TailApply(tmp19126, tmp19135)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp19215 := PrimTail(V2636)

																																tmp19216 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19215, V2783)

																																__e.TailApply(tmp19123, tmp19216)
																																return

																															}, 1)

																															tmp19217 := PrimHead(V2636)

																															__e.TailApply(tmp19122, tmp19217)
																															return

																														} else {
																															tmp19121 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2636)

																															if True == tmp19121 {
																																tmp19108 := MakeNative(func(__e *ControlFlow) {
																																	A := __e.Get(1)
																																	_ = A
																																	tmp19109 := MakeNative(func(__e *ControlFlow) {
																																		B := __e.Get(1)
																																		_ = B
																																		tmp19110 := PrimCons(B, Nil)

																																		tmp19111 := PrimCons(sym_d, tmp19110)

																																		tmp19112 := PrimCons(A, tmp19111)

																																		tmp19113 := Call(__e, PrimNS2Value(symshen_4bindv), V2636, tmp19112, V2783)

																																		_ = tmp19113

																																		tmp19114 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp19115 := Call(__e, PrimNS2Value(symshen_4unbindv), V2636, V2783)

																																			_ = tmp19115

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp19116 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																		_ = tmp19116

																																		tmp19117 := MakeNative(func(__e *ControlFlow) {
																																			__e.TailApply(PrimNS2Value(symshen_4th_d), Y, B, V2782, V2783, V2784)
																																			return
																																		}, 0)

																																		tmp19118 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19117)

																																		__e.TailApply(tmp19114, tmp19118)
																																		return

																																	}, 1)

																																	tmp19119 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																	__e.TailApply(tmp19109, tmp19119)
																																	return

																																}, 1)

																																tmp19120 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																__e.TailApply(tmp19108, tmp19120)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp19219 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																													__e.TailApply(tmp19105, tmp19219)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp19221 := PrimTail(V2634)

																											tmp19222 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19221, V2783)

																											__e.TailApply(tmp19103, tmp19222)
																											return

																										}, 1)

																										tmp19223 := PrimHead(V2634)

																										__e.TailApply(tmp19102, tmp19223)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp19225 := PrimTail(V2633)

																								tmp19226 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19225, V2783)

																								__e.TailApply(tmp19100, tmp19226)
																								return

																							}, 1)

																							tmp19227 := PrimHead(V2633)

																							__e.TailApply(tmp19099, tmp19227)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}, 1)

																					tmp19229 := PrimTail(V2631)

																					tmp19230 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19229, V2783)

																					__e.TailApply(tmp19097, tmp19230)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp19232 := PrimHead(V2631)

																			tmp19233 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19232, V2783)

																			__e.TailApply(tmp19095, tmp19233)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp19235 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

																	tmp19236 := Call(__e, tmp19093, tmp19235)

																	__e.TailApply(tmp18189, tmp19236)
																	return

																} else {
																	__e.Return(Case)
																	return
																}

															}, 1)

															tmp19238 := MakeNative(func(__e *ControlFlow) {
																V2620 := __e.Get(1)
																_ = V2620
																tmp19371 := PrimIsPair(V2620)

																if True == tmp19371 {
																	tmp19240 := MakeNative(func(__e *ControlFlow) {
																		V2621 := __e.Get(1)
																		_ = V2621
																		tmp19368 := PrimEqual(symcons, V2621)

																		if True == tmp19368 {
																			tmp19242 := MakeNative(func(__e *ControlFlow) {
																				V2622 := __e.Get(1)
																				_ = V2622
																				tmp19365 := PrimIsPair(V2622)

																				if True == tmp19365 {
																					tmp19244 := MakeNative(func(__e *ControlFlow) {
																						X := __e.Get(1)
																						_ = X
																						tmp19245 := MakeNative(func(__e *ControlFlow) {
																							V2623 := __e.Get(1)
																							_ = V2623
																							tmp19361 := PrimIsPair(V2623)

																							if True == tmp19361 {
																								tmp19247 := MakeNative(func(__e *ControlFlow) {
																									Y := __e.Get(1)
																									_ = Y
																									tmp19248 := MakeNative(func(__e *ControlFlow) {
																										V2624 := __e.Get(1)
																										_ = V2624
																										tmp19357 := PrimEqual(Nil, V2624)

																										if True == tmp19357 {
																											tmp19250 := MakeNative(func(__e *ControlFlow) {
																												V2625 := __e.Get(1)
																												_ = V2625
																												tmp19355 := PrimIsPair(V2625)

																												if True == tmp19355 {
																													tmp19266 := MakeNative(func(__e *ControlFlow) {
																														V2626 := __e.Get(1)
																														_ = V2626
																														tmp19352 := PrimEqual(symlist, V2626)

																														if True == tmp19352 {
																															tmp19313 := MakeNative(func(__e *ControlFlow) {
																																V2627 := __e.Get(1)
																																_ = V2627
																																tmp19349 := PrimIsPair(V2627)

																																if True == tmp19349 {
																																	tmp19328 := MakeNative(func(__e *ControlFlow) {
																																		A := __e.Get(1)
																																		_ = A
																																		tmp19329 := MakeNative(func(__e *ControlFlow) {
																																			V2628 := __e.Get(1)
																																			_ = V2628
																																			tmp19345 := PrimEqual(Nil, V2628)

																																			if True == tmp19345 {
																																				tmp19341 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19341

																																				tmp19342 := MakeNative(func(__e *ControlFlow) {
																																					tmp19343 := PrimCons(A, Nil)

																																					tmp19344 := PrimCons(symlist, tmp19343)

																																					__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19344, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19342)
																																				return

																																			} else {
																																				tmp19340 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2628)

																																				if True == tmp19340 {
																																					tmp19332 := Call(__e, PrimNS2Value(symshen_4bindv), V2628, Nil, V2783)

																																					_ = tmp19332

																																					tmp19333 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp19334 := Call(__e, PrimNS2Value(symshen_4unbindv), V2628, V2783)

																																						_ = tmp19334

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp19335 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp19335

																																					tmp19336 := MakeNative(func(__e *ControlFlow) {
																																						tmp19337 := PrimCons(A, Nil)

																																						tmp19338 := PrimCons(symlist, tmp19337)

																																						__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19338, V2782, V2783, V2784)
																																						return

																																					}, 0)

																																					tmp19339 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19336)

																																					__e.TailApply(tmp19333, tmp19339)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp19346 := PrimTail(V2627)

																																		tmp19347 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19346, V2783)

																																		__e.TailApply(tmp19329, tmp19347)
																																		return

																																	}, 1)

																																	tmp19348 := PrimHead(V2627)

																																	__e.TailApply(tmp19328, tmp19348)
																																	return

																																} else {
																																	tmp19327 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2627)

																																	if True == tmp19327 {
																																		tmp19316 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp19317 := PrimCons(A, Nil)

																																			tmp19318 := Call(__e, PrimNS2Value(symshen_4bindv), V2627, tmp19317, V2783)

																																			_ = tmp19318

																																			tmp19319 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp19320 := Call(__e, PrimNS2Value(symshen_4unbindv), V2627, V2783)

																																				_ = tmp19320

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp19321 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp19321

																																			tmp19322 := MakeNative(func(__e *ControlFlow) {
																																				tmp19323 := PrimCons(A, Nil)

																																				tmp19324 := PrimCons(symlist, tmp19323)

																																				__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19324, V2782, V2783, V2784)
																																				return

																																			}, 0)

																																			tmp19325 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19322)

																																			__e.TailApply(tmp19319, tmp19325)
																																			return

																																		}, 1)

																																		tmp19326 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																		__e.TailApply(tmp19316, tmp19326)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp19350 := PrimTail(V2625)

																															tmp19351 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19350, V2783)

																															__e.TailApply(tmp19313, tmp19351)
																															return

																														} else {
																															tmp19312 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2626)

																															if True == tmp19312 {
																																tmp19269 := Call(__e, PrimNS2Value(symshen_4bindv), V2626, symlist, V2783)

																																_ = tmp19269

																																tmp19270 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp19271 := Call(__e, PrimNS2Value(symshen_4unbindv), V2626, V2783)

																																	_ = tmp19271

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp19272 := MakeNative(func(__e *ControlFlow) {
																																	V2629 := __e.Get(1)
																																	_ = V2629
																																	tmp19308 := PrimIsPair(V2629)

																																	if True == tmp19308 {
																																		tmp19287 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp19288 := MakeNative(func(__e *ControlFlow) {
																																				V2630 := __e.Get(1)
																																				_ = V2630
																																				tmp19304 := PrimEqual(Nil, V2630)

																																				if True == tmp19304 {
																																					tmp19300 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp19300

																																					tmp19301 := MakeNative(func(__e *ControlFlow) {
																																						tmp19302 := PrimCons(A, Nil)

																																						tmp19303 := PrimCons(symlist, tmp19302)

																																						__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19303, V2782, V2783, V2784)
																																						return

																																					}, 0)

																																					__e.TailApply(PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19301)
																																					return

																																				} else {
																																					tmp19299 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2630)

																																					if True == tmp19299 {
																																						tmp19291 := Call(__e, PrimNS2Value(symshen_4bindv), V2630, Nil, V2783)

																																						_ = tmp19291

																																						tmp19292 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp19293 := Call(__e, PrimNS2Value(symshen_4unbindv), V2630, V2783)

																																							_ = tmp19293

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp19294 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																						_ = tmp19294

																																						tmp19295 := MakeNative(func(__e *ControlFlow) {
																																							tmp19296 := PrimCons(A, Nil)

																																							tmp19297 := PrimCons(symlist, tmp19296)

																																							__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19297, V2782, V2783, V2784)
																																							return

																																						}, 0)

																																						tmp19298 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19295)

																																						__e.TailApply(tmp19292, tmp19298)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp19305 := PrimTail(V2629)

																																			tmp19306 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19305, V2783)

																																			__e.TailApply(tmp19288, tmp19306)
																																			return

																																		}, 1)

																																		tmp19307 := PrimHead(V2629)

																																		__e.TailApply(tmp19287, tmp19307)
																																		return

																																	} else {
																																		tmp19286 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2629)

																																		if True == tmp19286 {
																																			tmp19275 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp19276 := PrimCons(A, Nil)

																																				tmp19277 := Call(__e, PrimNS2Value(symshen_4bindv), V2629, tmp19276, V2783)

																																				_ = tmp19277

																																				tmp19278 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp19279 := Call(__e, PrimNS2Value(symshen_4unbindv), V2629, V2783)

																																					_ = tmp19279

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp19280 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19280

																																				tmp19281 := MakeNative(func(__e *ControlFlow) {
																																					tmp19282 := PrimCons(A, Nil)

																																					tmp19283 := PrimCons(symlist, tmp19282)

																																					__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19283, V2782, V2783, V2784)
																																					return

																																				}, 0)

																																				tmp19284 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19281)

																																				__e.TailApply(tmp19278, tmp19284)
																																				return

																																			}, 1)

																																			tmp19285 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																																			__e.TailApply(tmp19275, tmp19285)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp19309 := PrimTail(V2625)

																																tmp19310 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19309, V2783)

																																tmp19311 := Call(__e, tmp19272, tmp19310)

																																__e.TailApply(tmp19270, tmp19311)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp19353 := PrimHead(V2625)

																													tmp19354 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19353, V2783)

																													__e.TailApply(tmp19266, tmp19354)
																													return

																												} else {
																													tmp19265 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2625)

																													if True == tmp19265 {
																														tmp19253 := MakeNative(func(__e *ControlFlow) {
																															A := __e.Get(1)
																															_ = A
																															tmp19254 := PrimCons(A, Nil)

																															tmp19255 := PrimCons(symlist, tmp19254)

																															tmp19256 := Call(__e, PrimNS2Value(symshen_4bindv), V2625, tmp19255, V2783)

																															_ = tmp19256

																															tmp19257 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp19258 := Call(__e, PrimNS2Value(symshen_4unbindv), V2625, V2783)

																																_ = tmp19258

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp19259 := Call(__e, PrimNS2Value(symshen_4incinfs))

																															_ = tmp19259

																															tmp19260 := MakeNative(func(__e *ControlFlow) {
																																tmp19261 := PrimCons(A, Nil)

																																tmp19262 := PrimCons(symlist, tmp19261)

																																__e.TailApply(PrimNS2Value(symshen_4th_d), Y, tmp19262, V2782, V2783, V2784)
																																return

																															}, 0)

																															tmp19263 := Call(__e, PrimNS2Value(symshen_4th_d), X, A, V2782, V2783, tmp19260)

																															__e.TailApply(tmp19257, tmp19263)
																															return

																														}, 1)

																														tmp19264 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																														__e.TailApply(tmp19253, tmp19264)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}

																											}, 1)

																											tmp19356 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2781, V2783)

																											__e.TailApply(tmp19250, tmp19356)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp19358 := PrimTail(V2623)

																									tmp19359 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19358, V2783)

																									__e.TailApply(tmp19248, tmp19359)
																									return

																								}, 1)

																								tmp19360 := PrimHead(V2623)

																								__e.TailApply(tmp19247, tmp19360)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp19362 := PrimTail(V2622)

																						tmp19363 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19362, V2783)

																						__e.TailApply(tmp19245, tmp19363)
																						return

																					}, 1)

																					tmp19364 := PrimHead(V2622)

																					__e.TailApply(tmp19244, tmp19364)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp19366 := PrimTail(V2620)

																			tmp19367 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19366, V2783)

																			__e.TailApply(tmp19242, tmp19367)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp19369 := PrimHead(V2620)

																	tmp19370 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19369, V2783)

																	__e.TailApply(tmp19240, tmp19370)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp19372 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

															tmp19373 := Call(__e, tmp19238, tmp19372)

															__e.TailApply(tmp18187, tmp19373)
															return

														} else {
															__e.Return(Case)
															return
														}

													}, 1)

													tmp19375 := MakeNative(func(__e *ControlFlow) {
														V2617 := __e.Get(1)
														_ = V2617
														tmp19398 := PrimIsPair(V2617)

														if True == tmp19398 {
															tmp19377 := MakeNative(func(__e *ControlFlow) {
																F := __e.Get(1)
																_ = F
																tmp19378 := MakeNative(func(__e *ControlFlow) {
																	V2618 := __e.Get(1)
																	_ = V2618
																	tmp19394 := PrimIsPair(V2618)

																	if True == tmp19394 {
																		tmp19380 := MakeNative(func(__e *ControlFlow) {
																			X := __e.Get(1)
																			_ = X
																			tmp19381 := MakeNative(func(__e *ControlFlow) {
																				V2619 := __e.Get(1)
																				_ = V2619
																				tmp19390 := PrimEqual(Nil, V2619)

																				if True == tmp19390 {
																					tmp19383 := MakeNative(func(__e *ControlFlow) {
																						B := __e.Get(1)
																						_ = B
																						tmp19384 := Call(__e, PrimNS2Value(symshen_4incinfs))

																						_ = tmp19384

																						tmp19385 := PrimCons(V2781, Nil)

																						tmp19386 := PrimCons(sym_1_1_6, tmp19385)

																						tmp19387 := PrimCons(B, tmp19386)

																						tmp19388 := MakeNative(func(__e *ControlFlow) {
																							__e.TailApply(PrimNS2Value(symshen_4th_d), X, B, V2782, V2783, V2784)
																							return
																						}, 0)

																						__e.TailApply(PrimNS2Value(symshen_4th_d), F, tmp19387, V2782, V2783, tmp19388)
																						return

																					}, 1)

																					tmp19389 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

																					__e.TailApply(tmp19383, tmp19389)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp19391 := PrimTail(V2618)

																			tmp19392 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19391, V2783)

																			__e.TailApply(tmp19381, tmp19392)
																			return

																		}, 1)

																		tmp19393 := PrimHead(V2618)

																		__e.TailApply(tmp19380, tmp19393)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp19395 := PrimTail(V2617)

																tmp19396 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19395, V2783)

																__e.TailApply(tmp19378, tmp19396)
																return

															}, 1)

															tmp19397 := PrimHead(V2617)

															__e.TailApply(tmp19377, tmp19397)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp19399 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

													tmp19400 := Call(__e, tmp19375, tmp19399)

													__e.TailApply(tmp18185, tmp19400)
													return

												} else {
													__e.Return(Case)
													return
												}

											}, 1)

											tmp19402 := MakeNative(func(__e *ControlFlow) {
												V2615 := __e.Get(1)
												_ = V2615
												tmp19414 := PrimIsPair(V2615)

												if True == tmp19414 {
													tmp19404 := MakeNative(func(__e *ControlFlow) {
														F := __e.Get(1)
														_ = F
														tmp19405 := MakeNative(func(__e *ControlFlow) {
															V2616 := __e.Get(1)
															_ = V2616
															tmp19410 := PrimEqual(Nil, V2616)

															if True == tmp19410 {
																tmp19407 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp19407

																tmp19408 := PrimCons(V2781, Nil)

																tmp19409 := PrimCons(sym_1_1_6, tmp19408)

																__e.TailApply(PrimNS2Value(symshen_4th_d), F, tmp19409, V2782, V2783, V2784)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp19411 := PrimTail(V2615)

														tmp19412 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19411, V2783)

														__e.TailApply(tmp19405, tmp19412)
														return

													}, 1)

													tmp19413 := PrimHead(V2615)

													__e.TailApply(tmp19404, tmp19413)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp19415 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

											tmp19416 := Call(__e, tmp19402, tmp19415)

											__e.TailApply(tmp18183, tmp19416)
											return

										} else {
											__e.Return(Case)
											return
										}

									}, 1)

									tmp19418 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp19418

									tmp19419 := Call(__e, PrimNS2Value(symshen_4by__hypothesis), V2780, V2781, V2782, V2783, V2784)

									__e.TailApply(tmp18181, tmp19419)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp19421 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp19421

							tmp19422 := Call(__e, PrimNS2Value(symshen_4base), V2780, V2781, V2783, V2784)

							__e.TailApply(tmp18179, tmp19422)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp19424 := MakeNative(func(__e *ControlFlow) {
						F := __e.Get(1)
						_ = F
						tmp19425 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp19425

						tmp19426 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

						tmp19427 := Call(__e, PrimNS2Value(symshen_4typedf_2), tmp19426)

						tmp19428 := MakeNative(func(__e *ControlFlow) {
							tmp19429 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2780, V2783)

							tmp19430 := Call(__e, PrimNS2Value(symshen_4sigf), tmp19429)

							tmp19431 := MakeNative(func(__e *ControlFlow) {
								tmp19432 := PrimCons(V2781, Nil)

								tmp19433 := PrimCons(F, tmp19432)

								__e.TailApply(PrimNS2Value(symcall), tmp19433, V2783, V2784)
								return

							}, 0)

							__e.TailApply(PrimNS2Value(symbind), F, tmp19430, V2783, tmp19431)
							return

						}, 0)

						__e.TailApply(PrimNS2Value(symfwhen), tmp19427, V2783, tmp19428)
						return

					}, 1)

					tmp19434 := Call(__e, PrimNS2Value(symshen_4newpv), V2783)

					tmp19435 := Call(__e, tmp19424, tmp19434)

					__e.TailApply(tmp18177, tmp19435)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp19437 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp19437

			tmp19438 := PrimCons(V2781, Nil)

			tmp19439 := PrimCons(sym_h, tmp19438)

			tmp19440 := PrimCons(V2780, tmp19439)

			tmp19441 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(symfwhen), False, V2783, V2784)
				return
			}, 0)

			tmp19442 := Call(__e, PrimNS2Value(symshen_4show), tmp19440, V2782, V2783, tmp19441)

			tmp19443 := Call(__e, tmp18175, tmp19442)

			__e.TailApply(PrimNS2Value(symshen_4cutpoint), Throwcontrol, tmp19443)
			return

		}, 1)

		tmp19444 := Call(__e, PrimNS2Value(symshen_4catchpoint))

		__e.TailApply(tmp18174, tmp19444)
		return

	}, 5)

	tmp19445 := Call(__e, PrimNS1Value(symns2_1set), symshen_4th_d, tmp18173)

	_ = tmp19445

	tmp19446 := MakeNative(func(__e *ControlFlow) {
		V2789 := __e.Get(1)
		_ = V2789
		V2790 := __e.Get(2)
		_ = V2790
		V2791 := __e.Get(3)
		_ = V2791
		V2792 := __e.Get(4)
		_ = V2792
		tmp19447 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp20577 := PrimEqual(Case, False)

			if True == tmp20577 {
				tmp19449 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp20078 := PrimEqual(Case, False)

					if True == tmp20078 {
						tmp19451 := MakeNative(func(__e *ControlFlow) {
							Case := __e.Get(1)
							_ = Case
							tmp19614 := PrimEqual(Case, False)

							if True == tmp19614 {
								tmp19453 := MakeNative(func(__e *ControlFlow) {
									Case := __e.Get(1)
									_ = Case
									tmp19470 := PrimEqual(Case, False)

									if True == tmp19470 {
										tmp19455 := MakeNative(func(__e *ControlFlow) {
											V2610 := __e.Get(1)
											_ = V2610
											tmp19468 := PrimIsPair(V2610)

											if True == tmp19468 {
												tmp19457 := MakeNative(func(__e *ControlFlow) {
													X := __e.Get(1)
													_ = X
													tmp19458 := MakeNative(func(__e *ControlFlow) {
														Hyp := __e.Get(1)
														_ = Hyp
														tmp19459 := MakeNative(func(__e *ControlFlow) {
															NewHyps := __e.Get(1)
															_ = NewHyps
															tmp19460 := Call(__e, PrimNS2Value(symshen_4incinfs))

															_ = tmp19460

															tmp19461 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

															tmp19462 := Call(__e, PrimNS2Value(symshen_4lazyderef), NewHyps, V2791)

															tmp19463 := PrimCons(tmp19461, tmp19462)

															tmp19464 := MakeNative(func(__e *ControlFlow) {
																__e.TailApply(PrimNS2Value(symshen_4t_d_1hyps), Hyp, NewHyps, V2791, V2792)
																return
															}, 0)

															__e.TailApply(PrimNS2Value(symbind), V2790, tmp19463, V2791, tmp19464)
															return

														}, 1)

														tmp19465 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

														__e.TailApply(tmp19459, tmp19465)
														return

													}, 1)

													tmp19466 := PrimTail(V2610)

													__e.TailApply(tmp19458, tmp19466)
													return

												}, 1)

												tmp19467 := PrimHead(V2610)

												__e.TailApply(tmp19457, tmp19467)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp19469 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2789, V2791)

										__e.TailApply(tmp19455, tmp19469)
										return

									} else {
										__e.Return(Case)
										return
									}

								}, 1)

								tmp19471 := MakeNative(func(__e *ControlFlow) {
									V2597 := __e.Get(1)
									_ = V2597
									tmp19611 := PrimIsPair(V2597)

									if True == tmp19611 {
										tmp19473 := MakeNative(func(__e *ControlFlow) {
											V2598 := __e.Get(1)
											_ = V2598
											tmp19608 := PrimIsPair(V2598)

											if True == tmp19608 {
												tmp19475 := MakeNative(func(__e *ControlFlow) {
													V2599 := __e.Get(1)
													_ = V2599
													tmp19605 := PrimIsPair(V2599)

													if True == tmp19605 {
														tmp19477 := MakeNative(func(__e *ControlFlow) {
															V2600 := __e.Get(1)
															_ = V2600
															tmp19602 := PrimEqual(sym_8s, V2600)

															if True == tmp19602 {
																tmp19479 := MakeNative(func(__e *ControlFlow) {
																	V2601 := __e.Get(1)
																	_ = V2601
																	tmp19599 := PrimIsPair(V2601)

																	if True == tmp19599 {
																		tmp19481 := MakeNative(func(__e *ControlFlow) {
																			X := __e.Get(1)
																			_ = X
																			tmp19482 := MakeNative(func(__e *ControlFlow) {
																				V2602 := __e.Get(1)
																				_ = V2602
																				tmp19595 := PrimIsPair(V2602)

																				if True == tmp19595 {
																					tmp19484 := MakeNative(func(__e *ControlFlow) {
																						Y := __e.Get(1)
																						_ = Y
																						tmp19485 := MakeNative(func(__e *ControlFlow) {
																							V2603 := __e.Get(1)
																							_ = V2603
																							tmp19591 := PrimEqual(Nil, V2603)

																							if True == tmp19591 {
																								tmp19487 := MakeNative(func(__e *ControlFlow) {
																									V2604 := __e.Get(1)
																									_ = V2604
																									tmp19588 := PrimIsPair(V2604)

																									if True == tmp19588 {
																										tmp19489 := MakeNative(func(__e *ControlFlow) {
																											V2605 := __e.Get(1)
																											_ = V2605
																											tmp19585 := PrimEqual(sym_h, V2605)

																											if True == tmp19585 {
																												tmp19491 := MakeNative(func(__e *ControlFlow) {
																													V2606 := __e.Get(1)
																													_ = V2606
																													tmp19582 := PrimIsPair(V2606)

																													if True == tmp19582 {
																														tmp19493 := MakeNative(func(__e *ControlFlow) {
																															V2607 := __e.Get(1)
																															_ = V2607
																															tmp19579 := PrimEqual(symstring, V2607)

																															if True == tmp19579 {
																																tmp19540 := MakeNative(func(__e *ControlFlow) {
																																	V2608 := __e.Get(1)
																																	_ = V2608
																																	tmp19576 := PrimEqual(Nil, V2608)

																																	if True == tmp19576 {
																																		tmp19562 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp19563 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp19563

																																			tmp19564 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																			tmp19565 := PrimCons(symstring, Nil)

																																			tmp19566 := PrimCons(sym_h, tmp19565)

																																			tmp19567 := PrimCons(tmp19564, tmp19566)

																																			tmp19568 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																			tmp19569 := PrimCons(symstring, Nil)

																																			tmp19570 := PrimCons(sym_h, tmp19569)

																																			tmp19571 := PrimCons(tmp19568, tmp19570)

																																			tmp19572 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																			tmp19573 := PrimCons(tmp19571, tmp19572)

																																			tmp19574 := PrimCons(tmp19567, tmp19573)

																																			__e.TailApply(PrimNS2Value(symbind), V2790, tmp19574, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp19575 := PrimTail(V2597)

																																		__e.TailApply(tmp19562, tmp19575)
																																		return

																																	} else {
																																		tmp19561 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2608)

																																		if True == tmp19561 {
																																			tmp19543 := Call(__e, PrimNS2Value(symshen_4bindv), V2608, Nil, V2791)

																																			_ = tmp19543

																																			tmp19544 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp19545 := Call(__e, PrimNS2Value(symshen_4unbindv), V2608, V2791)

																																				_ = tmp19545

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp19546 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp19547 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19547

																																				tmp19548 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp19549 := PrimCons(symstring, Nil)

																																				tmp19550 := PrimCons(sym_h, tmp19549)

																																				tmp19551 := PrimCons(tmp19548, tmp19550)

																																				tmp19552 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp19553 := PrimCons(symstring, Nil)

																																				tmp19554 := PrimCons(sym_h, tmp19553)

																																				tmp19555 := PrimCons(tmp19552, tmp19554)

																																				tmp19556 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp19557 := PrimCons(tmp19555, tmp19556)

																																				tmp19558 := PrimCons(tmp19551, tmp19557)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp19558, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp19559 := PrimTail(V2597)

																																			tmp19560 := Call(__e, tmp19546, tmp19559)

																																			__e.TailApply(tmp19544, tmp19560)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp19577 := PrimTail(V2606)

																																tmp19578 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19577, V2791)

																																__e.TailApply(tmp19540, tmp19578)
																																return

																															} else {
																																tmp19539 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2607)

																																if True == tmp19539 {
																																	tmp19496 := Call(__e, PrimNS2Value(symshen_4bindv), V2607, symstring, V2791)

																																	_ = tmp19496

																																	tmp19497 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp19498 := Call(__e, PrimNS2Value(symshen_4unbindv), V2607, V2791)

																																		_ = tmp19498

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp19499 := MakeNative(func(__e *ControlFlow) {
																																		V2609 := __e.Get(1)
																																		_ = V2609
																																		tmp19535 := PrimEqual(Nil, V2609)

																																		if True == tmp19535 {
																																			tmp19521 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp19522 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19522

																																				tmp19523 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp19524 := PrimCons(symstring, Nil)

																																				tmp19525 := PrimCons(sym_h, tmp19524)

																																				tmp19526 := PrimCons(tmp19523, tmp19525)

																																				tmp19527 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp19528 := PrimCons(symstring, Nil)

																																				tmp19529 := PrimCons(sym_h, tmp19528)

																																				tmp19530 := PrimCons(tmp19527, tmp19529)

																																				tmp19531 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp19532 := PrimCons(tmp19530, tmp19531)

																																				tmp19533 := PrimCons(tmp19526, tmp19532)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp19533, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp19534 := PrimTail(V2597)

																																			__e.TailApply(tmp19521, tmp19534)
																																			return

																																		} else {
																																			tmp19520 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2609)

																																			if True == tmp19520 {
																																				tmp19502 := Call(__e, PrimNS2Value(symshen_4bindv), V2609, Nil, V2791)

																																				_ = tmp19502

																																				tmp19503 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp19504 := Call(__e, PrimNS2Value(symshen_4unbindv), V2609, V2791)

																																					_ = tmp19504

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp19505 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp19506 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp19506

																																					tmp19507 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp19508 := PrimCons(symstring, Nil)

																																					tmp19509 := PrimCons(sym_h, tmp19508)

																																					tmp19510 := PrimCons(tmp19507, tmp19509)

																																					tmp19511 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp19512 := PrimCons(symstring, Nil)

																																					tmp19513 := PrimCons(sym_h, tmp19512)

																																					tmp19514 := PrimCons(tmp19511, tmp19513)

																																					tmp19515 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp19516 := PrimCons(tmp19514, tmp19515)

																																					tmp19517 := PrimCons(tmp19510, tmp19516)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp19517, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp19518 := PrimTail(V2597)

																																				tmp19519 := Call(__e, tmp19505, tmp19518)

																																				__e.TailApply(tmp19503, tmp19519)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp19536 := PrimTail(V2606)

																																	tmp19537 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19536, V2791)

																																	tmp19538 := Call(__e, tmp19499, tmp19537)

																																	__e.TailApply(tmp19497, tmp19538)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}

																														}, 1)

																														tmp19580 := PrimHead(V2606)

																														tmp19581 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19580, V2791)

																														__e.TailApply(tmp19493, tmp19581)
																														return

																													} else {
																														__e.Return(False)
																														return
																													}

																												}, 1)

																												tmp19583 := PrimTail(V2604)

																												tmp19584 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19583, V2791)

																												__e.TailApply(tmp19491, tmp19584)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp19586 := PrimHead(V2604)

																										tmp19587 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19586, V2791)

																										__e.TailApply(tmp19489, tmp19587)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp19589 := PrimTail(V2598)

																								tmp19590 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19589, V2791)

																								__e.TailApply(tmp19487, tmp19590)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp19592 := PrimTail(V2602)

																						tmp19593 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19592, V2791)

																						__e.TailApply(tmp19485, tmp19593)
																						return

																					}, 1)

																					tmp19594 := PrimHead(V2602)

																					__e.TailApply(tmp19484, tmp19594)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}, 1)

																			tmp19596 := PrimTail(V2601)

																			tmp19597 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19596, V2791)

																			__e.TailApply(tmp19482, tmp19597)
																			return

																		}, 1)

																		tmp19598 := PrimHead(V2601)

																		__e.TailApply(tmp19481, tmp19598)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp19600 := PrimTail(V2599)

																tmp19601 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19600, V2791)

																__e.TailApply(tmp19479, tmp19601)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp19603 := PrimHead(V2599)

														tmp19604 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19603, V2791)

														__e.TailApply(tmp19477, tmp19604)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp19606 := PrimHead(V2598)

												tmp19607 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19606, V2791)

												__e.TailApply(tmp19475, tmp19607)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp19609 := PrimHead(V2597)

										tmp19610 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19609, V2791)

										__e.TailApply(tmp19473, tmp19610)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp19612 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2789, V2791)

								tmp19613 := Call(__e, tmp19471, tmp19612)

								__e.TailApply(tmp19453, tmp19613)
								return

							} else {
								__e.Return(Case)
								return
							}

						}, 1)

						tmp19615 := MakeNative(func(__e *ControlFlow) {
							V2574 := __e.Get(1)
							_ = V2574
							tmp20075 := PrimIsPair(V2574)

							if True == tmp20075 {
								tmp19617 := MakeNative(func(__e *ControlFlow) {
									V2575 := __e.Get(1)
									_ = V2575
									tmp20072 := PrimIsPair(V2575)

									if True == tmp20072 {
										tmp19619 := MakeNative(func(__e *ControlFlow) {
											V2576 := __e.Get(1)
											_ = V2576
											tmp20069 := PrimIsPair(V2576)

											if True == tmp20069 {
												tmp19621 := MakeNative(func(__e *ControlFlow) {
													V2577 := __e.Get(1)
													_ = V2577
													tmp20066 := PrimEqual(sym_8v, V2577)

													if True == tmp20066 {
														tmp19623 := MakeNative(func(__e *ControlFlow) {
															V2578 := __e.Get(1)
															_ = V2578
															tmp20063 := PrimIsPair(V2578)

															if True == tmp20063 {
																tmp19625 := MakeNative(func(__e *ControlFlow) {
																	X := __e.Get(1)
																	_ = X
																	tmp19626 := MakeNative(func(__e *ControlFlow) {
																		V2579 := __e.Get(1)
																		_ = V2579
																		tmp20059 := PrimIsPair(V2579)

																		if True == tmp20059 {
																			tmp19628 := MakeNative(func(__e *ControlFlow) {
																				Y := __e.Get(1)
																				_ = Y
																				tmp19629 := MakeNative(func(__e *ControlFlow) {
																					V2580 := __e.Get(1)
																					_ = V2580
																					tmp20055 := PrimEqual(Nil, V2580)

																					if True == tmp20055 {
																						tmp19631 := MakeNative(func(__e *ControlFlow) {
																							V2581 := __e.Get(1)
																							_ = V2581
																							tmp20052 := PrimIsPair(V2581)

																							if True == tmp20052 {
																								tmp19633 := MakeNative(func(__e *ControlFlow) {
																									V2582 := __e.Get(1)
																									_ = V2582
																									tmp20049 := PrimEqual(sym_h, V2582)

																									if True == tmp20049 {
																										tmp19635 := MakeNative(func(__e *ControlFlow) {
																											V2583 := __e.Get(1)
																											_ = V2583
																											tmp20046 := PrimIsPair(V2583)

																											if True == tmp20046 {
																												tmp19637 := MakeNative(func(__e *ControlFlow) {
																													V2584 := __e.Get(1)
																													_ = V2584
																													tmp20043 := PrimIsPair(V2584)

																													if True == tmp20043 {
																														tmp19696 := MakeNative(func(__e *ControlFlow) {
																															V2585 := __e.Get(1)
																															_ = V2585
																															tmp20040 := PrimEqual(symvector, V2585)

																															if True == tmp20040 {
																																tmp19872 := MakeNative(func(__e *ControlFlow) {
																																	V2586 := __e.Get(1)
																																	_ = V2586
																																	tmp20037 := PrimIsPair(V2586)

																																	if True == tmp20037 {
																																		tmp19930 := MakeNative(func(__e *ControlFlow) {
																																			A := __e.Get(1)
																																			_ = A
																																			tmp19931 := MakeNative(func(__e *ControlFlow) {
																																				V2587 := __e.Get(1)
																																				_ = V2587
																																				tmp20033 := PrimEqual(Nil, V2587)

																																				if True == tmp20033 {
																																					tmp19986 := MakeNative(func(__e *ControlFlow) {
																																						V2588 := __e.Get(1)
																																						_ = V2588
																																						tmp20030 := PrimEqual(Nil, V2588)

																																						if True == tmp20030 {
																																							tmp20012 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp20013 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp20013

																																								tmp20014 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																								tmp20015 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp20016 := PrimCons(tmp20015, Nil)

																																								tmp20017 := PrimCons(sym_h, tmp20016)

																																								tmp20018 := PrimCons(tmp20014, tmp20017)

																																								tmp20019 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																								tmp20020 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp20021 := PrimCons(tmp20020, Nil)

																																								tmp20022 := PrimCons(symvector, tmp20021)

																																								tmp20023 := PrimCons(tmp20022, Nil)

																																								tmp20024 := PrimCons(sym_h, tmp20023)

																																								tmp20025 := PrimCons(tmp20019, tmp20024)

																																								tmp20026 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																								tmp20027 := PrimCons(tmp20025, tmp20026)

																																								tmp20028 := PrimCons(tmp20018, tmp20027)

																																								__e.TailApply(PrimNS2Value(symbind), V2790, tmp20028, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp20029 := PrimTail(V2574)

																																							__e.TailApply(tmp20012, tmp20029)
																																							return

																																						} else {
																																							tmp20011 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2588)

																																							if True == tmp20011 {
																																								tmp19989 := Call(__e, PrimNS2Value(symshen_4bindv), V2588, Nil, V2791)

																																								_ = tmp19989

																																								tmp19990 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp19991 := Call(__e, PrimNS2Value(symshen_4unbindv), V2588, V2791)

																																									_ = tmp19991

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp19992 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp19993 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19993

																																									tmp19994 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp19995 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19996 := PrimCons(tmp19995, Nil)

																																									tmp19997 := PrimCons(sym_h, tmp19996)

																																									tmp19998 := PrimCons(tmp19994, tmp19997)

																																									tmp19999 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp20000 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp20001 := PrimCons(tmp20000, Nil)

																																									tmp20002 := PrimCons(symvector, tmp20001)

																																									tmp20003 := PrimCons(tmp20002, Nil)

																																									tmp20004 := PrimCons(sym_h, tmp20003)

																																									tmp20005 := PrimCons(tmp19999, tmp20004)

																																									tmp20006 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp20007 := PrimCons(tmp20005, tmp20006)

																																									tmp20008 := PrimCons(tmp19998, tmp20007)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp20008, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp20009 := PrimTail(V2574)

																																								tmp20010 := Call(__e, tmp19992, tmp20009)

																																								__e.TailApply(tmp19990, tmp20010)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp20031 := PrimTail(V2583)

																																					tmp20032 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20031, V2791)

																																					__e.TailApply(tmp19986, tmp20032)
																																					return

																																				} else {
																																					tmp19985 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2587)

																																					if True == tmp19985 {
																																						tmp19934 := Call(__e, PrimNS2Value(symshen_4bindv), V2587, Nil, V2791)

																																						_ = tmp19934

																																						tmp19935 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp19936 := Call(__e, PrimNS2Value(symshen_4unbindv), V2587, V2791)

																																							_ = tmp19936

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp19937 := MakeNative(func(__e *ControlFlow) {
																																							V2589 := __e.Get(1)
																																							_ = V2589
																																							tmp19981 := PrimEqual(Nil, V2589)

																																							if True == tmp19981 {
																																								tmp19963 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp19964 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19964

																																									tmp19965 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp19966 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19967 := PrimCons(tmp19966, Nil)

																																									tmp19968 := PrimCons(sym_h, tmp19967)

																																									tmp19969 := PrimCons(tmp19965, tmp19968)

																																									tmp19970 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp19971 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19972 := PrimCons(tmp19971, Nil)

																																									tmp19973 := PrimCons(symvector, tmp19972)

																																									tmp19974 := PrimCons(tmp19973, Nil)

																																									tmp19975 := PrimCons(sym_h, tmp19974)

																																									tmp19976 := PrimCons(tmp19970, tmp19975)

																																									tmp19977 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp19978 := PrimCons(tmp19976, tmp19977)

																																									tmp19979 := PrimCons(tmp19969, tmp19978)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp19979, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp19980 := PrimTail(V2574)

																																								__e.TailApply(tmp19963, tmp19980)
																																								return

																																							} else {
																																								tmp19962 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2589)

																																								if True == tmp19962 {
																																									tmp19940 := Call(__e, PrimNS2Value(symshen_4bindv), V2589, Nil, V2791)

																																									_ = tmp19940

																																									tmp19941 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp19942 := Call(__e, PrimNS2Value(symshen_4unbindv), V2589, V2791)

																																										_ = tmp19942

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp19943 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp19944 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19944

																																										tmp19945 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp19946 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19947 := PrimCons(tmp19946, Nil)

																																										tmp19948 := PrimCons(sym_h, tmp19947)

																																										tmp19949 := PrimCons(tmp19945, tmp19948)

																																										tmp19950 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp19951 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19952 := PrimCons(tmp19951, Nil)

																																										tmp19953 := PrimCons(symvector, tmp19952)

																																										tmp19954 := PrimCons(tmp19953, Nil)

																																										tmp19955 := PrimCons(sym_h, tmp19954)

																																										tmp19956 := PrimCons(tmp19950, tmp19955)

																																										tmp19957 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp19958 := PrimCons(tmp19956, tmp19957)

																																										tmp19959 := PrimCons(tmp19949, tmp19958)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp19959, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp19960 := PrimTail(V2574)

																																									tmp19961 := Call(__e, tmp19943, tmp19960)

																																									__e.TailApply(tmp19941, tmp19961)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp19982 := PrimTail(V2583)

																																						tmp19983 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19982, V2791)

																																						tmp19984 := Call(__e, tmp19937, tmp19983)

																																						__e.TailApply(tmp19935, tmp19984)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp20034 := PrimTail(V2586)

																																			tmp20035 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20034, V2791)

																																			__e.TailApply(tmp19931, tmp20035)
																																			return

																																		}, 1)

																																		tmp20036 := PrimHead(V2586)

																																		__e.TailApply(tmp19930, tmp20036)
																																		return

																																	} else {
																																		tmp19929 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2586)

																																		if True == tmp19929 {
																																			tmp19875 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp19876 := PrimCons(A, Nil)

																																				tmp19877 := Call(__e, PrimNS2Value(symshen_4bindv), V2586, tmp19876, V2791)

																																				_ = tmp19877

																																				tmp19878 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp19879 := Call(__e, PrimNS2Value(symshen_4unbindv), V2586, V2791)

																																					_ = tmp19879

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp19880 := MakeNative(func(__e *ControlFlow) {
																																					V2590 := __e.Get(1)
																																					_ = V2590
																																					tmp19924 := PrimEqual(Nil, V2590)

																																					if True == tmp19924 {
																																						tmp19906 := MakeNative(func(__e *ControlFlow) {
																																							Hyp := __e.Get(1)
																																							_ = Hyp
																																							tmp19907 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																							_ = tmp19907

																																							tmp19908 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																							tmp19909 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																							tmp19910 := PrimCons(tmp19909, Nil)

																																							tmp19911 := PrimCons(sym_h, tmp19910)

																																							tmp19912 := PrimCons(tmp19908, tmp19911)

																																							tmp19913 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																							tmp19914 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																							tmp19915 := PrimCons(tmp19914, Nil)

																																							tmp19916 := PrimCons(symvector, tmp19915)

																																							tmp19917 := PrimCons(tmp19916, Nil)

																																							tmp19918 := PrimCons(sym_h, tmp19917)

																																							tmp19919 := PrimCons(tmp19913, tmp19918)

																																							tmp19920 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																							tmp19921 := PrimCons(tmp19919, tmp19920)

																																							tmp19922 := PrimCons(tmp19912, tmp19921)

																																							__e.TailApply(PrimNS2Value(symbind), V2790, tmp19922, V2791, V2792)
																																							return

																																						}, 1)

																																						tmp19923 := PrimTail(V2574)

																																						__e.TailApply(tmp19906, tmp19923)
																																						return

																																					} else {
																																						tmp19905 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2590)

																																						if True == tmp19905 {
																																							tmp19883 := Call(__e, PrimNS2Value(symshen_4bindv), V2590, Nil, V2791)

																																							_ = tmp19883

																																							tmp19884 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp19885 := Call(__e, PrimNS2Value(symshen_4unbindv), V2590, V2791)

																																								_ = tmp19885

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp19886 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp19887 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp19887

																																								tmp19888 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																								tmp19889 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp19890 := PrimCons(tmp19889, Nil)

																																								tmp19891 := PrimCons(sym_h, tmp19890)

																																								tmp19892 := PrimCons(tmp19888, tmp19891)

																																								tmp19893 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																								tmp19894 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp19895 := PrimCons(tmp19894, Nil)

																																								tmp19896 := PrimCons(symvector, tmp19895)

																																								tmp19897 := PrimCons(tmp19896, Nil)

																																								tmp19898 := PrimCons(sym_h, tmp19897)

																																								tmp19899 := PrimCons(tmp19893, tmp19898)

																																								tmp19900 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																								tmp19901 := PrimCons(tmp19899, tmp19900)

																																								tmp19902 := PrimCons(tmp19892, tmp19901)

																																								__e.TailApply(PrimNS2Value(symbind), V2790, tmp19902, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp19903 := PrimTail(V2574)

																																							tmp19904 := Call(__e, tmp19886, tmp19903)

																																							__e.TailApply(tmp19884, tmp19904)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp19925 := PrimTail(V2583)

																																				tmp19926 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19925, V2791)

																																				tmp19927 := Call(__e, tmp19880, tmp19926)

																																				__e.TailApply(tmp19878, tmp19927)
																																				return

																																			}, 1)

																																			tmp19928 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																			__e.TailApply(tmp19875, tmp19928)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp20038 := PrimTail(V2584)

																																tmp20039 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20038, V2791)

																																__e.TailApply(tmp19872, tmp20039)
																																return

																															} else {
																																tmp19871 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2585)

																																if True == tmp19871 {
																																	tmp19699 := Call(__e, PrimNS2Value(symshen_4bindv), V2585, symvector, V2791)

																																	_ = tmp19699

																																	tmp19700 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp19701 := Call(__e, PrimNS2Value(symshen_4unbindv), V2585, V2791)

																																		_ = tmp19701

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp19702 := MakeNative(func(__e *ControlFlow) {
																																		V2591 := __e.Get(1)
																																		_ = V2591
																																		tmp19867 := PrimIsPair(V2591)

																																		if True == tmp19867 {
																																			tmp19760 := MakeNative(func(__e *ControlFlow) {
																																				A := __e.Get(1)
																																				_ = A
																																				tmp19761 := MakeNative(func(__e *ControlFlow) {
																																					V2592 := __e.Get(1)
																																					_ = V2592
																																					tmp19863 := PrimEqual(Nil, V2592)

																																					if True == tmp19863 {
																																						tmp19816 := MakeNative(func(__e *ControlFlow) {
																																							V2593 := __e.Get(1)
																																							_ = V2593
																																							tmp19860 := PrimEqual(Nil, V2593)

																																							if True == tmp19860 {
																																								tmp19842 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp19843 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19843

																																									tmp19844 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp19845 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19846 := PrimCons(tmp19845, Nil)

																																									tmp19847 := PrimCons(sym_h, tmp19846)

																																									tmp19848 := PrimCons(tmp19844, tmp19847)

																																									tmp19849 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp19850 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19851 := PrimCons(tmp19850, Nil)

																																									tmp19852 := PrimCons(symvector, tmp19851)

																																									tmp19853 := PrimCons(tmp19852, Nil)

																																									tmp19854 := PrimCons(sym_h, tmp19853)

																																									tmp19855 := PrimCons(tmp19849, tmp19854)

																																									tmp19856 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp19857 := PrimCons(tmp19855, tmp19856)

																																									tmp19858 := PrimCons(tmp19848, tmp19857)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp19858, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp19859 := PrimTail(V2574)

																																								__e.TailApply(tmp19842, tmp19859)
																																								return

																																							} else {
																																								tmp19841 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2593)

																																								if True == tmp19841 {
																																									tmp19819 := Call(__e, PrimNS2Value(symshen_4bindv), V2593, Nil, V2791)

																																									_ = tmp19819

																																									tmp19820 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp19821 := Call(__e, PrimNS2Value(symshen_4unbindv), V2593, V2791)

																																										_ = tmp19821

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp19822 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp19823 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19823

																																										tmp19824 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp19825 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19826 := PrimCons(tmp19825, Nil)

																																										tmp19827 := PrimCons(sym_h, tmp19826)

																																										tmp19828 := PrimCons(tmp19824, tmp19827)

																																										tmp19829 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp19830 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19831 := PrimCons(tmp19830, Nil)

																																										tmp19832 := PrimCons(symvector, tmp19831)

																																										tmp19833 := PrimCons(tmp19832, Nil)

																																										tmp19834 := PrimCons(sym_h, tmp19833)

																																										tmp19835 := PrimCons(tmp19829, tmp19834)

																																										tmp19836 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp19837 := PrimCons(tmp19835, tmp19836)

																																										tmp19838 := PrimCons(tmp19828, tmp19837)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp19838, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp19839 := PrimTail(V2574)

																																									tmp19840 := Call(__e, tmp19822, tmp19839)

																																									__e.TailApply(tmp19820, tmp19840)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp19861 := PrimTail(V2583)

																																						tmp19862 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19861, V2791)

																																						__e.TailApply(tmp19816, tmp19862)
																																						return

																																					} else {
																																						tmp19815 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2592)

																																						if True == tmp19815 {
																																							tmp19764 := Call(__e, PrimNS2Value(symshen_4bindv), V2592, Nil, V2791)

																																							_ = tmp19764

																																							tmp19765 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp19766 := Call(__e, PrimNS2Value(symshen_4unbindv), V2592, V2791)

																																								_ = tmp19766

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp19767 := MakeNative(func(__e *ControlFlow) {
																																								V2594 := __e.Get(1)
																																								_ = V2594
																																								tmp19811 := PrimEqual(Nil, V2594)

																																								if True == tmp19811 {
																																									tmp19793 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp19794 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp19794

																																										tmp19795 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp19796 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19797 := PrimCons(tmp19796, Nil)

																																										tmp19798 := PrimCons(sym_h, tmp19797)

																																										tmp19799 := PrimCons(tmp19795, tmp19798)

																																										tmp19800 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp19801 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp19802 := PrimCons(tmp19801, Nil)

																																										tmp19803 := PrimCons(symvector, tmp19802)

																																										tmp19804 := PrimCons(tmp19803, Nil)

																																										tmp19805 := PrimCons(sym_h, tmp19804)

																																										tmp19806 := PrimCons(tmp19800, tmp19805)

																																										tmp19807 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp19808 := PrimCons(tmp19806, tmp19807)

																																										tmp19809 := PrimCons(tmp19799, tmp19808)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp19809, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp19810 := PrimTail(V2574)

																																									__e.TailApply(tmp19793, tmp19810)
																																									return

																																								} else {
																																									tmp19792 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2594)

																																									if True == tmp19792 {
																																										tmp19770 := Call(__e, PrimNS2Value(symshen_4bindv), V2594, Nil, V2791)

																																										_ = tmp19770

																																										tmp19771 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp19772 := Call(__e, PrimNS2Value(symshen_4unbindv), V2594, V2791)

																																											_ = tmp19772

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp19773 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp19774 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp19774

																																											tmp19775 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																											tmp19776 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																											tmp19777 := PrimCons(tmp19776, Nil)

																																											tmp19778 := PrimCons(sym_h, tmp19777)

																																											tmp19779 := PrimCons(tmp19775, tmp19778)

																																											tmp19780 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																											tmp19781 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																											tmp19782 := PrimCons(tmp19781, Nil)

																																											tmp19783 := PrimCons(symvector, tmp19782)

																																											tmp19784 := PrimCons(tmp19783, Nil)

																																											tmp19785 := PrimCons(sym_h, tmp19784)

																																											tmp19786 := PrimCons(tmp19780, tmp19785)

																																											tmp19787 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																											tmp19788 := PrimCons(tmp19786, tmp19787)

																																											tmp19789 := PrimCons(tmp19779, tmp19788)

																																											__e.TailApply(PrimNS2Value(symbind), V2790, tmp19789, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp19790 := PrimTail(V2574)

																																										tmp19791 := Call(__e, tmp19773, tmp19790)

																																										__e.TailApply(tmp19771, tmp19791)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp19812 := PrimTail(V2583)

																																							tmp19813 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19812, V2791)

																																							tmp19814 := Call(__e, tmp19767, tmp19813)

																																							__e.TailApply(tmp19765, tmp19814)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp19864 := PrimTail(V2591)

																																				tmp19865 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19864, V2791)

																																				__e.TailApply(tmp19761, tmp19865)
																																				return

																																			}, 1)

																																			tmp19866 := PrimHead(V2591)

																																			__e.TailApply(tmp19760, tmp19866)
																																			return

																																		} else {
																																			tmp19759 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2591)

																																			if True == tmp19759 {
																																				tmp19705 := MakeNative(func(__e *ControlFlow) {
																																					A := __e.Get(1)
																																					_ = A
																																					tmp19706 := PrimCons(A, Nil)

																																					tmp19707 := Call(__e, PrimNS2Value(symshen_4bindv), V2591, tmp19706, V2791)

																																					_ = tmp19707

																																					tmp19708 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp19709 := Call(__e, PrimNS2Value(symshen_4unbindv), V2591, V2791)

																																						_ = tmp19709

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp19710 := MakeNative(func(__e *ControlFlow) {
																																						V2595 := __e.Get(1)
																																						_ = V2595
																																						tmp19754 := PrimEqual(Nil, V2595)

																																						if True == tmp19754 {
																																							tmp19736 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp19737 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp19737

																																								tmp19738 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																								tmp19739 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp19740 := PrimCons(tmp19739, Nil)

																																								tmp19741 := PrimCons(sym_h, tmp19740)

																																								tmp19742 := PrimCons(tmp19738, tmp19741)

																																								tmp19743 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																								tmp19744 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp19745 := PrimCons(tmp19744, Nil)

																																								tmp19746 := PrimCons(symvector, tmp19745)

																																								tmp19747 := PrimCons(tmp19746, Nil)

																																								tmp19748 := PrimCons(sym_h, tmp19747)

																																								tmp19749 := PrimCons(tmp19743, tmp19748)

																																								tmp19750 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																								tmp19751 := PrimCons(tmp19749, tmp19750)

																																								tmp19752 := PrimCons(tmp19742, tmp19751)

																																								__e.TailApply(PrimNS2Value(symbind), V2790, tmp19752, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp19753 := PrimTail(V2574)

																																							__e.TailApply(tmp19736, tmp19753)
																																							return

																																						} else {
																																							tmp19735 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2595)

																																							if True == tmp19735 {
																																								tmp19713 := Call(__e, PrimNS2Value(symshen_4bindv), V2595, Nil, V2791)

																																								_ = tmp19713

																																								tmp19714 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp19715 := Call(__e, PrimNS2Value(symshen_4unbindv), V2595, V2791)

																																									_ = tmp19715

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp19716 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp19717 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp19717

																																									tmp19718 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp19719 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19720 := PrimCons(tmp19719, Nil)

																																									tmp19721 := PrimCons(sym_h, tmp19720)

																																									tmp19722 := PrimCons(tmp19718, tmp19721)

																																									tmp19723 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp19724 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp19725 := PrimCons(tmp19724, Nil)

																																									tmp19726 := PrimCons(symvector, tmp19725)

																																									tmp19727 := PrimCons(tmp19726, Nil)

																																									tmp19728 := PrimCons(sym_h, tmp19727)

																																									tmp19729 := PrimCons(tmp19723, tmp19728)

																																									tmp19730 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp19731 := PrimCons(tmp19729, tmp19730)

																																									tmp19732 := PrimCons(tmp19722, tmp19731)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp19732, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp19733 := PrimTail(V2574)

																																								tmp19734 := Call(__e, tmp19716, tmp19733)

																																								__e.TailApply(tmp19714, tmp19734)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp19755 := PrimTail(V2583)

																																					tmp19756 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19755, V2791)

																																					tmp19757 := Call(__e, tmp19710, tmp19756)

																																					__e.TailApply(tmp19708, tmp19757)
																																					return

																																				}, 1)

																																				tmp19758 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																				__e.TailApply(tmp19705, tmp19758)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp19868 := PrimTail(V2584)

																																	tmp19869 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19868, V2791)

																																	tmp19870 := Call(__e, tmp19702, tmp19869)

																																	__e.TailApply(tmp19700, tmp19870)
																																	return

																																} else {
																																	__e.Return(False)
																																	return
																																}

																															}

																														}, 1)

																														tmp20041 := PrimHead(V2584)

																														tmp20042 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20041, V2791)

																														__e.TailApply(tmp19696, tmp20042)
																														return

																													} else {
																														tmp19695 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2584)

																														if True == tmp19695 {
																															tmp19640 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp19641 := PrimCons(A, Nil)

																																tmp19642 := PrimCons(symvector, tmp19641)

																																tmp19643 := Call(__e, PrimNS2Value(symshen_4bindv), V2584, tmp19642, V2791)

																																_ = tmp19643

																																tmp19644 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp19645 := Call(__e, PrimNS2Value(symshen_4unbindv), V2584, V2791)

																																	_ = tmp19645

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp19646 := MakeNative(func(__e *ControlFlow) {
																																	V2596 := __e.Get(1)
																																	_ = V2596
																																	tmp19690 := PrimEqual(Nil, V2596)

																																	if True == tmp19690 {
																																		tmp19672 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp19673 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp19673

																																			tmp19674 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																			tmp19675 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																			tmp19676 := PrimCons(tmp19675, Nil)

																																			tmp19677 := PrimCons(sym_h, tmp19676)

																																			tmp19678 := PrimCons(tmp19674, tmp19677)

																																			tmp19679 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																			tmp19680 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																			tmp19681 := PrimCons(tmp19680, Nil)

																																			tmp19682 := PrimCons(symvector, tmp19681)

																																			tmp19683 := PrimCons(tmp19682, Nil)

																																			tmp19684 := PrimCons(sym_h, tmp19683)

																																			tmp19685 := PrimCons(tmp19679, tmp19684)

																																			tmp19686 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																			tmp19687 := PrimCons(tmp19685, tmp19686)

																																			tmp19688 := PrimCons(tmp19678, tmp19687)

																																			__e.TailApply(PrimNS2Value(symbind), V2790, tmp19688, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp19689 := PrimTail(V2574)

																																		__e.TailApply(tmp19672, tmp19689)
																																		return

																																	} else {
																																		tmp19671 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2596)

																																		if True == tmp19671 {
																																			tmp19649 := Call(__e, PrimNS2Value(symshen_4bindv), V2596, Nil, V2791)

																																			_ = tmp19649

																																			tmp19650 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp19651 := Call(__e, PrimNS2Value(symshen_4unbindv), V2596, V2791)

																																				_ = tmp19651

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp19652 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp19653 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp19653

																																				tmp19654 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp19655 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp19656 := PrimCons(tmp19655, Nil)

																																				tmp19657 := PrimCons(sym_h, tmp19656)

																																				tmp19658 := PrimCons(tmp19654, tmp19657)

																																				tmp19659 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp19660 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp19661 := PrimCons(tmp19660, Nil)

																																				tmp19662 := PrimCons(symvector, tmp19661)

																																				tmp19663 := PrimCons(tmp19662, Nil)

																																				tmp19664 := PrimCons(sym_h, tmp19663)

																																				tmp19665 := PrimCons(tmp19659, tmp19664)

																																				tmp19666 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp19667 := PrimCons(tmp19665, tmp19666)

																																				tmp19668 := PrimCons(tmp19658, tmp19667)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp19668, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp19669 := PrimTail(V2574)

																																			tmp19670 := Call(__e, tmp19652, tmp19669)

																																			__e.TailApply(tmp19650, tmp19670)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp19691 := PrimTail(V2583)

																																tmp19692 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp19691, V2791)

																																tmp19693 := Call(__e, tmp19646, tmp19692)

																																__e.TailApply(tmp19644, tmp19693)
																																return

																															}, 1)

																															tmp19694 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																															__e.TailApply(tmp19640, tmp19694)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp20044 := PrimHead(V2583)

																												tmp20045 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20044, V2791)

																												__e.TailApply(tmp19637, tmp20045)
																												return

																											} else {
																												__e.Return(False)
																												return
																											}

																										}, 1)

																										tmp20047 := PrimTail(V2581)

																										tmp20048 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20047, V2791)

																										__e.TailApply(tmp19635, tmp20048)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp20050 := PrimHead(V2581)

																								tmp20051 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20050, V2791)

																								__e.TailApply(tmp19633, tmp20051)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp20053 := PrimTail(V2575)

																						tmp20054 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20053, V2791)

																						__e.TailApply(tmp19631, tmp20054)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp20056 := PrimTail(V2579)

																				tmp20057 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20056, V2791)

																				__e.TailApply(tmp19629, tmp20057)
																				return

																			}, 1)

																			tmp20058 := PrimHead(V2579)

																			__e.TailApply(tmp19628, tmp20058)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp20060 := PrimTail(V2578)

																	tmp20061 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20060, V2791)

																	__e.TailApply(tmp19626, tmp20061)
																	return

																}, 1)

																tmp20062 := PrimHead(V2578)

																__e.TailApply(tmp19625, tmp20062)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp20064 := PrimTail(V2576)

														tmp20065 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20064, V2791)

														__e.TailApply(tmp19623, tmp20065)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp20067 := PrimHead(V2576)

												tmp20068 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20067, V2791)

												__e.TailApply(tmp19621, tmp20068)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp20070 := PrimHead(V2575)

										tmp20071 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20070, V2791)

										__e.TailApply(tmp19619, tmp20071)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp20073 := PrimHead(V2574)

								tmp20074 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20073, V2791)

								__e.TailApply(tmp19617, tmp20074)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp20076 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2789, V2791)

						tmp20077 := Call(__e, tmp19615, tmp20076)

						__e.TailApply(tmp19451, tmp20077)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp20079 := MakeNative(func(__e *ControlFlow) {
					V2549 := __e.Get(1)
					_ = V2549
					tmp20574 := PrimIsPair(V2549)

					if True == tmp20574 {
						tmp20081 := MakeNative(func(__e *ControlFlow) {
							V2550 := __e.Get(1)
							_ = V2550
							tmp20571 := PrimIsPair(V2550)

							if True == tmp20571 {
								tmp20083 := MakeNative(func(__e *ControlFlow) {
									V2551 := __e.Get(1)
									_ = V2551
									tmp20568 := PrimIsPair(V2551)

									if True == tmp20568 {
										tmp20085 := MakeNative(func(__e *ControlFlow) {
											V2552 := __e.Get(1)
											_ = V2552
											tmp20565 := PrimEqual(sym_8p, V2552)

											if True == tmp20565 {
												tmp20087 := MakeNative(func(__e *ControlFlow) {
													V2553 := __e.Get(1)
													_ = V2553
													tmp20562 := PrimIsPair(V2553)

													if True == tmp20562 {
														tmp20089 := MakeNative(func(__e *ControlFlow) {
															X := __e.Get(1)
															_ = X
															tmp20090 := MakeNative(func(__e *ControlFlow) {
																V2554 := __e.Get(1)
																_ = V2554
																tmp20558 := PrimIsPair(V2554)

																if True == tmp20558 {
																	tmp20092 := MakeNative(func(__e *ControlFlow) {
																		Y := __e.Get(1)
																		_ = Y
																		tmp20093 := MakeNative(func(__e *ControlFlow) {
																			V2555 := __e.Get(1)
																			_ = V2555
																			tmp20554 := PrimEqual(Nil, V2555)

																			if True == tmp20554 {
																				tmp20095 := MakeNative(func(__e *ControlFlow) {
																					V2556 := __e.Get(1)
																					_ = V2556
																					tmp20551 := PrimIsPair(V2556)

																					if True == tmp20551 {
																						tmp20097 := MakeNative(func(__e *ControlFlow) {
																							V2557 := __e.Get(1)
																							_ = V2557
																							tmp20548 := PrimEqual(sym_h, V2557)

																							if True == tmp20548 {
																								tmp20099 := MakeNative(func(__e *ControlFlow) {
																									V2558 := __e.Get(1)
																									_ = V2558
																									tmp20545 := PrimIsPair(V2558)

																									if True == tmp20545 {
																										tmp20101 := MakeNative(func(__e *ControlFlow) {
																											V2559 := __e.Get(1)
																											_ = V2559
																											tmp20542 := PrimIsPair(V2559)

																											if True == tmp20542 {
																												tmp20159 := MakeNative(func(__e *ControlFlow) {
																													A := __e.Get(1)
																													_ = A
																													tmp20160 := MakeNative(func(__e *ControlFlow) {
																														V2560 := __e.Get(1)
																														_ = V2560
																														tmp20538 := PrimIsPair(V2560)

																														if True == tmp20538 {
																															tmp20215 := MakeNative(func(__e *ControlFlow) {
																																V2561 := __e.Get(1)
																																_ = V2561
																																tmp20535 := PrimEqual(sym_d, V2561)

																																if True == tmp20535 {
																																	tmp20379 := MakeNative(func(__e *ControlFlow) {
																																		V2562 := __e.Get(1)
																																		_ = V2562
																																		tmp20532 := PrimIsPair(V2562)

																																		if True == tmp20532 {
																																			tmp20433 := MakeNative(func(__e *ControlFlow) {
																																				B := __e.Get(1)
																																				_ = B
																																				tmp20434 := MakeNative(func(__e *ControlFlow) {
																																					V2563 := __e.Get(1)
																																					_ = V2563
																																					tmp20528 := PrimEqual(Nil, V2563)

																																					if True == tmp20528 {
																																						tmp20485 := MakeNative(func(__e *ControlFlow) {
																																							V2564 := __e.Get(1)
																																							_ = V2564
																																							tmp20525 := PrimEqual(Nil, V2564)

																																							if True == tmp20525 {
																																								tmp20509 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp20510 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp20510

																																									tmp20511 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp20512 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp20513 := PrimCons(tmp20512, Nil)

																																									tmp20514 := PrimCons(sym_h, tmp20513)

																																									tmp20515 := PrimCons(tmp20511, tmp20514)

																																									tmp20516 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp20517 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																									tmp20518 := PrimCons(tmp20517, Nil)

																																									tmp20519 := PrimCons(sym_h, tmp20518)

																																									tmp20520 := PrimCons(tmp20516, tmp20519)

																																									tmp20521 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp20522 := PrimCons(tmp20520, tmp20521)

																																									tmp20523 := PrimCons(tmp20515, tmp20522)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp20523, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp20524 := PrimTail(V2549)

																																								__e.TailApply(tmp20509, tmp20524)
																																								return

																																							} else {
																																								tmp20508 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2564)

																																								if True == tmp20508 {
																																									tmp20488 := Call(__e, PrimNS2Value(symshen_4bindv), V2564, Nil, V2791)

																																									_ = tmp20488

																																									tmp20489 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp20490 := Call(__e, PrimNS2Value(symshen_4unbindv), V2564, V2791)

																																										_ = tmp20490

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp20491 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp20492 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp20492

																																										tmp20493 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp20494 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp20495 := PrimCons(tmp20494, Nil)

																																										tmp20496 := PrimCons(sym_h, tmp20495)

																																										tmp20497 := PrimCons(tmp20493, tmp20496)

																																										tmp20498 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp20499 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																										tmp20500 := PrimCons(tmp20499, Nil)

																																										tmp20501 := PrimCons(sym_h, tmp20500)

																																										tmp20502 := PrimCons(tmp20498, tmp20501)

																																										tmp20503 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp20504 := PrimCons(tmp20502, tmp20503)

																																										tmp20505 := PrimCons(tmp20497, tmp20504)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp20505, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp20506 := PrimTail(V2549)

																																									tmp20507 := Call(__e, tmp20491, tmp20506)

																																									__e.TailApply(tmp20489, tmp20507)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp20526 := PrimTail(V2558)

																																						tmp20527 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20526, V2791)

																																						__e.TailApply(tmp20485, tmp20527)
																																						return

																																					} else {
																																						tmp20484 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2563)

																																						if True == tmp20484 {
																																							tmp20437 := Call(__e, PrimNS2Value(symshen_4bindv), V2563, Nil, V2791)

																																							_ = tmp20437

																																							tmp20438 := MakeNative(func(__e *ControlFlow) {
																																								Result := __e.Get(1)
																																								_ = Result
																																								tmp20439 := Call(__e, PrimNS2Value(symshen_4unbindv), V2563, V2791)

																																								_ = tmp20439

																																								__e.Return(Result)
																																								return

																																							}, 1)

																																							tmp20440 := MakeNative(func(__e *ControlFlow) {
																																								V2565 := __e.Get(1)
																																								_ = V2565
																																								tmp20480 := PrimEqual(Nil, V2565)

																																								if True == tmp20480 {
																																									tmp20464 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp20465 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp20465

																																										tmp20466 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp20467 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp20468 := PrimCons(tmp20467, Nil)

																																										tmp20469 := PrimCons(sym_h, tmp20468)

																																										tmp20470 := PrimCons(tmp20466, tmp20469)

																																										tmp20471 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp20472 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																										tmp20473 := PrimCons(tmp20472, Nil)

																																										tmp20474 := PrimCons(sym_h, tmp20473)

																																										tmp20475 := PrimCons(tmp20471, tmp20474)

																																										tmp20476 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp20477 := PrimCons(tmp20475, tmp20476)

																																										tmp20478 := PrimCons(tmp20470, tmp20477)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp20478, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp20479 := PrimTail(V2549)

																																									__e.TailApply(tmp20464, tmp20479)
																																									return

																																								} else {
																																									tmp20463 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2565)

																																									if True == tmp20463 {
																																										tmp20443 := Call(__e, PrimNS2Value(symshen_4bindv), V2565, Nil, V2791)

																																										_ = tmp20443

																																										tmp20444 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp20445 := Call(__e, PrimNS2Value(symshen_4unbindv), V2565, V2791)

																																											_ = tmp20445

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp20446 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp20447 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp20447

																																											tmp20448 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																											tmp20449 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																											tmp20450 := PrimCons(tmp20449, Nil)

																																											tmp20451 := PrimCons(sym_h, tmp20450)

																																											tmp20452 := PrimCons(tmp20448, tmp20451)

																																											tmp20453 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																											tmp20454 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																											tmp20455 := PrimCons(tmp20454, Nil)

																																											tmp20456 := PrimCons(sym_h, tmp20455)

																																											tmp20457 := PrimCons(tmp20453, tmp20456)

																																											tmp20458 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																											tmp20459 := PrimCons(tmp20457, tmp20458)

																																											tmp20460 := PrimCons(tmp20452, tmp20459)

																																											__e.TailApply(PrimNS2Value(symbind), V2790, tmp20460, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp20461 := PrimTail(V2549)

																																										tmp20462 := Call(__e, tmp20446, tmp20461)

																																										__e.TailApply(tmp20444, tmp20462)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp20481 := PrimTail(V2558)

																																							tmp20482 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20481, V2791)

																																							tmp20483 := Call(__e, tmp20440, tmp20482)

																																							__e.TailApply(tmp20438, tmp20483)
																																							return

																																						} else {
																																							__e.Return(False)
																																							return
																																						}

																																					}

																																				}, 1)

																																				tmp20529 := PrimTail(V2562)

																																				tmp20530 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20529, V2791)

																																				__e.TailApply(tmp20434, tmp20530)
																																				return

																																			}, 1)

																																			tmp20531 := PrimHead(V2562)

																																			__e.TailApply(tmp20433, tmp20531)
																																			return

																																		} else {
																																			tmp20432 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2562)

																																			if True == tmp20432 {
																																				tmp20382 := MakeNative(func(__e *ControlFlow) {
																																					B := __e.Get(1)
																																					_ = B
																																					tmp20383 := PrimCons(B, Nil)

																																					tmp20384 := Call(__e, PrimNS2Value(symshen_4bindv), V2562, tmp20383, V2791)

																																					_ = tmp20384

																																					tmp20385 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp20386 := Call(__e, PrimNS2Value(symshen_4unbindv), V2562, V2791)

																																						_ = tmp20386

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp20387 := MakeNative(func(__e *ControlFlow) {
																																						V2566 := __e.Get(1)
																																						_ = V2566
																																						tmp20427 := PrimEqual(Nil, V2566)

																																						if True == tmp20427 {
																																							tmp20411 := MakeNative(func(__e *ControlFlow) {
																																								Hyp := __e.Get(1)
																																								_ = Hyp
																																								tmp20412 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																								_ = tmp20412

																																								tmp20413 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																								tmp20414 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																								tmp20415 := PrimCons(tmp20414, Nil)

																																								tmp20416 := PrimCons(sym_h, tmp20415)

																																								tmp20417 := PrimCons(tmp20413, tmp20416)

																																								tmp20418 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																								tmp20419 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																								tmp20420 := PrimCons(tmp20419, Nil)

																																								tmp20421 := PrimCons(sym_h, tmp20420)

																																								tmp20422 := PrimCons(tmp20418, tmp20421)

																																								tmp20423 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																								tmp20424 := PrimCons(tmp20422, tmp20423)

																																								tmp20425 := PrimCons(tmp20417, tmp20424)

																																								__e.TailApply(PrimNS2Value(symbind), V2790, tmp20425, V2791, V2792)
																																								return

																																							}, 1)

																																							tmp20426 := PrimTail(V2549)

																																							__e.TailApply(tmp20411, tmp20426)
																																							return

																																						} else {
																																							tmp20410 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2566)

																																							if True == tmp20410 {
																																								tmp20390 := Call(__e, PrimNS2Value(symshen_4bindv), V2566, Nil, V2791)

																																								_ = tmp20390

																																								tmp20391 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp20392 := Call(__e, PrimNS2Value(symshen_4unbindv), V2566, V2791)

																																									_ = tmp20392

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp20393 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp20394 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp20394

																																									tmp20395 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp20396 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp20397 := PrimCons(tmp20396, Nil)

																																									tmp20398 := PrimCons(sym_h, tmp20397)

																																									tmp20399 := PrimCons(tmp20395, tmp20398)

																																									tmp20400 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp20401 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																									tmp20402 := PrimCons(tmp20401, Nil)

																																									tmp20403 := PrimCons(sym_h, tmp20402)

																																									tmp20404 := PrimCons(tmp20400, tmp20403)

																																									tmp20405 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp20406 := PrimCons(tmp20404, tmp20405)

																																									tmp20407 := PrimCons(tmp20399, tmp20406)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp20407, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp20408 := PrimTail(V2549)

																																								tmp20409 := Call(__e, tmp20393, tmp20408)

																																								__e.TailApply(tmp20391, tmp20409)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp20428 := PrimTail(V2558)

																																					tmp20429 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20428, V2791)

																																					tmp20430 := Call(__e, tmp20387, tmp20429)

																																					__e.TailApply(tmp20385, tmp20430)
																																					return

																																				}, 1)

																																				tmp20431 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																				__e.TailApply(tmp20382, tmp20431)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp20533 := PrimTail(V2560)

																																	tmp20534 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20533, V2791)

																																	__e.TailApply(tmp20379, tmp20534)
																																	return

																																} else {
																																	tmp20378 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2561)

																																	if True == tmp20378 {
																																		tmp20218 := Call(__e, PrimNS2Value(symshen_4bindv), V2561, sym_d, V2791)

																																		_ = tmp20218

																																		tmp20219 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp20220 := Call(__e, PrimNS2Value(symshen_4unbindv), V2561, V2791)

																																			_ = tmp20220

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp20221 := MakeNative(func(__e *ControlFlow) {
																																			V2567 := __e.Get(1)
																																			_ = V2567
																																			tmp20374 := PrimIsPair(V2567)

																																			if True == tmp20374 {
																																				tmp20275 := MakeNative(func(__e *ControlFlow) {
																																					B := __e.Get(1)
																																					_ = B
																																					tmp20276 := MakeNative(func(__e *ControlFlow) {
																																						V2568 := __e.Get(1)
																																						_ = V2568
																																						tmp20370 := PrimEqual(Nil, V2568)

																																						if True == tmp20370 {
																																							tmp20327 := MakeNative(func(__e *ControlFlow) {
																																								V2569 := __e.Get(1)
																																								_ = V2569
																																								tmp20367 := PrimEqual(Nil, V2569)

																																								if True == tmp20367 {
																																									tmp20351 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp20352 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp20352

																																										tmp20353 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp20354 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp20355 := PrimCons(tmp20354, Nil)

																																										tmp20356 := PrimCons(sym_h, tmp20355)

																																										tmp20357 := PrimCons(tmp20353, tmp20356)

																																										tmp20358 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp20359 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																										tmp20360 := PrimCons(tmp20359, Nil)

																																										tmp20361 := PrimCons(sym_h, tmp20360)

																																										tmp20362 := PrimCons(tmp20358, tmp20361)

																																										tmp20363 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp20364 := PrimCons(tmp20362, tmp20363)

																																										tmp20365 := PrimCons(tmp20357, tmp20364)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp20365, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp20366 := PrimTail(V2549)

																																									__e.TailApply(tmp20351, tmp20366)
																																									return

																																								} else {
																																									tmp20350 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2569)

																																									if True == tmp20350 {
																																										tmp20330 := Call(__e, PrimNS2Value(symshen_4bindv), V2569, Nil, V2791)

																																										_ = tmp20330

																																										tmp20331 := MakeNative(func(__e *ControlFlow) {
																																											Result := __e.Get(1)
																																											_ = Result
																																											tmp20332 := Call(__e, PrimNS2Value(symshen_4unbindv), V2569, V2791)

																																											_ = tmp20332

																																											__e.Return(Result)
																																											return

																																										}, 1)

																																										tmp20333 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp20334 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp20334

																																											tmp20335 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																											tmp20336 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																											tmp20337 := PrimCons(tmp20336, Nil)

																																											tmp20338 := PrimCons(sym_h, tmp20337)

																																											tmp20339 := PrimCons(tmp20335, tmp20338)

																																											tmp20340 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																											tmp20341 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																											tmp20342 := PrimCons(tmp20341, Nil)

																																											tmp20343 := PrimCons(sym_h, tmp20342)

																																											tmp20344 := PrimCons(tmp20340, tmp20343)

																																											tmp20345 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																											tmp20346 := PrimCons(tmp20344, tmp20345)

																																											tmp20347 := PrimCons(tmp20339, tmp20346)

																																											__e.TailApply(PrimNS2Value(symbind), V2790, tmp20347, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp20348 := PrimTail(V2549)

																																										tmp20349 := Call(__e, tmp20333, tmp20348)

																																										__e.TailApply(tmp20331, tmp20349)
																																										return

																																									} else {
																																										__e.Return(False)
																																										return
																																									}

																																								}

																																							}, 1)

																																							tmp20368 := PrimTail(V2558)

																																							tmp20369 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20368, V2791)

																																							__e.TailApply(tmp20327, tmp20369)
																																							return

																																						} else {
																																							tmp20326 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2568)

																																							if True == tmp20326 {
																																								tmp20279 := Call(__e, PrimNS2Value(symshen_4bindv), V2568, Nil, V2791)

																																								_ = tmp20279

																																								tmp20280 := MakeNative(func(__e *ControlFlow) {
																																									Result := __e.Get(1)
																																									_ = Result
																																									tmp20281 := Call(__e, PrimNS2Value(symshen_4unbindv), V2568, V2791)

																																									_ = tmp20281

																																									__e.Return(Result)
																																									return

																																								}, 1)

																																								tmp20282 := MakeNative(func(__e *ControlFlow) {
																																									V2570 := __e.Get(1)
																																									_ = V2570
																																									tmp20322 := PrimEqual(Nil, V2570)

																																									if True == tmp20322 {
																																										tmp20306 := MakeNative(func(__e *ControlFlow) {
																																											Hyp := __e.Get(1)
																																											_ = Hyp
																																											tmp20307 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																											_ = tmp20307

																																											tmp20308 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																											tmp20309 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																											tmp20310 := PrimCons(tmp20309, Nil)

																																											tmp20311 := PrimCons(sym_h, tmp20310)

																																											tmp20312 := PrimCons(tmp20308, tmp20311)

																																											tmp20313 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																											tmp20314 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																											tmp20315 := PrimCons(tmp20314, Nil)

																																											tmp20316 := PrimCons(sym_h, tmp20315)

																																											tmp20317 := PrimCons(tmp20313, tmp20316)

																																											tmp20318 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																											tmp20319 := PrimCons(tmp20317, tmp20318)

																																											tmp20320 := PrimCons(tmp20312, tmp20319)

																																											__e.TailApply(PrimNS2Value(symbind), V2790, tmp20320, V2791, V2792)
																																											return

																																										}, 1)

																																										tmp20321 := PrimTail(V2549)

																																										__e.TailApply(tmp20306, tmp20321)
																																										return

																																									} else {
																																										tmp20305 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2570)

																																										if True == tmp20305 {
																																											tmp20285 := Call(__e, PrimNS2Value(symshen_4bindv), V2570, Nil, V2791)

																																											_ = tmp20285

																																											tmp20286 := MakeNative(func(__e *ControlFlow) {
																																												Result := __e.Get(1)
																																												_ = Result
																																												tmp20287 := Call(__e, PrimNS2Value(symshen_4unbindv), V2570, V2791)

																																												_ = tmp20287

																																												__e.Return(Result)
																																												return

																																											}, 1)

																																											tmp20288 := MakeNative(func(__e *ControlFlow) {
																																												Hyp := __e.Get(1)
																																												_ = Hyp
																																												tmp20289 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																												_ = tmp20289

																																												tmp20290 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																												tmp20291 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																												tmp20292 := PrimCons(tmp20291, Nil)

																																												tmp20293 := PrimCons(sym_h, tmp20292)

																																												tmp20294 := PrimCons(tmp20290, tmp20293)

																																												tmp20295 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																												tmp20296 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																												tmp20297 := PrimCons(tmp20296, Nil)

																																												tmp20298 := PrimCons(sym_h, tmp20297)

																																												tmp20299 := PrimCons(tmp20295, tmp20298)

																																												tmp20300 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																												tmp20301 := PrimCons(tmp20299, tmp20300)

																																												tmp20302 := PrimCons(tmp20294, tmp20301)

																																												__e.TailApply(PrimNS2Value(symbind), V2790, tmp20302, V2791, V2792)
																																												return

																																											}, 1)

																																											tmp20303 := PrimTail(V2549)

																																											tmp20304 := Call(__e, tmp20288, tmp20303)

																																											__e.TailApply(tmp20286, tmp20304)
																																											return

																																										} else {
																																											__e.Return(False)
																																											return
																																										}

																																									}

																																								}, 1)

																																								tmp20323 := PrimTail(V2558)

																																								tmp20324 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20323, V2791)

																																								tmp20325 := Call(__e, tmp20282, tmp20324)

																																								__e.TailApply(tmp20280, tmp20325)
																																								return

																																							} else {
																																								__e.Return(False)
																																								return
																																							}

																																						}

																																					}, 1)

																																					tmp20371 := PrimTail(V2567)

																																					tmp20372 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20371, V2791)

																																					__e.TailApply(tmp20276, tmp20372)
																																					return

																																				}, 1)

																																				tmp20373 := PrimHead(V2567)

																																				__e.TailApply(tmp20275, tmp20373)
																																				return

																																			} else {
																																				tmp20274 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2567)

																																				if True == tmp20274 {
																																					tmp20224 := MakeNative(func(__e *ControlFlow) {
																																						B := __e.Get(1)
																																						_ = B
																																						tmp20225 := PrimCons(B, Nil)

																																						tmp20226 := Call(__e, PrimNS2Value(symshen_4bindv), V2567, tmp20225, V2791)

																																						_ = tmp20226

																																						tmp20227 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp20228 := Call(__e, PrimNS2Value(symshen_4unbindv), V2567, V2791)

																																							_ = tmp20228

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp20229 := MakeNative(func(__e *ControlFlow) {
																																							V2571 := __e.Get(1)
																																							_ = V2571
																																							tmp20269 := PrimEqual(Nil, V2571)

																																							if True == tmp20269 {
																																								tmp20253 := MakeNative(func(__e *ControlFlow) {
																																									Hyp := __e.Get(1)
																																									_ = Hyp
																																									tmp20254 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																									_ = tmp20254

																																									tmp20255 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																									tmp20256 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																									tmp20257 := PrimCons(tmp20256, Nil)

																																									tmp20258 := PrimCons(sym_h, tmp20257)

																																									tmp20259 := PrimCons(tmp20255, tmp20258)

																																									tmp20260 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																									tmp20261 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																									tmp20262 := PrimCons(tmp20261, Nil)

																																									tmp20263 := PrimCons(sym_h, tmp20262)

																																									tmp20264 := PrimCons(tmp20260, tmp20263)

																																									tmp20265 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																									tmp20266 := PrimCons(tmp20264, tmp20265)

																																									tmp20267 := PrimCons(tmp20259, tmp20266)

																																									__e.TailApply(PrimNS2Value(symbind), V2790, tmp20267, V2791, V2792)
																																									return

																																								}, 1)

																																								tmp20268 := PrimTail(V2549)

																																								__e.TailApply(tmp20253, tmp20268)
																																								return

																																							} else {
																																								tmp20252 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2571)

																																								if True == tmp20252 {
																																									tmp20232 := Call(__e, PrimNS2Value(symshen_4bindv), V2571, Nil, V2791)

																																									_ = tmp20232

																																									tmp20233 := MakeNative(func(__e *ControlFlow) {
																																										Result := __e.Get(1)
																																										_ = Result
																																										tmp20234 := Call(__e, PrimNS2Value(symshen_4unbindv), V2571, V2791)

																																										_ = tmp20234

																																										__e.Return(Result)
																																										return

																																									}, 1)

																																									tmp20235 := MakeNative(func(__e *ControlFlow) {
																																										Hyp := __e.Get(1)
																																										_ = Hyp
																																										tmp20236 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																										_ = tmp20236

																																										tmp20237 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																										tmp20238 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																										tmp20239 := PrimCons(tmp20238, Nil)

																																										tmp20240 := PrimCons(sym_h, tmp20239)

																																										tmp20241 := PrimCons(tmp20237, tmp20240)

																																										tmp20242 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																										tmp20243 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																										tmp20244 := PrimCons(tmp20243, Nil)

																																										tmp20245 := PrimCons(sym_h, tmp20244)

																																										tmp20246 := PrimCons(tmp20242, tmp20245)

																																										tmp20247 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																										tmp20248 := PrimCons(tmp20246, tmp20247)

																																										tmp20249 := PrimCons(tmp20241, tmp20248)

																																										__e.TailApply(PrimNS2Value(symbind), V2790, tmp20249, V2791, V2792)
																																										return

																																									}, 1)

																																									tmp20250 := PrimTail(V2549)

																																									tmp20251 := Call(__e, tmp20235, tmp20250)

																																									__e.TailApply(tmp20233, tmp20251)
																																									return

																																								} else {
																																									__e.Return(False)
																																									return
																																								}

																																							}

																																						}, 1)

																																						tmp20270 := PrimTail(V2558)

																																						tmp20271 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20270, V2791)

																																						tmp20272 := Call(__e, tmp20229, tmp20271)

																																						__e.TailApply(tmp20227, tmp20272)
																																						return

																																					}, 1)

																																					tmp20273 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																					__e.TailApply(tmp20224, tmp20273)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp20375 := PrimTail(V2560)

																																		tmp20376 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20375, V2791)

																																		tmp20377 := Call(__e, tmp20221, tmp20376)

																																		__e.TailApply(tmp20219, tmp20377)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp20536 := PrimHead(V2560)

																															tmp20537 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20536, V2791)

																															__e.TailApply(tmp20215, tmp20537)
																															return

																														} else {
																															tmp20214 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2560)

																															if True == tmp20214 {
																																tmp20163 := MakeNative(func(__e *ControlFlow) {
																																	B := __e.Get(1)
																																	_ = B
																																	tmp20164 := PrimCons(B, Nil)

																																	tmp20165 := PrimCons(sym_d, tmp20164)

																																	tmp20166 := Call(__e, PrimNS2Value(symshen_4bindv), V2560, tmp20165, V2791)

																																	_ = tmp20166

																																	tmp20167 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp20168 := Call(__e, PrimNS2Value(symshen_4unbindv), V2560, V2791)

																																		_ = tmp20168

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp20169 := MakeNative(func(__e *ControlFlow) {
																																		V2572 := __e.Get(1)
																																		_ = V2572
																																		tmp20209 := PrimEqual(Nil, V2572)

																																		if True == tmp20209 {
																																			tmp20193 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp20194 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp20194

																																				tmp20195 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp20196 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20197 := PrimCons(tmp20196, Nil)

																																				tmp20198 := PrimCons(sym_h, tmp20197)

																																				tmp20199 := PrimCons(tmp20195, tmp20198)

																																				tmp20200 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp20201 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																				tmp20202 := PrimCons(tmp20201, Nil)

																																				tmp20203 := PrimCons(sym_h, tmp20202)

																																				tmp20204 := PrimCons(tmp20200, tmp20203)

																																				tmp20205 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp20206 := PrimCons(tmp20204, tmp20205)

																																				tmp20207 := PrimCons(tmp20199, tmp20206)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp20207, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp20208 := PrimTail(V2549)

																																			__e.TailApply(tmp20193, tmp20208)
																																			return

																																		} else {
																																			tmp20192 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2572)

																																			if True == tmp20192 {
																																				tmp20172 := Call(__e, PrimNS2Value(symshen_4bindv), V2572, Nil, V2791)

																																				_ = tmp20172

																																				tmp20173 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp20174 := Call(__e, PrimNS2Value(symshen_4unbindv), V2572, V2791)

																																					_ = tmp20174

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp20175 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp20176 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp20176

																																					tmp20177 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp20178 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20179 := PrimCons(tmp20178, Nil)

																																					tmp20180 := PrimCons(sym_h, tmp20179)

																																					tmp20181 := PrimCons(tmp20177, tmp20180)

																																					tmp20182 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp20183 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																					tmp20184 := PrimCons(tmp20183, Nil)

																																					tmp20185 := PrimCons(sym_h, tmp20184)

																																					tmp20186 := PrimCons(tmp20182, tmp20185)

																																					tmp20187 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp20188 := PrimCons(tmp20186, tmp20187)

																																					tmp20189 := PrimCons(tmp20181, tmp20188)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp20189, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp20190 := PrimTail(V2549)

																																				tmp20191 := Call(__e, tmp20175, tmp20190)

																																				__e.TailApply(tmp20173, tmp20191)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp20210 := PrimTail(V2558)

																																	tmp20211 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20210, V2791)

																																	tmp20212 := Call(__e, tmp20169, tmp20211)

																																	__e.TailApply(tmp20167, tmp20212)
																																	return

																																}, 1)

																																tmp20213 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																__e.TailApply(tmp20163, tmp20213)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp20539 := PrimTail(V2559)

																													tmp20540 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20539, V2791)

																													__e.TailApply(tmp20160, tmp20540)
																													return

																												}, 1)

																												tmp20541 := PrimHead(V2559)

																												__e.TailApply(tmp20159, tmp20541)
																												return

																											} else {
																												tmp20158 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2559)

																												if True == tmp20158 {
																													tmp20104 := MakeNative(func(__e *ControlFlow) {
																														A := __e.Get(1)
																														_ = A
																														tmp20105 := MakeNative(func(__e *ControlFlow) {
																															B := __e.Get(1)
																															_ = B
																															tmp20106 := PrimCons(B, Nil)

																															tmp20107 := PrimCons(sym_d, tmp20106)

																															tmp20108 := PrimCons(A, tmp20107)

																															tmp20109 := Call(__e, PrimNS2Value(symshen_4bindv), V2559, tmp20108, V2791)

																															_ = tmp20109

																															tmp20110 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp20111 := Call(__e, PrimNS2Value(symshen_4unbindv), V2559, V2791)

																																_ = tmp20111

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp20112 := MakeNative(func(__e *ControlFlow) {
																																V2573 := __e.Get(1)
																																_ = V2573
																																tmp20152 := PrimEqual(Nil, V2573)

																																if True == tmp20152 {
																																	tmp20136 := MakeNative(func(__e *ControlFlow) {
																																		Hyp := __e.Get(1)
																																		_ = Hyp
																																		tmp20137 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																		_ = tmp20137

																																		tmp20138 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																		tmp20139 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																		tmp20140 := PrimCons(tmp20139, Nil)

																																		tmp20141 := PrimCons(sym_h, tmp20140)

																																		tmp20142 := PrimCons(tmp20138, tmp20141)

																																		tmp20143 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																		tmp20144 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																		tmp20145 := PrimCons(tmp20144, Nil)

																																		tmp20146 := PrimCons(sym_h, tmp20145)

																																		tmp20147 := PrimCons(tmp20143, tmp20146)

																																		tmp20148 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																		tmp20149 := PrimCons(tmp20147, tmp20148)

																																		tmp20150 := PrimCons(tmp20142, tmp20149)

																																		__e.TailApply(PrimNS2Value(symbind), V2790, tmp20150, V2791, V2792)
																																		return

																																	}, 1)

																																	tmp20151 := PrimTail(V2549)

																																	__e.TailApply(tmp20136, tmp20151)
																																	return

																																} else {
																																	tmp20135 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2573)

																																	if True == tmp20135 {
																																		tmp20115 := Call(__e, PrimNS2Value(symshen_4bindv), V2573, Nil, V2791)

																																		_ = tmp20115

																																		tmp20116 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp20117 := Call(__e, PrimNS2Value(symshen_4unbindv), V2573, V2791)

																																			_ = tmp20117

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp20118 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp20119 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp20119

																																			tmp20120 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																			tmp20121 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																			tmp20122 := PrimCons(tmp20121, Nil)

																																			tmp20123 := PrimCons(sym_h, tmp20122)

																																			tmp20124 := PrimCons(tmp20120, tmp20123)

																																			tmp20125 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																			tmp20126 := Call(__e, PrimNS2Value(symshen_4lazyderef), B, V2791)

																																			tmp20127 := PrimCons(tmp20126, Nil)

																																			tmp20128 := PrimCons(sym_h, tmp20127)

																																			tmp20129 := PrimCons(tmp20125, tmp20128)

																																			tmp20130 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																			tmp20131 := PrimCons(tmp20129, tmp20130)

																																			tmp20132 := PrimCons(tmp20124, tmp20131)

																																			__e.TailApply(PrimNS2Value(symbind), V2790, tmp20132, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp20133 := PrimTail(V2549)

																																		tmp20134 := Call(__e, tmp20118, tmp20133)

																																		__e.TailApply(tmp20116, tmp20134)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp20153 := PrimTail(V2558)

																															tmp20154 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20153, V2791)

																															tmp20155 := Call(__e, tmp20112, tmp20154)

																															__e.TailApply(tmp20110, tmp20155)
																															return

																														}, 1)

																														tmp20156 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																														__e.TailApply(tmp20105, tmp20156)
																														return

																													}, 1)

																													tmp20157 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																													__e.TailApply(tmp20104, tmp20157)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}

																										}, 1)

																										tmp20543 := PrimHead(V2558)

																										tmp20544 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20543, V2791)

																										__e.TailApply(tmp20101, tmp20544)
																										return

																									} else {
																										__e.Return(False)
																										return
																									}

																								}, 1)

																								tmp20546 := PrimTail(V2556)

																								tmp20547 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20546, V2791)

																								__e.TailApply(tmp20099, tmp20547)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp20549 := PrimHead(V2556)

																						tmp20550 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20549, V2791)

																						__e.TailApply(tmp20097, tmp20550)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp20552 := PrimTail(V2550)

																				tmp20553 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20552, V2791)

																				__e.TailApply(tmp20095, tmp20553)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp20555 := PrimTail(V2554)

																		tmp20556 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20555, V2791)

																		__e.TailApply(tmp20093, tmp20556)
																		return

																	}, 1)

																	tmp20557 := PrimHead(V2554)

																	__e.TailApply(tmp20092, tmp20557)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp20559 := PrimTail(V2553)

															tmp20560 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20559, V2791)

															__e.TailApply(tmp20090, tmp20560)
															return

														}, 1)

														tmp20561 := PrimHead(V2553)

														__e.TailApply(tmp20089, tmp20561)
														return

													} else {
														__e.Return(False)
														return
													}

												}, 1)

												tmp20563 := PrimTail(V2551)

												tmp20564 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20563, V2791)

												__e.TailApply(tmp20087, tmp20564)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp20566 := PrimHead(V2551)

										tmp20567 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20566, V2791)

										__e.TailApply(tmp20085, tmp20567)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp20569 := PrimHead(V2550)

								tmp20570 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20569, V2791)

								__e.TailApply(tmp20083, tmp20570)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp20572 := PrimHead(V2549)

						tmp20573 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20572, V2791)

						__e.TailApply(tmp20081, tmp20573)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp20575 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2789, V2791)

				tmp20576 := Call(__e, tmp20079, tmp20575)

				__e.TailApply(tmp19449, tmp20576)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp20578 := MakeNative(func(__e *ControlFlow) {
			V2526 := __e.Get(1)
			_ = V2526
			tmp21038 := PrimIsPair(V2526)

			if True == tmp21038 {
				tmp20580 := MakeNative(func(__e *ControlFlow) {
					V2527 := __e.Get(1)
					_ = V2527
					tmp21035 := PrimIsPair(V2527)

					if True == tmp21035 {
						tmp20582 := MakeNative(func(__e *ControlFlow) {
							V2528 := __e.Get(1)
							_ = V2528
							tmp21032 := PrimIsPair(V2528)

							if True == tmp21032 {
								tmp20584 := MakeNative(func(__e *ControlFlow) {
									V2529 := __e.Get(1)
									_ = V2529
									tmp21029 := PrimEqual(symcons, V2529)

									if True == tmp21029 {
										tmp20586 := MakeNative(func(__e *ControlFlow) {
											V2530 := __e.Get(1)
											_ = V2530
											tmp21026 := PrimIsPair(V2530)

											if True == tmp21026 {
												tmp20588 := MakeNative(func(__e *ControlFlow) {
													X := __e.Get(1)
													_ = X
													tmp20589 := MakeNative(func(__e *ControlFlow) {
														V2531 := __e.Get(1)
														_ = V2531
														tmp21022 := PrimIsPair(V2531)

														if True == tmp21022 {
															tmp20591 := MakeNative(func(__e *ControlFlow) {
																Y := __e.Get(1)
																_ = Y
																tmp20592 := MakeNative(func(__e *ControlFlow) {
																	V2532 := __e.Get(1)
																	_ = V2532
																	tmp21018 := PrimEqual(Nil, V2532)

																	if True == tmp21018 {
																		tmp20594 := MakeNative(func(__e *ControlFlow) {
																			V2533 := __e.Get(1)
																			_ = V2533
																			tmp21015 := PrimIsPair(V2533)

																			if True == tmp21015 {
																				tmp20596 := MakeNative(func(__e *ControlFlow) {
																					V2534 := __e.Get(1)
																					_ = V2534
																					tmp21012 := PrimEqual(sym_h, V2534)

																					if True == tmp21012 {
																						tmp20598 := MakeNative(func(__e *ControlFlow) {
																							V2535 := __e.Get(1)
																							_ = V2535
																							tmp21009 := PrimIsPair(V2535)

																							if True == tmp21009 {
																								tmp20600 := MakeNative(func(__e *ControlFlow) {
																									V2536 := __e.Get(1)
																									_ = V2536
																									tmp21006 := PrimIsPair(V2536)

																									if True == tmp21006 {
																										tmp20659 := MakeNative(func(__e *ControlFlow) {
																											V2537 := __e.Get(1)
																											_ = V2537
																											tmp21003 := PrimEqual(symlist, V2537)

																											if True == tmp21003 {
																												tmp20835 := MakeNative(func(__e *ControlFlow) {
																													V2538 := __e.Get(1)
																													_ = V2538
																													tmp21000 := PrimIsPair(V2538)

																													if True == tmp21000 {
																														tmp20893 := MakeNative(func(__e *ControlFlow) {
																															A := __e.Get(1)
																															_ = A
																															tmp20894 := MakeNative(func(__e *ControlFlow) {
																																V2539 := __e.Get(1)
																																_ = V2539
																																tmp20996 := PrimEqual(Nil, V2539)

																																if True == tmp20996 {
																																	tmp20949 := MakeNative(func(__e *ControlFlow) {
																																		V2540 := __e.Get(1)
																																		_ = V2540
																																		tmp20993 := PrimEqual(Nil, V2540)

																																		if True == tmp20993 {
																																			tmp20975 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp20976 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp20976

																																				tmp20977 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp20978 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20979 := PrimCons(tmp20978, Nil)

																																				tmp20980 := PrimCons(sym_h, tmp20979)

																																				tmp20981 := PrimCons(tmp20977, tmp20980)

																																				tmp20982 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp20983 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20984 := PrimCons(tmp20983, Nil)

																																				tmp20985 := PrimCons(symlist, tmp20984)

																																				tmp20986 := PrimCons(tmp20985, Nil)

																																				tmp20987 := PrimCons(sym_h, tmp20986)

																																				tmp20988 := PrimCons(tmp20982, tmp20987)

																																				tmp20989 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp20990 := PrimCons(tmp20988, tmp20989)

																																				tmp20991 := PrimCons(tmp20981, tmp20990)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp20991, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp20992 := PrimTail(V2526)

																																			__e.TailApply(tmp20975, tmp20992)
																																			return

																																		} else {
																																			tmp20974 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2540)

																																			if True == tmp20974 {
																																				tmp20952 := Call(__e, PrimNS2Value(symshen_4bindv), V2540, Nil, V2791)

																																				_ = tmp20952

																																				tmp20953 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp20954 := Call(__e, PrimNS2Value(symshen_4unbindv), V2540, V2791)

																																					_ = tmp20954

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp20955 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp20956 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp20956

																																					tmp20957 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp20958 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20959 := PrimCons(tmp20958, Nil)

																																					tmp20960 := PrimCons(sym_h, tmp20959)

																																					tmp20961 := PrimCons(tmp20957, tmp20960)

																																					tmp20962 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp20963 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20964 := PrimCons(tmp20963, Nil)

																																					tmp20965 := PrimCons(symlist, tmp20964)

																																					tmp20966 := PrimCons(tmp20965, Nil)

																																					tmp20967 := PrimCons(sym_h, tmp20966)

																																					tmp20968 := PrimCons(tmp20962, tmp20967)

																																					tmp20969 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp20970 := PrimCons(tmp20968, tmp20969)

																																					tmp20971 := PrimCons(tmp20961, tmp20970)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp20971, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp20972 := PrimTail(V2526)

																																				tmp20973 := Call(__e, tmp20955, tmp20972)

																																				__e.TailApply(tmp20953, tmp20973)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp20994 := PrimTail(V2535)

																																	tmp20995 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20994, V2791)

																																	__e.TailApply(tmp20949, tmp20995)
																																	return

																																} else {
																																	tmp20948 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2539)

																																	if True == tmp20948 {
																																		tmp20897 := Call(__e, PrimNS2Value(symshen_4bindv), V2539, Nil, V2791)

																																		_ = tmp20897

																																		tmp20898 := MakeNative(func(__e *ControlFlow) {
																																			Result := __e.Get(1)
																																			_ = Result
																																			tmp20899 := Call(__e, PrimNS2Value(symshen_4unbindv), V2539, V2791)

																																			_ = tmp20899

																																			__e.Return(Result)
																																			return

																																		}, 1)

																																		tmp20900 := MakeNative(func(__e *ControlFlow) {
																																			V2541 := __e.Get(1)
																																			_ = V2541
																																			tmp20944 := PrimEqual(Nil, V2541)

																																			if True == tmp20944 {
																																				tmp20926 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp20927 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp20927

																																					tmp20928 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp20929 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20930 := PrimCons(tmp20929, Nil)

																																					tmp20931 := PrimCons(sym_h, tmp20930)

																																					tmp20932 := PrimCons(tmp20928, tmp20931)

																																					tmp20933 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp20934 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20935 := PrimCons(tmp20934, Nil)

																																					tmp20936 := PrimCons(symlist, tmp20935)

																																					tmp20937 := PrimCons(tmp20936, Nil)

																																					tmp20938 := PrimCons(sym_h, tmp20937)

																																					tmp20939 := PrimCons(tmp20933, tmp20938)

																																					tmp20940 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp20941 := PrimCons(tmp20939, tmp20940)

																																					tmp20942 := PrimCons(tmp20932, tmp20941)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp20942, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp20943 := PrimTail(V2526)

																																				__e.TailApply(tmp20926, tmp20943)
																																				return

																																			} else {
																																				tmp20925 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2541)

																																				if True == tmp20925 {
																																					tmp20903 := Call(__e, PrimNS2Value(symshen_4bindv), V2541, Nil, V2791)

																																					_ = tmp20903

																																					tmp20904 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp20905 := Call(__e, PrimNS2Value(symshen_4unbindv), V2541, V2791)

																																						_ = tmp20905

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp20906 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp20907 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																						_ = tmp20907

																																						tmp20908 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																						tmp20909 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20910 := PrimCons(tmp20909, Nil)

																																						tmp20911 := PrimCons(sym_h, tmp20910)

																																						tmp20912 := PrimCons(tmp20908, tmp20911)

																																						tmp20913 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																						tmp20914 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20915 := PrimCons(tmp20914, Nil)

																																						tmp20916 := PrimCons(symlist, tmp20915)

																																						tmp20917 := PrimCons(tmp20916, Nil)

																																						tmp20918 := PrimCons(sym_h, tmp20917)

																																						tmp20919 := PrimCons(tmp20913, tmp20918)

																																						tmp20920 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																						tmp20921 := PrimCons(tmp20919, tmp20920)

																																						tmp20922 := PrimCons(tmp20912, tmp20921)

																																						__e.TailApply(PrimNS2Value(symbind), V2790, tmp20922, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp20923 := PrimTail(V2526)

																																					tmp20924 := Call(__e, tmp20906, tmp20923)

																																					__e.TailApply(tmp20904, tmp20924)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp20945 := PrimTail(V2535)

																																		tmp20946 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20945, V2791)

																																		tmp20947 := Call(__e, tmp20900, tmp20946)

																																		__e.TailApply(tmp20898, tmp20947)
																																		return

																																	} else {
																																		__e.Return(False)
																																		return
																																	}

																																}

																															}, 1)

																															tmp20997 := PrimTail(V2538)

																															tmp20998 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20997, V2791)

																															__e.TailApply(tmp20894, tmp20998)
																															return

																														}, 1)

																														tmp20999 := PrimHead(V2538)

																														__e.TailApply(tmp20893, tmp20999)
																														return

																													} else {
																														tmp20892 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2538)

																														if True == tmp20892 {
																															tmp20838 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp20839 := PrimCons(A, Nil)

																																tmp20840 := Call(__e, PrimNS2Value(symshen_4bindv), V2538, tmp20839, V2791)

																																_ = tmp20840

																																tmp20841 := MakeNative(func(__e *ControlFlow) {
																																	Result := __e.Get(1)
																																	_ = Result
																																	tmp20842 := Call(__e, PrimNS2Value(symshen_4unbindv), V2538, V2791)

																																	_ = tmp20842

																																	__e.Return(Result)
																																	return

																																}, 1)

																																tmp20843 := MakeNative(func(__e *ControlFlow) {
																																	V2542 := __e.Get(1)
																																	_ = V2542
																																	tmp20887 := PrimEqual(Nil, V2542)

																																	if True == tmp20887 {
																																		tmp20869 := MakeNative(func(__e *ControlFlow) {
																																			Hyp := __e.Get(1)
																																			_ = Hyp
																																			tmp20870 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																			_ = tmp20870

																																			tmp20871 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																			tmp20872 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																			tmp20873 := PrimCons(tmp20872, Nil)

																																			tmp20874 := PrimCons(sym_h, tmp20873)

																																			tmp20875 := PrimCons(tmp20871, tmp20874)

																																			tmp20876 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																			tmp20877 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																			tmp20878 := PrimCons(tmp20877, Nil)

																																			tmp20879 := PrimCons(symlist, tmp20878)

																																			tmp20880 := PrimCons(tmp20879, Nil)

																																			tmp20881 := PrimCons(sym_h, tmp20880)

																																			tmp20882 := PrimCons(tmp20876, tmp20881)

																																			tmp20883 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																			tmp20884 := PrimCons(tmp20882, tmp20883)

																																			tmp20885 := PrimCons(tmp20875, tmp20884)

																																			__e.TailApply(PrimNS2Value(symbind), V2790, tmp20885, V2791, V2792)
																																			return

																																		}, 1)

																																		tmp20886 := PrimTail(V2526)

																																		__e.TailApply(tmp20869, tmp20886)
																																		return

																																	} else {
																																		tmp20868 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2542)

																																		if True == tmp20868 {
																																			tmp20846 := Call(__e, PrimNS2Value(symshen_4bindv), V2542, Nil, V2791)

																																			_ = tmp20846

																																			tmp20847 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp20848 := Call(__e, PrimNS2Value(symshen_4unbindv), V2542, V2791)

																																				_ = tmp20848

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp20849 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp20850 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp20850

																																				tmp20851 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp20852 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20853 := PrimCons(tmp20852, Nil)

																																				tmp20854 := PrimCons(sym_h, tmp20853)

																																				tmp20855 := PrimCons(tmp20851, tmp20854)

																																				tmp20856 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp20857 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20858 := PrimCons(tmp20857, Nil)

																																				tmp20859 := PrimCons(symlist, tmp20858)

																																				tmp20860 := PrimCons(tmp20859, Nil)

																																				tmp20861 := PrimCons(sym_h, tmp20860)

																																				tmp20862 := PrimCons(tmp20856, tmp20861)

																																				tmp20863 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp20864 := PrimCons(tmp20862, tmp20863)

																																				tmp20865 := PrimCons(tmp20855, tmp20864)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp20865, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp20866 := PrimTail(V2526)

																																			tmp20867 := Call(__e, tmp20849, tmp20866)

																																			__e.TailApply(tmp20847, tmp20867)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp20888 := PrimTail(V2535)

																																tmp20889 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20888, V2791)

																																tmp20890 := Call(__e, tmp20843, tmp20889)

																																__e.TailApply(tmp20841, tmp20890)
																																return

																															}, 1)

																															tmp20891 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																															__e.TailApply(tmp20838, tmp20891)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp21001 := PrimTail(V2536)

																												tmp21002 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21001, V2791)

																												__e.TailApply(tmp20835, tmp21002)
																												return

																											} else {
																												tmp20834 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2537)

																												if True == tmp20834 {
																													tmp20662 := Call(__e, PrimNS2Value(symshen_4bindv), V2537, symlist, V2791)

																													_ = tmp20662

																													tmp20663 := MakeNative(func(__e *ControlFlow) {
																														Result := __e.Get(1)
																														_ = Result
																														tmp20664 := Call(__e, PrimNS2Value(symshen_4unbindv), V2537, V2791)

																														_ = tmp20664

																														__e.Return(Result)
																														return

																													}, 1)

																													tmp20665 := MakeNative(func(__e *ControlFlow) {
																														V2543 := __e.Get(1)
																														_ = V2543
																														tmp20830 := PrimIsPair(V2543)

																														if True == tmp20830 {
																															tmp20723 := MakeNative(func(__e *ControlFlow) {
																																A := __e.Get(1)
																																_ = A
																																tmp20724 := MakeNative(func(__e *ControlFlow) {
																																	V2544 := __e.Get(1)
																																	_ = V2544
																																	tmp20826 := PrimEqual(Nil, V2544)

																																	if True == tmp20826 {
																																		tmp20779 := MakeNative(func(__e *ControlFlow) {
																																			V2545 := __e.Get(1)
																																			_ = V2545
																																			tmp20823 := PrimEqual(Nil, V2545)

																																			if True == tmp20823 {
																																				tmp20805 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp20806 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp20806

																																					tmp20807 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp20808 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20809 := PrimCons(tmp20808, Nil)

																																					tmp20810 := PrimCons(sym_h, tmp20809)

																																					tmp20811 := PrimCons(tmp20807, tmp20810)

																																					tmp20812 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp20813 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20814 := PrimCons(tmp20813, Nil)

																																					tmp20815 := PrimCons(symlist, tmp20814)

																																					tmp20816 := PrimCons(tmp20815, Nil)

																																					tmp20817 := PrimCons(sym_h, tmp20816)

																																					tmp20818 := PrimCons(tmp20812, tmp20817)

																																					tmp20819 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp20820 := PrimCons(tmp20818, tmp20819)

																																					tmp20821 := PrimCons(tmp20811, tmp20820)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp20821, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp20822 := PrimTail(V2526)

																																				__e.TailApply(tmp20805, tmp20822)
																																				return

																																			} else {
																																				tmp20804 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2545)

																																				if True == tmp20804 {
																																					tmp20782 := Call(__e, PrimNS2Value(symshen_4bindv), V2545, Nil, V2791)

																																					_ = tmp20782

																																					tmp20783 := MakeNative(func(__e *ControlFlow) {
																																						Result := __e.Get(1)
																																						_ = Result
																																						tmp20784 := Call(__e, PrimNS2Value(symshen_4unbindv), V2545, V2791)

																																						_ = tmp20784

																																						__e.Return(Result)
																																						return

																																					}, 1)

																																					tmp20785 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp20786 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																						_ = tmp20786

																																						tmp20787 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																						tmp20788 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20789 := PrimCons(tmp20788, Nil)

																																						tmp20790 := PrimCons(sym_h, tmp20789)

																																						tmp20791 := PrimCons(tmp20787, tmp20790)

																																						tmp20792 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																						tmp20793 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20794 := PrimCons(tmp20793, Nil)

																																						tmp20795 := PrimCons(symlist, tmp20794)

																																						tmp20796 := PrimCons(tmp20795, Nil)

																																						tmp20797 := PrimCons(sym_h, tmp20796)

																																						tmp20798 := PrimCons(tmp20792, tmp20797)

																																						tmp20799 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																						tmp20800 := PrimCons(tmp20798, tmp20799)

																																						tmp20801 := PrimCons(tmp20791, tmp20800)

																																						__e.TailApply(PrimNS2Value(symbind), V2790, tmp20801, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp20802 := PrimTail(V2526)

																																					tmp20803 := Call(__e, tmp20785, tmp20802)

																																					__e.TailApply(tmp20783, tmp20803)
																																					return

																																				} else {
																																					__e.Return(False)
																																					return
																																				}

																																			}

																																		}, 1)

																																		tmp20824 := PrimTail(V2535)

																																		tmp20825 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20824, V2791)

																																		__e.TailApply(tmp20779, tmp20825)
																																		return

																																	} else {
																																		tmp20778 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2544)

																																		if True == tmp20778 {
																																			tmp20727 := Call(__e, PrimNS2Value(symshen_4bindv), V2544, Nil, V2791)

																																			_ = tmp20727

																																			tmp20728 := MakeNative(func(__e *ControlFlow) {
																																				Result := __e.Get(1)
																																				_ = Result
																																				tmp20729 := Call(__e, PrimNS2Value(symshen_4unbindv), V2544, V2791)

																																				_ = tmp20729

																																				__e.Return(Result)
																																				return

																																			}, 1)

																																			tmp20730 := MakeNative(func(__e *ControlFlow) {
																																				V2546 := __e.Get(1)
																																				_ = V2546
																																				tmp20774 := PrimEqual(Nil, V2546)

																																				if True == tmp20774 {
																																					tmp20756 := MakeNative(func(__e *ControlFlow) {
																																						Hyp := __e.Get(1)
																																						_ = Hyp
																																						tmp20757 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																						_ = tmp20757

																																						tmp20758 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																						tmp20759 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20760 := PrimCons(tmp20759, Nil)

																																						tmp20761 := PrimCons(sym_h, tmp20760)

																																						tmp20762 := PrimCons(tmp20758, tmp20761)

																																						tmp20763 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																						tmp20764 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																						tmp20765 := PrimCons(tmp20764, Nil)

																																						tmp20766 := PrimCons(symlist, tmp20765)

																																						tmp20767 := PrimCons(tmp20766, Nil)

																																						tmp20768 := PrimCons(sym_h, tmp20767)

																																						tmp20769 := PrimCons(tmp20763, tmp20768)

																																						tmp20770 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																						tmp20771 := PrimCons(tmp20769, tmp20770)

																																						tmp20772 := PrimCons(tmp20762, tmp20771)

																																						__e.TailApply(PrimNS2Value(symbind), V2790, tmp20772, V2791, V2792)
																																						return

																																					}, 1)

																																					tmp20773 := PrimTail(V2526)

																																					__e.TailApply(tmp20756, tmp20773)
																																					return

																																				} else {
																																					tmp20755 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2546)

																																					if True == tmp20755 {
																																						tmp20733 := Call(__e, PrimNS2Value(symshen_4bindv), V2546, Nil, V2791)

																																						_ = tmp20733

																																						tmp20734 := MakeNative(func(__e *ControlFlow) {
																																							Result := __e.Get(1)
																																							_ = Result
																																							tmp20735 := Call(__e, PrimNS2Value(symshen_4unbindv), V2546, V2791)

																																							_ = tmp20735

																																							__e.Return(Result)
																																							return

																																						}, 1)

																																						tmp20736 := MakeNative(func(__e *ControlFlow) {
																																							Hyp := __e.Get(1)
																																							_ = Hyp
																																							tmp20737 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																							_ = tmp20737

																																							tmp20738 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																							tmp20739 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																							tmp20740 := PrimCons(tmp20739, Nil)

																																							tmp20741 := PrimCons(sym_h, tmp20740)

																																							tmp20742 := PrimCons(tmp20738, tmp20741)

																																							tmp20743 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																							tmp20744 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																							tmp20745 := PrimCons(tmp20744, Nil)

																																							tmp20746 := PrimCons(symlist, tmp20745)

																																							tmp20747 := PrimCons(tmp20746, Nil)

																																							tmp20748 := PrimCons(sym_h, tmp20747)

																																							tmp20749 := PrimCons(tmp20743, tmp20748)

																																							tmp20750 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																							tmp20751 := PrimCons(tmp20749, tmp20750)

																																							tmp20752 := PrimCons(tmp20742, tmp20751)

																																							__e.TailApply(PrimNS2Value(symbind), V2790, tmp20752, V2791, V2792)
																																							return

																																						}, 1)

																																						tmp20753 := PrimTail(V2526)

																																						tmp20754 := Call(__e, tmp20736, tmp20753)

																																						__e.TailApply(tmp20734, tmp20754)
																																						return

																																					} else {
																																						__e.Return(False)
																																						return
																																					}

																																				}

																																			}, 1)

																																			tmp20775 := PrimTail(V2535)

																																			tmp20776 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20775, V2791)

																																			tmp20777 := Call(__e, tmp20730, tmp20776)

																																			__e.TailApply(tmp20728, tmp20777)
																																			return

																																		} else {
																																			__e.Return(False)
																																			return
																																		}

																																	}

																																}, 1)

																																tmp20827 := PrimTail(V2543)

																																tmp20828 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20827, V2791)

																																__e.TailApply(tmp20724, tmp20828)
																																return

																															}, 1)

																															tmp20829 := PrimHead(V2543)

																															__e.TailApply(tmp20723, tmp20829)
																															return

																														} else {
																															tmp20722 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2543)

																															if True == tmp20722 {
																																tmp20668 := MakeNative(func(__e *ControlFlow) {
																																	A := __e.Get(1)
																																	_ = A
																																	tmp20669 := PrimCons(A, Nil)

																																	tmp20670 := Call(__e, PrimNS2Value(symshen_4bindv), V2543, tmp20669, V2791)

																																	_ = tmp20670

																																	tmp20671 := MakeNative(func(__e *ControlFlow) {
																																		Result := __e.Get(1)
																																		_ = Result
																																		tmp20672 := Call(__e, PrimNS2Value(symshen_4unbindv), V2543, V2791)

																																		_ = tmp20672

																																		__e.Return(Result)
																																		return

																																	}, 1)

																																	tmp20673 := MakeNative(func(__e *ControlFlow) {
																																		V2547 := __e.Get(1)
																																		_ = V2547
																																		tmp20717 := PrimEqual(Nil, V2547)

																																		if True == tmp20717 {
																																			tmp20699 := MakeNative(func(__e *ControlFlow) {
																																				Hyp := __e.Get(1)
																																				_ = Hyp
																																				tmp20700 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																				_ = tmp20700

																																				tmp20701 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																				tmp20702 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20703 := PrimCons(tmp20702, Nil)

																																				tmp20704 := PrimCons(sym_h, tmp20703)

																																				tmp20705 := PrimCons(tmp20701, tmp20704)

																																				tmp20706 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																				tmp20707 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																				tmp20708 := PrimCons(tmp20707, Nil)

																																				tmp20709 := PrimCons(symlist, tmp20708)

																																				tmp20710 := PrimCons(tmp20709, Nil)

																																				tmp20711 := PrimCons(sym_h, tmp20710)

																																				tmp20712 := PrimCons(tmp20706, tmp20711)

																																				tmp20713 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																				tmp20714 := PrimCons(tmp20712, tmp20713)

																																				tmp20715 := PrimCons(tmp20705, tmp20714)

																																				__e.TailApply(PrimNS2Value(symbind), V2790, tmp20715, V2791, V2792)
																																				return

																																			}, 1)

																																			tmp20716 := PrimTail(V2526)

																																			__e.TailApply(tmp20699, tmp20716)
																																			return

																																		} else {
																																			tmp20698 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2547)

																																			if True == tmp20698 {
																																				tmp20676 := Call(__e, PrimNS2Value(symshen_4bindv), V2547, Nil, V2791)

																																				_ = tmp20676

																																				tmp20677 := MakeNative(func(__e *ControlFlow) {
																																					Result := __e.Get(1)
																																					_ = Result
																																					tmp20678 := Call(__e, PrimNS2Value(symshen_4unbindv), V2547, V2791)

																																					_ = tmp20678

																																					__e.Return(Result)
																																					return

																																				}, 1)

																																				tmp20679 := MakeNative(func(__e *ControlFlow) {
																																					Hyp := __e.Get(1)
																																					_ = Hyp
																																					tmp20680 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																					_ = tmp20680

																																					tmp20681 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																					tmp20682 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20683 := PrimCons(tmp20682, Nil)

																																					tmp20684 := PrimCons(sym_h, tmp20683)

																																					tmp20685 := PrimCons(tmp20681, tmp20684)

																																					tmp20686 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																					tmp20687 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																					tmp20688 := PrimCons(tmp20687, Nil)

																																					tmp20689 := PrimCons(symlist, tmp20688)

																																					tmp20690 := PrimCons(tmp20689, Nil)

																																					tmp20691 := PrimCons(sym_h, tmp20690)

																																					tmp20692 := PrimCons(tmp20686, tmp20691)

																																					tmp20693 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																					tmp20694 := PrimCons(tmp20692, tmp20693)

																																					tmp20695 := PrimCons(tmp20685, tmp20694)

																																					__e.TailApply(PrimNS2Value(symbind), V2790, tmp20695, V2791, V2792)
																																					return

																																				}, 1)

																																				tmp20696 := PrimTail(V2526)

																																				tmp20697 := Call(__e, tmp20679, tmp20696)

																																				__e.TailApply(tmp20677, tmp20697)
																																				return

																																			} else {
																																				__e.Return(False)
																																				return
																																			}

																																		}

																																	}, 1)

																																	tmp20718 := PrimTail(V2535)

																																	tmp20719 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20718, V2791)

																																	tmp20720 := Call(__e, tmp20673, tmp20719)

																																	__e.TailApply(tmp20671, tmp20720)
																																	return

																																}, 1)

																																tmp20721 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																																__e.TailApply(tmp20668, tmp20721)
																																return

																															} else {
																																__e.Return(False)
																																return
																															}

																														}

																													}, 1)

																													tmp20831 := PrimTail(V2536)

																													tmp20832 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20831, V2791)

																													tmp20833 := Call(__e, tmp20665, tmp20832)

																													__e.TailApply(tmp20663, tmp20833)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}

																										}, 1)

																										tmp21004 := PrimHead(V2536)

																										tmp21005 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21004, V2791)

																										__e.TailApply(tmp20659, tmp21005)
																										return

																									} else {
																										tmp20658 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2536)

																										if True == tmp20658 {
																											tmp20603 := MakeNative(func(__e *ControlFlow) {
																												A := __e.Get(1)
																												_ = A
																												tmp20604 := PrimCons(A, Nil)

																												tmp20605 := PrimCons(symlist, tmp20604)

																												tmp20606 := Call(__e, PrimNS2Value(symshen_4bindv), V2536, tmp20605, V2791)

																												_ = tmp20606

																												tmp20607 := MakeNative(func(__e *ControlFlow) {
																													Result := __e.Get(1)
																													_ = Result
																													tmp20608 := Call(__e, PrimNS2Value(symshen_4unbindv), V2536, V2791)

																													_ = tmp20608

																													__e.Return(Result)
																													return

																												}, 1)

																												tmp20609 := MakeNative(func(__e *ControlFlow) {
																													V2548 := __e.Get(1)
																													_ = V2548
																													tmp20653 := PrimEqual(Nil, V2548)

																													if True == tmp20653 {
																														tmp20635 := MakeNative(func(__e *ControlFlow) {
																															Hyp := __e.Get(1)
																															_ = Hyp
																															tmp20636 := Call(__e, PrimNS2Value(symshen_4incinfs))

																															_ = tmp20636

																															tmp20637 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																															tmp20638 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																															tmp20639 := PrimCons(tmp20638, Nil)

																															tmp20640 := PrimCons(sym_h, tmp20639)

																															tmp20641 := PrimCons(tmp20637, tmp20640)

																															tmp20642 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																															tmp20643 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																															tmp20644 := PrimCons(tmp20643, Nil)

																															tmp20645 := PrimCons(symlist, tmp20644)

																															tmp20646 := PrimCons(tmp20645, Nil)

																															tmp20647 := PrimCons(sym_h, tmp20646)

																															tmp20648 := PrimCons(tmp20642, tmp20647)

																															tmp20649 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																															tmp20650 := PrimCons(tmp20648, tmp20649)

																															tmp20651 := PrimCons(tmp20641, tmp20650)

																															__e.TailApply(PrimNS2Value(symbind), V2790, tmp20651, V2791, V2792)
																															return

																														}, 1)

																														tmp20652 := PrimTail(V2526)

																														__e.TailApply(tmp20635, tmp20652)
																														return

																													} else {
																														tmp20634 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2548)

																														if True == tmp20634 {
																															tmp20612 := Call(__e, PrimNS2Value(symshen_4bindv), V2548, Nil, V2791)

																															_ = tmp20612

																															tmp20613 := MakeNative(func(__e *ControlFlow) {
																																Result := __e.Get(1)
																																_ = Result
																																tmp20614 := Call(__e, PrimNS2Value(symshen_4unbindv), V2548, V2791)

																																_ = tmp20614

																																__e.Return(Result)
																																return

																															}, 1)

																															tmp20615 := MakeNative(func(__e *ControlFlow) {
																																Hyp := __e.Get(1)
																																_ = Hyp
																																tmp20616 := Call(__e, PrimNS2Value(symshen_4incinfs))

																																_ = tmp20616

																																tmp20617 := Call(__e, PrimNS2Value(symshen_4lazyderef), X, V2791)

																																tmp20618 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																tmp20619 := PrimCons(tmp20618, Nil)

																																tmp20620 := PrimCons(sym_h, tmp20619)

																																tmp20621 := PrimCons(tmp20617, tmp20620)

																																tmp20622 := Call(__e, PrimNS2Value(symshen_4lazyderef), Y, V2791)

																																tmp20623 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2791)

																																tmp20624 := PrimCons(tmp20623, Nil)

																																tmp20625 := PrimCons(symlist, tmp20624)

																																tmp20626 := PrimCons(tmp20625, Nil)

																																tmp20627 := PrimCons(sym_h, tmp20626)

																																tmp20628 := PrimCons(tmp20622, tmp20627)

																																tmp20629 := Call(__e, PrimNS2Value(symshen_4lazyderef), Hyp, V2791)

																																tmp20630 := PrimCons(tmp20628, tmp20629)

																																tmp20631 := PrimCons(tmp20621, tmp20630)

																																__e.TailApply(PrimNS2Value(symbind), V2790, tmp20631, V2791, V2792)
																																return

																															}, 1)

																															tmp20632 := PrimTail(V2526)

																															tmp20633 := Call(__e, tmp20615, tmp20632)

																															__e.TailApply(tmp20613, tmp20633)
																															return

																														} else {
																															__e.Return(False)
																															return
																														}

																													}

																												}, 1)

																												tmp20654 := PrimTail(V2535)

																												tmp20655 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp20654, V2791)

																												tmp20656 := Call(__e, tmp20609, tmp20655)

																												__e.TailApply(tmp20607, tmp20656)
																												return

																											}, 1)

																											tmp20657 := Call(__e, PrimNS2Value(symshen_4newpv), V2791)

																											__e.TailApply(tmp20603, tmp20657)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}

																								}, 1)

																								tmp21007 := PrimHead(V2535)

																								tmp21008 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21007, V2791)

																								__e.TailApply(tmp20600, tmp21008)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp21010 := PrimTail(V2533)

																						tmp21011 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21010, V2791)

																						__e.TailApply(tmp20598, tmp21011)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp21013 := PrimHead(V2533)

																				tmp21014 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21013, V2791)

																				__e.TailApply(tmp20596, tmp21014)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp21016 := PrimTail(V2527)

																		tmp21017 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21016, V2791)

																		__e.TailApply(tmp20594, tmp21017)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}, 1)

																tmp21019 := PrimTail(V2531)

																tmp21020 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21019, V2791)

																__e.TailApply(tmp20592, tmp21020)
																return

															}, 1)

															tmp21021 := PrimHead(V2531)

															__e.TailApply(tmp20591, tmp21021)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp21023 := PrimTail(V2530)

													tmp21024 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21023, V2791)

													__e.TailApply(tmp20589, tmp21024)
													return

												}, 1)

												tmp21025 := PrimHead(V2530)

												__e.TailApply(tmp20588, tmp21025)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp21027 := PrimTail(V2528)

										tmp21028 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21027, V2791)

										__e.TailApply(tmp20586, tmp21028)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp21030 := PrimHead(V2528)

								tmp21031 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21030, V2791)

								__e.TailApply(tmp20584, tmp21031)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp21033 := PrimHead(V2527)

						tmp21034 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21033, V2791)

						__e.TailApply(tmp20582, tmp21034)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21036 := PrimHead(V2526)

				tmp21037 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21036, V2791)

				__e.TailApply(tmp20580, tmp21037)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21039 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2789, V2791)

		tmp21040 := Call(__e, tmp20578, tmp21039)

		__e.TailApply(tmp19447, tmp21040)
		return

	}, 4)

	tmp21041 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1hyps, tmp19446)

	_ = tmp21041

	tmp21042 := MakeNative(func(__e *ControlFlow) {
		V2809 := __e.Get(1)
		_ = V2809
		V2810 := __e.Get(2)
		_ = V2810
		V2811 := __e.Get(3)
		_ = V2811
		V2812 := __e.Get(4)
		_ = V2812
		tmp21054 := PrimNS3Value(symshen_4_dspy_d)

		if True == tmp21054 {
			tmp21044 := Call(__e, PrimNS2Value(symshen_4line))

			_ = tmp21044

			tmp21045 := Call(__e, PrimNS2Value(symshen_4deref), V2809, V2811)

			tmp21046 := Call(__e, PrimNS2Value(symshen_4show_1p), tmp21045)

			_ = tmp21046

			tmp21047 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

			_ = tmp21047

			tmp21048 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

			_ = tmp21048

			tmp21049 := Call(__e, PrimNS2Value(symshen_4deref), V2810, V2811)

			tmp21050 := Call(__e, PrimNS2Value(symshen_4show_1assumptions), tmp21049, MakeNumber(1))

			_ = tmp21050

			tmp21051 := Call(__e, PrimNS2Value(symstoutput))

			tmp21052 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString("\n> "), tmp21051)

			_ = tmp21052

			tmp21053 := Call(__e, PrimNS2Value(symshen_4pause_1for_1user))

			_ = tmp21053

			__e.TailApply(PrimNS2Value(symthaw), V2812)
			return

		} else {
			__e.TailApply(PrimNS2Value(symthaw), V2812)
			return
		}

	}, 4)

	tmp21055 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show, tmp21042)

	_ = tmp21055

	tmp21056 := MakeNative(func(__e *ControlFlow) {
		tmp21057 := MakeNative(func(__e *ControlFlow) {
			Infs := __e.Get(1)
			_ = Infs
			tmp21059 := PrimEqual(MakeNumber(1), Infs)

			var ifres21058 Obj

			if True == tmp21059 {
				ifres21058 = MakeString("")

			} else {
				ifres21058 = MakeString("s")

			}

			tmp21060 := Call(__e, PrimNS2Value(symshen_4app), ifres21058, MakeString(" \n?- "), symshen_4a)

			tmp21061 := PrimStringConcat(MakeString(" inference"), tmp21060)

			tmp21062 := Call(__e, PrimNS2Value(symshen_4app), Infs, tmp21061, symshen_4a)

			tmp21063 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp21062)

			tmp21064 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), tmp21063, tmp21064)
			return

		}, 1)

		tmp21065 := Call(__e, PrimNS2Value(syminferences))

		__e.TailApply(tmp21057, tmp21065)
		return

	}, 0)

	tmp21066 := Call(__e, PrimNS1Value(symns2_1set), symshen_4line, tmp21056)

	_ = tmp21066

	tmp21067 := MakeNative(func(__e *ControlFlow) {
		V2814 := __e.Get(1)
		_ = V2814
		tmp21099 := PrimIsPair(V2814)

		var ifres21079 Obj

		if True == tmp21099 {
			tmp21097 := PrimTail(V2814)

			tmp21098 := PrimIsPair(tmp21097)

			var ifres21081 Obj

			if True == tmp21098 {
				tmp21094 := PrimTail(V2814)

				tmp21095 := PrimHead(tmp21094)

				tmp21096 := PrimEqual(sym_h, tmp21095)

				var ifres21083 Obj

				if True == tmp21096 {
					tmp21091 := PrimTail(V2814)

					tmp21092 := PrimTail(tmp21091)

					tmp21093 := PrimIsPair(tmp21092)

					var ifres21085 Obj

					if True == tmp21093 {
						tmp21087 := PrimTail(V2814)

						tmp21088 := PrimTail(tmp21087)

						tmp21089 := PrimTail(tmp21088)

						tmp21090 := PrimEqual(Nil, tmp21089)

						var ifres21086 Obj

						if True == tmp21090 {
							ifres21086 = True

						} else {
							ifres21086 = False

						}

						ifres21085 = ifres21086

					} else {
						ifres21085 = False

					}

					var ifres21084 Obj

					if True == ifres21085 {
						ifres21084 = True

					} else {
						ifres21084 = False

					}

					ifres21083 = ifres21084

				} else {
					ifres21083 = False

				}

				var ifres21082 Obj

				if True == ifres21083 {
					ifres21082 = True

				} else {
					ifres21082 = False

				}

				ifres21081 = ifres21082

			} else {
				ifres21081 = False

			}

			var ifres21080 Obj

			if True == ifres21081 {
				ifres21080 = True

			} else {
				ifres21080 = False

			}

			ifres21079 = ifres21080

		} else {
			ifres21079 = False

		}

		if True == ifres21079 {
			tmp21071 := PrimHead(V2814)

			tmp21072 := PrimTail(V2814)

			tmp21073 := PrimTail(tmp21072)

			tmp21074 := PrimHead(tmp21073)

			tmp21075 := Call(__e, PrimNS2Value(symshen_4app), tmp21074, MakeString(""), symshen_4r)

			tmp21076 := PrimStringConcat(MakeString(" : "), tmp21075)

			tmp21077 := Call(__e, PrimNS2Value(symshen_4app), tmp21071, tmp21076, symshen_4r)

			tmp21078 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), tmp21077, tmp21078)
			return

		} else {
			tmp21069 := Call(__e, PrimNS2Value(symshen_4app), V2814, MakeString(""), symshen_4r)

			tmp21070 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), tmp21069, tmp21070)
			return

		}

	}, 1)

	tmp21100 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show_1p, tmp21067)

	_ = tmp21100

	tmp21101 := MakeNative(func(__e *ControlFlow) {
		V2819 := __e.Get(1)
		_ = V2819
		V2820 := __e.Get(2)
		_ = V2820
		tmp21113 := PrimEqual(Nil, V2819)

		if True == tmp21113 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp21112 := PrimIsPair(V2819)

			if True == tmp21112 {
				tmp21104 := Call(__e, PrimNS2Value(symshen_4app), V2820, MakeString(". "), symshen_4a)

				tmp21105 := Call(__e, PrimNS2Value(symstoutput))

				tmp21106 := Call(__e, PrimNS2Value(symshen_4prhush), tmp21104, tmp21105)

				_ = tmp21106

				tmp21107 := PrimHead(V2819)

				tmp21108 := Call(__e, PrimNS2Value(symshen_4show_1p), tmp21107)

				_ = tmp21108

				tmp21109 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

				_ = tmp21109

				tmp21110 := PrimTail(V2819)

				tmp21111 := PrimNumberAdd(V2820, MakeNumber(1))

				__e.TailApply(PrimNS2Value(symshen_4show_1assumptions), tmp21110, tmp21111)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4show_1assumptions)
				return
			}

		}

	}, 2)

	tmp21114 := Call(__e, PrimNS1Value(symns2_1set), symshen_4show_1assumptions, tmp21101)

	_ = tmp21114

	tmp21115 := MakeNative(func(__e *ControlFlow) {
		tmp21116 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			tmp21118 := PrimEqual(Byte, MakeNumber(94))

			if True == tmp21118 {
				__e.Return(PrimSimpleError(MakeString("input aborted\n")))
				return
			} else {
				__e.TailApply(PrimNS2Value(symnl), MakeNumber(1))
				return
			}

		}, 1)

		tmp21119 := Call(__e, PrimNS2Value(symstinput))

		tmp21120 := PrimReadByte(tmp21119)

		__e.TailApply(tmp21116, tmp21120)
		return

	}, 0)

	tmp21121 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pause_1for_1user, tmp21115)

	_ = tmp21121

	tmp21122 := MakeNative(func(__e *ControlFlow) {
		V2822 := __e.Get(1)
		_ = V2822
		tmp21123 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp21124 := Call(__e, PrimNS2Value(symassoc), V2822, tmp21123)

		__e.Return(PrimIsPair(tmp21124))
		return

	}, 1)

	tmp21125 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typedf_2, tmp21122)

	_ = tmp21125

	tmp21126 := MakeNative(func(__e *ControlFlow) {
		V2824 := __e.Get(1)
		_ = V2824
		__e.TailApply(PrimNS2Value(symconcat), symshen_4type_1signature_1of_1, V2824)
		return
	}, 1)

	tmp21127 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sigf, tmp21126)

	_ = tmp21127

	tmp21128 := MakeNative(func(__e *ControlFlow) {
		__e.TailApply(PrimNS2Value(symgensym), sym_e_e)
		return
	}, 0)

	tmp21129 := Call(__e, PrimNS1Value(symns2_1set), symshen_4placeholder, tmp21128)

	_ = tmp21129

	tmp21130 := MakeNative(func(__e *ControlFlow) {
		V2829 := __e.Get(1)
		_ = V2829
		V2830 := __e.Get(2)
		_ = V2830
		V2831 := __e.Get(3)
		_ = V2831
		V2832 := __e.Get(4)
		_ = V2832
		tmp21131 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp21291 := PrimEqual(Case, False)

			if True == tmp21291 {
				tmp21133 := MakeNative(func(__e *ControlFlow) {
					Case := __e.Get(1)
					_ = Case
					tmp21273 := PrimEqual(Case, False)

					if True == tmp21273 {
						tmp21135 := MakeNative(func(__e *ControlFlow) {
							Case := __e.Get(1)
							_ = Case
							tmp21255 := PrimEqual(Case, False)

							if True == tmp21255 {
								tmp21137 := MakeNative(func(__e *ControlFlow) {
									Case := __e.Get(1)
									_ = Case
									tmp21229 := PrimEqual(Case, False)

									if True == tmp21229 {
										tmp21139 := MakeNative(func(__e *ControlFlow) {
											V2517 := __e.Get(1)
											_ = V2517
											tmp21227 := PrimEqual(Nil, V2517)

											if True == tmp21227 {
												tmp21141 := MakeNative(func(__e *ControlFlow) {
													V2518 := __e.Get(1)
													_ = V2518
													tmp21225 := PrimIsPair(V2518)

													if True == tmp21225 {
														tmp21154 := MakeNative(func(__e *ControlFlow) {
															V2519 := __e.Get(1)
															_ = V2519
															tmp21222 := PrimEqual(symlist, V2519)

															if True == tmp21222 {
																tmp21192 := MakeNative(func(__e *ControlFlow) {
																	V2520 := __e.Get(1)
																	_ = V2520
																	tmp21219 := PrimIsPair(V2520)

																	if True == tmp21219 {
																		tmp21204 := MakeNative(func(__e *ControlFlow) {
																			A := __e.Get(1)
																			_ = A
																			tmp21205 := MakeNative(func(__e *ControlFlow) {
																				V2521 := __e.Get(1)
																				_ = V2521
																				tmp21215 := PrimEqual(Nil, V2521)

																				if True == tmp21215 {
																					tmp21214 := Call(__e, PrimNS2Value(symshen_4incinfs))

																					_ = tmp21214

																					__e.TailApply(PrimNS2Value(symthaw), V2832)
																					return

																				} else {
																					tmp21213 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2521)

																					if True == tmp21213 {
																						tmp21208 := Call(__e, PrimNS2Value(symshen_4bindv), V2521, Nil, V2831)

																						_ = tmp21208

																						tmp21209 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp21210 := Call(__e, PrimNS2Value(symshen_4unbindv), V2521, V2831)

																							_ = tmp21210

																							__e.Return(Result)
																							return

																						}, 1)

																						tmp21211 := Call(__e, PrimNS2Value(symshen_4incinfs))

																						_ = tmp21211

																						tmp21212 := Call(__e, PrimNS2Value(symthaw), V2832)

																						__e.TailApply(tmp21209, tmp21212)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}

																			}, 1)

																			tmp21216 := PrimTail(V2520)

																			tmp21217 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21216, V2831)

																			__e.TailApply(tmp21205, tmp21217)
																			return

																		}, 1)

																		tmp21218 := PrimHead(V2520)

																		__e.TailApply(tmp21204, tmp21218)
																		return

																	} else {
																		tmp21203 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2520)

																		if True == tmp21203 {
																			tmp21195 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp21196 := PrimCons(A, Nil)

																				tmp21197 := Call(__e, PrimNS2Value(symshen_4bindv), V2520, tmp21196, V2831)

																				_ = tmp21197

																				tmp21198 := MakeNative(func(__e *ControlFlow) {
																					Result := __e.Get(1)
																					_ = Result
																					tmp21199 := Call(__e, PrimNS2Value(symshen_4unbindv), V2520, V2831)

																					_ = tmp21199

																					__e.Return(Result)
																					return

																				}, 1)

																				tmp21200 := Call(__e, PrimNS2Value(symshen_4incinfs))

																				_ = tmp21200

																				tmp21201 := Call(__e, PrimNS2Value(symthaw), V2832)

																				__e.TailApply(tmp21198, tmp21201)
																				return

																			}, 1)

																			tmp21202 := Call(__e, PrimNS2Value(symshen_4newpv), V2831)

																			__e.TailApply(tmp21195, tmp21202)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}

																}, 1)

																tmp21220 := PrimTail(V2518)

																tmp21221 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21220, V2831)

																__e.TailApply(tmp21192, tmp21221)
																return

															} else {
																tmp21191 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2519)

																if True == tmp21191 {
																	tmp21157 := Call(__e, PrimNS2Value(symshen_4bindv), V2519, symlist, V2831)

																	_ = tmp21157

																	tmp21158 := MakeNative(func(__e *ControlFlow) {
																		Result := __e.Get(1)
																		_ = Result
																		tmp21159 := Call(__e, PrimNS2Value(symshen_4unbindv), V2519, V2831)

																		_ = tmp21159

																		__e.Return(Result)
																		return

																	}, 1)

																	tmp21160 := MakeNative(func(__e *ControlFlow) {
																		V2522 := __e.Get(1)
																		_ = V2522
																		tmp21187 := PrimIsPair(V2522)

																		if True == tmp21187 {
																			tmp21172 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp21173 := MakeNative(func(__e *ControlFlow) {
																					V2523 := __e.Get(1)
																					_ = V2523
																					tmp21183 := PrimEqual(Nil, V2523)

																					if True == tmp21183 {
																						tmp21182 := Call(__e, PrimNS2Value(symshen_4incinfs))

																						_ = tmp21182

																						__e.TailApply(PrimNS2Value(symthaw), V2832)
																						return

																					} else {
																						tmp21181 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2523)

																						if True == tmp21181 {
																							tmp21176 := Call(__e, PrimNS2Value(symshen_4bindv), V2523, Nil, V2831)

																							_ = tmp21176

																							tmp21177 := MakeNative(func(__e *ControlFlow) {
																								Result := __e.Get(1)
																								_ = Result
																								tmp21178 := Call(__e, PrimNS2Value(symshen_4unbindv), V2523, V2831)

																								_ = tmp21178

																								__e.Return(Result)
																								return

																							}, 1)

																							tmp21179 := Call(__e, PrimNS2Value(symshen_4incinfs))

																							_ = tmp21179

																							tmp21180 := Call(__e, PrimNS2Value(symthaw), V2832)

																							__e.TailApply(tmp21177, tmp21180)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}

																				}, 1)

																				tmp21184 := PrimTail(V2522)

																				tmp21185 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21184, V2831)

																				__e.TailApply(tmp21173, tmp21185)
																				return

																			}, 1)

																			tmp21186 := PrimHead(V2522)

																			__e.TailApply(tmp21172, tmp21186)
																			return

																		} else {
																			tmp21171 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2522)

																			if True == tmp21171 {
																				tmp21163 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp21164 := PrimCons(A, Nil)

																					tmp21165 := Call(__e, PrimNS2Value(symshen_4bindv), V2522, tmp21164, V2831)

																					_ = tmp21165

																					tmp21166 := MakeNative(func(__e *ControlFlow) {
																						Result := __e.Get(1)
																						_ = Result
																						tmp21167 := Call(__e, PrimNS2Value(symshen_4unbindv), V2522, V2831)

																						_ = tmp21167

																						__e.Return(Result)
																						return

																					}, 1)

																					tmp21168 := Call(__e, PrimNS2Value(symshen_4incinfs))

																					_ = tmp21168

																					tmp21169 := Call(__e, PrimNS2Value(symthaw), V2832)

																					__e.TailApply(tmp21166, tmp21169)
																					return

																				}, 1)

																				tmp21170 := Call(__e, PrimNS2Value(symshen_4newpv), V2831)

																				__e.TailApply(tmp21163, tmp21170)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp21188 := PrimTail(V2518)

																	tmp21189 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21188, V2831)

																	tmp21190 := Call(__e, tmp21160, tmp21189)

																	__e.TailApply(tmp21158, tmp21190)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}

														}, 1)

														tmp21223 := PrimHead(V2518)

														tmp21224 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21223, V2831)

														__e.TailApply(tmp21154, tmp21224)
														return

													} else {
														tmp21153 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2518)

														if True == tmp21153 {
															tmp21144 := MakeNative(func(__e *ControlFlow) {
																A := __e.Get(1)
																_ = A
																tmp21145 := PrimCons(A, Nil)

																tmp21146 := PrimCons(symlist, tmp21145)

																tmp21147 := Call(__e, PrimNS2Value(symshen_4bindv), V2518, tmp21146, V2831)

																_ = tmp21147

																tmp21148 := MakeNative(func(__e *ControlFlow) {
																	Result := __e.Get(1)
																	_ = Result
																	tmp21149 := Call(__e, PrimNS2Value(symshen_4unbindv), V2518, V2831)

																	_ = tmp21149

																	__e.Return(Result)
																	return

																}, 1)

																tmp21150 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp21150

																tmp21151 := Call(__e, PrimNS2Value(symthaw), V2832)

																__e.TailApply(tmp21148, tmp21151)
																return

															}, 1)

															tmp21152 := Call(__e, PrimNS2Value(symshen_4newpv), V2831)

															__e.TailApply(tmp21144, tmp21152)
															return

														} else {
															__e.Return(False)
															return
														}

													}

												}, 1)

												tmp21226 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2830, V2831)

												__e.TailApply(tmp21141, tmp21226)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp21228 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

										__e.TailApply(tmp21139, tmp21228)
										return

									} else {
										__e.Return(Case)
										return
									}

								}, 1)

								tmp21230 := MakeNative(func(__e *ControlFlow) {
									V2516 := __e.Get(1)
									_ = V2516
									tmp21252 := PrimEqual(symsymbol, V2516)

									if True == tmp21252 {
										tmp21245 := Call(__e, PrimNS2Value(symshen_4incinfs))

										_ = tmp21245

										tmp21246 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

										tmp21247 := PrimIsSymbol(tmp21246)

										tmp21248 := MakeNative(func(__e *ControlFlow) {
											tmp21249 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

											tmp21250 := Call(__e, PrimNS2Value(symshen_4ue_2), tmp21249)

											tmp21251 := PrimNot(tmp21250)

											__e.TailApply(PrimNS2Value(symfwhen), tmp21251, V2831, V2832)
											return

										}, 0)

										__e.TailApply(PrimNS2Value(symfwhen), tmp21247, V2831, tmp21248)
										return

									} else {
										tmp21244 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2516)

										if True == tmp21244 {
											tmp21233 := Call(__e, PrimNS2Value(symshen_4bindv), V2516, symsymbol, V2831)

											_ = tmp21233

											tmp21234 := MakeNative(func(__e *ControlFlow) {
												Result := __e.Get(1)
												_ = Result
												tmp21235 := Call(__e, PrimNS2Value(symshen_4unbindv), V2516, V2831)

												_ = tmp21235

												__e.Return(Result)
												return

											}, 1)

											tmp21236 := Call(__e, PrimNS2Value(symshen_4incinfs))

											_ = tmp21236

											tmp21237 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

											tmp21238 := PrimIsSymbol(tmp21237)

											tmp21239 := MakeNative(func(__e *ControlFlow) {
												tmp21240 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

												tmp21241 := Call(__e, PrimNS2Value(symshen_4ue_2), tmp21240)

												tmp21242 := PrimNot(tmp21241)

												__e.TailApply(PrimNS2Value(symfwhen), tmp21242, V2831, V2832)
												return

											}, 0)

											tmp21243 := Call(__e, PrimNS2Value(symfwhen), tmp21238, V2831, tmp21239)

											__e.TailApply(tmp21234, tmp21243)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp21253 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2830, V2831)

								tmp21254 := Call(__e, tmp21230, tmp21253)

								__e.TailApply(tmp21137, tmp21254)
								return

							} else {
								__e.Return(Case)
								return
							}

						}, 1)

						tmp21256 := MakeNative(func(__e *ControlFlow) {
							V2515 := __e.Get(1)
							_ = V2515
							tmp21270 := PrimEqual(symstring, V2515)

							if True == tmp21270 {
								tmp21267 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp21267

								tmp21268 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

								tmp21269 := PrimIsString(tmp21268)

								__e.TailApply(PrimNS2Value(symfwhen), tmp21269, V2831, V2832)
								return

							} else {
								tmp21266 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2515)

								if True == tmp21266 {
									tmp21259 := Call(__e, PrimNS2Value(symshen_4bindv), V2515, symstring, V2831)

									_ = tmp21259

									tmp21260 := MakeNative(func(__e *ControlFlow) {
										Result := __e.Get(1)
										_ = Result
										tmp21261 := Call(__e, PrimNS2Value(symshen_4unbindv), V2515, V2831)

										_ = tmp21261

										__e.Return(Result)
										return

									}, 1)

									tmp21262 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp21262

									tmp21263 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

									tmp21264 := PrimIsString(tmp21263)

									tmp21265 := Call(__e, PrimNS2Value(symfwhen), tmp21264, V2831, V2832)

									__e.TailApply(tmp21260, tmp21265)
									return

								} else {
									__e.Return(False)
									return
								}

							}

						}, 1)

						tmp21271 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2830, V2831)

						tmp21272 := Call(__e, tmp21256, tmp21271)

						__e.TailApply(tmp21135, tmp21272)
						return

					} else {
						__e.Return(Case)
						return
					}

				}, 1)

				tmp21274 := MakeNative(func(__e *ControlFlow) {
					V2514 := __e.Get(1)
					_ = V2514
					tmp21288 := PrimEqual(symboolean, V2514)

					if True == tmp21288 {
						tmp21285 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp21285

						tmp21286 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

						tmp21287 := Call(__e, PrimNS2Value(symboolean_2), tmp21286)

						__e.TailApply(PrimNS2Value(symfwhen), tmp21287, V2831, V2832)
						return

					} else {
						tmp21284 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2514)

						if True == tmp21284 {
							tmp21277 := Call(__e, PrimNS2Value(symshen_4bindv), V2514, symboolean, V2831)

							_ = tmp21277

							tmp21278 := MakeNative(func(__e *ControlFlow) {
								Result := __e.Get(1)
								_ = Result
								tmp21279 := Call(__e, PrimNS2Value(symshen_4unbindv), V2514, V2831)

								_ = tmp21279

								__e.Return(Result)
								return

							}, 1)

							tmp21280 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp21280

							tmp21281 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

							tmp21282 := Call(__e, PrimNS2Value(symboolean_2), tmp21281)

							tmp21283 := Call(__e, PrimNS2Value(symfwhen), tmp21282, V2831, V2832)

							__e.TailApply(tmp21278, tmp21283)
							return

						} else {
							__e.Return(False)
							return
						}

					}

				}, 1)

				tmp21289 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2830, V2831)

				tmp21290 := Call(__e, tmp21274, tmp21289)

				__e.TailApply(tmp21133, tmp21290)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp21292 := MakeNative(func(__e *ControlFlow) {
			V2513 := __e.Get(1)
			_ = V2513
			tmp21306 := PrimEqual(symnumber, V2513)

			if True == tmp21306 {
				tmp21303 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp21303

				tmp21304 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

				tmp21305 := PrimIsNumber(tmp21304)

				__e.TailApply(PrimNS2Value(symfwhen), tmp21305, V2831, V2832)
				return

			} else {
				tmp21302 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2513)

				if True == tmp21302 {
					tmp21295 := Call(__e, PrimNS2Value(symshen_4bindv), V2513, symnumber, V2831)

					_ = tmp21295

					tmp21296 := MakeNative(func(__e *ControlFlow) {
						Result := __e.Get(1)
						_ = Result
						tmp21297 := Call(__e, PrimNS2Value(symshen_4unbindv), V2513, V2831)

						_ = tmp21297

						__e.Return(Result)
						return

					}, 1)

					tmp21298 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp21298

					tmp21299 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2829, V2831)

					tmp21300 := PrimIsNumber(tmp21299)

					tmp21301 := Call(__e, PrimNS2Value(symfwhen), tmp21300, V2831, V2832)

					__e.TailApply(tmp21296, tmp21301)
					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 1)

		tmp21307 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2830, V2831)

		tmp21308 := Call(__e, tmp21292, tmp21307)

		__e.TailApply(tmp21131, tmp21308)
		return

	}, 4)

	tmp21309 := Call(__e, PrimNS1Value(symns2_1set), symshen_4base, tmp21130)

	_ = tmp21309

	tmp21310 := MakeNative(func(__e *ControlFlow) {
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
		tmp21311 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp21320 := PrimEqual(Case, False)

			if True == tmp21320 {
				tmp21313 := MakeNative(func(__e *ControlFlow) {
					V2510 := __e.Get(1)
					_ = V2510
					tmp21318 := PrimIsPair(V2510)

					if True == tmp21318 {
						tmp21315 := MakeNative(func(__e *ControlFlow) {
							Hyp := __e.Get(1)
							_ = Hyp
							tmp21316 := Call(__e, PrimNS2Value(symshen_4incinfs))

							_ = tmp21316

							__e.TailApply(PrimNS2Value(symshen_4by__hypothesis), V2838, V2839, Hyp, V2841, V2842)
							return

						}, 1)

						tmp21317 := PrimTail(V2510)

						__e.TailApply(tmp21315, tmp21317)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21319 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2840, V2841)

				__e.TailApply(tmp21313, tmp21319)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp21321 := MakeNative(func(__e *ControlFlow) {
			V2504 := __e.Get(1)
			_ = V2504
			tmp21354 := PrimIsPair(V2504)

			if True == tmp21354 {
				tmp21323 := MakeNative(func(__e *ControlFlow) {
					V2505 := __e.Get(1)
					_ = V2505
					tmp21351 := PrimIsPair(V2505)

					if True == tmp21351 {
						tmp21325 := MakeNative(func(__e *ControlFlow) {
							Y := __e.Get(1)
							_ = Y
							tmp21326 := MakeNative(func(__e *ControlFlow) {
								V2506 := __e.Get(1)
								_ = V2506
								tmp21347 := PrimIsPair(V2506)

								if True == tmp21347 {
									tmp21328 := MakeNative(func(__e *ControlFlow) {
										V2507 := __e.Get(1)
										_ = V2507
										tmp21344 := PrimEqual(sym_h, V2507)

										if True == tmp21344 {
											tmp21330 := MakeNative(func(__e *ControlFlow) {
												V2508 := __e.Get(1)
												_ = V2508
												tmp21341 := PrimIsPair(V2508)

												if True == tmp21341 {
													tmp21332 := MakeNative(func(__e *ControlFlow) {
														B := __e.Get(1)
														_ = B
														tmp21333 := MakeNative(func(__e *ControlFlow) {
															V2509 := __e.Get(1)
															_ = V2509
															tmp21337 := PrimEqual(Nil, V2509)

															if True == tmp21337 {
																tmp21335 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp21335

																tmp21336 := MakeNative(func(__e *ControlFlow) {
																	__e.TailApply(PrimNS2Value(symunify_b), V2839, B, V2841, V2842)
																	return
																}, 0)

																__e.TailApply(PrimNS2Value(symidentical), V2838, Y, V2841, tmp21336)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp21338 := PrimTail(V2508)

														tmp21339 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21338, V2841)

														__e.TailApply(tmp21333, tmp21339)
														return

													}, 1)

													tmp21340 := PrimHead(V2508)

													__e.TailApply(tmp21332, tmp21340)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp21342 := PrimTail(V2506)

											tmp21343 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21342, V2841)

											__e.TailApply(tmp21330, tmp21343)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp21345 := PrimHead(V2506)

									tmp21346 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21345, V2841)

									__e.TailApply(tmp21328, tmp21346)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp21348 := PrimTail(V2505)

							tmp21349 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21348, V2841)

							__e.TailApply(tmp21326, tmp21349)
							return

						}, 1)

						tmp21350 := PrimHead(V2505)

						__e.TailApply(tmp21325, tmp21350)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21352 := PrimHead(V2504)

				tmp21353 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21352, V2841)

				__e.TailApply(tmp21323, tmp21353)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21355 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2840, V2841)

		tmp21356 := Call(__e, tmp21321, tmp21355)

		__e.TailApply(tmp21311, tmp21356)
		return

	}, 5)

	tmp21357 := Call(__e, PrimNS1Value(symns2_1set), symshen_4by__hypothesis, tmp21310)

	_ = tmp21357

	tmp21358 := MakeNative(func(__e *ControlFlow) {
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
		tmp21359 := MakeNative(func(__e *ControlFlow) {
			V2498 := __e.Get(1)
			_ = V2498
			tmp21387 := PrimIsPair(V2498)

			if True == tmp21387 {
				tmp21361 := MakeNative(func(__e *ControlFlow) {
					V2499 := __e.Get(1)
					_ = V2499
					tmp21384 := PrimEqual(symdefine, V2499)

					if True == tmp21384 {
						tmp21363 := MakeNative(func(__e *ControlFlow) {
							V2500 := __e.Get(1)
							_ = V2500
							tmp21381 := PrimIsPair(V2500)

							if True == tmp21381 {
								tmp21365 := MakeNative(func(__e *ControlFlow) {
									F := __e.Get(1)
									_ = F
									tmp21366 := MakeNative(func(__e *ControlFlow) {
										X := __e.Get(1)
										_ = X
										tmp21367 := MakeNative(func(__e *ControlFlow) {
											Y := __e.Get(1)
											_ = Y
											tmp21368 := MakeNative(func(__e *ControlFlow) {
												E := __e.Get(1)
												_ = E
												tmp21369 := Call(__e, PrimNS2Value(symshen_4incinfs))

												_ = tmp21369

												tmp21370 := MakeNative(func(__e *ControlFlow) {
													Y := __e.Get(1)
													_ = Y
													__e.TailApply(PrimNS2Value(symshen_4_5sig_7rules_6), Y)
													return
												}, 1)

												tmp21371 := MakeNative(func(__e *ControlFlow) {
													E := __e.Get(1)
													_ = E
													tmp21375 := PrimIsPair(E)

													if True == tmp21375 {
														tmp21373 := Call(__e, PrimNS2Value(symshen_4app), E, MakeString("\n"), symshen_4s)

														tmp21374 := PrimStringConcat(MakeString("parse error here: "), tmp21373)

														__e.Return(PrimSimpleError(tmp21374))
														return

													} else {
														__e.Return(PrimSimpleError(MakeString("parse error\n")))
														return
													}

												}, 1)

												tmp21376 := Call(__e, PrimNS2Value(symcompile), tmp21370, X, tmp21371)

												__e.TailApply(PrimNS2Value(symshen_4t_d_1defh), tmp21376, F, V2849, V2850, V2851, V2852)
												return

											}, 1)

											tmp21377 := Call(__e, PrimNS2Value(symshen_4newpv), V2851)

											__e.TailApply(tmp21368, tmp21377)
											return

										}, 1)

										tmp21378 := Call(__e, PrimNS2Value(symshen_4newpv), V2851)

										__e.TailApply(tmp21367, tmp21378)
										return

									}, 1)

									tmp21379 := PrimTail(V2500)

									__e.TailApply(tmp21366, tmp21379)
									return

								}, 1)

								tmp21380 := PrimHead(V2500)

								__e.TailApply(tmp21365, tmp21380)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp21382 := PrimTail(V2498)

						tmp21383 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21382, V2851)

						__e.TailApply(tmp21363, tmp21383)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21385 := PrimHead(V2498)

				tmp21386 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21385, V2851)

				__e.TailApply(tmp21361, tmp21386)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21388 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2848, V2851)

		__e.TailApply(tmp21359, tmp21388)
		return

	}, 5)

	tmp21389 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1def, tmp21358)

	_ = tmp21389

	tmp21390 := MakeNative(func(__e *ControlFlow) {
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
		tmp21391 := MakeNative(func(__e *ControlFlow) {
			V2494 := __e.Get(1)
			_ = V2494
			tmp21399 := PrimIsPair(V2494)

			if True == tmp21399 {
				tmp21393 := MakeNative(func(__e *ControlFlow) {
					Sig := __e.Get(1)
					_ = Sig
					tmp21394 := MakeNative(func(__e *ControlFlow) {
						Rules := __e.Get(1)
						_ = Rules
						tmp21395 := Call(__e, PrimNS2Value(symshen_4incinfs))

						_ = tmp21395

						tmp21396 := Call(__e, PrimNS2Value(symshen_4ue_1sig), Sig)

						__e.TailApply(PrimNS2Value(symshen_4t_d_1defhh), Sig, tmp21396, V2860, V2861, V2862, Rules, V2863, V2864)
						return

					}, 1)

					tmp21397 := PrimTail(V2494)

					__e.TailApply(tmp21394, tmp21397)
					return

				}, 1)

				tmp21398 := PrimHead(V2494)

				__e.TailApply(tmp21393, tmp21398)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21400 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2859, V2863)

		__e.TailApply(tmp21391, tmp21400)
		return

	}, 6)

	tmp21401 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1defh, tmp21390)

	_ = tmp21401

	tmp21402 := MakeNative(func(__e *ControlFlow) {
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
		tmp21403 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp21403

		tmp21404 := PrimCons(V2874, Nil)

		tmp21405 := PrimCons(sym_h, tmp21404)

		tmp21406 := PrimCons(V2875, tmp21405)

		tmp21407 := PrimCons(tmp21406, V2877)

		tmp21408 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4memo), V2875, V2873, V2876, V2879, V2880)
			return
		}, 0)

		__e.TailApply(PrimNS2Value(symshen_4t_d_1rules), V2878, V2874, MakeNumber(1), V2875, tmp21407, V2879, tmp21408)
		return

	}, 8)

	tmp21409 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1defhh, tmp21402)

	_ = tmp21409

	tmp21410 := MakeNative(func(__e *ControlFlow) {
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
		tmp21411 := MakeNative(func(__e *ControlFlow) {
			Jnk := __e.Get(1)
			_ = Jnk
			tmp21412 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp21412

			tmp21413 := MakeNative(func(__e *ControlFlow) {
				tmp21414 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2886, V2889)

				tmp21415 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2888, V2889)

				tmp21416 := Call(__e, PrimNS2Value(symdeclare), tmp21414, tmp21415)

				__e.TailApply(PrimNS2Value(symbind), Jnk, tmp21416, V2889, V2890)
				return

			}, 0)

			__e.TailApply(PrimNS2Value(symunify_b), V2888, V2887, V2889, tmp21413)
			return

		}, 1)

		tmp21417 := Call(__e, PrimNS2Value(symshen_4newpv), V2889)

		__e.TailApply(tmp21411, tmp21417)
		return

	}, 5)

	tmp21418 := Call(__e, PrimNS1Value(symns2_1set), symshen_4memo, tmp21410)

	_ = tmp21418

	tmp21419 := MakeNative(func(__e *ControlFlow) {
		V2892 := __e.Get(1)
		_ = V2892
		tmp21420 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5signature_6 := __e.Get(1)
			_ = Parse__shen_4_5signature_6
			tmp21432 := Call(__e, PrimNS2Value(symfail))

			tmp21433 := PrimEqual(tmp21432, Parse__shen_4_5signature_6)

			tmp21434 := PrimNot(tmp21433)

			if True == tmp21434 {
				tmp21422 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5non_1ll_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5non_1ll_1rules_6
					tmp21428 := Call(__e, PrimNS2Value(symfail))

					tmp21429 := PrimEqual(tmp21428, Parse__shen_4_5non_1ll_1rules_6)

					tmp21430 := PrimNot(tmp21429)

					if True == tmp21430 {
						tmp21424 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)

						tmp21425 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5signature_6)

						tmp21426 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

						tmp21427 := PrimCons(tmp21425, tmp21426)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp21424, tmp21427)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp21431 := Call(__e, PrimNS2Value(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5signature_6)

				__e.TailApply(tmp21422, tmp21431)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp21435 := Call(__e, PrimNS2Value(symshen_4_5signature_6), V2892)

		__e.TailApply(tmp21420, tmp21435)
		return

	}, 1)

	tmp21436 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sig_7rules_6, tmp21419)

	_ = tmp21436

	tmp21437 := MakeNative(func(__e *ControlFlow) {
		V2894 := __e.Get(1)
		_ = V2894
		tmp21438 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp21449 := Call(__e, PrimNS2Value(symfail))

			tmp21450 := PrimEqual(YaccParse, tmp21449)

			if True == tmp21450 {
				tmp21440 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5rule_6 := __e.Get(1)
					_ = Parse__shen_4_5rule_6
					tmp21445 := Call(__e, PrimNS2Value(symfail))

					tmp21446 := PrimEqual(tmp21445, Parse__shen_4_5rule_6)

					tmp21447 := PrimNot(tmp21446)

					if True == tmp21447 {
						tmp21442 := PrimHead(Parse__shen_4_5rule_6)

						tmp21443 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rule_6)

						tmp21444 := PrimCons(tmp21443, Nil)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp21442, tmp21444)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp21448 := Call(__e, PrimNS2Value(symshen_4_5rule_6), V2894)

				__e.TailApply(tmp21440, tmp21448)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp21451 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5rule_6 := __e.Get(1)
			_ = Parse__shen_4_5rule_6
			tmp21463 := Call(__e, PrimNS2Value(symfail))

			tmp21464 := PrimEqual(tmp21463, Parse__shen_4_5rule_6)

			tmp21465 := PrimNot(tmp21464)

			if True == tmp21465 {
				tmp21453 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5non_1ll_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5non_1ll_1rules_6
					tmp21459 := Call(__e, PrimNS2Value(symfail))

					tmp21460 := PrimEqual(tmp21459, Parse__shen_4_5non_1ll_1rules_6)

					tmp21461 := PrimNot(tmp21460)

					if True == tmp21461 {
						tmp21455 := PrimHead(Parse__shen_4_5non_1ll_1rules_6)

						tmp21456 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5rule_6)

						tmp21457 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

						tmp21458 := PrimCons(tmp21456, tmp21457)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp21455, tmp21458)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp21462 := Call(__e, PrimNS2Value(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5rule_6)

				__e.TailApply(tmp21453, tmp21462)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp21466 := Call(__e, PrimNS2Value(symshen_4_5rule_6), V2894)

		tmp21467 := Call(__e, tmp21451, tmp21466)

		__e.TailApply(tmp21438, tmp21467)
		return

	}, 1)

	tmp21468 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5non_1ll_1rules_6, tmp21437)

	_ = tmp21468

	tmp21469 := MakeNative(func(__e *ControlFlow) {
		V2896 := __e.Get(1)
		_ = V2896
		tmp21489 := PrimIsPair(V2896)

		var ifres21476 Obj

		if True == tmp21489 {
			tmp21487 := PrimTail(V2896)

			tmp21488 := PrimIsPair(tmp21487)

			var ifres21478 Obj

			if True == tmp21488 {
				tmp21484 := PrimTail(V2896)

				tmp21485 := PrimTail(tmp21484)

				tmp21486 := PrimEqual(Nil, tmp21485)

				var ifres21480 Obj

				if True == tmp21486 {
					tmp21482 := PrimHead(V2896)

					tmp21483 := PrimEqual(tmp21482, symprotect)

					var ifres21481 Obj

					if True == tmp21483 {
						ifres21481 = True

					} else {
						ifres21481 = False

					}

					ifres21480 = ifres21481

				} else {
					ifres21480 = False

				}

				var ifres21479 Obj

				if True == ifres21480 {
					ifres21479 = True

				} else {
					ifres21479 = False

				}

				ifres21478 = ifres21479

			} else {
				ifres21478 = False

			}

			var ifres21477 Obj

			if True == ifres21478 {
				ifres21477 = True

			} else {
				ifres21477 = False

			}

			ifres21476 = ifres21477

		} else {
			ifres21476 = False

		}

		if True == ifres21476 {
			__e.Return(V2896)
			return
		} else {
			tmp21475 := PrimIsPair(V2896)

			if True == tmp21475 {
				tmp21474 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4ue), Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp21474, V2896)
				return

			} else {
				tmp21473 := PrimIsVariable(V2896)

				if True == tmp21473 {
					__e.TailApply(PrimNS2Value(symconcat), sym_e_e, V2896)
					return
				} else {
					__e.Return(V2896)
					return
				}

			}

		}

	}, 1)

	tmp21490 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue, tmp21469)

	_ = tmp21490

	tmp21491 := MakeNative(func(__e *ControlFlow) {
		V2898 := __e.Get(1)
		_ = V2898
		tmp21496 := PrimIsPair(V2898)

		if True == tmp21496 {
			tmp21495 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4ue_1sig), Z)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp21495, V2898)
			return

		} else {
			tmp21494 := PrimIsVariable(V2898)

			if True == tmp21494 {
				__e.TailApply(PrimNS2Value(symconcat), sym_e_e_e, V2898)
				return
			} else {
				__e.Return(V2898)
				return
			}

		}

	}, 1)

	tmp21497 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_1sig, tmp21491)

	_ = tmp21497

	tmp21498 := MakeNative(func(__e *ControlFlow) {
		V2904 := __e.Get(1)
		_ = V2904
		tmp21506 := Call(__e, PrimNS2Value(symshen_4ue_2), V2904)

		if True == tmp21506 {
			__e.Return(PrimCons(V2904, Nil))
			return
		} else {
			tmp21505 := PrimIsPair(V2904)

			if True == tmp21505 {
				tmp21501 := PrimHead(V2904)

				tmp21502 := Call(__e, PrimNS2Value(symshen_4ues), tmp21501)

				tmp21503 := PrimTail(V2904)

				tmp21504 := Call(__e, PrimNS2Value(symshen_4ues), tmp21503)

				__e.TailApply(PrimNS2Value(symunion), tmp21502, tmp21504)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp21507 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ues, tmp21498)

	_ = tmp21507

	tmp21508 := MakeNative(func(__e *ControlFlow) {
		V2906 := __e.Get(1)
		_ = V2906
		tmp21513 := PrimIsSymbol(V2906)

		if True == tmp21513 {
			tmp21511 := PrimStr(V2906)

			tmp21512 := Call(__e, PrimNS2Value(symshen_4ue_1h_2), tmp21511)

			if True == tmp21512 {
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

	tmp21514 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_2, tmp21508)

	_ = tmp21514

	tmp21515 := MakeNative(func(__e *ControlFlow) {
		V2914 := __e.Get(1)
		_ = V2914
		tmp21530 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2914)

		var ifres21517 Obj

		if True == tmp21530 {
			tmp21528 := PrimPos(V2914, MakeNumber(0))

			tmp21529 := PrimEqual(MakeString("&"), tmp21528)

			var ifres21519 Obj

			if True == tmp21529 {
				tmp21526 := PrimTailString(V2914)

				tmp21527 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp21526)

				var ifres21521 Obj

				if True == tmp21527 {
					tmp21523 := PrimTailString(V2914)

					tmp21524 := PrimPos(tmp21523, MakeNumber(0))

					tmp21525 := PrimEqual(MakeString("&"), tmp21524)

					var ifres21522 Obj

					if True == tmp21525 {
						ifres21522 = True

					} else {
						ifres21522 = False

					}

					ifres21521 = ifres21522

				} else {
					ifres21521 = False

				}

				var ifres21520 Obj

				if True == ifres21521 {
					ifres21520 = True

				} else {
					ifres21520 = False

				}

				ifres21519 = ifres21520

			} else {
				ifres21519 = False

			}

			var ifres21518 Obj

			if True == ifres21519 {
				ifres21518 = True

			} else {
				ifres21518 = False

			}

			ifres21517 = ifres21518

		} else {
			ifres21517 = False

		}

		if True == ifres21517 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp21531 := Call(__e, PrimNS1Value(symns2_1set), symshen_4ue_1h_2, tmp21515)

	_ = tmp21531

	tmp21532 := MakeNative(func(__e *ControlFlow) {
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
		tmp21533 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp21534 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp21563 := PrimEqual(Case, False)

				if True == tmp21563 {
					tmp21536 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp21548 := PrimEqual(Case, False)

						if True == tmp21548 {
							tmp21538 := MakeNative(func(__e *ControlFlow) {
								Err := __e.Get(1)
								_ = Err
								tmp21539 := Call(__e, PrimNS2Value(symshen_4incinfs))

								_ = tmp21539

								tmp21540 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2924, V2927)

								tmp21541 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2925, V2927)

								tmp21542 := Call(__e, PrimNS2Value(symshen_4app), tmp21541, MakeString(""), symshen_4a)

								tmp21543 := PrimStringConcat(MakeString(" of "), tmp21542)

								tmp21544 := Call(__e, PrimNS2Value(symshen_4app), tmp21540, tmp21543, symshen_4a)

								tmp21545 := PrimStringConcat(MakeString("type error in rule "), tmp21544)

								tmp21546 := PrimSimpleError(tmp21545)

								__e.TailApply(PrimNS2Value(symbind), Err, tmp21546, V2927, V2928)
								return

							}, 1)

							tmp21547 := Call(__e, PrimNS2Value(symshen_4newpv), V2927)

							__e.TailApply(tmp21538, tmp21547)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp21549 := MakeNative(func(__e *ControlFlow) {
						V2479 := __e.Get(1)
						_ = V2479
						tmp21560 := PrimIsPair(V2479)

						if True == tmp21560 {
							tmp21551 := MakeNative(func(__e *ControlFlow) {
								Rule := __e.Get(1)
								_ = Rule
								tmp21552 := MakeNative(func(__e *ControlFlow) {
									Rules := __e.Get(1)
									_ = Rules
									tmp21553 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp21553

									tmp21554 := Call(__e, PrimNS2Value(symshen_4ue), Rule)

									tmp21555 := MakeNative(func(__e *ControlFlow) {
										tmp21556 := MakeNative(func(__e *ControlFlow) {
											tmp21557 := PrimNumberAdd(V2924, MakeNumber(1))

											__e.TailApply(PrimNS2Value(symshen_4t_d_1rules), Rules, V2923, tmp21557, V2925, V2926, V2927, V2928)
											return

										}, 0)

										__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2927, tmp21556)
										return

									}, 0)

									__e.TailApply(PrimNS2Value(symshen_4t_d_1rule), tmp21554, V2923, V2926, V2927, tmp21555)
									return

								}, 1)

								tmp21558 := PrimTail(V2479)

								__e.TailApply(tmp21552, tmp21558)
								return

							}, 1)

							tmp21559 := PrimHead(V2479)

							__e.TailApply(tmp21551, tmp21559)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp21561 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2922, V2927)

					tmp21562 := Call(__e, tmp21549, tmp21561)

					__e.TailApply(tmp21536, tmp21562)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp21564 := MakeNative(func(__e *ControlFlow) {
				V2478 := __e.Get(1)
				_ = V2478
				tmp21567 := PrimEqual(Nil, V2478)

				if True == tmp21567 {
					tmp21566 := Call(__e, PrimNS2Value(symshen_4incinfs))

					_ = tmp21566

					__e.TailApply(PrimNS2Value(symthaw), V2928)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp21568 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2922, V2927)

			tmp21569 := Call(__e, tmp21564, tmp21568)

			tmp21570 := Call(__e, tmp21534, tmp21569)

			__e.TailApply(PrimNS2Value(symshen_4cutpoint), Throwcontrol, tmp21570)
			return

		}, 1)

		tmp21571 := Call(__e, PrimNS2Value(symshen_4catchpoint))

		__e.TailApply(tmp21533, tmp21571)
		return

	}, 7)

	tmp21572 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1rules, tmp21532)

	_ = tmp21572

	tmp21573 := MakeNative(func(__e *ControlFlow) {
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
		tmp21574 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp21575 := MakeNative(func(__e *ControlFlow) {
				V2470 := __e.Get(1)
				_ = V2470
				tmp21602 := PrimIsPair(V2470)

				if True == tmp21602 {
					tmp21577 := MakeNative(func(__e *ControlFlow) {
						Patterns := __e.Get(1)
						_ = Patterns
						tmp21578 := MakeNative(func(__e *ControlFlow) {
							V2471 := __e.Get(1)
							_ = V2471
							tmp21598 := PrimIsPair(V2471)

							if True == tmp21598 {
								tmp21580 := MakeNative(func(__e *ControlFlow) {
									Action := __e.Get(1)
									_ = Action
									tmp21581 := MakeNative(func(__e *ControlFlow) {
										V2472 := __e.Get(1)
										_ = V2472
										tmp21594 := PrimEqual(Nil, V2472)

										if True == tmp21594 {
											tmp21583 := MakeNative(func(__e *ControlFlow) {
												NewHyps := __e.Get(1)
												_ = NewHyps
												tmp21584 := Call(__e, PrimNS2Value(symshen_4incinfs))

												_ = tmp21584

												tmp21585 := Call(__e, PrimNS2Value(symshen_4placeholders), Patterns)

												tmp21586 := MakeNative(func(__e *ControlFlow) {
													tmp21587 := MakeNative(func(__e *ControlFlow) {
														tmp21588 := MakeNative(func(__e *ControlFlow) {
															tmp21589 := Call(__e, PrimNS2Value(symshen_4ue), Action)

															tmp21590 := Call(__e, PrimNS2Value(symshen_4curry), tmp21589)

															tmp21591 := Call(__e, PrimNS2Value(symshen_4result_1type), Patterns, V2935)

															tmp21592 := Call(__e, PrimNS2Value(symshen_4patthyps), Patterns, V2935, V2936)

															__e.TailApply(PrimNS2Value(symshen_4t_d_1action), tmp21590, tmp21591, tmp21592, V2937, V2938)
															return

														}, 0)

														__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2937, tmp21588)
														return

													}, 0)

													__e.TailApply(PrimNS2Value(symshen_4t_d_1patterns), Patterns, V2935, NewHyps, V2937, tmp21587)
													return

												}, 0)

												__e.TailApply(PrimNS2Value(symshen_4newhyps), tmp21585, V2936, NewHyps, V2937, tmp21586)
												return

											}, 1)

											tmp21593 := Call(__e, PrimNS2Value(symshen_4newpv), V2937)

											__e.TailApply(tmp21583, tmp21593)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp21595 := PrimTail(V2471)

									tmp21596 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21595, V2937)

									__e.TailApply(tmp21581, tmp21596)
									return

								}, 1)

								tmp21597 := PrimHead(V2471)

								__e.TailApply(tmp21580, tmp21597)
								return

							} else {
								__e.Return(False)
								return
							}

						}, 1)

						tmp21599 := PrimTail(V2470)

						tmp21600 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21599, V2937)

						__e.TailApply(tmp21578, tmp21600)
						return

					}, 1)

					tmp21601 := PrimHead(V2470)

					__e.TailApply(tmp21577, tmp21601)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp21603 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2934, V2937)

			tmp21604 := Call(__e, tmp21575, tmp21603)

			__e.TailApply(PrimNS2Value(symshen_4cutpoint), Throwcontrol, tmp21604)
			return

		}, 1)

		tmp21605 := Call(__e, PrimNS2Value(symshen_4catchpoint))

		__e.TailApply(tmp21574, tmp21605)
		return

	}, 5)

	tmp21606 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1rule, tmp21573)

	_ = tmp21606

	tmp21607 := MakeNative(func(__e *ControlFlow) {
		V2944 := __e.Get(1)
		_ = V2944
		tmp21615 := Call(__e, PrimNS2Value(symshen_4ue_2), V2944)

		if True == tmp21615 {
			__e.Return(PrimCons(V2944, Nil))
			return
		} else {
			tmp21614 := PrimIsPair(V2944)

			if True == tmp21614 {
				tmp21610 := PrimHead(V2944)

				tmp21611 := Call(__e, PrimNS2Value(symshen_4placeholders), tmp21610)

				tmp21612 := PrimTail(V2944)

				tmp21613 := Call(__e, PrimNS2Value(symshen_4placeholders), tmp21612)

				__e.TailApply(PrimNS2Value(symunion), tmp21611, tmp21613)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp21616 := Call(__e, PrimNS1Value(symns2_1set), symshen_4placeholders, tmp21607)

	_ = tmp21616

	tmp21617 := MakeNative(func(__e *ControlFlow) {
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
		tmp21618 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp21782 := PrimEqual(Case, False)

			if True == tmp21782 {
				tmp21620 := MakeNative(func(__e *ControlFlow) {
					V2458 := __e.Get(1)
					_ = V2458
					tmp21780 := PrimIsPair(V2458)

					if True == tmp21780 {
						tmp21622 := MakeNative(func(__e *ControlFlow) {
							V2453 := __e.Get(1)
							_ = V2453
							tmp21623 := MakeNative(func(__e *ControlFlow) {
								Vs := __e.Get(1)
								_ = Vs
								tmp21624 := MakeNative(func(__e *ControlFlow) {
									V2459 := __e.Get(1)
									_ = V2459
									tmp21776 := PrimIsPair(V2459)

									if True == tmp21776 {
										tmp21644 := MakeNative(func(__e *ControlFlow) {
											V2460 := __e.Get(1)
											_ = V2460
											tmp21773 := PrimIsPair(V2460)

											if True == tmp21773 {
												tmp21663 := MakeNative(func(__e *ControlFlow) {
													V := __e.Get(1)
													_ = V
													tmp21664 := MakeNative(func(__e *ControlFlow) {
														V2461 := __e.Get(1)
														_ = V2461
														tmp21769 := PrimIsPair(V2461)

														if True == tmp21769 {
															tmp21680 := MakeNative(func(__e *ControlFlow) {
																V2462 := __e.Get(1)
																_ = V2462
																tmp21766 := PrimEqual(sym_h, V2462)

																if True == tmp21766 {
																	tmp21727 := MakeNative(func(__e *ControlFlow) {
																		V2463 := __e.Get(1)
																		_ = V2463
																		tmp21763 := PrimIsPair(V2463)

																		if True == tmp21763 {
																			tmp21742 := MakeNative(func(__e *ControlFlow) {
																				A := __e.Get(1)
																				_ = A
																				tmp21743 := MakeNative(func(__e *ControlFlow) {
																					V2464 := __e.Get(1)
																					_ = V2464
																					tmp21759 := PrimEqual(Nil, V2464)

																					if True == tmp21759 {
																						tmp21755 := MakeNative(func(__e *ControlFlow) {
																							NewHyp := __e.Get(1)
																							_ = NewHyp
																							tmp21756 := Call(__e, PrimNS2Value(symshen_4incinfs))

																							_ = tmp21756

																							tmp21757 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																								return
																							}, 0)

																							__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21757)
																							return

																						}, 1)

																						tmp21758 := PrimTail(V2459)

																						__e.TailApply(tmp21755, tmp21758)
																						return

																					} else {
																						tmp21754 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2464)

																						if True == tmp21754 {
																							tmp21746 := Call(__e, PrimNS2Value(symshen_4bindv), V2464, Nil, V2953)

																							_ = tmp21746

																							tmp21747 := MakeNative(func(__e *ControlFlow) {
																								Result := __e.Get(1)
																								_ = Result
																								tmp21748 := Call(__e, PrimNS2Value(symshen_4unbindv), V2464, V2953)

																								_ = tmp21748

																								__e.Return(Result)
																								return

																							}, 1)

																							tmp21749 := MakeNative(func(__e *ControlFlow) {
																								NewHyp := __e.Get(1)
																								_ = NewHyp
																								tmp21750 := Call(__e, PrimNS2Value(symshen_4incinfs))

																								_ = tmp21750

																								tmp21751 := MakeNative(func(__e *ControlFlow) {
																									__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																									return
																								}, 0)

																								__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21751)
																								return

																							}, 1)

																							tmp21752 := PrimTail(V2459)

																							tmp21753 := Call(__e, tmp21749, tmp21752)

																							__e.TailApply(tmp21747, tmp21753)
																							return

																						} else {
																							__e.Return(False)
																							return
																						}

																					}

																				}, 1)

																				tmp21760 := PrimTail(V2463)

																				tmp21761 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21760, V2953)

																				__e.TailApply(tmp21743, tmp21761)
																				return

																			}, 1)

																			tmp21762 := PrimHead(V2463)

																			__e.TailApply(tmp21742, tmp21762)
																			return

																		} else {
																			tmp21741 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2463)

																			if True == tmp21741 {
																				tmp21730 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp21731 := PrimCons(A, Nil)

																					tmp21732 := Call(__e, PrimNS2Value(symshen_4bindv), V2463, tmp21731, V2953)

																					_ = tmp21732

																					tmp21733 := MakeNative(func(__e *ControlFlow) {
																						Result := __e.Get(1)
																						_ = Result
																						tmp21734 := Call(__e, PrimNS2Value(symshen_4unbindv), V2463, V2953)

																						_ = tmp21734

																						__e.Return(Result)
																						return

																					}, 1)

																					tmp21735 := MakeNative(func(__e *ControlFlow) {
																						NewHyp := __e.Get(1)
																						_ = NewHyp
																						tmp21736 := Call(__e, PrimNS2Value(symshen_4incinfs))

																						_ = tmp21736

																						tmp21737 := MakeNative(func(__e *ControlFlow) {
																							__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																							return
																						}, 0)

																						__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21737)
																						return

																					}, 1)

																					tmp21738 := PrimTail(V2459)

																					tmp21739 := Call(__e, tmp21735, tmp21738)

																					__e.TailApply(tmp21733, tmp21739)
																					return

																				}, 1)

																				tmp21740 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

																				__e.TailApply(tmp21730, tmp21740)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}

																	}, 1)

																	tmp21764 := PrimTail(V2461)

																	tmp21765 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21764, V2953)

																	__e.TailApply(tmp21727, tmp21765)
																	return

																} else {
																	tmp21726 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2462)

																	if True == tmp21726 {
																		tmp21683 := Call(__e, PrimNS2Value(symshen_4bindv), V2462, sym_h, V2953)

																		_ = tmp21683

																		tmp21684 := MakeNative(func(__e *ControlFlow) {
																			Result := __e.Get(1)
																			_ = Result
																			tmp21685 := Call(__e, PrimNS2Value(symshen_4unbindv), V2462, V2953)

																			_ = tmp21685

																			__e.Return(Result)
																			return

																		}, 1)

																		tmp21686 := MakeNative(func(__e *ControlFlow) {
																			V2465 := __e.Get(1)
																			_ = V2465
																			tmp21722 := PrimIsPair(V2465)

																			if True == tmp21722 {
																				tmp21701 := MakeNative(func(__e *ControlFlow) {
																					A := __e.Get(1)
																					_ = A
																					tmp21702 := MakeNative(func(__e *ControlFlow) {
																						V2466 := __e.Get(1)
																						_ = V2466
																						tmp21718 := PrimEqual(Nil, V2466)

																						if True == tmp21718 {
																							tmp21714 := MakeNative(func(__e *ControlFlow) {
																								NewHyp := __e.Get(1)
																								_ = NewHyp
																								tmp21715 := Call(__e, PrimNS2Value(symshen_4incinfs))

																								_ = tmp21715

																								tmp21716 := MakeNative(func(__e *ControlFlow) {
																									__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																									return
																								}, 0)

																								__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21716)
																								return

																							}, 1)

																							tmp21717 := PrimTail(V2459)

																							__e.TailApply(tmp21714, tmp21717)
																							return

																						} else {
																							tmp21713 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2466)

																							if True == tmp21713 {
																								tmp21705 := Call(__e, PrimNS2Value(symshen_4bindv), V2466, Nil, V2953)

																								_ = tmp21705

																								tmp21706 := MakeNative(func(__e *ControlFlow) {
																									Result := __e.Get(1)
																									_ = Result
																									tmp21707 := Call(__e, PrimNS2Value(symshen_4unbindv), V2466, V2953)

																									_ = tmp21707

																									__e.Return(Result)
																									return

																								}, 1)

																								tmp21708 := MakeNative(func(__e *ControlFlow) {
																									NewHyp := __e.Get(1)
																									_ = NewHyp
																									tmp21709 := Call(__e, PrimNS2Value(symshen_4incinfs))

																									_ = tmp21709

																									tmp21710 := MakeNative(func(__e *ControlFlow) {
																										__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																										return
																									}, 0)

																									__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21710)
																									return

																								}, 1)

																								tmp21711 := PrimTail(V2459)

																								tmp21712 := Call(__e, tmp21708, tmp21711)

																								__e.TailApply(tmp21706, tmp21712)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}

																					}, 1)

																					tmp21719 := PrimTail(V2465)

																					tmp21720 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21719, V2953)

																					__e.TailApply(tmp21702, tmp21720)
																					return

																				}, 1)

																				tmp21721 := PrimHead(V2465)

																				__e.TailApply(tmp21701, tmp21721)
																				return

																			} else {
																				tmp21700 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2465)

																				if True == tmp21700 {
																					tmp21689 := MakeNative(func(__e *ControlFlow) {
																						A := __e.Get(1)
																						_ = A
																						tmp21690 := PrimCons(A, Nil)

																						tmp21691 := Call(__e, PrimNS2Value(symshen_4bindv), V2465, tmp21690, V2953)

																						_ = tmp21691

																						tmp21692 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp21693 := Call(__e, PrimNS2Value(symshen_4unbindv), V2465, V2953)

																							_ = tmp21693

																							__e.Return(Result)
																							return

																						}, 1)

																						tmp21694 := MakeNative(func(__e *ControlFlow) {
																							NewHyp := __e.Get(1)
																							_ = NewHyp
																							tmp21695 := Call(__e, PrimNS2Value(symshen_4incinfs))

																							_ = tmp21695

																							tmp21696 := MakeNative(func(__e *ControlFlow) {
																								__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																								return
																							}, 0)

																							__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21696)
																							return

																						}, 1)

																						tmp21697 := PrimTail(V2459)

																						tmp21698 := Call(__e, tmp21694, tmp21697)

																						__e.TailApply(tmp21692, tmp21698)
																						return

																					}, 1)

																					tmp21699 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

																					__e.TailApply(tmp21689, tmp21699)
																					return

																				} else {
																					__e.Return(False)
																					return
																				}

																			}

																		}, 1)

																		tmp21723 := PrimTail(V2461)

																		tmp21724 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21723, V2953)

																		tmp21725 := Call(__e, tmp21686, tmp21724)

																		__e.TailApply(tmp21684, tmp21725)
																		return

																	} else {
																		__e.Return(False)
																		return
																	}

																}

															}, 1)

															tmp21767 := PrimHead(V2461)

															tmp21768 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21767, V2953)

															__e.TailApply(tmp21680, tmp21768)
															return

														} else {
															tmp21679 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2461)

															if True == tmp21679 {
																tmp21667 := MakeNative(func(__e *ControlFlow) {
																	A := __e.Get(1)
																	_ = A
																	tmp21668 := PrimCons(A, Nil)

																	tmp21669 := PrimCons(sym_h, tmp21668)

																	tmp21670 := Call(__e, PrimNS2Value(symshen_4bindv), V2461, tmp21669, V2953)

																	_ = tmp21670

																	tmp21671 := MakeNative(func(__e *ControlFlow) {
																		Result := __e.Get(1)
																		_ = Result
																		tmp21672 := Call(__e, PrimNS2Value(symshen_4unbindv), V2461, V2953)

																		_ = tmp21672

																		__e.Return(Result)
																		return

																	}, 1)

																	tmp21673 := MakeNative(func(__e *ControlFlow) {
																		NewHyp := __e.Get(1)
																		_ = NewHyp
																		tmp21674 := Call(__e, PrimNS2Value(symshen_4incinfs))

																		_ = tmp21674

																		tmp21675 := MakeNative(func(__e *ControlFlow) {
																			__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																			return
																		}, 0)

																		__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21675)
																		return

																	}, 1)

																	tmp21676 := PrimTail(V2459)

																	tmp21677 := Call(__e, tmp21673, tmp21676)

																	__e.TailApply(tmp21671, tmp21677)
																	return

																}, 1)

																tmp21678 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

																__e.TailApply(tmp21667, tmp21678)
																return

															} else {
																__e.Return(False)
																return
															}

														}

													}, 1)

													tmp21770 := PrimTail(V2460)

													tmp21771 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21770, V2953)

													__e.TailApply(tmp21664, tmp21771)
													return

												}, 1)

												tmp21772 := PrimHead(V2460)

												__e.TailApply(tmp21663, tmp21772)
												return

											} else {
												tmp21662 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2460)

												if True == tmp21662 {
													tmp21647 := MakeNative(func(__e *ControlFlow) {
														V := __e.Get(1)
														_ = V
														tmp21648 := MakeNative(func(__e *ControlFlow) {
															A := __e.Get(1)
															_ = A
															tmp21649 := PrimCons(A, Nil)

															tmp21650 := PrimCons(sym_h, tmp21649)

															tmp21651 := PrimCons(V, tmp21650)

															tmp21652 := Call(__e, PrimNS2Value(symshen_4bindv), V2460, tmp21651, V2953)

															_ = tmp21652

															tmp21653 := MakeNative(func(__e *ControlFlow) {
																Result := __e.Get(1)
																_ = Result
																tmp21654 := Call(__e, PrimNS2Value(symshen_4unbindv), V2460, V2953)

																_ = tmp21654

																__e.Return(Result)
																return

															}, 1)

															tmp21655 := MakeNative(func(__e *ControlFlow) {
																NewHyp := __e.Get(1)
																_ = NewHyp
																tmp21656 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp21656

																tmp21657 := MakeNative(func(__e *ControlFlow) {
																	__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
																	return
																}, 0)

																__e.TailApply(PrimNS2Value(symunify_b), V, V2453, V2953, tmp21657)
																return

															}, 1)

															tmp21658 := PrimTail(V2459)

															tmp21659 := Call(__e, tmp21655, tmp21658)

															__e.TailApply(tmp21653, tmp21659)
															return

														}, 1)

														tmp21660 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

														__e.TailApply(tmp21648, tmp21660)
														return

													}, 1)

													tmp21661 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

													__e.TailApply(tmp21647, tmp21661)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										}, 1)

										tmp21774 := PrimHead(V2459)

										tmp21775 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21774, V2953)

										__e.TailApply(tmp21644, tmp21775)
										return

									} else {
										tmp21643 := Call(__e, PrimNS2Value(symshen_4pvar_2), V2459)

										if True == tmp21643 {
											tmp21627 := MakeNative(func(__e *ControlFlow) {
												V := __e.Get(1)
												_ = V
												tmp21628 := MakeNative(func(__e *ControlFlow) {
													A := __e.Get(1)
													_ = A
													tmp21629 := MakeNative(func(__e *ControlFlow) {
														NewHyp := __e.Get(1)
														_ = NewHyp
														tmp21630 := PrimCons(A, Nil)

														tmp21631 := PrimCons(sym_h, tmp21630)

														tmp21632 := PrimCons(V, tmp21631)

														tmp21633 := PrimCons(tmp21632, NewHyp)

														tmp21634 := Call(__e, PrimNS2Value(symshen_4bindv), V2459, tmp21633, V2953)

														_ = tmp21634

														tmp21635 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp21636 := Call(__e, PrimNS2Value(symshen_4unbindv), V2459, V2953)

															_ = tmp21636

															__e.Return(Result)
															return

														}, 1)

														tmp21637 := Call(__e, PrimNS2Value(symshen_4incinfs))

														_ = tmp21637

														tmp21638 := MakeNative(func(__e *ControlFlow) {
															__e.TailApply(PrimNS2Value(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)
															return
														}, 0)

														tmp21639 := Call(__e, PrimNS2Value(symunify_b), V, V2453, V2953, tmp21638)

														__e.TailApply(tmp21635, tmp21639)
														return

													}, 1)

													tmp21640 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

													__e.TailApply(tmp21629, tmp21640)
													return

												}, 1)

												tmp21641 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

												__e.TailApply(tmp21628, tmp21641)
												return

											}, 1)

											tmp21642 := Call(__e, PrimNS2Value(symshen_4newpv), V2953)

											__e.TailApply(tmp21627, tmp21642)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								}, 1)

								tmp21777 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2952, V2953)

								__e.TailApply(tmp21624, tmp21777)
								return

							}, 1)

							tmp21778 := PrimTail(V2458)

							__e.TailApply(tmp21623, tmp21778)
							return

						}, 1)

						tmp21779 := PrimHead(V2458)

						__e.TailApply(tmp21622, tmp21779)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21781 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2950, V2953)

				__e.TailApply(tmp21620, tmp21781)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp21783 := MakeNative(func(__e *ControlFlow) {
			V2457 := __e.Get(1)
			_ = V2457
			tmp21786 := PrimEqual(Nil, V2457)

			if True == tmp21786 {
				tmp21785 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp21785

				__e.TailApply(PrimNS2Value(symunify_b), V2952, V2951, V2953, V2954)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21787 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2950, V2953)

		tmp21788 := Call(__e, tmp21783, tmp21787)

		__e.TailApply(tmp21618, tmp21788)
		return

	}, 5)

	tmp21789 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newhyps, tmp21617)

	_ = tmp21789

	tmp21790 := MakeNative(func(__e *ControlFlow) {
		V2960 := __e.Get(1)
		_ = V2960
		V2961 := __e.Get(2)
		_ = V2961
		V2962 := __e.Get(3)
		_ = V2962
		tmp21828 := PrimEqual(Nil, V2960)

		if True == tmp21828 {
			__e.Return(V2962)
			return
		} else {
			tmp21827 := PrimIsPair(V2960)

			var ifres21804 Obj

			if True == tmp21827 {
				tmp21826 := PrimIsPair(V2961)

				var ifres21806 Obj

				if True == tmp21826 {
					tmp21824 := PrimTail(V2961)

					tmp21825 := PrimIsPair(tmp21824)

					var ifres21808 Obj

					if True == tmp21825 {
						tmp21821 := PrimTail(V2961)

						tmp21822 := PrimHead(tmp21821)

						tmp21823 := PrimEqual(sym_1_1_6, tmp21822)

						var ifres21810 Obj

						if True == tmp21823 {
							tmp21818 := PrimTail(V2961)

							tmp21819 := PrimTail(tmp21818)

							tmp21820 := PrimIsPair(tmp21819)

							var ifres21812 Obj

							if True == tmp21820 {
								tmp21814 := PrimTail(V2961)

								tmp21815 := PrimTail(tmp21814)

								tmp21816 := PrimTail(tmp21815)

								tmp21817 := PrimEqual(Nil, tmp21816)

								var ifres21813 Obj

								if True == tmp21817 {
									ifres21813 = True

								} else {
									ifres21813 = False

								}

								ifres21812 = ifres21813

							} else {
								ifres21812 = False

							}

							var ifres21811 Obj

							if True == ifres21812 {
								ifres21811 = True

							} else {
								ifres21811 = False

							}

							ifres21810 = ifres21811

						} else {
							ifres21810 = False

						}

						var ifres21809 Obj

						if True == ifres21810 {
							ifres21809 = True

						} else {
							ifres21809 = False

						}

						ifres21808 = ifres21809

					} else {
						ifres21808 = False

					}

					var ifres21807 Obj

					if True == ifres21808 {
						ifres21807 = True

					} else {
						ifres21807 = False

					}

					ifres21806 = ifres21807

				} else {
					ifres21806 = False

				}

				var ifres21805 Obj

				if True == ifres21806 {
					ifres21805 = True

				} else {
					ifres21805 = False

				}

				ifres21804 = ifres21805

			} else {
				ifres21804 = False

			}

			if True == ifres21804 {
				tmp21793 := PrimHead(V2960)

				tmp21794 := Call(__e, PrimNS2Value(symshen_4curry), tmp21793)

				tmp21795 := PrimHead(V2961)

				tmp21796 := PrimCons(tmp21795, Nil)

				tmp21797 := PrimCons(sym_h, tmp21796)

				tmp21798 := PrimCons(tmp21794, tmp21797)

				tmp21799 := PrimTail(V2960)

				tmp21800 := PrimTail(V2961)

				tmp21801 := PrimTail(tmp21800)

				tmp21802 := PrimHead(tmp21801)

				tmp21803 := Call(__e, PrimNS2Value(symshen_4patthyps), tmp21799, tmp21802, V2962)

				__e.TailApply(PrimNS2Value(symadjoin), tmp21798, tmp21803)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4patthyps)
				return
			}

		}

	}, 3)

	tmp21829 := Call(__e, PrimNS1Value(symns2_1set), symshen_4patthyps, tmp21790)

	_ = tmp21829

	tmp21830 := MakeNative(func(__e *ControlFlow) {
		V2969 := __e.Get(1)
		_ = V2969
		V2970 := __e.Get(2)
		_ = V2970
		tmp21880 := PrimEqual(Nil, V2969)

		var ifres21864 Obj

		if True == tmp21880 {
			tmp21879 := PrimIsPair(V2970)

			var ifres21866 Obj

			if True == tmp21879 {
				tmp21877 := PrimHead(V2970)

				tmp21878 := PrimEqual(sym_1_1_6, tmp21877)

				var ifres21868 Obj

				if True == tmp21878 {
					tmp21875 := PrimTail(V2970)

					tmp21876 := PrimIsPair(tmp21875)

					var ifres21870 Obj

					if True == tmp21876 {
						tmp21872 := PrimTail(V2970)

						tmp21873 := PrimTail(tmp21872)

						tmp21874 := PrimEqual(Nil, tmp21873)

						var ifres21871 Obj

						if True == tmp21874 {
							ifres21871 = True

						} else {
							ifres21871 = False

						}

						ifres21870 = ifres21871

					} else {
						ifres21870 = False

					}

					var ifres21869 Obj

					if True == ifres21870 {
						ifres21869 = True

					} else {
						ifres21869 = False

					}

					ifres21868 = ifres21869

				} else {
					ifres21868 = False

				}

				var ifres21867 Obj

				if True == ifres21868 {
					ifres21867 = True

				} else {
					ifres21867 = False

				}

				ifres21866 = ifres21867

			} else {
				ifres21866 = False

			}

			var ifres21865 Obj

			if True == ifres21866 {
				ifres21865 = True

			} else {
				ifres21865 = False

			}

			ifres21864 = ifres21865

		} else {
			ifres21864 = False

		}

		if True == ifres21864 {
			tmp21863 := PrimTail(V2970)

			__e.Return(PrimHead(tmp21863))
			return

		} else {
			tmp21862 := PrimEqual(Nil, V2969)

			if True == tmp21862 {
				__e.Return(V2970)
				return
			} else {
				tmp21861 := PrimIsPair(V2969)

				var ifres21838 Obj

				if True == tmp21861 {
					tmp21860 := PrimIsPair(V2970)

					var ifres21840 Obj

					if True == tmp21860 {
						tmp21858 := PrimTail(V2970)

						tmp21859 := PrimIsPair(tmp21858)

						var ifres21842 Obj

						if True == tmp21859 {
							tmp21855 := PrimTail(V2970)

							tmp21856 := PrimHead(tmp21855)

							tmp21857 := PrimEqual(sym_1_1_6, tmp21856)

							var ifres21844 Obj

							if True == tmp21857 {
								tmp21852 := PrimTail(V2970)

								tmp21853 := PrimTail(tmp21852)

								tmp21854 := PrimIsPair(tmp21853)

								var ifres21846 Obj

								if True == tmp21854 {
									tmp21848 := PrimTail(V2970)

									tmp21849 := PrimTail(tmp21848)

									tmp21850 := PrimTail(tmp21849)

									tmp21851 := PrimEqual(Nil, tmp21850)

									var ifres21847 Obj

									if True == tmp21851 {
										ifres21847 = True

									} else {
										ifres21847 = False

									}

									ifres21846 = ifres21847

								} else {
									ifres21846 = False

								}

								var ifres21845 Obj

								if True == ifres21846 {
									ifres21845 = True

								} else {
									ifres21845 = False

								}

								ifres21844 = ifres21845

							} else {
								ifres21844 = False

							}

							var ifres21843 Obj

							if True == ifres21844 {
								ifres21843 = True

							} else {
								ifres21843 = False

							}

							ifres21842 = ifres21843

						} else {
							ifres21842 = False

						}

						var ifres21841 Obj

						if True == ifres21842 {
							ifres21841 = True

						} else {
							ifres21841 = False

						}

						ifres21840 = ifres21841

					} else {
						ifres21840 = False

					}

					var ifres21839 Obj

					if True == ifres21840 {
						ifres21839 = True

					} else {
						ifres21839 = False

					}

					ifres21838 = ifres21839

				} else {
					ifres21838 = False

				}

				if True == ifres21838 {
					tmp21834 := PrimTail(V2969)

					tmp21835 := PrimTail(V2970)

					tmp21836 := PrimTail(tmp21835)

					tmp21837 := PrimHead(tmp21836)

					__e.TailApply(PrimNS2Value(symshen_4result_1type), tmp21834, tmp21837)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4result_1type)
					return
				}

			}

		}

	}, 2)

	tmp21881 := Call(__e, PrimNS1Value(symns2_1set), symshen_4result_1type, tmp21830)

	_ = tmp21881

	tmp21882 := MakeNative(func(__e *ControlFlow) {
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
		tmp21883 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp21927 := PrimEqual(Case, False)

			if True == tmp21927 {
				tmp21885 := MakeNative(func(__e *ControlFlow) {
					V2446 := __e.Get(1)
					_ = V2446
					tmp21925 := PrimIsPair(V2446)

					if True == tmp21925 {
						tmp21887 := MakeNative(func(__e *ControlFlow) {
							Pattern := __e.Get(1)
							_ = Pattern
							tmp21888 := MakeNative(func(__e *ControlFlow) {
								Patterns := __e.Get(1)
								_ = Patterns
								tmp21889 := MakeNative(func(__e *ControlFlow) {
									V2447 := __e.Get(1)
									_ = V2447
									tmp21921 := PrimIsPair(V2447)

									if True == tmp21921 {
										tmp21891 := MakeNative(func(__e *ControlFlow) {
											A := __e.Get(1)
											_ = A
											tmp21892 := MakeNative(func(__e *ControlFlow) {
												V2448 := __e.Get(1)
												_ = V2448
												tmp21917 := PrimIsPair(V2448)

												if True == tmp21917 {
													tmp21894 := MakeNative(func(__e *ControlFlow) {
														V2449 := __e.Get(1)
														_ = V2449
														tmp21914 := PrimEqual(sym_1_1_6, V2449)

														if True == tmp21914 {
															tmp21896 := MakeNative(func(__e *ControlFlow) {
																V2450 := __e.Get(1)
																_ = V2450
																tmp21911 := PrimIsPair(V2450)

																if True == tmp21911 {
																	tmp21898 := MakeNative(func(__e *ControlFlow) {
																		B := __e.Get(1)
																		_ = B
																		tmp21899 := MakeNative(func(__e *ControlFlow) {
																			V2451 := __e.Get(1)
																			_ = V2451
																			tmp21907 := PrimEqual(Nil, V2451)

																			if True == tmp21907 {
																				tmp21901 := Call(__e, PrimNS2Value(symshen_4incinfs))

																				_ = tmp21901

																				tmp21902 := Call(__e, PrimNS2Value(symshen_4curry), Pattern)

																				tmp21903 := PrimCons(A, Nil)

																				tmp21904 := PrimCons(sym_h, tmp21903)

																				tmp21905 := PrimCons(tmp21902, tmp21904)

																				tmp21906 := MakeNative(func(__e *ControlFlow) {
																					__e.TailApply(PrimNS2Value(symshen_4t_d_1patterns), Patterns, B, V2978, V2979, V2980)
																					return
																				}, 0)

																				__e.TailApply(PrimNS2Value(symshen_4t_d), tmp21905, V2978, V2979, tmp21906)
																				return

																			} else {
																				__e.Return(False)
																				return
																			}

																		}, 1)

																		tmp21908 := PrimTail(V2450)

																		tmp21909 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21908, V2979)

																		__e.TailApply(tmp21899, tmp21909)
																		return

																	}, 1)

																	tmp21910 := PrimHead(V2450)

																	__e.TailApply(tmp21898, tmp21910)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp21912 := PrimTail(V2448)

															tmp21913 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21912, V2979)

															__e.TailApply(tmp21896, tmp21913)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp21915 := PrimHead(V2448)

													tmp21916 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21915, V2979)

													__e.TailApply(tmp21894, tmp21916)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp21918 := PrimTail(V2447)

											tmp21919 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21918, V2979)

											__e.TailApply(tmp21892, tmp21919)
											return

										}, 1)

										tmp21920 := PrimHead(V2447)

										__e.TailApply(tmp21891, tmp21920)
										return

									} else {
										__e.Return(False)
										return
									}

								}, 1)

								tmp21922 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2977, V2979)

								__e.TailApply(tmp21889, tmp21922)
								return

							}, 1)

							tmp21923 := PrimTail(V2446)

							__e.TailApply(tmp21888, tmp21923)
							return

						}, 1)

						tmp21924 := PrimHead(V2446)

						__e.TailApply(tmp21887, tmp21924)
						return

					} else {
						__e.Return(False)
						return
					}

				}, 1)

				tmp21926 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2976, V2979)

				__e.TailApply(tmp21885, tmp21926)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp21928 := MakeNative(func(__e *ControlFlow) {
			V2445 := __e.Get(1)
			_ = V2445
			tmp21931 := PrimEqual(Nil, V2445)

			if True == tmp21931 {
				tmp21930 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp21930

				__e.TailApply(PrimNS2Value(symthaw), V2980)
				return

			} else {
				__e.Return(False)
				return
			}

		}, 1)

		tmp21932 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2976, V2979)

		tmp21933 := Call(__e, tmp21928, tmp21932)

		__e.TailApply(tmp21883, tmp21933)
		return

	}, 5)

	tmp21934 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1patterns, tmp21882)

	_ = tmp21934

	tmp21935 := MakeNative(func(__e *ControlFlow) {
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
		tmp21936 := MakeNative(func(__e *ControlFlow) {
			Throwcontrol := __e.Get(1)
			_ = Throwcontrol
			tmp21937 := MakeNative(func(__e *ControlFlow) {
				Case := __e.Get(1)
				_ = Case
				tmp22051 := PrimEqual(Case, False)

				if True == tmp22051 {
					tmp21939 := MakeNative(func(__e *ControlFlow) {
						Case := __e.Get(1)
						_ = Case
						tmp21982 := PrimEqual(Case, False)

						if True == tmp21982 {
							tmp21941 := MakeNative(func(__e *ControlFlow) {
								Case := __e.Get(1)
								_ = Case
								tmp21947 := PrimEqual(Case, False)

								if True == tmp21947 {
									tmp21943 := Call(__e, PrimNS2Value(symshen_4incinfs))

									_ = tmp21943

									tmp21944 := PrimCons(V2987, Nil)

									tmp21945 := PrimCons(sym_h, tmp21944)

									tmp21946 := PrimCons(V2986, tmp21945)

									__e.TailApply(PrimNS2Value(symshen_4t_d), tmp21946, V2988, V2989, V2990)
									return

								} else {
									__e.Return(Case)
									return
								}

							}, 1)

							tmp21948 := MakeNative(func(__e *ControlFlow) {
								V2438 := __e.Get(1)
								_ = V2438
								tmp21979 := PrimIsPair(V2438)

								if True == tmp21979 {
									tmp21950 := MakeNative(func(__e *ControlFlow) {
										V2439 := __e.Get(1)
										_ = V2439
										tmp21976 := PrimEqual(symshen_4choicepoint_b, V2439)

										if True == tmp21976 {
											tmp21952 := MakeNative(func(__e *ControlFlow) {
												V2440 := __e.Get(1)
												_ = V2440
												tmp21973 := PrimIsPair(V2440)

												if True == tmp21973 {
													tmp21954 := MakeNative(func(__e *ControlFlow) {
														Action := __e.Get(1)
														_ = Action
														tmp21955 := MakeNative(func(__e *ControlFlow) {
															V2441 := __e.Get(1)
															_ = V2441
															tmp21969 := PrimEqual(Nil, V2441)

															if True == tmp21969 {
																tmp21957 := Call(__e, PrimNS2Value(symshen_4incinfs))

																_ = tmp21957

																tmp21958 := MakeNative(func(__e *ControlFlow) {
																	tmp21959 := PrimCons(Action, Nil)

																	tmp21960 := PrimCons(sym_a, tmp21959)

																	tmp21961 := PrimCons(symfail, Nil)

																	tmp21962 := PrimCons(tmp21961, Nil)

																	tmp21963 := PrimCons(tmp21960, tmp21962)

																	tmp21964 := PrimCons(tmp21963, Nil)

																	tmp21965 := PrimCons(symnot, tmp21964)

																	tmp21966 := PrimCons(Action, Nil)

																	tmp21967 := PrimCons(tmp21965, tmp21966)

																	tmp21968 := PrimCons(symwhere, tmp21967)

																	__e.TailApply(PrimNS2Value(symshen_4t_d_1action), tmp21968, V2987, V2988, V2989, V2990)
																	return

																}, 0)

																__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2989, tmp21958)
																return

															} else {
																__e.Return(False)
																return
															}

														}, 1)

														tmp21970 := PrimTail(V2440)

														tmp21971 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21970, V2989)

														__e.TailApply(tmp21955, tmp21971)
														return

													}, 1)

													tmp21972 := PrimHead(V2440)

													__e.TailApply(tmp21954, tmp21972)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp21974 := PrimTail(V2438)

											tmp21975 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21974, V2989)

											__e.TailApply(tmp21952, tmp21975)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp21977 := PrimHead(V2438)

									tmp21978 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp21977, V2989)

									__e.TailApply(tmp21950, tmp21978)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp21980 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2986, V2989)

							tmp21981 := Call(__e, tmp21948, tmp21980)

							__e.TailApply(tmp21941, tmp21981)
							return

						} else {
							__e.Return(Case)
							return
						}

					}, 1)

					tmp21983 := MakeNative(func(__e *ControlFlow) {
						V2427 := __e.Get(1)
						_ = V2427
						tmp22048 := PrimIsPair(V2427)

						if True == tmp22048 {
							tmp21985 := MakeNative(func(__e *ControlFlow) {
								V2428 := __e.Get(1)
								_ = V2428
								tmp22045 := PrimEqual(symshen_4choicepoint_b, V2428)

								if True == tmp22045 {
									tmp21987 := MakeNative(func(__e *ControlFlow) {
										V2429 := __e.Get(1)
										_ = V2429
										tmp22042 := PrimIsPair(V2429)

										if True == tmp22042 {
											tmp21989 := MakeNative(func(__e *ControlFlow) {
												V2430 := __e.Get(1)
												_ = V2430
												tmp22039 := PrimIsPair(V2430)

												if True == tmp22039 {
													tmp21991 := MakeNative(func(__e *ControlFlow) {
														V2431 := __e.Get(1)
														_ = V2431
														tmp22036 := PrimIsPair(V2431)

														if True == tmp22036 {
															tmp21993 := MakeNative(func(__e *ControlFlow) {
																V2432 := __e.Get(1)
																_ = V2432
																tmp22033 := PrimEqual(symfail_1if, V2432)

																if True == tmp22033 {
																	tmp21995 := MakeNative(func(__e *ControlFlow) {
																		V2433 := __e.Get(1)
																		_ = V2433
																		tmp22030 := PrimIsPair(V2433)

																		if True == tmp22030 {
																			tmp21997 := MakeNative(func(__e *ControlFlow) {
																				F := __e.Get(1)
																				_ = F
																				tmp21998 := MakeNative(func(__e *ControlFlow) {
																					V2434 := __e.Get(1)
																					_ = V2434
																					tmp22026 := PrimEqual(Nil, V2434)

																					if True == tmp22026 {
																						tmp22000 := MakeNative(func(__e *ControlFlow) {
																							V2435 := __e.Get(1)
																							_ = V2435
																							tmp22023 := PrimIsPair(V2435)

																							if True == tmp22023 {
																								tmp22002 := MakeNative(func(__e *ControlFlow) {
																									Action := __e.Get(1)
																									_ = Action
																									tmp22003 := MakeNative(func(__e *ControlFlow) {
																										V2436 := __e.Get(1)
																										_ = V2436
																										tmp22019 := PrimEqual(Nil, V2436)

																										if True == tmp22019 {
																											tmp22005 := MakeNative(func(__e *ControlFlow) {
																												V2437 := __e.Get(1)
																												_ = V2437
																												tmp22016 := PrimEqual(Nil, V2437)

																												if True == tmp22016 {
																													tmp22007 := Call(__e, PrimNS2Value(symshen_4incinfs))

																													_ = tmp22007

																													tmp22008 := MakeNative(func(__e *ControlFlow) {
																														tmp22009 := PrimCons(Action, Nil)

																														tmp22010 := PrimCons(F, tmp22009)

																														tmp22011 := PrimCons(tmp22010, Nil)

																														tmp22012 := PrimCons(symnot, tmp22011)

																														tmp22013 := PrimCons(Action, Nil)

																														tmp22014 := PrimCons(tmp22012, tmp22013)

																														tmp22015 := PrimCons(symwhere, tmp22014)

																														__e.TailApply(PrimNS2Value(symshen_4t_d_1action), tmp22015, V2987, V2988, V2989, V2990)
																														return

																													}, 0)

																													__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2989, tmp22008)
																													return

																												} else {
																													__e.Return(False)
																													return
																												}

																											}, 1)

																											tmp22017 := PrimTail(V2429)

																											tmp22018 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22017, V2989)

																											__e.TailApply(tmp22005, tmp22018)
																											return

																										} else {
																											__e.Return(False)
																											return
																										}

																									}, 1)

																									tmp22020 := PrimTail(V2435)

																									tmp22021 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22020, V2989)

																									__e.TailApply(tmp22003, tmp22021)
																									return

																								}, 1)

																								tmp22022 := PrimHead(V2435)

																								__e.TailApply(tmp22002, tmp22022)
																								return

																							} else {
																								__e.Return(False)
																								return
																							}

																						}, 1)

																						tmp22024 := PrimTail(V2430)

																						tmp22025 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22024, V2989)

																						__e.TailApply(tmp22000, tmp22025)
																						return

																					} else {
																						__e.Return(False)
																						return
																					}

																				}, 1)

																				tmp22027 := PrimTail(V2433)

																				tmp22028 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22027, V2989)

																				__e.TailApply(tmp21998, tmp22028)
																				return

																			}, 1)

																			tmp22029 := PrimHead(V2433)

																			__e.TailApply(tmp21997, tmp22029)
																			return

																		} else {
																			__e.Return(False)
																			return
																		}

																	}, 1)

																	tmp22031 := PrimTail(V2431)

																	tmp22032 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22031, V2989)

																	__e.TailApply(tmp21995, tmp22032)
																	return

																} else {
																	__e.Return(False)
																	return
																}

															}, 1)

															tmp22034 := PrimHead(V2431)

															tmp22035 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22034, V2989)

															__e.TailApply(tmp21993, tmp22035)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp22037 := PrimHead(V2430)

													tmp22038 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22037, V2989)

													__e.TailApply(tmp21991, tmp22038)
													return

												} else {
													__e.Return(False)
													return
												}

											}, 1)

											tmp22040 := PrimHead(V2429)

											tmp22041 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22040, V2989)

											__e.TailApply(tmp21989, tmp22041)
											return

										} else {
											__e.Return(False)
											return
										}

									}, 1)

									tmp22043 := PrimTail(V2427)

									tmp22044 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22043, V2989)

									__e.TailApply(tmp21987, tmp22044)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp22046 := PrimHead(V2427)

							tmp22047 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22046, V2989)

							__e.TailApply(tmp21985, tmp22047)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp22049 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2986, V2989)

					tmp22050 := Call(__e, tmp21983, tmp22049)

					__e.TailApply(tmp21939, tmp22050)
					return

				} else {
					__e.Return(Case)
					return
				}

			}, 1)

			tmp22052 := MakeNative(func(__e *ControlFlow) {
				V2422 := __e.Get(1)
				_ = V2422
				tmp22089 := PrimIsPair(V2422)

				if True == tmp22089 {
					tmp22054 := MakeNative(func(__e *ControlFlow) {
						V2423 := __e.Get(1)
						_ = V2423
						tmp22086 := PrimEqual(symwhere, V2423)

						if True == tmp22086 {
							tmp22056 := MakeNative(func(__e *ControlFlow) {
								V2424 := __e.Get(1)
								_ = V2424
								tmp22083 := PrimIsPair(V2424)

								if True == tmp22083 {
									tmp22058 := MakeNative(func(__e *ControlFlow) {
										P := __e.Get(1)
										_ = P
										tmp22059 := MakeNative(func(__e *ControlFlow) {
											V2425 := __e.Get(1)
											_ = V2425
											tmp22079 := PrimIsPair(V2425)

											if True == tmp22079 {
												tmp22061 := MakeNative(func(__e *ControlFlow) {
													Action := __e.Get(1)
													_ = Action
													tmp22062 := MakeNative(func(__e *ControlFlow) {
														V2426 := __e.Get(1)
														_ = V2426
														tmp22075 := PrimEqual(Nil, V2426)

														if True == tmp22075 {
															tmp22064 := Call(__e, PrimNS2Value(symshen_4incinfs))

															_ = tmp22064

															tmp22065 := MakeNative(func(__e *ControlFlow) {
																tmp22066 := PrimCons(symboolean, Nil)

																tmp22067 := PrimCons(sym_h, tmp22066)

																tmp22068 := PrimCons(P, tmp22067)

																tmp22069 := MakeNative(func(__e *ControlFlow) {
																	tmp22070 := MakeNative(func(__e *ControlFlow) {
																		tmp22071 := PrimCons(symverified, Nil)

																		tmp22072 := PrimCons(sym_h, tmp22071)

																		tmp22073 := PrimCons(P, tmp22072)

																		tmp22074 := PrimCons(tmp22073, V2988)

																		__e.TailApply(PrimNS2Value(symshen_4t_d_1action), Action, V2987, tmp22074, V2989, V2990)
																		return

																	}, 0)

																	__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2989, tmp22070)
																	return

																}, 0)

																__e.TailApply(PrimNS2Value(symshen_4t_d), tmp22068, V2988, V2989, tmp22069)
																return

															}, 0)

															__e.TailApply(PrimNS2Value(symcut), Throwcontrol, V2989, tmp22065)
															return

														} else {
															__e.Return(False)
															return
														}

													}, 1)

													tmp22076 := PrimTail(V2425)

													tmp22077 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22076, V2989)

													__e.TailApply(tmp22062, tmp22077)
													return

												}, 1)

												tmp22078 := PrimHead(V2425)

												__e.TailApply(tmp22061, tmp22078)
												return

											} else {
												__e.Return(False)
												return
											}

										}, 1)

										tmp22080 := PrimTail(V2424)

										tmp22081 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22080, V2989)

										__e.TailApply(tmp22059, tmp22081)
										return

									}, 1)

									tmp22082 := PrimHead(V2424)

									__e.TailApply(tmp22058, tmp22082)
									return

								} else {
									__e.Return(False)
									return
								}

							}, 1)

							tmp22084 := PrimTail(V2422)

							tmp22085 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22084, V2989)

							__e.TailApply(tmp22056, tmp22085)
							return

						} else {
							__e.Return(False)
							return
						}

					}, 1)

					tmp22087 := PrimHead(V2422)

					tmp22088 := Call(__e, PrimNS2Value(symshen_4lazyderef), tmp22087, V2989)

					__e.TailApply(tmp22054, tmp22088)
					return

				} else {
					__e.Return(False)
					return
				}

			}, 1)

			tmp22090 := Call(__e, PrimNS2Value(symshen_4lazyderef), V2986, V2989)

			tmp22091 := Call(__e, tmp22052, tmp22090)

			tmp22092 := Call(__e, tmp21937, tmp22091)

			__e.TailApply(PrimNS2Value(symshen_4cutpoint), Throwcontrol, tmp22092)
			return

		}, 1)

		tmp22093 := Call(__e, PrimNS2Value(symshen_4catchpoint))

		__e.TailApply(tmp21936, tmp22093)
		return

	}, 5)

	tmp22094 := Call(__e, PrimNS1Value(symns2_1set), symshen_4t_d_1action, tmp21935)

	_ = tmp22094

	tmp22095 := MakeNative(func(__e *ControlFlow) {
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
		tmp22096 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp22097 := MakeNative(func(__e *ControlFlow) {
				A := __e.Get(1)
				_ = A
				tmp22098 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22098

				tmp22099 := Call(__e, PrimNS2Value(symgensym), symshen_4a)

				tmp22100 := MakeNative(func(__e *ControlFlow) {
					tmp22101 := Call(__e, PrimNS2Value(symshen_4lazyderef), A, V2999)

					tmp22102 := PrimNS3Set(tmp22101, Nil)

					tmp22103 := MakeNative(func(__e *ControlFlow) {
						__e.TailApply(PrimNS2Value(symshen_4findallhelp), V2996, V2997, V2998, A, V2999, V3000)
						return
					}, 0)

					__e.TailApply(PrimNS2Value(symbind), B, tmp22102, V2999, tmp22103)
					return

				}, 0)

				__e.TailApply(PrimNS2Value(symbind), A, tmp22099, V2999, tmp22100)
				return

			}, 1)

			tmp22104 := Call(__e, PrimNS2Value(symshen_4newpv), V2999)

			__e.TailApply(tmp22097, tmp22104)
			return

		}, 1)

		tmp22105 := Call(__e, PrimNS2Value(symshen_4newpv), V2999)

		__e.TailApply(tmp22096, tmp22105)
		return

	}, 5)

	tmp22106 := Call(__e, PrimNS1Value(symns2_1set), symfindall, tmp22095)

	_ = tmp22106

	tmp22107 := MakeNative(func(__e *ControlFlow) {
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
		tmp22108 := MakeNative(func(__e *ControlFlow) {
			Case := __e.Get(1)
			_ = Case
			tmp22113 := PrimEqual(Case, False)

			if True == tmp22113 {
				tmp22110 := Call(__e, PrimNS2Value(symshen_4incinfs))

				_ = tmp22110

				tmp22111 := Call(__e, PrimNS2Value(symshen_4lazyderef), V3010, V3011)

				tmp22112 := PrimNS3Value(tmp22111)

				__e.TailApply(PrimNS2Value(symbind), V3009, tmp22112, V3011, V3012)
				return

			} else {
				__e.Return(Case)
				return
			}

		}, 1)

		tmp22114 := Call(__e, PrimNS2Value(symshen_4incinfs))

		_ = tmp22114

		tmp22115 := MakeNative(func(__e *ControlFlow) {
			tmp22116 := MakeNative(func(__e *ControlFlow) {
				__e.TailApply(PrimNS2Value(symfwhen), False, V3011, V3012)
				return
			}, 0)

			__e.TailApply(PrimNS2Value(symshen_4remember), V3010, V3007, V3011, tmp22116)
			return

		}, 0)

		tmp22117 := Call(__e, PrimNS2Value(symcall), V3008, V3011, tmp22115)

		__e.TailApply(tmp22108, tmp22117)
		return

	}, 6)

	tmp22118 := Call(__e, PrimNS1Value(symns2_1set), symshen_4findallhelp, tmp22107)

	_ = tmp22118

	tmp22119 := MakeNative(func(__e *ControlFlow) {
		V3017 := __e.Get(1)
		_ = V3017
		V3018 := __e.Get(2)
		_ = V3018
		V3019 := __e.Get(3)
		_ = V3019
		V3020 := __e.Get(4)
		_ = V3020
		tmp22120 := MakeNative(func(__e *ControlFlow) {
			B := __e.Get(1)
			_ = B
			tmp22121 := Call(__e, PrimNS2Value(symshen_4incinfs))

			_ = tmp22121

			tmp22122 := Call(__e, PrimNS2Value(symshen_4deref), V3017, V3019)

			tmp22123 := Call(__e, PrimNS2Value(symshen_4deref), V3018, V3019)

			tmp22124 := Call(__e, PrimNS2Value(symshen_4deref), V3017, V3019)

			tmp22125 := PrimNS3Value(tmp22124)

			tmp22126 := PrimCons(tmp22123, tmp22125)

			tmp22127 := PrimNS3Set(tmp22122, tmp22126)

			__e.TailApply(PrimNS2Value(symbind), B, tmp22127, V3019, V3020)
			return

		}, 1)

		tmp22128 := Call(__e, PrimNS2Value(symshen_4newpv), V3019)

		__e.TailApply(tmp22120, tmp22128)
		return

	}, 4)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4remember, tmp22119)
	return

}, 0)
