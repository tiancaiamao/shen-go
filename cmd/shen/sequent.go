package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen14957 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1722 := __args[0]
			_ = V1722
			gen14955 := Call(__e, ShenFunc(symcons_2), V1722)

			var gen14956 Obj
			if True == gen14955 {
				gen14952 := Call(__e, ShenFunc(symtl), V1722)

				gen14953 := Call(__e, ShenFunc(symcons_2), gen14952)

				var gen14954 Obj
				if True == gen14953 {
					gen14949 := Call(__e, ShenFunc(symtl), V1722)

					gen14950 := Call(__e, ShenFunc(symtl), gen14949)

					gen14951 := Call(__e, ShenFunc(sym_a), Nil, gen14950)

					if True == gen14951 {
						gen14954 = True
					} else {
						gen14954 = False
					}

				} else {
					gen14954 = False
				}
				if True == gen14954 {
					gen14956 = True
				} else {
					gen14956 = False
				}

			} else {
				gen14956 = False
			}
			if True == gen14956 {
				gen14945 := Call(__e, ShenFunc(symhd), V1722)

				gen14946 := Call(__e, ShenFunc(symshen_4next_150), MakeNumber(50), gen14945)

				gen14947 := Call(__e, ShenFunc(symshen_4app), gen14946, MakeString("\n"), MakeSymbol("shen.a"))

				gen14948 := Call(__e, ShenFunc(symcn), MakeString("datatype syntax error here:\n\n "), gen14947)

				__e.TailApply(ShenFunc(symsimple_1error), gen14948)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.datatype-error"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.datatype-error"), gen14957)

		gen14978 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1724 := __args[0]
			_ = V1724
			gen14958 := Call(__e, ShenFunc(symshen_4_5datatype_1rule_6), V1724)

			Parse__shen_4_5datatype_1rule_6 := gen14958
			gen14967 := Call(__e, ShenFunc(symfail))

			gen14968 := Call(__e, ShenFunc(sym_a), gen14967, Parse__shen_4_5datatype_1rule_6)

			gen14969 := Call(__e, ShenFunc(symnot), gen14968)

			var gen14970 Obj
			if True == gen14969 {
				gen14959 := Call(__e, ShenFunc(symshen_4_5datatype_1rules_6), Parse__shen_4_5datatype_1rule_6)

				Parse__shen_4_5datatype_1rules_6 := gen14959
				gen14964 := Call(__e, ShenFunc(symfail))

				gen14965 := Call(__e, ShenFunc(sym_a), gen14964, Parse__shen_4_5datatype_1rules_6)

				gen14966 := Call(__e, ShenFunc(symnot), gen14965)

				if True == gen14966 {
					gen14960 := Call(__e, ShenFunc(symhd), Parse__shen_4_5datatype_1rules_6)

					gen14961 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5datatype_1rule_6)

					gen14962 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5datatype_1rules_6)

					gen14963 := Call(__e, ShenFunc(symcons), gen14961, gen14962)

					gen14970 = Call(__e, ShenFunc(symshen_4pair), gen14960, gen14963)

				} else {
					gen14970 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen14970 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen14970
			gen14976 := Call(__e, ShenFunc(symfail))

			gen14977 := Call(__e, ShenFunc(sym_a), YaccParse, gen14976)

			if True == gen14977 {
				gen14971 := Call(__e, ShenFunc(sym_5e_6), V1724)

				Parse___5e_6 := gen14971
				gen14973 := Call(__e, ShenFunc(symfail))

				gen14974 := Call(__e, ShenFunc(sym_a), gen14973, Parse___5e_6)

				gen14975 := Call(__e, ShenFunc(symnot), gen14974)

				if True == gen14975 {
					gen14972 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen14972, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<datatype-rules>"), gen14978)

		gen15030 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1726 := __args[0]
			_ = V1726
			gen14979 := Call(__e, ShenFunc(symshen_4_5side_1conditions_6), V1726)

			Parse__shen_4_5side_1conditions_6 := gen14979
			gen15000 := Call(__e, ShenFunc(symfail))

			gen15001 := Call(__e, ShenFunc(sym_a), gen15000, Parse__shen_4_5side_1conditions_6)

			gen15002 := Call(__e, ShenFunc(symnot), gen15001)

			var gen15003 Obj
			if True == gen15002 {
				gen14980 := Call(__e, ShenFunc(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

				Parse__shen_4_5premises_6 := gen14980
				gen14997 := Call(__e, ShenFunc(symfail))

				gen14998 := Call(__e, ShenFunc(sym_a), gen14997, Parse__shen_4_5premises_6)

				gen14999 := Call(__e, ShenFunc(symnot), gen14998)

				if True == gen14999 {
					gen14981 := Call(__e, ShenFunc(symshen_4_5singleunderline_6), Parse__shen_4_5premises_6)

					Parse__shen_4_5singleunderline_6 := gen14981
					gen14994 := Call(__e, ShenFunc(symfail))

					gen14995 := Call(__e, ShenFunc(sym_a), gen14994, Parse__shen_4_5singleunderline_6)

					gen14996 := Call(__e, ShenFunc(symnot), gen14995)

					if True == gen14996 {
						gen14982 := Call(__e, ShenFunc(symshen_4_5conclusion_6), Parse__shen_4_5singleunderline_6)

						Parse__shen_4_5conclusion_6 := gen14982
						gen14991 := Call(__e, ShenFunc(symfail))

						gen14992 := Call(__e, ShenFunc(sym_a), gen14991, Parse__shen_4_5conclusion_6)

						gen14993 := Call(__e, ShenFunc(symnot), gen14992)

						if True == gen14993 {
							gen14983 := Call(__e, ShenFunc(symhd), Parse__shen_4_5conclusion_6)

							gen14984 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

							gen14985 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5premises_6)

							gen14986 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5conclusion_6)

							gen14987 := Call(__e, ShenFunc(symcons), gen14986, Nil)

							gen14988 := Call(__e, ShenFunc(symcons), gen14985, gen14987)

							gen14989 := Call(__e, ShenFunc(symcons), gen14984, gen14988)

							gen14990 := Call(__e, ShenFunc(symshen_4sequent), MakeSymbol("shen.single"), gen14989)

							gen15003 = Call(__e, ShenFunc(symshen_4pair), gen14983, gen14990)

						} else {
							gen15003 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen15003 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15003 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15003 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15003
			gen15028 := Call(__e, ShenFunc(symfail))

			gen15029 := Call(__e, ShenFunc(sym_a), YaccParse, gen15028)

			if True == gen15029 {
				gen15004 := Call(__e, ShenFunc(symshen_4_5side_1conditions_6), V1726)

				Parse__shen_4_5side_1conditions_6 := gen15004
				gen15025 := Call(__e, ShenFunc(symfail))

				gen15026 := Call(__e, ShenFunc(sym_a), gen15025, Parse__shen_4_5side_1conditions_6)

				gen15027 := Call(__e, ShenFunc(symnot), gen15026)

				if True == gen15027 {
					gen15005 := Call(__e, ShenFunc(symshen_4_5premises_6), Parse__shen_4_5side_1conditions_6)

					Parse__shen_4_5premises_6 := gen15005
					gen15022 := Call(__e, ShenFunc(symfail))

					gen15023 := Call(__e, ShenFunc(sym_a), gen15022, Parse__shen_4_5premises_6)

					gen15024 := Call(__e, ShenFunc(symnot), gen15023)

					if True == gen15024 {
						gen15006 := Call(__e, ShenFunc(symshen_4_5doubleunderline_6), Parse__shen_4_5premises_6)

						Parse__shen_4_5doubleunderline_6 := gen15006
						gen15019 := Call(__e, ShenFunc(symfail))

						gen15020 := Call(__e, ShenFunc(sym_a), gen15019, Parse__shen_4_5doubleunderline_6)

						gen15021 := Call(__e, ShenFunc(symnot), gen15020)

						if True == gen15021 {
							gen15007 := Call(__e, ShenFunc(symshen_4_5conclusion_6), Parse__shen_4_5doubleunderline_6)

							Parse__shen_4_5conclusion_6 := gen15007
							gen15016 := Call(__e, ShenFunc(symfail))

							gen15017 := Call(__e, ShenFunc(sym_a), gen15016, Parse__shen_4_5conclusion_6)

							gen15018 := Call(__e, ShenFunc(symnot), gen15017)

							if True == gen15018 {
								gen15008 := Call(__e, ShenFunc(symhd), Parse__shen_4_5conclusion_6)

								gen15009 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

								gen15010 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5premises_6)

								gen15011 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5conclusion_6)

								gen15012 := Call(__e, ShenFunc(symcons), gen15011, Nil)

								gen15013 := Call(__e, ShenFunc(symcons), gen15010, gen15012)

								gen15014 := Call(__e, ShenFunc(symcons), gen15009, gen15013)

								gen15015 := Call(__e, ShenFunc(symshen_4sequent), MakeSymbol("shen.double"), gen15014)

								__e.TailApply(ShenFunc(symshen_4pair), gen15008, gen15015)

								return

							} else {
								__e.TailApply(ShenFunc(symfail))

								return
							}

						} else {
							__e.TailApply(ShenFunc(symfail))

							return
						}

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<datatype-rule>"), gen15030)

		gen15051 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1728 := __args[0]
			_ = V1728
			gen15031 := Call(__e, ShenFunc(symshen_4_5side_1condition_6), V1728)

			Parse__shen_4_5side_1condition_6 := gen15031
			gen15040 := Call(__e, ShenFunc(symfail))

			gen15041 := Call(__e, ShenFunc(sym_a), gen15040, Parse__shen_4_5side_1condition_6)

			gen15042 := Call(__e, ShenFunc(symnot), gen15041)

			var gen15043 Obj
			if True == gen15042 {
				gen15032 := Call(__e, ShenFunc(symshen_4_5side_1conditions_6), Parse__shen_4_5side_1condition_6)

				Parse__shen_4_5side_1conditions_6 := gen15032
				gen15037 := Call(__e, ShenFunc(symfail))

				gen15038 := Call(__e, ShenFunc(sym_a), gen15037, Parse__shen_4_5side_1conditions_6)

				gen15039 := Call(__e, ShenFunc(symnot), gen15038)

				if True == gen15039 {
					gen15033 := Call(__e, ShenFunc(symhd), Parse__shen_4_5side_1conditions_6)

					gen15034 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5side_1condition_6)

					gen15035 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5side_1conditions_6)

					gen15036 := Call(__e, ShenFunc(symcons), gen15034, gen15035)

					gen15043 = Call(__e, ShenFunc(symshen_4pair), gen15033, gen15036)

				} else {
					gen15043 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15043 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15043
			gen15049 := Call(__e, ShenFunc(symfail))

			gen15050 := Call(__e, ShenFunc(sym_a), YaccParse, gen15049)

			if True == gen15050 {
				gen15044 := Call(__e, ShenFunc(sym_5e_6), V1728)

				Parse___5e_6 := gen15044
				gen15046 := Call(__e, ShenFunc(symfail))

				gen15047 := Call(__e, ShenFunc(sym_a), gen15046, Parse___5e_6)

				gen15048 := Call(__e, ShenFunc(symnot), gen15047)

				if True == gen15048 {
					gen15045 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen15045, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<side-conditions>"), gen15051)

		gen15093 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1732 := __args[0]
			_ = V1732
			gen15065 := Call(__e, ShenFunc(symhd), V1732)

			gen15066 := Call(__e, ShenFunc(symcons_2), gen15065)

			var gen15067 Obj
			if True == gen15066 {
				gen15063 := Call(__e, ShenFunc(symshen_4hdhd), V1732)

				gen15064 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen15063)

				if True == gen15064 {
					gen15067 = True
				} else {
					gen15067 = False
				}

			} else {
				gen15067 = False
			}
			var gen15068 Obj
			if True == gen15067 {
				gen15052 := Call(__e, ShenFunc(symshen_4tlhd), V1732)

				gen15053 := Call(__e, ShenFunc(symshen_4hdtl), V1732)

				gen15054 := Call(__e, ShenFunc(symshen_4pair), gen15052, gen15053)

				NewStream1729 := gen15054
				gen15055 := Call(__e, ShenFunc(symshen_4_5expr_6), NewStream1729)

				Parse__shen_4_5expr_6 := gen15055
				gen15060 := Call(__e, ShenFunc(symfail))

				gen15061 := Call(__e, ShenFunc(sym_a), gen15060, Parse__shen_4_5expr_6)

				gen15062 := Call(__e, ShenFunc(symnot), gen15061)

				if True == gen15062 {
					gen15056 := Call(__e, ShenFunc(symhd), Parse__shen_4_5expr_6)

					gen15057 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

					gen15058 := Call(__e, ShenFunc(symcons), gen15057, Nil)

					gen15059 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen15058)

					gen15068 = Call(__e, ShenFunc(symshen_4pair), gen15056, gen15059)

				} else {
					gen15068 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15068 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15068
			gen15091 := Call(__e, ShenFunc(symfail))

			gen15092 := Call(__e, ShenFunc(sym_a), YaccParse, gen15091)

			if True == gen15092 {
				gen15088 := Call(__e, ShenFunc(symhd), V1732)

				gen15089 := Call(__e, ShenFunc(symcons_2), gen15088)

				var gen15090 Obj
				if True == gen15089 {
					gen15086 := Call(__e, ShenFunc(symshen_4hdhd), V1732)

					gen15087 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen15086)

					if True == gen15087 {
						gen15090 = True
					} else {
						gen15090 = False
					}

				} else {
					gen15090 = False
				}
				if True == gen15090 {
					gen15069 := Call(__e, ShenFunc(symshen_4tlhd), V1732)

					gen15070 := Call(__e, ShenFunc(symshen_4hdtl), V1732)

					gen15071 := Call(__e, ShenFunc(symshen_4pair), gen15069, gen15070)

					NewStream1730 := gen15071
					gen15072 := Call(__e, ShenFunc(symshen_4_5variable_2_6), NewStream1730)

					Parse__shen_4_5variable_2_6 := gen15072
					gen15083 := Call(__e, ShenFunc(symfail))

					gen15084 := Call(__e, ShenFunc(sym_a), gen15083, Parse__shen_4_5variable_2_6)

					gen15085 := Call(__e, ShenFunc(symnot), gen15084)

					if True == gen15085 {
						gen15073 := Call(__e, ShenFunc(symshen_4_5expr_6), Parse__shen_4_5variable_2_6)

						Parse__shen_4_5expr_6 := gen15073
						gen15080 := Call(__e, ShenFunc(symfail))

						gen15081 := Call(__e, ShenFunc(sym_a), gen15080, Parse__shen_4_5expr_6)

						gen15082 := Call(__e, ShenFunc(symnot), gen15081)

						if True == gen15082 {
							gen15074 := Call(__e, ShenFunc(symhd), Parse__shen_4_5expr_6)

							gen15075 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5variable_2_6)

							gen15076 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

							gen15077 := Call(__e, ShenFunc(symcons), gen15076, Nil)

							gen15078 := Call(__e, ShenFunc(symcons), gen15075, gen15077)

							gen15079 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen15078)

							__e.TailApply(ShenFunc(symshen_4pair), gen15074, gen15079)

							return

						} else {
							__e.TailApply(ShenFunc(symfail))

							return
						}

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<side-condition>"), gen15093)

		gen15102 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1734 := __args[0]
			_ = V1734
			gen15100 := Call(__e, ShenFunc(symhd), V1734)

			gen15101 := Call(__e, ShenFunc(symcons_2), gen15100)

			if True == gen15101 {
				gen15094 := Call(__e, ShenFunc(symshen_4hdhd), V1734)

				Parse__X := gen15094
				gen15099 := Call(__e, ShenFunc(symvariable_2), Parse__X)

				if True == gen15099 {
					gen15095 := Call(__e, ShenFunc(symshen_4tlhd), V1734)

					gen15096 := Call(__e, ShenFunc(symshen_4hdtl), V1734)

					gen15097 := Call(__e, ShenFunc(symshen_4pair), gen15095, gen15096)

					gen15098 := Call(__e, ShenFunc(symhd), gen15097)

					__e.TailApply(ShenFunc(symshen_4pair), gen15098, Parse__X)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<variable?>"), gen15102)

		gen15119 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1736 := __args[0]
			_ = V1736
			gen15117 := Call(__e, ShenFunc(symhd), V1736)

			gen15118 := Call(__e, ShenFunc(symcons_2), gen15117)

			if True == gen15118 {
				gen15103 := Call(__e, ShenFunc(symshen_4hdhd), V1736)

				Parse__X := gen15103
				gen15112 := Call(__e, ShenFunc(symcons), MakeSymbol(";"), Nil)

				gen15113 := Call(__e, ShenFunc(symcons), MakeSymbol(">>"), gen15112)

				gen15114 := Call(__e, ShenFunc(symelement_2), Parse__X, gen15113)

				var gen15115 Obj
				if True == gen15114 {
					gen15115 = True
				} else {
					gen15110 := Call(__e, ShenFunc(symshen_4singleunderline_2), Parse__X)

					var gen15111 Obj
					if True == gen15110 {
						gen15111 = True
					} else {
						gen15109 := Call(__e, ShenFunc(symshen_4doubleunderline_2), Parse__X)

						if True == gen15109 {
							gen15111 = True
						} else {
							gen15111 = False
						}

					}
					if True == gen15111 {
						gen15115 = True
					} else {
						gen15115 = False
					}

				}
				gen15116 := Call(__e, ShenFunc(symnot), gen15115)

				if True == gen15116 {
					gen15104 := Call(__e, ShenFunc(symshen_4tlhd), V1736)

					gen15105 := Call(__e, ShenFunc(symshen_4hdtl), V1736)

					gen15106 := Call(__e, ShenFunc(symshen_4pair), gen15104, gen15105)

					gen15107 := Call(__e, ShenFunc(symhd), gen15106)

					gen15108 := Call(__e, ShenFunc(symshen_4remove_1bar), Parse__X)

					__e.TailApply(ShenFunc(symshen_4pair), gen15107, gen15108)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<expr>"), gen15119)

		gen15146 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1738 := __args[0]
			_ = V1738
			gen15144 := Call(__e, ShenFunc(symcons_2), V1738)

			var gen15145 Obj
			if True == gen15144 {
				gen15141 := Call(__e, ShenFunc(symtl), V1738)

				gen15142 := Call(__e, ShenFunc(symcons_2), gen15141)

				var gen15143 Obj
				if True == gen15142 {
					gen15137 := Call(__e, ShenFunc(symtl), V1738)

					gen15138 := Call(__e, ShenFunc(symtl), gen15137)

					gen15139 := Call(__e, ShenFunc(symcons_2), gen15138)

					var gen15140 Obj
					if True == gen15139 {
						gen15132 := Call(__e, ShenFunc(symtl), V1738)

						gen15133 := Call(__e, ShenFunc(symtl), gen15132)

						gen15134 := Call(__e, ShenFunc(symtl), gen15133)

						gen15135 := Call(__e, ShenFunc(sym_a), Nil, gen15134)

						var gen15136 Obj
						if True == gen15135 {
							gen15129 := Call(__e, ShenFunc(symtl), V1738)

							gen15130 := Call(__e, ShenFunc(symhd), gen15129)

							gen15131 := Call(__e, ShenFunc(sym_a), gen15130, MakeSymbol("bar!"))

							if True == gen15131 {
								gen15136 = True
							} else {
								gen15136 = False
							}

						} else {
							gen15136 = False
						}
						if True == gen15136 {
							gen15140 = True
						} else {
							gen15140 = False
						}

					} else {
						gen15140 = False
					}
					if True == gen15140 {
						gen15143 = True
					} else {
						gen15143 = False
					}

				} else {
					gen15143 = False
				}
				if True == gen15143 {
					gen15145 = True
				} else {
					gen15145 = False
				}

			} else {
				gen15145 = False
			}
			if True == gen15145 {
				gen15125 := Call(__e, ShenFunc(symhd), V1738)

				gen15126 := Call(__e, ShenFunc(symtl), V1738)

				gen15127 := Call(__e, ShenFunc(symtl), gen15126)

				gen15128 := Call(__e, ShenFunc(symhd), gen15127)

				__e.TailApply(ShenFunc(symcons), gen15125, gen15128)

				return

			} else {
				gen15124 := Call(__e, ShenFunc(symcons_2), V1738)

				if True == gen15124 {
					gen15120 := Call(__e, ShenFunc(symhd), V1738)

					gen15121 := Call(__e, ShenFunc(symshen_4remove_1bar), gen15120)

					gen15122 := Call(__e, ShenFunc(symtl), V1738)

					gen15123 := Call(__e, ShenFunc(symshen_4remove_1bar), gen15122)

					__e.TailApply(ShenFunc(symcons), gen15121, gen15123)

					return

				} else {
					__e.Return(V1738)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove-bar"), gen15146)

		gen15171 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1740 := __args[0]
			_ = V1740
			gen15147 := Call(__e, ShenFunc(symshen_4_5premise_6), V1740)

			Parse__shen_4_5premise_6 := gen15147
			gen15160 := Call(__e, ShenFunc(symfail))

			gen15161 := Call(__e, ShenFunc(sym_a), gen15160, Parse__shen_4_5premise_6)

			gen15162 := Call(__e, ShenFunc(symnot), gen15161)

			var gen15163 Obj
			if True == gen15162 {
				gen15148 := Call(__e, ShenFunc(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5premise_6)

				Parse__shen_4_5semicolon_1symbol_6 := gen15148
				gen15157 := Call(__e, ShenFunc(symfail))

				gen15158 := Call(__e, ShenFunc(sym_a), gen15157, Parse__shen_4_5semicolon_1symbol_6)

				gen15159 := Call(__e, ShenFunc(symnot), gen15158)

				if True == gen15159 {
					gen15149 := Call(__e, ShenFunc(symshen_4_5premises_6), Parse__shen_4_5semicolon_1symbol_6)

					Parse__shen_4_5premises_6 := gen15149
					gen15154 := Call(__e, ShenFunc(symfail))

					gen15155 := Call(__e, ShenFunc(sym_a), gen15154, Parse__shen_4_5premises_6)

					gen15156 := Call(__e, ShenFunc(symnot), gen15155)

					if True == gen15156 {
						gen15150 := Call(__e, ShenFunc(symhd), Parse__shen_4_5premises_6)

						gen15151 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5premise_6)

						gen15152 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5premises_6)

						gen15153 := Call(__e, ShenFunc(symcons), gen15151, gen15152)

						gen15163 = Call(__e, ShenFunc(symshen_4pair), gen15150, gen15153)

					} else {
						gen15163 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15163 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15163 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15163
			gen15169 := Call(__e, ShenFunc(symfail))

			gen15170 := Call(__e, ShenFunc(sym_a), YaccParse, gen15169)

			if True == gen15170 {
				gen15164 := Call(__e, ShenFunc(sym_5e_6), V1740)

				Parse___5e_6 := gen15164
				gen15166 := Call(__e, ShenFunc(symfail))

				gen15167 := Call(__e, ShenFunc(sym_a), gen15166, Parse___5e_6)

				gen15168 := Call(__e, ShenFunc(symnot), gen15167)

				if True == gen15168 {
					gen15165 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen15165, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<premises>"), gen15171)

		gen15180 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1742 := __args[0]
			_ = V1742
			gen15178 := Call(__e, ShenFunc(symhd), V1742)

			gen15179 := Call(__e, ShenFunc(symcons_2), gen15178)

			if True == gen15179 {
				gen15172 := Call(__e, ShenFunc(symshen_4hdhd), V1742)

				Parse__X := gen15172
				gen15177 := Call(__e, ShenFunc(sym_a), Parse__X, MakeSymbol(";"))

				if True == gen15177 {
					gen15173 := Call(__e, ShenFunc(symshen_4tlhd), V1742)

					gen15174 := Call(__e, ShenFunc(symshen_4hdtl), V1742)

					gen15175 := Call(__e, ShenFunc(symshen_4pair), gen15173, gen15174)

					gen15176 := Call(__e, ShenFunc(symhd), gen15175)

					__e.TailApply(ShenFunc(symshen_4pair), gen15176, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<semicolon-symbol>"), gen15180)

		gen15223 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1746 := __args[0]
			_ = V1746
			gen15187 := Call(__e, ShenFunc(symhd), V1746)

			gen15188 := Call(__e, ShenFunc(symcons_2), gen15187)

			var gen15189 Obj
			if True == gen15188 {
				gen15185 := Call(__e, ShenFunc(symshen_4hdhd), V1746)

				gen15186 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), gen15185)

				if True == gen15186 {
					gen15189 = True
				} else {
					gen15189 = False
				}

			} else {
				gen15189 = False
			}
			var gen15190 Obj
			if True == gen15189 {
				gen15181 := Call(__e, ShenFunc(symshen_4tlhd), V1746)

				gen15182 := Call(__e, ShenFunc(symshen_4hdtl), V1746)

				gen15183 := Call(__e, ShenFunc(symshen_4pair), gen15181, gen15182)

				NewStream1743 := gen15183
				gen15184 := Call(__e, ShenFunc(symhd), NewStream1743)

				gen15190 = Call(__e, ShenFunc(symshen_4pair), gen15184, MakeSymbol("!"))

			} else {
				gen15190 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15190
			gen15221 := Call(__e, ShenFunc(symfail))

			gen15222 := Call(__e, ShenFunc(sym_a), YaccParse, gen15221)

			if True == gen15222 {
				gen15191 := Call(__e, ShenFunc(symshen_4_5formulae_6), V1746)

				Parse__shen_4_5formulae_6 := gen15191
				gen15208 := Call(__e, ShenFunc(symfail))

				gen15209 := Call(__e, ShenFunc(sym_a), gen15208, Parse__shen_4_5formulae_6)

				gen15210 := Call(__e, ShenFunc(symnot), gen15209)

				var gen15211 Obj
				if True == gen15210 {
					gen15205 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formulae_6)

					gen15206 := Call(__e, ShenFunc(symcons_2), gen15205)

					var gen15207 Obj
					if True == gen15206 {
						gen15203 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5formulae_6)

						gen15204 := Call(__e, ShenFunc(sym_a), MakeSymbol(">>"), gen15203)

						if True == gen15204 {
							gen15207 = True
						} else {
							gen15207 = False
						}

					} else {
						gen15207 = False
					}
					if True == gen15207 {
						gen15192 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5formulae_6)

						gen15193 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formulae_6)

						gen15194 := Call(__e, ShenFunc(symshen_4pair), gen15192, gen15193)

						NewStream1744 := gen15194
						gen15195 := Call(__e, ShenFunc(symshen_4_5formula_6), NewStream1744)

						Parse__shen_4_5formula_6 := gen15195
						gen15200 := Call(__e, ShenFunc(symfail))

						gen15201 := Call(__e, ShenFunc(sym_a), gen15200, Parse__shen_4_5formula_6)

						gen15202 := Call(__e, ShenFunc(symnot), gen15201)

						if True == gen15202 {
							gen15196 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formula_6)

							gen15197 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formulae_6)

							gen15198 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

							gen15199 := Call(__e, ShenFunc(symshen_4sequent), gen15197, gen15198)

							gen15211 = Call(__e, ShenFunc(symshen_4pair), gen15196, gen15199)

						} else {
							gen15211 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen15211 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15211 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen15211
				gen15219 := Call(__e, ShenFunc(symfail))

				gen15220 := Call(__e, ShenFunc(sym_a), YaccParse, gen15219)

				if True == gen15220 {
					gen15212 := Call(__e, ShenFunc(symshen_4_5formula_6), V1746)

					Parse__shen_4_5formula_6 := gen15212
					gen15216 := Call(__e, ShenFunc(symfail))

					gen15217 := Call(__e, ShenFunc(sym_a), gen15216, Parse__shen_4_5formula_6)

					gen15218 := Call(__e, ShenFunc(symnot), gen15217)

					if True == gen15218 {
						gen15213 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formula_6)

						gen15214 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen15215 := Call(__e, ShenFunc(symshen_4sequent), Nil, gen15214)

						__e.TailApply(ShenFunc(symshen_4pair), gen15213, gen15215)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<premise>"), gen15223)

		gen15262 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1749 := __args[0]
			_ = V1749
			gen15224 := Call(__e, ShenFunc(symshen_4_5formulae_6), V1749)

			Parse__shen_4_5formulae_6 := gen15224
			gen15245 := Call(__e, ShenFunc(symfail))

			gen15246 := Call(__e, ShenFunc(sym_a), gen15245, Parse__shen_4_5formulae_6)

			gen15247 := Call(__e, ShenFunc(symnot), gen15246)

			var gen15248 Obj
			if True == gen15247 {
				gen15242 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formulae_6)

				gen15243 := Call(__e, ShenFunc(symcons_2), gen15242)

				var gen15244 Obj
				if True == gen15243 {
					gen15240 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5formulae_6)

					gen15241 := Call(__e, ShenFunc(sym_a), MakeSymbol(">>"), gen15240)

					if True == gen15241 {
						gen15244 = True
					} else {
						gen15244 = False
					}

				} else {
					gen15244 = False
				}
				if True == gen15244 {
					gen15225 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5formulae_6)

					gen15226 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formulae_6)

					gen15227 := Call(__e, ShenFunc(symshen_4pair), gen15225, gen15226)

					NewStream1747 := gen15227
					gen15228 := Call(__e, ShenFunc(symshen_4_5formula_6), NewStream1747)

					Parse__shen_4_5formula_6 := gen15228
					gen15237 := Call(__e, ShenFunc(symfail))

					gen15238 := Call(__e, ShenFunc(sym_a), gen15237, Parse__shen_4_5formula_6)

					gen15239 := Call(__e, ShenFunc(symnot), gen15238)

					if True == gen15239 {
						gen15229 := Call(__e, ShenFunc(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

						Parse__shen_4_5semicolon_1symbol_6 := gen15229
						gen15234 := Call(__e, ShenFunc(symfail))

						gen15235 := Call(__e, ShenFunc(sym_a), gen15234, Parse__shen_4_5semicolon_1symbol_6)

						gen15236 := Call(__e, ShenFunc(symnot), gen15235)

						if True == gen15236 {
							gen15230 := Call(__e, ShenFunc(symhd), Parse__shen_4_5semicolon_1symbol_6)

							gen15231 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formulae_6)

							gen15232 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

							gen15233 := Call(__e, ShenFunc(symshen_4sequent), gen15231, gen15232)

							gen15248 = Call(__e, ShenFunc(symshen_4pair), gen15230, gen15233)

						} else {
							gen15248 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen15248 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15248 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15248 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15248
			gen15260 := Call(__e, ShenFunc(symfail))

			gen15261 := Call(__e, ShenFunc(sym_a), YaccParse, gen15260)

			if True == gen15261 {
				gen15249 := Call(__e, ShenFunc(symshen_4_5formula_6), V1749)

				Parse__shen_4_5formula_6 := gen15249
				gen15257 := Call(__e, ShenFunc(symfail))

				gen15258 := Call(__e, ShenFunc(sym_a), gen15257, Parse__shen_4_5formula_6)

				gen15259 := Call(__e, ShenFunc(symnot), gen15258)

				if True == gen15259 {
					gen15250 := Call(__e, ShenFunc(symshen_4_5semicolon_1symbol_6), Parse__shen_4_5formula_6)

					Parse__shen_4_5semicolon_1symbol_6 := gen15250
					gen15254 := Call(__e, ShenFunc(symfail))

					gen15255 := Call(__e, ShenFunc(sym_a), gen15254, Parse__shen_4_5semicolon_1symbol_6)

					gen15256 := Call(__e, ShenFunc(symnot), gen15255)

					if True == gen15256 {
						gen15251 := Call(__e, ShenFunc(symhd), Parse__shen_4_5semicolon_1symbol_6)

						gen15252 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen15253 := Call(__e, ShenFunc(symshen_4sequent), Nil, gen15252)

						__e.TailApply(ShenFunc(symshen_4pair), gen15251, gen15253)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

						return
					}

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<conclusion>"), gen15262)

		gen15263 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1752 := __args[0]
			_ = V1752
			V1753 := __args[1]
			_ = V1753
			__e.TailApply(ShenFunc(sym_8p), V1752, V1753)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sequent"), gen15263)

		gen15298 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1755 := __args[0]
			_ = V1755
			gen15264 := Call(__e, ShenFunc(symshen_4_5formula_6), V1755)

			Parse__shen_4_5formula_6 := gen15264
			gen15277 := Call(__e, ShenFunc(symfail))

			gen15278 := Call(__e, ShenFunc(sym_a), gen15277, Parse__shen_4_5formula_6)

			gen15279 := Call(__e, ShenFunc(symnot), gen15278)

			var gen15280 Obj
			if True == gen15279 {
				gen15265 := Call(__e, ShenFunc(symshen_4_5comma_1symbol_6), Parse__shen_4_5formula_6)

				Parse__shen_4_5comma_1symbol_6 := gen15265
				gen15274 := Call(__e, ShenFunc(symfail))

				gen15275 := Call(__e, ShenFunc(sym_a), gen15274, Parse__shen_4_5comma_1symbol_6)

				gen15276 := Call(__e, ShenFunc(symnot), gen15275)

				if True == gen15276 {
					gen15266 := Call(__e, ShenFunc(symshen_4_5formulae_6), Parse__shen_4_5comma_1symbol_6)

					Parse__shen_4_5formulae_6 := gen15266
					gen15271 := Call(__e, ShenFunc(symfail))

					gen15272 := Call(__e, ShenFunc(sym_a), gen15271, Parse__shen_4_5formulae_6)

					gen15273 := Call(__e, ShenFunc(symnot), gen15272)

					if True == gen15273 {
						gen15267 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formulae_6)

						gen15268 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

						gen15269 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formulae_6)

						gen15270 := Call(__e, ShenFunc(symcons), gen15268, gen15269)

						gen15280 = Call(__e, ShenFunc(symshen_4pair), gen15267, gen15270)

					} else {
						gen15280 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15280 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15280 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15280
			gen15296 := Call(__e, ShenFunc(symfail))

			gen15297 := Call(__e, ShenFunc(sym_a), YaccParse, gen15296)

			if True == gen15297 {
				gen15281 := Call(__e, ShenFunc(symshen_4_5formula_6), V1755)

				Parse__shen_4_5formula_6 := gen15281
				gen15285 := Call(__e, ShenFunc(symfail))

				gen15286 := Call(__e, ShenFunc(sym_a), gen15285, Parse__shen_4_5formula_6)

				gen15287 := Call(__e, ShenFunc(symnot), gen15286)

				var gen15288 Obj
				if True == gen15287 {
					gen15282 := Call(__e, ShenFunc(symhd), Parse__shen_4_5formula_6)

					gen15283 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5formula_6)

					gen15284 := Call(__e, ShenFunc(symcons), gen15283, Nil)

					gen15288 = Call(__e, ShenFunc(symshen_4pair), gen15282, gen15284)

				} else {
					gen15288 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen15288
				gen15294 := Call(__e, ShenFunc(symfail))

				gen15295 := Call(__e, ShenFunc(sym_a), YaccParse, gen15294)

				if True == gen15295 {
					gen15289 := Call(__e, ShenFunc(sym_5e_6), V1755)

					Parse___5e_6 := gen15289
					gen15291 := Call(__e, ShenFunc(symfail))

					gen15292 := Call(__e, ShenFunc(sym_a), gen15291, Parse___5e_6)

					gen15293 := Call(__e, ShenFunc(symnot), gen15292)

					if True == gen15293 {
						gen15290 := Call(__e, ShenFunc(symhd), Parse___5e_6)

						__e.TailApply(ShenFunc(symshen_4pair), gen15290, Nil)

						return

					} else {
						__e.TailApply(ShenFunc(symfail))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<formulae>"), gen15298)

		gen15308 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1757 := __args[0]
			_ = V1757
			gen15306 := Call(__e, ShenFunc(symhd), V1757)

			gen15307 := Call(__e, ShenFunc(symcons_2), gen15306)

			if True == gen15307 {
				gen15299 := Call(__e, ShenFunc(symshen_4hdhd), V1757)

				Parse__X := gen15299
				gen15304 := Call(__e, ShenFunc(symintern), MakeString(","))

				gen15305 := Call(__e, ShenFunc(sym_a), Parse__X, gen15304)

				if True == gen15305 {
					gen15300 := Call(__e, ShenFunc(symshen_4tlhd), V1757)

					gen15301 := Call(__e, ShenFunc(symshen_4hdtl), V1757)

					gen15302 := Call(__e, ShenFunc(symshen_4pair), gen15300, gen15301)

					gen15303 := Call(__e, ShenFunc(symhd), gen15302)

					__e.TailApply(ShenFunc(symshen_4pair), gen15303, MakeSymbol("shen.skip"))

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<comma-symbol>"), gen15308)

		gen15342 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1760 := __args[0]
			_ = V1760
			gen15309 := Call(__e, ShenFunc(symshen_4_5expr_6), V1760)

			Parse__shen_4_5expr_6 := gen15309
			gen15330 := Call(__e, ShenFunc(symfail))

			gen15331 := Call(__e, ShenFunc(sym_a), gen15330, Parse__shen_4_5expr_6)

			gen15332 := Call(__e, ShenFunc(symnot), gen15331)

			var gen15333 Obj
			if True == gen15332 {
				gen15327 := Call(__e, ShenFunc(symhd), Parse__shen_4_5expr_6)

				gen15328 := Call(__e, ShenFunc(symcons_2), gen15327)

				var gen15329 Obj
				if True == gen15328 {
					gen15325 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5expr_6)

					gen15326 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen15325)

					if True == gen15326 {
						gen15329 = True
					} else {
						gen15329 = False
					}

				} else {
					gen15329 = False
				}
				if True == gen15329 {
					gen15310 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5expr_6)

					gen15311 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

					gen15312 := Call(__e, ShenFunc(symshen_4pair), gen15310, gen15311)

					NewStream1758 := gen15312
					gen15313 := Call(__e, ShenFunc(symshen_4_5type_6), NewStream1758)

					Parse__shen_4_5type_6 := gen15313
					gen15322 := Call(__e, ShenFunc(symfail))

					gen15323 := Call(__e, ShenFunc(sym_a), gen15322, Parse__shen_4_5type_6)

					gen15324 := Call(__e, ShenFunc(symnot), gen15323)

					if True == gen15324 {
						gen15314 := Call(__e, ShenFunc(symhd), Parse__shen_4_5type_6)

						gen15315 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

						gen15316 := Call(__e, ShenFunc(symshen_4curry), gen15315)

						gen15317 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5type_6)

						gen15318 := Call(__e, ShenFunc(symshen_4demodulate), gen15317)

						gen15319 := Call(__e, ShenFunc(symcons), gen15318, Nil)

						gen15320 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen15319)

						gen15321 := Call(__e, ShenFunc(symcons), gen15316, gen15320)

						gen15333 = Call(__e, ShenFunc(symshen_4pair), gen15314, gen15321)

					} else {
						gen15333 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen15333 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen15333 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen15333
			gen15340 := Call(__e, ShenFunc(symfail))

			gen15341 := Call(__e, ShenFunc(sym_a), YaccParse, gen15340)

			if True == gen15341 {
				gen15334 := Call(__e, ShenFunc(symshen_4_5expr_6), V1760)

				Parse__shen_4_5expr_6 := gen15334
				gen15337 := Call(__e, ShenFunc(symfail))

				gen15338 := Call(__e, ShenFunc(sym_a), gen15337, Parse__shen_4_5expr_6)

				gen15339 := Call(__e, ShenFunc(symnot), gen15338)

				if True == gen15339 {
					gen15335 := Call(__e, ShenFunc(symhd), Parse__shen_4_5expr_6)

					gen15336 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen15335, gen15336)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<formula>"), gen15342)

		gen15350 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1762 := __args[0]
			_ = V1762
			gen15343 := Call(__e, ShenFunc(symshen_4_5expr_6), V1762)

			Parse__shen_4_5expr_6 := gen15343
			gen15347 := Call(__e, ShenFunc(symfail))

			gen15348 := Call(__e, ShenFunc(sym_a), gen15347, Parse__shen_4_5expr_6)

			gen15349 := Call(__e, ShenFunc(symnot), gen15348)

			if True == gen15349 {
				gen15344 := Call(__e, ShenFunc(symhd), Parse__shen_4_5expr_6)

				gen15345 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5expr_6)

				gen15346 := Call(__e, ShenFunc(symshen_4curry_1type), gen15345)

				__e.TailApply(ShenFunc(symshen_4pair), gen15344, gen15346)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<type>"), gen15350)

		gen15359 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1764 := __args[0]
			_ = V1764
			gen15357 := Call(__e, ShenFunc(symhd), V1764)

			gen15358 := Call(__e, ShenFunc(symcons_2), gen15357)

			if True == gen15358 {
				gen15351 := Call(__e, ShenFunc(symshen_4hdhd), V1764)

				Parse__X := gen15351
				gen15356 := Call(__e, ShenFunc(symshen_4doubleunderline_2), Parse__X)

				if True == gen15356 {
					gen15352 := Call(__e, ShenFunc(symshen_4tlhd), V1764)

					gen15353 := Call(__e, ShenFunc(symshen_4hdtl), V1764)

					gen15354 := Call(__e, ShenFunc(symshen_4pair), gen15352, gen15353)

					gen15355 := Call(__e, ShenFunc(symhd), gen15354)

					__e.TailApply(ShenFunc(symshen_4pair), gen15355, Parse__X)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<doubleunderline>"), gen15359)

		gen15368 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1766 := __args[0]
			_ = V1766
			gen15366 := Call(__e, ShenFunc(symhd), V1766)

			gen15367 := Call(__e, ShenFunc(symcons_2), gen15366)

			if True == gen15367 {
				gen15360 := Call(__e, ShenFunc(symshen_4hdhd), V1766)

				Parse__X := gen15360
				gen15365 := Call(__e, ShenFunc(symshen_4singleunderline_2), Parse__X)

				if True == gen15365 {
					gen15361 := Call(__e, ShenFunc(symshen_4tlhd), V1766)

					gen15362 := Call(__e, ShenFunc(symshen_4hdtl), V1766)

					gen15363 := Call(__e, ShenFunc(symshen_4pair), gen15361, gen15362)

					gen15364 := Call(__e, ShenFunc(symhd), gen15363)

					__e.TailApply(ShenFunc(symshen_4pair), gen15364, Parse__X)

					return

				} else {
					__e.TailApply(ShenFunc(symfail))

					return
				}

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<singleunderline>"), gen15368)

		gen15372 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1768 := __args[0]
			_ = V1768
			gen15371 := Call(__e, ShenFunc(symsymbol_2), V1768)

			if True == gen15371 {
				gen15369 := Call(__e, ShenFunc(symstr), V1768)

				gen15370 := Call(__e, ShenFunc(symshen_4sh_2), gen15369)

				if True == gen15370 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.singleunderline?"), gen15372)

		gen15378 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1770 := __args[0]
			_ = V1770
			gen15377 := Call(__e, ShenFunc(sym_a), MakeString("_"), V1770)

			if True == gen15377 {
				__e.Return(True)
				return
			} else {
				gen15375 := Call(__e, ShenFunc(sympos), V1770, MakeNumber(0))

				gen15376 := Call(__e, ShenFunc(sym_a), gen15375, MakeString("_"))

				if True == gen15376 {
					gen15373 := Call(__e, ShenFunc(symtlstr), V1770)

					gen15374 := Call(__e, ShenFunc(symshen_4sh_2), gen15373)

					if True == gen15374 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sh?"), gen15378)

		gen15382 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1772 := __args[0]
			_ = V1772
			gen15381 := Call(__e, ShenFunc(symsymbol_2), V1772)

			if True == gen15381 {
				gen15379 := Call(__e, ShenFunc(symstr), V1772)

				gen15380 := Call(__e, ShenFunc(symshen_4dh_2), gen15379)

				if True == gen15380 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.doubleunderline?"), gen15382)

		gen15388 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1774 := __args[0]
			_ = V1774
			gen15387 := Call(__e, ShenFunc(sym_a), MakeString("="), V1774)

			if True == gen15387 {
				__e.Return(True)
				return
			} else {
				gen15385 := Call(__e, ShenFunc(sympos), V1774, MakeNumber(0))

				gen15386 := Call(__e, ShenFunc(sym_a), gen15385, MakeString("="))

				if True == gen15386 {
					gen15383 := Call(__e, ShenFunc(symtlstr), V1774)

					gen15384 := Call(__e, ShenFunc(symshen_4dh_2), gen15383)

					if True == gen15384 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.dh?"), gen15388)

		gen15391 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1777 := __args[0]
			_ = V1777
			V1778 := __args[1]
			_ = V1778
			gen15389 := Call(__e, ShenFunc(symshen_4rules_1_6horn_1clauses), V1777, V1778)

			gen15390 := Call(__e, ShenFunc(symshen_4s_1prolog), gen15389)

			__e.TailApply(ShenFunc(symshen_4remember_1datatype), gen15390)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.process-datatype"), gen15391)

		gen15399 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1784 := __args[0]
			_ = V1784
			gen15398 := Call(__e, ShenFunc(symcons_2), V1784)

			if True == gen15398 {
				gen15392 := Call(__e, ShenFunc(symhd), V1784)

				gen15393 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

				gen15394 := Call(__e, ShenFunc(symadjoin), gen15392, gen15393)

				Call(__e, ShenFunc(symset), MakeSymbol("shen.*datatypes*"), gen15394)

				gen15395 := Call(__e, ShenFunc(symhd), V1784)

				gen15396 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*alldatatypes*"))

				gen15397 := Call(__e, ShenFunc(symadjoin), gen15395, gen15396)

				Call(__e, ShenFunc(symset), MakeSymbol("shen.*alldatatypes*"), gen15397)

				__e.TailApply(ShenFunc(symhd), V1784)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.remember-datatype"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remember-datatype"), gen15399)

		gen15427 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1789 := __args[0]
			_ = V1789
			V1790 := __args[1]
			_ = V1790
			gen15426 := Call(__e, ShenFunc(sym_a), Nil, V1790)

			if True == gen15426 {
				__e.Return(Nil)
				return
			} else {
				gen15424 := Call(__e, ShenFunc(symcons_2), V1790)

				var gen15425 Obj
				if True == gen15424 {
					gen15421 := Call(__e, ShenFunc(symhd), V1790)

					gen15422 := Call(__e, ShenFunc(symtuple_2), gen15421)

					var gen15423 Obj
					if True == gen15422 {
						gen15418 := Call(__e, ShenFunc(symhd), V1790)

						gen15419 := Call(__e, ShenFunc(symfst), gen15418)

						gen15420 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.single"), gen15419)

						if True == gen15420 {
							gen15423 = True
						} else {
							gen15423 = False
						}

					} else {
						gen15423 = False
					}
					if True == gen15423 {
						gen15425 = True
					} else {
						gen15425 = False
					}

				} else {
					gen15425 = False
				}
				if True == gen15425 {
					gen15413 := Call(__e, ShenFunc(symhd), V1790)

					gen15414 := Call(__e, ShenFunc(symsnd), gen15413)

					gen15415 := Call(__e, ShenFunc(symshen_4rule_1_6horn_1clause), V1789, gen15414)

					gen15416 := Call(__e, ShenFunc(symtl), V1790)

					gen15417 := Call(__e, ShenFunc(symshen_4rules_1_6horn_1clauses), V1789, gen15416)

					__e.TailApply(ShenFunc(symcons), gen15415, gen15417)

					return

				} else {
					gen15411 := Call(__e, ShenFunc(symcons_2), V1790)

					var gen15412 Obj
					if True == gen15411 {
						gen15408 := Call(__e, ShenFunc(symhd), V1790)

						gen15409 := Call(__e, ShenFunc(symtuple_2), gen15408)

						var gen15410 Obj
						if True == gen15409 {
							gen15405 := Call(__e, ShenFunc(symhd), V1790)

							gen15406 := Call(__e, ShenFunc(symfst), gen15405)

							gen15407 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.double"), gen15406)

							if True == gen15407 {
								gen15410 = True
							} else {
								gen15410 = False
							}

						} else {
							gen15410 = False
						}
						if True == gen15410 {
							gen15412 = True
						} else {
							gen15412 = False
						}

					} else {
						gen15412 = False
					}
					if True == gen15412 {
						gen15400 := Call(__e, ShenFunc(symhd), V1790)

						gen15401 := Call(__e, ShenFunc(symsnd), gen15400)

						gen15402 := Call(__e, ShenFunc(symshen_4double_1_6singles), gen15401)

						gen15403 := Call(__e, ShenFunc(symtl), V1790)

						gen15404 := Call(__e, ShenFunc(symappend), gen15402, gen15403)

						__e.TailApply(ShenFunc(symshen_4rules_1_6horn_1clauses), V1789, gen15404)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.rules->horn-clauses"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rules->horn-clauses"), gen15427)

		gen15431 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1792 := __args[0]
			_ = V1792
			gen15428 := Call(__e, ShenFunc(symshen_4right_1rule), V1792)

			gen15429 := Call(__e, ShenFunc(symshen_4left_1rule), V1792)

			gen15430 := Call(__e, ShenFunc(symcons), gen15429, Nil)

			__e.TailApply(ShenFunc(symcons), gen15428, gen15430)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.double->singles"), gen15431)

		gen15432 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1794 := __args[0]
			_ = V1794
			__e.TailApply(ShenFunc(sym_8p), MakeSymbol("shen.single"), V1794)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.right-rule"), gen15432)

		gen15474 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1796 := __args[0]
			_ = V1796
			gen15472 := Call(__e, ShenFunc(symcons_2), V1796)

			var gen15473 Obj
			if True == gen15472 {
				gen15469 := Call(__e, ShenFunc(symtl), V1796)

				gen15470 := Call(__e, ShenFunc(symcons_2), gen15469)

				var gen15471 Obj
				if True == gen15470 {
					gen15465 := Call(__e, ShenFunc(symtl), V1796)

					gen15466 := Call(__e, ShenFunc(symtl), gen15465)

					gen15467 := Call(__e, ShenFunc(symcons_2), gen15466)

					var gen15468 Obj
					if True == gen15467 {
						gen15460 := Call(__e, ShenFunc(symtl), V1796)

						gen15461 := Call(__e, ShenFunc(symtl), gen15460)

						gen15462 := Call(__e, ShenFunc(symhd), gen15461)

						gen15463 := Call(__e, ShenFunc(symtuple_2), gen15462)

						var gen15464 Obj
						if True == gen15463 {
							gen15454 := Call(__e, ShenFunc(symtl), V1796)

							gen15455 := Call(__e, ShenFunc(symtl), gen15454)

							gen15456 := Call(__e, ShenFunc(symhd), gen15455)

							gen15457 := Call(__e, ShenFunc(symfst), gen15456)

							gen15458 := Call(__e, ShenFunc(sym_a), Nil, gen15457)

							var gen15459 Obj
							if True == gen15458 {
								gen15450 := Call(__e, ShenFunc(symtl), V1796)

								gen15451 := Call(__e, ShenFunc(symtl), gen15450)

								gen15452 := Call(__e, ShenFunc(symtl), gen15451)

								gen15453 := Call(__e, ShenFunc(sym_a), Nil, gen15452)

								if True == gen15453 {
									gen15459 = True
								} else {
									gen15459 = False
								}

							} else {
								gen15459 = False
							}
							if True == gen15459 {
								gen15464 = True
							} else {
								gen15464 = False
							}

						} else {
							gen15464 = False
						}
						if True == gen15464 {
							gen15468 = True
						} else {
							gen15468 = False
						}

					} else {
						gen15468 = False
					}
					if True == gen15468 {
						gen15471 = True
					} else {
						gen15471 = False
					}

				} else {
					gen15471 = False
				}
				if True == gen15471 {
					gen15473 = True
				} else {
					gen15473 = False
				}

			} else {
				gen15473 = False
			}
			if True == gen15473 {
				gen15433 := Call(__e, ShenFunc(symgensym), MakeSymbol("Qv"))

				Q := gen15433
				gen15434 := Call(__e, ShenFunc(symtl), V1796)

				gen15435 := Call(__e, ShenFunc(symtl), gen15434)

				gen15436 := Call(__e, ShenFunc(symhd), gen15435)

				gen15437 := Call(__e, ShenFunc(symsnd), gen15436)

				gen15438 := Call(__e, ShenFunc(symcons), gen15437, Nil)

				gen15439 := Call(__e, ShenFunc(sym_8p), gen15438, Q)

				NewConclusion := gen15439
				gen15440 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4right_1_6left), X)

					return
				}, 1)
				gen15441 := Call(__e, ShenFunc(symtl), V1796)

				gen15442 := Call(__e, ShenFunc(symhd), gen15441)

				gen15443 := Call(__e, ShenFunc(symmap), gen15440, gen15442)

				gen15444 := Call(__e, ShenFunc(sym_8p), gen15443, Q)

				gen15445 := Call(__e, ShenFunc(symcons), gen15444, Nil)

				NewPremises := gen15445
				gen15446 := Call(__e, ShenFunc(symhd), V1796)

				gen15447 := Call(__e, ShenFunc(symcons), NewConclusion, Nil)

				gen15448 := Call(__e, ShenFunc(symcons), NewPremises, gen15447)

				gen15449 := Call(__e, ShenFunc(symcons), gen15446, gen15448)

				__e.TailApply(ShenFunc(sym_8p), MakeSymbol("shen.single"), gen15449)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.left-rule"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.left-rule"), gen15474)

		gen15479 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1802 := __args[0]
			_ = V1802
			gen15477 := Call(__e, ShenFunc(symtuple_2), V1802)

			var gen15478 Obj
			if True == gen15477 {
				gen15475 := Call(__e, ShenFunc(symfst), V1802)

				gen15476 := Call(__e, ShenFunc(sym_a), Nil, gen15475)

				if True == gen15476 {
					gen15478 = True
				} else {
					gen15478 = False
				}

			} else {
				gen15478 = False
			}
			if True == gen15478 {
				__e.TailApply(ShenFunc(symsnd), V1802)

				return
			} else {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("syntax error with ==========\n"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.right->left"), gen15479)

		gen15513 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1805 := __args[0]
			_ = V1805
			V1806 := __args[1]
			_ = V1806
			gen15511 := Call(__e, ShenFunc(symcons_2), V1806)

			var gen15512 Obj
			if True == gen15511 {
				gen15508 := Call(__e, ShenFunc(symtl), V1806)

				gen15509 := Call(__e, ShenFunc(symcons_2), gen15508)

				var gen15510 Obj
				if True == gen15509 {
					gen15504 := Call(__e, ShenFunc(symtl), V1806)

					gen15505 := Call(__e, ShenFunc(symtl), gen15504)

					gen15506 := Call(__e, ShenFunc(symcons_2), gen15505)

					var gen15507 Obj
					if True == gen15506 {
						gen15499 := Call(__e, ShenFunc(symtl), V1806)

						gen15500 := Call(__e, ShenFunc(symtl), gen15499)

						gen15501 := Call(__e, ShenFunc(symhd), gen15500)

						gen15502 := Call(__e, ShenFunc(symtuple_2), gen15501)

						var gen15503 Obj
						if True == gen15502 {
							gen15495 := Call(__e, ShenFunc(symtl), V1806)

							gen15496 := Call(__e, ShenFunc(symtl), gen15495)

							gen15497 := Call(__e, ShenFunc(symtl), gen15496)

							gen15498 := Call(__e, ShenFunc(sym_a), Nil, gen15497)

							if True == gen15498 {
								gen15503 = True
							} else {
								gen15503 = False
							}

						} else {
							gen15503 = False
						}
						if True == gen15503 {
							gen15507 = True
						} else {
							gen15507 = False
						}

					} else {
						gen15507 = False
					}
					if True == gen15507 {
						gen15510 = True
					} else {
						gen15510 = False
					}

				} else {
					gen15510 = False
				}
				if True == gen15510 {
					gen15512 = True
				} else {
					gen15512 = False
				}

			} else {
				gen15512 = False
			}
			if True == gen15512 {
				gen15480 := Call(__e, ShenFunc(symtl), V1806)

				gen15481 := Call(__e, ShenFunc(symtl), gen15480)

				gen15482 := Call(__e, ShenFunc(symhd), gen15481)

				gen15483 := Call(__e, ShenFunc(symsnd), gen15482)

				gen15484 := Call(__e, ShenFunc(symshen_4rule_1_6horn_1clause_1head), V1805, gen15483)

				gen15485 := Call(__e, ShenFunc(symhd), V1806)

				gen15486 := Call(__e, ShenFunc(symtl), V1806)

				gen15487 := Call(__e, ShenFunc(symhd), gen15486)

				gen15488 := Call(__e, ShenFunc(symtl), V1806)

				gen15489 := Call(__e, ShenFunc(symtl), gen15488)

				gen15490 := Call(__e, ShenFunc(symhd), gen15489)

				gen15491 := Call(__e, ShenFunc(symfst), gen15490)

				gen15492 := Call(__e, ShenFunc(symshen_4rule_1_6horn_1clause_1body), gen15485, gen15487, gen15491)

				gen15493 := Call(__e, ShenFunc(symcons), gen15492, Nil)

				gen15494 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen15493)

				__e.TailApply(ShenFunc(symcons), gen15484, gen15494)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.rule->horn-clause"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rule->horn-clause"), gen15513)

		gen15517 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1809 := __args[0]
			_ = V1809
			V1810 := __args[1]
			_ = V1810
			gen15514 := Call(__e, ShenFunc(symshen_4mode_1ify), V1810)

			gen15515 := Call(__e, ShenFunc(symcons), MakeSymbol("Context_1957"), Nil)

			gen15516 := Call(__e, ShenFunc(symcons), gen15514, gen15515)

			__e.TailApply(ShenFunc(symcons), V1809, gen15516)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rule->horn-clause-head"), gen15517)

		gen15547 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1812 := __args[0]
			_ = V1812
			gen15545 := Call(__e, ShenFunc(symcons_2), V1812)

			var gen15546 Obj
			if True == gen15545 {
				gen15542 := Call(__e, ShenFunc(symtl), V1812)

				gen15543 := Call(__e, ShenFunc(symcons_2), gen15542)

				var gen15544 Obj
				if True == gen15543 {
					gen15538 := Call(__e, ShenFunc(symtl), V1812)

					gen15539 := Call(__e, ShenFunc(symhd), gen15538)

					gen15540 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen15539)

					var gen15541 Obj
					if True == gen15540 {
						gen15534 := Call(__e, ShenFunc(symtl), V1812)

						gen15535 := Call(__e, ShenFunc(symtl), gen15534)

						gen15536 := Call(__e, ShenFunc(symcons_2), gen15535)

						var gen15537 Obj
						if True == gen15536 {
							gen15530 := Call(__e, ShenFunc(symtl), V1812)

							gen15531 := Call(__e, ShenFunc(symtl), gen15530)

							gen15532 := Call(__e, ShenFunc(symtl), gen15531)

							gen15533 := Call(__e, ShenFunc(sym_a), Nil, gen15532)

							if True == gen15533 {
								gen15537 = True
							} else {
								gen15537 = False
							}

						} else {
							gen15537 = False
						}
						if True == gen15537 {
							gen15541 = True
						} else {
							gen15541 = False
						}

					} else {
						gen15541 = False
					}
					if True == gen15541 {
						gen15544 = True
					} else {
						gen15544 = False
					}

				} else {
					gen15544 = False
				}
				if True == gen15544 {
					gen15546 = True
				} else {
					gen15546 = False
				}

			} else {
				gen15546 = False
			}
			if True == gen15546 {
				gen15518 := Call(__e, ShenFunc(symhd), V1812)

				gen15519 := Call(__e, ShenFunc(symtl), V1812)

				gen15520 := Call(__e, ShenFunc(symtl), gen15519)

				gen15521 := Call(__e, ShenFunc(symhd), gen15520)

				gen15522 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

				gen15523 := Call(__e, ShenFunc(symcons), gen15521, gen15522)

				gen15524 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen15523)

				gen15525 := Call(__e, ShenFunc(symcons), gen15524, Nil)

				gen15526 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen15525)

				gen15527 := Call(__e, ShenFunc(symcons), gen15518, gen15526)

				gen15528 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), Nil)

				gen15529 := Call(__e, ShenFunc(symcons), gen15527, gen15528)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("mode"), gen15529)

				return

			} else {
				__e.Return(V1812)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mode-ify"), gen15547)

		gen15559 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1816 := __args[0]
			_ = V1816
			V1817 := __args[1]
			_ = V1817
			V1818 := __args[2]
			_ = V1818
			gen15548 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4extract__vars), X)

				return
			}, 1)
			gen15549 := Call(__e, ShenFunc(symmap), gen15548, V1818)

			Variables := gen15549
			gen15550 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symgensym), MakeSymbol("shen.cl"))

				return
			}, 1)
			gen15551 := Call(__e, ShenFunc(symmap), gen15550, V1818)

			Predicates := gen15551
			gen15552 := Call(__e, ShenFunc(symshen_4construct_1search_1literals), Predicates, Variables, MakeSymbol("Context_1957"), MakeSymbol("Context1_1957"))

			SearchLiterals := gen15552
			gen15553 := Call(__e, ShenFunc(symshen_4construct_1search_1clauses), Predicates, V1818, Variables)

			SearchClauses := gen15553
			_ = SearchClauses
			gen15554 := Call(__e, ShenFunc(symshen_4construct_1side_1literals), V1816)

			SideLiterals := gen15554
			gen15556 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				gen15555 := Call(__e, ShenFunc(symempty_2), V1818)

				__e.TailApply(ShenFunc(symshen_4construct_1premiss_1literal), X, gen15555)

				return

			}, 1)
			gen15557 := Call(__e, ShenFunc(symmap), gen15556, V1817)

			PremissLiterals := gen15557
			gen15558 := Call(__e, ShenFunc(symappend), SideLiterals, PremissLiterals)

			__e.TailApply(ShenFunc(symappend), SearchLiterals, gen15558)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rule->horn-clause-body"), gen15559)

		gen15563 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1827 := __args[0]
			_ = V1827
			V1828 := __args[1]
			_ = V1828
			V1829 := __args[2]
			_ = V1829
			V1830 := __args[3]
			_ = V1830
			gen15561 := Call(__e, ShenFunc(sym_a), Nil, V1827)

			var gen15562 Obj
			if True == gen15561 {
				gen15560 := Call(__e, ShenFunc(sym_a), Nil, V1828)

				if True == gen15560 {
					gen15562 = True
				} else {
					gen15562 = False
				}

			} else {
				gen15562 = False
			}
			if True == gen15562 {
				__e.Return(Nil)
				return
			} else {
				__e.TailApply(ShenFunc(symshen_4csl_1help), V1827, V1828, V1829, V1830)

				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-search-literals"), gen15563)

		gen15582 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1837 := __args[0]
			_ = V1837
			V1838 := __args[1]
			_ = V1838
			V1839 := __args[2]
			_ = V1839
			V1840 := __args[3]
			_ = V1840
			gen15580 := Call(__e, ShenFunc(sym_a), Nil, V1837)

			var gen15581 Obj
			if True == gen15580 {
				gen15579 := Call(__e, ShenFunc(sym_a), Nil, V1838)

				if True == gen15579 {
					gen15581 = True
				} else {
					gen15581 = False
				}

			} else {
				gen15581 = False
			}
			if True == gen15581 {
				gen15576 := Call(__e, ShenFunc(symcons), V1839, Nil)

				gen15577 := Call(__e, ShenFunc(symcons), MakeSymbol("ContextOut_1957"), gen15576)

				gen15578 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen15577)

				__e.TailApply(ShenFunc(symcons), gen15578, Nil)

				return

			} else {
				gen15574 := Call(__e, ShenFunc(symcons_2), V1837)

				var gen15575 Obj
				if True == gen15574 {
					gen15573 := Call(__e, ShenFunc(symcons_2), V1838)

					if True == gen15573 {
						gen15575 = True
					} else {
						gen15575 = False
					}

				} else {
					gen15575 = False
				}
				if True == gen15575 {
					gen15564 := Call(__e, ShenFunc(symhd), V1837)

					gen15565 := Call(__e, ShenFunc(symhd), V1838)

					gen15566 := Call(__e, ShenFunc(symcons), V1840, gen15565)

					gen15567 := Call(__e, ShenFunc(symcons), V1839, gen15566)

					gen15568 := Call(__e, ShenFunc(symcons), gen15564, gen15567)

					gen15569 := Call(__e, ShenFunc(symtl), V1837)

					gen15570 := Call(__e, ShenFunc(symtl), V1838)

					gen15571 := Call(__e, ShenFunc(symgensym), MakeSymbol("Context"))

					gen15572 := Call(__e, ShenFunc(symshen_4csl_1help), gen15569, gen15570, V1840, gen15571)

					__e.TailApply(ShenFunc(symcons), gen15568, gen15572)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.csl-help"))

					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.csl-help"), gen15582)

		gen15599 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1844 := __args[0]
			_ = V1844
			V1845 := __args[1]
			_ = V1845
			V1846 := __args[2]
			_ = V1846
			gen15597 := Call(__e, ShenFunc(sym_a), Nil, V1844)

			var gen15598 Obj
			if True == gen15597 {
				gen15595 := Call(__e, ShenFunc(sym_a), Nil, V1845)

				var gen15596 Obj
				if True == gen15595 {
					gen15594 := Call(__e, ShenFunc(sym_a), Nil, V1846)

					if True == gen15594 {
						gen15596 = True
					} else {
						gen15596 = False
					}

				} else {
					gen15596 = False
				}
				if True == gen15596 {
					gen15598 = True
				} else {
					gen15598 = False
				}

			} else {
				gen15598 = False
			}
			if True == gen15598 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen15592 := Call(__e, ShenFunc(symcons_2), V1844)

				var gen15593 Obj
				if True == gen15592 {
					gen15590 := Call(__e, ShenFunc(symcons_2), V1845)

					var gen15591 Obj
					if True == gen15590 {
						gen15589 := Call(__e, ShenFunc(symcons_2), V1846)

						if True == gen15589 {
							gen15591 = True
						} else {
							gen15591 = False
						}

					} else {
						gen15591 = False
					}
					if True == gen15591 {
						gen15593 = True
					} else {
						gen15593 = False
					}

				} else {
					gen15593 = False
				}
				if True == gen15593 {
					gen15583 := Call(__e, ShenFunc(symhd), V1844)

					gen15584 := Call(__e, ShenFunc(symhd), V1845)

					gen15585 := Call(__e, ShenFunc(symhd), V1846)

					Call(__e, ShenFunc(symshen_4construct_1search_1clause), gen15583, gen15584, gen15585)

					gen15586 := Call(__e, ShenFunc(symtl), V1844)

					gen15587 := Call(__e, ShenFunc(symtl), V1845)

					gen15588 := Call(__e, ShenFunc(symtl), V1846)

					__e.TailApply(ShenFunc(symshen_4construct_1search_1clauses), gen15586, gen15587, gen15588)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.construct-search-clauses"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-search-clauses"), gen15599)

		gen15604 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1850 := __args[0]
			_ = V1850
			V1851 := __args[1]
			_ = V1851
			V1852 := __args[2]
			_ = V1852
			gen15600 := Call(__e, ShenFunc(symshen_4construct_1base_1search_1clause), V1850, V1851, V1852)

			gen15601 := Call(__e, ShenFunc(symshen_4construct_1recursive_1search_1clause), V1850, V1851, V1852)

			gen15602 := Call(__e, ShenFunc(symcons), gen15601, Nil)

			gen15603 := Call(__e, ShenFunc(symcons), gen15600, gen15602)

			__e.TailApply(ShenFunc(symshen_4s_1prolog), gen15603)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-search-clause"), gen15604)

		gen15612 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1856 := __args[0]
			_ = V1856
			V1857 := __args[1]
			_ = V1857
			V1858 := __args[2]
			_ = V1858
			gen15605 := Call(__e, ShenFunc(symshen_4mode_1ify), V1857)

			gen15606 := Call(__e, ShenFunc(symcons), gen15605, MakeSymbol("In_1957"))

			gen15607 := Call(__e, ShenFunc(symcons), MakeSymbol("In_1957"), V1858)

			gen15608 := Call(__e, ShenFunc(symcons), gen15606, gen15607)

			gen15609 := Call(__e, ShenFunc(symcons), V1856, gen15608)

			gen15610 := Call(__e, ShenFunc(symcons), Nil, Nil)

			gen15611 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen15610)

			__e.TailApply(ShenFunc(symcons), gen15609, gen15611)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-base-search-clause"), gen15612)

		gen15624 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1862 := __args[0]
			_ = V1862
			V1863 := __args[1]
			_ = V1863
			V1864 := __args[2]
			_ = V1864
			gen15613 := Call(__e, ShenFunc(symcons), MakeSymbol("Assumption_1957"), MakeSymbol("Assumptions_1957"))

			gen15614 := Call(__e, ShenFunc(symcons), MakeSymbol("Assumption_1957"), MakeSymbol("Out_1957"))

			gen15615 := Call(__e, ShenFunc(symcons), gen15614, V1864)

			gen15616 := Call(__e, ShenFunc(symcons), gen15613, gen15615)

			gen15617 := Call(__e, ShenFunc(symcons), V1862, gen15616)

			gen15618 := Call(__e, ShenFunc(symcons), MakeSymbol("Out_1957"), V1864)

			gen15619 := Call(__e, ShenFunc(symcons), MakeSymbol("Assumptions_1957"), gen15618)

			gen15620 := Call(__e, ShenFunc(symcons), V1862, gen15619)

			gen15621 := Call(__e, ShenFunc(symcons), gen15620, Nil)

			gen15622 := Call(__e, ShenFunc(symcons), gen15621, Nil)

			gen15623 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen15622)

			__e.TailApply(ShenFunc(symcons), gen15617, gen15623)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-recursive-search-clause"), gen15624)

		gen15678 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1870 := __args[0]
			_ = V1870
			gen15677 := Call(__e, ShenFunc(sym_a), Nil, V1870)

			if True == gen15677 {
				__e.Return(Nil)
				return
			} else {
				gen15675 := Call(__e, ShenFunc(symcons_2), V1870)

				var gen15676 Obj
				if True == gen15675 {
					gen15672 := Call(__e, ShenFunc(symhd), V1870)

					gen15673 := Call(__e, ShenFunc(symcons_2), gen15672)

					var gen15674 Obj
					if True == gen15673 {
						gen15668 := Call(__e, ShenFunc(symhd), V1870)

						gen15669 := Call(__e, ShenFunc(symhd), gen15668)

						gen15670 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen15669)

						var gen15671 Obj
						if True == gen15670 {
							gen15664 := Call(__e, ShenFunc(symhd), V1870)

							gen15665 := Call(__e, ShenFunc(symtl), gen15664)

							gen15666 := Call(__e, ShenFunc(symcons_2), gen15665)

							var gen15667 Obj
							if True == gen15666 {
								gen15660 := Call(__e, ShenFunc(symhd), V1870)

								gen15661 := Call(__e, ShenFunc(symtl), gen15660)

								gen15662 := Call(__e, ShenFunc(symtl), gen15661)

								gen15663 := Call(__e, ShenFunc(sym_a), Nil, gen15662)

								if True == gen15663 {
									gen15667 = True
								} else {
									gen15667 = False
								}

							} else {
								gen15667 = False
							}
							if True == gen15667 {
								gen15671 = True
							} else {
								gen15671 = False
							}

						} else {
							gen15671 = False
						}
						if True == gen15671 {
							gen15674 = True
						} else {
							gen15674 = False
						}

					} else {
						gen15674 = False
					}
					if True == gen15674 {
						gen15676 = True
					} else {
						gen15676 = False
					}

				} else {
					gen15676 = False
				}
				if True == gen15676 {
					gen15655 := Call(__e, ShenFunc(symhd), V1870)

					gen15656 := Call(__e, ShenFunc(symtl), gen15655)

					gen15657 := Call(__e, ShenFunc(symcons), MakeSymbol("when"), gen15656)

					gen15658 := Call(__e, ShenFunc(symtl), V1870)

					gen15659 := Call(__e, ShenFunc(symshen_4construct_1side_1literals), gen15658)

					__e.TailApply(ShenFunc(symcons), gen15657, gen15659)

					return

				} else {
					gen15653 := Call(__e, ShenFunc(symcons_2), V1870)

					var gen15654 Obj
					if True == gen15653 {
						gen15650 := Call(__e, ShenFunc(symhd), V1870)

						gen15651 := Call(__e, ShenFunc(symcons_2), gen15650)

						var gen15652 Obj
						if True == gen15651 {
							gen15646 := Call(__e, ShenFunc(symhd), V1870)

							gen15647 := Call(__e, ShenFunc(symhd), gen15646)

							gen15648 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen15647)

							var gen15649 Obj
							if True == gen15648 {
								gen15642 := Call(__e, ShenFunc(symhd), V1870)

								gen15643 := Call(__e, ShenFunc(symtl), gen15642)

								gen15644 := Call(__e, ShenFunc(symcons_2), gen15643)

								var gen15645 Obj
								if True == gen15644 {
									gen15637 := Call(__e, ShenFunc(symhd), V1870)

									gen15638 := Call(__e, ShenFunc(symtl), gen15637)

									gen15639 := Call(__e, ShenFunc(symtl), gen15638)

									gen15640 := Call(__e, ShenFunc(symcons_2), gen15639)

									var gen15641 Obj
									if True == gen15640 {
										gen15632 := Call(__e, ShenFunc(symhd), V1870)

										gen15633 := Call(__e, ShenFunc(symtl), gen15632)

										gen15634 := Call(__e, ShenFunc(symtl), gen15633)

										gen15635 := Call(__e, ShenFunc(symtl), gen15634)

										gen15636 := Call(__e, ShenFunc(sym_a), Nil, gen15635)

										if True == gen15636 {
											gen15641 = True
										} else {
											gen15641 = False
										}

									} else {
										gen15641 = False
									}
									if True == gen15641 {
										gen15645 = True
									} else {
										gen15645 = False
									}

								} else {
									gen15645 = False
								}
								if True == gen15645 {
									gen15649 = True
								} else {
									gen15649 = False
								}

							} else {
								gen15649 = False
							}
							if True == gen15649 {
								gen15652 = True
							} else {
								gen15652 = False
							}

						} else {
							gen15652 = False
						}
						if True == gen15652 {
							gen15654 = True
						} else {
							gen15654 = False
						}

					} else {
						gen15654 = False
					}
					if True == gen15654 {
						gen15627 := Call(__e, ShenFunc(symhd), V1870)

						gen15628 := Call(__e, ShenFunc(symtl), gen15627)

						gen15629 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen15628)

						gen15630 := Call(__e, ShenFunc(symtl), V1870)

						gen15631 := Call(__e, ShenFunc(symshen_4construct_1side_1literals), gen15630)

						__e.TailApply(ShenFunc(symcons), gen15629, gen15631)

						return

					} else {
						gen15626 := Call(__e, ShenFunc(symcons_2), V1870)

						if True == gen15626 {
							gen15625 := Call(__e, ShenFunc(symtl), V1870)

							__e.TailApply(ShenFunc(symshen_4construct_1side_1literals), gen15625)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.construct-side-literals"))

							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-side-literals"), gen15678)

		gen15688 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1877 := __args[0]
			_ = V1877
			V1878 := __args[1]
			_ = V1878
			gen15687 := Call(__e, ShenFunc(symtuple_2), V1877)

			if True == gen15687 {
				gen15681 := Call(__e, ShenFunc(symsnd), V1877)

				gen15682 := Call(__e, ShenFunc(symshen_4recursive__cons__form), gen15681)

				gen15683 := Call(__e, ShenFunc(symfst), V1877)

				gen15684 := Call(__e, ShenFunc(symshen_4construct_1context), V1878, gen15683)

				gen15685 := Call(__e, ShenFunc(symcons), gen15684, Nil)

				gen15686 := Call(__e, ShenFunc(symcons), gen15682, gen15685)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.t*"), gen15686)

				return

			} else {
				gen15680 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), V1877)

				if True == gen15680 {
					gen15679 := Call(__e, ShenFunc(symcons), MakeSymbol("Throwcontrol"), Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("cut"), gen15679)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.construct-premiss-literal"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-premiss-literal"), gen15688)

		gen15702 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1881 := __args[0]
			_ = V1881
			V1882 := __args[1]
			_ = V1882
			gen15700 := Call(__e, ShenFunc(sym_a), True, V1881)

			var gen15701 Obj
			if True == gen15700 {
				gen15699 := Call(__e, ShenFunc(sym_a), Nil, V1882)

				if True == gen15699 {
					gen15701 = True
				} else {
					gen15701 = False
				}

			} else {
				gen15701 = False
			}
			if True == gen15701 {
				__e.Return(MakeSymbol("Context_1957"))
				return
			} else {
				gen15697 := Call(__e, ShenFunc(sym_a), False, V1881)

				var gen15698 Obj
				if True == gen15697 {
					gen15696 := Call(__e, ShenFunc(sym_a), Nil, V1882)

					if True == gen15696 {
						gen15698 = True
					} else {
						gen15698 = False
					}

				} else {
					gen15698 = False
				}
				if True == gen15698 {
					__e.Return(MakeSymbol("ContextOut_1957"))
					return
				} else {
					gen15695 := Call(__e, ShenFunc(symcons_2), V1882)

					if True == gen15695 {
						gen15689 := Call(__e, ShenFunc(symhd), V1882)

						gen15690 := Call(__e, ShenFunc(symshen_4recursive__cons__form), gen15689)

						gen15691 := Call(__e, ShenFunc(symtl), V1882)

						gen15692 := Call(__e, ShenFunc(symshen_4construct_1context), V1881, gen15691)

						gen15693 := Call(__e, ShenFunc(symcons), gen15692, Nil)

						gen15694 := Call(__e, ShenFunc(symcons), gen15690, gen15693)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen15694)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.construct-context"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.construct-context"), gen15702)

		gen15710 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1884 := __args[0]
			_ = V1884
			gen15709 := Call(__e, ShenFunc(symcons_2), V1884)

			if True == gen15709 {
				gen15703 := Call(__e, ShenFunc(symhd), V1884)

				gen15704 := Call(__e, ShenFunc(symshen_4recursive__cons__form), gen15703)

				gen15705 := Call(__e, ShenFunc(symtl), V1884)

				gen15706 := Call(__e, ShenFunc(symshen_4recursive__cons__form), gen15705)

				gen15707 := Call(__e, ShenFunc(symcons), gen15706, Nil)

				gen15708 := Call(__e, ShenFunc(symcons), gen15704, gen15707)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen15708)

				return

			} else {
				__e.Return(V1884)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.recursive_cons_form"), gen15710)

		gen15713 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1886 := __args[0]
			_ = V1886
			gen15711 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4intern_1type), X)

				return
			}, 1)
			gen15712 := Call(__e, ShenFunc(symmap), gen15711, V1886)

			__e.TailApply(ShenFunc(symshen_4preclude_1h), gen15712)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("preclude"), gen15713)

		gen15717 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1888 := __args[0]
			_ = V1888
			gen15714 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

			gen15715 := Call(__e, ShenFunc(symdifference), gen15714, V1888)

			gen15716 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*datatypes*"), gen15715)

			FilterDatatypes := gen15716
			_ = FilterDatatypes
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.preclude-h"), gen15717)

		gen15720 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1890 := __args[0]
			_ = V1890
			gen15718 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4intern_1type), X)

				return
			}, 1)
			gen15719 := Call(__e, ShenFunc(symmap), gen15718, V1890)

			__e.TailApply(ShenFunc(symshen_4include_1h), gen15719)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("include"), gen15720)

		gen15726 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1892 := __args[0]
			_ = V1892
			gen15721 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen15722 := Call(__e, ShenFunc(symintersection), V1892, gen15721)

			ValidTypes := gen15722
			gen15723 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

			gen15724 := Call(__e, ShenFunc(symunion), ValidTypes, gen15723)

			gen15725 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*datatypes*"), gen15724)

			NewDatatypes := gen15725
			_ = NewDatatypes
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.include-h"), gen15726)

		gen15731 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1894 := __args[0]
			_ = V1894
			gen15727 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen15728 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4intern_1type), X)

				return
			}, 1)
			gen15729 := Call(__e, ShenFunc(symmap), gen15728, V1894)

			gen15730 := Call(__e, ShenFunc(symdifference), gen15727, gen15729)

			__e.TailApply(ShenFunc(symshen_4preclude_1h), gen15730)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("preclude-all-but"), gen15731)

		gen15736 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1896 := __args[0]
			_ = V1896
			gen15732 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*alldatatypes*"))

			gen15733 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4intern_1type), X)

				return
			}, 1)
			gen15734 := Call(__e, ShenFunc(symmap), gen15733, V1896)

			gen15735 := Call(__e, ShenFunc(symdifference), gen15732, gen15734)

			__e.TailApply(ShenFunc(symshen_4include_1h), gen15735)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("include-all-but"), gen15736)

		gen15762 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1902 := __args[0]
			_ = V1902
			gen15761 := Call(__e, ShenFunc(sym_a), Nil, V1902)

			if True == gen15761 {
				gen15757 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*tc*"))

				gen15758 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4demod_1rule), X)

					return
				}, 1)
				gen15759 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*synonyms*"))

				gen15760 := Call(__e, ShenFunc(symmapcan), gen15758, gen15759)

				__e.TailApply(ShenFunc(symshen_4update_1demodulation_1function), gen15757, gen15760)

				return

			} else {
				gen15755 := Call(__e, ShenFunc(symcons_2), V1902)

				var gen15756 Obj
				if True == gen15755 {
					gen15753 := Call(__e, ShenFunc(symtl), V1902)

					gen15754 := Call(__e, ShenFunc(symcons_2), gen15753)

					if True == gen15754 {
						gen15756 = True
					} else {
						gen15756 = False
					}

				} else {
					gen15756 = False
				}
				if True == gen15756 {
					gen15737 := Call(__e, ShenFunc(symtl), V1902)

					gen15738 := Call(__e, ShenFunc(symhd), gen15737)

					gen15739 := Call(__e, ShenFunc(symshen_4extract__vars), gen15738)

					gen15740 := Call(__e, ShenFunc(symhd), V1902)

					gen15741 := Call(__e, ShenFunc(symshen_4extract__vars), gen15740)

					gen15742 := Call(__e, ShenFunc(symdifference), gen15739, gen15741)

					Vs := gen15742
					gen15752 := Call(__e, ShenFunc(symempty_2), Vs)

					if True == gen15752 {
						gen15745 := Call(__e, ShenFunc(symhd), V1902)

						gen15746 := Call(__e, ShenFunc(symtl), V1902)

						gen15747 := Call(__e, ShenFunc(symhd), gen15746)

						gen15748 := Call(__e, ShenFunc(symcons), gen15747, Nil)

						gen15749 := Call(__e, ShenFunc(symcons), gen15745, gen15748)

						Call(__e, ShenFunc(symshen_4pushnew), gen15749, MakeSymbol("shen.*synonyms*"))

						gen15750 := Call(__e, ShenFunc(symtl), V1902)

						gen15751 := Call(__e, ShenFunc(symtl), gen15750)

						__e.TailApply(ShenFunc(symshen_4synonyms_1help), gen15751)

						return

					} else {
						gen15743 := Call(__e, ShenFunc(symtl), V1902)

						gen15744 := Call(__e, ShenFunc(symhd), gen15743)

						__e.TailApply(ShenFunc(symshen_4free__variable__warnings), gen15744, Vs)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("odd number of synonyms\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.synonyms-help"), gen15762)

		gen15767 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1905 := __args[0]
			_ = V1905
			V1906 := __args[1]
			_ = V1906
			gen15765 := Call(__e, ShenFunc(symvalue), V1906)

			gen15766 := Call(__e, ShenFunc(symelement_2), V1905, gen15765)

			if True == gen15766 {
				__e.TailApply(ShenFunc(symvalue), V1906)

				return
			} else {
				gen15763 := Call(__e, ShenFunc(symvalue), V1906)

				gen15764 := Call(__e, ShenFunc(symcons), V1905, gen15763)

				__e.TailApply(ShenFunc(symset), V1906, gen15764)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pushnew"), gen15767)

		gen15783 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1908 := __args[0]
			_ = V1908
			gen15781 := Call(__e, ShenFunc(symcons_2), V1908)

			var gen15782 Obj
			if True == gen15781 {
				gen15778 := Call(__e, ShenFunc(symtl), V1908)

				gen15779 := Call(__e, ShenFunc(symcons_2), gen15778)

				var gen15780 Obj
				if True == gen15779 {
					gen15775 := Call(__e, ShenFunc(symtl), V1908)

					gen15776 := Call(__e, ShenFunc(symtl), gen15775)

					gen15777 := Call(__e, ShenFunc(sym_a), Nil, gen15776)

					if True == gen15777 {
						gen15780 = True
					} else {
						gen15780 = False
					}

				} else {
					gen15780 = False
				}
				if True == gen15780 {
					gen15782 = True
				} else {
					gen15782 = False
				}

			} else {
				gen15782 = False
			}
			if True == gen15782 {
				gen15768 := Call(__e, ShenFunc(symhd), V1908)

				gen15769 := Call(__e, ShenFunc(symshen_4rcons__form), gen15768)

				gen15770 := Call(__e, ShenFunc(symtl), V1908)

				gen15771 := Call(__e, ShenFunc(symhd), gen15770)

				gen15772 := Call(__e, ShenFunc(symshen_4rcons__form), gen15771)

				gen15773 := Call(__e, ShenFunc(symcons), gen15772, Nil)

				gen15774 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen15773)

				__e.TailApply(ShenFunc(symcons), gen15769, gen15774)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.demod-rule"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.demod-rule"), gen15783)

		gen15826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1914 := __args[0]
			_ = V1914
			gen15824 := Call(__e, ShenFunc(symcons_2), V1914)

			var gen15825 Obj
			if True == gen15824 {
				gen15821 := Call(__e, ShenFunc(symhd), V1914)

				gen15822 := Call(__e, ShenFunc(sym_a), MakeSymbol("defun"), gen15821)

				var gen15823 Obj
				if True == gen15822 {
					gen15818 := Call(__e, ShenFunc(symtl), V1914)

					gen15819 := Call(__e, ShenFunc(symcons_2), gen15818)

					var gen15820 Obj
					if True == gen15819 {
						gen15814 := Call(__e, ShenFunc(symtl), V1914)

						gen15815 := Call(__e, ShenFunc(symtl), gen15814)

						gen15816 := Call(__e, ShenFunc(symcons_2), gen15815)

						var gen15817 Obj
						if True == gen15816 {
							gen15809 := Call(__e, ShenFunc(symtl), V1914)

							gen15810 := Call(__e, ShenFunc(symtl), gen15809)

							gen15811 := Call(__e, ShenFunc(symhd), gen15810)

							gen15812 := Call(__e, ShenFunc(symcons_2), gen15811)

							var gen15813 Obj
							if True == gen15812 {
								gen15803 := Call(__e, ShenFunc(symtl), V1914)

								gen15804 := Call(__e, ShenFunc(symtl), gen15803)

								gen15805 := Call(__e, ShenFunc(symhd), gen15804)

								gen15806 := Call(__e, ShenFunc(symtl), gen15805)

								gen15807 := Call(__e, ShenFunc(sym_a), Nil, gen15806)

								var gen15808 Obj
								if True == gen15807 {
									gen15798 := Call(__e, ShenFunc(symtl), V1914)

									gen15799 := Call(__e, ShenFunc(symtl), gen15798)

									gen15800 := Call(__e, ShenFunc(symtl), gen15799)

									gen15801 := Call(__e, ShenFunc(symcons_2), gen15800)

									var gen15802 Obj
									if True == gen15801 {
										gen15793 := Call(__e, ShenFunc(symtl), V1914)

										gen15794 := Call(__e, ShenFunc(symtl), gen15793)

										gen15795 := Call(__e, ShenFunc(symtl), gen15794)

										gen15796 := Call(__e, ShenFunc(symtl), gen15795)

										gen15797 := Call(__e, ShenFunc(sym_a), Nil, gen15796)

										if True == gen15797 {
											gen15802 = True
										} else {
											gen15802 = False
										}

									} else {
										gen15802 = False
									}
									if True == gen15802 {
										gen15808 = True
									} else {
										gen15808 = False
									}

								} else {
									gen15808 = False
								}
								if True == gen15808 {
									gen15813 = True
								} else {
									gen15813 = False
								}

							} else {
								gen15813 = False
							}
							if True == gen15813 {
								gen15817 = True
							} else {
								gen15817 = False
							}

						} else {
							gen15817 = False
						}
						if True == gen15817 {
							gen15820 = True
						} else {
							gen15820 = False
						}

					} else {
						gen15820 = False
					}
					if True == gen15820 {
						gen15823 = True
					} else {
						gen15823 = False
					}

				} else {
					gen15823 = False
				}
				if True == gen15823 {
					gen15825 = True
				} else {
					gen15825 = False
				}

			} else {
				gen15825 = False
			}
			if True == gen15825 {
				gen15784 := Call(__e, ShenFunc(symtl), V1914)

				gen15785 := Call(__e, ShenFunc(symtl), gen15784)

				gen15786 := Call(__e, ShenFunc(symhd), gen15785)

				gen15787 := Call(__e, ShenFunc(symhd), gen15786)

				gen15788 := Call(__e, ShenFunc(symtl), V1914)

				gen15789 := Call(__e, ShenFunc(symtl), gen15788)

				gen15790 := Call(__e, ShenFunc(symtl), gen15789)

				gen15791 := Call(__e, ShenFunc(symcons), gen15787, gen15790)

				gen15792 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen15791)

				__e.TailApply(ShenFunc(symeval), gen15792)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.lambda-of-defun"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lambda-of-defun"), gen15826)

		gen15833 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1917 := __args[0]
			_ = V1917
			V1918 := __args[1]
			_ = V1918
			Call(__e, ShenFunc(symtc), MakeSymbol("-"))
			gen15827 := Call(__e, ShenFunc(symshen_4default_1rule))

			gen15828 := Call(__e, ShenFunc(symappend), V1918, gen15827)

			gen15829 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.demod"), gen15828)

			gen15830 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen15829)

			gen15831 := Call(__e, ShenFunc(symshen_4elim_1def), gen15830)

			gen15832 := Call(__e, ShenFunc(symshen_4lambda_1of_1defun), gen15831)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*demodulation-function*"), gen15832)

			if True == V1917 {
				Call(__e, ShenFunc(symtc), MakeSymbol("+"))
			} else {
				MakeSymbol("shen.skip")
			}
			__e.Return(MakeSymbol("synonyms"))
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.update-demodulation-function"), gen15833)

		gen15836 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen15834 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

			gen15835 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen15834)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("X"), gen15835)

			return

		}, 0)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.default-rule"), gen15836)

		return

	}, 0))
}
