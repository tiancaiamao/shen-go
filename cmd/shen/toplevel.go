package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen1 := MakeNative(func(__e Evaluator) {
			Call(__e, __e.Global(symshen_4credits))
			__e.TailApply(__e.Global(symshen_4loop))

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.repl"), gen1)

		gen4 := MakeNative(func(__e Evaluator) {
			Call(__e, __e.Global(symshen_4initialise__environment))
			Call(__e, __e.Global(symshen_4prompt))
			gen2 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				__e.TailApply(__e.Global(symshen_4toplevel_1display_1exception), E)

				return
			}, 1)
			gen3 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symshen_4read_1evaluate_1print))

				return
			}, 0)
			Try(__e, gen3).Catch(gen2)

			__e.TailApply(__e.Global(symshen_4loop))

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.loop"), gen4)

		gen7 := MakeNative(func(__e Evaluator) {
			V3022 := __e.Get(1)
			_ = V3022
			gen5 := Call(__e, __e.Global(symerror_1to_1string), V3022)

			gen6 := Call(__e, __e.Global(symstoutput))

			__e.TailApply(__e.Global(sympr), gen5, gen6)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.toplevel-display-exception"), gen7)

		gen27 := MakeNative(func(__e Evaluator) {
			gen8 := Call(__e, __e.Global(symstoutput))

			Call(__e, __e.Global(symshen_4prhush), MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n"), gen8)

			gen9 := Call(__e, __e.Global(symvalue), MakeSymbol("*version*"))

			gen10 := Call(__e, __e.Global(symshen_4app), gen9, MakeString("\n"), MakeSymbol("shen.a"))

			gen11 := Call(__e, __e.Global(symcn), MakeString("www.shenlanguage.org, "), gen10)

			gen12 := Call(__e, __e.Global(symstoutput))

			Call(__e, __e.Global(symshen_4prhush), gen11, gen12)

			gen13 := Call(__e, __e.Global(symvalue), MakeSymbol("*language*"))

			gen14 := Call(__e, __e.Global(symvalue), MakeSymbol("*implementation*"))

			gen15 := Call(__e, __e.Global(symshen_4app), gen14, MakeString(""), MakeSymbol("shen.a"))

			gen16 := Call(__e, __e.Global(symcn), MakeString(", implementation: "), gen15)

			gen17 := Call(__e, __e.Global(symshen_4app), gen13, gen16, MakeSymbol("shen.a"))

			gen18 := Call(__e, __e.Global(symcn), MakeString("running under "), gen17)

			gen19 := Call(__e, __e.Global(symstoutput))

			Call(__e, __e.Global(symshen_4prhush), gen18, gen19)

			gen20 := Call(__e, __e.Global(symvalue), MakeSymbol("*port*"))

			gen21 := Call(__e, __e.Global(symvalue), MakeSymbol("*porters*"))

			gen22 := Call(__e, __e.Global(symshen_4app), gen21, MakeString("\n"), MakeSymbol("shen.a"))

			gen23 := Call(__e, __e.Global(symcn), MakeString(" ported by "), gen22)

			gen24 := Call(__e, __e.Global(symshen_4app), gen20, gen23, MakeSymbol("shen.a"))

			gen25 := Call(__e, __e.Global(symcn), MakeString("\nport "), gen24)

			gen26 := Call(__e, __e.Global(symstoutput))

			__e.TailApply(__e.Global(symshen_4prhush), gen25, gen26)

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.credits"), gen27)

		gen36 := MakeNative(func(__e Evaluator) {
			gen28 := Call(__e, __e.Global(symcons), MakeNumber(0), Nil)

			gen29 := Call(__e, __e.Global(symcons), MakeSymbol("shen.*catch*"), gen28)

			gen30 := Call(__e, __e.Global(symcons), MakeNumber(0), gen29)

			gen31 := Call(__e, __e.Global(symcons), MakeSymbol("shen.*process-counter*"), gen30)

			gen32 := Call(__e, __e.Global(symcons), MakeNumber(0), gen31)

			gen33 := Call(__e, __e.Global(symcons), MakeSymbol("shen.*infs*"), gen32)

			gen34 := Call(__e, __e.Global(symcons), MakeNumber(0), gen33)

			gen35 := Call(__e, __e.Global(symcons), MakeSymbol("shen.*call*"), gen34)

			__e.TailApply(__e.Global(symshen_4multiple_1set), gen35)

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.initialise_environment"), gen36)

		gen47 := MakeNative(func(__e Evaluator) {
			V3024 := __e.Get(1)
			_ = V3024
			gen46 := Call(__e, __e.Global(sym_a), Nil, V3024)

			if True == gen46 {
				__e.Return(Nil)
				return
			} else {
				gen44 := Call(__e, __e.Global(symcons_2), V3024)

				var gen45 Obj
				if True == gen44 {
					gen42 := Call(__e, __e.Global(symtl), V3024)

					gen43 := Call(__e, __e.Global(symcons_2), gen42)

					if True == gen43 {
						gen45 = True
					} else {
						gen45 = False
					}

				} else {
					gen45 = False
				}
				if True == gen45 {
					gen37 := Call(__e, __e.Global(symhd), V3024)

					gen38 := Call(__e, __e.Global(symtl), V3024)

					gen39 := Call(__e, __e.Global(symhd), gen38)

					Call(__e, __e.Global(symset), gen37, gen39)

					gen40 := Call(__e, __e.Global(symtl), V3024)

					gen41 := Call(__e, __e.Global(symtl), gen40)

					__e.TailApply(__e.Global(symshen_4multiple_1set), gen41)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.multiple-set"))

					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.multiple-set"), gen47)

		gen48 := MakeNative(func(__e Evaluator) {
			V3026 := __e.Get(1)
			_ = V3026
			__e.TailApply(__e.Global(symdeclare), V3026, MakeSymbol("symbol"))

			return
		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("destroy"), gen48)

		gen54 := MakeNative(func(__e Evaluator) {
			gen49 := Call(__e, __e.Global(symshen_4toplineread))

			Lineread := gen49
			gen50 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*history*"))

			History := gen50
			gen51 := Call(__e, __e.Global(symshen_4retrieve_1from_1history_1if_1needed), Lineread, History)

			NewLineread := gen51
			gen52 := Call(__e, __e.Global(symshen_4update__history), NewLineread, History)

			NewHistory := gen52
			_ = NewHistory
			gen53 := Call(__e, __e.Global(symfst), NewLineread)

			Parsed := gen53
			__e.TailApply(__e.Global(symshen_4toplevel), Parsed)

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.read-evaluate-print"), gen54)

		gen145 := MakeNative(func(__e Evaluator) {
			V3038 := __e.Get(1)
			_ = V3038
			V3039 := __e.Get(2)
			_ = V3039
			gen143 := Call(__e, __e.Global(symtuple_2), V3038)

			var gen144 Obj
			if True == gen143 {
				gen140 := Call(__e, __e.Global(symsnd), V3038)

				gen141 := Call(__e, __e.Global(symcons_2), gen140)

				var gen142 Obj
				if True == gen141 {
					gen133 := Call(__e, __e.Global(symsnd), V3038)

					gen134 := Call(__e, __e.Global(symhd), gen133)

					gen135 := Call(__e, __e.Global(symshen_4space))

					gen136 := Call(__e, __e.Global(symshen_4newline))

					gen137 := Call(__e, __e.Global(symcons), gen136, Nil)

					gen138 := Call(__e, __e.Global(symcons), gen135, gen137)

					gen139 := Call(__e, __e.Global(symelement_2), gen134, gen138)

					if True == gen139 {
						gen142 = True
					} else {
						gen142 = False
					}

				} else {
					gen142 = False
				}
				if True == gen142 {
					gen144 = True
				} else {
					gen144 = False
				}

			} else {
				gen144 = False
			}
			if True == gen144 {
				gen129 := Call(__e, __e.Global(symfst), V3038)

				gen130 := Call(__e, __e.Global(symsnd), V3038)

				gen131 := Call(__e, __e.Global(symtl), gen130)

				gen132 := Call(__e, __e.Global(sym_8p), gen129, gen131)

				__e.TailApply(__e.Global(symshen_4retrieve_1from_1history_1if_1needed), gen132, V3039)

				return

			} else {
				gen127 := Call(__e, __e.Global(symtuple_2), V3038)

				var gen128 Obj
				if True == gen127 {
					gen124 := Call(__e, __e.Global(symsnd), V3038)

					gen125 := Call(__e, __e.Global(symcons_2), gen124)

					var gen126 Obj
					if True == gen125 {
						gen120 := Call(__e, __e.Global(symsnd), V3038)

						gen121 := Call(__e, __e.Global(symtl), gen120)

						gen122 := Call(__e, __e.Global(symcons_2), gen121)

						var gen123 Obj
						if True == gen122 {
							gen115 := Call(__e, __e.Global(symsnd), V3038)

							gen116 := Call(__e, __e.Global(symtl), gen115)

							gen117 := Call(__e, __e.Global(symtl), gen116)

							gen118 := Call(__e, __e.Global(sym_a), Nil, gen117)

							var gen119 Obj
							if True == gen118 {
								gen113 := Call(__e, __e.Global(symcons_2), V3039)

								var gen114 Obj
								if True == gen113 {
									gen108 := Call(__e, __e.Global(symsnd), V3038)

									gen109 := Call(__e, __e.Global(symhd), gen108)

									gen110 := Call(__e, __e.Global(symshen_4exclamation))

									gen111 := Call(__e, __e.Global(sym_a), gen109, gen110)

									var gen112 Obj
									if True == gen111 {
										gen103 := Call(__e, __e.Global(symsnd), V3038)

										gen104 := Call(__e, __e.Global(symtl), gen103)

										gen105 := Call(__e, __e.Global(symhd), gen104)

										gen106 := Call(__e, __e.Global(symshen_4exclamation))

										gen107 := Call(__e, __e.Global(sym_a), gen105, gen106)

										if True == gen107 {
											gen112 = True
										} else {
											gen112 = False
										}

									} else {
										gen112 = False
									}
									if True == gen112 {
										gen114 = True
									} else {
										gen114 = False
									}

								} else {
									gen114 = False
								}
								if True == gen114 {
									gen119 = True
								} else {
									gen119 = False
								}

							} else {
								gen119 = False
							}
							if True == gen119 {
								gen123 = True
							} else {
								gen123 = False
							}

						} else {
							gen123 = False
						}
						if True == gen123 {
							gen126 = True
						} else {
							gen126 = False
						}

					} else {
						gen126 = False
					}
					if True == gen126 {
						gen128 = True
					} else {
						gen128 = False
					}

				} else {
					gen128 = False
				}
				if True == gen128 {
					gen100 := Call(__e, __e.Global(symhd), V3039)

					gen101 := Call(__e, __e.Global(symsnd), gen100)

					gen102 := Call(__e, __e.Global(symshen_4prbytes), gen101)

					PastPrint := gen102
					_ = PastPrint
					__e.TailApply(__e.Global(symhd), V3039)

					return

				} else {
					gen98 := Call(__e, __e.Global(symtuple_2), V3038)

					var gen99 Obj
					if True == gen98 {
						gen95 := Call(__e, __e.Global(symsnd), V3038)

						gen96 := Call(__e, __e.Global(symcons_2), gen95)

						var gen97 Obj
						if True == gen96 {
							gen91 := Call(__e, __e.Global(symsnd), V3038)

							gen92 := Call(__e, __e.Global(symhd), gen91)

							gen93 := Call(__e, __e.Global(symshen_4exclamation))

							gen94 := Call(__e, __e.Global(sym_a), gen92, gen93)

							if True == gen94 {
								gen97 = True
							} else {
								gen97 = False
							}

						} else {
							gen97 = False
						}
						if True == gen97 {
							gen99 = True
						} else {
							gen99 = False
						}

					} else {
						gen99 = False
					}
					if True == gen99 {
						gen84 := Call(__e, __e.Global(symsnd), V3038)

						gen85 := Call(__e, __e.Global(symtl), gen84)

						gen86 := Call(__e, __e.Global(symshen_4make_1key), gen85, V3039)

						Key_2 := gen86
						gen87 := Call(__e, __e.Global(symshen_4find_1past_1inputs), Key_2, V3039)

						gen88 := Call(__e, __e.Global(symhead), gen87)

						Find := gen88
						gen89 := Call(__e, __e.Global(symsnd), Find)

						gen90 := Call(__e, __e.Global(symshen_4prbytes), gen89)

						PastPrint := gen90
						_ = PastPrint
						__e.Return(Find)
						return

					} else {
						gen82 := Call(__e, __e.Global(symtuple_2), V3038)

						var gen83 Obj
						if True == gen82 {
							gen79 := Call(__e, __e.Global(symsnd), V3038)

							gen80 := Call(__e, __e.Global(symcons_2), gen79)

							var gen81 Obj
							if True == gen80 {
								gen75 := Call(__e, __e.Global(symsnd), V3038)

								gen76 := Call(__e, __e.Global(symtl), gen75)

								gen77 := Call(__e, __e.Global(sym_a), Nil, gen76)

								var gen78 Obj
								if True == gen77 {
									gen71 := Call(__e, __e.Global(symsnd), V3038)

									gen72 := Call(__e, __e.Global(symhd), gen71)

									gen73 := Call(__e, __e.Global(symshen_4percent))

									gen74 := Call(__e, __e.Global(sym_a), gen72, gen73)

									if True == gen74 {
										gen78 = True
									} else {
										gen78 = False
									}

								} else {
									gen78 = False
								}
								if True == gen78 {
									gen81 = True
								} else {
									gen81 = False
								}

							} else {
								gen81 = False
							}
							if True == gen81 {
								gen83 = True
							} else {
								gen83 = False
							}

						} else {
							gen83 = False
						}
						if True == gen83 {
							gen69 := MakeNative(func(__e Evaluator) {
								X := __e.Get(1)
								_ = X
								__e.Return(True)
								return
							}, 1)
							gen70 := Call(__e, __e.Global(symreverse), V3039)

							Call(__e, __e.Global(symshen_4print_1past_1inputs), gen69, gen70, MakeNumber(0))

							__e.TailApply(__e.Global(symabort))

							return

						} else {
							gen67 := Call(__e, __e.Global(symtuple_2), V3038)

							var gen68 Obj
							if True == gen67 {
								gen64 := Call(__e, __e.Global(symsnd), V3038)

								gen65 := Call(__e, __e.Global(symcons_2), gen64)

								var gen66 Obj
								if True == gen65 {
									gen60 := Call(__e, __e.Global(symsnd), V3038)

									gen61 := Call(__e, __e.Global(symhd), gen60)

									gen62 := Call(__e, __e.Global(symshen_4percent))

									gen63 := Call(__e, __e.Global(sym_a), gen61, gen62)

									if True == gen63 {
										gen66 = True
									} else {
										gen66 = False
									}

								} else {
									gen66 = False
								}
								if True == gen66 {
									gen68 = True
								} else {
									gen68 = False
								}

							} else {
								gen68 = False
							}
							if True == gen68 {
								gen55 := Call(__e, __e.Global(symsnd), V3038)

								gen56 := Call(__e, __e.Global(symtl), gen55)

								gen57 := Call(__e, __e.Global(symshen_4make_1key), gen56, V3039)

								Key_2 := gen57
								gen58 := Call(__e, __e.Global(symreverse), V3039)

								gen59 := Call(__e, __e.Global(symshen_4print_1past_1inputs), Key_2, gen58, MakeNumber(0))

								Pastprint := gen59
								_ = Pastprint
								__e.TailApply(__e.Global(symabort))

								return

							} else {
								__e.Return(V3038)
								return
							}

						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.retrieve-from-history-if-needed"), gen145)

		gen146 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(37))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.percent"), gen146)

		gen147 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(33))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.exclamation"), gen147)

		gen151 := MakeNative(func(__e Evaluator) {
			V3041 := __e.Get(1)
			_ = V3041
			gen150 := MakeNative(func(__e Evaluator) {
				Byte := __e.Get(1)
				_ = Byte
				gen148 := Call(__e, __e.Global(symn_1_6string), Byte)

				gen149 := Call(__e, __e.Global(symstoutput))

				__e.TailApply(__e.Global(sympr), gen148, gen149)

				return

			}, 1)
			Call(__e, __e.Global(symshen_4for_1each), gen150, V3041)

			__e.TailApply(__e.Global(symnl), MakeNumber(1))

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.prbytes"), gen151)

		gen153 := MakeNative(func(__e Evaluator) {
			V3044 := __e.Get(1)
			_ = V3044
			V3045 := __e.Get(2)
			_ = V3045
			gen152 := Call(__e, __e.Global(symcons), V3044, V3045)

			__e.TailApply(__e.Global(symset), MakeSymbol("shen.*history*"), gen152)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.update_history"), gen153)

		gen156 := MakeNative(func(__e Evaluator) {
			gen154 := Call(__e, __e.Global(symstinput))

			gen155 := Call(__e, __e.Global(symshen_4read_1char_1code), gen154)

			__e.TailApply(__e.Global(symshen_4toplineread__loop), gen155, Nil)

			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.toplineread"), gen156)

		gen180 := MakeNative(func(__e Evaluator) {
			V3049 := __e.Get(1)
			_ = V3049
			V3050 := __e.Get(2)
			_ = V3050
			gen178 := Call(__e, __e.Global(symshen_4hat))

			gen179 := Call(__e, __e.Global(sym_a), V3049, gen178)

			if True == gen179 {
				__e.TailApply(__e.Global(symsimple_1error), MakeString("line read aborted"))

				return
			} else {
				gen173 := Call(__e, __e.Global(symshen_4newline))

				gen174 := Call(__e, __e.Global(symshen_4carriage_1return))

				gen175 := Call(__e, __e.Global(symcons), gen174, Nil)

				gen176 := Call(__e, __e.Global(symcons), gen173, gen175)

				gen177 := Call(__e, __e.Global(symelement_2), V3049, gen176)

				if True == gen177 {
					gen162 := MakeNative(func(__e Evaluator) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(__e.Global(symshen_4_5st__input_6), X)

						return
					}, 1)
					gen163 := MakeNative(func(__e Evaluator) {
						E := __e.Get(1)
						_ = E
						__e.Return(MakeSymbol("shen.nextline"))
						return
					}, 1)
					gen164 := Call(__e, __e.Global(symcompile), gen162, V3050, gen163)

					Line := gen164
					gen165 := Call(__e, __e.Global(symshen_4record_1it), V3050)

					It := gen165
					_ = It
					gen171 := Call(__e, __e.Global(sym_a), Line, MakeSymbol("shen.nextline"))

					var gen172 Obj
					if True == gen171 {
						gen172 = True
					} else {
						gen170 := Call(__e, __e.Global(symempty_2), Line)

						if True == gen170 {
							gen172 = True
						} else {
							gen172 = False
						}

					}
					if True == gen172 {
						gen166 := Call(__e, __e.Global(symstinput))

						gen167 := Call(__e, __e.Global(symshen_4read_1char_1code), gen166)

						gen168 := Call(__e, __e.Global(symcons), V3049, Nil)

						gen169 := Call(__e, __e.Global(symappend), V3050, gen168)

						__e.TailApply(__e.Global(symshen_4toplineread__loop), gen167, gen169)

						return

					} else {
						__e.TailApply(__e.Global(sym_8p), Line, V3050)

						return
					}

				} else {
					gen157 := Call(__e, __e.Global(symstinput))

					gen158 := Call(__e, __e.Global(symshen_4read_1char_1code), gen157)

					gen160 := Call(__e, __e.Global(sym_a), V3049, MakeNumber(-1))

					var gen161 Obj
					if True == gen160 {
						gen161 = V3050
					} else {
						gen159 := Call(__e, __e.Global(symcons), V3049, Nil)

						gen161 = Call(__e, __e.Global(symappend), V3050, gen159)

					}
					__e.TailApply(__e.Global(symshen_4toplineread__loop), gen158, gen161)

					return

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.toplineread_loop"), gen180)

		gen181 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(94))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.hat"), gen181)

		gen182 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(10))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.newline"), gen182)

		gen183 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(13))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.carriage-return"), gen183)

		gen186 := MakeNative(func(__e Evaluator) {
			V3056 := __e.Get(1)
			_ = V3056
			gen185 := Call(__e, __e.Global(sym_a), MakeSymbol("+"), V3056)

			if True == gen185 {
				__e.TailApply(__e.Global(symset), MakeSymbol("shen.*tc*"), True)

				return
			} else {
				gen184 := Call(__e, __e.Global(sym_a), MakeSymbol("-"), V3056)

				if True == gen184 {
					__e.TailApply(__e.Global(symset), MakeSymbol("shen.*tc*"), False)

					return
				} else {
					__e.TailApply(__e.Global(symsimple_1error), MakeString("tc expects a + or -"))

					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("tc"), gen186)

		gen198 := MakeNative(func(__e Evaluator) {
			gen197 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*tc*"))

			if True == gen197 {
				gen192 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*history*"))

				gen193 := Call(__e, __e.Global(symlength), gen192)

				gen194 := Call(__e, __e.Global(symshen_4app), gen193, MakeString("+) "), MakeSymbol("shen.a"))

				gen195 := Call(__e, __e.Global(symcn), MakeString("\n\n("), gen194)

				gen196 := Call(__e, __e.Global(symstoutput))

				__e.TailApply(__e.Global(symshen_4prhush), gen195, gen196)

				return

			} else {
				gen187 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*history*"))

				gen188 := Call(__e, __e.Global(symlength), gen187)

				gen189 := Call(__e, __e.Global(symshen_4app), gen188, MakeString("-) "), MakeSymbol("shen.a"))

				gen190 := Call(__e, __e.Global(symcn), MakeString("\n\n("), gen189)

				gen191 := Call(__e, __e.Global(symstoutput))

				__e.TailApply(__e.Global(symshen_4prhush), gen190, gen191)

				return

			}

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.prompt"), gen198)

		gen200 := MakeNative(func(__e Evaluator) {
			V3058 := __e.Get(1)
			_ = V3058
			gen199 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*tc*"))

			__e.TailApply(__e.Global(symshen_4toplevel__evaluate), V3058, gen199)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.toplevel"), gen200)

		gen203 := MakeNative(func(__e Evaluator) {
			V3061 := __e.Get(1)
			_ = V3061
			V3062 := __e.Get(2)
			_ = V3062
			gen201 := Call(__e, __e.Global(symshen_4find), V3061, V3062)

			F := gen201
			gen202 := Call(__e, __e.Global(symempty_2), F)

			if True == gen202 {
				__e.TailApply(__e.Global(symsimple_1error), MakeString("input not found\n"))

				return
			} else {
				__e.Return(F)
				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.find-past-inputs"), gen203)

		gen217 := MakeNative(func(__e Evaluator) {
			V3065 := __e.Get(1)
			_ = V3065
			V3066 := __e.Get(2)
			_ = V3066
			gen204 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(__e.Global(symshen_4_5st__input_6), X)

				return
			}, 1)
			gen208 := MakeNative(func(__e Evaluator) {
				E := __e.Get(1)
				_ = E
				gen207 := Call(__e, __e.Global(symcons_2), E)

				if True == gen207 {
					gen205 := Call(__e, __e.Global(symshen_4app), E, MakeString("\n"), MakeSymbol("shen.s"))

					gen206 := Call(__e, __e.Global(symcn), MakeString("parse error here: "), gen205)

					__e.TailApply(__e.Global(symsimple_1error), gen206)

					return

				} else {
					__e.TailApply(__e.Global(symsimple_1error), MakeString("parse error\n"))

					return
				}

			}, 1)
			gen209 := Call(__e, __e.Global(symcompile), gen204, V3065, gen208)

			gen210 := Call(__e, __e.Global(symhd), gen209)

			Atom := gen210
			gen216 := Call(__e, __e.Global(syminteger_2), Atom)

			if True == gen216 {
				__e.Return(MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					gen213 := Call(__e, __e.Global(sym_7), Atom, MakeNumber(1))

					gen214 := Call(__e, __e.Global(symreverse), V3066)

					gen215 := Call(__e, __e.Global(symnth), gen213, gen214)

					__e.TailApply(__e.Global(sym_a), X, gen215)

					return

				}, 1))
				return
			} else {
				__e.Return(MakeNative(func(__e Evaluator) {
					X := __e.Get(1)
					_ = X
					gen211 := Call(__e, __e.Global(symsnd), X)

					gen212 := Call(__e, __e.Global(symshen_4trim_1gubbins), gen211)

					__e.TailApply(__e.Global(symshen_4prefix_2), V3065, gen212)

					return

				}, 1))
				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.make-key"), gen217)

		gen248 := MakeNative(func(__e Evaluator) {
			V3068 := __e.Get(1)
			_ = V3068
			gen246 := Call(__e, __e.Global(symcons_2), V3068)

			var gen247 Obj
			if True == gen246 {
				gen243 := Call(__e, __e.Global(symhd), V3068)

				gen244 := Call(__e, __e.Global(symshen_4space))

				gen245 := Call(__e, __e.Global(sym_a), gen243, gen244)

				if True == gen245 {
					gen247 = True
				} else {
					gen247 = False
				}

			} else {
				gen247 = False
			}
			if True == gen247 {
				gen242 := Call(__e, __e.Global(symtl), V3068)

				__e.TailApply(__e.Global(symshen_4trim_1gubbins), gen242)

				return

			} else {
				gen240 := Call(__e, __e.Global(symcons_2), V3068)

				var gen241 Obj
				if True == gen240 {
					gen237 := Call(__e, __e.Global(symhd), V3068)

					gen238 := Call(__e, __e.Global(symshen_4newline))

					gen239 := Call(__e, __e.Global(sym_a), gen237, gen238)

					if True == gen239 {
						gen241 = True
					} else {
						gen241 = False
					}

				} else {
					gen241 = False
				}
				if True == gen241 {
					gen236 := Call(__e, __e.Global(symtl), V3068)

					__e.TailApply(__e.Global(symshen_4trim_1gubbins), gen236)

					return

				} else {
					gen234 := Call(__e, __e.Global(symcons_2), V3068)

					var gen235 Obj
					if True == gen234 {
						gen231 := Call(__e, __e.Global(symhd), V3068)

						gen232 := Call(__e, __e.Global(symshen_4carriage_1return))

						gen233 := Call(__e, __e.Global(sym_a), gen231, gen232)

						if True == gen233 {
							gen235 = True
						} else {
							gen235 = False
						}

					} else {
						gen235 = False
					}
					if True == gen235 {
						gen230 := Call(__e, __e.Global(symtl), V3068)

						__e.TailApply(__e.Global(symshen_4trim_1gubbins), gen230)

						return

					} else {
						gen228 := Call(__e, __e.Global(symcons_2), V3068)

						var gen229 Obj
						if True == gen228 {
							gen225 := Call(__e, __e.Global(symhd), V3068)

							gen226 := Call(__e, __e.Global(symshen_4tab))

							gen227 := Call(__e, __e.Global(sym_a), gen225, gen226)

							if True == gen227 {
								gen229 = True
							} else {
								gen229 = False
							}

						} else {
							gen229 = False
						}
						if True == gen229 {
							gen224 := Call(__e, __e.Global(symtl), V3068)

							__e.TailApply(__e.Global(symshen_4trim_1gubbins), gen224)

							return

						} else {
							gen222 := Call(__e, __e.Global(symcons_2), V3068)

							var gen223 Obj
							if True == gen222 {
								gen219 := Call(__e, __e.Global(symhd), V3068)

								gen220 := Call(__e, __e.Global(symshen_4left_1round))

								gen221 := Call(__e, __e.Global(sym_a), gen219, gen220)

								if True == gen221 {
									gen223 = True
								} else {
									gen223 = False
								}

							} else {
								gen223 = False
							}
							if True == gen223 {
								gen218 := Call(__e, __e.Global(symtl), V3068)

								__e.TailApply(__e.Global(symshen_4trim_1gubbins), gen218)

								return

							} else {
								__e.Return(V3068)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.trim-gubbins"), gen248)

		gen249 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(32))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.space"), gen249)

		gen250 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(9))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.tab"), gen250)

		gen251 := MakeNative(func(__e Evaluator) {
			__e.Return(MakeNumber(40))
			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.left-round"), gen251)

		gen262 := MakeNative(func(__e Evaluator) {
			V3077 := __e.Get(1)
			_ = V3077
			V3078 := __e.Get(2)
			_ = V3078
			gen261 := Call(__e, __e.Global(sym_a), Nil, V3078)

			if True == gen261 {
				__e.Return(Nil)
				return
			} else {
				gen259 := Call(__e, __e.Global(symcons_2), V3078)

				var gen260 Obj
				if True == gen259 {
					gen257 := Call(__e, __e.Global(symhd), V3078)

					gen258 := Call(__e, V3077, gen257)

					if True == gen258 {
						gen260 = True
					} else {
						gen260 = False
					}

				} else {
					gen260 = False
				}
				if True == gen260 {
					gen254 := Call(__e, __e.Global(symhd), V3078)

					gen255 := Call(__e, __e.Global(symtl), V3078)

					gen256 := Call(__e, __e.Global(symshen_4find), V3077, gen255)

					__e.TailApply(__e.Global(symcons), gen254, gen256)

					return

				} else {
					gen253 := Call(__e, __e.Global(symcons_2), V3078)

					if True == gen253 {
						gen252 := Call(__e, __e.Global(symtl), V3078)

						__e.TailApply(__e.Global(symshen_4find), V3077, gen252)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.find"))

						return
					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.find"), gen262)

		gen273 := MakeNative(func(__e Evaluator) {
			V3092 := __e.Get(1)
			_ = V3092
			V3093 := __e.Get(2)
			_ = V3093
			gen272 := Call(__e, __e.Global(sym_a), Nil, V3092)

			if True == gen272 {
				__e.Return(True)
				return
			} else {
				gen270 := Call(__e, __e.Global(symcons_2), V3092)

				var gen271 Obj
				if True == gen270 {
					gen268 := Call(__e, __e.Global(symcons_2), V3093)

					var gen269 Obj
					if True == gen268 {
						gen265 := Call(__e, __e.Global(symhd), V3093)

						gen266 := Call(__e, __e.Global(symhd), V3092)

						gen267 := Call(__e, __e.Global(sym_a), gen265, gen266)

						if True == gen267 {
							gen269 = True
						} else {
							gen269 = False
						}

					} else {
						gen269 = False
					}
					if True == gen269 {
						gen271 = True
					} else {
						gen271 = False
					}

				} else {
					gen271 = False
				}
				if True == gen271 {
					gen263 := Call(__e, __e.Global(symtl), V3092)

					gen264 := Call(__e, __e.Global(symtl), V3093)

					__e.TailApply(__e.Global(symshen_4prefix_2), gen263, gen264)

					return

				} else {
					__e.Return(False)
					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.prefix?"), gen273)

		gen292 := MakeNative(func(__e Evaluator) {
			V3105 := __e.Get(1)
			_ = V3105
			V3106 := __e.Get(2)
			_ = V3106
			V3107 := __e.Get(3)
			_ = V3107
			gen291 := Call(__e, __e.Global(sym_a), Nil, V3106)

			if True == gen291 {
				__e.Return(MakeSymbol("_"))
				return
			} else {
				gen289 := Call(__e, __e.Global(symcons_2), V3106)

				var gen290 Obj
				if True == gen289 {
					gen286 := Call(__e, __e.Global(symhd), V3106)

					gen287 := Call(__e, V3105, gen286)

					gen288 := Call(__e, __e.Global(symnot), gen287)

					if True == gen288 {
						gen290 = True
					} else {
						gen290 = False
					}

				} else {
					gen290 = False
				}
				if True == gen290 {
					gen284 := Call(__e, __e.Global(symtl), V3106)

					gen285 := Call(__e, __e.Global(sym_7), V3107, MakeNumber(1))

					__e.TailApply(__e.Global(symshen_4print_1past_1inputs), V3105, gen284, gen285)

					return

				} else {
					gen282 := Call(__e, __e.Global(symcons_2), V3106)

					var gen283 Obj
					if True == gen282 {
						gen280 := Call(__e, __e.Global(symhd), V3106)

						gen281 := Call(__e, __e.Global(symtuple_2), gen280)

						if True == gen281 {
							gen283 = True
						} else {
							gen283 = False
						}

					} else {
						gen283 = False
					}
					if True == gen283 {
						gen274 := Call(__e, __e.Global(symshen_4app), V3107, MakeString(". "), MakeSymbol("shen.a"))

						gen275 := Call(__e, __e.Global(symstoutput))

						Call(__e, __e.Global(symshen_4prhush), gen274, gen275)

						gen276 := Call(__e, __e.Global(symhd), V3106)

						gen277 := Call(__e, __e.Global(symsnd), gen276)

						Call(__e, __e.Global(symshen_4prbytes), gen277)

						gen278 := Call(__e, __e.Global(symtl), V3106)

						gen279 := Call(__e, __e.Global(sym_7), V3107, MakeNumber(1))

						__e.TailApply(__e.Global(symshen_4print_1past_1inputs), V3105, gen278, gen279)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.print-past-inputs"))

						return
					}

				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.print-past-inputs"), gen292)

		gen339 := MakeNative(func(__e Evaluator) {
			V3110 := __e.Get(1)
			_ = V3110
			V3111 := __e.Get(2)
			_ = V3111
			gen337 := Call(__e, __e.Global(symcons_2), V3110)

			var gen338 Obj
			if True == gen337 {
				gen334 := Call(__e, __e.Global(symtl), V3110)

				gen335 := Call(__e, __e.Global(symcons_2), gen334)

				var gen336 Obj
				if True == gen335 {
					gen330 := Call(__e, __e.Global(symtl), V3110)

					gen331 := Call(__e, __e.Global(symhd), gen330)

					gen332 := Call(__e, __e.Global(sym_a), MakeSymbol(":"), gen331)

					var gen333 Obj
					if True == gen332 {
						gen326 := Call(__e, __e.Global(symtl), V3110)

						gen327 := Call(__e, __e.Global(symtl), gen326)

						gen328 := Call(__e, __e.Global(symcons_2), gen327)

						var gen329 Obj
						if True == gen328 {
							gen321 := Call(__e, __e.Global(symtl), V3110)

							gen322 := Call(__e, __e.Global(symtl), gen321)

							gen323 := Call(__e, __e.Global(symtl), gen322)

							gen324 := Call(__e, __e.Global(sym_a), Nil, gen323)

							var gen325 Obj
							if True == gen324 {
								gen320 := Call(__e, __e.Global(sym_a), True, V3111)

								if True == gen320 {
									gen325 = True
								} else {
									gen325 = False
								}

							} else {
								gen325 = False
							}
							if True == gen325 {
								gen329 = True
							} else {
								gen329 = False
							}

						} else {
							gen329 = False
						}
						if True == gen329 {
							gen333 = True
						} else {
							gen333 = False
						}

					} else {
						gen333 = False
					}
					if True == gen333 {
						gen336 = True
					} else {
						gen336 = False
					}

				} else {
					gen336 = False
				}
				if True == gen336 {
					gen338 = True
				} else {
					gen338 = False
				}

			} else {
				gen338 = False
			}
			if True == gen338 {
				gen316 := Call(__e, __e.Global(symhd), V3110)

				gen317 := Call(__e, __e.Global(symtl), V3110)

				gen318 := Call(__e, __e.Global(symtl), gen317)

				gen319 := Call(__e, __e.Global(symhd), gen318)

				__e.TailApply(__e.Global(symshen_4typecheck_1and_1evaluate), gen316, gen319)

				return

			} else {
				gen314 := Call(__e, __e.Global(symcons_2), V3110)

				var gen315 Obj
				if True == gen314 {
					gen312 := Call(__e, __e.Global(symtl), V3110)

					gen313 := Call(__e, __e.Global(symcons_2), gen312)

					if True == gen313 {
						gen315 = True
					} else {
						gen315 = False
					}

				} else {
					gen315 = False
				}
				if True == gen315 {
					gen309 := Call(__e, __e.Global(symhd), V3110)

					gen310 := Call(__e, __e.Global(symcons), gen309, Nil)

					Call(__e, __e.Global(symshen_4toplevel__evaluate), gen310, V3111)

					Call(__e, __e.Global(symnl), MakeNumber(1))
					gen311 := Call(__e, __e.Global(symtl), V3110)

					__e.TailApply(__e.Global(symshen_4toplevel__evaluate), gen311, V3111)

					return

				} else {
					gen307 := Call(__e, __e.Global(symcons_2), V3110)

					var gen308 Obj
					if True == gen307 {
						gen304 := Call(__e, __e.Global(symtl), V3110)

						gen305 := Call(__e, __e.Global(sym_a), Nil, gen304)

						var gen306 Obj
						if True == gen305 {
							gen303 := Call(__e, __e.Global(sym_a), True, V3111)

							if True == gen303 {
								gen306 = True
							} else {
								gen306 = False
							}

						} else {
							gen306 = False
						}
						if True == gen306 {
							gen308 = True
						} else {
							gen308 = False
						}

					} else {
						gen308 = False
					}
					if True == gen308 {
						gen301 := Call(__e, __e.Global(symhd), V3110)

						gen302 := Call(__e, __e.Global(symgensym), MakeSymbol("A"))

						__e.TailApply(__e.Global(symshen_4typecheck_1and_1evaluate), gen301, gen302)

						return

					} else {
						gen299 := Call(__e, __e.Global(symcons_2), V3110)

						var gen300 Obj
						if True == gen299 {
							gen296 := Call(__e, __e.Global(symtl), V3110)

							gen297 := Call(__e, __e.Global(sym_a), Nil, gen296)

							var gen298 Obj
							if True == gen297 {
								gen295 := Call(__e, __e.Global(sym_a), False, V3111)

								if True == gen295 {
									gen298 = True
								} else {
									gen298 = False
								}

							} else {
								gen298 = False
							}
							if True == gen298 {
								gen300 = True
							} else {
								gen300 = False
							}

						} else {
							gen300 = False
						}
						if True == gen300 {
							gen293 := Call(__e, __e.Global(symhd), V3110)

							gen294 := Call(__e, __e.Global(symshen_4eval_1without_1macros), gen293)

							Eval := gen294
							__e.TailApply(__e.Global(symprint), Eval)

							return

						} else {
							__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.toplevel_evaluate"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.toplevel_evaluate"), gen339)

		gen348 := MakeNative(func(__e Evaluator) {
			V3114 := __e.Get(1)
			_ = V3114
			V3115 := __e.Get(2)
			_ = V3115
			gen340 := Call(__e, __e.Global(symshen_4typecheck), V3114, V3115)

			Typecheck := gen340
			gen347 := Call(__e, __e.Global(sym_a), Typecheck, False)

			if True == gen347 {
				__e.TailApply(__e.Global(symsimple_1error), MakeString("type error\n"))

				return
			} else {
				gen341 := Call(__e, __e.Global(symshen_4eval_1without_1macros), V3114)

				Eval := gen341
				gen342 := Call(__e, __e.Global(symshen_4pretty_1type), Typecheck)

				Type := gen342
				gen343 := Call(__e, __e.Global(symshen_4app), Type, MakeString(""), MakeSymbol("shen.r"))

				gen344 := Call(__e, __e.Global(symcn), MakeString(" : "), gen343)

				gen345 := Call(__e, __e.Global(symshen_4app), Eval, gen344, MakeSymbol("shen.s"))

				gen346 := Call(__e, __e.Global(symstoutput))

				__e.TailApply(__e.Global(symshen_4prhush), gen345, gen346)

				return

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.typecheck-and-evaluate"), gen348)

		gen351 := MakeNative(func(__e Evaluator) {
			V3117 := __e.Get(1)
			_ = V3117
			gen349 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.*alphabet*"))

			gen350 := Call(__e, __e.Global(symshen_4extract_1pvars), V3117)

			__e.TailApply(__e.Global(symshen_4mult__subst), gen349, gen350, V3117)

			return

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.pretty-type"), gen351)

		gen358 := MakeNative(func(__e Evaluator) {
			V3123 := __e.Get(1)
			_ = V3123
			gen357 := Call(__e, __e.Global(symshen_4pvar_2), V3123)

			if True == gen357 {
				__e.TailApply(__e.Global(symcons), V3123, Nil)

				return
			} else {
				gen356 := Call(__e, __e.Global(symcons_2), V3123)

				if True == gen356 {
					gen352 := Call(__e, __e.Global(symhd), V3123)

					gen353 := Call(__e, __e.Global(symshen_4extract_1pvars), gen352)

					gen354 := Call(__e, __e.Global(symtl), V3123)

					gen355 := Call(__e, __e.Global(symshen_4extract_1pvars), gen354)

					__e.TailApply(__e.Global(symunion), gen353, gen355)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.extract-pvars"), gen358)

		gen369 := MakeNative(func(__e Evaluator) {
			V3131 := __e.Get(1)
			_ = V3131
			V3132 := __e.Get(2)
			_ = V3132
			V3133 := __e.Get(3)
			_ = V3133
			gen368 := Call(__e, __e.Global(sym_a), Nil, V3131)

			if True == gen368 {
				__e.Return(V3133)
				return
			} else {
				gen367 := Call(__e, __e.Global(sym_a), Nil, V3132)

				if True == gen367 {
					__e.Return(V3133)
					return
				} else {
					gen365 := Call(__e, __e.Global(symcons_2), V3131)

					var gen366 Obj
					if True == gen365 {
						gen364 := Call(__e, __e.Global(symcons_2), V3132)

						if True == gen364 {
							gen366 = True
						} else {
							gen366 = False
						}

					} else {
						gen366 = False
					}
					if True == gen366 {
						gen359 := Call(__e, __e.Global(symtl), V3131)

						gen360 := Call(__e, __e.Global(symtl), V3132)

						gen361 := Call(__e, __e.Global(symhd), V3131)

						gen362 := Call(__e, __e.Global(symhd), V3132)

						gen363 := Call(__e, __e.Global(symsubst), gen361, gen362, V3133)

						__e.TailApply(__e.Global(symshen_4mult__subst), gen359, gen360, gen363)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.mult_subst"))

						return
					}

				}

			}

		}, 3)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.mult_subst"), gen369)

		return

	}, 0))
}
