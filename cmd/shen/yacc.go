package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp105362 := MakeNative(func(__e *ControlFlow) {
		V4724 := __e.Get(1)
		_ = V4724
		tmp105386 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp105387 := Call(__e, tmp105386, V4724)

		var ifres105374 Obj

		if True == tmp105387 {
			tmp105382 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105383 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105384 := Call(__e, tmp105383, V4724)

			tmp105385 := Call(__e, tmp105382, symdefcc, tmp105384)

			var ifres105376 Obj

			if True == tmp105385 {
				tmp105378 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105379 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105380 := Call(__e, tmp105379, V4724)

				tmp105381 := Call(__e, tmp105378, tmp105380)

				var ifres105377 Obj

				if True == tmp105381 {
					ifres105377 = True

				} else {
					ifres105377 = False

				}

				ifres105376 = ifres105377

			} else {
				ifres105376 = False

			}

			var ifres105375 Obj

			if True == ifres105376 {
				ifres105375 = True

			} else {
				ifres105375 = False

			}

			ifres105374 = ifres105375

		} else {
			ifres105374 = False

		}

		if True == ifres105374 {
			tmp105365 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc_1_6shen)

			tmp105366 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105367 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105368 := Call(__e, tmp105367, V4724)

			tmp105369 := Call(__e, tmp105366, tmp105368)

			tmp105370 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105371 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105372 := Call(__e, tmp105371, V4724)

			tmp105373 := Call(__e, tmp105370, tmp105372)

			__e.TailApply(tmp105365, tmp105369, tmp105373)
			return

		} else {
			tmp105364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp105364, symshen_4yacc)
			return

		}

	}, 1)

	tmp105388 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc, tmp105362)

	_ = tmp105388

	tmp105389 := MakeNative(func(__e *ControlFlow) {
		V4727 := __e.Get(1)
		_ = V4727
		V4728 := __e.Get(2)
		_ = V4728
		tmp105390 := MakeNative(func(__e *ControlFlow) {
			CCRules := __e.Get(1)
			_ = CCRules
			tmp105391 := MakeNative(func(__e *ControlFlow) {
				CCBody := __e.Get(1)
				_ = CCBody
				tmp105392 := MakeNative(func(__e *ControlFlow) {
					YaccCases := __e.Get(1)
					_ = YaccCases
					tmp105393 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105394 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105395 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105396 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105397 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4kill_1code)

					tmp105399 := Call(__e, tmp105398, YaccCases)

					tmp105400 := Call(__e, tmp105397, tmp105399, Nil)

					tmp105401 := Call(__e, tmp105396, sym_1_6, tmp105400)

					tmp105402 := Call(__e, tmp105395, symStream, tmp105401)

					tmp105403 := Call(__e, tmp105394, V4727, tmp105402)

					__e.TailApply(tmp105393, symdefine, tmp105403)
					return

				}, 1)

				tmp105404 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc__cases)

				tmp105405 := Call(__e, tmp105404, CCBody)

				__e.TailApply(tmp105392, tmp105405)
				return

			}, 1)

			tmp105406 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp105407 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp105408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cc__body)

				__e.TailApply(tmp105408, X)
				return

			}, 1)

			tmp105409 := Call(__e, tmp105406, tmp105407, CCRules)

			__e.TailApply(tmp105391, tmp105409)
			return

		}, 1)

		tmp105410 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

		tmp105411 := Call(__e, tmp105410, True, V4728, Nil)

		__e.TailApply(tmp105390, tmp105411)
		return

	}, 2)

	tmp105412 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc_1_6shen, tmp105389)

	_ = tmp105412

	tmp105413 := MakeNative(func(__e *ControlFlow) {
		V4730 := __e.Get(1)
		_ = V4730
		tmp105430 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		tmp105431 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

		tmp105432 := Call(__e, tmp105431, symkill, V4730)

		tmp105433 := Call(__e, tmp105430, tmp105432, MakeNumber(0))

		if True == tmp105433 {
			tmp105415 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105416 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105417 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105418 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105419 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105420 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105421 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105422 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105423 := Call(__e, tmp105422, symE, Nil)

			tmp105424 := Call(__e, tmp105421, symshen_4analyse_1kill, tmp105423)

			tmp105425 := Call(__e, tmp105420, tmp105424, Nil)

			tmp105426 := Call(__e, tmp105419, symE, tmp105425)

			tmp105427 := Call(__e, tmp105418, symlambda, tmp105426)

			tmp105428 := Call(__e, tmp105417, tmp105427, Nil)

			tmp105429 := Call(__e, tmp105416, V4730, tmp105428)

			__e.TailApply(tmp105415, symtrap_1error, tmp105429)
			return

		} else {
			__e.Return(V4730)
			return
		}

	}, 1)

	tmp105434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4kill_1code, tmp105413)

	_ = tmp105434

	tmp105435 := MakeNative(func(__e *ControlFlow) {
		tmp105436 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(tmp105436, MakeString("yacc kill"))
		return

	}, 0)

	tmp105437 := Call(__e, PrimNS1Value(symns2_1set), symkill, tmp105435)

	_ = tmp105437

	tmp105438 := MakeNative(func(__e *ControlFlow) {
		V4732 := __e.Get(1)
		_ = V4732
		tmp105439 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp105442 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105443 := Call(__e, tmp105442, String, MakeString("yacc kill"))

			if True == tmp105443 {
				tmp105441 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp105441)
				return

			} else {
				__e.Return(V4732)
				return
			}

		}, 1)

		tmp105444 := Call(__e, PrimNS1Value(symns2_1value), symerror_1to_1string)

		tmp105445 := Call(__e, tmp105444, V4732)

		__e.TailApply(tmp105439, tmp105445)
		return

	}, 1)

	tmp105446 := Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1kill, tmp105438)

	_ = tmp105446

	tmp105447 := MakeNative(func(__e *ControlFlow) {
		V4738 := __e.Get(1)
		_ = V4738
		V4739 := __e.Get(2)
		_ = V4739
		V4740 := __e.Get(3)
		_ = V4740
		tmp105490 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105491 := Call(__e, tmp105490, Nil, V4739)

		var ifres105486 Obj

		if True == tmp105491 {
			tmp105488 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105489 := Call(__e, tmp105488, Nil, V4740)

			var ifres105487 Obj

			if True == tmp105489 {
				ifres105487 = True

			} else {
				ifres105487 = False

			}

			ifres105486 = ifres105487

		} else {
			ifres105486 = False

		}

		if True == ifres105486 {
			__e.Return(Nil)
			return
		} else {
			tmp105484 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105485 := Call(__e, tmp105484, Nil, V4739)

			if True == tmp105485 {
				tmp105479 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

				tmp105481 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				tmp105482 := Call(__e, tmp105481, V4740)

				tmp105483 := Call(__e, tmp105480, V4738, tmp105482, Nil)

				__e.TailApply(tmp105479, tmp105483, Nil)
				return

			} else {
				tmp105477 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105478 := Call(__e, tmp105477, V4739)

				var ifres105471 Obj

				if True == tmp105478 {
					tmp105473 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105474 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105475 := Call(__e, tmp105474, V4739)

					tmp105476 := Call(__e, tmp105473, sym_k, tmp105475)

					var ifres105472 Obj

					if True == tmp105476 {
						ifres105472 = True

					} else {
						ifres105472 = False

					}

					ifres105471 = ifres105472

				} else {
					ifres105471 = False

				}

				if True == ifres105471 {
					tmp105462 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

					tmp105464 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					tmp105465 := Call(__e, tmp105464, V4740)

					tmp105466 := Call(__e, tmp105463, V4738, tmp105465, Nil)

					tmp105467 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

					tmp105468 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105469 := Call(__e, tmp105468, V4739)

					tmp105470 := Call(__e, tmp105467, V4738, tmp105469, Nil)

					__e.TailApply(tmp105462, tmp105466, tmp105470)
					return

				} else {
					tmp105460 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105461 := Call(__e, tmp105460, V4739)

					if True == tmp105461 {
						tmp105453 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rules)

						tmp105454 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105455 := Call(__e, tmp105454, V4739)

						tmp105456 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105457 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp105458 := Call(__e, tmp105457, V4739)

						tmp105459 := Call(__e, tmp105456, tmp105458, V4740)

						__e.TailApply(tmp105453, V4738, tmp105455, tmp105459)
						return

					} else {
						tmp105452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp105452, symshen_4split__cc__rules)
						return

					}

				}

			}

		}

	}, 3)

	tmp105492 := Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rules, tmp105447)

	_ = tmp105492

	tmp105493 := MakeNative(func(__e *ControlFlow) {
		V4748 := __e.Get(1)
		_ = V4748
		V4749 := __e.Get(2)
		_ = V4749
		V4750 := __e.Get(3)
		_ = V4750
		tmp105623 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp105624 := Call(__e, tmp105623, V4749)

		var ifres105603 Obj

		if True == tmp105624 {
			tmp105619 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105620 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105621 := Call(__e, tmp105620, V4749)

			tmp105622 := Call(__e, tmp105619, sym_h_a, tmp105621)

			var ifres105605 Obj

			if True == tmp105622 {
				tmp105615 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105616 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105617 := Call(__e, tmp105616, V4749)

				tmp105618 := Call(__e, tmp105615, tmp105617)

				var ifres105607 Obj

				if True == tmp105618 {
					tmp105609 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105610 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105611 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105612 := Call(__e, tmp105611, V4749)

					tmp105613 := Call(__e, tmp105610, tmp105612)

					tmp105614 := Call(__e, tmp105609, Nil, tmp105613)

					var ifres105608 Obj

					if True == tmp105614 {
						ifres105608 = True

					} else {
						ifres105608 = False

					}

					ifres105607 = ifres105608

				} else {
					ifres105607 = False

				}

				var ifres105606 Obj

				if True == ifres105607 {
					ifres105606 = True

				} else {
					ifres105606 = False

				}

				ifres105605 = ifres105606

			} else {
				ifres105605 = False

			}

			var ifres105604 Obj

			if True == ifres105605 {
				ifres105604 = True

			} else {
				ifres105604 = False

			}

			ifres105603 = ifres105604

		} else {
			ifres105603 = False

		}

		if True == ifres105603 {
			tmp105598 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105599 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			tmp105600 := Call(__e, tmp105599, V4750)

			tmp105601 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105602 := Call(__e, tmp105601, V4749)

			__e.TailApply(tmp105598, tmp105600, tmp105602)
			return

		} else {
			tmp105596 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105597 := Call(__e, tmp105596, V4749)

			var ifres105544 Obj

			if True == tmp105597 {
				tmp105592 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105593 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp105594 := Call(__e, tmp105593, V4749)

				tmp105595 := Call(__e, tmp105592, sym_h_a, tmp105594)

				var ifres105546 Obj

				if True == tmp105595 {
					tmp105588 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105589 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105590 := Call(__e, tmp105589, V4749)

					tmp105591 := Call(__e, tmp105588, tmp105590)

					var ifres105548 Obj

					if True == tmp105591 {
						tmp105582 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp105583 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105584 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105585 := Call(__e, tmp105584, V4749)

						tmp105586 := Call(__e, tmp105583, tmp105585)

						tmp105587 := Call(__e, tmp105582, tmp105586)

						var ifres105550 Obj

						if True == tmp105587 {
							tmp105574 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp105575 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp105576 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp105577 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp105578 := Call(__e, tmp105577, V4749)

							tmp105579 := Call(__e, tmp105576, tmp105578)

							tmp105580 := Call(__e, tmp105575, tmp105579)

							tmp105581 := Call(__e, tmp105574, symwhere, tmp105580)

							var ifres105552 Obj

							if True == tmp105581 {
								tmp105566 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp105567 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp105568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp105569 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp105570 := Call(__e, tmp105569, V4749)

								tmp105571 := Call(__e, tmp105568, tmp105570)

								tmp105572 := Call(__e, tmp105567, tmp105571)

								tmp105573 := Call(__e, tmp105566, tmp105572)

								var ifres105554 Obj

								if True == tmp105573 {
									tmp105556 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp105557 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp105558 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp105559 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp105560 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp105561 := Call(__e, tmp105560, V4749)

									tmp105562 := Call(__e, tmp105559, tmp105561)

									tmp105563 := Call(__e, tmp105558, tmp105562)

									tmp105564 := Call(__e, tmp105557, tmp105563)

									tmp105565 := Call(__e, tmp105556, Nil, tmp105564)

									var ifres105555 Obj

									if True == tmp105565 {
										ifres105555 = True

									} else {
										ifres105555 = False

									}

									ifres105554 = ifres105555

								} else {
									ifres105554 = False

								}

								var ifres105553 Obj

								if True == ifres105554 {
									ifres105553 = True

								} else {
									ifres105553 = False

								}

								ifres105552 = ifres105553

							} else {
								ifres105552 = False

							}

							var ifres105551 Obj

							if True == ifres105552 {
								ifres105551 = True

							} else {
								ifres105551 = False

							}

							ifres105550 = ifres105551

						} else {
							ifres105550 = False

						}

						var ifres105549 Obj

						if True == ifres105550 {
							ifres105549 = True

						} else {
							ifres105549 = False

						}

						ifres105548 = ifres105549

					} else {
						ifres105548 = False

					}

					var ifres105547 Obj

					if True == ifres105548 {
						ifres105547 = True

					} else {
						ifres105547 = False

					}

					ifres105546 = ifres105547

				} else {
					ifres105546 = False

				}

				var ifres105545 Obj

				if True == ifres105546 {
					ifres105545 = True

				} else {
					ifres105545 = False

				}

				ifres105544 = ifres105545

			} else {
				ifres105544 = False

			}

			if True == ifres105544 {
				tmp105521 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105522 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

				tmp105523 := Call(__e, tmp105522, V4750)

				tmp105524 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105525 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105526 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105527 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp105528 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105529 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105530 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105531 := Call(__e, tmp105530, V4749)

				tmp105532 := Call(__e, tmp105529, tmp105531)

				tmp105533 := Call(__e, tmp105528, tmp105532)

				tmp105534 := Call(__e, tmp105527, tmp105533)

				tmp105535 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105536 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp105537 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105538 := Call(__e, tmp105537, V4749)

				tmp105539 := Call(__e, tmp105536, tmp105538)

				tmp105540 := Call(__e, tmp105535, tmp105539, Nil)

				tmp105541 := Call(__e, tmp105526, tmp105534, tmp105540)

				tmp105542 := Call(__e, tmp105525, symwhere, tmp105541)

				tmp105543 := Call(__e, tmp105524, tmp105542, Nil)

				__e.TailApply(tmp105521, tmp105523, tmp105543)
				return

			} else {
				tmp105519 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105520 := Call(__e, tmp105519, Nil, V4749)

				if True == tmp105520 {
					tmp105508 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantic_1completion_1warning)

					tmp105509 := Call(__e, tmp105508, V4748, V4750)

					_ = tmp105509

					tmp105510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

					tmp105511 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105512 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105513 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

					tmp105514 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					tmp105515 := Call(__e, tmp105514, V4750)

					tmp105516 := Call(__e, tmp105513, tmp105515)

					tmp105517 := Call(__e, tmp105512, tmp105516, Nil)

					tmp105518 := Call(__e, tmp105511, sym_h_a, tmp105517)

					__e.TailApply(tmp105510, V4748, tmp105518, V4750)
					return

				} else {
					tmp105506 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105507 := Call(__e, tmp105506, V4749)

					if True == tmp105507 {
						tmp105499 := Call(__e, PrimNS1Value(symns2_1value), symshen_4split__cc__rule)

						tmp105500 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105501 := Call(__e, tmp105500, V4749)

						tmp105502 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105503 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp105504 := Call(__e, tmp105503, V4749)

						tmp105505 := Call(__e, tmp105502, tmp105504, V4750)

						__e.TailApply(tmp105499, V4748, tmp105501, tmp105505)
						return

					} else {
						tmp105498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp105498, symshen_4split__cc__rule)
						return

					}

				}

			}

		}

	}, 3)

	tmp105625 := Call(__e, PrimNS1Value(symns2_1set), symshen_4split__cc__rule, tmp105493)

	_ = tmp105625

	tmp105626 := MakeNative(func(__e *ControlFlow) {
		V4761 := __e.Get(1)
		_ = V4761
		V4762 := __e.Get(2)
		_ = V4762
		tmp105645 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105646 := Call(__e, tmp105645, True, V4761)

		if True == tmp105646 {
			tmp105628 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp105629 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp105630 := Call(__e, tmp105629)

			tmp105631 := Call(__e, tmp105628, MakeString("warning: "), tmp105630)

			_ = tmp105631

			tmp105632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

			tmp105633 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp105634 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp105635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				tmp105636 := Call(__e, tmp105635, X, MakeString(" "), symshen_4a)

				tmp105637 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp105638 := Call(__e, tmp105637)

				__e.TailApply(tmp105634, tmp105636, tmp105638)
				return

			}, 1)

			tmp105639 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

			tmp105640 := Call(__e, tmp105639, V4762)

			tmp105641 := Call(__e, tmp105632, tmp105633, tmp105640)

			_ = tmp105641

			tmp105642 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp105643 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp105644 := Call(__e, tmp105643)

			__e.TailApply(tmp105642, MakeString("has no semantics.\n"), tmp105644)
			return

		} else {
			__e.Return(symshen_4skip)
			return
		}

	}, 2)

	tmp105647 := Call(__e, PrimNS1Value(symns2_1set), symshen_4semantic_1completion_1warning, tmp105626)

	_ = tmp105647

	tmp105648 := MakeNative(func(__e *ControlFlow) {
		V4764 := __e.Get(1)
		_ = V4764
		tmp105701 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105702 := Call(__e, tmp105701, Nil, V4764)

		if True == tmp105702 {
			__e.Return(Nil)
			return
		} else {
			tmp105699 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105700 := Call(__e, tmp105699, V4764)

			var ifres105687 Obj

			if True == tmp105700 {
				tmp105695 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105696 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105697 := Call(__e, tmp105696, V4764)

				tmp105698 := Call(__e, tmp105695, Nil, tmp105697)

				var ifres105689 Obj

				if True == tmp105698 {
					tmp105691 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					tmp105692 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105693 := Call(__e, tmp105692, V4764)

					tmp105694 := Call(__e, tmp105691, tmp105693)

					var ifres105690 Obj

					if True == tmp105694 {
						ifres105690 = True

					} else {
						ifres105690 = False

					}

					ifres105689 = ifres105690

				} else {
					ifres105689 = False

				}

				var ifres105688 Obj

				if True == ifres105689 {
					ifres105688 = True

				} else {
					ifres105688 = False

				}

				ifres105687 = ifres105688

			} else {
				ifres105687 = False

			}

			if True == ifres105687 {
				tmp105686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				__e.TailApply(tmp105686, V4764)
				return

			} else {
				tmp105684 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105685 := Call(__e, tmp105684, V4764)

				var ifres105678 Obj

				if True == tmp105685 {
					tmp105680 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					tmp105681 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105682 := Call(__e, tmp105681, V4764)

					tmp105683 := Call(__e, tmp105680, tmp105682)

					var ifres105679 Obj

					if True == tmp105683 {
						ifres105679 = True

					} else {
						ifres105679 = False

					}

					ifres105678 = ifres105679

				} else {
					ifres105678 = False

				}

				if True == ifres105678 {
					tmp105667 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105668 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105669 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105670 := Call(__e, tmp105669, V4764)

					tmp105671 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

					tmp105673 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105674 := Call(__e, tmp105673, V4764)

					tmp105675 := Call(__e, tmp105672, tmp105674)

					tmp105676 := Call(__e, tmp105671, tmp105675, Nil)

					tmp105677 := Call(__e, tmp105668, tmp105670, tmp105676)

					__e.TailApply(tmp105667, symappend, tmp105677)
					return

				} else {
					tmp105665 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105666 := Call(__e, tmp105665, V4764)

					if True == tmp105666 {
						tmp105654 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105655 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105656 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp105657 := Call(__e, tmp105656, V4764)

						tmp105658 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105659 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default__semantics)

						tmp105660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105661 := Call(__e, tmp105660, V4764)

						tmp105662 := Call(__e, tmp105659, tmp105661)

						tmp105663 := Call(__e, tmp105658, tmp105662, Nil)

						tmp105664 := Call(__e, tmp105655, tmp105657, tmp105663)

						__e.TailApply(tmp105654, symcons, tmp105664)
						return

					} else {
						tmp105653 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp105653, symshen_4default__semantics)
						return

					}

				}

			}

		}

	}, 1)

	tmp105703 := Call(__e, PrimNS1Value(symns2_1set), symshen_4default__semantics, tmp105648)

	_ = tmp105703

	tmp105704 := MakeNative(func(__e *ControlFlow) {
		V4766 := __e.Get(1)
		_ = V4766
		tmp105725 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp105726 := Call(__e, tmp105725, V4766)

		if True == tmp105726 {
			tmp105707 := MakeNative(func(__e *ControlFlow) {
				Cs := __e.Get(1)
				_ = Cs
				tmp105716 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105717 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp105718 := Call(__e, tmp105717, Cs)

				tmp105719 := Call(__e, tmp105716, tmp105718, MakeString("<"))

				if True == tmp105719 {
					tmp105710 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp105711 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105712 := Call(__e, PrimNS1Value(symns2_1value), symreverse)

					tmp105713 := Call(__e, tmp105712, Cs)

					tmp105714 := Call(__e, tmp105711, tmp105713)

					tmp105715 := Call(__e, tmp105710, tmp105714, MakeString(">"))

					if True == tmp105715 {
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

			tmp105720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1pathname)

			tmp105721 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

			tmp105722 := Call(__e, tmp105721, V4766)

			tmp105723 := Call(__e, tmp105720, tmp105722)

			tmp105724 := Call(__e, tmp105707, tmp105723)

			if True == tmp105724 {
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

	tmp105727 := Call(__e, PrimNS1Value(symns2_1set), symshen_4grammar__symbol_2, tmp105704)

	_ = tmp105727

	tmp105728 := MakeNative(func(__e *ControlFlow) {
		V4768 := __e.Get(1)
		_ = V4768
		tmp105771 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp105772 := Call(__e, tmp105771, V4768)

		var ifres105765 Obj

		if True == tmp105772 {
			tmp105767 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105769 := Call(__e, tmp105768, V4768)

			tmp105770 := Call(__e, tmp105767, Nil, tmp105769)

			var ifres105766 Obj

			if True == tmp105770 {
				ifres105766 = True

			} else {
				ifres105766 = False

			}

			ifres105765 = ifres105766

		} else {
			ifres105765 = False

		}

		if True == ifres105765 {
			tmp105764 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp105764, V4768)
			return

		} else {
			tmp105762 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105763 := Call(__e, tmp105762, V4768)

			if True == tmp105763 {
				tmp105732 := MakeNative(func(__e *ControlFlow) {
					P := __e.Get(1)
					_ = P
					tmp105733 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105734 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105735 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105736 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105737 := Call(__e, tmp105736, V4768)

					tmp105738 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105739 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105742 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105743 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105744 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105745 := Call(__e, tmp105744, symfail, Nil)

					tmp105746 := Call(__e, tmp105743, tmp105745, Nil)

					tmp105747 := Call(__e, tmp105742, P, tmp105746)

					tmp105748 := Call(__e, tmp105741, sym_a, tmp105747)

					tmp105749 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105750 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc__cases)

					tmp105751 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105752 := Call(__e, tmp105751, V4768)

					tmp105753 := Call(__e, tmp105750, tmp105752)

					tmp105754 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105755 := Call(__e, tmp105754, P, Nil)

					tmp105756 := Call(__e, tmp105749, tmp105753, tmp105755)

					tmp105757 := Call(__e, tmp105740, tmp105748, tmp105756)

					tmp105758 := Call(__e, tmp105739, symif, tmp105757)

					tmp105759 := Call(__e, tmp105738, tmp105758, Nil)

					tmp105760 := Call(__e, tmp105735, tmp105737, tmp105759)

					tmp105761 := Call(__e, tmp105734, P, tmp105760)

					__e.TailApply(tmp105733, symlet, tmp105761)
					return

				}, 1)

				__e.TailApply(tmp105732, symYaccParse)
				return

			} else {
				tmp105731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp105731, symshen_4yacc__cases)
				return

			}

		}

	}, 1)

	tmp105773 := Call(__e, PrimNS1Value(symns2_1set), symshen_4yacc__cases, tmp105728)

	_ = tmp105773

	tmp105774 := MakeNative(func(__e *ControlFlow) {
		V4770 := __e.Get(1)
		_ = V4770
		tmp105798 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp105799 := Call(__e, tmp105798, V4770)

		var ifres105784 Obj

		if True == tmp105799 {
			tmp105794 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105795 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105796 := Call(__e, tmp105795, V4770)

			tmp105797 := Call(__e, tmp105794, tmp105796)

			var ifres105786 Obj

			if True == tmp105797 {
				tmp105788 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105789 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105790 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp105791 := Call(__e, tmp105790, V4770)

				tmp105792 := Call(__e, tmp105789, tmp105791)

				tmp105793 := Call(__e, tmp105788, Nil, tmp105792)

				var ifres105787 Obj

				if True == tmp105793 {
					ifres105787 = True

				} else {
					ifres105787 = False

				}

				ifres105786 = ifres105787

			} else {
				ifres105786 = False

			}

			var ifres105785 Obj

			if True == ifres105786 {
				ifres105785 = True

			} else {
				ifres105785 = False

			}

			ifres105784 = ifres105785

		} else {
			ifres105784 = False

		}

		if True == ifres105784 {
			tmp105777 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

			tmp105778 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105779 := Call(__e, tmp105778, V4770)

			tmp105780 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105781 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105782 := Call(__e, tmp105781, V4770)

			tmp105783 := Call(__e, tmp105780, tmp105782)

			__e.TailApply(tmp105777, tmp105779, symStream, tmp105783)
			return

		} else {
			tmp105776 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp105776, symshen_4cc__body)
			return

		}

	}, 1)

	tmp105800 := Call(__e, PrimNS1Value(symns2_1set), symshen_4cc__body, tmp105774)

	_ = tmp105800

	tmp105801 := MakeNative(func(__e *ControlFlow) {
		V4774 := __e.Get(1)
		_ = V4774
		V4775 := __e.Get(2)
		_ = V4775
		V4776 := __e.Get(3)
		_ = V4776
		tmp105929 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp105930 := Call(__e, tmp105929, Nil, V4774)

		var ifres105895 Obj

		if True == tmp105930 {
			tmp105927 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp105928 := Call(__e, tmp105927, V4776)

			var ifres105897 Obj

			if True == tmp105928 {
				tmp105923 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp105924 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp105925 := Call(__e, tmp105924, V4776)

				tmp105926 := Call(__e, tmp105923, symwhere, tmp105925)

				var ifres105899 Obj

				if True == tmp105926 {
					tmp105919 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp105920 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp105921 := Call(__e, tmp105920, V4776)

					tmp105922 := Call(__e, tmp105919, tmp105921)

					var ifres105901 Obj

					if True == tmp105922 {
						tmp105913 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp105914 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp105916 := Call(__e, tmp105915, V4776)

						tmp105917 := Call(__e, tmp105914, tmp105916)

						tmp105918 := Call(__e, tmp105913, tmp105917)

						var ifres105903 Obj

						if True == tmp105918 {
							tmp105905 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp105906 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp105907 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp105908 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp105909 := Call(__e, tmp105908, V4776)

							tmp105910 := Call(__e, tmp105907, tmp105909)

							tmp105911 := Call(__e, tmp105906, tmp105910)

							tmp105912 := Call(__e, tmp105905, Nil, tmp105911)

							var ifres105904 Obj

							if True == tmp105912 {
								ifres105904 = True

							} else {
								ifres105904 = False

							}

							ifres105903 = ifres105904

						} else {
							ifres105903 = False

						}

						var ifres105902 Obj

						if True == ifres105903 {
							ifres105902 = True

						} else {
							ifres105902 = False

						}

						ifres105901 = ifres105902

					} else {
						ifres105901 = False

					}

					var ifres105900 Obj

					if True == ifres105901 {
						ifres105900 = True

					} else {
						ifres105900 = False

					}

					ifres105899 = ifres105900

				} else {
					ifres105899 = False

				}

				var ifres105898 Obj

				if True == ifres105899 {
					ifres105898 = True

				} else {
					ifres105898 = False

				}

				ifres105897 = ifres105898

			} else {
				ifres105897 = False

			}

			var ifres105896 Obj

			if True == ifres105897 {
				ifres105896 = True

			} else {
				ifres105896 = False

			}

			ifres105895 = ifres105896

		} else {
			ifres105895 = False

		}

		if True == ifres105895 {
			tmp105862 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105863 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

			tmp105865 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105866 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105867 := Call(__e, tmp105866, V4776)

			tmp105868 := Call(__e, tmp105865, tmp105867)

			tmp105869 := Call(__e, tmp105864, tmp105868)

			tmp105870 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105871 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105872 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105873 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105874 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105875 := Call(__e, tmp105874, V4775, Nil)

			tmp105876 := Call(__e, tmp105873, symhd, tmp105875)

			tmp105877 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105878 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

			tmp105879 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp105880 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105881 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp105882 := Call(__e, tmp105881, V4776)

			tmp105883 := Call(__e, tmp105880, tmp105882)

			tmp105884 := Call(__e, tmp105879, tmp105883)

			tmp105885 := Call(__e, tmp105878, tmp105884)

			tmp105886 := Call(__e, tmp105877, tmp105885, Nil)

			tmp105887 := Call(__e, tmp105872, tmp105876, tmp105886)

			tmp105888 := Call(__e, tmp105871, symshen_4pair, tmp105887)

			tmp105889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp105891 := Call(__e, tmp105890, symfail, Nil)

			tmp105892 := Call(__e, tmp105889, tmp105891, Nil)

			tmp105893 := Call(__e, tmp105870, tmp105888, tmp105892)

			tmp105894 := Call(__e, tmp105863, tmp105869, tmp105893)

			__e.TailApply(tmp105862, symif, tmp105894)
			return

		} else {
			tmp105860 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp105861 := Call(__e, tmp105860, Nil, V4774)

			if True == tmp105861 {
				tmp105849 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105851 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105852 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105853 := Call(__e, tmp105852, V4775, Nil)

				tmp105854 := Call(__e, tmp105851, symhd, tmp105853)

				tmp105855 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

				tmp105857 := Call(__e, tmp105856, V4776)

				tmp105858 := Call(__e, tmp105855, tmp105857, Nil)

				tmp105859 := Call(__e, tmp105850, tmp105854, tmp105858)

				__e.TailApply(tmp105849, symshen_4pair, tmp105859)
				return

			} else {
				tmp105847 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp105848 := Call(__e, tmp105847, V4774)

				if True == tmp105848 {
					tmp105843 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

					tmp105844 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp105845 := Call(__e, tmp105844, V4774)

					tmp105846 := Call(__e, tmp105843, tmp105845)

					if True == tmp105846 {
						tmp105842 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__descent)

						__e.TailApply(tmp105842, V4774, V4775, V4776)
						return

					} else {
						tmp105838 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

						tmp105839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp105840 := Call(__e, tmp105839, V4774)

						tmp105841 := Call(__e, tmp105838, tmp105840)

						if True == tmp105841 {
							tmp105837 := Call(__e, PrimNS1Value(symns2_1value), symshen_4variable_1match)

							__e.TailApply(tmp105837, V4774, V4775, V4776)
							return

						} else {
							tmp105833 := Call(__e, PrimNS1Value(symns2_1value), symshen_4jump__stream_2)

							tmp105834 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp105835 := Call(__e, tmp105834, V4774)

							tmp105836 := Call(__e, tmp105833, tmp105835)

							if True == tmp105836 {
								tmp105832 := Call(__e, PrimNS1Value(symns2_1value), symshen_4jump__stream)

								__e.TailApply(tmp105832, V4774, V4775, V4776)
								return

							} else {
								tmp105828 := Call(__e, PrimNS1Value(symns2_1value), symshen_4terminal_2)

								tmp105829 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp105830 := Call(__e, tmp105829, V4774)

								tmp105831 := Call(__e, tmp105828, tmp105830)

								if True == tmp105831 {
									tmp105827 := Call(__e, PrimNS1Value(symns2_1value), symshen_4check__stream)

									__e.TailApply(tmp105827, V4774, V4775, V4776)
									return

								} else {
									tmp105823 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

									tmp105824 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp105825 := Call(__e, tmp105824, V4774)

									tmp105826 := Call(__e, tmp105823, tmp105825)

									if True == tmp105826 {
										tmp105816 := Call(__e, PrimNS1Value(symns2_1value), symshen_4list_1stream)

										tmp105817 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decons)

										tmp105818 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp105819 := Call(__e, tmp105818, V4774)

										tmp105820 := Call(__e, tmp105817, tmp105819)

										tmp105821 := Call(__e, PrimNS1Value(symns2_1value), symtl)

										tmp105822 := Call(__e, tmp105821, V4774)

										__e.TailApply(tmp105816, tmp105820, tmp105822, V4775, V4776)
										return

									} else {
										tmp105811 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

										tmp105812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

										tmp105813 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp105814 := Call(__e, tmp105813, V4774)

										tmp105815 := Call(__e, tmp105812, tmp105814, MakeString(" is not legal syntax\n"), symshen_4a)

										__e.TailApply(tmp105811, tmp105815)
										return

									}

								}

							}

						}

					}

				} else {
					tmp105805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp105805, symshen_4syntax)
					return

				}

			}

		}

	}, 3)

	tmp105931 := Call(__e, PrimNS1Value(symns2_1set), symshen_4syntax, tmp105801)

	_ = tmp105931

	tmp105932 := MakeNative(func(__e *ControlFlow) {
		V4781 := __e.Get(1)
		_ = V4781
		V4782 := __e.Get(2)
		_ = V4782
		V4783 := __e.Get(3)
		_ = V4783
		V4784 := __e.Get(4)
		_ = V4784
		tmp105933 := MakeNative(func(__e *ControlFlow) {
			Test := __e.Get(1)
			_ = Test
			tmp105934 := MakeNative(func(__e *ControlFlow) {
				Placeholder := __e.Get(1)
				_ = Placeholder
				tmp105935 := MakeNative(func(__e *ControlFlow) {
					RunOn := __e.Get(1)
					_ = RunOn
					tmp105936 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						tmp105937 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105938 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105939 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105940 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105941 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp105942 := Call(__e, tmp105941, symfail, Nil)

						tmp105943 := Call(__e, tmp105940, tmp105942, Nil)

						tmp105944 := Call(__e, tmp105939, Action, tmp105943)

						tmp105945 := Call(__e, tmp105938, Test, tmp105944)

						__e.TailApply(tmp105937, symif, tmp105945)
						return

					}, 1)

					tmp105946 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1runon)

					tmp105947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

					tmp105948 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105949 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105950 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105951 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105952 := Call(__e, tmp105951, V4783, Nil)

					tmp105953 := Call(__e, tmp105950, symshen_4hdhd, tmp105952)

					tmp105954 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105955 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105956 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp105957 := Call(__e, tmp105956, V4783, Nil)

					tmp105958 := Call(__e, tmp105955, symshen_4hdtl, tmp105957)

					tmp105959 := Call(__e, tmp105954, tmp105958, Nil)

					tmp105960 := Call(__e, tmp105949, tmp105953, tmp105959)

					tmp105961 := Call(__e, tmp105948, symshen_4pair, tmp105960)

					tmp105962 := Call(__e, tmp105947, V4781, tmp105961, Placeholder)

					tmp105963 := Call(__e, tmp105946, RunOn, Placeholder, tmp105962)

					__e.TailApply(tmp105936, tmp105963)
					return

				}, 1)

				tmp105964 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				tmp105965 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105966 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105967 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105969 := Call(__e, tmp105968, V4783, Nil)

				tmp105970 := Call(__e, tmp105967, symshen_4tlhd, tmp105969)

				tmp105971 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105972 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105973 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp105974 := Call(__e, tmp105973, V4783, Nil)

				tmp105975 := Call(__e, tmp105972, symshen_4hdtl, tmp105974)

				tmp105976 := Call(__e, tmp105971, tmp105975, Nil)

				tmp105977 := Call(__e, tmp105966, tmp105970, tmp105976)

				tmp105978 := Call(__e, tmp105965, symshen_4pair, tmp105977)

				tmp105979 := Call(__e, tmp105964, V4782, tmp105978, V4784)

				__e.TailApply(tmp105935, tmp105979)
				return

			}, 1)

			tmp105980 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp105981 := Call(__e, tmp105980, symshen_4place)

			__e.TailApply(tmp105934, tmp105981)
			return

		}, 1)

		tmp105982 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105983 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105985 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105986 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105988 := Call(__e, tmp105987, V4783, Nil)

		tmp105989 := Call(__e, tmp105986, symhd, tmp105988)

		tmp105990 := Call(__e, tmp105985, tmp105989, Nil)

		tmp105991 := Call(__e, tmp105984, symcons_2, tmp105990)

		tmp105992 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105993 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105994 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105995 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105996 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp105997 := Call(__e, tmp105996, V4783, Nil)

		tmp105998 := Call(__e, tmp105995, symshen_4hdhd, tmp105997)

		tmp105999 := Call(__e, tmp105994, tmp105998, Nil)

		tmp106000 := Call(__e, tmp105993, symcons_2, tmp105999)

		tmp106001 := Call(__e, tmp105992, tmp106000, Nil)

		tmp106002 := Call(__e, tmp105983, tmp105991, tmp106001)

		tmp106003 := Call(__e, tmp105982, symand, tmp106002)

		__e.TailApply(tmp105933, tmp106003)
		return

	}, 4)

	tmp106004 := Call(__e, PrimNS1Value(symns2_1set), symshen_4list_1stream, tmp105932)

	_ = tmp106004

	tmp106005 := MakeNative(func(__e *ControlFlow) {
		V4786 := __e.Get(1)
		_ = V4786
		tmp106098 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106099 := Call(__e, tmp106098, V4786)

		var ifres106058 Obj

		if True == tmp106099 {
			tmp106094 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp106095 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106096 := Call(__e, tmp106095, V4786)

			tmp106097 := Call(__e, tmp106094, symcons, tmp106096)

			var ifres106060 Obj

			if True == tmp106097 {
				tmp106090 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp106091 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106092 := Call(__e, tmp106091, V4786)

				tmp106093 := Call(__e, tmp106090, tmp106092)

				var ifres106062 Obj

				if True == tmp106093 {
					tmp106084 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp106085 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106086 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106087 := Call(__e, tmp106086, V4786)

					tmp106088 := Call(__e, tmp106085, tmp106087)

					tmp106089 := Call(__e, tmp106084, tmp106088)

					var ifres106064 Obj

					if True == tmp106089 {
						tmp106076 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp106077 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp106078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106079 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106080 := Call(__e, tmp106079, V4786)

						tmp106081 := Call(__e, tmp106078, tmp106080)

						tmp106082 := Call(__e, tmp106077, tmp106081)

						tmp106083 := Call(__e, tmp106076, Nil, tmp106082)

						var ifres106066 Obj

						if True == tmp106083 {
							tmp106068 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp106069 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106071 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106072 := Call(__e, tmp106071, V4786)

							tmp106073 := Call(__e, tmp106070, tmp106072)

							tmp106074 := Call(__e, tmp106069, tmp106073)

							tmp106075 := Call(__e, tmp106068, Nil, tmp106074)

							var ifres106067 Obj

							if True == tmp106075 {
								ifres106067 = True

							} else {
								ifres106067 = False

							}

							ifres106066 = ifres106067

						} else {
							ifres106066 = False

						}

						var ifres106065 Obj

						if True == ifres106066 {
							ifres106065 = True

						} else {
							ifres106065 = False

						}

						ifres106064 = ifres106065

					} else {
						ifres106064 = False

					}

					var ifres106063 Obj

					if True == ifres106064 {
						ifres106063 = True

					} else {
						ifres106063 = False

					}

					ifres106062 = ifres106063

				} else {
					ifres106062 = False

				}

				var ifres106061 Obj

				if True == ifres106062 {
					ifres106061 = True

				} else {
					ifres106061 = False

				}

				ifres106060 = ifres106061

			} else {
				ifres106060 = False

			}

			var ifres106059 Obj

			if True == ifres106060 {
				ifres106059 = True

			} else {
				ifres106059 = False

			}

			ifres106058 = ifres106059

		} else {
			ifres106058 = False

		}

		if True == ifres106058 {
			tmp106053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106054 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106055 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp106056 := Call(__e, tmp106055, V4786)

			tmp106057 := Call(__e, tmp106054, tmp106056)

			__e.TailApply(tmp106053, tmp106057, Nil)
			return

		} else {
			tmp106051 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106052 := Call(__e, tmp106051, V4786)

			var ifres106021 Obj

			if True == tmp106052 {
				tmp106047 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp106048 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp106049 := Call(__e, tmp106048, V4786)

				tmp106050 := Call(__e, tmp106047, symcons, tmp106049)

				var ifres106023 Obj

				if True == tmp106050 {
					tmp106043 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp106044 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106045 := Call(__e, tmp106044, V4786)

					tmp106046 := Call(__e, tmp106043, tmp106045)

					var ifres106025 Obj

					if True == tmp106046 {
						tmp106037 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp106038 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106039 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106040 := Call(__e, tmp106039, V4786)

						tmp106041 := Call(__e, tmp106038, tmp106040)

						tmp106042 := Call(__e, tmp106037, tmp106041)

						var ifres106027 Obj

						if True == tmp106042 {
							tmp106029 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp106030 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106031 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106032 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106033 := Call(__e, tmp106032, V4786)

							tmp106034 := Call(__e, tmp106031, tmp106033)

							tmp106035 := Call(__e, tmp106030, tmp106034)

							tmp106036 := Call(__e, tmp106029, Nil, tmp106035)

							var ifres106028 Obj

							if True == tmp106036 {
								ifres106028 = True

							} else {
								ifres106028 = False

							}

							ifres106027 = ifres106028

						} else {
							ifres106027 = False

						}

						var ifres106026 Obj

						if True == ifres106027 {
							ifres106026 = True

						} else {
							ifres106026 = False

						}

						ifres106025 = ifres106026

					} else {
						ifres106025 = False

					}

					var ifres106024 Obj

					if True == ifres106025 {
						ifres106024 = True

					} else {
						ifres106024 = False

					}

					ifres106023 = ifres106024

				} else {
					ifres106023 = False

				}

				var ifres106022 Obj

				if True == ifres106023 {
					ifres106022 = True

				} else {
					ifres106022 = False

				}

				ifres106021 = ifres106022

			} else {
				ifres106021 = False

			}

			if True == ifres106021 {
				tmp106008 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106009 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp106010 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106011 := Call(__e, tmp106010, V4786)

				tmp106012 := Call(__e, tmp106009, tmp106011)

				tmp106013 := Call(__e, PrimNS1Value(symns2_1value), symshen_4decons)

				tmp106014 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp106015 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106016 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106017 := Call(__e, tmp106016, V4786)

				tmp106018 := Call(__e, tmp106015, tmp106017)

				tmp106019 := Call(__e, tmp106014, tmp106018)

				tmp106020 := Call(__e, tmp106013, tmp106019)

				__e.TailApply(tmp106008, tmp106012, tmp106020)
				return

			} else {
				__e.Return(V4786)
				return
			}

		}

	}, 1)

	tmp106100 := Call(__e, PrimNS1Value(symns2_1set), symshen_4decons, tmp106005)

	_ = tmp106100

	tmp106101 := MakeNative(func(__e *ControlFlow) {
		V4801 := __e.Get(1)
		_ = V4801
		V4802 := __e.Get(2)
		_ = V4802
		V4803 := __e.Get(3)
		_ = V4803
		tmp106149 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106150 := Call(__e, tmp106149, V4803)

		var ifres106109 Obj

		if True == tmp106150 {
			tmp106145 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp106146 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106147 := Call(__e, tmp106146, V4803)

			tmp106148 := Call(__e, tmp106145, symshen_4pair, tmp106147)

			var ifres106111 Obj

			if True == tmp106148 {
				tmp106141 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp106142 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106143 := Call(__e, tmp106142, V4803)

				tmp106144 := Call(__e, tmp106141, tmp106143)

				var ifres106113 Obj

				if True == tmp106144 {
					tmp106135 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp106136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106137 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106138 := Call(__e, tmp106137, V4803)

					tmp106139 := Call(__e, tmp106136, tmp106138)

					tmp106140 := Call(__e, tmp106135, tmp106139)

					var ifres106115 Obj

					if True == tmp106140 {
						tmp106127 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp106128 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106129 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106130 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp106131 := Call(__e, tmp106130, V4803)

						tmp106132 := Call(__e, tmp106129, tmp106131)

						tmp106133 := Call(__e, tmp106128, tmp106132)

						tmp106134 := Call(__e, tmp106127, Nil, tmp106133)

						var ifres106117 Obj

						if True == tmp106134 {
							tmp106119 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp106120 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp106121 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106122 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp106123 := Call(__e, tmp106122, V4803)

							tmp106124 := Call(__e, tmp106121, tmp106123)

							tmp106125 := Call(__e, tmp106120, tmp106124)

							tmp106126 := Call(__e, tmp106119, tmp106125, V4802)

							var ifres106118 Obj

							if True == tmp106126 {
								ifres106118 = True

							} else {
								ifres106118 = False

							}

							ifres106117 = ifres106118

						} else {
							ifres106117 = False

						}

						var ifres106116 Obj

						if True == ifres106117 {
							ifres106116 = True

						} else {
							ifres106116 = False

						}

						ifres106115 = ifres106116

					} else {
						ifres106115 = False

					}

					var ifres106114 Obj

					if True == ifres106115 {
						ifres106114 = True

					} else {
						ifres106114 = False

					}

					ifres106113 = ifres106114

				} else {
					ifres106113 = False

				}

				var ifres106112 Obj

				if True == ifres106113 {
					ifres106112 = True

				} else {
					ifres106112 = False

				}

				ifres106111 = ifres106112

			} else {
				ifres106111 = False

			}

			var ifres106110 Obj

			if True == ifres106111 {
				ifres106110 = True

			} else {
				ifres106110 = False

			}

			ifres106109 = ifres106110

		} else {
			ifres106109 = False

		}

		if True == ifres106109 {
			__e.Return(V4801)
			return
		} else {
			tmp106107 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106108 := Call(__e, tmp106107, V4803)

			if True == tmp106108 {
				tmp106104 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				tmp106105 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					tmp106106 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1runon)

					__e.TailApply(tmp106106, V4801, V4802, Z)
					return

				}, 1)

				__e.TailApply(tmp106104, tmp106105, V4803)
				return

			} else {
				__e.Return(V4803)
				return
			}

		}

	}, 3)

	tmp106151 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1runon, tmp106101)

	_ = tmp106151

	tmp106152 := MakeNative(func(__e *ControlFlow) {
		V4809 := __e.Get(1)
		_ = V4809
		tmp106161 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		tmp106162 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp106163 := Call(__e, tmp106162, MakeString("."), V4809)

		tmp106164 := Call(__e, tmp106161, tmp106163)

		if True == tmp106164 {
			__e.Return(V4809)
			return
		} else {
			tmp106159 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106160 := Call(__e, tmp106159, V4809)

			if True == tmp106160 {
				tmp106156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4strip_1pathname)

				tmp106157 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106158 := Call(__e, tmp106157, V4809)

				__e.TailApply(tmp106156, tmp106158)
				return

			} else {
				tmp106155 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp106155, symshen_4strip_1pathname)
				return

			}

		}

	}, 1)

	tmp106165 := Call(__e, PrimNS1Value(symns2_1set), symshen_4strip_1pathname, tmp106152)

	_ = tmp106165

	tmp106166 := MakeNative(func(__e *ControlFlow) {
		V4813 := __e.Get(1)
		_ = V4813
		V4814 := __e.Get(2)
		_ = V4814
		V4815 := __e.Get(3)
		_ = V4815
		tmp106223 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106224 := Call(__e, tmp106223, V4813)

		if True == tmp106224 {
			tmp106169 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp106170 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp106171 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp106172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106173 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106174 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

						tmp106175 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp106176 := Call(__e, tmp106175, V4813)

						tmp106177 := Call(__e, tmp106174, symParse__, tmp106176)

						tmp106178 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106179 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106180 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106181 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106182 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106183 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106184 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106185 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106186 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106187 := Call(__e, tmp106186, symfail, Nil)

						tmp106188 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106189 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

						tmp106190 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp106191 := Call(__e, tmp106190, V4813)

						tmp106192 := Call(__e, tmp106189, symParse__, tmp106191)

						tmp106193 := Call(__e, tmp106188, tmp106192, Nil)

						tmp106194 := Call(__e, tmp106185, tmp106187, tmp106193)

						tmp106195 := Call(__e, tmp106184, sym_a, tmp106194)

						tmp106196 := Call(__e, tmp106183, tmp106195, Nil)

						tmp106197 := Call(__e, tmp106182, symnot, tmp106196)

						tmp106198 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106199 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106200 := Call(__e, tmp106199, Else, Nil)

						tmp106201 := Call(__e, tmp106198, Action, tmp106200)

						tmp106202 := Call(__e, tmp106181, tmp106197, tmp106201)

						tmp106203 := Call(__e, tmp106180, symif, tmp106202)

						tmp106204 := Call(__e, tmp106179, tmp106203, Nil)

						tmp106205 := Call(__e, tmp106178, Test, tmp106204)

						tmp106206 := Call(__e, tmp106173, tmp106177, tmp106205)

						__e.TailApply(tmp106172, symlet, tmp106206)
						return

					}, 1)

					tmp106207 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106208 := Call(__e, tmp106207, symfail, Nil)

					__e.TailApply(tmp106171, tmp106208)
					return

				}, 1)

				tmp106209 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				tmp106210 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106211 := Call(__e, tmp106210, V4813)

				tmp106212 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				tmp106213 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp106214 := Call(__e, tmp106213, V4813)

				tmp106215 := Call(__e, tmp106212, symParse__, tmp106214)

				tmp106216 := Call(__e, tmp106209, tmp106211, tmp106215, V4815)

				__e.TailApply(tmp106170, tmp106216)
				return

			}, 1)

			tmp106217 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106218 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106219 := Call(__e, tmp106218, V4813)

			tmp106220 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106221 := Call(__e, tmp106220, V4814, Nil)

			tmp106222 := Call(__e, tmp106217, tmp106219, tmp106221)

			__e.TailApply(tmp106169, tmp106222)
			return

		} else {
			tmp106168 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp106168, symshen_4recursive__descent)
			return

		}

	}, 3)

	tmp106225 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__descent, tmp106166)

	_ = tmp106225

	tmp106226 := MakeNative(func(__e *ControlFlow) {
		V4819 := __e.Get(1)
		_ = V4819
		V4820 := __e.Get(2)
		_ = V4820
		V4821 := __e.Get(3)
		_ = V4821
		tmp106283 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106284 := Call(__e, tmp106283, V4819)

		if True == tmp106284 {
			tmp106229 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp106230 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp106231 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp106232 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106233 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106234 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106235 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106236 := Call(__e, tmp106235, Else, Nil)

						tmp106237 := Call(__e, tmp106234, Action, tmp106236)

						tmp106238 := Call(__e, tmp106233, Test, tmp106237)

						__e.TailApply(tmp106232, symif, tmp106238)
						return

					}, 1)

					tmp106239 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106240 := Call(__e, tmp106239, symfail, Nil)

					__e.TailApply(tmp106231, tmp106240)
					return

				}, 1)

				tmp106241 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106242 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106243 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				tmp106244 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp106245 := Call(__e, tmp106244, V4819)

				tmp106246 := Call(__e, tmp106243, symParse__, tmp106245)

				tmp106247 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106248 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106249 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106250 := Call(__e, tmp106249, V4820, Nil)

				tmp106251 := Call(__e, tmp106248, symshen_4hdhd, tmp106250)

				tmp106252 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				tmp106254 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106255 := Call(__e, tmp106254, V4819)

				tmp106256 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106257 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106258 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106260 := Call(__e, tmp106259, V4820, Nil)

				tmp106261 := Call(__e, tmp106258, symshen_4tlhd, tmp106260)

				tmp106262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106265 := Call(__e, tmp106264, V4820, Nil)

				tmp106266 := Call(__e, tmp106263, symshen_4hdtl, tmp106265)

				tmp106267 := Call(__e, tmp106262, tmp106266, Nil)

				tmp106268 := Call(__e, tmp106257, tmp106261, tmp106267)

				tmp106269 := Call(__e, tmp106256, symshen_4pair, tmp106268)

				tmp106270 := Call(__e, tmp106253, tmp106255, tmp106269, V4821)

				tmp106271 := Call(__e, tmp106252, tmp106270, Nil)

				tmp106272 := Call(__e, tmp106247, tmp106251, tmp106271)

				tmp106273 := Call(__e, tmp106242, tmp106246, tmp106272)

				tmp106274 := Call(__e, tmp106241, symlet, tmp106273)

				__e.TailApply(tmp106230, tmp106274)
				return

			}, 1)

			tmp106275 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106276 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106279 := Call(__e, tmp106278, V4820, Nil)

			tmp106280 := Call(__e, tmp106277, symhd, tmp106279)

			tmp106281 := Call(__e, tmp106276, tmp106280, Nil)

			tmp106282 := Call(__e, tmp106275, symcons_2, tmp106281)

			__e.TailApply(tmp106229, tmp106282)
			return

		} else {
			tmp106228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp106228, symshen_4variable_1match)
			return

		}

	}, 3)

	tmp106285 := Call(__e, PrimNS1Value(symns2_1set), symshen_4variable_1match, tmp106226)

	_ = tmp106285

	tmp106286 := MakeNative(func(__e *ControlFlow) {
		V4831 := __e.Get(1)
		_ = V4831
		tmp106291 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106292 := Call(__e, tmp106291, V4831)

		if True == tmp106292 {
			__e.Return(False)
			return
		} else {
			tmp106289 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

			tmp106290 := Call(__e, tmp106289, V4831)

			if True == tmp106290 {
				__e.Return(False)
				return
			} else {
				__e.Return(True)
				return
			}

		}

	}, 1)

	tmp106293 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terminal_2, tmp106286)

	_ = tmp106293

	tmp106294 := MakeNative(func(__e *ControlFlow) {
		V4837 := __e.Get(1)
		_ = V4837
		tmp106296 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp106297 := Call(__e, tmp106296, V4837, sym__)

		if True == tmp106297 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp106298 := Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream_2, tmp106294)

	_ = tmp106298

	tmp106299 := MakeNative(func(__e *ControlFlow) {
		V4841 := __e.Get(1)
		_ = V4841
		V4842 := __e.Get(2)
		_ = V4842
		V4843 := __e.Get(3)
		_ = V4843
		tmp106369 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106370 := Call(__e, tmp106369, V4841)

		if True == tmp106370 {
			tmp106302 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp106303 := MakeNative(func(__e *ControlFlow) {
					NewStr := __e.Get(1)
					_ = NewStr
					tmp106304 := MakeNative(func(__e *ControlFlow) {
						Action := __e.Get(1)
						_ = Action
						tmp106305 := MakeNative(func(__e *ControlFlow) {
							Else := __e.Get(1)
							_ = Else
							tmp106306 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp106307 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp106308 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp106309 := Call(__e, PrimNS1Value(symns2_1value), symcons)

							tmp106310 := Call(__e, tmp106309, Else, Nil)

							tmp106311 := Call(__e, tmp106308, Action, tmp106310)

							tmp106312 := Call(__e, tmp106307, Test, tmp106311)

							__e.TailApply(tmp106306, symif, tmp106312)
							return

						}, 1)

						tmp106313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106314 := Call(__e, tmp106313, symfail, Nil)

						__e.TailApply(tmp106305, tmp106314)
						return

					}, 1)

					tmp106315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106322 := Call(__e, tmp106321, V4842, Nil)

					tmp106323 := Call(__e, tmp106320, symshen_4tlhd, tmp106322)

					tmp106324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106327 := Call(__e, tmp106326, V4842, Nil)

					tmp106328 := Call(__e, tmp106325, symshen_4hdtl, tmp106327)

					tmp106329 := Call(__e, tmp106324, tmp106328, Nil)

					tmp106330 := Call(__e, tmp106319, tmp106323, tmp106329)

					tmp106331 := Call(__e, tmp106318, symshen_4pair, tmp106330)

					tmp106332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106333 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

					tmp106334 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp106335 := Call(__e, tmp106334, V4841)

					tmp106336 := Call(__e, tmp106333, tmp106335, NewStr, V4843)

					tmp106337 := Call(__e, tmp106332, tmp106336, Nil)

					tmp106338 := Call(__e, tmp106317, tmp106331, tmp106337)

					tmp106339 := Call(__e, tmp106316, NewStr, tmp106338)

					tmp106340 := Call(__e, tmp106315, symlet, tmp106339)

					__e.TailApply(tmp106304, tmp106340)
					return

				}, 1)

				tmp106341 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				tmp106342 := Call(__e, tmp106341, symNewStream)

				__e.TailApply(tmp106303, tmp106342)
				return

			}, 1)

			tmp106343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106344 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106345 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106347 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106349 := Call(__e, tmp106348, V4842, Nil)

			tmp106350 := Call(__e, tmp106347, symhd, tmp106349)

			tmp106351 := Call(__e, tmp106346, tmp106350, Nil)

			tmp106352 := Call(__e, tmp106345, symcons_2, tmp106351)

			tmp106353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106356 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106357 := Call(__e, tmp106356, V4841)

			tmp106358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106359 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106360 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106361 := Call(__e, tmp106360, V4842, Nil)

			tmp106362 := Call(__e, tmp106359, symshen_4hdhd, tmp106361)

			tmp106363 := Call(__e, tmp106358, tmp106362, Nil)

			tmp106364 := Call(__e, tmp106355, tmp106357, tmp106363)

			tmp106365 := Call(__e, tmp106354, sym_a, tmp106364)

			tmp106366 := Call(__e, tmp106353, tmp106365, Nil)

			tmp106367 := Call(__e, tmp106344, tmp106352, tmp106366)

			tmp106368 := Call(__e, tmp106343, symand, tmp106367)

			__e.TailApply(tmp106302, tmp106368)
			return

		} else {
			tmp106301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp106301, symshen_4check__stream)
			return

		}

	}, 3)

	tmp106371 := Call(__e, PrimNS1Value(symns2_1set), symshen_4check__stream, tmp106299)

	_ = tmp106371

	tmp106372 := MakeNative(func(__e *ControlFlow) {
		V4847 := __e.Get(1)
		_ = V4847
		V4848 := __e.Get(2)
		_ = V4848
		V4849 := __e.Get(3)
		_ = V4849
		tmp106413 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106414 := Call(__e, tmp106413, V4847)

		if True == tmp106414 {
			tmp106375 := MakeNative(func(__e *ControlFlow) {
				Test := __e.Get(1)
				_ = Test
				tmp106376 := MakeNative(func(__e *ControlFlow) {
					Action := __e.Get(1)
					_ = Action
					tmp106377 := MakeNative(func(__e *ControlFlow) {
						Else := __e.Get(1)
						_ = Else
						tmp106378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106379 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106380 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106381 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp106382 := Call(__e, tmp106381, Else, Nil)

						tmp106383 := Call(__e, tmp106380, Action, tmp106382)

						tmp106384 := Call(__e, tmp106379, Test, tmp106383)

						__e.TailApply(tmp106378, symif, tmp106384)
						return

					}, 1)

					tmp106385 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp106386 := Call(__e, tmp106385, symfail, Nil)

					__e.TailApply(tmp106377, tmp106386)
					return

				}, 1)

				tmp106387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4syntax)

				tmp106388 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106389 := Call(__e, tmp106388, V4847)

				tmp106390 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106391 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106392 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106393 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106394 := Call(__e, tmp106393, V4848, Nil)

				tmp106395 := Call(__e, tmp106392, symshen_4tlhd, tmp106394)

				tmp106396 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106397 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106398 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106399 := Call(__e, tmp106398, V4848, Nil)

				tmp106400 := Call(__e, tmp106397, symshen_4hdtl, tmp106399)

				tmp106401 := Call(__e, tmp106396, tmp106400, Nil)

				tmp106402 := Call(__e, tmp106391, tmp106395, tmp106401)

				tmp106403 := Call(__e, tmp106390, symshen_4pair, tmp106402)

				tmp106404 := Call(__e, tmp106387, tmp106389, tmp106403, V4849)

				__e.TailApply(tmp106376, tmp106404)
				return

			}, 1)

			tmp106405 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106406 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106407 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106408 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106409 := Call(__e, tmp106408, V4848, Nil)

			tmp106410 := Call(__e, tmp106407, symhd, tmp106409)

			tmp106411 := Call(__e, tmp106406, tmp106410, Nil)

			tmp106412 := Call(__e, tmp106405, symcons_2, tmp106411)

			__e.TailApply(tmp106375, tmp106412)
			return

		} else {
			tmp106374 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp106374, symshen_4jump__stream)
			return

		}

	}, 3)

	tmp106415 := Call(__e, PrimNS1Value(symns2_1set), symshen_4jump__stream, tmp106372)

	_ = tmp106415

	tmp106416 := MakeNative(func(__e *ControlFlow) {
		V4851 := __e.Get(1)
		_ = V4851
		tmp106436 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp106437 := Call(__e, tmp106436, Nil, V4851)

		if True == tmp106437 {
			__e.Return(Nil)
			return
		} else {
			tmp106434 := Call(__e, PrimNS1Value(symns2_1value), symshen_4grammar__symbol_2)

			tmp106435 := Call(__e, tmp106434, V4851)

			if True == tmp106435 {
				tmp106429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106430 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp106431 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

				tmp106432 := Call(__e, tmp106431, symParse__, V4851)

				tmp106433 := Call(__e, tmp106430, tmp106432, Nil)

				__e.TailApply(tmp106429, symshen_4hdtl, tmp106433)
				return

			} else {
				tmp106427 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp106428 := Call(__e, tmp106427, V4851)

				if True == tmp106428 {
					tmp106426 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

					__e.TailApply(tmp106426, symParse__, V4851)
					return

				} else {
					tmp106424 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp106425 := Call(__e, tmp106424, V4851)

					if True == tmp106425 {
						tmp106421 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						tmp106422 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							tmp106423 := Call(__e, PrimNS1Value(symns2_1value), symshen_4semantics)

							__e.TailApply(tmp106423, Z)
							return

						}, 1)

						__e.TailApply(tmp106421, tmp106422, V4851)
						return

					} else {
						__e.Return(V4851)
						return
					}

				}

			}

		}

	}, 1)

	tmp106438 := Call(__e, PrimNS1Value(symns2_1set), symshen_4semantics, tmp106416)

	_ = tmp106438

	tmp106439 := MakeNative(func(__e *ControlFlow) {
		V4854 := __e.Get(1)
		_ = V4854
		V4855 := __e.Get(2)
		_ = V4855
		tmp106440 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp106441 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp106442 := Call(__e, tmp106441, V4855, Nil)

		__e.TailApply(tmp106440, V4854, tmp106442)
		return

	}, 2)

	tmp106443 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pair, tmp106439)

	_ = tmp106443

	tmp106444 := MakeNative(func(__e *ControlFlow) {
		V4857 := __e.Get(1)
		_ = V4857
		tmp106445 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp106446 := Call(__e, PrimNS1Value(symns2_1value), symtl)

		tmp106447 := Call(__e, tmp106446, V4857)

		__e.TailApply(tmp106445, tmp106447)
		return

	}, 1)

	tmp106448 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hdtl, tmp106444)

	_ = tmp106448

	tmp106449 := MakeNative(func(__e *ControlFlow) {
		V4859 := __e.Get(1)
		_ = V4859
		tmp106450 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp106451 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp106452 := Call(__e, tmp106451, V4859)

		__e.TailApply(tmp106450, tmp106452)
		return

	}, 1)

	tmp106453 := Call(__e, PrimNS1Value(symns2_1set), symshen_4hdhd, tmp106449)

	_ = tmp106453

	tmp106454 := MakeNative(func(__e *ControlFlow) {
		V4861 := __e.Get(1)
		_ = V4861
		tmp106455 := Call(__e, PrimNS1Value(symns2_1value), symtl)

		tmp106456 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp106457 := Call(__e, tmp106456, V4861)

		__e.TailApply(tmp106455, tmp106457)
		return

	}, 1)

	tmp106458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tlhd, tmp106454)

	_ = tmp106458

	tmp106459 := MakeNative(func(__e *ControlFlow) {
		V4869 := __e.Get(1)
		_ = V4869
		tmp106479 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106480 := Call(__e, tmp106479, V4869)

		var ifres106465 Obj

		if True == tmp106480 {
			tmp106475 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106476 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp106477 := Call(__e, tmp106476, V4869)

			tmp106478 := Call(__e, tmp106475, tmp106477)

			var ifres106467 Obj

			if True == tmp106478 {
				tmp106469 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp106470 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106471 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106472 := Call(__e, tmp106471, V4869)

				tmp106473 := Call(__e, tmp106470, tmp106472)

				tmp106474 := Call(__e, tmp106469, Nil, tmp106473)

				var ifres106468 Obj

				if True == tmp106474 {
					ifres106468 = True

				} else {
					ifres106468 = False

				}

				ifres106467 = ifres106468

			} else {
				ifres106467 = False

			}

			var ifres106466 Obj

			if True == ifres106467 {
				ifres106466 = True

			} else {
				ifres106466 = False

			}

			ifres106465 = ifres106466

		} else {
			ifres106465 = False

		}

		if True == ifres106465 {
			tmp106462 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106463 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp106464 := Call(__e, tmp106463, V4869)

			__e.TailApply(tmp106462, tmp106464)
			return

		} else {
			tmp106461 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp106461)
			return

		}

	}, 1)

	tmp106481 := Call(__e, PrimNS1Value(symns2_1set), symshen_4snd_1or_1fail, tmp106459)

	_ = tmp106481

	tmp106482 := MakeNative(func(__e *ControlFlow) {
		__e.Return(symshen_4fail_b)
		return
	}, 0)

	tmp106483 := Call(__e, PrimNS1Value(symns2_1set), symfail, tmp106482)

	_ = tmp106483

	tmp106484 := MakeNative(func(__e *ControlFlow) {
		V4877 := __e.Get(1)
		_ = V4877
		tmp106506 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106507 := Call(__e, tmp106506, V4877)

		var ifres106492 Obj

		if True == tmp106507 {
			tmp106502 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106503 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp106504 := Call(__e, tmp106503, V4877)

			tmp106505 := Call(__e, tmp106502, tmp106504)

			var ifres106494 Obj

			if True == tmp106505 {
				tmp106496 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp106497 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106498 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106499 := Call(__e, tmp106498, V4877)

				tmp106500 := Call(__e, tmp106497, tmp106499)

				tmp106501 := Call(__e, tmp106496, Nil, tmp106500)

				var ifres106495 Obj

				if True == tmp106501 {
					ifres106495 = True

				} else {
					ifres106495 = False

				}

				ifres106494 = ifres106495

			} else {
				ifres106494 = False

			}

			var ifres106493 Obj

			if True == ifres106494 {
				ifres106493 = True

			} else {
				ifres106493 = False

			}

			ifres106492 = ifres106493

		} else {
			ifres106492 = False

		}

		if True == ifres106492 {
			tmp106487 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106488 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106489 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106490 := Call(__e, tmp106489, V4877)

			tmp106491 := Call(__e, tmp106488, tmp106490, Nil)

			__e.TailApply(tmp106487, Nil, tmp106491)
			return

		} else {
			tmp106486 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp106486)
			return

		}

	}, 1)

	tmp106508 := Call(__e, PrimNS1Value(symns2_1set), sym_5_b_6, tmp106484)

	_ = tmp106508

	tmp106509 := MakeNative(func(__e *ControlFlow) {
		V4883 := __e.Get(1)
		_ = V4883
		tmp106531 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp106532 := Call(__e, tmp106531, V4883)

		var ifres106517 Obj

		if True == tmp106532 {
			tmp106527 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp106528 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp106529 := Call(__e, tmp106528, V4883)

			tmp106530 := Call(__e, tmp106527, tmp106529)

			var ifres106519 Obj

			if True == tmp106530 {
				tmp106521 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp106522 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106523 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp106524 := Call(__e, tmp106523, V4883)

				tmp106525 := Call(__e, tmp106522, tmp106524)

				tmp106526 := Call(__e, tmp106521, Nil, tmp106525)

				var ifres106520 Obj

				if True == tmp106526 {
					ifres106520 = True

				} else {
					ifres106520 = False

				}

				ifres106519 = ifres106520

			} else {
				ifres106519 = False

			}

			var ifres106518 Obj

			if True == ifres106519 {
				ifres106518 = True

			} else {
				ifres106518 = False

			}

			ifres106517 = ifres106518

		} else {
			ifres106517 = False

		}

		if True == ifres106517 {
			tmp106512 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106513 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp106514 := Call(__e, tmp106513, V4883)

			tmp106515 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp106516 := Call(__e, tmp106515, Nil, Nil)

			__e.TailApply(tmp106512, tmp106514, tmp106516)
			return

		} else {
			tmp106511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp106511, sym_5e_6)
			return

		}

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), sym_5e_6, tmp106509)
	return

}, 0)
