package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp17114 := MakeNative(func(__e *ControlFlow) {
		tmp17115 := Call(__e, PrimNS2Value(symshen_4credits))

		_ = tmp17115

		__e.TailApply(PrimNS2Value(symshen_4loop))
		return

	}, 0)

	tmp17116 := Call(__e, PrimNS1Value(symns2_1set), symshen_4repl, tmp17114)

	_ = tmp17116

	tmp17117 := MakeNative(func(__e *ControlFlow) {
		tmp17118 := Call(__e, PrimNS2Value(symshen_4initialise__environment))

		_ = tmp17118

		tmp17119 := Call(__e, PrimNS2Value(symshen_4prompt))

		_ = tmp17119

		tmp17120 := MakeNative(func(__e *ControlFlow) {
			__e.TailApply(PrimNS2Value(symshen_4read_1evaluate_1print))
			return
		}, 0)

		tmp17121 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.TailApply(PrimNS2Value(symshen_4toplevel_1display_1exception), E)
			return
		}, 1)

		tmp17122 := Call(__e, PrimNS1Value(symtry_1catch), tmp17120, tmp17121)

		_ = tmp17122

		__e.TailApply(PrimNS2Value(symshen_4loop))
		return

	}, 0)

	tmp17123 := Call(__e, PrimNS1Value(symns2_1set), symshen_4loop, tmp17117)

	_ = tmp17123

	tmp17124 := MakeNative(func(__e *ControlFlow) {
		V3022 := __e.Get(1)
		_ = V3022
		tmp17125 := PrimErrorToString(V3022)

		tmp17126 := Call(__e, PrimNS2Value(symstoutput))

		__e.TailApply(PrimNS2Value(sympr), tmp17125, tmp17126)
		return

	}, 1)

	tmp17127 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel_1display_1exception, tmp17124)

	_ = tmp17127

	tmp17128 := MakeNative(func(__e *ControlFlow) {
		tmp17129 := Call(__e, PrimNS2Value(symstoutput))

		tmp17130 := Call(__e, PrimNS2Value(symshen_4prhush), MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n"), tmp17129)

		_ = tmp17130

		tmp17131 := PrimNS3Value(sym_dversion_d)

		tmp17132 := Call(__e, PrimNS2Value(symshen_4app), tmp17131, MakeString("\n"), symshen_4a)

		tmp17133 := PrimStringConcat(MakeString("www.shenlanguage.org, "), tmp17132)

		tmp17134 := Call(__e, PrimNS2Value(symstoutput))

		tmp17135 := Call(__e, PrimNS2Value(symshen_4prhush), tmp17133, tmp17134)

		_ = tmp17135

		tmp17136 := PrimNS3Value(sym_dlanguage_d)

		tmp17137 := PrimNS3Value(sym_dimplementation_d)

		tmp17138 := Call(__e, PrimNS2Value(symshen_4app), tmp17137, MakeString(""), symshen_4a)

		tmp17139 := PrimStringConcat(MakeString(", implementation: "), tmp17138)

		tmp17140 := Call(__e, PrimNS2Value(symshen_4app), tmp17136, tmp17139, symshen_4a)

		tmp17141 := PrimStringConcat(MakeString("running under "), tmp17140)

		tmp17142 := Call(__e, PrimNS2Value(symstoutput))

		tmp17143 := Call(__e, PrimNS2Value(symshen_4prhush), tmp17141, tmp17142)

		_ = tmp17143

		tmp17144 := PrimNS3Value(sym_dport_d)

		tmp17145 := PrimNS3Value(sym_dporters_d)

		tmp17146 := Call(__e, PrimNS2Value(symshen_4app), tmp17145, MakeString("\n"), symshen_4a)

		tmp17147 := PrimStringConcat(MakeString(" ported by "), tmp17146)

		tmp17148 := Call(__e, PrimNS2Value(symshen_4app), tmp17144, tmp17147, symshen_4a)

		tmp17149 := PrimStringConcat(MakeString("\nport "), tmp17148)

		tmp17150 := Call(__e, PrimNS2Value(symstoutput))

		__e.TailApply(PrimNS2Value(symshen_4prhush), tmp17149, tmp17150)
		return

	}, 0)

	tmp17151 := Call(__e, PrimNS1Value(symns2_1set), symshen_4credits, tmp17128)

	_ = tmp17151

	tmp17152 := MakeNative(func(__e *ControlFlow) {
		tmp17153 := PrimCons(MakeNumber(0), Nil)

		tmp17154 := PrimCons(symshen_4_dcatch_d, tmp17153)

		tmp17155 := PrimCons(MakeNumber(0), tmp17154)

		tmp17156 := PrimCons(symshen_4_dprocess_1counter_d, tmp17155)

		tmp17157 := PrimCons(MakeNumber(0), tmp17156)

		tmp17158 := PrimCons(symshen_4_dinfs_d, tmp17157)

		tmp17159 := PrimCons(MakeNumber(0), tmp17158)

		tmp17160 := PrimCons(symshen_4_dcall_d, tmp17159)

		__e.TailApply(PrimNS2Value(symshen_4multiple_1set), tmp17160)
		return

	}, 0)

	tmp17161 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__environment, tmp17152)

	_ = tmp17161

	tmp17162 := MakeNative(func(__e *ControlFlow) {
		V3024 := __e.Get(1)
		_ = V3024
		tmp17176 := PrimEqual(Nil, V3024)

		if True == tmp17176 {
			__e.Return(Nil)
			return
		} else {
			tmp17175 := PrimIsPair(V3024)

			var ifres17171 Obj

			if True == tmp17175 {
				tmp17173 := PrimTail(V3024)

				tmp17174 := PrimIsPair(tmp17173)

				var ifres17172 Obj

				if True == tmp17174 {
					ifres17172 = True

				} else {
					ifres17172 = False

				}

				ifres17171 = ifres17172

			} else {
				ifres17171 = False

			}

			if True == ifres17171 {
				tmp17165 := PrimHead(V3024)

				tmp17166 := PrimTail(V3024)

				tmp17167 := PrimHead(tmp17166)

				tmp17168 := PrimNS3Set(tmp17165, tmp17167)

				_ = tmp17168

				tmp17169 := PrimTail(V3024)

				tmp17170 := PrimTail(tmp17169)

				__e.TailApply(PrimNS2Value(symshen_4multiple_1set), tmp17170)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4multiple_1set)
				return
			}

		}

	}, 1)

	tmp17177 := Call(__e, PrimNS1Value(symns2_1set), symshen_4multiple_1set, tmp17162)

	_ = tmp17177

	tmp17178 := MakeNative(func(__e *ControlFlow) {
		V3026 := __e.Get(1)
		_ = V3026
		__e.TailApply(PrimNS2Value(symdeclare), V3026, symsymbol)
		return
	}, 1)

	tmp17179 := Call(__e, PrimNS1Value(symns2_1set), symdestroy, tmp17178)

	_ = tmp17179

	tmp17180 := MakeNative(func(__e *ControlFlow) {
		tmp17181 := MakeNative(func(__e *ControlFlow) {
			Lineread := __e.Get(1)
			_ = Lineread
			tmp17182 := MakeNative(func(__e *ControlFlow) {
				History := __e.Get(1)
				_ = History
				tmp17183 := MakeNative(func(__e *ControlFlow) {
					NewLineread := __e.Get(1)
					_ = NewLineread
					tmp17184 := MakeNative(func(__e *ControlFlow) {
						NewHistory := __e.Get(1)
						_ = NewHistory
						tmp17185 := MakeNative(func(__e *ControlFlow) {
							Parsed := __e.Get(1)
							_ = Parsed
							__e.TailApply(PrimNS2Value(symshen_4toplevel), Parsed)
							return
						}, 1)

						tmp17186 := Call(__e, PrimNS2Value(symfst), NewLineread)

						__e.TailApply(tmp17185, tmp17186)
						return

					}, 1)

					tmp17187 := Call(__e, PrimNS2Value(symshen_4update__history), NewLineread, History)

					__e.TailApply(tmp17184, tmp17187)
					return

				}, 1)

				tmp17188 := Call(__e, PrimNS2Value(symshen_4retrieve_1from_1history_1if_1needed), Lineread, History)

				__e.TailApply(tmp17183, tmp17188)
				return

			}, 1)

			tmp17189 := PrimNS3Value(symshen_4_dhistory_d)

			__e.TailApply(tmp17182, tmp17189)
			return

		}, 1)

		tmp17190 := Call(__e, PrimNS2Value(symshen_4toplineread))

		__e.TailApply(tmp17181, tmp17190)
		return

	}, 0)

	tmp17191 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1evaluate_1print, tmp17180)

	_ = tmp17191

	tmp17192 := MakeNative(func(__e *ControlFlow) {
		V3038 := __e.Get(1)
		_ = V3038
		V3039 := __e.Get(2)
		_ = V3039
		tmp17309 := Call(__e, PrimNS2Value(symtuple_2), V3038)

		var ifres17296 Obj

		if True == tmp17309 {
			tmp17307 := Call(__e, PrimNS2Value(symsnd), V3038)

			tmp17308 := PrimIsPair(tmp17307)

			var ifres17298 Obj

			if True == tmp17308 {
				tmp17300 := Call(__e, PrimNS2Value(symsnd), V3038)

				tmp17301 := PrimHead(tmp17300)

				tmp17302 := Call(__e, PrimNS2Value(symshen_4space))

				tmp17303 := Call(__e, PrimNS2Value(symshen_4newline))

				tmp17304 := PrimCons(tmp17303, Nil)

				tmp17305 := PrimCons(tmp17302, tmp17304)

				tmp17306 := Call(__e, PrimNS2Value(symelement_2), tmp17301, tmp17305)

				var ifres17299 Obj

				if True == tmp17306 {
					ifres17299 = True

				} else {
					ifres17299 = False

				}

				ifres17298 = ifres17299

			} else {
				ifres17298 = False

			}

			var ifres17297 Obj

			if True == ifres17298 {
				ifres17297 = True

			} else {
				ifres17297 = False

			}

			ifres17296 = ifres17297

		} else {
			ifres17296 = False

		}

		if True == ifres17296 {
			tmp17292 := Call(__e, PrimNS2Value(symfst), V3038)

			tmp17293 := Call(__e, PrimNS2Value(symsnd), V3038)

			tmp17294 := PrimTail(tmp17293)

			tmp17295 := Call(__e, PrimNS2Value(sym_8p), tmp17292, tmp17294)

			__e.TailApply(PrimNS2Value(symshen_4retrieve_1from_1history_1if_1needed), tmp17295, V3039)
			return

		} else {
			tmp17291 := Call(__e, PrimNS2Value(symtuple_2), V3038)

			var ifres17260 Obj

			if True == tmp17291 {
				tmp17289 := Call(__e, PrimNS2Value(symsnd), V3038)

				tmp17290 := PrimIsPair(tmp17289)

				var ifres17262 Obj

				if True == tmp17290 {
					tmp17286 := Call(__e, PrimNS2Value(symsnd), V3038)

					tmp17287 := PrimTail(tmp17286)

					tmp17288 := PrimIsPair(tmp17287)

					var ifres17264 Obj

					if True == tmp17288 {
						tmp17282 := Call(__e, PrimNS2Value(symsnd), V3038)

						tmp17283 := PrimTail(tmp17282)

						tmp17284 := PrimTail(tmp17283)

						tmp17285 := PrimEqual(Nil, tmp17284)

						var ifres17266 Obj

						if True == tmp17285 {
							tmp17281 := PrimIsPair(V3039)

							var ifres17268 Obj

							if True == tmp17281 {
								tmp17277 := Call(__e, PrimNS2Value(symsnd), V3038)

								tmp17278 := PrimHead(tmp17277)

								tmp17279 := Call(__e, PrimNS2Value(symshen_4exclamation))

								tmp17280 := PrimEqual(tmp17278, tmp17279)

								var ifres17270 Obj

								if True == tmp17280 {
									tmp17272 := Call(__e, PrimNS2Value(symsnd), V3038)

									tmp17273 := PrimTail(tmp17272)

									tmp17274 := PrimHead(tmp17273)

									tmp17275 := Call(__e, PrimNS2Value(symshen_4exclamation))

									tmp17276 := PrimEqual(tmp17274, tmp17275)

									var ifres17271 Obj

									if True == tmp17276 {
										ifres17271 = True

									} else {
										ifres17271 = False

									}

									ifres17270 = ifres17271

								} else {
									ifres17270 = False

								}

								var ifres17269 Obj

								if True == ifres17270 {
									ifres17269 = True

								} else {
									ifres17269 = False

								}

								ifres17268 = ifres17269

							} else {
								ifres17268 = False

							}

							var ifres17267 Obj

							if True == ifres17268 {
								ifres17267 = True

							} else {
								ifres17267 = False

							}

							ifres17266 = ifres17267

						} else {
							ifres17266 = False

						}

						var ifres17265 Obj

						if True == ifres17266 {
							ifres17265 = True

						} else {
							ifres17265 = False

						}

						ifres17264 = ifres17265

					} else {
						ifres17264 = False

					}

					var ifres17263 Obj

					if True == ifres17264 {
						ifres17263 = True

					} else {
						ifres17263 = False

					}

					ifres17262 = ifres17263

				} else {
					ifres17262 = False

				}

				var ifres17261 Obj

				if True == ifres17262 {
					ifres17261 = True

				} else {
					ifres17261 = False

				}

				ifres17260 = ifres17261

			} else {
				ifres17260 = False

			}

			if True == ifres17260 {
				tmp17256 := MakeNative(func(__e *ControlFlow) {
					PastPrint := __e.Get(1)
					_ = PastPrint
					__e.Return(PrimHead(V3039))
					return
				}, 1)

				tmp17257 := PrimHead(V3039)

				tmp17258 := Call(__e, PrimNS2Value(symsnd), tmp17257)

				tmp17259 := Call(__e, PrimNS2Value(symshen_4prbytes), tmp17258)

				__e.TailApply(tmp17256, tmp17259)
				return

			} else {
				tmp17255 := Call(__e, PrimNS2Value(symtuple_2), V3038)

				var ifres17245 Obj

				if True == tmp17255 {
					tmp17253 := Call(__e, PrimNS2Value(symsnd), V3038)

					tmp17254 := PrimIsPair(tmp17253)

					var ifres17247 Obj

					if True == tmp17254 {
						tmp17249 := Call(__e, PrimNS2Value(symsnd), V3038)

						tmp17250 := PrimHead(tmp17249)

						tmp17251 := Call(__e, PrimNS2Value(symshen_4exclamation))

						tmp17252 := PrimEqual(tmp17250, tmp17251)

						var ifres17248 Obj

						if True == tmp17252 {
							ifres17248 = True

						} else {
							ifres17248 = False

						}

						ifres17247 = ifres17248

					} else {
						ifres17247 = False

					}

					var ifres17246 Obj

					if True == ifres17247 {
						ifres17246 = True

					} else {
						ifres17246 = False

					}

					ifres17245 = ifres17246

				} else {
					ifres17245 = False

				}

				if True == ifres17245 {
					tmp17235 := MakeNative(func(__e *ControlFlow) {
						Key_2 := __e.Get(1)
						_ = Key_2
						tmp17236 := MakeNative(func(__e *ControlFlow) {
							Find := __e.Get(1)
							_ = Find
							tmp17237 := MakeNative(func(__e *ControlFlow) {
								PastPrint := __e.Get(1)
								_ = PastPrint
								__e.Return(Find)
								return
							}, 1)

							tmp17238 := Call(__e, PrimNS2Value(symsnd), Find)

							tmp17239 := Call(__e, PrimNS2Value(symshen_4prbytes), tmp17238)

							__e.TailApply(tmp17237, tmp17239)
							return

						}, 1)

						tmp17240 := Call(__e, PrimNS2Value(symshen_4find_1past_1inputs), Key_2, V3039)

						tmp17241 := Call(__e, PrimNS2Value(symhead), tmp17240)

						__e.TailApply(tmp17236, tmp17241)
						return

					}, 1)

					tmp17242 := Call(__e, PrimNS2Value(symsnd), V3038)

					tmp17243 := PrimTail(tmp17242)

					tmp17244 := Call(__e, PrimNS2Value(symshen_4make_1key), tmp17243, V3039)

					__e.TailApply(tmp17235, tmp17244)
					return

				} else {
					tmp17234 := Call(__e, PrimNS2Value(symtuple_2), V3038)

					var ifres17219 Obj

					if True == tmp17234 {
						tmp17232 := Call(__e, PrimNS2Value(symsnd), V3038)

						tmp17233 := PrimIsPair(tmp17232)

						var ifres17221 Obj

						if True == tmp17233 {
							tmp17229 := Call(__e, PrimNS2Value(symsnd), V3038)

							tmp17230 := PrimTail(tmp17229)

							tmp17231 := PrimEqual(Nil, tmp17230)

							var ifres17223 Obj

							if True == tmp17231 {
								tmp17225 := Call(__e, PrimNS2Value(symsnd), V3038)

								tmp17226 := PrimHead(tmp17225)

								tmp17227 := Call(__e, PrimNS2Value(symshen_4percent))

								tmp17228 := PrimEqual(tmp17226, tmp17227)

								var ifres17224 Obj

								if True == tmp17228 {
									ifres17224 = True

								} else {
									ifres17224 = False

								}

								ifres17223 = ifres17224

							} else {
								ifres17223 = False

							}

							var ifres17222 Obj

							if True == ifres17223 {
								ifres17222 = True

							} else {
								ifres17222 = False

							}

							ifres17221 = ifres17222

						} else {
							ifres17221 = False

						}

						var ifres17220 Obj

						if True == ifres17221 {
							ifres17220 = True

						} else {
							ifres17220 = False

						}

						ifres17219 = ifres17220

					} else {
						ifres17219 = False

					}

					if True == ifres17219 {
						tmp17216 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.Return(True)
							return
						}, 1)

						tmp17217 := Call(__e, PrimNS2Value(symreverse), V3039)

						tmp17218 := Call(__e, PrimNS2Value(symshen_4print_1past_1inputs), tmp17216, tmp17217, MakeNumber(0))

						_ = tmp17218

						__e.TailApply(PrimNS2Value(symabort))
						return

					} else {
						tmp17215 := Call(__e, PrimNS2Value(symtuple_2), V3038)

						var ifres17205 Obj

						if True == tmp17215 {
							tmp17213 := Call(__e, PrimNS2Value(symsnd), V3038)

							tmp17214 := PrimIsPair(tmp17213)

							var ifres17207 Obj

							if True == tmp17214 {
								tmp17209 := Call(__e, PrimNS2Value(symsnd), V3038)

								tmp17210 := PrimHead(tmp17209)

								tmp17211 := Call(__e, PrimNS2Value(symshen_4percent))

								tmp17212 := PrimEqual(tmp17210, tmp17211)

								var ifres17208 Obj

								if True == tmp17212 {
									ifres17208 = True

								} else {
									ifres17208 = False

								}

								ifres17207 = ifres17208

							} else {
								ifres17207 = False

							}

							var ifres17206 Obj

							if True == ifres17207 {
								ifres17206 = True

							} else {
								ifres17206 = False

							}

							ifres17205 = ifres17206

						} else {
							ifres17205 = False

						}

						if True == ifres17205 {
							tmp17198 := MakeNative(func(__e *ControlFlow) {
								Key_2 := __e.Get(1)
								_ = Key_2
								tmp17199 := MakeNative(func(__e *ControlFlow) {
									Pastprint := __e.Get(1)
									_ = Pastprint
									__e.TailApply(PrimNS2Value(symabort))
									return
								}, 1)

								tmp17200 := Call(__e, PrimNS2Value(symreverse), V3039)

								tmp17201 := Call(__e, PrimNS2Value(symshen_4print_1past_1inputs), Key_2, tmp17200, MakeNumber(0))

								__e.TailApply(tmp17199, tmp17201)
								return

							}, 1)

							tmp17202 := Call(__e, PrimNS2Value(symsnd), V3038)

							tmp17203 := PrimTail(tmp17202)

							tmp17204 := Call(__e, PrimNS2Value(symshen_4make_1key), tmp17203, V3039)

							__e.TailApply(tmp17198, tmp17204)
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

	tmp17310 := Call(__e, PrimNS1Value(symns2_1set), symshen_4retrieve_1from_1history_1if_1needed, tmp17192)

	_ = tmp17310

	tmp17311 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(37))
		return
	}, 0)

	tmp17312 := Call(__e, PrimNS1Value(symns2_1set), symshen_4percent, tmp17311)

	_ = tmp17312

	tmp17313 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(33))
		return
	}, 0)

	tmp17314 := Call(__e, PrimNS1Value(symns2_1set), symshen_4exclamation, tmp17313)

	_ = tmp17314

	tmp17315 := MakeNative(func(__e *ControlFlow) {
		V3041 := __e.Get(1)
		_ = V3041
		tmp17316 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			tmp17317 := PrimNumberToString(Byte)

			tmp17318 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), tmp17317, tmp17318)
			return

		}, 1)

		tmp17319 := Call(__e, PrimNS2Value(symshen_4for_1each), tmp17316, V3041)

		_ = tmp17319

		__e.TailApply(PrimNS2Value(symnl), MakeNumber(1))
		return

	}, 1)

	tmp17320 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prbytes, tmp17315)

	_ = tmp17320

	tmp17321 := MakeNative(func(__e *ControlFlow) {
		V3044 := __e.Get(1)
		_ = V3044
		V3045 := __e.Get(2)
		_ = V3045
		tmp17322 := PrimCons(V3044, V3045)

		__e.Return(PrimNS3Set(symshen_4_dhistory_d, tmp17322))
		return

	}, 2)

	tmp17323 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update__history, tmp17321)

	_ = tmp17323

	tmp17324 := MakeNative(func(__e *ControlFlow) {
		tmp17325 := Call(__e, PrimNS2Value(symstinput))

		tmp17326 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), tmp17325)

		__e.TailApply(PrimNS2Value(symshen_4toplineread__loop), tmp17326, Nil)
		return

	}, 0)

	tmp17327 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread, tmp17324)

	_ = tmp17327

	tmp17328 := MakeNative(func(__e *ControlFlow) {
		V3049 := __e.Get(1)
		_ = V3049
		V3050 := __e.Get(2)
		_ = V3050
		tmp17357 := Call(__e, PrimNS2Value(symshen_4hat))

		tmp17358 := PrimEqual(V3049, tmp17357)

		if True == tmp17358 {
			__e.Return(PrimSimpleError(MakeString("line read aborted")))
			return
		} else {
			tmp17352 := Call(__e, PrimNS2Value(symshen_4newline))

			tmp17353 := Call(__e, PrimNS2Value(symshen_4carriage_1return))

			tmp17354 := PrimCons(tmp17353, Nil)

			tmp17355 := PrimCons(tmp17352, tmp17354)

			tmp17356 := Call(__e, PrimNS2Value(symelement_2), V3049, tmp17355)

			if True == tmp17356 {
				tmp17337 := MakeNative(func(__e *ControlFlow) {
					Line := __e.Get(1)
					_ = Line
					tmp17338 := MakeNative(func(__e *ControlFlow) {
						It := __e.Get(1)
						_ = It
						tmp17347 := PrimEqual(Line, symshen_4nextline)

						var ifres17344 Obj

						if True == tmp17347 {
							ifres17344 = True

						} else {
							tmp17346 := Call(__e, PrimNS2Value(symempty_2), Line)

							var ifres17345 Obj

							if True == tmp17346 {
								ifres17345 = True

							} else {
								ifres17345 = False

							}

							ifres17344 = ifres17345

						}

						if True == ifres17344 {
							tmp17340 := Call(__e, PrimNS2Value(symstinput))

							tmp17341 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), tmp17340)

							tmp17342 := PrimCons(V3049, Nil)

							tmp17343 := Call(__e, PrimNS2Value(symappend), V3050, tmp17342)

							__e.TailApply(PrimNS2Value(symshen_4toplineread__loop), tmp17341, tmp17343)
							return

						} else {
							__e.TailApply(PrimNS2Value(sym_8p), Line, V3050)
							return
						}

					}, 1)

					tmp17348 := Call(__e, PrimNS2Value(symshen_4record_1it), V3050)

					__e.TailApply(tmp17338, tmp17348)
					return

				}, 1)

				tmp17349 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
					return
				}, 1)

				tmp17350 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4nextline)
					return
				}, 1)

				tmp17351 := Call(__e, PrimNS2Value(symcompile), tmp17349, V3050, tmp17350)

				__e.TailApply(tmp17337, tmp17351)
				return

			} else {
				tmp17331 := Call(__e, PrimNS2Value(symstinput))

				tmp17332 := Call(__e, PrimNS2Value(symshen_4read_1char_1code), tmp17331)

				tmp17336 := PrimEqual(V3049, MakeNumber(-1))

				var ifres17333 Obj

				if True == tmp17336 {
					ifres17333 = V3050

				} else {
					tmp17334 := PrimCons(V3049, Nil)

					tmp17335 := Call(__e, PrimNS2Value(symappend), V3050, tmp17334)

					ifres17333 = tmp17335

				}

				__e.TailApply(PrimNS2Value(symshen_4toplineread__loop), tmp17332, ifres17333)
				return

			}

		}

	}, 2)

	tmp17359 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread__loop, tmp17328)

	_ = tmp17359

	tmp17360 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(94))
		return
	}, 0)

	tmp17361 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hat, tmp17360)

	_ = tmp17361

	tmp17362 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(10))
		return
	}, 0)

	tmp17363 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newline, tmp17362)

	_ = tmp17363

	tmp17364 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(13))
		return
	}, 0)

	tmp17365 := Call(__e, PrimNS1Value(symns2_1set), symshen_4carriage_1return, tmp17364)

	_ = tmp17365

	tmp17366 := MakeNative(func(__e *ControlFlow) {
		V3056 := __e.Get(1)
		_ = V3056
		tmp17370 := PrimEqual(sym_7, V3056)

		if True == tmp17370 {
			__e.Return(PrimNS3Set(symshen_4_dtc_d, True))
			return
		} else {
			tmp17369 := PrimEqual(sym_1, V3056)

			if True == tmp17369 {
				__e.Return(PrimNS3Set(symshen_4_dtc_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
				return
			}

		}

	}, 1)

	tmp17371 := Call(__e, PrimNS1Value(symns2_1set), symtc, tmp17366)

	_ = tmp17371

	tmp17372 := MakeNative(func(__e *ControlFlow) {
		tmp17384 := PrimNS3Value(symshen_4_dtc_d)

		if True == tmp17384 {
			tmp17379 := PrimNS3Value(symshen_4_dhistory_d)

			tmp17380 := Call(__e, PrimNS2Value(symlength), tmp17379)

			tmp17381 := Call(__e, PrimNS2Value(symshen_4app), tmp17380, MakeString("+) "), symshen_4a)

			tmp17382 := PrimStringConcat(MakeString("\n\n("), tmp17381)

			tmp17383 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), tmp17382, tmp17383)
			return

		} else {
			tmp17374 := PrimNS3Value(symshen_4_dhistory_d)

			tmp17375 := Call(__e, PrimNS2Value(symlength), tmp17374)

			tmp17376 := Call(__e, PrimNS2Value(symshen_4app), tmp17375, MakeString("-) "), symshen_4a)

			tmp17377 := PrimStringConcat(MakeString("\n\n("), tmp17376)

			tmp17378 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(symshen_4prhush), tmp17377, tmp17378)
			return

		}

	}, 0)

	tmp17385 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prompt, tmp17372)

	_ = tmp17385

	tmp17386 := MakeNative(func(__e *ControlFlow) {
		V3058 := __e.Get(1)
		_ = V3058
		tmp17387 := PrimNS3Value(symshen_4_dtc_d)

		__e.TailApply(PrimNS2Value(symshen_4toplevel__evaluate), V3058, tmp17387)
		return

	}, 1)

	tmp17388 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel, tmp17386)

	_ = tmp17388

	tmp17389 := MakeNative(func(__e *ControlFlow) {
		V3061 := __e.Get(1)
		_ = V3061
		V3062 := __e.Get(2)
		_ = V3062
		tmp17390 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp17392 := Call(__e, PrimNS2Value(symempty_2), F)

			if True == tmp17392 {
				__e.Return(PrimSimpleError(MakeString("input not found\n")))
				return
			} else {
				__e.Return(F)
				return
			}

		}, 1)

		tmp17393 := Call(__e, PrimNS2Value(symshen_4find), V3061, V3062)

		__e.TailApply(tmp17390, tmp17393)
		return

	}, 2)

	tmp17394 := Call(__e, PrimNS1Value(symns2_1set), symshen_4find_1past_1inputs, tmp17389)

	_ = tmp17394

	tmp17395 := MakeNative(func(__e *ControlFlow) {
		V3065 := __e.Get(1)
		_ = V3065
		V3066 := __e.Get(2)
		_ = V3066
		tmp17396 := MakeNative(func(__e *ControlFlow) {
			Atom := __e.Get(1)
			_ = Atom
			tmp17403 := PrimIsInteger(Atom)

			if True == tmp17403 {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp17400 := PrimNumberAdd(Atom, MakeNumber(1))

					tmp17401 := Call(__e, PrimNS2Value(symreverse), V3066)

					tmp17402 := Call(__e, PrimNS2Value(symnth), tmp17400, tmp17401)

					__e.Return(PrimEqual(X, tmp17402))
					return

				}, 1))
				return
			} else {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp17398 := Call(__e, PrimNS2Value(symsnd), X)

					tmp17399 := Call(__e, PrimNS2Value(symshen_4trim_1gubbins), tmp17398)

					__e.TailApply(PrimNS2Value(symshen_4prefix_2), V3065, tmp17399)
					return

				}, 1))
				return
			}

		}, 1)

		tmp17404 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4_5st__input_6), X)
			return
		}, 1)

		tmp17405 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp17409 := PrimIsPair(E)

			if True == tmp17409 {
				tmp17407 := Call(__e, PrimNS2Value(symshen_4app), E, MakeString("\n"), symshen_4s)

				tmp17408 := PrimStringConcat(MakeString("parse error here: "), tmp17407)

				__e.Return(PrimSimpleError(tmp17408))
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("parse error\n")))
				return
			}

		}, 1)

		tmp17410 := Call(__e, PrimNS2Value(symcompile), tmp17404, V3065, tmp17405)

		tmp17411 := PrimHead(tmp17410)

		__e.TailApply(tmp17396, tmp17411)
		return

	}, 2)

	tmp17412 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1key, tmp17395)

	_ = tmp17412

	tmp17413 := MakeNative(func(__e *ControlFlow) {
		V3068 := __e.Get(1)
		_ = V3068
		tmp17453 := PrimIsPair(V3068)

		var ifres17448 Obj

		if True == tmp17453 {
			tmp17450 := PrimHead(V3068)

			tmp17451 := Call(__e, PrimNS2Value(symshen_4space))

			tmp17452 := PrimEqual(tmp17450, tmp17451)

			var ifres17449 Obj

			if True == tmp17452 {
				ifres17449 = True

			} else {
				ifres17449 = False

			}

			ifres17448 = ifres17449

		} else {
			ifres17448 = False

		}

		if True == ifres17448 {
			tmp17447 := PrimTail(V3068)

			__e.TailApply(PrimNS2Value(symshen_4trim_1gubbins), tmp17447)
			return

		} else {
			tmp17446 := PrimIsPair(V3068)

			var ifres17441 Obj

			if True == tmp17446 {
				tmp17443 := PrimHead(V3068)

				tmp17444 := Call(__e, PrimNS2Value(symshen_4newline))

				tmp17445 := PrimEqual(tmp17443, tmp17444)

				var ifres17442 Obj

				if True == tmp17445 {
					ifres17442 = True

				} else {
					ifres17442 = False

				}

				ifres17441 = ifres17442

			} else {
				ifres17441 = False

			}

			if True == ifres17441 {
				tmp17440 := PrimTail(V3068)

				__e.TailApply(PrimNS2Value(symshen_4trim_1gubbins), tmp17440)
				return

			} else {
				tmp17439 := PrimIsPair(V3068)

				var ifres17434 Obj

				if True == tmp17439 {
					tmp17436 := PrimHead(V3068)

					tmp17437 := Call(__e, PrimNS2Value(symshen_4carriage_1return))

					tmp17438 := PrimEqual(tmp17436, tmp17437)

					var ifres17435 Obj

					if True == tmp17438 {
						ifres17435 = True

					} else {
						ifres17435 = False

					}

					ifres17434 = ifres17435

				} else {
					ifres17434 = False

				}

				if True == ifres17434 {
					tmp17433 := PrimTail(V3068)

					__e.TailApply(PrimNS2Value(symshen_4trim_1gubbins), tmp17433)
					return

				} else {
					tmp17432 := PrimIsPair(V3068)

					var ifres17427 Obj

					if True == tmp17432 {
						tmp17429 := PrimHead(V3068)

						tmp17430 := Call(__e, PrimNS2Value(symshen_4tab))

						tmp17431 := PrimEqual(tmp17429, tmp17430)

						var ifres17428 Obj

						if True == tmp17431 {
							ifres17428 = True

						} else {
							ifres17428 = False

						}

						ifres17427 = ifres17428

					} else {
						ifres17427 = False

					}

					if True == ifres17427 {
						tmp17426 := PrimTail(V3068)

						__e.TailApply(PrimNS2Value(symshen_4trim_1gubbins), tmp17426)
						return

					} else {
						tmp17425 := PrimIsPair(V3068)

						var ifres17420 Obj

						if True == tmp17425 {
							tmp17422 := PrimHead(V3068)

							tmp17423 := Call(__e, PrimNS2Value(symshen_4left_1round))

							tmp17424 := PrimEqual(tmp17422, tmp17423)

							var ifres17421 Obj

							if True == tmp17424 {
								ifres17421 = True

							} else {
								ifres17421 = False

							}

							ifres17420 = ifres17421

						} else {
							ifres17420 = False

						}

						if True == ifres17420 {
							tmp17419 := PrimTail(V3068)

							__e.TailApply(PrimNS2Value(symshen_4trim_1gubbins), tmp17419)
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

	tmp17454 := Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1gubbins, tmp17413)

	_ = tmp17454

	tmp17455 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(32))
		return
	}, 0)

	tmp17456 := Call(__e, PrimNS1Value(symns2_1set), symshen_4space, tmp17455)

	_ = tmp17456

	tmp17457 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(9))
		return
	}, 0)

	tmp17458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tab, tmp17457)

	_ = tmp17458

	tmp17459 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(40))
		return
	}, 0)

	tmp17460 := Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1round, tmp17459)

	_ = tmp17460

	tmp17461 := MakeNative(func(__e *ControlFlow) {
		V3077 := __e.Get(1)
		_ = V3077
		V3078 := __e.Get(2)
		_ = V3078
		tmp17475 := PrimEqual(Nil, V3078)

		if True == tmp17475 {
			__e.Return(Nil)
			return
		} else {
			tmp17474 := PrimIsPair(V3078)

			var ifres17470 Obj

			if True == tmp17474 {
				tmp17472 := PrimHead(V3078)

				tmp17473 := Call(__e, V3077, tmp17472)

				var ifres17471 Obj

				if True == tmp17473 {
					ifres17471 = True

				} else {
					ifres17471 = False

				}

				ifres17470 = ifres17471

			} else {
				ifres17470 = False

			}

			if True == ifres17470 {
				tmp17467 := PrimHead(V3078)

				tmp17468 := PrimTail(V3078)

				tmp17469 := Call(__e, PrimNS2Value(symshen_4find), V3077, tmp17468)

				__e.Return(PrimCons(tmp17467, tmp17469))
				return

			} else {
				tmp17466 := PrimIsPair(V3078)

				if True == tmp17466 {
					tmp17465 := PrimTail(V3078)

					__e.TailApply(PrimNS2Value(symshen_4find), V3077, tmp17465)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4find)
					return
				}

			}

		}

	}, 2)

	tmp17476 := Call(__e, PrimNS1Value(symns2_1set), symshen_4find, tmp17461)

	_ = tmp17476

	tmp17477 := MakeNative(func(__e *ControlFlow) {
		V3092 := __e.Get(1)
		_ = V3092
		V3093 := __e.Get(2)
		_ = V3093
		tmp17491 := PrimEqual(Nil, V3092)

		if True == tmp17491 {
			__e.Return(True)
			return
		} else {
			tmp17490 := PrimIsPair(V3092)

			var ifres17482 Obj

			if True == tmp17490 {
				tmp17489 := PrimIsPair(V3093)

				var ifres17484 Obj

				if True == tmp17489 {
					tmp17486 := PrimHead(V3093)

					tmp17487 := PrimHead(V3092)

					tmp17488 := PrimEqual(tmp17486, tmp17487)

					var ifres17485 Obj

					if True == tmp17488 {
						ifres17485 = True

					} else {
						ifres17485 = False

					}

					ifres17484 = ifres17485

				} else {
					ifres17484 = False

				}

				var ifres17483 Obj

				if True == ifres17484 {
					ifres17483 = True

				} else {
					ifres17483 = False

				}

				ifres17482 = ifres17483

			} else {
				ifres17482 = False

			}

			if True == ifres17482 {
				tmp17480 := PrimTail(V3092)

				tmp17481 := PrimTail(V3093)

				__e.TailApply(PrimNS2Value(symshen_4prefix_2), tmp17480, tmp17481)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp17492 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prefix_2, tmp17477)

	_ = tmp17492

	tmp17493 := MakeNative(func(__e *ControlFlow) {
		V3105 := __e.Get(1)
		_ = V3105
		V3106 := __e.Get(2)
		_ = V3106
		V3107 := __e.Get(3)
		_ = V3107
		tmp17518 := PrimEqual(Nil, V3106)

		if True == tmp17518 {
			__e.Return(sym__)
			return
		} else {
			tmp17517 := PrimIsPair(V3106)

			var ifres17512 Obj

			if True == tmp17517 {
				tmp17514 := PrimHead(V3106)

				tmp17515 := Call(__e, V3105, tmp17514)

				tmp17516 := PrimNot(tmp17515)

				var ifres17513 Obj

				if True == tmp17516 {
					ifres17513 = True

				} else {
					ifres17513 = False

				}

				ifres17512 = ifres17513

			} else {
				ifres17512 = False

			}

			if True == ifres17512 {
				tmp17510 := PrimTail(V3106)

				tmp17511 := PrimNumberAdd(V3107, MakeNumber(1))

				__e.TailApply(PrimNS2Value(symshen_4print_1past_1inputs), V3105, tmp17510, tmp17511)
				return

			} else {
				tmp17509 := PrimIsPair(V3106)

				var ifres17505 Obj

				if True == tmp17509 {
					tmp17507 := PrimHead(V3106)

					tmp17508 := Call(__e, PrimNS2Value(symtuple_2), tmp17507)

					var ifres17506 Obj

					if True == tmp17508 {
						ifres17506 = True

					} else {
						ifres17506 = False

					}

					ifres17505 = ifres17506

				} else {
					ifres17505 = False

				}

				if True == ifres17505 {
					tmp17497 := Call(__e, PrimNS2Value(symshen_4app), V3107, MakeString(". "), symshen_4a)

					tmp17498 := Call(__e, PrimNS2Value(symstoutput))

					tmp17499 := Call(__e, PrimNS2Value(symshen_4prhush), tmp17497, tmp17498)

					_ = tmp17499

					tmp17500 := PrimHead(V3106)

					tmp17501 := Call(__e, PrimNS2Value(symsnd), tmp17500)

					tmp17502 := Call(__e, PrimNS2Value(symshen_4prbytes), tmp17501)

					_ = tmp17502

					tmp17503 := PrimTail(V3106)

					tmp17504 := PrimNumberAdd(V3107, MakeNumber(1))

					__e.TailApply(PrimNS2Value(symshen_4print_1past_1inputs), V3105, tmp17503, tmp17504)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4print_1past_1inputs)
					return
				}

			}

		}

	}, 3)

	tmp17519 := Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1past_1inputs, tmp17493)

	_ = tmp17519

	tmp17520 := MakeNative(func(__e *ControlFlow) {
		V3110 := __e.Get(1)
		_ = V3110
		V3111 := __e.Get(2)
		_ = V3111
		tmp17583 := PrimIsPair(V3110)

		var ifres17560 Obj

		if True == tmp17583 {
			tmp17581 := PrimTail(V3110)

			tmp17582 := PrimIsPair(tmp17581)

			var ifres17562 Obj

			if True == tmp17582 {
				tmp17578 := PrimTail(V3110)

				tmp17579 := PrimHead(tmp17578)

				tmp17580 := PrimEqual(sym_h, tmp17579)

				var ifres17564 Obj

				if True == tmp17580 {
					tmp17575 := PrimTail(V3110)

					tmp17576 := PrimTail(tmp17575)

					tmp17577 := PrimIsPair(tmp17576)

					var ifres17566 Obj

					if True == tmp17577 {
						tmp17571 := PrimTail(V3110)

						tmp17572 := PrimTail(tmp17571)

						tmp17573 := PrimTail(tmp17572)

						tmp17574 := PrimEqual(Nil, tmp17573)

						var ifres17568 Obj

						if True == tmp17574 {
							tmp17570 := PrimEqual(True, V3111)

							var ifres17569 Obj

							if True == tmp17570 {
								ifres17569 = True

							} else {
								ifres17569 = False

							}

							ifres17568 = ifres17569

						} else {
							ifres17568 = False

						}

						var ifres17567 Obj

						if True == ifres17568 {
							ifres17567 = True

						} else {
							ifres17567 = False

						}

						ifres17566 = ifres17567

					} else {
						ifres17566 = False

					}

					var ifres17565 Obj

					if True == ifres17566 {
						ifres17565 = True

					} else {
						ifres17565 = False

					}

					ifres17564 = ifres17565

				} else {
					ifres17564 = False

				}

				var ifres17563 Obj

				if True == ifres17564 {
					ifres17563 = True

				} else {
					ifres17563 = False

				}

				ifres17562 = ifres17563

			} else {
				ifres17562 = False

			}

			var ifres17561 Obj

			if True == ifres17562 {
				ifres17561 = True

			} else {
				ifres17561 = False

			}

			ifres17560 = ifres17561

		} else {
			ifres17560 = False

		}

		if True == ifres17560 {
			tmp17556 := PrimHead(V3110)

			tmp17557 := PrimTail(V3110)

			tmp17558 := PrimTail(tmp17557)

			tmp17559 := PrimHead(tmp17558)

			__e.TailApply(PrimNS2Value(symshen_4typecheck_1and_1evaluate), tmp17556, tmp17559)
			return

		} else {
			tmp17555 := PrimIsPair(V3110)

			var ifres17551 Obj

			if True == tmp17555 {
				tmp17553 := PrimTail(V3110)

				tmp17554 := PrimIsPair(tmp17553)

				var ifres17552 Obj

				if True == tmp17554 {
					ifres17552 = True

				} else {
					ifres17552 = False

				}

				ifres17551 = ifres17552

			} else {
				ifres17551 = False

			}

			if True == ifres17551 {
				tmp17546 := PrimHead(V3110)

				tmp17547 := PrimCons(tmp17546, Nil)

				tmp17548 := Call(__e, PrimNS2Value(symshen_4toplevel__evaluate), tmp17547, V3111)

				_ = tmp17548

				tmp17549 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

				_ = tmp17549

				tmp17550 := PrimTail(V3110)

				__e.TailApply(PrimNS2Value(symshen_4toplevel__evaluate), tmp17550, V3111)
				return

			} else {
				tmp17545 := PrimIsPair(V3110)

				var ifres17538 Obj

				if True == tmp17545 {
					tmp17543 := PrimTail(V3110)

					tmp17544 := PrimEqual(Nil, tmp17543)

					var ifres17540 Obj

					if True == tmp17544 {
						tmp17542 := PrimEqual(True, V3111)

						var ifres17541 Obj

						if True == tmp17542 {
							ifres17541 = True

						} else {
							ifres17541 = False

						}

						ifres17540 = ifres17541

					} else {
						ifres17540 = False

					}

					var ifres17539 Obj

					if True == ifres17540 {
						ifres17539 = True

					} else {
						ifres17539 = False

					}

					ifres17538 = ifres17539

				} else {
					ifres17538 = False

				}

				if True == ifres17538 {
					tmp17536 := PrimHead(V3110)

					tmp17537 := Call(__e, PrimNS2Value(symgensym), symA)

					__e.TailApply(PrimNS2Value(symshen_4typecheck_1and_1evaluate), tmp17536, tmp17537)
					return

				} else {
					tmp17535 := PrimIsPair(V3110)

					var ifres17528 Obj

					if True == tmp17535 {
						tmp17533 := PrimTail(V3110)

						tmp17534 := PrimEqual(Nil, tmp17533)

						var ifres17530 Obj

						if True == tmp17534 {
							tmp17532 := PrimEqual(False, V3111)

							var ifres17531 Obj

							if True == tmp17532 {
								ifres17531 = True

							} else {
								ifres17531 = False

							}

							ifres17530 = ifres17531

						} else {
							ifres17530 = False

						}

						var ifres17529 Obj

						if True == ifres17530 {
							ifres17529 = True

						} else {
							ifres17529 = False

						}

						ifres17528 = ifres17529

					} else {
						ifres17528 = False

					}

					if True == ifres17528 {
						tmp17525 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							__e.TailApply(PrimNS2Value(symprint), Eval)
							return
						}, 1)

						tmp17526 := PrimHead(V3110)

						tmp17527 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), tmp17526)

						__e.TailApply(tmp17525, tmp17527)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4toplevel__evaluate)
						return
					}

				}

			}

		}

	}, 2)

	tmp17584 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel__evaluate, tmp17520)

	_ = tmp17584

	tmp17585 := MakeNative(func(__e *ControlFlow) {
		V3114 := __e.Get(1)
		_ = V3114
		V3115 := __e.Get(2)
		_ = V3115
		tmp17586 := MakeNative(func(__e *ControlFlow) {
			Typecheck := __e.Get(1)
			_ = Typecheck
			tmp17596 := PrimEqual(Typecheck, False)

			if True == tmp17596 {
				__e.Return(PrimSimpleError(MakeString("type error\n")))
				return
			} else {
				tmp17588 := MakeNative(func(__e *ControlFlow) {
					Eval := __e.Get(1)
					_ = Eval
					tmp17589 := MakeNative(func(__e *ControlFlow) {
						Type := __e.Get(1)
						_ = Type
						tmp17590 := Call(__e, PrimNS2Value(symshen_4app), Type, MakeString(""), symshen_4r)

						tmp17591 := PrimStringConcat(MakeString(" : "), tmp17590)

						tmp17592 := Call(__e, PrimNS2Value(symshen_4app), Eval, tmp17591, symshen_4s)

						tmp17593 := Call(__e, PrimNS2Value(symstoutput))

						__e.TailApply(PrimNS2Value(symshen_4prhush), tmp17592, tmp17593)
						return

					}, 1)

					tmp17594 := Call(__e, PrimNS2Value(symshen_4pretty_1type), Typecheck)

					__e.TailApply(tmp17589, tmp17594)
					return

				}, 1)

				tmp17595 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), V3114)

				__e.TailApply(tmp17588, tmp17595)
				return

			}

		}, 1)

		tmp17597 := Call(__e, PrimNS2Value(symshen_4typecheck), V3114, V3115)

		__e.TailApply(tmp17586, tmp17597)
		return

	}, 2)

	tmp17598 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1evaluate, tmp17585)

	_ = tmp17598

	tmp17599 := MakeNative(func(__e *ControlFlow) {
		V3117 := __e.Get(1)
		_ = V3117
		tmp17600 := PrimNS3Value(symshen_4_dalphabet_d)

		tmp17601 := Call(__e, PrimNS2Value(symshen_4extract_1pvars), V3117)

		__e.TailApply(PrimNS2Value(symshen_4mult__subst), tmp17600, tmp17601, V3117)
		return

	}, 1)

	tmp17602 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pretty_1type, tmp17599)

	_ = tmp17602

	tmp17603 := MakeNative(func(__e *ControlFlow) {
		V3123 := __e.Get(1)
		_ = V3123
		tmp17611 := Call(__e, PrimNS2Value(symshen_4pvar_2), V3123)

		if True == tmp17611 {
			__e.Return(PrimCons(V3123, Nil))
			return
		} else {
			tmp17610 := PrimIsPair(V3123)

			if True == tmp17610 {
				tmp17606 := PrimHead(V3123)

				tmp17607 := Call(__e, PrimNS2Value(symshen_4extract_1pvars), tmp17606)

				tmp17608 := PrimTail(V3123)

				tmp17609 := Call(__e, PrimNS2Value(symshen_4extract_1pvars), tmp17608)

				__e.TailApply(PrimNS2Value(symunion), tmp17607, tmp17609)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp17612 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract_1pvars, tmp17603)

	_ = tmp17612

	tmp17613 := MakeNative(func(__e *ControlFlow) {
		V3131 := __e.Get(1)
		_ = V3131
		V3132 := __e.Get(2)
		_ = V3132
		V3133 := __e.Get(3)
		_ = V3133
		tmp17627 := PrimEqual(Nil, V3131)

		if True == tmp17627 {
			__e.Return(V3133)
			return
		} else {
			tmp17626 := PrimEqual(Nil, V3132)

			if True == tmp17626 {
				__e.Return(V3133)
				return
			} else {
				tmp17625 := PrimIsPair(V3131)

				var ifres17622 Obj

				if True == tmp17625 {
					tmp17624 := PrimIsPair(V3132)

					var ifres17623 Obj

					if True == tmp17624 {
						ifres17623 = True

					} else {
						ifres17623 = False

					}

					ifres17622 = ifres17623

				} else {
					ifres17622 = False

				}

				if True == ifres17622 {
					tmp17617 := PrimTail(V3131)

					tmp17618 := PrimTail(V3132)

					tmp17619 := PrimHead(V3131)

					tmp17620 := PrimHead(V3132)

					tmp17621 := Call(__e, PrimNS2Value(symsubst), tmp17619, tmp17620, V3133)

					__e.TailApply(PrimNS2Value(symshen_4mult__subst), tmp17617, tmp17618, tmp17621)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4mult__subst)
					return
				}

			}

		}

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4mult__subst, tmp17613)
	return

}, 0)
