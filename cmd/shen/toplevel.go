package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp93265 := MakeNative(func(__e *ControlFlow) {
		tmp93266 := Call(__e, PrimNS1Value(symns2_1value), symshen_4credits)

		tmp93267 := Call(__e, tmp93266)

		_ = tmp93267

		tmp93268 := Call(__e, PrimNS1Value(symns2_1value), symshen_4loop)

		__e.TailApply(tmp93268)
		return

	}, 0)

	tmp93269 := Call(__e, PrimNS1Value(symns2_1set), symshen_4repl, tmp93265)

	_ = tmp93269

	tmp93270 := MakeNative(func(__e *ControlFlow) {
		tmp93271 := Call(__e, PrimNS1Value(symns2_1value), symshen_4initialise__environment)

		tmp93272 := Call(__e, tmp93271)

		_ = tmp93272

		tmp93273 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prompt)

		tmp93274 := Call(__e, tmp93273)

		_ = tmp93274

		tmp93275 := MakeNative(func(__e *ControlFlow) {
			tmp93276 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1evaluate_1print)

			__e.TailApply(tmp93276)
			return

		}, 0)

		tmp93277 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp93278 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel_1display_1exception)

			__e.TailApply(tmp93278, E)
			return

		}, 1)

		tmp93279 := Call(__e, PrimNS1Value(symtry_1catch), tmp93275, tmp93277)

		_ = tmp93279

		tmp93280 := Call(__e, PrimNS1Value(symns2_1value), symshen_4loop)

		__e.TailApply(tmp93280)
		return

	}, 0)

	tmp93281 := Call(__e, PrimNS1Value(symns2_1set), symshen_4loop, tmp93270)

	_ = tmp93281

	tmp93282 := MakeNative(func(__e *ControlFlow) {
		V3022 := __e.Get(1)
		_ = V3022
		tmp93283 := Call(__e, PrimNS1Value(symns2_1value), sympr)

		tmp93284 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

		tmp93285 := Call(__e, tmp93284, V3022)

		tmp93286 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp93287 := Call(__e, tmp93286)

		__e.TailApply(tmp93283, tmp93285, tmp93287)
		return

	}, 1)

	tmp93288 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel_1display_1exception, tmp93282)

	_ = tmp93288

	tmp93289 := MakeNative(func(__e *ControlFlow) {
		tmp93290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp93291 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp93292 := Call(__e, tmp93291)

		tmp93293 := Call(__e, tmp93290, MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n"), tmp93292)

		_ = tmp93293

		tmp93294 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp93295 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp93296 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp93297 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93298 := Call(__e, tmp93297, sym_dversion_d)

		tmp93299 := Call(__e, tmp93296, tmp93298, MakeString("\n"), symshen_4a)

		tmp93300 := Call(__e, tmp93295, MakeString("www.shenlanguage.org, "), tmp93299)

		tmp93301 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp93302 := Call(__e, tmp93301)

		tmp93303 := Call(__e, tmp93294, tmp93300, tmp93302)

		_ = tmp93303

		tmp93304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp93305 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp93306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp93307 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93308 := Call(__e, tmp93307, sym_dlanguage_d)

		tmp93309 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp93310 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp93311 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93312 := Call(__e, tmp93311, sym_dimplementation_d)

		tmp93313 := Call(__e, tmp93310, tmp93312, MakeString(""), symshen_4a)

		tmp93314 := Call(__e, tmp93309, MakeString(", implementation: "), tmp93313)

		tmp93315 := Call(__e, tmp93306, tmp93308, tmp93314, symshen_4a)

		tmp93316 := Call(__e, tmp93305, MakeString("running under "), tmp93315)

		tmp93317 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp93318 := Call(__e, tmp93317)

		tmp93319 := Call(__e, tmp93304, tmp93316, tmp93318)

		_ = tmp93319

		tmp93320 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp93321 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp93322 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp93323 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93324 := Call(__e, tmp93323, sym_dport_d)

		tmp93325 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp93326 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp93327 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93328 := Call(__e, tmp93327, sym_dporters_d)

		tmp93329 := Call(__e, tmp93326, tmp93328, MakeString("\n"), symshen_4a)

		tmp93330 := Call(__e, tmp93325, MakeString(" ported by "), tmp93329)

		tmp93331 := Call(__e, tmp93322, tmp93324, tmp93330, symshen_4a)

		tmp93332 := Call(__e, tmp93321, MakeString("\nport "), tmp93331)

		tmp93333 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp93334 := Call(__e, tmp93333)

		__e.TailApply(tmp93320, tmp93332, tmp93334)
		return

	}, 0)

	tmp93335 := Call(__e, PrimNS1Value(symns2_1set), symshen_4credits, tmp93289)

	_ = tmp93335

	tmp93336 := MakeNative(func(__e *ControlFlow) {
		tmp93337 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiple_1set)

		tmp93338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93346 := Call(__e, tmp93345, MakeNumber(0), Nil)

		tmp93347 := Call(__e, tmp93344, symshen_4_dcatch_d, tmp93346)

		tmp93348 := Call(__e, tmp93343, MakeNumber(0), tmp93347)

		tmp93349 := Call(__e, tmp93342, symshen_4_dprocess_1counter_d, tmp93348)

		tmp93350 := Call(__e, tmp93341, MakeNumber(0), tmp93349)

		tmp93351 := Call(__e, tmp93340, symshen_4_dinfs_d, tmp93350)

		tmp93352 := Call(__e, tmp93339, MakeNumber(0), tmp93351)

		tmp93353 := Call(__e, tmp93338, symshen_4_dcall_d, tmp93352)

		__e.TailApply(tmp93337, tmp93353)
		return

	}, 0)

	tmp93354 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__environment, tmp93336)

	_ = tmp93354

	tmp93355 := MakeNative(func(__e *ControlFlow) {
		V3024 := __e.Get(1)
		_ = V3024
		tmp93380 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93381 := Call(__e, tmp93380, Nil, V3024)

		if True == tmp93381 {
			__e.Return(Nil)
			return
		} else {
			tmp93378 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93379 := Call(__e, tmp93378, V3024)

			var ifres93372 Obj

			if True == tmp93379 {
				tmp93374 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93376 := Call(__e, tmp93375, V3024)

				tmp93377 := Call(__e, tmp93374, tmp93376)

				var ifres93373 Obj

				if True == tmp93377 {
					ifres93373 = True

				} else {
					ifres93373 = False

				}

				ifres93372 = ifres93373

			} else {
				ifres93372 = False

			}

			if True == ifres93372 {
				tmp93359 := Call(__e, PrimNS1Value(symns2_1value), symset)

				tmp93360 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93361 := Call(__e, tmp93360, V3024)

				tmp93362 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93363 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93364 := Call(__e, tmp93363, V3024)

				tmp93365 := Call(__e, tmp93362, tmp93364)

				tmp93366 := Call(__e, tmp93359, tmp93361, tmp93365)

				_ = tmp93366

				tmp93367 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiple_1set)

				tmp93368 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93369 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93370 := Call(__e, tmp93369, V3024)

				tmp93371 := Call(__e, tmp93368, tmp93370)

				__e.TailApply(tmp93367, tmp93371)
				return

			} else {
				tmp93358 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp93358, symshen_4multiple_1set)
				return

			}

		}

	}, 1)

	tmp93382 := Call(__e, PrimNS1Value(symns2_1set), symshen_4multiple_1set, tmp93355)

	_ = tmp93382

	tmp93383 := MakeNative(func(__e *ControlFlow) {
		V3026 := __e.Get(1)
		_ = V3026
		tmp93384 := Call(__e, PrimNS1Value(symns2_1value), symdeclare)

		__e.TailApply(tmp93384, V3026, symsymbol)
		return

	}, 1)

	tmp93385 := Call(__e, PrimNS1Value(symns2_1set), symdestroy, tmp93383)

	_ = tmp93385

	tmp93386 := MakeNative(func(__e *ControlFlow) {
		tmp93387 := MakeNative(func(__e *ControlFlow) {
			Lineread := __e.Get(1)
			_ = Lineread
			tmp93388 := MakeNative(func(__e *ControlFlow) {
				History := __e.Get(1)
				_ = History
				tmp93389 := MakeNative(func(__e *ControlFlow) {
					NewLineread := __e.Get(1)
					_ = NewLineread
					tmp93390 := MakeNative(func(__e *ControlFlow) {
						NewHistory := __e.Get(1)
						_ = NewHistory
						tmp93391 := MakeNative(func(__e *ControlFlow) {
							Parsed := __e.Get(1)
							_ = Parsed
							tmp93392 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel)

							__e.TailApply(tmp93392, Parsed)
							return

						}, 1)

						tmp93393 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						tmp93394 := Call(__e, tmp93393, NewLineread)

						__e.TailApply(tmp93391, tmp93394)
						return

					}, 1)

					tmp93395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4update__history)

					tmp93396 := Call(__e, tmp93395, NewLineread, History)

					__e.TailApply(tmp93390, tmp93396)
					return

				}, 1)

				tmp93397 := Call(__e, PrimNS1Value(symns2_1value), symshen_4retrieve_1from_1history_1if_1needed)

				tmp93398 := Call(__e, tmp93397, Lineread, History)

				__e.TailApply(tmp93389, tmp93398)
				return

			}, 1)

			tmp93399 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp93400 := Call(__e, tmp93399, symshen_4_dhistory_d)

			__e.TailApply(tmp93388, tmp93400)
			return

		}, 1)

		tmp93401 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread)

		tmp93402 := Call(__e, tmp93401)

		__e.TailApply(tmp93387, tmp93402)
		return

	}, 0)

	tmp93403 := Call(__e, PrimNS1Value(symns2_1set), symshen_4read_1evaluate_1print, tmp93386)

	_ = tmp93403

	tmp93404 := MakeNative(func(__e *ControlFlow) {
		V3038 := __e.Get(1)
		_ = V3038
		V3039 := __e.Get(2)
		_ = V3039
		tmp93599 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		tmp93600 := Call(__e, tmp93599, V3038)

		var ifres93577 Obj

		if True == tmp93600 {
			tmp93595 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93596 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			tmp93597 := Call(__e, tmp93596, V3038)

			tmp93598 := Call(__e, tmp93595, tmp93597)

			var ifres93579 Obj

			if True == tmp93598 {
				tmp93581 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp93582 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93583 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				tmp93584 := Call(__e, tmp93583, V3038)

				tmp93585 := Call(__e, tmp93582, tmp93584)

				tmp93586 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp93587 := Call(__e, PrimNS1Value(symns2_1value), symshen_4space)

				tmp93588 := Call(__e, tmp93587)

				tmp93589 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp93590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				tmp93591 := Call(__e, tmp93590)

				tmp93592 := Call(__e, tmp93589, tmp93591, Nil)

				tmp93593 := Call(__e, tmp93586, tmp93588, tmp93592)

				tmp93594 := Call(__e, tmp93581, tmp93585, tmp93593)

				var ifres93580 Obj

				if True == tmp93594 {
					ifres93580 = True

				} else {
					ifres93580 = False

				}

				ifres93579 = ifres93580

			} else {
				ifres93579 = False

			}

			var ifres93578 Obj

			if True == ifres93579 {
				ifres93578 = True

			} else {
				ifres93578 = False

			}

			ifres93577 = ifres93578

		} else {
			ifres93577 = False

		}

		if True == ifres93577 {
			tmp93568 := Call(__e, PrimNS1Value(symns2_1value), symshen_4retrieve_1from_1history_1if_1needed)

			tmp93569 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

			tmp93570 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			tmp93571 := Call(__e, tmp93570, V3038)

			tmp93572 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp93573 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			tmp93574 := Call(__e, tmp93573, V3038)

			tmp93575 := Call(__e, tmp93572, tmp93574)

			tmp93576 := Call(__e, tmp93569, tmp93571, tmp93575)

			__e.TailApply(tmp93568, tmp93576, V3039)
			return

		} else {
			tmp93566 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

			tmp93567 := Call(__e, tmp93566, V3038)

			var ifres93516 Obj

			if True == tmp93567 {
				tmp93562 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93563 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				tmp93564 := Call(__e, tmp93563, V3038)

				tmp93565 := Call(__e, tmp93562, tmp93564)

				var ifres93518 Obj

				if True == tmp93565 {
					tmp93556 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp93557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93558 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp93559 := Call(__e, tmp93558, V3038)

					tmp93560 := Call(__e, tmp93557, tmp93559)

					tmp93561 := Call(__e, tmp93556, tmp93560)

					var ifres93520 Obj

					if True == tmp93561 {
						tmp93548 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp93550 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp93551 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						tmp93552 := Call(__e, tmp93551, V3038)

						tmp93553 := Call(__e, tmp93550, tmp93552)

						tmp93554 := Call(__e, tmp93549, tmp93553)

						tmp93555 := Call(__e, tmp93548, Nil, tmp93554)

						var ifres93522 Obj

						if True == tmp93555 {
							tmp93546 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp93547 := Call(__e, tmp93546, V3039)

							var ifres93524 Obj

							if True == tmp93547 {
								tmp93538 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp93539 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp93540 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								tmp93541 := Call(__e, tmp93540, V3038)

								tmp93542 := Call(__e, tmp93539, tmp93541)

								tmp93543 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

								tmp93544 := Call(__e, tmp93543)

								tmp93545 := Call(__e, tmp93538, tmp93542, tmp93544)

								var ifres93526 Obj

								if True == tmp93545 {
									tmp93528 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp93529 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp93530 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp93531 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

									tmp93532 := Call(__e, tmp93531, V3038)

									tmp93533 := Call(__e, tmp93530, tmp93532)

									tmp93534 := Call(__e, tmp93529, tmp93533)

									tmp93535 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

									tmp93536 := Call(__e, tmp93535)

									tmp93537 := Call(__e, tmp93528, tmp93534, tmp93536)

									var ifres93527 Obj

									if True == tmp93537 {
										ifres93527 = True

									} else {
										ifres93527 = False

									}

									ifres93526 = ifres93527

								} else {
									ifres93526 = False

								}

								var ifres93525 Obj

								if True == ifres93526 {
									ifres93525 = True

								} else {
									ifres93525 = False

								}

								ifres93524 = ifres93525

							} else {
								ifres93524 = False

							}

							var ifres93523 Obj

							if True == ifres93524 {
								ifres93523 = True

							} else {
								ifres93523 = False

							}

							ifres93522 = ifres93523

						} else {
							ifres93522 = False

						}

						var ifres93521 Obj

						if True == ifres93522 {
							ifres93521 = True

						} else {
							ifres93521 = False

						}

						ifres93520 = ifres93521

					} else {
						ifres93520 = False

					}

					var ifres93519 Obj

					if True == ifres93520 {
						ifres93519 = True

					} else {
						ifres93519 = False

					}

					ifres93518 = ifres93519

				} else {
					ifres93518 = False

				}

				var ifres93517 Obj

				if True == ifres93518 {
					ifres93517 = True

				} else {
					ifres93517 = False

				}

				ifres93516 = ifres93517

			} else {
				ifres93516 = False

			}

			if True == ifres93516 {
				tmp93508 := MakeNative(func(__e *ControlFlow) {
					PastPrint := __e.Get(1)
					_ = PastPrint
					tmp93509 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					__e.TailApply(tmp93509, V3039)
					return

				}, 1)

				tmp93510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

				tmp93511 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				tmp93512 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93513 := Call(__e, tmp93512, V3039)

				tmp93514 := Call(__e, tmp93511, tmp93513)

				tmp93515 := Call(__e, tmp93510, tmp93514)

				__e.TailApply(tmp93508, tmp93515)
				return

			} else {
				tmp93506 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

				tmp93507 := Call(__e, tmp93506, V3038)

				var ifres93490 Obj

				if True == tmp93507 {
					tmp93502 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp93503 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp93504 := Call(__e, tmp93503, V3038)

					tmp93505 := Call(__e, tmp93502, tmp93504)

					var ifres93492 Obj

					if True == tmp93505 {
						tmp93494 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93495 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp93496 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						tmp93497 := Call(__e, tmp93496, V3038)

						tmp93498 := Call(__e, tmp93495, tmp93497)

						tmp93499 := Call(__e, PrimNS1Value(symns2_1value), symshen_4exclamation)

						tmp93500 := Call(__e, tmp93499)

						tmp93501 := Call(__e, tmp93494, tmp93498, tmp93500)

						var ifres93493 Obj

						if True == tmp93501 {
							ifres93493 = True

						} else {
							ifres93493 = False

						}

						ifres93492 = ifres93493

					} else {
						ifres93492 = False

					}

					var ifres93491 Obj

					if True == ifres93492 {
						ifres93491 = True

					} else {
						ifres93491 = False

					}

					ifres93490 = ifres93491

				} else {
					ifres93490 = False

				}

				if True == ifres93490 {
					tmp93473 := MakeNative(func(__e *ControlFlow) {
						Key_2 := __e.Get(1)
						_ = Key_2
						tmp93474 := MakeNative(func(__e *ControlFlow) {
							Find := __e.Get(1)
							_ = Find
							tmp93475 := MakeNative(func(__e *ControlFlow) {
								PastPrint := __e.Get(1)
								_ = PastPrint
								__e.Return(Find)
								return
							}, 1)

							tmp93476 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

							tmp93477 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							tmp93478 := Call(__e, tmp93477, Find)

							tmp93479 := Call(__e, tmp93476, tmp93478)

							__e.TailApply(tmp93475, tmp93479)
							return

						}, 1)

						tmp93480 := Call(__e, PrimNS1Value(symns2_1value), symhead)

						tmp93481 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find_1past_1inputs)

						tmp93482 := Call(__e, tmp93481, Key_2, V3039)

						tmp93483 := Call(__e, tmp93480, tmp93482)

						__e.TailApply(tmp93474, tmp93483)
						return

					}, 1)

					tmp93484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make_1key)

					tmp93485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93486 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp93487 := Call(__e, tmp93486, V3038)

					tmp93488 := Call(__e, tmp93485, tmp93487)

					tmp93489 := Call(__e, tmp93484, tmp93488, V3039)

					__e.TailApply(tmp93473, tmp93489)
					return

				} else {
					tmp93471 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					tmp93472 := Call(__e, tmp93471, V3038)

					var ifres93447 Obj

					if True == tmp93472 {
						tmp93467 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp93468 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

						tmp93469 := Call(__e, tmp93468, V3038)

						tmp93470 := Call(__e, tmp93467, tmp93469)

						var ifres93449 Obj

						if True == tmp93470 {
							tmp93461 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp93462 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp93463 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							tmp93464 := Call(__e, tmp93463, V3038)

							tmp93465 := Call(__e, tmp93462, tmp93464)

							tmp93466 := Call(__e, tmp93461, Nil, tmp93465)

							var ifres93451 Obj

							if True == tmp93466 {
								tmp93453 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp93454 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp93455 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								tmp93456 := Call(__e, tmp93455, V3038)

								tmp93457 := Call(__e, tmp93454, tmp93456)

								tmp93458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4percent)

								tmp93459 := Call(__e, tmp93458)

								tmp93460 := Call(__e, tmp93453, tmp93457, tmp93459)

								var ifres93452 Obj

								if True == tmp93460 {
									ifres93452 = True

								} else {
									ifres93452 = False

								}

								ifres93451 = ifres93452

							} else {
								ifres93451 = False

							}

							var ifres93450 Obj

							if True == ifres93451 {
								ifres93450 = True

							} else {
								ifres93450 = False

							}

							ifres93449 = ifres93450

						} else {
							ifres93449 = False

						}

						var ifres93448 Obj

						if True == ifres93449 {
							ifres93448 = True

						} else {
							ifres93448 = False

						}

						ifres93447 = ifres93448

					} else {
						ifres93447 = False

					}

					if True == ifres93447 {
						tmp93441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

						tmp93442 := MakeNative(func(__e *ControlFlow) {
							X := __e.Get(1)
							_ = X
							__e.Return(True)
							return
						}, 1)

						tmp93443 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

						tmp93444 := Call(__e, tmp93443, V3039)

						tmp93445 := Call(__e, tmp93441, tmp93442, tmp93444, MakeNumber(0))

						_ = tmp93445

						tmp93446 := Call(__e, PrimNS1Value(symns2_1value), symabort)

						__e.TailApply(tmp93446)
						return

					} else {
						tmp93439 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

						tmp93440 := Call(__e, tmp93439, V3038)

						var ifres93423 Obj

						if True == tmp93440 {
							tmp93435 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp93436 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							tmp93437 := Call(__e, tmp93436, V3038)

							tmp93438 := Call(__e, tmp93435, tmp93437)

							var ifres93425 Obj

							if True == tmp93438 {
								tmp93427 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

								tmp93428 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp93429 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

								tmp93430 := Call(__e, tmp93429, V3038)

								tmp93431 := Call(__e, tmp93428, tmp93430)

								tmp93432 := Call(__e, PrimNS1Value(symns2_1value), symshen_4percent)

								tmp93433 := Call(__e, tmp93432)

								tmp93434 := Call(__e, tmp93427, tmp93431, tmp93433)

								var ifres93426 Obj

								if True == tmp93434 {
									ifres93426 = True

								} else {
									ifres93426 = False

								}

								ifres93425 = ifres93426

							} else {
								ifres93425 = False

							}

							var ifres93424 Obj

							if True == ifres93425 {
								ifres93424 = True

							} else {
								ifres93424 = False

							}

							ifres93423 = ifres93424

						} else {
							ifres93423 = False

						}

						if True == ifres93423 {
							tmp93410 := MakeNative(func(__e *ControlFlow) {
								Key_2 := __e.Get(1)
								_ = Key_2
								tmp93411 := MakeNative(func(__e *ControlFlow) {
									Pastprint := __e.Get(1)
									_ = Pastprint
									tmp93412 := Call(__e, PrimNS1Value(symns2_1value), symabort)

									__e.TailApply(tmp93412)
									return

								}, 1)

								tmp93413 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

								tmp93414 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

								tmp93415 := Call(__e, tmp93414, V3039)

								tmp93416 := Call(__e, tmp93413, Key_2, tmp93415, MakeNumber(0))

								__e.TailApply(tmp93411, tmp93416)
								return

							}, 1)

							tmp93417 := Call(__e, PrimNS1Value(symns2_1value), symshen_4make_1key)

							tmp93418 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp93419 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

							tmp93420 := Call(__e, tmp93419, V3038)

							tmp93421 := Call(__e, tmp93418, tmp93420)

							tmp93422 := Call(__e, tmp93417, tmp93421, V3039)

							__e.TailApply(tmp93410, tmp93422)
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

	tmp93601 := Call(__e, PrimNS1Value(symns2_1set), symshen_4retrieve_1from_1history_1if_1needed, tmp93404)

	_ = tmp93601

	tmp93602 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(37))
		return
	}, 0)

	tmp93603 := Call(__e, PrimNS1Value(symns2_1set), symshen_4percent, tmp93602)

	_ = tmp93603

	tmp93604 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(33))
		return
	}, 0)

	tmp93605 := Call(__e, PrimNS1Value(symns2_1set), symshen_4exclamation, tmp93604)

	_ = tmp93605

	tmp93606 := MakeNative(func(__e *ControlFlow) {
		V3041 := __e.Get(1)
		_ = V3041
		tmp93607 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

		tmp93608 := MakeNative(func(__e *ControlFlow) {
			Byte := __e.Get(1)
			_ = Byte
			tmp93609 := Call(__e, PrimNS1Value(symns2_1value), sympr)

			tmp93610 := Call(__e, PrimNS1Value(symns2_1value), symn_1_6string)

			tmp93611 := Call(__e, tmp93610, Byte)

			tmp93612 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp93613 := Call(__e, tmp93612)

			__e.TailApply(tmp93609, tmp93611, tmp93613)
			return

		}, 1)

		tmp93614 := Call(__e, tmp93607, tmp93608, V3041)

		_ = tmp93614

		tmp93615 := Call(__e, PrimNS1Value(symns2_1value), symnl)

		__e.TailApply(tmp93615, MakeNumber(1))
		return

	}, 1)

	tmp93616 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prbytes, tmp93606)

	_ = tmp93616

	tmp93617 := MakeNative(func(__e *ControlFlow) {
		V3044 := __e.Get(1)
		_ = V3044
		V3045 := __e.Get(2)
		_ = V3045
		tmp93618 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp93619 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp93620 := Call(__e, tmp93619, V3044, V3045)

		__e.TailApply(tmp93618, symshen_4_dhistory_d, tmp93620)
		return

	}, 2)

	tmp93621 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update__history, tmp93617)

	_ = tmp93621

	tmp93622 := MakeNative(func(__e *ControlFlow) {
		tmp93623 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

		tmp93624 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

		tmp93625 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

		tmp93626 := Call(__e, tmp93625)

		tmp93627 := Call(__e, tmp93624, tmp93626)

		__e.TailApply(tmp93623, tmp93627, Nil)
		return

	}, 0)

	tmp93628 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread, tmp93622)

	_ = tmp93628

	tmp93629 := MakeNative(func(__e *ControlFlow) {
		V3049 := __e.Get(1)
		_ = V3049
		V3050 := __e.Get(2)
		_ = V3050
		tmp93681 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93682 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hat)

		tmp93683 := Call(__e, tmp93682)

		tmp93684 := Call(__e, tmp93681, V3049, tmp93683)

		if True == tmp93684 {
			tmp93680 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp93680, MakeString("line read aborted"))
			return

		} else {
			tmp93670 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

			tmp93671 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp93672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

			tmp93673 := Call(__e, tmp93672)

			tmp93674 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp93675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

			tmp93676 := Call(__e, tmp93675)

			tmp93677 := Call(__e, tmp93674, tmp93676, Nil)

			tmp93678 := Call(__e, tmp93671, tmp93673, tmp93677)

			tmp93679 := Call(__e, tmp93670, V3049, tmp93678)

			if True == tmp93679 {
				tmp93644 := MakeNative(func(__e *ControlFlow) {
					Line := __e.Get(1)
					_ = Line
					tmp93645 := MakeNative(func(__e *ControlFlow) {
						It := __e.Get(1)
						_ = It
						tmp93661 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93662 := Call(__e, tmp93661, Line, symshen_4nextline)

						var ifres93657 Obj

						if True == tmp93662 {
							ifres93657 = True

						} else {
							tmp93659 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

							tmp93660 := Call(__e, tmp93659, Line)

							var ifres93658 Obj

							if True == tmp93660 {
								ifres93658 = True

							} else {
								ifres93658 = False

							}

							ifres93657 = ifres93658

						}

						if True == ifres93657 {
							tmp93648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

							tmp93649 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

							tmp93650 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

							tmp93651 := Call(__e, tmp93650)

							tmp93652 := Call(__e, tmp93649, tmp93651)

							tmp93653 := Call(__e, PrimNS1Value(symns2_1value), symappend)

							tmp93654 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp93655 := Call(__e, tmp93654, V3049, Nil)

							tmp93656 := Call(__e, tmp93653, V3050, tmp93655)

							__e.TailApply(tmp93648, tmp93652, tmp93656)
							return

						} else {
							tmp93647 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

							__e.TailApply(tmp93647, Line, V3050)
							return

						}

					}, 1)

					tmp93663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4record_1it)

					tmp93664 := Call(__e, tmp93663, V3050)

					__e.TailApply(tmp93645, tmp93664)
					return

				}, 1)

				tmp93665 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

				tmp93666 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp93667 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

					__e.TailApply(tmp93667, X)
					return

				}, 1)

				tmp93668 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4nextline)
					return
				}, 1)

				tmp93669 := Call(__e, tmp93665, tmp93666, V3050, tmp93668)

				__e.TailApply(tmp93644, tmp93669)
				return

			} else {
				tmp93632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplineread__loop)

				tmp93633 := Call(__e, PrimNS1Value(symns2_1value), symshen_4read_1char_1code)

				tmp93634 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

				tmp93635 := Call(__e, tmp93634)

				tmp93636 := Call(__e, tmp93633, tmp93635)

				tmp93642 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp93643 := Call(__e, tmp93642, V3049, MakeNumber(-1))

				var ifres93637 Obj

				if True == tmp93643 {
					ifres93637 = V3050

				} else {
					tmp93638 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp93639 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp93640 := Call(__e, tmp93639, V3049, Nil)

					tmp93641 := Call(__e, tmp93638, V3050, tmp93640)

					ifres93637 = tmp93641

				}

				__e.TailApply(tmp93632, tmp93636, ifres93637)
				return

			}

		}

	}, 2)

	tmp93685 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplineread__loop, tmp93629)

	_ = tmp93685

	tmp93686 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(94))
		return
	}, 0)

	tmp93687 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hat, tmp93686)

	_ = tmp93687

	tmp93688 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(10))
		return
	}, 0)

	tmp93689 := Call(__e, PrimNS1Value(symns2_1set), symshen_4newline, tmp93688)

	_ = tmp93689

	tmp93690 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(13))
		return
	}, 0)

	tmp93691 := Call(__e, PrimNS1Value(symns2_1set), symshen_4carriage_1return, tmp93690)

	_ = tmp93691

	tmp93692 := MakeNative(func(__e *ControlFlow) {
		V3056 := __e.Get(1)
		_ = V3056
		tmp93700 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93701 := Call(__e, tmp93700, sym_7, V3056)

		if True == tmp93701 {
			tmp93699 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp93699, symshen_4_dtc_d, True)
			return

		} else {
			tmp93697 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp93698 := Call(__e, tmp93697, sym_1, V3056)

			if True == tmp93698 {
				tmp93696 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp93696, symshen_4_dtc_d, False)
				return

			} else {
				tmp93695 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp93695, MakeString("tc expects a + or -"))
				return

			}

		}

	}, 1)

	tmp93702 := Call(__e, PrimNS1Value(symns2_1set), symtc, tmp93692)

	_ = tmp93702

	tmp93703 := MakeNative(func(__e *ControlFlow) {
		tmp93727 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93728 := Call(__e, tmp93727, symshen_4_dtc_d)

		if True == tmp93728 {
			tmp93716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp93717 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp93718 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp93719 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp93720 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp93721 := Call(__e, tmp93720, symshen_4_dhistory_d)

			tmp93722 := Call(__e, tmp93719, tmp93721)

			tmp93723 := Call(__e, tmp93718, tmp93722, MakeString("+) "), symshen_4a)

			tmp93724 := Call(__e, tmp93717, MakeString("\n\n("), tmp93723)

			tmp93725 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp93726 := Call(__e, tmp93725)

			__e.TailApply(tmp93716, tmp93724, tmp93726)
			return

		} else {
			tmp93705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp93706 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp93707 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp93708 := Call(__e, PrimNS1Value(symns2_1value), symlength)

			tmp93709 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp93710 := Call(__e, tmp93709, symshen_4_dhistory_d)

			tmp93711 := Call(__e, tmp93708, tmp93710)

			tmp93712 := Call(__e, tmp93707, tmp93711, MakeString("-) "), symshen_4a)

			tmp93713 := Call(__e, tmp93706, MakeString("\n\n("), tmp93712)

			tmp93714 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp93715 := Call(__e, tmp93714)

			__e.TailApply(tmp93705, tmp93713, tmp93715)
			return

		}

	}, 0)

	tmp93729 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prompt, tmp93703)

	_ = tmp93729

	tmp93730 := MakeNative(func(__e *ControlFlow) {
		V3058 := __e.Get(1)
		_ = V3058
		tmp93731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

		tmp93732 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp93733 := Call(__e, tmp93732, symshen_4_dtc_d)

		__e.TailApply(tmp93731, V3058, tmp93733)
		return

	}, 1)

	tmp93734 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel, tmp93730)

	_ = tmp93734

	tmp93735 := MakeNative(func(__e *ControlFlow) {
		V3061 := __e.Get(1)
		_ = V3061
		V3062 := __e.Get(2)
		_ = V3062
		tmp93736 := MakeNative(func(__e *ControlFlow) {
			F := __e.Get(1)
			_ = F
			tmp93739 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

			tmp93740 := Call(__e, tmp93739, F)

			if True == tmp93740 {
				tmp93738 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp93738, MakeString("input not found\n"))
				return

			} else {
				__e.Return(F)
				return
			}

		}, 1)

		tmp93741 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

		tmp93742 := Call(__e, tmp93741, V3061, V3062)

		__e.TailApply(tmp93736, tmp93742)
		return

	}, 2)

	tmp93743 := Call(__e, PrimNS1Value(symns2_1set), symshen_4find_1past_1inputs, tmp93735)

	_ = tmp93743

	tmp93744 := MakeNative(func(__e *ControlFlow) {
		V3065 := __e.Get(1)
		_ = V3065
		V3066 := __e.Get(2)
		_ = V3066
		tmp93745 := MakeNative(func(__e *ControlFlow) {
			Atom := __e.Get(1)
			_ = Atom
			tmp93759 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

			tmp93760 := Call(__e, tmp93759, Atom)

			if True == tmp93760 {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp93752 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp93753 := Call(__e, PrimNS1Value(symns2_1value), symnth)

					tmp93754 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					tmp93755 := Call(__e, tmp93754, Atom, MakeNumber(1))

					tmp93756 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					tmp93757 := Call(__e, tmp93756, V3066)

					tmp93758 := Call(__e, tmp93753, tmp93755, tmp93757)

					__e.TailApply(tmp93752, X, tmp93758)
					return

				}, 1))
				return
			} else {
				__e.Return(MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp93747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

					tmp93748 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

					tmp93749 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp93750 := Call(__e, tmp93749, X)

					tmp93751 := Call(__e, tmp93748, tmp93750)

					__e.TailApply(tmp93747, V3065, tmp93751)
					return

				}, 1))
				return
			}

		}, 1)

		tmp93761 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp93762 := Call(__e, PrimNS1Value(symns2_1value), symcompile)

		tmp93763 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp93764 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5st__input_6)

			__e.TailApply(tmp93764, X)
			return

		}, 1)

		tmp93765 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			tmp93773 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93774 := Call(__e, tmp93773, E)

			if True == tmp93774 {
				tmp93768 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				tmp93769 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				tmp93770 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp93771 := Call(__e, tmp93770, E, MakeString("\n"), symshen_4s)

				tmp93772 := Call(__e, tmp93769, MakeString("parse error here: "), tmp93771)

				__e.TailApply(tmp93768, tmp93772)
				return

			} else {
				tmp93767 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp93767, MakeString("parse error\n"))
				return

			}

		}, 1)

		tmp93775 := Call(__e, tmp93762, tmp93763, V3065, tmp93765)

		tmp93776 := Call(__e, tmp93761, tmp93775)

		__e.TailApply(tmp93745, tmp93776)
		return

	}, 2)

	tmp93777 := Call(__e, PrimNS1Value(symns2_1set), symshen_4make_1key, tmp93744)

	_ = tmp93777

	tmp93778 := MakeNative(func(__e *ControlFlow) {
		V3068 := __e.Get(1)
		_ = V3068
		tmp93847 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp93848 := Call(__e, tmp93847, V3068)

		var ifres93839 Obj

		if True == tmp93848 {
			tmp93841 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp93842 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp93843 := Call(__e, tmp93842, V3068)

			tmp93844 := Call(__e, PrimNS1Value(symns2_1value), symshen_4space)

			tmp93845 := Call(__e, tmp93844)

			tmp93846 := Call(__e, tmp93841, tmp93843, tmp93845)

			var ifres93840 Obj

			if True == tmp93846 {
				ifres93840 = True

			} else {
				ifres93840 = False

			}

			ifres93839 = ifres93840

		} else {
			ifres93839 = False

		}

		if True == ifres93839 {
			tmp93836 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

			tmp93837 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp93838 := Call(__e, tmp93837, V3068)

			__e.TailApply(tmp93836, tmp93838)
			return

		} else {
			tmp93834 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93835 := Call(__e, tmp93834, V3068)

			var ifres93826 Obj

			if True == tmp93835 {
				tmp93828 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp93829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93830 := Call(__e, tmp93829, V3068)

				tmp93831 := Call(__e, PrimNS1Value(symns2_1value), symshen_4newline)

				tmp93832 := Call(__e, tmp93831)

				tmp93833 := Call(__e, tmp93828, tmp93830, tmp93832)

				var ifres93827 Obj

				if True == tmp93833 {
					ifres93827 = True

				} else {
					ifres93827 = False

				}

				ifres93826 = ifres93827

			} else {
				ifres93826 = False

			}

			if True == ifres93826 {
				tmp93823 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

				tmp93824 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93825 := Call(__e, tmp93824, V3068)

				__e.TailApply(tmp93823, tmp93825)
				return

			} else {
				tmp93821 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93822 := Call(__e, tmp93821, V3068)

				var ifres93813 Obj

				if True == tmp93822 {
					tmp93815 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp93816 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93817 := Call(__e, tmp93816, V3068)

					tmp93818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4carriage_1return)

					tmp93819 := Call(__e, tmp93818)

					tmp93820 := Call(__e, tmp93815, tmp93817, tmp93819)

					var ifres93814 Obj

					if True == tmp93820 {
						ifres93814 = True

					} else {
						ifres93814 = False

					}

					ifres93813 = ifres93814

				} else {
					ifres93813 = False

				}

				if True == ifres93813 {
					tmp93810 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

					tmp93811 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93812 := Call(__e, tmp93811, V3068)

					__e.TailApply(tmp93810, tmp93812)
					return

				} else {
					tmp93808 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp93809 := Call(__e, tmp93808, V3068)

					var ifres93800 Obj

					if True == tmp93809 {
						tmp93802 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93803 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp93804 := Call(__e, tmp93803, V3068)

						tmp93805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tab)

						tmp93806 := Call(__e, tmp93805)

						tmp93807 := Call(__e, tmp93802, tmp93804, tmp93806)

						var ifres93801 Obj

						if True == tmp93807 {
							ifres93801 = True

						} else {
							ifres93801 = False

						}

						ifres93800 = ifres93801

					} else {
						ifres93800 = False

					}

					if True == ifres93800 {
						tmp93797 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

						tmp93798 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp93799 := Call(__e, tmp93798, V3068)

						__e.TailApply(tmp93797, tmp93799)
						return

					} else {
						tmp93795 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp93796 := Call(__e, tmp93795, V3068)

						var ifres93787 Obj

						if True == tmp93796 {
							tmp93789 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp93790 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp93791 := Call(__e, tmp93790, V3068)

							tmp93792 := Call(__e, PrimNS1Value(symns2_1value), symshen_4left_1round)

							tmp93793 := Call(__e, tmp93792)

							tmp93794 := Call(__e, tmp93789, tmp93791, tmp93793)

							var ifres93788 Obj

							if True == tmp93794 {
								ifres93788 = True

							} else {
								ifres93788 = False

							}

							ifres93787 = ifres93788

						} else {
							ifres93787 = False

						}

						if True == ifres93787 {
							tmp93784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4trim_1gubbins)

							tmp93785 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp93786 := Call(__e, tmp93785, V3068)

							__e.TailApply(tmp93784, tmp93786)
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

	tmp93849 := Call(__e, PrimNS1Value(symns2_1set), symshen_4trim_1gubbins, tmp93778)

	_ = tmp93849

	tmp93850 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(32))
		return
	}, 0)

	tmp93851 := Call(__e, PrimNS1Value(symns2_1set), symshen_4space, tmp93850)

	_ = tmp93851

	tmp93852 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(9))
		return
	}, 0)

	tmp93853 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tab, tmp93852)

	_ = tmp93853

	tmp93854 := MakeNative(func(__e *ControlFlow) {
		__e.Return(MakeNumber(40))
		return
	}, 0)

	tmp93855 := Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1round, tmp93854)

	_ = tmp93855

	tmp93856 := MakeNative(func(__e *ControlFlow) {
		V3077 := __e.Get(1)
		_ = V3077
		V3078 := __e.Get(2)
		_ = V3078
		tmp93880 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93881 := Call(__e, tmp93880, Nil, V3078)

		if True == tmp93881 {
			__e.Return(Nil)
			return
		} else {
			tmp93878 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93879 := Call(__e, tmp93878, V3078)

			var ifres93873 Obj

			if True == tmp93879 {
				tmp93875 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93876 := Call(__e, tmp93875, V3078)

				tmp93877 := Call(__e, V3077, tmp93876)

				var ifres93874 Obj

				if True == tmp93877 {
					ifres93874 = True

				} else {
					ifres93874 = False

				}

				ifres93873 = ifres93874

			} else {
				ifres93873 = False

			}

			if True == ifres93873 {
				tmp93866 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp93867 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93868 := Call(__e, tmp93867, V3078)

				tmp93869 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

				tmp93870 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93871 := Call(__e, tmp93870, V3078)

				tmp93872 := Call(__e, tmp93869, V3077, tmp93871)

				__e.TailApply(tmp93866, tmp93868, tmp93872)
				return

			} else {
				tmp93864 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93865 := Call(__e, tmp93864, V3078)

				if True == tmp93865 {
					tmp93861 := Call(__e, PrimNS1Value(symns2_1value), symshen_4find)

					tmp93862 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93863 := Call(__e, tmp93862, V3078)

					__e.TailApply(tmp93861, V3077, tmp93863)
					return

				} else {
					tmp93860 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp93860, symshen_4find)
					return

				}

			}

		}

	}, 2)

	tmp93882 := Call(__e, PrimNS1Value(symns2_1set), symshen_4find, tmp93856)

	_ = tmp93882

	tmp93883 := MakeNative(func(__e *ControlFlow) {
		V3092 := __e.Get(1)
		_ = V3092
		V3093 := __e.Get(2)
		_ = V3093
		tmp93905 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93906 := Call(__e, tmp93905, Nil, V3092)

		if True == tmp93906 {
			__e.Return(True)
			return
		} else {
			tmp93903 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93904 := Call(__e, tmp93903, V3092)

			var ifres93891 Obj

			if True == tmp93904 {
				tmp93901 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93902 := Call(__e, tmp93901, V3093)

				var ifres93893 Obj

				if True == tmp93902 {
					tmp93895 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp93896 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93897 := Call(__e, tmp93896, V3093)

					tmp93898 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93899 := Call(__e, tmp93898, V3092)

					tmp93900 := Call(__e, tmp93895, tmp93897, tmp93899)

					var ifres93894 Obj

					if True == tmp93900 {
						ifres93894 = True

					} else {
						ifres93894 = False

					}

					ifres93893 = ifres93894

				} else {
					ifres93893 = False

				}

				var ifres93892 Obj

				if True == ifres93893 {
					ifres93892 = True

				} else {
					ifres93892 = False

				}

				ifres93891 = ifres93892

			} else {
				ifres93891 = False

			}

			if True == ifres93891 {
				tmp93886 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prefix_2)

				tmp93887 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93888 := Call(__e, tmp93887, V3092)

				tmp93889 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93890 := Call(__e, tmp93889, V3093)

				__e.TailApply(tmp93886, tmp93888, tmp93890)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp93907 := Call(__e, PrimNS1Value(symns2_1set), symshen_4prefix_2, tmp93883)

	_ = tmp93907

	tmp93908 := MakeNative(func(__e *ControlFlow) {
		V3105 := __e.Get(1)
		_ = V3105
		V3106 := __e.Get(2)
		_ = V3106
		V3107 := __e.Get(3)
		_ = V3107
		tmp93952 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp93953 := Call(__e, tmp93952, Nil, V3106)

		if True == tmp93953 {
			__e.Return(sym__)
			return
		} else {
			tmp93950 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp93951 := Call(__e, tmp93950, V3106)

			var ifres93943 Obj

			if True == tmp93951 {
				tmp93945 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp93946 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93947 := Call(__e, tmp93946, V3106)

				tmp93948 := Call(__e, V3105, tmp93947)

				tmp93949 := Call(__e, tmp93945, tmp93948)

				var ifres93944 Obj

				if True == tmp93949 {
					ifres93944 = True

				} else {
					ifres93944 = False

				}

				ifres93943 = ifres93944

			} else {
				ifres93943 = False

			}

			if True == ifres93943 {
				tmp93938 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

				tmp93939 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp93940 := Call(__e, tmp93939, V3106)

				tmp93941 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				tmp93942 := Call(__e, tmp93941, V3107, MakeNumber(1))

				__e.TailApply(tmp93938, V3105, tmp93940, tmp93942)
				return

			} else {
				tmp93936 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93937 := Call(__e, tmp93936, V3106)

				var ifres93930 Obj

				if True == tmp93937 {
					tmp93932 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					tmp93933 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93934 := Call(__e, tmp93933, V3106)

					tmp93935 := Call(__e, tmp93932, tmp93934)

					var ifres93931 Obj

					if True == tmp93935 {
						ifres93931 = True

					} else {
						ifres93931 = False

					}

					ifres93930 = ifres93931

				} else {
					ifres93930 = False

				}

				if True == ifres93930 {
					tmp93913 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

					tmp93914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					tmp93915 := Call(__e, tmp93914, V3107, MakeString(". "), symshen_4a)

					tmp93916 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

					tmp93917 := Call(__e, tmp93916)

					tmp93918 := Call(__e, tmp93913, tmp93915, tmp93917)

					_ = tmp93918

					tmp93919 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prbytes)

					tmp93920 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp93921 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93922 := Call(__e, tmp93921, V3106)

					tmp93923 := Call(__e, tmp93920, tmp93922)

					tmp93924 := Call(__e, tmp93919, tmp93923)

					_ = tmp93924

					tmp93925 := Call(__e, PrimNS1Value(symns2_1value), symshen_4print_1past_1inputs)

					tmp93926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93927 := Call(__e, tmp93926, V3106)

					tmp93928 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

					tmp93929 := Call(__e, tmp93928, V3107, MakeNumber(1))

					__e.TailApply(tmp93925, V3105, tmp93927, tmp93929)
					return

				} else {
					tmp93912 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp93912, symshen_4print_1past_1inputs)
					return

				}

			}

		}

	}, 3)

	tmp93954 := Call(__e, PrimNS1Value(symns2_1set), symshen_4print_1past_1inputs, tmp93908)

	_ = tmp93954

	tmp93955 := MakeNative(func(__e *ControlFlow) {
		V3110 := __e.Get(1)
		_ = V3110
		V3111 := __e.Get(2)
		_ = V3111
		tmp94060 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp94061 := Call(__e, tmp94060, V3110)

		var ifres94024 Obj

		if True == tmp94061 {
			tmp94056 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp94057 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94058 := Call(__e, tmp94057, V3110)

			tmp94059 := Call(__e, tmp94056, tmp94058)

			var ifres94026 Obj

			if True == tmp94059 {
				tmp94050 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp94051 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94052 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94053 := Call(__e, tmp94052, V3110)

				tmp94054 := Call(__e, tmp94051, tmp94053)

				tmp94055 := Call(__e, tmp94050, sym_h, tmp94054)

				var ifres94028 Obj

				if True == tmp94055 {
					tmp94044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp94045 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94046 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94047 := Call(__e, tmp94046, V3110)

					tmp94048 := Call(__e, tmp94045, tmp94047)

					tmp94049 := Call(__e, tmp94044, tmp94048)

					var ifres94030 Obj

					if True == tmp94049 {
						tmp94036 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp94037 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94038 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94039 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94040 := Call(__e, tmp94039, V3110)

						tmp94041 := Call(__e, tmp94038, tmp94040)

						tmp94042 := Call(__e, tmp94037, tmp94041)

						tmp94043 := Call(__e, tmp94036, Nil, tmp94042)

						var ifres94032 Obj

						if True == tmp94043 {
							tmp94034 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp94035 := Call(__e, tmp94034, True, V3111)

							var ifres94033 Obj

							if True == tmp94035 {
								ifres94033 = True

							} else {
								ifres94033 = False

							}

							ifres94032 = ifres94033

						} else {
							ifres94032 = False

						}

						var ifres94031 Obj

						if True == ifres94032 {
							ifres94031 = True

						} else {
							ifres94031 = False

						}

						ifres94030 = ifres94031

					} else {
						ifres94030 = False

					}

					var ifres94029 Obj

					if True == ifres94030 {
						ifres94029 = True

					} else {
						ifres94029 = False

					}

					ifres94028 = ifres94029

				} else {
					ifres94028 = False

				}

				var ifres94027 Obj

				if True == ifres94028 {
					ifres94027 = True

				} else {
					ifres94027 = False

				}

				ifres94026 = ifres94027

			} else {
				ifres94026 = False

			}

			var ifres94025 Obj

			if True == ifres94026 {
				ifres94025 = True

			} else {
				ifres94025 = False

			}

			ifres94024 = ifres94025

		} else {
			ifres94024 = False

		}

		if True == ifres94024 {
			tmp94015 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

			tmp94016 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94017 := Call(__e, tmp94016, V3110)

			tmp94018 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94019 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94020 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94021 := Call(__e, tmp94020, V3110)

			tmp94022 := Call(__e, tmp94019, tmp94021)

			tmp94023 := Call(__e, tmp94018, tmp94022)

			__e.TailApply(tmp94015, tmp94017, tmp94023)
			return

		} else {
			tmp94013 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp94014 := Call(__e, tmp94013, V3110)

			var ifres94007 Obj

			if True == tmp94014 {
				tmp94009 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94010 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94011 := Call(__e, tmp94010, V3110)

				tmp94012 := Call(__e, tmp94009, tmp94011)

				var ifres94008 Obj

				if True == tmp94012 {
					ifres94008 = True

				} else {
					ifres94008 = False

				}

				ifres94007 = ifres94008

			} else {
				ifres94007 = False

			}

			if True == ifres94007 {
				tmp93996 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

				tmp93997 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp93998 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp93999 := Call(__e, tmp93998, V3110)

				tmp94000 := Call(__e, tmp93997, tmp93999, Nil)

				tmp94001 := Call(__e, tmp93996, tmp94000, V3111)

				_ = tmp94001

				tmp94002 := Call(__e, PrimNS1Value(symns2_1value), symnl)

				tmp94003 := Call(__e, tmp94002, MakeNumber(1))

				_ = tmp94003

				tmp94004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4toplevel__evaluate)

				tmp94005 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94006 := Call(__e, tmp94005, V3110)

				__e.TailApply(tmp94004, tmp94006, V3111)
				return

			} else {
				tmp93994 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp93995 := Call(__e, tmp93994, V3110)

				var ifres93984 Obj

				if True == tmp93995 {
					tmp93990 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp93991 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp93992 := Call(__e, tmp93991, V3110)

					tmp93993 := Call(__e, tmp93990, Nil, tmp93992)

					var ifres93986 Obj

					if True == tmp93993 {
						tmp93988 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93989 := Call(__e, tmp93988, True, V3111)

						var ifres93987 Obj

						if True == tmp93989 {
							ifres93987 = True

						} else {
							ifres93987 = False

						}

						ifres93986 = ifres93987

					} else {
						ifres93986 = False

					}

					var ifres93985 Obj

					if True == ifres93986 {
						ifres93985 = True

					} else {
						ifres93985 = False

					}

					ifres93984 = ifres93985

				} else {
					ifres93984 = False

				}

				if True == ifres93984 {
					tmp93979 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck_1and_1evaluate)

					tmp93980 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp93981 := Call(__e, tmp93980, V3110)

					tmp93982 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

					tmp93983 := Call(__e, tmp93982, symA)

					__e.TailApply(tmp93979, tmp93981, tmp93983)
					return

				} else {
					tmp93977 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp93978 := Call(__e, tmp93977, V3110)

					var ifres93967 Obj

					if True == tmp93978 {
						tmp93973 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp93974 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp93975 := Call(__e, tmp93974, V3110)

						tmp93976 := Call(__e, tmp93973, Nil, tmp93975)

						var ifres93969 Obj

						if True == tmp93976 {
							tmp93971 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp93972 := Call(__e, tmp93971, False, V3111)

							var ifres93970 Obj

							if True == tmp93972 {
								ifres93970 = True

							} else {
								ifres93970 = False

							}

							ifres93969 = ifres93970

						} else {
							ifres93969 = False

						}

						var ifres93968 Obj

						if True == ifres93969 {
							ifres93968 = True

						} else {
							ifres93968 = False

						}

						ifres93967 = ifres93968

					} else {
						ifres93967 = False

					}

					if True == ifres93967 {
						tmp93961 := MakeNative(func(__e *ControlFlow) {
							Eval := __e.Get(1)
							_ = Eval
							tmp93962 := Call(__e, PrimNS1Value(symns2_1value), symprint)

							__e.TailApply(tmp93962, Eval)
							return

						}, 1)

						tmp93963 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

						tmp93964 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp93965 := Call(__e, tmp93964, V3110)

						tmp93966 := Call(__e, tmp93963, tmp93965)

						__e.TailApply(tmp93961, tmp93966)
						return

					} else {
						tmp93960 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp93960, symshen_4toplevel__evaluate)
						return

					}

				}

			}

		}

	}, 2)

	tmp94062 := Call(__e, PrimNS1Value(symns2_1set), symshen_4toplevel__evaluate, tmp93955)

	_ = tmp94062

	tmp94063 := MakeNative(func(__e *ControlFlow) {
		V3114 := __e.Get(1)
		_ = V3114
		V3115 := __e.Get(2)
		_ = V3115
		tmp94064 := MakeNative(func(__e *ControlFlow) {
			Typecheck := __e.Get(1)
			_ = Typecheck
			tmp94082 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94083 := Call(__e, tmp94082, Typecheck, False)

			if True == tmp94083 {
				tmp94081 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp94081, MakeString("type error\n"))
				return

			} else {
				tmp94066 := MakeNative(func(__e *ControlFlow) {
					Eval := __e.Get(1)
					_ = Eval
					tmp94067 := MakeNative(func(__e *ControlFlow) {
						Type := __e.Get(1)
						_ = Type
						tmp94068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

						tmp94069 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						tmp94070 := Call(__e, PrimNS1Value(symns2_1value), symcn)

						tmp94071 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

						tmp94072 := Call(__e, tmp94071, Type, MakeString(""), symshen_4r)

						tmp94073 := Call(__e, tmp94070, MakeString(" : "), tmp94072)

						tmp94074 := Call(__e, tmp94069, Eval, tmp94073, symshen_4s)

						tmp94075 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

						tmp94076 := Call(__e, tmp94075)

						__e.TailApply(tmp94068, tmp94074, tmp94076)
						return

					}, 1)

					tmp94077 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pretty_1type)

					tmp94078 := Call(__e, tmp94077, Typecheck)

					__e.TailApply(tmp94067, tmp94078)
					return

				}, 1)

				tmp94079 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				tmp94080 := Call(__e, tmp94079, V3114)

				__e.TailApply(tmp94066, tmp94080)
				return

			}

		}, 1)

		tmp94084 := Call(__e, PrimNS1Value(symns2_1value), symshen_4typecheck)

		tmp94085 := Call(__e, tmp94084, V3114, V3115)

		__e.TailApply(tmp94064, tmp94085)
		return

	}, 2)

	tmp94086 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1evaluate, tmp94063)

	_ = tmp94086

	tmp94087 := MakeNative(func(__e *ControlFlow) {
		V3117 := __e.Get(1)
		_ = V3117
		tmp94088 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mult__subst)

		tmp94089 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94090 := Call(__e, tmp94089, symshen_4_dalphabet_d)

		tmp94091 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

		tmp94092 := Call(__e, tmp94091, V3117)

		__e.TailApply(tmp94088, tmp94090, tmp94092, V3117)
		return

	}, 1)

	tmp94093 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pretty_1type, tmp94087)

	_ = tmp94093

	tmp94094 := MakeNative(func(__e *ControlFlow) {
		V3123 := __e.Get(1)
		_ = V3123
		tmp94109 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pvar_2)

		tmp94110 := Call(__e, tmp94109, V3123)

		if True == tmp94110 {
			tmp94108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			__e.TailApply(tmp94108, V3123, Nil)
			return

		} else {
			tmp94106 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp94107 := Call(__e, tmp94106, V3123)

			if True == tmp94107 {
				tmp94097 := Call(__e, PrimNS1Value(symns2_1value), symunion)

				tmp94098 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

				tmp94099 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94100 := Call(__e, tmp94099, V3123)

				tmp94101 := Call(__e, tmp94098, tmp94100)

				tmp94102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract_1pvars)

				tmp94103 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94104 := Call(__e, tmp94103, V3123)

				tmp94105 := Call(__e, tmp94102, tmp94104)

				__e.TailApply(tmp94097, tmp94101, tmp94105)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp94111 := Call(__e, PrimNS1Value(symns2_1set), symshen_4extract_1pvars, tmp94094)

	_ = tmp94111

	tmp94112 := MakeNative(func(__e *ControlFlow) {
		V3131 := __e.Get(1)
		_ = V3131
		V3132 := __e.Get(2)
		_ = V3132
		V3133 := __e.Get(3)
		_ = V3133
		tmp94136 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94137 := Call(__e, tmp94136, Nil, V3131)

		if True == tmp94137 {
			__e.Return(V3133)
			return
		} else {
			tmp94134 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94135 := Call(__e, tmp94134, Nil, V3132)

			if True == tmp94135 {
				__e.Return(V3133)
				return
			} else {
				tmp94132 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94133 := Call(__e, tmp94132, V3131)

				var ifres94128 Obj

				if True == tmp94133 {
					tmp94130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp94131 := Call(__e, tmp94130, V3132)

					var ifres94129 Obj

					if True == tmp94131 {
						ifres94129 = True

					} else {
						ifres94129 = False

					}

					ifres94128 = ifres94129

				} else {
					ifres94128 = False

				}

				if True == ifres94128 {
					tmp94117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mult__subst)

					tmp94118 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94119 := Call(__e, tmp94118, V3131)

					tmp94120 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94121 := Call(__e, tmp94120, V3132)

					tmp94122 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp94123 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94124 := Call(__e, tmp94123, V3131)

					tmp94125 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94126 := Call(__e, tmp94125, V3132)

					tmp94127 := Call(__e, tmp94122, tmp94124, tmp94126, V3133)

					__e.TailApply(tmp94117, tmp94119, tmp94121, tmp94127)
					return

				} else {
					tmp94116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp94116, symshen_4mult__subst)
					return

				}

			}

		}

	}, 3)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4mult__subst, tmp94112)
	return

}, 0)
