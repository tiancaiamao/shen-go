package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp89401 := MakeNative(func(__e *ControlFlow) {
		V1722 := __e.Get(1)
		_ = V1722
		tmp89427 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89428 := Call(__e, tmp89427, V1722)

		var ifres89413 Obj

		if True == tmp89428 {
			tmp89423 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89424 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89425 := Call(__e, tmp89424, V1722)

			tmp89426 := Call(__e, tmp89423, tmp89425)

			var ifres89415 Obj

			if True == tmp89426 {
				tmp89417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp89418 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89419 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89420 := Call(__e, tmp89419, V1722)

				tmp89421 := Call(__e, tmp89418, tmp89420)

				tmp89422 := Call(__e, tmp89417, Nil, tmp89421)

				var ifres89416 Obj

				if True == tmp89422 {
					ifres89416 = True

				} else {
					ifres89416 = False

				}

				ifres89415 = ifres89416

			} else {
				ifres89415 = False

			}

			var ifres89414 Obj

			if True == ifres89415 {
				ifres89414 = True

			} else {
				ifres89414 = False

			}

			ifres89413 = ifres89414

		} else {
			ifres89413 = False

		}

		if True == ifres89413 {
			tmp89404 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			tmp89405 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp89406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp89407 := Call(__e, PrimNS1Value(symns2_1value), symshen_4next_150)

			tmp89408 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89409 := Call(__e, tmp89408, V1722)

			tmp89410 := Call(__e, tmp89407, MakeNumber(50), tmp89409)

			tmp89411 := Call(__e, tmp89406, tmp89410, MakeString("\n"), symshen_4a)

			tmp89412 := Call(__e, tmp89405, MakeString("datatype syntax error here:\n\n "), tmp89411)

			__e.TailApply(tmp89404, tmp89412)
			return

		} else {
			tmp89403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp89403, symshen_4datatype_1error)
			return

		}

	}, 1)

	tmp89429 := Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1error, tmp89401)

	_ = tmp89429

	tmp89430 := MakeNative(func(__e *ControlFlow) {
		V1724 := __e.Get(1)
		_ = V1724
		tmp89431 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89447 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89448 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89449 := Call(__e, tmp89448)

			tmp89450 := Call(__e, tmp89447, YaccParse, tmp89449)

			if True == tmp89450 {
				tmp89433 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp89439 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89440 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89441 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89442 := Call(__e, tmp89441)

					tmp89443 := Call(__e, tmp89440, tmp89442, Parse___5e_6)

					tmp89444 := Call(__e, tmp89439, tmp89443)

					if True == tmp89444 {
						tmp89436 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89437 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89438 := Call(__e, tmp89437, Parse___5e_6)

						__e.TailApply(tmp89436, tmp89438, Nil)
						return

					} else {
						tmp89435 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89435)
						return

					}

				}, 1)

				tmp89445 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp89446 := Call(__e, tmp89445, V1724)

				__e.TailApply(tmp89433, tmp89446)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89451 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5datatype_1rule_6 := __e.Get(1)
			_ = Parse__shen_4_5datatype_1rule_6
			tmp89474 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp89475 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89476 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89477 := Call(__e, tmp89476)

			tmp89478 := Call(__e, tmp89475, tmp89477, Parse__shen_4_5datatype_1rule_6)

			tmp89479 := Call(__e, tmp89474, tmp89478)

			if True == tmp89479 {
				tmp89454 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5datatype_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5datatype_1rules_6
					tmp89466 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89468 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89469 := Call(__e, tmp89468)

					tmp89470 := Call(__e, tmp89467, tmp89469, Parse__shen_4_5datatype_1rules_6)

					tmp89471 := Call(__e, tmp89466, tmp89470)

					if True == tmp89471 {
						tmp89457 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89458 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89459 := Call(__e, tmp89458, Parse__shen_4_5datatype_1rules_6)

						tmp89460 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp89462 := Call(__e, tmp89461, Parse__shen_4_5datatype_1rule_6)

						tmp89463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp89464 := Call(__e, tmp89463, Parse__shen_4_5datatype_1rules_6)

						tmp89465 := Call(__e, tmp89460, tmp89462, tmp89464)

						__e.TailApply(tmp89457, tmp89459, tmp89465)
						return

					} else {
						tmp89456 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89456)
						return

					}

				}, 1)

				tmp89472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5datatype_1rules_6)

				tmp89473 := Call(__e, tmp89472, Parse__shen_4_5datatype_1rule_6)

				__e.TailApply(tmp89454, tmp89473)
				return

			} else {
				tmp89453 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp89453)
				return

			}

		}, 1)

		tmp89480 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5datatype_1rule_6)

		tmp89481 := Call(__e, tmp89480, V1724)

		tmp89482 := Call(__e, tmp89451, tmp89481)

		__e.TailApply(tmp89431, tmp89482)
		return

	}, 1)

	tmp89483 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rules_6, tmp89430)

	_ = tmp89483

	tmp89484 := MakeNative(func(__e *ControlFlow) {
		V1726 := __e.Get(1)
		_ = V1726
		tmp89485 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89548 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89549 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89550 := Call(__e, tmp89549)

			tmp89551 := Call(__e, tmp89548, YaccParse, tmp89550)

			if True == tmp89551 {
				tmp89487 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					tmp89540 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89541 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89542 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89543 := Call(__e, tmp89542)

					tmp89544 := Call(__e, tmp89541, tmp89543, Parse__shen_4_5side_1conditions_6)

					tmp89545 := Call(__e, tmp89540, tmp89544)

					if True == tmp89545 {
						tmp89490 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							tmp89532 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89533 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89534 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp89535 := Call(__e, tmp89534)

							tmp89536 := Call(__e, tmp89533, tmp89535, Parse__shen_4_5premises_6)

							tmp89537 := Call(__e, tmp89532, tmp89536)

							if True == tmp89537 {
								tmp89493 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5doubleunderline_6 := __e.Get(1)
									_ = Parse__shen_4_5doubleunderline_6
									tmp89524 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp89525 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp89526 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp89527 := Call(__e, tmp89526)

									tmp89528 := Call(__e, tmp89525, tmp89527, Parse__shen_4_5doubleunderline_6)

									tmp89529 := Call(__e, tmp89524, tmp89528)

									if True == tmp89529 {
										tmp89496 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5conclusion_6 := __e.Get(1)
											_ = Parse__shen_4_5conclusion_6
											tmp89516 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											tmp89517 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											tmp89518 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											tmp89519 := Call(__e, tmp89518)

											tmp89520 := Call(__e, tmp89517, tmp89519, Parse__shen_4_5conclusion_6)

											tmp89521 := Call(__e, tmp89516, tmp89520)

											if True == tmp89521 {
												tmp89499 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												tmp89500 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												tmp89501 := Call(__e, tmp89500, Parse__shen_4_5conclusion_6)

												tmp89502 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

												tmp89503 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp89504 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp89505 := Call(__e, tmp89504, Parse__shen_4_5side_1conditions_6)

												tmp89506 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp89507 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp89508 := Call(__e, tmp89507, Parse__shen_4_5premises_6)

												tmp89509 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												tmp89510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												tmp89511 := Call(__e, tmp89510, Parse__shen_4_5conclusion_6)

												tmp89512 := Call(__e, tmp89509, tmp89511, Nil)

												tmp89513 := Call(__e, tmp89506, tmp89508, tmp89512)

												tmp89514 := Call(__e, tmp89503, tmp89505, tmp89513)

												tmp89515 := Call(__e, tmp89502, symshen_4double, tmp89514)

												__e.TailApply(tmp89499, tmp89501, tmp89515)
												return

											} else {
												tmp89498 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(tmp89498)
												return

											}

										}, 1)

										tmp89522 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5conclusion_6)

										tmp89523 := Call(__e, tmp89522, Parse__shen_4_5doubleunderline_6)

										__e.TailApply(tmp89496, tmp89523)
										return

									} else {
										tmp89495 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp89495)
										return

									}

								}, 1)

								tmp89530 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5doubleunderline_6)

								tmp89531 := Call(__e, tmp89530, Parse__shen_4_5premises_6)

								__e.TailApply(tmp89493, tmp89531)
								return

							} else {
								tmp89492 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp89492)
								return

							}

						}, 1)

						tmp89538 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

						tmp89539 := Call(__e, tmp89538, Parse__shen_4_5side_1conditions_6)

						__e.TailApply(tmp89490, tmp89539)
						return

					} else {
						tmp89489 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89489)
						return

					}

				}, 1)

				tmp89546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

				tmp89547 := Call(__e, tmp89546, V1726)

				__e.TailApply(tmp89487, tmp89547)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89552 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1conditions_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1conditions_6
			tmp89605 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp89606 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89607 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89608 := Call(__e, tmp89607)

			tmp89609 := Call(__e, tmp89606, tmp89608, Parse__shen_4_5side_1conditions_6)

			tmp89610 := Call(__e, tmp89605, tmp89609)

			if True == tmp89610 {
				tmp89555 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5premises_6 := __e.Get(1)
					_ = Parse__shen_4_5premises_6
					tmp89597 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89598 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89599 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89600 := Call(__e, tmp89599)

					tmp89601 := Call(__e, tmp89598, tmp89600, Parse__shen_4_5premises_6)

					tmp89602 := Call(__e, tmp89597, tmp89601)

					if True == tmp89602 {
						tmp89558 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5singleunderline_6 := __e.Get(1)
							_ = Parse__shen_4_5singleunderline_6
							tmp89589 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89590 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89591 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp89592 := Call(__e, tmp89591)

							tmp89593 := Call(__e, tmp89590, tmp89592, Parse__shen_4_5singleunderline_6)

							tmp89594 := Call(__e, tmp89589, tmp89593)

							if True == tmp89594 {
								tmp89561 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5conclusion_6 := __e.Get(1)
									_ = Parse__shen_4_5conclusion_6
									tmp89581 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp89582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp89583 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp89584 := Call(__e, tmp89583)

									tmp89585 := Call(__e, tmp89582, tmp89584, Parse__shen_4_5conclusion_6)

									tmp89586 := Call(__e, tmp89581, tmp89585)

									if True == tmp89586 {
										tmp89564 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp89565 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp89566 := Call(__e, tmp89565, Parse__shen_4_5conclusion_6)

										tmp89567 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										tmp89568 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89569 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp89570 := Call(__e, tmp89569, Parse__shen_4_5side_1conditions_6)

										tmp89571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp89573 := Call(__e, tmp89572, Parse__shen_4_5premises_6)

										tmp89574 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89575 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp89576 := Call(__e, tmp89575, Parse__shen_4_5conclusion_6)

										tmp89577 := Call(__e, tmp89574, tmp89576, Nil)

										tmp89578 := Call(__e, tmp89571, tmp89573, tmp89577)

										tmp89579 := Call(__e, tmp89568, tmp89570, tmp89578)

										tmp89580 := Call(__e, tmp89567, symshen_4single, tmp89579)

										__e.TailApply(tmp89564, tmp89566, tmp89580)
										return

									} else {
										tmp89563 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp89563)
										return

									}

								}, 1)

								tmp89587 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5conclusion_6)

								tmp89588 := Call(__e, tmp89587, Parse__shen_4_5singleunderline_6)

								__e.TailApply(tmp89561, tmp89588)
								return

							} else {
								tmp89560 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp89560)
								return

							}

						}, 1)

						tmp89595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5singleunderline_6)

						tmp89596 := Call(__e, tmp89595, Parse__shen_4_5premises_6)

						__e.TailApply(tmp89558, tmp89596)
						return

					} else {
						tmp89557 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89557)
						return

					}

				}, 1)

				tmp89603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

				tmp89604 := Call(__e, tmp89603, Parse__shen_4_5side_1conditions_6)

				__e.TailApply(tmp89555, tmp89604)
				return

			} else {
				tmp89554 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp89554)
				return

			}

		}, 1)

		tmp89611 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

		tmp89612 := Call(__e, tmp89611, V1726)

		tmp89613 := Call(__e, tmp89552, tmp89612)

		__e.TailApply(tmp89485, tmp89613)
		return

	}, 1)

	tmp89614 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rule_6, tmp89484)

	_ = tmp89614

	tmp89615 := MakeNative(func(__e *ControlFlow) {
		V1728 := __e.Get(1)
		_ = V1728
		tmp89616 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89632 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89633 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89634 := Call(__e, tmp89633)

			tmp89635 := Call(__e, tmp89632, YaccParse, tmp89634)

			if True == tmp89635 {
				tmp89618 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp89624 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89625 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89626 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89627 := Call(__e, tmp89626)

					tmp89628 := Call(__e, tmp89625, tmp89627, Parse___5e_6)

					tmp89629 := Call(__e, tmp89624, tmp89628)

					if True == tmp89629 {
						tmp89621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89622 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89623 := Call(__e, tmp89622, Parse___5e_6)

						__e.TailApply(tmp89621, tmp89623, Nil)
						return

					} else {
						tmp89620 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89620)
						return

					}

				}, 1)

				tmp89630 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp89631 := Call(__e, tmp89630, V1728)

				__e.TailApply(tmp89618, tmp89631)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89636 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1condition_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1condition_6
			tmp89659 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp89660 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89661 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89662 := Call(__e, tmp89661)

			tmp89663 := Call(__e, tmp89660, tmp89662, Parse__shen_4_5side_1condition_6)

			tmp89664 := Call(__e, tmp89659, tmp89663)

			if True == tmp89664 {
				tmp89639 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					tmp89651 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89652 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89653 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89654 := Call(__e, tmp89653)

					tmp89655 := Call(__e, tmp89652, tmp89654, Parse__shen_4_5side_1conditions_6)

					tmp89656 := Call(__e, tmp89651, tmp89655)

					if True == tmp89656 {
						tmp89642 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89643 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89644 := Call(__e, tmp89643, Parse__shen_4_5side_1conditions_6)

						tmp89645 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89646 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp89647 := Call(__e, tmp89646, Parse__shen_4_5side_1condition_6)

						tmp89648 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp89649 := Call(__e, tmp89648, Parse__shen_4_5side_1conditions_6)

						tmp89650 := Call(__e, tmp89645, tmp89647, tmp89649)

						__e.TailApply(tmp89642, tmp89644, tmp89650)
						return

					} else {
						tmp89641 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89641)
						return

					}

				}, 1)

				tmp89657 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

				tmp89658 := Call(__e, tmp89657, Parse__shen_4_5side_1condition_6)

				__e.TailApply(tmp89639, tmp89658)
				return

			} else {
				tmp89638 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp89638)
				return

			}

		}, 1)

		tmp89665 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1condition_6)

		tmp89666 := Call(__e, tmp89665, V1728)

		tmp89667 := Call(__e, tmp89636, tmp89666)

		__e.TailApply(tmp89616, tmp89667)
		return

	}, 1)

	tmp89668 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1conditions_6, tmp89615)

	_ = tmp89668

	tmp89669 := MakeNative(func(__e *ControlFlow) {
		V1732 := __e.Get(1)
		_ = V1732
		tmp89670 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89726 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89727 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89728 := Call(__e, tmp89727)

			tmp89729 := Call(__e, tmp89726, YaccParse, tmp89728)

			if True == tmp89729 {
				tmp89722 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89723 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp89724 := Call(__e, tmp89723, V1732)

				tmp89725 := Call(__e, tmp89722, tmp89724)

				var ifres89716 Obj

				if True == tmp89725 {
					tmp89718 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89719 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp89720 := Call(__e, tmp89719, V1732)

					tmp89721 := Call(__e, tmp89718, symlet, tmp89720)

					var ifres89717 Obj

					if True == tmp89721 {
						ifres89717 = True

					} else {
						ifres89717 = False

					}

					ifres89716 = ifres89717

				} else {
					ifres89716 = False

				}

				if True == ifres89716 {
					tmp89674 := MakeNative(func(__e *ControlFlow) {
						NewStream1730 := __e.Get(1)
						_ = NewStream1730
						tmp89675 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5variable_2_6 := __e.Get(1)
							_ = Parse__shen_4_5variable_2_6
							tmp89702 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89703 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89704 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp89705 := Call(__e, tmp89704)

							tmp89706 := Call(__e, tmp89703, tmp89705, Parse__shen_4_5variable_2_6)

							tmp89707 := Call(__e, tmp89702, tmp89706)

							if True == tmp89707 {
								tmp89678 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5expr_6 := __e.Get(1)
									_ = Parse__shen_4_5expr_6
									tmp89694 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp89695 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp89696 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp89697 := Call(__e, tmp89696)

									tmp89698 := Call(__e, tmp89695, tmp89697, Parse__shen_4_5expr_6)

									tmp89699 := Call(__e, tmp89694, tmp89698)

									if True == tmp89699 {
										tmp89681 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp89682 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp89683 := Call(__e, tmp89682, Parse__shen_4_5expr_6)

										tmp89684 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89685 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89686 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp89687 := Call(__e, tmp89686, Parse__shen_4_5variable_2_6)

										tmp89688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										tmp89689 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp89690 := Call(__e, tmp89689, Parse__shen_4_5expr_6)

										tmp89691 := Call(__e, tmp89688, tmp89690, Nil)

										tmp89692 := Call(__e, tmp89685, tmp89687, tmp89691)

										tmp89693 := Call(__e, tmp89684, symlet, tmp89692)

										__e.TailApply(tmp89681, tmp89683, tmp89693)
										return

									} else {
										tmp89680 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp89680)
										return

									}

								}, 1)

								tmp89700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

								tmp89701 := Call(__e, tmp89700, Parse__shen_4_5variable_2_6)

								__e.TailApply(tmp89678, tmp89701)
								return

							} else {
								tmp89677 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp89677)
								return

							}

						}, 1)

						tmp89708 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5variable_2_6)

						tmp89709 := Call(__e, tmp89708, NewStream1730)

						__e.TailApply(tmp89675, tmp89709)
						return

					}, 1)

					tmp89710 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89711 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp89712 := Call(__e, tmp89711, V1732)

					tmp89713 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp89714 := Call(__e, tmp89713, V1732)

					tmp89715 := Call(__e, tmp89710, tmp89712, tmp89714)

					__e.TailApply(tmp89674, tmp89715)
					return

				} else {
					tmp89673 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp89673)
					return

				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89767 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89768 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp89769 := Call(__e, tmp89768, V1732)

		tmp89770 := Call(__e, tmp89767, tmp89769)

		var ifres89761 Obj

		if True == tmp89770 {
			tmp89763 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89764 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp89765 := Call(__e, tmp89764, V1732)

			tmp89766 := Call(__e, tmp89763, symif, tmp89765)

			var ifres89762 Obj

			if True == tmp89766 {
				ifres89762 = True

			} else {
				ifres89762 = False

			}

			ifres89761 = ifres89762

		} else {
			ifres89761 = False

		}

		var ifres89730 Obj

		if True == ifres89761 {
			tmp89733 := MakeNative(func(__e *ControlFlow) {
				NewStream1729 := __e.Get(1)
				_ = NewStream1729
				tmp89734 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					tmp89746 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89747 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89748 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89749 := Call(__e, tmp89748)

					tmp89750 := Call(__e, tmp89747, tmp89749, Parse__shen_4_5expr_6)

					tmp89751 := Call(__e, tmp89746, tmp89750)

					if True == tmp89751 {
						tmp89737 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89738 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89739 := Call(__e, tmp89738, Parse__shen_4_5expr_6)

						tmp89740 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89741 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp89742 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp89743 := Call(__e, tmp89742, Parse__shen_4_5expr_6)

						tmp89744 := Call(__e, tmp89741, tmp89743, Nil)

						tmp89745 := Call(__e, tmp89740, symif, tmp89744)

						__e.TailApply(tmp89737, tmp89739, tmp89745)
						return

					} else {
						tmp89736 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89736)
						return

					}

				}, 1)

				tmp89752 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

				tmp89753 := Call(__e, tmp89752, NewStream1729)

				__e.TailApply(tmp89734, tmp89753)
				return

			}, 1)

			tmp89754 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp89755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp89756 := Call(__e, tmp89755, V1732)

			tmp89757 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp89758 := Call(__e, tmp89757, V1732)

			tmp89759 := Call(__e, tmp89754, tmp89756, tmp89758)

			tmp89760 := Call(__e, tmp89733, tmp89759)

			ifres89730 = tmp89760

		} else {
			tmp89731 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89732 := Call(__e, tmp89731)

			ifres89730 = tmp89732

		}

		__e.TailApply(tmp89670, ifres89730)
		return

	}, 1)

	tmp89771 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1condition_6, tmp89669)

	_ = tmp89771

	tmp89772 := MakeNative(func(__e *ControlFlow) {
		V1734 := __e.Get(1)
		_ = V1734
		tmp89791 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89792 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp89793 := Call(__e, tmp89792, V1734)

		tmp89794 := Call(__e, tmp89791, tmp89793)

		if True == tmp89794 {
			tmp89775 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp89787 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				tmp89788 := Call(__e, tmp89787, Parse__X)

				if True == tmp89788 {
					tmp89778 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89779 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89780 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp89782 := Call(__e, tmp89781, V1734)

					tmp89783 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp89784 := Call(__e, tmp89783, V1734)

					tmp89785 := Call(__e, tmp89780, tmp89782, tmp89784)

					tmp89786 := Call(__e, tmp89779, tmp89785)

					__e.TailApply(tmp89778, tmp89786, Parse__X)
					return

				} else {
					tmp89777 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp89777)
					return

				}

			}, 1)

			tmp89789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp89790 := Call(__e, tmp89789, V1734)

			__e.TailApply(tmp89775, tmp89790)
			return

		} else {
			tmp89774 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp89774)
			return

		}

	}, 1)

	tmp89795 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5variable_2_6, tmp89772)

	_ = tmp89795

	tmp89796 := MakeNative(func(__e *ControlFlow) {
		V1736 := __e.Get(1)
		_ = V1736
		tmp89831 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89832 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp89833 := Call(__e, tmp89832, V1736)

		tmp89834 := Call(__e, tmp89831, tmp89833)

		if True == tmp89834 {
			tmp89799 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp89813 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				tmp89822 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				tmp89823 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp89824 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp89825 := Call(__e, tmp89824, sym_k, Nil)

				tmp89826 := Call(__e, tmp89823, sym_6_6, tmp89825)

				tmp89827 := Call(__e, tmp89822, Parse__X, tmp89826)

				var ifres89814 Obj

				if True == tmp89827 {
					ifres89814 = True

				} else {
					tmp89820 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

					tmp89821 := Call(__e, tmp89820, Parse__X)

					var ifres89816 Obj

					if True == tmp89821 {
						ifres89816 = True

					} else {
						tmp89818 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

						tmp89819 := Call(__e, tmp89818, Parse__X)

						var ifres89817 Obj

						if True == tmp89819 {
							ifres89817 = True

						} else {
							ifres89817 = False

						}

						ifres89816 = ifres89817

					}

					var ifres89815 Obj

					if True == ifres89816 {
						ifres89815 = True

					} else {
						ifres89815 = False

					}

					ifres89814 = ifres89815

				}

				tmp89828 := Call(__e, tmp89813, ifres89814)

				if True == tmp89828 {
					tmp89802 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89803 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89805 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp89806 := Call(__e, tmp89805, V1736)

					tmp89807 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp89808 := Call(__e, tmp89807, V1736)

					tmp89809 := Call(__e, tmp89804, tmp89806, tmp89808)

					tmp89810 := Call(__e, tmp89803, tmp89809)

					tmp89811 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

					tmp89812 := Call(__e, tmp89811, Parse__X)

					__e.TailApply(tmp89802, tmp89810, tmp89812)
					return

				} else {
					tmp89801 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp89801)
					return

				}

			}, 1)

			tmp89829 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp89830 := Call(__e, tmp89829, V1736)

			__e.TailApply(tmp89799, tmp89830)
			return

		} else {
			tmp89798 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp89798)
			return

		}

	}, 1)

	tmp89835 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5expr_6, tmp89796)

	_ = tmp89835

	tmp89836 := MakeNative(func(__e *ControlFlow) {
		V1738 := __e.Get(1)
		_ = V1738
		tmp89891 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89892 := Call(__e, tmp89891, V1738)

		var ifres89859 Obj

		if True == tmp89892 {
			tmp89887 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89888 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89889 := Call(__e, tmp89888, V1738)

			tmp89890 := Call(__e, tmp89887, tmp89889)

			var ifres89861 Obj

			if True == tmp89890 {
				tmp89881 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp89882 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89883 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89884 := Call(__e, tmp89883, V1738)

				tmp89885 := Call(__e, tmp89882, tmp89884)

				tmp89886 := Call(__e, tmp89881, tmp89885)

				var ifres89863 Obj

				if True == tmp89886 {
					tmp89873 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89874 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89875 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89876 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp89877 := Call(__e, tmp89876, V1738)

					tmp89878 := Call(__e, tmp89875, tmp89877)

					tmp89879 := Call(__e, tmp89874, tmp89878)

					tmp89880 := Call(__e, tmp89873, Nil, tmp89879)

					var ifres89865 Obj

					if True == tmp89880 {
						tmp89867 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp89868 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89869 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp89870 := Call(__e, tmp89869, V1738)

						tmp89871 := Call(__e, tmp89868, tmp89870)

						tmp89872 := Call(__e, tmp89867, tmp89871, symbar_b)

						var ifres89866 Obj

						if True == tmp89872 {
							ifres89866 = True

						} else {
							ifres89866 = False

						}

						ifres89865 = ifres89866

					} else {
						ifres89865 = False

					}

					var ifres89864 Obj

					if True == ifres89865 {
						ifres89864 = True

					} else {
						ifres89864 = False

					}

					ifres89863 = ifres89864

				} else {
					ifres89863 = False

				}

				var ifres89862 Obj

				if True == ifres89863 {
					ifres89862 = True

				} else {
					ifres89862 = False

				}

				ifres89861 = ifres89862

			} else {
				ifres89861 = False

			}

			var ifres89860 Obj

			if True == ifres89861 {
				ifres89860 = True

			} else {
				ifres89860 = False

			}

			ifres89859 = ifres89860

		} else {
			ifres89859 = False

		}

		if True == ifres89859 {
			tmp89850 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp89851 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89852 := Call(__e, tmp89851, V1738)

			tmp89853 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp89854 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89855 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp89856 := Call(__e, tmp89855, V1738)

			tmp89857 := Call(__e, tmp89854, tmp89856)

			tmp89858 := Call(__e, tmp89853, tmp89857)

			__e.TailApply(tmp89850, tmp89852, tmp89858)
			return

		} else {
			tmp89848 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp89849 := Call(__e, tmp89848, V1738)

			if True == tmp89849 {
				tmp89839 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp89840 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

				tmp89841 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp89842 := Call(__e, tmp89841, V1738)

				tmp89843 := Call(__e, tmp89840, tmp89842)

				tmp89844 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

				tmp89845 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp89846 := Call(__e, tmp89845, V1738)

				tmp89847 := Call(__e, tmp89844, tmp89846)

				__e.TailApply(tmp89839, tmp89843, tmp89847)
				return

			} else {
				__e.Return(V1738)
				return
			}

		}

	}, 1)

	tmp89893 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1bar, tmp89836)

	_ = tmp89893

	tmp89894 := MakeNative(func(__e *ControlFlow) {
		V1740 := __e.Get(1)
		_ = V1740
		tmp89895 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp89911 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89912 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89913 := Call(__e, tmp89912)

			tmp89914 := Call(__e, tmp89911, YaccParse, tmp89913)

			if True == tmp89914 {
				tmp89897 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					tmp89903 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89904 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89905 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89906 := Call(__e, tmp89905)

					tmp89907 := Call(__e, tmp89904, tmp89906, Parse___5e_6)

					tmp89908 := Call(__e, tmp89903, tmp89907)

					if True == tmp89908 {
						tmp89900 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp89901 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp89902 := Call(__e, tmp89901, Parse___5e_6)

						__e.TailApply(tmp89900, tmp89902, Nil)
						return

					} else {
						tmp89899 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89899)
						return

					}

				}, 1)

				tmp89909 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				tmp89910 := Call(__e, tmp89909, V1740)

				__e.TailApply(tmp89897, tmp89910)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp89915 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5premise_6 := __e.Get(1)
			_ = Parse__shen_4_5premise_6
			tmp89949 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp89950 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp89951 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp89952 := Call(__e, tmp89951)

			tmp89953 := Call(__e, tmp89950, tmp89952, Parse__shen_4_5premise_6)

			tmp89954 := Call(__e, tmp89949, tmp89953)

			if True == tmp89954 {
				tmp89918 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5semicolon_1symbol_6
					tmp89941 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp89942 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp89943 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp89944 := Call(__e, tmp89943)

					tmp89945 := Call(__e, tmp89942, tmp89944, Parse__shen_4_5semicolon_1symbol_6)

					tmp89946 := Call(__e, tmp89941, tmp89945)

					if True == tmp89946 {
						tmp89921 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							tmp89933 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89934 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp89935 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp89936 := Call(__e, tmp89935)

							tmp89937 := Call(__e, tmp89934, tmp89936, Parse__shen_4_5premises_6)

							tmp89938 := Call(__e, tmp89933, tmp89937)

							if True == tmp89938 {
								tmp89924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp89925 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp89926 := Call(__e, tmp89925, Parse__shen_4_5premises_6)

								tmp89927 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp89928 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp89929 := Call(__e, tmp89928, Parse__shen_4_5premise_6)

								tmp89930 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp89931 := Call(__e, tmp89930, Parse__shen_4_5premises_6)

								tmp89932 := Call(__e, tmp89927, tmp89929, tmp89931)

								__e.TailApply(tmp89924, tmp89926, tmp89932)
								return

							} else {
								tmp89923 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp89923)
								return

							}

						}, 1)

						tmp89939 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

						tmp89940 := Call(__e, tmp89939, Parse__shen_4_5semicolon_1symbol_6)

						__e.TailApply(tmp89921, tmp89940)
						return

					} else {
						tmp89920 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp89920)
						return

					}

				}, 1)

				tmp89947 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

				tmp89948 := Call(__e, tmp89947, Parse__shen_4_5premise_6)

				__e.TailApply(tmp89918, tmp89948)
				return

			} else {
				tmp89917 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp89917)
				return

			}

		}, 1)

		tmp89955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premise_6)

		tmp89956 := Call(__e, tmp89955, V1740)

		tmp89957 := Call(__e, tmp89915, tmp89956)

		__e.TailApply(tmp89895, tmp89957)
		return

	}, 1)

	tmp89958 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premises_6, tmp89894)

	_ = tmp89958

	tmp89959 := MakeNative(func(__e *ControlFlow) {
		V1742 := __e.Get(1)
		_ = V1742
		tmp89978 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp89979 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp89980 := Call(__e, tmp89979, V1742)

		tmp89981 := Call(__e, tmp89978, tmp89980)

		if True == tmp89981 {
			tmp89962 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp89974 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp89975 := Call(__e, tmp89974, Parse__X, sym_k)

				if True == tmp89975 {
					tmp89965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp89967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp89968 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp89969 := Call(__e, tmp89968, V1742)

					tmp89970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp89971 := Call(__e, tmp89970, V1742)

					tmp89972 := Call(__e, tmp89967, tmp89969, tmp89971)

					tmp89973 := Call(__e, tmp89966, tmp89972)

					__e.TailApply(tmp89965, tmp89973, symshen_4skip)
					return

				} else {
					tmp89964 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp89964)
					return

				}

			}, 1)

			tmp89976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp89977 := Call(__e, tmp89976, V1742)

			__e.TailApply(tmp89962, tmp89977)
			return

		} else {
			tmp89961 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp89961)
			return

		}

	}, 1)

	tmp89982 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_1symbol_6, tmp89959)

	_ = tmp89982

	tmp89983 := MakeNative(func(__e *ControlFlow) {
		V1746 := __e.Get(1)
		_ = V1746
		tmp89984 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp90061 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90062 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90063 := Call(__e, tmp90062)

			tmp90064 := Call(__e, tmp90061, YaccParse, tmp90063)

			if True == tmp90064 {
				tmp89986 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp90006 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90007 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90008 := Call(__e, tmp90007)

					tmp90009 := Call(__e, tmp90006, YaccParse, tmp90008)

					if True == tmp90009 {
						tmp89988 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							tmp89998 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp89999 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90000 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90001 := Call(__e, tmp90000)

							tmp90002 := Call(__e, tmp89999, tmp90001, Parse__shen_4_5formula_6)

							tmp90003 := Call(__e, tmp89998, tmp90002)

							if True == tmp90003 {
								tmp89991 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp89992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp89993 := Call(__e, tmp89992, Parse__shen_4_5formula_6)

								tmp89994 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

								tmp89995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp89996 := Call(__e, tmp89995, Parse__shen_4_5formula_6)

								tmp89997 := Call(__e, tmp89994, Nil, tmp89996)

								__e.TailApply(tmp89991, tmp89993, tmp89997)
								return

							} else {
								tmp89990 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp89990)
								return

							}

						}, 1)

						tmp90004 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

						tmp90005 := Call(__e, tmp90004, V1746)

						__e.TailApply(tmp89988, tmp90005)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp90010 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formulae_6 := __e.Get(1)
					_ = Parse__shen_4_5formulae_6
					tmp90052 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp90053 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90054 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90055 := Call(__e, tmp90054)

					tmp90056 := Call(__e, tmp90053, tmp90055, Parse__shen_4_5formulae_6)

					tmp90057 := Call(__e, tmp90052, tmp90056)

					if True == tmp90057 {
						tmp90048 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp90049 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90050 := Call(__e, tmp90049, Parse__shen_4_5formulae_6)

						tmp90051 := Call(__e, tmp90048, tmp90050)

						var ifres90042 Obj

						if True == tmp90051 {
							tmp90044 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90045 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							tmp90046 := Call(__e, tmp90045, Parse__shen_4_5formulae_6)

							tmp90047 := Call(__e, tmp90044, sym_6_6, tmp90046)

							var ifres90043 Obj

							if True == tmp90047 {
								ifres90043 = True

							} else {
								ifres90043 = False

							}

							ifres90042 = ifres90043

						} else {
							ifres90042 = False

						}

						if True == ifres90042 {
							tmp90015 := MakeNative(func(__e *ControlFlow) {
								NewStream1744 := __e.Get(1)
								_ = NewStream1744
								tmp90016 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5formula_6 := __e.Get(1)
									_ = Parse__shen_4_5formula_6
									tmp90028 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp90029 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp90030 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp90031 := Call(__e, tmp90030)

									tmp90032 := Call(__e, tmp90029, tmp90031, Parse__shen_4_5formula_6)

									tmp90033 := Call(__e, tmp90028, tmp90032)

									if True == tmp90033 {
										tmp90019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp90020 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp90021 := Call(__e, tmp90020, Parse__shen_4_5formula_6)

										tmp90022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										tmp90023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp90024 := Call(__e, tmp90023, Parse__shen_4_5formulae_6)

										tmp90025 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp90026 := Call(__e, tmp90025, Parse__shen_4_5formula_6)

										tmp90027 := Call(__e, tmp90022, tmp90024, tmp90026)

										__e.TailApply(tmp90019, tmp90021, tmp90027)
										return

									} else {
										tmp90018 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp90018)
										return

									}

								}, 1)

								tmp90034 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

								tmp90035 := Call(__e, tmp90034, NewStream1744)

								__e.TailApply(tmp90016, tmp90035)
								return

							}, 1)

							tmp90036 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							tmp90037 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							tmp90038 := Call(__e, tmp90037, Parse__shen_4_5formulae_6)

							tmp90039 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							tmp90040 := Call(__e, tmp90039, Parse__shen_4_5formulae_6)

							tmp90041 := Call(__e, tmp90036, tmp90038, tmp90040)

							__e.TailApply(tmp90015, tmp90041)
							return

						} else {
							tmp90014 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(tmp90014)
							return

						}

					} else {
						tmp90012 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp90012)
						return

					}

				}, 1)

				tmp90058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

				tmp90059 := Call(__e, tmp90058, V1746)

				tmp90060 := Call(__e, tmp90010, tmp90059)

				__e.TailApply(tmp89986, tmp90060)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp90085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90086 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp90087 := Call(__e, tmp90086, V1746)

		tmp90088 := Call(__e, tmp90085, tmp90087)

		var ifres90079 Obj

		if True == tmp90088 {
			tmp90081 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90082 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp90083 := Call(__e, tmp90082, V1746)

			tmp90084 := Call(__e, tmp90081, sym_b, tmp90083)

			var ifres90080 Obj

			if True == tmp90084 {
				ifres90080 = True

			} else {
				ifres90080 = False

			}

			ifres90079 = ifres90080

		} else {
			ifres90079 = False

		}

		var ifres90065 Obj

		if True == ifres90079 {
			tmp90068 := MakeNative(func(__e *ControlFlow) {
				NewStream1743 := __e.Get(1)
				_ = NewStream1743
				tmp90069 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp90070 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90071 := Call(__e, tmp90070, NewStream1743)

				__e.TailApply(tmp90069, tmp90071, sym_b)
				return

			}, 1)

			tmp90072 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			tmp90073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			tmp90074 := Call(__e, tmp90073, V1746)

			tmp90075 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			tmp90076 := Call(__e, tmp90075, V1746)

			tmp90077 := Call(__e, tmp90072, tmp90074, tmp90076)

			tmp90078 := Call(__e, tmp90068, tmp90077)

			ifres90065 = tmp90078

		} else {
			tmp90066 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90067 := Call(__e, tmp90066)

			ifres90065 = tmp90067

		}

		__e.TailApply(tmp89984, ifres90065)
		return

	}, 1)

	tmp90089 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premise_6, tmp89983)

	_ = tmp90089

	tmp90090 := MakeNative(func(__e *ControlFlow) {
		V1749 := __e.Get(1)
		_ = V1749
		tmp90091 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp90122 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90123 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90124 := Call(__e, tmp90123)

			tmp90125 := Call(__e, tmp90122, YaccParse, tmp90124)

			if True == tmp90125 {
				tmp90093 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					tmp90114 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp90115 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90116 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90117 := Call(__e, tmp90116)

					tmp90118 := Call(__e, tmp90115, tmp90117, Parse__shen_4_5formula_6)

					tmp90119 := Call(__e, tmp90114, tmp90118)

					if True == tmp90119 {
						tmp90096 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
							_ = Parse__shen_4_5semicolon_1symbol_6
							tmp90106 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp90107 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90108 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90109 := Call(__e, tmp90108)

							tmp90110 := Call(__e, tmp90107, tmp90109, Parse__shen_4_5semicolon_1symbol_6)

							tmp90111 := Call(__e, tmp90106, tmp90110)

							if True == tmp90111 {
								tmp90099 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp90100 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp90101 := Call(__e, tmp90100, Parse__shen_4_5semicolon_1symbol_6)

								tmp90102 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

								tmp90103 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp90104 := Call(__e, tmp90103, Parse__shen_4_5formula_6)

								tmp90105 := Call(__e, tmp90102, Nil, tmp90104)

								__e.TailApply(tmp90099, tmp90101, tmp90105)
								return

							} else {
								tmp90098 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp90098)
								return

							}

						}, 1)

						tmp90112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

						tmp90113 := Call(__e, tmp90112, Parse__shen_4_5formula_6)

						__e.TailApply(tmp90096, tmp90113)
						return

					} else {
						tmp90095 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp90095)
						return

					}

				}, 1)

				tmp90120 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

				tmp90121 := Call(__e, tmp90120, V1749)

				__e.TailApply(tmp90093, tmp90121)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp90126 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formulae_6 := __e.Get(1)
			_ = Parse__shen_4_5formulae_6
			tmp90179 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp90180 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90181 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90182 := Call(__e, tmp90181)

			tmp90183 := Call(__e, tmp90180, tmp90182, Parse__shen_4_5formulae_6)

			tmp90184 := Call(__e, tmp90179, tmp90183)

			if True == tmp90184 {
				tmp90175 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90176 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90177 := Call(__e, tmp90176, Parse__shen_4_5formulae_6)

				tmp90178 := Call(__e, tmp90175, tmp90177)

				var ifres90169 Obj

				if True == tmp90178 {
					tmp90171 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90172 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp90173 := Call(__e, tmp90172, Parse__shen_4_5formulae_6)

					tmp90174 := Call(__e, tmp90171, sym_6_6, tmp90173)

					var ifres90170 Obj

					if True == tmp90174 {
						ifres90170 = True

					} else {
						ifres90170 = False

					}

					ifres90169 = ifres90170

				} else {
					ifres90169 = False

				}

				if True == ifres90169 {
					tmp90131 := MakeNative(func(__e *ControlFlow) {
						NewStream1747 := __e.Get(1)
						_ = NewStream1747
						tmp90132 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							tmp90155 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp90156 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90157 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90158 := Call(__e, tmp90157)

							tmp90159 := Call(__e, tmp90156, tmp90158, Parse__shen_4_5formula_6)

							tmp90160 := Call(__e, tmp90155, tmp90159)

							if True == tmp90160 {
								tmp90135 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
									_ = Parse__shen_4_5semicolon_1symbol_6
									tmp90147 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									tmp90148 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp90149 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									tmp90150 := Call(__e, tmp90149)

									tmp90151 := Call(__e, tmp90148, tmp90150, Parse__shen_4_5semicolon_1symbol_6)

									tmp90152 := Call(__e, tmp90147, tmp90151)

									if True == tmp90152 {
										tmp90138 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										tmp90139 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										tmp90140 := Call(__e, tmp90139, Parse__shen_4_5semicolon_1symbol_6)

										tmp90141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										tmp90142 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp90143 := Call(__e, tmp90142, Parse__shen_4_5formulae_6)

										tmp90144 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										tmp90145 := Call(__e, tmp90144, Parse__shen_4_5formula_6)

										tmp90146 := Call(__e, tmp90141, tmp90143, tmp90145)

										__e.TailApply(tmp90138, tmp90140, tmp90146)
										return

									} else {
										tmp90137 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(tmp90137)
										return

									}

								}, 1)

								tmp90153 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

								tmp90154 := Call(__e, tmp90153, Parse__shen_4_5formula_6)

								__e.TailApply(tmp90135, tmp90154)
								return

							} else {
								tmp90134 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp90134)
								return

							}

						}, 1)

						tmp90161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

						tmp90162 := Call(__e, tmp90161, NewStream1747)

						__e.TailApply(tmp90132, tmp90162)
						return

					}, 1)

					tmp90163 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90164 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp90165 := Call(__e, tmp90164, Parse__shen_4_5formulae_6)

					tmp90166 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp90167 := Call(__e, tmp90166, Parse__shen_4_5formulae_6)

					tmp90168 := Call(__e, tmp90163, tmp90165, tmp90167)

					__e.TailApply(tmp90131, tmp90168)
					return

				} else {
					tmp90130 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp90130)
					return

				}

			} else {
				tmp90128 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp90128)
				return

			}

		}, 1)

		tmp90185 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

		tmp90186 := Call(__e, tmp90185, V1749)

		tmp90187 := Call(__e, tmp90126, tmp90186)

		__e.TailApply(tmp90091, tmp90187)
		return

	}, 1)

	tmp90188 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5conclusion_6, tmp90090)

	_ = tmp90188

	tmp90189 := MakeNative(func(__e *ControlFlow) {
		V1752 := __e.Get(1)
		_ = V1752
		V1753 := __e.Get(2)
		_ = V1753
		tmp90190 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

		__e.TailApply(tmp90190, V1752, V1753)
		return

	}, 2)

	tmp90191 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sequent, tmp90189)

	_ = tmp90191

	tmp90192 := MakeNative(func(__e *ControlFlow) {
		V1755 := __e.Get(1)
		_ = V1755
		tmp90193 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp90234 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90235 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90236 := Call(__e, tmp90235)

			tmp90237 := Call(__e, tmp90234, YaccParse, tmp90236)

			if True == tmp90237 {
				tmp90195 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					tmp90211 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90212 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90213 := Call(__e, tmp90212)

					tmp90214 := Call(__e, tmp90211, YaccParse, tmp90213)

					if True == tmp90214 {
						tmp90197 := MakeNative(func(__e *ControlFlow) {
							Parse___5e_6 := __e.Get(1)
							_ = Parse___5e_6
							tmp90203 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp90204 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90205 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90206 := Call(__e, tmp90205)

							tmp90207 := Call(__e, tmp90204, tmp90206, Parse___5e_6)

							tmp90208 := Call(__e, tmp90203, tmp90207)

							if True == tmp90208 {
								tmp90200 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp90201 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp90202 := Call(__e, tmp90201, Parse___5e_6)

								__e.TailApply(tmp90200, tmp90202, Nil)
								return

							} else {
								tmp90199 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp90199)
								return

							}

						}, 1)

						tmp90209 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

						tmp90210 := Call(__e, tmp90209, V1755)

						__e.TailApply(tmp90197, tmp90210)
						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				tmp90215 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					tmp90225 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp90226 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90227 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90228 := Call(__e, tmp90227)

					tmp90229 := Call(__e, tmp90226, tmp90228, Parse__shen_4_5formula_6)

					tmp90230 := Call(__e, tmp90225, tmp90229)

					if True == tmp90230 {
						tmp90218 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp90219 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90220 := Call(__e, tmp90219, Parse__shen_4_5formula_6)

						tmp90221 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp90222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp90223 := Call(__e, tmp90222, Parse__shen_4_5formula_6)

						tmp90224 := Call(__e, tmp90221, tmp90223, Nil)

						__e.TailApply(tmp90218, tmp90220, tmp90224)
						return

					} else {
						tmp90217 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp90217)
						return

					}

				}, 1)

				tmp90231 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

				tmp90232 := Call(__e, tmp90231, V1755)

				tmp90233 := Call(__e, tmp90215, tmp90232)

				__e.TailApply(tmp90195, tmp90233)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp90238 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formula_6 := __e.Get(1)
			_ = Parse__shen_4_5formula_6
			tmp90272 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp90273 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90274 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90275 := Call(__e, tmp90274)

			tmp90276 := Call(__e, tmp90273, tmp90275, Parse__shen_4_5formula_6)

			tmp90277 := Call(__e, tmp90272, tmp90276)

			if True == tmp90277 {
				tmp90241 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5comma_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5comma_1symbol_6
					tmp90264 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp90265 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90266 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90267 := Call(__e, tmp90266)

					tmp90268 := Call(__e, tmp90265, tmp90267, Parse__shen_4_5comma_1symbol_6)

					tmp90269 := Call(__e, tmp90264, tmp90268)

					if True == tmp90269 {
						tmp90244 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formulae_6 := __e.Get(1)
							_ = Parse__shen_4_5formulae_6
							tmp90256 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp90257 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90258 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90259 := Call(__e, tmp90258)

							tmp90260 := Call(__e, tmp90257, tmp90259, Parse__shen_4_5formulae_6)

							tmp90261 := Call(__e, tmp90256, tmp90260)

							if True == tmp90261 {
								tmp90247 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp90248 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp90249 := Call(__e, tmp90248, Parse__shen_4_5formulae_6)

								tmp90250 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp90251 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp90252 := Call(__e, tmp90251, Parse__shen_4_5formula_6)

								tmp90253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp90254 := Call(__e, tmp90253, Parse__shen_4_5formulae_6)

								tmp90255 := Call(__e, tmp90250, tmp90252, tmp90254)

								__e.TailApply(tmp90247, tmp90249, tmp90255)
								return

							} else {
								tmp90246 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp90246)
								return

							}

						}, 1)

						tmp90262 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

						tmp90263 := Call(__e, tmp90262, Parse__shen_4_5comma_1symbol_6)

						__e.TailApply(tmp90244, tmp90263)
						return

					} else {
						tmp90243 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp90243)
						return

					}

				}, 1)

				tmp90270 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comma_1symbol_6)

				tmp90271 := Call(__e, tmp90270, Parse__shen_4_5formula_6)

				__e.TailApply(tmp90241, tmp90271)
				return

			} else {
				tmp90240 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp90240)
				return

			}

		}, 1)

		tmp90278 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

		tmp90279 := Call(__e, tmp90278, V1755)

		tmp90280 := Call(__e, tmp90238, tmp90279)

		__e.TailApply(tmp90193, tmp90280)
		return

	}, 1)

	tmp90281 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formulae_6, tmp90192)

	_ = tmp90281

	tmp90282 := MakeNative(func(__e *ControlFlow) {
		V1757 := __e.Get(1)
		_ = V1757
		tmp90303 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp90305 := Call(__e, tmp90304, V1757)

		tmp90306 := Call(__e, tmp90303, tmp90305)

		if True == tmp90306 {
			tmp90285 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp90297 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp90298 := Call(__e, PrimNS1Value(symns2_1value), symintern)

				tmp90299 := Call(__e, tmp90298, MakeString(","))

				tmp90300 := Call(__e, tmp90297, Parse__X, tmp90299)

				if True == tmp90300 {
					tmp90288 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90289 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90290 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90291 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp90292 := Call(__e, tmp90291, V1757)

					tmp90293 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp90294 := Call(__e, tmp90293, V1757)

					tmp90295 := Call(__e, tmp90290, tmp90292, tmp90294)

					tmp90296 := Call(__e, tmp90289, tmp90295)

					__e.TailApply(tmp90288, tmp90296, symshen_4skip)
					return

				} else {
					tmp90287 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp90287)
					return

				}

			}, 1)

			tmp90301 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp90302 := Call(__e, tmp90301, V1757)

			__e.TailApply(tmp90285, tmp90302)
			return

		} else {
			tmp90284 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp90284)
			return

		}

	}, 1)

	tmp90307 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_1symbol_6, tmp90282)

	_ = tmp90307

	tmp90308 := MakeNative(func(__e *ControlFlow) {
		V1760 := __e.Get(1)
		_ = V1760
		tmp90309 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			tmp90327 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90328 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90329 := Call(__e, tmp90328)

			tmp90330 := Call(__e, tmp90327, YaccParse, tmp90329)

			if True == tmp90330 {
				tmp90311 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					tmp90319 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					tmp90320 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90321 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					tmp90322 := Call(__e, tmp90321)

					tmp90323 := Call(__e, tmp90320, tmp90322, Parse__shen_4_5expr_6)

					tmp90324 := Call(__e, tmp90319, tmp90323)

					if True == tmp90324 {
						tmp90314 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						tmp90315 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90316 := Call(__e, tmp90315, Parse__shen_4_5expr_6)

						tmp90317 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						tmp90318 := Call(__e, tmp90317, Parse__shen_4_5expr_6)

						__e.TailApply(tmp90314, tmp90316, tmp90318)
						return

					} else {
						tmp90313 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(tmp90313)
						return

					}

				}, 1)

				tmp90325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

				tmp90326 := Call(__e, tmp90325, V1760)

				__e.TailApply(tmp90311, tmp90326)
				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		tmp90331 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			tmp90381 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp90382 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90383 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90384 := Call(__e, tmp90383)

			tmp90385 := Call(__e, tmp90382, tmp90384, Parse__shen_4_5expr_6)

			tmp90386 := Call(__e, tmp90381, tmp90385)

			if True == tmp90386 {
				tmp90377 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90378 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90379 := Call(__e, tmp90378, Parse__shen_4_5expr_6)

				tmp90380 := Call(__e, tmp90377, tmp90379)

				var ifres90371 Obj

				if True == tmp90380 {
					tmp90373 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90374 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					tmp90375 := Call(__e, tmp90374, Parse__shen_4_5expr_6)

					tmp90376 := Call(__e, tmp90373, sym_h, tmp90375)

					var ifres90372 Obj

					if True == tmp90376 {
						ifres90372 = True

					} else {
						ifres90372 = False

					}

					ifres90371 = ifres90372

				} else {
					ifres90371 = False

				}

				if True == ifres90371 {
					tmp90336 := MakeNative(func(__e *ControlFlow) {
						NewStream1758 := __e.Get(1)
						_ = NewStream1758
						tmp90337 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5type_6 := __e.Get(1)
							_ = Parse__shen_4_5type_6
							tmp90357 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							tmp90358 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90359 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							tmp90360 := Call(__e, tmp90359)

							tmp90361 := Call(__e, tmp90358, tmp90360, Parse__shen_4_5type_6)

							tmp90362 := Call(__e, tmp90357, tmp90361)

							if True == tmp90362 {
								tmp90340 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								tmp90341 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp90342 := Call(__e, tmp90341, Parse__shen_4_5type_6)

								tmp90343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp90344 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

								tmp90345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp90346 := Call(__e, tmp90345, Parse__shen_4_5expr_6)

								tmp90347 := Call(__e, tmp90344, tmp90346)

								tmp90348 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp90349 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								tmp90350 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

								tmp90351 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								tmp90352 := Call(__e, tmp90351, Parse__shen_4_5type_6)

								tmp90353 := Call(__e, tmp90350, tmp90352)

								tmp90354 := Call(__e, tmp90349, tmp90353, Nil)

								tmp90355 := Call(__e, tmp90348, sym_h, tmp90354)

								tmp90356 := Call(__e, tmp90343, tmp90347, tmp90355)

								__e.TailApply(tmp90340, tmp90342, tmp90356)
								return

							} else {
								tmp90339 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(tmp90339)
								return

							}

						}, 1)

						tmp90363 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5type_6)

						tmp90364 := Call(__e, tmp90363, NewStream1758)

						__e.TailApply(tmp90337, tmp90364)
						return

					}, 1)

					tmp90365 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90366 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp90367 := Call(__e, tmp90366, Parse__shen_4_5expr_6)

					tmp90368 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp90369 := Call(__e, tmp90368, Parse__shen_4_5expr_6)

					tmp90370 := Call(__e, tmp90365, tmp90367, tmp90369)

					__e.TailApply(tmp90336, tmp90370)
					return

				} else {
					tmp90335 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp90335)
					return

				}

			} else {
				tmp90333 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp90333)
				return

			}

		}, 1)

		tmp90387 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

		tmp90388 := Call(__e, tmp90387, V1760)

		tmp90389 := Call(__e, tmp90331, tmp90388)

		__e.TailApply(tmp90309, tmp90389)
		return

	}, 1)

	tmp90390 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formula_6, tmp90308)

	_ = tmp90390

	tmp90391 := MakeNative(func(__e *ControlFlow) {
		V1762 := __e.Get(1)
		_ = V1762
		tmp90392 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			tmp90402 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			tmp90403 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90404 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			tmp90405 := Call(__e, tmp90404)

			tmp90406 := Call(__e, tmp90403, tmp90405, Parse__shen_4_5expr_6)

			tmp90407 := Call(__e, tmp90402, tmp90406)

			if True == tmp90407 {
				tmp90395 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				tmp90396 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90397 := Call(__e, tmp90396, Parse__shen_4_5expr_6)

				tmp90398 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

				tmp90399 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				tmp90400 := Call(__e, tmp90399, Parse__shen_4_5expr_6)

				tmp90401 := Call(__e, tmp90398, tmp90400)

				__e.TailApply(tmp90395, tmp90397, tmp90401)
				return

			} else {
				tmp90394 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(tmp90394)
				return

			}

		}, 1)

		tmp90408 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

		tmp90409 := Call(__e, tmp90408, V1762)

		__e.TailApply(tmp90392, tmp90409)
		return

	}, 1)

	tmp90410 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5type_6, tmp90391)

	_ = tmp90410

	tmp90411 := MakeNative(func(__e *ControlFlow) {
		V1764 := __e.Get(1)
		_ = V1764
		tmp90430 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90431 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp90432 := Call(__e, tmp90431, V1764)

		tmp90433 := Call(__e, tmp90430, tmp90432)

		if True == tmp90433 {
			tmp90414 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp90426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

				tmp90427 := Call(__e, tmp90426, Parse__X)

				if True == tmp90427 {
					tmp90417 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90418 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90419 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90420 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp90421 := Call(__e, tmp90420, V1764)

					tmp90422 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp90423 := Call(__e, tmp90422, V1764)

					tmp90424 := Call(__e, tmp90419, tmp90421, tmp90423)

					tmp90425 := Call(__e, tmp90418, tmp90424)

					__e.TailApply(tmp90417, tmp90425, Parse__X)
					return

				} else {
					tmp90416 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp90416)
					return

				}

			}, 1)

			tmp90428 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp90429 := Call(__e, tmp90428, V1764)

			__e.TailApply(tmp90414, tmp90429)
			return

		} else {
			tmp90413 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp90413)
			return

		}

	}, 1)

	tmp90434 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5doubleunderline_6, tmp90411)

	_ = tmp90434

	tmp90435 := MakeNative(func(__e *ControlFlow) {
		V1766 := __e.Get(1)
		_ = V1766
		tmp90454 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90455 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		tmp90456 := Call(__e, tmp90455, V1766)

		tmp90457 := Call(__e, tmp90454, tmp90456)

		if True == tmp90457 {
			tmp90438 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				tmp90450 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

				tmp90451 := Call(__e, tmp90450, Parse__X)

				if True == tmp90451 {
					tmp90441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90442 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90443 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					tmp90444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					tmp90445 := Call(__e, tmp90444, V1766)

					tmp90446 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					tmp90447 := Call(__e, tmp90446, V1766)

					tmp90448 := Call(__e, tmp90443, tmp90445, tmp90447)

					tmp90449 := Call(__e, tmp90442, tmp90448)

					__e.TailApply(tmp90441, tmp90449, Parse__X)
					return

				} else {
					tmp90440 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(tmp90440)
					return

				}

			}, 1)

			tmp90452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			tmp90453 := Call(__e, tmp90452, V1766)

			__e.TailApply(tmp90438, tmp90453)
			return

		} else {
			tmp90437 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(tmp90437)
			return

		}

	}, 1)

	tmp90458 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleunderline_6, tmp90435)

	_ = tmp90458

	tmp90459 := MakeNative(func(__e *ControlFlow) {
		V1768 := __e.Get(1)
		_ = V1768
		tmp90466 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp90467 := Call(__e, tmp90466, V1768)

		if True == tmp90467 {
			tmp90462 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sh_2)

			tmp90463 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			tmp90464 := Call(__e, tmp90463, V1768)

			tmp90465 := Call(__e, tmp90462, tmp90464)

			if True == tmp90465 {
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

	tmp90468 := Call(__e, PrimNS1Value(symns2_1set), symshen_4singleunderline_2, tmp90459)

	_ = tmp90468

	tmp90469 := MakeNative(func(__e *ControlFlow) {
		V1770 := __e.Get(1)
		_ = V1770
		tmp90481 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90482 := Call(__e, tmp90481, MakeString("_"), V1770)

		if True == tmp90482 {
			__e.Return(True)
			return
		} else {
			tmp90477 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90478 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			tmp90479 := Call(__e, tmp90478, V1770, MakeNumber(0))

			tmp90480 := Call(__e, tmp90477, tmp90479, MakeString("_"))

			if True == tmp90480 {
				tmp90473 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sh_2)

				tmp90474 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp90475 := Call(__e, tmp90474, V1770)

				tmp90476 := Call(__e, tmp90473, tmp90475)

				if True == tmp90476 {
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

		}

	}, 1)

	tmp90483 := Call(__e, PrimNS1Value(symns2_1set), symshen_4sh_2, tmp90469)

	_ = tmp90483

	tmp90484 := MakeNative(func(__e *ControlFlow) {
		V1772 := __e.Get(1)
		_ = V1772
		tmp90491 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		tmp90492 := Call(__e, tmp90491, V1772)

		if True == tmp90492 {
			tmp90487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dh_2)

			tmp90488 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			tmp90489 := Call(__e, tmp90488, V1772)

			tmp90490 := Call(__e, tmp90487, tmp90489)

			if True == tmp90490 {
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

	tmp90493 := Call(__e, PrimNS1Value(symns2_1set), symshen_4doubleunderline_2, tmp90484)

	_ = tmp90493

	tmp90494 := MakeNative(func(__e *ControlFlow) {
		V1774 := __e.Get(1)
		_ = V1774
		tmp90506 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90507 := Call(__e, tmp90506, MakeString("="), V1774)

		if True == tmp90507 {
			__e.Return(True)
			return
		} else {
			tmp90502 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90503 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			tmp90504 := Call(__e, tmp90503, V1774, MakeNumber(0))

			tmp90505 := Call(__e, tmp90502, tmp90504, MakeString("="))

			if True == tmp90505 {
				tmp90498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dh_2)

				tmp90499 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				tmp90500 := Call(__e, tmp90499, V1774)

				tmp90501 := Call(__e, tmp90498, tmp90500)

				if True == tmp90501 {
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

		}

	}, 1)

	tmp90508 := Call(__e, PrimNS1Value(symns2_1set), symshen_4dh_2, tmp90494)

	_ = tmp90508

	tmp90509 := MakeNative(func(__e *ControlFlow) {
		V1777 := __e.Get(1)
		_ = V1777
		V1778 := __e.Get(2)
		_ = V1778
		tmp90510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remember_1datatype)

		tmp90511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog)

		tmp90512 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

		tmp90513 := Call(__e, tmp90512, V1777, V1778)

		tmp90514 := Call(__e, tmp90511, tmp90513)

		__e.TailApply(tmp90510, tmp90514)
		return

	}, 2)

	tmp90515 := Call(__e, PrimNS1Value(symns2_1set), symshen_4process_1datatype, tmp90509)

	_ = tmp90515

	tmp90516 := MakeNative(func(__e *ControlFlow) {
		V1784 := __e.Get(1)
		_ = V1784
		tmp90536 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90537 := Call(__e, tmp90536, V1784)

		if True == tmp90537 {
			tmp90519 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp90520 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			tmp90521 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90522 := Call(__e, tmp90521, V1784)

			tmp90523 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp90524 := Call(__e, tmp90523, symshen_4_ddatatypes_d)

			tmp90525 := Call(__e, tmp90520, tmp90522, tmp90524)

			tmp90526 := Call(__e, tmp90519, symshen_4_ddatatypes_d, tmp90525)

			_ = tmp90526

			tmp90527 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp90528 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			tmp90529 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90530 := Call(__e, tmp90529, V1784)

			tmp90531 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp90532 := Call(__e, tmp90531, symshen_4_dalldatatypes_d)

			tmp90533 := Call(__e, tmp90528, tmp90530, tmp90532)

			tmp90534 := Call(__e, tmp90527, symshen_4_dalldatatypes_d, tmp90533)

			_ = tmp90534

			tmp90535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(tmp90535, V1784)
			return

		} else {
			tmp90518 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp90518, symshen_4remember_1datatype)
			return

		}

	}, 1)

	tmp90538 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remember_1datatype, tmp90516)

	_ = tmp90538

	tmp90539 := MakeNative(func(__e *ControlFlow) {
		V1789 := __e.Get(1)
		_ = V1789
		V1790 := __e.Get(2)
		_ = V1790
		tmp90598 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90599 := Call(__e, tmp90598, Nil, V1790)

		if True == tmp90599 {
			__e.Return(Nil)
			return
		} else {
			tmp90596 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90597 := Call(__e, tmp90596, V1790)

			var ifres90582 Obj

			if True == tmp90597 {
				tmp90592 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

				tmp90593 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90594 := Call(__e, tmp90593, V1790)

				tmp90595 := Call(__e, tmp90592, tmp90594)

				var ifres90584 Obj

				if True == tmp90595 {
					tmp90586 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp90587 := Call(__e, PrimNS1Value(symns2_1value), symfst)

					tmp90588 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90589 := Call(__e, tmp90588, V1790)

					tmp90590 := Call(__e, tmp90587, tmp90589)

					tmp90591 := Call(__e, tmp90586, symshen_4single, tmp90590)

					var ifres90585 Obj

					if True == tmp90591 {
						ifres90585 = True

					} else {
						ifres90585 = False

					}

					ifres90584 = ifres90585

				} else {
					ifres90584 = False

				}

				var ifres90583 Obj

				if True == ifres90584 {
					ifres90583 = True

				} else {
					ifres90583 = False

				}

				ifres90582 = ifres90583

			} else {
				ifres90582 = False

			}

			if True == ifres90582 {
				tmp90571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90572 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause)

				tmp90573 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				tmp90574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90575 := Call(__e, tmp90574, V1790)

				tmp90576 := Call(__e, tmp90573, tmp90575)

				tmp90577 := Call(__e, tmp90572, V1789, tmp90576)

				tmp90578 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

				tmp90579 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90580 := Call(__e, tmp90579, V1790)

				tmp90581 := Call(__e, tmp90578, V1789, tmp90580)

				__e.TailApply(tmp90571, tmp90577, tmp90581)
				return

			} else {
				tmp90569 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90570 := Call(__e, tmp90569, V1790)

				var ifres90555 Obj

				if True == tmp90570 {
					tmp90565 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					tmp90566 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90567 := Call(__e, tmp90566, V1790)

					tmp90568 := Call(__e, tmp90565, tmp90567)

					var ifres90557 Obj

					if True == tmp90568 {
						tmp90559 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp90560 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						tmp90561 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90562 := Call(__e, tmp90561, V1790)

						tmp90563 := Call(__e, tmp90560, tmp90562)

						tmp90564 := Call(__e, tmp90559, symshen_4double, tmp90563)

						var ifres90558 Obj

						if True == tmp90564 {
							ifres90558 = True

						} else {
							ifres90558 = False

						}

						ifres90557 = ifres90558

					} else {
						ifres90557 = False

					}

					var ifres90556 Obj

					if True == ifres90557 {
						ifres90556 = True

					} else {
						ifres90556 = False

					}

					ifres90555 = ifres90556

				} else {
					ifres90555 = False

				}

				if True == ifres90555 {
					tmp90544 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

					tmp90545 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					tmp90546 := Call(__e, PrimNS1Value(symns2_1value), symshen_4double_1_6singles)

					tmp90547 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					tmp90548 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90549 := Call(__e, tmp90548, V1790)

					tmp90550 := Call(__e, tmp90547, tmp90549)

					tmp90551 := Call(__e, tmp90546, tmp90550)

					tmp90552 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90553 := Call(__e, tmp90552, V1790)

					tmp90554 := Call(__e, tmp90545, tmp90551, tmp90553)

					__e.TailApply(tmp90544, V1789, tmp90554)
					return

				} else {
					tmp90543 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp90543, symshen_4rules_1_6horn_1clauses)
					return

				}

			}

		}

	}, 2)

	tmp90600 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rules_1_6horn_1clauses, tmp90539)

	_ = tmp90600

	tmp90601 := MakeNative(func(__e *ControlFlow) {
		V1792 := __e.Get(1)
		_ = V1792
		tmp90602 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90603 := Call(__e, PrimNS1Value(symns2_1value), symshen_4right_1rule)

		tmp90604 := Call(__e, tmp90603, V1792)

		tmp90605 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4left_1rule)

		tmp90607 := Call(__e, tmp90606, V1792)

		tmp90608 := Call(__e, tmp90605, tmp90607, Nil)

		__e.TailApply(tmp90602, tmp90604, tmp90608)
		return

	}, 1)

	tmp90609 := Call(__e, PrimNS1Value(symns2_1set), symshen_4double_1_6singles, tmp90601)

	_ = tmp90609

	tmp90610 := MakeNative(func(__e *ControlFlow) {
		V1794 := __e.Get(1)
		_ = V1794
		tmp90611 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

		__e.TailApply(tmp90611, symshen_4single, V1794)
		return

	}, 1)

	tmp90612 := Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1rule, tmp90610)

	_ = tmp90612

	tmp90613 := MakeNative(func(__e *ControlFlow) {
		V1796 := __e.Get(1)
		_ = V1796
		tmp90700 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90701 := Call(__e, tmp90700, V1796)

		var ifres90654 Obj

		if True == tmp90701 {
			tmp90696 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90697 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90698 := Call(__e, tmp90697, V1796)

			tmp90699 := Call(__e, tmp90696, tmp90698)

			var ifres90656 Obj

			if True == tmp90699 {
				tmp90690 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90691 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90692 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90693 := Call(__e, tmp90692, V1796)

				tmp90694 := Call(__e, tmp90691, tmp90693)

				tmp90695 := Call(__e, tmp90690, tmp90694)

				var ifres90658 Obj

				if True == tmp90695 {
					tmp90682 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					tmp90683 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90684 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90685 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90686 := Call(__e, tmp90685, V1796)

					tmp90687 := Call(__e, tmp90684, tmp90686)

					tmp90688 := Call(__e, tmp90683, tmp90687)

					tmp90689 := Call(__e, tmp90682, tmp90688)

					var ifres90660 Obj

					if True == tmp90689 {
						tmp90672 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp90673 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						tmp90674 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90675 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90676 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90677 := Call(__e, tmp90676, V1796)

						tmp90678 := Call(__e, tmp90675, tmp90677)

						tmp90679 := Call(__e, tmp90674, tmp90678)

						tmp90680 := Call(__e, tmp90673, tmp90679)

						tmp90681 := Call(__e, tmp90672, Nil, tmp90680)

						var ifres90662 Obj

						if True == tmp90681 {
							tmp90664 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp90665 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp90666 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp90667 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp90668 := Call(__e, tmp90667, V1796)

							tmp90669 := Call(__e, tmp90666, tmp90668)

							tmp90670 := Call(__e, tmp90665, tmp90669)

							tmp90671 := Call(__e, tmp90664, Nil, tmp90670)

							var ifres90663 Obj

							if True == tmp90671 {
								ifres90663 = True

							} else {
								ifres90663 = False

							}

							ifres90662 = ifres90663

						} else {
							ifres90662 = False

						}

						var ifres90661 Obj

						if True == ifres90662 {
							ifres90661 = True

						} else {
							ifres90661 = False

						}

						ifres90660 = ifres90661

					} else {
						ifres90660 = False

					}

					var ifres90659 Obj

					if True == ifres90660 {
						ifres90659 = True

					} else {
						ifres90659 = False

					}

					ifres90658 = ifres90659

				} else {
					ifres90658 = False

				}

				var ifres90657 Obj

				if True == ifres90658 {
					ifres90657 = True

				} else {
					ifres90657 = False

				}

				ifres90656 = ifres90657

			} else {
				ifres90656 = False

			}

			var ifres90655 Obj

			if True == ifres90656 {
				ifres90655 = True

			} else {
				ifres90655 = False

			}

			ifres90654 = ifres90655

		} else {
			ifres90654 = False

		}

		if True == ifres90654 {
			tmp90616 := MakeNative(func(__e *ControlFlow) {
				Q := __e.Get(1)
				_ = Q
				tmp90617 := MakeNative(func(__e *ControlFlow) {
					NewConclusion := __e.Get(1)
					_ = NewConclusion
					tmp90618 := MakeNative(func(__e *ControlFlow) {
						NewPremises := __e.Get(1)
						_ = NewPremises
						tmp90619 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

						tmp90620 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp90621 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp90622 := Call(__e, tmp90621, V1796)

						tmp90623 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp90624 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp90625 := Call(__e, tmp90624, NewConclusion, Nil)

						tmp90626 := Call(__e, tmp90623, NewPremises, tmp90625)

						tmp90627 := Call(__e, tmp90620, tmp90622, tmp90626)

						__e.TailApply(tmp90619, symshen_4single, tmp90627)
						return

					}, 1)

					tmp90628 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp90629 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

					tmp90630 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					tmp90631 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						tmp90632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4right_1_6left)

						__e.TailApply(tmp90632, X)
						return

					}, 1)

					tmp90633 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90634 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90635 := Call(__e, tmp90634, V1796)

					tmp90636 := Call(__e, tmp90633, tmp90635)

					tmp90637 := Call(__e, tmp90630, tmp90631, tmp90636)

					tmp90638 := Call(__e, tmp90629, tmp90637, Q)

					tmp90639 := Call(__e, tmp90628, tmp90638, Nil)

					__e.TailApply(tmp90618, tmp90639)
					return

				}, 1)

				tmp90640 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

				tmp90641 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90642 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				tmp90643 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90644 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90646 := Call(__e, tmp90645, V1796)

				tmp90647 := Call(__e, tmp90644, tmp90646)

				tmp90648 := Call(__e, tmp90643, tmp90647)

				tmp90649 := Call(__e, tmp90642, tmp90648)

				tmp90650 := Call(__e, tmp90641, tmp90649, Nil)

				tmp90651 := Call(__e, tmp90640, tmp90650, Q)

				__e.TailApply(tmp90617, tmp90651)
				return

			}, 1)

			tmp90652 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp90653 := Call(__e, tmp90652, symQv)

			__e.TailApply(tmp90616, tmp90653)
			return

		} else {
			tmp90615 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp90615, symshen_4left_1rule)
			return

		}

	}, 1)

	tmp90702 := Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1rule, tmp90613)

	_ = tmp90702

	tmp90703 := MakeNative(func(__e *ControlFlow) {
		V1802 := __e.Get(1)
		_ = V1802
		tmp90713 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		tmp90714 := Call(__e, tmp90713, V1802)

		var ifres90707 Obj

		if True == tmp90714 {
			tmp90709 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90710 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			tmp90711 := Call(__e, tmp90710, V1802)

			tmp90712 := Call(__e, tmp90709, Nil, tmp90711)

			var ifres90708 Obj

			if True == tmp90712 {
				ifres90708 = True

			} else {
				ifres90708 = False

			}

			ifres90707 = ifres90708

		} else {
			ifres90707 = False

		}

		if True == ifres90707 {
			tmp90706 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			__e.TailApply(tmp90706, V1802)
			return

		} else {
			tmp90705 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp90705, MakeString("syntax error with ==========\n"))
			return

		}

	}, 1)

	tmp90715 := Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1_6left, tmp90703)

	_ = tmp90715

	tmp90716 := MakeNative(func(__e *ControlFlow) {
		V1805 := __e.Get(1)
		_ = V1805
		V1806 := __e.Get(2)
		_ = V1806
		tmp90784 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90785 := Call(__e, tmp90784, V1806)

		var ifres90750 Obj

		if True == tmp90785 {
			tmp90780 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90781 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90782 := Call(__e, tmp90781, V1806)

			tmp90783 := Call(__e, tmp90780, tmp90782)

			var ifres90752 Obj

			if True == tmp90783 {
				tmp90774 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90775 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90776 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90777 := Call(__e, tmp90776, V1806)

				tmp90778 := Call(__e, tmp90775, tmp90777)

				tmp90779 := Call(__e, tmp90774, tmp90778)

				var ifres90754 Obj

				if True == tmp90779 {
					tmp90766 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					tmp90767 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp90768 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90769 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90770 := Call(__e, tmp90769, V1806)

					tmp90771 := Call(__e, tmp90768, tmp90770)

					tmp90772 := Call(__e, tmp90767, tmp90771)

					tmp90773 := Call(__e, tmp90766, tmp90772)

					var ifres90756 Obj

					if True == tmp90773 {
						tmp90758 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp90759 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90760 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90761 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90762 := Call(__e, tmp90761, V1806)

						tmp90763 := Call(__e, tmp90760, tmp90762)

						tmp90764 := Call(__e, tmp90759, tmp90763)

						tmp90765 := Call(__e, tmp90758, Nil, tmp90764)

						var ifres90757 Obj

						if True == tmp90765 {
							ifres90757 = True

						} else {
							ifres90757 = False

						}

						ifres90756 = ifres90757

					} else {
						ifres90756 = False

					}

					var ifres90755 Obj

					if True == ifres90756 {
						ifres90755 = True

					} else {
						ifres90755 = False

					}

					ifres90754 = ifres90755

				} else {
					ifres90754 = False

				}

				var ifres90753 Obj

				if True == ifres90754 {
					ifres90753 = True

				} else {
					ifres90753 = False

				}

				ifres90752 = ifres90753

			} else {
				ifres90752 = False

			}

			var ifres90751 Obj

			if True == ifres90752 {
				ifres90751 = True

			} else {
				ifres90751 = False

			}

			ifres90750 = ifres90751

		} else {
			ifres90750 = False

		}

		if True == ifres90750 {
			tmp90719 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause_1head)

			tmp90721 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			tmp90722 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90723 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90724 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90725 := Call(__e, tmp90724, V1806)

			tmp90726 := Call(__e, tmp90723, tmp90725)

			tmp90727 := Call(__e, tmp90722, tmp90726)

			tmp90728 := Call(__e, tmp90721, tmp90727)

			tmp90729 := Call(__e, tmp90720, V1805, tmp90728)

			tmp90730 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90731 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90732 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause_1body)

			tmp90733 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90734 := Call(__e, tmp90733, V1806)

			tmp90735 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90736 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90737 := Call(__e, tmp90736, V1806)

			tmp90738 := Call(__e, tmp90735, tmp90737)

			tmp90739 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			tmp90740 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90741 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90742 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90743 := Call(__e, tmp90742, V1806)

			tmp90744 := Call(__e, tmp90741, tmp90743)

			tmp90745 := Call(__e, tmp90740, tmp90744)

			tmp90746 := Call(__e, tmp90739, tmp90745)

			tmp90747 := Call(__e, tmp90732, tmp90734, tmp90738, tmp90746)

			tmp90748 := Call(__e, tmp90731, tmp90747, Nil)

			tmp90749 := Call(__e, tmp90730, sym_h_1, tmp90748)

			__e.TailApply(tmp90719, tmp90729, tmp90749)
			return

		} else {
			tmp90718 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp90718, symshen_4rule_1_6horn_1clause)
			return

		}

	}, 2)

	tmp90786 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause, tmp90716)

	_ = tmp90786

	tmp90787 := MakeNative(func(__e *ControlFlow) {
		V1809 := __e.Get(1)
		_ = V1809
		V1810 := __e.Get(2)
		_ = V1810
		tmp90788 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90789 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90790 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mode_1ify)

		tmp90791 := Call(__e, tmp90790, V1810)

		tmp90792 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90793 := Call(__e, tmp90792, symContext__1957, Nil)

		tmp90794 := Call(__e, tmp90789, tmp90791, tmp90793)

		__e.TailApply(tmp90788, V1809, tmp90794)
		return

	}, 2)

	tmp90795 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1head, tmp90787)

	_ = tmp90795

	tmp90796 := MakeNative(func(__e *ControlFlow) {
		V1812 := __e.Get(1)
		_ = V1812
		tmp90855 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp90856 := Call(__e, tmp90855, V1812)

		var ifres90823 Obj

		if True == tmp90856 {
			tmp90851 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90852 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90853 := Call(__e, tmp90852, V1812)

			tmp90854 := Call(__e, tmp90851, tmp90853)

			var ifres90825 Obj

			if True == tmp90854 {
				tmp90845 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp90846 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90847 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90848 := Call(__e, tmp90847, V1812)

				tmp90849 := Call(__e, tmp90846, tmp90848)

				tmp90850 := Call(__e, tmp90845, sym_h, tmp90849)

				var ifres90827 Obj

				if True == tmp90850 {
					tmp90839 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp90840 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90841 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp90842 := Call(__e, tmp90841, V1812)

					tmp90843 := Call(__e, tmp90840, tmp90842)

					tmp90844 := Call(__e, tmp90839, tmp90843)

					var ifres90829 Obj

					if True == tmp90844 {
						tmp90831 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp90832 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90833 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90834 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp90835 := Call(__e, tmp90834, V1812)

						tmp90836 := Call(__e, tmp90833, tmp90835)

						tmp90837 := Call(__e, tmp90832, tmp90836)

						tmp90838 := Call(__e, tmp90831, Nil, tmp90837)

						var ifres90830 Obj

						if True == tmp90838 {
							ifres90830 = True

						} else {
							ifres90830 = False

						}

						ifres90829 = ifres90830

					} else {
						ifres90829 = False

					}

					var ifres90828 Obj

					if True == ifres90829 {
						ifres90828 = True

					} else {
						ifres90828 = False

					}

					ifres90827 = ifres90828

				} else {
					ifres90827 = False

				}

				var ifres90826 Obj

				if True == ifres90827 {
					ifres90826 = True

				} else {
					ifres90826 = False

				}

				ifres90825 = ifres90826

			} else {
				ifres90825 = False

			}

			var ifres90824 Obj

			if True == ifres90825 {
				ifres90824 = True

			} else {
				ifres90824 = False

			}

			ifres90823 = ifres90824

		} else {
			ifres90823 = False

		}

		if True == ifres90823 {
			tmp90798 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90799 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90800 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90801 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90802 := Call(__e, tmp90801, V1812)

			tmp90803 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90804 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90805 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90806 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90807 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp90808 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90809 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp90810 := Call(__e, tmp90809, V1812)

			tmp90811 := Call(__e, tmp90808, tmp90810)

			tmp90812 := Call(__e, tmp90807, tmp90811)

			tmp90813 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90814 := Call(__e, tmp90813, sym_7, Nil)

			tmp90815 := Call(__e, tmp90806, tmp90812, tmp90814)

			tmp90816 := Call(__e, tmp90805, symmode, tmp90815)

			tmp90817 := Call(__e, tmp90804, tmp90816, Nil)

			tmp90818 := Call(__e, tmp90803, sym_h, tmp90817)

			tmp90819 := Call(__e, tmp90800, tmp90802, tmp90818)

			tmp90820 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90821 := Call(__e, tmp90820, sym_1, Nil)

			tmp90822 := Call(__e, tmp90799, tmp90819, tmp90821)

			__e.TailApply(tmp90798, symmode, tmp90822)
			return

		} else {
			__e.Return(V1812)
			return
		}

	}, 1)

	tmp90857 := Call(__e, PrimNS1Value(symns2_1set), symshen_4mode_1ify, tmp90796)

	_ = tmp90857

	tmp90858 := MakeNative(func(__e *ControlFlow) {
		V1816 := __e.Get(1)
		_ = V1816
		V1817 := __e.Get(2)
		_ = V1817
		V1818 := __e.Get(3)
		_ = V1818
		tmp90859 := MakeNative(func(__e *ControlFlow) {
			Variables := __e.Get(1)
			_ = Variables
			tmp90860 := MakeNative(func(__e *ControlFlow) {
				Predicates := __e.Get(1)
				_ = Predicates
				tmp90861 := MakeNative(func(__e *ControlFlow) {
					SearchLiterals := __e.Get(1)
					_ = SearchLiterals
					tmp90862 := MakeNative(func(__e *ControlFlow) {
						SearchClauses := __e.Get(1)
						_ = SearchClauses
						tmp90863 := MakeNative(func(__e *ControlFlow) {
							SideLiterals := __e.Get(1)
							_ = SideLiterals
							tmp90864 := MakeNative(func(__e *ControlFlow) {
								PremissLiterals := __e.Get(1)
								_ = PremissLiterals
								tmp90865 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								tmp90866 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								tmp90867 := Call(__e, tmp90866, SideLiterals, PremissLiterals)

								__e.TailApply(tmp90865, SearchLiterals, tmp90867)
								return

							}, 1)

							tmp90868 := Call(__e, PrimNS1Value(symns2_1value), symmap)

							tmp90869 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								tmp90870 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1premiss_1literal)

								tmp90871 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

								tmp90872 := Call(__e, tmp90871, V1818)

								__e.TailApply(tmp90870, X, tmp90872)
								return

							}, 1)

							tmp90873 := Call(__e, tmp90868, tmp90869, V1817)

							__e.TailApply(tmp90864, tmp90873)
							return

						}, 1)

						tmp90874 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

						tmp90875 := Call(__e, tmp90874, V1816)

						__e.TailApply(tmp90863, tmp90875)
						return

					}, 1)

					tmp90876 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clauses)

					tmp90877 := Call(__e, tmp90876, Predicates, V1818, Variables)

					__e.TailApply(tmp90862, tmp90877)
					return

				}, 1)

				tmp90878 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1literals)

				tmp90879 := Call(__e, tmp90878, Predicates, Variables, symContext__1957, symContext1__1957)

				__e.TailApply(tmp90861, tmp90879)
				return

			}, 1)

			tmp90880 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			tmp90881 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp90882 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				__e.TailApply(tmp90882, symshen_4cl)
				return

			}, 1)

			tmp90883 := Call(__e, tmp90880, tmp90881, V1818)

			__e.TailApply(tmp90860, tmp90883)
			return

		}, 1)

		tmp90884 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp90885 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp90886 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

			__e.TailApply(tmp90886, X)
			return

		}, 1)

		tmp90887 := Call(__e, tmp90884, tmp90885, V1818)

		__e.TailApply(tmp90859, tmp90887)
		return

	}, 3)

	tmp90888 := Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1body, tmp90858)

	_ = tmp90888

	tmp90889 := MakeNative(func(__e *ControlFlow) {
		V1827 := __e.Get(1)
		_ = V1827
		V1828 := __e.Get(2)
		_ = V1828
		V1829 := __e.Get(3)
		_ = V1829
		V1830 := __e.Get(4)
		_ = V1830
		tmp90896 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90897 := Call(__e, tmp90896, Nil, V1827)

		var ifres90892 Obj

		if True == tmp90897 {
			tmp90894 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90895 := Call(__e, tmp90894, Nil, V1828)

			var ifres90893 Obj

			if True == tmp90895 {
				ifres90893 = True

			} else {
				ifres90893 = False

			}

			ifres90892 = ifres90893

		} else {
			ifres90892 = False

		}

		if True == ifres90892 {
			__e.Return(Nil)
			return
		} else {
			tmp90891 := Call(__e, PrimNS1Value(symns2_1value), symshen_4csl_1help)

			__e.TailApply(tmp90891, V1827, V1828, V1829, V1830)
			return

		}

	}, 4)

	tmp90898 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1literals, tmp90889)

	_ = tmp90898

	tmp90899 := MakeNative(func(__e *ControlFlow) {
		V1837 := __e.Get(1)
		_ = V1837
		V1838 := __e.Get(2)
		_ = V1838
		V1839 := __e.Get(3)
		_ = V1839
		V1840 := __e.Get(4)
		_ = V1840
		tmp90939 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90940 := Call(__e, tmp90939, Nil, V1837)

		var ifres90935 Obj

		if True == tmp90940 {
			tmp90937 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90938 := Call(__e, tmp90937, Nil, V1838)

			var ifres90936 Obj

			if True == tmp90938 {
				ifres90936 = True

			} else {
				ifres90936 = False

			}

			ifres90935 = ifres90936

		} else {
			ifres90935 = False

		}

		if True == ifres90935 {
			tmp90928 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90929 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90930 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90931 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp90932 := Call(__e, tmp90931, V1839, Nil)

			tmp90933 := Call(__e, tmp90930, symContextOut__1957, tmp90932)

			tmp90934 := Call(__e, tmp90929, symbind, tmp90933)

			__e.TailApply(tmp90928, tmp90934, Nil)
			return

		} else {
			tmp90926 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90927 := Call(__e, tmp90926, V1837)

			var ifres90922 Obj

			if True == tmp90927 {
				tmp90924 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90925 := Call(__e, tmp90924, V1838)

				var ifres90923 Obj

				if True == tmp90925 {
					ifres90923 = True

				} else {
					ifres90923 = False

				}

				ifres90922 = ifres90923

			} else {
				ifres90922 = False

			}

			if True == ifres90922 {
				tmp90903 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90904 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90905 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90906 := Call(__e, tmp90905, V1837)

				tmp90907 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90908 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp90909 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90910 := Call(__e, tmp90909, V1838)

				tmp90911 := Call(__e, tmp90908, V1840, tmp90910)

				tmp90912 := Call(__e, tmp90907, V1839, tmp90911)

				tmp90913 := Call(__e, tmp90904, tmp90906, tmp90912)

				tmp90914 := Call(__e, PrimNS1Value(symns2_1value), symshen_4csl_1help)

				tmp90915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90916 := Call(__e, tmp90915, V1837)

				tmp90917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90918 := Call(__e, tmp90917, V1838)

				tmp90919 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				tmp90920 := Call(__e, tmp90919, symContext)

				tmp90921 := Call(__e, tmp90914, tmp90916, tmp90918, V1840, tmp90920)

				__e.TailApply(tmp90903, tmp90913, tmp90921)
				return

			} else {
				tmp90902 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp90902, symshen_4csl_1help)
				return

			}

		}

	}, 4)

	tmp90941 := Call(__e, PrimNS1Value(symns2_1set), symshen_4csl_1help, tmp90899)

	_ = tmp90941

	tmp90942 := MakeNative(func(__e *ControlFlow) {
		V1844 := __e.Get(1)
		_ = V1844
		V1845 := __e.Get(2)
		_ = V1845
		V1846 := __e.Get(3)
		_ = V1846
		tmp90979 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp90980 := Call(__e, tmp90979, Nil, V1844)

		var ifres90971 Obj

		if True == tmp90980 {
			tmp90977 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp90978 := Call(__e, tmp90977, Nil, V1845)

			var ifres90973 Obj

			if True == tmp90978 {
				tmp90975 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp90976 := Call(__e, tmp90975, Nil, V1846)

				var ifres90974 Obj

				if True == tmp90976 {
					ifres90974 = True

				} else {
					ifres90974 = False

				}

				ifres90973 = ifres90974

			} else {
				ifres90973 = False

			}

			var ifres90972 Obj

			if True == ifres90973 {
				ifres90972 = True

			} else {
				ifres90972 = False

			}

			ifres90971 = ifres90972

		} else {
			ifres90971 = False

		}

		if True == ifres90971 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp90969 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp90970 := Call(__e, tmp90969, V1844)

			var ifres90961 Obj

			if True == tmp90970 {
				tmp90967 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp90968 := Call(__e, tmp90967, V1845)

				var ifres90963 Obj

				if True == tmp90968 {
					tmp90965 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp90966 := Call(__e, tmp90965, V1846)

					var ifres90964 Obj

					if True == tmp90966 {
						ifres90964 = True

					} else {
						ifres90964 = False

					}

					ifres90963 = ifres90964

				} else {
					ifres90963 = False

				}

				var ifres90962 Obj

				if True == ifres90963 {
					ifres90962 = True

				} else {
					ifres90962 = False

				}

				ifres90961 = ifres90962

			} else {
				ifres90961 = False

			}

			if True == ifres90961 {
				tmp90946 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clause)

				tmp90947 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90948 := Call(__e, tmp90947, V1844)

				tmp90949 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90950 := Call(__e, tmp90949, V1845)

				tmp90951 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp90952 := Call(__e, tmp90951, V1846)

				tmp90953 := Call(__e, tmp90946, tmp90948, tmp90950, tmp90952)

				_ = tmp90953

				tmp90954 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clauses)

				tmp90955 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90956 := Call(__e, tmp90955, V1844)

				tmp90957 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90958 := Call(__e, tmp90957, V1845)

				tmp90959 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp90960 := Call(__e, tmp90959, V1846)

				__e.TailApply(tmp90954, tmp90956, tmp90958, tmp90960)
				return

			} else {
				tmp90945 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp90945, symshen_4construct_1search_1clauses)
				return

			}

		}

	}, 3)

	tmp90981 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clauses, tmp90942)

	_ = tmp90981

	tmp90982 := MakeNative(func(__e *ControlFlow) {
		V1850 := __e.Get(1)
		_ = V1850
		V1851 := __e.Get(2)
		_ = V1851
		V1852 := __e.Get(3)
		_ = V1852
		tmp90983 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog)

		tmp90984 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90985 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1base_1search_1clause)

		tmp90986 := Call(__e, tmp90985, V1850, V1851, V1852)

		tmp90987 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90988 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1recursive_1search_1clause)

		tmp90989 := Call(__e, tmp90988, V1850, V1851, V1852)

		tmp90990 := Call(__e, tmp90987, tmp90989, Nil)

		tmp90991 := Call(__e, tmp90984, tmp90986, tmp90990)

		__e.TailApply(tmp90983, tmp90991)
		return

	}, 3)

	tmp90992 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clause, tmp90982)

	_ = tmp90992

	tmp90993 := MakeNative(func(__e *ControlFlow) {
		V1856 := __e.Get(1)
		_ = V1856
		V1857 := __e.Get(2)
		_ = V1857
		V1858 := __e.Get(3)
		_ = V1858
		tmp90994 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90995 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90996 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90997 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp90998 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mode_1ify)

		tmp90999 := Call(__e, tmp90998, V1857)

		tmp91000 := Call(__e, tmp90997, tmp90999, symIn__1957)

		tmp91001 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91002 := Call(__e, tmp91001, symIn__1957, V1858)

		tmp91003 := Call(__e, tmp90996, tmp91000, tmp91002)

		tmp91004 := Call(__e, tmp90995, V1856, tmp91003)

		tmp91005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91006 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91007 := Call(__e, tmp91006, Nil, Nil)

		tmp91008 := Call(__e, tmp91005, sym_h_1, tmp91007)

		__e.TailApply(tmp90994, tmp91004, tmp91008)
		return

	}, 3)

	tmp91009 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1base_1search_1clause, tmp90993)

	_ = tmp91009

	tmp91010 := MakeNative(func(__e *ControlFlow) {
		V1862 := __e.Get(1)
		_ = V1862
		V1863 := __e.Get(2)
		_ = V1863
		V1864 := __e.Get(3)
		_ = V1864
		tmp91011 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91012 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91013 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91014 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91015 := Call(__e, tmp91014, symAssumption__1957, symAssumptions__1957)

		tmp91016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91018 := Call(__e, tmp91017, symAssumption__1957, symOut__1957)

		tmp91019 := Call(__e, tmp91016, tmp91018, V1864)

		tmp91020 := Call(__e, tmp91013, tmp91015, tmp91019)

		tmp91021 := Call(__e, tmp91012, V1862, tmp91020)

		tmp91022 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91023 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91024 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91025 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91026 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91027 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91028 := Call(__e, tmp91027, symOut__1957, V1864)

		tmp91029 := Call(__e, tmp91026, symAssumptions__1957, tmp91028)

		tmp91030 := Call(__e, tmp91025, V1862, tmp91029)

		tmp91031 := Call(__e, tmp91024, tmp91030, Nil)

		tmp91032 := Call(__e, tmp91023, tmp91031, Nil)

		tmp91033 := Call(__e, tmp91022, sym_h_1, tmp91032)

		__e.TailApply(tmp91011, tmp91021, tmp91033)
		return

	}, 3)

	tmp91034 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1recursive_1search_1clause, tmp91010)

	_ = tmp91034

	tmp91035 := MakeNative(func(__e *ControlFlow) {
		V1870 := __e.Get(1)
		_ = V1870
		tmp91148 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp91149 := Call(__e, tmp91148, Nil, V1870)

		if True == tmp91149 {
			__e.Return(Nil)
			return
		} else {
			tmp91146 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91147 := Call(__e, tmp91146, V1870)

			var ifres91114 Obj

			if True == tmp91147 {
				tmp91142 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91143 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91144 := Call(__e, tmp91143, V1870)

				tmp91145 := Call(__e, tmp91142, tmp91144)

				var ifres91116 Obj

				if True == tmp91145 {
					tmp91136 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					tmp91137 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91138 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91139 := Call(__e, tmp91138, V1870)

					tmp91140 := Call(__e, tmp91137, tmp91139)

					tmp91141 := Call(__e, tmp91136, symif, tmp91140)

					var ifres91118 Obj

					if True == tmp91141 {
						tmp91130 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91131 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91132 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91133 := Call(__e, tmp91132, V1870)

						tmp91134 := Call(__e, tmp91131, tmp91133)

						tmp91135 := Call(__e, tmp91130, tmp91134)

						var ifres91120 Obj

						if True == tmp91135 {
							tmp91122 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp91123 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91124 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91125 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp91126 := Call(__e, tmp91125, V1870)

							tmp91127 := Call(__e, tmp91124, tmp91126)

							tmp91128 := Call(__e, tmp91123, tmp91127)

							tmp91129 := Call(__e, tmp91122, Nil, tmp91128)

							var ifres91121 Obj

							if True == tmp91129 {
								ifres91121 = True

							} else {
								ifres91121 = False

							}

							ifres91120 = ifres91121

						} else {
							ifres91120 = False

						}

						var ifres91119 Obj

						if True == ifres91120 {
							ifres91119 = True

						} else {
							ifres91119 = False

						}

						ifres91118 = ifres91119

					} else {
						ifres91118 = False

					}

					var ifres91117 Obj

					if True == ifres91118 {
						ifres91117 = True

					} else {
						ifres91117 = False

					}

					ifres91116 = ifres91117

				} else {
					ifres91116 = False

				}

				var ifres91115 Obj

				if True == ifres91116 {
					ifres91115 = True

				} else {
					ifres91115 = False

				}

				ifres91114 = ifres91115

			} else {
				ifres91114 = False

			}

			if True == ifres91114 {
				tmp91103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91105 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91106 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91107 := Call(__e, tmp91106, V1870)

				tmp91108 := Call(__e, tmp91105, tmp91107)

				tmp91109 := Call(__e, tmp91104, symwhen, tmp91108)

				tmp91110 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

				tmp91111 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91112 := Call(__e, tmp91111, V1870)

				tmp91113 := Call(__e, tmp91110, tmp91112)

				__e.TailApply(tmp91103, tmp91109, tmp91113)
				return

			} else {
				tmp91101 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91102 := Call(__e, tmp91101, V1870)

				var ifres91057 Obj

				if True == tmp91102 {
					tmp91097 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91098 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91099 := Call(__e, tmp91098, V1870)

					tmp91100 := Call(__e, tmp91097, tmp91099)

					var ifres91059 Obj

					if True == tmp91100 {
						tmp91091 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						tmp91092 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91093 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91094 := Call(__e, tmp91093, V1870)

						tmp91095 := Call(__e, tmp91092, tmp91094)

						tmp91096 := Call(__e, tmp91091, symlet, tmp91095)

						var ifres91061 Obj

						if True == tmp91096 {
							tmp91085 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							tmp91086 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91087 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp91088 := Call(__e, tmp91087, V1870)

							tmp91089 := Call(__e, tmp91086, tmp91088)

							tmp91090 := Call(__e, tmp91085, tmp91089)

							var ifres91063 Obj

							if True == tmp91090 {
								tmp91077 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp91078 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp91079 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp91080 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp91081 := Call(__e, tmp91080, V1870)

								tmp91082 := Call(__e, tmp91079, tmp91081)

								tmp91083 := Call(__e, tmp91078, tmp91082)

								tmp91084 := Call(__e, tmp91077, tmp91083)

								var ifres91065 Obj

								if True == tmp91084 {
									tmp91067 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp91068 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91069 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91070 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91071 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									tmp91072 := Call(__e, tmp91071, V1870)

									tmp91073 := Call(__e, tmp91070, tmp91072)

									tmp91074 := Call(__e, tmp91069, tmp91073)

									tmp91075 := Call(__e, tmp91068, tmp91074)

									tmp91076 := Call(__e, tmp91067, Nil, tmp91075)

									var ifres91066 Obj

									if True == tmp91076 {
										ifres91066 = True

									} else {
										ifres91066 = False

									}

									ifres91065 = ifres91066

								} else {
									ifres91065 = False

								}

								var ifres91064 Obj

								if True == ifres91065 {
									ifres91064 = True

								} else {
									ifres91064 = False

								}

								ifres91063 = ifres91064

							} else {
								ifres91063 = False

							}

							var ifres91062 Obj

							if True == ifres91063 {
								ifres91062 = True

							} else {
								ifres91062 = False

							}

							ifres91061 = ifres91062

						} else {
							ifres91061 = False

						}

						var ifres91060 Obj

						if True == ifres91061 {
							ifres91060 = True

						} else {
							ifres91060 = False

						}

						ifres91059 = ifres91060

					} else {
						ifres91059 = False

					}

					var ifres91058 Obj

					if True == ifres91059 {
						ifres91058 = True

					} else {
						ifres91058 = False

					}

					ifres91057 = ifres91058

				} else {
					ifres91057 = False

				}

				if True == ifres91057 {
					tmp91046 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91047 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91048 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91049 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91050 := Call(__e, tmp91049, V1870)

					tmp91051 := Call(__e, tmp91048, tmp91050)

					tmp91052 := Call(__e, tmp91047, symis, tmp91051)

					tmp91053 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

					tmp91054 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91055 := Call(__e, tmp91054, V1870)

					tmp91056 := Call(__e, tmp91053, tmp91055)

					__e.TailApply(tmp91046, tmp91052, tmp91056)
					return

				} else {
					tmp91044 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91045 := Call(__e, tmp91044, V1870)

					if True == tmp91045 {
						tmp91041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

						tmp91042 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91043 := Call(__e, tmp91042, V1870)

						__e.TailApply(tmp91041, tmp91043)
						return

					} else {
						tmp91040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(tmp91040, symshen_4construct_1side_1literals)
						return

					}

				}

			}

		}

	}, 1)

	tmp91150 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1side_1literals, tmp91035)

	_ = tmp91150

	tmp91151 := MakeNative(func(__e *ControlFlow) {
		V1877 := __e.Get(1)
		_ = V1877
		V1878 := __e.Get(2)
		_ = V1878
		tmp91173 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		tmp91174 := Call(__e, tmp91173, V1877)

		if True == tmp91174 {
			tmp91160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91162 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			tmp91163 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			tmp91164 := Call(__e, tmp91163, V1877)

			tmp91165 := Call(__e, tmp91162, tmp91164)

			tmp91166 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91167 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1context)

			tmp91168 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			tmp91169 := Call(__e, tmp91168, V1877)

			tmp91170 := Call(__e, tmp91167, V1878, tmp91169)

			tmp91171 := Call(__e, tmp91166, tmp91170, Nil)

			tmp91172 := Call(__e, tmp91161, tmp91165, tmp91171)

			__e.TailApply(tmp91160, symshen_4t_d, tmp91172)
			return

		} else {
			tmp91158 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91159 := Call(__e, tmp91158, sym_b, V1877)

			if True == tmp91159 {
				tmp91155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp91157 := Call(__e, tmp91156, symThrowcontrol, Nil)

				__e.TailApply(tmp91155, symcut, tmp91157)
				return

			} else {
				tmp91154 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp91154, symshen_4construct_1premiss_1literal)
				return

			}

		}

	}, 2)

	tmp91175 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1premiss_1literal, tmp91151)

	_ = tmp91175

	tmp91176 := MakeNative(func(__e *ControlFlow) {
		V1881 := __e.Get(1)
		_ = V1881
		V1882 := __e.Get(2)
		_ = V1882
		tmp91206 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp91207 := Call(__e, tmp91206, True, V1881)

		var ifres91202 Obj

		if True == tmp91207 {
			tmp91204 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91205 := Call(__e, tmp91204, Nil, V1882)

			var ifres91203 Obj

			if True == tmp91205 {
				ifres91203 = True

			} else {
				ifres91203 = False

			}

			ifres91202 = ifres91203

		} else {
			ifres91202 = False

		}

		if True == ifres91202 {
			__e.Return(symContext__1957)
			return
		} else {
			tmp91200 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91201 := Call(__e, tmp91200, False, V1881)

			var ifres91196 Obj

			if True == tmp91201 {
				tmp91198 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91199 := Call(__e, tmp91198, Nil, V1882)

				var ifres91197 Obj

				if True == tmp91199 {
					ifres91197 = True

				} else {
					ifres91197 = False

				}

				ifres91196 = ifres91197

			} else {
				ifres91196 = False

			}

			if True == ifres91196 {
				__e.Return(symContextOut__1957)
				return
			} else {
				tmp91194 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91195 := Call(__e, tmp91194, V1882)

				if True == tmp91195 {
					tmp91181 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91182 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91183 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

					tmp91184 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp91185 := Call(__e, tmp91184, V1882)

					tmp91186 := Call(__e, tmp91183, tmp91185)

					tmp91187 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp91188 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1context)

					tmp91189 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91190 := Call(__e, tmp91189, V1882)

					tmp91191 := Call(__e, tmp91188, V1881, tmp91190)

					tmp91192 := Call(__e, tmp91187, tmp91191, Nil)

					tmp91193 := Call(__e, tmp91182, tmp91186, tmp91192)

					__e.TailApply(tmp91181, symcons, tmp91193)
					return

				} else {
					tmp91180 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(tmp91180, symshen_4construct_1context)
					return

				}

			}

		}

	}, 2)

	tmp91208 := Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1context, tmp91176)

	_ = tmp91208

	tmp91209 := MakeNative(func(__e *ControlFlow) {
		V1884 := __e.Get(1)
		_ = V1884
		tmp91224 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91225 := Call(__e, tmp91224, V1884)

		if True == tmp91225 {
			tmp91211 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			tmp91214 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91215 := Call(__e, tmp91214, V1884)

			tmp91216 := Call(__e, tmp91213, tmp91215)

			tmp91217 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91218 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			tmp91219 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91220 := Call(__e, tmp91219, V1884)

			tmp91221 := Call(__e, tmp91218, tmp91220)

			tmp91222 := Call(__e, tmp91217, tmp91221, Nil)

			tmp91223 := Call(__e, tmp91212, tmp91216, tmp91222)

			__e.TailApply(tmp91211, symcons, tmp91223)
			return

		} else {
			__e.Return(V1884)
			return
		}

	}, 1)

	tmp91226 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__cons__form, tmp91209)

	_ = tmp91226

	tmp91227 := MakeNative(func(__e *ControlFlow) {
		V1886 := __e.Get(1)
		_ = V1886
		tmp91228 := Call(__e, PrimNS1Value(symns2_1value), symshen_4preclude_1h)

		tmp91229 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp91230 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp91231 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(tmp91231, X)
			return

		}, 1)

		tmp91232 := Call(__e, tmp91229, tmp91230, V1886)

		__e.TailApply(tmp91228, tmp91232)
		return

	}, 1)

	tmp91233 := Call(__e, PrimNS1Value(symns2_1set), sympreclude, tmp91227)

	_ = tmp91233

	tmp91234 := MakeNative(func(__e *ControlFlow) {
		V1888 := __e.Get(1)
		_ = V1888
		tmp91235 := MakeNative(func(__e *ControlFlow) {
			FilterDatatypes := __e.Get(1)
			_ = FilterDatatypes
			tmp91236 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(tmp91236, symshen_4_ddatatypes_d)
			return

		}, 1)

		tmp91237 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp91238 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		tmp91239 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91240 := Call(__e, tmp91239, symshen_4_ddatatypes_d)

		tmp91241 := Call(__e, tmp91238, tmp91240, V1888)

		tmp91242 := Call(__e, tmp91237, symshen_4_ddatatypes_d, tmp91241)

		__e.TailApply(tmp91235, tmp91242)
		return

	}, 1)

	tmp91243 := Call(__e, PrimNS1Value(symns2_1set), symshen_4preclude_1h, tmp91234)

	_ = tmp91243

	tmp91244 := MakeNative(func(__e *ControlFlow) {
		V1890 := __e.Get(1)
		_ = V1890
		tmp91245 := Call(__e, PrimNS1Value(symns2_1value), symshen_4include_1h)

		tmp91246 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp91247 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp91248 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(tmp91248, X)
			return

		}, 1)

		tmp91249 := Call(__e, tmp91246, tmp91247, V1890)

		__e.TailApply(tmp91245, tmp91249)
		return

	}, 1)

	tmp91250 := Call(__e, PrimNS1Value(symns2_1set), syminclude, tmp91244)

	_ = tmp91250

	tmp91251 := MakeNative(func(__e *ControlFlow) {
		V1892 := __e.Get(1)
		_ = V1892
		tmp91252 := MakeNative(func(__e *ControlFlow) {
			ValidTypes := __e.Get(1)
			_ = ValidTypes
			tmp91253 := MakeNative(func(__e *ControlFlow) {
				NewDatatypes := __e.Get(1)
				_ = NewDatatypes
				tmp91254 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				__e.TailApply(tmp91254, symshen_4_ddatatypes_d)
				return

			}, 1)

			tmp91255 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp91256 := Call(__e, PrimNS1Value(symns2_1value), symunion)

			tmp91257 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91258 := Call(__e, tmp91257, symshen_4_ddatatypes_d)

			tmp91259 := Call(__e, tmp91256, ValidTypes, tmp91258)

			tmp91260 := Call(__e, tmp91255, symshen_4_ddatatypes_d, tmp91259)

			__e.TailApply(tmp91253, tmp91260)
			return

		}, 1)

		tmp91261 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

		tmp91262 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91263 := Call(__e, tmp91262, symshen_4_dalldatatypes_d)

		tmp91264 := Call(__e, tmp91261, V1892, tmp91263)

		__e.TailApply(tmp91252, tmp91264)
		return

	}, 1)

	tmp91265 := Call(__e, PrimNS1Value(symns2_1set), symshen_4include_1h, tmp91251)

	_ = tmp91265

	tmp91266 := MakeNative(func(__e *ControlFlow) {
		V1894 := __e.Get(1)
		_ = V1894
		tmp91267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4preclude_1h)

		tmp91268 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		tmp91269 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91270 := Call(__e, tmp91269, symshen_4_dalldatatypes_d)

		tmp91271 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp91272 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp91273 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(tmp91273, X)
			return

		}, 1)

		tmp91274 := Call(__e, tmp91271, tmp91272, V1894)

		tmp91275 := Call(__e, tmp91268, tmp91270, tmp91274)

		__e.TailApply(tmp91267, tmp91275)
		return

	}, 1)

	tmp91276 := Call(__e, PrimNS1Value(symns2_1set), sympreclude_1all_1but, tmp91266)

	_ = tmp91276

	tmp91277 := MakeNative(func(__e *ControlFlow) {
		V1896 := __e.Get(1)
		_ = V1896
		tmp91278 := Call(__e, PrimNS1Value(symns2_1value), symshen_4include_1h)

		tmp91279 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		tmp91280 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91281 := Call(__e, tmp91280, symshen_4_dalldatatypes_d)

		tmp91282 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		tmp91283 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			tmp91284 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(tmp91284, X)
			return

		}, 1)

		tmp91285 := Call(__e, tmp91282, tmp91283, V1896)

		tmp91286 := Call(__e, tmp91279, tmp91281, tmp91285)

		__e.TailApply(tmp91278, tmp91286)
		return

	}, 1)

	tmp91287 := Call(__e, PrimNS1Value(symns2_1set), syminclude_1all_1but, tmp91277)

	_ = tmp91287

	tmp91288 := MakeNative(func(__e *ControlFlow) {
		V1902 := __e.Get(1)
		_ = V1902
		tmp91347 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp91348 := Call(__e, tmp91347, Nil, V1902)

		if True == tmp91348 {
			tmp91338 := Call(__e, PrimNS1Value(symns2_1value), symshen_4update_1demodulation_1function)

			tmp91339 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91340 := Call(__e, tmp91339, symshen_4_dtc_d)

			tmp91341 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

			tmp91342 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp91343 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demod_1rule)

				__e.TailApply(tmp91343, X)
				return

			}, 1)

			tmp91344 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91345 := Call(__e, tmp91344, symshen_4_dsynonyms_d)

			tmp91346 := Call(__e, tmp91341, tmp91342, tmp91345)

			__e.TailApply(tmp91338, tmp91340, tmp91346)
			return

		} else {
			tmp91336 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91337 := Call(__e, tmp91336, V1902)

			var ifres91330 Obj

			if True == tmp91337 {
				tmp91332 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91333 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91334 := Call(__e, tmp91333, V1902)

				tmp91335 := Call(__e, tmp91332, tmp91334)

				var ifres91331 Obj

				if True == tmp91335 {
					ifres91331 = True

				} else {
					ifres91331 = False

				}

				ifres91330 = ifres91331

			} else {
				ifres91330 = False

			}

			if True == ifres91330 {
				tmp91292 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					tmp91316 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					tmp91317 := Call(__e, tmp91316, Vs)

					if True == tmp91317 {
						tmp91299 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pushnew)

						tmp91300 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp91301 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91302 := Call(__e, tmp91301, V1902)

						tmp91303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						tmp91304 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91305 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91306 := Call(__e, tmp91305, V1902)

						tmp91307 := Call(__e, tmp91304, tmp91306)

						tmp91308 := Call(__e, tmp91303, tmp91307, Nil)

						tmp91309 := Call(__e, tmp91300, tmp91302, tmp91308)

						tmp91310 := Call(__e, tmp91299, tmp91309, symshen_4_dsynonyms_d)

						_ = tmp91310

						tmp91311 := Call(__e, PrimNS1Value(symns2_1value), symshen_4synonyms_1help)

						tmp91312 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91313 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91314 := Call(__e, tmp91313, V1902)

						tmp91315 := Call(__e, tmp91312, tmp91314)

						__e.TailApply(tmp91311, tmp91315)
						return

					} else {
						tmp91294 := Call(__e, PrimNS1Value(symns2_1value), symshen_4free__variable__warnings)

						tmp91295 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91296 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91297 := Call(__e, tmp91296, V1902)

						tmp91298 := Call(__e, tmp91295, tmp91297)

						__e.TailApply(tmp91294, tmp91298, Vs)
						return

					}

				}, 1)

				tmp91318 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

				tmp91319 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				tmp91320 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91321 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91322 := Call(__e, tmp91321, V1902)

				tmp91323 := Call(__e, tmp91320, tmp91322)

				tmp91324 := Call(__e, tmp91319, tmp91323)

				tmp91325 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				tmp91326 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp91327 := Call(__e, tmp91326, V1902)

				tmp91328 := Call(__e, tmp91325, tmp91327)

				tmp91329 := Call(__e, tmp91318, tmp91324, tmp91328)

				__e.TailApply(tmp91292, tmp91329)
				return

			} else {
				tmp91291 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp91291, MakeString("odd number of synonyms\n"))
				return

			}

		}

	}, 1)

	tmp91349 := Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1help, tmp91288)

	_ = tmp91349

	tmp91350 := MakeNative(func(__e *ControlFlow) {
		V1905 := __e.Get(1)
		_ = V1905
		V1906 := __e.Get(2)
		_ = V1906
		tmp91358 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp91359 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp91360 := Call(__e, tmp91359, V1906)

		tmp91361 := Call(__e, tmp91358, V1905, tmp91360)

		if True == tmp91361 {
			tmp91357 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(tmp91357, V1906)
			return

		} else {
			tmp91352 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp91353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91354 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp91355 := Call(__e, tmp91354, V1906)

			tmp91356 := Call(__e, tmp91353, V1905, tmp91355)

			__e.TailApply(tmp91352, V1906, tmp91356)
			return

		}

	}, 2)

	tmp91362 := Call(__e, PrimNS1Value(symns2_1set), symshen_4pushnew, tmp91350)

	_ = tmp91362

	tmp91363 := MakeNative(func(__e *ControlFlow) {
		V1908 := __e.Get(1)
		_ = V1908
		tmp91395 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91396 := Call(__e, tmp91395, V1908)

		var ifres91381 Obj

		if True == tmp91396 {
			tmp91391 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp91392 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91393 := Call(__e, tmp91392, V1908)

			tmp91394 := Call(__e, tmp91391, tmp91393)

			var ifres91383 Obj

			if True == tmp91394 {
				tmp91385 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				tmp91386 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91387 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91388 := Call(__e, tmp91387, V1908)

				tmp91389 := Call(__e, tmp91386, tmp91388)

				tmp91390 := Call(__e, tmp91385, Nil, tmp91389)

				var ifres91384 Obj

				if True == tmp91390 {
					ifres91384 = True

				} else {
					ifres91384 = False

				}

				ifres91383 = ifres91384

			} else {
				ifres91383 = False

			}

			var ifres91382 Obj

			if True == ifres91383 {
				ifres91382 = True

			} else {
				ifres91382 = False

			}

			ifres91381 = ifres91382

		} else {
			ifres91381 = False

		}

		if True == ifres91381 {
			tmp91366 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91367 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp91368 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91369 := Call(__e, tmp91368, V1908)

			tmp91370 := Call(__e, tmp91367, tmp91369)

			tmp91371 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91372 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91373 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			tmp91374 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91375 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91376 := Call(__e, tmp91375, V1908)

			tmp91377 := Call(__e, tmp91374, tmp91376)

			tmp91378 := Call(__e, tmp91373, tmp91377)

			tmp91379 := Call(__e, tmp91372, tmp91378, Nil)

			tmp91380 := Call(__e, tmp91371, sym_1_6, tmp91379)

			__e.TailApply(tmp91366, tmp91370, tmp91380)
			return

		} else {
			tmp91365 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp91365, symshen_4demod_1rule)
			return

		}

	}, 1)

	tmp91397 := Call(__e, PrimNS1Value(symns2_1set), symshen_4demod_1rule, tmp91363)

	_ = tmp91397

	tmp91398 := MakeNative(func(__e *ControlFlow) {
		V1914 := __e.Get(1)
		_ = V1914
		tmp91484 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp91485 := Call(__e, tmp91484, V1914)

		var ifres91420 Obj

		if True == tmp91485 {
			tmp91480 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp91481 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91482 := Call(__e, tmp91481, V1914)

			tmp91483 := Call(__e, tmp91480, symdefun, tmp91482)

			var ifres91422 Obj

			if True == tmp91483 {
				tmp91476 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp91477 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp91478 := Call(__e, tmp91477, V1914)

				tmp91479 := Call(__e, tmp91476, tmp91478)

				var ifres91424 Obj

				if True == tmp91479 {
					tmp91470 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp91471 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91472 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp91473 := Call(__e, tmp91472, V1914)

					tmp91474 := Call(__e, tmp91471, tmp91473)

					tmp91475 := Call(__e, tmp91470, tmp91474)

					var ifres91426 Obj

					if True == tmp91475 {
						tmp91462 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp91463 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						tmp91464 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91465 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp91466 := Call(__e, tmp91465, V1914)

						tmp91467 := Call(__e, tmp91464, tmp91466)

						tmp91468 := Call(__e, tmp91463, tmp91467)

						tmp91469 := Call(__e, tmp91462, tmp91468)

						var ifres91428 Obj

						if True == tmp91469 {
							tmp91452 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp91453 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91454 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							tmp91455 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91456 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp91457 := Call(__e, tmp91456, V1914)

							tmp91458 := Call(__e, tmp91455, tmp91457)

							tmp91459 := Call(__e, tmp91454, tmp91458)

							tmp91460 := Call(__e, tmp91453, tmp91459)

							tmp91461 := Call(__e, tmp91452, Nil, tmp91460)

							var ifres91430 Obj

							if True == tmp91461 {
								tmp91444 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								tmp91445 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp91446 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp91447 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp91448 := Call(__e, tmp91447, V1914)

								tmp91449 := Call(__e, tmp91446, tmp91448)

								tmp91450 := Call(__e, tmp91445, tmp91449)

								tmp91451 := Call(__e, tmp91444, tmp91450)

								var ifres91432 Obj

								if True == tmp91451 {
									tmp91434 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									tmp91435 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91436 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91437 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91438 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									tmp91439 := Call(__e, tmp91438, V1914)

									tmp91440 := Call(__e, tmp91437, tmp91439)

									tmp91441 := Call(__e, tmp91436, tmp91440)

									tmp91442 := Call(__e, tmp91435, tmp91441)

									tmp91443 := Call(__e, tmp91434, Nil, tmp91442)

									var ifres91433 Obj

									if True == tmp91443 {
										ifres91433 = True

									} else {
										ifres91433 = False

									}

									ifres91432 = ifres91433

								} else {
									ifres91432 = False

								}

								var ifres91431 Obj

								if True == ifres91432 {
									ifres91431 = True

								} else {
									ifres91431 = False

								}

								ifres91430 = ifres91431

							} else {
								ifres91430 = False

							}

							var ifres91429 Obj

							if True == ifres91430 {
								ifres91429 = True

							} else {
								ifres91429 = False

							}

							ifres91428 = ifres91429

						} else {
							ifres91428 = False

						}

						var ifres91427 Obj

						if True == ifres91428 {
							ifres91427 = True

						} else {
							ifres91427 = False

						}

						ifres91426 = ifres91427

					} else {
						ifres91426 = False

					}

					var ifres91425 Obj

					if True == ifres91426 {
						ifres91425 = True

					} else {
						ifres91425 = False

					}

					ifres91424 = ifres91425

				} else {
					ifres91424 = False

				}

				var ifres91423 Obj

				if True == ifres91424 {
					ifres91423 = True

				} else {
					ifres91423 = False

				}

				ifres91422 = ifres91423

			} else {
				ifres91422 = False

			}

			var ifres91421 Obj

			if True == ifres91422 {
				ifres91421 = True

			} else {
				ifres91421 = False

			}

			ifres91420 = ifres91421

		} else {
			ifres91420 = False

		}

		if True == ifres91420 {
			tmp91401 := Call(__e, PrimNS1Value(symns2_1value), symeval)

			tmp91402 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91403 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp91404 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91405 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp91406 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91407 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91408 := Call(__e, tmp91407, V1914)

			tmp91409 := Call(__e, tmp91406, tmp91408)

			tmp91410 := Call(__e, tmp91405, tmp91409)

			tmp91411 := Call(__e, tmp91404, tmp91410)

			tmp91412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91413 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91414 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp91415 := Call(__e, tmp91414, V1914)

			tmp91416 := Call(__e, tmp91413, tmp91415)

			tmp91417 := Call(__e, tmp91412, tmp91416)

			tmp91418 := Call(__e, tmp91403, tmp91411, tmp91417)

			tmp91419 := Call(__e, tmp91402, sym_c_4, tmp91418)

			__e.TailApply(tmp91401, tmp91419)
			return

		} else {
			tmp91400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp91400, symshen_4lambda_1of_1defun)
			return

		}

	}, 1)

	tmp91486 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1of_1defun, tmp91398)

	_ = tmp91486

	tmp91487 := MakeNative(func(__e *ControlFlow) {
		V1917 := __e.Get(1)
		_ = V1917
		V1918 := __e.Get(2)
		_ = V1918
		tmp91488 := Call(__e, PrimNS1Value(symns2_1value), symtc)

		tmp91489 := Call(__e, tmp91488, sym_1)

		_ = tmp91489

		tmp91490 := Call(__e, PrimNS1Value(symns2_1value), symset)

		tmp91491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1of_1defun)

		tmp91492 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

		tmp91493 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91494 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91495 := Call(__e, PrimNS1Value(symns2_1value), symappend)

		tmp91496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default_1rule)

		tmp91497 := Call(__e, tmp91496)

		tmp91498 := Call(__e, tmp91495, V1918, tmp91497)

		tmp91499 := Call(__e, tmp91494, symshen_4demod, tmp91498)

		tmp91500 := Call(__e, tmp91493, symdefine, tmp91499)

		tmp91501 := Call(__e, tmp91492, tmp91500)

		tmp91502 := Call(__e, tmp91491, tmp91501)

		tmp91503 := Call(__e, tmp91490, symshen_4_ddemodulation_1function_d, tmp91502)

		_ = tmp91503

		var ifres91504 Obj

		if True == V1917 {
			tmp91505 := Call(__e, PrimNS1Value(symns2_1value), symtc)

			tmp91506 := Call(__e, tmp91505, sym_7)

			ifres91504 = tmp91506

		} else {
			ifres91504 = symshen_4skip

		}

		_ = ifres91504

		__e.Return(symsynonyms)
		return

	}, 2)

	tmp91507 := Call(__e, PrimNS1Value(symns2_1set), symshen_4update_1demodulation_1function, tmp91487)

	_ = tmp91507

	tmp91508 := MakeNative(func(__e *ControlFlow) {
		tmp91509 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91510 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91511 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp91512 := Call(__e, tmp91511, symX, Nil)

		tmp91513 := Call(__e, tmp91510, sym_1_6, tmp91512)

		__e.TailApply(tmp91509, symX, tmp91513)
		return

	}, 0)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4default_1rule, tmp91508)
	return

}, 0)
