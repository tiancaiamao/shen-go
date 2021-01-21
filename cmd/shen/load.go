package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp77248 := MakeNative(func(__e *ControlFlow) {
		V701 := __e.Get(1)
		_ = V701
		tmp77249 := MakeNative(func(__e *ControlFlow) {
			Load := __e.Get(1)
			_ = Load
			tmp77250 := MakeNative(func(__e *ControlFlow) {
				Infs := __e.Get(1)
				_ = Infs
				__e.Return(symloaded)
				return
			}, 1)

			tmp77262 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp77263 := Call(__e, tmp77262, symshen_4_dtc_d)

			var ifres77251 Obj

			if True == tmp77263 {
				tmp77252 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp77253 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp77254 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp77255 := Call(__e, PrimNS1Value(symns2_1value), syminferences)

				tmp77256 := Call(__e, tmp77255)

				tmp77257 := Call(__e, tmp77254, tmp77256, MakeString(" inferences\n"), symshen_4a)

				tmp77258 := Call(__e, tmp77253, MakeString("\ntypechecked in "), tmp77257)

				tmp77259 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp77260 := Call(__e, tmp77259)

				tmp77261 := Call(__e, tmp77252, tmp77258, tmp77260)

				ifres77251 = tmp77261

			} else {
				ifres77251 = symshen_4skip

			}

			__e.TailApply(tmp77250, ifres77251)
			return

		}, 1)

		tmp77264 := MakeNative(func(__e *ControlFlow) {
			Start := __e.Get(1)
			_ = Start
			tmp77265 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				tmp77266 := MakeNative(func(__e *ControlFlow) {
					Finish := __e.Get(1)
					_ = Finish
					tmp77267 := MakeNative(func(__e *ControlFlow) {
						Time := __e.Get(1)
						_ = Time
						tmp77268 := MakeNative(func(__e *ControlFlow) {
							Message := __e.Get(1)
							_ = Message
							__e.Return(Result)
							return
						}, 1)

						tmp77269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

						tmp77270 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						tmp77271 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						tmp77272 := Call(__e, PrimNS1Value(symns2_1value), symstr)

						tmp77273 := Call(__e, tmp77272, Time)

						tmp77274 := Call(__e, tmp77271, tmp77273, MakeString(" secs\n"))

						tmp77275 := Call(__e, tmp77270, MakeString("\nrun time: "), tmp77274)

						tmp77276 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

						tmp77277 := Call(__e, tmp77276)

						tmp77278 := Call(__e, tmp77269, tmp77275, tmp77277)

						__e.TailApply(tmp77268, tmp77278)
						return

					}, 1)

					tmp77279 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					tmp77280 := Call(__e, tmp77279, Finish, Start)

					__e.TailApply(tmp77267, tmp77280)
					return

				}, 1)

				tmp77281 := Call(__e, PrimNS1Value(symns2_1value), symget_1time)

				tmp77282 := Call(__e, tmp77281, symrun)

				__e.TailApply(tmp77266, tmp77282)
				return

			}, 1)

			tmp77283 := Call(__e, PrimNS1Value(symns2_1value), symshen_4load_1help)

			tmp77284 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp77285 := Call(__e, tmp77284, symshen_4_dtc_d)

			tmp77286 := Call(__e, PrimNS1Value(symns2_1value), symread_1file)

			tmp77287 := Call(__e, tmp77286, V701)

			tmp77288 := Call(__e, tmp77283, tmp77285, tmp77287)

			__e.TailApply(tmp77265, tmp77288)
			return

		}, 1)

		tmp77289 := Call(__e, PrimNS1Value(symns2_1value), symget_1time)

		tmp77290 := Call(__e, tmp77289, symrun)

		tmp77291 := Call(__e, tmp77264, tmp77290)

		__e.TailApply(tmp77249, tmp77291)
		return

	}, 1)

	tmp77292 := Call(__e, PrimNS1Value(symns2_1set), symload, tmp77248)

	_ = tmp77292

	tmp77293 := MakeNative(func(__e *ControlFlow) {
		V708 := __e.Get(1)
		_ = V708
		V709 := __e.Get(2)
		_ = V709
		tmp77325 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp77326 := Call(__e, tmp77325, False, V708)

		if True == tmp77326 {
			tmp77316 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

			tmp77317 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp77318 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp77319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp77320 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				tmp77321 := Call(__e, tmp77320, X)

				tmp77322 := Call(__e, tmp77319, tmp77321, MakeString("\n"), symshen_4s)

				tmp77323 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp77324 := Call(__e, tmp77323)

				__e.TailApply(tmp77318, tmp77322, tmp77324)
				return

			}, 1)

			__e.TailApply(tmp77316, tmp77317, V709)
			return

		} else {
			tmp77295 := MakeNative(func(__e *ControlFlow) {
				RemoveSynonyms := __e.Get(1)
				_ = RemoveSynonyms
				tmp77296 := MakeNative(func(__e *ControlFlow) {
					Table := __e.Get(1)
					_ = Table
					tmp77297 := MakeNative(func(__e *ControlFlow) {
						Assume := __e.Get(1)
						_ = Assume
						tmp77298 := MakeNative(func(__e *ControlFlow) {
							tmp77299 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

							tmp77300 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp77301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1load)

								__e.TailApply(tmp77301, X)
								return

							}, 1)

							__e.TailApply(tmp77299, tmp77300, RemoveSynonyms)
							return

						}, 0)

						tmp77302 := MakeNative(func(__e *ControlFlow) {
							E := __e.Get(1)
							_ = E
							tmp77303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unwind_1types)

							__e.TailApply(tmp77303, E, Table)
							return

						}, 1)

						__e.TailApply(PrimNS1Value(symtry_1catch), tmp77298, tmp77302)
						return

					}, 1)

					tmp77304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

					tmp77305 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp77306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assumetype)

						__e.TailApply(tmp77306, X)
						return

					}, 1)

					tmp77307 := Call(__e, tmp77304, tmp77305, Table)

					__e.TailApply(tmp77297, tmp77307)
					return

				}, 1)

				tmp77308 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

				tmp77309 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp77310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typetable)

					__e.TailApply(tmp77310, X)
					return

				}, 1)

				tmp77311 := Call(__e, tmp77308, tmp77309, RemoveSynonyms)

				__e.TailApply(tmp77296, tmp77311)
				return

			}, 1)

			tmp77312 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

			tmp77313 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp77314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1synonyms)

				__e.TailApply(tmp77314, X)
				return

			}, 1)

			tmp77315 := Call(__e, tmp77312, tmp77313, V709)

			__e.TailApply(tmp77295, tmp77315)
			return

		}

	}, 2)

	tmp77327 := Call(__e, PrimNS1Value(symns2_1set), symshen_4load_1help, tmp77293)

	_ = tmp77327

	tmp77328 := MakeNative(func(__e *ControlFlow) {
		V711 := __e.Get(1)
		_ = V711
		tmp77339 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77340 := Call(__e, tmp77339, V711)

		var ifres77333 Obj

		if True == tmp77340 {
			tmp77335 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77336 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77337 := Call(__e, tmp77336, V711)

			tmp77338 := Call(__e, tmp77335, symshen_4synonyms_1help, tmp77337)

			var ifres77334 Obj

			if True == tmp77338 {
				ifres77334 = True

			} else {
				ifres77334 = False

			}

			ifres77333 = ifres77334

		} else {
			ifres77333 = False

		}

		if True == ifres77333 {
			tmp77331 := Call(__e, PrimNS1Value(symns2_1value), symeval)

			tmp77332 := Call(__e, tmp77331, V711)

			_ = tmp77332

			__e.Return(Nil)
			return

		} else {
			tmp77330 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp77330, V711, Nil)
			return

		}

	}, 1)

	tmp77341 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1synonyms, tmp77328)

	_ = tmp77341

	tmp77342 := MakeNative(func(__e *ControlFlow) {
		V713 := __e.Get(1)
		_ = V713
		tmp77343 := Call(__e, PrimNS1Value(symns2_1value), symnl)

		tmp77344 := Call(__e, tmp77343, MakeNumber(1))

		_ = tmp77344

		tmp77345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

		tmp77346 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

		tmp77347 := Call(__e, tmp77346, symA)

		__e.TailApply(tmp77345, V713, tmp77347)
		return

	}, 1)

	tmp77348 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1load, tmp77342)

	_ = tmp77348

	tmp77349 := MakeNative(func(__e *ControlFlow) {
		V719 := __e.Get(1)
		_ = V719
		tmp77387 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77388 := Call(__e, tmp77387, V719)

		var ifres77375 Obj

		if True == tmp77388 {
			tmp77383 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77384 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77385 := Call(__e, tmp77384, V719)

			tmp77386 := Call(__e, tmp77383, symdefine, tmp77385)

			var ifres77377 Obj

			if True == tmp77386 {
				tmp77379 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77380 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77381 := Call(__e, tmp77380, V719)

				tmp77382 := Call(__e, tmp77379, tmp77381)

				var ifres77378 Obj

				if True == tmp77382 {
					ifres77378 = True

				} else {
					ifres77378 = False

				}

				ifres77377 = ifres77378

			} else {
				ifres77377 = False

			}

			var ifres77376 Obj

			if True == ifres77377 {
				ifres77376 = True

			} else {
				ifres77376 = False

			}

			ifres77375 = ifres77376

		} else {
			ifres77375 = False

		}

		if True == ifres77375 {
			tmp77351 := MakeNative(func(__e *ControlFlow) {
				Sig := __e.Get(1)
				_ = Sig
				tmp77352 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp77354 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77355 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77356 := Call(__e, tmp77355, V719)

				tmp77357 := Call(__e, tmp77354, tmp77356)

				tmp77358 := Call(__e, tmp77353, tmp77357, Sig)

				__e.TailApply(tmp77352, tmp77358, Nil)
				return

			}, 1)

			tmp77359 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

			tmp77360 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				tmp77361 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5sig_7rest_6)

				__e.TailApply(tmp77361, Y)
				return

			}, 1)

			tmp77362 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77363 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77364 := Call(__e, tmp77363, V719)

			tmp77365 := Call(__e, tmp77362, tmp77364)

			tmp77366 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp77367 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp77368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp77369 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77370 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77371 := Call(__e, tmp77370, V719)

				tmp77372 := Call(__e, tmp77369, tmp77371)

				tmp77373 := Call(__e, tmp77368, tmp77372, MakeString(" lacks a proper signature.\n"), symshen_4a)

				__e.TailApply(tmp77367, tmp77373)
				return

			}, 1)

			tmp77374 := Call(__e, tmp77359, tmp77360, tmp77365, tmp77366)

			__e.TailApply(tmp77351, tmp77374)
			return

		} else {
			__e.Return(Nil)
			return
		}

	}, 1)

	tmp77389 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typetable, tmp77349)

	_ = tmp77389

	tmp77390 := MakeNative(func(__e *ControlFlow) {
		V721 := __e.Get(1)
		_ = V721
		tmp77398 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp77399 := Call(__e, tmp77398, V721)

		if True == tmp77399 {
			tmp77393 := Call(__e, PrimNS1Value(symns2_1value), symdeclare)

			tmp77394 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp77395 := Call(__e, tmp77394, V721)

			tmp77396 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp77397 := Call(__e, tmp77396, V721)

			__e.TailApply(tmp77393, tmp77395, tmp77397)
			return

		} else {
			tmp77392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp77392, symshen_4assumetype)
			return

		}

	}, 1)

	tmp77400 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assumetype, tmp77390)

	_ = tmp77400

	tmp77401 := MakeNative(func(__e *ControlFlow) {
		V728 := __e.Get(1)
		_ = V728
		V729 := __e.Get(2)
		_ = V729
		tmp77425 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp77426 := Call(__e, tmp77425, Nil, V729)

		if True == tmp77426 {
			tmp77422 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp77423 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

			tmp77424 := Call(__e, tmp77423, V728)

			__e.TailApply(tmp77422, tmp77424)
			return

		} else {
			tmp77420 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77421 := Call(__e, tmp77420, V729)

			var ifres77414 Obj

			if True == tmp77421 {
				tmp77416 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77417 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77418 := Call(__e, tmp77417, V729)

				tmp77419 := Call(__e, tmp77416, tmp77418)

				var ifres77415 Obj

				if True == tmp77419 {
					ifres77415 = True

				} else {
					ifres77415 = False

				}

				ifres77414 = ifres77415

			} else {
				ifres77414 = False

			}

			if True == ifres77414 {
				tmp77405 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remtype)

				tmp77406 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77407 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77408 := Call(__e, tmp77407, V729)

				tmp77409 := Call(__e, tmp77406, tmp77408)

				tmp77410 := Call(__e, tmp77405, tmp77409)

				_ = tmp77410

				tmp77411 := Call(__e, PrimNS1Value(symns2_1value), symshen_4unwind_1types)

				tmp77412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77413 := Call(__e, tmp77412, V729)

				__e.TailApply(tmp77411, V728, tmp77413)
				return

			} else {
				tmp77404 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp77404, symshen_4unwind_1types)
				return

			}

		}

	}, 2)

	tmp77427 := Call(__e, PrimNS1Value(symns2_1set), symshen_4unwind_1types, tmp77401)

	_ = tmp77427

	tmp77428 := MakeNative(func(__e *ControlFlow) {
		V731 := __e.Get(1)
		_ = V731
		tmp77429 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp77430 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

		tmp77431 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp77432 := Call(__e, tmp77431, symshen_4_dsignedfuncs_d)

		tmp77433 := Call(__e, tmp77430, V731, tmp77432)

		__e.TailApply(tmp77429, symshen_4_dsignedfuncs_d, tmp77433)
		return

	}, 1)

	tmp77434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remtype, tmp77428)

	_ = tmp77434

	tmp77435 := MakeNative(func(__e *ControlFlow) {
		V739 := __e.Get(1)
		_ = V739
		V740 := __e.Get(2)
		_ = V740
		tmp77472 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp77473 := Call(__e, tmp77472, Nil, V740)

		if True == tmp77473 {
			__e.Return(Nil)
			return
		} else {
			tmp77470 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp77471 := Call(__e, tmp77470, V740)

			var ifres77456 Obj

			if True == tmp77471 {
				tmp77466 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77467 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77468 := Call(__e, tmp77467, V740)

				tmp77469 := Call(__e, tmp77466, tmp77468)

				var ifres77458 Obj

				if True == tmp77469 {
					tmp77460 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp77461 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp77462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp77463 := Call(__e, tmp77462, V740)

					tmp77464 := Call(__e, tmp77461, tmp77463)

					tmp77465 := Call(__e, tmp77460, tmp77464, V739)

					var ifres77459 Obj

					if True == tmp77465 {
						ifres77459 = True

					} else {
						ifres77459 = False

					}

					ifres77458 = ifres77459

				} else {
					ifres77458 = False

				}

				var ifres77457 Obj

				if True == ifres77458 {
					ifres77457 = True

				} else {
					ifres77457 = False

				}

				ifres77456 = ifres77457

			} else {
				ifres77456 = False

			}

			if True == ifres77456 {
				tmp77449 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

				tmp77450 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77451 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp77452 := Call(__e, tmp77451, V740)

				tmp77453 := Call(__e, tmp77450, tmp77452)

				tmp77454 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp77455 := Call(__e, tmp77454, V740)

				__e.TailApply(tmp77449, tmp77453, tmp77455)
				return

			} else {
				tmp77447 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp77448 := Call(__e, tmp77447, V740)

				if True == tmp77448 {
					tmp77440 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp77441 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp77442 := Call(__e, tmp77441, V740)

					tmp77443 := Call(__e, PrimNS1Value(symns2_1value), symshen_4removetype)

					tmp77444 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp77445 := Call(__e, tmp77444, V740)

					tmp77446 := Call(__e, tmp77443, V739, tmp77445)

					__e.TailApply(tmp77440, tmp77442, tmp77446)
					return

				} else {
					tmp77439 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp77439, symshen_4removetype)
					return

				}

			}

		}

	}, 2)

	tmp77474 := Call(__e, PrimNS1Value(symns2_1set), symshen_4removetype, tmp77435)

	_ = tmp77474

	tmp77475 := MakeNative(func(__e *ControlFlow) {
		V742 := __e.Get(1)
		_ = V742
		tmp77476 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5signature_6 := __e.Get(1)
			_ = Parse__shen_4_5signature_6
			tmp77495 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp77496 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp77497 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp77498 := Call(__e, tmp77497)

			tmp77499 := Call(__e, tmp77496, tmp77498, Parse__shen_4_5signature_6)

			tmp77500 := Call(__e, tmp77495, tmp77499)

			if True == tmp77500 {
				tmp77479 := MakeNative(func(__e *ControlFlow) {
					Parse___5_b_6 := __e.Get(1)
					_ = Parse___5_b_6
					tmp77487 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp77488 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp77489 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp77490 := Call(__e, tmp77489)

					tmp77491 := Call(__e, tmp77488, tmp77490, Parse___5_b_6)

					tmp77492 := Call(__e, tmp77487, tmp77491)

					if True == tmp77492 {
						tmp77482 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp77483 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp77484 := Call(__e, tmp77483, Parse___5_b_6)

						tmp77485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp77486 := Call(__e, tmp77485, Parse__shen_4_5signature_6)

						__e.TailApply(tmp77482, tmp77484, tmp77486)
						return

					} else {
						tmp77481 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp77481)
						return

					}

				}, 1)

				tmp77493 := Call(__e, PrimNS1Value(symns2_1value), sym_5_b_6)

				tmp77494 := Call(__e, tmp77493, Parse__shen_4_5signature_6)

				__e.TailApply(tmp77479, tmp77494)
				return

			} else {
				tmp77478 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp77478)
				return

			}

		}, 1)

		tmp77501 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5signature_6)

		tmp77502 := Call(__e, tmp77501, V742)

		__e.TailApply(tmp77476, tmp77502)
		return

	}, 1)

	tmp77503 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sig_7rest_6, tmp77475)

	_ = tmp77503

	tmp77504 := MakeNative(func(__e *ControlFlow) {
		V745 := __e.Get(1)
		_ = V745
		V746 := __e.Get(2)
		_ = V746
		tmp77505 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp77506 := MakeNative(func(__e *ControlFlow) {
				String := __e.Get(1)
				_ = String
				tmp77507 := MakeNative(func(__e *ControlFlow) {
					Write := __e.Get(1)
					_ = Write
					tmp77508 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.Return(V746)
						return
					}, 1)

					tmp77509 := Call(__e, PrimNS1Value(symns2_1value), symclose)

					tmp77510 := Call(__e, tmp77509, Stream)

					__e.TailApply(tmp77508, tmp77510)
					return

				}, 1)

				tmp77511 := Call(__e, PrimNS1Value(symns2_1value), sympr)

				tmp77512 := Call(__e, tmp77511, String, Stream)

				__e.TailApply(tmp77507, tmp77512)
				return

			}, 1)

			tmp77518 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

			tmp77519 := Call(__e, tmp77518, V746)

			var ifres77513 Obj

			if True == tmp77519 {
				tmp77516 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp77517 := Call(__e, tmp77516, V746, MakeString("\n\n"), symshen_4a)

				ifres77513 = tmp77517

			} else {
				tmp77514 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp77515 := Call(__e, tmp77514, V746, MakeString("\n\n"), symshen_4s)

				ifres77513 = tmp77515

			}

			__e.TailApply(tmp77506, ifres77513)
			return

		}, 1)

		tmp77520 := Call(__e, PrimNS1Value(symns2_1value), symopen)

		tmp77521 := Call(__e, tmp77520, V745, symout)

		__e.TailApply(tmp77505, tmp77521)
		return

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symwrite_1to_1file, tmp77504)
	return

}, 0)
