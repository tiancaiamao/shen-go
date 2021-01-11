package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen9045 := MakeNative(func(__e *ControlFlow) {
		V4724 := __e.Get(1)
		_ = V4724
		gen9042 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9043 := Call(__e, gen9042, V4724)

		var gen9044 Obj

		if True == gen9043 {
			gen9037 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9038 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9039 := Call(__e, gen9038, V4724)

			gen9040 := Call(__e, gen9037, symdefcc, gen9039)

			var gen9041 Obj

			if True == gen9040 {
				gen9033 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9034 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9035 := Call(__e, gen9034, V4724)

				gen9036 := Call(__e, gen9033, gen9035)

				if True == gen9036 {
					gen9041 = True
				} else {
					gen9041 = False
				}

			} else {
				gen9041 = False
			}

			if True == gen9041 {
				gen9044 = True
			} else {
				gen9044 = False
			}

		} else {
			gen9044 = False
		}

		if True == gen9044 {
			gen9024 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc_1_6shen)

			gen9025 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9026 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9027 := Call(__e, gen9026, V4724)

			gen9028 := Call(__e, gen9025, gen9027)

			gen9029 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9031 := Call(__e, gen9030, V4724)

			gen9032 := Call(__e, gen9029, gen9031)

			__e.TailApply(gen9024, gen9028, gen9032)

			return

		} else {
			if True == True {
				gen9023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9023, symshen_4yacc)

				return

			} else {
				gen9022 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9022, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc, gen9045)

	gen9068 := MakeNative(func(__e *ControlFlow) {
		V4727 := __e.Get(1)
		_ = V4727
		V4728 := __e.Get(2)
		_ = V4728
		gen9065 := MakeNative(func(__e *ControlFlow) {
			CCRules := __e.Get(1)
			_ = CCRules
			gen9060 := MakeNative(func(__e *ControlFlow) {
				CCBody := __e.Get(1)
				_ = CCBody
				gen9057 := MakeNative(func(__e *ControlFlow) {
					YaccCases := __e.Get(1)
					_ = YaccCases
					gen9046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9047 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9048 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9049 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9050 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9051 := Call(__e, PrimNS1Value(symns2_1value), symshen_4kill_1code)

					gen9052 := Call(__e, gen9051, YaccCases)

					gen9053 := Call(__e, gen9050, gen9052, Nil)

					gen9054 := Call(__e, gen9049, sym_1_6, gen9053)

					gen9055 := Call(__e, gen9048, symStream, gen9054)

					gen9056 := Call(__e, gen9047, V4727, gen9055)

					__e.TailApply(gen9046, symdefine, gen9056)

					return

				}, 1)

				gen9058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc__cases)

				gen9059 := Call(__e, gen9058, CCBody)

				__e.TailApply(gen9057, gen9059)

				return

			}, 1)

			gen9061 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			gen9063 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen9062 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cc__body)

				__e.TailApply(gen9062, X)

				return

			}, 1)

			gen9064 := Call(__e, gen9061, gen9063, CCRules)

			__e.TailApply(gen9060, gen9064)

			return

		}, 1)

		gen9066 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

		gen9067 := Call(__e, gen9066, True, V4728, Nil)

		__e.TailApply(gen9065, gen9067)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc_1_6shen, gen9068)

	gen9089 := MakeNative(func(__e *ControlFlow) {
		V4730 := __e.Get(1)
		_ = V4730
		gen9085 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		gen9086 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

		gen9087 := Call(__e, gen9086, symkill, V4730)

		gen9088 := Call(__e, gen9085, gen9087, MakeNumber(0))

		if True == gen9088 {
			gen9070 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9071 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9072 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9073 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9074 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9075 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9076 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9077 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9078 := Call(__e, gen9077, symE, Nil)

			gen9079 := Call(__e, gen9076, symshen_4analyse_1kill, gen9078)

			gen9080 := Call(__e, gen9075, gen9079, Nil)

			gen9081 := Call(__e, gen9074, symE, gen9080)

			gen9082 := Call(__e, gen9073, symlambda, gen9081)

			gen9083 := Call(__e, gen9072, gen9082, Nil)

			gen9084 := Call(__e, gen9071, V4730, gen9083)

			__e.TailApply(gen9070, symtrap_1error, gen9084)

			return

		} else {
			if True == True {
				__e.Return(V4730)
				return
			} else {
				gen9069 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9069, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4kill_1code, gen9089)

	gen9091 := MakeNative(func(__e *ControlFlow) {
		gen9090 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(gen9090, MakeString("yacc kill"))

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symkill, gen9091)

	gen9098 := MakeNative(func(__e *ControlFlow) {
		V4732 := __e.Get(1)
		_ = V4732
		gen9095 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			gen9093 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9094 := Call(__e, gen9093, String, MakeString("yacc kill"))

			if True == gen9094 {
				gen9092 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen9092)

				return

			} else {
				__e.Return(V4732)
				return
			}

		}, 1)

		gen9096 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

		gen9097 := Call(__e, gen9096, V4732)

		__e.TailApply(gen9095, gen9097)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1kill, gen9098)

	gen9138 := MakeNative(func(__e *ControlFlow) {
		V4738 := __e.Get(1)
		_ = V4738
		V4739 := __e.Get(2)
		_ = V4739
		V4740 := __e.Get(3)
		_ = V4740
		gen9135 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen9136 := Call(__e, gen9135, Nil, V4739)

		var gen9137 Obj

		if True == gen9136 {
			gen9133 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9134 := Call(__e, gen9133, Nil, V4740)

			if True == gen9134 {
				gen9137 = True
			} else {
				gen9137 = False
			}

		} else {
			gen9137 = False
		}

		if True == gen9137 {
			__e.Return(Nil)
			return
		} else {
			gen9131 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9132 := Call(__e, gen9131, Nil, V4739)

			if True == gen9132 {
				gen9126 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

				gen9128 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				gen9129 := Call(__e, gen9128, V4740)

				gen9130 := Call(__e, gen9127, V4738, gen9129, Nil)

				__e.TailApply(gen9126, gen9130, Nil)

				return

			} else {
				gen9123 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9124 := Call(__e, gen9123, V4739)

				var gen9125 Obj

				if True == gen9124 {
					gen9119 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen9120 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9121 := Call(__e, gen9120, V4739)

					gen9122 := Call(__e, gen9119, sym_k, gen9121)

					if True == gen9122 {
						gen9125 = True
					} else {
						gen9125 = False
					}

				} else {
					gen9125 = False
				}

				if True == gen9125 {
					gen9110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9111 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

					gen9112 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					gen9113 := Call(__e, gen9112, V4740)

					gen9114 := Call(__e, gen9111, V4738, gen9113, Nil)

					gen9115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

					gen9116 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9117 := Call(__e, gen9116, V4739)

					gen9118 := Call(__e, gen9115, V4738, gen9117, Nil)

					__e.TailApply(gen9110, gen9114, gen9118)

					return

				} else {
					gen9108 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9109 := Call(__e, gen9108, V4739)

					if True == gen9109 {
						gen9101 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

						gen9102 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9103 := Call(__e, gen9102, V4739)

						gen9104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9105 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9106 := Call(__e, gen9105, V4739)

						gen9107 := Call(__e, gen9104, gen9106, V4740)

						__e.TailApply(gen9101, V4738, gen9103, gen9107)

						return

					} else {
						if True == True {
							gen9100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen9100, symshen_4split__cc__rules)

							return

						} else {
							gen9099 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen9099, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rules, gen9138)

	gen9257 := MakeNative(func(__e *ControlFlow) {
		V4748 := __e.Get(1)
		_ = V4748
		V4749 := __e.Get(2)
		_ = V4749
		V4750 := __e.Get(3)
		_ = V4750
		gen9254 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9255 := Call(__e, gen9254, V4749)

		var gen9256 Obj

		if True == gen9255 {
			gen9249 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9250 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9251 := Call(__e, gen9250, V4749)

			gen9252 := Call(__e, gen9249, sym_h_a, gen9251)

			var gen9253 Obj

			if True == gen9252 {
				gen9244 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9245 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9246 := Call(__e, gen9245, V4749)

				gen9247 := Call(__e, gen9244, gen9246)

				var gen9248 Obj

				if True == gen9247 {
					gen9238 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen9239 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9241 := Call(__e, gen9240, V4749)

					gen9242 := Call(__e, gen9239, gen9241)

					gen9243 := Call(__e, gen9238, Nil, gen9242)

					if True == gen9243 {
						gen9248 = True
					} else {
						gen9248 = False
					}

				} else {
					gen9248 = False
				}

				if True == gen9248 {
					gen9253 = True
				} else {
					gen9253 = False
				}

			} else {
				gen9253 = False
			}

			if True == gen9253 {
				gen9256 = True
			} else {
				gen9256 = False
			}

		} else {
			gen9256 = False
		}

		if True == gen9256 {
			gen9233 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9234 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			gen9235 := Call(__e, gen9234, V4750)

			gen9236 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9237 := Call(__e, gen9236, V4749)

			__e.TailApply(gen9233, gen9235, gen9237)

			return

		} else {
			gen9230 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9231 := Call(__e, gen9230, V4749)

			var gen9232 Obj

			if True == gen9231 {
				gen9225 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9226 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9227 := Call(__e, gen9226, V4749)

				gen9228 := Call(__e, gen9225, sym_h_a, gen9227)

				var gen9229 Obj

				if True == gen9228 {
					gen9220 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9221 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9222 := Call(__e, gen9221, V4749)

					gen9223 := Call(__e, gen9220, gen9222)

					var gen9224 Obj

					if True == gen9223 {
						gen9213 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen9214 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9215 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9216 := Call(__e, gen9215, V4749)

						gen9217 := Call(__e, gen9214, gen9216)

						gen9218 := Call(__e, gen9213, gen9217)

						var gen9219 Obj

						if True == gen9218 {
							gen9204 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen9205 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen9206 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9208 := Call(__e, gen9207, V4749)

							gen9209 := Call(__e, gen9206, gen9208)

							gen9210 := Call(__e, gen9205, gen9209)

							gen9211 := Call(__e, gen9204, symwhere, gen9210)

							var gen9212 Obj

							if True == gen9211 {
								gen9195 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen9196 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen9197 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen9198 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen9199 := Call(__e, gen9198, V4749)

								gen9200 := Call(__e, gen9197, gen9199)

								gen9201 := Call(__e, gen9196, gen9200)

								gen9202 := Call(__e, gen9195, gen9201)

								var gen9203 Obj

								if True == gen9202 {
									gen9185 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen9186 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen9187 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen9188 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen9189 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen9190 := Call(__e, gen9189, V4749)

									gen9191 := Call(__e, gen9188, gen9190)

									gen9192 := Call(__e, gen9187, gen9191)

									gen9193 := Call(__e, gen9186, gen9192)

									gen9194 := Call(__e, gen9185, Nil, gen9193)

									if True == gen9194 {
										gen9203 = True
									} else {
										gen9203 = False
									}

								} else {
									gen9203 = False
								}

								if True == gen9203 {
									gen9212 = True
								} else {
									gen9212 = False
								}

							} else {
								gen9212 = False
							}

							if True == gen9212 {
								gen9219 = True
							} else {
								gen9219 = False
							}

						} else {
							gen9219 = False
						}

						if True == gen9219 {
							gen9224 = True
						} else {
							gen9224 = False
						}

					} else {
						gen9224 = False
					}

					if True == gen9224 {
						gen9229 = True
					} else {
						gen9229 = False
					}

				} else {
					gen9229 = False
				}

				if True == gen9229 {
					gen9232 = True
				} else {
					gen9232 = False
				}

			} else {
				gen9232 = False
			}

			if True == gen9232 {
				gen9162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9163 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				gen9164 := Call(__e, gen9163, V4750)

				gen9165 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9167 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9168 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9169 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9170 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9171 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9172 := Call(__e, gen9171, V4749)

				gen9173 := Call(__e, gen9170, gen9172)

				gen9174 := Call(__e, gen9169, gen9173)

				gen9175 := Call(__e, gen9168, gen9174)

				gen9176 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9177 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9178 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9179 := Call(__e, gen9178, V4749)

				gen9180 := Call(__e, gen9177, gen9179)

				gen9181 := Call(__e, gen9176, gen9180, Nil)

				gen9182 := Call(__e, gen9167, gen9175, gen9181)

				gen9183 := Call(__e, gen9166, symwhere, gen9182)

				gen9184 := Call(__e, gen9165, gen9183, Nil)

				__e.TailApply(gen9162, gen9164, gen9184)

				return

			} else {
				gen9160 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9161 := Call(__e, gen9160, Nil, V4749)

				if True == gen9161 {
					gen9150 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantic_1completion_1warning)

					Call(__e, gen9150, V4748, V4750)

					gen9151 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

					gen9152 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9153 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

					gen9155 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					gen9156 := Call(__e, gen9155, V4750)

					gen9157 := Call(__e, gen9154, gen9156)

					gen9158 := Call(__e, gen9153, gen9157, Nil)

					gen9159 := Call(__e, gen9152, sym_h_a, gen9158)

					__e.TailApply(gen9151, V4748, gen9159, V4750)

					return

				} else {
					gen9148 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9149 := Call(__e, gen9148, V4749)

					if True == gen9149 {
						gen9141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

						gen9142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9143 := Call(__e, gen9142, V4749)

						gen9144 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9145 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9146 := Call(__e, gen9145, V4749)

						gen9147 := Call(__e, gen9144, gen9146, V4750)

						__e.TailApply(gen9141, V4748, gen9143, gen9147)

						return

					} else {
						if True == True {
							gen9140 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen9140, symshen_4split__cc__rule)

							return

						} else {
							gen9139 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen9139, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rule, gen9257)

	gen9276 := MakeNative(func(__e *ControlFlow) {
		V4761 := __e.Get(1)
		_ = V4761
		V4762 := __e.Get(2)
		_ = V4762
		gen9274 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen9275 := Call(__e, gen9274, True, V4761)

		if True == gen9275 {
			gen9259 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen9260 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen9261 := Call(__e, gen9260)

			Call(__e, gen9259, MakeString("warning: "), gen9261)

			gen9262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

			gen9268 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen9263 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				gen9264 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen9265 := Call(__e, gen9264, X, MakeString(" "), symshen_4a)

				gen9266 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				gen9267 := Call(__e, gen9266)

				__e.TailApply(gen9263, gen9265, gen9267)

				return

			}, 1)

			gen9269 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			gen9270 := Call(__e, gen9269, V4762)

			Call(__e, gen9262, gen9268, gen9270)

			gen9271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen9272 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen9273 := Call(__e, gen9272)

			__e.TailApply(gen9271, MakeString("has no semantics.\n"), gen9273)

			return

		} else {
			if True == True {
				__e.Return(symshen_4skip)
				return
			} else {
				gen9258 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9258, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4semantic_1completion_1warning, gen9276)

	gen9325 := MakeNative(func(__e *ControlFlow) {
		V4764 := __e.Get(1)
		_ = V4764
		gen9323 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen9324 := Call(__e, gen9323, Nil, V4764)

		if True == gen9324 {
			__e.Return(Nil)
			return
		} else {
			gen9320 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9321 := Call(__e, gen9320, V4764)

			var gen9322 Obj

			if True == gen9321 {
				gen9315 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9316 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9317 := Call(__e, gen9316, V4764)

				gen9318 := Call(__e, gen9315, Nil, gen9317)

				var gen9319 Obj

				if True == gen9318 {
					gen9311 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					gen9312 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9313 := Call(__e, gen9312, V4764)

					gen9314 := Call(__e, gen9311, gen9313)

					if True == gen9314 {
						gen9319 = True
					} else {
						gen9319 = False
					}

				} else {
					gen9319 = False
				}

				if True == gen9319 {
					gen9322 = True
				} else {
					gen9322 = False
				}

			} else {
				gen9322 = False
			}

			if True == gen9322 {
				gen9310 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				__e.TailApply(gen9310, V4764)

				return

			} else {
				gen9307 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9308 := Call(__e, gen9307, V4764)

				var gen9309 Obj

				if True == gen9308 {
					gen9303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					gen9304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9305 := Call(__e, gen9304, V4764)

					gen9306 := Call(__e, gen9303, gen9305)

					if True == gen9306 {
						gen9309 = True
					} else {
						gen9309 = False
					}

				} else {
					gen9309 = False
				}

				if True == gen9309 {
					gen9292 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9293 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9294 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9295 := Call(__e, gen9294, V4764)

					gen9296 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9297 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

					gen9298 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9299 := Call(__e, gen9298, V4764)

					gen9300 := Call(__e, gen9297, gen9299)

					gen9301 := Call(__e, gen9296, gen9300, Nil)

					gen9302 := Call(__e, gen9293, gen9295, gen9301)

					__e.TailApply(gen9292, symappend, gen9302)

					return

				} else {
					gen9290 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9291 := Call(__e, gen9290, V4764)

					if True == gen9291 {
						gen9279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9281 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9282 := Call(__e, gen9281, V4764)

						gen9283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

						gen9285 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9286 := Call(__e, gen9285, V4764)

						gen9287 := Call(__e, gen9284, gen9286)

						gen9288 := Call(__e, gen9283, gen9287, Nil)

						gen9289 := Call(__e, gen9280, gen9282, gen9288)

						__e.TailApply(gen9279, symcons, gen9289)

						return

					} else {
						if True == True {
							gen9278 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen9278, symshen_4default__semantics)

							return

						} else {
							gen9277 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen9277, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4default__semantics, gen9325)

	gen9344 := MakeNative(func(__e *ControlFlow) {
		V4766 := __e.Get(1)
		_ = V4766
		gen9342 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		gen9343 := Call(__e, gen9342, V4766)

		if True == gen9343 {
			gen9336 := MakeNative(func(__e *ControlFlow) {
				Cs := __e.Get(1)
				_ = Cs
				gen9332 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9333 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9334 := Call(__e, gen9333, Cs)

				gen9335 := Call(__e, gen9332, gen9334, MakeString("<"))

				if True == gen9335 {
					gen9326 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen9327 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9328 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					gen9329 := Call(__e, gen9328, Cs)

					gen9330 := Call(__e, gen9327, gen9329)

					gen9331 := Call(__e, gen9326, gen9330, MakeString(">"))

					if True == gen9331 {
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

			gen9337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1pathname)

			gen9338 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			gen9339 := Call(__e, gen9338, V4766)

			gen9340 := Call(__e, gen9337, gen9339)

			gen9341 := Call(__e, gen9336, gen9340)

			if True == gen9341 {
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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4grammar__symbol_2, gen9344)

	gen9387 := MakeNative(func(__e *ControlFlow) {
		V4768 := __e.Get(1)
		_ = V4768
		gen9384 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9385 := Call(__e, gen9384, V4768)

		var gen9386 Obj

		if True == gen9385 {
			gen9380 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9381 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9382 := Call(__e, gen9381, V4768)

			gen9383 := Call(__e, gen9380, Nil, gen9382)

			if True == gen9383 {
				gen9386 = True
			} else {
				gen9386 = False
			}

		} else {
			gen9386 = False
		}

		if True == gen9386 {
			gen9379 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(gen9379, V4768)

			return

		} else {
			gen9377 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9378 := Call(__e, gen9377, V4768)

			if True == gen9378 {
				gen9376 := MakeNative(func(__e *ControlFlow) {
					P := __e.Get(1)
					_ = P
					gen9347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9350 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9351 := Call(__e, gen9350, V4768)

					gen9352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9357 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9359 := Call(__e, gen9358, symfail, Nil)

					gen9360 := Call(__e, gen9357, gen9359, Nil)

					gen9361 := Call(__e, gen9356, P, gen9360)

					gen9362 := Call(__e, gen9355, sym_a, gen9361)

					gen9363 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc__cases)

					gen9365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9366 := Call(__e, gen9365, V4768)

					gen9367 := Call(__e, gen9364, gen9366)

					gen9368 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9369 := Call(__e, gen9368, P, Nil)

					gen9370 := Call(__e, gen9363, gen9367, gen9369)

					gen9371 := Call(__e, gen9354, gen9362, gen9370)

					gen9372 := Call(__e, gen9353, symif, gen9371)

					gen9373 := Call(__e, gen9352, gen9372, Nil)

					gen9374 := Call(__e, gen9349, gen9351, gen9373)

					gen9375 := Call(__e, gen9348, P, gen9374)

					__e.TailApply(gen9347, symlet, gen9375)

					return

				}, 1)

				__e.TailApply(gen9376, symYaccParse)

				return

			} else {
				if True == True {
					gen9346 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen9346, symshen_4yacc__cases)

					return

				} else {
					gen9345 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen9345, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc__cases, gen9387)

	gen9411 := MakeNative(func(__e *ControlFlow) {
		V4770 := __e.Get(1)
		_ = V4770
		gen9408 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9409 := Call(__e, gen9408, V4770)

		var gen9410 Obj

		if True == gen9409 {
			gen9403 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9404 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9405 := Call(__e, gen9404, V4770)

			gen9406 := Call(__e, gen9403, gen9405)

			var gen9407 Obj

			if True == gen9406 {
				gen9397 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9398 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9399 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9400 := Call(__e, gen9399, V4770)

				gen9401 := Call(__e, gen9398, gen9400)

				gen9402 := Call(__e, gen9397, Nil, gen9401)

				if True == gen9402 {
					gen9407 = True
				} else {
					gen9407 = False
				}

			} else {
				gen9407 = False
			}

			if True == gen9407 {
				gen9410 = True
			} else {
				gen9410 = False
			}

		} else {
			gen9410 = False
		}

		if True == gen9410 {
			gen9390 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

			gen9391 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9392 := Call(__e, gen9391, V4770)

			gen9393 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9394 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9395 := Call(__e, gen9394, V4770)

			gen9396 := Call(__e, gen9393, gen9395)

			__e.TailApply(gen9390, gen9392, symStream, gen9396)

			return

		} else {
			if True == True {
				gen9389 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9389, symshen_4cc__body)

				return

			} else {
				gen9388 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9388, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4cc__body, gen9411)

	gen9529 := MakeNative(func(__e *ControlFlow) {
		V4774 := __e.Get(1)
		_ = V4774
		V4775 := __e.Get(2)
		_ = V4775
		V4776 := __e.Get(3)
		_ = V4776
		gen9526 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen9527 := Call(__e, gen9526, Nil, V4774)

		var gen9528 Obj

		if True == gen9527 {
			gen9523 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9524 := Call(__e, gen9523, V4776)

			var gen9525 Obj

			if True == gen9524 {
				gen9518 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9519 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9520 := Call(__e, gen9519, V4776)

				gen9521 := Call(__e, gen9518, symwhere, gen9520)

				var gen9522 Obj

				if True == gen9521 {
					gen9513 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9514 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9515 := Call(__e, gen9514, V4776)

					gen9516 := Call(__e, gen9513, gen9515)

					var gen9517 Obj

					if True == gen9516 {
						gen9506 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen9507 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9508 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9509 := Call(__e, gen9508, V4776)

						gen9510 := Call(__e, gen9507, gen9509)

						gen9511 := Call(__e, gen9506, gen9510)

						var gen9512 Obj

						if True == gen9511 {
							gen9498 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen9499 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9500 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9501 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9502 := Call(__e, gen9501, V4776)

							gen9503 := Call(__e, gen9500, gen9502)

							gen9504 := Call(__e, gen9499, gen9503)

							gen9505 := Call(__e, gen9498, Nil, gen9504)

							if True == gen9505 {
								gen9512 = True
							} else {
								gen9512 = False
							}

						} else {
							gen9512 = False
						}

						if True == gen9512 {
							gen9517 = True
						} else {
							gen9517 = False
						}

					} else {
						gen9517 = False
					}

					if True == gen9517 {
						gen9522 = True
					} else {
						gen9522 = False
					}

				} else {
					gen9522 = False
				}

				if True == gen9522 {
					gen9525 = True
				} else {
					gen9525 = False
				}

			} else {
				gen9525 = False
			}

			if True == gen9525 {
				gen9528 = True
			} else {
				gen9528 = False
			}

		} else {
			gen9528 = False
		}

		if True == gen9528 {
			gen9465 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9466 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9467 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

			gen9468 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9469 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9470 := Call(__e, gen9469, V4776)

			gen9471 := Call(__e, gen9468, gen9470)

			gen9472 := Call(__e, gen9467, gen9471)

			gen9473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9475 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9476 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9477 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9478 := Call(__e, gen9477, V4775, Nil)

			gen9479 := Call(__e, gen9476, symhd, gen9478)

			gen9480 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9481 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

			gen9482 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9483 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9484 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9485 := Call(__e, gen9484, V4776)

			gen9486 := Call(__e, gen9483, gen9485)

			gen9487 := Call(__e, gen9482, gen9486)

			gen9488 := Call(__e, gen9481, gen9487)

			gen9489 := Call(__e, gen9480, gen9488, Nil)

			gen9490 := Call(__e, gen9475, gen9479, gen9489)

			gen9491 := Call(__e, gen9474, symshen_4pair, gen9490)

			gen9492 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9493 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9494 := Call(__e, gen9493, symfail, Nil)

			gen9495 := Call(__e, gen9492, gen9494, Nil)

			gen9496 := Call(__e, gen9473, gen9491, gen9495)

			gen9497 := Call(__e, gen9466, gen9472, gen9496)

			__e.TailApply(gen9465, symif, gen9497)

			return

		} else {
			gen9463 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9464 := Call(__e, gen9463, Nil, V4774)

			if True == gen9464 {
				gen9452 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9453 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9455 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9456 := Call(__e, gen9455, V4775, Nil)

				gen9457 := Call(__e, gen9454, symhd, gen9456)

				gen9458 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9459 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

				gen9460 := Call(__e, gen9459, V4776)

				gen9461 := Call(__e, gen9458, gen9460, Nil)

				gen9462 := Call(__e, gen9453, gen9457, gen9461)

				__e.TailApply(gen9452, symshen_4pair, gen9462)

				return

			} else {
				gen9450 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9451 := Call(__e, gen9450, V4774)

				if True == gen9451 {
					gen9446 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					gen9447 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen9448 := Call(__e, gen9447, V4774)

					gen9449 := Call(__e, gen9446, gen9448)

					if True == gen9449 {
						gen9445 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__descent)

						__e.TailApply(gen9445, V4774, V4775, V4776)

						return

					} else {
						gen9441 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

						gen9442 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9443 := Call(__e, gen9442, V4774)

						gen9444 := Call(__e, gen9441, gen9443)

						if True == gen9444 {
							gen9440 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variable_1match)

							__e.TailApply(gen9440, V4774, V4775, V4776)

							return

						} else {
							gen9436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4jump__stream_2)

							gen9437 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen9438 := Call(__e, gen9437, V4774)

							gen9439 := Call(__e, gen9436, gen9438)

							if True == gen9439 {
								gen9435 := Call(__e, PrimNS1Value(symns2_1value), symshen_4jump__stream)

								__e.TailApply(gen9435, V4774, V4775, V4776)

								return

							} else {
								gen9431 := Call(__e, PrimNS1Value(symns2_1value), symshen_4terminal_2)

								gen9432 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen9433 := Call(__e, gen9432, V4774)

								gen9434 := Call(__e, gen9431, gen9433)

								if True == gen9434 {
									gen9430 := Call(__e, PrimNS1Value(symns2_1value), symshen_4check__stream)

									__e.TailApply(gen9430, V4774, V4775, V4776)

									return

								} else {
									gen9426 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									gen9427 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen9428 := Call(__e, gen9427, V4774)

									gen9429 := Call(__e, gen9426, gen9428)

									if True == gen9429 {
										gen9419 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_1stream)

										gen9420 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decons)

										gen9421 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen9422 := Call(__e, gen9421, V4774)

										gen9423 := Call(__e, gen9420, gen9422)

										gen9424 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										gen9425 := Call(__e, gen9424, V4774)

										__e.TailApply(gen9419, gen9423, gen9425, V4775, V4776)

										return

									} else {
										gen9414 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

										gen9415 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

										gen9416 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen9417 := Call(__e, gen9416, V4774)

										gen9418 := Call(__e, gen9415, gen9417, MakeString(" is not legal syntax\n"), symshen_4a)

										__e.TailApply(gen9414, gen9418)

										return

									}

								}

							}

						}

					}

				} else {
					if True == True {
						gen9413 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen9413, symshen_4syntax)

						return

					} else {
						gen9412 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen9412, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4syntax, gen9529)

	gen9601 := MakeNative(func(__e *ControlFlow) {
		V4781 := __e.Get(1)
		_ = V4781
		V4782 := __e.Get(2)
		_ = V4782
		V4783 := __e.Get(3)
		_ = V4783
		V4784 := __e.Get(4)
		_ = V4784
		gen9578 := MakeNative(func(__e *ControlFlow) {
			Test := __e.Get(1)
			_ = Test
			gen9575 := MakeNative(func(__e *ControlFlow) {
				Placeholder := __e.Get(1)
				_ = Placeholder
				gen9558 := MakeNative(func(__e *ControlFlow) {
					RunOn := __e.Get(1)
					_ = RunOn
					gen9539 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						gen9530 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9531 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9532 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9533 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9535 := Call(__e, gen9534, symfail, Nil)

						gen9536 := Call(__e, gen9533, gen9535, Nil)

						gen9537 := Call(__e, gen9532, Action, gen9536)

						gen9538 := Call(__e, gen9531, Test, gen9537)

						__e.TailApply(gen9530, symif, gen9538)

						return

					}, 1)

					gen9540 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1runon)

					gen9541 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

					gen9542 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9543 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9545 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9546 := Call(__e, gen9545, V4783, Nil)

					gen9547 := Call(__e, gen9544, symshen_4hdhd, gen9546)

					gen9548 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9549 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9550 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9551 := Call(__e, gen9550, V4783, Nil)

					gen9552 := Call(__e, gen9549, symshen_4hdtl, gen9551)

					gen9553 := Call(__e, gen9548, gen9552, Nil)

					gen9554 := Call(__e, gen9543, gen9547, gen9553)

					gen9555 := Call(__e, gen9542, symshen_4pair, gen9554)

					gen9556 := Call(__e, gen9541, V4781, gen9555, Placeholder)

					gen9557 := Call(__e, gen9540, RunOn, Placeholder, gen9556)

					__e.TailApply(gen9539, gen9557)

					return

				}, 1)

				gen9559 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				gen9560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9561 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9562 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9563 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9564 := Call(__e, gen9563, V4783, Nil)

				gen9565 := Call(__e, gen9562, symshen_4tlhd, gen9564)

				gen9566 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9567 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9569 := Call(__e, gen9568, V4783, Nil)

				gen9570 := Call(__e, gen9567, symshen_4hdtl, gen9569)

				gen9571 := Call(__e, gen9566, gen9570, Nil)

				gen9572 := Call(__e, gen9561, gen9565, gen9571)

				gen9573 := Call(__e, gen9560, symshen_4pair, gen9572)

				gen9574 := Call(__e, gen9559, V4782, gen9573, V4784)

				__e.TailApply(gen9558, gen9574)

				return

			}, 1)

			gen9576 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			gen9577 := Call(__e, gen9576, symshen_4place)

			__e.TailApply(gen9575, gen9577)

			return

		}, 1)

		gen9579 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9580 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9583 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9584 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9585 := Call(__e, gen9584, V4783, Nil)

		gen9586 := Call(__e, gen9583, symhd, gen9585)

		gen9587 := Call(__e, gen9582, gen9586, Nil)

		gen9588 := Call(__e, gen9581, symcons_2, gen9587)

		gen9589 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9590 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9591 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9592 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9594 := Call(__e, gen9593, V4783, Nil)

		gen9595 := Call(__e, gen9592, symshen_4hdhd, gen9594)

		gen9596 := Call(__e, gen9591, gen9595, Nil)

		gen9597 := Call(__e, gen9590, symcons_2, gen9596)

		gen9598 := Call(__e, gen9589, gen9597, Nil)

		gen9599 := Call(__e, gen9580, gen9588, gen9598)

		gen9600 := Call(__e, gen9579, symand, gen9599)

		__e.TailApply(gen9578, gen9600)

		return

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1stream, gen9601)

	gen9686 := MakeNative(func(__e *ControlFlow) {
		V4786 := __e.Get(1)
		_ = V4786
		gen9683 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9684 := Call(__e, gen9683, V4786)

		var gen9685 Obj

		if True == gen9684 {
			gen9678 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9680 := Call(__e, gen9679, V4786)

			gen9681 := Call(__e, gen9678, symcons, gen9680)

			var gen9682 Obj

			if True == gen9681 {
				gen9673 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9674 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9675 := Call(__e, gen9674, V4786)

				gen9676 := Call(__e, gen9673, gen9675)

				var gen9677 Obj

				if True == gen9676 {
					gen9666 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9668 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9669 := Call(__e, gen9668, V4786)

					gen9670 := Call(__e, gen9667, gen9669)

					gen9671 := Call(__e, gen9666, gen9670)

					var gen9672 Obj

					if True == gen9671 {
						gen9657 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen9658 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9659 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9661 := Call(__e, gen9660, V4786)

						gen9662 := Call(__e, gen9659, gen9661)

						gen9663 := Call(__e, gen9658, gen9662)

						gen9664 := Call(__e, gen9657, Nil, gen9663)

						var gen9665 Obj

						if True == gen9664 {
							gen9649 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen9650 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9651 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9652 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9653 := Call(__e, gen9652, V4786)

							gen9654 := Call(__e, gen9651, gen9653)

							gen9655 := Call(__e, gen9650, gen9654)

							gen9656 := Call(__e, gen9649, Nil, gen9655)

							if True == gen9656 {
								gen9665 = True
							} else {
								gen9665 = False
							}

						} else {
							gen9665 = False
						}

						if True == gen9665 {
							gen9672 = True
						} else {
							gen9672 = False
						}

					} else {
						gen9672 = False
					}

					if True == gen9672 {
						gen9677 = True
					} else {
						gen9677 = False
					}

				} else {
					gen9677 = False
				}

				if True == gen9677 {
					gen9682 = True
				} else {
					gen9682 = False
				}

			} else {
				gen9682 = False
			}

			if True == gen9682 {
				gen9685 = True
			} else {
				gen9685 = False
			}

		} else {
			gen9685 = False
		}

		if True == gen9685 {
			gen9644 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9645 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen9647 := Call(__e, gen9646, V4786)

			gen9648 := Call(__e, gen9645, gen9647)

			__e.TailApply(gen9644, gen9648, Nil)

			return

		} else {
			gen9641 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9642 := Call(__e, gen9641, V4786)

			var gen9643 Obj

			if True == gen9642 {
				gen9636 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen9637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9638 := Call(__e, gen9637, V4786)

				gen9639 := Call(__e, gen9636, symcons, gen9638)

				var gen9640 Obj

				if True == gen9639 {
					gen9631 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9632 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9633 := Call(__e, gen9632, V4786)

					gen9634 := Call(__e, gen9631, gen9633)

					var gen9635 Obj

					if True == gen9634 {
						gen9624 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen9625 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9626 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9627 := Call(__e, gen9626, V4786)

						gen9628 := Call(__e, gen9625, gen9627)

						gen9629 := Call(__e, gen9624, gen9628)

						var gen9630 Obj

						if True == gen9629 {
							gen9616 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen9617 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9618 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9619 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9620 := Call(__e, gen9619, V4786)

							gen9621 := Call(__e, gen9618, gen9620)

							gen9622 := Call(__e, gen9617, gen9621)

							gen9623 := Call(__e, gen9616, Nil, gen9622)

							if True == gen9623 {
								gen9630 = True
							} else {
								gen9630 = False
							}

						} else {
							gen9630 = False
						}

						if True == gen9630 {
							gen9635 = True
						} else {
							gen9635 = False
						}

					} else {
						gen9635 = False
					}

					if True == gen9635 {
						gen9640 = True
					} else {
						gen9640 = False
					}

				} else {
					gen9640 = False
				}

				if True == gen9640 {
					gen9643 = True
				} else {
					gen9643 = False
				}

			} else {
				gen9643 = False
			}

			if True == gen9643 {
				gen9603 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9604 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9605 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9606 := Call(__e, gen9605, V4786)

				gen9607 := Call(__e, gen9604, gen9606)

				gen9608 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decons)

				gen9609 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9610 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9612 := Call(__e, gen9611, V4786)

				gen9613 := Call(__e, gen9610, gen9612)

				gen9614 := Call(__e, gen9609, gen9613)

				gen9615 := Call(__e, gen9608, gen9614)

				__e.TailApply(gen9603, gen9607, gen9615)

				return

			} else {
				if True == True {
					__e.Return(V4786)
					return
				} else {
					gen9602 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen9602, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4decons, gen9686)

	gen9730 := MakeNative(func(__e *ControlFlow) {
		V4801 := __e.Get(1)
		_ = V4801
		V4802 := __e.Get(2)
		_ = V4802
		V4803 := __e.Get(3)
		_ = V4803
		gen9727 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9728 := Call(__e, gen9727, V4803)

		var gen9729 Obj

		if True == gen9728 {
			gen9722 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen9723 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9724 := Call(__e, gen9723, V4803)

			gen9725 := Call(__e, gen9722, symshen_4pair, gen9724)

			var gen9726 Obj

			if True == gen9725 {
				gen9717 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen9718 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9719 := Call(__e, gen9718, V4803)

				gen9720 := Call(__e, gen9717, gen9719)

				var gen9721 Obj

				if True == gen9720 {
					gen9710 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9711 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9712 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9713 := Call(__e, gen9712, V4803)

					gen9714 := Call(__e, gen9711, gen9713)

					gen9715 := Call(__e, gen9710, gen9714)

					var gen9716 Obj

					if True == gen9715 {
						gen9701 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen9702 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9703 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9704 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen9705 := Call(__e, gen9704, V4803)

						gen9706 := Call(__e, gen9703, gen9705)

						gen9707 := Call(__e, gen9702, gen9706)

						gen9708 := Call(__e, gen9701, Nil, gen9707)

						var gen9709 Obj

						if True == gen9708 {
							gen9693 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen9694 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen9695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9696 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen9697 := Call(__e, gen9696, V4803)

							gen9698 := Call(__e, gen9695, gen9697)

							gen9699 := Call(__e, gen9694, gen9698)

							gen9700 := Call(__e, gen9693, gen9699, V4802)

							if True == gen9700 {
								gen9709 = True
							} else {
								gen9709 = False
							}

						} else {
							gen9709 = False
						}

						if True == gen9709 {
							gen9716 = True
						} else {
							gen9716 = False
						}

					} else {
						gen9716 = False
					}

					if True == gen9716 {
						gen9721 = True
					} else {
						gen9721 = False
					}

				} else {
					gen9721 = False
				}

				if True == gen9721 {
					gen9726 = True
				} else {
					gen9726 = False
				}

			} else {
				gen9726 = False
			}

			if True == gen9726 {
				gen9729 = True
			} else {
				gen9729 = False
			}

		} else {
			gen9729 = False
		}

		if True == gen9729 {
			__e.Return(V4801)
			return
		} else {
			gen9691 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9692 := Call(__e, gen9691, V4803)

			if True == gen9692 {
				gen9688 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				gen9690 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					gen9689 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1runon)

					__e.TailApply(gen9689, V4801, V4802, Z)

					return

				}, 1)

				__e.TailApply(gen9688, gen9690, V4803)

				return

			} else {
				if True == True {
					__e.Return(V4803)
					return
				} else {
					gen9687 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen9687, MakeString("no cond match"))

					return

				}
			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1runon, gen9730)

	gen9742 := MakeNative(func(__e *ControlFlow) {
		V4809 := __e.Get(1)
		_ = V4809
		gen9738 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		gen9739 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen9740 := Call(__e, gen9739, MakeString("."), V4809)

		gen9741 := Call(__e, gen9738, gen9740)

		if True == gen9741 {
			__e.Return(V4809)
			return
		} else {
			gen9736 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen9737 := Call(__e, gen9736, V4809)

			if True == gen9737 {
				gen9733 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1pathname)

				gen9734 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9735 := Call(__e, gen9734, V4809)

				__e.TailApply(gen9733, gen9735)

				return

			} else {
				if True == True {
					gen9732 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen9732, symshen_4strip_1pathname)

					return

				} else {
					gen9731 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen9731, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4strip_1pathname, gen9742)

	gen9801 := MakeNative(func(__e *ControlFlow) {
		V4813 := __e.Get(1)
		_ = V4813
		V4814 := __e.Get(2)
		_ = V4814
		V4815 := __e.Get(3)
		_ = V4815
		gen9799 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9800 := Call(__e, gen9799, V4813)

		if True == gen9800 {
			gen9792 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				gen9783 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					gen9780 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						gen9745 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9746 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9747 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

						gen9748 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9749 := Call(__e, gen9748, V4813)

						gen9750 := Call(__e, gen9747, symParse__, gen9749)

						gen9751 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9752 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9753 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9754 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9755 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9756 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9757 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9758 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9759 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9760 := Call(__e, gen9759, symfail, Nil)

						gen9761 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9762 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

						gen9763 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen9764 := Call(__e, gen9763, V4813)

						gen9765 := Call(__e, gen9762, symParse__, gen9764)

						gen9766 := Call(__e, gen9761, gen9765, Nil)

						gen9767 := Call(__e, gen9758, gen9760, gen9766)

						gen9768 := Call(__e, gen9757, sym_a, gen9767)

						gen9769 := Call(__e, gen9756, gen9768, Nil)

						gen9770 := Call(__e, gen9755, symnot, gen9769)

						gen9771 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9772 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9773 := Call(__e, gen9772, Else, Nil)

						gen9774 := Call(__e, gen9771, Action, gen9773)

						gen9775 := Call(__e, gen9754, gen9770, gen9774)

						gen9776 := Call(__e, gen9753, symif, gen9775)

						gen9777 := Call(__e, gen9752, gen9776, Nil)

						gen9778 := Call(__e, gen9751, Test, gen9777)

						gen9779 := Call(__e, gen9746, gen9750, gen9778)

						__e.TailApply(gen9745, symlet, gen9779)

						return

					}, 1)

					gen9781 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9782 := Call(__e, gen9781, symfail, Nil)

					__e.TailApply(gen9780, gen9782)

					return

				}, 1)

				gen9784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				gen9785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9786 := Call(__e, gen9785, V4813)

				gen9787 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				gen9788 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9789 := Call(__e, gen9788, V4813)

				gen9790 := Call(__e, gen9787, symParse__, gen9789)

				gen9791 := Call(__e, gen9784, gen9786, gen9790, V4815)

				__e.TailApply(gen9783, gen9791)

				return

			}, 1)

			gen9793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9795 := Call(__e, gen9794, V4813)

			gen9796 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9797 := Call(__e, gen9796, V4814, Nil)

			gen9798 := Call(__e, gen9793, gen9795, gen9797)

			__e.TailApply(gen9792, gen9798)

			return

		} else {
			if True == True {
				gen9744 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9744, symshen_4recursive__descent)

				return

			} else {
				gen9743 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9743, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__descent, gen9801)

	gen9860 := MakeNative(func(__e *ControlFlow) {
		V4819 := __e.Get(1)
		_ = V4819
		V4820 := __e.Get(2)
		_ = V4820
		V4821 := __e.Get(3)
		_ = V4821
		gen9858 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9859 := Call(__e, gen9858, V4819)

		if True == gen9859 {
			gen9849 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				gen9814 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					gen9811 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						gen9804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9807 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9808 := Call(__e, gen9807, Else, Nil)

						gen9809 := Call(__e, gen9806, Action, gen9808)

						gen9810 := Call(__e, gen9805, Test, gen9809)

						__e.TailApply(gen9804, symif, gen9810)

						return

					}, 1)

					gen9812 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9813 := Call(__e, gen9812, symfail, Nil)

					__e.TailApply(gen9811, gen9813)

					return

				}, 1)

				gen9815 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9817 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				gen9818 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen9819 := Call(__e, gen9818, V4819)

				gen9820 := Call(__e, gen9817, symParse__, gen9819)

				gen9821 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9822 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9823 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9824 := Call(__e, gen9823, V4820, Nil)

				gen9825 := Call(__e, gen9822, symshen_4hdhd, gen9824)

				gen9826 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9827 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				gen9828 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9829 := Call(__e, gen9828, V4819)

				gen9830 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9831 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9832 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9833 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9834 := Call(__e, gen9833, V4820, Nil)

				gen9835 := Call(__e, gen9832, symshen_4tlhd, gen9834)

				gen9836 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9837 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9838 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9839 := Call(__e, gen9838, V4820, Nil)

				gen9840 := Call(__e, gen9837, symshen_4hdtl, gen9839)

				gen9841 := Call(__e, gen9836, gen9840, Nil)

				gen9842 := Call(__e, gen9831, gen9835, gen9841)

				gen9843 := Call(__e, gen9830, symshen_4pair, gen9842)

				gen9844 := Call(__e, gen9827, gen9829, gen9843, V4821)

				gen9845 := Call(__e, gen9826, gen9844, Nil)

				gen9846 := Call(__e, gen9821, gen9825, gen9845)

				gen9847 := Call(__e, gen9816, gen9820, gen9846)

				gen9848 := Call(__e, gen9815, symlet, gen9847)

				__e.TailApply(gen9814, gen9848)

				return

			}, 1)

			gen9850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9851 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9852 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9853 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9854 := Call(__e, gen9853, V4820, Nil)

			gen9855 := Call(__e, gen9852, symhd, gen9854)

			gen9856 := Call(__e, gen9851, gen9855, Nil)

			gen9857 := Call(__e, gen9850, symcons_2, gen9856)

			__e.TailApply(gen9849, gen9857)

			return

		} else {
			if True == True {
				gen9803 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9803, symshen_4variable_1match)

				return

			} else {
				gen9802 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9802, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4variable_1match, gen9860)

	gen9866 := MakeNative(func(__e *ControlFlow) {
		V4831 := __e.Get(1)
		_ = V4831
		gen9864 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9865 := Call(__e, gen9864, V4831)

		if True == gen9865 {
			__e.Return(False)
			return
		} else {
			gen9862 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			gen9863 := Call(__e, gen9862, V4831)

			if True == gen9863 {
				__e.Return(False)
				return
			} else {
				if True == True {
					__e.Return(True)
					return
				} else {
					gen9861 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen9861, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4terminal_2, gen9866)

	gen9870 := MakeNative(func(__e *ControlFlow) {
		V4837 := __e.Get(1)
		_ = V4837
		gen9868 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen9869 := Call(__e, gen9868, V4837, sym__)

		if True == gen9869 {
			__e.Return(True)
			return
		} else {
			if True == True {
				__e.Return(False)
				return
			} else {
				gen9867 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9867, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream_2, gen9870)

	gen9942 := MakeNative(func(__e *ControlFlow) {
		V4841 := __e.Get(1)
		_ = V4841
		V4842 := __e.Get(2)
		_ = V4842
		V4843 := __e.Get(3)
		_ = V4843
		gen9940 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9941 := Call(__e, gen9940, V4841)

		if True == gen9941 {
			gen9913 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				gen9910 := MakeNative(func(__e *ControlFlow) {
					NewStr := __e.Get(1)
					_ = NewStr
					gen9883 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						gen9880 := MakeNative(func(__e *ControlFlow) {
							Else := __e.Get(1)
							_ = Else
							gen9873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen9874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen9875 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen9876 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen9877 := Call(__e, gen9876, Else, Nil)

							gen9878 := Call(__e, gen9875, Action, gen9877)

							gen9879 := Call(__e, gen9874, Test, gen9878)

							__e.TailApply(gen9873, symif, gen9879)

							return

						}, 1)

						gen9881 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9882 := Call(__e, gen9881, symfail, Nil)

						__e.TailApply(gen9880, gen9882)

						return

					}, 1)

					gen9884 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9885 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9886 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9887 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9888 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9891 := Call(__e, gen9890, V4842, Nil)

					gen9892 := Call(__e, gen9889, symshen_4tlhd, gen9891)

					gen9893 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9896 := Call(__e, gen9895, V4842, Nil)

					gen9897 := Call(__e, gen9894, symshen_4hdtl, gen9896)

					gen9898 := Call(__e, gen9893, gen9897, Nil)

					gen9899 := Call(__e, gen9888, gen9892, gen9898)

					gen9900 := Call(__e, gen9887, symshen_4pair, gen9899)

					gen9901 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9902 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

					gen9903 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen9904 := Call(__e, gen9903, V4841)

					gen9905 := Call(__e, gen9902, gen9904, NewStr, V4843)

					gen9906 := Call(__e, gen9901, gen9905, Nil)

					gen9907 := Call(__e, gen9886, gen9900, gen9906)

					gen9908 := Call(__e, gen9885, NewStr, gen9907)

					gen9909 := Call(__e, gen9884, symlet, gen9908)

					__e.TailApply(gen9883, gen9909)

					return

				}, 1)

				gen9911 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				gen9912 := Call(__e, gen9911, symNewStream)

				__e.TailApply(gen9910, gen9912)

				return

			}, 1)

			gen9914 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9915 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9916 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9917 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9918 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9919 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9920 := Call(__e, gen9919, V4842, Nil)

			gen9921 := Call(__e, gen9918, symhd, gen9920)

			gen9922 := Call(__e, gen9917, gen9921, Nil)

			gen9923 := Call(__e, gen9916, symcons_2, gen9922)

			gen9924 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9925 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9926 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9927 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen9928 := Call(__e, gen9927, V4841)

			gen9929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9931 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9932 := Call(__e, gen9931, V4842, Nil)

			gen9933 := Call(__e, gen9930, symshen_4hdhd, gen9932)

			gen9934 := Call(__e, gen9929, gen9933, Nil)

			gen9935 := Call(__e, gen9926, gen9928, gen9934)

			gen9936 := Call(__e, gen9925, sym_a, gen9935)

			gen9937 := Call(__e, gen9924, gen9936, Nil)

			gen9938 := Call(__e, gen9915, gen9923, gen9937)

			gen9939 := Call(__e, gen9914, symand, gen9938)

			__e.TailApply(gen9913, gen9939)

			return

		} else {
			if True == True {
				gen9872 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9872, symshen_4check__stream)

				return

			} else {
				gen9871 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9871, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4check__stream, gen9942)

	gen9985 := MakeNative(func(__e *ControlFlow) {
		V4847 := __e.Get(1)
		_ = V4847
		V4848 := __e.Get(2)
		_ = V4848
		V4849 := __e.Get(3)
		_ = V4849
		gen9983 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen9984 := Call(__e, gen9983, V4847)

		if True == gen9984 {
			gen9974 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				gen9955 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					gen9952 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						gen9945 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9946 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9947 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9948 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen9949 := Call(__e, gen9948, Else, Nil)

						gen9950 := Call(__e, gen9947, Action, gen9949)

						gen9951 := Call(__e, gen9946, Test, gen9950)

						__e.TailApply(gen9945, symif, gen9951)

						return

					}, 1)

					gen9953 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen9954 := Call(__e, gen9953, symfail, Nil)

					__e.TailApply(gen9952, gen9954)

					return

				}, 1)

				gen9956 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				gen9957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen9958 := Call(__e, gen9957, V4847)

				gen9959 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9960 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9961 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9962 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9963 := Call(__e, gen9962, V4848, Nil)

				gen9964 := Call(__e, gen9961, symshen_4tlhd, gen9963)

				gen9965 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9966 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9967 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9968 := Call(__e, gen9967, V4848, Nil)

				gen9969 := Call(__e, gen9966, symshen_4hdtl, gen9968)

				gen9970 := Call(__e, gen9965, gen9969, Nil)

				gen9971 := Call(__e, gen9960, gen9964, gen9970)

				gen9972 := Call(__e, gen9959, symshen_4pair, gen9971)

				gen9973 := Call(__e, gen9956, gen9958, gen9972, V4849)

				__e.TailApply(gen9955, gen9973)

				return

			}, 1)

			gen9975 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9976 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9977 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9978 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen9979 := Call(__e, gen9978, V4848, Nil)

			gen9980 := Call(__e, gen9977, symhd, gen9979)

			gen9981 := Call(__e, gen9976, gen9980, Nil)

			gen9982 := Call(__e, gen9975, symcons_2, gen9981)

			__e.TailApply(gen9974, gen9982)

			return

		} else {
			if True == True {
				gen9944 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen9944, symshen_4jump__stream)

				return

			} else {
				gen9943 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen9943, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream, gen9985)

	gen10004 := MakeNative(func(__e *ControlFlow) {
		V4851 := __e.Get(1)
		_ = V4851
		gen10002 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen10003 := Call(__e, gen10002, Nil, V4851)

		if True == gen10003 {
			__e.Return(Nil)
			return
		} else {
			gen10000 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

			gen10001 := Call(__e, gen10000, V4851)

			if True == gen10001 {
				gen9995 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9996 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen9997 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				gen9998 := Call(__e, gen9997, symParse__, V4851)

				gen9999 := Call(__e, gen9996, gen9998, Nil)

				__e.TailApply(gen9995, symshen_4hdtl, gen9999)

				return

			} else {
				gen9993 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				gen9994 := Call(__e, gen9993, V4851)

				if True == gen9994 {
					gen9992 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					__e.TailApply(gen9992, symParse__, V4851)

					return

				} else {
					gen9990 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen9991 := Call(__e, gen9990, V4851)

					if True == gen9991 {
						gen9987 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						gen9989 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							gen9988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

							__e.TailApply(gen9988, Z)

							return

						}, 1)

						__e.TailApply(gen9987, gen9989, V4851)

						return

					} else {
						if True == True {
							__e.Return(V4851)
							return
						} else {
							gen9986 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen9986, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4semantics, gen10004)

	gen10008 := MakeNative(func(__e *ControlFlow) {
		V4854 := __e.Get(1)
		_ = V4854
		V4855 := __e.Get(2)
		_ = V4855
		gen10005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen10007 := Call(__e, gen10006, V4855, Nil)

		__e.TailApply(gen10005, V4854, gen10007)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4pair, gen10008)

	gen10012 := MakeNative(func(__e *ControlFlow) {
		V4857 := __e.Get(1)
		_ = V4857
		gen10009 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10010 := Call(__e, PrimNS1Value(symns2_1value), symtl)

		gen10011 := Call(__e, gen10010, V4857)

		__e.TailApply(gen10009, gen10011)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4hdtl, gen10012)

	gen10016 := MakeNative(func(__e *ControlFlow) {
		V4859 := __e.Get(1)
		_ = V4859
		gen10013 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10014 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10015 := Call(__e, gen10014, V4859)

		__e.TailApply(gen10013, gen10015)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4hdhd, gen10016)

	gen10020 := MakeNative(func(__e *ControlFlow) {
		V4861 := __e.Get(1)
		_ = V4861
		gen10017 := Call(__e, PrimNS1Value(symns2_1value), symtl)

		gen10018 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen10019 := Call(__e, gen10018, V4861)

		__e.TailApply(gen10017, gen10019)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4tlhd, gen10020)

	gen10040 := MakeNative(func(__e *ControlFlow) {
		V4869 := __e.Get(1)
		_ = V4869
		gen10037 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10038 := Call(__e, gen10037, V4869)

		var gen10039 Obj

		if True == gen10038 {
			gen10032 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen10033 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen10034 := Call(__e, gen10033, V4869)

			gen10035 := Call(__e, gen10032, gen10034)

			var gen10036 Obj

			if True == gen10035 {
				gen10026 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen10027 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10028 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10029 := Call(__e, gen10028, V4869)

				gen10030 := Call(__e, gen10027, gen10029)

				gen10031 := Call(__e, gen10026, Nil, gen10030)

				if True == gen10031 {
					gen10036 = True
				} else {
					gen10036 = False
				}

			} else {
				gen10036 = False
			}

			if True == gen10036 {
				gen10039 = True
			} else {
				gen10039 = False
			}

		} else {
			gen10039 = False
		}

		if True == gen10039 {
			gen10023 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10024 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen10025 := Call(__e, gen10024, V4869)

			__e.TailApply(gen10023, gen10025)

			return

		} else {
			if True == True {
				gen10022 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen10022)

				return

			} else {
				gen10021 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10021, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4snd_1or_1fail, gen10040)

	gen10041 := MakeNative(func(__e *ControlFlow) {
		__e.Return(symshen_4fail_b)
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symfail, gen10041)

	gen10063 := MakeNative(func(__e *ControlFlow) {
		V4877 := __e.Get(1)
		_ = V4877
		gen10060 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10061 := Call(__e, gen10060, V4877)

		var gen10062 Obj

		if True == gen10061 {
			gen10055 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen10056 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen10057 := Call(__e, gen10056, V4877)

			gen10058 := Call(__e, gen10055, gen10057)

			var gen10059 Obj

			if True == gen10058 {
				gen10049 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen10050 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10051 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10052 := Call(__e, gen10051, V4877)

				gen10053 := Call(__e, gen10050, gen10052)

				gen10054 := Call(__e, gen10049, Nil, gen10053)

				if True == gen10054 {
					gen10059 = True
				} else {
					gen10059 = False
				}

			} else {
				gen10059 = False
			}

			if True == gen10059 {
				gen10062 = True
			} else {
				gen10062 = False
			}

		} else {
			gen10062 = False
		}

		if True == gen10062 {
			gen10044 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10045 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10046 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10047 := Call(__e, gen10046, V4877)

			gen10048 := Call(__e, gen10045, gen10047, Nil)

			__e.TailApply(gen10044, Nil, gen10048)

			return

		} else {
			if True == True {
				gen10043 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen10043)

				return

			} else {
				gen10042 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10042, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), sym_5_b_6, gen10063)

	gen10085 := MakeNative(func(__e *ControlFlow) {
		V4883 := __e.Get(1)
		_ = V4883
		gen10082 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen10083 := Call(__e, gen10082, V4883)

		var gen10084 Obj

		if True == gen10083 {
			gen10077 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen10078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen10079 := Call(__e, gen10078, V4883)

			gen10080 := Call(__e, gen10077, gen10079)

			var gen10081 Obj

			if True == gen10080 {
				gen10071 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen10072 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10073 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen10074 := Call(__e, gen10073, V4883)

				gen10075 := Call(__e, gen10072, gen10074)

				gen10076 := Call(__e, gen10071, Nil, gen10075)

				if True == gen10076 {
					gen10081 = True
				} else {
					gen10081 = False
				}

			} else {
				gen10081 = False
			}

			if True == gen10081 {
				gen10084 = True
			} else {
				gen10084 = False
			}

		} else {
			gen10084 = False
		}

		if True == gen10084 {
			gen10066 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10067 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen10068 := Call(__e, gen10067, V4883)

			gen10069 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen10070 := Call(__e, gen10069, Nil, Nil)

			__e.TailApply(gen10066, gen10068, gen10070)

			return

		} else {
			if True == True {
				gen10065 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen10065, sym_5e_6)

				return

			} else {
				gen10064 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen10064, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), sym_5e_6, gen10085)

	return

}, 0)
