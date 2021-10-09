package main

import . "github.com/tiancaiamao/shen-go/cora"

var YaccMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp24058 := MakeNative(func(__e *ControlFlow) {
		V4724 := __e.Get(1)
		_ = V4724
		tmp24072 := PrimIsPair(V4724)

		var ifres24064 Obj

		if True == tmp24072 {
			tmp24070 := PrimHead(V4724)

			tmp24071 := PrimEqual(symdefcc, tmp24070)

			var ifres24066 Obj

			if True == tmp24071 {
				tmp24068 := PrimTail(V4724)

				tmp24069 := PrimIsPair(tmp24068)

				var ifres24067 Obj

				if True == tmp24069 {
					ifres24067 = True

				} else {
					ifres24067 = False

				}

				ifres24066 = ifres24067

			} else {
				ifres24066 = False

			}

			var ifres24065 Obj

			if True == ifres24066 {
				ifres24065 = True

			} else {
				ifres24065 = False

			}

			ifres24064 = ifres24065

		} else {
			ifres24064 = False

		}

		if True == ifres24064 {
			tmp24060 := PrimTail(V4724)

			tmp24061 := PrimHead(tmp24060)

			tmp24062 := PrimTail(V4724)

			tmp24063 := PrimTail(tmp24062)

			__e.TailApply(PrimNS2Value(symshen_4yacc_1_6shen), tmp24061, tmp24063)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4yacc)
			return
		}

	}, 1)

	tmp24073 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc, tmp24058)

	_ = tmp24073

	tmp24074 := MakeNative(func(__e *ControlFlow) {
		V4727 := __e.Get(1)
		_ = V4727
		V4728 := __e.Get(2)
		_ = V4728
		tmp24075 := MakeNative(func(__e *ControlFlow) {
			CCRules := __e.Get(1)
			_ = CCRules
			tmp24076 := MakeNative(func(__e *ControlFlow) {
				CCBody := __e.Get(1)
				_ = CCBody
				tmp24077 := MakeNative(func(__e *ControlFlow) {
					YaccCases := __e.Get(1)
					_ = YaccCases
					tmp24078 := Call(__e, PrimNS2Value(symshen_4kill_1code), YaccCases)

					tmp24079 := PrimCons(tmp24078, Nil)

					tmp24080 := PrimCons(sym_1_6, tmp24079)

					tmp24081 := PrimCons(symStream, tmp24080)

					tmp24082 := PrimCons(V4727, tmp24081)

					__e.Return(PrimCons(symdefine, tmp24082))
					return

				}, 1)

				tmp24083 := Call(__e, PrimNS2Value(symshen_4yacc__cases), CCBody)

				__e.TailApply(tmp24077, tmp24083)
				return

			}, 1)

			tmp24084 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4cc__body), X)
				return
			}, 1)

			tmp24085 := Call(__e, PrimNS2Value(symmap), tmp24084, CCRules)

			__e.TailApply(tmp24076, tmp24085)
			return

		}, 1)

		tmp24086 := Call(__e, PrimNS2Value(symshen_4split__cc__rules), True, V4728, Nil)

		__e.TailApply(tmp24075, tmp24086)
		return

	}, 2)

	tmp24087 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc_1_6shen, tmp24074)

	_ = tmp24087

	tmp24088 := MakeNative(func(__e *ControlFlow) {
		V4730 := __e.Get(1)
		_ = V4730
		tmp24097 := Call(__e, PrimNS2Value(symoccurrences), symkill, V4730)

		tmp24098 := PrimGreatThan(tmp24097, MakeNumber(0))

		if True == tmp24098 {
			tmp24090 := PrimCons(symE, Nil)

			tmp24091 := PrimCons(symshen_4analyse_1kill, tmp24090)

			tmp24092 := PrimCons(tmp24091, Nil)

			tmp24093 := PrimCons(symE, tmp24092)

			tmp24094 := PrimCons(symlambda, tmp24093)

			tmp24095 := PrimCons(tmp24094, Nil)

			tmp24096 := PrimCons(V4730, tmp24095)

			__e.Return(PrimCons(symtrap_1error, tmp24096))
			return

		} else {
			__e.Return(V4730)
			return
		}

	}, 1)

	tmp24099 := Call(__e, PrimNS1Value(symns2_1set), symshen_4kill_1code, tmp24088)

	_ = tmp24099

	tmp24100 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimSimpleError(MakeString("yacc kill")))
		return
	}, 0)

	tmp24101 := Call(__e, PrimNS1Value(symns2_1set), symkill, tmp24100)

	_ = tmp24101

	tmp24102 := MakeNative(func(__e *ControlFlow) {
		V4732 := __e.Get(1)
		_ = V4732
		tmp24103 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp24105 := PrimEqual(String, MakeString("yacc kill"))

			if True == tmp24105 {
				__e.TailApply(PrimNS2Value(symfail))
				return
			} else {
				__e.Return(V4732)
				return
			}

		}, 1)

		tmp24106 := PrimErrorToString(V4732)

		__e.TailApply(tmp24103, tmp24106)
		return

	}, 1)

	tmp24107 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1kill, tmp24102)

	_ = tmp24107

	tmp24108 := MakeNative(func(__e *ControlFlow) {
		V4738 := __e.Get(1)
		_ = V4738
		V4739 := __e.Get(2)
		_ = V4739
		V4740 := __e.Get(3)
		_ = V4740
		tmp24132 := PrimEqual(Nil, V4739)

		var ifres24129 Obj

		if True == tmp24132 {
			tmp24131 := PrimEqual(Nil, V4740)

			var ifres24130 Obj

			if True == tmp24131 {
				ifres24130 = True

			} else {
				ifres24130 = False

			}

			ifres24129 = ifres24130

		} else {
			ifres24129 = False

		}

		if True == ifres24129 {
			__e.Return(Nil)
			return
		} else {
			tmp24128 := PrimEqual(Nil, V4739)

			if True == tmp24128 {
				tmp24126 := Call(__e, PrimNS2Value(symreverse), V4740)

				tmp24127 := Call(__e, PrimNS2Value(symshen_4split__cc__rule), V4738, tmp24126, Nil)

				__e.Return(PrimCons(tmp24127, Nil))
				return

			} else {
				tmp24125 := PrimIsPair(V4739)

				var ifres24121 Obj

				if True == tmp24125 {
					tmp24123 := PrimHead(V4739)

					tmp24124 := PrimEqual(sym_k, tmp24123)

					var ifres24122 Obj

					if True == tmp24124 {
						ifres24122 = True

					} else {
						ifres24122 = False

					}

					ifres24121 = ifres24122

				} else {
					ifres24121 = False

				}

				if True == ifres24121 {
					tmp24117 := Call(__e, PrimNS2Value(symreverse), V4740)

					tmp24118 := Call(__e, PrimNS2Value(symshen_4split__cc__rule), V4738, tmp24117, Nil)

					tmp24119 := PrimTail(V4739)

					tmp24120 := Call(__e, PrimNS2Value(symshen_4split__cc__rules), V4738, tmp24119, Nil)

					__e.Return(PrimCons(tmp24118, tmp24120))
					return

				} else {
					tmp24116 := PrimIsPair(V4739)

					if True == tmp24116 {
						tmp24113 := PrimTail(V4739)

						tmp24114 := PrimHead(V4739)

						tmp24115 := PrimCons(tmp24114, V4740)

						__e.TailApply(PrimNS2Value(symshen_4split__cc__rules), V4738, tmp24113, tmp24115)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4split__cc__rules)
						return
					}

				}

			}

		}

	}, 3)

	tmp24133 := Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rules, tmp24108)

	_ = tmp24133

	tmp24134 := MakeNative(func(__e *ControlFlow) {
		V4748 := __e.Get(1)
		_ = V4748
		V4749 := __e.Get(2)
		_ = V4749
		V4750 := __e.Get(3)
		_ = V4750
		tmp24208 := PrimIsPair(V4749)

		var ifres24195 Obj

		if True == tmp24208 {
			tmp24206 := PrimHead(V4749)

			tmp24207 := PrimEqual(sym_h_a, tmp24206)

			var ifres24197 Obj

			if True == tmp24207 {
				tmp24204 := PrimTail(V4749)

				tmp24205 := PrimIsPair(tmp24204)

				var ifres24199 Obj

				if True == tmp24205 {
					tmp24201 := PrimTail(V4749)

					tmp24202 := PrimTail(tmp24201)

					tmp24203 := PrimEqual(Nil, tmp24202)

					var ifres24200 Obj

					if True == tmp24203 {
						ifres24200 = True

					} else {
						ifres24200 = False

					}

					ifres24199 = ifres24200

				} else {
					ifres24199 = False

				}

				var ifres24198 Obj

				if True == ifres24199 {
					ifres24198 = True

				} else {
					ifres24198 = False

				}

				ifres24197 = ifres24198

			} else {
				ifres24197 = False

			}

			var ifres24196 Obj

			if True == ifres24197 {
				ifres24196 = True

			} else {
				ifres24196 = False

			}

			ifres24195 = ifres24196

		} else {
			ifres24195 = False

		}

		if True == ifres24195 {
			tmp24193 := Call(__e, PrimNS2Value(symreverse), V4750)

			tmp24194 := PrimTail(V4749)

			__e.Return(PrimCons(tmp24193, tmp24194))
			return

		} else {
			tmp24192 := PrimIsPair(V4749)

			var ifres24160 Obj

			if True == tmp24192 {
				tmp24190 := PrimHead(V4749)

				tmp24191 := PrimEqual(sym_h_a, tmp24190)

				var ifres24162 Obj

				if True == tmp24191 {
					tmp24188 := PrimTail(V4749)

					tmp24189 := PrimIsPair(tmp24188)

					var ifres24164 Obj

					if True == tmp24189 {
						tmp24185 := PrimTail(V4749)

						tmp24186 := PrimTail(tmp24185)

						tmp24187 := PrimIsPair(tmp24186)

						var ifres24166 Obj

						if True == tmp24187 {
							tmp24181 := PrimTail(V4749)

							tmp24182 := PrimTail(tmp24181)

							tmp24183 := PrimHead(tmp24182)

							tmp24184 := PrimEqual(symwhere, tmp24183)

							var ifres24168 Obj

							if True == tmp24184 {
								tmp24177 := PrimTail(V4749)

								tmp24178 := PrimTail(tmp24177)

								tmp24179 := PrimTail(tmp24178)

								tmp24180 := PrimIsPair(tmp24179)

								var ifres24170 Obj

								if True == tmp24180 {
									tmp24172 := PrimTail(V4749)

									tmp24173 := PrimTail(tmp24172)

									tmp24174 := PrimTail(tmp24173)

									tmp24175 := PrimTail(tmp24174)

									tmp24176 := PrimEqual(Nil, tmp24175)

									var ifres24171 Obj

									if True == tmp24176 {
										ifres24171 = True

									} else {
										ifres24171 = False

									}

									ifres24170 = ifres24171

								} else {
									ifres24170 = False

								}

								var ifres24169 Obj

								if True == ifres24170 {
									ifres24169 = True

								} else {
									ifres24169 = False

								}

								ifres24168 = ifres24169

							} else {
								ifres24168 = False

							}

							var ifres24167 Obj

							if True == ifres24168 {
								ifres24167 = True

							} else {
								ifres24167 = False

							}

							ifres24166 = ifres24167

						} else {
							ifres24166 = False

						}

						var ifres24165 Obj

						if True == ifres24166 {
							ifres24165 = True

						} else {
							ifres24165 = False

						}

						ifres24164 = ifres24165

					} else {
						ifres24164 = False

					}

					var ifres24163 Obj

					if True == ifres24164 {
						ifres24163 = True

					} else {
						ifres24163 = False

					}

					ifres24162 = ifres24163

				} else {
					ifres24162 = False

				}

				var ifres24161 Obj

				if True == ifres24162 {
					ifres24161 = True

				} else {
					ifres24161 = False

				}

				ifres24160 = ifres24161

			} else {
				ifres24160 = False

			}

			if True == ifres24160 {
				tmp24149 := Call(__e, PrimNS2Value(symreverse), V4750)

				tmp24150 := PrimTail(V4749)

				tmp24151 := PrimTail(tmp24150)

				tmp24152 := PrimTail(tmp24151)

				tmp24153 := PrimHead(tmp24152)

				tmp24154 := PrimTail(V4749)

				tmp24155 := PrimHead(tmp24154)

				tmp24156 := PrimCons(tmp24155, Nil)

				tmp24157 := PrimCons(tmp24153, tmp24156)

				tmp24158 := PrimCons(symwhere, tmp24157)

				tmp24159 := PrimCons(tmp24158, Nil)

				__e.Return(PrimCons(tmp24149, tmp24159))
				return

			} else {
				tmp24148 := PrimEqual(Nil, V4749)

				if True == tmp24148 {
					tmp24143 := Call(__e, PrimNS2Value(symshen_4semantic_1completion_1warning), V4748, V4750)

					_ = tmp24143

					tmp24144 := Call(__e, PrimNS2Value(symreverse), V4750)

					tmp24145 := Call(__e, PrimNS2Value(symshen_4default__semantics), tmp24144)

					tmp24146 := PrimCons(tmp24145, Nil)

					tmp24147 := PrimCons(sym_h_a, tmp24146)

					__e.TailApply(PrimNS2Value(symshen_4split__cc__rule), V4748, tmp24147, V4750)
					return

				} else {
					tmp24142 := PrimIsPair(V4749)

					if True == tmp24142 {
						tmp24139 := PrimTail(V4749)

						tmp24140 := PrimHead(V4749)

						tmp24141 := PrimCons(tmp24140, V4750)

						__e.TailApply(PrimNS2Value(symshen_4split__cc__rule), V4748, tmp24139, tmp24141)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4split__cc__rule)
						return
					}

				}

			}

		}

	}, 3)

	tmp24209 := Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rule, tmp24134)

	_ = tmp24209

	tmp24210 := MakeNative(func(__e *ControlFlow) {
		V4761 := __e.Get(1)
		_ = V4761
		V4762 := __e.Get(2)
		_ = V4762
		tmp24220 := PrimEqual(True, V4761)

		if True == tmp24220 {
			tmp24212 := Call(__e, PrimNS2Value(symstoutput))

			tmp24213 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString("warning: "), tmp24212)

			_ = tmp24213

			tmp24214 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp24215 := Call(__e, PrimNS2Value(symshen_4app), X, MakeString(" "), symshen_4a)

				tmp24216 := Call(__e, PrimNS2Value(symstoutput))

				__e.TailApply(PrimNS2Value(symshen_4prhush), tmp24215, tmp24216)
				return

			}, 1)

			tmp24217 := Call(__e, PrimNS2Value(symreverse), V4762)

			tmp24218 := Call(__e, PrimNS2Value(symshen_4for_1each), tmp24214, tmp24217)

			_ = tmp24218

			tmp24219 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), MakeString("has no semantics.\n"), tmp24219)
			return

		} else {
			__e.Return(symshen_4skip)
			return
		}

	}, 2)

	tmp24221 := Call(__e, PrimNS1Value(symns2_1set), symshen_4semantic_1completion_1warning, tmp24210)

	_ = tmp24221

	tmp24222 := MakeNative(func(__e *ControlFlow) {
		V4764 := __e.Get(1)
		_ = V4764
		tmp24252 := PrimEqual(Nil, V4764)

		if True == tmp24252 {
			__e.Return(Nil)
			return
		} else {
			tmp24251 := PrimIsPair(V4764)

			var ifres24243 Obj

			if True == tmp24251 {
				tmp24249 := PrimTail(V4764)

				tmp24250 := PrimEqual(Nil, tmp24249)

				var ifres24245 Obj

				if True == tmp24250 {
					tmp24247 := PrimHead(V4764)

					tmp24248 := Call(__e, PrimNS2Value(symshen_4grammar__symbol_2), tmp24247)

					var ifres24246 Obj

					if True == tmp24248 {
						ifres24246 = True

					} else {
						ifres24246 = False

					}

					ifres24245 = ifres24246

				} else {
					ifres24245 = False

				}

				var ifres24244 Obj

				if True == ifres24245 {
					ifres24244 = True

				} else {
					ifres24244 = False

				}

				ifres24243 = ifres24244

			} else {
				ifres24243 = False

			}

			if True == ifres24243 {
				__e.Return(PrimHead(V4764))
				return
			} else {
				tmp24242 := PrimIsPair(V4764)

				var ifres24238 Obj

				if True == tmp24242 {
					tmp24240 := PrimHead(V4764)

					tmp24241 := Call(__e, PrimNS2Value(symshen_4grammar__symbol_2), tmp24240)

					var ifres24239 Obj

					if True == tmp24241 {
						ifres24239 = True

					} else {
						ifres24239 = False

					}

					ifres24238 = ifres24239

				} else {
					ifres24238 = False

				}

				if True == ifres24238 {
					tmp24233 := PrimHead(V4764)

					tmp24234 := PrimTail(V4764)

					tmp24235 := Call(__e, PrimNS2Value(symshen_4default__semantics), tmp24234)

					tmp24236 := PrimCons(tmp24235, Nil)

					tmp24237 := PrimCons(tmp24233, tmp24236)

					__e.Return(PrimCons(symappend, tmp24237))
					return

				} else {
					tmp24232 := PrimIsPair(V4764)

					if True == tmp24232 {
						tmp24227 := PrimHead(V4764)

						tmp24228 := PrimTail(V4764)

						tmp24229 := Call(__e, PrimNS2Value(symshen_4default__semantics), tmp24228)

						tmp24230 := PrimCons(tmp24229, Nil)

						tmp24231 := PrimCons(tmp24227, tmp24230)

						__e.Return(PrimCons(symcons, tmp24231))
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4default__semantics)
						return
					}

				}

			}

		}

	}, 1)

	tmp24253 := Call(__e, PrimNS1Value(symns2_1set), symshen_4default__semantics, tmp24222)

	_ = tmp24253

	tmp24254 := MakeNative(func(__e *ControlFlow) {
		V4766 := __e.Get(1)
		_ = V4766
		tmp24268 := PrimIsSymbol(V4766)

		if True == tmp24268 {
			tmp24257 := MakeNative(func(__e *ControlFlow) {
				Cs := __e.Get(1)
				_ = Cs
				tmp24263 := PrimHead(Cs)

				tmp24264 := PrimEqual(tmp24263, MakeString("<"))

				if True == tmp24264 {
					tmp24260 := Call(__e, PrimNS2Value(symreverse), Cs)

					tmp24261 := PrimHead(tmp24260)

					tmp24262 := PrimEqual(tmp24261, MakeString(">"))

					if True == tmp24262 {
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

			tmp24265 := Call(__e, PrimNS2Value(symexplode), V4766)

			tmp24266 := Call(__e, PrimNS2Value(symshen_4strip_1pathname), tmp24265)

			tmp24267 := Call(__e, tmp24257, tmp24266)

			if True == tmp24267 {
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

	tmp24269 := Call(__e, PrimNS1Value(symns2_1set), symshen_4grammar__symbol_2, tmp24254)

	_ = tmp24269

	tmp24270 := MakeNative(func(__e *ControlFlow) {
		V4768 := __e.Get(1)
		_ = V4768
		tmp24293 := PrimIsPair(V4768)

		var ifres24289 Obj

		if True == tmp24293 {
			tmp24291 := PrimTail(V4768)

			tmp24292 := PrimEqual(Nil, tmp24291)

			var ifres24290 Obj

			if True == tmp24292 {
				ifres24290 = True

			} else {
				ifres24290 = False

			}

			ifres24289 = ifres24290

		} else {
			ifres24289 = False

		}

		if True == ifres24289 {
			__e.Return(PrimHead(V4768))
			return
		} else {
			tmp24288 := PrimIsPair(V4768)

			if True == tmp24288 {
				tmp24273 := MakeNative(func(__e *ControlFlow) {
					P := __e.Get(1)
					_ = P
					tmp24274 := PrimHead(V4768)

					tmp24275 := PrimCons(symfail, Nil)

					tmp24276 := PrimCons(tmp24275, Nil)

					tmp24277 := PrimCons(P, tmp24276)

					tmp24278 := PrimCons(sym_a, tmp24277)

					tmp24279 := PrimTail(V4768)

					tmp24280 := Call(__e, PrimNS2Value(symshen_4yacc__cases), tmp24279)

					tmp24281 := PrimCons(P, Nil)

					tmp24282 := PrimCons(tmp24280, tmp24281)

					tmp24283 := PrimCons(tmp24278, tmp24282)

					tmp24284 := PrimCons(symif, tmp24283)

					tmp24285 := PrimCons(tmp24284, Nil)

					tmp24286 := PrimCons(tmp24274, tmp24285)

					tmp24287 := PrimCons(P, tmp24286)

					__e.Return(PrimCons(symlet, tmp24287))
					return

				}, 1)

				__e.TailApply(tmp24273, symYaccParse)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4yacc__cases)
				return
			}

		}

	}, 1)

	tmp24294 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc__cases, tmp24270)

	_ = tmp24294

	tmp24295 := MakeNative(func(__e *ControlFlow) {
		V4770 := __e.Get(1)
		_ = V4770
		tmp24309 := PrimIsPair(V4770)

		var ifres24300 Obj

		if True == tmp24309 {
			tmp24307 := PrimTail(V4770)

			tmp24308 := PrimIsPair(tmp24307)

			var ifres24302 Obj

			if True == tmp24308 {
				tmp24304 := PrimTail(V4770)

				tmp24305 := PrimTail(tmp24304)

				tmp24306 := PrimEqual(Nil, tmp24305)

				var ifres24303 Obj

				if True == tmp24306 {
					ifres24303 = True

				} else {
					ifres24303 = False

				}

				ifres24302 = ifres24303

			} else {
				ifres24302 = False

			}

			var ifres24301 Obj

			if True == ifres24302 {
				ifres24301 = True

			} else {
				ifres24301 = False

			}

			ifres24300 = ifres24301

		} else {
			ifres24300 = False

		}

		if True == ifres24300 {
			tmp24297 := PrimHead(V4770)

			tmp24298 := PrimTail(V4770)

			tmp24299 := PrimHead(tmp24298)

			__e.TailApply(PrimNS2Value(symshen_4syntax), tmp24297, symStream, tmp24299)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4cc__body)
			return
		}

	}, 1)

	tmp24310 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cc__body, tmp24295)

	_ = tmp24310

	tmp24311 := MakeNative(func(__e *ControlFlow) {
		V4774 := __e.Get(1)
		_ = V4774
		V4775 := __e.Get(2)
		_ = V4775
		V4776 := __e.Get(3)
		_ = V4776
		tmp24380 := PrimEqual(Nil, V4774)

		var ifres24358 Obj

		if True == tmp24380 {
			tmp24379 := PrimIsPair(V4776)

			var ifres24360 Obj

			if True == tmp24379 {
				tmp24377 := PrimHead(V4776)

				tmp24378 := PrimEqual(symwhere, tmp24377)

				var ifres24362 Obj

				if True == tmp24378 {
					tmp24375 := PrimTail(V4776)

					tmp24376 := PrimIsPair(tmp24375)

					var ifres24364 Obj

					if True == tmp24376 {
						tmp24372 := PrimTail(V4776)

						tmp24373 := PrimTail(tmp24372)

						tmp24374 := PrimIsPair(tmp24373)

						var ifres24366 Obj

						if True == tmp24374 {
							tmp24368 := PrimTail(V4776)

							tmp24369 := PrimTail(tmp24368)

							tmp24370 := PrimTail(tmp24369)

							tmp24371 := PrimEqual(Nil, tmp24370)

							var ifres24367 Obj

							if True == tmp24371 {
								ifres24367 = True

							} else {
								ifres24367 = False

							}

							ifres24366 = ifres24367

						} else {
							ifres24366 = False

						}

						var ifres24365 Obj

						if True == ifres24366 {
							ifres24365 = True

						} else {
							ifres24365 = False

						}

						ifres24364 = ifres24365

					} else {
						ifres24364 = False

					}

					var ifres24363 Obj

					if True == ifres24364 {
						ifres24363 = True

					} else {
						ifres24363 = False

					}

					ifres24362 = ifres24363

				} else {
					ifres24362 = False

				}

				var ifres24361 Obj

				if True == ifres24362 {
					ifres24361 = True

				} else {
					ifres24361 = False

				}

				ifres24360 = ifres24361

			} else {
				ifres24360 = False

			}

			var ifres24359 Obj

			if True == ifres24360 {
				ifres24359 = True

			} else {
				ifres24359 = False

			}

			ifres24358 = ifres24359

		} else {
			ifres24358 = False

		}

		if True == ifres24358 {
			tmp24342 := PrimTail(V4776)

			tmp24343 := PrimHead(tmp24342)

			tmp24344 := Call(__e, PrimNS2Value(symshen_4semantics), tmp24343)

			tmp24345 := PrimCons(V4775, Nil)

			tmp24346 := PrimCons(symhd, tmp24345)

			tmp24347 := PrimTail(V4776)

			tmp24348 := PrimTail(tmp24347)

			tmp24349 := PrimHead(tmp24348)

			tmp24350 := Call(__e, PrimNS2Value(symshen_4semantics), tmp24349)

			tmp24351 := PrimCons(tmp24350, Nil)

			tmp24352 := PrimCons(tmp24346, tmp24351)

			tmp24353 := PrimCons(symshen_4pair, tmp24352)

			tmp24354 := PrimCons(symfail, Nil)

			tmp24355 := PrimCons(tmp24354, Nil)

			tmp24356 := PrimCons(tmp24353, tmp24355)

			tmp24357 := PrimCons(tmp24344, tmp24356)

			__e.Return(PrimCons(symif, tmp24357))
			return

		} else {
			tmp24341 := PrimEqual(Nil, V4774)

			if True == tmp24341 {
				tmp24336 := PrimCons(V4775, Nil)

				tmp24337 := PrimCons(symhd, tmp24336)

				tmp24338 := Call(__e, PrimNS2Value(symshen_4semantics), V4776)

				tmp24339 := PrimCons(tmp24338, Nil)

				tmp24340 := PrimCons(tmp24337, tmp24339)

				__e.Return(PrimCons(symshen_4pair, tmp24340))
				return

			} else {
				tmp24335 := PrimIsPair(V4774)

				if True == tmp24335 {
					tmp24333 := PrimHead(V4774)

					tmp24334 := Call(__e, PrimNS2Value(symshen_4grammar__symbol_2), tmp24333)

					if True == tmp24334 {
						__e.TailApply(PrimNS2Value(symshen_4recursive__descent), V4774, V4775, V4776)
						return
					} else {
						tmp24331 := PrimHead(V4774)

						tmp24332 := PrimIsVariable(tmp24331)

						if True == tmp24332 {
							__e.TailApply(PrimNS2Value(symshen_4variable_1match), V4774, V4775, V4776)
							return
						} else {
							tmp24329 := PrimHead(V4774)

							tmp24330 := Call(__e, PrimNS2Value(symshen_4jump__stream_2), tmp24329)

							if True == tmp24330 {
								__e.TailApply(PrimNS2Value(symshen_4jump__stream), V4774, V4775, V4776)
								return
							} else {
								tmp24327 := PrimHead(V4774)

								tmp24328 := Call(__e, PrimNS2Value(symshen_4terminal_2), tmp24327)

								if True == tmp24328 {
									__e.TailApply(PrimNS2Value(symshen_4check__stream), V4774, V4775, V4776)
									return
								} else {
									tmp24325 := PrimHead(V4774)

									tmp24326 := PrimIsPair(tmp24325)

									if True == tmp24326 {
										tmp24322 := PrimHead(V4774)

										tmp24323 := Call(__e, PrimNS2Value(symshen_4decons), tmp24322)

										tmp24324 := PrimTail(V4774)

										__e.TailApply(PrimNS2Value(symshen_4list_1stream), tmp24323, tmp24324, V4775, V4776)
										return

									} else {
										tmp24320 := PrimHead(V4774)

										tmp24321 := Call(__e, PrimNS2Value(symshen_4app), tmp24320, MakeString(" is not legal syntax\n"), symshen_4a)

										__e.Return(PrimSimpleError(tmp24321))
										return

									}

								}

							}

						}

					}

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4syntax)
					return
				}

			}

		}

	}, 3)

	tmp24381 := Call(__e, PrimNS1Value(symns2_1set), symshen_4syntax, tmp24311)

	_ = tmp24381

	tmp24382 := MakeNative(func(__e *ControlFlow) {
		V4781 := __e.Get(1)
		_ = V4781
		V4782 := __e.Get(2)
		_ = V4782
		V4783 := __e.Get(3)
		_ = V4783
		V4784 := __e.Get(4)
		_ = V4784
		tmp24383 := MakeNative(func(__e *ControlFlow) {
			Test := __e.Get(1)
			_ = Test
			tmp24384 := MakeNative(func(__e *ControlFlow) {
				Placeholder := __e.Get(1)
				_ = Placeholder
				tmp24385 := MakeNative(func(__e *ControlFlow) {
					RunOn := __e.Get(1)
					_ = RunOn
					tmp24386 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						tmp24387 := PrimCons(symfail, Nil)

						tmp24388 := PrimCons(tmp24387, Nil)

						tmp24389 := PrimCons(Action, tmp24388)

						tmp24390 := PrimCons(Test, tmp24389)

						__e.Return(PrimCons(symif, tmp24390))
						return

					}, 1)

					tmp24391 := PrimCons(V4783, Nil)

					tmp24392 := PrimCons(symshen_4hdhd, tmp24391)

					tmp24393 := PrimCons(V4783, Nil)

					tmp24394 := PrimCons(symshen_4hdtl, tmp24393)

					tmp24395 := PrimCons(tmp24394, Nil)

					tmp24396 := PrimCons(tmp24392, tmp24395)

					tmp24397 := PrimCons(symshen_4pair, tmp24396)

					tmp24398 := Call(__e, PrimNS2Value(symshen_4syntax), V4781, tmp24397, Placeholder)

					tmp24399 := Call(__e, PrimNS2Value(symshen_4insert_1runon), RunOn, Placeholder, tmp24398)

					__e.TailApply(tmp24386, tmp24399)
					return

				}, 1)

				tmp24400 := PrimCons(V4783, Nil)

				tmp24401 := PrimCons(symshen_4tlhd, tmp24400)

				tmp24402 := PrimCons(V4783, Nil)

				tmp24403 := PrimCons(symshen_4hdtl, tmp24402)

				tmp24404 := PrimCons(tmp24403, Nil)

				tmp24405 := PrimCons(tmp24401, tmp24404)

				tmp24406 := PrimCons(symshen_4pair, tmp24405)

				tmp24407 := Call(__e, PrimNS2Value(symshen_4syntax), V4782, tmp24406, V4784)

				__e.TailApply(tmp24385, tmp24407)
				return

			}, 1)

			tmp24408 := Call(__e, PrimNS2Value(symgensym), symshen_4place)

			__e.TailApply(tmp24384, tmp24408)
			return

		}, 1)

		tmp24409 := PrimCons(V4783, Nil)

		tmp24410 := PrimCons(symhd, tmp24409)

		tmp24411 := PrimCons(tmp24410, Nil)

		tmp24412 := PrimCons(symcons_2, tmp24411)

		tmp24413 := PrimCons(V4783, Nil)

		tmp24414 := PrimCons(symshen_4hdhd, tmp24413)

		tmp24415 := PrimCons(tmp24414, Nil)

		tmp24416 := PrimCons(symcons_2, tmp24415)

		tmp24417 := PrimCons(tmp24416, Nil)

		tmp24418 := PrimCons(tmp24412, tmp24417)

		tmp24419 := PrimCons(symand, tmp24418)

		__e.TailApply(tmp24383, tmp24419)
		return

	}, 4)

	tmp24420 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1stream, tmp24382)

	_ = tmp24420

	tmp24421 := MakeNative(func(__e *ControlFlow) {
		V4786 := __e.Get(1)
		_ = V4786
		tmp24477 := PrimIsPair(V4786)

		var ifres24452 Obj

		if True == tmp24477 {
			tmp24475 := PrimHead(V4786)

			tmp24476 := PrimEqual(symcons, tmp24475)

			var ifres24454 Obj

			if True == tmp24476 {
				tmp24473 := PrimTail(V4786)

				tmp24474 := PrimIsPair(tmp24473)

				var ifres24456 Obj

				if True == tmp24474 {
					tmp24470 := PrimTail(V4786)

					tmp24471 := PrimTail(tmp24470)

					tmp24472 := PrimIsPair(tmp24471)

					var ifres24458 Obj

					if True == tmp24472 {
						tmp24466 := PrimTail(V4786)

						tmp24467 := PrimTail(tmp24466)

						tmp24468 := PrimHead(tmp24467)

						tmp24469 := PrimEqual(Nil, tmp24468)

						var ifres24460 Obj

						if True == tmp24469 {
							tmp24462 := PrimTail(V4786)

							tmp24463 := PrimTail(tmp24462)

							tmp24464 := PrimTail(tmp24463)

							tmp24465 := PrimEqual(Nil, tmp24464)

							var ifres24461 Obj

							if True == tmp24465 {
								ifres24461 = True

							} else {
								ifres24461 = False

							}

							ifres24460 = ifres24461

						} else {
							ifres24460 = False

						}

						var ifres24459 Obj

						if True == ifres24460 {
							ifres24459 = True

						} else {
							ifres24459 = False

						}

						ifres24458 = ifres24459

					} else {
						ifres24458 = False

					}

					var ifres24457 Obj

					if True == ifres24458 {
						ifres24457 = True

					} else {
						ifres24457 = False

					}

					ifres24456 = ifres24457

				} else {
					ifres24456 = False

				}

				var ifres24455 Obj

				if True == ifres24456 {
					ifres24455 = True

				} else {
					ifres24455 = False

				}

				ifres24454 = ifres24455

			} else {
				ifres24454 = False

			}

			var ifres24453 Obj

			if True == ifres24454 {
				ifres24453 = True

			} else {
				ifres24453 = False

			}

			ifres24452 = ifres24453

		} else {
			ifres24452 = False

		}

		if True == ifres24452 {
			tmp24450 := PrimTail(V4786)

			tmp24451 := PrimHead(tmp24450)

			__e.Return(PrimCons(tmp24451, Nil))
			return

		} else {
			tmp24449 := PrimIsPair(V4786)

			var ifres24430 Obj

			if True == tmp24449 {
				tmp24447 := PrimHead(V4786)

				tmp24448 := PrimEqual(symcons, tmp24447)

				var ifres24432 Obj

				if True == tmp24448 {
					tmp24445 := PrimTail(V4786)

					tmp24446 := PrimIsPair(tmp24445)

					var ifres24434 Obj

					if True == tmp24446 {
						tmp24442 := PrimTail(V4786)

						tmp24443 := PrimTail(tmp24442)

						tmp24444 := PrimIsPair(tmp24443)

						var ifres24436 Obj

						if True == tmp24444 {
							tmp24438 := PrimTail(V4786)

							tmp24439 := PrimTail(tmp24438)

							tmp24440 := PrimTail(tmp24439)

							tmp24441 := PrimEqual(Nil, tmp24440)

							var ifres24437 Obj

							if True == tmp24441 {
								ifres24437 = True

							} else {
								ifres24437 = False

							}

							ifres24436 = ifres24437

						} else {
							ifres24436 = False

						}

						var ifres24435 Obj

						if True == ifres24436 {
							ifres24435 = True

						} else {
							ifres24435 = False

						}

						ifres24434 = ifres24435

					} else {
						ifres24434 = False

					}

					var ifres24433 Obj

					if True == ifres24434 {
						ifres24433 = True

					} else {
						ifres24433 = False

					}

					ifres24432 = ifres24433

				} else {
					ifres24432 = False

				}

				var ifres24431 Obj

				if True == ifres24432 {
					ifres24431 = True

				} else {
					ifres24431 = False

				}

				ifres24430 = ifres24431

			} else {
				ifres24430 = False

			}

			if True == ifres24430 {
				tmp24424 := PrimTail(V4786)

				tmp24425 := PrimHead(tmp24424)

				tmp24426 := PrimTail(V4786)

				tmp24427 := PrimTail(tmp24426)

				tmp24428 := PrimHead(tmp24427)

				tmp24429 := Call(__e, PrimNS2Value(symshen_4decons), tmp24428)

				__e.Return(PrimCons(tmp24425, tmp24429))
				return

			} else {
				__e.Return(V4786)
				return
			}

		}

	}, 1)

	tmp24478 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decons, tmp24421)

	_ = tmp24478

	tmp24479 := MakeNative(func(__e *ControlFlow) {
		V4801 := __e.Get(1)
		_ = V4801
		V4802 := __e.Get(2)
		_ = V4802
		V4803 := __e.Get(3)
		_ = V4803
		tmp24509 := PrimIsPair(V4803)

		var ifres24484 Obj

		if True == tmp24509 {
			tmp24507 := PrimHead(V4803)

			tmp24508 := PrimEqual(symshen_4pair, tmp24507)

			var ifres24486 Obj

			if True == tmp24508 {
				tmp24505 := PrimTail(V4803)

				tmp24506 := PrimIsPair(tmp24505)

				var ifres24488 Obj

				if True == tmp24506 {
					tmp24502 := PrimTail(V4803)

					tmp24503 := PrimTail(tmp24502)

					tmp24504 := PrimIsPair(tmp24503)

					var ifres24490 Obj

					if True == tmp24504 {
						tmp24498 := PrimTail(V4803)

						tmp24499 := PrimTail(tmp24498)

						tmp24500 := PrimTail(tmp24499)

						tmp24501 := PrimEqual(Nil, tmp24500)

						var ifres24492 Obj

						if True == tmp24501 {
							tmp24494 := PrimTail(V4803)

							tmp24495 := PrimTail(tmp24494)

							tmp24496 := PrimHead(tmp24495)

							tmp24497 := PrimEqual(tmp24496, V4802)

							var ifres24493 Obj

							if True == tmp24497 {
								ifres24493 = True

							} else {
								ifres24493 = False

							}

							ifres24492 = ifres24493

						} else {
							ifres24492 = False

						}

						var ifres24491 Obj

						if True == ifres24492 {
							ifres24491 = True

						} else {
							ifres24491 = False

						}

						ifres24490 = ifres24491

					} else {
						ifres24490 = False

					}

					var ifres24489 Obj

					if True == ifres24490 {
						ifres24489 = True

					} else {
						ifres24489 = False

					}

					ifres24488 = ifres24489

				} else {
					ifres24488 = False

				}

				var ifres24487 Obj

				if True == ifres24488 {
					ifres24487 = True

				} else {
					ifres24487 = False

				}

				ifres24486 = ifres24487

			} else {
				ifres24486 = False

			}

			var ifres24485 Obj

			if True == ifres24486 {
				ifres24485 = True

			} else {
				ifres24485 = False

			}

			ifres24484 = ifres24485

		} else {
			ifres24484 = False

		}

		if True == ifres24484 {
			__e.Return(V4801)
			return
		} else {
			tmp24483 := PrimIsPair(V4803)

			if True == tmp24483 {
				tmp24482 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					__e.TailApply(PrimNS2Value(symshen_4insert_1runon), V4801, V4802, Z)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symmap), tmp24482, V4803)
				return

			} else {
				__e.Return(V4803)
				return
			}

		}

	}, 3)

	tmp24510 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1runon, tmp24479)

	_ = tmp24510

	tmp24511 := MakeNative(func(__e *ControlFlow) {
		V4809 := __e.Get(1)
		_ = V4809
		tmp24516 := Call(__e, PrimNS2Value(symelement_2), MakeString("."), V4809)

		tmp24517 := PrimNot(tmp24516)

		if True == tmp24517 {
			__e.Return(V4809)
			return
		} else {
			tmp24515 := PrimIsPair(V4809)

			if True == tmp24515 {
				tmp24514 := PrimTail(V4809)

				__e.TailApply(PrimNS2Value(symshen_4strip_1pathname), tmp24514)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4strip_1pathname)
				return
			}

		}

	}, 1)

	tmp24518 := Call(__e, PrimNS1Value(symns2_1set), symshen_4strip_1pathname, tmp24511)

	_ = tmp24518

	tmp24519 := MakeNative(func(__e *ControlFlow) {
		V4813 := __e.Get(1)
		_ = V4813
		V4814 := __e.Get(2)
		_ = V4814
		V4815 := __e.Get(3)
		_ = V4815
		tmp24549 := PrimIsPair(V4813)

		if True == tmp24549 {
			tmp24521 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp24522 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp24523 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp24524 := PrimHead(V4813)

						tmp24525 := Call(__e, PrimNS2Value(symconcat), symParse__, tmp24524)

						tmp24526 := PrimCons(symfail, Nil)

						tmp24527 := PrimHead(V4813)

						tmp24528 := Call(__e, PrimNS2Value(symconcat), symParse__, tmp24527)

						tmp24529 := PrimCons(tmp24528, Nil)

						tmp24530 := PrimCons(tmp24526, tmp24529)

						tmp24531 := PrimCons(sym_a, tmp24530)

						tmp24532 := PrimCons(tmp24531, Nil)

						tmp24533 := PrimCons(symnot, tmp24532)

						tmp24534 := PrimCons(Else, Nil)

						tmp24535 := PrimCons(Action, tmp24534)

						tmp24536 := PrimCons(tmp24533, tmp24535)

						tmp24537 := PrimCons(symif, tmp24536)

						tmp24538 := PrimCons(tmp24537, Nil)

						tmp24539 := PrimCons(Test, tmp24538)

						tmp24540 := PrimCons(tmp24525, tmp24539)

						__e.Return(PrimCons(symlet, tmp24540))
						return

					}, 1)

					tmp24541 := PrimCons(symfail, Nil)

					__e.TailApply(tmp24523, tmp24541)
					return

				}, 1)

				tmp24542 := PrimTail(V4813)

				tmp24543 := PrimHead(V4813)

				tmp24544 := Call(__e, PrimNS2Value(symconcat), symParse__, tmp24543)

				tmp24545 := Call(__e, PrimNS2Value(symshen_4syntax), tmp24542, tmp24544, V4815)

				__e.TailApply(tmp24522, tmp24545)
				return

			}, 1)

			tmp24546 := PrimHead(V4813)

			tmp24547 := PrimCons(V4814, Nil)

			tmp24548 := PrimCons(tmp24546, tmp24547)

			__e.TailApply(tmp24521, tmp24548)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4recursive__descent)
			return
		}

	}, 3)

	tmp24550 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__descent, tmp24519)

	_ = tmp24550

	tmp24551 := MakeNative(func(__e *ControlFlow) {
		V4819 := __e.Get(1)
		_ = V4819
		V4820 := __e.Get(2)
		_ = V4820
		V4821 := __e.Get(3)
		_ = V4821
		tmp24581 := PrimIsPair(V4819)

		if True == tmp24581 {
			tmp24553 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp24554 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp24555 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp24556 := PrimCons(Else, Nil)

						tmp24557 := PrimCons(Action, tmp24556)

						tmp24558 := PrimCons(Test, tmp24557)

						__e.Return(PrimCons(symif, tmp24558))
						return

					}, 1)

					tmp24559 := PrimCons(symfail, Nil)

					__e.TailApply(tmp24555, tmp24559)
					return

				}, 1)

				tmp24560 := PrimHead(V4819)

				tmp24561 := Call(__e, PrimNS2Value(symconcat), symParse__, tmp24560)

				tmp24562 := PrimCons(V4820, Nil)

				tmp24563 := PrimCons(symshen_4hdhd, tmp24562)

				tmp24564 := PrimTail(V4819)

				tmp24565 := PrimCons(V4820, Nil)

				tmp24566 := PrimCons(symshen_4tlhd, tmp24565)

				tmp24567 := PrimCons(V4820, Nil)

				tmp24568 := PrimCons(symshen_4hdtl, tmp24567)

				tmp24569 := PrimCons(tmp24568, Nil)

				tmp24570 := PrimCons(tmp24566, tmp24569)

				tmp24571 := PrimCons(symshen_4pair, tmp24570)

				tmp24572 := Call(__e, PrimNS2Value(symshen_4syntax), tmp24564, tmp24571, V4821)

				tmp24573 := PrimCons(tmp24572, Nil)

				tmp24574 := PrimCons(tmp24563, tmp24573)

				tmp24575 := PrimCons(tmp24561, tmp24574)

				tmp24576 := PrimCons(symlet, tmp24575)

				__e.TailApply(tmp24554, tmp24576)
				return

			}, 1)

			tmp24577 := PrimCons(V4820, Nil)

			tmp24578 := PrimCons(symhd, tmp24577)

			tmp24579 := PrimCons(tmp24578, Nil)

			tmp24580 := PrimCons(symcons_2, tmp24579)

			__e.TailApply(tmp24553, tmp24580)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4variable_1match)
			return
		}

	}, 3)

	tmp24582 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variable_1match, tmp24551)

	_ = tmp24582

	tmp24583 := MakeNative(func(__e *ControlFlow) {
		V4831 := __e.Get(1)
		_ = V4831
		tmp24587 := PrimIsPair(V4831)

		if True == tmp24587 {
			__e.Return(False)
			return
		} else {
			tmp24586 := PrimIsVariable(V4831)

			if True == tmp24586 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}

	}, 1)

	tmp24588 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terminal_2, tmp24583)

	_ = tmp24588

	tmp24589 := MakeNative(func(__e *ControlFlow) {
		V4837 := __e.Get(1)
		_ = V4837
		tmp24591 := PrimEqual(V4837, sym__)

		if True == tmp24591 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp24592 := Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream_2, tmp24589)

	_ = tmp24592

	tmp24593 := MakeNative(func(__e *ControlFlow) {
		V4841 := __e.Get(1)
		_ = V4841
		V4842 := __e.Get(2)
		_ = V4842
		V4843 := __e.Get(3)
		_ = V4843
		tmp24630 := PrimIsPair(V4841)

		if True == tmp24630 {
			tmp24595 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp24596 := MakeNative(func(__e *ControlFlow) {
					NewStr := __e.Get(1)
					_ = NewStr
					tmp24597 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						tmp24598 := MakeNative(func(__e *ControlFlow) {
							Else := __e.Get(1)
							_ = Else
							tmp24599 := PrimCons(Else, Nil)

							tmp24600 := PrimCons(Action, tmp24599)

							tmp24601 := PrimCons(Test, tmp24600)

							__e.Return(PrimCons(symif, tmp24601))
							return

						}, 1)

						tmp24602 := PrimCons(symfail, Nil)

						__e.TailApply(tmp24598, tmp24602)
						return

					}, 1)

					tmp24603 := PrimCons(V4842, Nil)

					tmp24604 := PrimCons(symshen_4tlhd, tmp24603)

					tmp24605 := PrimCons(V4842, Nil)

					tmp24606 := PrimCons(symshen_4hdtl, tmp24605)

					tmp24607 := PrimCons(tmp24606, Nil)

					tmp24608 := PrimCons(tmp24604, tmp24607)

					tmp24609 := PrimCons(symshen_4pair, tmp24608)

					tmp24610 := PrimTail(V4841)

					tmp24611 := Call(__e, PrimNS2Value(symshen_4syntax), tmp24610, NewStr, V4843)

					tmp24612 := PrimCons(tmp24611, Nil)

					tmp24613 := PrimCons(tmp24609, tmp24612)

					tmp24614 := PrimCons(NewStr, tmp24613)

					tmp24615 := PrimCons(symlet, tmp24614)

					__e.TailApply(tmp24597, tmp24615)
					return

				}, 1)

				tmp24616 := Call(__e, PrimNS2Value(symgensym), symNewStream)

				__e.TailApply(tmp24596, tmp24616)
				return

			}, 1)

			tmp24617 := PrimCons(V4842, Nil)

			tmp24618 := PrimCons(symhd, tmp24617)

			tmp24619 := PrimCons(tmp24618, Nil)

			tmp24620 := PrimCons(symcons_2, tmp24619)

			tmp24621 := PrimHead(V4841)

			tmp24622 := PrimCons(V4842, Nil)

			tmp24623 := PrimCons(symshen_4hdhd, tmp24622)

			tmp24624 := PrimCons(tmp24623, Nil)

			tmp24625 := PrimCons(tmp24621, tmp24624)

			tmp24626 := PrimCons(sym_a, tmp24625)

			tmp24627 := PrimCons(tmp24626, Nil)

			tmp24628 := PrimCons(tmp24620, tmp24627)

			tmp24629 := PrimCons(symand, tmp24628)

			__e.TailApply(tmp24595, tmp24629)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4check__stream)
			return
		}

	}, 3)

	tmp24631 := Call(__e, PrimNS1Value(symns2_1set), symshen_4check__stream, tmp24593)

	_ = tmp24631

	tmp24632 := MakeNative(func(__e *ControlFlow) {
		V4847 := __e.Get(1)
		_ = V4847
		V4848 := __e.Get(2)
		_ = V4848
		V4849 := __e.Get(3)
		_ = V4849
		tmp24654 := PrimIsPair(V4847)

		if True == tmp24654 {
			tmp24634 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp24635 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp24636 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp24637 := PrimCons(Else, Nil)

						tmp24638 := PrimCons(Action, tmp24637)

						tmp24639 := PrimCons(Test, tmp24638)

						__e.Return(PrimCons(symif, tmp24639))
						return

					}, 1)

					tmp24640 := PrimCons(symfail, Nil)

					__e.TailApply(tmp24636, tmp24640)
					return

				}, 1)

				tmp24641 := PrimTail(V4847)

				tmp24642 := PrimCons(V4848, Nil)

				tmp24643 := PrimCons(symshen_4tlhd, tmp24642)

				tmp24644 := PrimCons(V4848, Nil)

				tmp24645 := PrimCons(symshen_4hdtl, tmp24644)

				tmp24646 := PrimCons(tmp24645, Nil)

				tmp24647 := PrimCons(tmp24643, tmp24646)

				tmp24648 := PrimCons(symshen_4pair, tmp24647)

				tmp24649 := Call(__e, PrimNS2Value(symshen_4syntax), tmp24641, tmp24648, V4849)

				__e.TailApply(tmp24635, tmp24649)
				return

			}, 1)

			tmp24650 := PrimCons(V4848, Nil)

			tmp24651 := PrimCons(symhd, tmp24650)

			tmp24652 := PrimCons(tmp24651, Nil)

			tmp24653 := PrimCons(symcons_2, tmp24652)

			__e.TailApply(tmp24634, tmp24653)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4jump__stream)
			return
		}

	}, 3)

	tmp24655 := Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream, tmp24632)

	_ = tmp24655

	tmp24656 := MakeNative(func(__e *ControlFlow) {
		V4851 := __e.Get(1)
		_ = V4851
		tmp24667 := PrimEqual(Nil, V4851)

		if True == tmp24667 {
			__e.Return(Nil)
			return
		} else {
			tmp24666 := Call(__e, PrimNS2Value(symshen_4grammar__symbol_2), V4851)

			if True == tmp24666 {
				tmp24664 := Call(__e, PrimNS2Value(symconcat), symParse__, V4851)

				tmp24665 := PrimCons(tmp24664, Nil)

				__e.Return(PrimCons(symshen_4hdtl, tmp24665))
				return

			} else {
				tmp24663 := PrimIsVariable(V4851)

				if True == tmp24663 {
					__e.TailApply(PrimNS2Value(symconcat), symParse__, V4851)
					return
				} else {
					tmp24662 := PrimIsPair(V4851)

					if True == tmp24662 {
						tmp24661 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							__e.TailApply(PrimNS2Value(symshen_4semantics), Z)
							return
						}, 1)

						__e.TailApply(PrimNS2Value(symmap), tmp24661, V4851)
						return

					} else {
						__e.Return(V4851)
						return
					}

				}

			}

		}

	}, 1)

	tmp24668 := Call(__e, PrimNS1Value(symns2_1set), symshen_4semantics, tmp24656)

	_ = tmp24668

	tmp24669 := MakeNative(func(__e *ControlFlow) {
		V4854 := __e.Get(1)
		_ = V4854
		V4855 := __e.Get(2)
		_ = V4855
		tmp24670 := PrimCons(V4855, Nil)

		__e.Return(PrimCons(V4854, tmp24670))
		return

	}, 2)

	tmp24671 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pair, tmp24669)

	_ = tmp24671

	tmp24672 := MakeNative(func(__e *ControlFlow) {
		V4857 := __e.Get(1)
		_ = V4857
		tmp24673 := PrimTail(V4857)

		__e.Return(PrimHead(tmp24673))
		return

	}, 1)

	tmp24674 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hdtl, tmp24672)

	_ = tmp24674

	tmp24675 := MakeNative(func(__e *ControlFlow) {
		V4859 := __e.Get(1)
		_ = V4859
		tmp24676 := PrimHead(V4859)

		__e.Return(PrimHead(tmp24676))
		return

	}, 1)

	tmp24677 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hdhd, tmp24675)

	_ = tmp24677

	tmp24678 := MakeNative(func(__e *ControlFlow) {
		V4861 := __e.Get(1)
		_ = V4861
		tmp24679 := PrimHead(V4861)

		__e.Return(PrimTail(tmp24679))
		return

	}, 1)

	tmp24680 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tlhd, tmp24678)

	_ = tmp24680

	tmp24681 := MakeNative(func(__e *ControlFlow) {
		V4869 := __e.Get(1)
		_ = V4869
		tmp24693 := PrimIsPair(V4869)

		var ifres24684 Obj

		if True == tmp24693 {
			tmp24691 := PrimTail(V4869)

			tmp24692 := PrimIsPair(tmp24691)

			var ifres24686 Obj

			if True == tmp24692 {
				tmp24688 := PrimTail(V4869)

				tmp24689 := PrimTail(tmp24688)

				tmp24690 := PrimEqual(Nil, tmp24689)

				var ifres24687 Obj

				if True == tmp24690 {
					ifres24687 = True

				} else {
					ifres24687 = False

				}

				ifres24686 = ifres24687

			} else {
				ifres24686 = False

			}

			var ifres24685 Obj

			if True == ifres24686 {
				ifres24685 = True

			} else {
				ifres24685 = False

			}

			ifres24684 = ifres24685

		} else {
			ifres24684 = False

		}

		if True == ifres24684 {
			tmp24683 := PrimTail(V4869)

			__e.Return(PrimHead(tmp24683))
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp24694 := Call(__e, PrimNS1Value(symns2_1set), symshen_4snd_1or_1fail, tmp24681)

	_ = tmp24694

	tmp24695 := MakeNative(func(__e *ControlFlow) {
		__e.Return(symshen_4fail_b)
		return
	}, 0)

	tmp24696 := Call(__e, PrimNS1Value(symns2_1set), symfail, tmp24695)

	_ = tmp24696

	tmp24697 := MakeNative(func(__e *ControlFlow) {
		V4877 := __e.Get(1)
		_ = V4877
		tmp24710 := PrimIsPair(V4877)

		var ifres24701 Obj

		if True == tmp24710 {
			tmp24708 := PrimTail(V4877)

			tmp24709 := PrimIsPair(tmp24708)

			var ifres24703 Obj

			if True == tmp24709 {
				tmp24705 := PrimTail(V4877)

				tmp24706 := PrimTail(tmp24705)

				tmp24707 := PrimEqual(Nil, tmp24706)

				var ifres24704 Obj

				if True == tmp24707 {
					ifres24704 = True

				} else {
					ifres24704 = False

				}

				ifres24703 = ifres24704

			} else {
				ifres24703 = False

			}

			var ifres24702 Obj

			if True == ifres24703 {
				ifres24702 = True

			} else {
				ifres24702 = False

			}

			ifres24701 = ifres24702

		} else {
			ifres24701 = False

		}

		if True == ifres24701 {
			tmp24699 := PrimHead(V4877)

			tmp24700 := PrimCons(tmp24699, Nil)

			__e.Return(PrimCons(Nil, tmp24700))
			return

		} else {
			__e.TailApply(PrimNS2Value(symfail))
			return
		}

	}, 1)

	tmp24711 := Call(__e, PrimNS1Value(symns2_1set), sym_5_b_6, tmp24697)

	_ = tmp24711

	tmp24712 := MakeNative(func(__e *ControlFlow) {
		V4883 := __e.Get(1)
		_ = V4883
		tmp24725 := PrimIsPair(V4883)

		var ifres24716 Obj

		if True == tmp24725 {
			tmp24723 := PrimTail(V4883)

			tmp24724 := PrimIsPair(tmp24723)

			var ifres24718 Obj

			if True == tmp24724 {
				tmp24720 := PrimTail(V4883)

				tmp24721 := PrimTail(tmp24720)

				tmp24722 := PrimEqual(Nil, tmp24721)

				var ifres24719 Obj

				if True == tmp24722 {
					ifres24719 = True

				} else {
					ifres24719 = False

				}

				ifres24718 = ifres24719

			} else {
				ifres24718 = False

			}

			var ifres24717 Obj

			if True == ifres24718 {
				ifres24717 = True

			} else {
				ifres24717 = False

			}

			ifres24716 = ifres24717

		} else {
			ifres24716 = False

		}

		if True == ifres24716 {
			tmp24714 := PrimHead(V4883)

			tmp24715 := PrimCons(Nil, Nil)

			__e.Return(PrimCons(tmp24714, tmp24715))
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), sym_5e_6)
			return
		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), sym_5e_6, tmp24712)
	return

}, 0)
