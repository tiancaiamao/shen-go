package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp14858 := MakeNative(func(__e *ControlFlow) {
		V1722 := __e.Get(1)
		_ = V1722
		tmp14873 := PrimIsPair(V1722)

		var ifres14864 Obj

		if True == tmp14873 {
			tmp14871 := PrimTail(V1722)

			tmp14872 := PrimIsPair(tmp14871)

			var ifres14866 Obj

			if True == tmp14872 {
				tmp14868 := PrimTail(V1722)

				tmp14869 := PrimTail(tmp14868)

				tmp14870 := PrimEqual(Nil, tmp14869)

				var ifres14867 Obj

				if True == tmp14870 {
					ifres14867 = True

				} else {
					ifres14867 = False

				}

				ifres14866 = ifres14867

			} else {
				ifres14866 = False

			}

			var ifres14865 Obj

			if True == ifres14866 {
				ifres14865 = True

			} else {
				ifres14865 = False

			}

			ifres14864 = ifres14865

		} else {
			ifres14864 = False

		}

		if True == ifres14864 {
			tmp14860 := PrimHead(V1722)

			tmp14861 := Call(__e, PrimNS2Value(symshen_4next_150), MakeNumber(50), tmp14860)

			tmp14862 := Call(__e, PrimNS2Value(symshen_4app), tmp14861, MakeString("\n"), symshen_4a)

			tmp14863 := PrimStringConcat(MakeString("datatype syntax error here:\n\n "), tmp14862)

			__e.Return(PrimSimpleError(tmp14863))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4datatype_1error)
			return
		}

	}, 1)

	tmp14874 := Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1error, tmp14858)

	_ = tmp14874

	tmp14875 := MakeNative(func(__e *ControlFlow) {
		V1724 := __e.Get(1)
		_ = V1724
		tmp14876 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14885 := Call(__e, PrimNS2Value(symfail))

			tmp14886 := PrimEqual(YaccParse, tmp14885)

			if True == tmp14886 {
				tmp14878 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp14881 := Call(__e, PrimNS2Value(symfail))

					tmp14882 := PrimEqual(tmp14881, Parse___5e_6)

					tmp14883 := PrimNot(tmp14882)

					if True == tmp14883 {
						tmp14880 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14880, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14884 := Call(__e, PrimNS2Value(sym_5e_6), V1724)

				__e.TailApply(tmp14878, tmp14884)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14887 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5datatype_1rule_6 := __e.Get(1)
			_ = Parse__shen_4_5datatype_1rule_6
			tmp14899 := Call(__e, PrimNS2Value(symfail))

			tmp14900 := PrimEqual(tmp14899, Parse__shen_4_5datatype_1rule_6)

			tmp14901 := PrimNot(tmp14900)

			if True == tmp14901 {
				tmp14889 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5datatype_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5datatype_1rules_6
					tmp14895 := Call(__e, PrimNS2Value(symfail))

					tmp14896 := PrimEqual(tmp14895, Parse__shen_4_5datatype_1rules_6)

					tmp14897 := PrimNot(tmp14896)

					if True == tmp14897 {
						tmp14891 := PrimHead(Parse__shen_4_5datatype_1rules_6)

						tmp14892 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5datatype_1rule_6)

						tmp14893 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5datatype_1rules_6)

						tmp14894 := PrimCons(tmp14892, tmp14893)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14891, tmp14894)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14898 := Call(__e, PrimNS2Value(symshen_4_5datatype_1rules_6), Parse__shen_4_5datatype_1rule_6)

				__e.TailApply(tmp14889, tmp14898)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14902 := Call(__e, PrimNS2Value(symshen_4_5datatype_1rule_6), V1724)

		tmp14903 := Call(__e, tmp14887, tmp14902)

		__e.TailApply(tmp14876, tmp14903)
		return

	}, 1)

	tmp14904 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rules_6, tmp14875)

	_ = tmp14904

	tmp14905 := MakeNative(func(__e *ControlFlow) {
		V1726 := __e.Get(1)
		_ = V1726
		tmp14906 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14940 := Call(__e, PrimNS2Value(symfail))

			tmp14941 := PrimEqual(YaccParse, tmp14940)

			if True == tmp14941 {
				tmp14908 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					tmp14936 := Call(__e, PrimNS2Value(symfail))

					tmp14937 := PrimEqual(tmp14936, Parse__shen_4_5side_1conditions_6)

					tmp14938 := PrimNot(tmp14937)

					if True == tmp14938 {
						tmp14910 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							tmp14932 := Call(__e, PrimNS2Value(symfail))

							tmp14933 := PrimEqual(tmp14932, Parse__shen_4_5premises_6)

							tmp14934 := PrimNot(tmp14933)

							if True == tmp14934 {
								tmp14912 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5doubleunderline_6 := __e.Get(1)
									_ = Parse__shen_4_5doubleunderline_6
									tmp14928 := Call(__e, PrimNS2Value(symfail))

									tmp14929 := PrimEqual(tmp14928, Parse__shen_4_5doubleunderline_6)

									tmp14930 := PrimNot(tmp14929)

									if True == tmp14930 {
										tmp14914 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5conclusion_6 := __e.Get(1)
											_ = Parse__shen_4_5conclusion_6
											tmp14924 := Call(__e, PrimNS2Value(symfail))

											tmp14925 := PrimEqual(tmp14924, Parse__shen_4_5conclusion_6)

											tmp14926 := PrimNot(tmp14925)

											if True == tmp14926 {
												tmp14916 := PrimHead(Parse__shen_4_5conclusion_6)

												tmp14917 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

												tmp14918 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5premises_6)

												tmp14919 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5conclusion_6)

												tmp14920 := PrimCons(tmp14919, Nil)

												tmp14921 := PrimCons(tmp14918, tmp14920)

												tmp14922 := PrimCons(tmp14917, tmp14921)

												tmp14923 := Call(__e, PrimNS2Value(symshen_4sequent), symshen_4double, tmp14922)

												__e.TailApply(PrimNS2Value(symshen_4pair), tmp14916, tmp14923)
												return

											} else {
												__e.TailApply(PrimNS2Value(symfail))
												return
											}

										}, 1)

										tmp14927 := Call(__e, PrimNS2Value(symshen_4_5conclusion_6), Parse__shen_4_5doubleunderline_6)

										__e.TailApply(tmp14914, tmp14927)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14931 := Call(__e, PrimNS2Value(symshen_4_5doubleunderline_6), Parse__shen_4_5premises_6)

								__e.TailApply(tmp14912, tmp14931)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14935 := Call(__e, PrimNS2Value(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

						__e.TailApply(tmp14910, tmp14935)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14939 := Call(__e, PrimNS2Value(symshen_4_5side_1conditions_6), V1726)

				__e.TailApply(tmp14908, tmp14939)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14942 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1conditions_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1conditions_6
			tmp14970 := Call(__e, PrimNS2Value(symfail))

			tmp14971 := PrimEqual(tmp14970, Parse__shen_4_5side_1conditions_6)

			tmp14972 := PrimNot(tmp14971)

			if True == tmp14972 {
				tmp14944 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5premises_6 := __e.Get(1)
					_ = Parse__shen_4_5premises_6
					tmp14966 := Call(__e, PrimNS2Value(symfail))

					tmp14967 := PrimEqual(tmp14966, Parse__shen_4_5premises_6)

					tmp14968 := PrimNot(tmp14967)

					if True == tmp14968 {
						tmp14946 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5singleunderline_6 := __e.Get(1)
							_ = Parse__shen_4_5singleunderline_6
							tmp14962 := Call(__e, PrimNS2Value(symfail))

							tmp14963 := PrimEqual(tmp14962, Parse__shen_4_5singleunderline_6)

							tmp14964 := PrimNot(tmp14963)

							if True == tmp14964 {
								tmp14948 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5conclusion_6 := __e.Get(1)
									_ = Parse__shen_4_5conclusion_6
									tmp14958 := Call(__e, PrimNS2Value(symfail))

									tmp14959 := PrimEqual(tmp14958, Parse__shen_4_5conclusion_6)

									tmp14960 := PrimNot(tmp14959)

									if True == tmp14960 {
										tmp14950 := PrimHead(Parse__shen_4_5conclusion_6)

										tmp14951 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

										tmp14952 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5premises_6)

										tmp14953 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5conclusion_6)

										tmp14954 := PrimCons(tmp14953, Nil)

										tmp14955 := PrimCons(tmp14952, tmp14954)

										tmp14956 := PrimCons(tmp14951, tmp14955)

										tmp14957 := Call(__e, PrimNS2Value(symshen_4sequent), symshen_4single, tmp14956)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp14950, tmp14957)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp14961 := Call(__e, PrimNS2Value(symshen_4_5conclusion_6), Parse__shen_4_5singleunderline_6)

								__e.TailApply(tmp14948, tmp14961)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp14965 := Call(__e, PrimNS2Value(symshen_4_5singleunderline_6), Parse__shen_4_5premises_6)

						__e.TailApply(tmp14946, tmp14965)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14969 := Call(__e, PrimNS2Value(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

				__e.TailApply(tmp14944, tmp14969)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp14973 := Call(__e, PrimNS2Value(symshen_4_5side_1conditions_6), V1726)

		tmp14974 := Call(__e, tmp14942, tmp14973)

		__e.TailApply(tmp14906, tmp14974)
		return

	}, 1)

	tmp14975 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rule_6, tmp14905)

	_ = tmp14975

	tmp14976 := MakeNative(func(__e *ControlFlow) {
		V1728 := __e.Get(1)
		_ = V1728
		tmp14977 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp14986 := Call(__e, PrimNS2Value(symfail))

			tmp14987 := PrimEqual(YaccParse, tmp14986)

			if True == tmp14987 {
				tmp14979 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp14982 := Call(__e, PrimNS2Value(symfail))

					tmp14983 := PrimEqual(tmp14982, Parse___5e_6)

					tmp14984 := PrimNot(tmp14983)

					if True == tmp14984 {
						tmp14981 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14981, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14985 := Call(__e, PrimNS2Value(sym_5e_6), V1728)

				__e.TailApply(tmp14979, tmp14985)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp14988 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1condition_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1condition_6
			tmp15000 := Call(__e, PrimNS2Value(symfail))

			tmp15001 := PrimEqual(tmp15000, Parse__shen_4_5side_1condition_6)

			tmp15002 := PrimNot(tmp15001)

			if True == tmp15002 {
				tmp14990 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					tmp14996 := Call(__e, PrimNS2Value(symfail))

					tmp14997 := PrimEqual(tmp14996, Parse__shen_4_5side_1conditions_6)

					tmp14998 := PrimNot(tmp14997)

					if True == tmp14998 {
						tmp14992 := PrimHead(Parse__shen_4_5side_1conditions_6)

						tmp14993 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5side_1condition_6)

						tmp14994 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

						tmp14995 := PrimCons(tmp14993, tmp14994)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp14992, tmp14995)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp14999 := Call(__e, PrimNS2Value(symshen_4_5side_1conditions_6), Parse__shen_4_5side_1condition_6)

				__e.TailApply(tmp14990, tmp14999)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15003 := Call(__e, PrimNS2Value(symshen_4_5side_1condition_6), V1728)

		tmp15004 := Call(__e, tmp14988, tmp15003)

		__e.TailApply(tmp14977, tmp15004)
		return

	}, 1)

	tmp15005 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1conditions_6, tmp14976)

	_ = tmp15005

	tmp15006 := MakeNative(func(__e *ControlFlow) {
		V1732 := __e.Get(1)
		_ = V1732
		tmp15007 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15038 := Call(__e, PrimNS2Value(symfail))

			tmp15039 := PrimEqual(YaccParse, tmp15038)

			if True == tmp15039 {
				tmp15036 := PrimHead(V1732)

				tmp15037 := PrimIsPair(tmp15036)

				var ifres15032 Obj

				if True == tmp15037 {
					tmp15034 := Call(__e, PrimNS2Value(symshen_4hdhd), V1732)

					tmp15035 := PrimEqual(symlet, tmp15034)

					var ifres15033 Obj

					if True == tmp15035 {
						ifres15033 = True

					} else {
						ifres15033 = False

					}

					ifres15032 = ifres15033

				} else {
					ifres15032 = False

				}

				if True == ifres15032 {
					tmp15010 := MakeNative(func(__e *ControlFlow) {
						NewStream1730 := __e.Get(1)
						_ = NewStream1730
						tmp15011 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5variable_2_6 := __e.Get(1)
							_ = Parse__shen_4_5variable_2_6
							tmp15025 := Call(__e, PrimNS2Value(symfail))

							tmp15026 := PrimEqual(tmp15025, Parse__shen_4_5variable_2_6)

							tmp15027 := PrimNot(tmp15026)

							if True == tmp15027 {
								tmp15013 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5expr_6 := __e.Get(1)
									_ = Parse__shen_4_5expr_6
									tmp15021 := Call(__e, PrimNS2Value(symfail))

									tmp15022 := PrimEqual(tmp15021, Parse__shen_4_5expr_6)

									tmp15023 := PrimNot(tmp15022)

									if True == tmp15023 {
										tmp15015 := PrimHead(Parse__shen_4_5expr_6)

										tmp15016 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5variable_2_6)

										tmp15017 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

										tmp15018 := PrimCons(tmp15017, Nil)

										tmp15019 := PrimCons(tmp15016, tmp15018)

										tmp15020 := PrimCons(symlet, tmp15019)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp15015, tmp15020)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp15024 := Call(__e, PrimNS2Value(symshen_4_5expr_6), Parse__shen_4_5variable_2_6)

								__e.TailApply(tmp15013, tmp15024)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15028 := Call(__e, PrimNS2Value(symshen_4_5variable_2_6), NewStream1730)

						__e.TailApply(tmp15011, tmp15028)
						return

					}, 1)

					tmp15029 := Call(__e, PrimNS2Value(symshen_4tlhd), V1732)

					tmp15030 := Call(__e, PrimNS2Value(symshen_4hdtl), V1732)

					tmp15031 := Call(__e, PrimNS2Value(symshen_4pair), tmp15029, tmp15030)

					__e.TailApply(tmp15010, tmp15031)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15061 := PrimHead(V1732)

		tmp15062 := PrimIsPair(tmp15061)

		var ifres15057 Obj

		if True == tmp15062 {
			tmp15059 := Call(__e, PrimNS2Value(symshen_4hdhd), V1732)

			tmp15060 := PrimEqual(symif, tmp15059)

			var ifres15058 Obj

			if True == tmp15060 {
				ifres15058 = True

			} else {
				ifres15058 = False

			}

			ifres15057 = ifres15058

		} else {
			ifres15057 = False

		}

		var ifres15040 Obj

		if True == ifres15057 {
			tmp15042 := MakeNative(func(__e *ControlFlow) {
				NewStream1729 := __e.Get(1)
				_ = NewStream1729
				tmp15043 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					tmp15049 := Call(__e, PrimNS2Value(symfail))

					tmp15050 := PrimEqual(tmp15049, Parse__shen_4_5expr_6)

					tmp15051 := PrimNot(tmp15050)

					if True == tmp15051 {
						tmp15045 := PrimHead(Parse__shen_4_5expr_6)

						tmp15046 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

						tmp15047 := PrimCons(tmp15046, Nil)

						tmp15048 := PrimCons(symif, tmp15047)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp15045, tmp15048)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15052 := Call(__e, PrimNS2Value(symshen_4_5expr_6), NewStream1729)

				__e.TailApply(tmp15043, tmp15052)
				return

			}, 1)

			tmp15053 := Call(__e, PrimNS2Value(symshen_4tlhd), V1732)

			tmp15054 := Call(__e, PrimNS2Value(symshen_4hdtl), V1732)

			tmp15055 := Call(__e, PrimNS2Value(symshen_4pair), tmp15053, tmp15054)

			tmp15056 := Call(__e, tmp15042, tmp15055)

			ifres15040 = tmp15056

		} else {
			tmp15041 := Call(__e, PrimNS2Value(symfail))

			ifres15040 = tmp15041

		}

		__e.TailApply(tmp15007, ifres15040)
		return

	}, 1)

	tmp15063 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1condition_6, tmp15006)

	_ = tmp15063

	tmp15064 := MakeNative(func(__e *ControlFlow) {
		V1734 := __e.Get(1)
		_ = V1734
		tmp15074 := PrimHead(V1734)

		tmp15075 := PrimIsPair(tmp15074)

		if True == tmp15075 {
			tmp15066 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15072 := PrimIsVariable(Parse__X)

				if True == tmp15072 {
					tmp15068 := Call(__e, PrimNS2Value(symshen_4tlhd), V1734)

					tmp15069 := Call(__e, PrimNS2Value(symshen_4hdtl), V1734)

					tmp15070 := Call(__e, PrimNS2Value(symshen_4pair), tmp15068, tmp15069)

					tmp15071 := PrimHead(tmp15070)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15071, Parse__X)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15073 := Call(__e, PrimNS2Value(symshen_4hdhd), V1734)

			__e.TailApply(tmp15066, tmp15073)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15076 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5variable_2_6, tmp15064)

	_ = tmp15076

	tmp15077 := MakeNative(func(__e *ControlFlow) {
		V1736 := __e.Get(1)
		_ = V1736
		tmp15097 := PrimHead(V1736)

		tmp15098 := PrimIsPair(tmp15097)

		if True == tmp15098 {
			tmp15079 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15092 := PrimCons(sym_k, Nil)

				tmp15093 := PrimCons(sym_6_6, tmp15092)

				tmp15094 := Call(__e, PrimNS2Value(symelement_2), Parse__X, tmp15093)

				var ifres15086 Obj

				if True == tmp15094 {
					ifres15086 = True

				} else {
					tmp15091 := Call(__e, PrimNS2Value(symshen_4singleunderline_2), Parse__X)

					var ifres15088 Obj

					if True == tmp15091 {
						ifres15088 = True

					} else {
						tmp15090 := Call(__e, PrimNS2Value(symshen_4doubleunderline_2), Parse__X)

						var ifres15089 Obj

						if True == tmp15090 {
							ifres15089 = True

						} else {
							ifres15089 = False

						}

						ifres15088 = ifres15089

					}

					var ifres15087 Obj

					if True == ifres15088 {
						ifres15087 = True

					} else {
						ifres15087 = False

					}

					ifres15086 = ifres15087

				}

				tmp15095 := PrimNot(ifres15086)

				if True == tmp15095 {
					tmp15081 := Call(__e, PrimNS2Value(symshen_4tlhd), V1736)

					tmp15082 := Call(__e, PrimNS2Value(symshen_4hdtl), V1736)

					tmp15083 := Call(__e, PrimNS2Value(symshen_4pair), tmp15081, tmp15082)

					tmp15084 := PrimHead(tmp15083)

					tmp15085 := Call(__e, PrimNS2Value(symshen_4remove_1bar), Parse__X)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15084, tmp15085)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15096 := Call(__e, PrimNS2Value(symshen_4hdhd), V1736)

			__e.TailApply(tmp15079, tmp15096)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15099 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5expr_6, tmp15077)

	_ = tmp15099

	tmp15100 := MakeNative(func(__e *ControlFlow) {
		V1738 := __e.Get(1)
		_ = V1738
		tmp15132 := PrimIsPair(V1738)

		var ifres15112 Obj

		if True == tmp15132 {
			tmp15130 := PrimTail(V1738)

			tmp15131 := PrimIsPair(tmp15130)

			var ifres15114 Obj

			if True == tmp15131 {
				tmp15127 := PrimTail(V1738)

				tmp15128 := PrimTail(tmp15127)

				tmp15129 := PrimIsPair(tmp15128)

				var ifres15116 Obj

				if True == tmp15129 {
					tmp15123 := PrimTail(V1738)

					tmp15124 := PrimTail(tmp15123)

					tmp15125 := PrimTail(tmp15124)

					tmp15126 := PrimEqual(Nil, tmp15125)

					var ifres15118 Obj

					if True == tmp15126 {
						tmp15120 := PrimTail(V1738)

						tmp15121 := PrimHead(tmp15120)

						tmp15122 := PrimEqual(tmp15121, symbar_b)

						var ifres15119 Obj

						if True == tmp15122 {
							ifres15119 = True

						} else {
							ifres15119 = False

						}

						ifres15118 = ifres15119

					} else {
						ifres15118 = False

					}

					var ifres15117 Obj

					if True == ifres15118 {
						ifres15117 = True

					} else {
						ifres15117 = False

					}

					ifres15116 = ifres15117

				} else {
					ifres15116 = False

				}

				var ifres15115 Obj

				if True == ifres15116 {
					ifres15115 = True

				} else {
					ifres15115 = False

				}

				ifres15114 = ifres15115

			} else {
				ifres15114 = False

			}

			var ifres15113 Obj

			if True == ifres15114 {
				ifres15113 = True

			} else {
				ifres15113 = False

			}

			ifres15112 = ifres15113

		} else {
			ifres15112 = False

		}

		if True == ifres15112 {
			tmp15108 := PrimHead(V1738)

			tmp15109 := PrimTail(V1738)

			tmp15110 := PrimTail(tmp15109)

			tmp15111 := PrimHead(tmp15110)

			__e.Return(PrimCons(tmp15108, tmp15111))
			return

		} else {
			tmp15107 := PrimIsPair(V1738)

			if True == tmp15107 {
				tmp15103 := PrimHead(V1738)

				tmp15104 := Call(__e, PrimNS2Value(symshen_4remove_1bar), tmp15103)

				tmp15105 := PrimTail(V1738)

				tmp15106 := Call(__e, PrimNS2Value(symshen_4remove_1bar), tmp15105)

				__e.Return(PrimCons(tmp15104, tmp15106))
				return

			} else {
				__e.Return(V1738)
				return
			}

		}

	}, 1)

	tmp15133 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1bar, tmp15100)

	_ = tmp15133

	tmp15134 := MakeNative(func(__e *ControlFlow) {
		V1740 := __e.Get(1)
		_ = V1740
		tmp15135 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15144 := Call(__e, PrimNS2Value(symfail))

			tmp15145 := PrimEqual(YaccParse, tmp15144)

			if True == tmp15145 {
				tmp15137 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp15140 := Call(__e, PrimNS2Value(symfail))

					tmp15141 := PrimEqual(tmp15140, Parse___5e_6)

					tmp15142 := PrimNot(tmp15141)

					if True == tmp15142 {
						tmp15139 := PrimHead(Parse___5e_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp15139, Nil)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15143 := Call(__e, PrimNS2Value(sym_5e_6), V1740)

				__e.TailApply(tmp15137, tmp15143)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15146 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5premise_6 := __e.Get(1)
			_ = Parse__shen_4_5premise_6
			tmp15164 := Call(__e, PrimNS2Value(symfail))

			tmp15165 := PrimEqual(tmp15164, Parse__shen_4_5premise_6)

			tmp15166 := PrimNot(tmp15165)

			if True == tmp15166 {
				tmp15148 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5semicolon_1symbol_6
					tmp15160 := Call(__e, PrimNS2Value(symfail))

					tmp15161 := PrimEqual(tmp15160, Parse__shen_4_5semicolon_1symbol_6)

					tmp15162 := PrimNot(tmp15161)

					if True == tmp15162 {
						tmp15150 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							tmp15156 := Call(__e, PrimNS2Value(symfail))

							tmp15157 := PrimEqual(tmp15156, Parse__shen_4_5premises_6)

							tmp15158 := PrimNot(tmp15157)

							if True == tmp15158 {
								tmp15152 := PrimHead(Parse__shen_4_5premises_6)

								tmp15153 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5premise_6)

								tmp15154 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5premises_6)

								tmp15155 := PrimCons(tmp15153, tmp15154)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15152, tmp15155)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15159 := Call(__e, PrimNS2Value(symshen_4_5premises_6), Parse__shen_4_5semicolon_1symbol_6)

						__e.TailApply(tmp15150, tmp15159)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15163 := Call(__e, PrimNS2Value(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5premise_6)

				__e.TailApply(tmp15148, tmp15163)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15167 := Call(__e, PrimNS2Value(symshen_4_5premise_6), V1740)

		tmp15168 := Call(__e, tmp15146, tmp15167)

		__e.TailApply(tmp15135, tmp15168)
		return

	}, 1)

	tmp15169 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premises_6, tmp15134)

	_ = tmp15169

	tmp15170 := MakeNative(func(__e *ControlFlow) {
		V1742 := __e.Get(1)
		_ = V1742
		tmp15180 := PrimHead(V1742)

		tmp15181 := PrimIsPair(tmp15180)

		if True == tmp15181 {
			tmp15172 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15178 := PrimEqual(Parse__X, sym_k)

				if True == tmp15178 {
					tmp15174 := Call(__e, PrimNS2Value(symshen_4tlhd), V1742)

					tmp15175 := Call(__e, PrimNS2Value(symshen_4hdtl), V1742)

					tmp15176 := Call(__e, PrimNS2Value(symshen_4pair), tmp15174, tmp15175)

					tmp15177 := PrimHead(tmp15176)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15177, symshen_4skip)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15179 := Call(__e, PrimNS2Value(symshen_4hdhd), V1742)

			__e.TailApply(tmp15172, tmp15179)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15182 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_1symbol_6, tmp15170)

	_ = tmp15182

	tmp15183 := MakeNative(func(__e *ControlFlow) {
		V1746 := __e.Get(1)
		_ = V1746
		tmp15184 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15227 := Call(__e, PrimNS2Value(symfail))

			tmp15228 := PrimEqual(YaccParse, tmp15227)

			if True == tmp15228 {
				tmp15186 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp15197 := Call(__e, PrimNS2Value(symfail))

					tmp15198 := PrimEqual(YaccParse, tmp15197)

					if True == tmp15198 {
						tmp15188 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							tmp15193 := Call(__e, PrimNS2Value(symfail))

							tmp15194 := PrimEqual(tmp15193, Parse__shen_4_5formula_6)

							tmp15195 := PrimNot(tmp15194)

							if True == tmp15195 {
								tmp15190 := PrimHead(Parse__shen_4_5formula_6)

								tmp15191 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

								tmp15192 := Call(__e, PrimNS2Value(symshen_4sequent), Nil, tmp15191)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15190, tmp15192)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15196 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V1746)

						__e.TailApply(tmp15188, tmp15196)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp15199 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formulae_6 := __e.Get(1)
					_ = Parse__shen_4_5formulae_6
					tmp15222 := Call(__e, PrimNS2Value(symfail))

					tmp15223 := PrimEqual(tmp15222, Parse__shen_4_5formulae_6)

					tmp15224 := PrimNot(tmp15223)

					if True == tmp15224 {
						tmp15220 := PrimHead(Parse__shen_4_5formulae_6)

						tmp15221 := PrimIsPair(tmp15220)

						var ifres15216 Obj

						if True == tmp15221 {
							tmp15218 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5formulae_6)

							tmp15219 := PrimEqual(sym_6_6, tmp15218)

							var ifres15217 Obj

							if True == tmp15219 {
								ifres15217 = True

							} else {
								ifres15217 = False

							}

							ifres15216 = ifres15217

						} else {
							ifres15216 = False

						}

						if True == ifres15216 {
							tmp15202 := MakeNative(func(__e *ControlFlow) {
								NewStream1744 := __e.Get(1)
								_ = NewStream1744
								tmp15203 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5formula_6 := __e.Get(1)
									_ = Parse__shen_4_5formula_6
									tmp15209 := Call(__e, PrimNS2Value(symfail))

									tmp15210 := PrimEqual(tmp15209, Parse__shen_4_5formula_6)

									tmp15211 := PrimNot(tmp15210)

									if True == tmp15211 {
										tmp15205 := PrimHead(Parse__shen_4_5formula_6)

										tmp15206 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formulae_6)

										tmp15207 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

										tmp15208 := Call(__e, PrimNS2Value(symshen_4sequent), tmp15206, tmp15207)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp15205, tmp15208)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp15212 := Call(__e, PrimNS2Value(symshen_4_5formula_6), NewStream1744)

								__e.TailApply(tmp15203, tmp15212)
								return

							}, 1)

							tmp15213 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5formulae_6)

							tmp15214 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formulae_6)

							tmp15215 := Call(__e, PrimNS2Value(symshen_4pair), tmp15213, tmp15214)

							__e.TailApply(tmp15202, tmp15215)
							return

						} else {
							__e.TailApply(PrimNS2Value(symfail))
							return
						}

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15225 := Call(__e, PrimNS2Value(symshen_4_5formulae_6), V1746)

				tmp15226 := Call(__e, tmp15199, tmp15225)

				__e.TailApply(tmp15186, tmp15226)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15241 := PrimHead(V1746)

		tmp15242 := PrimIsPair(tmp15241)

		var ifres15237 Obj

		if True == tmp15242 {
			tmp15239 := Call(__e, PrimNS2Value(symshen_4hdhd), V1746)

			tmp15240 := PrimEqual(sym_b, tmp15239)

			var ifres15238 Obj

			if True == tmp15240 {
				ifres15238 = True

			} else {
				ifres15238 = False

			}

			ifres15237 = ifres15238

		} else {
			ifres15237 = False

		}

		var ifres15229 Obj

		if True == ifres15237 {
			tmp15231 := MakeNative(func(__e *ControlFlow) {
				NewStream1743 := __e.Get(1)
				_ = NewStream1743
				tmp15232 := PrimHead(NewStream1743)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp15232, sym_b)
				return

			}, 1)

			tmp15233 := Call(__e, PrimNS2Value(symshen_4tlhd), V1746)

			tmp15234 := Call(__e, PrimNS2Value(symshen_4hdtl), V1746)

			tmp15235 := Call(__e, PrimNS2Value(symshen_4pair), tmp15233, tmp15234)

			tmp15236 := Call(__e, tmp15231, tmp15235)

			ifres15229 = tmp15236

		} else {
			tmp15230 := Call(__e, PrimNS2Value(symfail))

			ifres15229 = tmp15230

		}

		__e.TailApply(tmp15184, ifres15229)
		return

	}, 1)

	tmp15243 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premise_6, tmp15183)

	_ = tmp15243

	tmp15244 := MakeNative(func(__e *ControlFlow) {
		V1749 := __e.Get(1)
		_ = V1749
		tmp15245 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15262 := Call(__e, PrimNS2Value(symfail))

			tmp15263 := PrimEqual(YaccParse, tmp15262)

			if True == tmp15263 {
				tmp15247 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					tmp15258 := Call(__e, PrimNS2Value(symfail))

					tmp15259 := PrimEqual(tmp15258, Parse__shen_4_5formula_6)

					tmp15260 := PrimNot(tmp15259)

					if True == tmp15260 {
						tmp15249 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
							_ = Parse__shen_4_5semicolon_1symbol_6
							tmp15254 := Call(__e, PrimNS2Value(symfail))

							tmp15255 := PrimEqual(tmp15254, Parse__shen_4_5semicolon_1symbol_6)

							tmp15256 := PrimNot(tmp15255)

							if True == tmp15256 {
								tmp15251 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)

								tmp15252 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

								tmp15253 := Call(__e, PrimNS2Value(symshen_4sequent), Nil, tmp15252)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15251, tmp15253)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15257 := Call(__e, PrimNS2Value(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

						__e.TailApply(tmp15249, tmp15257)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15261 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V1749)

				__e.TailApply(tmp15247, tmp15261)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15264 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formulae_6 := __e.Get(1)
			_ = Parse__shen_4_5formulae_6
			tmp15293 := Call(__e, PrimNS2Value(symfail))

			tmp15294 := PrimEqual(tmp15293, Parse__shen_4_5formulae_6)

			tmp15295 := PrimNot(tmp15294)

			if True == tmp15295 {
				tmp15291 := PrimHead(Parse__shen_4_5formulae_6)

				tmp15292 := PrimIsPair(tmp15291)

				var ifres15287 Obj

				if True == tmp15292 {
					tmp15289 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5formulae_6)

					tmp15290 := PrimEqual(sym_6_6, tmp15289)

					var ifres15288 Obj

					if True == tmp15290 {
						ifres15288 = True

					} else {
						ifres15288 = False

					}

					ifres15287 = ifres15288

				} else {
					ifres15287 = False

				}

				if True == ifres15287 {
					tmp15267 := MakeNative(func(__e *ControlFlow) {
						NewStream1747 := __e.Get(1)
						_ = NewStream1747
						tmp15268 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							tmp15280 := Call(__e, PrimNS2Value(symfail))

							tmp15281 := PrimEqual(tmp15280, Parse__shen_4_5formula_6)

							tmp15282 := PrimNot(tmp15281)

							if True == tmp15282 {
								tmp15270 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
									_ = Parse__shen_4_5semicolon_1symbol_6
									tmp15276 := Call(__e, PrimNS2Value(symfail))

									tmp15277 := PrimEqual(tmp15276, Parse__shen_4_5semicolon_1symbol_6)

									tmp15278 := PrimNot(tmp15277)

									if True == tmp15278 {
										tmp15272 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)

										tmp15273 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formulae_6)

										tmp15274 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

										tmp15275 := Call(__e, PrimNS2Value(symshen_4sequent), tmp15273, tmp15274)

										__e.TailApply(PrimNS2Value(symshen_4pair), tmp15272, tmp15275)
										return

									} else {
										__e.TailApply(PrimNS2Value(symfail))
										return
									}

								}, 1)

								tmp15279 := Call(__e, PrimNS2Value(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

								__e.TailApply(tmp15270, tmp15279)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15283 := Call(__e, PrimNS2Value(symshen_4_5formula_6), NewStream1747)

						__e.TailApply(tmp15268, tmp15283)
						return

					}, 1)

					tmp15284 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5formulae_6)

					tmp15285 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formulae_6)

					tmp15286 := Call(__e, PrimNS2Value(symshen_4pair), tmp15284, tmp15285)

					__e.TailApply(tmp15267, tmp15286)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15296 := Call(__e, PrimNS2Value(symshen_4_5formulae_6), V1749)

		tmp15297 := Call(__e, tmp15264, tmp15296)

		__e.TailApply(tmp15245, tmp15297)
		return

	}, 1)

	tmp15298 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5conclusion_6, tmp15244)

	_ = tmp15298

	tmp15299 := MakeNative(func(__e *ControlFlow) {
		V1752 := __e.Get(1)
		_ = V1752
		V1753 := __e.Get(2)
		_ = V1753
		__e.TailApply(PrimNS2Value(sym_8p), V1752, V1753)
		return
	}, 2)

	tmp15300 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sequent, tmp15299)

	_ = tmp15300

	tmp15301 := MakeNative(func(__e *ControlFlow) {
		V1755 := __e.Get(1)
		_ = V1755
		tmp15302 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15325 := Call(__e, PrimNS2Value(symfail))

			tmp15326 := PrimEqual(YaccParse, tmp15325)

			if True == tmp15326 {
				tmp15304 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp15313 := Call(__e, PrimNS2Value(symfail))

					tmp15314 := PrimEqual(YaccParse, tmp15313)

					if True == tmp15314 {
						tmp15306 := MakeNative(func(__e *ControlFlow) {
							Parse___5e_6 := __e.Get(1)
							_ = Parse___5e_6
							tmp15309 := Call(__e, PrimNS2Value(symfail))

							tmp15310 := PrimEqual(tmp15309, Parse___5e_6)

							tmp15311 := PrimNot(tmp15310)

							if True == tmp15311 {
								tmp15308 := PrimHead(Parse___5e_6)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15308, Nil)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15312 := Call(__e, PrimNS2Value(sym_5e_6), V1755)

						__e.TailApply(tmp15306, tmp15312)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp15315 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					tmp15320 := Call(__e, PrimNS2Value(symfail))

					tmp15321 := PrimEqual(tmp15320, Parse__shen_4_5formula_6)

					tmp15322 := PrimNot(tmp15321)

					if True == tmp15322 {
						tmp15317 := PrimHead(Parse__shen_4_5formula_6)

						tmp15318 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

						tmp15319 := PrimCons(tmp15318, Nil)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp15317, tmp15319)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15323 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V1755)

				tmp15324 := Call(__e, tmp15315, tmp15323)

				__e.TailApply(tmp15304, tmp15324)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15327 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formula_6 := __e.Get(1)
			_ = Parse__shen_4_5formula_6
			tmp15345 := Call(__e, PrimNS2Value(symfail))

			tmp15346 := PrimEqual(tmp15345, Parse__shen_4_5formula_6)

			tmp15347 := PrimNot(tmp15346)

			if True == tmp15347 {
				tmp15329 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5comma_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5comma_1symbol_6
					tmp15341 := Call(__e, PrimNS2Value(symfail))

					tmp15342 := PrimEqual(tmp15341, Parse__shen_4_5comma_1symbol_6)

					tmp15343 := PrimNot(tmp15342)

					if True == tmp15343 {
						tmp15331 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formulae_6 := __e.Get(1)
							_ = Parse__shen_4_5formulae_6
							tmp15337 := Call(__e, PrimNS2Value(symfail))

							tmp15338 := PrimEqual(tmp15337, Parse__shen_4_5formulae_6)

							tmp15339 := PrimNot(tmp15338)

							if True == tmp15339 {
								tmp15333 := PrimHead(Parse__shen_4_5formulae_6)

								tmp15334 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formula_6)

								tmp15335 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5formulae_6)

								tmp15336 := PrimCons(tmp15334, tmp15335)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15333, tmp15336)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15340 := Call(__e, PrimNS2Value(symshen_4_5formulae_6), Parse__shen_4_5comma_1symbol_6)

						__e.TailApply(tmp15331, tmp15340)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15344 := Call(__e, PrimNS2Value(symshen_4_5comma_1symbol_6), Parse__shen_4_5formula_6)

				__e.TailApply(tmp15329, tmp15344)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15348 := Call(__e, PrimNS2Value(symshen_4_5formula_6), V1755)

		tmp15349 := Call(__e, tmp15327, tmp15348)

		__e.TailApply(tmp15302, tmp15349)
		return

	}, 1)

	tmp15350 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formulae_6, tmp15301)

	_ = tmp15350

	tmp15351 := MakeNative(func(__e *ControlFlow) {
		V1757 := __e.Get(1)
		_ = V1757
		tmp15362 := PrimHead(V1757)

		tmp15363 := PrimIsPair(tmp15362)

		if True == tmp15363 {
			tmp15353 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15359 := PrimIntern(MakeString(","))

				tmp15360 := PrimEqual(Parse__X, tmp15359)

				if True == tmp15360 {
					tmp15355 := Call(__e, PrimNS2Value(symshen_4tlhd), V1757)

					tmp15356 := Call(__e, PrimNS2Value(symshen_4hdtl), V1757)

					tmp15357 := Call(__e, PrimNS2Value(symshen_4pair), tmp15355, tmp15356)

					tmp15358 := PrimHead(tmp15357)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15358, symshen_4skip)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15361 := Call(__e, PrimNS2Value(symshen_4hdhd), V1757)

			__e.TailApply(tmp15353, tmp15361)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15364 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_1symbol_6, tmp15351)

	_ = tmp15364

	tmp15365 := MakeNative(func(__e *ControlFlow) {
		V1760 := __e.Get(1)
		_ = V1760
		tmp15366 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp15376 := Call(__e, PrimNS2Value(symfail))

			tmp15377 := PrimEqual(YaccParse, tmp15376)

			if True == tmp15377 {
				tmp15368 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					tmp15372 := Call(__e, PrimNS2Value(symfail))

					tmp15373 := PrimEqual(tmp15372, Parse__shen_4_5expr_6)

					tmp15374 := PrimNot(tmp15373)

					if True == tmp15374 {
						tmp15370 := PrimHead(Parse__shen_4_5expr_6)

						tmp15371 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp15370, tmp15371)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp15375 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V1760)

				__e.TailApply(tmp15368, tmp15375)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp15378 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			tmp15405 := Call(__e, PrimNS2Value(symfail))

			tmp15406 := PrimEqual(tmp15405, Parse__shen_4_5expr_6)

			tmp15407 := PrimNot(tmp15406)

			if True == tmp15407 {
				tmp15403 := PrimHead(Parse__shen_4_5expr_6)

				tmp15404 := PrimIsPair(tmp15403)

				var ifres15399 Obj

				if True == tmp15404 {
					tmp15401 := Call(__e, PrimNS2Value(symshen_4hdhd), Parse__shen_4_5expr_6)

					tmp15402 := PrimEqual(sym_h, tmp15401)

					var ifres15400 Obj

					if True == tmp15402 {
						ifres15400 = True

					} else {
						ifres15400 = False

					}

					ifres15399 = ifres15400

				} else {
					ifres15399 = False

				}

				if True == ifres15399 {
					tmp15381 := MakeNative(func(__e *ControlFlow) {
						NewStream1758 := __e.Get(1)
						_ = NewStream1758
						tmp15382 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5type_6 := __e.Get(1)
							_ = Parse__shen_4_5type_6
							tmp15392 := Call(__e, PrimNS2Value(symfail))

							tmp15393 := PrimEqual(tmp15392, Parse__shen_4_5type_6)

							tmp15394 := PrimNot(tmp15393)

							if True == tmp15394 {
								tmp15384 := PrimHead(Parse__shen_4_5type_6)

								tmp15385 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

								tmp15386 := Call(__e, PrimNS2Value(symshen_4curry), tmp15385)

								tmp15387 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5type_6)

								tmp15388 := Call(__e, PrimNS2Value(symshen_4demodulate), tmp15387)

								tmp15389 := PrimCons(tmp15388, Nil)

								tmp15390 := PrimCons(sym_h, tmp15389)

								tmp15391 := PrimCons(tmp15386, tmp15390)

								__e.TailApply(PrimNS2Value(symshen_4pair), tmp15384, tmp15391)
								return

							} else {
								__e.TailApply(PrimNS2Value(symfail))
								return
							}

						}, 1)

						tmp15395 := Call(__e, PrimNS2Value(symshen_4_5type_6), NewStream1758)

						__e.TailApply(tmp15382, tmp15395)
						return

					}, 1)

					tmp15396 := Call(__e, PrimNS2Value(symshen_4tlhd), Parse__shen_4_5expr_6)

					tmp15397 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

					tmp15398 := Call(__e, PrimNS2Value(symshen_4pair), tmp15396, tmp15397)

					__e.TailApply(tmp15381, tmp15398)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15408 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V1760)

		tmp15409 := Call(__e, tmp15378, tmp15408)

		__e.TailApply(tmp15366, tmp15409)
		return

	}, 1)

	tmp15410 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formula_6, tmp15365)

	_ = tmp15410

	tmp15411 := MakeNative(func(__e *ControlFlow) {
		V1762 := __e.Get(1)
		_ = V1762
		tmp15412 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			tmp15417 := Call(__e, PrimNS2Value(symfail))

			tmp15418 := PrimEqual(tmp15417, Parse__shen_4_5expr_6)

			tmp15419 := PrimNot(tmp15418)

			if True == tmp15419 {
				tmp15414 := PrimHead(Parse__shen_4_5expr_6)

				tmp15415 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5expr_6)

				tmp15416 := Call(__e, PrimNS2Value(symshen_4curry_1type), tmp15415)

				__e.TailApply(PrimNS2Value(symshen_4pair), tmp15414, tmp15416)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp15420 := Call(__e, PrimNS2Value(symshen_4_5expr_6), V1762)

		__e.TailApply(tmp15412, tmp15420)
		return

	}, 1)

	tmp15421 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5type_6, tmp15411)

	_ = tmp15421

	tmp15422 := MakeNative(func(__e *ControlFlow) {
		V1764 := __e.Get(1)
		_ = V1764
		tmp15432 := PrimHead(V1764)

		tmp15433 := PrimIsPair(tmp15432)

		if True == tmp15433 {
			tmp15424 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15430 := Call(__e, PrimNS2Value(symshen_4doubleunderline_2), Parse__X)

				if True == tmp15430 {
					tmp15426 := Call(__e, PrimNS2Value(symshen_4tlhd), V1764)

					tmp15427 := Call(__e, PrimNS2Value(symshen_4hdtl), V1764)

					tmp15428 := Call(__e, PrimNS2Value(symshen_4pair), tmp15426, tmp15427)

					tmp15429 := PrimHead(tmp15428)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15429, Parse__X)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15431 := Call(__e, PrimNS2Value(symshen_4hdhd), V1764)

			__e.TailApply(tmp15424, tmp15431)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5doubleunderline_6, tmp15422)

	_ = tmp15434

	tmp15435 := MakeNative(func(__e *ControlFlow) {
		V1766 := __e.Get(1)
		_ = V1766
		tmp15445 := PrimHead(V1766)

		tmp15446 := PrimIsPair(tmp15445)

		if True == tmp15446 {
			tmp15437 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp15443 := Call(__e, PrimNS2Value(symshen_4singleunderline_2), Parse__X)

				if True == tmp15443 {
					tmp15439 := Call(__e, PrimNS2Value(symshen_4tlhd), V1766)

					tmp15440 := Call(__e, PrimNS2Value(symshen_4hdtl), V1766)

					tmp15441 := Call(__e, PrimNS2Value(symshen_4pair), tmp15439, tmp15440)

					tmp15442 := PrimHead(tmp15441)

					__e.TailApply(PrimNS2Value(symshen_4pair), tmp15442, Parse__X)
					return

				} else {
					__e.TailApply(PrimNS2Value(symfail))
					return
				}

			}, 1)

			tmp15444 := Call(__e, PrimNS2Value(symshen_4hdhd), V1766)

			__e.TailApply(tmp15437, tmp15444)
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp15447 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleunderline_6, tmp15435)

	_ = tmp15447

	tmp15448 := MakeNative(func(__e *ControlFlow) {
		V1768 := __e.Get(1)
		_ = V1768
		tmp15453 := PrimIsSymbol(V1768)

		if True == tmp15453 {
			tmp15451 := PrimStr(V1768)

			tmp15452 := Call(__e, PrimNS2Value(symshen_4sh_2), tmp15451)

			if True == tmp15452 {
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

	tmp15454 := Call(__e, PrimNS1Value(symns2_1set), symshen_4singleunderline_2, tmp15448)

	_ = tmp15454

	tmp15455 := MakeNative(func(__e *ControlFlow) {
		V1770 := __e.Get(1)
		_ = V1770
		tmp15463 := PrimEqual(MakeString("_"), V1770)

		if True == tmp15463 {
			__e.Return(True)
			return
		} else {
			tmp15461 := PrimPos(V1770, MakeNumber(0))

			tmp15462 := PrimEqual(tmp15461, MakeString("_"))

			if True == tmp15462 {
				tmp15459 := PrimTailString(V1770)

				tmp15460 := Call(__e, PrimNS2Value(symshen_4sh_2), tmp15459)

				if True == tmp15460 {
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

	tmp15464 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sh_2, tmp15455)

	_ = tmp15464

	tmp15465 := MakeNative(func(__e *ControlFlow) {
		V1772 := __e.Get(1)
		_ = V1772
		tmp15470 := PrimIsSymbol(V1772)

		if True == tmp15470 {
			tmp15468 := PrimStr(V1772)

			tmp15469 := Call(__e, PrimNS2Value(symshen_4dh_2), tmp15468)

			if True == tmp15469 {
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

	tmp15471 := Call(__e, PrimNS1Value(symns2_1set), symshen_4doubleunderline_2, tmp15465)

	_ = tmp15471

	tmp15472 := MakeNative(func(__e *ControlFlow) {
		V1774 := __e.Get(1)
		_ = V1774
		tmp15480 := PrimEqual(MakeString("="), V1774)

		if True == tmp15480 {
			__e.Return(True)
			return
		} else {
			tmp15478 := PrimPos(V1774, MakeNumber(0))

			tmp15479 := PrimEqual(tmp15478, MakeString("="))

			if True == tmp15479 {
				tmp15476 := PrimTailString(V1774)

				tmp15477 := Call(__e, PrimNS2Value(symshen_4dh_2), tmp15476)

				if True == tmp15477 {
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

	tmp15481 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dh_2, tmp15472)

	_ = tmp15481

	tmp15482 := MakeNative(func(__e *ControlFlow) {
		V1777 := __e.Get(1)
		_ = V1777
		V1778 := __e.Get(2)
		_ = V1778
		tmp15483 := Call(__e, PrimNS2Value(symshen_4rules_1_6horn_1clauses), V1777, V1778)

		tmp15484 := Call(__e, PrimNS2Value(symshen_4s_1prolog), tmp15483)

		__e.TailApply(PrimNS2Value(symshen_4remember_1datatype), tmp15484)
		return

	}, 2)

	tmp15485 := Call(__e, PrimNS1Value(symns2_1set), symshen_4process_1datatype, tmp15482)

	_ = tmp15485

	tmp15486 := MakeNative(func(__e *ControlFlow) {
		V1784 := __e.Get(1)
		_ = V1784
		tmp15496 := PrimIsPair(V1784)

		if True == tmp15496 {
			tmp15488 := PrimHead(V1784)

			tmp15489 := PrimNS3Value(symshen_4_ddatatypes_d)

			tmp15490 := Call(__e, PrimNS2Value(symadjoin), tmp15488, tmp15489)

			tmp15491 := PrimNS3Set(symshen_4_ddatatypes_d, tmp15490)

			_ = tmp15491

			tmp15492 := PrimHead(V1784)

			tmp15493 := PrimNS3Value(symshen_4_dalldatatypes_d)

			tmp15494 := Call(__e, PrimNS2Value(symadjoin), tmp15492, tmp15493)

			tmp15495 := PrimNS3Set(symshen_4_dalldatatypes_d, tmp15494)

			_ = tmp15495

			__e.Return(PrimHead(V1784))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4remember_1datatype)
			return
		}

	}, 1)

	tmp15497 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remember_1datatype, tmp15486)

	_ = tmp15497

	tmp15498 := MakeNative(func(__e *ControlFlow) {
		V1789 := __e.Get(1)
		_ = V1789
		V1790 := __e.Get(2)
		_ = V1790
		tmp15532 := PrimEqual(Nil, V1790)

		if True == tmp15532 {
			__e.Return(Nil)
			return
		} else {
			tmp15531 := PrimIsPair(V1790)

			var ifres15522 Obj

			if True == tmp15531 {
				tmp15529 := PrimHead(V1790)

				tmp15530 := Call(__e, PrimNS2Value(symtuple_2), tmp15529)

				var ifres15524 Obj

				if True == tmp15530 {
					tmp15526 := PrimHead(V1790)

					tmp15527 := Call(__e, PrimNS2Value(symfst), tmp15526)

					tmp15528 := PrimEqual(symshen_4single, tmp15527)

					var ifres15525 Obj

					if True == tmp15528 {
						ifres15525 = True

					} else {
						ifres15525 = False

					}

					ifres15524 = ifres15525

				} else {
					ifres15524 = False

				}

				var ifres15523 Obj

				if True == ifres15524 {
					ifres15523 = True

				} else {
					ifres15523 = False

				}

				ifres15522 = ifres15523

			} else {
				ifres15522 = False

			}

			if True == ifres15522 {
				tmp15517 := PrimHead(V1790)

				tmp15518 := Call(__e, PrimNS2Value(symsnd), tmp15517)

				tmp15519 := Call(__e, PrimNS2Value(symshen_4rule_1_6horn_1clause), V1789, tmp15518)

				tmp15520 := PrimTail(V1790)

				tmp15521 := Call(__e, PrimNS2Value(symshen_4rules_1_6horn_1clauses), V1789, tmp15520)

				__e.Return(PrimCons(tmp15519, tmp15521))
				return

			} else {
				tmp15516 := PrimIsPair(V1790)

				var ifres15507 Obj

				if True == tmp15516 {
					tmp15514 := PrimHead(V1790)

					tmp15515 := Call(__e, PrimNS2Value(symtuple_2), tmp15514)

					var ifres15509 Obj

					if True == tmp15515 {
						tmp15511 := PrimHead(V1790)

						tmp15512 := Call(__e, PrimNS2Value(symfst), tmp15511)

						tmp15513 := PrimEqual(symshen_4double, tmp15512)

						var ifres15510 Obj

						if True == tmp15513 {
							ifres15510 = True

						} else {
							ifres15510 = False

						}

						ifres15509 = ifres15510

					} else {
						ifres15509 = False

					}

					var ifres15508 Obj

					if True == ifres15509 {
						ifres15508 = True

					} else {
						ifres15508 = False

					}

					ifres15507 = ifres15508

				} else {
					ifres15507 = False

				}

				if True == ifres15507 {
					tmp15502 := PrimHead(V1790)

					tmp15503 := Call(__e, PrimNS2Value(symsnd), tmp15502)

					tmp15504 := Call(__e, PrimNS2Value(symshen_4double_1_6singles), tmp15503)

					tmp15505 := PrimTail(V1790)

					tmp15506 := Call(__e, PrimNS2Value(symappend), tmp15504, tmp15505)

					__e.TailApply(PrimNS2Value(symshen_4rules_1_6horn_1clauses), V1789, tmp15506)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4rules_1_6horn_1clauses)
					return
				}

			}

		}

	}, 2)

	tmp15533 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rules_1_6horn_1clauses, tmp15498)

	_ = tmp15533

	tmp15534 := MakeNative(func(__e *ControlFlow) {
		V1792 := __e.Get(1)
		_ = V1792
		tmp15535 := Call(__e, PrimNS2Value(symshen_4right_1rule), V1792)

		tmp15536 := Call(__e, PrimNS2Value(symshen_4left_1rule), V1792)

		tmp15537 := PrimCons(tmp15536, Nil)

		__e.Return(PrimCons(tmp15535, tmp15537))
		return

	}, 1)

	tmp15538 := Call(__e, PrimNS1Value(symns2_1set), symshen_4double_1_6singles, tmp15534)

	_ = tmp15538

	tmp15539 := MakeNative(func(__e *ControlFlow) {
		V1794 := __e.Get(1)
		_ = V1794
		__e.TailApply(PrimNS2Value(sym_8p), symshen_4single, V1794)
		return
	}, 1)

	tmp15540 := Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1rule, tmp15539)

	_ = tmp15540

	tmp15541 := MakeNative(func(__e *ControlFlow) {
		V1796 := __e.Get(1)
		_ = V1796
		tmp15591 := PrimIsPair(V1796)

		var ifres15563 Obj

		if True == tmp15591 {
			tmp15589 := PrimTail(V1796)

			tmp15590 := PrimIsPair(tmp15589)

			var ifres15565 Obj

			if True == tmp15590 {
				tmp15586 := PrimTail(V1796)

				tmp15587 := PrimTail(tmp15586)

				tmp15588 := PrimIsPair(tmp15587)

				var ifres15567 Obj

				if True == tmp15588 {
					tmp15582 := PrimTail(V1796)

					tmp15583 := PrimTail(tmp15582)

					tmp15584 := PrimHead(tmp15583)

					tmp15585 := Call(__e, PrimNS2Value(symtuple_2), tmp15584)

					var ifres15569 Obj

					if True == tmp15585 {
						tmp15577 := PrimTail(V1796)

						tmp15578 := PrimTail(tmp15577)

						tmp15579 := PrimHead(tmp15578)

						tmp15580 := Call(__e, PrimNS2Value(symfst), tmp15579)

						tmp15581 := PrimEqual(Nil, tmp15580)

						var ifres15571 Obj

						if True == tmp15581 {
							tmp15573 := PrimTail(V1796)

							tmp15574 := PrimTail(tmp15573)

							tmp15575 := PrimTail(tmp15574)

							tmp15576 := PrimEqual(Nil, tmp15575)

							var ifres15572 Obj

							if True == tmp15576 {
								ifres15572 = True

							} else {
								ifres15572 = False

							}

							ifres15571 = ifres15572

						} else {
							ifres15571 = False

						}

						var ifres15570 Obj

						if True == ifres15571 {
							ifres15570 = True

						} else {
							ifres15570 = False

						}

						ifres15569 = ifres15570

					} else {
						ifres15569 = False

					}

					var ifres15568 Obj

					if True == ifres15569 {
						ifres15568 = True

					} else {
						ifres15568 = False

					}

					ifres15567 = ifres15568

				} else {
					ifres15567 = False

				}

				var ifres15566 Obj

				if True == ifres15567 {
					ifres15566 = True

				} else {
					ifres15566 = False

				}

				ifres15565 = ifres15566

			} else {
				ifres15565 = False

			}

			var ifres15564 Obj

			if True == ifres15565 {
				ifres15564 = True

			} else {
				ifres15564 = False

			}

			ifres15563 = ifres15564

		} else {
			ifres15563 = False

		}

		if True == ifres15563 {
			tmp15543 := MakeNative(func(__e *ControlFlow) {
				Q := __e.Get(1)
				_ = Q
				tmp15544 := MakeNative(func(__e *ControlFlow) {
					NewConclusion := __e.Get(1)
					_ = NewConclusion
					tmp15545 := MakeNative(func(__e *ControlFlow) {
						NewPremises := __e.Get(1)
						_ = NewPremises
						tmp15546 := PrimHead(V1796)

						tmp15547 := PrimCons(NewConclusion, Nil)

						tmp15548 := PrimCons(NewPremises, tmp15547)

						tmp15549 := PrimCons(tmp15546, tmp15548)

						__e.TailApply(PrimNS2Value(sym_8p), symshen_4single, tmp15549)
						return

					}, 1)

					tmp15550 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4right_1_6left), X)
						return
					}, 1)

					tmp15551 := PrimTail(V1796)

					tmp15552 := PrimHead(tmp15551)

					tmp15553 := Call(__e, PrimNS2Value(symmap), tmp15550, tmp15552)

					tmp15554 := Call(__e, PrimNS2Value(sym_8p), tmp15553, Q)

					tmp15555 := PrimCons(tmp15554, Nil)

					__e.TailApply(tmp15545, tmp15555)
					return

				}, 1)

				tmp15556 := PrimTail(V1796)

				tmp15557 := PrimTail(tmp15556)

				tmp15558 := PrimHead(tmp15557)

				tmp15559 := Call(__e, PrimNS2Value(symsnd), tmp15558)

				tmp15560 := PrimCons(tmp15559, Nil)

				tmp15561 := Call(__e, PrimNS2Value(sym_8p), tmp15560, Q)

				__e.TailApply(tmp15544, tmp15561)
				return

			}, 1)

			tmp15562 := Call(__e, PrimNS2Value(symgensym), symQv)

			__e.TailApply(tmp15543, tmp15562)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4left_1rule)
			return
		}

	}, 1)

	tmp15592 := Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1rule, tmp15541)

	_ = tmp15592

	tmp15593 := MakeNative(func(__e *ControlFlow) {
		V1802 := __e.Get(1)
		_ = V1802
		tmp15599 := Call(__e, PrimNS2Value(symtuple_2), V1802)

		var ifres15595 Obj

		if True == tmp15599 {
			tmp15597 := Call(__e, PrimNS2Value(symfst), V1802)

			tmp15598 := PrimEqual(Nil, tmp15597)

			var ifres15596 Obj

			if True == tmp15598 {
				ifres15596 = True

			} else {
				ifres15596 = False

			}

			ifres15595 = ifres15596

		} else {
			ifres15595 = False

		}

		if True == ifres15595 {
			__e.TailApply(PrimNS2Value(symsnd), V1802)
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("syntax error with ==========\n")))
			return
		}

	}, 1)

	tmp15600 := Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1_6left, tmp15593)

	_ = tmp15600

	tmp15601 := MakeNative(func(__e *ControlFlow) {
		V1805 := __e.Get(1)
		_ = V1805
		V1806 := __e.Get(2)
		_ = V1806
		tmp15639 := PrimIsPair(V1806)

		var ifres15618 Obj

		if True == tmp15639 {
			tmp15637 := PrimTail(V1806)

			tmp15638 := PrimIsPair(tmp15637)

			var ifres15620 Obj

			if True == tmp15638 {
				tmp15634 := PrimTail(V1806)

				tmp15635 := PrimTail(tmp15634)

				tmp15636 := PrimIsPair(tmp15635)

				var ifres15622 Obj

				if True == tmp15636 {
					tmp15630 := PrimTail(V1806)

					tmp15631 := PrimTail(tmp15630)

					tmp15632 := PrimHead(tmp15631)

					tmp15633 := Call(__e, PrimNS2Value(symtuple_2), tmp15632)

					var ifres15624 Obj

					if True == tmp15633 {
						tmp15626 := PrimTail(V1806)

						tmp15627 := PrimTail(tmp15626)

						tmp15628 := PrimTail(tmp15627)

						tmp15629 := PrimEqual(Nil, tmp15628)

						var ifres15625 Obj

						if True == tmp15629 {
							ifres15625 = True

						} else {
							ifres15625 = False

						}

						ifres15624 = ifres15625

					} else {
						ifres15624 = False

					}

					var ifres15623 Obj

					if True == ifres15624 {
						ifres15623 = True

					} else {
						ifres15623 = False

					}

					ifres15622 = ifres15623

				} else {
					ifres15622 = False

				}

				var ifres15621 Obj

				if True == ifres15622 {
					ifres15621 = True

				} else {
					ifres15621 = False

				}

				ifres15620 = ifres15621

			} else {
				ifres15620 = False

			}

			var ifres15619 Obj

			if True == ifres15620 {
				ifres15619 = True

			} else {
				ifres15619 = False

			}

			ifres15618 = ifres15619

		} else {
			ifres15618 = False

		}

		if True == ifres15618 {
			tmp15603 := PrimTail(V1806)

			tmp15604 := PrimTail(tmp15603)

			tmp15605 := PrimHead(tmp15604)

			tmp15606 := Call(__e, PrimNS2Value(symsnd), tmp15605)

			tmp15607 := Call(__e, PrimNS2Value(symshen_4rule_1_6horn_1clause_1head), V1805, tmp15606)

			tmp15608 := PrimHead(V1806)

			tmp15609 := PrimTail(V1806)

			tmp15610 := PrimHead(tmp15609)

			tmp15611 := PrimTail(V1806)

			tmp15612 := PrimTail(tmp15611)

			tmp15613 := PrimHead(tmp15612)

			tmp15614 := Call(__e, PrimNS2Value(symfst), tmp15613)

			tmp15615 := Call(__e, PrimNS2Value(symshen_4rule_1_6horn_1clause_1body), tmp15608, tmp15610, tmp15614)

			tmp15616 := PrimCons(tmp15615, Nil)

			tmp15617 := PrimCons(sym_h_1, tmp15616)

			__e.Return(PrimCons(tmp15607, tmp15617))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4rule_1_6horn_1clause)
			return
		}

	}, 2)

	tmp15640 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause, tmp15601)

	_ = tmp15640

	tmp15641 := MakeNative(func(__e *ControlFlow) {
		V1809 := __e.Get(1)
		_ = V1809
		V1810 := __e.Get(2)
		_ = V1810
		tmp15642 := Call(__e, PrimNS2Value(symshen_4mode_1ify), V1810)

		tmp15643 := PrimCons(symContext__1957, Nil)

		tmp15644 := PrimCons(tmp15642, tmp15643)

		__e.Return(PrimCons(V1809, tmp15644))
		return

	}, 2)

	tmp15645 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1head, tmp15641)

	_ = tmp15645

	tmp15646 := MakeNative(func(__e *ControlFlow) {
		V1812 := __e.Get(1)
		_ = V1812
		tmp15680 := PrimIsPair(V1812)

		var ifres15660 Obj

		if True == tmp15680 {
			tmp15678 := PrimTail(V1812)

			tmp15679 := PrimIsPair(tmp15678)

			var ifres15662 Obj

			if True == tmp15679 {
				tmp15675 := PrimTail(V1812)

				tmp15676 := PrimHead(tmp15675)

				tmp15677 := PrimEqual(sym_h, tmp15676)

				var ifres15664 Obj

				if True == tmp15677 {
					tmp15672 := PrimTail(V1812)

					tmp15673 := PrimTail(tmp15672)

					tmp15674 := PrimIsPair(tmp15673)

					var ifres15666 Obj

					if True == tmp15674 {
						tmp15668 := PrimTail(V1812)

						tmp15669 := PrimTail(tmp15668)

						tmp15670 := PrimTail(tmp15669)

						tmp15671 := PrimEqual(Nil, tmp15670)

						var ifres15667 Obj

						if True == tmp15671 {
							ifres15667 = True

						} else {
							ifres15667 = False

						}

						ifres15666 = ifres15667

					} else {
						ifres15666 = False

					}

					var ifres15665 Obj

					if True == ifres15666 {
						ifres15665 = True

					} else {
						ifres15665 = False

					}

					ifres15664 = ifres15665

				} else {
					ifres15664 = False

				}

				var ifres15663 Obj

				if True == ifres15664 {
					ifres15663 = True

				} else {
					ifres15663 = False

				}

				ifres15662 = ifres15663

			} else {
				ifres15662 = False

			}

			var ifres15661 Obj

			if True == ifres15662 {
				ifres15661 = True

			} else {
				ifres15661 = False

			}

			ifres15660 = ifres15661

		} else {
			ifres15660 = False

		}

		if True == ifres15660 {
			tmp15648 := PrimHead(V1812)

			tmp15649 := PrimTail(V1812)

			tmp15650 := PrimTail(tmp15649)

			tmp15651 := PrimHead(tmp15650)

			tmp15652 := PrimCons(sym_7, Nil)

			tmp15653 := PrimCons(tmp15651, tmp15652)

			tmp15654 := PrimCons(symmode, tmp15653)

			tmp15655 := PrimCons(tmp15654, Nil)

			tmp15656 := PrimCons(sym_h, tmp15655)

			tmp15657 := PrimCons(tmp15648, tmp15656)

			tmp15658 := PrimCons(sym_1, Nil)

			tmp15659 := PrimCons(tmp15657, tmp15658)

			__e.Return(PrimCons(symmode, tmp15659))
			return

		} else {
			__e.Return(V1812)
			return
		}

	}, 1)

	tmp15681 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mode_1ify, tmp15646)

	_ = tmp15681

	tmp15682 := MakeNative(func(__e *ControlFlow) {
		V1816 := __e.Get(1)
		_ = V1816
		V1817 := __e.Get(2)
		_ = V1817
		V1818 := __e.Get(3)
		_ = V1818
		tmp15683 := MakeNative(func(__e *ControlFlow) {
			Variables := __e.Get(1)
			_ = Variables
			tmp15684 := MakeNative(func(__e *ControlFlow) {
				Predicates := __e.Get(1)
				_ = Predicates
				tmp15685 := MakeNative(func(__e *ControlFlow) {
					SearchLiterals := __e.Get(1)
					_ = SearchLiterals
					tmp15686 := MakeNative(func(__e *ControlFlow) {
						SearchClauses := __e.Get(1)
						_ = SearchClauses
						tmp15687 := MakeNative(func(__e *ControlFlow) {
							SideLiterals := __e.Get(1)
							_ = SideLiterals
							tmp15688 := MakeNative(func(__e *ControlFlow) {
								PremissLiterals := __e.Get(1)
								_ = PremissLiterals
								tmp15689 := Call(__e, PrimNS2Value(symappend), SideLiterals, PremissLiterals)

								__e.TailApply(PrimNS2Value(symappend), SearchLiterals, tmp15689)
								return

							}, 1)

							tmp15690 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp15691 := Call(__e, PrimNS2Value(symempty_2), V1818)

								__e.TailApply(PrimNS2Value(symshen_4construct_1premiss_1literal), X, tmp15691)
								return

							}, 1)

							tmp15692 := Call(__e, PrimNS2Value(symmap), tmp15690, V1817)

							__e.TailApply(tmp15688, tmp15692)
							return

						}, 1)

						tmp15693 := Call(__e, PrimNS2Value(symshen_4construct_1side_1literals), V1816)

						__e.TailApply(tmp15687, tmp15693)
						return

					}, 1)

					tmp15694 := Call(__e, PrimNS2Value(symshen_4construct_1search_1clauses), Predicates, V1818, Variables)

					__e.TailApply(tmp15686, tmp15694)
					return

				}, 1)

				tmp15695 := Call(__e, PrimNS2Value(symshen_4construct_1search_1literals), Predicates, Variables, symContext__1957, symContext1__1957)

				__e.TailApply(tmp15685, tmp15695)
				return

			}, 1)

			tmp15696 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symgensym), symshen_4cl)
				return
			}, 1)

			tmp15697 := Call(__e, PrimNS2Value(symmap), tmp15696, V1818)

			__e.TailApply(tmp15684, tmp15697)
			return

		}, 1)

		tmp15698 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4extract__vars), X)
			return
		}, 1)

		tmp15699 := Call(__e, PrimNS2Value(symmap), tmp15698, V1818)

		__e.TailApply(tmp15683, tmp15699)
		return

	}, 3)

	tmp15700 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1body, tmp15682)

	_ = tmp15700

	tmp15701 := MakeNative(func(__e *ControlFlow) {
		V1827 := __e.Get(1)
		_ = V1827
		V1828 := __e.Get(2)
		_ = V1828
		V1829 := __e.Get(3)
		_ = V1829
		V1830 := __e.Get(4)
		_ = V1830
		tmp15706 := PrimEqual(Nil, V1827)

		var ifres15703 Obj

		if True == tmp15706 {
			tmp15705 := PrimEqual(Nil, V1828)

			var ifres15704 Obj

			if True == tmp15705 {
				ifres15704 = True

			} else {
				ifres15704 = False

			}

			ifres15703 = ifres15704

		} else {
			ifres15703 = False

		}

		if True == ifres15703 {
			__e.Return(Nil)
			return
		} else {
			__e.TailApply(PrimNS2Value(symshen_4csl_1help), V1827, V1828, V1829, V1830)
			return
		}

	}, 4)

	tmp15707 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1literals, tmp15701)

	_ = tmp15707

	tmp15708 := MakeNative(func(__e *ControlFlow) {
		V1837 := __e.Get(1)
		_ = V1837
		V1838 := __e.Get(2)
		_ = V1838
		V1839 := __e.Get(3)
		_ = V1839
		V1840 := __e.Get(4)
		_ = V1840
		tmp15730 := PrimEqual(Nil, V1837)

		var ifres15727 Obj

		if True == tmp15730 {
			tmp15729 := PrimEqual(Nil, V1838)

			var ifres15728 Obj

			if True == tmp15729 {
				ifres15728 = True

			} else {
				ifres15728 = False

			}

			ifres15727 = ifres15728

		} else {
			ifres15727 = False

		}

		if True == ifres15727 {
			tmp15724 := PrimCons(V1839, Nil)

			tmp15725 := PrimCons(symContextOut__1957, tmp15724)

			tmp15726 := PrimCons(symbind, tmp15725)

			__e.Return(PrimCons(tmp15726, Nil))
			return

		} else {
			tmp15723 := PrimIsPair(V1837)

			var ifres15720 Obj

			if True == tmp15723 {
				tmp15722 := PrimIsPair(V1838)

				var ifres15721 Obj

				if True == tmp15722 {
					ifres15721 = True

				} else {
					ifres15721 = False

				}

				ifres15720 = ifres15721

			} else {
				ifres15720 = False

			}

			if True == ifres15720 {
				tmp15711 := PrimHead(V1837)

				tmp15712 := PrimHead(V1838)

				tmp15713 := PrimCons(V1840, tmp15712)

				tmp15714 := PrimCons(V1839, tmp15713)

				tmp15715 := PrimCons(tmp15711, tmp15714)

				tmp15716 := PrimTail(V1837)

				tmp15717 := PrimTail(V1838)

				tmp15718 := Call(__e, PrimNS2Value(symgensym), symContext)

				tmp15719 := Call(__e, PrimNS2Value(symshen_4csl_1help), tmp15716, tmp15717, V1840, tmp15718)

				__e.Return(PrimCons(tmp15715, tmp15719))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4csl_1help)
				return
			}

		}

	}, 4)

	tmp15731 := Call(__e, PrimNS1Value(symns2_1set), symshen_4csl_1help, tmp15708)

	_ = tmp15731

	tmp15732 := MakeNative(func(__e *ControlFlow) {
		V1844 := __e.Get(1)
		_ = V1844
		V1845 := __e.Get(2)
		_ = V1845
		V1846 := __e.Get(3)
		_ = V1846
		tmp15755 := PrimEqual(Nil, V1844)

		var ifres15749 Obj

		if True == tmp15755 {
			tmp15754 := PrimEqual(Nil, V1845)

			var ifres15751 Obj

			if True == tmp15754 {
				tmp15753 := PrimEqual(Nil, V1846)

				var ifres15752 Obj

				if True == tmp15753 {
					ifres15752 = True

				} else {
					ifres15752 = False

				}

				ifres15751 = ifres15752

			} else {
				ifres15751 = False

			}

			var ifres15750 Obj

			if True == ifres15751 {
				ifres15750 = True

			} else {
				ifres15750 = False

			}

			ifres15749 = ifres15750

		} else {
			ifres15749 = False

		}

		if True == ifres15749 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp15748 := PrimIsPair(V1844)

			var ifres15742 Obj

			if True == tmp15748 {
				tmp15747 := PrimIsPair(V1845)

				var ifres15744 Obj

				if True == tmp15747 {
					tmp15746 := PrimIsPair(V1846)

					var ifres15745 Obj

					if True == tmp15746 {
						ifres15745 = True

					} else {
						ifres15745 = False

					}

					ifres15744 = ifres15745

				} else {
					ifres15744 = False

				}

				var ifres15743 Obj

				if True == ifres15744 {
					ifres15743 = True

				} else {
					ifres15743 = False

				}

				ifres15742 = ifres15743

			} else {
				ifres15742 = False

			}

			if True == ifres15742 {
				tmp15735 := PrimHead(V1844)

				tmp15736 := PrimHead(V1845)

				tmp15737 := PrimHead(V1846)

				tmp15738 := Call(__e, PrimNS2Value(symshen_4construct_1search_1clause), tmp15735, tmp15736, tmp15737)

				_ = tmp15738

				tmp15739 := PrimTail(V1844)

				tmp15740 := PrimTail(V1845)

				tmp15741 := PrimTail(V1846)

				__e.TailApply(PrimNS2Value(symshen_4construct_1search_1clauses), tmp15739, tmp15740, tmp15741)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4construct_1search_1clauses)
				return
			}

		}

	}, 3)

	tmp15756 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clauses, tmp15732)

	_ = tmp15756

	tmp15757 := MakeNative(func(__e *ControlFlow) {
		V1850 := __e.Get(1)
		_ = V1850
		V1851 := __e.Get(2)
		_ = V1851
		V1852 := __e.Get(3)
		_ = V1852
		tmp15758 := Call(__e, PrimNS2Value(symshen_4construct_1base_1search_1clause), V1850, V1851, V1852)

		tmp15759 := Call(__e, PrimNS2Value(symshen_4construct_1recursive_1search_1clause), V1850, V1851, V1852)

		tmp15760 := PrimCons(tmp15759, Nil)

		tmp15761 := PrimCons(tmp15758, tmp15760)

		__e.TailApply(PrimNS2Value(symshen_4s_1prolog), tmp15761)
		return

	}, 3)

	tmp15762 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clause, tmp15757)

	_ = tmp15762

	tmp15763 := MakeNative(func(__e *ControlFlow) {
		V1856 := __e.Get(1)
		_ = V1856
		V1857 := __e.Get(2)
		_ = V1857
		V1858 := __e.Get(3)
		_ = V1858
		tmp15764 := Call(__e, PrimNS2Value(symshen_4mode_1ify), V1857)

		tmp15765 := PrimCons(tmp15764, symIn__1957)

		tmp15766 := PrimCons(symIn__1957, V1858)

		tmp15767 := PrimCons(tmp15765, tmp15766)

		tmp15768 := PrimCons(V1856, tmp15767)

		tmp15769 := PrimCons(Nil, Nil)

		tmp15770 := PrimCons(sym_h_1, tmp15769)

		__e.Return(PrimCons(tmp15768, tmp15770))
		return

	}, 3)

	tmp15771 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1base_1search_1clause, tmp15763)

	_ = tmp15771

	tmp15772 := MakeNative(func(__e *ControlFlow) {
		V1862 := __e.Get(1)
		_ = V1862
		V1863 := __e.Get(2)
		_ = V1863
		V1864 := __e.Get(3)
		_ = V1864
		tmp15773 := PrimCons(symAssumption__1957, symAssumptions__1957)

		tmp15774 := PrimCons(symAssumption__1957, symOut__1957)

		tmp15775 := PrimCons(tmp15774, V1864)

		tmp15776 := PrimCons(tmp15773, tmp15775)

		tmp15777 := PrimCons(V1862, tmp15776)

		tmp15778 := PrimCons(symOut__1957, V1864)

		tmp15779 := PrimCons(symAssumptions__1957, tmp15778)

		tmp15780 := PrimCons(V1862, tmp15779)

		tmp15781 := PrimCons(tmp15780, Nil)

		tmp15782 := PrimCons(tmp15781, Nil)

		tmp15783 := PrimCons(sym_h_1, tmp15782)

		__e.Return(PrimCons(tmp15777, tmp15783))
		return

	}, 3)

	tmp15784 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1recursive_1search_1clause, tmp15772)

	_ = tmp15784

	tmp15785 := MakeNative(func(__e *ControlFlow) {
		V1870 := __e.Get(1)
		_ = V1870
		tmp15851 := PrimEqual(Nil, V1870)

		if True == tmp15851 {
			__e.Return(Nil)
			return
		} else {
			tmp15850 := PrimIsPair(V1870)

			var ifres15830 Obj

			if True == tmp15850 {
				tmp15848 := PrimHead(V1870)

				tmp15849 := PrimIsPair(tmp15848)

				var ifres15832 Obj

				if True == tmp15849 {
					tmp15845 := PrimHead(V1870)

					tmp15846 := PrimHead(tmp15845)

					tmp15847 := PrimEqual(symif, tmp15846)

					var ifres15834 Obj

					if True == tmp15847 {
						tmp15842 := PrimHead(V1870)

						tmp15843 := PrimTail(tmp15842)

						tmp15844 := PrimIsPair(tmp15843)

						var ifres15836 Obj

						if True == tmp15844 {
							tmp15838 := PrimHead(V1870)

							tmp15839 := PrimTail(tmp15838)

							tmp15840 := PrimTail(tmp15839)

							tmp15841 := PrimEqual(Nil, tmp15840)

							var ifres15837 Obj

							if True == tmp15841 {
								ifres15837 = True

							} else {
								ifres15837 = False

							}

							ifres15836 = ifres15837

						} else {
							ifres15836 = False

						}

						var ifres15835 Obj

						if True == ifres15836 {
							ifres15835 = True

						} else {
							ifres15835 = False

						}

						ifres15834 = ifres15835

					} else {
						ifres15834 = False

					}

					var ifres15833 Obj

					if True == ifres15834 {
						ifres15833 = True

					} else {
						ifres15833 = False

					}

					ifres15832 = ifres15833

				} else {
					ifres15832 = False

				}

				var ifres15831 Obj

				if True == ifres15832 {
					ifres15831 = True

				} else {
					ifres15831 = False

				}

				ifres15830 = ifres15831

			} else {
				ifres15830 = False

			}

			if True == ifres15830 {
				tmp15825 := PrimHead(V1870)

				tmp15826 := PrimTail(tmp15825)

				tmp15827 := PrimCons(symwhen, tmp15826)

				tmp15828 := PrimTail(V1870)

				tmp15829 := Call(__e, PrimNS2Value(symshen_4construct_1side_1literals), tmp15828)

				__e.Return(PrimCons(tmp15827, tmp15829))
				return

			} else {
				tmp15824 := PrimIsPair(V1870)

				var ifres15797 Obj

				if True == tmp15824 {
					tmp15822 := PrimHead(V1870)

					tmp15823 := PrimIsPair(tmp15822)

					var ifres15799 Obj

					if True == tmp15823 {
						tmp15819 := PrimHead(V1870)

						tmp15820 := PrimHead(tmp15819)

						tmp15821 := PrimEqual(symlet, tmp15820)

						var ifres15801 Obj

						if True == tmp15821 {
							tmp15816 := PrimHead(V1870)

							tmp15817 := PrimTail(tmp15816)

							tmp15818 := PrimIsPair(tmp15817)

							var ifres15803 Obj

							if True == tmp15818 {
								tmp15812 := PrimHead(V1870)

								tmp15813 := PrimTail(tmp15812)

								tmp15814 := PrimTail(tmp15813)

								tmp15815 := PrimIsPair(tmp15814)

								var ifres15805 Obj

								if True == tmp15815 {
									tmp15807 := PrimHead(V1870)

									tmp15808 := PrimTail(tmp15807)

									tmp15809 := PrimTail(tmp15808)

									tmp15810 := PrimTail(tmp15809)

									tmp15811 := PrimEqual(Nil, tmp15810)

									var ifres15806 Obj

									if True == tmp15811 {
										ifres15806 = True

									} else {
										ifres15806 = False

									}

									ifres15805 = ifres15806

								} else {
									ifres15805 = False

								}

								var ifres15804 Obj

								if True == ifres15805 {
									ifres15804 = True

								} else {
									ifres15804 = False

								}

								ifres15803 = ifres15804

							} else {
								ifres15803 = False

							}

							var ifres15802 Obj

							if True == ifres15803 {
								ifres15802 = True

							} else {
								ifres15802 = False

							}

							ifres15801 = ifres15802

						} else {
							ifres15801 = False

						}

						var ifres15800 Obj

						if True == ifres15801 {
							ifres15800 = True

						} else {
							ifres15800 = False

						}

						ifres15799 = ifres15800

					} else {
						ifres15799 = False

					}

					var ifres15798 Obj

					if True == ifres15799 {
						ifres15798 = True

					} else {
						ifres15798 = False

					}

					ifres15797 = ifres15798

				} else {
					ifres15797 = False

				}

				if True == ifres15797 {
					tmp15792 := PrimHead(V1870)

					tmp15793 := PrimTail(tmp15792)

					tmp15794 := PrimCons(symis, tmp15793)

					tmp15795 := PrimTail(V1870)

					tmp15796 := Call(__e, PrimNS2Value(symshen_4construct_1side_1literals), tmp15795)

					__e.Return(PrimCons(tmp15794, tmp15796))
					return

				} else {
					tmp15791 := PrimIsPair(V1870)

					if True == tmp15791 {
						tmp15790 := PrimTail(V1870)

						__e.TailApply(PrimNS2Value(symshen_4construct_1side_1literals), tmp15790)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4construct_1side_1literals)
						return
					}

				}

			}

		}

	}, 1)

	tmp15852 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1side_1literals, tmp15785)

	_ = tmp15852

	tmp15853 := MakeNative(func(__e *ControlFlow) {
		V1877 := __e.Get(1)
		_ = V1877
		V1878 := __e.Get(2)
		_ = V1878
		tmp15864 := Call(__e, PrimNS2Value(symtuple_2), V1877)

		if True == tmp15864 {
			tmp15858 := Call(__e, PrimNS2Value(symsnd), V1877)

			tmp15859 := Call(__e, PrimNS2Value(symshen_4recursive__cons__form), tmp15858)

			tmp15860 := Call(__e, PrimNS2Value(symfst), V1877)

			tmp15861 := Call(__e, PrimNS2Value(symshen_4construct_1context), V1878, tmp15860)

			tmp15862 := PrimCons(tmp15861, Nil)

			tmp15863 := PrimCons(tmp15859, tmp15862)

			__e.Return(PrimCons(symshen_4t_d, tmp15863))
			return

		} else {
			tmp15857 := PrimEqual(sym_b, V1877)

			if True == tmp15857 {
				tmp15856 := PrimCons(symThrowcontrol, Nil)

				__e.Return(PrimCons(symcut, tmp15856))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4construct_1premiss_1literal)
				return
			}

		}

	}, 2)

	tmp15865 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1premiss_1literal, tmp15853)

	_ = tmp15865

	tmp15866 := MakeNative(func(__e *ControlFlow) {
		V1881 := __e.Get(1)
		_ = V1881
		V1882 := __e.Get(2)
		_ = V1882
		tmp15884 := PrimEqual(True, V1881)

		var ifres15881 Obj

		if True == tmp15884 {
			tmp15883 := PrimEqual(Nil, V1882)

			var ifres15882 Obj

			if True == tmp15883 {
				ifres15882 = True

			} else {
				ifres15882 = False

			}

			ifres15881 = ifres15882

		} else {
			ifres15881 = False

		}

		if True == ifres15881 {
			__e.Return(symContext__1957)
			return
		} else {
			tmp15880 := PrimEqual(False, V1881)

			var ifres15877 Obj

			if True == tmp15880 {
				tmp15879 := PrimEqual(Nil, V1882)

				var ifres15878 Obj

				if True == tmp15879 {
					ifres15878 = True

				} else {
					ifres15878 = False

				}

				ifres15877 = ifres15878

			} else {
				ifres15877 = False

			}

			if True == ifres15877 {
				__e.Return(symContextOut__1957)
				return
			} else {
				tmp15876 := PrimIsPair(V1882)

				if True == tmp15876 {
					tmp15870 := PrimHead(V1882)

					tmp15871 := Call(__e, PrimNS2Value(symshen_4recursive__cons__form), tmp15870)

					tmp15872 := PrimTail(V1882)

					tmp15873 := Call(__e, PrimNS2Value(symshen_4construct_1context), V1881, tmp15872)

					tmp15874 := PrimCons(tmp15873, Nil)

					tmp15875 := PrimCons(tmp15871, tmp15874)

					__e.Return(PrimCons(symcons, tmp15875))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4construct_1context)
					return
				}

			}

		}

	}, 2)

	tmp15885 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1context, tmp15866)

	_ = tmp15885

	tmp15886 := MakeNative(func(__e *ControlFlow) {
		V1884 := __e.Get(1)
		_ = V1884
		tmp15894 := PrimIsPair(V1884)

		if True == tmp15894 {
			tmp15888 := PrimHead(V1884)

			tmp15889 := Call(__e, PrimNS2Value(symshen_4recursive__cons__form), tmp15888)

			tmp15890 := PrimTail(V1884)

			tmp15891 := Call(__e, PrimNS2Value(symshen_4recursive__cons__form), tmp15890)

			tmp15892 := PrimCons(tmp15891, Nil)

			tmp15893 := PrimCons(tmp15889, tmp15892)

			__e.Return(PrimCons(symcons, tmp15893))
			return

		} else {
			__e.Return(V1884)
			return
		}

	}, 1)

	tmp15895 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__cons__form, tmp15886)

	_ = tmp15895

	tmp15896 := MakeNative(func(__e *ControlFlow) {
		V1886 := __e.Get(1)
		_ = V1886
		tmp15897 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp15898 := Call(__e, PrimNS2Value(symmap), tmp15897, V1886)

		__e.TailApply(PrimNS2Value(symshen_4preclude_1h), tmp15898)
		return

	}, 1)

	tmp15899 := Call(__e, PrimNS1Value(symns2_1set), sympreclude, tmp15896)

	_ = tmp15899

	tmp15900 := MakeNative(func(__e *ControlFlow) {
		V1888 := __e.Get(1)
		_ = V1888
		tmp15901 := MakeNative(func(__e *ControlFlow) {
			FilterDatatypes := __e.Get(1)
			_ = FilterDatatypes
			__e.Return(PrimNS3Value(symshen_4_ddatatypes_d))
			return
		}, 1)

		tmp15902 := PrimNS3Value(symshen_4_ddatatypes_d)

		tmp15903 := Call(__e, PrimNS2Value(symdifference), tmp15902, V1888)

		tmp15904 := PrimNS3Set(symshen_4_ddatatypes_d, tmp15903)

		__e.TailApply(tmp15901, tmp15904)
		return

	}, 1)

	tmp15905 := Call(__e, PrimNS1Value(symns2_1set), symshen_4preclude_1h, tmp15900)

	_ = tmp15905

	tmp15906 := MakeNative(func(__e *ControlFlow) {
		V1890 := __e.Get(1)
		_ = V1890
		tmp15907 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp15908 := Call(__e, PrimNS2Value(symmap), tmp15907, V1890)

		__e.TailApply(PrimNS2Value(symshen_4include_1h), tmp15908)
		return

	}, 1)

	tmp15909 := Call(__e, PrimNS1Value(symns2_1set), syminclude, tmp15906)

	_ = tmp15909

	tmp15910 := MakeNative(func(__e *ControlFlow) {
		V1892 := __e.Get(1)
		_ = V1892
		tmp15911 := MakeNative(func(__e *ControlFlow) {
			ValidTypes := __e.Get(1)
			_ = ValidTypes
			tmp15912 := MakeNative(func(__e *ControlFlow) {
				NewDatatypes := __e.Get(1)
				_ = NewDatatypes
				__e.Return(PrimNS3Value(symshen_4_ddatatypes_d))
				return
			}, 1)

			tmp15913 := PrimNS3Value(symshen_4_ddatatypes_d)

			tmp15914 := Call(__e, PrimNS2Value(symunion), ValidTypes, tmp15913)

			tmp15915 := PrimNS3Set(symshen_4_ddatatypes_d, tmp15914)

			__e.TailApply(tmp15912, tmp15915)
			return

		}, 1)

		tmp15916 := PrimNS3Value(symshen_4_dalldatatypes_d)

		tmp15917 := Call(__e, PrimNS2Value(symintersection), V1892, tmp15916)

		__e.TailApply(tmp15911, tmp15917)
		return

	}, 1)

	tmp15918 := Call(__e, PrimNS1Value(symns2_1set), symshen_4include_1h, tmp15910)

	_ = tmp15918

	tmp15919 := MakeNative(func(__e *ControlFlow) {
		V1894 := __e.Get(1)
		_ = V1894
		tmp15920 := PrimNS3Value(symshen_4_dalldatatypes_d)

		tmp15921 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp15922 := Call(__e, PrimNS2Value(symmap), tmp15921, V1894)

		tmp15923 := Call(__e, PrimNS2Value(symdifference), tmp15920, tmp15922)

		__e.TailApply(PrimNS2Value(symshen_4preclude_1h), tmp15923)
		return

	}, 1)

	tmp15924 := Call(__e, PrimNS1Value(symns2_1set), sympreclude_1all_1but, tmp15919)

	_ = tmp15924

	tmp15925 := MakeNative(func(__e *ControlFlow) {
		V1896 := __e.Get(1)
		_ = V1896
		tmp15926 := PrimNS3Value(symshen_4_dalldatatypes_d)

		tmp15927 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4intern_1type), X)
			return
		}, 1)

		tmp15928 := Call(__e, PrimNS2Value(symmap), tmp15927, V1896)

		tmp15929 := Call(__e, PrimNS2Value(symdifference), tmp15926, tmp15928)

		__e.TailApply(PrimNS2Value(symshen_4include_1h), tmp15929)
		return

	}, 1)

	tmp15930 := Call(__e, PrimNS1Value(symns2_1set), syminclude_1all_1but, tmp15925)

	_ = tmp15930

	tmp15931 := MakeNative(func(__e *ControlFlow) {
		V1902 := __e.Get(1)
		_ = V1902
		tmp15962 := PrimEqual(Nil, V1902)

		if True == tmp15962 {
			tmp15958 := PrimNS3Value(symshen_4_dtc_d)

			tmp15959 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4demod_1rule), X)
				return
			}, 1)

			tmp15960 := PrimNS3Value(symshen_4_dsynonyms_d)

			tmp15961 := Call(__e, PrimNS2Value(symmapcan), tmp15959, tmp15960)

			__e.TailApply(PrimNS2Value(symshen_4update_1demodulation_1function), tmp15958, tmp15961)
			return

		} else {
			tmp15957 := PrimIsPair(V1902)

			var ifres15953 Obj

			if True == tmp15957 {
				tmp15955 := PrimTail(V1902)

				tmp15956 := PrimIsPair(tmp15955)

				var ifres15954 Obj

				if True == tmp15956 {
					ifres15954 = True

				} else {
					ifres15954 = False

				}

				ifres15953 = ifres15954

			} else {
				ifres15953 = False

			}

			if True == ifres15953 {
				tmp15934 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					tmp15946 := Call(__e, PrimNS2Value(symempty_2), Vs)

					if True == tmp15946 {
						tmp15938 := PrimHead(V1902)

						tmp15939 := PrimTail(V1902)

						tmp15940 := PrimHead(tmp15939)

						tmp15941 := PrimCons(tmp15940, Nil)

						tmp15942 := PrimCons(tmp15938, tmp15941)

						tmp15943 := Call(__e, PrimNS2Value(symshen_4pushnew), tmp15942, symshen_4_dsynonyms_d)

						_ = tmp15943

						tmp15944 := PrimTail(V1902)

						tmp15945 := PrimTail(tmp15944)

						__e.TailApply(PrimNS2Value(symshen_4synonyms_1help), tmp15945)
						return

					} else {
						tmp15936 := PrimTail(V1902)

						tmp15937 := PrimHead(tmp15936)

						__e.TailApply(PrimNS2Value(symshen_4free__variable__warnings), tmp15937, Vs)
						return

					}

				}, 1)

				tmp15947 := PrimTail(V1902)

				tmp15948 := PrimHead(tmp15947)

				tmp15949 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp15948)

				tmp15950 := PrimHead(V1902)

				tmp15951 := Call(__e, PrimNS2Value(symshen_4extract__vars), tmp15950)

				tmp15952 := Call(__e, PrimNS2Value(symdifference), tmp15949, tmp15951)

				__e.TailApply(tmp15934, tmp15952)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("odd number of synonyms\n")))
				return
			}

		}

	}, 1)

	tmp15963 := Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1help, tmp15931)

	_ = tmp15963

	tmp15964 := MakeNative(func(__e *ControlFlow) {
		V1905 := __e.Get(1)
		_ = V1905
		V1906 := __e.Get(2)
		_ = V1906
		tmp15968 := PrimNS3Value(V1906)

		tmp15969 := Call(__e, PrimNS2Value(symelement_2), V1905, tmp15968)

		if True == tmp15969 {
			__e.Return(PrimNS3Value(V1906))
			return
		} else {
			tmp15966 := PrimNS3Value(V1906)

			tmp15967 := PrimCons(V1905, tmp15966)

			__e.Return(PrimNS3Set(V1906, tmp15967))
			return

		}

	}, 2)

	tmp15970 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pushnew, tmp15964)

	_ = tmp15970

	tmp15971 := MakeNative(func(__e *ControlFlow) {
		V1908 := __e.Get(1)
		_ = V1908
		tmp15989 := PrimIsPair(V1908)

		var ifres15980 Obj

		if True == tmp15989 {
			tmp15987 := PrimTail(V1908)

			tmp15988 := PrimIsPair(tmp15987)

			var ifres15982 Obj

			if True == tmp15988 {
				tmp15984 := PrimTail(V1908)

				tmp15985 := PrimTail(tmp15984)

				tmp15986 := PrimEqual(Nil, tmp15985)

				var ifres15983 Obj

				if True == tmp15986 {
					ifres15983 = True

				} else {
					ifres15983 = False

				}

				ifres15982 = ifres15983

			} else {
				ifres15982 = False

			}

			var ifres15981 Obj

			if True == ifres15982 {
				ifres15981 = True

			} else {
				ifres15981 = False

			}

			ifres15980 = ifres15981

		} else {
			ifres15980 = False

		}

		if True == ifres15980 {
			tmp15973 := PrimHead(V1908)

			tmp15974 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp15973)

			tmp15975 := PrimTail(V1908)

			tmp15976 := PrimHead(tmp15975)

			tmp15977 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp15976)

			tmp15978 := PrimCons(tmp15977, Nil)

			tmp15979 := PrimCons(sym_1_6, tmp15978)

			__e.Return(PrimCons(tmp15974, tmp15979))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4demod_1rule)
			return
		}

	}, 1)

	tmp15990 := Call(__e, PrimNS1Value(symns2_1set), symshen_4demod_1rule, tmp15971)

	_ = tmp15990

	tmp15991 := MakeNative(func(__e *ControlFlow) {
		V1914 := __e.Get(1)
		_ = V1914
		tmp16041 := PrimIsPair(V1914)

		var ifres16002 Obj

		if True == tmp16041 {
			tmp16039 := PrimHead(V1914)

			tmp16040 := PrimEqual(symdefun, tmp16039)

			var ifres16004 Obj

			if True == tmp16040 {
				tmp16037 := PrimTail(V1914)

				tmp16038 := PrimIsPair(tmp16037)

				var ifres16006 Obj

				if True == tmp16038 {
					tmp16034 := PrimTail(V1914)

					tmp16035 := PrimTail(tmp16034)

					tmp16036 := PrimIsPair(tmp16035)

					var ifres16008 Obj

					if True == tmp16036 {
						tmp16030 := PrimTail(V1914)

						tmp16031 := PrimTail(tmp16030)

						tmp16032 := PrimHead(tmp16031)

						tmp16033 := PrimIsPair(tmp16032)

						var ifres16010 Obj

						if True == tmp16033 {
							tmp16025 := PrimTail(V1914)

							tmp16026 := PrimTail(tmp16025)

							tmp16027 := PrimHead(tmp16026)

							tmp16028 := PrimTail(tmp16027)

							tmp16029 := PrimEqual(Nil, tmp16028)

							var ifres16012 Obj

							if True == tmp16029 {
								tmp16021 := PrimTail(V1914)

								tmp16022 := PrimTail(tmp16021)

								tmp16023 := PrimTail(tmp16022)

								tmp16024 := PrimIsPair(tmp16023)

								var ifres16014 Obj

								if True == tmp16024 {
									tmp16016 := PrimTail(V1914)

									tmp16017 := PrimTail(tmp16016)

									tmp16018 := PrimTail(tmp16017)

									tmp16019 := PrimTail(tmp16018)

									tmp16020 := PrimEqual(Nil, tmp16019)

									var ifres16015 Obj

									if True == tmp16020 {
										ifres16015 = True

									} else {
										ifres16015 = False

									}

									ifres16014 = ifres16015

								} else {
									ifres16014 = False

								}

								var ifres16013 Obj

								if True == ifres16014 {
									ifres16013 = True

								} else {
									ifres16013 = False

								}

								ifres16012 = ifres16013

							} else {
								ifres16012 = False

							}

							var ifres16011 Obj

							if True == ifres16012 {
								ifres16011 = True

							} else {
								ifres16011 = False

							}

							ifres16010 = ifres16011

						} else {
							ifres16010 = False

						}

						var ifres16009 Obj

						if True == ifres16010 {
							ifres16009 = True

						} else {
							ifres16009 = False

						}

						ifres16008 = ifres16009

					} else {
						ifres16008 = False

					}

					var ifres16007 Obj

					if True == ifres16008 {
						ifres16007 = True

					} else {
						ifres16007 = False

					}

					ifres16006 = ifres16007

				} else {
					ifres16006 = False

				}

				var ifres16005 Obj

				if True == ifres16006 {
					ifres16005 = True

				} else {
					ifres16005 = False

				}

				ifres16004 = ifres16005

			} else {
				ifres16004 = False

			}

			var ifres16003 Obj

			if True == ifres16004 {
				ifres16003 = True

			} else {
				ifres16003 = False

			}

			ifres16002 = ifres16003

		} else {
			ifres16002 = False

		}

		if True == ifres16002 {
			tmp15993 := PrimTail(V1914)

			tmp15994 := PrimTail(tmp15993)

			tmp15995 := PrimHead(tmp15994)

			tmp15996 := PrimHead(tmp15995)

			tmp15997 := PrimTail(V1914)

			tmp15998 := PrimTail(tmp15997)

			tmp15999 := PrimTail(tmp15998)

			tmp16000 := PrimCons(tmp15996, tmp15999)

			tmp16001 := PrimCons(sym_c_4, tmp16000)

			__e.TailApply(PrimNS2Value(symeval), tmp16001)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4lambda_1of_1defun)
			return
		}

	}, 1)

	tmp16042 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1of_1defun, tmp15991)

	_ = tmp16042

	tmp16043 := MakeNative(func(__e *ControlFlow) {
		V1917 := __e.Get(1)
		_ = V1917
		V1918 := __e.Get(2)
		_ = V1918
		tmp16044 := Call(__e, PrimNS2Value(symtc), sym_1)

		_ = tmp16044

		tmp16045 := Call(__e, PrimNS2Value(symshen_4default_1rule))

		tmp16046 := Call(__e, PrimNS2Value(symappend), V1918, tmp16045)

		tmp16047 := PrimCons(symshen_4demod, tmp16046)

		tmp16048 := PrimCons(symdefine, tmp16047)

		tmp16049 := Call(__e, PrimNS2Value(symshen_4elim_1def), tmp16048)

		tmp16050 := Call(__e, PrimNS2Value(symshen_4lambda_1of_1defun), tmp16049)

		tmp16051 := PrimNS3Set(symshen_4_ddemodulation_1function_d, tmp16050)

		_ = tmp16051

		var ifres16052 Obj

		if True == V1917 {
			tmp16053 := Call(__e, PrimNS2Value(symtc), sym_7)

			ifres16052 = tmp16053

		} else {
			ifres16052 = symshen_4skip

		}

		_ = ifres16052

		__e.Return(symsynonyms)
		return

	}, 2)

	tmp16054 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update_1demodulation_1function, tmp16043)

	_ = tmp16054

	tmp16055 := MakeNative(func(__e *ControlFlow) {
		tmp16056 := PrimCons(symX, Nil)

		tmp16057 := PrimCons(sym_1_6, tmp16056)

		__e.Return(PrimCons(symX, tmp16057))
		return

	}, 0)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4default_1rule, tmp16055)
	return

}, 0)
