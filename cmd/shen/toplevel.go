package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen97 := MakeNative(func(__e *ControlFlow) {
		gen95 := Call(__e, PrimNS1Value(symns2_1value), symshen_4credits)

		Call(__e, gen95)

		gen96 := Call(__e, PrimNS1Value(symns2_1value), symshen_4loop)

		__e.TailApply(gen96)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4repl, gen97)

	gen105 := MakeNative(func(__e *ControlFlow) {
		gen98 := Call(__e, PrimNS1Value(symns2_1value), symshen_4initialise__environment)

		Call(__e, gen98)

		gen99 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prompt)

		Call(__e, gen99)

		gen101 := MakeNative(func(__e *ControlFlow) {
			gen100 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1evaluate_1print)

			__e.TailApply(gen100)

			return

		}, 0)

		gen103 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel_1display_1exception)

			__e.TailApply(gen102, E)

			return

		}, 1)

		Call(__e, PrimNS1Value(symtry_1catch), gen101, gen103)

		gen104 := Call(__e, PrimNS1Value(symns2_1value), symshen_4loop)

		__e.TailApply(gen104)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4loop, gen105)

	gen111 := MakeNative(func(__e *ControlFlow) {
		V3022 := __e.Get(1)
		_ = V3022
		gen106 := Call(__e, PrimNS1Value(symns2_1value), sympr)

		gen107 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

		gen108 := Call(__e, gen107, V3022)

		gen109 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen110 := Call(__e, gen109)

		__e.TailApply(gen106, gen108, gen110)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel_1display_1exception, gen111)

	gen154 := MakeNative(func(__e *ControlFlow) {
		gen112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		gen113 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen114 := Call(__e, gen113)

		Call(__e, gen112, MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n"), gen114)

		gen115 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		gen116 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen118 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen119 := Call(__e, gen118, sym_dversion_d)

		gen120 := Call(__e, gen117, gen119, MakeString("\n"), symshen_4a)

		gen121 := Call(__e, gen116, MakeString("www.shenlanguage.org, "), gen120)

		gen122 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen123 := Call(__e, gen122)

		Call(__e, gen115, gen121, gen123)

		gen124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		gen125 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen126 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen127 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen128 := Call(__e, gen127, sym_dlanguage_d)

		gen129 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen130 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen131 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen132 := Call(__e, gen131, sym_dimplementation_d)

		gen133 := Call(__e, gen130, gen132, MakeString(""), symshen_4a)

		gen134 := Call(__e, gen129, MakeString(", implementation: "), gen133)

		gen135 := Call(__e, gen126, gen128, gen134, symshen_4a)

		gen136 := Call(__e, gen125, MakeString("running under "), gen135)

		gen137 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen138 := Call(__e, gen137)

		Call(__e, gen124, gen136, gen138)

		gen139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		gen140 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen142 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen143 := Call(__e, gen142, sym_dport_d)

		gen144 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen145 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen146 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen147 := Call(__e, gen146, sym_dporters_d)

		gen148 := Call(__e, gen145, gen147, MakeString("\n"), symshen_4a)

		gen149 := Call(__e, gen144, MakeString(" ported by "), gen148)

		gen150 := Call(__e, gen141, gen143, gen149, symshen_4a)

		gen151 := Call(__e, gen140, MakeString("\nport "), gen150)

		gen152 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen153 := Call(__e, gen152)

		__e.TailApply(gen139, gen151, gen153)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4credits, gen154)

	gen172 := MakeNative(func(__e *ControlFlow) {
		gen155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiple_1set)

		gen156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen158 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen162 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen163 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen164 := Call(__e, gen163, MakeNumber(0), Nil)

		gen165 := Call(__e, gen162, symshen_4_dcatch_d, gen164)

		gen166 := Call(__e, gen161, MakeNumber(0), gen165)

		gen167 := Call(__e, gen160, symshen_4_dprocess_1counter_d, gen166)

		gen168 := Call(__e, gen159, MakeNumber(0), gen167)

		gen169 := Call(__e, gen158, symshen_4_dinfs_d, gen168)

		gen170 := Call(__e, gen157, MakeNumber(0), gen169)

		gen171 := Call(__e, gen156, symshen_4_dcall_d, gen170)

		__e.TailApply(gen155, gen171)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__environment, gen172)

	gen196 := MakeNative(func(__e *ControlFlow) {
		V3024 := __e.Get(1)
		_ = V3024
		gen194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen195 := Call(__e, gen194, Nil, V3024)

		if True == gen195 {
			__e.Return(Nil)
			return
		} else {
			gen191 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen192 := Call(__e, gen191, V3024)

			var gen193 Obj

			if True == gen192 {
				gen187 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen188 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen189 := Call(__e, gen188, V3024)

				gen190 := Call(__e, gen187, gen189)

				if True == gen190 {
					gen193 = True
				} else {
					gen193 = False
				}

			} else {
				gen193 = False
			}

			if True == gen193 {
				gen175 := Call(__e, PrimNS1Value(symns2_1value), symset)

				gen176 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen177 := Call(__e, gen176, V3024)

				gen178 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen179 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen180 := Call(__e, gen179, V3024)

				gen181 := Call(__e, gen178, gen180)

				Call(__e, gen175, gen177, gen181)

				gen182 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiple_1set)

				gen183 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen184 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen185 := Call(__e, gen184, V3024)

				gen186 := Call(__e, gen183, gen185)

				__e.TailApply(gen182, gen186)

				return

			} else {
				if True == True {
					gen174 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen174, symshen_4multiple_1set)

					return

				} else {
					gen173 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen173, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4multiple_1set, gen196)

	gen198 := MakeNative(func(__e *ControlFlow) {
		V3026 := __e.Get(1)
		_ = V3026
		gen197 := Call(__e, PrimNS1Value(symns2_1value), symdeclare)

		__e.TailApply(gen197, V3026, symsymbol)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symdestroy, gen198)

	gen215 := MakeNative(func(__e *ControlFlow) {
		gen212 := MakeNative(func(__e *ControlFlow) {
			Lineread := __e.Get(1)
			_ = Lineread
			gen209 := MakeNative(func(__e *ControlFlow) {
				History := __e.Get(1)
				_ = History
				gen206 := MakeNative(func(__e *ControlFlow) {
					NewLineread := __e.Get(1)
					_ = NewLineread
					gen203 := MakeNative(func(__e *ControlFlow) {
						NewHistory := __e.Get(1)
						_ = NewHistory
						gen200 := MakeNative(func(__e *ControlFlow) {
							Parsed := __e.Get(1)
							_ = Parsed
							gen199 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel)

							__e.TailApply(gen199, Parsed)

							return

						}, 1)

						gen201 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						gen202 := Call(__e, gen201, NewLineread)

						__e.TailApply(gen200, gen202)

						return

					}, 1)

					gen204 := Call(__e, PrimNS1Value(symns2_1value), symshen_4update__history)

					gen205 := Call(__e, gen204, NewLineread, History)

					__e.TailApply(gen203, gen205)

					return

				}, 1)

				gen207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4retrieve_1from_1history_1if_1needed)

				gen208 := Call(__e, gen207, Lineread, History)

				__e.TailApply(gen206, gen208)

				return

			}, 1)

			gen210 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen211 := Call(__e, gen210, symshen_4_dhistory_d)

			__e.TailApply(gen209, gen211)

			return

		}, 1)

		gen213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread)

		gen214 := Call(__e, gen213)

		__e.TailApply(gen212, gen214)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1evaluate_1print, gen215)

	gen392 := MakeNative(func(__e *ControlFlow) {
		V3038 := __e.Get(1)
		_ = V3038
		V3039 := __e.Get(2)
		_ = V3039
		gen389 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		gen390 := Call(__e, gen389, V3038)

		var gen391 Obj

		if True == gen390 {
			gen384 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen385 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			gen386 := Call(__e, gen385, V3038)

			gen387 := Call(__e, gen384, gen386)

			var gen388 Obj

			if True == gen387 {
				gen370 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen371 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen372 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				gen373 := Call(__e, gen372, V3038)

				gen374 := Call(__e, gen371, gen373)

				gen375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen376 := Call(__e, PrimNS1Value(symns2_1value), symshen_4space)

				gen377 := Call(__e, gen376)

				gen378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				gen380 := Call(__e, gen379)

				gen381 := Call(__e, gen378, gen380, Nil)

				gen382 := Call(__e, gen375, gen377, gen381)

				gen383 := Call(__e, gen370, gen374, gen382)

				if True == gen383 {
					gen388 = True
				} else {
					gen388 = False
				}

			} else {
				gen388 = False
			}

			if True == gen388 {
				gen391 = True
			} else {
				gen391 = False
			}

		} else {
			gen391 = False
		}

		if True == gen391 {
			gen361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4retrieve_1from_1history_1if_1needed)

			gen362 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

			gen363 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			gen364 := Call(__e, gen363, V3038)

			gen365 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen366 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			gen367 := Call(__e, gen366, V3038)

			gen368 := Call(__e, gen365, gen367)

			gen369 := Call(__e, gen362, gen364, gen368)

			__e.TailApply(gen361, gen369, V3039)

			return

		} else {
			gen358 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

			gen359 := Call(__e, gen358, V3038)

			var gen360 Obj

			if True == gen359 {
				gen353 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen354 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				gen355 := Call(__e, gen354, V3038)

				gen356 := Call(__e, gen353, gen355)

				var gen357 Obj

				if True == gen356 {
					gen346 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen347 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen348 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen349 := Call(__e, gen348, V3038)

					gen350 := Call(__e, gen347, gen349)

					gen351 := Call(__e, gen346, gen350)

					var gen352 Obj

					if True == gen351 {
						gen337 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen339 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen340 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						gen341 := Call(__e, gen340, V3038)

						gen342 := Call(__e, gen339, gen341)

						gen343 := Call(__e, gen338, gen342)

						gen344 := Call(__e, gen337, Nil, gen343)

						var gen345 Obj

						if True == gen344 {
							gen334 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen335 := Call(__e, gen334, V3039)

							var gen336 Obj

							if True == gen335 {
								gen325 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen326 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen327 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								gen328 := Call(__e, gen327, V3038)

								gen329 := Call(__e, gen326, gen328)

								gen330 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

								gen331 := Call(__e, gen330)

								gen332 := Call(__e, gen325, gen329, gen331)

								var gen333 Obj

								if True == gen332 {
									gen315 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen316 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen317 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen318 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

									gen319 := Call(__e, gen318, V3038)

									gen320 := Call(__e, gen317, gen319)

									gen321 := Call(__e, gen316, gen320)

									gen322 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

									gen323 := Call(__e, gen322)

									gen324 := Call(__e, gen315, gen321, gen323)

									if True == gen324 {
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
								gen345 = True
							} else {
								gen345 = False
							}

						} else {
							gen345 = False
						}

						if True == gen345 {
							gen352 = True
						} else {
							gen352 = False
						}

					} else {
						gen352 = False
					}

					if True == gen352 {
						gen357 = True
					} else {
						gen357 = False
					}

				} else {
					gen357 = False
				}

				if True == gen357 {
					gen360 = True
				} else {
					gen360 = False
				}

			} else {
				gen360 = False
			}

			if True == gen360 {
				gen308 := MakeNative(func(__e *ControlFlow) {
					PastPrint := __e.Get(1)
					_ = PastPrint
					gen307 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					__e.TailApply(gen307, V3039)

					return

				}, 1)

				gen309 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

				gen310 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				gen311 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen312 := Call(__e, gen311, V3039)

				gen313 := Call(__e, gen310, gen312)

				gen314 := Call(__e, gen309, gen313)

				__e.TailApply(gen308, gen314)

				return

			} else {
				gen304 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

				gen305 := Call(__e, gen304, V3038)

				var gen306 Obj

				if True == gen305 {
					gen299 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen300 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen301 := Call(__e, gen300, V3038)

					gen302 := Call(__e, gen299, gen301)

					var gen303 Obj

					if True == gen302 {
						gen291 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen292 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen293 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						gen294 := Call(__e, gen293, V3038)

						gen295 := Call(__e, gen292, gen294)

						gen296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

						gen297 := Call(__e, gen296)

						gen298 := Call(__e, gen291, gen295, gen297)

						if True == gen298 {
							gen303 = True
						} else {
							gen303 = False
						}

					} else {
						gen303 = False
					}

					if True == gen303 {
						gen306 = True
					} else {
						gen306 = False
					}

				} else {
					gen306 = False
				}

				if True == gen306 {
					gen284 := MakeNative(func(__e *ControlFlow) {
						Key_2 := __e.Get(1)
						_ = Key_2
						gen279 := MakeNative(func(__e *ControlFlow) {
							Find := __e.Get(1)
							_ = Find
							gen274 := MakeNative(func(__e *ControlFlow) {
								PastPrint := __e.Get(1)
								_ = PastPrint
								__e.Return(Find)
								return
							}, 1)

							gen275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

							gen276 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							gen277 := Call(__e, gen276, Find)

							gen278 := Call(__e, gen275, gen277)

							__e.TailApply(gen274, gen278)

							return

						}, 1)

						gen280 := Call(__e, PrimNS1Value(symns2_1value), symhead)

						gen281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find_1past_1inputs)

						gen282 := Call(__e, gen281, Key_2, V3039)

						gen283 := Call(__e, gen280, gen282)

						__e.TailApply(gen279, gen283)

						return

					}, 1)

					gen285 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make_1key)

					gen286 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen287 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen288 := Call(__e, gen287, V3038)

					gen289 := Call(__e, gen286, gen288)

					gen290 := Call(__e, gen285, gen289, V3039)

					__e.TailApply(gen284, gen290)

					return

				} else {
					gen271 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					gen272 := Call(__e, gen271, V3038)

					var gen273 Obj

					if True == gen272 {
						gen266 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen267 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						gen268 := Call(__e, gen267, V3038)

						gen269 := Call(__e, gen266, gen268)

						var gen270 Obj

						if True == gen269 {
							gen259 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen260 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen261 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							gen262 := Call(__e, gen261, V3038)

							gen263 := Call(__e, gen260, gen262)

							gen264 := Call(__e, gen259, Nil, gen263)

							var gen265 Obj

							if True == gen264 {
								gen251 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen252 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen253 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								gen254 := Call(__e, gen253, V3038)

								gen255 := Call(__e, gen252, gen254)

								gen256 := Call(__e, PrimNS1Value(symns2_1value), symshen_4percent)

								gen257 := Call(__e, gen256)

								gen258 := Call(__e, gen251, gen255, gen257)

								if True == gen258 {
									gen265 = True
								} else {
									gen265 = False
								}

							} else {
								gen265 = False
							}

							if True == gen265 {
								gen270 = True
							} else {
								gen270 = False
							}

						} else {
							gen270 = False
						}

						if True == gen270 {
							gen273 = True
						} else {
							gen273 = False
						}

					} else {
						gen273 = False
					}

					if True == gen273 {
						gen246 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

						gen247 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.Return(True)
							return
						}, 1)

						gen248 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						gen249 := Call(__e, gen248, V3039)

						Call(__e, gen246, gen247, gen249, MakeNumber(0))

						gen250 := Call(__e, PrimNS1Value(symns2_1value), symabort)

						__e.TailApply(gen250)

						return

					} else {
						gen243 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

						gen244 := Call(__e, gen243, V3038)

						var gen245 Obj

						if True == gen244 {
							gen238 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen239 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							gen240 := Call(__e, gen239, V3038)

							gen241 := Call(__e, gen238, gen240)

							var gen242 Obj

							if True == gen241 {
								gen230 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								gen231 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen232 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								gen233 := Call(__e, gen232, V3038)

								gen234 := Call(__e, gen231, gen233)

								gen235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4percent)

								gen236 := Call(__e, gen235)

								gen237 := Call(__e, gen230, gen234, gen236)

								if True == gen237 {
									gen242 = True
								} else {
									gen242 = False
								}

							} else {
								gen242 = False
							}

							if True == gen242 {
								gen245 = True
							} else {
								gen245 = False
							}

						} else {
							gen245 = False
						}

						if True == gen245 {
							gen223 := MakeNative(func(__e *ControlFlow) {
								Key_2 := __e.Get(1)
								_ = Key_2
								gen218 := MakeNative(func(__e *ControlFlow) {
									Pastprint := __e.Get(1)
									_ = Pastprint
									gen217 := Call(__e, PrimNS1Value(symns2_1value), symabort)

									__e.TailApply(gen217)

									return

								}, 1)

								gen219 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

								gen220 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

								gen221 := Call(__e, gen220, V3039)

								gen222 := Call(__e, gen219, Key_2, gen221, MakeNumber(0))

								__e.TailApply(gen218, gen222)

								return

							}, 1)

							gen224 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make_1key)

							gen225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen226 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							gen227 := Call(__e, gen226, V3038)

							gen228 := Call(__e, gen225, gen227)

							gen229 := Call(__e, gen224, gen228, V3039)

							__e.TailApply(gen223, gen229)

							return

						} else {
							if True == True {
								__e.Return(V3038)
								return
							} else {
								gen216 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								__e.TailApply(gen216, MakeString("no cond match"))

								return

							}
						}

					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4retrieve_1from_1history_1if_1needed, gen392)

	gen393 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(37))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4percent, gen393)

	gen394 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(33))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4exclamation, gen394)

	gen403 := MakeNative(func(__e *ControlFlow) {
		V3041 := __e.Get(1)
		_ = V3041
		gen395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

		gen401 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			gen396 := Call(__e, PrimNS1Value(symns2_1value), sympr)

			gen397 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			gen398 := Call(__e, gen397, Byte)

			gen399 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen400 := Call(__e, gen399)

			__e.TailApply(gen396, gen398, gen400)

			return

		}, 1)

		Call(__e, gen395, gen401, V3041)

		gen402 := Call(__e, PrimNS1Value(symns2_1value), symnl)

		__e.TailApply(gen402, MakeNumber(1))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prbytes, gen403)

	gen407 := MakeNative(func(__e *ControlFlow) {
		V3044 := __e.Get(1)
		_ = V3044
		V3045 := __e.Get(2)
		_ = V3045
		gen404 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen405 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen406 := Call(__e, gen405, V3044, V3045)

		__e.TailApply(gen404, symshen_4_dhistory_d, gen406)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4update__history, gen407)

	gen413 := MakeNative(func(__e *ControlFlow) {
		gen408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

		gen409 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		gen410 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

		gen411 := Call(__e, gen410)

		gen412 := Call(__e, gen409, gen411)

		__e.TailApply(gen408, gen412, Nil)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread, gen413)

	gen465 := MakeNative(func(__e *ControlFlow) {
		V3049 := __e.Get(1)
		_ = V3049
		V3050 := __e.Get(2)
		_ = V3050
		gen461 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hat)

		gen463 := Call(__e, gen462)

		gen464 := Call(__e, gen461, V3049, gen463)

		if True == gen464 {
			gen460 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(gen460, MakeString("line read aborted"))

			return

		} else {
			gen450 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

			gen451 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

			gen453 := Call(__e, gen452)

			gen454 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen455 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

			gen456 := Call(__e, gen455)

			gen457 := Call(__e, gen454, gen456, Nil)

			gen458 := Call(__e, gen451, gen453, gen457)

			gen459 := Call(__e, gen450, V3049, gen458)

			if True == gen459 {
				gen444 := MakeNative(func(__e *ControlFlow) {
					Line := __e.Get(1)
					_ = Line
					gen441 := MakeNative(func(__e *ControlFlow) {
						It := __e.Get(1)
						_ = It
						gen438 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen439 := Call(__e, gen438, Line, symshen_4nextline)

						var gen440 Obj

						if True == gen439 {
							gen440 = True
						} else {
							gen436 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

							gen437 := Call(__e, gen436, Line)

							if True == gen437 {
								gen440 = True
							} else {
								gen440 = False
							}

						}

						if True == gen440 {
							gen427 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

							gen428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

							gen429 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

							gen430 := Call(__e, gen429)

							gen431 := Call(__e, gen428, gen430)

							gen432 := Call(__e, PrimNS1Value(symns2_1value), symappend)

							gen433 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							gen434 := Call(__e, gen433, V3049, Nil)

							gen435 := Call(__e, gen432, V3050, gen434)

							__e.TailApply(gen427, gen431, gen435)

							return

						} else {
							gen426 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

							__e.TailApply(gen426, Line, V3050)

							return

						}

					}, 1)

					gen442 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

					gen443 := Call(__e, gen442, V3050)

					__e.TailApply(gen441, gen443)

					return

				}, 1)

				gen445 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

				gen447 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen446 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

					__e.TailApply(gen446, X)

					return

				}, 1)

				gen448 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4nextline)
					return
				}, 1)

				gen449 := Call(__e, gen445, gen447, V3050, gen448)

				__e.TailApply(gen444, gen449)

				return

			} else {
				if True == True {
					gen415 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

					gen416 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

					gen417 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

					gen418 := Call(__e, gen417)

					gen419 := Call(__e, gen416, gen418)

					gen423 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen424 := Call(__e, gen423, V3049, MakeNumber(-1))

					var gen425 Obj

					if True == gen424 {
						gen425 = V3050
					} else {
						gen420 := Call(__e, PrimNS1Value(symns2_1value), symappend)

						gen421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen422 := Call(__e, gen421, V3049, Nil)

						gen425 = Call(__e, gen420, V3050, gen422)

					}

					__e.TailApply(gen415, gen419, gen425)

					return

				} else {
					gen414 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen414, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread__loop, gen465)

	gen466 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(94))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4hat, gen466)

	gen467 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(10))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4newline, gen467)

	gen468 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(13))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4carriage_1return, gen468)

	gen477 := MakeNative(func(__e *ControlFlow) {
		V3056 := __e.Get(1)
		_ = V3056
		gen475 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen476 := Call(__e, gen475, sym_7, V3056)

		if True == gen476 {
			gen474 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(gen474, symshen_4_dtc_d, True)

			return

		} else {
			gen472 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen473 := Call(__e, gen472, sym_1, V3056)

			if True == gen473 {
				gen471 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(gen471, symshen_4_dtc_d, False)

				return

			} else {
				if True == True {
					gen470 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen470, MakeString("tc expects a + or -"))

					return

				} else {
					gen469 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen469, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symtc, gen477)

	gen502 := MakeNative(func(__e *ControlFlow) {
		gen500 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen501 := Call(__e, gen500, symshen_4_dtc_d)

		if True == gen501 {
			gen489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen490 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen492 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			gen493 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen494 := Call(__e, gen493, symshen_4_dhistory_d)

			gen495 := Call(__e, gen492, gen494)

			gen496 := Call(__e, gen491, gen495, MakeString("+) "), symshen_4a)

			gen497 := Call(__e, gen490, MakeString("\n\n("), gen496)

			gen498 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen499 := Call(__e, gen498)

			__e.TailApply(gen489, gen497, gen499)

			return

		} else {
			gen478 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen479 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen481 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			gen482 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen483 := Call(__e, gen482, symshen_4_dhistory_d)

			gen484 := Call(__e, gen481, gen483)

			gen485 := Call(__e, gen480, gen484, MakeString("-) "), symshen_4a)

			gen486 := Call(__e, gen479, MakeString("\n\n("), gen485)

			gen487 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen488 := Call(__e, gen487)

			__e.TailApply(gen478, gen486, gen488)

			return

		}

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prompt, gen502)

	gen506 := MakeNative(func(__e *ControlFlow) {
		V3058 := __e.Get(1)
		_ = V3058
		gen503 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

		gen504 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen505 := Call(__e, gen504, symshen_4_dtc_d)

		__e.TailApply(gen503, V3058, gen505)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel, gen506)

	gen513 := MakeNative(func(__e *ControlFlow) {
		V3061 := __e.Get(1)
		_ = V3061
		V3062 := __e.Get(2)
		_ = V3062
		gen510 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			gen508 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

			gen509 := Call(__e, gen508, F)

			if True == gen509 {
				gen507 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen507, MakeString("input not found\n"))

				return

			} else {
				__e.Return(F)
				return
			}

		}, 1)

		gen511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

		gen512 := Call(__e, gen511, V3061, V3062)

		__e.TailApply(gen510, gen512)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4find_1past_1inputs, gen513)

	gen544 := MakeNative(func(__e *ControlFlow) {
		V3065 := __e.Get(1)
		_ = V3065
		V3066 := __e.Get(2)
		_ = V3066
		gen528 := MakeNative(func(__e *ControlFlow) {
			Atom := __e.Get(1)
			_ = Atom
			gen526 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

			gen527 := Call(__e, gen526, Atom)

			if True == gen527 {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen519 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen520 := Call(__e, PrimNS1Value(symns2_1value), symnth)

					gen521 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					gen522 := Call(__e, gen521, Atom, MakeNumber(1))

					gen523 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					gen524 := Call(__e, gen523, V3066)

					gen525 := Call(__e, gen520, gen522, gen524)

					__e.TailApply(gen519, X, gen525)

					return

				}, 1))
				return
			} else {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					gen514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

					gen515 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

					gen516 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen517 := Call(__e, gen516, X)

					gen518 := Call(__e, gen515, gen517)

					__e.TailApply(gen514, V3065, gen518)

					return

				}, 1))
				return
			}

		}, 1)

		gen529 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen530 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

		gen532 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen531 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

			__e.TailApply(gen531, X)

			return

		}, 1)

		gen541 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen539 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen540 := Call(__e, gen539, E)

			if True == gen540 {
				gen534 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				gen535 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen536 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen537 := Call(__e, gen536, E, MakeString("\n"), symshen_4s)

				gen538 := Call(__e, gen535, MakeString("parse error here: "), gen537)

				__e.TailApply(gen534, gen538)

				return

			} else {
				gen533 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen533, MakeString("parse error\n"))

				return

			}

		}, 1)

		gen542 := Call(__e, gen530, gen532, V3065, gen541)

		gen543 := Call(__e, gen529, gen542)

		__e.TailApply(gen528, gen543)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1key, gen544)

	gen606 := MakeNative(func(__e *ControlFlow) {
		V3068 := __e.Get(1)
		_ = V3068
		gen603 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen604 := Call(__e, gen603, V3068)

		var gen605 Obj

		if True == gen604 {
			gen597 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen598 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen599 := Call(__e, gen598, V3068)

			gen600 := Call(__e, PrimNS1Value(symns2_1value), symshen_4space)

			gen601 := Call(__e, gen600)

			gen602 := Call(__e, gen597, gen599, gen601)

			if True == gen602 {
				gen605 = True
			} else {
				gen605 = False
			}

		} else {
			gen605 = False
		}

		if True == gen605 {
			gen594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

			gen595 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen596 := Call(__e, gen595, V3068)

			__e.TailApply(gen594, gen596)

			return

		} else {
			gen591 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen592 := Call(__e, gen591, V3068)

			var gen593 Obj

			if True == gen592 {
				gen585 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen586 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen587 := Call(__e, gen586, V3068)

				gen588 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				gen589 := Call(__e, gen588)

				gen590 := Call(__e, gen585, gen587, gen589)

				if True == gen590 {
					gen593 = True
				} else {
					gen593 = False
				}

			} else {
				gen593 = False
			}

			if True == gen593 {
				gen582 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

				gen583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen584 := Call(__e, gen583, V3068)

				__e.TailApply(gen582, gen584)

				return

			} else {
				gen579 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen580 := Call(__e, gen579, V3068)

				var gen581 Obj

				if True == gen580 {
					gen573 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen575 := Call(__e, gen574, V3068)

					gen576 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

					gen577 := Call(__e, gen576)

					gen578 := Call(__e, gen573, gen575, gen577)

					if True == gen578 {
						gen581 = True
					} else {
						gen581 = False
					}

				} else {
					gen581 = False
				}

				if True == gen581 {
					gen570 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

					gen571 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen572 := Call(__e, gen571, V3068)

					__e.TailApply(gen570, gen572)

					return

				} else {
					gen567 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen568 := Call(__e, gen567, V3068)

					var gen569 Obj

					if True == gen568 {
						gen561 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen562 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen563 := Call(__e, gen562, V3068)

						gen564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tab)

						gen565 := Call(__e, gen564)

						gen566 := Call(__e, gen561, gen563, gen565)

						if True == gen566 {
							gen569 = True
						} else {
							gen569 = False
						}

					} else {
						gen569 = False
					}

					if True == gen569 {
						gen558 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

						gen559 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen560 := Call(__e, gen559, V3068)

						__e.TailApply(gen558, gen560)

						return

					} else {
						gen555 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen556 := Call(__e, gen555, V3068)

						var gen557 Obj

						if True == gen556 {
							gen549 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen550 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen551 := Call(__e, gen550, V3068)

							gen552 := Call(__e, PrimNS1Value(symns2_1value), symshen_4left_1round)

							gen553 := Call(__e, gen552)

							gen554 := Call(__e, gen549, gen551, gen553)

							if True == gen554 {
								gen557 = True
							} else {
								gen557 = False
							}

						} else {
							gen557 = False
						}

						if True == gen557 {
							gen546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

							gen547 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen548 := Call(__e, gen547, V3068)

							__e.TailApply(gen546, gen548)

							return

						} else {
							if True == True {
								__e.Return(V3068)
								return
							} else {
								gen545 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

								__e.TailApply(gen545, MakeString("no cond match"))

								return

							}
						}

					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1gubbins, gen606)

	gen607 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(32))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4space, gen607)

	gen608 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(9))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4tab, gen608)

	gen609 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(40))
		return
	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1round, gen609)

	gen632 := MakeNative(func(__e *ControlFlow) {
		V3077 := __e.Get(1)
		_ = V3077
		V3078 := __e.Get(2)
		_ = V3078
		gen630 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen631 := Call(__e, gen630, Nil, V3078)

		if True == gen631 {
			__e.Return(Nil)
			return
		} else {
			gen627 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen628 := Call(__e, gen627, V3078)

			var gen629 Obj

			if True == gen628 {
				gen624 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen625 := Call(__e, gen624, V3078)

				gen626 := Call(__e, V3077, gen625)

				if True == gen626 {
					gen629 = True
				} else {
					gen629 = False
				}

			} else {
				gen629 = False
			}

			if True == gen629 {
				gen617 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen618 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen619 := Call(__e, gen618, V3078)

				gen620 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

				gen621 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen622 := Call(__e, gen621, V3078)

				gen623 := Call(__e, gen620, V3077, gen622)

				__e.TailApply(gen617, gen619, gen623)

				return

			} else {
				gen615 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen616 := Call(__e, gen615, V3078)

				if True == gen616 {
					gen612 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

					gen613 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen614 := Call(__e, gen613, V3078)

					__e.TailApply(gen612, V3077, gen614)

					return

				} else {
					if True == True {
						gen611 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen611, symshen_4find)

						return

					} else {
						gen610 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen610, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4find, gen632)

	gen653 := MakeNative(func(__e *ControlFlow) {
		V3092 := __e.Get(1)
		_ = V3092
		V3093 := __e.Get(2)
		_ = V3093
		gen651 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen652 := Call(__e, gen651, Nil, V3092)

		if True == gen652 {
			__e.Return(True)
			return
		} else {
			gen648 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen649 := Call(__e, gen648, V3092)

			var gen650 Obj

			if True == gen649 {
				gen645 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen646 := Call(__e, gen645, V3093)

				var gen647 Obj

				if True == gen646 {
					gen639 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen640 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen641 := Call(__e, gen640, V3093)

					gen642 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen643 := Call(__e, gen642, V3092)

					gen644 := Call(__e, gen639, gen641, gen643)

					if True == gen644 {
						gen647 = True
					} else {
						gen647 = False
					}

				} else {
					gen647 = False
				}

				if True == gen647 {
					gen650 = True
				} else {
					gen650 = False
				}

			} else {
				gen650 = False
			}

			if True == gen650 {
				gen634 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

				gen635 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen636 := Call(__e, gen635, V3092)

				gen637 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen638 := Call(__e, gen637, V3093)

				__e.TailApply(gen634, gen636, gen638)

				return

			} else {
				if True == True {
					__e.Return(False)
					return
				} else {
					gen633 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen633, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4prefix_2, gen653)

	gen693 := MakeNative(func(__e *ControlFlow) {
		V3105 := __e.Get(1)
		_ = V3105
		V3106 := __e.Get(2)
		_ = V3106
		V3107 := __e.Get(3)
		_ = V3107
		gen691 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen692 := Call(__e, gen691, Nil, V3106)

		if True == gen692 {
			__e.Return(sym__)
			return
		} else {
			gen688 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen689 := Call(__e, gen688, V3106)

			var gen690 Obj

			if True == gen689 {
				gen683 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				gen684 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen685 := Call(__e, gen684, V3106)

				gen686 := Call(__e, V3105, gen685)

				gen687 := Call(__e, gen683, gen686)

				if True == gen687 {
					gen690 = True
				} else {
					gen690 = False
				}

			} else {
				gen690 = False
			}

			if True == gen690 {
				gen678 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

				gen679 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen680 := Call(__e, gen679, V3106)

				gen681 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen682 := Call(__e, gen681, V3107, MakeNumber(1))

				__e.TailApply(gen678, V3105, gen680, gen682)

				return

			} else {
				gen675 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen676 := Call(__e, gen675, V3106)

				var gen677 Obj

				if True == gen676 {
					gen671 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					gen672 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen673 := Call(__e, gen672, V3106)

					gen674 := Call(__e, gen671, gen673)

					if True == gen674 {
						gen677 = True
					} else {
						gen677 = False
					}

				} else {
					gen677 = False
				}

				if True == gen677 {
					gen656 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					gen657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen658 := Call(__e, gen657, V3107, MakeString(". "), symshen_4a)

					gen659 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					gen660 := Call(__e, gen659)

					Call(__e, gen656, gen658, gen660)

					gen661 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

					gen662 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen663 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen664 := Call(__e, gen663, V3106)

					gen665 := Call(__e, gen662, gen664)

					Call(__e, gen661, gen665)

					gen666 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

					gen667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen668 := Call(__e, gen667, V3106)

					gen669 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					gen670 := Call(__e, gen669, V3107, MakeNumber(1))

					__e.TailApply(gen666, V3105, gen668, gen670)

					return

				} else {
					if True == True {
						gen655 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen655, symshen_4print_1past_1inputs)

						return

					} else {
						gen654 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen654, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1past_1inputs, gen693)

	gen785 := MakeNative(func(__e *ControlFlow) {
		V3110 := __e.Get(1)
		_ = V3110
		V3111 := __e.Get(2)
		_ = V3111
		gen782 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen783 := Call(__e, gen782, V3110)

		var gen784 Obj

		if True == gen783 {
			gen777 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen778 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen779 := Call(__e, gen778, V3110)

			gen780 := Call(__e, gen777, gen779)

			var gen781 Obj

			if True == gen780 {
				gen770 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen771 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen772 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen773 := Call(__e, gen772, V3110)

				gen774 := Call(__e, gen771, gen773)

				gen775 := Call(__e, gen770, sym_h, gen774)

				var gen776 Obj

				if True == gen775 {
					gen763 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen764 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen765 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen766 := Call(__e, gen765, V3110)

					gen767 := Call(__e, gen764, gen766)

					gen768 := Call(__e, gen763, gen767)

					var gen769 Obj

					if True == gen768 {
						gen754 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen755 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen756 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen757 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen758 := Call(__e, gen757, V3110)

						gen759 := Call(__e, gen756, gen758)

						gen760 := Call(__e, gen755, gen759)

						gen761 := Call(__e, gen754, Nil, gen760)

						var gen762 Obj

						if True == gen761 {
							gen752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen753 := Call(__e, gen752, True, V3111)

							if True == gen753 {
								gen762 = True
							} else {
								gen762 = False
							}

						} else {
							gen762 = False
						}

						if True == gen762 {
							gen769 = True
						} else {
							gen769 = False
						}

					} else {
						gen769 = False
					}

					if True == gen769 {
						gen776 = True
					} else {
						gen776 = False
					}

				} else {
					gen776 = False
				}

				if True == gen776 {
					gen781 = True
				} else {
					gen781 = False
				}

			} else {
				gen781 = False
			}

			if True == gen781 {
				gen784 = True
			} else {
				gen784 = False
			}

		} else {
			gen784 = False
		}

		if True == gen784 {
			gen743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

			gen744 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen745 := Call(__e, gen744, V3110)

			gen746 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen748 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen749 := Call(__e, gen748, V3110)

			gen750 := Call(__e, gen747, gen749)

			gen751 := Call(__e, gen746, gen750)

			__e.TailApply(gen743, gen745, gen751)

			return

		} else {
			gen740 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen741 := Call(__e, gen740, V3110)

			var gen742 Obj

			if True == gen741 {
				gen736 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen737 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen738 := Call(__e, gen737, V3110)

				gen739 := Call(__e, gen736, gen738)

				if True == gen739 {
					gen742 = True
				} else {
					gen742 = False
				}

			} else {
				gen742 = False
			}

			if True == gen742 {
				gen727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

				gen728 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen729 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen730 := Call(__e, gen729, V3110)

				gen731 := Call(__e, gen728, gen730, Nil)

				Call(__e, gen727, gen731, V3111)

				gen732 := Call(__e, PrimNS1Value(symns2_1value), symnl)

				Call(__e, gen732, MakeNumber(1))

				gen733 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

				gen734 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen735 := Call(__e, gen734, V3110)

				__e.TailApply(gen733, gen735, V3111)

				return

			} else {
				gen724 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen725 := Call(__e, gen724, V3110)

				var gen726 Obj

				if True == gen725 {
					gen719 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen720 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen721 := Call(__e, gen720, V3110)

					gen722 := Call(__e, gen719, Nil, gen721)

					var gen723 Obj

					if True == gen722 {
						gen717 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen718 := Call(__e, gen717, True, V3111)

						if True == gen718 {
							gen723 = True
						} else {
							gen723 = False
						}

					} else {
						gen723 = False
					}

					if True == gen723 {
						gen726 = True
					} else {
						gen726 = False
					}

				} else {
					gen726 = False
				}

				if True == gen726 {
					gen712 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

					gen713 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen714 := Call(__e, gen713, V3110)

					gen715 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

					gen716 := Call(__e, gen715, symA)

					__e.TailApply(gen712, gen714, gen716)

					return

				} else {
					gen709 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen710 := Call(__e, gen709, V3110)

					var gen711 Obj

					if True == gen710 {
						gen704 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen705 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen706 := Call(__e, gen705, V3110)

						gen707 := Call(__e, gen704, Nil, gen706)

						var gen708 Obj

						if True == gen707 {
							gen702 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen703 := Call(__e, gen702, False, V3111)

							if True == gen703 {
								gen708 = True
							} else {
								gen708 = False
							}

						} else {
							gen708 = False
						}

						if True == gen708 {
							gen711 = True
						} else {
							gen711 = False
						}

					} else {
						gen711 = False
					}

					if True == gen711 {
						gen697 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							gen696 := Call(__e, PrimNS1Value(symns2_1value), symprint)

							__e.TailApply(gen696, Eval)

							return

						}, 1)

						gen698 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

						gen699 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen700 := Call(__e, gen699, V3110)

						gen701 := Call(__e, gen698, gen700)

						__e.TailApply(gen697, gen701)

						return

					} else {
						if True == True {
							gen695 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen695, symshen_4toplevel__evaluate)

							return

						} else {
							gen694 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen694, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel__evaluate, gen785)

	gen807 := MakeNative(func(__e *ControlFlow) {
		V3114 := __e.Get(1)
		_ = V3114
		V3115 := __e.Get(2)
		_ = V3115
		gen804 := MakeNative(func(__e *ControlFlow) {
			Typecheck := __e.Get(1)
			_ = Typecheck
			gen802 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen803 := Call(__e, gen802, Typecheck, False)

			if True == gen803 {
				gen801 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen801, MakeString("type error\n"))

				return

			} else {
				gen798 := MakeNative(func(__e *ControlFlow) {
					Eval := __e.Get(1)
					_ = Eval
					gen795 := MakeNative(func(__e *ControlFlow) {
						Type := __e.Get(1)
						_ = Type
						gen786 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

						gen787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						gen788 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						gen789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						gen790 := Call(__e, gen789, Type, MakeString(""), symshen_4r)

						gen791 := Call(__e, gen788, MakeString(" : "), gen790)

						gen792 := Call(__e, gen787, Eval, gen791, symshen_4s)

						gen793 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

						gen794 := Call(__e, gen793)

						__e.TailApply(gen786, gen792, gen794)

						return

					}, 1)

					gen796 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pretty_1type)

					gen797 := Call(__e, gen796, Typecheck)

					__e.TailApply(gen795, gen797)

					return

				}, 1)

				gen799 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				gen800 := Call(__e, gen799, V3114)

				__e.TailApply(gen798, gen800)

				return

			}

		}, 1)

		gen805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

		gen806 := Call(__e, gen805, V3114, V3115)

		__e.TailApply(gen804, gen806)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1evaluate, gen807)

	gen813 := MakeNative(func(__e *ControlFlow) {
		V3117 := __e.Get(1)
		_ = V3117
		gen808 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mult__subst)

		gen809 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen810 := Call(__e, gen809, symshen_4_dalphabet_d)

		gen811 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

		gen812 := Call(__e, gen811, V3117)

		__e.TailApply(gen808, gen810, gen812, V3117)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4pretty_1type, gen813)

	gen829 := MakeNative(func(__e *ControlFlow) {
		V3123 := __e.Get(1)
		_ = V3123
		gen827 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

		gen828 := Call(__e, gen827, V3123)

		if True == gen828 {
			gen826 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(gen826, V3123, Nil)

			return

		} else {
			gen824 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen825 := Call(__e, gen824, V3123)

			if True == gen825 {
				gen815 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				gen816 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

				gen817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen818 := Call(__e, gen817, V3123)

				gen819 := Call(__e, gen816, gen818)

				gen820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

				gen821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen822 := Call(__e, gen821, V3123)

				gen823 := Call(__e, gen820, gen822)

				__e.TailApply(gen815, gen819, gen823)

				return

			} else {
				if True == True {
					__e.Return(Nil)
					return
				} else {
					gen814 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen814, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4extract_1pvars, gen829)

	gen852 := MakeNative(func(__e *ControlFlow) {
		V3131 := __e.Get(1)
		_ = V3131
		V3132 := __e.Get(2)
		_ = V3132
		V3133 := __e.Get(3)
		_ = V3133
		gen850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen851 := Call(__e, gen850, Nil, V3131)

		if True == gen851 {
			__e.Return(V3133)
			return
		} else {
			gen848 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen849 := Call(__e, gen848, Nil, V3132)

			if True == gen849 {
				__e.Return(V3133)
				return
			} else {
				gen845 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen846 := Call(__e, gen845, V3131)

				var gen847 Obj

				if True == gen846 {
					gen843 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen844 := Call(__e, gen843, V3132)

					if True == gen844 {
						gen847 = True
					} else {
						gen847 = False
					}

				} else {
					gen847 = False
				}

				if True == gen847 {
					gen832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mult__subst)

					gen833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen834 := Call(__e, gen833, V3131)

					gen835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen836 := Call(__e, gen835, V3132)

					gen837 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					gen838 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen839 := Call(__e, gen838, V3131)

					gen840 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen841 := Call(__e, gen840, V3132)

					gen842 := Call(__e, gen837, gen839, gen841, V3133)

					__e.TailApply(gen832, gen834, gen836, gen842)

					return

				} else {
					if True == True {
						gen831 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen831, symshen_4mult__subst)

						return

					} else {
						gen830 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen830, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4mult__subst, gen852)

	return

}, 0)
