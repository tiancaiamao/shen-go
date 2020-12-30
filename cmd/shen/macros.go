package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen10082 := MakeNative(func(__e Evaluator) {
			V748 := __e.Get(1)
			_ = V748
			gen10078 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

			gen10079 := Call(__e, ShenFunc(symshen_4compose), gen10078, V748)

			Y := gen10079
			gen10081 := Call(__e, ShenFunc(sym_a), V748, Y)

			if True == gen10081 {
				__e.Return(V748)
				return
			} else {
				gen10080 := MakeNative(func(__e Evaluator) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(ShenFunc(symmacroexpand), Z)

					return
				}, 1)
				__e.TailApply(ShenFunc(symshen_4walk), gen10080, Y)

				return

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("macroexpand"), gen10082)

		gen10096 := MakeNative(func(__e Evaluator) {
			V750 := __e.Get(1)
			_ = V750
			gen10094 := Call(__e, ShenFunc(symcons_2), V750)

			var gen10095 Obj
			if True == gen10094 {
				gen10091 := Call(__e, ShenFunc(symhd), V750)

				gen10092 := Call(__e, ShenFunc(sym_a), MakeSymbol("error"), gen10091)

				var gen10093 Obj
				if True == gen10092 {
					gen10089 := Call(__e, ShenFunc(symtl), V750)

					gen10090 := Call(__e, ShenFunc(symcons_2), gen10089)

					if True == gen10090 {
						gen10093 = True
					} else {
						gen10093 = False
					}

				} else {
					gen10093 = False
				}
				if True == gen10093 {
					gen10095 = True
				} else {
					gen10095 = False
				}

			} else {
				gen10095 = False
			}
			if True == gen10095 {
				gen10083 := Call(__e, ShenFunc(symtl), V750)

				gen10084 := Call(__e, ShenFunc(symhd), gen10083)

				gen10085 := Call(__e, ShenFunc(symtl), V750)

				gen10086 := Call(__e, ShenFunc(symtl), gen10085)

				gen10087 := Call(__e, ShenFunc(symshen_4mkstr), gen10084, gen10086)

				gen10088 := Call(__e, ShenFunc(symcons), gen10087, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("simple-error"), gen10088)

				return

			} else {
				__e.Return(V750)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.error-macro"), gen10096)

		gen10128 := MakeNative(func(__e Evaluator) {
			V752 := __e.Get(1)
			_ = V752
			gen10126 := Call(__e, ShenFunc(symcons_2), V752)

			var gen10127 Obj
			if True == gen10126 {
				gen10123 := Call(__e, ShenFunc(symhd), V752)

				gen10124 := Call(__e, ShenFunc(sym_a), MakeSymbol("output"), gen10123)

				var gen10125 Obj
				if True == gen10124 {
					gen10121 := Call(__e, ShenFunc(symtl), V752)

					gen10122 := Call(__e, ShenFunc(symcons_2), gen10121)

					if True == gen10122 {
						gen10125 = True
					} else {
						gen10125 = False
					}

				} else {
					gen10125 = False
				}
				if True == gen10125 {
					gen10127 = True
				} else {
					gen10127 = False
				}

			} else {
				gen10127 = False
			}
			if True == gen10127 {
				gen10113 := Call(__e, ShenFunc(symtl), V752)

				gen10114 := Call(__e, ShenFunc(symhd), gen10113)

				gen10115 := Call(__e, ShenFunc(symtl), V752)

				gen10116 := Call(__e, ShenFunc(symtl), gen10115)

				gen10117 := Call(__e, ShenFunc(symshen_4mkstr), gen10114, gen10116)

				gen10118 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

				gen10119 := Call(__e, ShenFunc(symcons), gen10118, Nil)

				gen10120 := Call(__e, ShenFunc(symcons), gen10117, gen10119)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.prhush"), gen10120)

				return

			} else {
				gen10111 := Call(__e, ShenFunc(symcons_2), V752)

				var gen10112 Obj
				if True == gen10111 {
					gen10108 := Call(__e, ShenFunc(symhd), V752)

					gen10109 := Call(__e, ShenFunc(sym_a), MakeSymbol("pr"), gen10108)

					var gen10110 Obj
					if True == gen10109 {
						gen10105 := Call(__e, ShenFunc(symtl), V752)

						gen10106 := Call(__e, ShenFunc(symcons_2), gen10105)

						var gen10107 Obj
						if True == gen10106 {
							gen10102 := Call(__e, ShenFunc(symtl), V752)

							gen10103 := Call(__e, ShenFunc(symtl), gen10102)

							gen10104 := Call(__e, ShenFunc(sym_a), Nil, gen10103)

							if True == gen10104 {
								gen10107 = True
							} else {
								gen10107 = False
							}

						} else {
							gen10107 = False
						}
						if True == gen10107 {
							gen10110 = True
						} else {
							gen10110 = False
						}

					} else {
						gen10110 = False
					}
					if True == gen10110 {
						gen10112 = True
					} else {
						gen10112 = False
					}

				} else {
					gen10112 = False
				}
				if True == gen10112 {
					gen10097 := Call(__e, ShenFunc(symtl), V752)

					gen10098 := Call(__e, ShenFunc(symhd), gen10097)

					gen10099 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

					gen10100 := Call(__e, ShenFunc(symcons), gen10099, Nil)

					gen10101 := Call(__e, ShenFunc(symcons), gen10098, gen10100)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("pr"), gen10101)

					return

				} else {
					__e.Return(V752)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.output-macro"), gen10128)

		gen10140 := MakeNative(func(__e Evaluator) {
			V754 := __e.Get(1)
			_ = V754
			gen10138 := Call(__e, ShenFunc(symcons_2), V754)

			var gen10139 Obj
			if True == gen10138 {
				gen10135 := Call(__e, ShenFunc(symhd), V754)

				gen10136 := Call(__e, ShenFunc(sym_a), MakeSymbol("make-string"), gen10135)

				var gen10137 Obj
				if True == gen10136 {
					gen10133 := Call(__e, ShenFunc(symtl), V754)

					gen10134 := Call(__e, ShenFunc(symcons_2), gen10133)

					if True == gen10134 {
						gen10137 = True
					} else {
						gen10137 = False
					}

				} else {
					gen10137 = False
				}
				if True == gen10137 {
					gen10139 = True
				} else {
					gen10139 = False
				}

			} else {
				gen10139 = False
			}
			if True == gen10139 {
				gen10129 := Call(__e, ShenFunc(symtl), V754)

				gen10130 := Call(__e, ShenFunc(symhd), gen10129)

				gen10131 := Call(__e, ShenFunc(symtl), V754)

				gen10132 := Call(__e, ShenFunc(symtl), gen10131)

				__e.TailApply(ShenFunc(symshen_4mkstr), gen10130, gen10132)

				return

			} else {
				__e.Return(V754)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.make-string-macro"), gen10140)

		gen10193 := MakeNative(func(__e Evaluator) {
			V756 := __e.Get(1)
			_ = V756
			gen10191 := Call(__e, ShenFunc(symcons_2), V756)

			var gen10192 Obj
			if True == gen10191 {
				gen10188 := Call(__e, ShenFunc(symhd), V756)

				gen10189 := Call(__e, ShenFunc(sym_a), MakeSymbol("lineread"), gen10188)

				var gen10190 Obj
				if True == gen10189 {
					gen10186 := Call(__e, ShenFunc(symtl), V756)

					gen10187 := Call(__e, ShenFunc(sym_a), Nil, gen10186)

					if True == gen10187 {
						gen10190 = True
					} else {
						gen10190 = False
					}

				} else {
					gen10190 = False
				}
				if True == gen10190 {
					gen10192 = True
				} else {
					gen10192 = False
				}

			} else {
				gen10192 = False
			}
			if True == gen10192 {
				gen10184 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

				gen10185 := Call(__e, ShenFunc(symcons), gen10184, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("lineread"), gen10185)

				return

			} else {
				gen10182 := Call(__e, ShenFunc(symcons_2), V756)

				var gen10183 Obj
				if True == gen10182 {
					gen10179 := Call(__e, ShenFunc(symhd), V756)

					gen10180 := Call(__e, ShenFunc(sym_a), MakeSymbol("input"), gen10179)

					var gen10181 Obj
					if True == gen10180 {
						gen10177 := Call(__e, ShenFunc(symtl), V756)

						gen10178 := Call(__e, ShenFunc(sym_a), Nil, gen10177)

						if True == gen10178 {
							gen10181 = True
						} else {
							gen10181 = False
						}

					} else {
						gen10181 = False
					}
					if True == gen10181 {
						gen10183 = True
					} else {
						gen10183 = False
					}

				} else {
					gen10183 = False
				}
				if True == gen10183 {
					gen10175 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

					gen10176 := Call(__e, ShenFunc(symcons), gen10175, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("input"), gen10176)

					return

				} else {
					gen10173 := Call(__e, ShenFunc(symcons_2), V756)

					var gen10174 Obj
					if True == gen10173 {
						gen10170 := Call(__e, ShenFunc(symhd), V756)

						gen10171 := Call(__e, ShenFunc(sym_a), MakeSymbol("read"), gen10170)

						var gen10172 Obj
						if True == gen10171 {
							gen10168 := Call(__e, ShenFunc(symtl), V756)

							gen10169 := Call(__e, ShenFunc(sym_a), Nil, gen10168)

							if True == gen10169 {
								gen10172 = True
							} else {
								gen10172 = False
							}

						} else {
							gen10172 = False
						}
						if True == gen10172 {
							gen10174 = True
						} else {
							gen10174 = False
						}

					} else {
						gen10174 = False
					}
					if True == gen10174 {
						gen10166 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

						gen10167 := Call(__e, ShenFunc(symcons), gen10166, Nil)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("read"), gen10167)

						return

					} else {
						gen10164 := Call(__e, ShenFunc(symcons_2), V756)

						var gen10165 Obj
						if True == gen10164 {
							gen10161 := Call(__e, ShenFunc(symhd), V756)

							gen10162 := Call(__e, ShenFunc(sym_a), MakeSymbol("input+"), gen10161)

							var gen10163 Obj
							if True == gen10162 {
								gen10158 := Call(__e, ShenFunc(symtl), V756)

								gen10159 := Call(__e, ShenFunc(symcons_2), gen10158)

								var gen10160 Obj
								if True == gen10159 {
									gen10155 := Call(__e, ShenFunc(symtl), V756)

									gen10156 := Call(__e, ShenFunc(symtl), gen10155)

									gen10157 := Call(__e, ShenFunc(sym_a), Nil, gen10156)

									if True == gen10157 {
										gen10160 = True
									} else {
										gen10160 = False
									}

								} else {
									gen10160 = False
								}
								if True == gen10160 {
									gen10163 = True
								} else {
									gen10163 = False
								}

							} else {
								gen10163 = False
							}
							if True == gen10163 {
								gen10165 = True
							} else {
								gen10165 = False
							}

						} else {
							gen10165 = False
						}
						if True == gen10165 {
							gen10150 := Call(__e, ShenFunc(symtl), V756)

							gen10151 := Call(__e, ShenFunc(symhd), gen10150)

							gen10152 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

							gen10153 := Call(__e, ShenFunc(symcons), gen10152, Nil)

							gen10154 := Call(__e, ShenFunc(symcons), gen10151, gen10153)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("input+"), gen10154)

							return

						} else {
							gen10148 := Call(__e, ShenFunc(symcons_2), V756)

							var gen10149 Obj
							if True == gen10148 {
								gen10145 := Call(__e, ShenFunc(symhd), V756)

								gen10146 := Call(__e, ShenFunc(sym_a), MakeSymbol("read-byte"), gen10145)

								var gen10147 Obj
								if True == gen10146 {
									gen10143 := Call(__e, ShenFunc(symtl), V756)

									gen10144 := Call(__e, ShenFunc(sym_a), Nil, gen10143)

									if True == gen10144 {
										gen10147 = True
									} else {
										gen10147 = False
									}

								} else {
									gen10147 = False
								}
								if True == gen10147 {
									gen10149 = True
								} else {
									gen10149 = False
								}

							} else {
								gen10149 = False
							}
							if True == gen10149 {
								gen10141 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), Nil)

								gen10142 := Call(__e, ShenFunc(symcons), gen10141, Nil)

								__e.TailApply(ShenFunc(symcons), MakeSymbol("read-byte"), gen10142)

								return

							} else {
								__e.Return(V756)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.input-macro"), gen10193)

		gen10199 := MakeNative(func(__e Evaluator) {
			V759 := __e.Get(1)
			_ = V759
			V760 := __e.Get(2)
			_ = V760
			gen10198 := Call(__e, ShenFunc(sym_a), Nil, V759)

			if True == gen10198 {
				__e.Return(V760)
				return
			} else {
				gen10197 := Call(__e, ShenFunc(symcons_2), V759)

				if True == gen10197 {
					gen10194 := Call(__e, ShenFunc(symtl), V759)

					gen10195 := Call(__e, ShenFunc(symhd), V759)

					f33 := gen10195
					gen10196 := Call(__e, f33, V760)

					__e.TailApply(ShenFunc(symshen_4compose), gen10194, gen10196)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.compose"))

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compose"), gen10199)

		gen10238 := MakeNative(func(__e Evaluator) {
			V762 := __e.Get(1)
			_ = V762
			gen10236 := Call(__e, ShenFunc(symcons_2), V762)

			var gen10237 Obj
			if True == gen10236 {
				gen10233 := Call(__e, ShenFunc(symhd), V762)

				gen10234 := Call(__e, ShenFunc(sym_a), MakeSymbol("compile"), gen10233)

				var gen10235 Obj
				if True == gen10234 {
					gen10230 := Call(__e, ShenFunc(symtl), V762)

					gen10231 := Call(__e, ShenFunc(symcons_2), gen10230)

					var gen10232 Obj
					if True == gen10231 {
						gen10226 := Call(__e, ShenFunc(symtl), V762)

						gen10227 := Call(__e, ShenFunc(symtl), gen10226)

						gen10228 := Call(__e, ShenFunc(symcons_2), gen10227)

						var gen10229 Obj
						if True == gen10228 {
							gen10222 := Call(__e, ShenFunc(symtl), V762)

							gen10223 := Call(__e, ShenFunc(symtl), gen10222)

							gen10224 := Call(__e, ShenFunc(symtl), gen10223)

							gen10225 := Call(__e, ShenFunc(sym_a), Nil, gen10224)

							if True == gen10225 {
								gen10229 = True
							} else {
								gen10229 = False
							}

						} else {
							gen10229 = False
						}
						if True == gen10229 {
							gen10232 = True
						} else {
							gen10232 = False
						}

					} else {
						gen10232 = False
					}
					if True == gen10232 {
						gen10235 = True
					} else {
						gen10235 = False
					}

				} else {
					gen10235 = False
				}
				if True == gen10235 {
					gen10237 = True
				} else {
					gen10237 = False
				}

			} else {
				gen10237 = False
			}
			if True == gen10237 {
				gen10200 := Call(__e, ShenFunc(symtl), V762)

				gen10201 := Call(__e, ShenFunc(symhd), gen10200)

				gen10202 := Call(__e, ShenFunc(symtl), V762)

				gen10203 := Call(__e, ShenFunc(symtl), gen10202)

				gen10204 := Call(__e, ShenFunc(symhd), gen10203)

				gen10205 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen10206 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen10205)

				gen10207 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), Nil)

				gen10208 := Call(__e, ShenFunc(symcons), MakeString("parse error here: ~S~%"), gen10207)

				gen10209 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen10208)

				gen10210 := Call(__e, ShenFunc(symcons), MakeString("parse error~%"), Nil)

				gen10211 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen10210)

				gen10212 := Call(__e, ShenFunc(symcons), gen10211, Nil)

				gen10213 := Call(__e, ShenFunc(symcons), gen10209, gen10212)

				gen10214 := Call(__e, ShenFunc(symcons), gen10206, gen10213)

				gen10215 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen10214)

				gen10216 := Call(__e, ShenFunc(symcons), gen10215, Nil)

				gen10217 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen10216)

				gen10218 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen10217)

				gen10219 := Call(__e, ShenFunc(symcons), gen10218, Nil)

				gen10220 := Call(__e, ShenFunc(symcons), gen10204, gen10219)

				gen10221 := Call(__e, ShenFunc(symcons), gen10201, gen10220)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("compile"), gen10221)

				return

			} else {
				__e.Return(V762)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.compile-macro"), gen10238)

		gen10255 := MakeNative(func(__e Evaluator) {
			V764 := __e.Get(1)
			_ = V764
			gen10253 := Call(__e, ShenFunc(symcons_2), V764)

			var gen10254 Obj
			if True == gen10253 {
				gen10251 := Call(__e, ShenFunc(symhd), V764)

				gen10252 := Call(__e, ShenFunc(sym_a), MakeSymbol("prolog?"), gen10251)

				if True == gen10252 {
					gen10254 = True
				} else {
					gen10254 = False
				}

			} else {
				gen10254 = False
			}
			if True == gen10254 {
				gen10239 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.start-new-prolog-process"), Nil)

				gen10240 := Call(__e, ShenFunc(symtl), V764)

				gen10241 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), MakeSymbol("NPP"), gen10240)

				Calls := gen10241
				gen10242 := Call(__e, ShenFunc(symtl), V764)

				gen10243 := Call(__e, ShenFunc(symshen_4extract__vars), gen10242)

				Vs := gen10243
				gen10244 := Call(__e, ShenFunc(symtl), V764)

				gen10245 := Call(__e, ShenFunc(symshen_4externally_1bound), gen10244)

				External := gen10245
				gen10246 := Call(__e, ShenFunc(symdifference), Vs, External)

				PrologVs := gen10246
				gen10247 := Call(__e, ShenFunc(symshen_4locally_1bind_1prolog_1vs), MakeSymbol("NPP"), PrologVs, Calls)

				gen10248 := Call(__e, ShenFunc(symcons), gen10247, Nil)

				gen10249 := Call(__e, ShenFunc(symcons), gen10239, gen10248)

				gen10250 := Call(__e, ShenFunc(symcons), MakeSymbol("NPP"), gen10249)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen10250)

				return

			} else {
				__e.Return(V764)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.prolog-macro"), gen10255)

		gen10272 := MakeNative(func(__e Evaluator) {
			V770 := __e.Get(1)
			_ = V770
			gen10270 := Call(__e, ShenFunc(symcons_2), V770)

			var gen10271 Obj
			if True == gen10270 {
				gen10267 := Call(__e, ShenFunc(symhd), V770)

				gen10268 := Call(__e, ShenFunc(sym_a), MakeSymbol("receive"), gen10267)

				var gen10269 Obj
				if True == gen10268 {
					gen10264 := Call(__e, ShenFunc(symtl), V770)

					gen10265 := Call(__e, ShenFunc(symcons_2), gen10264)

					var gen10266 Obj
					if True == gen10265 {
						gen10261 := Call(__e, ShenFunc(symtl), V770)

						gen10262 := Call(__e, ShenFunc(symtl), gen10261)

						gen10263 := Call(__e, ShenFunc(sym_a), Nil, gen10262)

						if True == gen10263 {
							gen10266 = True
						} else {
							gen10266 = False
						}

					} else {
						gen10266 = False
					}
					if True == gen10266 {
						gen10269 = True
					} else {
						gen10269 = False
					}

				} else {
					gen10269 = False
				}
				if True == gen10269 {
					gen10271 = True
				} else {
					gen10271 = False
				}

			} else {
				gen10271 = False
			}
			if True == gen10271 {
				__e.TailApply(ShenFunc(symtl), V770)

				return
			} else {
				gen10260 := Call(__e, ShenFunc(symcons_2), V770)

				if True == gen10260 {
					gen10256 := Call(__e, ShenFunc(symhd), V770)

					gen10257 := Call(__e, ShenFunc(symshen_4externally_1bound), gen10256)

					gen10258 := Call(__e, ShenFunc(symtl), V770)

					gen10259 := Call(__e, ShenFunc(symshen_4externally_1bound), gen10258)

					__e.TailApply(ShenFunc(symunion), gen10257, gen10259)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.externally-bound"), gen10272)

		gen10283 := MakeNative(func(__e Evaluator) {
			V788 := __e.Get(1)
			_ = V788
			V789 := __e.Get(2)
			_ = V789
			V790 := __e.Get(3)
			_ = V790
			gen10282 := Call(__e, ShenFunc(sym_a), Nil, V789)

			if True == gen10282 {
				__e.Return(V790)
				return
			} else {
				gen10281 := Call(__e, ShenFunc(symcons_2), V789)

				if True == gen10281 {
					gen10273 := Call(__e, ShenFunc(symhd), V789)

					gen10274 := Call(__e, ShenFunc(symcons), V788, Nil)

					gen10275 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.newpv"), gen10274)

					gen10276 := Call(__e, ShenFunc(symtl), V789)

					gen10277 := Call(__e, ShenFunc(symshen_4locally_1bind_1prolog_1vs), V788, gen10276, V790)

					gen10278 := Call(__e, ShenFunc(symcons), gen10277, Nil)

					gen10279 := Call(__e, ShenFunc(symcons), gen10275, gen10278)

					gen10280 := Call(__e, ShenFunc(symcons), gen10273, gen10279)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen10280)

					return

				} else {
					__e.TailApply(ShenFunc(symsimple_1error), MakeString("implementation error inp locally-bind-prolog-vs"))

					return
				}

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.locally-bind-prolog-vs"), gen10283)

		gen10456 := MakeNative(func(__e Evaluator) {
			V803 := __e.Get(1)
			_ = V803
			V804 := __e.Get(2)
			_ = V804
			gen10455 := Call(__e, ShenFunc(sym_a), Nil, V804)

			if True == gen10455 {
				__e.Return(True)
				return
			} else {
				gen10453 := Call(__e, ShenFunc(symcons_2), V804)

				var gen10454 Obj
				if True == gen10453 {
					gen10451 := Call(__e, ShenFunc(symhd), V804)

					gen10452 := Call(__e, ShenFunc(sym_a), MakeSymbol("!"), gen10451)

					if True == gen10452 {
						gen10454 = True
					} else {
						gen10454 = False
					}

				} else {
					gen10454 = False
				}
				if True == gen10454 {
					gen10444 := Call(__e, ShenFunc(symtl), V804)

					gen10445 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10444)

					gen10446 := Call(__e, ShenFunc(symcons), gen10445, Nil)

					gen10447 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10446)

					gen10448 := Call(__e, ShenFunc(symcons), gen10447, Nil)

					gen10449 := Call(__e, ShenFunc(symcons), V803, gen10448)

					gen10450 := Call(__e, ShenFunc(symcons), False, gen10449)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("cut"), gen10450)

					return

				} else {
					gen10442 := Call(__e, ShenFunc(symcons_2), V804)

					var gen10443 Obj
					if True == gen10442 {
						gen10439 := Call(__e, ShenFunc(symhd), V804)

						gen10440 := Call(__e, ShenFunc(symcons_2), gen10439)

						var gen10441 Obj
						if True == gen10440 {
							gen10435 := Call(__e, ShenFunc(symhd), V804)

							gen10436 := Call(__e, ShenFunc(symhd), gen10435)

							gen10437 := Call(__e, ShenFunc(sym_a), MakeSymbol("when"), gen10436)

							var gen10438 Obj
							if True == gen10437 {
								gen10431 := Call(__e, ShenFunc(symhd), V804)

								gen10432 := Call(__e, ShenFunc(symtl), gen10431)

								gen10433 := Call(__e, ShenFunc(symcons_2), gen10432)

								var gen10434 Obj
								if True == gen10433 {
									gen10427 := Call(__e, ShenFunc(symhd), V804)

									gen10428 := Call(__e, ShenFunc(symtl), gen10427)

									gen10429 := Call(__e, ShenFunc(symtl), gen10428)

									gen10430 := Call(__e, ShenFunc(sym_a), Nil, gen10429)

									if True == gen10430 {
										gen10434 = True
									} else {
										gen10434 = False
									}

								} else {
									gen10434 = False
								}
								if True == gen10434 {
									gen10438 = True
								} else {
									gen10438 = False
								}

							} else {
								gen10438 = False
							}
							if True == gen10438 {
								gen10441 = True
							} else {
								gen10441 = False
							}

						} else {
							gen10441 = False
						}
						if True == gen10441 {
							gen10443 = True
						} else {
							gen10443 = False
						}

					} else {
						gen10443 = False
					}
					if True == gen10443 {
						gen10416 := Call(__e, ShenFunc(symhd), V804)

						gen10417 := Call(__e, ShenFunc(symtl), gen10416)

						gen10418 := Call(__e, ShenFunc(symhd), gen10417)

						gen10419 := Call(__e, ShenFunc(symshen_4insert_1deref), gen10418, V803)

						gen10420 := Call(__e, ShenFunc(symtl), V804)

						gen10421 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10420)

						gen10422 := Call(__e, ShenFunc(symcons), gen10421, Nil)

						gen10423 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10422)

						gen10424 := Call(__e, ShenFunc(symcons), gen10423, Nil)

						gen10425 := Call(__e, ShenFunc(symcons), V803, gen10424)

						gen10426 := Call(__e, ShenFunc(symcons), gen10419, gen10425)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen10426)

						return

					} else {
						gen10414 := Call(__e, ShenFunc(symcons_2), V804)

						var gen10415 Obj
						if True == gen10414 {
							gen10411 := Call(__e, ShenFunc(symhd), V804)

							gen10412 := Call(__e, ShenFunc(symcons_2), gen10411)

							var gen10413 Obj
							if True == gen10412 {
								gen10407 := Call(__e, ShenFunc(symhd), V804)

								gen10408 := Call(__e, ShenFunc(symhd), gen10407)

								gen10409 := Call(__e, ShenFunc(sym_a), MakeSymbol("is"), gen10408)

								var gen10410 Obj
								if True == gen10409 {
									gen10403 := Call(__e, ShenFunc(symhd), V804)

									gen10404 := Call(__e, ShenFunc(symtl), gen10403)

									gen10405 := Call(__e, ShenFunc(symcons_2), gen10404)

									var gen10406 Obj
									if True == gen10405 {
										gen10398 := Call(__e, ShenFunc(symhd), V804)

										gen10399 := Call(__e, ShenFunc(symtl), gen10398)

										gen10400 := Call(__e, ShenFunc(symtl), gen10399)

										gen10401 := Call(__e, ShenFunc(symcons_2), gen10400)

										var gen10402 Obj
										if True == gen10401 {
											gen10393 := Call(__e, ShenFunc(symhd), V804)

											gen10394 := Call(__e, ShenFunc(symtl), gen10393)

											gen10395 := Call(__e, ShenFunc(symtl), gen10394)

											gen10396 := Call(__e, ShenFunc(symtl), gen10395)

											gen10397 := Call(__e, ShenFunc(sym_a), Nil, gen10396)

											if True == gen10397 {
												gen10402 = True
											} else {
												gen10402 = False
											}

										} else {
											gen10402 = False
										}
										if True == gen10402 {
											gen10406 = True
										} else {
											gen10406 = False
										}

									} else {
										gen10406 = False
									}
									if True == gen10406 {
										gen10410 = True
									} else {
										gen10410 = False
									}

								} else {
									gen10410 = False
								}
								if True == gen10410 {
									gen10413 = True
								} else {
									gen10413 = False
								}

							} else {
								gen10413 = False
							}
							if True == gen10413 {
								gen10415 = True
							} else {
								gen10415 = False
							}

						} else {
							gen10415 = False
						}
						if True == gen10415 {
							gen10377 := Call(__e, ShenFunc(symhd), V804)

							gen10378 := Call(__e, ShenFunc(symtl), gen10377)

							gen10379 := Call(__e, ShenFunc(symhd), gen10378)

							gen10380 := Call(__e, ShenFunc(symhd), V804)

							gen10381 := Call(__e, ShenFunc(symtl), gen10380)

							gen10382 := Call(__e, ShenFunc(symtl), gen10381)

							gen10383 := Call(__e, ShenFunc(symhd), gen10382)

							gen10384 := Call(__e, ShenFunc(symshen_4insert_1deref), gen10383, V803)

							gen10385 := Call(__e, ShenFunc(symtl), V804)

							gen10386 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10385)

							gen10387 := Call(__e, ShenFunc(symcons), gen10386, Nil)

							gen10388 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10387)

							gen10389 := Call(__e, ShenFunc(symcons), gen10388, Nil)

							gen10390 := Call(__e, ShenFunc(symcons), V803, gen10389)

							gen10391 := Call(__e, ShenFunc(symcons), gen10384, gen10390)

							gen10392 := Call(__e, ShenFunc(symcons), gen10379, gen10391)

							__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen10392)

							return

						} else {
							gen10375 := Call(__e, ShenFunc(symcons_2), V804)

							var gen10376 Obj
							if True == gen10375 {
								gen10372 := Call(__e, ShenFunc(symhd), V804)

								gen10373 := Call(__e, ShenFunc(symcons_2), gen10372)

								var gen10374 Obj
								if True == gen10373 {
									gen10368 := Call(__e, ShenFunc(symhd), V804)

									gen10369 := Call(__e, ShenFunc(symhd), gen10368)

									gen10370 := Call(__e, ShenFunc(sym_a), MakeSymbol("receive"), gen10369)

									var gen10371 Obj
									if True == gen10370 {
										gen10364 := Call(__e, ShenFunc(symhd), V804)

										gen10365 := Call(__e, ShenFunc(symtl), gen10364)

										gen10366 := Call(__e, ShenFunc(symcons_2), gen10365)

										var gen10367 Obj
										if True == gen10366 {
											gen10360 := Call(__e, ShenFunc(symhd), V804)

											gen10361 := Call(__e, ShenFunc(symtl), gen10360)

											gen10362 := Call(__e, ShenFunc(symtl), gen10361)

											gen10363 := Call(__e, ShenFunc(sym_a), Nil, gen10362)

											if True == gen10363 {
												gen10367 = True
											} else {
												gen10367 = False
											}

										} else {
											gen10367 = False
										}
										if True == gen10367 {
											gen10371 = True
										} else {
											gen10371 = False
										}

									} else {
										gen10371 = False
									}
									if True == gen10371 {
										gen10374 = True
									} else {
										gen10374 = False
									}

								} else {
									gen10374 = False
								}
								if True == gen10374 {
									gen10376 = True
								} else {
									gen10376 = False
								}

							} else {
								gen10376 = False
							}
							if True == gen10376 {
								gen10359 := Call(__e, ShenFunc(symtl), V804)

								__e.TailApply(ShenFunc(symshen_4bld_1prolog_1call), V803, gen10359)

								return

							} else {
								gen10357 := Call(__e, ShenFunc(symcons_2), V804)

								var gen10358 Obj
								if True == gen10357 {
									gen10354 := Call(__e, ShenFunc(symhd), V804)

									gen10355 := Call(__e, ShenFunc(symcons_2), gen10354)

									var gen10356 Obj
									if True == gen10355 {
										gen10350 := Call(__e, ShenFunc(symhd), V804)

										gen10351 := Call(__e, ShenFunc(symhd), gen10350)

										gen10352 := Call(__e, ShenFunc(sym_a), MakeSymbol("bind"), gen10351)

										var gen10353 Obj
										if True == gen10352 {
											gen10346 := Call(__e, ShenFunc(symhd), V804)

											gen10347 := Call(__e, ShenFunc(symtl), gen10346)

											gen10348 := Call(__e, ShenFunc(symcons_2), gen10347)

											var gen10349 Obj
											if True == gen10348 {
												gen10341 := Call(__e, ShenFunc(symhd), V804)

												gen10342 := Call(__e, ShenFunc(symtl), gen10341)

												gen10343 := Call(__e, ShenFunc(symtl), gen10342)

												gen10344 := Call(__e, ShenFunc(symcons_2), gen10343)

												var gen10345 Obj
												if True == gen10344 {
													gen10336 := Call(__e, ShenFunc(symhd), V804)

													gen10337 := Call(__e, ShenFunc(symtl), gen10336)

													gen10338 := Call(__e, ShenFunc(symtl), gen10337)

													gen10339 := Call(__e, ShenFunc(symtl), gen10338)

													gen10340 := Call(__e, ShenFunc(sym_a), Nil, gen10339)

													if True == gen10340 {
														gen10345 = True
													} else {
														gen10345 = False
													}

												} else {
													gen10345 = False
												}
												if True == gen10345 {
													gen10349 = True
												} else {
													gen10349 = False
												}

											} else {
												gen10349 = False
											}
											if True == gen10349 {
												gen10353 = True
											} else {
												gen10353 = False
											}

										} else {
											gen10353 = False
										}
										if True == gen10353 {
											gen10356 = True
										} else {
											gen10356 = False
										}

									} else {
										gen10356 = False
									}
									if True == gen10356 {
										gen10358 = True
									} else {
										gen10358 = False
									}

								} else {
									gen10358 = False
								}
								if True == gen10358 {
									gen10320 := Call(__e, ShenFunc(symhd), V804)

									gen10321 := Call(__e, ShenFunc(symtl), gen10320)

									gen10322 := Call(__e, ShenFunc(symhd), gen10321)

									gen10323 := Call(__e, ShenFunc(symhd), V804)

									gen10324 := Call(__e, ShenFunc(symtl), gen10323)

									gen10325 := Call(__e, ShenFunc(symtl), gen10324)

									gen10326 := Call(__e, ShenFunc(symhd), gen10325)

									gen10327 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen10326, V803)

									gen10328 := Call(__e, ShenFunc(symtl), V804)

									gen10329 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10328)

									gen10330 := Call(__e, ShenFunc(symcons), gen10329, Nil)

									gen10331 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10330)

									gen10332 := Call(__e, ShenFunc(symcons), gen10331, Nil)

									gen10333 := Call(__e, ShenFunc(symcons), V803, gen10332)

									gen10334 := Call(__e, ShenFunc(symcons), gen10327, gen10333)

									gen10335 := Call(__e, ShenFunc(symcons), gen10322, gen10334)

									__e.TailApply(ShenFunc(symcons), MakeSymbol("bind"), gen10335)

									return

								} else {
									gen10318 := Call(__e, ShenFunc(symcons_2), V804)

									var gen10319 Obj
									if True == gen10318 {
										gen10315 := Call(__e, ShenFunc(symhd), V804)

										gen10316 := Call(__e, ShenFunc(symcons_2), gen10315)

										var gen10317 Obj
										if True == gen10316 {
											gen10311 := Call(__e, ShenFunc(symhd), V804)

											gen10312 := Call(__e, ShenFunc(symhd), gen10311)

											gen10313 := Call(__e, ShenFunc(sym_a), MakeSymbol("fwhen"), gen10312)

											var gen10314 Obj
											if True == gen10313 {
												gen10307 := Call(__e, ShenFunc(symhd), V804)

												gen10308 := Call(__e, ShenFunc(symtl), gen10307)

												gen10309 := Call(__e, ShenFunc(symcons_2), gen10308)

												var gen10310 Obj
												if True == gen10309 {
													gen10303 := Call(__e, ShenFunc(symhd), V804)

													gen10304 := Call(__e, ShenFunc(symtl), gen10303)

													gen10305 := Call(__e, ShenFunc(symtl), gen10304)

													gen10306 := Call(__e, ShenFunc(sym_a), Nil, gen10305)

													if True == gen10306 {
														gen10310 = True
													} else {
														gen10310 = False
													}

												} else {
													gen10310 = False
												}
												if True == gen10310 {
													gen10314 = True
												} else {
													gen10314 = False
												}

											} else {
												gen10314 = False
											}
											if True == gen10314 {
												gen10317 = True
											} else {
												gen10317 = False
											}

										} else {
											gen10317 = False
										}
										if True == gen10317 {
											gen10319 = True
										} else {
											gen10319 = False
										}

									} else {
										gen10319 = False
									}
									if True == gen10319 {
										gen10292 := Call(__e, ShenFunc(symhd), V804)

										gen10293 := Call(__e, ShenFunc(symtl), gen10292)

										gen10294 := Call(__e, ShenFunc(symhd), gen10293)

										gen10295 := Call(__e, ShenFunc(symshen_4insert_1lazyderef), gen10294, V803)

										gen10296 := Call(__e, ShenFunc(symtl), V804)

										gen10297 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10296)

										gen10298 := Call(__e, ShenFunc(symcons), gen10297, Nil)

										gen10299 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10298)

										gen10300 := Call(__e, ShenFunc(symcons), gen10299, Nil)

										gen10301 := Call(__e, ShenFunc(symcons), V803, gen10300)

										gen10302 := Call(__e, ShenFunc(symcons), gen10295, gen10301)

										__e.TailApply(ShenFunc(symcons), MakeSymbol("fwhen"), gen10302)

										return

									} else {
										gen10291 := Call(__e, ShenFunc(symcons_2), V804)

										if True == gen10291 {
											gen10284 := Call(__e, ShenFunc(symhd), V804)

											gen10285 := Call(__e, ShenFunc(symtl), V804)

											gen10286 := Call(__e, ShenFunc(symshen_4bld_1prolog_1call), V803, gen10285)

											gen10287 := Call(__e, ShenFunc(symcons), gen10286, Nil)

											gen10288 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen10287)

											gen10289 := Call(__e, ShenFunc(symcons), gen10288, Nil)

											gen10290 := Call(__e, ShenFunc(symcons), V803, gen10289)

											__e.TailApply(ShenFunc(symappend), gen10284, gen10290)

											return

										} else {
											__e.TailApply(ShenFunc(symsimple_1error), MakeString("implementation error in bld-prolog-call"))

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
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.bld-prolog-call"), gen10456)

		gen10469 := MakeNative(func(__e Evaluator) {
			V806 := __e.Get(1)
			_ = V806
			gen10467 := Call(__e, ShenFunc(symcons_2), V806)

			var gen10468 Obj
			if True == gen10467 {
				gen10464 := Call(__e, ShenFunc(symhd), V806)

				gen10465 := Call(__e, ShenFunc(sym_a), MakeSymbol("defprolog"), gen10464)

				var gen10466 Obj
				if True == gen10465 {
					gen10462 := Call(__e, ShenFunc(symtl), V806)

					gen10463 := Call(__e, ShenFunc(symcons_2), gen10462)

					if True == gen10463 {
						gen10466 = True
					} else {
						gen10466 = False
					}

				} else {
					gen10466 = False
				}
				if True == gen10466 {
					gen10468 = True
				} else {
					gen10468 = False
				}

			} else {
				gen10468 = False
			}
			if True == gen10468 {
				gen10457 := MakeNative(func(__e Evaluator) {
					Y := __e.Get(1)
					_ = Y
					__e.TailApply(ShenFunc(symshen_4_5defprolog_6), Y)

					return
				}, 1)
				gen10458 := Call(__e, ShenFunc(symtl), V806)

				gen10461 := MakeNative(func(__e Evaluator) {
					Y := __e.Get(1)
					_ = Y
					gen10459 := Call(__e, ShenFunc(symtl), V806)

					gen10460 := Call(__e, ShenFunc(symhd), gen10459)

					__e.TailApply(ShenFunc(symshen_4prolog_1error), gen10460, Y)

					return

				}, 1)
				__e.TailApply(ShenFunc(symcompile), gen10457, gen10458, gen10461)

				return

			} else {
				__e.Return(V806)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.defprolog-macro"), gen10469)

		gen10496 := MakeNative(func(__e Evaluator) {
			V808 := __e.Get(1)
			_ = V808
			gen10494 := Call(__e, ShenFunc(symcons_2), V808)

			var gen10495 Obj
			if True == gen10494 {
				gen10491 := Call(__e, ShenFunc(symhd), V808)

				gen10492 := Call(__e, ShenFunc(sym_a), MakeSymbol("datatype"), gen10491)

				var gen10493 Obj
				if True == gen10492 {
					gen10489 := Call(__e, ShenFunc(symtl), V808)

					gen10490 := Call(__e, ShenFunc(symcons_2), gen10489)

					if True == gen10490 {
						gen10493 = True
					} else {
						gen10493 = False
					}

				} else {
					gen10493 = False
				}
				if True == gen10493 {
					gen10495 = True
				} else {
					gen10495 = False
				}

			} else {
				gen10495 = False
			}
			if True == gen10495 {
				gen10470 := Call(__e, ShenFunc(symtl), V808)

				gen10471 := Call(__e, ShenFunc(symhd), gen10470)

				gen10472 := Call(__e, ShenFunc(symshen_4intern_1type), gen10471)

				gen10473 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), Nil)

				gen10474 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.<datatype-rules>"), gen10473)

				gen10475 := Call(__e, ShenFunc(symcons), gen10474, Nil)

				gen10476 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen10475)

				gen10477 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen10476)

				gen10478 := Call(__e, ShenFunc(symtl), V808)

				gen10479 := Call(__e, ShenFunc(symtl), gen10478)

				gen10480 := Call(__e, ShenFunc(symshen_4rcons__form), gen10479)

				gen10481 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-error"), Nil)

				gen10482 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen10481)

				gen10483 := Call(__e, ShenFunc(symcons), gen10482, Nil)

				gen10484 := Call(__e, ShenFunc(symcons), gen10480, gen10483)

				gen10485 := Call(__e, ShenFunc(symcons), gen10477, gen10484)

				gen10486 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen10485)

				gen10487 := Call(__e, ShenFunc(symcons), gen10486, Nil)

				gen10488 := Call(__e, ShenFunc(symcons), gen10472, gen10487)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.process-datatype"), gen10488)

				return

			} else {
				__e.Return(V808)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.datatype-macro"), gen10496)

		gen10499 := MakeNative(func(__e Evaluator) {
			V810 := __e.Get(1)
			_ = V810
			gen10497 := Call(__e, ShenFunc(symstr), V810)

			gen10498 := Call(__e, ShenFunc(symcn), gen10497, MakeString("#type"))

			__e.TailApply(ShenFunc(symintern), gen10498)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.intern-type"), gen10499)

		gen10553 := MakeNative(func(__e Evaluator) {
			V812 := __e.Get(1)
			_ = V812
			gen10551 := Call(__e, ShenFunc(symcons_2), V812)

			var gen10552 Obj
			if True == gen10551 {
				gen10548 := Call(__e, ShenFunc(symhd), V812)

				gen10549 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen10548)

				var gen10550 Obj
				if True == gen10549 {
					gen10545 := Call(__e, ShenFunc(symtl), V812)

					gen10546 := Call(__e, ShenFunc(symcons_2), gen10545)

					var gen10547 Obj
					if True == gen10546 {
						gen10541 := Call(__e, ShenFunc(symtl), V812)

						gen10542 := Call(__e, ShenFunc(symtl), gen10541)

						gen10543 := Call(__e, ShenFunc(symcons_2), gen10542)

						var gen10544 Obj
						if True == gen10543 {
							gen10537 := Call(__e, ShenFunc(symtl), V812)

							gen10538 := Call(__e, ShenFunc(symtl), gen10537)

							gen10539 := Call(__e, ShenFunc(symtl), gen10538)

							gen10540 := Call(__e, ShenFunc(symcons_2), gen10539)

							if True == gen10540 {
								gen10544 = True
							} else {
								gen10544 = False
							}

						} else {
							gen10544 = False
						}
						if True == gen10544 {
							gen10547 = True
						} else {
							gen10547 = False
						}

					} else {
						gen10547 = False
					}
					if True == gen10547 {
						gen10550 = True
					} else {
						gen10550 = False
					}

				} else {
					gen10550 = False
				}
				if True == gen10550 {
					gen10552 = True
				} else {
					gen10552 = False
				}

			} else {
				gen10552 = False
			}
			if True == gen10552 {
				gen10529 := Call(__e, ShenFunc(symtl), V812)

				gen10530 := Call(__e, ShenFunc(symhd), gen10529)

				gen10531 := Call(__e, ShenFunc(symtl), V812)

				gen10532 := Call(__e, ShenFunc(symtl), gen10531)

				gen10533 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen10532)

				gen10534 := Call(__e, ShenFunc(symshen_4_8s_1macro), gen10533)

				gen10535 := Call(__e, ShenFunc(symcons), gen10534, Nil)

				gen10536 := Call(__e, ShenFunc(symcons), gen10530, gen10535)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("@s"), gen10536)

				return

			} else {
				gen10527 := Call(__e, ShenFunc(symcons_2), V812)

				var gen10528 Obj
				if True == gen10527 {
					gen10524 := Call(__e, ShenFunc(symhd), V812)

					gen10525 := Call(__e, ShenFunc(sym_a), MakeSymbol("@s"), gen10524)

					var gen10526 Obj
					if True == gen10525 {
						gen10521 := Call(__e, ShenFunc(symtl), V812)

						gen10522 := Call(__e, ShenFunc(symcons_2), gen10521)

						var gen10523 Obj
						if True == gen10522 {
							gen10517 := Call(__e, ShenFunc(symtl), V812)

							gen10518 := Call(__e, ShenFunc(symtl), gen10517)

							gen10519 := Call(__e, ShenFunc(symcons_2), gen10518)

							var gen10520 Obj
							if True == gen10519 {
								gen10512 := Call(__e, ShenFunc(symtl), V812)

								gen10513 := Call(__e, ShenFunc(symtl), gen10512)

								gen10514 := Call(__e, ShenFunc(symtl), gen10513)

								gen10515 := Call(__e, ShenFunc(sym_a), Nil, gen10514)

								var gen10516 Obj
								if True == gen10515 {
									gen10509 := Call(__e, ShenFunc(symtl), V812)

									gen10510 := Call(__e, ShenFunc(symhd), gen10509)

									gen10511 := Call(__e, ShenFunc(symstring_2), gen10510)

									if True == gen10511 {
										gen10516 = True
									} else {
										gen10516 = False
									}

								} else {
									gen10516 = False
								}
								if True == gen10516 {
									gen10520 = True
								} else {
									gen10520 = False
								}

							} else {
								gen10520 = False
							}
							if True == gen10520 {
								gen10523 = True
							} else {
								gen10523 = False
							}

						} else {
							gen10523 = False
						}
						if True == gen10523 {
							gen10526 = True
						} else {
							gen10526 = False
						}

					} else {
						gen10526 = False
					}
					if True == gen10526 {
						gen10528 = True
					} else {
						gen10528 = False
					}

				} else {
					gen10528 = False
				}
				if True == gen10528 {
					gen10500 := Call(__e, ShenFunc(symtl), V812)

					gen10501 := Call(__e, ShenFunc(symhd), gen10500)

					gen10502 := Call(__e, ShenFunc(symexplode), gen10501)

					E := gen10502
					gen10507 := Call(__e, ShenFunc(symlength), E)

					gen10508 := Call(__e, ShenFunc(sym_6), gen10507, MakeNumber(1))

					if True == gen10508 {
						gen10503 := Call(__e, ShenFunc(symtl), V812)

						gen10504 := Call(__e, ShenFunc(symtl), gen10503)

						gen10505 := Call(__e, ShenFunc(symappend), E, gen10504)

						gen10506 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen10505)

						__e.TailApply(ShenFunc(symshen_4_8s_1macro), gen10506)

						return

					} else {
						__e.Return(V812)
						return
					}

				} else {
					__e.Return(V812)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.@s-macro"), gen10553)

		gen10562 := MakeNative(func(__e Evaluator) {
			V814 := __e.Get(1)
			_ = V814
			gen10560 := Call(__e, ShenFunc(symcons_2), V814)

			var gen10561 Obj
			if True == gen10560 {
				gen10558 := Call(__e, ShenFunc(symhd), V814)

				gen10559 := Call(__e, ShenFunc(sym_a), MakeSymbol("synonyms"), gen10558)

				if True == gen10559 {
					gen10561 = True
				} else {
					gen10561 = False
				}

			} else {
				gen10561 = False
			}
			if True == gen10561 {
				gen10554 := Call(__e, ShenFunc(symtl), V814)

				gen10555 := Call(__e, ShenFunc(symshen_4curry_1synonyms), gen10554)

				gen10556 := Call(__e, ShenFunc(symshen_4rcons__form), gen10555)

				gen10557 := Call(__e, ShenFunc(symcons), gen10556, Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("shen.synonyms-help"), gen10557)

				return

			} else {
				__e.Return(V814)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.synonyms-macro"), gen10562)

		gen10564 := MakeNative(func(__e Evaluator) {
			V816 := __e.Get(1)
			_ = V816
			gen10563 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4curry_1type), X)

				return
			}, 1)
			__e.TailApply(ShenFunc(symmap), gen10563, V816)

			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.curry-synonyms"), gen10564)

		gen10573 := MakeNative(func(__e Evaluator) {
			V818 := __e.Get(1)
			_ = V818
			gen10571 := Call(__e, ShenFunc(symcons_2), V818)

			var gen10572 Obj
			if True == gen10571 {
				gen10568 := Call(__e, ShenFunc(symhd), V818)

				gen10569 := Call(__e, ShenFunc(sym_a), MakeSymbol("nl"), gen10568)

				var gen10570 Obj
				if True == gen10569 {
					gen10566 := Call(__e, ShenFunc(symtl), V818)

					gen10567 := Call(__e, ShenFunc(sym_a), Nil, gen10566)

					if True == gen10567 {
						gen10570 = True
					} else {
						gen10570 = False
					}

				} else {
					gen10570 = False
				}
				if True == gen10570 {
					gen10572 = True
				} else {
					gen10572 = False
				}

			} else {
				gen10572 = False
			}
			if True == gen10572 {
				gen10565 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("nl"), gen10565)

				return

			} else {
				__e.Return(V818)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.nl-macro"), gen10573)

		gen10608 := MakeNative(func(__e Evaluator) {
			V820 := __e.Get(1)
			_ = V820
			gen10606 := Call(__e, ShenFunc(symcons_2), V820)

			var gen10607 Obj
			if True == gen10606 {
				gen10603 := Call(__e, ShenFunc(symtl), V820)

				gen10604 := Call(__e, ShenFunc(symcons_2), gen10603)

				var gen10605 Obj
				if True == gen10604 {
					gen10599 := Call(__e, ShenFunc(symtl), V820)

					gen10600 := Call(__e, ShenFunc(symtl), gen10599)

					gen10601 := Call(__e, ShenFunc(symcons_2), gen10600)

					var gen10602 Obj
					if True == gen10601 {
						gen10594 := Call(__e, ShenFunc(symtl), V820)

						gen10595 := Call(__e, ShenFunc(symtl), gen10594)

						gen10596 := Call(__e, ShenFunc(symtl), gen10595)

						gen10597 := Call(__e, ShenFunc(symcons_2), gen10596)

						var gen10598 Obj
						if True == gen10597 {
							gen10584 := Call(__e, ShenFunc(symhd), V820)

							gen10585 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), Nil)

							gen10586 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen10585)

							gen10587 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen10586)

							gen10588 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen10587)

							gen10589 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen10588)

							gen10590 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen10589)

							gen10591 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen10590)

							gen10592 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen10591)

							gen10593 := Call(__e, ShenFunc(symelement_2), gen10584, gen10592)

							if True == gen10593 {
								gen10598 = True
							} else {
								gen10598 = False
							}

						} else {
							gen10598 = False
						}
						if True == gen10598 {
							gen10602 = True
						} else {
							gen10602 = False
						}

					} else {
						gen10602 = False
					}
					if True == gen10602 {
						gen10605 = True
					} else {
						gen10605 = False
					}

				} else {
					gen10605 = False
				}
				if True == gen10605 {
					gen10607 = True
				} else {
					gen10607 = False
				}

			} else {
				gen10607 = False
			}
			if True == gen10607 {
				gen10574 := Call(__e, ShenFunc(symhd), V820)

				gen10575 := Call(__e, ShenFunc(symtl), V820)

				gen10576 := Call(__e, ShenFunc(symhd), gen10575)

				gen10577 := Call(__e, ShenFunc(symhd), V820)

				gen10578 := Call(__e, ShenFunc(symtl), V820)

				gen10579 := Call(__e, ShenFunc(symtl), gen10578)

				gen10580 := Call(__e, ShenFunc(symcons), gen10577, gen10579)

				gen10581 := Call(__e, ShenFunc(symshen_4assoc_1macro), gen10580)

				gen10582 := Call(__e, ShenFunc(symcons), gen10581, Nil)

				gen10583 := Call(__e, ShenFunc(symcons), gen10576, gen10582)

				__e.TailApply(ShenFunc(symcons), gen10574, gen10583)

				return

			} else {
				__e.Return(V820)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.assoc-macro"), gen10608)

		gen10644 := MakeNative(func(__e Evaluator) {
			V822 := __e.Get(1)
			_ = V822
			gen10642 := Call(__e, ShenFunc(symcons_2), V822)

			var gen10643 Obj
			if True == gen10642 {
				gen10639 := Call(__e, ShenFunc(symhd), V822)

				gen10640 := Call(__e, ShenFunc(sym_a), MakeSymbol("let"), gen10639)

				var gen10641 Obj
				if True == gen10640 {
					gen10636 := Call(__e, ShenFunc(symtl), V822)

					gen10637 := Call(__e, ShenFunc(symcons_2), gen10636)

					var gen10638 Obj
					if True == gen10637 {
						gen10632 := Call(__e, ShenFunc(symtl), V822)

						gen10633 := Call(__e, ShenFunc(symtl), gen10632)

						gen10634 := Call(__e, ShenFunc(symcons_2), gen10633)

						var gen10635 Obj
						if True == gen10634 {
							gen10627 := Call(__e, ShenFunc(symtl), V822)

							gen10628 := Call(__e, ShenFunc(symtl), gen10627)

							gen10629 := Call(__e, ShenFunc(symtl), gen10628)

							gen10630 := Call(__e, ShenFunc(symcons_2), gen10629)

							var gen10631 Obj
							if True == gen10630 {
								gen10622 := Call(__e, ShenFunc(symtl), V822)

								gen10623 := Call(__e, ShenFunc(symtl), gen10622)

								gen10624 := Call(__e, ShenFunc(symtl), gen10623)

								gen10625 := Call(__e, ShenFunc(symtl), gen10624)

								gen10626 := Call(__e, ShenFunc(symcons_2), gen10625)

								if True == gen10626 {
									gen10631 = True
								} else {
									gen10631 = False
								}

							} else {
								gen10631 = False
							}
							if True == gen10631 {
								gen10635 = True
							} else {
								gen10635 = False
							}

						} else {
							gen10635 = False
						}
						if True == gen10635 {
							gen10638 = True
						} else {
							gen10638 = False
						}

					} else {
						gen10638 = False
					}
					if True == gen10638 {
						gen10641 = True
					} else {
						gen10641 = False
					}

				} else {
					gen10641 = False
				}
				if True == gen10641 {
					gen10643 = True
				} else {
					gen10643 = False
				}

			} else {
				gen10643 = False
			}
			if True == gen10643 {
				gen10609 := Call(__e, ShenFunc(symtl), V822)

				gen10610 := Call(__e, ShenFunc(symhd), gen10609)

				gen10611 := Call(__e, ShenFunc(symtl), V822)

				gen10612 := Call(__e, ShenFunc(symtl), gen10611)

				gen10613 := Call(__e, ShenFunc(symhd), gen10612)

				gen10614 := Call(__e, ShenFunc(symtl), V822)

				gen10615 := Call(__e, ShenFunc(symtl), gen10614)

				gen10616 := Call(__e, ShenFunc(symtl), gen10615)

				gen10617 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen10616)

				gen10618 := Call(__e, ShenFunc(symshen_4let_1macro), gen10617)

				gen10619 := Call(__e, ShenFunc(symcons), gen10618, Nil)

				gen10620 := Call(__e, ShenFunc(symcons), gen10613, gen10619)

				gen10621 := Call(__e, ShenFunc(symcons), gen10610, gen10620)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("let"), gen10621)

				return

			} else {
				__e.Return(V822)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.let-macro"), gen10644)

		gen10686 := MakeNative(func(__e Evaluator) {
			V824 := __e.Get(1)
			_ = V824
			gen10684 := Call(__e, ShenFunc(symcons_2), V824)

			var gen10685 Obj
			if True == gen10684 {
				gen10681 := Call(__e, ShenFunc(symhd), V824)

				gen10682 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen10681)

				var gen10683 Obj
				if True == gen10682 {
					gen10678 := Call(__e, ShenFunc(symtl), V824)

					gen10679 := Call(__e, ShenFunc(symcons_2), gen10678)

					var gen10680 Obj
					if True == gen10679 {
						gen10674 := Call(__e, ShenFunc(symtl), V824)

						gen10675 := Call(__e, ShenFunc(symtl), gen10674)

						gen10676 := Call(__e, ShenFunc(symcons_2), gen10675)

						var gen10677 Obj
						if True == gen10676 {
							gen10670 := Call(__e, ShenFunc(symtl), V824)

							gen10671 := Call(__e, ShenFunc(symtl), gen10670)

							gen10672 := Call(__e, ShenFunc(symtl), gen10671)

							gen10673 := Call(__e, ShenFunc(symcons_2), gen10672)

							if True == gen10673 {
								gen10677 = True
							} else {
								gen10677 = False
							}

						} else {
							gen10677 = False
						}
						if True == gen10677 {
							gen10680 = True
						} else {
							gen10680 = False
						}

					} else {
						gen10680 = False
					}
					if True == gen10680 {
						gen10683 = True
					} else {
						gen10683 = False
					}

				} else {
					gen10683 = False
				}
				if True == gen10683 {
					gen10685 = True
				} else {
					gen10685 = False
				}

			} else {
				gen10685 = False
			}
			if True == gen10685 {
				gen10662 := Call(__e, ShenFunc(symtl), V824)

				gen10663 := Call(__e, ShenFunc(symhd), gen10662)

				gen10664 := Call(__e, ShenFunc(symtl), V824)

				gen10665 := Call(__e, ShenFunc(symtl), gen10664)

				gen10666 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen10665)

				gen10667 := Call(__e, ShenFunc(symshen_4abs_1macro), gen10666)

				gen10668 := Call(__e, ShenFunc(symcons), gen10667, Nil)

				gen10669 := Call(__e, ShenFunc(symcons), gen10663, gen10668)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen10669)

				return

			} else {
				gen10660 := Call(__e, ShenFunc(symcons_2), V824)

				var gen10661 Obj
				if True == gen10660 {
					gen10657 := Call(__e, ShenFunc(symhd), V824)

					gen10658 := Call(__e, ShenFunc(sym_a), MakeSymbol("/."), gen10657)

					var gen10659 Obj
					if True == gen10658 {
						gen10654 := Call(__e, ShenFunc(symtl), V824)

						gen10655 := Call(__e, ShenFunc(symcons_2), gen10654)

						var gen10656 Obj
						if True == gen10655 {
							gen10650 := Call(__e, ShenFunc(symtl), V824)

							gen10651 := Call(__e, ShenFunc(symtl), gen10650)

							gen10652 := Call(__e, ShenFunc(symcons_2), gen10651)

							var gen10653 Obj
							if True == gen10652 {
								gen10646 := Call(__e, ShenFunc(symtl), V824)

								gen10647 := Call(__e, ShenFunc(symtl), gen10646)

								gen10648 := Call(__e, ShenFunc(symtl), gen10647)

								gen10649 := Call(__e, ShenFunc(sym_a), Nil, gen10648)

								if True == gen10649 {
									gen10653 = True
								} else {
									gen10653 = False
								}

							} else {
								gen10653 = False
							}
							if True == gen10653 {
								gen10656 = True
							} else {
								gen10656 = False
							}

						} else {
							gen10656 = False
						}
						if True == gen10656 {
							gen10659 = True
						} else {
							gen10659 = False
						}

					} else {
						gen10659 = False
					}
					if True == gen10659 {
						gen10661 = True
					} else {
						gen10661 = False
					}

				} else {
					gen10661 = False
				}
				if True == gen10661 {
					gen10645 := Call(__e, ShenFunc(symtl), V824)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("lambda"), gen10645)

					return

				} else {
					__e.Return(V824)
					return
				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.abs-macro"), gen10686)

		gen10765 := MakeNative(func(__e Evaluator) {
			V828 := __e.Get(1)
			_ = V828
			gen10763 := Call(__e, ShenFunc(symcons_2), V828)

			var gen10764 Obj
			if True == gen10763 {
				gen10760 := Call(__e, ShenFunc(symhd), V828)

				gen10761 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen10760)

				var gen10762 Obj
				if True == gen10761 {
					gen10757 := Call(__e, ShenFunc(symtl), V828)

					gen10758 := Call(__e, ShenFunc(symcons_2), gen10757)

					var gen10759 Obj
					if True == gen10758 {
						gen10753 := Call(__e, ShenFunc(symtl), V828)

						gen10754 := Call(__e, ShenFunc(symhd), gen10753)

						gen10755 := Call(__e, ShenFunc(sym_a), True, gen10754)

						var gen10756 Obj
						if True == gen10755 {
							gen10750 := Call(__e, ShenFunc(symtl), V828)

							gen10751 := Call(__e, ShenFunc(symtl), gen10750)

							gen10752 := Call(__e, ShenFunc(symcons_2), gen10751)

							if True == gen10752 {
								gen10756 = True
							} else {
								gen10756 = False
							}

						} else {
							gen10756 = False
						}
						if True == gen10756 {
							gen10759 = True
						} else {
							gen10759 = False
						}

					} else {
						gen10759 = False
					}
					if True == gen10759 {
						gen10762 = True
					} else {
						gen10762 = False
					}

				} else {
					gen10762 = False
				}
				if True == gen10762 {
					gen10764 = True
				} else {
					gen10764 = False
				}

			} else {
				gen10764 = False
			}
			if True == gen10764 {
				gen10748 := Call(__e, ShenFunc(symtl), V828)

				gen10749 := Call(__e, ShenFunc(symtl), gen10748)

				__e.TailApply(ShenFunc(symhd), gen10749)

				return

			} else {
				gen10746 := Call(__e, ShenFunc(symcons_2), V828)

				var gen10747 Obj
				if True == gen10746 {
					gen10743 := Call(__e, ShenFunc(symhd), V828)

					gen10744 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen10743)

					var gen10745 Obj
					if True == gen10744 {
						gen10740 := Call(__e, ShenFunc(symtl), V828)

						gen10741 := Call(__e, ShenFunc(symcons_2), gen10740)

						var gen10742 Obj
						if True == gen10741 {
							gen10736 := Call(__e, ShenFunc(symtl), V828)

							gen10737 := Call(__e, ShenFunc(symtl), gen10736)

							gen10738 := Call(__e, ShenFunc(symcons_2), gen10737)

							var gen10739 Obj
							if True == gen10738 {
								gen10732 := Call(__e, ShenFunc(symtl), V828)

								gen10733 := Call(__e, ShenFunc(symtl), gen10732)

								gen10734 := Call(__e, ShenFunc(symtl), gen10733)

								gen10735 := Call(__e, ShenFunc(sym_a), Nil, gen10734)

								if True == gen10735 {
									gen10739 = True
								} else {
									gen10739 = False
								}

							} else {
								gen10739 = False
							}
							if True == gen10739 {
								gen10742 = True
							} else {
								gen10742 = False
							}

						} else {
							gen10742 = False
						}
						if True == gen10742 {
							gen10745 = True
						} else {
							gen10745 = False
						}

					} else {
						gen10745 = False
					}
					if True == gen10745 {
						gen10747 = True
					} else {
						gen10747 = False
					}

				} else {
					gen10747 = False
				}
				if True == gen10747 {
					gen10722 := Call(__e, ShenFunc(symtl), V828)

					gen10723 := Call(__e, ShenFunc(symhd), gen10722)

					gen10724 := Call(__e, ShenFunc(symtl), V828)

					gen10725 := Call(__e, ShenFunc(symtl), gen10724)

					gen10726 := Call(__e, ShenFunc(symhd), gen10725)

					gen10727 := Call(__e, ShenFunc(symcons), MakeString("error: cases exhausted"), Nil)

					gen10728 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen10727)

					gen10729 := Call(__e, ShenFunc(symcons), gen10728, Nil)

					gen10730 := Call(__e, ShenFunc(symcons), gen10726, gen10729)

					gen10731 := Call(__e, ShenFunc(symcons), gen10723, gen10730)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10731)

					return

				} else {
					gen10720 := Call(__e, ShenFunc(symcons_2), V828)

					var gen10721 Obj
					if True == gen10720 {
						gen10717 := Call(__e, ShenFunc(symhd), V828)

						gen10718 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen10717)

						var gen10719 Obj
						if True == gen10718 {
							gen10714 := Call(__e, ShenFunc(symtl), V828)

							gen10715 := Call(__e, ShenFunc(symcons_2), gen10714)

							var gen10716 Obj
							if True == gen10715 {
								gen10711 := Call(__e, ShenFunc(symtl), V828)

								gen10712 := Call(__e, ShenFunc(symtl), gen10711)

								gen10713 := Call(__e, ShenFunc(symcons_2), gen10712)

								if True == gen10713 {
									gen10716 = True
								} else {
									gen10716 = False
								}

							} else {
								gen10716 = False
							}
							if True == gen10716 {
								gen10719 = True
							} else {
								gen10719 = False
							}

						} else {
							gen10719 = False
						}
						if True == gen10719 {
							gen10721 = True
						} else {
							gen10721 = False
						}

					} else {
						gen10721 = False
					}
					if True == gen10721 {
						gen10698 := Call(__e, ShenFunc(symtl), V828)

						gen10699 := Call(__e, ShenFunc(symhd), gen10698)

						gen10700 := Call(__e, ShenFunc(symtl), V828)

						gen10701 := Call(__e, ShenFunc(symtl), gen10700)

						gen10702 := Call(__e, ShenFunc(symhd), gen10701)

						gen10703 := Call(__e, ShenFunc(symtl), V828)

						gen10704 := Call(__e, ShenFunc(symtl), gen10703)

						gen10705 := Call(__e, ShenFunc(symtl), gen10704)

						gen10706 := Call(__e, ShenFunc(symcons), MakeSymbol("cases"), gen10705)

						gen10707 := Call(__e, ShenFunc(symshen_4cases_1macro), gen10706)

						gen10708 := Call(__e, ShenFunc(symcons), gen10707, Nil)

						gen10709 := Call(__e, ShenFunc(symcons), gen10702, gen10708)

						gen10710 := Call(__e, ShenFunc(symcons), gen10699, gen10709)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("if"), gen10710)

						return

					} else {
						gen10696 := Call(__e, ShenFunc(symcons_2), V828)

						var gen10697 Obj
						if True == gen10696 {
							gen10693 := Call(__e, ShenFunc(symhd), V828)

							gen10694 := Call(__e, ShenFunc(sym_a), MakeSymbol("cases"), gen10693)

							var gen10695 Obj
							if True == gen10694 {
								gen10690 := Call(__e, ShenFunc(symtl), V828)

								gen10691 := Call(__e, ShenFunc(symcons_2), gen10690)

								var gen10692 Obj
								if True == gen10691 {
									gen10687 := Call(__e, ShenFunc(symtl), V828)

									gen10688 := Call(__e, ShenFunc(symtl), gen10687)

									gen10689 := Call(__e, ShenFunc(sym_a), Nil, gen10688)

									if True == gen10689 {
										gen10692 = True
									} else {
										gen10692 = False
									}

								} else {
									gen10692 = False
								}
								if True == gen10692 {
									gen10695 = True
								} else {
									gen10695 = False
								}

							} else {
								gen10695 = False
							}
							if True == gen10695 {
								gen10697 = True
							} else {
								gen10697 = False
							}

						} else {
							gen10697 = False
						}
						if True == gen10697 {
							__e.TailApply(ShenFunc(symsimple_1error), MakeString("error: odd number of case elements\n"))

							return
						} else {
							__e.Return(V828)
							return
						}

					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.cases-macro"), gen10765)

		gen10810 := MakeNative(func(__e Evaluator) {
			V830 := __e.Get(1)
			_ = V830
			gen10808 := Call(__e, ShenFunc(symcons_2), V830)

			var gen10809 Obj
			if True == gen10808 {
				gen10805 := Call(__e, ShenFunc(symhd), V830)

				gen10806 := Call(__e, ShenFunc(sym_a), MakeSymbol("time"), gen10805)

				var gen10807 Obj
				if True == gen10806 {
					gen10802 := Call(__e, ShenFunc(symtl), V830)

					gen10803 := Call(__e, ShenFunc(symcons_2), gen10802)

					var gen10804 Obj
					if True == gen10803 {
						gen10799 := Call(__e, ShenFunc(symtl), V830)

						gen10800 := Call(__e, ShenFunc(symtl), gen10799)

						gen10801 := Call(__e, ShenFunc(sym_a), Nil, gen10800)

						if True == gen10801 {
							gen10804 = True
						} else {
							gen10804 = False
						}

					} else {
						gen10804 = False
					}
					if True == gen10804 {
						gen10807 = True
					} else {
						gen10807 = False
					}

				} else {
					gen10807 = False
				}
				if True == gen10807 {
					gen10809 = True
				} else {
					gen10809 = False
				}

			} else {
				gen10809 = False
			}
			if True == gen10809 {
				gen10766 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

				gen10767 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen10766)

				gen10768 := Call(__e, ShenFunc(symtl), V830)

				gen10769 := Call(__e, ShenFunc(symhd), gen10768)

				gen10770 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), Nil)

				gen10771 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen10770)

				gen10772 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), Nil)

				gen10773 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen10772)

				gen10774 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen10773)

				gen10775 := Call(__e, ShenFunc(symcons), MakeSymbol("Time"), Nil)

				gen10776 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen10775)

				gen10777 := Call(__e, ShenFunc(symcons), MakeString(" secs\n"), Nil)

				gen10778 := Call(__e, ShenFunc(symcons), gen10776, gen10777)

				gen10779 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen10778)

				gen10780 := Call(__e, ShenFunc(symcons), gen10779, Nil)

				gen10781 := Call(__e, ShenFunc(symcons), MakeString("\nrun time: "), gen10780)

				gen10782 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen10781)

				gen10783 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), Nil)

				gen10784 := Call(__e, ShenFunc(symcons), gen10783, Nil)

				gen10785 := Call(__e, ShenFunc(symcons), gen10782, gen10784)

				gen10786 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prhush"), gen10785)

				gen10787 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), Nil)

				gen10788 := Call(__e, ShenFunc(symcons), gen10786, gen10787)

				gen10789 := Call(__e, ShenFunc(symcons), MakeSymbol("Message"), gen10788)

				gen10790 := Call(__e, ShenFunc(symcons), gen10774, gen10789)

				gen10791 := Call(__e, ShenFunc(symcons), MakeSymbol("Time"), gen10790)

				gen10792 := Call(__e, ShenFunc(symcons), gen10771, gen10791)

				gen10793 := Call(__e, ShenFunc(symcons), MakeSymbol("Finish"), gen10792)

				gen10794 := Call(__e, ShenFunc(symcons), gen10769, gen10793)

				gen10795 := Call(__e, ShenFunc(symcons), MakeSymbol("Result"), gen10794)

				gen10796 := Call(__e, ShenFunc(symcons), gen10767, gen10795)

				gen10797 := Call(__e, ShenFunc(symcons), MakeSymbol("Start"), gen10796)

				gen10798 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen10797)

				__e.TailApply(ShenFunc(symshen_4let_1macro), gen10798)

				return

			} else {
				__e.Return(V830)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.timer-macro"), gen10810)

		gen10817 := MakeNative(func(__e Evaluator) {
			V832 := __e.Get(1)
			_ = V832
			gen10816 := Call(__e, ShenFunc(symcons_2), V832)

			if True == gen10816 {
				gen10811 := Call(__e, ShenFunc(symhd), V832)

				gen10812 := Call(__e, ShenFunc(symtl), V832)

				gen10813 := Call(__e, ShenFunc(symshen_4tuple_1up), gen10812)

				gen10814 := Call(__e, ShenFunc(symcons), gen10813, Nil)

				gen10815 := Call(__e, ShenFunc(symcons), gen10811, gen10814)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("@p"), gen10815)

				return

			} else {
				__e.Return(V832)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.tuple-up"), gen10817)

		gen10907 := MakeNative(func(__e Evaluator) {
			V834 := __e.Get(1)
			_ = V834
			gen10905 := Call(__e, ShenFunc(symcons_2), V834)

			var gen10906 Obj
			if True == gen10905 {
				gen10902 := Call(__e, ShenFunc(symhd), V834)

				gen10903 := Call(__e, ShenFunc(sym_a), MakeSymbol("put"), gen10902)

				var gen10904 Obj
				if True == gen10903 {
					gen10899 := Call(__e, ShenFunc(symtl), V834)

					gen10900 := Call(__e, ShenFunc(symcons_2), gen10899)

					var gen10901 Obj
					if True == gen10900 {
						gen10895 := Call(__e, ShenFunc(symtl), V834)

						gen10896 := Call(__e, ShenFunc(symtl), gen10895)

						gen10897 := Call(__e, ShenFunc(symcons_2), gen10896)

						var gen10898 Obj
						if True == gen10897 {
							gen10890 := Call(__e, ShenFunc(symtl), V834)

							gen10891 := Call(__e, ShenFunc(symtl), gen10890)

							gen10892 := Call(__e, ShenFunc(symtl), gen10891)

							gen10893 := Call(__e, ShenFunc(symcons_2), gen10892)

							var gen10894 Obj
							if True == gen10893 {
								gen10885 := Call(__e, ShenFunc(symtl), V834)

								gen10886 := Call(__e, ShenFunc(symtl), gen10885)

								gen10887 := Call(__e, ShenFunc(symtl), gen10886)

								gen10888 := Call(__e, ShenFunc(symtl), gen10887)

								gen10889 := Call(__e, ShenFunc(sym_a), Nil, gen10888)

								if True == gen10889 {
									gen10894 = True
								} else {
									gen10894 = False
								}

							} else {
								gen10894 = False
							}
							if True == gen10894 {
								gen10898 = True
							} else {
								gen10898 = False
							}

						} else {
							gen10898 = False
						}
						if True == gen10898 {
							gen10901 = True
						} else {
							gen10901 = False
						}

					} else {
						gen10901 = False
					}
					if True == gen10901 {
						gen10904 = True
					} else {
						gen10904 = False
					}

				} else {
					gen10904 = False
				}
				if True == gen10904 {
					gen10906 = True
				} else {
					gen10906 = False
				}

			} else {
				gen10906 = False
			}
			if True == gen10906 {
				gen10870 := Call(__e, ShenFunc(symtl), V834)

				gen10871 := Call(__e, ShenFunc(symhd), gen10870)

				gen10872 := Call(__e, ShenFunc(symtl), V834)

				gen10873 := Call(__e, ShenFunc(symtl), gen10872)

				gen10874 := Call(__e, ShenFunc(symhd), gen10873)

				gen10875 := Call(__e, ShenFunc(symtl), V834)

				gen10876 := Call(__e, ShenFunc(symtl), gen10875)

				gen10877 := Call(__e, ShenFunc(symtl), gen10876)

				gen10878 := Call(__e, ShenFunc(symhd), gen10877)

				gen10879 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

				gen10880 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen10879)

				gen10881 := Call(__e, ShenFunc(symcons), gen10880, Nil)

				gen10882 := Call(__e, ShenFunc(symcons), gen10878, gen10881)

				gen10883 := Call(__e, ShenFunc(symcons), gen10874, gen10882)

				gen10884 := Call(__e, ShenFunc(symcons), gen10871, gen10883)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("put"), gen10884)

				return

			} else {
				gen10868 := Call(__e, ShenFunc(symcons_2), V834)

				var gen10869 Obj
				if True == gen10868 {
					gen10865 := Call(__e, ShenFunc(symhd), V834)

					gen10866 := Call(__e, ShenFunc(sym_a), MakeSymbol("get"), gen10865)

					var gen10867 Obj
					if True == gen10866 {
						gen10862 := Call(__e, ShenFunc(symtl), V834)

						gen10863 := Call(__e, ShenFunc(symcons_2), gen10862)

						var gen10864 Obj
						if True == gen10863 {
							gen10858 := Call(__e, ShenFunc(symtl), V834)

							gen10859 := Call(__e, ShenFunc(symtl), gen10858)

							gen10860 := Call(__e, ShenFunc(symcons_2), gen10859)

							var gen10861 Obj
							if True == gen10860 {
								gen10854 := Call(__e, ShenFunc(symtl), V834)

								gen10855 := Call(__e, ShenFunc(symtl), gen10854)

								gen10856 := Call(__e, ShenFunc(symtl), gen10855)

								gen10857 := Call(__e, ShenFunc(sym_a), Nil, gen10856)

								if True == gen10857 {
									gen10861 = True
								} else {
									gen10861 = False
								}

							} else {
								gen10861 = False
							}
							if True == gen10861 {
								gen10864 = True
							} else {
								gen10864 = False
							}

						} else {
							gen10864 = False
						}
						if True == gen10864 {
							gen10867 = True
						} else {
							gen10867 = False
						}

					} else {
						gen10867 = False
					}
					if True == gen10867 {
						gen10869 = True
					} else {
						gen10869 = False
					}

				} else {
					gen10869 = False
				}
				if True == gen10869 {
					gen10844 := Call(__e, ShenFunc(symtl), V834)

					gen10845 := Call(__e, ShenFunc(symhd), gen10844)

					gen10846 := Call(__e, ShenFunc(symtl), V834)

					gen10847 := Call(__e, ShenFunc(symtl), gen10846)

					gen10848 := Call(__e, ShenFunc(symhd), gen10847)

					gen10849 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

					gen10850 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen10849)

					gen10851 := Call(__e, ShenFunc(symcons), gen10850, Nil)

					gen10852 := Call(__e, ShenFunc(symcons), gen10848, gen10851)

					gen10853 := Call(__e, ShenFunc(symcons), gen10845, gen10852)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("get"), gen10853)

					return

				} else {
					gen10842 := Call(__e, ShenFunc(symcons_2), V834)

					var gen10843 Obj
					if True == gen10842 {
						gen10839 := Call(__e, ShenFunc(symhd), V834)

						gen10840 := Call(__e, ShenFunc(sym_a), MakeSymbol("unput"), gen10839)

						var gen10841 Obj
						if True == gen10840 {
							gen10836 := Call(__e, ShenFunc(symtl), V834)

							gen10837 := Call(__e, ShenFunc(symcons_2), gen10836)

							var gen10838 Obj
							if True == gen10837 {
								gen10832 := Call(__e, ShenFunc(symtl), V834)

								gen10833 := Call(__e, ShenFunc(symtl), gen10832)

								gen10834 := Call(__e, ShenFunc(symcons_2), gen10833)

								var gen10835 Obj
								if True == gen10834 {
									gen10828 := Call(__e, ShenFunc(symtl), V834)

									gen10829 := Call(__e, ShenFunc(symtl), gen10828)

									gen10830 := Call(__e, ShenFunc(symtl), gen10829)

									gen10831 := Call(__e, ShenFunc(sym_a), Nil, gen10830)

									if True == gen10831 {
										gen10835 = True
									} else {
										gen10835 = False
									}

								} else {
									gen10835 = False
								}
								if True == gen10835 {
									gen10838 = True
								} else {
									gen10838 = False
								}

							} else {
								gen10838 = False
							}
							if True == gen10838 {
								gen10841 = True
							} else {
								gen10841 = False
							}

						} else {
							gen10841 = False
						}
						if True == gen10841 {
							gen10843 = True
						} else {
							gen10843 = False
						}

					} else {
						gen10843 = False
					}
					if True == gen10843 {
						gen10818 := Call(__e, ShenFunc(symtl), V834)

						gen10819 := Call(__e, ShenFunc(symhd), gen10818)

						gen10820 := Call(__e, ShenFunc(symtl), V834)

						gen10821 := Call(__e, ShenFunc(symtl), gen10820)

						gen10822 := Call(__e, ShenFunc(symhd), gen10821)

						gen10823 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), Nil)

						gen10824 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen10823)

						gen10825 := Call(__e, ShenFunc(symcons), gen10824, Nil)

						gen10826 := Call(__e, ShenFunc(symcons), gen10822, gen10825)

						gen10827 := Call(__e, ShenFunc(symcons), gen10819, gen10826)

						__e.TailApply(ShenFunc(symcons), MakeSymbol("unput"), gen10827)

						return

					} else {
						__e.Return(V834)
						return
					}

				}

			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.put/get-macro"), gen10907)

		gen10924 := MakeNative(func(__e Evaluator) {
			V836 := __e.Get(1)
			_ = V836
			gen10922 := Call(__e, ShenFunc(symcons_2), V836)

			var gen10923 Obj
			if True == gen10922 {
				gen10919 := Call(__e, ShenFunc(symhd), V836)

				gen10920 := Call(__e, ShenFunc(sym_a), MakeSymbol("function"), gen10919)

				var gen10921 Obj
				if True == gen10920 {
					gen10916 := Call(__e, ShenFunc(symtl), V836)

					gen10917 := Call(__e, ShenFunc(symcons_2), gen10916)

					var gen10918 Obj
					if True == gen10917 {
						gen10913 := Call(__e, ShenFunc(symtl), V836)

						gen10914 := Call(__e, ShenFunc(symtl), gen10913)

						gen10915 := Call(__e, ShenFunc(sym_a), Nil, gen10914)

						if True == gen10915 {
							gen10918 = True
						} else {
							gen10918 = False
						}

					} else {
						gen10918 = False
					}
					if True == gen10918 {
						gen10921 = True
					} else {
						gen10921 = False
					}

				} else {
					gen10921 = False
				}
				if True == gen10921 {
					gen10923 = True
				} else {
					gen10923 = False
				}

			} else {
				gen10923 = False
			}
			if True == gen10923 {
				gen10908 := Call(__e, ShenFunc(symtl), V836)

				gen10909 := Call(__e, ShenFunc(symhd), gen10908)

				gen10910 := Call(__e, ShenFunc(symtl), V836)

				gen10911 := Call(__e, ShenFunc(symhd), gen10910)

				gen10912 := Call(__e, ShenFunc(symarity), gen10911)

				__e.TailApply(ShenFunc(symshen_4function_1abstraction), gen10909, gen10912)

				return

			} else {
				__e.Return(V836)
				return
			}

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-macro"), gen10924)

		gen10929 := MakeNative(func(__e Evaluator) {
			V839 := __e.Get(1)
			_ = V839
			V840 := __e.Get(2)
			_ = V840
			gen10928 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V840)

			if True == gen10928 {
				gen10927 := Call(__e, ShenFunc(symshen_4app), V839, MakeString(" has no lambda form\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen10927)

				return

			} else {
				gen10926 := Call(__e, ShenFunc(sym_a), MakeNumber(-1), V840)

				if True == gen10926 {
					gen10925 := Call(__e, ShenFunc(symcons), V839, Nil)

					__e.TailApply(ShenFunc(symcons), MakeSymbol("function"), gen10925)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4function_1abstraction_1help), V839, V840, Nil)

					return
				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-abstraction"), gen10929)

		gen10938 := MakeNative(func(__e Evaluator) {
			V844 := __e.Get(1)
			_ = V844
			V845 := __e.Get(2)
			_ = V845
			V846 := __e.Get(3)
			_ = V846
			gen10937 := Call(__e, ShenFunc(sym_a), MakeNumber(0), V845)

			if True == gen10937 {
				__e.TailApply(ShenFunc(symcons), V844, V846)

				return
			} else {
				gen10930 := Call(__e, ShenFunc(symgensym), MakeSymbol("V"))

				X := gen10930
				gen10931 := Call(__e, ShenFunc(sym_1), V845, MakeNumber(1))

				gen10932 := Call(__e, ShenFunc(symcons), X, Nil)

				gen10933 := Call(__e, ShenFunc(symappend), V846, gen10932)

				gen10934 := Call(__e, ShenFunc(symshen_4function_1abstraction_1help), V844, gen10931, gen10933)

				gen10935 := Call(__e, ShenFunc(symcons), gen10934, Nil)

				gen10936 := Call(__e, ShenFunc(symcons), X, gen10935)

				__e.TailApply(ShenFunc(symcons), MakeSymbol("/."), gen10936)

				return

			}

		}, 3)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.function-abstraction-help"), gen10938)

		gen10946 := MakeNative(func(__e Evaluator) {
			V848 := __e.Get(1)
			_ = V848
			gen10939 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*macroreg*"))

			MacroReg := gen10939
			gen10940 := Call(__e, ShenFunc(symshen_4findpos), V848, MacroReg)

			Pos := gen10940
			gen10941 := Call(__e, ShenFunc(symremove), V848, MacroReg)

			gen10942 := Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen10941)

			Remove1 := gen10942
			_ = Remove1
			gen10943 := Call(__e, ShenFunc(symvalue), MakeSymbol("*macros*"))

			gen10944 := Call(__e, ShenFunc(symshen_4remove_1nth), Pos, gen10943)

			gen10945 := Call(__e, ShenFunc(symset), MakeSymbol("*macros*"), gen10944)

			Remove2 := gen10945
			_ = Remove2
			__e.Return(V848)
			return

		}, 1)
		Call(__e, ShenFunc(symdefun), MakeSymbol("undefmacro"), gen10946)

		gen10956 := MakeNative(func(__e Evaluator) {
			V858 := __e.Get(1)
			_ = V858
			V859 := __e.Get(2)
			_ = V859
			gen10955 := Call(__e, ShenFunc(sym_a), Nil, V859)

			if True == gen10955 {
				gen10954 := Call(__e, ShenFunc(symshen_4app), V858, MakeString(" is not a macro\n"), MakeSymbol("shen.a"))

				__e.TailApply(ShenFunc(symsimple_1error), gen10954)

				return

			} else {
				gen10952 := Call(__e, ShenFunc(symcons_2), V859)

				var gen10953 Obj
				if True == gen10952 {
					gen10950 := Call(__e, ShenFunc(symhd), V859)

					gen10951 := Call(__e, ShenFunc(sym_a), gen10950, V858)

					if True == gen10951 {
						gen10953 = True
					} else {
						gen10953 = False
					}

				} else {
					gen10953 = False
				}
				if True == gen10953 {
					__e.Return(MakeNumber(1))
					return
				} else {
					gen10949 := Call(__e, ShenFunc(symcons_2), V859)

					if True == gen10949 {
						gen10947 := Call(__e, ShenFunc(symtl), V859)

						gen10948 := Call(__e, ShenFunc(symshen_4findpos), V858, gen10947)

						__e.TailApply(ShenFunc(sym_7), MakeNumber(1), gen10948)

						return

					} else {
						__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.findpos"))

						return
					}

				}

			}

		}, 2)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.findpos"), gen10956)

		gen10965 := MakeNative(func(__e Evaluator) {
			V864 := __e.Get(1)
			_ = V864
			V865 := __e.Get(2)
			_ = V865
			gen10963 := Call(__e, ShenFunc(sym_a), MakeNumber(1), V864)

			var gen10964 Obj
			if True == gen10963 {
				gen10962 := Call(__e, ShenFunc(symcons_2), V865)

				if True == gen10962 {
					gen10964 = True
				} else {
					gen10964 = False
				}

			} else {
				gen10964 = False
			}
			if True == gen10964 {
				__e.TailApply(ShenFunc(symtl), V865)

				return
			} else {
				gen10961 := Call(__e, ShenFunc(symcons_2), V865)

				if True == gen10961 {
					gen10957 := Call(__e, ShenFunc(symhd), V865)

					gen10958 := Call(__e, ShenFunc(sym_1), V864, MakeNumber(1))

					gen10959 := Call(__e, ShenFunc(symtl), V865)

					gen10960 := Call(__e, ShenFunc(symshen_4remove_1nth), gen10958, gen10959)

					__e.TailApply(ShenFunc(symcons), gen10957, gen10960)

					return

				} else {
					__e.TailApply(ShenFunc(symshen_4f__error), MakeSymbol("shen.remove-nth"))

					return
				}

			}

		}, 2)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.remove-nth"), gen10965)

		return

	}, 0))
}
