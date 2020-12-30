package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen11033 := MakeNative(func(__e Evaluator) {
			V2731 := __e.Get(1)
			_ = V2731
			V2732 := __e.Get(2)
			_ = V2732
			gen11024 := Call(__e, ShenFunc(symshen_4curry), V2731)

			Curry := gen11024
			gen11025 := Call(__e, ShenFunc(symshen_4start_1new_1prolog_1process))

			ProcessN := gen11025
			gen11026 := Call(__e, ShenFunc(symshen_4curry_1type), V2732)

			gen11027 := Call(__e, ShenFunc(symshen_4demodulate), gen11026)

			gen11028 := Call(__e, ShenFunc(symshen_4insert_1prolog_1variables), gen11027, ProcessN)

			Type := gen11028
			gen11029 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symreturn), Type, ProcessN, MakeSymbol("shen.void"))

				return
			}, 0)
			Continuation := gen11029
			gen11030 := Call(__e, ShenFunc(symcons), Type, Nil)

			gen11031 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11030)

			gen11032 := Call(__e, ShenFunc(symcons), Curry, gen11031)

			__e.TailApply(ShenFunc(symshen_4t_d), gen11032, Nil, ProcessN, Continuation)

			return

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typecheck"), gen11033)

		gen11101 := MakeNative(func(__e Evaluator) {
			V2734 := __e.Get(1)
			_ = V2734
			gen11099 := Call(__e, ShenFunc(symcons_2), V2734)

			var gen11100 Obj
			if True == gen11099 {
				gen11097 := Call(__e, ShenFunc(symhd), V2734)

				gen11098 := Call(__e, ShenFunc(symshen_4special_2), gen11097)

				if True == gen11098 {
					gen11100 = True
				} else {
					gen11100 = False
				}

			} else {
				gen11100 = False
			}
			if True == gen11100 {
				gen11093 := Call(__e, ShenFunc(symhd), V2734)

				gen11094 := MakeNative(func(__e Evaluator) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(ShenFunc(symshen_4curry), Y)

					return
				}, 1)
				gen11095 := Call(__e, ShenFunc(symtl), V2734)

				gen11096 := Call(__e, ShenFunc(symmap), gen11094, gen11095)

				__e.TailApply(ShenFunc(symcons), gen11093, gen11096)

				return

			} else {
				gen11091 := Call(__e, ShenFunc(symcons_2), V2734)

				var gen11092 Obj
				if True == gen11091 {
					gen11088 := Call(__e, ShenFunc(symtl), V2734)

					gen11089 := Call(__e, ShenFunc(symcons_2), gen11088)

					var gen11090 Obj
					if True == gen11089 {
						gen11086 := Call(__e, ShenFunc(symhd), V2734)

						gen11087 := Call(__e, ShenFunc(symshen_4extraspecial_2), gen11086)

						if True == gen11087 {
							gen11090 = True
						} else {
							gen11090 = False
						}

					} else {
						gen11090 = False
					}
					if True == gen11090 {
						gen11092 = True
					} else {
						gen11092 = False
					}

				} else {
					gen11092 = False
				}
				if True == gen11092 {
					__e.Return(V2734)
					return
				} else {
					gen11084 := Call(__e, ShenFunc(symcons_2), V2734)

					var gen11085 Obj
					if True == gen11084 {
						gen11081 := Call(__e, ShenFunc(symhd), V2734)

						gen11082 := Call(__e, ShenFunc(sym_a), MakeSymbol("type"), gen11081)

						var gen11083 Obj
						if True == gen11082 {
							gen11078 := Call(__e, ShenFunc(symtl), V2734)

							gen11079 := Call(__e, ShenFunc(symcons_2), gen11078)

							var gen11080 Obj
							if True == gen11079 {
								gen11074 := Call(__e, ShenFunc(symtl), V2734)

								gen11075 := Call(__e, ShenFunc(symtl), gen11074)

								gen11076 := Call(__e, ShenFunc(symcons_2), gen11075)

								var gen11077 Obj
								if True == gen11076 {
									gen11070 := Call(__e, ShenFunc(symtl), V2734)

									gen11071 := Call(__e, ShenFunc(symtl), gen11070)

									gen11072 := Call(__e, ShenFunc(symtl), gen11071)

									gen11073 := Call(__e, ShenFunc(sym_a), Nil, gen11072)

									if True == gen11073 {
										gen11077 = True
									} else {
										gen11077 = False
									}

								} else {
									gen11077 = False
								}
								if True == gen11077 {
									gen11080 = True
								} else {
									gen11080 = False
								}

							} else {
								gen11080 = False
							}
							if True == gen11080 {
								gen11083 = True
							} else {
								gen11083 = False
							}

						} else {
							gen11083 = False
						}
						if True == gen11083 {
							gen11085 = True
						} else {
							gen11085 = False
						}

					} else {
						gen11085 = False
					}
					if True == gen11085 {
						gen11064 := Call(__e, ShenFunc(symtl), V2734)

						gen11065 := Call(__e, ShenFunc(symhd), gen11064)

						gen11066 := Call(__e, ShenFunc(symshen_4curry), gen11065)

						gen11067 := Call(__e, ShenFunc(symtl), V2734)

						gen11068 := Call(__e, ShenFunc(symtl), gen11067)

						gen11069 := Call(__e, ShenFunc(symcons), gen11066, gen11068)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("type"), gen11069)

						return

					} else {
						gen11062 := Call(__e, ShenFunc(symcons_2), V2734)

						var gen11063 Obj
						if True == gen11062 {
							gen11059 := Call(__e, ShenFunc(symtl), V2734)

							gen11060 := Call(__e, ShenFunc(symcons_2), gen11059)

							var gen11061 Obj
							if True == gen11060 {
								gen11056 := Call(__e, ShenFunc(symtl), V2734)

								gen11057 := Call(__e, ShenFunc(symtl), gen11056)

								gen11058 := Call(__e, ShenFunc(symcons_2), gen11057)

								if True == gen11058 {
									gen11061 = True
								} else {
									gen11061 = False
								}

							} else {
								gen11061 = False
							}
							if True == gen11061 {
								gen11063 = True
							} else {
								gen11063 = False
							}

						} else {
							gen11063 = False
						}
						if True == gen11063 {
							gen11048 := Call(__e, ShenFunc(symhd), V2734)

							gen11049 := Call(__e, ShenFunc(symtl), V2734)

							gen11050 := Call(__e, ShenFunc(symhd), gen11049)

							gen11051 := Call(__e, ShenFunc(symcons), gen11050, Nil)

							gen11052 := Call(__e, ShenFunc(symcons), gen11048, gen11051)

							gen11053 := Call(__e, ShenFunc(symtl), V2734)

							gen11054 := Call(__e, ShenFunc(symtl), gen11053)

							gen11055 := Call(__e, ShenFunc(symcons), gen11052, gen11054)

							__e.TailApply(ShenFunc(symshen_4curry), gen11055)

							return

						} else {
							gen11046 := Call(__e, ShenFunc(symcons_2), V2734)

							var gen11047 Obj
							if True == gen11046 {
								gen11043 := Call(__e, ShenFunc(symtl), V2734)

								gen11044 := Call(__e, ShenFunc(symcons_2), gen11043)

								var gen11045 Obj
								if True == gen11044 {
									gen11040 := Call(__e, ShenFunc(symtl), V2734)

									gen11041 := Call(__e, ShenFunc(symtl), gen11040)

									gen11042 := Call(__e, ShenFunc(sym_a), Nil, gen11041)

									if True == gen11042 {
										gen11045 = True
									} else {
										gen11045 = False
									}

								} else {
									gen11045 = False
								}
								if True == gen11045 {
									gen11047 = True
								} else {
									gen11047 = False
								}

							} else {
								gen11047 = False
							}
							if True == gen11047 {
								gen11034 := Call(__e, ShenFunc(symhd), V2734)

								gen11035 := Call(__e, ShenFunc(symshen_4curry), gen11034)

								gen11036 := Call(__e, ShenFunc(symtl), V2734)

								gen11037 := Call(__e, ShenFunc(symhd), gen11036)

								gen11038 := Call(__e, ShenFunc(symshen_4curry), gen11037)

								gen11039 := Call(__e, ShenFunc(symcons), gen11038, Nil)

								__e.TailApply(ShenFunc(symcons), gen11035, gen11039)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry"), gen11101)

		gen11103 := MakeNative(func(__e Evaluator) {
			V2736 := __e.Get(1)
			_ = V2736
			gen11102 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*special*"))

			__e.TailApply(ShenFunc(symelement_2), V2736, gen11102)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.special?"), gen11103)

		gen11105 := MakeNative(func(__e Evaluator) {
			V2738 := __e.Get(1)
			_ = V2738
			gen11104 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*extraspecial*"))

			__e.TailApply(ShenFunc(symelement_2), V2738, gen11104)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.extraspecial?"), gen11105)

		gen11144 := MakeNative(func(__e Evaluator) {
			V2743 := __e.Get(1)
			_ = V2743
			V2744 := __e.Get(2)
			_ = V2744
			V2745 := __e.Get(3)
			_ = V2745
			V2746 := __e.Get(4)
			_ = V2746
			gen11106 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen11106
			gen11107 := Call(__e, ShenFunc(symshen_4newpv), V2745)

			Error := gen11107
			Call(__e, ShenFunc(symshen_4incinfs))
			gen11108 := Call(__e, ShenFunc(symshen_4maxinfexceeded_2))

			gen11110 := MakeNative(func(__e Evaluator) {
				gen11109 := Call(__e, ShenFunc(symshen_4errormaxinfs))

				__e.TailApply(ShenFunc(symbind), Error, gen11109, V2745, V2746)

				return

			}, 0)
			gen11111 := Call(__e, ShenFunc(symfwhen), gen11108, V2745, gen11110)

			Case := gen11111
			gen11142 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen11143 Obj
			if True == gen11142 {
				gen11112 := Call(__e, ShenFunc(symshen_4lazyderef), V2743, V2745)

				V2723 := gen11112
				gen11114 := Call(__e, ShenFunc(sym_a), MakeSymbol("fail"), V2723)

				var gen11115 Obj
				if True == gen11114 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen11113 := MakeNative(func(__e Evaluator) {
						__e.TailApply(ShenFunc(symshen_4prolog_1failure), V2745, V2746)

						return
					}, 0)
					gen11115 = Call(__e, ShenFunc(symcut), Throwcontrol, V2745, gen11113)

				} else {
					gen11115 = False
				}
				Case := gen11115
				gen11141 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen11141 {
					gen11116 := Call(__e, ShenFunc(symshen_4lazyderef), V2743, V2745)

					V2724 := gen11116
					gen11134 := Call(__e, ShenFunc(symcons_2), V2724)

					var gen11135 Obj
					if True == gen11134 {
						gen11117 := Call(__e, ShenFunc(symhd), V2724)

						X := gen11117
						gen11118 := Call(__e, ShenFunc(symtl), V2724)

						gen11119 := Call(__e, ShenFunc(symshen_4lazyderef), gen11118, V2745)

						V2725 := gen11119
						gen11133 := Call(__e, ShenFunc(symcons_2), V2725)

						if True == gen11133 {
							gen11120 := Call(__e, ShenFunc(symhd), V2725)

							gen11121 := Call(__e, ShenFunc(symshen_4lazyderef), gen11120, V2745)

							V2726 := gen11121
							gen11132 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2726)

							if True == gen11132 {
								gen11122 := Call(__e, ShenFunc(symtl), V2725)

								gen11123 := Call(__e, ShenFunc(symshen_4lazyderef), gen11122, V2745)

								V2727 := gen11123
								gen11131 := Call(__e, ShenFunc(symcons_2), V2727)

								if True == gen11131 {
									gen11124 := Call(__e, ShenFunc(symhd), V2727)

									A := gen11124
									gen11125 := Call(__e, ShenFunc(symtl), V2727)

									gen11126 := Call(__e, ShenFunc(symshen_4lazyderef), gen11125, V2745)

									V2728 := gen11126
									gen11130 := Call(__e, ShenFunc(sym_a), Nil, V2728)

									if True == gen11130 {
										Call(__e, ShenFunc(symshen_4incinfs))
										gen11127 := Call(__e, ShenFunc(symshen_4type_1theory_1enabled_2))

										gen11129 := MakeNative(func(__e Evaluator) {
											gen11128 := MakeNative(func(__e Evaluator) {
												__e.TailApply(ShenFunc(symshen_4th_d), X, A, V2744, V2745, V2746)

												return
											}, 0)
											__e.TailApply(ShenFunc(symcut), Throwcontrol, V2745, gen11128)

											return

										}, 0)
										gen11135 = Call(__e, ShenFunc(symfwhen), gen11127, V2745, gen11129)

									} else {
										gen11135 = False
									}

								} else {
									gen11135 = False
								}

							} else {
								gen11135 = False
							}

						} else {
							gen11135 = False
						}

					} else {
						gen11135 = False
					}
					Case := gen11135
					gen11140 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen11140 {
						gen11136 := Call(__e, ShenFunc(symshen_4newpv), V2745)

						Datatypes := gen11136
						Call(__e, ShenFunc(symshen_4incinfs))
						gen11139 := MakeNative(func(__e Evaluator) {
							gen11137 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

							gen11138 := MakeNative(func(__e Evaluator) {
								__e.TailApply(ShenFunc(symshen_4udefs_d), V2743, V2744, Datatypes, V2745, V2746)

								return
							}, 0)
							__e.TailApply(ShenFunc(symbind), Datatypes, gen11137, V2745, gen11138)

							return

						}, 0)
						gen11143 = Call(__e, ShenFunc(symshen_4show), V2743, V2744, V2745, gen11139)

					} else {
						gen11143 = Case
					}

				} else {
					gen11143 = Case
				}

			} else {
				gen11143 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen11143)

			return

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*"), gen11144)

		gen11145 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symvalue), MakeSymbol("shen.*shen-type-theory-enabled?*"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.type-theory-enabled?"), gen11145)

		gen11148 := MakeNative(func(__e Evaluator) {
			V2752 := __e.Get(1)
			_ = V2752
			gen11147 := Call(__e, ShenFunc(sym_a), MakeSymbol("+"), V2752)

			if True == gen11147 {
				__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), True)

				return
			} else {
				gen11146 := Call(__e, ShenFunc(sym_a), MakeSymbol("-"), V2752)

				if True == gen11146 {
					__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), False)

					return
				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("enable-type-theory expects a + or a -\n"))

					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("enable-type-theory"), gen11148)

		gen11149 := MakeNative(func(__e Evaluator) {
			V2763 := __e.Get(1)
			_ = V2763
			V2764 := __e.Get(2)
			_ = V2764
			__e.Return(False)
			return
		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-failure"), gen11149)

		gen11152 := MakeNative(func(__e Evaluator) {
			gen11150 := Call(__e, ShenFunc(syminferences))

			gen11151 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*maxinferences*"))

			__e.TailApply(ShenFunc(sym_6), gen11150, gen11151)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.maxinfexceeded?"), gen11152)

		gen11153 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symsimple_1error), MakeString("maximum inferences exceeded~%"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.errormaxinfs"), gen11153)

		gen11165 := MakeNative(func(__e Evaluator) {
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
			gen11154 := Call(__e, ShenFunc(symshen_4lazyderef), V2772, V2773)

			V2719 := gen11154
			gen11159 := Call(__e, ShenFunc(symcons_2), V2719)

			var gen11160 Obj
			if True == gen11159 {
				gen11155 := Call(__e, ShenFunc(symhd), V2719)

				D := gen11155
				Call(__e, ShenFunc(symshen_4incinfs))
				gen11156 := Call(__e, ShenFunc(symcons), V2771, Nil)

				gen11157 := Call(__e, ShenFunc(symcons), V2770, gen11156)

				gen11158 := Call(__e, ShenFunc(symcons), D, gen11157)

				gen11160 = Call(__e, ShenFunc(symcall), gen11158, V2773, V2774)

			} else {
				gen11160 = False
			}
			Case := gen11160
			gen11164 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen11164 {
				gen11161 := Call(__e, ShenFunc(symshen_4lazyderef), V2772, V2773)

				V2720 := gen11161
				gen11163 := Call(__e, ShenFunc(symcons_2), V2720)

				if True == gen11163 {
					gen11162 := Call(__e, ShenFunc(symtl), V2720)

					Ds := gen11162
					Call(__e, ShenFunc(symshen_4incinfs))
					__e.TailApply(ShenFunc(symshen_4udefs_d), V2770, V2771, Ds, V2773, V2774)

					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.udefs*"), gen11165)

		gen11909 := MakeNative(func(__e Evaluator) {
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
			gen11166 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen11166
			Call(__e, ShenFunc(symshen_4incinfs))
			gen11167 := Call(__e, ShenFunc(symcons), V2781, Nil)

			gen11168 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11167)

			gen11169 := Call(__e, ShenFunc(symcons), V2780, gen11168)

			gen11170 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symfwhen), False, V2783, V2784)

				return
			}, 0)
			gen11171 := Call(__e, ShenFunc(symshen_4show), gen11169, V2782, V2783, gen11170)

			Case := gen11171
			gen11907 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen11908 Obj
			if True == gen11907 {
				gen11172 := Call(__e, ShenFunc(symshen_4newpv), V2783)

				F := gen11172
				Call(__e, ShenFunc(symshen_4incinfs))
				gen11173 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

				gen11174 := Call(__e, ShenFunc(symshen_4typedf_2), gen11173)

				gen11180 := MakeNative(func(__e Evaluator) {
					gen11175 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

					gen11176 := Call(__e, ShenFunc(symshen_4sigf), gen11175)

					gen11179 := MakeNative(func(__e Evaluator) {
						gen11177 := Call(__e, ShenFunc(symcons), V2781, Nil)

						gen11178 := Call(__e, ShenFunc(symcons), F, gen11177)

						__e.TailApply(ShenFunc(symcall), gen11178, V2783, V2784)

						return

					}, 0)
					__e.TailApply(ShenFunc(symbind), F, gen11176, V2783, gen11179)

					return

				}, 0)
				gen11181 := Call(__e, ShenFunc(symfwhen), gen11174, V2783, gen11180)

				Case := gen11181
				gen11906 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen11906 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen11182 := Call(__e, ShenFunc(symshen_4base), V2780, V2781, V2783, V2784)

					Case := gen11182
					gen11905 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen11905 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen11183 := Call(__e, ShenFunc(symshen_4by__hypothesis), V2780, V2781, V2782, V2783, V2784)

						Case := gen11183
						gen11904 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen11904 {
							gen11184 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

							V2615 := gen11184
							gen11191 := Call(__e, ShenFunc(symcons_2), V2615)

							var gen11192 Obj
							if True == gen11191 {
								gen11185 := Call(__e, ShenFunc(symhd), V2615)

								F := gen11185
								gen11186 := Call(__e, ShenFunc(symtl), V2615)

								gen11187 := Call(__e, ShenFunc(symshen_4lazyderef), gen11186, V2783)

								V2616 := gen11187
								gen11190 := Call(__e, ShenFunc(sym_a), Nil, V2616)

								if True == gen11190 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen11188 := Call(__e, ShenFunc(symcons), V2781, Nil)

									gen11189 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen11188)

									gen11192 = Call(__e, ShenFunc(symshen_4th_d), F, gen11189, V2782, V2783, V2784)

								} else {
									gen11192 = False
								}

							} else {
								gen11192 = False
							}
							Case := gen11192
							gen11903 := Call(__e, ShenFunc(sym_a), Case, False)

							if True == gen11903 {
								gen11193 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

								V2617 := gen11193
								gen11207 := Call(__e, ShenFunc(symcons_2), V2617)

								var gen11208 Obj
								if True == gen11207 {
									gen11194 := Call(__e, ShenFunc(symhd), V2617)

									F := gen11194
									gen11195 := Call(__e, ShenFunc(symtl), V2617)

									gen11196 := Call(__e, ShenFunc(symshen_4lazyderef), gen11195, V2783)

									V2618 := gen11196
									gen11206 := Call(__e, ShenFunc(symcons_2), V2618)

									if True == gen11206 {
										gen11197 := Call(__e, ShenFunc(symhd), V2618)

										X := gen11197
										gen11198 := Call(__e, ShenFunc(symtl), V2618)

										gen11199 := Call(__e, ShenFunc(symshen_4lazyderef), gen11198, V2783)

										V2619 := gen11199
										gen11205 := Call(__e, ShenFunc(sym_a), Nil, V2619)

										if True == gen11205 {
											gen11200 := Call(__e, ShenFunc(symshen_4newpv), V2783)

											B := gen11200
											Call(__e, ShenFunc(symshen_4incinfs))
											gen11201 := Call(__e, ShenFunc(symcons), V2781, Nil)

											gen11202 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen11201)

											gen11203 := Call(__e, ShenFunc(symcons), B, gen11202)

											gen11204 := MakeNative(func(__e Evaluator) {
												__e.TailApply(ShenFunc(symshen_4th_d), X, B, V2782, V2783, V2784)

												return
											}, 0)
											gen11208 = Call(__e, ShenFunc(symshen_4th_d), F, gen11203, V2782, V2783, gen11204)

										} else {
											gen11208 = False
										}

									} else {
										gen11208 = False
									}

								} else {
									gen11208 = False
								}
								Case := gen11208
								gen11902 := Call(__e, ShenFunc(sym_a), Case, False)

								if True == gen11902 {
									gen11209 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

									V2620 := gen11209
									gen11283 := Call(__e, ShenFunc(symcons_2), V2620)

									var gen11284 Obj
									if True == gen11283 {
										gen11210 := Call(__e, ShenFunc(symhd), V2620)

										gen11211 := Call(__e, ShenFunc(symshen_4lazyderef), gen11210, V2783)

										V2621 := gen11211
										gen11282 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), V2621)

										if True == gen11282 {
											gen11212 := Call(__e, ShenFunc(symtl), V2620)

											gen11213 := Call(__e, ShenFunc(symshen_4lazyderef), gen11212, V2783)

											V2622 := gen11213
											gen11281 := Call(__e, ShenFunc(symcons_2), V2622)

											if True == gen11281 {
												gen11214 := Call(__e, ShenFunc(symhd), V2622)

												X := gen11214
												gen11215 := Call(__e, ShenFunc(symtl), V2622)

												gen11216 := Call(__e, ShenFunc(symshen_4lazyderef), gen11215, V2783)

												V2623 := gen11216
												gen11280 := Call(__e, ShenFunc(symcons_2), V2623)

												if True == gen11280 {
													gen11217 := Call(__e, ShenFunc(symhd), V2623)

													Y := gen11217
													gen11218 := Call(__e, ShenFunc(symtl), V2623)

													gen11219 := Call(__e, ShenFunc(symshen_4lazyderef), gen11218, V2783)

													V2624 := gen11219
													gen11279 := Call(__e, ShenFunc(sym_a), Nil, V2624)

													if True == gen11279 {
														gen11220 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

														V2625 := gen11220
														gen11278 := Call(__e, ShenFunc(symcons_2), V2625)

														if True == gen11278 {
															gen11229 := Call(__e, ShenFunc(symhd), V2625)

															gen11230 := Call(__e, ShenFunc(symshen_4lazyderef), gen11229, V2783)

															V2626 := gen11230
															gen11277 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2626)

															if True == gen11277 {
																gen11255 := Call(__e, ShenFunc(symtl), V2625)

																gen11256 := Call(__e, ShenFunc(symshen_4lazyderef), gen11255, V2783)

																V2627 := gen11256
																gen11276 := Call(__e, ShenFunc(symcons_2), V2627)

																if True == gen11276 {
																	gen11264 := Call(__e, ShenFunc(symhd), V2627)

																	A := gen11264
																	gen11265 := Call(__e, ShenFunc(symtl), V2627)

																	gen11266 := Call(__e, ShenFunc(symshen_4lazyderef), gen11265, V2783)

																	V2628 := gen11266
																	gen11275 := Call(__e, ShenFunc(sym_a), Nil, V2628)

																	if True == gen11275 {
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen11274 := MakeNative(func(__e Evaluator) {
																			gen11272 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen11273 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11272)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11273, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen11284 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11274)

																	} else {
																		gen11271 := Call(__e, ShenFunc(symshen_4pvar_2), V2628)

																		if True == gen11271 {
																			Call(__e, ShenFunc(symshen_4bindv), V2628, Nil, V2783)
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen11269 := MakeNative(func(__e Evaluator) {
																				gen11267 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen11268 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11267)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11268, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen11270 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11269)

																			Result := gen11270
																			Call(__e, ShenFunc(symshen_4unbindv), V2628, V2783)
																			gen11284 = Result

																		} else {
																			gen11284 = False
																		}

																	}

																} else {
																	gen11263 := Call(__e, ShenFunc(symshen_4pvar_2), V2627)

																	if True == gen11263 {
																		gen11257 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		A := gen11257
																		gen11258 := Call(__e, ShenFunc(symcons), A, Nil)

																		Call(__e, ShenFunc(symshen_4bindv), V2627, gen11258, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen11261 := MakeNative(func(__e Evaluator) {
																			gen11259 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen11260 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11259)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11260, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen11262 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11261)

																		Result := gen11262
																		Call(__e, ShenFunc(symshen_4unbindv), V2627, V2783)
																		gen11284 = Result

																	} else {
																		gen11284 = False
																	}

																}

															} else {
																gen11254 := Call(__e, ShenFunc(symshen_4pvar_2), V2626)

																if True == gen11254 {
																	Call(__e, ShenFunc(symshen_4bindv), V2626, MakeSymbol("list"), V2783)
																	gen11231 := Call(__e, ShenFunc(symtl), V2625)

																	gen11232 := Call(__e, ShenFunc(symshen_4lazyderef), gen11231, V2783)

																	V2629 := gen11232
																	gen11252 := Call(__e, ShenFunc(symcons_2), V2629)

																	var gen11253 Obj
																	if True == gen11252 {
																		gen11240 := Call(__e, ShenFunc(symhd), V2629)

																		A := gen11240
																		gen11241 := Call(__e, ShenFunc(symtl), V2629)

																		gen11242 := Call(__e, ShenFunc(symshen_4lazyderef), gen11241, V2783)

																		V2630 := gen11242
																		gen11251 := Call(__e, ShenFunc(sym_a), Nil, V2630)

																		if True == gen11251 {
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen11250 := MakeNative(func(__e Evaluator) {
																				gen11248 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen11249 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11248)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11249, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen11253 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11250)

																		} else {
																			gen11247 := Call(__e, ShenFunc(symshen_4pvar_2), V2630)

																			if True == gen11247 {
																				Call(__e, ShenFunc(symshen_4bindv), V2630, Nil, V2783)
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11245 := MakeNative(func(__e Evaluator) {
																					gen11243 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen11244 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11243)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11244, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen11246 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11245)

																				Result := gen11246
																				Call(__e, ShenFunc(symshen_4unbindv), V2630, V2783)
																				gen11253 = Result

																			} else {
																				gen11253 = False
																			}

																		}

																	} else {
																		gen11239 := Call(__e, ShenFunc(symshen_4pvar_2), V2629)

																		if True == gen11239 {
																			gen11233 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																			A := gen11233
																			gen11234 := Call(__e, ShenFunc(symcons), A, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2629, gen11234, V2783)

																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen11237 := MakeNative(func(__e Evaluator) {
																				gen11235 := Call(__e, ShenFunc(symcons), A, Nil)

																				gen11236 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11235)

																				__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11236, V2782, V2783, V2784)

																				return

																			}, 0)
																			gen11238 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11237)

																			Result := gen11238
																			Call(__e, ShenFunc(symshen_4unbindv), V2629, V2783)
																			gen11253 = Result

																		} else {
																			gen11253 = False
																		}

																	}
																	Result := gen11253
																	Call(__e, ShenFunc(symshen_4unbindv), V2626, V2783)
																	gen11284 = Result

																} else {
																	gen11284 = False
																}

															}

														} else {
															gen11228 := Call(__e, ShenFunc(symshen_4pvar_2), V2625)

															if True == gen11228 {
																gen11221 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																A := gen11221
																gen11222 := Call(__e, ShenFunc(symcons), A, Nil)

																gen11223 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11222)

																Call(__e, ShenFunc(symshen_4bindv), V2625, gen11223, V2783)

																Call(__e, ShenFunc(symshen_4incinfs))
																gen11226 := MakeNative(func(__e Evaluator) {
																	gen11224 := Call(__e, ShenFunc(symcons), A, Nil)

																	gen11225 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11224)

																	__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11225, V2782, V2783, V2784)

																	return

																}, 0)
																gen11227 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11226)

																Result := gen11227
																Call(__e, ShenFunc(symshen_4unbindv), V2625, V2783)
																gen11284 = Result

															} else {
																gen11284 = False
															}

														}

													} else {
														gen11284 = False
													}

												} else {
													gen11284 = False
												}

											} else {
												gen11284 = False
											}

										} else {
											gen11284 = False
										}

									} else {
										gen11284 = False
									}
									Case := gen11284
									gen11901 := Call(__e, ShenFunc(sym_a), Case, False)

									if True == gen11901 {
										gen11285 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

										V2631 := gen11285
										gen11357 := Call(__e, ShenFunc(symcons_2), V2631)

										var gen11358 Obj
										if True == gen11357 {
											gen11286 := Call(__e, ShenFunc(symhd), V2631)

											gen11287 := Call(__e, ShenFunc(symshen_4lazyderef), gen11286, V2783)

											V2632 := gen11287
											gen11356 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), V2632)

											if True == gen11356 {
												gen11288 := Call(__e, ShenFunc(symtl), V2631)

												gen11289 := Call(__e, ShenFunc(symshen_4lazyderef), gen11288, V2783)

												V2633 := gen11289
												gen11355 := Call(__e, ShenFunc(symcons_2), V2633)

												if True == gen11355 {
													gen11290 := Call(__e, ShenFunc(symhd), V2633)

													X := gen11290
													gen11291 := Call(__e, ShenFunc(symtl), V2633)

													gen11292 := Call(__e, ShenFunc(symshen_4lazyderef), gen11291, V2783)

													V2634 := gen11292
													gen11354 := Call(__e, ShenFunc(symcons_2), V2634)

													if True == gen11354 {
														gen11293 := Call(__e, ShenFunc(symhd), V2634)

														Y := gen11293
														gen11294 := Call(__e, ShenFunc(symtl), V2634)

														gen11295 := Call(__e, ShenFunc(symshen_4lazyderef), gen11294, V2783)

														V2635 := gen11295
														gen11353 := Call(__e, ShenFunc(sym_a), Nil, V2635)

														if True == gen11353 {
															gen11296 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

															V2636 := gen11296
															gen11352 := Call(__e, ShenFunc(symcons_2), V2636)

															if True == gen11352 {
																gen11305 := Call(__e, ShenFunc(symhd), V2636)

																A := gen11305
																gen11306 := Call(__e, ShenFunc(symtl), V2636)

																gen11307 := Call(__e, ShenFunc(symshen_4lazyderef), gen11306, V2783)

																V2637 := gen11307
																gen11351 := Call(__e, ShenFunc(symcons_2), V2637)

																if True == gen11351 {
																	gen11314 := Call(__e, ShenFunc(symhd), V2637)

																	gen11315 := Call(__e, ShenFunc(symshen_4lazyderef), gen11314, V2783)

																	V2638 := gen11315
																	gen11350 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), V2638)

																	if True == gen11350 {
																		gen11334 := Call(__e, ShenFunc(symtl), V2637)

																		gen11335 := Call(__e, ShenFunc(symshen_4lazyderef), gen11334, V2783)

																		V2639 := gen11335
																		gen11349 := Call(__e, ShenFunc(symcons_2), V2639)

																		if True == gen11349 {
																			gen11341 := Call(__e, ShenFunc(symhd), V2639)

																			B := gen11341
																			gen11342 := Call(__e, ShenFunc(symtl), V2639)

																			gen11343 := Call(__e, ShenFunc(symshen_4lazyderef), gen11342, V2783)

																			V2640 := gen11343
																			gen11348 := Call(__e, ShenFunc(sym_a), Nil, V2640)

																			if True == gen11348 {
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11347 := MakeNative(func(__e Evaluator) {
																					__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																					return
																				}, 0)
																				gen11358 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11347)

																			} else {
																				gen11346 := Call(__e, ShenFunc(symshen_4pvar_2), V2640)

																				if True == gen11346 {
																					Call(__e, ShenFunc(symshen_4bindv), V2640, Nil, V2783)
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11344 := MakeNative(func(__e Evaluator) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen11345 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11344)

																					Result := gen11345
																					Call(__e, ShenFunc(symshen_4unbindv), V2640, V2783)
																					gen11358 = Result

																				} else {
																					gen11358 = False
																				}

																			}

																		} else {
																			gen11340 := Call(__e, ShenFunc(symshen_4pvar_2), V2639)

																			if True == gen11340 {
																				gen11336 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen11336
																				gen11337 := Call(__e, ShenFunc(symcons), B, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2639, gen11337, V2783)

																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11338 := MakeNative(func(__e Evaluator) {
																					__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																					return
																				}, 0)
																				gen11339 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11338)

																				Result := gen11339
																				Call(__e, ShenFunc(symshen_4unbindv), V2639, V2783)
																				gen11358 = Result

																			} else {
																				gen11358 = False
																			}

																		}

																	} else {
																		gen11333 := Call(__e, ShenFunc(symshen_4pvar_2), V2638)

																		if True == gen11333 {
																			Call(__e, ShenFunc(symshen_4bindv), V2638, MakeSymbol("*"), V2783)
																			gen11316 := Call(__e, ShenFunc(symtl), V2637)

																			gen11317 := Call(__e, ShenFunc(symshen_4lazyderef), gen11316, V2783)

																			V2641 := gen11317
																			gen11331 := Call(__e, ShenFunc(symcons_2), V2641)

																			var gen11332 Obj
																			if True == gen11331 {
																				gen11323 := Call(__e, ShenFunc(symhd), V2641)

																				B := gen11323
																				gen11324 := Call(__e, ShenFunc(symtl), V2641)

																				gen11325 := Call(__e, ShenFunc(symshen_4lazyderef), gen11324, V2783)

																				V2642 := gen11325
																				gen11330 := Call(__e, ShenFunc(sym_a), Nil, V2642)

																				if True == gen11330 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11329 := MakeNative(func(__e Evaluator) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen11332 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11329)

																				} else {
																					gen11328 := Call(__e, ShenFunc(symshen_4pvar_2), V2642)

																					if True == gen11328 {
																						Call(__e, ShenFunc(symshen_4bindv), V2642, Nil, V2783)
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen11326 := MakeNative(func(__e Evaluator) {
																							__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																							return
																						}, 0)
																						gen11327 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11326)

																						Result := gen11327
																						Call(__e, ShenFunc(symshen_4unbindv), V2642, V2783)
																						gen11332 = Result

																					} else {
																						gen11332 = False
																					}

																				}

																			} else {
																				gen11322 := Call(__e, ShenFunc(symshen_4pvar_2), V2641)

																				if True == gen11322 {
																					gen11318 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					B := gen11318
																					gen11319 := Call(__e, ShenFunc(symcons), B, Nil)

																					Call(__e, ShenFunc(symshen_4bindv), V2641, gen11319, V2783)

																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11320 := MakeNative(func(__e Evaluator) {
																						__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																						return
																					}, 0)
																					gen11321 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11320)

																					Result := gen11321
																					Call(__e, ShenFunc(symshen_4unbindv), V2641, V2783)
																					gen11332 = Result

																				} else {
																					gen11332 = False
																				}

																			}
																			Result := gen11332
																			Call(__e, ShenFunc(symshen_4unbindv), V2638, V2783)
																			gen11358 = Result

																		} else {
																			gen11358 = False
																		}

																	}

																} else {
																	gen11313 := Call(__e, ShenFunc(symshen_4pvar_2), V2637)

																	if True == gen11313 {
																		gen11308 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		B := gen11308
																		gen11309 := Call(__e, ShenFunc(symcons), B, Nil)

																		gen11310 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen11309)

																		Call(__e, ShenFunc(symshen_4bindv), V2637, gen11310, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen11311 := MakeNative(func(__e Evaluator) {
																			__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																			return
																		}, 0)
																		gen11312 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11311)

																		Result := gen11312
																		Call(__e, ShenFunc(symshen_4unbindv), V2637, V2783)
																		gen11358 = Result

																	} else {
																		gen11358 = False
																	}

																}

															} else {
																gen11304 := Call(__e, ShenFunc(symshen_4pvar_2), V2636)

																if True == gen11304 {
																	gen11297 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																	A := gen11297
																	gen11298 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																	B := gen11298
																	gen11299 := Call(__e, ShenFunc(symcons), B, Nil)

																	gen11300 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen11299)

																	gen11301 := Call(__e, ShenFunc(symcons), A, gen11300)

																	Call(__e, ShenFunc(symshen_4bindv), V2636, gen11301, V2783)

																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen11302 := MakeNative(func(__e Evaluator) {
																		__e.TailApply(ShenFunc(symshen_4th_d), Y, B, V2782, V2783, V2784)

																		return
																	}, 0)
																	gen11303 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11302)

																	Result := gen11303
																	Call(__e, ShenFunc(symshen_4unbindv), V2636, V2783)
																	gen11358 = Result

																} else {
																	gen11358 = False
																}

															}

														} else {
															gen11358 = False
														}

													} else {
														gen11358 = False
													}

												} else {
													gen11358 = False
												}

											} else {
												gen11358 = False
											}

										} else {
											gen11358 = False
										}
										Case := gen11358
										gen11900 := Call(__e, ShenFunc(sym_a), Case, False)

										if True == gen11900 {
											gen11359 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

											V2643 := gen11359
											gen11433 := Call(__e, ShenFunc(symcons_2), V2643)

											var gen11434 Obj
											if True == gen11433 {
												gen11360 := Call(__e, ShenFunc(symhd), V2643)

												gen11361 := Call(__e, ShenFunc(symshen_4lazyderef), gen11360, V2783)

												V2644 := gen11361
												gen11432 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), V2644)

												if True == gen11432 {
													gen11362 := Call(__e, ShenFunc(symtl), V2643)

													gen11363 := Call(__e, ShenFunc(symshen_4lazyderef), gen11362, V2783)

													V2645 := gen11363
													gen11431 := Call(__e, ShenFunc(symcons_2), V2645)

													if True == gen11431 {
														gen11364 := Call(__e, ShenFunc(symhd), V2645)

														X := gen11364
														gen11365 := Call(__e, ShenFunc(symtl), V2645)

														gen11366 := Call(__e, ShenFunc(symshen_4lazyderef), gen11365, V2783)

														V2646 := gen11366
														gen11430 := Call(__e, ShenFunc(symcons_2), V2646)

														if True == gen11430 {
															gen11367 := Call(__e, ShenFunc(symhd), V2646)

															Y := gen11367
															gen11368 := Call(__e, ShenFunc(symtl), V2646)

															gen11369 := Call(__e, ShenFunc(symshen_4lazyderef), gen11368, V2783)

															V2647 := gen11369
															gen11429 := Call(__e, ShenFunc(sym_a), Nil, V2647)

															if True == gen11429 {
																gen11370 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																V2648 := gen11370
																gen11428 := Call(__e, ShenFunc(symcons_2), V2648)

																if True == gen11428 {
																	gen11379 := Call(__e, ShenFunc(symhd), V2648)

																	gen11380 := Call(__e, ShenFunc(symshen_4lazyderef), gen11379, V2783)

																	V2649 := gen11380
																	gen11427 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), V2649)

																	if True == gen11427 {
																		gen11405 := Call(__e, ShenFunc(symtl), V2648)

																		gen11406 := Call(__e, ShenFunc(symshen_4lazyderef), gen11405, V2783)

																		V2650 := gen11406
																		gen11426 := Call(__e, ShenFunc(symcons_2), V2650)

																		if True == gen11426 {
																			gen11414 := Call(__e, ShenFunc(symhd), V2650)

																			A := gen11414
																			gen11415 := Call(__e, ShenFunc(symtl), V2650)

																			gen11416 := Call(__e, ShenFunc(symshen_4lazyderef), gen11415, V2783)

																			V2651 := gen11416
																			gen11425 := Call(__e, ShenFunc(sym_a), Nil, V2651)

																			if True == gen11425 {
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11424 := MakeNative(func(__e Evaluator) {
																					gen11422 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen11423 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11422)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11423, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen11434 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11424)

																			} else {
																				gen11421 := Call(__e, ShenFunc(symshen_4pvar_2), V2651)

																				if True == gen11421 {
																					Call(__e, ShenFunc(symshen_4bindv), V2651, Nil, V2783)
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11419 := MakeNative(func(__e Evaluator) {
																						gen11417 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen11418 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11417)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11418, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen11420 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11419)

																					Result := gen11420
																					Call(__e, ShenFunc(symshen_4unbindv), V2651, V2783)
																					gen11434 = Result

																				} else {
																					gen11434 = False
																				}

																			}

																		} else {
																			gen11413 := Call(__e, ShenFunc(symshen_4pvar_2), V2650)

																			if True == gen11413 {
																				gen11407 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				A := gen11407
																				gen11408 := Call(__e, ShenFunc(symcons), A, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2650, gen11408, V2783)

																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11411 := MakeNative(func(__e Evaluator) {
																					gen11409 := Call(__e, ShenFunc(symcons), A, Nil)

																					gen11410 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11409)

																					__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11410, V2782, V2783, V2784)

																					return

																				}, 0)
																				gen11412 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11411)

																				Result := gen11412
																				Call(__e, ShenFunc(symshen_4unbindv), V2650, V2783)
																				gen11434 = Result

																			} else {
																				gen11434 = False
																			}

																		}

																	} else {
																		gen11404 := Call(__e, ShenFunc(symshen_4pvar_2), V2649)

																		if True == gen11404 {
																			Call(__e, ShenFunc(symshen_4bindv), V2649, MakeSymbol("vector"), V2783)
																			gen11381 := Call(__e, ShenFunc(symtl), V2648)

																			gen11382 := Call(__e, ShenFunc(symshen_4lazyderef), gen11381, V2783)

																			V2652 := gen11382
																			gen11402 := Call(__e, ShenFunc(symcons_2), V2652)

																			var gen11403 Obj
																			if True == gen11402 {
																				gen11390 := Call(__e, ShenFunc(symhd), V2652)

																				A := gen11390
																				gen11391 := Call(__e, ShenFunc(symtl), V2652)

																				gen11392 := Call(__e, ShenFunc(symshen_4lazyderef), gen11391, V2783)

																				V2653 := gen11392
																				gen11401 := Call(__e, ShenFunc(sym_a), Nil, V2653)

																				if True == gen11401 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11400 := MakeNative(func(__e Evaluator) {
																						gen11398 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen11399 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11398)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11399, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen11403 = Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11400)

																				} else {
																					gen11397 := Call(__e, ShenFunc(symshen_4pvar_2), V2653)

																					if True == gen11397 {
																						Call(__e, ShenFunc(symshen_4bindv), V2653, Nil, V2783)
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen11395 := MakeNative(func(__e Evaluator) {
																							gen11393 := Call(__e, ShenFunc(symcons), A, Nil)

																							gen11394 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11393)

																							__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11394, V2782, V2783, V2784)

																							return

																						}, 0)
																						gen11396 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11395)

																						Result := gen11396
																						Call(__e, ShenFunc(symshen_4unbindv), V2653, V2783)
																						gen11403 = Result

																					} else {
																						gen11403 = False
																					}

																				}

																			} else {
																				gen11389 := Call(__e, ShenFunc(symshen_4pvar_2), V2652)

																				if True == gen11389 {
																					gen11383 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					A := gen11383
																					gen11384 := Call(__e, ShenFunc(symcons), A, Nil)

																					Call(__e, ShenFunc(symshen_4bindv), V2652, gen11384, V2783)

																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11387 := MakeNative(func(__e Evaluator) {
																						gen11385 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen11386 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11385)

																						__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11386, V2782, V2783, V2784)

																						return

																					}, 0)
																					gen11388 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11387)

																					Result := gen11388
																					Call(__e, ShenFunc(symshen_4unbindv), V2652, V2783)
																					gen11403 = Result

																				} else {
																					gen11403 = False
																				}

																			}
																			Result := gen11403
																			Call(__e, ShenFunc(symshen_4unbindv), V2649, V2783)
																			gen11434 = Result

																		} else {
																			gen11434 = False
																		}

																	}

																} else {
																	gen11378 := Call(__e, ShenFunc(symshen_4pvar_2), V2648)

																	if True == gen11378 {
																		gen11371 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																		A := gen11371
																		gen11372 := Call(__e, ShenFunc(symcons), A, Nil)

																		gen11373 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11372)

																		Call(__e, ShenFunc(symshen_4bindv), V2648, gen11373, V2783)

																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen11376 := MakeNative(func(__e Evaluator) {
																			gen11374 := Call(__e, ShenFunc(symcons), A, Nil)

																			gen11375 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen11374)

																			__e.TailApply(ShenFunc(symshen_4th_d), Y, gen11375, V2782, V2783, V2784)

																			return

																		}, 0)
																		gen11377 := Call(__e, ShenFunc(symshen_4th_d), X, A, V2782, V2783, gen11376)

																		Result := gen11377
																		Call(__e, ShenFunc(symshen_4unbindv), V2648, V2783)
																		gen11434 = Result

																	} else {
																		gen11434 = False
																	}

																}

															} else {
																gen11434 = False
															}

														} else {
															gen11434 = False
														}

													} else {
														gen11434 = False
													}

												} else {
													gen11434 = False
												}

											} else {
												gen11434 = False
											}
											Case := gen11434
											gen11899 := Call(__e, ShenFunc(sym_a), Case, False)

											if True == gen11899 {
												gen11435 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

												V2654 := gen11435
												gen11456 := Call(__e, ShenFunc(symcons_2), V2654)

												var gen11457 Obj
												if True == gen11456 {
													gen11436 := Call(__e, ShenFunc(symhd), V2654)

													gen11437 := Call(__e, ShenFunc(symshen_4lazyderef), gen11436, V2783)

													V2655 := gen11437
													gen11455 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), V2655)

													if True == gen11455 {
														gen11438 := Call(__e, ShenFunc(symtl), V2654)

														gen11439 := Call(__e, ShenFunc(symshen_4lazyderef), gen11438, V2783)

														V2656 := gen11439
														gen11454 := Call(__e, ShenFunc(symcons_2), V2656)

														if True == gen11454 {
															gen11440 := Call(__e, ShenFunc(symhd), V2656)

															X := gen11440
															gen11441 := Call(__e, ShenFunc(symtl), V2656)

															gen11442 := Call(__e, ShenFunc(symshen_4lazyderef), gen11441, V2783)

															V2657 := gen11442
															gen11453 := Call(__e, ShenFunc(symcons_2), V2657)

															if True == gen11453 {
																gen11443 := Call(__e, ShenFunc(symhd), V2657)

																Y := gen11443
																gen11444 := Call(__e, ShenFunc(symtl), V2657)

																gen11445 := Call(__e, ShenFunc(symshen_4lazyderef), gen11444, V2783)

																V2658 := gen11445
																gen11452 := Call(__e, ShenFunc(sym_a), Nil, V2658)

																if True == gen11452 {
																	gen11446 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																	V2659 := gen11446
																	gen11451 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2659)

																	if True == gen11451 {
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen11450 := MakeNative(func(__e Evaluator) {
																			__e.TailApply(ShenFunc(symshen_4th_d), Y, MakeSymbol("string"), V2782, V2783, V2784)

																			return
																		}, 0)
																		gen11457 = Call(__e, ShenFunc(symshen_4th_d), X, MakeSymbol("string"), V2782, V2783, gen11450)

																	} else {
																		gen11449 := Call(__e, ShenFunc(symshen_4pvar_2), V2659)

																		if True == gen11449 {
																			Call(__e, ShenFunc(symshen_4bindv), V2659, MakeSymbol("string"), V2783)
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen11447 := MakeNative(func(__e Evaluator) {
																				__e.TailApply(ShenFunc(symshen_4th_d), Y, MakeSymbol("string"), V2782, V2783, V2784)

																				return
																			}, 0)
																			gen11448 := Call(__e, ShenFunc(symshen_4th_d), X, MakeSymbol("string"), V2782, V2783, gen11447)

																			Result := gen11448
																			Call(__e, ShenFunc(symshen_4unbindv), V2659, V2783)
																			gen11457 = Result

																		} else {
																			gen11457 = False
																		}

																	}

																} else {
																	gen11457 = False
																}

															} else {
																gen11457 = False
															}

														} else {
															gen11457 = False
														}

													} else {
														gen11457 = False
													}

												} else {
													gen11457 = False
												}
												Case := gen11457
												gen11898 := Call(__e, ShenFunc(sym_a), Case, False)

												if True == gen11898 {
													gen11458 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

													V2660 := gen11458
													gen11626 := Call(__e, ShenFunc(symcons_2), V2660)

													var gen11627 Obj
													if True == gen11626 {
														gen11459 := Call(__e, ShenFunc(symhd), V2660)

														gen11460 := Call(__e, ShenFunc(symshen_4lazyderef), gen11459, V2783)

														V2661 := gen11460
														gen11625 := Call(__e, ShenFunc(sym_a), MakeSymbol("lambda"), V2661)

														if True == gen11625 {
															gen11461 := Call(__e, ShenFunc(symtl), V2660)

															gen11462 := Call(__e, ShenFunc(symshen_4lazyderef), gen11461, V2783)

															V2662 := gen11462
															gen11624 := Call(__e, ShenFunc(symcons_2), V2662)

															if True == gen11624 {
																gen11463 := Call(__e, ShenFunc(symhd), V2662)

																X := gen11463
																gen11464 := Call(__e, ShenFunc(symtl), V2662)

																gen11465 := Call(__e, ShenFunc(symshen_4lazyderef), gen11464, V2783)

																V2663 := gen11465
																gen11623 := Call(__e, ShenFunc(symcons_2), V2663)

																if True == gen11623 {
																	gen11466 := Call(__e, ShenFunc(symhd), V2663)

																	Y := gen11466
																	gen11467 := Call(__e, ShenFunc(symtl), V2663)

																	gen11468 := Call(__e, ShenFunc(symshen_4lazyderef), gen11467, V2783)

																	V2664 := gen11468
																	gen11622 := Call(__e, ShenFunc(sym_a), Nil, V2664)

																	if True == gen11622 {
																		gen11469 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																		V2665 := gen11469
																		gen11621 := Call(__e, ShenFunc(symcons_2), V2665)

																		if True == gen11621 {
																			gen11490 := Call(__e, ShenFunc(symhd), V2665)

																			A := gen11490
																			gen11491 := Call(__e, ShenFunc(symtl), V2665)

																			gen11492 := Call(__e, ShenFunc(symshen_4lazyderef), gen11491, V2783)

																			V2666 := gen11492
																			gen11620 := Call(__e, ShenFunc(symcons_2), V2666)

																			if True == gen11620 {
																				gen11511 := Call(__e, ShenFunc(symhd), V2666)

																				gen11512 := Call(__e, ShenFunc(symshen_4lazyderef), gen11511, V2783)

																				V2667 := gen11512
																				gen11619 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), V2667)

																				if True == gen11619 {
																					gen11567 := Call(__e, ShenFunc(symtl), V2666)

																					gen11568 := Call(__e, ShenFunc(symshen_4lazyderef), gen11567, V2783)

																					V2668 := gen11568
																					gen11618 := Call(__e, ShenFunc(symcons_2), V2668)

																					if True == gen11618 {
																						gen11586 := Call(__e, ShenFunc(symhd), V2668)

																						B := gen11586
																						gen11587 := Call(__e, ShenFunc(symtl), V2668)

																						gen11588 := Call(__e, ShenFunc(symshen_4lazyderef), gen11587, V2783)

																						V2669 := gen11588
																						gen11617 := Call(__e, ShenFunc(sym_a), Nil, V2669)

																						if True == gen11617 {
																							gen11604 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							Z := gen11604
																							gen11605 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							X_e_e := gen11605
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen11606 := Call(__e, ShenFunc(symshen_4placeholder))

																							gen11616 := MakeNative(func(__e Evaluator) {
																								gen11607 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																								gen11608 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																								gen11609 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																								gen11610 := Call(__e, ShenFunc(symshen_4ebr), gen11607, gen11608, gen11609)

																								gen11615 := MakeNative(func(__e Evaluator) {
																									gen11611 := Call(__e, ShenFunc(symcons), A, Nil)

																									gen11612 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11611)

																									gen11613 := Call(__e, ShenFunc(symcons), X_e_e, gen11612)

																									gen11614 := Call(__e, ShenFunc(symcons), gen11613, V2782)

																									__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11614, V2783, V2784)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symbind), Z, gen11610, V2783, gen11615)

																								return

																							}, 0)
																							gen11627 = Call(__e, ShenFunc(symbind), X_e_e, gen11606, V2783, gen11616)

																						} else {
																							gen11603 := Call(__e, ShenFunc(symshen_4pvar_2), V2669)

																							if True == gen11603 {
																								Call(__e, ShenFunc(symshen_4bindv), V2669, Nil, V2783)
																								gen11589 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen11589
																								gen11590 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen11590
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11591 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen11601 := MakeNative(func(__e Evaluator) {
																									gen11592 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen11593 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen11594 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen11595 := Call(__e, ShenFunc(symshen_4ebr), gen11592, gen11593, gen11594)

																									gen11600 := MakeNative(func(__e Evaluator) {
																										gen11596 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen11597 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11596)

																										gen11598 := Call(__e, ShenFunc(symcons), X_e_e, gen11597)

																										gen11599 := Call(__e, ShenFunc(symcons), gen11598, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11599, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen11595, V2783, gen11600)

																									return

																								}, 0)
																								gen11602 := Call(__e, ShenFunc(symbind), X_e_e, gen11591, V2783, gen11601)

																								Result := gen11602
																								Call(__e, ShenFunc(symshen_4unbindv), V2669, V2783)
																								gen11627 = Result

																							} else {
																								gen11627 = False
																							}

																						}

																					} else {
																						gen11585 := Call(__e, ShenFunc(symshen_4pvar_2), V2668)

																						if True == gen11585 {
																							gen11569 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							B := gen11569
																							gen11570 := Call(__e, ShenFunc(symcons), B, Nil)

																							Call(__e, ShenFunc(symshen_4bindv), V2668, gen11570, V2783)

																							gen11571 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							Z := gen11571
																							gen11572 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																							X_e_e := gen11572
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen11573 := Call(__e, ShenFunc(symshen_4placeholder))

																							gen11583 := MakeNative(func(__e Evaluator) {
																								gen11574 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																								gen11575 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																								gen11576 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																								gen11577 := Call(__e, ShenFunc(symshen_4ebr), gen11574, gen11575, gen11576)

																								gen11582 := MakeNative(func(__e Evaluator) {
																									gen11578 := Call(__e, ShenFunc(symcons), A, Nil)

																									gen11579 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11578)

																									gen11580 := Call(__e, ShenFunc(symcons), X_e_e, gen11579)

																									gen11581 := Call(__e, ShenFunc(symcons), gen11580, V2782)

																									__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11581, V2783, V2784)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symbind), Z, gen11577, V2783, gen11582)

																								return

																							}, 0)
																							gen11584 := Call(__e, ShenFunc(symbind), X_e_e, gen11573, V2783, gen11583)

																							Result := gen11584
																							Call(__e, ShenFunc(symshen_4unbindv), V2668, V2783)
																							gen11627 = Result

																						} else {
																							gen11627 = False
																						}

																					}

																				} else {
																					gen11566 := Call(__e, ShenFunc(symshen_4pvar_2), V2667)

																					if True == gen11566 {
																						Call(__e, ShenFunc(symshen_4bindv), V2667, MakeSymbol("-->"), V2783)
																						gen11513 := Call(__e, ShenFunc(symtl), V2666)

																						gen11514 := Call(__e, ShenFunc(symshen_4lazyderef), gen11513, V2783)

																						V2670 := gen11514
																						gen11564 := Call(__e, ShenFunc(symcons_2), V2670)

																						var gen11565 Obj
																						if True == gen11564 {
																							gen11532 := Call(__e, ShenFunc(symhd), V2670)

																							B := gen11532
																							gen11533 := Call(__e, ShenFunc(symtl), V2670)

																							gen11534 := Call(__e, ShenFunc(symshen_4lazyderef), gen11533, V2783)

																							V2671 := gen11534
																							gen11563 := Call(__e, ShenFunc(sym_a), Nil, V2671)

																							if True == gen11563 {
																								gen11550 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen11550
																								gen11551 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen11551
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11552 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen11562 := MakeNative(func(__e Evaluator) {
																									gen11553 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen11554 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen11555 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen11556 := Call(__e, ShenFunc(symshen_4ebr), gen11553, gen11554, gen11555)

																									gen11561 := MakeNative(func(__e Evaluator) {
																										gen11557 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen11558 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11557)

																										gen11559 := Call(__e, ShenFunc(symcons), X_e_e, gen11558)

																										gen11560 := Call(__e, ShenFunc(symcons), gen11559, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11560, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen11556, V2783, gen11561)

																									return

																								}, 0)
																								gen11565 = Call(__e, ShenFunc(symbind), X_e_e, gen11552, V2783, gen11562)

																							} else {
																								gen11549 := Call(__e, ShenFunc(symshen_4pvar_2), V2671)

																								if True == gen11549 {
																									Call(__e, ShenFunc(symshen_4bindv), V2671, Nil, V2783)
																									gen11535 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									Z := gen11535
																									gen11536 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									X_e_e := gen11536
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11537 := Call(__e, ShenFunc(symshen_4placeholder))

																									gen11547 := MakeNative(func(__e Evaluator) {
																										gen11538 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																										gen11539 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																										gen11540 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																										gen11541 := Call(__e, ShenFunc(symshen_4ebr), gen11538, gen11539, gen11540)

																										gen11546 := MakeNative(func(__e Evaluator) {
																											gen11542 := Call(__e, ShenFunc(symcons), A, Nil)

																											gen11543 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11542)

																											gen11544 := Call(__e, ShenFunc(symcons), X_e_e, gen11543)

																											gen11545 := Call(__e, ShenFunc(symcons), gen11544, V2782)

																											__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11545, V2783, V2784)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symbind), Z, gen11541, V2783, gen11546)

																										return

																									}, 0)
																									gen11548 := Call(__e, ShenFunc(symbind), X_e_e, gen11537, V2783, gen11547)

																									Result := gen11548
																									Call(__e, ShenFunc(symshen_4unbindv), V2671, V2783)
																									gen11565 = Result

																								} else {
																									gen11565 = False
																								}

																							}

																						} else {
																							gen11531 := Call(__e, ShenFunc(symshen_4pvar_2), V2670)

																							if True == gen11531 {
																								gen11515 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								B := gen11515
																								gen11516 := Call(__e, ShenFunc(symcons), B, Nil)

																								Call(__e, ShenFunc(symshen_4bindv), V2670, gen11516, V2783)

																								gen11517 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Z := gen11517
																								gen11518 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								X_e_e := gen11518
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11519 := Call(__e, ShenFunc(symshen_4placeholder))

																								gen11529 := MakeNative(func(__e Evaluator) {
																									gen11520 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																									gen11521 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																									gen11522 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																									gen11523 := Call(__e, ShenFunc(symshen_4ebr), gen11520, gen11521, gen11522)

																									gen11528 := MakeNative(func(__e Evaluator) {
																										gen11524 := Call(__e, ShenFunc(symcons), A, Nil)

																										gen11525 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11524)

																										gen11526 := Call(__e, ShenFunc(symcons), X_e_e, gen11525)

																										gen11527 := Call(__e, ShenFunc(symcons), gen11526, V2782)

																										__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11527, V2783, V2784)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symbind), Z, gen11523, V2783, gen11528)

																									return

																								}, 0)
																								gen11530 := Call(__e, ShenFunc(symbind), X_e_e, gen11519, V2783, gen11529)

																								Result := gen11530
																								Call(__e, ShenFunc(symshen_4unbindv), V2670, V2783)
																								gen11565 = Result

																							} else {
																								gen11565 = False
																							}

																						}
																						Result := gen11565
																						Call(__e, ShenFunc(symshen_4unbindv), V2667, V2783)
																						gen11627 = Result

																					} else {
																						gen11627 = False
																					}

																				}

																			} else {
																				gen11510 := Call(__e, ShenFunc(symshen_4pvar_2), V2666)

																				if True == gen11510 {
																					gen11493 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					B := gen11493
																					gen11494 := Call(__e, ShenFunc(symcons), B, Nil)

																					gen11495 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen11494)

																					Call(__e, ShenFunc(symshen_4bindv), V2666, gen11495, V2783)

																					gen11496 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					Z := gen11496
																					gen11497 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																					X_e_e := gen11497
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11498 := Call(__e, ShenFunc(symshen_4placeholder))

																					gen11508 := MakeNative(func(__e Evaluator) {
																						gen11499 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																						gen11500 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																						gen11501 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																						gen11502 := Call(__e, ShenFunc(symshen_4ebr), gen11499, gen11500, gen11501)

																						gen11507 := MakeNative(func(__e Evaluator) {
																							gen11503 := Call(__e, ShenFunc(symcons), A, Nil)

																							gen11504 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11503)

																							gen11505 := Call(__e, ShenFunc(symcons), X_e_e, gen11504)

																							gen11506 := Call(__e, ShenFunc(symcons), gen11505, V2782)

																							__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11506, V2783, V2784)

																							return

																						}, 0)
																						__e.TailApply(ShenFunc(symbind), Z, gen11502, V2783, gen11507)

																						return

																					}, 0)
																					gen11509 := Call(__e, ShenFunc(symbind), X_e_e, gen11498, V2783, gen11508)

																					Result := gen11509
																					Call(__e, ShenFunc(symshen_4unbindv), V2666, V2783)
																					gen11627 = Result

																				} else {
																					gen11627 = False
																				}

																			}

																		} else {
																			gen11489 := Call(__e, ShenFunc(symshen_4pvar_2), V2665)

																			if True == gen11489 {
																				gen11470 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				A := gen11470
																				gen11471 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen11471
																				gen11472 := Call(__e, ShenFunc(symcons), B, Nil)

																				gen11473 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen11472)

																				gen11474 := Call(__e, ShenFunc(symcons), A, gen11473)

																				Call(__e, ShenFunc(symshen_4bindv), V2665, gen11474, V2783)

																				gen11475 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				Z := gen11475
																				gen11476 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				X_e_e := gen11476
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11477 := Call(__e, ShenFunc(symshen_4placeholder))

																				gen11487 := MakeNative(func(__e Evaluator) {
																					gen11478 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																					gen11479 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																					gen11480 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2783)

																					gen11481 := Call(__e, ShenFunc(symshen_4ebr), gen11478, gen11479, gen11480)

																					gen11486 := MakeNative(func(__e Evaluator) {
																						gen11482 := Call(__e, ShenFunc(symcons), A, Nil)

																						gen11483 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11482)

																						gen11484 := Call(__e, ShenFunc(symcons), X_e_e, gen11483)

																						gen11485 := Call(__e, ShenFunc(symcons), gen11484, V2782)

																						__e.TailApply(ShenFunc(symshen_4th_d), Z, B, gen11485, V2783, V2784)

																						return

																					}, 0)
																					__e.TailApply(ShenFunc(symbind), Z, gen11481, V2783, gen11486)

																					return

																				}, 0)
																				gen11488 := Call(__e, ShenFunc(symbind), X_e_e, gen11477, V2783, gen11487)

																				Result := gen11488
																				Call(__e, ShenFunc(symshen_4unbindv), V2665, V2783)
																				gen11627 = Result

																			} else {
																				gen11627 = False
																			}

																		}

																	} else {
																		gen11627 = False
																	}

																} else {
																	gen11627 = False
																}

															} else {
																gen11627 = False
															}

														} else {
															gen11627 = False
														}

													} else {
														gen11627 = False
													}
													Case := gen11627
													gen11897 := Call(__e, ShenFunc(sym_a), Case, False)

													if True == gen11897 {
														gen11628 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

														V2672 := gen11628
														gen11662 := Call(__e, ShenFunc(symcons_2), V2672)

														var gen11663 Obj
														if True == gen11662 {
															gen11629 := Call(__e, ShenFunc(symhd), V2672)

															gen11630 := Call(__e, ShenFunc(symshen_4lazyderef), gen11629, V2783)

															V2673 := gen11630
															gen11661 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), V2673)

															if True == gen11661 {
																gen11631 := Call(__e, ShenFunc(symtl), V2672)

																gen11632 := Call(__e, ShenFunc(symshen_4lazyderef), gen11631, V2783)

																V2674 := gen11632
																gen11660 := Call(__e, ShenFunc(symcons_2), V2674)

																if True == gen11660 {
																	gen11633 := Call(__e, ShenFunc(symhd), V2674)

																	X := gen11633
																	gen11634 := Call(__e, ShenFunc(symtl), V2674)

																	gen11635 := Call(__e, ShenFunc(symshen_4lazyderef), gen11634, V2783)

																	V2675 := gen11635
																	gen11659 := Call(__e, ShenFunc(symcons_2), V2675)

																	if True == gen11659 {
																		gen11636 := Call(__e, ShenFunc(symhd), V2675)

																		Y := gen11636
																		gen11637 := Call(__e, ShenFunc(symtl), V2675)

																		gen11638 := Call(__e, ShenFunc(symshen_4lazyderef), gen11637, V2783)

																		V2676 := gen11638
																		gen11658 := Call(__e, ShenFunc(symcons_2), V2676)

																		if True == gen11658 {
																			gen11639 := Call(__e, ShenFunc(symhd), V2676)

																			Z := gen11639
																			gen11640 := Call(__e, ShenFunc(symtl), V2676)

																			gen11641 := Call(__e, ShenFunc(symshen_4lazyderef), gen11640, V2783)

																			V2677 := gen11641
																			gen11657 := Call(__e, ShenFunc(sym_a), Nil, V2677)

																			if True == gen11657 {
																				gen11642 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				W := gen11642
																				gen11643 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				X_e_e := gen11643
																				gen11644 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																				B := gen11644
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11656 := MakeNative(func(__e Evaluator) {
																					gen11645 := Call(__e, ShenFunc(symshen_4placeholder))

																					gen11655 := MakeNative(func(__e Evaluator) {
																						gen11646 := Call(__e, ShenFunc(symshen_4lazyderef), X_e_e, V2783)

																						gen11647 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2783)

																						gen11648 := Call(__e, ShenFunc(symshen_4lazyderef), Z, V2783)

																						gen11649 := Call(__e, ShenFunc(symshen_4ebr), gen11646, gen11647, gen11648)

																						gen11654 := MakeNative(func(__e Evaluator) {
																							gen11650 := Call(__e, ShenFunc(symcons), B, Nil)

																							gen11651 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11650)

																							gen11652 := Call(__e, ShenFunc(symcons), X_e_e, gen11651)

																							gen11653 := Call(__e, ShenFunc(symcons), gen11652, V2782)

																							__e.TailApply(ShenFunc(symshen_4th_d), W, V2781, gen11653, V2783, V2784)

																							return

																						}, 0)
																						__e.TailApply(ShenFunc(symbind), W, gen11649, V2783, gen11654)

																						return

																					}, 0)
																					__e.TailApply(ShenFunc(symbind), X_e_e, gen11645, V2783, gen11655)

																					return

																				}, 0)
																				gen11663 = Call(__e, ShenFunc(symshen_4th_d), Y, B, V2782, V2783, gen11656)

																			} else {
																				gen11663 = False
																			}

																		} else {
																			gen11663 = False
																		}

																	} else {
																		gen11663 = False
																	}

																} else {
																	gen11663 = False
																}

															} else {
																gen11663 = False
															}

														} else {
															gen11663 = False
														}
														Case := gen11663
														gen11896 := Call(__e, ShenFunc(sym_a), Case, False)

														if True == gen11896 {
															gen11664 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

															V2678 := gen11664
															gen11766 := Call(__e, ShenFunc(symcons_2), V2678)

															var gen11767 Obj
															if True == gen11766 {
																gen11665 := Call(__e, ShenFunc(symhd), V2678)

																gen11666 := Call(__e, ShenFunc(symshen_4lazyderef), gen11665, V2783)

																V2679 := gen11666
																gen11765 := Call(__e, ShenFunc(sym_a), MakeSymbol("open"), V2679)

																if True == gen11765 {
																	gen11667 := Call(__e, ShenFunc(symtl), V2678)

																	gen11668 := Call(__e, ShenFunc(symshen_4lazyderef), gen11667, V2783)

																	V2680 := gen11668
																	gen11764 := Call(__e, ShenFunc(symcons_2), V2680)

																	if True == gen11764 {
																		gen11669 := Call(__e, ShenFunc(symhd), V2680)

																		FileName := gen11669
																		gen11670 := Call(__e, ShenFunc(symtl), V2680)

																		gen11671 := Call(__e, ShenFunc(symshen_4lazyderef), gen11670, V2783)

																		V2681 := gen11671
																		gen11763 := Call(__e, ShenFunc(symcons_2), V2681)

																		if True == gen11763 {
																			gen11672 := Call(__e, ShenFunc(symhd), V2681)

																			Direction2611 := gen11672
																			gen11673 := Call(__e, ShenFunc(symtl), V2681)

																			gen11674 := Call(__e, ShenFunc(symshen_4lazyderef), gen11673, V2783)

																			V2682 := gen11674
																			gen11762 := Call(__e, ShenFunc(sym_a), Nil, V2682)

																			if True == gen11762 {
																				gen11675 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																				V2683 := gen11675
																				gen11761 := Call(__e, ShenFunc(symcons_2), V2683)

																				if True == gen11761 {
																					gen11688 := Call(__e, ShenFunc(symhd), V2683)

																					gen11689 := Call(__e, ShenFunc(symshen_4lazyderef), gen11688, V2783)

																					V2684 := gen11689
																					gen11760 := Call(__e, ShenFunc(sym_a), MakeSymbol("stream"), V2684)

																					if True == gen11760 {
																						gen11726 := Call(__e, ShenFunc(symtl), V2683)

																						gen11727 := Call(__e, ShenFunc(symshen_4lazyderef), gen11726, V2783)

																						V2685 := gen11727
																						gen11759 := Call(__e, ShenFunc(symcons_2), V2685)

																						if True == gen11759 {
																							gen11739 := Call(__e, ShenFunc(symhd), V2685)

																							Direction := gen11739
																							gen11740 := Call(__e, ShenFunc(symtl), V2685)

																							gen11741 := Call(__e, ShenFunc(symshen_4lazyderef), gen11740, V2783)

																							V2686 := gen11741
																							gen11758 := Call(__e, ShenFunc(sym_a), Nil, V2686)

																							if True == gen11758 {
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11757 := MakeNative(func(__e Evaluator) {
																									gen11756 := MakeNative(func(__e Evaluator) {
																										gen11751 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																										gen11752 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																										gen11753 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11752)

																										gen11754 := Call(__e, ShenFunc(symelement_2), gen11751, gen11753)

																										gen11755 := MakeNative(func(__e Evaluator) {
																											__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symfwhen), gen11754, V2783, gen11755)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11756)

																									return

																								}, 0)
																								gen11767 = Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11757)

																							} else {
																								gen11750 := Call(__e, ShenFunc(symshen_4pvar_2), V2686)

																								if True == gen11750 {
																									Call(__e, ShenFunc(symshen_4bindv), V2686, Nil, V2783)
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11748 := MakeNative(func(__e Evaluator) {
																										gen11747 := MakeNative(func(__e Evaluator) {
																											gen11742 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen11743 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen11744 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11743)

																											gen11745 := Call(__e, ShenFunc(symelement_2), gen11742, gen11744)

																											gen11746 := MakeNative(func(__e Evaluator) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen11745, V2783, gen11746)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11747)

																										return

																									}, 0)
																									gen11749 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11748)

																									Result := gen11749
																									Call(__e, ShenFunc(symshen_4unbindv), V2686, V2783)
																									gen11767 = Result

																								} else {
																									gen11767 = False
																								}

																							}

																						} else {
																							gen11738 := Call(__e, ShenFunc(symshen_4pvar_2), V2685)

																							if True == gen11738 {
																								gen11728 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Direction := gen11728
																								gen11729 := Call(__e, ShenFunc(symcons), Direction, Nil)

																								Call(__e, ShenFunc(symshen_4bindv), V2685, gen11729, V2783)

																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11736 := MakeNative(func(__e Evaluator) {
																									gen11735 := MakeNative(func(__e Evaluator) {
																										gen11730 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																										gen11731 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																										gen11732 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11731)

																										gen11733 := Call(__e, ShenFunc(symelement_2), gen11730, gen11732)

																										gen11734 := MakeNative(func(__e Evaluator) {
																											__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symfwhen), gen11733, V2783, gen11734)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11735)

																									return

																								}, 0)
																								gen11737 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11736)

																								Result := gen11737
																								Call(__e, ShenFunc(symshen_4unbindv), V2685, V2783)
																								gen11767 = Result

																							} else {
																								gen11767 = False
																							}

																						}

																					} else {
																						gen11725 := Call(__e, ShenFunc(symshen_4pvar_2), V2684)

																						if True == gen11725 {
																							Call(__e, ShenFunc(symshen_4bindv), V2684, MakeSymbol("stream"), V2783)
																							gen11690 := Call(__e, ShenFunc(symtl), V2683)

																							gen11691 := Call(__e, ShenFunc(symshen_4lazyderef), gen11690, V2783)

																							V2687 := gen11691
																							gen11723 := Call(__e, ShenFunc(symcons_2), V2687)

																							var gen11724 Obj
																							if True == gen11723 {
																								gen11703 := Call(__e, ShenFunc(symhd), V2687)

																								Direction := gen11703
																								gen11704 := Call(__e, ShenFunc(symtl), V2687)

																								gen11705 := Call(__e, ShenFunc(symshen_4lazyderef), gen11704, V2783)

																								V2688 := gen11705
																								gen11722 := Call(__e, ShenFunc(sym_a), Nil, V2688)

																								if True == gen11722 {
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11721 := MakeNative(func(__e Evaluator) {
																										gen11720 := MakeNative(func(__e Evaluator) {
																											gen11715 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen11716 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen11717 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11716)

																											gen11718 := Call(__e, ShenFunc(symelement_2), gen11715, gen11717)

																											gen11719 := MakeNative(func(__e Evaluator) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen11718, V2783, gen11719)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11720)

																										return

																									}, 0)
																									gen11724 = Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11721)

																								} else {
																									gen11714 := Call(__e, ShenFunc(symshen_4pvar_2), V2688)

																									if True == gen11714 {
																										Call(__e, ShenFunc(symshen_4bindv), V2688, Nil, V2783)
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen11712 := MakeNative(func(__e Evaluator) {
																											gen11711 := MakeNative(func(__e Evaluator) {
																												gen11706 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																												gen11707 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																												gen11708 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11707)

																												gen11709 := Call(__e, ShenFunc(symelement_2), gen11706, gen11708)

																												gen11710 := MakeNative(func(__e Evaluator) {
																													__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																													return
																												}, 0)
																												__e.TailApply(ShenFunc(symfwhen), gen11709, V2783, gen11710)

																												return

																											}, 0)
																											__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11711)

																											return

																										}, 0)
																										gen11713 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11712)

																										Result := gen11713
																										Call(__e, ShenFunc(symshen_4unbindv), V2688, V2783)
																										gen11724 = Result

																									} else {
																										gen11724 = False
																									}

																								}

																							} else {
																								gen11702 := Call(__e, ShenFunc(symshen_4pvar_2), V2687)

																								if True == gen11702 {
																									gen11692 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																									Direction := gen11692
																									gen11693 := Call(__e, ShenFunc(symcons), Direction, Nil)

																									Call(__e, ShenFunc(symshen_4bindv), V2687, gen11693, V2783)

																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11700 := MakeNative(func(__e Evaluator) {
																										gen11699 := MakeNative(func(__e Evaluator) {
																											gen11694 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																											gen11695 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																											gen11696 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11695)

																											gen11697 := Call(__e, ShenFunc(symelement_2), gen11694, gen11696)

																											gen11698 := MakeNative(func(__e Evaluator) {
																												__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																												return
																											}, 0)
																											__e.TailApply(ShenFunc(symfwhen), gen11697, V2783, gen11698)

																											return

																										}, 0)
																										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11699)

																										return

																									}, 0)
																									gen11701 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11700)

																									Result := gen11701
																									Call(__e, ShenFunc(symshen_4unbindv), V2687, V2783)
																									gen11724 = Result

																								} else {
																									gen11724 = False
																								}

																							}
																							Result := gen11724
																							Call(__e, ShenFunc(symshen_4unbindv), V2684, V2783)
																							gen11767 = Result

																						} else {
																							gen11767 = False
																						}

																					}

																				} else {
																					gen11687 := Call(__e, ShenFunc(symshen_4pvar_2), V2683)

																					if True == gen11687 {
																						gen11676 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																						Direction := gen11676
																						gen11677 := Call(__e, ShenFunc(symcons), Direction, Nil)

																						gen11678 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen11677)

																						Call(__e, ShenFunc(symshen_4bindv), V2683, gen11678, V2783)

																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen11685 := MakeNative(func(__e Evaluator) {
																							gen11684 := MakeNative(func(__e Evaluator) {
																								gen11679 := Call(__e, ShenFunc(symshen_4lazyderef), Direction, V2783)

																								gen11680 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

																								gen11681 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen11680)

																								gen11682 := Call(__e, ShenFunc(symelement_2), gen11679, gen11681)

																								gen11683 := MakeNative(func(__e Evaluator) {
																									__e.TailApply(ShenFunc(symshen_4th_d), FileName, MakeSymbol("string"), V2782, V2783, V2784)

																									return
																								}, 0)
																								__e.TailApply(ShenFunc(symfwhen), gen11682, V2783, gen11683)

																								return

																							}, 0)
																							__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11684)

																							return

																						}, 0)
																						gen11686 := Call(__e, ShenFunc(symunify_b), Direction, Direction2611, V2783, gen11685)

																						Result := gen11686
																						Call(__e, ShenFunc(symshen_4unbindv), V2683, V2783)
																						gen11767 = Result

																					} else {
																						gen11767 = False
																					}

																				}

																			} else {
																				gen11767 = False
																			}

																		} else {
																			gen11767 = False
																		}

																	} else {
																		gen11767 = False
																	}

																} else {
																	gen11767 = False
																}

															} else {
																gen11767 = False
															}
															Case := gen11767
															gen11895 := Call(__e, ShenFunc(sym_a), Case, False)

															if True == gen11895 {
																gen11768 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																V2689 := gen11768
																gen11785 := Call(__e, ShenFunc(symcons_2), V2689)

																var gen11786 Obj
																if True == gen11785 {
																	gen11769 := Call(__e, ShenFunc(symhd), V2689)

																	gen11770 := Call(__e, ShenFunc(symshen_4lazyderef), gen11769, V2783)

																	V2690 := gen11770
																	gen11784 := Call(__e, ShenFunc(sym_a), MakeSymbol("type"), V2690)

																	if True == gen11784 {
																		gen11771 := Call(__e, ShenFunc(symtl), V2689)

																		gen11772 := Call(__e, ShenFunc(symshen_4lazyderef), gen11771, V2783)

																		V2691 := gen11772
																		gen11783 := Call(__e, ShenFunc(symcons_2), V2691)

																		if True == gen11783 {
																			gen11773 := Call(__e, ShenFunc(symhd), V2691)

																			X := gen11773
																			gen11774 := Call(__e, ShenFunc(symtl), V2691)

																			gen11775 := Call(__e, ShenFunc(symshen_4lazyderef), gen11774, V2783)

																			V2692 := gen11775
																			gen11782 := Call(__e, ShenFunc(symcons_2), V2692)

																			if True == gen11782 {
																				gen11776 := Call(__e, ShenFunc(symhd), V2692)

																				A := gen11776
																				gen11777 := Call(__e, ShenFunc(symtl), V2692)

																				gen11778 := Call(__e, ShenFunc(symshen_4lazyderef), gen11777, V2783)

																				V2693 := gen11778
																				gen11781 := Call(__e, ShenFunc(sym_a), Nil, V2693)

																				if True == gen11781 {
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen11780 := MakeNative(func(__e Evaluator) {
																						gen11779 := MakeNative(func(__e Evaluator) {
																							__e.TailApply(ShenFunc(symshen_4th_d), X, A, V2782, V2783, V2784)

																							return
																						}, 0)
																						__e.TailApply(ShenFunc(symunify), A, V2781, V2783, gen11779)

																						return

																					}, 0)
																					gen11786 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen11780)

																				} else {
																					gen11786 = False
																				}

																			} else {
																				gen11786 = False
																			}

																		} else {
																			gen11786 = False
																		}

																	} else {
																		gen11786 = False
																	}

																} else {
																	gen11786 = False
																}
																Case := gen11786
																gen11894 := Call(__e, ShenFunc(sym_a), Case, False)

																if True == gen11894 {
																	gen11787 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																	V2694 := gen11787
																	gen11809 := Call(__e, ShenFunc(symcons_2), V2694)

																	var gen11810 Obj
																	if True == gen11809 {
																		gen11788 := Call(__e, ShenFunc(symhd), V2694)

																		gen11789 := Call(__e, ShenFunc(symshen_4lazyderef), gen11788, V2783)

																		V2695 := gen11789
																		gen11808 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), V2695)

																		if True == gen11808 {
																			gen11790 := Call(__e, ShenFunc(symtl), V2694)

																			gen11791 := Call(__e, ShenFunc(symshen_4lazyderef), gen11790, V2783)

																			V2696 := gen11791
																			gen11807 := Call(__e, ShenFunc(symcons_2), V2696)

																			if True == gen11807 {
																				gen11792 := Call(__e, ShenFunc(symhd), V2696)

																				A := gen11792
																				gen11793 := Call(__e, ShenFunc(symtl), V2696)

																				gen11794 := Call(__e, ShenFunc(symshen_4lazyderef), gen11793, V2783)

																				V2697 := gen11794
																				gen11806 := Call(__e, ShenFunc(symcons_2), V2697)

																				if True == gen11806 {
																					gen11795 := Call(__e, ShenFunc(symhd), V2697)

																					Stream := gen11795
																					gen11796 := Call(__e, ShenFunc(symtl), V2697)

																					gen11797 := Call(__e, ShenFunc(symshen_4lazyderef), gen11796, V2783)

																					V2698 := gen11797
																					gen11805 := Call(__e, ShenFunc(sym_a), Nil, V2698)

																					if True == gen11805 {
																						gen11798 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																						C := gen11798
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen11799 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2783)

																						gen11800 := Call(__e, ShenFunc(symshen_4demodulate), gen11799)

																						gen11804 := MakeNative(func(__e Evaluator) {
																							gen11803 := MakeNative(func(__e Evaluator) {
																								gen11801 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

																								gen11802 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen11801)

																								__e.TailApply(ShenFunc(symshen_4th_d), Stream, gen11802, V2782, V2783, V2784)

																								return

																							}, 0)
																							__e.TailApply(ShenFunc(symunify), V2781, C, V2783, gen11803)

																							return

																						}, 0)
																						gen11810 = Call(__e, ShenFunc(symbind), C, gen11800, V2783, gen11804)

																					} else {
																						gen11810 = False
																					}

																				} else {
																					gen11810 = False
																				}

																			} else {
																				gen11810 = False
																			}

																		} else {
																			gen11810 = False
																		}

																	} else {
																		gen11810 = False
																	}
																	Case := gen11810
																	gen11893 := Call(__e, ShenFunc(sym_a), Case, False)

																	if True == gen11893 {
																		gen11811 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																		V2699 := gen11811
																		gen11832 := Call(__e, ShenFunc(symcons_2), V2699)

																		var gen11833 Obj
																		if True == gen11832 {
																			gen11812 := Call(__e, ShenFunc(symhd), V2699)

																			gen11813 := Call(__e, ShenFunc(symshen_4lazyderef), gen11812, V2783)

																			V2700 := gen11813
																			gen11831 := Call(__e, ShenFunc(sym_a), MakeSymbol("set"), V2700)

																			if True == gen11831 {
																				gen11814 := Call(__e, ShenFunc(symtl), V2699)

																				gen11815 := Call(__e, ShenFunc(symshen_4lazyderef), gen11814, V2783)

																				V2701 := gen11815
																				gen11830 := Call(__e, ShenFunc(symcons_2), V2701)

																				if True == gen11830 {
																					gen11816 := Call(__e, ShenFunc(symhd), V2701)

																					Var := gen11816
																					gen11817 := Call(__e, ShenFunc(symtl), V2701)

																					gen11818 := Call(__e, ShenFunc(symshen_4lazyderef), gen11817, V2783)

																					V2702 := gen11818
																					gen11829 := Call(__e, ShenFunc(symcons_2), V2702)

																					if True == gen11829 {
																						gen11819 := Call(__e, ShenFunc(symhd), V2702)

																						Val := gen11819
																						gen11820 := Call(__e, ShenFunc(symtl), V2702)

																						gen11821 := Call(__e, ShenFunc(symshen_4lazyderef), gen11820, V2783)

																						V2703 := gen11821
																						gen11828 := Call(__e, ShenFunc(sym_a), Nil, V2703)

																						if True == gen11828 {
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen11827 := MakeNative(func(__e Evaluator) {
																								gen11826 := MakeNative(func(__e Evaluator) {
																									gen11825 := MakeNative(func(__e Evaluator) {
																										gen11822 := Call(__e, ShenFunc(symcons), Var, Nil)

																										gen11823 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen11822)

																										gen11824 := MakeNative(func(__e Evaluator) {
																											__e.TailApply(ShenFunc(symshen_4th_d), Val, V2781, V2782, V2783, V2784)

																											return
																										}, 0)
																										__e.TailApply(ShenFunc(symshen_4th_d), gen11823, V2781, V2782, V2783, gen11824)

																										return

																									}, 0)
																									__e.TailApply(ShenFunc(symcut), Throwcontrol, V2783, gen11825)

																									return

																								}, 0)
																								__e.TailApply(ShenFunc(symshen_4th_d), Var, MakeSymbol("symbol"), V2782, V2783, gen11826)

																								return

																							}, 0)
																							gen11833 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen11827)

																						} else {
																							gen11833 = False
																						}

																					} else {
																						gen11833 = False
																					}

																				} else {
																					gen11833 = False
																				}

																			} else {
																				gen11833 = False
																			}

																		} else {
																			gen11833 = False
																		}
																		Case := gen11833
																		gen11892 := Call(__e, ShenFunc(sym_a), Case, False)

																		if True == gen11892 {
																			gen11834 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																			NewHyp := gen11834
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen11835 := MakeNative(func(__e Evaluator) {
																				__e.TailApply(ShenFunc(symshen_4th_d), V2780, V2781, NewHyp, V2783, V2784)

																				return
																			}, 0)
																			gen11836 := Call(__e, ShenFunc(symshen_4t_d_1hyps), V2782, NewHyp, V2783, gen11835)

																			Case := gen11836
																			gen11891 := Call(__e, ShenFunc(sym_a), Case, False)

																			if True == gen11891 {
																				gen11837 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																				V2704 := gen11837
																				gen11849 := Call(__e, ShenFunc(symcons_2), V2704)

																				var gen11850 Obj
																				if True == gen11849 {
																					gen11838 := Call(__e, ShenFunc(symhd), V2704)

																					gen11839 := Call(__e, ShenFunc(symshen_4lazyderef), gen11838, V2783)

																					V2705 := gen11839
																					gen11848 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), V2705)

																					if True == gen11848 {
																						gen11840 := Call(__e, ShenFunc(symtl), V2704)

																						gen11841 := Call(__e, ShenFunc(symshen_4lazyderef), gen11840, V2783)

																						V2706 := gen11841
																						gen11847 := Call(__e, ShenFunc(symcons_2), V2706)

																						if True == gen11847 {
																							gen11842 := Call(__e, ShenFunc(symhd), V2706)

																							F := gen11842
																							gen11843 := Call(__e, ShenFunc(symtl), V2706)

																							X := gen11843
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen11846 := MakeNative(func(__e Evaluator) {
																								gen11844 := Call(__e, ShenFunc(symcons), F, X)

																								gen11845 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen11844)

																								__e.TailApply(ShenFunc(symshen_4t_d_1def), gen11845, V2781, V2782, V2783, V2784)

																								return

																							}, 0)
																							gen11850 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, gen11846)

																						} else {
																							gen11850 = False
																						}

																					} else {
																						gen11850 = False
																					}

																				} else {
																					gen11850 = False
																				}
																				Case := gen11850
																				gen11890 := Call(__e, ShenFunc(sym_a), Case, False)

																				if True == gen11890 {
																					gen11851 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																					V2707 := gen11851
																					gen11859 := Call(__e, ShenFunc(symcons_2), V2707)

																					var gen11860 Obj
																					if True == gen11859 {
																						gen11852 := Call(__e, ShenFunc(symhd), V2707)

																						gen11853 := Call(__e, ShenFunc(symshen_4lazyderef), gen11852, V2783)

																						V2708 := gen11853
																						gen11858 := Call(__e, ShenFunc(sym_a), MakeSymbol("defmacro"), V2708)

																						if True == gen11858 {
																							gen11854 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																							V2709 := gen11854
																							gen11857 := Call(__e, ShenFunc(sym_a), MakeSymbol("unit"), V2709)

																							if True == gen11857 {
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11860 = Call(__e, ShenFunc(symcut), Throwcontrol, V2783, V2784)

																							} else {
																								gen11856 := Call(__e, ShenFunc(symshen_4pvar_2), V2709)

																								if True == gen11856 {
																									Call(__e, ShenFunc(symshen_4bindv), V2709, MakeSymbol("unit"), V2783)
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11855 := Call(__e, ShenFunc(symcut), Throwcontrol, V2783, V2784)

																									Result := gen11855
																									Call(__e, ShenFunc(symshen_4unbindv), V2709, V2783)
																									gen11860 = Result

																								} else {
																									gen11860 = False
																								}

																							}

																						} else {
																							gen11860 = False
																						}

																					} else {
																						gen11860 = False
																					}
																					Case := gen11860
																					gen11889 := Call(__e, ShenFunc(sym_a), Case, False)

																					if True == gen11889 {
																						gen11861 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																						V2710 := gen11861
																						gen11869 := Call(__e, ShenFunc(symcons_2), V2710)

																						var gen11870 Obj
																						if True == gen11869 {
																							gen11862 := Call(__e, ShenFunc(symhd), V2710)

																							gen11863 := Call(__e, ShenFunc(symshen_4lazyderef), gen11862, V2783)

																							V2711 := gen11863
																							gen11868 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.process-datatype"), V2711)

																							if True == gen11868 {
																								gen11864 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																								V2712 := gen11864
																								gen11867 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2712)

																								if True == gen11867 {
																									Call(__e, ShenFunc(symshen_4incinfs))
																									gen11870 = Call(__e, ShenFunc(symthaw), V2784)

																								} else {
																									gen11866 := Call(__e, ShenFunc(symshen_4pvar_2), V2712)

																									if True == gen11866 {
																										Call(__e, ShenFunc(symshen_4bindv), V2712, MakeSymbol("symbol"), V2783)
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen11865 := Call(__e, ShenFunc(symthaw), V2784)

																										Result := gen11865
																										Call(__e, ShenFunc(symshen_4unbindv), V2712, V2783)
																										gen11870 = Result

																									} else {
																										gen11870 = False
																									}

																								}

																							} else {
																								gen11870 = False
																							}

																						} else {
																							gen11870 = False
																						}
																						Case := gen11870
																						gen11888 := Call(__e, ShenFunc(sym_a), Case, False)

																						if True == gen11888 {
																							gen11871 := Call(__e, ShenFunc(symshen_4lazyderef), V2780, V2783)

																							V2713 := gen11871
																							gen11879 := Call(__e, ShenFunc(symcons_2), V2713)

																							var gen11880 Obj
																							if True == gen11879 {
																								gen11872 := Call(__e, ShenFunc(symhd), V2713)

																								gen11873 := Call(__e, ShenFunc(symshen_4lazyderef), gen11872, V2783)

																								V2714 := gen11873
																								gen11878 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.synonyms-help"), V2714)

																								if True == gen11878 {
																									gen11874 := Call(__e, ShenFunc(symshen_4lazyderef), V2781, V2783)

																									V2715 := gen11874
																									gen11877 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2715)

																									if True == gen11877 {
																										Call(__e, ShenFunc(symshen_4incinfs))
																										gen11880 = Call(__e, ShenFunc(symthaw), V2784)

																									} else {
																										gen11876 := Call(__e, ShenFunc(symshen_4pvar_2), V2715)

																										if True == gen11876 {
																											Call(__e, ShenFunc(symshen_4bindv), V2715, MakeSymbol("symbol"), V2783)
																											Call(__e, ShenFunc(symshen_4incinfs))
																											gen11875 := Call(__e, ShenFunc(symthaw), V2784)

																											Result := gen11875
																											Call(__e, ShenFunc(symshen_4unbindv), V2715, V2783)
																											gen11880 = Result

																										} else {
																											gen11880 = False
																										}

																									}

																								} else {
																									gen11880 = False
																								}

																							} else {
																								gen11880 = False
																							}
																							Case := gen11880
																							gen11887 := Call(__e, ShenFunc(sym_a), Case, False)

																							if True == gen11887 {
																								gen11881 := Call(__e, ShenFunc(symshen_4newpv), V2783)

																								Datatypes := gen11881
																								Call(__e, ShenFunc(symshen_4incinfs))
																								gen11882 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*datatypes*"))

																								gen11886 := MakeNative(func(__e Evaluator) {
																									gen11883 := Call(__e, ShenFunc(symcons), V2781, Nil)

																									gen11884 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11883)

																									gen11885 := Call(__e, ShenFunc(symcons), V2780, gen11884)

																									__e.TailApply(ShenFunc(symshen_4udefs_d), gen11885, V2782, Datatypes, V2783, V2784)

																									return

																								}, 0)
																								gen11908 = Call(__e, ShenFunc(symbind), Datatypes, gen11882, V2783, gen11886)

																							} else {
																								gen11908 = Case
																							}

																						} else {
																							gen11908 = Case
																						}

																					} else {
																						gen11908 = Case
																					}

																				} else {
																					gen11908 = Case
																				}

																			} else {
																				gen11908 = Case
																			}

																		} else {
																			gen11908 = Case
																		}

																	} else {
																		gen11908 = Case
																	}

																} else {
																	gen11908 = Case
																}

															} else {
																gen11908 = Case
															}

														} else {
															gen11908 = Case
														}

													} else {
														gen11908 = Case
													}

												} else {
													gen11908 = Case
												}

											} else {
												gen11908 = Case
											}

										} else {
											gen11908 = Case
										}

									} else {
										gen11908 = Case
									}

								} else {
									gen11908 = Case
								}

							} else {
								gen11908 = Case
							}

						} else {
							gen11908 = Case
						}

					} else {
						gen11908 = Case
					}

				} else {
					gen11908 = Case
				}

			} else {
				gen11908 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen11908)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.th*"), gen11909)

		gen13024 := MakeNative(func(__e Evaluator) {
			V2789 := __e.Get(1)
			_ = V2789
			V2790 := __e.Get(2)
			_ = V2790
			V2791 := __e.Get(3)
			_ = V2791
			V2792 := __e.Get(4)
			_ = V2792
			gen11910 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

			V2526 := gen11910
			gen12238 := Call(__e, ShenFunc(symcons_2), V2526)

			var gen12239 Obj
			if True == gen12238 {
				gen11911 := Call(__e, ShenFunc(symhd), V2526)

				gen11912 := Call(__e, ShenFunc(symshen_4lazyderef), gen11911, V2791)

				V2527 := gen11912
				gen12237 := Call(__e, ShenFunc(symcons_2), V2527)

				if True == gen12237 {
					gen11913 := Call(__e, ShenFunc(symhd), V2527)

					gen11914 := Call(__e, ShenFunc(symshen_4lazyderef), gen11913, V2791)

					V2528 := gen11914
					gen12236 := Call(__e, ShenFunc(symcons_2), V2528)

					if True == gen12236 {
						gen11915 := Call(__e, ShenFunc(symhd), V2528)

						gen11916 := Call(__e, ShenFunc(symshen_4lazyderef), gen11915, V2791)

						V2529 := gen11916
						gen12235 := Call(__e, ShenFunc(sym_a), MakeSymbol("cons"), V2529)

						if True == gen12235 {
							gen11917 := Call(__e, ShenFunc(symtl), V2528)

							gen11918 := Call(__e, ShenFunc(symshen_4lazyderef), gen11917, V2791)

							V2530 := gen11918
							gen12234 := Call(__e, ShenFunc(symcons_2), V2530)

							if True == gen12234 {
								gen11919 := Call(__e, ShenFunc(symhd), V2530)

								X := gen11919
								gen11920 := Call(__e, ShenFunc(symtl), V2530)

								gen11921 := Call(__e, ShenFunc(symshen_4lazyderef), gen11920, V2791)

								V2531 := gen11921
								gen12233 := Call(__e, ShenFunc(symcons_2), V2531)

								if True == gen12233 {
									gen11922 := Call(__e, ShenFunc(symhd), V2531)

									Y := gen11922
									gen11923 := Call(__e, ShenFunc(symtl), V2531)

									gen11924 := Call(__e, ShenFunc(symshen_4lazyderef), gen11923, V2791)

									V2532 := gen11924
									gen12232 := Call(__e, ShenFunc(sym_a), Nil, V2532)

									if True == gen12232 {
										gen11925 := Call(__e, ShenFunc(symtl), V2527)

										gen11926 := Call(__e, ShenFunc(symshen_4lazyderef), gen11925, V2791)

										V2533 := gen11926
										gen12231 := Call(__e, ShenFunc(symcons_2), V2533)

										if True == gen12231 {
											gen11927 := Call(__e, ShenFunc(symhd), V2533)

											gen11928 := Call(__e, ShenFunc(symshen_4lazyderef), gen11927, V2791)

											V2534 := gen11928
											gen12230 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2534)

											if True == gen12230 {
												gen11929 := Call(__e, ShenFunc(symtl), V2533)

												gen11930 := Call(__e, ShenFunc(symshen_4lazyderef), gen11929, V2791)

												V2535 := gen11930
												gen12229 := Call(__e, ShenFunc(symcons_2), V2535)

												if True == gen12229 {
													gen11931 := Call(__e, ShenFunc(symhd), V2535)

													gen11932 := Call(__e, ShenFunc(symshen_4lazyderef), gen11931, V2791)

													V2536 := gen11932
													gen12228 := Call(__e, ShenFunc(symcons_2), V2536)

													if True == gen12228 {
														gen11975 := Call(__e, ShenFunc(symhd), V2536)

														gen11976 := Call(__e, ShenFunc(symshen_4lazyderef), gen11975, V2791)

														V2537 := gen11976
														gen12227 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2537)

														if True == gen12227 {
															gen12103 := Call(__e, ShenFunc(symtl), V2536)

															gen12104 := Call(__e, ShenFunc(symshen_4lazyderef), gen12103, V2791)

															V2538 := gen12104
															gen12226 := Call(__e, ShenFunc(symcons_2), V2538)

															if True == gen12226 {
																gen12146 := Call(__e, ShenFunc(symhd), V2538)

																A := gen12146
																gen12147 := Call(__e, ShenFunc(symtl), V2538)

																gen12148 := Call(__e, ShenFunc(symshen_4lazyderef), gen12147, V2791)

																V2539 := gen12148
																gen12225 := Call(__e, ShenFunc(sym_a), Nil, V2539)

																if True == gen12225 {
																	gen12188 := Call(__e, ShenFunc(symtl), V2535)

																	gen12189 := Call(__e, ShenFunc(symshen_4lazyderef), gen12188, V2791)

																	V2540 := gen12189
																	gen12224 := Call(__e, ShenFunc(sym_a), Nil, V2540)

																	if True == gen12224 {
																		gen12208 := Call(__e, ShenFunc(symtl), V2526)

																		Hyp := gen12208
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12209 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12210 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12211 := Call(__e, ShenFunc(symcons), gen12210, Nil)

																		gen12212 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12211)

																		gen12213 := Call(__e, ShenFunc(symcons), gen12209, gen12212)

																		gen12214 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12215 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12216 := Call(__e, ShenFunc(symcons), gen12215, Nil)

																		gen12217 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12216)

																		gen12218 := Call(__e, ShenFunc(symcons), gen12217, Nil)

																		gen12219 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12218)

																		gen12220 := Call(__e, ShenFunc(symcons), gen12214, gen12219)

																		gen12221 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12222 := Call(__e, ShenFunc(symcons), gen12220, gen12221)

																		gen12223 := Call(__e, ShenFunc(symcons), gen12213, gen12222)

																		gen12239 = Call(__e, ShenFunc(symbind), V2790, gen12223, V2791, V2792)

																	} else {
																		gen12207 := Call(__e, ShenFunc(symshen_4pvar_2), V2540)

																		if True == gen12207 {
																			Call(__e, ShenFunc(symshen_4bindv), V2540, Nil, V2791)
																			gen12190 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen12190
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12191 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12192 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12193 := Call(__e, ShenFunc(symcons), gen12192, Nil)

																			gen12194 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12193)

																			gen12195 := Call(__e, ShenFunc(symcons), gen12191, gen12194)

																			gen12196 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12197 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12198 := Call(__e, ShenFunc(symcons), gen12197, Nil)

																			gen12199 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12198)

																			gen12200 := Call(__e, ShenFunc(symcons), gen12199, Nil)

																			gen12201 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12200)

																			gen12202 := Call(__e, ShenFunc(symcons), gen12196, gen12201)

																			gen12203 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12204 := Call(__e, ShenFunc(symcons), gen12202, gen12203)

																			gen12205 := Call(__e, ShenFunc(symcons), gen12195, gen12204)

																			gen12206 := Call(__e, ShenFunc(symbind), V2790, gen12205, V2791, V2792)

																			Result := gen12206
																			Call(__e, ShenFunc(symshen_4unbindv), V2540, V2791)
																			gen12239 = Result

																		} else {
																			gen12239 = False
																		}

																	}

																} else {
																	gen12187 := Call(__e, ShenFunc(symshen_4pvar_2), V2539)

																	if True == gen12187 {
																		Call(__e, ShenFunc(symshen_4bindv), V2539, Nil, V2791)
																		gen12149 := Call(__e, ShenFunc(symtl), V2535)

																		gen12150 := Call(__e, ShenFunc(symshen_4lazyderef), gen12149, V2791)

																		V2541 := gen12150
																		gen12185 := Call(__e, ShenFunc(sym_a), Nil, V2541)

																		var gen12186 Obj
																		if True == gen12185 {
																			gen12169 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen12169
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12170 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12171 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12172 := Call(__e, ShenFunc(symcons), gen12171, Nil)

																			gen12173 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12172)

																			gen12174 := Call(__e, ShenFunc(symcons), gen12170, gen12173)

																			gen12175 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12176 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12177 := Call(__e, ShenFunc(symcons), gen12176, Nil)

																			gen12178 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12177)

																			gen12179 := Call(__e, ShenFunc(symcons), gen12178, Nil)

																			gen12180 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12179)

																			gen12181 := Call(__e, ShenFunc(symcons), gen12175, gen12180)

																			gen12182 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12183 := Call(__e, ShenFunc(symcons), gen12181, gen12182)

																			gen12184 := Call(__e, ShenFunc(symcons), gen12174, gen12183)

																			gen12186 = Call(__e, ShenFunc(symbind), V2790, gen12184, V2791, V2792)

																		} else {
																			gen12168 := Call(__e, ShenFunc(symshen_4pvar_2), V2541)

																			if True == gen12168 {
																				Call(__e, ShenFunc(symshen_4bindv), V2541, Nil, V2791)
																				gen12151 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen12151
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12152 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12153 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12154 := Call(__e, ShenFunc(symcons), gen12153, Nil)

																				gen12155 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12154)

																				gen12156 := Call(__e, ShenFunc(symcons), gen12152, gen12155)

																				gen12157 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12158 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12159 := Call(__e, ShenFunc(symcons), gen12158, Nil)

																				gen12160 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12159)

																				gen12161 := Call(__e, ShenFunc(symcons), gen12160, Nil)

																				gen12162 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12161)

																				gen12163 := Call(__e, ShenFunc(symcons), gen12157, gen12162)

																				gen12164 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12165 := Call(__e, ShenFunc(symcons), gen12163, gen12164)

																				gen12166 := Call(__e, ShenFunc(symcons), gen12156, gen12165)

																				gen12167 := Call(__e, ShenFunc(symbind), V2790, gen12166, V2791, V2792)

																				Result := gen12167
																				Call(__e, ShenFunc(symshen_4unbindv), V2541, V2791)
																				gen12186 = Result

																			} else {
																				gen12186 = False
																			}

																		}
																		Result := gen12186
																		Call(__e, ShenFunc(symshen_4unbindv), V2539, V2791)
																		gen12239 = Result

																	} else {
																		gen12239 = False
																	}

																}

															} else {
																gen12145 := Call(__e, ShenFunc(symshen_4pvar_2), V2538)

																if True == gen12145 {
																	gen12105 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	A := gen12105
																	gen12106 := Call(__e, ShenFunc(symcons), A, Nil)

																	Call(__e, ShenFunc(symshen_4bindv), V2538, gen12106, V2791)

																	gen12107 := Call(__e, ShenFunc(symtl), V2535)

																	gen12108 := Call(__e, ShenFunc(symshen_4lazyderef), gen12107, V2791)

																	V2542 := gen12108
																	gen12143 := Call(__e, ShenFunc(sym_a), Nil, V2542)

																	var gen12144 Obj
																	if True == gen12143 {
																		gen12127 := Call(__e, ShenFunc(symtl), V2526)

																		Hyp := gen12127
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12128 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12129 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12130 := Call(__e, ShenFunc(symcons), gen12129, Nil)

																		gen12131 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12130)

																		gen12132 := Call(__e, ShenFunc(symcons), gen12128, gen12131)

																		gen12133 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12134 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12135 := Call(__e, ShenFunc(symcons), gen12134, Nil)

																		gen12136 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12135)

																		gen12137 := Call(__e, ShenFunc(symcons), gen12136, Nil)

																		gen12138 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12137)

																		gen12139 := Call(__e, ShenFunc(symcons), gen12133, gen12138)

																		gen12140 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12141 := Call(__e, ShenFunc(symcons), gen12139, gen12140)

																		gen12142 := Call(__e, ShenFunc(symcons), gen12132, gen12141)

																		gen12144 = Call(__e, ShenFunc(symbind), V2790, gen12142, V2791, V2792)

																	} else {
																		gen12126 := Call(__e, ShenFunc(symshen_4pvar_2), V2542)

																		if True == gen12126 {
																			Call(__e, ShenFunc(symshen_4bindv), V2542, Nil, V2791)
																			gen12109 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen12109
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12110 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12111 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12112 := Call(__e, ShenFunc(symcons), gen12111, Nil)

																			gen12113 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12112)

																			gen12114 := Call(__e, ShenFunc(symcons), gen12110, gen12113)

																			gen12115 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12116 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12117 := Call(__e, ShenFunc(symcons), gen12116, Nil)

																			gen12118 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12117)

																			gen12119 := Call(__e, ShenFunc(symcons), gen12118, Nil)

																			gen12120 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12119)

																			gen12121 := Call(__e, ShenFunc(symcons), gen12115, gen12120)

																			gen12122 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12123 := Call(__e, ShenFunc(symcons), gen12121, gen12122)

																			gen12124 := Call(__e, ShenFunc(symcons), gen12114, gen12123)

																			gen12125 := Call(__e, ShenFunc(symbind), V2790, gen12124, V2791, V2792)

																			Result := gen12125
																			Call(__e, ShenFunc(symshen_4unbindv), V2542, V2791)
																			gen12144 = Result

																		} else {
																			gen12144 = False
																		}

																	}
																	Result := gen12144
																	Call(__e, ShenFunc(symshen_4unbindv), V2538, V2791)
																	gen12239 = Result

																} else {
																	gen12239 = False
																}

															}

														} else {
															gen12102 := Call(__e, ShenFunc(symshen_4pvar_2), V2537)

															if True == gen12102 {
																Call(__e, ShenFunc(symshen_4bindv), V2537, MakeSymbol("list"), V2791)
																gen11977 := Call(__e, ShenFunc(symtl), V2536)

																gen11978 := Call(__e, ShenFunc(symshen_4lazyderef), gen11977, V2791)

																V2543 := gen11978
																gen12100 := Call(__e, ShenFunc(symcons_2), V2543)

																var gen12101 Obj
																if True == gen12100 {
																	gen12020 := Call(__e, ShenFunc(symhd), V2543)

																	A := gen12020
																	gen12021 := Call(__e, ShenFunc(symtl), V2543)

																	gen12022 := Call(__e, ShenFunc(symshen_4lazyderef), gen12021, V2791)

																	V2544 := gen12022
																	gen12099 := Call(__e, ShenFunc(sym_a), Nil, V2544)

																	if True == gen12099 {
																		gen12062 := Call(__e, ShenFunc(symtl), V2535)

																		gen12063 := Call(__e, ShenFunc(symshen_4lazyderef), gen12062, V2791)

																		V2545 := gen12063
																		gen12098 := Call(__e, ShenFunc(sym_a), Nil, V2545)

																		if True == gen12098 {
																			gen12082 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen12082
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12083 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12084 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12085 := Call(__e, ShenFunc(symcons), gen12084, Nil)

																			gen12086 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12085)

																			gen12087 := Call(__e, ShenFunc(symcons), gen12083, gen12086)

																			gen12088 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12089 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12090 := Call(__e, ShenFunc(symcons), gen12089, Nil)

																			gen12091 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12090)

																			gen12092 := Call(__e, ShenFunc(symcons), gen12091, Nil)

																			gen12093 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12092)

																			gen12094 := Call(__e, ShenFunc(symcons), gen12088, gen12093)

																			gen12095 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12096 := Call(__e, ShenFunc(symcons), gen12094, gen12095)

																			gen12097 := Call(__e, ShenFunc(symcons), gen12087, gen12096)

																			gen12101 = Call(__e, ShenFunc(symbind), V2790, gen12097, V2791, V2792)

																		} else {
																			gen12081 := Call(__e, ShenFunc(symshen_4pvar_2), V2545)

																			if True == gen12081 {
																				Call(__e, ShenFunc(symshen_4bindv), V2545, Nil, V2791)
																				gen12064 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen12064
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12065 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12066 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12067 := Call(__e, ShenFunc(symcons), gen12066, Nil)

																				gen12068 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12067)

																				gen12069 := Call(__e, ShenFunc(symcons), gen12065, gen12068)

																				gen12070 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12071 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12072 := Call(__e, ShenFunc(symcons), gen12071, Nil)

																				gen12073 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12072)

																				gen12074 := Call(__e, ShenFunc(symcons), gen12073, Nil)

																				gen12075 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12074)

																				gen12076 := Call(__e, ShenFunc(symcons), gen12070, gen12075)

																				gen12077 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12078 := Call(__e, ShenFunc(symcons), gen12076, gen12077)

																				gen12079 := Call(__e, ShenFunc(symcons), gen12069, gen12078)

																				gen12080 := Call(__e, ShenFunc(symbind), V2790, gen12079, V2791, V2792)

																				Result := gen12080
																				Call(__e, ShenFunc(symshen_4unbindv), V2545, V2791)
																				gen12101 = Result

																			} else {
																				gen12101 = False
																			}

																		}

																	} else {
																		gen12061 := Call(__e, ShenFunc(symshen_4pvar_2), V2544)

																		if True == gen12061 {
																			Call(__e, ShenFunc(symshen_4bindv), V2544, Nil, V2791)
																			gen12023 := Call(__e, ShenFunc(symtl), V2535)

																			gen12024 := Call(__e, ShenFunc(symshen_4lazyderef), gen12023, V2791)

																			V2546 := gen12024
																			gen12059 := Call(__e, ShenFunc(sym_a), Nil, V2546)

																			var gen12060 Obj
																			if True == gen12059 {
																				gen12043 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen12043
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12044 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12045 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12046 := Call(__e, ShenFunc(symcons), gen12045, Nil)

																				gen12047 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12046)

																				gen12048 := Call(__e, ShenFunc(symcons), gen12044, gen12047)

																				gen12049 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12050 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12051 := Call(__e, ShenFunc(symcons), gen12050, Nil)

																				gen12052 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12051)

																				gen12053 := Call(__e, ShenFunc(symcons), gen12052, Nil)

																				gen12054 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12053)

																				gen12055 := Call(__e, ShenFunc(symcons), gen12049, gen12054)

																				gen12056 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12057 := Call(__e, ShenFunc(symcons), gen12055, gen12056)

																				gen12058 := Call(__e, ShenFunc(symcons), gen12048, gen12057)

																				gen12060 = Call(__e, ShenFunc(symbind), V2790, gen12058, V2791, V2792)

																			} else {
																				gen12042 := Call(__e, ShenFunc(symshen_4pvar_2), V2546)

																				if True == gen12042 {
																					Call(__e, ShenFunc(symshen_4bindv), V2546, Nil, V2791)
																					gen12025 := Call(__e, ShenFunc(symtl), V2526)

																					Hyp := gen12025
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12026 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12027 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12028 := Call(__e, ShenFunc(symcons), gen12027, Nil)

																					gen12029 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12028)

																					gen12030 := Call(__e, ShenFunc(symcons), gen12026, gen12029)

																					gen12031 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12032 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12033 := Call(__e, ShenFunc(symcons), gen12032, Nil)

																					gen12034 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12033)

																					gen12035 := Call(__e, ShenFunc(symcons), gen12034, Nil)

																					gen12036 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12035)

																					gen12037 := Call(__e, ShenFunc(symcons), gen12031, gen12036)

																					gen12038 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12039 := Call(__e, ShenFunc(symcons), gen12037, gen12038)

																					gen12040 := Call(__e, ShenFunc(symcons), gen12030, gen12039)

																					gen12041 := Call(__e, ShenFunc(symbind), V2790, gen12040, V2791, V2792)

																					Result := gen12041
																					Call(__e, ShenFunc(symshen_4unbindv), V2546, V2791)
																					gen12060 = Result

																				} else {
																					gen12060 = False
																				}

																			}
																			Result := gen12060
																			Call(__e, ShenFunc(symshen_4unbindv), V2544, V2791)
																			gen12101 = Result

																		} else {
																			gen12101 = False
																		}

																	}

																} else {
																	gen12019 := Call(__e, ShenFunc(symshen_4pvar_2), V2543)

																	if True == gen12019 {
																		gen11979 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																		A := gen11979
																		gen11980 := Call(__e, ShenFunc(symcons), A, Nil)

																		Call(__e, ShenFunc(symshen_4bindv), V2543, gen11980, V2791)

																		gen11981 := Call(__e, ShenFunc(symtl), V2535)

																		gen11982 := Call(__e, ShenFunc(symshen_4lazyderef), gen11981, V2791)

																		V2547 := gen11982
																		gen12017 := Call(__e, ShenFunc(sym_a), Nil, V2547)

																		var gen12018 Obj
																		if True == gen12017 {
																			gen12001 := Call(__e, ShenFunc(symtl), V2526)

																			Hyp := gen12001
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12002 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12003 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12004 := Call(__e, ShenFunc(symcons), gen12003, Nil)

																			gen12005 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12004)

																			gen12006 := Call(__e, ShenFunc(symcons), gen12002, gen12005)

																			gen12007 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12008 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12009 := Call(__e, ShenFunc(symcons), gen12008, Nil)

																			gen12010 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen12009)

																			gen12011 := Call(__e, ShenFunc(symcons), gen12010, Nil)

																			gen12012 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12011)

																			gen12013 := Call(__e, ShenFunc(symcons), gen12007, gen12012)

																			gen12014 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12015 := Call(__e, ShenFunc(symcons), gen12013, gen12014)

																			gen12016 := Call(__e, ShenFunc(symcons), gen12006, gen12015)

																			gen12018 = Call(__e, ShenFunc(symbind), V2790, gen12016, V2791, V2792)

																		} else {
																			gen12000 := Call(__e, ShenFunc(symshen_4pvar_2), V2547)

																			if True == gen12000 {
																				Call(__e, ShenFunc(symshen_4bindv), V2547, Nil, V2791)
																				gen11983 := Call(__e, ShenFunc(symtl), V2526)

																				Hyp := gen11983
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen11984 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen11985 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen11986 := Call(__e, ShenFunc(symcons), gen11985, Nil)

																				gen11987 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11986)

																				gen11988 := Call(__e, ShenFunc(symcons), gen11984, gen11987)

																				gen11989 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen11990 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen11991 := Call(__e, ShenFunc(symcons), gen11990, Nil)

																				gen11992 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11991)

																				gen11993 := Call(__e, ShenFunc(symcons), gen11992, Nil)

																				gen11994 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11993)

																				gen11995 := Call(__e, ShenFunc(symcons), gen11989, gen11994)

																				gen11996 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen11997 := Call(__e, ShenFunc(symcons), gen11995, gen11996)

																				gen11998 := Call(__e, ShenFunc(symcons), gen11988, gen11997)

																				gen11999 := Call(__e, ShenFunc(symbind), V2790, gen11998, V2791, V2792)

																				Result := gen11999
																				Call(__e, ShenFunc(symshen_4unbindv), V2547, V2791)
																				gen12018 = Result

																			} else {
																				gen12018 = False
																			}

																		}
																		Result := gen12018
																		Call(__e, ShenFunc(symshen_4unbindv), V2543, V2791)
																		gen12101 = Result

																	} else {
																		gen12101 = False
																	}

																}
																Result := gen12101
																Call(__e, ShenFunc(symshen_4unbindv), V2537, V2791)
																gen12239 = Result

															} else {
																gen12239 = False
															}

														}

													} else {
														gen11974 := Call(__e, ShenFunc(symshen_4pvar_2), V2536)

														if True == gen11974 {
															gen11933 := Call(__e, ShenFunc(symshen_4newpv), V2791)

															A := gen11933
															gen11934 := Call(__e, ShenFunc(symcons), A, Nil)

															gen11935 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11934)

															Call(__e, ShenFunc(symshen_4bindv), V2536, gen11935, V2791)

															gen11936 := Call(__e, ShenFunc(symtl), V2535)

															gen11937 := Call(__e, ShenFunc(symshen_4lazyderef), gen11936, V2791)

															V2548 := gen11937
															gen11972 := Call(__e, ShenFunc(sym_a), Nil, V2548)

															var gen11973 Obj
															if True == gen11972 {
																gen11956 := Call(__e, ShenFunc(symtl), V2526)

																Hyp := gen11956
																Call(__e, ShenFunc(symshen_4incinfs))
																gen11957 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																gen11958 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																gen11959 := Call(__e, ShenFunc(symcons), gen11958, Nil)

																gen11960 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11959)

																gen11961 := Call(__e, ShenFunc(symcons), gen11957, gen11960)

																gen11962 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																gen11963 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																gen11964 := Call(__e, ShenFunc(symcons), gen11963, Nil)

																gen11965 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11964)

																gen11966 := Call(__e, ShenFunc(symcons), gen11965, Nil)

																gen11967 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11966)

																gen11968 := Call(__e, ShenFunc(symcons), gen11962, gen11967)

																gen11969 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																gen11970 := Call(__e, ShenFunc(symcons), gen11968, gen11969)

																gen11971 := Call(__e, ShenFunc(symcons), gen11961, gen11970)

																gen11973 = Call(__e, ShenFunc(symbind), V2790, gen11971, V2791, V2792)

															} else {
																gen11955 := Call(__e, ShenFunc(symshen_4pvar_2), V2548)

																if True == gen11955 {
																	Call(__e, ShenFunc(symshen_4bindv), V2548, Nil, V2791)
																	gen11938 := Call(__e, ShenFunc(symtl), V2526)

																	Hyp := gen11938
																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen11939 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																	gen11940 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen11941 := Call(__e, ShenFunc(symcons), gen11940, Nil)

																	gen11942 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11941)

																	gen11943 := Call(__e, ShenFunc(symcons), gen11939, gen11942)

																	gen11944 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																	gen11945 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen11946 := Call(__e, ShenFunc(symcons), gen11945, Nil)

																	gen11947 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen11946)

																	gen11948 := Call(__e, ShenFunc(symcons), gen11947, Nil)

																	gen11949 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen11948)

																	gen11950 := Call(__e, ShenFunc(symcons), gen11944, gen11949)

																	gen11951 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																	gen11952 := Call(__e, ShenFunc(symcons), gen11950, gen11951)

																	gen11953 := Call(__e, ShenFunc(symcons), gen11943, gen11952)

																	gen11954 := Call(__e, ShenFunc(symbind), V2790, gen11953, V2791, V2792)

																	Result := gen11954
																	Call(__e, ShenFunc(symshen_4unbindv), V2548, V2791)
																	gen11973 = Result

																} else {
																	gen11973 = False
																}

															}
															Result := gen11973
															Call(__e, ShenFunc(symshen_4unbindv), V2536, V2791)
															gen12239 = Result

														} else {
															gen12239 = False
														}

													}

												} else {
													gen12239 = False
												}

											} else {
												gen12239 = False
											}

										} else {
											gen12239 = False
										}

									} else {
										gen12239 = False
									}

								} else {
									gen12239 = False
								}

							} else {
								gen12239 = False
							}

						} else {
							gen12239 = False
						}

					} else {
						gen12239 = False
					}

				} else {
					gen12239 = False
				}

			} else {
				gen12239 = False
			}
			Case := gen12239
			gen13023 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13023 {
				gen12240 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

				V2549 := gen12240
				gen12584 := Call(__e, ShenFunc(symcons_2), V2549)

				var gen12585 Obj
				if True == gen12584 {
					gen12241 := Call(__e, ShenFunc(symhd), V2549)

					gen12242 := Call(__e, ShenFunc(symshen_4lazyderef), gen12241, V2791)

					V2550 := gen12242
					gen12583 := Call(__e, ShenFunc(symcons_2), V2550)

					if True == gen12583 {
						gen12243 := Call(__e, ShenFunc(symhd), V2550)

						gen12244 := Call(__e, ShenFunc(symshen_4lazyderef), gen12243, V2791)

						V2551 := gen12244
						gen12582 := Call(__e, ShenFunc(symcons_2), V2551)

						if True == gen12582 {
							gen12245 := Call(__e, ShenFunc(symhd), V2551)

							gen12246 := Call(__e, ShenFunc(symshen_4lazyderef), gen12245, V2791)

							V2552 := gen12246
							gen12581 := Call(__e, ShenFunc(sym_a), MakeSymbol("@p"), V2552)

							if True == gen12581 {
								gen12247 := Call(__e, ShenFunc(symtl), V2551)

								gen12248 := Call(__e, ShenFunc(symshen_4lazyderef), gen12247, V2791)

								V2553 := gen12248
								gen12580 := Call(__e, ShenFunc(symcons_2), V2553)

								if True == gen12580 {
									gen12249 := Call(__e, ShenFunc(symhd), V2553)

									X := gen12249
									gen12250 := Call(__e, ShenFunc(symtl), V2553)

									gen12251 := Call(__e, ShenFunc(symshen_4lazyderef), gen12250, V2791)

									V2554 := gen12251
									gen12579 := Call(__e, ShenFunc(symcons_2), V2554)

									if True == gen12579 {
										gen12252 := Call(__e, ShenFunc(symhd), V2554)

										Y := gen12252
										gen12253 := Call(__e, ShenFunc(symtl), V2554)

										gen12254 := Call(__e, ShenFunc(symshen_4lazyderef), gen12253, V2791)

										V2555 := gen12254
										gen12578 := Call(__e, ShenFunc(sym_a), Nil, V2555)

										if True == gen12578 {
											gen12255 := Call(__e, ShenFunc(symtl), V2550)

											gen12256 := Call(__e, ShenFunc(symshen_4lazyderef), gen12255, V2791)

											V2556 := gen12256
											gen12577 := Call(__e, ShenFunc(symcons_2), V2556)

											if True == gen12577 {
												gen12257 := Call(__e, ShenFunc(symhd), V2556)

												gen12258 := Call(__e, ShenFunc(symshen_4lazyderef), gen12257, V2791)

												V2557 := gen12258
												gen12576 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2557)

												if True == gen12576 {
													gen12259 := Call(__e, ShenFunc(symtl), V2556)

													gen12260 := Call(__e, ShenFunc(symshen_4lazyderef), gen12259, V2791)

													V2558 := gen12260
													gen12575 := Call(__e, ShenFunc(symcons_2), V2558)

													if True == gen12575 {
														gen12261 := Call(__e, ShenFunc(symhd), V2558)

														gen12262 := Call(__e, ShenFunc(symshen_4lazyderef), gen12261, V2791)

														V2559 := gen12262
														gen12574 := Call(__e, ShenFunc(symcons_2), V2559)

														if True == gen12574 {
															gen12303 := Call(__e, ShenFunc(symhd), V2559)

															A := gen12303
															gen12304 := Call(__e, ShenFunc(symtl), V2559)

															gen12305 := Call(__e, ShenFunc(symshen_4lazyderef), gen12304, V2791)

															V2560 := gen12305
															gen12573 := Call(__e, ShenFunc(symcons_2), V2560)

															if True == gen12573 {
																gen12344 := Call(__e, ShenFunc(symhd), V2560)

																gen12345 := Call(__e, ShenFunc(symshen_4lazyderef), gen12344, V2791)

																V2561 := gen12345
																gen12572 := Call(__e, ShenFunc(sym_a), MakeSymbol("*"), V2561)

																if True == gen12572 {
																	gen12460 := Call(__e, ShenFunc(symtl), V2560)

																	gen12461 := Call(__e, ShenFunc(symshen_4lazyderef), gen12460, V2791)

																	V2562 := gen12461
																	gen12571 := Call(__e, ShenFunc(symcons_2), V2562)

																	if True == gen12571 {
																		gen12499 := Call(__e, ShenFunc(symhd), V2562)

																		B := gen12499
																		gen12500 := Call(__e, ShenFunc(symtl), V2562)

																		gen12501 := Call(__e, ShenFunc(symshen_4lazyderef), gen12500, V2791)

																		V2563 := gen12501
																		gen12570 := Call(__e, ShenFunc(sym_a), Nil, V2563)

																		if True == gen12570 {
																			gen12537 := Call(__e, ShenFunc(symtl), V2558)

																			gen12538 := Call(__e, ShenFunc(symshen_4lazyderef), gen12537, V2791)

																			V2564 := gen12538
																			gen12569 := Call(__e, ShenFunc(sym_a), Nil, V2564)

																			if True == gen12569 {
																				gen12555 := Call(__e, ShenFunc(symtl), V2549)

																				Hyp := gen12555
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12556 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12557 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12558 := Call(__e, ShenFunc(symcons), gen12557, Nil)

																				gen12559 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12558)

																				gen12560 := Call(__e, ShenFunc(symcons), gen12556, gen12559)

																				gen12561 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12562 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																				gen12563 := Call(__e, ShenFunc(symcons), gen12562, Nil)

																				gen12564 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12563)

																				gen12565 := Call(__e, ShenFunc(symcons), gen12561, gen12564)

																				gen12566 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12567 := Call(__e, ShenFunc(symcons), gen12565, gen12566)

																				gen12568 := Call(__e, ShenFunc(symcons), gen12560, gen12567)

																				gen12585 = Call(__e, ShenFunc(symbind), V2790, gen12568, V2791, V2792)

																			} else {
																				gen12554 := Call(__e, ShenFunc(symshen_4pvar_2), V2564)

																				if True == gen12554 {
																					Call(__e, ShenFunc(symshen_4bindv), V2564, Nil, V2791)
																					gen12539 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen12539
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12540 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12541 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12542 := Call(__e, ShenFunc(symcons), gen12541, Nil)

																					gen12543 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12542)

																					gen12544 := Call(__e, ShenFunc(symcons), gen12540, gen12543)

																					gen12545 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12546 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen12547 := Call(__e, ShenFunc(symcons), gen12546, Nil)

																					gen12548 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12547)

																					gen12549 := Call(__e, ShenFunc(symcons), gen12545, gen12548)

																					gen12550 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12551 := Call(__e, ShenFunc(symcons), gen12549, gen12550)

																					gen12552 := Call(__e, ShenFunc(symcons), gen12544, gen12551)

																					gen12553 := Call(__e, ShenFunc(symbind), V2790, gen12552, V2791, V2792)

																					Result := gen12553
																					Call(__e, ShenFunc(symshen_4unbindv), V2564, V2791)
																					gen12585 = Result

																				} else {
																					gen12585 = False
																				}

																			}

																		} else {
																			gen12536 := Call(__e, ShenFunc(symshen_4pvar_2), V2563)

																			if True == gen12536 {
																				Call(__e, ShenFunc(symshen_4bindv), V2563, Nil, V2791)
																				gen12502 := Call(__e, ShenFunc(symtl), V2558)

																				gen12503 := Call(__e, ShenFunc(symshen_4lazyderef), gen12502, V2791)

																				V2565 := gen12503
																				gen12534 := Call(__e, ShenFunc(sym_a), Nil, V2565)

																				var gen12535 Obj
																				if True == gen12534 {
																					gen12520 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen12520
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12521 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12522 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12523 := Call(__e, ShenFunc(symcons), gen12522, Nil)

																					gen12524 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12523)

																					gen12525 := Call(__e, ShenFunc(symcons), gen12521, gen12524)

																					gen12526 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12527 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen12528 := Call(__e, ShenFunc(symcons), gen12527, Nil)

																					gen12529 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12528)

																					gen12530 := Call(__e, ShenFunc(symcons), gen12526, gen12529)

																					gen12531 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12532 := Call(__e, ShenFunc(symcons), gen12530, gen12531)

																					gen12533 := Call(__e, ShenFunc(symcons), gen12525, gen12532)

																					gen12535 = Call(__e, ShenFunc(symbind), V2790, gen12533, V2791, V2792)

																				} else {
																					gen12519 := Call(__e, ShenFunc(symshen_4pvar_2), V2565)

																					if True == gen12519 {
																						Call(__e, ShenFunc(symshen_4bindv), V2565, Nil, V2791)
																						gen12504 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen12504
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12505 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12506 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12507 := Call(__e, ShenFunc(symcons), gen12506, Nil)

																						gen12508 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12507)

																						gen12509 := Call(__e, ShenFunc(symcons), gen12505, gen12508)

																						gen12510 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12511 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen12512 := Call(__e, ShenFunc(symcons), gen12511, Nil)

																						gen12513 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12512)

																						gen12514 := Call(__e, ShenFunc(symcons), gen12510, gen12513)

																						gen12515 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12516 := Call(__e, ShenFunc(symcons), gen12514, gen12515)

																						gen12517 := Call(__e, ShenFunc(symcons), gen12509, gen12516)

																						gen12518 := Call(__e, ShenFunc(symbind), V2790, gen12517, V2791, V2792)

																						Result := gen12518
																						Call(__e, ShenFunc(symshen_4unbindv), V2565, V2791)
																						gen12535 = Result

																					} else {
																						gen12535 = False
																					}

																				}
																				Result := gen12535
																				Call(__e, ShenFunc(symshen_4unbindv), V2563, V2791)
																				gen12585 = Result

																			} else {
																				gen12585 = False
																			}

																		}

																	} else {
																		gen12498 := Call(__e, ShenFunc(symshen_4pvar_2), V2562)

																		if True == gen12498 {
																			gen12462 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																			B := gen12462
																			gen12463 := Call(__e, ShenFunc(symcons), B, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2562, gen12463, V2791)

																			gen12464 := Call(__e, ShenFunc(symtl), V2558)

																			gen12465 := Call(__e, ShenFunc(symshen_4lazyderef), gen12464, V2791)

																			V2566 := gen12465
																			gen12496 := Call(__e, ShenFunc(sym_a), Nil, V2566)

																			var gen12497 Obj
																			if True == gen12496 {
																				gen12482 := Call(__e, ShenFunc(symtl), V2549)

																				Hyp := gen12482
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12483 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12484 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12485 := Call(__e, ShenFunc(symcons), gen12484, Nil)

																				gen12486 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12485)

																				gen12487 := Call(__e, ShenFunc(symcons), gen12483, gen12486)

																				gen12488 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12489 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																				gen12490 := Call(__e, ShenFunc(symcons), gen12489, Nil)

																				gen12491 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12490)

																				gen12492 := Call(__e, ShenFunc(symcons), gen12488, gen12491)

																				gen12493 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12494 := Call(__e, ShenFunc(symcons), gen12492, gen12493)

																				gen12495 := Call(__e, ShenFunc(symcons), gen12487, gen12494)

																				gen12497 = Call(__e, ShenFunc(symbind), V2790, gen12495, V2791, V2792)

																			} else {
																				gen12481 := Call(__e, ShenFunc(symshen_4pvar_2), V2566)

																				if True == gen12481 {
																					Call(__e, ShenFunc(symshen_4bindv), V2566, Nil, V2791)
																					gen12466 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen12466
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12467 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12468 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12469 := Call(__e, ShenFunc(symcons), gen12468, Nil)

																					gen12470 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12469)

																					gen12471 := Call(__e, ShenFunc(symcons), gen12467, gen12470)

																					gen12472 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12473 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen12474 := Call(__e, ShenFunc(symcons), gen12473, Nil)

																					gen12475 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12474)

																					gen12476 := Call(__e, ShenFunc(symcons), gen12472, gen12475)

																					gen12477 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12478 := Call(__e, ShenFunc(symcons), gen12476, gen12477)

																					gen12479 := Call(__e, ShenFunc(symcons), gen12471, gen12478)

																					gen12480 := Call(__e, ShenFunc(symbind), V2790, gen12479, V2791, V2792)

																					Result := gen12480
																					Call(__e, ShenFunc(symshen_4unbindv), V2566, V2791)
																					gen12497 = Result

																				} else {
																					gen12497 = False
																				}

																			}
																			Result := gen12497
																			Call(__e, ShenFunc(symshen_4unbindv), V2562, V2791)
																			gen12585 = Result

																		} else {
																			gen12585 = False
																		}

																	}

																} else {
																	gen12459 := Call(__e, ShenFunc(symshen_4pvar_2), V2561)

																	if True == gen12459 {
																		Call(__e, ShenFunc(symshen_4bindv), V2561, MakeSymbol("*"), V2791)
																		gen12346 := Call(__e, ShenFunc(symtl), V2560)

																		gen12347 := Call(__e, ShenFunc(symshen_4lazyderef), gen12346, V2791)

																		V2567 := gen12347
																		gen12457 := Call(__e, ShenFunc(symcons_2), V2567)

																		var gen12458 Obj
																		if True == gen12457 {
																			gen12385 := Call(__e, ShenFunc(symhd), V2567)

																			B := gen12385
																			gen12386 := Call(__e, ShenFunc(symtl), V2567)

																			gen12387 := Call(__e, ShenFunc(symshen_4lazyderef), gen12386, V2791)

																			V2568 := gen12387
																			gen12456 := Call(__e, ShenFunc(sym_a), Nil, V2568)

																			if True == gen12456 {
																				gen12423 := Call(__e, ShenFunc(symtl), V2558)

																				gen12424 := Call(__e, ShenFunc(symshen_4lazyderef), gen12423, V2791)

																				V2569 := gen12424
																				gen12455 := Call(__e, ShenFunc(sym_a), Nil, V2569)

																				if True == gen12455 {
																					gen12441 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen12441
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12442 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12443 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12444 := Call(__e, ShenFunc(symcons), gen12443, Nil)

																					gen12445 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12444)

																					gen12446 := Call(__e, ShenFunc(symcons), gen12442, gen12445)

																					gen12447 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12448 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen12449 := Call(__e, ShenFunc(symcons), gen12448, Nil)

																					gen12450 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12449)

																					gen12451 := Call(__e, ShenFunc(symcons), gen12447, gen12450)

																					gen12452 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12453 := Call(__e, ShenFunc(symcons), gen12451, gen12452)

																					gen12454 := Call(__e, ShenFunc(symcons), gen12446, gen12453)

																					gen12458 = Call(__e, ShenFunc(symbind), V2790, gen12454, V2791, V2792)

																				} else {
																					gen12440 := Call(__e, ShenFunc(symshen_4pvar_2), V2569)

																					if True == gen12440 {
																						Call(__e, ShenFunc(symshen_4bindv), V2569, Nil, V2791)
																						gen12425 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen12425
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12426 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12427 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12428 := Call(__e, ShenFunc(symcons), gen12427, Nil)

																						gen12429 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12428)

																						gen12430 := Call(__e, ShenFunc(symcons), gen12426, gen12429)

																						gen12431 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12432 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen12433 := Call(__e, ShenFunc(symcons), gen12432, Nil)

																						gen12434 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12433)

																						gen12435 := Call(__e, ShenFunc(symcons), gen12431, gen12434)

																						gen12436 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12437 := Call(__e, ShenFunc(symcons), gen12435, gen12436)

																						gen12438 := Call(__e, ShenFunc(symcons), gen12430, gen12437)

																						gen12439 := Call(__e, ShenFunc(symbind), V2790, gen12438, V2791, V2792)

																						Result := gen12439
																						Call(__e, ShenFunc(symshen_4unbindv), V2569, V2791)
																						gen12458 = Result

																					} else {
																						gen12458 = False
																					}

																				}

																			} else {
																				gen12422 := Call(__e, ShenFunc(symshen_4pvar_2), V2568)

																				if True == gen12422 {
																					Call(__e, ShenFunc(symshen_4bindv), V2568, Nil, V2791)
																					gen12388 := Call(__e, ShenFunc(symtl), V2558)

																					gen12389 := Call(__e, ShenFunc(symshen_4lazyderef), gen12388, V2791)

																					V2570 := gen12389
																					gen12420 := Call(__e, ShenFunc(sym_a), Nil, V2570)

																					var gen12421 Obj
																					if True == gen12420 {
																						gen12406 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen12406
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12407 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12408 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12409 := Call(__e, ShenFunc(symcons), gen12408, Nil)

																						gen12410 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12409)

																						gen12411 := Call(__e, ShenFunc(symcons), gen12407, gen12410)

																						gen12412 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12413 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen12414 := Call(__e, ShenFunc(symcons), gen12413, Nil)

																						gen12415 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12414)

																						gen12416 := Call(__e, ShenFunc(symcons), gen12412, gen12415)

																						gen12417 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12418 := Call(__e, ShenFunc(symcons), gen12416, gen12417)

																						gen12419 := Call(__e, ShenFunc(symcons), gen12411, gen12418)

																						gen12421 = Call(__e, ShenFunc(symbind), V2790, gen12419, V2791, V2792)

																					} else {
																						gen12405 := Call(__e, ShenFunc(symshen_4pvar_2), V2570)

																						if True == gen12405 {
																							Call(__e, ShenFunc(symshen_4bindv), V2570, Nil, V2791)
																							gen12390 := Call(__e, ShenFunc(symtl), V2549)

																							Hyp := gen12390
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen12391 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																							gen12392 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen12393 := Call(__e, ShenFunc(symcons), gen12392, Nil)

																							gen12394 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12393)

																							gen12395 := Call(__e, ShenFunc(symcons), gen12391, gen12394)

																							gen12396 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																							gen12397 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																							gen12398 := Call(__e, ShenFunc(symcons), gen12397, Nil)

																							gen12399 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12398)

																							gen12400 := Call(__e, ShenFunc(symcons), gen12396, gen12399)

																							gen12401 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																							gen12402 := Call(__e, ShenFunc(symcons), gen12400, gen12401)

																							gen12403 := Call(__e, ShenFunc(symcons), gen12395, gen12402)

																							gen12404 := Call(__e, ShenFunc(symbind), V2790, gen12403, V2791, V2792)

																							Result := gen12404
																							Call(__e, ShenFunc(symshen_4unbindv), V2570, V2791)
																							gen12421 = Result

																						} else {
																							gen12421 = False
																						}

																					}
																					Result := gen12421
																					Call(__e, ShenFunc(symshen_4unbindv), V2568, V2791)
																					gen12458 = Result

																				} else {
																					gen12458 = False
																				}

																			}

																		} else {
																			gen12384 := Call(__e, ShenFunc(symshen_4pvar_2), V2567)

																			if True == gen12384 {
																				gen12348 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																				B := gen12348
																				gen12349 := Call(__e, ShenFunc(symcons), B, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2567, gen12349, V2791)

																				gen12350 := Call(__e, ShenFunc(symtl), V2558)

																				gen12351 := Call(__e, ShenFunc(symshen_4lazyderef), gen12350, V2791)

																				V2571 := gen12351
																				gen12382 := Call(__e, ShenFunc(sym_a), Nil, V2571)

																				var gen12383 Obj
																				if True == gen12382 {
																					gen12368 := Call(__e, ShenFunc(symtl), V2549)

																					Hyp := gen12368
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12369 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12370 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12371 := Call(__e, ShenFunc(symcons), gen12370, Nil)

																					gen12372 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12371)

																					gen12373 := Call(__e, ShenFunc(symcons), gen12369, gen12372)

																					gen12374 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12375 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																					gen12376 := Call(__e, ShenFunc(symcons), gen12375, Nil)

																					gen12377 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12376)

																					gen12378 := Call(__e, ShenFunc(symcons), gen12374, gen12377)

																					gen12379 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12380 := Call(__e, ShenFunc(symcons), gen12378, gen12379)

																					gen12381 := Call(__e, ShenFunc(symcons), gen12373, gen12380)

																					gen12383 = Call(__e, ShenFunc(symbind), V2790, gen12381, V2791, V2792)

																				} else {
																					gen12367 := Call(__e, ShenFunc(symshen_4pvar_2), V2571)

																					if True == gen12367 {
																						Call(__e, ShenFunc(symshen_4bindv), V2571, Nil, V2791)
																						gen12352 := Call(__e, ShenFunc(symtl), V2549)

																						Hyp := gen12352
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12353 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12354 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12355 := Call(__e, ShenFunc(symcons), gen12354, Nil)

																						gen12356 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12355)

																						gen12357 := Call(__e, ShenFunc(symcons), gen12353, gen12356)

																						gen12358 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12359 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																						gen12360 := Call(__e, ShenFunc(symcons), gen12359, Nil)

																						gen12361 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12360)

																						gen12362 := Call(__e, ShenFunc(symcons), gen12358, gen12361)

																						gen12363 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12364 := Call(__e, ShenFunc(symcons), gen12362, gen12363)

																						gen12365 := Call(__e, ShenFunc(symcons), gen12357, gen12364)

																						gen12366 := Call(__e, ShenFunc(symbind), V2790, gen12365, V2791, V2792)

																						Result := gen12366
																						Call(__e, ShenFunc(symshen_4unbindv), V2571, V2791)
																						gen12383 = Result

																					} else {
																						gen12383 = False
																					}

																				}
																				Result := gen12383
																				Call(__e, ShenFunc(symshen_4unbindv), V2567, V2791)
																				gen12458 = Result

																			} else {
																				gen12458 = False
																			}

																		}
																		Result := gen12458
																		Call(__e, ShenFunc(symshen_4unbindv), V2561, V2791)
																		gen12585 = Result

																	} else {
																		gen12585 = False
																	}

																}

															} else {
																gen12343 := Call(__e, ShenFunc(symshen_4pvar_2), V2560)

																if True == gen12343 {
																	gen12306 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	B := gen12306
																	gen12307 := Call(__e, ShenFunc(symcons), B, Nil)

																	gen12308 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen12307)

																	Call(__e, ShenFunc(symshen_4bindv), V2560, gen12308, V2791)

																	gen12309 := Call(__e, ShenFunc(symtl), V2558)

																	gen12310 := Call(__e, ShenFunc(symshen_4lazyderef), gen12309, V2791)

																	V2572 := gen12310
																	gen12341 := Call(__e, ShenFunc(sym_a), Nil, V2572)

																	var gen12342 Obj
																	if True == gen12341 {
																		gen12327 := Call(__e, ShenFunc(symtl), V2549)

																		Hyp := gen12327
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12328 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12329 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12330 := Call(__e, ShenFunc(symcons), gen12329, Nil)

																		gen12331 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12330)

																		gen12332 := Call(__e, ShenFunc(symcons), gen12328, gen12331)

																		gen12333 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12334 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																		gen12335 := Call(__e, ShenFunc(symcons), gen12334, Nil)

																		gen12336 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12335)

																		gen12337 := Call(__e, ShenFunc(symcons), gen12333, gen12336)

																		gen12338 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12339 := Call(__e, ShenFunc(symcons), gen12337, gen12338)

																		gen12340 := Call(__e, ShenFunc(symcons), gen12332, gen12339)

																		gen12342 = Call(__e, ShenFunc(symbind), V2790, gen12340, V2791, V2792)

																	} else {
																		gen12326 := Call(__e, ShenFunc(symshen_4pvar_2), V2572)

																		if True == gen12326 {
																			Call(__e, ShenFunc(symshen_4bindv), V2572, Nil, V2791)
																			gen12311 := Call(__e, ShenFunc(symtl), V2549)

																			Hyp := gen12311
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12312 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12313 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12314 := Call(__e, ShenFunc(symcons), gen12313, Nil)

																			gen12315 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12314)

																			gen12316 := Call(__e, ShenFunc(symcons), gen12312, gen12315)

																			gen12317 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12318 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																			gen12319 := Call(__e, ShenFunc(symcons), gen12318, Nil)

																			gen12320 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12319)

																			gen12321 := Call(__e, ShenFunc(symcons), gen12317, gen12320)

																			gen12322 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12323 := Call(__e, ShenFunc(symcons), gen12321, gen12322)

																			gen12324 := Call(__e, ShenFunc(symcons), gen12316, gen12323)

																			gen12325 := Call(__e, ShenFunc(symbind), V2790, gen12324, V2791, V2792)

																			Result := gen12325
																			Call(__e, ShenFunc(symshen_4unbindv), V2572, V2791)
																			gen12342 = Result

																		} else {
																			gen12342 = False
																		}

																	}
																	Result := gen12342
																	Call(__e, ShenFunc(symshen_4unbindv), V2560, V2791)
																	gen12585 = Result

																} else {
																	gen12585 = False
																}

															}

														} else {
															gen12302 := Call(__e, ShenFunc(symshen_4pvar_2), V2559)

															if True == gen12302 {
																gen12263 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																A := gen12263
																gen12264 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																B := gen12264
																gen12265 := Call(__e, ShenFunc(symcons), B, Nil)

																gen12266 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen12265)

																gen12267 := Call(__e, ShenFunc(symcons), A, gen12266)

																Call(__e, ShenFunc(symshen_4bindv), V2559, gen12267, V2791)

																gen12268 := Call(__e, ShenFunc(symtl), V2558)

																gen12269 := Call(__e, ShenFunc(symshen_4lazyderef), gen12268, V2791)

																V2573 := gen12269
																gen12300 := Call(__e, ShenFunc(sym_a), Nil, V2573)

																var gen12301 Obj
																if True == gen12300 {
																	gen12286 := Call(__e, ShenFunc(symtl), V2549)

																	Hyp := gen12286
																	Call(__e, ShenFunc(symshen_4incinfs))
																	gen12287 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																	gen12288 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																	gen12289 := Call(__e, ShenFunc(symcons), gen12288, Nil)

																	gen12290 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12289)

																	gen12291 := Call(__e, ShenFunc(symcons), gen12287, gen12290)

																	gen12292 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																	gen12293 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																	gen12294 := Call(__e, ShenFunc(symcons), gen12293, Nil)

																	gen12295 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12294)

																	gen12296 := Call(__e, ShenFunc(symcons), gen12292, gen12295)

																	gen12297 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																	gen12298 := Call(__e, ShenFunc(symcons), gen12296, gen12297)

																	gen12299 := Call(__e, ShenFunc(symcons), gen12291, gen12298)

																	gen12301 = Call(__e, ShenFunc(symbind), V2790, gen12299, V2791, V2792)

																} else {
																	gen12285 := Call(__e, ShenFunc(symshen_4pvar_2), V2573)

																	if True == gen12285 {
																		Call(__e, ShenFunc(symshen_4bindv), V2573, Nil, V2791)
																		gen12270 := Call(__e, ShenFunc(symtl), V2549)

																		Hyp := gen12270
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12271 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12272 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12273 := Call(__e, ShenFunc(symcons), gen12272, Nil)

																		gen12274 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12273)

																		gen12275 := Call(__e, ShenFunc(symcons), gen12271, gen12274)

																		gen12276 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12277 := Call(__e, ShenFunc(symshen_4lazyderef), B, V2791)

																		gen12278 := Call(__e, ShenFunc(symcons), gen12277, Nil)

																		gen12279 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12278)

																		gen12280 := Call(__e, ShenFunc(symcons), gen12276, gen12279)

																		gen12281 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12282 := Call(__e, ShenFunc(symcons), gen12280, gen12281)

																		gen12283 := Call(__e, ShenFunc(symcons), gen12275, gen12282)

																		gen12284 := Call(__e, ShenFunc(symbind), V2790, gen12283, V2791, V2792)

																		Result := gen12284
																		Call(__e, ShenFunc(symshen_4unbindv), V2573, V2791)
																		gen12301 = Result

																	} else {
																		gen12301 = False
																	}

																}
																Result := gen12301
																Call(__e, ShenFunc(symshen_4unbindv), V2559, V2791)
																gen12585 = Result

															} else {
																gen12585 = False
															}

														}

													} else {
														gen12585 = False
													}

												} else {
													gen12585 = False
												}

											} else {
												gen12585 = False
											}

										} else {
											gen12585 = False
										}

									} else {
										gen12585 = False
									}

								} else {
									gen12585 = False
								}

							} else {
								gen12585 = False
							}

						} else {
							gen12585 = False
						}

					} else {
						gen12585 = False
					}

				} else {
					gen12585 = False
				}
				Case := gen12585
				gen13022 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen13022 {
					gen12586 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

					V2574 := gen12586
					gen12914 := Call(__e, ShenFunc(symcons_2), V2574)

					var gen12915 Obj
					if True == gen12914 {
						gen12587 := Call(__e, ShenFunc(symhd), V2574)

						gen12588 := Call(__e, ShenFunc(symshen_4lazyderef), gen12587, V2791)

						V2575 := gen12588
						gen12913 := Call(__e, ShenFunc(symcons_2), V2575)

						if True == gen12913 {
							gen12589 := Call(__e, ShenFunc(symhd), V2575)

							gen12590 := Call(__e, ShenFunc(symshen_4lazyderef), gen12589, V2791)

							V2576 := gen12590
							gen12912 := Call(__e, ShenFunc(symcons_2), V2576)

							if True == gen12912 {
								gen12591 := Call(__e, ShenFunc(symhd), V2576)

								gen12592 := Call(__e, ShenFunc(symshen_4lazyderef), gen12591, V2791)

								V2577 := gen12592
								gen12911 := Call(__e, ShenFunc(sym_a), MakeSymbol("@v"), V2577)

								if True == gen12911 {
									gen12593 := Call(__e, ShenFunc(symtl), V2576)

									gen12594 := Call(__e, ShenFunc(symshen_4lazyderef), gen12593, V2791)

									V2578 := gen12594
									gen12910 := Call(__e, ShenFunc(symcons_2), V2578)

									if True == gen12910 {
										gen12595 := Call(__e, ShenFunc(symhd), V2578)

										X := gen12595
										gen12596 := Call(__e, ShenFunc(symtl), V2578)

										gen12597 := Call(__e, ShenFunc(symshen_4lazyderef), gen12596, V2791)

										V2579 := gen12597
										gen12909 := Call(__e, ShenFunc(symcons_2), V2579)

										if True == gen12909 {
											gen12598 := Call(__e, ShenFunc(symhd), V2579)

											Y := gen12598
											gen12599 := Call(__e, ShenFunc(symtl), V2579)

											gen12600 := Call(__e, ShenFunc(symshen_4lazyderef), gen12599, V2791)

											V2580 := gen12600
											gen12908 := Call(__e, ShenFunc(sym_a), Nil, V2580)

											if True == gen12908 {
												gen12601 := Call(__e, ShenFunc(symtl), V2575)

												gen12602 := Call(__e, ShenFunc(symshen_4lazyderef), gen12601, V2791)

												V2581 := gen12602
												gen12907 := Call(__e, ShenFunc(symcons_2), V2581)

												if True == gen12907 {
													gen12603 := Call(__e, ShenFunc(symhd), V2581)

													gen12604 := Call(__e, ShenFunc(symshen_4lazyderef), gen12603, V2791)

													V2582 := gen12604
													gen12906 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2582)

													if True == gen12906 {
														gen12605 := Call(__e, ShenFunc(symtl), V2581)

														gen12606 := Call(__e, ShenFunc(symshen_4lazyderef), gen12605, V2791)

														V2583 := gen12606
														gen12905 := Call(__e, ShenFunc(symcons_2), V2583)

														if True == gen12905 {
															gen12607 := Call(__e, ShenFunc(symhd), V2583)

															gen12608 := Call(__e, ShenFunc(symshen_4lazyderef), gen12607, V2791)

															V2584 := gen12608
															gen12904 := Call(__e, ShenFunc(symcons_2), V2584)

															if True == gen12904 {
																gen12651 := Call(__e, ShenFunc(symhd), V2584)

																gen12652 := Call(__e, ShenFunc(symshen_4lazyderef), gen12651, V2791)

																V2585 := gen12652
																gen12903 := Call(__e, ShenFunc(sym_a), MakeSymbol("vector"), V2585)

																if True == gen12903 {
																	gen12779 := Call(__e, ShenFunc(symtl), V2584)

																	gen12780 := Call(__e, ShenFunc(symshen_4lazyderef), gen12779, V2791)

																	V2586 := gen12780
																	gen12902 := Call(__e, ShenFunc(symcons_2), V2586)

																	if True == gen12902 {
																		gen12822 := Call(__e, ShenFunc(symhd), V2586)

																		A := gen12822
																		gen12823 := Call(__e, ShenFunc(symtl), V2586)

																		gen12824 := Call(__e, ShenFunc(symshen_4lazyderef), gen12823, V2791)

																		V2587 := gen12824
																		gen12901 := Call(__e, ShenFunc(sym_a), Nil, V2587)

																		if True == gen12901 {
																			gen12864 := Call(__e, ShenFunc(symtl), V2583)

																			gen12865 := Call(__e, ShenFunc(symshen_4lazyderef), gen12864, V2791)

																			V2588 := gen12865
																			gen12900 := Call(__e, ShenFunc(sym_a), Nil, V2588)

																			if True == gen12900 {
																				gen12884 := Call(__e, ShenFunc(symtl), V2574)

																				Hyp := gen12884
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12885 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12886 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12887 := Call(__e, ShenFunc(symcons), gen12886, Nil)

																				gen12888 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12887)

																				gen12889 := Call(__e, ShenFunc(symcons), gen12885, gen12888)

																				gen12890 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12891 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12892 := Call(__e, ShenFunc(symcons), gen12891, Nil)

																				gen12893 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12892)

																				gen12894 := Call(__e, ShenFunc(symcons), gen12893, Nil)

																				gen12895 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12894)

																				gen12896 := Call(__e, ShenFunc(symcons), gen12890, gen12895)

																				gen12897 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12898 := Call(__e, ShenFunc(symcons), gen12896, gen12897)

																				gen12899 := Call(__e, ShenFunc(symcons), gen12889, gen12898)

																				gen12915 = Call(__e, ShenFunc(symbind), V2790, gen12899, V2791, V2792)

																			} else {
																				gen12883 := Call(__e, ShenFunc(symshen_4pvar_2), V2588)

																				if True == gen12883 {
																					Call(__e, ShenFunc(symshen_4bindv), V2588, Nil, V2791)
																					gen12866 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen12866
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12867 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12868 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12869 := Call(__e, ShenFunc(symcons), gen12868, Nil)

																					gen12870 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12869)

																					gen12871 := Call(__e, ShenFunc(symcons), gen12867, gen12870)

																					gen12872 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12873 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12874 := Call(__e, ShenFunc(symcons), gen12873, Nil)

																					gen12875 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12874)

																					gen12876 := Call(__e, ShenFunc(symcons), gen12875, Nil)

																					gen12877 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12876)

																					gen12878 := Call(__e, ShenFunc(symcons), gen12872, gen12877)

																					gen12879 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12880 := Call(__e, ShenFunc(symcons), gen12878, gen12879)

																					gen12881 := Call(__e, ShenFunc(symcons), gen12871, gen12880)

																					gen12882 := Call(__e, ShenFunc(symbind), V2790, gen12881, V2791, V2792)

																					Result := gen12882
																					Call(__e, ShenFunc(symshen_4unbindv), V2588, V2791)
																					gen12915 = Result

																				} else {
																					gen12915 = False
																				}

																			}

																		} else {
																			gen12863 := Call(__e, ShenFunc(symshen_4pvar_2), V2587)

																			if True == gen12863 {
																				Call(__e, ShenFunc(symshen_4bindv), V2587, Nil, V2791)
																				gen12825 := Call(__e, ShenFunc(symtl), V2583)

																				gen12826 := Call(__e, ShenFunc(symshen_4lazyderef), gen12825, V2791)

																				V2589 := gen12826
																				gen12861 := Call(__e, ShenFunc(sym_a), Nil, V2589)

																				var gen12862 Obj
																				if True == gen12861 {
																					gen12845 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen12845
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12846 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12847 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12848 := Call(__e, ShenFunc(symcons), gen12847, Nil)

																					gen12849 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12848)

																					gen12850 := Call(__e, ShenFunc(symcons), gen12846, gen12849)

																					gen12851 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12852 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12853 := Call(__e, ShenFunc(symcons), gen12852, Nil)

																					gen12854 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12853)

																					gen12855 := Call(__e, ShenFunc(symcons), gen12854, Nil)

																					gen12856 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12855)

																					gen12857 := Call(__e, ShenFunc(symcons), gen12851, gen12856)

																					gen12858 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12859 := Call(__e, ShenFunc(symcons), gen12857, gen12858)

																					gen12860 := Call(__e, ShenFunc(symcons), gen12850, gen12859)

																					gen12862 = Call(__e, ShenFunc(symbind), V2790, gen12860, V2791, V2792)

																				} else {
																					gen12844 := Call(__e, ShenFunc(symshen_4pvar_2), V2589)

																					if True == gen12844 {
																						Call(__e, ShenFunc(symshen_4bindv), V2589, Nil, V2791)
																						gen12827 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen12827
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12828 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12829 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12830 := Call(__e, ShenFunc(symcons), gen12829, Nil)

																						gen12831 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12830)

																						gen12832 := Call(__e, ShenFunc(symcons), gen12828, gen12831)

																						gen12833 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12834 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12835 := Call(__e, ShenFunc(symcons), gen12834, Nil)

																						gen12836 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12835)

																						gen12837 := Call(__e, ShenFunc(symcons), gen12836, Nil)

																						gen12838 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12837)

																						gen12839 := Call(__e, ShenFunc(symcons), gen12833, gen12838)

																						gen12840 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12841 := Call(__e, ShenFunc(symcons), gen12839, gen12840)

																						gen12842 := Call(__e, ShenFunc(symcons), gen12832, gen12841)

																						gen12843 := Call(__e, ShenFunc(symbind), V2790, gen12842, V2791, V2792)

																						Result := gen12843
																						Call(__e, ShenFunc(symshen_4unbindv), V2589, V2791)
																						gen12862 = Result

																					} else {
																						gen12862 = False
																					}

																				}
																				Result := gen12862
																				Call(__e, ShenFunc(symshen_4unbindv), V2587, V2791)
																				gen12915 = Result

																			} else {
																				gen12915 = False
																			}

																		}

																	} else {
																		gen12821 := Call(__e, ShenFunc(symshen_4pvar_2), V2586)

																		if True == gen12821 {
																			gen12781 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																			A := gen12781
																			gen12782 := Call(__e, ShenFunc(symcons), A, Nil)

																			Call(__e, ShenFunc(symshen_4bindv), V2586, gen12782, V2791)

																			gen12783 := Call(__e, ShenFunc(symtl), V2583)

																			gen12784 := Call(__e, ShenFunc(symshen_4lazyderef), gen12783, V2791)

																			V2590 := gen12784
																			gen12819 := Call(__e, ShenFunc(sym_a), Nil, V2590)

																			var gen12820 Obj
																			if True == gen12819 {
																				gen12803 := Call(__e, ShenFunc(symtl), V2574)

																				Hyp := gen12803
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12804 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12805 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12806 := Call(__e, ShenFunc(symcons), gen12805, Nil)

																				gen12807 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12806)

																				gen12808 := Call(__e, ShenFunc(symcons), gen12804, gen12807)

																				gen12809 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12810 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																				gen12811 := Call(__e, ShenFunc(symcons), gen12810, Nil)

																				gen12812 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12811)

																				gen12813 := Call(__e, ShenFunc(symcons), gen12812, Nil)

																				gen12814 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12813)

																				gen12815 := Call(__e, ShenFunc(symcons), gen12809, gen12814)

																				gen12816 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12817 := Call(__e, ShenFunc(symcons), gen12815, gen12816)

																				gen12818 := Call(__e, ShenFunc(symcons), gen12808, gen12817)

																				gen12820 = Call(__e, ShenFunc(symbind), V2790, gen12818, V2791, V2792)

																			} else {
																				gen12802 := Call(__e, ShenFunc(symshen_4pvar_2), V2590)

																				if True == gen12802 {
																					Call(__e, ShenFunc(symshen_4bindv), V2590, Nil, V2791)
																					gen12785 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen12785
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12786 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12787 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12788 := Call(__e, ShenFunc(symcons), gen12787, Nil)

																					gen12789 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12788)

																					gen12790 := Call(__e, ShenFunc(symcons), gen12786, gen12789)

																					gen12791 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12792 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12793 := Call(__e, ShenFunc(symcons), gen12792, Nil)

																					gen12794 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12793)

																					gen12795 := Call(__e, ShenFunc(symcons), gen12794, Nil)

																					gen12796 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12795)

																					gen12797 := Call(__e, ShenFunc(symcons), gen12791, gen12796)

																					gen12798 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12799 := Call(__e, ShenFunc(symcons), gen12797, gen12798)

																					gen12800 := Call(__e, ShenFunc(symcons), gen12790, gen12799)

																					gen12801 := Call(__e, ShenFunc(symbind), V2790, gen12800, V2791, V2792)

																					Result := gen12801
																					Call(__e, ShenFunc(symshen_4unbindv), V2590, V2791)
																					gen12820 = Result

																				} else {
																					gen12820 = False
																				}

																			}
																			Result := gen12820
																			Call(__e, ShenFunc(symshen_4unbindv), V2586, V2791)
																			gen12915 = Result

																		} else {
																			gen12915 = False
																		}

																	}

																} else {
																	gen12778 := Call(__e, ShenFunc(symshen_4pvar_2), V2585)

																	if True == gen12778 {
																		Call(__e, ShenFunc(symshen_4bindv), V2585, MakeSymbol("vector"), V2791)
																		gen12653 := Call(__e, ShenFunc(symtl), V2584)

																		gen12654 := Call(__e, ShenFunc(symshen_4lazyderef), gen12653, V2791)

																		V2591 := gen12654
																		gen12776 := Call(__e, ShenFunc(symcons_2), V2591)

																		var gen12777 Obj
																		if True == gen12776 {
																			gen12696 := Call(__e, ShenFunc(symhd), V2591)

																			A := gen12696
																			gen12697 := Call(__e, ShenFunc(symtl), V2591)

																			gen12698 := Call(__e, ShenFunc(symshen_4lazyderef), gen12697, V2791)

																			V2592 := gen12698
																			gen12775 := Call(__e, ShenFunc(sym_a), Nil, V2592)

																			if True == gen12775 {
																				gen12738 := Call(__e, ShenFunc(symtl), V2583)

																				gen12739 := Call(__e, ShenFunc(symshen_4lazyderef), gen12738, V2791)

																				V2593 := gen12739
																				gen12774 := Call(__e, ShenFunc(sym_a), Nil, V2593)

																				if True == gen12774 {
																					gen12758 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen12758
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12759 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12760 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12761 := Call(__e, ShenFunc(symcons), gen12760, Nil)

																					gen12762 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12761)

																					gen12763 := Call(__e, ShenFunc(symcons), gen12759, gen12762)

																					gen12764 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12765 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12766 := Call(__e, ShenFunc(symcons), gen12765, Nil)

																					gen12767 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12766)

																					gen12768 := Call(__e, ShenFunc(symcons), gen12767, Nil)

																					gen12769 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12768)

																					gen12770 := Call(__e, ShenFunc(symcons), gen12764, gen12769)

																					gen12771 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12772 := Call(__e, ShenFunc(symcons), gen12770, gen12771)

																					gen12773 := Call(__e, ShenFunc(symcons), gen12763, gen12772)

																					gen12777 = Call(__e, ShenFunc(symbind), V2790, gen12773, V2791, V2792)

																				} else {
																					gen12757 := Call(__e, ShenFunc(symshen_4pvar_2), V2593)

																					if True == gen12757 {
																						Call(__e, ShenFunc(symshen_4bindv), V2593, Nil, V2791)
																						gen12740 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen12740
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12741 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12742 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12743 := Call(__e, ShenFunc(symcons), gen12742, Nil)

																						gen12744 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12743)

																						gen12745 := Call(__e, ShenFunc(symcons), gen12741, gen12744)

																						gen12746 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12747 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12748 := Call(__e, ShenFunc(symcons), gen12747, Nil)

																						gen12749 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12748)

																						gen12750 := Call(__e, ShenFunc(symcons), gen12749, Nil)

																						gen12751 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12750)

																						gen12752 := Call(__e, ShenFunc(symcons), gen12746, gen12751)

																						gen12753 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12754 := Call(__e, ShenFunc(symcons), gen12752, gen12753)

																						gen12755 := Call(__e, ShenFunc(symcons), gen12745, gen12754)

																						gen12756 := Call(__e, ShenFunc(symbind), V2790, gen12755, V2791, V2792)

																						Result := gen12756
																						Call(__e, ShenFunc(symshen_4unbindv), V2593, V2791)
																						gen12777 = Result

																					} else {
																						gen12777 = False
																					}

																				}

																			} else {
																				gen12737 := Call(__e, ShenFunc(symshen_4pvar_2), V2592)

																				if True == gen12737 {
																					Call(__e, ShenFunc(symshen_4bindv), V2592, Nil, V2791)
																					gen12699 := Call(__e, ShenFunc(symtl), V2583)

																					gen12700 := Call(__e, ShenFunc(symshen_4lazyderef), gen12699, V2791)

																					V2594 := gen12700
																					gen12735 := Call(__e, ShenFunc(sym_a), Nil, V2594)

																					var gen12736 Obj
																					if True == gen12735 {
																						gen12719 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen12719
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12720 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12721 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12722 := Call(__e, ShenFunc(symcons), gen12721, Nil)

																						gen12723 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12722)

																						gen12724 := Call(__e, ShenFunc(symcons), gen12720, gen12723)

																						gen12725 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12726 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12727 := Call(__e, ShenFunc(symcons), gen12726, Nil)

																						gen12728 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12727)

																						gen12729 := Call(__e, ShenFunc(symcons), gen12728, Nil)

																						gen12730 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12729)

																						gen12731 := Call(__e, ShenFunc(symcons), gen12725, gen12730)

																						gen12732 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12733 := Call(__e, ShenFunc(symcons), gen12731, gen12732)

																						gen12734 := Call(__e, ShenFunc(symcons), gen12724, gen12733)

																						gen12736 = Call(__e, ShenFunc(symbind), V2790, gen12734, V2791, V2792)

																					} else {
																						gen12718 := Call(__e, ShenFunc(symshen_4pvar_2), V2594)

																						if True == gen12718 {
																							Call(__e, ShenFunc(symshen_4bindv), V2594, Nil, V2791)
																							gen12701 := Call(__e, ShenFunc(symtl), V2574)

																							Hyp := gen12701
																							Call(__e, ShenFunc(symshen_4incinfs))
																							gen12702 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																							gen12703 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen12704 := Call(__e, ShenFunc(symcons), gen12703, Nil)

																							gen12705 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12704)

																							gen12706 := Call(__e, ShenFunc(symcons), gen12702, gen12705)

																							gen12707 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																							gen12708 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																							gen12709 := Call(__e, ShenFunc(symcons), gen12708, Nil)

																							gen12710 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12709)

																							gen12711 := Call(__e, ShenFunc(symcons), gen12710, Nil)

																							gen12712 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12711)

																							gen12713 := Call(__e, ShenFunc(symcons), gen12707, gen12712)

																							gen12714 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																							gen12715 := Call(__e, ShenFunc(symcons), gen12713, gen12714)

																							gen12716 := Call(__e, ShenFunc(symcons), gen12706, gen12715)

																							gen12717 := Call(__e, ShenFunc(symbind), V2790, gen12716, V2791, V2792)

																							Result := gen12717
																							Call(__e, ShenFunc(symshen_4unbindv), V2594, V2791)
																							gen12736 = Result

																						} else {
																							gen12736 = False
																						}

																					}
																					Result := gen12736
																					Call(__e, ShenFunc(symshen_4unbindv), V2592, V2791)
																					gen12777 = Result

																				} else {
																					gen12777 = False
																				}

																			}

																		} else {
																			gen12695 := Call(__e, ShenFunc(symshen_4pvar_2), V2591)

																			if True == gen12695 {
																				gen12655 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																				A := gen12655
																				gen12656 := Call(__e, ShenFunc(symcons), A, Nil)

																				Call(__e, ShenFunc(symshen_4bindv), V2591, gen12656, V2791)

																				gen12657 := Call(__e, ShenFunc(symtl), V2583)

																				gen12658 := Call(__e, ShenFunc(symshen_4lazyderef), gen12657, V2791)

																				V2595 := gen12658
																				gen12693 := Call(__e, ShenFunc(sym_a), Nil, V2595)

																				var gen12694 Obj
																				if True == gen12693 {
																					gen12677 := Call(__e, ShenFunc(symtl), V2574)

																					Hyp := gen12677
																					Call(__e, ShenFunc(symshen_4incinfs))
																					gen12678 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																					gen12679 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12680 := Call(__e, ShenFunc(symcons), gen12679, Nil)

																					gen12681 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12680)

																					gen12682 := Call(__e, ShenFunc(symcons), gen12678, gen12681)

																					gen12683 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																					gen12684 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																					gen12685 := Call(__e, ShenFunc(symcons), gen12684, Nil)

																					gen12686 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12685)

																					gen12687 := Call(__e, ShenFunc(symcons), gen12686, Nil)

																					gen12688 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12687)

																					gen12689 := Call(__e, ShenFunc(symcons), gen12683, gen12688)

																					gen12690 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																					gen12691 := Call(__e, ShenFunc(symcons), gen12689, gen12690)

																					gen12692 := Call(__e, ShenFunc(symcons), gen12682, gen12691)

																					gen12694 = Call(__e, ShenFunc(symbind), V2790, gen12692, V2791, V2792)

																				} else {
																					gen12676 := Call(__e, ShenFunc(symshen_4pvar_2), V2595)

																					if True == gen12676 {
																						Call(__e, ShenFunc(symshen_4bindv), V2595, Nil, V2791)
																						gen12659 := Call(__e, ShenFunc(symtl), V2574)

																						Hyp := gen12659
																						Call(__e, ShenFunc(symshen_4incinfs))
																						gen12660 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																						gen12661 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12662 := Call(__e, ShenFunc(symcons), gen12661, Nil)

																						gen12663 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12662)

																						gen12664 := Call(__e, ShenFunc(symcons), gen12660, gen12663)

																						gen12665 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																						gen12666 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																						gen12667 := Call(__e, ShenFunc(symcons), gen12666, Nil)

																						gen12668 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12667)

																						gen12669 := Call(__e, ShenFunc(symcons), gen12668, Nil)

																						gen12670 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12669)

																						gen12671 := Call(__e, ShenFunc(symcons), gen12665, gen12670)

																						gen12672 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																						gen12673 := Call(__e, ShenFunc(symcons), gen12671, gen12672)

																						gen12674 := Call(__e, ShenFunc(symcons), gen12664, gen12673)

																						gen12675 := Call(__e, ShenFunc(symbind), V2790, gen12674, V2791, V2792)

																						Result := gen12675
																						Call(__e, ShenFunc(symshen_4unbindv), V2595, V2791)
																						gen12694 = Result

																					} else {
																						gen12694 = False
																					}

																				}
																				Result := gen12694
																				Call(__e, ShenFunc(symshen_4unbindv), V2591, V2791)
																				gen12777 = Result

																			} else {
																				gen12777 = False
																			}

																		}
																		Result := gen12777
																		Call(__e, ShenFunc(symshen_4unbindv), V2585, V2791)
																		gen12915 = Result

																	} else {
																		gen12915 = False
																	}

																}

															} else {
																gen12650 := Call(__e, ShenFunc(symshen_4pvar_2), V2584)

																if True == gen12650 {
																	gen12609 := Call(__e, ShenFunc(symshen_4newpv), V2791)

																	A := gen12609
																	gen12610 := Call(__e, ShenFunc(symcons), A, Nil)

																	gen12611 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12610)

																	Call(__e, ShenFunc(symshen_4bindv), V2584, gen12611, V2791)

																	gen12612 := Call(__e, ShenFunc(symtl), V2583)

																	gen12613 := Call(__e, ShenFunc(symshen_4lazyderef), gen12612, V2791)

																	V2596 := gen12613
																	gen12648 := Call(__e, ShenFunc(sym_a), Nil, V2596)

																	var gen12649 Obj
																	if True == gen12648 {
																		gen12632 := Call(__e, ShenFunc(symtl), V2574)

																		Hyp := gen12632
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12633 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12634 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12635 := Call(__e, ShenFunc(symcons), gen12634, Nil)

																		gen12636 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12635)

																		gen12637 := Call(__e, ShenFunc(symcons), gen12633, gen12636)

																		gen12638 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12639 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																		gen12640 := Call(__e, ShenFunc(symcons), gen12639, Nil)

																		gen12641 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12640)

																		gen12642 := Call(__e, ShenFunc(symcons), gen12641, Nil)

																		gen12643 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12642)

																		gen12644 := Call(__e, ShenFunc(symcons), gen12638, gen12643)

																		gen12645 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12646 := Call(__e, ShenFunc(symcons), gen12644, gen12645)

																		gen12647 := Call(__e, ShenFunc(symcons), gen12637, gen12646)

																		gen12649 = Call(__e, ShenFunc(symbind), V2790, gen12647, V2791, V2792)

																	} else {
																		gen12631 := Call(__e, ShenFunc(symshen_4pvar_2), V2596)

																		if True == gen12631 {
																			Call(__e, ShenFunc(symshen_4bindv), V2596, Nil, V2791)
																			gen12614 := Call(__e, ShenFunc(symtl), V2574)

																			Hyp := gen12614
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12615 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12616 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12617 := Call(__e, ShenFunc(symcons), gen12616, Nil)

																			gen12618 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12617)

																			gen12619 := Call(__e, ShenFunc(symcons), gen12615, gen12618)

																			gen12620 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12621 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2791)

																			gen12622 := Call(__e, ShenFunc(symcons), gen12621, Nil)

																			gen12623 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen12622)

																			gen12624 := Call(__e, ShenFunc(symcons), gen12623, Nil)

																			gen12625 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12624)

																			gen12626 := Call(__e, ShenFunc(symcons), gen12620, gen12625)

																			gen12627 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12628 := Call(__e, ShenFunc(symcons), gen12626, gen12627)

																			gen12629 := Call(__e, ShenFunc(symcons), gen12619, gen12628)

																			gen12630 := Call(__e, ShenFunc(symbind), V2790, gen12629, V2791, V2792)

																			Result := gen12630
																			Call(__e, ShenFunc(symshen_4unbindv), V2596, V2791)
																			gen12649 = Result

																		} else {
																			gen12649 = False
																		}

																	}
																	Result := gen12649
																	Call(__e, ShenFunc(symshen_4unbindv), V2584, V2791)
																	gen12915 = Result

																} else {
																	gen12915 = False
																}

															}

														} else {
															gen12915 = False
														}

													} else {
														gen12915 = False
													}

												} else {
													gen12915 = False
												}

											} else {
												gen12915 = False
											}

										} else {
											gen12915 = False
										}

									} else {
										gen12915 = False
									}

								} else {
									gen12915 = False
								}

							} else {
								gen12915 = False
							}

						} else {
							gen12915 = False
						}

					} else {
						gen12915 = False
					}
					Case := gen12915
					gen13021 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen13021 {
						gen12916 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

						V2597 := gen12916
						gen13009 := Call(__e, ShenFunc(symcons_2), V2597)

						var gen13010 Obj
						if True == gen13009 {
							gen12917 := Call(__e, ShenFunc(symhd), V2597)

							gen12918 := Call(__e, ShenFunc(symshen_4lazyderef), gen12917, V2791)

							V2598 := gen12918
							gen13008 := Call(__e, ShenFunc(symcons_2), V2598)

							if True == gen13008 {
								gen12919 := Call(__e, ShenFunc(symhd), V2598)

								gen12920 := Call(__e, ShenFunc(symshen_4lazyderef), gen12919, V2791)

								V2599 := gen12920
								gen13007 := Call(__e, ShenFunc(symcons_2), V2599)

								if True == gen13007 {
									gen12921 := Call(__e, ShenFunc(symhd), V2599)

									gen12922 := Call(__e, ShenFunc(symshen_4lazyderef), gen12921, V2791)

									V2600 := gen12922
									gen13006 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), V2600)

									if True == gen13006 {
										gen12923 := Call(__e, ShenFunc(symtl), V2599)

										gen12924 := Call(__e, ShenFunc(symshen_4lazyderef), gen12923, V2791)

										V2601 := gen12924
										gen13005 := Call(__e, ShenFunc(symcons_2), V2601)

										if True == gen13005 {
											gen12925 := Call(__e, ShenFunc(symhd), V2601)

											X := gen12925
											gen12926 := Call(__e, ShenFunc(symtl), V2601)

											gen12927 := Call(__e, ShenFunc(symshen_4lazyderef), gen12926, V2791)

											V2602 := gen12927
											gen13004 := Call(__e, ShenFunc(symcons_2), V2602)

											if True == gen13004 {
												gen12928 := Call(__e, ShenFunc(symhd), V2602)

												Y := gen12928
												gen12929 := Call(__e, ShenFunc(symtl), V2602)

												gen12930 := Call(__e, ShenFunc(symshen_4lazyderef), gen12929, V2791)

												V2603 := gen12930
												gen13003 := Call(__e, ShenFunc(sym_a), Nil, V2603)

												if True == gen13003 {
													gen12931 := Call(__e, ShenFunc(symtl), V2598)

													gen12932 := Call(__e, ShenFunc(symshen_4lazyderef), gen12931, V2791)

													V2604 := gen12932
													gen13002 := Call(__e, ShenFunc(symcons_2), V2604)

													if True == gen13002 {
														gen12933 := Call(__e, ShenFunc(symhd), V2604)

														gen12934 := Call(__e, ShenFunc(symshen_4lazyderef), gen12933, V2791)

														V2605 := gen12934
														gen13001 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2605)

														if True == gen13001 {
															gen12935 := Call(__e, ShenFunc(symtl), V2604)

															gen12936 := Call(__e, ShenFunc(symshen_4lazyderef), gen12935, V2791)

															V2606 := gen12936
															gen13000 := Call(__e, ShenFunc(symcons_2), V2606)

															if True == gen13000 {
																gen12937 := Call(__e, ShenFunc(symhd), V2606)

																gen12938 := Call(__e, ShenFunc(symshen_4lazyderef), gen12937, V2791)

																V2607 := gen12938
																gen12999 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2607)

																if True == gen12999 {
																	gen12970 := Call(__e, ShenFunc(symtl), V2606)

																	gen12971 := Call(__e, ShenFunc(symshen_4lazyderef), gen12970, V2791)

																	V2608 := gen12971
																	gen12998 := Call(__e, ShenFunc(sym_a), Nil, V2608)

																	if True == gen12998 {
																		gen12986 := Call(__e, ShenFunc(symtl), V2597)

																		Hyp := gen12986
																		Call(__e, ShenFunc(symshen_4incinfs))
																		gen12987 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																		gen12988 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																		gen12989 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12988)

																		gen12990 := Call(__e, ShenFunc(symcons), gen12987, gen12989)

																		gen12991 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																		gen12992 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																		gen12993 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12992)

																		gen12994 := Call(__e, ShenFunc(symcons), gen12991, gen12993)

																		gen12995 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																		gen12996 := Call(__e, ShenFunc(symcons), gen12994, gen12995)

																		gen12997 := Call(__e, ShenFunc(symcons), gen12990, gen12996)

																		gen13010 = Call(__e, ShenFunc(symbind), V2790, gen12997, V2791, V2792)

																	} else {
																		gen12985 := Call(__e, ShenFunc(symshen_4pvar_2), V2608)

																		if True == gen12985 {
																			Call(__e, ShenFunc(symshen_4bindv), V2608, Nil, V2791)
																			gen12972 := Call(__e, ShenFunc(symtl), V2597)

																			Hyp := gen12972
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12973 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12974 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen12975 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12974)

																			gen12976 := Call(__e, ShenFunc(symcons), gen12973, gen12975)

																			gen12977 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12978 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen12979 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12978)

																			gen12980 := Call(__e, ShenFunc(symcons), gen12977, gen12979)

																			gen12981 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12982 := Call(__e, ShenFunc(symcons), gen12980, gen12981)

																			gen12983 := Call(__e, ShenFunc(symcons), gen12976, gen12982)

																			gen12984 := Call(__e, ShenFunc(symbind), V2790, gen12983, V2791, V2792)

																			Result := gen12984
																			Call(__e, ShenFunc(symshen_4unbindv), V2608, V2791)
																			gen13010 = Result

																		} else {
																			gen13010 = False
																		}

																	}

																} else {
																	gen12969 := Call(__e, ShenFunc(symshen_4pvar_2), V2607)

																	if True == gen12969 {
																		Call(__e, ShenFunc(symshen_4bindv), V2607, MakeSymbol("string"), V2791)
																		gen12939 := Call(__e, ShenFunc(symtl), V2606)

																		gen12940 := Call(__e, ShenFunc(symshen_4lazyderef), gen12939, V2791)

																		V2609 := gen12940
																		gen12967 := Call(__e, ShenFunc(sym_a), Nil, V2609)

																		var gen12968 Obj
																		if True == gen12967 {
																			gen12955 := Call(__e, ShenFunc(symtl), V2597)

																			Hyp := gen12955
																			Call(__e, ShenFunc(symshen_4incinfs))
																			gen12956 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																			gen12957 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen12958 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12957)

																			gen12959 := Call(__e, ShenFunc(symcons), gen12956, gen12958)

																			gen12960 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																			gen12961 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																			gen12962 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12961)

																			gen12963 := Call(__e, ShenFunc(symcons), gen12960, gen12962)

																			gen12964 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																			gen12965 := Call(__e, ShenFunc(symcons), gen12963, gen12964)

																			gen12966 := Call(__e, ShenFunc(symcons), gen12959, gen12965)

																			gen12968 = Call(__e, ShenFunc(symbind), V2790, gen12966, V2791, V2792)

																		} else {
																			gen12954 := Call(__e, ShenFunc(symshen_4pvar_2), V2609)

																			if True == gen12954 {
																				Call(__e, ShenFunc(symshen_4bindv), V2609, Nil, V2791)
																				gen12941 := Call(__e, ShenFunc(symtl), V2597)

																				Hyp := gen12941
																				Call(__e, ShenFunc(symshen_4incinfs))
																				gen12942 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

																				gen12943 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																				gen12944 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12943)

																				gen12945 := Call(__e, ShenFunc(symcons), gen12942, gen12944)

																				gen12946 := Call(__e, ShenFunc(symshen_4lazyderef), Y, V2791)

																				gen12947 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

																				gen12948 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen12947)

																				gen12949 := Call(__e, ShenFunc(symcons), gen12946, gen12948)

																				gen12950 := Call(__e, ShenFunc(symshen_4lazyderef), Hyp, V2791)

																				gen12951 := Call(__e, ShenFunc(symcons), gen12949, gen12950)

																				gen12952 := Call(__e, ShenFunc(symcons), gen12945, gen12951)

																				gen12953 := Call(__e, ShenFunc(symbind), V2790, gen12952, V2791, V2792)

																				Result := gen12953
																				Call(__e, ShenFunc(symshen_4unbindv), V2609, V2791)
																				gen12968 = Result

																			} else {
																				gen12968 = False
																			}

																		}
																		Result := gen12968
																		Call(__e, ShenFunc(symshen_4unbindv), V2607, V2791)
																		gen13010 = Result

																	} else {
																		gen13010 = False
																	}

																}

															} else {
																gen13010 = False
															}

														} else {
															gen13010 = False
														}

													} else {
														gen13010 = False
													}

												} else {
													gen13010 = False
												}

											} else {
												gen13010 = False
											}

										} else {
											gen13010 = False
										}

									} else {
										gen13010 = False
									}

								} else {
									gen13010 = False
								}

							} else {
								gen13010 = False
							}

						} else {
							gen13010 = False
						}
						Case := gen13010
						gen13020 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen13020 {
							gen13011 := Call(__e, ShenFunc(symshen_4lazyderef), V2789, V2791)

							V2610 := gen13011
							gen13019 := Call(__e, ShenFunc(symcons_2), V2610)

							if True == gen13019 {
								gen13012 := Call(__e, ShenFunc(symhd), V2610)

								X := gen13012
								gen13013 := Call(__e, ShenFunc(symtl), V2610)

								Hyp := gen13013
								gen13014 := Call(__e, ShenFunc(symshen_4newpv), V2791)

								NewHyps := gen13014
								Call(__e, ShenFunc(symshen_4incinfs))
								gen13015 := Call(__e, ShenFunc(symshen_4lazyderef), X, V2791)

								gen13016 := Call(__e, ShenFunc(symshen_4lazyderef), NewHyps, V2791)

								gen13017 := Call(__e, ShenFunc(symcons), gen13015, gen13016)

								gen13018 := MakeNative(func(__e Evaluator) {
									__e.TailApply(ShenFunc(symshen_4t_d_1hyps), Hyp, NewHyps, V2791, V2792)

									return
								}, 0)
								__e.TailApply(ShenFunc(symbind), V2790, gen13017, V2791, gen13018)

								return

							} else {
								__e.Return(False)
								return
							}

						} else {
							__e.Return(Case)
							return
						}

					} else {
						__e.Return(Case)
						return
					}

				} else {
					__e.Return(Case)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-hyps"), gen13024)

		gen13029 := MakeNative(func(__e Evaluator) {
			V2809 := __e.Get(1)
			_ = V2809
			V2810 := __e.Get(2)
			_ = V2810
			V2811 := __e.Get(3)
			_ = V2811
			V2812 := __e.Get(4)
			_ = V2812
			gen13028 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*spy*"))

			if True == gen13028 {
				Call(__e, ShenFunc(symshen_4line))
				gen13025 := Call(__e, ShenFunc(symshen_4deref), V2809, V2811)

				Call(__e, ShenFunc(symshen_4show_1p), gen13025)

				Call(__e, ShenFunc(symnl), MakeNumber(1))
				Call(__e, ShenFunc(symnl), MakeNumber(1))
				gen13026 := Call(__e, ShenFunc(symshen_4deref), V2810, V2811)

				Call(__e, ShenFunc(symshen_4show_1assumptions), gen13026, MakeNumber(1))

				gen13027 := Call(__e, ShenFunc(symstoutput))

				Call(__e, ShenFunc(symshen_4prhush), MakeString("\n> "), gen13027)

				Call(__e, ShenFunc(symshen_4pause_1for_1user))
				__e.TailApply(ShenFunc(symthaw), V2812)

				return

			} else {
				__e.TailApply(ShenFunc(symthaw), V2812)

				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show"), gen13029)

		gen13038 := MakeNative(func(__e Evaluator) {
			gen13030 := Call(__e, ShenFunc(syminferences))

			Infs := gen13030
			gen13031 := Call(__e, ShenFunc(sym_a), MakeNumber(1), Infs)

			var gen13032 Obj
			if True == gen13031 {
				gen13032 = MakeString("")
			} else {
				gen13032 = MakeString("s")
			}
			gen13033 := Call(__e, ShenFunc(symshen_4app), gen13032, MakeString(" \n?- "), MakeSymbol("shen.a"))

			gen13034 := Call(__e, ShenFunc(symcn), MakeString(" inference"), gen13033)

			gen13035 := Call(__e, ShenFunc(symshen_4app), Infs, gen13034, MakeSymbol("shen.a"))

			gen13036 := Call(__e, ShenFunc(symcn), MakeString("____________________________________________________________ "), gen13035)

			gen13037 := Call(__e, ShenFunc(symstoutput))

			__e.TailApply(ShenFunc(symshen_4prhush), gen13036, gen13037)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.line"), gen13038)

		gen13066 := MakeNative(func(__e Evaluator) {
			V2814 := __e.Get(1)
			_ = V2814
			gen13064 := Call(__e, ShenFunc(symcons_2), V2814)

			var gen13065 Obj
			if True == gen13064 {
				gen13061 := Call(__e, ShenFunc(symtl), V2814)

				gen13062 := Call(__e, ShenFunc(symcons_2), gen13061)

				var gen13063 Obj
				if True == gen13062 {
					gen13057 := Call(__e, ShenFunc(symtl), V2814)

					gen13058 := Call(__e, ShenFunc(symhd), gen13057)

					gen13059 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), gen13058)

					var gen13060 Obj
					if True == gen13059 {
						gen13053 := Call(__e, ShenFunc(symtl), V2814)

						gen13054 := Call(__e, ShenFunc(symtl), gen13053)

						gen13055 := Call(__e, ShenFunc(symcons_2), gen13054)

						var gen13056 Obj
						if True == gen13055 {
							gen13049 := Call(__e, ShenFunc(symtl), V2814)

							gen13050 := Call(__e, ShenFunc(symtl), gen13049)

							gen13051 := Call(__e, ShenFunc(symtl), gen13050)

							gen13052 := Call(__e, ShenFunc(sym_a), Nil, gen13051)

							if True == gen13052 {
								gen13056 = True
							} else {
								gen13056 = False
							}

						} else {
							gen13056 = False
						}
						if True == gen13056 {
							gen13060 = True
						} else {
							gen13060 = False
						}

					} else {
						gen13060 = False
					}
					if True == gen13060 {
						gen13063 = True
					} else {
						gen13063 = False
					}

				} else {
					gen13063 = False
				}
				if True == gen13063 {
					gen13065 = True
				} else {
					gen13065 = False
				}

			} else {
				gen13065 = False
			}
			if True == gen13065 {
				gen13041 := Call(__e, ShenFunc(symhd), V2814)

				gen13042 := Call(__e, ShenFunc(symtl), V2814)

				gen13043 := Call(__e, ShenFunc(symtl), gen13042)

				gen13044 := Call(__e, ShenFunc(symhd), gen13043)

				gen13045 := Call(__e, ShenFunc(symshen_4app), gen13044, MakeString(""), MakeSymbol("shen.r"))

				gen13046 := Call(__e, ShenFunc(symcn), MakeString(" : "), gen13045)

				gen13047 := Call(__e, ShenFunc(symshen_4app), gen13041, gen13046, MakeSymbol("shen.r"))

				gen13048 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen13047, gen13048)

				return

			} else {
				gen13039 := Call(__e, ShenFunc(symshen_4app), V2814, MakeString(""), MakeSymbol("shen.r"))

				gen13040 := Call(__e, ShenFunc(symstoutput))

				__e.TailApply(ShenFunc(symshen_4prhush), gen13039, gen13040)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show-p"), gen13066)

		gen13074 := MakeNative(func(__e Evaluator) {
			V2819 := __e.Get(1)
			_ = V2819
			V2820 := __e.Get(2)
			_ = V2820
			gen13073 := Call(__e, ShenFunc(sym_a), Nil, V2819)

			if True == gen13073 {
				__e.Return(MakeSymbol("shen.skip"))
				return
			} else {
				gen13072 := Call(__e, ShenFunc(symcons_2), V2819)

				if True == gen13072 {
					gen13067 := Call(__e, ShenFunc(symshen_4app), V2820, MakeString(". "), MakeSymbol("shen.a"))

					gen13068 := Call(__e, ShenFunc(symstoutput))

					Call(__e, ShenFunc(symshen_4prhush), gen13067, gen13068)

					gen13069 := Call(__e, ShenFunc(symhd), V2819)

					Call(__e, ShenFunc(symshen_4show_1p), gen13069)

					Call(__e, ShenFunc(symnl), MakeNumber(1))
					gen13070 := Call(__e, ShenFunc(symtl), V2819)

					gen13071 := Call(__e, ShenFunc(sym_7), V2820, MakeNumber(1))

					__e.TailApply(ShenFunc(symshen_4show_1assumptions), gen13070, gen13071)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.show-assumptions"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.show-assumptions"), gen13074)

		gen13078 := MakeNative(func(__e Evaluator) {
			gen13075 := Call(__e, ShenFunc(symstinput))

			gen13076 := Call(__e, ShenFunc(symread_1byte), gen13075)

			Byte := gen13076
			gen13077 := Call(__e, ShenFunc(sym_a), Byte, MakeNumber(94))

			if True == gen13077 {
				__e.TailApply(ShenFunc(symsimple_1error), MakeString("input aborted\n"))

				return
			} else {
				__e.TailApply(ShenFunc(symnl), MakeNumber(1))

				return
			}

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.pause-for-user"), gen13078)

		gen13081 := MakeNative(func(__e Evaluator) {
			V2822 := __e.Get(1)
			_ = V2822
			gen13079 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen13080 := Call(__e, ShenFunc(symassoc), V2822, gen13079)

			__e.TailApply(ShenFunc(symcons_2), gen13080)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.typedf?"), gen13081)

		gen13082 := MakeNative(func(__e Evaluator) {
			V2824 := __e.Get(1)
			_ = V2824
			__e.TailApply(ShenFunc(symconcat), MakeSymbol("shen.type-signature-of-"), V2824)

			return
		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.sigf"), gen13082)

		gen13083 := MakeNative(func(__e Evaluator) {
			__e.TailApply(ShenFunc(symgensym), MakeSymbol("&&"))

			return
		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.placeholder"), gen13083)

		gen13172 := MakeNative(func(__e Evaluator) {
			V2829 := __e.Get(1)
			_ = V2829
			V2830 := __e.Get(2)
			_ = V2830
			V2831 := __e.Get(3)
			_ = V2831
			V2832 := __e.Get(4)
			_ = V2832
			gen13084 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

			V2513 := gen13084
			gen13091 := Call(__e, ShenFunc(sym_a), MakeSymbol("number"), V2513)

			var gen13092 Obj
			if True == gen13091 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13089 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

				gen13090 := Call(__e, ShenFunc(symnumber_2), gen13089)

				gen13092 = Call(__e, ShenFunc(symfwhen), gen13090, V2831, V2832)

			} else {
				gen13088 := Call(__e, ShenFunc(symshen_4pvar_2), V2513)

				if True == gen13088 {
					Call(__e, ShenFunc(symshen_4bindv), V2513, MakeSymbol("number"), V2831)
					Call(__e, ShenFunc(symshen_4incinfs))
					gen13085 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

					gen13086 := Call(__e, ShenFunc(symnumber_2), gen13085)

					gen13087 := Call(__e, ShenFunc(symfwhen), gen13086, V2831, V2832)

					Result := gen13087
					Call(__e, ShenFunc(symshen_4unbindv), V2513, V2831)
					gen13092 = Result

				} else {
					gen13092 = False
				}

			}
			Case := gen13092
			gen13171 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13171 {
				gen13093 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

				V2514 := gen13093
				gen13100 := Call(__e, ShenFunc(sym_a), MakeSymbol("boolean"), V2514)

				var gen13101 Obj
				if True == gen13100 {
					Call(__e, ShenFunc(symshen_4incinfs))
					gen13098 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

					gen13099 := Call(__e, ShenFunc(symboolean_2), gen13098)

					gen13101 = Call(__e, ShenFunc(symfwhen), gen13099, V2831, V2832)

				} else {
					gen13097 := Call(__e, ShenFunc(symshen_4pvar_2), V2514)

					if True == gen13097 {
						Call(__e, ShenFunc(symshen_4bindv), V2514, MakeSymbol("boolean"), V2831)
						Call(__e, ShenFunc(symshen_4incinfs))
						gen13094 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

						gen13095 := Call(__e, ShenFunc(symboolean_2), gen13094)

						gen13096 := Call(__e, ShenFunc(symfwhen), gen13095, V2831, V2832)

						Result := gen13096
						Call(__e, ShenFunc(symshen_4unbindv), V2514, V2831)
						gen13101 = Result

					} else {
						gen13101 = False
					}

				}
				Case := gen13101
				gen13170 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen13170 {
					gen13102 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

					V2515 := gen13102
					gen13109 := Call(__e, ShenFunc(sym_a), MakeSymbol("string"), V2515)

					var gen13110 Obj
					if True == gen13109 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen13107 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

						gen13108 := Call(__e, ShenFunc(symstring_2), gen13107)

						gen13110 = Call(__e, ShenFunc(symfwhen), gen13108, V2831, V2832)

					} else {
						gen13106 := Call(__e, ShenFunc(symshen_4pvar_2), V2515)

						if True == gen13106 {
							Call(__e, ShenFunc(symshen_4bindv), V2515, MakeSymbol("string"), V2831)
							Call(__e, ShenFunc(symshen_4incinfs))
							gen13103 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							gen13104 := Call(__e, ShenFunc(symstring_2), gen13103)

							gen13105 := Call(__e, ShenFunc(symfwhen), gen13104, V2831, V2832)

							Result := gen13105
							Call(__e, ShenFunc(symshen_4unbindv), V2515, V2831)
							gen13110 = Result

						} else {
							gen13110 = False
						}

					}
					Case := gen13110
					gen13169 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen13169 {
						gen13111 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

						V2516 := gen13111
						gen13126 := Call(__e, ShenFunc(sym_a), MakeSymbol("symbol"), V2516)

						var gen13127 Obj
						if True == gen13126 {
							Call(__e, ShenFunc(symshen_4incinfs))
							gen13120 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							gen13121 := Call(__e, ShenFunc(symsymbol_2), gen13120)

							gen13125 := MakeNative(func(__e Evaluator) {
								gen13122 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

								gen13123 := Call(__e, ShenFunc(symshen_4ue_2), gen13122)

								gen13124 := Call(__e, ShenFunc(symnot), gen13123)

								__e.TailApply(ShenFunc(symfwhen), gen13124, V2831, V2832)

								return

							}, 0)
							gen13127 = Call(__e, ShenFunc(symfwhen), gen13121, V2831, gen13125)

						} else {
							gen13119 := Call(__e, ShenFunc(symshen_4pvar_2), V2516)

							if True == gen13119 {
								Call(__e, ShenFunc(symshen_4bindv), V2516, MakeSymbol("symbol"), V2831)
								Call(__e, ShenFunc(symshen_4incinfs))
								gen13112 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

								gen13113 := Call(__e, ShenFunc(symsymbol_2), gen13112)

								gen13117 := MakeNative(func(__e Evaluator) {
									gen13114 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

									gen13115 := Call(__e, ShenFunc(symshen_4ue_2), gen13114)

									gen13116 := Call(__e, ShenFunc(symnot), gen13115)

									__e.TailApply(ShenFunc(symfwhen), gen13116, V2831, V2832)

									return

								}, 0)
								gen13118 := Call(__e, ShenFunc(symfwhen), gen13113, V2831, gen13117)

								Result := gen13118
								Call(__e, ShenFunc(symshen_4unbindv), V2516, V2831)
								gen13127 = Result

							} else {
								gen13127 = False
							}

						}
						Case := gen13127
						gen13168 := Call(__e, ShenFunc(sym_a), Case, False)

						if True == gen13168 {
							gen13128 := Call(__e, ShenFunc(symshen_4lazyderef), V2829, V2831)

							V2517 := gen13128
							gen13167 := Call(__e, ShenFunc(sym_a), Nil, V2517)

							if True == gen13167 {
								gen13129 := Call(__e, ShenFunc(symshen_4lazyderef), V2830, V2831)

								V2518 := gen13129
								gen13166 := Call(__e, ShenFunc(symcons_2), V2518)

								if True == gen13166 {
									gen13135 := Call(__e, ShenFunc(symhd), V2518)

									gen13136 := Call(__e, ShenFunc(symshen_4lazyderef), gen13135, V2831)

									V2519 := gen13136
									gen13165 := Call(__e, ShenFunc(sym_a), MakeSymbol("list"), V2519)

									if True == gen13165 {
										gen13152 := Call(__e, ShenFunc(symtl), V2518)

										gen13153 := Call(__e, ShenFunc(symshen_4lazyderef), gen13152, V2831)

										V2520 := gen13153
										gen13164 := Call(__e, ShenFunc(symcons_2), V2520)

										if True == gen13164 {
											gen13158 := Call(__e, ShenFunc(symhd), V2520)

											A := gen13158
											_ = A
											gen13159 := Call(__e, ShenFunc(symtl), V2520)

											gen13160 := Call(__e, ShenFunc(symshen_4lazyderef), gen13159, V2831)

											V2521 := gen13160
											gen13163 := Call(__e, ShenFunc(sym_a), Nil, V2521)

											if True == gen13163 {
												Call(__e, ShenFunc(symshen_4incinfs))
												__e.TailApply(ShenFunc(symthaw), V2832)

												return

											} else {
												gen13162 := Call(__e, ShenFunc(symshen_4pvar_2), V2521)

												if True == gen13162 {
													Call(__e, ShenFunc(symshen_4bindv), V2521, Nil, V2831)
													Call(__e, ShenFunc(symshen_4incinfs))
													gen13161 := Call(__e, ShenFunc(symthaw), V2832)

													Result := gen13161
													Call(__e, ShenFunc(symshen_4unbindv), V2521, V2831)
													__e.Return(Result)
													return

												} else {
													__e.Return(False)
													return
												}

											}

										} else {
											gen13157 := Call(__e, ShenFunc(symshen_4pvar_2), V2520)

											if True == gen13157 {
												gen13154 := Call(__e, ShenFunc(symshen_4newpv), V2831)

												A := gen13154
												gen13155 := Call(__e, ShenFunc(symcons), A, Nil)

												Call(__e, ShenFunc(symshen_4bindv), V2520, gen13155, V2831)

												Call(__e, ShenFunc(symshen_4incinfs))
												gen13156 := Call(__e, ShenFunc(symthaw), V2832)

												Result := gen13156
												Call(__e, ShenFunc(symshen_4unbindv), V2520, V2831)
												__e.Return(Result)
												return

											} else {
												__e.Return(False)
												return
											}

										}

									} else {
										gen13151 := Call(__e, ShenFunc(symshen_4pvar_2), V2519)

										if True == gen13151 {
											Call(__e, ShenFunc(symshen_4bindv), V2519, MakeSymbol("list"), V2831)
											gen13137 := Call(__e, ShenFunc(symtl), V2518)

											gen13138 := Call(__e, ShenFunc(symshen_4lazyderef), gen13137, V2831)

											V2522 := gen13138
											gen13149 := Call(__e, ShenFunc(symcons_2), V2522)

											var gen13150 Obj
											if True == gen13149 {
												gen13143 := Call(__e, ShenFunc(symhd), V2522)

												A := gen13143
												_ = A
												gen13144 := Call(__e, ShenFunc(symtl), V2522)

												gen13145 := Call(__e, ShenFunc(symshen_4lazyderef), gen13144, V2831)

												V2523 := gen13145
												gen13148 := Call(__e, ShenFunc(sym_a), Nil, V2523)

												if True == gen13148 {
													Call(__e, ShenFunc(symshen_4incinfs))
													gen13150 = Call(__e, ShenFunc(symthaw), V2832)

												} else {
													gen13147 := Call(__e, ShenFunc(symshen_4pvar_2), V2523)

													if True == gen13147 {
														Call(__e, ShenFunc(symshen_4bindv), V2523, Nil, V2831)
														Call(__e, ShenFunc(symshen_4incinfs))
														gen13146 := Call(__e, ShenFunc(symthaw), V2832)

														Result := gen13146
														Call(__e, ShenFunc(symshen_4unbindv), V2523, V2831)
														gen13150 = Result

													} else {
														gen13150 = False
													}

												}

											} else {
												gen13142 := Call(__e, ShenFunc(symshen_4pvar_2), V2522)

												if True == gen13142 {
													gen13139 := Call(__e, ShenFunc(symshen_4newpv), V2831)

													A := gen13139
													gen13140 := Call(__e, ShenFunc(symcons), A, Nil)

													Call(__e, ShenFunc(symshen_4bindv), V2522, gen13140, V2831)

													Call(__e, ShenFunc(symshen_4incinfs))
													gen13141 := Call(__e, ShenFunc(symthaw), V2832)

													Result := gen13141
													Call(__e, ShenFunc(symshen_4unbindv), V2522, V2831)
													gen13150 = Result

												} else {
													gen13150 = False
												}

											}
											Result := gen13150
											Call(__e, ShenFunc(symshen_4unbindv), V2519, V2831)
											__e.Return(Result)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								} else {
									gen13134 := Call(__e, ShenFunc(symshen_4pvar_2), V2518)

									if True == gen13134 {
										gen13130 := Call(__e, ShenFunc(symshen_4newpv), V2831)

										A := gen13130
										gen13131 := Call(__e, ShenFunc(symcons), A, Nil)

										gen13132 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen13131)

										Call(__e, ShenFunc(symshen_4bindv), V2518, gen13132, V2831)

										Call(__e, ShenFunc(symshen_4incinfs))
										gen13133 := Call(__e, ShenFunc(symthaw), V2832)

										Result := gen13133
										Call(__e, ShenFunc(symshen_4unbindv), V2518, V2831)
										__e.Return(Result)
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

						} else {
							__e.Return(Case)
							return
						}

					} else {
						__e.Return(Case)
						return
					}

				} else {
					__e.Return(Case)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 4)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.base"), gen13172)

		gen13198 := MakeNative(func(__e Evaluator) {
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
			gen13173 := Call(__e, ShenFunc(symshen_4lazyderef), V2840, V2841)

			V2504 := gen13173
			gen13192 := Call(__e, ShenFunc(symcons_2), V2504)

			var gen13193 Obj
			if True == gen13192 {
				gen13174 := Call(__e, ShenFunc(symhd), V2504)

				gen13175 := Call(__e, ShenFunc(symshen_4lazyderef), gen13174, V2841)

				V2505 := gen13175
				gen13191 := Call(__e, ShenFunc(symcons_2), V2505)

				if True == gen13191 {
					gen13176 := Call(__e, ShenFunc(symhd), V2505)

					Y := gen13176
					gen13177 := Call(__e, ShenFunc(symtl), V2505)

					gen13178 := Call(__e, ShenFunc(symshen_4lazyderef), gen13177, V2841)

					V2506 := gen13178
					gen13190 := Call(__e, ShenFunc(symcons_2), V2506)

					if True == gen13190 {
						gen13179 := Call(__e, ShenFunc(symhd), V2506)

						gen13180 := Call(__e, ShenFunc(symshen_4lazyderef), gen13179, V2841)

						V2507 := gen13180
						gen13189 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2507)

						if True == gen13189 {
							gen13181 := Call(__e, ShenFunc(symtl), V2506)

							gen13182 := Call(__e, ShenFunc(symshen_4lazyderef), gen13181, V2841)

							V2508 := gen13182
							gen13188 := Call(__e, ShenFunc(symcons_2), V2508)

							if True == gen13188 {
								gen13183 := Call(__e, ShenFunc(symhd), V2508)

								B := gen13183
								gen13184 := Call(__e, ShenFunc(symtl), V2508)

								gen13185 := Call(__e, ShenFunc(symshen_4lazyderef), gen13184, V2841)

								V2509 := gen13185
								gen13187 := Call(__e, ShenFunc(sym_a), Nil, V2509)

								if True == gen13187 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen13186 := MakeNative(func(__e Evaluator) {
										__e.TailApply(ShenFunc(symunify_b), V2839, B, V2841, V2842)

										return
									}, 0)
									gen13193 = Call(__e, ShenFunc(symidentical), V2838, Y, V2841, gen13186)

								} else {
									gen13193 = False
								}

							} else {
								gen13193 = False
							}

						} else {
							gen13193 = False
						}

					} else {
						gen13193 = False
					}

				} else {
					gen13193 = False
				}

			} else {
				gen13193 = False
			}
			Case := gen13193
			gen13197 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13197 {
				gen13194 := Call(__e, ShenFunc(symshen_4lazyderef), V2840, V2841)

				V2510 := gen13194
				gen13196 := Call(__e, ShenFunc(symcons_2), V2510)

				if True == gen13196 {
					gen13195 := Call(__e, ShenFunc(symtl), V2510)

					Hyp := gen13195
					Call(__e, ShenFunc(symshen_4incinfs))
					__e.TailApply(ShenFunc(symshen_4by__hypothesis), V2838, V2839, Hyp, V2841, V2842)

					return

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.by_hypothesis"), gen13198)

		gen13217 := MakeNative(func(__e Evaluator) {
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
			gen13199 := Call(__e, ShenFunc(symshen_4lazyderef), V2848, V2851)

			V2498 := gen13199
			gen13216 := Call(__e, ShenFunc(symcons_2), V2498)

			if True == gen13216 {
				gen13200 := Call(__e, ShenFunc(symhd), V2498)

				gen13201 := Call(__e, ShenFunc(symshen_4lazyderef), gen13200, V2851)

				V2499 := gen13201
				gen13215 := Call(__e, ShenFunc(sym_a), MakeSymbol("define"), V2499)

				if True == gen13215 {
					gen13202 := Call(__e, ShenFunc(symtl), V2498)

					gen13203 := Call(__e, ShenFunc(symshen_4lazyderef), gen13202, V2851)

					V2500 := gen13203
					gen13214 := Call(__e, ShenFunc(symcons_2), V2500)

					if True == gen13214 {
						gen13204 := Call(__e, ShenFunc(symhd), V2500)

						F := gen13204
						gen13205 := Call(__e, ShenFunc(symtl), V2500)

						X := gen13205
						gen13206 := Call(__e, ShenFunc(symshen_4newpv), V2851)

						Y := gen13206
						_ = Y
						gen13207 := Call(__e, ShenFunc(symshen_4newpv), V2851)

						E := gen13207
						_ = E
						Call(__e, ShenFunc(symshen_4incinfs))
						gen13208 := MakeNative(func(__e Evaluator) {
							Y := __e.Get(1)
							_ = Y
							__e.TailApply(ShenFunc(symshen_4_5sig_7rules_6), Y)

							return
						}, 1)
						gen13212 := MakeNative(func(__e Evaluator) {
							E := __e.Get(1)
							_ = E
							gen13211 := Call(__e, ShenFunc(symcons_2), E)

							if True == gen13211 {
								gen13209 := Call(__e, ShenFunc(symshen_4app), E, MakeString("\n"), MakeSymbol("shen.s"))

								gen13210 := Call(__e, ShenFunc(symcn), MakeString("parse error here: "), gen13209)

								__e.TailApply(ShenFunc(symsimple_1error), gen13210)

								return

							} else {
								__e.TailApply(ShenFunc(symsimple_1error), MakeString("parse error\n"))

								return
							}

						}, 1)
						gen13213 := Call(__e, ShenFunc(symcompile), gen13208, X, gen13212)

						__e.TailApply(ShenFunc(symshen_4t_d_1defh), gen13213, F, V2849, V2850, V2851, V2852)

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
				__e.Return(False)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-def"), gen13217)

		gen13223 := MakeNative(func(__e Evaluator) {
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
			gen13218 := Call(__e, ShenFunc(symshen_4lazyderef), V2859, V2863)

			V2494 := gen13218
			gen13222 := Call(__e, ShenFunc(symcons_2), V2494)

			if True == gen13222 {
				gen13219 := Call(__e, ShenFunc(symhd), V2494)

				Sig := gen13219
				gen13220 := Call(__e, ShenFunc(symtl), V2494)

				Rules := gen13220
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13221 := Call(__e, ShenFunc(symshen_4ue_1sig), Sig)

				__e.TailApply(ShenFunc(symshen_4t_d_1defhh), Sig, gen13221, V2860, V2861, V2862, Rules, V2863, V2864)

				return

			} else {
				__e.Return(False)
				return
			}

		}, 6)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-defh"), gen13223)

		gen13229 := MakeNative(func(__e Evaluator) {
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
			Call(__e, ShenFunc(symshen_4incinfs))
			gen13224 := Call(__e, ShenFunc(symcons), V2874, Nil)

			gen13225 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13224)

			gen13226 := Call(__e, ShenFunc(symcons), V2875, gen13225)

			gen13227 := Call(__e, ShenFunc(symcons), gen13226, V2877)

			gen13228 := MakeNative(func(__e Evaluator) {
				__e.TailApply(ShenFunc(symshen_4memo), V2875, V2873, V2876, V2879, V2880)

				return
			}, 0)
			__e.TailApply(ShenFunc(symshen_4t_d_1rules), V2878, V2874, MakeNumber(1), V2875, gen13227, V2879, gen13228)

			return

		}, 8)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-defhh"), gen13229)

		gen13235 := MakeNative(func(__e Evaluator) {
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
			gen13230 := Call(__e, ShenFunc(symshen_4newpv), V2889)

			Jnk := gen13230
			Call(__e, ShenFunc(symshen_4incinfs))
			gen13234 := MakeNative(func(__e Evaluator) {
				gen13231 := Call(__e, ShenFunc(symshen_4lazyderef), V2886, V2889)

				gen13232 := Call(__e, ShenFunc(symshen_4lazyderef), V2888, V2889)

				gen13233 := Call(__e, ShenFunc(symdeclare), gen13231, gen13232)

				__e.TailApply(ShenFunc(symbind), Jnk, gen13233, V2889, V2890)

				return

			}, 0)
			__e.TailApply(ShenFunc(symunify_b), V2888, V2887, V2889, gen13234)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.memo"), gen13235)

		gen13248 := MakeNative(func(__e Evaluator) {
			V2892 := __e.Get(1)
			_ = V2892
			gen13236 := Call(__e, ShenFunc(symshen_4_5signature_6), V2892)

			Parse__shen_4_5signature_6 := gen13236
			gen13245 := Call(__e, ShenFunc(symfail))

			gen13246 := Call(__e, ShenFunc(sym_a), gen13245, Parse__shen_4_5signature_6)

			gen13247 := Call(__e, ShenFunc(symnot), gen13246)

			if True == gen13247 {
				gen13237 := Call(__e, ShenFunc(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5signature_6)

				Parse__shen_4_5non_1ll_1rules_6 := gen13237
				gen13242 := Call(__e, ShenFunc(symfail))

				gen13243 := Call(__e, ShenFunc(sym_a), gen13242, Parse__shen_4_5non_1ll_1rules_6)

				gen13244 := Call(__e, ShenFunc(symnot), gen13243)

				if True == gen13244 {
					gen13238 := Call(__e, ShenFunc(symhd), Parse__shen_4_5non_1ll_1rules_6)

					gen13239 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5signature_6)

					gen13240 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

					gen13241 := Call(__e, ShenFunc(symcons), gen13239, gen13240)

					__e.TailApply(ShenFunc(symshen_4pair), gen13238, gen13241)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<sig+rules>"), gen13248)

		gen13271 := MakeNative(func(__e Evaluator) {
			V2894 := __e.Get(1)
			_ = V2894
			gen13249 := Call(__e, ShenFunc(symshen_4_5rule_6), V2894)

			Parse__shen_4_5rule_6 := gen13249
			gen13258 := Call(__e, ShenFunc(symfail))

			gen13259 := Call(__e, ShenFunc(sym_a), gen13258, Parse__shen_4_5rule_6)

			gen13260 := Call(__e, ShenFunc(symnot), gen13259)

			var gen13261 Obj
			if True == gen13260 {
				gen13250 := Call(__e, ShenFunc(symshen_4_5non_1ll_1rules_6), Parse__shen_4_5rule_6)

				Parse__shen_4_5non_1ll_1rules_6 := gen13250
				gen13255 := Call(__e, ShenFunc(symfail))

				gen13256 := Call(__e, ShenFunc(sym_a), gen13255, Parse__shen_4_5non_1ll_1rules_6)

				gen13257 := Call(__e, ShenFunc(symnot), gen13256)

				if True == gen13257 {
					gen13251 := Call(__e, ShenFunc(symhd), Parse__shen_4_5non_1ll_1rules_6)

					gen13252 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen13253 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5non_1ll_1rules_6)

					gen13254 := Call(__e, ShenFunc(symcons), gen13252, gen13253)

					gen13261 = Call(__e, ShenFunc(symshen_4pair), gen13251, gen13254)

				} else {
					gen13261 = Call(__e, ShenFunc(symfail))

				}

			} else {
				gen13261 = Call(__e, ShenFunc(symfail))

			}
			YaccParse := gen13261
			gen13269 := Call(__e, ShenFunc(symfail))

			gen13270 := Call(__e, ShenFunc(sym_a), YaccParse, gen13269)

			if True == gen13270 {
				gen13262 := Call(__e, ShenFunc(symshen_4_5rule_6), V2894)

				Parse__shen_4_5rule_6 := gen13262
				gen13266 := Call(__e, ShenFunc(symfail))

				gen13267 := Call(__e, ShenFunc(sym_a), gen13266, Parse__shen_4_5rule_6)

				gen13268 := Call(__e, ShenFunc(symnot), gen13267)

				if True == gen13268 {
					gen13263 := Call(__e, ShenFunc(symhd), Parse__shen_4_5rule_6)

					gen13264 := Call(__e, ShenFunc(symshen_4hdtl), Parse__shen_4_5rule_6)

					gen13265 := Call(__e, ShenFunc(symcons), gen13264, Nil)

					__e.TailApply(ShenFunc(symshen_4pair), gen13263, gen13265)

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.<non-ll-rules>"), gen13271)

		gen13286 := MakeNative(func(__e Evaluator) {
			V2896 := __e.Get(1)
			_ = V2896
			gen13284 := Call(__e, ShenFunc(symcons_2), V2896)

			var gen13285 Obj
			if True == gen13284 {
				gen13281 := Call(__e, ShenFunc(symtl), V2896)

				gen13282 := Call(__e, ShenFunc(symcons_2), gen13281)

				var gen13283 Obj
				if True == gen13282 {
					gen13277 := Call(__e, ShenFunc(symtl), V2896)

					gen13278 := Call(__e, ShenFunc(symtl), gen13277)

					gen13279 := Call(__e, ShenFunc(sym_a), Nil, gen13278)

					var gen13280 Obj
					if True == gen13279 {
						gen13275 := Call(__e, ShenFunc(symhd), V2896)

						gen13276 := Call(__e, ShenFunc(sym_a), gen13275, MakeSymbol("protect"))

						if True == gen13276 {
							gen13280 = True
						} else {
							gen13280 = False
						}

					} else {
						gen13280 = False
					}
					if True == gen13280 {
						gen13283 = True
					} else {
						gen13283 = False
					}

				} else {
					gen13283 = False
				}
				if True == gen13283 {
					gen13285 = True
				} else {
					gen13285 = False
				}

			} else {
				gen13285 = False
			}
			if True == gen13285 {
				__e.Return(V2896)
				return
			} else {
				gen13274 := Call(__e, ShenFunc(symcons_2), V2896)

				if True == gen13274 {
					gen13273 := MakeNative(func(__e Evaluator) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(ShenFunc(symshen_4ue), Z)

						return
					}, 1)
					__e.TailApply(ShenFunc(symmap), gen13273, V2896)

					return

				} else {
					gen13272 := Call(__e, ShenFunc(symvariable_2), V2896)

					if True == gen13272 {
						__e.TailApply(ShenFunc(symconcat), MakeSymbol("&&"), V2896)

						return
					} else {
						__e.Return(V2896)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue"), gen13286)

		gen13290 := MakeNative(func(__e Evaluator) {
			V2898 := __e.Get(1)
			_ = V2898
			gen13289 := Call(__e, ShenFunc(symcons_2), V2898)

			if True == gen13289 {
				gen13288 := MakeNative(func(__e Evaluator) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(ShenFunc(symshen_4ue_1sig), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symmap), gen13288, V2898)

				return

			} else {
				gen13287 := Call(__e, ShenFunc(symvariable_2), V2898)

				if True == gen13287 {
					__e.TailApply(ShenFunc(symconcat), MakeSymbol("&&&"), V2898)

					return
				} else {
					__e.Return(V2898)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue-sig"), gen13290)

		gen13297 := MakeNative(func(__e Evaluator) {
			V2904 := __e.Get(1)
			_ = V2904
			gen13296 := Call(__e, ShenFunc(symshen_4ue_2), V2904)

			if True == gen13296 {
				__e.TailApply(ShenFunc(symcons), V2904, Nil)

				return
			} else {
				gen13295 := Call(__e, ShenFunc(symcons_2), V2904)

				if True == gen13295 {
					gen13291 := Call(__e, ShenFunc(symhd), V2904)

					gen13292 := Call(__e, ShenFunc(symshen_4ues), gen13291)

					gen13293 := Call(__e, ShenFunc(symtl), V2904)

					gen13294 := Call(__e, ShenFunc(symshen_4ues), gen13293)

					__e.TailApply(ShenFunc(symunion), gen13292, gen13294)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ues"), gen13297)

		gen13301 := MakeNative(func(__e Evaluator) {
			V2906 := __e.Get(1)
			_ = V2906
			gen13300 := Call(__e, ShenFunc(symsymbol_2), V2906)

			if True == gen13300 {
				gen13298 := Call(__e, ShenFunc(symstr), V2906)

				gen13299 := Call(__e, ShenFunc(symshen_4ue_1h_2), gen13298)

				if True == gen13299 {
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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue?"), gen13301)

		gen13313 := MakeNative(func(__e Evaluator) {
			V2914 := __e.Get(1)
			_ = V2914
			gen13311 := Call(__e, ShenFunc(symshen_4_7string_2), V2914)

			var gen13312 Obj
			if True == gen13311 {
				gen13308 := Call(__e, ShenFunc(sympos), V2914, MakeNumber(0))

				gen13309 := Call(__e, ShenFunc(sym_a), MakeString("&"), gen13308)

				var gen13310 Obj
				if True == gen13309 {
					gen13305 := Call(__e, ShenFunc(symtlstr), V2914)

					gen13306 := Call(__e, ShenFunc(symshen_4_7string_2), gen13305)

					var gen13307 Obj
					if True == gen13306 {
						gen13302 := Call(__e, ShenFunc(symtlstr), V2914)

						gen13303 := Call(__e, ShenFunc(sympos), gen13302, MakeNumber(0))

						gen13304 := Call(__e, ShenFunc(sym_a), MakeString("&"), gen13303)

						if True == gen13304 {
							gen13307 = True
						} else {
							gen13307 = False
						}

					} else {
						gen13307 = False
					}
					if True == gen13307 {
						gen13310 = True
					} else {
						gen13310 = False
					}

				} else {
					gen13310 = False
				}
				if True == gen13310 {
					gen13312 = True
				} else {
					gen13312 = False
				}

			} else {
				gen13312 = False
			}
			if True == gen13312 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.ue-h?"), gen13313)

		gen13338 := MakeNative(func(__e Evaluator) {
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
			gen13314 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen13314
			gen13315 := Call(__e, ShenFunc(symshen_4lazyderef), V2922, V2927)

			V2478 := gen13315
			gen13316 := Call(__e, ShenFunc(sym_a), Nil, V2478)

			var gen13317 Obj
			if True == gen13316 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13317 = Call(__e, ShenFunc(symthaw), V2928)

			} else {
				gen13317 = False
			}
			Case := gen13317
			gen13336 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen13337 Obj
			if True == gen13336 {
				gen13318 := Call(__e, ShenFunc(symshen_4lazyderef), V2922, V2927)

				V2479 := gen13318
				gen13325 := Call(__e, ShenFunc(symcons_2), V2479)

				var gen13326 Obj
				if True == gen13325 {
					gen13319 := Call(__e, ShenFunc(symhd), V2479)

					Rule := gen13319
					gen13320 := Call(__e, ShenFunc(symtl), V2479)

					Rules := gen13320
					Call(__e, ShenFunc(symshen_4incinfs))
					gen13321 := Call(__e, ShenFunc(symshen_4ue), Rule)

					gen13324 := MakeNative(func(__e Evaluator) {
						gen13323 := MakeNative(func(__e Evaluator) {
							gen13322 := Call(__e, ShenFunc(sym_7), V2924, MakeNumber(1))

							__e.TailApply(ShenFunc(symshen_4t_d_1rules), Rules, V2923, gen13322, V2925, V2926, V2927, V2928)

							return

						}, 0)
						__e.TailApply(ShenFunc(symcut), Throwcontrol, V2927, gen13323)

						return

					}, 0)
					gen13326 = Call(__e, ShenFunc(symshen_4t_d_1rule), gen13321, V2923, V2926, V2927, gen13324)

				} else {
					gen13326 = False
				}
				Case := gen13326
				gen13335 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen13335 {
					gen13327 := Call(__e, ShenFunc(symshen_4newpv), V2927)

					Err := gen13327
					Call(__e, ShenFunc(symshen_4incinfs))
					gen13328 := Call(__e, ShenFunc(symshen_4lazyderef), V2924, V2927)

					gen13329 := Call(__e, ShenFunc(symshen_4lazyderef), V2925, V2927)

					gen13330 := Call(__e, ShenFunc(symshen_4app), gen13329, MakeString(""), MakeSymbol("shen.a"))

					gen13331 := Call(__e, ShenFunc(symcn), MakeString(" of "), gen13330)

					gen13332 := Call(__e, ShenFunc(symshen_4app), gen13328, gen13331, MakeSymbol("shen.a"))

					gen13333 := Call(__e, ShenFunc(symcn), MakeString("type error in rule "), gen13332)

					gen13334 := Call(__e, ShenFunc(symsimple_1error), gen13333)

					gen13337 = Call(__e, ShenFunc(symbind), Err, gen13334, V2927, V2928)

				} else {
					gen13337 = Case
				}

			} else {
				gen13337 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen13337)

			return

		}, 7)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-rules"), gen13338)

		gen13360 := MakeNative(func(__e Evaluator) {
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
			gen13339 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen13339
			gen13340 := Call(__e, ShenFunc(symshen_4lazyderef), V2934, V2937)

			V2470 := gen13340
			gen13358 := Call(__e, ShenFunc(symcons_2), V2470)

			var gen13359 Obj
			if True == gen13358 {
				gen13341 := Call(__e, ShenFunc(symhd), V2470)

				Patterns := gen13341
				gen13342 := Call(__e, ShenFunc(symtl), V2470)

				gen13343 := Call(__e, ShenFunc(symshen_4lazyderef), gen13342, V2937)

				V2471 := gen13343
				gen13357 := Call(__e, ShenFunc(symcons_2), V2471)

				if True == gen13357 {
					gen13344 := Call(__e, ShenFunc(symhd), V2471)

					Action := gen13344
					gen13345 := Call(__e, ShenFunc(symtl), V2471)

					gen13346 := Call(__e, ShenFunc(symshen_4lazyderef), gen13345, V2937)

					V2472 := gen13346
					gen13356 := Call(__e, ShenFunc(sym_a), Nil, V2472)

					if True == gen13356 {
						gen13347 := Call(__e, ShenFunc(symshen_4newpv), V2937)

						NewHyps := gen13347
						Call(__e, ShenFunc(symshen_4incinfs))
						gen13348 := Call(__e, ShenFunc(symshen_4placeholders), Patterns)

						gen13355 := MakeNative(func(__e Evaluator) {
							gen13354 := MakeNative(func(__e Evaluator) {
								gen13353 := MakeNative(func(__e Evaluator) {
									gen13349 := Call(__e, ShenFunc(symshen_4ue), Action)

									gen13350 := Call(__e, ShenFunc(symshen_4curry), gen13349)

									gen13351 := Call(__e, ShenFunc(symshen_4result_1type), Patterns, V2935)

									gen13352 := Call(__e, ShenFunc(symshen_4patthyps), Patterns, V2935, V2936)

									__e.TailApply(ShenFunc(symshen_4t_d_1action), gen13350, gen13351, gen13352, V2937, V2938)

									return

								}, 0)
								__e.TailApply(ShenFunc(symcut), Throwcontrol, V2937, gen13353)

								return

							}, 0)
							__e.TailApply(ShenFunc(symshen_4t_d_1patterns), Patterns, V2935, NewHyps, V2937, gen13354)

							return

						}, 0)
						gen13359 = Call(__e, ShenFunc(symshen_4newhyps), gen13348, V2936, NewHyps, V2937, gen13355)

					} else {
						gen13359 = False
					}

				} else {
					gen13359 = False
				}

			} else {
				gen13359 = False
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen13359)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-rule"), gen13360)

		gen13367 := MakeNative(func(__e Evaluator) {
			V2944 := __e.Get(1)
			_ = V2944
			gen13366 := Call(__e, ShenFunc(symshen_4ue_2), V2944)

			if True == gen13366 {
				__e.TailApply(ShenFunc(symcons), V2944, Nil)

				return
			} else {
				gen13365 := Call(__e, ShenFunc(symcons_2), V2944)

				if True == gen13365 {
					gen13361 := Call(__e, ShenFunc(symhd), V2944)

					gen13362 := Call(__e, ShenFunc(symshen_4placeholders), gen13361)

					gen13363 := Call(__e, ShenFunc(symtl), V2944)

					gen13364 := Call(__e, ShenFunc(symshen_4placeholders), gen13363)

					__e.TailApply(ShenFunc(symunion), gen13362, gen13364)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.placeholders"), gen13367)

		gen13454 := MakeNative(func(__e Evaluator) {
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
			gen13368 := Call(__e, ShenFunc(symshen_4lazyderef), V2950, V2953)

			V2457 := gen13368
			gen13369 := Call(__e, ShenFunc(sym_a), Nil, V2457)

			var gen13370 Obj
			if True == gen13369 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13370 = Call(__e, ShenFunc(symunify_b), V2952, V2951, V2953, V2954)

			} else {
				gen13370 = False
			}
			Case := gen13370
			gen13453 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13453 {
				gen13371 := Call(__e, ShenFunc(symshen_4lazyderef), V2950, V2953)

				V2458 := gen13371
				gen13452 := Call(__e, ShenFunc(symcons_2), V2458)

				if True == gen13452 {
					gen13372 := Call(__e, ShenFunc(symhd), V2458)

					V2453 := gen13372
					gen13373 := Call(__e, ShenFunc(symtl), V2458)

					Vs := gen13373
					gen13374 := Call(__e, ShenFunc(symshen_4lazyderef), V2952, V2953)

					V2459 := gen13374
					gen13451 := Call(__e, ShenFunc(symcons_2), V2459)

					if True == gen13451 {
						gen13385 := Call(__e, ShenFunc(symhd), V2459)

						gen13386 := Call(__e, ShenFunc(symshen_4lazyderef), gen13385, V2953)

						V2460 := gen13386
						gen13450 := Call(__e, ShenFunc(symcons_2), V2460)

						if True == gen13450 {
							gen13396 := Call(__e, ShenFunc(symhd), V2460)

							V := gen13396
							gen13397 := Call(__e, ShenFunc(symtl), V2460)

							gen13398 := Call(__e, ShenFunc(symshen_4lazyderef), gen13397, V2953)

							V2461 := gen13398
							gen13449 := Call(__e, ShenFunc(symcons_2), V2461)

							if True == gen13449 {
								gen13406 := Call(__e, ShenFunc(symhd), V2461)

								gen13407 := Call(__e, ShenFunc(symshen_4lazyderef), gen13406, V2953)

								V2462 := gen13407
								gen13448 := Call(__e, ShenFunc(sym_a), MakeSymbol(":"), V2462)

								if True == gen13448 {
									gen13429 := Call(__e, ShenFunc(symtl), V2461)

									gen13430 := Call(__e, ShenFunc(symshen_4lazyderef), gen13429, V2953)

									V2463 := gen13430
									gen13447 := Call(__e, ShenFunc(symcons_2), V2463)

									if True == gen13447 {
										gen13437 := Call(__e, ShenFunc(symhd), V2463)

										A := gen13437
										_ = A
										gen13438 := Call(__e, ShenFunc(symtl), V2463)

										gen13439 := Call(__e, ShenFunc(symshen_4lazyderef), gen13438, V2953)

										V2464 := gen13439
										gen13446 := Call(__e, ShenFunc(sym_a), Nil, V2464)

										if True == gen13446 {
											gen13444 := Call(__e, ShenFunc(symtl), V2459)

											NewHyp := gen13444
											Call(__e, ShenFunc(symshen_4incinfs))
											gen13445 := MakeNative(func(__e Evaluator) {
												__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

												return
											}, 0)
											__e.TailApply(ShenFunc(symunify_b), V, V2453, V2953, gen13445)

											return

										} else {
											gen13443 := Call(__e, ShenFunc(symshen_4pvar_2), V2464)

											if True == gen13443 {
												Call(__e, ShenFunc(symshen_4bindv), V2464, Nil, V2953)
												gen13440 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen13440
												Call(__e, ShenFunc(symshen_4incinfs))
												gen13441 := MakeNative(func(__e Evaluator) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen13442 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13441)

												Result := gen13442
												Call(__e, ShenFunc(symshen_4unbindv), V2464, V2953)
												__e.Return(Result)
												return

											} else {
												__e.Return(False)
												return
											}

										}

									} else {
										gen13436 := Call(__e, ShenFunc(symshen_4pvar_2), V2463)

										if True == gen13436 {
											gen13431 := Call(__e, ShenFunc(symshen_4newpv), V2953)

											A := gen13431
											gen13432 := Call(__e, ShenFunc(symcons), A, Nil)

											Call(__e, ShenFunc(symshen_4bindv), V2463, gen13432, V2953)

											gen13433 := Call(__e, ShenFunc(symtl), V2459)

											NewHyp := gen13433
											Call(__e, ShenFunc(symshen_4incinfs))
											gen13434 := MakeNative(func(__e Evaluator) {
												__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

												return
											}, 0)
											gen13435 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13434)

											Result := gen13435
											Call(__e, ShenFunc(symshen_4unbindv), V2463, V2953)
											__e.Return(Result)
											return

										} else {
											__e.Return(False)
											return
										}

									}

								} else {
									gen13428 := Call(__e, ShenFunc(symshen_4pvar_2), V2462)

									if True == gen13428 {
										Call(__e, ShenFunc(symshen_4bindv), V2462, MakeSymbol(":"), V2953)
										gen13408 := Call(__e, ShenFunc(symtl), V2461)

										gen13409 := Call(__e, ShenFunc(symshen_4lazyderef), gen13408, V2953)

										V2465 := gen13409
										gen13426 := Call(__e, ShenFunc(symcons_2), V2465)

										var gen13427 Obj
										if True == gen13426 {
											gen13416 := Call(__e, ShenFunc(symhd), V2465)

											A := gen13416
											_ = A
											gen13417 := Call(__e, ShenFunc(symtl), V2465)

											gen13418 := Call(__e, ShenFunc(symshen_4lazyderef), gen13417, V2953)

											V2466 := gen13418
											gen13425 := Call(__e, ShenFunc(sym_a), Nil, V2466)

											if True == gen13425 {
												gen13423 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen13423
												Call(__e, ShenFunc(symshen_4incinfs))
												gen13424 := MakeNative(func(__e Evaluator) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen13427 = Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13424)

											} else {
												gen13422 := Call(__e, ShenFunc(symshen_4pvar_2), V2466)

												if True == gen13422 {
													Call(__e, ShenFunc(symshen_4bindv), V2466, Nil, V2953)
													gen13419 := Call(__e, ShenFunc(symtl), V2459)

													NewHyp := gen13419
													Call(__e, ShenFunc(symshen_4incinfs))
													gen13420 := MakeNative(func(__e Evaluator) {
														__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

														return
													}, 0)
													gen13421 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13420)

													Result := gen13421
													Call(__e, ShenFunc(symshen_4unbindv), V2466, V2953)
													gen13427 = Result

												} else {
													gen13427 = False
												}

											}

										} else {
											gen13415 := Call(__e, ShenFunc(symshen_4pvar_2), V2465)

											if True == gen13415 {
												gen13410 := Call(__e, ShenFunc(symshen_4newpv), V2953)

												A := gen13410
												gen13411 := Call(__e, ShenFunc(symcons), A, Nil)

												Call(__e, ShenFunc(symshen_4bindv), V2465, gen13411, V2953)

												gen13412 := Call(__e, ShenFunc(symtl), V2459)

												NewHyp := gen13412
												Call(__e, ShenFunc(symshen_4incinfs))
												gen13413 := MakeNative(func(__e Evaluator) {
													__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

													return
												}, 0)
												gen13414 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13413)

												Result := gen13414
												Call(__e, ShenFunc(symshen_4unbindv), V2465, V2953)
												gen13427 = Result

											} else {
												gen13427 = False
											}

										}
										Result := gen13427
										Call(__e, ShenFunc(symshen_4unbindv), V2462, V2953)
										__e.Return(Result)
										return

									} else {
										__e.Return(False)
										return
									}

								}

							} else {
								gen13405 := Call(__e, ShenFunc(symshen_4pvar_2), V2461)

								if True == gen13405 {
									gen13399 := Call(__e, ShenFunc(symshen_4newpv), V2953)

									A := gen13399
									gen13400 := Call(__e, ShenFunc(symcons), A, Nil)

									gen13401 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13400)

									Call(__e, ShenFunc(symshen_4bindv), V2461, gen13401, V2953)

									gen13402 := Call(__e, ShenFunc(symtl), V2459)

									NewHyp := gen13402
									Call(__e, ShenFunc(symshen_4incinfs))
									gen13403 := MakeNative(func(__e Evaluator) {
										__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

										return
									}, 0)
									gen13404 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13403)

									Result := gen13404
									Call(__e, ShenFunc(symshen_4unbindv), V2461, V2953)
									__e.Return(Result)
									return

								} else {
									__e.Return(False)
									return
								}

							}

						} else {
							gen13395 := Call(__e, ShenFunc(symshen_4pvar_2), V2460)

							if True == gen13395 {
								gen13387 := Call(__e, ShenFunc(symshen_4newpv), V2953)

								V := gen13387
								gen13388 := Call(__e, ShenFunc(symshen_4newpv), V2953)

								A := gen13388
								gen13389 := Call(__e, ShenFunc(symcons), A, Nil)

								gen13390 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13389)

								gen13391 := Call(__e, ShenFunc(symcons), V, gen13390)

								Call(__e, ShenFunc(symshen_4bindv), V2460, gen13391, V2953)

								gen13392 := Call(__e, ShenFunc(symtl), V2459)

								NewHyp := gen13392
								Call(__e, ShenFunc(symshen_4incinfs))
								gen13393 := MakeNative(func(__e Evaluator) {
									__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

									return
								}, 0)
								gen13394 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13393)

								Result := gen13394
								Call(__e, ShenFunc(symshen_4unbindv), V2460, V2953)
								__e.Return(Result)
								return

							} else {
								__e.Return(False)
								return
							}

						}

					} else {
						gen13384 := Call(__e, ShenFunc(symshen_4pvar_2), V2459)

						if True == gen13384 {
							gen13375 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							V := gen13375
							gen13376 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							A := gen13376
							gen13377 := Call(__e, ShenFunc(symshen_4newpv), V2953)

							NewHyp := gen13377
							gen13378 := Call(__e, ShenFunc(symcons), A, Nil)

							gen13379 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13378)

							gen13380 := Call(__e, ShenFunc(symcons), V, gen13379)

							gen13381 := Call(__e, ShenFunc(symcons), gen13380, NewHyp)

							Call(__e, ShenFunc(symshen_4bindv), V2459, gen13381, V2953)

							Call(__e, ShenFunc(symshen_4incinfs))
							gen13382 := MakeNative(func(__e Evaluator) {
								__e.TailApply(ShenFunc(symshen_4newhyps), Vs, V2951, NewHyp, V2953, V2954)

								return
							}, 0)
							gen13383 := Call(__e, ShenFunc(symunify_b), V, V2453, V2953, gen13382)

							Result := gen13383
							Call(__e, ShenFunc(symshen_4unbindv), V2459, V2953)
							__e.Return(Result)
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

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.newhyps"), gen13454)

		gen13486 := MakeNative(func(__e Evaluator) {
			V2960 := __e.Get(1)
			_ = V2960
			V2961 := __e.Get(2)
			_ = V2961
			V2962 := __e.Get(3)
			_ = V2962
			gen13485 := Call(__e, ShenFunc(sym_a), Nil, V2960)

			if True == gen13485 {
				__e.Return(V2962)
				return
			} else {
				gen13483 := Call(__e, ShenFunc(symcons_2), V2960)

				var gen13484 Obj
				if True == gen13483 {
					gen13481 := Call(__e, ShenFunc(symcons_2), V2961)

					var gen13482 Obj
					if True == gen13481 {
						gen13478 := Call(__e, ShenFunc(symtl), V2961)

						gen13479 := Call(__e, ShenFunc(symcons_2), gen13478)

						var gen13480 Obj
						if True == gen13479 {
							gen13474 := Call(__e, ShenFunc(symtl), V2961)

							gen13475 := Call(__e, ShenFunc(symhd), gen13474)

							gen13476 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen13475)

							var gen13477 Obj
							if True == gen13476 {
								gen13470 := Call(__e, ShenFunc(symtl), V2961)

								gen13471 := Call(__e, ShenFunc(symtl), gen13470)

								gen13472 := Call(__e, ShenFunc(symcons_2), gen13471)

								var gen13473 Obj
								if True == gen13472 {
									gen13466 := Call(__e, ShenFunc(symtl), V2961)

									gen13467 := Call(__e, ShenFunc(symtl), gen13466)

									gen13468 := Call(__e, ShenFunc(symtl), gen13467)

									gen13469 := Call(__e, ShenFunc(sym_a), Nil, gen13468)

									if True == gen13469 {
										gen13473 = True
									} else {
										gen13473 = False
									}

								} else {
									gen13473 = False
								}
								if True == gen13473 {
									gen13477 = True
								} else {
									gen13477 = False
								}

							} else {
								gen13477 = False
							}
							if True == gen13477 {
								gen13480 = True
							} else {
								gen13480 = False
							}

						} else {
							gen13480 = False
						}
						if True == gen13480 {
							gen13482 = True
						} else {
							gen13482 = False
						}

					} else {
						gen13482 = False
					}
					if True == gen13482 {
						gen13484 = True
					} else {
						gen13484 = False
					}

				} else {
					gen13484 = False
				}
				if True == gen13484 {
					gen13455 := Call(__e, ShenFunc(symhd), V2960)

					gen13456 := Call(__e, ShenFunc(symshen_4curry), gen13455)

					gen13457 := Call(__e, ShenFunc(symhd), V2961)

					gen13458 := Call(__e, ShenFunc(symcons), gen13457, Nil)

					gen13459 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13458)

					gen13460 := Call(__e, ShenFunc(symcons), gen13456, gen13459)

					gen13461 := Call(__e, ShenFunc(symtl), V2960)

					gen13462 := Call(__e, ShenFunc(symtl), V2961)

					gen13463 := Call(__e, ShenFunc(symtl), gen13462)

					gen13464 := Call(__e, ShenFunc(symhd), gen13463)

					gen13465 := Call(__e, ShenFunc(symshen_4patthyps), gen13461, gen13464, V2962)

					__e.TailApply(ShenFunc(symadjoin), gen13460, gen13465)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.patthyps"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.patthyps"), gen13486)

		gen13525 := MakeNative(func(__e Evaluator) {
			V2969 := __e.Get(1)
			_ = V2969
			V2970 := __e.Get(2)
			_ = V2970
			gen13523 := Call(__e, ShenFunc(sym_a), Nil, V2969)

			var gen13524 Obj
			if True == gen13523 {
				gen13521 := Call(__e, ShenFunc(symcons_2), V2970)

				var gen13522 Obj
				if True == gen13521 {
					gen13518 := Call(__e, ShenFunc(symhd), V2970)

					gen13519 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen13518)

					var gen13520 Obj
					if True == gen13519 {
						gen13515 := Call(__e, ShenFunc(symtl), V2970)

						gen13516 := Call(__e, ShenFunc(symcons_2), gen13515)

						var gen13517 Obj
						if True == gen13516 {
							gen13512 := Call(__e, ShenFunc(symtl), V2970)

							gen13513 := Call(__e, ShenFunc(symtl), gen13512)

							gen13514 := Call(__e, ShenFunc(sym_a), Nil, gen13513)

							if True == gen13514 {
								gen13517 = True
							} else {
								gen13517 = False
							}

						} else {
							gen13517 = False
						}
						if True == gen13517 {
							gen13520 = True
						} else {
							gen13520 = False
						}

					} else {
						gen13520 = False
					}
					if True == gen13520 {
						gen13522 = True
					} else {
						gen13522 = False
					}

				} else {
					gen13522 = False
				}
				if True == gen13522 {
					gen13524 = True
				} else {
					gen13524 = False
				}

			} else {
				gen13524 = False
			}
			if True == gen13524 {
				gen13511 := Call(__e, ShenFunc(symtl), V2970)

				__e.TailApply(ShenFunc(symhd), gen13511)

				return

			} else {
				gen13510 := Call(__e, ShenFunc(sym_a), Nil, V2969)

				if True == gen13510 {
					__e.Return(V2970)
					return
				} else {
					gen13508 := Call(__e, ShenFunc(symcons_2), V2969)

					var gen13509 Obj
					if True == gen13508 {
						gen13506 := Call(__e, ShenFunc(symcons_2), V2970)

						var gen13507 Obj
						if True == gen13506 {
							gen13503 := Call(__e, ShenFunc(symtl), V2970)

							gen13504 := Call(__e, ShenFunc(symcons_2), gen13503)

							var gen13505 Obj
							if True == gen13504 {
								gen13499 := Call(__e, ShenFunc(symtl), V2970)

								gen13500 := Call(__e, ShenFunc(symhd), gen13499)

								gen13501 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), gen13500)

								var gen13502 Obj
								if True == gen13501 {
									gen13495 := Call(__e, ShenFunc(symtl), V2970)

									gen13496 := Call(__e, ShenFunc(symtl), gen13495)

									gen13497 := Call(__e, ShenFunc(symcons_2), gen13496)

									var gen13498 Obj
									if True == gen13497 {
										gen13491 := Call(__e, ShenFunc(symtl), V2970)

										gen13492 := Call(__e, ShenFunc(symtl), gen13491)

										gen13493 := Call(__e, ShenFunc(symtl), gen13492)

										gen13494 := Call(__e, ShenFunc(sym_a), Nil, gen13493)

										if True == gen13494 {
											gen13498 = True
										} else {
											gen13498 = False
										}

									} else {
										gen13498 = False
									}
									if True == gen13498 {
										gen13502 = True
									} else {
										gen13502 = False
									}

								} else {
									gen13502 = False
								}
								if True == gen13502 {
									gen13505 = True
								} else {
									gen13505 = False
								}

							} else {
								gen13505 = False
							}
							if True == gen13505 {
								gen13507 = True
							} else {
								gen13507 = False
							}

						} else {
							gen13507 = False
						}
						if True == gen13507 {
							gen13509 = True
						} else {
							gen13509 = False
						}

					} else {
						gen13509 = False
					}
					if True == gen13509 {
						gen13487 := Call(__e, ShenFunc(symtl), V2969)

						gen13488 := Call(__e, ShenFunc(symtl), V2970)

						gen13489 := Call(__e, ShenFunc(symtl), gen13488)

						gen13490 := Call(__e, ShenFunc(symhd), gen13489)

						__e.TailApply(ShenFunc(symshen_4result_1type), gen13487, gen13490)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.result-type"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.result-type"), gen13525)

		gen13555 := MakeNative(func(__e Evaluator) {
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
			gen13526 := Call(__e, ShenFunc(symshen_4lazyderef), V2976, V2979)

			V2445 := gen13526
			gen13527 := Call(__e, ShenFunc(sym_a), Nil, V2445)

			var gen13528 Obj
			if True == gen13527 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13528 = Call(__e, ShenFunc(symthaw), V2980)

			} else {
				gen13528 = False
			}
			Case := gen13528
			gen13554 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13554 {
				gen13529 := Call(__e, ShenFunc(symshen_4lazyderef), V2976, V2979)

				V2446 := gen13529
				gen13553 := Call(__e, ShenFunc(symcons_2), V2446)

				if True == gen13553 {
					gen13530 := Call(__e, ShenFunc(symhd), V2446)

					Pattern := gen13530
					gen13531 := Call(__e, ShenFunc(symtl), V2446)

					Patterns := gen13531
					gen13532 := Call(__e, ShenFunc(symshen_4lazyderef), V2977, V2979)

					V2447 := gen13532
					gen13552 := Call(__e, ShenFunc(symcons_2), V2447)

					if True == gen13552 {
						gen13533 := Call(__e, ShenFunc(symhd), V2447)

						A := gen13533
						gen13534 := Call(__e, ShenFunc(symtl), V2447)

						gen13535 := Call(__e, ShenFunc(symshen_4lazyderef), gen13534, V2979)

						V2448 := gen13535
						gen13551 := Call(__e, ShenFunc(symcons_2), V2448)

						if True == gen13551 {
							gen13536 := Call(__e, ShenFunc(symhd), V2448)

							gen13537 := Call(__e, ShenFunc(symshen_4lazyderef), gen13536, V2979)

							V2449 := gen13537
							gen13550 := Call(__e, ShenFunc(sym_a), MakeSymbol("-->"), V2449)

							if True == gen13550 {
								gen13538 := Call(__e, ShenFunc(symtl), V2448)

								gen13539 := Call(__e, ShenFunc(symshen_4lazyderef), gen13538, V2979)

								V2450 := gen13539
								gen13549 := Call(__e, ShenFunc(symcons_2), V2450)

								if True == gen13549 {
									gen13540 := Call(__e, ShenFunc(symhd), V2450)

									B := gen13540
									gen13541 := Call(__e, ShenFunc(symtl), V2450)

									gen13542 := Call(__e, ShenFunc(symshen_4lazyderef), gen13541, V2979)

									V2451 := gen13542
									gen13548 := Call(__e, ShenFunc(sym_a), Nil, V2451)

									if True == gen13548 {
										Call(__e, ShenFunc(symshen_4incinfs))
										gen13543 := Call(__e, ShenFunc(symshen_4curry), Pattern)

										gen13544 := Call(__e, ShenFunc(symcons), A, Nil)

										gen13545 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13544)

										gen13546 := Call(__e, ShenFunc(symcons), gen13543, gen13545)

										gen13547 := MakeNative(func(__e Evaluator) {
											__e.TailApply(ShenFunc(symshen_4t_d_1patterns), Patterns, B, V2978, V2979, V2980)

											return
										}, 0)
										__e.TailApply(ShenFunc(symshen_4t_d), gen13546, V2978, V2979, gen13547)

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
								__e.Return(False)
								return
							}

						} else {
							__e.Return(False)
							return
						}

					} else {
						__e.Return(False)
						return
					}

				} else {
					__e.Return(False)
					return
				}

			} else {
				__e.Return(Case)
				return
			}

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-patterns"), gen13555)

		gen13658 := MakeNative(func(__e Evaluator) {
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
			gen13556 := Call(__e, ShenFunc(symshen_4catchpoint))

			Throwcontrol := gen13556
			gen13557 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

			V2422 := gen13557
			gen13582 := Call(__e, ShenFunc(symcons_2), V2422)

			var gen13583 Obj
			if True == gen13582 {
				gen13558 := Call(__e, ShenFunc(symhd), V2422)

				gen13559 := Call(__e, ShenFunc(symshen_4lazyderef), gen13558, V2989)

				V2423 := gen13559
				gen13581 := Call(__e, ShenFunc(sym_a), MakeSymbol("where"), V2423)

				if True == gen13581 {
					gen13560 := Call(__e, ShenFunc(symtl), V2422)

					gen13561 := Call(__e, ShenFunc(symshen_4lazyderef), gen13560, V2989)

					V2424 := gen13561
					gen13580 := Call(__e, ShenFunc(symcons_2), V2424)

					if True == gen13580 {
						gen13562 := Call(__e, ShenFunc(symhd), V2424)

						P := gen13562
						gen13563 := Call(__e, ShenFunc(symtl), V2424)

						gen13564 := Call(__e, ShenFunc(symshen_4lazyderef), gen13563, V2989)

						V2425 := gen13564
						gen13579 := Call(__e, ShenFunc(symcons_2), V2425)

						if True == gen13579 {
							gen13565 := Call(__e, ShenFunc(symhd), V2425)

							Action := gen13565
							gen13566 := Call(__e, ShenFunc(symtl), V2425)

							gen13567 := Call(__e, ShenFunc(symshen_4lazyderef), gen13566, V2989)

							V2426 := gen13567
							gen13578 := Call(__e, ShenFunc(sym_a), Nil, V2426)

							if True == gen13578 {
								Call(__e, ShenFunc(symshen_4incinfs))
								gen13577 := MakeNative(func(__e Evaluator) {
									gen13568 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

									gen13569 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13568)

									gen13570 := Call(__e, ShenFunc(symcons), P, gen13569)

									gen13576 := MakeNative(func(__e Evaluator) {
										gen13575 := MakeNative(func(__e Evaluator) {
											gen13571 := Call(__e, ShenFunc(symcons), MakeSymbol("verified"), Nil)

											gen13572 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13571)

											gen13573 := Call(__e, ShenFunc(symcons), P, gen13572)

											gen13574 := Call(__e, ShenFunc(symcons), gen13573, V2988)

											__e.TailApply(ShenFunc(symshen_4t_d_1action), Action, V2987, gen13574, V2989, V2990)

											return

										}, 0)
										__e.TailApply(ShenFunc(symcut), Throwcontrol, V2989, gen13575)

										return

									}, 0)
									__e.TailApply(ShenFunc(symshen_4t_d), gen13570, V2988, V2989, gen13576)

									return

								}, 0)
								gen13583 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen13577)

							} else {
								gen13583 = False
							}

						} else {
							gen13583 = False
						}

					} else {
						gen13583 = False
					}

				} else {
					gen13583 = False
				}

			} else {
				gen13583 = False
			}
			Case := gen13583
			gen13656 := Call(__e, ShenFunc(sym_a), Case, False)

			var gen13657 Obj
			if True == gen13656 {
				gen13584 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

				V2427 := gen13584
				gen13625 := Call(__e, ShenFunc(symcons_2), V2427)

				var gen13626 Obj
				if True == gen13625 {
					gen13585 := Call(__e, ShenFunc(symhd), V2427)

					gen13586 := Call(__e, ShenFunc(symshen_4lazyderef), gen13585, V2989)

					V2428 := gen13586
					gen13624 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), V2428)

					if True == gen13624 {
						gen13587 := Call(__e, ShenFunc(symtl), V2427)

						gen13588 := Call(__e, ShenFunc(symshen_4lazyderef), gen13587, V2989)

						V2429 := gen13588
						gen13623 := Call(__e, ShenFunc(symcons_2), V2429)

						if True == gen13623 {
							gen13589 := Call(__e, ShenFunc(symhd), V2429)

							gen13590 := Call(__e, ShenFunc(symshen_4lazyderef), gen13589, V2989)

							V2430 := gen13590
							gen13622 := Call(__e, ShenFunc(symcons_2), V2430)

							if True == gen13622 {
								gen13591 := Call(__e, ShenFunc(symhd), V2430)

								gen13592 := Call(__e, ShenFunc(symshen_4lazyderef), gen13591, V2989)

								V2431 := gen13592
								gen13621 := Call(__e, ShenFunc(symcons_2), V2431)

								if True == gen13621 {
									gen13593 := Call(__e, ShenFunc(symhd), V2431)

									gen13594 := Call(__e, ShenFunc(symshen_4lazyderef), gen13593, V2989)

									V2432 := gen13594
									gen13620 := Call(__e, ShenFunc(sym_a), MakeSymbol("fail-if"), V2432)

									if True == gen13620 {
										gen13595 := Call(__e, ShenFunc(symtl), V2431)

										gen13596 := Call(__e, ShenFunc(symshen_4lazyderef), gen13595, V2989)

										V2433 := gen13596
										gen13619 := Call(__e, ShenFunc(symcons_2), V2433)

										if True == gen13619 {
											gen13597 := Call(__e, ShenFunc(symhd), V2433)

											F := gen13597
											gen13598 := Call(__e, ShenFunc(symtl), V2433)

											gen13599 := Call(__e, ShenFunc(symshen_4lazyderef), gen13598, V2989)

											V2434 := gen13599
											gen13618 := Call(__e, ShenFunc(sym_a), Nil, V2434)

											if True == gen13618 {
												gen13600 := Call(__e, ShenFunc(symtl), V2430)

												gen13601 := Call(__e, ShenFunc(symshen_4lazyderef), gen13600, V2989)

												V2435 := gen13601
												gen13617 := Call(__e, ShenFunc(symcons_2), V2435)

												if True == gen13617 {
													gen13602 := Call(__e, ShenFunc(symhd), V2435)

													Action := gen13602
													gen13603 := Call(__e, ShenFunc(symtl), V2435)

													gen13604 := Call(__e, ShenFunc(symshen_4lazyderef), gen13603, V2989)

													V2436 := gen13604
													gen13616 := Call(__e, ShenFunc(sym_a), Nil, V2436)

													if True == gen13616 {
														gen13605 := Call(__e, ShenFunc(symtl), V2429)

														gen13606 := Call(__e, ShenFunc(symshen_4lazyderef), gen13605, V2989)

														V2437 := gen13606
														gen13615 := Call(__e, ShenFunc(sym_a), Nil, V2437)

														if True == gen13615 {
															Call(__e, ShenFunc(symshen_4incinfs))
															gen13614 := MakeNative(func(__e Evaluator) {
																gen13607 := Call(__e, ShenFunc(symcons), Action, Nil)

																gen13608 := Call(__e, ShenFunc(symcons), F, gen13607)

																gen13609 := Call(__e, ShenFunc(symcons), gen13608, Nil)

																gen13610 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen13609)

																gen13611 := Call(__e, ShenFunc(symcons), Action, Nil)

																gen13612 := Call(__e, ShenFunc(symcons), gen13610, gen13611)

																gen13613 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen13612)

																__e.TailApply(ShenFunc(symshen_4t_d_1action), gen13613, V2987, V2988, V2989, V2990)

																return

															}, 0)
															gen13626 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen13614)

														} else {
															gen13626 = False
														}

													} else {
														gen13626 = False
													}

												} else {
													gen13626 = False
												}

											} else {
												gen13626 = False
											}

										} else {
											gen13626 = False
										}

									} else {
										gen13626 = False
									}

								} else {
									gen13626 = False
								}

							} else {
								gen13626 = False
							}

						} else {
							gen13626 = False
						}

					} else {
						gen13626 = False
					}

				} else {
					gen13626 = False
				}
				Case := gen13626
				gen13655 := Call(__e, ShenFunc(sym_a), Case, False)

				if True == gen13655 {
					gen13627 := Call(__e, ShenFunc(symshen_4lazyderef), V2986, V2989)

					V2438 := gen13627
					gen13649 := Call(__e, ShenFunc(symcons_2), V2438)

					var gen13650 Obj
					if True == gen13649 {
						gen13628 := Call(__e, ShenFunc(symhd), V2438)

						gen13629 := Call(__e, ShenFunc(symshen_4lazyderef), gen13628, V2989)

						V2439 := gen13629
						gen13648 := Call(__e, ShenFunc(sym_a), MakeSymbol("shen.choicepoint!"), V2439)

						if True == gen13648 {
							gen13630 := Call(__e, ShenFunc(symtl), V2438)

							gen13631 := Call(__e, ShenFunc(symshen_4lazyderef), gen13630, V2989)

							V2440 := gen13631
							gen13647 := Call(__e, ShenFunc(symcons_2), V2440)

							if True == gen13647 {
								gen13632 := Call(__e, ShenFunc(symhd), V2440)

								Action := gen13632
								gen13633 := Call(__e, ShenFunc(symtl), V2440)

								gen13634 := Call(__e, ShenFunc(symshen_4lazyderef), gen13633, V2989)

								V2441 := gen13634
								gen13646 := Call(__e, ShenFunc(sym_a), Nil, V2441)

								if True == gen13646 {
									Call(__e, ShenFunc(symshen_4incinfs))
									gen13645 := MakeNative(func(__e Evaluator) {
										gen13635 := Call(__e, ShenFunc(symcons), Action, Nil)

										gen13636 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen13635)

										gen13637 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), Nil)

										gen13638 := Call(__e, ShenFunc(symcons), gen13637, Nil)

										gen13639 := Call(__e, ShenFunc(symcons), gen13636, gen13638)

										gen13640 := Call(__e, ShenFunc(symcons), gen13639, Nil)

										gen13641 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen13640)

										gen13642 := Call(__e, ShenFunc(symcons), Action, Nil)

										gen13643 := Call(__e, ShenFunc(symcons), gen13641, gen13642)

										gen13644 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen13643)

										__e.TailApply(ShenFunc(symshen_4t_d_1action), gen13644, V2987, V2988, V2989, V2990)

										return

									}, 0)
									gen13650 = Call(__e, ShenFunc(symcut), Throwcontrol, V2989, gen13645)

								} else {
									gen13650 = False
								}

							} else {
								gen13650 = False
							}

						} else {
							gen13650 = False
						}

					} else {
						gen13650 = False
					}
					Case := gen13650
					gen13654 := Call(__e, ShenFunc(sym_a), Case, False)

					if True == gen13654 {
						Call(__e, ShenFunc(symshen_4incinfs))
						gen13651 := Call(__e, ShenFunc(symcons), V2987, Nil)

						gen13652 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen13651)

						gen13653 := Call(__e, ShenFunc(symcons), V2986, gen13652)

						gen13657 = Call(__e, ShenFunc(symshen_4t_d), gen13653, V2988, V2989, V2990)

					} else {
						gen13657 = Case
					}

				} else {
					gen13657 = Case
				}

			} else {
				gen13657 = Case
			}
			__e.TailApply(ShenFunc(symshen_4cutpoint), Throwcontrol, gen13657)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.t*-action"), gen13658)

		gen13666 := MakeNative(func(__e Evaluator) {
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
			gen13659 := Call(__e, ShenFunc(symshen_4newpv), V2999)

			B := gen13659
			gen13660 := Call(__e, ShenFunc(symshen_4newpv), V2999)

			A := gen13660
			Call(__e, ShenFunc(symshen_4incinfs))
			gen13661 := Call(__e, ShenFunc(symgensym), MakeSymbol("shen.a"))

			gen13665 := MakeNative(func(__e Evaluator) {
				gen13662 := Call(__e, ShenFunc(symshen_4lazyderef), A, V2999)

				gen13663 := Call(__e, ShenFunc(symset), gen13662, Nil)

				gen13664 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(symshen_4findallhelp), V2996, V2997, V2998, A, V2999, V3000)

					return
				}, 0)
				__e.TailApply(ShenFunc(symbind), B, gen13663, V2999, gen13664)

				return

			}, 0)
			__e.TailApply(ShenFunc(symbind), A, gen13661, V2999, gen13665)

			return

		}, 5)
		Call(__e, ShenFunc(symdefun), MakeSymbol("findall"), gen13666)

		gen13673 := MakeNative(func(__e Evaluator) {
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
			Call(__e, ShenFunc(symshen_4incinfs))
			gen13668 := MakeNative(func(__e Evaluator) {
				gen13667 := MakeNative(func(__e Evaluator) {
					__e.TailApply(ShenFunc(symfwhen), False, V3011, V3012)

					return
				}, 0)
				__e.TailApply(ShenFunc(symshen_4remember), V3010, V3007, V3011, gen13667)

				return

			}, 0)
			gen13669 := Call(__e, ShenFunc(symcall), V3008, V3011, gen13668)

			Case := gen13669
			gen13672 := Call(__e, ShenFunc(sym_a), Case, False)

			if True == gen13672 {
				Call(__e, ShenFunc(symshen_4incinfs))
				gen13670 := Call(__e, ShenFunc(symshen_4lazyderef), V3010, V3011)

				gen13671 := Call(__e, ShenFunc(symvalue), gen13670)

				__e.TailApply(ShenFunc(symbind), V3009, gen13671, V3011, V3012)

				return

			} else {
				__e.Return(Case)
				return
			}

		}, 6)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.findallhelp"), gen13673)

		gen13681 := MakeNative(func(__e Evaluator) {
			V3017 := __e.Get(1)
			_ = V3017
			V3018 := __e.Get(2)
			_ = V3018
			V3019 := __e.Get(3)
			_ = V3019
			V3020 := __e.Get(4)
			_ = V3020
			gen13674 := Call(__e, ShenFunc(symshen_4newpv), V3019)

			B := gen13674
			Call(__e, ShenFunc(symshen_4incinfs))
			gen13675 := Call(__e, ShenFunc(symshen_4deref), V3017, V3019)

			gen13676 := Call(__e, ShenFunc(symshen_4deref), V3018, V3019)

			gen13677 := Call(__e, ShenFunc(symshen_4deref), V3017, V3019)

			gen13678 := Call(__e, ShenFunc(symvalue), gen13677)

			gen13679 := Call(__e, ShenFunc(symcons), gen13676, gen13678)

			gen13680 := Call(__e, ShenFunc(symset), gen13675, gen13679)

			__e.TailApply(ShenFunc(symbind), B, gen13680, V3019, V3020)

			return

		}, 4)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.remember"), gen13681)

		return

	}, 0))
}
