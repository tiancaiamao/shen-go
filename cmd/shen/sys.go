package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp16058 := MakeNative(func(__e *ControlFlow) {
		V1920 := __e.Get(1)
		_ = V1920
		__e.TailApply(V1920)
		return
	}, 1)

	tmp16059 := Call(__e, PrimNS1Value(symns2_1set), symthaw, tmp16058)

	_ = tmp16059

	tmp16060 := MakeNative(func(__e *ControlFlow) {
		V1922 := __e.Get(1)
		_ = V1922
		tmp16061 := MakeNative(func(__e *ControlFlow) {
			Macroexpand := __e.Get(1)
			_ = Macroexpand
			tmp16065 := Call(__e, PrimNS2Value(symshen_4packaged_2), Macroexpand)

			if True == tmp16065 {
				tmp16063 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4eval_1without_1macros), Z)
					return
				}, 1)

				tmp16064 := Call(__e, PrimNS2Value(symshen_4package_1contents), Macroexpand)

				__e.TailApply(PrimNS2Value(symmap), tmp16063, tmp16064)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4eval_1without_1macros), Macroexpand)
				return
			}

		}, 1)

		tmp16066 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			__e.TailApply(PrimNS2Value(symmacroexpand), Y)
			return
		}, 1)

		tmp16067 := Call(__e, PrimNS2Value(symshen_4walk), tmp16066, V1922)

		__e.TailApply(tmp16061, tmp16067)
		return

	}, 1)

	tmp16068 := Call(__e, PrimNS1Value(symns2_1set), symeval, tmp16060)

	_ = tmp16068

	tmp16069 := MakeNative(func(__e *ControlFlow) {
		V1924 := __e.Get(1)
		_ = V1924
		tmp16070 := Call(__e, PrimNS2Value(symshen_4proc_1input_7), V1924)

		tmp16071 := Call(__e, PrimNS2Value(symshen_4elim_1def), tmp16070)

		__e.TailApply(PrimNS2Value(symeval_1kl), tmp16071)
		return

	}, 1)

	tmp16072 := Call(__e, PrimNS1Value(symns2_1set), symshen_4eval_1without_1macros, tmp16069)

	_ = tmp16072

	tmp16073 := MakeNative(func(__e *ControlFlow) {
		V1926 := __e.Get(1)
		_ = V1926
		tmp16130 := PrimIsPair(V1926)

		var ifres16111 Obj

		if True == tmp16130 {
			tmp16128 := PrimHead(V1926)

			tmp16129 := PrimEqual(syminput_7, tmp16128)

			var ifres16113 Obj

			if True == tmp16129 {
				tmp16126 := PrimTail(V1926)

				tmp16127 := PrimIsPair(tmp16126)

				var ifres16115 Obj

				if True == tmp16127 {
					tmp16123 := PrimTail(V1926)

					tmp16124 := PrimTail(tmp16123)

					tmp16125 := PrimIsPair(tmp16124)

					var ifres16117 Obj

					if True == tmp16125 {
						tmp16119 := PrimTail(V1926)

						tmp16120 := PrimTail(tmp16119)

						tmp16121 := PrimTail(tmp16120)

						tmp16122 := PrimEqual(Nil, tmp16121)

						var ifres16118 Obj

						if True == tmp16122 {
							ifres16118 = True

						} else {
							ifres16118 = False

						}

						ifres16117 = ifres16118

					} else {
						ifres16117 = False

					}

					var ifres16116 Obj

					if True == ifres16117 {
						ifres16116 = True

					} else {
						ifres16116 = False

					}

					ifres16115 = ifres16116

				} else {
					ifres16115 = False

				}

				var ifres16114 Obj

				if True == ifres16115 {
					ifres16114 = True

				} else {
					ifres16114 = False

				}

				ifres16113 = ifres16114

			} else {
				ifres16113 = False

			}

			var ifres16112 Obj

			if True == ifres16113 {
				ifres16112 = True

			} else {
				ifres16112 = False

			}

			ifres16111 = ifres16112

		} else {
			ifres16111 = False

		}

		if True == ifres16111 {
			tmp16105 := PrimTail(V1926)

			tmp16106 := PrimHead(tmp16105)

			tmp16107 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp16106)

			tmp16108 := PrimTail(V1926)

			tmp16109 := PrimTail(tmp16108)

			tmp16110 := PrimCons(tmp16107, tmp16109)

			__e.Return(PrimCons(syminput_7, tmp16110))
			return

		} else {
			tmp16104 := PrimIsPair(V1926)

			var ifres16085 Obj

			if True == tmp16104 {
				tmp16102 := PrimHead(V1926)

				tmp16103 := PrimEqual(symshen_4read_7, tmp16102)

				var ifres16087 Obj

				if True == tmp16103 {
					tmp16100 := PrimTail(V1926)

					tmp16101 := PrimIsPair(tmp16100)

					var ifres16089 Obj

					if True == tmp16101 {
						tmp16097 := PrimTail(V1926)

						tmp16098 := PrimTail(tmp16097)

						tmp16099 := PrimIsPair(tmp16098)

						var ifres16091 Obj

						if True == tmp16099 {
							tmp16093 := PrimTail(V1926)

							tmp16094 := PrimTail(tmp16093)

							tmp16095 := PrimTail(tmp16094)

							tmp16096 := PrimEqual(Nil, tmp16095)

							var ifres16092 Obj

							if True == tmp16096 {
								ifres16092 = True

							} else {
								ifres16092 = False

							}

							ifres16091 = ifres16092

						} else {
							ifres16091 = False

						}

						var ifres16090 Obj

						if True == ifres16091 {
							ifres16090 = True

						} else {
							ifres16090 = False

						}

						ifres16089 = ifres16090

					} else {
						ifres16089 = False

					}

					var ifres16088 Obj

					if True == ifres16089 {
						ifres16088 = True

					} else {
						ifres16088 = False

					}

					ifres16087 = ifres16088

				} else {
					ifres16087 = False

				}

				var ifres16086 Obj

				if True == ifres16087 {
					ifres16086 = True

				} else {
					ifres16086 = False

				}

				ifres16085 = ifres16086

			} else {
				ifres16085 = False

			}

			if True == ifres16085 {
				tmp16079 := PrimTail(V1926)

				tmp16080 := PrimHead(tmp16079)

				tmp16081 := Call(__e, PrimNS2Value(symshen_4rcons__form), tmp16080)

				tmp16082 := PrimTail(V1926)

				tmp16083 := PrimTail(tmp16082)

				tmp16084 := PrimCons(tmp16081, tmp16083)

				__e.Return(PrimCons(symshen_4read_7, tmp16084))
				return

			} else {
				tmp16078 := PrimIsPair(V1926)

				if True == tmp16078 {
					tmp16077 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimNS2Value(symshen_4proc_1input_7), Z)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp16077, V1926)
					return

				} else {
					__e.Return(V1926)
					return
				}

			}

		}

	}, 1)

	tmp16131 := Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1input_7, tmp16073)

	_ = tmp16131

	tmp16132 := MakeNative(func(__e *ControlFlow) {
		V1928 := __e.Get(1)
		_ = V1928
		tmp16187 := PrimIsPair(V1928)

		var ifres16179 Obj

		if True == tmp16187 {
			tmp16185 := PrimHead(V1928)

			tmp16186 := PrimEqual(symdefine, tmp16185)

			var ifres16181 Obj

			if True == tmp16186 {
				tmp16183 := PrimTail(V1928)

				tmp16184 := PrimIsPair(tmp16183)

				var ifres16182 Obj

				if True == tmp16184 {
					ifres16182 = True

				} else {
					ifres16182 = False

				}

				ifres16181 = ifres16182

			} else {
				ifres16181 = False

			}

			var ifres16180 Obj

			if True == ifres16181 {
				ifres16180 = True

			} else {
				ifres16180 = False

			}

			ifres16179 = ifres16180

		} else {
			ifres16179 = False

		}

		if True == ifres16179 {
			tmp16175 := PrimTail(V1928)

			tmp16176 := PrimHead(tmp16175)

			tmp16177 := PrimTail(V1928)

			tmp16178 := PrimTail(tmp16177)

			__e.TailApply(PrimNS2Value(symshen_4shen_1_6kl), tmp16176, tmp16178)
			return

		} else {
			tmp16174 := PrimIsPair(V1928)

			var ifres16166 Obj

			if True == tmp16174 {
				tmp16172 := PrimHead(V1928)

				tmp16173 := PrimEqual(symdefmacro, tmp16172)

				var ifres16168 Obj

				if True == tmp16173 {
					tmp16170 := PrimTail(V1928)

					tmp16171 := PrimIsPair(tmp16170)

					var ifres16169 Obj

					if True == tmp16171 {
						ifres16169 = True

					} else {
						ifres16169 = False

					}

					ifres16168 = ifres16169

				} else {
					ifres16168 = False

				}

				var ifres16167 Obj

				if True == ifres16168 {
					ifres16167 = True

				} else {
					ifres16167 = False

				}

				ifres16166 = ifres16167

			} else {
				ifres16166 = False

			}

			if True == ifres16166 {
				tmp16149 := MakeNative(func(__e *ControlFlow) {
					Default := __e.Get(1)
					_ = Default
					tmp16150 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						tmp16151 := MakeNative(func(__e *ControlFlow) {
							MacroAdd := __e.Get(1)
							_ = MacroAdd
							__e.Return(Def)
							return
						}, 1)

						tmp16152 := PrimTail(V1928)

						tmp16153 := PrimHead(tmp16152)

						tmp16154 := Call(__e, PrimNS2Value(symshen_4add_1macro), tmp16153)

						__e.TailApply(tmp16151, tmp16154)
						return

					}, 1)

					tmp16155 := PrimTail(V1928)

					tmp16156 := PrimHead(tmp16155)

					tmp16157 := PrimTail(V1928)

					tmp16158 := PrimTail(tmp16157)

					tmp16159 := Call(__e, PrimNS2Value(symappend), tmp16158, Default)

					tmp16160 := PrimCons(tmp16156, tmp16159)

					tmp16161 := PrimCons(symdefine, tmp16160)

					tmp16162 := Call(__e, PrimNS2Value(symshen_4elim_1def), tmp16161)

					__e.TailApply(tmp16150, tmp16162)
					return

				}, 1)

				tmp16163 := PrimCons(symX, Nil)

				tmp16164 := PrimCons(sym_1_6, tmp16163)

				tmp16165 := PrimCons(symX, tmp16164)

				__e.TailApply(tmp16149, tmp16165)
				return

			} else {
				tmp16148 := PrimIsPair(V1928)

				var ifres16140 Obj

				if True == tmp16148 {
					tmp16146 := PrimHead(V1928)

					tmp16147 := PrimEqual(symdefcc, tmp16146)

					var ifres16142 Obj

					if True == tmp16147 {
						tmp16144 := PrimTail(V1928)

						tmp16145 := PrimIsPair(tmp16144)

						var ifres16143 Obj

						if True == tmp16145 {
							ifres16143 = True

						} else {
							ifres16143 = False

						}

						ifres16142 = ifres16143

					} else {
						ifres16142 = False

					}

					var ifres16141 Obj

					if True == ifres16142 {
						ifres16141 = True

					} else {
						ifres16141 = False

					}

					ifres16140 = ifres16141

				} else {
					ifres16140 = False

				}

				if True == ifres16140 {
					tmp16139 := Call(__e, PrimNS2Value(symshen_4yacc), V1928)

					__e.TailApply(PrimNS2Value(symshen_4elim_1def), tmp16139)
					return

				} else {
					tmp16138 := PrimIsPair(V1928)

					if True == tmp16138 {
						tmp16137 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(PrimNS2Value(symshen_4elim_1def), Z)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp16137, V1928)
						return

					} else {
						__e.Return(V1928)
						return
					}

				}

			}

		}

	}, 1)

	tmp16188 := Call(__e, PrimNS1Value(symns2_1set), symshen_4elim_1def, tmp16132)

	_ = tmp16188

	tmp16189 := MakeNative(func(__e *ControlFlow) {
		V1930 := __e.Get(1)
		_ = V1930
		tmp16190 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			tmp16191 := MakeNative(func(__e *ControlFlow) {
				NewMacroReg := __e.Get(1)
				_ = NewMacroReg
				tmp16196 := PrimEqual(MacroReg, NewMacroReg)

				if True == tmp16196 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp16193 := Call(__e, PrimNS2Value(symfunction), V1930)

					tmp16194 := PrimNS3Value(sym_dmacros_d)

					tmp16195 := PrimCons(tmp16193, tmp16194)

					__e.Return(PrimNS3Set(sym_dmacros_d, tmp16195))
					return

				}

			}, 1)

			tmp16197 := PrimNS3Value(symshen_4_dmacroreg_d)

			tmp16198 := Call(__e, PrimNS2Value(symadjoin), V1930, tmp16197)

			tmp16199 := PrimNS3Set(symshen_4_dmacroreg_d, tmp16198)

			__e.TailApply(tmp16191, tmp16199)
			return

		}, 1)

		tmp16200 := PrimNS3Value(symshen_4_dmacroreg_d)

		__e.TailApply(tmp16190, tmp16200)
		return

	}, 1)

	tmp16201 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1macro, tmp16189)

	_ = tmp16201

	tmp16202 := MakeNative(func(__e *ControlFlow) {
		V1938 := __e.Get(1)
		_ = V1938
		tmp16217 := PrimIsPair(V1938)

		var ifres16204 Obj

		if True == tmp16217 {
			tmp16215 := PrimHead(V1938)

			tmp16216 := PrimEqual(sympackage, tmp16215)

			var ifres16206 Obj

			if True == tmp16216 {
				tmp16213 := PrimTail(V1938)

				tmp16214 := PrimIsPair(tmp16213)

				var ifres16208 Obj

				if True == tmp16214 {
					tmp16210 := PrimTail(V1938)

					tmp16211 := PrimTail(tmp16210)

					tmp16212 := PrimIsPair(tmp16211)

					var ifres16209 Obj

					if True == tmp16212 {
						ifres16209 = True

					} else {
						ifres16209 = False

					}

					ifres16208 = ifres16209

				} else {
					ifres16208 = False

				}

				var ifres16207 Obj

				if True == ifres16208 {
					ifres16207 = True

				} else {
					ifres16207 = False

				}

				ifres16206 = ifres16207

			} else {
				ifres16206 = False

			}

			var ifres16205 Obj

			if True == ifres16206 {
				ifres16205 = True

			} else {
				ifres16205 = False

			}

			ifres16204 = ifres16205

		} else {
			ifres16204 = False

		}

		if True == ifres16204 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp16218 := Call(__e, PrimNS1Value(symns2_1set), symshen_4packaged_2, tmp16202)

	_ = tmp16218

	tmp16219 := MakeNative(func(__e *ControlFlow) {
		V1940 := __e.Get(1)
		_ = V1940
		tmp16220 := MakeNative(func(__e *ControlFlow) {
			tmp16221 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1940, symshen_4external_1symbols, tmp16221)
			return

		}, 0)

		tmp16222 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp16223 := Call(__e, PrimNS2Value(symshen_4app), V1940, MakeString(" has not been used.\n"), symshen_4a)

			tmp16224 := PrimStringConcat(MakeString("package "), tmp16223)

			__e.Return(PrimSimpleError(tmp16224))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp16220, tmp16222)
		return

	}, 1)

	tmp16225 := Call(__e, PrimNS1Value(symns2_1set), symexternal, tmp16219)

	_ = tmp16225

	tmp16226 := MakeNative(func(__e *ControlFlow) {
		V1942 := __e.Get(1)
		_ = V1942
		tmp16227 := MakeNative(func(__e *ControlFlow) {
			tmp16228 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1942, symshen_4internal_1symbols, tmp16228)
			return

		}, 0)

		tmp16229 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp16230 := Call(__e, PrimNS2Value(symshen_4app), V1942, MakeString(" has not been used.\n"), symshen_4a)

			tmp16231 := PrimStringConcat(MakeString("package "), tmp16230)

			__e.Return(PrimSimpleError(tmp16231))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp16227, tmp16229)
		return

	}, 1)

	tmp16232 := Call(__e, PrimNS1Value(symns2_1set), syminternal, tmp16226)

	_ = tmp16232

	tmp16233 := MakeNative(func(__e *ControlFlow) {
		V1946 := __e.Get(1)
		_ = V1946
		tmp16286 := PrimIsPair(V1946)

		var ifres16268 Obj

		if True == tmp16286 {
			tmp16284 := PrimHead(V1946)

			tmp16285 := PrimEqual(sympackage, tmp16284)

			var ifres16270 Obj

			if True == tmp16285 {
				tmp16282 := PrimTail(V1946)

				tmp16283 := PrimIsPair(tmp16282)

				var ifres16272 Obj

				if True == tmp16283 {
					tmp16279 := PrimTail(V1946)

					tmp16280 := PrimHead(tmp16279)

					tmp16281 := PrimEqual(symnull, tmp16280)

					var ifres16274 Obj

					if True == tmp16281 {
						tmp16276 := PrimTail(V1946)

						tmp16277 := PrimTail(tmp16276)

						tmp16278 := PrimIsPair(tmp16277)

						var ifres16275 Obj

						if True == tmp16278 {
							ifres16275 = True

						} else {
							ifres16275 = False

						}

						ifres16274 = ifres16275

					} else {
						ifres16274 = False

					}

					var ifres16273 Obj

					if True == ifres16274 {
						ifres16273 = True

					} else {
						ifres16273 = False

					}

					ifres16272 = ifres16273

				} else {
					ifres16272 = False

				}

				var ifres16271 Obj

				if True == ifres16272 {
					ifres16271 = True

				} else {
					ifres16271 = False

				}

				ifres16270 = ifres16271

			} else {
				ifres16270 = False

			}

			var ifres16269 Obj

			if True == ifres16270 {
				ifres16269 = True

			} else {
				ifres16269 = False

			}

			ifres16268 = ifres16269

		} else {
			ifres16268 = False

		}

		if True == ifres16268 {
			tmp16266 := PrimTail(V1946)

			tmp16267 := PrimTail(tmp16266)

			__e.Return(PrimTail(tmp16267))
			return

		} else {
			tmp16265 := PrimIsPair(V1946)

			var ifres16252 Obj

			if True == tmp16265 {
				tmp16263 := PrimHead(V1946)

				tmp16264 := PrimEqual(sympackage, tmp16263)

				var ifres16254 Obj

				if True == tmp16264 {
					tmp16261 := PrimTail(V1946)

					tmp16262 := PrimIsPair(tmp16261)

					var ifres16256 Obj

					if True == tmp16262 {
						tmp16258 := PrimTail(V1946)

						tmp16259 := PrimTail(tmp16258)

						tmp16260 := PrimIsPair(tmp16259)

						var ifres16257 Obj

						if True == tmp16260 {
							ifres16257 = True

						} else {
							ifres16257 = False

						}

						ifres16256 = ifres16257

					} else {
						ifres16256 = False

					}

					var ifres16255 Obj

					if True == ifres16256 {
						ifres16255 = True

					} else {
						ifres16255 = False

					}

					ifres16254 = ifres16255

				} else {
					ifres16254 = False

				}

				var ifres16253 Obj

				if True == ifres16254 {
					ifres16253 = True

				} else {
					ifres16253 = False

				}

				ifres16252 = ifres16253

			} else {
				ifres16252 = False

			}

			if True == ifres16252 {
				tmp16236 := MakeNative(func(__e *ControlFlow) {
					PackageNameDot := __e.Get(1)
					_ = PackageNameDot
					tmp16237 := MakeNative(func(__e *ControlFlow) {
						ExpPackageNameDot := __e.Get(1)
						_ = ExpPackageNameDot
						tmp16238 := PrimTail(V1946)

						tmp16239 := PrimHead(tmp16238)

						tmp16240 := PrimTail(V1946)

						tmp16241 := PrimTail(tmp16240)

						tmp16242 := PrimHead(tmp16241)

						tmp16243 := PrimTail(V1946)

						tmp16244 := PrimTail(tmp16243)

						tmp16245 := PrimTail(tmp16244)

						__e.TailApply(PrimNS2Value(symshen_4packageh), tmp16239, tmp16242, tmp16245, ExpPackageNameDot)
						return

					}, 1)

					tmp16246 := Call(__e, PrimNS2Value(symexplode), PackageNameDot)

					__e.TailApply(tmp16237, tmp16246)
					return

				}, 1)

				tmp16247 := PrimTail(V1946)

				tmp16248 := PrimHead(tmp16247)

				tmp16249 := PrimStr(tmp16248)

				tmp16250 := PrimStringConcat(tmp16249, MakeString("."))

				tmp16251 := PrimIntern(tmp16250)

				__e.TailApply(tmp16236, tmp16251)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4package_1contents)
				return
			}

		}

	}, 1)

	tmp16287 := Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1contents, tmp16233)

	_ = tmp16287

	tmp16288 := MakeNative(func(__e *ControlFlow) {
		V1949 := __e.Get(1)
		_ = V1949
		V1950 := __e.Get(2)
		_ = V1950
		tmp16292 := PrimIsPair(V1950)

		if True == tmp16292 {
			tmp16290 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4walk), V1949, Z)
				return
			}, 1)

			tmp16291 := Call(__e, PrimNS2Value(symmap), tmp16290, V1950)

			__e.TailApply(V1949, tmp16291)
			return

		} else {
			__e.TailApply(V1949, V1950)
			return
		}

	}, 2)

	tmp16293 := Call(__e, PrimNS1Value(symns2_1set), symshen_4walk, tmp16288)

	_ = tmp16293

	tmp16294 := MakeNative(func(__e *ControlFlow) {
		V1954 := __e.Get(1)
		_ = V1954
		V1955 := __e.Get(2)
		_ = V1955
		V1956 := __e.Get(3)
		_ = V1956
		tmp16295 := MakeNative(func(__e *ControlFlow) {
			O := __e.Get(1)
			_ = O
			tmp16302 := Call(__e, PrimNS2Value(symfail))

			tmp16303 := PrimEqual(tmp16302, O)

			var ifres16297 Obj

			if True == tmp16303 {
				ifres16297 = True

			} else {
				tmp16299 := PrimHead(O)

				tmp16300 := Call(__e, PrimNS2Value(symempty_2), tmp16299)

				tmp16301 := PrimNot(tmp16300)

				var ifres16298 Obj

				if True == tmp16301 {
					ifres16298 = True

				} else {
					ifres16298 = False

				}

				ifres16297 = ifres16298

			}

			if True == ifres16297 {
				__e.TailApply(V1956, O)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4hdtl), O)
				return
			}

		}, 1)

		tmp16304 := PrimCons(Nil, Nil)

		tmp16305 := PrimCons(V1955, tmp16304)

		tmp16306 := Call(__e, V1954, tmp16305)

		__e.TailApply(tmp16295, tmp16306)
		return

	}, 3)

	tmp16307 := Call(__e, PrimNS1Value(symns2_1set), symcompile, tmp16294)

	_ = tmp16307

	tmp16308 := MakeNative(func(__e *ControlFlow) {
		V1959 := __e.Get(1)
		_ = V1959
		V1960 := __e.Get(2)
		_ = V1960
		tmp16310 := Call(__e, V1959, V1960)

		if True == tmp16310 {
			__e.TailApply(PrimNS2Value(symfail))
			return
		} else {
			__e.Return(V1960)
			return
		}

	}, 2)

	tmp16311 := Call(__e, PrimNS1Value(symns2_1set), symfail_1if, tmp16308)

	_ = tmp16311

	tmp16312 := MakeNative(func(__e *ControlFlow) {
		V1963 := __e.Get(1)
		_ = V1963
		V1964 := __e.Get(2)
		_ = V1964
		__e.Return(PrimStringConcat(V1963, V1964))
		return
	}, 2)

	tmp16313 := Call(__e, PrimNS1Value(symns2_1set), sym_8s, tmp16312)

	_ = tmp16313

	tmp16314 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dtc_d))
		return
	}, 0)

	tmp16315 := Call(__e, PrimNS1Value(symns2_1set), symtc_2, tmp16314)

	_ = tmp16315

	tmp16316 := MakeNative(func(__e *ControlFlow) {
		V1966 := __e.Get(1)
		_ = V1966
		tmp16317 := MakeNative(func(__e *ControlFlow) {
			tmp16318 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V1966, symshen_4source, tmp16318)
			return

		}, 0)

		tmp16319 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp16320 := Call(__e, PrimNS2Value(symshen_4app), V1966, MakeString(" not found.\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp16320))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp16317, tmp16319)
		return

	}, 1)

	tmp16321 := Call(__e, PrimNS1Value(symns2_1set), symps, tmp16316)

	_ = tmp16321

	tmp16322 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dstinput_d))
		return
	}, 0)

	tmp16323 := Call(__e, PrimNS1Value(symns2_1set), symstinput, tmp16322)

	_ = tmp16323

	tmp16324 := MakeNative(func(__e *ControlFlow) {
		V1968 := __e.Get(1)
		_ = V1968
		tmp16325 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp16326 := MakeNative(func(__e *ControlFlow) {
				ZeroStamp := __e.Get(1)
				_ = ZeroStamp
				tmp16327 := MakeNative(func(__e *ControlFlow) {
					Standard := __e.Get(1)
					_ = Standard
					__e.Return(Standard)
					return
				}, 1)

				tmp16331 := PrimEqual(V1968, MakeNumber(0))

				var ifres16328 Obj

				if True == tmp16331 {
					ifres16328 = ZeroStamp

				} else {
					tmp16329 := Call(__e, PrimNS2Value(symfail))

					tmp16330 := Call(__e, PrimNS2Value(symshen_4fillvector), ZeroStamp, MakeNumber(1), V1968, tmp16329)

					ifres16328 = tmp16330

				}

				__e.TailApply(tmp16327, ifres16328)
				return

			}, 1)

			tmp16332 := PrimVectorSet(Vector, MakeNumber(0), V1968)

			__e.TailApply(tmp16326, tmp16332)
			return

		}, 1)

		tmp16333 := PrimNumberAdd(V1968, MakeNumber(1))

		tmp16334 := PrimAbsvector(tmp16333)

		__e.TailApply(tmp16325, tmp16334)
		return

	}, 1)

	tmp16335 := Call(__e, PrimNS1Value(symns2_1set), symvector, tmp16324)

	_ = tmp16335

	tmp16336 := MakeNative(func(__e *ControlFlow) {
		V1974 := __e.Get(1)
		_ = V1974
		V1975 := __e.Get(2)
		_ = V1975
		V1976 := __e.Get(3)
		_ = V1976
		V1977 := __e.Get(4)
		_ = V1977
		tmp16340 := PrimEqual(V1976, V1975)

		if True == tmp16340 {
			__e.Return(PrimVectorSet(V1974, V1976, V1977))
			return
		} else {
			tmp16338 := PrimVectorSet(V1974, V1975, V1977)

			tmp16339 := PrimNumberAdd(MakeNumber(1), V1975)

			__e.TailApply(PrimNS2Value(symshen_4fillvector), tmp16338, tmp16339, V1976, V1977)
			return

		}

	}, 4)

	tmp16341 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fillvector, tmp16336)

	_ = tmp16341

	tmp16342 := MakeNative(func(__e *ControlFlow) {
		V1979 := __e.Get(1)
		_ = V1979
		tmp16354 := PrimIsVector(V1979)

		if True == tmp16354 {
			tmp16345 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp16349 := PrimIsNumber(X)

				if True == tmp16349 {
					tmp16348 := PrimGreatEqual(X, MakeNumber(0))

					if True == tmp16348 {
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

			tmp16350 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V1979, MakeNumber(0)))
				return
			}, 0)

			tmp16351 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeNumber(-1))
				return
			}, 1)

			tmp16352 := Call(__e, PrimNS1Value(symtry_1catch), tmp16350, tmp16351)

			tmp16353 := Call(__e, tmp16345, tmp16352)

			if True == tmp16353 {
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

	tmp16355 := Call(__e, PrimNS1Value(symns2_1set), symvector_2, tmp16342)

	_ = tmp16355

	tmp16356 := MakeNative(func(__e *ControlFlow) {
		V1983 := __e.Get(1)
		_ = V1983
		V1984 := __e.Get(2)
		_ = V1984
		V1985 := __e.Get(3)
		_ = V1985
		tmp16358 := PrimEqual(V1984, MakeNumber(0))

		if True == tmp16358 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			__e.Return(PrimVectorSet(V1983, V1984, V1985))
			return
		}

	}, 3)

	tmp16359 := Call(__e, PrimNS1Value(symns2_1set), symvector_1_6, tmp16356)

	_ = tmp16359

	tmp16360 := MakeNative(func(__e *ControlFlow) {
		V1988 := __e.Get(1)
		_ = V1988
		V1989 := __e.Get(2)
		_ = V1989
		tmp16367 := PrimEqual(V1989, MakeNumber(0))

		if True == tmp16367 {
			__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
			return
		} else {
			tmp16362 := MakeNative(func(__e *ControlFlow) {
				VectorElement := __e.Get(1)
				_ = VectorElement
				tmp16364 := Call(__e, PrimNS2Value(symfail))

				tmp16365 := PrimEqual(VectorElement, tmp16364)

				if True == tmp16365 {
					__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
					return
				} else {
					__e.Return(VectorElement)
					return
				}

			}, 1)

			tmp16366 := PrimVectorGet(V1988, V1989)

			__e.TailApply(tmp16362, tmp16366)
			return

		}

	}, 2)

	tmp16368 := Call(__e, PrimNS1Value(symns2_1set), sym_5_1vector, tmp16360)

	_ = tmp16368

	tmp16369 := MakeNative(func(__e *ControlFlow) {
		V1991 := __e.Get(1)
		_ = V1991
		tmp16373 := PrimIsInteger(V1991)

		if True == tmp16373 {
			tmp16372 := PrimGreatEqual(V1991, MakeNumber(0))

			if True == tmp16372 {
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

	tmp16374 := Call(__e, PrimNS1Value(symns2_1set), symshen_4posint_2, tmp16369)

	_ = tmp16374

	tmp16375 := MakeNative(func(__e *ControlFlow) {
		V1993 := __e.Get(1)
		_ = V1993
		__e.Return(PrimVectorGet(V1993, MakeNumber(0)))
		return
	}, 1)

	tmp16376 := Call(__e, PrimNS1Value(symns2_1set), symlimit, tmp16375)

	_ = tmp16376

	tmp16377 := MakeNative(func(__e *ControlFlow) {
		V1995 := __e.Get(1)
		_ = V1995
		tmp16389 := Call(__e, PrimNS2Value(symboolean_2), V1995)

		var ifres16383 Obj

		if True == tmp16389 {
			ifres16383 = True

		} else {
			tmp16388 := PrimIsNumber(V1995)

			var ifres16385 Obj

			if True == tmp16388 {
				ifres16385 = True

			} else {
				tmp16387 := PrimIsString(V1995)

				var ifres16386 Obj

				if True == tmp16387 {
					ifres16386 = True

				} else {
					ifres16386 = False

				}

				ifres16385 = ifres16386

			}

			var ifres16384 Obj

			if True == ifres16385 {
				ifres16384 = True

			} else {
				ifres16384 = False

			}

			ifres16383 = ifres16384

		}

		if True == ifres16383 {
			__e.Return(False)
			return
		} else {
			tmp16379 := MakeNative(func(__e *ControlFlow) {
				tmp16380 := MakeNative(func(__e *ControlFlow) {
					String := __e.Get(1)
					_ = String
					__e.TailApply(PrimNS2Value(symshen_4analyse_1symbol_2), String)
					return
				}, 1)

				tmp16381 := PrimStr(V1995)

				__e.TailApply(tmp16380, tmp16381)
				return

			}, 0)

			tmp16382 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp16379, tmp16382)
			return

		}

	}, 1)

	tmp16390 := Call(__e, PrimNS1Value(symns2_1set), symsymbol_2, tmp16377)

	_ = tmp16390

	tmp16391 := MakeNative(func(__e *ControlFlow) {
		V1997 := __e.Get(1)
		_ = V1997
		tmp16401 := PrimEqual(MakeString(""), V1997)

		if True == tmp16401 {
			__e.Return(False)
			return
		} else {
			tmp16400 := Call(__e, PrimNS2Value(symshen_4_7string_2), V1997)

			if True == tmp16400 {
				tmp16398 := PrimPos(V1997, MakeNumber(0))

				tmp16399 := Call(__e, PrimNS2Value(symshen_4alpha_2), tmp16398)

				if True == tmp16399 {
					tmp16396 := PrimTailString(V1997)

					tmp16397 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp16396)

					if True == tmp16397 {
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
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4analyse_1symbol_2)
				return
			}

		}

	}, 1)

	tmp16402 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1symbol_2, tmp16391)

	_ = tmp16402

	tmp16403 := MakeNative(func(__e *ControlFlow) {
		V1999 := __e.Get(1)
		_ = V1999
		tmp16404 := PrimCons(MakeString("."), Nil)

		tmp16405 := PrimCons(MakeString("'"), tmp16404)

		tmp16406 := PrimCons(MakeString("#"), tmp16405)

		tmp16407 := PrimCons(MakeString("`"), tmp16406)

		tmp16408 := PrimCons(MakeString(";"), tmp16407)

		tmp16409 := PrimCons(MakeString(":"), tmp16408)

		tmp16410 := PrimCons(MakeString("}"), tmp16409)

		tmp16411 := PrimCons(MakeString("{"), tmp16410)

		tmp16412 := PrimCons(MakeString("%"), tmp16411)

		tmp16413 := PrimCons(MakeString("&"), tmp16412)

		tmp16414 := PrimCons(MakeString("<"), tmp16413)

		tmp16415 := PrimCons(MakeString(">"), tmp16414)

		tmp16416 := PrimCons(MakeString("~"), tmp16415)

		tmp16417 := PrimCons(MakeString("@"), tmp16416)

		tmp16418 := PrimCons(MakeString("!"), tmp16417)

		tmp16419 := PrimCons(MakeString("$"), tmp16418)

		tmp16420 := PrimCons(MakeString("?"), tmp16419)

		tmp16421 := PrimCons(MakeString("_"), tmp16420)

		tmp16422 := PrimCons(MakeString("-"), tmp16421)

		tmp16423 := PrimCons(MakeString("+"), tmp16422)

		tmp16424 := PrimCons(MakeString("/"), tmp16423)

		tmp16425 := PrimCons(MakeString("*"), tmp16424)

		tmp16426 := PrimCons(MakeString("="), tmp16425)

		tmp16427 := PrimCons(MakeString("z"), tmp16426)

		tmp16428 := PrimCons(MakeString("y"), tmp16427)

		tmp16429 := PrimCons(MakeString("x"), tmp16428)

		tmp16430 := PrimCons(MakeString("w"), tmp16429)

		tmp16431 := PrimCons(MakeString("v"), tmp16430)

		tmp16432 := PrimCons(MakeString("u"), tmp16431)

		tmp16433 := PrimCons(MakeString("t"), tmp16432)

		tmp16434 := PrimCons(MakeString("s"), tmp16433)

		tmp16435 := PrimCons(MakeString("r"), tmp16434)

		tmp16436 := PrimCons(MakeString("q"), tmp16435)

		tmp16437 := PrimCons(MakeString("p"), tmp16436)

		tmp16438 := PrimCons(MakeString("o"), tmp16437)

		tmp16439 := PrimCons(MakeString("n"), tmp16438)

		tmp16440 := PrimCons(MakeString("m"), tmp16439)

		tmp16441 := PrimCons(MakeString("l"), tmp16440)

		tmp16442 := PrimCons(MakeString("k"), tmp16441)

		tmp16443 := PrimCons(MakeString("j"), tmp16442)

		tmp16444 := PrimCons(MakeString("i"), tmp16443)

		tmp16445 := PrimCons(MakeString("h"), tmp16444)

		tmp16446 := PrimCons(MakeString("g"), tmp16445)

		tmp16447 := PrimCons(MakeString("f"), tmp16446)

		tmp16448 := PrimCons(MakeString("e"), tmp16447)

		tmp16449 := PrimCons(MakeString("d"), tmp16448)

		tmp16450 := PrimCons(MakeString("c"), tmp16449)

		tmp16451 := PrimCons(MakeString("b"), tmp16450)

		tmp16452 := PrimCons(MakeString("a"), tmp16451)

		tmp16453 := PrimCons(MakeString("Z"), tmp16452)

		tmp16454 := PrimCons(MakeString("Y"), tmp16453)

		tmp16455 := PrimCons(MakeString("X"), tmp16454)

		tmp16456 := PrimCons(MakeString("W"), tmp16455)

		tmp16457 := PrimCons(MakeString("V"), tmp16456)

		tmp16458 := PrimCons(MakeString("U"), tmp16457)

		tmp16459 := PrimCons(MakeString("T"), tmp16458)

		tmp16460 := PrimCons(MakeString("S"), tmp16459)

		tmp16461 := PrimCons(MakeString("R"), tmp16460)

		tmp16462 := PrimCons(MakeString("Q"), tmp16461)

		tmp16463 := PrimCons(MakeString("P"), tmp16462)

		tmp16464 := PrimCons(MakeString("O"), tmp16463)

		tmp16465 := PrimCons(MakeString("N"), tmp16464)

		tmp16466 := PrimCons(MakeString("M"), tmp16465)

		tmp16467 := PrimCons(MakeString("L"), tmp16466)

		tmp16468 := PrimCons(MakeString("K"), tmp16467)

		tmp16469 := PrimCons(MakeString("J"), tmp16468)

		tmp16470 := PrimCons(MakeString("I"), tmp16469)

		tmp16471 := PrimCons(MakeString("H"), tmp16470)

		tmp16472 := PrimCons(MakeString("G"), tmp16471)

		tmp16473 := PrimCons(MakeString("F"), tmp16472)

		tmp16474 := PrimCons(MakeString("E"), tmp16473)

		tmp16475 := PrimCons(MakeString("D"), tmp16474)

		tmp16476 := PrimCons(MakeString("C"), tmp16475)

		tmp16477 := PrimCons(MakeString("B"), tmp16476)

		tmp16478 := PrimCons(MakeString("A"), tmp16477)

		__e.TailApply(PrimNS2Value(symelement_2), V1999, tmp16478)
		return

	}, 1)

	tmp16479 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alpha_2, tmp16403)

	_ = tmp16479

	tmp16480 := MakeNative(func(__e *ControlFlow) {
		V2001 := __e.Get(1)
		_ = V2001
		tmp16490 := PrimEqual(MakeString(""), V2001)

		if True == tmp16490 {
			__e.Return(True)
			return
		} else {
			tmp16489 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2001)

			if True == tmp16489 {
				tmp16487 := PrimPos(V2001, MakeNumber(0))

				tmp16488 := Call(__e, PrimNS2Value(symshen_4alphanum_2), tmp16487)

				if True == tmp16488 {
					tmp16485 := PrimTailString(V2001)

					tmp16486 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp16485)

					if True == tmp16486 {
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
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4alphanums_2)
				return
			}

		}

	}, 1)

	tmp16491 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanums_2, tmp16480)

	_ = tmp16491

	tmp16492 := MakeNative(func(__e *ControlFlow) {
		V2003 := __e.Get(1)
		_ = V2003
		tmp16496 := Call(__e, PrimNS2Value(symshen_4alpha_2), V2003)

		if True == tmp16496 {
			__e.Return(True)
			return
		} else {
			tmp16495 := Call(__e, PrimNS2Value(symshen_4digit_2), V2003)

			if True == tmp16495 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp16497 := Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanum_2, tmp16492)

	_ = tmp16497

	tmp16498 := MakeNative(func(__e *ControlFlow) {
		V2005 := __e.Get(1)
		_ = V2005
		tmp16499 := PrimCons(MakeString("0"), Nil)

		tmp16500 := PrimCons(MakeString("9"), tmp16499)

		tmp16501 := PrimCons(MakeString("8"), tmp16500)

		tmp16502 := PrimCons(MakeString("7"), tmp16501)

		tmp16503 := PrimCons(MakeString("6"), tmp16502)

		tmp16504 := PrimCons(MakeString("5"), tmp16503)

		tmp16505 := PrimCons(MakeString("4"), tmp16504)

		tmp16506 := PrimCons(MakeString("3"), tmp16505)

		tmp16507 := PrimCons(MakeString("2"), tmp16506)

		tmp16508 := PrimCons(MakeString("1"), tmp16507)

		__e.TailApply(PrimNS2Value(symelement_2), V2005, tmp16508)
		return

	}, 1)

	tmp16509 := Call(__e, PrimNS1Value(symns2_1set), symshen_4digit_2, tmp16498)

	_ = tmp16509

	tmp16510 := MakeNative(func(__e *ControlFlow) {
		V2007 := __e.Get(1)
		_ = V2007
		tmp16522 := Call(__e, PrimNS2Value(symboolean_2), V2007)

		var ifres16516 Obj

		if True == tmp16522 {
			ifres16516 = True

		} else {
			tmp16521 := PrimIsNumber(V2007)

			var ifres16518 Obj

			if True == tmp16521 {
				ifres16518 = True

			} else {
				tmp16520 := PrimIsString(V2007)

				var ifres16519 Obj

				if True == tmp16520 {
					ifres16519 = True

				} else {
					ifres16519 = False

				}

				ifres16518 = ifres16519

			}

			var ifres16517 Obj

			if True == ifres16518 {
				ifres16517 = True

			} else {
				ifres16517 = False

			}

			ifres16516 = ifres16517

		}

		if True == ifres16516 {
			__e.Return(False)
			return
		} else {
			tmp16512 := MakeNative(func(__e *ControlFlow) {
				tmp16513 := MakeNative(func(__e *ControlFlow) {
					String := __e.Get(1)
					_ = String
					__e.TailApply(PrimNS2Value(symshen_4analyse_1variable_2), String)
					return
				}, 1)

				tmp16514 := PrimStr(V2007)

				__e.TailApply(tmp16513, tmp16514)
				return

			}, 0)

			tmp16515 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(False)
				return
			}, 1)

			__e.TailApply(PrimNS1Value(symtry_1catch), tmp16512, tmp16515)
			return

		}

	}, 1)

	tmp16523 := Call(__e, PrimNS1Value(symns2_1set), symvariable_2, tmp16510)

	_ = tmp16523

	tmp16524 := MakeNative(func(__e *ControlFlow) {
		V2009 := __e.Get(1)
		_ = V2009
		tmp16532 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2009)

		if True == tmp16532 {
			tmp16530 := PrimPos(V2009, MakeNumber(0))

			tmp16531 := Call(__e, PrimNS2Value(symshen_4uppercase_2), tmp16530)

			if True == tmp16531 {
				tmp16528 := PrimTailString(V2009)

				tmp16529 := Call(__e, PrimNS2Value(symshen_4alphanums_2), tmp16528)

				if True == tmp16529 {
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
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4analyse_1variable_2)
			return
		}

	}, 1)

	tmp16533 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1variable_2, tmp16524)

	_ = tmp16533

	tmp16534 := MakeNative(func(__e *ControlFlow) {
		V2011 := __e.Get(1)
		_ = V2011
		tmp16535 := PrimCons(MakeString("Z"), Nil)

		tmp16536 := PrimCons(MakeString("Y"), tmp16535)

		tmp16537 := PrimCons(MakeString("X"), tmp16536)

		tmp16538 := PrimCons(MakeString("W"), tmp16537)

		tmp16539 := PrimCons(MakeString("V"), tmp16538)

		tmp16540 := PrimCons(MakeString("U"), tmp16539)

		tmp16541 := PrimCons(MakeString("T"), tmp16540)

		tmp16542 := PrimCons(MakeString("S"), tmp16541)

		tmp16543 := PrimCons(MakeString("R"), tmp16542)

		tmp16544 := PrimCons(MakeString("Q"), tmp16543)

		tmp16545 := PrimCons(MakeString("P"), tmp16544)

		tmp16546 := PrimCons(MakeString("O"), tmp16545)

		tmp16547 := PrimCons(MakeString("N"), tmp16546)

		tmp16548 := PrimCons(MakeString("M"), tmp16547)

		tmp16549 := PrimCons(MakeString("L"), tmp16548)

		tmp16550 := PrimCons(MakeString("K"), tmp16549)

		tmp16551 := PrimCons(MakeString("J"), tmp16550)

		tmp16552 := PrimCons(MakeString("I"), tmp16551)

		tmp16553 := PrimCons(MakeString("H"), tmp16552)

		tmp16554 := PrimCons(MakeString("G"), tmp16553)

		tmp16555 := PrimCons(MakeString("F"), tmp16554)

		tmp16556 := PrimCons(MakeString("E"), tmp16555)

		tmp16557 := PrimCons(MakeString("D"), tmp16556)

		tmp16558 := PrimCons(MakeString("C"), tmp16557)

		tmp16559 := PrimCons(MakeString("B"), tmp16558)

		tmp16560 := PrimCons(MakeString("A"), tmp16559)

		__e.TailApply(PrimNS2Value(symelement_2), V2011, tmp16560)
		return

	}, 1)

	tmp16561 := Call(__e, PrimNS1Value(symns2_1set), symshen_4uppercase_2, tmp16534)

	_ = tmp16561

	tmp16562 := MakeNative(func(__e *ControlFlow) {
		V2013 := __e.Get(1)
		_ = V2013
		tmp16563 := PrimNS3Value(symshen_4_dgensym_d)

		tmp16564 := PrimNumberAdd(MakeNumber(1), tmp16563)

		tmp16565 := PrimNS3Set(symshen_4_dgensym_d, tmp16564)

		__e.TailApply(PrimNS2Value(symconcat), V2013, tmp16565)
		return

	}, 1)

	tmp16566 := Call(__e, PrimNS1Value(symns2_1set), symgensym, tmp16562)

	_ = tmp16566

	tmp16567 := MakeNative(func(__e *ControlFlow) {
		V2016 := __e.Get(1)
		_ = V2016
		V2017 := __e.Get(2)
		_ = V2017
		tmp16568 := PrimStr(V2016)

		tmp16569 := PrimStr(V2017)

		tmp16570 := PrimStringConcat(tmp16568, tmp16569)

		__e.Return(PrimIntern(tmp16570))
		return

	}, 2)

	tmp16571 := Call(__e, PrimNS1Value(symns2_1set), symconcat, tmp16567)

	_ = tmp16571

	tmp16572 := MakeNative(func(__e *ControlFlow) {
		V2020 := __e.Get(1)
		_ = V2020
		V2021 := __e.Get(2)
		_ = V2021
		tmp16573 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			tmp16574 := MakeNative(func(__e *ControlFlow) {
				Tag := __e.Get(1)
				_ = Tag
				tmp16575 := MakeNative(func(__e *ControlFlow) {
					Fst := __e.Get(1)
					_ = Fst
					tmp16576 := MakeNative(func(__e *ControlFlow) {
						Snd := __e.Get(1)
						_ = Snd
						__e.Return(Vector)
						return
					}, 1)

					tmp16577 := PrimVectorSet(Vector, MakeNumber(2), V2021)

					__e.TailApply(tmp16576, tmp16577)
					return

				}, 1)

				tmp16578 := PrimVectorSet(Vector, MakeNumber(1), V2020)

				__e.TailApply(tmp16575, tmp16578)
				return

			}, 1)

			tmp16579 := PrimVectorSet(Vector, MakeNumber(0), symshen_4tuple)

			__e.TailApply(tmp16574, tmp16579)
			return

		}, 1)

		tmp16580 := PrimAbsvector(MakeNumber(3))

		__e.TailApply(tmp16573, tmp16580)
		return

	}, 2)

	tmp16581 := Call(__e, PrimNS1Value(symns2_1set), sym_8p, tmp16572)

	_ = tmp16581

	tmp16582 := MakeNative(func(__e *ControlFlow) {
		V2023 := __e.Get(1)
		_ = V2023
		__e.Return(PrimVectorGet(V2023, MakeNumber(1)))
		return
	}, 1)

	tmp16583 := Call(__e, PrimNS1Value(symns2_1set), symfst, tmp16582)

	_ = tmp16583

	tmp16584 := MakeNative(func(__e *ControlFlow) {
		V2025 := __e.Get(1)
		_ = V2025
		__e.Return(PrimVectorGet(V2025, MakeNumber(2)))
		return
	}, 1)

	tmp16585 := Call(__e, PrimNS1Value(symns2_1set), symsnd, tmp16584)

	_ = tmp16585

	tmp16586 := MakeNative(func(__e *ControlFlow) {
		V2027 := __e.Get(1)
		_ = V2027
		tmp16593 := PrimIsVector(V2027)

		if True == tmp16593 {
			tmp16589 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V2027, MakeNumber(0)))
				return
			}, 0)

			tmp16590 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1tuple)
				return
			}, 1)

			tmp16591 := Call(__e, PrimNS1Value(symtry_1catch), tmp16589, tmp16590)

			tmp16592 := PrimEqual(symshen_4tuple, tmp16591)

			if True == tmp16592 {
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

	tmp16594 := Call(__e, PrimNS1Value(symns2_1set), symtuple_2, tmp16586)

	_ = tmp16594

	tmp16595 := MakeNative(func(__e *ControlFlow) {
		V2030 := __e.Get(1)
		_ = V2030
		V2031 := __e.Get(2)
		_ = V2031
		tmp16602 := PrimEqual(Nil, V2030)

		if True == tmp16602 {
			__e.Return(V2031)
			return
		} else {
			tmp16601 := PrimIsPair(V2030)

			if True == tmp16601 {
				tmp16598 := PrimHead(V2030)

				tmp16599 := PrimTail(V2030)

				tmp16600 := Call(__e, PrimNS2Value(symappend), tmp16599, V2031)

				__e.Return(PrimCons(tmp16598, tmp16600))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symappend)
				return
			}

		}

	}, 2)

	tmp16603 := Call(__e, PrimNS1Value(symns2_1set), symappend, tmp16595)

	_ = tmp16603

	tmp16604 := MakeNative(func(__e *ControlFlow) {
		V2034 := __e.Get(1)
		_ = V2034
		V2035 := __e.Get(2)
		_ = V2035
		tmp16605 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp16606 := MakeNative(func(__e *ControlFlow) {
				NewVector := __e.Get(1)
				_ = NewVector
				tmp16607 := MakeNative(func(__e *ControlFlow) {
					X_7NewVector := __e.Get(1)
					_ = X_7NewVector
					tmp16609 := PrimEqual(Limit, MakeNumber(0))

					if True == tmp16609 {
						__e.Return(X_7NewVector)
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4_8v_1help), V2035, MakeNumber(1), Limit, X_7NewVector)
						return
					}

				}, 1)

				tmp16610 := Call(__e, PrimNS2Value(symvector_1_6), NewVector, MakeNumber(1), V2034)

				__e.TailApply(tmp16607, tmp16610)
				return

			}, 1)

			tmp16611 := PrimNumberAdd(Limit, MakeNumber(1))

			tmp16612 := Call(__e, PrimNS2Value(symvector), tmp16611)

			__e.TailApply(tmp16606, tmp16612)
			return

		}, 1)

		tmp16613 := Call(__e, PrimNS2Value(symlimit), V2035)

		__e.TailApply(tmp16605, tmp16613)
		return

	}, 2)

	tmp16614 := Call(__e, PrimNS1Value(symns2_1set), sym_8v, tmp16604)

	_ = tmp16614

	tmp16615 := MakeNative(func(__e *ControlFlow) {
		V2041 := __e.Get(1)
		_ = V2041
		V2042 := __e.Get(2)
		_ = V2042
		V2043 := __e.Get(3)
		_ = V2043
		V2044 := __e.Get(4)
		_ = V2044
		tmp16621 := PrimEqual(V2043, V2042)

		if True == tmp16621 {
			tmp16620 := PrimNumberAdd(V2043, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4copyfromvector), V2041, V2044, V2043, tmp16620)
			return

		} else {
			tmp16617 := PrimNumberAdd(V2042, MakeNumber(1))

			tmp16618 := PrimNumberAdd(V2042, MakeNumber(1))

			tmp16619 := Call(__e, PrimNS2Value(symshen_4copyfromvector), V2041, V2044, V2042, tmp16618)

			__e.TailApply(PrimNS2Value(symshen_4_8v_1help), V2041, tmp16617, V2043, tmp16619)
			return

		}

	}, 4)

	tmp16622 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_8v_1help, tmp16615)

	_ = tmp16622

	tmp16623 := MakeNative(func(__e *ControlFlow) {
		V2049 := __e.Get(1)
		_ = V2049
		V2050 := __e.Get(2)
		_ = V2050
		V2051 := __e.Get(3)
		_ = V2051
		V2052 := __e.Get(4)
		_ = V2052
		tmp16624 := MakeNative(func(__e *ControlFlow) {
			tmp16625 := Call(__e, PrimNS2Value(sym_5_1vector), V2049, V2051)

			__e.TailApply(PrimNS2Value(symvector_1_6), V2050, V2052, tmp16625)
			return

		}, 0)

		tmp16626 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V2050)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp16624, tmp16626)
		return

	}, 4)

	tmp16627 := Call(__e, PrimNS1Value(symns2_1set), symshen_4copyfromvector, tmp16623)

	_ = tmp16627

	tmp16628 := MakeNative(func(__e *ControlFlow) {
		V2054 := __e.Get(1)
		_ = V2054
		tmp16629 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(sym_5_1vector), V2054, MakeNumber(1))
			return
		}, 0)

		tmp16630 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp16631 := Call(__e, PrimNS2Value(symshen_4app), V2054, MakeString("\n"), symshen_4s)

			tmp16632 := PrimStringConcat(MakeString("hdv needs a non-empty vector as an argument; not "), tmp16631)

			__e.Return(PrimSimpleError(tmp16632))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp16629, tmp16630)
		return

	}, 1)

	tmp16633 := Call(__e, PrimNS1Value(symns2_1set), symhdv, tmp16628)

	_ = tmp16633

	tmp16634 := MakeNative(func(__e *ControlFlow) {
		V2056 := __e.Get(1)
		_ = V2056
		tmp16635 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			tmp16644 := PrimEqual(Limit, MakeNumber(0))

			if True == tmp16644 {
				__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
				return
			} else {
				tmp16643 := PrimEqual(Limit, MakeNumber(1))

				if True == tmp16643 {
					__e.TailApply(PrimNS2Value(symvector), MakeNumber(0))
					return
				} else {
					tmp16638 := MakeNative(func(__e *ControlFlow) {
						NewVector := __e.Get(1)
						_ = NewVector
						tmp16639 := PrimNumberSubtract(Limit, MakeNumber(1))

						tmp16640 := Call(__e, PrimNS2Value(symvector), tmp16639)

						__e.TailApply(PrimNS2Value(symshen_4tlv_1help), V2056, MakeNumber(2), Limit, tmp16640)
						return

					}, 1)

					tmp16641 := PrimNumberSubtract(Limit, MakeNumber(1))

					tmp16642 := Call(__e, PrimNS2Value(symvector), tmp16641)

					__e.TailApply(tmp16638, tmp16642)
					return

				}

			}

		}, 1)

		tmp16645 := Call(__e, PrimNS2Value(symlimit), V2056)

		__e.TailApply(tmp16635, tmp16645)
		return

	}, 1)

	tmp16646 := Call(__e, PrimNS1Value(symns2_1set), symtlv, tmp16634)

	_ = tmp16646

	tmp16647 := MakeNative(func(__e *ControlFlow) {
		V2062 := __e.Get(1)
		_ = V2062
		V2063 := __e.Get(2)
		_ = V2063
		V2064 := __e.Get(3)
		_ = V2064
		V2065 := __e.Get(4)
		_ = V2065
		tmp16653 := PrimEqual(V2064, V2063)

		if True == tmp16653 {
			tmp16652 := PrimNumberSubtract(V2064, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4copyfromvector), V2062, V2065, V2064, tmp16652)
			return

		} else {
			tmp16649 := PrimNumberAdd(V2063, MakeNumber(1))

			tmp16650 := PrimNumberSubtract(V2063, MakeNumber(1))

			tmp16651 := Call(__e, PrimNS2Value(symshen_4copyfromvector), V2062, V2065, V2063, tmp16650)

			__e.TailApply(PrimNS2Value(symshen_4tlv_1help), V2062, tmp16649, V2064, tmp16651)
			return

		}

	}, 4)

	tmp16654 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tlv_1help, tmp16647)

	_ = tmp16654

	tmp16655 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		V2078 := __e.Get(2)
		_ = V2078
		tmp16671 := PrimEqual(Nil, V2078)

		if True == tmp16671 {
			__e.Return(Nil)
			return
		} else {
			tmp16670 := PrimIsPair(V2078)

			var ifres16661 Obj

			if True == tmp16670 {
				tmp16668 := PrimHead(V2078)

				tmp16669 := PrimIsPair(tmp16668)

				var ifres16663 Obj

				if True == tmp16669 {
					tmp16665 := PrimHead(V2078)

					tmp16666 := PrimHead(tmp16665)

					tmp16667 := PrimEqual(tmp16666, V2077)

					var ifres16664 Obj

					if True == tmp16667 {
						ifres16664 = True

					} else {
						ifres16664 = False

					}

					ifres16663 = ifres16664

				} else {
					ifres16663 = False

				}

				var ifres16662 Obj

				if True == ifres16663 {
					ifres16662 = True

				} else {
					ifres16662 = False

				}

				ifres16661 = ifres16662

			} else {
				ifres16661 = False

			}

			if True == ifres16661 {
				__e.Return(PrimHead(V2078))
				return
			} else {
				tmp16660 := PrimIsPair(V2078)

				if True == tmp16660 {
					tmp16659 := PrimTail(V2078)

					__e.TailApply(PrimNS2Value(symassoc), V2077, tmp16659)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symassoc)
					return
				}

			}

		}

	}, 2)

	tmp16672 := Call(__e, PrimNS1Value(symns2_1set), symassoc, tmp16655)

	_ = tmp16672

	tmp16673 := MakeNative(func(__e *ControlFlow) {
		V2085 := __e.Get(1)
		_ = V2085
		V2086 := __e.Get(2)
		_ = V2086
		V2087 := __e.Get(3)
		_ = V2087
		tmp16696 := PrimEqual(Nil, V2087)

		if True == tmp16696 {
			tmp16695 := PrimCons(V2085, V2086)

			__e.Return(PrimCons(tmp16695, Nil))
			return

		} else {
			tmp16694 := PrimIsPair(V2087)

			var ifres16685 Obj

			if True == tmp16694 {
				tmp16692 := PrimHead(V2087)

				tmp16693 := PrimIsPair(tmp16692)

				var ifres16687 Obj

				if True == tmp16693 {
					tmp16689 := PrimHead(V2087)

					tmp16690 := PrimHead(tmp16689)

					tmp16691 := PrimEqual(tmp16690, V2085)

					var ifres16688 Obj

					if True == tmp16691 {
						ifres16688 = True

					} else {
						ifres16688 = False

					}

					ifres16687 = ifres16688

				} else {
					ifres16687 = False

				}

				var ifres16686 Obj

				if True == ifres16687 {
					ifres16686 = True

				} else {
					ifres16686 = False

				}

				ifres16685 = ifres16686

			} else {
				ifres16685 = False

			}

			if True == ifres16685 {
				tmp16681 := PrimHead(V2087)

				tmp16682 := PrimHead(tmp16681)

				tmp16683 := PrimCons(tmp16682, V2086)

				tmp16684 := PrimTail(V2087)

				__e.Return(PrimCons(tmp16683, tmp16684))
				return

			} else {
				tmp16680 := PrimIsPair(V2087)

				if True == tmp16680 {
					tmp16677 := PrimHead(V2087)

					tmp16678 := PrimTail(V2087)

					tmp16679 := Call(__e, PrimNS2Value(symshen_4assoc_1set), V2085, V2086, tmp16678)

					__e.Return(PrimCons(tmp16677, tmp16679))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4assoc_1set)
					return
				}

			}

		}

	}, 3)

	tmp16697 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1set, tmp16673)

	_ = tmp16697

	tmp16698 := MakeNative(func(__e *ControlFlow) {
		V2093 := __e.Get(1)
		_ = V2093
		V2094 := __e.Get(2)
		_ = V2094
		tmp16716 := PrimEqual(Nil, V2094)

		if True == tmp16716 {
			__e.Return(Nil)
			return
		} else {
			tmp16715 := PrimIsPair(V2094)

			var ifres16706 Obj

			if True == tmp16715 {
				tmp16713 := PrimHead(V2094)

				tmp16714 := PrimIsPair(tmp16713)

				var ifres16708 Obj

				if True == tmp16714 {
					tmp16710 := PrimHead(V2094)

					tmp16711 := PrimHead(tmp16710)

					tmp16712 := PrimEqual(tmp16711, V2093)

					var ifres16709 Obj

					if True == tmp16712 {
						ifres16709 = True

					} else {
						ifres16709 = False

					}

					ifres16708 = ifres16709

				} else {
					ifres16708 = False

				}

				var ifres16707 Obj

				if True == ifres16708 {
					ifres16707 = True

				} else {
					ifres16707 = False

				}

				ifres16706 = ifres16707

			} else {
				ifres16706 = False

			}

			if True == ifres16706 {
				__e.Return(PrimTail(V2094))
				return
			} else {
				tmp16705 := PrimIsPair(V2094)

				if True == tmp16705 {
					tmp16702 := PrimHead(V2094)

					tmp16703 := PrimTail(V2094)

					tmp16704 := Call(__e, PrimNS2Value(symshen_4assoc_1rm), V2093, tmp16703)

					__e.Return(PrimCons(tmp16702, tmp16704))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4assoc_1rm)
					return
				}

			}

		}

	}, 2)

	tmp16717 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1rm, tmp16698)

	_ = tmp16717

	tmp16718 := MakeNative(func(__e *ControlFlow) {
		V2100 := __e.Get(1)
		_ = V2100
		tmp16722 := PrimEqual(True, V2100)

		if True == tmp16722 {
			__e.Return(True)
			return
		} else {
			tmp16721 := PrimEqual(False, V2100)

			if True == tmp16721 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp16723 := Call(__e, PrimNS1Value(symns2_1set), symboolean_2, tmp16718)

	_ = tmp16723

	tmp16724 := MakeNative(func(__e *ControlFlow) {
		V2102 := __e.Get(1)
		_ = V2102
		tmp16729 := PrimEqual(MakeNumber(0), V2102)

		if True == tmp16729 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp16726 := Call(__e, PrimNS2Value(symstoutput))

			tmp16727 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString("\n"), tmp16726)

			_ = tmp16727

			tmp16728 := PrimNumberSubtract(V2102, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symnl), tmp16728)
			return

		}

	}, 1)

	tmp16730 := Call(__e, PrimNS1Value(symns2_1set), symnl, tmp16724)

	_ = tmp16730

	tmp16731 := MakeNative(func(__e *ControlFlow) {
		V2107 := __e.Get(1)
		_ = V2107
		V2108 := __e.Get(2)
		_ = V2108
		tmp16742 := PrimEqual(Nil, V2107)

		if True == tmp16742 {
			__e.Return(Nil)
			return
		} else {
			tmp16741 := PrimIsPair(V2107)

			if True == tmp16741 {
				tmp16739 := PrimHead(V2107)

				tmp16740 := Call(__e, PrimNS2Value(symelement_2), tmp16739, V2108)

				if True == tmp16740 {
					tmp16738 := PrimTail(V2107)

					__e.TailApply(PrimNS2Value(symdifference), tmp16738, V2108)
					return

				} else {
					tmp16735 := PrimHead(V2107)

					tmp16736 := PrimTail(V2107)

					tmp16737 := Call(__e, PrimNS2Value(symdifference), tmp16736, V2108)

					__e.Return(PrimCons(tmp16735, tmp16737))
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symdifference)
				return
			}

		}

	}, 2)

	tmp16743 := Call(__e, PrimNS1Value(symns2_1set), symdifference, tmp16731)

	_ = tmp16743

	tmp16744 := MakeNative(func(__e *ControlFlow) {
		V2111 := __e.Get(1)
		_ = V2111
		V2112 := __e.Get(2)
		_ = V2112
		__e.Return(V2112)
		return
	}, 2)

	tmp16745 := Call(__e, PrimNS1Value(symns2_1set), symdo, tmp16744)

	_ = tmp16745

	tmp16746 := MakeNative(func(__e *ControlFlow) {
		V2124 := __e.Get(1)
		_ = V2124
		V2125 := __e.Get(2)
		_ = V2125
		tmp16757 := PrimEqual(Nil, V2125)

		if True == tmp16757 {
			__e.Return(False)
			return
		} else {
			tmp16756 := PrimIsPair(V2125)

			var ifres16752 Obj

			if True == tmp16756 {
				tmp16754 := PrimHead(V2125)

				tmp16755 := PrimEqual(tmp16754, V2124)

				var ifres16753 Obj

				if True == tmp16755 {
					ifres16753 = True

				} else {
					ifres16753 = False

				}

				ifres16752 = ifres16753

			} else {
				ifres16752 = False

			}

			if True == ifres16752 {
				__e.Return(True)
				return
			} else {
				tmp16751 := PrimIsPair(V2125)

				if True == tmp16751 {
					tmp16750 := PrimTail(V2125)

					__e.TailApply(PrimNS2Value(symelement_2), V2124, tmp16750)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symelement_2)
					return
				}

			}

		}

	}, 2)

	tmp16758 := Call(__e, PrimNS1Value(symns2_1set), symelement_2, tmp16746)

	_ = tmp16758

	tmp16759 := MakeNative(func(__e *ControlFlow) {
		V2131 := __e.Get(1)
		_ = V2131
		tmp16761 := PrimEqual(Nil, V2131)

		if True == tmp16761 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp16762 := Call(__e, PrimNS1Value(symns2_1set), symempty_2, tmp16759)

	_ = tmp16762

	tmp16763 := MakeNative(func(__e *ControlFlow) {
		V2134 := __e.Get(1)
		_ = V2134
		V2135 := __e.Get(2)
		_ = V2135
		tmp16764 := Call(__e, V2134, V2135)

		__e.TailApply(PrimNS2Value(symshen_4fix_1help), V2134, V2135, tmp16764)
		return

	}, 2)

	tmp16765 := Call(__e, PrimNS1Value(symns2_1set), symfix, tmp16763)

	_ = tmp16765

	tmp16766 := MakeNative(func(__e *ControlFlow) {
		V2146 := __e.Get(1)
		_ = V2146
		V2147 := __e.Get(2)
		_ = V2147
		V2148 := __e.Get(3)
		_ = V2148
		tmp16769 := PrimEqual(V2148, V2147)

		if True == tmp16769 {
			__e.Return(V2148)
			return
		} else {
			tmp16768 := Call(__e, V2146, V2148)

			__e.TailApply(PrimNS2Value(symshen_4fix_1help), V2146, V2148, tmp16768)
			return

		}

	}, 3)

	tmp16770 := Call(__e, PrimNS1Value(symns2_1set), symshen_4fix_1help, tmp16766)

	_ = tmp16770

	tmp16771 := MakeNative(func(__e *ControlFlow) {
		V2153 := __e.Get(1)
		_ = V2153
		V2154 := __e.Get(2)
		_ = V2154
		V2155 := __e.Get(3)
		_ = V2155
		V2156 := __e.Get(4)
		_ = V2156
		tmp16772 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			tmp16773 := MakeNative(func(__e *ControlFlow) {
				Added := __e.Get(1)
				_ = Added
				tmp16774 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2155)
					return
				}, 1)

				tmp16775 := Call(__e, PrimNS2Value(symshen_4dict_1_6), V2156, V2153, Added)

				__e.TailApply(tmp16774, tmp16775)
				return

			}, 1)

			tmp16776 := Call(__e, PrimNS2Value(symshen_4assoc_1set), V2154, V2155, Curr)

			__e.TailApply(tmp16773, tmp16776)
			return

		}, 1)

		tmp16777 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4_5_1dict), V2156, V2153)
			return
		}, 0)

		tmp16778 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp16779 := Call(__e, PrimNS1Value(symtry_1catch), tmp16777, tmp16778)

		__e.TailApply(tmp16772, tmp16779)
		return

	}, 4)

	tmp16780 := Call(__e, PrimNS1Value(symns2_1set), symput, tmp16771)

	_ = tmp16780

	tmp16781 := MakeNative(func(__e *ControlFlow) {
		V2160 := __e.Get(1)
		_ = V2160
		V2161 := __e.Get(2)
		_ = V2161
		V2162 := __e.Get(3)
		_ = V2162
		tmp16782 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			tmp16783 := MakeNative(func(__e *ControlFlow) {
				Removed := __e.Get(1)
				_ = Removed
				tmp16784 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2160)
					return
				}, 1)

				tmp16785 := Call(__e, PrimNS2Value(symshen_4dict_1_6), V2162, V2160, Removed)

				__e.TailApply(tmp16784, tmp16785)
				return

			}, 1)

			tmp16786 := Call(__e, PrimNS2Value(symshen_4assoc_1rm), V2161, Curr)

			__e.TailApply(tmp16783, tmp16786)
			return

		}, 1)

		tmp16787 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4_5_1dict), V2162, V2160)
			return
		}, 0)

		tmp16788 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp16789 := Call(__e, PrimNS1Value(symtry_1catch), tmp16787, tmp16788)

		__e.TailApply(tmp16782, tmp16789)
		return

	}, 3)

	tmp16790 := Call(__e, PrimNS1Value(symns2_1set), symunput, tmp16781)

	_ = tmp16790

	tmp16791 := MakeNative(func(__e *ControlFlow) {
		V2166 := __e.Get(1)
		_ = V2166
		V2167 := __e.Get(2)
		_ = V2167
		V2168 := __e.Get(3)
		_ = V2168
		tmp16792 := MakeNative(func(__e *ControlFlow) {
			Entry := __e.Get(1)
			_ = Entry
			tmp16793 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				tmp16795 := Call(__e, PrimNS2Value(symempty_2), Result)

				if True == tmp16795 {
					__e.Return(PrimSimpleError(MakeString("value not found\n")))
					return
				} else {
					__e.Return(PrimTail(Result))
					return
				}

			}, 1)

			tmp16796 := Call(__e, PrimNS2Value(symassoc), V2167, Entry)

			__e.TailApply(tmp16793, tmp16796)
			return

		}, 1)

		tmp16797 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4_5_1dict), V2168, V2166)
			return
		}, 0)

		tmp16798 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp16799 := Call(__e, PrimNS1Value(symtry_1catch), tmp16797, tmp16798)

		__e.TailApply(tmp16792, tmp16799)
		return

	}, 3)

	tmp16800 := Call(__e, PrimNS1Value(symns2_1set), symget, tmp16791)

	_ = tmp16800

	tmp16801 := MakeNative(func(__e *ControlFlow) {
		V2171 := __e.Get(1)
		_ = V2171
		V2172 := __e.Get(2)
		_ = V2172
		tmp16802 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(PrimStringToNumber(X))
			return
		}, 1)

		tmp16803 := Call(__e, PrimNS2Value(symexplode), V2171)

		tmp16804 := Call(__e, PrimNS2Value(symmap), tmp16802, tmp16803)

		tmp16805 := Call(__e, PrimNS2Value(symsum), tmp16804)

		__e.TailApply(PrimNS2Value(symshen_4mod), tmp16805, V2172)
		return

	}, 2)

	tmp16806 := Call(__e, PrimNS1Value(symns2_1set), symhash, tmp16801)

	_ = tmp16806

	tmp16807 := MakeNative(func(__e *ControlFlow) {
		V2175 := __e.Get(1)
		_ = V2175
		V2176 := __e.Get(2)
		_ = V2176
		tmp16808 := PrimCons(V2176, Nil)

		tmp16809 := Call(__e, PrimNS2Value(symshen_4multiples), V2175, tmp16808)

		__e.TailApply(PrimNS2Value(symshen_4modh), V2175, tmp16809)
		return

	}, 2)

	tmp16810 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mod, tmp16807)

	_ = tmp16810

	tmp16811 := MakeNative(func(__e *ControlFlow) {
		V2179 := __e.Get(1)
		_ = V2179
		V2180 := __e.Get(2)
		_ = V2180
		tmp16822 := PrimIsPair(V2180)

		var ifres16818 Obj

		if True == tmp16822 {
			tmp16820 := PrimHead(V2180)

			tmp16821 := PrimGreatThan(tmp16820, V2179)

			var ifres16819 Obj

			if True == tmp16821 {
				ifres16819 = True

			} else {
				ifres16819 = False

			}

			ifres16818 = ifres16819

		} else {
			ifres16818 = False

		}

		if True == ifres16818 {
			__e.Return(PrimTail(V2180))
			return
		} else {
			tmp16817 := PrimIsPair(V2180)

			if True == tmp16817 {
				tmp16814 := PrimHead(V2180)

				tmp16815 := PrimNumberMultiply(MakeNumber(2), tmp16814)

				tmp16816 := PrimCons(tmp16815, V2180)

				__e.TailApply(PrimNS2Value(symshen_4multiples), V2179, tmp16816)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4multiples)
				return
			}

		}

	}, 2)

	tmp16823 := Call(__e, PrimNS1Value(symns2_1set), symshen_4multiples, tmp16811)

	_ = tmp16823

	tmp16824 := MakeNative(func(__e *ControlFlow) {
		V2185 := __e.Get(1)
		_ = V2185
		V2186 := __e.Get(2)
		_ = V2186
		tmp16842 := PrimEqual(MakeNumber(0), V2185)

		if True == tmp16842 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp16841 := PrimEqual(Nil, V2186)

			if True == tmp16841 {
				__e.Return(V2185)
				return
			} else {
				tmp16840 := PrimIsPair(V2186)

				var ifres16836 Obj

				if True == tmp16840 {
					tmp16838 := PrimHead(V2186)

					tmp16839 := PrimGreatThan(tmp16838, V2185)

					var ifres16837 Obj

					if True == tmp16839 {
						ifres16837 = True

					} else {
						ifres16837 = False

					}

					ifres16836 = ifres16837

				} else {
					ifres16836 = False

				}

				if True == ifres16836 {
					tmp16834 := PrimTail(V2186)

					tmp16835 := Call(__e, PrimNS2Value(symempty_2), tmp16834)

					if True == tmp16835 {
						__e.Return(V2185)
						return
					} else {
						tmp16833 := PrimTail(V2186)

						__e.TailApply(PrimNS2Value(symshen_4modh), V2185, tmp16833)
						return

					}

				} else {
					tmp16831 := PrimIsPair(V2186)

					if True == tmp16831 {
						tmp16829 := PrimHead(V2186)

						tmp16830 := PrimNumberSubtract(V2185, tmp16829)

						__e.TailApply(PrimNS2Value(symshen_4modh), tmp16830, V2186)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4modh)
						return
					}

				}

			}

		}

	}, 2)

	tmp16843 := Call(__e, PrimNS1Value(symns2_1set), symshen_4modh, tmp16824)

	_ = tmp16843

	tmp16844 := MakeNative(func(__e *ControlFlow) {
		V2188 := __e.Get(1)
		_ = V2188
		tmp16851 := PrimEqual(Nil, V2188)

		if True == tmp16851 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp16850 := PrimIsPair(V2188)

			if True == tmp16850 {
				tmp16847 := PrimHead(V2188)

				tmp16848 := PrimTail(V2188)

				tmp16849 := Call(__e, PrimNS2Value(symsum), tmp16848)

				__e.Return(PrimNumberAdd(tmp16847, tmp16849))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symsum)
				return
			}

		}

	}, 1)

	tmp16852 := Call(__e, PrimNS1Value(symns2_1set), symsum, tmp16844)

	_ = tmp16852

	tmp16853 := MakeNative(func(__e *ControlFlow) {
		V2196 := __e.Get(1)
		_ = V2196
		tmp16855 := PrimIsPair(V2196)

		if True == tmp16855 {
			__e.Return(PrimHead(V2196))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("head expects a non-empty list")))
			return
		}

	}, 1)

	tmp16856 := Call(__e, PrimNS1Value(symns2_1set), symhead, tmp16853)

	_ = tmp16856

	tmp16857 := MakeNative(func(__e *ControlFlow) {
		V2204 := __e.Get(1)
		_ = V2204
		tmp16859 := PrimIsPair(V2204)

		if True == tmp16859 {
			__e.Return(PrimTail(V2204))
			return
		} else {
			__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list")))
			return
		}

	}, 1)

	tmp16860 := Call(__e, PrimNS1Value(symns2_1set), symtail, tmp16857)

	_ = tmp16860

	tmp16861 := MakeNative(func(__e *ControlFlow) {
		V2206 := __e.Get(1)
		_ = V2206
		__e.Return(PrimPos(V2206, MakeNumber(0)))
		return
	}, 1)

	tmp16862 := Call(__e, PrimNS1Value(symns2_1set), symhdstr, tmp16861)

	_ = tmp16862

	tmp16863 := MakeNative(func(__e *ControlFlow) {
		V2211 := __e.Get(1)
		_ = V2211
		V2212 := __e.Get(2)
		_ = V2212
		tmp16874 := PrimEqual(Nil, V2211)

		if True == tmp16874 {
			__e.Return(Nil)
			return
		} else {
			tmp16873 := PrimIsPair(V2211)

			if True == tmp16873 {
				tmp16871 := PrimHead(V2211)

				tmp16872 := Call(__e, PrimNS2Value(symelement_2), tmp16871, V2212)

				if True == tmp16872 {
					tmp16868 := PrimHead(V2211)

					tmp16869 := PrimTail(V2211)

					tmp16870 := Call(__e, PrimNS2Value(symintersection), tmp16869, V2212)

					__e.Return(PrimCons(tmp16868, tmp16870))
					return

				} else {
					tmp16867 := PrimTail(V2211)

					__e.TailApply(PrimNS2Value(symintersection), tmp16867, V2212)
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symintersection)
				return
			}

		}

	}, 2)

	tmp16875 := Call(__e, PrimNS1Value(symns2_1set), symintersection, tmp16863)

	_ = tmp16875

	tmp16876 := MakeNative(func(__e *ControlFlow) {
		V2214 := __e.Get(1)
		_ = V2214
		__e.TailApply(PrimNS2Value(symshen_4reverse__help), V2214, Nil)
		return
	}, 1)

	tmp16877 := Call(__e, PrimNS1Value(symns2_1set), symreverse, tmp16876)

	_ = tmp16877

	tmp16878 := MakeNative(func(__e *ControlFlow) {
		V2217 := __e.Get(1)
		_ = V2217
		V2218 := __e.Get(2)
		_ = V2218
		tmp16885 := PrimEqual(Nil, V2217)

		if True == tmp16885 {
			__e.Return(V2218)
			return
		} else {
			tmp16884 := PrimIsPair(V2217)

			if True == tmp16884 {
				tmp16881 := PrimTail(V2217)

				tmp16882 := PrimHead(V2217)

				tmp16883 := PrimCons(tmp16882, V2218)

				__e.TailApply(PrimNS2Value(symshen_4reverse__help), tmp16881, tmp16883)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4reverse__help)
				return
			}

		}

	}, 2)

	tmp16886 := Call(__e, PrimNS1Value(symns2_1set), symshen_4reverse__help, tmp16878)

	_ = tmp16886

	tmp16887 := MakeNative(func(__e *ControlFlow) {
		V2221 := __e.Get(1)
		_ = V2221
		V2222 := __e.Get(2)
		_ = V2222
		tmp16898 := PrimEqual(Nil, V2221)

		if True == tmp16898 {
			__e.Return(V2222)
			return
		} else {
			tmp16897 := PrimIsPair(V2221)

			if True == tmp16897 {
				tmp16895 := PrimHead(V2221)

				tmp16896 := Call(__e, PrimNS2Value(symelement_2), tmp16895, V2222)

				if True == tmp16896 {
					tmp16894 := PrimTail(V2221)

					__e.TailApply(PrimNS2Value(symunion), tmp16894, V2222)
					return

				} else {
					tmp16891 := PrimHead(V2221)

					tmp16892 := PrimTail(V2221)

					tmp16893 := Call(__e, PrimNS2Value(symunion), tmp16892, V2222)

					__e.Return(PrimCons(tmp16891, tmp16893))
					return

				}

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symunion)
				return
			}

		}

	}, 2)

	tmp16899 := Call(__e, PrimNS1Value(symns2_1set), symunion, tmp16887)

	_ = tmp16899

	tmp16900 := MakeNative(func(__e *ControlFlow) {
		V2224 := __e.Get(1)
		_ = V2224
		tmp16901 := MakeNative(func(__e *ControlFlow) {
			Message := __e.Get(1)
			_ = Message
			tmp16902 := MakeNative(func(__e *ControlFlow) {
				Y_1or_1N := __e.Get(1)
				_ = Y_1or_1N
				tmp16903 := MakeNative(func(__e *ControlFlow) {
					Input := __e.Get(1)
					_ = Input
					tmp16909 := PrimEqual(MakeString("y"), Input)

					if True == tmp16909 {
						__e.Return(True)
						return
					} else {
						tmp16908 := PrimEqual(MakeString("n"), Input)

						if True == tmp16908 {
							__e.Return(False)
							return
						} else {
							tmp16906 := Call(__e, PrimNS2Value(symstoutput))

							tmp16907 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString("please answer y or n\n"), tmp16906)

							_ = tmp16907

							__e.TailApply(PrimNS2Value(symy_1or_1n_2), V2224)
							return

						}

					}

				}, 1)

				tmp16910 := Call(__e, PrimNS2Value(symstinput))

				tmp16911 := Call(__e, PrimNS2Value(symread), tmp16910)

				tmp16912 := Call(__e, PrimNS2Value(symshen_4app), tmp16911, MakeString(""), symshen_4s)

				__e.TailApply(tmp16903, tmp16912)
				return

			}, 1)

			tmp16913 := Call(__e, PrimNS2Value(symstoutput))

			tmp16914 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString(" (y/n) "), tmp16913)

			__e.TailApply(tmp16902, tmp16914)
			return

		}, 1)

		tmp16915 := Call(__e, PrimNS2Value(symshen_4proc_1nl), V2224)

		tmp16916 := Call(__e, PrimNS2Value(symstoutput))

		tmp16917 := Call(__e, PrimNS2Value(symshen_4prhush), tmp16915, tmp16916)

		__e.TailApply(tmp16901, tmp16917)
		return

	}, 1)

	tmp16918 := Call(__e, PrimNS1Value(symns2_1set), symy_1or_1n_2, tmp16900)

	_ = tmp16918

	tmp16919 := MakeNative(func(__e *ControlFlow) {
		V2226 := __e.Get(1)
		_ = V2226
		if True == V2226 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}
	}, 1)

	tmp16921 := Call(__e, PrimNS1Value(symns2_1set), symnot, tmp16919)

	_ = tmp16921

	tmp16922 := MakeNative(func(__e *ControlFlow) {
		V2239 := __e.Get(1)
		_ = V2239
		V2240 := __e.Get(2)
		_ = V2240
		V2241 := __e.Get(3)
		_ = V2241
		tmp16927 := PrimEqual(V2241, V2240)

		if True == tmp16927 {
			__e.Return(V2239)
			return
		} else {
			tmp16926 := PrimIsPair(V2241)

			if True == tmp16926 {
				tmp16925 := MakeNative(func(__e *ControlFlow) {
					W := __e.Get(1)
					_ = W
					__e.TailApply(PrimNS2Value(symsubst), V2239, V2240, W)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp16925, V2241)
				return

			} else {
				__e.Return(V2241)
				return
			}

		}

	}, 3)

	tmp16928 := Call(__e, PrimNS1Value(symns2_1set), symsubst, tmp16922)

	_ = tmp16928

	tmp16929 := MakeNative(func(__e *ControlFlow) {
		V2243 := __e.Get(1)
		_ = V2243
		tmp16930 := Call(__e, PrimNS2Value(symshen_4app), V2243, MakeString(""), symshen_4a)

		__e.TailApply(PrimNS2Value(symshen_4explode_1h), tmp16930)
		return

	}, 1)

	tmp16931 := Call(__e, PrimNS1Value(symns2_1set), symexplode, tmp16929)

	_ = tmp16931

	tmp16932 := MakeNative(func(__e *ControlFlow) {
		V2245 := __e.Get(1)
		_ = V2245
		tmp16939 := PrimEqual(MakeString(""), V2245)

		if True == tmp16939 {
			__e.Return(Nil)
			return
		} else {
			tmp16938 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2245)

			if True == tmp16938 {
				tmp16935 := PrimPos(V2245, MakeNumber(0))

				tmp16936 := PrimTailString(V2245)

				tmp16937 := Call(__e, PrimNS2Value(symshen_4explode_1h), tmp16936)

				__e.Return(PrimCons(tmp16935, tmp16937))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4explode_1h)
				return
			}

		}

	}, 1)

	tmp16940 := Call(__e, PrimNS1Value(symns2_1set), symshen_4explode_1h, tmp16932)

	_ = tmp16940

	tmp16941 := MakeNative(func(__e *ControlFlow) {
		V2247 := __e.Get(1)
		_ = V2247
		tmp16944 := PrimEqual(V2247, MakeString(""))

		var ifres16942 Obj

		if True == tmp16944 {
			ifres16942 = MakeString("")

		} else {
			tmp16943 := Call(__e, PrimNS2Value(symshen_4app), V2247, MakeString("/"), symshen_4a)

			ifres16942 = tmp16943

		}

		__e.Return(PrimNS3Set(sym_dhome_1directory_d, ifres16942))
		return

	}, 1)

	tmp16945 := Call(__e, PrimNS1Value(symns2_1set), symcd, tmp16941)

	_ = tmp16945

	tmp16946 := MakeNative(func(__e *ControlFlow) {
		V2250 := __e.Get(1)
		_ = V2250
		V2251 := __e.Get(2)
		_ = V2251
		tmp16954 := PrimEqual(Nil, V2251)

		if True == tmp16954 {
			__e.Return(True)
			return
		} else {
			tmp16953 := PrimIsPair(V2251)

			if True == tmp16953 {
				tmp16949 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					tmp16950 := PrimTail(V2251)

					__e.TailApply(PrimNS2Value(symshen_4for_1each), V2250, tmp16950)
					return

				}, 1)

				tmp16951 := PrimHead(V2251)

				tmp16952 := Call(__e, V2250, tmp16951)

				__e.TailApply(tmp16949, tmp16952)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4for_1each)
				return
			}

		}

	}, 2)

	tmp16955 := Call(__e, PrimNS1Value(symns2_1set), symshen_4for_1each, tmp16946)

	_ = tmp16955

	tmp16956 := MakeNative(func(__e *ControlFlow) {
		V2256 := __e.Get(1)
		_ = V2256
		V2257 := __e.Get(2)
		_ = V2257
		tmp16964 := PrimEqual(Nil, V2257)

		if True == tmp16964 {
			__e.Return(Nil)
			return
		} else {
			tmp16963 := PrimIsPair(V2257)

			if True == tmp16963 {
				tmp16959 := PrimHead(V2257)

				tmp16960 := Call(__e, V2256, tmp16959)

				tmp16961 := PrimTail(V2257)

				tmp16962 := Call(__e, PrimNS2Value(symmap), V2256, tmp16961)

				__e.Return(PrimCons(tmp16960, tmp16962))
				return

			} else {
				__e.TailApply(V2256, V2257)
				return
			}

		}

	}, 2)

	tmp16965 := Call(__e, PrimNS1Value(symns2_1set), symmap, tmp16956)

	_ = tmp16965

	tmp16966 := MakeNative(func(__e *ControlFlow) {
		V2259 := __e.Get(1)
		_ = V2259
		__e.TailApply(PrimNS2Value(symshen_4length_1h), V2259, MakeNumber(0))
		return
	}, 1)

	tmp16967 := Call(__e, PrimNS1Value(symns2_1set), symlength, tmp16966)

	_ = tmp16967

	tmp16968 := MakeNative(func(__e *ControlFlow) {
		V2262 := __e.Get(1)
		_ = V2262
		V2263 := __e.Get(2)
		_ = V2263
		tmp16972 := PrimEqual(Nil, V2262)

		if True == tmp16972 {
			__e.Return(V2263)
			return
		} else {
			tmp16970 := PrimTail(V2262)

			tmp16971 := PrimNumberAdd(V2263, MakeNumber(1))

			__e.TailApply(PrimNS2Value(symshen_4length_1h), tmp16970, tmp16971)
			return

		}

	}, 2)

	tmp16973 := Call(__e, PrimNS1Value(symns2_1set), symshen_4length_1h, tmp16968)

	_ = tmp16973

	tmp16974 := MakeNative(func(__e *ControlFlow) {
		V2275 := __e.Get(1)
		_ = V2275
		V2276 := __e.Get(2)
		_ = V2276
		tmp16982 := PrimEqual(V2276, V2275)

		if True == tmp16982 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp16981 := PrimIsPair(V2276)

			if True == tmp16981 {
				tmp16977 := PrimHead(V2276)

				tmp16978 := Call(__e, PrimNS2Value(symoccurrences), V2275, tmp16977)

				tmp16979 := PrimTail(V2276)

				tmp16980 := Call(__e, PrimNS2Value(symoccurrences), V2275, tmp16979)

				__e.Return(PrimNumberAdd(tmp16978, tmp16980))
				return

			} else {
				__e.Return(MakeNumber(0))
				return
			}

		}

	}, 2)

	tmp16983 := Call(__e, PrimNS1Value(symns2_1set), symoccurrences, tmp16974)

	_ = tmp16983

	tmp16984 := MakeNative(func(__e *ControlFlow) {
		V2283 := __e.Get(1)
		_ = V2283
		V2284 := __e.Get(2)
		_ = V2284
		tmp16997 := PrimEqual(MakeNumber(1), V2283)

		var ifres16994 Obj

		if True == tmp16997 {
			tmp16996 := PrimIsPair(V2284)

			var ifres16995 Obj

			if True == tmp16996 {
				ifres16995 = True

			} else {
				ifres16995 = False

			}

			ifres16994 = ifres16995

		} else {
			ifres16994 = False

		}

		if True == ifres16994 {
			__e.Return(PrimHead(V2284))
			return
		} else {
			tmp16993 := PrimIsPair(V2284)

			if True == tmp16993 {
				tmp16991 := PrimNumberSubtract(V2283, MakeNumber(1))

				tmp16992 := PrimTail(V2284)

				__e.TailApply(PrimNS2Value(symnth), tmp16991, tmp16992)
				return

			} else {
				tmp16987 := Call(__e, PrimNS2Value(symshen_4app), V2284, MakeString("\n"), symshen_4a)

				tmp16988 := PrimStringConcat(MakeString(", "), tmp16987)

				tmp16989 := Call(__e, PrimNS2Value(symshen_4app), V2283, tmp16988, symshen_4a)

				tmp16990 := PrimStringConcat(MakeString("nth applied to "), tmp16989)

				__e.Return(PrimSimpleError(tmp16990))
				return

			}

		}

	}, 2)

	tmp16998 := Call(__e, PrimNS1Value(symns2_1set), symnth, tmp16984)

	_ = tmp16998

	tmp16999 := MakeNative(func(__e *ControlFlow) {
		V2286 := __e.Get(1)
		_ = V2286
		tmp17006 := PrimIsNumber(V2286)

		if True == tmp17006 {
			tmp17002 := MakeNative(func(__e *ControlFlow) {
				Abs := __e.Get(1)
				_ = Abs
				tmp17003 := Call(__e, PrimNS2Value(symshen_4magless), Abs, MakeNumber(1))

				__e.TailApply(PrimNS2Value(symshen_4integer_1test_2), Abs, tmp17003)
				return

			}, 1)

			tmp17004 := Call(__e, PrimNS2Value(symshen_4abs), V2286)

			tmp17005 := Call(__e, tmp17002, tmp17004)

			if True == tmp17005 {
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

	tmp17007 := Call(__e, PrimNS1Value(symns2_1set), syminteger_2, tmp16999)

	_ = tmp17007

	tmp17008 := MakeNative(func(__e *ControlFlow) {
		V2288 := __e.Get(1)
		_ = V2288
		tmp17010 := PrimGreatThan(V2288, MakeNumber(0))

		if True == tmp17010 {
			__e.Return(V2288)
			return
		} else {
			__e.Return(PrimNumberSubtract(MakeNumber(0), V2288))
			return
		}

	}, 1)

	tmp17011 := Call(__e, PrimNS1Value(symns2_1set), symshen_4abs, tmp17008)

	_ = tmp17011

	tmp17012 := MakeNative(func(__e *ControlFlow) {
		V2291 := __e.Get(1)
		_ = V2291
		V2292 := __e.Get(2)
		_ = V2292
		tmp17013 := MakeNative(func(__e *ControlFlow) {
			Nx2 := __e.Get(1)
			_ = Nx2
			tmp17015 := PrimGreatThan(Nx2, V2291)

			if True == tmp17015 {
				__e.Return(V2292)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4magless), V2291, Nx2)
				return
			}

		}, 1)

		tmp17016 := PrimNumberMultiply(V2292, MakeNumber(2))

		__e.TailApply(tmp17013, tmp17016)
		return

	}, 2)

	tmp17017 := Call(__e, PrimNS1Value(symns2_1set), symshen_4magless, tmp17012)

	_ = tmp17017

	tmp17018 := MakeNative(func(__e *ControlFlow) {
		V2298 := __e.Get(1)
		_ = V2298
		V2299 := __e.Get(2)
		_ = V2299
		tmp17026 := PrimEqual(MakeNumber(0), V2298)

		if True == tmp17026 {
			__e.Return(True)
			return
		} else {
			tmp17025 := PrimGreatThan(MakeNumber(1), V2298)

			if True == tmp17025 {
				__e.Return(False)
				return
			} else {
				tmp17021 := MakeNative(func(__e *ControlFlow) {
					Abs_1N := __e.Get(1)
					_ = Abs_1N
					tmp17023 := PrimGreatThan(MakeNumber(0), Abs_1N)

					if True == tmp17023 {
						__e.Return(PrimIsInteger(V2298))
						return
					} else {
						__e.TailApply(PrimNS2Value(symshen_4integer_1test_2), Abs_1N, V2299)
						return
					}

				}, 1)

				tmp17024 := PrimNumberSubtract(V2298, V2299)

				__e.TailApply(tmp17021, tmp17024)
				return

			}

		}

	}, 2)

	tmp17027 := Call(__e, PrimNS1Value(symns2_1set), symshen_4integer_1test_2, tmp17018)

	_ = tmp17027

	tmp17028 := MakeNative(func(__e *ControlFlow) {
		V2304 := __e.Get(1)
		_ = V2304
		V2305 := __e.Get(2)
		_ = V2305
		tmp17036 := PrimEqual(Nil, V2305)

		if True == tmp17036 {
			__e.Return(Nil)
			return
		} else {
			tmp17035 := PrimIsPair(V2305)

			if True == tmp17035 {
				tmp17031 := PrimHead(V2305)

				tmp17032 := Call(__e, V2304, tmp17031)

				tmp17033 := PrimTail(V2305)

				tmp17034 := Call(__e, PrimNS2Value(symmapcan), V2304, tmp17033)

				__e.TailApply(PrimNS2Value(symappend), tmp17032, tmp17034)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symmapcan)
				return
			}

		}

	}, 2)

	tmp17037 := Call(__e, PrimNS1Value(symns2_1set), symmapcan, tmp17028)

	_ = tmp17037

	tmp17038 := MakeNative(func(__e *ControlFlow) {
		V2317 := __e.Get(1)
		_ = V2317
		V2318 := __e.Get(2)
		_ = V2318
		tmp17040 := PrimEqual(V2318, V2317)

		if True == tmp17040 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp17041 := Call(__e, PrimNS1Value(symns2_1set), sym_a_a, tmp17038)

	_ = tmp17041

	tmp17042 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("")))
		return
	}, 0)

	tmp17043 := Call(__e, PrimNS1Value(symns2_1set), symabort, tmp17042)

	_ = tmp17043

	tmp17044 := MakeNative(func(__e *ControlFlow) {
		V2320 := __e.Get(1)
		_ = V2320
		tmp17054 := PrimIsSymbol(V2320)

		if True == tmp17054 {
			tmp17047 := MakeNative(func(__e *ControlFlow) {
				Val := __e.Get(1)
				_ = Val
				tmp17049 := PrimEqual(Val, symshen_4this_1symbol_1is_1unbound)

				if True == tmp17049 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}, 1)

			tmp17050 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimNS3Value(V2320))
				return
			}, 0)

			tmp17051 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4this_1symbol_1is_1unbound)
				return
			}, 1)

			tmp17052 := Call(__e, PrimNS1Value(symtry_1catch), tmp17050, tmp17051)

			tmp17053 := Call(__e, tmp17047, tmp17052)

			if True == tmp17053 {
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

	tmp17055 := Call(__e, PrimNS1Value(symns2_1set), symbound_2, tmp17044)

	_ = tmp17055

	tmp17056 := MakeNative(func(__e *ControlFlow) {
		V2322 := __e.Get(1)
		_ = V2322
		tmp17062 := PrimEqual(MakeString(""), V2322)

		if True == tmp17062 {
			__e.Return(Nil)
			return
		} else {
			tmp17058 := PrimPos(V2322, MakeNumber(0))

			tmp17059 := PrimStringToNumber(tmp17058)

			tmp17060 := PrimTailString(V2322)

			tmp17061 := Call(__e, PrimNS2Value(symshen_4string_1_6bytes), tmp17060)

			__e.Return(PrimCons(tmp17059, tmp17061))
			return

		}

	}, 1)

	tmp17063 := Call(__e, PrimNS1Value(symns2_1set), symshen_4string_1_6bytes, tmp17056)

	_ = tmp17063

	tmp17064 := MakeNative(func(__e *ControlFlow) {
		V2324 := __e.Get(1)
		_ = V2324
		__e.Return(PrimNS3Set(symshen_4_dmaxinferences_d, V2324))
		return
	}, 1)

	tmp17065 := Call(__e, PrimNS1Value(symns2_1set), symmaxinferences, tmp17064)

	_ = tmp17065

	tmp17066 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dinfs_d))
		return
	}, 0)

	tmp17067 := Call(__e, PrimNS1Value(symns2_1set), syminferences, tmp17066)

	_ = tmp17067

	tmp17068 := MakeNative(func(__e *ControlFlow) {
		V2326 := __e.Get(1)
		_ = V2326
		__e.Return(V2326)
		return
	}, 1)

	tmp17069 := Call(__e, PrimNS1Value(symns2_1set), symprotect, tmp17068)

	_ = tmp17069

	tmp17070 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dstoutput_d))
		return
	}, 0)

	tmp17071 := Call(__e, PrimNS1Value(symns2_1set), symstoutput, tmp17070)

	_ = tmp17071

	tmp17072 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dsterror_d))
		return
	}, 0)

	tmp17073 := Call(__e, PrimNS1Value(symns2_1set), symsterror, tmp17072)

	_ = tmp17073

	tmp17074 := MakeNative(func(__e *ControlFlow) {
		V2328 := __e.Get(1)
		_ = V2328
		tmp17075 := MakeNative(func(__e *ControlFlow) {
			Symbol := __e.Get(1)
			_ = Symbol
			tmp17079 := PrimIsSymbol(Symbol)

			if True == tmp17079 {
				__e.Return(Symbol)
				return
			} else {
				tmp17077 := Call(__e, PrimNS2Value(symshen_4app), V2328, MakeString(" to a symbol"), symshen_4s)

				tmp17078 := PrimStringConcat(MakeString("cannot intern "), tmp17077)

				__e.Return(PrimSimpleError(tmp17078))
				return

			}

		}, 1)

		tmp17080 := PrimIntern(V2328)

		__e.TailApply(tmp17075, tmp17080)
		return

	}, 1)

	tmp17081 := Call(__e, PrimNS1Value(symns2_1set), symstring_1_6symbol, tmp17074)

	_ = tmp17081

	tmp17082 := MakeNative(func(__e *ControlFlow) {
		V2334 := __e.Get(1)
		_ = V2334
		tmp17086 := PrimEqual(sym_7, V2334)

		if True == tmp17086 {
			__e.Return(PrimNS3Set(symshen_4_doptimise_d, True))
			return
		} else {
			tmp17085 := PrimEqual(sym_1, V2334)

			if True == tmp17085 {
				__e.Return(PrimNS3Set(symshen_4_doptimise_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp17087 := Call(__e, PrimNS1Value(symns2_1set), symoptimise, tmp17082)

	_ = tmp17087

	tmp17088 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dos_d))
		return
	}, 0)

	tmp17089 := Call(__e, PrimNS1Value(symns2_1set), symos, tmp17088)

	_ = tmp17089

	tmp17090 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dlanguage_d))
		return
	}, 0)

	tmp17091 := Call(__e, PrimNS1Value(symns2_1set), symlanguage, tmp17090)

	_ = tmp17091

	tmp17092 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dversion_d))
		return
	}, 0)

	tmp17093 := Call(__e, PrimNS1Value(symns2_1set), symversion, tmp17092)

	_ = tmp17093

	tmp17094 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dport_d))
		return
	}, 0)

	tmp17095 := Call(__e, PrimNS1Value(symns2_1set), symport, tmp17094)

	_ = tmp17095

	tmp17096 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dporters_d))
		return
	}, 0)

	tmp17097 := Call(__e, PrimNS1Value(symns2_1set), symporters, tmp17096)

	_ = tmp17097

	tmp17098 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_dimplementation_d))
		return
	}, 0)

	tmp17099 := Call(__e, PrimNS1Value(symns2_1set), symimplementation, tmp17098)

	_ = tmp17099

	tmp17100 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(sym_drelease_d))
		return
	}, 0)

	tmp17101 := Call(__e, PrimNS1Value(symns2_1set), symrelease, tmp17100)

	_ = tmp17101

	tmp17102 := MakeNative(func(__e *ControlFlow) {
		V2336 := __e.Get(1)
		_ = V2336
		tmp17103 := MakeNative(func(__e *ControlFlow) {
			tmp17104 := Call(__e, PrimNS2Value(symexternal), V2336)

			_ = tmp17104

			__e.Return(True)
			return

		}, 0)

		tmp17105 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp17103, tmp17105)
		return

	}, 1)

	tmp17106 := Call(__e, PrimNS1Value(symns2_1set), sympackage_2, tmp17102)

	_ = tmp17106

	tmp17107 := MakeNative(func(__e *ControlFlow) {
		V2338 := __e.Get(1)
		_ = V2338
		__e.TailApply(PrimNS2Value(symshen_4lookup_1func), V2338)
		return
	}, 1)

	tmp17108 := Call(__e, PrimNS1Value(symns2_1set), symfunction, tmp17107)

	_ = tmp17108

	tmp17109 := MakeNative(func(__e *ControlFlow) {
		V2340 := __e.Get(1)
		_ = V2340
		tmp17110 := MakeNative(func(__e *ControlFlow) {
			tmp17111 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V2340, symshen_4lambda_1form, tmp17111)
			return

		}, 0)

		tmp17112 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp17113 := Call(__e, PrimNS2Value(symshen_4app), V2340, MakeString(" has no lambda expansion\n"), symshen_4a)

			__e.Return(PrimSimpleError(tmp17113))
			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp17110, tmp17112)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4lookup_1func, tmp17109)
	return

}, 0)
