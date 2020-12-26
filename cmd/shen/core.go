package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen4 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V152 := __args[0]
			_ = V152
			V153 := __args[1]
			_ = V153
			gen1 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4_5define_6), X)

				return
			}, 1)
			gen2 := Call(__e, ShenFunc(symcons), V152, V153)

			gen3 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4shen_1syntax_1error), V152, X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symcompile), gen1, gen2, gen3)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.shen->kl"), gen4)

		gen14 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V160 := __args[0]
			_ = V160
			V161 := __args[1]
			_ = V161
			gen13 := Call(__e, ShenFunc(symcons_2), V161)

			if True == gen13 {
				gen7 := Call(__e, ShenFunc(symhd), V161)

				gen8 := Call(__e, ShenFunc(symshen_4next_150), MakeNumber(50), gen7)

				gen9 := Call(__e, ShenFunc(symshen_4app), gen8, MakeString("\n"), MakeSymbol("shen.a"))

				gen10 := Call(__e, ShenFunc(symcn), MakeString(" here:\n\n "), gen9)

				gen11 := Call(__e, ShenFunc(symshen_4app), V160, gen10, MakeSymbol("shen.a"))

				gen12 := Call(__e, ShenFunc(symcn), MakeString("syntax error in "), gen11)

				__e.TailApply(ShenFunc(symsimple_1error), gen12)

				return

			} else {
				gen5 := Call(__e, ShenFunc(symshen_4app), V160, MakeString("\n"), MakeSymbol("shen.a"))

				gen6 := Call(__e, ShenFunc(symcn), MakeString("syntax error in "), gen5)

				__e.TailApply(ShenFunc(symsimple_1error), gen6)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.shen-syntax-error"), gen14)

		gen46 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V163 := __args[0]
			_ = V163
			gen15 := Call(__e, ShenFunc(symshen_4_5name_6), V163)

			Parse__shen_4_5name_6 := gen15
			gen28 := Call(__e, ShenFunc(symfail))

			gen29 := Call(__e, ShenFunc(sym_a), gen28, Parse__shen_4_5name_6)

			gen30 := Call(__e, ShenFunc(symnot), gen29)

			var gen31 Obj
			if True == gen30 {
				gen16 := Call(__e, ShenFunc(symshen_4_5signature_6), Parse__shen_4_5name_6)

				Parse__shen_4_5signature_6 := gen16
				gen25 := Call(__e, ShenFunc(symfail))

				gen26 := Call(__e, ShenFunc(sym_a), gen25, Parse__shen_4_5signature_6)

				gen27 := Call(__e, ShenFunc(symnot), gen26)

				if True == gen27 {
					gen17 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5signature_6)

					Parse__shen_4_5rules_6 := gen17
					gen22 := Call(__e, ShenFunc(symfail))

					gen23 := Call(__e, ShenFunc(sym_a), gen22, Parse__shen_4_5rules_6)

					gen24 := Call(__e, ShenFunc(symnot), gen23)

					if True == gen24 {
						gen18 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

						gen19 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5name_6)

						gen20 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

						gen21 := Call(__e, ShenFunc(symshen_4compile__to__machine__code), gen19, gen20)

						gen31 = Call(__e, ShenFunc(symshen_4pair), gen18, gen21)

					} else {
						gen31 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen31 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen31 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen31
			gen44 := Call(__e, ShenFunc(symfail))

			gen45 := Call(__e, ShenFunc(sym_a), YaccParse, gen44)

			if True == gen45 {
				gen32 := Call(__e, ShenFunc(symshen_4_5name_6), V163)

				Parse__shen_4_5name_6 := gen32
				gen41 := Call(__e, ShenFunc(symfail))

				gen42 := Call(__e, ShenFunc(sym_a), gen41, Parse__shen_4_5name_6)

				gen43 := Call(__e, ShenFunc(symnot), gen42)

				if True == gen43 {
					gen33 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5name_6)

					Parse__shen_4_5rules_6 := gen33
					gen38 := Call(__e, ShenFunc(symfail))

					gen39 := Call(__e, ShenFunc(sym_a), gen38, Parse__shen_4_5rules_6)

					gen40 := Call(__e, ShenFunc(symnot), gen39)

					if True == gen40 {
						gen34 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

						gen35 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5name_6)

						gen36 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

						gen37 := Call(__e, ShenFunc(symshen_4compile__to__machine__code), gen35, gen36)

						__e.TailApply(ShenFunc(symshen_4pair), gen34, gen37)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<define>"), gen46)

		gen60 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V165 := __args[0]
			_ = V165
			gen58 := Call(__e, ShenFunc(symhd), V165)

			gen59 := Call(__e, ShenFunc(symcons_2), gen58)

			if True == gen59 {
				gen47 := Call(__e, ShenFunc(symshen_4hdhd), V165)

				Parse__X := gen47
				gen48 := Call(__e, ShenFunc(symshen_4tlhd), V165)

				gen49 := Call(__e, ShenFunc(symshen_4hdtl), V165)

				gen50 := Call(__e, ShenFunc(symshen_4pair), gen48, gen49)

				gen51 := Call(__e, ShenFunc(symhd), gen50)

				gen55 := Call(__e, ShenFunc(symsymbol_2), Parse__X)

				var gen56 Obj
				if True == gen55 {
					gen53 := Call(__e, ShenFunc(symshen_4sysfunc_2), Parse__X)

					gen54 := Call(__e, ShenFunc(symnot), gen53)

					if True == gen54 {
						gen56 = True
					} else {
						gen56 = False
					}

				} else {
					gen56 = False
				}
				var gen57 Obj
				if True == gen56 {
					gen57 = Parse__X
				} else {
					gen52 := Call(__e, ShenFunc(symshen_4app), Parse__X, MakeString(" is not a legitimate function name.\n"), MakeSymbol("shen.a"))

					gen57 = Call(__e, ShenFunc(symsimple_1error), gen52)

				}
				__e.TailApply(ShenFunc(symshen_4pair), gen51, gen57)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<name>"), gen60)

		gen64 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V167 := __args[0]
			_ = V167
			gen61 := Call(__e, ShenFunc(symintern), MakeString("shen"))

			gen62 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			gen63 := Call(__e, ShenFunc(symget), gen61, MakeSymbol("shen.external-symbols"), gen62)

			__e.TailApply(ShenFunc(symelement_2), V167, gen63)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sysfunc?"), gen64)

		gen89 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V171 := __args[0]
			_ = V171
			gen86 := Call(__e, ShenFunc(symhd), V171)

			gen87 := Call(__e, ShenFunc(symcons_2), gen86)

			var gen88 Obj
			if True == gen87 {
				gen84 := Call(__e, ShenFunc(symshen_4hdhd), V171)

				gen85 := Call(__e, ShenFunc(sym_a), MakeSymbol("{"), gen84)

				if True == gen85 {
					gen88 = True
				} else {
					gen88 = False
				}

			} else {
				gen88 = False
			}
			if True == gen88 {
				gen65 := Call(__e, ShenFunc(symshen_4tlhd), V171)

				gen66 := Call(__e, ShenFunc(symshen_4hdtl), V171)

				gen67 := Call(__e, ShenFunc(symshen_4pair), gen65, gen66)

				NewStream168 := gen67
				gen68 := Call(__e, ShenFunc(symshen_4_5signature_1help_6), NewStream168)

				Parse__shen_4_5signature_1help_6 := gen68
				gen81 := Call(__e, ShenFunc(symfail))

				gen82 := Call(__e, ShenFunc(sym_a), gen81, Parse__shen_4_5signature_1help_6)

				gen83 := Call(__e, ShenFunc(symnot), gen82)

				if True == gen83 {
					gen78 := Call(__e, ShenFunc(symhd), Parse__shen_4_5signature_1help_6)

					gen79 := Call(__e, ShenFunc(symcons_2), gen78)

					var gen80 Obj
					if True == gen79 {
						gen76 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5signature_1help_6)

						gen77 := Call(__e, ShenFunc(sym_a), MakeSymbol("}"), gen76)

						if True == gen77 {
							gen80 = True
						} else {
							gen80 = False
						}

					} else {
						gen80 = False
					}
					if True == gen80 {
						gen69 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5signature_1help_6)

						gen70 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen71 := Call(__e, ShenFunc(symshen_4pair), gen69, gen70)

						NewStream169 := gen71
						gen72 := Call(__e, ShenFunc(symhd), NewStream169)

						gen73 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen74 := Call(__e, ShenFunc(symshen_4curry_1type), gen73)

						gen75 := Call(__e, ShenFunc(symshen_4demodulate), gen74)

						__e.TailApply(ShenFunc(symshen_4pair), gen72, gen75)

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

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<signature>"), gen89)

		gen91 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V173 := __args[0]
			_ = V173
			gen90 := Call(__e, ShenFunc(symshen_4curry_1type_1h), V173)

			__e.TailApply(ShenFunc(symshen_4active_1cons), gen90)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-type"), gen91)

		gen120 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V175 := __args[0]
			_ = V175
			gen118 := Call(__e, ShenFunc(symcons_2), V175)

			var gen119 Obj
			if True == gen118 {
				gen115 := Call(__e, ShenFunc(symtl), V175)

				gen116 := Call(__e, ShenFunc(symcons_2), gen115)

				var gen117 Obj
				if True == gen116 {
					gen111 := Call(__e, ShenFunc(symtl), V175)

					gen112 := Call(__e, ShenFunc(symtl), gen111)

					gen113 := Call(__e, ShenFunc(symcons_2), gen112)

					var gen114 Obj
					if True == gen113 {
						gen106 := Call(__e, ShenFunc(symtl), V175)

						gen107 := Call(__e, ShenFunc(symtl), gen106)

						gen108 := Call(__e, ShenFunc(symtl), gen107)

						gen109 := Call(__e, ShenFunc(sym_a), Nil, gen108)

						var gen110 Obj
						if True == gen109 {
							gen103 := Call(__e, ShenFunc(symtl), V175)

							gen104 := Call(__e, ShenFunc(symhd), gen103)

							gen105 := Call(__e, ShenFunc(sym_a), gen104, MakeSymbol("bar!"))

							if True == gen105 {
								gen110 = True
							} else {
								gen110 = False
							}

						} else {
							gen110 = False
						}
						if True == gen110 {
							gen114 = True
						} else {
							gen114 = False
						}

					} else {
						gen114 = False
					}
					if True == gen114 {
						gen117 = True
					} else {
						gen117 = False
					}

				} else {
					gen117 = False
				}
				if True == gen117 {
					gen119 = True
				} else {
					gen119 = False
				}

			} else {
				gen119 = False
			}
			if True == gen119 {
				gen97 := Call(__e, ShenFunc(symhd), V175)

				gen98 := Call(__e, ShenFunc(symshen_4active_1cons), gen97)

				gen99 := Call(__e, ShenFunc(symtl), V175)

				gen100 := Call(__e, ShenFunc(symtl), gen99)

				gen101 := Call(__e, ShenFunc(symhd), gen100)

				gen102 := Call(__e, ShenFunc(symshen_4active_1cons), gen101)

				__e.TailApply(ShenFunc(symcons), gen98, gen102)

				return

			} else {
				gen96 := Call(__e, ShenFunc(symcons_2), V175)

				if True == gen96 {
					gen92 := Call(__e, ShenFunc(symhd), V175)

					gen93 := Call(__e, ShenFunc(symshen_4active_1cons), gen92)

					gen94 := Call(__e, ShenFunc(symtl), V175)

					gen95 := Call(__e, ShenFunc(symshen_4active_1cons), gen94)

					__e.TailApply(ShenFunc(symcons), gen93, gen95)

					return

				} else {
					__e.Return(V175)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.active-cons"), gen120)

		gen181 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V177 := __args[0]
			_ = V177
			gen179 := Call(__e, ShenFunc(symcons_2), V177)

			var gen180 Obj
			if True == gen179 {
				gen176 := Call(__e, ShenFunc(symtl), V177)

				gen177 := Call(__e, ShenFunc(symcons_2), gen176)

				var gen178 Obj
				if True == gen177 {
					gen172 := Call(__e, ShenFunc(symtl), V177)

					gen173 := Call(__e, ShenFunc(symhd), gen172)

					gen174 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen173)

					var gen175 Obj
					if True == gen174 {
						gen168 := Call(__e, ShenFunc(symtl), V177)

						gen169 := Call(__e, ShenFunc(symtl), gen168)

						gen170 := Call(__e, ShenFunc(symcons_2), gen169)

						var gen171 Obj
						if True == gen170 {
							gen163 := Call(__e, ShenFunc(symtl), V177)

							gen164 := Call(__e, ShenFunc(symtl), gen163)

							gen165 := Call(__e, ShenFunc(symtl), gen164)

							gen166 := Call(__e, ShenFunc(symcons_2), gen165)

							var gen167 Obj
							if True == gen166 {
								gen158 := Call(__e, ShenFunc(symtl), V177)

								gen159 := Call(__e, ShenFunc(symtl), gen158)

								gen160 := Call(__e, ShenFunc(symtl), gen159)

								gen161 := Call(__e, ShenFunc(symhd), gen160)

								gen162 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen161)

								if True == gen162 {
									gen167 = True
								} else {
									gen167 = False
								}

							} else {
								gen167 = False
							}
							if True == gen167 {
								gen171 = True
							} else {
								gen171 = False
							}

						} else {
							gen171 = False
						}
						if True == gen171 {
							gen175 = True
						} else {
							gen175 = False
						}

					} else {
						gen175 = False
					}
					if True == gen175 {
						gen178 = True
					} else {
						gen178 = False
					}

				} else {
					gen178 = False
				}
				if True == gen178 {
					gen180 = True
				} else {
					gen180 = False
				}

			} else {
				gen180 = False
			}
			if True == gen180 {
				gen152 := Call(__e, ShenFunc(symhd), V177)

				gen153 := Call(__e, ShenFunc(symtl), V177)

				gen154 := Call(__e, ShenFunc(symtl), gen153)

				gen155 := Call(__e, ShenFunc(symcons), gen154, Nil)

				gen156 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen155)

				gen157 := Call(__e, ShenFunc(symcons), gen152, gen156)

				__e.TailApply(ShenFunc(symshen_4curry_1type_1h), gen157)

				return

			} else {
				gen150 := Call(__e, ShenFunc(symcons_2), V177)

				var gen151 Obj
				if True == gen150 {
					gen147 := Call(__e, ShenFunc(symtl), V177)

					gen148 := Call(__e, ShenFunc(symcons_2), gen147)

					var gen149 Obj
					if True == gen148 {
						gen143 := Call(__e, ShenFunc(symtl), V177)

						gen144 := Call(__e, ShenFunc(symhd), gen143)

						gen145 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), gen144)

						var gen146 Obj
						if True == gen145 {
							gen139 := Call(__e, ShenFunc(symtl), V177)

							gen140 := Call(__e, ShenFunc(symtl), gen139)

							gen141 := Call(__e, ShenFunc(symcons_2), gen140)

							var gen142 Obj
							if True == gen141 {
								gen134 := Call(__e, ShenFunc(symtl), V177)

								gen135 := Call(__e, ShenFunc(symtl), gen134)

								gen136 := Call(__e, ShenFunc(symtl), gen135)

								gen137 := Call(__e, ShenFunc(symcons_2), gen136)

								var gen138 Obj
								if True == gen137 {
									gen129 := Call(__e, ShenFunc(symtl), V177)

									gen130 := Call(__e, ShenFunc(symtl), gen129)

									gen131 := Call(__e, ShenFunc(symtl), gen130)

									gen132 := Call(__e, ShenFunc(symhd), gen131)

									gen133 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), gen132)

									if True == gen133 {
										gen138 = True
									} else {
										gen138 = False
									}

								} else {
									gen138 = False
								}
								if True == gen138 {
									gen142 = True
								} else {
									gen142 = False
								}

							} else {
								gen142 = False
							}
							if True == gen142 {
								gen146 = True
							} else {
								gen146 = False
							}

						} else {
							gen146 = False
						}
						if True == gen146 {
							gen149 = True
						} else {
							gen149 = False
						}

					} else {
						gen149 = False
					}
					if True == gen149 {
						gen151 = True
					} else {
						gen151 = False
					}

				} else {
					gen151 = False
				}
				if True == gen151 {
					gen123 := Call(__e, ShenFunc(symhd), V177)

					gen124 := Call(__e, ShenFunc(symtl), V177)

					gen125 := Call(__e, ShenFunc(symtl), gen124)

					gen126 := Call(__e, ShenFunc(symcons), gen125, Nil)

					gen127 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen126)

					gen128 := Call(__e, ShenFunc(symcons), gen123, gen127)

					__e.TailApply(ShenFunc(symshen_4curry_1type_1h), gen128)

					return

				} else {
					gen122 := Call(__e, ShenFunc(symcons_2), V177)

					if True == gen122 {
						gen121 := MakeNative(func(__e Evaluator, __args ...Obj) {
							Z := __args[0]
							_ = Z
							__e.TailApply(ShenFunc(symshen_4curry_1type_1h), Z)

							return
						}, 1)
						__e.TailApply(ShenFunc(symmap), gen121, V177)

						return

					} else {
						__e.Return(V177)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-type-h"), gen181)

		gen207 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V179 := __args[0]
			_ = V179
			gen197 := Call(__e, ShenFunc(symhd), V179)

			gen198 := Call(__e, ShenFunc(symcons_2), gen197)

			var gen199 Obj
			if True == gen198 {
				gen182 := Call(__e, ShenFunc(symshen_4hdhd), V179)

				Parse__X := gen182
				gen183 := Call(__e, ShenFunc(symshen_4tlhd), V179)

				gen184 := Call(__e, ShenFunc(symshen_4hdtl), V179)

				gen185 := Call(__e, ShenFunc(symshen_4pair), gen183, gen184)

				gen186 := Call(__e, ShenFunc(symshen_4_5signature_1help_6), gen185)

				Parse__shen_4_5signature_1help_6 := gen186
				gen194 := Call(__e, ShenFunc(symfail))

				gen195 := Call(__e, ShenFunc(sym_a), gen194, Parse__shen_4_5signature_1help_6)

				gen196 := Call(__e, ShenFunc(symnot), gen195)

				if True == gen196 {
					gen190 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), Nil)

					gen191 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen190)

					gen192 := Call(__e, ShenFunc(symelement_2), Parse__X, gen191)

					gen193 := Call(__e, ShenFunc(symnot), gen192)

					if True == gen193 {
						gen187 := Call(__e, ShenFunc(symhd), Parse__shen_4_5signature_1help_6)

						gen188 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_1help_6)

						gen189 := Call(__e, ShenFunc(symcons), Parse__X, gen188)

						gen199 = Call(__e, ShenFunc(symshen_4pair), gen187, gen189)

					} else {
						gen199 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen199 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen199 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen199
			gen205 := Call(__e, ShenFunc(symfail))

			gen206 := Call(__e, ShenFunc(sym_a), YaccParse, gen205)

			if True == gen206 {
				gen200 := Call(__e, ShenFunc(sym_5e_6), V179)

				Parse___5e_6 := gen200
				gen202 := Call(__e, ShenFunc(symfail))

				gen203 := Call(__e, ShenFunc(sym_a), gen202, Parse___5e_6)

				gen204 := Call(__e, ShenFunc(symnot), gen203)

				if True == gen204 {
					gen201 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen201, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<signature-help>"), gen207)

		gen232 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V181 := __args[0]
			_ = V181
			gen208 := Call(__e, ShenFunc(symshen_4_5rule_6), V181)

			Parse__shen_4_5rule_6 := gen208
			gen218 := Call(__e, ShenFunc(symfail))

			gen219 := Call(__e, ShenFunc(sym_a), gen218, Parse__shen_4_5rule_6)

			gen220 := Call(__e, ShenFunc(symnot), gen219)

			var gen221 Obj
			if True == gen220 {
				gen209 := Call(__e, ShenFunc(symshen_4_5rules_6), Parse__shen_4_5rule_6)

				Parse__shen_4_5rules_6 := gen209
				gen215 := Call(__e, ShenFunc(symfail))

				gen216 := Call(__e, ShenFunc(sym_a), gen215, Parse__shen_4_5rules_6)

				gen217 := Call(__e, ShenFunc(symnot), gen216)

				if True == gen217 {
					gen210 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rules_6)

					gen211 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen212 := Call(__e, ShenFunc(symshen_4linearise), gen211)

					gen213 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rules_6)

					gen214 := Call(__e, ShenFunc(symcons), gen212, gen213)

					gen221 = Call(__e, ShenFunc(symshen_4pair), gen210, gen214)

				} else {
					gen221 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen221 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen221
			gen230 := Call(__e, ShenFunc(symfail))

			gen231 := Call(__e, ShenFunc(sym_a), YaccParse, gen230)

			if True == gen231 {
				gen222 := Call(__e, ShenFunc(symshen_4_5rule_6), V181)

				Parse__shen_4_5rule_6 := gen222
				gen227 := Call(__e, ShenFunc(symfail))

				gen228 := Call(__e, ShenFunc(sym_a), gen227, Parse__shen_4_5rule_6)

				gen229 := Call(__e, ShenFunc(symnot), gen228)

				if True == gen229 {
					gen223 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rule_6)

					gen224 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen225 := Call(__e, ShenFunc(symshen_4linearise), gen224)

					gen226 := Call(__e, ShenFunc(symcons), gen225, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen223, gen226)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rules>"), gen232)

		gen362 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V189 := __args[0]
			_ = V189
			gen233 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

			Parse__shen_4_5patterns_6 := gen233
			gen267 := Call(__e, ShenFunc(symfail))

			gen268 := Call(__e, ShenFunc(sym_a), gen267, Parse__shen_4_5patterns_6)

			gen269 := Call(__e, ShenFunc(symnot), gen268)

			var gen270 Obj
			if True == gen269 {
				gen264 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

				gen265 := Call(__e, ShenFunc(symcons_2), gen264)

				var gen266 Obj
				if True == gen265 {
					gen262 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

					gen263 := Call(__e, ShenFunc(sym_a), MakeSymbol("->"), gen262)

					if True == gen263 {
						gen266 = True
					} else {
						gen266 = False
					}

				} else {
					gen266 = False
				}
				if True == gen266 {
					gen234 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

					gen235 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

					gen236 := Call(__e, ShenFunc(symshen_4pair), gen234, gen235)

					NewStream182 := gen236
					gen237 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream182)

					Parse__shen_4_5action_6 := gen237
					gen259 := Call(__e, ShenFunc(symfail))

					gen260 := Call(__e, ShenFunc(sym_a), gen259, Parse__shen_4_5action_6)

					gen261 := Call(__e, ShenFunc(symnot), gen260)

					if True == gen261 {
						gen256 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

						gen257 := Call(__e, ShenFunc(symcons_2), gen256)

						var gen258 Obj
						if True == gen257 {
							gen254 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5action_6)

							gen255 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen254)

							if True == gen255 {
								gen258 = True
							} else {
								gen258 = False
							}

						} else {
							gen258 = False
						}
						if True == gen258 {
							gen238 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5action_6)

							gen239 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

							gen240 := Call(__e, ShenFunc(symshen_4pair), gen238, gen239)

							NewStream183 := gen240
							gen241 := Call(__e, ShenFunc(symshen_4_5guard_6), NewStream183)

							Parse__shen_4_5guard_6 := gen241
							gen251 := Call(__e, ShenFunc(symfail))

							gen252 := Call(__e, ShenFunc(sym_a), gen251, Parse__shen_4_5guard_6)

							gen253 := Call(__e, ShenFunc(symnot), gen252)

							if True == gen253 {
								gen242 := Call(__e, ShenFunc(symhd), Parse__shen_4_5guard_6)

								gen243 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

								gen244 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5guard_6)

								gen245 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

								gen246 := Call(__e, ShenFunc(symcons), gen245, Nil)

								gen247 := Call(__e, ShenFunc(symcons), gen244, gen246)

								gen248 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen247)

								gen249 := Call(__e, ShenFunc(symcons), gen248, Nil)

								gen250 := Call(__e, ShenFunc(symcons), gen243, gen249)

								gen270 = Call(__e, ShenFunc(symshen_4pair), gen242, gen250)

							} else {
								gen270 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen270 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen270 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen270 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen270 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen270
			gen360 := Call(__e, ShenFunc(symfail))

			gen361 := Call(__e, ShenFunc(sym_a), YaccParse, gen360)

			if True == gen361 {
				gen271 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

				Parse__shen_4_5patterns_6 := gen271
				gen289 := Call(__e, ShenFunc(symfail))

				gen290 := Call(__e, ShenFunc(sym_a), gen289, Parse__shen_4_5patterns_6)

				gen291 := Call(__e, ShenFunc(symnot), gen290)

				var gen292 Obj
				if True == gen291 {
					gen286 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

					gen287 := Call(__e, ShenFunc(symcons_2), gen286)

					var gen288 Obj
					if True == gen287 {
						gen284 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

						gen285 := Call(__e, ShenFunc(sym_a), MakeSymbol("->"), gen284)

						if True == gen285 {
							gen288 = True
						} else {
							gen288 = False
						}

					} else {
						gen288 = False
					}
					if True == gen288 {
						gen272 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

						gen273 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

						gen274 := Call(__e, ShenFunc(symshen_4pair), gen272, gen273)

						NewStream184 := gen274
						gen275 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream184)

						Parse__shen_4_5action_6 := gen275
						gen281 := Call(__e, ShenFunc(symfail))

						gen282 := Call(__e, ShenFunc(sym_a), gen281, Parse__shen_4_5action_6)

						gen283 := Call(__e, ShenFunc(symnot), gen282)

						if True == gen283 {
							gen276 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

							gen277 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

							gen278 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

							gen279 := Call(__e, ShenFunc(symcons), gen278, Nil)

							gen280 := Call(__e, ShenFunc(symcons), gen277, gen279)

							gen292 = Call(__e, ShenFunc(symshen_4pair), gen276, gen280)

						} else {
							gen292 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen292 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen292 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen292
				gen358 := Call(__e, ShenFunc(symfail))

				gen359 := Call(__e, ShenFunc(sym_a), YaccParse, gen358)

				if True == gen359 {
					gen293 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

					Parse__shen_4_5patterns_6 := gen293
					gen329 := Call(__e, ShenFunc(symfail))

					gen330 := Call(__e, ShenFunc(sym_a), gen329, Parse__shen_4_5patterns_6)

					gen331 := Call(__e, ShenFunc(symnot), gen330)

					var gen332 Obj
					if True == gen331 {
						gen326 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

						gen327 := Call(__e, ShenFunc(symcons_2), gen326)

						var gen328 Obj
						if True == gen327 {
							gen324 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

							gen325 := Call(__e, ShenFunc(sym_a), MakeSymbol("<-"), gen324)

							if True == gen325 {
								gen328 = True
							} else {
								gen328 = False
							}

						} else {
							gen328 = False
						}
						if True == gen328 {
							gen294 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

							gen295 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

							gen296 := Call(__e, ShenFunc(symshen_4pair), gen294, gen295)

							NewStream185 := gen296
							gen297 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream185)

							Parse__shen_4_5action_6 := gen297
							gen321 := Call(__e, ShenFunc(symfail))

							gen322 := Call(__e, ShenFunc(sym_a), gen321, Parse__shen_4_5action_6)

							gen323 := Call(__e, ShenFunc(symnot), gen322)

							if True == gen323 {
								gen318 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

								gen319 := Call(__e, ShenFunc(symcons_2), gen318)

								var gen320 Obj
								if True == gen319 {
									gen316 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5action_6)

									gen317 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen316)

									if True == gen317 {
										gen320 = True
									} else {
										gen320 = False
									}

								} else {
									gen320 = False
								}
								if True == gen320 {
									gen298 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5action_6)

									gen299 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

									gen300 := Call(__e, ShenFunc(symshen_4pair), gen298, gen299)

									NewStream186 := gen300
									gen301 := Call(__e, ShenFunc(symshen_4_5guard_6), NewStream186)

									Parse__shen_4_5guard_6 := gen301
									gen313 := Call(__e, ShenFunc(symfail))

									gen314 := Call(__e, ShenFunc(sym_a), gen313, Parse__shen_4_5guard_6)

									gen315 := Call(__e, ShenFunc(symnot), gen314)

									if True == gen315 {
										gen302 := Call(__e, ShenFunc(symhd), Parse__shen_4_5guard_6)

										gen303 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

										gen304 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5guard_6)

										gen305 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

										gen306 := Call(__e, ShenFunc(symcons), gen305, Nil)

										gen307 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.choicepoint!"), gen306)

										gen308 := Call(__e, ShenFunc(symcons), gen307, Nil)

										gen309 := Call(__e, ShenFunc(symcons), gen304, gen308)

										gen310 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen309)

										gen311 := Call(__e, ShenFunc(symcons), gen310, Nil)

										gen312 := Call(__e, ShenFunc(symcons), gen303, gen311)

										gen332 = Call(__e, ShenFunc(symshen_4pair), gen302, gen312)

									} else {
										gen332 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen332 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen332 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen332 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen332 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen332
					gen356 := Call(__e, ShenFunc(symfail))

					gen357 := Call(__e, ShenFunc(sym_a), YaccParse, gen356)

					if True == gen357 {
						gen333 := Call(__e, ShenFunc(symshen_4_5patterns_6), V189)

						Parse__shen_4_5patterns_6 := gen333
						gen353 := Call(__e, ShenFunc(symfail))

						gen354 := Call(__e, ShenFunc(sym_a), gen353, Parse__shen_4_5patterns_6)

						gen355 := Call(__e, ShenFunc(symnot), gen354)

						if True == gen355 {
							gen350 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

							gen351 := Call(__e, ShenFunc(symcons_2), gen350)

							var gen352 Obj
							if True == gen351 {
								gen348 := Call(__e, ShenFunc(symshen_4hdhd), Parse__shen_4_5patterns_6)

								gen349 := Call(__e, ShenFunc(sym_a), MakeSymbol("<-"), gen348)

								if True == gen349 {
									gen352 = True
								} else {
									gen352 = False
								}

							} else {
								gen352 = False
							}
							if True == gen352 {
								gen334 := Call(__e, ShenFunc(symshen_4tlhd), Parse__shen_4_5patterns_6)

								gen335 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

								gen336 := Call(__e, ShenFunc(symshen_4pair), gen334, gen335)

								NewStream187 := gen336
								gen337 := Call(__e, ShenFunc(symshen_4_5action_6), NewStream187)

								Parse__shen_4_5action_6 := gen337
								gen345 := Call(__e, ShenFunc(symfail))

								gen346 := Call(__e, ShenFunc(sym_a), gen345, Parse__shen_4_5action_6)

								gen347 := Call(__e, ShenFunc(symnot), gen346)

								if True == gen347 {
									gen338 := Call(__e, ShenFunc(symhd), Parse__shen_4_5action_6)

									gen339 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

									gen340 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5action_6)

									gen341 := Call(__e, ShenFunc(symcons), gen340, Nil)

									gen342 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.choicepoint!"), gen341)

									gen343 := Call(__e, ShenFunc(symcons), gen342, Nil)

									gen344 := Call(__e, ShenFunc(symcons), gen339, gen343)

									__e.TailApply(ShenFunc(symshen_4pair), gen338, gen344)

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

				} else {
					__e.Return(YaccParse)
					return
				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<rule>"), gen362)

		gen364 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V192 := __args[0]
			_ = V192
			V193 := __args[1]
			_ = V193
			gen363 := Call(__e, V192, V193)

			if True == gen363 {
				__e.TailApply(ShenFunc(symfail))

				return
			} else {
				__e.Return(V193)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.fail_if"), gen364)

		gen367 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V199 := __args[0]
			_ = V199
			gen365 := Call(__e, ShenFunc(symfail))

			gen366 := Call(__e, ShenFunc(sym_a), V199, gen365)

			if True == gen366 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.succeeds?"), gen367)

		gen369 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V202 := __args[0]
			_ = V202
			V203 := __args[1]
			_ = V203
			gen368 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*custom-pattern-compiler*"))

			f28 := gen368
			__e.TailApply(f28, V202, V203)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.custom-pattern-compiler"), gen369)

		gen371 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V205 := __args[0]
			_ = V205
			gen370 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*custom-pattern-reducer*"))

			f29 := gen370
			__e.TailApply(f29, V205)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.custom-pattern-reducer"), gen371)

		gen392 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V207 := __args[0]
			_ = V207
			gen372 := Call(__e, ShenFunc(symshen_4_5pattern_6), V207)

			Parse__shen_4_5pattern_6 := gen372
			gen381 := Call(__e, ShenFunc(symfail))

			gen382 := Call(__e, ShenFunc(sym_a), gen381, Parse__shen_4_5pattern_6)

			gen383 := Call(__e, ShenFunc(symnot), gen382)

			var gen384 Obj
			if True == gen383 {
				gen373 := Call(__e, ShenFunc(symshen_4_5patterns_6), Parse__shen_4_5pattern_6)

				Parse__shen_4_5patterns_6 := gen373
				gen378 := Call(__e, ShenFunc(symfail))

				gen379 := Call(__e, ShenFunc(sym_a), gen378, Parse__shen_4_5patterns_6)

				gen380 := Call(__e, ShenFunc(symnot), gen379)

				if True == gen380 {
					gen374 := Call(__e, ShenFunc(symhd), Parse__shen_4_5patterns_6)

					gen375 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

					gen376 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5patterns_6)

					gen377 := Call(__e, ShenFunc(symcons), gen375, gen376)

					gen384 = Call(__e, ShenFunc(symshen_4pair), gen374, gen377)

				} else {
					gen384 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen384 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen384
			gen390 := Call(__e, ShenFunc(symfail))

			gen391 := Call(__e, ShenFunc(sym_a), YaccParse, gen390)

			if True == gen391 {
				gen385 := Call(__e, ShenFunc(sym_5e_6), V207)

				Parse___5e_6 := gen385
				gen387 := Call(__e, ShenFunc(symfail))

				gen388 := Call(__e, ShenFunc(sym_a), gen387, Parse___5e_6)

				gen389 := Call(__e, ShenFunc(symnot), gen388)

				if True == gen389 {
					gen386 := Call(__e, ShenFunc(symhd), Parse___5e_6)

					__e.TailApply(ShenFunc(symshen_4pair), gen386, Nil)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<patterns>"), gen392)

		gen634 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V220 := __args[0]
			_ = V220
			gen432 := Call(__e, ShenFunc(symhd), V220)

			gen433 := Call(__e, ShenFunc(symcons_2), gen432)

			var gen434 Obj
			if True == gen433 {
				gen430 := Call(__e, ShenFunc(symshen_4hdhd), V220)

				gen431 := Call(__e, ShenFunc(symcons_2), gen430)

				if True == gen431 {
					gen434 = True
				} else {
					gen434 = False
				}

			} else {
				gen434 = False
			}
			var gen435 Obj
			if True == gen434 {
				gen424 := Call(__e, ShenFunc(symshen_4hdhd), V220)

				gen425 := Call(__e, ShenFunc(symshen_4hdtl), V220)

				gen426 := Call(__e, ShenFunc(symshen_4pair), gen424, gen425)

				gen427 := Call(__e, ShenFunc(symhd), gen426)

				gen428 := Call(__e, ShenFunc(symcons_2), gen427)

				var gen429 Obj
				if True == gen428 {
					gen419 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen420 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen421 := Call(__e, ShenFunc(symshen_4pair), gen419, gen420)

					gen422 := Call(__e, ShenFunc(symshen_4hdhd), gen421)

					gen423 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), gen422)

					if True == gen423 {
						gen429 = True
					} else {
						gen429 = False
					}

				} else {
					gen429 = False
				}
				if True == gen429 {
					gen393 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen394 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen395 := Call(__e, ShenFunc(symshen_4pair), gen393, gen394)

					gen396 := Call(__e, ShenFunc(symshen_4tlhd), gen395)

					gen397 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen398 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen399 := Call(__e, ShenFunc(symshen_4pair), gen397, gen398)

					gen400 := Call(__e, ShenFunc(symshen_4hdtl), gen399)

					gen401 := Call(__e, ShenFunc(symshen_4pair), gen396, gen400)

					NewStream209 := gen401
					gen402 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream209)

					Parse__shen_4_5pattern1_6 := gen402
					gen416 := Call(__e, ShenFunc(symfail))

					gen417 := Call(__e, ShenFunc(sym_a), gen416, Parse__shen_4_5pattern1_6)

					gen418 := Call(__e, ShenFunc(symnot), gen417)

					if True == gen418 {
						gen403 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

						Parse__shen_4_5pattern2_6 := gen403
						gen413 := Call(__e, ShenFunc(symfail))

						gen414 := Call(__e, ShenFunc(sym_a), gen413, Parse__shen_4_5pattern2_6)

						gen415 := Call(__e, ShenFunc(symnot), gen414)

						if True == gen415 {
							gen404 := Call(__e, ShenFunc(symshen_4tlhd), V220)

							gen405 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen406 := Call(__e, ShenFunc(symshen_4pair), gen404, gen405)

							gen407 := Call(__e, ShenFunc(symhd), gen406)

							gen408 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

							gen409 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

							gen410 := Call(__e, ShenFunc(symcons), gen409, Nil)

							gen411 := Call(__e, ShenFunc(symcons), gen408, gen410)

							gen412 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen411)

							gen435 = Call(__e, ShenFunc(symshen_4pair), gen407, gen412)

						} else {
							gen435 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen435 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen435 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen435 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen435
			gen632 := Call(__e, ShenFunc(symfail))

			gen633 := Call(__e, ShenFunc(sym_a), YaccParse, gen632)

			if True == gen633 {
				gen475 := Call(__e, ShenFunc(symhd), V220)

				gen476 := Call(__e, ShenFunc(symcons_2), gen475)

				var gen477 Obj
				if True == gen476 {
					gen473 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen474 := Call(__e, ShenFunc(symcons_2), gen473)

					if True == gen474 {
						gen477 = True
					} else {
						gen477 = False
					}

				} else {
					gen477 = False
				}
				var gen478 Obj
				if True == gen477 {
					gen467 := Call(__e, ShenFunc(symshen_4hdhd), V220)

					gen468 := Call(__e, ShenFunc(symshen_4hdtl), V220)

					gen469 := Call(__e, ShenFunc(symshen_4pair), gen467, gen468)

					gen470 := Call(__e, ShenFunc(symhd), gen469)

					gen471 := Call(__e, ShenFunc(symcons_2), gen470)

					var gen472 Obj
					if True == gen471 {
						gen462 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen463 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen464 := Call(__e, ShenFunc(symshen_4pair), gen462, gen463)

						gen465 := Call(__e, ShenFunc(symshen_4hdhd), gen464)

						gen466 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen465)

						if True == gen466 {
							gen472 = True
						} else {
							gen472 = False
						}

					} else {
						gen472 = False
					}
					if True == gen472 {
						gen436 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen437 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen438 := Call(__e, ShenFunc(symshen_4pair), gen436, gen437)

						gen439 := Call(__e, ShenFunc(symshen_4tlhd), gen438)

						gen440 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen441 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen442 := Call(__e, ShenFunc(symshen_4pair), gen440, gen441)

						gen443 := Call(__e, ShenFunc(symshen_4hdtl), gen442)

						gen444 := Call(__e, ShenFunc(symshen_4pair), gen439, gen443)

						NewStream211 := gen444
						gen445 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream211)

						Parse__shen_4_5pattern1_6 := gen445
						gen459 := Call(__e, ShenFunc(symfail))

						gen460 := Call(__e, ShenFunc(sym_a), gen459, Parse__shen_4_5pattern1_6)

						gen461 := Call(__e, ShenFunc(symnot), gen460)

						if True == gen461 {
							gen446 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

							Parse__shen_4_5pattern2_6 := gen446
							gen456 := Call(__e, ShenFunc(symfail))

							gen457 := Call(__e, ShenFunc(sym_a), gen456, Parse__shen_4_5pattern2_6)

							gen458 := Call(__e, ShenFunc(symnot), gen457)

							if True == gen458 {
								gen447 := Call(__e, ShenFunc(symshen_4tlhd), V220)

								gen448 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen449 := Call(__e, ShenFunc(symshen_4pair), gen447, gen448)

								gen450 := Call(__e, ShenFunc(symhd), gen449)

								gen451 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

								gen452 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

								gen453 := Call(__e, ShenFunc(symcons), gen452, Nil)

								gen454 := Call(__e, ShenFunc(symcons), gen451, gen453)

								gen455 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen454)

								gen478 = Call(__e, ShenFunc(symshen_4pair), gen450, gen455)

							} else {
								gen478 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen478 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen478 = Call(__e, ShenFunc(symfail))

					}

				} else {
					gen478 = Call(__e, ShenFunc(symfail))

				}
				YaccParse := gen478
				gen630 := Call(__e, ShenFunc(symfail))

				gen631 := Call(__e, ShenFunc(sym_a), YaccParse, gen630)

				if True == gen631 {
					gen518 := Call(__e, ShenFunc(symhd), V220)

					gen519 := Call(__e, ShenFunc(symcons_2), gen518)

					var gen520 Obj
					if True == gen519 {
						gen516 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen517 := Call(__e, ShenFunc(symcons_2), gen516)

						if True == gen517 {
							gen520 = True
						} else {
							gen520 = False
						}

					} else {
						gen520 = False
					}
					var gen521 Obj
					if True == gen520 {
						gen510 := Call(__e, ShenFunc(symshen_4hdhd), V220)

						gen511 := Call(__e, ShenFunc(symshen_4hdtl), V220)

						gen512 := Call(__e, ShenFunc(symshen_4pair), gen510, gen511)

						gen513 := Call(__e, ShenFunc(symhd), gen512)

						gen514 := Call(__e, ShenFunc(symcons_2), gen513)

						var gen515 Obj
						if True == gen514 {
							gen505 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen506 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen507 := Call(__e, ShenFunc(symshen_4pair), gen505, gen506)

							gen508 := Call(__e, ShenFunc(symshen_4hdhd), gen507)

							gen509 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), gen508)

							if True == gen509 {
								gen515 = True
							} else {
								gen515 = False
							}

						} else {
							gen515 = False
						}
						if True == gen515 {
							gen479 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen480 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen481 := Call(__e, ShenFunc(symshen_4pair), gen479, gen480)

							gen482 := Call(__e, ShenFunc(symshen_4tlhd), gen481)

							gen483 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen484 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen485 := Call(__e, ShenFunc(symshen_4pair), gen483, gen484)

							gen486 := Call(__e, ShenFunc(symshen_4hdtl), gen485)

							gen487 := Call(__e, ShenFunc(symshen_4pair), gen482, gen486)

							NewStream213 := gen487
							gen488 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream213)

							Parse__shen_4_5pattern1_6 := gen488
							gen502 := Call(__e, ShenFunc(symfail))

							gen503 := Call(__e, ShenFunc(sym_a), gen502, Parse__shen_4_5pattern1_6)

							gen504 := Call(__e, ShenFunc(symnot), gen503)

							if True == gen504 {
								gen489 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

								Parse__shen_4_5pattern2_6 := gen489
								gen499 := Call(__e, ShenFunc(symfail))

								gen500 := Call(__e, ShenFunc(sym_a), gen499, Parse__shen_4_5pattern2_6)

								gen501 := Call(__e, ShenFunc(symnot), gen500)

								if True == gen501 {
									gen490 := Call(__e, ShenFunc(symshen_4tlhd), V220)

									gen491 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen492 := Call(__e, ShenFunc(symshen_4pair), gen490, gen491)

									gen493 := Call(__e, ShenFunc(symhd), gen492)

									gen494 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

									gen495 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

									gen496 := Call(__e, ShenFunc(symcons), gen495, Nil)

									gen497 := Call(__e, ShenFunc(symcons), gen494, gen496)

									gen498 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen497)

									gen521 = Call(__e, ShenFunc(symshen_4pair), gen493, gen498)

								} else {
									gen521 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen521 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen521 = Call(__e, ShenFunc(symfail))

						}

					} else {
						gen521 = Call(__e, ShenFunc(symfail))

					}
					YaccParse := gen521
					gen628 := Call(__e, ShenFunc(symfail))

					gen629 := Call(__e, ShenFunc(sym_a), YaccParse, gen628)

					if True == gen629 {
						gen561 := Call(__e, ShenFunc(symhd), V220)

						gen562 := Call(__e, ShenFunc(symcons_2), gen561)

						var gen563 Obj
						if True == gen562 {
							gen559 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen560 := Call(__e, ShenFunc(symcons_2), gen559)

							if True == gen560 {
								gen563 = True
							} else {
								gen563 = False
							}

						} else {
							gen563 = False
						}
						var gen564 Obj
						if True == gen563 {
							gen553 := Call(__e, ShenFunc(symshen_4hdhd), V220)

							gen554 := Call(__e, ShenFunc(symshen_4hdtl), V220)

							gen555 := Call(__e, ShenFunc(symshen_4pair), gen553, gen554)

							gen556 := Call(__e, ShenFunc(symhd), gen555)

							gen557 := Call(__e, ShenFunc(symcons_2), gen556)

							var gen558 Obj
							if True == gen557 {
								gen548 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen549 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen550 := Call(__e, ShenFunc(symshen_4pair), gen548, gen549)

								gen551 := Call(__e, ShenFunc(symshen_4hdhd), gen550)

								gen552 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen551)

								if True == gen552 {
									gen558 = True
								} else {
									gen558 = False
								}

							} else {
								gen558 = False
							}
							if True == gen558 {
								gen522 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen523 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen524 := Call(__e, ShenFunc(symshen_4pair), gen522, gen523)

								gen525 := Call(__e, ShenFunc(symshen_4tlhd), gen524)

								gen526 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen527 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen528 := Call(__e, ShenFunc(symshen_4pair), gen526, gen527)

								gen529 := Call(__e, ShenFunc(symshen_4hdtl), gen528)

								gen530 := Call(__e, ShenFunc(symshen_4pair), gen525, gen529)

								NewStream215 := gen530
								gen531 := Call(__e, ShenFunc(symshen_4_5pattern1_6), NewStream215)

								Parse__shen_4_5pattern1_6 := gen531
								gen545 := Call(__e, ShenFunc(symfail))

								gen546 := Call(__e, ShenFunc(sym_a), gen545, Parse__shen_4_5pattern1_6)

								gen547 := Call(__e, ShenFunc(symnot), gen546)

								if True == gen547 {
									gen532 := Call(__e, ShenFunc(symshen_4_5pattern2_6), Parse__shen_4_5pattern1_6)

									Parse__shen_4_5pattern2_6 := gen532
									gen542 := Call(__e, ShenFunc(symfail))

									gen543 := Call(__e, ShenFunc(sym_a), gen542, Parse__shen_4_5pattern2_6)

									gen544 := Call(__e, ShenFunc(symnot), gen543)

									if True == gen544 {
										gen533 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen534 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen535 := Call(__e, ShenFunc(symshen_4pair), gen533, gen534)

										gen536 := Call(__e, ShenFunc(symhd), gen535)

										gen537 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern1_6)

										gen538 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern2_6)

										gen539 := Call(__e, ShenFunc(symcons), gen538, Nil)

										gen540 := Call(__e, ShenFunc(symcons), gen537, gen539)

										gen541 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen540)

										gen564 = Call(__e, ShenFunc(symshen_4pair), gen536, gen541)

									} else {
										gen564 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen564 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen564 = Call(__e, ShenFunc(symfail))

							}

						} else {
							gen564 = Call(__e, ShenFunc(symfail))

						}
						YaccParse := gen564
						gen626 := Call(__e, ShenFunc(symfail))

						gen627 := Call(__e, ShenFunc(sym_a), YaccParse, gen626)

						if True == gen627 {
							gen601 := Call(__e, ShenFunc(symhd), V220)

							gen602 := Call(__e, ShenFunc(symcons_2), gen601)

							var gen603 Obj
							if True == gen602 {
								gen599 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen600 := Call(__e, ShenFunc(symcons_2), gen599)

								if True == gen600 {
									gen603 = True
								} else {
									gen603 = False
								}

							} else {
								gen603 = False
							}
							var gen604 Obj
							if True == gen603 {
								gen593 := Call(__e, ShenFunc(symshen_4hdhd), V220)

								gen594 := Call(__e, ShenFunc(symshen_4hdtl), V220)

								gen595 := Call(__e, ShenFunc(symshen_4pair), gen593, gen594)

								gen596 := Call(__e, ShenFunc(symhd), gen595)

								gen597 := Call(__e, ShenFunc(symcons_2), gen596)

								var gen598 Obj
								if True == gen597 {
									gen588 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen589 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen590 := Call(__e, ShenFunc(symshen_4pair), gen588, gen589)

									gen591 := Call(__e, ShenFunc(symshen_4hdhd), gen590)

									gen592 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), gen591)

									if True == gen592 {
										gen598 = True
									} else {
										gen598 = False
									}

								} else {
									gen598 = False
								}
								if True == gen598 {
									gen565 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen566 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen567 := Call(__e, ShenFunc(symshen_4pair), gen565, gen566)

									gen568 := Call(__e, ShenFunc(symshen_4tlhd), gen567)

									gen569 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									gen570 := Call(__e, ShenFunc(symshen_4hdtl), V220)

									gen571 := Call(__e, ShenFunc(symshen_4pair), gen569, gen570)

									gen572 := Call(__e, ShenFunc(symshen_4hdtl), gen571)

									gen573 := Call(__e, ShenFunc(symshen_4pair), gen568, gen572)

									NewStream217 := gen573
									gen585 := Call(__e, ShenFunc(symhd), NewStream217)

									gen586 := Call(__e, ShenFunc(symcons_2), gen585)

									var gen587 Obj
									if True == gen586 {
										gen583 := Call(__e, ShenFunc(symshen_4hdhd), NewStream217)

										gen584 := Call(__e, ShenFunc(sym_a), MakeNumber(0), gen583)

										if True == gen584 {
											gen587 = True
										} else {
											gen587 = False
										}

									} else {
										gen587 = False
									}
									if True == gen587 {
										gen574 := Call(__e, ShenFunc(symshen_4tlhd), NewStream217)

										gen575 := Call(__e, ShenFunc(symshen_4hdtl), NewStream217)

										gen576 := Call(__e, ShenFunc(symshen_4pair), gen574, gen575)

										NewStream218 := gen576
										_ = NewStream218
										gen577 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen578 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen579 := Call(__e, ShenFunc(symshen_4pair), gen577, gen578)

										gen580 := Call(__e, ShenFunc(symhd), gen579)

										gen581 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

										gen582 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen581)

										gen604 = Call(__e, ShenFunc(symshen_4pair), gen580, gen582)

									} else {
										gen604 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen604 = Call(__e, ShenFunc(symfail))

								}

							} else {
								gen604 = Call(__e, ShenFunc(symfail))

							}
							YaccParse := gen604
							gen624 := Call(__e, ShenFunc(symfail))

							gen625 := Call(__e, ShenFunc(sym_a), YaccParse, gen624)

							if True == gen625 {
								gen613 := Call(__e, ShenFunc(symhd), V220)

								gen614 := Call(__e, ShenFunc(symcons_2), gen613)

								var gen615 Obj
								if True == gen614 {
									gen605 := Call(__e, ShenFunc(symshen_4hdhd), V220)

									Parse__X := gen605
									gen612 := Call(__e, ShenFunc(symcons_2), Parse__X)

									if True == gen612 {
										gen606 := Call(__e, ShenFunc(symshen_4tlhd), V220)

										gen607 := Call(__e, ShenFunc(symshen_4hdtl), V220)

										gen608 := Call(__e, ShenFunc(symshen_4pair), gen606, gen607)

										gen609 := Call(__e, ShenFunc(symhd), gen608)

										gen610 := MakeNative(func(__e Evaluator, __args ...Obj) {
											__e.TailApply(ShenFunc(symshen_4constructor_1error), Parse__X)

											return
										}, 0)
										gen611 := Call(__e, ShenFunc(symshen_4custom_1pattern_1compiler), Parse__X, gen610)

										gen615 = Call(__e, ShenFunc(symshen_4pair), gen609, gen611)

									} else {
										gen615 = Call(__e, ShenFunc(symfail))

									}

								} else {
									gen615 = Call(__e, ShenFunc(symfail))

								}
								YaccParse := gen615
								gen622 := Call(__e, ShenFunc(symfail))

								gen623 := Call(__e, ShenFunc(sym_a), YaccParse, gen622)

								if True == gen623 {
									gen616 := Call(__e, ShenFunc(symshen_4_5simple__pattern_6), V220)

									Parse__shen_4_5simple__pattern_6 := gen616
									gen619 := Call(__e, ShenFunc(symfail))

									gen620 := Call(__e, ShenFunc(sym_a), gen619, Parse__shen_4_5simple__pattern_6)

									gen621 := Call(__e, ShenFunc(symnot), gen620)

									if True == gen621 {
										gen617 := Call(__e, ShenFunc(symhd), Parse__shen_4_5simple__pattern_6)

										gen618 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5simple__pattern_6)

										__e.TailApply(ShenFunc(symshen_4pair), gen617, gen618)

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

						} else {
							__e.Return(YaccParse)
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

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern>"), gen634)

		gen636 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V222 := __args[0]
			_ = V222
			gen635 := Call(__e, ShenFunc(symshen_4app), V222, MakeString(" is not a legitimate constructor\n"), MakeSymbol("shen.a"))

			__e.TailApply(ShenFunc(symsimple_1error), gen635)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.constructor-error"), gen636)

		gen660 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V224 := __args[0]
			_ = V224
			gen644 := Call(__e, ShenFunc(symhd), V224)

			gen645 := Call(__e, ShenFunc(symcons_2), gen644)

			var gen646 Obj
			if True == gen645 {
				gen637 := Call(__e, ShenFunc(symshen_4hdhd), V224)

				Parse__X := gen637
				gen643 := Call(__e, ShenFunc(sym_a), Parse__X, MakeSymbol("_"))

				if True == gen643 {
					gen638 := Call(__e, ShenFunc(symshen_4tlhd), V224)

					gen639 := Call(__e, ShenFunc(symshen_4hdtl), V224)

					gen640 := Call(__e, ShenFunc(symshen_4pair), gen638, gen639)

					gen641 := Call(__e, ShenFunc(symhd), gen640)

					gen642 := Call(__e, ShenFunc(symgensym), MakeSymbol("Parse_Y"))

					gen646 = Call(__e, ShenFunc(symshen_4pair), gen641, gen642)

				} else {
					gen646 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen646 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen646
			gen658 := Call(__e, ShenFunc(symfail))

			gen659 := Call(__e, ShenFunc(sym_a), YaccParse, gen658)

			if True == gen659 {
				gen656 := Call(__e, ShenFunc(symhd), V224)

				gen657 := Call(__e, ShenFunc(symcons_2), gen656)

				if True == gen657 {
					gen647 := Call(__e, ShenFunc(symshen_4hdhd), V224)

					Parse__X := gen647
					gen652 := Call(__e, ShenFunc(symcons), MakeSymbol("<-"), Nil)

					gen653 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen652)

					gen654 := Call(__e, ShenFunc(symelement_2), Parse__X, gen653)

					gen655 := Call(__e, ShenFunc(symnot), gen654)

					if True == gen655 {
						gen648 := Call(__e, ShenFunc(symshen_4tlhd), V224)

						gen649 := Call(__e, ShenFunc(symshen_4hdtl), V224)

						gen650 := Call(__e, ShenFunc(symshen_4pair), gen648, gen649)

						gen651 := Call(__e, ShenFunc(symhd), gen650)

						__e.TailApply(ShenFunc(symshen_4pair), gen651, Parse__X)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<simple_pattern>"), gen660)

		gen667 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V226 := __args[0]
			_ = V226
			gen661 := Call(__e, ShenFunc(symshen_4_5pattern_6), V226)

			Parse__shen_4_5pattern_6 := gen661
			gen664 := Call(__e, ShenFunc(symfail))

			gen665 := Call(__e, ShenFunc(sym_a), gen664, Parse__shen_4_5pattern_6)

			gen666 := Call(__e, ShenFunc(symnot), gen665)

			if True == gen666 {
				gen662 := Call(__e, ShenFunc(symhd), Parse__shen_4_5pattern_6)

				gen663 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen662, gen663)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern1>"), gen667)

		gen674 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V228 := __args[0]
			_ = V228
			gen668 := Call(__e, ShenFunc(symshen_4_5pattern_6), V228)

			Parse__shen_4_5pattern_6 := gen668
			gen671 := Call(__e, ShenFunc(symfail))

			gen672 := Call(__e, ShenFunc(sym_a), gen671, Parse__shen_4_5pattern_6)

			gen673 := Call(__e, ShenFunc(symnot), gen672)

			if True == gen673 {
				gen669 := Call(__e, ShenFunc(symhd), Parse__shen_4_5pattern_6)

				gen670 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5pattern_6)

				__e.TailApply(ShenFunc(symshen_4pair), gen669, gen670)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<pattern2>"), gen674)

		gen682 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V230 := __args[0]
			_ = V230
			gen680 := Call(__e, ShenFunc(symhd), V230)

			gen681 := Call(__e, ShenFunc(symcons_2), gen680)

			if True == gen681 {
				gen675 := Call(__e, ShenFunc(symshen_4hdhd), V230)

				Parse__X := gen675
				gen676 := Call(__e, ShenFunc(symshen_4tlhd), V230)

				gen677 := Call(__e, ShenFunc(symshen_4hdtl), V230)

				gen678 := Call(__e, ShenFunc(symshen_4pair), gen676, gen677)

				gen679 := Call(__e, ShenFunc(symhd), gen678)

				__e.TailApply(ShenFunc(symshen_4pair), gen679, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<action>"), gen682)

		gen690 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V232 := __args[0]
			_ = V232
			gen688 := Call(__e, ShenFunc(symhd), V232)

			gen689 := Call(__e, ShenFunc(symcons_2), gen688)

			if True == gen689 {
				gen683 := Call(__e, ShenFunc(symshen_4hdhd), V232)

				Parse__X := gen683
				gen684 := Call(__e, ShenFunc(symshen_4tlhd), V232)

				gen685 := Call(__e, ShenFunc(symshen_4hdtl), V232)

				gen686 := Call(__e, ShenFunc(symshen_4pair), gen684, gen685)

				gen687 := Call(__e, ShenFunc(symhd), gen686)

				__e.TailApply(ShenFunc(symshen_4pair), gen687, Parse__X)

				return

			} else {
				__e.TailApply(ShenFunc(symfail))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<guard>"), gen690)

		gen694 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V235 := __args[0]
			_ = V235
			V236 := __args[1]
			_ = V236
			gen691 := Call(__e, ShenFunc(symshen_4compile__to__lambda_7), V235, V236)

			Lambda_7 := gen691
			gen692 := Call(__e, ShenFunc(symshen_4compile__to__kl), V235, Lambda_7)

			KL := gen692
			gen693 := Call(__e, ShenFunc(symshen_4record_1source), V235, KL)

			Record := gen693
			_ = Record
			__e.Return(KL)
			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_machine_code"), gen694)

		gen697 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V241 := __args[0]
			_ = V241
			V242 := __args[1]
			_ = V242
			gen696 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

			if True == gen696 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen695 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V241, MakeSymbol("shen.source"), V242, gen695)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.record-source"), gen697)

		gen710 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V245 := __args[0]
			_ = V245
			V246 := __args[1]
			_ = V246
			gen698 := Call(__e, ShenFunc(symshen_4aritycheck), V245, V246)

			Arity := gen698
			_ = Arity
			gen699 := Call(__e, ShenFunc(symshen_4update_1symbol_1table), V245, Arity)

			UpDateSymbolTable := gen699
			_ = UpDateSymbolTable
			gen700 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Rule := __args[0]
				_ = Rule
				__e.TailApply(ShenFunc(symshen_4free__variable__check), V245, Rule)

				return
			}, 1)
			gen701 := Call(__e, ShenFunc(symshen_4for_1each), gen700, V246)

			Free := gen701
			_ = Free
			gen702 := Call(__e, ShenFunc(symshen_4parameters), Arity)

			Variables := gen702
			gen703 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4strip_1protect), X)

				return
			}, 1)
			gen704 := Call(__e, ShenFunc(symmap), gen703, V246)

			Strip := gen704
			gen705 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4abstract__rule), X)

				return
			}, 1)
			gen706 := Call(__e, ShenFunc(symmap), gen705, Strip)

			Abstractions := gen706
			gen707 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4application__build), Variables, X)

				return
			}, 1)
			gen708 := Call(__e, ShenFunc(symmap), gen707, Abstractions)

			Applications := gen708
			gen709 := Call(__e, ShenFunc(symcons), Applications, Nil)

			__e.TailApply(ShenFunc(symcons), Variables, gen709)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_lambda+"), gen710)

		gen715 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V249 := __args[0]
			_ = V249
			V250 := __args[1]
			_ = V250
			gen714 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V250)

			if True == gen714 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen711 := Call(__e, ShenFunc(symshen_4lambda_1form), V249, V250)

				gen712 := Call(__e, ShenFunc(symeval_1kl), gen711)

				gen713 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V249, MakeSymbol("shen.lambda-form"), gen712, gen713)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.update-symbol-table"), gen715)

		gen729 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V253 := __args[0]
			_ = V253
			V254 := __args[1]
			_ = V254
			gen727 := Call(__e, ShenFunc(symcons_2), V254)

			var gen728 Obj
			if True == gen727 {
				gen724 := Call(__e, ShenFunc(symtl), V254)

				gen725 := Call(__e, ShenFunc(symcons_2), gen724)

				var gen726 Obj
				if True == gen725 {
					gen721 := Call(__e, ShenFunc(symtl), V254)

					gen722 := Call(__e, ShenFunc(symtl), gen721)

					gen723 := Call(__e, ShenFunc(sym_a), Nil, gen722)

					if True == gen723 {
						gen726 = True
					} else {
						gen726 = False
					}

				} else {
					gen726 = False
				}
				if True == gen726 {
					gen728 = True
				} else {
					gen728 = False
				}

			} else {
				gen728 = False
			}
			if True == gen728 {
				gen716 := Call(__e, ShenFunc(symhd), V254)

				gen717 := Call(__e, ShenFunc(symshen_4extract__vars), gen716)

				Bound := gen717
				gen718 := Call(__e, ShenFunc(symtl), V254)

				gen719 := Call(__e, ShenFunc(symhd), gen718)

				gen720 := Call(__e, ShenFunc(symshen_4extract__free__vars), Bound, gen719)

				Free := gen720
				__e.TailApply(ShenFunc(symshen_4free__variable__warnings), V253, Free)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.free_variable_check"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.free_variable_check"), gen729)

		gen736 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V256 := __args[0]
			_ = V256
			gen735 := Call(__e, ShenFunc(symvariable_2), V256)

			if True == gen735 {
				__e.TailApply(ShenFunc(symcons), V256, Nil)

				return
			} else {
				gen734 := Call(__e, ShenFunc(symcons_2), V256)

				if True == gen734 {
					gen730 := Call(__e, ShenFunc(symhd), V256)

					gen731 := Call(__e, ShenFunc(symshen_4extract__vars), gen730)

					gen732 := Call(__e, ShenFunc(symtl), V256)

					gen733 := Call(__e, ShenFunc(symshen_4extract__vars), gen732)

					__e.TailApply(ShenFunc(symunion), gen731, gen733)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extract_vars"), gen736)

		gen813 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V268 := __args[0]
			_ = V268
			V269 := __args[1]
			_ = V269
			gen811 := Call(__e, ShenFunc(symcons_2), V269)

			var gen812 Obj
			if True == gen811 {
				gen808 := Call(__e, ShenFunc(symtl), V269)

				gen809 := Call(__e, ShenFunc(symcons_2), gen808)

				var gen810 Obj
				if True == gen809 {
					gen804 := Call(__e, ShenFunc(symtl), V269)

					gen805 := Call(__e, ShenFunc(symtl), gen804)

					gen806 := Call(__e, ShenFunc(sym_a), Nil, gen805)

					var gen807 Obj
					if True == gen806 {
						gen802 := Call(__e, ShenFunc(symhd), V269)

						gen803 := Call(__e, ShenFunc(sym_a), gen802, MakeSymbol("protect"))

						if True == gen803 {
							gen807 = True
						} else {
							gen807 = False
						}

					} else {
						gen807 = False
					}
					if True == gen807 {
						gen810 = True
					} else {
						gen810 = False
					}

				} else {
					gen810 = False
				}
				if True == gen810 {
					gen812 = True
				} else {
					gen812 = False
				}

			} else {
				gen812 = False
			}
			if True == gen812 {
				__e.Return(Nil)
				return
			} else {
				gen800 := Call(__e, ShenFunc(symvariable_2), V269)

				var gen801 Obj
				if True == gen800 {
					gen798 := Call(__e, ShenFunc(symelement_2), V269, V268)

					gen799 := Call(__e, ShenFunc(symnot), gen798)

					if True == gen799 {
						gen801 = True
					} else {
						gen801 = False
					}

				} else {
					gen801 = False
				}
				if True == gen801 {
					__e.TailApply(ShenFunc(symcons), V269, Nil)

					return
				} else {
					gen796 := Call(__e, ShenFunc(symcons_2), V269)

					var gen797 Obj
					if True == gen796 {
						gen793 := Call(__e, ShenFunc(symhd), V269)

						gen794 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen793)

						var gen795 Obj
						if True == gen794 {
							gen790 := Call(__e, ShenFunc(symtl), V269)

							gen791 := Call(__e, ShenFunc(symcons_2), gen790)

							var gen792 Obj
							if True == gen791 {
								gen786 := Call(__e, ShenFunc(symtl), V269)

								gen787 := Call(__e, ShenFunc(symtl), gen786)

								gen788 := Call(__e, ShenFunc(symcons_2), gen787)

								var gen789 Obj
								if True == gen788 {
									gen782 := Call(__e, ShenFunc(symtl), V269)

									gen783 := Call(__e, ShenFunc(symtl), gen782)

									gen784 := Call(__e, ShenFunc(symtl), gen783)

									gen785 := Call(__e, ShenFunc(sym_a), Nil, gen784)

									if True == gen785 {
										gen789 = True
									} else {
										gen789 = False
									}

								} else {
									gen789 = False
								}
								if True == gen789 {
									gen792 = True
								} else {
									gen792 = False
								}

							} else {
								gen792 = False
							}
							if True == gen792 {
								gen795 = True
							} else {
								gen795 = False
							}

						} else {
							gen795 = False
						}
						if True == gen795 {
							gen797 = True
						} else {
							gen797 = False
						}

					} else {
						gen797 = False
					}
					if True == gen797 {
						gen776 := Call(__e, ShenFunc(symtl), V269)

						gen777 := Call(__e, ShenFunc(symhd), gen776)

						gen778 := Call(__e, ShenFunc(symcons), gen777, V268)

						gen779 := Call(__e, ShenFunc(symtl), V269)

						gen780 := Call(__e, ShenFunc(symtl), gen779)

						gen781 := Call(__e, ShenFunc(symhd), gen780)

						__e.TailApply(ShenFunc(symshen_4extract__free__vars), gen778, gen781)

						return

					} else {
						gen774 := Call(__e, ShenFunc(symcons_2), V269)

						var gen775 Obj
						if True == gen774 {
							gen771 := Call(__e, ShenFunc(symhd), V269)

							gen772 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen771)

							var gen773 Obj
							if True == gen772 {
								gen768 := Call(__e, ShenFunc(symtl), V269)

								gen769 := Call(__e, ShenFunc(symcons_2), gen768)

								var gen770 Obj
								if True == gen769 {
									gen764 := Call(__e, ShenFunc(symtl), V269)

									gen765 := Call(__e, ShenFunc(symtl), gen764)

									gen766 := Call(__e, ShenFunc(symcons_2), gen765)

									var gen767 Obj
									if True == gen766 {
										gen759 := Call(__e, ShenFunc(symtl), V269)

										gen760 := Call(__e, ShenFunc(symtl), gen759)

										gen761 := Call(__e, ShenFunc(symtl), gen760)

										gen762 := Call(__e, ShenFunc(symcons_2), gen761)

										var gen763 Obj
										if True == gen762 {
											gen754 := Call(__e, ShenFunc(symtl), V269)

											gen755 := Call(__e, ShenFunc(symtl), gen754)

											gen756 := Call(__e, ShenFunc(symtl), gen755)

											gen757 := Call(__e, ShenFunc(symtl), gen756)

											gen758 := Call(__e, ShenFunc(sym_a), Nil, gen757)

											if True == gen758 {
												gen763 = True
											} else {
												gen763 = False
											}

										} else {
											gen763 = False
										}
										if True == gen763 {
											gen767 = True
										} else {
											gen767 = False
										}

									} else {
										gen767 = False
									}
									if True == gen767 {
										gen770 = True
									} else {
										gen770 = False
									}

								} else {
									gen770 = False
								}
								if True == gen770 {
									gen773 = True
								} else {
									gen773 = False
								}

							} else {
								gen773 = False
							}
							if True == gen773 {
								gen775 = True
							} else {
								gen775 = False
							}

						} else {
							gen775 = False
						}
						if True == gen775 {
							gen742 := Call(__e, ShenFunc(symtl), V269)

							gen743 := Call(__e, ShenFunc(symtl), gen742)

							gen744 := Call(__e, ShenFunc(symhd), gen743)

							gen745 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen744)

							gen746 := Call(__e, ShenFunc(symtl), V269)

							gen747 := Call(__e, ShenFunc(symhd), gen746)

							gen748 := Call(__e, ShenFunc(symcons), gen747, V268)

							gen749 := Call(__e, ShenFunc(symtl), V269)

							gen750 := Call(__e, ShenFunc(symtl), gen749)

							gen751 := Call(__e, ShenFunc(symtl), gen750)

							gen752 := Call(__e, ShenFunc(symhd), gen751)

							gen753 := Call(__e, ShenFunc(symshen_4extract__free__vars), gen748, gen752)

							__e.TailApply(ShenFunc(symunion), gen745, gen753)

							return

						} else {
							gen741 := Call(__e, ShenFunc(symcons_2), V269)

							if True == gen741 {
								gen737 := Call(__e, ShenFunc(symhd), V269)

								gen738 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen737)

								gen739 := Call(__e, ShenFunc(symtl), V269)

								gen740 := Call(__e, ShenFunc(symshen_4extract__free__vars), V268, gen739)

								__e.TailApply(ShenFunc(symunion), gen738, gen740)

								return

							} else {
								__e.Return(Nil)
								return
							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extract_free_vars"), gen813)

		gen820 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V274 := __args[0]
			_ = V274
			V275 := __args[1]
			_ = V275
			gen819 := Call(__e, ShenFunc(sym_a), Nil, V275)

			if True == gen819 {
				__e.Return(MakeSymbol("_"))
				return
			} else {
				gen814 := Call(__e, ShenFunc(symshen_4list__variables), V275)

				gen815 := Call(__e, ShenFunc(symshen_4app), gen814, MakeString(""), MakeSymbol("shen.a"))

				gen816 := Call(__e, ShenFunc(symcn), MakeString(": "), gen815)

				gen817 := Call(__e, ShenFunc(symshen_4app), V274, gen816, MakeSymbol("shen.a"))

				gen818 := Call(__e, ShenFunc(symcn), MakeString("error: the following variables are free in "), gen817)

				__e.TailApply(ShenFunc(symsimple_1error), gen818)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.free_variable_warnings"), gen820)

		gen833 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V277 := __args[0]
			_ = V277
			gen831 := Call(__e, ShenFunc(symcons_2), V277)

			var gen832 Obj
			if True == gen831 {
				gen829 := Call(__e, ShenFunc(symtl), V277)

				gen830 := Call(__e, ShenFunc(sym_a), Nil, gen829)

				if True == gen830 {
					gen832 = True
				} else {
					gen832 = False
				}

			} else {
				gen832 = False
			}
			if True == gen832 {
				gen827 := Call(__e, ShenFunc(symhd), V277)

				gen828 := Call(__e, ShenFunc(symstr), gen827)

				__e.TailApply(ShenFunc(symcn), gen828, MakeString("."))

				return

			} else {
				gen826 := Call(__e, ShenFunc(symcons_2), V277)

				if True == gen826 {
					gen821 := Call(__e, ShenFunc(symhd), V277)

					gen822 := Call(__e, ShenFunc(symstr), gen821)

					gen823 := Call(__e, ShenFunc(symtl), V277)

					gen824 := Call(__e, ShenFunc(symshen_4list__variables), gen823)

					gen825 := Call(__e, ShenFunc(symcn), MakeString(", "), gen824)

					__e.TailApply(ShenFunc(symcn), gen822, gen825)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.list_variables"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.list_variables"), gen833)

		gen849 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V279 := __args[0]
			_ = V279
			gen847 := Call(__e, ShenFunc(symcons_2), V279)

			var gen848 Obj
			if True == gen847 {
				gen844 := Call(__e, ShenFunc(symtl), V279)

				gen845 := Call(__e, ShenFunc(symcons_2), gen844)

				var gen846 Obj
				if True == gen845 {
					gen840 := Call(__e, ShenFunc(symtl), V279)

					gen841 := Call(__e, ShenFunc(symtl), gen840)

					gen842 := Call(__e, ShenFunc(sym_a), Nil, gen841)

					var gen843 Obj
					if True == gen842 {
						gen838 := Call(__e, ShenFunc(symhd), V279)

						gen839 := Call(__e, ShenFunc(sym_a), gen838, MakeSymbol("protect"))

						if True == gen839 {
							gen843 = True
						} else {
							gen843 = False
						}

					} else {
						gen843 = False
					}
					if True == gen843 {
						gen846 = True
					} else {
						gen846 = False
					}

				} else {
					gen846 = False
				}
				if True == gen846 {
					gen848 = True
				} else {
					gen848 = False
				}

			} else {
				gen848 = False
			}
			if True == gen848 {
				gen836 := Call(__e, ShenFunc(symtl), V279)

				gen837 := Call(__e, ShenFunc(symhd), gen836)

				__e.TailApply(ShenFunc(symshen_4strip_1protect), gen837)

				return

			} else {
				gen835 := Call(__e, ShenFunc(symcons_2), V279)

				if True == gen835 {
					gen834 := MakeNative(func(__e Evaluator, __args ...Obj) {
						Z := __args[0]
						_ = Z
						__e.TailApply(ShenFunc(symshen_4strip_1protect), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen834, V279)

					return

				} else {
					__e.Return(V279)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.strip-protect"), gen849)

		gen863 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V281 := __args[0]
			_ = V281
			gen861 := Call(__e, ShenFunc(symcons_2), V281)

			var gen862 Obj
			if True == gen861 {
				gen858 := Call(__e, ShenFunc(symtl), V281)

				gen859 := Call(__e, ShenFunc(symcons_2), gen858)

				var gen860 Obj
				if True == gen859 {
					gen855 := Call(__e, ShenFunc(symtl), V281)

					gen856 := Call(__e, ShenFunc(symtl), gen855)

					gen857 := Call(__e, ShenFunc(sym_a), Nil, gen856)

					if True == gen857 {
						gen860 = True
					} else {
						gen860 = False
					}

				} else {
					gen860 = False
				}
				if True == gen860 {
					gen862 = True
				} else {
					gen862 = False
				}

			} else {
				gen862 = False
			}
			if True == gen862 {
				gen850 := Call(__e, ShenFunc(symhd), V281)

				gen851 := Call(__e, ShenFunc(symshen_4flatten), gen850)

				gen852 := Call(__e, ShenFunc(symhd), V281)

				gen853 := Call(__e, ShenFunc(symtl), V281)

				gen854 := Call(__e, ShenFunc(symhd), gen853)

				__e.TailApply(ShenFunc(symshen_4linearise__help), gen851, gen852, gen854)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise"), gen863)

		gen870 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V283 := __args[0]
			_ = V283
			gen869 := Call(__e, ShenFunc(sym_a), Nil, V283)

			if True == gen869 {
				__e.Return(Nil)
				return
			} else {
				gen868 := Call(__e, ShenFunc(symcons_2), V283)

				if True == gen868 {
					gen864 := Call(__e, ShenFunc(symhd), V283)

					gen865 := Call(__e, ShenFunc(symshen_4flatten), gen864)

					gen866 := Call(__e, ShenFunc(symtl), V283)

					gen867 := Call(__e, ShenFunc(symshen_4flatten), gen866)

					__e.TailApply(ShenFunc(symappend), gen865, gen867)

					return

				} else {
					__e.TailApply(ShenFunc(symcons), V283, Nil)

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.flatten"), gen870)

		gen893 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V287 := __args[0]
			_ = V287
			V288 := __args[1]
			_ = V288
			V289 := __args[2]
			_ = V289
			gen892 := Call(__e, ShenFunc(sym_a), Nil, V287)

			if True == gen892 {
				gen891 := Call(__e, ShenFunc(symcons), V289, Nil)

				__e.TailApply(ShenFunc(symcons), V288, gen891)

				return

			} else {
				gen890 := Call(__e, ShenFunc(symcons_2), V287)

				if True == gen890 {
					gen887 := Call(__e, ShenFunc(symhd), V287)

					gen888 := Call(__e, ShenFunc(symvariable_2), gen887)

					var gen889 Obj
					if True == gen888 {
						gen884 := Call(__e, ShenFunc(symhd), V287)

						gen885 := Call(__e, ShenFunc(symtl), V287)

						gen886 := Call(__e, ShenFunc(symelement_2), gen884, gen885)

						if True == gen886 {
							gen889 = True
						} else {
							gen889 = False
						}

					} else {
						gen889 = False
					}
					if True == gen889 {
						gen872 := Call(__e, ShenFunc(symhd), V287)

						gen873 := Call(__e, ShenFunc(symgensym), gen872)

						Var := gen873
						gen874 := Call(__e, ShenFunc(symhd), V287)

						gen875 := Call(__e, ShenFunc(symcons), Var, Nil)

						gen876 := Call(__e, ShenFunc(symcons), gen874, gen875)

						gen877 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen876)

						gen878 := Call(__e, ShenFunc(symcons), V289, Nil)

						gen879 := Call(__e, ShenFunc(symcons), gen877, gen878)

						gen880 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen879)

						NewAction := gen880
						gen881 := Call(__e, ShenFunc(symhd), V287)

						gen882 := Call(__e, ShenFunc(symshen_4linearise__X), gen881, Var, V288)

						NewPatts := gen882
						gen883 := Call(__e, ShenFunc(symtl), V287)

						__e.TailApply(ShenFunc(symshen_4linearise__help), gen883, NewPatts, NewAction)

						return

					} else {
						gen871 := Call(__e, ShenFunc(symtl), V287)

						__e.TailApply(ShenFunc(symshen_4linearise__help), gen871, V288, V289)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.linearise_help"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise_help"), gen893)

		gen904 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V302 := __args[0]
			_ = V302
			V303 := __args[1]
			_ = V303
			V304 := __args[2]
			_ = V304
			gen903 := Call(__e, ShenFunc(sym_a), V304, V302)

			if True == gen903 {
				__e.Return(V303)
				return
			} else {
				gen902 := Call(__e, ShenFunc(symcons_2), V304)

				if True == gen902 {
					gen894 := Call(__e, ShenFunc(symhd), V304)

					gen895 := Call(__e, ShenFunc(symshen_4linearise__X), V302, V303, gen894)

					L := gen895
					gen900 := Call(__e, ShenFunc(symhd), V304)

					gen901 := Call(__e, ShenFunc(sym_a), L, gen900)

					if True == gen901 {
						gen897 := Call(__e, ShenFunc(symhd), V304)

						gen898 := Call(__e, ShenFunc(symtl), V304)

						gen899 := Call(__e, ShenFunc(symshen_4linearise__X), V302, V303, gen898)

						__e.TailApply(ShenFunc(symcons), gen897, gen899)

						return

					} else {
						gen896 := Call(__e, ShenFunc(symtl), V304)

						__e.TailApply(ShenFunc(symcons), L, gen896)

						return

					}

				} else {
					__e.Return(V304)
					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.linearise_X"), gen904)

		gen973 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V307 := __args[0]
			_ = V307
			V308 := __args[1]
			_ = V308
			gen971 := Call(__e, ShenFunc(symcons_2), V308)

			var gen972 Obj
			if True == gen971 {
				gen968 := Call(__e, ShenFunc(symhd), V308)

				gen969 := Call(__e, ShenFunc(symcons_2), gen968)

				var gen970 Obj
				if True == gen969 {
					gen964 := Call(__e, ShenFunc(symhd), V308)

					gen965 := Call(__e, ShenFunc(symtl), gen964)

					gen966 := Call(__e, ShenFunc(symcons_2), gen965)

					var gen967 Obj
					if True == gen966 {
						gen959 := Call(__e, ShenFunc(symhd), V308)

						gen960 := Call(__e, ShenFunc(symtl), gen959)

						gen961 := Call(__e, ShenFunc(symtl), gen960)

						gen962 := Call(__e, ShenFunc(sym_a), Nil, gen961)

						var gen963 Obj
						if True == gen962 {
							gen957 := Call(__e, ShenFunc(symtl), V308)

							gen958 := Call(__e, ShenFunc(sym_a), Nil, gen957)

							if True == gen958 {
								gen963 = True
							} else {
								gen963 = False
							}

						} else {
							gen963 = False
						}
						if True == gen963 {
							gen967 = True
						} else {
							gen967 = False
						}

					} else {
						gen967 = False
					}
					if True == gen967 {
						gen970 = True
					} else {
						gen970 = False
					}

				} else {
					gen970 = False
				}
				if True == gen970 {
					gen972 = True
				} else {
					gen972 = False
				}

			} else {
				gen972 = False
			}
			if True == gen972 {
				gen950 := Call(__e, ShenFunc(symhd), V308)

				gen951 := Call(__e, ShenFunc(symtl), gen950)

				gen952 := Call(__e, ShenFunc(symhd), gen951)

				Call(__e, ShenFunc(symshen_4aritycheck_1action), gen952)

				gen953 := Call(__e, ShenFunc(symarity), V307)

				gen954 := Call(__e, ShenFunc(symhd), V308)

				gen955 := Call(__e, ShenFunc(symhd), gen954)

				gen956 := Call(__e, ShenFunc(symlength), gen955)

				__e.TailApply(ShenFunc(symshen_4aritycheck_1name), V307, gen953, gen956)

				return

			} else {
				gen948 := Call(__e, ShenFunc(symcons_2), V308)

				var gen949 Obj
				if True == gen948 {
					gen945 := Call(__e, ShenFunc(symhd), V308)

					gen946 := Call(__e, ShenFunc(symcons_2), gen945)

					var gen947 Obj
					if True == gen946 {
						gen941 := Call(__e, ShenFunc(symhd), V308)

						gen942 := Call(__e, ShenFunc(symtl), gen941)

						gen943 := Call(__e, ShenFunc(symcons_2), gen942)

						var gen944 Obj
						if True == gen943 {
							gen936 := Call(__e, ShenFunc(symhd), V308)

							gen937 := Call(__e, ShenFunc(symtl), gen936)

							gen938 := Call(__e, ShenFunc(symtl), gen937)

							gen939 := Call(__e, ShenFunc(sym_a), Nil, gen938)

							var gen940 Obj
							if True == gen939 {
								gen933 := Call(__e, ShenFunc(symtl), V308)

								gen934 := Call(__e, ShenFunc(symcons_2), gen933)

								var gen935 Obj
								if True == gen934 {
									gen929 := Call(__e, ShenFunc(symtl), V308)

									gen930 := Call(__e, ShenFunc(symhd), gen929)

									gen931 := Call(__e, ShenFunc(symcons_2), gen930)

									var gen932 Obj
									if True == gen931 {
										gen924 := Call(__e, ShenFunc(symtl), V308)

										gen925 := Call(__e, ShenFunc(symhd), gen924)

										gen926 := Call(__e, ShenFunc(symtl), gen925)

										gen927 := Call(__e, ShenFunc(symcons_2), gen926)

										var gen928 Obj
										if True == gen927 {
											gen919 := Call(__e, ShenFunc(symtl), V308)

											gen920 := Call(__e, ShenFunc(symhd), gen919)

											gen921 := Call(__e, ShenFunc(symtl), gen920)

											gen922 := Call(__e, ShenFunc(symtl), gen921)

											gen923 := Call(__e, ShenFunc(sym_a), Nil, gen922)

											if True == gen923 {
												gen928 = True
											} else {
												gen928 = False
											}

										} else {
											gen928 = False
										}
										if True == gen928 {
											gen932 = True
										} else {
											gen932 = False
										}

									} else {
										gen932 = False
									}
									if True == gen932 {
										gen935 = True
									} else {
										gen935 = False
									}

								} else {
									gen935 = False
								}
								if True == gen935 {
									gen940 = True
								} else {
									gen940 = False
								}

							} else {
								gen940 = False
							}
							if True == gen940 {
								gen944 = True
							} else {
								gen944 = False
							}

						} else {
							gen944 = False
						}
						if True == gen944 {
							gen947 = True
						} else {
							gen947 = False
						}

					} else {
						gen947 = False
					}
					if True == gen947 {
						gen949 = True
					} else {
						gen949 = False
					}

				} else {
					gen949 = False
				}
				if True == gen949 {
					gen911 := Call(__e, ShenFunc(symhd), V308)

					gen912 := Call(__e, ShenFunc(symhd), gen911)

					gen913 := Call(__e, ShenFunc(symlength), gen912)

					gen914 := Call(__e, ShenFunc(symtl), V308)

					gen915 := Call(__e, ShenFunc(symhd), gen914)

					gen916 := Call(__e, ShenFunc(symhd), gen915)

					gen917 := Call(__e, ShenFunc(symlength), gen916)

					gen918 := Call(__e, ShenFunc(sym_a), gen913, gen917)

					if True == gen918 {
						gen907 := Call(__e, ShenFunc(symhd), V308)

						gen908 := Call(__e, ShenFunc(symtl), gen907)

						gen909 := Call(__e, ShenFunc(symhd), gen908)

						Call(__e, ShenFunc(symshen_4aritycheck_1action), gen909)

						gen910 := Call(__e, ShenFunc(symtl), V308)

						__e.TailApply(ShenFunc(symshen_4aritycheck), V307, gen910)

						return

					} else {
						gen905 := Call(__e, ShenFunc(symshen_4app), V307, MakeString("\n"), MakeSymbol("shen.a"))

						gen906 := Call(__e, ShenFunc(symcn), MakeString("arity error in "), gen905)

						__e.TailApply(ShenFunc(symsimple_1error), gen906)

						return

					}

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.aritycheck"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck"), gen973)

		gen979 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V321 := __args[0]
			_ = V321
			V322 := __args[1]
			_ = V322
			V323 := __args[2]
			_ = V323
			gen978 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V322)

			if True == gen978 {
				__e.Return(V323)
				return
			} else {
				gen977 := Call(__e, ShenFunc(sym_a), V323, V322)

				if True == gen977 {
					__e.Return(V323)
					return
				} else {
					gen974 := Call(__e, ShenFunc(symshen_4app), V321, MakeString(" can cause errors.\n"), MakeSymbol("shen.a"))

					gen975 := Call(__e, ShenFunc(symcn), MakeString("\nwarning: changing the arity of "), gen974)

					gen976 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), gen975, gen976)

					__e.Return(V323)
					return

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck-name"), gen979)

		gen984 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V329 := __args[0]
			_ = V329
			gen983 := Call(__e, ShenFunc(symcons_2), V329)

			if True == gen983 {
				gen980 := Call(__e, ShenFunc(symhd), V329)

				gen981 := Call(__e, ShenFunc(symtl), V329)

				Call(__e, ShenFunc(symshen_4aah), gen980, gen981)

				gen982 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Y := __args[0]
					_ = Y
					__e.TailApply(ShenFunc(symshen_4aritycheck_1action), Y)

					return
				}, 1)
				__e.TailApply(ShenFunc(symshen_4for_1each), gen982, V329)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aritycheck-action"), gen984)

		gen999 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V332 := __args[0]
			_ = V332
			V333 := __args[1]
			_ = V333
			gen985 := Call(__e, ShenFunc(symarity), V332)

			Arity := gen985
			gen986 := Call(__e, ShenFunc(symlength), V333)

			Len := gen986
			gen997 := Call(__e, ShenFunc(sym_6), Arity, MakeNumber(-1))

			var gen998 Obj
			if True == gen997 {
				gen996 := Call(__e, ShenFunc(sym_6), Len, Arity)

				if True == gen996 {
					gen998 = True
				} else {
					gen998 = False
				}

			} else {
				gen998 = False
			}
			if True == gen998 {
				gen987 := Call(__e, ShenFunc(sym_6), Len, MakeNumber(1))

				var gen988 Obj
				if True == gen987 {
					gen988 = MakeString("s")
				} else {
					gen988 = MakeString("")
				}
				gen989 := Call(__e, ShenFunc(symshen_4app), gen988, MakeString(".\n"), MakeSymbol("shen.a"))

				gen990 := Call(__e, ShenFunc(symcn), MakeString(" argument"), gen989)

				gen991 := Call(__e, ShenFunc(symshen_4app), Len, gen990, MakeSymbol("shen.a"))

				gen992 := Call(__e, ShenFunc(symcn), MakeString(" might not like "), gen991)

				gen993 := Call(__e, ShenFunc(symshen_4app), V332, gen992, MakeSymbol("shen.a"))

				gen994 := Call(__e, ShenFunc(symcn), MakeString("warning: "), gen993)

				gen995 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen994, gen995)

				return

			} else {
				__e.Return(MakeSymbol("shen.skip"))
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.aah"), gen999)

		gen1011 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V335 := __args[0]
			_ = V335
			gen1009 := Call(__e, ShenFunc(symcons_2), V335)

			var gen1010 Obj
			if True == gen1009 {
				gen1006 := Call(__e, ShenFunc(symtl), V335)

				gen1007 := Call(__e, ShenFunc(symcons_2), gen1006)

				var gen1008 Obj
				if True == gen1007 {
					gen1003 := Call(__e, ShenFunc(symtl), V335)

					gen1004 := Call(__e, ShenFunc(symtl), gen1003)

					gen1005 := Call(__e, ShenFunc(sym_a), Nil, gen1004)

					if True == gen1005 {
						gen1008 = True
					} else {
						gen1008 = False
					}

				} else {
					gen1008 = False
				}
				if True == gen1008 {
					gen1010 = True
				} else {
					gen1010 = False
				}

			} else {
				gen1010 = False
			}
			if True == gen1010 {
				gen1000 := Call(__e, ShenFunc(symhd), V335)

				gen1001 := Call(__e, ShenFunc(symtl), V335)

				gen1002 := Call(__e, ShenFunc(symhd), gen1001)

				__e.TailApply(ShenFunc(symshen_4abstraction__build), gen1000, gen1002)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.abstract_rule"))

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abstract_rule"), gen1011)

		gen1019 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V338 := __args[0]
			_ = V338
			V339 := __args[1]
			_ = V339
			gen1018 := Call(__e, ShenFunc(sym_a), Nil, V338)

			if True == gen1018 {
				__e.Return(V339)
				return
			} else {
				gen1017 := Call(__e, ShenFunc(symcons_2), V338)

				if True == gen1017 {
					gen1012 := Call(__e, ShenFunc(symhd), V338)

					gen1013 := Call(__e, ShenFunc(symtl), V338)

					gen1014 := Call(__e, ShenFunc(symshen_4abstraction__build), gen1013, V339)

					gen1015 := Call(__e, ShenFunc(symcons), gen1014, Nil)

					gen1016 := Call(__e, ShenFunc(symcons), gen1012, gen1015)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("/."), gen1016)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.abstraction_build"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abstraction_build"), gen1019)

		gen1024 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V341 := __args[0]
			_ = V341
			gen1023 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V341)

			if True == gen1023 {
				__e.Return(Nil)
				return
			} else {
				gen1020 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

				gen1021 := Call(__e, ShenFunc(sym_1), V341, MakeNumber(1))

				gen1022 := Call(__e, ShenFunc(symshen_4parameters), gen1021)

				__e.TailApply(ShenFunc(symcons), gen1020, gen1022)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.parameters"), gen1024)

		gen1031 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V344 := __args[0]
			_ = V344
			V345 := __args[1]
			_ = V345
			gen1030 := Call(__e, ShenFunc(sym_a), Nil, V344)

			if True == gen1030 {
				__e.Return(V345)
				return
			} else {
				gen1029 := Call(__e, ShenFunc(symcons_2), V344)

				if True == gen1029 {
					gen1025 := Call(__e, ShenFunc(symtl), V344)

					gen1026 := Call(__e, ShenFunc(symhd), V344)

					gen1027 := Call(__e, ShenFunc(symcons), gen1026, Nil)

					gen1028 := Call(__e, ShenFunc(symcons), V345, gen1027)

					__e.TailApply(ShenFunc(symshen_4application__build), gen1025, gen1028)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.application_build"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.application_build"), gen1031)

		gen1060 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V348 := __args[0]
			_ = V348
			V349 := __args[1]
			_ = V349
			gen1058 := Call(__e, ShenFunc(symcons_2), V349)

			var gen1059 Obj
			if True == gen1058 {
				gen1055 := Call(__e, ShenFunc(symtl), V349)

				gen1056 := Call(__e, ShenFunc(symcons_2), gen1055)

				var gen1057 Obj
				if True == gen1056 {
					gen1052 := Call(__e, ShenFunc(symtl), V349)

					gen1053 := Call(__e, ShenFunc(symtl), gen1052)

					gen1054 := Call(__e, ShenFunc(sym_a), Nil, gen1053)

					if True == gen1054 {
						gen1057 = True
					} else {
						gen1057 = False
					}

				} else {
					gen1057 = False
				}
				if True == gen1057 {
					gen1059 = True
				} else {
					gen1059 = False
				}

			} else {
				gen1059 = False
			}
			if True == gen1059 {
				gen1032 := Call(__e, ShenFunc(symhd), V349)

				gen1033 := Call(__e, ShenFunc(symlength), gen1032)

				gen1034 := Call(__e, ShenFunc(symshen_4store_1arity), V348, gen1033)

				Arity := gen1034
				_ = Arity
				gen1035 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(ShenFunc(symshen_4reduce), X)

					return
				}, 1)
				gen1036 := Call(__e, ShenFunc(symtl), V349)

				gen1037 := Call(__e, ShenFunc(symhd), gen1036)

				gen1038 := Call(__e, ShenFunc(symmap), gen1035, gen1037)

				Reduce := gen1038
				gen1039 := Call(__e, ShenFunc(symhd), V349)

				gen1040 := Call(__e, ShenFunc(symshen_4cond_1expression), V348, gen1039, Reduce)

				CondExpression := gen1040
				gen1043 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*optimise*"))

				var gen1044 Obj
				if True == gen1043 {
					gen1041 := Call(__e, ShenFunc(symshen_4get_1type), V348)

					gen1042 := Call(__e, ShenFunc(symhd), V349)

					gen1044 = Call(__e, ShenFunc(symshen_4typextable), gen1041, gen1042)

				} else {
					gen1044 = MakeSymbol("shen.skip")
				}
				TypeTable := gen1044
				gen1046 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*optimise*"))

				var gen1047 Obj
				if True == gen1046 {
					gen1045 := Call(__e, ShenFunc(symhd), V349)

					gen1047 = Call(__e, ShenFunc(symshen_4assign_1types), gen1045, TypeTable, CondExpression)

				} else {
					gen1047 = CondExpression
				}
				TypedCondExpression := gen1047
				gen1048 := Call(__e, ShenFunc(symhd), V349)

				gen1049 := Call(__e, ShenFunc(symcons), TypedCondExpression, Nil)

				gen1050 := Call(__e, ShenFunc(symcons), gen1048, gen1049)

				gen1051 := Call(__e, ShenFunc(symcons), V348, gen1050)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("defun"), gen1051)

				return

			} else {
				__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compile_to_kl"))

				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile_to_kl"), gen1060)

		gen1065 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V355 := __args[0]
			_ = V355
			gen1064 := Call(__e, ShenFunc(symcons_2), V355)

			if True == gen1064 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1061 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

				gen1062 := Call(__e, ShenFunc(symassoc), V355, gen1061)

				FType := gen1062
				gen1063 := Call(__e, ShenFunc(symempty_2), FType)

				if True == gen1063 {
					__e.Return(MakeSymbol("shen.skip"))
					return
				} else {
					__e.TailApply(ShenFunc(symtl), FType)

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.get-type"), gen1065)

		gen1099 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V366 := __args[0]
			_ = V366
			V367 := __args[1]
			_ = V367
			gen1097 := Call(__e, ShenFunc(symcons_2), V366)

			var gen1098 Obj
			if True == gen1097 {
				gen1094 := Call(__e, ShenFunc(symtl), V366)

				gen1095 := Call(__e, ShenFunc(symcons_2), gen1094)

				var gen1096 Obj
				if True == gen1095 {
					gen1090 := Call(__e, ShenFunc(symtl), V366)

					gen1091 := Call(__e, ShenFunc(symhd), gen1090)

					gen1092 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen1091)

					var gen1093 Obj
					if True == gen1092 {
						gen1086 := Call(__e, ShenFunc(symtl), V366)

						gen1087 := Call(__e, ShenFunc(symtl), gen1086)

						gen1088 := Call(__e, ShenFunc(symcons_2), gen1087)

						var gen1089 Obj
						if True == gen1088 {
							gen1081 := Call(__e, ShenFunc(symtl), V366)

							gen1082 := Call(__e, ShenFunc(symtl), gen1081)

							gen1083 := Call(__e, ShenFunc(symtl), gen1082)

							gen1084 := Call(__e, ShenFunc(sym_a), Nil, gen1083)

							var gen1085 Obj
							if True == gen1084 {
								gen1080 := Call(__e, ShenFunc(symcons_2), V367)

								if True == gen1080 {
									gen1085 = True
								} else {
									gen1085 = False
								}

							} else {
								gen1085 = False
							}
							if True == gen1085 {
								gen1089 = True
							} else {
								gen1089 = False
							}

						} else {
							gen1089 = False
						}
						if True == gen1089 {
							gen1093 = True
						} else {
							gen1093 = False
						}

					} else {
						gen1093 = False
					}
					if True == gen1093 {
						gen1096 = True
					} else {
						gen1096 = False
					}

				} else {
					gen1096 = False
				}
				if True == gen1096 {
					gen1098 = True
				} else {
					gen1098 = False
				}

			} else {
				gen1098 = False
			}
			if True == gen1098 {
				gen1078 := Call(__e, ShenFunc(symhd), V366)

				gen1079 := Call(__e, ShenFunc(symvariable_2), gen1078)

				if True == gen1079 {
					gen1074 := Call(__e, ShenFunc(symtl), V366)

					gen1075 := Call(__e, ShenFunc(symtl), gen1074)

					gen1076 := Call(__e, ShenFunc(symhd), gen1075)

					gen1077 := Call(__e, ShenFunc(symtl), V367)

					__e.TailApply(ShenFunc(symshen_4typextable), gen1076, gen1077)

					return

				} else {
					gen1066 := Call(__e, ShenFunc(symhd), V367)

					gen1067 := Call(__e, ShenFunc(symhd), V366)

					gen1068 := Call(__e, ShenFunc(symcons), gen1066, gen1067)

					gen1069 := Call(__e, ShenFunc(symtl), V366)

					gen1070 := Call(__e, ShenFunc(symtl), gen1069)

					gen1071 := Call(__e, ShenFunc(symhd), gen1070)

					gen1072 := Call(__e, ShenFunc(symtl), V367)

					gen1073 := Call(__e, ShenFunc(symshen_4typextable), gen1071, gen1072)

					__e.TailApply(ShenFunc(symcons), gen1068, gen1073)

					return

				}

			} else {
				__e.Return(Nil)
				return
			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typextable"), gen1099)

		gen1195 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V371 := __args[0]
			_ = V371
			V372 := __args[1]
			_ = V372
			V373 := __args[2]
			_ = V373
			gen1193 := Call(__e, ShenFunc(symcons_2), V373)

			var gen1194 Obj
			if True == gen1193 {
				gen1190 := Call(__e, ShenFunc(symhd), V373)

				gen1191 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen1190)

				var gen1192 Obj
				if True == gen1191 {
					gen1187 := Call(__e, ShenFunc(symtl), V373)

					gen1188 := Call(__e, ShenFunc(symcons_2), gen1187)

					var gen1189 Obj
					if True == gen1188 {
						gen1183 := Call(__e, ShenFunc(symtl), V373)

						gen1184 := Call(__e, ShenFunc(symtl), gen1183)

						gen1185 := Call(__e, ShenFunc(symcons_2), gen1184)

						var gen1186 Obj
						if True == gen1185 {
							gen1178 := Call(__e, ShenFunc(symtl), V373)

							gen1179 := Call(__e, ShenFunc(symtl), gen1178)

							gen1180 := Call(__e, ShenFunc(symtl), gen1179)

							gen1181 := Call(__e, ShenFunc(symcons_2), gen1180)

							var gen1182 Obj
							if True == gen1181 {
								gen1173 := Call(__e, ShenFunc(symtl), V373)

								gen1174 := Call(__e, ShenFunc(symtl), gen1173)

								gen1175 := Call(__e, ShenFunc(symtl), gen1174)

								gen1176 := Call(__e, ShenFunc(symtl), gen1175)

								gen1177 := Call(__e, ShenFunc(sym_a), Nil, gen1176)

								if True == gen1177 {
									gen1182 = True
								} else {
									gen1182 = False
								}

							} else {
								gen1182 = False
							}
							if True == gen1182 {
								gen1186 = True
							} else {
								gen1186 = False
							}

						} else {
							gen1186 = False
						}
						if True == gen1186 {
							gen1189 = True
						} else {
							gen1189 = False
						}

					} else {
						gen1189 = False
					}
					if True == gen1189 {
						gen1192 = True
					} else {
						gen1192 = False
					}

				} else {
					gen1192 = False
				}
				if True == gen1192 {
					gen1194 = True
				} else {
					gen1194 = False
				}

			} else {
				gen1194 = False
			}
			if True == gen1194 {
				gen1156 := Call(__e, ShenFunc(symtl), V373)

				gen1157 := Call(__e, ShenFunc(symhd), gen1156)

				gen1158 := Call(__e, ShenFunc(symtl), V373)

				gen1159 := Call(__e, ShenFunc(symtl), gen1158)

				gen1160 := Call(__e, ShenFunc(symhd), gen1159)

				gen1161 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1160)

				gen1162 := Call(__e, ShenFunc(symtl), V373)

				gen1163 := Call(__e, ShenFunc(symhd), gen1162)

				gen1164 := Call(__e, ShenFunc(symcons), gen1163, V371)

				gen1165 := Call(__e, ShenFunc(symtl), V373)

				gen1166 := Call(__e, ShenFunc(symtl), gen1165)

				gen1167 := Call(__e, ShenFunc(symtl), gen1166)

				gen1168 := Call(__e, ShenFunc(symhd), gen1167)

				gen1169 := Call(__e, ShenFunc(symshen_4assign_1types), gen1164, V372, gen1168)

				gen1170 := Call(__e, ShenFunc(symcons), gen1169, Nil)

				gen1171 := Call(__e, ShenFunc(symcons), gen1161, gen1170)

				gen1172 := Call(__e, ShenFunc(symcons), gen1157, gen1171)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen1172)

				return

			} else {
				gen1154 := Call(__e, ShenFunc(symcons_2), V373)

				var gen1155 Obj
				if True == gen1154 {
					gen1151 := Call(__e, ShenFunc(symhd), V373)

					gen1152 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen1151)

					var gen1153 Obj
					if True == gen1152 {
						gen1148 := Call(__e, ShenFunc(symtl), V373)

						gen1149 := Call(__e, ShenFunc(symcons_2), gen1148)

						var gen1150 Obj
						if True == gen1149 {
							gen1144 := Call(__e, ShenFunc(symtl), V373)

							gen1145 := Call(__e, ShenFunc(symtl), gen1144)

							gen1146 := Call(__e, ShenFunc(symcons_2), gen1145)

							var gen1147 Obj
							if True == gen1146 {
								gen1140 := Call(__e, ShenFunc(symtl), V373)

								gen1141 := Call(__e, ShenFunc(symtl), gen1140)

								gen1142 := Call(__e, ShenFunc(symtl), gen1141)

								gen1143 := Call(__e, ShenFunc(sym_a), Nil, gen1142)

								if True == gen1143 {
									gen1147 = True
								} else {
									gen1147 = False
								}

							} else {
								gen1147 = False
							}
							if True == gen1147 {
								gen1150 = True
							} else {
								gen1150 = False
							}

						} else {
							gen1150 = False
						}
						if True == gen1150 {
							gen1153 = True
						} else {
							gen1153 = False
						}

					} else {
						gen1153 = False
					}
					if True == gen1153 {
						gen1155 = True
					} else {
						gen1155 = False
					}

				} else {
					gen1155 = False
				}
				if True == gen1155 {
					gen1129 := Call(__e, ShenFunc(symtl), V373)

					gen1130 := Call(__e, ShenFunc(symhd), gen1129)

					gen1131 := Call(__e, ShenFunc(symtl), V373)

					gen1132 := Call(__e, ShenFunc(symhd), gen1131)

					gen1133 := Call(__e, ShenFunc(symcons), gen1132, V371)

					gen1134 := Call(__e, ShenFunc(symtl), V373)

					gen1135 := Call(__e, ShenFunc(symtl), gen1134)

					gen1136 := Call(__e, ShenFunc(symhd), gen1135)

					gen1137 := Call(__e, ShenFunc(symshen_4assign_1types), gen1133, V372, gen1136)

					gen1138 := Call(__e, ShenFunc(symcons), gen1137, Nil)

					gen1139 := Call(__e, ShenFunc(symcons), gen1130, gen1138)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen1139)

					return

				} else {
					gen1127 := Call(__e, ShenFunc(symcons_2), V373)

					var gen1128 Obj
					if True == gen1127 {
						gen1125 := Call(__e, ShenFunc(symhd), V373)

						gen1126 := Call(__e, ShenFunc(sym_a), MakeSymbol("cond"), gen1125)

						if True == gen1126 {
							gen1128 = True
						} else {
							gen1128 = False
						}

					} else {
						gen1128 = False
					}
					if True == gen1128 {
						gen1122 := MakeNative(func(__e Evaluator, __args ...Obj) {
							Y := __args[0]
							_ = Y
							gen1116 := Call(__e, ShenFunc(symhd), Y)

							gen1117 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1116)

							gen1118 := Call(__e, ShenFunc(symtl), Y)

							gen1119 := Call(__e, ShenFunc(symhd), gen1118)

							gen1120 := Call(__e, ShenFunc(symshen_4assign_1types), V371, V372, gen1119)

							gen1121 := Call(__e, ShenFunc(symcons), gen1120, Nil)

							__e.TailApply(ShenFunc(symcons), gen1117, gen1121)

							return

						}, 1)
						gen1123 := Call(__e, ShenFunc(symtl), V373)

						gen1124 := Call(__e, ShenFunc(symmap), gen1122, gen1123)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("cond"), gen1124)

						return

					} else {
						gen1115 := Call(__e, ShenFunc(symcons_2), V373)

						if True == gen1115 {
							gen1106 := Call(__e, ShenFunc(symhd), V373)

							gen1107 := Call(__e, ShenFunc(symshen_4get_1type), gen1106)

							gen1108 := Call(__e, ShenFunc(symtl), V373)

							gen1109 := Call(__e, ShenFunc(symshen_4typextable), gen1107, gen1108)

							NewTable := gen1109
							gen1110 := Call(__e, ShenFunc(symhd), V373)

							gen1112 := MakeNative(func(__e Evaluator, __args ...Obj) {
								Y := __args[0]
								_ = Y
								gen1111 := Call(__e, ShenFunc(symappend), V372, NewTable)

								__e.TailApply(ShenFunc(symshen_4assign_1types), V371, gen1111, Y)

								return

							}, 1)
							gen1113 := Call(__e, ShenFunc(symtl), V373)

							gen1114 := Call(__e, ShenFunc(symmap), gen1112, gen1113)

							__e.TailApply(ShenFunc(symcons), gen1110, gen1114)

							return

						} else {
							gen1100 := Call(__e, ShenFunc(symassoc), V373, V372)

							AtomType := gen1100
							gen1105 := Call(__e, ShenFunc(symcons_2), AtomType)

							if True == gen1105 {
								gen1102 := Call(__e, ShenFunc(symtl), AtomType)

								gen1103 := Call(__e, ShenFunc(symcons), gen1102, Nil)

								gen1104 := Call(__e, ShenFunc(symcons), V373, gen1103)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1104)

								return

							} else {
								gen1101 := Call(__e, ShenFunc(symelement_2), V373, V371)

								if True == gen1101 {
									__e.Return(V373)
									return
								} else {
									__e.TailApply(ShenFunc(symshen_4atom_1type), V373)

									return
								}

							}

						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assign-types"), gen1195)

		gen1208 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V375 := __args[0]
			_ = V375
			gen1207 := Call(__e, ShenFunc(symstring_2), V375)

			if True == gen1207 {
				gen1205 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

				gen1206 := Call(__e, ShenFunc(symcons), V375, gen1205)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1206)

				return

			} else {
				gen1204 := Call(__e, ShenFunc(symnumber_2), V375)

				if True == gen1204 {
					gen1202 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

					gen1203 := Call(__e, ShenFunc(symcons), V375, gen1202)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1203)

					return

				} else {
					gen1201 := Call(__e, ShenFunc(symboolean_2), V375)

					if True == gen1201 {
						gen1199 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

						gen1200 := Call(__e, ShenFunc(symcons), V375, gen1199)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1200)

						return

					} else {
						gen1198 := Call(__e, ShenFunc(symsymbol_2), V375)

						if True == gen1198 {
							gen1196 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

							gen1197 := Call(__e, ShenFunc(symcons), V375, gen1196)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen1197)

							return

						} else {
							__e.Return(V375)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.atom-type"), gen1208)

		gen1211 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V380 := __args[0]
			_ = V380
			V381 := __args[1]
			_ = V381
			gen1210 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

			if True == gen1210 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen1209 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

				__e.TailApply(ShenFunc(symput), V380, MakeSymbol("arity"), V381, gen1209)

				return

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.store-arity"), gen1211)

		gen1218 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V383 := __args[0]
			_ = V383
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*teststack*"), Nil)
			gen1212 := Call(__e, ShenFunc(symshen_4reduce__help), V383)

			Result := gen1212
			gen1213 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*teststack*"))

			gen1214 := Call(__e, ShenFunc(symreverse), gen1213)

			gen1215 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tests"), gen1214)

			gen1216 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen1215)

			gen1217 := Call(__e, ShenFunc(symcons), Result, Nil)

			__e.TailApply(ShenFunc(symcons), gen1216, gen1217)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reduce"), gen1218)

		gen1774 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V385 := __args[0]
			_ = V385
			gen1772 := Call(__e, ShenFunc(symcons_2), V385)

			var gen1773 Obj
			if True == gen1772 {
				gen1769 := Call(__e, ShenFunc(symhd), V385)

				gen1770 := Call(__e, ShenFunc(symcons_2), gen1769)

				var gen1771 Obj
				if True == gen1770 {
					gen1765 := Call(__e, ShenFunc(symhd), V385)

					gen1766 := Call(__e, ShenFunc(symhd), gen1765)

					gen1767 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1766)

					var gen1768 Obj
					if True == gen1767 {
						gen1761 := Call(__e, ShenFunc(symhd), V385)

						gen1762 := Call(__e, ShenFunc(symtl), gen1761)

						gen1763 := Call(__e, ShenFunc(symcons_2), gen1762)

						var gen1764 Obj
						if True == gen1763 {
							gen1756 := Call(__e, ShenFunc(symhd), V385)

							gen1757 := Call(__e, ShenFunc(symtl), gen1756)

							gen1758 := Call(__e, ShenFunc(symhd), gen1757)

							gen1759 := Call(__e, ShenFunc(symcons_2), gen1758)

							var gen1760 Obj
							if True == gen1759 {
								gen1750 := Call(__e, ShenFunc(symhd), V385)

								gen1751 := Call(__e, ShenFunc(symtl), gen1750)

								gen1752 := Call(__e, ShenFunc(symhd), gen1751)

								gen1753 := Call(__e, ShenFunc(symhd), gen1752)

								gen1754 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), gen1753)

								var gen1755 Obj
								if True == gen1754 {
									gen1744 := Call(__e, ShenFunc(symhd), V385)

									gen1745 := Call(__e, ShenFunc(symtl), gen1744)

									gen1746 := Call(__e, ShenFunc(symhd), gen1745)

									gen1747 := Call(__e, ShenFunc(symtl), gen1746)

									gen1748 := Call(__e, ShenFunc(symcons_2), gen1747)

									var gen1749 Obj
									if True == gen1748 {
										gen1737 := Call(__e, ShenFunc(symhd), V385)

										gen1738 := Call(__e, ShenFunc(symtl), gen1737)

										gen1739 := Call(__e, ShenFunc(symhd), gen1738)

										gen1740 := Call(__e, ShenFunc(symtl), gen1739)

										gen1741 := Call(__e, ShenFunc(symtl), gen1740)

										gen1742 := Call(__e, ShenFunc(symcons_2), gen1741)

										var gen1743 Obj
										if True == gen1742 {
											gen1729 := Call(__e, ShenFunc(symhd), V385)

											gen1730 := Call(__e, ShenFunc(symtl), gen1729)

											gen1731 := Call(__e, ShenFunc(symhd), gen1730)

											gen1732 := Call(__e, ShenFunc(symtl), gen1731)

											gen1733 := Call(__e, ShenFunc(symtl), gen1732)

											gen1734 := Call(__e, ShenFunc(symtl), gen1733)

											gen1735 := Call(__e, ShenFunc(sym_a), Nil, gen1734)

											var gen1736 Obj
											if True == gen1735 {
												gen1724 := Call(__e, ShenFunc(symhd), V385)

												gen1725 := Call(__e, ShenFunc(symtl), gen1724)

												gen1726 := Call(__e, ShenFunc(symtl), gen1725)

												gen1727 := Call(__e, ShenFunc(symcons_2), gen1726)

												var gen1728 Obj
												if True == gen1727 {
													gen1718 := Call(__e, ShenFunc(symhd), V385)

													gen1719 := Call(__e, ShenFunc(symtl), gen1718)

													gen1720 := Call(__e, ShenFunc(symtl), gen1719)

													gen1721 := Call(__e, ShenFunc(symtl), gen1720)

													gen1722 := Call(__e, ShenFunc(sym_a), Nil, gen1721)

													var gen1723 Obj
													if True == gen1722 {
														gen1715 := Call(__e, ShenFunc(symtl), V385)

														gen1716 := Call(__e, ShenFunc(symcons_2), gen1715)

														var gen1717 Obj
														if True == gen1716 {
															gen1712 := Call(__e, ShenFunc(symtl), V385)

															gen1713 := Call(__e, ShenFunc(symtl), gen1712)

															gen1714 := Call(__e, ShenFunc(sym_a), Nil, gen1713)

															if True == gen1714 {
																gen1717 = True
															} else {
																gen1717 = False
															}

														} else {
															gen1717 = False
														}
														if True == gen1717 {
															gen1723 = True
														} else {
															gen1723 = False
														}

													} else {
														gen1723 = False
													}
													if True == gen1723 {
														gen1728 = True
													} else {
														gen1728 = False
													}

												} else {
													gen1728 = False
												}
												if True == gen1728 {
													gen1736 = True
												} else {
													gen1736 = False
												}

											} else {
												gen1736 = False
											}
											if True == gen1736 {
												gen1743 = True
											} else {
												gen1743 = False
											}

										} else {
											gen1743 = False
										}
										if True == gen1743 {
											gen1749 = True
										} else {
											gen1749 = False
										}

									} else {
										gen1749 = False
									}
									if True == gen1749 {
										gen1755 = True
									} else {
										gen1755 = False
									}

								} else {
									gen1755 = False
								}
								if True == gen1755 {
									gen1760 = True
								} else {
									gen1760 = False
								}

							} else {
								gen1760 = False
							}
							if True == gen1760 {
								gen1764 = True
							} else {
								gen1764 = False
							}

						} else {
							gen1764 = False
						}
						if True == gen1764 {
							gen1768 = True
						} else {
							gen1768 = False
						}

					} else {
						gen1768 = False
					}
					if True == gen1768 {
						gen1771 = True
					} else {
						gen1771 = False
					}

				} else {
					gen1771 = False
				}
				if True == gen1771 {
					gen1773 = True
				} else {
					gen1773 = False
				}

			} else {
				gen1773 = False
			}
			if True == gen1773 {
				gen1675 := Call(__e, ShenFunc(symtl), V385)

				gen1676 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen1675)

				Call(__e, ShenFunc(symshen_4add__test), gen1676)

				gen1677 := Call(__e, ShenFunc(symhd), V385)

				gen1678 := Call(__e, ShenFunc(symtl), gen1677)

				gen1679 := Call(__e, ShenFunc(symhd), gen1678)

				gen1680 := Call(__e, ShenFunc(symtl), gen1679)

				gen1681 := Call(__e, ShenFunc(symhd), gen1680)

				gen1682 := Call(__e, ShenFunc(symhd), V385)

				gen1683 := Call(__e, ShenFunc(symtl), gen1682)

				gen1684 := Call(__e, ShenFunc(symhd), gen1683)

				gen1685 := Call(__e, ShenFunc(symtl), gen1684)

				gen1686 := Call(__e, ShenFunc(symtl), gen1685)

				gen1687 := Call(__e, ShenFunc(symhd), gen1686)

				gen1688 := Call(__e, ShenFunc(symtl), V385)

				gen1689 := Call(__e, ShenFunc(symhd), gen1688)

				gen1690 := Call(__e, ShenFunc(symhd), V385)

				gen1691 := Call(__e, ShenFunc(symtl), gen1690)

				gen1692 := Call(__e, ShenFunc(symhd), gen1691)

				gen1693 := Call(__e, ShenFunc(symhd), V385)

				gen1694 := Call(__e, ShenFunc(symtl), gen1693)

				gen1695 := Call(__e, ShenFunc(symtl), gen1694)

				gen1696 := Call(__e, ShenFunc(symhd), gen1695)

				gen1697 := Call(__e, ShenFunc(symshen_4ebr), gen1689, gen1692, gen1696)

				gen1698 := Call(__e, ShenFunc(symcons), gen1697, Nil)

				gen1699 := Call(__e, ShenFunc(symcons), gen1687, gen1698)

				gen1700 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1699)

				gen1701 := Call(__e, ShenFunc(symcons), gen1700, Nil)

				gen1702 := Call(__e, ShenFunc(symcons), gen1681, gen1701)

				gen1703 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1702)

				Abstraction := gen1703
				gen1704 := Call(__e, ShenFunc(symtl), V385)

				gen1705 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen1704)

				gen1706 := Call(__e, ShenFunc(symcons), gen1705, Nil)

				gen1707 := Call(__e, ShenFunc(symcons), Abstraction, gen1706)

				gen1708 := Call(__e, ShenFunc(symtl), V385)

				gen1709 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen1708)

				gen1710 := Call(__e, ShenFunc(symcons), gen1709, Nil)

				gen1711 := Call(__e, ShenFunc(symcons), gen1707, gen1710)

				Application := gen1711
				__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

				return

			} else {
				gen1673 := Call(__e, ShenFunc(symcons_2), V385)

				var gen1674 Obj
				if True == gen1673 {
					gen1670 := Call(__e, ShenFunc(symhd), V385)

					gen1671 := Call(__e, ShenFunc(symcons_2), gen1670)

					var gen1672 Obj
					if True == gen1671 {
						gen1666 := Call(__e, ShenFunc(symhd), V385)

						gen1667 := Call(__e, ShenFunc(symhd), gen1666)

						gen1668 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1667)

						var gen1669 Obj
						if True == gen1668 {
							gen1662 := Call(__e, ShenFunc(symhd), V385)

							gen1663 := Call(__e, ShenFunc(symtl), gen1662)

							gen1664 := Call(__e, ShenFunc(symcons_2), gen1663)

							var gen1665 Obj
							if True == gen1664 {
								gen1657 := Call(__e, ShenFunc(symhd), V385)

								gen1658 := Call(__e, ShenFunc(symtl), gen1657)

								gen1659 := Call(__e, ShenFunc(symhd), gen1658)

								gen1660 := Call(__e, ShenFunc(symcons_2), gen1659)

								var gen1661 Obj
								if True == gen1660 {
									gen1651 := Call(__e, ShenFunc(symhd), V385)

									gen1652 := Call(__e, ShenFunc(symtl), gen1651)

									gen1653 := Call(__e, ShenFunc(symhd), gen1652)

									gen1654 := Call(__e, ShenFunc(symhd), gen1653)

									gen1655 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), gen1654)

									var gen1656 Obj
									if True == gen1655 {
										gen1645 := Call(__e, ShenFunc(symhd), V385)

										gen1646 := Call(__e, ShenFunc(symtl), gen1645)

										gen1647 := Call(__e, ShenFunc(symhd), gen1646)

										gen1648 := Call(__e, ShenFunc(symtl), gen1647)

										gen1649 := Call(__e, ShenFunc(symcons_2), gen1648)

										var gen1650 Obj
										if True == gen1649 {
											gen1638 := Call(__e, ShenFunc(symhd), V385)

											gen1639 := Call(__e, ShenFunc(symtl), gen1638)

											gen1640 := Call(__e, ShenFunc(symhd), gen1639)

											gen1641 := Call(__e, ShenFunc(symtl), gen1640)

											gen1642 := Call(__e, ShenFunc(symtl), gen1641)

											gen1643 := Call(__e, ShenFunc(symcons_2), gen1642)

											var gen1644 Obj
											if True == gen1643 {
												gen1630 := Call(__e, ShenFunc(symhd), V385)

												gen1631 := Call(__e, ShenFunc(symtl), gen1630)

												gen1632 := Call(__e, ShenFunc(symhd), gen1631)

												gen1633 := Call(__e, ShenFunc(symtl), gen1632)

												gen1634 := Call(__e, ShenFunc(symtl), gen1633)

												gen1635 := Call(__e, ShenFunc(symtl), gen1634)

												gen1636 := Call(__e, ShenFunc(sym_a), Nil, gen1635)

												var gen1637 Obj
												if True == gen1636 {
													gen1625 := Call(__e, ShenFunc(symhd), V385)

													gen1626 := Call(__e, ShenFunc(symtl), gen1625)

													gen1627 := Call(__e, ShenFunc(symtl), gen1626)

													gen1628 := Call(__e, ShenFunc(symcons_2), gen1627)

													var gen1629 Obj
													if True == gen1628 {
														gen1619 := Call(__e, ShenFunc(symhd), V385)

														gen1620 := Call(__e, ShenFunc(symtl), gen1619)

														gen1621 := Call(__e, ShenFunc(symtl), gen1620)

														gen1622 := Call(__e, ShenFunc(symtl), gen1621)

														gen1623 := Call(__e, ShenFunc(sym_a), Nil, gen1622)

														var gen1624 Obj
														if True == gen1623 {
															gen1616 := Call(__e, ShenFunc(symtl), V385)

															gen1617 := Call(__e, ShenFunc(symcons_2), gen1616)

															var gen1618 Obj
															if True == gen1617 {
																gen1613 := Call(__e, ShenFunc(symtl), V385)

																gen1614 := Call(__e, ShenFunc(symtl), gen1613)

																gen1615 := Call(__e, ShenFunc(sym_a), Nil, gen1614)

																if True == gen1615 {
																	gen1618 = True
																} else {
																	gen1618 = False
																}

															} else {
																gen1618 = False
															}
															if True == gen1618 {
																gen1624 = True
															} else {
																gen1624 = False
															}

														} else {
															gen1624 = False
														}
														if True == gen1624 {
															gen1629 = True
														} else {
															gen1629 = False
														}

													} else {
														gen1629 = False
													}
													if True == gen1629 {
														gen1637 = True
													} else {
														gen1637 = False
													}

												} else {
													gen1637 = False
												}
												if True == gen1637 {
													gen1644 = True
												} else {
													gen1644 = False
												}

											} else {
												gen1644 = False
											}
											if True == gen1644 {
												gen1650 = True
											} else {
												gen1650 = False
											}

										} else {
											gen1650 = False
										}
										if True == gen1650 {
											gen1656 = True
										} else {
											gen1656 = False
										}

									} else {
										gen1656 = False
									}
									if True == gen1656 {
										gen1661 = True
									} else {
										gen1661 = False
									}

								} else {
									gen1661 = False
								}
								if True == gen1661 {
									gen1665 = True
								} else {
									gen1665 = False
								}

							} else {
								gen1665 = False
							}
							if True == gen1665 {
								gen1669 = True
							} else {
								gen1669 = False
							}

						} else {
							gen1669 = False
						}
						if True == gen1669 {
							gen1672 = True
						} else {
							gen1672 = False
						}

					} else {
						gen1672 = False
					}
					if True == gen1672 {
						gen1674 = True
					} else {
						gen1674 = False
					}

				} else {
					gen1674 = False
				}
				if True == gen1674 {
					gen1576 := Call(__e, ShenFunc(symtl), V385)

					gen1577 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen1576)

					Call(__e, ShenFunc(symshen_4add__test), gen1577)

					gen1578 := Call(__e, ShenFunc(symhd), V385)

					gen1579 := Call(__e, ShenFunc(symtl), gen1578)

					gen1580 := Call(__e, ShenFunc(symhd), gen1579)

					gen1581 := Call(__e, ShenFunc(symtl), gen1580)

					gen1582 := Call(__e, ShenFunc(symhd), gen1581)

					gen1583 := Call(__e, ShenFunc(symhd), V385)

					gen1584 := Call(__e, ShenFunc(symtl), gen1583)

					gen1585 := Call(__e, ShenFunc(symhd), gen1584)

					gen1586 := Call(__e, ShenFunc(symtl), gen1585)

					gen1587 := Call(__e, ShenFunc(symtl), gen1586)

					gen1588 := Call(__e, ShenFunc(symhd), gen1587)

					gen1589 := Call(__e, ShenFunc(symtl), V385)

					gen1590 := Call(__e, ShenFunc(symhd), gen1589)

					gen1591 := Call(__e, ShenFunc(symhd), V385)

					gen1592 := Call(__e, ShenFunc(symtl), gen1591)

					gen1593 := Call(__e, ShenFunc(symhd), gen1592)

					gen1594 := Call(__e, ShenFunc(symhd), V385)

					gen1595 := Call(__e, ShenFunc(symtl), gen1594)

					gen1596 := Call(__e, ShenFunc(symtl), gen1595)

					gen1597 := Call(__e, ShenFunc(symhd), gen1596)

					gen1598 := Call(__e, ShenFunc(symshen_4ebr), gen1590, gen1593, gen1597)

					gen1599 := Call(__e, ShenFunc(symcons), gen1598, Nil)

					gen1600 := Call(__e, ShenFunc(symcons), gen1588, gen1599)

					gen1601 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1600)

					gen1602 := Call(__e, ShenFunc(symcons), gen1601, Nil)

					gen1603 := Call(__e, ShenFunc(symcons), gen1582, gen1602)

					gen1604 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1603)

					Abstraction := gen1604
					gen1605 := Call(__e, ShenFunc(symtl), V385)

					gen1606 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen1605)

					gen1607 := Call(__e, ShenFunc(symcons), gen1606, Nil)

					gen1608 := Call(__e, ShenFunc(symcons), Abstraction, gen1607)

					gen1609 := Call(__e, ShenFunc(symtl), V385)

					gen1610 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen1609)

					gen1611 := Call(__e, ShenFunc(symcons), gen1610, Nil)

					gen1612 := Call(__e, ShenFunc(symcons), gen1608, gen1611)

					Application := gen1612
					__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

					return

				} else {
					gen1574 := Call(__e, ShenFunc(symcons_2), V385)

					var gen1575 Obj
					if True == gen1574 {
						gen1571 := Call(__e, ShenFunc(symhd), V385)

						gen1572 := Call(__e, ShenFunc(symcons_2), gen1571)

						var gen1573 Obj
						if True == gen1572 {
							gen1567 := Call(__e, ShenFunc(symhd), V385)

							gen1568 := Call(__e, ShenFunc(symhd), gen1567)

							gen1569 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1568)

							var gen1570 Obj
							if True == gen1569 {
								gen1563 := Call(__e, ShenFunc(symhd), V385)

								gen1564 := Call(__e, ShenFunc(symtl), gen1563)

								gen1565 := Call(__e, ShenFunc(symcons_2), gen1564)

								var gen1566 Obj
								if True == gen1565 {
									gen1558 := Call(__e, ShenFunc(symhd), V385)

									gen1559 := Call(__e, ShenFunc(symtl), gen1558)

									gen1560 := Call(__e, ShenFunc(symhd), gen1559)

									gen1561 := Call(__e, ShenFunc(symcons_2), gen1560)

									var gen1562 Obj
									if True == gen1561 {
										gen1552 := Call(__e, ShenFunc(symhd), V385)

										gen1553 := Call(__e, ShenFunc(symtl), gen1552)

										gen1554 := Call(__e, ShenFunc(symhd), gen1553)

										gen1555 := Call(__e, ShenFunc(symhd), gen1554)

										gen1556 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), gen1555)

										var gen1557 Obj
										if True == gen1556 {
											gen1546 := Call(__e, ShenFunc(symhd), V385)

											gen1547 := Call(__e, ShenFunc(symtl), gen1546)

											gen1548 := Call(__e, ShenFunc(symhd), gen1547)

											gen1549 := Call(__e, ShenFunc(symtl), gen1548)

											gen1550 := Call(__e, ShenFunc(symcons_2), gen1549)

											var gen1551 Obj
											if True == gen1550 {
												gen1539 := Call(__e, ShenFunc(symhd), V385)

												gen1540 := Call(__e, ShenFunc(symtl), gen1539)

												gen1541 := Call(__e, ShenFunc(symhd), gen1540)

												gen1542 := Call(__e, ShenFunc(symtl), gen1541)

												gen1543 := Call(__e, ShenFunc(symtl), gen1542)

												gen1544 := Call(__e, ShenFunc(symcons_2), gen1543)

												var gen1545 Obj
												if True == gen1544 {
													gen1531 := Call(__e, ShenFunc(symhd), V385)

													gen1532 := Call(__e, ShenFunc(symtl), gen1531)

													gen1533 := Call(__e, ShenFunc(symhd), gen1532)

													gen1534 := Call(__e, ShenFunc(symtl), gen1533)

													gen1535 := Call(__e, ShenFunc(symtl), gen1534)

													gen1536 := Call(__e, ShenFunc(symtl), gen1535)

													gen1537 := Call(__e, ShenFunc(sym_a), Nil, gen1536)

													var gen1538 Obj
													if True == gen1537 {
														gen1526 := Call(__e, ShenFunc(symhd), V385)

														gen1527 := Call(__e, ShenFunc(symtl), gen1526)

														gen1528 := Call(__e, ShenFunc(symtl), gen1527)

														gen1529 := Call(__e, ShenFunc(symcons_2), gen1528)

														var gen1530 Obj
														if True == gen1529 {
															gen1520 := Call(__e, ShenFunc(symhd), V385)

															gen1521 := Call(__e, ShenFunc(symtl), gen1520)

															gen1522 := Call(__e, ShenFunc(symtl), gen1521)

															gen1523 := Call(__e, ShenFunc(symtl), gen1522)

															gen1524 := Call(__e, ShenFunc(sym_a), Nil, gen1523)

															var gen1525 Obj
															if True == gen1524 {
																gen1517 := Call(__e, ShenFunc(symtl), V385)

																gen1518 := Call(__e, ShenFunc(symcons_2), gen1517)

																var gen1519 Obj
																if True == gen1518 {
																	gen1514 := Call(__e, ShenFunc(symtl), V385)

																	gen1515 := Call(__e, ShenFunc(symtl), gen1514)

																	gen1516 := Call(__e, ShenFunc(sym_a), Nil, gen1515)

																	if True == gen1516 {
																		gen1519 = True
																	} else {
																		gen1519 = False
																	}

																} else {
																	gen1519 = False
																}
																if True == gen1519 {
																	gen1525 = True
																} else {
																	gen1525 = False
																}

															} else {
																gen1525 = False
															}
															if True == gen1525 {
																gen1530 = True
															} else {
																gen1530 = False
															}

														} else {
															gen1530 = False
														}
														if True == gen1530 {
															gen1538 = True
														} else {
															gen1538 = False
														}

													} else {
														gen1538 = False
													}
													if True == gen1538 {
														gen1545 = True
													} else {
														gen1545 = False
													}

												} else {
													gen1545 = False
												}
												if True == gen1545 {
													gen1551 = True
												} else {
													gen1551 = False
												}

											} else {
												gen1551 = False
											}
											if True == gen1551 {
												gen1557 = True
											} else {
												gen1557 = False
											}

										} else {
											gen1557 = False
										}
										if True == gen1557 {
											gen1562 = True
										} else {
											gen1562 = False
										}

									} else {
										gen1562 = False
									}
									if True == gen1562 {
										gen1566 = True
									} else {
										gen1566 = False
									}

								} else {
									gen1566 = False
								}
								if True == gen1566 {
									gen1570 = True
								} else {
									gen1570 = False
								}

							} else {
								gen1570 = False
							}
							if True == gen1570 {
								gen1573 = True
							} else {
								gen1573 = False
							}

						} else {
							gen1573 = False
						}
						if True == gen1573 {
							gen1575 = True
						} else {
							gen1575 = False
						}

					} else {
						gen1575 = False
					}
					if True == gen1575 {
						gen1477 := Call(__e, ShenFunc(symtl), V385)

						gen1478 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.+vector?"), gen1477)

						Call(__e, ShenFunc(symshen_4add__test), gen1478)

						gen1479 := Call(__e, ShenFunc(symhd), V385)

						gen1480 := Call(__e, ShenFunc(symtl), gen1479)

						gen1481 := Call(__e, ShenFunc(symhd), gen1480)

						gen1482 := Call(__e, ShenFunc(symtl), gen1481)

						gen1483 := Call(__e, ShenFunc(symhd), gen1482)

						gen1484 := Call(__e, ShenFunc(symhd), V385)

						gen1485 := Call(__e, ShenFunc(symtl), gen1484)

						gen1486 := Call(__e, ShenFunc(symhd), gen1485)

						gen1487 := Call(__e, ShenFunc(symtl), gen1486)

						gen1488 := Call(__e, ShenFunc(symtl), gen1487)

						gen1489 := Call(__e, ShenFunc(symhd), gen1488)

						gen1490 := Call(__e, ShenFunc(symtl), V385)

						gen1491 := Call(__e, ShenFunc(symhd), gen1490)

						gen1492 := Call(__e, ShenFunc(symhd), V385)

						gen1493 := Call(__e, ShenFunc(symtl), gen1492)

						gen1494 := Call(__e, ShenFunc(symhd), gen1493)

						gen1495 := Call(__e, ShenFunc(symhd), V385)

						gen1496 := Call(__e, ShenFunc(symtl), gen1495)

						gen1497 := Call(__e, ShenFunc(symtl), gen1496)

						gen1498 := Call(__e, ShenFunc(symhd), gen1497)

						gen1499 := Call(__e, ShenFunc(symshen_4ebr), gen1491, gen1494, gen1498)

						gen1500 := Call(__e, ShenFunc(symcons), gen1499, Nil)

						gen1501 := Call(__e, ShenFunc(symcons), gen1489, gen1500)

						gen1502 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1501)

						gen1503 := Call(__e, ShenFunc(symcons), gen1502, Nil)

						gen1504 := Call(__e, ShenFunc(symcons), gen1483, gen1503)

						gen1505 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1504)

						Abstraction := gen1505
						gen1506 := Call(__e, ShenFunc(symtl), V385)

						gen1507 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen1506)

						gen1508 := Call(__e, ShenFunc(symcons), gen1507, Nil)

						gen1509 := Call(__e, ShenFunc(symcons), Abstraction, gen1508)

						gen1510 := Call(__e, ShenFunc(symtl), V385)

						gen1511 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen1510)

						gen1512 := Call(__e, ShenFunc(symcons), gen1511, Nil)

						gen1513 := Call(__e, ShenFunc(symcons), gen1509, gen1512)

						Application := gen1513
						__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

						return

					} else {
						gen1475 := Call(__e, ShenFunc(symcons_2), V385)

						var gen1476 Obj
						if True == gen1475 {
							gen1472 := Call(__e, ShenFunc(symhd), V385)

							gen1473 := Call(__e, ShenFunc(symcons_2), gen1472)

							var gen1474 Obj
							if True == gen1473 {
								gen1468 := Call(__e, ShenFunc(symhd), V385)

								gen1469 := Call(__e, ShenFunc(symhd), gen1468)

								gen1470 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1469)

								var gen1471 Obj
								if True == gen1470 {
									gen1464 := Call(__e, ShenFunc(symhd), V385)

									gen1465 := Call(__e, ShenFunc(symtl), gen1464)

									gen1466 := Call(__e, ShenFunc(symcons_2), gen1465)

									var gen1467 Obj
									if True == gen1466 {
										gen1459 := Call(__e, ShenFunc(symhd), V385)

										gen1460 := Call(__e, ShenFunc(symtl), gen1459)

										gen1461 := Call(__e, ShenFunc(symhd), gen1460)

										gen1462 := Call(__e, ShenFunc(symcons_2), gen1461)

										var gen1463 Obj
										if True == gen1462 {
											gen1453 := Call(__e, ShenFunc(symhd), V385)

											gen1454 := Call(__e, ShenFunc(symtl), gen1453)

											gen1455 := Call(__e, ShenFunc(symhd), gen1454)

											gen1456 := Call(__e, ShenFunc(symhd), gen1455)

											gen1457 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen1456)

											var gen1458 Obj
											if True == gen1457 {
												gen1447 := Call(__e, ShenFunc(symhd), V385)

												gen1448 := Call(__e, ShenFunc(symtl), gen1447)

												gen1449 := Call(__e, ShenFunc(symhd), gen1448)

												gen1450 := Call(__e, ShenFunc(symtl), gen1449)

												gen1451 := Call(__e, ShenFunc(symcons_2), gen1450)

												var gen1452 Obj
												if True == gen1451 {
													gen1440 := Call(__e, ShenFunc(symhd), V385)

													gen1441 := Call(__e, ShenFunc(symtl), gen1440)

													gen1442 := Call(__e, ShenFunc(symhd), gen1441)

													gen1443 := Call(__e, ShenFunc(symtl), gen1442)

													gen1444 := Call(__e, ShenFunc(symtl), gen1443)

													gen1445 := Call(__e, ShenFunc(symcons_2), gen1444)

													var gen1446 Obj
													if True == gen1445 {
														gen1432 := Call(__e, ShenFunc(symhd), V385)

														gen1433 := Call(__e, ShenFunc(symtl), gen1432)

														gen1434 := Call(__e, ShenFunc(symhd), gen1433)

														gen1435 := Call(__e, ShenFunc(symtl), gen1434)

														gen1436 := Call(__e, ShenFunc(symtl), gen1435)

														gen1437 := Call(__e, ShenFunc(symtl), gen1436)

														gen1438 := Call(__e, ShenFunc(sym_a), Nil, gen1437)

														var gen1439 Obj
														if True == gen1438 {
															gen1427 := Call(__e, ShenFunc(symhd), V385)

															gen1428 := Call(__e, ShenFunc(symtl), gen1427)

															gen1429 := Call(__e, ShenFunc(symtl), gen1428)

															gen1430 := Call(__e, ShenFunc(symcons_2), gen1429)

															var gen1431 Obj
															if True == gen1430 {
																gen1421 := Call(__e, ShenFunc(symhd), V385)

																gen1422 := Call(__e, ShenFunc(symtl), gen1421)

																gen1423 := Call(__e, ShenFunc(symtl), gen1422)

																gen1424 := Call(__e, ShenFunc(symtl), gen1423)

																gen1425 := Call(__e, ShenFunc(sym_a), Nil, gen1424)

																var gen1426 Obj
																if True == gen1425 {
																	gen1418 := Call(__e, ShenFunc(symtl), V385)

																	gen1419 := Call(__e, ShenFunc(symcons_2), gen1418)

																	var gen1420 Obj
																	if True == gen1419 {
																		gen1415 := Call(__e, ShenFunc(symtl), V385)

																		gen1416 := Call(__e, ShenFunc(symtl), gen1415)

																		gen1417 := Call(__e, ShenFunc(sym_a), Nil, gen1416)

																		if True == gen1417 {
																			gen1420 = True
																		} else {
																			gen1420 = False
																		}

																	} else {
																		gen1420 = False
																	}
																	if True == gen1420 {
																		gen1426 = True
																	} else {
																		gen1426 = False
																	}

																} else {
																	gen1426 = False
																}
																if True == gen1426 {
																	gen1431 = True
																} else {
																	gen1431 = False
																}

															} else {
																gen1431 = False
															}
															if True == gen1431 {
																gen1439 = True
															} else {
																gen1439 = False
															}

														} else {
															gen1439 = False
														}
														if True == gen1439 {
															gen1446 = True
														} else {
															gen1446 = False
														}

													} else {
														gen1446 = False
													}
													if True == gen1446 {
														gen1452 = True
													} else {
														gen1452 = False
													}

												} else {
													gen1452 = False
												}
												if True == gen1452 {
													gen1458 = True
												} else {
													gen1458 = False
												}

											} else {
												gen1458 = False
											}
											if True == gen1458 {
												gen1463 = True
											} else {
												gen1463 = False
											}

										} else {
											gen1463 = False
										}
										if True == gen1463 {
											gen1467 = True
										} else {
											gen1467 = False
										}

									} else {
										gen1467 = False
									}
									if True == gen1467 {
										gen1471 = True
									} else {
										gen1471 = False
									}

								} else {
									gen1471 = False
								}
								if True == gen1471 {
									gen1474 = True
								} else {
									gen1474 = False
								}

							} else {
								gen1474 = False
							}
							if True == gen1474 {
								gen1476 = True
							} else {
								gen1476 = False
							}

						} else {
							gen1476 = False
						}
						if True == gen1476 {
							gen1375 := Call(__e, ShenFunc(symtl), V385)

							gen1376 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.+string?"), gen1375)

							Call(__e, ShenFunc(symshen_4add__test), gen1376)

							gen1377 := Call(__e, ShenFunc(symhd), V385)

							gen1378 := Call(__e, ShenFunc(symtl), gen1377)

							gen1379 := Call(__e, ShenFunc(symhd), gen1378)

							gen1380 := Call(__e, ShenFunc(symtl), gen1379)

							gen1381 := Call(__e, ShenFunc(symhd), gen1380)

							gen1382 := Call(__e, ShenFunc(symhd), V385)

							gen1383 := Call(__e, ShenFunc(symtl), gen1382)

							gen1384 := Call(__e, ShenFunc(symhd), gen1383)

							gen1385 := Call(__e, ShenFunc(symtl), gen1384)

							gen1386 := Call(__e, ShenFunc(symtl), gen1385)

							gen1387 := Call(__e, ShenFunc(symhd), gen1386)

							gen1388 := Call(__e, ShenFunc(symtl), V385)

							gen1389 := Call(__e, ShenFunc(symhd), gen1388)

							gen1390 := Call(__e, ShenFunc(symhd), V385)

							gen1391 := Call(__e, ShenFunc(symtl), gen1390)

							gen1392 := Call(__e, ShenFunc(symhd), gen1391)

							gen1393 := Call(__e, ShenFunc(symhd), V385)

							gen1394 := Call(__e, ShenFunc(symtl), gen1393)

							gen1395 := Call(__e, ShenFunc(symtl), gen1394)

							gen1396 := Call(__e, ShenFunc(symhd), gen1395)

							gen1397 := Call(__e, ShenFunc(symshen_4ebr), gen1389, gen1392, gen1396)

							gen1398 := Call(__e, ShenFunc(symcons), gen1397, Nil)

							gen1399 := Call(__e, ShenFunc(symcons), gen1387, gen1398)

							gen1400 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1399)

							gen1401 := Call(__e, ShenFunc(symcons), gen1400, Nil)

							gen1402 := Call(__e, ShenFunc(symcons), gen1381, gen1401)

							gen1403 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen1402)

							Abstraction := gen1403
							gen1404 := Call(__e, ShenFunc(symtl), V385)

							gen1405 := Call(__e, ShenFunc(symhd), gen1404)

							gen1406 := Call(__e, ShenFunc(symcons), MakeNumber(0), Nil)

							gen1407 := Call(__e, ShenFunc(symcons), gen1405, gen1406)

							gen1408 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen1407)

							gen1409 := Call(__e, ShenFunc(symcons), gen1408, Nil)

							gen1410 := Call(__e, ShenFunc(symcons), Abstraction, gen1409)

							gen1411 := Call(__e, ShenFunc(symtl), V385)

							gen1412 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen1411)

							gen1413 := Call(__e, ShenFunc(symcons), gen1412, Nil)

							gen1414 := Call(__e, ShenFunc(symcons), gen1410, gen1413)

							Application := gen1414
							__e.TailApply(ShenFunc(symshen_4reduce__help), Application)

							return

						} else {
							gen1373 := Call(__e, ShenFunc(symcons_2), V385)

							var gen1374 Obj
							if True == gen1373 {
								gen1370 := Call(__e, ShenFunc(symhd), V385)

								gen1371 := Call(__e, ShenFunc(symcons_2), gen1370)

								var gen1372 Obj
								if True == gen1371 {
									gen1366 := Call(__e, ShenFunc(symhd), V385)

									gen1367 := Call(__e, ShenFunc(symhd), gen1366)

									gen1368 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1367)

									var gen1369 Obj
									if True == gen1368 {
										gen1362 := Call(__e, ShenFunc(symhd), V385)

										gen1363 := Call(__e, ShenFunc(symtl), gen1362)

										gen1364 := Call(__e, ShenFunc(symcons_2), gen1363)

										var gen1365 Obj
										if True == gen1364 {
											gen1357 := Call(__e, ShenFunc(symhd), V385)

											gen1358 := Call(__e, ShenFunc(symtl), gen1357)

											gen1359 := Call(__e, ShenFunc(symhd), gen1358)

											gen1360 := Call(__e, ShenFunc(symcons_2), gen1359)

											var gen1361 Obj
											if True == gen1360 {
												gen1352 := Call(__e, ShenFunc(symhd), V385)

												gen1353 := Call(__e, ShenFunc(symtl), gen1352)

												gen1354 := Call(__e, ShenFunc(symtl), gen1353)

												gen1355 := Call(__e, ShenFunc(symcons_2), gen1354)

												var gen1356 Obj
												if True == gen1355 {
													gen1346 := Call(__e, ShenFunc(symhd), V385)

													gen1347 := Call(__e, ShenFunc(symtl), gen1346)

													gen1348 := Call(__e, ShenFunc(symtl), gen1347)

													gen1349 := Call(__e, ShenFunc(symtl), gen1348)

													gen1350 := Call(__e, ShenFunc(sym_a), Nil, gen1349)

													var gen1351 Obj
													if True == gen1350 {
														gen1343 := Call(__e, ShenFunc(symtl), V385)

														gen1344 := Call(__e, ShenFunc(symcons_2), gen1343)

														var gen1345 Obj
														if True == gen1344 {
															gen1340 := Call(__e, ShenFunc(symtl), V385)

															gen1341 := Call(__e, ShenFunc(symtl), gen1340)

															gen1342 := Call(__e, ShenFunc(sym_a), Nil, gen1341)

															if True == gen1342 {
																gen1345 = True
															} else {
																gen1345 = False
															}

														} else {
															gen1345 = False
														}
														if True == gen1345 {
															gen1351 = True
														} else {
															gen1351 = False
														}

													} else {
														gen1351 = False
													}
													if True == gen1351 {
														gen1356 = True
													} else {
														gen1356 = False
													}

												} else {
													gen1356 = False
												}
												if True == gen1356 {
													gen1361 = True
												} else {
													gen1361 = False
												}

											} else {
												gen1361 = False
											}
											if True == gen1361 {
												gen1365 = True
											} else {
												gen1365 = False
											}

										} else {
											gen1365 = False
										}
										if True == gen1365 {
											gen1369 = True
										} else {
											gen1369 = False
										}

									} else {
										gen1369 = False
									}
									if True == gen1369 {
										gen1372 = True
									} else {
										gen1372 = False
									}

								} else {
									gen1372 = False
								}
								if True == gen1372 {
									gen1374 = True
								} else {
									gen1374 = False
								}

							} else {
								gen1374 = False
							}
							if True == gen1374 {
								__e.TailApply(ShenFunc(symshen_4custom_1pattern_1reducer), V385)

								return
							} else {
								gen1338 := Call(__e, ShenFunc(symcons_2), V385)

								var gen1339 Obj
								if True == gen1338 {
									gen1335 := Call(__e, ShenFunc(symhd), V385)

									gen1336 := Call(__e, ShenFunc(symcons_2), gen1335)

									var gen1337 Obj
									if True == gen1336 {
										gen1331 := Call(__e, ShenFunc(symhd), V385)

										gen1332 := Call(__e, ShenFunc(symhd), gen1331)

										gen1333 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1332)

										var gen1334 Obj
										if True == gen1333 {
											gen1327 := Call(__e, ShenFunc(symhd), V385)

											gen1328 := Call(__e, ShenFunc(symtl), gen1327)

											gen1329 := Call(__e, ShenFunc(symcons_2), gen1328)

											var gen1330 Obj
											if True == gen1329 {
												gen1322 := Call(__e, ShenFunc(symhd), V385)

												gen1323 := Call(__e, ShenFunc(symtl), gen1322)

												gen1324 := Call(__e, ShenFunc(symtl), gen1323)

												gen1325 := Call(__e, ShenFunc(symcons_2), gen1324)

												var gen1326 Obj
												if True == gen1325 {
													gen1316 := Call(__e, ShenFunc(symhd), V385)

													gen1317 := Call(__e, ShenFunc(symtl), gen1316)

													gen1318 := Call(__e, ShenFunc(symtl), gen1317)

													gen1319 := Call(__e, ShenFunc(symtl), gen1318)

													gen1320 := Call(__e, ShenFunc(sym_a), Nil, gen1319)

													var gen1321 Obj
													if True == gen1320 {
														gen1313 := Call(__e, ShenFunc(symtl), V385)

														gen1314 := Call(__e, ShenFunc(symcons_2), gen1313)

														var gen1315 Obj
														if True == gen1314 {
															gen1309 := Call(__e, ShenFunc(symtl), V385)

															gen1310 := Call(__e, ShenFunc(symtl), gen1309)

															gen1311 := Call(__e, ShenFunc(sym_a), Nil, gen1310)

															var gen1312 Obj
															if True == gen1311 {
																gen1304 := Call(__e, ShenFunc(symhd), V385)

																gen1305 := Call(__e, ShenFunc(symtl), gen1304)

																gen1306 := Call(__e, ShenFunc(symhd), gen1305)

																gen1307 := Call(__e, ShenFunc(symvariable_2), gen1306)

																gen1308 := Call(__e, ShenFunc(symnot), gen1307)

																if True == gen1308 {
																	gen1312 = True
																} else {
																	gen1312 = False
																}

															} else {
																gen1312 = False
															}
															if True == gen1312 {
																gen1315 = True
															} else {
																gen1315 = False
															}

														} else {
															gen1315 = False
														}
														if True == gen1315 {
															gen1321 = True
														} else {
															gen1321 = False
														}

													} else {
														gen1321 = False
													}
													if True == gen1321 {
														gen1326 = True
													} else {
														gen1326 = False
													}

												} else {
													gen1326 = False
												}
												if True == gen1326 {
													gen1330 = True
												} else {
													gen1330 = False
												}

											} else {
												gen1330 = False
											}
											if True == gen1330 {
												gen1334 = True
											} else {
												gen1334 = False
											}

										} else {
											gen1334 = False
										}
										if True == gen1334 {
											gen1337 = True
										} else {
											gen1337 = False
										}

									} else {
										gen1337 = False
									}
									if True == gen1337 {
										gen1339 = True
									} else {
										gen1339 = False
									}

								} else {
									gen1339 = False
								}
								if True == gen1339 {
									gen1294 := Call(__e, ShenFunc(symhd), V385)

									gen1295 := Call(__e, ShenFunc(symtl), gen1294)

									gen1296 := Call(__e, ShenFunc(symhd), gen1295)

									gen1297 := Call(__e, ShenFunc(symtl), V385)

									gen1298 := Call(__e, ShenFunc(symcons), gen1296, gen1297)

									gen1299 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen1298)

									Call(__e, ShenFunc(symshen_4add__test), gen1299)

									gen1300 := Call(__e, ShenFunc(symhd), V385)

									gen1301 := Call(__e, ShenFunc(symtl), gen1300)

									gen1302 := Call(__e, ShenFunc(symtl), gen1301)

									gen1303 := Call(__e, ShenFunc(symhd), gen1302)

									__e.TailApply(ShenFunc(symshen_4reduce__help), gen1303)

									return

								} else {
									gen1292 := Call(__e, ShenFunc(symcons_2), V385)

									var gen1293 Obj
									if True == gen1292 {
										gen1289 := Call(__e, ShenFunc(symhd), V385)

										gen1290 := Call(__e, ShenFunc(symcons_2), gen1289)

										var gen1291 Obj
										if True == gen1290 {
											gen1285 := Call(__e, ShenFunc(symhd), V385)

											gen1286 := Call(__e, ShenFunc(symhd), gen1285)

											gen1287 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen1286)

											var gen1288 Obj
											if True == gen1287 {
												gen1281 := Call(__e, ShenFunc(symhd), V385)

												gen1282 := Call(__e, ShenFunc(symtl), gen1281)

												gen1283 := Call(__e, ShenFunc(symcons_2), gen1282)

												var gen1284 Obj
												if True == gen1283 {
													gen1276 := Call(__e, ShenFunc(symhd), V385)

													gen1277 := Call(__e, ShenFunc(symtl), gen1276)

													gen1278 := Call(__e, ShenFunc(symtl), gen1277)

													gen1279 := Call(__e, ShenFunc(symcons_2), gen1278)

													var gen1280 Obj
													if True == gen1279 {
														gen1270 := Call(__e, ShenFunc(symhd), V385)

														gen1271 := Call(__e, ShenFunc(symtl), gen1270)

														gen1272 := Call(__e, ShenFunc(symtl), gen1271)

														gen1273 := Call(__e, ShenFunc(symtl), gen1272)

														gen1274 := Call(__e, ShenFunc(sym_a), Nil, gen1273)

														var gen1275 Obj
														if True == gen1274 {
															gen1267 := Call(__e, ShenFunc(symtl), V385)

															gen1268 := Call(__e, ShenFunc(symcons_2), gen1267)

															var gen1269 Obj
															if True == gen1268 {
																gen1264 := Call(__e, ShenFunc(symtl), V385)

																gen1265 := Call(__e, ShenFunc(symtl), gen1264)

																gen1266 := Call(__e, ShenFunc(sym_a), Nil, gen1265)

																if True == gen1266 {
																	gen1269 = True
																} else {
																	gen1269 = False
																}

															} else {
																gen1269 = False
															}
															if True == gen1269 {
																gen1275 = True
															} else {
																gen1275 = False
															}

														} else {
															gen1275 = False
														}
														if True == gen1275 {
															gen1280 = True
														} else {
															gen1280 = False
														}

													} else {
														gen1280 = False
													}
													if True == gen1280 {
														gen1284 = True
													} else {
														gen1284 = False
													}

												} else {
													gen1284 = False
												}
												if True == gen1284 {
													gen1288 = True
												} else {
													gen1288 = False
												}

											} else {
												gen1288 = False
											}
											if True == gen1288 {
												gen1291 = True
											} else {
												gen1291 = False
											}

										} else {
											gen1291 = False
										}
										if True == gen1291 {
											gen1293 = True
										} else {
											gen1293 = False
										}

									} else {
										gen1293 = False
									}
									if True == gen1293 {
										gen1254 := Call(__e, ShenFunc(symtl), V385)

										gen1255 := Call(__e, ShenFunc(symhd), gen1254)

										gen1256 := Call(__e, ShenFunc(symhd), V385)

										gen1257 := Call(__e, ShenFunc(symtl), gen1256)

										gen1258 := Call(__e, ShenFunc(symhd), gen1257)

										gen1259 := Call(__e, ShenFunc(symhd), V385)

										gen1260 := Call(__e, ShenFunc(symtl), gen1259)

										gen1261 := Call(__e, ShenFunc(symtl), gen1260)

										gen1262 := Call(__e, ShenFunc(symhd), gen1261)

										gen1263 := Call(__e, ShenFunc(symshen_4ebr), gen1255, gen1258, gen1262)

										__e.TailApply(ShenFunc(symshen_4reduce__help), gen1263)

										return

									} else {
										gen1252 := Call(__e, ShenFunc(symcons_2), V385)

										var gen1253 Obj
										if True == gen1252 {
											gen1249 := Call(__e, ShenFunc(symhd), V385)

											gen1250 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), gen1249)

											var gen1251 Obj
											if True == gen1250 {
												gen1246 := Call(__e, ShenFunc(symtl), V385)

												gen1247 := Call(__e, ShenFunc(symcons_2), gen1246)

												var gen1248 Obj
												if True == gen1247 {
													gen1242 := Call(__e, ShenFunc(symtl), V385)

													gen1243 := Call(__e, ShenFunc(symtl), gen1242)

													gen1244 := Call(__e, ShenFunc(symcons_2), gen1243)

													var gen1245 Obj
													if True == gen1244 {
														gen1238 := Call(__e, ShenFunc(symtl), V385)

														gen1239 := Call(__e, ShenFunc(symtl), gen1238)

														gen1240 := Call(__e, ShenFunc(symtl), gen1239)

														gen1241 := Call(__e, ShenFunc(sym_a), Nil, gen1240)

														if True == gen1241 {
															gen1245 = True
														} else {
															gen1245 = False
														}

													} else {
														gen1245 = False
													}
													if True == gen1245 {
														gen1248 = True
													} else {
														gen1248 = False
													}

												} else {
													gen1248 = False
												}
												if True == gen1248 {
													gen1251 = True
												} else {
													gen1251 = False
												}

											} else {
												gen1251 = False
											}
											if True == gen1251 {
												gen1253 = True
											} else {
												gen1253 = False
											}

										} else {
											gen1253 = False
										}
										if True == gen1253 {
											gen1233 := Call(__e, ShenFunc(symtl), V385)

											gen1234 := Call(__e, ShenFunc(symhd), gen1233)

											Call(__e, ShenFunc(symshen_4add__test), gen1234)

											gen1235 := Call(__e, ShenFunc(symtl), V385)

											gen1236 := Call(__e, ShenFunc(symtl), gen1235)

											gen1237 := Call(__e, ShenFunc(symhd), gen1236)

											__e.TailApply(ShenFunc(symshen_4reduce__help), gen1237)

											return

										} else {
											gen1231 := Call(__e, ShenFunc(symcons_2), V385)

											var gen1232 Obj
											if True == gen1231 {
												gen1228 := Call(__e, ShenFunc(symtl), V385)

												gen1229 := Call(__e, ShenFunc(symcons_2), gen1228)

												var gen1230 Obj
												if True == gen1229 {
													gen1225 := Call(__e, ShenFunc(symtl), V385)

													gen1226 := Call(__e, ShenFunc(symtl), gen1225)

													gen1227 := Call(__e, ShenFunc(sym_a), Nil, gen1226)

													if True == gen1227 {
														gen1230 = True
													} else {
														gen1230 = False
													}

												} else {
													gen1230 = False
												}
												if True == gen1230 {
													gen1232 = True
												} else {
													gen1232 = False
												}

											} else {
												gen1232 = False
											}
											if True == gen1232 {
												gen1219 := Call(__e, ShenFunc(symhd), V385)

												gen1220 := Call(__e, ShenFunc(symshen_4reduce__help), gen1219)

												Z := gen1220
												gen1223 := Call(__e, ShenFunc(symhd), V385)

												gen1224 := Call(__e, ShenFunc(sym_a), gen1223, Z)

												if True == gen1224 {
													__e.Return(V385)
													return
												} else {
													gen1221 := Call(__e, ShenFunc(symtl), V385)

													gen1222 := Call(__e, ShenFunc(symcons), Z, gen1221)

													__e.TailApply(ShenFunc(symshen_4reduce__help), gen1222)

													return

												}

											} else {
												__e.Return(V385)
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

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.reduce_help"), gen1774)

		gen1776 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V387 := __args[0]
			_ = V387
			gen1775 := Call(__e, ShenFunc(sym_a), MakeString(""), V387)

			if True == gen1775 {
				__e.Return(False)
				return
			} else {
				__e.TailApply(ShenFunc(symstring_2), V387)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.+string?"), gen1776)

		gen1780 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V389 := __args[0]
			_ = V389
			gen1779 := Call(__e, ShenFunc(symabsvector_2), V389)

			if True == gen1779 {
				gen1777 := Call(__e, ShenFunc(sym_5_1address), V389, MakeNumber(0))

				gen1778 := Call(__e, ShenFunc(sym_6), gen1777, MakeNumber(0))

				if True == gen1778 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.+vector?"), gen1780)

		gen1844 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V402 := __args[0]
			_ = V402
			V403 := __args[1]
			_ = V403
			V404 := __args[2]
			_ = V404
			gen1843 := Call(__e, ShenFunc(sym_a), V404, V403)

			if True == gen1843 {
				__e.Return(V402)
				return
			} else {
				gen1841 := Call(__e, ShenFunc(symcons_2), V404)

				var gen1842 Obj
				if True == gen1841 {
					gen1838 := Call(__e, ShenFunc(symhd), V404)

					gen1839 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), gen1838)

					var gen1840 Obj
					if True == gen1839 {
						gen1835 := Call(__e, ShenFunc(symtl), V404)

						gen1836 := Call(__e, ShenFunc(symcons_2), gen1835)

						var gen1837 Obj
						if True == gen1836 {
							gen1831 := Call(__e, ShenFunc(symtl), V404)

							gen1832 := Call(__e, ShenFunc(symtl), gen1831)

							gen1833 := Call(__e, ShenFunc(symcons_2), gen1832)

							var gen1834 Obj
							if True == gen1833 {
								gen1826 := Call(__e, ShenFunc(symtl), V404)

								gen1827 := Call(__e, ShenFunc(symtl), gen1826)

								gen1828 := Call(__e, ShenFunc(symtl), gen1827)

								gen1829 := Call(__e, ShenFunc(sym_a), Nil, gen1828)

								var gen1830 Obj
								if True == gen1829 {
									gen1823 := Call(__e, ShenFunc(symtl), V404)

									gen1824 := Call(__e, ShenFunc(symhd), gen1823)

									gen1825 := Call(__e, ShenFunc(symshen_4clash_2), gen1824, V403)

									if True == gen1825 {
										gen1830 = True
									} else {
										gen1830 = False
									}

								} else {
									gen1830 = False
								}
								if True == gen1830 {
									gen1834 = True
								} else {
									gen1834 = False
								}

							} else {
								gen1834 = False
							}
							if True == gen1834 {
								gen1837 = True
							} else {
								gen1837 = False
							}

						} else {
							gen1837 = False
						}
						if True == gen1837 {
							gen1840 = True
						} else {
							gen1840 = False
						}

					} else {
						gen1840 = False
					}
					if True == gen1840 {
						gen1842 = True
					} else {
						gen1842 = False
					}

				} else {
					gen1842 = False
				}
				if True == gen1842 {
					__e.Return(V404)
					return
				} else {
					gen1821 := Call(__e, ShenFunc(symcons_2), V404)

					var gen1822 Obj
					if True == gen1821 {
						gen1818 := Call(__e, ShenFunc(symhd), V404)

						gen1819 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen1818)

						var gen1820 Obj
						if True == gen1819 {
							gen1815 := Call(__e, ShenFunc(symtl), V404)

							gen1816 := Call(__e, ShenFunc(symcons_2), gen1815)

							var gen1817 Obj
							if True == gen1816 {
								gen1811 := Call(__e, ShenFunc(symtl), V404)

								gen1812 := Call(__e, ShenFunc(symtl), gen1811)

								gen1813 := Call(__e, ShenFunc(symcons_2), gen1812)

								var gen1814 Obj
								if True == gen1813 {
									gen1806 := Call(__e, ShenFunc(symtl), V404)

									gen1807 := Call(__e, ShenFunc(symtl), gen1806)

									gen1808 := Call(__e, ShenFunc(symtl), gen1807)

									gen1809 := Call(__e, ShenFunc(symcons_2), gen1808)

									var gen1810 Obj
									if True == gen1809 {
										gen1800 := Call(__e, ShenFunc(symtl), V404)

										gen1801 := Call(__e, ShenFunc(symtl), gen1800)

										gen1802 := Call(__e, ShenFunc(symtl), gen1801)

										gen1803 := Call(__e, ShenFunc(symtl), gen1802)

										gen1804 := Call(__e, ShenFunc(sym_a), Nil, gen1803)

										var gen1805 Obj
										if True == gen1804 {
											gen1797 := Call(__e, ShenFunc(symtl), V404)

											gen1798 := Call(__e, ShenFunc(symhd), gen1797)

											gen1799 := Call(__e, ShenFunc(symshen_4clash_2), gen1798, V403)

											if True == gen1799 {
												gen1805 = True
											} else {
												gen1805 = False
											}

										} else {
											gen1805 = False
										}
										if True == gen1805 {
											gen1810 = True
										} else {
											gen1810 = False
										}

									} else {
										gen1810 = False
									}
									if True == gen1810 {
										gen1814 = True
									} else {
										gen1814 = False
									}

								} else {
									gen1814 = False
								}
								if True == gen1814 {
									gen1817 = True
								} else {
									gen1817 = False
								}

							} else {
								gen1817 = False
							}
							if True == gen1817 {
								gen1820 = True
							} else {
								gen1820 = False
							}

						} else {
							gen1820 = False
						}
						if True == gen1820 {
							gen1822 = True
						} else {
							gen1822 = False
						}

					} else {
						gen1822 = False
					}
					if True == gen1822 {
						gen1786 := Call(__e, ShenFunc(symtl), V404)

						gen1787 := Call(__e, ShenFunc(symhd), gen1786)

						gen1788 := Call(__e, ShenFunc(symtl), V404)

						gen1789 := Call(__e, ShenFunc(symtl), gen1788)

						gen1790 := Call(__e, ShenFunc(symhd), gen1789)

						gen1791 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen1790)

						gen1792 := Call(__e, ShenFunc(symtl), V404)

						gen1793 := Call(__e, ShenFunc(symtl), gen1792)

						gen1794 := Call(__e, ShenFunc(symtl), gen1793)

						gen1795 := Call(__e, ShenFunc(symcons), gen1791, gen1794)

						gen1796 := Call(__e, ShenFunc(symcons), gen1787, gen1795)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen1796)

						return

					} else {
						gen1785 := Call(__e, ShenFunc(symcons_2), V404)

						if True == gen1785 {
							gen1781 := Call(__e, ShenFunc(symhd), V404)

							gen1782 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen1781)

							gen1783 := Call(__e, ShenFunc(symtl), V404)

							gen1784 := Call(__e, ShenFunc(symshen_4ebr), V402, V403, gen1783)

							__e.TailApply(ShenFunc(symcons), gen1782, gen1784)

							return

						} else {
							__e.Return(V404)
							return
						}

					}

				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ebr"), gen1844)

		gen1851 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V416 := __args[0]
			_ = V416
			V417 := __args[1]
			_ = V417
			gen1850 := Call(__e, ShenFunc(sym_a), V417, V416)

			if True == gen1850 {
				__e.Return(True)
				return
			} else {
				gen1849 := Call(__e, ShenFunc(symcons_2), V417)

				if True == gen1849 {
					gen1847 := Call(__e, ShenFunc(symhd), V417)

					gen1848 := Call(__e, ShenFunc(symshen_4clash_2), V416, gen1847)

					if True == gen1848 {
						__e.Return(True)
						return
					} else {
						gen1845 := Call(__e, ShenFunc(symtl), V417)

						gen1846 := Call(__e, ShenFunc(symshen_4clash_2), V416, gen1845)

						if True == gen1846 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.clash?"), gen1851)

		gen1854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V419 := __args[0]
			_ = V419
			gen1852 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*teststack*"))

			gen1853 := Call(__e, ShenFunc(symcons), V419, gen1852)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*teststack*"), gen1853)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.add_test"), gen1854)

		gen1858 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V423 := __args[0]
			_ = V423
			V424 := __args[1]
			_ = V424
			V425 := __args[2]
			_ = V425
			gen1855 := Call(__e, ShenFunc(symshen_4err_1condition), V423)

			Err := gen1855
			gen1856 := Call(__e, ShenFunc(symshen_4case_1form), V425, Err)

			Cases := gen1856
			gen1857 := Call(__e, ShenFunc(symshen_4encode_1choices), Cases, V423)

			EncodeChoices := gen1857
			__e.TailApply(ShenFunc(symshen_4cond_1form), EncodeChoices)

			return

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cond-expression"), gen1858)

		gen1878 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V429 := __args[0]
			_ = V429
			gen1876 := Call(__e, ShenFunc(symcons_2), V429)

			var gen1877 Obj
			if True == gen1876 {
				gen1873 := Call(__e, ShenFunc(symhd), V429)

				gen1874 := Call(__e, ShenFunc(symcons_2), gen1873)

				var gen1875 Obj
				if True == gen1874 {
					gen1869 := Call(__e, ShenFunc(symhd), V429)

					gen1870 := Call(__e, ShenFunc(symhd), gen1869)

					gen1871 := Call(__e, ShenFunc(sym_a), True, gen1870)

					var gen1872 Obj
					if True == gen1871 {
						gen1865 := Call(__e, ShenFunc(symhd), V429)

						gen1866 := Call(__e, ShenFunc(symtl), gen1865)

						gen1867 := Call(__e, ShenFunc(symcons_2), gen1866)

						var gen1868 Obj
						if True == gen1867 {
							gen1861 := Call(__e, ShenFunc(symhd), V429)

							gen1862 := Call(__e, ShenFunc(symtl), gen1861)

							gen1863 := Call(__e, ShenFunc(symtl), gen1862)

							gen1864 := Call(__e, ShenFunc(sym_a), Nil, gen1863)

							if True == gen1864 {
								gen1868 = True
							} else {
								gen1868 = False
							}

						} else {
							gen1868 = False
						}
						if True == gen1868 {
							gen1872 = True
						} else {
							gen1872 = False
						}

					} else {
						gen1872 = False
					}
					if True == gen1872 {
						gen1875 = True
					} else {
						gen1875 = False
					}

				} else {
					gen1875 = False
				}
				if True == gen1875 {
					gen1877 = True
				} else {
					gen1877 = False
				}

			} else {
				gen1877 = False
			}
			if True == gen1877 {
				gen1859 := Call(__e, ShenFunc(symhd), V429)

				gen1860 := Call(__e, ShenFunc(symtl), gen1859)

				__e.TailApply(ShenFunc(symhd), gen1860)

				return

			} else {
				__e.TailApply(ShenFunc(symcons), MakeSymbol("cond"), V429)

				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cond-form"), gen1878)

		gen2101 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V434 := __args[0]
			_ = V434
			V435 := __args[1]
			_ = V435
			gen2100 := Call(__e, ShenFunc(sym_a), Nil, V434)

			if True == gen2100 {
				__e.Return(Nil)
				return
			} else {
				gen2098 := Call(__e, ShenFunc(symcons_2), V434)

				var gen2099 Obj
				if True == gen2098 {
					gen2095 := Call(__e, ShenFunc(symhd), V434)

					gen2096 := Call(__e, ShenFunc(symcons_2), gen2095)

					var gen2097 Obj
					if True == gen2096 {
						gen2091 := Call(__e, ShenFunc(symhd), V434)

						gen2092 := Call(__e, ShenFunc(symhd), gen2091)

						gen2093 := Call(__e, ShenFunc(sym_a), True, gen2092)

						var gen2094 Obj
						if True == gen2093 {
							gen2087 := Call(__e, ShenFunc(symhd), V434)

							gen2088 := Call(__e, ShenFunc(symtl), gen2087)

							gen2089 := Call(__e, ShenFunc(symcons_2), gen2088)

							var gen2090 Obj
							if True == gen2089 {
								gen2082 := Call(__e, ShenFunc(symhd), V434)

								gen2083 := Call(__e, ShenFunc(symtl), gen2082)

								gen2084 := Call(__e, ShenFunc(symhd), gen2083)

								gen2085 := Call(__e, ShenFunc(symcons_2), gen2084)

								var gen2086 Obj
								if True == gen2085 {
									gen2076 := Call(__e, ShenFunc(symhd), V434)

									gen2077 := Call(__e, ShenFunc(symtl), gen2076)

									gen2078 := Call(__e, ShenFunc(symhd), gen2077)

									gen2079 := Call(__e, ShenFunc(symhd), gen2078)

									gen2080 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2079)

									var gen2081 Obj
									if True == gen2080 {
										gen2070 := Call(__e, ShenFunc(symhd), V434)

										gen2071 := Call(__e, ShenFunc(symtl), gen2070)

										gen2072 := Call(__e, ShenFunc(symhd), gen2071)

										gen2073 := Call(__e, ShenFunc(symtl), gen2072)

										gen2074 := Call(__e, ShenFunc(symcons_2), gen2073)

										var gen2075 Obj
										if True == gen2074 {
											gen2063 := Call(__e, ShenFunc(symhd), V434)

											gen2064 := Call(__e, ShenFunc(symtl), gen2063)

											gen2065 := Call(__e, ShenFunc(symhd), gen2064)

											gen2066 := Call(__e, ShenFunc(symtl), gen2065)

											gen2067 := Call(__e, ShenFunc(symtl), gen2066)

											gen2068 := Call(__e, ShenFunc(sym_a), Nil, gen2067)

											var gen2069 Obj
											if True == gen2068 {
												gen2058 := Call(__e, ShenFunc(symhd), V434)

												gen2059 := Call(__e, ShenFunc(symtl), gen2058)

												gen2060 := Call(__e, ShenFunc(symtl), gen2059)

												gen2061 := Call(__e, ShenFunc(sym_a), Nil, gen2060)

												var gen2062 Obj
												if True == gen2061 {
													gen2056 := Call(__e, ShenFunc(symtl), V434)

													gen2057 := Call(__e, ShenFunc(sym_a), Nil, gen2056)

													if True == gen2057 {
														gen2062 = True
													} else {
														gen2062 = False
													}

												} else {
													gen2062 = False
												}
												if True == gen2062 {
													gen2069 = True
												} else {
													gen2069 = False
												}

											} else {
												gen2069 = False
											}
											if True == gen2069 {
												gen2075 = True
											} else {
												gen2075 = False
											}

										} else {
											gen2075 = False
										}
										if True == gen2075 {
											gen2081 = True
										} else {
											gen2081 = False
										}

									} else {
										gen2081 = False
									}
									if True == gen2081 {
										gen2086 = True
									} else {
										gen2086 = False
									}

								} else {
									gen2086 = False
								}
								if True == gen2086 {
									gen2090 = True
								} else {
									gen2090 = False
								}

							} else {
								gen2090 = False
							}
							if True == gen2090 {
								gen2094 = True
							} else {
								gen2094 = False
							}

						} else {
							gen2094 = False
						}
						if True == gen2094 {
							gen2097 = True
						} else {
							gen2097 = False
						}

					} else {
						gen2097 = False
					}
					if True == gen2097 {
						gen2099 = True
					} else {
						gen2099 = False
					}

				} else {
					gen2099 = False
				}
				if True == gen2099 {
					gen2033 := Call(__e, ShenFunc(symhd), V434)

					gen2034 := Call(__e, ShenFunc(symtl), gen2033)

					gen2035 := Call(__e, ShenFunc(symhd), gen2034)

					gen2036 := Call(__e, ShenFunc(symtl), gen2035)

					gen2037 := Call(__e, ShenFunc(symhd), gen2036)

					gen2038 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

					gen2039 := Call(__e, ShenFunc(symcons), gen2038, Nil)

					gen2040 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2039)

					gen2041 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen2040)

					gen2044 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*installing-kl*"))

					var gen2045 Obj
					if True == gen2044 {
						gen2043 := Call(__e, ShenFunc(symcons), V435, Nil)

						gen2045 = Call(__e, ShenFunc(symcons), MakeSymbol("shen.sys-error"), gen2043)

					} else {
						gen2042 := Call(__e, ShenFunc(symcons), V435, Nil)

						gen2045 = Call(__e, ShenFunc(symcons), MakeSymbol("shen.f_error"), gen2042)

					}
					gen2046 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

					gen2047 := Call(__e, ShenFunc(symcons), gen2045, gen2046)

					gen2048 := Call(__e, ShenFunc(symcons), gen2041, gen2047)

					gen2049 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen2048)

					gen2050 := Call(__e, ShenFunc(symcons), gen2049, Nil)

					gen2051 := Call(__e, ShenFunc(symcons), gen2037, gen2050)

					gen2052 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen2051)

					gen2053 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen2052)

					gen2054 := Call(__e, ShenFunc(symcons), gen2053, Nil)

					gen2055 := Call(__e, ShenFunc(symcons), True, gen2054)

					__e.TailApply(ShenFunc(symcons), gen2055, Nil)

					return

				} else {
					gen2031 := Call(__e, ShenFunc(symcons_2), V434)

					var gen2032 Obj
					if True == gen2031 {
						gen2028 := Call(__e, ShenFunc(symhd), V434)

						gen2029 := Call(__e, ShenFunc(symcons_2), gen2028)

						var gen2030 Obj
						if True == gen2029 {
							gen2024 := Call(__e, ShenFunc(symhd), V434)

							gen2025 := Call(__e, ShenFunc(symhd), gen2024)

							gen2026 := Call(__e, ShenFunc(sym_a), True, gen2025)

							var gen2027 Obj
							if True == gen2026 {
								gen2020 := Call(__e, ShenFunc(symhd), V434)

								gen2021 := Call(__e, ShenFunc(symtl), gen2020)

								gen2022 := Call(__e, ShenFunc(symcons_2), gen2021)

								var gen2023 Obj
								if True == gen2022 {
									gen2015 := Call(__e, ShenFunc(symhd), V434)

									gen2016 := Call(__e, ShenFunc(symtl), gen2015)

									gen2017 := Call(__e, ShenFunc(symhd), gen2016)

									gen2018 := Call(__e, ShenFunc(symcons_2), gen2017)

									var gen2019 Obj
									if True == gen2018 {
										gen2009 := Call(__e, ShenFunc(symhd), V434)

										gen2010 := Call(__e, ShenFunc(symtl), gen2009)

										gen2011 := Call(__e, ShenFunc(symhd), gen2010)

										gen2012 := Call(__e, ShenFunc(symhd), gen2011)

										gen2013 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2012)

										var gen2014 Obj
										if True == gen2013 {
											gen2003 := Call(__e, ShenFunc(symhd), V434)

											gen2004 := Call(__e, ShenFunc(symtl), gen2003)

											gen2005 := Call(__e, ShenFunc(symhd), gen2004)

											gen2006 := Call(__e, ShenFunc(symtl), gen2005)

											gen2007 := Call(__e, ShenFunc(symcons_2), gen2006)

											var gen2008 Obj
											if True == gen2007 {
												gen1996 := Call(__e, ShenFunc(symhd), V434)

												gen1997 := Call(__e, ShenFunc(symtl), gen1996)

												gen1998 := Call(__e, ShenFunc(symhd), gen1997)

												gen1999 := Call(__e, ShenFunc(symtl), gen1998)

												gen2000 := Call(__e, ShenFunc(symtl), gen1999)

												gen2001 := Call(__e, ShenFunc(sym_a), Nil, gen2000)

												var gen2002 Obj
												if True == gen2001 {
													gen1992 := Call(__e, ShenFunc(symhd), V434)

													gen1993 := Call(__e, ShenFunc(symtl), gen1992)

													gen1994 := Call(__e, ShenFunc(symtl), gen1993)

													gen1995 := Call(__e, ShenFunc(sym_a), Nil, gen1994)

													if True == gen1995 {
														gen2002 = True
													} else {
														gen2002 = False
													}

												} else {
													gen2002 = False
												}
												if True == gen2002 {
													gen2008 = True
												} else {
													gen2008 = False
												}

											} else {
												gen2008 = False
											}
											if True == gen2008 {
												gen2014 = True
											} else {
												gen2014 = False
											}

										} else {
											gen2014 = False
										}
										if True == gen2014 {
											gen2019 = True
										} else {
											gen2019 = False
										}

									} else {
										gen2019 = False
									}
									if True == gen2019 {
										gen2023 = True
									} else {
										gen2023 = False
									}

								} else {
									gen2023 = False
								}
								if True == gen2023 {
									gen2027 = True
								} else {
									gen2027 = False
								}

							} else {
								gen2027 = False
							}
							if True == gen2027 {
								gen2030 = True
							} else {
								gen2030 = False
							}

						} else {
							gen2030 = False
						}
						if True == gen2030 {
							gen2032 = True
						} else {
							gen2032 = False
						}

					} else {
						gen2032 = False
					}
					if True == gen2032 {
						gen1970 := Call(__e, ShenFunc(symhd), V434)

						gen1971 := Call(__e, ShenFunc(symtl), gen1970)

						gen1972 := Call(__e, ShenFunc(symhd), gen1971)

						gen1973 := Call(__e, ShenFunc(symtl), gen1972)

						gen1974 := Call(__e, ShenFunc(symhd), gen1973)

						gen1975 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

						gen1976 := Call(__e, ShenFunc(symcons), gen1975, Nil)

						gen1977 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen1976)

						gen1978 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen1977)

						gen1979 := Call(__e, ShenFunc(symtl), V434)

						gen1980 := Call(__e, ShenFunc(symshen_4encode_1choices), gen1979, V435)

						gen1981 := Call(__e, ShenFunc(symshen_4cond_1form), gen1980)

						gen1982 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

						gen1983 := Call(__e, ShenFunc(symcons), gen1981, gen1982)

						gen1984 := Call(__e, ShenFunc(symcons), gen1978, gen1983)

						gen1985 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen1984)

						gen1986 := Call(__e, ShenFunc(symcons), gen1985, Nil)

						gen1987 := Call(__e, ShenFunc(symcons), gen1974, gen1986)

						gen1988 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen1987)

						gen1989 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen1988)

						gen1990 := Call(__e, ShenFunc(symcons), gen1989, Nil)

						gen1991 := Call(__e, ShenFunc(symcons), True, gen1990)

						__e.TailApply(ShenFunc(symcons), gen1991, Nil)

						return

					} else {
						gen1968 := Call(__e, ShenFunc(symcons_2), V434)

						var gen1969 Obj
						if True == gen1968 {
							gen1965 := Call(__e, ShenFunc(symhd), V434)

							gen1966 := Call(__e, ShenFunc(symcons_2), gen1965)

							var gen1967 Obj
							if True == gen1966 {
								gen1961 := Call(__e, ShenFunc(symhd), V434)

								gen1962 := Call(__e, ShenFunc(symtl), gen1961)

								gen1963 := Call(__e, ShenFunc(symcons_2), gen1962)

								var gen1964 Obj
								if True == gen1963 {
									gen1956 := Call(__e, ShenFunc(symhd), V434)

									gen1957 := Call(__e, ShenFunc(symtl), gen1956)

									gen1958 := Call(__e, ShenFunc(symhd), gen1957)

									gen1959 := Call(__e, ShenFunc(symcons_2), gen1958)

									var gen1960 Obj
									if True == gen1959 {
										gen1950 := Call(__e, ShenFunc(symhd), V434)

										gen1951 := Call(__e, ShenFunc(symtl), gen1950)

										gen1952 := Call(__e, ShenFunc(symhd), gen1951)

										gen1953 := Call(__e, ShenFunc(symhd), gen1952)

										gen1954 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen1953)

										var gen1955 Obj
										if True == gen1954 {
											gen1944 := Call(__e, ShenFunc(symhd), V434)

											gen1945 := Call(__e, ShenFunc(symtl), gen1944)

											gen1946 := Call(__e, ShenFunc(symhd), gen1945)

											gen1947 := Call(__e, ShenFunc(symtl), gen1946)

											gen1948 := Call(__e, ShenFunc(symcons_2), gen1947)

											var gen1949 Obj
											if True == gen1948 {
												gen1937 := Call(__e, ShenFunc(symhd), V434)

												gen1938 := Call(__e, ShenFunc(symtl), gen1937)

												gen1939 := Call(__e, ShenFunc(symhd), gen1938)

												gen1940 := Call(__e, ShenFunc(symtl), gen1939)

												gen1941 := Call(__e, ShenFunc(symtl), gen1940)

												gen1942 := Call(__e, ShenFunc(sym_a), Nil, gen1941)

												var gen1943 Obj
												if True == gen1942 {
													gen1933 := Call(__e, ShenFunc(symhd), V434)

													gen1934 := Call(__e, ShenFunc(symtl), gen1933)

													gen1935 := Call(__e, ShenFunc(symtl), gen1934)

													gen1936 := Call(__e, ShenFunc(sym_a), Nil, gen1935)

													if True == gen1936 {
														gen1943 = True
													} else {
														gen1943 = False
													}

												} else {
													gen1943 = False
												}
												if True == gen1943 {
													gen1949 = True
												} else {
													gen1949 = False
												}

											} else {
												gen1949 = False
											}
											if True == gen1949 {
												gen1955 = True
											} else {
												gen1955 = False
											}

										} else {
											gen1955 = False
										}
										if True == gen1955 {
											gen1960 = True
										} else {
											gen1960 = False
										}

									} else {
										gen1960 = False
									}
									if True == gen1960 {
										gen1964 = True
									} else {
										gen1964 = False
									}

								} else {
									gen1964 = False
								}
								if True == gen1964 {
									gen1967 = True
								} else {
									gen1967 = False
								}

							} else {
								gen1967 = False
							}
							if True == gen1967 {
								gen1969 = True
							} else {
								gen1969 = False
							}

						} else {
							gen1969 = False
						}
						if True == gen1969 {
							gen1895 := Call(__e, ShenFunc(symtl), V434)

							gen1896 := Call(__e, ShenFunc(symshen_4encode_1choices), gen1895, V435)

							gen1897 := Call(__e, ShenFunc(symshen_4cond_1form), gen1896)

							gen1898 := Call(__e, ShenFunc(symcons), gen1897, Nil)

							gen1899 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen1898)

							gen1900 := Call(__e, ShenFunc(symhd), V434)

							gen1901 := Call(__e, ShenFunc(symhd), gen1900)

							gen1902 := Call(__e, ShenFunc(symhd), V434)

							gen1903 := Call(__e, ShenFunc(symtl), gen1902)

							gen1904 := Call(__e, ShenFunc(symhd), gen1903)

							gen1905 := Call(__e, ShenFunc(symtl), gen1904)

							gen1906 := Call(__e, ShenFunc(symhd), gen1905)

							gen1907 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

							gen1908 := Call(__e, ShenFunc(symcons), gen1907, Nil)

							gen1909 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen1908)

							gen1910 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen1909)

							gen1911 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), Nil)

							gen1912 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen1911)

							gen1913 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

							gen1914 := Call(__e, ShenFunc(symcons), gen1912, gen1913)

							gen1915 := Call(__e, ShenFunc(symcons), gen1910, gen1914)

							gen1916 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen1915)

							gen1917 := Call(__e, ShenFunc(symcons), gen1916, Nil)

							gen1918 := Call(__e, ShenFunc(symcons), gen1906, gen1917)

							gen1919 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen1918)

							gen1920 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen1919)

							gen1921 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), Nil)

							gen1922 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen1921)

							gen1923 := Call(__e, ShenFunc(symcons), gen1922, Nil)

							gen1924 := Call(__e, ShenFunc(symcons), gen1920, gen1923)

							gen1925 := Call(__e, ShenFunc(symcons), gen1901, gen1924)

							gen1926 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen1925)

							gen1927 := Call(__e, ShenFunc(symcons), gen1926, Nil)

							gen1928 := Call(__e, ShenFunc(symcons), gen1899, gen1927)

							gen1929 := Call(__e, ShenFunc(symcons), MakeSymbol("Freeze"), gen1928)

							gen1930 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen1929)

							gen1931 := Call(__e, ShenFunc(symcons), gen1930, Nil)

							gen1932 := Call(__e, ShenFunc(symcons), True, gen1931)

							__e.TailApply(ShenFunc(symcons), gen1932, Nil)

							return

						} else {
							gen1893 := Call(__e, ShenFunc(symcons_2), V434)

							var gen1894 Obj
							if True == gen1893 {
								gen1890 := Call(__e, ShenFunc(symhd), V434)

								gen1891 := Call(__e, ShenFunc(symcons_2), gen1890)

								var gen1892 Obj
								if True == gen1891 {
									gen1886 := Call(__e, ShenFunc(symhd), V434)

									gen1887 := Call(__e, ShenFunc(symtl), gen1886)

									gen1888 := Call(__e, ShenFunc(symcons_2), gen1887)

									var gen1889 Obj
									if True == gen1888 {
										gen1882 := Call(__e, ShenFunc(symhd), V434)

										gen1883 := Call(__e, ShenFunc(symtl), gen1882)

										gen1884 := Call(__e, ShenFunc(symtl), gen1883)

										gen1885 := Call(__e, ShenFunc(sym_a), Nil, gen1884)

										if True == gen1885 {
											gen1889 = True
										} else {
											gen1889 = False
										}

									} else {
										gen1889 = False
									}
									if True == gen1889 {
										gen1892 = True
									} else {
										gen1892 = False
									}

								} else {
									gen1892 = False
								}
								if True == gen1892 {
									gen1894 = True
								} else {
									gen1894 = False
								}

							} else {
								gen1894 = False
							}
							if True == gen1894 {
								gen1879 := Call(__e, ShenFunc(symhd), V434)

								gen1880 := Call(__e, ShenFunc(symtl), V434)

								gen1881 := Call(__e, ShenFunc(symshen_4encode_1choices), gen1880, V435)

								__e.TailApply(ShenFunc(symcons), gen1879, gen1881)

								return

							} else {
								__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.encode-choices"))

								return
							}

						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.encode-choices"), gen2101)

		gen2256 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V442 := __args[0]
			_ = V442
			V443 := __args[1]
			_ = V443
			gen2255 := Call(__e, ShenFunc(sym_a), Nil, V442)

			if True == gen2255 {
				__e.TailApply(ShenFunc(symcons), V443, Nil)

				return
			} else {
				gen2253 := Call(__e, ShenFunc(symcons_2), V442)

				var gen2254 Obj
				if True == gen2253 {
					gen2250 := Call(__e, ShenFunc(symhd), V442)

					gen2251 := Call(__e, ShenFunc(symcons_2), gen2250)

					var gen2252 Obj
					if True == gen2251 {
						gen2246 := Call(__e, ShenFunc(symhd), V442)

						gen2247 := Call(__e, ShenFunc(symhd), gen2246)

						gen2248 := Call(__e, ShenFunc(symcons_2), gen2247)

						var gen2249 Obj
						if True == gen2248 {
							gen2241 := Call(__e, ShenFunc(symhd), V442)

							gen2242 := Call(__e, ShenFunc(symhd), gen2241)

							gen2243 := Call(__e, ShenFunc(symhd), gen2242)

							gen2244 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2243)

							var gen2245 Obj
							if True == gen2244 {
								gen2236 := Call(__e, ShenFunc(symhd), V442)

								gen2237 := Call(__e, ShenFunc(symhd), gen2236)

								gen2238 := Call(__e, ShenFunc(symtl), gen2237)

								gen2239 := Call(__e, ShenFunc(symcons_2), gen2238)

								var gen2240 Obj
								if True == gen2239 {
									gen2230 := Call(__e, ShenFunc(symhd), V442)

									gen2231 := Call(__e, ShenFunc(symhd), gen2230)

									gen2232 := Call(__e, ShenFunc(symtl), gen2231)

									gen2233 := Call(__e, ShenFunc(symhd), gen2232)

									gen2234 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2233)

									var gen2235 Obj
									if True == gen2234 {
										gen2224 := Call(__e, ShenFunc(symhd), V442)

										gen2225 := Call(__e, ShenFunc(symhd), gen2224)

										gen2226 := Call(__e, ShenFunc(symtl), gen2225)

										gen2227 := Call(__e, ShenFunc(symtl), gen2226)

										gen2228 := Call(__e, ShenFunc(sym_a), Nil, gen2227)

										var gen2229 Obj
										if True == gen2228 {
											gen2220 := Call(__e, ShenFunc(symhd), V442)

											gen2221 := Call(__e, ShenFunc(symtl), gen2220)

											gen2222 := Call(__e, ShenFunc(symcons_2), gen2221)

											var gen2223 Obj
											if True == gen2222 {
												gen2215 := Call(__e, ShenFunc(symhd), V442)

												gen2216 := Call(__e, ShenFunc(symtl), gen2215)

												gen2217 := Call(__e, ShenFunc(symhd), gen2216)

												gen2218 := Call(__e, ShenFunc(symcons_2), gen2217)

												var gen2219 Obj
												if True == gen2218 {
													gen2209 := Call(__e, ShenFunc(symhd), V442)

													gen2210 := Call(__e, ShenFunc(symtl), gen2209)

													gen2211 := Call(__e, ShenFunc(symhd), gen2210)

													gen2212 := Call(__e, ShenFunc(symhd), gen2211)

													gen2213 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), gen2212)

													var gen2214 Obj
													if True == gen2213 {
														gen2203 := Call(__e, ShenFunc(symhd), V442)

														gen2204 := Call(__e, ShenFunc(symtl), gen2203)

														gen2205 := Call(__e, ShenFunc(symhd), gen2204)

														gen2206 := Call(__e, ShenFunc(symtl), gen2205)

														gen2207 := Call(__e, ShenFunc(symcons_2), gen2206)

														var gen2208 Obj
														if True == gen2207 {
															gen2196 := Call(__e, ShenFunc(symhd), V442)

															gen2197 := Call(__e, ShenFunc(symtl), gen2196)

															gen2198 := Call(__e, ShenFunc(symhd), gen2197)

															gen2199 := Call(__e, ShenFunc(symtl), gen2198)

															gen2200 := Call(__e, ShenFunc(symtl), gen2199)

															gen2201 := Call(__e, ShenFunc(sym_a), Nil, gen2200)

															var gen2202 Obj
															if True == gen2201 {
																gen2192 := Call(__e, ShenFunc(symhd), V442)

																gen2193 := Call(__e, ShenFunc(symtl), gen2192)

																gen2194 := Call(__e, ShenFunc(symtl), gen2193)

																gen2195 := Call(__e, ShenFunc(sym_a), Nil, gen2194)

																if True == gen2195 {
																	gen2202 = True
																} else {
																	gen2202 = False
																}

															} else {
																gen2202 = False
															}
															if True == gen2202 {
																gen2208 = True
															} else {
																gen2208 = False
															}

														} else {
															gen2208 = False
														}
														if True == gen2208 {
															gen2214 = True
														} else {
															gen2214 = False
														}

													} else {
														gen2214 = False
													}
													if True == gen2214 {
														gen2219 = True
													} else {
														gen2219 = False
													}

												} else {
													gen2219 = False
												}
												if True == gen2219 {
													gen2223 = True
												} else {
													gen2223 = False
												}

											} else {
												gen2223 = False
											}
											if True == gen2223 {
												gen2229 = True
											} else {
												gen2229 = False
											}

										} else {
											gen2229 = False
										}
										if True == gen2229 {
											gen2235 = True
										} else {
											gen2235 = False
										}

									} else {
										gen2235 = False
									}
									if True == gen2235 {
										gen2240 = True
									} else {
										gen2240 = False
									}

								} else {
									gen2240 = False
								}
								if True == gen2240 {
									gen2245 = True
								} else {
									gen2245 = False
								}

							} else {
								gen2245 = False
							}
							if True == gen2245 {
								gen2249 = True
							} else {
								gen2249 = False
							}

						} else {
							gen2249 = False
						}
						if True == gen2249 {
							gen2252 = True
						} else {
							gen2252 = False
						}

					} else {
						gen2252 = False
					}
					if True == gen2252 {
						gen2254 = True
					} else {
						gen2254 = False
					}

				} else {
					gen2254 = False
				}
				if True == gen2254 {
					gen2187 := Call(__e, ShenFunc(symhd), V442)

					gen2188 := Call(__e, ShenFunc(symtl), gen2187)

					gen2189 := Call(__e, ShenFunc(symcons), True, gen2188)

					gen2190 := Call(__e, ShenFunc(symtl), V442)

					gen2191 := Call(__e, ShenFunc(symshen_4case_1form), gen2190, V443)

					__e.TailApply(ShenFunc(symcons), gen2189, gen2191)

					return

				} else {
					gen2185 := Call(__e, ShenFunc(symcons_2), V442)

					var gen2186 Obj
					if True == gen2185 {
						gen2182 := Call(__e, ShenFunc(symhd), V442)

						gen2183 := Call(__e, ShenFunc(symcons_2), gen2182)

						var gen2184 Obj
						if True == gen2183 {
							gen2178 := Call(__e, ShenFunc(symhd), V442)

							gen2179 := Call(__e, ShenFunc(symhd), gen2178)

							gen2180 := Call(__e, ShenFunc(symcons_2), gen2179)

							var gen2181 Obj
							if True == gen2180 {
								gen2173 := Call(__e, ShenFunc(symhd), V442)

								gen2174 := Call(__e, ShenFunc(symhd), gen2173)

								gen2175 := Call(__e, ShenFunc(symhd), gen2174)

								gen2176 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2175)

								var gen2177 Obj
								if True == gen2176 {
									gen2168 := Call(__e, ShenFunc(symhd), V442)

									gen2169 := Call(__e, ShenFunc(symhd), gen2168)

									gen2170 := Call(__e, ShenFunc(symtl), gen2169)

									gen2171 := Call(__e, ShenFunc(symcons_2), gen2170)

									var gen2172 Obj
									if True == gen2171 {
										gen2162 := Call(__e, ShenFunc(symhd), V442)

										gen2163 := Call(__e, ShenFunc(symhd), gen2162)

										gen2164 := Call(__e, ShenFunc(symtl), gen2163)

										gen2165 := Call(__e, ShenFunc(symhd), gen2164)

										gen2166 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2165)

										var gen2167 Obj
										if True == gen2166 {
											gen2156 := Call(__e, ShenFunc(symhd), V442)

											gen2157 := Call(__e, ShenFunc(symhd), gen2156)

											gen2158 := Call(__e, ShenFunc(symtl), gen2157)

											gen2159 := Call(__e, ShenFunc(symtl), gen2158)

											gen2160 := Call(__e, ShenFunc(sym_a), Nil, gen2159)

											var gen2161 Obj
											if True == gen2160 {
												gen2152 := Call(__e, ShenFunc(symhd), V442)

												gen2153 := Call(__e, ShenFunc(symtl), gen2152)

												gen2154 := Call(__e, ShenFunc(symcons_2), gen2153)

												var gen2155 Obj
												if True == gen2154 {
													gen2148 := Call(__e, ShenFunc(symhd), V442)

													gen2149 := Call(__e, ShenFunc(symtl), gen2148)

													gen2150 := Call(__e, ShenFunc(symtl), gen2149)

													gen2151 := Call(__e, ShenFunc(sym_a), Nil, gen2150)

													if True == gen2151 {
														gen2155 = True
													} else {
														gen2155 = False
													}

												} else {
													gen2155 = False
												}
												if True == gen2155 {
													gen2161 = True
												} else {
													gen2161 = False
												}

											} else {
												gen2161 = False
											}
											if True == gen2161 {
												gen2167 = True
											} else {
												gen2167 = False
											}

										} else {
											gen2167 = False
										}
										if True == gen2167 {
											gen2172 = True
										} else {
											gen2172 = False
										}

									} else {
										gen2172 = False
									}
									if True == gen2172 {
										gen2177 = True
									} else {
										gen2177 = False
									}

								} else {
									gen2177 = False
								}
								if True == gen2177 {
									gen2181 = True
								} else {
									gen2181 = False
								}

							} else {
								gen2181 = False
							}
							if True == gen2181 {
								gen2184 = True
							} else {
								gen2184 = False
							}

						} else {
							gen2184 = False
						}
						if True == gen2184 {
							gen2186 = True
						} else {
							gen2186 = False
						}

					} else {
						gen2186 = False
					}
					if True == gen2186 {
						gen2145 := Call(__e, ShenFunc(symhd), V442)

						gen2146 := Call(__e, ShenFunc(symtl), gen2145)

						gen2147 := Call(__e, ShenFunc(symcons), True, gen2146)

						__e.TailApply(ShenFunc(symcons), gen2147, Nil)

						return

					} else {
						gen2143 := Call(__e, ShenFunc(symcons_2), V442)

						var gen2144 Obj
						if True == gen2143 {
							gen2140 := Call(__e, ShenFunc(symhd), V442)

							gen2141 := Call(__e, ShenFunc(symcons_2), gen2140)

							var gen2142 Obj
							if True == gen2141 {
								gen2136 := Call(__e, ShenFunc(symhd), V442)

								gen2137 := Call(__e, ShenFunc(symhd), gen2136)

								gen2138 := Call(__e, ShenFunc(symcons_2), gen2137)

								var gen2139 Obj
								if True == gen2138 {
									gen2131 := Call(__e, ShenFunc(symhd), V442)

									gen2132 := Call(__e, ShenFunc(symhd), gen2131)

									gen2133 := Call(__e, ShenFunc(symhd), gen2132)

									gen2134 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen2133)

									var gen2135 Obj
									if True == gen2134 {
										gen2126 := Call(__e, ShenFunc(symhd), V442)

										gen2127 := Call(__e, ShenFunc(symhd), gen2126)

										gen2128 := Call(__e, ShenFunc(symtl), gen2127)

										gen2129 := Call(__e, ShenFunc(symcons_2), gen2128)

										var gen2130 Obj
										if True == gen2129 {
											gen2120 := Call(__e, ShenFunc(symhd), V442)

											gen2121 := Call(__e, ShenFunc(symhd), gen2120)

											gen2122 := Call(__e, ShenFunc(symtl), gen2121)

											gen2123 := Call(__e, ShenFunc(symhd), gen2122)

											gen2124 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.tests"), gen2123)

											var gen2125 Obj
											if True == gen2124 {
												gen2116 := Call(__e, ShenFunc(symhd), V442)

												gen2117 := Call(__e, ShenFunc(symtl), gen2116)

												gen2118 := Call(__e, ShenFunc(symcons_2), gen2117)

												var gen2119 Obj
												if True == gen2118 {
													gen2112 := Call(__e, ShenFunc(symhd), V442)

													gen2113 := Call(__e, ShenFunc(symtl), gen2112)

													gen2114 := Call(__e, ShenFunc(symtl), gen2113)

													gen2115 := Call(__e, ShenFunc(sym_a), Nil, gen2114)

													if True == gen2115 {
														gen2119 = True
													} else {
														gen2119 = False
													}

												} else {
													gen2119 = False
												}
												if True == gen2119 {
													gen2125 = True
												} else {
													gen2125 = False
												}

											} else {
												gen2125 = False
											}
											if True == gen2125 {
												gen2130 = True
											} else {
												gen2130 = False
											}

										} else {
											gen2130 = False
										}
										if True == gen2130 {
											gen2135 = True
										} else {
											gen2135 = False
										}

									} else {
										gen2135 = False
									}
									if True == gen2135 {
										gen2139 = True
									} else {
										gen2139 = False
									}

								} else {
									gen2139 = False
								}
								if True == gen2139 {
									gen2142 = True
								} else {
									gen2142 = False
								}

							} else {
								gen2142 = False
							}
							if True == gen2142 {
								gen2144 = True
							} else {
								gen2144 = False
							}

						} else {
							gen2144 = False
						}
						if True == gen2144 {
							gen2102 := Call(__e, ShenFunc(symhd), V442)

							gen2103 := Call(__e, ShenFunc(symhd), gen2102)

							gen2104 := Call(__e, ShenFunc(symtl), gen2103)

							gen2105 := Call(__e, ShenFunc(symtl), gen2104)

							gen2106 := Call(__e, ShenFunc(symshen_4embed_1and), gen2105)

							gen2107 := Call(__e, ShenFunc(symhd), V442)

							gen2108 := Call(__e, ShenFunc(symtl), gen2107)

							gen2109 := Call(__e, ShenFunc(symcons), gen2106, gen2108)

							gen2110 := Call(__e, ShenFunc(symtl), V442)

							gen2111 := Call(__e, ShenFunc(symshen_4case_1form), gen2110, V443)

							__e.TailApply(ShenFunc(symcons), gen2109, gen2111)

							return

						} else {
							__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.case-form"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.case-form"), gen2256)

		gen2267 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V445 := __args[0]
			_ = V445
			gen2265 := Call(__e, ShenFunc(symcons_2), V445)

			var gen2266 Obj
			if True == gen2265 {
				gen2263 := Call(__e, ShenFunc(symtl), V445)

				gen2264 := Call(__e, ShenFunc(sym_a), Nil, gen2263)

				if True == gen2264 {
					gen2266 = True
				} else {
					gen2266 = False
				}

			} else {
				gen2266 = False
			}
			if True == gen2266 {
				__e.TailApply(ShenFunc(symhd), V445)

				return
			} else {
				gen2262 := Call(__e, ShenFunc(symcons_2), V445)

				if True == gen2262 {
					gen2257 := Call(__e, ShenFunc(symhd), V445)

					gen2258 := Call(__e, ShenFunc(symtl), V445)

					gen2259 := Call(__e, ShenFunc(symshen_4embed_1and), gen2258)

					gen2260 := Call(__e, ShenFunc(symcons), gen2259, Nil)

					gen2261 := Call(__e, ShenFunc(symcons), gen2257, gen2260)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("and"), gen2261)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.embed-and"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.embed-and"), gen2267)

		gen2271 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V447 := __args[0]
			_ = V447
			gen2268 := Call(__e, ShenFunc(symcons), V447, Nil)

			gen2269 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.f_error"), gen2268)

			gen2270 := Call(__e, ShenFunc(symcons), gen2269, Nil)

			__e.TailApply(ShenFunc(symcons), True, gen2270)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.err-condition"), gen2271)

		gen2274 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V449 := __args[0]
			_ = V449
			gen2272 := Call(__e, ShenFunc(symshen_4app), V449, MakeString(": unexpected argument\n"), MakeSymbol("shen.a"))

			gen2273 := Call(__e, ShenFunc(symcn), MakeString("system function "), gen2272)

			__e.TailApply(ShenFunc(symsimple_1error), gen2273)

			return

		}, 1)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.sys-error"), gen2274)

		return

	}, 0))
}
