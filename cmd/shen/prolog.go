package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen11942 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V867 := __args[0]
			_ = V867
			gen11927 := Call(__e, ShenFunc(symshen_4_5predicate_d_6), V867)

			Parse__shen_4_5predicate_d_6 := gen11927
			gen11939 := Call(__e, ShenFunc(symfail))

			gen11940 := Call(__e, ShenFunc(sym_a), gen11939, Parse__shen_4_5predicate_d_6)

			gen11941 := Call(__e, ShenFunc(symnot), gen11940)

			if True == gen11941 {
				gen11928 := Call(__e, ShenFunc(symshen_4_5clauses_d_6), Parse__shen_4_5predicate_d_6)

				Parse__shen_4_5clauses_d_6 := gen11928
				gen11936 := Call(__e, ShenFunc(symfail))

				gen11937 := Call(__e, ShenFunc(sym_a), gen11936, Parse__shen_4_5clauses_d_6)

				gen11938 := Call(__e, ShenFunc(symnot), gen11937)

				if True == gen11938 {
					gen11929 := Call(__e, ShenFunc(symhd), Parse__shen_4_5clauses_d_6)

					gen11931 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Parse__X := __args[0]
						_ = Parse__X
						gen11930 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5predicate_d_6)

						__e.TailApply(ShenFunc(symshen_4insert_1predicate), gen11930, Parse__X)

						return

					}, 1)
					gen11932 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

					gen11933 := Call(__e, ShenFunc(symmap), gen11931, gen11932)

					gen11934 := Call(__e, ShenFunc(symshen_4prolog_1_6shen), gen11933)

					gen11935 := Call(__e, ShenFunc(symhd), gen11934)

					__e.TailApply(ShenFunc(symshen_4pair), gen11929, gen11935)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<defprolog>"), gen11942)

		gen11959 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V876 := __args[0]
			_ = V876
			V877 := __args[1]
			_ = V877
			gen11957 := Call(__e, ShenFunc(symcons_2), V877)

			var gen11958 Obj
			if True == gen11957 {
				gen11954 := Call(__e, ShenFunc(symtl), V877)

				gen11955 := Call(__e, ShenFunc(symcons_2), gen11954)

				var gen11956 Obj
				if True == gen11955 {
					gen11951 := Call(__e, ShenFunc(symtl), V877)

					gen11952 := Call(__e, ShenFunc(symtl), gen11951)

					gen11953 := Call(__e, ShenFunc(sym_a), Nil, gen11952)

					if True == gen11953 {
						gen11956 = True
					} else {
						gen11956 = False
					}

				} else {
					gen11956 = False
				}
				if True == gen11956 {
					gen11958 = True
				} else {
					gen11958 = False
				}

			} else {
				gen11958 = False
			}
			if True == gen11958 {
				gen11945 := Call(__e, ShenFunc(symhd), V877)

				gen11946 := Call(__e, ShenFunc(symshen_4next_150), MakeNumber(50), gen11945)

				gen11947 := Call(__e, ShenFunc(symshen_4app), gen11946, MakeString("\n"), MakeSymbol("shen.a"))

				gen11948 := Call(__e, ShenFunc(symcn), MakeString(" here:\n\n "), gen11947)

				gen11949 := Call(__e, ShenFunc(symshen_4app), V876, gen11948, MakeSymbol("shen.a"))

				gen11950 := Call(__e, ShenFunc(symcn), MakeString("prolog syntax error in "), gen11949)

				__e.TailApply(ShenFunc(symsimple_1error), gen11950)

				return

			} else {
				gen11943 := Call(__e, ShenFunc(symshen_4app), V876, MakeString("\n"), MakeSymbol("shen.a"))

				gen11944 := Call(__e, ShenFunc(symcn), MakeString("prolog syntax error in "), gen11943)

				__e.TailApply(ShenFunc(symsimple_1error), gen11944)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-error"), gen11959)

		gen11968 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V884 := __args[0]
			_ = V884
			V885 := __args[1]
			_ = V885
			gen11967 := Call(__e, ShenFunc(sym_a), Nil, V885)

			if True == gen11967 {
				__e.Return(MakeString(""))
				return
			} else {
				gen11966 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V884)

				if True == gen11966 {
					__e.Return(MakeString(""))
					return
				} else {
					gen11965 := Call(__e, ShenFunc(symcons_2), V885)

					if True == gen11965 {
						gen11960 := Call(__e, ShenFunc(symhd), V885)

						gen11961 := Call(__e, ShenFunc(symshen_4decons_1string), gen11960)

						gen11962 := Call(__e, ShenFunc(sym_1), V884, MakeNumber(1))

						gen11963 := Call(__e, ShenFunc(symtl), V885)

						gen11964 := Call(__e, ShenFunc(symshen_4next_150), gen11962, gen11963)

						__e.TailApply(ShenFunc(symcn), gen11961, gen11964)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.next-50"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.next-50"), gen11968)

		gen11986 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V887 := __args[0]
			_ = V887
			gen11984 := Call(__e, ShenFunc(symcons_2), V887)

			var gen11985 Obj
			if True == gen11984 {
				gen11981 := Call(__e, ShenFunc(symhd), V887)

				gen11982 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen11981)

				var gen11983 Obj
				if True == gen11982 {
					gen11978 := Call(__e, ShenFunc(symtl), V887)

					gen11979 := Call(__e, ShenFunc(symcons_2), gen11978)

					var gen11980 Obj
					if True == gen11979 {
						gen11974 := Call(__e, ShenFunc(symtl), V887)

						gen11975 := Call(__e, ShenFunc(symtl), gen11974)

						gen11976 := Call(__e, ShenFunc(symcons_2), gen11975)

						var gen11977 Obj
						if True == gen11976 {
							gen11970 := Call(__e, ShenFunc(symtl), V887)

							gen11971 := Call(__e, ShenFunc(symtl), gen11970)

							gen11972 := Call(__e, ShenFunc(symtl), gen11971)

							gen11973 := Call(__e, ShenFunc(sym_a), Nil, gen11972)

							if True == gen11973 {
								gen11977 = True
							} else {
								gen11977 = False
							}

						} else {
							gen11977 = False
						}
						if True == gen11977 {
							gen11980 = True
						} else {
							gen11980 = False
						}

					} else {
						gen11980 = False
					}
					if True == gen11980 {
						gen11983 = True
					} else {
						gen11983 = False
					}

				} else {
					gen11983 = False
				}
				if True == gen11983 {
					gen11985 = True
				} else {
					gen11985 = False
				}

			} else {
				gen11985 = False
			}
			if True == gen11985 {
				gen11969 := Call(__e, ShenFunc(symshen_4eval_1cons), V887)

				__e.TailApply(ShenFunc(symshen_4app), gen11969, MakeString(" "), MakeSymbol("shen.s"))

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4app), V887, MakeString(" "), MakeSymbol("shen.r"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.decons-string"), gen11986)

		gen11999 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V890 := __args[0]
			_ = V890
			V891 := __args[1]
			_ = V891
			gen11997 := Call(__e, ShenFunc(symcons_2), V891)

			var gen11998 Obj
			if True == gen11997 {
				gen11994 := Call(__e, ShenFunc(symtl), V891)

				gen11995 := Call(__e, ShenFunc(symcons_2), gen11994)

				var gen11996 Obj
				if True == gen11995 {
					gen11991 := Call(__e, ShenFunc(symtl), V891)

					gen11992 := Call(__e, ShenFunc(symtl), gen11991)

					gen11993 := Call(__e, ShenFunc(sym_a), Nil, gen11992)

					if True == gen11993 {
						gen11996 = True
					} else {
						gen11996 = False
					}

				} else {
					gen11996 = False
				}
				if True == gen11996 {
					gen11998 = True
				} else {
					gen11998 = False
				}

			} else {
				gen11998 = False
			}
			if True == gen11998 {
				gen11987 := Call(__e, ShenFunc(symhd), V891)

				gen11988 := Call(__e, ShenFunc(symcons), V890, gen11987)

				gen11989 := Call(__e, ShenFunc(symtl), V891)

				gen11990 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen11989)

				__e.TailApply(ShenFunc(symcons), gen11988, gen11990)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-predicate"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-predicate"), gen11999)

		gen12007 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V893 := __args[0]
			_ = V893
			gen12005 := Call(__e, ShenFunc(symhd), V893)

			gen12006 := Call(__e, ShenFunc(symcons_2), gen12005)

			if True == gen12006 {
				gen12000 := Call(__e, ShenFunc(symshen_4hdhd), V893)

				Parse__X := gen12000
				gen12001 := Call(__e, ShenFunc(symshen_4tlhd), V893)

				gen12002 := Call(__e, ShenFunc(symshen_4hdtl), V893)

				gen12003 := Call(__e, ShenFunc(symshen_4pair), gen12001, gen12002)

				gen12004 := Call(__e, ShenFunc(symhd), gen12003)

				__e.TailApply(ShenFunc(symshen_4pair), gen12004, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<predicate*>"), gen12007)

		gen12028 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V895 := __args[0]
			_ = V895
			gen12008 := Call(__e, ShenFunc(symshen_4_5clause_d_6), V895)

			Parse__shen_4_5clause_d_6 := gen12008
			gen12017 := Call(__e, ShenFunc(symfail))

			gen12018 := Call(__e, ShenFunc(sym_a), gen12017, Parse__shen_4_5clause_d_6)

			gen12019 := Call(__e, ShenFunc(symnot), gen12018)

			var gen12020 Obj
			if True == gen12019 {
				gen12009 := Call(__e, ShenFunc(symshen_4_5clauses_d_6), Parse__shen_4_5clause_d_6)

				Parse__shen_4_5clauses_d_6 := gen12009
				gen12014 := Call(__e, ShenFunc(symfail))

				gen12015 := Call(__e, ShenFunc(sym_a), gen12014, Parse__shen_4_5clauses_d_6)

				gen12016 := Call(__e, ShenFunc(symnot), gen12015)

				if True == gen12016 {
					gen12010 := Call(__e, ShenFunc(symhd), Parse__shen_4_5clauses_d_6)

					gen12011 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clause_d_6)

					gen12012 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5clauses_d_6)

					gen12013 := Call(__e, ShenFunc(symcons), gen12011, gen12012)

					gen12020 = Call(__e, ShenFunc(symshen_4pair), gen12010, gen12013)

				} else {
					gen12020 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen12020 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen12020
			gen12026 := Call(__e, ShenFunc(symfail))

			gen12027 := Call(__e, ShenFunc(sym_a), YaccParse, gen12026)

			if True == gen12027 {
				gen12021 := Call(__e, ShenFunc(sym_5e_6), V895)

				Parse___5e_6 := gen12021
				gen12023 := Call(__e, ShenFunc(symfail))

				gen12024 := Call(__e, ShenFunc(sym_a), gen12023, Parse___5e_6)

				gen12025 := Call(__e, ShenFunc(symnot), gen12024)

				if True == gen12025 {
					gen12022 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen12022, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<clauses*>"), gen12028)

		gen12054 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V898 := __args[0]
			_ = V898
			gen12029 := Call(__e, ShenFunc(symshen_4_5head_d_6), V898)

			Parse__shen_4_5head_d_6 := gen12029
			gen12051 := Call(__e, ShenFunc(symfail))

			gen12052 := Call(__e, ShenFunc(sym_a), gen12051, Parse__shen_4_5head_d_6)

			gen12053 := Call(__e, ShenFunc(symnot), gen12052)

			if True == gen12053 {
				gen12048 := Call(__e, ShenFunc(symhd), Parse__shen_4_5head_d_6)

				gen12049 := Call(__e, ShenFunc(symcons_2), gen12048)

				var gen12050 Obj
				if True == gen12049 {
					gen12046 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5head_d_6)

					gen12047 := Call(__e, ShenFunc(sym_a), MakeSymbol("<--"), gen12046)

					if True == gen12047 {
						gen12050 = True
					} else {
						gen12050 = False
					}

				} else {
					gen12050 = False
				}
				if True == gen12050 {
					gen12030 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5head_d_6)

					gen12031 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

					gen12032 := Call(__e, ShenFunc(symshen_4pair), gen12030, gen12031)

					NewStream896 := gen12032
					gen12033 := Call(__e, ShenFunc(symshen_4_5body_d_6), NewStream896)

					Parse__shen_4_5body_d_6 := gen12033
					gen12043 := Call(__e, ShenFunc(symfail))

					gen12044 := Call(__e, ShenFunc(sym_a), gen12043, Parse__shen_4_5body_d_6)

					gen12045 := Call(__e, ShenFunc(symnot), gen12044)

					if True == gen12045 {
						gen12034 := Call(__e, ShenFunc(symshen_4_5end_d_6), Parse__shen_4_5body_d_6)

						Parse__shen_4_5end_d_6 := gen12034
						gen12040 := Call(__e, ShenFunc(symfail))

						gen12041 := Call(__e, ShenFunc(sym_a), gen12040, Parse__shen_4_5end_d_6)

						gen12042 := Call(__e, ShenFunc(symnot), gen12041)

						if True == gen12042 {
							gen12035 := Call(__e, ShenFunc(symhd), Parse__shen_4_5end_d_6)

							gen12036 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

							gen12037 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5body_d_6)

							gen12038 := Call(__e, ShenFunc(symcons), gen12037, Nil)

							gen12039 := Call(__e, ShenFunc(symcons), gen12036, gen12038)

							__e.TailApply(ShenFunc(symshen_4pair), gen12035, gen12039)

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

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<clause*>"), gen12054)

		gen12075 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V900 := __args[0]
			_ = V900
			gen12055 := Call(__e, ShenFunc(symshen_4_5term_d_6), V900)

			Parse__shen_4_5term_d_6 := gen12055
			gen12064 := Call(__e, ShenFunc(symfail))

			gen12065 := Call(__e, ShenFunc(sym_a), gen12064, Parse__shen_4_5term_d_6)

			gen12066 := Call(__e, ShenFunc(symnot), gen12065)

			var gen12067 Obj
			if True == gen12066 {
				gen12056 := Call(__e, ShenFunc(symshen_4_5head_d_6), Parse__shen_4_5term_d_6)

				Parse__shen_4_5head_d_6 := gen12056
				gen12061 := Call(__e, ShenFunc(symfail))

				gen12062 := Call(__e, ShenFunc(sym_a), gen12061, Parse__shen_4_5head_d_6)

				gen12063 := Call(__e, ShenFunc(symnot), gen12062)

				if True == gen12063 {
					gen12057 := Call(__e, ShenFunc(symhd), Parse__shen_4_5head_d_6)

					gen12058 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5term_d_6)

					gen12059 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5head_d_6)

					gen12060 := Call(__e, ShenFunc(symcons), gen12058, gen12059)

					gen12067 = Call(__e, ShenFunc(symshen_4pair), gen12057, gen12060)

				} else {
					gen12067 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen12067 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen12067
			gen12073 := Call(__e, ShenFunc(symfail))

			gen12074 := Call(__e, ShenFunc(sym_a), YaccParse, gen12073)

			if True == gen12074 {
				gen12068 := Call(__e, ShenFunc(sym_5e_6), V900)

				Parse___5e_6 := gen12068
				gen12070 := Call(__e, ShenFunc(symfail))

				gen12071 := Call(__e, ShenFunc(sym_a), gen12070, Parse___5e_6)

				gen12072 := Call(__e, ShenFunc(symnot), gen12071)

				if True == gen12072 {
					gen12069 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen12069, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<head*>"), gen12075)

		gen12088 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V902 := __args[0]
			_ = V902
			gen12086 := Call(__e, ShenFunc(symhd), V902)

			gen12087 := Call(__e, ShenFunc(symcons_2), gen12086)

			if True == gen12087 {
				gen12076 := Call(__e, ShenFunc(symshen_4hdhd), V902)

				Parse__X := gen12076
				gen12083 := Call(__e, ShenFunc(sym_a), MakeSymbol("<--"), Parse__X)

				gen12084 := Call(__e, ShenFunc(symnot), gen12083)

				var gen12085 Obj
				if True == gen12084 {
					gen12082 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), Parse__X)

					if True == gen12082 {
						gen12085 = True
					} else {
						gen12085 = False
					}

				} else {
					gen12085 = False
				}
				if True == gen12085 {
					gen12077 := Call(__e, ShenFunc(symshen_4tlhd), V902)

					gen12078 := Call(__e, ShenFunc(symshen_4hdtl), V902)

					gen12079 := Call(__e, ShenFunc(symshen_4pair), gen12077, gen12078)

					gen12080 := Call(__e, ShenFunc(symhd), gen12079)

					gen12081 := Call(__e, ShenFunc(symshen_4eval_1cons), Parse__X)

					__e.TailApply(ShenFunc(symshen_4pair), gen12080, gen12081)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<term*>"), gen12088)

		gen12159 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V908 := __args[0]
			_ = V908
			gen12157 := Call(__e, ShenFunc(symcons_2), V908)

			var gen12158 Obj
			if True == gen12157 {
				gen12154 := Call(__e, ShenFunc(symhd), V908)

				gen12155 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen12154)

				var gen12156 Obj
				if True == gen12155 {
					gen12151 := Call(__e, ShenFunc(symtl), V908)

					gen12152 := Call(__e, ShenFunc(symcons_2), gen12151)

					var gen12153 Obj
					if True == gen12152 {
						gen12147 := Call(__e, ShenFunc(symtl), V908)

						gen12148 := Call(__e, ShenFunc(symtl), gen12147)

						gen12149 := Call(__e, ShenFunc(symcons_2), gen12148)

						var gen12150 Obj
						if True == gen12149 {
							gen12143 := Call(__e, ShenFunc(symtl), V908)

							gen12144 := Call(__e, ShenFunc(symtl), gen12143)

							gen12145 := Call(__e, ShenFunc(symtl), gen12144)

							gen12146 := Call(__e, ShenFunc(sym_a), Nil, gen12145)

							if True == gen12146 {
								gen12150 = True
							} else {
								gen12150 = False
							}

						} else {
							gen12150 = False
						}
						if True == gen12150 {
							gen12153 = True
						} else {
							gen12153 = False
						}

					} else {
						gen12153 = False
					}
					if True == gen12153 {
						gen12156 = True
					} else {
						gen12156 = False
					}

				} else {
					gen12156 = False
				}
				if True == gen12156 {
					gen12158 = True
				} else {
					gen12158 = False
				}

			} else {
				gen12158 = False
			}
			if True == gen12158 {
				gen12140 := Call(__e, ShenFunc(symtl), V908)

				gen12141 := Call(__e, ShenFunc(symhd), gen12140)

				gen12142 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), gen12141)

				if True == gen12142 {
					gen12136 := Call(__e, ShenFunc(symtl), V908)

					gen12137 := Call(__e, ShenFunc(symtl), gen12136)

					gen12138 := Call(__e, ShenFunc(symhd), gen12137)

					gen12139 := Call(__e, ShenFunc(symshen_4legitimate_1term_2), gen12138)

					if True == gen12139 {
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
				gen12134 := Call(__e, ShenFunc(symcons_2), V908)

				var gen12135 Obj
				if True == gen12134 {
					gen12131 := Call(__e, ShenFunc(symhd), V908)

					gen12132 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12131)

					var gen12133 Obj
					if True == gen12132 {
						gen12128 := Call(__e, ShenFunc(symtl), V908)

						gen12129 := Call(__e, ShenFunc(symcons_2), gen12128)

						var gen12130 Obj
						if True == gen12129 {
							gen12124 := Call(__e, ShenFunc(symtl), V908)

							gen12125 := Call(__e, ShenFunc(symtl), gen12124)

							gen12126 := Call(__e, ShenFunc(symcons_2), gen12125)

							var gen12127 Obj
							if True == gen12126 {
								gen12119 := Call(__e, ShenFunc(symtl), V908)

								gen12120 := Call(__e, ShenFunc(symtl), gen12119)

								gen12121 := Call(__e, ShenFunc(symhd), gen12120)

								gen12122 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen12121)

								var gen12123 Obj
								if True == gen12122 {
									gen12115 := Call(__e, ShenFunc(symtl), V908)

									gen12116 := Call(__e, ShenFunc(symtl), gen12115)

									gen12117 := Call(__e, ShenFunc(symtl), gen12116)

									gen12118 := Call(__e, ShenFunc(sym_a), Nil, gen12117)

									if True == gen12118 {
										gen12123 = True
									} else {
										gen12123 = False
									}

								} else {
									gen12123 = False
								}
								if True == gen12123 {
									gen12127 = True
								} else {
									gen12127 = False
								}

							} else {
								gen12127 = False
							}
							if True == gen12127 {
								gen12130 = True
							} else {
								gen12130 = False
							}

						} else {
							gen12130 = False
						}
						if True == gen12130 {
							gen12133 = True
						} else {
							gen12133 = False
						}

					} else {
						gen12133 = False
					}
					if True == gen12133 {
						gen12135 = True
					} else {
						gen12135 = False
					}

				} else {
					gen12135 = False
				}
				if True == gen12135 {
					gen12113 := Call(__e, ShenFunc(symtl), V908)

					gen12114 := Call(__e, ShenFunc(symhd), gen12113)

					__e.TailApply(ShenFunc(symshen_4legitimate_1term_2), gen12114)

					return

				} else {
					gen12111 := Call(__e, ShenFunc(symcons_2), V908)

					var gen12112 Obj
					if True == gen12111 {
						gen12108 := Call(__e, ShenFunc(symhd), V908)

						gen12109 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12108)

						var gen12110 Obj
						if True == gen12109 {
							gen12105 := Call(__e, ShenFunc(symtl), V908)

							gen12106 := Call(__e, ShenFunc(symcons_2), gen12105)

							var gen12107 Obj
							if True == gen12106 {
								gen12101 := Call(__e, ShenFunc(symtl), V908)

								gen12102 := Call(__e, ShenFunc(symtl), gen12101)

								gen12103 := Call(__e, ShenFunc(symcons_2), gen12102)

								var gen12104 Obj
								if True == gen12103 {
									gen12096 := Call(__e, ShenFunc(symtl), V908)

									gen12097 := Call(__e, ShenFunc(symtl), gen12096)

									gen12098 := Call(__e, ShenFunc(symhd), gen12097)

									gen12099 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen12098)

									var gen12100 Obj
									if True == gen12099 {
										gen12092 := Call(__e, ShenFunc(symtl), V908)

										gen12093 := Call(__e, ShenFunc(symtl), gen12092)

										gen12094 := Call(__e, ShenFunc(symtl), gen12093)

										gen12095 := Call(__e, ShenFunc(sym_a), Nil, gen12094)

										if True == gen12095 {
											gen12100 = True
										} else {
											gen12100 = False
										}

									} else {
										gen12100 = False
									}
									if True == gen12100 {
										gen12104 = True
									} else {
										gen12104 = False
									}

								} else {
									gen12104 = False
								}
								if True == gen12104 {
									gen12107 = True
								} else {
									gen12107 = False
								}

							} else {
								gen12107 = False
							}
							if True == gen12107 {
								gen12110 = True
							} else {
								gen12110 = False
							}

						} else {
							gen12110 = False
						}
						if True == gen12110 {
							gen12112 = True
						} else {
							gen12112 = False
						}

					} else {
						gen12112 = False
					}
					if True == gen12112 {
						gen12090 := Call(__e, ShenFunc(symtl), V908)

						gen12091 := Call(__e, ShenFunc(symhd), gen12090)

						__e.TailApply(ShenFunc(symshen_4legitimate_1term_2), gen12091)

						return

					} else {
						gen12089 := Call(__e, ShenFunc(symcons_2), V908)

						if True == gen12089 {
							__e.Return(False)
							return
						} else {
							__e.Return(True)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.legitimate-term?"), gen12159)

		gen12205 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V910 := __args[0]
			_ = V910
			gen12203 := Call(__e, ShenFunc(symcons_2), V910)

			var gen12204 Obj
			if True == gen12203 {
				gen12200 := Call(__e, ShenFunc(symhd), V910)

				gen12201 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen12200)

				var gen12202 Obj
				if True == gen12201 {
					gen12197 := Call(__e, ShenFunc(symtl), V910)

					gen12198 := Call(__e, ShenFunc(symcons_2), gen12197)

					var gen12199 Obj
					if True == gen12198 {
						gen12193 := Call(__e, ShenFunc(symtl), V910)

						gen12194 := Call(__e, ShenFunc(symtl), gen12193)

						gen12195 := Call(__e, ShenFunc(symcons_2), gen12194)

						var gen12196 Obj
						if True == gen12195 {
							gen12189 := Call(__e, ShenFunc(symtl), V910)

							gen12190 := Call(__e, ShenFunc(symtl), gen12189)

							gen12191 := Call(__e, ShenFunc(symtl), gen12190)

							gen12192 := Call(__e, ShenFunc(sym_a), Nil, gen12191)

							if True == gen12192 {
								gen12196 = True
							} else {
								gen12196 = False
							}

						} else {
							gen12196 = False
						}
						if True == gen12196 {
							gen12199 = True
						} else {
							gen12199 = False
						}

					} else {
						gen12199 = False
					}
					if True == gen12199 {
						gen12202 = True
					} else {
						gen12202 = False
					}

				} else {
					gen12202 = False
				}
				if True == gen12202 {
					gen12204 = True
				} else {
					gen12204 = False
				}

			} else {
				gen12204 = False
			}
			if True == gen12204 {
				gen12182 := Call(__e, ShenFunc(symtl), V910)

				gen12183 := Call(__e, ShenFunc(symhd), gen12182)

				gen12184 := Call(__e, ShenFunc(symshen_4eval_1cons), gen12183)

				gen12185 := Call(__e, ShenFunc(symtl), V910)

				gen12186 := Call(__e, ShenFunc(symtl), gen12185)

				gen12187 := Call(__e, ShenFunc(symhd), gen12186)

				gen12188 := Call(__e, ShenFunc(symshen_4eval_1cons), gen12187)

				__e.TailApply(ShenFunc(symcons), gen12184, gen12188)

				return

			} else {
				gen12180 := Call(__e, ShenFunc(symcons_2), V910)

				var gen12181 Obj
				if True == gen12180 {
					gen12177 := Call(__e, ShenFunc(symhd), V910)

					gen12178 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12177)

					var gen12179 Obj
					if True == gen12178 {
						gen12174 := Call(__e, ShenFunc(symtl), V910)

						gen12175 := Call(__e, ShenFunc(symcons_2), gen12174)

						var gen12176 Obj
						if True == gen12175 {
							gen12170 := Call(__e, ShenFunc(symtl), V910)

							gen12171 := Call(__e, ShenFunc(symtl), gen12170)

							gen12172 := Call(__e, ShenFunc(symcons_2), gen12171)

							var gen12173 Obj
							if True == gen12172 {
								gen12166 := Call(__e, ShenFunc(symtl), V910)

								gen12167 := Call(__e, ShenFunc(symtl), gen12166)

								gen12168 := Call(__e, ShenFunc(symtl), gen12167)

								gen12169 := Call(__e, ShenFunc(sym_a), Nil, gen12168)

								if True == gen12169 {
									gen12173 = True
								} else {
									gen12173 = False
								}

							} else {
								gen12173 = False
							}
							if True == gen12173 {
								gen12176 = True
							} else {
								gen12176 = False
							}

						} else {
							gen12176 = False
						}
						if True == gen12176 {
							gen12179 = True
						} else {
							gen12179 = False
						}

					} else {
						gen12179 = False
					}
					if True == gen12179 {
						gen12181 = True
					} else {
						gen12181 = False
					}

				} else {
					gen12181 = False
				}
				if True == gen12181 {
					gen12160 := Call(__e, ShenFunc(symtl), V910)

					gen12161 := Call(__e, ShenFunc(symhd), gen12160)

					gen12162 := Call(__e, ShenFunc(symshen_4eval_1cons), gen12161)

					gen12163 := Call(__e, ShenFunc(symtl), V910)

					gen12164 := Call(__e, ShenFunc(symtl), gen12163)

					gen12165 := Call(__e, ShenFunc(symcons), gen12162, gen12164)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("mode"), gen12165)

					return

				} else {
					__e.Return(V910)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.eval-cons"), gen12205)

		gen12226 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V912 := __args[0]
			_ = V912
			gen12206 := Call(__e, ShenFunc(symshen_4_5literal_d_6), V912)

			Parse__shen_4_5literal_d_6 := gen12206
			gen12215 := Call(__e, ShenFunc(symfail))

			gen12216 := Call(__e, ShenFunc(sym_a), gen12215, Parse__shen_4_5literal_d_6)

			gen12217 := Call(__e, ShenFunc(symnot), gen12216)

			var gen12218 Obj
			if True == gen12217 {
				gen12207 := Call(__e, ShenFunc(symshen_4_5body_d_6), Parse__shen_4_5literal_d_6)

				Parse__shen_4_5body_d_6 := gen12207
				gen12212 := Call(__e, ShenFunc(symfail))

				gen12213 := Call(__e, ShenFunc(sym_a), gen12212, Parse__shen_4_5body_d_6)

				gen12214 := Call(__e, ShenFunc(symnot), gen12213)

				if True == gen12214 {
					gen12208 := Call(__e, ShenFunc(symhd), Parse__shen_4_5body_d_6)

					gen12209 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5literal_d_6)

					gen12210 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5body_d_6)

					gen12211 := Call(__e, ShenFunc(symcons), gen12209, gen12210)

					gen12218 = Call(__e, ShenFunc(symshen_4pair), gen12208, gen12211)

				} else {
					gen12218 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen12218 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen12218
			gen12224 := Call(__e, ShenFunc(symfail))

			gen12225 := Call(__e, ShenFunc(sym_a), YaccParse, gen12224)

			if True == gen12225 {
				gen12219 := Call(__e, ShenFunc(sym_5e_6), V912)

				Parse___5e_6 := gen12219
				gen12221 := Call(__e, ShenFunc(symfail))

				gen12222 := Call(__e, ShenFunc(sym_a), gen12221, Parse___5e_6)

				gen12223 := Call(__e, ShenFunc(symnot), gen12222)

				if True == gen12223 {
					gen12220 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen12220, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<body*>"), gen12226)

		gen12250 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V915 := __args[0]
			_ = V915
			gen12236 := Call(__e, ShenFunc(symhd), V915)

			gen12237 := Call(__e, ShenFunc(symcons_2), gen12236)

			var gen12238 Obj
			if True == gen12237 {
				gen12234 := Call(__e, ShenFunc(symshen_4hdhd), V915)

				gen12235 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), gen12234)

				if True == gen12235 {
					gen12238 = True
				} else {
					gen12238 = False
				}

			} else {
				gen12238 = False
			}
			var gen12239 Obj
			if True == gen12238 {
				gen12227 := Call(__e, ShenFunc(symshen_4tlhd), V915)

				gen12228 := Call(__e, ShenFunc(symshen_4hdtl), V915)

				gen12229 := Call(__e, ShenFunc(symshen_4pair), gen12227, gen12228)

				NewStream913 := gen12229
				gen12230 := Call(__e, ShenFunc(symhd), NewStream913)

				gen12231 := Call(__e, ShenFunc(symintern), MakeString("Throwcontrol"))

				gen12232 := Call(__e, ShenFunc(symcons), gen12231, Nil)

				gen12233 := Call(__e, ShenFunc(symcons), MakeSymbol("cut"), gen12232)

				gen12239 = Call(__e, ShenFunc(symshen_4pair), gen12230, gen12233)

			} else {
				gen12239 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen12239
			gen12248 := Call(__e, ShenFunc(symfail))

			gen12249 := Call(__e, ShenFunc(sym_a), YaccParse, gen12248)

			if True == gen12249 {
				gen12246 := Call(__e, ShenFunc(symhd), V915)

				gen12247 := Call(__e, ShenFunc(symcons_2), gen12246)

				if True == gen12247 {
					gen12240 := Call(__e, ShenFunc(symshen_4hdhd), V915)

					Parse__X := gen12240
					gen12245 := Call(__e, ShenFunc(symcons_2), Parse__X)

					if True == gen12245 {
						gen12241 := Call(__e, ShenFunc(symshen_4tlhd), V915)

						gen12242 := Call(__e, ShenFunc(symshen_4hdtl), V915)

						gen12243 := Call(__e, ShenFunc(symshen_4pair), gen12241, gen12242)

						gen12244 := Call(__e, ShenFunc(symhd), gen12243)

						__e.TailApply(ShenFunc(symshen_4pair), gen12244, Parse__X)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<literal*>"), gen12250)

		gen12259 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V917 := __args[0]
			_ = V917
			gen12257 := Call(__e, ShenFunc(symhd), V917)

			gen12258 := Call(__e, ShenFunc(symcons_2), gen12257)

			if True == gen12258 {
				gen12251 := Call(__e, ShenFunc(symshen_4hdhd), V917)

				Parse__X := gen12251
				gen12256 := Call(__e, ShenFunc(sym_a), Parse__X, MakeSymbol(";"))

				if True == gen12256 {
					gen12252 := Call(__e, ShenFunc(symshen_4tlhd), V917)

					gen12253 := Call(__e, ShenFunc(symshen_4hdtl), V917)

					gen12254 := Call(__e, ShenFunc(symshen_4pair), gen12252, gen12253)

					gen12255 := Call(__e, ShenFunc(symhd), gen12254)

					__e.TailApply(ShenFunc(symshen_4pair), gen12255, Parse__X)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<end*>"), gen12259)

		gen12262 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V921 := __args[0]
			_ = V921
			V922 := __args[1]
			_ = V922
			V923 := __args[2]
			_ = V923
			gen12260 := Call(__e, ShenFunc(symthaw), V923)

			Result := gen12260
			gen12261 := Call(__e, ShenFunc(sym_a), Result, False)

			if True == gen12261 {
				__e.Return(V921)
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("cut"), gen12262)

		gen12290 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V925 := __args[0]
			_ = V925
			gen12288 := Call(__e, ShenFunc(symcons_2), V925)

			var gen12289 Obj
			if True == gen12288 {
				gen12285 := Call(__e, ShenFunc(symhd), V925)

				gen12286 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12285)

				var gen12287 Obj
				if True == gen12286 {
					gen12282 := Call(__e, ShenFunc(symtl), V925)

					gen12283 := Call(__e, ShenFunc(symcons_2), gen12282)

					var gen12284 Obj
					if True == gen12283 {
						gen12278 := Call(__e, ShenFunc(symtl), V925)

						gen12279 := Call(__e, ShenFunc(symtl), gen12278)

						gen12280 := Call(__e, ShenFunc(symcons_2), gen12279)

						var gen12281 Obj
						if True == gen12280 {
							gen12274 := Call(__e, ShenFunc(symtl), V925)

							gen12275 := Call(__e, ShenFunc(symtl), gen12274)

							gen12276 := Call(__e, ShenFunc(symtl), gen12275)

							gen12277 := Call(__e, ShenFunc(sym_a), Nil, gen12276)

							if True == gen12277 {
								gen12281 = True
							} else {
								gen12281 = False
							}

						} else {
							gen12281 = False
						}
						if True == gen12281 {
							gen12284 = True
						} else {
							gen12284 = False
						}

					} else {
						gen12284 = False
					}
					if True == gen12284 {
						gen12287 = True
					} else {
						gen12287 = False
					}

				} else {
					gen12287 = False
				}
				if True == gen12287 {
					gen12289 = True
				} else {
					gen12289 = False
				}

			} else {
				gen12289 = False
			}
			if True == gen12289 {
				__e.Return(V925)
				return
			} else {
				gen12273 := Call(__e, ShenFunc(sym_a), Nil, V925)

				if True == gen12273 {
					__e.Return(Nil)
					return
				} else {
					gen12272 := Call(__e, ShenFunc(symcons_2), V925)

					if True == gen12272 {
						gen12263 := Call(__e, ShenFunc(symhd), V925)

						gen12264 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

						gen12265 := Call(__e, ShenFunc(symcons), gen12263, gen12264)

						gen12266 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12265)

						gen12267 := Call(__e, ShenFunc(symtl), V925)

						gen12268 := Call(__e, ShenFunc(symshen_4insert__modes), gen12267)

						gen12269 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), Nil)

						gen12270 := Call(__e, ShenFunc(symcons), gen12268, gen12269)

						gen12271 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12270)

						__e.TailApply(ShenFunc(symcons), gen12266, gen12271)

						return

					} else {
						__e.Return(V925)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert_modes"), gen12290)

		gen12293 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V927 := __args[0]
			_ = V927
			gen12291 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symeval), X)

				return
			}, 1)
			gen12292 := Call(__e, ShenFunc(symshen_4prolog_1_6shen), V927)

			__e.TailApply(ShenFunc(symmap), gen12291, gen12292)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog"), gen12293)

		gen12300 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V929 := __args[0]
			_ = V929
			gen12294 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4compile__prolog__procedure), X)

				return
			}, 1)
			gen12295 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4s_1prolog__clause), X)

				return
			}, 1)
			gen12296 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4head__abstraction), X)

				return
			}, 1)
			gen12297 := Call(__e, ShenFunc(symmapcan), gen12296, V929)

			gen12298 := Call(__e, ShenFunc(symmap), gen12295, gen12297)

			gen12299 := Call(__e, ShenFunc(symshen_4group__clauses), gen12298)

			__e.TailApply(ShenFunc(symmap), gen12294, gen12299)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog->shen"), gen12300)

		gen12326 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V931 := __args[0]
			_ = V931
			gen12324 := Call(__e, ShenFunc(symcons_2), V931)

			var gen12325 Obj
			if True == gen12324 {
				gen12321 := Call(__e, ShenFunc(symtl), V931)

				gen12322 := Call(__e, ShenFunc(symcons_2), gen12321)

				var gen12323 Obj
				if True == gen12322 {
					gen12317 := Call(__e, ShenFunc(symtl), V931)

					gen12318 := Call(__e, ShenFunc(symhd), gen12317)

					gen12319 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen12318)

					var gen12320 Obj
					if True == gen12319 {
						gen12313 := Call(__e, ShenFunc(symtl), V931)

						gen12314 := Call(__e, ShenFunc(symtl), gen12313)

						gen12315 := Call(__e, ShenFunc(symcons_2), gen12314)

						var gen12316 Obj
						if True == gen12315 {
							gen12309 := Call(__e, ShenFunc(symtl), V931)

							gen12310 := Call(__e, ShenFunc(symtl), gen12309)

							gen12311 := Call(__e, ShenFunc(symtl), gen12310)

							gen12312 := Call(__e, ShenFunc(sym_a), Nil, gen12311)

							if True == gen12312 {
								gen12316 = True
							} else {
								gen12316 = False
							}

						} else {
							gen12316 = False
						}
						if True == gen12316 {
							gen12320 = True
						} else {
							gen12320 = False
						}

					} else {
						gen12320 = False
					}
					if True == gen12320 {
						gen12323 = True
					} else {
						gen12323 = False
					}

				} else {
					gen12323 = False
				}
				if True == gen12323 {
					gen12325 = True
				} else {
					gen12325 = False
				}

			} else {
				gen12325 = False
			}
			if True == gen12325 {
				gen12301 := Call(__e, ShenFunc(symhd), V931)

				gen12302 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4s_1prolog__literal), X)

					return
				}, 1)
				gen12303 := Call(__e, ShenFunc(symtl), V931)

				gen12304 := Call(__e, ShenFunc(symtl), gen12303)

				gen12305 := Call(__e, ShenFunc(symhd), gen12304)

				gen12306 := Call(__e, ShenFunc(symmap), gen12302, gen12305)

				gen12307 := Call(__e, ShenFunc(symcons), gen12306, Nil)

				gen12308 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen12307)

				__e.TailApply(ShenFunc(symcons), gen12301, gen12308)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.s-prolog_clause"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog_clause"), gen12326)

		gen12393 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V933 := __args[0]
			_ = V933
			gen12391 := Call(__e, ShenFunc(symcons_2), V933)

			var gen12392 Obj
			if True == gen12391 {
				gen12388 := Call(__e, ShenFunc(symtl), V933)

				gen12389 := Call(__e, ShenFunc(symcons_2), gen12388)

				var gen12390 Obj
				if True == gen12389 {
					gen12384 := Call(__e, ShenFunc(symtl), V933)

					gen12385 := Call(__e, ShenFunc(symhd), gen12384)

					gen12386 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen12385)

					var gen12387 Obj
					if True == gen12386 {
						gen12380 := Call(__e, ShenFunc(symtl), V933)

						gen12381 := Call(__e, ShenFunc(symtl), gen12380)

						gen12382 := Call(__e, ShenFunc(symcons_2), gen12381)

						var gen12383 Obj
						if True == gen12382 {
							gen12375 := Call(__e, ShenFunc(symtl), V933)

							gen12376 := Call(__e, ShenFunc(symtl), gen12375)

							gen12377 := Call(__e, ShenFunc(symtl), gen12376)

							gen12378 := Call(__e, ShenFunc(sym_a), Nil, gen12377)

							var gen12379 Obj
							if True == gen12378 {
								gen12369 := MakeNative(func(__e Evaluator, __args ...Obj) {
									__ := __args[0]
									_ = __
									__e.Return(False)
									return
								}, 1)
								gen12373 := MakeNative(func(__e Evaluator, __args ...Obj) {
									gen12370 := Call(__e, ShenFunc(symhd), V933)

									gen12371 := Call(__e, ShenFunc(symshen_4complexity__head), gen12370)

									gen12372 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*maxcomplexity*"))

									__e.TailApply(ShenFunc(sym_5), gen12371, gen12372)

									return

								}, 0)
								gen12374 := Try(__e, gen12373).Catch(gen12369)
								if True == gen12374 {
									gen12379 = True
								} else {
									gen12379 = False
								}

							} else {
								gen12379 = False
							}
							if True == gen12379 {
								gen12383 = True
							} else {
								gen12383 = False
							}

						} else {
							gen12383 = False
						}
						if True == gen12383 {
							gen12387 = True
						} else {
							gen12387 = False
						}

					} else {
						gen12387 = False
					}
					if True == gen12387 {
						gen12390 = True
					} else {
						gen12390 = False
					}

				} else {
					gen12390 = False
				}
				if True == gen12390 {
					gen12392 = True
				} else {
					gen12392 = False
				}

			} else {
				gen12392 = False
			}
			if True == gen12392 {
				__e.TailApply(ShenFunc(symcons), V933, Nil)

				return
			} else {
				gen12367 := Call(__e, ShenFunc(symcons_2), V933)

				var gen12368 Obj
				if True == gen12367 {
					gen12364 := Call(__e, ShenFunc(symhd), V933)

					gen12365 := Call(__e, ShenFunc(symcons_2), gen12364)

					var gen12366 Obj
					if True == gen12365 {
						gen12361 := Call(__e, ShenFunc(symtl), V933)

						gen12362 := Call(__e, ShenFunc(symcons_2), gen12361)

						var gen12363 Obj
						if True == gen12362 {
							gen12357 := Call(__e, ShenFunc(symtl), V933)

							gen12358 := Call(__e, ShenFunc(symhd), gen12357)

							gen12359 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen12358)

							var gen12360 Obj
							if True == gen12359 {
								gen12353 := Call(__e, ShenFunc(symtl), V933)

								gen12354 := Call(__e, ShenFunc(symtl), gen12353)

								gen12355 := Call(__e, ShenFunc(symcons_2), gen12354)

								var gen12356 Obj
								if True == gen12355 {
									gen12349 := Call(__e, ShenFunc(symtl), V933)

									gen12350 := Call(__e, ShenFunc(symtl), gen12349)

									gen12351 := Call(__e, ShenFunc(symtl), gen12350)

									gen12352 := Call(__e, ShenFunc(sym_a), Nil, gen12351)

									if True == gen12352 {
										gen12356 = True
									} else {
										gen12356 = False
									}

								} else {
									gen12356 = False
								}
								if True == gen12356 {
									gen12360 = True
								} else {
									gen12360 = False
								}

							} else {
								gen12360 = False
							}
							if True == gen12360 {
								gen12363 = True
							} else {
								gen12363 = False
							}

						} else {
							gen12363 = False
						}
						if True == gen12363 {
							gen12366 = True
						} else {
							gen12366 = False
						}

					} else {
						gen12366 = False
					}
					if True == gen12366 {
						gen12368 = True
					} else {
						gen12368 = False
					}

				} else {
					gen12368 = False
				}
				if True == gen12368 {
					gen12327 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Y := __args[0]
						_ = Y
						__e.TailApply(ShenFunc(symgensym), MakeSymbol("V"))

						return
					}, 1)
					gen12328 := Call(__e, ShenFunc(symhd), V933)

					gen12329 := Call(__e, ShenFunc(symtl), gen12328)

					gen12330 := Call(__e, ShenFunc(symmap), gen12327, gen12329)

					Terms := gen12330
					gen12331 := Call(__e, ShenFunc(symhd), V933)

					gen12332 := Call(__e, ShenFunc(symtl), gen12331)

					gen12333 := Call(__e, ShenFunc(symshen_4remove__modes), gen12332)

					gen12334 := Call(__e, ShenFunc(symshen_4rcons__form), gen12333)

					XTerms := gen12334
					gen12335 := Call(__e, ShenFunc(symshen_4cons__form), Terms)

					gen12336 := Call(__e, ShenFunc(symcons), XTerms, Nil)

					gen12337 := Call(__e, ShenFunc(symcons), gen12335, gen12336)

					gen12338 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen12337)

					Literal := gen12338
					gen12339 := Call(__e, ShenFunc(symhd), V933)

					gen12340 := Call(__e, ShenFunc(symhd), gen12339)

					gen12341 := Call(__e, ShenFunc(symcons), gen12340, Terms)

					gen12342 := Call(__e, ShenFunc(symtl), V933)

					gen12343 := Call(__e, ShenFunc(symtl), gen12342)

					gen12344 := Call(__e, ShenFunc(symhd), gen12343)

					gen12345 := Call(__e, ShenFunc(symcons), Literal, gen12344)

					gen12346 := Call(__e, ShenFunc(symcons), gen12345, Nil)

					gen12347 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen12346)

					gen12348 := Call(__e, ShenFunc(symcons), gen12341, gen12347)

					Clause := gen12348
					__e.TailApply(ShenFunc(symcons), Clause, Nil)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.head_abstraction"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.head_abstraction"), gen12393)

		gen12398 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V939 := __args[0]
			_ = V939
			gen12397 := Call(__e, ShenFunc(symcons_2), V939)

			if True == gen12397 {
				gen12394 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4complexity), X)

					return
				}, 1)
				gen12395 := Call(__e, ShenFunc(symtl), V939)

				gen12396 := Call(__e, ShenFunc(symmap), gen12394, gen12395)

				__e.TailApply(ShenFunc(symshen_4safe_1product), gen12396)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.complexity_head"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.complexity_head"), gen12398)

		gen12399 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V942 := __args[0]
			_ = V942
			V943 := __args[1]
			_ = V943
			__e.TailApply(ShenFunc(sym_d), V942, V943)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.safe-multiply"), gen12399)

		gen12593 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V952 := __args[0]
			_ = V952
			gen12591 := Call(__e, ShenFunc(symcons_2), V952)

			var gen12592 Obj
			if True == gen12591 {
				gen12588 := Call(__e, ShenFunc(symhd), V952)

				gen12589 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12588)

				var gen12590 Obj
				if True == gen12589 {
					gen12585 := Call(__e, ShenFunc(symtl), V952)

					gen12586 := Call(__e, ShenFunc(symcons_2), gen12585)

					var gen12587 Obj
					if True == gen12586 {
						gen12581 := Call(__e, ShenFunc(symtl), V952)

						gen12582 := Call(__e, ShenFunc(symhd), gen12581)

						gen12583 := Call(__e, ShenFunc(symcons_2), gen12582)

						var gen12584 Obj
						if True == gen12583 {
							gen12576 := Call(__e, ShenFunc(symtl), V952)

							gen12577 := Call(__e, ShenFunc(symhd), gen12576)

							gen12578 := Call(__e, ShenFunc(symhd), gen12577)

							gen12579 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12578)

							var gen12580 Obj
							if True == gen12579 {
								gen12571 := Call(__e, ShenFunc(symtl), V952)

								gen12572 := Call(__e, ShenFunc(symhd), gen12571)

								gen12573 := Call(__e, ShenFunc(symtl), gen12572)

								gen12574 := Call(__e, ShenFunc(symcons_2), gen12573)

								var gen12575 Obj
								if True == gen12574 {
									gen12565 := Call(__e, ShenFunc(symtl), V952)

									gen12566 := Call(__e, ShenFunc(symhd), gen12565)

									gen12567 := Call(__e, ShenFunc(symtl), gen12566)

									gen12568 := Call(__e, ShenFunc(symtl), gen12567)

									gen12569 := Call(__e, ShenFunc(symcons_2), gen12568)

									var gen12570 Obj
									if True == gen12569 {
										gen12558 := Call(__e, ShenFunc(symtl), V952)

										gen12559 := Call(__e, ShenFunc(symhd), gen12558)

										gen12560 := Call(__e, ShenFunc(symtl), gen12559)

										gen12561 := Call(__e, ShenFunc(symtl), gen12560)

										gen12562 := Call(__e, ShenFunc(symtl), gen12561)

										gen12563 := Call(__e, ShenFunc(sym_a), Nil, gen12562)

										var gen12564 Obj
										if True == gen12563 {
											gen12554 := Call(__e, ShenFunc(symtl), V952)

											gen12555 := Call(__e, ShenFunc(symtl), gen12554)

											gen12556 := Call(__e, ShenFunc(symcons_2), gen12555)

											var gen12557 Obj
											if True == gen12556 {
												gen12550 := Call(__e, ShenFunc(symtl), V952)

												gen12551 := Call(__e, ShenFunc(symtl), gen12550)

												gen12552 := Call(__e, ShenFunc(symtl), gen12551)

												gen12553 := Call(__e, ShenFunc(sym_a), Nil, gen12552)

												if True == gen12553 {
													gen12557 = True
												} else {
													gen12557 = False
												}

											} else {
												gen12557 = False
											}
											if True == gen12557 {
												gen12564 = True
											} else {
												gen12564 = False
											}

										} else {
											gen12564 = False
										}
										if True == gen12564 {
											gen12570 = True
										} else {
											gen12570 = False
										}

									} else {
										gen12570 = False
									}
									if True == gen12570 {
										gen12575 = True
									} else {
										gen12575 = False
									}

								} else {
									gen12575 = False
								}
								if True == gen12575 {
									gen12580 = True
								} else {
									gen12580 = False
								}

							} else {
								gen12580 = False
							}
							if True == gen12580 {
								gen12584 = True
							} else {
								gen12584 = False
							}

						} else {
							gen12584 = False
						}
						if True == gen12584 {
							gen12587 = True
						} else {
							gen12587 = False
						}

					} else {
						gen12587 = False
					}
					if True == gen12587 {
						gen12590 = True
					} else {
						gen12590 = False
					}

				} else {
					gen12590 = False
				}
				if True == gen12590 {
					gen12592 = True
				} else {
					gen12592 = False
				}

			} else {
				gen12592 = False
			}
			if True == gen12592 {
				gen12548 := Call(__e, ShenFunc(symtl), V952)

				gen12549 := Call(__e, ShenFunc(symhd), gen12548)

				__e.TailApply(ShenFunc(symshen_4complexity), gen12549)

				return

			} else {
				gen12546 := Call(__e, ShenFunc(symcons_2), V952)

				var gen12547 Obj
				if True == gen12546 {
					gen12543 := Call(__e, ShenFunc(symhd), V952)

					gen12544 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12543)

					var gen12545 Obj
					if True == gen12544 {
						gen12540 := Call(__e, ShenFunc(symtl), V952)

						gen12541 := Call(__e, ShenFunc(symcons_2), gen12540)

						var gen12542 Obj
						if True == gen12541 {
							gen12536 := Call(__e, ShenFunc(symtl), V952)

							gen12537 := Call(__e, ShenFunc(symhd), gen12536)

							gen12538 := Call(__e, ShenFunc(symcons_2), gen12537)

							var gen12539 Obj
							if True == gen12538 {
								gen12532 := Call(__e, ShenFunc(symtl), V952)

								gen12533 := Call(__e, ShenFunc(symtl), gen12532)

								gen12534 := Call(__e, ShenFunc(symcons_2), gen12533)

								var gen12535 Obj
								if True == gen12534 {
									gen12527 := Call(__e, ShenFunc(symtl), V952)

									gen12528 := Call(__e, ShenFunc(symtl), gen12527)

									gen12529 := Call(__e, ShenFunc(symhd), gen12528)

									gen12530 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen12529)

									var gen12531 Obj
									if True == gen12530 {
										gen12523 := Call(__e, ShenFunc(symtl), V952)

										gen12524 := Call(__e, ShenFunc(symtl), gen12523)

										gen12525 := Call(__e, ShenFunc(symtl), gen12524)

										gen12526 := Call(__e, ShenFunc(sym_a), Nil, gen12525)

										if True == gen12526 {
											gen12531 = True
										} else {
											gen12531 = False
										}

									} else {
										gen12531 = False
									}
									if True == gen12531 {
										gen12535 = True
									} else {
										gen12535 = False
									}

								} else {
									gen12535 = False
								}
								if True == gen12535 {
									gen12539 = True
								} else {
									gen12539 = False
								}

							} else {
								gen12539 = False
							}
							if True == gen12539 {
								gen12542 = True
							} else {
								gen12542 = False
							}

						} else {
							gen12542 = False
						}
						if True == gen12542 {
							gen12545 = True
						} else {
							gen12545 = False
						}

					} else {
						gen12545 = False
					}
					if True == gen12545 {
						gen12547 = True
					} else {
						gen12547 = False
					}

				} else {
					gen12547 = False
				}
				if True == gen12547 {
					gen12506 := Call(__e, ShenFunc(symtl), V952)

					gen12507 := Call(__e, ShenFunc(symhd), gen12506)

					gen12508 := Call(__e, ShenFunc(symhd), gen12507)

					gen12509 := Call(__e, ShenFunc(symtl), V952)

					gen12510 := Call(__e, ShenFunc(symtl), gen12509)

					gen12511 := Call(__e, ShenFunc(symcons), gen12508, gen12510)

					gen12512 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12511)

					gen12513 := Call(__e, ShenFunc(symshen_4complexity), gen12512)

					gen12514 := Call(__e, ShenFunc(symtl), V952)

					gen12515 := Call(__e, ShenFunc(symhd), gen12514)

					gen12516 := Call(__e, ShenFunc(symtl), gen12515)

					gen12517 := Call(__e, ShenFunc(symtl), V952)

					gen12518 := Call(__e, ShenFunc(symtl), gen12517)

					gen12519 := Call(__e, ShenFunc(symcons), gen12516, gen12518)

					gen12520 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12519)

					gen12521 := Call(__e, ShenFunc(symshen_4complexity), gen12520)

					gen12522 := Call(__e, ShenFunc(symshen_4safe_1multiply), gen12513, gen12521)

					__e.TailApply(ShenFunc(symshen_4safe_1multiply), MakeNumber(2), gen12522)

					return

				} else {
					gen12504 := Call(__e, ShenFunc(symcons_2), V952)

					var gen12505 Obj
					if True == gen12504 {
						gen12501 := Call(__e, ShenFunc(symhd), V952)

						gen12502 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12501)

						var gen12503 Obj
						if True == gen12502 {
							gen12498 := Call(__e, ShenFunc(symtl), V952)

							gen12499 := Call(__e, ShenFunc(symcons_2), gen12498)

							var gen12500 Obj
							if True == gen12499 {
								gen12494 := Call(__e, ShenFunc(symtl), V952)

								gen12495 := Call(__e, ShenFunc(symhd), gen12494)

								gen12496 := Call(__e, ShenFunc(symcons_2), gen12495)

								var gen12497 Obj
								if True == gen12496 {
									gen12490 := Call(__e, ShenFunc(symtl), V952)

									gen12491 := Call(__e, ShenFunc(symtl), gen12490)

									gen12492 := Call(__e, ShenFunc(symcons_2), gen12491)

									var gen12493 Obj
									if True == gen12492 {
										gen12485 := Call(__e, ShenFunc(symtl), V952)

										gen12486 := Call(__e, ShenFunc(symtl), gen12485)

										gen12487 := Call(__e, ShenFunc(symhd), gen12486)

										gen12488 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen12487)

										var gen12489 Obj
										if True == gen12488 {
											gen12481 := Call(__e, ShenFunc(symtl), V952)

											gen12482 := Call(__e, ShenFunc(symtl), gen12481)

											gen12483 := Call(__e, ShenFunc(symtl), gen12482)

											gen12484 := Call(__e, ShenFunc(sym_a), Nil, gen12483)

											if True == gen12484 {
												gen12489 = True
											} else {
												gen12489 = False
											}

										} else {
											gen12489 = False
										}
										if True == gen12489 {
											gen12493 = True
										} else {
											gen12493 = False
										}

									} else {
										gen12493 = False
									}
									if True == gen12493 {
										gen12497 = True
									} else {
										gen12497 = False
									}

								} else {
									gen12497 = False
								}
								if True == gen12497 {
									gen12500 = True
								} else {
									gen12500 = False
								}

							} else {
								gen12500 = False
							}
							if True == gen12500 {
								gen12503 = True
							} else {
								gen12503 = False
							}

						} else {
							gen12503 = False
						}
						if True == gen12503 {
							gen12505 = True
						} else {
							gen12505 = False
						}

					} else {
						gen12505 = False
					}
					if True == gen12505 {
						gen12465 := Call(__e, ShenFunc(symtl), V952)

						gen12466 := Call(__e, ShenFunc(symhd), gen12465)

						gen12467 := Call(__e, ShenFunc(symhd), gen12466)

						gen12468 := Call(__e, ShenFunc(symtl), V952)

						gen12469 := Call(__e, ShenFunc(symtl), gen12468)

						gen12470 := Call(__e, ShenFunc(symcons), gen12467, gen12469)

						gen12471 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12470)

						gen12472 := Call(__e, ShenFunc(symshen_4complexity), gen12471)

						gen12473 := Call(__e, ShenFunc(symtl), V952)

						gen12474 := Call(__e, ShenFunc(symhd), gen12473)

						gen12475 := Call(__e, ShenFunc(symtl), gen12474)

						gen12476 := Call(__e, ShenFunc(symtl), V952)

						gen12477 := Call(__e, ShenFunc(symtl), gen12476)

						gen12478 := Call(__e, ShenFunc(symcons), gen12475, gen12477)

						gen12479 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12478)

						gen12480 := Call(__e, ShenFunc(symshen_4complexity), gen12479)

						__e.TailApply(ShenFunc(symshen_4safe_1multiply), gen12472, gen12480)

						return

					} else {
						gen12463 := Call(__e, ShenFunc(symcons_2), V952)

						var gen12464 Obj
						if True == gen12463 {
							gen12460 := Call(__e, ShenFunc(symhd), V952)

							gen12461 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12460)

							var gen12462 Obj
							if True == gen12461 {
								gen12457 := Call(__e, ShenFunc(symtl), V952)

								gen12458 := Call(__e, ShenFunc(symcons_2), gen12457)

								var gen12459 Obj
								if True == gen12458 {
									gen12453 := Call(__e, ShenFunc(symtl), V952)

									gen12454 := Call(__e, ShenFunc(symtl), gen12453)

									gen12455 := Call(__e, ShenFunc(symcons_2), gen12454)

									var gen12456 Obj
									if True == gen12455 {
										gen12448 := Call(__e, ShenFunc(symtl), V952)

										gen12449 := Call(__e, ShenFunc(symtl), gen12448)

										gen12450 := Call(__e, ShenFunc(symtl), gen12449)

										gen12451 := Call(__e, ShenFunc(sym_a), Nil, gen12450)

										var gen12452 Obj
										if True == gen12451 {
											gen12445 := Call(__e, ShenFunc(symtl), V952)

											gen12446 := Call(__e, ShenFunc(symhd), gen12445)

											gen12447 := Call(__e, ShenFunc(symvariable_2), gen12446)

											if True == gen12447 {
												gen12452 = True
											} else {
												gen12452 = False
											}

										} else {
											gen12452 = False
										}
										if True == gen12452 {
											gen12456 = True
										} else {
											gen12456 = False
										}

									} else {
										gen12456 = False
									}
									if True == gen12456 {
										gen12459 = True
									} else {
										gen12459 = False
									}

								} else {
									gen12459 = False
								}
								if True == gen12459 {
									gen12462 = True
								} else {
									gen12462 = False
								}

							} else {
								gen12462 = False
							}
							if True == gen12462 {
								gen12464 = True
							} else {
								gen12464 = False
							}

						} else {
							gen12464 = False
						}
						if True == gen12464 {
							__e.Return(MakeNumber(1))
							return
						} else {
							gen12443 := Call(__e, ShenFunc(symcons_2), V952)

							var gen12444 Obj
							if True == gen12443 {
								gen12440 := Call(__e, ShenFunc(symhd), V952)

								gen12441 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12440)

								var gen12442 Obj
								if True == gen12441 {
									gen12437 := Call(__e, ShenFunc(symtl), V952)

									gen12438 := Call(__e, ShenFunc(symcons_2), gen12437)

									var gen12439 Obj
									if True == gen12438 {
										gen12433 := Call(__e, ShenFunc(symtl), V952)

										gen12434 := Call(__e, ShenFunc(symtl), gen12433)

										gen12435 := Call(__e, ShenFunc(symcons_2), gen12434)

										var gen12436 Obj
										if True == gen12435 {
											gen12428 := Call(__e, ShenFunc(symtl), V952)

											gen12429 := Call(__e, ShenFunc(symtl), gen12428)

											gen12430 := Call(__e, ShenFunc(symhd), gen12429)

											gen12431 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen12430)

											var gen12432 Obj
											if True == gen12431 {
												gen12424 := Call(__e, ShenFunc(symtl), V952)

												gen12425 := Call(__e, ShenFunc(symtl), gen12424)

												gen12426 := Call(__e, ShenFunc(symtl), gen12425)

												gen12427 := Call(__e, ShenFunc(sym_a), Nil, gen12426)

												if True == gen12427 {
													gen12432 = True
												} else {
													gen12432 = False
												}

											} else {
												gen12432 = False
											}
											if True == gen12432 {
												gen12436 = True
											} else {
												gen12436 = False
											}

										} else {
											gen12436 = False
										}
										if True == gen12436 {
											gen12439 = True
										} else {
											gen12439 = False
										}

									} else {
										gen12439 = False
									}
									if True == gen12439 {
										gen12442 = True
									} else {
										gen12442 = False
									}

								} else {
									gen12442 = False
								}
								if True == gen12442 {
									gen12444 = True
								} else {
									gen12444 = False
								}

							} else {
								gen12444 = False
							}
							if True == gen12444 {
								__e.Return(MakeNumber(2))
								return
							} else {
								gen12422 := Call(__e, ShenFunc(symcons_2), V952)

								var gen12423 Obj
								if True == gen12422 {
									gen12419 := Call(__e, ShenFunc(symhd), V952)

									gen12420 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12419)

									var gen12421 Obj
									if True == gen12420 {
										gen12416 := Call(__e, ShenFunc(symtl), V952)

										gen12417 := Call(__e, ShenFunc(symcons_2), gen12416)

										var gen12418 Obj
										if True == gen12417 {
											gen12412 := Call(__e, ShenFunc(symtl), V952)

											gen12413 := Call(__e, ShenFunc(symtl), gen12412)

											gen12414 := Call(__e, ShenFunc(symcons_2), gen12413)

											var gen12415 Obj
											if True == gen12414 {
												gen12407 := Call(__e, ShenFunc(symtl), V952)

												gen12408 := Call(__e, ShenFunc(symtl), gen12407)

												gen12409 := Call(__e, ShenFunc(symhd), gen12408)

												gen12410 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen12409)

												var gen12411 Obj
												if True == gen12410 {
													gen12403 := Call(__e, ShenFunc(symtl), V952)

													gen12404 := Call(__e, ShenFunc(symtl), gen12403)

													gen12405 := Call(__e, ShenFunc(symtl), gen12404)

													gen12406 := Call(__e, ShenFunc(sym_a), Nil, gen12405)

													if True == gen12406 {
														gen12411 = True
													} else {
														gen12411 = False
													}

												} else {
													gen12411 = False
												}
												if True == gen12411 {
													gen12415 = True
												} else {
													gen12415 = False
												}

											} else {
												gen12415 = False
											}
											if True == gen12415 {
												gen12418 = True
											} else {
												gen12418 = False
											}

										} else {
											gen12418 = False
										}
										if True == gen12418 {
											gen12421 = True
										} else {
											gen12421 = False
										}

									} else {
										gen12421 = False
									}
									if True == gen12421 {
										gen12423 = True
									} else {
										gen12423 = False
									}

								} else {
									gen12423 = False
								}
								if True == gen12423 {
									__e.Return(MakeNumber(1))
									return
								} else {
									gen12400 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

									gen12401 := Call(__e, ShenFunc(symcons), V952, gen12400)

									gen12402 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen12401)

									__e.TailApply(ShenFunc(symshen_4complexity), gen12402)

									return

								}

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.complexity"), gen12593)

		gen12599 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V954 := __args[0]
			_ = V954
			gen12598 := Call(__e, ShenFunc(sym_a), Nil, V954)

			if True == gen12598 {
				__e.Return(MakeNumber(1))
				return
			} else {
				gen12597 := Call(__e, ShenFunc(symcons_2), V954)

				if True == gen12597 {
					gen12594 := Call(__e, ShenFunc(symhd), V954)

					gen12595 := Call(__e, ShenFunc(symtl), V954)

					gen12596 := Call(__e, ShenFunc(symshen_4safe_1product), gen12595)

					__e.TailApply(ShenFunc(symshen_4safe_1multiply), gen12594, gen12596)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.safe-product"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.safe-product"), gen12599)

		gen12679 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V956 := __args[0]
			_ = V956
			gen12677 := Call(__e, ShenFunc(symcons_2), V956)

			var gen12678 Obj
			if True == gen12677 {
				gen12674 := Call(__e, ShenFunc(symhd), V956)

				gen12675 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen12674)

				var gen12676 Obj
				if True == gen12675 {
					gen12671 := Call(__e, ShenFunc(symtl), V956)

					gen12672 := Call(__e, ShenFunc(symcons_2), gen12671)

					var gen12673 Obj
					if True == gen12672 {
						gen12667 := Call(__e, ShenFunc(symtl), V956)

						gen12668 := Call(__e, ShenFunc(symtl), gen12667)

						gen12669 := Call(__e, ShenFunc(symcons_2), gen12668)

						var gen12670 Obj
						if True == gen12669 {
							gen12663 := Call(__e, ShenFunc(symtl), V956)

							gen12664 := Call(__e, ShenFunc(symtl), gen12663)

							gen12665 := Call(__e, ShenFunc(symtl), gen12664)

							gen12666 := Call(__e, ShenFunc(sym_a), Nil, gen12665)

							if True == gen12666 {
								gen12670 = True
							} else {
								gen12670 = False
							}

						} else {
							gen12670 = False
						}
						if True == gen12670 {
							gen12673 = True
						} else {
							gen12673 = False
						}

					} else {
						gen12673 = False
					}
					if True == gen12673 {
						gen12676 = True
					} else {
						gen12676 = False
					}

				} else {
					gen12676 = False
				}
				if True == gen12676 {
					gen12678 = True
				} else {
					gen12678 = False
				}

			} else {
				gen12678 = False
			}
			if True == gen12678 {
				gen12655 := Call(__e, ShenFunc(symtl), V956)

				gen12656 := Call(__e, ShenFunc(symhd), gen12655)

				gen12657 := Call(__e, ShenFunc(symtl), V956)

				gen12658 := Call(__e, ShenFunc(symtl), gen12657)

				gen12659 := Call(__e, ShenFunc(symhd), gen12658)

				gen12660 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12659, MakeSymbol("ProcessN"))

				gen12661 := Call(__e, ShenFunc(symcons), gen12660, Nil)

				gen12662 := Call(__e, ShenFunc(symcons), gen12656, gen12661)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen12662)

				return

			} else {
				gen12653 := Call(__e, ShenFunc(symcons_2), V956)

				var gen12654 Obj
				if True == gen12653 {
					gen12650 := Call(__e, ShenFunc(symhd), V956)

					gen12651 := Call(__e, ShenFunc(sym_a), MakeSymbol("when"), gen12650)

					var gen12652 Obj
					if True == gen12651 {
						gen12647 := Call(__e, ShenFunc(symtl), V956)

						gen12648 := Call(__e, ShenFunc(symcons_2), gen12647)

						var gen12649 Obj
						if True == gen12648 {
							gen12644 := Call(__e, ShenFunc(symtl), V956)

							gen12645 := Call(__e, ShenFunc(symtl), gen12644)

							gen12646 := Call(__e, ShenFunc(sym_a), Nil, gen12645)

							if True == gen12646 {
								gen12649 = True
							} else {
								gen12649 = False
							}

						} else {
							gen12649 = False
						}
						if True == gen12649 {
							gen12652 = True
						} else {
							gen12652 = False
						}

					} else {
						gen12652 = False
					}
					if True == gen12652 {
						gen12654 = True
					} else {
						gen12654 = False
					}

				} else {
					gen12654 = False
				}
				if True == gen12654 {
					gen12640 := Call(__e, ShenFunc(symtl), V956)

					gen12641 := Call(__e, ShenFunc(symhd), gen12640)

					gen12642 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12641, MakeSymbol("ProcessN"))

					gen12643 := Call(__e, ShenFunc(symcons), gen12642, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen12643)

					return

				} else {
					gen12638 := Call(__e, ShenFunc(symcons_2), V956)

					var gen12639 Obj
					if True == gen12638 {
						gen12635 := Call(__e, ShenFunc(symhd), V956)

						gen12636 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen12635)

						var gen12637 Obj
						if True == gen12636 {
							gen12632 := Call(__e, ShenFunc(symtl), V956)

							gen12633 := Call(__e, ShenFunc(symcons_2), gen12632)

							var gen12634 Obj
							if True == gen12633 {
								gen12628 := Call(__e, ShenFunc(symtl), V956)

								gen12629 := Call(__e, ShenFunc(symtl), gen12628)

								gen12630 := Call(__e, ShenFunc(symcons_2), gen12629)

								var gen12631 Obj
								if True == gen12630 {
									gen12624 := Call(__e, ShenFunc(symtl), V956)

									gen12625 := Call(__e, ShenFunc(symtl), gen12624)

									gen12626 := Call(__e, ShenFunc(symtl), gen12625)

									gen12627 := Call(__e, ShenFunc(sym_a), Nil, gen12626)

									if True == gen12627 {
										gen12631 = True
									} else {
										gen12631 = False
									}

								} else {
									gen12631 = False
								}
								if True == gen12631 {
									gen12634 = True
								} else {
									gen12634 = False
								}

							} else {
								gen12634 = False
							}
							if True == gen12634 {
								gen12637 = True
							} else {
								gen12637 = False
							}

						} else {
							gen12637 = False
						}
						if True == gen12637 {
							gen12639 = True
						} else {
							gen12639 = False
						}

					} else {
						gen12639 = False
					}
					if True == gen12639 {
						gen12616 := Call(__e, ShenFunc(symtl), V956)

						gen12617 := Call(__e, ShenFunc(symhd), gen12616)

						gen12618 := Call(__e, ShenFunc(symtl), V956)

						gen12619 := Call(__e, ShenFunc(symtl), gen12618)

						gen12620 := Call(__e, ShenFunc(symhd), gen12619)

						gen12621 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12620, MakeSymbol("ProcessN"))

						gen12622 := Call(__e, ShenFunc(symcons), gen12621, Nil)

						gen12623 := Call(__e, ShenFunc(symcons), gen12617, gen12622)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen12623)

						return

					} else {
						gen12614 := Call(__e, ShenFunc(symcons_2), V956)

						var gen12615 Obj
						if True == gen12614 {
							gen12611 := Call(__e, ShenFunc(symhd), V956)

							gen12612 := Call(__e, ShenFunc(sym_a), MakeSymbol("fwhen"), gen12611)

							var gen12613 Obj
							if True == gen12612 {
								gen12608 := Call(__e, ShenFunc(symtl), V956)

								gen12609 := Call(__e, ShenFunc(symcons_2), gen12608)

								var gen12610 Obj
								if True == gen12609 {
									gen12605 := Call(__e, ShenFunc(symtl), V956)

									gen12606 := Call(__e, ShenFunc(symtl), gen12605)

									gen12607 := Call(__e, ShenFunc(sym_a), Nil, gen12606)

									if True == gen12607 {
										gen12610 = True
									} else {
										gen12610 = False
									}

								} else {
									gen12610 = False
								}
								if True == gen12610 {
									gen12613 = True
								} else {
									gen12613 = False
								}

							} else {
								gen12613 = False
							}
							if True == gen12613 {
								gen12615 = True
							} else {
								gen12615 = False
							}

						} else {
							gen12615 = False
						}
						if True == gen12615 {
							gen12601 := Call(__e, ShenFunc(symtl), V956)

							gen12602 := Call(__e, ShenFunc(symhd), gen12601)

							gen12603 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12602, MakeSymbol("ProcessN"))

							gen12604 := Call(__e, ShenFunc(symcons), gen12603, Nil)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen12604)

							return

						} else {
							gen12600 := Call(__e, ShenFunc(symcons_2), V956)

							if True == gen12600 {
								__e.Return(V956)
								return
							} else {
								__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.s-prolog_literal"))

								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.s-prolog_literal"), gen12679)

		gen12748 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V963 := __args[0]
			_ = V963
			V964 := __args[1]
			_ = V964
			gen12747 := Call(__e, ShenFunc(symvariable_2), V963)

			if True == gen12747 {
				gen12745 := Call(__e, ShenFunc(symcons), V964, Nil)

				gen12746 := Call(__e, ShenFunc(symcons), V963, gen12745)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.deref"), gen12746)

				return

			} else {
				gen12743 := Call(__e, ShenFunc(symcons_2), V963)

				var gen12744 Obj
				if True == gen12743 {
					gen12740 := Call(__e, ShenFunc(symhd), V963)

					gen12741 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen12740)

					var gen12742 Obj
					if True == gen12741 {
						gen12737 := Call(__e, ShenFunc(symtl), V963)

						gen12738 := Call(__e, ShenFunc(symcons_2), gen12737)

						var gen12739 Obj
						if True == gen12738 {
							gen12733 := Call(__e, ShenFunc(symtl), V963)

							gen12734 := Call(__e, ShenFunc(symtl), gen12733)

							gen12735 := Call(__e, ShenFunc(symcons_2), gen12734)

							var gen12736 Obj
							if True == gen12735 {
								gen12729 := Call(__e, ShenFunc(symtl), V963)

								gen12730 := Call(__e, ShenFunc(symtl), gen12729)

								gen12731 := Call(__e, ShenFunc(symtl), gen12730)

								gen12732 := Call(__e, ShenFunc(sym_a), Nil, gen12731)

								if True == gen12732 {
									gen12736 = True
								} else {
									gen12736 = False
								}

							} else {
								gen12736 = False
							}
							if True == gen12736 {
								gen12739 = True
							} else {
								gen12739 = False
							}

						} else {
							gen12739 = False
						}
						if True == gen12739 {
							gen12742 = True
						} else {
							gen12742 = False
						}

					} else {
						gen12742 = False
					}
					if True == gen12742 {
						gen12744 = True
					} else {
						gen12744 = False
					}

				} else {
					gen12744 = False
				}
				if True == gen12744 {
					gen12721 := Call(__e, ShenFunc(symtl), V963)

					gen12722 := Call(__e, ShenFunc(symhd), gen12721)

					gen12723 := Call(__e, ShenFunc(symtl), V963)

					gen12724 := Call(__e, ShenFunc(symtl), gen12723)

					gen12725 := Call(__e, ShenFunc(symhd), gen12724)

					gen12726 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12725, V964)

					gen12727 := Call(__e, ShenFunc(symcons), gen12726, Nil)

					gen12728 := Call(__e, ShenFunc(symcons), gen12722, gen12727)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen12728)

					return

				} else {
					gen12719 := Call(__e, ShenFunc(symcons_2), V963)

					var gen12720 Obj
					if True == gen12719 {
						gen12716 := Call(__e, ShenFunc(symhd), V963)

						gen12717 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen12716)

						var gen12718 Obj
						if True == gen12717 {
							gen12713 := Call(__e, ShenFunc(symtl), V963)

							gen12714 := Call(__e, ShenFunc(symcons_2), gen12713)

							var gen12715 Obj
							if True == gen12714 {
								gen12709 := Call(__e, ShenFunc(symtl), V963)

								gen12710 := Call(__e, ShenFunc(symtl), gen12709)

								gen12711 := Call(__e, ShenFunc(symcons_2), gen12710)

								var gen12712 Obj
								if True == gen12711 {
									gen12704 := Call(__e, ShenFunc(symtl), V963)

									gen12705 := Call(__e, ShenFunc(symtl), gen12704)

									gen12706 := Call(__e, ShenFunc(symtl), gen12705)

									gen12707 := Call(__e, ShenFunc(symcons_2), gen12706)

									var gen12708 Obj
									if True == gen12707 {
										gen12699 := Call(__e, ShenFunc(symtl), V963)

										gen12700 := Call(__e, ShenFunc(symtl), gen12699)

										gen12701 := Call(__e, ShenFunc(symtl), gen12700)

										gen12702 := Call(__e, ShenFunc(symtl), gen12701)

										gen12703 := Call(__e, ShenFunc(sym_a), Nil, gen12702)

										if True == gen12703 {
											gen12708 = True
										} else {
											gen12708 = False
										}

									} else {
										gen12708 = False
									}
									if True == gen12708 {
										gen12712 = True
									} else {
										gen12712 = False
									}

								} else {
									gen12712 = False
								}
								if True == gen12712 {
									gen12715 = True
								} else {
									gen12715 = False
								}

							} else {
								gen12715 = False
							}
							if True == gen12715 {
								gen12718 = True
							} else {
								gen12718 = False
							}

						} else {
							gen12718 = False
						}
						if True == gen12718 {
							gen12720 = True
						} else {
							gen12720 = False
						}

					} else {
						gen12720 = False
					}
					if True == gen12720 {
						gen12685 := Call(__e, ShenFunc(symtl), V963)

						gen12686 := Call(__e, ShenFunc(symhd), gen12685)

						gen12687 := Call(__e, ShenFunc(symtl), V963)

						gen12688 := Call(__e, ShenFunc(symtl), gen12687)

						gen12689 := Call(__e, ShenFunc(symhd), gen12688)

						gen12690 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12689, V964)

						gen12691 := Call(__e, ShenFunc(symtl), V963)

						gen12692 := Call(__e, ShenFunc(symtl), gen12691)

						gen12693 := Call(__e, ShenFunc(symtl), gen12692)

						gen12694 := Call(__e, ShenFunc(symhd), gen12693)

						gen12695 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12694, V964)

						gen12696 := Call(__e, ShenFunc(symcons), gen12695, Nil)

						gen12697 := Call(__e, ShenFunc(symcons), gen12690, gen12696)

						gen12698 := Call(__e, ShenFunc(symcons), gen12686, gen12697)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen12698)

						return

					} else {
						gen12684 := Call(__e, ShenFunc(symcons_2), V963)

						if True == gen12684 {
							gen12680 := Call(__e, ShenFunc(symhd), V963)

							gen12681 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12680, V964)

							gen12682 := Call(__e, ShenFunc(symtl), V963)

							gen12683 := Call(__e, ShenFunc(symshen_4insert_1deref), gen12682, V964)

							__e.TailApply(ShenFunc(symcons), gen12681, gen12683)

							return

						} else {
							__e.Return(V963)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-deref"), gen12748)

		gen12817 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V971 := __args[0]
			_ = V971
			V972 := __args[1]
			_ = V972
			gen12816 := Call(__e, ShenFunc(symvariable_2), V971)

			if True == gen12816 {
				gen12814 := Call(__e, ShenFunc(symcons), V972, Nil)

				gen12815 := Call(__e, ShenFunc(symcons), V971, gen12814)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.lazyderef"), gen12815)

				return

			} else {
				gen12812 := Call(__e, ShenFunc(symcons_2), V971)

				var gen12813 Obj
				if True == gen12812 {
					gen12809 := Call(__e, ShenFunc(symhd), V971)

					gen12810 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen12809)

					var gen12811 Obj
					if True == gen12810 {
						gen12806 := Call(__e, ShenFunc(symtl), V971)

						gen12807 := Call(__e, ShenFunc(symcons_2), gen12806)

						var gen12808 Obj
						if True == gen12807 {
							gen12802 := Call(__e, ShenFunc(symtl), V971)

							gen12803 := Call(__e, ShenFunc(symtl), gen12802)

							gen12804 := Call(__e, ShenFunc(symcons_2), gen12803)

							var gen12805 Obj
							if True == gen12804 {
								gen12798 := Call(__e, ShenFunc(symtl), V971)

								gen12799 := Call(__e, ShenFunc(symtl), gen12798)

								gen12800 := Call(__e, ShenFunc(symtl), gen12799)

								gen12801 := Call(__e, ShenFunc(sym_a), Nil, gen12800)

								if True == gen12801 {
									gen12805 = True
								} else {
									gen12805 = False
								}

							} else {
								gen12805 = False
							}
							if True == gen12805 {
								gen12808 = True
							} else {
								gen12808 = False
							}

						} else {
							gen12808 = False
						}
						if True == gen12808 {
							gen12811 = True
						} else {
							gen12811 = False
						}

					} else {
						gen12811 = False
					}
					if True == gen12811 {
						gen12813 = True
					} else {
						gen12813 = False
					}

				} else {
					gen12813 = False
				}
				if True == gen12813 {
					gen12790 := Call(__e, ShenFunc(symtl), V971)

					gen12791 := Call(__e, ShenFunc(symhd), gen12790)

					gen12792 := Call(__e, ShenFunc(symtl), V971)

					gen12793 := Call(__e, ShenFunc(symtl), gen12792)

					gen12794 := Call(__e, ShenFunc(symhd), gen12793)

					gen12795 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12794, V972)

					gen12796 := Call(__e, ShenFunc(symcons), gen12795, Nil)

					gen12797 := Call(__e, ShenFunc(symcons), gen12791, gen12796)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen12797)

					return

				} else {
					gen12788 := Call(__e, ShenFunc(symcons_2), V971)

					var gen12789 Obj
					if True == gen12788 {
						gen12785 := Call(__e, ShenFunc(symhd), V971)

						gen12786 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen12785)

						var gen12787 Obj
						if True == gen12786 {
							gen12782 := Call(__e, ShenFunc(symtl), V971)

							gen12783 := Call(__e, ShenFunc(symcons_2), gen12782)

							var gen12784 Obj
							if True == gen12783 {
								gen12778 := Call(__e, ShenFunc(symtl), V971)

								gen12779 := Call(__e, ShenFunc(symtl), gen12778)

								gen12780 := Call(__e, ShenFunc(symcons_2), gen12779)

								var gen12781 Obj
								if True == gen12780 {
									gen12773 := Call(__e, ShenFunc(symtl), V971)

									gen12774 := Call(__e, ShenFunc(symtl), gen12773)

									gen12775 := Call(__e, ShenFunc(symtl), gen12774)

									gen12776 := Call(__e, ShenFunc(symcons_2), gen12775)

									var gen12777 Obj
									if True == gen12776 {
										gen12768 := Call(__e, ShenFunc(symtl), V971)

										gen12769 := Call(__e, ShenFunc(symtl), gen12768)

										gen12770 := Call(__e, ShenFunc(symtl), gen12769)

										gen12771 := Call(__e, ShenFunc(symtl), gen12770)

										gen12772 := Call(__e, ShenFunc(sym_a), Nil, gen12771)

										if True == gen12772 {
											gen12777 = True
										} else {
											gen12777 = False
										}

									} else {
										gen12777 = False
									}
									if True == gen12777 {
										gen12781 = True
									} else {
										gen12781 = False
									}

								} else {
									gen12781 = False
								}
								if True == gen12781 {
									gen12784 = True
								} else {
									gen12784 = False
								}

							} else {
								gen12784 = False
							}
							if True == gen12784 {
								gen12787 = True
							} else {
								gen12787 = False
							}

						} else {
							gen12787 = False
						}
						if True == gen12787 {
							gen12789 = True
						} else {
							gen12789 = False
						}

					} else {
						gen12789 = False
					}
					if True == gen12789 {
						gen12754 := Call(__e, ShenFunc(symtl), V971)

						gen12755 := Call(__e, ShenFunc(symhd), gen12754)

						gen12756 := Call(__e, ShenFunc(symtl), V971)

						gen12757 := Call(__e, ShenFunc(symtl), gen12756)

						gen12758 := Call(__e, ShenFunc(symhd), gen12757)

						gen12759 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12758, V972)

						gen12760 := Call(__e, ShenFunc(symtl), V971)

						gen12761 := Call(__e, ShenFunc(symtl), gen12760)

						gen12762 := Call(__e, ShenFunc(symtl), gen12761)

						gen12763 := Call(__e, ShenFunc(symhd), gen12762)

						gen12764 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12763, V972)

						gen12765 := Call(__e, ShenFunc(symcons), gen12764, Nil)

						gen12766 := Call(__e, ShenFunc(symcons), gen12759, gen12765)

						gen12767 := Call(__e, ShenFunc(symcons), gen12755, gen12766)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen12767)

						return

					} else {
						gen12753 := Call(__e, ShenFunc(symcons_2), V971)

						if True == gen12753 {
							gen12749 := Call(__e, ShenFunc(symhd), V971)

							gen12750 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12749, V972)

							gen12751 := Call(__e, ShenFunc(symtl), V971)

							gen12752 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen12751, V972)

							__e.TailApply(ShenFunc(symcons), gen12750, gen12752)

							return

						} else {
							__e.Return(V971)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-lazyderef"), gen12817)

		gen12825 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V974 := __args[0]
			_ = V974
			gen12824 := Call(__e, ShenFunc(sym_a), Nil, V974)

			if True == gen12824 {
				__e.Return(Nil)
				return
			} else {
				gen12823 := Call(__e, ShenFunc(symcons_2), V974)

				if True == gen12823 {
					gen12819 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						gen12818 := Call(__e, ShenFunc(symhd), V974)

						__e.TailApply(ShenFunc(symshen_4same__predicate_2), gen12818, X)

						return

					}, 1)
					gen12820 := Call(__e, ShenFunc(symshen_4collect), gen12819, V974)

					Group := gen12820
					gen12821 := Call(__e, ShenFunc(symdifference), V974, Group)

					Rest := gen12821
					gen12822 := Call(__e, ShenFunc(symshen_4group__clauses), Rest)

					__e.TailApply(ShenFunc(symcons), Group, gen12822)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.group_clauses"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.group_clauses"), gen12825)

		gen12834 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V979 := __args[0]
			_ = V979
			V980 := __args[1]
			_ = V980
			gen12833 := Call(__e, ShenFunc(sym_a), Nil, V980)

			if True == gen12833 {
				__e.Return(Nil)
				return
			} else {
				gen12832 := Call(__e, ShenFunc(symcons_2), V980)

				if True == gen12832 {
					gen12830 := Call(__e, ShenFunc(symhd), V980)

					gen12831 := Call(__e, V979, gen12830)

					if True == gen12831 {
						gen12827 := Call(__e, ShenFunc(symhd), V980)

						gen12828 := Call(__e, ShenFunc(symtl), V980)

						gen12829 := Call(__e, ShenFunc(symshen_4collect), V979, gen12828)

						__e.TailApply(ShenFunc(symcons), gen12827, gen12829)

						return

					} else {
						gen12826 := Call(__e, ShenFunc(symtl), V980)

						__e.TailApply(ShenFunc(symshen_4collect), V979, gen12826)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.collect"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.collect"), gen12834)

		gen12848 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V999 := __args[0]
			_ = V999
			V1000 := __args[1]
			_ = V1000
			gen12846 := Call(__e, ShenFunc(symcons_2), V999)

			var gen12847 Obj
			if True == gen12846 {
				gen12843 := Call(__e, ShenFunc(symhd), V999)

				gen12844 := Call(__e, ShenFunc(symcons_2), gen12843)

				var gen12845 Obj
				if True == gen12844 {
					gen12841 := Call(__e, ShenFunc(symcons_2), V1000)

					var gen12842 Obj
					if True == gen12841 {
						gen12839 := Call(__e, ShenFunc(symhd), V1000)

						gen12840 := Call(__e, ShenFunc(symcons_2), gen12839)

						if True == gen12840 {
							gen12842 = True
						} else {
							gen12842 = False
						}

					} else {
						gen12842 = False
					}
					if True == gen12842 {
						gen12845 = True
					} else {
						gen12845 = False
					}

				} else {
					gen12845 = False
				}
				if True == gen12845 {
					gen12847 = True
				} else {
					gen12847 = False
				}

			} else {
				gen12847 = False
			}
			if True == gen12847 {
				gen12835 := Call(__e, ShenFunc(symhd), V999)

				gen12836 := Call(__e, ShenFunc(symhd), gen12835)

				gen12837 := Call(__e, ShenFunc(symhd), V1000)

				gen12838 := Call(__e, ShenFunc(symhd), gen12837)

				__e.TailApply(ShenFunc(sym_a), gen12836, gen12838)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.same_predicate?"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.same_predicate?"), gen12848)

		gen12851 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1002 := __args[0]
			_ = V1002
			gen12849 := Call(__e, ShenFunc(symshen_4procedure__name), V1002)

			F := gen12849
			gen12850 := Call(__e, ShenFunc(symshen_4clauses_1to_1shen), F, V1002)

			Shen := gen12850
			__e.Return(Shen)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_prolog_procedure"), gen12851)

		gen12862 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1016 := __args[0]
			_ = V1016
			gen12860 := Call(__e, ShenFunc(symcons_2), V1016)

			var gen12861 Obj
			if True == gen12860 {
				gen12857 := Call(__e, ShenFunc(symhd), V1016)

				gen12858 := Call(__e, ShenFunc(symcons_2), gen12857)

				var gen12859 Obj
				if True == gen12858 {
					gen12854 := Call(__e, ShenFunc(symhd), V1016)

					gen12855 := Call(__e, ShenFunc(symhd), gen12854)

					gen12856 := Call(__e, ShenFunc(symcons_2), gen12855)

					if True == gen12856 {
						gen12859 = True
					} else {
						gen12859 = False
					}

				} else {
					gen12859 = False
				}
				if True == gen12859 {
					gen12861 = True
				} else {
					gen12861 = False
				}

			} else {
				gen12861 = False
			}
			if True == gen12861 {
				gen12852 := Call(__e, ShenFunc(symhd), V1016)

				gen12853 := Call(__e, ShenFunc(symhd), gen12852)

				__e.TailApply(ShenFunc(symhd), gen12853)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.procedure_name"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.procedure_name"), gen12862)

		gen12883 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1019 := __args[0]
			_ = V1019
			V1020 := __args[1]
			_ = V1020
			gen12863 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4linearise_1clause), X)

				return
			}, 1)
			gen12864 := Call(__e, ShenFunc(symmap), gen12863, V1020)

			Linear := gen12864
			gen12865 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symhead), X)

				return
			}, 1)
			gen12866 := Call(__e, ShenFunc(symmap), gen12865, V1020)

			gen12867 := Call(__e, ShenFunc(symshen_4prolog_1aritycheck), V1019, gen12866)

			Arity := gen12867
			gen12868 := Call(__e, ShenFunc(symshen_4parameters), Arity)

			Parameters := gen12868
			gen12869 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4aum), X, Parameters)

				return
			}, 1)
			gen12870 := Call(__e, ShenFunc(symmap), gen12869, Linear)

			AUM__instructions := gen12870
			gen12871 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4aum__to__shen), X)

				return
			}, 1)
			gen12872 := Call(__e, ShenFunc(symmap), gen12871, AUM__instructions)

			gen12873 := Call(__e, ShenFunc(symshen_4nest_1disjunct), gen12872)

			gen12874 := Call(__e, ShenFunc(symshen_4catch_1cut), gen12873)

			Code := gen12874
			gen12875 := Call(__e, ShenFunc(symcons), MakeSymbol("Continuation"), Nil)

			gen12876 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), gen12875)

			gen12877 := Call(__e, ShenFunc(symcons), Code, Nil)

			gen12878 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen12877)

			gen12879 := Call(__e, ShenFunc(symappend), gen12876, gen12878)

			gen12880 := Call(__e, ShenFunc(symappend), Parameters, gen12879)

			gen12881 := Call(__e, ShenFunc(symcons), V1019, gen12880)

			gen12882 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen12881)

			ShenDef := gen12882
			__e.Return(ShenDef)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clauses-to-shen"), gen12883)

		gen12893 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1022 := __args[0]
			_ = V1022
			gen12891 := Call(__e, ShenFunc(symshen_4occurs_2), MakeSymbol("cut"), V1022)

			gen12892 := Call(__e, ShenFunc(symnot), gen12891)

			if True == gen12892 {
				__e.Return(V1022)
				return
			} else {
				gen12884 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.catchpoint"), Nil)

				gen12885 := Call(__e, ShenFunc(symcons), V1022, Nil)

				gen12886 := Call(__e, ShenFunc(symcons), MakeSymbol("Throwcontrol"), gen12885)

				gen12887 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.cutpoint"), gen12886)

				gen12888 := Call(__e, ShenFunc(symcons), gen12887, Nil)

				gen12889 := Call(__e, ShenFunc(symcons), gen12884, gen12888)

				gen12890 := Call(__e, ShenFunc(symcons), MakeSymbol("Throwcontrol"), gen12889)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen12890)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.catch-cut"), gen12893)

		gen12897 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen12894 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*catch*"))

			gen12895 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen12894)

			gen12896 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*catch*"), gen12895)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.catchpoint!"), gen12896)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.catchpoint"), gen12897)

		gen12899 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1030 := __args[0]
			_ = V1030
			V1031 := __args[1]
			_ = V1031
			gen12898 := Call(__e, ShenFunc(sym_a), V1031, V1030)

			if True == gen12898 {
				__e.Return(False)
				return
			} else {
				__e.Return(V1031)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cutpoint"), gen12899)

		gen12908 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1033 := __args[0]
			_ = V1033
			gen12906 := Call(__e, ShenFunc(symcons_2), V1033)

			var gen12907 Obj
			if True == gen12906 {
				gen12904 := Call(__e, ShenFunc(symtl), V1033)

				gen12905 := Call(__e, ShenFunc(sym_a), Nil, gen12904)

				if True == gen12905 {
					gen12907 = True
				} else {
					gen12907 = False
				}

			} else {
				gen12907 = False
			}
			if True == gen12907 {
				__e.TailApply(ShenFunc(symhd), V1033)

				return
			} else {
				gen12903 := Call(__e, ShenFunc(symcons_2), V1033)

				if True == gen12903 {
					gen12900 := Call(__e, ShenFunc(symhd), V1033)

					gen12901 := Call(__e, ShenFunc(symtl), V1033)

					gen12902 := Call(__e, ShenFunc(symshen_4nest_1disjunct), gen12901)

					__e.TailApply(ShenFunc(symshen_4lisp_1or), gen12900, gen12902)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.nest-disjunct"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.nest-disjunct"), gen12908)

		gen12919 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1036 := __args[0]
			_ = V1036
			V1037 := __args[1]
			_ = V1037
			gen12909 := Call(__e, ShenFunc(symcons), False, Nil)

			gen12910 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), gen12909)

			gen12911 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen12910)

			gen12912 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), Nil)

			gen12913 := Call(__e, ShenFunc(symcons), V1037, gen12912)

			gen12914 := Call(__e, ShenFunc(symcons), gen12911, gen12913)

			gen12915 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen12914)

			gen12916 := Call(__e, ShenFunc(symcons), gen12915, Nil)

			gen12917 := Call(__e, ShenFunc(symcons), V1036, gen12916)

			gen12918 := Call(__e, ShenFunc(symcons), MakeSymbol("Case"), gen12917)

			__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen12918)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lisp-or"), gen12919)

		gen12940 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1042 := __args[0]
			_ = V1042
			V1043 := __args[1]
			_ = V1043
			gen12938 := Call(__e, ShenFunc(symcons_2), V1043)

			var gen12939 Obj
			if True == gen12938 {
				gen12936 := Call(__e, ShenFunc(symtl), V1043)

				gen12937 := Call(__e, ShenFunc(sym_a), Nil, gen12936)

				if True == gen12937 {
					gen12939 = True
				} else {
					gen12939 = False
				}

			} else {
				gen12939 = False
			}
			if True == gen12939 {
				gen12934 := Call(__e, ShenFunc(symhd), V1043)

				gen12935 := Call(__e, ShenFunc(symlength), gen12934)

				__e.TailApply(ShenFunc(sym_1), gen12935, MakeNumber(1))

				return

			} else {
				gen12932 := Call(__e, ShenFunc(symcons_2), V1043)

				var gen12933 Obj
				if True == gen12932 {
					gen12930 := Call(__e, ShenFunc(symtl), V1043)

					gen12931 := Call(__e, ShenFunc(symcons_2), gen12930)

					if True == gen12931 {
						gen12933 = True
					} else {
						gen12933 = False
					}

				} else {
					gen12933 = False
				}
				if True == gen12933 {
					gen12924 := Call(__e, ShenFunc(symhd), V1043)

					gen12925 := Call(__e, ShenFunc(symlength), gen12924)

					gen12926 := Call(__e, ShenFunc(symtl), V1043)

					gen12927 := Call(__e, ShenFunc(symhd), gen12926)

					gen12928 := Call(__e, ShenFunc(symlength), gen12927)

					gen12929 := Call(__e, ShenFunc(sym_a), gen12925, gen12928)

					if True == gen12929 {
						gen12923 := Call(__e, ShenFunc(symtl), V1043)

						__e.TailApply(ShenFunc(symshen_4prolog_1aritycheck), V1042, gen12923)

						return

					} else {
						gen12920 := Call(__e, ShenFunc(symcons), V1042, Nil)

						gen12921 := Call(__e, ShenFunc(symshen_4app), gen12920, MakeString("\n"), MakeSymbol("shen.a"))

						gen12922 := Call(__e, ShenFunc(symcn), MakeString("arity error in prolog procedure "), gen12921)

						__e.TailApply(ShenFunc(symsimple_1error), gen12922)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.prolog-aritycheck"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-aritycheck"), gen12940)

		gen12963 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1045 := __args[0]
			_ = V1045
			gen12961 := Call(__e, ShenFunc(symcons_2), V1045)

			var gen12962 Obj
			if True == gen12961 {
				gen12958 := Call(__e, ShenFunc(symtl), V1045)

				gen12959 := Call(__e, ShenFunc(symcons_2), gen12958)

				var gen12960 Obj
				if True == gen12959 {
					gen12954 := Call(__e, ShenFunc(symtl), V1045)

					gen12955 := Call(__e, ShenFunc(symhd), gen12954)

					gen12956 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen12955)

					var gen12957 Obj
					if True == gen12956 {
						gen12950 := Call(__e, ShenFunc(symtl), V1045)

						gen12951 := Call(__e, ShenFunc(symtl), gen12950)

						gen12952 := Call(__e, ShenFunc(symcons_2), gen12951)

						var gen12953 Obj
						if True == gen12952 {
							gen12946 := Call(__e, ShenFunc(symtl), V1045)

							gen12947 := Call(__e, ShenFunc(symtl), gen12946)

							gen12948 := Call(__e, ShenFunc(symtl), gen12947)

							gen12949 := Call(__e, ShenFunc(sym_a), Nil, gen12948)

							if True == gen12949 {
								gen12953 = True
							} else {
								gen12953 = False
							}

						} else {
							gen12953 = False
						}
						if True == gen12953 {
							gen12957 = True
						} else {
							gen12957 = False
						}

					} else {
						gen12957 = False
					}
					if True == gen12957 {
						gen12960 = True
					} else {
						gen12960 = False
					}

				} else {
					gen12960 = False
				}
				if True == gen12960 {
					gen12962 = True
				} else {
					gen12962 = False
				}

			} else {
				gen12962 = False
			}
			if True == gen12962 {
				gen12941 := Call(__e, ShenFunc(symhd), V1045)

				gen12942 := Call(__e, ShenFunc(symtl), V1045)

				gen12943 := Call(__e, ShenFunc(symtl), gen12942)

				gen12944 := Call(__e, ShenFunc(symcons), gen12941, gen12943)

				gen12945 := Call(__e, ShenFunc(symshen_4linearise), gen12944)

				Linear := gen12945
				__e.TailApply(ShenFunc(symshen_4clause__form), Linear)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise-clause"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise-clause"), gen12963)

		gen12979 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1047 := __args[0]
			_ = V1047
			gen12977 := Call(__e, ShenFunc(symcons_2), V1047)

			var gen12978 Obj
			if True == gen12977 {
				gen12974 := Call(__e, ShenFunc(symtl), V1047)

				gen12975 := Call(__e, ShenFunc(symcons_2), gen12974)

				var gen12976 Obj
				if True == gen12975 {
					gen12971 := Call(__e, ShenFunc(symtl), V1047)

					gen12972 := Call(__e, ShenFunc(symtl), gen12971)

					gen12973 := Call(__e, ShenFunc(sym_a), Nil, gen12972)

					if True == gen12973 {
						gen12976 = True
					} else {
						gen12976 = False
					}

				} else {
					gen12976 = False
				}
				if True == gen12976 {
					gen12978 = True
				} else {
					gen12978 = False
				}

			} else {
				gen12978 = False
			}
			if True == gen12978 {
				gen12964 := Call(__e, ShenFunc(symhd), V1047)

				gen12965 := Call(__e, ShenFunc(symshen_4explicit__modes), gen12964)

				gen12966 := Call(__e, ShenFunc(symtl), V1047)

				gen12967 := Call(__e, ShenFunc(symhd), gen12966)

				gen12968 := Call(__e, ShenFunc(symshen_4cf__help), gen12967)

				gen12969 := Call(__e, ShenFunc(symcons), gen12968, Nil)

				gen12970 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen12969)

				__e.TailApply(ShenFunc(symcons), gen12965, gen12970)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.clause_form"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clause_form"), gen12979)

		gen12985 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1049 := __args[0]
			_ = V1049
			gen12984 := Call(__e, ShenFunc(symcons_2), V1049)

			if True == gen12984 {
				gen12980 := Call(__e, ShenFunc(symhd), V1049)

				gen12981 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4em__help), X)

					return
				}, 1)
				gen12982 := Call(__e, ShenFunc(symtl), V1049)

				gen12983 := Call(__e, ShenFunc(symmap), gen12981, gen12982)

				__e.TailApply(ShenFunc(symcons), gen12980, gen12983)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.explicit_modes"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.explicit_modes"), gen12985)

		gen13004 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1051 := __args[0]
			_ = V1051
			gen13002 := Call(__e, ShenFunc(symcons_2), V1051)

			var gen13003 Obj
			if True == gen13002 {
				gen12999 := Call(__e, ShenFunc(symhd), V1051)

				gen13000 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen12999)

				var gen13001 Obj
				if True == gen13000 {
					gen12996 := Call(__e, ShenFunc(symtl), V1051)

					gen12997 := Call(__e, ShenFunc(symcons_2), gen12996)

					var gen12998 Obj
					if True == gen12997 {
						gen12992 := Call(__e, ShenFunc(symtl), V1051)

						gen12993 := Call(__e, ShenFunc(symtl), gen12992)

						gen12994 := Call(__e, ShenFunc(symcons_2), gen12993)

						var gen12995 Obj
						if True == gen12994 {
							gen12988 := Call(__e, ShenFunc(symtl), V1051)

							gen12989 := Call(__e, ShenFunc(symtl), gen12988)

							gen12990 := Call(__e, ShenFunc(symtl), gen12989)

							gen12991 := Call(__e, ShenFunc(sym_a), Nil, gen12990)

							if True == gen12991 {
								gen12995 = True
							} else {
								gen12995 = False
							}

						} else {
							gen12995 = False
						}
						if True == gen12995 {
							gen12998 = True
						} else {
							gen12998 = False
						}

					} else {
						gen12998 = False
					}
					if True == gen12998 {
						gen13001 = True
					} else {
						gen13001 = False
					}

				} else {
					gen13001 = False
				}
				if True == gen13001 {
					gen13003 = True
				} else {
					gen13003 = False
				}

			} else {
				gen13003 = False
			}
			if True == gen13003 {
				__e.Return(V1051)
				return
			} else {
				gen12986 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), Nil)

				gen12987 := Call(__e, ShenFunc(symcons), V1051, gen12986)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("mode"), gen12987)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.em_help"), gen13004)

		gen13058 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1053 := __args[0]
			_ = V1053
			gen13056 := Call(__e, ShenFunc(symcons_2), V1053)

			var gen13057 Obj
			if True == gen13056 {
				gen13053 := Call(__e, ShenFunc(symhd), V1053)

				gen13054 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen13053)

				var gen13055 Obj
				if True == gen13054 {
					gen13050 := Call(__e, ShenFunc(symtl), V1053)

					gen13051 := Call(__e, ShenFunc(symcons_2), gen13050)

					var gen13052 Obj
					if True == gen13051 {
						gen13046 := Call(__e, ShenFunc(symtl), V1053)

						gen13047 := Call(__e, ShenFunc(symhd), gen13046)

						gen13048 := Call(__e, ShenFunc(symcons_2), gen13047)

						var gen13049 Obj
						if True == gen13048 {
							gen13041 := Call(__e, ShenFunc(symtl), V1053)

							gen13042 := Call(__e, ShenFunc(symhd), gen13041)

							gen13043 := Call(__e, ShenFunc(symhd), gen13042)

							gen13044 := Call(__e, ShenFunc(sym_a), MakeSymbol("="), gen13043)

							var gen13045 Obj
							if True == gen13044 {
								gen13036 := Call(__e, ShenFunc(symtl), V1053)

								gen13037 := Call(__e, ShenFunc(symhd), gen13036)

								gen13038 := Call(__e, ShenFunc(symtl), gen13037)

								gen13039 := Call(__e, ShenFunc(symcons_2), gen13038)

								var gen13040 Obj
								if True == gen13039 {
									gen13030 := Call(__e, ShenFunc(symtl), V1053)

									gen13031 := Call(__e, ShenFunc(symhd), gen13030)

									gen13032 := Call(__e, ShenFunc(symtl), gen13031)

									gen13033 := Call(__e, ShenFunc(symtl), gen13032)

									gen13034 := Call(__e, ShenFunc(symcons_2), gen13033)

									var gen13035 Obj
									if True == gen13034 {
										gen13023 := Call(__e, ShenFunc(symtl), V1053)

										gen13024 := Call(__e, ShenFunc(symhd), gen13023)

										gen13025 := Call(__e, ShenFunc(symtl), gen13024)

										gen13026 := Call(__e, ShenFunc(symtl), gen13025)

										gen13027 := Call(__e, ShenFunc(symtl), gen13026)

										gen13028 := Call(__e, ShenFunc(sym_a), Nil, gen13027)

										var gen13029 Obj
										if True == gen13028 {
											gen13019 := Call(__e, ShenFunc(symtl), V1053)

											gen13020 := Call(__e, ShenFunc(symtl), gen13019)

											gen13021 := Call(__e, ShenFunc(symcons_2), gen13020)

											var gen13022 Obj
											if True == gen13021 {
												gen13015 := Call(__e, ShenFunc(symtl), V1053)

												gen13016 := Call(__e, ShenFunc(symtl), gen13015)

												gen13017 := Call(__e, ShenFunc(symtl), gen13016)

												gen13018 := Call(__e, ShenFunc(sym_a), Nil, gen13017)

												if True == gen13018 {
													gen13022 = True
												} else {
													gen13022 = False
												}

											} else {
												gen13022 = False
											}
											if True == gen13022 {
												gen13029 = True
											} else {
												gen13029 = False
											}

										} else {
											gen13029 = False
										}
										if True == gen13029 {
											gen13035 = True
										} else {
											gen13035 = False
										}

									} else {
										gen13035 = False
									}
									if True == gen13035 {
										gen13040 = True
									} else {
										gen13040 = False
									}

								} else {
									gen13040 = False
								}
								if True == gen13040 {
									gen13045 = True
								} else {
									gen13045 = False
								}

							} else {
								gen13045 = False
							}
							if True == gen13045 {
								gen13049 = True
							} else {
								gen13049 = False
							}

						} else {
							gen13049 = False
						}
						if True == gen13049 {
							gen13052 = True
						} else {
							gen13052 = False
						}

					} else {
						gen13052 = False
					}
					if True == gen13052 {
						gen13055 = True
					} else {
						gen13055 = False
					}

				} else {
					gen13055 = False
				}
				if True == gen13055 {
					gen13057 = True
				} else {
					gen13057 = False
				}

			} else {
				gen13057 = False
			}
			if True == gen13057 {
				gen13005 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*occurs*"))

				var gen13006 Obj
				if True == gen13005 {
					gen13006 = MakeSymbol("unify!")
				} else {
					gen13006 = MakeSymbol("unify")
				}
				gen13007 := Call(__e, ShenFunc(symtl), V1053)

				gen13008 := Call(__e, ShenFunc(symhd), gen13007)

				gen13009 := Call(__e, ShenFunc(symtl), gen13008)

				gen13010 := Call(__e, ShenFunc(symcons), gen13006, gen13009)

				gen13011 := Call(__e, ShenFunc(symtl), V1053)

				gen13012 := Call(__e, ShenFunc(symtl), gen13011)

				gen13013 := Call(__e, ShenFunc(symhd), gen13012)

				gen13014 := Call(__e, ShenFunc(symshen_4cf__help), gen13013)

				__e.TailApply(ShenFunc(symcons), gen13010, gen13014)

				return

			} else {
				__e.Return(V1053)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cf_help"), gen13058)

		gen13061 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1059 := __args[0]
			_ = V1059
			gen13060 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1059)

			if True == gen13060 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*occurs*"), True)

				return
			} else {
				gen13059 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1059)

				if True == gen13059 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*occurs*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("occurs-check expects + or -\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("occurs-check"), gen13061)

		gen13094 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1062 := __args[0]
			_ = V1062
			V1063 := __args[1]
			_ = V1063
			gen13092 := Call(__e, ShenFunc(symcons_2), V1062)

			var gen13093 Obj
			if True == gen13092 {
				gen13089 := Call(__e, ShenFunc(symhd), V1062)

				gen13090 := Call(__e, ShenFunc(symcons_2), gen13089)

				var gen13091 Obj
				if True == gen13090 {
					gen13086 := Call(__e, ShenFunc(symtl), V1062)

					gen13087 := Call(__e, ShenFunc(symcons_2), gen13086)

					var gen13088 Obj
					if True == gen13087 {
						gen13082 := Call(__e, ShenFunc(symtl), V1062)

						gen13083 := Call(__e, ShenFunc(symhd), gen13082)

						gen13084 := Call(__e, ShenFunc(sym_a), MakeSymbol(":-"), gen13083)

						var gen13085 Obj
						if True == gen13084 {
							gen13078 := Call(__e, ShenFunc(symtl), V1062)

							gen13079 := Call(__e, ShenFunc(symtl), gen13078)

							gen13080 := Call(__e, ShenFunc(symcons_2), gen13079)

							var gen13081 Obj
							if True == gen13080 {
								gen13074 := Call(__e, ShenFunc(symtl), V1062)

								gen13075 := Call(__e, ShenFunc(symtl), gen13074)

								gen13076 := Call(__e, ShenFunc(symtl), gen13075)

								gen13077 := Call(__e, ShenFunc(sym_a), Nil, gen13076)

								if True == gen13077 {
									gen13081 = True
								} else {
									gen13081 = False
								}

							} else {
								gen13081 = False
							}
							if True == gen13081 {
								gen13085 = True
							} else {
								gen13085 = False
							}

						} else {
							gen13085 = False
						}
						if True == gen13085 {
							gen13088 = True
						} else {
							gen13088 = False
						}

					} else {
						gen13088 = False
					}
					if True == gen13088 {
						gen13091 = True
					} else {
						gen13091 = False
					}

				} else {
					gen13091 = False
				}
				if True == gen13091 {
					gen13093 = True
				} else {
					gen13093 = False
				}

			} else {
				gen13093 = False
			}
			if True == gen13093 {
				gen13062 := Call(__e, ShenFunc(symhd), V1062)

				gen13063 := Call(__e, ShenFunc(symtl), gen13062)

				gen13064 := Call(__e, ShenFunc(symhd), V1062)

				gen13065 := Call(__e, ShenFunc(symtl), gen13064)

				gen13066 := Call(__e, ShenFunc(symtl), V1062)

				gen13067 := Call(__e, ShenFunc(symtl), gen13066)

				gen13068 := Call(__e, ShenFunc(symhd), gen13067)

				gen13069 := Call(__e, ShenFunc(symshen_4continuation__call), gen13065, gen13068)

				gen13070 := Call(__e, ShenFunc(symcons), gen13069, Nil)

				gen13071 := Call(__e, ShenFunc(symcons), gen13063, gen13070)

				gen13072 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13071)

				gen13073 := Call(__e, ShenFunc(symshen_4make__mu__application), gen13072, V1063)

				MuApplication := gen13073
				__e.TailApply(ShenFunc(symshen_4mu__reduction), MuApplication, MakeSymbol("+"))

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.aum"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aum"), gen13094)

		gen13100 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1066 := __args[0]
			_ = V1066
			V1067 := __args[1]
			_ = V1067
			gen13095 := Call(__e, ShenFunc(symshen_4extract__vars), V1066)

			gen13096 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), gen13095)

			VTerms := gen13096
			gen13097 := Call(__e, ShenFunc(symshen_4extract__vars), V1067)

			VBody := gen13097
			gen13098 := Call(__e, ShenFunc(symdifference), VBody, VTerms)

			gen13099 := Call(__e, ShenFunc(symremove), MakeSymbol("Throwcontrol"), gen13098)

			Free := gen13099
			__e.TailApply(ShenFunc(symshen_4cc__help), Free, V1067)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.continuation_call"), gen13100)

		gen13101 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1070 := __args[0]
			_ = V1070
			V1071 := __args[1]
			_ = V1071
			__e.TailApply(ShenFunc(symshen_4remove_1h), V1070, V1071, Nil)

			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("remove"), gen13101)

		gen13113 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1078 := __args[0]
			_ = V1078
			V1079 := __args[1]
			_ = V1079
			V1080 := __args[2]
			_ = V1080
			gen13112 := Call(__e, ShenFunc(sym_a), Nil, V1079)

			if True == gen13112 {
				__e.TailApply(ShenFunc(symreverse), V1080)

				return
			} else {
				gen13110 := Call(__e, ShenFunc(symcons_2), V1079)

				var gen13111 Obj
				if True == gen13110 {
					gen13108 := Call(__e, ShenFunc(symhd), V1079)

					gen13109 := Call(__e, ShenFunc(sym_a), gen13108, V1078)

					if True == gen13109 {
						gen13111 = True
					} else {
						gen13111 = False
					}

				} else {
					gen13111 = False
				}
				if True == gen13111 {
					gen13106 := Call(__e, ShenFunc(symhd), V1079)

					gen13107 := Call(__e, ShenFunc(symtl), V1079)

					__e.TailApply(ShenFunc(symshen_4remove_1h), gen13106, gen13107, V1080)

					return

				} else {
					gen13105 := Call(__e, ShenFunc(symcons_2), V1079)

					if True == gen13105 {
						gen13102 := Call(__e, ShenFunc(symtl), V1079)

						gen13103 := Call(__e, ShenFunc(symhd), V1079)

						gen13104 := Call(__e, ShenFunc(symcons), gen13103, V1080)

						__e.TailApply(ShenFunc(symshen_4remove_1h), V1078, gen13102, gen13104)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.remove-h"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove-h"), gen13113)

		gen13145 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1083 := __args[0]
			_ = V1083
			V1084 := __args[1]
			_ = V1084
			gen13143 := Call(__e, ShenFunc(sym_a), Nil, V1083)

			var gen13144 Obj
			if True == gen13143 {
				gen13142 := Call(__e, ShenFunc(sym_a), Nil, V1084)

				if True == gen13142 {
					gen13144 = True
				} else {
					gen13144 = False
				}

			} else {
				gen13144 = False
			}
			if True == gen13144 {
				gen13140 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.stack"), Nil)

				gen13141 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13140)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pop"), gen13141)

				return

			} else {
				gen13139 := Call(__e, ShenFunc(sym_a), Nil, V1084)

				if True == gen13139 {
					gen13129 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.stack"), Nil)

					gen13130 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13129)

					gen13131 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pop"), gen13130)

					gen13132 := Call(__e, ShenFunc(symcons), gen13131, Nil)

					gen13133 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13132)

					gen13134 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen13133)

					gen13135 := Call(__e, ShenFunc(symcons), V1083, gen13134)

					gen13136 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13135)

					gen13137 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen13136)

					gen13138 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13137)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.rename"), gen13138)

					return

				} else {
					gen13128 := Call(__e, ShenFunc(sym_a), Nil, V1083)

					if True == gen13128 {
						gen13125 := Call(__e, ShenFunc(symcons), V1084, Nil)

						gen13126 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.continuation"), gen13125)

						gen13127 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13126)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("call"), gen13127)

						return

					} else {
						gen13114 := Call(__e, ShenFunc(symcons), V1084, Nil)

						gen13115 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.continuation"), gen13114)

						gen13116 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13115)

						gen13117 := Call(__e, ShenFunc(symcons), MakeSymbol("call"), gen13116)

						gen13118 := Call(__e, ShenFunc(symcons), gen13117, Nil)

						gen13119 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13118)

						gen13120 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen13119)

						gen13121 := Call(__e, ShenFunc(symcons), V1083, gen13120)

						gen13122 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13121)

						gen13123 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen13122)

						gen13124 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13123)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.rename"), gen13124)

						return

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cc_help"), gen13145)

		gen13209 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1087 := __args[0]
			_ = V1087
			V1088 := __args[1]
			_ = V1088
			gen13207 := Call(__e, ShenFunc(symcons_2), V1087)

			var gen13208 Obj
			if True == gen13207 {
				gen13204 := Call(__e, ShenFunc(symhd), V1087)

				gen13205 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13204)

				var gen13206 Obj
				if True == gen13205 {
					gen13201 := Call(__e, ShenFunc(symtl), V1087)

					gen13202 := Call(__e, ShenFunc(symcons_2), gen13201)

					var gen13203 Obj
					if True == gen13202 {
						gen13197 := Call(__e, ShenFunc(symtl), V1087)

						gen13198 := Call(__e, ShenFunc(symhd), gen13197)

						gen13199 := Call(__e, ShenFunc(sym_a), Nil, gen13198)

						var gen13200 Obj
						if True == gen13199 {
							gen13193 := Call(__e, ShenFunc(symtl), V1087)

							gen13194 := Call(__e, ShenFunc(symtl), gen13193)

							gen13195 := Call(__e, ShenFunc(symcons_2), gen13194)

							var gen13196 Obj
							if True == gen13195 {
								gen13188 := Call(__e, ShenFunc(symtl), V1087)

								gen13189 := Call(__e, ShenFunc(symtl), gen13188)

								gen13190 := Call(__e, ShenFunc(symtl), gen13189)

								gen13191 := Call(__e, ShenFunc(sym_a), Nil, gen13190)

								var gen13192 Obj
								if True == gen13191 {
									gen13187 := Call(__e, ShenFunc(sym_a), Nil, V1088)

									if True == gen13187 {
										gen13192 = True
									} else {
										gen13192 = False
									}

								} else {
									gen13192 = False
								}
								if True == gen13192 {
									gen13196 = True
								} else {
									gen13196 = False
								}

							} else {
								gen13196 = False
							}
							if True == gen13196 {
								gen13200 = True
							} else {
								gen13200 = False
							}

						} else {
							gen13200 = False
						}
						if True == gen13200 {
							gen13203 = True
						} else {
							gen13203 = False
						}

					} else {
						gen13203 = False
					}
					if True == gen13203 {
						gen13206 = True
					} else {
						gen13206 = False
					}

				} else {
					gen13206 = False
				}
				if True == gen13206 {
					gen13208 = True
				} else {
					gen13208 = False
				}

			} else {
				gen13208 = False
			}
			if True == gen13208 {
				gen13185 := Call(__e, ShenFunc(symtl), V1087)

				gen13186 := Call(__e, ShenFunc(symtl), gen13185)

				__e.TailApply(ShenFunc(symhd), gen13186)

				return

			} else {
				gen13183 := Call(__e, ShenFunc(symcons_2), V1087)

				var gen13184 Obj
				if True == gen13183 {
					gen13180 := Call(__e, ShenFunc(symhd), V1087)

					gen13181 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13180)

					var gen13182 Obj
					if True == gen13181 {
						gen13177 := Call(__e, ShenFunc(symtl), V1087)

						gen13178 := Call(__e, ShenFunc(symcons_2), gen13177)

						var gen13179 Obj
						if True == gen13178 {
							gen13173 := Call(__e, ShenFunc(symtl), V1087)

							gen13174 := Call(__e, ShenFunc(symhd), gen13173)

							gen13175 := Call(__e, ShenFunc(symcons_2), gen13174)

							var gen13176 Obj
							if True == gen13175 {
								gen13169 := Call(__e, ShenFunc(symtl), V1087)

								gen13170 := Call(__e, ShenFunc(symtl), gen13169)

								gen13171 := Call(__e, ShenFunc(symcons_2), gen13170)

								var gen13172 Obj
								if True == gen13171 {
									gen13164 := Call(__e, ShenFunc(symtl), V1087)

									gen13165 := Call(__e, ShenFunc(symtl), gen13164)

									gen13166 := Call(__e, ShenFunc(symtl), gen13165)

									gen13167 := Call(__e, ShenFunc(sym_a), Nil, gen13166)

									var gen13168 Obj
									if True == gen13167 {
										gen13163 := Call(__e, ShenFunc(symcons_2), V1088)

										if True == gen13163 {
											gen13168 = True
										} else {
											gen13168 = False
										}

									} else {
										gen13168 = False
									}
									if True == gen13168 {
										gen13172 = True
									} else {
										gen13172 = False
									}

								} else {
									gen13172 = False
								}
								if True == gen13172 {
									gen13176 = True
								} else {
									gen13176 = False
								}

							} else {
								gen13176 = False
							}
							if True == gen13176 {
								gen13179 = True
							} else {
								gen13179 = False
							}

						} else {
							gen13179 = False
						}
						if True == gen13179 {
							gen13182 = True
						} else {
							gen13182 = False
						}

					} else {
						gen13182 = False
					}
					if True == gen13182 {
						gen13184 = True
					} else {
						gen13184 = False
					}

				} else {
					gen13184 = False
				}
				if True == gen13184 {
					gen13146 := Call(__e, ShenFunc(symtl), V1087)

					gen13147 := Call(__e, ShenFunc(symhd), gen13146)

					gen13148 := Call(__e, ShenFunc(symhd), gen13147)

					gen13149 := Call(__e, ShenFunc(symtl), V1087)

					gen13150 := Call(__e, ShenFunc(symhd), gen13149)

					gen13151 := Call(__e, ShenFunc(symtl), gen13150)

					gen13152 := Call(__e, ShenFunc(symtl), V1087)

					gen13153 := Call(__e, ShenFunc(symtl), gen13152)

					gen13154 := Call(__e, ShenFunc(symcons), gen13151, gen13153)

					gen13155 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13154)

					gen13156 := Call(__e, ShenFunc(symtl), V1088)

					gen13157 := Call(__e, ShenFunc(symshen_4make__mu__application), gen13155, gen13156)

					gen13158 := Call(__e, ShenFunc(symcons), gen13157, Nil)

					gen13159 := Call(__e, ShenFunc(symcons), gen13148, gen13158)

					gen13160 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13159)

					gen13161 := Call(__e, ShenFunc(symhd), V1088)

					gen13162 := Call(__e, ShenFunc(symcons), gen13161, Nil)

					__e.TailApply(ShenFunc(symcons), gen13160, gen13162)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.make_mu_application"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.make_mu_application"), gen13209)

		gen13798 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1097 := __args[0]
			_ = V1097
			V1098 := __args[1]
			_ = V1098
			gen13796 := Call(__e, ShenFunc(symcons_2), V1097)

			var gen13797 Obj
			if True == gen13796 {
				gen13793 := Call(__e, ShenFunc(symhd), V1097)

				gen13794 := Call(__e, ShenFunc(symcons_2), gen13793)

				var gen13795 Obj
				if True == gen13794 {
					gen13789 := Call(__e, ShenFunc(symhd), V1097)

					gen13790 := Call(__e, ShenFunc(symhd), gen13789)

					gen13791 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13790)

					var gen13792 Obj
					if True == gen13791 {
						gen13785 := Call(__e, ShenFunc(symhd), V1097)

						gen13786 := Call(__e, ShenFunc(symtl), gen13785)

						gen13787 := Call(__e, ShenFunc(symcons_2), gen13786)

						var gen13788 Obj
						if True == gen13787 {
							gen13780 := Call(__e, ShenFunc(symhd), V1097)

							gen13781 := Call(__e, ShenFunc(symtl), gen13780)

							gen13782 := Call(__e, ShenFunc(symhd), gen13781)

							gen13783 := Call(__e, ShenFunc(symcons_2), gen13782)

							var gen13784 Obj
							if True == gen13783 {
								gen13774 := Call(__e, ShenFunc(symhd), V1097)

								gen13775 := Call(__e, ShenFunc(symtl), gen13774)

								gen13776 := Call(__e, ShenFunc(symhd), gen13775)

								gen13777 := Call(__e, ShenFunc(symhd), gen13776)

								gen13778 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen13777)

								var gen13779 Obj
								if True == gen13778 {
									gen13768 := Call(__e, ShenFunc(symhd), V1097)

									gen13769 := Call(__e, ShenFunc(symtl), gen13768)

									gen13770 := Call(__e, ShenFunc(symhd), gen13769)

									gen13771 := Call(__e, ShenFunc(symtl), gen13770)

									gen13772 := Call(__e, ShenFunc(symcons_2), gen13771)

									var gen13773 Obj
									if True == gen13772 {
										gen13761 := Call(__e, ShenFunc(symhd), V1097)

										gen13762 := Call(__e, ShenFunc(symtl), gen13761)

										gen13763 := Call(__e, ShenFunc(symhd), gen13762)

										gen13764 := Call(__e, ShenFunc(symtl), gen13763)

										gen13765 := Call(__e, ShenFunc(symtl), gen13764)

										gen13766 := Call(__e, ShenFunc(symcons_2), gen13765)

										var gen13767 Obj
										if True == gen13766 {
											gen13753 := Call(__e, ShenFunc(symhd), V1097)

											gen13754 := Call(__e, ShenFunc(symtl), gen13753)

											gen13755 := Call(__e, ShenFunc(symhd), gen13754)

											gen13756 := Call(__e, ShenFunc(symtl), gen13755)

											gen13757 := Call(__e, ShenFunc(symtl), gen13756)

											gen13758 := Call(__e, ShenFunc(symtl), gen13757)

											gen13759 := Call(__e, ShenFunc(sym_a), Nil, gen13758)

											var gen13760 Obj
											if True == gen13759 {
												gen13748 := Call(__e, ShenFunc(symhd), V1097)

												gen13749 := Call(__e, ShenFunc(symtl), gen13748)

												gen13750 := Call(__e, ShenFunc(symtl), gen13749)

												gen13751 := Call(__e, ShenFunc(symcons_2), gen13750)

												var gen13752 Obj
												if True == gen13751 {
													gen13742 := Call(__e, ShenFunc(symhd), V1097)

													gen13743 := Call(__e, ShenFunc(symtl), gen13742)

													gen13744 := Call(__e, ShenFunc(symtl), gen13743)

													gen13745 := Call(__e, ShenFunc(symtl), gen13744)

													gen13746 := Call(__e, ShenFunc(sym_a), Nil, gen13745)

													var gen13747 Obj
													if True == gen13746 {
														gen13739 := Call(__e, ShenFunc(symtl), V1097)

														gen13740 := Call(__e, ShenFunc(symcons_2), gen13739)

														var gen13741 Obj
														if True == gen13740 {
															gen13736 := Call(__e, ShenFunc(symtl), V1097)

															gen13737 := Call(__e, ShenFunc(symtl), gen13736)

															gen13738 := Call(__e, ShenFunc(sym_a), Nil, gen13737)

															if True == gen13738 {
																gen13741 = True
															} else {
																gen13741 = False
															}

														} else {
															gen13741 = False
														}
														if True == gen13741 {
															gen13747 = True
														} else {
															gen13747 = False
														}

													} else {
														gen13747 = False
													}
													if True == gen13747 {
														gen13752 = True
													} else {
														gen13752 = False
													}

												} else {
													gen13752 = False
												}
												if True == gen13752 {
													gen13760 = True
												} else {
													gen13760 = False
												}

											} else {
												gen13760 = False
											}
											if True == gen13760 {
												gen13767 = True
											} else {
												gen13767 = False
											}

										} else {
											gen13767 = False
										}
										if True == gen13767 {
											gen13773 = True
										} else {
											gen13773 = False
										}

									} else {
										gen13773 = False
									}
									if True == gen13773 {
										gen13779 = True
									} else {
										gen13779 = False
									}

								} else {
									gen13779 = False
								}
								if True == gen13779 {
									gen13784 = True
								} else {
									gen13784 = False
								}

							} else {
								gen13784 = False
							}
							if True == gen13784 {
								gen13788 = True
							} else {
								gen13788 = False
							}

						} else {
							gen13788 = False
						}
						if True == gen13788 {
							gen13792 = True
						} else {
							gen13792 = False
						}

					} else {
						gen13792 = False
					}
					if True == gen13792 {
						gen13795 = True
					} else {
						gen13795 = False
					}

				} else {
					gen13795 = False
				}
				if True == gen13795 {
					gen13797 = True
				} else {
					gen13797 = False
				}

			} else {
				gen13797 = False
			}
			if True == gen13797 {
				gen13718 := Call(__e, ShenFunc(symhd), V1097)

				gen13719 := Call(__e, ShenFunc(symtl), gen13718)

				gen13720 := Call(__e, ShenFunc(symhd), gen13719)

				gen13721 := Call(__e, ShenFunc(symtl), gen13720)

				gen13722 := Call(__e, ShenFunc(symhd), gen13721)

				gen13723 := Call(__e, ShenFunc(symhd), V1097)

				gen13724 := Call(__e, ShenFunc(symtl), gen13723)

				gen13725 := Call(__e, ShenFunc(symtl), gen13724)

				gen13726 := Call(__e, ShenFunc(symcons), gen13722, gen13725)

				gen13727 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13726)

				gen13728 := Call(__e, ShenFunc(symtl), V1097)

				gen13729 := Call(__e, ShenFunc(symcons), gen13727, gen13728)

				gen13730 := Call(__e, ShenFunc(symhd), V1097)

				gen13731 := Call(__e, ShenFunc(symtl), gen13730)

				gen13732 := Call(__e, ShenFunc(symhd), gen13731)

				gen13733 := Call(__e, ShenFunc(symtl), gen13732)

				gen13734 := Call(__e, ShenFunc(symtl), gen13733)

				gen13735 := Call(__e, ShenFunc(symhd), gen13734)

				__e.TailApply(ShenFunc(symshen_4mu__reduction), gen13729, gen13735)

				return

			} else {
				gen13716 := Call(__e, ShenFunc(symcons_2), V1097)

				var gen13717 Obj
				if True == gen13716 {
					gen13713 := Call(__e, ShenFunc(symhd), V1097)

					gen13714 := Call(__e, ShenFunc(symcons_2), gen13713)

					var gen13715 Obj
					if True == gen13714 {
						gen13709 := Call(__e, ShenFunc(symhd), V1097)

						gen13710 := Call(__e, ShenFunc(symhd), gen13709)

						gen13711 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13710)

						var gen13712 Obj
						if True == gen13711 {
							gen13705 := Call(__e, ShenFunc(symhd), V1097)

							gen13706 := Call(__e, ShenFunc(symtl), gen13705)

							gen13707 := Call(__e, ShenFunc(symcons_2), gen13706)

							var gen13708 Obj
							if True == gen13707 {
								gen13700 := Call(__e, ShenFunc(symhd), V1097)

								gen13701 := Call(__e, ShenFunc(symtl), gen13700)

								gen13702 := Call(__e, ShenFunc(symtl), gen13701)

								gen13703 := Call(__e, ShenFunc(symcons_2), gen13702)

								var gen13704 Obj
								if True == gen13703 {
									gen13694 := Call(__e, ShenFunc(symhd), V1097)

									gen13695 := Call(__e, ShenFunc(symtl), gen13694)

									gen13696 := Call(__e, ShenFunc(symtl), gen13695)

									gen13697 := Call(__e, ShenFunc(symtl), gen13696)

									gen13698 := Call(__e, ShenFunc(sym_a), Nil, gen13697)

									var gen13699 Obj
									if True == gen13698 {
										gen13691 := Call(__e, ShenFunc(symtl), V1097)

										gen13692 := Call(__e, ShenFunc(symcons_2), gen13691)

										var gen13693 Obj
										if True == gen13692 {
											gen13687 := Call(__e, ShenFunc(symtl), V1097)

											gen13688 := Call(__e, ShenFunc(symtl), gen13687)

											gen13689 := Call(__e, ShenFunc(sym_a), Nil, gen13688)

											var gen13690 Obj
											if True == gen13689 {
												gen13683 := Call(__e, ShenFunc(symhd), V1097)

												gen13684 := Call(__e, ShenFunc(symtl), gen13683)

												gen13685 := Call(__e, ShenFunc(symhd), gen13684)

												gen13686 := Call(__e, ShenFunc(sym_a), MakeSymbol("_"), gen13685)

												if True == gen13686 {
													gen13690 = True
												} else {
													gen13690 = False
												}

											} else {
												gen13690 = False
											}
											if True == gen13690 {
												gen13693 = True
											} else {
												gen13693 = False
											}

										} else {
											gen13693 = False
										}
										if True == gen13693 {
											gen13699 = True
										} else {
											gen13699 = False
										}

									} else {
										gen13699 = False
									}
									if True == gen13699 {
										gen13704 = True
									} else {
										gen13704 = False
									}

								} else {
									gen13704 = False
								}
								if True == gen13704 {
									gen13708 = True
								} else {
									gen13708 = False
								}

							} else {
								gen13708 = False
							}
							if True == gen13708 {
								gen13712 = True
							} else {
								gen13712 = False
							}

						} else {
							gen13712 = False
						}
						if True == gen13712 {
							gen13715 = True
						} else {
							gen13715 = False
						}

					} else {
						gen13715 = False
					}
					if True == gen13715 {
						gen13717 = True
					} else {
						gen13717 = False
					}

				} else {
					gen13717 = False
				}
				if True == gen13717 {
					gen13679 := Call(__e, ShenFunc(symhd), V1097)

					gen13680 := Call(__e, ShenFunc(symtl), gen13679)

					gen13681 := Call(__e, ShenFunc(symtl), gen13680)

					gen13682 := Call(__e, ShenFunc(symhd), gen13681)

					__e.TailApply(ShenFunc(symshen_4mu__reduction), gen13682, V1098)

					return

				} else {
					gen13677 := Call(__e, ShenFunc(symcons_2), V1097)

					var gen13678 Obj
					if True == gen13677 {
						gen13674 := Call(__e, ShenFunc(symhd), V1097)

						gen13675 := Call(__e, ShenFunc(symcons_2), gen13674)

						var gen13676 Obj
						if True == gen13675 {
							gen13670 := Call(__e, ShenFunc(symhd), V1097)

							gen13671 := Call(__e, ShenFunc(symhd), gen13670)

							gen13672 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13671)

							var gen13673 Obj
							if True == gen13672 {
								gen13666 := Call(__e, ShenFunc(symhd), V1097)

								gen13667 := Call(__e, ShenFunc(symtl), gen13666)

								gen13668 := Call(__e, ShenFunc(symcons_2), gen13667)

								var gen13669 Obj
								if True == gen13668 {
									gen13661 := Call(__e, ShenFunc(symhd), V1097)

									gen13662 := Call(__e, ShenFunc(symtl), gen13661)

									gen13663 := Call(__e, ShenFunc(symtl), gen13662)

									gen13664 := Call(__e, ShenFunc(symcons_2), gen13663)

									var gen13665 Obj
									if True == gen13664 {
										gen13655 := Call(__e, ShenFunc(symhd), V1097)

										gen13656 := Call(__e, ShenFunc(symtl), gen13655)

										gen13657 := Call(__e, ShenFunc(symtl), gen13656)

										gen13658 := Call(__e, ShenFunc(symtl), gen13657)

										gen13659 := Call(__e, ShenFunc(sym_a), Nil, gen13658)

										var gen13660 Obj
										if True == gen13659 {
											gen13652 := Call(__e, ShenFunc(symtl), V1097)

											gen13653 := Call(__e, ShenFunc(symcons_2), gen13652)

											var gen13654 Obj
											if True == gen13653 {
												gen13648 := Call(__e, ShenFunc(symtl), V1097)

												gen13649 := Call(__e, ShenFunc(symtl), gen13648)

												gen13650 := Call(__e, ShenFunc(sym_a), Nil, gen13649)

												var gen13651 Obj
												if True == gen13650 {
													gen13642 := Call(__e, ShenFunc(symhd), V1097)

													gen13643 := Call(__e, ShenFunc(symtl), gen13642)

													gen13644 := Call(__e, ShenFunc(symhd), gen13643)

													gen13645 := Call(__e, ShenFunc(symtl), V1097)

													gen13646 := Call(__e, ShenFunc(symhd), gen13645)

													gen13647 := Call(__e, ShenFunc(symshen_4ephemeral__variable_2), gen13644, gen13646)

													if True == gen13647 {
														gen13651 = True
													} else {
														gen13651 = False
													}

												} else {
													gen13651 = False
												}
												if True == gen13651 {
													gen13654 = True
												} else {
													gen13654 = False
												}

											} else {
												gen13654 = False
											}
											if True == gen13654 {
												gen13660 = True
											} else {
												gen13660 = False
											}

										} else {
											gen13660 = False
										}
										if True == gen13660 {
											gen13665 = True
										} else {
											gen13665 = False
										}

									} else {
										gen13665 = False
									}
									if True == gen13665 {
										gen13669 = True
									} else {
										gen13669 = False
									}

								} else {
									gen13669 = False
								}
								if True == gen13669 {
									gen13673 = True
								} else {
									gen13673 = False
								}

							} else {
								gen13673 = False
							}
							if True == gen13673 {
								gen13676 = True
							} else {
								gen13676 = False
							}

						} else {
							gen13676 = False
						}
						if True == gen13676 {
							gen13678 = True
						} else {
							gen13678 = False
						}

					} else {
						gen13678 = False
					}
					if True == gen13678 {
						gen13632 := Call(__e, ShenFunc(symtl), V1097)

						gen13633 := Call(__e, ShenFunc(symhd), gen13632)

						gen13634 := Call(__e, ShenFunc(symhd), V1097)

						gen13635 := Call(__e, ShenFunc(symtl), gen13634)

						gen13636 := Call(__e, ShenFunc(symhd), gen13635)

						gen13637 := Call(__e, ShenFunc(symhd), V1097)

						gen13638 := Call(__e, ShenFunc(symtl), gen13637)

						gen13639 := Call(__e, ShenFunc(symtl), gen13638)

						gen13640 := Call(__e, ShenFunc(symhd), gen13639)

						gen13641 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13640, V1098)

						__e.TailApply(ShenFunc(symsubst), gen13633, gen13636, gen13641)

						return

					} else {
						gen13630 := Call(__e, ShenFunc(symcons_2), V1097)

						var gen13631 Obj
						if True == gen13630 {
							gen13627 := Call(__e, ShenFunc(symhd), V1097)

							gen13628 := Call(__e, ShenFunc(symcons_2), gen13627)

							var gen13629 Obj
							if True == gen13628 {
								gen13623 := Call(__e, ShenFunc(symhd), V1097)

								gen13624 := Call(__e, ShenFunc(symhd), gen13623)

								gen13625 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13624)

								var gen13626 Obj
								if True == gen13625 {
									gen13619 := Call(__e, ShenFunc(symhd), V1097)

									gen13620 := Call(__e, ShenFunc(symtl), gen13619)

									gen13621 := Call(__e, ShenFunc(symcons_2), gen13620)

									var gen13622 Obj
									if True == gen13621 {
										gen13614 := Call(__e, ShenFunc(symhd), V1097)

										gen13615 := Call(__e, ShenFunc(symtl), gen13614)

										gen13616 := Call(__e, ShenFunc(symtl), gen13615)

										gen13617 := Call(__e, ShenFunc(symcons_2), gen13616)

										var gen13618 Obj
										if True == gen13617 {
											gen13608 := Call(__e, ShenFunc(symhd), V1097)

											gen13609 := Call(__e, ShenFunc(symtl), gen13608)

											gen13610 := Call(__e, ShenFunc(symtl), gen13609)

											gen13611 := Call(__e, ShenFunc(symtl), gen13610)

											gen13612 := Call(__e, ShenFunc(sym_a), Nil, gen13611)

											var gen13613 Obj
											if True == gen13612 {
												gen13605 := Call(__e, ShenFunc(symtl), V1097)

												gen13606 := Call(__e, ShenFunc(symcons_2), gen13605)

												var gen13607 Obj
												if True == gen13606 {
													gen13601 := Call(__e, ShenFunc(symtl), V1097)

													gen13602 := Call(__e, ShenFunc(symtl), gen13601)

													gen13603 := Call(__e, ShenFunc(sym_a), Nil, gen13602)

													var gen13604 Obj
													if True == gen13603 {
														gen13597 := Call(__e, ShenFunc(symhd), V1097)

														gen13598 := Call(__e, ShenFunc(symtl), gen13597)

														gen13599 := Call(__e, ShenFunc(symhd), gen13598)

														gen13600 := Call(__e, ShenFunc(symvariable_2), gen13599)

														if True == gen13600 {
															gen13604 = True
														} else {
															gen13604 = False
														}

													} else {
														gen13604 = False
													}
													if True == gen13604 {
														gen13607 = True
													} else {
														gen13607 = False
													}

												} else {
													gen13607 = False
												}
												if True == gen13607 {
													gen13613 = True
												} else {
													gen13613 = False
												}

											} else {
												gen13613 = False
											}
											if True == gen13613 {
												gen13618 = True
											} else {
												gen13618 = False
											}

										} else {
											gen13618 = False
										}
										if True == gen13618 {
											gen13622 = True
										} else {
											gen13622 = False
										}

									} else {
										gen13622 = False
									}
									if True == gen13622 {
										gen13626 = True
									} else {
										gen13626 = False
									}

								} else {
									gen13626 = False
								}
								if True == gen13626 {
									gen13629 = True
								} else {
									gen13629 = False
								}

							} else {
								gen13629 = False
							}
							if True == gen13629 {
								gen13631 = True
							} else {
								gen13631 = False
							}

						} else {
							gen13631 = False
						}
						if True == gen13631 {
							gen13582 := Call(__e, ShenFunc(symhd), V1097)

							gen13583 := Call(__e, ShenFunc(symtl), gen13582)

							gen13584 := Call(__e, ShenFunc(symhd), gen13583)

							gen13585 := Call(__e, ShenFunc(symtl), V1097)

							gen13586 := Call(__e, ShenFunc(symhd), gen13585)

							gen13587 := Call(__e, ShenFunc(symhd), V1097)

							gen13588 := Call(__e, ShenFunc(symtl), gen13587)

							gen13589 := Call(__e, ShenFunc(symtl), gen13588)

							gen13590 := Call(__e, ShenFunc(symhd), gen13589)

							gen13591 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13590, V1098)

							gen13592 := Call(__e, ShenFunc(symcons), gen13591, Nil)

							gen13593 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13592)

							gen13594 := Call(__e, ShenFunc(symcons), gen13586, gen13593)

							gen13595 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen13594)

							gen13596 := Call(__e, ShenFunc(symcons), gen13584, gen13595)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen13596)

							return

						} else {
							gen13580 := Call(__e, ShenFunc(symcons_2), V1097)

							var gen13581 Obj
							if True == gen13580 {
								gen13577 := Call(__e, ShenFunc(symhd), V1097)

								gen13578 := Call(__e, ShenFunc(symcons_2), gen13577)

								var gen13579 Obj
								if True == gen13578 {
									gen13573 := Call(__e, ShenFunc(symhd), V1097)

									gen13574 := Call(__e, ShenFunc(symhd), gen13573)

									gen13575 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13574)

									var gen13576 Obj
									if True == gen13575 {
										gen13569 := Call(__e, ShenFunc(symhd), V1097)

										gen13570 := Call(__e, ShenFunc(symtl), gen13569)

										gen13571 := Call(__e, ShenFunc(symcons_2), gen13570)

										var gen13572 Obj
										if True == gen13571 {
											gen13564 := Call(__e, ShenFunc(symhd), V1097)

											gen13565 := Call(__e, ShenFunc(symtl), gen13564)

											gen13566 := Call(__e, ShenFunc(symtl), gen13565)

											gen13567 := Call(__e, ShenFunc(symcons_2), gen13566)

											var gen13568 Obj
											if True == gen13567 {
												gen13558 := Call(__e, ShenFunc(symhd), V1097)

												gen13559 := Call(__e, ShenFunc(symtl), gen13558)

												gen13560 := Call(__e, ShenFunc(symtl), gen13559)

												gen13561 := Call(__e, ShenFunc(symtl), gen13560)

												gen13562 := Call(__e, ShenFunc(sym_a), Nil, gen13561)

												var gen13563 Obj
												if True == gen13562 {
													gen13555 := Call(__e, ShenFunc(symtl), V1097)

													gen13556 := Call(__e, ShenFunc(symcons_2), gen13555)

													var gen13557 Obj
													if True == gen13556 {
														gen13551 := Call(__e, ShenFunc(symtl), V1097)

														gen13552 := Call(__e, ShenFunc(symtl), gen13551)

														gen13553 := Call(__e, ShenFunc(sym_a), Nil, gen13552)

														var gen13554 Obj
														if True == gen13553 {
															gen13549 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1098)

															var gen13550 Obj
															if True == gen13549 {
																gen13545 := Call(__e, ShenFunc(symhd), V1097)

																gen13546 := Call(__e, ShenFunc(symtl), gen13545)

																gen13547 := Call(__e, ShenFunc(symhd), gen13546)

																gen13548 := Call(__e, ShenFunc(symshen_4prolog__constant_2), gen13547)

																if True == gen13548 {
																	gen13550 = True
																} else {
																	gen13550 = False
																}

															} else {
																gen13550 = False
															}
															if True == gen13550 {
																gen13554 = True
															} else {
																gen13554 = False
															}

														} else {
															gen13554 = False
														}
														if True == gen13554 {
															gen13557 = True
														} else {
															gen13557 = False
														}

													} else {
														gen13557 = False
													}
													if True == gen13557 {
														gen13563 = True
													} else {
														gen13563 = False
													}

												} else {
													gen13563 = False
												}
												if True == gen13563 {
													gen13568 = True
												} else {
													gen13568 = False
												}

											} else {
												gen13568 = False
											}
											if True == gen13568 {
												gen13572 = True
											} else {
												gen13572 = False
											}

										} else {
											gen13572 = False
										}
										if True == gen13572 {
											gen13576 = True
										} else {
											gen13576 = False
										}

									} else {
										gen13576 = False
									}
									if True == gen13576 {
										gen13579 = True
									} else {
										gen13579 = False
									}

								} else {
									gen13579 = False
								}
								if True == gen13579 {
									gen13581 = True
								} else {
									gen13581 = False
								}

							} else {
								gen13581 = False
							}
							if True == gen13581 {
								gen13515 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

								Z := gen13515
								gen13516 := Call(__e, ShenFunc(symtl), V1097)

								gen13517 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen13516)

								gen13518 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13517)

								gen13519 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen13518)

								gen13520 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13519)

								gen13521 := Call(__e, ShenFunc(symhd), V1097)

								gen13522 := Call(__e, ShenFunc(symtl), gen13521)

								gen13523 := Call(__e, ShenFunc(symhd), gen13522)

								gen13524 := Call(__e, ShenFunc(symcons), gen13523, Nil)

								gen13525 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen13524)

								gen13526 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen13525)

								gen13527 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13526)

								gen13528 := Call(__e, ShenFunc(symcons), Z, gen13527)

								gen13529 := Call(__e, ShenFunc(symhd), V1097)

								gen13530 := Call(__e, ShenFunc(symtl), gen13529)

								gen13531 := Call(__e, ShenFunc(symtl), gen13530)

								gen13532 := Call(__e, ShenFunc(symhd), gen13531)

								gen13533 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13532, MakeSymbol("-"))

								gen13534 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

								gen13535 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13534)

								gen13536 := Call(__e, ShenFunc(symcons), gen13533, gen13535)

								gen13537 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13536)

								gen13538 := Call(__e, ShenFunc(symcons), gen13528, gen13537)

								gen13539 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13538)

								gen13540 := Call(__e, ShenFunc(symcons), gen13539, Nil)

								gen13541 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13540)

								gen13542 := Call(__e, ShenFunc(symcons), gen13520, gen13541)

								gen13543 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen13542)

								gen13544 := Call(__e, ShenFunc(symcons), Z, gen13543)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen13544)

								return

							} else {
								gen13513 := Call(__e, ShenFunc(symcons_2), V1097)

								var gen13514 Obj
								if True == gen13513 {
									gen13510 := Call(__e, ShenFunc(symhd), V1097)

									gen13511 := Call(__e, ShenFunc(symcons_2), gen13510)

									var gen13512 Obj
									if True == gen13511 {
										gen13506 := Call(__e, ShenFunc(symhd), V1097)

										gen13507 := Call(__e, ShenFunc(symhd), gen13506)

										gen13508 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13507)

										var gen13509 Obj
										if True == gen13508 {
											gen13502 := Call(__e, ShenFunc(symhd), V1097)

											gen13503 := Call(__e, ShenFunc(symtl), gen13502)

											gen13504 := Call(__e, ShenFunc(symcons_2), gen13503)

											var gen13505 Obj
											if True == gen13504 {
												gen13497 := Call(__e, ShenFunc(symhd), V1097)

												gen13498 := Call(__e, ShenFunc(symtl), gen13497)

												gen13499 := Call(__e, ShenFunc(symtl), gen13498)

												gen13500 := Call(__e, ShenFunc(symcons_2), gen13499)

												var gen13501 Obj
												if True == gen13500 {
													gen13491 := Call(__e, ShenFunc(symhd), V1097)

													gen13492 := Call(__e, ShenFunc(symtl), gen13491)

													gen13493 := Call(__e, ShenFunc(symtl), gen13492)

													gen13494 := Call(__e, ShenFunc(symtl), gen13493)

													gen13495 := Call(__e, ShenFunc(sym_a), Nil, gen13494)

													var gen13496 Obj
													if True == gen13495 {
														gen13488 := Call(__e, ShenFunc(symtl), V1097)

														gen13489 := Call(__e, ShenFunc(symcons_2), gen13488)

														var gen13490 Obj
														if True == gen13489 {
															gen13484 := Call(__e, ShenFunc(symtl), V1097)

															gen13485 := Call(__e, ShenFunc(symtl), gen13484)

															gen13486 := Call(__e, ShenFunc(sym_a), Nil, gen13485)

															var gen13487 Obj
															if True == gen13486 {
																gen13482 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1098)

																var gen13483 Obj
																if True == gen13482 {
																	gen13478 := Call(__e, ShenFunc(symhd), V1097)

																	gen13479 := Call(__e, ShenFunc(symtl), gen13478)

																	gen13480 := Call(__e, ShenFunc(symhd), gen13479)

																	gen13481 := Call(__e, ShenFunc(symshen_4prolog__constant_2), gen13480)

																	if True == gen13481 {
																		gen13483 = True
																	} else {
																		gen13483 = False
																	}

																} else {
																	gen13483 = False
																}
																if True == gen13483 {
																	gen13487 = True
																} else {
																	gen13487 = False
																}

															} else {
																gen13487 = False
															}
															if True == gen13487 {
																gen13490 = True
															} else {
																gen13490 = False
															}

														} else {
															gen13490 = False
														}
														if True == gen13490 {
															gen13496 = True
														} else {
															gen13496 = False
														}

													} else {
														gen13496 = False
													}
													if True == gen13496 {
														gen13501 = True
													} else {
														gen13501 = False
													}

												} else {
													gen13501 = False
												}
												if True == gen13501 {
													gen13505 = True
												} else {
													gen13505 = False
												}

											} else {
												gen13505 = False
											}
											if True == gen13505 {
												gen13509 = True
											} else {
												gen13509 = False
											}

										} else {
											gen13509 = False
										}
										if True == gen13509 {
											gen13512 = True
										} else {
											gen13512 = False
										}

									} else {
										gen13512 = False
									}
									if True == gen13512 {
										gen13514 = True
									} else {
										gen13514 = False
									}

								} else {
									gen13514 = False
								}
								if True == gen13514 {
									gen13424 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

									Z := gen13424
									gen13425 := Call(__e, ShenFunc(symtl), V1097)

									gen13426 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen13425)

									gen13427 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13426)

									gen13428 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen13427)

									gen13429 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13428)

									gen13430 := Call(__e, ShenFunc(symhd), V1097)

									gen13431 := Call(__e, ShenFunc(symtl), gen13430)

									gen13432 := Call(__e, ShenFunc(symhd), gen13431)

									gen13433 := Call(__e, ShenFunc(symcons), gen13432, Nil)

									gen13434 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen13433)

									gen13435 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen13434)

									gen13436 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13435)

									gen13437 := Call(__e, ShenFunc(symcons), Z, gen13436)

									gen13438 := Call(__e, ShenFunc(symhd), V1097)

									gen13439 := Call(__e, ShenFunc(symtl), gen13438)

									gen13440 := Call(__e, ShenFunc(symtl), gen13439)

									gen13441 := Call(__e, ShenFunc(symhd), gen13440)

									gen13442 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13441, MakeSymbol("+"))

									gen13443 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variable"), Nil)

									gen13444 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen13443)

									gen13445 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13444)

									gen13446 := Call(__e, ShenFunc(symcons), Z, gen13445)

									gen13447 := Call(__e, ShenFunc(symhd), V1097)

									gen13448 := Call(__e, ShenFunc(symtl), gen13447)

									gen13449 := Call(__e, ShenFunc(symhd), gen13448)

									gen13450 := Call(__e, ShenFunc(symhd), V1097)

									gen13451 := Call(__e, ShenFunc(symtl), gen13450)

									gen13452 := Call(__e, ShenFunc(symtl), gen13451)

									gen13453 := Call(__e, ShenFunc(symhd), gen13452)

									gen13454 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13453, MakeSymbol("+"))

									gen13455 := Call(__e, ShenFunc(symcons), gen13454, Nil)

									gen13456 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13455)

									gen13457 := Call(__e, ShenFunc(symcons), gen13449, gen13456)

									gen13458 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen13457)

									gen13459 := Call(__e, ShenFunc(symcons), Z, gen13458)

									gen13460 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen13459)

									gen13461 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

									gen13462 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13461)

									gen13463 := Call(__e, ShenFunc(symcons), gen13460, gen13462)

									gen13464 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13463)

									gen13465 := Call(__e, ShenFunc(symcons), gen13446, gen13464)

									gen13466 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13465)

									gen13467 := Call(__e, ShenFunc(symcons), gen13466, Nil)

									gen13468 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13467)

									gen13469 := Call(__e, ShenFunc(symcons), gen13442, gen13468)

									gen13470 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13469)

									gen13471 := Call(__e, ShenFunc(symcons), gen13437, gen13470)

									gen13472 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13471)

									gen13473 := Call(__e, ShenFunc(symcons), gen13472, Nil)

									gen13474 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13473)

									gen13475 := Call(__e, ShenFunc(symcons), gen13429, gen13474)

									gen13476 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen13475)

									gen13477 := Call(__e, ShenFunc(symcons), Z, gen13476)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen13477)

									return

								} else {
									gen13422 := Call(__e, ShenFunc(symcons_2), V1097)

									var gen13423 Obj
									if True == gen13422 {
										gen13419 := Call(__e, ShenFunc(symhd), V1097)

										gen13420 := Call(__e, ShenFunc(symcons_2), gen13419)

										var gen13421 Obj
										if True == gen13420 {
											gen13415 := Call(__e, ShenFunc(symhd), V1097)

											gen13416 := Call(__e, ShenFunc(symhd), gen13415)

											gen13417 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13416)

											var gen13418 Obj
											if True == gen13417 {
												gen13411 := Call(__e, ShenFunc(symhd), V1097)

												gen13412 := Call(__e, ShenFunc(symtl), gen13411)

												gen13413 := Call(__e, ShenFunc(symcons_2), gen13412)

												var gen13414 Obj
												if True == gen13413 {
													gen13406 := Call(__e, ShenFunc(symhd), V1097)

													gen13407 := Call(__e, ShenFunc(symtl), gen13406)

													gen13408 := Call(__e, ShenFunc(symhd), gen13407)

													gen13409 := Call(__e, ShenFunc(symcons_2), gen13408)

													var gen13410 Obj
													if True == gen13409 {
														gen13401 := Call(__e, ShenFunc(symhd), V1097)

														gen13402 := Call(__e, ShenFunc(symtl), gen13401)

														gen13403 := Call(__e, ShenFunc(symtl), gen13402)

														gen13404 := Call(__e, ShenFunc(symcons_2), gen13403)

														var gen13405 Obj
														if True == gen13404 {
															gen13395 := Call(__e, ShenFunc(symhd), V1097)

															gen13396 := Call(__e, ShenFunc(symtl), gen13395)

															gen13397 := Call(__e, ShenFunc(symtl), gen13396)

															gen13398 := Call(__e, ShenFunc(symtl), gen13397)

															gen13399 := Call(__e, ShenFunc(sym_a), Nil, gen13398)

															var gen13400 Obj
															if True == gen13399 {
																gen13392 := Call(__e, ShenFunc(symtl), V1097)

																gen13393 := Call(__e, ShenFunc(symcons_2), gen13392)

																var gen13394 Obj
																if True == gen13393 {
																	gen13388 := Call(__e, ShenFunc(symtl), V1097)

																	gen13389 := Call(__e, ShenFunc(symtl), gen13388)

																	gen13390 := Call(__e, ShenFunc(sym_a), Nil, gen13389)

																	var gen13391 Obj
																	if True == gen13390 {
																		gen13387 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V1098)

																		if True == gen13387 {
																			gen13391 = True
																		} else {
																			gen13391 = False
																		}

																	} else {
																		gen13391 = False
																	}
																	if True == gen13391 {
																		gen13394 = True
																	} else {
																		gen13394 = False
																	}

																} else {
																	gen13394 = False
																}
																if True == gen13394 {
																	gen13400 = True
																} else {
																	gen13400 = False
																}

															} else {
																gen13400 = False
															}
															if True == gen13400 {
																gen13405 = True
															} else {
																gen13405 = False
															}

														} else {
															gen13405 = False
														}
														if True == gen13405 {
															gen13410 = True
														} else {
															gen13410 = False
														}

													} else {
														gen13410 = False
													}
													if True == gen13410 {
														gen13414 = True
													} else {
														gen13414 = False
													}

												} else {
													gen13414 = False
												}
												if True == gen13414 {
													gen13418 = True
												} else {
													gen13418 = False
												}

											} else {
												gen13418 = False
											}
											if True == gen13418 {
												gen13421 = True
											} else {
												gen13421 = False
											}

										} else {
											gen13421 = False
										}
										if True == gen13421 {
											gen13423 = True
										} else {
											gen13423 = False
										}

									} else {
										gen13423 = False
									}
									if True == gen13423 {
										gen13336 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

										Z := gen13336
										gen13337 := Call(__e, ShenFunc(symtl), V1097)

										gen13338 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen13337)

										gen13339 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13338)

										gen13340 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen13339)

										gen13341 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13340)

										gen13342 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), Nil)

										gen13343 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.non-empty"), gen13342)

										gen13344 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen13343)

										gen13345 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13344)

										gen13346 := Call(__e, ShenFunc(symcons), Z, gen13345)

										gen13347 := Call(__e, ShenFunc(symhd), V1097)

										gen13348 := Call(__e, ShenFunc(symtl), gen13347)

										gen13349 := Call(__e, ShenFunc(symhd), gen13348)

										gen13350 := Call(__e, ShenFunc(symhd), gen13349)

										gen13351 := Call(__e, ShenFunc(symhd), V1097)

										gen13352 := Call(__e, ShenFunc(symtl), gen13351)

										gen13353 := Call(__e, ShenFunc(symhd), gen13352)

										gen13354 := Call(__e, ShenFunc(symtl), gen13353)

										gen13355 := Call(__e, ShenFunc(symhd), V1097)

										gen13356 := Call(__e, ShenFunc(symtl), gen13355)

										gen13357 := Call(__e, ShenFunc(symtl), gen13356)

										gen13358 := Call(__e, ShenFunc(symcons), gen13354, gen13357)

										gen13359 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13358)

										gen13360 := Call(__e, ShenFunc(symcons), Z, Nil)

										gen13361 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13360)

										gen13362 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen13361)

										gen13363 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13362)

										gen13364 := Call(__e, ShenFunc(symcons), gen13363, Nil)

										gen13365 := Call(__e, ShenFunc(symcons), gen13359, gen13364)

										gen13366 := Call(__e, ShenFunc(symcons), gen13365, Nil)

										gen13367 := Call(__e, ShenFunc(symcons), gen13350, gen13366)

										gen13368 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13367)

										gen13369 := Call(__e, ShenFunc(symcons), Z, Nil)

										gen13370 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13369)

										gen13371 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen13370)

										gen13372 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13371)

										gen13373 := Call(__e, ShenFunc(symcons), gen13372, Nil)

										gen13374 := Call(__e, ShenFunc(symcons), gen13368, gen13373)

										gen13375 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13374, MakeSymbol("-"))

										gen13376 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

										gen13377 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13376)

										gen13378 := Call(__e, ShenFunc(symcons), gen13375, gen13377)

										gen13379 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13378)

										gen13380 := Call(__e, ShenFunc(symcons), gen13346, gen13379)

										gen13381 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13380)

										gen13382 := Call(__e, ShenFunc(symcons), gen13381, Nil)

										gen13383 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13382)

										gen13384 := Call(__e, ShenFunc(symcons), gen13341, gen13383)

										gen13385 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen13384)

										gen13386 := Call(__e, ShenFunc(symcons), Z, gen13385)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen13386)

										return

									} else {
										gen13334 := Call(__e, ShenFunc(symcons_2), V1097)

										var gen13335 Obj
										if True == gen13334 {
											gen13331 := Call(__e, ShenFunc(symhd), V1097)

											gen13332 := Call(__e, ShenFunc(symcons_2), gen13331)

											var gen13333 Obj
											if True == gen13332 {
												gen13327 := Call(__e, ShenFunc(symhd), V1097)

												gen13328 := Call(__e, ShenFunc(symhd), gen13327)

												gen13329 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.mu"), gen13328)

												var gen13330 Obj
												if True == gen13329 {
													gen13323 := Call(__e, ShenFunc(symhd), V1097)

													gen13324 := Call(__e, ShenFunc(symtl), gen13323)

													gen13325 := Call(__e, ShenFunc(symcons_2), gen13324)

													var gen13326 Obj
													if True == gen13325 {
														gen13318 := Call(__e, ShenFunc(symhd), V1097)

														gen13319 := Call(__e, ShenFunc(symtl), gen13318)

														gen13320 := Call(__e, ShenFunc(symhd), gen13319)

														gen13321 := Call(__e, ShenFunc(symcons_2), gen13320)

														var gen13322 Obj
														if True == gen13321 {
															gen13313 := Call(__e, ShenFunc(symhd), V1097)

															gen13314 := Call(__e, ShenFunc(symtl), gen13313)

															gen13315 := Call(__e, ShenFunc(symtl), gen13314)

															gen13316 := Call(__e, ShenFunc(symcons_2), gen13315)

															var gen13317 Obj
															if True == gen13316 {
																gen13307 := Call(__e, ShenFunc(symhd), V1097)

																gen13308 := Call(__e, ShenFunc(symtl), gen13307)

																gen13309 := Call(__e, ShenFunc(symtl), gen13308)

																gen13310 := Call(__e, ShenFunc(symtl), gen13309)

																gen13311 := Call(__e, ShenFunc(sym_a), Nil, gen13310)

																var gen13312 Obj
																if True == gen13311 {
																	gen13304 := Call(__e, ShenFunc(symtl), V1097)

																	gen13305 := Call(__e, ShenFunc(symcons_2), gen13304)

																	var gen13306 Obj
																	if True == gen13305 {
																		gen13300 := Call(__e, ShenFunc(symtl), V1097)

																		gen13301 := Call(__e, ShenFunc(symtl), gen13300)

																		gen13302 := Call(__e, ShenFunc(sym_a), Nil, gen13301)

																		var gen13303 Obj
																		if True == gen13302 {
																			gen13299 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V1098)

																			if True == gen13299 {
																				gen13303 = True
																			} else {
																				gen13303 = False
																			}

																		} else {
																			gen13303 = False
																		}
																		if True == gen13303 {
																			gen13306 = True
																		} else {
																			gen13306 = False
																		}

																	} else {
																		gen13306 = False
																	}
																	if True == gen13306 {
																		gen13312 = True
																	} else {
																		gen13312 = False
																	}

																} else {
																	gen13312 = False
																}
																if True == gen13312 {
																	gen13317 = True
																} else {
																	gen13317 = False
																}

															} else {
																gen13317 = False
															}
															if True == gen13317 {
																gen13322 = True
															} else {
																gen13322 = False
															}

														} else {
															gen13322 = False
														}
														if True == gen13322 {
															gen13326 = True
														} else {
															gen13326 = False
														}

													} else {
														gen13326 = False
													}
													if True == gen13326 {
														gen13330 = True
													} else {
														gen13330 = False
													}

												} else {
													gen13330 = False
												}
												if True == gen13330 {
													gen13333 = True
												} else {
													gen13333 = False
												}

											} else {
												gen13333 = False
											}
											if True == gen13333 {
												gen13335 = True
											} else {
												gen13335 = False
											}

										} else {
											gen13335 = False
										}
										if True == gen13335 {
											gen13210 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

											Z := gen13210
											gen13211 := Call(__e, ShenFunc(symtl), V1097)

											gen13212 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dereferencing"), gen13211)

											gen13213 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13212)

											gen13214 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.result"), gen13213)

											gen13215 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13214)

											gen13216 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), Nil)

											gen13217 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.non-empty"), gen13216)

											gen13218 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen13217)

											gen13219 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13218)

											gen13220 := Call(__e, ShenFunc(symcons), Z, gen13219)

											gen13221 := Call(__e, ShenFunc(symhd), V1097)

											gen13222 := Call(__e, ShenFunc(symtl), gen13221)

											gen13223 := Call(__e, ShenFunc(symhd), gen13222)

											gen13224 := Call(__e, ShenFunc(symhd), gen13223)

											gen13225 := Call(__e, ShenFunc(symhd), V1097)

											gen13226 := Call(__e, ShenFunc(symtl), gen13225)

											gen13227 := Call(__e, ShenFunc(symhd), gen13226)

											gen13228 := Call(__e, ShenFunc(symtl), gen13227)

											gen13229 := Call(__e, ShenFunc(symhd), V1097)

											gen13230 := Call(__e, ShenFunc(symtl), gen13229)

											gen13231 := Call(__e, ShenFunc(symtl), gen13230)

											gen13232 := Call(__e, ShenFunc(symcons), gen13228, gen13231)

											gen13233 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13232)

											gen13234 := Call(__e, ShenFunc(symcons), Z, Nil)

											gen13235 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13234)

											gen13236 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen13235)

											gen13237 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13236)

											gen13238 := Call(__e, ShenFunc(symcons), gen13237, Nil)

											gen13239 := Call(__e, ShenFunc(symcons), gen13233, gen13238)

											gen13240 := Call(__e, ShenFunc(symcons), gen13239, Nil)

											gen13241 := Call(__e, ShenFunc(symcons), gen13224, gen13240)

											gen13242 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.mu"), gen13241)

											gen13243 := Call(__e, ShenFunc(symcons), Z, Nil)

											gen13244 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.of"), gen13243)

											gen13245 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen13244)

											gen13246 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13245)

											gen13247 := Call(__e, ShenFunc(symcons), gen13246, Nil)

											gen13248 := Call(__e, ShenFunc(symcons), gen13242, gen13247)

											gen13249 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13248, MakeSymbol("+"))

											gen13250 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variable"), Nil)

											gen13251 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.a"), gen13250)

											gen13252 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen13251)

											gen13253 := Call(__e, ShenFunc(symcons), Z, gen13252)

											gen13254 := Call(__e, ShenFunc(symhd), V1097)

											gen13255 := Call(__e, ShenFunc(symtl), gen13254)

											gen13256 := Call(__e, ShenFunc(symhd), gen13255)

											gen13257 := Call(__e, ShenFunc(symshen_4extract__vars), gen13256)

											gen13258 := Call(__e, ShenFunc(symhd), V1097)

											gen13259 := Call(__e, ShenFunc(symtl), gen13258)

											gen13260 := Call(__e, ShenFunc(symhd), gen13259)

											gen13261 := Call(__e, ShenFunc(symshen_4remove__modes), gen13260)

											gen13262 := Call(__e, ShenFunc(symshen_4rcons__form), gen13261)

											gen13263 := Call(__e, ShenFunc(symhd), V1097)

											gen13264 := Call(__e, ShenFunc(symtl), gen13263)

											gen13265 := Call(__e, ShenFunc(symtl), gen13264)

											gen13266 := Call(__e, ShenFunc(symhd), gen13265)

											gen13267 := Call(__e, ShenFunc(symshen_4mu__reduction), gen13266, MakeSymbol("+"))

											gen13268 := Call(__e, ShenFunc(symcons), gen13267, Nil)

											gen13269 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13268)

											gen13270 := Call(__e, ShenFunc(symcons), gen13262, gen13269)

											gen13271 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.to"), gen13270)

											gen13272 := Call(__e, ShenFunc(symcons), Z, gen13271)

											gen13273 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen13272)

											gen13274 := Call(__e, ShenFunc(symcons), gen13273, Nil)

											gen13275 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13274)

											gen13276 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen13275)

											gen13277 := Call(__e, ShenFunc(symcons), gen13257, gen13276)

											gen13278 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13277)

											gen13279 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen13278)

											gen13280 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen13279)

											gen13281 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.rename"), gen13280)

											gen13282 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.failed!"), Nil)

											gen13283 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13282)

											gen13284 := Call(__e, ShenFunc(symcons), gen13281, gen13283)

											gen13285 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13284)

											gen13286 := Call(__e, ShenFunc(symcons), gen13253, gen13285)

											gen13287 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13286)

											gen13288 := Call(__e, ShenFunc(symcons), gen13287, Nil)

											gen13289 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.else"), gen13288)

											gen13290 := Call(__e, ShenFunc(symcons), gen13249, gen13289)

											gen13291 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.then"), gen13290)

											gen13292 := Call(__e, ShenFunc(symcons), gen13220, gen13291)

											gen13293 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen13292)

											gen13294 := Call(__e, ShenFunc(symcons), gen13293, Nil)

											gen13295 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen13294)

											gen13296 := Call(__e, ShenFunc(symcons), gen13215, gen13295)

											gen13297 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.be"), gen13296)

											gen13298 := Call(__e, ShenFunc(symcons), Z, gen13297)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen13298)

											return

										} else {
											__e.Return(V1097)
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mu_reduction"), gen13798)

		gen13806 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1100 := __args[0]
			_ = V1100
			gen13805 := Call(__e, ShenFunc(symcons_2), V1100)

			if True == gen13805 {
				gen13799 := Call(__e, ShenFunc(symhd), V1100)

				gen13800 := Call(__e, ShenFunc(symshen_4rcons__form), gen13799)

				gen13801 := Call(__e, ShenFunc(symtl), V1100)

				gen13802 := Call(__e, ShenFunc(symshen_4rcons__form), gen13801)

				gen13803 := Call(__e, ShenFunc(symcons), gen13802, Nil)

				gen13804 := Call(__e, ShenFunc(symcons), gen13800, gen13803)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("cons"), gen13804)

				return

			} else {
				__e.Return(V1100)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.rcons_form"), gen13806)

		gen13858 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1102 := __args[0]
			_ = V1102
			gen13856 := Call(__e, ShenFunc(symcons_2), V1102)

			var gen13857 Obj
			if True == gen13856 {
				gen13853 := Call(__e, ShenFunc(symhd), V1102)

				gen13854 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen13853)

				var gen13855 Obj
				if True == gen13854 {
					gen13850 := Call(__e, ShenFunc(symtl), V1102)

					gen13851 := Call(__e, ShenFunc(symcons_2), gen13850)

					var gen13852 Obj
					if True == gen13851 {
						gen13846 := Call(__e, ShenFunc(symtl), V1102)

						gen13847 := Call(__e, ShenFunc(symtl), gen13846)

						gen13848 := Call(__e, ShenFunc(symcons_2), gen13847)

						var gen13849 Obj
						if True == gen13848 {
							gen13841 := Call(__e, ShenFunc(symtl), V1102)

							gen13842 := Call(__e, ShenFunc(symtl), gen13841)

							gen13843 := Call(__e, ShenFunc(symhd), gen13842)

							gen13844 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), gen13843)

							var gen13845 Obj
							if True == gen13844 {
								gen13837 := Call(__e, ShenFunc(symtl), V1102)

								gen13838 := Call(__e, ShenFunc(symtl), gen13837)

								gen13839 := Call(__e, ShenFunc(symtl), gen13838)

								gen13840 := Call(__e, ShenFunc(sym_a), Nil, gen13839)

								if True == gen13840 {
									gen13845 = True
								} else {
									gen13845 = False
								}

							} else {
								gen13845 = False
							}
							if True == gen13845 {
								gen13849 = True
							} else {
								gen13849 = False
							}

						} else {
							gen13849 = False
						}
						if True == gen13849 {
							gen13852 = True
						} else {
							gen13852 = False
						}

					} else {
						gen13852 = False
					}
					if True == gen13852 {
						gen13855 = True
					} else {
						gen13855 = False
					}

				} else {
					gen13855 = False
				}
				if True == gen13855 {
					gen13857 = True
				} else {
					gen13857 = False
				}

			} else {
				gen13857 = False
			}
			if True == gen13857 {
				gen13835 := Call(__e, ShenFunc(symtl), V1102)

				gen13836 := Call(__e, ShenFunc(symhd), gen13835)

				__e.TailApply(ShenFunc(symshen_4remove__modes), gen13836)

				return

			} else {
				gen13833 := Call(__e, ShenFunc(symcons_2), V1102)

				var gen13834 Obj
				if True == gen13833 {
					gen13830 := Call(__e, ShenFunc(symhd), V1102)

					gen13831 := Call(__e, ShenFunc(sym_a), MakeSymbol("mode"), gen13830)

					var gen13832 Obj
					if True == gen13831 {
						gen13827 := Call(__e, ShenFunc(symtl), V1102)

						gen13828 := Call(__e, ShenFunc(symcons_2), gen13827)

						var gen13829 Obj
						if True == gen13828 {
							gen13823 := Call(__e, ShenFunc(symtl), V1102)

							gen13824 := Call(__e, ShenFunc(symtl), gen13823)

							gen13825 := Call(__e, ShenFunc(symcons_2), gen13824)

							var gen13826 Obj
							if True == gen13825 {
								gen13818 := Call(__e, ShenFunc(symtl), V1102)

								gen13819 := Call(__e, ShenFunc(symtl), gen13818)

								gen13820 := Call(__e, ShenFunc(symhd), gen13819)

								gen13821 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), gen13820)

								var gen13822 Obj
								if True == gen13821 {
									gen13814 := Call(__e, ShenFunc(symtl), V1102)

									gen13815 := Call(__e, ShenFunc(symtl), gen13814)

									gen13816 := Call(__e, ShenFunc(symtl), gen13815)

									gen13817 := Call(__e, ShenFunc(sym_a), Nil, gen13816)

									if True == gen13817 {
										gen13822 = True
									} else {
										gen13822 = False
									}

								} else {
									gen13822 = False
								}
								if True == gen13822 {
									gen13826 = True
								} else {
									gen13826 = False
								}

							} else {
								gen13826 = False
							}
							if True == gen13826 {
								gen13829 = True
							} else {
								gen13829 = False
							}

						} else {
							gen13829 = False
						}
						if True == gen13829 {
							gen13832 = True
						} else {
							gen13832 = False
						}

					} else {
						gen13832 = False
					}
					if True == gen13832 {
						gen13834 = True
					} else {
						gen13834 = False
					}

				} else {
					gen13834 = False
				}
				if True == gen13834 {
					gen13812 := Call(__e, ShenFunc(symtl), V1102)

					gen13813 := Call(__e, ShenFunc(symhd), gen13812)

					__e.TailApply(ShenFunc(symshen_4remove__modes), gen13813)

					return

				} else {
					gen13811 := Call(__e, ShenFunc(symcons_2), V1102)

					if True == gen13811 {
						gen13807 := Call(__e, ShenFunc(symhd), V1102)

						gen13808 := Call(__e, ShenFunc(symshen_4remove__modes), gen13807)

						gen13809 := Call(__e, ShenFunc(symtl), V1102)

						gen13810 := Call(__e, ShenFunc(symshen_4remove__modes), gen13809)

						__e.TailApply(ShenFunc(symcons), gen13808, gen13810)

						return

					} else {
						__e.Return(V1102)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.remove_modes"), gen13858)

		gen13861 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1105 := __args[0]
			_ = V1105
			V1106 := __args[1]
			_ = V1106
			gen13860 := Call(__e, ShenFunc(symvariable_2), V1105)

			if True == gen13860 {
				gen13859 := Call(__e, ShenFunc(symvariable_2), V1106)

				if True == gen13859 {
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

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ephemeral_variable?"), gen13861)

		gen13863 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1116 := __args[0]
			_ = V1116
			gen13862 := Call(__e, ShenFunc(symcons_2), V1116)

			if True == gen13862 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog_constant?"), gen13863)

		gen14630 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1118 := __args[0]
			_ = V1118
			gen14628 := Call(__e, ShenFunc(symcons_2), V1118)

			var gen14629 Obj
			if True == gen14628 {
				gen14625 := Call(__e, ShenFunc(symhd), V1118)

				gen14626 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen14625)

				var gen14627 Obj
				if True == gen14626 {
					gen14622 := Call(__e, ShenFunc(symtl), V1118)

					gen14623 := Call(__e, ShenFunc(symcons_2), gen14622)

					var gen14624 Obj
					if True == gen14623 {
						gen14618 := Call(__e, ShenFunc(symtl), V1118)

						gen14619 := Call(__e, ShenFunc(symtl), gen14618)

						gen14620 := Call(__e, ShenFunc(symcons_2), gen14619)

						var gen14621 Obj
						if True == gen14620 {
							gen14613 := Call(__e, ShenFunc(symtl), V1118)

							gen14614 := Call(__e, ShenFunc(symtl), gen14613)

							gen14615 := Call(__e, ShenFunc(symhd), gen14614)

							gen14616 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.be"), gen14615)

							var gen14617 Obj
							if True == gen14616 {
								gen14608 := Call(__e, ShenFunc(symtl), V1118)

								gen14609 := Call(__e, ShenFunc(symtl), gen14608)

								gen14610 := Call(__e, ShenFunc(symtl), gen14609)

								gen14611 := Call(__e, ShenFunc(symcons_2), gen14610)

								var gen14612 Obj
								if True == gen14611 {
									gen14602 := Call(__e, ShenFunc(symtl), V1118)

									gen14603 := Call(__e, ShenFunc(symtl), gen14602)

									gen14604 := Call(__e, ShenFunc(symtl), gen14603)

									gen14605 := Call(__e, ShenFunc(symtl), gen14604)

									gen14606 := Call(__e, ShenFunc(symcons_2), gen14605)

									var gen14607 Obj
									if True == gen14606 {
										gen14595 := Call(__e, ShenFunc(symtl), V1118)

										gen14596 := Call(__e, ShenFunc(symtl), gen14595)

										gen14597 := Call(__e, ShenFunc(symtl), gen14596)

										gen14598 := Call(__e, ShenFunc(symtl), gen14597)

										gen14599 := Call(__e, ShenFunc(symhd), gen14598)

										gen14600 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen14599)

										var gen14601 Obj
										if True == gen14600 {
											gen14588 := Call(__e, ShenFunc(symtl), V1118)

											gen14589 := Call(__e, ShenFunc(symtl), gen14588)

											gen14590 := Call(__e, ShenFunc(symtl), gen14589)

											gen14591 := Call(__e, ShenFunc(symtl), gen14590)

											gen14592 := Call(__e, ShenFunc(symtl), gen14591)

											gen14593 := Call(__e, ShenFunc(symcons_2), gen14592)

											var gen14594 Obj
											if True == gen14593 {
												gen14581 := Call(__e, ShenFunc(symtl), V1118)

												gen14582 := Call(__e, ShenFunc(symtl), gen14581)

												gen14583 := Call(__e, ShenFunc(symtl), gen14582)

												gen14584 := Call(__e, ShenFunc(symtl), gen14583)

												gen14585 := Call(__e, ShenFunc(symtl), gen14584)

												gen14586 := Call(__e, ShenFunc(symtl), gen14585)

												gen14587 := Call(__e, ShenFunc(sym_a), Nil, gen14586)

												if True == gen14587 {
													gen14594 = True
												} else {
													gen14594 = False
												}

											} else {
												gen14594 = False
											}
											if True == gen14594 {
												gen14601 = True
											} else {
												gen14601 = False
											}

										} else {
											gen14601 = False
										}
										if True == gen14601 {
											gen14607 = True
										} else {
											gen14607 = False
										}

									} else {
										gen14607 = False
									}
									if True == gen14607 {
										gen14612 = True
									} else {
										gen14612 = False
									}

								} else {
									gen14612 = False
								}
								if True == gen14612 {
									gen14617 = True
								} else {
									gen14617 = False
								}

							} else {
								gen14617 = False
							}
							if True == gen14617 {
								gen14621 = True
							} else {
								gen14621 = False
							}

						} else {
							gen14621 = False
						}
						if True == gen14621 {
							gen14624 = True
						} else {
							gen14624 = False
						}

					} else {
						gen14624 = False
					}
					if True == gen14624 {
						gen14627 = True
					} else {
						gen14627 = False
					}

				} else {
					gen14627 = False
				}
				if True == gen14627 {
					gen14629 = True
				} else {
					gen14629 = False
				}

			} else {
				gen14629 = False
			}
			if True == gen14629 {
				gen14564 := Call(__e, ShenFunc(symtl), V1118)

				gen14565 := Call(__e, ShenFunc(symhd), gen14564)

				gen14566 := Call(__e, ShenFunc(symtl), V1118)

				gen14567 := Call(__e, ShenFunc(symtl), gen14566)

				gen14568 := Call(__e, ShenFunc(symtl), gen14567)

				gen14569 := Call(__e, ShenFunc(symhd), gen14568)

				gen14570 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14569)

				gen14571 := Call(__e, ShenFunc(symtl), V1118)

				gen14572 := Call(__e, ShenFunc(symtl), gen14571)

				gen14573 := Call(__e, ShenFunc(symtl), gen14572)

				gen14574 := Call(__e, ShenFunc(symtl), gen14573)

				gen14575 := Call(__e, ShenFunc(symtl), gen14574)

				gen14576 := Call(__e, ShenFunc(symhd), gen14575)

				gen14577 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14576)

				gen14578 := Call(__e, ShenFunc(symcons), gen14577, Nil)

				gen14579 := Call(__e, ShenFunc(symcons), gen14570, gen14578)

				gen14580 := Call(__e, ShenFunc(symcons), gen14565, gen14579)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen14580)

				return

			} else {
				gen14562 := Call(__e, ShenFunc(symcons_2), V1118)

				var gen14563 Obj
				if True == gen14562 {
					gen14559 := Call(__e, ShenFunc(symhd), V1118)

					gen14560 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen14559)

					var gen14561 Obj
					if True == gen14560 {
						gen14556 := Call(__e, ShenFunc(symtl), V1118)

						gen14557 := Call(__e, ShenFunc(symcons_2), gen14556)

						var gen14558 Obj
						if True == gen14557 {
							gen14552 := Call(__e, ShenFunc(symtl), V1118)

							gen14553 := Call(__e, ShenFunc(symhd), gen14552)

							gen14554 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.result"), gen14553)

							var gen14555 Obj
							if True == gen14554 {
								gen14548 := Call(__e, ShenFunc(symtl), V1118)

								gen14549 := Call(__e, ShenFunc(symtl), gen14548)

								gen14550 := Call(__e, ShenFunc(symcons_2), gen14549)

								var gen14551 Obj
								if True == gen14550 {
									gen14543 := Call(__e, ShenFunc(symtl), V1118)

									gen14544 := Call(__e, ShenFunc(symtl), gen14543)

									gen14545 := Call(__e, ShenFunc(symhd), gen14544)

									gen14546 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen14545)

									var gen14547 Obj
									if True == gen14546 {
										gen14538 := Call(__e, ShenFunc(symtl), V1118)

										gen14539 := Call(__e, ShenFunc(symtl), gen14538)

										gen14540 := Call(__e, ShenFunc(symtl), gen14539)

										gen14541 := Call(__e, ShenFunc(symcons_2), gen14540)

										var gen14542 Obj
										if True == gen14541 {
											gen14532 := Call(__e, ShenFunc(symtl), V1118)

											gen14533 := Call(__e, ShenFunc(symtl), gen14532)

											gen14534 := Call(__e, ShenFunc(symtl), gen14533)

											gen14535 := Call(__e, ShenFunc(symhd), gen14534)

											gen14536 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.dereferencing"), gen14535)

											var gen14537 Obj
											if True == gen14536 {
												gen14526 := Call(__e, ShenFunc(symtl), V1118)

												gen14527 := Call(__e, ShenFunc(symtl), gen14526)

												gen14528 := Call(__e, ShenFunc(symtl), gen14527)

												gen14529 := Call(__e, ShenFunc(symtl), gen14528)

												gen14530 := Call(__e, ShenFunc(symcons_2), gen14529)

												var gen14531 Obj
												if True == gen14530 {
													gen14520 := Call(__e, ShenFunc(symtl), V1118)

													gen14521 := Call(__e, ShenFunc(symtl), gen14520)

													gen14522 := Call(__e, ShenFunc(symtl), gen14521)

													gen14523 := Call(__e, ShenFunc(symtl), gen14522)

													gen14524 := Call(__e, ShenFunc(symtl), gen14523)

													gen14525 := Call(__e, ShenFunc(sym_a), Nil, gen14524)

													if True == gen14525 {
														gen14531 = True
													} else {
														gen14531 = False
													}

												} else {
													gen14531 = False
												}
												if True == gen14531 {
													gen14537 = True
												} else {
													gen14537 = False
												}

											} else {
												gen14537 = False
											}
											if True == gen14537 {
												gen14542 = True
											} else {
												gen14542 = False
											}

										} else {
											gen14542 = False
										}
										if True == gen14542 {
											gen14547 = True
										} else {
											gen14547 = False
										}

									} else {
										gen14547 = False
									}
									if True == gen14547 {
										gen14551 = True
									} else {
										gen14551 = False
									}

								} else {
									gen14551 = False
								}
								if True == gen14551 {
									gen14555 = True
								} else {
									gen14555 = False
								}

							} else {
								gen14555 = False
							}
							if True == gen14555 {
								gen14558 = True
							} else {
								gen14558 = False
							}

						} else {
							gen14558 = False
						}
						if True == gen14558 {
							gen14561 = True
						} else {
							gen14561 = False
						}

					} else {
						gen14561 = False
					}
					if True == gen14561 {
						gen14563 = True
					} else {
						gen14563 = False
					}

				} else {
					gen14563 = False
				}
				if True == gen14563 {
					gen14512 := Call(__e, ShenFunc(symtl), V1118)

					gen14513 := Call(__e, ShenFunc(symtl), gen14512)

					gen14514 := Call(__e, ShenFunc(symtl), gen14513)

					gen14515 := Call(__e, ShenFunc(symtl), gen14514)

					gen14516 := Call(__e, ShenFunc(symhd), gen14515)

					gen14517 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14516)

					gen14518 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

					gen14519 := Call(__e, ShenFunc(symcons), gen14517, gen14518)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.lazyderef"), gen14519)

					return

				} else {
					gen14510 := Call(__e, ShenFunc(symcons_2), V1118)

					var gen14511 Obj
					if True == gen14510 {
						gen14507 := Call(__e, ShenFunc(symhd), V1118)

						gen14508 := Call(__e, ShenFunc(sym_a), MakeSymbol("if"), gen14507)

						var gen14509 Obj
						if True == gen14508 {
							gen14504 := Call(__e, ShenFunc(symtl), V1118)

							gen14505 := Call(__e, ShenFunc(symcons_2), gen14504)

							var gen14506 Obj
							if True == gen14505 {
								gen14500 := Call(__e, ShenFunc(symtl), V1118)

								gen14501 := Call(__e, ShenFunc(symtl), gen14500)

								gen14502 := Call(__e, ShenFunc(symcons_2), gen14501)

								var gen14503 Obj
								if True == gen14502 {
									gen14495 := Call(__e, ShenFunc(symtl), V1118)

									gen14496 := Call(__e, ShenFunc(symtl), gen14495)

									gen14497 := Call(__e, ShenFunc(symhd), gen14496)

									gen14498 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen14497)

									var gen14499 Obj
									if True == gen14498 {
										gen14490 := Call(__e, ShenFunc(symtl), V1118)

										gen14491 := Call(__e, ShenFunc(symtl), gen14490)

										gen14492 := Call(__e, ShenFunc(symtl), gen14491)

										gen14493 := Call(__e, ShenFunc(symcons_2), gen14492)

										var gen14494 Obj
										if True == gen14493 {
											gen14484 := Call(__e, ShenFunc(symtl), V1118)

											gen14485 := Call(__e, ShenFunc(symtl), gen14484)

											gen14486 := Call(__e, ShenFunc(symtl), gen14485)

											gen14487 := Call(__e, ShenFunc(symtl), gen14486)

											gen14488 := Call(__e, ShenFunc(symcons_2), gen14487)

											var gen14489 Obj
											if True == gen14488 {
												gen14477 := Call(__e, ShenFunc(symtl), V1118)

												gen14478 := Call(__e, ShenFunc(symtl), gen14477)

												gen14479 := Call(__e, ShenFunc(symtl), gen14478)

												gen14480 := Call(__e, ShenFunc(symtl), gen14479)

												gen14481 := Call(__e, ShenFunc(symhd), gen14480)

												gen14482 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.else"), gen14481)

												var gen14483 Obj
												if True == gen14482 {
													gen14470 := Call(__e, ShenFunc(symtl), V1118)

													gen14471 := Call(__e, ShenFunc(symtl), gen14470)

													gen14472 := Call(__e, ShenFunc(symtl), gen14471)

													gen14473 := Call(__e, ShenFunc(symtl), gen14472)

													gen14474 := Call(__e, ShenFunc(symtl), gen14473)

													gen14475 := Call(__e, ShenFunc(symcons_2), gen14474)

													var gen14476 Obj
													if True == gen14475 {
														gen14463 := Call(__e, ShenFunc(symtl), V1118)

														gen14464 := Call(__e, ShenFunc(symtl), gen14463)

														gen14465 := Call(__e, ShenFunc(symtl), gen14464)

														gen14466 := Call(__e, ShenFunc(symtl), gen14465)

														gen14467 := Call(__e, ShenFunc(symtl), gen14466)

														gen14468 := Call(__e, ShenFunc(symtl), gen14467)

														gen14469 := Call(__e, ShenFunc(sym_a), Nil, gen14468)

														if True == gen14469 {
															gen14476 = True
														} else {
															gen14476 = False
														}

													} else {
														gen14476 = False
													}
													if True == gen14476 {
														gen14483 = True
													} else {
														gen14483 = False
													}

												} else {
													gen14483 = False
												}
												if True == gen14483 {
													gen14489 = True
												} else {
													gen14489 = False
												}

											} else {
												gen14489 = False
											}
											if True == gen14489 {
												gen14494 = True
											} else {
												gen14494 = False
											}

										} else {
											gen14494 = False
										}
										if True == gen14494 {
											gen14499 = True
										} else {
											gen14499 = False
										}

									} else {
										gen14499 = False
									}
									if True == gen14499 {
										gen14503 = True
									} else {
										gen14503 = False
									}

								} else {
									gen14503 = False
								}
								if True == gen14503 {
									gen14506 = True
								} else {
									gen14506 = False
								}

							} else {
								gen14506 = False
							}
							if True == gen14506 {
								gen14509 = True
							} else {
								gen14509 = False
							}

						} else {
							gen14509 = False
						}
						if True == gen14509 {
							gen14511 = True
						} else {
							gen14511 = False
						}

					} else {
						gen14511 = False
					}
					if True == gen14511 {
						gen14445 := Call(__e, ShenFunc(symtl), V1118)

						gen14446 := Call(__e, ShenFunc(symhd), gen14445)

						gen14447 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14446)

						gen14448 := Call(__e, ShenFunc(symtl), V1118)

						gen14449 := Call(__e, ShenFunc(symtl), gen14448)

						gen14450 := Call(__e, ShenFunc(symtl), gen14449)

						gen14451 := Call(__e, ShenFunc(symhd), gen14450)

						gen14452 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14451)

						gen14453 := Call(__e, ShenFunc(symtl), V1118)

						gen14454 := Call(__e, ShenFunc(symtl), gen14453)

						gen14455 := Call(__e, ShenFunc(symtl), gen14454)

						gen14456 := Call(__e, ShenFunc(symtl), gen14455)

						gen14457 := Call(__e, ShenFunc(symtl), gen14456)

						gen14458 := Call(__e, ShenFunc(symhd), gen14457)

						gen14459 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14458)

						gen14460 := Call(__e, ShenFunc(symcons), gen14459, Nil)

						gen14461 := Call(__e, ShenFunc(symcons), gen14452, gen14460)

						gen14462 := Call(__e, ShenFunc(symcons), gen14447, gen14461)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen14462)

						return

					} else {
						gen14443 := Call(__e, ShenFunc(symcons_2), V1118)

						var gen14444 Obj
						if True == gen14443 {
							gen14440 := Call(__e, ShenFunc(symtl), V1118)

							gen14441 := Call(__e, ShenFunc(symcons_2), gen14440)

							var gen14442 Obj
							if True == gen14441 {
								gen14436 := Call(__e, ShenFunc(symtl), V1118)

								gen14437 := Call(__e, ShenFunc(symhd), gen14436)

								gen14438 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen14437)

								var gen14439 Obj
								if True == gen14438 {
									gen14432 := Call(__e, ShenFunc(symtl), V1118)

									gen14433 := Call(__e, ShenFunc(symtl), gen14432)

									gen14434 := Call(__e, ShenFunc(symcons_2), gen14433)

									var gen14435 Obj
									if True == gen14434 {
										gen14427 := Call(__e, ShenFunc(symtl), V1118)

										gen14428 := Call(__e, ShenFunc(symtl), gen14427)

										gen14429 := Call(__e, ShenFunc(symhd), gen14428)

										gen14430 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.a"), gen14429)

										var gen14431 Obj
										if True == gen14430 {
											gen14422 := Call(__e, ShenFunc(symtl), V1118)

											gen14423 := Call(__e, ShenFunc(symtl), gen14422)

											gen14424 := Call(__e, ShenFunc(symtl), gen14423)

											gen14425 := Call(__e, ShenFunc(symcons_2), gen14424)

											var gen14426 Obj
											if True == gen14425 {
												gen14416 := Call(__e, ShenFunc(symtl), V1118)

												gen14417 := Call(__e, ShenFunc(symtl), gen14416)

												gen14418 := Call(__e, ShenFunc(symtl), gen14417)

												gen14419 := Call(__e, ShenFunc(symhd), gen14418)

												gen14420 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variable"), gen14419)

												var gen14421 Obj
												if True == gen14420 {
													gen14411 := Call(__e, ShenFunc(symtl), V1118)

													gen14412 := Call(__e, ShenFunc(symtl), gen14411)

													gen14413 := Call(__e, ShenFunc(symtl), gen14412)

													gen14414 := Call(__e, ShenFunc(symtl), gen14413)

													gen14415 := Call(__e, ShenFunc(sym_a), Nil, gen14414)

													if True == gen14415 {
														gen14421 = True
													} else {
														gen14421 = False
													}

												} else {
													gen14421 = False
												}
												if True == gen14421 {
													gen14426 = True
												} else {
													gen14426 = False
												}

											} else {
												gen14426 = False
											}
											if True == gen14426 {
												gen14431 = True
											} else {
												gen14431 = False
											}

										} else {
											gen14431 = False
										}
										if True == gen14431 {
											gen14435 = True
										} else {
											gen14435 = False
										}

									} else {
										gen14435 = False
									}
									if True == gen14435 {
										gen14439 = True
									} else {
										gen14439 = False
									}

								} else {
									gen14439 = False
								}
								if True == gen14439 {
									gen14442 = True
								} else {
									gen14442 = False
								}

							} else {
								gen14442 = False
							}
							if True == gen14442 {
								gen14444 = True
							} else {
								gen14444 = False
							}

						} else {
							gen14444 = False
						}
						if True == gen14444 {
							gen14409 := Call(__e, ShenFunc(symhd), V1118)

							gen14410 := Call(__e, ShenFunc(symcons), gen14409, Nil)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.pvar?"), gen14410)

							return

						} else {
							gen14407 := Call(__e, ShenFunc(symcons_2), V1118)

							var gen14408 Obj
							if True == gen14407 {
								gen14404 := Call(__e, ShenFunc(symtl), V1118)

								gen14405 := Call(__e, ShenFunc(symcons_2), gen14404)

								var gen14406 Obj
								if True == gen14405 {
									gen14400 := Call(__e, ShenFunc(symtl), V1118)

									gen14401 := Call(__e, ShenFunc(symhd), gen14400)

									gen14402 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen14401)

									var gen14403 Obj
									if True == gen14402 {
										gen14396 := Call(__e, ShenFunc(symtl), V1118)

										gen14397 := Call(__e, ShenFunc(symtl), gen14396)

										gen14398 := Call(__e, ShenFunc(symcons_2), gen14397)

										var gen14399 Obj
										if True == gen14398 {
											gen14391 := Call(__e, ShenFunc(symtl), V1118)

											gen14392 := Call(__e, ShenFunc(symtl), gen14391)

											gen14393 := Call(__e, ShenFunc(symhd), gen14392)

											gen14394 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.a"), gen14393)

											var gen14395 Obj
											if True == gen14394 {
												gen14386 := Call(__e, ShenFunc(symtl), V1118)

												gen14387 := Call(__e, ShenFunc(symtl), gen14386)

												gen14388 := Call(__e, ShenFunc(symtl), gen14387)

												gen14389 := Call(__e, ShenFunc(symcons_2), gen14388)

												var gen14390 Obj
												if True == gen14389 {
													gen14380 := Call(__e, ShenFunc(symtl), V1118)

													gen14381 := Call(__e, ShenFunc(symtl), gen14380)

													gen14382 := Call(__e, ShenFunc(symtl), gen14381)

													gen14383 := Call(__e, ShenFunc(symhd), gen14382)

													gen14384 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.non-empty"), gen14383)

													var gen14385 Obj
													if True == gen14384 {
														gen14374 := Call(__e, ShenFunc(symtl), V1118)

														gen14375 := Call(__e, ShenFunc(symtl), gen14374)

														gen14376 := Call(__e, ShenFunc(symtl), gen14375)

														gen14377 := Call(__e, ShenFunc(symtl), gen14376)

														gen14378 := Call(__e, ShenFunc(symcons_2), gen14377)

														var gen14379 Obj
														if True == gen14378 {
															gen14367 := Call(__e, ShenFunc(symtl), V1118)

															gen14368 := Call(__e, ShenFunc(symtl), gen14367)

															gen14369 := Call(__e, ShenFunc(symtl), gen14368)

															gen14370 := Call(__e, ShenFunc(symtl), gen14369)

															gen14371 := Call(__e, ShenFunc(symhd), gen14370)

															gen14372 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), gen14371)

															var gen14373 Obj
															if True == gen14372 {
																gen14361 := Call(__e, ShenFunc(symtl), V1118)

																gen14362 := Call(__e, ShenFunc(symtl), gen14361)

																gen14363 := Call(__e, ShenFunc(symtl), gen14362)

																gen14364 := Call(__e, ShenFunc(symtl), gen14363)

																gen14365 := Call(__e, ShenFunc(symtl), gen14364)

																gen14366 := Call(__e, ShenFunc(sym_a), Nil, gen14365)

																if True == gen14366 {
																	gen14373 = True
																} else {
																	gen14373 = False
																}

															} else {
																gen14373 = False
															}
															if True == gen14373 {
																gen14379 = True
															} else {
																gen14379 = False
															}

														} else {
															gen14379 = False
														}
														if True == gen14379 {
															gen14385 = True
														} else {
															gen14385 = False
														}

													} else {
														gen14385 = False
													}
													if True == gen14385 {
														gen14390 = True
													} else {
														gen14390 = False
													}

												} else {
													gen14390 = False
												}
												if True == gen14390 {
													gen14395 = True
												} else {
													gen14395 = False
												}

											} else {
												gen14395 = False
											}
											if True == gen14395 {
												gen14399 = True
											} else {
												gen14399 = False
											}

										} else {
											gen14399 = False
										}
										if True == gen14399 {
											gen14403 = True
										} else {
											gen14403 = False
										}

									} else {
										gen14403 = False
									}
									if True == gen14403 {
										gen14406 = True
									} else {
										gen14406 = False
									}

								} else {
									gen14406 = False
								}
								if True == gen14406 {
									gen14408 = True
								} else {
									gen14408 = False
								}

							} else {
								gen14408 = False
							}
							if True == gen14408 {
								gen14359 := Call(__e, ShenFunc(symhd), V1118)

								gen14360 := Call(__e, ShenFunc(symcons), gen14359, Nil)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("cons?"), gen14360)

								return

							} else {
								gen14357 := Call(__e, ShenFunc(symcons_2), V1118)

								var gen14358 Obj
								if True == gen14357 {
									gen14354 := Call(__e, ShenFunc(symhd), V1118)

									gen14355 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.rename"), gen14354)

									var gen14356 Obj
									if True == gen14355 {
										gen14351 := Call(__e, ShenFunc(symtl), V1118)

										gen14352 := Call(__e, ShenFunc(symcons_2), gen14351)

										var gen14353 Obj
										if True == gen14352 {
											gen14347 := Call(__e, ShenFunc(symtl), V1118)

											gen14348 := Call(__e, ShenFunc(symhd), gen14347)

											gen14349 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen14348)

											var gen14350 Obj
											if True == gen14349 {
												gen14343 := Call(__e, ShenFunc(symtl), V1118)

												gen14344 := Call(__e, ShenFunc(symtl), gen14343)

												gen14345 := Call(__e, ShenFunc(symcons_2), gen14344)

												var gen14346 Obj
												if True == gen14345 {
													gen14338 := Call(__e, ShenFunc(symtl), V1118)

													gen14339 := Call(__e, ShenFunc(symtl), gen14338)

													gen14340 := Call(__e, ShenFunc(symhd), gen14339)

													gen14341 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variables"), gen14340)

													var gen14342 Obj
													if True == gen14341 {
														gen14333 := Call(__e, ShenFunc(symtl), V1118)

														gen14334 := Call(__e, ShenFunc(symtl), gen14333)

														gen14335 := Call(__e, ShenFunc(symtl), gen14334)

														gen14336 := Call(__e, ShenFunc(symcons_2), gen14335)

														var gen14337 Obj
														if True == gen14336 {
															gen14327 := Call(__e, ShenFunc(symtl), V1118)

															gen14328 := Call(__e, ShenFunc(symtl), gen14327)

															gen14329 := Call(__e, ShenFunc(symtl), gen14328)

															gen14330 := Call(__e, ShenFunc(symhd), gen14329)

															gen14331 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen14330)

															var gen14332 Obj
															if True == gen14331 {
																gen14321 := Call(__e, ShenFunc(symtl), V1118)

																gen14322 := Call(__e, ShenFunc(symtl), gen14321)

																gen14323 := Call(__e, ShenFunc(symtl), gen14322)

																gen14324 := Call(__e, ShenFunc(symtl), gen14323)

																gen14325 := Call(__e, ShenFunc(symcons_2), gen14324)

																var gen14326 Obj
																if True == gen14325 {
																	gen14314 := Call(__e, ShenFunc(symtl), V1118)

																	gen14315 := Call(__e, ShenFunc(symtl), gen14314)

																	gen14316 := Call(__e, ShenFunc(symtl), gen14315)

																	gen14317 := Call(__e, ShenFunc(symtl), gen14316)

																	gen14318 := Call(__e, ShenFunc(symhd), gen14317)

																	gen14319 := Call(__e, ShenFunc(sym_a), Nil, gen14318)

																	var gen14320 Obj
																	if True == gen14319 {
																		gen14307 := Call(__e, ShenFunc(symtl), V1118)

																		gen14308 := Call(__e, ShenFunc(symtl), gen14307)

																		gen14309 := Call(__e, ShenFunc(symtl), gen14308)

																		gen14310 := Call(__e, ShenFunc(symtl), gen14309)

																		gen14311 := Call(__e, ShenFunc(symtl), gen14310)

																		gen14312 := Call(__e, ShenFunc(symcons_2), gen14311)

																		var gen14313 Obj
																		if True == gen14312 {
																			gen14299 := Call(__e, ShenFunc(symtl), V1118)

																			gen14300 := Call(__e, ShenFunc(symtl), gen14299)

																			gen14301 := Call(__e, ShenFunc(symtl), gen14300)

																			gen14302 := Call(__e, ShenFunc(symtl), gen14301)

																			gen14303 := Call(__e, ShenFunc(symtl), gen14302)

																			gen14304 := Call(__e, ShenFunc(symhd), gen14303)

																			gen14305 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen14304)

																			var gen14306 Obj
																			if True == gen14305 {
																				gen14291 := Call(__e, ShenFunc(symtl), V1118)

																				gen14292 := Call(__e, ShenFunc(symtl), gen14291)

																				gen14293 := Call(__e, ShenFunc(symtl), gen14292)

																				gen14294 := Call(__e, ShenFunc(symtl), gen14293)

																				gen14295 := Call(__e, ShenFunc(symtl), gen14294)

																				gen14296 := Call(__e, ShenFunc(symtl), gen14295)

																				gen14297 := Call(__e, ShenFunc(symcons_2), gen14296)

																				var gen14298 Obj
																				if True == gen14297 {
																					gen14282 := Call(__e, ShenFunc(symtl), V1118)

																					gen14283 := Call(__e, ShenFunc(symtl), gen14282)

																					gen14284 := Call(__e, ShenFunc(symtl), gen14283)

																					gen14285 := Call(__e, ShenFunc(symtl), gen14284)

																					gen14286 := Call(__e, ShenFunc(symtl), gen14285)

																					gen14287 := Call(__e, ShenFunc(symtl), gen14286)

																					gen14288 := Call(__e, ShenFunc(symhd), gen14287)

																					gen14289 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen14288)

																					var gen14290 Obj
																					if True == gen14289 {
																						gen14273 := Call(__e, ShenFunc(symtl), V1118)

																						gen14274 := Call(__e, ShenFunc(symtl), gen14273)

																						gen14275 := Call(__e, ShenFunc(symtl), gen14274)

																						gen14276 := Call(__e, ShenFunc(symtl), gen14275)

																						gen14277 := Call(__e, ShenFunc(symtl), gen14276)

																						gen14278 := Call(__e, ShenFunc(symtl), gen14277)

																						gen14279 := Call(__e, ShenFunc(symtl), gen14278)

																						gen14280 := Call(__e, ShenFunc(symcons_2), gen14279)

																						var gen14281 Obj
																						if True == gen14280 {
																							gen14264 := Call(__e, ShenFunc(symtl), V1118)

																							gen14265 := Call(__e, ShenFunc(symtl), gen14264)

																							gen14266 := Call(__e, ShenFunc(symtl), gen14265)

																							gen14267 := Call(__e, ShenFunc(symtl), gen14266)

																							gen14268 := Call(__e, ShenFunc(symtl), gen14267)

																							gen14269 := Call(__e, ShenFunc(symtl), gen14268)

																							gen14270 := Call(__e, ShenFunc(symtl), gen14269)

																							gen14271 := Call(__e, ShenFunc(symtl), gen14270)

																							gen14272 := Call(__e, ShenFunc(sym_a), Nil, gen14271)

																							if True == gen14272 {
																								gen14281 = True
																							} else {
																								gen14281 = False
																							}

																						} else {
																							gen14281 = False
																						}
																						if True == gen14281 {
																							gen14290 = True
																						} else {
																							gen14290 = False
																						}

																					} else {
																						gen14290 = False
																					}
																					if True == gen14290 {
																						gen14298 = True
																					} else {
																						gen14298 = False
																					}

																				} else {
																					gen14298 = False
																				}
																				if True == gen14298 {
																					gen14306 = True
																				} else {
																					gen14306 = False
																				}

																			} else {
																				gen14306 = False
																			}
																			if True == gen14306 {
																				gen14313 = True
																			} else {
																				gen14313 = False
																			}

																		} else {
																			gen14313 = False
																		}
																		if True == gen14313 {
																			gen14320 = True
																		} else {
																			gen14320 = False
																		}

																	} else {
																		gen14320 = False
																	}
																	if True == gen14320 {
																		gen14326 = True
																	} else {
																		gen14326 = False
																	}

																} else {
																	gen14326 = False
																}
																if True == gen14326 {
																	gen14332 = True
																} else {
																	gen14332 = False
																}

															} else {
																gen14332 = False
															}
															if True == gen14332 {
																gen14337 = True
															} else {
																gen14337 = False
															}

														} else {
															gen14337 = False
														}
														if True == gen14337 {
															gen14342 = True
														} else {
															gen14342 = False
														}

													} else {
														gen14342 = False
													}
													if True == gen14342 {
														gen14346 = True
													} else {
														gen14346 = False
													}

												} else {
													gen14346 = False
												}
												if True == gen14346 {
													gen14350 = True
												} else {
													gen14350 = False
												}

											} else {
												gen14350 = False
											}
											if True == gen14350 {
												gen14353 = True
											} else {
												gen14353 = False
											}

										} else {
											gen14353 = False
										}
										if True == gen14353 {
											gen14356 = True
										} else {
											gen14356 = False
										}

									} else {
										gen14356 = False
									}
									if True == gen14356 {
										gen14358 = True
									} else {
										gen14358 = False
									}

								} else {
									gen14358 = False
								}
								if True == gen14358 {
									gen14256 := Call(__e, ShenFunc(symtl), V1118)

									gen14257 := Call(__e, ShenFunc(symtl), gen14256)

									gen14258 := Call(__e, ShenFunc(symtl), gen14257)

									gen14259 := Call(__e, ShenFunc(symtl), gen14258)

									gen14260 := Call(__e, ShenFunc(symtl), gen14259)

									gen14261 := Call(__e, ShenFunc(symtl), gen14260)

									gen14262 := Call(__e, ShenFunc(symtl), gen14261)

									gen14263 := Call(__e, ShenFunc(symhd), gen14262)

									__e.TailApply(ShenFunc(symshen_4aum__to__shen), gen14263)

									return

								} else {
									gen14254 := Call(__e, ShenFunc(symcons_2), V1118)

									var gen14255 Obj
									if True == gen14254 {
										gen14251 := Call(__e, ShenFunc(symhd), V1118)

										gen14252 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.rename"), gen14251)

										var gen14253 Obj
										if True == gen14252 {
											gen14248 := Call(__e, ShenFunc(symtl), V1118)

											gen14249 := Call(__e, ShenFunc(symcons_2), gen14248)

											var gen14250 Obj
											if True == gen14249 {
												gen14244 := Call(__e, ShenFunc(symtl), V1118)

												gen14245 := Call(__e, ShenFunc(symhd), gen14244)

												gen14246 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen14245)

												var gen14247 Obj
												if True == gen14246 {
													gen14240 := Call(__e, ShenFunc(symtl), V1118)

													gen14241 := Call(__e, ShenFunc(symtl), gen14240)

													gen14242 := Call(__e, ShenFunc(symcons_2), gen14241)

													var gen14243 Obj
													if True == gen14242 {
														gen14235 := Call(__e, ShenFunc(symtl), V1118)

														gen14236 := Call(__e, ShenFunc(symtl), gen14235)

														gen14237 := Call(__e, ShenFunc(symhd), gen14236)

														gen14238 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.variables"), gen14237)

														var gen14239 Obj
														if True == gen14238 {
															gen14230 := Call(__e, ShenFunc(symtl), V1118)

															gen14231 := Call(__e, ShenFunc(symtl), gen14230)

															gen14232 := Call(__e, ShenFunc(symtl), gen14231)

															gen14233 := Call(__e, ShenFunc(symcons_2), gen14232)

															var gen14234 Obj
															if True == gen14233 {
																gen14224 := Call(__e, ShenFunc(symtl), V1118)

																gen14225 := Call(__e, ShenFunc(symtl), gen14224)

																gen14226 := Call(__e, ShenFunc(symtl), gen14225)

																gen14227 := Call(__e, ShenFunc(symhd), gen14226)

																gen14228 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen14227)

																var gen14229 Obj
																if True == gen14228 {
																	gen14218 := Call(__e, ShenFunc(symtl), V1118)

																	gen14219 := Call(__e, ShenFunc(symtl), gen14218)

																	gen14220 := Call(__e, ShenFunc(symtl), gen14219)

																	gen14221 := Call(__e, ShenFunc(symtl), gen14220)

																	gen14222 := Call(__e, ShenFunc(symcons_2), gen14221)

																	var gen14223 Obj
																	if True == gen14222 {
																		gen14211 := Call(__e, ShenFunc(symtl), V1118)

																		gen14212 := Call(__e, ShenFunc(symtl), gen14211)

																		gen14213 := Call(__e, ShenFunc(symtl), gen14212)

																		gen14214 := Call(__e, ShenFunc(symtl), gen14213)

																		gen14215 := Call(__e, ShenFunc(symhd), gen14214)

																		gen14216 := Call(__e, ShenFunc(symcons_2), gen14215)

																		var gen14217 Obj
																		if True == gen14216 {
																			gen14204 := Call(__e, ShenFunc(symtl), V1118)

																			gen14205 := Call(__e, ShenFunc(symtl), gen14204)

																			gen14206 := Call(__e, ShenFunc(symtl), gen14205)

																			gen14207 := Call(__e, ShenFunc(symtl), gen14206)

																			gen14208 := Call(__e, ShenFunc(symtl), gen14207)

																			gen14209 := Call(__e, ShenFunc(symcons_2), gen14208)

																			var gen14210 Obj
																			if True == gen14209 {
																				gen14196 := Call(__e, ShenFunc(symtl), V1118)

																				gen14197 := Call(__e, ShenFunc(symtl), gen14196)

																				gen14198 := Call(__e, ShenFunc(symtl), gen14197)

																				gen14199 := Call(__e, ShenFunc(symtl), gen14198)

																				gen14200 := Call(__e, ShenFunc(symtl), gen14199)

																				gen14201 := Call(__e, ShenFunc(symhd), gen14200)

																				gen14202 := Call(__e, ShenFunc(sym_a), MakeSymbol("and"), gen14201)

																				var gen14203 Obj
																				if True == gen14202 {
																					gen14188 := Call(__e, ShenFunc(symtl), V1118)

																					gen14189 := Call(__e, ShenFunc(symtl), gen14188)

																					gen14190 := Call(__e, ShenFunc(symtl), gen14189)

																					gen14191 := Call(__e, ShenFunc(symtl), gen14190)

																					gen14192 := Call(__e, ShenFunc(symtl), gen14191)

																					gen14193 := Call(__e, ShenFunc(symtl), gen14192)

																					gen14194 := Call(__e, ShenFunc(symcons_2), gen14193)

																					var gen14195 Obj
																					if True == gen14194 {
																						gen14179 := Call(__e, ShenFunc(symtl), V1118)

																						gen14180 := Call(__e, ShenFunc(symtl), gen14179)

																						gen14181 := Call(__e, ShenFunc(symtl), gen14180)

																						gen14182 := Call(__e, ShenFunc(symtl), gen14181)

																						gen14183 := Call(__e, ShenFunc(symtl), gen14182)

																						gen14184 := Call(__e, ShenFunc(symtl), gen14183)

																						gen14185 := Call(__e, ShenFunc(symhd), gen14184)

																						gen14186 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.then"), gen14185)

																						var gen14187 Obj
																						if True == gen14186 {
																							gen14170 := Call(__e, ShenFunc(symtl), V1118)

																							gen14171 := Call(__e, ShenFunc(symtl), gen14170)

																							gen14172 := Call(__e, ShenFunc(symtl), gen14171)

																							gen14173 := Call(__e, ShenFunc(symtl), gen14172)

																							gen14174 := Call(__e, ShenFunc(symtl), gen14173)

																							gen14175 := Call(__e, ShenFunc(symtl), gen14174)

																							gen14176 := Call(__e, ShenFunc(symtl), gen14175)

																							gen14177 := Call(__e, ShenFunc(symcons_2), gen14176)

																							var gen14178 Obj
																							if True == gen14177 {
																								gen14161 := Call(__e, ShenFunc(symtl), V1118)

																								gen14162 := Call(__e, ShenFunc(symtl), gen14161)

																								gen14163 := Call(__e, ShenFunc(symtl), gen14162)

																								gen14164 := Call(__e, ShenFunc(symtl), gen14163)

																								gen14165 := Call(__e, ShenFunc(symtl), gen14164)

																								gen14166 := Call(__e, ShenFunc(symtl), gen14165)

																								gen14167 := Call(__e, ShenFunc(symtl), gen14166)

																								gen14168 := Call(__e, ShenFunc(symtl), gen14167)

																								gen14169 := Call(__e, ShenFunc(sym_a), Nil, gen14168)

																								if True == gen14169 {
																									gen14178 = True
																								} else {
																									gen14178 = False
																								}

																							} else {
																								gen14178 = False
																							}
																							if True == gen14178 {
																								gen14187 = True
																							} else {
																								gen14187 = False
																							}

																						} else {
																							gen14187 = False
																						}
																						if True == gen14187 {
																							gen14195 = True
																						} else {
																							gen14195 = False
																						}

																					} else {
																						gen14195 = False
																					}
																					if True == gen14195 {
																						gen14203 = True
																					} else {
																						gen14203 = False
																					}

																				} else {
																					gen14203 = False
																				}
																				if True == gen14203 {
																					gen14210 = True
																				} else {
																					gen14210 = False
																				}

																			} else {
																				gen14210 = False
																			}
																			if True == gen14210 {
																				gen14217 = True
																			} else {
																				gen14217 = False
																			}

																		} else {
																			gen14217 = False
																		}
																		if True == gen14217 {
																			gen14223 = True
																		} else {
																			gen14223 = False
																		}

																	} else {
																		gen14223 = False
																	}
																	if True == gen14223 {
																		gen14229 = True
																	} else {
																		gen14229 = False
																	}

																} else {
																	gen14229 = False
																}
																if True == gen14229 {
																	gen14234 = True
																} else {
																	gen14234 = False
																}

															} else {
																gen14234 = False
															}
															if True == gen14234 {
																gen14239 = True
															} else {
																gen14239 = False
															}

														} else {
															gen14239 = False
														}
														if True == gen14239 {
															gen14243 = True
														} else {
															gen14243 = False
														}

													} else {
														gen14243 = False
													}
													if True == gen14243 {
														gen14247 = True
													} else {
														gen14247 = False
													}

												} else {
													gen14247 = False
												}
												if True == gen14247 {
													gen14250 = True
												} else {
													gen14250 = False
												}

											} else {
												gen14250 = False
											}
											if True == gen14250 {
												gen14253 = True
											} else {
												gen14253 = False
											}

										} else {
											gen14253 = False
										}
										if True == gen14253 {
											gen14255 = True
										} else {
											gen14255 = False
										}

									} else {
										gen14255 = False
									}
									if True == gen14255 {
										gen14133 := Call(__e, ShenFunc(symtl), V1118)

										gen14134 := Call(__e, ShenFunc(symtl), gen14133)

										gen14135 := Call(__e, ShenFunc(symtl), gen14134)

										gen14136 := Call(__e, ShenFunc(symtl), gen14135)

										gen14137 := Call(__e, ShenFunc(symhd), gen14136)

										gen14138 := Call(__e, ShenFunc(symhd), gen14137)

										gen14139 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

										gen14140 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.newpv"), gen14139)

										gen14141 := Call(__e, ShenFunc(symtl), V1118)

										gen14142 := Call(__e, ShenFunc(symtl), gen14141)

										gen14143 := Call(__e, ShenFunc(symtl), gen14142)

										gen14144 := Call(__e, ShenFunc(symtl), gen14143)

										gen14145 := Call(__e, ShenFunc(symhd), gen14144)

										gen14146 := Call(__e, ShenFunc(symtl), gen14145)

										gen14147 := Call(__e, ShenFunc(symtl), V1118)

										gen14148 := Call(__e, ShenFunc(symtl), gen14147)

										gen14149 := Call(__e, ShenFunc(symtl), gen14148)

										gen14150 := Call(__e, ShenFunc(symtl), gen14149)

										gen14151 := Call(__e, ShenFunc(symtl), gen14150)

										gen14152 := Call(__e, ShenFunc(symcons), gen14146, gen14151)

										gen14153 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen14152)

										gen14154 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.variables"), gen14153)

										gen14155 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.the"), gen14154)

										gen14156 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.rename"), gen14155)

										gen14157 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14156)

										gen14158 := Call(__e, ShenFunc(symcons), gen14157, Nil)

										gen14159 := Call(__e, ShenFunc(symcons), gen14140, gen14158)

										gen14160 := Call(__e, ShenFunc(symcons), gen14138, gen14159)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen14160)

										return

									} else {
										gen14131 := Call(__e, ShenFunc(symcons_2), V1118)

										var gen14132 Obj
										if True == gen14131 {
											gen14128 := Call(__e, ShenFunc(symhd), V1118)

											gen14129 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen14128)

											var gen14130 Obj
											if True == gen14129 {
												gen14125 := Call(__e, ShenFunc(symtl), V1118)

												gen14126 := Call(__e, ShenFunc(symcons_2), gen14125)

												var gen14127 Obj
												if True == gen14126 {
													gen14121 := Call(__e, ShenFunc(symtl), V1118)

													gen14122 := Call(__e, ShenFunc(symtl), gen14121)

													gen14123 := Call(__e, ShenFunc(symcons_2), gen14122)

													var gen14124 Obj
													if True == gen14123 {
														gen14116 := Call(__e, ShenFunc(symtl), V1118)

														gen14117 := Call(__e, ShenFunc(symtl), gen14116)

														gen14118 := Call(__e, ShenFunc(symhd), gen14117)

														gen14119 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.to"), gen14118)

														var gen14120 Obj
														if True == gen14119 {
															gen14111 := Call(__e, ShenFunc(symtl), V1118)

															gen14112 := Call(__e, ShenFunc(symtl), gen14111)

															gen14113 := Call(__e, ShenFunc(symtl), gen14112)

															gen14114 := Call(__e, ShenFunc(symcons_2), gen14113)

															var gen14115 Obj
															if True == gen14114 {
																gen14105 := Call(__e, ShenFunc(symtl), V1118)

																gen14106 := Call(__e, ShenFunc(symtl), gen14105)

																gen14107 := Call(__e, ShenFunc(symtl), gen14106)

																gen14108 := Call(__e, ShenFunc(symtl), gen14107)

																gen14109 := Call(__e, ShenFunc(symcons_2), gen14108)

																var gen14110 Obj
																if True == gen14109 {
																	gen14098 := Call(__e, ShenFunc(symtl), V1118)

																	gen14099 := Call(__e, ShenFunc(symtl), gen14098)

																	gen14100 := Call(__e, ShenFunc(symtl), gen14099)

																	gen14101 := Call(__e, ShenFunc(symtl), gen14100)

																	gen14102 := Call(__e, ShenFunc(symhd), gen14101)

																	gen14103 := Call(__e, ShenFunc(sym_a), MakeSymbol("in"), gen14102)

																	var gen14104 Obj
																	if True == gen14103 {
																		gen14091 := Call(__e, ShenFunc(symtl), V1118)

																		gen14092 := Call(__e, ShenFunc(symtl), gen14091)

																		gen14093 := Call(__e, ShenFunc(symtl), gen14092)

																		gen14094 := Call(__e, ShenFunc(symtl), gen14093)

																		gen14095 := Call(__e, ShenFunc(symtl), gen14094)

																		gen14096 := Call(__e, ShenFunc(symcons_2), gen14095)

																		var gen14097 Obj
																		if True == gen14096 {
																			gen14084 := Call(__e, ShenFunc(symtl), V1118)

																			gen14085 := Call(__e, ShenFunc(symtl), gen14084)

																			gen14086 := Call(__e, ShenFunc(symtl), gen14085)

																			gen14087 := Call(__e, ShenFunc(symtl), gen14086)

																			gen14088 := Call(__e, ShenFunc(symtl), gen14087)

																			gen14089 := Call(__e, ShenFunc(symtl), gen14088)

																			gen14090 := Call(__e, ShenFunc(sym_a), Nil, gen14089)

																			if True == gen14090 {
																				gen14097 = True
																			} else {
																				gen14097 = False
																			}

																		} else {
																			gen14097 = False
																		}
																		if True == gen14097 {
																			gen14104 = True
																		} else {
																			gen14104 = False
																		}

																	} else {
																		gen14104 = False
																	}
																	if True == gen14104 {
																		gen14110 = True
																	} else {
																		gen14110 = False
																	}

																} else {
																	gen14110 = False
																}
																if True == gen14110 {
																	gen14115 = True
																} else {
																	gen14115 = False
																}

															} else {
																gen14115 = False
															}
															if True == gen14115 {
																gen14120 = True
															} else {
																gen14120 = False
															}

														} else {
															gen14120 = False
														}
														if True == gen14120 {
															gen14124 = True
														} else {
															gen14124 = False
														}

													} else {
														gen14124 = False
													}
													if True == gen14124 {
														gen14127 = True
													} else {
														gen14127 = False
													}

												} else {
													gen14127 = False
												}
												if True == gen14127 {
													gen14130 = True
												} else {
													gen14130 = False
												}

											} else {
												gen14130 = False
											}
											if True == gen14130 {
												gen14132 = True
											} else {
												gen14132 = False
											}

										} else {
											gen14132 = False
										}
										if True == gen14132 {
											gen14052 := Call(__e, ShenFunc(symtl), V1118)

											gen14053 := Call(__e, ShenFunc(symhd), gen14052)

											gen14054 := Call(__e, ShenFunc(symtl), V1118)

											gen14055 := Call(__e, ShenFunc(symtl), gen14054)

											gen14056 := Call(__e, ShenFunc(symtl), gen14055)

											gen14057 := Call(__e, ShenFunc(symhd), gen14056)

											gen14058 := Call(__e, ShenFunc(symshen_4chwild), gen14057)

											gen14059 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

											gen14060 := Call(__e, ShenFunc(symcons), gen14058, gen14059)

											gen14061 := Call(__e, ShenFunc(symcons), gen14053, gen14060)

											gen14062 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.bindv"), gen14061)

											gen14063 := Call(__e, ShenFunc(symtl), V1118)

											gen14064 := Call(__e, ShenFunc(symtl), gen14063)

											gen14065 := Call(__e, ShenFunc(symtl), gen14064)

											gen14066 := Call(__e, ShenFunc(symtl), gen14065)

											gen14067 := Call(__e, ShenFunc(symtl), gen14066)

											gen14068 := Call(__e, ShenFunc(symhd), gen14067)

											gen14069 := Call(__e, ShenFunc(symshen_4aum__to__shen), gen14068)

											gen14070 := Call(__e, ShenFunc(symtl), V1118)

											gen14071 := Call(__e, ShenFunc(symhd), gen14070)

											gen14072 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

											gen14073 := Call(__e, ShenFunc(symcons), gen14071, gen14072)

											gen14074 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.unbindv"), gen14073)

											gen14075 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

											gen14076 := Call(__e, ShenFunc(symcons), gen14074, gen14075)

											gen14077 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen14076)

											gen14078 := Call(__e, ShenFunc(symcons), gen14077, Nil)

											gen14079 := Call(__e, ShenFunc(symcons), gen14069, gen14078)

											gen14080 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen14079)

											gen14081 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen14080)

											gen14082 := Call(__e, ShenFunc(symcons), gen14081, Nil)

											gen14083 := Call(__e, ShenFunc(symcons), gen14062, gen14082)

											__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen14083)

											return

										} else {
											gen14050 := Call(__e, ShenFunc(symcons_2), V1118)

											var gen14051 Obj
											if True == gen14050 {
												gen14047 := Call(__e, ShenFunc(symtl), V1118)

												gen14048 := Call(__e, ShenFunc(symcons_2), gen14047)

												var gen14049 Obj
												if True == gen14048 {
													gen14043 := Call(__e, ShenFunc(symtl), V1118)

													gen14044 := Call(__e, ShenFunc(symhd), gen14043)

													gen14045 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen14044)

													var gen14046 Obj
													if True == gen14045 {
														gen14039 := Call(__e, ShenFunc(symtl), V1118)

														gen14040 := Call(__e, ShenFunc(symtl), gen14039)

														gen14041 := Call(__e, ShenFunc(symcons_2), gen14040)

														var gen14042 Obj
														if True == gen14041 {
															gen14034 := Call(__e, ShenFunc(symtl), V1118)

															gen14035 := Call(__e, ShenFunc(symtl), gen14034)

															gen14036 := Call(__e, ShenFunc(symhd), gen14035)

															gen14037 := Call(__e, ShenFunc(sym_a), MakeSymbol("identical"), gen14036)

															var gen14038 Obj
															if True == gen14037 {
																gen14029 := Call(__e, ShenFunc(symtl), V1118)

																gen14030 := Call(__e, ShenFunc(symtl), gen14029)

																gen14031 := Call(__e, ShenFunc(symtl), gen14030)

																gen14032 := Call(__e, ShenFunc(symcons_2), gen14031)

																var gen14033 Obj
																if True == gen14032 {
																	gen14023 := Call(__e, ShenFunc(symtl), V1118)

																	gen14024 := Call(__e, ShenFunc(symtl), gen14023)

																	gen14025 := Call(__e, ShenFunc(symtl), gen14024)

																	gen14026 := Call(__e, ShenFunc(symhd), gen14025)

																	gen14027 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.to"), gen14026)

																	var gen14028 Obj
																	if True == gen14027 {
																		gen14017 := Call(__e, ShenFunc(symtl), V1118)

																		gen14018 := Call(__e, ShenFunc(symtl), gen14017)

																		gen14019 := Call(__e, ShenFunc(symtl), gen14018)

																		gen14020 := Call(__e, ShenFunc(symtl), gen14019)

																		gen14021 := Call(__e, ShenFunc(symcons_2), gen14020)

																		var gen14022 Obj
																		if True == gen14021 {
																			gen14011 := Call(__e, ShenFunc(symtl), V1118)

																			gen14012 := Call(__e, ShenFunc(symtl), gen14011)

																			gen14013 := Call(__e, ShenFunc(symtl), gen14012)

																			gen14014 := Call(__e, ShenFunc(symtl), gen14013)

																			gen14015 := Call(__e, ShenFunc(symtl), gen14014)

																			gen14016 := Call(__e, ShenFunc(sym_a), Nil, gen14015)

																			if True == gen14016 {
																				gen14022 = True
																			} else {
																				gen14022 = False
																			}

																		} else {
																			gen14022 = False
																		}
																		if True == gen14022 {
																			gen14028 = True
																		} else {
																			gen14028 = False
																		}

																	} else {
																		gen14028 = False
																	}
																	if True == gen14028 {
																		gen14033 = True
																	} else {
																		gen14033 = False
																	}

																} else {
																	gen14033 = False
																}
																if True == gen14033 {
																	gen14038 = True
																} else {
																	gen14038 = False
																}

															} else {
																gen14038 = False
															}
															if True == gen14038 {
																gen14042 = True
															} else {
																gen14042 = False
															}

														} else {
															gen14042 = False
														}
														if True == gen14042 {
															gen14046 = True
														} else {
															gen14046 = False
														}

													} else {
														gen14046 = False
													}
													if True == gen14046 {
														gen14049 = True
													} else {
														gen14049 = False
													}

												} else {
													gen14049 = False
												}
												if True == gen14049 {
													gen14051 = True
												} else {
													gen14051 = False
												}

											} else {
												gen14051 = False
											}
											if True == gen14051 {
												gen14003 := Call(__e, ShenFunc(symtl), V1118)

												gen14004 := Call(__e, ShenFunc(symtl), gen14003)

												gen14005 := Call(__e, ShenFunc(symtl), gen14004)

												gen14006 := Call(__e, ShenFunc(symtl), gen14005)

												gen14007 := Call(__e, ShenFunc(symhd), gen14006)

												gen14008 := Call(__e, ShenFunc(symhd), V1118)

												gen14009 := Call(__e, ShenFunc(symcons), gen14008, Nil)

												gen14010 := Call(__e, ShenFunc(symcons), gen14007, gen14009)

												__e.TailApply(ShenFunc(symcons), MakeSymbol("="), gen14010)

												return

											} else {
												gen14002 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.failed!"), V1118)

												if True == gen14002 {
													__e.Return(False)
													return
												} else {
													gen14000 := Call(__e, ShenFunc(symcons_2), V1118)

													var gen14001 Obj
													if True == gen14000 {
														gen13997 := Call(__e, ShenFunc(symhd), V1118)

														gen13998 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen13997)

														var gen13999 Obj
														if True == gen13998 {
															gen13994 := Call(__e, ShenFunc(symtl), V1118)

															gen13995 := Call(__e, ShenFunc(symcons_2), gen13994)

															var gen13996 Obj
															if True == gen13995 {
																gen13990 := Call(__e, ShenFunc(symtl), V1118)

																gen13991 := Call(__e, ShenFunc(symhd), gen13990)

																gen13992 := Call(__e, ShenFunc(sym_a), MakeSymbol("head"), gen13991)

																var gen13993 Obj
																if True == gen13992 {
																	gen13986 := Call(__e, ShenFunc(symtl), V1118)

																	gen13987 := Call(__e, ShenFunc(symtl), gen13986)

																	gen13988 := Call(__e, ShenFunc(symcons_2), gen13987)

																	var gen13989 Obj
																	if True == gen13988 {
																		gen13981 := Call(__e, ShenFunc(symtl), V1118)

																		gen13982 := Call(__e, ShenFunc(symtl), gen13981)

																		gen13983 := Call(__e, ShenFunc(symhd), gen13982)

																		gen13984 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen13983)

																		var gen13985 Obj
																		if True == gen13984 {
																			gen13976 := Call(__e, ShenFunc(symtl), V1118)

																			gen13977 := Call(__e, ShenFunc(symtl), gen13976)

																			gen13978 := Call(__e, ShenFunc(symtl), gen13977)

																			gen13979 := Call(__e, ShenFunc(symcons_2), gen13978)

																			var gen13980 Obj
																			if True == gen13979 {
																				gen13971 := Call(__e, ShenFunc(symtl), V1118)

																				gen13972 := Call(__e, ShenFunc(symtl), gen13971)

																				gen13973 := Call(__e, ShenFunc(symtl), gen13972)

																				gen13974 := Call(__e, ShenFunc(symtl), gen13973)

																				gen13975 := Call(__e, ShenFunc(sym_a), Nil, gen13974)

																				if True == gen13975 {
																					gen13980 = True
																				} else {
																					gen13980 = False
																				}

																			} else {
																				gen13980 = False
																			}
																			if True == gen13980 {
																				gen13985 = True
																			} else {
																				gen13985 = False
																			}

																		} else {
																			gen13985 = False
																		}
																		if True == gen13985 {
																			gen13989 = True
																		} else {
																			gen13989 = False
																		}

																	} else {
																		gen13989 = False
																	}
																	if True == gen13989 {
																		gen13993 = True
																	} else {
																		gen13993 = False
																	}

																} else {
																	gen13993 = False
																}
																if True == gen13993 {
																	gen13996 = True
																} else {
																	gen13996 = False
																}

															} else {
																gen13996 = False
															}
															if True == gen13996 {
																gen13999 = True
															} else {
																gen13999 = False
															}

														} else {
															gen13999 = False
														}
														if True == gen13999 {
															gen14001 = True
														} else {
															gen14001 = False
														}

													} else {
														gen14001 = False
													}
													if True == gen14001 {
														gen13968 := Call(__e, ShenFunc(symtl), V1118)

														gen13969 := Call(__e, ShenFunc(symtl), gen13968)

														gen13970 := Call(__e, ShenFunc(symtl), gen13969)

														__e.TailApply(ShenFunc(symcons), MakeSymbol("hd"), gen13970)

														return

													} else {
														gen13966 := Call(__e, ShenFunc(symcons_2), V1118)

														var gen13967 Obj
														if True == gen13966 {
															gen13963 := Call(__e, ShenFunc(symhd), V1118)

															gen13964 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen13963)

															var gen13965 Obj
															if True == gen13964 {
																gen13960 := Call(__e, ShenFunc(symtl), V1118)

																gen13961 := Call(__e, ShenFunc(symcons_2), gen13960)

																var gen13962 Obj
																if True == gen13961 {
																	gen13956 := Call(__e, ShenFunc(symtl), V1118)

																	gen13957 := Call(__e, ShenFunc(symhd), gen13956)

																	gen13958 := Call(__e, ShenFunc(sym_a), MakeSymbol("tail"), gen13957)

																	var gen13959 Obj
																	if True == gen13958 {
																		gen13952 := Call(__e, ShenFunc(symtl), V1118)

																		gen13953 := Call(__e, ShenFunc(symtl), gen13952)

																		gen13954 := Call(__e, ShenFunc(symcons_2), gen13953)

																		var gen13955 Obj
																		if True == gen13954 {
																			gen13947 := Call(__e, ShenFunc(symtl), V1118)

																			gen13948 := Call(__e, ShenFunc(symtl), gen13947)

																			gen13949 := Call(__e, ShenFunc(symhd), gen13948)

																			gen13950 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.of"), gen13949)

																			var gen13951 Obj
																			if True == gen13950 {
																				gen13942 := Call(__e, ShenFunc(symtl), V1118)

																				gen13943 := Call(__e, ShenFunc(symtl), gen13942)

																				gen13944 := Call(__e, ShenFunc(symtl), gen13943)

																				gen13945 := Call(__e, ShenFunc(symcons_2), gen13944)

																				var gen13946 Obj
																				if True == gen13945 {
																					gen13937 := Call(__e, ShenFunc(symtl), V1118)

																					gen13938 := Call(__e, ShenFunc(symtl), gen13937)

																					gen13939 := Call(__e, ShenFunc(symtl), gen13938)

																					gen13940 := Call(__e, ShenFunc(symtl), gen13939)

																					gen13941 := Call(__e, ShenFunc(sym_a), Nil, gen13940)

																					if True == gen13941 {
																						gen13946 = True
																					} else {
																						gen13946 = False
																					}

																				} else {
																					gen13946 = False
																				}
																				if True == gen13946 {
																					gen13951 = True
																				} else {
																					gen13951 = False
																				}

																			} else {
																				gen13951 = False
																			}
																			if True == gen13951 {
																				gen13955 = True
																			} else {
																				gen13955 = False
																			}

																		} else {
																			gen13955 = False
																		}
																		if True == gen13955 {
																			gen13959 = True
																		} else {
																			gen13959 = False
																		}

																	} else {
																		gen13959 = False
																	}
																	if True == gen13959 {
																		gen13962 = True
																	} else {
																		gen13962 = False
																	}

																} else {
																	gen13962 = False
																}
																if True == gen13962 {
																	gen13965 = True
																} else {
																	gen13965 = False
																}

															} else {
																gen13965 = False
															}
															if True == gen13965 {
																gen13967 = True
															} else {
																gen13967 = False
															}

														} else {
															gen13967 = False
														}
														if True == gen13967 {
															gen13934 := Call(__e, ShenFunc(symtl), V1118)

															gen13935 := Call(__e, ShenFunc(symtl), gen13934)

															gen13936 := Call(__e, ShenFunc(symtl), gen13935)

															__e.TailApply(ShenFunc(symcons), MakeSymbol("tl"), gen13936)

															return

														} else {
															gen13932 := Call(__e, ShenFunc(symcons_2), V1118)

															var gen13933 Obj
															if True == gen13932 {
																gen13929 := Call(__e, ShenFunc(symhd), V1118)

																gen13930 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.pop"), gen13929)

																var gen13931 Obj
																if True == gen13930 {
																	gen13926 := Call(__e, ShenFunc(symtl), V1118)

																	gen13927 := Call(__e, ShenFunc(symcons_2), gen13926)

																	var gen13928 Obj
																	if True == gen13927 {
																		gen13922 := Call(__e, ShenFunc(symtl), V1118)

																		gen13923 := Call(__e, ShenFunc(symhd), gen13922)

																		gen13924 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen13923)

																		var gen13925 Obj
																		if True == gen13924 {
																			gen13918 := Call(__e, ShenFunc(symtl), V1118)

																			gen13919 := Call(__e, ShenFunc(symtl), gen13918)

																			gen13920 := Call(__e, ShenFunc(symcons_2), gen13919)

																			var gen13921 Obj
																			if True == gen13920 {
																				gen13913 := Call(__e, ShenFunc(symtl), V1118)

																				gen13914 := Call(__e, ShenFunc(symtl), gen13913)

																				gen13915 := Call(__e, ShenFunc(symhd), gen13914)

																				gen13916 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.stack"), gen13915)

																				var gen13917 Obj
																				if True == gen13916 {
																					gen13909 := Call(__e, ShenFunc(symtl), V1118)

																					gen13910 := Call(__e, ShenFunc(symtl), gen13909)

																					gen13911 := Call(__e, ShenFunc(symtl), gen13910)

																					gen13912 := Call(__e, ShenFunc(sym_a), Nil, gen13911)

																					if True == gen13912 {
																						gen13917 = True
																					} else {
																						gen13917 = False
																					}

																				} else {
																					gen13917 = False
																				}
																				if True == gen13917 {
																					gen13921 = True
																				} else {
																					gen13921 = False
																				}

																			} else {
																				gen13921 = False
																			}
																			if True == gen13921 {
																				gen13925 = True
																			} else {
																				gen13925 = False
																			}

																		} else {
																			gen13925 = False
																		}
																		if True == gen13925 {
																			gen13928 = True
																		} else {
																			gen13928 = False
																		}

																	} else {
																		gen13928 = False
																	}
																	if True == gen13928 {
																		gen13931 = True
																	} else {
																		gen13931 = False
																	}

																} else {
																	gen13931 = False
																}
																if True == gen13931 {
																	gen13933 = True
																} else {
																	gen13933 = False
																}

															} else {
																gen13933 = False
															}
															if True == gen13933 {
																gen13904 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.incinfs"), Nil)

																gen13905 := Call(__e, ShenFunc(symcons), MakeSymbol("Continuation"), Nil)

																gen13906 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen13905)

																gen13907 := Call(__e, ShenFunc(symcons), gen13906, Nil)

																gen13908 := Call(__e, ShenFunc(symcons), gen13904, gen13907)

																__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen13908)

																return

															} else {
																gen13902 := Call(__e, ShenFunc(symcons_2), V1118)

																var gen13903 Obj
																if True == gen13902 {
																	gen13899 := Call(__e, ShenFunc(symhd), V1118)

																	gen13900 := Call(__e, ShenFunc(sym_a), MakeSymbol("call"), gen13899)

																	var gen13901 Obj
																	if True == gen13900 {
																		gen13896 := Call(__e, ShenFunc(symtl), V1118)

																		gen13897 := Call(__e, ShenFunc(symcons_2), gen13896)

																		var gen13898 Obj
																		if True == gen13897 {
																			gen13892 := Call(__e, ShenFunc(symtl), V1118)

																			gen13893 := Call(__e, ShenFunc(symhd), gen13892)

																			gen13894 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.the"), gen13893)

																			var gen13895 Obj
																			if True == gen13894 {
																				gen13888 := Call(__e, ShenFunc(symtl), V1118)

																				gen13889 := Call(__e, ShenFunc(symtl), gen13888)

																				gen13890 := Call(__e, ShenFunc(symcons_2), gen13889)

																				var gen13891 Obj
																				if True == gen13890 {
																					gen13883 := Call(__e, ShenFunc(symtl), V1118)

																					gen13884 := Call(__e, ShenFunc(symtl), gen13883)

																					gen13885 := Call(__e, ShenFunc(symhd), gen13884)

																					gen13886 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.continuation"), gen13885)

																					var gen13887 Obj
																					if True == gen13886 {
																						gen13878 := Call(__e, ShenFunc(symtl), V1118)

																						gen13879 := Call(__e, ShenFunc(symtl), gen13878)

																						gen13880 := Call(__e, ShenFunc(symtl), gen13879)

																						gen13881 := Call(__e, ShenFunc(symcons_2), gen13880)

																						var gen13882 Obj
																						if True == gen13881 {
																							gen13873 := Call(__e, ShenFunc(symtl), V1118)

																							gen13874 := Call(__e, ShenFunc(symtl), gen13873)

																							gen13875 := Call(__e, ShenFunc(symtl), gen13874)

																							gen13876 := Call(__e, ShenFunc(symtl), gen13875)

																							gen13877 := Call(__e, ShenFunc(sym_a), Nil, gen13876)

																							if True == gen13877 {
																								gen13882 = True
																							} else {
																								gen13882 = False
																							}

																						} else {
																							gen13882 = False
																						}
																						if True == gen13882 {
																							gen13887 = True
																						} else {
																							gen13887 = False
																						}

																					} else {
																						gen13887 = False
																					}
																					if True == gen13887 {
																						gen13891 = True
																					} else {
																						gen13891 = False
																					}

																				} else {
																					gen13891 = False
																				}
																				if True == gen13891 {
																					gen13895 = True
																				} else {
																					gen13895 = False
																				}

																			} else {
																				gen13895 = False
																			}
																			if True == gen13895 {
																				gen13898 = True
																			} else {
																				gen13898 = False
																			}

																		} else {
																			gen13898 = False
																		}
																		if True == gen13898 {
																			gen13901 = True
																		} else {
																			gen13901 = False
																		}

																	} else {
																		gen13901 = False
																	}
																	if True == gen13901 {
																		gen13903 = True
																	} else {
																		gen13903 = False
																	}

																} else {
																	gen13903 = False
																}
																if True == gen13903 {
																	gen13864 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.incinfs"), Nil)

																	gen13865 := Call(__e, ShenFunc(symtl), V1118)

																	gen13866 := Call(__e, ShenFunc(symtl), gen13865)

																	gen13867 := Call(__e, ShenFunc(symtl), gen13866)

																	gen13868 := Call(__e, ShenFunc(symhd), gen13867)

																	gen13869 := Call(__e, ShenFunc(symshen_4chwild), gen13868)

																	gen13870 := Call(__e, ShenFunc(symshen_4call__the__continuation), gen13869, MakeSymbol("ProcessN"), MakeSymbol("Continuation"))

																	gen13871 := Call(__e, ShenFunc(symcons), gen13870, Nil)

																	gen13872 := Call(__e, ShenFunc(symcons), gen13864, gen13871)

																	__e.TailApply(ShenFunc(symcons), MakeSymbol("do"), gen13872)

																	return

																} else {
																	__e.Return(V1118)
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

							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aum_to_shen"), gen14630)

		gen14635 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1120 := __args[0]
			_ = V1120
			gen14634 := Call(__e, ShenFunc(sym_a), V1120, MakeSymbol("_"))

			if True == gen14634 {
				gen14633 := Call(__e, ShenFunc(symcons), MakeSymbol("ProcessN"), Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.newpv"), gen14633)

				return

			} else {
				gen14632 := Call(__e, ShenFunc(symcons_2), V1120)

				if True == gen14632 {
					gen14631 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Z := __args[0]
						_ = Z
						__e.TailApply(ShenFunc(symshen_4chwild), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen14631, V1120)

					return

				} else {
					__e.Return(V1120)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.chwild"), gen14635)

		gen14646 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1122 := __args[0]
			_ = V1122
			gen14636 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen14637 := Call(__e, ShenFunc(sym_5_1address), gen14636, V1122)

			gen14638 := Call(__e, ShenFunc(sym_7), gen14637, MakeNumber(1))

			Count_71 := gen14638
			gen14639 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen14640 := Call(__e, ShenFunc(symaddress_1_6), gen14639, V1122, Count_71)

			IncVar := gen14640
			_ = IncVar
			gen14641 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14642 := Call(__e, ShenFunc(sym_5_1address), gen14641, V1122)

			Vector := gen14642
			gen14643 := Call(__e, ShenFunc(symlimit), Vector)

			gen14644 := Call(__e, ShenFunc(sym_a), Count_71, gen14643)

			var gen14645 Obj
			if True == gen14644 {
				gen14645 = Call(__e, ShenFunc(symshen_4resizeprocessvector), V1122, Count_71)

			} else {
				gen14645 = MakeSymbol("shen.skip")
			}
			ResizeVectorIfNeeded := gen14645
			_ = ResizeVectorIfNeeded
			__e.TailApply(ShenFunc(symshen_4mk_1pvar), Count_71)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newpv"), gen14646)

		gen14652 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1125 := __args[0]
			_ = V1125
			V1126 := __args[1]
			_ = V1126
			gen14647 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14648 := Call(__e, ShenFunc(sym_5_1address), gen14647, V1125)

			Vector := gen14648
			gen14649 := Call(__e, ShenFunc(sym_7), V1126, V1126)

			gen14650 := Call(__e, ShenFunc(symshen_4resize_1vector), Vector, gen14649, MakeSymbol("shen.-null-"))

			BigVector := gen14650
			gen14651 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			__e.TailApply(ShenFunc(symaddress_1_6), gen14651, V1125, BigVector)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.resizeprocessvector"), gen14652)

		gen14657 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1130 := __args[0]
			_ = V1130
			V1131 := __args[1]
			_ = V1131
			V1132 := __args[2]
			_ = V1132
			gen14653 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1131)

			gen14654 := Call(__e, ShenFunc(symabsvector), gen14653)

			gen14655 := Call(__e, ShenFunc(symaddress_1_6), gen14654, MakeNumber(0), V1131)

			BigVector := gen14655
			gen14656 := Call(__e, ShenFunc(symlimit), V1130)

			__e.TailApply(ShenFunc(symshen_4copy_1vector), V1130, BigVector, gen14656, V1131, V1132)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.resize-vector"), gen14657)

		gen14662 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1138 := __args[0]
			_ = V1138
			V1139 := __args[1]
			_ = V1139
			V1140 := __args[2]
			_ = V1140
			V1141 := __args[3]
			_ = V1141
			V1142 := __args[4]
			_ = V1142
			gen14658 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1140)

			gen14659 := Call(__e, ShenFunc(sym_7), V1141, MakeNumber(1))

			gen14660 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1140)

			gen14661 := Call(__e, ShenFunc(symshen_4copy_1vector_1stage_11), MakeNumber(1), V1138, V1139, gen14660)

			__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_12), gen14658, gen14659, V1142, gen14661)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector"), gen14662)

		gen14667 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1150 := __args[0]
			_ = V1150
			V1151 := __args[1]
			_ = V1151
			V1152 := __args[2]
			_ = V1152
			V1153 := __args[3]
			_ = V1153
			gen14666 := Call(__e, ShenFunc(sym_a), V1153, V1150)

			if True == gen14666 {
				__e.Return(V1152)
				return
			} else {
				gen14663 := Call(__e, ShenFunc(sym_7), MakeNumber(1), V1150)

				gen14664 := Call(__e, ShenFunc(sym_5_1address), V1151, V1150)

				gen14665 := Call(__e, ShenFunc(symaddress_1_6), V1152, V1150, gen14664)

				__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_11), gen14663, V1151, gen14665, V1153)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector-stage-1"), gen14667)

		gen14671 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1161 := __args[0]
			_ = V1161
			V1162 := __args[1]
			_ = V1162
			V1163 := __args[2]
			_ = V1163
			V1164 := __args[3]
			_ = V1164
			gen14670 := Call(__e, ShenFunc(sym_a), V1162, V1161)

			if True == gen14670 {
				__e.Return(V1164)
				return
			} else {
				gen14668 := Call(__e, ShenFunc(sym_7), V1161, MakeNumber(1))

				gen14669 := Call(__e, ShenFunc(symaddress_1_6), V1164, V1161, V1163)

				__e.TailApply(ShenFunc(symshen_4copy_1vector_1stage_12), gen14668, V1162, V1163, gen14669)

				return

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.copy-vector-stage-2"), gen14671)

		gen14674 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1166 := __args[0]
			_ = V1166
			gen14672 := Call(__e, ShenFunc(symabsvector), MakeNumber(2))

			gen14673 := Call(__e, ShenFunc(symaddress_1_6), gen14672, MakeNumber(0), MakeSymbol("shen.pvar"))

			__e.TailApply(ShenFunc(symaddress_1_6), gen14673, MakeNumber(1), V1166)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.mk-pvar"), gen14674)

		gen14680 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1168 := __args[0]
			_ = V1168
			gen14679 := Call(__e, ShenFunc(symabsvector_2), V1168)

			if True == gen14679 {
				gen14675 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(MakeSymbol("shen.not-pvar"))
					return
				}, 1)
				gen14676 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(sym_5_1address), V1168, MakeNumber(0))

					return
				}, 0)
				gen14677 := Try(__e, gen14676).Catch(gen14675)
				gen14678 := Call(__e, ShenFunc(sym_a), gen14677, MakeSymbol("shen.pvar"))

				if True == gen14678 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pvar?"), gen14680)

		gen14684 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1172 := __args[0]
			_ = V1172
			V1173 := __args[1]
			_ = V1173
			V1174 := __args[2]
			_ = V1174
			gen14681 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14682 := Call(__e, ShenFunc(sym_5_1address), gen14681, V1174)

			Vector := gen14682
			gen14683 := Call(__e, ShenFunc(sym_5_1address), V1172, MakeNumber(1))

			__e.TailApply(ShenFunc(symaddress_1_6), Vector, gen14683, V1173)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.bindv"), gen14684)

		gen14688 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1177 := __args[0]
			_ = V1177
			V1178 := __args[1]
			_ = V1178
			gen14685 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14686 := Call(__e, ShenFunc(sym_5_1address), gen14685, V1178)

			Vector := gen14686
			gen14687 := Call(__e, ShenFunc(sym_5_1address), V1177, MakeNumber(1))

			__e.TailApply(ShenFunc(symaddress_1_6), Vector, gen14687, MakeSymbol("shen.-null-"))

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.unbindv"), gen14688)

		gen14691 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen14689 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			gen14690 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen14689)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*infs*"), gen14690)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.incinfs"), gen14691)

		gen14719 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1182 := __args[0]
			_ = V1182
			V1183 := __args[1]
			_ = V1183
			V1184 := __args[2]
			_ = V1184
			gen14717 := Call(__e, ShenFunc(symcons_2), V1182)

			var gen14718 Obj
			if True == gen14717 {
				gen14714 := Call(__e, ShenFunc(symhd), V1182)

				gen14715 := Call(__e, ShenFunc(symcons_2), gen14714)

				var gen14716 Obj
				if True == gen14715 {
					gen14712 := Call(__e, ShenFunc(symtl), V1182)

					gen14713 := Call(__e, ShenFunc(sym_a), Nil, gen14712)

					if True == gen14713 {
						gen14716 = True
					} else {
						gen14716 = False
					}

				} else {
					gen14716 = False
				}
				if True == gen14716 {
					gen14718 = True
				} else {
					gen14718 = False
				}

			} else {
				gen14718 = False
			}
			if True == gen14718 {
				gen14705 := Call(__e, ShenFunc(symhd), V1182)

				gen14706 := Call(__e, ShenFunc(symhd), gen14705)

				gen14707 := Call(__e, ShenFunc(symhd), V1182)

				gen14708 := Call(__e, ShenFunc(symtl), gen14707)

				gen14709 := Call(__e, ShenFunc(symcons), V1184, Nil)

				gen14710 := Call(__e, ShenFunc(symcons), V1183, gen14709)

				gen14711 := Call(__e, ShenFunc(symappend), gen14708, gen14710)

				__e.TailApply(ShenFunc(symcons), gen14706, gen14711)

				return

			} else {
				gen14703 := Call(__e, ShenFunc(symcons_2), V1182)

				var gen14704 Obj
				if True == gen14703 {
					gen14701 := Call(__e, ShenFunc(symhd), V1182)

					gen14702 := Call(__e, ShenFunc(symcons_2), gen14701)

					if True == gen14702 {
						gen14704 = True
					} else {
						gen14704 = False
					}

				} else {
					gen14704 = False
				}
				if True == gen14704 {
					gen14692 := Call(__e, ShenFunc(symtl), V1182)

					gen14693 := Call(__e, ShenFunc(symshen_4newcontinuation), gen14692, V1183, V1184)

					NewContinuation := gen14693
					gen14694 := Call(__e, ShenFunc(symhd), V1182)

					gen14695 := Call(__e, ShenFunc(symhd), gen14694)

					gen14696 := Call(__e, ShenFunc(symhd), V1182)

					gen14697 := Call(__e, ShenFunc(symtl), gen14696)

					gen14698 := Call(__e, ShenFunc(symcons), NewContinuation, Nil)

					gen14699 := Call(__e, ShenFunc(symcons), V1183, gen14698)

					gen14700 := Call(__e, ShenFunc(symappend), gen14697, gen14699)

					__e.TailApply(ShenFunc(symcons), gen14695, gen14700)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call_the_continuation"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call_the_continuation"), gen14719)

		gen14736 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1188 := __args[0]
			_ = V1188
			V1189 := __args[1]
			_ = V1189
			V1190 := __args[2]
			_ = V1190
			gen14735 := Call(__e, ShenFunc(sym_a), Nil, V1188)

			if True == gen14735 {
				__e.Return(V1190)
				return
			} else {
				gen14733 := Call(__e, ShenFunc(symcons_2), V1188)

				var gen14734 Obj
				if True == gen14733 {
					gen14731 := Call(__e, ShenFunc(symhd), V1188)

					gen14732 := Call(__e, ShenFunc(symcons_2), gen14731)

					if True == gen14732 {
						gen14734 = True
					} else {
						gen14734 = False
					}

				} else {
					gen14734 = False
				}
				if True == gen14734 {
					gen14720 := Call(__e, ShenFunc(symhd), V1188)

					gen14721 := Call(__e, ShenFunc(symhd), gen14720)

					gen14722 := Call(__e, ShenFunc(symhd), V1188)

					gen14723 := Call(__e, ShenFunc(symtl), gen14722)

					gen14724 := Call(__e, ShenFunc(symtl), V1188)

					gen14725 := Call(__e, ShenFunc(symshen_4newcontinuation), gen14724, V1189, V1190)

					gen14726 := Call(__e, ShenFunc(symcons), gen14725, Nil)

					gen14727 := Call(__e, ShenFunc(symcons), V1189, gen14726)

					gen14728 := Call(__e, ShenFunc(symappend), gen14723, gen14727)

					gen14729 := Call(__e, ShenFunc(symcons), gen14721, gen14728)

					gen14730 := Call(__e, ShenFunc(symcons), gen14729, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("freeze"), gen14730)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.newcontinuation"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newcontinuation"), gen14736)

		gen14737 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1198 := __args[0]
			_ = V1198
			V1199 := __args[1]
			_ = V1199
			V1200 := __args[2]
			_ = V1200
			__e.TailApply(ShenFunc(symshen_4deref), V1198, V1199)

			return
		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("return"), gen14737)

		gen14741 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1208 := __args[0]
			_ = V1208
			V1209 := __args[1]
			_ = V1209
			V1210 := __args[2]
			_ = V1210
			gen14738 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*infs*"))

			gen14739 := Call(__e, ShenFunc(symshen_4app), gen14738, MakeString(" inferences\n"), MakeSymbol("shen.a"))

			gen14740 := Call(__e, ShenFunc(symstoutput))

			Call(__e, ShenFunc(symshen_4prhush), gen14739, gen14740)

			__e.TailApply(ShenFunc(symshen_4deref), V1208, V1209)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.measure&return"), gen14741)

		gen14744 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1215 := __args[0]
			_ = V1215
			V1216 := __args[1]
			_ = V1216
			V1217 := __args[2]
			_ = V1217
			V1218 := __args[3]
			_ = V1218
			gen14742 := Call(__e, ShenFunc(symshen_4lazyderef), V1215, V1217)

			gen14743 := Call(__e, ShenFunc(symshen_4lazyderef), V1216, V1217)

			__e.TailApply(ShenFunc(symshen_4lzy_a), gen14742, gen14743, V1217, V1218)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unify"), gen14744)

		gen14760 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1240 := __args[0]
			_ = V1240
			V1241 := __args[1]
			_ = V1241
			V1242 := __args[2]
			_ = V1242
			V1243 := __args[3]
			_ = V1243
			gen14759 := Call(__e, ShenFunc(sym_a), V1241, V1240)

			if True == gen14759 {
				__e.TailApply(ShenFunc(symthaw), V1243)

				return
			} else {
				gen14758 := Call(__e, ShenFunc(symshen_4pvar_2), V1240)

				if True == gen14758 {
					__e.TailApply(ShenFunc(symbind), V1240, V1241, V1242, V1243)

					return
				} else {
					gen14757 := Call(__e, ShenFunc(symshen_4pvar_2), V1241)

					if True == gen14757 {
						__e.TailApply(ShenFunc(symbind), V1241, V1240, V1242, V1243)

						return
					} else {
						gen14755 := Call(__e, ShenFunc(symcons_2), V1240)

						var gen14756 Obj
						if True == gen14755 {
							gen14754 := Call(__e, ShenFunc(symcons_2), V1241)

							if True == gen14754 {
								gen14756 = True
							} else {
								gen14756 = False
							}

						} else {
							gen14756 = False
						}
						if True == gen14756 {
							gen14745 := Call(__e, ShenFunc(symhd), V1240)

							gen14746 := Call(__e, ShenFunc(symshen_4lazyderef), gen14745, V1242)

							gen14747 := Call(__e, ShenFunc(symhd), V1241)

							gen14748 := Call(__e, ShenFunc(symshen_4lazyderef), gen14747, V1242)

							gen14753 := MakeNative(func(__e Evaluator, __args ...Obj) {
								gen14749 := Call(__e, ShenFunc(symtl), V1240)

								gen14750 := Call(__e, ShenFunc(symshen_4lazyderef), gen14749, V1242)

								gen14751 := Call(__e, ShenFunc(symtl), V1241)

								gen14752 := Call(__e, ShenFunc(symshen_4lazyderef), gen14751, V1242)

								__e.TailApply(ShenFunc(symshen_4lzy_a), gen14750, gen14752, V1242, V1243)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4lzy_a), gen14746, gen14748, V1242, gen14753)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy="), gen14760)

		gen14769 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1246 := __args[0]
			_ = V1246
			V1247 := __args[1]
			_ = V1247
			gen14768 := Call(__e, ShenFunc(symcons_2), V1246)

			if True == gen14768 {
				gen14764 := Call(__e, ShenFunc(symhd), V1246)

				gen14765 := Call(__e, ShenFunc(symshen_4deref), gen14764, V1247)

				gen14766 := Call(__e, ShenFunc(symtl), V1246)

				gen14767 := Call(__e, ShenFunc(symshen_4deref), gen14766, V1247)

				__e.TailApply(ShenFunc(symcons), gen14765, gen14767)

				return

			} else {
				gen14763 := Call(__e, ShenFunc(symshen_4pvar_2), V1246)

				if True == gen14763 {
					gen14761 := Call(__e, ShenFunc(symshen_4valvector), V1246, V1247)

					Value := gen14761
					gen14762 := Call(__e, ShenFunc(sym_a), Value, MakeSymbol("shen.-null-"))

					if True == gen14762 {
						__e.Return(V1246)
						return
					} else {
						__e.TailApply(ShenFunc(symshen_4deref), Value, V1247)

						return
					}

				} else {
					__e.Return(V1246)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.deref"), gen14769)

		gen14773 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1250 := __args[0]
			_ = V1250
			V1251 := __args[1]
			_ = V1251
			gen14772 := Call(__e, ShenFunc(symshen_4pvar_2), V1250)

			if True == gen14772 {
				gen14770 := Call(__e, ShenFunc(symshen_4valvector), V1250, V1251)

				Value := gen14770
				gen14771 := Call(__e, ShenFunc(sym_a), Value, MakeSymbol("shen.-null-"))

				if True == gen14771 {
					__e.Return(V1250)
					return
				} else {
					__e.TailApply(ShenFunc(symshen_4lazyderef), Value, V1251)

					return
				}

			} else {
				__e.Return(V1250)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lazyderef"), gen14773)

		gen14777 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1254 := __args[0]
			_ = V1254
			V1255 := __args[1]
			_ = V1255
			gen14774 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14775 := Call(__e, ShenFunc(sym_5_1address), gen14774, V1255)

			gen14776 := Call(__e, ShenFunc(sym_5_1address), V1254, MakeNumber(1))

			__e.TailApply(ShenFunc(sym_5_1address), gen14775, gen14776)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.valvector"), gen14777)

		gen14780 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1260 := __args[0]
			_ = V1260
			V1261 := __args[1]
			_ = V1261
			V1262 := __args[2]
			_ = V1262
			V1263 := __args[3]
			_ = V1263
			gen14778 := Call(__e, ShenFunc(symshen_4lazyderef), V1260, V1262)

			gen14779 := Call(__e, ShenFunc(symshen_4lazyderef), V1261, V1262)

			__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen14778, gen14779, V1262, V1263)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("unify!"), gen14780)

		gen14804 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1285 := __args[0]
			_ = V1285
			V1286 := __args[1]
			_ = V1286
			V1287 := __args[2]
			_ = V1287
			V1288 := __args[3]
			_ = V1288
			gen14803 := Call(__e, ShenFunc(sym_a), V1286, V1285)

			if True == gen14803 {
				__e.TailApply(ShenFunc(symthaw), V1288)

				return
			} else {
				gen14801 := Call(__e, ShenFunc(symshen_4pvar_2), V1285)

				var gen14802 Obj
				if True == gen14801 {
					gen14798 := Call(__e, ShenFunc(symshen_4deref), V1286, V1287)

					gen14799 := Call(__e, ShenFunc(symshen_4occurs_2), V1285, gen14798)

					gen14800 := Call(__e, ShenFunc(symnot), gen14799)

					if True == gen14800 {
						gen14802 = True
					} else {
						gen14802 = False
					}

				} else {
					gen14802 = False
				}
				if True == gen14802 {
					__e.TailApply(ShenFunc(symbind), V1285, V1286, V1287, V1288)

					return
				} else {
					gen14796 := Call(__e, ShenFunc(symshen_4pvar_2), V1286)

					var gen14797 Obj
					if True == gen14796 {
						gen14793 := Call(__e, ShenFunc(symshen_4deref), V1285, V1287)

						gen14794 := Call(__e, ShenFunc(symshen_4occurs_2), V1286, gen14793)

						gen14795 := Call(__e, ShenFunc(symnot), gen14794)

						if True == gen14795 {
							gen14797 = True
						} else {
							gen14797 = False
						}

					} else {
						gen14797 = False
					}
					if True == gen14797 {
						__e.TailApply(ShenFunc(symbind), V1286, V1285, V1287, V1288)

						return
					} else {
						gen14791 := Call(__e, ShenFunc(symcons_2), V1285)

						var gen14792 Obj
						if True == gen14791 {
							gen14790 := Call(__e, ShenFunc(symcons_2), V1286)

							if True == gen14790 {
								gen14792 = True
							} else {
								gen14792 = False
							}

						} else {
							gen14792 = False
						}
						if True == gen14792 {
							gen14781 := Call(__e, ShenFunc(symhd), V1285)

							gen14782 := Call(__e, ShenFunc(symshen_4lazyderef), gen14781, V1287)

							gen14783 := Call(__e, ShenFunc(symhd), V1286)

							gen14784 := Call(__e, ShenFunc(symshen_4lazyderef), gen14783, V1287)

							gen14789 := MakeNative(func(__e Evaluator, __args ...Obj) {
								gen14785 := Call(__e, ShenFunc(symtl), V1285)

								gen14786 := Call(__e, ShenFunc(symshen_4lazyderef), gen14785, V1287)

								gen14787 := Call(__e, ShenFunc(symtl), V1286)

								gen14788 := Call(__e, ShenFunc(symshen_4lazyderef), gen14787, V1287)

								__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen14786, gen14788, V1287, V1288)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4lzy_a_b), gen14782, gen14784, V1287, gen14789)

							return

						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy=!"), gen14804)

		gen14811 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1300 := __args[0]
			_ = V1300
			V1301 := __args[1]
			_ = V1301
			gen14810 := Call(__e, ShenFunc(sym_a), V1301, V1300)

			if True == gen14810 {
				__e.Return(True)
				return
			} else {
				gen14809 := Call(__e, ShenFunc(symcons_2), V1301)

				if True == gen14809 {
					gen14807 := Call(__e, ShenFunc(symhd), V1301)

					gen14808 := Call(__e, ShenFunc(symshen_4occurs_2), V1300, gen14807)

					if True == gen14808 {
						__e.Return(True)
						return
					} else {
						gen14805 := Call(__e, ShenFunc(symtl), V1301)

						gen14806 := Call(__e, ShenFunc(symshen_4occurs_2), V1300, gen14805)

						if True == gen14806 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				} else {
					__e.Return(False)
					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.occurs?"), gen14811)

		gen14814 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1306 := __args[0]
			_ = V1306
			V1307 := __args[1]
			_ = V1307
			V1308 := __args[2]
			_ = V1308
			V1309 := __args[3]
			_ = V1309
			gen14812 := Call(__e, ShenFunc(symshen_4lazyderef), V1306, V1308)

			gen14813 := Call(__e, ShenFunc(symshen_4lazyderef), V1307, V1308)

			__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen14812, gen14813, V1308, V1309)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("identical"), gen14814)

		gen14826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1331 := __args[0]
			_ = V1331
			V1332 := __args[1]
			_ = V1332
			V1333 := __args[2]
			_ = V1333
			V1334 := __args[3]
			_ = V1334
			gen14825 := Call(__e, ShenFunc(sym_a), V1332, V1331)

			if True == gen14825 {
				__e.TailApply(ShenFunc(symthaw), V1334)

				return
			} else {
				gen14823 := Call(__e, ShenFunc(symcons_2), V1331)

				var gen14824 Obj
				if True == gen14823 {
					gen14822 := Call(__e, ShenFunc(symcons_2), V1332)

					if True == gen14822 {
						gen14824 = True
					} else {
						gen14824 = False
					}

				} else {
					gen14824 = False
				}
				if True == gen14824 {
					gen14815 := Call(__e, ShenFunc(symhd), V1331)

					gen14816 := Call(__e, ShenFunc(symshen_4lazyderef), gen14815, V1333)

					gen14817 := Call(__e, ShenFunc(symhd), V1332)

					gen14818 := Call(__e, ShenFunc(symshen_4lazyderef), gen14817, V1333)

					gen14821 := MakeNative(func(__e Evaluator, __args ...Obj) {
						gen14819 := Call(__e, ShenFunc(symtl), V1331)

						gen14820 := Call(__e, ShenFunc(symtl), V1332)

						__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen14819, gen14820, V1333, V1334)

						return

					}, 0)
					__e.TailApply(ShenFunc(symshen_4lzy_a_a), gen14816, gen14818, V1333, gen14821)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.lzy=="), gen14826)

		gen14829 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1336 := __args[0]
			_ = V1336
			gen14827 := Call(__e, ShenFunc(sym_5_1address), V1336, MakeNumber(1))

			gen14828 := Call(__e, ShenFunc(symshen_4app), gen14827, MakeString(""), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symcn), MakeString("Var"), gen14828)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pvar"), gen14829)

		gen14831 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1341 := __args[0]
			_ = V1341
			V1342 := __args[1]
			_ = V1342
			V1343 := __args[2]
			_ = V1343
			V1344 := __args[3]
			_ = V1344
			Call(__e, ShenFunc(symshen_4bindv), V1341, V1342, V1343)
			gen14830 := Call(__e, ShenFunc(symthaw), V1344)

			Result := gen14830
			Call(__e, ShenFunc(symshen_4unbindv), V1341, V1343)
			__e.Return(Result)
			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("bind"), gen14831)

		gen14836 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1362 := __args[0]
			_ = V1362
			V1363 := __args[1]
			_ = V1363
			V1364 := __args[2]
			_ = V1364
			gen14835 := Call(__e, ShenFunc(sym_a), True, V1362)

			if True == gen14835 {
				__e.TailApply(ShenFunc(symthaw), V1364)

				return
			} else {
				gen14834 := Call(__e, ShenFunc(sym_a), False, V1362)

				if True == gen14834 {
					__e.Return(False)
					return
				} else {
					gen14832 := Call(__e, ShenFunc(symshen_4app), V1362, MakeString("%"), MakeSymbol("shen.s"))

					gen14833 := Call(__e, ShenFunc(symcn), MakeString("fwhen expects a boolean: not "), gen14832)

					__e.TailApply(ShenFunc(symsimple_1error), gen14833)

					return

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("fwhen"), gen14836)

		gen14844 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1380 := __args[0]
			_ = V1380
			V1381 := __args[1]
			_ = V1381
			V1382 := __args[2]
			_ = V1382
			gen14843 := Call(__e, ShenFunc(symcons_2), V1380)

			if True == gen14843 {
				gen14839 := Call(__e, ShenFunc(symhd), V1380)

				gen14840 := Call(__e, ShenFunc(symshen_4lazyderef), gen14839, V1381)

				gen14841 := Call(__e, ShenFunc(symfunction), gen14840)

				gen14842 := Call(__e, ShenFunc(symtl), V1380)

				__e.TailApply(ShenFunc(symshen_4call_1help), gen14841, gen14842, V1381, V1382)

				return

			} else {
				gen14838 := Call(__e, ShenFunc(symshen_4pvar_2), V1380)

				if True == gen14838 {
					gen14837 := Call(__e, ShenFunc(symshen_4lazyderef), V1380, V1381)

					__e.TailApply(ShenFunc(symcall), gen14837, V1381, V1382)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("call"), gen14844)

		gen14850 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1387 := __args[0]
			_ = V1387
			V1388 := __args[1]
			_ = V1388
			V1389 := __args[2]
			_ = V1389
			V1390 := __args[3]
			_ = V1390
			gen14849 := Call(__e, ShenFunc(sym_a), Nil, V1388)

			if True == gen14849 {
				__e.TailApply(V1387, V1389, V1390)

				return
			} else {
				gen14848 := Call(__e, ShenFunc(symcons_2), V1388)

				if True == gen14848 {
					gen14845 := Call(__e, ShenFunc(symhd), V1388)

					gen14846 := Call(__e, V1387, gen14845)

					gen14847 := Call(__e, ShenFunc(symtl), V1388)

					__e.TailApply(ShenFunc(symshen_4call_1help), gen14846, gen14847, V1389, V1390)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call-help"))

					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call-help"), gen14850)

		gen14864 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1392 := __args[0]
			_ = V1392
			gen14862 := Call(__e, ShenFunc(symcons_2), V1392)

			var gen14863 Obj
			if True == gen14862 {
				gen14860 := Call(__e, ShenFunc(symhd), V1392)

				gen14861 := Call(__e, ShenFunc(symcons_2), gen14860)

				if True == gen14861 {
					gen14863 = True
				} else {
					gen14863 = False
				}

			} else {
				gen14863 = False
			}
			if True == gen14863 {
				gen14851 := Call(__e, ShenFunc(symshen_4start_1new_1prolog_1process))

				ProcessN := gen14851
				gen14852 := Call(__e, ShenFunc(symhd), V1392)

				gen14853 := Call(__e, ShenFunc(symhd), gen14852)

				gen14854 := Call(__e, ShenFunc(symhd), V1392)

				gen14855 := Call(__e, ShenFunc(symtl), gen14854)

				gen14856 := Call(__e, ShenFunc(symtl), V1392)

				gen14857 := Call(__e, ShenFunc(symcons), gen14856, Nil)

				gen14858 := Call(__e, ShenFunc(symcons), gen14855, gen14857)

				gen14859 := Call(__e, ShenFunc(symshen_4insert_1prolog_1variables), gen14858, ProcessN)

				__e.TailApply(ShenFunc(symshen_4intprolog_1help), gen14853, gen14859, ProcessN)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog"), gen14864)

		gen14876 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1396 := __args[0]
			_ = V1396
			V1397 := __args[1]
			_ = V1397
			V1398 := __args[2]
			_ = V1398
			gen14874 := Call(__e, ShenFunc(symcons_2), V1397)

			var gen14875 Obj
			if True == gen14874 {
				gen14871 := Call(__e, ShenFunc(symtl), V1397)

				gen14872 := Call(__e, ShenFunc(symcons_2), gen14871)

				var gen14873 Obj
				if True == gen14872 {
					gen14868 := Call(__e, ShenFunc(symtl), V1397)

					gen14869 := Call(__e, ShenFunc(symtl), gen14868)

					gen14870 := Call(__e, ShenFunc(sym_a), Nil, gen14869)

					if True == gen14870 {
						gen14873 = True
					} else {
						gen14873 = False
					}

				} else {
					gen14873 = False
				}
				if True == gen14873 {
					gen14875 = True
				} else {
					gen14875 = False
				}

			} else {
				gen14875 = False
			}
			if True == gen14875 {
				gen14865 := Call(__e, ShenFunc(symhd), V1397)

				gen14866 := Call(__e, ShenFunc(symtl), V1397)

				gen14867 := Call(__e, ShenFunc(symhd), gen14866)

				__e.TailApply(ShenFunc(symshen_4intprolog_1help_1help), V1396, gen14865, gen14867, V1398)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog-help"))

				return
			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog-help"), gen14876)

		gen14883 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1403 := __args[0]
			_ = V1403
			V1404 := __args[1]
			_ = V1404
			V1405 := __args[2]
			_ = V1405
			V1406 := __args[3]
			_ = V1406
			gen14882 := Call(__e, ShenFunc(sym_a), Nil, V1404)

			if True == gen14882 {
				gen14881 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.TailApply(ShenFunc(symshen_4call_1rest), V1405, V1406)

					return
				}, 0)
				__e.TailApply(V1403, V1406, gen14881)

				return

			} else {
				gen14880 := Call(__e, ShenFunc(symcons_2), V1404)

				if True == gen14880 {
					gen14877 := Call(__e, ShenFunc(symhd), V1404)

					gen14878 := Call(__e, V1403, gen14877)

					gen14879 := Call(__e, ShenFunc(symtl), V1404)

					__e.TailApply(ShenFunc(symshen_4intprolog_1help_1help), gen14878, gen14879, V1405, V1406)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.intprolog-help-help"))

					return
				}

			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intprolog-help-help"), gen14883)

		gen14917 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1411 := __args[0]
			_ = V1411
			V1412 := __args[1]
			_ = V1412
			gen14916 := Call(__e, ShenFunc(sym_a), Nil, V1411)

			if True == gen14916 {
				__e.Return(True)
				return
			} else {
				gen14914 := Call(__e, ShenFunc(symcons_2), V1411)

				var gen14915 Obj
				if True == gen14914 {
					gen14911 := Call(__e, ShenFunc(symhd), V1411)

					gen14912 := Call(__e, ShenFunc(symcons_2), gen14911)

					var gen14913 Obj
					if True == gen14912 {
						gen14908 := Call(__e, ShenFunc(symhd), V1411)

						gen14909 := Call(__e, ShenFunc(symtl), gen14908)

						gen14910 := Call(__e, ShenFunc(symcons_2), gen14909)

						if True == gen14910 {
							gen14913 = True
						} else {
							gen14913 = False
						}

					} else {
						gen14913 = False
					}
					if True == gen14913 {
						gen14915 = True
					} else {
						gen14915 = False
					}

				} else {
					gen14915 = False
				}
				if True == gen14915 {
					gen14896 := Call(__e, ShenFunc(symhd), V1411)

					gen14897 := Call(__e, ShenFunc(symhd), gen14896)

					f33 := gen14897
					gen14898 := Call(__e, ShenFunc(symhd), V1411)

					gen14899 := Call(__e, ShenFunc(symtl), gen14898)

					gen14900 := Call(__e, ShenFunc(symhd), gen14899)

					gen14901 := Call(__e, f33, gen14900)

					gen14902 := Call(__e, ShenFunc(symhd), V1411)

					gen14903 := Call(__e, ShenFunc(symtl), gen14902)

					gen14904 := Call(__e, ShenFunc(symtl), gen14903)

					gen14905 := Call(__e, ShenFunc(symcons), gen14901, gen14904)

					gen14906 := Call(__e, ShenFunc(symtl), V1411)

					gen14907 := Call(__e, ShenFunc(symcons), gen14905, gen14906)

					__e.TailApply(ShenFunc(symshen_4call_1rest), gen14907, V1412)

					return

				} else {
					gen14894 := Call(__e, ShenFunc(symcons_2), V1411)

					var gen14895 Obj
					if True == gen14894 {
						gen14891 := Call(__e, ShenFunc(symhd), V1411)

						gen14892 := Call(__e, ShenFunc(symcons_2), gen14891)

						var gen14893 Obj
						if True == gen14892 {
							gen14888 := Call(__e, ShenFunc(symhd), V1411)

							gen14889 := Call(__e, ShenFunc(symtl), gen14888)

							gen14890 := Call(__e, ShenFunc(sym_a), Nil, gen14889)

							if True == gen14890 {
								gen14893 = True
							} else {
								gen14893 = False
							}

						} else {
							gen14893 = False
						}
						if True == gen14893 {
							gen14895 = True
						} else {
							gen14895 = False
						}

					} else {
						gen14895 = False
					}
					if True == gen14895 {
						gen14884 := Call(__e, ShenFunc(symhd), V1411)

						gen14885 := Call(__e, ShenFunc(symhd), gen14884)

						f34 := gen14885
						gen14887 := MakeNative(func(__e Evaluator, __args ...Obj) {
							gen14886 := Call(__e, ShenFunc(symtl), V1411)

							__e.TailApply(ShenFunc(symshen_4call_1rest), gen14886, V1412)

							return

						}, 0)
						__e.TailApply(f34, V1412, gen14887)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.call-rest"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.call-rest"), gen14917)

		gen14921 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen14918 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*process-counter*"))

			gen14919 := Call(__e, ShenFunc(sym_7), MakeNumber(1), gen14918)

			gen14920 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*process-counter*"), gen14919)

			IncrementProcessCounter := gen14920
			__e.TailApply(ShenFunc(symshen_4initialise_1prolog), IncrementProcessCounter)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.start-new-prolog-process"), gen14921)

		gen14923 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1415 := __args[0]
			_ = V1415
			V1416 := __args[1]
			_ = V1416
			gen14922 := Call(__e, ShenFunc(symshen_4flatten), V1415)

			__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), V1415, gen14922, V1416)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-prolog-variables"), gen14923)

		gen14937 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1424 := __args[0]
			_ = V1424
			V1425 := __args[1]
			_ = V1425
			V1426 := __args[2]
			_ = V1426
			gen14936 := Call(__e, ShenFunc(sym_a), Nil, V1425)

			if True == gen14936 {
				__e.Return(V1424)
				return
			} else {
				gen14934 := Call(__e, ShenFunc(symcons_2), V1425)

				var gen14935 Obj
				if True == gen14934 {
					gen14932 := Call(__e, ShenFunc(symhd), V1425)

					gen14933 := Call(__e, ShenFunc(symvariable_2), gen14932)

					if True == gen14933 {
						gen14935 = True
					} else {
						gen14935 = False
					}

				} else {
					gen14935 = False
				}
				if True == gen14935 {
					gen14926 := Call(__e, ShenFunc(symshen_4newpv), V1426)

					V := gen14926
					gen14927 := Call(__e, ShenFunc(symhd), V1425)

					gen14928 := Call(__e, ShenFunc(symsubst), V, gen14927, V1424)

					XV_cY := gen14928
					gen14929 := Call(__e, ShenFunc(symhd), V1425)

					gen14930 := Call(__e, ShenFunc(symtl), V1425)

					gen14931 := Call(__e, ShenFunc(symremove), gen14929, gen14930)

					Z_1Y := gen14931
					__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), XV_cY, Z_1Y, V1426)

					return

				} else {
					gen14925 := Call(__e, ShenFunc(symcons_2), V1425)

					if True == gen14925 {
						gen14924 := Call(__e, ShenFunc(symtl), V1425)

						__e.TailApply(ShenFunc(symshen_4insert_1prolog_1variables_1help), V1424, gen14924, V1426)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.insert-prolog-variables-help"))

						return
					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.insert-prolog-variables-help"), gen14937)

		gen14944 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V1428 := __args[0]
			_ = V1428
			gen14938 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*prologvectors*"))

			gen14939 := Call(__e, ShenFunc(symvector), MakeNumber(10))

			gen14940 := Call(__e, ShenFunc(symshen_4fillvector), gen14939, MakeNumber(1), MakeNumber(10), MakeSymbol("shen.-null-"))

			gen14941 := Call(__e, ShenFunc(symaddress_1_6), gen14938, V1428, gen14940)

			Vector := gen14941
			_ = Vector
			gen14942 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*varcounter*"))

			gen14943 := Call(__e, ShenFunc(symaddress_1_6), gen14942, V1428, MakeNumber(1))

			Counter := gen14943
			_ = Counter
			__e.Return(V1428)
			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.initialise-prolog"), gen14944)

		return

	}, 0))
}
